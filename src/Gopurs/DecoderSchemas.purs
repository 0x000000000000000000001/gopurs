module Gopurs.DecoderSchemas (specializeDecoderSchemas) where

import Prelude hiding (one)

import Control.Alternative (guard)
import Control.Monad.State (get, put, runState)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (all, foldl)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe, isJust)
import Data.Newtype (unwrap)
import Data.Set as Set
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.String.CodeUnits as CU
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Gopurs.CodegenState (CodegenMetadata)
import Gopurs.GoAst (GoDecl(..), GoExpr(..), GoType(..), rawGo, sanitizeName)
import Gopurs.Printer (printGoExpr)
import PureScript.Backend.Optimizer.Convert (BackendModule)
import PureScript.Backend.Optimizer.CoreFn (Ident(..), Literal(..), ModuleName(..), Prop(..), Qualified(..))
import PureScript.Backend.Optimizer.CoreFn as CoreFn
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendAccessor(..), BackendOperator(..), BackendOperator1(..), BackendOperator2(..), BackendOperatorOrd(..), BackendSyntax(..), Level, Pair(..))

data Schema = Scalar String | Sequence Schema | Optional Schema | ObjectSchema (Array (Tuple String Schema)) | Custom | Derived (Program Schema)

-- A small, proven subset of a real custom decoder body. Every read, branch and
-- constructor comes from that body; result types and label names cannot invent
-- a variant decoder. Error continuations must forward the exact Left payload.
data Program a
  = ReadField Level String Boolean Boolean a (Program a)
  | Choices (Array (Tuple (Tuple Level String) (Program a))) (Program a)
  | ReturnValue NeutralExpr
  | TypeMismatch String

data Atom = Input | ObjectInput | Decoded Level | Failure Level | Result Boolean Atom

type Build = { next :: Int, names :: Set.Set String, sources :: Array (Tuple Ident NeutralExpr), declarations :: Array GoDecl }

-- Only actual standard dictionary applications justify a worker. Types alone
-- do not identify a decoder. Unknown methods remain leaves of the same native
-- plan, and nonconstant symbol dictionaries prevent record specialization.
specializeDecoderSchemas :: CodegenMetadata -> BackendModule -> { module :: BackendModule, declarations :: Array GoDecl }
specializeDecoderSchemas metadata mod
  | not (Map.member "Data.Argonaut.Decode.Internal.Record.schemaDecoderABI1" metadata.globalTypes) = { module: mod, declarations: [] }
  | otherwise =
      let
        definitions = Map.fromFoldable (Array.concatMap (\g -> if g.recursive then [] else g.bindings) mod.bindings)
        names = Set.fromFoldable (map (\(Tuple (Ident name) _) -> name) (Array.concatMap _.bindings mod.bindings))
        initial :: Build
        initial = { next: 0, names, sources: [], declarations: [] }
        Tuple bindings result = runState (traverse (rewriteGroup definitions) mod.bindings) initial
      in { module: mod { bindings = bindings <> if Array.null result.sources then [] else [ { recursive: false, bindings: result.sources } ] }, declarations: result.declarations }
  where
  prefix = String.replaceAll (Pattern ".") (Replacement "_") (unwrap mod.name)
  textEnabled = Map.member "Data.Argonaut.Decode.Parser.textDecoderABI1" metadata.globalTypes
  rewriteGroup definitions group
    | group.recursive = pure group
    | otherwise = do
        bindings <- traverse (\(Tuple ident expr) -> Tuple ident <$> rewrite definitions expr) group.bindings
        pure group { bindings = bindings }
  rewrite definitions original@(NeutralExpr syn) =
    case candidate definitions original of
      Just schema -> allocate schema original
      Nothing -> case syn of
        LetRec _ _ _ -> pure original
        _ -> NeutralExpr <$> traverse (rewrite definitions) syn
  candidate definitions expr = case strip expr of
    App fn args | standard "decodeJson" fn && closed expr -> case NEA.toArray args of
      [ dictionary ] -> do
        schema <- decoder mod.name definitions 64 dictionary
        if weight schema > 1 && weight schema <= 128 && tagged schema then Just schema else Nothing
      _ -> Nothing
    _ -> Nothing
  allocate schema original = do
    state <- get
    let
      name = "__json_schema_" <> show state.next
      source = name <> "_source"
      next = state { next = state.next + 1 }
    if Array.any (\existing -> existing == name || isJust (String.stripPrefix (Pattern (name <> "_")) existing)) (Set.toUnfoldable state.names) then put next *> allocate schema original
    else do
      let
        qualified = prefix <> "_" <> sanitizeName name
        worker = qualified <> "_decode"
        withText = textEnabled && textComplete schema
        generated = emit prefix worker false false schema <> if withText then emit prefix worker false true schema else []
        compiled = GoCall (GoVar "argonautCompileSchema")
          [ GoCall (GoVar ("Get_" <> prefix <> "_" <> sanitizeName source)) []
          , GoVar worker, GoVar (worker <> "_accepts")
          ]
        getter = GoCachedValue
          { identifier: qualified, goType: TypeValue
          , expression: if withText then GoCall (GoVar "argonautCompileTextSchema")
              [ compiled, GoVar (worker <> "_text"), GoVar (worker <> "_text_accepts") ] else compiled
          }
      put next
        { names = Set.insert source (Set.insert name state.names)
        , sources = state.sources <> [ Tuple (Ident source) (NeutralExpr (Typed CoreFn.Any original)) ] <> programSources prefix worker schema
        , declarations = state.declarations <> generated <> [ getter ]
        }
      pure (replace original (NeutralExpr (Var (Qualified (Just mod.name) (Ident name)))))

