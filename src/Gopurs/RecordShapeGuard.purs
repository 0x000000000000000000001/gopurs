module Gopurs.RecordShapeGuard (integerReads, integerReadsWithShape, integerReadsStrictShape, internShapeExpr, internKeysExpr) where

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

-- Register a package-level canonical *gopurs_runtime.RecordKeysN for a 1..5
-- field layout so hot guards can validate keys with one pointer comparison.
internKeysExpr :: Ref CodegenState -> String -> Array String -> GoExpr
internKeysExpr stateRef modNameStr keys = unsafePerformEffect do
  state <- Ref.read stateRef
  let
    count = show (Array.length keys)
    name = "__gopurs_record_keys_" <> modNameStr <> "_" <> show state.globalId
    decl = GoCachedValue
      { identifier: generated modNameStr name
      , expression: GoCall (GoSelector (GoVar "gopurs_runtime") ("InternRecordKeys" <> count)) (map GoString keys)
      , goType: TypeInterface ("*gopurs_runtime.RecordKeys" <> count)
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
    compactKeyCheck = case mbShape of
      Just shapeExpr ->
        keysName <> " == " <> printGoExpr shapeExpr <> " || (" <> keysName <> " != nil && " <> compactLabels <> ")"
      Nothing ->
        keysName <> " != nil && " <> compactLabels
    compactValues = Array.mapWithIndex (\index _ ->
      if size == 2 then GoVar (valueName index)
      else GoSelector (GoSelector (GoVar dataName) ("V" <> show index)) "IntVal") keys
    compactLoad = if size == 2 then
      keysName <> ", " <> valueName 0 <> ", " <> valueName 1 <> " := gopurs_runtime.RecordIntFields2(" <> input <> ")\n"
      else dataName <> " := (*gopurs_runtime.RecordData" <> count <> ")(" <> input <> ".UnsafePtr)\n"
    labels = String.joinWith " && " (Array.mapWithIndex
      (\index key -> dataName <> ".Keys[" <> show index <> "] == " <> quote key) keys)
    values = Array.mapWithIndex (\index _ ->
      GoSelector (GoIndex (GoSelector (GoVar dataName) "Vals") (GoInt index)) "IntVal") keys
    intValues = Array.mapWithIndex (\index _ ->
      GoIndex (GoSelector (GoVar dataName) "IntVals") (GoInt index)) keys
  in
    if size >= 1 && size <= 5 then
      "if " <> input <> ".Type == gopurs_runtime.TypeRecord" <> count <> " && " <> input <> ".UnsafePtr != nil {\n"
        <> compactLoad <> "if " <> compactKeyCheck <> " {\n"
        <> body compactValues <> "\n}\n}\n"
    else if size > 5 then
      case mbShape of
        Just shapeExpr ->
          let shapeStr = printGoExpr shapeExpr
          in "if " <> input <> ".Type == gopurs_runtime.TypeRecordData && " <> input <> ".UnsafePtr != nil {\n"
            <> dataName <> " := (*gopurs_runtime.RecordData)(" <> input <> ".UnsafePtr)\n"
            <> "if " <> dataName <> ".Shape == " <> shapeStr <> " && len(" <> dataName <> ".IntVals) == " <> count <> " {\n"
            <> "_ = " <> dataName <> ".IntVals[" <> show (size - 1) <> "]\n"
            <> body intValues <> "\n"
            <> "} else if len(" <> dataName <> ".Vals) >= " <> count
            <> " && (" <> dataName <> ".Shape == " <> shapeStr
            <> " || (len(" <> dataName <> ".Keys) == " <> count <> " && " <> labels <> ")) {\n"
            <> "_ = " <> dataName <> ".Vals[" <> show (size - 1) <> "]\n"
            <> body values <> "\n}\n}\n"
        Nothing ->
          "if " <> input <> ".Type == gopurs_runtime.TypeRecordData && " <> input <> ".UnsafePtr != nil {\n"
            <> dataName <> " := (*gopurs_runtime.RecordData)(" <> input <> ".UnsafePtr)\n"
            <> "if len(" <> dataName <> ".Keys) == " <> count <> " && len(" <> dataName <> ".Vals) >= " <> count <> " && " <> labels <> " {\n"
            <> body values <> "\n}\n}\n"
    else ""

-- Strict single-pointer-equality fast guard for callers that already provide a
-- cold fallback covering non-interned and mixed-value records.
integerReadsStrictShape :: Maybe GoExpr -> String -> String -> Array String -> (Array GoExpr -> String) -> String
integerReadsStrictShape mbShape prefix input keys body =
  let
    size = Array.length keys
    count = show size
    dataName = prefix <> "_data"
    keysName = if size == 2 then prefix <> "_keys" else dataName <> ".RecordKeys" <> count
    valueName index = prefix <> "_value" <> show index
    compactValues = Array.mapWithIndex (\index _ ->
      if size == 2 then GoVar (valueName index)
      else GoSelector (GoSelector (GoVar dataName) ("V" <> show index)) "IntVal") keys
    compactLoad = if size == 2 then
      keysName <> ", " <> valueName 0 <> ", " <> valueName 1 <> " := gopurs_runtime.RecordIntFields2(" <> input <> ")\n"
      else dataName <> " := (*gopurs_runtime.RecordData" <> count <> ")(" <> input <> ".UnsafePtr)\n"
    intValues = Array.mapWithIndex (\index _ ->
      GoIndex (GoSelector (GoVar dataName) "IntVals") (GoInt index)) keys
  in
    case mbShape of
      Just shapeExpr | size >= 1 && size <= 5 ->
        "if " <> input <> ".Type == gopurs_runtime.TypeRecord" <> count <> " && " <> input <> ".UnsafePtr != nil {\n"
          <> compactLoad <> "if " <> keysName <> " == " <> printGoExpr shapeExpr <> " {\n"
          <> body compactValues <> "\n}\n}\n"
      Just shapeExpr | size > 5 ->
        "if " <> input <> ".Type == gopurs_runtime.TypeRecordData && " <> input <> ".UnsafePtr != nil {\n"
          <> dataName <> " := (*gopurs_runtime.RecordData)(" <> input <> ".UnsafePtr)\n"
          <> "if " <> dataName <> ".Shape == " <> printGoExpr shapeExpr <> " && len(" <> dataName <> ".IntVals) == " <> count <> " {\n"
          <> "_ = " <> dataName <> ".IntVals[" <> show (size - 1) <> "]\n"
          <> body intValues <> "\n}\n}\n"
      _ -> integerReadsWithShape mbShape prefix input keys body
