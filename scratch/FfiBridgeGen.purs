module Gopurs.FfiBridgeGen where

import Prelude
import Data.Array as Array
import Data.Maybe (Maybe(..), fromMaybe, isJust)
import Data.String as String
import Data.String.Pattern (Pattern(..))
import Data.Tuple (Tuple(..))
import Gopurs.FfiTypes (FfiDecl, TypeNode(..))
import PureScript.Backend.Optimizer.CoreFn (Ident(..))
import PureScript.Backend.Optimizer.FreeVars (sanitizeName)
import Data.Newtype (unwrap)
import Data.Map as Map

capitalize :: String -> String
capitalize s =
  case String.uncons s of
    Nothing -> ""
    Just { head, tail } -> String.singleton (String.toUpper head) <> tail

printTypeNode :: TypeNode -> String
printTypeNode (Named n) = n
printTypeNode (Func args ret) = 
  let 
    argsStr = String.joinWith ", " (map printTypeNode args)
    retStr = case ret of
      Nothing -> ""
      Just r -> " " <> printTypeNode r
  in "func(" <> argsStr <> ")" <> retStr
printTypeNode (Array elem) = "[]" <> printTypeNode elem
printTypeNode (Map k v) = "map[" <> printTypeNode k <> "]" <> printTypeNode v
printTypeNode (Unknown s) = s

unwrapValueToFunc :: TypeNode -> String -> Int -> Int -> String
unwrapValueToFunc (Func args ret) valName depth _ = 
  let
    paramsArr = Array.mapWithIndex (\cidx atype -> "p" <> show depth <> "_" <> show cidx <> " " <> printTypeNode atype) args
    applyArgsArr = Array.mapWithIndex (\cidx atype ->
      case atype of
        Named "gopurs_runtime.Value" -> "p" <> show depth <> "_" <> show cidx
        Func _ _ -> 
          let wrapped = wrapReturn atype ("p" <> show depth <> "_" <> show cidx)
          in String.replaceAll (Pattern "\n") (Pattern "\n\t\t") wrapped
        _ -> "gopurs_runtime.Box(p" <> show depth <> "_" <> show cidx <> ")"
    ) args
    params = String.joinWith ", " paramsArr
    applyArgs = String.joinWith ", " applyArgsArr
    
    applyCall =
      if Array.length args == 1 then
        "gopurs_runtime.Apply(" <> valName <> ", " <> applyArgs <> ")"
      else if Array.length args > 1 then
        "gopurs_runtime.Apply" <> show (Array.length args) <> "(" <> valName <> ", " <> applyArgs <> ")"
      else
        "gopurs_runtime.Apply(" <> valName <> ", gopurs_runtime.Value{})"
        
  in case ret of
    Nothing ->
      "func(" <> params <> ") {\n\t\t" <> applyCall <> "\n\t}"
    Just (Array elem) | printTypeNode elem /= "gopurs_runtime.Value" ->
      let elemType = printTypeNode elem
          retStr = printTypeNode (Array elem)
      in "func(" <> params <> ") " <> retStr <> " {\n\t\tinner_res" <> show depth <> " := " <> applyCall <> "\n\t\tres_arr" <> show depth <> " := *(*[]gopurs_runtime.Value)(inner_res" <> show depth <> ".UnsafePtr)\n\t\tres_go" <> show depth <> " := make(" <> retStr <> ", len(res_arr" <> show depth <> "))\n\t\tfor i, v := range res_arr" <> show depth <> " { res_go" <> show depth <> "[i] = gopurs_runtime.Unbox[" <> elemType <> "](v) }\n\t\treturn res_go" <> show depth <> "\n\t}"
    Just (Named "any") -> "func(" <> params <> ") any {\n\t\treturn " <> applyCall <> "\n\t}"
    Just (Named "interface{}") -> "func(" <> params <> ") interface{} {\n\t\treturn " <> applyCall <> "\n\t}"
    Just (Named "gopurs_runtime.Value") -> "func(" <> params <> ") gopurs_runtime.Value {\n\t\treturn " <> applyCall <> "\n\t}"
    Just f@(Func _ _) ->
      let innerUnwrap = unwrapValueToFunc f ("inner_res" <> show depth) (depth + 1) 0
      in "func(" <> params <> ") " <> printTypeNode f <> " {\n\t\tinner_res" <> show depth <> " := " <> applyCall <> "\n\t\treturn " <> innerUnwrap <> "\n\t}"
    Just r ->
      "func(" <> params <> ") " <> printTypeNode r <> " {\n\t\tinner_res" <> show depth <> " := " <> applyCall <> "\n\t\treturn gopurs_runtime.Unbox[" <> printTypeNode r <> "](inner_res" <> show depth <> ")\n\t}"
unwrapValueToFunc t valName _ _ = "gopurs_runtime.Unbox[" <> printTypeNode t <> "](" <> valName <> ")"

