module Gopurs.FfiBridge
  ( generateFfiBridge
  , ffiFunctionInfos
  , ffiValueWorkers
  ) where

import Prelude

import Data.Array as Array
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Newtype (unwrap)
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Gopurs.FfiTypes (TypeNode(..), FfiDecl)
import Gopurs.CodegenState (FunctionInfo)
import Gopurs.GoAst (GoType(..), capitalize, sanitizeName)
import Gopurs.GoTypes (isClosedRowTail, printExprType)
import PureScript.Backend.Optimizer.CoreFn (DataDecl, ExprType(..), Ident(..))

printTypeNode :: TypeNode -> String
printTypeNode (TNamed n) = n
printTypeNode (TFunc args ret) =
  let
    argsStr = String.joinWith ", " (map printTypeNode args)
    retStr = case ret of
      Nothing -> ""
      Just r -> " " <> printTypeNode r
  in
    "func(" <> argsStr <> ")" <> retStr
printTypeNode (TArray elem) = "[]" <> printTypeNode elem
printTypeNode (TMap k v) = "map[" <> printTypeNode k <> "]" <> printTypeNode v
printTypeNode (TUnknown s) = s

exprTypeToDummyTypeNode :: ExprType -> TypeNode
exprTypeToDummyTypeNode (Func args ret) = TFunc (map exprTypeToDummyTypeNode args) (Just (exprTypeToDummyTypeNode ret))
exprTypeToDummyTypeNode (Array elem) = TArray (exprTypeToDummyTypeNode elem)
exprTypeToDummyTypeNode (Record _) = TMap (TNamed "string") (TNamed "any")
exprTypeToDummyTypeNode _ = TNamed "any"

getTastReturnType :: ExprType -> Maybe ExprType
getTastReturnType (Func _ ret) = Just ret
getTastReturnType _ = Nothing

resolveNewtype :: Array DataDecl -> ExprType -> ExprType
resolveNewtype dataDecls (ADT fullName path args) =
  let
    typeName = fullName
    mbDecl = Array.find (\d -> d.name == typeName || (Array.length path > 0 && d.name == fromMaybe "" (Array.last path))) dataDecls
  in
    case mbDecl of
      Just decl ->
        if Array.length decl.constructors == 1 then
          case Array.head decl.constructors of
            Just ctor ->
              if Array.length ctor.fields == 1 then
                case Array.head ctor.fields of
                  Just fieldT -> resolveNewtype dataDecls fieldT
                  Nothing -> ADT fullName path args
              else ADT fullName path args
            Nothing -> ADT fullName path args
        else ADT fullName path args
      Nothing -> ADT fullName path args
resolveNewtype dataDecls (Func fArgs ret) = Func (map (resolveNewtype dataDecls) fArgs) (resolveNewtype dataDecls ret)
resolveNewtype dataDecls (Array elem) = Array (resolveNewtype dataDecls elem)
resolveNewtype dataDecls (Record row) = Record (resolveNewtype dataDecls row)
resolveNewtype dataDecls (Row fields tail) = Row (map (\(Tuple k v) -> Tuple k (resolveNewtype dataDecls v)) fields) (map (resolveNewtype dataDecls) tail)
resolveNewtype dataDecls (TypeApp c args) = TypeApp (resolveNewtype dataDecls c) (map (resolveNewtype dataDecls) args)
resolveNewtype dataDecls (ForAll vars body) = ForAll vars (resolveNewtype dataDecls body)
resolveNewtype dataDecls (ConstrainedType constraints body) = ConstrainedType (map (\(Tuple c a) -> Tuple c (map (resolveNewtype dataDecls) a)) constraints) (resolveNewtype dataDecls body)
resolveNewtype _ other = other

flattenFuncArgs :: ExprType -> Array ExprType
flattenFuncArgs (Func args ret) = args <> flattenFuncArgs ret
flattenFuncArgs _ = []

