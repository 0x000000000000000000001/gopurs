import * as Control_Category from "../Control.Category/index.js";
import * as Data_Boolean from "../Data.Boolean/index.js";

// | The `on` function is used to change the domain of a binary operator.
// |
// | For example, we can create a function which compares two records based on the values of their `x` properties:
// |
// | ```purescript
// | compareX :: forall r. { x :: Number | r } -> { x :: Number | r } -> Ordering
// | compareX = compare `on` _.x
// | ```
var on = function (f) {
    return function (g) {
        return function (x) {
            return function (y) {
                return f(g(x))(g(y));
            };
        };
    };
};

// | Given a function that takes two arguments, applies the arguments
// | to the function in a swapped order.
// |
// | ```purescript
// | flip append "1" "2" == append "2" "1" == "21"
// |
// | const 1 "two" == 1
// |
// | flip const 1 "two" == const "two" 1 == "two"
// | ```
var flip = function (f) {
    return function (b) {
        return function (a) {
            return f(a)(b);
        };
    };
};

// | Returns its first argument and ignores its second.
// |
// | ```purescript
// | const 1 "hello" = 1
// | ```
// |
// | It can also be thought of as creating a function that ignores its argument:
// |
// | ```purescript
// | const 1 = \_ -> 1
// | ```
var $$const = function (a) {
    return function (v) {
        return a;
    };
};

// | `applyN f n` applies the function `f` to its argument `n` times.
// |
// | If n is less than or equal to 0, the function is not applied.
// |
// | ```purescript
// | applyN (_ + 1) 10 0 == 10
// | ```
var applyN = function (f) {
    var go = function ($copy_n) {
        return function ($copy_acc) {
            var $tco_var_n = $copy_n;
            var $tco_done = false;
            var $tco_result;
            function $tco_loop(n, acc) {
                if (n <= 0) {
                    $tco_done = true;
                    return acc;
                };
                if (Data_Boolean.otherwise) {
                    $tco_var_n = n - 1 | 0;
                    $copy_acc = f(acc);
                    return;
                };
                throw new Error("Failed pattern match at Data.Function (line 107, column 3 - line 109, column 37): " + [ n.constructor.name, acc.constructor.name ]);
            };
            while (!$tco_done) {
                $tco_result = $tco_loop($tco_var_n, $copy_acc);
            };
            return $tco_result;
        };
    };
    return go;
};

// | Applies an argument to a function. This is primarily used as the `(#)`
// | operator, which allows parentheses to be omitted in some cases, or as a
// | natural way to apply a value to a chain of composed functions.
var applyFlipped = function (x) {
    return function (f) {
        return f(x);
    };
};

// | Applies a function to an argument. This is primarily used as the operator
// | `($)` which allows parentheses to be omitted in some cases, or as a
// | natural way to apply a chain of composed functions to a value.
var apply = function (f) {
    return function (x) {
        return f(x);
    };
};
export {
    flip,
    $$const as const,
    apply,
    applyFlipped,
    applyN,
    on
};
export {
    compose,
    identity
} from "../Control.Category/index.js";
