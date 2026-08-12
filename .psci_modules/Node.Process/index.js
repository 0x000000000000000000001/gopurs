// | Bindings to the global `process` object in Node.js. See also [the Node API documentation](https://nodejs.org/api/process.html)
import * as $foreign from "./foreign.js";
import * as Control_Category from "../Control.Category/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Nullable from "../Data.Nullable/index.js";
import * as Data_Posix_Signal from "../Data.Posix.Signal/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Data_String_Common from "../Data.String.Common/index.js";
import * as Effect from "../Effect/index.js";
import * as Effect_Uncurried from "../Effect.Uncurried/index.js";
import * as Foreign_Object from "../Foreign.Object/index.js";
import * as Node_EventEmitter from "../Node.EventEmitter/index.js";
import * as Node_Platform from "../Node.Platform/index.js";
var showRecord = /* #__PURE__ */ Data_Show.showRecord()()(/* #__PURE__ */ Data_Show.showRecordFieldsCons({
    reflectSymbol: function () {
        return "system";
    }
})(/* #__PURE__ */ Data_Show.showRecordFieldsConsNil({
    reflectSymbol: function () {
        return "user";
    }
})(Data_Show.showInt))(Data_Show.showInt));
var identity = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var CpuUsage = function (x) {
    return x;
};
var showCpuUsage = {
    show: function (v) {
        return "(CpuUsage " + (Data_Show.show(showRecord)(v) + ")");
    }
};
var eqCpuUsage = {
    eq: function (x) {
        return function (y) {
            return x.system === y.system && x.user === y.user;
        };
    }
};

// | Args:
// | - `worker` <Worker> The <Worker> that was created.
// | 
// | The 'worker' event is emitted after a new <Worker> thread has been created.
var workerH = /* #__PURE__ */ (function () {
    return new Node_EventEmitter.EventHandle("worker", Effect_Uncurried.mkEffectFn1);
})();

// | Args:
// | - `warning` <Error> Key properties of the warning are:
// |   - `name` <string> The name of the warning. Default: 'Warning'.
// |   - `message` <string> A system-provided description of the warning.
// |   - `stack` <string> A stack trace to the location in the code where the warning was issued.
// | 
// | The 'warning' event is emitted whenever Node.js emits a process warning.
// | 
// | A process warning is similar to an error in that it describes exceptional conditions that are being brought to the user's attention. However, warnings are not part of the normal Node.js and JavaScript error handling flow. Node.js can emit warnings whenever it detects bad coding practices that could lead to sub-optimal application performance, bugs, or security vulnerabilities.
// | By default, Node.js will print process warnings to stderr. The --no-warnings command-line option can be used to suppress the default console output but the 'warning' event will still be emitted by the process object.
var warningH = /* #__PURE__ */ (function () {
    return new Node_EventEmitter.EventHandle("warning", Effect_Uncurried.mkEffectFn1);
})();

// | Delete an environment variable.
// | Use case: to hide secret environment variable from child processes.
var unsetEnv = function (key) {
    return function () {
        return $foreign.unsetEnvImpl(key);
    };
};

// | Unsafe because process must be a Node child process and an IPC channel must exist.
var unsafeSendOptsCb = function () {
    return function (msg) {
        return function (handle) {
            return function (opts) {
                return function (cb) {
                    return function () {
                        return $foreign.sendOptsCbImpl(msg, handle, opts, function (err) {
                            return cb(Data_Nullable.toMaybe(err))();
                        });
                    };
                };
            };
        };
    };
};

// | Unsafe because process must be a Node child process and an IPC channel must exist.
var unsafeSendOpts = function () {
    return function (msg) {
        return function (handle) {
            return function (opts) {
                return function () {
                    return $foreign.sendOptsImpl(msg, handle, opts);
                };
            };
        };
    };
};

