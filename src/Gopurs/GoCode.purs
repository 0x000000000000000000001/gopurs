module Gopurs.GoCode
  ( GoCode
  , BoundGoCode
  , SourceReference
  , SourceDeclaration
  , renderBoundCode
  , boundImports
  , opaqueCode
  , referencedImports
  ) where

import Prelude
import Data.Array as Array
import Data.Char as Char
import Data.Argonaut.Core (Json)
import Data.Either (Either(..))
import Data.Foldable (foldM)
import Data.List (List(..), (:))
import Data.List as List
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Set as Set
import Data.String as String
import Data.String.CodeUnits as CodeUnits
import Data.Traversable (traverse)
import Gopurs.GoSymbol (GoReference(..), GoSymbol, RenderContext, renderReference)
import Partial.Unsafe (unsafePartial)

type GoCode = { text :: String, imports :: Array String }

type SourceReference =
  { start :: Int, end :: Int, role :: String, name :: String, symbol :: GoSymbol }
type SourceDeclaration =
  { start :: Int, end :: Int, kind :: String, symbols :: Array GoSymbol }

-- Complete source units are bound after opaque fragments acquire their lexical
-- context. Positions address the exact text, not a gofmt reconstruction. Field
-- and method edges are retained separately from qualifiable package references.
type BoundGoCode =
  { owner :: String
  , path :: String
  , text :: String
  , imports :: Array { path :: String, alias :: String }
  , identifiers :: Array String
  , references :: Array SourceReference
  , declarations :: Array SourceDeclaration
  , uses :: Json
  }

boundImports :: RenderContext -> BoundGoCode -> Either String (Array String)
boundImports context code = do
  paths <- traverse dependency code.references
  pure (Set.toUnfoldable (Set.fromFoldable (map _.path code.imports <> Array.catMaybes paths)))
  where
  dependency reference | reference.role == "reference" = case Map.lookup reference.symbol context.symbols of
    Nothing -> Left ("Unplaced Go symbol: " <> reference.symbol.owner <> ":" <> reference.symbol.name)
    Just target | target.packagePath /= context.consumer -> Right (Just target.packagePath)
    _ -> Right Nothing
  dependency _ = Right Nothing

renderBoundCode :: RenderContext -> BoundGoCode -> Either String String
renderBoundCode context code = do
  result <- foldM render { end: 0, checked: 0, chunks: Nil, changed: false } code.references
  pure if result.changed then String.joinWith "" (Array.fromFoldable (List.reverse result.chunks) <> [ CodeUnits.drop result.end code.text ]) else code.text
  where
  -- Native strings are UTF-8. Convert once so checking ordered UTF-16 spans
  -- never rescans the full prefix for every identifier in a large source file.
  chars = CodeUnits.toCharArray code.text
  slice start end = CodeUnits.fromCharArray (Array.slice start end chars)

  render state reference = do
    if reference.start < state.checked || reference.end <= reference.start || reference.end > Array.length chars
      || slice reference.start reference.end /= reference.name then
      Left ("Invalid Go source reference in " <> code.path <> ": " <> reference.name)
    else do
      name <- case reference.role of
        "reference" -> do
          rendered <- renderReference context (SymbolReference reference.symbol)
          case Map.lookup reference.symbol context.symbols of
            Just target | target.packagePath /= context.consumer -> case Map.lookup target.packagePath context.aliases of
              Just alias | Array.elem alias code.identifiers ->
                  Left ("Go import alias shadows an identifier in " <> code.path <> ": " <> alias)
              _ -> Right rendered
            _ -> Right rendered
        "definition" -> case Map.lookup reference.symbol context.symbols of
          Just target | target.packagePath == context.consumer -> Right target.name
          _ -> Left ("Go definition is not owned by the consumer: " <> reference.symbol.name)
        "field" -> Right reference.name
        "method" -> Right reference.name
        "import" -> Right reference.name
        "external" -> Right reference.name
        _ -> Left ("Unknown Go reference role: " <> reference.role)
      pure if name == reference.name then state { checked = reference.end }
        else { end: reference.end
             , checked: reference.end
             , chunks: name : slice state.end reference.start : state.chunks
             , changed: true }

