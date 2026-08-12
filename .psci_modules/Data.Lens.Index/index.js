import * as Data_Array from "../Data.Array/index.js";
import * as Data_Array_NonEmpty from "../Data.Array.NonEmpty/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Lens_AffineTraversal from "../Data.Lens.AffineTraversal/index.js";
import * as Data_Lens_Iso_Newtype from "../Data.Lens.Iso.Newtype/index.js";
import * as Data_Lens_Lens from "../Data.Lens.Lens/index.js";
import * as Data_Lens_Prism_Maybe from "../Data.Lens.Prism.Maybe/index.js";
import * as Data_List from "../Data.List/index.js";
import * as Data_Map_Internal from "../Data.Map.Internal/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Set from "../Data.Set/index.js";
import * as Data_Unit from "../Data.Unit/index.js";
import * as Foreign_Object from "../Foreign.Object/index.js";
var _Newtype = /* #__PURE__ */ Data_Lens_Iso_Newtype["_Newtype"]()();
var ix = function (dict) {
    return dict.ix;
};
var indexSet = function (dictOrd) {
    return {
        ix: function (x) {
            return function (dictStrong) {
                return function (dictChoice) {
                    var set = function (xs) {
                        return function (v) {
                            return xs;
                        };
                    };
                    var pre = function (xs) {
                        var $32 = Data_Set.member(dictOrd)(x)(xs);
                        if ($32) {
                            return new Data_Either.Right(Data_Unit.unit);
                        };
                        return new Data_Either.Left(xs);
                    };
                    return Data_Lens_AffineTraversal.affineTraversal(set)(pre)(dictStrong)(dictChoice);
                };
            };
        }
    };
};
var indexNonEmptyArray = {
    ix: function (n) {
        return function (dictStrong) {
            return function (dictChoice) {
                var set = function (s) {
                    return function (b) {
                        return Data_Maybe.fromMaybe(s)(Data_Array_NonEmpty.updateAt(n)(b)(s));
                    };
                };
                var pre = function (s) {
                    return Data_Maybe.maybe(new Data_Either.Left(s))(Data_Either.Right.create)(Data_Array_NonEmpty.index(s)(n));
                };
                return Data_Lens_AffineTraversal.affineTraversal(set)(pre)(dictStrong)(dictChoice);
            };
        };
    }
};
var indexMaybe = {
    ix: function (v) {
        return function (dictStrong) {
            return function (dictChoice) {
                return Data_Lens_Prism_Maybe["_Just"](dictChoice);
            };
        };
    }
};
var indexMap = function (dictOrd) {
    return {
        ix: function (k) {
            return function (dictStrong) {
                return function (dictChoice) {
                    var set = function (s) {
                        return function (b) {
                            return Data_Map_Internal.update(dictOrd)(function (v) {
                                return new Data_Maybe.Just(b);
                            })(k)(s);
                        };
                    };
                    var pre = function (s) {
                        return Data_Maybe.maybe(new Data_Either.Left(s))(Data_Either.Right.create)(Data_Map_Internal.lookup(dictOrd)(k)(s));
                    };
                    return Data_Lens_AffineTraversal.affineTraversal(set)(pre)(dictStrong)(dictChoice);
                };
            };
        }
    };
};
var indexList = {
    ix: function (n) {
        return function (dictStrong) {
            return function (dictChoice) {
                var set = function (s) {
                    return function (b) {
                        return Data_Maybe.fromMaybe(s)(Data_List.updateAt(n)(b)(s));
                    };
                };
                var pre = function (s) {
                    return Data_Maybe.maybe(new Data_Either.Left(s))(Data_Either.Right.create)(Data_List.index(s)(n));
                };
                return Data_Lens_AffineTraversal.affineTraversal(set)(pre)(dictStrong)(dictChoice);
            };
        };
    }
};
var indexIdentity = {
    ix: function (v) {
        return function (dictStrong) {
            return function (dictChoice) {
                return _Newtype(dictChoice.Profunctor0());
            };
        };
    }
};
var indexForeignObject = {
    ix: function (k) {
        return function (dictStrong) {
            return function (dictChoice) {
                var set = function (s) {
                    return function (b) {
                        return Foreign_Object.update(function (v) {
                            return new Data_Maybe.Just(b);
                        })(k)(s);
                    };
                };
                var pre = function (s) {
                    return Data_Maybe.maybe(new Data_Either.Left(s))(Data_Either.Right.create)(Foreign_Object.lookup(k)(s));
                };
                return Data_Lens_AffineTraversal.affineTraversal(set)(pre)(dictStrong)(dictChoice);
            };
        };
    }
};
var indexFn = function (dictEq) {
    return {
        ix: function (i) {
            return function (dictStrong) {
                return function (dictChoice) {
                    return Data_Lens_Lens.lens(function (f) {
                        return f(i);
                    })(function (f) {
                        return function (a) {
                            return function (j) {
                                var $33 = Data_Eq.eq(dictEq)(i)(j);
                                if ($33) {
                                    return a;
                                };
                                return f(j);
                            };
                        };
                    })(dictStrong);
                };
            };
        }
    };
};
var indexArray = {
    ix: function (n) {
        return function (dictStrong) {
            return function (dictChoice) {
                var set = function (s) {
                    return function (b) {
                        return Data_Maybe.fromMaybe(s)(Data_Array.updateAt(n)(b)(s));
                    };
                };
                var pre = function (s) {
                    return Data_Maybe.maybe(new Data_Either.Left(s))(Data_Either.Right.create)(Data_Array.index(s)(n));
                };
                return Data_Lens_AffineTraversal.affineTraversal(set)(pre)(dictStrong)(dictChoice);
            };
        };
    }
};
export {
    ix,
    indexFn,
    indexMaybe,
    indexIdentity,
    indexArray,
    indexNonEmptyArray,
    indexList,
    indexSet,
    indexMap,
    indexForeignObject
};
