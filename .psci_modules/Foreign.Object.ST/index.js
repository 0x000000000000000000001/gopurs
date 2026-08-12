// | Helper functions for working with mutable objects using the `ST` effect.
// |
// | This module can be used when performance is important and mutation is a
// | local effect.
import * as $foreign from "./foreign.js";
import * as Data_Maybe from "../Data.Maybe/index.js";

// | Get the value for a key in a mutable object
var peek = /* #__PURE__ */ (function () {
    return $foreign.peekImpl(Data_Maybe.Just.create)(Data_Maybe.Nothing.value);
})();
export {
    new,
    poke,
    delete
} from "./foreign.js";
export {
    peek
};
