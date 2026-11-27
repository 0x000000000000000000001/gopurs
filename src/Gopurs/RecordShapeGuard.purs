module Gopurs.RecordShapeGuard (integerReads, integerReadsWithShape, internShapeExpr) where

import Prelude

import Data.Array as Array
import Data.Maybe (Maybe(..))
import Data.String as String
import Effect.Ref (Ref)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)
import Gopurs.CodegenState (CodegenState)
import Gopurs.GoAst (GoDecl(..), GoExpr(..), GoType(..))
import Gopurs.GoSymbol (GoReference(..), generated)
import Gopurs.Printer (printGoExpr)

-- Register a package-level canonical *gopurs_runtime.RecordShape for a >5-field
-- key layout so hot guards can validate all keys with one pointer comparison.
internShapeExpr :: Ref CodegenState -> String -> Array String -> GoExpr
internShapeExpr stateRef modNameStr keys = unsafePerformEffect do
  state <- Ref.read stateRef
  let
    name = "__gopurs_record_shape_" <> modNameStr <> "_" <> show state.globalId
    decl = GoCachedValue
      { identifier: generated modNameStr name
      , expression: GoCall (GoSelector (GoVar "gopurs_runtime") "InternRecordShape") (map GoString keys)
      , goType: TypeInterface "*gopurs_runtime.RecordShape"
      }
  Ref.modify_ (\current -> current
    { globalId = current.globalId + 1
    , declarations = current.declarations <> [ decl ]
    }) stateRef
  pure (GoCall (GoRef (SymbolReference (generated modNameStr ("Get_" <> name)))) [])

-- Check the complete shape at each consumption site. Neither TAST field types
-- nor a previously checked alias establish an opaque Value's current layout.
-- Callers admit unique labels and use only side-effect-free IntVal projections.
integerReads :: String -> String -> Array String -> (Array GoExpr -> String) -> String
integerReads = integerReadsWithShape Nothing

integerReadsWithShape :: Maybe GoExpr -> String -> String -> Array String -> (Array GoExpr -> String) -> String
integerReadsWithShape mbShape prefix input keys body =
  let
    size = Array.length keys
    count = show size
    dataName = prefix <> "_data"
    keysName = if size == 2 then prefix <> "_keys" else dataName <> ".RecordKeys" <> count
    valueName index = prefix <> "_value" <> show index
    quote key = printGoExpr (GoString key)
    compactLabels = String.joinWith " && " (Array.mapWithIndex
      (\index key -> keysName <> ".K" <> show index <> " == " <> quote key) keys)
    compactValues = Array.mapWithIndex (\index _ -> GoSelector
      (if size == 2 then GoVar (valueName index) else GoSelector (GoVar dataName) ("V" <> show index)) "IntVal") keys
    compactLoad = if size == 2 then
      keysName <> ", " <> valueName 0 <> ", " <> valueName 1 <> " := gopurs_runtime.RecordFields2(" <> input <> ")\n"
      else dataName <> " := (*gopurs_runtime.RecordData" <> count <> ")(" <> input <> ".UnsafePtr)\n"
    labels = String.joinWith " && " (Array.mapWithIndex
      (\index key -> dataName <> ".Keys[" <> show index <> "] == " <> quote key) keys)
    shapeCheck = case mbShape of
      Just shapeExpr ->
        "len(" <> dataName <> ".Vals) >= " <> count
          <> " && (" <> dataName <> ".Shape == " <> printGoExpr shapeExpr
          <> " || (len(" <> dataName <> ".Keys) == " <> count <> " && " <> labels <> "))"
      Nothing ->
        "len(" <> dataName <> ".Keys) == " <> count <> " && len(" <> dataName <> ".Vals) >= " <> count <> " && " <> labels
    values = Array.mapWithIndex (\index _ ->
      GoSelector (GoIndex (GoSelector (GoVar dataName) "Vals") (GoInt index)) "IntVal") keys
  in
    if size >= 1 && size <= 5 then
      "if " <> input <> ".Type == gopurs_runtime.TypeRecord" <> count <> " && " <> input <> ".UnsafePtr != nil {\n"
        <> compactLoad <> "if " <> keysName <> " != nil && " <> compactLabels <> " {\n"
        <> body compactValues <> "\n}\n}\n"
    else if size > 5 then
      "if " <> input <> ".Type == gopurs_runtime.TypeRecordData && " <> input <> ".UnsafePtr != nil {\n"
        <> dataName <> " := (*gopurs_runtime.RecordData)(" <> input <> ".UnsafePtr)\n"
        <> "if " <> shapeCheck <> " {\n" <> body values <> "\n}\n}\n"
    else ""
