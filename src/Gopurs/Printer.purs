module Gopurs.Printer where

import Prelude
import Data.String as String
import Gopurs.GoAst (GoExpr(..), GoDecl, GoFile)
import Data.Tuple (Tuple(..))

printGoExpr :: GoExpr -> String
printGoExpr expr = case expr of
  GoVar name ->
    name
  GoString s ->
    show s
  GoInt i ->
    show i
  GoCall f args ->
    printGoExpr f <> "(" <> String.joinWith ", " (map printGoExpr args) <> ")"
  GoSelector obj field ->
    printGoExpr obj <> "." <> field
  GoFunc arg body ->
    "gopurs_runtime.Func(func(" <> arg <> " gopurs_runtime.Value) gopurs_runtime.Value {\nreturn " <> printGoExpr body <> "\n})"
  GoBlock stmts ->
    String.joinWith "\n" (map printGoExpr stmts)
  GoReturn e ->
    "return " <> printGoExpr e
  GoAssign name e ->
    name <> " := " <> printGoExpr e
  GoMap props ->
    "map[string]gopurs_runtime.Value{" <> String.joinWith ", " (map (\(Tuple k v) -> "\"" <> k <> "\": " <> printGoExpr v) props) <> "}"
  GoIIFE name binding body ->
    let assignment = if name == "_" then name <> " = " <> printGoExpr binding else name <> " := " <> printGoExpr binding <> "\n_ = " <> name
    in "func() gopurs_runtime.Value {\n" <> assignment <> "\nreturn " <> printGoExpr body <> "\n}()"
  GoLetRec bindings body ->
    "func() gopurs_runtime.Value {\n" <>
    String.joinWith "\n" (map (\(Tuple name _) -> "var " <> name <> " gopurs_runtime.Value") bindings) <> "\n" <>
    String.joinWith "\n" (map (\(Tuple name _) -> "_ = " <> name) bindings) <> "\n" <>
    String.joinWith "\n" (map (\(Tuple name expr) -> name <> " = " <> printGoExpr expr) bindings) <> "\n" <>
    "return " <> printGoExpr body <> "\n}()"
  GoRecordAccess obj prop ->
    printGoExpr obj <> ".PtrVal.(map[string]gopurs_runtime.Value)[\"" <> prop <> "\"]"
  GoBranch branches def ->
    "func() gopurs_runtime.Value {\n" <>
    String.joinWith "\n" (map (\(Tuple cond t) -> "if (" <> printGoExpr cond <> ").IntVal != 0 {\nreturn " <> printGoExpr t <> "\n}") branches) <>
    "\nreturn " <> printGoExpr def <> "\n}()"
  GoBinOp op left right ->
    printGoExpr left <> " " <> op <> " " <> printGoExpr right
  GoTypeAssertion expr t ->
    printGoExpr expr <> ".(" <> t <> ")"
  GoRaw raw ->
    raw

printGoDeclVar :: GoDecl -> String
printGoDeclVar { identifier } =
  "var " <> identifier <> " gopurs_runtime.Value"

printGoDeclInit :: GoDecl -> String
printGoDeclInit { identifier, expression } =
  "\t" <> identifier <> " = " <> printGoExpr expression

printGoFile :: GoFile -> String
printGoFile { packageName, imports, decls } =
  "package " <> packageName <> "\n\n" <>
  "import (\n" <>
  String.joinWith "\n" (map (\i -> "\t\"" <> i <> "\"") imports) <> "\n" <>
  ")\n\n" <>
  String.joinWith "\n" (map printGoDeclVar decls) <> "\n\n" <>
  "func init() {\n" <>
  String.joinWith "\n" (map printGoDeclInit decls) <> "\n" <>
  "}\n"