unwrapValueToFunc :: Array DataDecl -> TypeNode -> Maybe ExprType -> String -> Int -> Int -> String
unwrapValueToFunc dataDecls (TFunc args ret) mbTast valName depth _ =
  let
    paramsArr = Array.mapWithIndex (\cidx atype -> "p" <> show depth <> "_" <> show cidx <> " " <> printTypeNode atype) args
    applyArgsArr = Array.mapWithIndex
      ( \cidx atype ->
          case atype of
            TNamed "gopurs_runtime.Value" -> "p" <> show depth <> "_" <> show cidx
            TFunc _ _ ->
              let
                wrapped = wrapReturn dataDecls atype (mbTast >>= \tast -> getTastArgType tast cidx) ("p" <> show depth <> "_" <> show cidx)
              in
                String.replaceAll (Pattern "\n") (Replacement "\n\t\t") wrapped
            _ -> "gopurs_runtime.Box(p" <> show depth <> "_" <> show cidx <> ")"
      )
      args
    params = String.joinWith ", " paramsArr
    applyArgs = String.joinWith ", " applyArgsArr

    applyCall =
      if Array.length args == 1 then
        "gopurs_runtime.Apply(" <> valName <> ", " <> applyArgs <> ")"
      else if Array.length args > 1 then
        "gopurs_runtime.Apply" <> show (Array.length args) <> "(" <> valName <> ", " <> applyArgs <> ")"
      else
        "gopurs_runtime.Apply(" <> valName <> ", gopurs_runtime.Value{})"

  in
    case ret of
      Nothing ->
        "func(" <> params <> ") {\n\t\t" <> applyCall <> "\n\t}"
      Just (TArray elem) | printTypeNode elem /= "gopurs_runtime.Value" ->
        let
          elemType = printTypeNode elem
          retStr = printTypeNode (TArray elem)
        in
          "func(" <> params <> ") " <> retStr <> " {\n\t\tinner_res" <> show depth <> " := " <> applyCall <> "\n\t\tres_arr" <> show depth <> " := *(*[]gopurs_runtime.Value)(inner_res" <> show depth <> ".UnsafePtr)\n\t\tres_go" <> show depth <> " := make(" <> retStr <> ", len(res_arr" <> show depth <> "))\n\t\tfor i, v := range res_arr" <> show depth <> " { res_go" <> show depth <> "[i] = gopurs_runtime.Unbox[" <> elemType <> "](v) }\n\t\treturn res_go" <> show depth <> "\n\t}"
      Just (TNamed "any") -> "func(" <> params <> ") any {\n\t\treturn " <> applyCall <> "\n\t}"
      Just (TNamed "interface{}") -> "func(" <> params <> ") interface{} {\n\t\treturn " <> applyCall <> "\n\t}"
      Just (TNamed "gopurs_runtime.Value") -> "func(" <> params <> ") gopurs_runtime.Value {\n\t\treturn " <> applyCall <> "\n\t}"
      Just f@(TFunc _ _) ->
        let
          innerUnwrap = unwrapValueToFunc dataDecls f (mbTast >>= getTastReturnType) ("inner_res" <> show depth) (depth + 1) 0
        in
          "func(" <> params <> ") " <> printTypeNode f <> " {\n\t\tinner_res" <> show depth <> " := " <> applyCall <> "\n\t\treturn " <> innerUnwrap <> "\n\t}"
      Just r ->
        "func(" <> params <> ") " <> printTypeNode r <> " {\n\t\tinner_res" <> show depth <> " := " <> applyCall <> "\n\t\treturn gopurs_runtime.Unbox[" <> printTypeNode r <> "](inner_res" <> show depth <> ")\n\t}"
unwrapValueToFunc dataDecls (TNamed anyT) mbTast valName depth cidx | anyT == "any" || anyT == "interface{}" || anyT == "gopurs_runtime.Value" =
  let
    resolvedTast = map (resolveNewtype dataDecls) mbTast
  in
    case resolvedTast of
      Just (Record (Row fields tail)) | isClosedRowTail tail ->
        let
          fieldStr = Array.mapWithIndex
            ( \i (Tuple fK fT) ->
                "\t\t\t\tres_map[\"" <> fK <> "\"] = " <> unwrapValueToFunc dataDecls (TNamed "any") (Just fT) ("_raw[\"" <> fK <> "\"]") (depth + 1) i
            )
            fields
        in
          "func() map[string]any {\n\t\t\t_raw := gopurs_runtime.RecordToMap(" <> valName <> ")\n\t\t\tres_map := make(map[string]any)\n" <> String.joinWith "\n" fieldStr <> "\n\t\t\treturn res_map\n\t\t}()"
      Just f@(Func _ _) ->
        unwrapValueToFunc dataDecls (exprTypeToDummyTypeNode f) resolvedTast valName depth cidx
      _ -> "gopurs_runtime.Unbox[" <> anyT <> "](" <> valName <> ")"
