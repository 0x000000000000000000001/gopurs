import * as Control_Biapplicative from "../Control.Biapplicative/index.js";
import * as Control_Biapply from "../Control.Biapply/index.js";
import * as Control_Category from "../Control.Category/index.js";
import * as Control_Semigroupoid from "../Control.Semigroupoid/index.js";
import * as Data_Bifunctor from "../Data.Bifunctor/index.js";
import * as Data_Profunctor from "../Data.Profunctor/index.js";
import * as Data_Show from "../Data.Show/index.js";
var Flip = function (x) {
    return x;
};
var showFlip = function (dictShow) {
    return {
        show: function (v) {
            return "(Flip " + (Data_Show.show(dictShow)(v) + ")");
        }
    };
};
var semigroupoidFlip = function (dictSemigroupoid) {
    return {
        compose: function (v) {
            return function (v1) {
                return Control_Semigroupoid.compose(dictSemigroupoid)(v1)(v);
            };
        }
    };
};
var ordFlip = function (dictOrd) {
    return dictOrd;
};
var newtypeFlip = {
    Coercible0: function () {
        return undefined;
    }
};
var functorFlip = function (dictBifunctor) {
    return {
        map: function (f) {
            return function (v) {
                return Data_Bifunctor.lmap(dictBifunctor)(f)(v);
            };
        }
    };
};
var eqFlip = function (dictEq) {
    return dictEq;
};
var contravariantFlip = function (dictProfunctor) {
    return {
        cmap: function (f) {
            return function (v) {
                return Data_Profunctor.lcmap(dictProfunctor)(f)(v);
            };
        }
    };
};
var categoryFlip = function (dictCategory) {
    var semigroupoidFlip1 = semigroupoidFlip(dictCategory.Semigroupoid0());
    return {
        identity: Control_Category.identity(dictCategory),
        Semigroupoid0: function () {
            return semigroupoidFlip1;
        }
    };
};
var bifunctorFlip = function (dictBifunctor) {
    return {
        bimap: function (f) {
            return function (g) {
                return function (v) {
                    return Data_Bifunctor.bimap(dictBifunctor)(g)(f)(v);
                };
            };
        }
    };
};
var biapplyFlip = function (dictBiapply) {
    var bifunctorFlip1 = bifunctorFlip(dictBiapply.Bifunctor0());
    return {
        biapply: function (v) {
            return function (v1) {
                return Control_Biapply.biapply(dictBiapply)(v)(v1);
            };
        },
        Bifunctor0: function () {
            return bifunctorFlip1;
        }
    };
};
var biapplicativeFlip = function (dictBiapplicative) {
    var biapplyFlip1 = biapplyFlip(dictBiapplicative.Biapply0());
    return {
        bipure: function (a) {
            return function (b) {
                return Control_Biapplicative.bipure(dictBiapplicative)(b)(a);
            };
        },
        Biapply0: function () {
            return biapplyFlip1;
        }
    };
};
export {
    Flip,
    newtypeFlip,
    eqFlip,
    ordFlip,
    showFlip,
    functorFlip,
    bifunctorFlip,
    biapplyFlip,
    biapplicativeFlip,
    contravariantFlip,
    semigroupoidFlip,
    categoryFlip
};
