// | This module defines common lenses and prisms.
import * as Data_Lens_Lens_Tuple from "../Data.Lens.Lens.Tuple/index.js";
import * as Data_Lens_Lens_Unit from "../Data.Lens.Lens.Unit/index.js";
import * as Data_Lens_Prism_Either from "../Data.Lens.Prism.Either/index.js";
import * as Data_Lens_Prism_Maybe from "../Data.Lens.Prism.Maybe/index.js";

// | This is useful for when you want to restrict the type of another optic.
// | For example, suppose you have the following declarations:
// | ```purescript
// | newtype X = X Int
// | derive instance newtypeX :: Newtype X _
// | ```
// |
// | Attempting to view with the `_Newtype` optic:
// | ```purescript
// | X 42 ^. _Newtype
// | ```
// | Will result in a type error:
// | ```
// |  The inferred type
// |    forall t3 t5. Newtype t3 t5 => Int
// |  has type variables which are not mentioned in the body of the type.
// |  Consider adding a type annotation.
// | ```
// |
// | However, if we apply the `simple` function:
// | ```purescript
// |  X 42 ^. simple _Newtype
// | ```
// | We get the expected result `42`.
var simple = function (x) {
    return x;
};
export {
    simple
};
export {
    _1,
    _2,
    first,
    second
} from "../Data.Lens.Lens.Tuple/index.js";
export {
    united
} from "../Data.Lens.Lens.Unit/index.js";
export {
    _Left,
    _Right,
    left,
    right
} from "../Data.Lens.Prism.Either/index.js";
export {
    _Just,
    _Nothing
} from "../Data.Lens.Prism.Maybe/index.js";
