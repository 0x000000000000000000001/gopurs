import * as $runtime from "../runtime.js";
import * as Control$dBind from "../Control.Bind/index.js";
import * as Data$dArray from "../Data.Array/index.js";
import * as Data$dFoldable from "../Data.Foldable/index.js";
import * as Data$dFunction from "../Data.Function/index.js";
import * as Data$dFunctor from "../Data.Functor/index.js";
import * as Data$dList$dTypes from "../Data.List.Types/index.js";
import * as Data$dMap$dInternal from "../Data.Map.Internal/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dMonoid from "../Data.Monoid/index.js";
import * as Data$dOrd from "../Data.Ord/index.js";
import * as Data$dOrdering from "../Data.Ordering/index.js";
import * as Data$dSemiring from "../Data.Semiring/index.js";
import * as Data$dSet from "../Data.Set/index.js";
import * as Data$dTraversable from "../Data.Traversable/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
import * as Data$dUnfoldable from "../Data.Unfoldable/index.js";
import * as PureScript$dBackend$dOptimizer$dCoreFn from "../PureScript.Backend.Optimizer.CoreFn/index.js";
import * as PureScript$dBackend$dOptimizer$dSyntax from "../PureScript.Backend.Optimizer.Syntax/index.js";
const $TcoExpr = (_1, _2) => ({tag: "TcoExpr", _1, _2});
const $TcoRef = (tag, _1, _2) => ({tag, _1, _2});
const ordQualified = /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.ordQualified(Data$dOrd.ordString);
const compare1 = x => y => {
  if (x.tag === "Nothing") {
    if (y.tag === "Nothing") { return Data$dOrdering.EQ; }
    return Data$dOrdering.LT;
  }
  if (y.tag === "Nothing") { return Data$dOrdering.GT; }
  if (x.tag === "Just" && y.tag === "Just") { return Data$dOrd.ordString.compare(x._1)(y._1); }
  $runtime.fail();
};
const traverse = /* #__PURE__ */ (() => Data$dTraversable.traversableArray.traverse(Data$dMaybe.applicativeMaybe))();
const traverse1 = /* #__PURE__ */ (() => Data$dTraversable.traversableArray.traverse(Data$dMaybe.applicativeMaybe))();
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
const identity = x => x;
const foldMap1 = /* #__PURE__ */ (() => Data$dFoldable.foldableArray.foldMap(Data$dMonoid.monoidArray))();
const TcoUsage = x => x;
const TcoTopLevel = value0 => $TcoRef("TcoTopLevel", value0);
const TcoLocal = value0 => value1 => $TcoRef("TcoLocal", value0, value1);
const TcoAnalysis = x => x;
const TcoExpr = value0 => value1 => $TcoExpr(value0, value1);
const semigroupTcoUsage = {
  append: v => v1 => (
    {
      total: v.total + v1.total | 0,
      arities: Data$dMap$dInternal.unsafeUnionWith(Data$dOrd.ordInt.compare, Data$dFunction.const, v.arities, v1.arities),
      call: v.call + v1.call | 0,
      readWrite: v.readWrite + v1.readWrite | 0
    }
  )
};
const newtypeTcoAnalysis_ = {Coercible0: () => {}};
const monoidTcoUsage = {mempty: {total: 0, arities: Data$dMap$dInternal.Leaf, call: 0, readWrite: 0}, Semigroup0: () => semigroupTcoUsage};
const eqTcoRef = {
  eq: x => y => {
    if (x.tag === "TcoTopLevel") {
      return y.tag === "TcoTopLevel" && (x._1._1.tag === "Nothing" ? y._1._1.tag === "Nothing" : x._1._1.tag === "Just" && y._1._1.tag === "Just" && x._1._1._1 === y._1._1._1) && x._1._2 === y._1._2;
    }
    return x.tag === "TcoLocal" && y.tag === "TcoLocal" && (x._1.tag === "Nothing" ? y._1.tag === "Nothing" : x._1.tag === "Just" && y._1.tag === "Just" && x._1._1 === y._1._1) && x._2 === y._2;
  }
};
const ordTcoRef = {
  compare: x => y => {
    if (x.tag === "TcoTopLevel") {
      if (y.tag === "TcoTopLevel") { return ordQualified.compare(x._1)(y._1); }
      return Data$dOrdering.LT;
    }
    if (y.tag === "TcoTopLevel") { return Data$dOrdering.GT; }
    if (x.tag === "TcoLocal" && y.tag === "TcoLocal") {
      const v = compare1(x._1)(y._1);
      if (v === "LT") { return Data$dOrdering.LT; }
      if (v === "GT") { return Data$dOrdering.GT; }
      return Data$dOrd.ordInt.compare(x._2)(y._2);
    }
    $runtime.fail();
  },
  Eq0: () => eqTcoRef
};
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
        const v1 = ordTcoRef.compare(k)(v._3);
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
const withTcoRole = role => v => ({...v, role});
const usedTopLevel = v => Data$dSet.mapMaybe(ordQualified)(v1 => {
  if (v1.tag === "TcoTopLevel") { return Data$dMaybe.$Maybe("Just", v1._1); }
  if (v1.tag === "TcoLocal") { return Data$dMaybe.Nothing; }
  $runtime.fail();
})((() => {
  const go = v$1 => {
    if (v$1.tag === "Leaf") { return Data$dMap$dInternal.Leaf; }
    if (v$1.tag === "Node") { return Data$dMap$dInternal.$$$Map("Node", v$1._1, v$1._2, v$1._3, undefined, go(v$1._5), go(v$1._6)); }
    $runtime.fail();
  };
  return go(v.usages);
})());
const unwindTcoScope = /* #__PURE__ */ (() => {
  const go = go$a0$copy => go$a1$copy => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
    while (go$c) {
      const pop = go$a0, v = go$a1;
      if (v.tag === "Cons") {
        if (v._2.tag === "Nil") {
          go$c = false;
          go$r = Data$dMaybe.$Maybe("Just", Data$dTuple.$Tuple(v._1._1, pop));
          continue;
        }
        go$a0 = Data$dList$dTypes.$List("Cons", v._1._1, pop);
        go$a1 = v._2;
        continue;
      }
      if (v.tag === "Nil") {
        go$c = false;
        go$r = Data$dMaybe.Nothing;
        continue;
      }
      $runtime.fail();
    }
    return go$r;
  };
  return go(Data$dList$dTypes.Nil);
})();
const unTcoExpr = v => v._2;
const tcoRefEffect = ident => v => (
  {...v, usages: Data$dMap$dInternal.insertWith(ordTcoRef)(semigroupTcoUsage.append)(ident)({total: 1, call: 0, arities: Data$dMap$dInternal.Leaf, readWrite: 1})(v.usages)}
);
const tcoRefBinding = ref => v => {
  if (v._2.tag === "Abs") { return Data$dMaybe.$Maybe("Just", {ref, analysis: v._2._2._1, arity: v._2._1.length}); }
  if (v._2.tag === "UncurriedAbs") { return Data$dMaybe.$Maybe("Just", {ref, analysis: v._2._2._1, arity: v._2._1.length}); }
  return Data$dMaybe.Nothing;
};
const tcoRefBindings = toTcoRef => traverse(v => {
  const $0 = toTcoRef(v._1);
  if (v._2._2.tag === "Abs") { return Data$dMaybe.$Maybe("Just", {ref: $0, analysis: v._2._2._2._1, arity: v._2._2._1.length}); }
  if (v._2._2.tag === "UncurriedAbs") { return Data$dMaybe.$Maybe("Just", {ref: $0, analysis: v._2._2._2._1, arity: v._2._2._1.length}); }
  return Data$dMaybe.Nothing;
});
const topLevelTcoRefBindings = mod => tcoRefBindings((() => {
  const $0 = PureScript$dBackend$dOptimizer$dCoreFn.Qualified(Data$dMaybe.$Maybe("Just", mod));
  return x => $TcoRef("TcoTopLevel", $0(x));
})());
const tcoNoTailCalls = v => ({...v, tailCalls: Data$dMap$dInternal.Leaf, role: {...v.role, joins: []}});
const tcoCall = ident => arity => v => (
  {
    ...v,
    usages: Data$dMap$dInternal.insertWith(ordTcoRef)(semigroupTcoUsage.append)(ident)({
      total: 1,
      call: 1,
      arities: Data$dMap$dInternal.$$$Map("Node", 1, 1, arity, undefined, Data$dMap$dInternal.Leaf, Data$dMap$dInternal.Leaf),
      readWrite: 0
    })(v.usages),
    tailCalls: Data$dMap$dInternal.insert(ordTcoRef)(ident)(1)(v.tailCalls)
  }
);
const tcoAnalysisOf = v => v._1;
const syntacticArity = v => {
  if (v.tag === "Abs") { return Data$dMaybe.$Maybe("Just", v._1.length); }
  if (v.tag === "UncurriedAbs") { return Data$dMaybe.$Maybe("Just", v._1.length); }
  return Data$dMaybe.Nothing;
};
const tcoEnvGroup = toTcoRef => {
  const $0 = traverse1(v => {
    const $0 = Data$dTuple.Tuple(toTcoRef(v._1));
    if (v._2.tag === "Abs") { return Data$dMaybe.$Maybe("Just", $0(v._2._1.length)); }
    if (v._2.tag === "UncurriedAbs") { return Data$dMaybe.$Maybe("Just", $0(v._2._1.length)); }
    return Data$dMaybe.Nothing;
  });
  return x => {
    const $1 = $0(x);
    if ($1.tag === "Nothing") { return []; }
    if ($1.tag === "Just") { return $1._1; }
    $runtime.fail();
  };
};
const topLevelTcoEnvGroup = mod => tcoEnvGroup((() => {
  const $0 = PureScript$dBackend$dOptimizer$dCoreFn.Qualified(Data$dMaybe.$Maybe("Just", mod));
  return x => $TcoRef("TcoTopLevel", $0(x));
})());
const popTcoScope = ref => {
  const go = go$a0$copy => go$a1$copy => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
    while (go$c) {
      const stack = go$a0, v = go$a1;
      if (v.tag === "Cons") {
        const $0 = Data$dArray.findIndexImpl(Data$dMaybe.Just, Data$dMaybe.Nothing, eqTcoRef.eq(ref), v._1._2);
        if ($0.tag === "Just") {
          go$c = false;
          go$r = Data$dMaybe.$Maybe("Just", {ident: v._1._1, group: v._1._2, index: $0._1, stack});
          continue;
        }
      }
      if (v.tag === "Cons") {
        go$a0 = Data$dList$dTypes.$List("Cons", v._1._1, stack);
        go$a1 = v._2;
        continue;
      }
      go$c = false;
      go$r = Data$dMaybe.Nothing;
    }
    return go$r;
  };
  return go(Data$dList$dTypes.Nil);
};
const overTcoAnalysis = f => v => $TcoExpr(f(v._1), v._2);
const noTcoRole = {joins: [], isLoop: false};
const semigroupTcoAnalysis = {
  append: v => v1 => (
    {
      usages: Data$dMap$dInternal.unsafeUnionWith(ordTcoRef.compare, semigroupTcoUsage.append, v.usages, v1.usages),
      tailCalls: Data$dMap$dInternal.unsafeUnionWith(ordTcoRef.compare, Data$dSemiring.intAdd, v.tailCalls, v1.tailCalls),
      role: noTcoRole
    }
  )
};
const monoidTcoAnalysis = {mempty: {usages: Data$dMap$dInternal.Leaf, tailCalls: Data$dMap$dInternal.Leaf, role: noTcoRole}, Semigroup0: () => semigroupTcoAnalysis};
const foldMap4 = /* #__PURE__ */ (() => Data$dFoldable.foldableArray.foldMap(monoidTcoAnalysis))();
const foldMap5 = /* #__PURE__ */ (() => Data$dFoldable.foldableArray.foldMap(monoidTcoAnalysis))();
const foldMap6 = f => v => semigroupTcoAnalysis.append(f(v._1))(f(v._2));
const foldMap9 = /* #__PURE__ */ (() => PureScript$dBackend$dOptimizer$dSyntax.foldableBackendSyntax.foldMap(monoidTcoAnalysis))();
const localTcoRefBindings = level => tcoRefBindings(ident => $TcoRef("TcoLocal", Data$dMaybe.$Maybe("Just", ident), level));
const localTcoEnvGroup = level => tcoEnvGroup(ident => $TcoRef("TcoLocal", Data$dMaybe.$Maybe("Just", ident), level));
const isUniformTailCall = v => ref => arity => {
  const $0 = lookup(ref)(v.tailCalls);
  if ($0.tag === "Just") {
    const $1 = lookup(ref)(v.usages);
    if ($1.tag === "Just") {
      const v2 = toUnfoldable($1._1.arities);
      if (v2.length === 1) { return Data$dMaybe.$Maybe("Just", v2[0] === arity && $1._1.total === $0._1); }
      return Data$dMaybe.Nothing;
    }
    if ($1.tag === "Nothing") { return Data$dMaybe.Nothing; }
    $runtime.fail();
  }
  if ($0.tag === "Nothing") { return Data$dMaybe.Nothing; }
  $runtime.fail();
};
const isTailCalledIn = analysis => group => {
  const tailCalled = Data$dArray.mapMaybe(b => isUniformTailCall(analysis)(b.ref)(b.arity))(group);
  return tailCalled.length !== 0 && Data$dArray.allImpl(identity, tailCalled);
};
const tcoRoleIsLoop = group => Data$dArray.allImpl(x => isTailCalledIn(x.analysis)(group), group);
const tcoRoleJoins = env => analysis => group => Control$dBind.arrayBind(isTailCalledIn(analysis)(group) ? [undefined] : [])(() => Data$dArray.nubBy(ordTcoRef.compare)(foldMap1(b => Data$dArray.mapMaybe(v => {
  const $0 = isUniformTailCall(b.analysis)(v._1)(v._2);
  if (
    (() => {
      if ($0.tag === "Just") { return $0._1; }
      if ($0.tag === "Nothing") { return false; }
      $runtime.fail();
    })()
  ) {
    return Data$dMaybe.$Maybe("Just", v._1);
  }
  return Data$dMaybe.Nothing;
})(env))(group)));
const inTcoScope = ref => {
  const go = v => v.tag === "Cons" && (Data$dArray.elem(eqTcoRef)(ref)(v._1._2) ? true : go(v._2));
  return go;
};
const hasTcoRole = v => v.isLoop || v.joins.length !== 0;
const analyze = env => v => {
  const $0 = () => {
    const expr$p = PureScript$dBackend$dOptimizer$dSyntax.functorBackendSyntax.map(analyze(env))(v);
    return $TcoExpr(
      (() => {
        const $0 = foldMap9(tcoAnalysisOf)(expr$p);
        return {...$0, tailCalls: Data$dMap$dInternal.Leaf, role: {...$0.role, joins: []}};
      })(),
      expr$p
    );
  };
  if (v.tag === "Var") { return $TcoExpr(tcoCall($TcoRef("TcoTopLevel", v._1))(0)(monoidTcoAnalysis.mempty), PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Var", v._1)); }
  if (v.tag === "Local") {
    return $TcoExpr(tcoCall($TcoRef("TcoLocal", v._1, v._2))(0)(monoidTcoAnalysis.mempty), PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Local", v._1, v._2));
  }
  if (v.tag === "App") {
    if (v._1.tag === "Local") {
      const tl$p = Data$dFunctor.arrayMap((() => {
        const $1 = analyze(env);
        return x => {
          const $2 = $1(x);
          return $TcoExpr({...$2._1, tailCalls: Data$dMap$dInternal.Leaf, role: {...$2._1.role, joins: []}}, $2._2);
        };
      })())(v._2);
      return $TcoExpr(
        tcoCall($TcoRef("TcoLocal", v._1._1, v._1._2))(tl$p.length)(foldMap4(tcoAnalysisOf)(tl$p)),
        PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("App", analyze(env)(v._1), tl$p)
      );
    }
    if (v._1.tag === "Var") {
      const tl$p = Data$dFunctor.arrayMap((() => {
        const $1 = analyze(env);
        return x => {
          const $2 = $1(x);
          return $TcoExpr({...$2._1, tailCalls: Data$dMap$dInternal.Leaf, role: {...$2._1.role, joins: []}}, $2._2);
        };
      })())(v._2);
      return $TcoExpr(
        tcoCall($TcoRef("TcoTopLevel", v._1._1))(tl$p.length)(foldMap4(tcoAnalysisOf)(tl$p)),
        PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("App", analyze(env)(v._1), tl$p)
      );
    }
    return $0();
  }
  if (v.tag === "UncurriedApp") {
    if (v._1.tag === "Local") {
      const tl$p = Data$dFunctor.arrayMap((() => {
        const $1 = analyze(env);
        return x => {
          const $2 = $1(x);
          return $TcoExpr({...$2._1, tailCalls: Data$dMap$dInternal.Leaf, role: {...$2._1.role, joins: []}}, $2._2);
        };
      })())(v._2);
      return $TcoExpr(
        tcoCall($TcoRef("TcoLocal", v._1._1, v._1._2))(tl$p.length)(foldMap5(tcoAnalysisOf)(tl$p)),
        PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("UncurriedApp", analyze(env)(v._1), tl$p)
      );
    }
    if (v._1.tag === "Var") {
      const tl$p = Data$dFunctor.arrayMap((() => {
        const $1 = analyze(env);
        return x => {
          const $2 = $1(x);
          return $TcoExpr({...$2._1, tailCalls: Data$dMap$dInternal.Leaf, role: {...$2._1.role, joins: []}}, $2._2);
        };
      })())(v._2);
      return $TcoExpr(
        tcoCall($TcoRef("TcoTopLevel", v._1._1))(tl$p.length)(foldMap5(tcoAnalysisOf)(tl$p)),
        PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("UncurriedApp", analyze(env)(v._1), tl$p)
      );
    }
    return $0();
  }
  if (v.tag === "PrimEffect") {
    if (v._1.tag === "EffectRefRead") {
      if (v._1._1.tag === "Local") {
        return $TcoExpr(
          tcoRefEffect($TcoRef("TcoLocal", v._1._1._1, v._1._1._2))(monoidTcoAnalysis.mempty),
          PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("PrimEffect", PureScript$dBackend$dOptimizer$dSyntax.$BackendEffect("EffectRefRead", analyze(env)(v._1._1)))
        );
      }
      return $0();
    }
    if (v._1.tag === "EffectRefWrite" && v._1._1.tag === "Local") {
      const val$p = analyze(env)(v._1._2);
      return $TcoExpr(
        tcoRefEffect($TcoRef("TcoLocal", v._1._1._1, v._1._1._2))(val$p._1),
        PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("PrimEffect", PureScript$dBackend$dOptimizer$dSyntax.$BackendEffect("EffectRefWrite", analyze(env)(v._1._1), val$p))
      );
    }
    return $0();
  }
  if (v.tag === "Branch") {
    const branches$p = Data$dFunctor.arrayMap(v1 => PureScript$dBackend$dOptimizer$dSyntax.$Pair(
      (() => {
        const $1 = analyze(env)(v1._1);
        return $TcoExpr({...$1._1, tailCalls: Data$dMap$dInternal.Leaf, role: {...$1._1.role, joins: []}}, $1._2);
      })(),
      analyze(env)(v1._2)
    ))(v._1);
    const def$p = analyze(env)(v._2);
    return $TcoExpr(
      semigroupTcoAnalysis.append(foldMap4(foldMap6(tcoAnalysisOf))(branches$p))(def$p._1),
      PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Branch", branches$p, def$p)
    );
  }
  if (v.tag === "Let") {
    const binding$p = analyze(env)(v._3);
    const body$p = analyze(env)(v._4);
    if (binding$p._2.tag === "Abs") {
      const role = {isLoop: false, joins: tcoRoleJoins(env)(body$p._1)([{ref: $TcoRef("TcoLocal", v._1, v._2), analysis: binding$p._2._2._1, arity: binding$p._2._1.length}])};
      if (role.joins.length !== 0) {
        return $TcoExpr(
          {...semigroupTcoAnalysis.append(binding$p._2._2._1)(body$p._1), role},
          PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Let", v._1, v._2, binding$p, body$p)
        );
      }
      return $TcoExpr(
        semigroupTcoAnalysis.append({...binding$p._1, tailCalls: Data$dMap$dInternal.Leaf, role: {...binding$p._1.role, joins: []}})(body$p._1),
        PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Let", v._1, v._2, binding$p, body$p)
      );
    }
    if (binding$p._2.tag === "UncurriedAbs") {
      const role = {isLoop: false, joins: tcoRoleJoins(env)(body$p._1)([{ref: $TcoRef("TcoLocal", v._1, v._2), analysis: binding$p._2._2._1, arity: binding$p._2._1.length}])};
      if (role.joins.length !== 0) {
        return $TcoExpr(
          {...semigroupTcoAnalysis.append(binding$p._2._2._1)(body$p._1), role},
          PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Let", v._1, v._2, binding$p, body$p)
        );
      }
    }
    return $TcoExpr(
      semigroupTcoAnalysis.append({...binding$p._1, tailCalls: Data$dMap$dInternal.Leaf, role: {...binding$p._1.role, joins: []}})(body$p._1),
      PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Let", v._1, v._2, binding$p, body$p)
    );
  }
  if (v.tag === "LetRec") {
    const $1 = v._1;
    const env$p = [...tcoEnvGroup(ident => $TcoRef("TcoLocal", Data$dMaybe.$Maybe("Just", ident), $1))(v._2), ...env];
    const bindings$p = Data$dFunctor.arrayMap((() => {
      const $2 = analyze(env$p);
      return m => Data$dTuple.$Tuple(m._1, $2(m._2));
    })())(v._2);
    const body$p = analyze(env$p)(v._3);
    const refBindings = tcoRefBindings(ident => $TcoRef("TcoLocal", Data$dMaybe.$Maybe("Just", ident), $1))(bindings$p);
    const role = {
      isLoop: (() => {
        if (refBindings.tag === "Nothing") { return false; }
        if (refBindings.tag === "Just") {
          const $2 = refBindings._1;
          return Data$dArray.allImpl(x => isTailCalledIn(x.analysis)($2), $2);
        }
        $runtime.fail();
      })(),
      joins: (() => {
        if (refBindings.tag === "Nothing") { return []; }
        if (refBindings.tag === "Just") { return tcoRoleJoins(env)(body$p._1)(refBindings._1); }
        $runtime.fail();
      })()
    };
    if (role.isLoop || role.joins.length !== 0) {
      return $TcoExpr(
        {
          ...semigroupTcoAnalysis.append((() => {
            const $2 = foldMap4(v1 => v1.analysis);
            if (refBindings.tag === "Nothing") { return monoidTcoAnalysis.mempty; }
            if (refBindings.tag === "Just") { return $2(refBindings._1); }
            $runtime.fail();
          })())(body$p._1),
          role
        },
        PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("LetRec", $1, bindings$p, body$p)
      );
    }
    return $TcoExpr(
      semigroupTcoAnalysis.append((() => {
        const $2 = foldMap4(v$1 => v$1._2._1)(bindings$p);
        return {...$2, tailCalls: Data$dMap$dInternal.Leaf, role: {...$2.role, joins: []}};
      })())(body$p._1),
      PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("LetRec", $1, bindings$p, body$p)
    );
  }
  return $0();
};
export {
  $TcoExpr,
  $TcoRef,
  TcoAnalysis,
  TcoExpr,
  TcoLocal,
  TcoTopLevel,
  TcoUsage,
  analyze,
  compare1,
  eqTcoRef,
  foldMap1,
  foldMap4,
  foldMap5,
  foldMap6,
  foldMap9,
  hasTcoRole,
  identity,
  inTcoScope,
  isTailCalledIn,
  isUniformTailCall,
  localTcoEnvGroup,
  localTcoRefBindings,
  lookup,
  monoidTcoAnalysis,
  monoidTcoUsage,
  newtypeTcoAnalysis_,
  noTcoRole,
  ordQualified,
  ordTcoRef,
  overTcoAnalysis,
  popTcoScope,
  semigroupTcoAnalysis,
  semigroupTcoUsage,
  syntacticArity,
  tcoAnalysisOf,
  tcoCall,
  tcoEnvGroup,
  tcoNoTailCalls,
  tcoRefBinding,
  tcoRefBindings,
  tcoRefEffect,
  tcoRoleIsLoop,
  tcoRoleJoins,
  toUnfoldable,
  topLevelTcoEnvGroup,
  topLevelTcoRefBindings,
  traverse,
  traverse1,
  unTcoExpr,
  unwindTcoScope,
  usedTopLevel,
  withTcoRole
};
