import * as Data_Lens_Iso_Newtype from "../Data.Lens.Iso.Newtype/index.js";
import * as Data_Lens_Lens_Tuple from "../Data.Lens.Lens.Tuple/index.js";
var _Newtype = /* #__PURE__ */ Data_Lens_Iso_Newtype["_Newtype"]()();

// | Lens for the second component of a `Product`.
var _2 = function (dictStrong) {
    var $10 = _Newtype(dictStrong.Profunctor0());
    var $11 = Data_Lens_Lens_Tuple["_2"](dictStrong);
    return function ($12) {
        return $10($11($12));
    };
};

// | Lens for the first component of a `Product`.
var _1 = function (dictStrong) {
    var $13 = _Newtype(dictStrong.Profunctor0());
    var $14 = Data_Lens_Lens_Tuple["_1"](dictStrong);
    return function ($15) {
        return $13($14($15));
    };
};
export {
    _1,
    _2
};
