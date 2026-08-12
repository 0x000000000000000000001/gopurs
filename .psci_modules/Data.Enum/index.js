import * as $foreign from "./foreign.js";
import * as Control_Alternative from "../Control.Alternative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Data_Boolean from "../Data.Boolean/index.js";
import * as Data_Bounded from "../Data.Bounded/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Ordering from "../Data.Ordering/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Data_Unfoldable from "../Data.Unfoldable/index.js";
import * as Data_Unfoldable1 from "../Data.Unfoldable1/index.js";
import * as Data_Unit from "../Data.Unit/index.js";
var bottom = /* #__PURE__ */ Data_Bounded.bottom(Data_Bounded.boundedInt);
var fromJust = /* #__PURE__ */ Data_Maybe.fromJust();
var bottom1 = /* #__PURE__ */ Data_Bounded.bottom(Data_Bounded.boundedChar);
var top = /* #__PURE__ */ Data_Bounded.top(Data_Bounded.boundedChar);
var Cardinality = function (x) {
    return x;
};
var toEnum = function (dict) {
    return dict.toEnum;
};
var succ = function (dict) {
    return dict.succ;
};

// | Produces all successors of an `Enum` value, including the start value.
// |
// | `upFromIncluding bottom` will return all values in an `Enum`.
var upFromIncluding = function (dictEnum) {
    var succ1 = succ(dictEnum);
    return function (dictUnfoldable1) {
        return Data_Unfoldable1.unfoldr1(dictUnfoldable1)(Control_Apply.apply(Control_Apply.applyFn)(Data_Tuple.Tuple.create)(succ1));
    };
};
var showCardinality = {
    show: function (v) {
        return "(Cardinality " + (Data_Show.show(Data_Show.showInt)(v) + ")");
    }
};
var pred = function (dict) {
    return dict.pred;
};
var ordCardinality = Data_Ord.ordInt;
var newtypeCardinality = {
    Coercible0: function () {
        return undefined;
    }
};
var fromEnum = function (dict) {
    return dict.fromEnum;
};

