import * as Control_Apply from "../Control.Apply/index.js";
import * as Data_Date_Gen from "../Data.Date.Gen/index.js";
import * as Data_DateTime from "../Data.DateTime/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Time_Gen from "../Data.Time.Gen/index.js";

// | Generates a random `DateTime` between 1st Jan 1900 00:00:00 and
// | 31st Dec 2100 23:59:59, inclusive.
var genDateTime = function (dictMonadGen) {
    var Bind1 = (dictMonadGen.Monad0()).Bind1();
    return Control_Apply.apply(Bind1.Apply0())(Data_Functor.map((Bind1.Apply0()).Functor0())(Data_DateTime.DateTime.create)(Data_Date_Gen.genDate(dictMonadGen)))(Data_Time_Gen.genTime(dictMonadGen));
};
export {
    genDateTime
};
export {
    genDate,
    genDay,
    genMonth,
    genWeekday,
    genYear
} from "../Data.Date.Gen/index.js";
export {
    genHour,
    genMillisecond,
    genMinute,
    genSecond,
    genTime
} from "../Data.Time.Gen/index.js";
