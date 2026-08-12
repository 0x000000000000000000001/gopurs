import * as $foreign from "./foreign.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Nullable from "../Data.Nullable/index.js";
import * as Effect from "../Effect/index.js";
var showSymbol = function (s) {
    return $foreign.showSymbolImpl(s);
};
var showJsSymbol = {
    show: function (x) {
        return showSymbol(x);
    }
};
var keyFor = function (s) {
    return Data_Functor.map(Effect.functorEffect)(Data_Nullable.toMaybe)(function () {
        return $foreign.keyForImpl(s);
    });
};
var $$for = function (s) {
    return function () {
        return $foreign.forImpl(s);
    };
};
export {
    showSymbolImpl,
    forImpl,
    keyForImpl
} from "./foreign.js";
export {
    showSymbol,
    $$for as for,
    keyFor,
    showJsSymbol
};