// | Like `toEnum` but returns the first argument if `x` is less than
// | `fromEnum bottom` and the second argument if `x` is greater than
// | `fromEnum top`.
// |
// | ``` purescript
// | toEnumWithDefaults False True (-1) -- False
// | toEnumWithDefaults False True 0    -- False
// | toEnumWithDefaults False True 1    -- True
// | toEnumWithDefaults False True 2    -- True
// | ```
var toEnumWithDefaults = function (dictBoundedEnum) {
    var bottom2 = Data_Bounded.bottom(dictBoundedEnum.Bounded0());
    return function (low) {
        return function (high) {
            return function (x) {
                var v = toEnum(dictBoundedEnum)(x);
                if (v instanceof Data_Maybe.Just) {
                    return v.value0;
                };
                if (v instanceof Data_Maybe.Nothing) {
                    var $95 = x < fromEnum(dictBoundedEnum)(bottom2);
                    if ($95) {
                        return low;
                    };
                    return high;
                };
                throw new Error("Failed pattern match at Data.Enum (line 158, column 33 - line 160, column 62): " + [ v.constructor.name ]);
            };
        };
    };
};
var eqCardinality = Data_Eq.eqInt;
var enumUnit = /* #__PURE__ */ (function () {
    return {
        succ: Data_Function["const"](Data_Maybe.Nothing.value),
        pred: Data_Function["const"](Data_Maybe.Nothing.value),
        Ord0: function () {
            return Data_Ord.ordUnit;
        }
    };
})();
var enumTuple = function (dictEnum) {
    var ordTuple = Data_Tuple.ordTuple(dictEnum.Ord0());
    return function (dictBoundedEnum) {
        var Bounded0 = dictBoundedEnum.Bounded0();
        var Enum1 = dictBoundedEnum.Enum1();
        var ordTuple1 = ordTuple((dictBoundedEnum.Enum1()).Ord0());
        return {
            succ: function (v) {
                return Data_Maybe.maybe(Data_Functor.map(Data_Maybe.functorMaybe)(Data_Function.flip(Data_Tuple.Tuple.create)(Data_Bounded.bottom(Bounded0)))(succ(dictEnum)(v.value0)))((function () {
                    var $138 = Data_Tuple.Tuple.create(v.value0);
                    return function ($139) {
                        return Data_Maybe.Just.create($138($139));
                    };
                })())(succ(Enum1)(v.value1));
            },
            pred: function (v) {
                return Data_Maybe.maybe(Data_Functor.map(Data_Maybe.functorMaybe)(Data_Function.flip(Data_Tuple.Tuple.create)(Data_Bounded.top(Bounded0)))(pred(dictEnum)(v.value0)))((function () {
                    var $140 = Data_Tuple.Tuple.create(v.value0);
                    return function ($141) {
                        return Data_Maybe.Just.create($140($141));
                    };
                })())(pred(Enum1)(v.value1));
            },
            Ord0: function () {
                return ordTuple1;
            }
        };
    };
};
var enumOrdering = {
    succ: function (v) {
        if (v instanceof Data_Ordering.LT) {
            return new Data_Maybe.Just(Data_Ordering.EQ.value);
        };
        if (v instanceof Data_Ordering.EQ) {
            return new Data_Maybe.Just(Data_Ordering.GT.value);
        };
        if (v instanceof Data_Ordering.GT) {
            return Data_Maybe.Nothing.value;
        };
        throw new Error("Failed pattern match at Data.Enum (line 72, column 1 - line 78, column 20): " + [ v.constructor.name ]);
    },
    pred: function (v) {
        if (v instanceof Data_Ordering.LT) {
            return Data_Maybe.Nothing.value;
        };
        if (v instanceof Data_Ordering.EQ) {
            return new Data_Maybe.Just(Data_Ordering.LT.value);
        };
        if (v instanceof Data_Ordering.GT) {
            return new Data_Maybe.Just(Data_Ordering.EQ.value);
        };
        throw new Error("Failed pattern match at Data.Enum (line 72, column 1 - line 78, column 20): " + [ v.constructor.name ]);
    },
    Ord0: function () {
        return Data_Ord.ordOrdering;
    }
};
var enumMaybe = function (dictBoundedEnum) {
    var Bounded0 = dictBoundedEnum.Bounded0();
    var Enum1 = dictBoundedEnum.Enum1();
    var ordMaybe = Data_Maybe.ordMaybe((dictBoundedEnum.Enum1()).Ord0());
    return {
        succ: function (v) {
            if (v instanceof Data_Maybe.Nothing) {
                return new Data_Maybe.Just(new Data_Maybe.Just(Data_Bounded.bottom(Bounded0)));
            };
            if (v instanceof Data_Maybe.Just) {
                return Data_Functor.map(Data_Maybe.functorMaybe)(Data_Maybe.Just.create)(succ(Enum1)(v.value0));
            };
            throw new Error("Failed pattern match at Data.Enum (line 80, column 1 - line 84, column 32): " + [ v.constructor.name ]);
        },
        pred: function (v) {
            if (v instanceof Data_Maybe.Nothing) {
                return Data_Maybe.Nothing.value;
            };
            if (v instanceof Data_Maybe.Just) {
                return new Data_Maybe.Just(pred(Enum1)(v.value0));
            };
            throw new Error("Failed pattern match at Data.Enum (line 80, column 1 - line 84, column 32): " + [ v.constructor.name ]);
        },
        Ord0: function () {
            return ordMaybe;
        }
    };
};
var enumInt = {
    succ: function (n) {
        var $108 = n < Data_Bounded.top(Data_Bounded.boundedInt);
        if ($108) {
            return new Data_Maybe.Just(n + 1 | 0);
        };
        return Data_Maybe.Nothing.value;
    },
    pred: function (n) {
        var $109 = n > bottom;
        if ($109) {
            return new Data_Maybe.Just(n - 1 | 0);
        };
        return Data_Maybe.Nothing.value;
    },
    Ord0: function () {
        return Data_Ord.ordInt;
    }
};

