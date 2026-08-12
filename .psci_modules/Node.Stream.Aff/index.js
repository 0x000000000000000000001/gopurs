// | Asynchronous I/O with the [*Node.js* Stream API](https://nodejs.org/docs/latest/api/stream.html).
// |
// | Open __file streams__ with
// | [__Node.FS.Stream__](https://pursuit.purescript.org/packages/purescript-node-fs/docs/Node.FS.Stream).
// |
// | Open __process streams__ with
// | [__Node.Process__](https://pursuit.purescript.org/packages/purescript-node-process/docs/Node.Process).
// |
// | All __I/O errors__ will be thrown through the `Aff` `MonadError` class
// | instance.
// |
// | `Aff` __cancellation__ will clean up all *Node.js* event listeners.
// |
// | All of these `Aff` functions will prevent the *Node.js* __event loop__ from
// | exiting until the `Aff` function completes.
// |
// | ## Reading
// |
// | #### Implementation
// |
// | The `read*` functions (not to be confused with the `readable*` functions)
// | in this module all operate on a `Readable` stream
// | in
// | [“paused mode”](https://nodejs.org/docs/latest/api/stream.html#stream_two_reading_modes).
// |
// | Internally the reading functions use the
// | [`readable.read([size])`](https://nodejs.org/docs/latest/api/stream.html#readablereadsize)
// | function and are subject to the caveats of that function.
// |
// | #### Result Buffers
// |
// | The result of a reading function may be chunked into more than one `Buffer`.
// | The `buffers` element of the result is an `Array Buffer` of what
// | was read.
// | To concatenate the result into a single `Buffer`, use
// | [`Node.Buffer.concat :: Array Buffer -> m Buffer`](https://pursuit.purescript.org/packages/purescript-node-buffer/docs/Node.Buffer#t:MutableBuffer).
// |
// | ```
// | input :: Buffer
// |     <- liftEffect <<< concat <<< _.buffers =<< readSome stdin
// | ```
// |
// | To calculate the number of bytes read, use
// | `Node.Buffer.size :: Buffer -> m Int`.
// |
// | ```
// | { buffers } :: Array Buffer <- readSome stdin
// | bytesRead :: Int
// |     <- liftEffect $ Array.foldM (\a b -> (a+_) <$> size b) 0 buffers
// | ```
// |
// | #### Result `readagain` flag
// |
// | The `readagain` field of the result is a `Boolean` flag which
// | is `true` if the stream has not reached End-Of-File (and also if the stream
// | has not errored or been destroyed), so we know we can read again.
// | If the flag is `false` then the stream is not `readable`;
// | no more bytes will ever be produced by the stream.
// |
// | Reading from an ended, closed, errored, or destroyed stream
// | will complete immediately with `{ buffers: [], readagain: false }`.
// |
// | The `readagain` flag will give the same answer as a
// | subsequent call to `readable`.
// |
// | ## Writing
// |
// | #### Implementation
// |
// | The writing functions in this module all operate on a `Writeable` stream.
// |
// | Internally the writing functions will call the
// | [`writable.write(chunk[, encoding][, callback])`](https://nodejs.org/docs/latest/api/stream.html#writablewritechunk-encoding-callback)
// | function on each of the `Buffer`s,
// | and will asychronously wait if there is “backpressure” from the stream.
// |
// | #### Result
// |
// | The writing functions will complete after all the data is flushed to the
// | stream.
// |
// | If a write fails then it will `throwError` in the `Aff`.
import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Monad_Rec_Class from "../Control.Monad.Rec.Class/index.js";
import * as Control_Monad_ST_Class from "../Control.Monad.ST.Class/index.js";
import * as Data_Array from "../Data.Array/index.js";
import * as Data_Array_ST from "../Data.Array.ST/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Unit from "../Data.Unit/index.js";
import * as Effect from "../Effect/index.js";
import * as Effect_Aff from "../Effect.Aff/index.js";
import * as Effect_Aff_Class from "../Effect.Aff.Class/index.js";
import * as Effect_Class from "../Effect.Class/index.js";
import * as Effect_Exception from "../Effect.Exception/index.js";
import * as Effect_Ref from "../Effect.Ref/index.js";
import * as Node_Buffer from "../Node.Buffer/index.js";
import * as Node_Encoding from "../Node.Encoding/index.js";
import * as Node_EventEmitter from "../Node.EventEmitter/index.js";
import * as Node_Stream from "../Node.Stream/index.js";
var $runtime_lazy = function (name, moduleName, init) {
    var state = 0;
    var val;
    return function (lineNumber) {
        if (state === 2) return val;
        if (state === 1) throw new ReferenceError(name + " was needed before it finished initializing (module " + moduleName + ", line " + lineNumber + ")", moduleName, lineNumber);
        state = 1;
        val = init();
        state = 2;
        return val;
    };
};
var tailRecM = /* #__PURE__ */ Control_Monad_Rec_Class.tailRecM(Control_Monad_Rec_Class.monadRecEffect);

