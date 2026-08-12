// | A _trampoline_ monad, which can be used at the bottom of
// | a monad transformer stack to avoid stack overflows in large
// | monadic computations.
import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Monad_Free from "../Control.Monad.Free/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Unit from "../Data.Unit/index.js";

// | Run a computation in the `Trampoline` monad.
var runTrampoline = /* #__PURE__ */ Control_Monad_Free.runFree(Data_Functor.functorFn)(function (v) {
    return v(Data_Unit.unit);
});

// | Return a value immediately
var done = /* #__PURE__ */ Control_Applicative.pure(Control_Monad_Free.freeApplicative);

// | Use the `Trampoline` monad to represent the delayed evaluation of a value.
var delay = Control_Monad_Free.liftF;
export {
    done,
    delay,
    runTrampoline
};
