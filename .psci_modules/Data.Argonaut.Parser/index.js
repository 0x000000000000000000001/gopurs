import * as $foreign from "./foreign.js";
import * as Data_Either from "../Data.Either/index.js";

// | Parse a JSON string, constructing the `Json` value described by the string.
// | To convert a string into a `Json` string, see `fromString`.
var jsonParser = function (j) {
    return $foreign["_jsonParser"](Data_Either.Left.create, Data_Either.Right.create, j);
};
export {
    jsonParser
};
