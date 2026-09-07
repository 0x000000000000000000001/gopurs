module Main where

-- @dependencies: assert prelude effect console refs
-- @snapshot-ffi

import Prelude

import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Test.Assert (assertEqual)

foreign import returnInt64 :: Int -> Int
foreign import returnInt :: Int -> Int
foreign import returnInt64Array :: Array Int -> Array Int
foreign import returnIntArray :: Array Int -> Array Int

-- These Go functions return native values through any, not pre-boxed Values.
foreign import dynamicInt64 :: Int -> Int
foreign import dynamicInt :: Int -> Int
foreign import dynamicString :: String -> String
foreign import dynamicBoolean :: Boolean -> Boolean

checkConsumer :: (Int -> Int) -> Effect Unit
checkConsumer consume = do
  assertEqual { expected: 50, actual: consume (returnInt64 (-7)) }
  assertEqual { expected: 26, actual: consume (returnInt 5) }
  assertEqual { expected: 10, actual: consume (dynamicInt64 (-3)) }
  assertEqual { expected: 1, actual: consume (dynamicInt 0) }

main :: Effect Unit
main = do
  -- PureScript Int bounds are shared by the Go and JavaScript fixtures.
  -- No arithmetic overflow or narrowing beyond that range is assumed.
  assertEqual { expected: 0, actual: returnInt64 0 }
  assertEqual { expected: -7, actual: returnInt64 (-7) }
  assertEqual { expected: 42, actual: returnInt64 42 }
  assertEqual { expected: -2147483648, actual: returnInt64 (-2147483648) }
  assertEqual { expected: 2147483647, actual: returnInt64 2147483647 }

  assertEqual { expected: 0, actual: returnInt 0 }
  assertEqual { expected: -7, actual: returnInt (-7) }
  assertEqual { expected: 42, actual: returnInt 42 }
  assertEqual { expected: -2147483648, actual: returnInt (-2147483648) }
  assertEqual { expected: 2147483647, actual: returnInt 2147483647 }

  assertEqual { expected: [], actual: returnInt64Array [] }
  assertEqual
    { expected: [0, -7, 42, -2147483648, 2147483647]
    , actual: returnInt64Array [0, -7, 42, -2147483648, 2147483647]
    }
  assertEqual { expected: [], actual: returnIntArray [] }
  assertEqual
    { expected: [0, -7, 42, -2147483648, 2147483647]
    , actual: returnIntArray [0, -7, 42, -2147483648, 2147483647]
    }

  -- Keep the dynamic-return fallback valid for numeric and non-numeric values.
  assertEqual { expected: 0, actual: dynamicInt64 0 }
  assertEqual { expected: -2147483648, actual: dynamicInt64 (-2147483648) }
  assertEqual { expected: 2147483647, actual: dynamicInt64 2147483647 }
  assertEqual { expected: -7, actual: dynamicInt (-7) }
  assertEqual { expected: -2147483648, actual: dynamicInt (-2147483648) }
  assertEqual { expected: 2147483647, actual: dynamicInt 2147483647 }
  assertEqual { expected: "fallback", actual: dynamicString "fallback" }
  assertEqual { expected: false, actual: dynamicBoolean false }
  assertEqual { expected: true, actual: dynamicBoolean true }

  -- Keep the consumer opaque so generated Go must pass the FFI result to it.
  consumerRef <- Ref.new (\value -> value * value + 1)
  consumer <- Ref.read consumerRef
  checkConsumer consumer

  log "Done"