unwrapValueToFunc _ t _ valName _ _ = "gopurs_runtime.Unbox[" <> printTypeNode t <> "](" <> valName <> ")"

boxFfiValue :: TypeNode -> String -> String
boxFfiValue (TNamed "gopurs_runtime.Value") valName = valName
boxFfiValue (TNamed "int64") valName = "gopurs_runtime.Int(" <> valName <> ")"
boxFfiValue (TNamed "int") valName = "gopurs_runtime.Int(int64(" <> valName <> "))"
boxFfiValue _ valName = "gopurs_runtime.Box(" <> valName <> ")"

wrapReturn :: Array DataDecl -> TypeNode -> Maybe ExprType -> String -> String
-- A runtime Value already has the required representation, including when its
-- PureScript type is a record or function. Native callbacks with this exact
-- signature likewise need no argument/result conversion or forwarding closure.
wrapReturn _ (TNamed "gopurs_runtime.Value") _ valName = valName
wrapReturn _ (TFunc [ TNamed "gopurs_runtime.Value" ] (Just (TNamed "gopurs_runtime.Value"))) _ valName =
  "gopurs_runtime.Func(" <> valName <> ")"
wrapReturn dataDecls (TFunc args ret) mbTast valName =
  let
    innerT = ret
    argT = Array.head args

    genInner val innerWrap =
      "gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {\n\t\t\t" <> val <> "\n\t\t\treturn " <> innerWrap <> "\n\t\t})"

    genInnerArg val innerWrap argUnwrap =
      "gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {\n\t\t\t" <> val <> "(" <> argUnwrap <> ")\n\t\t\treturn " <> innerWrap <> "\n\t\t})"
  in
    case argT of
      Nothing ->
        case innerT of
          Nothing -> genInner (valName <> "()") "gopurs_runtime.Value{}"
          Just r -> genInner ("inner_res := " <> valName <> "()") (wrapReturn dataDecls r (mbTast >>= getTastReturnType) "inner_res")
      Just a ->
        let
          argUnwrap = case a of
            TNamed "any" -> "arg"
            TNamed "interface{}" -> "arg"
            TNamed "gopurs_runtime.Value" -> "arg"
            f@(TFunc _ _) -> String.replaceAll (Pattern "\n") (Replacement "\n\t\t\t") (unwrapValueToFunc dataDecls f (mbTast >>= \tast -> getTastArgType tast 0) "arg" 99 0)
            _ -> "gopurs_runtime.Unbox[" <> printTypeNode a <> "](arg)"
        in
          case innerT of
            Nothing -> genInnerArg valName "gopurs_runtime.Value{}" argUnwrap
            Just r -> genInnerArg ("inner_res := " <> valName) (wrapReturn dataDecls r (mbTast >>= getTastReturnType) "inner_res") argUnwrap
wrapReturn _ (TArray elem) _ valName | printTypeNode elem /= "gopurs_runtime.Value" =
  "func() gopurs_runtime.Value {\n\t\t\tres_arr := make([]gopurs_runtime.Value, len(" <> valName <> "))\n\t\t\tfor i, v := range " <> valName <> " { res_arr[i] = " <> boxFfiValue elem "v" <> " }\n\t\t\treturn gopurs_runtime.Array(res_arr)\n\t\t}()"
wrapReturn dataDecls (TMap _ _) (Just (Record (Row fields tail))) valName | isClosedRowTail tail =
  let
    fieldStr = Array.mapWithIndex
      ( \_ (Tuple fK fT) ->
          "\t\t\t\tres_map[\"" <> fK <> "\"] = " <> wrapReturn dataDecls (TNamed "any") (Just fT) ("_raw[\"" <> fK <> "\"]")
      )
      fields
  in
    "func() gopurs_runtime.Value {\n\t\t\t_raw := " <> valName <> "\n\t\t\tres_map := make(map[string]gopurs_runtime.Value)\n" <> String.joinWith "\n" fieldStr <> "\n\t\t\treturn gopurs_runtime.Record(res_map)\n\t\t}()"
