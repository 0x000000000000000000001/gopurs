-- | Consuming workers for a closed, first-order forest language. These proofs
-- | belong to the final backend IR and never use source usage annotations.
module Gopurs.Ownership (prepare) where

import Prelude

import Control.Alternative (guard)
import Control.Monad.State (StateT, evalStateT, get, put)
import Control.Monad.Trans.Class (lift)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (all, any, foldMap, foldl)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe, isJust)
import Data.Newtype (unwrap)
import Data.Set as Set
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..), fst)
import Gopurs.CodegenState (CodegenMetadata, FunctionInfo)
import Gopurs.GoAst (GoDecl(..), GoExpr(..), GoType(..), rawGo, sanitizeName, goTypeToStr)
import Gopurs.GoTypes (exprTypeToGoType)
import Gopurs.PrimitiveExprs as Primitive
import PureScript.Backend.Optimizer.Convert (BackendModule)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Ident(..), ModuleName, Qualified(..))
import PureScript.Backend.Optimizer.FfiSupport (hashString)
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendAccessor(..), BackendOperator(..), BackendOperator1(..), BackendOperator2(..), BackendOperatorNum(..), BackendOperatorOrd(..), BackendSyntax(..), Level, Pair(..))
import PureScript.Backend.Optimizer.Syntax as Syn

type LocalRef = Tuple (Maybe Ident) Level

data Path = Path String (Array Int)
derive instance eqPath :: Eq Path
derive instance ordPath :: Ord Path

type Scalar = { expr :: GoExpr, goType :: GoType, reads :: Array Path }
data Value = Tree Path | Scalar Scalar
type Env = Map.Map LocalRef Value

type TreeSpec =
  { ty :: ExprType
  , goType :: GoType
  , constructor :: Qualified Ident
  , leaf :: Maybe (Qualified Ident)
  , fields :: Array GoType
  }

type Candidate =
  { original :: Ident
  , worker :: Ident
  , native :: String
  , consume :: String
  , spec :: TreeSpec
  , args :: Array LocalRef
  , argTypes :: Array GoType
  , body :: NeutralExpr
  }

type Candidates = Map.Map Ident Candidate
type Context = { metadata :: CodegenMetadata, moduleName :: ModuleName, candidate :: Candidate, candidates :: Candidates }

-- Tree leaves name disjoint subtrees. Scalar reads are borrowed and all are
-- materialized before any worker call or mutation in the same terminal term.
data TreeTerm
  = Keep Path
  | Empty
  | Construct (Array Argument)
  | Call Ident (Array Argument)
  | Existing GoExpr
data Argument = TreeArg TreeTerm | ScalarArg Scalar

type Generated = { stmts :: Array GoExpr, expr :: GoExpr }
type Gen = StateT Int Maybe

strip :: NeutralExpr -> BackendSyntax NeutralExpr
strip (NeutralExpr syn) = case syn of
  Typed _ inner -> strip inner
  Syn.TypeApp inner _ -> strip inner
  _ -> syn

annotation :: NeutralExpr -> Maybe ExprType
annotation (NeutralExpr syn) = case syn of
  Typed ty _ -> Just ty
  Syn.TypeApp inner _ -> annotation inner
  _ -> Nothing

arrow :: ExprType -> { args :: Array ExprType, result :: ExprType }
arrow (Func args result) = let next = arrow result in { args: args <> next.args, result: next.result }
arrow ty = { args: [], result: ty }

abstractions :: NeutralExpr -> { args :: Array LocalRef, body :: NeutralExpr }
abstractions expr = case strip expr of
  Abs args body -> let next = abstractions body in { args: NEA.toArray args <> next.args, body: next.body }
  UncurriedAbs args body | not (Array.null args) ->
    let next = abstractions body in { args: args <> next.args, body: next.body }
  _ -> { args: [], body: expr }

modulePrefix :: ModuleName -> String
modulePrefix = String.replaceAll (Pattern ".") (Replacement "_") <<< unwrap

qualify :: ModuleName -> Qualified Ident -> Qualified Ident
qualify current (Qualified mod ident) = Qualified (Just (fromMaybe current mod)) ident

