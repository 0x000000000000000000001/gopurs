-- @inline export runSummary never
-- @inline export useSummary never
module Worker where

import Prelude
import Data.Either (Either(..))

type Summary = { id :: Int, name :: String, active :: Boolean }

-- Shape 1: class-constraint dictionary used by a generic helper.
-- The newtype is erased in CoreFn, so the decoded payload is a record.
newtype Named = Named Summary

class Decode a where
  decode :: Int -> Either String a

instance decodeNamed :: Decode Named where
  decode n =
    if n < 0 then Left "negative"
    else Right (Named { id: n + 1, name: "alpha", active: true })

runDecode :: forall a. Decode a => Int -> Either String a
runDecode n = decode n

runSummary :: Int -> Either String Summary
runSummary n = case runDecode n of
  Left e -> Left e
  Right (Named s) -> Right s

-- Shape 2: explicit record dictionary, as the Argonaut combinators build them.
type Decoder a = { decode :: Int -> Either String a }

summaryDecoder :: Decoder Summary
summaryDecoder = { decode: \n -> if n < 0 then Left "negative" else Right { id: n + 2, name: "beta", active: false } }

useDecoder :: forall a. Decoder a -> Int -> Either String a
useDecoder d n = d.decode n

useSummary :: Int -> Either String Summary
useSummary = useDecoder summaryDecoder
