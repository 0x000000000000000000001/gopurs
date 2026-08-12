import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Effect_Class from "../Effect.Class/index.js";
import * as Effect_Console from "../Effect.Console/index.js";
var warnShow = function (dictMonadEffect) {
    var liftEffect = Effect_Class.liftEffect(dictMonadEffect);
    return function (dictShow) {
        var $47 = Effect_Console.warnShow(dictShow);
        return function ($48) {
            return liftEffect($47($48));
        };
    };
};
var warn = function (dictMonadEffect) {
    var $49 = Effect_Class.liftEffect(dictMonadEffect);
    return function ($50) {
        return $49(Effect_Console.warn($50));
    };
};
var timeLog = function (dictMonadEffect) {
    var $51 = Effect_Class.liftEffect(dictMonadEffect);
    return function ($52) {
        return $51(Effect_Console.timeLog($52));
    };
};
var timeEnd = function (dictMonadEffect) {
    var $53 = Effect_Class.liftEffect(dictMonadEffect);
    return function ($54) {
        return $53(Effect_Console.timeEnd($54));
    };
};
var time = function (dictMonadEffect) {
    var $55 = Effect_Class.liftEffect(dictMonadEffect);
    return function ($56) {
        return $55(Effect_Console.time($56));
    };
};
var logShow = function (dictMonadEffect) {
    var liftEffect = Effect_Class.liftEffect(dictMonadEffect);
    return function (dictShow) {
        var $57 = Effect_Console.logShow(dictShow);
        return function ($58) {
            return liftEffect($57($58));
        };
    };
};
var log = function (dictMonadEffect) {
    var $59 = Effect_Class.liftEffect(dictMonadEffect);
    return function ($60) {
        return $59(Effect_Console.log($60));
    };
};
var infoShow = function (dictMonadEffect) {
    var liftEffect = Effect_Class.liftEffect(dictMonadEffect);
    return function (dictShow) {
        var $61 = Effect_Console.infoShow(dictShow);
        return function ($62) {
            return liftEffect($61($62));
        };
    };
};
var info = function (dictMonadEffect) {
    var $63 = Effect_Class.liftEffect(dictMonadEffect);
    return function ($64) {
        return $63(Effect_Console.info($64));
    };
};
var groupEnd = function (dictMonadEffect) {
    return Effect_Class.liftEffect(dictMonadEffect)(Effect_Console.groupEnd);
};
var groupCollapsed = function (dictMonadEffect) {
    var $65 = Effect_Class.liftEffect(dictMonadEffect);
    return function ($66) {
        return $65(Effect_Console.groupCollapsed($66));
    };
};
var group = function (dictMonadEffect) {
    var $67 = Effect_Class.liftEffect(dictMonadEffect);
    return function ($68) {
        return $67(Effect_Console.group($68));
    };
};
var grouped = function (dictMonadEffect) {
    var Monad0 = dictMonadEffect.Monad0();
    var Bind1 = Monad0.Bind1();
    var groupEnd1 = groupEnd(dictMonadEffect);
    var Applicative0 = Monad0.Applicative0();
    return function (name) {
        return function (inner) {
            return Control_Bind.discard(Control_Bind.discardUnit)(Bind1)(group(dictMonadEffect)(name))(function () {
                return Control_Bind.bind(Bind1)(inner)(function (result) {
                    return Control_Bind.discard(Control_Bind.discardUnit)(Bind1)(groupEnd1)(function () {
                        return Control_Applicative.pure(Applicative0)(result);
                    });
                });
            });
        };
    };
};
var errorShow = function (dictMonadEffect) {
    var liftEffect = Effect_Class.liftEffect(dictMonadEffect);
    return function (dictShow) {
        var $69 = Effect_Console.errorShow(dictShow);
        return function ($70) {
            return liftEffect($69($70));
        };
    };
};
var error = function (dictMonadEffect) {
    var $71 = Effect_Class.liftEffect(dictMonadEffect);
    return function ($72) {
        return $71(Effect_Console.error($72));
    };
};
var debugShow = function (dictMonadEffect) {
    var liftEffect = Effect_Class.liftEffect(dictMonadEffect);
    return function (dictShow) {
        var $73 = Effect_Console.debugShow(dictShow);
        return function ($74) {
            return liftEffect($73($74));
        };
    };
};
var debug = function (dictMonadEffect) {
    var $75 = Effect_Class.liftEffect(dictMonadEffect);
    return function ($76) {
        return $75(Effect_Console.debug($76));
    };
};
var clear = function (dictMonadEffect) {
    return Effect_Class.liftEffect(dictMonadEffect)(Effect_Console.clear);
};
export {
    log,
    logShow,
    warn,
    warnShow,
    error,
    errorShow,
    info,
    infoShow,
    debug,
    debugShow,
    time,
    timeLog,
    timeEnd,
    clear,
    group,
    groupCollapsed,
    groupEnd,
    grouped
};
