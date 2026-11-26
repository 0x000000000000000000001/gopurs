module Consumer (consume, seedRunsCount) where

import Prelude
import Data.FoldableWithIndex (foldlWithIndex)
import Effect (Effect)
import Effect.Ref as Ref
import Effect.Unsafe (unsafePerformEffect)

seedRuns :: Ref.Ref Int
seedRuns = unsafePerformEffect (Ref.new 0)

seedRunsCount :: Effect Int
seedRunsCount = Ref.read seedRuns

consume :: Array Int -> Int
consume = foldlWithIndex (\index total value -> total + (1 + index `mod` 2) * value) 0
