import * as Control_Category from "../Control.Category/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Identity from "../Data.Identity/index.js";
import * as Data_Newtype from "../Data.Newtype/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Data_Unit from "../Data.Unit/index.js";
import * as Type_Equality from "../Type.Equality/index.js";
var unwrap = /* #__PURE__ */ Data_Newtype.unwrap();
var unwrap1 = /* #__PURE__ */ Data_Newtype.unwrap();
var identity = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var distributiveIdentity = {
    distribute: function (dictFunctor) {
        var $24 = Data_Functor.map(dictFunctor)(unwrap);
        return function ($25) {
            return Data_Identity.Identity($24($25));
        };
    },
    collect: function (dictFunctor) {
        return function (f) {
            var $26 = Data_Functor.map(dictFunctor)(function ($28) {
                return unwrap1(f($28));
            });
            return function ($27) {
                return Data_Identity.Identity($26($27));
            };
        };
    },
    Functor0: function () {
        return Data_Identity.functorIdentity;
    }
};
var distribute = function (dict) {
    return dict.distribute;
};
var distributiveFunction = {
    distribute: function (dictFunctor) {
        return function (a) {
            return function (e) {
                return Data_Functor.map(dictFunctor)(function (v) {
                    return v(e);
                })(a);
            };
        };
    },
    collect: function (dictFunctor) {
        return function (f) {
            var $29 = distribute(distributiveFunction)(dictFunctor);
            var $30 = Data_Functor.map(dictFunctor)(f);
            return function ($31) {
                return $29($30($31));
            };
        };
    },
    Functor0: function () {
        return Data_Functor.functorFn;
    }
};

// | Zip an arbitrary collection of containers and summarize the results
var cotraverse = function (dictDistributive) {
    var Functor0 = dictDistributive.Functor0();
    var distribute1 = distribute(dictDistributive);
    return function (dictFunctor) {
        var distribute2 = distribute1(dictFunctor);
        return function (f) {
            var $32 = Data_Functor.map(Functor0)(f);
            return function ($33) {
                return $32(distribute2($33));
            };
        };
    };
};

// | A default implementation of `collect`, based on `distribute`.
var collectDefault = function (dictDistributive) {
    var distribute1 = distribute(dictDistributive);
    return function (dictFunctor) {
        var distribute2 = distribute1(dictFunctor);
        return function (f) {
            var $34 = Data_Functor.map(dictFunctor)(f);
            return function ($35) {
                return distribute2($34($35));
            };
        };
    };
};
var distributiveTuple = function (dictTypeEquals) {
    return {
        collect: function (dictFunctor) {
            return collectDefault(distributiveTuple(dictTypeEquals))(dictFunctor);
        },
        distribute: function (dictFunctor) {
            var $36 = Data_Tuple.Tuple.create(Type_Equality.from(dictTypeEquals)(Data_Unit.unit));
            var $37 = Data_Functor.map(dictFunctor)(Data_Tuple.snd);
            return function ($38) {
                return $36($37($38));
            };
        },
        Functor0: function () {
            return Data_Tuple.functorTuple;
        }
    };
};
var collect = function (dict) {
    return dict.collect;
};

// | A default implementation of `distribute`, based on `collect`.
var distributeDefault = function (dictDistributive) {
    return function (dictFunctor) {
        return collect(dictDistributive)(dictFunctor)(identity);
    };
};
export {
    collect,
    distribute,
    distributeDefault,
    collectDefault,
    cotraverse,
    distributiveIdentity,
    distributiveFunction,
    distributiveTuple
};
