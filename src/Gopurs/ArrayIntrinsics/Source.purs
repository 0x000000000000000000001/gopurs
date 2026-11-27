module Gopurs.ArrayIntrinsics.Source (arraySource) where

import Gopurs.GoAst (GoExpr(..), GoType(..))

-- Bind an already evaluated array without copying its elements. Native slices
-- retain their element type; boxed arrays expose their existing Value buffer.
-- Callers own the value binding and when this source binding is emitted.
arraySource :: String -> String -> GoType -> { assignment :: GoExpr, target :: String, elementType :: GoType }
arraySource valueName arrayName = case _ of
  TypeNativeArray inner ->
    { assignment: GoAssign arrayName (GoVar valueName), target: arrayName, elementType: inner }
  _ ->
    { assignment: GoAssign arrayName (GoCall (GoSelector (GoVar "gopurs_runtime") "ArrayValues") [ GoVar valueName ])
    , target: arrayName, elementType: TypeValue
    }
