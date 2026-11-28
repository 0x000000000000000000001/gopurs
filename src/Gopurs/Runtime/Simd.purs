module Gopurs.Runtime.Simd
  ( simdAsmGoCode
  , simdGenericGoCode
  , simdArm64AsmCode
  , simdAmd64AsmCode
  , runtimeSimdFiles
  , obsoleteRuntimeSimdFiles
  ) where

foreign import simdAsmGoCode :: String
foreign import simdGenericGoCode :: String
foreign import simdArm64AsmCode :: String
foreign import simdAmd64AsmCode :: String

runtimeSimdFiles :: Array { name :: String, content :: String }
runtimeSimdFiles =
  [ { name: "simd_asm.go", content: simdAsmGoCode }
  , { name: "simd_generic.go", content: simdGenericGoCode }
  , { name: "simd_arm64.s", content: simdArm64AsmCode }
  , { name: "simd_amd64.s", content: simdAmd64AsmCode }
  ]

obsoleteRuntimeSimdFiles :: Array String
obsoleteRuntimeSimdFiles = []
