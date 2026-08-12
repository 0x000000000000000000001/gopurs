// | This module defines types for STf uncurried functions, as well as
// | functions for converting back and forth between them.
// |
// | The general naming scheme for functions and types in this module is as
// | follows:
// |
// | * `STFn{N}` means, an uncurried function which accepts N arguments and
// |   performs some STs. The first N arguments are the actual function's
// |   argument. The last type argument is the return type.
// | * `runSTFn{N}` takes an `STFn` of N arguments, and converts it into
// |   the normal PureScript form: a curried function which returns an ST
// |   action.
// | * `mkSTFn{N}` is the inverse of `runSTFn{N}`. It can be useful for
// |   callbacks.
// |
import * as $foreign from "./foreign.js";

export {
    mkSTFn1,
    mkSTFn2,
    mkSTFn3,
    mkSTFn4,
    mkSTFn5,
    mkSTFn6,
    mkSTFn7,
    mkSTFn8,
    mkSTFn9,
    mkSTFn10,
    runSTFn1,
    runSTFn2,
    runSTFn3,
    runSTFn4,
    runSTFn5,
    runSTFn6,
    runSTFn7,
    runSTFn8,
    runSTFn9,
    runSTFn10
} from "./foreign.js";
