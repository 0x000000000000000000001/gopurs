import * as Control_Apply from "../Control.Apply/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Unit from "../Data.Unit/index.js";
import * as Type_Proxy from "../Type.Proxy/index.js";
var pure = function (dict) {
    return dict.pure;
};

// | Perform an applicative action unless a condition is true.
var unless = function (dictApplicative) {
    return function (v) {
        return function (v1) {
            if (!v) {
                return v1;
            };
            if (v) {
                return pure(dictApplicative)(Data_Unit.unit);
            };
            throw new Error("Failed pattern match at Control.Applicative (line 68, column 1 - line 68, column 65): " + [ v.constructor.name, v1.constructor.name ]);
        };
    };
};

// | Perform an applicative action when a condition is true.
var when = function (dictApplicative) {
    return function (v) {
        return function (v1) {
            if (v) {
                return v1;
            };
            if (!v) {
                return pure(dictApplicative)(Data_Unit.unit);
            };
            throw new Error("Failed pattern match at Control.Applicative (line 63, column 1 - line 63, column 63): " + [ v.constructor.name, v1.constructor.name ]);
        };
    };
};

// | `liftA1` provides a default implementation of `(<$>)` for any
// | [`Applicative`](#applicative) functor, without using `(<$>)` as provided
// | by the [`Functor`](#functor)-[`Applicative`](#applicative) superclass
// | relationship.
// |
// | `liftA1` can therefore be used to write [`Functor`](#functor) instances
// | as follows:
// |
// | ```purescript
// | instance functorF :: Functor F where
// |   map = liftA1
// | ```
var liftA1 = function (dictApplicative) {
    var Apply0 = dictApplicative.Apply0();
    return function (f) {
        return function (a) {
            return Control_Apply.apply(Apply0)(pure(dictApplicative)(f))(a);
        };
    };
};
var applicativeProxy = {
    pure: function (v) {
        return Type_Proxy["Proxy"].value;
    },
    Apply0: function () {
        return Control_Apply.applyProxy;
    }
};
var applicativeFn = {
    pure: function (x) {
        return function (v) {
            return x;
        };
    },
    Apply0: function () {
        return Control_Apply.applyFn;
    }
};
var applicativeArray = {
    pure: function (x) {
        return [ x ];
    },
    Apply0: function () {
        return Control_Apply.applyArray;
    }
};
export {
    pure,
    liftA1,
    unless,
    when,
    applicativeFn,
    applicativeArray,
    applicativeProxy
};
export {
    apply
} from "../Control.Apply/index.js";
export {
    map,
    void
} from "../Data.Functor/index.js";