wrapReturn :: TypeNode -> String -> String
wrapReturn (Func args ret) valName = 
  let
    innerT = ret
    argT = Array.head args
    
    genInner val innerWrap = 
      "gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {\n\t\t\t" <> val <> "\n\t\t\treturn " <> innerWrap <> "\n\t\t})"
      
    genInnerArg val innerWrap argUnwrap =
      "gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {\n\t\t\t" <> val <> "(" <> argUnwrap <> ")\n\t\t\treturn " <> innerWrap <> "\n\t\t})"
  in case argT of
    Nothing ->
      case innerT of
        Nothing -> genInner (valName <> "()") "gopurs_runtime.Value{}"
        Just r -> genInner ("inner_res := " <> valName <> "()") (wrapReturn r "inner_res")
    Just a ->
      let 
        argUnwrap = case a of
          Named "any" -> "arg"
          Named "interface{}" -> "arg"
          Named "gopurs_runtime.Value" -> "arg"
          f@(Func _ _) -> String.replaceAll (Pattern "\n") (Pattern "\n\t\t\t") (unwrapValueToFunc f "arg" 99 0)
          _ -> "gopurs_runtime.Unbox[" <> printTypeNode a <> "](arg)"
      in case innerT of
        Nothing -> genInnerArg valName "gopurs_runtime.Value{}" argUnwrap
        Just r -> genInnerArg ("inner_res := " <> valName) (wrapReturn r "inner_res") argUnwrap
wrapReturn (Array elem) valName | printTypeNode elem /= "gopurs_runtime.Value" = 
  "func() gopurs_runtime.Value {\n\t\t\tres_arr := make([]gopurs_runtime.Value, len(" <> valName <> "))\n\t\t\tfor i, v := range " <> valName <> " { res_arr[i] = gopurs_runtime.Box(v) }\n\t\t\treturn gopurs_runtime.Array(res_arr)\n\t\t}()"
wrapReturn (Named "gopurs_runtime.Value") valName = valName
wrapReturn (Map _ _) valName = 
  "func() gopurs_runtime.Value {\n\t\t\tres_map := make(map[string]gopurs_runtime.Value)\n\t\t\tfor k, v := range " <> valName <> " { res_map[k] = gopurs_runtime.Box(v) }\n\t\t\treturn gopurs_runtime.Record(res_map)\n\t\t}()"
wrapReturn t valName = "gopurs_runtime.Box(" <> valName <> ")"

generateWrapperFunc :: FfiDecl -> String
generateWrapperFunc d = 
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
          
      substituteGeneric :: TypeNode -> TypeNode
      substituteGeneric t = 
        let s = printTypeNode t 
            res = Array.foldl (\acc tp -> String.replaceAll (Pattern ("\\b" <> tp <> "\\b")) (Replacement "gopurs_runtime.Value") acc) s d.typeParams
        in Unknown res -- Very simple hack to substitute generic params, works for most cases
        -- We actually just string replace it like the old code did if it's not perfectly matching
        -- In PureScript, if we just use printTypeNode it's fine, we can replace the string output instead.

      newLines = []
      
      callFunc = if Array.length d.typeParams > 0 then
          d.name <> "[" <> String.joinWith ", " (map (const "gopurs_runtime.Value") d.typeParams) <> "]"
        else d.name

      processArg i t =
        let typStr = printTypeNode t 
            elemType = String.drop 2 typStr
        in case t of
          Func _ _ -> 
            let unwrapStr = unwrapValueToFunc t ("arg" <> show i) 0 0
                indented = String.replaceAll (Pattern "\n") (Pattern "\n\t") unwrapStr
            in [ "\tgo_arg" <> show i <> " := " <> indented ]
          Array _ | typStr /= "[]gopurs_runtime.Value" ->
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
          Named "any" -> [ "\tgo_arg" <> show i <> " := arg" <> show i ]
          Named "interface{}" -> [ "\tgo_arg" <> show i <> " := arg" <> show i ]
          Named "gopurs_runtime.Value" -> [ "\tgo_arg" <> show i <> " := arg" <> show i ]
          Map _ _ ->
            let 
              et = String.drop (String.indexOf (Pattern "]") typStr # fromMaybe 0 # add 1) typStr
            in if et == "any" || et == "interface{}" then
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
          let wrapCode = wrapReturn r "go_res"
          in [ "\tgo_res := " <> callFunc <> "(" <> callArgs <> ")"
             , "\treturn " <> wrapCode
             ]
             
      fullCode = "gopurs_runtime." <> funcConstructor <> "(func(" <> boxedArgs <> ") gopurs_runtime.Value {\n" <>
                 String.joinWith "\n" argsCode <> "\n" <>
                 String.joinWith "\n" retCode <> "\n" <>
                 "})"
    in fullCode

generateFfiBridge :: Array FfiDecl -> Array (Tuple Ident a) -> String
generateFfiBridge decls foreigns = 
  String.joinWith "\n" (map genBridge foreigns)
  where
  genBridge (Tuple ident _) = 
    let 
      pursName = unwrap ident
      sanitized = sanitizeName pursName
      capName = capitalize sanitized
      exportName = "_Gopurs_" <> capName
      
      fallback1 = capitalize pursName
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
          "var " <> exportName <> " = " <> generateWrapperFunc d
