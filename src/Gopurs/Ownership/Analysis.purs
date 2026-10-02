module Gopurs.Ownership.Analysis
  ( strip
  , spine
  , modulePrefix
  , qualify
  , pathExpr
  , prefix
  , overlap
  , prefixes
  , pathValue
  , scalar
  , treeTerm
  , leaves
  , disjoint
  , continuationPaths
  , freshTree
  ) where

import Prelude

import Control.Alternative (guard)
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
import Data.Tuple (Tuple(..))
import Gopurs.CodegenState (CodegenMetadata)
import Gopurs.GoAst (GoExpr(..), GoType(..), rawGo, sanitizeName)
import Gopurs.Ownership.Types (Argument(..), Candidate, Context, Env, Path(..), Scalar, TreeTerm(..), Value(..))
import Gopurs.PrimitiveExprs as Primitive
import PureScript.Backend.Optimizer.CoreFn (Ident(..), ModuleName, Qualified(..))
import PureScript.Backend.Optimizer.FfiSupport (hashString)
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendAccessor(..), BackendOperator(..), BackendOperator1(..), BackendOperator2(..), BackendOperatorNum(..), BackendOperatorOrd(..), BackendSyntax(..))

-- Read only the final backend IR. Typed/TypeApp wrappers guide candidate
-- signatures elsewhere; neither wrapper certifies ownership of an expression.
strip :: NeutralExpr -> BackendSyntax NeutralExpr
strip (NeutralExpr syn) = case syn of
  Typed _ inner -> strip inner
  TypeApp inner _ -> strip inner
  _ -> syn

spine :: NeutralExpr -> { head :: NeutralExpr, args :: Array NeutralExpr }
spine expr = case strip expr of
  App fn args -> let inner = spine fn in inner { args = inner.args <> NEA.toArray args }
  UncurriedApp fn args -> let inner = spine fn in inner { args = inner.args <> args }
  _ -> { head: expr, args: [] }

modulePrefix :: ModuleName -> String
modulePrefix = String.replaceAll (Pattern ".") (Replacement "_") <<< unwrap

qualify :: ModuleName -> Qualified Ident -> Qualified Ident
qualify current (Qualified mod ident) = Qualified (Just (fromMaybe current mod)) ident

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

enumTag :: CodegenMetadata -> ModuleName -> Qualified Ident -> Maybe GoExpr
enumTag metadata current name = do
  let Qualified mod (Ident ident) = qualify current name
      base = "Data_" <> modulePrefix (fromMaybe current mod) <> "_" <> sanitizeName ident
  guard (Set.member base metadata.enumCtors)
  pure (rawGo $ hashString base)

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

-- A projection through an old root counts as a use of that root too. This
-- conservative continuation check may reject reuse but cannot excuse aliasing.
continuationPaths :: Env -> NeutralExpr -> Array Path
continuationPaths env (NeutralExpr syn) = case syn of
  Local name level -> case Map.lookup (Tuple name level) env of
    Just (Tree path) -> [ path ]
    Just (Scalar value) -> value.reads
    _ -> []
  _ -> foldMap (continuationPaths env) syn

-- Entry calls need closed constructor trees. A fresh parent containing a local
-- subtree is still borrowed; only nil leaves may be shared without a proof.
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
