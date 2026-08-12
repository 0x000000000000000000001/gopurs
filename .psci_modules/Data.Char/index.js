// | A type and functions for single characters.
import * as Data_Enum from "../Data.Enum/index.js";

// | Returns the numeric Unicode value of the character.
var toCharCode = /* #__PURE__ */ Data_Enum.fromEnum(Data_Enum.boundedEnumChar);

// | Constructs a character from the given Unicode numeric value.
var fromCharCode = /* #__PURE__ */ Data_Enum.toEnum(Data_Enum.boundedEnumChar);
export {
    toCharCode,
    fromCharCode
};
