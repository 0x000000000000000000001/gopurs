-- Transactional decoding for schemas already proved by Admission/Programs.
-- Only successful complete parses escape; every rejection restarts the old text
-- decoder. Speculation therefore admits constructors and total scalar decoding,
-- never opaque decoders or error callbacks.
module Gopurs.DecoderSchemas.Stream (build, supported) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Foldable (all, foldl)
import Data.Maybe (Maybe(..), fromMaybe, isJust)
import Data.Newtype (unwrap)
import Data.String as String
import Data.Tuple (Tuple(..), fst)
import Gopurs.DecoderSchemas.Types (Program(..), Schema(..), complete)
import Gopurs.GoAst (GoDecl(..), GoExpr(..), GoType(..), rawGo)
import Gopurs.GoSymbol as Symbol
import Gopurs.Printer (printGoExpr)
import PureScript.Backend.Optimizer.CoreFn (Literal(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (Level)
import PureScript.Backend.Optimizer.Syntax as Syn

type Rendering =
  { owner :: String
  , recordValue :: Array (Tuple String Schema) -> String
  , constructValue :: String -> NeutralExpr -> String
  }

type Read = { key :: String, path :: String, schema :: Schema, nullable :: Boolean }

supported :: Schema -> Boolean
supported = case _ of
  Scalar tag -> Array.elem tag [ "Int", "Number", "String", "Boolean" ]
  Sequence child -> supported child
  Optional child -> supported child
  ObjectSchema fields ->
    let keys = map fst fields
    in Array.length (Array.nub keys) == Array.length keys && all (supported <<< \(Tuple _ child) -> child) fields
  Derived program -> isJust (readFields "" program)
  Custom -> false

-- Derived reads are already complete (no nested Derived or Custom). Compare
-- their actual shapes, not decoder names or tags, before sharing a value slot.
sameSchema :: Schema -> Schema -> Boolean
sameSchema left right = case left, right of
  Scalar a, Scalar b -> a == b
  Sequence a, Sequence b -> sameSchema a b
  Optional a, Optional b -> sameSchema a b
  ObjectSchema a, ObjectSchema b -> Array.length a == Array.length b
    && all identity (Array.zipWith (\(Tuple ka va) (Tuple kb vb) -> ka == kb && sameSchema va vb) a b)
  _, _ -> false

programReads :: String -> Program Schema -> Array Read
programReads path = case _ of
  ReadField _ key _ nullable schema next ->
    [ { key, path: path <> "_read", schema, nullable } ] <> programReads (path <> "_next") next
  Choices branches other -> Array.concat (Array.mapWithIndex (\i (Tuple _ branch) ->
    programReads (path <> "_case" <> show i) branch) branches) <> programReads (path <> "_else") other
  _ -> []

readFields :: String -> Program Schema -> Maybe (Array Read)
readFields path program = do
  guard (pureProgram program && uniqueReads [] program)
  foldl insert (Just []) (programReads path program)
  where
  insert previous field = do
    fields <- previous
    guard (complete field.schema && supported field.schema)
    case Array.findIndex (\existing -> existing.key == field.key) fields of
      Nothing -> pure (Array.snoc fields field)
      Just index -> do
        existing <- Array.index fields index
        guard (sameSchema existing.schema field.schema)
        Array.updateAt index (existing { nullable = existing.nullable || field.nullable }) fields

-- Decoding one container twice creates independent results in the old worker.
-- Sharing a slot for those reads would introduce a mutable alias in the output.
-- Branches have independent paths; sharing their slot is safe because only one
-- branch executes. Repeated uses of one decoded local keep their original alias.
uniqueReads :: Array String -> Program Schema -> Boolean
uniqueReads seen = case _ of
  ReadField _ key _ _ _ next ->
    not (Array.elem key seen) && uniqueReads (Array.snoc seen key) next
  Choices branches other -> all (\(Tuple _ branch) -> uniqueReads seen branch) branches && uniqueReads seen other
  _ -> true

-- Keep the construction proof explicit at this speculative boundary as well
-- as in Programs.valueExpr. A future Program extension cannot silently admit
-- calls or effects before the complete input has been accepted.
pureProgram :: Program Schema -> Boolean
pureProgram = case _ of
  ReadField _ _ _ _ _ next -> pureProgram next
  Choices branches other -> all (\(Tuple _ branch) -> pureProgram branch) branches && pureProgram other
  ReturnValue value -> pureConstruction value
  TypeMismatch _ -> true

pureConstruction :: NeutralExpr -> Boolean
pureConstruction (NeutralExpr syntax) = case syntax of
  Syn.Typed _ inner -> pureConstruction inner
  Syn.TypeApp inner _ -> pureConstruction inner
  Syn.Local _ _ -> true
  Syn.Lit (LitString _) -> true
  Syn.Lit (LitInt _) -> true
  Syn.Lit (LitNumber _) -> true
  Syn.Lit (LitBoolean _) -> true
  Syn.CtorSaturated _ _ _ _ fields -> all (\(Tuple _ value) -> pureConstruction value) fields
  _ -> false

children :: String -> Schema -> Array (Tuple String Schema)
children name = case _ of
  Sequence child -> [ Tuple (name <> "_item") child ]
  Optional child -> [ Tuple (name <> "_item") child ]
  ObjectSchema fields -> Array.mapWithIndex (\i (Tuple _ child) -> Tuple (name <> "_field" <> show i) child) fields
  Derived program -> map (\read -> Tuple read.path read.schema) (programReads name program)
  _ -> []

build :: Rendering -> String -> Schema -> Array GoDecl
build rendering name schema
  | supported schema = emit rendering false name schema
  | otherwise = []

emit :: Rendering -> Boolean -> String -> Schema -> Array GoDecl
emit rendering@{ owner } direct name schema = [ worker ] <> nested
  where
  recordType = TypeNamed <<< Symbol.foreignSymbol "Data_Argonaut_Decode_Internal_Record"
  worker = GoFunctionDecl
    { name: Symbol.generated owner (name <> "_stream")
    , params:
        [ Tuple "plan" (TypePointer (recordType "recordDecodePlan"))
        , Tuple "kind" (TypePointer (recordType "fieldKind"))
        , Tuple "p" (TypePointer (TypeNamed (Symbol.foreignSymbol "Data_Argonaut_Decode_Parser" "argonautStreamParser")))
        ]
    , result: TypeResults [ TypeValue, TypeBool ]
    , body: rawGo (body rendering direct name schema)
    }
  childDirect = direct || case schema of
    Derived _ -> true
    _ -> false
  nested = Array.concatMap (\(Tuple child value) -> emit rendering childDirect child value) (children name schema)

failure :: String
failure = "return gopurs_runtime.Value{}, false"

quote :: String -> String
quote = printGoExpr <<< GoString

entry :: String -> String
entry name = name <> "_stream"

isDerived :: Schema -> Boolean
isDerived (Derived _) = true
isDerived _ = false

call :: String -> String -> String
call name kind = entry name <> "(plan, " <> kind <> ", p)"

enter :: String -> String
enter token = "if !p.enter() { " <> failure <> " }; if !p.take('" <> token <> "') { " <> failure <> " }\n"

body :: Rendering -> Boolean -> String -> Schema -> String
body rendering direct name = case _ of
  Scalar "Int" -> "number, ok := p.number(); if !ok { " <> failure <> " }; value, ok := intFromNumber(number); if !ok { " <> failure <> " }; return gopurs_runtime.Int(value), true"
  Scalar tag ->
    let Tuple method box = case tag of
          "String" -> Tuple "string(true)" "Str"
          "Number" -> Tuple "number()" "Float"
          _ -> Tuple "boolean()" "Bool"
    in "value, ok := p." <> method <> "; if !ok { " <> failure <> " }; return gopurs_runtime." <> box <> "(value), true"
  Optional child ->
    "if p.peek() == 'n' { if !p.null() { " <> failure <> " }; return plan.mkNothing(), true }\n"
      <> "value, ok := " <> call (name <> "_item") (if direct || isDerived child then "nil" else "kind.inner")
      <> "; if !ok { " <> failure <> " }; return plan.mkJust(value), true"
  Sequence child -> enter "["
    <> "values := make([]gopurs_runtime.Value, 0)\nif !p.take(']') { for {\n"
    <> "value, ok := " <> call (name <> "_item") (if direct || isDerived child then "nil" else "kind.inner")
    <> "; if !ok { " <> failure <> " }\n"
    <> "if len(values) == 0 { values = make([]gopurs_runtime.Value, 0, 4) }\n"
    <> "values = append(values, value)\nif p.take(']') { break }; if !p.take(',') { " <> failure <> " }\n} }\n"
    <> "p.leave(); return gopurs_runtime.Array(values), true"
  ObjectSchema fields -> objectBody rendering direct name fields
  Derived program -> case readFields name program of
    Just fields -> derivedBody rendering fields name program
    Nothing -> failure
  Custom -> failure

-- Only these containers entered a depth frame. Rejected parsers are discarded
-- at the root, so restoring depth on failure adds no observable behavior.
objectBody :: Rendering -> Boolean -> String -> Array (Tuple String Schema) -> String
objectBody rendering direct name fields =
  (if direct then "" else "plan = kind.plan\n") <> enter "{"
    <> "var seen [" <> show (Array.length fields) <> "]bool\n_ = seen\n"
    <> String.joinWith "\n" (Array.mapWithIndex (\i _ -> "var value" <> show i <> " gopurs_runtime.Value") fields)
    <> "\nif !p.take('}') { for {\nkey, ok := p.string(false); if !ok || !p.take(':') { " <> failure <> " }\nswitch key {\n"
    <> String.joinWith "\n" (Array.mapWithIndex field fields)
    <> "\ndefault: " <> failure <> "\n}\nif p.take('}') { break }; if !p.take(',') { " <> failure <> " }\n} }\n"
    <> String.joinWith "\n" (Array.mapWithIndex absent fields)
    <> "\np.leave(); return " <> rendering.recordValue fields <> ", true"
  where
  field i (Tuple key child) =
    let index = show i
        kind = if direct || isDerived child then "nil" else "plan.fields[" <> index <> "].kind"
    in "case " <> quote key <> ":\nif seen[" <> index <> "] { " <> failure <> " }; seen[" <> index <> "] = true\n"
      <> "value, ok := " <> call (name <> "_field" <> index) kind <> "; if !ok { " <> failure <> " }; value" <> index <> " = value"
  absent i (Tuple _ child) =
    let fallback = case child of
          Optional _ -> "value" <> show i <> " = plan.mkNothing()"
          _ -> failure
    in "if !seen[" <> show i <> "] { " <> fallback <> " }"

derivedBody :: Rendering -> Array Read -> String -> Program Schema -> String
derivedBody rendering fields name program = enter "{"
  <> "var values [" <> count <> "]gopurs_runtime.Value; var seen, nulls [" <> count <> "]bool\n_ = values; _ = seen; _ = nulls\n"
  <> "if !p.take('}') { for {\nkey, ok := p.string(false); if !ok || !p.take(':') { " <> failure <> " }\nswitch key {\n"
  <> String.joinWith "\n" (Array.mapWithIndex field fields)
  <> "\ndefault: " <> failure <> "\n}\nif p.take('}') { break }; if !p.take(',') { " <> failure <> " }\n} }\n"
  <> programCode rendering fields name program
  where
  count = show (Array.length fields)
  field i read =
    let index = show i
        decode = "value, ok := " <> call read.path "nil" <> "; if !ok { " <> failure <> " }; values[" <> index <> "] = value"
    in "case " <> quote read.key <> ":\nif seen[" <> index <> "] { " <> failure <> " }; seen[" <> index <> "] = true\n"
      <> if read.nullable then
        "if p.peek() == 'n' { if !p.null() { " <> failure <> " }; nulls[" <> index <> "] = true } else { " <> decode <> " }"
      else decode

programCode :: Rendering -> Array Read -> String -> Program Schema -> String
programCode rendering fields name = case _ of
  ReadField level key optional nullable _ next ->
    let index = show (fromMaybe 0 (Array.findIndex (\field -> field.key == key) fields))
        variable = local level
        missing = "!seen[" <> index <> "]" <> if nullable then " || nulls[" <> index <> "]" else ""
        absent = if optional then variable <> " = plan.mkNothing()" else failure
        value = "values[" <> index <> "]"
    in "var " <> variable <> " gopurs_runtime.Value\nif " <> missing <> " { " <> absent <> " } else {\n"
      <> "if nulls[" <> index <> "] { " <> failure <> " }\n"
      <> variable <> " = " <> (if optional then "plan.mkJust(" <> value <> ")" else value) <> "\n}\n_ = " <> variable <> "\n"
      <> programCode rendering fields (name <> "_next") next
  Choices branches other -> String.joinWith "\n" (Array.mapWithIndex (\i (Tuple (Tuple level key) branch) ->
    "if " <> local level <> ".StrVal() == " <> quote key <> " {\n" <> programCode rendering fields (name <> "_case" <> show i) branch <> "\n}") branches)
    <> "\n" <> programCode rendering fields (name <> "_else") other
  ReturnValue value -> "p.leave(); return " <> rendering.constructValue name value <> ", true"
  TypeMismatch _ -> failure
  where
  local :: Level -> String
  local level = "field_" <> show (unwrap level)
