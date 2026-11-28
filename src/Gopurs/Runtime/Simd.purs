module Gopurs.Runtime.Simd
  ( simdGenericGoCode
  , runtimeSimdFiles
  , obsoleteRuntimeSimdFiles
  ) where

foreign import simdGenericGoCode :: String

runtimeSimdFiles :: Array { name :: String, content :: String }
runtimeSimdFiles =
  [ { name: "simd_generic.go", content: simdGenericGoCode }
  ]

obsoleteRuntimeSimdFiles :: Array String
obsoleteRuntimeSimdFiles =
  [ "simd_asm.go"
  , "simd_arm64.s"
  , "simd_amd64.s"
  ]
