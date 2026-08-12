import * as $foreign from "./foreign.js";
import * as Type_Proxy from "../Type.Proxy/index.js";
var reifiableString = {};
var reifiableOrdering = {};
var reifiableInt = {};
var reifiableBoolean = {};

// | Reify a value of type `t` such that it can be consumed by a
// | function constrained by the `Reflectable` type class. For
// | example:
// |
// | ```purs
// | twiceFromType :: forall v. Reflectable v Int => Proxy v -> Int
// | twiceFromType = (_ * 2) <<< reflectType
// |
// | twiceOfTerm :: Int
// | twiceOfTerm = reifyType 21 twiceFromType
// | ```
var reifyType = function () {
    return function (s) {
        return function (f) {
            return $foreign.unsafeCoerce(function (dictReflectable) {
                return f(dictReflectable);
            })({
                reflectType: function (v) {
                    return s;
                }
            })(Type_Proxy["Proxy"].value);
        };
    };
};

// | Reflect a type `v` to its term-level representation.
var reflectType = function (dict) {
    return dict.reflectType;
};
export {
    reflectType,
    reifyType,
    reifiableBoolean,
    reifiableInt,
    reifiableOrdering,
    reifiableString
};
