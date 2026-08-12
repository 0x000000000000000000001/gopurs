import * as Data_Function from "../Data.Function/index.js";
import * as Data_Lens_Lens from "../Data.Lens.Lens/index.js";
import * as Data_Unit from "../Data.Unit/index.js";

// | There is a `Unit` in everything.
// | ```purescript
// | > view united [1,2,3]
// | unit
// | > over united (\a -> a :: Unit) [1,2,3]
// | [1 2 3]
// | ```
var united = function (dictStrong) {
    return Data_Lens_Lens.lens(Data_Function["const"](Data_Unit.unit))(Data_Function["const"])(dictStrong);
};
export {
    united
};
