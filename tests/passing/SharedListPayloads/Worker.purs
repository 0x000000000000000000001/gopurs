-- @inline export invoke never
-- @inline export prepend never
-- @inline export shift never
-- @inline export mapStack never
-- @inline export keep never
-- @inline export summarize never
module Worker where

import Prelude

import Data.List (List(..))
import Data.Tuple (Tuple(..))

data Phase = Begin | Between | End

derive instance eqPhase :: Eq Phase

instance showPhase :: Show Phase where
  show Begin = "B"
  show Between = "M"
  show End = "E"

type Position = { byteOffset :: Int, rowIndex :: Int }
type Stack = List (Tuple Position Phase)

invoke :: forall @a @b. (a -> b) -> a -> b
invoke callback value = callback value

prepend :: Position -> Phase -> Stack -> Stack
prepend position phase rest = Cons (Tuple position phase) rest

shift :: Int -> Stack -> Stack
shift amount = case _ of
  Nil -> Nil
  Cons (Tuple position phase) rest ->
    prepend (position { byteOffset = position.byteOffset + amount }) phase rest

mapStack :: (Tuple Position Phase -> Tuple Position Phase) -> Stack -> Stack
mapStack callback = case _ of
  Nil -> Nil
  Cons entry rest -> Cons (callback entry) (mapStack callback rest)

keep :: forall a. a -> Unit -> a
keep value _ = value

summarize :: Stack -> String
summarize = case _ of
  Nil -> ""
  Cons (Tuple position phase) rest ->
    show position.byteOffset <> ":" <> show position.rowIndex <> ":" <> show phase <> ";" <> summarize rest
