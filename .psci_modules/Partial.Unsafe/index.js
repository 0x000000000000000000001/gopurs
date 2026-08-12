// | Utilities for working with partial functions.
// | See the README for more documentation.
import * as $foreign from "./foreign.js";
import * as Partial from "../Partial/index.js";

// | Discharge a partiality constraint, unsafely.
var unsafePartial = $foreign["_unsafePartial"];

// | A function which crashes with the specified error message.
var unsafeCrashWith = function (msg) {
    return unsafePartial(function () {
        return Partial.crashWith()(msg);
    });
};
export {
    unsafePartial,
    unsafeCrashWith
};
