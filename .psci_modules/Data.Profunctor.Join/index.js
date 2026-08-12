import * as Control_Category from "../Control.Category/index.js";
import * as Control_Semigroupoid from "../Control.Semigroupoid/index.js";
import * as Data_Profunctor from "../Data.Profunctor/index.js";
import * as Data_Show from "../Data.Show/index.js";
var Join = function (x) {
    return x;
};
var showJoin = function (dictShow) {
    return {
        show: function (v) {
            return "(Join " + (Data_Show.show(dictShow)(v) + ")");
        }
    };
};
var semigroupJoin = function (dictSemigroupoid) {
    return {
        append: function (v) {
            return function (v1) {
                return Control_Semigroupoid.compose(dictSemigroupoid)(v)(v1);
            };
        }
    };
};
var ordJoin = function (dictOrd) {
    return dictOrd;
};
var newtypeJoin = {
    Coercible0: function () {
        return undefined;
    }
};
var monoidJoin = function (dictCategory) {
    var semigroupJoin1 = semigroupJoin(dictCategory.Semigroupoid0());
    return {
        mempty: Control_Category.identity(dictCategory),
        Semigroup0: function () {
            return semigroupJoin1;
        }
    };
};
var invariantJoin = function (dictProfunctor) {
    return {
        imap: function (f) {
            return function (g) {
                return function (v) {
                    return Data_Profunctor.dimap(dictProfunctor)(g)(f)(v);
                };
            };
        }
    };
};
var eqJoin = function (dictEq) {
    return dictEq;
};
export {
    Join,
    newtypeJoin,
    eqJoin,
    ordJoin,
    showJoin,
    semigroupJoin,
    monoidJoin,
    invariantJoin
};
