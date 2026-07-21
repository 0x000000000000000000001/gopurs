-- | ## Handling events emitted by an `EventEmitter`
-- |
-- | One can add callbacks to an `EventEmitter` on two major axes:
-- | - whether listener is added to the **end** (i.e. `on`) or **start** (i.e. `prependListener`) of the array
-- | - whether a listener is automatically removed after the first event (i.e. `once` or `prependOnceListener`).
-- |
-- | This module provides functions for each of the above 4 callback-adding functions
-- | If `<fn>` is either `on`, `once`, `prependListener`, or `prependOnceListener`, then this module exposes
-- | 1. `<fn>` - returns a callback that removes the listener
-- | 2. `<fn>_` - does not return a callback that can remove the listener
-- |
-- | ## Defining events emitted by an `EventEmitter`
-- |
-- | Below, we'll provide an example for how to define an event handler for a type. Let's assume the following:
-- | - There is a type `Foo` that exends `EventEmitter`
-- | - `Foo` values can handle "bar" events
-- | - a "bar" event takes the following callback: `EffectFn2 (Nullable Error) String Unit`
-- | - the `String` value is always either "red", "green", or "blue"
-- |
-- | Then we would write
-- | ```
-- | data Color 
-- |   = Red 
-- |   | Green 
-- |   | Blue
-- |
-- | -- Note: see docs on `EventHandle` 
-- | -- for the below naming convention justification 
-- | -- of suffixing an event name with `H`.
-- | barH 
-- |   :: EventHandle 
-- |        Foo 
-- |        (Maybe Error -> Color -> Effect Unit) 
-- |        (EffectFn1 (Nullable Error) String Unit)
-- | barH = EventHandle "bar" $ \psCb -> 
-- |   mkEffectFn2 \nullableError str ->
-- |     psCb (toMaybe nullableError) case str of
-- |       "red" -> Red
-- |       "green" -> Green
-- |       "blue" -> Blue
-- |       _ -> 
-- |         unsafeCrashWith $ 
-- |           "Impossible String value for event 'bar': " <> show str
-- | ```
-- |
-- | ## Emitting events via an `EventEmitter`
-- |
-- | Unfortunately, there isn't a good way to emit events safely in PureScript. If one wants to emit an event
-- | in PureScript code that will be consumed by PureScript code, there are better abstractions to use than `EventEmitter`.
-- | If one wants to emit an event in PureScript code that will be consumed by JavaScript code, then
-- | the `unsafeEmitFn` function can be used to call n-ary functions. However, this is very unsafe. See its docs for more context.
module Node.EventEmitter
  ( EventEmitter
  , new
  , SymbolOrStr
  , eventNames
  , getMaxListeners
  , listenerCount
  , setMaxListeners
  , setUnlimitedListeners
  , unsafeEmitFn
  , EventHandle(..)
  , newListenerH
  , removeListenerH
  , on
  , on_
  , once
  , once_
  , prependListener
  , prependListener_
  , prependOnceListener
  , prependOnceListener_
  ) where

import Prelude

import Data.Either (Either(..))
import Data.Function.Uncurried (Fn3, runFn3)
import Effect (Effect)
import Effect.Uncurried (EffectFn1, EffectFn2, EffectFn3, EffectFn4, mkEffectFn1, mkEffectFn4, runEffectFn1, runEffectFn2, runEffectFn3, runEffectFn4)
import Node.Symbol (JsSymbol)
import Unsafe.Coerce (unsafeCoerce)

foreign import data EventEmitter :: Type

-- | Create a new event emitter
foreign import new :: Effect EventEmitter

foreign import data SymbolOrStr :: Type

foreign import eventNamesImpl :: EventEmitter -> Array SymbolOrStr

eventNames :: EventEmitter -> Array (Either JsSymbol String)
eventNames ee = map (\x -> runFn3 symbolOrStr Left Right x) $ eventNamesImpl ee

foreign import symbolOrStr
  :: Fn3
       (forall a. JsSymbol -> Either JsSymbol a)
       (forall b. String -> Either b String)
       SymbolOrStr
       (Either JsSymbol String)

