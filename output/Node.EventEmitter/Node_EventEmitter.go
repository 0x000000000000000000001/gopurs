package Node_EventEmitter

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var eventNamesImpl = gopurs_runtime.Value{}
var getMaxListenersImpl = gopurs_runtime.Value{}
var listenerCountImpl = gopurs_runtime.Value{}
var new = gopurs_runtime.Value{}
var setMaxListenersImpl = gopurs_runtime.Value{}
var symbolOrStr = gopurs_runtime.Value{}
var unsafeEmitFn = gopurs_runtime.Value{}
var unsafeOff = gopurs_runtime.Value{}
var unsafeOn = gopurs_runtime.Value{}
var unsafeOnce = gopurs_runtime.Value{}
var unsafePrependListener = gopurs_runtime.Value{}
var unsafePrependOnceListener = gopurs_runtime.Value{}
var EventHandle = gopurs_runtime.Value{}
var subscribeSameFunction = gopurs_runtime.Apply(mkEffectFn4, gopurs_runtime.Value{})
var setMaxListeners = gopurs_runtime.Value{}
var setUnlimitedListeners = gopurs_runtime.Apply(setMaxListeners, gopurs_runtime.Value{})
var removeListenerH = gopurs_runtime.Value{}
var prependOnceListener_ = gopurs_runtime.Value{}
var prependOnceListener = gopurs_runtime.Value{}
var prependListener_ = gopurs_runtime.Value{}
var prependListener = gopurs_runtime.Value{}
var once_ = gopurs_runtime.Value{}
var once = gopurs_runtime.Value{}
var on_ = gopurs_runtime.Value{}
var on = gopurs_runtime.Value{}
var newListenerH = gopurs_runtime.Value{}
var listenerCount = gopurs_runtime.Value{}
var getMaxListeners = gopurs_runtime.Apply(runEffectFn1, getMaxListenersImpl)
var eventNames = gopurs_runtime.Value{}
