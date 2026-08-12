// | ### Algorithm Summary for Optimized Pattern Matching Conversion
// |
// | The algorithm used for converting `ExprCase` into `BackendExpr` is based on two papers:
// | 1. https://www.cs.tufts.edu/comp/150FP/archive/luc-maranget/jun08.pdf - "Compiling Pattern Matching to Good Decision Trees" (CPMtGDT paper)
// | 2. https://julesjacobs.com/notes/patternmatching/patternmatching.pdf - "How to compile pattern matching"
// |
// | The algorithm uses the composition of heuristics `p`, `b`, `a`, and pseudo-heuristic `N` described in
// | the CPMtGDT paper.
// |
// | The algorithm can be summarized as:
// |
// | 1. Entry point preprocessing steps:
// |     1. let-bind the expressions in the case head and refer to these as `caseHeadIdents`
// |     2. for each case row, convert each column's `Binder` into its corresponding `Pattern` type, removing Newtypes completely
// |     3. for each case row, zip the corresponding ident in `caseHeadIdents` with its corresponding `Pattern` value
// |     4. calculate the references introduced in a case row's binders, sort them by their name, store the result with the case row expression, and reference these as `leafFnArgs`
// |     5. let-bind the "leafs" in the case row's expression:
// |         1. if `Unconditional`, let-bind the expression as a function, using `leafFnArgs` to determine the number, order, and names of the function args, and using the original expression as the function's body.
// |         2. if `Guarded`, do the same let-bind-expression-as-function described in the `Unconditional` step for each `Guard` but do not yet convert the guards predicate. Since the references introduced by the case row's binders are not yet in scope, the predicate won't reference the correct values
// | 2. Start the recursive algorithm
// |     1. preprocess all record patterns as described previously, so that each case row's corresponding column has all fields referenced in that column and orders its fields in the same order throughout all case rows
// |     2. If the clause matrix has 0 rows, then we produce a pattern match failure
// |     3. Otherwise, there's at least 1 row. If the clause matrix's first row contains only wildcard patterns (e.g. `value is _`) or is otherwise empty
// |         1. calculate the `allReferences` value by combining the case row's "References" array with the references introduced by each remaining column (if any)
// |         2. Sort the `allReferences` array by the reference names, so that the order of the references matches the order originally calculated in `leafFnArgs`
// |         3. case on the guard
// |             1. if `Unconditional`, call the function it references with the ordered `allReferences` args
// |             2. if `Guarded`
// |                 1. add `allReferences` to the current scope and then convert the predicate.
// |                 2. call the function it references with the ordered `allReferences` args
// |     4. Otherwise, there's at least one column in the first row against which we still need to test (i.e. there is a `value is pattern` test where the `pattern` is not a wildcard/`_`).
// |         1. From among the remaining non-wildcard patterns we could test, use a heuristic to determine which column's `value is pattern` test from the first row will produce the smallest tree
// |             1. if the chosen column is a value that can always be expanded (e.g. a `Product` type or `Record` type), use that column
// |             2. Otherwise, use heuristic `pbaN`
// |         2. Build 2 new clause matrices, Problem A and Problem B, using the below rules. Problem A contains rows where a match occurred. Problem B contains rows where a match did not occur.
// |             1. If a row's corresponding column uses the same pattern as the chosen one (e.g. `chosen: a is 1; row's: a is 1`), then
// |                 1. in the case row, add the references row's corresponding pattern introduces the case rows' "References" array
// |                 2. in the case row's columns, replace the parent pattern with its subterm patterns (if any)
// |                 3. put the case row into Problem A because a match occurred
// |             2. If a row's corresponding column differs from the chosen one (e.g. `chosen: a is 1; row's: a is 2`), then put it in Problem B; a match did not occur.
// |             3. If a row's corresponding column is a wildcard (e.g. `chosen: a is 1; row's: a is _`)
// |                 1. follow the instructions above as if there was a normal match and put the resulting row in Problem A
// |                 2. put a copy of the row in Problem B
// |         3. If the chosen column is an expandable type, recurse on Problem A
// |         4. Otherwise, guard against the chosen pattern, recursing on Problem A if it succeeds and recursing on Problem B if it fails.
import * as Control_Alt from "../Control.Alt/index.js";
import * as Control_Alternative from "../Control.Alternative/index.js";
import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Category from "../Control.Category/index.js";
import * as Control_Monad_Reader_Class from "../Control.Monad.Reader.Class/index.js";
import * as Data_Array from "../Data.Array/index.js";
import * as Data_Array_NonEmpty from "../Data.Array.NonEmpty/index.js";
import * as Data_Array_NonEmpty_Internal from "../Data.Array.NonEmpty.Internal/index.js";
import * as Data_Boolean from "../Data.Boolean/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_FoldableWithIndex from "../Data.FoldableWithIndex/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_FunctorWithIndex from "../Data.FunctorWithIndex/index.js";
import * as Data_HeytingAlgebra from "../Data.HeytingAlgebra/index.js";
import * as Data_Map from "../Data.Map/index.js";
import * as Data_Map_Internal from "../Data.Map.Internal/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Monoid_Additive from "../Data.Monoid.Additive/index.js";
import * as Data_Newtype from "../Data.Newtype/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Ordering from "../Data.Ordering/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Semigroup_First from "../Data.Semigroup.First/index.js";
import * as Data_Semigroup_Foldable from "../Data.Semigroup.Foldable/index.js";
import * as Data_Semiring from "../Data.Semiring/index.js";
import * as Data_Set from "../Data.Set/index.js";
import * as Data_Traversable from "../Data.Traversable/index.js";
import * as Data_TraversableWithIndex from "../Data.TraversableWithIndex/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Data_Unfoldable from "../Data.Unfoldable/index.js";
import * as Partial_Unsafe from "../Partial.Unsafe/index.js";
import * as PureScript_Backend_Optimizer_Analysis from "../PureScript.Backend.Optimizer.Analysis/index.js";
import * as PureScript_Backend_Optimizer_CoreFn from "../PureScript.Backend.Optimizer.CoreFn/index.js";
import * as PureScript_Backend_Optimizer_Directives from "../PureScript.Backend.Optimizer.Directives/index.js";
import * as PureScript_Backend_Optimizer_Semantics from "../PureScript.Backend.Optimizer.Semantics/index.js";
import * as PureScript_Backend_Optimizer_Syntax from "../PureScript.Backend.Optimizer.Syntax/index.js";
import * as PureScript_Backend_Optimizer_Utils from "../PureScript.Backend.Optimizer.Utils/index.js";
import * as Safe_Coerce from "../Safe.Coerce/index.js";
var $runtime_lazy = function (name, moduleName, init) {
    var state = 0;
    var val;
    return function (lineNumber) {
        if (state === 2) return val;
        if (state === 1) throw new ReferenceError(name + " was needed before it finished initializing (module " + moduleName + ", line " + lineNumber + ")", moduleName, lineNumber);
        state = 1;
        val = init();
        state = 2;
        return val;
    };
};
var eqArray = /* #__PURE__ */ Data_Eq.eqArray(Data_Eq.eqString);
var eqQualified = /* #__PURE__ */ PureScript_Backend_Optimizer_CoreFn.eqQualified(PureScript_Backend_Optimizer_CoreFn.eqProperName);
var eqQualified1 = /* #__PURE__ */ PureScript_Backend_Optimizer_CoreFn.eqQualified(PureScript_Backend_Optimizer_CoreFn.eqIdent);
var ordArray = /* #__PURE__ */ Data_Ord.ordArray(Data_Ord.ordString);
var ordQualified = /* #__PURE__ */ PureScript_Backend_Optimizer_CoreFn.ordQualified(PureScript_Backend_Optimizer_CoreFn.ordProperName);
var ordQualified1 = /* #__PURE__ */ PureScript_Backend_Optimizer_CoreFn.ordQualified(PureScript_Backend_Optimizer_CoreFn.ordIdent);
var fromJust = /* #__PURE__ */ Data_Maybe.fromJust();
var monoidSemigroupMap = /* #__PURE__ */ Data_Map.monoidSemigroupMap(PureScript_Backend_Optimizer_CoreFn.ordIdent)(Data_Semigroup_First.semigroupFirst);
var foldl = /* #__PURE__ */ Data_Foldable.foldl(Data_Foldable.foldableArray);
var semigroupSet = /* #__PURE__ */ Data_Set.semigroupSet(Data_Ord.ordString);
var insert = /* #__PURE__ */ Data_Map_Internal.insert(PureScript_Backend_Optimizer_CoreFn.ordIdent);
var monoidRecord = /* #__PURE__ */ Data_Monoid.monoidRecord();
var monoidRecord1 = /* #__PURE__ */ monoidRecord(/* #__PURE__ */ Data_Monoid.monoidRecordCons({
    reflectSymbol: function () {
        return "rowsNoMatch";
    }
})(Data_Monoid.monoidArray)()(/* #__PURE__ */ Data_Monoid.monoidRecordCons({
    reflectSymbol: function () {
        return "rowsWithMatch";
    }
})(Data_Monoid.monoidArray)()(Data_Monoid.monoidRecordNil)));
var pure = /* #__PURE__ */ Control_Applicative.pure(Control_Applicative.applicativeFn);
var ask = /* #__PURE__ */ Control_Monad_Reader_Class.ask(Control_Monad_Reader_Class.monadAskFun);
var identity = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var sort = /* #__PURE__ */ Data_Array.sort(PureScript_Backend_Optimizer_CoreFn.ordIdent);
var ask1 = /* #__PURE__ */ Control_Monad_Reader_Class.ask(Control_Monad_Reader_Class.monadAskFun);
var ask2 = /* #__PURE__ */ Control_Monad_Reader_Class.ask(Control_Monad_Reader_Class.monadAskFun);
var join = /* #__PURE__ */ Control_Bind.join(Control_Bind.bindFn);
var semigroupSemigroupMap = /* #__PURE__ */ Data_Map.semigroupSemigroupMap(PureScript_Backend_Optimizer_CoreFn.ordIdent)(Data_Semigroup_First.semigroupFirst);
var toUnfoldable = /* #__PURE__ */ Data_Map_Internal.toUnfoldable(Data_Unfoldable.unfoldableArray);
var coerce = /* #__PURE__ */ Safe_Coerce.coerce();
var conj = /* #__PURE__ */ Data_HeytingAlgebra.conj(Data_HeytingAlgebra.heytingAlgebraBoolean);
var pure1 = /* #__PURE__ */ Control_Applicative.pure(Control_Applicative.applicativeArray);
var fromFoldable = /* #__PURE__ */ Data_Map_Internal.fromFoldable(PureScript_Backend_Optimizer_CoreFn.ordProperName)(Data_Foldable.foldableArray);
var fromFoldable1 = /* #__PURE__ */ Data_Map_Internal.fromFoldable(PureScript_Backend_Optimizer_CoreFn.ordIdent)(Data_Array_NonEmpty_Internal.foldableNonEmptyArray);
var ordQualified2 = /* #__PURE__ */ PureScript_Backend_Optimizer_CoreFn.ordQualified(PureScript_Backend_Optimizer_CoreFn.ordIdent);
var semigroupSet1 = /* #__PURE__ */ Data_Set.semigroupSet(ordQualified2);
var monoidSet = /* #__PURE__ */ Data_Set.monoidSet(ordQualified2);
var toUnfoldable1 = /* #__PURE__ */ Data_Map_Internal.toUnfoldable(Data_Unfoldable.unfoldableArray);
var PatWild = /* #__PURE__ */ (function () {
    function PatWild() {

    };
    PatWild.value = new PatWild();
    return PatWild;
})();
var PatRecord = /* #__PURE__ */ (function () {
    function PatRecord(value0) {
        this.value0 = value0;
    };
    PatRecord.create = function (value0) {
        return new PatRecord(value0);
    };
    return PatRecord;
})();
var PatProduct = /* #__PURE__ */ (function () {
    function PatProduct(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    PatProduct.create = function (value0) {
        return function (value1) {
            return new PatProduct(value0, value1);
        };
    };
    return PatProduct;
})();
var PatArray = /* #__PURE__ */ (function () {
    function PatArray(value0) {
        this.value0 = value0;
    };
    PatArray.create = function (value0) {
        return new PatArray(value0);
    };
    return PatArray;
})();
var PatSum = /* #__PURE__ */ (function () {
    function PatSum(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    PatSum.create = function (value0) {
        return function (value1) {
            return new PatSum(value0, value1);
        };
    };
    return PatSum;
})();
var PatInt = /* #__PURE__ */ (function () {
    function PatInt(value0) {
        this.value0 = value0;
    };
    PatInt.create = function (value0) {
        return new PatInt(value0);
    };
    return PatInt;
})();
var PatNumber = /* #__PURE__ */ (function () {
    function PatNumber(value0) {
        this.value0 = value0;
    };
    PatNumber.create = function (value0) {
        return new PatNumber(value0);
    };
    return PatNumber;
})();
var PatString = /* #__PURE__ */ (function () {
    function PatString(value0) {
        this.value0 = value0;
    };
    PatString.create = function (value0) {
        return new PatString(value0);
    };
    return PatString;
})();
var PatChar = /* #__PURE__ */ (function () {
    function PatChar(value0) {
        this.value0 = value0;
    };
    PatChar.create = function (value0) {
        return new PatChar(value0);
    };
    return PatChar;
})();
var PatBoolean = /* #__PURE__ */ (function () {
    function PatBoolean(value0) {
        this.value0 = value0;
    };
    PatBoolean.create = function (value0) {
        return new PatBoolean(value0);
    };
    return PatBoolean;
})();

// | vars - the references introduced at this pattern.
// | pattern - the actual pattern match to test
// | subterms - the subterm patterns to match only once this pattern matches.
var Pattern = function (x) {
    return x;
};
var UnconditionalFn = /* #__PURE__ */ (function () {
    function UnconditionalFn(value0) {
        this.value0 = value0;
    };
    UnconditionalFn.create = function (value0) {
        return new UnconditionalFn(value0);
    };
    return UnconditionalFn;
})();
var GuardedFn = /* #__PURE__ */ (function () {
    function GuardedFn(value0) {
        this.value0 = value0;
    };
    GuardedFn.create = function (value0) {
        return new GuardedFn(value0);
    };
    return GuardedFn;
})();
var newtypePattern_ = {
    Coercible0: function () {
        return undefined;
    }
};
var eqPatternCase = {
    eq: function (x) {
        return function (y) {
            if (x instanceof PatWild && y instanceof PatWild) {
                return true;
            };
            if (x instanceof PatRecord && y instanceof PatRecord) {
                return Data_Eq.eq(eqArray)(x.value0)(y.value0);
            };
            if (x instanceof PatProduct && y instanceof PatProduct) {
                return Data_Eq.eq(eqQualified)(x.value0)(y.value0) && Data_Eq.eq(eqQualified1)(x.value1)(y.value1);
            };
            if (x instanceof PatArray && y instanceof PatArray) {
                return x.value0 === y.value0;
            };
            if (x instanceof PatSum && y instanceof PatSum) {
                return Data_Eq.eq(eqQualified)(x.value0)(y.value0) && Data_Eq.eq(eqQualified1)(x.value1)(y.value1);
            };
            if (x instanceof PatInt && y instanceof PatInt) {
                return x.value0 === y.value0;
            };
            if (x instanceof PatNumber && y instanceof PatNumber) {
                return x.value0 === y.value0;
            };
            if (x instanceof PatString && y instanceof PatString) {
                return x.value0 === y.value0;
            };
            if (x instanceof PatChar && y instanceof PatChar) {
                return x.value0 === y.value0;
            };
            if (x instanceof PatBoolean && y instanceof PatBoolean) {
                return x.value0 === y.value0;
            };
            return false;
        };
    }
};
var eq = /* #__PURE__ */ Data_Eq.eq(eqPatternCase);
var ordPatternCase = {
    compare: function (x) {
        return function (y) {
            if (x instanceof PatWild && y instanceof PatWild) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof PatWild) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof PatWild) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof PatRecord && y instanceof PatRecord) {
                return Data_Ord.compare(ordArray)(x.value0)(y.value0);
            };
            if (x instanceof PatRecord) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof PatRecord) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof PatProduct && y instanceof PatProduct) {
                var v = Data_Ord.compare(ordQualified)(x.value0)(y.value0);
                if (v instanceof Data_Ordering.LT) {
                    return Data_Ordering.LT.value;
                };
                if (v instanceof Data_Ordering.GT) {
                    return Data_Ordering.GT.value;
                };
                return Data_Ord.compare(ordQualified1)(x.value1)(y.value1);
            };
            if (x instanceof PatProduct) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof PatProduct) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof PatArray && y instanceof PatArray) {
                return Data_Ord.compare(Data_Ord.ordInt)(x.value0)(y.value0);
            };
            if (x instanceof PatArray) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof PatArray) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof PatSum && y instanceof PatSum) {
                var v = Data_Ord.compare(ordQualified)(x.value0)(y.value0);
                if (v instanceof Data_Ordering.LT) {
                    return Data_Ordering.LT.value;
                };
                if (v instanceof Data_Ordering.GT) {
                    return Data_Ordering.GT.value;
                };
                return Data_Ord.compare(ordQualified1)(x.value1)(y.value1);
            };
            if (x instanceof PatSum) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof PatSum) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof PatInt && y instanceof PatInt) {
                return Data_Ord.compare(Data_Ord.ordInt)(x.value0)(y.value0);
            };
            if (x instanceof PatInt) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof PatInt) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof PatNumber && y instanceof PatNumber) {
                return Data_Ord.compare(Data_Ord.ordNumber)(x.value0)(y.value0);
            };
            if (x instanceof PatNumber) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof PatNumber) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof PatString && y instanceof PatString) {
                return Data_Ord.compare(Data_Ord.ordString)(x.value0)(y.value0);
            };
            if (x instanceof PatString) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof PatString) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof PatChar && y instanceof PatChar) {
                return Data_Ord.compare(Data_Ord.ordChar)(x.value0)(y.value0);
            };
            if (x instanceof PatChar) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof PatChar) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof PatBoolean && y instanceof PatBoolean) {
                return Data_Ord.compare(Data_Ord.ordBoolean)(x.value0)(y.value0);
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 0, column 0 - line 0, column 0): " + [ x.constructor.name, y.constructor.name ]);
        };
    },
    Eq0: function () {
        return eqPatternCase;
    }
};
var monoidRecord2 = /* #__PURE__ */ monoidRecord(/* #__PURE__ */ Data_Monoid.monoidRecordCons({
    reflectSymbol: function () {
        return "aScore";
    }
})(/* #__PURE__ */ Data_Monoid_Additive.monoidAdditive(Data_Semiring.semiringInt))()(/* #__PURE__ */ Data_Monoid.monoidRecordCons({
    reflectSymbol: function () {
        return "ctors";
    }
})(/* #__PURE__ */ Data_Set.monoidSet(ordPatternCase))()(/* #__PURE__ */ Data_Monoid.monoidRecordCons({
    reflectSymbol: function () {
        return "tailRowIndices";
    }
})(Data_Monoid.monoidArray)()(Data_Monoid.monoidRecordNil))));
var monoidSet1 = /* #__PURE__ */ Data_Set.monoidSet(ordPatternCase);
var topEnv = function (v) {
    return {
        currentModule: v.currentModule,
        evalExternRef: v.evalExternRef,
        evalExternSpine: v.evalExternSpine,
        directives: v.directives,
        locals: [  ]
    };
};
var toExternImpl = function (env) {
    return function (group) {
        return function (expr) {
            if (expr instanceof PureScript_Backend_Optimizer_Semantics.ExprSyntax && (expr.value1 instanceof PureScript_Backend_Optimizer_Syntax.Lit && expr.value1.value0 instanceof PureScript_Backend_Optimizer_CoreFn.LitRecord)) {
                var propsWithAnalysis = Data_Functor.map(Data_Functor.functorArray)(Data_Functor.map(PureScript_Backend_Optimizer_CoreFn.functorProp)(PureScript_Backend_Optimizer_Semantics.freeze))(expr.value1.value0.value0);
                return new Data_Tuple.Tuple(new Data_Tuple.Tuple(expr.value0, new PureScript_Backend_Optimizer_Semantics.ExternDict(group, propsWithAnalysis)), new PureScript_Backend_Optimizer_Syntax.Lit(new PureScript_Backend_Optimizer_CoreFn.LitRecord(Data_Functor.map(Data_Functor.functorArray)(Data_Functor.map(PureScript_Backend_Optimizer_CoreFn.functorProp)(Data_Tuple.snd))(propsWithAnalysis))));
            };
            if (expr instanceof PureScript_Backend_Optimizer_Semantics.ExprSyntax && expr.value1 instanceof PureScript_Backend_Optimizer_Syntax.CtorDef) {
                var v = PureScript_Backend_Optimizer_Semantics.freeze(expr);
                var meta = fromJust(Data_Map_Internal.lookup(PureScript_Backend_Optimizer_CoreFn.ordProperName)(expr.value1.value1)(env.dataTypes));
                return new Data_Tuple.Tuple(new Data_Tuple.Tuple(v.value0, new PureScript_Backend_Optimizer_Semantics.ExternCtor(meta, expr.value1.value0, expr.value1.value1, expr.value1.value2, expr.value1.value3)), v.value1);
            };
            var v = PureScript_Backend_Optimizer_Semantics.freeze(expr);
            return new Data_Tuple.Tuple(new Data_Tuple.Tuple(v.value0, new PureScript_Backend_Optimizer_Semantics.ExternExpr(group, v.value1)), v.value1);
        };
    };
};
var toCaseRowVars = function (v) {
    return Data_Foldable.foldMap(Data_Set.foldableSet)(monoidSemigroupMap)((function () {
        var $753 = Data_Function.flip(Data_Map_Internal.singleton)(v.column);
        return function ($754) {
            return Data_Map.SemigroupMap($753($754));
        };
    })())(v.pattern.vars);
};
var patternVars = function (v) {
    return Data_Semigroup.append(Data_Semigroup.semigroupArray)(Data_Set.toUnfoldable(Data_Unfoldable.unfoldableArray)(v.pattern.vars))(Data_Foldable.foldMap(Data_Foldable.foldableArray)(Data_Monoid.monoidArray)(patternVars)(v.pattern.subterms));
};
var patternSubterms = function (v) {
    return v.pattern.subterms;
};

