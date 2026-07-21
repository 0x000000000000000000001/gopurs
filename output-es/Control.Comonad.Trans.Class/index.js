// | This module defines the `ComonadTrans` type class of _comonad transformers_.
import * as Control$dMonad$dIdentity$dTrans from "../Control.Monad.Identity.Trans/index.js";
const lower = dict => dict.lower;
const comonadTransIdentityT = {lower: dictComonad => Control$dMonad$dIdentity$dTrans.runIdentityT};
export {comonadTransIdentityT, lower};
