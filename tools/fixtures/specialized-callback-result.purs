-- @inline Main.analyze never
-- @inline Main.lookup never
module Main where

-- @dependencies: prelude effect console maybe arrays foldable-traversable
import Prelude
import Data.Array (length)
import Data.Foldable (class Foldable, foldl)
import Data.Maybe (Maybe(..))
import Effect (Effect)
import Effect.Console (logShow)

newtype Analysis = Analysis { args :: Array Int }

class HasAnalysis a where
  analysisOf :: a -> Analysis

instance hasAnalysisInt :: HasAnalysis Int where
  analysisOf n = Analysis { args: [ n ] }

data Syntax a = Var Int | Other a

analyze :: forall a. HasAnalysis a => (Int -> Maybe String -> Maybe Analysis) -> Syntax a -> Int
analyze lookup expr = case expr of
  Var key -> case lookup key Nothing of
    Just (Analysis { args }) -> length args
    Nothing -> 0
  Other value -> case analysisOf value of
    Analysis { args } -> length args

lookup :: Int -> Maybe String -> Maybe Analysis
lookup key _ = if key == 0 then Nothing else Just (Analysis { args: [ 1, 2, 3 ] })

data Ints = Many (Array Int) | One (Maybe Int)

largest :: Ints -> Int
largest values = case values of
  Many xs -> refs xs
  One x -> refs x
  where
  refs :: forall f. Foldable f => f Int -> Int
  refs xs = foldl max (-1) xs

main :: Effect Unit
main = do
  logShow (analyze lookup (Var 0 :: Syntax Int))
  logShow (analyze lookup (Var 1 :: Syntax Int))
  logShow (analyze lookup (Other 1))
  logShow (largest (Many [ 1, 4 ]))
  logShow (largest (One (Just 7)))
