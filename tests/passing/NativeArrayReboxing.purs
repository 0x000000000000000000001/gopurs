-- @inline Main.flatten never
module Main where

-- @dependencies: prelude arrays effect console refs assert

import Prelude
import Data.Array as Array
import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)

data Binding a = Binding a String
data Bind a = NonRec (Binding a) | Rec (Array (Binding a))

type Annotation = { line :: Int, label :: String }

flatten :: Bind Annotation -> Array (Binding Annotation)
flatten = case _ of
  NonRec binding -> [ binding ]
  Rec bindings -> bindings

describe :: Binding Annotation -> String
describe (Binding ann name) = name <> ":" <> ann.label <> ":" <> show ann.line

main :: Effect Unit
main = do
  input <- Ref.new (NonRec (Binding { line: 42, label: "typed" } "x"))
  value <- Ref.read input
  assertEqual { expected: [ "x:typed:42" ], actual: map describe (flatten value) }
  assertEqual { expected: 0, actual: Array.length (flatten (Rec [])) }
  log "Done"