replace :: NeutralExpr -> NeutralExpr -> NeutralExpr
replace (NeutralExpr syn) replacement = case syn of
  Typed ty inner -> NeutralExpr (Typed ty (replace inner replacement))
  TypeApp inner ty -> NeutralExpr (TypeApp (replace inner replacement) ty)
  _ -> replacement

strip :: NeutralExpr -> BackendSyntax NeutralExpr
strip (NeutralExpr syn) = case syn of
  Typed _ inner -> strip inner
  TypeApp inner _ -> strip inner
  _ -> syn

closed :: NeutralExpr -> Boolean
closed (NeutralExpr syn) = case syn of
  Local _ _ -> false
  LetRec _ _ _ -> false
  PrimEffect _ -> false
  EffectBind _ _ _ _ -> false
  EffectPure _ -> false
  EffectDefer _ -> false
  UncurriedEffectAbs _ _ -> false
  UncurriedEffectApp _ _ -> false
  _ -> all closed syn

standard :: String -> NeutralExpr -> Boolean
standard name expr = case strip expr of
  Var (Qualified (Just (ModuleName "Data.Argonaut.Decode.Class")) (Ident found)) -> name == found
  _ -> false

-- Flatten only applications/aliases of nonrecursive module values. The original
-- decoder construction is preserved in its source getter, with argument order.
resolve :: ModuleName -> Map.Map Ident NeutralExpr -> Int -> NeutralExpr -> Maybe (Tuple NeutralExpr (Array NeutralExpr))
resolve owner definitions fuel expr
  | fuel <= 0 = Nothing
  | otherwise = case strip expr of
      Var (Qualified (Just moduleName) name) | moduleName == owner -> case Map.lookup name definitions of
        Just value -> resolve owner definitions (fuel - 1) value
        Nothing -> Just (Tuple expr [])
      App fn args -> do
        Tuple head prior <- resolve owner definitions (fuel - 1) fn
        pure (Tuple head (prior <> NEA.toArray args))
      _ -> Just (Tuple expr [])

decoder :: ModuleName -> Map.Map Ident NeutralExpr -> Int -> NeutralExpr -> Maybe Schema
decoder owner definitions fuel expr = do
  Tuple head args <- resolve owner definitions fuel expr
  let recur = decoder owner definitions (fuel - 1)
  case args of
    [] | standard "decodeJsonInt" head -> pure (Scalar "Int")
    [] | standard "decodeJsonNumber" head -> pure (Scalar "Number")
    [] | standard "decodeJsonString" head -> pure (Scalar "String")
    [] | standard "decodeJsonBoolean" head -> pure (Scalar "Boolean")
    [] | standard "decodeJsonJson" head -> pure (Scalar "Json")
    [ inner ] | standard "decodeArray" head -> Sequence <$> recur inner
    [ inner ] | standard "decodeJsonMaybe" head -> Optional <$> recur inner
    [ row, proxy ] | standard "decodeRecord" head && erased proxy -> ObjectSchema <$> record owner definitions (fuel - 1) row
    [] -> case strip expr of
      Var (Qualified (Just moduleName) ident) -> do
        guard (moduleName /= owner || Map.member ident definitions)
        guard (case strip head of
          Var (Qualified (Just resolvedOwner) resolvedName) | resolvedOwner == owner -> Map.member resolvedName definitions
          _ -> true)
        pure (fromMaybe Custom (Derived <$> deriveProgram owner definitions (fuel - 1) head))
      _ -> Nothing
    _ -> Nothing

