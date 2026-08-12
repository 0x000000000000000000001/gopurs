import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Show from "../Data.Show/index.js";

// | A newtype used in cases to specify a replacement for a pattern.
var Replacement = function (x) {
    return x;
};

// | A newtype used in cases where there is a string to be matched.
// |
// | ```purescript
// | pursPattern = Pattern ".purs"
// | --can be used like this:
// | contains pursPattern "Test.purs"
// |    == true
// | ```
// |
var Pattern = function (x) {
    return x;
};
var showReplacement = {
    show: function (v) {
        return "(Replacement " + (Data_Show.show(Data_Show.showString)(v) + ")");
    }
};
var showPattern = {
    show: function (v) {
        return "(Pattern " + (Data_Show.show(Data_Show.showString)(v) + ")");
    }
};
var newtypeReplacement = {
    Coercible0: function () {
        return undefined;
    }
};
var newtypePattern = {
    Coercible0: function () {
        return undefined;
    }
};
var eqReplacement = {
    eq: function (x) {
        return function (y) {
            return x === y;
        };
    }
};
var ordReplacement = {
    compare: function (x) {
        return function (y) {
            return Data_Ord.compare(Data_Ord.ordString)(x)(y);
        };
    },
    Eq0: function () {
        return eqReplacement;
    }
};
var eqPattern = {
    eq: function (x) {
        return function (y) {
            return x === y;
        };
    }
};
var ordPattern = {
    compare: function (x) {
        return function (y) {
            return Data_Ord.compare(Data_Ord.ordString)(x)(y);
        };
    },
    Eq0: function () {
        return eqPattern;
    }
};
export {
    Pattern,
    Replacement,
    eqPattern,
    ordPattern,
    newtypePattern,
    showPattern,
    eqReplacement,
    ordReplacement,
    newtypeReplacement,
    showReplacement
};
