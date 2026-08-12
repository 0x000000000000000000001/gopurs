import * as Data_Array from "../Data.Array/index.js";
import * as Data_Array_NonEmpty_Internal from "../Data.Array.NonEmpty.Internal/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Data_String_Common from "../Data.String.Common/index.js";
import * as PureScript_Backend_Optimizer_CoreFn from "../PureScript.Backend.Optimizer.CoreFn/index.js";
import * as PureScript_Backend_Optimizer_Syntax from "../PureScript.Backend.Optimizer.Syntax/index.js";
var showExprType = function (v) {
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.Int) {
        return "Int";
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn["Number"]) {
        return "Number";
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn["String"]) {
        return "String";
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.Char) {
        return "Char";
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn["Boolean"]) {
        return "Boolean";
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.Any) {
        return "Any";
    };
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.Func) {
        return "Func([" + (Data_String_Common.joinWith(",")(Data_Functor.map(Data_Functor.functorArray)(showExprType)(v.value0)) + ("]," + (showExprType(v.value1) + ")")));
    };
    return "OtherType";
};
var showTco = function (v) {
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Var) {
        return "Var";
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Local) {
        return "Local";
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Lit) {
        return "Lit";
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.App) {
        return "App(" + (showTco(v.value1.value0) + ("," + (Data_Show.show(Data_Show.showInt)(Data_Array.length(Data_Array.fromFoldable(Data_Array_NonEmpty_Internal.foldableNonEmptyArray)(v.value1.value1))) + ")")));
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Typed) {
        return "Typed(" + (showExprType(v.value1.value0) + ("," + (showTco(v.value1.value1) + ")")));
    };
    return "Other";
};
export {
    showExprType,
    showTco
};
