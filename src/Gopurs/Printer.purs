module Gopurs.Printer where

import Prelude
import Data.String as String
import Gopurs.GoAst (GoExpr(..), GoDecl, GoFile, GoType(..), goTypeToStr)
import Data.Tuple (Tuple(..))
import Data.Array as Array
import Data.Maybe (Maybe(..), fromMaybe)

foreign import escapeGoStringImpl :: String -> String
escapeGoString :: String -> String
escapeGoString = escapeGoStringImpl

printGoExpr :: GoExpr -> String
printGoExpr expr = case expr of
  GoVar name ->
    name
  GoString s ->
    "\"" <> escapeGoString s <> "\""
  GoInt i ->
    show i
  GoCall f args ->
    printGoExpr f <> "(" <> String.joinWith ", " (map printGoExpr args) <> ")"
  GoSelector obj field ->
    printGoExpr obj <> "." <> field
  GoFunc arg argType retType body -> 
    let
      flattenGoFunc (GoFunc a atype rtype b) acc = flattenGoFunc b (Array.snoc acc (Tuple a atype))
      flattenGoFunc b acc = Tuple acc b
      Tuple allArgs finalBody = flattenGoFunc body [Tuple arg argType]
      len = Array.length allArgs
      funcName = if len == 1 then "Func" else "Func" <> show len
      argStr = String.joinWith ", " (map (\(Tuple a t) -> a <> " " <> goTypeToStr t) allArgs)
    in
      if len > 10 then
        case body of
          GoBlock _ -> "gopurs_runtime.Func(func(" <> arg <> " " <> goTypeToStr argType <> ") " <> goTypeToStr retType <> " {\n" <> printGoExpr body <> "\n})"
          _ -> "gopurs_runtime.Func(func(" <> arg <> " " <> goTypeToStr argType <> ") " <> goTypeToStr retType <> " {\nreturn " <> printGoExpr body <> "\n})"
      else
        case finalBody of
          GoBlock _ -> "gopurs_runtime." <> funcName <> "(func(" <> argStr <> ") " <> goTypeToStr retType <> " {\n" <> printGoExpr finalBody <> "\n})"
          _ -> "gopurs_runtime." <> funcName <> "(func(" <> argStr <> ") " <> goTypeToStr retType <> " {\nreturn " <> printGoExpr finalBody <> "\n})"
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
      len = Array.length props
    in
      if len == 1 then
        case props of
          [Tuple k v] -> "gopurs_runtime.RecordUpdate1(" <> printGoExpr orig <> ", \"" <> k <> "\", " <> printGoExpr v <> ")"
          _ -> "gopurs_runtime.RecordUpdateDict(" <> printGoExpr orig <> ", []string{" <> keysStr <> "}, []gopurs_runtime.Value{" <> valsStr <> "})"
      else if len == 2 then
        case props of
          [Tuple k1 v1, Tuple k2 v2] -> "gopurs_runtime.RecordUpdate2(" <> printGoExpr orig <> ", \"" <> k1 <> "\", " <> printGoExpr v1 <> ", \"" <> k2 <> "\", " <> printGoExpr v2 <> ")"
          _ -> "gopurs_runtime.RecordUpdateDict(" <> printGoExpr orig <> ", []string{" <> keysStr <> "}, []gopurs_runtime.Value{" <> valsStr <> "})"
      else if len == 3 then
        case props of
          [Tuple k1 v1, Tuple k2 v2, Tuple k3 v3] -> "gopurs_runtime.RecordUpdate3(" <> printGoExpr orig <> ", \"" <> k1 <> "\", " <> printGoExpr v1 <> ", \"" <> k2 <> "\", " <> printGoExpr v2 <> ", \"" <> k3 <> "\", " <> printGoExpr v3 <> ")"
          _ -> "gopurs_runtime.RecordUpdateDict(" <> printGoExpr orig <> ", []string{" <> keysStr <> "}, []gopurs_runtime.Value{" <> valsStr <> "})"
      else
        "gopurs_runtime.RecordUpdateDict(" <> printGoExpr orig <> ", []string{" <> keysStr <> "}, []gopurs_runtime.Value{" <> valsStr <> "})"
  GoRecordUpdateStatic orig size updates fallbackUpdates ->
    let
      structName = if size >= 6 then "gopurs_runtime.RecordData" else "gopurs_runtime.RecordData" <> show size
      typeVal = if size >= 6 then "gopurs_runtime.TypeRecordData" else "gopurs_runtime.TypeRecord" <> show size
      
      fallbackKeys = String.joinWith ", " (map (\(Tuple k _) -> "\"" <> k <> "\"") fallbackUpdates)
      fallbackVals = String.joinWith ", " (map (\(Tuple _ v) -> printGoExpr v) fallbackUpdates)
      fallbackCall = "gopurs_runtime.RecordUpdateDict(origVal, []string{" <> fallbackKeys <> "}, []gopurs_runtime.Value{" <> fallbackVals <> "})"
    in
      if size >= 6 then
        let
          assignments = String.joinWith "\n" (map (\(Tuple idx val) -> "newVals[" <> show idx <> "] = " <> printGoExpr val) updates)
        in
          "func() gopurs_runtime.Value {\norigVal := " <> printGoExpr orig <> "\nif origVal.Type != " <> typeVal <> " {\nreturn " <> fallbackCall <> "\n}\nr := (*" <> structName <> ")(origVal.UnsafePtr)\nnewVals := make([]gopurs_runtime.Value, len(r.Vals))\ncopy(newVals, r.Vals)\n" <> assignments <> "\nnewR := gopurs_runtime.RecordData{Keys: r.Keys, Vals: newVals}\nreturn gopurs_runtime.Value{Type: " <> typeVal <> ", UnsafePtr: unsafe.Pointer(&newR)}\n}()"
      else
        let
          assignments = String.joinWith "\n" (map (\(Tuple idx val) -> "clone.V" <> show idx <> " = " <> printGoExpr val) updates)
        in
          "func() gopurs_runtime.Value {\norigVal := " <> printGoExpr orig <> "\nif origVal.Type != " <> typeVal <> " {\nreturn " <> fallbackCall <> "\n}\nclone := *((*" <> structName <> ")(origVal.UnsafePtr))\n" <> assignments <> "\nreturn gopurs_runtime.Value{Type: " <> typeVal <> ", UnsafePtr: unsafe.Pointer(&clone)}\n}()"
  GoIIFE name binding body ->
    let assignment = if name == "_" then name <> " = " <> printGoExpr binding else name <> " := " <> printGoExpr binding <> "\n_ = " <> name
    in case body of
      GoBlock _ -> "func() gopurs_runtime.Value {\n" <> assignment <> "\n" <> printGoExpr body <> "\n}()"
      _ -> "func() gopurs_runtime.Value {\n" <> assignment <> "\nreturn " <> printGoExpr body <> "\n}()"
  GoLetRec bindings body ->
    "func() gopurs_runtime.Value {\n" <>
    String.joinWith "\n" (map (\(Tuple name _) -> "var " <> name <> " gopurs_runtime.Value") bindings) <> "\n" <>
    String.joinWith "\n" (map (\(Tuple name _) -> "_ = " <> name) bindings) <> "\n" <>
    String.joinWith "\n" (map (\(Tuple name expr) -> name <> " = " <> printGoExpr expr) bindings) <> "\n" <>
    "return " <> printGoExpr body <> "\n}()"
  GoRecordAccess obj prop ->
    "gopurs_runtime.RecordGet(" <> printGoExpr obj <> ", \"" <> prop <> "\")"
  GoStructAccess obj prop ->
    printGoExpr obj <> "." <> prop
  GoRecordAccessStatic obj size idx ->
    if size >= 6 then
      "((*gopurs_runtime.RecordData)(" <> printGoExpr obj <> ".UnsafePtr)).Vals[" <> show idx <> "]"
    else
      "((*gopurs_runtime.RecordData" <> show size <> ")(" <> printGoExpr obj <> ".UnsafePtr)).V" <> show idx
  GoConstructor hashStr structName typeArgs args ->
    let typeArgsStr = ""
    in if Array.null args then
      "nil"
    else
      "(&" <> structName <> typeArgsStr <> "{1, " <> String.joinWith ", " (map printGoExpr args) <> "})"
  GoConstructorDict tag args ->
    let len = Array.length args
    in if len <= 5 then
         "gopurs_runtime.Constructor" <> show len <> "(\"" <> tag <> "\"" <> (if len > 0 then ", " else "") <> String.joinWith ", " (map printGoExpr args) <> ")"
       else
         "gopurs_runtime.Constructor(\"" <> tag <> "\", []gopurs_runtime.Value{" <> String.joinWith ", " (map printGoExpr args) <> "})"
  GoConstructorAccess obj structName typeArgs idx isNative ->
    if isNative then
      "(" <> printGoExpr obj <> ").V" <> show idx
    else
      let typeArgsStr = ""
      in "(*" <> structName <> typeArgsStr <> ")(" <> printGoExpr obj <> ".UnsafePtr).V" <> show idx
  GoBranch branches def ->
    "func() gopurs_runtime.Value {\n" <>
    String.joinWith "\n" (map (\(Tuple cond t) -> "if (" <> printGoExpr cond <> ").IntVal != 0 {\nreturn " <> printGoExpr t <> "\n}") branches) <>
    "\nreturn " <> printGoExpr def <> "\n}()"
  GoBinOp op left right ->
    "(" <> printGoExpr left <> ") " <> op <> " (" <> printGoExpr right <> ")"
  GoPrefixOp op expr ->
    op <> "(" <> printGoExpr expr <> ")"
  GoTypeAssertion expr t ->
    printGoExpr expr <> ".(" <> t <> ")"
  GoIndex expr index ->
    "(" <> printGoExpr expr <> ")[" <> printGoExpr index <> "]"
  GoRaw raw ->
    raw
  GoFor label stmts ->
    label <> ":\nfor {\nif false { continue " <> label <> " }\n" <> String.joinWith "\n" (map printGoExpr stmts) <> "\n}"
  GoForRange rangeStmt stmts ->
    "for " <> rangeStmt <> " {\n" <> String.joinWith "\n" (map printGoExpr stmts) <> "\n}"
  GoContinue label ->
    "continue " <> label
  GoIfElse cond trueStmts falseStmts ->
    "if " <> printGoExpr cond <> " {\n" <>
      String.joinWith "\n" (map printGoExpr trueStmts) <>
    "\n} else {\n" <>
      String.joinWith "\n" (map printGoExpr falseStmts) <>
    "\n}"
  GoMutate name expr ->
    name <> " = " <> printGoExpr expr
  GoFuncLit params stmts retExpr retType ->
    "func(" <> String.joinWith ", " (map (\(Tuple p goT) -> p <> " " <> goTypeToStr goT) params) <> ") " <> goTypeToStr retType <> " {\n" <> printGoExpr (GoBlock (stmts <> [ GoReturn retExpr ])) <> "\n}"

