import * as $foreign from "./foreign.js";
import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Function_Uncurried from "../Data.Function.Uncurried/index.js";
import * as Data_Unit from "../Data.Unit/index.js";
var warn = function () {
    return {};
};

// | Measures the time it takes the given function to run and prints it out,
// | then returns the function's result. This is handy for diagnosing
// | performance problems by wrapping suspected parts of the code in
// | `traceTime`.
// |
// | For example:
// | ```purescript
// | bunchOfThings =
// |   [ traceTime "one" \_ -> one x y
// |   , traceTime "two" \_ -> two z
// |   , traceTime "three" \_ -> three a b c
// |   ]
// | ```
// |
// | Console output would look something like this:
// | ```
// | one took 3.456ms
// | two took 562.0023ms
// | three took 42.0111ms
// | ```
// |
// | Note that the timing precision may differ depending on whether the
// | Performance API is supported. Where supported (on most modern browsers and
// | versions of Node), the Performance API offers timing resolution of 5
// | microseconds. Where Performance API is not supported, this function will
// | fall back on standard JavaScript Date object, which only offers a
// | 1-millisecond resolution.
var traceTime = function () {
    return Data_Function_Uncurried.runFn2($foreign["_traceTime"]);
};

// | Log any PureScript value to the console for debugging purposes and then
// | return a value. This will log the value's underlying representation for
// | low-level debugging, so it may be desireable to `show` the value first.
// |
// | The return value is thunked so it is not evaluated until after the
// | message has been printed, to preserve a predictable console output.
// |
// | For example:
// | ``` purescript
// | doSomething = trace "Hello" \_ -> ... some value or computation ...
// | ```
var trace = function () {
    return function (a) {
        return function (k) {
            return $foreign["_trace"](a, k);
        };
    };
};

// | Log any PureScript value to the console and return the unit value of the
// | Monad `m`.
var traceM = function () {
    return function (dictMonad) {
        var Bind1 = dictMonad.Bind1();
        var Applicative0 = dictMonad.Applicative0();
        return function (s) {
            return Control_Bind.discard(Control_Bind.discardUnit)(Bind1)(Control_Applicative.pure(Applicative0)(Data_Unit.unit))(function () {
                return trace()(s)(function (v) {
                    return Control_Applicative.pure(Applicative0)(Data_Unit.unit);
                });
            });
        };
    };
};

// | Logs any value and returns it, using a "tag" or key value to annotate the
// | traced value. Useful when debugging something in the middle of a
// | expression, as you can insert this into the expression without having to
// | break it up.
var spy = function () {
    return function (tag) {
        return function (a) {
            return $foreign["_spy"](tag, a);
        };
    };
};

// | Similar to `spy`, but allows a function to be passed in to alter the value
// | that will be printed. Useful in cases where the raw printed form of a value
// | is inconvenient to read - for example, when spying on a `Set`, passing
// | `Array.fromFoldable` here will print it in a more useful form.
var spyWith = function () {
    return function (msg) {
        return function (f) {
            return function (a) {
                return Data_Function["const"](a)(spy()(msg)(f(a)));
            };
        };
    };
};

// | Triggers any available debugging features in the current runtime - in a
// | web browser with the debug tools open, this acts like setting a breakpoint
// | in the script. If no debugging feature are available nothing will occur,
// | although the passed contination will still be evaluated.
// |
// | Generally this works best by passing in a block of code to debug as the
// | continuation argument, as stepping forward in the debugger will then drop
// | straight into the passed code block.
var $$debugger = function () {
    return function (f) {
        return $foreign["_debugger"](f);
    };
};
export {
    trace,
    traceM,
    traceTime,
    spy,
    spyWith,
    $$debugger as debugger,
    warn
};
