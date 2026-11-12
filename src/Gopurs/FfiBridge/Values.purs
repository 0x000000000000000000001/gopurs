module Gopurs.FfiBridge.Values
  ( boxedArgument
  , nativeArgument
  , toValue
  , workerReturn
  ) where

import Prelude

import Data.Array as Array
import Data.Maybe (Maybe(..), fromMaybe)
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Tuple (Tuple(..))
import Gopurs.FfiBridge.Structured as Structured
import Gopurs.FfiBridge.TypeSupport (flattenFuncArgs, getTastArgType, getTastReturnType, isRuntimeValueNode, isStandardPursFunc, isUncurriedFunction, printTypeNode, resolveNewtype)
import Gopurs.FfiTypes (TypeNode(..))
import Gopurs.GoTypes (isClosedRowTail)
import PureScript.Backend.Optimizer.CoreFn (DataDecl, ExprType(..))

boxedArgument :: Array DataDecl -> Maybe ExprType -> Int -> TypeNode -> Array String
boxedArgument dataDecls mbTast i t
  | Just tast <- mbTast >>= \ft -> getTastArgType ft i
  , Structured.needsConversion t tast =
      [ "\tgo_arg" <> show i <> " := " <> Structured.fromValue (\go ps value -> fromValue dataDecls go (Just ps) value 0) t tast ("arg" <> show i) ]
boxedArgument dataDecls mbTast i t =
  let
    typStr = printTypeNode t
    tastArg = mbTast >>= \tast -> getTastArgType tast i
  in
    case t of
      TNamed "any" | isForeign tastArg -> [ "\tgo_arg" <> show i <> " := gopurs_runtime.ForeignValue(arg" <> show i <> ")" ]
      TNamed "interface{}" | isForeign tastArg -> [ "\tgo_arg" <> show i <> " := gopurs_runtime.ForeignValue(arg" <> show i <> ")" ]
      TArray (TNamed "any") | isForeignArray tastArg -> [ "\tgo_arg" <> show i <> " := gopurs_runtime.ForeignValue(arg" <> show i <> ").([]any)" ]
      TArray (TNamed "interface{}") | isForeignArray tastArg -> [ "\tgo_arg" <> show i <> " := gopurs_runtime.ForeignValue(arg" <> show i <> ").([]any)" ]
      TArray _ | isOpaqueType tastArg -> [ "\tgo_arg" <> show i <> " := gopurs_runtime.Unbox[" <> typStr <> "](arg" <> show i <> ")" ]
      TFunc _ _ -> callbackArgument dataDecls tastArg i t
      TArray _ | typStr /= "[]gopurs_runtime.Value" ->
        [ "\tgo_arg" <> show i <> " := " <> fromValue dataDecls t tastArg ("arg" <> show i) 0 ]
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

-- Worker scalar/container arguments already have their native representation.
-- Callback parameters still arrive as Value and use the wrapper's adapter.
nativeArgument :: Array DataDecl -> ExprType -> Int -> TypeNode -> Array String
nativeArgument dataDecls tast i t = case t of
  TFunc _ _ -> callbackArgument dataDecls (getTastArgType tast i) i t
  _ -> [ "\tgo_arg" <> show i <> " := arg" <> show i ]

callbackArgument :: Array DataDecl -> Maybe ExprType -> Int -> TypeNode -> Array String
callbackArgument dataDecls tastArg i t =
  let
    -- Opaque newtypes can carry native functions rather than boxed PS callbacks.
    isOpaque = case tastArg of
      Just (ADT name _ _) -> not (isUncurriedFunction name || name == "Effect.Effect")
      Just Any -> true
      Just (TypeVar _) -> true
      _ -> false
  in
    if isOpaque && not (isStandardPursFunc t) then
      [ "\tgo_arg" <> show i <> " := (*(*any)(arg" <> show i <> ".UnsafePtr)).(" <> printTypeNode t <> ")" ]
    else
      let
        adapter = fromValue dataDecls t tastArg ("arg" <> show i) 0
        indented = String.replaceAll (Pattern "\n") (Replacement "\n\t") adapter
      in
        [ "\tgo_arg" <> show i <> " := " <> indented ]

-- Adapt values flowing from boxed PureScript callbacks into native Go types.
fromValue :: Array DataDecl -> TypeNode -> Maybe ExprType -> String -> Int -> String
fromValue dataDecls t (Just tast) valName depth
  | Structured.needsConversion t tast = Structured.fromValue (\go ps value -> fromValue dataDecls go (Just ps) value depth) t tast valName
