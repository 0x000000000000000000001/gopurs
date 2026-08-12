// | This module defines types and functions for working with nullable types
// | using the FFI.
import * as $foreign from "./foreign.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Show from "../Data.Show/index.js";

// | Takes `Nothing` to `null`, and `Just a` to `a`.
var toNullable = /* #__PURE__ */ Data_Maybe.maybe($foreign["null"])($foreign.notNull);

// | Represent `null` using `Maybe a` as `Nothing`. Note that this function
// | can violate parametricity, as it inspects the runtime representation of
// | its argument (see the warning about the pitfall of `Nullable` above).
var toMaybe = function (n) {
    return $foreign.nullable(n, Data_Maybe.Nothing.value, Data_Maybe.Just.create);
};
var showNullable = function (dictShow) {
    return {
        show: (function () {
            var $16 = Data_Maybe.maybe("null")(Data_Show.show(dictShow));
            return function ($17) {
                return $16(toMaybe($17));
            };
        })()
    };
};
var eqNullable = function (dictEq) {
    return {
        eq: Data_Function.on(Data_Eq.eq(Data_Maybe.eqMaybe(dictEq)))(toMaybe)
    };
};
var ordNullable = function (dictOrd) {
    var eqNullable1 = eqNullable(dictOrd.Eq0());
    return {
        compare: Data_Function.on(Data_Ord.compare(Data_Maybe.ordMaybe(dictOrd)))(toMaybe),
        Eq0: function () {
            return eqNullable1;
        }
    };
};
var eq1Nullable = {
    eq1: function (dictEq) {
        return Data_Eq.eq(eqNullable(dictEq));
    }
};
var ord1Nullable = {
    compare1: function (dictOrd) {
        return Data_Ord.compare(ordNullable(dictOrd));
    },
    Eq10: function () {
        return eq1Nullable;
    }
};
export {
    null,
    notNull
} from "./foreign.js";
export {
    toMaybe,
    toNullable,
    showNullable,
    eqNullable,
    eq1Nullable,
    ordNullable,
    ord1Nullable
};
