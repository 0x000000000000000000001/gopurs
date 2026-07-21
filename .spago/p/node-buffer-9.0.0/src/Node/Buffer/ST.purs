module Node.Buffer.ST
  ( STBuffer
  , run
  , create
  , alloc
  , allocUnsafe
  , allocUnsafeSlow
  , compareParts
  , freeze
  , unsafeFreeze
  , thaw
  , unsafeThaw
  , fromArray
  , fromString
  , fromArrayBuffer
  , toArrayBuffer
  , read
  , readString
  , toString
  , toString'
  , write
  , writeString
  , toArray
  , getAtOffset
  , setAtOffset
  , slice
  , size
  , concat
  , concat'
  , copy
  , fill
  , swap16
  , swap32
  , swap64
  , transcode
  ) where

import Prelude

import Control.Monad.ST (ST, Region)
import Control.Monad.ST as ST
import Data.ArrayBuffer.Types (ArrayBuffer)
import Data.Maybe (Maybe)
import Node.Buffer (BufferValueType, Octet, Offset)
import Node.Buffer as Buffer
import Node.Buffer.Class (class MutableBuffer)
import Node.Buffer.Immutable (ImmutableBuffer)
import Node.Encoding (Encoding)
import Unsafe.Coerce (unsafeCoerce)

-- | A reference to a mutable Buffer for use with `ST`
-- |
-- | The type parameter represents the memory region which the (STBuffer h) belongs to.
foreign import data STBuffer :: Region -> Type

-- | Runs an ST creating an `STBuffer` then freezes the (STBuffer h) and returns
-- | it, without unneccessary copying.
run :: (forall h. ST h (STBuffer h)) -> ImmutableBuffer
run st = ST.run (st >>= unsafeFreeze)

instance mutableBufferST :: MutableBuffer (STBuffer h) (ST h) where
  create = create
  freeze = freeze
  unsafeFreeze = unsafeFreeze
  thaw = thaw
  unsafeThaw = unsafeThaw
  fromArray = fromArray
  fromString = fromString
  fromArrayBuffer = fromArrayBuffer
  toArrayBuffer = toArrayBuffer
  read = read
  readString = readString
  toString = toString
  write = write
  writeString = writeString
  toArray = toArray
  getAtOffset = getAtOffset
  setAtOffset = setAtOffset
  slice = slice
  size = size
  concat = concat
  concat' = concat'
  copy = copy
  fill = fill

-- | Creates a new `STBuffer`. Alias of `alloc`.
create :: forall h. Int -> ST h (STBuffer h)
create = alloc

-- | Creates a new `STBuffer`.
alloc :: forall h. Int -> ST h (STBuffer h)
alloc = unsafeCoerce Buffer.alloc

-- | Creates a new `STBuffer` of the specified size. Unsafe because it reuses memory from a pool
-- | and may contain sensitive data. See the Node docs: 
-- | https://nodejs.org/docs/latest-v18.x/api/buffer.html#what-makes-bufferallocunsafe-and-bufferallocunsafeslow-unsafe
allocUnsafe :: forall h. Int -> ST h (STBuffer h)
allocUnsafe = unsafeCoerce Buffer.allocUnsafe

-- | Creates a new `STBuffer` of the specified size. Unsafe because it reuses memory from a pool
-- | and may contain sensitive data. See the Node docs: 
-- | https://nodejs.org/docs/latest-v18.x/api/buffer.html#what-makes-bufferallocunsafe-and-bufferallocunsafeslow-unsafe
allocUnsafeSlow :: forall h. Int -> ST h (STBuffer h)
allocUnsafeSlow = unsafeCoerce Buffer.allocUnsafeSlow

compareParts :: forall h. STBuffer h -> STBuffer h -> Int -> Int -> Int -> Int -> ST h Ordering
compareParts = unsafeCoerce Buffer.compareParts

freeze :: forall h. (STBuffer h) -> ST h ImmutableBuffer
freeze = unsafeCoerce Buffer.freeze

unsafeFreeze :: forall h. STBuffer h -> ST h ImmutableBuffer
unsafeFreeze = unsafeCoerce Buffer.unsafeFreeze

thaw :: forall h. ImmutableBuffer -> ST h (STBuffer h)
thaw = unsafeCoerce Buffer.thaw

unsafeThaw :: forall h. ImmutableBuffer -> ST h (STBuffer h)
unsafeThaw = unsafeCoerce Buffer.unsafeThaw

fromArray :: forall h. Array Octet -> ST h (STBuffer h)
fromArray = unsafeCoerce Buffer.fromArray

fromString :: forall h. String -> Encoding -> ST h (STBuffer h)
fromString = unsafeCoerce Buffer.fromString

fromArrayBuffer :: forall h. ArrayBuffer -> ST h (STBuffer h)
fromArrayBuffer = unsafeCoerce Buffer.fromArrayBuffer

toArrayBuffer :: forall h. STBuffer h -> ST h ArrayBuffer
toArrayBuffer = unsafeCoerce Buffer.toArrayBuffer

read :: forall h. BufferValueType -> Offset -> STBuffer h -> ST h Number
read = unsafeCoerce Buffer.read

readString :: forall h. Encoding -> Offset -> Offset -> STBuffer h -> ST h String
readString = unsafeCoerce Buffer.readString

toString :: forall h. Encoding -> STBuffer h -> ST h String
toString = unsafeCoerce Buffer.toString

toString' :: forall h. Encoding -> Offset -> Offset -> STBuffer h -> ST h String
toString' = unsafeCoerce Buffer.toString'

write :: forall h. BufferValueType -> Number -> Offset -> STBuffer h -> ST h Unit
write = unsafeCoerce Buffer.write

writeString :: forall h. Encoding -> Offset -> Int -> String -> STBuffer h -> ST h Int
writeString =
  unsafeCoerce Buffer.writeString

toArray :: forall h. STBuffer h -> ST h (Array Octet)
toArray = unsafeCoerce Buffer.toArray

getAtOffset :: forall h. Offset -> STBuffer h -> ST h (Maybe Octet)
getAtOffset = unsafeCoerce Buffer.getAtOffset

setAtOffset :: forall h. Octet -> Offset -> STBuffer h -> ST h Unit
setAtOffset = unsafeCoerce Buffer.setAtOffset

slice :: forall h. Offset -> Offset -> STBuffer h -> STBuffer h
slice = unsafeCoerce Buffer.slice

size :: forall h. STBuffer h -> ST h Int
size = unsafeCoerce Buffer.size

concat :: forall h. Array (STBuffer h) -> ST h (STBuffer h)
concat = unsafeCoerce Buffer.concat

concat' :: forall h. Array (STBuffer h) -> Int -> ST h (STBuffer h)
concat' = unsafeCoerce Buffer.concat'

copy :: forall h. Offset -> Offset -> STBuffer h -> Offset -> STBuffer h -> ST h Int
copy = unsafeCoerce Buffer.copy

fill :: forall h. Octet -> Offset -> Offset -> STBuffer h -> ST h Unit
fill = unsafeCoerce Buffer.fill

swap16 :: forall h. STBuffer h -> ST h (STBuffer h)
swap16 = unsafeCoerce Buffer.swap16

swap32 :: forall h. STBuffer h -> ST h (STBuffer h)
swap32 = unsafeCoerce Buffer.swap32

swap64 :: forall h. STBuffer h -> ST h (STBuffer h)
swap64 = unsafeCoerce Buffer.swap64

transcode :: forall h. STBuffer h -> Encoding -> Encoding -> ST h (STBuffer h)
transcode = unsafeCoerce Buffer.transcode