// canceller might be called while waiting for `onceDrain`
var pure = /* #__PURE__ */ Control_Applicative.pure(Effect.applicativeEffect);
var join = /* #__PURE__ */ Control_Bind.join(Effect.bindEffect);
var liftST = /* #__PURE__ */ Control_Monad_ST_Class.liftST(Control_Monad_ST_Class.monadSTEffect);
var mempty = /* #__PURE__ */ Data_Monoid.mempty(/* #__PURE__ */ Effect.monoidEffect(Data_Monoid.monoidUnit));
var liftST1 = /* #__PURE__ */ Control_Monad_ST_Class.liftST(Control_Monad_ST_Class.monadSTEffect);
var $$void = /* #__PURE__ */ Data_Functor["void"](Effect.functorEffect);
var liftST2 = /* #__PURE__ */ Control_Monad_ST_Class.liftST(Control_Monad_ST_Class.monadSTEffect);
var pure1 = /* #__PURE__ */ Control_Applicative.pure(Control_Applicative.applicativeArray);

// | Write to a stream.
// |
// | Will complete after the data is flushed to the stream.
var write = function (dictMonadAff) {
    var liftAff = Effect_Aff_Class.liftAff(dictMonadAff);
    return function (w) {
        return function (bs) {
            return liftAff(Effect_Aff.makeAff(function (complete) {
                return function __do() {
                    var removeDrain = Effect_Ref["new"](Control_Applicative.pure(Effect.applicativeEffect)(Data_Unit.unit))();
                    var oneWrite = function (i$prime) {
                        return Data_Function.flip(tailRecM)(i$prime)(function (i) {
                            var v = Data_Array.index(bs)(i);
                            if (v instanceof Data_Maybe.Nothing) {
                                return function __do() {
                                    complete(new Data_Either.Right(Data_Unit.unit))();
                                    return new Control_Monad_Rec_Class.Done(Data_Unit.unit);
                                };
                            };
                            if (v instanceof Data_Maybe.Just) {
                                return function __do() {
                                    var nobackpressure = Node_Stream["write$prime"](w)(v.value0)(function (v1) {
                                        if (v1 instanceof Data_Maybe.Nothing) {
                                            return Control_Applicative.pure(Effect.applicativeEffect)(Data_Unit.unit);
                                        };
                                        if (v1 instanceof Data_Maybe.Just) {
                                            return complete(new Data_Either.Left(v1.value0));
                                        };
                                        throw new Error("Failed pattern match at Node.Stream.Aff (line 465, column 49 - line 469, column 34): " + [ v1.constructor.name ]);
                                    })();
                                    if (nobackpressure) {
                                        return new Control_Monad_Rec_Class.Loop(i + 1 | 0);
                                    };
                                    var removeDrain$prime = Node_EventEmitter.once(Node_Stream.drainH)(oneWrite(i + 1 | 0))(w)();
                                    Effect_Ref.write(removeDrain$prime)(removeDrain)();
                                    return new Control_Monad_Rec_Class.Done(Data_Unit.unit);
                                };
                            };
                            throw new Error("Failed pattern match at Node.Stream.Aff (line 456, column 7 - line 476, column 29): " + [ v.constructor.name ]);
                        });
                    };
                    oneWrite(0)();
                    return pure(Effect_Aff.effectCanceler(join(Effect_Ref.read(removeDrain))))();
                };
            }));
        };
    };
};

// | Concatenate an `Array` of UTF-8 encoded `Buffer`s into a `String`.
// |
// | Example:
// |
// | ```
// | inputstring <- toStringUTF8 =<< readAll stream
// | ```
var toStringUTF8 = function (dictMonadEffect) {
    var liftEffect = Effect_Class.liftEffect(dictMonadEffect);
    return function (bs) {
        return liftEffect(Control_Bind.bindFlipped(Effect.bindEffect)(Node_Buffer.toString(Node_Encoding.UTF8.value))(Node_Buffer.concat(bs)));
    };
};