foreign import getMaxListenersImpl :: EffectFn1 EventEmitter Int

-- | By default, an event emitter can only have a maximum of 10 listeners
-- | for a given event.
getMaxListeners :: EventEmitter -> Effect Int
getMaxListeners = runEffectFn1 getMaxListenersImpl

foreign import listenerCountImpl :: EffectFn2 EventEmitter String Int

listenerCount :: EventEmitter -> String -> Effect Int
listenerCount emitter eventName = runEffectFn2 listenerCountImpl emitter eventName

foreign import setMaxListenersImpl :: EffectFn2 EventEmitter Int Unit

setMaxListeners :: Int -> EventEmitter -> Effect Unit
setMaxListeners max emitter = runEffectFn2 setMaxListenersImpl emitter max

setUnlimitedListeners :: EventEmitter -> Effect Unit
setUnlimitedListeners = setMaxListeners 0

-- | THIS IS UNSAFE! REALLY UNSAFE!
-- | Gets the `emit` function for a particular `EventEmitter`, so that one can call n-ary functions.
-- |
-- | Given `http2session.goaway([code[, lastStreamID[, opaqueData]]])` as an example...
-- | - https://nodejs.org/dist/latest-v18.x/docs/api/http2.html#event-goaway
-- | - https://nodejs.org/dist/latest-v18.x/docs/api/http2.html#http2sessiongoawaycode-laststreamid-opaquedata
-- |
-- | We can then write a single function that handles all four cases:
-- | ```
-- | goAway
-- |   :: Http2Session
-- |   -> Maybe Code
-- |   -> Maybe LastStreamId
-- |   -> Maybe OpaqueData
-- |   -> Effect Unit
-- | goAway h2s = case _, _, _ of
-- |   Just c, Just id, Just d ->
-- |     runEffectFn4 (unsafeEmitFn h2s :: EffectFn4 String Code LastStreamId OpaqueData Unit) "goaway" c id d
-- |   Just c, Just id, Nothing ->
-- |     -- If you're feeling lucky, omit the type annotations completely
-- |     runEffectFn3 (unsafeEmitFn h2s) "goaway" c id
-- |   Just c, Nothing, Nothing ->
-- |     runEffectFn2 (unsafeEmitFn h2s :: EffectFn2 String Code LastStreamId Unit) "goaway" c
-- |   _, _, _ ->
-- |     runEffectFn1 (unsafeEmitFn h2s :: EffectFn1 String Unit) "goaway"
-- | ```
-- | 
-- | Synchronously calls each of the listeners registered for the event named `eventName`, 
-- | in the order they were registered, passing the supplied arguments to each.
-- | Returns `true` if the event had listeners, `false` otherwise.
foreign import unsafeEmitFn :: forall f. EventEmitter -> f Boolean

