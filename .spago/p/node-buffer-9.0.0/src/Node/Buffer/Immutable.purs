-- | Immutable buffers and associated operations.
module Node.Buffer.Immutable
  ( ImmutableBuffer
  , compareParts
  , create
  , alloc
  , fromArray
  , fromString
  , fromArrayBuffer
  , read
  , readString
  , toString
  , toString'
  , toArray
  , toArrayBuffer
  , getAtOffset
  , concat
  , concat'
  , slice
  , size
  ) where

import Prelude

import Data.ArrayBuffer.Types (ArrayBuffer)
import Data.Function.Uncurried (Fn2, Fn3, Fn4, runFn2, runFn3, runFn4)
import Data.Maybe (Maybe)
import Data.Nullable (Nullable, toMaybe)
import Effect (Effect)
import Effect.Uncurried (EffectFn6, runEffectFn6)
import Node.Buffer.Types (BufferValueType, Octet, Offset)
import Node.Encoding (Encoding, encodingToNode)
import Partial.Unsafe (unsafeCrashWith)

-- | An immutable buffer that exists independently of any memory region or effect.
foreign import data ImmutableBuffer :: Type

instance showBuffer :: Show ImmutableBuffer where
  show = showImpl

foreign import showImpl :: ImmutableBuffer -> String

instance eqBuffer :: Eq ImmutableBuffer where
  eq a b = runFn2 eqImpl a b

foreign import eqImpl :: Fn2 ImmutableBuffer ImmutableBuffer Boolean

instance ordBuffer :: Ord ImmutableBuffer where
  compare a b =
    case runFn2 compareImpl a b of
      x | x < 0 -> LT
      x | x > 0 -> GT
      _ -> EQ

foreign import compareImpl :: Fn2 ImmutableBuffer ImmutableBuffer Int

-- | Creates a new buffer of the specified size. Alias for `alloc`.
create :: Int -> ImmutableBuffer
create = alloc

-- | Creates a new buffer of the specified size.
foreign import alloc :: Int -> ImmutableBuffer

-- | Creates a new buffer from an array of octets, sized to match the array.
foreign import fromArray :: Array Octet -> ImmutableBuffer

-- | Creates a buffer view from a JS ArrayByffer without copying data.
foreign import fromArrayBuffer :: ArrayBuffer -> ImmutableBuffer

-- | Creates a new buffer from a string with the specified encoding, sized to match the string.
fromString :: String -> Encoding -> ImmutableBuffer
fromString str enc = runFn2 fromStringImpl str $ encodingToNode enc

foreign import fromStringImpl :: Fn2 String String ImmutableBuffer

compareParts :: ImmutableBuffer -> ImmutableBuffer -> Offset -> Offset -> Offset -> Offset -> Effect Ordering
compareParts src target targetStart targetEnd sourceStart sourceEnd =
  runEffectFn6 comparePartsImpl src target targetStart targetEnd sourceStart sourceEnd <#> case _ of
    -1 -> LT
    0 -> EQ
    1 -> GT
    x -> unsafeCrashWith $ "Impossible: Invalid value: " <> show x

foreign import comparePartsImpl :: EffectFn6 ImmutableBuffer ImmutableBuffer Int Int Int Int Int

-- | Reads a numeric value from a buffer at the specified offset.
read :: BufferValueType -> Offset -> ImmutableBuffer -> Number
read ty off buf = runFn3 readImpl (show ty) off buf

foreign import readImpl :: Fn3 String Offset ImmutableBuffer Number

-- | Reads a section of a buffer as a string with the specified encoding.
readString :: Encoding -> Offset -> Offset -> ImmutableBuffer -> String
readString enc start end buf = runFn4 readStringImpl (encodingToNode enc) start end buf

foreign import readStringImpl :: Fn4 String Offset Offset ImmutableBuffer String

-- | Reads the buffer as a string with the specified encoding.
toString :: Encoding -> ImmutableBuffer -> String
toString enc buf = runFn2 toStringImpl (encodingToNode enc) buf

foreign import toStringImpl :: Fn2 String ImmutableBuffer String

toString' :: Encoding -> Offset -> Offset -> ImmutableBuffer -> String
toString' enc start end buf = runFn4 toStringSubImpl enc start end buf

foreign import toStringSubImpl :: Fn4 Encoding Offset Offset ImmutableBuffer String

-- | Creates an array of octets from a buffer's contents.
foreign import toArray :: ImmutableBuffer -> Array Octet

-- | Creates an `ArrayBuffer` by copying a buffer's contents.
foreign import toArrayBuffer :: ImmutableBuffer -> ArrayBuffer

-- | Reads an octet from a buffer at the specified offset.
getAtOffset :: Offset -> ImmutableBuffer -> Maybe Octet
getAtOffset off buf = toMaybe $ runFn2 getAtOffsetImpl off buf

foreign import getAtOffsetImpl :: Fn2 Offset ImmutableBuffer (Nullable Octet)

-- | Concatenates a list of buffers.
foreign import concat :: Array ImmutableBuffer -> ImmutableBuffer

-- | Concatenates a list of buffers, combining them into a new buffer of the
-- | specified length.
concat' :: Array ImmutableBuffer -> Int -> ImmutableBuffer
concat' a i = runFn2 concatToLength a i

foreign import concatToLength :: Fn2 (Array ImmutableBuffer) Int ImmutableBuffer

-- | Creates a new buffer slice that shares the memory of the original buffer.
slice :: Offset -> Offset -> ImmutableBuffer -> ImmutableBuffer
slice start end buf = runFn3 sliceImpl start end buf

foreign import sliceImpl :: Fn3 Offset Offset ImmutableBuffer ImmutableBuffer

-- | Returns the size of a buffer.
foreign import size :: ImmutableBuffer -> Int
