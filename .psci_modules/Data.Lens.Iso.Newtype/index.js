import * as Data_Lens_Iso from "../Data.Lens.Iso/index.js";
var coerced = /* #__PURE__ */ Data_Lens_Iso.coerced()();

// | An Iso between a newtype and its inner type. This is a specialization of
// | `coerced` restricted to newtypes. Supports switching between different
// | types that have instances of the Newtype type class.
// | If you don't need to change types, you may have a better experience with
// | type inference if you use `simple _Newtype`.
var _Newtype = function () {
    return function () {
        return function (dictProfunctor) {
            return coerced(dictProfunctor);
        };
    };
};
var _Newtype1 = /* #__PURE__ */ _Newtype()();

// | A variant of `_Newtype` which takes the constructor as an argument
// | and infers its inverse.
// |
// | This is useful as an aid to type inference in certain situations.
var unto = function () {
    return function (v) {
        return function (dictProfunctor) {
            return _Newtype1(dictProfunctor);
        };
    };
};
export {
    _Newtype,
    unto
};
