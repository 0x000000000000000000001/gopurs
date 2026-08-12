import * as Data_Functor from "../Data.Functor/index.js";
var invariantMultiplicative = {
    imap: function (f) {
        return function (v) {
            return function (v1) {
                return f(v1);
            };
        };
    }
};
var invariantEndo = {
    imap: function (ab) {
        return function (ba) {
            return function (v) {
                return function ($39) {
                    return ab(v(ba($39)));
                };
            };
        };
    }
};
var invariantDual = {
    imap: function (f) {
        return function (v) {
            return function (v1) {
                return f(v1);
            };
        };
    }
};
var invariantDisj = {
    imap: function (f) {
        return function (v) {
            return function (v1) {
                return f(v1);
            };
        };
    }
};
var invariantConj = {
    imap: function (f) {
        return function (v) {
            return function (v1) {
                return f(v1);
            };
        };
    }
};
var invariantAdditive = {
    imap: function (f) {
        return function (v) {
            return function (v1) {
                return f(v1);
            };
        };
    }
};

// | As all `Functor`s are also trivially `Invariant`, this function can be
// | used as the `imap` implementation for any types that has an existing
// | `Functor` instance.
var imapF = function (dictFunctor) {
    return function (f) {
        return function (v) {
            return Data_Functor.map(dictFunctor)(f);
        };
    };
};
var invariantArray = {
    imap: /* #__PURE__ */ imapF(Data_Functor.functorArray)
};
var invariantFn = {
    imap: /* #__PURE__ */ imapF(Data_Functor.functorFn)
};
var imap = function (dict) {
    return dict.imap;
};
var invariantAlternate = function (dictInvariant) {
    return {
        imap: function (f) {
            return function (g) {
                return function (v) {
                    return imap(dictInvariant)(f)(g)(v);
                };
            };
        }
    };
};
export {
    imap,
    imapF,
    invariantFn,
    invariantArray,
    invariantAdditive,
    invariantConj,
    invariantDisj,
    invariantDual,
    invariantEndo,
    invariantMultiplicative,
    invariantAlternate
};
