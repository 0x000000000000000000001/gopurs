import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_List from "../Data.List/index.js";
import * as Data_List_Types from "../Data.List.Types/index.js";
var Buffer = function (x) {
    return x;
};
var revert = function (v) {
    return {
        buffer: v.buffer,
        queue: Data_List.drop(1)(v.queue)
    };
};
var $$new = function (buffer) {
    return {
        buffer: buffer,
        queue: Data_List_Types.Nil.value
    };
};
var modify = function (f) {
    return function (v) {
        if (v.queue instanceof Data_List_Types.Cons) {
            return {
                buffer: v.buffer,
                queue: new Data_List_Types.Cons(new Data_List_Types.Cons(f, v.queue.value0), v.queue.value1)
            };
        };
        return {
            queue: v.queue,
            buffer: f(v.buffer)
        };
    };
};
var isBranching = function (v) {
    return !Data_List["null"](v.queue);
};
var commit = function (v) {
    return {
        buffer: Data_Foldable.foldr(Data_List_Types.foldableList)(Data_Function.flip(Data_Foldable.foldr(Data_List_Types.foldableList)(Data_Function.apply)))(v.buffer)(v.queue),
        queue: Data_List_Types.Nil.value
    };
};
var get = function ($21) {
    return (function (v) {
        return v.buffer;
    })(commit($21));
};
var branch = function (v) {
    return {
        buffer: v.buffer,
        queue: new Data_List_Types.Cons(Data_List_Types.Nil.value, v.queue)
    };
};
export {
    Buffer,
    branch,
    commit,
    revert,
    modify,
    get,
    $$new as new,
    isBranching
};