record :: ModuleName -> Map.Map Ident NeutralExpr -> Int -> NeutralExpr -> Maybe (Array (Tuple String Schema))
record owner definitions fuel expr = do
  Tuple head args <- resolve owner definitions fuel expr
  case args of
    [] | standard "gDecodeJsonNil" head -> pure []
    [ field, tail, symbol, cons, lacks ] | standard "gDecodeJsonCons" head && erased cons && erased lacks -> do
      key <- symbolName owner definitions (fuel - 1) symbol
      Tuple fieldHead fieldArgs <- resolve owner definitions (fuel - 1) field
      schema <- case fieldArgs of
        [ inner ] | standard "decodeFieldId" fieldHead -> decoder owner definitions (fuel - 1) inner
        [ inner ] | standard "decodeFieldMaybe" fieldHead -> Optional <$> decoder owner definitions (fuel - 1) inner
        _ -> Nothing
      rest <- record owner definitions (fuel - 1) tail
      pure (Array.cons (Tuple key schema) rest)
    _ -> Nothing

erased :: NeutralExpr -> Boolean
erased expr = case strip expr of
  PrimUndefined -> true
  _ -> false

symbolName :: ModuleName -> Map.Map Ident NeutralExpr -> Int -> NeutralExpr -> Maybe String
symbolName owner definitions fuel expr = do
  Tuple head args <- resolve owner definitions fuel expr
  case args, strip head of
    [], Lit (LitRecord [ Prop "reflectSymbol" method ]) -> case strip method of
      Abs refs body | NEA.length refs == 1 -> case strip body of
        Lit (LitString value) -> Just value
        _ -> Nothing
      _ -> Nothing
    _, _ -> Nothing

weight :: Schema -> Int
weight = case _ of
  Scalar _ -> 1
  Custom -> 0
  Derived program -> 1 + programWeight program
  Sequence inner -> 1 + weight inner
  Optional inner -> 1 + weight inner
  ObjectSchema fields -> 1 + foldl (\n (Tuple _ child) -> n + weight child) 0 fields

tagged :: Schema -> Boolean
tagged = case _ of
  Custom -> false
  Derived _ -> false
  _ -> true

complete :: Schema -> Boolean
complete = case _ of
  Custom -> false
  Derived _ -> false
  Sequence child -> complete child
  Optional child -> complete child
  ObjectSchema fields -> all (\(Tuple _ child) -> complete child) fields
  _ -> true

-- An opaque callback may mutate or retain its Json argument. Parsing only a
-- subtree at its call site would not preserve aliases across repeated reads.
-- Keep the complete ordinary composition unless every operation is proven.
-- Derived programs already require complete schemas for all their reads.
textComplete :: Schema -> Boolean
textComplete = case _ of
  Custom -> false
  Sequence child -> textComplete child
  Optional child -> textComplete child
  ObjectSchema fields -> all (\(Tuple _ child) -> textComplete child) fields
  _ -> true

programWeight :: Program Schema -> Int
programWeight = case _ of
  ReadField _ _ _ _ schema next -> 1 + weight schema + programWeight next
  Choices cases other -> 1 + foldl (\n (Tuple _ branch) -> n + programWeight branch) (programWeight other) cases
  _ -> 1

isGlobal :: String -> String -> NeutralExpr -> Boolean
isGlobal owner name expr = case strip expr of
  Var (Qualified (Just (ModuleName m)) (Ident n)) -> m == owner && n == name
  _ -> false

isEither :: String -> Qualified Ident -> Boolean
isEither name = (_ == Qualified (Just (ModuleName "Data.Either")) (Ident name))

resultAtom :: Atom -> Maybe { success :: Boolean, payload :: Atom }
resultAtom = case _ of
  Result success payload -> Just { success, payload }
  _ -> Nothing

decodedAtom :: Atom -> Maybe Level
decodedAtom = case _ of
  Decoded level -> Just level
  _ -> Nothing

