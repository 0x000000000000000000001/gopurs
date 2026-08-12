import * as Data_Function from "../Data.Function/index.js";
import * as Data_Lens_Lens from "../Data.Lens.Lens/index.js";
import * as Data_Void from "../Data.Void/index.js";

// | There is everything in `Void`.
var devoid = function (dictStrong) {
    return Data_Lens_Lens.lens(Data_Void.absurd)(Data_Function["const"])(dictStrong);
};
export {
    devoid
};
