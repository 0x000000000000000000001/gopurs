// | This module defines functions for working with setters.
import * as Control_Monad_State_Class from "../Control.Monad.State.Class/index.js";
import * as Data_EuclideanRing from "../Data.EuclideanRing/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_HeytingAlgebra from "../Data.HeytingAlgebra/index.js";
import * as Data_Lens_Types from "../Data.Lens.Types/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Ring from "../Data.Ring/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Semiring from "../Data.Semiring/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";

// | Apply a function to the foci of a `Setter`.
var over = function (l) {
    return l;
};

// | Set the foci of a `Setter` to a constant value.
var set = function (l) {
    return function (b) {
        return over(l)(Data_Function["const"](b));
    };
};
var setJust = function (p) {
    var $46 = set(p);
    return function ($47) {
        return $46(Data_Maybe.Just.create($47));
    };
};
var subOver = function (dictRing) {
    var sub = Data_Ring.sub(dictRing);
    return function (p) {
        var $48 = over(p);
        var $49 = Data_Function.flip(sub);
        return function ($50) {
            return $48($49($50));
        };
    };
};
var mulOver = function (dictSemiring) {
    var mul = Data_Semiring.mul(dictSemiring);
    return function (p) {
        var $51 = over(p);
        var $52 = Data_Function.flip(mul);
        return function ($53) {
            return $51($52($53));
        };
    };
};

// | Modify the foci of a `Setter` in a monadic state.
var modifying = function (dictMonadState) {
    var Functor0 = (((dictMonadState.Monad0()).Bind1()).Apply0()).Functor0();
    return function (p) {
        return function (f) {
            return Data_Functor["void"](Functor0)(Control_Monad_State_Class.modify(dictMonadState)(over(p)(f)));
        };
    };
};
var mulModifying = function (dictMonadState) {
    return function (dictSemiring) {
        var mul = Data_Semiring.mul(dictSemiring);
        return function (p) {
            var $54 = modifying(dictMonadState)(p);
            var $55 = Data_Function.flip(mul);
            return function ($56) {
                return $54($55($56));
            };
        };
    };
};
var subModifying = function (dictMonadState) {
    return function (dictRing) {
        var sub = Data_Ring.sub(dictRing);
        return function (p) {
            var $57 = modifying(dictMonadState)(p);
            var $58 = Data_Function.flip(sub);
            return function ($59) {
                return $57($58($59));
            };
        };
    };
};

// | Apply a function to the foci of a `Setter` that may vary with the index.
var iover = function (l) {
    return function (f) {
        return l(Data_Tuple.uncurry(f));
    };
};
var divOver = function (dictEuclideanRing) {
    var div = Data_EuclideanRing.div(dictEuclideanRing);
    return function (p) {
        var $60 = over(p);
        var $61 = Data_Function.flip(div);
        return function ($62) {
            return $60($61($62));
        };
    };
};
var divModifying = function (dictMonadState) {
    return function (dictEuclideanRing) {
        var div = Data_EuclideanRing.div(dictEuclideanRing);
        return function (p) {
            var $63 = modifying(dictMonadState)(p);
            var $64 = Data_Function.flip(div);
            return function ($65) {
                return $63($64($65));
            };
        };
    };
};
var disjOver = function (dictHeytingAlgebra) {
    var disj = Data_HeytingAlgebra.disj(dictHeytingAlgebra);
    return function (p) {
        var $66 = over(p);
        var $67 = Data_Function.flip(disj);
        return function ($68) {
            return $66($67($68));
        };
    };
};
var disjModifying = function (dictMonadState) {
    return function (dictHeytingAlgebra) {
        var disj = Data_HeytingAlgebra.disj(dictHeytingAlgebra);
        return function (p) {
            var $69 = modifying(dictMonadState)(p);
            var $70 = Data_Function.flip(disj);
            return function ($71) {
                return $69($70($71));
            };
        };
    };
};
var conjOver = function (dictHeytingAlgebra) {
    var conj = Data_HeytingAlgebra.conj(dictHeytingAlgebra);
    return function (p) {
        var $72 = over(p);
        var $73 = Data_Function.flip(conj);
        return function ($74) {
            return $72($73($74));
        };
    };
};
var conjModifying = function (dictMonadState) {
    return function (dictHeytingAlgebra) {
        var conj = Data_HeytingAlgebra.conj(dictHeytingAlgebra);
        return function (p) {
            var $75 = modifying(dictMonadState)(p);
            var $76 = Data_Function.flip(conj);
            return function ($77) {
                return $75($76($77));
            };
        };
    };
};

// Stateful
// | Set the foci of a `Setter` in a monadic state to a constant value.
var assign = function (dictMonadState) {
    var Functor0 = (((dictMonadState.Monad0()).Bind1()).Apply0()).Functor0();
    return function (p) {
        return function (b) {
            return Data_Functor["void"](Functor0)(Control_Monad_State_Class.modify(dictMonadState)(set(p)(b)));
        };
    };
};
var assignJust = function (dictMonadState) {
    return function (p) {
        var $78 = assign(dictMonadState)(p);
        return function ($79) {
            return $78(Data_Maybe.Just.create($79));
        };
    };
};
var appendOver = function (dictSemigroup) {
    var append = Data_Semigroup.append(dictSemigroup);
    return function (p) {
        var $80 = over(p);
        var $81 = Data_Function.flip(append);
        return function ($82) {
            return $80($81($82));
        };
    };
};
var appendModifying = function (dictMonadState) {
    return function (dictSemigroup) {
        var append = Data_Semigroup.append(dictSemigroup);
        return function (p) {
            var $83 = modifying(dictMonadState)(p);
            var $84 = Data_Function.flip(append);
            return function ($85) {
                return $83($84($85));
            };
        };
    };
};
var addOver = function (dictSemiring) {
    var add = Data_Semiring.add(dictSemiring);
    return function (p) {
        var $86 = over(p);
        return function ($87) {
            return $86(add($87));
        };
    };
};
var addModifying = function (dictMonadState) {
    return function (dictSemiring) {
        var add = Data_Semiring.add(dictSemiring);
        return function (p) {
            var $88 = modifying(dictMonadState)(p);
            return function ($89) {
                return $88(add($89));
            };
        };
    };
};
export {
    over,
    iover,
    set,
    addOver,
    subOver,
    mulOver,
    divOver,
    disjOver,
    conjOver,
    appendOver,
    setJust,
    assign,
    modifying,
    addModifying,
    mulModifying,
    subModifying,
    divModifying,
    disjModifying,
    conjModifying,
    appendModifying,
    assignJust
};
export {
    Indexed
} from "../Data.Lens.Types/index.js";
