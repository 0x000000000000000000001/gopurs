module Gopurs.ArrayIntrinsics.Source (arraySource) where

import Prelude

import Gopurs.GoAst (rawGo, GoExpr(..), GoType(..))

-- Bind an already evaluated array without copying its elements. Native slices
-- retain their element type; boxed arrays expose their existing Value buffer.
-- Callers own the value binding and when this source binding is emitted.
arraySource :: String -> String -> GoType -> { assignment :: GoExpr, target :: String, elementType :: GoType }
arraySource valueName arrayName = case _ of
  TypeNativeArray inner ->
    { assignment: GoAssign arrayName (GoVar valueName), target: arrayName, elementType: inner }
  _ ->
    { assignment: GoAssign arrayName (GoCall (rawGo "(*[]gopurs_runtime.Value)") [ GoSelector (GoVar valueName) "UnsafePtr" ])
    , target: "*" <> arrayName, elementType: TypeValue
    }