enumTag :: CodegenMetadata -> ModuleName -> Qualified Ident -> Maybe GoExpr
enumTag metadata current name = do
  let Qualified mod (Ident ident) = qualify current name
      base = "Data_" <> modulePrefix (fromMaybe current mod) <> "_" <> sanitizeName ident
  guard (Set.member base metadata.enumCtors)
  pure (rawGo $ hashString base)

scalarType :: CodegenMetadata -> ModuleName -> ExprType -> Maybe GoType
scalarType metadata current ty = do
  let goType = exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors (modulePrefix current) ty
  guard case goType of
    TypeInt64 -> true
    TypeFloat64 -> true
    TypeBool -> true
    TypeString -> true
    TypeUint32 -> true
    _ -> false
  pure goType

treeSpecs :: CodegenMetadata -> BackendModule -> Array TreeSpec
treeSpecs metadata mod = Array.mapMaybe make mod.dataDecls
  where
  make decl = do
    guard (Array.null decl.vars)
    let fullName = unwrap mod.name <> "." <> decl.name
        ty = ADT fullName (String.split (Pattern ".") fullName) []
        nodes = Array.filter (not <<< Array.null <<< _.fields) decl.constructors
        nullary = Array.filter (Array.null <<< _.fields) decl.constructors
    node <- case nodes of
      [ node ] | Array.length nullary <= 1 -> Just node
      _ -> Nothing
    pointer <- case exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors (modulePrefix mod.name) ty of
      result@(TypeStructPointer _) -> Just result
      _ -> Nothing
    fields <- traverse (\field -> if field == ty then Just pointer else scalarType metadata mod.name field) node.fields
    guard (Array.elem pointer fields)
    pure
      { ty, goType: pointer, fields
      , constructor: Qualified (Just mod.name) (Ident node.name)
      , leaf: map (\leaf -> Qualified (Just mod.name) (Ident leaf.name)) (Array.head nullary)
      }

candidate :: CodegenMetadata -> BackendModule -> Array TreeSpec -> Tuple Ident NeutralExpr -> Maybe Candidate
candidate metadata mod specs (Tuple original@(Ident name) expr) = do
  signature <- arrow <$> annotation expr
  spec <- Array.find (\spec -> spec.ty == signature.result) specs
  let lambda = abstractions expr
  guard (not (Array.null lambda.args) && Array.length lambda.args == Array.length signature.args)
  argTypes <- traverse (\ty -> if ty == spec.ty then Just spec.goType else scalarType metadata mod.name ty) signature.args
  guard (Array.elem spec.goType argTypes)
  let worker = Ident ("__gopurs_owned_" <> name)
      native = "Call_" <> modulePrefix mod.name <> "_" <> sanitizeName (unwrap worker)
  pure { original, worker, native, consume: native <> "_consume", spec, args: lambda.args, argTypes, body: lambda.body }

reserveNames :: BackendModule -> Array Candidate -> Array Candidate
reserveNames mod functions = _.functions $ foldl assign { reserved, functions: [] } functions
  where
  names = map fst (Array.concatMap _.bindings mod.bindings)
    <> Array.fromFoldable (Map.keys mod.foreign)
  reserved = Set.fromFoldable $ map (\name -> "Call_" <> modulePrefix mod.name <> "_" <> sanitizeName (unwrap name)) names
  assign acc fn =
    let choose index =
          let worker = Ident ("__gopurs_owned_" <> unwrap fn.original <> "_" <> show index)
              native = "Call_" <> modulePrefix mod.name <> "_" <> sanitizeName (unwrap worker)
              consume = native <> "_consume"
          in if Set.member native acc.reserved || Set.member consume acc.reserved then choose (index + 1)
             else fn { worker = worker, native = native, consume = consume }
        renamed = choose 0
    in { reserved: Set.insert renamed.native $ Set.insert renamed.consume acc.reserved
       , functions: Array.snoc acc.functions renamed }

pathExpr :: Path -> GoExpr
pathExpr (Path root fields) = foldl (\value field -> GoStructAccess value ("V" <> show field)) (GoVar root) fields

prefix :: Path -> Path -> Boolean
prefix (Path a xs) (Path b ys) = a == b && Array.take (Array.length xs) ys == xs

