import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_String_Common from "../Data.String.Common/index.js";
var fullName = function (mbMod) {
    return function (name) {
        return Data_Maybe.fromMaybe("")(Data_Functor.map(Data_Maybe.functorMaybe)(function (m) {
            return Data_String_Common.joinWith("_")(Data_String_Common.split(".")(m)) + "_";
        })(mbMod)) + name;
    };
};
var main = /* #__PURE__ */ (function () {
    return fullName(new Data_Maybe.Just("Test.TCO"))("deepTailRec");
})();
export {
    fullName,
    main
};
