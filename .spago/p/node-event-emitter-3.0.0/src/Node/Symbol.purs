module Node.Symbol where

import Prelude

import Data.Function.Uncurried (Fn1, runFn1)
import Data.Maybe (Maybe)
import Data.Nullable (Nullable, toMaybe)
import Effect (Effect)
import Effect.Uncurried (EffectFn1, runEffectFn1)

foreign import data JsSymbol :: Type

instance Show JsSymbol where
  show x = showSymbol x

showSymbol :: JsSymbol -> String
showSymbol s = runFn1 showSymbolImpl s

foreign import showSymbolImpl :: Fn1 (JsSymbol) (String)

for :: String -> Effect JsSymbol
for s = runEffectFn1 forImpl s

foreign import forImpl :: EffectFn1 (String) (JsSymbol)

keyFor :: JsSymbol -> Effect (Maybe String)
keyFor s = map toMaybe $ runEffectFn1 keyForImpl s

foreign import keyForImpl :: EffectFn1 (JsSymbol) (Nullable String)
