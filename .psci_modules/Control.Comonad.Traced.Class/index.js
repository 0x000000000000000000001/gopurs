// | This module defines the `ComonadTraced` type class and its instances.
import * as Control_Comonad from "../Control.Comonad/index.js";
import * as Control_Comonad_Env_Trans from "../Control.Comonad.Env.Trans/index.js";
import * as Control_Comonad_Store_Trans from "../Control.Comonad.Store.Trans/index.js";
import * as Control_Comonad_Traced_Trans from "../Control.Comonad.Traced.Trans/index.js";
import * as Control_Comonad_Trans_Class from "../Control.Comonad.Trans.Class/index.js";
import * as Control_Monad_Identity_Trans from "../Control.Monad.Identity.Trans/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
var track = function (dict) {
    return dict.track;
};

// | Extracts a value at a relative position which depends on the current value.
var tracks = function (dictComonadTraced) {
    var Comonad0 = dictComonadTraced.Comonad0();
    return function (f) {
        return function (w) {
            return track(dictComonadTraced)(f(Control_Comonad.extract(Comonad0)(w)))(w);
        };
    };
};
var lowerTrack = function (dictComonadTrans) {
    var lower = Control_Comonad_Trans_Class.lower(dictComonadTrans);
    return function (dictComonadTraced) {
        var lower1 = lower(dictComonadTraced.Comonad0());
        return function (m) {
            var $42 = track(dictComonadTraced)(m);
            return function ($43) {
                return $42(lower1($43));
            };
        };
    };
};
var lowerTrack1 = /* #__PURE__ */ lowerTrack(Control_Comonad_Store_Trans.comonadTransStoreT);
var lowerTrack2 = /* #__PURE__ */ lowerTrack(Control_Comonad_Trans_Class.comonadTransIdentityT);
var lowerTrack3 = /* #__PURE__ */ lowerTrack(Control_Comonad_Env_Trans.comonadTransEnvT);

// | Get a value which depends on the current position.
var listens = function (dictFunctor) {
    return function (f) {
        return function (v) {
            return Data_Functor.map(dictFunctor)(function (g) {
                return function (t) {
                    return new Data_Tuple.Tuple(g(t), f(t));
                };
            })(v);
        };
    };
};

// | Get the current position.
var listen = function (dictFunctor) {
    return function (v) {
        return Data_Functor.map(dictFunctor)(function (f) {
            return function (t) {
                return new Data_Tuple.Tuple(f(t), t);
            };
        })(v);
    };
};
var comonadTracedTracedT = function (dictComonad) {
    var comonadTracedT = Control_Comonad_Traced_Trans.comonadTracedT(dictComonad);
    return function (dictMonoid) {
        var comonadTracedT1 = comonadTracedT(dictMonoid);
        return {
            track: function (t) {
                return function (v) {
                    return Control_Comonad.extract(dictComonad)(v)(t);
                };
            },
            Comonad0: function () {
                return comonadTracedT1;
            }
        };
    };
};
var comonadTracedStoreT = function (dictComonadTraced) {
    var comonadStoreT = Control_Comonad_Store_Trans.comonadStoreT(dictComonadTraced.Comonad0());
    return {
        track: lowerTrack1(dictComonadTraced),
        Comonad0: function () {
            return comonadStoreT;
        }
    };
};
var comonadTracedIdentityT = function (dictComonadTraced) {
    var comonadIdentityT = Control_Monad_Identity_Trans.comonadIdentityT(dictComonadTraced.Comonad0());
    return {
        track: lowerTrack2(dictComonadTraced),
        Comonad0: function () {
            return comonadIdentityT;
        }
    };
};
var comonadTracedEnvT = function (dictComonadTraced) {
    var comonadEnvT = Control_Comonad_Env_Trans.comonadEnvT(dictComonadTraced.Comonad0());
    return {
        track: lowerTrack3(dictComonadTraced),
        Comonad0: function () {
            return comonadEnvT;
        }
    };
};

// | Apply a function to the current position.
var censor = function (dictFunctor) {
    return function (f) {
        return function (v) {
            return Data_Functor.map(dictFunctor)(function (v1) {
                return function ($44) {
                    return v1(f($44));
                };
            })(v);
        };
    };
};
export {
    track,
    tracks,
    listen,
    listens,
    censor,
    comonadTracedTracedT,
    comonadTracedIdentityT,
    comonadTracedEnvT,
    comonadTracedStoreT
};
