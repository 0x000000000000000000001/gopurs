import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Monad_ST_Internal from "../Control.Monad.ST.Internal/index.js";
import * as Node_Buffer from "../Node.Buffer/index.js";
var writeString = Node_Buffer.writeString;
var write = Node_Buffer.write;
var unsafeThaw = Node_Buffer.unsafeThaw;
var unsafeFreeze = Node_Buffer.unsafeFreeze;
var transcode = Node_Buffer.transcode;
var toString$prime = Node_Buffer["toString$prime"];
var toString = Node_Buffer.toString;
var toArrayBuffer = Node_Buffer.toArrayBuffer;
var toArray = Node_Buffer.toArray;
var thaw = Node_Buffer.thaw;
var swap64 = Node_Buffer.swap64;
var swap32 = Node_Buffer.swap32;
var swap16 = Node_Buffer.swap16;
var slice = Node_Buffer.slice;
var size = Node_Buffer.size;
var setAtOffset = Node_Buffer.setAtOffset;

// | Runs an ST creating an `STBuffer` then freezes the (STBuffer h) and returns
// | it, without unneccessary copying.
var run = function (st) {
    return Control_Bind.bind(Control_Monad_ST_Internal.bindST)(st)(unsafeFreeze)();
};
var readString = Node_Buffer.readString;
var read = Node_Buffer.read;
var getAtOffset = Node_Buffer.getAtOffset;
var fromString = Node_Buffer.fromString;
var fromArrayBuffer = Node_Buffer.fromArrayBuffer;
var fromArray = Node_Buffer.fromArray;
var freeze = Node_Buffer.freeze;
var fill = Node_Buffer.fill;
var copy = Node_Buffer.copy;
var concat$prime = Node_Buffer["concat$prime"];
var concat = Node_Buffer.concat;
var compareParts = Node_Buffer.compareParts;

// | Creates a new `STBuffer` of the specified size. Unsafe because it reuses memory from a pool
// | and may contain sensitive data. See the Node docs: 
// | https://nodejs.org/docs/latest-v18.x/api/buffer.html#what-makes-bufferallocunsafe-and-bufferallocunsafeslow-unsafe
var allocUnsafeSlow = Node_Buffer.allocUnsafeSlow;

// | Creates a new `STBuffer` of the specified size. Unsafe because it reuses memory from a pool
// | and may contain sensitive data. See the Node docs: 
// | https://nodejs.org/docs/latest-v18.x/api/buffer.html#what-makes-bufferallocunsafe-and-bufferallocunsafeslow-unsafe
var allocUnsafe = Node_Buffer.allocUnsafe;

// | Creates a new `STBuffer`.
var alloc = Node_Buffer.alloc;

// | Creates a new `STBuffer`. Alias of `alloc`.
var create = alloc;
var mutableBufferST = {
    create: create,
    freeze: freeze,
    unsafeFreeze: unsafeFreeze,
    thaw: thaw,
    unsafeThaw: unsafeThaw,
    fromArray: fromArray,
    fromString: fromString,
    fromArrayBuffer: fromArrayBuffer,
    toArrayBuffer: toArrayBuffer,
    read: read,
    readString: readString,
    toString: toString,
    write: write,
    writeString: writeString,
    toArray: toArray,
    getAtOffset: getAtOffset,
    setAtOffset: setAtOffset,
    slice: slice,
    size: size,
    concat: concat,
    "concat'": concat$prime,
    copy: copy,
    fill: fill,
    Monad0: function () {
        return Control_Monad_ST_Internal.monadST;
    }
};
export {
    run,
    create,
    alloc,
    allocUnsafe,
    allocUnsafeSlow,
    compareParts,
    freeze,
    unsafeFreeze,
    thaw,
    unsafeThaw,
    fromArray,
    fromString,
    fromArrayBuffer,
    toArrayBuffer,
    read,
    readString,
    toString,
    toString$prime,
    write,
    writeString,
    toArray,
    getAtOffset,
    setAtOffset,
    slice,
    size,
    concat,
    concat$prime,
    copy,
    fill,
    swap16,
    swap32,
    swap64,
    transcode,
    mutableBufferST
};