overlap :: Path -> Path -> Boolean
overlap a b = prefix a b || prefix b a

prefixes :: Path -> Array Path
prefixes (Path root fields) =
  if Array.null fields then []
  else map (Path root <<< flip Array.take fields) (Array.range 0 (Array.length fields - 1))

appendField :: Path -> Int -> Path
appendField (Path root fields) index = Path root (Array.snoc fields index)

pathValue :: Context -> Env -> NeutralExpr -> Maybe Path
pathValue context env expr = case strip expr of
  Local name level -> case Map.lookup (Tuple name level) env of
    Just (Tree path) -> Just path
    _ -> Nothing
  Accessor base (GetCtorField ctor _ _ _ _ index) -> do
    guard (qualify context.moduleName ctor == context.candidate.spec.constructor)
    fieldType <- Array.index context.candidate.spec.fields index
    guard (fieldType == context.candidate.spec.goType)
    path <- pathValue context env base
    pure (appendField path index)
  _ -> Nothing

ordOperator :: BackendOperatorOrd -> String
ordOperator = case _ of
  OpEq -> "=="
  OpNotEq -> "!="
  OpLt -> "<"
  OpLte -> "<="
  OpGt -> ">"
  OpGte -> ">="

scalar :: Context -> Env -> NeutralExpr -> Maybe Scalar
scalar context env expr = case strip expr of
  Local name level -> case Map.lookup (Tuple name level) env of
    Just (Scalar value) -> Just value
    _ -> Nothing
  Lit lit -> do
    value <- Primitive.literal lit
    pure { expr: value.expr, goType: value.exprType, reads: [] }
  CtorSaturated ctor _ _ _ fields | Array.null fields -> do
    value <- enumTag context.metadata context.moduleName ctor
    pure { expr: value, goType: TypeUint32, reads: [] }
  Var ctor -> do
    value <- enumTag context.metadata context.moduleName ctor
    pure { expr: value, goType: TypeUint32, reads: [] }
  Accessor base (GetCtorField ctor _ _ _ _ index) -> do
    guard (qualify context.moduleName ctor == context.candidate.spec.constructor)
    goType <- Array.index context.candidate.spec.fields index
    guard (goType /= context.candidate.spec.goType)
    path <- pathValue context env base
    pure { expr: GoStructAccess (pathExpr path) ("V" <> show index), goType, reads: [ path ] }
  PrimOp (Op1 (OpIsTag ctor) value) -> case pathValue context env value of
    Just path -> do
      let qualified = qualify context.moduleName ctor
      operator <- if qualified == context.candidate.spec.constructor then Just "!="
        else if Just qualified == context.candidate.spec.leaf then Just "==" else Nothing
      pure { expr: GoBinOp operator (pathExpr path) (rawGo "nil"), goType: TypeBool, reads: [ path ] }
    Nothing -> do
      value' <- scalar context env value
      guard (value'.goType == TypeUint32)
      tag <- enumTag context.metadata context.moduleName ctor
      pure { expr: GoBinOp "==" value'.expr tag, goType: TypeBool, reads: value'.reads }
  PrimOp (Op1 op value) -> do
    value' <- scalar context env value
    operator <- case op, value'.goType of
      OpBooleanNot, TypeBool -> Just "!"
      OpIntNegate, TypeInt64 -> Just "-"
      OpNumberNegate, TypeFloat64 -> Just "-"
      _ , _ -> Nothing
    pure (value' { expr = GoPrefixOp operator value'.expr })
  PrimOp (Op2 op left right) -> do
    left' <- scalar context env left
    right' <- scalar context env right
    let binary symbol expected result = do
          guard (left'.goType == expected && right'.goType == expected)
          pure { expr: GoBinOp symbol left'.expr right'.expr, goType: result, reads: left'.reads <> right'.reads }
    case op of
      OpIntNum OpAdd -> binary "+" TypeInt64 TypeInt64
      OpIntNum OpSubtract -> binary "-" TypeInt64 TypeInt64
      OpIntNum OpMultiply -> binary "*" TypeInt64 TypeInt64
      OpIntOrd order -> binary (ordOperator order) TypeInt64 TypeBool
      OpNumberOrd order -> binary (ordOperator order) TypeFloat64 TypeBool
      OpStringOrd order -> binary (ordOperator order) TypeString TypeBool
      OpBooleanAnd -> binary "&&" TypeBool TypeBool
      OpBooleanOr -> binary "||" TypeBool TypeBool
      OpBooleanOrd OpEq -> binary "==" TypeBool TypeBool
      OpBooleanOrd OpNotEq -> binary "!=" TypeBool TypeBool
      _ -> Nothing
  _ -> Nothing