// | Unsafe because process must be a Node child process and an IPC channel must exist.
var unsafeSendCb = function (msg) {
    return function (handle) {
        return function (cb) {
            return function () {
                return $foreign.sendCbImpl(msg, handle, function (err) {
                    return cb(Data_Nullable.toMaybe(err))();
                });
            };
        };
    };
};

// | Unsafe because child process must be a Node child process and an IPC channel must exist.
var unsafeSend = function (msg) {
    return function (handle) {
        return function () {
            return $foreign.sendImpl(msg, handle);
        };
    };
};

// | Args:
// | - `reason` <Error> | <any> The object with which the promise was rejected (typically an Error object).
// | - `promise` <Promise> The rejected promise.
// |
// | The 'unhandledRejection' event is emitted whenever a Promise is rejected and no error handler is attached to the promise within a turn of the event loop. When programming with Promises, exceptions are encapsulated as "rejected promises". Rejections can be caught and handled using promise.catch() and are propagated through a Promise chain. The 'unhandledRejection' event is useful for detecting and keeping track of promises that were rejected whose rejections have not yet been handled.
var unhandledRejectionH = /* #__PURE__ */ (function () {
    return new Node_EventEmitter.EventHandle("unhandledRejection", function (cb) {
        return function (a, b) {
            return cb(a)(b)();
        };
    });
})();

// | Args:
// | - `err` <Error> The uncaught exception.
// | - `origin` <string> Indicates if the exception originates from an unhandled rejection or from a synchronous error. 
// |    Can either be 'uncaughtException' or 'unhandledRejection'. 
// |    The latter is used when an exception happens in a Promise based async context (or if a Promise is rejected) 
// |    and `--unhandled-rejections` flag set to `strict` or `throw` (which is the default) and 
// |    the rejection is not handled, or when a rejection happens during the command line entry point's 
// |    ES module static loading phase.
// |
// | The 'uncaughtException' event is emitted when an uncaught JavaScript exception bubbles 
// | all the way back to the event loop. By default, Node.js handles such exceptions 
// | by printing the stack trace to `stderr` and exiting with code 1, 
// | overriding any previously set `process.exitCode`. 
// | Adding a handler for the 'uncaughtException' event overrides this default behavior. 
// | Alternatively, change the process.exitCode in the 'uncaughtException' handler which will 
// | result in the process exiting with the provided exit code. Otherwise, in the presence of 
// | such handler the process will exit with 0.
// |
// | 'uncaughtException' is a crude mechanism for exception handling intended to be used only as a last resort. The event should not be used as an equivalent to On Error Resume Next. Unhandled exceptions inherently mean that an application is in an undefined state. Attempting to resume application code without properly recovering from the exception can cause additional unforeseen and unpredictable issues.
// | 
// | Exceptions thrown from within the event handler will not be caught. Instead the process will exit with a non-zero exit code and the stack trace will be printed. This is to avoid infinite recursion.
// | 
// | Attempting to resume normally after an uncaught exception can be similar to pulling out the power cord when upgrading a computer. Nine out of ten times, nothing happens. But the tenth time, the system becomes corrupted.
// | 
// | The correct use of 'uncaughtException' is to perform synchronous cleanup of allocated resources (e.g. file descriptors, handles, etc) before shutting down the process. It is not safe to resume normal operation after 'uncaughtException'.
// | 
// | To restart a crashed application in a more reliable way, whether 'uncaughtException' is emitted or not, an external monitor should be employed in a separate process to detect application failures and recover or restart as needed.
var uncaughtExceptionH = /* #__PURE__ */ (function () {
    return new Node_EventEmitter.EventHandle("uncaughtException", function (cb) {
        return function (a, b) {
            return cb(a)(b)();
        };
    });
})();
var setUncaughtExceptionCaptureCallback = function (cb) {
    return function () {
        return $foreign.setUncaughtExceptionCaptureCallbackImpl(cb);
    };
};

