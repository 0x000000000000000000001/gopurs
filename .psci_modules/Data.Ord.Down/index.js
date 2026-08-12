import * as Data_Bounded from "../Data.Bounded/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Ordering from "../Data.Ordering/index.js";
import * as Data_Show from "../Data.Show/index.js";

// | A newtype wrapper which provides a reversed `Ord` instance. For example:
// |
// |     sortBy (comparing Down) [1,2,3] = [3,2,1]
// |
var Down = function (x) {
    return x;
};
var showDown = function (dictShow) {
    return {
        show: function (v) {
            return "(Down " + (Data_Show.show(dictShow)(v) + ")");
        }
    };
};
var newtypeDown = {
    Coercible0: function () {
        return undefined;
    }
};
var eqDown = function (dictEq) {
    return dictEq;
};
var ordDown = function (dictOrd) {
    var eqDown1 = eqDown(dictOrd.Eq0());
    return {
        compare: function (v) {
            return function (v1) {
                return Data_Ordering.invert(Data_Ord.compare(dictOrd)(v)(v1));
            };
        },
        Eq0: function () {
            return eqDown1;
        }
    };
};
var boundedDown = function (dictBounded) {
    var ordDown1 = ordDown(dictBounded.Ord0());
    return {
        top: Data_Bounded.bottom(dictBounded),
        bottom: Data_Bounded.top(dictBounded),
        Ord0: function () {
            return ordDown1;
        }
    };
};
export {
    Down,
    newtypeDown,
    eqDown,
    ordDown,
    boundedDown,
    showDown
};
