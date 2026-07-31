module TestMangle where
import Prelude
import Effect.Console (log)
import PureScript.Backend.Optimizer.FreeVars (sanitizeName, localId)
import PureScript.Backend.Optimizer.CoreFn (Ident(..))
import PureScript.Backend.Optimizer.Syntax (Level(..))
import Data.Maybe (Maybe(..))

main = do
  log (sanitizeName "go")
  log (localId (Just (Ident "go")) (Level 1))
  log (sanitizeName "go__go")
  log (localId (Just (Ident "go__go")) (Level 1))
