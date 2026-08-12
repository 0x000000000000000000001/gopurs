// | Some partial helper functions. See the README for more documentation.
import * as $foreign from "./foreign.js";

// | A partial function which crashes on any input with the specified message.
var crashWith = function () {
    return $foreign["_crashWith"];
};

// | A partial function which crashes on any input with a default message.
var crash = function () {
    return crashWith()("Partial.crash: partial function");
};
export {
    crash,
    crashWith
};