spine :: NeutralExpr -> { head :: NeutralExpr, args :: Array NeutralExpr }
spine expr = case strip expr of
  App fn args -> let inner = spine fn in inner { args = inner.args <> NEA.toArray args }
  UncurriedApp fn args -> let inner = spine fn in inner { args = inner.args <> args }
  _ -> { head: expr, args: [] }

knownCall :: Context -> NeutralExpr -> Maybe { fn :: Candidate, args :: Array NeutralExpr }
knownCall context expr = do
  let call = spine expr
  name <- case strip call.head of
    Var qualified -> case qualify context.moduleName qualified of
      Qualified (Just mod) name | mod == context.moduleName -> Just name
      _ -> Nothing
    _ -> Nothing
  fn <- Map.lookup name context.candidates
  guard (fn.spec.ty == context.candidate.spec.ty && Array.length call.args == Array.length fn.args)
  pure { fn, args: call.args }

treeTerm :: Context -> Env -> NeutralExpr -> Maybe TreeTerm
treeTerm context env expr = case pathValue context env expr of
  Just path -> Just (Keep path)
  Nothing -> case strip expr of
    CtorSaturated ctor _ _ _ fields -> do
      let qualified = qualify context.moduleName ctor
      if Just qualified == context.candidate.spec.leaf && Array.null fields then pure Empty
      else do
        guard (qualified == context.candidate.spec.constructor)
        guard (Array.length fields == Array.length context.candidate.spec.fields)
        Construct <$> traverse (\(Tuple ty (Tuple _ value)) -> argument context env ty value)
          (Array.zip context.candidate.spec.fields fields)
    Var ctor | Just (qualify context.moduleName ctor) == context.candidate.spec.leaf -> Just Empty
    _ -> do
      call <- knownCall context expr
      Call call.fn.original <$> traverse (\(Tuple ty value) -> argument context env ty value)
        (Array.zip call.fn.argTypes call.args)

argument :: Context -> Env -> GoType -> NeutralExpr -> Maybe Argument
argument context env expected expr =
  if expected == context.candidate.spec.goType then TreeArg <$> treeTerm context env expr
  else do
    value <- scalar context env expr
    guard (value.goType == expected)
    pure (ScalarArg value)

leaves :: TreeTerm -> Array Path
leaves = case _ of
  Keep path -> [ path ]
  Construct args -> foldMap argumentLeaves args
  Call _ args -> foldMap argumentLeaves args
  _ -> []
  where
  argumentLeaves (TreeArg value) = leaves value
  argumentLeaves _ = []

disjoint :: Array Path -> Boolean
disjoint paths = all identity (Array.mapWithIndex
  (\index path -> not (any (overlap path) (Array.drop (index + 1) paths))) paths)

-- Conservative continuation check. A projection through an old root counts as
-- a use of that root too; rejecting it loses opportunities, never aliases.
continuationPaths :: Env -> NeutralExpr -> Array Path
continuationPaths env (NeutralExpr syn) = case syn of
  Local name level -> case Map.lookup (Tuple name level) env of
    Just (Tree path) -> [ path ]
    Just (Scalar value) -> value.reads
    _ -> []
  _ -> foldMap (continuationPaths env) syn

freshName :: String -> Gen String
freshName prefix_ = do
  next <- get
  put (next + 1)
  pure (prefix_ <> show next)

snapshot :: TreeTerm -> Gen { stmts :: Array GoExpr, term :: TreeTerm }
snapshot = case _ of
  Keep path -> do
    name <- freshName "__read_"
    pure { stmts: [ GoAssign name (pathExpr path) ], term: Existing (GoVar name) }
  Construct args -> do
    result <- traverse snapshotArgument args
    pure { stmts: foldMap _.stmts result, term: Construct (map _.arg result) }
  Call name args -> do
    result <- traverse snapshotArgument args
    pure { stmts: foldMap _.stmts result, term: Call name (map _.arg result) }
  term -> pure { stmts: [], term }

