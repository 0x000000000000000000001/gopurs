import * as Data_Generic_Rep from "../Data.Generic.Rep/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
var genericMonoidNoArguments = /* #__PURE__ */ (function () {
    return {
        "genericMempty'": Data_Generic_Rep.NoArguments.value
    };
})();
var genericMonoidArgument = function (dictMonoid) {
    return {
        "genericMempty'": Data_Monoid.mempty(dictMonoid)
    };
};
var genericMempty$prime = function (dict) {
    return dict["genericMempty'"];
};
var genericMonoidConstructor = function (dictGenericMonoid) {
    return {
        "genericMempty'": genericMempty$prime(dictGenericMonoid)
    };
};
var genericMonoidProduct = function (dictGenericMonoid) {
    return function (dictGenericMonoid1) {
        return {
            "genericMempty'": new Data_Generic_Rep.Product(genericMempty$prime(dictGenericMonoid), genericMempty$prime(dictGenericMonoid1))
        };
    };
};

// | A `Generic` implementation of the `mempty` member from the `Monoid` type class.
var genericMempty = function (dictGeneric) {
    return function (dictGenericMonoid) {
        return Data_Generic_Rep.to(dictGeneric)(genericMempty$prime(dictGenericMonoid));
    };
};
export {
    genericMempty$prime,
    genericMempty,
    genericMonoidNoArguments,
    genericMonoidProduct,
    genericMonoidConstructor,
    genericMonoidArgument
};
