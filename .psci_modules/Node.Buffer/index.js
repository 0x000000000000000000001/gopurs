// | Mutable buffers and associated operations.
import * as $foreign from "./foreign.js";
import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Effect from "../Effect/index.js";
import * as Effect_Uncurried from "../Effect.Uncurried/index.js";
import * as Node_Buffer_Immutable from "../Node.Buffer.Immutable/index.js";
import * as Node_Buffer_Types from "../Node.Buffer.Types/index.js";
import * as Node_Encoding from "../Node.Encoding/index.js";
var writeString = function (enc) {
    return function (offset) {
        return function (len) {
            return function (value) {
                return function (buf) {
                    return function () {
                        return $foreign.writeStringInternal(Node_Encoding.encodingToNode(enc), offset, len, value, buf);
                    };
                };
            };
        };
    };
};
var write = function (ty) {
    return function (value) {
        return function (offset) {
            return function (buf) {
                return function () {
                    return $foreign.writeInternal(Data_Show.show(Node_Buffer_Types.showBufferValueType)(ty), value, offset, buf);
                };
            };
        };
    };
};
var unsafeThaw = /* #__PURE__ */ (function () {
    var $4 = Control_Applicative.pure(Effect.applicativeEffect);
    return function ($5) {
        return $4($5);
    };
})();
var usingToImmutable = function (f) {
    return function (x) {
        return unsafeThaw(f(x));
    };
};
var unsafeFreeze = /* #__PURE__ */ (function () {
    var $6 = Control_Applicative.pure(Effect.applicativeEffect);
    return function ($7) {
        return $6($7);
    };
})();
var usingFromImmutable = function (f) {
    return function (buf) {
        return Data_Functor.map(Effect.functorEffect)(f)(unsafeFreeze(buf));
    };
};
var transcode = function (buf) {
    return function (from) {
        return function (to) {
            return function () {
                return $foreign.transcodeImpl(buf, Node_Encoding.encodingToNode(from), Node_Encoding.encodingToNode(to));
            };
        };
    };
};
var toString$prime = function (enc) {
    return function (start) {
        return function (end) {
            return usingFromImmutable(Node_Buffer_Immutable["toString$prime"](enc)(start)(end));
        };
    };
};
var toString = function (enc) {
    return usingFromImmutable(Node_Buffer_Immutable.toString(enc));
};
var toArrayBuffer = /* #__PURE__ */ usingFromImmutable(Node_Buffer_Immutable.toArrayBuffer);
var toArray = /* #__PURE__ */ usingFromImmutable(Node_Buffer_Immutable.toArray);
var thaw = /* #__PURE__ */ Effect_Uncurried.runEffectFn1($foreign.thawImpl);
var swap64 = function (b) {
    return function () {
        return $foreign.swap64Impl(b);
    };
};
var swap32 = function (b) {
    return function () {
        return $foreign.swap32Impl(b);
    };
};
var swap16 = function (b) {
    return function () {
        return $foreign.swap16Impl(b);
    };
};
var slice = Node_Buffer_Immutable.slice;
var size = /* #__PURE__ */ usingFromImmutable(Node_Buffer_Immutable.size);
var setPoolSize = function (sizeInBytes) {
    return function () {
        return $foreign.setPoolSizeImpl(sizeInBytes);
    };
};
var setAtOffset = function (val) {
    return function (off) {
        return function (buff) {
            return function () {
                return $foreign.setAtOffsetImpl(val, off, buff);
            };
        };
    };
};
var readString = function (enc) {
    return function (o) {
        return function (o$prime) {
            return usingFromImmutable(Node_Buffer_Immutable.readString(enc)(o)(o$prime));
        };
    };
};
var read = function (t) {
    return function (o) {
        return usingFromImmutable(Node_Buffer_Immutable.read(t)(o));
    };
};
var getAtOffset = function (o) {
    return usingFromImmutable(Node_Buffer_Immutable.getAtOffset(o));
};
var fromString = function (s) {
    return usingToImmutable(Node_Buffer_Immutable.fromString(s));
};
var fromArrayBuffer = /* #__PURE__ */ usingToImmutable(Node_Buffer_Immutable.fromArrayBuffer);
var fromArray = /* #__PURE__ */ usingToImmutable(Node_Buffer_Immutable.fromArray);
var freeze = /* #__PURE__ */ Effect_Uncurried.runEffectFn1($foreign.freezeImpl);
var fill = function (octet) {
    return function (start) {
        return function (end) {
            return function (buf) {
                return function () {
                    return $foreign.fillImpl(octet, start, end, buf);
                };
            };
        };
    };
};
var copy = function (srcStart) {
    return function (srcEnd) {
        return function (src) {
            return function (targStart) {
                return function (targ) {
                    return function () {
                        return $foreign.copyImpl(srcStart, srcEnd, src, targStart, targ);
                    };
                };
            };
        };
    };
};
var concat$prime = function (arrs) {
    return function (n) {
        return function (v) {
            return Node_Buffer_Immutable["concat$prime"](arrs)(n);
        };
    };
};
var concat = function (arrs) {
    return function (v) {
        return Node_Buffer_Immutable.concat(arrs);
    };
};
var compareParts = function (src) {
    return function (target) {
        return function (targetSrc) {
            return function (targetEnd) {
                return function (srcStart) {
                    return function (srcEnd) {
                        return function __do() {
                            var src$prime = unsafeFreeze(src)();
                            var target$prime = unsafeFreeze(target)();
                            return Node_Buffer_Immutable.compareParts(src$prime)(target$prime)(targetSrc)(targetEnd)(srcStart)(srcEnd)();
                        };
                    };
                };
            };
        };
    };
};

// | Creates a new buffer of the specified size. Unsafe because it reuses memory from a pool
// | and may contain sensitive data. See the Node docs: 
// | https://nodejs.org/docs/latest-v18.x/api/buffer.html#what-makes-bufferallocunsafe-and-bufferallocunsafeslow-unsafe
var allocUnsafeSlow = function (s) {
    return function () {
        return $foreign.allocUnsafeSlowImpl(s);
    };
};

// | Creates a new buffer of the specified size. Unsafe because it reuses memory from a pool
// | and may contain sensitive data. See the Node docs: 
// | https://nodejs.org/docs/latest-v18.x/api/buffer.html#what-makes-bufferallocunsafe-and-bufferallocunsafeslow-unsafe
var allocUnsafe = function (s) {
    return function () {
        return $foreign.allocUnsafeImpl(s);
    };
};

// | Creates a new buffer of the specified size.
var alloc = /* #__PURE__ */ usingToImmutable(Node_Buffer_Immutable.alloc);

// | Creates a new buffer of the specified size. Alias to `alloc`.
var create = alloc;
var mutableBufferEffect = {
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
        return Effect.monadEffect;
    }
};
export {
    poolSize
} from "./foreign.js";
export {
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
    setPoolSize,
    swap16,
    swap32,
    swap64,
    transcode,
    mutableBufferEffect
};
export {
    DoubleBE,
    DoubleLE,
    FloatBE,
    FloatLE,
    Int16BE,
    Int16LE,
    Int32BE,
    Int32LE,
    Int8,
    UInt16BE,
    UInt16LE,
    UInt32BE,
    UInt32LE,
    UInt8
} from "../Node.Buffer.Types/index.js";