-- Nested native arrays require element conversion at every level. A slice of
-- boxed Values cannot be reinterpreted as a slice of scalar Go elements.
fromValue dataDecls array@(TArray item) mbTast valName depth
  | not (isOpaqueType mbTast) && not (isRuntimeValueNode item) =
      let
        itemTast = case mbTast of
          Just (Array tast) -> Just tast
          _ -> Nothing
        element = "gopurs_runtime.ArrayAccess(_source, _i)"
        converted = case item of
          TNamed name | (name == "any" || name == "interface{}") && not (isForeign itemTast) -> element
          _ -> fromValue dataDecls item itemTast element (depth + 1)
      in
        "func() " <> printTypeNode array <> " { _source := " <> valName
          <> "; _out := make(" <> printTypeNode array <> ", gopurs_runtime.ArrayLength(_source)); for _i := range _out { _out[_i] = "
          <> converted <> " }; return _out }()"
fromValue dataDecls (TFunc args ret) mbTast valName depth =
  let
    paramsArr = Array.mapWithIndex (\cidx atype -> "p" <> show depth <> "_" <> show cidx <> " " <> printTypeNode atype) args
    applyArgsArr = Array.mapWithIndex
      ( \cidx atype ->
          case atype of
            TNamed "gopurs_runtime.Value" -> "p" <> show depth <> "_" <> show cidx
            TFunc _ _ ->
              let
                wrapped = toValue dataDecls atype (mbTast >>= \tast -> getTastArgType tast cidx) ("p" <> show depth <> "_" <> show cidx)
              in
                String.replaceAll (Pattern "\n") (Replacement "\n\t\t") wrapped
            _ -> toValue dataDecls atype (mbTast >>= \tast -> getTastArgType tast cidx) ("p" <> show depth <> "_" <> show cidx)
      )
      args
    params = String.joinWith ", " paramsArr
    applyArgs = String.joinWith ", " applyArgsArr
    rawApplyCall =
      if Array.length args == 1 then
        "gopurs_runtime.Apply(" <> valName <> ", " <> applyArgs <> ")"
      else if Array.length args > 1 then
        "gopurs_runtime.Apply" <> show (Array.length args) <> "(" <> valName <> ", " <> applyArgs <> ")"
      else
        "gopurs_runtime.Apply(" <> valName <> ", gopurs_runtime.Value{})"
    applyCall = case mbTast >>= getTastReturnType of
      Just (ADT "Effect.Effect" _ _) | nativeCallbackResult ret -> "gopurs_runtime.Apply(" <> rawApplyCall <> ", gopurs_runtime.Value{})"
      _ -> rawApplyCall
    resultTast = case mbTast >>= getTastReturnType of
      Just (ADT "Effect.Effect" _ [result]) | nativeCallbackResult ret -> Just result
      other -> other
    nativeReturn result =
      "func(" <> params <> ") " <> printTypeNode result <> " {\n\t\tinner_res" <> show depth <> " := " <> applyCall
        <> "\n\t\treturn " <> fromValue dataDecls result resultTast ("inner_res" <> show depth) (depth + 1) <> "\n\t}"

  in
    case ret of
      Nothing ->
        "func(" <> params <> ") {\n\t\t" <> applyCall <> "\n\t}"
      Just array@(TArray _) -> nativeReturn array
      Just dynamic@(TNamed "any") | isForeign resultTast -> nativeReturn dynamic
      Just dynamic@(TNamed "interface{}") | isForeign resultTast -> nativeReturn dynamic
      Just (TNamed "any") -> "func(" <> params <> ") any {\n\t\treturn " <> applyCall <> "\n\t}"
      Just (TNamed "interface{}") -> "func(" <> params <> ") interface{} {\n\t\treturn " <> applyCall <> "\n\t}"
      Just (TNamed "gopurs_runtime.Value") -> "func(" <> params <> ") gopurs_runtime.Value {\n\t\treturn " <> applyCall <> "\n\t}"
      Just f@(TFunc _ _) ->
        let
          innerUnwrap = fromValue dataDecls f (mbTast >>= getTastReturnType) ("inner_res" <> show depth) (depth + 1)
        in
          "func(" <> params <> ") " <> printTypeNode f <> " {\n\t\tinner_res" <> show depth <> " := " <> applyCall <> "\n\t\treturn " <> innerUnwrap <> "\n\t}"
      Just r -> nativeReturn r
