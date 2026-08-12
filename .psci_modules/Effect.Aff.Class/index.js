import * as Control_Category from "../Control.Category/index.js";
import * as Control_Monad_Cont_Trans from "../Control.Monad.Cont.Trans/index.js";
import * as Control_Monad_Except_Trans from "../Control.Monad.Except.Trans/index.js";
import * as Control_Monad_List_Trans from "../Control.Monad.List.Trans/index.js";
import * as Control_Monad_Maybe_Trans from "../Control.Monad.Maybe.Trans/index.js";
import * as Control_Monad_RWS_Trans from "../Control.Monad.RWS.Trans/index.js";
import * as Control_Monad_Reader_Trans from "../Control.Monad.Reader.Trans/index.js";
import * as Control_Monad_State_Trans from "../Control.Monad.State.Trans/index.js";
import * as Control_Monad_Trans_Class from "../Control.Monad.Trans.Class/index.js";
import * as Control_Monad_Writer_Trans from "../Control.Monad.Writer.Trans/index.js";
import * as Effect_Aff from "../Effect.Aff/index.js";
var lift = /* #__PURE__ */ Control_Monad_Trans_Class.lift(Control_Monad_Cont_Trans.monadTransContT);
var lift1 = /* #__PURE__ */ Control_Monad_Trans_Class.lift(Control_Monad_Except_Trans.monadTransExceptT);
var lift2 = /* #__PURE__ */ Control_Monad_Trans_Class.lift(Control_Monad_List_Trans.monadTransListT);
var lift3 = /* #__PURE__ */ Control_Monad_Trans_Class.lift(Control_Monad_Maybe_Trans.monadTransMaybeT);
var lift4 = /* #__PURE__ */ Control_Monad_Trans_Class.lift(Control_Monad_Reader_Trans.monadTransReaderT);
var lift5 = /* #__PURE__ */ Control_Monad_Trans_Class.lift(Control_Monad_State_Trans.monadTransStateT);
var monadAffAff = {
    liftAff: /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn),
    MonadEffect0: function () {
        return Effect_Aff.monadEffectAff;
    }
};
var liftAff = function (dict) {
    return dict.liftAff;
};
var monadAffContT = function (dictMonadAff) {
    var MonadEffect0 = dictMonadAff.MonadEffect0();
    var monadEffectContT = Control_Monad_Cont_Trans.monadEffectContT(MonadEffect0);
    return {
        liftAff: (function () {
            var $64 = lift(MonadEffect0.Monad0());
            var $65 = liftAff(dictMonadAff);
            return function ($66) {
                return $64($65($66));
            };
        })(),
        MonadEffect0: function () {
            return monadEffectContT;
        }
    };
};
var monadAffExceptT = function (dictMonadAff) {
    var MonadEffect0 = dictMonadAff.MonadEffect0();
    var monadEffectExceptT = Control_Monad_Except_Trans.monadEffectExceptT(MonadEffect0);
    return {
        liftAff: (function () {
            var $67 = lift1(MonadEffect0.Monad0());
            var $68 = liftAff(dictMonadAff);
            return function ($69) {
                return $67($68($69));
            };
        })(),
        MonadEffect0: function () {
            return monadEffectExceptT;
        }
    };
};
var monadAffListT = function (dictMonadAff) {
    var MonadEffect0 = dictMonadAff.MonadEffect0();
    var monadEffectListT = Control_Monad_List_Trans.monadEffectListT(MonadEffect0);
    return {
        liftAff: (function () {
            var $70 = lift2(MonadEffect0.Monad0());
            var $71 = liftAff(dictMonadAff);
            return function ($72) {
                return $70($71($72));
            };
        })(),
        MonadEffect0: function () {
            return monadEffectListT;
        }
    };
};
var monadAffMaybe = function (dictMonadAff) {
    var MonadEffect0 = dictMonadAff.MonadEffect0();
    var monadEffectMaybe = Control_Monad_Maybe_Trans.monadEffectMaybe(MonadEffect0);
    return {
        liftAff: (function () {
            var $73 = lift3(MonadEffect0.Monad0());
            var $74 = liftAff(dictMonadAff);
            return function ($75) {
                return $73($74($75));
            };
        })(),
        MonadEffect0: function () {
            return monadEffectMaybe;
        }
    };
};
var monadAffRWS = function (dictMonadAff) {
    var MonadEffect0 = dictMonadAff.MonadEffect0();
    var Monad0 = MonadEffect0.Monad0();
    var liftAff1 = liftAff(dictMonadAff);
    return function (dictMonoid) {
        var monadEffectRWS = Control_Monad_RWS_Trans.monadEffectRWS(dictMonoid)(MonadEffect0);
        return {
            liftAff: (function () {
                var $76 = Control_Monad_Trans_Class.lift(Control_Monad_RWS_Trans.monadTransRWST(dictMonoid))(Monad0);
                return function ($77) {
                    return $76(liftAff1($77));
                };
            })(),
            MonadEffect0: function () {
                return monadEffectRWS;
            }
        };
    };
};
var monadAffReader = function (dictMonadAff) {
    var MonadEffect0 = dictMonadAff.MonadEffect0();
    var monadEffectReader = Control_Monad_Reader_Trans.monadEffectReader(MonadEffect0);
    return {
        liftAff: (function () {
            var $78 = lift4(MonadEffect0.Monad0());
            var $79 = liftAff(dictMonadAff);
            return function ($80) {
                return $78($79($80));
            };
        })(),
        MonadEffect0: function () {
            return monadEffectReader;
        }
    };
};
var monadAffState = function (dictMonadAff) {
    var MonadEffect0 = dictMonadAff.MonadEffect0();
    var monadEffectState = Control_Monad_State_Trans.monadEffectState(MonadEffect0);
    return {
        liftAff: (function () {
            var $81 = lift5(MonadEffect0.Monad0());
            var $82 = liftAff(dictMonadAff);
            return function ($83) {
                return $81($82($83));
            };
        })(),
        MonadEffect0: function () {
            return monadEffectState;
        }
    };
};
var monadAffWriter = function (dictMonadAff) {
    var MonadEffect0 = dictMonadAff.MonadEffect0();
    var Monad0 = MonadEffect0.Monad0();
    var liftAff1 = liftAff(dictMonadAff);
    return function (dictMonoid) {
        var monadEffectWriter = Control_Monad_Writer_Trans.monadEffectWriter(dictMonoid)(MonadEffect0);
        return {
            liftAff: (function () {
                var $84 = Control_Monad_Trans_Class.lift(Control_Monad_Writer_Trans.monadTransWriterT(dictMonoid))(Monad0);
                return function ($85) {
                    return $84(liftAff1($85));
                };
            })(),
            MonadEffect0: function () {
                return monadEffectWriter;
            }
        };
    };
};
export {
    liftAff,
    monadAffAff,
    monadAffContT,
    monadAffExceptT,
    monadAffListT,
    monadAffMaybe,
    monadAffReader,
    monadAffRWS,
    monadAffState,
    monadAffWriter
};
