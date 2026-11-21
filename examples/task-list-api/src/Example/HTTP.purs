module Example.HTTP
  ( Request
  , Response
  , serve
  ) where

import Prelude

import Effect (Effect)
import Effect.Aff (Aff, bracket)
import Effect.Aff.Compat (EffectFnAff, fromEffectFnAff)
import Effect.Uncurried (EffectFn1, EffectFn2, mkEffectFn2, runEffectFn1)

type Request = { method :: String, path :: String, body :: String }
type Response = { status :: Int, body :: String }
foreign import data Server :: Type

-- | Runs until SIGINT, SIGTERM or Aff cancellation. Every accepted request
-- | must eventually invoke its reply callback, including on application errors.
-- | The server drains these callbacks before returning to its resource bracket.
serve :: String -> (Request -> (Response -> Effect Unit) -> Effect Unit) -> Aff Unit
serve address handler = bracket
  (fromEffectFnAff $ startImpl address $ mkEffectFn2 \request reply -> handler request (runEffectFn1 reply))
  (fromEffectFnAff <<< stopImpl)
  (fromEffectFnAff <<< waitImpl)

foreign import startImpl
  :: String
  -> EffectFn2 Request (EffectFn1 Response Unit) Unit
  -> EffectFnAff Server
foreign import waitImpl :: Server -> EffectFnAff Unit
foreign import stopImpl :: Server -> EffectFnAff Unit
