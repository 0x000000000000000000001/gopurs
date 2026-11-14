module Gopurs.FfiBridge
  ( generateFfiBridge
  , missingImplementations
  , ffiFunctionInfos
  , ffiProblems
  , ffiProblemsWithTypes
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
import Gopurs.FfiBridge.Effects as Effects
import Gopurs.FfiBridge.Render as Render
import Gopurs.FfiBridge.Signatures (NativeCandidate)
import Gopurs.FfiBridge.Signatures as Signatures
import Gopurs.FfiTypes (FfiDecl)
import Gopurs.GoTypes (printExprType)
import Gopurs.FfiBridge.TypeSupport (printTypeNode, resolveNewtype)
import Gopurs.GoAst (capitalize, sanitizeName)
import PureScript.Backend.Optimizer.CoreFn (DataDecl, ExprType(..), Ident(..))

-- Every foreign binding keeps a boxed entry, including first-class uses and
-- declarations that cannot publish a direct native-call signature.
generateFfiBridge :: String -> Array DataDecl -> Array FfiDecl -> Array (Tuple Ident (Maybe ExprType)) -> String
generateFfiBridge modNameStr dataDecls decls foreigns = String.joinWith "\n" (map bridge foreigns)
  where
  bridge (Tuple (Ident name) tast) =
    let exportName = foreignExportName modNameStr name
    in
      case Signatures.matchDeclaration modNameStr decls name of
        Nothing | unitConstant modNameStr name tast -> "var " <> exportName <> " = gopurs_runtime.Value{}"
        Nothing -> Render.missingImplementation exportName name
        Just declaration -> "var " <> exportName <> " = " <> Render.wrapper dataDecls declaration (map (resolveNewtype dataDecls) tast)

-- The driver records these alongside the generated stubs. Reachability is
-- checked after translation, when intrinsic calls have already disappeared.
missingImplementations :: String -> Array FfiDecl -> Array (Tuple Ident (Maybe ExprType)) -> Array { binding :: String, symbol :: String }
missingImplementations modNameStr decls = Array.mapMaybe \(Tuple (Ident name) tast) ->
  case Signatures.matchDeclaration modNameStr decls name of
    Nothing | unitConstant modNameStr name tast -> Nothing
    Nothing -> Just { binding: foreignExportName modNameStr name, symbol: name }
    Just _ -> Nothing

foreignExportName :: String -> String -> String
foreignExportName modNameStr name = "_Gopurs_" <> modNameStr <> "_" <> capitalize (sanitizeName name)

-- The canonical Unit value has a single representation. Materialize that
-- value when using stock prelude, while keeping explicit Go FFI authoritative.
unitConstant :: String -> String -> Maybe ExprType -> Boolean
unitConstant "Data_Unit" "unit" (Just Unit) = true
unitConstant _ _ _ = false

-- Only publish signatures whose arguments and result can cross a direct call.
ffiFunctionInfos :: String -> Array (Tuple Ident (Maybe ExprType)) -> Array FfiDecl -> Map String FunctionInfo
ffiFunctionInfos modNameStr foreigns decls = Map.fromFoldable (Array.mapMaybe signature foreigns)
  where
  signature foreignBinding@(Tuple (Ident name) _) = do
    candidate <- nativeForeign modNameStr decls foreignBinding
    info <- Signatures.functionInfo modNameStr candidate
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

-- Reject incompatible contracts before publishing any Go output. The driver
-- supplies the original module and handwritten source path around these facts.
ffiProblems :: String -> Array FfiDecl -> Array (Tuple Ident (Maybe ExprType)) -> Array String
ffiProblems modName decls = Array.mapMaybe \(Tuple (Ident name) tast) -> do
  d <- Signatures.matchDeclaration modName decls name
  t <- tast
  reason <- Effects.problem d t
  pure (name <> " at " <> show d.line <> ":" <> show d.column
    <> ": Go (" <> String.joinWith ", " (map printTypeNode d.args)
    <> ") -> (" <> String.joinWith ", " (map printTypeNode d.results)
    <> "); PureScript " <> printExprType t <> ". " <> reason)

ffiProblemsWithTypes :: String -> Array DataDecl -> Array FfiDecl -> Array (Tuple Ident (Maybe ExprType)) -> Array String
ffiProblemsWithTypes modName dataDecls decls = ffiProblems modName decls <<< map (map (map (resolveNewtype dataDecls)))
