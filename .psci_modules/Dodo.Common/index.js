import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Dodo from "../Dodo/index.js";
import * as Dodo_Internal from "../Dodo.Internal/index.js";
var trailingComma = /* #__PURE__ */ Dodo.flexAlt(/* #__PURE__ */ Dodo.text(", "))(/* #__PURE__ */ Data_Semigroup.append(Dodo_Internal.semigroupDoc)(/* #__PURE__ */ Dodo.text(","))(Dodo["break"]));
var pursSquares = /* #__PURE__ */ (function () {
    var open = Dodo.text("[ ");
    var close = Dodo.flexAlt(Dodo.text(" ]"))(Data_Semigroup.append(Dodo_Internal.semigroupDoc)(Dodo["break"])(Dodo.text("]")));
    var $0 = Dodo.encloseEmptyAlt(open)(close)(Dodo.text("[]"));
    return function ($1) {
        return Dodo.flexGroup($0($1));
    };
})();
var pursParensExpr = /* #__PURE__ */ (function () {
    var open = Dodo.text("(");
    var close = Dodo.text(")");
    var $2 = Dodo.enclose(open)(close);
    return function ($3) {
        return Dodo.flexGroup($2(Dodo.indent($3)));
    };
})();
var pursParens = /* #__PURE__ */ (function () {
    var open = Dodo.flexAlt(Dodo.text("("))(Dodo.text("( "));
    var close = Dodo.flexAlt(Dodo.text(")"))(Data_Semigroup.append(Dodo_Internal.semigroupDoc)(Dodo["break"])(Dodo.text(")")));
    var $4 = Dodo.encloseEmptyAlt(open)(close)(Dodo.text("()"));
    return function ($5) {
        return Dodo.flexGroup($4($5));
    };
})();
var pursCurlies = /* #__PURE__ */ (function () {
    var open = Dodo.text("{ ");
    var close = Dodo.flexAlt(Dodo.text(" }"))(Data_Semigroup.append(Dodo_Internal.semigroupDoc)(Dodo["break"])(Dodo.text("}")));
    var $6 = Dodo.encloseEmptyAlt(open)(close)(Dodo.text("{}"));
    return function ($7) {
        return Dodo.flexGroup($6($7));
    };
})();
var leadingComma = /* #__PURE__ */ Dodo.flexAlt(/* #__PURE__ */ Dodo.text(", "))(/* #__PURE__ */ Data_Semigroup.append(Dodo_Internal.semigroupDoc)(Dodo["break"])(/* #__PURE__ */ Dodo.text(", ")));
var jsSquares = /* #__PURE__ */ (function () {
    var open = Dodo.flexAlt(Dodo.text("["))(Data_Semigroup.append(Dodo_Internal.semigroupDoc)(Dodo.text("["))(Dodo["break"]));
    var close = Dodo.flexAlt(Dodo.text("]"))(Data_Semigroup.append(Dodo_Internal.semigroupDoc)(Dodo["break"])(Dodo.text("]")));
    var $8 = Dodo.encloseEmptyAlt(open)(close)(Dodo.text("[]"));
    return function ($9) {
        return Dodo.flexGroup($8(Dodo.indent($9)));
    };
})();
var jsParens = /* #__PURE__ */ (function () {
    var open = Dodo.flexAlt(Dodo.text("("))(Data_Semigroup.append(Dodo_Internal.semigroupDoc)(Dodo.text("("))(Dodo["break"]));
    var close = Dodo.flexAlt(Dodo.text(")"))(Data_Semigroup.append(Dodo_Internal.semigroupDoc)(Dodo["break"])(Dodo.text(")")));
    var $10 = Dodo.encloseEmptyAlt(open)(close)(Dodo.text("()"));
    return function ($11) {
        return Dodo.flexGroup($10(Dodo.indent($11)));
    };
})();
var jsCurlies = /* #__PURE__ */ (function () {
    var open = Dodo.flexAlt(Dodo.text("{"))(Data_Semigroup.append(Dodo_Internal.semigroupDoc)(Dodo.text("{"))(Dodo["break"]));
    var close = Dodo.flexAlt(Dodo.text("}"))(Data_Semigroup.append(Dodo_Internal.semigroupDoc)(Dodo["break"])(Dodo.text("}")));
    var $12 = Dodo.encloseEmptyAlt(open)(close)(Dodo.text("{}"));
    return function ($13) {
        return Dodo.flexGroup($12(Dodo.indent($13)));
    };
})();
export {
    leadingComma,
    trailingComma,
    pursCurlies,
    pursSquares,
    pursParens,
    pursParensExpr,
    jsCurlies,
    jsSquares,
    jsParens
};
