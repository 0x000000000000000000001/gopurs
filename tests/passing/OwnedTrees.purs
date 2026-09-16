-- @inline export put never
-- @inline export rebalance never
-- @inline export descend never
-- @inline export blacken never
-- @inline export build never
-- @inline export mixed never
-- @inline export render never
-- @inline export graftAndPut never
module Main where

-- @dependencies: assert prelude effect console refs

import Prelude

import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)

-- Deliberately different names from the benchmark: reuse must follow proofs.
data Paint = Crimson | Onyx
data SearchTree = Tip | Branch Paint SearchTree Int SearchTree

rebalance :: Paint -> SearchTree -> Int -> SearchTree -> SearchTree
rebalance Onyx (Branch Crimson (Branch Crimson a x b) y c) z d =
  Branch Crimson (Branch Onyx a x b) y (Branch Onyx c z d)
rebalance Onyx (Branch Crimson a x (Branch Crimson b y c)) z d =
  Branch Crimson (Branch Onyx a x b) y (Branch Onyx c z d)
rebalance Onyx a x (Branch Crimson (Branch Crimson b y c) z d) =
  Branch Crimson (Branch Onyx a x b) y (Branch Onyx c z d)
rebalance Onyx a x (Branch Crimson b y (Branch Crimson c z d)) =
  Branch Crimson (Branch Onyx a x b) y (Branch Onyx c z d)
rebalance color a x b = Branch color a x b

put :: Int -> SearchTree -> SearchTree
put value tree = blacken (descend value tree)

descend :: Int -> SearchTree -> SearchTree
descend value Tip = Branch Crimson Tip value Tip
descend value (Branch color left key right)
  | value < key = rebalance color (descend value left) key right
  | value > key = rebalance color left key (descend value right)
  | otherwise = Branch color left key right

blacken :: SearchTree -> SearchTree
blacken Tip = Tip
blacken (Branch _ left key right) = Branch Onyx left key right

isRed :: SearchTree -> Boolean
isRed (Branch Crimson _ _ _) = true
isRed _ = false

render :: SearchTree -> String
render Tip = "E"
render (Branch color left key right) =
  (case color of
    Crimson -> "R"
    Onyx -> "B") <> show key <> "(" <> render left <> ")(" <> render right <> ")"

audit :: Int -> Int -> SearchTree -> { valid :: Boolean, blackHeight :: Int, size :: Int }
audit _ _ Tip = { valid: true, blackHeight: 1, size: 0 }
audit lower upper tree@(Branch color left key right) =
  let
    l = audit lower key left
    r = audit key upper right
    black = case color of
      Onyx -> 1
      Crimson -> 0
  in
    { valid: l.valid && r.valid && lower < key && key < upper
        && l.blackHeight == r.blackHeight
        && (not (isRed tree) || (not (isRed left) && not (isRed right)))
    , blackHeight: l.blackHeight + black
    , size: l.size + r.size + 1
    }

check :: Int -> SearchTree -> Effect Unit
check expectedSize tree = do
  let result = audit (-1000000) 1000000 tree
  assertEqual { expected: true, actual: result.valid }
  assertEqual { expected: false, actual: isRed tree }
  assertEqual { expected: expectedSize, actual: result.size }

three :: Int -> Int -> Int -> Effect Unit
three a b c = do
  input <- Ref.new { a, b, c }
  values <- Ref.read input
  output <- Ref.new (put values.c (put values.b (put values.a Tip)))
  tree <- Ref.read output
  check 3 tree
  assertEqual { expected: "B2(B1(E)(E))(B3(E)(E))", actual: render tree }

build :: Int -> Int -> SearchTree -> SearchTree
build 0 _ acc = acc
build remaining direction acc =
  build (remaining - 1) direction (put (remaining * direction) acc)

mixed :: Int -> SearchTree -> SearchTree
mixed 0 acc = acc
mixed remaining acc = mixed (remaining - 1) (put (mod (remaining * 37) 127) acc)

-- Every prior version stays observable after all later recursive insertions.
snapshots :: Int -> SearchTree -> Effect Unit
snapshots 0 tree = check 32 tree
snapshots remaining tree = do
  oldRef <- Ref.new tree
  old <- Ref.read oldRef
  expectedRef <- Ref.new (render old)
  expected <- Ref.read expectedRef
  nextRef <- Ref.new (put remaining old)
  next <- Ref.read nextRef
  check (33 - remaining) next
  snapshots (remaining - 1) next
  assertEqual { expected, actual: render old }

leftChild :: SearchTree -> SearchTree
leftChild Tip = Tip
leftChild (Branch _ left _ _) = left

-- The parent is freshly constructed, but its left child is externally shared.
graftAndPut :: SearchTree -> SearchTree
graftAndPut child = put 5 (Branch Onyx child 50 (Branch Onyx Tip 90 Tip))

main :: Effect Unit
main = do
  three 3 2 1
  three 3 1 2
  three 1 3 2
  three 1 2 3

  countRef <- Ref.new 127
  count <- Ref.read countRef
  ascRef <- Ref.new (build count (-1) Tip)
  descRef <- Ref.new (build count 1 Tip)
  mixRef <- Ref.new (mixed count Tip)
  asc <- Ref.read ascRef
  desc <- Ref.read descRef
  shuffled <- Ref.read mixRef
  check count asc
  check count desc
  check count shuffled
  -- Repeated keys must leave both contents and size unchanged.
  expectedRef <- Ref.new (render shuffled)
  expected <- Ref.read expectedRef
  duplicateRef <- Ref.new (mixed count shuffled)
  duplicate <- Ref.read duplicateRef
  check count duplicate
  assertEqual { expected, actual: render duplicate }
  assertEqual { expected, actual: render shuffled }

  snapshots 32 Tip

  rootRef <- Ref.new (Branch Onyx (Branch Onyx Tip 10 Tip) 50 (Branch Onyx Tip 90 Tip))
  root <- Ref.read rootRef
  childRef <- Ref.new (leftChild root)
  child <- Ref.read childRef
  rootTextRef <- Ref.new (render root)
  childTextRef <- Ref.new (render child)
  rootText <- Ref.read rootTextRef
  childText <- Ref.read childTextRef
  updatedRef <- Ref.new (put 5 root)
  updated <- Ref.read updatedRef
  check 4 updated
  assertEqual { expected: rootText, actual: render root }
  assertEqual { expected: childText, actual: render child }
  graftedRef <- Ref.new (graftAndPut child)
  grafted <- Ref.read graftedRef
  check 4 grafted
  assertEqual { expected: childText, actual: render child }
  assertEqual { expected: rootText, actual: render root }

  log "Done"
