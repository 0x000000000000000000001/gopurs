import * as Data_Bounded from "../Data.Bounded/index.js";
import * as Data_Generic_Rep from "../Data.Generic.Rep/index.js";
var genericTopNoArguments = /* #__PURE__ */ (function () {
    return {
        "genericTop'": Data_Generic_Rep.NoArguments.value
    };
})();
var genericTopArgument = function (dictBounded) {
    return {
        "genericTop'": Data_Bounded.top(dictBounded)
    };
};
var genericTop$prime = function (dict) {
    return dict["genericTop'"];
};
var genericTopConstructor = function (dictGenericTop) {
    return {
        "genericTop'": genericTop$prime(dictGenericTop)
    };
};
var genericTopProduct = function (dictGenericTop) {
    return function (dictGenericTop1) {
        return {
            "genericTop'": new Data_Generic_Rep.Product(genericTop$prime(dictGenericTop), genericTop$prime(dictGenericTop1))
        };
    };
};
var genericTopSum = function (dictGenericTop) {
    return {
        "genericTop'": new Data_Generic_Rep.Inr(genericTop$prime(dictGenericTop))
    };
};

// | A `Generic` implementation of the `top` member from the `Bounded` type class.
var genericTop = function (dictGeneric) {
    return function (dictGenericTop) {
        return Data_Generic_Rep.to(dictGeneric)(genericTop$prime(dictGenericTop));
    };
};
var genericBottomNoArguments = /* #__PURE__ */ (function () {
    return {
        "genericBottom'": Data_Generic_Rep.NoArguments.value
    };
})();
var genericBottomArgument = function (dictBounded) {
    return {
        "genericBottom'": Data_Bounded.bottom(dictBounded)
    };
};
var genericBottom$prime = function (dict) {
    return dict["genericBottom'"];
};
var genericBottomConstructor = function (dictGenericBottom) {
    return {
        "genericBottom'": genericBottom$prime(dictGenericBottom)
    };
};
var genericBottomProduct = function (dictGenericBottom) {
    return function (dictGenericBottom1) {
        return {
            "genericBottom'": new Data_Generic_Rep.Product(genericBottom$prime(dictGenericBottom), genericBottom$prime(dictGenericBottom1))
        };
    };
};
var genericBottomSum = function (dictGenericBottom) {
    return {
        "genericBottom'": new Data_Generic_Rep.Inl(genericBottom$prime(dictGenericBottom))
    };
};

// | A `Generic` implementation of the `bottom` member from the `Bounded` type class.
var genericBottom = function (dictGeneric) {
    return function (dictGenericBottom) {
        return Data_Generic_Rep.to(dictGeneric)(genericBottom$prime(dictGenericBottom));
    };
};
export {
    genericBottom$prime,
    genericBottom,
    genericTop$prime,
    genericTop,
    genericBottomNoArguments,
    genericBottomArgument,
    genericBottomSum,
    genericBottomProduct,
    genericBottomConstructor,
    genericTopNoArguments,
    genericTopArgument,
    genericTopSum,
    genericTopProduct,
    genericTopConstructor
};
