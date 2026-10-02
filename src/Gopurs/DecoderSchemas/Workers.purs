module Gopurs.DecoderSchemas.Workers (build) where

import Prelude

import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (foldl)
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Newtype (unwrap)
import Data.Set as Set
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Tuple (Tuple(..))
import Gopurs.DecoderSchemas.Types (Program(..), Schema(..), textComplete)
import Gopurs.GoAst (GoDecl(..), GoExpr(..), GoType(..), rawGo, sanitizeName)
import Gopurs.Printer (printGoExpr)
import PureScript.Backend.Optimizer.CoreFn (Ident(..))
import PureScript.Backend.Optimizer.CoreFn as CoreFn
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), Level)

-- One reserved family produces DOM workers, optional text workers and their
-- cached entry. Both modes call the same ordinary constructor source getters.
-- The facade supplies the untouched original decoder through the source getter.
build
  :: { prefix :: String, name :: String, source :: String, textEnabled :: Boolean }
  -> Schema
  -> { sources :: Array (Tuple Ident NeutralExpr), declarations :: Array GoDecl }
build { prefix, name, source, textEnabled } schema =
  let
    qualified = prefix <> "_" <> sanitizeName name
    worker = qualified <> "_decode"
    withText = textEnabled && textComplete schema
    generated = emit { direct: false, text: false } worker schema
      <> if withText then emit { direct: false, text: true } worker schema else []
    compiled = GoCall (GoVar "argonautCompileSchema")
      [ GoCall (GoVar ("Get_" <> prefix <> "_" <> sanitizeName source)) []
      , GoVar worker, GoVar (worker <> "_accepts")
      ]
    getter = GoCachedValue
      { identifier: qualified, goType: TypeValue
      , expression: if withText then GoCall (GoVar "argonautCompileTextSchema")
          [ compiled, GoVar (worker <> "_text"), GoVar (worker <> "_text_accepts") ] else compiled
      }
  in { sources: programSources prefix worker schema, declarations: generated <> [ getter ] }

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

-- Sorted, unique levels must agree between a constructor's lambda parameters
-- and every DOM/text call to that constructor, regardless of first-use order.
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

-- Ordinary compositions consult the runtime plan and retain opaque callbacks.
-- Reads in a proven program are complete and dispatch directly with no child tag.
-- Text mode has stricter guards: it cannot fall back on partially parsed input.
emit :: { direct :: Boolean, text :: Boolean } -> String -> Schema -> Array GoDecl
emit mode@{ direct, text } name schema = [ worker, checker ] <> nested
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
    Derived program -> Array.concatMap (\(Tuple child value) -> emit (mode { direct = true }) child value) (programSchemas name program)
    _ -> Array.concatMap (\(Tuple child value) -> emit mode child value) children
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

-- Reproduce row-list insertion: tail slots come first, then an earlier field
-- replaces the value of a duplicate label without moving its output slot.
recordValue :: Array (Tuple String Schema) -> String
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
