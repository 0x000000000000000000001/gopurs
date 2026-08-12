// | This module defines the `Traced` comonad.
import * as Control_Comonad_Traced_Class from "../Control.Comonad.Traced.Class/index.js";
import * as Control_Comonad_Traced_Trans from "../Control.Comonad.Traced.Trans/index.js";
import * as Data_Identity from "../Data.Identity/index.js";
import * as Data_Newtype from "../Data.Newtype/index.js";

// | Create a value in context in the `Traced` comonad.
var traced = function ($2) {
    return Control_Comonad_Traced_Trans.TracedT(Data_Identity.Identity($2));
};

// | Unwrap a value in the `Traced` comonad.
var runTraced = function (v) {
    return Data_Newtype.unwrap()(v);
};
export {
    runTraced,
    traced
};
export {
    censor,
    listen,
    listens,
    track,
    tracks
} from "../Control.Comonad.Traced.Class/index.js";
export {
    TracedT,
    runTracedT
} from "../Control.Comonad.Traced.Trans/index.js";
