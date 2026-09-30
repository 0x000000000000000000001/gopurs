module Gopurs.FfiBridge
  ( generateFfiBridge
  , ffiFunctionInfos
  , ffiValueWorkers
  ) where

import Prelude

import Data.Array as Array
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.String as String
import Data.Tuple (Tuple(..))
import Gopurs.CodegenState (FunctionInfo)
import Gopurs.FfiBridge.Render as Render
import Gopurs.FfiBridge.Signatures (NativeCandidate)
import Gopurs.FfiBridge.Signatures as Signatures
import Gopurs.FfiTypes (FfiDecl)
import Gopurs.GoAst (capitalize, sanitizeName)
import PureScript.Backend.Optimizer.CoreFn (DataDecl, ExprType, Ident(..))

-- Every foreign binding keeps a boxed entry, including first-class uses and
-- declarations that cannot publish a direct native-call signature.
generateFfiBridge :: String -> Array DataDecl -> Array FfiDecl -> Array (Tuple Ident (Maybe ExprType)) -> String
generateFfiBridge modNameStr dataDecls decls foreigns = String.joinWith "\n" (map bridge foreigns)
  where
  bridge (Tuple (Ident name) tast) =
    let exportName = "_Gopurs_" <> modNameStr <> "_" <> capitalize (sanitizeName name)
    in
      case Signatures.matchDeclaration modNameStr decls name of
        Nothing -> Render.missingImplementation exportName name
        Just declaration -> "var " <> exportName <> " = " <> Render.wrapper dataDecls declaration tast

-- Only publish signatures whose arguments and result can cross a direct call.
ffiFunctionInfos :: String -> Array (Tuple Ident (Maybe ExprType)) -> Array FfiDecl -> Map String FunctionInfo
ffiFunctionInfos modNameStr foreigns decls = Map.fromFoldable (Array.mapMaybe signature foreigns)
  where
  signature foreignBinding@(Tuple (Ident name) _) = do
    candidate <- nativeForeign modNameStr decls foreignBinding
    info <- Signatures.functionInfo candidate
    pure (Tuple name info)

-- Worker emission uses candidate eligibility, before the stricter signature
-- check above. Some generated workers therefore remain unused by direct calls.
ffiValueWorkers :: String -> Array DataDecl -> Array (Tuple Ident (Maybe ExprType)) -> Array FfiDecl -> String
ffiValueWorkers modNameStr dataDecls foreigns decls = String.joinWith "\n" (Array.mapMaybe worker foreigns)
  where
  worker foreignBinding = do
    candidate <- nativeForeign modNameStr decls foreignBinding
    if candidate.requiresWorker then Just (Render.valueWorker dataDecls candidate) else Nothing

nativeForeign :: String -> Array FfiDecl -> Tuple Ident (Maybe ExprType) -> Maybe NativeCandidate
nativeForeign modNameStr decls (Tuple (Ident name) mbTast) = do
  tast <- mbTast
  declaration <- Signatures.matchDeclaration modNameStr decls name
  Signatures.nativeCandidate declaration tast
