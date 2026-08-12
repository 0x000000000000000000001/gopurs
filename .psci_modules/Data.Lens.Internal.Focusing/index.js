// | This module defines the `Focusing` functor
import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";

// | The functor used to zoom into `StateT`.
var Focusing = function (x) {
    return x;
};
var newtypeFocusing = {
    Coercible0: function () {
        return undefined;
    }
};
var focusingFunctor = function (dictFunctor) {
    return {
        map: function (f) {
            return function (v) {
                return Data_Functor.map(dictFunctor)(Data_Functor.map(Data_Tuple.functorTuple)(f))(v);
            };
        }
    };
};
var focusingApply = function (dictApply) {
    var Functor0 = dictApply.Functor0();
    var focusingFunctor1 = focusingFunctor(dictApply.Functor0());
    return function (dictSemigroup) {
        var apply = Control_Apply.apply(Data_Tuple.applyTuple(dictSemigroup));
        return {
            apply: function (v) {
                return function (v1) {
                    return Control_Apply.apply(dictApply)(Data_Functor.map(Functor0)(apply)(v))(v1);
                };
            },
            Functor0: function () {
                return focusingFunctor1;
            }
        };
    };
};
var focusingApplicative = function (dictApplicative) {
    var pure = Control_Applicative.pure(dictApplicative);
    var focusingApply1 = focusingApply(dictApplicative.Apply0());
    return function (dictMonoid) {
        var focusingApply2 = focusingApply1(dictMonoid.Semigroup0());
        return {
            pure: (function () {
                var $24 = Control_Applicative.pure(Data_Tuple.applicativeTuple(dictMonoid));
                return function ($25) {
                    return Focusing(pure($24($25)));
                };
            })(),
            Apply0: function () {
                return focusingApply2;
            }
        };
    };
};
export {
    Focusing,
    newtypeFocusing,
    focusingFunctor,
    focusingApply,
    focusingApplicative
};
