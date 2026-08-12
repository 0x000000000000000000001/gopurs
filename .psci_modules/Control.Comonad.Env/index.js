// | This module defines the `Env` comonad.
import * as Control_Comonad_Env_Class from "../Control.Comonad.Env.Class/index.js";
import * as Control_Comonad_Env_Trans from "../Control.Comonad.Env.Trans/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Identity from "../Data.Identity/index.js";
import * as Data_Newtype from "../Data.Newtype/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
var unwrap = /* #__PURE__ */ Data_Newtype.unwrap();

// | Change the environment type in an `Env` computation.
var withEnv = Control_Comonad_Env_Trans.withEnvT;

// | Unwrap a value in the `Env` comonad.
var runEnv = function (v) {
    return Data_Functor.map(Data_Tuple.functorTuple)(unwrap)(v);
};

// | Change the data type in an `Env` computation.
var mapEnv = /* #__PURE__ */ Data_Functor.map(/* #__PURE__ */ Control_Comonad_Env_Trans.functorEnvT(Data_Identity.functorIdentity));

// | Create a value in context in the `Env` comonad.
var env = function (e) {
    return function (a) {
        return new Data_Tuple.Tuple(e, a);
    };
};
export {
    runEnv,
    withEnv,
    mapEnv,
    env
};
export {
    ask,
    asks,
    local
} from "../Control.Comonad.Env.Class/index.js";
export {
    EnvT,
    mapEnvT,
    runEnvT,
    withEnvT
} from "../Control.Comonad.Env.Trans/index.js";
