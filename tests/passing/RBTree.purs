module Main where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)

-- The ultimate functional benchmark: Chris Okasaki's Red-Black Tree.
-- This heavily tests complex pattern matching, structural sharing,
-- rapid allocation/deallocation of ADTs, and balancing rotations.

data Color = R | B
data Tree = E | T Color Tree Int Tree

balance :: Color -> Tree -> Int -> Tree -> Tree
balance B (T R (T R a x b) y c) z d = T R (T B a x b) y (T B c z d)
balance B (T R a x (T R b y c)) z d = T R (T B a x b) y (T B c z d)
balance B a x (T R (T R b y c) z d) = T R (T B a x b) y (T B c z d)
balance B a x (T R b y (T R c z d)) = T R (T B a x b) y (T B c z d)
balance color a x b = T color a x b

insert :: Int -> Tree -> Tree
insert x s = makeBlack (ins s)
  where
  ins E = T R E x E
  ins (T color a y b) =
    if x < y then balance color (ins a) y b
    else if x > y then balance color a y (ins b)
    else T color a y b
    
  makeBlack (T _ a y b) = T B a y b
  makeBlack E = E

-- Tail-recursive insertion loop
buildTree :: Int -> Tree -> Tree
buildTree 0 acc = acc
buildTree n acc = buildTree (n - 1) (insert n acc)

-- Computes the maximum depth of the tree to force traversal
depth :: Tree -> Int
depth E = 0
depth (T _ a _ b) = 1 + max (depth a) (depth b)

max :: Int -> Int -> Int
max x y = if x > y then x else y

describe :: Effect Unit
describe = log "Red-Black Tree (100k Worst-Case Insertions):"

act :: Effect Unit
act = logShow $ depth (buildTree 100000 E)
main :: Effect Unit
main = do
  describe
  act
