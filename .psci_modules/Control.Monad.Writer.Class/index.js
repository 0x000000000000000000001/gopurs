// | This module defines the `MonadWriter` type class and its instances.
import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
var tell = function (dict) {
    return dict.tell;
};
var pass = function (dict) {
    return dict.pass;
};
var listen = function (dict) {
    return dict.listen;
};

// | Projects a value from modifications made to the accumulator during an
// | action.
var listens = function (dictMonadWriter) {
    var Monad1 = (dictMonadWriter.MonadTell1()).Monad1();
    var Bind1 = Monad1.Bind1();
    var pure = Control_Applicative.pure(Monad1.Applicative0());
    return function (f) {
        return function (m) {
            return Control_Bind.bind(Bind1)(listen(dictMonadWriter)(m))(function (v) {
                return pure(new Data_Tuple.Tuple(v.value0, f(v.value1)));
            });
        };
    };
};

// | Modify the final accumulator value by applying a function.
var censor = function (dictMonadWriter) {
    var Monad1 = (dictMonadWriter.MonadTell1()).Monad1();
    var Bind1 = Monad1.Bind1();
    var pure = Control_Applicative.pure(Monad1.Applicative0());
    return function (f) {
        return function (m) {
            return pass(dictMonadWriter)(Control_Bind.bind(Bind1)(m)(function (a) {
                return pure(new Data_Tuple.Tuple(a, f));
            }));
        };
    };
};
export {
    listen,
    pass,
    tell,
    listens,
    censor
};
