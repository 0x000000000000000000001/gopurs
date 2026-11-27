-- @inline Subject.consume never
module Subject (consume) where

import Prelude
import Data.FoldableWithIndex (foldlWithIndex)

consume :: Array Int -> Int
consume values = foldlWithIndex (\index total value -> total + (index + 1) * value) 7 values