printGoDeclVar :: GoDecl -> String
printGoDeclVar { identifier, expression, goType } =
  let typeStr = goTypeToStr goType
  in
  "var cache_" <> identifier <> " " <> typeStr <> "\n" <>
  "var once_" <> identifier <> " sync.Once\n" <>
  "func Get_" <> identifier <> "() " <> typeStr <> " {\n" <>
  "\tonce_" <> identifier <> ".Do(func() {\n" <>
  "\t\tcache_" <> identifier <> " = " <> printGoExpr expression <> "\n" <>
  "\t})\n" <>
  "\treturn cache_" <> identifier <> "\n" <>
  "}"

printGoFile :: GoFile -> String
printGoFile { packageName, imports, decls, rawDecls, foreigns } =
  let declsStr = String.joinWith "\n\n" (map printGoDeclVar decls) <> "\n\n" <> String.joinWith "\n\n" rawDecls <> "\n\n" <> String.joinWith "\n\n" (map (\f -> "func Get_" <> f.pursName <> "() gopurs_runtime.Value {\n\treturn " <> f.goName <> "\n}") foreigns) <> "\n"
      unsafeImport = if String.contains (String.Pattern "unsafe.") declsStr then ["unsafe"] else []
      hasDecls = Array.length decls > 0
      hasForeigns = Array.length foreigns > 0
      needsRuntime = hasDecls || hasForeigns || String.contains (String.Pattern "gopurs_runtime.") declsStr
      usedImports = (if needsRuntime then ["gopurs/output/gopurs_runtime"] else []) <> (if hasDecls then ["sync"] else []) <> unsafeImport
  in
  "package " <> packageName <> "\n\n" <>
  "import (\n" <>
  String.joinWith "\n" (map (\i -> 
      let pkg = Array.last (String.split (String.Pattern "/") i)
          pkgAlias = fromMaybe "" pkg
      in "\t" <> pkgAlias <> " \"" <> i <> "\""
  ) usedImports) <> "\n" <>
  ")\n\n" <> declsStr
