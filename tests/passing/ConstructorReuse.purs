-- Keep the reconstruction patterns visible in the generated Go snapshot.
-- @inline export setBlack never
-- @inline export setSeven never
-- @inline export setTrue never
-- @inline export swapChildren never
-- @inline export mixSources never
-- @inline export setBlackTrue never
-- @inline export toSecond never
-- @inline export positiveZero never
module Main where

-- @dependencies: assert prelude effect console refs

import Prelude

import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)

data Color = Red | Black
data Tree = Empty | Node Color Tree Int Boolean Tree

setBlack :: Tree -> Tree
setBlack (Node _ left key active right) = Node Black left key active right
setBlack Empty = Empty

setSeven :: Tree -> Tree
setSeven (Node color left _ active right) = Node color left 7 active right
setSeven Empty = Empty

setTrue :: Tree -> Tree
setTrue (Node color left key _ right) = Node color left key true right
setTrue Empty = Empty

-- None of these three patterns satisfies the single-field replacement rule.
swapChildren :: Tree -> Tree
swapChildren (Node _ left key active right) = Node Black right key active left
swapChildren Empty = Empty

mixSources :: Tree -> Tree -> Tree
mixSources (Node _ left key active _) (Node _ _ _ _ right) =
  Node Black left key active right
mixSources source _ = source

setBlackTrue :: Tree -> Tree
setBlackTrue (Node _ left key _ right) = Node Black left key true right
setBlackTrue Empty = Empty

data Tagged = First Color Int | Second Color Int

toSecond :: Tagged -> Tagged
toSecond (First _ key) = Second Black key
toSecond (Second _ key) = Second Black key

data NumberCell = NumberCell Number Int

-- Numeric equality cannot justify reusing a cell containing negative zero.
positiveZero :: NumberCell -> NumberCell
positiveZero (NumberCell _ marker) = NumberCell 0.0 marker

isNegativeZero :: NumberCell -> Boolean
isNegativeZero (NumberCell value _) = 1.0 / value < 0.0

numberMarker :: NumberCell -> Int
numberMarker (NumberCell _ marker) = marker

colorName :: Color -> String
colorName Red = "R"
colorName Black = "B"

nodeText :: String -> Int -> Boolean -> String -> String -> String
nodeText color key active left right =
  color <> ":" <> show key <> ":" <> show active <> "[" <> left <> "][" <> right <> "]"

render :: Tree -> String
render Empty = "E"
render (Node color left key active right) =
  nodeText (colorName color) key active (render left) (render right)

renderTagged :: Tagged -> String
renderTagged (First color key) = "first:" <> colorName color <> ":" <> show key
renderTagged (Second color key) = "second:" <> colorName color <> ":" <> show key

expectTree :: String -> Tree -> Effect Unit
expectTree expected actual = assertEqual { expected, actual: render actual }

main :: Effect Unit
main = do
  let
    left = Node Black Empty (-3) true Empty
    right = Node Red Empty 9 false Empty
    alternateRight = Node Black Empty 42 true Empty
    leftText = "B:-3:true[E][E]"
    rightText = "R:9:false[E][E]"
    redText = nodeText "R" 17 false leftText rightText
    blackText = nodeText "B" 17 false leftText rightText
    sevenText = nodeText "R" 7 false leftText rightText
    trueText = nodeText "R" 17 true leftText rightText
  -- Ref.read makes the constructor, tag and field values runtime inputs.
  inputsRef <- Ref.new
    { red: Node Red left 17 false right
    , black: Node Black left 17 false right
    , seven: Node Red left 7 false right
    , flagged: Node Red left 17 true right
    , empty: Empty
    , other: Node Black Empty 91 true alternateRight
    }
  inputs <- Ref.read inputsRef

  expectTree blackText (setBlack inputs.red)
  expectTree blackText (setBlack inputs.black)
  expectTree "E" (setBlack inputs.empty)
  expectTree sevenText (setSeven inputs.red)
  expectTree sevenText (setSeven inputs.seven)
  expectTree trueText (setTrue inputs.red)
  expectTree trueText (setTrue inputs.flagged)
  -- Expected strings contain no tree references that an incorrect mutation
  -- could change together with the actual value.
  expectTree redText inputs.red
  expectTree blackText inputs.black
  expectTree "E" (setSeven inputs.empty)
  expectTree "E" (setTrue inputs.empty)

  -- Keep an old version across a subsequent update of its shared root.
  oldVersionRef <- Ref.new (setBlack inputs.red)
  oldVersion <- Ref.read oldVersionRef
  newVersionRef <- Ref.new (setSeven oldVersion)
  newVersion <- Ref.read newVersionRef
  expectTree (nodeText "B" 7 false leftText rightText) newVersion
  expectTree blackText oldVersion
  expectTree redText inputs.red

  expectTree (nodeText "B" 17 false rightText leftText) (swapChildren inputs.black)
  expectTree (nodeText "B" 17 false leftText "B:42:true[E][E]")
    (mixSources inputs.black inputs.other)
  expectTree blackText inputs.black
  expectTree (nodeText "B" 17 true leftText rightText) (setBlackTrue inputs.black)
  expectTree blackText inputs.black

  taggedRef <- Ref.new (First Black 41)
  tagged <- Ref.read taggedRef
  assertEqual { expected: "second:B:41", actual: renderTagged (toSecond tagged) }
  assertEqual { expected: "first:B:41", actual: renderTagged tagged }

  -- Negate a runtime zero: a Go constant -0.0 would lose the sign already.
  zeroRef <- Ref.new 0.0
  zero <- Ref.read zeroRef
  numberRef <- Ref.new (NumberCell (negate zero) 5)
  number <- Ref.read numberRef
  assertEqual { expected: true, actual: isNegativeZero number }
  changedRef <- Ref.new (positiveZero number)
  changed <- Ref.read changedRef
  assertEqual { expected: false, actual: isNegativeZero changed }
  assertEqual { expected: true, actual: isNegativeZero number }
  assertEqual { expected: 5, actual: numberMarker changed }

  log "Done"
