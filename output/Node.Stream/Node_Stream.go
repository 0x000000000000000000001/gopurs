package Node_Stream

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var writeableNeedDrain = gopurs_runtime.Value{}
var writeableLength = gopurs_runtime.Value{}
var writeableHighWaterMark = gopurs_runtime.Value{}
var writeableFinished = gopurs_runtime.Value{}
var writeableEnded = gopurs_runtime.Value{}
var writeableCorked = gopurs_runtime.Value{}
var writeable = gopurs_runtime.Value{}
var writeString' = gopurs_runtime.Value{}
var writeString = gopurs_runtime.Value{}
var write' = gopurs_runtime.Value{}
var write = gopurs_runtime.Value{}
var unpipeH = gopurs_runtime.Value{}
var unpipeAll = gopurs_runtime.Value{}
var unpipe = gopurs_runtime.Value{}
var uncork = gopurs_runtime.Value{}
var toEventEmitter = unsafeCoerce
var setEncoding = gopurs_runtime.Value{}
var setDefaultEncoding = gopurs_runtime.Value{}
var resumeH = gopurs_runtime.Value{}
var resume = gopurs_runtime.Value{}
var readableLength = gopurs_runtime.Value{}
var readableHighWaterMark = gopurs_runtime.Value{}
var readableH = gopurs_runtime.Value{}
var readableFromString = gopurs_runtime.Value{}
var readableFromBuffer = gopurs_runtime.Value{}
var readableFlowing = gopurs_runtime.Value{}
var readableEnded = gopurs_runtime.Value{}
var readable = gopurs_runtime.Value{}
var readEither' = gopurs_runtime.Value{}
var readEither = gopurs_runtime.Value{}
var read' = gopurs_runtime.Value{}
var readString' = gopurs_runtime.Value{}
var read = gopurs_runtime.Value{}
var readString = gopurs_runtime.Value{}
var pipeline = gopurs_runtime.Value{}
var pipeH = gopurs_runtime.Value{}
var pipe' = gopurs_runtime.Value{}
var pipe = gopurs_runtime.Value{}
var pauseH = gopurs_runtime.Value{}
var pause = gopurs_runtime.Value{}
var isPaused = gopurs_runtime.Value{}
var finishH = gopurs_runtime.Value{}
var errored = gopurs_runtime.Value{}
var errorH = gopurs_runtime.Value{}
var endH = gopurs_runtime.Value{}
var end' = gopurs_runtime.Value{}
var end = gopurs_runtime.Value{}
var drainH = gopurs_runtime.Value{}
var destroyed = gopurs_runtime.Value{}
var destroy' = gopurs_runtime.Value{}
var destroy = gopurs_runtime.Value{}
var dataHStr = gopurs_runtime.Value{}
var dataHEither = gopurs_runtime.Value{}
var dataH = gopurs_runtime.Value{}
var cork = gopurs_runtime.Value{}
var closed = gopurs_runtime.Value{}
var closeH = gopurs_runtime.Value{}
var allowHalfOpen = gopurs_runtime.Value{}