// `patternCase` has a naming clash with record puns
var patternPatCase = function (v) {
    return v.pattern.patternCase;
};
var normalizeCaseRows = /* #__PURE__ */ (function () {
    var normalizeProps = function (allFieldNames) {
        var addBinders = function (allFieldsSet) {
            return function (pat) {
                var v = patternPatCase(pat);
                if (v instanceof PatRecord) {
                    var currentFieldsWithSubterms = Data_Array.zip(v.value0)(patternSubterms(pat));
                    var allFieldsWithWildSubterms = Data_Functor.mapFlipped(Data_Functor.functorArray)(Data_Set.toUnfoldable(Data_Unfoldable.unfoldableArray)(allFieldsSet))(function (fieldName) {
                        return new Data_Tuple.Tuple(fieldName, {
                            accessor: new PureScript_Backend_Optimizer_Syntax.GetProp(fieldName),
                            pattern: {
                                vars: Data_Set.empty,
                                patternCase: PatWild.value,
                                subterms: [  ]
                            }
                        });
                    });
                    var v1 = Data_Array.unzip(Data_Functor.map(Data_Functor.functorArray)(Data_Array_NonEmpty.head)(Data_Array.groupAllBy(Data_Ord.comparing(Data_Ord.ordString)(Data_Tuple.fst))(Data_Semigroup.append(Data_Semigroup.semigroupArray)(currentFieldsWithSubterms)(allFieldsWithWildSubterms))));
                    return {
                        column: pat.column,
                        pattern: Data_Newtype.over()()(Pattern)(function (v3) {
                            return {
                                vars: v3.vars,
                                patternCase: new PatRecord(v1.value0),
                                subterms: v1.value1
                            };
                        })(pat.pattern)
                    };
                };
                return pat;
            };
        };
        return Data_Functor.map(Data_Functor.functorArray)(function (nextRow) {
            return {
                guardFn: nextRow.guardFn,
                vars: nextRow.vars,
                patterns: Data_Array.zipWith(addBinders)(allFieldNames)(nextRow.patterns)
            };
        });
    };
    var columnProps = function (caseRows) {
        var go = function ($copy_columnIdx) {
            return function ($copy_columnsAcc) {
                var $tco_var_columnIdx = $copy_columnIdx;
                var $tco_done = false;
                var $tco_result;
                function $tco_loop(columnIdx, columnsAcc) {
                    var nextColumnFields = Data_Function.flip(foldl)(Data_Maybe.Nothing.value)(function (acc) {
                        return function (next) {
                            return Control_Bind.bind(Data_Maybe.bindMaybe)(Data_Array.index(next.patterns)(columnIdx))(function (pat) {
                                return Control_Applicative.pure(Data_Maybe.applicativeMaybe)((function () {
                                    var v = patternPatCase(pat);
                                    if (v instanceof PatRecord) {
                                        var keys = Data_Set.fromFoldable(Data_Foldable.foldableArray)(Data_Ord.ordString)(v.value0);
                                        return Data_Maybe.maybe(keys)(Data_Semigroup.append(semigroupSet)(keys))(acc);
                                    };
                                    return Data_Maybe.fromMaybe(Data_Set.empty)(acc);
                                })());
                            });
                        };
                    })(caseRows);
                    if (nextColumnFields instanceof Data_Maybe.Nothing) {
                        $tco_done = true;
                        return columnsAcc;
                    };
                    if (nextColumnFields instanceof Data_Maybe.Just) {
                        $tco_var_columnIdx = columnIdx + 1 | 0;
                        $copy_columnsAcc = Data_Array.snoc(columnsAcc)(nextColumnFields.value0);
                        return;
                    };
                    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 716, column 7 - line 718, column 63): " + [ nextColumnFields.constructor.name ]);
                };
                while (!$tco_done) {
                    $tco_result = $tco_loop($tco_var_columnIdx, $copy_columnsAcc);
                };
                return $tco_result;
            };
        };
        return go(0)([  ]);
    };
    return Control_Bind.bindFlipped(Control_Bind.bindFn)(normalizeProps)(columnProps);
})();
var makeExternEvalSpine = function (conv) {
    return function (env) {
        return function (qual) {
            return function (spine) {
                var result = Control_Bind.bind(Data_Maybe.bindMaybe)(Data_Map_Internal.lookup(ordQualified1)(qual)(conv.foreignSemantics))(function (fn) {
                    return fn(env)(qual)(spine);
                });
                if (result instanceof Data_Maybe.Nothing) {
                    return Control_Bind.bind(Data_Maybe.bindMaybe)(Data_Map_Internal.lookup(ordQualified1)(qual)(conv.implementations))(function (impl) {
                        return PureScript_Backend_Optimizer_Semantics.evalExternFromImpl(topEnv(env))(qual)(impl)(spine);
                    });
                };
                return result;
            };
        };
    };
};
var makeExternEvalRef = function (conv) {
    return function (env) {
        return function (qual) {
            return Data_Functor.map(Data_Maybe.functorMaybe)(PureScript_Backend_Optimizer_Semantics.evalExternRefFromImpl(env)(qual))(Data_Map_Internal.lookup(ordQualified1)(qual)(conv.implementations));
        };
    };
};
var levelUp = function (f) {
    return function (env) {
        return f({
            analyzeCustom: env.analyzeCustom,
            currentModule: env.currentModule,
            dataTypes: env.dataTypes,
            toLevel: env.toLevel,
            implementations: env.implementations,
            moduleImplementations: env.moduleImplementations,
            optimizationSteps: env.optimizationSteps,
            directives: env.directives,
            foreignSemantics: env.foreignSemantics,
            rewriteLimit: env.rewriteLimit,
            traceIdents: env.traceIdents,
            currentLevel: env.currentLevel + 1 | 0
        });
    };
};
var intro = function (dictFoldable) {
    return function (ident) {
        return function (lvl) {
            return function (f) {
                return function (env) {
                    return f({
                        analyzeCustom: env.analyzeCustom,
                        currentModule: env.currentModule,
                        dataTypes: env.dataTypes,
                        implementations: env.implementations,
                        moduleImplementations: env.moduleImplementations,
                        optimizationSteps: env.optimizationSteps,
                        directives: env.directives,
                        foreignSemantics: env.foreignSemantics,
                        rewriteLimit: env.rewriteLimit,
                        traceIdents: env.traceIdents,
                        currentLevel: env.currentLevel + 1 | 0,
                        toLevel: Data_Foldable.foldr(dictFoldable)(Data_Function.flip(insert)(lvl))(env.toLevel)(ident)
                    });
                };
            };
        };
    };
};
var inferTransitiveDirective = function (directives) {
    return function (impl) {
        return function (backendExpr) {
            return function (cfn) {
                var fromImpl = (function () {
                    if (impl instanceof PureScript_Backend_Optimizer_Semantics.ExternExpr && (impl.value1 instanceof PureScript_Backend_Optimizer_Syntax.App && impl.value1.value0 instanceof PureScript_Backend_Optimizer_Syntax.Var)) {
                        var v = Data_Map_Internal.lookup(PureScript_Backend_Optimizer_Semantics.ordEvalRef)(new PureScript_Backend_Optimizer_Semantics.EvalExtern(impl.value1.value0.value0))(directives);
                        if (v instanceof Data_Maybe.Just) {
                            var newDirs = Data_FoldableWithIndex.foldrWithIndex(Data_Map_Internal.foldableWithIndexMap)(function (ix) {
                                return function (dir) {
                                    return function (accum) {
                                        if (ix instanceof PureScript_Backend_Optimizer_Semantics.InlineRef && dir instanceof PureScript_Backend_Optimizer_Semantics.InlineArity) {
                                            return Data_Map_Internal.insert(PureScript_Backend_Optimizer_Semantics.ordInlineAccessor)(PureScript_Backend_Optimizer_Semantics.InlineRef.value)(new PureScript_Backend_Optimizer_Semantics.InlineArity(dir.value0 - Data_Array_NonEmpty.length(impl.value1.value1) | 0))(accum);
                                        };
                                        if (ix instanceof PureScript_Backend_Optimizer_Semantics.InlineSpineProp) {
                                            return Data_Map_Internal.insert(PureScript_Backend_Optimizer_Semantics.ordInlineAccessor)(new PureScript_Backend_Optimizer_Semantics.InlineSpineProp(ix.value0))(dir)(Data_Map_Internal.insert(PureScript_Backend_Optimizer_Semantics.ordInlineAccessor)(new PureScript_Backend_Optimizer_Semantics.InlineProp(ix.value0))(dir)(accum));
                                        };
                                        return accum;
                                    };
                                };
                            })(Data_Map_Internal.empty)(v.value0);
                            var $358 = Data_Map_Internal.isEmpty(newDirs);
                            if ($358) {
                                return Data_Maybe.Nothing.value;
                            };
                            return new Data_Maybe.Just(newDirs);
                        };
                        return Data_Maybe.Nothing.value;
                    };
                    if (impl instanceof PureScript_Backend_Optimizer_Semantics.ExternExpr && (impl.value1 instanceof PureScript_Backend_Optimizer_Syntax.Accessor && (impl.value1.value0 instanceof PureScript_Backend_Optimizer_Syntax.App && (impl.value1.value0.value0 instanceof PureScript_Backend_Optimizer_Syntax.Var && impl.value1.value1 instanceof PureScript_Backend_Optimizer_Syntax.GetProp)))) {
                        var v = Control_Bind.bind(Data_Maybe.bindMaybe)(Data_Map_Internal.lookup(PureScript_Backend_Optimizer_Semantics.ordEvalRef)(new PureScript_Backend_Optimizer_Semantics.EvalExtern(impl.value1.value0.value0.value0))(directives))(Data_Map_Internal.lookup(PureScript_Backend_Optimizer_Semantics.ordInlineAccessor)(new PureScript_Backend_Optimizer_Semantics.InlineSpineProp(impl.value1.value1.value0)));
                        if (v instanceof Data_Maybe.Just && v.value0 instanceof PureScript_Backend_Optimizer_Semantics.InlineArity) {
                            return new Data_Maybe.Just(Data_Map_Internal.singleton(PureScript_Backend_Optimizer_Semantics.InlineRef.value)(new PureScript_Backend_Optimizer_Semantics.InlineArity(v.value0.value0)));
                        };
                        return Data_Maybe.Nothing.value;
                    };
                    return Data_Maybe.Nothing.value;
                })();
                var fromBackendExpr = (function () {
                    if (backendExpr instanceof PureScript_Backend_Optimizer_Semantics.ExprSyntax && (backendExpr.value1 instanceof PureScript_Backend_Optimizer_Syntax.App && (backendExpr.value1.value0 instanceof PureScript_Backend_Optimizer_Semantics.ExprSyntax && backendExpr.value1.value0.value1 instanceof PureScript_Backend_Optimizer_Syntax.Var))) {
                        var v = Control_Bind.bind(Data_Maybe.bindMaybe)(Data_Map_Internal.lookup(PureScript_Backend_Optimizer_Semantics.ordEvalRef)(new PureScript_Backend_Optimizer_Semantics.EvalExtern(backendExpr.value1.value0.value1.value0))(directives))(Data_Map_Internal.lookup(PureScript_Backend_Optimizer_Semantics.ordInlineAccessor)(PureScript_Backend_Optimizer_Semantics.InlineRef.value));
                        var v1 = function (v2) {
                            return Data_Maybe.Nothing.value;
                        };
                        if (v instanceof Data_Maybe.Just && v.value0 instanceof PureScript_Backend_Optimizer_Semantics.InlineArity) {
                            if (cfn instanceof PureScript_Backend_Optimizer_CoreFn.ExprApp && (cfn.value0.meta instanceof Data_Maybe.Just && cfn.value0.meta.value0 instanceof PureScript_Backend_Optimizer_CoreFn.IsSyntheticApp)) {
                                var $379 = Data_Array_NonEmpty.length(backendExpr.value1.value1);
                                var $380 = $379 >= v.value0.value0;
                                if ($380) {
                                    return new Data_Maybe.Just(Data_Map_Internal.singleton(PureScript_Backend_Optimizer_Semantics.InlineRef.value)(PureScript_Backend_Optimizer_Semantics.InlineAlways.value));
                                };
                                return v1(true);
                            };
                            return v1(true);
                        };
                        return v1(true);
                    };
                    return Data_Maybe.Nothing.value;
                })();
                return Control_Alt.alt(Data_Maybe.altMaybe)(fromImpl)(fromBackendExpr);
            };
        };
    };
};
var guardTag = function (n) {
    return function (lhs) {
        return new PureScript_Backend_Optimizer_Syntax.PrimOp(new PureScript_Backend_Optimizer_Syntax.Op1(new PureScript_Backend_Optimizer_Syntax.OpIsTag(n), lhs));
    };
};
var getReturnType = function ($copy_v) {
    var $tco_done = false;
    var $tco_result;
    function $tco_loop(v) {
        if (v instanceof PureScript_Backend_Optimizer_CoreFn.ForAll) {
            $copy_v = v.value1;
            return;
        };
        if (v instanceof PureScript_Backend_Optimizer_CoreFn.ConstrainedType) {
            $copy_v = v.value1;
            return;
        };
        if (v instanceof PureScript_Backend_Optimizer_CoreFn.Func) {
            $tco_done = true;
            return new Data_Maybe.Just(v.value1);
        };
        $tco_done = true;
        return Data_Maybe.Nothing.value;
    };
    while (!$tco_done) {
        $tco_result = $tco_loop($copy_v);
    };
    return $tco_result;
};
var inferExprType = function (v) {
    if (v instanceof PureScript_Backend_Optimizer_CoreFn.ExprApp) {
        var v1 = PureScript_Backend_Optimizer_CoreFn.exprAnn(v.value1);
        if (v1.type instanceof Data_Maybe.Just) {
            return getReturnType(v1.type.value0);
        };
        var v2 = inferExprType(v.value1);
        if (v2 instanceof Data_Maybe.Just) {
            return getReturnType(v2.value0);
        };
        if (v2 instanceof Data_Maybe.Nothing) {
            return Data_Maybe.Nothing.value;
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 1037, column 8 - line 1039, column 28): " + [ v2.constructor.name ]);
    };
    return Data_Maybe.Nothing.value;
};
var getCtx = function (env) {
    var lookupExtern = function (qual) {
        return function (acc) {
            return Control_Bind.bind(Data_Maybe.bindMaybe)(Data_Map_Internal.lookup(ordQualified1)(qual)(env.implementations))(function (v) {
                if (v.value1 instanceof PureScript_Backend_Optimizer_Semantics.ExternExpr) {
                    if (acc instanceof Data_Maybe.Nothing) {
                        return new Data_Maybe.Just(v.value0);
                    };
                    return Data_Maybe.Nothing.value;
                };
                if (v.value1 instanceof PureScript_Backend_Optimizer_Semantics.ExternDict) {
                    if (acc instanceof Data_Maybe.Just) {
                        return Data_Functor.map(Data_Maybe.functorMaybe)(Data_Tuple.fst)(PureScript_Backend_Optimizer_CoreFn.findProp(acc.value0)(v.value1.value1));
                    };
                    if (acc instanceof Data_Maybe.Nothing) {
                        return new Data_Maybe.Just(v.value0);
                    };
                    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 395, column 9 - line 399, column 19): " + [ acc.constructor.name ]);
                };
                if (v.value1 instanceof PureScript_Backend_Optimizer_Semantics.ExternCtor) {
                    return Data_Maybe.Nothing.value;
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 389, column 5 - line 401, column 16): " + [ v.value1.constructor.name ]);
            });
        };
    };
    return {
        currentLevel: env.currentLevel,
        lookupExtern: lookupExtern,
        analyze: function (v) {
            return function (expr) {
                var v1 = env.analyzeCustom(v)(expr);
                if (v1 instanceof Data_Maybe.Just) {
                    return v1.value0;
                };
                if (v1 instanceof Data_Maybe.Nothing) {
                    if (v.effect) {
                        return PureScript_Backend_Optimizer_Analysis.analyzeEffectBlock(PureScript_Backend_Optimizer_Semantics.hasAnalysisBackendExpr)(PureScript_Backend_Optimizer_Semantics.hasSyntaxBackendExpr)(lookupExtern)(expr);
                    };
                    return PureScript_Backend_Optimizer_Analysis.analyze(PureScript_Backend_Optimizer_Semantics.hasAnalysisBackendExpr)(PureScript_Backend_Optimizer_Semantics.hasSyntaxBackendExpr)(lookupExtern)(expr);
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 377, column 7 - line 383, column 38): " + [ v1.constructor.name ]);
            };
        },
        effect: false
    };
};
var fromExternImpl = function (v) {
    if (v instanceof PureScript_Backend_Optimizer_Semantics.ExternExpr) {
        return new Data_Maybe.Just(v.value1);
    };
    if (v instanceof PureScript_Backend_Optimizer_Semantics.ExternDict) {
        return Data_Maybe.Nothing.value;
    };
    if (v instanceof PureScript_Backend_Optimizer_Semantics.ExternCtor) {
        return Data_Maybe.Nothing.value;
    };
    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 404, column 18 - line 407, column 34): " + [ v.constructor.name ]);
};
var decompose = function (chosenColumn) {
    var mergeResults = function (l) {
        return function (r) {
            if (l.match instanceof Data_Maybe.Just && r.match instanceof Data_Maybe.Just) {
                return Partial_Unsafe.unsafeCrashWith("mergeResults - impossible: cannot match the same column twice in the same row");
            };
            if (l.match instanceof Data_Maybe.Nothing && r.match instanceof Data_Maybe.Nothing) {
                return {
                    match: r.match,
                    nonMatchesBefore: Data_Semigroup.append(Data_Semigroup.semigroupArray)(l.nonMatchesBefore)(r.nonMatchesBefore)
                };
            };
            if (l.match instanceof Data_Maybe.Nothing && r.match instanceof Data_Maybe.Just) {
                return {
                    match: r.match,
                    nonMatchesBefore: Data_Semigroup.append(Data_Semigroup.semigroupArray)(l.nonMatchesBefore)(r.nonMatchesBefore)
                };
            };
            if (l.match instanceof Data_Maybe.Just && r.match instanceof Data_Maybe.Nothing) {
                return {
                    nonMatchesBefore: l.nonMatchesBefore,
                    match: new Data_Maybe.Just({
                        match: l.match.value0.match,
                        nonMatchesAfter: Data_Semigroup.append(Data_Semigroup.semigroupArray)(l.match.value0.nonMatchesAfter)(r.nonMatchesBefore)
                    })
                };
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 964, column 22 - line 968, column 123): " + [ l.match.constructor.name, r.match.constructor.name ]);
        };
    };
    var checkMatch = function (p) {
        var v = function (v1) {
            if (Data_Boolean.otherwise) {
                return {
                    nonMatchesBefore: [ p ],
                    match: Data_Maybe.Nothing.value
                };
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 930, column 1 - line 930, column 119): " + [ p.constructor.name ]);
        };
        var $450 = Data_Eq.eq(PureScript_Backend_Optimizer_Syntax.eqLevel)(p.column)(chosenColumn.column);
        if ($450) {
            var $451 = Data_Eq.eq(eqPatternCase)(patternPatCase(p))(PatWild.value) || Data_Function.on(eq)(patternPatCase)(chosenColumn)(p);
            if ($451) {
                return {
                    nonMatchesBefore: [  ],
                    match: new Data_Maybe.Just({
                        match: p,
                        nonMatchesAfter: [  ]
                    })
                };
            };
            return v(true);
        };
        return v(true);
    };
    return Data_Foldable.foldMap(Data_Foldable.foldableArray)(monoidRecord1)(function (row) {
        var v = Data_Array_NonEmpty.fromArray(row.patterns);
        if (v instanceof Data_Maybe.Nothing) {
            return Partial_Unsafe.unsafeCrashWith("decompose - nextRow.patterns cannot be empty since the first row contains at least one `PatCtor` patternCase");
        };
        if (v instanceof Data_Maybe.Just) {
            var v1 = PureScript_Backend_Optimizer_Utils.foldl1Array(function (l) {
                var $755 = mergeResults(l);
                return function ($756) {
                    return $755(checkMatch($756));
                };
            })(checkMatch)(v.value0);
            if (v1.match instanceof Data_Maybe.Just) {
                return {
                    rowsWithMatch: [ {
                        guardFn: row.guardFn,
                        vars: row.vars,
                        nonMatchesBefore: v1.nonMatchesBefore,
                        match: v1.match.value0.match,
                        nonMatchesAfter: v1.match.value0.nonMatchesAfter
                    } ],
                    rowsNoMatch: (function () {
                        var $455 = Data_Eq.eq(eqPatternCase)(patternPatCase(v1.match.value0.match))(PatWild.value);
                        if ($455) {
                            return [ row ];
                        };
                        return [  ];
                    })()
                };
            };
            if (v1.match instanceof Data_Maybe.Nothing) {
                return {
                    rowsWithMatch: [  ],
                    rowsNoMatch: [ row ]
                };
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 936, column 7 - line 944, column 12): " + [ v1.match.constructor.name ]);
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 932, column 3 - line 944, column 12): " + [ v.constructor.name ]);
    });
};
var currentLevel = function (env) {
    return env.currentLevel;
};
var chooseNextPattern = function (row0Patterns) {
    return function (tailRows) {
        var maximumByAll = function (dictFoldable) {
            return function (f) {
                var keepAllMax = function (acc) {
                    return function (next) {
                        if (acc instanceof Data_Maybe.Nothing) {
                            return new Data_Maybe.Just(Data_Array_NonEmpty.singleton(next));
                        };
                        if (acc instanceof Data_Maybe.Just) {
                            var v = f(Data_Array_NonEmpty.head(acc.value0))(next);
                            if (v instanceof Data_Ordering.GT) {
                                return acc;
                            };
                            if (v instanceof Data_Ordering.EQ) {
                                return new Data_Maybe.Just(Data_Array_NonEmpty.snoc(acc.value0)(next));
                            };
                            if (v instanceof Data_Ordering.LT) {
                                return new Data_Maybe.Just(Data_Array_NonEmpty.singleton(next));
                            };
                            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 820, column 19 - line 823, column 52): " + [ v.constructor.name ]);
                        };
                        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 818, column 7 - line 823, column 52): " + [ acc.constructor.name ]);
                    };
                };
                return Data_Foldable.foldl(dictFoldable)(keepAllMax)(Data_Maybe.Nothing.value);
            };
        };
        var expandIfPossible = Data_Array_NonEmpty.findMap(function (v) {
            var v1 = patternPatCase(v.value1);
            if (v1 instanceof PatRecord) {
                return new Data_Maybe.Just(v.value1);
            };
            if (v1 instanceof PatProduct) {
                return new Data_Maybe.Just(v.value1);
            };
            return Data_Maybe.Nothing.value;
        })(row0Patterns);
        if (expandIfPossible instanceof Data_Maybe.Just) {
            return expandIfPossible.value0;
        };
        if (expandIfPossible instanceof Data_Maybe.Nothing) {
            var matchingPatternGroups = Data_Functor.mapFlipped(Data_Array_NonEmpty_Internal.functorNonEmptyArray)(row0Patterns)(function (v) {
                var matchingCols = Data_FoldableWithIndex.foldMapWithIndex(Data_FoldableWithIndex.foldableWithIndexArray)(monoidRecord2)(function (rowIdx) {
                    return function (row) {
                        var v1 = Data_Array.index(row.patterns)(v.value0);
                        if (v1 instanceof Data_Maybe.Nothing) {
                            return Partial_Unsafe.unsafeCrashWith("Impossible: rows' column lengths differ in pattern match");
                        };
                        if (v1 instanceof Data_Maybe.Just) {
                            return {
                                tailRowIndices: Data_Monoid.guard(Data_Monoid.monoidArray)(Data_Function.on(eq)(patternPatCase)(v.value1)(v1.value0))([ rowIdx + 1 | 0 ]),
                                ctors: Data_Monoid.guard(monoidSet1)(Data_Eq.notEq(eqPatternCase)(patternPatCase(v1.value0))(PatWild.value))(Data_Set.singleton(patternPatCase(v1.value0))),
                                aScore: -Data_Array.length(Data_Array.filter((function () {
                                    var $757 = Data_Eq.notEq(eqPatternCase)(PatWild.value);
                                    return function ($758) {
                                        return $757(patternPatCase($758));
                                    };
                                })())(patternSubterms(v1.value0))) | 0
                            };
                        };
                        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 792, column 15 - line 799, column 20): " + [ v1.constructor.name ]);
                    };
                })(tailRows);
                return {
                    pattern: v.value1,
                    pScore: Data_Foldable.foldl(Data_Foldable.foldableArray)(function (l) {
                        return function (r) {
                            var $476 = (l + 1 | 0) === r;
                            if ($476) {
                                return r;
                            };
                            return l;
                        };
                    })(0)(matchingCols.tailRowIndices),
                    bScore: -Data_Set.size(Data_Set.insert(ordPatternCase)(patternPatCase(v.value1))(matchingCols.ctors)) | 0,
                    aScore: Data_Newtype.unwrap()(matchingCols.aScore)
                };
            });
            var heuristic = Control_Bind.composeKleisli(Data_Maybe.bindMaybe)(maximumByAll(Data_Array_NonEmpty_Internal.foldableNonEmptyArray)(Data_Ord.comparing(Data_Ord.ordInt)(function (v) {
                return v.pScore;
            })))(Control_Bind.composeKleisli(Data_Maybe.bindMaybe)(maximumByAll(Data_Array_NonEmpty_Internal.foldableNonEmptyArray)(Data_Ord.comparing(Data_Ord.ordInt)(function (v) {
                return v.bScore;
            })))((function () {
                var $759 = Data_Functor.map(Data_Maybe.functorMaybe)(function ($762) {
                    return (function (v) {
                        return v.pattern;
                    })(Data_Array_NonEmpty.head($762));
                });
                var $760 = maximumByAll(Data_Array_NonEmpty_Internal.foldableNonEmptyArray)(Data_Ord.comparing(Data_Ord.ordInt)(function (v) {
                    return v.aScore;
                }));
                return function ($761) {
                    return $759($760($761));
                };
            })()));
            var v = heuristic(matchingPatternGroups);
            if (v instanceof Data_Maybe.Just) {
                return v.value0;
            };
            if (v instanceof Data_Maybe.Nothing) {
                return Data_Tuple.snd(Data_Array_NonEmpty.head(row0Patterns));
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 810, column 7 - line 812, column 57): " + [ v.constructor.name ]);
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 783, column 3 - line 812, column 57): " + [ expandIfPossible.constructor.name ]);
    };
};
var buildM = function (a) {
    return function (env) {
        return PureScript_Backend_Optimizer_Semantics.build(getCtx(env))(a);
    };
};
var make = function (a) {
    return Control_Bind.bindFlipped(Control_Bind.bindFn)(buildM)(Data_Traversable.sequence(PureScript_Backend_Optimizer_Syntax.traversableBackendSyntax)(Control_Applicative.applicativeFn)(a));
};
var guardBoolean = function (n) {
    return function (lhs) {
        return new PureScript_Backend_Optimizer_Syntax.PrimOp(new PureScript_Backend_Optimizer_Syntax.Op2(new PureScript_Backend_Optimizer_Syntax.OpBooleanOrd(PureScript_Backend_Optimizer_Syntax.OpEq.value), lhs, make(new PureScript_Backend_Optimizer_Syntax.Lit(new PureScript_Backend_Optimizer_CoreFn.LitBoolean(n)))));
    };
};
var guardChar = function (n) {
    return function (lhs) {
        return new PureScript_Backend_Optimizer_Syntax.PrimOp(new PureScript_Backend_Optimizer_Syntax.Op2(new PureScript_Backend_Optimizer_Syntax.OpCharOrd(PureScript_Backend_Optimizer_Syntax.OpEq.value), lhs, make(new PureScript_Backend_Optimizer_Syntax.Lit(new PureScript_Backend_Optimizer_CoreFn.LitChar(n)))));
    };
};
var guardInt = function (n) {
    return function (lhs) {
        return new PureScript_Backend_Optimizer_Syntax.PrimOp(new PureScript_Backend_Optimizer_Syntax.Op2(new PureScript_Backend_Optimizer_Syntax.OpIntOrd(PureScript_Backend_Optimizer_Syntax.OpEq.value), lhs, make(new PureScript_Backend_Optimizer_Syntax.Lit(new PureScript_Backend_Optimizer_CoreFn.LitInt(n)))));
    };
};
var guardArrayLength = function (n) {
    return function (lhs) {
        return guardInt(n)(make(new PureScript_Backend_Optimizer_Syntax.PrimOp(new PureScript_Backend_Optimizer_Syntax.Op1(PureScript_Backend_Optimizer_Syntax.OpArrayLength.value, lhs))));
    };
};
var guardNumber = function (n) {
    return function (lhs) {
        return new PureScript_Backend_Optimizer_Syntax.PrimOp(new PureScript_Backend_Optimizer_Syntax.Op2(new PureScript_Backend_Optimizer_Syntax.OpNumberOrd(PureScript_Backend_Optimizer_Syntax.OpEq.value), lhs, make(new PureScript_Backend_Optimizer_Syntax.Lit(new PureScript_Backend_Optimizer_CoreFn.LitNumber(n)))));
    };
};
var guardString = function (n) {
    return function (lhs) {
        return new PureScript_Backend_Optimizer_Syntax.PrimOp(new PureScript_Backend_Optimizer_Syntax.Op2(new PureScript_Backend_Optimizer_Syntax.OpStringOrd(PureScript_Backend_Optimizer_Syntax.OpEq.value), lhs, make(new PureScript_Backend_Optimizer_Syntax.Lit(new PureScript_Backend_Optimizer_CoreFn.LitString(n)))));
    };
};
var makeGuard = function (lvl) {
    return function (g) {
        return function (inner) {
            return function (def) {
                return make(new PureScript_Backend_Optimizer_Syntax.Branch(Data_Array_NonEmpty.singleton(new PureScript_Backend_Optimizer_Syntax.Pair(make(g(make(new PureScript_Backend_Optimizer_Syntax.Local(Data_Maybe.Nothing.value, lvl)))), inner)), def));
            };
        };
    };
};
var makeLet = function (id) {
    return function (a) {
        return function (k) {
            return Control_Bind.bind(Control_Bind.bindFn)(currentLevel)(function (lvl) {
                if (id instanceof Data_Maybe.Nothing) {
                    return make(new PureScript_Backend_Optimizer_Syntax.Let(id, lvl, a, levelUp(k(lvl))));
                };
                if (id instanceof Data_Maybe.Just) {
                    return make(new PureScript_Backend_Optimizer_Syntax.Let(id, lvl, a, intro(Data_Foldable.foldableArray)([ id.value0 ])(lvl)(k(lvl))));
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 976, column 3 - line 980, column 56): " + [ id.constructor.name ]);
            });
        };
    };
};
var makeUncurriedAbs = function (args) {
    return function (cb) {
        return Data_Foldable.foldr(Data_Foldable.foldableArray)(function (ident) {
            return function (next) {
                return function (tmps) {
                    return Control_Bind.bind(Control_Bind.bindFn)(currentLevel)(function (lvl) {
                        return intro(Data_Foldable.foldableArray)([ ident ])(lvl)(next(Data_Array.snoc(tmps)(new Data_Tuple.Tuple(new Data_Maybe.Just(ident), lvl))));
                    });
                };
            };
        })(function (tmps) {
            return make(new PureScript_Backend_Optimizer_Syntax.UncurriedAbs(tmps, cb(tmps)));
        })(args)([  ]);
    };
};
var patternFail = /* #__PURE__ */ (function () {
    return make(new PureScript_Backend_Optimizer_Syntax.Fail("Failed pattern match"));
})();
var $lazy_binderToPattern = /* #__PURE__ */ $runtime_lazy("binderToPattern", "PureScript.Backend.Optimizer.Convert", function () {
    var primitivePattern = function (patternCase) {
        return pure({
            vars: Data_Set.empty,
            patternCase: patternCase,
            subterms: [  ]
        });
    };
    var lookupCtorFields = function (ty) {
        return function (ctor) {
            var localCtorFields = function (dataTypes) {
                return Control_Bind.bind(Data_Maybe.bindMaybe)(Data_Map_Internal.lookup(PureScript_Backend_Optimizer_CoreFn.ordProperName)(PureScript_Backend_Optimizer_CoreFn.unQualified(ty))(dataTypes))(function (v) {
                    return Data_Functor.map(Data_Maybe.functorMaybe)(function (v1) {
                        return v1.fields;
                    })(Data_Map_Internal.lookup(PureScript_Backend_Optimizer_CoreFn.ordIdent)(PureScript_Backend_Optimizer_CoreFn.unQualified(ctor))(v.constructors));
                });
            };
            var importedCtorFields = function (implementations) {
                var v = Data_Map_Internal.lookup(ordQualified1)(ctor)(implementations);
                if (v instanceof Data_Maybe.Just && v.value0.value1 instanceof PureScript_Backend_Optimizer_Semantics.ExternCtor) {
                    return new Data_Maybe.Just(v.value0.value1.value4);
                };
                return Data_Maybe.Nothing.value;
            };
            return Control_Bind.bind(Control_Bind.bindFn)(ask)(function (v) {
                var v1 = Control_Alt.alt(Data_Maybe.altMaybe)(importedCtorFields(v.implementations))(localCtorFields(v.dataTypes));
                if (v1 instanceof Data_Maybe.Just) {
                    return Control_Applicative.pure(Control_Applicative.applicativeFn)(v1.value0);
                };
                if (v1 instanceof Data_Maybe.Nothing) {
                    return Partial_Unsafe.unsafeCrashWith("Invariant broken: could not determine pattern matched constructor's fields during conversion.");
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 664, column 5 - line 666, column 129): " + [ v1.constructor.name ]);
            });
        };
    };
    var ctorPattern = function (patternCase) {
        return function (args) {
            return function (buildAccessor) {
                return function (toBinder) {
                    return Data_Functor.map(Data_Functor.functorFn)(function (v) {
                        return {
                            vars: Data_Set.empty,
                            patternCase: patternCase,
                            subterms: v
                        };
                    })(Data_TraversableWithIndex.forWithIndex(Control_Applicative.applicativeFn)(Data_TraversableWithIndex.traversableWithIndexArray)(args)(function (idx) {
                        return function (nextArg) {
                            return Data_Functor.map(Data_Functor.functorFn)(function (v) {
                                return {
                                    accessor: buildAccessor(idx)(nextArg),
                                    pattern: v
                                };
                            })($lazy_binderToPattern(649)(toBinder(nextArg)));
                        };
                    }));
                };
            };
        };
    };
    return function (v) {
        if (v instanceof PureScript_Backend_Optimizer_CoreFn.BinderNull) {
            return primitivePattern(PatWild.value);
        };
        if (v instanceof PureScript_Backend_Optimizer_CoreFn.BinderVar) {
            return pure({
                vars: Data_Set.singleton(v.value1),
                patternCase: PatWild.value,
                subterms: [  ]
            });
        };
        if (v instanceof PureScript_Backend_Optimizer_CoreFn.BinderNamed) {
            return Data_Functor.map(Data_Functor.functorFn)(Data_Newtype.over()()(Pattern)(function (r) {
                return {
                    patternCase: r.patternCase,
                    subterms: r.subterms,
                    vars: Data_Set.insert(PureScript_Backend_Optimizer_CoreFn.ordIdent)(v.value1)(r.vars)
                };
            }))($lazy_binderToPattern(586)(v.value2));
        };
        if (v instanceof PureScript_Backend_Optimizer_CoreFn.BinderLit) {
            if (v.value1 instanceof PureScript_Backend_Optimizer_CoreFn.LitInt) {
                return primitivePattern(new PatInt(v.value1.value0));
            };
            if (v.value1 instanceof PureScript_Backend_Optimizer_CoreFn.LitNumber) {
                return primitivePattern(new PatNumber(v.value1.value0));
            };
            if (v.value1 instanceof PureScript_Backend_Optimizer_CoreFn.LitString) {
                return primitivePattern(new PatString(v.value1.value0));
            };
            if (v.value1 instanceof PureScript_Backend_Optimizer_CoreFn.LitChar) {
                return primitivePattern(new PatChar(v.value1.value0));
            };
            if (v.value1 instanceof PureScript_Backend_Optimizer_CoreFn.LitBoolean) {
                return primitivePattern(new PatBoolean(v.value1.value0));
            };
            if (v.value1 instanceof PureScript_Backend_Optimizer_CoreFn.LitArray) {
                return ctorPattern(new PatArray(Data_Array.length(v.value1.value0)))(v.value1.value0)(function (idx) {
                    return function (v1) {
                        return new PureScript_Backend_Optimizer_Syntax.GetIndex(idx);
                    };
                })(identity);
            };
            if (v.value1 instanceof PureScript_Backend_Optimizer_CoreFn.LitRecord) {
                return ctorPattern(new PatRecord(Data_Functor.map(Data_Functor.functorArray)(PureScript_Backend_Optimizer_CoreFn.propKey)(v.value1.value0)))(v.value1.value0)(function (v1) {
                    return function (p) {
                        return new PureScript_Backend_Optimizer_Syntax.GetProp(PureScript_Backend_Optimizer_CoreFn.propKey(p));
                    };
                })(PureScript_Backend_Optimizer_CoreFn.propValue);
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 587, column 22 - line 607, column 18): " + [ v.value1.constructor.name ]);
        };
        if (v instanceof PureScript_Backend_Optimizer_CoreFn.BinderConstructor) {
            var v1 = function (v2) {
                if (v.value0.meta instanceof Data_Maybe.Just && (v.value0.meta.value0 instanceof PureScript_Backend_Optimizer_CoreFn.IsNewtype && Data_Boolean.otherwise)) {
                    return Partial_Unsafe.unsafeCrashWith("Newtype binder didn't wrap 1 arg");
                };
                if (v.value0.meta instanceof Data_Maybe.Just && (v.value0.meta.value0 instanceof PureScript_Backend_Optimizer_CoreFn.IsConstructor && v.value0.meta.value0.value0 instanceof PureScript_Backend_Optimizer_CoreFn.ProductType)) {
                    return Control_Bind.bind(Control_Bind.bindFn)(lookupCtorFields(v.value1)(v.value2))(function (ctorFields) {
                        var argsWithNames = Data_Array.zip(v.value3)(ctorFields);
                        return ctorPattern(new PatProduct(v.value1, v.value2))(argsWithNames)(function (idx) {
                            return function (v3) {
                                return new PureScript_Backend_Optimizer_Syntax.GetCtorField(v.value2, PureScript_Backend_Optimizer_CoreFn.ProductType.value, PureScript_Backend_Optimizer_CoreFn.unQualified(v.value1), PureScript_Backend_Optimizer_CoreFn.unQualified(v.value2), v3.value1, idx);
                            };
                        })(Data_Tuple.fst);
                    });
                };
                if (v.value0.meta instanceof Data_Maybe.Just && (v.value0.meta.value0 instanceof PureScript_Backend_Optimizer_CoreFn.IsConstructor && v.value0.meta.value0.value0 instanceof PureScript_Backend_Optimizer_CoreFn.SumType)) {
                    return Control_Bind.bind(Control_Bind.bindFn)(lookupCtorFields(v.value1)(v.value2))(function (ctorFields) {
                        var argsWithNames = Data_Array.zip(v.value3)(ctorFields);
                        return ctorPattern(new PatSum(v.value1, v.value2))(argsWithNames)(function (idx) {
                            return function (v3) {
                                return new PureScript_Backend_Optimizer_Syntax.GetCtorField(v.value2, PureScript_Backend_Optimizer_CoreFn.SumType.value, PureScript_Backend_Optimizer_CoreFn.unQualified(v.value1), PureScript_Backend_Optimizer_CoreFn.unQualified(v.value2), v3.value1, idx);
                            };
                        })(Data_Tuple.fst);
                    });
                };
                return Partial_Unsafe.unsafeCrashWith("binderToPattern - invalid meta");
            };
            if (v.value0.meta instanceof Data_Maybe.Just && v.value0.meta.value0 instanceof PureScript_Backend_Optimizer_CoreFn.IsNewtype) {
                if (v.value3.length === 1) {
                    return $lazy_binderToPattern(613)(v["value3"][0]);
                };
                return v1(true);
            };
            return v1(true);
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 581, column 19 - line 635, column 55): " + [ v.constructor.name ]);
    };
});
var binderToPattern = /* #__PURE__ */ $lazy_binderToPattern(580);
var toBackendExpr = function (expr) {
    var go = (function () {
        var toInitialCaseRows = function (idents) {
            return function (alts) {
                return function (useCaseRowsCb) {
                    return Data_Foldable.foldr(Data_Foldable.foldableArray)(function (v) {
                        return function (mainCb) {
                            return function (caseRows) {
                                return Control_Bind.bind(Control_Bind.bindFn)(Data_Array.zipWithA(Control_Applicative.applicativeFn)(function (ident) {
                                    return function (b) {
                                        return Data_Functor.map(Data_Functor.functorFn)(function (v1) {
                                            return {
                                                column: ident,
                                                pattern: v1
                                            };
                                        })(binderToPattern(b));
                                    };
                                })(idents)(v.value0))(function (patterns) {
                                    var buildCaseRow = function (guardFn) {
                                        return {
                                            patterns: patterns,
                                            guardFn: guardFn,
                                            vars: Data_Map_Internal.empty
                                        };
                                    };
                                    var args = sort(Data_Foldable.foldMap(Data_Foldable.foldableArray)(Data_Monoid.monoidArray)(patternVars)(patterns));
                                    if (v.value1 instanceof PureScript_Backend_Optimizer_CoreFn.Unconditional) {
                                        return makeLet(Data_Maybe.Nothing.value)(makeUncurriedAbs(args)(function (v1) {
                                            return toBackendExpr(v.value1.value0);
                                        }))(function (tmp) {
                                            return mainCb(Data_Array.snoc(caseRows)(buildCaseRow(new UnconditionalFn(tmp))));
                                        });
                                    };
                                    if (v.value1 instanceof PureScript_Backend_Optimizer_CoreFn.Guarded) {
                                        return Data_Foldable.foldr(Data_Foldable.foldableArray)(function (v1) {
                                            return function (cb) {
                                                return function (xs) {
                                                    return makeLet(Data_Maybe.Nothing.value)(makeUncurriedAbs(args)(function (v2) {
                                                        return toBackendExpr(v1.value1);
                                                    }))(function (tmp) {
                                                        return cb(Data_Array.snoc(xs)(new Data_Tuple.Tuple(v1.value0, tmp)));
                                                    });
                                                };
                                            };
                                        })(function (xs) {
                                            var v1 = Data_Array_NonEmpty.fromArray(xs);
                                            if (v1 instanceof Data_Maybe.Nothing) {
                                                return Partial_Unsafe.unsafeCrashWith("CoreFn empty Guarded");
                                            };
                                            if (v1 instanceof Data_Maybe.Just) {
                                                return mainCb(Data_Array.snoc(caseRows)(buildCaseRow(new GuardedFn(v1.value0))));
                                            };
                                            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 517, column 23 - line 520, column 86): " + [ v1.constructor.name ]);
                                        })(v.value1.value0)([  ]);
                                    };
                                    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 506, column 13 - line 523, column 19): " + [ v.value1.constructor.name ]);
                                });
                            };
                        };
                    })(useCaseRowsCb)(alts)([  ]);
                };
            };
        };
        return function (v) {
            if (v instanceof PureScript_Backend_Optimizer_CoreFn.ExprVar) {
                return Control_Bind.bind(Control_Bind.bindFn)(ask1)(function (v1) {
                    var v2 = function (v3) {
                        var v4 = function (v5) {
                            if (v.value1.value0 instanceof Data_Maybe.Just && (v.value1.value0.value0 === "Prim" && v.value1.value1 === "undefined")) {
                                return buildM(PureScript_Backend_Optimizer_Syntax.PrimUndefined.value);
                            };
                            if (v.value1.value0 instanceof Data_Maybe.Nothing) {
                                return buildM(new PureScript_Backend_Optimizer_Syntax.Var(new PureScript_Backend_Optimizer_CoreFn.Qualified(new Data_Maybe.Just(v1.currentModule), v.value1.value1)));
                            };
                            return buildM(new PureScript_Backend_Optimizer_Syntax.Var(v.value1));
                        };
                        if (v.value1.value0 instanceof Data_Maybe.Just) {
                            var $561 = Data_Eq.eq(PureScript_Backend_Optimizer_CoreFn.eqModuleName)(v.value1.value0.value0)(v1.currentModule);
                            if ($561) {
                                var $562 = Data_Map_Internal.lookup(PureScript_Backend_Optimizer_CoreFn.ordIdent)(v.value1.value1)(v1.toLevel);
                                if ($562 instanceof Data_Maybe.Just) {
                                    return buildM(new PureScript_Backend_Optimizer_Syntax.Local(new Data_Maybe.Just(v.value1.value1), $562.value0));
                                };
                                return v4(true);
                            };
                            return v4(true);
                        };
                        return v4(true);
                    };
                    if (v.value1.value0 instanceof Data_Maybe.Nothing) {
                        var $568 = Data_Map_Internal.lookup(PureScript_Backend_Optimizer_CoreFn.ordIdent)(v.value1.value1)(v1.toLevel);
                        if ($568 instanceof Data_Maybe.Just) {
                            return buildM(new PureScript_Backend_Optimizer_Syntax.Local(new Data_Maybe.Just(v.value1.value1), $568.value0));
                        };
                        return v2(true);
                    };
                    return v2(true);
                });
            };
            if (v instanceof PureScript_Backend_Optimizer_CoreFn.ExprLit) {
                return Control_Bind.bindFlipped(Control_Bind.bindFn)(function ($763) {
                    return buildM(PureScript_Backend_Optimizer_Syntax.Lit.create($763));
                })(Data_Traversable.traverse(PureScript_Backend_Optimizer_CoreFn.traversableLiteral)(Control_Applicative.applicativeFn)(toBackendExpr)(v.value1));
            };
            if (v instanceof PureScript_Backend_Optimizer_CoreFn.ExprConstructor) {
                return Control_Bind.bind(Control_Bind.bindFn)(ask2)(function (v1) {
                    var ct = (function () {
                        var v2 = Data_Map_Internal.lookup(PureScript_Backend_Optimizer_CoreFn.ordProperName)(v.value1)(v1.dataTypes);
                        if (v2 instanceof Data_Maybe.Just && Data_Map_Internal.size(v2.value0.constructors) === 1) {
                            return PureScript_Backend_Optimizer_CoreFn.ProductType.value;
                        };
                        return PureScript_Backend_Optimizer_CoreFn.SumType.value;
                    })();
                    return buildM(new PureScript_Backend_Optimizer_Syntax.CtorDef(ct, v.value1, v.value2, v.value3));
                });
            };
            if (v instanceof PureScript_Backend_Optimizer_CoreFn.ExprAccessor) {
                return Control_Bind.bindFlipped(Control_Bind.bindFn)((function () {
                    var $764 = Data_Function.flip(PureScript_Backend_Optimizer_Syntax.Accessor.create)(new PureScript_Backend_Optimizer_Syntax.GetProp(v.value2));
                    return function ($765) {
                        return buildM($764($765));
                    };
                })())(toBackendExpr(v.value1));
            };
            if (v instanceof PureScript_Backend_Optimizer_CoreFn.ExprUpdate) {
                return join(Control_Apply.apply(Control_Apply.applyFn)(Data_Functor.map(Data_Functor.functorFn)(function (x) {
                    return function (y) {
                        return buildM(new PureScript_Backend_Optimizer_Syntax.Update(x, y));
                    };
                })(toBackendExpr(v.value1)))(Data_Traversable.traverse(Data_Traversable.traversableArray)(Control_Applicative.applicativeFn)(Data_Traversable.traverse(PureScript_Backend_Optimizer_CoreFn.traversableProp)(Control_Applicative.applicativeFn)(toBackendExpr))(v.value2)));
            };
            if (v instanceof PureScript_Backend_Optimizer_CoreFn.ExprAbs) {
                return Control_Bind.bind(Control_Bind.bindFn)(currentLevel)(function (lvl) {
                    return make(new PureScript_Backend_Optimizer_Syntax.Abs(Data_Array_NonEmpty.singleton(new Data_Tuple.Tuple(new Data_Maybe.Just(v.value1), lvl)), intro(Data_Foldable.foldableArray)([ v.value1 ])(lvl)(toBackendExpr(v.value2))));
                });
            };
            var v1 = function (v2) {
                if (v instanceof PureScript_Backend_Optimizer_CoreFn.ExprApp && Data_Boolean.otherwise) {
                    return make(new PureScript_Backend_Optimizer_Syntax.App(toBackendExpr(v.value1), Data_Array_NonEmpty.singleton(toBackendExpr(v.value2))));
                };
                if (v instanceof PureScript_Backend_Optimizer_CoreFn.ExprLet) {
                    var go1 = function (bind$prime) {
                        return function (next) {
                            if (bind$prime instanceof PureScript_Backend_Optimizer_CoreFn.NonRec) {
                                return makeLet(new Data_Maybe.Just(bind$prime.value0.value1))(toBackendExpr(bind$prime.value0.value2))(function (v3) {
                                    return next;
                                });
                            };
                            var v3 = function (v4) {
                                if (bind$prime instanceof PureScript_Backend_Optimizer_CoreFn.Rec) {
                                    return Partial_Unsafe.unsafeCrashWith("CoreFn empty Rec binding group");
                                };
                                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 423, column 1 - line 423, column 50): " + [ bind$prime.constructor.name ]);
                            };
                            if (bind$prime instanceof PureScript_Backend_Optimizer_CoreFn.Rec) {
                                var $608 = Data_Array_NonEmpty.fromArray(bind$prime.value0);
                                if ($608 instanceof Data_Maybe.Just) {
                                    return Control_Bind.bind(Control_Bind.bindFn)(currentLevel)(function (lvl) {
                                        var idents = Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)(function (v4) {
                                            return v4.value1;
                                        })($608.value0);
                                        return join(Control_Apply.apply(Control_Apply.applyFn)(Data_Functor.map(Data_Functor.functorFn)(function (x) {
                                            return function (y) {
                                                return buildM(new PureScript_Backend_Optimizer_Syntax.LetRec(lvl, x, y));
                                            };
                                        })(intro(Data_Array_NonEmpty_Internal.foldableNonEmptyArray)(idents)(lvl)(Data_Traversable.traverse(Data_Array_NonEmpty_Internal.traversableNonEmptyArray)(Control_Applicative.applicativeFn)(toBackendBinding)($608.value0))))(intro(Data_Array_NonEmpty_Internal.foldableNonEmptyArray)(idents)(lvl)(next)));
                                    });
                                };
                                return v3(true);
                            };
                            return v3(true);
                        };
                    };
                    return Data_Foldable.foldr(Data_Foldable.foldableArray)(go1)(toBackendExpr(v.value2))(v.value1);
                };
                if (v instanceof PureScript_Backend_Optimizer_CoreFn.ExprCase) {
                    return Data_Foldable.foldr(Data_Foldable.foldableArray)(function (expr1) {
                        return function (next) {
                            return function (idents) {
                                return makeLet(Data_Maybe.Nothing.value)(toBackendExpr(expr1))(function (tmp) {
                                    return next(Data_Array.snoc(idents)(tmp));
                                });
                            };
                        };
                    })(function (idents) {
                        return toInitialCaseRows(idents)(v.value2)(function (caseRows) {
                            return buildCaseTreeFromRows(caseRows);
                        });
                    })(v.value1)([  ]);
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 423, column 1 - line 423, column 50): " + [ v.constructor.name ]);
            };
            if (v instanceof PureScript_Backend_Optimizer_CoreFn.ExprApp) {
                if (v.value1 instanceof PureScript_Backend_Optimizer_CoreFn.ExprVar && (v.value1.value0.meta instanceof Data_Maybe.Just && v.value1.value0.meta.value0 instanceof PureScript_Backend_Optimizer_CoreFn.IsNewtype)) {
                    return toBackendExpr(v.value2);
                };
                return v1(true);
            };
            return v1(true);
        };
    })();
    var v = PureScript_Backend_Optimizer_CoreFn.exprAnn(expr);
    return Control_Bind.bind(Control_Bind.bindFn)(go(expr))(function (backendExpr) {
        return Control_Applicative.pure(Control_Applicative.applicativeFn)((function () {
            var v1 = (function () {
                if (v.type instanceof Data_Maybe.Just) {
                    return new Data_Maybe.Just(v.type.value0);
                };
                if (v.type instanceof Data_Maybe.Nothing) {
                    return inferExprType(expr);
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 427, column 13 - line 429, column 45): " + [ v.type.constructor.name ]);
            })();
            if (v1 instanceof Data_Maybe.Just) {
                return new PureScript_Backend_Optimizer_Semantics.ExprSyntax(PureScript_Backend_Optimizer_Analysis.analysisOf(PureScript_Backend_Optimizer_Semantics.hasAnalysisBackendExpr)(backendExpr), new PureScript_Backend_Optimizer_Syntax.Typed(v1.value0, backendExpr));
            };
            if (v1 instanceof Data_Maybe.Nothing) {
                return backendExpr;
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 427, column 8 - line 431, column 27): " + [ v1.constructor.name ]);
        })());
    });
};
var toBackendBinding = function (v) {
    return Data_Functor.map(Data_Functor.functorFn)(Data_Tuple.Tuple.create(v.value1))(toBackendExpr(v.value2));
};
var buildCaseTreeFromRows = function (denormalizedRows) {
    var v = Data_Array_NonEmpty.fromArray(normalizeCaseRows(denormalizedRows));
    if (v instanceof Data_Maybe.Nothing) {
        return patternFail;
    };
    if (v instanceof Data_Maybe.Just) {
        var v1 = Data_Array_NonEmpty.uncons(v.value0);
        var row0NonPatWildPatterns = Data_Array_NonEmpty.fromArray(Data_FoldableWithIndex.foldlWithIndex(Data_FoldableWithIndex.foldableWithIndexArray)(function (idx) {
            return function (acc) {
                return function (p) {
                    var $641 = Data_Eq.notEq(eqPatternCase)(patternPatCase(p))(PatWild.value);
                    if ($641) {
                        return Data_Array.snoc(acc)(new Data_Tuple.Tuple(idx, p));
                    };
                    return acc;
                };
            };
        })([  ])(v1.head.patterns));
        if (row0NonPatWildPatterns instanceof Data_Maybe.Nothing) {
            return buildCaseLeaf(v1.head)(v1.tail);
        };
        if (row0NonPatWildPatterns instanceof Data_Maybe.Just) {
            return buildCasePattern(chooseNextPattern(row0NonPatWildPatterns.value0)(v1.tail))(Data_Array_NonEmpty.toArray(v.value0));
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 702, column 5 - line 706, column 95): " + [ row0NonPatWildPatterns.constructor.name ]);
    };
    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 691, column 42 - line 706, column 95): " + [ v.constructor.name ]);
};
var buildCasePattern = function (chosenColumn) {
    return function (rows) {
        
        // | Rebuilds the case row by doing two things:
        // | 1. replacing the matched pattern with its subterm patterns
        // | 2. adding the `vars` exposed by this case row's corresponding binder to the case row's `vars`
var rebuildCaseRow = function (idents) {
            var inlineWildSubterms = Data_Functor.mapFlipped(Data_Functor.functorArray)(idents)(function (column) {
                return {
                    column: column,
                    pattern: {
                        vars: Data_Set.empty,
                        patternCase: PatWild.value,
                        subterms: [  ]
                    }
                };
            });
            var convertSubtermToPattern = function (column) {
                return function (v) {
                    return {
                        column: column,
                        pattern: v.pattern
                    };
                };
            };
            return Data_Functor.map(Data_Functor.functorArray)(function (v) {
                var subtermPatterns = (function () {
                    var v1 = patternPatCase(v.match);
                    if (v1 instanceof PatWild) {
                        return inlineWildSubterms;
                    };
                    return Data_Array.zipWith(convertSubtermToPattern)(idents)(patternSubterms(v.match));
                })();
                return {
                    guardFn: v.guardFn,
                    vars: Data_Semigroup.append(semigroupSemigroupMap)(v.vars)(toCaseRowVars(v.match)),
                    patterns: Data_Semigroup.append(Data_Semigroup.semigroupArray)(v.nonMatchesBefore)(Data_Semigroup.append(Data_Semigroup.semigroupArray)(subtermPatterns)(v.nonMatchesAfter))
                };
            });
        };
        var letBindSubterm = function (v) {
            return function (nextCb) {
                return function (idents) {
                    var parentExpr = make(new PureScript_Backend_Optimizer_Syntax.Local(Data_Maybe.Nothing.value, chosenColumn.column));
                    return makeLet(Data_Maybe.Nothing.value)(make(new PureScript_Backend_Optimizer_Syntax.Accessor(parentExpr, v.accessor)))(function (tmp) {
                        return nextCb(Data_Array.snoc(idents)(tmp));
                    });
                };
            };
        };
        
        // There's no guard to make here. We just expose all subterms as patterns in the following expressions.
var expandSubterms = (function () {
            var v = decompose(chosenColumn)(rows);
            return Data_Foldable.foldr(Data_Foldable.foldableArray)(letBindSubterm)(function (idents) {
                return buildCaseTreeFromRows(rebuildCaseRow(idents)(v.rowsWithMatch));
            })(patternSubterms(chosenColumn))([  ]);
        })();
        var buildCaseBranch = function (guardExpr) {
            var v = decompose(chosenColumn)(rows);
            var exprOnPatternMiss = buildCaseTreeFromRows(v.rowsNoMatch);
            var exprOnPatternMatch = Data_Foldable.foldr(Data_Foldable.foldableArray)(letBindSubterm)(function (idents) {
                return buildCaseTreeFromRows(rebuildCaseRow(idents)(v.rowsWithMatch));
            })(patternSubterms(chosenColumn))([  ]);
            return makeGuard(chosenColumn.column)(guardExpr)(exprOnPatternMatch)(exprOnPatternMiss);
        };
        var v = patternPatCase(chosenColumn);
        if (v instanceof PatWild) {
            return Partial_Unsafe.unsafeCrashWith("Impossible: chosen column cannot be wild pattern");
        };
        if (v instanceof PatRecord) {
            return expandSubterms;
        };
        if (v instanceof PatProduct) {
            return expandSubterms;
        };
        if (v instanceof PatSum) {
            return buildCaseBranch(guardTag(v.value1));
        };
        if (v instanceof PatArray) {
            return buildCaseBranch(guardArrayLength(v.value0));
        };
        if (v instanceof PatInt) {
            return buildCaseBranch(guardInt(v.value0));
        };
        if (v instanceof PatNumber) {
            return buildCaseBranch(guardNumber(v.value0));
        };
        if (v instanceof PatString) {
            return buildCaseBranch(guardString(v.value0));
        };
        if (v instanceof PatChar) {
            return buildCaseBranch(guardChar(v.value0));
        };
        if (v instanceof PatBoolean) {
            return buildCaseBranch(guardBoolean(v.value0));
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 832, column 38 - line 852, column 37): " + [ v.constructor.name ]);
    };
};
var buildCaseLeaf = function (row0) {
    return function (tailRows) {
        var orderedArgs = toUnfoldable(coerce(Data_Semigroup.append(semigroupSemigroupMap)(row0.vars)(Data_Foldable.foldMap(Data_Foldable.foldableArray)(monoidSemigroupMap)(toCaseRowVars)(row0.patterns))));
        var callFn = function (fn) {
            return function (args) {
                return make(new PureScript_Backend_Optimizer_Syntax.UncurriedApp(make(new PureScript_Backend_Optimizer_Syntax.Local(Data_Maybe.Nothing.value, fn)), Data_Functor.map(Data_Functor.functorArray)(function (v) {
                    return make(new PureScript_Backend_Optimizer_Syntax.Local(new Data_Maybe.Just(v.value0), v.value1));
                })(args)));
            };
        };
        if (row0.guardFn instanceof UnconditionalFn) {
            return callFn(row0.guardFn.value0)(orderedArgs);
        };
        if (row0.guardFn instanceof GuardedFn) {
            return Data_Foldable.foldr(Data_Foldable.foldableArray)(function (v) {
                return function (cb) {
                    return function (args) {
                        return makeLet(new Data_Maybe.Just(v.value0))(make(new PureScript_Backend_Optimizer_Syntax.Local(Data_Maybe.Nothing.value, v.value1)))(function (tmp) {
                            return cb(Data_Array.snoc(args)(new Data_Tuple.Tuple(v.value0, tmp)));
                        });
                    };
                };
            })(function (args) {
                return Control_Bind.bind(Control_Bind.bindFn)(Data_Traversable["for"](Control_Applicative.applicativeFn)(Data_Array_NonEmpty_Internal.traversableNonEmptyArray)(row0.guardFn.value0)(function (v) {
                    return Control_Apply.apply(Control_Apply.applyFn)(Data_Functor.map(Data_Functor.functorFn)(PureScript_Backend_Optimizer_Syntax.Pair.create)(toBackendExpr(v.value0)))(callFn(v.value1)(args));
                }))(function (pairs) {
                    return Control_Bind.bind(Control_Bind.bindFn)(buildCaseTreeFromRows(tailRows))(function (fallback) {
                        return buildM(new PureScript_Backend_Optimizer_Syntax.Branch(pairs, fallback));
                    });
                });
            })(orderedArgs)([  ]);
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 757, column 3 - line 779, column 11): " + [ row0.guardFn.constructor.name ]);
    };
};
var toTopLevelBackendBinding = function (group) {
    return function (env) {
        return function (v) {
            var evalEnv = {
                currentModule: env.currentModule,
                evalExternRef: makeExternEvalRef(env),
                evalExternSpine: makeExternEvalSpine(env),
                locals: [  ],
                directives: env.directives
            };
            var qualifiedIdent = new PureScript_Backend_Optimizer_CoreFn.Qualified(new Data_Maybe.Just(env.currentModule), v.value1);
            var backendExpr = toBackendExpr(v.value2)(env);
            var enableTracing = Data_Set.member(ordQualified1)(qualifiedIdent)(env.traceIdents);
            var mbType = (function () {
                if (backendExpr instanceof PureScript_Backend_Optimizer_Semantics.ExprSyntax && backendExpr.value1 instanceof PureScript_Backend_Optimizer_Syntax.Typed) {
                    return new Data_Maybe.Just(backendExpr.value1.value0);
                };
                return Data_Maybe.Nothing.value;
            })();
            var v1 = PureScript_Backend_Optimizer_Semantics.optimize(enableTracing)(getCtx(env))(evalEnv)(qualifiedIdent)(env.rewriteLimit)(backendExpr);
            var optimizedExprWithTy = (function () {
                if (mbType instanceof Data_Maybe.Just) {
                    return new PureScript_Backend_Optimizer_Semantics.ExprSyntax(PureScript_Backend_Optimizer_Analysis.analysisOf(PureScript_Backend_Optimizer_Semantics.hasAnalysisBackendExpr)(v1.value1), new PureScript_Backend_Optimizer_Syntax.Typed(mbType.value0, v1.value1));
                };
                if (mbType instanceof Data_Maybe.Nothing) {
                    return v1.value1;
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 261, column 27 - line 263, column 31): " + [ mbType.constructor.name ]);
            })();
            var v2 = toExternImpl(env)(group)(optimizedExprWithTy);
            return {
                accum: {
                    analyzeCustom: env.analyzeCustom,
                    currentLevel: env.currentLevel,
                    currentModule: env.currentModule,
                    dataTypes: env.dataTypes,
                    toLevel: env.toLevel,
                    foreignSemantics: env.foreignSemantics,
                    rewriteLimit: env.rewriteLimit,
                    traceIdents: env.traceIdents,
                    implementations: Data_Map_Internal.insert(ordQualified1)(qualifiedIdent)(v2.value0)(env.implementations),
                    moduleImplementations: Data_Map_Internal.insert(ordQualified1)(qualifiedIdent)(v2.value0)(env.moduleImplementations),
                    optimizationSteps: Data_Maybe.maybe(env.optimizationSteps)((function () {
                        var $766 = Data_Array.snoc(env.optimizationSteps);
                        var $767 = Data_Tuple.Tuple.create(qualifiedIdent);
                        return function ($768) {
                            return $766($767($768));
                        };
                    })())(Data_Array_NonEmpty.fromArray(v1.value0)),
                    directives: (function () {
                        var v4 = inferTransitiveDirective(env.directives)(Data_Tuple.snd(v2.value0))(backendExpr)(v.value2);
                        if (v4 instanceof Data_Maybe.Just) {
                            return Data_Map_Internal.alter(PureScript_Backend_Optimizer_Semantics.ordEvalRef)(function (v5) {
                                if (v5 instanceof Data_Maybe.Just) {
                                    return new Data_Maybe.Just(Data_Map_Internal.union(PureScript_Backend_Optimizer_Semantics.ordInlineAccessor)(v5.value0)(v4.value0));
                                };
                                if (v5 instanceof Data_Maybe.Nothing) {
                                    return new Data_Maybe.Just(v4.value0);
                                };
                                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 273, column 17 - line 277, column 30): " + [ v5.constructor.name ]);
                            })(new PureScript_Backend_Optimizer_Semantics.EvalExtern(new PureScript_Backend_Optimizer_CoreFn.Qualified(new Data_Maybe.Just(env.currentModule), v.value1)))(env.directives);
                        };
                        if (v4 instanceof Data_Maybe.Nothing) {
                            return env.directives;
                        };
                        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 270, column 11 - line 281, column 29): " + [ v4.constructor.name ]);
                    })()
                },
                value: new Data_Tuple.Tuple(v.value1, new Data_Tuple.Tuple((Data_Newtype.unwrap()(Data_Tuple.fst(v2.value0))).deps, v2.value1))
            };
        };
    };
};
var toBackendTopLevelBindingGroup = function (env) {
    var overValue = function (f) {
        return function (a) {
            return {
                accum: a.accum,
                value: f(a.value)
            };
        };
    };
    return function (v) {
        if (v instanceof PureScript_Backend_Optimizer_CoreFn.Rec) {
            var group = Data_Functor.map(Data_Functor.functorArray)(function (v1) {
                return new PureScript_Backend_Optimizer_CoreFn.Qualified(new Data_Maybe.Just(env.currentModule), v1.value1);
            })(v.value0);
            return overValue(function (v1) {
                return {
                    recursive: true,
                    bindings: v1
                };
            })(Data_Traversable.mapAccumL(Data_Traversable.traversableArray)(toTopLevelBackendBinding(group))(env)(v.value0));
        };
        if (v instanceof PureScript_Backend_Optimizer_CoreFn.NonRec) {
            return overValue(function (v1) {
                return {
                    recursive: false,
                    bindings: v1
                };
            })(Data_Traversable.mapAccumL(Data_Traversable.traversableArray)(toTopLevelBackendBinding([  ]))(env)([ v.value0 ]));
        };
        throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 231, column 37 - line 238, column 52): " + [ v.constructor.name ]);
    };
};
var toBackendTopLevelBindingGroups = function (binds) {
    return function (env) {
        var result = Data_Traversable.mapAccumL(Data_Traversable.traversableArray)(toBackendTopLevelBindingGroup)(env)(binds);
        return {
            accum: result.accum,
            value: Data_Functor.map(Data_Functor.functorArray)(function (as) {
                return {
                    recursive: (Data_Array_NonEmpty.head(as)).recursive,
                    bindings: Control_Bind.bindFlipped(Control_Bind.bindArray)(function (v1) {
                        return v1.bindings;
                    })(Data_Array_NonEmpty.toArray(as))
                };
            })(Data_Array.groupBy(Data_Function.on(conj)(function ($769) {
                return !(function (v1) {
                    return v1.recursive;
                })($769);
            }))(result.value))
        };
    };
};
var toBackendModule = function (v) {
    return function (env) {
        var localExports = Data_Set.fromFoldable(Data_Foldable.foldableArray)(PureScript_Backend_Optimizer_CoreFn.ordIdent)(v.exports);
        var isBindingUsed = function (deps) {
            return function (v1) {
                return Data_Set.member(PureScript_Backend_Optimizer_CoreFn.ordIdent)(v1.value0)(localExports) || Data_Set.member(ordQualified1)(new PureScript_Backend_Optimizer_CoreFn.Qualified(new Data_Maybe.Just(v.name), v1.value0))(deps);
            };
        };
        var directives = PureScript_Backend_Optimizer_Directives.parseDirectiveHeader(v.name)(v.comments);
        var ctors = Control_Bind.bind(Control_Bind.bindArray)(Control_Bind.bind(Control_Bind.bindArray)(v.decls)(function (v1) {
            if (v1 instanceof PureScript_Backend_Optimizer_CoreFn.Rec) {
                return v1.value0;
            };
            if (v1 instanceof PureScript_Backend_Optimizer_CoreFn.NonRec) {
                return Control_Applicative.pure(Control_Applicative.applicativeArray)(v1.value0);
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.Convert (line 129, column 42 - line 131, column 39): " + [ v1.constructor.name ]);
        }))(function (v1) {
            if (v1.value2 instanceof PureScript_Backend_Optimizer_CoreFn.ExprConstructor) {
                return pure1(new Data_Tuple.Tuple(v1.value2.value1, new Data_Tuple.Tuple(v1.value2.value2, v1.value2.value3)));
            };
            return [  ];
        });
        var dataTypes = fromFoldable(Data_Functor.map(Data_Functor.functorArray)(function (group) {
            var proper = Data_Tuple.fst(Data_Array_NonEmpty.head(group));
            var constructors = fromFoldable1(Data_FunctorWithIndex.mapWithIndex(Data_Array_NonEmpty_Internal.functorWithIndexNonEmptyArray)(function (tag) {
                return function (v1) {
                    return new Data_Tuple.Tuple(v1.value1.value0, {
                        fields: v1.value1.value1,
                        tag: tag
                    });
                };
            })(group));
            var sizes = Data_Functor.map(Data_Array_NonEmpty_Internal.functorNonEmptyArray)(function ($770) {
                return Data_Array.length(Data_Tuple.snd(Data_Tuple.snd($770)));
            })(group);
            return new Data_Tuple.Tuple(proper, {
                constructors: constructors,
                size: Data_Semigroup_Foldable.maximum(Data_Ord.ordInt)(Data_Array_NonEmpty_Internal.foldable1NonEmptyArray)(sizes)
            });
        })(Data_Array.groupAllBy(Data_Ord.comparing(PureScript_Backend_Optimizer_CoreFn.ordProperName)(Data_Tuple.fst))(ctors)));
        var moduleBindings = toBackendTopLevelBindingGroups(v.decls)({
            analyzeCustom: env.analyzeCustom,
            currentLevel: env.currentLevel,
            currentModule: env.currentModule,
            toLevel: env.toLevel,
            implementations: env.implementations,
            optimizationSteps: env.optimizationSteps,
            foreignSemantics: env.foreignSemantics,
            rewriteLimit: env.rewriteLimit,
            traceIdents: env.traceIdents,
            dataTypes: dataTypes,
            directives: Data_FoldableWithIndex.foldlWithIndex(Data_Map_Internal.foldableWithIndexMap)(function (qual) {
                return function (dirs) {
                    return function (dir) {
                        return Data_Map_Internal.alter(PureScript_Backend_Optimizer_Semantics.ordEvalRef)(Data_Maybe.maybe(new Data_Maybe.Just(dir))(Data_Maybe.Just.create))(qual)(dirs);
                    };
                };
            })(Data_Map_Internal.union(PureScript_Backend_Optimizer_Semantics.ordEvalRef)(directives.locals)(env.directives))(directives.exports),
            moduleImplementations: Data_Map_Internal.empty
        });
        var usedBindings = Data_Traversable.mapAccumR(Data_Traversable.traversableArray)(function (deps) {
            return function (group) {
                var v1 = (function () {
                    if (group.recursive) {
                        var $743 = Data_Array.any(isBindingUsed(deps))(group.bindings);
                        if ($743) {
                            return {
                                accum: Data_Semigroup.append(semigroupSet1)(Data_Foldable.foldMap(Data_Foldable.foldableArray)(monoidSet)(function ($771) {
                                    return Data_Tuple.fst(Data_Tuple.snd($771));
                                })(group.bindings))(deps),
                                value: Data_Functor.map(Data_Functor.functorArray)((function () {
                                    var $772 = Data_Functor.map(Data_Tuple.functorTuple)(Data_Tuple.snd);
                                    return function ($773) {
                                        return Data_Maybe.Just.create($772($773));
                                    };
                                })())(group.bindings)
                            };
                        };
                        return {
                            accum: deps,
                            value: [  ]
                        };
                    };
                    return Data_Traversable.mapAccumR(Data_Traversable.traversableArray)(function (deps$prime) {
                        return function (v2) {
                            var $745 = isBindingUsed(deps$prime)(v2);
                            if ($745) {
                                return {
                                    accum: Data_Semigroup.append(semigroupSet1)(v2.value1.value0)(deps$prime),
                                    value: new Data_Maybe.Just(new Data_Tuple.Tuple(v2.value0, v2.value1.value1))
                                };
                            };
                            return {
                                accum: deps$prime,
                                value: Data_Maybe.Nothing.value
                            };
                        };
                    })(deps)(group.bindings);
                })();
                return {
                    accum: v1.accum,
                    value: {
                        recursive: group.recursive,
                        bindings: Data_Array.catMaybes(v1.value)
                    }
                };
            };
        })(Data_Set.empty)(moduleBindings.value);
        var usedImports = Data_Set.mapMaybe(PureScript_Backend_Optimizer_CoreFn.ordModuleName)(function (qi) {
            return Control_Bind.bind(Data_Maybe.bindMaybe)(PureScript_Backend_Optimizer_CoreFn.qualifiedModuleName(qi))(function (mn) {
                return Data_Functor.voidRight(Data_Maybe.functorMaybe)(mn)(Control_Alternative.guard(Data_Maybe.alternativeMaybe)(Data_Eq.notEq(PureScript_Backend_Optimizer_CoreFn.eqModuleName)(mn)(v.name) && Data_Eq.notEq(PureScript_Backend_Optimizer_CoreFn.eqModuleName)(mn)("Prim")));
            });
        })(usedBindings.accum);
        return new Data_Tuple.Tuple(moduleBindings.accum.optimizationSteps, {
            name: v.name,
            comments: v.comments,
            dataDecls: v.dataDecls,
            classDecls: v.classDecls,
            imports: usedImports,
            dataTypes: Data_Map_Internal.filter(PureScript_Backend_Optimizer_CoreFn.ordProperName)((function () {
                var $774 = Data_Array.any(isBindingUsed(usedBindings.accum));
                return function ($775) {
                    return $774(toUnfoldable1((function (v1) {
                        return v1.constructors;
                    })($775)));
                };
            })())(dataTypes),
            bindings: usedBindings.value,
            exports: localExports,
            reExports: Data_Set.fromFoldable(Data_Foldable.foldableArray)(PureScript_Backend_Optimizer_CoreFn.ordReExport)(v.reExports),
            implementations: moduleBindings.accum.moduleImplementations,
            directives: directives.exports,
            foreign: v.foreign
        });
    };
};
export {
    toBackendModule,
    toBackendTopLevelBindingGroups,
    toBackendTopLevelBindingGroup,
    toTopLevelBackendBinding,
    inferTransitiveDirective,
    toExternImpl,
    topEnv,
    makeExternEvalSpine,
    makeExternEvalRef,
    buildM,
    getCtx,
    fromExternImpl,
    levelUp,
    intro,
    currentLevel,
    toBackendExpr,
    UnconditionalFn,
    GuardedFn,
    Pattern,
    PatWild,
    PatRecord,
    PatProduct,
    PatArray,
    PatSum,
    PatInt,
    PatNumber,
    PatString,
    PatChar,
    PatBoolean,
    binderToPattern,
    patternVars,
    toCaseRowVars,
    patternPatCase,
    patternSubterms,
    buildCaseTreeFromRows,
    normalizeCaseRows,
    buildCaseLeaf,
    chooseNextPattern,
    buildCasePattern,
    decompose,
    patternFail,
    makeLet,
    guardInt,
    guardNumber,
    guardString,
    guardChar,
    guardBoolean,
    guardArrayLength,
    guardTag,
    makeGuard,
    makeUncurriedAbs,
    make,
    toBackendBinding,
    inferExprType,
    getReturnType,
    newtypePattern_,
    eqPatternCase,
    ordPatternCase
};