-- Legacy opaque fragments retain their dependencies when they enter the AST.
-- This fallback inspects only the fragment, never a rendered module.
opaqueCode :: String -> GoCode
opaqueCode text = { text, imports: referencedImports text }

-- A separate typed worker keeps this frequent lookup unboxed in native code.
charAtOrSpace :: Array Char -> Int -> Char
charAtOrSpace chars index =
  if index >= 0 && index < Array.length chars then
    unsafePartial (Array.unsafeIndex chars index)
  else ' '

referencedImports :: String -> Array String
referencedImports text = referencedImportsImpl referencedImportsPS text

-- The Go backend scans in one byte pass without decoding the text to code
-- points or building intermediate import arrays. The JavaScript backend calls
-- the PureScript implementation passed as the first argument, so the JS bundle
-- keeps the exact previous behaviour.
foreign import referencedImportsImpl :: (String -> Array String) -> String -> Array String

referencedImportsPS :: String -> Array String
referencedImportsPS text = Array.sort (scan 0 [])
  where
  -- Native strings are UTF-8: repeatedly looking up a UTF-16 position would
  -- rescan their prefix. Decode once, then retain constant-time indexing.
  chars = CodeUnits.toCharArray text
  size = Array.length chars
  charAt index = charAtOrSpace chars index

  isIdentifierChar char =
    -- Compare one numeric code instead of repeatedly boxing Char bounds.
    -- Explicit branches also avoid closures for partially applied booleans.
    let code = Char.toCharCode char
    in if code > 127 then true
       else if code >= 97 then code <= 122
       else if code >= 65 then if code <= 90 then true else code == 95
       else if code >= 48 then code <= 57
       else false

  skipIdentifier index
    | index < size && isIdentifierChar (charAt index) = skipIdentifier (index + 1)
    | otherwise = index

  skipQuoted quote index
    | index >= size = size
    | charAt index == quote = index + 1
    | quote /= '`' && charAt index == '\\' = skipQuoted quote (index + 2)
    | otherwise = skipQuoted quote (index + 1)

  skipLine index
    | index >= size || charAt index == '\n' = index
    | otherwise = skipLine (index + 1)

  skipComment index
    | index >= size = size
    | charAt index == '*' && charAt (index + 1) == '/' = index + 2
    | otherwise = skipComment (index + 1)

  scan index imports
    | index >= size = imports
    | otherwise =
        let
          next = case charAt index of
            '"' -> { index: skipQuoted '"' (index + 1), imports }
            '\'' -> { index: skipQuoted '\'' (index + 1), imports }
            '`' -> { index: skipQuoted '`' (index + 1), imports }
            '/' | charAt (index + 1) == '/' -> { index: skipLine (index + 2), imports }
            '/' | charAt (index + 1) == '*' -> { index: skipComment (index + 2), imports }
            char | isIdentifierChar char ->
              let
                end = skipIdentifier (index + 1)
                dependency = if charAt end == '.' then
                  case CodeUnits.fromCharArray (map charAt (Array.range index (end - 1))) of
                    "gopurs_runtime" -> Just "gopurs/output/gopurs_runtime"
                    "math" -> Just "math"
                    "sync" -> Just "sync"
                    "unsafe" -> Just "unsafe"
                    _ -> Nothing
                  else Nothing
                nextImports = case dependency of
                  Just path | not (Array.elem path imports) -> Array.snoc imports path
                  _ -> imports
              in { index: end, imports: nextImports }
            _ -> { index: index + 1, imports }
        -- Keep the only recursive call outside guarded branches: their local
        -- helper functions must return a step, never recurse back into scan.
        in scan next.index next.imports
