import * as Control_Apply from "../Control.Apply/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Time from "../Data.Time/index.js";
import * as Data_Time_Component_Gen from "../Data.Time.Component.Gen/index.js";

// | Generates a random `Time` between 00:00:00 and 23:59:59, inclusive.
var genTime = function (dictMonadGen) {
    var Bind1 = (dictMonadGen.Monad0()).Bind1();
    var Apply0 = Bind1.Apply0();
    return Control_Apply.apply(Apply0)(Control_Apply.apply(Apply0)(Control_Apply.apply(Apply0)(Data_Functor.map((Bind1.Apply0()).Functor0())(Data_Time.Time.create)(Data_Time_Component_Gen.genHour(dictMonadGen)))(Data_Time_Component_Gen.genMinute(dictMonadGen)))(Data_Time_Component_Gen.genSecond(dictMonadGen)))(Data_Time_Component_Gen.genMillisecond(dictMonadGen));
};
export {
    genTime
};
export {
    genHour,
    genMillisecond,
    genMinute,
    genSecond
} from "../Data.Time.Component.Gen/index.js";