callExpr :: NeutralExpr -> Maybe { fn :: NeutralExpr, args :: Array NeutralExpr }
callExpr expr = case strip expr of
  App fn args -> Just { fn, args: NEA.toArray args }
  _ -> Nothing

stringExpr :: NeutralExpr -> Maybe String
stringExpr expr = case strip expr of
  Lit (LitString value) -> Just value
  _ -> Nothing

one :: forall a. Array a -> Maybe a
one = case _ of
  [ value ] -> Just value
  _ -> Nothing

two :: forall a. Array a -> Maybe { first :: a, second :: a }
two = case _ of
  [ first, second ] -> Just { first, second }
  _ -> Nothing

three :: forall a. Array a -> Maybe { first :: a, second :: a, third :: a }
three = case _ of
  [ first, second, third ] -> Just { first, second, third }
  _ -> Nothing

atom :: Map.Map Level Atom -> NeutralExpr -> Maybe Atom
atom env expr = case strip expr of
  Local _ level -> Map.lookup level env
  Accessor value (GetCtorField ctor _ _ _ "value0" 0) -> do
    { success, payload } <- atom env value >>= resultAtom
    guard (isEither (if success then "Right" else "Left") ctor)
    pure payload
  _ -> Nothing

knownBranch :: Map.Map Level Atom -> Array (Pair NeutralExpr) -> NeutralExpr -> Maybe NeutralExpr
knownBranch env cases other = case Array.uncons cases of
  Nothing -> Just other
  Just { head: Pair condition body, tail } -> case strip condition of
    PrimOp (Op1 (OpIsTag ctor) value) -> do
      { success } <- atom env value >>= resultAtom
      guard (isEither "Right" ctor || isEither "Left" ctor)
      if success == isEither "Right" ctor then Just body else knownBranch env tail other
    _ -> Nothing

forwards :: Int -> Level -> Map.Map Level Atom -> NeutralExpr -> Boolean
forwards fuel expected env expr
  | fuel <= 0 = false
  | otherwise = case strip expr of
      Let _ level value body -> case atom env value of
        Just bound -> forwards (fuel - 1) expected (Map.insert level bound env) body
        _ -> false
      Branch cases other -> case knownBranch env (NEA.toArray cases) other of
        Just body -> forwards (fuel - 1) expected env body
        _ -> false
      CtorSaturated ctor _ _ _ [ Tuple "value0" value ] | isEither "Left" ctor -> case atom env value of
        Just (Failure level) -> level == expected
        _ -> false
      _ -> case atom env expr of
        Just (Result false (Failure level)) -> level == expected
        _ -> false

-- Restrict successful construction to literals, decoded values and real ADT
-- constructors. The ordinary code generator compiles this expression, retaining
-- its annotations, native field layouts, reboxing and constructor identities.
valueExpr :: Map.Map Level Atom -> NeutralExpr -> Maybe NeutralExpr
valueExpr env original@(NeutralExpr syn) = case syn of
  Typed ty inner -> NeutralExpr <<< Typed ty <$> valueExpr env inner
  TypeApp inner ty -> NeutralExpr <<< flip TypeApp ty <$> valueExpr env inner
  Lit (LitString _) -> Just original
  Lit (LitInt _) -> Just original
  Lit (LitNumber _) -> Just original
  Lit (LitBoolean _) -> Just original
  CtorSaturated ctor ct tn cn fields -> NeutralExpr <<< CtorSaturated ctor ct tn cn <$> traverse (\(Tuple key value) -> Tuple key <$> valueExpr env value) fields
  _ -> do
    level <- atom env original >>= decodedAtom
    pure (NeutralExpr (Local Nothing level))

reader :: NeutralExpr -> Maybe (Tuple Boolean Boolean)
reader expr = case strip expr of
  Var (Qualified (Just (ModuleName "Data.Argonaut.Decode.Decoders")) (Ident name)) -> do
    base <- Array.find (\base -> name == base || case String.stripPrefix (Pattern (base <> "__")) name of
      Just suffix -> suffix /= "" && all (\c -> c >= '0' && c <= '9') (CU.toCharArray suffix)
      _ -> false) [ "getField", "getFieldOptional", "getFieldOptional'" ]
    pure (Tuple (base /= "getField") (base == "getFieldOptional'"))
  _ -> Nothing

