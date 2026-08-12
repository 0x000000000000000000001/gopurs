import * as Data_Lens_Iso_Newtype from "../Data.Lens.Iso.Newtype/index.js";
import * as Data_Lens_Prism_Either from "../Data.Lens.Prism.Either/index.js";
var _Newtype = /* #__PURE__ */ Data_Lens_Iso_Newtype["_Newtype"]()();

// | Prism for the `right` of a `Coproduct`.
var _Right = function (dictChoice) {
    var $10 = _Newtype(dictChoice.Profunctor0());
    var $11 = Data_Lens_Prism_Either["_Right"](dictChoice);
    return function ($12) {
        return $10($11($12));
    };
};

// | Prism for the `left` of a `Coproduct`.
var _Left = function (dictChoice) {
    var $13 = _Newtype(dictChoice.Profunctor0());
    var $14 = Data_Lens_Prism_Either["_Left"](dictChoice);
    return function ($15) {
        return $13($14($15));
    };
};
export {
    _Left,
    _Right
};
