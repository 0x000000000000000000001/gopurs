// | This module defines the `Cont`inuation monad.
import * as Control_Monad_Cont_Class from "../Control.Monad.Cont.Class/index.js";
import * as Control_Monad_Cont_Trans from "../Control.Monad.Cont.Trans/index.js";
import * as Control_Semigroupoid from "../Control.Semigroupoid/index.js";
import * as Data_Identity from "../Data.Identity/index.js";
import * as Data_Newtype from "../Data.Newtype/index.js";
var unwrap = /* #__PURE__ */ Data_Newtype.unwrap();

// | Transform the continuation passed into the continuation-passing function.
var withCont = function (f) {
    return Control_Monad_Cont_Trans.withContT((function () {
        var $1 = Control_Semigroupoid.compose(Control_Semigroupoid.semigroupoidFn)(Data_Identity.Identity);
        var $2 = Control_Semigroupoid.compose(Control_Semigroupoid.semigroupoidFn)(unwrap);
        return function ($3) {
            return $1(f($2($3)));
        };
    })());
};

// | Runs a computation in the `Cont` monad.
var runCont = function (cc) {
    return function (k) {
        return Data_Newtype.unwrap()(Control_Monad_Cont_Trans.runContT(cc)(function ($4) {
            return Data_Identity.Identity(k($4));
        }));
    };
};

// | Transform the result of a continuation-passing function.
var mapCont = function (f) {
    return Control_Monad_Cont_Trans.mapContT(function ($5) {
        return Data_Identity.Identity(f(unwrap($5)));
    });
};

// | Creates a computation in the `Cont` monad.
var cont = function (f) {
    return function (c) {
        return f(function ($6) {
            return unwrap(c($6));
        });
    };
};
export {
    cont,
    runCont,
    mapCont,
    withCont
};
export {
    callCC
} from "../Control.Monad.Cont.Class/index.js";
export {
    ContT,
    lift,
    mapContT,
    runContT,
    withContT
} from "../Control.Monad.Cont.Trans/index.js";