fromValue _ (TNamed anyT) mbTast valName _ | (anyT == "any" || anyT == "interface{}") && isForeign mbTast = "gopurs_runtime.ForeignValue(" <> valName <> ")"
fromValue dataDecls (TNamed anyT) mbTast valName depth | anyT == "any" || anyT == "interface{}" || anyT == "gopurs_runtime.Value" =
  let
    resolvedTast = map (resolveNewtype dataDecls) mbTast
  in
    case resolvedTast of
      Just (Record (Row fields tail)) | isClosedRowTail tail ->
        let
          fieldStr = map
            ( \(Tuple fK fT) ->
                "\t\t\t\tres_map[\"" <> fK <> "\"] = " <> fromValue dataDecls (TNamed "any") (Just fT) ("_raw[\"" <> fK <> "\"]") (depth + 1)
            )
            fields
        in
          "func() map[string]any {\n\t\t\t_raw := gopurs_runtime.RecordToMap(" <> valName <> ")\n\t\t\tres_map := make(map[string]any)\n" <> String.joinWith "\n" fieldStr <> "\n\t\t\treturn res_map\n\t\t}()"
      Just f@(Func _ _) ->
        fromValue dataDecls (exprTypeToDummyTypeNode f) resolvedTast valName depth
      _ -> "gopurs_runtime.Unbox[" <> anyT <> "](" <> valName <> ")"
fromValue _ t _ valName _ = "gopurs_runtime.Unbox[" <> printTypeNode t <> "](" <> valName <> ")"

exprTypeToDummyTypeNode :: ExprType -> TypeNode
exprTypeToDummyTypeNode (Func args ret) = TFunc (map exprTypeToDummyTypeNode args) (Just (exprTypeToDummyTypeNode ret))
exprTypeToDummyTypeNode (Array elem) = TArray (exprTypeToDummyTypeNode elem)
exprTypeToDummyTypeNode (Record _) = TMap (TNamed "string") (TNamed "any")
exprTypeToDummyTypeNode _ = TNamed "any"

boxScalar :: TypeNode -> String -> String
boxScalar (TNamed "gopurs_runtime.Value") valName = valName
boxScalar (TNamed "int64") valName = "gopurs_runtime.Int(" <> valName <> ")"
boxScalar (TNamed "int") valName = "gopurs_runtime.Int(int64(" <> valName <> "))"
boxScalar _ valName = "gopurs_runtime.Box(" <> valName <> ")"

-- A Value already has the required representation, even for PS records and
-- functions. Exact Value -> Value callbacks need no forwarding closure.
toValue :: Array DataDecl -> TypeNode -> Maybe ExprType -> String -> String
toValue dataDecls t (Just tast) valName
  | Structured.needsConversion t tast = Structured.toValue (\go ps value -> toValue dataDecls go (Just ps) value) t tast valName
toValue _ (TNamed "gopurs_runtime.Value") _ valName = valName
toValue _ (TFunc [ TNamed "gopurs_runtime.Value" ] (Just (TNamed "gopurs_runtime.Value"))) _ valName =
  "gopurs_runtime.Func(" <> valName <> ")"
toValue dataDecls (TFunc args ret) mbTast valName =
  let
    genInner val innerWrap =
      "gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {\n\t\t\t" <> val <> "\n\t\t\treturn " <> innerWrap <> "\n\t\t})"
    genInnerArg val innerWrap argUnwrap =
      "gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {\n\t\t\t" <> val <> "(" <> argUnwrap <> ")\n\t\t\treturn " <> innerWrap <> "\n\t\t})"
  in
    case Array.head args of
      Nothing ->
        case ret of
          Nothing -> genInner (valName <> "()") "gopurs_runtime.Value{}"
          Just r -> genInner ("inner_res := " <> valName <> "()") (toValue dataDecls r (mbTast >>= getTastReturnType) "inner_res")
      Just a ->
        let
          argUnwrap = case a of
            TNamed "any" -> "arg"
            TNamed "interface{}" -> "arg"
            TNamed "gopurs_runtime.Value" -> "arg"
            f@(TFunc _ _) -> String.replaceAll (Pattern "\n") (Replacement "\n\t\t\t") (fromValue dataDecls f (mbTast >>= \tast -> getTastArgType tast 0) "arg" 99)
            _ -> "gopurs_runtime.Unbox[" <> printTypeNode a <> "](arg)"
        in
          case ret of
            Nothing -> genInnerArg valName "gopurs_runtime.Value{}" argUnwrap
            Just r -> genInnerArg ("inner_res := " <> valName) (toValue dataDecls r (mbTast >>= getTastReturnType) "inner_res") argUnwrap
toValue _ (TArray _) mbTast valName | isOpaqueType mbTast = "gopurs_runtime.Any(" <> valName <> ")"
toValue dataDecls (TArray elem) mbTast valName | printTypeNode elem /= "gopurs_runtime.Value" =
  let
    itemTast = case mbTast of
      Just (Array t) -> Just t
      _ -> Nothing
  in
  "func() gopurs_runtime.Value {\n\t\t\tres_arr := make([]gopurs_runtime.Value, len(" <> valName <> "))\n\t\t\tfor i, v := range " <> valName <> " { res_arr[i] = " <> toValue dataDecls elem itemTast "v" <> " }\n\t\t\treturn gopurs_runtime.Array(res_arr)\n\t\t}()"