wrapReturn dataDecls (TMap _ _) mbTast valName =
  let
    resolvedTast = map (resolveNewtype dataDecls) mbTast
    isOpaque = case resolvedTast of
      Just (Record _) -> false
      _ -> true
  in
    if isOpaque then
      "gopurs_runtime.Any(" <> valName <> ")"
    else
      "func() gopurs_runtime.Value {\n\t\t\tres_map := make(map[string]gopurs_runtime.Value)\n\t\t\tfor k, v := range " <> valName <> " { res_map[k] = gopurs_runtime.Box(v) }\n\t\t\treturn gopurs_runtime.Record(res_map)\n\t\t}()"
wrapReturn dataDecls (TNamed anyT) mbTast valName | anyT == "any" || anyT == "interface{}" =
  let
    resolvedTast = map (resolveNewtype dataDecls) mbTast
  in
    case resolvedTast of
      Just (Record (Row fields tail)) | isClosedRowTail tail ->
        let
          fieldStr = Array.mapWithIndex
            ( \_ (Tuple fK fT) ->
                "\t\t\t\tres_map[\"" <> fK <> "\"] = " <> wrapReturn dataDecls (TNamed "any") (Just fT) ("_raw[\"" <> fK <> "\"]")
            )
            fields
        in
          "func() gopurs_runtime.Value {\n\t\t\t_raw := gopurs_runtime.RecordToMap(gopurs_runtime.Box(" <> valName <> "))\n\t\t\tres_map := make(map[string]gopurs_runtime.Value)\n" <> String.joinWith "\n" fieldStr <> "\n\t\t\treturn gopurs_runtime.Record(res_map)\n\t\t}()"
      Just f@(Func _ _) ->
        let
          fArgs = flattenFuncArgs f
          arity = Array.length fArgs

          genWrap args remaining depth =
            if remaining == 0 then
              let
                castArgs = String.joinWith ", " (Array.replicate arity "any")
                castType = "func(" <> castArgs <> ") any"

                invokeArgs = Array.mapWithIndex
                  ( \i argT ->
                      let
                        pName = "p" <> show (depth - arity + i)
                      in
                        case argT of
                          Func _ _ ->
                            let
                              cbArgs = flattenFuncArgs argT
                              cbParams = Array.mapWithIndex (\ci _ -> "cb_arg" <> show ci) cbArgs
                              cbParamsDecl = String.joinWith ", " (map (\n -> n <> " any") cbParams)
                              applyChain = Array.foldl (\acc a -> "gopurs_runtime.Apply(" <> acc <> ", gopurs_runtime.Box(" <> a <> "))") pName cbParams
                            in
                              "func(" <> cbParamsDecl <> ") any { return " <> applyChain <> " }"
                          _ -> "gopurs_runtime.Unbox[any](" <> pName <> ")"
                  )
                  fArgs

                invokeCall = "fn(" <> String.joinWith ", " invokeArgs <> ")"
                fallbackChain = Array.foldl (\acc p -> "gopurs_runtime.Apply(" <> acc <> ", " <> p <> ")") ("(" <> valName <> ".(gopurs_runtime.Value))") (Array.mapWithIndex (\i _ -> "p" <> show (depth - arity + i)) fArgs)

              in
                "func() gopurs_runtime.Value {\n\t\t\t\tif fn, ok := " <> valName <> ".(" <> castType <> "); ok {\n\t\t\t\t\treturn gopurs_runtime.Box(" <> invokeCall <> ")\n\t\t\t\t}\n\t\t\t\treturn " <> fallbackChain <> "\n\t\t\t}()"
            else
              let
                currArg = "p" <> show depth
              in
                "gopurs_runtime.Func(func(" <> currArg <> " gopurs_runtime.Value) gopurs_runtime.Value {\n\t\t\treturn " <> genWrap args (remaining - 1) (depth + 1) <> "\n\t\t})"

        in
          genWrap fArgs arity 0
      _ -> "gopurs_runtime.Box(" <> valName <> ")"
wrapReturn _ typ _ valName = boxFfiValue typ valName

getTastArgType :: ExprType -> Int -> Maybe ExprType
getTastArgType (Func args _) i = Array.index args i
getTastArgType _ _ = Nothing

