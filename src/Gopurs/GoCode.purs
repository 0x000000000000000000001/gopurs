module Gopurs.GoCode
  ( GoCode
  , opaqueCode
  , referencedImports
  ) where

import Data.Array as Array
import Data.Maybe (Maybe(..))
import Data.String as String
import Data.String.Pattern (Pattern(..))

type GoCode = { text :: String, imports :: Array String }

-- Legacy opaque fragments retain their dependencies when they enter the AST.
-- This fallback inspects only the fragment, never a rendered module.
opaqueCode :: String -> GoCode
opaqueCode text = { text, imports: referencedImports text }

referencedImports :: String -> Array String
referencedImports text = Array.mapMaybe
  (\{ qualifier, path } -> if String.contains (Pattern qualifier) text then Just path else Nothing)
  [ { qualifier: "gopurs_runtime.", path: "gopurs/output/gopurs_runtime" }
  , { qualifier: "math.", path: "math" }
  , { qualifier: "sync.", path: "sync" }
  , { qualifier: "unsafe.", path: "unsafe" }
  ]
