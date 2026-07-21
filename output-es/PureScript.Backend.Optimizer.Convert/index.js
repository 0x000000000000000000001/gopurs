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
import * as $runtime from "../runtime.js";
import * as Control$dApplicative from "../Control.Applicative/index.js";
import * as Control$dBind from "../Control.Bind/index.js";
import * as Data$dArray from "../Data.Array/index.js";
import * as Data$dArray$dNonEmpty from "../Data.Array.NonEmpty/index.js";
import * as Data$dArray$dNonEmpty$dInternal from "../Data.Array.NonEmpty.Internal/index.js";
import * as Data$dEq from "../Data.Eq/index.js";
import * as Data$dFoldable from "../Data.Foldable/index.js";
import * as Data$dFoldableWithIndex from "../Data.FoldableWithIndex/index.js";
import * as Data$dFunction from "../Data.Function/index.js";
import * as Data$dFunctor from "../Data.Functor/index.js";
import * as Data$dFunctorWithIndex from "../Data.FunctorWithIndex/index.js";
import * as Data$dList$dTypes from "../Data.List.Types/index.js";
import * as Data$dMap from "../Data.Map/index.js";
import * as Data$dMap$dInternal from "../Data.Map.Internal/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dMonoid from "../Data.Monoid/index.js";
import * as Data$dOrd from "../Data.Ord/index.js";
import * as Data$dOrdering from "../Data.Ordering/index.js";
import * as Data$dSemigroup$dFirst from "../Data.Semigroup.First/index.js";
import * as Data$dSemigroup$dFoldable from "../Data.Semigroup.Foldable/index.js";
import * as Data$dSet from "../Data.Set/index.js";
import * as Data$dTraversable from "../Data.Traversable/index.js";
import * as Data$dTraversableWithIndex from "../Data.TraversableWithIndex/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
import * as Data$dUnfoldable from "../Data.Unfoldable/index.js";
import * as Partial from "../Partial/index.js";
import * as PureScript$dBackend$dOptimizer$dAnalysis from "../PureScript.Backend.Optimizer.Analysis/index.js";
import * as PureScript$dBackend$dOptimizer$dCoreFn from "../PureScript.Backend.Optimizer.CoreFn/index.js";
import * as PureScript$dBackend$dOptimizer$dDirectives from "../PureScript.Backend.Optimizer.Directives/index.js";
import * as PureScript$dBackend$dOptimizer$dSemantics from "../PureScript.Backend.Optimizer.Semantics/index.js";
import * as PureScript$dBackend$dOptimizer$dSyntax from "../PureScript.Backend.Optimizer.Syntax/index.js";
const $CaseRowGuardedExpr = (tag, _1) => ({tag, _1});
const $PatternCase = (tag, _1, _2) => ({tag, _1, _2});
const eq = /* #__PURE__ */ Data$dEq.eqArrayImpl(Data$dEq.eqStringImpl);
const compare = /* #__PURE__ */ (() => Data$dOrd.ordArray(Data$dOrd.ordString).compare)();
const compare1 = /* #__PURE__ */ (() => PureScript$dBackend$dOptimizer$dCoreFn.ordQualified(Data$dOrd.ordString).compare)();
const ordQualified = /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.ordQualified(Data$dOrd.ordString);
const lookup = k => {
  const go = go$a0$copy => {
    let go$a0 = go$a0$copy, go$c = true, go$r;
    while (go$c) {
      const v = go$a0;
      if (v.tag === "Leaf") {
        go$c = false;
        go$r = Data$dMaybe.Nothing;
        continue;
      }
      if (v.tag === "Node") {
        const v1 = Data$dOrd.ordString.compare(k)(v._3);
        if (v1 === "LT") {
          go$a0 = v._5;
          continue;
        }
        if (v1 === "GT") {
          go$a0 = v._6;
          continue;
        }
        if (v1 === "EQ") {
          go$c = false;
          go$r = Data$dMaybe.$Maybe("Just", v._4);
          continue;
        }
      }
      $runtime.fail();
    }
    return go$r;
  };
  return go;
};
const monoidSemigroupMap = /* #__PURE__ */ Data$dMap.monoidSemigroupMap(Data$dOrd.ordString)(Data$dSemigroup$dFirst.semigroupFirst);
const foldMap = /* #__PURE__ */ (() => Data$dSet.foldableSet.foldMap(monoidSemigroupMap))();
const toUnfoldable = /* #__PURE__ */ (() => {
  const $0 = Data$dUnfoldable.unfoldableArray.unfoldr(xs => {
    if (xs.tag === "Nil") { return Data$dMaybe.Nothing; }
    if (xs.tag === "Cons") { return Data$dMaybe.$Maybe("Just", Data$dTuple.$Tuple(xs._1, xs._2)); }
    $runtime.fail();
  });
  return x => $0((() => {
    const go = (m$p, z$p) => {
      if (m$p.tag === "Leaf") { return z$p; }
      if (m$p.tag === "Node") { return go(m$p._5, Data$dList$dTypes.$List("Cons", m$p._3, go(m$p._6, z$p))); }
      $runtime.fail();
    };
    return go(x, Data$dList$dTypes.Nil);
  })());
})();
const foldMap2 = /* #__PURE__ */ (() => Data$dFoldable.foldableArray.foldMap(Data$dMonoid.monoidArray))();
const fromFoldable1 = /* #__PURE__ */ Data$dFoldable.foldlArray(m => a => Data$dMap$dInternal.insert(Data$dOrd.ordString)(a)()(m))(Data$dMap$dInternal.Leaf);
const lookup1 = k => {
  const go = go$a0$copy => {
    let go$a0 = go$a0$copy, go$c = true, go$r;
    while (go$c) {
      const v = go$a0;
      if (v.tag === "Leaf") {
        go$c = false;
        go$r = Data$dMaybe.Nothing;
        continue;
      }
      if (v.tag === "Node") {
        const v1 = ordQualified.compare(k)(v._3);
        if (v1 === "LT") {
          go$a0 = v._5;
          continue;
        }
        if (v1 === "GT") {
          go$a0 = v._6;
          continue;
        }
        if (v1 === "EQ") {
          go$c = false;
          go$r = Data$dMaybe.$Maybe("Just", v._4);
          continue;
        }
      }
      $runtime.fail();
    }
    return go$r;
  };
  return go;
};
const lookup2 = k => {
  const go = go$a0$copy => {
    let go$a0 = go$a0$copy, go$c = true, go$r;
    while (go$c) {
      const v = go$a0;
      if (v.tag === "Leaf") {
        go$c = false;
        go$r = Data$dMaybe.Nothing;
        continue;
      }
      if (v.tag === "Node") {
        const v1 = PureScript$dBackend$dOptimizer$dSemantics.ordEvalRef.compare(k)(v._3);
        if (v1 === "LT") {
          go$a0 = v._5;
          continue;
        }
        if (v1 === "GT") {
          go$a0 = v._6;
          continue;
        }
        if (v1 === "EQ") {
          go$c = false;
          go$r = Data$dMaybe.$Maybe("Just", v._4);
          continue;
        }
      }
      $runtime.fail();
    }
    return go$r;
  };
  return go;
};
const lookup3 = k => {
  const go = go$a0$copy => {
    let go$a0 = go$a0$copy, go$c = true, go$r;
    while (go$c) {
      const v = go$a0;
      if (v.tag === "Leaf") {
        go$c = false;
        go$r = Data$dMaybe.Nothing;
        continue;
      }
      if (v.tag === "Node") {
        const v1 = PureScript$dBackend$dOptimizer$dSemantics.ordInlineAccessor.compare(k)(v._3);
        if (v1 === "LT") {
          go$a0 = v._5;
          continue;
        }
        if (v1 === "GT") {
          go$a0 = v._6;
          continue;
        }
        if (v1 === "EQ") {
          go$c = false;
          go$r = Data$dMaybe.$Maybe("Just", v._4);
          continue;
        }
      }
      $runtime.fail();
    }
    return go$r;
  };
  return go;
};
const analyzeEffectBlock = /* #__PURE__ */ PureScript$dBackend$dOptimizer$dAnalysis.analyzeEffectBlock(PureScript$dBackend$dOptimizer$dSemantics.hasAnalysisBackendExpr)(PureScript$dBackend$dOptimizer$dSemantics.hasSyntaxBackendExpr);
const analyze = /* #__PURE__ */ PureScript$dBackend$dOptimizer$dAnalysis.analyze(PureScript$dBackend$dOptimizer$dSemantics.hasAnalysisBackendExpr)(PureScript$dBackend$dOptimizer$dSemantics.hasSyntaxBackendExpr);
const foldMap3 = /* #__PURE__ */ (() => Data$dFoldable.foldableArray.foldMap((() => {
  const semigroupRecord1 = {append: ra => rb => ({rowsNoMatch: [...ra.rowsNoMatch, ...rb.rowsNoMatch], rowsWithMatch: [...ra.rowsWithMatch, ...rb.rowsWithMatch]})};
  return {mempty: {rowsNoMatch: [], rowsWithMatch: []}, Semigroup0: () => semigroupRecord1};
})()))();
const lookup4 = k => {
  const go = go$a0$copy => {
    let go$a0 = go$a0$copy, go$c = true, go$r;
    while (go$c) {
      const v = go$a0;
      if (v.tag === "Leaf") {
        go$c = false;
        go$r = Data$dMaybe.Nothing;
        continue;
      }
      if (v.tag === "Node") {
        const v1 = Data$dOrd.ordString.compare(k)(v._3);
        if (v1 === "LT") {
          go$a0 = v._5;
          continue;
        }
        if (v1 === "GT") {
          go$a0 = v._6;
          continue;
        }
        if (v1 === "EQ") {
          go$c = false;
          go$r = Data$dMaybe.$Maybe("Just", v._4);
          continue;
        }
      }
      $runtime.fail();
    }
    return go$r;
  };
  return go;
};
const forWithIndex = /* #__PURE__ */ (() => {
  const $0 = Data$dTraversableWithIndex.traversableWithIndexArray.traverseWithIndex(Control$dApplicative.applicativeFn);
  return b => a => $0(a)(b);
})();
const zipWithA = /* #__PURE__ */ Data$dArray.zipWithA(Control$dApplicative.applicativeFn);
const traverse = /* #__PURE__ */ (() => PureScript$dBackend$dOptimizer$dCoreFn.traversableLiteral.traverse(Control$dApplicative.applicativeFn))();
const traverse1 = /* #__PURE__ */ (() => Data$dTraversable.traversableArray.traverse(Control$dApplicative.applicativeFn))();
const traverse3 = /* #__PURE__ */ (() => Data$dTraversable.traversableArray.traverse(Control$dApplicative.applicativeFn))();
const append2 = /* #__PURE__ */ (() => Data$dMap.semigroupSemigroupMap(Data$dOrd.ordString)(Data$dSemigroup$dFirst.semigroupFirst).append)();
const toUnfoldable1 = /* #__PURE__ */ (() => {
  const $0 = Data$dUnfoldable.unfoldableArray.unfoldr(Data$dMap$dInternal.stepUnfoldr);
  return x => $0(Data$dMap$dInternal.$MapIter("IterNode", x, Data$dMap$dInternal.IterLeaf));
})();
const foldMap4 = /* #__PURE__ */ (() => Data$dFoldable.foldableArray.foldMap(monoidSemigroupMap))();
const $$for = /* #__PURE__ */ (() => {
  const traverse2 = Data$dTraversable.traversableArray.traverse(Control$dApplicative.applicativeFn);
  return x => f => traverse2(f)(x);
})();
const member = k => {
  const go = go$a0$copy => {
    let go$a0 = go$a0$copy, go$c = true, go$r;
    while (go$c) {
      const v = go$a0;
      if (v.tag === "Leaf") {
        go$c = false;
        go$r = false;
        continue;
      }
      if (v.tag === "Node") {
        const v1 = ordQualified.compare(k)(v._3);
        if (v1 === "LT") {
          go$a0 = v._5;
          continue;
        }
        if (v1 === "GT") {
          go$a0 = v._6;
          continue;
        }
        if (v1 === "EQ") {
          go$c = false;
          go$r = true;
          continue;
        }
      }
      $runtime.fail();
    }
    return go$r;
  };
  return go;
};
const alter = /* #__PURE__ */ Data$dMap$dInternal.alter(PureScript$dBackend$dOptimizer$dSemantics.ordEvalRef);
const mapAccumL = /* #__PURE__ */ Data$dTraversable.mapAccumL(Data$dTraversable.traversableArray);
const fromFoldable2 = /* #__PURE__ */ Data$dFoldable.foldlArray(m => a => Data$dMap$dInternal.insert(Data$dOrd.ordString)(a)()(m))(Data$dMap$dInternal.Leaf);
const member1 = k => {
  const go = go$a0$copy => {
    let go$a0 = go$a0$copy, go$c = true, go$r;
    while (go$c) {
      const v = go$a0;
      if (v.tag === "Leaf") {
        go$c = false;
        go$r = false;
        continue;
      }
      if (v.tag === "Node") {
        const v1 = Data$dOrd.ordString.compare(k)(v._3);
        if (v1 === "LT") {
          go$a0 = v._5;
          continue;
        }
        if (v1 === "GT") {
          go$a0 = v._6;
          continue;
        }
        if (v1 === "EQ") {
          go$c = false;
          go$r = true;
          continue;
        }
      }
      $runtime.fail();
    }
    return go$r;
  };
  return go;
};
const fromFoldable3 = /* #__PURE__ */ Data$dMap$dInternal.fromFoldable(Data$dOrd.ordString)(Data$dFoldable.foldableArray);
const fromFoldable4 = /* #__PURE__ */ Data$dMap$dInternal.fromFoldable(Data$dOrd.ordString)(Data$dFoldable.foldableArray);
const maximum = /* #__PURE__ */ Data$dSemigroup$dFoldable.maximum(Data$dOrd.ordInt)(Data$dArray$dNonEmpty$dInternal.foldable1NonEmptyArray);
const mapAccumR = /* #__PURE__ */ Data$dTraversable.mapAccumR(Data$dTraversable.traversableArray);
const foldMap5 = /* #__PURE__ */ (() => Data$dFoldable.foldableArray.foldMap(Data$dSet.monoidSet(ordQualified)))();
const fromFoldable5 = /* #__PURE__ */ Data$dFoldable.foldlArray(m => a => Data$dMap$dInternal.insert(PureScript$dBackend$dOptimizer$dCoreFn.ordReExport)(a)()(m))(Data$dMap$dInternal.Leaf);
const PatWild = /* #__PURE__ */ $PatternCase("PatWild");
const PatRecord = value0 => $PatternCase("PatRecord", value0);
const PatProduct = value0 => value1 => $PatternCase("PatProduct", value0, value1);
const PatArray = value0 => $PatternCase("PatArray", value0);
const PatSum = value0 => value1 => $PatternCase("PatSum", value0, value1);
const PatInt = value0 => $PatternCase("PatInt", value0);
const PatNumber = value0 => $PatternCase("PatNumber", value0);
const PatString = value0 => $PatternCase("PatString", value0);
const PatChar = value0 => $PatternCase("PatChar", value0);
const PatBoolean = value0 => $PatternCase("PatBoolean", value0);
const Pattern = x => x;
const UnconditionalFn = value0 => $CaseRowGuardedExpr("UnconditionalFn", value0);
const GuardedFn = value0 => $CaseRowGuardedExpr("GuardedFn", value0);
const newtypePattern_ = {Coercible0: () => {}};
const eqPatternCase = {
  eq: x => y => {
    if (x.tag === "PatWild") { return y.tag === "PatWild"; }
    if (x.tag === "PatRecord") { return y.tag === "PatRecord" && eq(x._1)(y._1); }
    if (x.tag === "PatProduct") {
      return y.tag === "PatProduct" && (x._1._1.tag === "Nothing" ? y._1._1.tag === "Nothing" : x._1._1.tag === "Just" && y._1._1.tag === "Just" && x._1._1._1 === y._1._1._1) && x._1._2 === y._1._2 && (x._2._1.tag === "Nothing"
        ? y._2._1.tag === "Nothing"
        : x._2._1.tag === "Just" && y._2._1.tag === "Just" && x._2._1._1 === y._2._1._1) && x._2._2 === y._2._2;
    }
    if (x.tag === "PatArray") { return y.tag === "PatArray" && x._1 === y._1; }
    if (x.tag === "PatSum") {
      return y.tag === "PatSum" && (x._1._1.tag === "Nothing" ? y._1._1.tag === "Nothing" : x._1._1.tag === "Just" && y._1._1.tag === "Just" && x._1._1._1 === y._1._1._1) && x._1._2 === y._1._2 && (x._2._1.tag === "Nothing"
        ? y._2._1.tag === "Nothing"
        : x._2._1.tag === "Just" && y._2._1.tag === "Just" && x._2._1._1 === y._2._1._1) && x._2._2 === y._2._2;
    }
    if (x.tag === "PatInt") { return y.tag === "PatInt" && x._1 === y._1; }
    if (x.tag === "PatNumber") { return y.tag === "PatNumber" && x._1 === y._1; }
    if (x.tag === "PatString") { return y.tag === "PatString" && x._1 === y._1; }
    if (x.tag === "PatChar") { return y.tag === "PatChar" && x._1 === y._1; }
    return x.tag === "PatBoolean" && y.tag === "PatBoolean" && x._1 === y._1;
  }
};
const ordPatternCase = {
  compare: x => y => {
    if (x.tag === "PatWild") {
      if (y.tag === "PatWild") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y.tag === "PatWild") { return Data$dOrdering.GT; }
    if (x.tag === "PatRecord") {
      if (y.tag === "PatRecord") { return compare(x._1)(y._1); }
      return Data$dOrdering.LT;
    }
    if (y.tag === "PatRecord") { return Data$dOrdering.GT; }
    if (x.tag === "PatProduct") {
      if (y.tag === "PatProduct") {
        const v = compare1(x._1)(y._1);
        if (v === "LT") { return Data$dOrdering.LT; }
        if (v === "GT") { return Data$dOrdering.GT; }
        return ordQualified.compare(x._2)(y._2);
      }
      return Data$dOrdering.LT;
    }
    if (y.tag === "PatProduct") { return Data$dOrdering.GT; }
    if (x.tag === "PatArray") {
      if (y.tag === "PatArray") { return Data$dOrd.ordInt.compare(x._1)(y._1); }
      return Data$dOrdering.LT;
    }
    if (y.tag === "PatArray") { return Data$dOrdering.GT; }
    if (x.tag === "PatSum") {
      if (y.tag === "PatSum") {
        const v = compare1(x._1)(y._1);
        if (v === "LT") { return Data$dOrdering.LT; }
        if (v === "GT") { return Data$dOrdering.GT; }
        return ordQualified.compare(x._2)(y._2);
      }
      return Data$dOrdering.LT;
    }
    if (y.tag === "PatSum") { return Data$dOrdering.GT; }
    if (x.tag === "PatInt") {
      if (y.tag === "PatInt") { return Data$dOrd.ordInt.compare(x._1)(y._1); }
      return Data$dOrdering.LT;
    }
    if (y.tag === "PatInt") { return Data$dOrdering.GT; }
    if (x.tag === "PatNumber") {
      if (y.tag === "PatNumber") { return Data$dOrd.ordNumber.compare(x._1)(y._1); }
      return Data$dOrdering.LT;
    }
    if (y.tag === "PatNumber") { return Data$dOrdering.GT; }
    if (x.tag === "PatString") {
      if (y.tag === "PatString") { return Data$dOrd.ordString.compare(x._1)(y._1); }
      return Data$dOrdering.LT;
    }
    if (y.tag === "PatString") { return Data$dOrdering.GT; }
    if (x.tag === "PatChar") {
      if (y.tag === "PatChar") { return Data$dOrd.ordChar.compare(x._1)(y._1); }
      return Data$dOrdering.LT;
    }
    if (y.tag === "PatChar") { return Data$dOrdering.GT; }
    if (x.tag === "PatBoolean" && y.tag === "PatBoolean") { return Data$dOrd.ordBoolean.compare(x._1)(y._1); }
    $runtime.fail();
  },
  Eq0: () => eqPatternCase
};
const monoidSet = /* #__PURE__ */ Data$dSet.monoidSet(ordPatternCase);
const foldMapWithIndex = /* #__PURE__ */ (() => Data$dFoldableWithIndex.foldableWithIndexArray.foldMapWithIndex((() => {
  const Semigroup0 = monoidSet.Semigroup0();
  const semigroupRecord1 = {
    append: ra => rb => ({aScore: ra.aScore + rb.aScore | 0, ctors: Semigroup0.append(ra.ctors)(rb.ctors), tailRowIndices: [...ra.tailRowIndices, ...rb.tailRowIndices]})
  };
  return {mempty: {aScore: 0, ctors: monoidSet.mempty, tailRowIndices: []}, Semigroup0: () => semigroupRecord1};
})()))();
const topEnv = v => ({...v, locals: []});
const toExternImpl = env => group => expr => {
  const $0 = () => {
    const v = PureScript$dBackend$dOptimizer$dSemantics.freeze(expr);
    return Data$dTuple.$Tuple(Data$dTuple.$Tuple(v._1, PureScript$dBackend$dOptimizer$dSemantics.$ExternImpl("ExternExpr", group, v._2)), v._2);
  };
  if (expr.tag === "ExprSyntax") {
    if (expr._2.tag === "Lit") {
      if (expr._2._1.tag === "LitRecord") {
        const propsWithAnalysis = Data$dFunctor.arrayMap(m => PureScript$dBackend$dOptimizer$dCoreFn.$Prop(m._1, PureScript$dBackend$dOptimizer$dSemantics.freeze(m._2)))(expr._2._1._1);
        return Data$dTuple.$Tuple(
          Data$dTuple.$Tuple(expr._1, PureScript$dBackend$dOptimizer$dSemantics.$ExternImpl("ExternDict", group, propsWithAnalysis)),
          PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
            "Lit",
            PureScript$dBackend$dOptimizer$dCoreFn.$Literal(
              "LitRecord",
              Data$dFunctor.arrayMap(m => PureScript$dBackend$dOptimizer$dCoreFn.$Prop(m._1, m._2._2))(propsWithAnalysis)
            )
          )
        );
      }
      return $0();
    }
    if (expr._2.tag === "CtorDef") {
      const v = PureScript$dBackend$dOptimizer$dSemantics.freeze(expr);
      return Data$dTuple.$Tuple(
        Data$dTuple.$Tuple(
          v._1,
          PureScript$dBackend$dOptimizer$dSemantics.$ExternImpl(
            "ExternCtor",
            (() => {
              const $1 = lookup(expr._2._2)(env.dataTypes);
              if ($1.tag === "Just") { return $1._1; }
              $runtime.fail();
            })(),
            expr._2._1,
            expr._2._2,
            expr._2._3,
            expr._2._4
          )
        ),
        v._2
      );
    }
  }
  return $0();
};
const toCaseRowVars = v => {
  const $0 = v.column;
  return foldMap(x => Data$dMap$dInternal.$$$Map("Node", 1, 1, x, $0, Data$dMap$dInternal.Leaf, Data$dMap$dInternal.Leaf))(v.pattern.vars);
};
const patternVars = v => [...toUnfoldable(v.pattern.vars), ...foldMap2(patternVars)(v.pattern.subterms)];
const patternSubterms = v => v.pattern.subterms;
const patternPatCase = v => v.pattern.patternCase;
const normalizeCaseRows = x => {
  const go = go$a0$copy => go$a1$copy => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
    while (go$c) {
      const columnIdx = go$a0, columnsAcc = go$a1;
      const nextColumnFields = Data$dFoldable.foldlArray(acc => next => {
        if (columnIdx >= 0 && columnIdx < next.patterns.length) {
          const $0 = next.patterns[columnIdx];
          return Data$dMaybe.$Maybe(
            "Just",
            (() => {
              if ($0.pattern.patternCase.tag === "PatRecord") {
                const keys = fromFoldable1($0.pattern.patternCase._1);
                if (acc.tag === "Nothing") { return keys; }
                if (acc.tag === "Just") { return Data$dMap$dInternal.unsafeUnionWith(Data$dOrd.ordString.compare, Data$dFunction.const, keys, acc._1); }
                $runtime.fail();
              }
              if (acc.tag === "Nothing") { return Data$dMap$dInternal.Leaf; }
              if (acc.tag === "Just") { return acc._1; }
              $runtime.fail();
            })()
          );
        }
        return Data$dMaybe.Nothing;
      })(Data$dMaybe.Nothing)(x);
      if (nextColumnFields.tag === "Nothing") {
        go$c = false;
        go$r = columnsAcc;
        continue;
      }
      if (nextColumnFields.tag === "Just") {
        go$a0 = columnIdx + 1 | 0;
        go$a1 = Data$dArray.snoc(columnsAcc)(nextColumnFields._1);
        continue;
      }
      $runtime.fail();
    }
    return go$r;
  };
  const $0 = go(0)([]);
  return Data$dFunctor.arrayMap(nextRow => (
    {
      ...nextRow,
      patterns: Data$dArray.zipWithImpl(
        allFieldsSet => pat => {
          if (pat.pattern.patternCase.tag === "PatRecord") {
            const v1 = Data$dArray.unzip(Data$dFunctor.arrayMap(Data$dArray$dNonEmpty.head)(Data$dArray.groupAllBy(x$1 => y => Data$dOrd.ordString.compare(x$1._1)(y._1))([
              ...Data$dArray.zipWithImpl(Data$dTuple.Tuple, pat.pattern.patternCase._1, pat.pattern.subterms),
              ...Data$dFunctor.arrayMap(fieldName => Data$dTuple.$Tuple(
                fieldName,
                {
                  accessor: PureScript$dBackend$dOptimizer$dSyntax.$BackendAccessor("GetProp", fieldName),
                  pattern: {vars: Data$dMap$dInternal.Leaf, patternCase: PatWild, subterms: []}
                }
              ))(toUnfoldable(allFieldsSet))
            ])));
            return {...pat, pattern: {...pat.pattern, patternCase: $PatternCase("PatRecord", v1._1), subterms: v1._2}};
          }
          return pat;
        },
        $0,
        nextRow.patterns
      )
    }
  ))(x);
};
const makeExternEvalSpine = conv => env => qual => spine => {
  const $0 = lookup1(qual)(conv.foreignSemantics);
  const result = (() => {
    if ($0.tag === "Just") { return $0._1(env)(qual)(spine); }
    if ($0.tag === "Nothing") { return Data$dMaybe.Nothing; }
    $runtime.fail();
  })();
  if (result.tag === "Nothing") {
    const $1 = lookup1(qual)(conv.implementations);
    if ($1.tag === "Just") { return PureScript$dBackend$dOptimizer$dSemantics.evalExternFromImpl({...env, locals: []})(qual)($1._1)(spine); }
    if ($1.tag === "Nothing") { return Data$dMaybe.Nothing; }
    $runtime.fail();
  }
  return result;
};
const makeExternEvalRef = conv => env => qual => {
  const $0 = lookup1(qual)(conv.implementations);
  if ($0.tag === "Just") { return Data$dMaybe.$Maybe("Just", PureScript$dBackend$dOptimizer$dSemantics.evalExternRefFromImpl(env)(qual)($0._1)); }
  return Data$dMaybe.Nothing;
};
const levelUp = f => env => f({...env, currentLevel: env.currentLevel + 1 | 0});
const intro = dictFoldable => ident => lvl => f => env => f({
  ...env,
  currentLevel: env.currentLevel + 1 | 0,
  toLevel: dictFoldable.foldr(a => Data$dMap$dInternal.insert(Data$dOrd.ordString)(a)(lvl))(env.toLevel)(ident)
});
const inferTransitiveDirective = directives => impl => backendExpr => cfn => {
  const $0 = (() => {
    if (impl.tag === "ExternExpr") {
      if (impl._2.tag === "App") {
        if (impl._2._1.tag === "Var") {
          const v = lookup2(PureScript$dBackend$dOptimizer$dSemantics.$EvalRef("EvalExtern", impl._2._1._1))(directives);
          if (v.tag === "Just") {
            const go = (m$p, z$p) => {
              if (m$p.tag === "Leaf") { return z$p; }
              if (m$p.tag === "Node") {
                return go(
                  m$p._5,
                  (() => {
                    const $0 = m$p._4;
                    const $1 = go(m$p._6, z$p);
                    const $2 = prop => Data$dMap$dInternal.insert(PureScript$dBackend$dOptimizer$dSemantics.ordInlineAccessor)(PureScript$dBackend$dOptimizer$dSemantics.$InlineAccessor(
                      "InlineSpineProp",
                      prop
                    ))($0)(Data$dMap$dInternal.insert(PureScript$dBackend$dOptimizer$dSemantics.ordInlineAccessor)(PureScript$dBackend$dOptimizer$dSemantics.$InlineAccessor(
                      "InlineProp",
                      prop
                    ))($0)($1));
                    if ($0.tag === "InlineArity" && m$p._3.tag === "InlineRef") {
                      return Data$dMap$dInternal.insert(PureScript$dBackend$dOptimizer$dSemantics.ordInlineAccessor)(PureScript$dBackend$dOptimizer$dSemantics.InlineRef)(PureScript$dBackend$dOptimizer$dSemantics.$InlineDirective(
                        "InlineArity",
                        $0._1 - impl._2._2.length | 0
                      ))($1);
                    }
                    if (m$p._3.tag === "InlineSpineProp") { return $2(m$p._3._1); }
                    return $1;
                  })()
                );
              }
              $runtime.fail();
            };
            const newDirs = go(v._1, Data$dMap$dInternal.Leaf);
            if (newDirs.tag === "Leaf") { return Data$dMaybe.Nothing; }
            return Data$dMaybe.$Maybe("Just", newDirs);
          }
        }
        return Data$dMaybe.Nothing;
      }
      if (impl._2.tag === "Accessor" && impl._2._1.tag === "App" && impl._2._1._1.tag === "Var" && impl._2._2.tag === "GetProp") {
        const $0 = lookup2(PureScript$dBackend$dOptimizer$dSemantics.$EvalRef("EvalExtern", impl._2._1._1._1))(directives);
        const $1 = lookup3(PureScript$dBackend$dOptimizer$dSemantics.$InlineAccessor("InlineSpineProp", impl._2._2._1));
        const v = (() => {
          if ($0.tag === "Just") { return $1($0._1); }
          if ($0.tag === "Nothing") { return Data$dMaybe.Nothing; }
          $runtime.fail();
        })();
        if (v.tag === "Just" && v._1.tag === "InlineArity") {
          return Data$dMaybe.$Maybe(
            "Just",
            Data$dMap$dInternal.$$$Map(
              "Node",
              1,
              1,
              PureScript$dBackend$dOptimizer$dSemantics.InlineRef,
              PureScript$dBackend$dOptimizer$dSemantics.$InlineDirective("InlineArity", v._1._1),
              Data$dMap$dInternal.Leaf,
              Data$dMap$dInternal.Leaf
            )
          );
        }
      }
    }
    return Data$dMaybe.Nothing;
  })();
  const $1 = (() => {
    if (backendExpr.tag === "ExprSyntax" && backendExpr._2.tag === "App" && backendExpr._2._1.tag === "ExprSyntax" && backendExpr._2._1._2.tag === "Var") {
      const $1 = lookup2(PureScript$dBackend$dOptimizer$dSemantics.$EvalRef("EvalExtern", backendExpr._2._1._2._1))(directives);
      const $2 = lookup3(PureScript$dBackend$dOptimizer$dSemantics.InlineRef);
      const v = (() => {
        if ($1.tag === "Just") { return $2($1._1); }
        if ($1.tag === "Nothing") { return Data$dMaybe.Nothing; }
        $runtime.fail();
      })();
      if (
        v.tag === "Just" && v._1.tag === "InlineArity" && cfn.tag === "ExprApp" && cfn._1.meta.tag === "Just" && cfn._1.meta._1.tag === "IsSyntheticApp" && backendExpr._2._2.length >= v._1._1
      ) {
        return Data$dMaybe.$Maybe(
          "Just",
          Data$dMap$dInternal.$$$Map(
            "Node",
            1,
            1,
            PureScript$dBackend$dOptimizer$dSemantics.InlineRef,
            PureScript$dBackend$dOptimizer$dSemantics.InlineAlways,
            Data$dMap$dInternal.Leaf,
            Data$dMap$dInternal.Leaf
          )
        );
      }
    }
    return Data$dMaybe.Nothing;
  })();
  if ($0.tag === "Nothing") { return $1; }
  return $0;
};
const guardTag = n => lhs => PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
  "PrimOp",
  PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op1", PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator1("OpIsTag", n), lhs)
);
const getCtx = env => {
  const lookupExtern = qual => acc => {
    const $0 = lookup1(qual)(env.implementations);
    if ($0.tag === "Just") {
      if ($0._1._2.tag === "ExternExpr") {
        if (acc.tag === "Nothing") { return Data$dMaybe.$Maybe("Just", $0._1._1); }
        return Data$dMaybe.Nothing;
      }
      if ($0._1._2.tag === "ExternDict") {
        if (acc.tag === "Just") {
          const $1 = acc._1;
          const $2 = Data$dArray.findMapImpl(
            Data$dMaybe.Nothing,
            Data$dMaybe.isJust,
            v => {
              if ($1 === v._1) { return Data$dMaybe.$Maybe("Just", v._2); }
              return Data$dMaybe.Nothing;
            },
            $0._1._2._2
          );
          if ($2.tag === "Just") { return Data$dMaybe.$Maybe("Just", $2._1._1); }
          return Data$dMaybe.Nothing;
        }
        if (acc.tag === "Nothing") { return Data$dMaybe.$Maybe("Just", $0._1._1); }
        $runtime.fail();
      }
      if ($0._1._2.tag === "ExternCtor") { return Data$dMaybe.Nothing; }
      $runtime.fail();
    }
    if ($0.tag === "Nothing") { return Data$dMaybe.Nothing; }
    $runtime.fail();
  };
  return {
    currentLevel: env.currentLevel,
    lookupExtern,
    analyze: v => {
      const $0 = v.effect;
      return expr => {
        const v1 = env.analyzeCustom(v)(expr);
        if (v1.tag === "Just") { return v1._1; }
        if (v1.tag === "Nothing") {
          if ($0) { return analyzeEffectBlock(lookupExtern)(expr); }
          return analyze(lookupExtern)(expr);
        }
        $runtime.fail();
      };
    },
    effect: false
  };
};
const fromExternImpl = v => {
  if (v.tag === "ExternExpr") { return Data$dMaybe.$Maybe("Just", v._2); }
  if (v.tag === "ExternDict") { return Data$dMaybe.Nothing; }
  if (v.tag === "ExternCtor") { return Data$dMaybe.Nothing; }
  $runtime.fail();
};
const decompose = chosenColumn => {
  const checkMatch = p => {
    if (p.column === chosenColumn.column && (eqPatternCase.eq(p.pattern.patternCase)(PatWild) || eqPatternCase.eq(chosenColumn.pattern.patternCase)(p.pattern.patternCase))) {
      return {nonMatchesBefore: [], match: Data$dMaybe.$Maybe("Just", {match: p, nonMatchesAfter: []})};
    }
    return {nonMatchesBefore: [p], match: Data$dMaybe.Nothing};
  };
  return foldMap3(row => {
    const v = row.patterns.length > 0 ? Data$dMaybe.$Maybe("Just", row.patterns) : Data$dMaybe.Nothing;
    if (v.tag === "Nothing") { return Partial._crashWith("decompose - nextRow.patterns cannot be empty since the first row contains at least one `PatCtor` patternCase"); }
    if (v.tag === "Just") {
      const $0 = v._1;
      const len = $0.length;
      const v1 = (() => {
        const go = go$a0$copy => go$a1$copy => {
          let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
          while (go$c) {
            const ix = go$a0, acc = go$a1;
            if (ix === len) {
              go$c = false;
              go$r = acc;
              continue;
            }
            go$a0 = ix + 1 | 0;
            go$a1 = (() => {
              const $1 = checkMatch($0[ix]);
              if (acc.match.tag === "Just") {
                if ($1.match.tag === "Just") { return Partial._crashWith("mergeResults - impossible: cannot match the same column twice in the same row"); }
                if ($1.match.tag === "Nothing") {
                  return {...acc, match: Data$dMaybe.$Maybe("Just", {...acc.match._1, nonMatchesAfter: [...acc.match._1.nonMatchesAfter, ...$1.nonMatchesBefore]})};
                }
                $runtime.fail();
              }
              if (acc.match.tag === "Nothing") {
                if ($1.match.tag === "Nothing") { return {...$1, nonMatchesBefore: [...acc.nonMatchesBefore, ...$1.nonMatchesBefore]}; }
                if ($1.match.tag === "Just") { return {...$1, nonMatchesBefore: [...acc.nonMatchesBefore, ...$1.nonMatchesBefore]}; }
              }
              $runtime.fail();
            })();
          }
          return go$r;
        };
        return go(1)(checkMatch((() => {
          if (0 < $0.length) { return $0[0]; }
          $runtime.fail();
        })()));
      })();
      if (v1.match.tag === "Just") {
        return {
          rowsWithMatch: [{guardFn: row.guardFn, vars: row.vars, nonMatchesBefore: v1.nonMatchesBefore, match: v1.match._1.match, nonMatchesAfter: v1.match._1.nonMatchesAfter}],
          rowsNoMatch: eqPatternCase.eq(v1.match._1.match.pattern.patternCase)(PatWild) ? [row] : []
        };
      }
      if (v1.match.tag === "Nothing") { return {rowsWithMatch: [], rowsNoMatch: [row]}; }
    }
    $runtime.fail();
  });
};
const currentLevel = env => env.currentLevel;
const chooseNextPattern = row0Patterns => tailRows => {
  const expandIfPossible = Data$dArray.findMapImpl(
    Data$dMaybe.Nothing,
    Data$dMaybe.isJust,
    v => {
      if (v._2.pattern.patternCase.tag === "PatRecord") { return Data$dMaybe.$Maybe("Just", v._2); }
      if (v._2.pattern.patternCase.tag === "PatProduct") { return Data$dMaybe.$Maybe("Just", v._2); }
      return Data$dMaybe.Nothing;
    },
    row0Patterns
  );
  if (expandIfPossible.tag === "Just") { return expandIfPossible._1; }
  if (expandIfPossible.tag === "Nothing") {
    const $0 = Data$dFoldable.foldlArray(acc => next => {
      if (acc.tag === "Nothing") { return Data$dMaybe.$Maybe("Just", [next]); }
      if (acc.tag === "Just") {
        const v = Data$dOrd.ordInt.compare((() => {
          if (0 < acc._1.length) { return acc._1[0].bScore; }
          $runtime.fail();
        })())(next.bScore);
        if (v === "GT") { return acc; }
        if (v === "EQ") { return Data$dMaybe.$Maybe("Just", Data$dArray.snoc(acc._1)(next)); }
        if (v === "LT") { return Data$dMaybe.$Maybe("Just", [next]); }
      }
      $runtime.fail();
    })(Data$dMaybe.Nothing);
    const $1 = Data$dFoldable.foldlArray(acc => next => {
      if (acc.tag === "Nothing") { return Data$dMaybe.$Maybe("Just", [next]); }
      if (acc.tag === "Just") {
        const v = Data$dOrd.ordInt.compare((() => {
          if (0 < acc._1.length) { return acc._1[0].aScore; }
          $runtime.fail();
        })())(next.aScore);
        if (v === "GT") { return acc; }
        if (v === "EQ") { return Data$dMaybe.$Maybe("Just", Data$dArray.snoc(acc._1)(next)); }
        if (v === "LT") { return Data$dMaybe.$Maybe("Just", [next]); }
      }
      $runtime.fail();
    })(Data$dMaybe.Nothing);
    const $2 = Data$dFoldable.foldlArray(acc => next => {
      if (acc.tag === "Nothing") { return Data$dMaybe.$Maybe("Just", [next]); }
      if (acc.tag === "Just") {
        const v = Data$dOrd.ordInt.compare((() => {
          if (0 < acc._1.length) { return acc._1[0].pScore; }
          $runtime.fail();
        })())(next.pScore);
        if (v === "GT") { return acc; }
        if (v === "EQ") { return Data$dMaybe.$Maybe("Just", Data$dArray.snoc(acc._1)(next)); }
        if (v === "LT") { return Data$dMaybe.$Maybe("Just", [next]); }
      }
      $runtime.fail();
    })(Data$dMaybe.Nothing)(Data$dFunctor.arrayMap(v => {
      const $2 = v._1;
      const $3 = v._2;
      const matchingCols = foldMapWithIndex(rowIdx => row => {
        if ($2 >= 0 && $2 < row.patterns.length) {
          return {
            tailRowIndices: (() => {
              const $4 = [rowIdx + 1 | 0];
              if (eqPatternCase.eq($3.pattern.patternCase)(row.patterns[$2].pattern.patternCase)) { return $4; }
              return [];
            })(),
            ctors: (() => {
              const $4 = Data$dMap$dInternal.$$$Map("Node", 1, 1, row.patterns[$2].pattern.patternCase, undefined, Data$dMap$dInternal.Leaf, Data$dMap$dInternal.Leaf);
              if (!eqPatternCase.eq(row.patterns[$2].pattern.patternCase)(PatWild)) { return $4; }
              return monoidSet.mempty;
            })(),
            aScore: -Data$dArray.filterImpl(x => !eqPatternCase.eq(PatWild)(x.pattern.patternCase), row.patterns[$2].pattern.subterms).length
          };
        }
        return Partial._crashWith("Impossible: rows' column lengths differ in pattern match");
      })(tailRows);
      return {
        pattern: $3,
        pScore: Data$dFoldable.foldlArray(l => r => {
          if ((l + 1 | 0) === r) { return r; }
          return l;
        })(0)(matchingCols.tailRowIndices),
        bScore: (() => {
          const $4 = Data$dMap$dInternal.insert(ordPatternCase)($3.pattern.patternCase)()(matchingCols.ctors);
          if ($4.tag === "Leaf") { return 0; }
          if ($4.tag === "Node") { return -$4._2; }
          $runtime.fail();
        })(),
        aScore: matchingCols.aScore
      };
    })(row0Patterns));
    if ($2.tag === "Just") {
      const $3 = $0($2._1);
      if ($3.tag === "Just") {
        const $4 = $1($3._1);
        if ($4.tag === "Just") {
          if (0 < $4._1.length) { return $4._1[0].pattern; }
          $runtime.fail();
        }
        if (0 < row0Patterns.length) { return row0Patterns[0]._2; }
        $runtime.fail();
      }
      if ($3.tag === "Nothing" && 0 < row0Patterns.length) { return row0Patterns[0]._2; }
      $runtime.fail();
    }
    if ($2.tag === "Nothing" && 0 < row0Patterns.length) { return row0Patterns[0]._2; }
  }
  $runtime.fail();
};
const buildM = a => env => PureScript$dBackend$dOptimizer$dSemantics.build(getCtx(env))(a);
const make = a => {
  const $0 = PureScript$dBackend$dOptimizer$dSyntax.traversableBackendSyntax.traverse(Control$dApplicative.applicativeFn)(Data$dTraversable.identity)(a);
  return x => PureScript$dBackend$dOptimizer$dSemantics.build(getCtx(x))($0(x));
};
const guardBoolean = n => lhs => PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
  "PrimOp",
  PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator(
    "Op2",
    PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpBooleanOrd", PureScript$dBackend$dOptimizer$dSyntax.OpEq),
    lhs,
    make(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Lit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitBoolean", n)))
  )
);
const guardChar = n => lhs => PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
  "PrimOp",
  PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator(
    "Op2",
    PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpCharOrd", PureScript$dBackend$dOptimizer$dSyntax.OpEq),
    lhs,
    make(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Lit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitChar", n)))
  )
);
const guardInt = n => lhs => PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
  "PrimOp",
  PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator(
    "Op2",
    PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpIntOrd", PureScript$dBackend$dOptimizer$dSyntax.OpEq),
    lhs,
    make(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Lit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitInt", n)))
  )
);
const guardArrayLength = n => lhs => PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
  "PrimOp",
  PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator(
    "Op2",
    PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpIntOrd", PureScript$dBackend$dOptimizer$dSyntax.OpEq),
    make(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
      "PrimOp",
      PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op1", PureScript$dBackend$dOptimizer$dSyntax.OpArrayLength, lhs)
    )),
    make(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Lit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitInt", n)))
  )
);
const guardNumber = n => lhs => PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
  "PrimOp",
  PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator(
    "Op2",
    PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpNumberOrd", PureScript$dBackend$dOptimizer$dSyntax.OpEq),
    lhs,
    make(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Lit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitNumber", n)))
  )
);
const guardString = n => lhs => PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
  "PrimOp",
  PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator(
    "Op2",
    PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpStringOrd", PureScript$dBackend$dOptimizer$dSyntax.OpEq),
    lhs,
    make(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Lit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitString", n)))
  )
);
const makeGuard = lvl => g => inner => def => make(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
  "Branch",
  [PureScript$dBackend$dOptimizer$dSyntax.$Pair(make(g(make(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Local", Data$dMaybe.Nothing, lvl)))), inner)],
  def
));
const makeLet = id => a => k => x => {
  if (id.tag === "Nothing") {
    return make(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
      "Let",
      id,
      x.currentLevel,
      a,
      (() => {
        const $0 = k(x.currentLevel);
        return env => $0({...env, currentLevel: env.currentLevel + 1 | 0});
      })()
    ))(x);
  }
  if (id.tag === "Just") {
    return make(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Let", id, x.currentLevel, a, intro(Data$dFoldable.foldableArray)([id._1])(x.currentLevel)(k(x.currentLevel))))(x);
  }
  $runtime.fail();
};
const makeUncurriedAbs = args => cb => Data$dFoldable.foldrArray(ident => next => tmps => x => intro(Data$dFoldable.foldableArray)([ident])(x.currentLevel)(next(Data$dArray.snoc(tmps)(Data$dTuple.$Tuple(
  Data$dMaybe.$Maybe("Just", ident),
  x.currentLevel
))))(x))(tmps => make(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("UncurriedAbs", tmps, cb(tmps))))(args)([]);
const patternFail = /* #__PURE__ */ make(/* #__PURE__ */ PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Fail", "Failed pattern match"));
const binderToPattern = v => {
  if (v.tag === "BinderNull") { return v$1 => ({vars: Data$dMap$dInternal.Leaf, patternCase: PatWild, subterms: []}); }
  if (v.tag === "BinderVar") {
    return v$1 => ({vars: Data$dMap$dInternal.$$$Map("Node", 1, 1, v._2, undefined, Data$dMap$dInternal.Leaf, Data$dMap$dInternal.Leaf), patternCase: PatWild, subterms: []});
  }
  if (v.tag === "BinderNamed") {
    const $0 = v._2;
    const $1 = binderToPattern(v._3);
    return x => {
      const $2 = $1(x);
      return {...$2, vars: Data$dMap$dInternal.insert(Data$dOrd.ordString)($0)()($2.vars)};
    };
  }
  if (v.tag === "BinderLit") {
    if (v._2.tag === "LitInt") { return v$1 => ({vars: Data$dMap$dInternal.Leaf, patternCase: $PatternCase("PatInt", v._2._1), subterms: []}); }
    if (v._2.tag === "LitNumber") { return v$1 => ({vars: Data$dMap$dInternal.Leaf, patternCase: $PatternCase("PatNumber", v._2._1), subterms: []}); }
    if (v._2.tag === "LitString") { return v$1 => ({vars: Data$dMap$dInternal.Leaf, patternCase: $PatternCase("PatString", v._2._1), subterms: []}); }
    if (v._2.tag === "LitChar") { return v$1 => ({vars: Data$dMap$dInternal.Leaf, patternCase: $PatternCase("PatChar", v._2._1), subterms: []}); }
    if (v._2.tag === "LitBoolean") { return v$1 => ({vars: Data$dMap$dInternal.Leaf, patternCase: $PatternCase("PatBoolean", v._2._1), subterms: []}); }
    if (v._2.tag === "LitArray") {
      const $0 = $PatternCase("PatArray", v._2._1.length);
      const $1 = forWithIndex(v._2._1)(idx => nextArg => {
        const $1 = binderToPattern(nextArg);
        return x => ({accessor: PureScript$dBackend$dOptimizer$dSyntax.$BackendAccessor("GetIndex", idx), pattern: $1(x)});
      });
      return x => ({vars: Data$dMap$dInternal.Leaf, patternCase: $0, subterms: $1(x)});
    }
    if (v._2.tag === "LitRecord") {
      const $0 = $PatternCase("PatRecord", Data$dFunctor.arrayMap(PureScript$dBackend$dOptimizer$dCoreFn.propKey)(v._2._1));
      const $1 = forWithIndex(v._2._1)(idx => nextArg => {
        const $1 = binderToPattern(nextArg._2);
        return x => ({accessor: PureScript$dBackend$dOptimizer$dSyntax.$BackendAccessor("GetProp", nextArg._1), pattern: $1(x)});
      });
      return x => ({vars: Data$dMap$dInternal.Leaf, patternCase: $0, subterms: $1(x)});
    }
    $runtime.fail();
  }
  if (v.tag === "BinderConstructor") {
    if (v._1.meta.tag === "Just" && v._1.meta._1.tag === "IsNewtype" && v._4.length === 1) { return binderToPattern(v._4[0]); }
    if (v._1.meta.tag === "Just") {
      if (v._1.meta._1.tag === "IsNewtype") { return Partial._crashWith("Newtype binder didn't wrap 1 arg"); }
      if (v._1.meta._1.tag === "IsConstructor") {
        if (v._1.meta._1._1 === "ProductType") {
          return x => {
            const v$1 = lookup1(v._3)(x.implementations);
            return {
              vars: Data$dMap$dInternal.Leaf,
              patternCase: $PatternCase("PatProduct", v._2, v._3),
              subterms: forWithIndex(Data$dArray.zipWithImpl(
                Data$dTuple.Tuple,
                v._4,
                (() => {
                  if (v$1.tag === "Just" && v$1._1._2.tag === "ExternCtor") {
                    const $0 = lookup(v._2._2)(x.dataTypes);
                    if ($0.tag === "Just") { return v$1._1._2._5; }
                    return v$1._1._2._5;
                  }
                  const $0 = lookup(v._2._2)(x.dataTypes);
                  if ($0.tag === "Just") {
                    const $1 = lookup4(v._3._2)($0._1.constructors);
                    if ($1.tag === "Just") { return $1._1.fields; }
                    return Partial._crashWith("Invariant broken: could not determine pattern matched constructor's fields during conversion.")(x);
                  }
                  if ($0.tag === "Nothing") { return Partial._crashWith("Invariant broken: could not determine pattern matched constructor's fields during conversion.")(x); }
                  $runtime.fail();
                })()
              ))(idx => nextArg => {
                const $0 = binderToPattern(nextArg._1);
                return x$1 => (
                  {
                    accessor: PureScript$dBackend$dOptimizer$dSyntax.$BackendAccessor(
                      "GetCtorField",
                      v._3,
                      PureScript$dBackend$dOptimizer$dCoreFn.ProductType,
                      v._2._2,
                      v._3._2,
                      nextArg._2,
                      idx
                    ),
                    pattern: $0(x$1)
                  }
                );
              })(x)
            };
          };
        }
        if (v._1.meta._1._1 === "SumType") {
          return x => {
            const v$1 = lookup1(v._3)(x.implementations);
            return {
              vars: Data$dMap$dInternal.Leaf,
              patternCase: $PatternCase("PatSum", v._2, v._3),
              subterms: forWithIndex(Data$dArray.zipWithImpl(
                Data$dTuple.Tuple,
                v._4,
                (() => {
                  if (v$1.tag === "Just" && v$1._1._2.tag === "ExternCtor") {
                    const $0 = lookup(v._2._2)(x.dataTypes);
                    if ($0.tag === "Just") { return v$1._1._2._5; }
                    return v$1._1._2._5;
                  }
                  const $0 = lookup(v._2._2)(x.dataTypes);
                  if ($0.tag === "Just") {
                    const $1 = lookup4(v._3._2)($0._1.constructors);
                    if ($1.tag === "Just") { return $1._1.fields; }
                    return Partial._crashWith("Invariant broken: could not determine pattern matched constructor's fields during conversion.")(x);
                  }
                  if ($0.tag === "Nothing") { return Partial._crashWith("Invariant broken: could not determine pattern matched constructor's fields during conversion.")(x); }
                  $runtime.fail();
                })()
              ))(idx => nextArg => {
                const $0 = binderToPattern(nextArg._1);
                return x$1 => (
                  {
                    accessor: PureScript$dBackend$dOptimizer$dSyntax.$BackendAccessor(
                      "GetCtorField",
                      v._3,
                      PureScript$dBackend$dOptimizer$dCoreFn.SumType,
                      v._2._2,
                      v._3._2,
                      nextArg._2,
                      idx
                    ),
                    pattern: $0(x$1)
                  }
                );
              })(x)
            };
          };
        }
      }
    }
    return Partial._crashWith("binderToPattern - invalid meta");
  }
  $runtime.fail();
};
const toBackendExpr = v => {
  if (v.tag === "ExprVar") {
    const $0 = v._2;
    return x => {
      const $1 = x.currentModule;
      const $2 = x.toLevel;
      const v2 = v3 => {
        const v4 = v5 => {
          if ($0._2 === "undefined" && $0._1.tag === "Just") {
            if ($0._1._1 === "Prim") { return buildM(PureScript$dBackend$dOptimizer$dSyntax.PrimUndefined); }
            return buildM(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Var", $0));
          }
          if ($0._1.tag === "Nothing") {
            return buildM(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Var", PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", $1), $0._2)));
          }
          return buildM(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Var", $0));
        };
        if ($0._1.tag === "Just" && $0._1._1 === $1) {
          const $3 = lookup4($0._2)($2);
          if ($3.tag === "Just") { return buildM(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Local", Data$dMaybe.$Maybe("Just", $0._2), $3._1)); }
        }
        return v4(true);
      };
      if ($0._1.tag === "Nothing") {
        const $3 = lookup4($0._2)($2);
        if ($3.tag === "Just") {
          return PureScript$dBackend$dOptimizer$dSemantics.build(getCtx(x))(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Local", Data$dMaybe.$Maybe("Just", $0._2), $3._1));
        }
      }
      return v2(true)(x);
    };
  }
  if (v.tag === "ExprLit") {
    const $0 = traverse(toBackendExpr)(v._2);
    return x => PureScript$dBackend$dOptimizer$dSemantics.build(getCtx(x))(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Lit", $0(x)));
  }
  if (v.tag === "ExprConstructor") {
    const $0 = v._4;
    const $1 = v._3;
    const $2 = v._2;
    return x => PureScript$dBackend$dOptimizer$dSemantics.build(getCtx(x))(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
      "CtorDef",
      (() => {
        const v2 = lookup($2)(x.dataTypes);
        if (
          v2.tag === "Just" && (() => {
            if (v2._1.constructors.tag === "Leaf") { return false; }
            if (v2._1.constructors.tag === "Node") { return v2._1.constructors._2 === 1; }
            $runtime.fail();
          })()
        ) {
          return PureScript$dBackend$dOptimizer$dCoreFn.ProductType;
        }
        return PureScript$dBackend$dOptimizer$dCoreFn.SumType;
      })(),
      $2,
      $1,
      $0
    ));
  }
  if (v.tag === "ExprAccessor") {
    const $0 = toBackendExpr(v._2);
    return x => PureScript$dBackend$dOptimizer$dSemantics.build(getCtx(x))(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
      "Accessor",
      $0(x),
      PureScript$dBackend$dOptimizer$dSyntax.$BackendAccessor("GetProp", v._3)
    ));
  }
  if (v.tag === "ExprUpdate") {
    const $0 = toBackendExpr(v._2);
    const $1 = traverse1(PureScript$dBackend$dOptimizer$dCoreFn.traversableProp.traverse(Control$dApplicative.applicativeFn)(toBackendExpr))(v._3);
    return x => PureScript$dBackend$dOptimizer$dSemantics.build(getCtx(x))(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Update", $0(x), $1(x)));
  }
  if (v.tag === "ExprAbs") {
    const $0 = v._2;
    const $1 = v._3;
    return x => make(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
      "Abs",
      [Data$dTuple.$Tuple(Data$dMaybe.$Maybe("Just", $0), x.currentLevel)],
      intro(Data$dFoldable.foldableArray)([$0])(x.currentLevel)(toBackendExpr($1))
    ))(x);
  }
  if (v.tag === "ExprApp" && v._2.tag === "ExprVar" && v._2._1.meta.tag === "Just" && v._2._1.meta._1.tag === "IsNewtype") { return toBackendExpr(v._3); }
  if (v.tag === "ExprApp") { return make(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("App", toBackendExpr(v._2), [toBackendExpr(v._3)])); }
  if (v.tag === "ExprLet") {
    return Data$dFoldable.foldrArray(bind$p => next => {
      if (bind$p.tag === "NonRec") { return makeLet(Data$dMaybe.$Maybe("Just", bind$p._1._2))(toBackendExpr(bind$p._1._3))(v3 => next); }
      if (bind$p.tag === "Rec" && bind$p._1.length > 0) {
        const $0 = bind$p._1;
        return x => {
          const idents = Data$dFunctor.arrayMap(v4 => v4._2)($0);
          return PureScript$dBackend$dOptimizer$dSemantics.build(getCtx(x))(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
            "LetRec",
            x.currentLevel,
            intro(Data$dFoldable.foldableArray)(idents)(x.currentLevel)(traverse3(toBackendBinding)($0))(x),
            intro(Data$dFoldable.foldableArray)(idents)(x.currentLevel)(next)(x)
          ));
        };
      }
      if (bind$p.tag === "Rec") { return Partial._crashWith("CoreFn empty Rec binding group"); }
      $runtime.fail();
    })(toBackendExpr(v._3))(v._2);
  }
  if (v.tag === "ExprCase") {
    const $0 = v._3;
    return Data$dFoldable.foldrArray(expr => next => idents => makeLet(Data$dMaybe.Nothing)(toBackendExpr(expr))(tmp => next(Data$dArray.snoc(idents)(tmp))))(idents => Data$dFoldable.foldrArray(v$1 => {
      const $1 = v$1._1;
      const $2 = v$1._2;
      return mainCb => caseRows => {
        const $3 = zipWithA(ident => b => {
          const $3 = binderToPattern(b);
          return x => ({column: ident, pattern: $3(x)});
        })(idents)($1);
        return x => {
          const $4 = $3(x);
          const args = Data$dArray.sortBy(Data$dOrd.ordString.compare)(foldMap2(patternVars)($4));
          if ($2.tag === "Unconditional") {
            const $5 = $2._1;
            return makeLet(Data$dMaybe.Nothing)(makeUncurriedAbs(args)(v1 => toBackendExpr($5)))(tmp => mainCb(Data$dArray.snoc(caseRows)({
              patterns: $4,
              guardFn: $CaseRowGuardedExpr("UnconditionalFn", tmp),
              vars: Data$dMap$dInternal.Leaf
            })))(x);
          }
          if ($2.tag === "Guarded") {
            return Data$dFoldable.foldrArray(v1 => {
              const $5 = v1._2;
              const $6 = v1._1;
              return cb => xs => makeLet(Data$dMaybe.Nothing)(makeUncurriedAbs(args)(v2 => toBackendExpr($5)))(tmp => cb(Data$dArray.snoc(xs)(Data$dTuple.$Tuple($6, tmp))));
            })(xs => {
              if (xs.length > 0) { return mainCb(Data$dArray.snoc(caseRows)({patterns: $4, guardFn: $CaseRowGuardedExpr("GuardedFn", xs), vars: Data$dMap$dInternal.Leaf})); }
              return Partial._crashWith("CoreFn empty Guarded");
            })($2._1)([])(x);
          }
          $runtime.fail();
        };
      };
    })(caseRows => buildCaseTreeFromRows(caseRows))($0)([]))(v._2)([]);
  }
  $runtime.fail();
};
const toBackendBinding = v => {
  const $0 = Data$dTuple.Tuple(v._2);
  const $1 = toBackendExpr(v._3);
  return x => $0($1(x));
};
const buildCaseTreeFromRows = denormalizedRows => {
  const $0 = normalizeCaseRows(denormalizedRows);
  if ($0.length > 0) {
    const v1 = Data$dArray$dNonEmpty.uncons($0);
    const $1 = Data$dFoldableWithIndex.foldableWithIndexArray.foldlWithIndex(idx => acc => p => {
      if (!eqPatternCase.eq(p.pattern.patternCase)(PatWild)) { return Data$dArray.snoc(acc)(Data$dTuple.$Tuple(idx, p)); }
      return acc;
    })([])(v1.head.patterns);
    if ($1.length > 0) { return buildCasePattern(chooseNextPattern($1)(v1.tail))($0); }
    return buildCaseLeaf(v1.head)(v1.tail);
  }
  return patternFail;
};
const buildCasePattern = chosenColumn => rows => {
  const letBindSubterm = v => nextCb => idents => makeLet(Data$dMaybe.Nothing)(make(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
    "Accessor",
    make(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Local", Data$dMaybe.Nothing, chosenColumn.column)),
    v.accessor
  )))(tmp => nextCb(Data$dArray.snoc(idents)(tmp)));
  const $0 = decompose(chosenColumn)(rows).rowsWithMatch;
  const expandSubterms = Data$dFoldable.foldrArray(letBindSubterm)(idents => buildCaseTreeFromRows((() => {
    const inlineWildSubterms = Data$dFunctor.arrayMap(column => ({column, pattern: {vars: Data$dMap$dInternal.Leaf, patternCase: PatWild, subterms: []}}))(idents);
    return Data$dFunctor.arrayMap(v => (
      {
        guardFn: v.guardFn,
        vars: append2(v.vars)(toCaseRowVars(v.match)),
        patterns: [
          ...v.nonMatchesBefore,
          ...v.match.pattern.patternCase.tag === "PatWild"
            ? inlineWildSubterms
            : Data$dArray.zipWithImpl(column => v$1 => ({column, pattern: v$1.pattern}), idents, v.match.pattern.subterms),
          ...v.nonMatchesAfter
        ]
      }
    ))($0);
  })()))(chosenColumn.pattern.subterms)([]);
  const buildCaseBranch = guardExpr => {
    const v = decompose(chosenColumn)(rows);
    const $1 = v.rowsWithMatch;
    return makeGuard(chosenColumn.column)(guardExpr)(Data$dFoldable.foldrArray(letBindSubterm)(idents => buildCaseTreeFromRows((() => {
      const inlineWildSubterms = Data$dFunctor.arrayMap(column => ({column, pattern: {vars: Data$dMap$dInternal.Leaf, patternCase: PatWild, subterms: []}}))(idents);
      return Data$dFunctor.arrayMap(v$1 => (
        {
          guardFn: v$1.guardFn,
          vars: append2(v$1.vars)(toCaseRowVars(v$1.match)),
          patterns: [
            ...v$1.nonMatchesBefore,
            ...v$1.match.pattern.patternCase.tag === "PatWild"
              ? inlineWildSubterms
              : Data$dArray.zipWithImpl(column => v$2 => ({column, pattern: v$2.pattern}), idents, v$1.match.pattern.subterms),
            ...v$1.nonMatchesAfter
          ]
        }
      ))($1);
    })()))(chosenColumn.pattern.subterms)([]))(buildCaseTreeFromRows(v.rowsNoMatch));
  };
  if (chosenColumn.pattern.patternCase.tag === "PatWild") { return Partial._crashWith("Impossible: chosen column cannot be wild pattern"); }
  if (chosenColumn.pattern.patternCase.tag === "PatRecord") { return expandSubterms; }
  if (chosenColumn.pattern.patternCase.tag === "PatProduct") { return expandSubterms; }
  if (chosenColumn.pattern.patternCase.tag === "PatSum") {
    const $1 = chosenColumn.pattern.patternCase._2;
    return buildCaseBranch(lhs => PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
      "PrimOp",
      PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op1", PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator1("OpIsTag", $1), lhs)
    ));
  }
  if (chosenColumn.pattern.patternCase.tag === "PatArray") { return buildCaseBranch(guardArrayLength(chosenColumn.pattern.patternCase._1)); }
  if (chosenColumn.pattern.patternCase.tag === "PatInt") { return buildCaseBranch(guardInt(chosenColumn.pattern.patternCase._1)); }
  if (chosenColumn.pattern.patternCase.tag === "PatNumber") { return buildCaseBranch(guardNumber(chosenColumn.pattern.patternCase._1)); }
  if (chosenColumn.pattern.patternCase.tag === "PatString") { return buildCaseBranch(guardString(chosenColumn.pattern.patternCase._1)); }
  if (chosenColumn.pattern.patternCase.tag === "PatChar") { return buildCaseBranch(guardChar(chosenColumn.pattern.patternCase._1)); }
  if (chosenColumn.pattern.patternCase.tag === "PatBoolean") { return buildCaseBranch(guardBoolean(chosenColumn.pattern.patternCase._1)); }
  $runtime.fail();
};
const buildCaseLeaf = row0 => tailRows => {
  const orderedArgs = toUnfoldable1(append2(row0.vars)(foldMap4(toCaseRowVars)(row0.patterns)));
  if (row0.guardFn.tag === "UnconditionalFn") {
    return make(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
      "UncurriedApp",
      make(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Local", Data$dMaybe.Nothing, row0.guardFn._1)),
      Data$dFunctor.arrayMap(v => make(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Local", Data$dMaybe.$Maybe("Just", v._1), v._2)))(orderedArgs)
    ));
  }
  if (row0.guardFn.tag === "GuardedFn") {
    const $0 = row0.guardFn._1;
    return Data$dFoldable.foldrArray(v => {
      const $1 = v._1;
      const $2 = v._2;
      return cb => args => makeLet(Data$dMaybe.$Maybe("Just", $1))(make(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Local", Data$dMaybe.Nothing, $2)))(tmp => cb(Data$dArray.snoc(args)(Data$dTuple.$Tuple(
        $1,
        tmp
      ))));
    })(args => {
      const $1 = $$for($0)(v => {
        const $1 = toBackendExpr(v._1);
        const $2 = make(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
          "UncurriedApp",
          make(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Local", Data$dMaybe.Nothing, v._2)),
          Data$dFunctor.arrayMap(v$1 => make(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Local", Data$dMaybe.$Maybe("Just", v$1._1), v$1._2)))(args)
        ));
        return x => PureScript$dBackend$dOptimizer$dSyntax.$Pair($1(x), $2(x));
      });
      return x => PureScript$dBackend$dOptimizer$dSemantics.build(getCtx(x))(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
        "Branch",
        $1(x),
        buildCaseTreeFromRows(tailRows)(x)
      ));
    })(orderedArgs)([]);
  }
  $runtime.fail();
};
const toTopLevelBackendBinding = group => env => v => {
  const qualifiedIdent = PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", env.currentModule), v._2);
  const backendExpr = toBackendExpr(v._3)(env);
  const v1 = PureScript$dBackend$dOptimizer$dSemantics.optimize(member(qualifiedIdent)(env.traceIdents))(getCtx(env))({
    currentModule: env.currentModule,
    evalExternRef: makeExternEvalRef(env),
    evalExternSpine: makeExternEvalSpine(env),
    locals: [],
    directives: env.directives
  })(qualifiedIdent)(env.rewriteLimit)(backendExpr);
  const v2 = toExternImpl(env)(group)(v1._2);
  return {
    accum: {
      ...env,
      implementations: Data$dMap$dInternal.insert(ordQualified)(qualifiedIdent)(v2._1)(env.implementations),
      moduleImplementations: Data$dMap$dInternal.insert(ordQualified)(qualifiedIdent)(v2._1)(env.moduleImplementations),
      optimizationSteps: (() => {
        const $0 = Data$dTuple.Tuple(qualifiedIdent);
        if (v1._1.length > 0) { return Data$dArray.snoc(env.optimizationSteps)($0(v1._1)); }
        return env.optimizationSteps;
      })(),
      directives: (() => {
        const v4 = inferTransitiveDirective(env.directives)(v2._1._2)(backendExpr)(v._3);
        if (v4.tag === "Just") {
          const $0 = v4._1;
          return alter(v5 => {
            if (v5.tag === "Just") {
              return Data$dMaybe.$Maybe(
                "Just",
                Data$dMap$dInternal.unsafeUnionWith(PureScript$dBackend$dOptimizer$dSemantics.ordInlineAccessor.compare, Data$dFunction.const, v5._1, $0)
              );
            }
            if (v5.tag === "Nothing") { return Data$dMaybe.$Maybe("Just", $0); }
            $runtime.fail();
          })(PureScript$dBackend$dOptimizer$dSemantics.$EvalRef(
            "EvalExtern",
            PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", env.currentModule), v._2)
          ))(env.directives);
        }
        if (v4.tag === "Nothing") { return env.directives; }
        $runtime.fail();
      })()
    },
    value: Data$dTuple.$Tuple(v._2, Data$dTuple.$Tuple(v2._1._1.deps, v2._2))
  };
};
const toBackendTopLevelBindingGroup = env => v => {
  if (v.tag === "Rec") {
    const $0 = mapAccumL(toTopLevelBackendBinding(Data$dFunctor.arrayMap(v1 => PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(
      Data$dMaybe.$Maybe("Just", env.currentModule),
      v1._2
    ))(v._1)))(env)(v._1);
    return {...$0, value: {recursive: true, bindings: $0.value}};
  }
  if (v.tag === "NonRec") {
    const $0 = mapAccumL(toTopLevelBackendBinding([]))(env)([v._1]);
    return {...$0, value: {recursive: false, bindings: $0.value}};
  }
  $runtime.fail();
};
const toBackendTopLevelBindingGroups = binds => env => {
  const result = mapAccumL(toBackendTopLevelBindingGroup)(env)(binds);
  return {
    ...result,
    value: Data$dFunctor.arrayMap(as => (
      {
        recursive: (() => {
          if (0 < as.length) { return as[0].recursive; }
          $runtime.fail();
        })(),
        bindings: Control$dBind.arrayBind(as)(v1 => v1.bindings)
      }
    ))(Data$dArray.groupBy(x => y => !x.recursive && !y.recursive)(result.value))
  };
};
const toBackendModule = v => env => {
  const localExports = fromFoldable2(v.exports);
  const isBindingUsed = deps => v1 => member1(v1._1)(localExports) || member(PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", v.name), v1._1))(deps);
  const directives = PureScript$dBackend$dOptimizer$dDirectives.parseDirectiveHeader(v.name)(v.comments);
  const dataTypes = fromFoldable3(Data$dFunctor.arrayMap(group => Data$dTuple.$Tuple(
    (() => {
      if (0 < group.length) { return group[0]._1; }
      $runtime.fail();
    })(),
    {
      constructors: fromFoldable4(Data$dFunctorWithIndex.mapWithIndexArray(tag => v1 => Data$dTuple.$Tuple(v1._2._1, {fields: v1._2._2, tag}))(group)),
      size: maximum(Data$dFunctor.arrayMap(x => x._2._2.length)(group))
    }
  ))(Data$dArray.groupAllBy(x => y => Data$dOrd.ordString.compare(x._1)(y._1))(Control$dBind.arrayBind(Control$dBind.arrayBind(v.decls)(v1 => {
    if (v1.tag === "Rec") { return v1._1; }
    if (v1.tag === "NonRec") { return [v1._1]; }
    $runtime.fail();
  }))(v1 => {
    if (v1._3.tag === "ExprConstructor") { return [Data$dTuple.$Tuple(v1._3._2, Data$dTuple.$Tuple(v1._3._3, v1._3._4))]; }
    return [];
  }))));
  const moduleBindings = toBackendTopLevelBindingGroups(v.decls)({
    ...env,
    dataTypes,
    directives: (() => {
      const go = (z$p, m$p) => {
        if (m$p.tag === "Leaf") { return z$p; }
        if (m$p.tag === "Node") {
          return go(
            (() => {
              const $0 = m$p._4;
              return alter(v2 => {
                if (v2.tag === "Nothing") { return Data$dMaybe.$Maybe("Just", $0); }
                if (v2.tag === "Just") { return Data$dMaybe.$Maybe("Just", v2._1); }
                $runtime.fail();
              })(m$p._3)(go(z$p, m$p._5));
            })(),
            m$p._6
          );
        }
        $runtime.fail();
      };
      return go(
        Data$dMap$dInternal.unsafeUnionWith(PureScript$dBackend$dOptimizer$dSemantics.ordEvalRef.compare, Data$dFunction.const, directives.locals, env.directives),
        directives.exports
      );
    })(),
    moduleImplementations: Data$dMap$dInternal.Leaf
  });
  const usedBindings = mapAccumR(deps => group => {
    if (group.recursive) {
      if (Data$dArray.anyImpl(isBindingUsed(deps), group.bindings)) {
        return {
          accum: Data$dMap$dInternal.unsafeUnionWith(ordQualified.compare, Data$dFunction.const, foldMap5(x => x._2._1)(group.bindings), deps),
          value: {...group, bindings: Data$dArray.mapMaybe(x => x)(Data$dFunctor.arrayMap(x => Data$dMaybe.$Maybe("Just", Data$dTuple.$Tuple(x._1, x._2._2)))(group.bindings))}
        };
      }
      return {accum: deps, value: {...group, bindings: Data$dArray.mapMaybe(x => x)([])}};
    }
    return {
      accum: mapAccumR(deps$p => v2 => {
        if (isBindingUsed(deps$p)(v2)) {
          return {
            accum: Data$dMap$dInternal.unsafeUnionWith(ordQualified.compare, Data$dFunction.const, v2._2._1, deps$p),
            value: Data$dMaybe.$Maybe("Just", Data$dTuple.$Tuple(v2._1, v2._2._2))
          };
        }
        return {accum: deps$p, value: Data$dMaybe.Nothing};
      })(deps)(group.bindings).accum,
      value: {
        ...group,
        bindings: Data$dArray.mapMaybe(x => x)(mapAccumR(deps$p => v2 => {
          if (isBindingUsed(deps$p)(v2)) {
            return {
              accum: Data$dMap$dInternal.unsafeUnionWith(ordQualified.compare, Data$dFunction.const, v2._2._1, deps$p),
              value: Data$dMaybe.$Maybe("Just", Data$dTuple.$Tuple(v2._1, v2._2._2))
            };
          }
          return {accum: deps$p, value: Data$dMaybe.Nothing};
        })(deps)(group.bindings).value)
      }
    };
  })(Data$dMap$dInternal.Leaf)(moduleBindings.value);
  return Data$dTuple.$Tuple(
    moduleBindings.accum.optimizationSteps,
    {
      name: v.name,
      comments: v.comments,
      imports: Data$dSet.mapMaybe(Data$dOrd.ordString)(qi => {
        if (qi._1.tag === "Just") {
          if (qi._1._1 !== v.name && qi._1._1 !== "Prim") { return Data$dMaybe.$Maybe("Just", qi._1._1); }
          return Data$dMaybe.Nothing;
        }
        if (qi._1.tag === "Nothing") { return Data$dMaybe.Nothing; }
        $runtime.fail();
      })(usedBindings.accum),
      dataTypes: (() => {
        const $0 = Data$dArray.any(isBindingUsed(usedBindings.accum));
        return Data$dMap$dInternal.filterWithKey(Data$dOrd.ordString)(v$1 => x => $0(toUnfoldable1(x.constructors)))(dataTypes);
      })(),
      bindings: usedBindings.value,
      exports: localExports,
      reExports: fromFoldable5(v.reExports),
      implementations: moduleBindings.accum.moduleImplementations,
      directives: directives.exports,
      foreign: fromFoldable2(v.foreign)
    }
  );
};
export {
  $CaseRowGuardedExpr,
  $PatternCase,
  GuardedFn,
  PatArray,
  PatBoolean,
  PatChar,
  PatInt,
  PatNumber,
  PatProduct,
  PatRecord,
  PatString,
  PatSum,
  PatWild,
  Pattern,
  UnconditionalFn,
  alter,
  analyze,
  analyzeEffectBlock,
  append2,
  binderToPattern,
  buildCaseLeaf,
  buildCasePattern,
  buildCaseTreeFromRows,
  buildM,
  chooseNextPattern,
  compare,
  compare1,
  currentLevel,
  decompose,
  eq,
  eqPatternCase,
  foldMap,
  foldMap2,
  foldMap3,
  foldMap4,
  foldMap5,
  foldMapWithIndex,
  $$for as for,
  forWithIndex,
  fromExternImpl,
  fromFoldable1,
  fromFoldable2,
  fromFoldable3,
  fromFoldable4,
  fromFoldable5,
  getCtx,
  guardArrayLength,
  guardBoolean,
  guardChar,
  guardInt,
  guardNumber,
  guardString,
  guardTag,
  inferTransitiveDirective,
  intro,
  levelUp,
  lookup,
  lookup1,
  lookup2,
  lookup3,
  lookup4,
  make,
  makeExternEvalRef,
  makeExternEvalSpine,
  makeGuard,
  makeLet,
  makeUncurriedAbs,
  mapAccumL,
  mapAccumR,
  maximum,
  member,
  member1,
  monoidSemigroupMap,
  monoidSet,
  newtypePattern_,
  normalizeCaseRows,
  ordPatternCase,
  ordQualified,
  patternFail,
  patternPatCase,
  patternSubterms,
  patternVars,
  toBackendBinding,
  toBackendExpr,
  toBackendModule,
  toBackendTopLevelBindingGroup,
  toBackendTopLevelBindingGroups,
  toCaseRowVars,
  toExternImpl,
  toTopLevelBackendBinding,
  toUnfoldable,
  toUnfoldable1,
  topEnv,
  traverse,
  traverse1,
  traverse3,
  zipWithA
};
