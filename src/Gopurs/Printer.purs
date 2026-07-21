module Gopurs.Printer where

import Prelude
import Data.String as String
import Gopurs.GoAst (GoExpr(..), GoDecl, GoFile)

printGoExpr :: GoExpr -> String
printGoExpr expr = case expr of
  GoVar name ->
    name
  GoString s ->
    "\"" <> s <> "\""
  GoInt i ->
    show i
  GoCall f args ->
    printGoExpr f <> "(" <> String.joinWith ", " (map printGoExpr args) <> ")"
  GoSelector obj field ->
    printGoExpr obj <> "." <> field
  GoFunc args body ->
    "func(" <> String.joinWith ", " (map (\a -> a <> " gopurs_runtime.Value") args) <> ") gopurs_runtime.Value {\n" <>
    printGoExpr body <>
    "\n}"
  GoBlock stmts ->
    String.joinWith "\n" (map printGoExpr stmts)
  GoReturn e ->
    "return " <> printGoExpr e
  GoAssign name e ->
    name <> " := " <> printGoExpr e
  GoRaw raw ->
    raw

printGoDecl :: GoDecl -> String
printGoDecl { identifier, expression } =
  "var " <> identifier <> " = " <> printGoExpr expression

printGoFile :: GoFile -> String
printGoFile { packageName, imports, decls } =
  "package " <> packageName <> "\n\n" <>
  "import (\n" <>
  String.joinWith "\n" (map (\i -> "\t\"" <> i <> "\"") imports) <> "\n" <>
  ")\n\n" <>
  String.joinWith "\n" (map printGoDecl decls) <> "\n"
