import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Monad_Except_Trans from "../Control.Monad.Except.Trans/index.js";
import * as Control_Monad_Free from "../Control.Monad.Free/index.js";
import * as Control_Monad_Maybe_Trans from "../Control.Monad.Maybe.Trans/index.js";
import * as Control_Monad_Reader_Trans from "../Control.Monad.Reader.Trans/index.js";
import * as Control_Monad_State_Trans from "../Control.Monad.State.Trans/index.js";
import * as Control_Monad_Writer_Trans from "../Control.Monad.Writer.Trans/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
var wrapFree = function (dict) {
    return dict.wrapFree;
};
var monadFreeWriterT = function (dictFunctor) {
    return function (dictMonadFree) {
        var Monad0 = dictMonadFree.Monad0();
        return function (dictMonoid) {
            var monadWriterT = Control_Monad_Writer_Trans.monadWriterT(dictMonoid)(Monad0);
            return {
                wrapFree: function (f) {
                    return wrapFree(dictMonadFree)(Data_Functor.map(dictFunctor)(Control_Monad_Writer_Trans.runWriterT)(f));
                },
                Monad0: function () {
                    return monadWriterT;
                }
            };
        };
    };
};
var monadFreeStateT = function (dictFunctor) {
    return function (dictMonadFree) {
        var monadStateT = Control_Monad_State_Trans.monadStateT(dictMonadFree.Monad0());
        return {
            wrapFree: function (f) {
                return function (s) {
                    return wrapFree(dictMonadFree)(Data_Functor.map(dictFunctor)(function (st) {
                        return Control_Monad_State_Trans.runStateT(st)(s);
                    })(f));
                };
            },
            Monad0: function () {
                return monadStateT;
            }
        };
    };
};
var monadFreeReaderT = function (dictFunctor) {
    return function (dictMonadFree) {
        var monadReaderT = Control_Monad_Reader_Trans.monadReaderT(dictMonadFree.Monad0());
        return {
            wrapFree: function (f) {
                return function (r) {
                    return wrapFree(dictMonadFree)(Data_Functor.map(dictFunctor)(function (rt) {
                        return Control_Monad_Reader_Trans.runReaderT(rt)(r);
                    })(f));
                };
            },
            Monad0: function () {
                return monadReaderT;
            }
        };
    };
};
var monadFreeMaybeT = function (dictFunctor) {
    return function (dictMonadFree) {
        var monadMaybeT = Control_Monad_Maybe_Trans.monadMaybeT(dictMonadFree.Monad0());
        return {
            wrapFree: function (f) {
                return wrapFree(dictMonadFree)(Data_Functor.map(dictFunctor)(Control_Monad_Maybe_Trans.runMaybeT)(f));
            },
            Monad0: function () {
                return monadMaybeT;
            }
        };
    };
};
var monadFreeFree = {
    wrapFree: /* #__PURE__ */ (function () {
        var $24 = Control_Bind.join(Control_Monad_Free.freeBind);
        return function ($25) {
            return $24(Control_Monad_Free.liftF($25));
        };
    })(),
    Monad0: function () {
        return Control_Monad_Free.freeMonad;
    }
};
var monadFreeExceptT = function (dictFunctor) {
    return function (dictMonadFree) {
        var monadExceptT = Control_Monad_Except_Trans.monadExceptT(dictMonadFree.Monad0());
        return {
            wrapFree: function (f) {
                return wrapFree(dictMonadFree)(Data_Functor.map(dictFunctor)(Control_Monad_Except_Trans.runExceptT)(f));
            },
            Monad0: function () {
                return monadExceptT;
            }
        };
    };
};
export {
    wrapFree,
    monadFreeFree,
    monadFreeReaderT,
    monadFreeStateT,
    monadFreeWriterT,
    monadFreeMaybeT,
    monadFreeExceptT
};