toValue dataDecls (TMap _ _) (Just (Record (Row fields tail))) valName | isClosedRowTail tail =
  let
    fieldStr = map
      ( \(Tuple fK fT) ->
          "\t\t\t\tres_map[\"" <> fK <> "\"] = " <> toValue dataDecls (TNamed "any") (Just fT) ("_raw[\"" <> fK <> "\"]")
      )
      fields
  in
    "func() gopurs_runtime.Value {\n\t\t\t_raw := " <> valName <> "\n\t\t\tres_map := make(map[string]gopurs_runtime.Value)\n" <> String.joinWith "\n" fieldStr <> "\n\t\t\treturn gopurs_runtime.Record(res_map)\n\t\t}()"
toValue dataDecls (TMap _ _) mbTast valName =
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
toValue _ (TNamed anyT) mbTast valName | (anyT == "any" || anyT == "interface{}") && isForeign mbTast = "gopurs_runtime.Any(" <> valName <> ")"
toValue dataDecls (TNamed anyT) mbTast valName | anyT == "any" || anyT == "interface{}" =
  let
    resolvedTast = map (resolveNewtype dataDecls) mbTast
  in
    case resolvedTast of
      Just (Record (Row fields tail)) | isClosedRowTail tail ->
        let
          fieldStr = map
            ( \(Tuple fK fT) ->
                "\t\t\t\tres_map[\"" <> fK <> "\"] = " <> toValue dataDecls (TNamed "any") (Just fT) ("_raw[\"" <> fK <> "\"]")
            )
            fields
        in
          "func() gopurs_runtime.Value {\n\t\t\t_raw := gopurs_runtime.RecordToMap(gopurs_runtime.Box(" <> valName <> "))\n\t\t\tres_map := make(map[string]gopurs_runtime.Value)\n" <> String.joinWith "\n" fieldStr <> "\n\t\t\treturn gopurs_runtime.Record(res_map)\n\t\t}()"
      Just f@(Func _ _) ->
        let
          fArgs = flattenFuncArgs f
          arity = Array.length fArgs
          genWrap remaining depth =
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
                "gopurs_runtime.Func(func(" <> currArg <> " gopurs_runtime.Value) gopurs_runtime.Value {\n\t\t\treturn " <> genWrap (remaining - 1) (depth + 1) <> "\n\t\t})"
        in
          genWrap arity 0
      _ -> "gopurs_runtime.Box(" <> valName <> ")"
toValue _ typ _ valName = boxScalar typ valName

-- Workers box their native return directly, except for callbacks and arrays.
-- Their array path deliberately keeps the generic Box convention.
workerReturn :: Array DataDecl -> ExprType -> Maybe TypeNode -> Array String
workerReturn dataDecls tast = case _ of
  Nothing -> [ "\treturn gopurs_runtime.Value{}" ]
  Just t -> case t of
    TFunc _ _ ->
      let
        wrapped = toValue dataDecls t (getTastReturnType tast) "go_res"
        indented = String.replaceAll (Pattern "\n") (Replacement "\n\t") wrapped
      in
        [ "\treturn " <> indented ]
    TArray elem | not (isRuntimeValueNode elem) ->
      [ "\tout := make([]gopurs_runtime.Value, len(go_res))"
      , "\tfor i, v := range go_res { out[i] = gopurs_runtime.Box(v) }"
      , "\treturn gopurs_runtime.Array(out)"
      ]
    _ -> [ "\treturn gopurs_runtime.Box(go_res)" ]

isForeign :: Maybe ExprType -> Boolean
isForeign (Just (ADT "Foreign.Foreign" _ _)) = true
isForeign _ = false

isForeignArray :: Maybe ExprType -> Boolean
isForeignArray (Just (Array t)) = isForeign (Just t)
isForeignArray _ = false

isOpaqueType :: Maybe ExprType -> Boolean
isOpaqueType (Just (ADT _ _ _)) = true
isOpaqueType _ = false

-- Historical callbacks return the Effect thunk explicitly. Only ordinary Go
-- results request its execution inside the adapter.
nativeCallbackResult :: Maybe TypeNode -> Boolean
nativeCallbackResult (Just (TFunc _ _)) = false
nativeCallbackResult (Just (TNamed "any")) = false
nativeCallbackResult (Just (TNamed "interface{}")) = false
nativeCallbackResult (Just (TNamed "gopurs_runtime.Value")) = false
nativeCallbackResult _ = true
