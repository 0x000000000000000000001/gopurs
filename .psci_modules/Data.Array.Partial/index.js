// | Partial helper functions for working with immutable arrays.
import * as Data_Array from "../Data.Array/index.js";

// | Get all but the first element of a non-empty array.
// |
// | Running time: `O(n)`, where `n` is the length of the array.
var tail = function () {
    return function (xs) {
        return Data_Array.slice(1)(Data_Array.length(xs))(xs);
    };
};

// | Get the last element of a non-empty array.
// |
// | Running time: `O(1)`.
var last = function () {
    return function (xs) {
        return xs[Data_Array.length(xs) - 1 | 0];
    };
};

// | Get all but the last element of a non-empty array.
// |
// | Running time: `O(n)`, where `n` is the length of the array.
var init = function () {
    return function (xs) {
        return Data_Array.slice(0)(Data_Array.length(xs) - 1 | 0)(xs);
    };
};

// | Get the first element of a non-empty array.
// |
// | Running time: `O(1)`.
var head = function () {
    return function (xs) {
        return xs[0];
    };
};
export {
    head,
    tail,
    last,
    init
};
