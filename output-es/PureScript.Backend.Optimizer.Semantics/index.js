import * as $runtime from "../runtime.js";
import * as Data$dArray from "../Data.Array/index.js";
import * as Data$dArray$dNonEmpty from "../Data.Array.NonEmpty/index.js";
import * as Data$dEither from "../Data.Either/index.js";
import * as Data$dEq from "../Data.Eq/index.js";
import * as Data$dFoldable from "../Data.Foldable/index.js";
import * as Data$dFunction from "../Data.Function/index.js";
import * as Data$dFunctor from "../Data.Functor/index.js";
import * as Data$dHeytingAlgebra from "../Data.HeytingAlgebra/index.js";
import * as Data$dLazy from "../Data.Lazy/index.js";
import * as Data$dList$dTypes from "../Data.List.Types/index.js";
import * as Data$dMap$dInternal from "../Data.Map.Internal/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dMonoid from "../Data.Monoid/index.js";
import * as Data$dOrd from "../Data.Ord/index.js";
import * as Data$dOrdering from "../Data.Ordering/index.js";
import * as Data$dShow from "../Data.Show/index.js";
import * as Data$dString$dCodePoints from "../Data.String.CodePoints/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
import * as Data$dUnfoldable from "../Data.Unfoldable/index.js";
import * as Partial from "../Partial/index.js";
import * as PureScript$dBackend$dOptimizer$dAnalysis from "../PureScript.Backend.Optimizer.Analysis/index.js";
import * as PureScript$dBackend$dOptimizer$dCoreFn from "../PureScript.Backend.Optimizer.CoreFn/index.js";
import * as PureScript$dBackend$dOptimizer$dSyntax from "../PureScript.Backend.Optimizer.Syntax/index.js";
const $BackendExpr = (tag, _1, _2) => ({tag, _1, _2});
const $BackendRewrite = (tag, _1, _2, _3, _4, _5) => ({tag, _1, _2, _3, _4, _5});
const $BackendSemantics = (tag, _1, _2, _3, _4, _5) => ({tag, _1, _2, _3, _4, _5});
const $DistOp = (tag, _1, _2) => ({tag, _1, _2});
const $EvalRef = (tag, _1, _2) => ({tag, _1, _2});
const $ExternImpl = (tag, _1, _2, _3, _4, _5) => ({tag, _1, _2, _3, _4, _5});
const $ExternSpine = (tag, _1) => ({tag, _1});
const $InlineAccessor = (tag, _1) => ({tag, _1});
const $InlineDirective = (tag, _1) => ({tag, _1});
const $LocalBinding = (tag, _1) => ({tag, _1});
const $MkFn = (tag, _1, _2) => ({tag, _1, _2});
const $SemConditional = (_1, _2) => ({tag: "SemConditional", _1, _2});
const $UnpackOp = (tag, _1, _2, _3, _4, _5) => ({tag, _1, _2, _3, _4, _5});
const compare1 = /* #__PURE__ */ (() => PureScript$dBackend$dOptimizer$dCoreFn.ordQualified(Data$dOrd.ordString).compare)();
const compare2 = x => y => {
  if (x.tag === "Nothing") {
    if (y.tag === "Nothing") { return Data$dOrdering.EQ; }
    return Data$dOrdering.LT;
  }
  if (y.tag === "Nothing") { return Data$dOrdering.GT; }
  if (x.tag === "Just" && y.tag === "Just") { return Data$dOrd.ordString.compare(x._1)(y._1); }
  $runtime.fail();
};
const eq10 = /* #__PURE__ */ Data$dEq.eqArrayImpl(x => y => (x._1.tag === "Nothing" ? y._1.tag === "Nothing" : x._1.tag === "Just" && y._1.tag === "Just" && x._1._1 === y._1._1) && x._2 === y._2);
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
        const v1 = Data$dOrd.ordInt.compare(k)(v._3);
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
const or = /* #__PURE__ */ Data$dFoldable.or(Data$dFoldable.foldableArray)(Data$dHeytingAlgebra.heytingAlgebraBoolean);
const and = /* #__PURE__ */ Data$dFoldable.and(Data$dFoldable.foldableArray)(Data$dHeytingAlgebra.heytingAlgebraBoolean);
const foldMap = /* #__PURE__ */ (() => Data$dFoldable.foldableArray.foldMap(PureScript$dBackend$dOptimizer$dAnalysis.monoidBackendAnalysis))();
const foldMap1 = /* #__PURE__ */ (() => Data$dFoldable.foldableArray.foldMap(PureScript$dBackend$dOptimizer$dAnalysis.monoidBackendAnalysis))();
const power = /* #__PURE__ */ Data$dMonoid.power(PureScript$dBackend$dOptimizer$dAnalysis.monoidBackendAnalysis);
const toUnfoldable1 = /* #__PURE__ */ (() => Data$dUnfoldable.unfoldableArray.unfoldr(xs => {
  if (xs.tag === "Nil") { return Data$dMaybe.Nothing; }
  if (xs.tag === "Cons") { return Data$dMaybe.$Maybe("Just", Data$dTuple.$Tuple(xs._1, xs._2)); }
  $runtime.fail();
}))();
const fromFoldable = /* #__PURE__ */ Data$dFoldable.foldrArray(Data$dList$dTypes.Cons)(Data$dList$dTypes.Nil);
const eq16 = x => y => {
  if (x.tag === "Left") {
    return y.tag === "Left" && (x._1._1.tag === "Nothing" ? y._1._1.tag === "Nothing" : x._1._1.tag === "Just" && y._1._1.tag === "Just" && x._1._1._1 === y._1._1._1) && x._1._2 === y._1._2;
  }
  return x.tag === "Right" && y.tag === "Right" && PureScript$dBackend$dOptimizer$dSyntax.eqBackendOperator2.eq(x._1)(y._1);
};
const identity = x => x;
const lookup1 = /* #__PURE__ */ Data$dFoldable.lookup(Data$dFoldable.foldableArray)(Data$dEq.eqString);
const UnpackRecord = value0 => $UnpackOp("UnpackRecord", value0);
const UnpackUpdate = value0 => value1 => $UnpackOp("UnpackUpdate", value0, value1);
const UnpackArray = value0 => $UnpackOp("UnpackArray", value0);
const UnpackData = value0 => value1 => value2 => value3 => value4 => $UnpackOp("UnpackData", value0, value1, value2, value3, value4);
const SemConditional = value0 => value1 => $SemConditional(value0, value1);
const NeutralExpr = x => x;
const MkFnApplied = value0 => $MkFn("MkFnApplied", value0);
const MkFnNext = value0 => value1 => $MkFn("MkFnNext", value0, value1);
const One = value0 => $LocalBinding("One", value0);
const Group = value0 => $LocalBinding("Group", value0);
const InlineDefault = /* #__PURE__ */ $InlineDirective("InlineDefault");
const InlineNever = /* #__PURE__ */ $InlineDirective("InlineNever");
const InlineAlways = /* #__PURE__ */ $InlineDirective("InlineAlways");
const InlineArity = value0 => $InlineDirective("InlineArity", value0);
const InlineProp = value0 => $InlineAccessor("InlineProp", value0);
const InlineSpineProp = value0 => $InlineAccessor("InlineSpineProp", value0);
const InlineRef = /* #__PURE__ */ $InlineAccessor("InlineRef");
const EvalExtern = value0 => $EvalRef("EvalExtern", value0);
const EvalLocal = value0 => value1 => $EvalRef("EvalLocal", value0, value1);
const DistApp = value0 => $DistOp("DistApp", value0);
const DistUncurriedApp = value0 => $DistOp("DistUncurriedApp", value0);
const DistAccessor = value0 => $DistOp("DistAccessor", value0);
const DistPrimOp1 = value0 => $DistOp("DistPrimOp1", value0);
const DistPrimOp2L = value0 => value1 => $DistOp("DistPrimOp2L", value0, value1);
const DistPrimOp2R = value0 => value1 => $DistOp("DistPrimOp2R", value0, value1);
const ExternExpr = value0 => value1 => $ExternImpl("ExternExpr", value0, value1);
const ExternDict = value0 => value1 => $ExternImpl("ExternDict", value0, value1);
const ExternCtor = value0 => value1 => value2 => value3 => value4 => $ExternImpl("ExternCtor", value0, value1, value2, value3, value4);
const SemRef = value0 => value1 => value2 => $BackendSemantics("SemRef", value0, value1, value2);
const SemLam = value0 => value1 => $BackendSemantics("SemLam", value0, value1);
const SemMkFn = value0 => $BackendSemantics("SemMkFn", value0);
const SemMkEffectFn = value0 => $BackendSemantics("SemMkEffectFn", value0);
const SemLet = value0 => value1 => value2 => $BackendSemantics("SemLet", value0, value1, value2);
const SemLetRec = value0 => value1 => $BackendSemantics("SemLetRec", value0, value1);
const SemEffectBind = value0 => value1 => value2 => $BackendSemantics("SemEffectBind", value0, value1, value2);
const SemEffectPure = value0 => $BackendSemantics("SemEffectPure", value0);
const SemEffectDefer = value0 => $BackendSemantics("SemEffectDefer", value0);
const SemBranch = value0 => value1 => $BackendSemantics("SemBranch", value0, value1);
const SemAssocOp = value0 => value1 => $BackendSemantics("SemAssocOp", value0, value1);
const NeutLocal = value0 => value1 => $BackendSemantics("NeutLocal", value0, value1);
const NeutVar = value0 => $BackendSemantics("NeutVar", value0);
const NeutStop = value0 => $BackendSemantics("NeutStop", value0);
const NeutData = value0 => value1 => value2 => value3 => value4 => $BackendSemantics("NeutData", value0, value1, value2, value3, value4);
const NeutCtorDef = value0 => value1 => value2 => value3 => value4 => $BackendSemantics("NeutCtorDef", value0, value1, value2, value3, value4);
const NeutApp = value0 => value1 => $BackendSemantics("NeutApp", value0, value1);
const NeutAccessor = value0 => value1 => $BackendSemantics("NeutAccessor", value0, value1);
const NeutUpdate = value0 => value1 => $BackendSemantics("NeutUpdate", value0, value1);
const NeutLit = value0 => $BackendSemantics("NeutLit", value0);
const NeutFail = value0 => $BackendSemantics("NeutFail", value0);
const NeutUncurriedApp = value0 => value1 => $BackendSemantics("NeutUncurriedApp", value0, value1);
const NeutUncurriedEffectApp = value0 => value1 => $BackendSemantics("NeutUncurriedEffectApp", value0, value1);
const NeutPrimOp = value0 => $BackendSemantics("NeutPrimOp", value0);
const NeutPrimEffect = value0 => $BackendSemantics("NeutPrimEffect", value0);
const NeutPrimUndefined = /* #__PURE__ */ $BackendSemantics("NeutPrimUndefined");
const ExternApp = value0 => $ExternSpine("ExternApp", value0);
const ExternUncurriedApp = value0 => $ExternSpine("ExternUncurriedApp", value0);
const ExternAccessor = value0 => $ExternSpine("ExternAccessor", value0);
const ExternPrimOp = value0 => $ExternSpine("ExternPrimOp", value0);
const Env = x => x;
const RewriteInline = value0 => value1 => value2 => value3 => $BackendRewrite("RewriteInline", value0, value1, value2, value3);
const RewriteUncurry = value0 => value1 => value2 => value3 => value4 => $BackendRewrite("RewriteUncurry", value0, value1, value2, value3, value4);
const RewriteStop = value0 => $BackendRewrite("RewriteStop", value0);
const RewriteUnpackOp = value0 => value1 => value2 => value3 => $BackendRewrite("RewriteUnpackOp", value0, value1, value2, value3);
const RewriteDistBranchesLet = value0 => value1 => value2 => value3 => value4 => $BackendRewrite("RewriteDistBranchesLet", value0, value1, value2, value3, value4);
const RewriteDistBranchesOp = value0 => value1 => value2 => $BackendRewrite("RewriteDistBranchesOp", value0, value1, value2);
const ExprSyntax = value0 => value1 => $BackendExpr("ExprSyntax", value0, value1);
const ExprRewrite = value0 => value1 => $BackendExpr("ExprRewrite", value0, value1);
const Ctx = x => x;
const newtypeNeutralExpr_ = {Coercible0: () => {}};
const newtypeEnv_ = {Coercible0: () => {}};
const hasSyntaxBackendExpr = {
  syntaxOf: v => {
    if (v.tag === "ExprSyntax") { return Data$dMaybe.$Maybe("Just", v._2); }
    return Data$dMaybe.Nothing;
  }
};
const hasAnalysisBackendExpr = {
  analysisOf: v => {
    if (v.tag === "ExprSyntax") { return v._1; }
    if (v.tag === "ExprRewrite") { return v._1; }
    $runtime.fail();
  }
};
const eqUnpackOp = dictEq => {
  const eq19 = Data$dEq.eqArrayImpl(x => y => x._1 === y._1 && dictEq.eq(x._2)(y._2));
  const eq22 = Data$dEq.eqArrayImpl(x => y => x._1 === y._1 && dictEq.eq(x._2)(y._2));
  return {
    eq: x => y => {
      if (x.tag === "UnpackRecord") { return y.tag === "UnpackRecord" && eq19(x._1)(y._1); }
      if (x.tag === "UnpackUpdate") { return y.tag === "UnpackUpdate" && dictEq.eq(x._1)(y._1) && eq19(x._2)(y._2); }
      if (x.tag === "UnpackArray") { return y.tag === "UnpackArray" && Data$dEq.eqArrayImpl(dictEq.eq)(x._1)(y._1); }
      return x.tag === "UnpackData" && y.tag === "UnpackData" && (x._1._1.tag === "Nothing"
        ? y._1._1.tag === "Nothing"
        : x._1._1.tag === "Just" && y._1._1.tag === "Just" && x._1._1._1 === y._1._1._1) && x._1._2 === y._1._2 && (x._2 === "ProductType"
        ? y._2 === "ProductType"
        : x._2 === "SumType" && y._2 === "SumType") && x._3 === y._3 && x._4 === y._4 && eq22(x._5)(y._5);
    }
  };
};
const eqInlineAccessor = {
  eq: x => y => {
    if (x.tag === "InlineProp") { return y.tag === "InlineProp" && x._1 === y._1; }
    if (x.tag === "InlineSpineProp") { return y.tag === "InlineSpineProp" && x._1 === y._1; }
    return x.tag === "InlineRef" && y.tag === "InlineRef";
  }
};
const ordInlineAccessor = {
  compare: x => y => {
    if (x.tag === "InlineProp") {
      if (y.tag === "InlineProp") { return Data$dOrd.ordString.compare(x._1)(y._1); }
      return Data$dOrdering.LT;
    }
    if (y.tag === "InlineProp") { return Data$dOrdering.GT; }
    if (x.tag === "InlineSpineProp") {
      if (y.tag === "InlineSpineProp") { return Data$dOrd.ordString.compare(x._1)(y._1); }
      return Data$dOrdering.LT;
    }
    if (y.tag === "InlineSpineProp") { return Data$dOrdering.GT; }
    if (x.tag === "InlineRef" && y.tag === "InlineRef") { return Data$dOrdering.EQ; }
    $runtime.fail();
  },
  Eq0: () => eqInlineAccessor
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
        const v1 = ordInlineAccessor.compare(k)(v._3);
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
const eqEvalRef = {
  eq: x => y => {
    if (x.tag === "EvalExtern") {
      return y.tag === "EvalExtern" && (x._1._1.tag === "Nothing" ? y._1._1.tag === "Nothing" : x._1._1.tag === "Just" && y._1._1.tag === "Just" && x._1._1._1 === y._1._1._1) && x._1._2 === y._1._2;
    }
    return x.tag === "EvalLocal" && y.tag === "EvalLocal" && (x._1.tag === "Nothing" ? y._1.tag === "Nothing" : x._1.tag === "Just" && y._1.tag === "Just" && x._1._1 === y._1._1) && x._2 === y._2;
  }
};
const ordEvalRef = {
  compare: x => y => {
    if (x.tag === "EvalExtern") {
      if (y.tag === "EvalExtern") { return compare1(x._1)(y._1); }
      return Data$dOrdering.LT;
    }
    if (y.tag === "EvalExtern") { return Data$dOrdering.GT; }
    if (x.tag === "EvalLocal" && y.tag === "EvalLocal") {
      const v = compare2(x._1)(y._1);
      if (v === "LT") { return Data$dOrdering.LT; }
      if (v === "GT") { return Data$dOrdering.GT; }
      return Data$dOrd.ordInt.compare(x._2)(y._2);
    }
    $runtime.fail();
  },
  Eq0: () => eqEvalRef
};
const alter = /* #__PURE__ */ Data$dMap$dInternal.alter(ordEvalRef);
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
        const v1 = ordEvalRef.compare(k)(v._3);
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
const eqDistOp = dictEq => (
  {
    eq: x => y => {
      if (x.tag === "DistApp") { return y.tag === "DistApp" && Data$dEq.eqArrayImpl(dictEq.eq)(x._1)(y._1); }
      if (x.tag === "DistUncurriedApp") { return y.tag === "DistUncurriedApp" && Data$dEq.eqArrayImpl(dictEq.eq)(x._1)(y._1); }
      if (x.tag === "DistAccessor") { return y.tag === "DistAccessor" && PureScript$dBackend$dOptimizer$dSyntax.eqBackendAccessor.eq(x._1)(y._1); }
      if (x.tag === "DistPrimOp1") {
        return y.tag === "DistPrimOp1" && (() => {
          if (x._1.tag === "OpBooleanNot") { return y._1.tag === "OpBooleanNot"; }
          if (x._1.tag === "OpIntBitNot") { return y._1.tag === "OpIntBitNot"; }
          if (x._1.tag === "OpIntNegate") { return y._1.tag === "OpIntNegate"; }
          if (x._1.tag === "OpNumberNegate") { return y._1.tag === "OpNumberNegate"; }
          if (x._1.tag === "OpArrayLength") { return y._1.tag === "OpArrayLength"; }
          return x._1.tag === "OpIsTag" && y._1.tag === "OpIsTag" && (x._1._1._1.tag === "Nothing"
            ? y._1._1._1.tag === "Nothing"
            : x._1._1._1.tag === "Just" && y._1._1._1.tag === "Just" && x._1._1._1._1 === y._1._1._1._1) && x._1._1._2 === y._1._1._2;
        })();
      }
      if (x.tag === "DistPrimOp2L") { return y.tag === "DistPrimOp2L" && PureScript$dBackend$dOptimizer$dSyntax.eqBackendOperator2.eq(x._1)(y._1) && dictEq.eq(x._2)(y._2); }
      return x.tag === "DistPrimOp2R" && y.tag === "DistPrimOp2R" && dictEq.eq(x._1)(y._1) && PureScript$dBackend$dOptimizer$dSyntax.eqBackendOperator2.eq(x._2)(y._2);
    }
  }
);
const eqBackendRewrite = dictEq => {
  const eq21 = Data$dEq.eqArrayImpl(x => y => dictEq.eq(x._1)(y._1) && dictEq.eq(x._2)(y._2));
  return {
    eq: x => y => {
      if (x.tag === "RewriteInline") {
        return y.tag === "RewriteInline" && (x._1.tag === "Nothing" ? y._1.tag === "Nothing" : x._1.tag === "Just" && y._1.tag === "Just" && x._1._1 === y._1._1) && x._2 === y._2 && dictEq.eq(x._3)(y._3) && dictEq.eq(x._4)(y._4);
      }
      if (x.tag === "RewriteUncurry") {
        return y.tag === "RewriteUncurry" && (x._1.tag === "Nothing" ? y._1.tag === "Nothing" : x._1.tag === "Just" && y._1.tag === "Just" && x._1._1 === y._1._1) && x._2 === y._2 && eq10(x._3)(y._3) && dictEq.eq(x._4)(y._4) && dictEq.eq(x._5)(y._5);
      }
      if (x.tag === "RewriteStop") {
        return y.tag === "RewriteStop" && (x._1._1.tag === "Nothing" ? y._1._1.tag === "Nothing" : x._1._1.tag === "Just" && y._1._1.tag === "Just" && x._1._1._1 === y._1._1._1) && x._1._2 === y._1._2;
      }
      if (x.tag === "RewriteUnpackOp") {
        return y.tag === "RewriteUnpackOp" && (x._1.tag === "Nothing" ? y._1.tag === "Nothing" : x._1.tag === "Just" && y._1.tag === "Just" && x._1._1 === y._1._1) && x._2 === y._2 && eqUnpackOp(dictEq).eq(x._3)(y._3) && dictEq.eq(x._4)(y._4);
      }
      if (x.tag === "RewriteDistBranchesLet") {
        return y.tag === "RewriteDistBranchesLet" && (x._1.tag === "Nothing" ? y._1.tag === "Nothing" : x._1.tag === "Just" && y._1.tag === "Just" && x._1._1 === y._1._1) && x._2 === y._2 && eq21(x._3)(y._3) && dictEq.eq(x._4)(y._4) && dictEq.eq(x._5)(y._5);
      }
      return x.tag === "RewriteDistBranchesOp" && y.tag === "RewriteDistBranchesOp" && eq21(x._1)(y._1) && dictEq.eq(x._2)(y._2) && eqDistOp(dictEq).eq(x._3)(y._3);
    }
  };
};
const eqBackendExpr = {
  eq: v => v1 => {
    if (v.tag === "ExprSyntax") {
      return v1.tag === "ExprSyntax" && v._1.size === v1._1.size && PureScript$dBackend$dOptimizer$dSyntax.eqBackendSyntax(eqBackendExpr).eq(v._2)(v1._2);
    }
    return v.tag === "ExprRewrite" && v1.tag === "ExprRewrite" && v._1.size === v1._1.size && eqBackendRewrite(eqBackendExpr).eq(v._2)(v1._2);
  }
};
const snocApp = prev => next => {
  const $0 = prev.length - 1 | 0;
  if ($0 >= 0 && $0 < prev.length && prev[$0].tag === "ExternApp") {
    return Data$dArray.snoc((() => {
      const $1 = prev.length - 1 | 0;
      if ($1 < 1) { return []; }
      return Data$dArray.sliceImpl(0, $1, prev);
    })())($ExternSpine("ExternApp", Data$dArray.snoc(prev[$0]._1)(next)));
  }
  return Data$dArray.snoc(prev)($ExternSpine("ExternApp", [next]));
};
const snocSpine = spine => v => {
  if (v.tag === "ExternApp") { return Data$dFoldable.foldlArray(snocApp)(spine)(v._1); }
  return Data$dArray.snoc(spine)(v);
};
const simplifyCondIsTag = v => v1 => v2 => {
  if (
    v1._1.tag === "ExprSyntax" && v1._1._2.tag === "PrimOp" && v1._1._2._1.tag === "Op1" && v1._1._2._1._1.tag === "OpIsTag" && v1._2.tag === "ExprSyntax" && v1._2._2.tag === "Lit" && v1._2._2._1.tag === "LitBoolean" && !v1._2._2._1._1 && v2.tag === "ExprSyntax" && v2._2.tag === "PrimOp" && v2._2._1.tag === "Op1" && v2._2._1._1.tag === "OpIsTag" && eqBackendExpr.eq(v1._1._2._1._2)(v2._2._1._2)
  ) {
    return Data$dMaybe.$Maybe("Just", v2);
  }
  return Data$dMaybe.Nothing;
};
const shouldUnpackUpdate = ident => level => binding => body => {
  const $0 = (() => {
    if (body.tag === "ExprSyntax") { return body._1; }
    if (body.tag === "ExprRewrite") { return body._1; }
    $runtime.fail();
  })();
  if (binding.tag === "ExprSyntax" && binding._2.tag === "Update") {
    const $1 = lookup(level)($0.usages);
    if ($1.tag === "Just" && $1._1.total === ($1._1.access + $1._1.update | 0)) {
      return Data$dMaybe.$Maybe(
        "Just",
        $BackendExpr(
          "ExprRewrite",
          {
            ...PureScript$dBackend$dOptimizer$dAnalysis.updated(level)(PureScript$dBackend$dOptimizer$dAnalysis.semigroupBackendAnalysis.append((() => {
              if (binding._2._1.tag === "ExprSyntax") { return binding._2._1._1; }
              if (binding._2._1.tag === "ExprRewrite") { return binding._2._1._1; }
              $runtime.fail();
            })())(Data$dFoldable.foldrArray(x => PureScript$dBackend$dOptimizer$dAnalysis.semigroupBackendAnalysis.append((() => {
              if (x._2.tag === "ExprSyntax") { return x._2._1; }
              if (x._2.tag === "ExprRewrite") { return x._2._1; }
              $runtime.fail();
            })()))({
              ...$0,
              complexity: (() => {
                if ($0.complexity === "Trivial") { return PureScript$dBackend$dOptimizer$dAnalysis.NonTrivial; }
                if ($0.complexity === "Deref") { return PureScript$dBackend$dOptimizer$dAnalysis.NonTrivial; }
                if ($0.complexity === "KnownSize") { return PureScript$dBackend$dOptimizer$dAnalysis.NonTrivial; }
                if ($0.complexity === "NonTrivial") { return $0.complexity; }
                $runtime.fail();
              })(),
              usages: Data$dMap$dInternal.delete(Data$dOrd.ordInt)(level)($0.usages)
            })(binding._2._2))),
            rewrite: true
          },
          $BackendRewrite("RewriteUnpackOp", ident, level, $UnpackOp("UnpackUpdate", binding._2._1, binding._2._2), body)
        )
      );
    }
  }
  return Data$dMaybe.Nothing;
};
const shouldUnpackRecord = ident => level => binding => body => {
  const $0 = (() => {
    if (body.tag === "ExprSyntax") { return body._1; }
    if (body.tag === "ExprRewrite") { return body._1; }
    $runtime.fail();
  })();
  if (binding.tag === "ExprSyntax" && binding._2.tag === "Lit" && binding._2._1.tag === "LitRecord") {
    const $1 = lookup(level)($0.usages);
    if ($1.tag === "Just" && $1._1.total === ($1._1.access + $1._1.update | 0)) {
      return Data$dMaybe.$Maybe(
        "Just",
        $BackendExpr(
          "ExprRewrite",
          {
            ...Data$dFoldable.foldrArray(x => PureScript$dBackend$dOptimizer$dAnalysis.semigroupBackendAnalysis.append((() => {
              if (x._2.tag === "ExprSyntax") { return x._2._1; }
              if (x._2.tag === "ExprRewrite") { return x._2._1; }
              $runtime.fail();
            })()))({
              ...$0,
              complexity: (() => {
                if ($0.complexity === "Trivial") { return PureScript$dBackend$dOptimizer$dAnalysis.NonTrivial; }
                if ($0.complexity === "Deref") { return PureScript$dBackend$dOptimizer$dAnalysis.NonTrivial; }
                if ($0.complexity === "KnownSize") { return PureScript$dBackend$dOptimizer$dAnalysis.NonTrivial; }
                if ($0.complexity === "NonTrivial") { return $0.complexity; }
                $runtime.fail();
              })(),
              usages: Data$dMap$dInternal.delete(Data$dOrd.ordInt)(level)($0.usages)
            })(binding._2._1._1),
            rewrite: true
          },
          $BackendRewrite("RewriteUnpackOp", ident, level, $UnpackOp("UnpackRecord", binding._2._1._1), body)
        )
      );
    }
  }
  return Data$dMaybe.Nothing;
};
const shouldUnpackCtor = ident => level => a => body => {
  const $0 = (() => {
    if (body.tag === "ExprSyntax") { return body._1; }
    if (body.tag === "ExprRewrite") { return body._1; }
    $runtime.fail();
  })();
  if (a.tag === "ExprSyntax" && a._2.tag === "CtorSaturated") {
    const $1 = lookup(level)($0.usages);
    if ($1.tag === "Just" && $1._1.total === ($1._1.access + $1._1.case | 0)) {
      return Data$dMaybe.$Maybe(
        "Just",
        $BackendExpr(
          "ExprRewrite",
          {
            ...Data$dFoldable.foldrArray(x => PureScript$dBackend$dOptimizer$dAnalysis.semigroupBackendAnalysis.append((() => {
              if (x._2.tag === "ExprSyntax") { return x._2._1; }
              if (x._2.tag === "ExprRewrite") { return x._2._1; }
              $runtime.fail();
            })()))({
              ...$0,
              complexity: (() => {
                if ($0.complexity === "Trivial") { return PureScript$dBackend$dOptimizer$dAnalysis.NonTrivial; }
                if ($0.complexity === "Deref") { return PureScript$dBackend$dOptimizer$dAnalysis.NonTrivial; }
                if ($0.complexity === "KnownSize") { return PureScript$dBackend$dOptimizer$dAnalysis.NonTrivial; }
                if ($0.complexity === "NonTrivial") { return $0.complexity; }
                $runtime.fail();
              })(),
              usages: Data$dMap$dInternal.delete(Data$dOrd.ordInt)(level)($0.usages)
            })(a._2._5),
            rewrite: true
          },
          $BackendRewrite("RewriteUnpackOp", ident, level, $UnpackOp("UnpackData", a._2._1, a._2._2, a._2._3, a._2._4, a._2._5), body)
        )
      );
    }
  }
  return Data$dMaybe.Nothing;
};
const shouldUnpackArray = ident => level => binding => body => {
  const $0 = (() => {
    if (body.tag === "ExprSyntax") { return body._1; }
    if (body.tag === "ExprRewrite") { return body._1; }
    $runtime.fail();
  })();
  if (binding.tag === "ExprSyntax" && binding._2.tag === "Lit" && binding._2._1.tag === "LitArray") {
    const $1 = lookup(level)($0.usages);
    if ($1.tag === "Just" && $1._1.total === $1._1.access) {
      return Data$dMaybe.$Maybe(
        "Just",
        $BackendExpr(
          "ExprRewrite",
          {
            ...Data$dFoldable.foldrArray(x => PureScript$dBackend$dOptimizer$dAnalysis.semigroupBackendAnalysis.append((() => {
              if (x.tag === "ExprSyntax") { return x._1; }
              if (x.tag === "ExprRewrite") { return x._1; }
              $runtime.fail();
            })()))({
              ...$0,
              complexity: (() => {
                if ($0.complexity === "Trivial") { return PureScript$dBackend$dOptimizer$dAnalysis.NonTrivial; }
                if ($0.complexity === "Deref") { return PureScript$dBackend$dOptimizer$dAnalysis.NonTrivial; }
                if ($0.complexity === "KnownSize") { return PureScript$dBackend$dOptimizer$dAnalysis.NonTrivial; }
                if ($0.complexity === "NonTrivial") { return $0.complexity; }
                $runtime.fail();
              })(),
              usages: Data$dMap$dInternal.delete(Data$dOrd.ordInt)(level)($0.usages)
            })(binding._2._1._1),
            rewrite: true
          },
          $BackendRewrite("RewriteUnpackOp", ident, level, $UnpackOp("UnpackArray", binding._2._1._1), body)
        )
      );
    }
  }
  return Data$dMaybe.Nothing;
};
const shouldUncurryAbs = ident => level => a => b => {
  const $0 = (() => {
    if (b.tag === "ExprSyntax") { return b._1; }
    if (b.tag === "ExprRewrite") { return b._1; }
    $runtime.fail();
  })();
  if (a.tag === "ExprSyntax" && a._2.tag === "Abs") {
    const $1 = lookup(level)($0.usages);
    if ($1.tag === "Just") {
      const $2 = toUnfoldable($1._1.arities);
      if ($2.length === 1 && $2[0] === a._2._1.length) {
        return Data$dMaybe.$Maybe(
          "Just",
          $BackendExpr(
            "ExprRewrite",
            (() => {
              const $3 = PureScript$dBackend$dOptimizer$dAnalysis.semigroupBackendAnalysis.append((() => {
                if (a.tag === "ExprSyntax") { return a._1; }
                if (a.tag === "ExprRewrite") { return a._1; }
                $runtime.fail();
              })())((() => {
                const $3 = (() => {
                  if (b.tag === "ExprSyntax") { return b._1; }
                  if (b.tag === "ExprRewrite") { return b._1; }
                  $runtime.fail();
                })();
                return {...$3, usages: Data$dMap$dInternal.delete(Data$dOrd.ordInt)(level)($3.usages)};
              })());
              return {
                ...$3,
                complexity: (() => {
                  if ($3.complexity === "Trivial") { return PureScript$dBackend$dOptimizer$dAnalysis.NonTrivial; }
                  if ($3.complexity === "Deref") { return PureScript$dBackend$dOptimizer$dAnalysis.NonTrivial; }
                  if ($3.complexity === "KnownSize") { return PureScript$dBackend$dOptimizer$dAnalysis.NonTrivial; }
                  if ($3.complexity === "NonTrivial") { return $3.complexity; }
                  $runtime.fail();
                })(),
                result: (() => {
                  if (b.tag === "ExprSyntax") { return b._1.result; }
                  if (b.tag === "ExprRewrite") { return b._1.result; }
                  $runtime.fail();
                })(),
                rewrite: true,
                size: $3.size + 1 | 0
              };
            })(),
            $BackendRewrite("RewriteUncurry", ident, level, a._2._1, a._2._2, b)
          )
        );
      }
    }
  }
  return Data$dMaybe.Nothing;
};
const shouldInlineExternReference = v => v1 => v2 => (v1.complexity === "Trivial" || v1.complexity === "Deref") && v1.size < 16;
const shouldInlineExternLiteral = v => {
  if (v.tag === "LitInt") { return true; }
  if (v.tag === "LitNumber") { return true; }
  if (v.tag === "LitString") { return Data$dString$dCodePoints.toCodePointArray(v._1).length <= 32; }
  if (v.tag === "LitChar") { return true; }
  if (v.tag === "LitBoolean") { return true; }
  if (v.tag === "LitArray") { return v._1.length === 0; }
  if (v.tag === "LitRecord") { return v._1.length === 0; }
  $runtime.fail();
};
const shouldInlineExternAppArg = v => v1 => v1.tag === "SemLam" && (v.captured === "CaptureNone" || v.captured === "CaptureBranch") && v.total > 0 && v.call === v.total;
const shouldInlineExternApp = v => v1 => v2 => args => {
  const delayed = v1.args.length > 0;
  return (v1.complexity === "Trivial" || v1.complexity === "Deref") && v1.size < 16 || v1.usages.tag === "Leaf" && !v1.externs && v1.size < 64 || delayed && v1.args.length <= args.length && v1.size < 16 || delayed && or(Data$dArray.zipWithImpl(
    shouldInlineExternAppArg,
    v1.args,
    args
  )) && v1.size < 16;
};
const shouldInlineExternAccessor = v => v1 => v2 => v3 => (v1.complexity === "Trivial" || v1.complexity === "Deref") && v1.size < 16;
const shouldEtaReduce = level1 => binding => v => {
  if (
    v.tag === "ExprSyntax" && v._2.tag === "Abs" && v._2._2.tag === "ExprSyntax" && v._2._2._2.tag === "App" && v._2._2._2._1.tag === "ExprSyntax" && v._2._2._2._1._2.tag === "Local"
  ) {
    const $0 = v._2._2._2._2;
    if (
      level1 === v._2._2._2._1._2._2 && v._2._1.length === $0.length && and(Data$dArray.zipWithImpl(
        v$1 => {
          const $1 = v$1._2;
          return v1 => v1.tag === "ExprSyntax" && v1._2.tag === "Local" && $1 === v1._2._2;
        },
        v._2._1,
        $0
      ))
    ) {
      return Data$dMaybe.$Maybe("Just", binding);
    }
  }
  return Data$dMaybe.Nothing;
};
const shouldDistributeBranches = ident => level => a => body => {
  const $0 = (() => {
    if (body.tag === "ExprSyntax") { return body._1; }
    if (body.tag === "ExprRewrite") { return body._1; }
    $runtime.fail();
  })();
  if (a.tag === "ExprSyntax" && a._2.tag === "Branch" && $0.size <= 128 && a._1.result === "KnownNeutral") {
    const $1 = lookup(level)($0.usages);
    if ($1.tag === "Just" && $1._1.total === ($1._1.access + $1._1.case | 0)) {
      return Data$dMaybe.$Maybe(
        "Just",
        $BackendExpr(
          "ExprRewrite",
          {
            ...PureScript$dBackend$dOptimizer$dAnalysis.semigroupBackendAnalysis.append((() => {
              if (a.tag === "ExprSyntax") { return a._1; }
              if (a.tag === "ExprRewrite") { return a._1; }
              $runtime.fail();
            })())((() => {
              const $2 = (() => {
                if (body.tag === "ExprSyntax") { return body._1; }
                if (body.tag === "ExprRewrite") { return body._1; }
                $runtime.fail();
              })();
              return {...$2, usages: Data$dMap$dInternal.delete(Data$dOrd.ordInt)(level)($2.usages)};
            })()),
            rewrite: true
          },
          $BackendRewrite("RewriteDistBranchesLet", ident, level, a._2._1, a._2._2, body)
        )
      );
    }
  }
  return Data$dMaybe.Nothing;
};
const shouldDistributeBranchUncurriedApps = analysis1 => branches => def => spine => {
  if (
    Data$dArray.allImpl(
      x => {
        const $0 = (() => {
          if (x.tag === "ExprSyntax") { return x._1; }
          if (x.tag === "ExprRewrite") { return x._1; }
          $runtime.fail();
        })();
        return $0.complexity === "Trivial" || $0.complexity === "Deref";
      },
      spine
    )
  ) {
    return Data$dMaybe.$Maybe(
      "Just",
      $BackendExpr(
        "ExprRewrite",
        {...PureScript$dBackend$dOptimizer$dAnalysis.semigroupBackendAnalysis.append(analysis1)(foldMap(hasAnalysisBackendExpr.analysisOf)(spine)), rewrite: true},
        $BackendRewrite("RewriteDistBranchesOp", branches, def, $DistOp("DistUncurriedApp", spine))
      )
    );
  }
  return Data$dMaybe.Nothing;
};
const shouldDistributeBranchPrimOp2R = analysis1 => branches => def => lhs => op2 => {
  if (
    (() => {
      const $0 = (() => {
        if (lhs.tag === "ExprSyntax") { return lhs._1.complexity; }
        if (lhs.tag === "ExprRewrite") { return lhs._1.complexity; }
        $runtime.fail();
      })();
      return $0 === "Trivial" || $0 === "Deref";
    })()
  ) {
    return Data$dMaybe.$Maybe(
      "Just",
      $BackendExpr(
        "ExprRewrite",
        (() => {
          const $0 = PureScript$dBackend$dOptimizer$dAnalysis.semigroupBackendAnalysis.append(analysis1)((() => {
            if (lhs.tag === "ExprSyntax") { return lhs._1; }
            if (lhs.tag === "ExprRewrite") { return lhs._1; }
            $runtime.fail();
          })());
          return {...$0, rewrite: true, size: $0.size + 1 | 0};
        })(),
        $BackendRewrite("RewriteDistBranchesOp", branches, def, $DistOp("DistPrimOp2R", lhs, op2))
      )
    );
  }
  return Data$dMaybe.Nothing;
};
const shouldDistributeBranchPrimOp2L = analysis1 => branches => def => op2 => rhs => {
  if (
    (() => {
      const $0 = (() => {
        if (rhs.tag === "ExprSyntax") { return rhs._1.complexity; }
        if (rhs.tag === "ExprRewrite") { return rhs._1.complexity; }
        $runtime.fail();
      })();
      return $0 === "Trivial" || $0 === "Deref";
    })()
  ) {
    return Data$dMaybe.$Maybe(
      "Just",
      $BackendExpr(
        "ExprRewrite",
        (() => {
          const $0 = PureScript$dBackend$dOptimizer$dAnalysis.semigroupBackendAnalysis.append(analysis1)((() => {
            if (rhs.tag === "ExprSyntax") { return rhs._1; }
            if (rhs.tag === "ExprRewrite") { return rhs._1; }
            $runtime.fail();
          })());
          return {...$0, rewrite: true, size: $0.size + 1 | 0};
        })(),
        $BackendRewrite("RewriteDistBranchesOp", branches, def, $DistOp("DistPrimOp2L", op2, rhs))
      )
    );
  }
  return Data$dMaybe.Nothing;
};
const shouldDistributeBranchPrimOp1 = analysis1 => branches => def => op => Data$dMaybe.$Maybe(
  "Just",
  $BackendExpr("ExprRewrite", {...analysis1, rewrite: true, size: analysis1.size + 1 | 0}, $BackendRewrite("RewriteDistBranchesOp", branches, def, $DistOp("DistPrimOp1", op)))
);
const shouldDistributeBranchApps = analysis1 => branches => def => spine => {
  if (
    Data$dArray.allImpl(
      x => {
        const $0 = (() => {
          if (x.tag === "ExprSyntax") { return x._1; }
          if (x.tag === "ExprRewrite") { return x._1; }
          $runtime.fail();
        })();
        return $0.complexity === "Trivial" || $0.complexity === "Deref";
      },
      spine
    )
  ) {
    return Data$dMaybe.$Maybe(
      "Just",
      $BackendExpr(
        "ExprRewrite",
        {...PureScript$dBackend$dOptimizer$dAnalysis.semigroupBackendAnalysis.append(analysis1)(foldMap1(hasAnalysisBackendExpr.analysisOf)(spine)), rewrite: true},
        $BackendRewrite("RewriteDistBranchesOp", branches, def, $DistOp("DistApp", spine))
      )
    );
  }
  return Data$dMaybe.Nothing;
};
const shouldDistributeBranchAccessor = analysis1 => branches => def => acc => Data$dMaybe.$Maybe(
  "Just",
  $BackendExpr("ExprRewrite", {...analysis1, rewrite: true, size: analysis1.size + 1 | 0}, $BackendRewrite("RewriteDistBranchesOp", branches, def, $DistOp("DistAccessor", acc)))
);
const rewriteInline = ident => level => binding => body => {
  const s2 = (() => {
    if (body.tag === "ExprSyntax") { return body._1; }
    if (body.tag === "ExprRewrite") { return body._1; }
    $runtime.fail();
  })();
  return $BackendExpr(
    "ExprRewrite",
    (() => {
      const v = lookup(level)(s2.usages);
      const $0 = (() => {
        if (v.tag === "Just") {
          return PureScript$dBackend$dOptimizer$dAnalysis.semigroupBackendAnalysis.append(s2)(power((() => {
            if (binding.tag === "ExprSyntax") { return binding._1; }
            if (binding.tag === "ExprRewrite") { return binding._1; }
            $runtime.fail();
          })())(v._1.total));
        }
        if (v.tag === "Nothing") { return s2; }
        $runtime.fail();
      })();
      return {...$0, rewrite: true, usages: Data$dMap$dInternal.delete(Data$dOrd.ordInt)(level)($0.usages)};
    })(),
    $BackendRewrite("RewriteInline", ident, level, binding, body)
  );
};
const rewriteBranches = k => {
  const go = v => {
    if (v.tag === "SemLet") { return $BackendSemantics("SemLet", v._1, v._2, x => go(v._3(x))); }
    if (v.tag === "SemLetRec") { return $BackendSemantics("SemLetRec", v._1, x => go(v._2(x))); }
    if (v.tag === "SemBranch") {
      const $0 = v._2;
      return $BackendSemantics(
        "SemBranch",
        Data$dFunctor.arrayMap(v1 => {
          const $1 = v1._2;
          return $SemConditional(v1._1, Data$dLazy.defer(v$1 => go(Data$dLazy.force($1))));
        })(v._1),
        Data$dLazy.defer(v$1 => go(Data$dLazy.force($0)))
      );
    }
    return k(v);
  };
  return go;
};
const purely = v => {
  if (v.effect) { return {...v, effect: false}; }
  return v;
};
const primOpOrdNot = v => {
  if (v === "OpEq") { return PureScript$dBackend$dOptimizer$dSyntax.OpNotEq; }
  if (v === "OpNotEq") { return PureScript$dBackend$dOptimizer$dSyntax.OpEq; }
  if (v === "OpLt") { return PureScript$dBackend$dOptimizer$dSyntax.OpGte; }
  if (v === "OpLte") { return PureScript$dBackend$dOptimizer$dSyntax.OpGt; }
  if (v === "OpGt") { return PureScript$dBackend$dOptimizer$dSyntax.OpLte; }
  if (v === "OpGte") { return PureScript$dBackend$dOptimizer$dSyntax.OpLt; }
  $runtime.fail();
};
const nextLevel = v => Data$dTuple.$Tuple(v.currentLevel, {...v, currentLevel: v.currentLevel + 1 | 0});
const neutralSpine = /* #__PURE__ */ Data$dFoldable.foldlArray(hd => v => {
  if (v.tag === "ExternApp") { return $BackendSemantics("NeutApp", hd, v._1); }
  if (v.tag === "ExternUncurriedApp") { return $BackendSemantics("NeutUncurriedApp", hd, v._1); }
  if (v.tag === "ExternAccessor") { return $BackendSemantics("NeutAccessor", hd, v._1); }
  if (v.tag === "ExternPrimOp") { return $BackendSemantics("NeutPrimOp", PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op1", v._1, hd)); }
  $runtime.fail();
});
const neutralApp = hd => spine => {
  if (spine.length === 0) { return hd; }
  if (hd.tag === "NeutApp") { return $BackendSemantics("NeutApp", hd._1, [...hd._2, ...spine]); }
  return $BackendSemantics("NeutApp", hd, spine);
};
const lookupLocal = v => v1 => {
  if (v1 >= 0 && v1 < v.locals.length) { return Data$dMaybe.$Maybe("Just", v.locals[v1]); }
  return Data$dMaybe.Nothing;
};
const liftString = x => $BackendSemantics("NeutLit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitString", x));
const liftOp2 = op => a => b => $BackendSemantics("NeutPrimOp", PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op2", op, a, b));
const liftOp1 = op => a => $BackendSemantics("NeutPrimOp", PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op1", op, a));
const liftNumber = x => $BackendSemantics("NeutLit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitNumber", x));
const liftInt = x => $BackendSemantics("NeutLit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitInt", x));
const liftBoolean = x => $BackendSemantics("NeutLit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitBoolean", x));
const isSimplePredicate = v => v.tag === "ExprSyntax" && (v._2.tag === "Lit" || v._2.tag === "Var" || v._2.tag === "Local" || v._2.tag === "PrimOp");
const isReference = v => v.tag === "Var" || v.tag === "Local";
const isRefExpr = v => v.tag === "Var" || v.tag === "Lit" || v.tag === "CtorSaturated" || v.tag === "Accessor" || v.tag === "Update" || v.tag === "PrimOp";
const isKnownEffect = x => x.tag === "ExprSyntax" && (x._2.tag === "PrimEffect" || x._2.tag === "UncurriedEffectApp" || x._2.tag === "EffectBind" || x._2.tag === "EffectDefer");
const isAssocPrimOp = v => {
  if (v.tag === "OpIntNum") { return v._1 === "OpAdd" || v._1 === "OpMultiply"; }
  if (v.tag === "OpNumberNum") { return v._1 === "OpAdd" || v._1 === "OpMultiply"; }
  return v.tag === "OpStringAppend";
};
const isAbs = x => x.tag === "ExprSyntax" && (x._2.tag === "Abs" || x._2.tag === "UncurriedAbs" || x._2.tag === "UncurriedEffectAbs" || x._2.tag === "EffectDefer");
const shouldInlineLet = level => a => b => {
  const $0 = (() => {
    if (a.tag === "ExprSyntax") { return a._1; }
    if (a.tag === "ExprRewrite") { return a._1; }
    $runtime.fail();
  })();
  const v2 = lookup(level)((() => {
    if (b.tag === "ExprSyntax") { return b._1.usages; }
    if (b.tag === "ExprRewrite") { return b._1.usages; }
    $runtime.fail();
  })());
  if (v2.tag === "Nothing") { return true; }
  if (v2.tag === "Just") {
    return $0.complexity === "Trivial" || v2._1.captured === "CaptureNone" && v2._1.total === 1 || (v2._1.captured === "CaptureNone" || v2._1.captured === "CaptureBranch") && (
      $0.complexity === "Trivial" || $0.complexity === "Deref"
    ) && $0.size < 5 || $0.complexity === "Deref" && v2._1.call === v2._1.total || $0.complexity === "KnownSize" && v2._1.total === 1 || a.tag === "ExprSyntax" && (
      a._2.tag === "Abs" || a._2.tag === "UncurriedAbs" || a._2.tag === "UncurriedEffectAbs" || a._2.tag === "EffectDefer"
    ) && (v2._1.total === 1 || $0.usages.tag === "Leaf" || $0.size < 16) || a.tag === "ExprSyntax" && (
      a._2.tag === "PrimEffect" || a._2.tag === "UncurriedEffectApp" || a._2.tag === "EffectBind" || a._2.tag === "EffectDefer"
    ) && v2._1.total === 1;
  }
  $runtime.fail();
};
const insertDirective = ref => acc => dir => alter(v => {
  if (v.tag === "Just") { return Data$dMaybe.$Maybe("Just", Data$dMap$dInternal.insert(ordInlineAccessor)(acc)(dir)(v._1)); }
  if (v.tag === "Nothing") { return Data$dMaybe.$Maybe("Just", Data$dMap$dInternal.$$$Map("Node", 1, 1, acc, dir, Data$dMap$dInternal.Leaf, Data$dMap$dInternal.Leaf)); }
  $runtime.fail();
})(ref);
const guardFailOver = dictFoldable => f => as => k => {
  const v = dictFoldable.foldl(v => v1 => {
    if (v.tag === "Nothing") {
      const $0 = f(v1);
      if ($0.tag === "NeutFail") { return Data$dMaybe.$Maybe("Just", $0); }
      return Data$dMaybe.Nothing;
    }
    return v;
  })(Data$dMaybe.Nothing)(as);
  if (v.tag === "Just") { return v._1; }
  if (v.tag === "Nothing") { return k(as); }
  $runtime.fail();
};
const guardFailOver1 = f => as => k => {
  const v = Data$dFoldable.foldlArray(v => v1 => {
    if (v.tag === "Nothing") {
      const $0 = f(v1);
      if ($0.tag === "NeutFail") { return Data$dMaybe.$Maybe("Just", $0); }
      return Data$dMaybe.Nothing;
    }
    return v;
  })(Data$dMaybe.Nothing)(as);
  if (v.tag === "Just") { return v._1; }
  if (v.tag === "Nothing") { return k(as); }
  $runtime.fail();
};
const guardFailOver2 = f => as => k => {
  const v = Data$dFoldable.foldlDefault(PureScript$dBackend$dOptimizer$dSyntax.foldableBackendEffect)(v => v1 => {
    if (v.tag === "Nothing") {
      const $0 = f(v1);
      if ($0.tag === "NeutFail") { return Data$dMaybe.$Maybe("Just", $0); }
      return Data$dMaybe.Nothing;
    }
    return v;
  })(Data$dMaybe.Nothing)(as);
  if (v.tag === "Just") { return v._1; }
  if (v.tag === "Nothing") { return k(as); }
  $runtime.fail();
};
const guardFailOver3 = f => as => k => {
  const v = Data$dFoldable.foldlDefault(PureScript$dBackend$dOptimizer$dCoreFn.foldableLiteral)(v => v1 => {
    if (v.tag === "Nothing") {
      const $0 = f(v1);
      if ($0.tag === "NeutFail") { return Data$dMaybe.$Maybe("Just", $0); }
      return Data$dMaybe.Nothing;
    }
    return v;
  })(Data$dMaybe.Nothing)(as);
  if (v.tag === "Just") { return v._1; }
  if (v.tag === "Nothing") { return k(as); }
  $runtime.fail();
};
const guardFail = sem => k => {
  if (sem.tag === "NeutFail") { return $BackendSemantics("NeutFail", sem._1); }
  return k(sem);
};
const foldBackendExpr = foldSyntax => foldRewrite => {
  const go = v => {
    if (v.tag === "ExprSyntax") { return foldSyntax(PureScript$dBackend$dOptimizer$dSyntax.functorBackendSyntax.map(go)(v._2)); }
    if (v.tag === "ExprRewrite") {
      return foldRewrite(v._2)((() => {
        if (v._2.tag === "RewriteInline") { return foldSyntax(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Let", v._2._1, v._2._2, go(v._2._3), go(v._2._4))); }
        if (v._2.tag === "RewriteUncurry") {
          return foldSyntax(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
            "Let",
            v._2._1,
            v._2._2,
            foldSyntax(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Abs", v._2._3, go(v._2._4))),
            go(v._2._5)
          ));
        }
        if (v._2.tag === "RewriteStop") { return foldSyntax(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Var", v._2._1)); }
        if (v._2.tag === "RewriteUnpackOp") {
          if (v._2._3.tag === "UnpackRecord") {
            return foldSyntax(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
              "Let",
              v._2._1,
              v._2._2,
              foldSyntax(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
                "Lit",
                PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitRecord", Data$dFunctor.arrayMap(m => PureScript$dBackend$dOptimizer$dCoreFn.$Prop(m._1, go(m._2)))(v._2._3._1))
              )),
              go(v._2._4)
            ));
          }
          if (v._2._3.tag === "UnpackUpdate") {
            return foldSyntax(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
              "Let",
              v._2._1,
              v._2._2,
              foldSyntax(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
                "Update",
                go(v._2._3._1),
                Data$dFunctor.arrayMap(m => PureScript$dBackend$dOptimizer$dCoreFn.$Prop(m._1, go(m._2)))(v._2._3._2)
              )),
              go(v._2._4)
            ));
          }
          if (v._2._3.tag === "UnpackArray") {
            return foldSyntax(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
              "Let",
              v._2._1,
              v._2._2,
              foldSyntax(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
                "Lit",
                PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitArray", Data$dFunctor.arrayMap(go)(v._2._3._1))
              )),
              go(v._2._4)
            ));
          }
          if (v._2._3.tag === "UnpackData") {
            return foldSyntax(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
              "Let",
              v._2._1,
              v._2._2,
              foldSyntax(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
                "CtorSaturated",
                v._2._3._1,
                v._2._3._2,
                v._2._3._3,
                v._2._3._4,
                Data$dFunctor.arrayMap(m => Data$dTuple.$Tuple(m._1, go(m._2)))(v._2._3._5)
              )),
              go(v._2._4)
            ));
          }
          $runtime.fail();
        }
        if (v._2.tag === "RewriteDistBranchesLet") {
          return foldSyntax(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
            "Let",
            v._2._1,
            v._2._2,
            foldSyntax(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
              "Branch",
              Data$dFunctor.arrayMap(m => PureScript$dBackend$dOptimizer$dSyntax.$Pair(go(m._1), go(m._2)))(v._2._3),
              go(v._2._4)
            )),
            go(v._2._5)
          ));
        }
        if (v._2.tag === "RewriteDistBranchesOp") {
          const branches$p = foldSyntax(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
            "Branch",
            Data$dFunctor.arrayMap(m => PureScript$dBackend$dOptimizer$dSyntax.$Pair(go(m._1), go(m._2)))(v._2._1),
            go(v._2._2)
          ));
          if (v._2._3.tag === "DistApp") { return foldSyntax(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("App", branches$p, Data$dFunctor.arrayMap(go)(v._2._3._1))); }
          if (v._2._3.tag === "DistUncurriedApp") {
            return foldSyntax(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("UncurriedApp", branches$p, Data$dFunctor.arrayMap(go)(v._2._3._1)));
          }
          if (v._2._3.tag === "DistAccessor") { return foldSyntax(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Accessor", branches$p, v._2._3._1)); }
          if (v._2._3.tag === "DistPrimOp1") {
            return foldSyntax(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
              "PrimOp",
              PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op1", v._2._3._1, branches$p)
            ));
          }
          if (v._2._3.tag === "DistPrimOp2L") {
            return foldSyntax(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
              "PrimOp",
              PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op2", v._2._3._1, branches$p, go(v._2._3._2))
            ));
          }
          if (v._2._3.tag === "DistPrimOp2R") {
            return foldSyntax(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
              "PrimOp",
              PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op2", v._2._3._2, go(v._2._3._1), branches$p)
            ));
          }
        }
        $runtime.fail();
      })());
    }
    $runtime.fail();
  };
  return go;
};
const freeze = init => Data$dTuple.$Tuple(
  (() => {
    if (init.tag === "ExprSyntax") { return init._1; }
    if (init.tag === "ExprRewrite") { return init._1; }
    $runtime.fail();
  })(),
  foldBackendExpr(NeutralExpr)(v => neutExpr => neutExpr)(init)
);
const floatLetWith$lazy = /* #__PURE__ */ $runtime.binding(() => {
  const go = go$a0$copy => go$a1$copy => go$a2$copy => go$a3$copy => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$a2 = go$a2$copy, go$a3 = go$a3$copy, go$c = true, go$r;
    while (go$c) {
      const f = go$a0, ident1 = go$a1, binding1 = go$a2, k1 = go$a3;
      if (binding1.tag === "SemLet") {
        go$a0 = makeLet$lazy();
        go$a1 = binding1._1;
        go$a2 = binding1._2;
        go$a3 = nextBinding2 => f(ident1)(binding1._3(nextBinding2))(k1);
        continue;
      }
      if (binding1.tag === "SemLetRec") {
        go$c = false;
        go$r = $BackendSemantics("SemLetRec", binding1._1, nextBindings => makeLet$lazy()(ident1)(binding1._2(nextBindings))(k1));
        continue;
      }
      if (binding1.tag === "NeutFail") {
        go$c = false;
        go$r = binding1;
        continue;
      }
      go$c = false;
      go$r = f(ident1)(binding1)(k1);
    }
    return go$r;
  };
  return go;
});
const makeLet$lazy = /* #__PURE__ */ $runtime.binding(() => floatLetWith$lazy()(ident => binding => k => {
  if (binding.tag === "SemRef") {
    if (binding._2.length === 0) { return k(binding); }
    return $BackendSemantics("SemLet", ident, binding, k);
  }
  if (binding.tag === "NeutLocal") { return k(binding); }
  if (binding.tag === "NeutStop") { return k(binding); }
  if (binding.tag === "NeutVar") { return k(binding); }
  return $BackendSemantics("SemLet", ident, binding, k);
}));
const floatLetWith = /* #__PURE__ */ floatLetWith$lazy();
const makeLet = /* #__PURE__ */ makeLet$lazy();
const floatLet = /* #__PURE__ */ floatLetWith(v => Data$dFunction.applyFlipped)(Data$dMaybe.Nothing);
const makeEffectBind$lazy = /* #__PURE__ */ $runtime.binding(() => {
  const go = go$a0$copy => go$a1$copy => go$a2$copy => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$a2 = go$a2$copy, go$c = true, go$r;
    while (go$c) {
      const ident1 = go$a0, binding1 = go$a1, k1 = go$a2;
      if (binding1.tag === "SemLet") {
        go$c = false;
        go$r = makeLet(binding1._1)(binding1._2)(nextBinding2 => makeEffectBind$lazy()(ident1)(binding1._3(nextBinding2))(k1));
        continue;
      }
      if (binding1.tag === "SemEffectBind") {
        go$a0 = binding1._1;
        go$a1 = binding1._2;
        go$a2 = nextBinding2 => makeEffectBind$lazy()(ident1)(binding1._3(nextBinding2))(k1);
        continue;
      }
      if (binding1.tag === "SemEffectDefer") {
        go$c = false;
        go$r = $BackendSemantics("SemEffectDefer", floatLet(binding1._1)(nextBinding2 => makeEffectBind$lazy()(ident1)(nextBinding2)(k1)));
        continue;
      }
      go$c = false;
      go$r = floatLet(binding1)(nextBinding2 => $BackendSemantics("SemEffectBind", ident1, nextBinding2, k1));
    }
    return go$r;
  };
  return go;
});
const makeEffectBind = /* #__PURE__ */ makeEffectBind$lazy();
const evalUpdate = lhs => props => floatLet(lhs)(v => {
  if (v.tag === "NeutLit") {
    if (v._1.tag === "LitRecord") {
      return $BackendSemantics(
        "NeutLit",
        PureScript$dBackend$dOptimizer$dCoreFn.$Literal(
          "LitRecord",
          Data$dFunctor.arrayMap(Data$dArray$dNonEmpty.head)(Data$dArray.groupAllBy(x => y => Data$dOrd.ordString.compare(x._1)(y._1))([...props, ...v._1._1]))
        )
      );
    }
    return $BackendSemantics("NeutUpdate", v, props);
  }
  if (v.tag === "NeutUpdate") {
    return $BackendSemantics(
      "NeutUpdate",
      v._1,
      Data$dFunctor.arrayMap(Data$dArray$dNonEmpty.head)(Data$dArray.groupAllBy(x => y => Data$dOrd.ordString.compare(x._1)(y._1))([...props, ...v._2]))
    );
  }
  return $BackendSemantics("NeutUpdate", v, props);
});
const evalUncurriedBeta = fn => mk => spine => {
  const go = v => v1 => {
    if (v.tag === "MkFnNext") {
      if (v1.tag === "Cons") {
        if (v1._1.tag === "NeutFail") { return $BackendSemantics("NeutFail", v1._1._1); }
        const $0 = v1._2;
        return makeLet(Data$dMaybe.Nothing)(v1._1)(nextArg => go(v._2(nextArg))($0));
      }
      return Partial._crashWith("Uncurried function applied to too few arguments");
    }
    if (v.tag === "MkFnApplied") {
      if (v1.tag === "Nil") { return v._1; }
      return fn(v._1)(toUnfoldable1(v1));
    }
    $runtime.fail();
  };
  return go(mk)(fromFoldable(spine));
};
const evalPrimOpOrdNumber = op => x => y => {
  if (op === "OpEq") { return x === y; }
  if (op === "OpNotEq") { return x !== y; }
  if (op === "OpGt") { return x > y; }
  if (op === "OpGte") { return x >= y; }
  if (op === "OpLt") { return x < y; }
  if (op === "OpLte") { return x <= y; }
  $runtime.fail();
};
const evalPrimOpOrd = dictOrd => {
  const Eq0 = dictOrd.Eq0();
  return op => x => y => {
    if (op === "OpEq") { return Eq0.eq(x)(y); }
    if (op === "OpNotEq") { return !Eq0.eq(x)(y); }
    if (op === "OpGt") { return dictOrd.compare(x)(y) === "GT"; }
    if (op === "OpGte") { return dictOrd.compare(x)(y) !== "LT"; }
    if (op === "OpLt") { return dictOrd.compare(x)(y) === "LT"; }
    if (op === "OpLte") { return dictOrd.compare(x)(y) !== "GT"; }
    $runtime.fail();
  };
};
const evalPrimOpOrd1 = /* #__PURE__ */ evalPrimOpOrd(Data$dOrd.ordString);
const evalPrimOpOrd2 = /* #__PURE__ */ evalPrimOpOrd(Data$dOrd.ordInt);
const evalPrimOpOrd3 = /* #__PURE__ */ evalPrimOpOrd(Data$dOrd.ordChar);
const evalPrimOpOrd4 = /* #__PURE__ */ evalPrimOpOrd(Data$dOrd.ordBoolean);
const evalPrimOpNot = v => {
  if (v.tag === "Op1") {
    if (v._1.tag === "OpBooleanNot") { return v._2; }
    return $BackendSemantics(
      "NeutPrimOp",
      PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator(
        "Op1",
        PureScript$dBackend$dOptimizer$dSyntax.OpBooleanNot,
        $BackendSemantics("NeutPrimOp", PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op1", v._1, v._2))
      )
    );
  }
  if (v.tag === "Op2") {
    if (v._1.tag === "OpIntOrd") {
      return $BackendSemantics(
        "NeutPrimOp",
        PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator(
          "Op2",
          PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2(
            "OpIntOrd",
            (() => {
              if (v._1._1 === "OpEq") { return PureScript$dBackend$dOptimizer$dSyntax.OpNotEq; }
              if (v._1._1 === "OpNotEq") { return PureScript$dBackend$dOptimizer$dSyntax.OpEq; }
              if (v._1._1 === "OpLt") { return PureScript$dBackend$dOptimizer$dSyntax.OpGte; }
              if (v._1._1 === "OpLte") { return PureScript$dBackend$dOptimizer$dSyntax.OpGt; }
              if (v._1._1 === "OpGt") { return PureScript$dBackend$dOptimizer$dSyntax.OpLte; }
              if (v._1._1 === "OpGte") { return PureScript$dBackend$dOptimizer$dSyntax.OpLt; }
              $runtime.fail();
            })()
          ),
          v._2,
          v._3
        )
      );
    }
    if (v._1.tag === "OpNumberOrd") {
      return $BackendSemantics(
        "NeutPrimOp",
        PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator(
          "Op2",
          PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2(
            "OpNumberOrd",
            (() => {
              if (v._1._1 === "OpEq") { return PureScript$dBackend$dOptimizer$dSyntax.OpNotEq; }
              if (v._1._1 === "OpNotEq") { return PureScript$dBackend$dOptimizer$dSyntax.OpEq; }
              if (v._1._1 === "OpLt") { return PureScript$dBackend$dOptimizer$dSyntax.OpGte; }
              if (v._1._1 === "OpLte") { return PureScript$dBackend$dOptimizer$dSyntax.OpGt; }
              if (v._1._1 === "OpGt") { return PureScript$dBackend$dOptimizer$dSyntax.OpLte; }
              if (v._1._1 === "OpGte") { return PureScript$dBackend$dOptimizer$dSyntax.OpLt; }
              $runtime.fail();
            })()
          ),
          v._2,
          v._3
        )
      );
    }
    if (v._1.tag === "OpStringOrd") {
      return $BackendSemantics(
        "NeutPrimOp",
        PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator(
          "Op2",
          PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2(
            "OpStringOrd",
            (() => {
              if (v._1._1 === "OpEq") { return PureScript$dBackend$dOptimizer$dSyntax.OpNotEq; }
              if (v._1._1 === "OpNotEq") { return PureScript$dBackend$dOptimizer$dSyntax.OpEq; }
              if (v._1._1 === "OpLt") { return PureScript$dBackend$dOptimizer$dSyntax.OpGte; }
              if (v._1._1 === "OpLte") { return PureScript$dBackend$dOptimizer$dSyntax.OpGt; }
              if (v._1._1 === "OpGt") { return PureScript$dBackend$dOptimizer$dSyntax.OpLte; }
              if (v._1._1 === "OpGte") { return PureScript$dBackend$dOptimizer$dSyntax.OpLt; }
              $runtime.fail();
            })()
          ),
          v._2,
          v._3
        )
      );
    }
    if (v._1.tag === "OpCharOrd") {
      return $BackendSemantics(
        "NeutPrimOp",
        PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator(
          "Op2",
          PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2(
            "OpCharOrd",
            (() => {
              if (v._1._1 === "OpEq") { return PureScript$dBackend$dOptimizer$dSyntax.OpNotEq; }
              if (v._1._1 === "OpNotEq") { return PureScript$dBackend$dOptimizer$dSyntax.OpEq; }
              if (v._1._1 === "OpLt") { return PureScript$dBackend$dOptimizer$dSyntax.OpGte; }
              if (v._1._1 === "OpLte") { return PureScript$dBackend$dOptimizer$dSyntax.OpGt; }
              if (v._1._1 === "OpGt") { return PureScript$dBackend$dOptimizer$dSyntax.OpLte; }
              if (v._1._1 === "OpGte") { return PureScript$dBackend$dOptimizer$dSyntax.OpLt; }
              $runtime.fail();
            })()
          ),
          v._2,
          v._3
        )
      );
    }
    if (v._1.tag === "OpBooleanOrd") {
      return $BackendSemantics(
        "NeutPrimOp",
        PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator(
          "Op2",
          PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2(
            "OpBooleanOrd",
            (() => {
              if (v._1._1 === "OpEq") { return PureScript$dBackend$dOptimizer$dSyntax.OpNotEq; }
              if (v._1._1 === "OpNotEq") { return PureScript$dBackend$dOptimizer$dSyntax.OpEq; }
              if (v._1._1 === "OpLt") { return PureScript$dBackend$dOptimizer$dSyntax.OpGte; }
              if (v._1._1 === "OpLte") { return PureScript$dBackend$dOptimizer$dSyntax.OpGt; }
              if (v._1._1 === "OpGt") { return PureScript$dBackend$dOptimizer$dSyntax.OpLte; }
              if (v._1._1 === "OpGte") { return PureScript$dBackend$dOptimizer$dSyntax.OpLt; }
              $runtime.fail();
            })()
          ),
          v._2,
          v._3
        )
      );
    }
    return $BackendSemantics(
      "NeutPrimOp",
      PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator(
        "Op1",
        PureScript$dBackend$dOptimizer$dSyntax.OpBooleanNot,
        $BackendSemantics("NeutPrimOp", PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op2", v._1, v._2, v._3))
      )
    );
  }
  $runtime.fail();
};
const evalEvalRef = v => {
  if (v.tag === "EvalExtern") { return $BackendSemantics("NeutVar", v._1); }
  if (v.tag === "EvalLocal") { return $BackendSemantics("NeutLocal", v._1, v._2); }
  $runtime.fail();
};
const $$eval = dict => dict.eval;
const evalPair = dictEval => env => v => {
  const $0 = v._1;
  const $1 = v._2;
  return $SemConditional(Data$dLazy.defer(v1 => dictEval.eval(env)($0)), Data$dLazy.defer(v1 => dictEval.eval(env)($1)));
};
const effectfully = v => {
  if (v.effect) { return v; }
  return {...v, effect: true};
};
const deref = v => {
  if (v.tag === "SemRef") { return Data$dLazy.force(v._3); }
  return v;
};
const evalBranches = v => initConds => initDef => {
  const go = go$a0$copy => go$a1$copy => go$a2$copy => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$a2 = go$a2$copy, go$c = true, go$r;
    while (go$c) {
      const acc = go$a0, conds = go$a1, def = go$a2;
      const v1 = Data$dArray.unconsImpl(v$1 => Data$dMaybe.Nothing, x => xs => Data$dMaybe.$Maybe("Just", {head: x, tail: xs}), conds);
      if (v1.tag === "Just") {
        const $0 = Data$dLazy.force(v1._1.head._1);
        const v2 = $0.tag === "SemRef" ? Data$dLazy.force($0._3) : $0;
        if (v2.tag === "NeutLit") {
          if (v2._1.tag === "LitBoolean") {
            if (v2._1._1) {
              go$a0 = acc;
              go$a1 = [];
              go$a2 = v1._1.head._2;
              continue;
            }
            go$a0 = acc;
            go$a1 = v1._1.tail;
            go$a2 = def;
            continue;
          }
          go$a0 = Data$dArray.snoc(acc)(v1._1.head);
          go$a1 = v1._1.tail;
          go$a2 = def;
          continue;
        }
        if (v2.tag === "NeutFail") {
          const $1 = v2._1;
          go$a0 = acc;
          go$a1 = [];
          go$a2 = Data$dLazy.defer(v3 => $BackendSemantics("NeutFail", $1));
          continue;
        }
        go$a0 = Data$dArray.snoc(acc)(v1._1.head);
        go$a1 = v1._1.tail;
        go$a2 = def;
        continue;
      }
      if (v1.tag === "Nothing") {
        if (acc.length > 0) {
          go$c = false;
          go$r = $BackendSemantics("SemBranch", acc, def);
          continue;
        }
        go$c = false;
        go$r = Data$dLazy.force(def);
        continue;
      }
      $runtime.fail();
    }
    return go$r;
  };
  return go([])(initConds)(initDef);
};
const evalPrimOpNumInt = op => x => y => {
  const $0 = x.tag === "SemRef" ? Data$dLazy.force(x._3) : x;
  if ($0.tag === "NeutLit" && $0._1.tag === "LitInt") {
    const $1 = y.tag === "SemRef" ? Data$dLazy.force(y._3) : y;
    if ($1.tag === "NeutLit" && $1._1.tag === "LitInt") {
      if (op === "OpAdd") {
        const res = $0._1._1 + $1._1._1 | 0;
        if ($1._1._1 > 0 && res < $0._1._1 || $1._1._1 < 0 && res > $0._1._1) { return Data$dMaybe.Nothing; }
        return Data$dMaybe.$Maybe("Just", $BackendSemantics("NeutLit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitInt", res)));
      }
      if (op === "OpMultiply") {
        const res = $0._1._1 * $1._1._1 | 0;
        if ($0._1._1 !== $runtime.intDiv(res, $1._1._1)) { return Data$dMaybe.Nothing; }
        return Data$dMaybe.$Maybe("Just", $BackendSemantics("NeutLit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitInt", res)));
      }
      if (op === "OpSubtract") {
        const res = $0._1._1 - $1._1._1 | 0;
        if ($1._1._1 > 0 && res > $0._1._1 || $1._1._1 < 0 && res < $0._1._1) { return Data$dMaybe.Nothing; }
        return Data$dMaybe.$Maybe("Just", $BackendSemantics("NeutLit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitInt", res)));
      }
      if (op === "OpDivide") {
        return Data$dMaybe.$Maybe("Just", $BackendSemantics("NeutLit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitInt", $runtime.intDiv($0._1._1, $1._1._1))));
      }
      $runtime.fail();
    }
  }
  return Data$dMaybe.Nothing;
};
const evalPrimOpNumNumber = op => x => y => {
  const $0 = x.tag === "SemRef" ? Data$dLazy.force(x._3) : x;
  if ($0.tag === "NeutLit" && $0._1.tag === "LitNumber") {
    const $1 = y.tag === "SemRef" ? Data$dLazy.force(y._3) : y;
    if ($1.tag === "NeutLit" && $1._1.tag === "LitNumber") {
      return Data$dMaybe.$Maybe(
        "Just",
        $BackendSemantics(
          "NeutLit",
          PureScript$dBackend$dOptimizer$dCoreFn.$Literal(
            "LitNumber",
            (() => {
              if (op === "OpAdd") { return $0._1._1 + $1._1._1; }
              if (op === "OpMultiply") { return $0._1._1 * $1._1._1; }
              if (op === "OpSubtract") { return $0._1._1 - $1._1._1; }
              if (op === "OpDivide") { return $0._1._1 / $1._1._1; }
              $runtime.fail();
            })()
          )
        )
      );
    }
  }
  return Data$dMaybe.Nothing;
};
const evalRefSpine = env => ref => spine => sem => v => {
  if (v.tag === "ExternApp") {
    return neutralSpine((() => {
      if (ref.tag === "EvalExtern") { return $BackendSemantics("NeutVar", ref._1); }
      if (ref.tag === "EvalLocal") { return $BackendSemantics("NeutLocal", ref._1, ref._2); }
      $runtime.fail();
    })())(spine);
  }
  if (v.tag === "ExternUncurriedApp") {
    return neutralSpine((() => {
      if (ref.tag === "EvalExtern") { return $BackendSemantics("NeutVar", ref._1); }
      if (ref.tag === "EvalLocal") { return $BackendSemantics("NeutLocal", ref._1, ref._2); }
      $runtime.fail();
    })())(spine);
  }
  if (v.tag === "ExternAccessor") { return evalAccessor(env)(Data$dLazy.force(sem))(v._1); }
  if (v.tag === "ExternPrimOp") { return evalPrimOp(env)(PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op1", v._1, Data$dLazy.force(sem))); }
  $runtime.fail();
};
const evalRef = v => ref => spine => last => sem => {
  const spine$p = last.tag === "ExternApp" ? Data$dFoldable.foldlArray(snocApp)(spine)(last._1) : Data$dArray.snoc(spine)(last);
  const v1 = v2 => $BackendSemantics(
    "SemRef",
    ref,
    spine$p,
    Data$dLazy.defer(v3 => {
      const $0 = evalRefSpine(v)(ref)(spine$p)(sem)(last);
      if ($0.tag === "SemRef") { return Data$dLazy.force($0._3); }
      return $0;
    })
  );
  if (ref.tag === "EvalExtern") {
    const $0 = v.evalExternSpine(v)(ref._1)(spine$p);
    if ($0.tag === "Just") { return $0._1; }
  }
  return v1(true);
};
const evalPrimOp = env => v => {
  if (v.tag === "Op1") {
    const $0 = v._1;
    const $1 = v._2;
    const v1 = v2 => {
      if ($0.tag === "OpBooleanNot" && $1.tag === "NeutPrimOp") { return evalPrimOpNot($1._1); }
      const v5 = v6 => {
        const v7 = v8 => {
          const v9 = v10 => {
            const v11 = v12 => {
              const v13 = v14 => {
                if ($1.tag === "SemRef") { return evalRef(env)($1._1)($1._2)($ExternSpine("ExternPrimOp", $0))($1._3); }
                if ($1.tag === "NeutFail") { return $BackendSemantics("NeutFail", $1._1); }
                return floatLet($1)((() => {
                  const $2 = PureScript$dBackend$dOptimizer$dSyntax.Op1($0);
                  return x => $BackendSemantics("NeutPrimOp", $2(x));
                })());
              };
              if ($0.tag === "OpNumberNegate") {
                const $2 = $1.tag === "SemRef" ? Data$dLazy.force($1._3) : $1;
                if ($2.tag === "NeutLit" && $2._1.tag === "LitNumber") {
                  return $BackendSemantics("NeutLit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitNumber", -$2._1._1));
                }
              }
              return v13(true);
            };
            if ($0.tag === "OpIntNegate") {
              const $2 = $1.tag === "SemRef" ? Data$dLazy.force($1._3) : $1;
              if ($2.tag === "NeutLit" && $2._1.tag === "LitInt") { return $BackendSemantics("NeutLit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitInt", -$2._1._1)); }
            }
            return v11(true);
          };
          if ($0.tag === "OpArrayLength") {
            const $2 = $1.tag === "SemRef" ? Data$dLazy.force($1._3) : $1;
            if ($2.tag === "NeutLit" && $2._1.tag === "LitArray") {
              return $BackendSemantics("NeutLit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitInt", $2._1._1.length));
            }
          }
          return v9(true);
        };
        if ($0.tag === "OpIsTag") {
          const $2 = $1.tag === "SemRef" ? Data$dLazy.force($1._3) : $1;
          if ($2.tag === "NeutData") {
            return $BackendSemantics(
              "NeutLit",
              PureScript$dBackend$dOptimizer$dCoreFn.$Literal(
                "LitBoolean",
                ($0._1._1.tag === "Nothing" ? $2._1._1.tag === "Nothing" : $0._1._1.tag === "Just" && $2._1._1.tag === "Just" && $0._1._1._1 === $2._1._1._1) && $0._1._2 === $2._1._2
              )
            );
          }
        }
        return v7(true);
      };
      if ($0.tag === "OpIntBitNot") {
        const $2 = $1.tag === "SemRef" ? Data$dLazy.force($1._3) : $1;
        if ($2.tag === "NeutLit" && $2._1.tag === "LitInt") { return $BackendSemantics("NeutLit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitInt", ~$2._1._1)); }
      }
      return v5(true);
    };
    if ($0.tag === "OpBooleanNot") {
      const $2 = $1.tag === "SemRef" ? Data$dLazy.force($1._3) : $1;
      if ($2.tag === "NeutLit" && $2._1.tag === "LitBoolean") { return $BackendSemantics("NeutLit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitBoolean", !$2._1._1)); }
    }
    return v1(true);
  }
  if (v.tag === "Op2") {
    const $0 = v._1;
    const $1 = v._2;
    const $2 = v._3;
    const v1 = v2 => {
      const v3 = v4 => {
        const v5 = v6 => {
          const v7 = v8 => {
            const v9 = v10 => {
              const v11 = v12 => {
                const v13 = v14 => {
                  const v15 = v16 => {
                    const v17 = v18 => {
                      const v19 = v20 => {
                        const v21 = v22 => {
                          const v23 = v24 => {
                            const v25 = v26 => {
                              const v27 = v28 => {
                                const v29 = v30 => {
                                  const v31 = v32 => {
                                    const v33 = v34 => {
                                      const v35 = v36 => {
                                        const v37 = v38 => {
                                          const v39 = v40 => {
                                            const v41 = v42 => {
                                              const v43 = v44 => {
                                                const v45 = v46 => {
                                                  const v47 = v48 => {
                                                    const v49 = v50 => {
                                                      if (
                                                        $0.tag === "OpStringAppend" && $1.tag === "NeutLit" && $1._1.tag === "LitString" && $2.tag === "NeutLit" && $2._1.tag === "LitString"
                                                      ) {
                                                        return $BackendSemantics("NeutLit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitString", $1._1._1 + $2._1._1));
                                                      }
                                                      if ($0.tag === "OpArrayIndex" && $2.tag === "NeutLit" && $2._1.tag === "LitInt") {
                                                        return evalAccessor(env)($1)(PureScript$dBackend$dOptimizer$dSyntax.$BackendAccessor("GetIndex", $2._1._1));
                                                      }
                                                      if ($0.tag === "OpBooleanAnd") {
                                                        if ($1.tag === "NeutFail") { return $BackendSemantics("NeutFail", $1._1); }
                                                        if ($2.tag === "NeutFail") { return $BackendSemantics("NeutFail", $2._1); }
                                                        return $BackendSemantics("NeutPrimOp", PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op2", $0, $1, $2));
                                                      }
                                                      if ($0.tag === "OpBooleanOr") {
                                                        if ($1.tag === "NeutFail") { return $BackendSemantics("NeutFail", $1._1); }
                                                        if ($2.tag === "NeutFail") { return $BackendSemantics("NeutFail", $2._1); }
                                                        return $BackendSemantics("NeutPrimOp", PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op2", $0, $1, $2));
                                                      }
                                                      if ($1.tag === "NeutFail") { return $BackendSemantics("NeutFail", $1._1); }
                                                      if ($2.tag === "NeutFail") { return $BackendSemantics("NeutFail", $2._1); }
                                                      return floatLet($1)(x$p => floatLet($2)(y$p => {
                                                        if (
                                                          (() => {
                                                            if ($0.tag === "OpIntNum") { return $0._1 === "OpAdd" || $0._1 === "OpMultiply"; }
                                                            if ($0.tag === "OpNumberNum") { return $0._1 === "OpAdd" || $0._1 === "OpMultiply"; }
                                                            return $0.tag === "OpStringAppend";
                                                          })()
                                                        ) {
                                                          return evalAssocOp(env)(Data$dEither.$Either("Right", $0))(x$p)(y$p);
                                                        }
                                                        return $BackendSemantics("NeutPrimOp", PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op2", $0, x$p, y$p));
                                                      }));
                                                    };
                                                    if ($0.tag === "OpStringOrd") {
                                                      const $3 = $1.tag === "SemRef" ? Data$dLazy.force($1._3) : $1;
                                                      if ($3.tag === "NeutLit" && $3._1.tag === "LitString") {
                                                        const $4 = $2.tag === "SemRef" ? Data$dLazy.force($2._3) : $2;
                                                        if ($4.tag === "NeutLit" && $4._1.tag === "LitString") {
                                                          return $BackendSemantics(
                                                            "NeutLit",
                                                            PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitBoolean", evalPrimOpOrd1($0._1)($3._1._1)($4._1._1))
                                                          );
                                                        }
                                                      }
                                                    }
                                                    return v49(true);
                                                  };
                                                  if ($0.tag === "OpNumberOrd") {
                                                    const $3 = $1.tag === "SemRef" ? Data$dLazy.force($1._3) : $1;
                                                    if ($3.tag === "NeutLit" && $3._1.tag === "LitNumber") {
                                                      const $4 = $2.tag === "SemRef" ? Data$dLazy.force($2._3) : $2;
                                                      if ($4.tag === "NeutLit" && $4._1.tag === "LitNumber") {
                                                        return $BackendSemantics(
                                                          "NeutLit",
                                                          PureScript$dBackend$dOptimizer$dCoreFn.$Literal(
                                                            "LitBoolean",
                                                            (() => {
                                                              if ($0._1 === "OpEq") { return $3._1._1 === $4._1._1; }
                                                              if ($0._1 === "OpNotEq") { return $3._1._1 !== $4._1._1; }
                                                              if ($0._1 === "OpGt") { return $3._1._1 > $4._1._1; }
                                                              if ($0._1 === "OpGte") { return $3._1._1 >= $4._1._1; }
                                                              if ($0._1 === "OpLt") { return $3._1._1 < $4._1._1; }
                                                              if ($0._1 === "OpLte") { return $3._1._1 <= $4._1._1; }
                                                              $runtime.fail();
                                                            })()
                                                          )
                                                        );
                                                      }
                                                    }
                                                  }
                                                  return v47(true);
                                                };
                                                if ($0.tag === "OpNumberNum") {
                                                  const $3 = evalPrimOpNumNumber($0._1)($1)($2);
                                                  if ($3.tag === "Just") { return $3._1; }
                                                }
                                                return v45(true);
                                              };
                                              if ($0.tag === "OpNumberNum" && $0._1 === "OpSubtract") {
                                                const $3 = $1.tag === "SemRef" ? Data$dLazy.force($1._3) : $1;
                                                if ($3.tag === "NeutLit" && $3._1.tag === "LitNumber" && $3._1._1 === 0.0) {
                                                  return evalPrimOp(env)(PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator(
                                                    "Op1",
                                                    PureScript$dBackend$dOptimizer$dSyntax.OpNumberNegate,
                                                    $2
                                                  ));
                                                }
                                              }
                                              return v43(true);
                                            };
                                            if ($0.tag === "OpIntOrd") {
                                              const $3 = $1.tag === "SemRef" ? Data$dLazy.force($1._3) : $1;
                                              if ($3.tag === "NeutLit" && $3._1.tag === "LitInt") {
                                                const $4 = $2.tag === "SemRef" ? Data$dLazy.force($2._3) : $2;
                                                if ($4.tag === "NeutLit" && $4._1.tag === "LitInt") {
                                                  return $BackendSemantics(
                                                    "NeutLit",
                                                    PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitBoolean", evalPrimOpOrd2($0._1)($3._1._1)($4._1._1))
                                                  );
                                                }
                                              }
                                            }
                                            return v41(true);
                                          };
                                          if ($0.tag === "OpIntNum") {
                                            const $3 = evalPrimOpNumInt($0._1)($1)($2);
                                            if ($3.tag === "Just") { return $3._1; }
                                          }
                                          return v39(true);
                                        };
                                        if ($0.tag === "OpIntNum" && $0._1 === "OpSubtract") {
                                          const $3 = $1.tag === "SemRef" ? Data$dLazy.force($1._3) : $1;
                                          if ($3.tag === "NeutLit" && $3._1.tag === "LitInt" && $3._1._1 === 0) {
                                            return evalPrimOp(env)(PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator(
                                              "Op1",
                                              PureScript$dBackend$dOptimizer$dSyntax.OpIntNegate,
                                              $2
                                            ));
                                          }
                                        }
                                        return v37(true);
                                      };
                                      if ($0.tag === "OpIntBitZeroFillShiftRight") {
                                        const $3 = $1.tag === "SemRef" ? Data$dLazy.force($1._3) : $1;
                                        if ($3.tag === "NeutLit" && $3._1.tag === "LitInt") {
                                          const $4 = $2.tag === "SemRef" ? Data$dLazy.force($2._3) : $2;
                                          if ($4.tag === "NeutLit" && $4._1.tag === "LitInt") {
                                            return $BackendSemantics("NeutLit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitInt", $3._1._1 >>> $4._1._1));
                                          }
                                        }
                                      }
                                      return v35(true);
                                    };
                                    if ($0.tag === "OpIntBitXor") {
                                      const $3 = $1.tag === "SemRef" ? Data$dLazy.force($1._3) : $1;
                                      if ($3.tag === "NeutLit" && $3._1.tag === "LitInt") {
                                        const $4 = $2.tag === "SemRef" ? Data$dLazy.force($2._3) : $2;
                                        if ($4.tag === "NeutLit" && $4._1.tag === "LitInt") {
                                          return $BackendSemantics("NeutLit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitInt", $3._1._1 ^ $4._1._1));
                                        }
                                      }
                                    }
                                    return v33(true);
                                  };
                                  if ($0.tag === "OpIntBitShiftRight") {
                                    const $3 = $1.tag === "SemRef" ? Data$dLazy.force($1._3) : $1;
                                    if ($3.tag === "NeutLit" && $3._1.tag === "LitInt") {
                                      const $4 = $2.tag === "SemRef" ? Data$dLazy.force($2._3) : $2;
                                      if ($4.tag === "NeutLit" && $4._1.tag === "LitInt") {
                                        return $BackendSemantics("NeutLit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitInt", $3._1._1 >> $4._1._1));
                                      }
                                    }
                                  }
                                  return v31(true);
                                };
                                if ($0.tag === "OpIntBitShiftLeft") {
                                  const $3 = $1.tag === "SemRef" ? Data$dLazy.force($1._3) : $1;
                                  if ($3.tag === "NeutLit" && $3._1.tag === "LitInt") {
                                    const $4 = $2.tag === "SemRef" ? Data$dLazy.force($2._3) : $2;
                                    if ($4.tag === "NeutLit" && $4._1.tag === "LitInt") {
                                      return $BackendSemantics("NeutLit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitInt", $3._1._1 << $4._1._1));
                                    }
                                  }
                                }
                                return v29(true);
                              };
                              if ($0.tag === "OpIntBitOr") {
                                const $3 = $1.tag === "SemRef" ? Data$dLazy.force($1._3) : $1;
                                if ($3.tag === "NeutLit" && $3._1.tag === "LitInt") {
                                  const $4 = $2.tag === "SemRef" ? Data$dLazy.force($2._3) : $2;
                                  if ($4.tag === "NeutLit" && $4._1.tag === "LitInt") {
                                    return $BackendSemantics("NeutLit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitInt", $3._1._1 | $4._1._1));
                                  }
                                }
                              }
                              return v27(true);
                            };
                            if ($0.tag === "OpIntBitAnd") {
                              const $3 = $1.tag === "SemRef" ? Data$dLazy.force($1._3) : $1;
                              if ($3.tag === "NeutLit" && $3._1.tag === "LitInt") {
                                const $4 = $2.tag === "SemRef" ? Data$dLazy.force($2._3) : $2;
                                if ($4.tag === "NeutLit" && $4._1.tag === "LitInt") {
                                  return $BackendSemantics("NeutLit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitInt", $3._1._1 & $4._1._1));
                                }
                              }
                            }
                            return v25(true);
                          };
                          if ($0.tag === "OpCharOrd") {
                            const $3 = $1.tag === "SemRef" ? Data$dLazy.force($1._3) : $1;
                            if ($3.tag === "NeutLit" && $3._1.tag === "LitChar") {
                              const $4 = $2.tag === "SemRef" ? Data$dLazy.force($2._3) : $2;
                              if ($4.tag === "NeutLit" && $4._1.tag === "LitChar") {
                                return $BackendSemantics("NeutLit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitBoolean", evalPrimOpOrd3($0._1)($3._1._1)($4._1._1)));
                              }
                            }
                          }
                          return v23(true);
                        };
                        if ($0.tag === "OpBooleanOrd") {
                          const $3 = $1.tag === "SemRef" ? Data$dLazy.force($1._3) : $1;
                          if ($3.tag === "NeutLit" && $3._1.tag === "LitBoolean") {
                            const $4 = $2.tag === "SemRef" ? Data$dLazy.force($2._3) : $2;
                            if ($4.tag === "NeutLit" && $4._1.tag === "LitBoolean") {
                              return $BackendSemantics("NeutLit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitBoolean", evalPrimOpOrd4($0._1)($3._1._1)($4._1._1)));
                            }
                          }
                        }
                        return v21(true);
                      };
                      if ($0.tag === "OpBooleanOrd" && $0._1 === "OpEq") {
                        const $3 = $2.tag === "SemRef" ? Data$dLazy.force($2._3) : $2;
                        if ($3.tag === "NeutLit" && $3._1.tag === "LitBoolean") {
                          if ($3._1._1) { return $1; }
                          return evalPrimOp(env)(PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op1", PureScript$dBackend$dOptimizer$dSyntax.OpBooleanNot, $1));
                        }
                      }
                      return v19(true);
                    };
                    if ($0.tag === "OpBooleanOrd" && $0._1 === "OpEq") {
                      const $3 = $1.tag === "SemRef" ? Data$dLazy.force($1._3) : $1;
                      if ($3.tag === "NeutLit" && $3._1.tag === "LitBoolean") {
                        if ($3._1._1) { return $2; }
                        return evalPrimOp(env)(PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op1", PureScript$dBackend$dOptimizer$dSyntax.OpBooleanNot, $2));
                      }
                    }
                    return v17(true);
                  };
                  if ($0.tag === "OpBooleanOr") {
                    const $3 = $2.tag === "SemRef" ? Data$dLazy.force($2._3) : $2;
                    if ($3.tag === "NeutLit" && $3._1.tag === "LitBoolean" && $3._1._1) { return $2; }
                  }
                  return v15(true);
                };
                if ($0.tag === "OpBooleanOr") {
                  const $3 = $1.tag === "SemRef" ? Data$dLazy.force($1._3) : $1;
                  if ($3.tag === "NeutLit" && $3._1.tag === "LitBoolean" && $3._1._1) { return $1; }
                }
                return v13(true);
              };
              if ($0.tag === "OpBooleanOr") {
                const $3 = $2.tag === "SemRef" ? Data$dLazy.force($2._3) : $2;
                if ($3.tag === "NeutLit" && $3._1.tag === "LitBoolean" && !$3._1._1) { return $1; }
              }
              return v11(true);
            };
            if ($0.tag === "OpBooleanOr") {
              const $3 = $1.tag === "SemRef" ? Data$dLazy.force($1._3) : $1;
              if ($3.tag === "NeutLit" && $3._1.tag === "LitBoolean" && !$3._1._1) { return $2; }
            }
            return v9(true);
          };
          if ($0.tag === "OpBooleanAnd") {
            const $3 = $2.tag === "SemRef" ? Data$dLazy.force($2._3) : $2;
            if ($3.tag === "NeutLit" && $3._1.tag === "LitBoolean" && $3._1._1) { return $1; }
          }
          return v7(true);
        };
        if ($0.tag === "OpBooleanAnd") {
          const $3 = $1.tag === "SemRef" ? Data$dLazy.force($1._3) : $1;
          if ($3.tag === "NeutLit" && $3._1.tag === "LitBoolean" && $3._1._1) { return $2; }
        }
        return v5(true);
      };
      if ($0.tag === "OpBooleanAnd") {
        const $3 = $2.tag === "SemRef" ? Data$dLazy.force($2._3) : $2;
        if ($3.tag === "NeutLit" && $3._1.tag === "LitBoolean" && !$3._1._1) { return $2; }
      }
      return v3(true);
    };
    if ($0.tag === "OpBooleanAnd") {
      const $3 = $1.tag === "SemRef" ? Data$dLazy.force($1._3) : $1;
      if ($3.tag === "NeutLit" && $3._1.tag === "LitBoolean" && !$3._1._1) { return $1; }
    }
    return v1(true);
  }
  $runtime.fail();
};
const evalAssocOp$p = v => op => a => b => {
  if (op.tag === "Left") {
    const v1 = v.evalExternSpine(v)(op._1)([$ExternSpine("ExternApp", [a, b])]);
    if (v1.tag === "Just") { return v1._1; }
    if (v1.tag === "Nothing") { return $BackendSemantics("SemAssocOp", op, [a, b]); }
    $runtime.fail();
  }
  if (op.tag === "Right") { return evalPrimOp(v)(PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op2", op._1, a, b)); }
  $runtime.fail();
};
const evalAssocOp = env => op1 => v => v1 => {
  if (v.tag === "SemAssocOp" && v1.tag === "SemAssocOp" && eq16(op1)(v._1) && eq16(v._1)(v1._1)) {
    const v3 = evalAssocOp$p(env)(op1)((() => {
      const $0 = v._2.length - 1 | 0;
      if ($0 >= 0 && $0 < v._2.length) { return v._2[$0]; }
      $runtime.fail();
    })())((() => {
      if (0 < v1._2.length) { return v1._2[0]; }
      $runtime.fail();
    })());
    if (v3.tag === "SemAssocOp" && eq16(v1._1)(v3._1)) {
      return $BackendSemantics(
        "SemAssocOp",
        op1,
        (() => {
          const $0 = Data$dArray.unconsImpl(v$1 => Data$dMaybe.Nothing, v$1 => xs => Data$dMaybe.$Maybe("Just", xs), v1._2);
          return [
            ...(() => {
              if (v._2.length === 0) { $runtime.fail(); }
              return Data$dArray.sliceImpl(0, v._2.length - 1 | 0, v._2);
            })(),
            ...v3._2,
            ...(() => {
              if ($0.tag === "Just") { return $0._1; }
              $runtime.fail();
            })()
          ];
        })()
      );
    }
    return $BackendSemantics(
      "SemAssocOp",
      op1,
      (() => {
        const $0 = Data$dArray.unconsImpl(v$1 => Data$dMaybe.Nothing, v$1 => xs => Data$dMaybe.$Maybe("Just", xs), v1._2);
        return [
          ...(() => {
            if (v._2.length === 0) { $runtime.fail(); }
            return Data$dArray.sliceImpl(0, v._2.length - 1 | 0, v._2);
          })(),
          v3,
          ...(() => {
            if ($0.tag === "Just") { return $0._1; }
            $runtime.fail();
          })()
        ];
      })()
    );
  }
  const $0 = (as, b, op2) => {
    const v4 = evalAssocOp$p(env)(op1)((() => {
      const $0 = as.length - 1 | 0;
      if ($0 >= 0 && $0 < as.length) { return as[$0]; }
      $runtime.fail();
    })())(b);
    if (v4.tag === "SemAssocOp" && eq16(op2)(v4._1)) {
      return $BackendSemantics(
        "SemAssocOp",
        op1,
        [
          ...(() => {
            if (as.length === 0) { $runtime.fail(); }
            return Data$dArray.sliceImpl(0, as.length - 1 | 0, as);
          })(),
          ...v4._2
        ]
      );
    }
    return $BackendSemantics(
      "SemAssocOp",
      op1,
      Data$dArray.snoc((() => {
        if (as.length === 0) { $runtime.fail(); }
        return Data$dArray.sliceImpl(0, as.length - 1 | 0, as);
      })())(v4)
    );
  };
  if (v1.tag === "SemAssocOp" && eq16(op1)(v1._1)) {
    const v4 = evalAssocOp$p(env)(op1)(v)((() => {
      if (0 < v1._2.length) { return v1._2[0]; }
      $runtime.fail();
    })());
    if (v4.tag === "SemAssocOp" && eq16(v1._1)(v4._1)) {
      return $BackendSemantics(
        "SemAssocOp",
        op1,
        (() => {
          const $1 = Data$dArray.unconsImpl(v$1 => Data$dMaybe.Nothing, v$1 => xs => Data$dMaybe.$Maybe("Just", xs), v1._2);
          return [
            ...v4._2,
            ...(() => {
              if ($1.tag === "Just") { return $1._1; }
              $runtime.fail();
            })()
          ];
        })()
      );
    }
    return $BackendSemantics(
      "SemAssocOp",
      op1,
      (() => {
        const $1 = Data$dArray.unconsImpl(v$1 => Data$dMaybe.Nothing, v$1 => xs => Data$dMaybe.$Maybe("Just", xs), v1._2);
        return [
          v4,
          ...(() => {
            if ($1.tag === "Just") { return $1._1; }
            $runtime.fail();
          })()
        ];
      })()
    );
  }
  if (v.tag === "SemAssocOp" && eq16(op1)(v._1)) { return $0(v._2, v1, v._1); }
  return $BackendSemantics("SemAssocOp", op1, [v, v1]);
};
const evalAccessor = env => lhs => accessor => floatLet(lhs)(v => {
  if (v.tag === "SemRef") { return evalRef(env)(v._1)(v._2)($ExternSpine("ExternAccessor", accessor))(v._3); }
  const v1 = v2 => {
    if (v.tag === "NeutUpdate" && accessor.tag === "GetProp") {
      const $0 = accessor._1;
      const v4 = Data$dArray.findMapImpl(
        Data$dMaybe.Nothing,
        Data$dMaybe.isJust,
        v5 => {
          if (v5._1 === $0) { return Data$dMaybe.$Maybe("Just", v5._2); }
          return Data$dMaybe.Nothing;
        },
        v._2
      );
      if (v4.tag === "Just") { return v4._1; }
      if (v4.tag === "Nothing") { return evalAccessor(env)(v._1)(accessor); }
      $runtime.fail();
    }
    if (v.tag === "NeutLit" && v._1.tag === "LitArray" && accessor.tag === "GetIndex" && accessor._1 >= 0 && accessor._1 < v._1._1.length) { return v._1._1[accessor._1]; }
    if (v.tag === "NeutData" && accessor.tag === "GetCtorField" && accessor._6 >= 0 && accessor._6 < v._5.length) { return v._5[accessor._6]._2; }
    if (v.tag === "NeutFail") { return $BackendSemantics("NeutFail", v._1); }
    return $BackendSemantics("NeutAccessor", v, accessor);
  };
  if (v.tag === "NeutLit" && v._1.tag === "LitRecord" && accessor.tag === "GetProp") {
    const $0 = accessor._1;
    const $1 = Data$dArray.findMapImpl(
      Data$dMaybe.Nothing,
      Data$dMaybe.isJust,
      v2 => {
        if (v2._1 === $0) { return Data$dMaybe.$Maybe("Just", v2._2); }
        return Data$dMaybe.Nothing;
      },
      v._1._1
    );
    if ($1.tag === "Just") { return $1._1; }
  }
  return v1(true);
});
const caseString = v => {
  if (v.tag === "NeutLit" && v._1.tag === "LitString") { return Data$dMaybe.$Maybe("Just", v._1._1); }
  return Data$dMaybe.Nothing;
};
const caseNumber = v => {
  if (v.tag === "NeutLit" && v._1.tag === "LitNumber") { return Data$dMaybe.$Maybe("Just", v._1._1); }
  return Data$dMaybe.Nothing;
};
const caseInt = v => {
  if (v.tag === "NeutLit" && v._1.tag === "LitInt") { return Data$dMaybe.$Maybe("Just", v._1._1); }
  return Data$dMaybe.Nothing;
};
const buildStop = v => stop => $BackendExpr("ExprRewrite", v.analyze(v)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Var", stop)), $BackendRewrite("RewriteStop", stop));
const buildDefault = v => expr => $BackendExpr("ExprSyntax", v.analyze(v)(expr), expr);
const build = ctx => v => {
  const $0 = () => {
    const v1 = v2 => {
      const v3 = v4 => {
        const v5 = v6 => {
          const v7 = v8 => {
            const v9 = v10 => {
              const v11 = v12 => {
                const v13 = v14 => {
                  const v15 = v16 => {
                    const v17 = v18 => {
                      if (v.tag === "Accessor" && v._1.tag === "ExprSyntax" && v._1._2.tag === "Branch") {
                        return $BackendExpr(
                          "ExprRewrite",
                          {...v._1._1, rewrite: true, size: v._1._1.size + 1 | 0},
                          $BackendRewrite("RewriteDistBranchesOp", v._1._2._1, v._1._2._2, $DistOp("DistAccessor", v._2))
                        );
                      }
                      if (v.tag === "PrimOp" && v._1.tag === "Op1" && v._1._2.tag === "ExprSyntax" && v._1._2._2.tag === "Branch") {
                        return $BackendExpr(
                          "ExprRewrite",
                          {...v._1._2._1, rewrite: true, size: v._1._2._1.size + 1 | 0},
                          $BackendRewrite("RewriteDistBranchesOp", v._1._2._2._1, v._1._2._2._2, $DistOp("DistPrimOp1", v._1._1))
                        );
                      }
                      const v23 = v24 => {
                        const v25 = v26 => {
                          if (v.tag === "EffectBind") {
                            if (v._3.tag === "ExprSyntax") {
                              if (v._3._2.tag === "EffectPure") {
                                return build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
                                  "EffectDefer",
                                  build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Let", v._1, v._2, v._3._2._1, v._4))
                                ));
                              }
                              if (v._3._2.tag === "EffectDefer") {
                                return build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("EffectBind", v._1, v._2, v._3._2._1, v._4));
                              }
                              if (v._4.tag === "ExprSyntax") {
                                if (v._4._2.tag === "EffectDefer") {
                                  return build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("EffectBind", v._1, v._2, v._3, v._4._2._1));
                                }
                                if (v._4._2.tag === "EffectPure" && v._4._2._1.tag === "ExprSyntax" && v._4._2._1._2.tag === "Local" && v._2 === v._4._2._1._2._2) { return v._3; }
                              }
                              return $BackendExpr("ExprSyntax", ctx.analyze(ctx)(v), v);
                            }
                            if (v._4.tag === "ExprSyntax") {
                              if (v._4._2.tag === "EffectDefer") {
                                return build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("EffectBind", v._1, v._2, v._3, v._4._2._1));
                              }
                              if (v._4._2.tag === "EffectPure" && v._4._2._1.tag === "ExprSyntax" && v._4._2._1._2.tag === "Local" && v._2 === v._4._2._1._2._2) { return v._3; }
                            }
                            return $BackendExpr("ExprSyntax", ctx.analyze(ctx)(v), v);
                          }
                          if (v.tag === "EffectDefer") {
                            if (v._1.tag === "ExprSyntax" && v._1._2.tag === "EffectDefer") { return v._1; }
                            return $BackendExpr("ExprSyntax", ctx.analyze(ctx)(v), v);
                          }
                          if (
                            v.tag === "PrimOp" && v._1.tag === "Op1" && v._1._1.tag === "OpBooleanNot" && v._1._2.tag === "ExprSyntax" && v._1._2._2.tag === "PrimOp" && v._1._2._2._1.tag === "Op1" && v._1._2._2._1._1.tag === "OpBooleanNot"
                          ) {
                            return v._1._2._2._1._2;
                          }
                          return $BackendExpr("ExprSyntax", ctx.analyze(ctx)(v), v);
                        };
                        if (v.tag === "PrimOp" && v._1.tag === "Op2" && v._1._3.tag === "ExprSyntax" && v._1._3._2.tag === "Branch") {
                          const $0 = shouldDistributeBranchPrimOp2R(v._1._3._1)(v._1._3._2._1)(v._1._3._2._2)(v._1._2)(v._1._1);
                          if ($0.tag === "Just") { return $0._1; }
                        }
                        return v25(true);
                      };
                      if (v.tag === "PrimOp" && v._1.tag === "Op2" && v._1._2.tag === "ExprSyntax" && v._1._2._2.tag === "Branch") {
                        const $0 = shouldDistributeBranchPrimOp2L(v._1._2._1)(v._1._2._2._1)(v._1._2._2._2)(v._1._1)(v._1._3);
                        if ($0.tag === "Just") { return $0._1; }
                      }
                      return v23(true);
                    };
                    if (v.tag === "UncurriedApp" && v._1.tag === "ExprSyntax" && v._1._2.tag === "Branch") {
                      const $0 = shouldDistributeBranchUncurriedApps(v._1._1)(v._1._2._1)(v._1._2._2)(v._2);
                      if ($0.tag === "Just") { return $0._1; }
                    }
                    return v17(true);
                  };
                  if (v.tag === "App" && v._1.tag === "ExprSyntax" && v._1._2.tag === "Branch") {
                    const $0 = shouldDistributeBranchApps(v._1._1)(v._1._2._1)(v._1._2._2)(v._2);
                    if ($0.tag === "Just") { return $0._1; }
                  }
                  return v15(true);
                };
                if (v.tag === "Let") {
                  const $0 = shouldEtaReduce(v._2)(v._3)(v._4);
                  if ($0.tag === "Just") { return $0._1; }
                }
                return v13(true);
              };
              if (v.tag === "Let") {
                const $0 = shouldDistributeBranches(v._1)(v._2)(v._3)(v._4);
                if ($0.tag === "Just") { return $0._1; }
              }
              return v11(true);
            };
            if (v.tag === "Let") {
              const $0 = shouldUnpackArray(v._1)(v._2)(v._3)(v._4);
              if ($0.tag === "Just") { return $0._1; }
            }
            return v9(true);
          };
          if (v.tag === "Let") {
            const $0 = shouldUnpackCtor(v._1)(v._2)(v._3)(v._4);
            if ($0.tag === "Just") { return $0._1; }
          }
          return v7(true);
        };
        if (v.tag === "Let") {
          const $0 = shouldUnpackUpdate(v._1)(v._2)(v._3)(v._4);
          if ($0.tag === "Just") { return $0._1; }
        }
        return v5(true);
      };
      if (v.tag === "Let") {
        const $0 = shouldUnpackRecord(v._1)(v._2)(v._3)(v._4);
        if ($0.tag === "Just") { return $0._1; }
      }
      return v3(true);
    };
    if (v.tag === "Let") {
      const $0 = shouldUncurryAbs(v._1)(v._2)(v._3)(v._4);
      if ($0.tag === "Just") { return $0._1; }
    }
    return v1(true);
  };
  if (v.tag === "App") {
    if (v._1.tag === "ExprSyntax" && v._1._2.tag === "App") {
      return build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("App", v._1._2._1, [...v._1._2._2, ...v._2]));
    }
    return $0();
  }
  if (v.tag === "Abs") {
    if (v._2.tag === "ExprSyntax" && v._2._2.tag === "Abs") {
      return build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Abs", [...v._1, ...v._2._2._1], v._2._2._2));
    }
    return $0();
  }
  if (v.tag === "Let" && shouldInlineLet(v._2)(v._3)(v._4)) { return rewriteInline(v._1)(v._2)(v._3)(v._4); }
  return $0();
};
const simplifyCondBoolean = ctx => v => v1 => {
  if (v._2.tag === "ExprSyntax" && v._2._2.tag === "Lit" && v._2._2._1.tag === "LitBoolean") {
    if (v1.tag === "ExprSyntax" && v1._2.tag === "Lit" && v1._2._1.tag === "LitBoolean") {
      if (v._2._2._1._1 === v1._2._1._1) { return Data$dMaybe.$Maybe("Just", v._2); }
      if (v._2._2._1._1 && !v1._2._1._1) { return Data$dMaybe.$Maybe("Just", v._1); }
      if (!v._2._2._1._1 && v1._2._1._1) {
        return Data$dMaybe.$Maybe(
          "Just",
          build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
            "PrimOp",
            PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op1", PureScript$dBackend$dOptimizer$dSyntax.OpBooleanNot, v._1)
          ))
        );
      }
      if (v._2._2._1._1 && v1.tag === "ExprSyntax" && (v1._2.tag === "Lit" || v1._2.tag === "Var" || v1._2.tag === "Local" || v1._2.tag === "PrimOp")) {
        return Data$dMaybe.$Maybe(
          "Just",
          build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
            "PrimOp",
            PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op2", PureScript$dBackend$dOptimizer$dSyntax.OpBooleanOr, v._1, v1)
          ))
        );
      }
      if (!v1._2._1._1) {
        return Data$dMaybe.$Maybe(
          "Just",
          build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
            "PrimOp",
            PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op2", PureScript$dBackend$dOptimizer$dSyntax.OpBooleanAnd, v._1, v._2)
          ))
        );
      }
      return Data$dMaybe.Nothing;
    }
    if (v._2._2._1._1 && v1.tag === "ExprSyntax" && (v1._2.tag === "Lit" || v1._2.tag === "Var" || v1._2.tag === "Local" || v1._2.tag === "PrimOp")) {
      return Data$dMaybe.$Maybe(
        "Just",
        build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
          "PrimOp",
          PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op2", PureScript$dBackend$dOptimizer$dSyntax.OpBooleanOr, v._1, v1)
        ))
      );
    }
    return Data$dMaybe.Nothing;
  }
  if (v1.tag === "ExprSyntax" && v1._2.tag === "Lit" && v1._2._1.tag === "LitBoolean" && !v1._2._1._1) {
    return Data$dMaybe.$Maybe(
      "Just",
      build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
        "PrimOp",
        PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op2", PureScript$dBackend$dOptimizer$dSyntax.OpBooleanAnd, v._1, v._2)
      ))
    );
  }
  return Data$dMaybe.Nothing;
};
const simplifyCondRedundantElse = ctx => v => v1 => {
  if (v1.tag === "ExprSyntax" && v1._2.tag === "Branch") {
    const $0 = (() => {
      if (0 < v1._2._1.length) { return v1._2._1[0]; }
      $runtime.fail();
    })();
    if ($0._1.tag === "ExprSyntax" && $0._1._2.tag === "PrimOp" && $0._1._2._1.tag === "Op1" && $0._1._2._1._1.tag === "OpBooleanNot" && eqBackendExpr.eq(v._1)($0._1._2._1._2)) {
      return Data$dMaybe.$Maybe("Just", buildBranchCond(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$Pair(v._1, v._2))($0._2));
    }
  }
  return Data$dMaybe.Nothing;
};
const simplifyCondLiftAnd = ctx => pair => def1 => {
  if (pair._2.tag === "ExprSyntax" && pair._2._2.tag === "Branch" && pair._2._2._1.length === 1 && eqBackendExpr.eq(def1)(pair._2._2._2)) {
    return Data$dMaybe.$Maybe(
      "Just",
      buildBranchCond(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$Pair(
        build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
          "PrimOp",
          PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op2", PureScript$dBackend$dOptimizer$dSyntax.OpBooleanAnd, pair._1, pair._2._2._1[0]._1)
        )),
        pair._2._2._1[0]._2
      ))(def1)
    );
  }
  return Data$dMaybe.Nothing;
};
const buildBranchCond = ctx => pair => def => {
  const $0 = simplifyCondIsTag(ctx)(pair)(def);
  if ($0.tag === "Just") { return $0._1; }
  const $1 = simplifyCondBoolean(ctx)(pair)(def);
  if ($1.tag === "Just") { return $1._1; }
  const $2 = simplifyCondLiftAnd(ctx)(pair)(def);
  if ($2.tag === "Just") { return $2._1; }
  const $3 = simplifyCondRedundantElse(ctx)(pair)(def);
  if ($3.tag === "Just") { return $3._1; }
  if (def.tag === "ExprSyntax" && def._2.tag === "Branch") { return build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Branch", [pair, ...def._2._1], def._2._2)); }
  return build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Branch", [pair], def));
};
const quote$lazy = /* #__PURE__ */ $runtime.binding(() => {
  const go = go$a0$copy => go$a1$copy => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
    while (go$c) {
      const ctx = go$a0, v = go$a1;
      if (v.tag === "SemLet") {
        const $0 = v._2;
        go$c = false;
        go$r = build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
          "Let",
          v._1,
          ctx.currentLevel,
          quote$lazy()(ctx.effect ? {...ctx, effect: false} : ctx)($0),
          quote$lazy()({...ctx, currentLevel: ctx.currentLevel + 1 | 0})(v._3($BackendSemantics(
            "SemRef",
            $EvalRef("EvalLocal", v._1, ctx.currentLevel),
            [],
            Data$dLazy.defer(v2 => {
              if ($0.tag === "SemRef") { return Data$dLazy.force($0._3); }
              return $0;
            })
          )))
        ));
        continue;
      }
      if (v.tag === "SemLetRec") {
        const $0 = ctx.currentLevel;
        const $1 = {...ctx, currentLevel: ctx.currentLevel + 1 | 0};
        const neutBindings = Data$dFunctor.arrayMap(v2 => {
          const $2 = v2._1;
          return Data$dTuple.$Tuple($2, Data$dLazy.defer(v3 => $BackendSemantics("NeutLocal", Data$dMaybe.$Maybe("Just", $2), $0)));
        })(v._1);
        go$c = false;
        go$r = build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
          "LetRec",
          $0,
          Data$dFunctor.arrayMap(m => Data$dTuple.$Tuple(m._1, quote$lazy()($1.effect ? {...$1, effect: false} : $1)(m._2(neutBindings))))(v._1),
          quote$lazy()($1)(v._2(neutBindings))
        ));
        continue;
      }
      if (v.tag === "SemEffectBind") {
        const ctx$p = ctx.effect ? ctx : {...ctx, effect: true};
        go$c = false;
        go$r = build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
          "EffectBind",
          v._1,
          ctx$p.currentLevel,
          quote$lazy()(ctx$p)(v._2),
          quote$lazy()({...ctx$p, currentLevel: ctx$p.currentLevel + 1 | 0})(v._3($BackendSemantics("NeutLocal", v._1, ctx$p.currentLevel)))
        ));
        continue;
      }
      if (v.tag === "SemEffectPure") {
        go$c = false;
        go$r = build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("EffectPure", quote$lazy()(ctx.effect ? {...ctx, effect: false} : ctx)(v._1)));
        continue;
      }
      if (v.tag === "SemEffectDefer") {
        go$c = false;
        go$r = build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("EffectDefer", quote$lazy()(ctx.effect ? ctx : {...ctx, effect: true})(v._1)));
        continue;
      }
      if (v.tag === "SemBranch") {
        const ctx$p = ctx.effect ? {...ctx, effect: false} : ctx;
        go$c = false;
        go$r = Data$dFoldable.foldrArray(buildBranchCond(ctx))(quote$lazy()(ctx)(Data$dLazy.force(v._2)))(Data$dFunctor.arrayMap(v1 => PureScript$dBackend$dOptimizer$dSyntax.$Pair(
          quote$lazy()(ctx$p)(Data$dLazy.force(v1._1)),
          quote$lazy()(ctx)(Data$dLazy.force(v1._2))
        ))(v._1));
        continue;
      }
      if (v.tag === "SemRef") {
        if (v._1.tag === "EvalExtern") {
          go$a0 = ctx;
          go$a1 = neutralSpine($BackendSemantics("NeutVar", v._1._1))(v._2);
          continue;
        }
        if (v._1.tag === "EvalLocal") {
          go$a0 = ctx;
          go$a1 = neutralSpine($BackendSemantics("NeutLocal", v._1._1, v._1._2))(v._2);
          continue;
        }
        $runtime.fail();
      }
      if (v.tag === "SemLam") {
        const $0 = {...ctx, currentLevel: ctx.currentLevel + 1 | 0};
        go$c = false;
        go$r = build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
          "Abs",
          [Data$dTuple.$Tuple(v._1, ctx.currentLevel)],
          quote$lazy()($0.effect ? {...$0, effect: false} : $0)(v._2($BackendSemantics("NeutLocal", v._1, ctx.currentLevel)))
        ));
        continue;
      }
      if (v.tag === "SemMkFn") {
        const loop = loop$a0$copy => loop$a1$copy => loop$a2$copy => {
          let loop$a0 = loop$a0$copy, loop$a1 = loop$a1$copy, loop$a2 = loop$a2$copy, loop$c = true, loop$r;
          while (loop$c) {
            const ctx$p = loop$a0, idents = loop$a1, v1 = loop$a2;
            if (v1.tag === "MkFnNext") {
              loop$a0 = {...ctx$p, currentLevel: ctx$p.currentLevel + 1 | 0};
              loop$a1 = Data$dArray.snoc(idents)(Data$dTuple.$Tuple(v1._1, ctx$p.currentLevel));
              loop$a2 = v1._2($BackendSemantics("NeutLocal", v1._1, ctx$p.currentLevel));
              continue;
            }
            if (v1.tag === "MkFnApplied") {
              loop$c = false;
              loop$r = build(ctx$p)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
                "UncurriedAbs",
                idents,
                quote$lazy()(ctx$p.effect ? {...ctx$p, effect: false} : ctx$p)(v1._1)
              ));
              continue;
            }
            $runtime.fail();
          }
          return loop$r;
        };
        go$c = false;
        go$r = loop(ctx)([])(v._1);
        continue;
      }
      if (v.tag === "SemMkEffectFn") {
        const loop = loop$a0$copy => loop$a1$copy => loop$a2$copy => {
          let loop$a0 = loop$a0$copy, loop$a1 = loop$a1$copy, loop$a2 = loop$a2$copy, loop$c = true, loop$r;
          while (loop$c) {
            const ctx$p = loop$a0, idents = loop$a1, v1 = loop$a2;
            if (v1.tag === "MkFnNext") {
              loop$a0 = {...ctx$p, currentLevel: ctx$p.currentLevel + 1 | 0};
              loop$a1 = Data$dArray.snoc(idents)(Data$dTuple.$Tuple(v1._1, ctx$p.currentLevel));
              loop$a2 = v1._2($BackendSemantics("NeutLocal", v1._1, ctx$p.currentLevel));
              continue;
            }
            if (v1.tag === "MkFnApplied") {
              loop$c = false;
              loop$r = build(ctx$p)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
                "UncurriedEffectAbs",
                idents,
                quote$lazy()(ctx$p.effect ? {...ctx$p, effect: false} : ctx$p)(v1._1)
              ));
              continue;
            }
            $runtime.fail();
          }
          return loop$r;
        };
        go$c = false;
        go$r = loop(ctx)([])(v._1);
        continue;
      }
      if (v.tag === "SemAssocOp") {
        const $0 = v._1;
        const $1 = v._2;
        const len = $1.length;
        const go$1 = go$1$a0$copy => go$1$a1$copy => {
          let go$1$a0 = go$1$a0$copy, go$1$a1 = go$1$a1$copy, go$1$c = true, go$1$r;
          while (go$1$c) {
            const ix = go$1$a0, acc = go$1$a1;
            if (ix === len) {
              go$1$c = false;
              go$1$r = acc;
              continue;
            }
            go$1$a0 = ix + 1 | 0;
            go$1$a1 = (() => {
              const $2 = $1[ix];
              if ($0.tag === "Left") {
                return build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
                  "App",
                  build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Var", $0._1)),
                  [acc, quote$lazy()(ctx)($2)]
                ));
              }
              if ($0.tag === "Right") {
                return build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
                  "PrimOp",
                  PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op2", $0._1, acc, quote$lazy()(ctx)($2))
                ));
              }
              $runtime.fail();
            })();
          }
          return go$1$r;
        };
        go$c = false;
        go$r = go$1(1)(quote$lazy()(ctx)((() => {
          if (0 < $1.length) { return $1[0]; }
          $runtime.fail();
        })()));
        continue;
      }
      if (v.tag === "NeutLocal") {
        go$c = false;
        go$r = build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Local", v._1, v._2));
        continue;
      }
      if (v.tag === "NeutVar") {
        go$c = false;
        go$r = build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Var", v._1));
        continue;
      }
      if (v.tag === "NeutStop") {
        go$c = false;
        go$r = $BackendExpr("ExprRewrite", ctx.analyze(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Var", v._1)), $BackendRewrite("RewriteStop", v._1));
        continue;
      }
      if (v.tag === "NeutData") {
        go$c = false;
        go$r = build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
          "CtorSaturated",
          v._1,
          v._2,
          v._3,
          v._4,
          Data$dFunctor.arrayMap((() => {
            const $0 = quote$lazy()(ctx);
            return m => Data$dTuple.$Tuple(m._1, $0(m._2));
          })())(v._5)
        ));
        continue;
      }
      if (v.tag === "NeutCtorDef") {
        go$c = false;
        go$r = build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("CtorDef", v._2, v._3, v._4, v._5));
        continue;
      }
      if (v.tag === "NeutUncurriedApp") {
        const ctx$p = ctx.effect ? {...ctx, effect: false} : ctx;
        go$c = false;
        go$r = build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("UncurriedApp", quote$lazy()(ctx$p)(v._1), Data$dFunctor.arrayMap(quote$lazy()(ctx$p))(v._2)));
        continue;
      }
      if (v.tag === "NeutUncurriedEffectApp") {
        const ctx$p = ctx.effect ? {...ctx, effect: false} : ctx;
        go$c = false;
        go$r = build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("UncurriedEffectApp", quote$lazy()(ctx$p)(v._1), Data$dFunctor.arrayMap(quote$lazy()(ctx$p))(v._2)));
        continue;
      }
      if (v.tag === "NeutApp") {
        const ctx$p = ctx.effect ? {...ctx, effect: false} : ctx;
        const hd$p = quote$lazy()(ctx$p)(v._1);
        const $0 = Data$dFunctor.arrayMap(quote$lazy()(ctx$p))(v._2);
        if ($0.length > 0) {
          go$c = false;
          go$r = build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("App", hd$p, $0));
          continue;
        }
        go$c = false;
        go$r = hd$p;
        continue;
      }
      if (v.tag === "NeutAccessor") {
        go$c = false;
        go$r = build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Accessor", quote$lazy()(ctx)(v._1), v._2));
        continue;
      }
      if (v.tag === "NeutUpdate") {
        go$c = false;
        go$r = build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
          "Update",
          quote$lazy()(ctx)(v._1),
          Data$dFunctor.arrayMap((() => {
            const $0 = quote$lazy()(ctx);
            return m => PureScript$dBackend$dOptimizer$dCoreFn.$Prop(m._1, $0(m._2));
          })())(v._2)
        ));
        continue;
      }
      if (v.tag === "NeutLit") {
        go$c = false;
        go$r = build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Lit", PureScript$dBackend$dOptimizer$dCoreFn.functorLiteral.map(quote$lazy()(ctx))(v._1)));
        continue;
      }
      if (v.tag === "NeutPrimOp") {
        go$c = false;
        go$r = build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
          "PrimOp",
          (() => {
            const $0 = quote$lazy()(ctx);
            if (v._1.tag === "Op1") { return PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op1", v._1._1, $0(v._1._2)); }
            if (v._1.tag === "Op2") { return PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op2", v._1._1, $0(v._1._2), $0(v._1._3)); }
            $runtime.fail();
          })()
        ));
        continue;
      }
      if (v.tag === "NeutPrimEffect") {
        go$c = false;
        go$r = build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax(
          "PrimEffect",
          (() => {
            const $0 = quote$lazy()(ctx.effect ? {...ctx, effect: false} : ctx);
            if (v._1.tag === "EffectRefNew") { return PureScript$dBackend$dOptimizer$dSyntax.$BackendEffect("EffectRefNew", $0(v._1._1)); }
            if (v._1.tag === "EffectRefRead") { return PureScript$dBackend$dOptimizer$dSyntax.$BackendEffect("EffectRefRead", $0(v._1._1)); }
            if (v._1.tag === "EffectRefWrite") { return PureScript$dBackend$dOptimizer$dSyntax.$BackendEffect("EffectRefWrite", $0(v._1._1), $0(v._1._2)); }
            $runtime.fail();
          })()
        ));
        continue;
      }
      if (v.tag === "NeutPrimUndefined") {
        go$c = false;
        go$r = build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.PrimUndefined);
        continue;
      }
      if (v.tag === "NeutFail") {
        go$c = false;
        go$r = build(ctx)(PureScript$dBackend$dOptimizer$dSyntax.$BackendSyntax("Fail", v._1));
        continue;
      }
      $runtime.fail();
    }
    return go$r;
  };
  return go;
});
const quote = /* #__PURE__ */ quote$lazy();
const bindLocal = v => sem => ({...v, locals: Data$dArray.snoc(v.locals)(sem)});
const evalApp = env => hd => spine => {
  const go = env$p => v => v1 => {
    const $0 = (args, ident, k, val) => $BackendSemantics(
      "SemLet",
      ident,
      val,
      nextVal => makeLet(Data$dMaybe.Nothing)(k(nextVal))(nextFn => go({
        ...env$p,
        locals: Data$dArray.snoc(Data$dArray.snoc(env$p.locals)($LocalBinding("One", nextVal)))($LocalBinding("One", nextFn))
      })(nextFn)(args))
    );
    const $1 = (args, k, vals) => $BackendSemantics(
      "SemLetRec",
      vals,
      nextVals => makeLet(Data$dMaybe.Nothing)(k(nextVals))(nextFn => go({
        ...env$p,
        locals: Data$dArray.snoc(Data$dArray.snoc(env$p.locals)($LocalBinding("Group", nextVals)))($LocalBinding("One", nextFn))
      })(nextFn)(args))
    );
    if (v1.tag === "Cons") {
      if (v1._1.tag === "NeutFail") { return $BackendSemantics("NeutFail", v1._1._1); }
      if (v.tag === "NeutFail") { return $BackendSemantics("NeutFail", v._1); }
      if (v.tag === "SemLam") {
        const $2 = v1._2;
        return makeLet(Data$dMaybe.Nothing)(v1._1)(nextArg => go(env$p)(v._2(nextArg))($2));
      }
      if (v.tag === "SemRef") { return go(env$p)(evalRef(env$p)(v._1)(v._2)($ExternSpine("ExternApp", [v1._1]))(v._3))(v1._2); }
      if (v.tag === "SemLet") { return $0(v1, v._1, v._3, v._2); }
      if (v.tag === "SemLetRec") { return $1(v1, v._2, v._1); }
      return $BackendSemantics("NeutApp", v, toUnfoldable1(v1));
    }
    if (v.tag === "NeutFail") { return $BackendSemantics("NeutFail", v._1); }
    if (v.tag === "SemLet") { return $0(v1, v._1, v._3, v._2); }
    if (v.tag === "SemLetRec") { return $1(v1, v._2, v._1); }
    if (v1.tag === "Nil") { return v; }
    return $BackendSemantics("NeutApp", v, toUnfoldable1(v1));
  };
  return go(env)(hd)(fromFoldable(spine));
};
const evalMkFn = env => n => sem => {
  if (n === 0) { return $MkFn("MkFnApplied", sem); }
  if (sem.tag === "SemLam") {
    return $MkFn(
      "MkFnNext",
      sem._1,
      (() => {
        const $0 = evalMkFn(env)(n - 1 | 0);
        return x => $0(sem._2(x));
      })()
    );
  }
  return $MkFn(
    "MkFnNext",
    Data$dMaybe.Nothing,
    nextArg => {
      const env$p = {...env, locals: Data$dArray.snoc(env.locals)($LocalBinding("One", nextArg))};
      return evalMkFn(env$p)(n - 1 | 0)(evalApp(env$p)(sem)([nextArg]));
    }
  );
};
const evalUncurriedApp = env => hd => spine => {
  if (hd.tag === "SemMkFn") { return evalUncurriedBeta(NeutUncurriedApp)(hd._1)(spine); }
  if (hd.tag === "SemRef") {
    const $0 = hd._1;
    const $1 = hd._3;
    const $2 = hd._2;
    return guardFailOver1(identity)(spine)(spine$p => evalRef(env)($0)($2)($ExternSpine("ExternUncurriedApp", spine$p))($1));
  }
  if (hd.tag === "SemLet") {
    return $BackendSemantics(
      "SemLet",
      hd._1,
      hd._2,
      nextVal => makeLet(Data$dMaybe.Nothing)(hd._3(nextVal))(nextFn => evalUncurriedApp({
        ...env,
        locals: Data$dArray.snoc(Data$dArray.snoc(env.locals)($LocalBinding("One", nextVal)))($LocalBinding("One", nextFn))
      })(nextFn)(spine))
    );
  }
  if (hd.tag === "NeutFail") { return $BackendSemantics("NeutFail", hd._1); }
  return guardFailOver1(identity)(spine)(NeutUncurriedApp(hd));
};
const evalSpine = env => Data$dFoldable.foldlArray(hd => v => {
  if (v.tag === "ExternApp") { return evalApp(env)(hd)(v._1); }
  if (v.tag === "ExternUncurriedApp") { return evalUncurriedApp(env)(hd)(v._1); }
  if (v.tag === "ExternAccessor") { return evalAccessor(env)(hd)(v._1); }
  if (v.tag === "ExternPrimOp") { return evalPrimOp(env)(PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op1", v._1, hd)); }
  $runtime.fail();
});
const mkUncurriedAppRewrite = env => hd => {
  const go = acc => n => {
    if (n === 0) { return evalUncurriedApp(env)(hd)(acc); }
    return $BackendSemantics("SemLam", Data$dMaybe.Nothing, arg => go(Data$dArray.snoc(acc)(arg))(n - 1 | 0));
  };
  return go([]);
};
const evalUncurriedEffectApp = env => hd => spine => {
  if (hd.tag === "SemMkEffectFn") { return evalUncurriedBeta(NeutUncurriedEffectApp)(hd._1)(spine); }
  if (hd.tag === "SemLet") {
    return $BackendSemantics(
      "SemLet",
      hd._1,
      hd._2,
      nextVal => makeLet(Data$dMaybe.Nothing)(hd._3(nextVal))(nextFn => evalUncurriedEffectApp({
        ...env,
        locals: Data$dArray.snoc(Data$dArray.snoc(env.locals)($LocalBinding("One", nextVal)))($LocalBinding("One", nextFn))
      })(nextFn)(spine))
    );
  }
  if (hd.tag === "NeutFail") { return $BackendSemantics("NeutFail", hd._1); }
  return guardFailOver1(identity)(spine)(NeutUncurriedEffectApp(hd));
};
const mkFnFromArgs = dictEval => env => args => body => $BackendSemantics(
  "SemMkFn",
  Data$dFoldable.foldrArray(v => {
    const $0 = v._1;
    return next => env$p => $MkFn("MkFnNext", $0, x => next({...env$p, locals: Data$dArray.snoc(env$p.locals)($LocalBinding("One", x))}));
  })(x => $MkFn("MkFnApplied", dictEval.eval(x)(body)))(args)(env)
);
const evalBackendSyntax = dictEval => (
  {
    eval: v => v1 => {
      if (v1.tag === "Var") {
        const v2 = v.evalExternSpine(v)(v1._1)([]);
        if (v2.tag === "Just") { return v2._1; }
        if (v2.tag === "Nothing") {
          return $BackendSemantics(
            "SemRef",
            $EvalRef("EvalExtern", v1._1),
            [],
            Data$dLazy.defer(v3 => {
              const v4 = v.evalExternRef(v)(v1._1);
              if (v4.tag === "Just") {
                if (v4._1.tag === "SemRef") { return Data$dLazy.force(v4._1._3); }
                return v4._1;
              }
              if (v4.tag === "Nothing") { return $BackendSemantics("NeutVar", v1._1); }
              $runtime.fail();
            })
          );
        }
        $runtime.fail();
      }
      if (v1.tag === "Local") {
        if (v1._2 >= 0 && v1._2 < v.locals.length) {
          if (v.locals[v1._2].tag === "One") { return v.locals[v1._2]._1; }
          if (v.locals[v1._2].tag === "Group") {
            const $0 = v.locals[v1._2]._1;
            const $1 = (() => {
              if (v1._1.tag === "Just") { return lookup1(v1._1._1)($0); }
              if (v1._1.tag === "Nothing") { return Data$dMaybe.Nothing; }
              $runtime.fail();
            })();
            if ($1.tag === "Just") { return Data$dLazy.force($1._1); }
          }
        }
        return Partial._crashWith("Unbound local at level " + Data$dShow.showIntImpl(v1._2));
      }
      if (v1.tag === "App") { return evalApp(v)(dictEval.eval(v)(v1._1))(Data$dFunctor.arrayMap(dictEval.eval(v))(v1._2)); }
      if (v1.tag === "UncurriedApp") { return evalUncurriedApp(v)(dictEval.eval(v)(v1._1))(Data$dFunctor.arrayMap(dictEval.eval(v))(v1._2)); }
      if (v1.tag === "UncurriedAbs") {
        const $0 = v1._2;
        const loop = env$p => v2 => {
          if (v2.tag === "Nil") { return $MkFn("MkFnApplied", dictEval.eval(env$p)($0)); }
          if (v2.tag === "Cons") {
            const $1 = v2._2;
            return $MkFn("MkFnNext", v2._1, nextArg => loop({...env$p, locals: Data$dArray.snoc(env$p.locals)($LocalBinding("One", nextArg))})($1));
          }
          $runtime.fail();
        };
        return $BackendSemantics(
          "SemMkFn",
          loop(v)((() => {
            const $1 = Data$dFunctor.arrayMap(Data$dTuple.fst)(v1._1);
            const len = $1.length;
            const go = go$a0$copy => go$a1$copy => {
              let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
              while (go$c) {
                const source = go$a0, memo = go$a1;
                if (source < len) {
                  go$a0 = source + 1 | 0;
                  go$a1 = Data$dList$dTypes.$List("Cons", $1[source], memo);
                  continue;
                }
                const go$1 = go$1$a0$copy => go$1$a1$copy => {
                  let go$1$a0 = go$1$a0$copy, go$1$a1 = go$1$a1$copy, go$1$c = true, go$1$r;
                  while (go$1$c) {
                    const b = go$1$a0, v$1 = go$1$a1;
                    if (v$1.tag === "Nil") {
                      go$1$c = false;
                      go$1$r = b;
                      continue;
                    }
                    if (v$1.tag === "Cons") {
                      go$1$a0 = Data$dList$dTypes.$List("Cons", v$1._1, b);
                      go$1$a1 = v$1._2;
                      continue;
                    }
                    $runtime.fail();
                  }
                  return go$1$r;
                };
                go$c = false;
                go$r = go$1(Data$dList$dTypes.Nil)(memo);
              }
              return go$r;
            };
            return go(0)(Data$dList$dTypes.Nil);
          })())
        );
      }
      if (v1.tag === "UncurriedEffectApp") { return evalUncurriedEffectApp(v)(dictEval.eval(v)(v1._1))(Data$dFunctor.arrayMap(dictEval.eval(v))(v1._2)); }
      if (v1.tag === "UncurriedEffectAbs") {
        const $0 = v1._2;
        const loop = env$p => v2 => {
          if (v2.tag === "Nil") { return $MkFn("MkFnApplied", dictEval.eval(env$p)($0)); }
          if (v2.tag === "Cons") {
            const $1 = v2._2;
            return $MkFn("MkFnNext", v2._1, nextArg => loop({...env$p, locals: Data$dArray.snoc(env$p.locals)($LocalBinding("One", nextArg))})($1));
          }
          $runtime.fail();
        };
        return $BackendSemantics(
          "SemMkEffectFn",
          loop(v)((() => {
            const $1 = Data$dFunctor.arrayMap(Data$dTuple.fst)(v1._1);
            const len = $1.length;
            const go = go$a0$copy => go$a1$copy => {
              let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
              while (go$c) {
                const source = go$a0, memo = go$a1;
                if (source < len) {
                  go$a0 = source + 1 | 0;
                  go$a1 = Data$dList$dTypes.$List("Cons", $1[source], memo);
                  continue;
                }
                const go$1 = go$1$a0$copy => go$1$a1$copy => {
                  let go$1$a0 = go$1$a0$copy, go$1$a1 = go$1$a1$copy, go$1$c = true, go$1$r;
                  while (go$1$c) {
                    const b = go$1$a0, v$1 = go$1$a1;
                    if (v$1.tag === "Nil") {
                      go$1$c = false;
                      go$1$r = b;
                      continue;
                    }
                    if (v$1.tag === "Cons") {
                      go$1$a0 = Data$dList$dTypes.$List("Cons", v$1._1, b);
                      go$1$a1 = v$1._2;
                      continue;
                    }
                    $runtime.fail();
                  }
                  return go$1$r;
                };
                go$c = false;
                go$r = go$1(Data$dList$dTypes.Nil)(memo);
              }
              return go$r;
            };
            return go(0)(Data$dList$dTypes.Nil);
          })())
        );
      }
      if (v1.tag === "Abs") {
        const $0 = v1._2;
        const $1 = v1._1;
        const go = go$a0$copy => go$a1$copy => {
          let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
          while (go$c) {
            const ix = go$a0, acc = go$a1;
            if (ix < 0) {
              go$c = false;
              go$r = acc;
              continue;
            }
            go$a0 = ix - 1 | 0;
            go$a1 = (() => {
              const $2 = $1[ix]._1;
              return env$p => $BackendSemantics("SemLam", $2, x => acc({...env$p, locals: Data$dArray.snoc(env$p.locals)($LocalBinding("One", x))}));
            })();
          }
          return go$r;
        };
        return go($1.length - 2 | 0)((() => {
          const $2 = $1.length - 1 | 0;
          const $3 = (() => {
            if ($2 >= 0 && $2 < $1.length) { return $1[$2]._1; }
            $runtime.fail();
          })();
          return env$p => $BackendSemantics("SemLam", $3, x => dictEval.eval({...env$p, locals: Data$dArray.snoc(env$p.locals)($LocalBinding("One", x))})($0));
        })())(v);
      }
      if (v1.tag === "Let") {
        const $0 = v1._4;
        return makeLet(v1._1)(dictEval.eval(v)(v1._3))(x => dictEval.eval({...v, locals: Data$dArray.snoc(v.locals)($LocalBinding("One", x))})($0));
      }
      if (v1.tag === "LetRec") {
        const bindGroup = sem => x => dictEval.eval({...v, locals: Data$dArray.snoc(v.locals)($LocalBinding("Group", x))})(sem);
        return $BackendSemantics("SemLetRec", Data$dFunctor.arrayMap(m => Data$dTuple.$Tuple(m._1, bindGroup(m._2)))(v1._2), bindGroup(v1._3));
      }
      if (v1.tag === "EffectBind") {
        const $0 = v1._4;
        return makeEffectBind(v1._1)(dictEval.eval(v)(v1._3))(x => dictEval.eval({...v, locals: Data$dArray.snoc(v.locals)($LocalBinding("One", x))})($0));
      }
      if (v1.tag === "EffectPure") {
        const $0 = dictEval.eval(v)(v1._1);
        if ($0.tag === "NeutFail") { return $BackendSemantics("NeutFail", $0._1); }
        return $BackendSemantics("SemEffectPure", $0);
      }
      if (v1.tag === "EffectDefer") {
        const $0 = dictEval.eval(v)(v1._1);
        if ($0.tag === "NeutFail") { return $BackendSemantics("NeutFail", $0._1); }
        return $BackendSemantics("SemEffectDefer", $0);
      }
      if (v1.tag === "Accessor") { return evalAccessor(v)(dictEval.eval(v)(v1._1))(v1._2); }
      if (v1.tag === "Update") {
        return evalUpdate(dictEval.eval(v)(v1._1))(Data$dFunctor.arrayMap((() => {
          const $0 = dictEval.eval(v);
          return m => PureScript$dBackend$dOptimizer$dCoreFn.$Prop(m._1, $0(m._2));
        })())(v1._2));
      }
      if (v1.tag === "Branch") {
        const $0 = v1._2;
        return evalBranches(v)(Data$dFunctor.arrayMap(evalPair(dictEval)(v))(v1._1))(Data$dLazy.defer(v2 => dictEval.eval(v)($0)));
      }
      if (v1.tag === "PrimOp") {
        return evalPrimOp(v)((() => {
          const $0 = dictEval.eval(v);
          if (v1._1.tag === "Op1") { return PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op1", v1._1._1, $0(v1._1._2)); }
          if (v1._1.tag === "Op2") { return PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op2", v1._1._1, $0(v1._1._2), $0(v1._1._3)); }
          $runtime.fail();
        })());
      }
      if (v1.tag === "PrimEffect") {
        return guardFailOver2(identity)((() => {
          const $0 = dictEval.eval(v);
          if (v1._1.tag === "EffectRefNew") { return PureScript$dBackend$dOptimizer$dSyntax.$BackendEffect("EffectRefNew", $0(v1._1._1)); }
          if (v1._1.tag === "EffectRefRead") { return PureScript$dBackend$dOptimizer$dSyntax.$BackendEffect("EffectRefRead", $0(v1._1._1)); }
          if (v1._1.tag === "EffectRefWrite") { return PureScript$dBackend$dOptimizer$dSyntax.$BackendEffect("EffectRefWrite", $0(v1._1._1), $0(v1._1._2)); }
          $runtime.fail();
        })())(NeutPrimEffect);
      }
      if (v1.tag === "PrimUndefined") { return NeutPrimUndefined; }
      if (v1.tag === "Lit") { return guardFailOver3(identity)(PureScript$dBackend$dOptimizer$dCoreFn.functorLiteral.map(dictEval.eval(v))(v1._1))(NeutLit); }
      if (v1.tag === "Fail") { return $BackendSemantics("NeutFail", v1._1); }
      if (v1.tag === "CtorDef") {
        return $BackendSemantics("NeutCtorDef", PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", v.currentModule), v1._3), v1._1, v1._2, v1._3, v1._4);
      }
      if (v1.tag === "CtorSaturated") {
        return guardFailOver1(Data$dTuple.snd)(Data$dFunctor.arrayMap((() => {
          const $0 = dictEval.eval(v);
          return m => Data$dTuple.$Tuple(m._1, $0(m._2));
        })())(v1._5))(NeutData(v1._1)(v1._2)(v1._3)(v1._4));
      }
      $runtime.fail();
    }
  }
);
const evalBackendExpr$lazy = /* #__PURE__ */ $runtime.binding(() => (
  {
    eval: (() => {
      const go = go$a0$copy => go$a1$copy => {
        let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
        while (go$c) {
          const env = go$a0, v = go$a1;
          if (v.tag === "ExprRewrite") {
            if (v._2.tag === "RewriteInline") {
              go$a0 = {...env, locals: Data$dArray.snoc(env.locals)($LocalBinding("One", evalBackendExpr$lazy().eval(env)(v._2._3)))};
              go$a1 = v._2._4;
              continue;
            }
            if (v._2.tag === "RewriteUncurry") {
              const $0 = v._2._3;
              const $1 = v._2._5;
              go$c = false;
              go$r = $BackendSemantics(
                "SemLet",
                v._2._1,
                mkFnFromArgs(evalBackendExpr$lazy())(env)($0)(v._2._4),
                newFn => evalBackendExpr$lazy().eval({...env, locals: Data$dArray.snoc(env.locals)($LocalBinding("One", mkUncurriedAppRewrite(env)(newFn)($0.length)))})($1)
              );
              continue;
            }
            if (v._2.tag === "RewriteStop") {
              go$c = false;
              go$r = $BackendSemantics("NeutStop", v._2._1);
              continue;
            }
            if (v._2.tag === "RewriteUnpackOp") {
              if (v._2._3.tag === "UnpackRecord") {
                go$c = false;
                go$r = Data$dFoldable.foldrArray(v1 => {
                  const $0 = v1._2;
                  const $1 = v1._1;
                  return next => props$p => makeLet(Data$dMaybe.Nothing)(evalBackendExpr$lazy().eval(env)($0))(val => next(Data$dArray.snoc(props$p)(PureScript$dBackend$dOptimizer$dCoreFn.$Prop(
                    $1,
                    val
                  ))));
                })((() => {
                  const $0 = v._2._4;
                  return x => evalBackendExpr$lazy().eval({
                    ...env,
                    locals: Data$dArray.snoc(env.locals)($LocalBinding("One", $BackendSemantics("NeutLit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitRecord", x))))
                  })($0);
                })())(v._2._3._1)([]);
                continue;
              }
              if (v._2._3.tag === "UnpackUpdate") {
                const $0 = v._2._3._2;
                go$c = false;
                go$r = makeLet(Data$dMaybe.Nothing)(evalBackendExpr$lazy().eval(env)(v._2._3._1))(hd$p => Data$dFoldable.foldrArray(v1 => {
                  const $1 = v1._2;
                  const $2 = v1._1;
                  return next => props$p => makeLet(Data$dMaybe.Nothing)(evalBackendExpr$lazy().eval(env)($1))(val => next(Data$dArray.snoc(props$p)(PureScript$dBackend$dOptimizer$dCoreFn.$Prop(
                    $2,
                    val
                  ))));
                })((() => {
                  const $1 = v._2._4;
                  const $2 = NeutUpdate(hd$p);
                  return x => evalBackendExpr$lazy().eval({...env, locals: Data$dArray.snoc(env.locals)($LocalBinding("One", $2(x)))})($1);
                })())($0)([]));
                continue;
              }
              if (v._2._3.tag === "UnpackArray") {
                go$c = false;
                go$r = Data$dFoldable.foldrArray(expr => next => exprs$p => makeLet(Data$dMaybe.Nothing)(evalBackendExpr$lazy().eval(env)(expr))(val => next(Data$dArray.snoc(exprs$p)(val))))((() => {
                  const $0 = v._2._4;
                  return x => evalBackendExpr$lazy().eval({
                    ...env,
                    locals: Data$dArray.snoc(env.locals)($LocalBinding("One", $BackendSemantics("NeutLit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitArray", x))))
                  })($0);
                })())(v._2._3._1)([]);
                continue;
              }
              if (v._2._3.tag === "UnpackData") {
                go$c = false;
                go$r = Data$dFoldable.foldrArray(v1 => {
                  const $0 = v1._2;
                  const $1 = v1._1;
                  return next => props$p => makeLet(Data$dMaybe.Nothing)(evalBackendExpr$lazy().eval(env)($0))(val => next(Data$dArray.snoc(props$p)(Data$dTuple.$Tuple($1, val))));
                })((() => {
                  const $0 = v._2._4;
                  const $1 = NeutData(v._2._3._1)(v._2._3._2)(v._2._3._3)(v._2._3._4);
                  return x => evalBackendExpr$lazy().eval({...env, locals: Data$dArray.snoc(env.locals)($LocalBinding("One", $1(x)))})($0);
                })())(v._2._3._5)([]);
                continue;
              }
              $runtime.fail();
            }
            if (v._2.tag === "RewriteDistBranchesLet") {
              const $0 = v._2._4;
              go$c = false;
              go$r = rewriteBranches((() => {
                const $1 = v._2._5;
                return x => evalBackendExpr$lazy().eval({...env, locals: Data$dArray.snoc(env.locals)($LocalBinding("One", x))})($1);
              })())(evalBranches(env)(Data$dFunctor.arrayMap(evalPair(evalBackendExpr$lazy())(env))(v._2._3))(Data$dLazy.defer(v1 => evalBackendExpr$lazy().eval(env)($0))));
              continue;
            }
            if (v._2.tag === "RewriteDistBranchesOp") {
              const $0 = v._2._2;
              go$c = false;
              go$r = rewriteBranches((() => {
                if (v._2._3.tag === "DistApp") {
                  const $1 = Data$dFunctor.arrayMap(evalBackendExpr$lazy().eval(env))(v._2._3._1);
                  return a => evalApp(env)(a)($1);
                }
                if (v._2._3.tag === "DistUncurriedApp") {
                  const $1 = Data$dFunctor.arrayMap(evalBackendExpr$lazy().eval(env))(v._2._3._1);
                  return a => evalUncurriedApp(env)(a)($1);
                }
                if (v._2._3.tag === "DistAccessor") {
                  const $1 = v._2._3._1;
                  return a => evalAccessor(env)(a)($1);
                }
                if (v._2._3.tag === "DistPrimOp1") {
                  const $1 = PureScript$dBackend$dOptimizer$dSyntax.Op1(v._2._3._1);
                  return x => evalPrimOp(env)($1(x));
                }
                if (v._2._3.tag === "DistPrimOp2L") {
                  const $1 = PureScript$dBackend$dOptimizer$dSyntax.Op2(v._2._3._1);
                  const $2 = evalBackendExpr$lazy().eval(env)(v._2._3._2);
                  return x => evalPrimOp(env)($1(x)($2));
                }
                if (v._2._3.tag === "DistPrimOp2R") {
                  const $1 = PureScript$dBackend$dOptimizer$dSyntax.Op2(v._2._3._2)(evalBackendExpr$lazy().eval(env)(v._2._3._1));
                  return x => evalPrimOp(env)($1(x));
                }
                $runtime.fail();
              })())(evalBranches(env)(Data$dFunctor.arrayMap(evalPair(evalBackendExpr$lazy())(env))(v._2._1))(Data$dLazy.defer(v1 => evalBackendExpr$lazy().eval(env)($0))));
              continue;
            }
            $runtime.fail();
          }
          if (v.tag === "ExprSyntax") {
            go$c = false;
            go$r = evalBackendSyntax(evalBackendExpr$lazy()).eval(env)(v._2);
            continue;
          }
          $runtime.fail();
        }
        return go$r;
      };
      return go;
    })()
  }
));
const evalBackendExpr = /* #__PURE__ */ evalBackendExpr$lazy();
const optimize = traceSteps => ctx => env => v => initN => originalExpr => {
  const $0 = v._2;
  const $1 = v._1;
  const go = go$a0$copy => go$a1$copy => go$a2$copy => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$a2 = go$a2$copy, go$c = true, go$r;
    while (go$c) {
      const steps = go$a0, n = go$a1, expr1 = go$a2;
      const v1 = (() => {
        if (n === 0) {
          return Partial._crashWith((() => {
            if ($1.tag === "Nothing") { return "" + $0 + ": Possible infinite optimization loop."; }
            if ($1.tag === "Just") { return $1._1 + "." + $0 + ": Possible infinite optimization loop."; }
            $runtime.fail();
          })());
        }
        const expr2 = quote(ctx)(evalBackendExpr.eval(env)(expr1));
        return Data$dTuple.$Tuple(
          (() => {
            if (expr2.tag === "ExprSyntax") { return expr2._1.rewrite; }
            if (expr2.tag === "ExprRewrite") { return expr2._1.rewrite; }
            $runtime.fail();
          })(),
          expr2
        );
      })();
      const newSteps = traceSteps ? Data$dList$dTypes.$List("Cons", v1._2, steps) : steps;
      if (v1._1) {
        go$a0 = newSteps;
        go$a1 = n - 1 | 0;
        go$a2 = v1._2;
        continue;
      }
      go$c = false;
      go$r = Data$dTuple.$Tuple(Data$dArray.reverse(toUnfoldable1(newSteps)), v1._2);
    }
    return go$r;
  };
  return go(traceSteps ? Data$dList$dTypes.$List("Cons", originalExpr, Data$dList$dTypes.Nil) : Data$dList$dTypes.Nil)(initN)(originalExpr);
};
const evalNeutralExpr = {eval: env => v => evalBackendSyntax(evalNeutralExpr).eval(env)(v)};
const eval3 = /* #__PURE__ */ (() => evalBackendSyntax(evalNeutralExpr).eval)();
const analysisFromDirective = v => v1 => {
  if (v1.tag === "InlineAlways") { return PureScript$dBackend$dOptimizer$dAnalysis.monoidBackendAnalysis.mempty; }
  if (v1.tag === "InlineNever") { return {...v, complexity: PureScript$dBackend$dOptimizer$dAnalysis.NonTrivial, size: 2147483647}; }
  if (v1.tag === "InlineArity") { return {...v, args: v1._1 < 1 ? [] : Data$dArray.sliceImpl(0, v1._1, v.args)}; }
  if (v1.tag === "InlineDefault") { return v; }
  $runtime.fail();
};
const addStop = v => ref => acc => (
  {
    ...v,
    directives: alter(v2 => {
      if (v2.tag === "Just") { return Data$dMaybe.$Maybe("Just", Data$dMap$dInternal.insert(ordInlineAccessor)(acc)(InlineNever)(v2._1)); }
      return Data$dMaybe.$Maybe("Just", Data$dMap$dInternal.$$$Map("Node", 1, 1, acc, InlineNever, Data$dMap$dInternal.Leaf, Data$dMap$dInternal.Leaf));
    })(ref)(v.directives)
  }
);
const envForGroup = env => ref => acc => group => {
  if (group.length === 0) { return env; }
  return addStop(env)(ref)(acc);
};
const evalExternFromImpl = v => qual => v1 => spine => {
  if (spine.length === 0) {
    if (v1._2.tag === "ExternExpr") {
      const $0 = v1._2._2;
      const $1 = v1._2._1;
      const $2 = lookup3($EvalRef("EvalExtern", qual))(v.directives);
      const $3 = lookup2(InlineRef);
      const v2 = (() => {
        if ($2.tag === "Just") { return $3($2._1); }
        if ($2.tag === "Nothing") { return Data$dMaybe.Nothing; }
        $runtime.fail();
      })();
      const $4 = () => {
        if ($0.tag === "Lit" && shouldInlineExternLiteral($0._1)) {
          return Data$dMaybe.$Maybe("Just", evalBackendSyntax(evalNeutralExpr).eval($1.length === 0 ? v : addStop(v)($EvalRef("EvalExtern", qual))(InlineRef))($0));
        }
        if ((v1._1.complexity === "Trivial" || v1._1.complexity === "Deref") && v1._1.size < 16) {
          return Data$dMaybe.$Maybe("Just", evalBackendSyntax(evalNeutralExpr).eval($1.length === 0 ? v : addStop(v)($EvalRef("EvalExtern", qual))(InlineRef))($0));
        }
        return Data$dMaybe.Nothing;
      };
      if (v2.tag === "Just") {
        if (v2._1.tag === "InlineNever") { return Data$dMaybe.$Maybe("Just", $BackendSemantics("NeutStop", qual)); }
        if (v2._1.tag === "InlineAlways") {
          return Data$dMaybe.$Maybe("Just", evalBackendSyntax(evalNeutralExpr).eval($1.length === 0 ? v : addStop(v)($EvalRef("EvalExtern", qual))(InlineRef))($0));
        }
        if (v2._1.tag === "InlineArity") { return Data$dMaybe.Nothing; }
      }
      return $4();
    }
    if (v1._2.tag === "ExternCtor" && v1._2._5.length === 0) { return Data$dMaybe.$Maybe("Just", $BackendSemantics("NeutData", qual, v1._2._2, v1._2._3, v1._2._4, [])); }
    return Data$dMaybe.Nothing;
  }
  if (spine.length === 1) {
    if (spine[0].tag === "ExternAccessor") {
      if (spine[0]._1.tag === "GetProp") {
        if (v1._2.tag === "ExternExpr") {
          const $0 = lookup3($EvalRef("EvalExtern", qual))(v.directives);
          const $1 = lookup2($InlineAccessor("InlineProp", spine[0]._1._1));
          const v2 = (() => {
            if ($0.tag === "Just") { return $1($0._1); }
            if ($0.tag === "Nothing") { return Data$dMaybe.Nothing; }
            $runtime.fail();
          })();
          if (v2.tag === "Just") {
            if (v2._1.tag === "InlineNever") { return Data$dMaybe.$Maybe("Just", neutralSpine($BackendSemantics("NeutStop", qual))(spine)); }
            if (v2._1.tag === "InlineAlways") {
              return Data$dMaybe.$Maybe(
                "Just",
                evalSpine(v)(evalBackendSyntax(evalNeutralExpr).eval(v1._2._1.length === 0
                  ? v
                  : addStop(v)($EvalRef("EvalExtern", qual))($InlineAccessor("InlineProp", spine[0]._1._1)))(v1._2._2))(spine)
              );
            }
          }
          return Data$dMaybe.Nothing;
        }
        if (v1._2.tag === "ExternDict") {
          const $0 = Data$dArray.findMapImpl(
            Data$dMaybe.Nothing,
            Data$dMaybe.isJust,
            v$1 => {
              if (spine[0]._1._1 === v$1._1) { return Data$dMaybe.$Maybe("Just", v$1._2); }
              return Data$dMaybe.Nothing;
            },
            v1._2._2
          );
          if ($0.tag === "Just") {
            const $1 = $0._1._2;
            const $2 = lookup3($EvalRef("EvalExtern", qual))(v.directives);
            const $3 = lookup2($InlineAccessor("InlineProp", spine[0]._1._1));
            const v3 = (() => {
              if ($2.tag === "Just") { return $3($2._1); }
              if ($2.tag === "Nothing") { return Data$dMaybe.Nothing; }
              $runtime.fail();
            })();
            const $4 = () => Data$dMaybe.$Maybe(
              "Just",
              evalBackendSyntax(evalNeutralExpr).eval(v1._2._1.length === 0 ? v : addStop(v)($EvalRef("EvalExtern", qual))($InlineAccessor("InlineProp", spine[0]._1._1)))($1)
            );
            if (v3.tag === "Just") {
              if (v3._1.tag === "InlineNever") { return Data$dMaybe.$Maybe("Just", neutralSpine($BackendSemantics("NeutStop", qual))(spine)); }
              if (v3._1.tag === "InlineAlways") {
                return Data$dMaybe.$Maybe(
                  "Just",
                  evalBackendSyntax(evalNeutralExpr).eval(v1._2._1.length === 0 ? v : addStop(v)($EvalRef("EvalExtern", qual))($InlineAccessor("InlineProp", spine[0]._1._1)))($1)
                );
              }
              if (v3._1.tag === "InlineArity") { return Data$dMaybe.Nothing; }
            }
            if (($0._1._1.complexity === "Trivial" || $0._1._1.complexity === "Deref") && $0._1._1.size < 16) { return $4(); }
          }
        }
      }
      return Data$dMaybe.Nothing;
    }
    if (spine[0].tag === "ExternApp") {
      if (v1._2.tag === "ExternExpr") {
        const $0 = v1._2._2;
        const $1 = v1._2._1;
        const $2 = lookup3($EvalRef("EvalExtern", qual))(v.directives);
        const $3 = lookup2(InlineRef);
        const v2 = (() => {
          if ($2.tag === "Just") { return $3($2._1); }
          if ($2.tag === "Nothing") { return Data$dMaybe.Nothing; }
          $runtime.fail();
        })();
        const $4 = () => Data$dMaybe.$Maybe(
          "Just",
          evalApp(v)(evalBackendSyntax(evalNeutralExpr).eval($1.length === 0 ? v : addStop(v)($EvalRef("EvalExtern", qual))(InlineRef))($0))(spine[0]._1)
        );
        if (v2.tag === "Just") {
          if (v2._1.tag === "InlineNever") { return Data$dMaybe.$Maybe("Just", neutralSpine($BackendSemantics("NeutStop", qual))(spine)); }
          if (v2._1.tag === "InlineAlways") {
            return Data$dMaybe.$Maybe(
              "Just",
              evalApp(v)(evalBackendSyntax(evalNeutralExpr).eval($1.length === 0 ? v : addStop(v)($EvalRef("EvalExtern", qual))(InlineRef))($0))(spine[0]._1)
            );
          }
          if (v2._1.tag === "InlineArity") {
            if (spine[0]._1.length >= v2._1._1) {
              return Data$dMaybe.$Maybe(
                "Just",
                evalApp(v)(evalBackendSyntax(evalNeutralExpr).eval($1.length === 0 ? v : addStop(v)($EvalRef("EvalExtern", qual))(InlineRef))($0))(spine[0]._1)
              );
            }
            return Data$dMaybe.Nothing;
          }
        }
        if (shouldInlineExternApp(qual)(v1._1)($0)(spine[0]._1)) { return $4(); }
        return Data$dMaybe.Nothing;
      }
      if (v1._2.tag === "ExternCtor" && v1._2._5.length === spine[0]._1.length) {
        return Data$dMaybe.$Maybe("Just", $BackendSemantics("NeutData", qual, v1._2._2, v1._2._3, v1._2._4, Data$dArray.zipWithImpl(Data$dTuple.Tuple, v1._2._5, spine[0]._1)));
      }
    }
    return Data$dMaybe.Nothing;
  }
  if (spine.length === 2) {
    if (spine[0].tag === "ExternAccessor") {
      if (spine[0]._1.tag === "GetProp" && spine[1].tag === "ExternApp") {
        if (v1._2.tag === "ExternExpr") {
          const $0 = lookup3($EvalRef("EvalExtern", qual))(v.directives);
          const $1 = lookup2($InlineAccessor("InlineProp", spine[0]._1._1));
          const v2 = (() => {
            if ($0.tag === "Just") { return $1($0._1); }
            if ($0.tag === "Nothing") { return Data$dMaybe.Nothing; }
            $runtime.fail();
          })();
          if (v2.tag === "Just") {
            if (v2._1.tag === "InlineNever") { return Data$dMaybe.$Maybe("Just", neutralSpine($BackendSemantics("NeutStop", qual))(spine)); }
            if (v2._1.tag === "InlineAlways") {
              return Data$dMaybe.$Maybe(
                "Just",
                evalSpine(v)(evalBackendSyntax(evalNeutralExpr).eval(v1._2._1.length === 0
                  ? v
                  : addStop(v)($EvalRef("EvalExtern", qual))($InlineAccessor("InlineProp", spine[0]._1._1)))(v1._2._2))(spine)
              );
            }
            if (v2._1.tag === "InlineArity" && spine[1]._1.length >= v2._1._1) {
              return Data$dMaybe.$Maybe(
                "Just",
                evalSpine(v)(evalBackendSyntax(evalNeutralExpr).eval(v1._2._1.length === 0
                  ? v
                  : addStop(v)($EvalRef("EvalExtern", qual))($InlineAccessor("InlineProp", spine[0]._1._1)))(v1._2._2))(spine)
              );
            }
          }
          return Data$dMaybe.Nothing;
        }
        if (v1._2.tag === "ExternDict") {
          const $0 = Data$dArray.findMapImpl(
            Data$dMaybe.Nothing,
            Data$dMaybe.isJust,
            v$1 => {
              if (spine[0]._1._1 === v$1._1) { return Data$dMaybe.$Maybe("Just", v$1._2); }
              return Data$dMaybe.Nothing;
            },
            v1._2._2
          );
          if ($0.tag === "Just") {
            const $1 = $0._1._2;
            const $2 = lookup3($EvalRef("EvalExtern", qual))(v.directives);
            const $3 = lookup2($InlineAccessor("InlineProp", spine[0]._1._1));
            const v3 = (() => {
              if ($2.tag === "Just") { return $3($2._1); }
              if ($2.tag === "Nothing") { return Data$dMaybe.Nothing; }
              $runtime.fail();
            })();
            const $4 = () => Data$dMaybe.$Maybe(
              "Just",
              evalApp(v)(evalBackendSyntax(evalNeutralExpr).eval(v1._2._1.length === 0 ? v : addStop(v)($EvalRef("EvalExtern", qual))($InlineAccessor("InlineProp", spine[0]._1._1)))($1))(spine[1]._1)
            );
            if (v3.tag === "Just") {
              if (v3._1.tag === "InlineNever") { return Data$dMaybe.$Maybe("Just", neutralSpine($BackendSemantics("NeutStop", qual))(spine)); }
              if (v3._1.tag === "InlineAlways") {
                return Data$dMaybe.$Maybe(
                  "Just",
                  evalApp(v)(evalBackendSyntax(evalNeutralExpr).eval(v1._2._1.length === 0
                    ? v
                    : addStop(v)($EvalRef("EvalExtern", qual))($InlineAccessor("InlineProp", spine[0]._1._1)))($1))(spine[1]._1)
                );
              }
              if (v3._1.tag === "InlineArity") {
                if (spine[1]._1.length >= v3._1._1) {
                  return Data$dMaybe.$Maybe(
                    "Just",
                    evalApp(v)(evalBackendSyntax(evalNeutralExpr).eval(v1._2._1.length === 0
                      ? v
                      : addStop(v)($EvalRef("EvalExtern", qual))($InlineAccessor("InlineProp", spine[0]._1._1)))($1))(spine[1]._1)
                  );
                }
                return Data$dMaybe.Nothing;
              }
            }
            if (shouldInlineExternApp(qual)($0._1._1)($1)(spine[1]._1)) { return $4(); }
          }
        }
      }
      return Data$dMaybe.Nothing;
    }
    if (spine[0].tag === "ExternApp" && spine[1].tag === "ExternAccessor" && spine[1]._1.tag === "GetProp" && v1._2.tag === "ExternExpr") {
      const $0 = lookup3($EvalRef("EvalExtern", qual))(v.directives);
      const $1 = lookup2($InlineAccessor("InlineSpineProp", spine[1]._1._1));
      const v2 = (() => {
        if ($0.tag === "Just") { return $1($0._1); }
        if ($0.tag === "Nothing") { return Data$dMaybe.Nothing; }
        $runtime.fail();
      })();
      if (v2.tag === "Just") {
        if (v2._1.tag === "InlineNever") { return Data$dMaybe.$Maybe("Just", neutralSpine($BackendSemantics("NeutStop", qual))(spine)); }
        if (v2._1.tag === "InlineAlways") {
          return Data$dMaybe.$Maybe(
            "Just",
            evalSpine(v)(evalBackendSyntax(evalNeutralExpr).eval(v1._2._1.length === 0
              ? v
              : addStop(v)($EvalRef("EvalExtern", qual))($InlineAccessor("InlineSpineProp", spine[1]._1._1)))(v1._2._2))(spine)
          );
        }
      }
    }
    return Data$dMaybe.Nothing;
  }
  if (
    spine.length === 3 && spine[0].tag === "ExternApp" && spine[1].tag === "ExternAccessor" && spine[1]._1.tag === "GetProp" && spine[2].tag === "ExternApp" && v1._2.tag === "ExternExpr"
  ) {
    const $0 = lookup3($EvalRef("EvalExtern", qual))(v.directives);
    const $1 = lookup2($InlineAccessor("InlineSpineProp", spine[1]._1._1));
    const v2 = (() => {
      if ($0.tag === "Just") { return $1($0._1); }
      if ($0.tag === "Nothing") { return Data$dMaybe.Nothing; }
      $runtime.fail();
    })();
    if (v2.tag === "Just") {
      if (v2._1.tag === "InlineNever") { return Data$dMaybe.$Maybe("Just", neutralSpine($BackendSemantics("NeutStop", qual))(spine)); }
      if (v2._1.tag === "InlineAlways") {
        return Data$dMaybe.$Maybe(
          "Just",
          evalSpine(v)(evalBackendSyntax(evalNeutralExpr).eval(v1._2._1.length === 0
            ? v
            : addStop(v)($EvalRef("EvalExtern", qual))($InlineAccessor("InlineSpineProp", spine[1]._1._1)))(v1._2._2))(spine)
        );
      }
      if (v2._1.tag === "InlineArity" && spine[2]._1.length >= v2._1._1) {
        return Data$dMaybe.$Maybe(
          "Just",
          evalSpine(v)(evalBackendSyntax(evalNeutralExpr).eval(v1._2._1.length === 0
            ? v
            : addStop(v)($EvalRef("EvalExtern", qual))($InlineAccessor("InlineSpineProp", spine[1]._1._1)))(v1._2._2))(spine)
        );
      }
    }
  }
  return Data$dMaybe.Nothing;
};
const evalExternRefFromImpl = env => qual => v => {
  if (v._2.tag === "ExternExpr") {
    if (v._2._2.tag === "Var" || v._2._2.tag === "Lit" || v._2._2.tag === "CtorSaturated" || v._2._2.tag === "Accessor" || v._2._2.tag === "Update" || v._2._2.tag === "PrimOp") {
      return eval3(v._2._1.length === 0 ? env : addStop(env)($EvalRef("EvalExtern", qual))(InlineRef))(v._2._2);
    }
    return $BackendSemantics("NeutVar", qual);
  }
  if (v._2.tag === "ExternDict") {
    const $0 = v._2._1;
    return $BackendSemantics(
      "NeutLit",
      PureScript$dBackend$dOptimizer$dCoreFn.$Literal(
        "LitRecord",
        Data$dFunctor.arrayMap(v1 => PureScript$dBackend$dOptimizer$dCoreFn.$Prop(
          v1._1,
          eval3($0.length === 0 ? env : addStop(env)($EvalRef("EvalExtern", qual))($InlineAccessor("InlineProp", v1._1)))(v1._2._2)
        ))(v._2._2)
      )
    );
  }
  return $BackendSemantics("NeutVar", qual);
};
export {
  $BackendExpr,
  $BackendRewrite,
  $BackendSemantics,
  $DistOp,
  $EvalRef,
  $ExternImpl,
  $ExternSpine,
  $InlineAccessor,
  $InlineDirective,
  $LocalBinding,
  $MkFn,
  $SemConditional,
  $UnpackOp,
  Ctx,
  DistAccessor,
  DistApp,
  DistPrimOp1,
  DistPrimOp2L,
  DistPrimOp2R,
  DistUncurriedApp,
  Env,
  EvalExtern,
  EvalLocal,
  ExprRewrite,
  ExprSyntax,
  ExternAccessor,
  ExternApp,
  ExternCtor,
  ExternDict,
  ExternExpr,
  ExternPrimOp,
  ExternUncurriedApp,
  Group,
  InlineAlways,
  InlineArity,
  InlineDefault,
  InlineNever,
  InlineProp,
  InlineRef,
  InlineSpineProp,
  MkFnApplied,
  MkFnNext,
  NeutAccessor,
  NeutApp,
  NeutCtorDef,
  NeutData,
  NeutFail,
  NeutLit,
  NeutLocal,
  NeutPrimEffect,
  NeutPrimOp,
  NeutPrimUndefined,
  NeutStop,
  NeutUncurriedApp,
  NeutUncurriedEffectApp,
  NeutUpdate,
  NeutVar,
  NeutralExpr,
  One,
  RewriteDistBranchesLet,
  RewriteDistBranchesOp,
  RewriteInline,
  RewriteStop,
  RewriteUncurry,
  RewriteUnpackOp,
  SemAssocOp,
  SemBranch,
  SemConditional,
  SemEffectBind,
  SemEffectDefer,
  SemEffectPure,
  SemLam,
  SemLet,
  SemLetRec,
  SemMkEffectFn,
  SemMkFn,
  SemRef,
  UnpackArray,
  UnpackData,
  UnpackRecord,
  UnpackUpdate,
  addStop,
  alter,
  analysisFromDirective,
  and,
  bindLocal,
  build,
  buildBranchCond,
  buildDefault,
  buildStop,
  caseInt,
  caseNumber,
  caseString,
  compare1,
  compare2,
  deref,
  effectfully,
  envForGroup,
  eq10,
  eq16,
  eqBackendExpr,
  eqBackendRewrite,
  eqDistOp,
  eqEvalRef,
  eqInlineAccessor,
  eqUnpackOp,
  $$eval as eval,
  eval3,
  evalAccessor,
  evalApp,
  evalAssocOp,
  evalAssocOp$p,
  evalBackendExpr,
  evalBackendSyntax,
  evalBranches,
  evalEvalRef,
  evalExternFromImpl,
  evalExternRefFromImpl,
  evalMkFn,
  evalNeutralExpr,
  evalPair,
  evalPrimOp,
  evalPrimOpNot,
  evalPrimOpNumInt,
  evalPrimOpNumNumber,
  evalPrimOpOrd,
  evalPrimOpOrd1,
  evalPrimOpOrd2,
  evalPrimOpOrd3,
  evalPrimOpOrd4,
  evalPrimOpOrdNumber,
  evalRef,
  evalRefSpine,
  evalSpine,
  evalUncurriedApp,
  evalUncurriedBeta,
  evalUncurriedEffectApp,
  evalUpdate,
  floatLet,
  floatLetWith,
  foldBackendExpr,
  foldMap,
  foldMap1,
  freeze,
  fromFoldable,
  guardFail,
  guardFailOver,
  guardFailOver1,
  guardFailOver2,
  guardFailOver3,
  hasAnalysisBackendExpr,
  hasSyntaxBackendExpr,
  identity,
  insertDirective,
  isAbs,
  isAssocPrimOp,
  isKnownEffect,
  isRefExpr,
  isReference,
  isSimplePredicate,
  liftBoolean,
  liftInt,
  liftNumber,
  liftOp1,
  liftOp2,
  liftString,
  lookup,
  lookup1,
  lookup2,
  lookup3,
  lookupLocal,
  makeEffectBind,
  makeLet,
  mkFnFromArgs,
  mkUncurriedAppRewrite,
  neutralApp,
  neutralSpine,
  newtypeEnv_,
  newtypeNeutralExpr_,
  nextLevel,
  optimize,
  or,
  ordEvalRef,
  ordInlineAccessor,
  power,
  primOpOrdNot,
  purely,
  quote,
  rewriteBranches,
  rewriteInline,
  shouldDistributeBranchAccessor,
  shouldDistributeBranchApps,
  shouldDistributeBranchPrimOp1,
  shouldDistributeBranchPrimOp2L,
  shouldDistributeBranchPrimOp2R,
  shouldDistributeBranchUncurriedApps,
  shouldDistributeBranches,
  shouldEtaReduce,
  shouldInlineExternAccessor,
  shouldInlineExternApp,
  shouldInlineExternAppArg,
  shouldInlineExternLiteral,
  shouldInlineExternReference,
  shouldInlineLet,
  shouldUncurryAbs,
  shouldUnpackArray,
  shouldUnpackCtor,
  shouldUnpackRecord,
  shouldUnpackUpdate,
  simplifyCondBoolean,
  simplifyCondIsTag,
  simplifyCondLiftAnd,
  simplifyCondRedundantElse,
  snocApp,
  snocSpine,
  toUnfoldable,
  toUnfoldable1
};
