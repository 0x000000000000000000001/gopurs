import * as Control_Monad_Error_Class from "../Control.Monad.Error.Class/index.js";
import * as Control_Monad_Except_Trans from "../Control.Monad.Except.Trans/index.js";
import * as Data_Identity from "../Data.Identity/index.js";
import * as Data_Newtype from "../Data.Newtype/index.js";
var unwrap = /* #__PURE__ */ Data_Newtype.unwrap();

// | Transform any exceptions thrown by an `Except` computation using the given function.
var withExcept = /* #__PURE__ */ Control_Monad_Except_Trans.withExceptT(Data_Identity.functorIdentity);

// | Run a computation in the `Except` monad. The inverse of `except`.
var runExcept = function ($2) {
    return unwrap(Control_Monad_Except_Trans.runExceptT($2));
};

// | Transform the unwrapped computation using the given function.
var mapExcept = function (f) {
    return Control_Monad_Except_Trans.mapExceptT(function ($3) {
        return Data_Identity.Identity(f(unwrap($3)));
    });
};
export {
    runExcept,
    mapExcept,
    withExcept
};
export {
    catchError,
    catchJust,
    throwError
} from "../Control.Monad.Error.Class/index.js";
export {
    ExceptT,
    except,
    lift,
    mapExceptT,
    runExceptT,
    withExceptT
} from "../Control.Monad.Except.Trans/index.js";