// | Works on streams in "flowing" mode.
// | Reads all of the stream's buffered contents into an array.
var readableToBuffers = function (dictMonadAff) {
    var liftAff = Effect_Aff_Class.liftAff(dictMonadAff);
    return function (r) {
        return liftAff(Effect_Aff.makeAff(function (complete) {
            return function __do() {
                var bufs = liftST(Data_Array_ST["new"])();
                var dataRef = Effect_Ref["new"](mempty)();
                var removeData = join(Effect_Ref.read(dataRef));
                var removeError = Node_EventEmitter.once(Node_Stream.errorH)(function (err) {
                    return function __do() {
                        removeData();
                        return complete(new Data_Either.Left(err))();
                    };
                })(r)();
                var removeClose = Node_EventEmitter.once(Node_Stream.closeH)(function __do() {
                    removeError();
                    removeData();
                    var result = liftST1(Data_Array_ST.unsafeFreeze(bufs))();
                    return complete(new Data_Either.Right(result))();
                })(r)();
                var removeEnd = Node_EventEmitter.once(Node_Stream.endH)(function __do() {
                    removeError();
                    removeClose();
                    removeData();
                    var result = liftST1(Data_Array_ST.unsafeFreeze(bufs))();
                    return complete(new Data_Either.Right(result))();
                })(r)();
                var rmData = Node_EventEmitter.on(Node_Stream.dataH)(function (buf) {
                    return $$void(liftST2(Data_Array_ST.push(buf)(bufs)));
                })(r)();
                Effect_Ref.write(rmData)(dataRef)();
                return pure(Effect_Aff.effectCanceler(function __do() {
                    removeError();
                    removeClose();
                    removeEnd();
                    return removeData();
                }))();
            };
        }));
    };
};

// | Works on streams in "flowing" mode. 
// | Reads all of the stream's contents into a buffer
// | and converts the result into a String using the provided encoding.
var readableToString = function (dictMonadAff) {
    var MonadEffect0 = dictMonadAff.MonadEffect0();
    var Bind1 = (MonadEffect0.Monad0()).Bind1();
    var liftEffect = Effect_Class.liftEffect(MonadEffect0);
    return function (r) {
        return function (enc) {
            return Control_Bind.bind(Bind1)(readableToBuffers(dictMonadAff)(r))(function (bufs) {
                return liftEffect(Control_Bind.bindFlipped(Effect.bindEffect)(Node_Buffer.toString(enc))(Node_Buffer.concat(bufs)));
            });
        };
    };
};

// | Works on streams in "flowing" mode. 
// | Reads all of the stream's contents into a buffer
// | and converts the result into a UTF8-encoded String.
var readableToStringUtf8 = function (dictMonadAff) {
    return function (r) {
        return readableToString(dictMonadAff)(r)(Node_Encoding.UTF8.value);
    };
};

