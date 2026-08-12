import * as $foreign from "./foreign.js";
import * as Data_Ordering from "../Data.Ordering/index.js";

// | Returns `true` if the given string is empty.
// |
// | ```purescript
// | null "" == true
// | null "Hi" == false
// | ```
var $$null = function (s) {
    return s === "";
};

// | Compare two strings in a locale-aware fashion. This is in contrast to
// | the `Ord` instance on `String` which treats strings as arrays of code
// | units:
// |
// | ```purescript
// | "ä" `localeCompare` "b" == LT
// | "ä" `compare` "b" == GT
// | ```
var localeCompare = /* #__PURE__ */ (function () {
    return $foreign["_localeCompare"](Data_Ordering.LT.value)(Data_Ordering.EQ.value)(Data_Ordering.GT.value);
})();
export {
    replace,
    replaceAll,
    split,
    toLower,
    toUpper,
    trim,
    joinWith
} from "./foreign.js";
export {
    $$null as null,
    localeCompare
};
