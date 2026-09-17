module Gopurs.GoCode
  ( GoCode
  , opaqueCode
  , referencedImports
  ) where

import Prelude
import Data.Array as Array
import Data.Maybe (Maybe(..), fromMaybe)
import Data.String.CodeUnits as CodeUnits

type GoCode = { text :: String, imports :: Array String }

-- Legacy opaque fragments retain their dependencies when they enter the AST.
-- This fallback inspects only the fragment, never a rendered module.
opaqueCode :: String -> GoCode
opaqueCode text = { text, imports: referencedImports text }

referencedImports :: String -> Array String
referencedImports text = Array.sort (scan 0 [])
  where
  size = CodeUnits.length text
  charAt index = fromMaybe ' ' (CodeUnits.charAt index text)

  isIdentifierChar char =
    (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
      || (char >= '0' && char <= '9') || char == '_' || char > '\x7F'

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
    | otherwise = case charAt index of
        '"' -> scan (skipQuoted '"' (index + 1)) imports
        '\'' -> scan (skipQuoted '\'' (index + 1)) imports
        '`' -> scan (skipQuoted '`' (index + 1)) imports
        '/' | charAt (index + 1) == '/' -> scan (skipLine (index + 2)) imports
        '/' | charAt (index + 1) == '*' -> scan (skipComment (index + 2)) imports
        char | isIdentifierChar char ->
          let
            end = skipIdentifier (index + 1)
            dependency = if charAt end == '.' then
              case CodeUnits.slice index end text of
                "gopurs_runtime" -> Just "gopurs/output/gopurs_runtime"
                "math" -> Just "math"
                "sync" -> Just "sync"
                "unsafe" -> Just "unsafe"
                _ -> Nothing
              else Nothing
            nextImports = case dependency of
              Just path | not (Array.elem path imports) -> Array.snoc imports path
              _ -> imports
          in scan end nextImports
        _ -> scan (index + 1) imports
