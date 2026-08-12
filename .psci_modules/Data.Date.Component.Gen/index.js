import * as Control_Monad_Gen_Class from "../Control.Monad.Gen.Class/index.js";
import * as Data_Date_Component from "../Data.Date.Component/index.js";
import * as Data_Enum from "../Data.Enum/index.js";
import * as Data_Enum_Gen from "../Data.Enum.Gen/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
var toEnum = /* #__PURE__ */ Data_Enum.toEnum(Data_Date_Component.boundedEnumYear);

// | Generates a random `Year` in the range 1900-2100, inclusive.
var genYear = function (dictMonadGen) {
    return Data_Functor.map((((dictMonadGen.Monad0()).Bind1()).Apply0()).Functor0())((function () {
        var $15 = Data_Maybe.fromJust();
        return function ($16) {
            return $15(toEnum($16));
        };
    })())(Control_Monad_Gen_Class.chooseInt(dictMonadGen)(1900)(2100));
};

// | Generates a random `Weekday` component.
var genWeekday = function (dictMonadGen) {
    return Data_Enum_Gen.genBoundedEnum(dictMonadGen)(Data_Date_Component.boundedEnumWeekday);
};

// | Generates a random `Month` component.
var genMonth = function (dictMonadGen) {
    return Data_Enum_Gen.genBoundedEnum(dictMonadGen)(Data_Date_Component.boundedEnumMonth);
};

// | Generates a random `Day` component.
var genDay = function (dictMonadGen) {
    return Data_Enum_Gen.genBoundedEnum(dictMonadGen)(Data_Date_Component.boundedEnumDay);
};
export {
    genYear,
    genMonth,
    genDay,
    genWeekday
};