// | The process.title property returns the current process title (i.e. returns the current value of ps). 
// | Assigning a new value to process.title modifies the current value of ps.
// | 
// | When a new value is assigned, different platforms will impose different maximum length restrictions 
// | on the title. Usually such restrictions are quite limited. For instance, on Linux and macOS, 
// | process.title is limited to the size of the binary name plus the length of the command-line arguments 
// | because setting the process.title overwrites the argv memory of the process. 
// | Node.js v0.8 allowed for longer process title strings by also overwriting the environ memory but 
// | that was potentially insecure and confusing in some (rather obscure) cases.
// | 
// | Assigning a value to process.title might not result in an accurate label within process manager 
// | applications such as macOS Activity Monitor or Windows Services Manager.
var setTitle = function (newTitle) {
    return function () {
        return $foreign.setTitleImpl(newTitle);
    };
};

// | Sets the exit code to use when the process exits.
// | An exit code of 0 is normally considered successful, and anything else is considered a
// | failure.
// |
// | In comparison to `exit`/`exit'`, this is the safer way 
// | to exit a Node.js process because any pending asynchronous operations
// | including I/O operations to `process.stdout` and `process.stderr`
// | will complete before the `Node.js` process exits.
var setExitCode = function (code) {
    return function () {
        return $foreign.setExitCodeImpl(code);
    };
};

// | Set an environment variable.
var setEnv = function (key) {
    return function (value) {
        return function () {
            return $foreign.setEnvImpl(key, value);
        };
    };
};

// | The 'rejectionHandled' event is emitted whenever a Promise has been rejected and an error handler was attached to it (using promise.catch(), for example) later than one turn of the Node.js event loop.
// | 
// | The Promise object would have previously been emitted in an 'unhandledRejection' event, but during the course of processing gained a rejection handler.
// | 
// | There is no notion of a top level for a Promise chain at which rejections can always be handled. Being inherently asynchronous in nature, a Promise rejection can be handled at a future point in time, possibly much later than the event loop turn it takes for the 'unhandledRejection' event to be emitted.
// | 
// | Another way of stating this is that, unlike in synchronous code where there is an ever-growing list of unhandled exceptions, with Promises there can be a growing-and-shrinking list of unhandled rejections.
// | 
// | In synchronous code, the 'uncaughtException' event is emitted when the list of unhandled exceptions grows.
// | 
// | In asynchronous code, the 'unhandledRejection' event is emitted when the list of unhandled rejections grows, and the 'rejectionHandled' event is emitted when the list of unhandled rejections shrinks.
var rejectionHandledH = /* #__PURE__ */ (function () {
    return new Node_EventEmitter.EventHandle("rejectionHandled", Effect_Uncurried.mkEffectFn1);
})();
var platform = /* #__PURE__ */ Node_Platform.fromString($foreign.platformStr);

// | Register a callback that will receive the record arg as an argument
// | to run as soon as the current event loop runs to completion.
// | One should use `queueMicroTask` instead for most situations instead of `nextTick`. 
// | See Node docs for more info.
var nextTick$prime = function (cb) {
    return function (args) {
        return function () {
            return $foreign.nextTickCbImpl(Effect_Uncurried.mkEffectFn1(cb), args);
        };
    };
};

// | Register a callback to run as soon as the current event loop runs to
// | completion.
// | One should use `queueMicroTask` instead for most situations instead of `nextTick`. 
// | See Node docs for more info.
var nextTick = function (cb) {
    return function () {
        return $foreign.nextTickImpl(cb);
    };
};

// | Same as `mkSignalH` but allows for more options in case the `Signal` ADT is missing any.
// |
// | See Node docs: https://nodejs.org/dist/latest-v18.x/docs/api/process.html#signal-events
var mkSignalH$prime = function (sig) {
    return new Node_EventEmitter.EventHandle(Data_String_Common.toUpper(sig), identity);
};

