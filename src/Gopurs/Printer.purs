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
    case Array.length props of
      0 -> "gopurs_runtime.RecordDict0()"
      1 -> case Array.index props 0 of
        Just (Tuple k0 v0) -> "gopurs_runtime.RecordDict1(\"" <> k0 <> "\", " <> printGoExpr v0 <> ")"
        Nothing -> ""
      2 -> case Tuple (Array.index props 0) (Array.index props 1) of
        Tuple (Just (Tuple k0 v0)) (Just (Tuple k1 v1)) ->
          "gopurs_runtime.RecordDict2(\"" <> k0 <> "\", \"" <> k1 <> "\", " <> printGoExpr v0 <> ", " <> printGoExpr v1 <> ")"
        _ -> ""
      3 -> case Tuple (Tuple (Array.index props 0) (Array.index props 1)) (Array.index props 2) of
        Tuple (Tuple (Just (Tuple k0 v0)) (Just (Tuple k1 v1))) (Just (Tuple k2 v2)) ->
          "gopurs_runtime.RecordDict3(\"" <> k0 <> "\", \"" <> k1 <> "\", \"" <> k2 <> "\", " <> printGoExpr v0 <> ", " <> printGoExpr v1 <> ", " <> printGoExpr v2 <> ")"
        _ -> ""
      4 -> case Tuple (Tuple (Array.index props 0) (Array.index props 1)) (Tuple (Array.index props 2) (Array.index props 3)) of
        Tuple (Tuple (Just (Tuple k0 v0)) (Just (Tuple k1 v1))) (Tuple (Just (Tuple k2 v2)) (Just (Tuple k3 v3))) ->
          "gopurs_runtime.RecordDict4(\"" <> k0 <> "\", \"" <> k1 <> "\", \"" <> k2 <> "\", \"" <> k3 <> "\", " <> printGoExpr v0 <> ", " <> printGoExpr v1 <> ", " <> printGoExpr v2 <> ", " <> printGoExpr v3 <> ")"
        _ -> ""
      5 -> case Tuple (Tuple (Array.index props 0) (Array.index props 1)) (Tuple (Tuple (Array.index props 2) (Array.index props 3)) (Array.index props 4)) of
        Tuple (Tuple (Just (Tuple k0 v0)) (Just (Tuple k1 v1))) (Tuple (Tuple (Just (Tuple k2 v2)) (Just (Tuple k3 v3))) (Just (Tuple k4 v4))) ->
          "gopurs_runtime.RecordDict5(\"" <> k0 <> "\", \"" <> k1 <> "\", \"" <> k2 <> "\", \"" <> k3 <> "\", \"" <> k4 <> "\", " <> printGoExpr v0 <> ", " <> printGoExpr v1 <> ", " <> printGoExpr v2 <> ", " <> printGoExpr v3 <> ", " <> printGoExpr v4 <> ")"
        _ -> ""
      _ ->
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
  GoConstructor tag args ->
    let
      len = Array.length args
    in
      if len <= 5 then
        "gopurs_runtime.Constructor" <> show len <> "(\"" <> tag <> "\"" <> (if len > 0 then ", " <> String.joinWith ", " (map printGoExpr args) else "") <> ")"
      else
        "gopurs_runtime.Constructor(\"" <> tag <> "\", []gopurs_runtime.Value{" <> String.joinWith ", " (map printGoExpr args) <> "})"
  GoConstructorAccess obj idx ->
    "gopurs_runtime.ConstructorGet(" <> printGoExpr obj <> ", " <> show idx <> ")"
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