// | Returns a contiguous sequence of elements from the first value to the
// | second value (inclusive).
// |
// | ``` purescript
// | enumFromTo 0 3 = [0, 1, 2, 3]
// | enumFromTo 'c' 'a' = ['c', 'b', 'a']
// | ```
// |
// | The example shows `Array` return values, but the result can be any type
// | with an `Unfoldable1` instance.
var enumFromTo = function (dictEnum) {
    var Ord0 = dictEnum.Ord0();
    var Eq0 = Ord0.Eq0();
    var Ord01 = dictEnum.Ord0();
    var succ1 = succ(dictEnum);
    var lessThanOrEq = Data_Ord.lessThanOrEq(Ord0);
    var pred1 = pred(dictEnum);
    var greaterThanOrEq = Data_Ord.greaterThanOrEq(Ord0);
    return function (dictUnfoldable1) {
        var go = function (step) {
            return function (op) {
                return function (to) {
                    return function (a) {
                        return new Data_Tuple.Tuple(a, Control_Bind.bind(Data_Maybe.bindMaybe)(step(a))(function (a$prime) {
                            return Data_Functor.voidLeft(Data_Maybe.functorMaybe)(Control_Alternative.guard(Data_Maybe.alternativeMaybe)(op(a$prime)(to)))(a$prime);
                        }));
                    };
                };
            };
        };
        return function (v) {
            return function (v1) {
                if (Data_Eq.eq(Eq0)(v)(v1)) {
                    return Data_Unfoldable1.singleton(dictUnfoldable1)(v);
                };
                if (Data_Ord.lessThan(Ord01)(v)(v1)) {
                    return Data_Unfoldable1.unfoldr1(dictUnfoldable1)(go(succ1)(lessThanOrEq)(v1))(v);
                };
                if (Data_Boolean.otherwise) {
                    return Data_Unfoldable1.unfoldr1(dictUnfoldable1)(go(pred1)(greaterThanOrEq)(v1))(v);
                };
                throw new Error("Failed pattern match at Data.Enum (line 186, column 14 - line 190, column 51): " + [ v.constructor.name, v1.constructor.name ]);
            };
        };
    };
};

// | Returns a sequence of elements from the first value, taking steps
// | according to the difference between the first and second value, up to
// | (but not exceeding) the third value.
// |
// | ``` purescript
// | enumFromThenTo 0 2 6 = [0, 2, 4, 6]
// | enumFromThenTo 0 3 5 = [0, 3]
// | ```
// |
// | Note that there is no `BoundedEnum` instance for integers, they're just
// | being used here for illustrative purposes to help clarify the behaviour.
// |
// | The example shows `Array` return values, but the result can be any type
// | with an `Unfoldable1` instance.
var enumFromThenTo = function (dictUnfoldable) {
    return function (dictFunctor) {
        return function (dictBoundedEnum) {
            var toEnum1 = toEnum(dictBoundedEnum);
            var go = function (step) {
                return function (to) {
                    return function (e) {
                        if (e <= to) {
                            return new Data_Maybe.Just(new Data_Tuple.Tuple(e, e + step | 0));
                        };
                        if (Data_Boolean.otherwise) {
                            return Data_Maybe.Nothing.value;
                        };
                        throw new Error("Failed pattern match at Data.Enum (line 217, column 5 - line 219, column 28): " + [ step.constructor.name, to.constructor.name, e.constructor.name ]);
                    };
                };
            };
            return function (a) {
                return function (b) {
                    return function (c) {
                        var c$prime = fromEnum(dictBoundedEnum)(c);
                        var b$prime = fromEnum(dictBoundedEnum)(b);
                        var a$prime = fromEnum(dictBoundedEnum)(a);
                        return Data_Functor.map(dictFunctor)(function ($142) {
                            return fromJust(toEnum1($142));
                        })(Data_Unfoldable.unfoldr(dictUnfoldable)(go(b$prime - a$prime | 0)(c$prime))(a$prime));
                    };
                };
            };
        };
    };
};
var enumEither = function (dictBoundedEnum) {
    var Enum1 = dictBoundedEnum.Enum1();
    var Bounded0 = dictBoundedEnum.Bounded0();
    var ordEither = Data_Either.ordEither((dictBoundedEnum.Enum1()).Ord0());
    return function (dictBoundedEnum1) {
        var Bounded01 = dictBoundedEnum1.Bounded0();
        var Enum11 = dictBoundedEnum1.Enum1();
        var ordEither1 = ordEither((dictBoundedEnum1.Enum1()).Ord0());
        return {
            succ: function (v) {
                if (v instanceof Data_Either.Left) {
                    return Data_Maybe.maybe(new Data_Maybe.Just(new Data_Either.Right(Data_Bounded.bottom(Bounded01))))(function ($143) {
                        return Data_Maybe.Just.create(Data_Either.Left.create($143));
                    })(succ(Enum1)(v.value0));
                };
                if (v instanceof Data_Either.Right) {
                    return Data_Maybe.maybe(Data_Maybe.Nothing.value)(function ($144) {
                        return Data_Maybe.Just.create(Data_Either.Right.create($144));
                    })(succ(Enum11)(v.value0));
                };
                throw new Error("Failed pattern match at Data.Enum (line 86, column 1 - line 90, column 69): " + [ v.constructor.name ]);
            },
            pred: function (v) {
                if (v instanceof Data_Either.Left) {
                    return Data_Maybe.maybe(Data_Maybe.Nothing.value)(function ($145) {
                        return Data_Maybe.Just.create(Data_Either.Left.create($145));
                    })(pred(Enum1)(v.value0));
                };
                if (v instanceof Data_Either.Right) {
                    return Data_Maybe.maybe(new Data_Maybe.Just(new Data_Either.Left(Data_Bounded.top(Bounded0))))(function ($146) {
                        return Data_Maybe.Just.create(Data_Either.Right.create($146));
                    })(pred(Enum11)(v.value0));
                };
                throw new Error("Failed pattern match at Data.Enum (line 86, column 1 - line 90, column 69): " + [ v.constructor.name ]);
            },
            Ord0: function () {
                return ordEither1;
            }
        };
    };
};
var enumBoolean = {
    succ: function (v) {
        if (!v) {
            return new Data_Maybe.Just(true);
        };
        return Data_Maybe.Nothing.value;
    },
    pred: function (v) {
        if (v) {
            return new Data_Maybe.Just(false);
        };
        return Data_Maybe.Nothing.value;
    },
    Ord0: function () {
        return Data_Ord.ordBoolean;
    }
};

