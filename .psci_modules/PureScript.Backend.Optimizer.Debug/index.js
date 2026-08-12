import * as $foreign from "./foreign.js";
var traceWhen = function () {
    return function (bool) {
        return function (a) {
            return function (b) {
                if (bool) {
                    return $foreign.traceImpl(a)(function (v) {
                        return b;
                    });
                };
                return b;
            };
        };
    };
};
var time = function () {
    return $foreign.time_;
};
var spyWhen = function () {
    return function (bool) {
        return function (a) {
            return traceWhen()(bool)(a)(a);
        };
    };
};
export {
    time_,
    traceImpl
} from "./foreign.js";
export {
    traceWhen,
    spyWhen,
    time
};