deriveProgram :: ModuleName -> Map.Map Ident NeutralExpr -> Int -> NeutralExpr -> Maybe (Program Schema)
deriveProgram owner definitions fuel expr = do
  guard (fuel > 0)
  method <- case strip expr of
    Lit (LitRecord [ Prop "decodeJson" value ]) -> Just value
    _ -> Nothing
  { refs, body } <- case strip method of
    Abs refs body -> Just { refs, body }
    _ -> Nothing
  Tuple _ input <- one (NEA.toArray refs)
  { result, producer, continuation } <- case strip body of
    Let _ result producer continuation -> Just { result, producer, continuation }
    _ -> Nothing
  { fn: helper, args } <- callExpr producer
  guard (isGlobal "Data.Argonaut.Decode.Internal.Record" "borrowObject" helper)
  { first: objectMethod, second: json } <- two args
  input' <- case strip json of
    Local _ level -> Just level
    _ -> Nothing
  guard (input == input')
  dictionary <- case strip objectMethod of
    Accessor dictionary (GetProp "decodeJson") -> Just dictionary
    _ -> Nothing
  Tuple head arguments <- resolve owner definitions fuel dictionary
  guard (standard "decodeForeignObject" head)
  inner <- one arguments
  guard (standard "decodeJsonJson" inner)
  let env = Map.singleton input Input
  guard (forwards 128 result (Map.insert result (Result false (Failure result)) env) continuation)
  lower 256 (Map.insert result (Result true ObjectInput) env) continuation
  where
  lower remaining env original
    | remaining <= 0 = Nothing
    | otherwise = case strip original of
        Let _ level value body -> case atom env value of
          Just bound -> lower (remaining - 1) (Map.insert level bound env) body
          Nothing -> do
            -- A shadowing producer can alias an earlier result. Decline this
            -- shape instead of conflating its generated local with that value.
            guard (not (Map.member level env))
            { fn, args } <- callExpr value
            Tuple optional nullable <- reader fn
            { first: method, second: object, third: label } <- three args
            guard (case atom env object of
              Just ObjectInput -> true
              _ -> false)
            key <- stringExpr label
            { fn: decode, args: dictionaries } <- callExpr method
            guard (standard "decodeJson" decode)
            dictionary <- one dictionaries
            schema <- decoder owner definitions (fuel - 1) dictionary
            guard (complete schema)
            guard (forwards 128 level (Map.insert level (Result false (Failure level)) env) body)
            next <- lower (remaining - 1) (Map.insert level (Result true (Decoded level)) env) body
            pure (ReadField level key optional nullable schema next)
        Branch cases other -> case knownBranch env (NEA.toArray cases) other of
          Just body -> lower (remaining - 1) env body
          Nothing -> do
            branches <- traverse (\(Pair condition body) -> do
              { left, right } <- case strip condition of
                PrimOp (Op2 (OpStringOrd OpEq) left right) -> Just { left, right }
                _ -> Nothing
              level <- atom env left >>= decodedAtom
              key <- stringExpr right
              Tuple (Tuple level key) <$> lower (remaining - 1) env body) (NEA.toArray cases)
            Choices branches <$> lower (remaining - 1) env other
        CtorSaturated ctor _ _ _ [ Tuple "value0" value ] | isEither "Right" ctor -> ReturnValue <$> valueExpr env value
        CtorSaturated ctor _ _ _ [ Tuple "value0" value ] | isEither "Left" ctor -> do
          { errorCtor, label } <- case strip value of
            CtorSaturated errorCtor _ _ _ [ Tuple "value0" label ] -> Just { errorCtor, label }
            _ -> Nothing
          guard (errorCtor == Qualified (Just (ModuleName "Data.Argonaut.Decode.Error")) (Ident "TypeMismatch"))
          message <- stringExpr label
          pure (TypeMismatch message)
        _ -> Nothing

quote :: String -> String
quote = printGoExpr <<< GoString

schemaChildren :: String -> Schema -> Array (Tuple String Schema)
schemaChildren name = case _ of
  Sequence inner -> [ Tuple (name <> "_item") inner ]
  Optional inner -> [ Tuple (name <> "_item") inner ]
  ObjectSchema fields -> Array.mapWithIndex (\i (Tuple _ child) -> Tuple (name <> "_field" <> show i) child) fields
  _ -> []

programSchemas :: String -> Program Schema -> Array (Tuple String Schema)
programSchemas name = case _ of
  ReadField _ _ _ _ schema next -> [ Tuple (name <> "_read") schema ] <> programSchemas (name <> "_next") next
  Choices branches other -> Array.concat (Array.mapWithIndex (\i (Tuple _ branch) -> programSchemas (name <> "_case" <> show i) branch) branches) <> programSchemas (name <> "_else") other
  _ -> []

valueLevels :: NeutralExpr -> Array Level
valueLevels = Set.toUnfoldable <<< collect
  where
  collect (NeutralExpr syn) = case syn of
    Local _ level -> Set.singleton level
    _ -> foldl (\levels child -> Set.union levels (collect child)) Set.empty syn

programSources :: String -> String -> Schema -> Array (Tuple Ident NeutralExpr)
programSources prefix name schema = case schema of
  Derived program -> sources name program
  _ -> Array.concatMap (\(Tuple child value) -> programSources prefix child value) (schemaChildren name schema)
  where
  sources path = case _ of
    ReadField _ _ _ _ _ next -> sources (path <> "_next") next
    Choices branches other -> Array.concat (Array.mapWithIndex (\i (Tuple _ branch) -> sources (path <> "_case" <> show i) branch) branches) <> sources (path <> "_else") other
    ReturnValue value ->
      let
        levels = valueLevels value
        body = case NEA.fromArray (map (Tuple Nothing) levels) of
          Nothing -> value
          Just refs -> NeutralExpr (Abs refs value)
        ident = fromMaybe path (String.stripPrefix (Pattern (prefix <> "_")) path) <> "_construct"
      in [ Tuple (Ident ident) (NeutralExpr (Typed CoreFn.Any body)) ]
    _ -> []

programKeys :: Program Schema -> Array String
programKeys = case _ of
  ReadField _ key _ _ _ next -> Array.cons key (programKeys next)
  Choices branches other -> Array.concatMap (\(Tuple _ branch) -> programKeys branch) branches <> programKeys other
  _ -> []

programCode :: Boolean -> (String -> String) -> String -> Program Schema -> String
programCode text lookup name = case _ of
  ReadField level key optional nullable _ next ->
    "var " <> local level <> " gopurs_runtime.Value\n{\ninput, present := " <> lookup key <> "\nif !present"
      <> (if nullable then " || " <> (if text then "argonautTextNull" else "domNull") <> "(input)" else "") <> " {\n"
      <> (if optional then local level <> " = plan.mkNothing()" else "return argonautSchemaResult{err:plan.mkAtKey(" <> quote key <> ", plan.missingValue)}")
      <> "\n} else {\nresult := " <> name <> "_read" <> (if text then "_text" else "") <> "(plan, nil, input)\nif !result.ok { return argonautSchemaResult{err:plan.mkAtKey(" <> quote key <> ", result.err)} }\n"
      <> local level <> " = " <> (if optional then "plan.mkJust(result.value)" else "result.value") <> "\n}\n}\n_ = " <> local level <> "\n"
      <> programCode text lookup (name <> "_next") next
  Choices branches other -> String.joinWith "\n" (Array.mapWithIndex (\i (Tuple (Tuple level key) branch) ->
    "if " <> local level <> ".StrVal() == " <> quote key <> " {\n" <> programCode text lookup (name <> "_case" <> show i) branch <> "\n}") branches)
    <> "\n" <> programCode text lookup (name <> "_else") other
  ReturnValue value ->
    let levels = valueLevels value
        getter = "Get_" <> name <> "_construct()"
        result = if Array.null levels then getter
          else if Array.length levels <= 10 then "gopurs_runtime.Apply" <> (if Array.length levels == 1 then "" else show (Array.length levels)) <> "(" <> getter <> ", " <> String.joinWith ", " (map local levels) <> ")"
          else foldl (\fn level -> "gopurs_runtime.Apply(" <> fn <> ", " <> local level <> ")") getter levels
    in "return argonautSchemaResult{value:" <> result <> ", ok:true}"
  TypeMismatch message -> "return argonautSchemaResult{err:plan.mkTypeMismatch(" <> quote message <> ")}"
  where
  local level = "field_" <> show (unwrap level)

emit :: String -> String -> Boolean -> Boolean -> Schema -> Array GoDecl
emit prefix name direct text schema = [ worker, checker ] <> nested
  where
  entry path = path <> if text then "_text" else ""
  inputFn fn = if text then String.replaceAll (Pattern "dom") (Replacement "argonautText") fn else fn
  materialized = if text then "raw.materialize()" else "raw"
  keys = Array.nub case schema of
    ObjectSchema fields -> map (\(Tuple key _) -> key) fields
    Derived program -> programKeys program
    _ -> []
  slot key = "inputs[" <> show (fromMaybe 0 (Array.elemIndex key keys)) <> "]"
  lookup key = if text then slot key <> ", " <> slot key <> ".document != nil" else "object.Lookup(" <> quote key <> ")"
  -- Reverse traversal keeps the last normalized input key. Only input offsets
  -- are collected here; decoding still executes in the proven schema order.
  inputSlots = if not text || Array.null keys then "" else
    "var inputs [" <> show (Array.length keys) <> "]argonautTextCursor\n"
      <> "for at := raw.document.tokens[raw.index].end; at != 0; at = raw.document.tokens[at].next {\n"
      <> "switch (argonautTextCursor{raw.document, at}).string(false) {\n"
      <> String.joinWith "\n" (map (\key -> "case " <> quote key <> ":\nif " <> slot key <> ".document == nil { " <> slot key <> " = argonautTextCursor{raw.document, at + 1} }") keys)
      <> "\n}\n}\n"
  children = schemaChildren name schema
  nested = case schema of
    Derived program -> Array.concatMap (\(Tuple child value) -> emit prefix child true text value) (programSchemas name program)
    _ -> Array.concatMap (\(Tuple child value) -> emit prefix child direct text value) children
  checker = GoFunctionDecl
    { name: entry name <> "_accepts", params: [ Tuple "kind" (TypeInterface "*fieldKind") ], result: TypeBool
    , body: rawGo ("return " <> accepts schema)
    }
  accepts = case _ of
    Custom -> "true"
    Derived _ -> "true"
    Scalar tag -> "kind != nil && kind.tag == " <> quote tag
    Sequence child -> containerCheck "Array" child
    Optional child -> containerCheck "Maybe" child
    ObjectSchema fields -> "kind != nil && kind.tag == kindRecord && kind.plan != nil && len(kind.plan.fields) == " <> show (Array.length fields)
      <> String.joinWith "" (Array.mapWithIndex (\i (Tuple _ child) ->
          let member = "kind.plan.fields[" <> show i <> "].kind"
          in childCheck member (name <> "_field" <> show i) child) fields)
  childCheck member path child
    | text = if isDerived child then "" else " && " <> member <> " != nil && " <> entry path <> "_accepts(" <> member <> ")"
    | otherwise = " && (" <> member <> " == nil || " <> path <> "_accepts(" <> member <> "))"
  containerCheck tag child = "kind != nil && kind.tag == " <> quote tag <> childCheck "kind.inner" (name <> "_item") child
  worker = GoFunctionDecl
    { name: entry name, params: [ Tuple "plan" (TypeInterface "*recordDecodePlan"), Tuple "kind" (TypeInterface "*fieldKind"), Tuple "raw" (TypeInterface (if text then "argonautTextCursor" else "any")) ]
    , result: TypeInterface "argonautSchemaResult", body: rawGo (body schema)
    }
  success value = "return argonautSchemaResult{value:" <> value <> ", ok:true}\n"
  failure message = "return argonautSchemaResult{err:plan.mkTypeMismatch(" <> quote message <> ")}\n"
  decodeItem input inner = if direct || isDerived inner
    then "result := " <> entry (name <> "_item") <> "(plan, nil, " <> input <> ")\n"
    else "var result argonautSchemaResult\nif kind.inner == nil { result = argonautSchemaElement(plan, kind, " <> input <> (if text then ".materialize()" else "") <> ") } else { result = " <> entry (name <> "_item") <> "(plan, kind.inner, " <> input <> ") }\n"
  isDerived = case _ of
    Derived _ -> true
    _ -> false
  body = case _ of
    Custom -> "return argonautSchemaCustom(plan, kind, " <> materialized <> ")"
    Derived program -> "object, ok := " <> inputFn "domObject" <> "(raw)\nif !ok {" <> failure "Object" <> "}\n_ = object\n" <> inputSlots <> programCode text lookup name program
    Scalar "Json" -> success ("gopurs_runtime.Box(" <> materialized <> ")")
    Scalar "Int" -> "number, ok := " <> inputFn "domNumber" <> "(raw)\nif !ok {" <> failure "Number" <> "}\nvalue, ok := intFromNumber(number)\nif !ok {" <> failure "Integer" <> "}\n" <> success "gopurs_runtime.Int(value)"
    Scalar tag ->
      let Tuple method box = case tag of
            "Number" -> Tuple "domNumber" "Float"
            "String" -> Tuple "domString" "Str"
            _ -> Tuple "domBool" "Bool"
       in "value, ok := " <> inputFn method <> "(raw)\nif !ok {" <> failure tag <> "}\n" <> success ("gopurs_runtime." <> box <> "(value)")
    Sequence inner -> "items, ok := " <> inputFn "domArrayValues" <> "(raw)\nif !ok {" <> failure "Array" <> "}\nvalues := make([]gopurs_runtime.Value, items.length())\n"
      <> (if text then "at := raw.index + 1\n" else "") <> "for index := range values {\n"
      <> (if text then "element := argonautTextCursor{raw.document, at}\nat = raw.document.tokens[at].next\n" else "")
      <> decodeItem (if text then "element" else "items.at(index)") inner <> "if !result.ok { return argonautSchemaResult{err:plan.mkNamed(\"Array\", plan.mkAtIndex(index, result.err))} }\nvalues[index] = result.value\n}\n" <> success "gopurs_runtime.Array(values)"
    Optional inner -> "if " <> inputFn "domNull" <> "(raw) {" <> success "plan.mkNothing()" <> "}\n" <> decodeItem "raw" inner
      <> "if !result.ok { return result }\n" <> success "plan.mkJust(result.value)"
    ObjectSchema fields -> (if direct then "" else "plan = kind.plan\n") <> "object, ok := " <> inputFn "domObject" <> "(raw)\nif !ok {" <> failure "Object" <> "}\n_ = object\n"
      <> inputSlots
      <> (if direct || Array.null fields then "" else "var boxed gopurs_runtime.Value\n_ = boxed\n")
      <> String.joinWith "\n" (Array.mapWithIndex fieldCode fields)
      <> success (recordValue fields)
  fieldCode index (Tuple key child) =
    let i = show index
        value = "value" <> i
        member = "plan.fields[" <> i <> "].kind"
        absent = case child of
          Optional _ -> value <> " = plan.mkNothing()"
          Custom -> "if " <> member <> ".tag == kindMaybe { " <> value <> " = plan.mkNothing() } else { return argonautSchemaResult{err:plan.mkAtKey(" <> quote key <> ", plan.missingValue)} }"
          _ -> "return argonautSchemaResult{err:plan.mkAtKey(" <> quote key <> ", plan.missingValue)}"
        fast = direct || isDerived child
    in "var " <> value <> " gopurs_runtime.Value\n"
      <> (if fast then "{\n" else "if " <> member <> " == nil {\nif boxed.Type == 0 { boxed = gopurs_runtime.Box(" <> materialized <> ") }\nresult := argonautSchemaField(plan, " <> i <> ", boxed, " <> quote key <> ")\nif !result.ok { return result }; " <> value <> " = result.value\n} else {\n")
      <> "input, present := " <> lookup key <> "\nif !present { " <> absent <> " } else {\nresult := " <> entry (name <> "_field" <> i) <> "(plan, " <> (if fast then "nil" else member) <> ", input)\nif !result.ok { return argonautSchemaResult{err:plan.mkAtKey(" <> quote key <> ", result.err)} }; " <> value <> " = result.value\n}\n}\n_ = " <> value <> "\n"
  recordValue fields =
    let
      slots = foldl insert [] (Array.reverse (Array.mapWithIndex (\i (Tuple key _) -> Tuple key ("value" <> show i)) fields))
      insert acc pair@(Tuple key _) = case Array.findIndex (\(Tuple existing _) -> existing == key) acc of
        Nothing -> Array.snoc acc pair
        Just index -> fromMaybe acc (Array.updateAt index pair acc)
      keys = map (\(Tuple key _) -> quote key) slots
      values = map (\(Tuple _ value) -> value) slots
      count = Array.length slots
    in if count <= 5 then "gopurs_runtime.RecordDict" <> show count <> "(" <> String.joinWith ", " (keys <> values) <> ")"
       else "gopurs_runtime.RecordDict([]string{" <> String.joinWith ", " keys <> "}, []gopurs_runtime.Value{" <> String.joinWith ", " values <> "})"
