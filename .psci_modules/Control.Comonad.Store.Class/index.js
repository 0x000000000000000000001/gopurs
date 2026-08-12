// | This module defines the `ComonadStore` type class and its instances.
import * as Control_Comonad from "../Control.Comonad/index.js";
import * as Control_Comonad_Env_Trans from "../Control.Comonad.Env.Trans/index.js";
import * as Control_Comonad_Store_Trans from "../Control.Comonad.Store.Trans/index.js";
import * as Control_Comonad_Traced_Trans from "../Control.Comonad.Traced.Trans/index.js";
import * as Control_Comonad_Trans_Class from "../Control.Comonad.Trans.Class/index.js";
import * as Control_Extend from "../Control.Extend/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
var lower = /* #__PURE__ */ Control_Comonad_Trans_Class.lower(Control_Comonad_Env_Trans.comonadTransEnvT);
var pos = function (dict) {
    return dict.pos;
};
var peek = function (dict) {
    return dict.peek;
};

// | Extract a value from a position which depends on the current position.
var peeks = function (dictComonadStore) {
    return function (f) {
        return function (x) {
            return peek(dictComonadStore)(f(pos(dictComonadStore)(x)))(x);
        };
    };
};

// | Reposition the focus at the specified position, which depends on the current position.
var seeks = function (dictComonadStore) {
    var duplicate = Control_Extend.duplicate((dictComonadStore.Comonad0()).Extend0());
    return function (f) {
        var $40 = peeks(dictComonadStore)(f);
        return function ($41) {
            return $40(duplicate($41));
        };
    };
};

// | Reposition the focus at the specified position.
var seek = function (dictComonadStore) {
    var duplicate = Control_Extend.duplicate((dictComonadStore.Comonad0()).Extend0());
    return function (s) {
        var $42 = peek(dictComonadStore)(s);
        return function ($43) {
            return $42(duplicate($43));
        };
    };
};

// | Extract a collection of values from positions which depend on the current position.
var experiment = function (dictComonadStore) {
    var peek1 = peek(dictComonadStore);
    return function (dictFunctor) {
        return function (f) {
            return function (x) {
                return Data_Functor.map(dictFunctor)(Data_Function.flip(peek1)(x))(f(pos(dictComonadStore)(x)));
            };
        };
    };
};
var comonadStoreTracedT = function (dictComonadStore) {
    var pos1 = pos(dictComonadStore);
    var Comonad0 = dictComonadStore.Comonad0();
    var comonadTracedT = Control_Comonad_Traced_Trans.comonadTracedT(Comonad0);
    return function (dictMonoid) {
        var lower1 = Control_Comonad_Trans_Class.lower(Control_Comonad_Traced_Trans.comonadTransTracedT(dictMonoid))(Comonad0);
        var comonadTracedT1 = comonadTracedT(dictMonoid);
        return {
            pos: function ($44) {
                return pos1(lower1($44));
            },
            peek: function (s) {
                var $45 = peek(dictComonadStore)(s);
                return function ($46) {
                    return $45(lower1($46));
                };
            },
            Comonad0: function () {
                return comonadTracedT1;
            }
        };
    };
};
var comonadStoreStoreT = function (dictComonad) {
    var comonadStoreT = Control_Comonad_Store_Trans.comonadStoreT(dictComonad);
    return {
        pos: function (v) {
            return v.value1;
        },
        peek: function (s) {
            return function (v) {
                return Control_Comonad.extract(dictComonad)(v.value0)(s);
            };
        },
        Comonad0: function () {
            return comonadStoreT;
        }
    };
};
var comonadStoreEnvT = function (dictComonadStore) {
    var Comonad0 = dictComonadStore.Comonad0();
    var lower1 = lower(Comonad0);
    var comonadEnvT = Control_Comonad_Env_Trans.comonadEnvT(Comonad0);
    return {
        pos: (function () {
            var $47 = pos(dictComonadStore);
            return function ($48) {
                return $47(lower1($48));
            };
        })(),
        peek: function (s) {
            var $49 = peek(dictComonadStore)(s);
            return function ($50) {
                return $49(lower1($50));
            };
        },
        Comonad0: function () {
            return comonadEnvT;
        }
    };
};
export {
    peek,
    pos,
    experiment,
    peeks,
    seek,
    seeks,
    comonadStoreStoreT,
    comonadStoreEnvT,
    comonadStoreTracedT
};
