module Gopurs.Printer where

import Prelude
import Data.String as String
import Gopurs.GoAst (GoExpr(..), GoDecl, GoFile)
import Data.Tuple (Tuple(..))
import Data.Array as Array
import Data.Maybe (Maybe(..), fromMaybe)

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
  GoFunc arg body -> case body of
    GoBlock _ -> "gopurs_runtime.Func(func(" <> arg <> " gopurs_runtime.Value) gopurs_runtime.Value {\n" <> printGoExpr body <> "\n})"
    _ -> "gopurs_runtime.Func(func(" <> arg <> " gopurs_runtime.Value) gopurs_runtime.Value {\nreturn " <> printGoExpr body <> "\n})"
  GoBlock stmts ->
    String.joinWith "\n" (map printGoExpr stmts)
  GoReturn e ->
    "return " <> printGoExpr e
  GoAssign name e ->
    name <> " := " <> printGoExpr e <> "\n_ = " <> name
  GoRecordDict props ->
    let
      keysStr = String.joinWith ", " (map (\(Tuple k _) -> "\"" <> k <> "\"") props)
      valsStr = String.joinWith ", " (map (\(Tuple _ v) -> printGoExpr v) props)
    in
      "gopurs_runtime.RecordDict([]string{" <> keysStr <> "}, []gopurs_runtime.Value{" <> valsStr <> "})"
  GoRecordUpdateDict orig props ->
    let
      keysStr = String.joinWith ", " (map (\(Tuple k _) -> "\"" <> k <> "\"") props)
      valsStr = String.joinWith ", " (map (\(Tuple _ v) -> printGoExpr v) props)
    in
      "gopurs_runtime.RecordUpdateDict(" <> printGoExpr orig <> ", []string{" <> keysStr <> "}, []gopurs_runtime.Value{" <> valsStr <> "})"
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
    "gopurs_runtime.RecordGet(" <> printGoExpr obj <> ", \"" <> prop <> "\")"
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
  GoFor label stmts ->
    label <> ":\nfor {\nif false { continue " <> label <> " }\n" <> String.joinWith "\n" (map printGoExpr stmts) <> "\n}"
  GoContinue label ->
    "continue " <> label
  GoIfElse cond trueStmts falseStmts ->
    "if (" <> printGoExpr cond <> ").IntVal != 0 {\n" <>
      String.joinWith "\n" (map printGoExpr trueStmts) <>
    "\n} else {\n" <>
      String.joinWith "\n" (map printGoExpr falseStmts) <>
    "\n}"
  GoMutate name expr ->
    name <> " = " <> printGoExpr expr



printGoDeclVar :: GoDecl -> String
printGoDeclVar { identifier, expression } =
  "var " <> identifier <> " gopurs_runtime.Value\n" <>
  "var once_" <> identifier <> " sync.Once\n" <>
  "func Get_" <> identifier <> "() gopurs_runtime.Value {\n" <>
  "\tonce_" <> identifier <> ".Do(func() {\n" <>
  "\t\t" <> identifier <> " = " <> printGoExpr expression <> "\n" <>
  "\t})\n" <>
  "\treturn " <> identifier <> "\n" <>
  "}"

printGoFile :: GoFile -> String
printGoFile { packageName, imports, decls, foreigns } =
  let declsStr = String.joinWith "\n\n" (map printGoDeclVar decls) <> "\n\n" <> String.joinWith "\n\n" (map (\f -> "func Get_" <> f.pursName <> "() gopurs_runtime.Value {\n\treturn " <> f.goName <> "\n}") foreigns) <> "\n"
      usedImports1 = Array.filter (\i -> 
        if i == "gopurs/output/gopurs_runtime" || i == "sync" then true
        else 
          let pkg = Array.last (String.split (String.Pattern "/") i)
              pkgAlias = "pkg_" <> String.replaceAll (String.Pattern ".") (String.Replacement "_") (fromMaybe "" pkg)
          in String.contains (String.Pattern (pkgAlias <> ".")) declsStr
      ) imports
      missingDeps = ["Unsafe.Coerce", "Data.Unit", "Control.Monad.ST.Internal", "Data.Eq", "Data.Function.Uncurried", "Control.Category"]
      injectedDeps = Array.filter (\dep ->
        let pkgAlias = "pkg_" <> String.replaceAll (String.Pattern ".") (String.Replacement "_") dep
        in String.contains (String.Pattern (pkgAlias <> ".")) declsStr && not (Array.elem ("gopurs/output/" <> dep) usedImports1)
      ) missingDeps
      usedImports = usedImports1 <> map (\dep -> "gopurs/output/" <> dep) injectedDeps
  in
  "package " <> packageName <> "\n\n" <>
  "import (\n" <>
  String.joinWith "\n" (map (\i -> 
      let pkg = Array.last (String.split (String.Pattern "/") i)
          pkgAlias = if i == "gopurs/output/gopurs_runtime" || i == "sync" then fromMaybe "" pkg else "pkg_" <> String.replaceAll (String.Pattern ".") (String.Replacement "_") (fromMaybe "" pkg)
      in "\t" <> pkgAlias <> " \"" <> i <> "\""
  ) usedImports) <> "\n" <>
  ")\n\n" <> declsStr