snapshotArgument :: Argument -> Gen { stmts :: Array GoExpr, arg :: Argument }
snapshotArgument = case _ of
  TreeArg value -> do
    result <- snapshot value
    pure { stmts: result.stmts, arg: TreeArg result.term }
  ScalarArg value -> do
    name <- freshName "__scalar_"
    pure { stmts: [ GoAssign name (GoCall (GoVar (goTypeToStr value.goType)) [ value.expr ]) ]
         , arg: ScalarArg (value { expr = GoVar name, reads = [] }) }

-- A finite, invocation-local collection of dead cells. The slots are consumed
-- by clearing them. Nil donors need no allocation until a constructor needs it.
takeCell :: Array String -> Gen Generated
takeCell slots = do
  name <- freshName "__cell_"
  let takeOne slot rest = [ GoIfElse (GoBinOp "!=" (GoVar slot) (rawGo "nil"))
        [ GoMutate name (GoVar slot), GoMutate slot (rawGo "nil") ] rest ]
  pure { stmts: [ rawGo ("var " <> name <> " " <> "__TREE_TYPE__") ] <> Array.foldr takeOne [] slots
       , expr: GoVar name }

-- Substitute only this internal declaration placeholder, never user syntax.
cellStatements :: TreeSpec -> Array GoExpr -> Array GoExpr
cellStatements spec = map replace
  where
  replace (GoRaw code) = rawGo $ String.replaceAll (Pattern "__TREE_TYPE__") (Replacement $ goTypeToStr spec.goType) code.text
  replace expr = expr

emitTree :: Context -> Array String -> Boolean -> TreeTerm -> Gen Generated
emitTree context pool outer = case _ of
  Existing expr -> pure { stmts: [], expr }
  Empty -> pure { stmts: [], expr: rawGo "nil" }
  Keep _ -> lift Nothing
  Construct args -> do
    values <- traverse (emitArgument context pool) args
    cell <- takeCell pool
    name <- case cell.expr of
      GoVar name -> pure name
      _ -> lift Nothing
    pointer <- case context.candidate.spec.goType of
      TypeStructPointer pointer -> pure pointer
      _ -> lift Nothing
    let initialize = GoIfElse (GoBinOp "==" cell.expr (rawGo "nil"))
          [ GoMutate name (GoCall (GoVar "new") [ GoVar pointer.fullPath ]) ] []
        assign = Array.mapWithIndex (\index value -> GoMutate (name <> ".V" <> show index) value.expr) values
    pure { stmts: foldMap _.stmts values <> cellStatements context.candidate.spec cell.stmts
             <> [ initialize, GoMutate (name <> ".Rc") (GoInt 1) ] <> assign
         , expr: cell.expr }
  Call name args -> do
    fn <- lift $ Map.lookup name context.candidates
    values <- traverse (emitArgument context pool) args
    donor <- if outer then takeCell pool else pure { stmts: [], expr: rawGo "nil" }
    result <- freshName "__result_"
    pure { stmts: foldMap _.stmts values <> cellStatements context.candidate.spec donor.stmts
             <> [ GoAssign result (GoCall (GoVar fn.consume) (map _.expr values <> [ donor.expr ])) ]
         , expr: GoVar result }

emitArgument :: Context -> Array String -> Argument -> Gen Generated
emitArgument context pool = case _ of
  TreeArg tree -> emitTree context pool false tree
  ScalarArg value -> pure { stmts: [], expr: value.expr }

plan :: Context -> Env -> Array Path -> TreeTerm -> Gen
  { stmts :: Array GoExpr, term :: TreeTerm, pool :: Array String, retired :: Array Path }