// | Works on streams in "paused" mode. 
// | Wait until there is some data available from the stream, then read it.
// |
// | This function is useful for streams like __stdin__ which never
// | reach End-Of-File.
var readSome = function (dictMonadAff) {
    var liftAff = Effect_Aff_Class.liftAff(dictMonadAff);
    return function (r) {
        return liftAff(Effect_Aff.makeAff(function (complete) {
            return function __do() {
                var isReadable = Node_Stream.readable(r)();
                if (isReadable) {
                    var bufs = liftST(Data_Array_ST["new"])();
                    var removeError = Node_EventEmitter.once(Node_Stream.errorH)(function (err) {
                        return complete(new Data_Either.Left(err));
                    })(r)();
                    var removeClose = Node_EventEmitter.once(Node_Stream.closeH)(function __do() {
                        removeError();
                        var ret = liftST1(Data_Array_ST.unsafeFreeze(bufs))();
                        return complete(new Data_Either.Right({
                            buffers: ret,
                            readagain: false
                        }))();
                    })(r)();
                    var removeEnd = Node_EventEmitter.once(Node_Stream.endH)(function __do() {
                        removeError();
                        removeClose();
                        var ret = liftST1(Data_Array_ST.unsafeFreeze(bufs))();
                        return complete(new Data_Either.Right({
                            buffers: ret,
                            readagain: false
                        }))();
                    })(r)();
                    var cleanupRethrow = function (err) {
                        return function __do() {
                            removeError();
                            removeClose();
                            removeEnd();
                            complete(new Data_Either.Left(err))();
                            return Effect_Aff.nonCanceler;
                        };
                    };
                    return Effect_Exception.catchException(cleanupRethrow)(function __do() {
                        (function () {
                            while (!(function __do() {
                                var v = Node_Stream.read(r)();
                                if (v instanceof Data_Maybe.Nothing) {
                                    return true;
                                };
                                if (v instanceof Data_Maybe.Just) {
                                    $$void(liftST2(Data_Array_ST.push(v.value0)(bufs)))();
                                    return false;
                                };
                                throw new Error("Failed pattern match at Node.Stream.Aff (line 221, column 27 - line 225, column 23): " + [ v.constructor.name ]);
                            })()) {

                            };
                            return {};
                        })();
                        var ret1 = liftST1(Data_Array_ST.unsafeFreeze(bufs))();
                        var readagain = Node_Stream.readable(r)();
                        var $50 = readagain && Data_Array.length(ret1) === 0;
                        if ($50) {
                            var removeReadable = Node_EventEmitter.once(Node_Stream.readableH)(function __do() {
                                (function () {
                                    while (!(function __do() {
                                        var v = Node_Stream.read(r)();
                                        if (v instanceof Data_Maybe.Nothing) {
                                            return true;
                                        };
                                        if (v instanceof Data_Maybe.Just) {
                                            $$void(liftST2(Data_Array_ST.push(v.value0)(bufs)))();
                                            return false;
                                        };
                                        throw new Error("Failed pattern match at Node.Stream.Aff (line 238, column 31 - line 242, column 27): " + [ v.constructor.name ]);
                                    })()) {

                                    };
                                    return {};
                                })();
                                var ret2 = liftST1(Data_Array_ST.unsafeFreeze(bufs))();
                                removeError();
                                removeClose();
                                removeEnd();
                                var readagain2 = Node_Stream.readable(r)();
                                return complete(new Data_Either.Right({
                                    buffers: ret2,
                                    readagain: readagain2
                                }))();
                            })(r)();
                            return pure(Effect_Aff.effectCanceler(function __do() {
                                removeError();
                                removeClose();
                                removeEnd();
                                return removeReadable();
                            }))();
                        };
                        removeError();
                        removeClose();
                        removeEnd();
                        complete(new Data_Either.Right({
                            buffers: ret1,
                            readagain: readagain
                        }))();
                        return Effect_Aff.nonCanceler;
                    })();
                };
                complete(new Data_Either.Right({
                    buffers: [  ],
                    readagain: false
                }))();
                return Effect_Aff.nonCanceler;
            };
        }));
    };
};

