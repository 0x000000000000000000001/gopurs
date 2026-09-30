module Gopurs.FfiBridge.Render
  ( wrapper
  , valueWorker
  , missingImplementation
  ) where

import Prelude

import Data.Array as Array
import Data.Maybe (Maybe(..))
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Gopurs.FfiBridge.Signatures (NativeCandidate)
import Gopurs.FfiBridge.TypeSupport (getTastReturnType, instantiateDeclaration, printTypeNode)
import Gopurs.FfiBridge.Values as Values
import Gopurs.FfiTypes (FfiDecl, TypeNode(..))
import Gopurs.GoTypes (printExprType)
import PureScript.Backend.Optimizer.CoreFn (DataDecl, ExprType)

wrapper :: Array DataDecl -> FfiDecl -> Maybe ExprType -> String
wrapper dataDecls declaration mbTast =
  let
    d = instantiateDeclaration declaration
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
            if arity == 0 then "_ gopurs_runtime.Value"
            else String.joinWith ", " (Array.mapWithIndex (\i _ -> "arg" <> show i <> " gopurs_runtime.Value") d.args)
          callFunc =
            if Array.null d.typeParams then d.name
            else d.name <> "[" <> String.joinWith ", " (map (const "gopurs_runtime.Value") d.typeParams) <> "]"
          argsCode = Array.concat (Array.mapWithIndex (Values.boxedArgument dataDecls mbTast) d.args)
          call = callFunc <> "(" <> callArguments d <> ")"
          retCode = case d.ret of
            Nothing ->
              [ "\t" <> call
              , "\treturn gopurs_runtime.Value{}"
              ]
            Just r ->
              let
                wrapped = Values.toValue dataDecls r (mbTast >>= getTastReturnType) "go_res"
                indented = String.replaceAll (Pattern "\n") (Replacement "\n\t") wrapped
              in
                [ "\tgo_res := " <> call
                , "\treturn " <> indented
                ]
        in
          "gopurs_runtime." <> funcConstructor <> "(func(" <> boxedArgs <> ") gopurs_runtime.Value {\n"
            <> String.joinWith "\n" argsCode <> "\n"
            <> String.joinWith "\n" retCode <> "\n})"

valueWorker :: Array DataDecl -> NativeCandidate -> String
valueWorker dataDecls { declaration, tast, callName } =
  let
    params = String.joinWith ", " (Array.mapWithIndex
      (\i t -> "arg" <> show i <> " " <> case t of
        TFunc _ _ -> "gopurs_runtime.Value"
        _ -> printTypeNode t)
      declaration.args)
    args = Array.concat (Array.mapWithIndex (Values.nativeArgument dataDecls tast) declaration.args)
    call = declaration.name <> "(" <> callArguments declaration <> ")"
    callLine = case declaration.ret of
      Nothing -> "\t" <> call
      Just _ -> "\tgo_res := " <> call
  in
    "func " <> callName <> "(" <> params <> ") gopurs_runtime.Value {\n"
      <> String.joinWith "\n" args
      <> "\n" <> callLine
      <> "\n" <> String.joinWith "\n" (Values.workerReturn dataDecls tast declaration.ret)
      <> "\n}\n"

callArguments :: FfiDecl -> String
callArguments declaration = String.joinWith ", " (Array.mapWithIndex (\i _ -> "go_arg" <> show i) declaration.args)

missingImplementation :: String -> String -> String
missingImplementation exportName pursName =
  "var " <> exportName <> " = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value { panic(\"FFI not implemented: " <> pursName <> "\"); return gopurs_runtime.Value{} })"
