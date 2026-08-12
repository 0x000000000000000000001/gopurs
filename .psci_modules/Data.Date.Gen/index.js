import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Monad_Gen_Class from "../Control.Monad.Gen.Class/index.js";
import * as Data_Bounded from "../Data.Bounded/index.js";
import * as Data_Date from "../Data.Date/index.js";
import * as Data_Date_Component from "../Data.Date.Component/index.js";
import * as Data_Date_Component_Gen from "../Data.Date.Component.Gen/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Int from "../Data.Int/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Time_Duration from "../Data.Time.Duration/index.js";
var bottom = /* #__PURE__ */ Data_Bounded.bottom(Data_Date_Component.boundedMonth);
var bottom1 = /* #__PURE__ */ Data_Bounded.bottom(Data_Date_Component.boundedDay);

// | Generates a random `Date` between 1st Jan 1900 and 31st Dec 2100,
// | inclusive.
var genDate = function (dictMonadGen) {
    var Monad0 = dictMonadGen.Monad0();
    var Bind1 = Monad0.Bind1();
    var Functor0 = ((Monad0.Bind1()).Apply0()).Functor0();
    var pure = Control_Applicative.pure(Monad0.Applicative0());
    return Control_Bind.bind(Bind1)(Data_Date_Component_Gen.genYear(dictMonadGen))(function (year) {
        var maxDays = (function () {
            var $11 = Data_Date.isLeapYear(year);
            if ($11) {
                return 365;
            };
            return 364;
        })();
        return Control_Bind.bind(Bind1)(Data_Functor.map(Functor0)(function ($12) {
            return Data_Time_Duration.Days(Data_Int.toNumber($12));
        })(Control_Monad_Gen_Class.chooseInt(dictMonadGen)(0)(maxDays)))(function (days) {
            return pure(Data_Maybe.fromJust()(Control_Bind.bind(Data_Maybe.bindMaybe)(Data_Date.exactDate(year)(bottom)(bottom1))(function (janFirst) {
                return Data_Date.adjust(days)(janFirst);
            })));
        });
    });
};
export {
    genDate
};
export {
    genDay,
    genMonth,
    genWeekday,
    genYear
} from "../Data.Date.Component.Gen/index.js";
