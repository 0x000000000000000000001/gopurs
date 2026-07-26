module Gopurs.FfiSupport
  ( appendFfiWrappers
  ) where

foreign import appendFfiWrappersImpl :: String -> String -> String

appendFfiWrappers :: String -> String -> String
appendFfiWrappers = appendFfiWrappersImpl