// | Rather than support an `EventHandle` for every possible `Signal`,
// | this function provides one a convenient way for constructing one for any given signal.
// |
// | See Node docs: https://nodejs.org/dist/latest-v18.x/docs/api/process.html#signal-events
var mkSignalH = function (sig) {
    return new Node_EventEmitter.EventHandle(Data_Posix_Signal.toString(sig), identity);
};
var messageH = /* #__PURE__ */ (function () {
    return new Node_EventEmitter.EventHandle("message", function (cb) {
        return function (a, b) {
            return cb(a)(Data_Nullable.toMaybe(b))();
        };
    });
})();

// | Lookup a particular environment variable.
var lookupEnv = function (k) {
    return Data_Functor.map(Effect.functorEffect)(Foreign_Object.lookup(k))($foreign.unsafeGetEnv);
};
var killStr = function (p) {
    return function (sig) {
        return function () {
            return $foreign.killStrImpl(p, sig);
        };
    };
};
var killInt = function (p) {
    return function (sig) {
        return function () {
            return $foreign.killIntImpl(p, sig);
        };
    };
};
var kill = function (p) {
    return function () {
        return $foreign.killImpl(p);
    };
};
var getUid = /* #__PURE__ */ Data_Functor.map(Effect.functorEffect)(Data_Nullable.toMaybe)($foreign.getUidImpl);
var getGid = /* #__PURE__ */ Data_Functor.map(Effect.functorEffect)(Data_Nullable.toMaybe)($foreign.getGidImpl);

// | Gets the currently set exit code. This will be `Nothing`
// | if the exit code has not been previously set.
var getExitCode = /* #__PURE__ */ Data_Functor.map(Effect.functorEffect)(Data_Nullable.toMaybe)($foreign.getExitCodeImpl);

// | The 'exit' event is emitted when the Node.js process is about to exit as a result of either:
// | - The process.exit() method being called explicitly;
// | - The Node.js event loop no longer having any additional work to perform.
// |
// | Listener functions **must** only perform **synchronous** operations. 
// | The Node.js process will exit immediately after calling the 'exit' event listeners causing 
// | any additional work still queued in the event loop to be abandoned.
// | (Maintainer note: I believe the above translates to
// | "Only synchronous (i.e. `Effect`) code can be run in the resulting handler.
// | If you need asynchronous (i.e. `Aff`) code, use `beforeExitH`.")
// |
// | There is no way to prevent the exiting of the event loop at this point, and once all 'exit' 
// | listeners have finished running the Node.js process will terminate.
// | The listener callback function is invoked with the exit code specified either by the 
// | `process.exitCode` property, or the exitCode argument passed to the `process.exit()` method.
var exitH = /* #__PURE__ */ (function () {
    return new Node_EventEmitter.EventHandle("exit", Effect_Uncurried.mkEffectFn1);
})();

// | Cause the process to exit immediately without completing any pending asynchronous operations
// | including I/O operations to `process.stdout` and `process.stderr`.
// | An exit code of 0 is normally considered successful, and anything else is considered a
// | failure.
// |
// | Rather than using this function to exit, one should set `process.exitCode` and
// | let Node gracefully exit once all pending asynchronous operations have completed.
// | If it is necessary to terminate the Node.js process due to an error condition, 
// | throwing an uncaught error and allowing the process to terminate accordingly 
// | is safer than calling process.exit().
var exit$prime = function (code) {
    return function () {
        return $foreign.exitImpl(code);
    };
};
var disconnectH = /* #__PURE__ */ (function () {
    return new Node_EventEmitter.EventHandle("disconnect", identity);
})();

// | If the Node.js process is spawned with an IPC channel (see the Child Process and Cluster documentation), the process.disconnect() method will close the IPC channel to the parent process, allowing the child process to exit gracefully once there are no other connections keeping it alive.
// | The effect of calling process.disconnect() is the same as calling ChildProcess.disconnect() from the parent process.
// |
// | If the Node.js process was not spawned with an IPC channel, process.disconnect() will be undefined.
var disconnect = /* #__PURE__ */ Data_Nullable.toMaybe($foreign.disconnectImpl);
var cpuUsageToRecord = function (v) {
    return v;
};
var cpuUsageDiff = function (prev) {
    return function () {
        return $foreign.cpuUsageDiffImpl(prev);
    };
};

