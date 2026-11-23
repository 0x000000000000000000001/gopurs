-- @inline export consume never
-- @inline export choose never
-- @inline export namedCallback never
module Main where

-- @dependencies: assert maybe prelude effect console refs

import Prelude

import Data.Maybe (Maybe(..))
import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)

type Input =
  { value :: Int
  , position :: { line :: Int, column :: Int }
  }

consume :: forall a. (Input -> Maybe a) -> Input -> Maybe a
consume callback input = callback input

choose :: Boolean -> Input -> Input -> Input
choose flag left right = if flag then left else right

namedCallback :: Input -> Maybe Int
namedCallback input = Just (input.value + input.position.line)

main :: Effect Unit
main = do
  inputRef <- Ref.new { value: 7, position: { line: 3, column: 2 } }
  input <- Ref.read inputRef
  callbackRef <- Ref.new \token ->
    if token.value > 0 then Just (token.value + token.position.column)
    else Nothing
  callback <- Ref.read callbackRef
  assertEqual { expected: Just 9, actual: consume callback input }
  assertEqual { expected: Nothing, actual: consume callback (input { value = 0 }) }

  -- A full record result must retain every field, including nested records.
  identityRef <- Ref.new (\token -> Just token)
  identity <- Ref.read identityRef
  assertEqual { expected: Just input, actual: consume identity input }

  -- Ordinary functions and partial application remain valid callback values.
  namedRef <- Ref.new namedCallback
  named <- Ref.read namedRef
  partialRef <- Ref.new (consume named)
  partial <- Ref.read partialRef
  assertEqual { expected: Just 10, actual: partial input }

  -- A closure may outlive the callback invocation that captured the record.
  factoryRef <- Ref.new (\token -> Just (\offset -> token.value + token.position.line + offset))
  factory <- Ref.read factoryRef
  case consume factory input of
    Nothing -> assertEqual { expected: true, actual: false }
    Just closure -> do
      closureRef <- Ref.new closure
      saved <- Ref.read closureRef
      assertEqual { expected: 14, actual: saved 4 }
      assertEqual { expected: 15, actual: saved 5 }

  flagRef <- Ref.new true
  flag <- Ref.read flagRef
  let other = input { value = 12 }
  assertEqual { expected: input, actual: choose flag input other }
  assertEqual { expected: other, actual: choose (not flag) input other }
  log "Done"
