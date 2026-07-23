module TestDecode where
import Prelude
import Data.Either
import Effect.Console (logShow)
import Data.Argonaut.Core (fromNumber)
import PureScript.Backend.Optimizer.CoreFn.Json (decodeModule) -- wait, decodeInt is not exported