// | Works on streams in "paused" mode. 
// | Wait for *N* bytes to become available from the stream.
// |
// | If more than *N* bytes are available on the stream, then
// | completes with *N* bytes and leaves the rest in the stream’s internal buffer.
// |
// | If the end of the stream is reached before *N* bytes are available,
// | then completes with less than *N* bytes.
var readN = function (dictMonadAff) {
    var liftAff = Effect_Aff_Class.liftAff(dictMonadAff);
    return function (r) {
        return function (n) {
            return liftAff(Effect_Aff.makeAff(function (complete) {
                var $53 = n < 0;
                if ($53) {
                    return Control_Apply.applySecond(Effect.applyEffect)(complete(new Data_Either.Left(Effect_Exception.error("read bytes must be > 0"))))(Control_Applicative.pure(Effect.applicativeEffect)(Effect_Aff.nonCanceler));
                };
                return function __do() {
                    var isReadable = Node_Stream.readable(r)();
                    if (isReadable) {
                        var redRef = Effect_Ref["new"](0)();
                        var bufs = liftST(Data_Array_ST["new"])();
                        var removeReadable = Effect_Ref["new"](Control_Applicative.pure(Effect.applicativeEffect)(Data_Unit.unit))();
                        var removeError = Node_EventEmitter.once(Node_Stream.errorH)(function (err) {
                            return function __do() {
                                join(Effect_Ref.read(removeReadable))();
                                return complete(new Data_Either.Left(err))();
                            };
                        })(r)();
                        var removeClose = Node_EventEmitter.once(Node_Stream.closeH)(function __do() {
                            removeError();
                            join(Effect_Ref.read(removeReadable))();
                            var ret = liftST1(Data_Array_ST.unsafeFreeze(bufs))();
                            return complete(new Data_Either.Right({
                                buffers: ret,
                                readagain: false
                            }))();
                        })(r)();
                        var removeEnd = Node_EventEmitter.once(Node_Stream.endH)(function __do() {
                            removeError();
                            removeClose();
                            var ret = liftST1(Data_Array_ST.unsafeFreeze(bufs))();
                            return complete(new Data_Either.Right({
                                buffers: ret,
                                readagain: false
                            }))();
                        })(r)();
                        
                        // try to read N bytes and then either return N bytes or run a continuation
var tryToRead = function (continuation) {
                            return function __do() {
                                (function () {
                                    while (!(function __do() {
                                        var red = Effect_Ref.read(redRef)();
                                        var v = Node_Stream["read$prime"](r)(n - red | 0)();
                                        if (v instanceof Data_Maybe.Nothing) {
                                            return true;
                                        };
                                        if (v instanceof Data_Maybe.Just) {
                                            liftST2(Data_Array_ST.push(v.value0)(bufs))();
                                            var s = Node_Buffer.size(v.value0)();
                                            var red$prime = Effect_Ref.modify(function (v1) {
                                                return v1 + s | 0;
                                            })(redRef)();
                                            var $56 = red$prime >= n;
                                            if ($56) {
                                                return true;
                                            };
                                            return false;
                                        };
                                        throw new Error("Failed pattern match at Node.Stream.Aff (line 401, column 42 - line 410, column 29): " + [ v.constructor.name ]);
                                    })()) {

                                    };
                                    return {};
                                })();
                                var red = Effect_Ref.read(redRef)();
                                var $58 = red >= n;
                                if ($58) {
                                    removeError();
                                    removeClose();
                                    removeEnd();
                                    var ret = liftST1(Data_Array_ST.unsafeFreeze(bufs))();
                                    var readagain = Node_Stream.readable(r)();
                                    return complete(new Data_Either.Right({
                                        buffers: ret,
                                        readagain: readagain
                                    }))();
                                };
                                return continuation(Data_Unit.unit)();
                            };
                        };
                        
                        // if there were not enough bytes right away, then wait for bytes to come in.
var waitToRead = function (v) {
                            return function __do() {
                                var removeReadable$prime = Node_EventEmitter.once(Node_Stream.readableH)(tryToRead(waitToRead))(r)();
                                return Effect_Ref.write(removeReadable$prime)(removeReadable)();
                            };
                        };
                        var cleanupRethrow = function (err) {
                            return function __do() {
                                removeError();
                                removeClose();
                                removeEnd();
                                join(Effect_Ref.read(removeReadable))();
                                complete(new Data_Either.Left(err))();
                                return Effect_Aff.nonCanceler;
                            };
                        };
                        return Effect_Exception.catchException(cleanupRethrow)(function __do() {
                            tryToRead(waitToRead)();
                            return pure(Effect_Aff.effectCanceler(function __do() {
                                removeError();
                                removeClose();
                                removeEnd();
                                return join(Effect_Ref.read(removeReadable))();
                            }))();
                        })();
                    };
                    complete(new Data_Either.Right({
                        buffers: [  ],
                        readagain: false
                    }))();
                    return Effect_Aff.nonCanceler;
                };
            }));
        };
    };
};

