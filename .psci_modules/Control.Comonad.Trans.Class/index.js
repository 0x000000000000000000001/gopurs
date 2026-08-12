// | This module defines the `ComonadTrans` type class of _comonad transformers_.
import * as Control_Monad_Identity_Trans from "../Control.Monad.Identity.Trans/index.js";
var lower = function (dict) {
    return dict.lower;
};
var comonadTransIdentityT = {
    lower: function (dictComonad) {
        return Control_Monad_Identity_Trans.runIdentityT;
    }
};
export {
    lower,
    comonadTransIdentityT
};
