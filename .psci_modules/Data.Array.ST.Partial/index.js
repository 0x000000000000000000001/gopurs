// | Partial functions for working with mutable arrays using the `ST` effect.
// |
// | This module is particularly helpful when performance is very important.
import * as $foreign from "./foreign.js";
import * as Control_Monad_ST_Uncurried from "../Control.Monad.ST.Uncurried/index.js";

// | Change the value at the specified index in a mutable array.
var poke = function () {
    return Control_Monad_ST_Uncurried.runSTFn3($foreign.pokeImpl);
};

// | Read the value at the specified index in a mutable array.
var peek = function () {
    return Control_Monad_ST_Uncurried.runSTFn2($foreign.peekImpl);
};
export {
    peek,
    poke
};
