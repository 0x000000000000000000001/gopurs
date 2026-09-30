module Main (main) where

import Effect (Effect)
import Effect.Aff (launchAff_)
import Gopurs.Driver (compile)

main :: Effect Unit
main = launchAff_ compile