// | Produces all predecessors of an `Enum` value, including the start value.
// |
// | `downFromIncluding top` will return all values in an `Enum`, in reverse
// | order.
var downFromIncluding = function (dictEnum) {
    var pred1 = pred(dictEnum);
    return function (dictUnfoldable1) {
        return Data_Unfoldable1.unfoldr1(dictUnfoldable1)(Control_Apply.apply(Control_Apply.applyFn)(Data_Tuple.Tuple.create)(pred1));
    };
};
var diag = function (a) {
    return new Data_Tuple.Tuple(a, a);
};

// | Produces all predecessors of an `Enum` value, excluding the start value.
var downFrom = function (dictEnum) {
    var pred1 = pred(dictEnum);
    return function (dictUnfoldable) {
        return Data_Unfoldable.unfoldr(dictUnfoldable)((function () {
            var $147 = Data_Functor.map(Data_Maybe.functorMaybe)(diag);
            return function ($148) {
                return $147(pred1($148));
            };
        })());
    };
};

// | Produces all successors of an `Enum` value, excluding the start value.
var upFrom = function (dictEnum) {
    var succ1 = succ(dictEnum);
    return function (dictUnfoldable) {
        return Data_Unfoldable.unfoldr(dictUnfoldable)((function () {
            var $149 = Data_Functor.map(Data_Maybe.functorMaybe)(diag);
            return function ($150) {
                return $149(succ1($150));
            };
        })());
    };
};

