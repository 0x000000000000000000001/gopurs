import * as Data_Argonaut_Core from "../Data.Argonaut.Core/index.js";
import * as Data_Lens_Fold from "../Data.Lens.Fold/index.js";
var _JsonString = function (dictWander) {
    var $13 = Data_Lens_Fold.filtered(dictWander.Choice1())(Data_Argonaut_Core.isString);
    return function ($14) {
        return $13($14);
    };
};
var _JsonObject = function (dictWander) {
    var $15 = Data_Lens_Fold.filtered(dictWander.Choice1())(Data_Argonaut_Core.isObject);
    return function ($16) {
        return $15($16);
    };
};
var _JsonNumber = function (dictWander) {
    var $17 = Data_Lens_Fold.filtered(dictWander.Choice1())(Data_Argonaut_Core.isNumber);
    return function ($18) {
        return $17($18);
    };
};
var _JsonNull = function (dictWander) {
    var $19 = Data_Lens_Fold.filtered(dictWander.Choice1())(Data_Argonaut_Core.isNull);
    return function ($20) {
        return $19($20);
    };
};
var _JsonBoolean = function (dictWander) {
    var $21 = Data_Lens_Fold.filtered(dictWander.Choice1())(Data_Argonaut_Core.isBoolean);
    return function ($22) {
        return $21($22);
    };
};
var _JsonArray = function (dictWander) {
    var $23 = Data_Lens_Fold.filtered(dictWander.Choice1())(Data_Argonaut_Core.isArray);
    return function ($24) {
        return $23($24);
    };
};
export {
    _JsonNull,
    _JsonBoolean,
    _JsonNumber,
    _JsonString,
    _JsonArray,
    _JsonObject
};
