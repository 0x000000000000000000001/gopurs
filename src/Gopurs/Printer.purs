module Gopurs.Printer where

import Prelude
import Data.Array as Array
import Data.Maybe (Maybe(..), fromMaybe)
import Data.String as String
import Data.Tuple (Tuple(..))
import Gopurs.GoAst (GoExpr(..), GoDecl(..), GoBinding, GoFile, GoType(..), goTypeToStr, recordFieldName)
import Gopurs.Printer.Builder (Out, emit, emitMany, withOut)

foreign import escapeGoStringImpl :: String -> String

escapeGoString :: String -> String
escapeGoString = escapeGoStringImpl

quoteGoString :: String -> String
quoteGoString s = "\"" <> escapeGoString s <> "\""

-- | Écrit des éléments séparés par `separator` (sans séparateur avant le
-- | premier).
writeJoined :: forall a. String -> (Out -> a -> Out) -> Out -> Array a -> Out
writeJoined separator write = go true
  where
  go _ out [] = out
  go first out xs = case Array.uncons xs of
    Nothing -> out
    Just { head: x, tail } -> go false (write (if first then out else emit out separator) x) tail

writeSeparated :: String -> Out -> Array GoExpr -> Out
writeSeparated separator out exprs = writeJoined separator writeGoExpr out exprs

writeStatements :: Out -> Array GoExpr -> Out
writeStatements out stmts = writeJoined "\n" writeGoExpr out stmts

writeExprArray :: Out -> Array GoExpr -> Out
writeExprArray out exprs = writeSeparated ", " out exprs

-- A short declaration (or blank assignment) provides no type for nil.
-- Other expression positions retain the surrounding Go type context.
printBindingRhs :: GoExpr -> String
printBindingRhs expr = withOut (\out -> writeBindingRhs out expr)

writeBindingRhs :: Out -> GoExpr -> Out
writeBindingRhs out (GoConstructor _ structName typeArgs []) =
  let typeArgsStr = if Array.null typeArgs then "" else "[" <> String.joinWith ", " (map goTypeToStr typeArgs) <> "]"
  in emit out ("(*" <> structName <> typeArgsStr <> ")(nil)")
writeBindingRhs out expr = writeGoExpr out expr

printGoExpr :: GoExpr -> String
printGoExpr goExpr = withOut (\out -> writeGoExpr out goExpr)

