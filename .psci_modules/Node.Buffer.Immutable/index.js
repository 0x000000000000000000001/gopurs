// | Immutable buffers and associated operations.
import * as $foreign from "./foreign.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Nullable from "../Data.Nullable/index.js";
import * as Data_Ordering from "../Data.Ordering/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Effect from "../Effect/index.js";
import * as Node_Buffer_Types from "../Node.Buffer.Types/index.js";
import * as Node_Encoding from "../Node.Encoding/index.js";
import * as Partial_Unsafe from "../Partial.Unsafe/index.js";
var toString$prime = function (enc) {
    return function (start) {
        return function (end) {
            return function (buf) {
                return $foreign.toStringSubImpl(enc, start, end, buf);
            };
        };
    };
};

// | Reads the buffer as a string with the specified encoding.
var toString = function (enc) {
    return function (buf) {
        return $foreign.toStringImpl(Node_Encoding.encodingToNode(enc), buf);
    };
};

// | Creates a new buffer slice that shares the memory of the original buffer.
var slice = function (start) {
    return function (end) {
        return function (buf) {
            return $foreign.sliceImpl(start, end, buf);
        };
    };
};
var showBuffer = {
    show: $foreign.showImpl
};

// | Reads a section of a buffer as a string with the specified encoding.
var readString = function (enc) {
    return function (start) {
        return function (end) {
            return function (buf) {
                return $foreign.readStringImpl(Node_Encoding.encodingToNode(enc), start, end, buf);
            };
        };
    };
};

// | Reads a numeric value from a buffer at the specified offset.
var read = function (ty) {
    return function (off) {
        return function (buf) {
            return $foreign.readImpl(Data_Show.show(Node_Buffer_Types.showBufferValueType)(ty), off, buf);
        };
    };
};

// | Reads an octet from a buffer at the specified offset.
var getAtOffset = function (off) {
    return function (buf) {
        return Data_Nullable.toMaybe($foreign.getAtOffsetImpl(off, buf));
    };
};

// | Creates a new buffer from a string with the specified encoding, sized to match the string.
var fromString = function (str) {
    return function (enc) {
        return $foreign.fromStringImpl(str, Node_Encoding.encodingToNode(enc));
    };
};
var eqBuffer = {
    eq: function (a) {
        return function (b) {
            return $foreign.eqImpl(a, b);
        };
    }
};
var ordBuffer = {
    compare: function (a) {
        return function (b) {
            var v = $foreign.compareImpl(a, b);
            if (v < 0) {
                return Data_Ordering.LT.value;
            };
            if (v > 0) {
                return Data_Ordering.GT.value;
            };
            return Data_Ordering.EQ.value;
        };
    },
    Eq0: function () {
        return eqBuffer;
    }
};

// | Creates a new buffer of the specified size. Alias for `alloc`.
var create = $foreign.alloc;

// | Concatenates a list of buffers, combining them into a new buffer of the
// | specified length.
var concat$prime = function (a) {
    return function (i) {
        return $foreign.concatToLength(a, i);
    };
};
var compareParts = function (src) {
    return function (target) {
        return function (targetStart) {
            return function (targetEnd) {
                return function (sourceStart) {
                    return function (sourceEnd) {
                        return Data_Functor.mapFlipped(Effect.functorEffect)(function () {
                            return $foreign.comparePartsImpl(src, target, targetStart, targetEnd, sourceStart, sourceEnd);
                        })(function (v) {
                            if (v === -1) {
                                return Data_Ordering.LT.value;
                            };
                            if (v === 0) {
                                return Data_Ordering.EQ.value;
                            };
                            if (v === 1) {
                                return Data_Ordering.GT.value;
                            };
                            return Partial_Unsafe.unsafeCrashWith("Impossible: Invalid value: " + Data_Show.show(Data_Show.showInt)(v));
                        });
                    };
                };
            };
        };
    };
};
export {
    alloc,
    fromArray,
    fromArrayBuffer,
    toArray,
    toArrayBuffer,
    concat,
    size
} from "./foreign.js";
export {
    compareParts,
    create,
    fromString,
    read,
    readString,
    toString,
    toString$prime,
    getAtOffset,
    concat$prime,
    slice,
    showBuffer,
    eqBuffer,
    ordBuffer
};
