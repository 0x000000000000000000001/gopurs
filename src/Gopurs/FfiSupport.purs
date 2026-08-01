module Gopurs.FfiSupport
  ( extractFfiAst
  ) where

import Effect (Effect)

foreign import extractFfiAstImpl :: String -> String -> Effect String

extractFfiAst :: String -> String -> Effect String
extractFfiAst = extractFfiAstImpl