writeGoExpr :: Out -> GoExpr -> Out
writeGoExpr out goExpr = case goExpr of
  GoVar name ->
    emit out name
  GoString s ->
    emitMany out [ "\"", escapeGoString s, "\"" ]
  GoInt i ->
    emit out (show i)
  GoCall f args ->
    emit (writeExprArray (emit (writeGoExpr out f) "(") args) ")"
  GoSelector obj field ->
    emitMany (writeGoExpr out obj) [ ".", field ]
  GoBlock stmts ->
    writeStatements out stmts
  GoReturn e ->
    writeGoExpr (emit out "return ") e
  GoAssign name e ->
    let
      out1 = emitMany out [ name, " := " ]
      out2 = writeBindingRhs out1 e
    in
      emitMany out2 [ "\n_ = ", name ]
  GoRecordDict goType props ->
    case goType of
      TypeRecord _ ->
        let
          out1 = emitMany out [ goTypeToStr goType, "{" ]
          out2 = writeExprArray out1 (map (\(Tuple _ v) -> v) props)
        in
          emit out2 "}"
      TypeStructPointer _ ->
        let
          capitalize s = String.toUpper (String.take 1 s) <> String.drop 1 s
          out1 = emitMany out [ "(&", String.drop 1 (goTypeToStr goType), "{" ]
          out2 = writeJoined ", " (\o (Tuple k v) -> writeGoExpr (emitMany o [ capitalize k, ": " ]) v) out1 props
        in
          emit out2 "})"
      _ ->
        case Array.length props of
          0 -> emit out "gopurs_runtime.RecordDict0()"
          1 -> case Array.index props 0 of
            Just (Tuple k0 v0) ->
              let out1 = emitMany out [ "gopurs_runtime.RecordDict1(", quoteGoString k0, ", " ]
              in emit (writeGoExpr out1 v0) ")"
            Nothing -> out
          2 -> case Tuple (Array.index props 0) (Array.index props 1) of
            Tuple (Just (Tuple k0 v0)) (Just (Tuple k1 v1)) ->
              let out1 = emitMany out [ "gopurs_runtime.RecordDict2(", quoteGoString k0, ", ", quoteGoString k1, ", " ]
              in emit (writeExprArray out1 [ v0, v1 ]) ")"
            _ -> out
          3 -> case Tuple (Tuple (Array.index props 0) (Array.index props 1)) (Array.index props 2) of
            Tuple (Tuple (Just (Tuple k0 v0)) (Just (Tuple k1 v1))) (Just (Tuple k2 v2)) ->
              let out1 = emitMany out [ "gopurs_runtime.RecordDict3(", quoteGoString k0, ", ", quoteGoString k1, ", ", quoteGoString k2, ", " ]
              in emit (writeExprArray out1 [ v0, v1, v2 ]) ")"
            _ -> out
          4 -> case Tuple (Tuple (Array.index props 0) (Array.index props 1)) (Tuple (Array.index props 2) (Array.index props 3)) of
            Tuple (Tuple (Just (Tuple k0 v0)) (Just (Tuple k1 v1))) (Tuple (Just (Tuple k2 v2)) (Just (Tuple k3 v3))) ->
              let out1 = emitMany out [ "gopurs_runtime.RecordDict4(", quoteGoString k0, ", ", quoteGoString k1, ", ", quoteGoString k2, ", ", quoteGoString k3, ", " ]
              in emit (writeExprArray out1 [ v0, v1, v2, v3 ]) ")"
            _ -> out
          5 -> case Tuple (Tuple (Array.index props 0) (Array.index props 1)) (Tuple (Tuple (Array.index props 2) (Array.index props 3)) (Array.index props 4)) of
            Tuple (Tuple (Just (Tuple k0 v0)) (Just (Tuple k1 v1))) (Tuple (Tuple (Just (Tuple k2 v2)) (Just (Tuple k3 v3))) (Just (Tuple k4 v4))) ->
              let out1 = emitMany out [ "gopurs_runtime.RecordDict5(", quoteGoString k0, ", ", quoteGoString k1, ", ", quoteGoString k2, ", ", quoteGoString k3, ", ", quoteGoString k4, ", " ]
              in emit (writeExprArray out1 [ v0, v1, v2, v3, v4 ]) ")"
            _ -> out
          _ ->
            let
              keysStr = String.joinWith ", " (map (\(Tuple k _) -> quoteGoString k) props)
              out1 = emitMany out [ "gopurs_runtime.RecordDict([]string{", keysStr, "}, []gopurs_runtime.Value{" ]
              out2 = writeExprArray out1 (map (\(Tuple _ v) -> v) props)
            in
              emit out2 "})"
  GoRecordUpdateDict orig props ->
    let
      keysStr = String.joinWith ", " (map (\(Tuple k _) -> quoteGoString k) props)
      writeGeneric o =
        let
          out1 = emitMany o [ "gopurs_runtime.RecordUpdateDict(" ]
          out2 = writeGoExpr out1 orig
          out3 = emitMany out2 [ ", []string{", keysStr, "}, []gopurs_runtime.Value{" ]
          out4 = writeExprArray out3 (map (\(Tuple _ v) -> v) props)
        in
          emit out4 "})"
    in
      case Array.length props of
        1 -> case props of
          [ Tuple k v ] ->
            let
              out1 = emitMany out [ "gopurs_runtime.RecordUpdate1(" ]
              out2 = writeGoExpr out1 orig
              out3 = emitMany out2 [ ", ", quoteGoString k, ", " ]
            in
              emit (writeGoExpr out3 v) ")"
          _ -> writeGeneric out
        2 -> case props of
          [ Tuple k1 v1, Tuple k2 v2 ] ->
            let
              out1 = emitMany out [ "gopurs_runtime.RecordUpdate2(" ]
              out2 = writeGoExpr out1 orig
              out3 = emitMany out2 [ ", ", quoteGoString k1, ", " ]
              out4 = writeGoExpr out3 v1
              out5 = emitMany out4 [ ", ", quoteGoString k2, ", " ]
            in
              emit (writeGoExpr out5 v2) ")"
          _ -> writeGeneric out
        3 -> case props of
          [ Tuple k1 v1, Tuple k2 v2, Tuple k3 v3 ] ->
            let
              out1 = emitMany out [ "gopurs_runtime.RecordUpdate3(" ]
              out2 = writeGoExpr out1 orig
              out3 = emitMany out2 [ ", ", quoteGoString k1, ", " ]
              out4 = writeGoExpr out3 v1
              out5 = emitMany out4 [ ", ", quoteGoString k2, ", " ]
              out6 = writeGoExpr out5 v2
              out7 = emitMany out6 [ ", ", quoteGoString k3, ", " ]
            in
              emit (writeGoExpr out7 v3) ")"
          _ -> writeGeneric out
        _ -> writeGeneric out
  GoRecordUpdateStatic orig size updates fallbackUpdates ->
    let
      structName = if size >= 6 then "gopurs_runtime.RecordData" else "gopurs_runtime.RecordData" <> show size
      typeVal = if size >= 6 then "gopurs_runtime.TypeRecordData" else "gopurs_runtime.TypeRecord" <> show size
      writeFallback o =
        let
          fallbackKeys = String.joinWith ", " (map (\(Tuple k _) -> quoteGoString k) fallbackUpdates)
          out1 = emitMany o [ "gopurs_runtime.RecordUpdateDict(origVal, []string{", fallbackKeys, "}, []gopurs_runtime.Value{" ]
        in
          emit (writeExprArray out1 (map (\(Tuple _ v) -> v) fallbackUpdates)) "})"
      writeUpdates withBracket o =
        writeJoined "\n"
          ( \acc (Tuple idx val) ->
              let
                head = if withBracket then "newVals[" else "clone.V"
                suffix = if withBracket then "] = " else " = "
              in
                writeGoExpr (emitMany acc [ head, show idx, suffix ]) val
          )
          o
          updates
    in
      if size >= 6 then
        let
          out1 = emit out "func() gopurs_runtime.Value {\norigVal := "
          out2 = writeGoExpr out1 orig
          out3 = emitMany out2 [ "\nif origVal.Type != ", typeVal, " {\nreturn " ]
          out4 = writeFallback out3
          out5 = emitMany out4 [ "\n}\nr := (*", structName, ")(origVal.UnsafePtr)\nnewVals := make([]gopurs_runtime.Value, len(r.Vals))\ncopy(newVals, r.Vals)\n" ]
          out6 = emitMany (writeUpdates true out5) [ "\n" ]
          out7 = emitMany out6 [ "newR := gopurs_runtime.RecordData{Keys: r.Keys, Vals: newVals}\nreturn gopurs_runtime.Value{Type: ", typeVal, ", UnsafePtr: unsafe.Pointer(&newR)}\n}()" ]
        in
          out7
      else
        let
          out1 = emit out "func() gopurs_runtime.Value {\norigVal := "
          out2 = writeGoExpr out1 orig
          out3 = emitMany out2 [ "\nif origVal.Type != ", typeVal, " {\nreturn " ]
          out4 = writeFallback out3
          out5 = emitMany out4 [ "\n}\nclone := *((*", structName, ")(origVal.UnsafePtr))\n" ]
          out6 = emitMany (writeUpdates false out5) [ "\n" ]
          out7 = emitMany out6 [ "return gopurs_runtime.Value{Type: ", typeVal, ", UnsafePtr: unsafe.Pointer(&clone)}\n}()" ]
        in
          out7
  GoRecordUpdateNative goType orig updates ->
    let
      out1 = emitMany out [ "func() ", goTypeToStr goType, " {\nclone := " ]
      out2 = writeGoExpr out1 orig
      out3 = writeJoined "\n" (\o (Tuple prop val) -> writeGoExpr (emitMany o [ "clone.", recordFieldName prop, " = " ]) val) (emit out2 "\n") updates
    in
      emit out3 "\nreturn clone\n}()"
  GoIIFE name binding body ->
    let
      out1 = emit out "func() gopurs_runtime.Value {\n"
      out2 = if name == "_" then
        writeBindingRhs (emitMany out1 [ name, " = " ]) binding
      else
        emitMany (writeBindingRhs (emitMany out1 [ name, " := " ]) binding) [ "\n_ = ", name ]
    in
      case body of
        GoBlock _ -> emit (writeGoExpr (emit out2 "\n") body) "\n}()"
        _ -> emit (writeGoExpr (emit out2 "\nreturn ") body) "\n}()"
  GoLetRec bindings body ->
    let
      out1 = emit out "func() gopurs_runtime.Value {\n"
      out2 = emitMany (writeJoined "\n" (\o (Tuple name _) -> emitMany o [ "var ", name, " gopurs_runtime.Value" ]) out1 bindings) [ "\n" ]
      out3 = emitMany (writeJoined "\n" (\o (Tuple name _) -> emitMany o [ "_ = ", name ]) out2 bindings) [ "\n" ]
      out4 = emitMany (writeJoined "\n" (\o (Tuple name expr) -> writeGoExpr (emitMany o [ name, " = " ]) expr) out3 bindings) [ "\n" ]
    in
      emit (writeGoExpr (emit out4 "return ") body) "\n}()"
  GoRecordAccess obj prop ->
    emitMany (writeGoExpr (emit out "gopurs_runtime.RecordGet(") obj) [ ", ", quoteGoString prop, ")" ]
  GoStructAccess obj prop ->
    emitMany (writeGoExpr out obj) [ ".", prop ]
  GoRecordAccessStatic obj size idx ->
    if size >= 6 then
      emitMany (writeGoExpr (emit out "((*gopurs_runtime.RecordData)(") obj) [ ".UnsafePtr)).Vals[", show idx, "]" ]
    else
      emitMany (writeGoExpr (emitMany out [ "((*gopurs_runtime.RecordData", show size, ")(" ]) obj) [ ".UnsafePtr)).V", show idx ]
  GoConstructor _ structName typeArgs args ->
    let typeArgsStr = if Array.length typeArgs > 0 then "[" <> String.joinWith ", " (map goTypeToStr typeArgs) <> "]" else ""
    in if Array.null args then
      emit out "nil"
    else
      let out1 = emitMany out [ "(&", structName, typeArgsStr, "{1, " ]
      in emit (writeExprArray out1 args) "})"
  GoConstructorDict tag args ->
    let len = Array.length args
    in if len <= 5 then
      let out1 = emitMany out [ "gopurs_runtime.Constructor", show len, "(\"", tag, "\"", (if len > 0 then ", " else "") ]
      in emit (writeExprArray out1 args) ")"
    else
      let out1 = emitMany out [ "gopurs_runtime.Constructor(\"", tag, "\", []gopurs_runtime.Value{" ]
      in emit (writeExprArray out1 args) "})"
  GoConstructorAccess obj structName typeArgs idx isNative ->
    if isNative then
      emitMany (writeGoExpr (emit out "(") obj) [ ").V", show idx ]
    else
      let typeArgsStr = if Array.length typeArgs > 0 then "[" <> String.joinWith ", " (map goTypeToStr typeArgs) <> "]" else ""
      in
        emitMany (writeGoExpr (emitMany out [ "(*", structName, typeArgsStr, ")(" ]) obj) [ ".UnsafePtr).V", show idx ]
  GoBranch branches def ->
    let
      out1 = emit out "func() gopurs_runtime.Value {\n"
      out2 = writeJoined "\n"
        ( \o (Tuple cond t) ->
            emit (writeGoExpr (emit (writeGoExpr (emitMany o [ "if (" ]) cond) ").IntVal != 0 {\nreturn ") t) "\n}"
        )
        out1
        branches
    in
      emit (writeGoExpr (emit out2 "\nreturn ") def) "\n}()"
  GoBinOp op left right ->
    let
      out1 = emitMany (writeGoExpr (emit out "(") left) [ ") ", op, " (" ]
    in
      emit (writeGoExpr out1 right) ")"
  GoPrefixOp op expr ->
    emit (writeGoExpr (emitMany out [ op, "(" ]) expr) ")"
  GoTypeAssertion expr t ->
    emitMany (writeGoExpr out expr) [ ".(", t, ")" ]
  GoIndex expr index ->
    emit (writeGoExpr (emitMany (writeGoExpr (emit out "(") expr) [ ")[" ]) index) "]"
  GoBoxStructPointer tag expr ->
    emit (writeGoExpr (emitMany out [ "gopurs_runtime.Value{Type: 9, IntVal: ", tag, ", UnsafePtr: unsafe.Pointer(" ]) expr) ")}"
  GoBoxIntArray expr ->
    let out1 = emitMany out [ "func() gopurs_runtime.Value {\n\t\t\t\t\tarr := " ]
    in
      emitMany (writeGoExpr out1 expr) [ "\n\t\t\t\t\tboxed := make([]gopurs_runtime.Value, len(arr))\n\t\t\t\t\tfor i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }\n\t\t\t\t\treturn gopurs_runtime.Array(boxed)\n\t\t\t\t}()" ]
  GoUnboxIntArray expr ->
    let out1 = emitMany out [ "func() []int64 {\n\t\t\t\t\tarr := *(*[]gopurs_runtime.Value)(" ]
    in
      emitMany (writeGoExpr out1 expr) [ ".UnsafePtr)\n\t\t\t\t\tunboxed := make([]int64, len(arr))\n\t\t\t\t\tfor i, v := range arr { unboxed[i] = v.IntVal }\n\t\t\t\t\treturn unboxed\n\t\t\t\t}()" ]
  GoFreshFilterArray expr ->
    writeGoExpr out expr
  GoRaw code ->
    emit out code.text
  GoFor label stmts ->
    let
      out1 = emitMany out [ label, ":\nfor {\nif false { continue ", label, " }\n" ]
    in
      emit (writeStatements out1 stmts) "\n}"
  GoForRange rangeStmt stmts ->
    let out1 = emitMany out [ "for ", rangeStmt, " {\n" ]
    in
      emit (writeStatements out1 stmts) "\n}"
  GoContinue label ->
    emitMany out [ "continue ", label ]
  GoIfElse cond trueStmts falseStmts ->
    let
      out1 = emit (writeGoExpr (emit out "if ") cond) " {\n"
      out2 = emit (writeStatements out1 trueStmts) "\n} else {\n"
    in
      emit (writeStatements out2 falseStmts) "\n}"
  GoMutate name expr ->
    writeGoExpr (emitMany out [ name, " = " ]) expr
  GoFuncBlock params stmts retType ->
    let
      out1 = emit out "func("
      out2 = writeParams out1 params
      out3 = emitMany out2 [ ") ", goTypeToStr retType, " {\n" ]
    in
      emit (writeGoExpr out3 (GoBlock stmts)) "\n}"
  GoFuncLit params stmts retExpr retType ->
    writeGoExpr out (GoFuncBlock params (stmts <> [ GoReturn retExpr ]) retType)
  GoStructValue adtName fields exprs ->
    let out1 = emitMany out [ goTypeToStr (TypeStructValue adtName fields), "{" ]
    in
      emit (writeExprArray out1 exprs) "}"