// | Provides a default implementation for `toEnum`.
// |
// | - Assumes `fromEnum bottom = 0`.
// | - Cannot be used in conjuction with `defaultSucc`.
// |
// | Runs in `O(n)` where `n` is `fromEnum a`.
var defaultToEnum = function (dictBounded) {
    return function (dictEnum) {
        return function (i$prime) {
            var go = function ($copy_i) {
                return function ($copy_x) {
                    var $tco_var_i = $copy_i;
                    var $tco_done = false;
                    var $tco_result;
                    function $tco_loop(i, x) {
                        var $123 = i === 0;
                        if ($123) {
                            $tco_done = true;
                            return new Data_Maybe.Just(x);
                        };
                        var v = succ(dictEnum)(x);
                        if (v instanceof Data_Maybe.Just) {
                            $tco_var_i = i - 1 | 0;
                            $copy_x = v.value0;
                            return;
                        };
                        if (v instanceof Data_Maybe.Nothing) {
                            $tco_done = true;
                            return Data_Maybe.Nothing.value;
                        };
                        throw new Error("Failed pattern match at Data.Enum (line 296, column 12 - line 298, column 33): " + [ v.constructor.name ]);
                    };
                    while (!$tco_done) {
                        $tco_result = $tco_loop($tco_var_i, $copy_x);
                    };
                    return $tco_result;
                };
            };
            var $126 = i$prime < 0;
            if ($126) {
                return Data_Maybe.Nothing.value;
            };
            return go(i$prime)(Data_Bounded.bottom(dictBounded));
        };
    };
};

// | Provides a default implementation for `succ`, given a function that maps
// | integers to values in the `Enum`, and a function that maps values in the
// | `Enum` back to integers. The integer mapping must agree in both directions
// | for this to implement a law-abiding `succ`.
// |
// | If a `BoundedEnum` instance exists for `a`, the `toEnum` and `fromEnum`
// | functions can be used here:
// |
// | ``` purescript
// | succ = defaultSucc toEnum fromEnum
// | ```
var defaultSucc = function (toEnum$prime) {
    return function (fromEnum$prime) {
        return function (a) {
            return toEnum$prime(fromEnum$prime(a) + 1 | 0);
        };
    };
};

// | Provides a default implementation for `pred`, given a function that maps
// | integers to values in the `Enum`, and a function that maps values in the
// | `Enum` back to integers. The integer mapping must agree in both directions
// | for this to implement a law-abiding `pred`.
// |
// | If a `BoundedEnum` instance exists for `a`, the `toEnum` and `fromEnum`
// | functions can be used here:
// |
// | ``` purescript
// | pred = defaultPred toEnum fromEnum
// | ```
var defaultPred = function (toEnum$prime) {
    return function (fromEnum$prime) {
        return function (a) {
            return toEnum$prime(fromEnum$prime(a) - 1 | 0);
        };
    };
};

// | Provides a default implementation for `fromEnum`.
// |
// | - Assumes `toEnum 0 = Just bottom`.
// | - Cannot be used in conjuction with `defaultPred`.
// |
// | Runs in `O(n)` where `n` is `fromEnum a`.
var defaultFromEnum = function (dictEnum) {
    var go = function ($copy_i) {
        return function ($copy_x) {
            var $tco_var_i = $copy_i;
            var $tco_done = false;
            var $tco_result;
            function $tco_loop(i, x) {
                var v = pred(dictEnum)(x);
                if (v instanceof Data_Maybe.Just) {
                    $tco_var_i = i + 1 | 0;
                    $copy_x = v.value0;
                    return;
                };
                if (v instanceof Data_Maybe.Nothing) {
                    $tco_done = true;
                    return i;
                };
                throw new Error("Failed pattern match at Data.Enum (line 309, column 5 - line 311, column 19): " + [ v.constructor.name ]);
            };
            while (!$tco_done) {
                $tco_result = $tco_loop($tco_var_i, $copy_x);
            };
            return $tco_result;
        };
    };
    return go(0);
};

