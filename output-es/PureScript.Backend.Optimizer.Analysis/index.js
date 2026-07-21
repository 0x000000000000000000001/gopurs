import * as $runtime from "../runtime.js";
import * as Data$dArray from "../Data.Array/index.js";
import * as Data$dFoldable from "../Data.Foldable/index.js";
import * as Data$dFunction from "../Data.Function/index.js";
import * as Data$dMap$dInternal from "../Data.Map.Internal/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dOrd from "../Data.Ord/index.js";
import * as Data$dOrdering from "../Data.Ordering/index.js";
import * as Data$dString$dCodeUnits from "../Data.String.CodeUnits/index.js";
import * as PureScript$dBackend$dOptimizer$dCoreFn from "../PureScript.Backend.Optimizer.CoreFn/index.js";
import * as PureScript$dBackend$dOptimizer$dSyntax from "../PureScript.Backend.Optimizer.Syntax/index.js";
const $Capture = tag => tag;
const $Complexity = tag => tag;
const $ResultTerm = tag => tag;
const ordQualified = /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.ordQualified(Data$dOrd.ordString);
const pop = /* #__PURE__ */ Data$dMap$dInternal.pop(Data$dOrd.ordInt);
const KnownNeutral = /* #__PURE__ */ $ResultTerm("KnownNeutral");
const Unknown = /* #__PURE__ */ $ResultTerm("Unknown");
const Trivial = /* #__PURE__ */ $Complexity("Trivial");
const Deref = /* #__PURE__ */ $Complexity("Deref");
const KnownSize = /* #__PURE__ */ $Complexity("KnownSize");
const NonTrivial = /* #__PURE__ */ $Complexity("NonTrivial");
const CaptureNone = /* #__PURE__ */ $Capture("CaptureNone");
const CaptureBranch = /* #__PURE__ */ $Capture("CaptureBranch");
const CaptureClosure = /* #__PURE__ */ $Capture("CaptureClosure");
const Usage = x => x;
const BackendAnalysis = x => x;
const semigroupResultTerm = {
  append: v => v1 => {
    if (v === "Unknown") { return Unknown; }
    if (v1 === "Unknown") { return Unknown; }
    return KnownNeutral;
  }
};
const newtypeUsage_ = {Coercible0: () => {}};
const newtypeBackendAnalysis_ = {Coercible0: () => {}};
const monoidResultTerm = {mempty: KnownNeutral, Semigroup0: () => semigroupResultTerm};
const foldMap1 = /* #__PURE__ */ (() => Data$dFoldable.foldableArray.foldMap(monoidResultTerm))();
const eqResultTerm = {
  eq: x => y => {
    if (x === "KnownNeutral") { return y === "KnownNeutral"; }
    return x === "Unknown" && y === "Unknown";
  }
};
const eqComplexity = {
  eq: x => y => {
    if (x === "Trivial") { return y === "Trivial"; }
    if (x === "Deref") { return y === "Deref"; }
    if (x === "KnownSize") { return y === "KnownSize"; }
    return x === "NonTrivial" && y === "NonTrivial";
  }
};
const ordComplexity = {
  compare: x => y => {
    if (x === "Trivial") {
      if (y === "Trivial") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "Trivial") { return Data$dOrdering.GT; }
    if (x === "Deref") {
      if (y === "Deref") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "Deref") { return Data$dOrdering.GT; }
    if (x === "KnownSize") {
      if (y === "KnownSize") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "KnownSize") { return Data$dOrdering.GT; }
    if (x === "NonTrivial" && y === "NonTrivial") { return Data$dOrdering.EQ; }
    $runtime.fail();
  },
  Eq0: () => eqComplexity
};
const semigroupComplexity = {
  append: x => y => {
    if (x === "Trivial") {
      if (y === "Trivial") { return x; }
      return y;
    }
    if (y === "Trivial") { return x; }
    if (x === "Deref") {
      if (y === "Deref") { return x; }
      return y;
    }
    if (y === "Deref") { return x; }
    if (x === "KnownSize") {
      if (y === "KnownSize") { return x; }
      return y;
    }
    if (y === "KnownSize") { return x; }
    if (x === "NonTrivial" && y === "NonTrivial") { return x; }
    $runtime.fail();
  }
};
const monoidComplexity = {mempty: Trivial, Semigroup0: () => semigroupComplexity};
const eqCapture = {
  eq: x => y => {
    if (x === "CaptureNone") { return y === "CaptureNone"; }
    if (x === "CaptureBranch") { return y === "CaptureBranch"; }
    return x === "CaptureClosure" && y === "CaptureClosure";
  }
};
const ordCapture = {
  compare: x => y => {
    if (x === "CaptureNone") {
      if (y === "CaptureNone") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "CaptureNone") { return Data$dOrdering.GT; }
    if (x === "CaptureBranch") {
      if (y === "CaptureBranch") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "CaptureBranch") { return Data$dOrdering.GT; }
    if (x === "CaptureClosure" && y === "CaptureClosure") { return Data$dOrdering.EQ; }
    $runtime.fail();
  },
  Eq0: () => eqCapture
};
const semigroupCapture = {
  append: x => y => {
    if (x === "CaptureNone") {
      if (y === "CaptureNone") { return x; }
      return y;
    }
    if (y === "CaptureNone") { return x; }
    if (x === "CaptureBranch") {
      if (y === "CaptureBranch") { return x; }
      return y;
    }
    if (y === "CaptureBranch") { return x; }
    if (x === "CaptureClosure" && y === "CaptureClosure") { return x; }
    $runtime.fail();
  }
};
const monoidCapture = {mempty: CaptureNone, Semigroup0: () => semigroupCapture};
const semigroupUsage = {
  append: v => v1 => (
    {
      total: v.total + v1.total | 0,
      captured: (() => {
        if (v.captured === "CaptureNone") {
          if (v1.captured === "CaptureNone") { return v.captured; }
          return v1.captured;
        }
        if (v1.captured === "CaptureNone") { return v.captured; }
        if (v.captured === "CaptureBranch") {
          if (v1.captured === "CaptureBranch") { return v.captured; }
          return v1.captured;
        }
        if (v1.captured === "CaptureBranch") { return v.captured; }
        if (v.captured === "CaptureClosure" && v1.captured === "CaptureClosure") { return v.captured; }
        $runtime.fail();
      })(),
      arities: Data$dMap$dInternal.unsafeUnionWith(Data$dOrd.ordInt.compare, Data$dFunction.const, v.arities, v1.arities),
      call: v.call + v1.call | 0,
      access: v.access + v1.access | 0,
      case: v.case + v1.case | 0,
      update: v.update + v1.update | 0
    }
  )
};
const monoidUsage = {mempty: {total: 0, captured: CaptureNone, arities: Data$dMap$dInternal.Leaf, call: 0, access: 0, case: 0, update: 0}, Semigroup0: () => semigroupUsage};
const semigroupBackendAnalysis = {
  append: v => v1 => (
    {
      usages: Data$dMap$dInternal.unsafeUnionWith(Data$dOrd.ordInt.compare, semigroupUsage.append, v.usages, v1.usages),
      size: v.size + v1.size | 0,
      complexity: (() => {
        if (v.complexity === "Trivial") {
          if (v1.complexity === "Trivial") { return v.complexity; }
          return v1.complexity;
        }
        if (v1.complexity === "Trivial") { return v.complexity; }
        if (v.complexity === "Deref") {
          if (v1.complexity === "Deref") { return v.complexity; }
          return v1.complexity;
        }
        if (v1.complexity === "Deref") { return v.complexity; }
        if (v.complexity === "KnownSize") {
          if (v1.complexity === "KnownSize") { return v.complexity; }
          return v1.complexity;
        }
        if (v1.complexity === "KnownSize") { return v.complexity; }
        if (v.complexity === "NonTrivial" && v1.complexity === "NonTrivial") { return v.complexity; }
        $runtime.fail();
      })(),
      args: [],
      rewrite: v.rewrite || v1.rewrite,
      deps: Data$dMap$dInternal.unsafeUnionWith(ordQualified.compare, Data$dFunction.const, v.deps, v1.deps),
      result: (() => {
        if (v.result === "Unknown") { return Unknown; }
        if (v1.result === "Unknown") { return Unknown; }
        return KnownNeutral;
      })(),
      externs: v.externs || v1.externs
    }
  )
};
const monoidBackendAnalysis = {
  mempty: {usages: Data$dMap$dInternal.Leaf, size: 0, complexity: Trivial, args: [], rewrite: false, deps: Data$dMap$dInternal.Leaf, result: KnownNeutral, externs: false},
  Semigroup0: () => semigroupBackendAnalysis
};
const foldMap2 = /* #__PURE__ */ (() => PureScript$dBackend$dOptimizer$dSyntax.foldableBackendSyntax.foldMap(monoidBackendAnalysis))();
const foldMap3 = /* #__PURE__ */ (() => Data$dFoldable.foldableArray.foldMap(monoidBackendAnalysis))();
const foldMap4 = /* #__PURE__ */ (() => Data$dFoldable.foldableArray.foldMap(monoidBackendAnalysis))();
const foldMap6 = f => v => semigroupBackendAnalysis.append(f(v._1))(f(v._2));
const withRewrite = v => ({...v, rewrite: true});
const withResult = r => v => ({...v, result: r});
const withArgs = args => v => ({...v, args});
const usedDep = dep => v => ({...v, deps: Data$dMap$dInternal.insert(ordQualified)(dep)()(v.deps)});
const used = level => (
  {
    ...monoidBackendAnalysis.mempty,
    usages: Data$dMap$dInternal.$$$Map(
      "Node",
      1,
      1,
      level,
      {total: 1, captured: CaptureNone, arities: Data$dMap$dInternal.Leaf, call: 0, access: 0, case: 0, update: 0},
      Data$dMap$dInternal.Leaf,
      Data$dMap$dInternal.Leaf
    )
  }
);
const updated = level => v => ({...v, usages: Data$dMap$dInternal.update(Data$dOrd.ordInt)(x => Data$dMaybe.$Maybe("Just", {...x, update: x.update + 1 | 0}))(level)(v.usages)});
const externs = v => ({...v, externs: true});
const complex = complexity => v => (
  {
    ...v,
    complexity: (() => {
      if (v.complexity === "Trivial") {
        if (complexity === "Trivial") { return v.complexity; }
        return complexity;
      }
      if (complexity === "Trivial") { return v.complexity; }
      if (v.complexity === "Deref") {
        if (complexity === "Deref") { return v.complexity; }
        return complexity;
      }
      if (complexity === "Deref") { return v.complexity; }
      if (v.complexity === "KnownSize") {
        if (complexity === "KnownSize") { return v.complexity; }
        return complexity;
      }
      if (complexity === "KnownSize") { return v.complexity; }
      if (v.complexity === "NonTrivial" && complexity === "NonTrivial") { return v.complexity; }
      $runtime.fail();
    })()
  }
);
const cased = level => v => ({...v, usages: Data$dMap$dInternal.update(Data$dOrd.ordInt)(x => Data$dMaybe.$Maybe("Just", {...x, case: x.case + 1 | 0}))(level)(v.usages)});
const capture = cap => v => (
  {
    ...v,
    usages: (() => {
      const go = v$1 => {
        if (v$1.tag === "Leaf") { return Data$dMap$dInternal.Leaf; }
        if (v$1.tag === "Node") { return Data$dMap$dInternal.$$$Map("Node", v$1._1, v$1._2, v$1._3, {...v$1._4, captured: cap}, go(v$1._5), go(v$1._6)); }
        $runtime.fail();
      };
      return go(v.usages);
    })()
  }
);
const callArity = lvl => arity => v => (
  {
    ...v,
    usages: Data$dMap$dInternal.update(Data$dOrd.ordInt)(x => Data$dMaybe.$Maybe(
      "Just",
      {...x, arities: Data$dMap$dInternal.insert(Data$dOrd.ordInt)(arity)()(x.arities), call: x.call + 1 | 0}
    ))(lvl)(v.usages)
  }
);
const bump = v => ({...v, size: v.size + 1 | 0});
const boundArg = level => v => {
  const v1 = pop(level)(v.usages);
  if (v1.tag === "Nothing") { return {...v, args: [monoidUsage.mempty, ...v.args]}; }
  if (v1.tag === "Just") { return {...v, usages: v1._1._2, args: [v1._1._1, ...v.args]}; }
  $runtime.fail();
};
const bound = level => v => ({...v, usages: Data$dMap$dInternal.delete(Data$dOrd.ordInt)(level)(v.usages)});
const analysisOf = dict => dict.analysisOf;
const analyzeDefault = dictHasAnalysis => {
  const $0 = foldMap2(dictHasAnalysis.analysisOf);
  return x => {
    const $1 = $0(x);
    return {...$1, size: $1.size + 1 | 0};
  };
};
const resultOf = dictHasAnalysis => x => dictHasAnalysis.analysisOf(x).result;
const accessed = level => v => ({...v, usages: Data$dMap$dInternal.update(Data$dOrd.ordInt)(x => Data$dMaybe.$Maybe("Just", {...x, access: x.access + 1 | 0}))(level)(v.usages)});
const analyze = dictHasAnalysis => {
  const analysisOf1 = dictHasAnalysis.analysisOf;
  const analyzeDefault1 = analyzeDefault(dictHasAnalysis);
  return dictHasSyntax => externAnalysis => expr => {
    if (expr.tag === "Var") {
      const analysis = {...monoidBackendAnalysis.mempty, deps: Data$dMap$dInternal.insert(ordQualified)(expr._1)()(monoidBackendAnalysis.mempty.deps), externs: true, size: 1};
      const v = externAnalysis(expr._1)(Data$dMaybe.Nothing);
      if (v.tag === "Just") { return {...analysis, args: v._1.args}; }
      if (v.tag === "Nothing") { return analysis; }
      $runtime.fail();
    }
    if (expr.tag === "Local") {
      const $0 = used(expr._2);
      return {...$0, size: $0.size + 1 | 0};
    }
    if (expr.tag === "Let") {
      const $0 = semigroupBackendAnalysis.append(analysisOf1(expr._3))((() => {
        const $0 = analysisOf1(expr._4);
        return {...$0, usages: Data$dMap$dInternal.delete(Data$dOrd.ordInt)(expr._2)($0.usages)};
      })());
      return {
        ...$0,
        complexity: (() => {
          if ($0.complexity === "Trivial") { return NonTrivial; }
          if ($0.complexity === "Deref") { return NonTrivial; }
          if ($0.complexity === "KnownSize") { return NonTrivial; }
          if ($0.complexity === "NonTrivial") { return $0.complexity; }
          $runtime.fail();
        })(),
        result: dictHasAnalysis.analysisOf(expr._4).result,
        size: $0.size + 1 | 0
      };
    }
    if (expr.tag === "LetRec") {
      const $0 = semigroupBackendAnalysis.append(foldMap3(x => analysisOf1(x._2))(expr._2))(analysisOf1(expr._3));
      return {
        ...$0,
        complexity: (() => {
          if ($0.complexity === "Trivial") { return NonTrivial; }
          if ($0.complexity === "Deref") { return NonTrivial; }
          if ($0.complexity === "KnownSize") { return NonTrivial; }
          if ($0.complexity === "NonTrivial") { return $0.complexity; }
          $runtime.fail();
        })(),
        result: dictHasAnalysis.analysisOf(expr._3).result,
        size: $0.size + 1 | 0,
        usages: Data$dMap$dInternal.delete(Data$dOrd.ordInt)(expr._1)($0.usages)
      };
    }
    if (expr.tag === "EffectBind") {
      const $0 = semigroupBackendAnalysis.append(analysisOf1(expr._3))((() => {
        const $0 = analysisOf1(expr._4);
        return {...$0, usages: Data$dMap$dInternal.delete(Data$dOrd.ordInt)(expr._2)($0.usages)};
      })());
      const go = v => {
        if (v.tag === "Leaf") { return Data$dMap$dInternal.Leaf; }
        if (v.tag === "Node") { return Data$dMap$dInternal.$$$Map("Node", v._1, v._2, v._3, {...v._4, captured: CaptureClosure}, go(v._5), go(v._6)); }
        $runtime.fail();
      };
      return {
        ...$0,
        complexity: (() => {
          if ($0.complexity === "Trivial") { return NonTrivial; }
          if ($0.complexity === "Deref") { return NonTrivial; }
          if ($0.complexity === "KnownSize") { return NonTrivial; }
          if ($0.complexity === "NonTrivial") { return $0.complexity; }
          $runtime.fail();
        })(),
        result: Unknown,
        size: $0.size + 1 | 0,
        usages: go($0.usages)
      };
    }
    if (expr.tag === "EffectPure") {
      const $0 = analysisOf1(expr._1);
      return {
        ...$0,
        result: Unknown,
        size: $0.size + 1 | 0,
        usages: (() => {
          const go = v => {
            if (v.tag === "Leaf") { return Data$dMap$dInternal.Leaf; }
            if (v.tag === "Node") { return Data$dMap$dInternal.$$$Map("Node", v._1, v._2, v._3, {...v._4, captured: CaptureClosure}, go(v._5), go(v._6)); }
            $runtime.fail();
          };
          return go($0.usages);
        })()
      };
    }
    if (expr.tag === "EffectDefer") {
      const $0 = analysisOf1(expr._1);
      return {
        ...$0,
        result: Unknown,
        size: $0.size + 1 | 0,
        usages: (() => {
          const go = v => {
            if (v.tag === "Leaf") { return Data$dMap$dInternal.Leaf; }
            if (v.tag === "Node") { return Data$dMap$dInternal.$$$Map("Node", v._1, v._2, v._3, {...v._4, captured: CaptureClosure}, go(v._5), go(v._6)); }
            $runtime.fail();
          };
          return go($0.usages);
        })()
      };
    }
    if (expr.tag === "Abs") {
      const $0 = Data$dFoldable.foldrArray(x => boundArg(x._2))(analyzeDefault1(expr))(expr._1);
      const go = v => {
        if (v.tag === "Leaf") { return Data$dMap$dInternal.Leaf; }
        if (v.tag === "Node") { return Data$dMap$dInternal.$$$Map("Node", v._1, v._2, v._3, {...v._4, captured: CaptureClosure}, go(v._5), go(v._6)); }
        $runtime.fail();
      };
      return {
        ...$0,
        complexity: (() => {
          if ($0.complexity === "Trivial") { return KnownSize; }
          if ($0.complexity === "Deref") { return KnownSize; }
          return $0.complexity;
        })(),
        result: KnownNeutral,
        usages: go($0.usages)
      };
    }
    if (expr.tag === "UncurriedAbs") {
      const $0 = Data$dFoldable.foldrArray(x => boundArg(x._2))(analyzeDefault1(expr))(expr._1);
      const go = v => {
        if (v.tag === "Leaf") { return Data$dMap$dInternal.Leaf; }
        if (v.tag === "Node") { return Data$dMap$dInternal.$$$Map("Node", v._1, v._2, v._3, {...v._4, captured: CaptureClosure}, go(v._5), go(v._6)); }
        $runtime.fail();
      };
      return {
        ...$0,
        complexity: (() => {
          if ($0.complexity === "Trivial") { return KnownSize; }
          if ($0.complexity === "Deref") { return KnownSize; }
          return $0.complexity;
        })(),
        result: KnownNeutral,
        usages: go($0.usages)
      };
    }
    if (expr.tag === "UncurriedApp") {
      const $0 = analyzeDefault1(expr);
      const analysis = {
        ...$0,
        complexity: (() => {
          if ($0.complexity === "Trivial") { return NonTrivial; }
          if ($0.complexity === "Deref") { return NonTrivial; }
          if ($0.complexity === "KnownSize") { return NonTrivial; }
          if ($0.complexity === "NonTrivial") { return $0.complexity; }
          $runtime.fail();
        })(),
        result: Unknown
      };
      const v = dictHasSyntax.syntaxOf(expr._1);
      if (v.tag === "Just" && v._1.tag === "Local") { return callArity(v._1._2)(expr._2.length)(analysis); }
      return analysis;
    }
    if (expr.tag === "UncurriedEffectAbs") {
      const $0 = Data$dFoldable.foldrArray(x => boundArg(x._2))(analyzeDefault1(expr))(expr._1);
      const go = v => {
        if (v.tag === "Leaf") { return Data$dMap$dInternal.Leaf; }
        if (v.tag === "Node") { return Data$dMap$dInternal.$$$Map("Node", v._1, v._2, v._3, {...v._4, captured: CaptureClosure}, go(v._5), go(v._6)); }
        $runtime.fail();
      };
      return {
        ...$0,
        complexity: (() => {
          if ($0.complexity === "Trivial") { return KnownSize; }
          if ($0.complexity === "Deref") { return KnownSize; }
          return $0.complexity;
        })(),
        result: KnownNeutral,
        usages: go($0.usages)
      };
    }
    if (expr.tag === "UncurriedEffectApp") {
      const $0 = analyzeDefault1(expr);
      const go = v => {
        if (v.tag === "Leaf") { return Data$dMap$dInternal.Leaf; }
        if (v.tag === "Node") { return Data$dMap$dInternal.$$$Map("Node", v._1, v._2, v._3, {...v._4, captured: CaptureClosure}, go(v._5), go(v._6)); }
        $runtime.fail();
      };
      const analysis = {
        ...$0,
        complexity: (() => {
          if ($0.complexity === "Trivial") { return NonTrivial; }
          if ($0.complexity === "Deref") { return NonTrivial; }
          if ($0.complexity === "KnownSize") { return NonTrivial; }
          if ($0.complexity === "NonTrivial") { return $0.complexity; }
          $runtime.fail();
        })(),
        result: Unknown,
        usages: go($0.usages)
      };
      const v = dictHasSyntax.syntaxOf(expr._1);
      if (v.tag === "Just" && v._1.tag === "Local") { return callArity(v._1._2)(expr._2.length)(analysis); }
      return analysis;
    }
    if (expr.tag === "App") {
      const $0 = analysisOf1(expr._1).args;
      const $1 = expr._2.length;
      const remainingArgs = $1 < 1 ? $0 : Data$dArray.sliceImpl($1, $0.length, $0);
      const analysis = (() => {
        if (remainingArgs.length === 0) {
          const $2 = analyzeDefault1(expr);
          return {
            ...$2,
            complexity: (() => {
              if ($2.complexity === "Trivial") { return NonTrivial; }
              if ($2.complexity === "Deref") { return NonTrivial; }
              if ($2.complexity === "KnownSize") { return NonTrivial; }
              if ($2.complexity === "NonTrivial") { return $2.complexity; }
              $runtime.fail();
            })()
          };
        }
        return analyzeDefault1(expr);
      })();
      const v1 = dictHasSyntax.syntaxOf(expr._1);
      return {
        ...v1.tag === "Just" && v1._1.tag === "Local"
          ? {...callArity(v1._1._2)(expr._2.length)({...analysis, size: analysis.size + 1 | 0}), result: Unknown}
          : {...analysis, result: Unknown, size: analysis.size + 1 | 0},
        args: remainingArgs
      };
    }
    if (expr.tag === "Update") {
      const $0 = analyzeDefault1(expr);
      const analysis = {
        ...$0,
        complexity: (() => {
          if ($0.complexity === "Trivial") { return NonTrivial; }
          if ($0.complexity === "Deref") { return NonTrivial; }
          if ($0.complexity === "KnownSize") { return NonTrivial; }
          if ($0.complexity === "NonTrivial") { return $0.complexity; }
          $runtime.fail();
        })(),
        result: Unknown
      };
      const v2 = dictHasSyntax.syntaxOf(expr._1);
      if (v2.tag === "Just" && v2._1.tag === "Local") { return updated(v2._1._2)(analysis); }
      return analysis;
    }
    if (expr.tag === "CtorSaturated") {
      const $0 = foldMap4(v => analysisOf1(v._2))(expr._5);
      return {...$0, deps: Data$dMap$dInternal.insert(ordQualified)(expr._1)()($0.deps), result: KnownNeutral, size: $0.size + 1 | 0};
    }
    if (expr.tag === "CtorDef") {
      const $0 = analyzeDefault1(expr);
      return {
        ...$0,
        complexity: (() => {
          if ($0.complexity === "Trivial") { return NonTrivial; }
          if ($0.complexity === "Deref") { return NonTrivial; }
          if ($0.complexity === "KnownSize") { return NonTrivial; }
          if ($0.complexity === "NonTrivial") { return $0.complexity; }
          $runtime.fail();
        })()
      };
    }
    if (expr.tag === "Branch") {
      const v2 = (() => {
        if (0 < expr._1.length) { return expr._1[0]; }
        $runtime.fail();
      })();
      const $0 = semigroupBackendAnalysis.append(analysisOf1(v2._1))(semigroupBackendAnalysis.append((() => {
        const $0 = analysisOf1(v2._2);
        return {
          ...$0,
          usages: (() => {
            const go = v => {
              if (v.tag === "Leaf") { return Data$dMap$dInternal.Leaf; }
              if (v.tag === "Node") { return Data$dMap$dInternal.$$$Map("Node", v._1, v._2, v._3, {...v._4, captured: CaptureBranch}, go(v._5), go(v._6)); }
              $runtime.fail();
            };
            return go($0.usages);
          })()
        };
      })())(semigroupBackendAnalysis.append((() => {
        const $0 = foldMap4(foldMap6(analysisOf1))((() => {
          const $0 = Data$dArray.unconsImpl(v => Data$dMaybe.Nothing, v => xs => Data$dMaybe.$Maybe("Just", xs), expr._1);
          if ($0.tag === "Just") { return $0._1; }
          $runtime.fail();
        })());
        return {
          ...$0,
          usages: (() => {
            const go = v => {
              if (v.tag === "Leaf") { return Data$dMap$dInternal.Leaf; }
              if (v.tag === "Node") { return Data$dMap$dInternal.$$$Map("Node", v._1, v._2, v._3, {...v._4, captured: CaptureBranch}, go(v._5), go(v._6)); }
              $runtime.fail();
            };
            return go($0.usages);
          })()
        };
      })())((() => {
        const $0 = analysisOf1(expr._2);
        return {
          ...$0,
          usages: (() => {
            const go = v => {
              if (v.tag === "Leaf") { return Data$dMap$dInternal.Leaf; }
              if (v.tag === "Node") { return Data$dMap$dInternal.$$$Map("Node", v._1, v._2, v._3, {...v._4, captured: CaptureBranch}, go(v._5), go(v._6)); }
              $runtime.fail();
            };
            return go($0.usages);
          })()
        };
      })())));
      return {
        ...$0,
        complexity: (() => {
          if ($0.complexity === "Trivial") { return NonTrivial; }
          if ($0.complexity === "Deref") { return NonTrivial; }
          if ($0.complexity === "KnownSize") { return NonTrivial; }
          if ($0.complexity === "NonTrivial") { return $0.complexity; }
          $runtime.fail();
        })(),
        result: foldMap1(x => dictHasAnalysis.analysisOf(x._2).result)(expr._1)
      };
    }
    if (expr.tag === "Fail") {
      const $0 = analyzeDefault1(expr);
      return {
        ...$0,
        complexity: (() => {
          if ($0.complexity === "Trivial") { return NonTrivial; }
          if ($0.complexity === "Deref") { return NonTrivial; }
          if ($0.complexity === "KnownSize") { return NonTrivial; }
          if ($0.complexity === "NonTrivial") { return $0.complexity; }
          $runtime.fail();
        })()
      };
    }
    if (expr.tag === "PrimOp") {
      const $0 = expr._1;
      const $1 = analyzeDefault1(expr);
      const analysis = {
        ...$1,
        complexity: (() => {
          if ($1.complexity === "Trivial") { return NonTrivial; }
          if ($1.complexity === "Deref") { return NonTrivial; }
          if ($1.complexity === "KnownSize") { return NonTrivial; }
          if ($1.complexity === "NonTrivial") { return $1.complexity; }
          $runtime.fail();
        })(),
        result: Unknown
      };
      const v2 = v3 => {
        if ($0.tag === "Op1" && $0._1.tag === "OpIsTag") { return {...analysis, deps: Data$dMap$dInternal.insert(ordQualified)($0._1._1)()(analysis.deps)}; }
        return analysis;
      };
      if ($0.tag === "Op1" && $0._1.tag === "OpIsTag") {
        const $2 = dictHasSyntax.syntaxOf($0._2);
        if ($2.tag === "Just" && $2._1.tag === "Local") { return cased($2._1._2)({...analysis, deps: Data$dMap$dInternal.insert(ordQualified)($0._1._1)()(analysis.deps)}); }
      }
      return v2(true);
    }
    if (expr.tag === "PrimEffect") {
      const $0 = analyzeDefault1(expr);
      const go = v => {
        if (v.tag === "Leaf") { return Data$dMap$dInternal.Leaf; }
        if (v.tag === "Node") { return Data$dMap$dInternal.$$$Map("Node", v._1, v._2, v._3, {...v._4, captured: CaptureClosure}, go(v._5), go(v._6)); }
        $runtime.fail();
      };
      return {
        ...$0,
        complexity: (() => {
          if ($0.complexity === "Trivial") { return NonTrivial; }
          if ($0.complexity === "Deref") { return NonTrivial; }
          if ($0.complexity === "KnownSize") { return NonTrivial; }
          if ($0.complexity === "NonTrivial") { return $0.complexity; }
          $runtime.fail();
        })(),
        result: Unknown,
        usages: go($0.usages)
      };
    }
    if (expr.tag === "PrimUndefined") { return analyzeDefault1(expr); }
    if (expr.tag === "Accessor") {
      const analysis = (() => {
        if (expr._2.tag === "GetCtorField") {
          const $0 = analyzeDefault1(expr);
          return {...$0, deps: Data$dMap$dInternal.insert(ordQualified)(expr._2._1)()($0.deps), result: Unknown};
        }
        return {...analyzeDefault1(expr), result: Unknown};
      })();
      const v2 = dictHasSyntax.syntaxOf(expr._1);
      if (v2.tag === "Just") {
        if (v2._1.tag === "Accessor") { return analysis; }
        if (v2._1.tag === "Local") { return accessed(v2._1._2)({...analysis, complexity: analysis.complexity === "Trivial" ? Deref : analysis.complexity}); }
        if (v2._1.tag === "Var") {
          if (expr._2.tag === "GetProp") {
            const $0 = externAnalysis(v2._1._1)(Data$dMaybe.$Maybe("Just", expr._2._1));
            if ($0.tag === "Just") { return {...analysis, args: $0._1.args, complexity: analysis.complexity}; }
          }
          return {...analysis, complexity: analysis.complexity};
        }
      }
      return {...analysis, complexity: analysis.complexity === "Trivial" ? Deref : analysis.complexity};
    }
    if (expr.tag === "Lit") {
      const analysis = {...analyzeDefault1(expr), result: KnownNeutral};
      if (expr._1.tag === "LitArray") {
        if (expr._1._1.length > 0) {
          return {
            ...analysis,
            complexity: (() => {
              if (analysis.complexity === "Trivial") { return KnownSize; }
              if (analysis.complexity === "Deref") { return KnownSize; }
              return analysis.complexity;
            })()
          };
        }
        return analysis;
      }
      if (expr._1.tag === "LitRecord") {
        if (expr._1._1.length > 0) {
          return {
            ...analysis,
            complexity: (() => {
              if (analysis.complexity === "Trivial") { return KnownSize; }
              if (analysis.complexity === "Deref") { return KnownSize; }
              return analysis.complexity;
            })()
          };
        }
        return analysis;
      }
      if (expr._1.tag === "LitString" && Data$dString$dCodeUnits.length(expr._1._1) > 128) {
        return {
          ...analysis,
          complexity: (() => {
            if (analysis.complexity === "Trivial") { return KnownSize; }
            if (analysis.complexity === "Deref") { return KnownSize; }
            return analysis.complexity;
          })()
        };
      }
      return analysis;
    }
    $runtime.fail();
  };
};
const analyzeEffectBlock = dictHasAnalysis => {
  const analyzeDefault1 = analyzeDefault(dictHasAnalysis);
  const analyze1 = analyze(dictHasAnalysis);
  return dictHasSyntax => {
    const analyze2 = analyze1(dictHasSyntax);
    return externAnalysis => expr => {
      if (expr.tag === "Let") {
        const $0 = semigroupBackendAnalysis.append(dictHasAnalysis.analysisOf(expr._3))((() => {
          const $0 = dictHasAnalysis.analysisOf(expr._4);
          return {...$0, usages: Data$dMap$dInternal.delete(Data$dOrd.ordInt)(expr._2)($0.usages)};
        })());
        return {
          ...$0,
          complexity: (() => {
            if ($0.complexity === "Trivial") { return NonTrivial; }
            if ($0.complexity === "Deref") { return NonTrivial; }
            if ($0.complexity === "KnownSize") { return NonTrivial; }
            if ($0.complexity === "NonTrivial") { return $0.complexity; }
            $runtime.fail();
          })(),
          result: dictHasAnalysis.analysisOf(expr._4).result,
          size: $0.size + 1 | 0
        };
      }
      if (expr.tag === "LetRec") {
        const $0 = semigroupBackendAnalysis.append(foldMap3(x => dictHasAnalysis.analysisOf(x._2))(expr._2))(dictHasAnalysis.analysisOf(expr._3));
        return {
          ...$0,
          complexity: (() => {
            if ($0.complexity === "Trivial") { return NonTrivial; }
            if ($0.complexity === "Deref") { return NonTrivial; }
            if ($0.complexity === "KnownSize") { return NonTrivial; }
            if ($0.complexity === "NonTrivial") { return $0.complexity; }
            $runtime.fail();
          })(),
          result: dictHasAnalysis.analysisOf(expr._3).result,
          size: $0.size + 1 | 0,
          usages: Data$dMap$dInternal.delete(Data$dOrd.ordInt)(expr._1)($0.usages)
        };
      }
      if (expr.tag === "EffectBind") {
        const $0 = semigroupBackendAnalysis.append(dictHasAnalysis.analysisOf(expr._3))((() => {
          const $0 = dictHasAnalysis.analysisOf(expr._4);
          return {...$0, usages: Data$dMap$dInternal.delete(Data$dOrd.ordInt)(expr._2)($0.usages)};
        })());
        return {
          ...$0,
          complexity: (() => {
            if ($0.complexity === "Trivial") { return NonTrivial; }
            if ($0.complexity === "Deref") { return NonTrivial; }
            if ($0.complexity === "KnownSize") { return NonTrivial; }
            if ($0.complexity === "NonTrivial") { return $0.complexity; }
            $runtime.fail();
          })(),
          result: Unknown,
          size: $0.size + 1 | 0
        };
      }
      if (expr.tag === "EffectPure") {
        const $0 = dictHasAnalysis.analysisOf(expr._1);
        return {...$0, result: Unknown, size: $0.size + 1 | 0};
      }
      if (expr.tag === "EffectDefer") {
        const $0 = dictHasAnalysis.analysisOf(expr._1);
        return {...$0, result: Unknown, size: $0.size + 1 | 0};
      }
      if (expr.tag === "UncurriedEffectApp") {
        const $0 = analyzeDefault1(expr);
        const analysis = {
          ...$0,
          complexity: (() => {
            if ($0.complexity === "Trivial") { return NonTrivial; }
            if ($0.complexity === "Deref") { return NonTrivial; }
            if ($0.complexity === "KnownSize") { return NonTrivial; }
            if ($0.complexity === "NonTrivial") { return $0.complexity; }
            $runtime.fail();
          })(),
          result: Unknown
        };
        const v = dictHasSyntax.syntaxOf(expr._1);
        if (v.tag === "Just" && v._1.tag === "Local") { return callArity(v._1._2)(expr._2.length)(analysis); }
        return analysis;
      }
      if (expr.tag === "PrimEffect") {
        const $0 = analyzeDefault1(expr);
        return {
          ...$0,
          complexity: (() => {
            if ($0.complexity === "Trivial") { return NonTrivial; }
            if ($0.complexity === "Deref") { return NonTrivial; }
            if ($0.complexity === "KnownSize") { return NonTrivial; }
            if ($0.complexity === "NonTrivial") { return $0.complexity; }
            $runtime.fail();
          })(),
          result: Unknown
        };
      }
      return analyze2(externAnalysis)(expr);
    };
  };
};
export {
  $Capture,
  $Complexity,
  $ResultTerm,
  BackendAnalysis,
  CaptureBranch,
  CaptureClosure,
  CaptureNone,
  Deref,
  KnownNeutral,
  KnownSize,
  NonTrivial,
  Trivial,
  Unknown,
  Usage,
  accessed,
  analysisOf,
  analyze,
  analyzeDefault,
  analyzeEffectBlock,
  bound,
  boundArg,
  bump,
  callArity,
  capture,
  cased,
  complex,
  eqCapture,
  eqComplexity,
  eqResultTerm,
  externs,
  foldMap1,
  foldMap2,
  foldMap3,
  foldMap4,
  foldMap6,
  monoidBackendAnalysis,
  monoidCapture,
  monoidComplexity,
  monoidResultTerm,
  monoidUsage,
  newtypeBackendAnalysis_,
  newtypeUsage_,
  ordCapture,
  ordComplexity,
  ordQualified,
  pop,
  resultOf,
  semigroupBackendAnalysis,
  semigroupCapture,
  semigroupComplexity,
  semigroupResultTerm,
  semigroupUsage,
  updated,
  used,
  usedDep,
  withArgs,
  withResult,
  withRewrite
};
