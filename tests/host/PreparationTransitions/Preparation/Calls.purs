module Preparation.Calls where

import Prelude
import Preparation.Apply as Apply
import Preparation.Cycle as Cycle
import Preparation.Provider as Provider

run :: Int -> Int
run value = Cycle.forward 2 value

provided :: Int -> Int
provided value = Apply.applyMethod Provider.method value
