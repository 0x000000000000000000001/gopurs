module Gopurs.RecordShapeGuard (integerReads) where

import Prelude

import Data.Array as Array
import Data.String as String
import Gopurs.GoAst (GoExpr(..))
import Gopurs.Printer (printGoExpr)

-- Check the complete shape at each consumption site. Neither TAST field types
-- nor a previously checked alias establish an opaque Value's current layout.
-- Callers admit unique labels and use only side-effect-free IntVal projections.
integerReads :: String -> String -> Array String -> (Array GoExpr -> String) -> String
integerReads prefix input keys body =
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
        <> "if len(" <> dataName <> ".Keys) == " <> count <> " && len(" <> dataName <> ".Vals) >= " <> count
        <> " && " <> labels <> " {\n" <> body values <> "\n}\n}\n"
    else ""
