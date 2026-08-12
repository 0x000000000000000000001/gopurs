import * as Data_Enum_Gen from "../Data.Enum.Gen/index.js";
import * as Data_Time_Component from "../Data.Time.Component/index.js";

// | Generates a random `Second` component.
var genSecond = function (dictMonadGen) {
    return Data_Enum_Gen.genBoundedEnum(dictMonadGen)(Data_Time_Component.boundedEnumSecond);
};

// | Generates a random `Minute` component.
var genMinute = function (dictMonadGen) {
    return Data_Enum_Gen.genBoundedEnum(dictMonadGen)(Data_Time_Component.boundedEnumMinute);
};

// | Generates a random `Millisecond` component.
var genMillisecond = function (dictMonadGen) {
    return Data_Enum_Gen.genBoundedEnum(dictMonadGen)(Data_Time_Component.boundedEnumMillisecond);
};

// | Generates a random `Hour` component.
var genHour = function (dictMonadGen) {
    return Data_Enum_Gen.genBoundedEnum(dictMonadGen)(Data_Time_Component.boundedEnumHour);
};
export {
    genHour,
    genMinute,
    genSecond,
    genMillisecond
};