plan _ env future term = do
  lift $ guard (disjoint $ leaves term)
  captured <- snapshot term
  let retained = leaves term
      roots = Array.mapMaybe (case _ of
        Tree (Path root _) -> Just (Path root [])
        _ -> Nothing) (Array.fromFoldable $ Map.values env)
      -- Only these paths were necessarily evaluated by the Keep snapshots.
      -- Scalar short-circuit expressions may leave deeper paths unevaluated.
      candidates = Array.nub $ roots <> foldMap prefixes retained
      dead = Array.filter (\path -> not (any (\kept -> prefix kept path) retained)
        && not (any (overlap path) future)) candidates
  slots <- traverse (\path -> do
    name <- freshName "__dead_"
    pure { name, stmt: GoAssign name (pathExpr path) }) dead
  donor <- freshName "__donor_slot_"
  pure { stmts: captured.stmts <> [ GoAssign donor (GoVar "__donor") ] <> map _.stmt slots
       , term: captured.term, pool: [ donor ] <> map _.name slots, retired: retained <> dead }

emitBody :: Context -> Env -> NeutralExpr -> Gen (Array GoExpr)
emitBody context env expr = case strip expr of
  Branch branches fallback -> do
    def <- emitBody context env fallback
    cases <- traverse (\(Pair condition body) -> do
      cond <- lift $ scalar context env condition
      lift $ guard (cond.goType == TypeBool)
      result <- emitBody context env body
      pure { condition: cond.expr, body: result }) (NEA.toArray branches)
    pure $ Array.foldr (\branch rest -> [ GoIfElse branch.condition branch.body rest ]) def cases
  Fail message -> pure [ GoCall (GoVar "panic") [ GoString message ] ]
  Let name level binding body -> case pathValue context env binding of
    Just path -> emitBody context (Map.insert (Tuple name level) (Tree path) env) body
    Nothing -> case scalar context env binding of
      Just value -> do
        temporary <- freshName "__let_scalar_"
        rest <- emitBody context (Map.insert (Tuple name level)
          (Scalar (value { expr = GoVar temporary, reads = [] })) env) body
        pure ([ GoAssign temporary (GoCall (GoVar (goTypeToStr value.goType)) [ value.expr ]) ] <> rest)
      Nothing -> do
        term <- lift $ treeTerm context env binding
        let consumed = leaves term
            future = continuationPaths env body
        lift $ guard (not (any (\path -> any (overlap path) future) consumed))
        prepared <- plan context env future term
        result <- emitTree context prepared.pool true prepared.term
        temporary <- freshName "__let_tree_"
        let remaining = Map.filter (case _ of
              Tree path -> not (any (overlap path) prepared.retired)
              Scalar value -> not (any (\path -> any (overlap path) prepared.retired) value.reads)) env
            nextEnv = Map.insert (Tuple name level) (Tree $ Path temporary []) remaining
        rest <- emitBody context nextEnv body
        pure (prepared.stmts <> result.stmts
          <> [ rawGo ("var " <> temporary <> " " <> goTypeToStr context.candidate.spec.goType)
             , GoMutate temporary result.expr, GoMutate "_" (GoVar temporary)
             , GoMutate "__donor" (rawGo "nil") ] <> rest)
  _ -> do
    term <- lift $ treeTerm context env expr
    prepared <- plan context env [] term
    case prepared.term of
      Call name args | name == context.candidate.original -> do
        values <- traverse (emitArgument context prepared.pool) args
        donor <- takeCell prepared.pool
        let assign = Array.mapWithIndex (\index value -> GoMutate ("__arg" <> show index) value.expr) values
        pure (prepared.stmts <> foldMap _.stmts values <> cellStatements context.candidate.spec donor.stmts
          <> assign <> [ GoMutate "__donor" donor.expr, GoContinue "__owned_loop" ])
      _ -> do
        result <- emitTree context prepared.pool true prepared.term
        pure (prepared.stmts <> result.stmts <> [ GoReturn result.expr ])

