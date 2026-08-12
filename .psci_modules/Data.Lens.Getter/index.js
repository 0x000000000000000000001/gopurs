// | This module defines functions for working with getters.
import * as Control_Category from "../Control.Category/index.js";
import * as Control_Monad_State_Class from "../Control.Monad.State.Class/index.js";
import * as Control_Semigroupoid from "../Control.Semigroupoid/index.js";
import * as Data_Lens_Types from "../Data.Lens.Types/index.js";
import * as Data_Newtype from "../Data.Newtype/index.js";
import * as Data_Profunctor_Strong from "../Data.Profunctor.Strong/index.js";
var identity = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var identity1 = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);

// | View the focus of a `Getter`.
var view = function (l) {
    return Data_Newtype.unwrap()(l(identity));
};

// | Synonym for `view`, flipped.
var viewOn = function (s) {
    return function (l) {
        return view(l)(s);
    };
};

// | View the focus of a `Getter` in the state of a monad.
var use = function (dictMonadState) {
    return function (p) {
        return Control_Monad_State_Class.gets(dictMonadState)(function (v) {
            return viewOn(v)(p);
        });
    };
};

// | Convert a function into a getter.
var to = function (f) {
    return function (p) {
        var $5 = Data_Newtype.unwrap()(p);
        return function ($6) {
            return $5(f($6));
        };
    };
};

// | Combine two getters.
var takeBoth = function (l) {
    return function (r) {
        return to(Data_Profunctor_Strong.fanout(Control_Semigroupoid.semigroupoidFn)(Data_Profunctor_Strong.strongFn)(view(l))(view(r)));
    };
};

// | View the focus of a `Getter` and its index.
var iview = function (l) {
    return Data_Newtype.unwrap()(l(identity1));
};

// | View the focus of a `Getter` and its index in the state of a monad.
var iuse = function (dictMonadState) {
    return function (p) {
        return Control_Monad_State_Class.gets(dictMonadState)(iview(p));
    };
};
var cloneGetter = function (g) {
    return to(view(g));
};
export {
    viewOn,
    view,
    to,
    takeBoth,
    use,
    iview,
    iuse,
    cloneGetter
};
export {
    Indexed
} from "../Data.Lens.Types/index.js";