isStandardPursFunc :: TypeNode -> Boolean
isStandardPursFunc (TFunc args ret) =
  let
    argsAny = Array.all (\a -> printTypeNode a == "any" || printTypeNode a == "interface{}" || printTypeNode a == "gopurs_runtime.Value") args
    retStr = case ret of
      Nothing -> ""
      Just r -> printTypeNode r
  in
    if Array.length args == 0 then
      retStr == "" || retStr == "bool" || retStr == "int" || retStr == "int64" || retStr == "string" || retStr == "float64" || retStr == "gopurs_runtime.Value" || retStr == "any" || retStr == "interface{}"
    else if argsAny then
      case ret of
        Nothing -> true
        Just r@(TFunc _ _) -> isStandardPursFunc r
        Just _ -> retStr == "any" || retStr == "interface{}" || retStr == "gopurs_runtime.Value"
    else
      false
isStandardPursFunc _ = false

generateWrapperFunc :: Array DataDecl -> FfiDecl -> Maybe ExprType -> String
generateWrapperFunc dataDecls declaration mbTast =
  let
    -- Go type parameters are instantiated with Value at the call below. Apply
    -- the same substitution inside callbacks and containers before adapting
    -- their arguments and results.
    instantiate = case _ of
      TNamed name | Array.elem name declaration.typeParams -> TNamed "gopurs_runtime.Value"
      TFunc args ret -> TFunc (map instantiate args) (map instantiate ret)
      TArray item -> TArray (instantiate item)
      TMap key value -> TMap (instantiate key) (instantiate value)
      other -> other
    d = declaration { args = map instantiate declaration.args, ret = map instantiate declaration.ret }
    tastComment = case mbTast of
      Just tast -> "// TAST: " <> printExprType tast <> "\n"
      Nothing -> "// TAST: Unknown\n"
  in
    tastComment <>
      if d.isVar then
        "gopurs_runtime.Box(" <> d.name <> ")"
      else
        let
          arity = Array.length d.args
          funcConstructor = if arity > 1 then "Func" <> show arity else "Func"

          boxedArgs =
            if arity == 0 then
              "_ gopurs_runtime.Value"
            else if arity == 1 then
              "arg0 gopurs_runtime.Value"
            else
              String.joinWith ", " (Array.mapWithIndex (\i _ -> "arg" <> show i <> " gopurs_runtime.Value") d.args)

          callFunc =
            if Array.length d.typeParams > 0 then
              d.name <> "[" <> String.joinWith ", " (map (const "gopurs_runtime.Value") d.typeParams) <> "]"
            else d.name

          processArg i t =
            let
              typStr = printTypeNode t
              elemType = String.drop 2 typStr
              tastArg = mbTast >>= \tast -> getTastArgType tast i
            in
              case t of
                TFunc _ _ ->
                  let
                    isOpaque = case tastArg of
                      Just (ADT _ _ _) -> true
                      Just Any -> true
                      Just (TypeVar _) -> true
                      _ -> false
                  in
                    if isOpaque && not (isStandardPursFunc t) then
                      [ "\tgo_arg" <> show i <> " := (*(*any)(arg" <> show i <> ".UnsafePtr)).(" <> typStr <> ")" ]
                    else
                      let
                        unwrapStr = unwrapValueToFunc dataDecls t tastArg ("arg" <> show i) 0 0
                        indented = String.replaceAll (Pattern "\n") (Replacement "\n\t") unwrapStr
                      in
                        [ "\tgo_arg" <> show i <> " := " <> indented ]
                TArray _ | typStr /= "[]gopurs_runtime.Value" ->
                  let
                    et = if elemType == "any" then "interface{}" else elemType
                  in
                    [ "\targ" <> show i <> "_arr := *(*[]gopurs_runtime.Value)(arg" <> show i <> ".UnsafePtr)"
                    , "\tgo_arg" <> show i <> " := make(" <> typStr <> ", len(arg" <> show i <> "_arr))"
                    , if et == "interface{}" then
                        "\tfor i, v := range arg" <> show i <> "_arr { go_arg" <> show i <> "[i] = v }"
                      else
                        "\tfor i, v := range arg" <> show i <> "_arr { go_arg" <> show i <> "[i] = gopurs_runtime.Unbox[" <> et <> "](v) }"
                    ]
                TNamed "any" -> [ "\tgo_arg" <> show i <> " := arg" <> show i ]
                TNamed "interface{}" -> [ "\tgo_arg" <> show i <> " := arg" <> show i ]
                TNamed "gopurs_runtime.Value" -> [ "\tgo_arg" <> show i <> " := arg" <> show i ]
                TMap _ _ ->
                  let
                    et = String.drop (String.indexOf (Pattern "]") typStr # fromMaybe 0 # add 1) typStr
                    resolvedTast = map (resolveNewtype dataDecls) tastArg
                    isOpaque = case resolvedTast of
                      Just (Record _) -> false
                      _ -> true
                  in
                    if isOpaque then
                      [ "\tgo_arg" <> show i <> " := gopurs_runtime.UnboxObject(arg" <> show i <> ")" ]
                    else if et == "any" || et == "interface{}" then
                      [ "\targ" <> show i <> "_map := gopurs_runtime.RecordToMap(arg" <> show i <> ")"
                      , "\tgo_arg" <> show i <> " := make(" <> typStr <> ")"
                      , "\tfor k, v := range arg" <> show i <> "_map { go_arg" <> show i <> "[k] = v }"
                      ]
                    else
                      [ "\tgo_arg" <> show i <> " := arg" <> show i <> ".PtrVal().(" <> typStr <> ")" ]
                _ -> [ "\tgo_arg" <> show i <> " := gopurs_runtime.Unbox[" <> typStr <> "](arg" <> show i <> ")" ]

          argsCode = Array.concat (Array.mapWithIndex processArg d.args)
          callArgs = String.joinWith ", " (Array.mapWithIndex (\i _ -> "go_arg" <> show i) d.args)

          retCode = case d.ret of
            Nothing ->
              [ "\t" <> callFunc <> "(" <> callArgs <> ")"
              , "\treturn gopurs_runtime.Value{}"
              ]
            Just r ->
              let
                wrapCode = wrapReturn dataDecls r (mbTast >>= getTastReturnType) "go_res"
                indentedWrap = String.replaceAll (Pattern "\n") (Replacement "\n\t") wrapCode
              in
                [ "\tgo_res := " <> callFunc <> "(" <> callArgs <> ")"
                , "\treturn " <> indentedWrap
                ]

          fullCode = "gopurs_runtime." <> funcConstructor <> "(func(" <> boxedArgs <> ") gopurs_runtime.Value {\n"
            <> String.joinWith "\n" argsCode
            <> "\n"
            <> String.joinWith "\n" retCode
            <> "\n"
            <>
              "})"
        in
          fullCode

generateFfiBridge :: String -> Array DataDecl -> Array FfiDecl -> Array (Tuple Ident (Maybe ExprType)) -> String
generateFfiBridge modNameStr dataDecls decls foreigns =
  String.joinWith "\n" (map genBridge foreigns)
  where
  genBridge (Tuple ident mbTast) =
    let
      pursName = unwrap ident
      sanitized = sanitizeName pursName
      capName = capitalize sanitized
      exportName = "_Gopurs_" <> modNameStr <> "_" <> capName

      fallback1 = modNameStr <> "_" <> capitalize pursName
      fallback2 = fallback1 <> "_"

      findDecl n = Array.find (\d -> d.name == n) decls

      match = case findDecl fallback1 of
        Just d -> Just d
        Nothing -> findDecl fallback2
    in
      case match of
        Nothing ->
          "var " <> exportName <> " = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value { panic(\"FFI not implemented: " <> pursName <> "\"); return gopurs_runtime.Value{} })"
        Just d ->
          "var " <> exportName <> " = " <> generateWrapperFunc dataDecls d mbTast


-- | Infos d'appel natif pour les FFI d'un module : le worker Go typé
-- | (`fullName`, tel qu'il est embarqué dans `<Module>_ffi.go`), ses
-- | paramètres/résultat natifs et son arité. `directFunction`/`nativeCall`
-- | s'en servent pour abaisser les appels saturés en appels directs, sans
-- | passer par l'enveloppe `_Gopurs_*` (qui reste utilisée pour les usages
-- | comme valeur).
-- |
-- | Prudence : seules les FFI **pures**, à résultat **concret**, dont toute
-- | la signature native se convertit sans perte sont abaissées. Les FFI
-- | `Effect`, les paramètres fonction/map/opaques et les résultats `any` ou
-- | sans retour restent sur l'enveloppe (sémantique inchangée).
ffiFunctionInfos :: String -> Array (Tuple Ident (Maybe ExprType)) -> Array FfiDecl -> Map String FunctionInfo
ffiFunctionInfos modNameStr foreigns decls = Map.fromFoldable (Array.mapMaybe mk foreigns)
  where
  mk (Tuple (Ident name) mbTast) = case mbTast of
    Just tast | isPureConcreteResult tast -> do
      decl <- matchFfiDecl modNameStr decls name
      guardArity decl
      let workerNeeded = hasFunctionParam decl.args || ffiNeedsValueWorker decl.ret
      args <- traverse (ffiArgType workerNeeded) decl.args
      ret <- case decl.ret of
        Just t | workerNeeded -> if isRecordResult tast then Nothing else Just TypeValue
        Just t -> typeNodeToGoType t
        Nothing -> if isRecordResult tast then Nothing else Just TypeValue
      Just (Tuple name
        { fullName: (if workerNeeded then decl.name <> "_nativeWorker" else decl.name)
        , fArgs: args
        , fRet: ret
        , arity: Array.length decl.args
        })
    _ -> Nothing

-- | Workers « résultat boxé » pour les FFI appariées dont le retour natif
-- | n'est pas directement une `Value` (any/interface/[]any/void). Générés
-- | dans `_ffi.go` ; le site d'appel les utilise via `ffiFunctionInfos`.
ffiValueWorkers :: String -> Array DataDecl -> Array (Tuple Ident (Maybe ExprType)) -> Array FfiDecl -> String
ffiValueWorkers modNameStr dataDecls foreigns decls = String.joinWith "\n" (Array.mapMaybe worker foreigns)
  where
  worker (Tuple (Ident name) mbTast) = case mbTast of
    Just tast | isPureConcreteResult tast -> do
      decl <- matchFfiDecl modNameStr decls name
      guardArity decl
      if hasFunctionParam decl.args || ffiNeedsValueWorker decl.ret then Just (declCode decl tast) else Nothing
    _ -> Nothing

  callCode d = d.name <> "(" <> String.joinWith ", " (Array.mapWithIndex (\i _ -> "go_arg" <> show i) d.args) <> ")"

  declCode d tast =
    "func " <> d.name <> "_nativeWorker(" <> paramsCode d <> ") gopurs_runtime.Value {\n"
      <> String.joinWith "\n" (Array.concat (argLines d tast))
      <> "\n" <> callLine d
      <> "\n" <> String.joinWith "\n" (retLines d tast d.ret)
      <> "\n}\n"

  paramsCode d = String.joinWith ", " (Array.mapWithIndex (\i t -> "arg" <> show i <> " " <> case t of
    TFunc _ _ -> "gopurs_runtime.Value"
    _ -> printTypeNode t) d.args)

  argLines d tast = Array.mapWithIndex (\i t -> case t of
    TFunc _ _ ->
      let
        typStr = printTypeNode t
        tastArg = getTastArgType tast i
        -- Newtypes opaques (Pattern, Replacement...) : la valeur PS porte une
        -- fonction native, pas une fonction PS boxée (cf. `processArg`).
        isOpaque = case tastArg of
          Just (ADT _ _ _) -> true
          Just Any -> true
          Just (TypeVar _) -> true
          _ -> false
      in
        if isOpaque && not (isStandardPursFunc t) then
          [ "\tgo_arg" <> show i <> " := (*(*any)(arg" <> show i <> ".UnsafePtr)).(" <> typStr <> ")" ]
        else
          let
            adapter = unwrapValueToFunc dataDecls t tastArg ("arg" <> show i) 0 0
            indented = String.replaceAll (Pattern "\n") (Replacement "\n\t") adapter
          in
            [ "\tgo_arg" <> show i <> " := " <> indented ]
    _ -> [ "\tgo_arg" <> show i <> " := arg" <> show i ]) d.args

  callLine d = case d.ret of
    Nothing -> "\t" <> callCode d
    Just _ -> "\tgo_res := " <> callCode d

  retLines d tast ret = case ret of
    Nothing -> [ "\treturn gopurs_runtime.Value{}" ]
    Just t -> case t of
      -- Retour fonction : même conversion que l'enveloppe (`wrapReturn`).
      TFunc _ _ ->
        let
          wrapped = wrapReturn dataDecls t (getTastReturnType tast) "go_res"
          indented = String.replaceAll (Pattern "\n") (Replacement "\n\t") wrapped
        in
          [ "\treturn " <> indented ]
      TArray elem | not (isRuntimeValueNode elem) ->
        [ "\tout := make([]gopurs_runtime.Value, len(go_res))"
        , "\tfor i, v := range go_res { out[i] = gopurs_runtime.Box(v) }"
        , "\treturn gopurs_runtime.Array(out)"
        ]
      _ -> [ "\treturn gopurs_runtime.Box(go_res)" ]

ffiArgType :: Boolean -> TypeNode -> Maybe GoType
ffiArgType workerNeeded t = case t of
  TFunc _ _ | workerNeeded -> Just TypeValue
  _ -> typeNodeToGoType t

hasFunctionParam :: Array TypeNode -> Boolean
hasFunctionParam = Array.any case _ of
  TFunc _ _ -> true
  _ -> false

isRuntimeValueNode :: TypeNode -> Boolean
isRuntimeValueNode = case _ of
  TNamed "gopurs_runtime.Value" -> true
  _ -> false

guardArity :: FfiDecl -> Maybe Unit
guardArity d = if not d.isVar && Array.length d.args >= 1 then Just unit else Nothing

ffiNeedsValueWorker :: Maybe TypeNode -> Boolean
ffiNeedsValueWorker = case _ of
  Nothing -> true
  Just t -> needsValueWorker t

ffiReturnType :: Maybe TypeNode -> ExprType -> Maybe GoType
ffiReturnType retTast tast = case retTast of
  Just t | needsValueWorker t -> if isRecordResult tast then Nothing else Just TypeValue
  Just t -> typeNodeToGoType t
  Nothing -> if isRecordResult tast then Nothing else Just TypeValue

matchFfiDecl :: String -> Array FfiDecl -> String -> Maybe FfiDecl
matchFfiDecl modNameStr decls pursName =
  let
    fallback1 = modNameStr <> "_" <> capitalize pursName
    fallback2 = fallback1 <> "_"
    findDecl n = Array.find (\d -> d.name == n) decls
  in
    case findDecl fallback1 of
      Just d -> Just d
      Nothing -> findDecl fallback2

isRecordResult :: ExprType -> Boolean
isRecordResult t = case resultType t of
  Record _ -> true
  Row _ _ -> true
  _ -> false

resultType :: ExprType -> ExprType
resultType = case _ of
  Func _ ret -> resultType ret
  t -> t

typeNodeToGoType :: TypeNode -> Maybe GoType
typeNodeToGoType = case _ of
  TNamed "gopurs_runtime.Value" -> Just TypeValue
  TNamed "string" -> Just TypeString
  TNamed "int64" -> Just TypeInt64
  TNamed "float64" -> Just TypeFloat64
  TNamed "bool" -> Just TypeBool
  TNamed "uint32" -> Just TypeUint32
  TNamed "any" -> Just (TypeInterface "any")
  TNamed "interface{}" -> Just (TypeInterface "interface{}")
  TArray elem -> TypeNativeArray <$> typeNodeToGoType elem
  _ -> Nothing

needsValueWorker :: TypeNode -> Boolean
needsValueWorker = case _ of
  TNamed "any" -> true
  TNamed "interface{}" -> true
  TArray elem -> needsValueWorker elem
  _ -> false

-- | Résultat PureScript « simple » : ni fonction, ni `Effect`, ni
-- | quantifié/contraint, ni variable de type.
isPureConcreteResult :: ExprType -> Boolean
isPureConcreteResult = case _ of
  -- Le type PS d'une FFI fonctionnelle est un `Func` ; on teste son résultat.
  Func _ ret -> isPureConcreteResult ret
  ForAll _ _ -> false
  ConstrainedType _ _ -> false
  TypeVar _ -> false
  Any -> false
  other -> not (mentionsEffect other)

mentionsEffect :: ExprType -> Boolean
mentionsEffect = case _ of
  ADT fullName _ args -> String.contains (Pattern "Effect") fullName || Array.any mentionsEffect args
  TypeApp c args -> mentionsEffect c || Array.any mentionsEffect args
  Array t -> mentionsEffect t
  Func args ret -> Array.any mentionsEffect args || mentionsEffect ret
  Record row -> mentionsEffect row
  Row fields tail ->
    Array.any (\(Tuple _ v) -> mentionsEffect v) fields
      || case tail of
        Just t -> mentionsEffect t
        Nothing -> false
  ForAll _ t -> mentionsEffect t
  ConstrainedType constraints t ->
    Array.any (\(Tuple _ as) -> Array.any mentionsEffect as) constraints || mentionsEffect t
  _ -> false
