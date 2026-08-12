// | This module defines the `Stall` profunctor
import * as Data_Bifunctor from "../Data.Bifunctor/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";

// | The `Stall` profunctor characterizes an `AffineTraversal`.
var Stall = /* #__PURE__ */ (function () {
    function Stall(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    Stall.create = function (value0) {
        return function (value1) {
            return new Stall(value0, value1);
        };
    };
    return Stall;
})();
var profunctorStall = {
    dimap: function (f) {
        return function (g) {
            return function (v) {
                return new Stall((function () {
                    var $61 = Data_Functor.map(Data_Functor.functorFn)(g);
                    return function ($62) {
                        return $61(v.value0(f($62)));
                    };
                })(), (function () {
                    var $63 = Data_Bifunctor.lmap(Data_Bifunctor.bifunctorEither)(g);
                    return function ($64) {
                        return $63(v.value1(f($64)));
                    };
                })());
            };
        };
    }
};
var strongStall = {
    first: function (v) {
        return new Stall(function (v1) {
            return function (b) {
                return new Data_Tuple.Tuple(v.value0(v1.value0)(b), v1.value1);
            };
        }, function (v1) {
            return Data_Bifunctor.lmap(Data_Bifunctor.bifunctorEither)(function (t) {
                return new Data_Tuple.Tuple(t, v1.value1);
            })(v.value1(v1.value0));
        });
    },
    second: function (v) {
        return new Stall(function (v1) {
            return function (b) {
                return new Data_Tuple.Tuple(v1.value0, v.value0(v1.value1)(b));
            };
        }, function (v1) {
            return Data_Bifunctor.lmap(Data_Bifunctor.bifunctorEither)(Data_Tuple.Tuple.create(v1.value0))(v.value1(v1.value1));
        });
    },
    Profunctor0: function () {
        return profunctorStall;
    }
};
var functorStall = {
    map: function (f) {
        return function (v) {
            return new Stall((function () {
                var $65 = Data_Functor.map(Data_Functor.functorFn)(f);
                return function ($66) {
                    return $65(v.value0($66));
                };
            })(), (function () {
                var $67 = Data_Bifunctor.lmap(Data_Bifunctor.bifunctorEither)(f);
                return function ($68) {
                    return $67(v.value1($68));
                };
            })());
        };
    }
};
var choiceStall = {
    left: function (v) {
        return new Stall(function (v1) {
            if (v1 instanceof Data_Either.Left) {
                return function (b) {
                    return new Data_Either.Left(v.value0(v1.value0)(b));
                };
            };
            if (v1 instanceof Data_Either.Right) {
                return function (v2) {
                    return new Data_Either.Right(v1.value0);
                };
            };
            throw new Error("Failed pattern match at Data.Lens.Internal.Stall (line 36, column 9 - line 38, column 35): " + [ v1.constructor.name ]);
        }, function (v1) {
            if (v1 instanceof Data_Either.Left) {
                return Data_Bifunctor.lmap(Data_Bifunctor.bifunctorEither)(Data_Either.Left.create)(v.value1(v1.value0));
            };
            if (v1 instanceof Data_Either.Right) {
                return new Data_Either.Left(new Data_Either.Right(v1.value0));
            };
            throw new Error("Failed pattern match at Data.Lens.Internal.Stall (line 40, column 9 - line 42, column 36): " + [ v1.constructor.name ]);
        });
    },
    right: function (v) {
        return new Stall(function (v1) {
            if (v1 instanceof Data_Either.Left) {
                return function (v2) {
                    return new Data_Either.Left(v1.value0);
                };
            };
            if (v1 instanceof Data_Either.Right) {
                return function (b) {
                    return new Data_Either.Right(v.value0(v1.value0)(b));
                };
            };
            throw new Error("Failed pattern match at Data.Lens.Internal.Stall (line 47, column 9 - line 49, column 41): " + [ v1.constructor.name ]);
        }, function (v1) {
            if (v1 instanceof Data_Either.Left) {
                return new Data_Either.Left(new Data_Either.Left(v1.value0));
            };
            if (v1 instanceof Data_Either.Right) {
                return Data_Bifunctor.lmap(Data_Bifunctor.bifunctorEither)(Data_Either.Right.create)(v.value1(v1.value0));
            };
            throw new Error("Failed pattern match at Data.Lens.Internal.Stall (line 51, column 9 - line 53, column 38): " + [ v1.constructor.name ]);
        });
    },
    Profunctor0: function () {
        return profunctorStall;
    }
};
export {
    Stall,
    functorStall,
    profunctorStall,
    strongStall,
    choiceStall
};
