// | This module defines the `Reader` monad.
import * as Control_Monad_Reader_Class from "../Control.Monad.Reader.Class/index.js";
import * as Control_Monad_Reader_Trans from "../Control.Monad.Reader.Trans/index.js";
import * as Data_Identity from "../Data.Identity/index.js";
import * as Data_Newtype from "../Data.Newtype/index.js";
var unwrap = /* #__PURE__ */ Data_Newtype.unwrap();

// | Change the type of the context in a `Reader` monad action.
var withReader = Control_Monad_Reader_Trans.withReaderT;

// | Run a computation in the `Reader` monad.
var runReader = function (v) {
    return function ($3) {
        return unwrap(v($3));
    };
};

// | Change the type of the result in a `Reader` monad action.
var mapReader = function (f) {
    return Control_Monad_Reader_Trans.mapReaderT(function ($4) {
        return Data_Identity.Identity(f(unwrap($4)));
    });
};
export {
    runReader,
    mapReader,
    withReader
};
export {
    ask,
    asks,
    local
} from "../Control.Monad.Reader.Class/index.js";
export {
    ReaderT,
    lift,
    mapReaderT,
    runReaderT,
    withReaderT
} from "../Control.Monad.Reader.Trans/index.js";