printGoDeclVar :: GoBinding -> String
printGoDeclVar binding = withOut (\out -> writeGoDeclVar out binding)

writeGoDeclVar :: Out -> GoBinding -> Out
writeGoDeclVar out { identifier, expression, goType } =
  let
    typeStr = goTypeToStr goType
    out1 = emitMany out
      [ "var cache_", identifier, " ", typeStr
      , "\nvar once_", identifier, " sync.Once\nfunc Get_", identifier, "() ", typeStr
      , " {\n\tonce_", identifier, ".Do(func() {\n\t\tcache_", identifier, " = "
      ]
    out2 = writeGoExpr out1 expression
  in
    emitMany out2 [ "\n\t})\n\treturn cache_", identifier, "\n}" ]

printGoDecl :: GoDecl -> String
printGoDecl decl = withOut (\out -> writeGoDecl out decl)

writeGoDecl :: Out -> GoDecl -> Out
writeGoDecl out = case _ of
  GoCachedValue binding -> writeGoDeclVar out binding
  GoStructDecl { name, typeParams, fields } ->
    let
      params = if Array.null typeParams then "" else "[" <> printParams typeParams <> "]"
      out1 = emitMany out [ "type ", name, params, " struct {\n\t" ]
      out2 = writeJoined "\n\t" (\o (Tuple field ty) -> emitMany o [ field, " ", goTypeToStr ty ]) out1 fields
    in
      emit out2 "\n}\n"
  GoFunctionDecl { name, params, result, body } ->
    let
      out1 = emitMany out [ "func ", name, "(" ]
      out2 = writeParams out1 params
      out3 = emitMany out2 [ ") ", goTypeToStr result, " {\n" ]
    in
      emit (writeGoExpr out3 body) "\n}"
  GoInitDecl body ->
    emit (writeGoExpr (emit out "func init() {\n") body) "\n}\n"
  GoForeignGetter { name, value } ->
    emitMany out [ "func ", name, "() gopurs_runtime.Value {\n\treturn ", value, "\n}" ]

printParams :: Array (Tuple String GoType) -> String
printParams params = withOut (\out -> writeParams out params)

writeParams :: Out -> Array (Tuple String GoType) -> Out
writeParams = writeJoined ", " (\o (Tuple name ty) -> emitMany o [ name, " ", goTypeToStr ty ])

printGoFile :: GoFile -> String
printGoFile file = withOut (\out -> writeGoFile out file)

writeGoFile :: Out -> GoFile -> Out
writeGoFile out { packageName, imports, declarationGroups } =
  let
    out1 = emitMany out [ "package ", packageName, "\n\nimport (\n" ]
    out2 = writeJoined "\n"
      ( \o path ->
          let alias = fromMaybe "" (Array.last (String.split (String.Pattern "/") path))
          in emitMany o [ "\t", alias, " \"", path, "\"" ]
      )
      out1
      imports
    out3 = emitMany out2 [ "\n)\n\n" ]
    writeGroup o group = writeJoined "\n\n" writeGoDecl o group
    out4 = writeJoined "\n\n" writeGroup out3 declarationGroups
  in
    emit out4 "\n"