workerDeclarations :: CodegenMetadata -> ModuleName -> Candidates -> Candidate -> Maybe (Array GoDecl)
workerDeclarations metadata moduleName candidates fn = do
  let context = { metadata, moduleName, candidates, candidate: fn }
      params = Array.mapWithIndex (\index ty -> Tuple ("__arg" <> show index) ty) fn.argTypes
      env = Map.fromFoldable $ Array.zipWith (\ref (Tuple name ty) -> Tuple ref
        (if ty == fn.spec.goType then Tree (Path name []) else Scalar { expr: GoVar name, goType: ty, reads: [] })) fn.args params
  body <- evalStateT (emitBody context env fn.body) 0
  pure
    [ GoFunctionDecl { name: fn.consume, params: params <> [ Tuple "__donor" fn.spec.goType ]
        , result: fn.spec.goType, body: GoFor "__owned_loop" body }
    , GoFunctionDecl { name: fn.native, params, result: fn.spec.goType
        , body: GoReturn (GoCall (GoVar fn.consume) (map (GoVar <<< fst) params <> [ rawGo "nil" ])) }
    ]

-- Removing a rejected function invalidates every caller which depended on its
-- consuming contract. The fixed point includes mutually recursive families.
validateCandidates :: CodegenMetadata -> ModuleName -> Candidates -> Candidates
validateCandidates metadata moduleName candidates =
  let accepted = Map.filter (isJust <<< workerDeclarations metadata moduleName candidates) candidates
  in if Map.size accepted == Map.size candidates then accepted else validateCandidates metadata moduleName accepted

freshTree :: Context -> NeutralExpr -> Boolean
freshTree context expr = case strip expr of
  CtorSaturated ctor _ _ _ fields ->
    let qualified = qualify context.moduleName ctor
    in if Just qualified == context.candidate.spec.leaf then Array.null fields
       else qualified == context.candidate.spec.constructor
         && Array.length fields == Array.length context.candidate.spec.fields
         && all (\(Tuple ty (Tuple _ value)) ->
              if ty == context.candidate.spec.goType then freshTree context value
              else isJust (scalar context Map.empty value)) (Array.zip context.candidate.spec.fields fields)
  Var ctor -> Just (qualify context.moduleName ctor) == context.candidate.spec.leaf
  _ -> false

rewrite :: CodegenMetadata -> ModuleName -> Candidates -> NeutralExpr -> NeutralExpr
rewrite metadata moduleName candidates original@(NeutralExpr syn) =
  let call = spine original
      target = case strip call.head of
        Var qualified -> case qualify moduleName qualified of
          Qualified (Just mod) name | mod == moduleName -> Map.lookup name candidates
          _ -> Nothing
        _ -> Nothing
      choose = do
        fn <- target
        guard (Array.length call.args == Array.length fn.args)
        let context = { metadata, moduleName, candidates, candidate: fn }
        guard (all (\(Tuple ty value) -> ty /= fn.spec.goType || freshTree context value)
          (Array.zip fn.argTypes call.args))
        args <- NEA.fromArray (map (rewrite metadata moduleName candidates) call.args)
        let signature = Func (map (\ty -> if ty == fn.spec.goType then fn.spec.ty else Any) fn.argTypes) fn.spec.ty
        pure $ NeutralExpr $ Typed fn.spec.ty $ NeutralExpr $ App
          (NeutralExpr $ Typed signature $ NeutralExpr $ Var $ Qualified (Just moduleName) fn.worker) args
  in fromMaybe (NeutralExpr $ map (rewrite metadata moduleName candidates) syn) choose

prepare :: CodegenMetadata -> BackendModule ->
  { module :: BackendModule, declarations :: Array GoDecl, functions :: Map.Map String FunctionInfo }
prepare metadata mod =
  let specs = treeSpecs metadata mod
      found = reserveNames mod $ Array.mapMaybe (candidate metadata mod specs) (Array.concatMap _.bindings mod.bindings)
      initial = Map.fromFoldable $ map (\fn -> Tuple fn.original fn) found
      accepted = validateCandidates metadata mod.name initial
      workers = Array.fromFoldable (Map.values accepted)
      declarations = foldMap (fromMaybe [] <<< workerDeclarations metadata mod.name accepted) workers
      functions = Map.fromFoldable $ map (\fn -> Tuple (sanitizeName $ unwrap fn.worker)
        { fullName: fn.native, fArgs: fn.argTypes, fRet: fn.spec.goType, arity: Array.length fn.args }) workers
      bindings = map (\group -> group { bindings = map (\(Tuple name expr) ->
        Tuple name (rewrite metadata mod.name accepted expr)) group.bindings }) mod.bindings
  in { module: mod { bindings = bindings }, declarations, functions }
