// | ## Handling events emitted by an `EventEmitter`
// |
// | One can add callbacks to an `EventEmitter` on two major axes:
// | - whether listener is added to the **end** (i.e. `on`) or **start** (i.e. `prependListener`) of the array
// | - whether a listener is automatically removed after the first event (i.e. `once` or `prependOnceListener`).
// |
// | This module provides functions for each of the above 4 callback-adding functions
// | If `<fn>` is either `on`, `once`, `prependListener`, or `prependOnceListener`, then this module exposes
// | 1. `<fn>` - returns a callback that removes the listener
// | 2. `<fn>_` - does not return a callback that can remove the listener
// |
// | ## Defining events emitted by an `EventEmitter`
// |
// | Below, we'll provide an example for how to define an event handler for a type. Let's assume the following:
// | - There is a type `Foo` that exends `EventEmitter`
// | - `Foo` values can handle "bar" events
// | - a "bar" event takes the following callback: `EffectFn2 (Nullable Error) String Unit`
// | - the `String` value is always either "red", "green", or "blue"
// |
// | Then we would write
// | ```
// | data Color 
// |   = Red 
// |   | Green 
// |   | Blue
// |
// | -- Note: see docs on `EventHandle` 
// | -- for the below naming convention justification 
// | -- of suffixing an event name with `H`.
// | barH 
// |   :: EventHandle 
// |        Foo 
// |        (Maybe Error -> Color -> Effect Unit) 
// |        (EffectFn1 (Nullable Error) String Unit)
// | barH = EventHandle "bar" $ \psCb -> 
// |   mkEffectFn2 \nullableError str ->
// |     psCb (toMaybe nullableError) case str of
// |       "red" -> Red
// |       "green" -> Green
// |       "blue" -> Blue
// |       _ -> 
// |         unsafeCrashWith $ 
// |           "Impossible String value for event 'bar': " <> show str
// | ```
// |
// | ## Emitting events via an `EventEmitter`
// |
// | Unfortunately, there isn't a good way to emit events safely in PureScript. If one wants to emit an event
// | in PureScript code that will be consumed by PureScript code, there are better abstractions to use than `EventEmitter`.
// | If one wants to emit an event in PureScript code that will be consumed by JavaScript code, then
// | the `unsafeEmitFn` function can be used to call n-ary functions. However, this is very unsafe. See its docs for more context.
import * as $foreign from "./foreign.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Effect_Uncurried from "../Effect.Uncurried/index.js";
var EventHandle = /* #__PURE__ */ (function () {
    function EventHandle(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    EventHandle.create = function (value0) {
        return function (value1) {
            return new EventHandle(value0, value1);
        };
    };
    return EventHandle;
})();

// | Internal function that ensures the JS callback function is the same one
// | used when both adding it and removing it from the listeners array.
// | Do not export this.
var subscribeSameFunction = function (onXFn, eventEmitter, eventName, jsCb) {
    onXFn(eventEmitter, eventName, jsCb);
    return function () {
        return $foreign.unsafeOff(eventEmitter, eventName, jsCb);
    };
};
var setMaxListeners = function (max) {
    return function (emitter) {
        return function () {
            return $foreign.setMaxListenersImpl(emitter, max);
        };
    };
};
var setUnlimitedListeners = /* #__PURE__ */ setMaxListeners(0);
var removeListenerH = /* #__PURE__ */ (function () {
    return new EventHandle("removeListener", function (cb) {
        return function (jsSymbol) {
            return cb($foreign.symbolOrStr(Data_Either.Left.create, Data_Either.Right.create, jsSymbol))();
        };
    });
})();

// | Adds the listener to the **start** of the `listeners` array. The listener will be removed after it is invoked once.
// | Returns a callback that will remove the listener from the event emitter's listeners array.
// | If you need a callback to remove the listener in the future, use `prependOnceListener`.
// |
// | Intended usage:
// | ```
// | eventEmitter # prependOnceListener_ errorHandle \error -> do
// |   log $ "Got error: " <> Exception.message error
// | ```
var prependOnceListener_ = function (v) {
    return function (psCb) {
        return function (eventEmitter) {
            return function () {
                return $foreign.unsafePrependOnceListener(eventEmitter, v.value0, v.value1(psCb));
            };
        };
    };
};

// | Adds the listener to the **start** of the `listeners` array. The listener will be removed after it is invoked once.
// | Returns a callback that will remove the listener from the event emitter's listeners array.
// | If the listener removal callback isn't needed, use `prependOnceListener_`.
// |
// | Intended usage:
// | ```
// | removeLoggerCallback <- eventEmitter # prependOnceListener errorHandle \error -> do
// |   log $ "Got error: " <> Exception.message error
// |   log $ "This listener will now be removed."
// | -- sometime later...
// | removeLoggerCallback
// | ```
var prependOnceListener = function (v) {
    return function (psCb) {
        return function (eventEmitter) {
            return function () {
                return subscribeSameFunction($foreign.unsafePrependOnceListener, eventEmitter, v.value0, v.value1(psCb));
            };
        };
    };
};

// | Adds the listener to the **start** of the `listeners` array.
// | Returns a callback that will remove the listener from the event emitter's listeners array.
// | If the listener removal callback isn't needed, use `prependListener`.
// |
// | Intended usage:
// | ```
// | eventEmitter # prependListener_ errorHandle \error -> do
// |   log $ "Got error: " <> Exception.message error
// | ```
var prependListener_ = function (v) {
    return function (psCb) {
        return function (eventEmitter) {
            return function () {
                return $foreign.unsafePrependListener(eventEmitter, v.value0, v.value1(psCb));
            };
        };
    };
};

// | Adds the listener to the **start** of the `listeners` array.
// | Returns a callback that will remove the listener from the event emitter's listeners array.
// | If the listener removal callback isn't needed, use `prependListener_`.
// |
// | Intended usage:
// | ```
// | removeLoggerCallback <- eventEmitter # prependListener errorHandle \error -> do
// |   log $ "Got error: " <> Exception.message error
// |   log $ "This listener will now be removed."
// | -- sometime later...
// | removeLoggerCallback
// | ```
var prependListener = function (v) {
    return function (psCb) {
        return function (eventEmitter) {
            return function () {
                return subscribeSameFunction($foreign.unsafePrependListener, eventEmitter, v.value0, v.value1(psCb));
            };
        };
    };
};

// | Adds the listener to the **end** of the `listeners` array. The listener will be removed after it is invoked once.
// | Returns a callback that will remove the listener from the event emitter's listeners array.
// | If you need a callback to remove the listener in the future, use `once`.
// |
// | Intended usage:
// | ```
// | eventEmitter # once_ errorHandle \error -> do
// |   log $ "Got error: " <> Exception.message error
// | ```
var once_ = function (v) {
    return function (psCb) {
        return function (eventEmitter) {
            return function () {
                return $foreign.unsafeOnce(eventEmitter, v.value0, v.value1(psCb));
            };
        };
    };
};

// | Adds the listener to the **end** of the `listeners` array. The listener will be removed after it is invoked once.
// | Returns a callback that will remove the listener from the event emitter's listeners array.
// | If the listener removal callback isn't needed, use `once_`.
// |
// | Intended usage:
// | ```
// | removeLoggerCallback <- eventEmitter # once errorHandle \error -> do
// |   log $ "Got error: " <> Exception.message error
// |   log $ "This listener will now be removed."
// | -- sometime later...
// | removeLoggerCallback
// | ```
var once = function (v) {
    return function (psCb) {
        return function (eventEmitter) {
            return function () {
                return subscribeSameFunction($foreign.unsafeOnce, eventEmitter, v.value0, v.value1(psCb));
            };
        };
    };
};

// | Adds the callback to the **end** of the `listeners` array and provides no way to remove the listener in the future.
// | If you need a callback to remove the listener in the future, use `on`.
// | Intended usage:
// | ```
// | eventEmitter # on_ errorHandle  \error -> do
// |   log $ "Got error: " <> Exception.message error
// | ```
var on_ = function (v) {
    return function (psCb) {
        return function (eventEmitter) {
            return function () {
                return $foreign.unsafeOn(eventEmitter, v.value0, v.value1(psCb));
            };
        };
    };
};

// | Adds the listener to the **end** of the `listeners` array.
// | Returns a callback that will remove the listener from the event emitter's `listeners` array.
// | If the listener removal callback isn't needed, use `on_`.
// |
// | Intended usage:
// | ```
// | removeLoggerCallback <- eventEmitter # on errorHandle \error -> do
// |   log $ "Got error: " <> Exception.message error
// |   log $ "This listener will now be removed."
// | -- sometime later...
// | removeLoggerCallback
// | ```
var on = function (v) {
    return function (psCb) {
        return function (eventEmitter) {
            return function () {
                return subscribeSameFunction($foreign.unsafeOn, eventEmitter, v.value0, v.value1(psCb));
            };
        };
    };
};
var newListenerH = /* #__PURE__ */ (function () {
    return new EventHandle("newListener", function (cb) {
        return function (jsSymbol) {
            return cb($foreign.symbolOrStr(Data_Either.Left.create, Data_Either.Right.create, jsSymbol))();
        };
    });
})();
var listenerCount = function (emitter) {
    return function (eventName) {
        return function () {
            return $foreign.listenerCountImpl(emitter, eventName);
        };
    };
};

// | By default, an event emitter can only have a maximum of 10 listeners
// | for a given event.
var getMaxListeners = /* #__PURE__ */ Effect_Uncurried.runEffectFn1($foreign.getMaxListenersImpl);
var eventNames = function (ee) {
    return Data_Functor.map(Data_Functor.functorArray)(function (x) {
        return $foreign.symbolOrStr(Data_Either.Left.create, Data_Either.Right.create, x);
    })($foreign.eventNamesImpl(ee));
};
export {
    new,
    unsafeEmitFn
} from "./foreign.js";
export {
    eventNames,
    getMaxListeners,
    listenerCount,
    setMaxListeners,
    setUnlimitedListeners,
    EventHandle,
    newListenerH,
    removeListenerH,
    on,
    on_,
    once,
    once_,
    prependListener,
    prependListener_,
    prependOnceListener,
    prependOnceListener_
};