// | Provides a default implementation for `cardinality`.
// |
// | Runs in `O(n)` where `n` is `fromEnum top`
var defaultCardinality = function (dictBounded) {
    var bottom2 = Data_Bounded.bottom(dictBounded);
    return function (dictEnum) {
        var go = function ($copy_i) {
            return function ($copy_x) {
                var $tco_var_i = $copy_i;
                var $tco_done = false;
                var $tco_result;
                function $tco_loop(i, x) {
                    var v = succ(dictEnum)(x);
                    if (v instanceof Data_Maybe.Just) {
                        $tco_var_i = i + 1 | 0;
                        $copy_x = v.value0;
                        return;
                    };
                    if (v instanceof Data_Maybe.Nothing) {
                        $tco_done = true;
                        return i;
                    };
                    throw new Error("Failed pattern match at Data.Enum (line 276, column 5 - line 278, column 19): " + [ v.constructor.name ]);
                };
                while (!$tco_done) {
                    $tco_result = $tco_loop($tco_var_i, $copy_x);
                };
                return $tco_result;
            };
        };
        return go(1)(bottom2);
    };
};
var charToEnum = function (v) {
    if (v >= $foreign.toCharCode(bottom1) && v <= $foreign.toCharCode(top)) {
        return new Data_Maybe.Just($foreign.fromCharCode(v));
    };
    return Data_Maybe.Nothing.value;
};
var enumChar = {
    succ: /* #__PURE__ */ defaultSucc(charToEnum)($foreign.toCharCode),
    pred: /* #__PURE__ */ defaultPred(charToEnum)($foreign.toCharCode),
    Ord0: function () {
        return Data_Ord.ordChar;
    }
};
var cardinality = function (dict) {
    return dict.cardinality;
};
var boundedEnumUnit = {
    cardinality: 1,
    toEnum: function (v) {
        if (v === 0) {
            return new Data_Maybe.Just(Data_Unit.unit);
        };
        return Data_Maybe.Nothing.value;
    },
    fromEnum: /* #__PURE__ */ Data_Function["const"](0),
    Bounded0: function () {
        return Data_Bounded.boundedUnit;
    },
    Enum1: function () {
        return enumUnit;
    }
};
var boundedEnumOrdering = {
    cardinality: 3,
    toEnum: function (v) {
        if (v === 0) {
            return new Data_Maybe.Just(Data_Ordering.LT.value);
        };
        if (v === 1) {
            return new Data_Maybe.Just(Data_Ordering.EQ.value);
        };
        if (v === 2) {
            return new Data_Maybe.Just(Data_Ordering.GT.value);
        };
        return Data_Maybe.Nothing.value;
    },
    fromEnum: function (v) {
        if (v instanceof Data_Ordering.LT) {
            return 0;
        };
        if (v instanceof Data_Ordering.EQ) {
            return 1;
        };
        if (v instanceof Data_Ordering.GT) {
            return 2;
        };
        throw new Error("Failed pattern match at Data.Enum (line 137, column 1 - line 145, column 18): " + [ v.constructor.name ]);
    },
    Bounded0: function () {
        return Data_Bounded.boundedOrdering;
    },
    Enum1: function () {
        return enumOrdering;
    }
};
var boundedEnumChar = /* #__PURE__ */ (function () {
    return {
        cardinality: $foreign.toCharCode(top) - $foreign.toCharCode(bottom1) | 0,
        toEnum: charToEnum,
        fromEnum: $foreign.toCharCode,
        Bounded0: function () {
            return Data_Bounded.boundedChar;
        },
        Enum1: function () {
            return enumChar;
        }
    };
})();
var boundedEnumBoolean = {
    cardinality: 2,
    toEnum: function (v) {
        if (v === 0) {
            return new Data_Maybe.Just(false);
        };
        if (v === 1) {
            return new Data_Maybe.Just(true);
        };
        return Data_Maybe.Nothing.value;
    },
    fromEnum: function (v) {
        if (!v) {
            return 0;
        };
        if (v) {
            return 1;
        };
        throw new Error("Failed pattern match at Data.Enum (line 118, column 1 - line 124, column 20): " + [ v.constructor.name ]);
    },
    Bounded0: function () {
        return Data_Bounded.boundedBoolean;
    },
    Enum1: function () {
        return enumBoolean;
    }
};
export {
    succ,
    pred,
    cardinality,
    toEnum,
    fromEnum,
    toEnumWithDefaults,
    Cardinality,
    enumFromTo,
    enumFromThenTo,
    upFrom,
    upFromIncluding,
    downFrom,
    downFromIncluding,
    defaultSucc,
    defaultPred,
    defaultCardinality,
    defaultToEnum,
    defaultFromEnum,
    enumBoolean,
    enumInt,
    enumChar,
    enumUnit,
    enumOrdering,
    enumMaybe,
    enumEither,
    enumTuple,
    boundedEnumBoolean,
    boundedEnumChar,
    boundedEnumUnit,
    boundedEnumOrdering,
    newtypeCardinality,
    eqCardinality,
    ordCardinality,
    showCardinality
};
