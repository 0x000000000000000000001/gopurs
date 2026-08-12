import * as Control_Alt from "../Control.Alt/index.js";
import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Plus from "../Control.Plus/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Unit from "../Data.Unit/index.js";

// | Fail using `Plus` if a condition does not hold, or
// | succeed using `Applicative` if it does.
// |
// | For example:
// |
// | ```purescript
// | import Prelude
// | import Control.Alternative (guard)
// | import Data.Array ((..))
// |
// | factors :: Int -> Array Int
// | factors n = do
// |   a <- 1..n
// |   b <- 1..n
// |   guard $ a * b == n
// |   pure a
// | ```
var guard = function (dictAlternative) {
    var Applicative0 = dictAlternative.Applicative0();
    var empty = Control_Plus.empty(dictAlternative.Plus1());
    return function (v) {
        if (v) {
            return Control_Applicative.pure(Applicative0)(Data_Unit.unit);
        };
        if (!v) {
            return empty;
        };
        throw new Error("Failed pattern match at Control.Alternative (line 48, column 1 - line 48, column 54): " + [ v.constructor.name ]);
    };
};
var alternativeArray = {
    Applicative0: function () {
        return Control_Applicative.applicativeArray;
    },
    Plus1: function () {
        return Control_Plus.plusArray;
    }
};
export {
    guard,
    alternativeArray
};
export {
    alt
} from "../Control.Alt/index.js";
export {
    liftA1,
    pure,
    unless,
    when
} from "../Control.Applicative/index.js";
export {
    apply
} from "../Control.Apply/index.js";
export {
    empty
} from "../Control.Plus/index.js";
export {
    map,
    void
} from "../Data.Functor/index.js";
