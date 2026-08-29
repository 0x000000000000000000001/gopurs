module Scratch where
import Prelude
import Gopurs.Printer (printGoExpr)
import Gopurs.GoAst (GoExpr(..), GoType(..))
import Effect.Console (log)

main = do
  log (printGoExpr (GoCall (GoFuncLit [] [] (GoRaw "ret") TypeValue) []))