// | Works on streams in "paused" mode. 
// | Read all data until the end of the stream. After completion the stream
// | will no longer be `readable`.
// |
// | Note that __stdin__ will never end.
var readAll = function (dictMonadAff) {
    var liftAff = Effect_Aff_Class.liftAff(dictMonadAff);
    return function (r) {
        return liftAff(Effect_Aff.makeAff(function (complete) {
            return function __do() {
                var isReadable = Node_Stream.readable(r)();
                if (isReadable) {
                    var bufs = liftST(Data_Array_ST["new"])();
                    var removeReadable = Effect_Ref["new"](Control_Applicative.pure(Effect.applicativeEffect)(Data_Unit.unit))();
                    var removeError = Node_EventEmitter.once(Node_Stream.errorH)(function (err) {
                        return function __do() {
                            join(Effect_Ref.read(removeReadable))();
                            return complete(new Data_Either.Left(err))();
                        };
                    })(r)();
                    var removeClose = Node_EventEmitter.once(Node_Stream.closeH)(function __do() {
                        removeError();
                        join(Effect_Ref.read(removeReadable))();
                        var ret = liftST1(Data_Array_ST.unsafeFreeze(bufs))();
                        return complete(new Data_Either.Right(ret))();
                    })(r)();
                    var removeEnd = Node_EventEmitter.once(Node_Stream.endH)(function __do() {
                        removeError();
                        removeClose();
                        var ret = liftST1(Data_Array_ST.unsafeFreeze(bufs))();
                        return complete(new Data_Either.Right(ret))();
                    })(r)();
                    var cleanupRethrow = function (err) {
                        return function __do() {
                            removeError();
                            removeClose();
                            removeEnd();
                            join(Effect_Ref.read(removeReadable))();
                            complete(new Data_Either.Left(err))();
                            return Effect_Aff.nonCanceler;
                        };
                    };
                    return Effect_Exception.catchException(cleanupRethrow)(function __do() {
                        (function () {
                            while (!(function __do() {
                                var v = Node_Stream.read(r)();
                                if (v instanceof Data_Maybe.Nothing) {
                                    return true;
                                };
                                if (v instanceof Data_Maybe.Just) {
                                    $$void(liftST2(Data_Array_ST.push(v.value0)(bufs)))();
                                    return false;
                                };
                                throw new Error("Failed pattern match at Node.Stream.Aff (line 311, column 27 - line 315, column 23): " + [ v.constructor.name ]);
                            })()) {

                            };
                            return {};
                        })();
                        var $lazy_waitToRead = $runtime_lazy("waitToRead", "Node.Stream.Aff", function () {
                            return function __do() {
                                var removeReadable$prime = Node_EventEmitter.once(Node_Stream.readableH)(function __do() {
                                    (function () {
                                        while (!(function __do() {
                                            var v = Node_Stream.read(r)();
                                            if (v instanceof Data_Maybe.Nothing) {
                                                return true;
                                            };
                                            if (v instanceof Data_Maybe.Just) {
                                                liftST2(Data_Array_ST.push(v.value0)(bufs))();
                                                return false;
                                            };
                                            throw new Error("Failed pattern match at Node.Stream.Aff (line 324, column 33 - line 328, column 29): " + [ v.constructor.name ]);
                                        })()) {

                                        };
                                        return {};
                                    })();
                                    return $lazy_waitToRead(329)();
                                })(r)();
                                return Effect_Ref.write(removeReadable$prime)(removeReadable)();
                            };
                        });
                        var waitToRead = $lazy_waitToRead(319);
                        waitToRead();
                        return pure(Effect_Aff.effectCanceler(function __do() {
                            removeError();
                            removeClose();
                            removeEnd();
                            return join(Effect_Ref.read(removeReadable))();
                        }))();
                    })();
                };
                complete(new Data_Either.Right([  ]))();
                return Effect_Aff.nonCanceler;
            };
        }));
    };
};

// | Encode a `String` as an `Array` containing one UTF-8 encoded `Buffer`.
// |
// | Example:
// |
// | ```
// | write stream =<< fromStringUTF8 "outputstring"
// | ```
var fromStringUTF8 = function (dictMonadEffect) {
    var liftEffect = Effect_Class.liftEffect(dictMonadEffect);
    return function (s) {
        return liftEffect(Data_Functor.map(Effect.functorEffect)(pure1)(Node_Buffer.fromString(s)(Node_Encoding.UTF8.value)));
    };
};

// | Signal that no more data will be written to the `Writable`. Will complete
// | after all data is written and flushed.
// |
// | When the `Writable` is an [__fs.WriteStream__](https://nodejs.org/api/fs.html#class-fswritestream)
// | then this will close the file descriptor because
// |
// | > “If `autoClose` is set to true (default behavior) on `'error'`
// | > or `'finish'` the file descriptor will be closed automatically.”
var end = function (dictMonadAff) {
    var liftAff = Effect_Aff_Class.liftAff(dictMonadAff);
    return function (w) {
        return liftAff(Effect_Aff.makeAff(function (complete) {
            return function __do() {
                Node_Stream["end$prime"](w)(function (v) {
                    if (v instanceof Data_Maybe.Nothing) {
                        return complete(new Data_Either.Right(Data_Unit.unit));
                    };
                    if (v instanceof Data_Maybe.Just) {
                        return complete(new Data_Either.Left(v.value0));
                    };
                    throw new Error("Failed pattern match at Node.Stream.Aff (line 497, column 19 - line 499, column 36): " + [ v.constructor.name ]);
                })();
                return Effect_Aff.nonCanceler;
            };
        }));
    };
};
export {
    readableToStringUtf8,
    readableToString,
    readableToBuffers,
    readSome,
    readAll,
    readN,
    write,
    end,
    toStringUTF8,
    fromStringUTF8
};