// | Change the current working directory of the process. If the current
// | directory could not be changed, an exception will be thrown.
var chdir = function (dir) {
    return function () {
        return $foreign.chdirImpl(dir);
    };
};

// | This method makes the IPC channel not keep the event loop of the process running, 
// | and lets it finish even while the channel is open.
// | Typically, this is managed through the number of 'disconnect' and 'message' listeners on the process object. 
// | However, this method can be used to explicitly request a specific behavior.
// |
// | Context: if the Node.js process was spawned with an IPC channel (see the Child Process documentation), 
// | the process.channel property is a reference to the IPC channel. If no IPC channel exists, this property is undefined.
var channelUnref = /* #__PURE__ */ Data_Nullable.toMaybe($foreign.channelUnrefImpl);

// | This method makes the IPC channel keep the event loop of the process running if .unref() has been called before.
// | Typically, this is managed through the number of 'disconnect' and 'message' listeners on the process object. 
// | However, this method can be used to explicitly request a specific behavior.
// |
// | Context: if the Node.js process was spawned with an IPC channel (see the Child Process documentation), 
// | the process.channel property is a reference to the IPC channel. If no IPC channel exists, this property is undefined.
var channelRef = /* #__PURE__ */ Data_Nullable.toMaybe($foreign.channelRefImpl);

// | The 'beforeExit' event is emitted when Node.js empties its event loop and has no additional work to schedule. 
// | Normally, the Node.js process will exit when there is no work scheduled, but a listener registered on the 
// | 'beforeExit' event can make asynchronous calls, and thereby cause the Node.js process to continue.
// | 
// | The listener callback function is invoked with the value of process.exitCode passed as the only argument.
// | The 'beforeExit' event is not emitted for conditions causing explicit termination, 
// | such as calling `process.exit()` or uncaught exceptions.
// | The 'beforeExit' should not be used as an alternative to the 'exit' event unless the 
// | intention is to schedule additional work.
var beforeExitH = /* #__PURE__ */ (function () {
    return new Node_EventEmitter.EventHandle("beforeExit", Effect_Uncurried.mkEffectFn1);
})();

// | The `process.abort()` method causes the Node.js process to exit immediately and generate a core file.
// | This feature is not available in Worker threads.
var abort = /* #__PURE__ */ Data_Nullable.toMaybe($foreign.abortImpl);
export {
    process,
    argv,
    argv0,
    config,
    connected,
    cpuUsage,
    cwd,
    debugPort,
    getEnv,
    execArgv,
    execPath,
    exit,
    hasUncaughtExceptionCaptureCallback,
    memoryUsage,
    memoryUsageRss,
    pid,
    platformStr,
    ppid,
    resourceUsage,
    clearUncaughtExceptionCaptureCallback,
    stdin,
    stdout,
    stderr,
    stdinIsTTY,
    stdoutIsTTY,
    stderrIsTTY,
    getTitle,
    uptime,
    version
} from "./foreign.js";
export {
    beforeExitH,
    disconnectH,
    exitH,
    messageH,
    rejectionHandledH,
    uncaughtExceptionH,
    unhandledRejectionH,
    mkSignalH,
    mkSignalH$prime,
    warningH,
    workerH,
    abort,
    channelRef,
    channelUnref,
    chdir,
    cpuUsageToRecord,
    cpuUsageDiff,
    disconnect,
    lookupEnv,
    setEnv,
    unsetEnv,
    exit$prime,
    setExitCode,
    getExitCode,
    getGid,
    getUid,
    kill,
    killStr,
    killInt,
    nextTick,
    nextTick$prime,
    platform,
    unsafeSend,
    unsafeSendOpts,
    unsafeSendCb,
    unsafeSendOptsCb,
    setUncaughtExceptionCaptureCallback,
    setTitle,
    eqCpuUsage,
    showCpuUsage
};
