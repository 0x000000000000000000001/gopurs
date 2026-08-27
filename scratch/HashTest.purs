module HashTest where
import Prelude
import Effect.Console as Console
import PureScript.Backend.Optimizer.Monomorphize (mangleType)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..))
import PureScript.Backend.Optimizer.Convert (hashString)
import Effect (Effect)

main :: Effect Unit
main = do
  Console.log (show (hashString "Any"))
  Console.log (show (hashString (mangleType Any)))
