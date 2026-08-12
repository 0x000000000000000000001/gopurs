// | This module defines the `Indexed` profunctor.
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Lens_Internal_Wander from "../Data.Lens.Internal.Wander/index.js";
import * as Data_Profunctor from "../Data.Profunctor/index.js";
import * as Data_Profunctor_Choice from "../Data.Profunctor.Choice/index.js";
import * as Data_Profunctor_Strong from "../Data.Profunctor.Strong/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
var Indexed = function (x) {
    return x;
};
var profunctorIndexed = function (dictProfunctor) {
    return {
        dimap: function (f) {
            return function (g) {
                return function (v) {
                    return Data_Profunctor.dimap(dictProfunctor)(Data_Profunctor_Strong.second(Data_Profunctor_Strong.strongFn)(f))(g)(v);
                };
            };
        }
    };
};
var strongIndexed = function (dictStrong) {
    var Profunctor0 = dictStrong.Profunctor0();
    var profunctorIndexed1 = profunctorIndexed(dictStrong.Profunctor0());
    return {
        first: function (v) {
            return Data_Profunctor.lcmap(Profunctor0)(function (v1) {
                return new Data_Tuple.Tuple(new Data_Tuple.Tuple(v1.value0, v1.value1.value0), v1.value1.value1);
            })(Data_Profunctor_Strong.first(dictStrong)(v));
        },
        second: function (v) {
            return Data_Profunctor.lcmap(Profunctor0)(function (v1) {
                return new Data_Tuple.Tuple(v1.value1.value0, new Data_Tuple.Tuple(v1.value0, v1.value1.value1));
            })(Data_Profunctor_Strong.second(dictStrong)(v));
        },
        Profunctor0: function () {
            return profunctorIndexed1;
        }
    };
};
var newtypeIndexed = {
    Coercible0: function () {
        return undefined;
    }
};
var choiceIndexed = function (dictChoice) {
    var Profunctor0 = dictChoice.Profunctor0();
    var profunctorIndexed1 = profunctorIndexed(dictChoice.Profunctor0());
    return {
        left: function (v) {
            return Data_Profunctor.lcmap(Profunctor0)(function (v1) {
                return Data_Either.either((function () {
                    var $54 = Data_Tuple.Tuple.create(v1.value0);
                    return function ($55) {
                        return Data_Either.Left.create($54($55));
                    };
                })())(Data_Either.Right.create)(v1.value1);
            })(Data_Profunctor_Choice.left(dictChoice)(v));
        },
        right: function (v) {
            return Data_Profunctor.lcmap(Profunctor0)(function (v1) {
                return Data_Either.either(Data_Either.Left.create)((function () {
                    var $56 = Data_Tuple.Tuple.create(v1.value0);
                    return function ($57) {
                        return Data_Either.Right.create($56($57));
                    };
                })())(v1.value1);
            })(Data_Profunctor_Choice.right(dictChoice)(v));
        },
        Profunctor0: function () {
            return profunctorIndexed1;
        }
    };
};
var wanderIndexed = function (dictWander) {
    var strongIndexed1 = strongIndexed(dictWander.Strong0());
    var choiceIndexed1 = choiceIndexed(dictWander.Choice1());
    return {
        wander: function (trav) {
            return function (v) {
                return Data_Lens_Internal_Wander.wander(dictWander)(function (dictApplicative) {
                    return function (ia2fb) {
                        return function (v1) {
                            return trav(dictApplicative)((function () {
                                var $58 = Data_Tuple.Tuple.create(v1.value0);
                                return function ($59) {
                                    return ia2fb($58($59));
                                };
                            })())(v1.value1);
                        };
                    };
                })(v);
            };
        },
        Strong0: function () {
            return strongIndexed1;
        },
        Choice1: function () {
            return choiceIndexed1;
        }
    };
};
export {
    Indexed,
    newtypeIndexed,
    profunctorIndexed,
    strongIndexed,
    choiceIndexed,
    wanderIndexed
};
