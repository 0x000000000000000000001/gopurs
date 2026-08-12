import * as $foreign from "./foreign.js";
import * as Data_CommutativeRing from "../Data.CommutativeRing/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Ring from "../Data.Ring/index.js";
import * as Data_Semiring from "../Data.Semiring/index.js";
var mod = function (dict) {
    return dict.mod;
};

// | The *greatest common divisor* of two values.
var gcd = function (dictEq) {
    return function (dictEuclideanRing) {
        var zero = Data_Semiring.zero(((dictEuclideanRing.CommutativeRing0()).Ring0()).Semiring0());
        return function (a) {
            return function (b) {
                var $17 = Data_Eq.eq(dictEq)(b)(zero);
                if ($17) {
                    return a;
                };
                return gcd(dictEq)(dictEuclideanRing)(b)(mod(dictEuclideanRing)(a)(b));
            };
        };
    };
};
var euclideanRingNumber = {
    degree: function (v) {
        return 1;
    },
    div: $foreign.numDiv,
    mod: function (v) {
        return function (v1) {
            return 0.0;
        };
    },
    CommutativeRing0: function () {
        return Data_CommutativeRing.commutativeRingNumber;
    }
};
var euclideanRingInt = {
    degree: $foreign.intDegree,
    div: $foreign.intDiv,
    mod: $foreign.intMod,
    CommutativeRing0: function () {
        return Data_CommutativeRing.commutativeRingInt;
    }
};
var div = function (dict) {
    return dict.div;
};

// | The *least common multiple* of two values.
var lcm = function (dictEq) {
    return function (dictEuclideanRing) {
        var Ring0 = (dictEuclideanRing.CommutativeRing0()).Ring0();
        var zero = Data_Semiring.zero(Ring0.Semiring0());
        var Semiring0 = Ring0.Semiring0();
        return function (a) {
            return function (b) {
                var $19 = Data_Eq.eq(dictEq)(a)(zero) || Data_Eq.eq(dictEq)(b)(zero);
                if ($19) {
                    return zero;
                };
                return div(dictEuclideanRing)(Data_Semiring.mul(Semiring0)(a)(b))(gcd(dictEq)(dictEuclideanRing)(a)(b));
            };
        };
    };
};
var degree = function (dict) {
    return dict.degree;
};
export {
    degree,
    div,
    mod,
    gcd,
    lcm,
    euclideanRingInt,
    euclideanRingNumber
};
export {
    sub
} from "../Data.Ring/index.js";
export {
    add,
    mul,
    one,
    zero
} from "../Data.Semiring/index.js";