-- | Packs all the type information we need to call `on`/`once`/`prependListener`/`prependOnceListener`
-- | with the correct callback function type.
-- |
-- | **Naming convention**: 
-- | If the name of an event is `foo`, 
-- | the corresponding PureScript `EventHandle` value should be called `fooH`.
-- | The `H` suffix is what prevent name conflicts in two situations:
-- | 1. similarly-named methods (e.g. the `"close"` event and the `close` method)
-- | 2. PureScript keywords (e.g. the `"data"` event)
-- |
-- | If an event, `foo`, can have two different kinds of callbacks, (e.g. See `Node.Stream`'s `data` event),
-- | one of two things should happen:
-- | 1. a suffix should follow the `H` to distinguish between the two (e.g. `dataHString`/`dataHBuffer`)
-- | 2. a prime character (i.e. `'`) should follow the `H` to distinguish between the two (e.g. `dataH`/`dataH'`)
data EventHandle :: Type -> Type -> Type -> Type
data EventHandle emitterType pureScriptCallback javaScriptCallback =
  EventHandle String (pureScriptCallback -> javaScriptCallback)

type role EventHandle representational representational representational

newListenerH :: EventHandle EventEmitter (Either JsSymbol String -> Effect Unit) (EffectFn1 SymbolOrStr Unit)
newListenerH = EventHandle "newListener" $ \cb -> mkEffectFn1 \jsSymbol ->
  cb $ runFn3 symbolOrStr Left Right jsSymbol

removeListenerH :: EventHandle EventEmitter (Either JsSymbol String -> Effect Unit) (EffectFn1 SymbolOrStr Unit)
removeListenerH = EventHandle "removeListener" $ \cb -> mkEffectFn1 \jsSymbol ->
  cb $ runFn3 symbolOrStr Left Right jsSymbol

-- | Adds the listener to the **end** of the `listeners` array.
-- | Returns a callback that will remove the listener from the event emitter's `listeners` array.
-- | If the listener removal callback isn't needed, use `on_`.
-- |
-- | Intended usage:
-- | ```
-- | removeLoggerCallback <- eventEmitter # on errorHandle \error -> do
-- |   log $ "Got error: " <> Exception.message error
-- |   log $ "This listener will now be removed."
-- | -- sometime later...
-- | removeLoggerCallback
-- | ```
on
  :: forall emitter psCb jsCb
   . EventHandle emitter psCb jsCb
  -> psCb
  -> emitter
  -> Effect (Effect Unit)
on (EventHandle eventName toJsCb) psCb eventEmitter =
  runEffectFn4 subscribeSameFunction unsafeOn (unsafeCoerce eventEmitter) eventName $ toJsCb psCb

-- | Adds the callback to the **end** of the `listeners` array and provides no way to remove the listener in the future.
-- | If you need a callback to remove the listener in the future, use `on`.
-- | Intended usage:
-- | ```
-- | eventEmitter # on_ errorHandle  \error -> do
-- |   log $ "Got error: " <> Exception.message error
-- | ```
on_
  :: forall emitter psCb jsCb
   . EventHandle emitter psCb jsCb
  -> psCb
  -> emitter
  -> Effect Unit
on_ (EventHandle eventName toJsCb) psCb eventEmitter =
  runEffectFn3 unsafeOn (unsafeCoerce eventEmitter) eventName $ toJsCb psCb

-- | Adds the listener to the **end** of the `listeners` array. The listener will be removed after it is invoked once.
-- | Returns a callback that will remove the listener from the event emitter's listeners array.
-- | If the listener removal callback isn't needed, use `once_`.
-- |
-- | Intended usage:
-- | ```
-- | removeLoggerCallback <- eventEmitter # once errorHandle \error -> do
-- |   log $ "Got error: " <> Exception.message error
-- |   log $ "This listener will now be removed."
-- | -- sometime later...
-- | removeLoggerCallback
-- | ```
once
  :: forall emitter psCb jsCb
   . EventHandle emitter psCb jsCb
  -> psCb
  -> emitter
  -> Effect (Effect Unit)
once (EventHandle eventName toJsCb) psCb eventEmitter =
  runEffectFn4 subscribeSameFunction unsafeOnce (unsafeCoerce eventEmitter) eventName $ toJsCb psCb

-- | Adds the listener to the **end** of the `listeners` array. The listener will be removed after it is invoked once.
-- | Returns a callback that will remove the listener from the event emitter's listeners array.
-- | If you need a callback to remove the listener in the future, use `once`.
-- |
-- | Intended usage:
-- | ```
-- | eventEmitter # once_ errorHandle \error -> do
-- |   log $ "Got error: " <> Exception.message error
-- | ```
once_
  :: forall emitter psCb jsCb
   . EventHandle emitter psCb jsCb
  -> psCb
  -> emitter
  -> Effect Unit
once_ (EventHandle eventName toJsCb) psCb eventEmitter =
  runEffectFn3 unsafeOnce (unsafeCoerce eventEmitter) eventName $ toJsCb psCb

-- | Adds the listener to the **start** of the `listeners` array.
-- | Returns a callback that will remove the listener from the event emitter's listeners array.
-- | If the listener removal callback isn't needed, use `prependListener_`.
-- |
-- | Intended usage:
-- | ```
-- | removeLoggerCallback <- eventEmitter # prependListener errorHandle \error -> do
-- |   log $ "Got error: " <> Exception.message error
-- |   log $ "This listener will now be removed."
-- | -- sometime later...
-- | removeLoggerCallback
-- | ```
prependListener
  :: forall emitter psCb jsCb
   . EventHandle emitter psCb jsCb
  -> psCb
  -> emitter
  -> Effect (Effect Unit)
prependListener (EventHandle eventName toJsCb) psCb eventEmitter =
  runEffectFn4 subscribeSameFunction unsafePrependListener (unsafeCoerce eventEmitter) eventName $ toJsCb psCb

-- | Adds the listener to the **start** of the `listeners` array.
-- | Returns a callback that will remove the listener from the event emitter's listeners array.
-- | If the listener removal callback isn't needed, use `prependListener`.
-- |
-- | Intended usage:
-- | ```
-- | eventEmitter # prependListener_ errorHandle \error -> do
-- |   log $ "Got error: " <> Exception.message error
-- | ```
prependListener_
  :: forall emitter psCb jsCb
   . EventHandle emitter psCb jsCb
  -> psCb
  -> emitter
  -> Effect Unit
prependListener_ (EventHandle eventName toJsCb) psCb eventEmitter =
  runEffectFn3 unsafePrependListener (unsafeCoerce eventEmitter) eventName $ toJsCb psCb

-- | Adds the listener to the **start** of the `listeners` array. The listener will be removed after it is invoked once.
-- | Returns a callback that will remove the listener from the event emitter's listeners array.
-- | If the listener removal callback isn't needed, use `prependOnceListener_`.
-- |
-- | Intended usage:
-- | ```
-- | removeLoggerCallback <- eventEmitter # prependOnceListener errorHandle \error -> do
-- |   log $ "Got error: " <> Exception.message error
-- |   log $ "This listener will now be removed."
-- | -- sometime later...
-- | removeLoggerCallback
-- | ```
prependOnceListener
  :: forall emitter psCb jsCb
   . EventHandle emitter psCb jsCb
  -> psCb
  -> emitter
  -> Effect (Effect Unit)
prependOnceListener (EventHandle eventName toJsCb) psCb eventEmitter =
  runEffectFn4 subscribeSameFunction unsafePrependOnceListener (unsafeCoerce eventEmitter) eventName $ toJsCb psCb

-- | Adds the listener to the **start** of the `listeners` array. The listener will be removed after it is invoked once.
-- | Returns a callback that will remove the listener from the event emitter's listeners array.
-- | If you need a callback to remove the listener in the future, use `prependOnceListener`.
-- |
-- | Intended usage:
-- | ```
-- | eventEmitter # prependOnceListener_ errorHandle \error -> do
-- |   log $ "Got error: " <> Exception.message error
-- | ```
prependOnceListener_
  :: forall emitter psCb jsCb
   . EventHandle emitter psCb jsCb
  -> psCb
  -> emitter
  -> Effect Unit
prependOnceListener_ (EventHandle eventName toJsCb) psCb eventEmitter =
  runEffectFn3 unsafePrependOnceListener (unsafeCoerce eventEmitter) eventName $ toJsCb psCb

-- | Internal function that ensures the JS callback function is the same one
-- | used when both adding it and removing it from the listeners array.
-- | Do not export this.
subscribeSameFunction
  :: forall emitter jsCb
   . EffectFn4
       (EffectFn3 emitter String jsCb Unit)
       emitter
       String
       jsCb
       (Effect Unit)
subscribeSameFunction = mkEffectFn4 \onXFn eventEmitter eventName jsCb -> do
  runEffectFn3 onXFn (unsafeCoerce eventEmitter) eventName jsCb
  pure $ runEffectFn3 unsafeOff (unsafeCoerce eventEmitter) eventName jsCb

foreign import unsafeOn :: forall f. EffectFn3 EventEmitter String f Unit
foreign import unsafeOff :: forall f. EffectFn3 EventEmitter String f Unit
foreign import unsafeOnce :: forall f. EffectFn3 EventEmitter String f Unit
foreign import unsafePrependListener :: forall f. EffectFn3 EventEmitter String f Unit
foreign import unsafePrependOnceListener :: forall f. EffectFn3 EventEmitter String f Unit
