import * as $runtime from "../runtime.js";
import * as Data$dEq from "../Data.Eq/index.js";
import * as Data$dFoldable from "../Data.Foldable/index.js";
import * as Data$dFunctor from "../Data.Functor/index.js";
import * as Data$dOrd from "../Data.Ord/index.js";
import * as Data$dOrdering from "../Data.Ordering/index.js";
import * as Data$dTraversable from "../Data.Traversable/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
import * as PureScript$dBackend$dOptimizer$dCoreFn from "../PureScript.Backend.Optimizer.CoreFn/index.js";
const $BackendAccessor = (tag, _1, _2, _3, _4, _5, _6) => ({tag, _1, _2, _3, _4, _5, _6});
const $BackendEffect = (tag, _1, _2) => ({tag, _1, _2});
const $BackendOperator = (tag, _1, _2, _3) => ({tag, _1, _2, _3});
const $BackendOperator1 = (tag, _1) => ({tag, _1});
const $BackendOperator2 = (tag, _1) => ({tag, _1});
const $BackendOperatorNum = tag => tag;
const $BackendOperatorOrd = tag => tag;
const $BackendSyntax = (tag, _1, _2, _3, _4, _5) => ({tag, _1, _2, _3, _4, _5});
const $Pair = (_1, _2) => ({tag: "Pair", _1, _2});
const compare = /* #__PURE__ */ (() => PureScript$dBackend$dOptimizer$dCoreFn.ordQualified(Data$dOrd.ordString).compare)();
const eq7 = /* #__PURE__ */ Data$dEq.eqArrayImpl(Data$dEq.eqStringImpl);
const Pair = value0 => value1 => $Pair(value0, value1);
const Level = x => x;
const OpEq = /* #__PURE__ */ $BackendOperatorOrd("OpEq");
const OpNotEq = /* #__PURE__ */ $BackendOperatorOrd("OpNotEq");
const OpGt = /* #__PURE__ */ $BackendOperatorOrd("OpGt");
const OpGte = /* #__PURE__ */ $BackendOperatorOrd("OpGte");
const OpLt = /* #__PURE__ */ $BackendOperatorOrd("OpLt");
const OpLte = /* #__PURE__ */ $BackendOperatorOrd("OpLte");
const OpAdd = /* #__PURE__ */ $BackendOperatorNum("OpAdd");
const OpDivide = /* #__PURE__ */ $BackendOperatorNum("OpDivide");
const OpMultiply = /* #__PURE__ */ $BackendOperatorNum("OpMultiply");
const OpSubtract = /* #__PURE__ */ $BackendOperatorNum("OpSubtract");
const OpArrayIndex = /* #__PURE__ */ $BackendOperator2("OpArrayIndex");
const OpBooleanAnd = /* #__PURE__ */ $BackendOperator2("OpBooleanAnd");
const OpBooleanOr = /* #__PURE__ */ $BackendOperator2("OpBooleanOr");
const OpBooleanOrd = value0 => $BackendOperator2("OpBooleanOrd", value0);
const OpCharOrd = value0 => $BackendOperator2("OpCharOrd", value0);
const OpIntBitAnd = /* #__PURE__ */ $BackendOperator2("OpIntBitAnd");
const OpIntBitOr = /* #__PURE__ */ $BackendOperator2("OpIntBitOr");
const OpIntBitShiftLeft = /* #__PURE__ */ $BackendOperator2("OpIntBitShiftLeft");
const OpIntBitShiftRight = /* #__PURE__ */ $BackendOperator2("OpIntBitShiftRight");
const OpIntBitXor = /* #__PURE__ */ $BackendOperator2("OpIntBitXor");
const OpIntBitZeroFillShiftRight = /* #__PURE__ */ $BackendOperator2("OpIntBitZeroFillShiftRight");
const OpIntNum = value0 => $BackendOperator2("OpIntNum", value0);
const OpIntOrd = value0 => $BackendOperator2("OpIntOrd", value0);
const OpNumberNum = value0 => $BackendOperator2("OpNumberNum", value0);
const OpNumberOrd = value0 => $BackendOperator2("OpNumberOrd", value0);
const OpStringAppend = /* #__PURE__ */ $BackendOperator2("OpStringAppend");
const OpStringOrd = value0 => $BackendOperator2("OpStringOrd", value0);
const OpBooleanNot = /* #__PURE__ */ $BackendOperator1("OpBooleanNot");
const OpIntBitNot = /* #__PURE__ */ $BackendOperator1("OpIntBitNot");
const OpIntNegate = /* #__PURE__ */ $BackendOperator1("OpIntNegate");
const OpNumberNegate = /* #__PURE__ */ $BackendOperator1("OpNumberNegate");
const OpArrayLength = /* #__PURE__ */ $BackendOperator1("OpArrayLength");
const OpIsTag = value0 => $BackendOperator1("OpIsTag", value0);
const Op1 = value0 => value1 => $BackendOperator("Op1", value0, value1);
const Op2 = value0 => value1 => value2 => $BackendOperator("Op2", value0, value1, value2);
const EffectRefNew = value0 => $BackendEffect("EffectRefNew", value0);
const EffectRefRead = value0 => $BackendEffect("EffectRefRead", value0);
const EffectRefWrite = value0 => value1 => $BackendEffect("EffectRefWrite", value0, value1);
const GetProp = value0 => $BackendAccessor("GetProp", value0);
const GetIndex = value0 => $BackendAccessor("GetIndex", value0);
const GetCtorField = value0 => value1 => value2 => value3 => value4 => value5 => $BackendAccessor("GetCtorField", value0, value1, value2, value3, value4, value5);
const Var = value0 => $BackendSyntax("Var", value0);
const Local = value0 => value1 => $BackendSyntax("Local", value0, value1);
const Lit = value0 => $BackendSyntax("Lit", value0);
const App = value0 => value1 => $BackendSyntax("App", value0, value1);
const Abs = value0 => value1 => $BackendSyntax("Abs", value0, value1);
const UncurriedApp = value0 => value1 => $BackendSyntax("UncurriedApp", value0, value1);
const UncurriedAbs = value0 => value1 => $BackendSyntax("UncurriedAbs", value0, value1);
const UncurriedEffectApp = value0 => value1 => $BackendSyntax("UncurriedEffectApp", value0, value1);
const UncurriedEffectAbs = value0 => value1 => $BackendSyntax("UncurriedEffectAbs", value0, value1);
const Accessor = value0 => value1 => $BackendSyntax("Accessor", value0, value1);
const Update = value0 => value1 => $BackendSyntax("Update", value0, value1);
const CtorSaturated = value0 => value1 => value2 => value3 => value4 => $BackendSyntax("CtorSaturated", value0, value1, value2, value3, value4);
const CtorDef = value0 => value1 => value2 => value3 => $BackendSyntax("CtorDef", value0, value1, value2, value3);
const LetRec = value0 => value1 => value2 => $BackendSyntax("LetRec", value0, value1, value2);
const Let = value0 => value1 => value2 => value3 => $BackendSyntax("Let", value0, value1, value2, value3);
const EffectBind = value0 => value1 => value2 => value3 => $BackendSyntax("EffectBind", value0, value1, value2, value3);
const EffectPure = value0 => $BackendSyntax("EffectPure", value0);
const EffectDefer = value0 => $BackendSyntax("EffectDefer", value0);
const Branch = value0 => value1 => $BackendSyntax("Branch", value0, value1);
const PrimOp = value0 => $BackendSyntax("PrimOp", value0);
const PrimEffect = value0 => $BackendSyntax("PrimEffect", value0);
const PrimUndefined = /* #__PURE__ */ $BackendSyntax("PrimUndefined");
const Fail = value0 => $BackendSyntax("Fail", value0);
const ordLevel = Data$dOrd.ordInt;
const newtypeLevel_ = {Coercible0: () => {}};
const functorPair = {map: f => m => $Pair(f(m._1), f(m._2))};
const functorBackendOperator = {
  map: f => m => {
    if (m.tag === "Op1") { return $BackendOperator("Op1", m._1, f(m._2)); }
    if (m.tag === "Op2") { return $BackendOperator("Op2", m._1, f(m._2), f(m._3)); }
    $runtime.fail();
  }
};
const functorBackendEffect = {
  map: f => m => {
    if (m.tag === "EffectRefNew") { return $BackendEffect("EffectRefNew", f(m._1)); }
    if (m.tag === "EffectRefRead") { return $BackendEffect("EffectRefRead", f(m._1)); }
    if (m.tag === "EffectRefWrite") { return $BackendEffect("EffectRefWrite", f(m._1), f(m._2)); }
    $runtime.fail();
  }
};
const functorBackendSyntax = {
  map: f => m => {
    if (m.tag === "Var") { return $BackendSyntax("Var", m._1); }
    if (m.tag === "Local") { return $BackendSyntax("Local", m._1, m._2); }
    if (m.tag === "Lit") { return $BackendSyntax("Lit", PureScript$dBackend$dOptimizer$dCoreFn.functorLiteral.map(f)(m._1)); }
    if (m.tag === "App") { return $BackendSyntax("App", f(m._1), Data$dFunctor.arrayMap(f)(m._2)); }
    if (m.tag === "Abs") { return $BackendSyntax("Abs", m._1, f(m._2)); }
    if (m.tag === "UncurriedApp") { return $BackendSyntax("UncurriedApp", f(m._1), Data$dFunctor.arrayMap(f)(m._2)); }
    if (m.tag === "UncurriedAbs") { return $BackendSyntax("UncurriedAbs", m._1, f(m._2)); }
    if (m.tag === "UncurriedEffectApp") { return $BackendSyntax("UncurriedEffectApp", f(m._1), Data$dFunctor.arrayMap(f)(m._2)); }
    if (m.tag === "UncurriedEffectAbs") { return $BackendSyntax("UncurriedEffectAbs", m._1, f(m._2)); }
    if (m.tag === "Accessor") { return $BackendSyntax("Accessor", f(m._1), m._2); }
    if (m.tag === "Update") { return $BackendSyntax("Update", f(m._1), Data$dFunctor.arrayMap(m$1 => PureScript$dBackend$dOptimizer$dCoreFn.$Prop(m$1._1, f(m$1._2)))(m._2)); }
    if (m.tag === "CtorSaturated") { return $BackendSyntax("CtorSaturated", m._1, m._2, m._3, m._4, Data$dFunctor.arrayMap(m$1 => Data$dTuple.$Tuple(m$1._1, f(m$1._2)))(m._5)); }
    if (m.tag === "CtorDef") { return $BackendSyntax("CtorDef", m._1, m._2, m._3, m._4); }
    if (m.tag === "LetRec") { return $BackendSyntax("LetRec", m._1, Data$dFunctor.arrayMap(m$1 => Data$dTuple.$Tuple(m$1._1, f(m$1._2)))(m._2), f(m._3)); }
    if (m.tag === "Let") { return $BackendSyntax("Let", m._1, m._2, f(m._3), f(m._4)); }
    if (m.tag === "EffectBind") { return $BackendSyntax("EffectBind", m._1, m._2, f(m._3), f(m._4)); }
    if (m.tag === "EffectPure") { return $BackendSyntax("EffectPure", f(m._1)); }
    if (m.tag === "EffectDefer") { return $BackendSyntax("EffectDefer", f(m._1)); }
    if (m.tag === "Branch") { return $BackendSyntax("Branch", Data$dFunctor.arrayMap(m$1 => $Pair(f(m$1._1), f(m$1._2)))(m._1), f(m._2)); }
    if (m.tag === "PrimOp") {
      return $BackendSyntax(
        "PrimOp",
        (() => {
          if (m._1.tag === "Op1") { return $BackendOperator("Op1", m._1._1, f(m._1._2)); }
          if (m._1.tag === "Op2") { return $BackendOperator("Op2", m._1._1, f(m._1._2), f(m._1._3)); }
          $runtime.fail();
        })()
      );
    }
    if (m.tag === "PrimEffect") {
      return $BackendSyntax(
        "PrimEffect",
        (() => {
          if (m._1.tag === "EffectRefNew") { return $BackendEffect("EffectRefNew", f(m._1._1)); }
          if (m._1.tag === "EffectRefRead") { return $BackendEffect("EffectRefRead", f(m._1._1)); }
          if (m._1.tag === "EffectRefWrite") { return $BackendEffect("EffectRefWrite", f(m._1._1), f(m._1._2)); }
          $runtime.fail();
        })()
      );
    }
    if (m.tag === "PrimUndefined") { return PrimUndefined; }
    if (m.tag === "Fail") { return $BackendSyntax("Fail", m._1); }
    $runtime.fail();
  }
};
const foldablePair = {
  foldl: f => acc => v => f(f(acc)(v._1))(v._2),
  foldr: f => acc => v => f(v._1)(f(v._2)(acc)),
  foldMap: dictMonoid => f => v => dictMonoid.Semigroup0().append(f(v._1))(f(v._2))
};
const traversablePair = {
  sequence: dictApplicative => a => traversablePair.traverse(dictApplicative)(Data$dTraversable.identity)(a),
  traverse: dictApplicative => {
    const Apply0 = dictApplicative.Apply0();
    return f => v => Apply0.apply(Apply0.Functor0().map(Pair)(f(v._1)))(f(v._2));
  },
  Functor0: () => functorPair,
  Foldable1: () => foldablePair
};
const foldableBackendOperator = {
  foldr: a => Data$dFoldable.foldrDefault(foldableBackendOperator)(a),
  foldl: a => Data$dFoldable.foldlDefault(foldableBackendOperator)(a),
  foldMap: dictMonoid => f => v => {
    if (v.tag === "Op1") { return f(v._2); }
    if (v.tag === "Op2") { return dictMonoid.Semigroup0().append(f(v._2))(f(v._3)); }
    $runtime.fail();
  }
};
const traversableBackendOperato = {
  sequence: dictApplicative => a => traversableBackendOperato.traverse(dictApplicative)(Data$dTraversable.identity)(a),
  traverse: dictApplicative => {
    const Apply0 = dictApplicative.Apply0();
    const $0 = Apply0.Functor0();
    return f => v => {
      if (v.tag === "Op1") { return $0.map(Op1(v._1))(f(v._2)); }
      if (v.tag === "Op2") { return Apply0.apply($0.map(Op2(v._1))(f(v._2)))(f(v._3)); }
      $runtime.fail();
    };
  },
  Functor0: () => functorBackendOperator,
  Foldable1: () => foldableBackendOperator
};
const foldableBackendEffect = {
  foldr: a => Data$dFoldable.foldrDefault(foldableBackendEffect)(a),
  foldl: a => Data$dFoldable.foldlDefault(foldableBackendEffect)(a),
  foldMap: dictMonoid => f => v => {
    if (v.tag === "EffectRefNew") { return f(v._1); }
    if (v.tag === "EffectRefRead") { return f(v._1); }
    if (v.tag === "EffectRefWrite") { return dictMonoid.Semigroup0().append(f(v._1))(f(v._2)); }
    $runtime.fail();
  }
};
const foldableBackendSyntax = {
  foldr: a => Data$dFoldable.foldrDefault(foldableBackendSyntax)(a),
  foldl: a => Data$dFoldable.foldlDefault(foldableBackendSyntax)(a),
  foldMap: dictMonoid => {
    const mempty = dictMonoid.mempty;
    const foldMap7 = Data$dFoldable.foldableArray.foldMap(dictMonoid);
    const $0 = dictMonoid.Semigroup0();
    const foldMap9 = Data$dFoldable.foldableArray.foldMap(dictMonoid);
    return f => v => {
      if (v.tag === "Var") { return mempty; }
      if (v.tag === "Local") { return mempty; }
      if (v.tag === "Lit") {
        if (v._1.tag === "LitArray") { return foldMap7(f)(v._1._1); }
        if (v._1.tag === "LitRecord") { return foldMap7(v$1 => f(v$1._2))(v._1._1); }
        return mempty;
      }
      if (v.tag === "App") { return $0.append(f(v._1))(foldMap9(f)(v._2)); }
      if (v.tag === "Abs") { return f(v._2); }
      if (v.tag === "UncurriedApp") { return $0.append(f(v._1))(foldMap7(f)(v._2)); }
      if (v.tag === "UncurriedAbs") { return f(v._2); }
      if (v.tag === "UncurriedEffectApp") { return $0.append(f(v._1))(foldMap7(f)(v._2)); }
      if (v.tag === "UncurriedEffectAbs") { return f(v._2); }
      if (v.tag === "Accessor") { return f(v._1); }
      if (v.tag === "Update") { return $0.append(f(v._1))(foldMap7(v$1 => f(v$1._2))(v._2)); }
      if (v.tag === "LetRec") { return $0.append(foldMap9(v$1 => f(v$1._2))(v._2))(f(v._3)); }
      if (v.tag === "Let") { return $0.append(f(v._3))(f(v._4)); }
      if (v.tag === "EffectBind") { return $0.append(f(v._3))(f(v._4)); }
      if (v.tag === "EffectPure") { return f(v._1); }
      if (v.tag === "EffectDefer") { return f(v._1); }
      if (v.tag === "Branch") { return $0.append(foldMap9(v$1 => dictMonoid.Semigroup0().append(f(v$1._1))(f(v$1._2)))(v._1))(f(v._2)); }
      if (v.tag === "PrimOp") {
        if (v._1.tag === "Op1") { return f(v._1._2); }
        if (v._1.tag === "Op2") { return dictMonoid.Semigroup0().append(f(v._1._2))(f(v._1._3)); }
        $runtime.fail();
      }
      if (v.tag === "PrimEffect") {
        if (v._1.tag === "EffectRefNew") { return f(v._1._1); }
        if (v._1.tag === "EffectRefRead") { return f(v._1._1); }
        if (v._1.tag === "EffectRefWrite") { return dictMonoid.Semigroup0().append(f(v._1._1))(f(v._1._2)); }
        $runtime.fail();
      }
      if (v.tag === "PrimUndefined") { return mempty; }
      if (v.tag === "CtorSaturated") { return foldMap7(v$1 => f(v$1._2))(v._5); }
      if (v.tag === "CtorDef") { return mempty; }
      if (v.tag === "Fail") { return mempty; }
      $runtime.fail();
    };
  }
};
const traversableBackendEffect = {
  sequence: dictApplicative => a => traversableBackendEffect.traverse(dictApplicative)(Data$dTraversable.identity)(a),
  traverse: dictApplicative => {
    const Apply0 = dictApplicative.Apply0();
    const $0 = Apply0.Functor0();
    return f => v => {
      if (v.tag === "EffectRefNew") { return $0.map(EffectRefNew)(f(v._1)); }
      if (v.tag === "EffectRefRead") { return $0.map(EffectRefRead)(f(v._1)); }
      if (v.tag === "EffectRefWrite") { return Apply0.apply($0.map(EffectRefWrite)(f(v._1)))(f(v._2)); }
      $runtime.fail();
    };
  },
  Functor0: () => functorBackendEffect,
  Foldable1: () => foldableBackendEffect
};
const traversableBackendSyntax = {
  sequence: dictApplicative => a => traversableBackendSyntax.traverse(dictApplicative)(Data$dTraversable.identity)(a),
  traverse: dictApplicative => {
    const Apply0 = dictApplicative.Apply0();
    const $0 = Apply0.Functor0();
    const traverse7 = Data$dTraversable.traversableArray.traverse(dictApplicative);
    const traverse9 = Data$dTraversable.traversableArray.traverse(dictApplicative);
    const traverse11 = traversablePair.traverse(dictApplicative);
    const traverse12 = traversableBackendOperato.traverse(dictApplicative);
    const traverse13 = traversableBackendEffect.traverse(dictApplicative);
    return f => v => {
      if (v.tag === "Var") { return dictApplicative.pure($BackendSyntax("Var", v._1)); }
      if (v.tag === "Local") { return dictApplicative.pure($BackendSyntax("Local", v._1, v._2)); }
      if (v.tag === "Lit") {
        if (v._1.tag === "LitInt") { return dictApplicative.pure($BackendSyntax("Lit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitInt", v._1._1))); }
        if (v._1.tag === "LitNumber") { return dictApplicative.pure($BackendSyntax("Lit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitNumber", v._1._1))); }
        if (v._1.tag === "LitString") { return dictApplicative.pure($BackendSyntax("Lit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitString", v._1._1))); }
        if (v._1.tag === "LitChar") { return dictApplicative.pure($BackendSyntax("Lit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitChar", v._1._1))); }
        if (v._1.tag === "LitBoolean") { return dictApplicative.pure($BackendSyntax("Lit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitBoolean", v._1._1))); }
        if (v._1.tag === "LitArray") { return $0.map(x => $BackendSyntax("Lit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitArray", x)))(traverse7(f)(v._1._1)); }
        if (v._1.tag === "LitRecord") {
          return $0.map(x => $BackendSyntax("Lit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitRecord", x)))(traverse7(PureScript$dBackend$dOptimizer$dCoreFn.traversableProp.traverse(dictApplicative)(f))(v._1._1));
        }
        $runtime.fail();
      }
      if (v.tag === "App") { return Apply0.apply($0.map(App)(f(v._1)))(traverse9(f)(v._2)); }
      if (v.tag === "Abs") { return $0.map(Abs(v._1))(f(v._2)); }
      if (v.tag === "UncurriedApp") { return Apply0.apply($0.map(UncurriedApp)(f(v._1)))(traverse7(f)(v._2)); }
      if (v.tag === "UncurriedAbs") { return $0.map(UncurriedAbs(v._1))(f(v._2)); }
      if (v.tag === "UncurriedEffectApp") { return Apply0.apply($0.map(UncurriedEffectApp)(f(v._1)))(traverse7(f)(v._2)); }
      if (v.tag === "UncurriedEffectAbs") { return $0.map(UncurriedEffectAbs(v._1))(f(v._2)); }
      if (v.tag === "Accessor") {
        const $1 = v._2;
        return $0.map(a => $BackendSyntax("Accessor", a, $1))(f(v._1));
      }
      if (v.tag === "Update") {
        return Apply0.apply($0.map(Update)(f(v._1)))(traverse7(PureScript$dBackend$dOptimizer$dCoreFn.traversableProp.traverse(dictApplicative)(f))(v._2));
      }
      if (v.tag === "CtorDef") { return dictApplicative.pure($BackendSyntax("CtorDef", v._1, v._2, v._3, v._4)); }
      if (v.tag === "CtorSaturated") { return $0.map(CtorSaturated(v._1)(v._2)(v._3)(v._4))(traverse7(Data$dTraversable.traversableTuple.traverse(dictApplicative)(f))(v._5)); }
      if (v.tag === "LetRec") { return Apply0.apply($0.map(LetRec(v._1))(traverse9(Data$dTraversable.traversableTuple.traverse(dictApplicative)(f))(v._2)))(f(v._3)); }
      if (v.tag === "Let") { return Apply0.apply($0.map(Let(v._1)(v._2))(f(v._3)))(f(v._4)); }
      if (v.tag === "EffectBind") { return Apply0.apply($0.map(EffectBind(v._1)(v._2))(f(v._3)))(f(v._4)); }
      if (v.tag === "EffectPure") { return $0.map(EffectPure)(f(v._1)); }
      if (v.tag === "EffectDefer") { return $0.map(EffectDefer)(f(v._1)); }
      if (v.tag === "Branch") { return Apply0.apply($0.map(Branch)(traverse9(traverse11(f))(v._1)))(f(v._2)); }
      if (v.tag === "PrimOp") { return $0.map(PrimOp)(traverse12(f)(v._1)); }
      if (v.tag === "PrimEffect") { return $0.map(PrimEffect)(traverse13(f)(v._1)); }
      if (v.tag === "PrimUndefined") { return dictApplicative.pure(PrimUndefined); }
      if (v.tag === "Fail") { return dictApplicative.pure($BackendSyntax("Fail", v._1)); }
      $runtime.fail();
    };
  },
  Functor0: () => functorBackendSyntax,
  Foldable1: () => foldableBackendSyntax
};
const eqPair = dictEq => ({eq: x => y => dictEq.eq(x._1)(y._1) && dictEq.eq(x._2)(y._2)});
const eqLevel = Data$dEq.eqInt;
const eqTuple2 = {eq: x => y => (x._1.tag === "Nothing" ? y._1.tag === "Nothing" : x._1.tag === "Just" && y._1.tag === "Just" && x._1._1 === y._1._1) && x._2 === y._2};
const eq9 = /* #__PURE__ */ (() => Data$dEq.eqArrayImpl(eqTuple2.eq))();
const eq10 = /* #__PURE__ */ (() => Data$dEq.eqArrayImpl(eqTuple2.eq))();
const eqBackendOperatorOrd = {
  eq: x => y => {
    if (x === "OpEq") { return y === "OpEq"; }
    if (x === "OpNotEq") { return y === "OpNotEq"; }
    if (x === "OpGt") { return y === "OpGt"; }
    if (x === "OpGte") { return y === "OpGte"; }
    if (x === "OpLt") { return y === "OpLt"; }
    return x === "OpLte" && y === "OpLte";
  }
};
const ordBackendOperatorOrd = {
  compare: x => y => {
    if (x === "OpEq") {
      if (y === "OpEq") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "OpEq") { return Data$dOrdering.GT; }
    if (x === "OpNotEq") {
      if (y === "OpNotEq") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "OpNotEq") { return Data$dOrdering.GT; }
    if (x === "OpGt") {
      if (y === "OpGt") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "OpGt") { return Data$dOrdering.GT; }
    if (x === "OpGte") {
      if (y === "OpGte") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "OpGte") { return Data$dOrdering.GT; }
    if (x === "OpLt") {
      if (y === "OpLt") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "OpLt") { return Data$dOrdering.GT; }
    if (x === "OpLte" && y === "OpLte") { return Data$dOrdering.EQ; }
    $runtime.fail();
  },
  Eq0: () => eqBackendOperatorOrd
};
const eqBackendOperatorNum = {
  eq: x => y => {
    if (x === "OpAdd") { return y === "OpAdd"; }
    if (x === "OpDivide") { return y === "OpDivide"; }
    if (x === "OpMultiply") { return y === "OpMultiply"; }
    return x === "OpSubtract" && y === "OpSubtract";
  }
};
const ordBackendOperatorNum = {
  compare: x => y => {
    if (x === "OpAdd") {
      if (y === "OpAdd") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "OpAdd") { return Data$dOrdering.GT; }
    if (x === "OpDivide") {
      if (y === "OpDivide") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "OpDivide") { return Data$dOrdering.GT; }
    if (x === "OpMultiply") {
      if (y === "OpMultiply") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "OpMultiply") { return Data$dOrdering.GT; }
    if (x === "OpSubtract" && y === "OpSubtract") { return Data$dOrdering.EQ; }
    $runtime.fail();
  },
  Eq0: () => eqBackendOperatorNum
};
const eqBackendOperator2 = {
  eq: x => y => {
    if (x.tag === "OpArrayIndex") { return y.tag === "OpArrayIndex"; }
    if (x.tag === "OpBooleanAnd") { return y.tag === "OpBooleanAnd"; }
    if (x.tag === "OpBooleanOr") { return y.tag === "OpBooleanOr"; }
    if (x.tag === "OpBooleanOrd") {
      return y.tag === "OpBooleanOrd" && (() => {
        if (x._1 === "OpEq") { return y._1 === "OpEq"; }
        if (x._1 === "OpNotEq") { return y._1 === "OpNotEq"; }
        if (x._1 === "OpGt") { return y._1 === "OpGt"; }
        if (x._1 === "OpGte") { return y._1 === "OpGte"; }
        if (x._1 === "OpLt") { return y._1 === "OpLt"; }
        return x._1 === "OpLte" && y._1 === "OpLte";
      })();
    }
    if (x.tag === "OpCharOrd") {
      return y.tag === "OpCharOrd" && (() => {
        if (x._1 === "OpEq") { return y._1 === "OpEq"; }
        if (x._1 === "OpNotEq") { return y._1 === "OpNotEq"; }
        if (x._1 === "OpGt") { return y._1 === "OpGt"; }
        if (x._1 === "OpGte") { return y._1 === "OpGte"; }
        if (x._1 === "OpLt") { return y._1 === "OpLt"; }
        return x._1 === "OpLte" && y._1 === "OpLte";
      })();
    }
    if (x.tag === "OpIntBitAnd") { return y.tag === "OpIntBitAnd"; }
    if (x.tag === "OpIntBitOr") { return y.tag === "OpIntBitOr"; }
    if (x.tag === "OpIntBitShiftLeft") { return y.tag === "OpIntBitShiftLeft"; }
    if (x.tag === "OpIntBitShiftRight") { return y.tag === "OpIntBitShiftRight"; }
    if (x.tag === "OpIntBitXor") { return y.tag === "OpIntBitXor"; }
    if (x.tag === "OpIntBitZeroFillShiftRight") { return y.tag === "OpIntBitZeroFillShiftRight"; }
    if (x.tag === "OpIntNum") {
      return y.tag === "OpIntNum" && (() => {
        if (x._1 === "OpAdd") { return y._1 === "OpAdd"; }
        if (x._1 === "OpDivide") { return y._1 === "OpDivide"; }
        if (x._1 === "OpMultiply") { return y._1 === "OpMultiply"; }
        return x._1 === "OpSubtract" && y._1 === "OpSubtract";
      })();
    }
    if (x.tag === "OpIntOrd") {
      return y.tag === "OpIntOrd" && (() => {
        if (x._1 === "OpEq") { return y._1 === "OpEq"; }
        if (x._1 === "OpNotEq") { return y._1 === "OpNotEq"; }
        if (x._1 === "OpGt") { return y._1 === "OpGt"; }
        if (x._1 === "OpGte") { return y._1 === "OpGte"; }
        if (x._1 === "OpLt") { return y._1 === "OpLt"; }
        return x._1 === "OpLte" && y._1 === "OpLte";
      })();
    }
    if (x.tag === "OpNumberNum") {
      return y.tag === "OpNumberNum" && (() => {
        if (x._1 === "OpAdd") { return y._1 === "OpAdd"; }
        if (x._1 === "OpDivide") { return y._1 === "OpDivide"; }
        if (x._1 === "OpMultiply") { return y._1 === "OpMultiply"; }
        return x._1 === "OpSubtract" && y._1 === "OpSubtract";
      })();
    }
    if (x.tag === "OpNumberOrd") {
      return y.tag === "OpNumberOrd" && (() => {
        if (x._1 === "OpEq") { return y._1 === "OpEq"; }
        if (x._1 === "OpNotEq") { return y._1 === "OpNotEq"; }
        if (x._1 === "OpGt") { return y._1 === "OpGt"; }
        if (x._1 === "OpGte") { return y._1 === "OpGte"; }
        if (x._1 === "OpLt") { return y._1 === "OpLt"; }
        return x._1 === "OpLte" && y._1 === "OpLte";
      })();
    }
    if (x.tag === "OpStringAppend") { return y.tag === "OpStringAppend"; }
    return x.tag === "OpStringOrd" && y.tag === "OpStringOrd" && (() => {
      if (x._1 === "OpEq") { return y._1 === "OpEq"; }
      if (x._1 === "OpNotEq") { return y._1 === "OpNotEq"; }
      if (x._1 === "OpGt") { return y._1 === "OpGt"; }
      if (x._1 === "OpGte") { return y._1 === "OpGte"; }
      if (x._1 === "OpLt") { return y._1 === "OpLt"; }
      return x._1 === "OpLte" && y._1 === "OpLte";
    })();
  }
};
const ordBackendOperator2 = {
  compare: x => y => {
    if (x.tag === "OpArrayIndex") {
      if (y.tag === "OpArrayIndex") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y.tag === "OpArrayIndex") { return Data$dOrdering.GT; }
    if (x.tag === "OpBooleanAnd") {
      if (y.tag === "OpBooleanAnd") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y.tag === "OpBooleanAnd") { return Data$dOrdering.GT; }
    if (x.tag === "OpBooleanOr") {
      if (y.tag === "OpBooleanOr") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y.tag === "OpBooleanOr") { return Data$dOrdering.GT; }
    if (x.tag === "OpBooleanOrd") {
      if (y.tag === "OpBooleanOrd") {
        if (x._1 === "OpEq") {
          if (y._1 === "OpEq") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "OpEq") { return Data$dOrdering.GT; }
        if (x._1 === "OpNotEq") {
          if (y._1 === "OpNotEq") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "OpNotEq") { return Data$dOrdering.GT; }
        if (x._1 === "OpGt") {
          if (y._1 === "OpGt") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "OpGt") { return Data$dOrdering.GT; }
        if (x._1 === "OpGte") {
          if (y._1 === "OpGte") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "OpGte") { return Data$dOrdering.GT; }
        if (x._1 === "OpLt") {
          if (y._1 === "OpLt") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "OpLt") { return Data$dOrdering.GT; }
        if (x._1 === "OpLte" && y._1 === "OpLte") { return Data$dOrdering.EQ; }
        $runtime.fail();
      }
      return Data$dOrdering.LT;
    }
    if (y.tag === "OpBooleanOrd") { return Data$dOrdering.GT; }
    if (x.tag === "OpCharOrd") {
      if (y.tag === "OpCharOrd") {
        if (x._1 === "OpEq") {
          if (y._1 === "OpEq") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "OpEq") { return Data$dOrdering.GT; }
        if (x._1 === "OpNotEq") {
          if (y._1 === "OpNotEq") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "OpNotEq") { return Data$dOrdering.GT; }
        if (x._1 === "OpGt") {
          if (y._1 === "OpGt") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "OpGt") { return Data$dOrdering.GT; }
        if (x._1 === "OpGte") {
          if (y._1 === "OpGte") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "OpGte") { return Data$dOrdering.GT; }
        if (x._1 === "OpLt") {
          if (y._1 === "OpLt") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "OpLt") { return Data$dOrdering.GT; }
        if (x._1 === "OpLte" && y._1 === "OpLte") { return Data$dOrdering.EQ; }
        $runtime.fail();
      }
      return Data$dOrdering.LT;
    }
    if (y.tag === "OpCharOrd") { return Data$dOrdering.GT; }
    if (x.tag === "OpIntBitAnd") {
      if (y.tag === "OpIntBitAnd") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y.tag === "OpIntBitAnd") { return Data$dOrdering.GT; }
    if (x.tag === "OpIntBitOr") {
      if (y.tag === "OpIntBitOr") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y.tag === "OpIntBitOr") { return Data$dOrdering.GT; }
    if (x.tag === "OpIntBitShiftLeft") {
      if (y.tag === "OpIntBitShiftLeft") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y.tag === "OpIntBitShiftLeft") { return Data$dOrdering.GT; }
    if (x.tag === "OpIntBitShiftRight") {
      if (y.tag === "OpIntBitShiftRight") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y.tag === "OpIntBitShiftRight") { return Data$dOrdering.GT; }
    if (x.tag === "OpIntBitXor") {
      if (y.tag === "OpIntBitXor") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y.tag === "OpIntBitXor") { return Data$dOrdering.GT; }
    if (x.tag === "OpIntBitZeroFillShiftRight") {
      if (y.tag === "OpIntBitZeroFillShiftRight") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y.tag === "OpIntBitZeroFillShiftRight") { return Data$dOrdering.GT; }
    if (x.tag === "OpIntNum") {
      if (y.tag === "OpIntNum") {
        if (x._1 === "OpAdd") {
          if (y._1 === "OpAdd") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "OpAdd") { return Data$dOrdering.GT; }
        if (x._1 === "OpDivide") {
          if (y._1 === "OpDivide") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "OpDivide") { return Data$dOrdering.GT; }
        if (x._1 === "OpMultiply") {
          if (y._1 === "OpMultiply") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "OpMultiply") { return Data$dOrdering.GT; }
        if (x._1 === "OpSubtract" && y._1 === "OpSubtract") { return Data$dOrdering.EQ; }
        $runtime.fail();
      }
      return Data$dOrdering.LT;
    }
    if (y.tag === "OpIntNum") { return Data$dOrdering.GT; }
    if (x.tag === "OpIntOrd") {
      if (y.tag === "OpIntOrd") {
        if (x._1 === "OpEq") {
          if (y._1 === "OpEq") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "OpEq") { return Data$dOrdering.GT; }
        if (x._1 === "OpNotEq") {
          if (y._1 === "OpNotEq") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "OpNotEq") { return Data$dOrdering.GT; }
        if (x._1 === "OpGt") {
          if (y._1 === "OpGt") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "OpGt") { return Data$dOrdering.GT; }
        if (x._1 === "OpGte") {
          if (y._1 === "OpGte") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "OpGte") { return Data$dOrdering.GT; }
        if (x._1 === "OpLt") {
          if (y._1 === "OpLt") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "OpLt") { return Data$dOrdering.GT; }
        if (x._1 === "OpLte" && y._1 === "OpLte") { return Data$dOrdering.EQ; }
        $runtime.fail();
      }
      return Data$dOrdering.LT;
    }
    if (y.tag === "OpIntOrd") { return Data$dOrdering.GT; }
    if (x.tag === "OpNumberNum") {
      if (y.tag === "OpNumberNum") {
        if (x._1 === "OpAdd") {
          if (y._1 === "OpAdd") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "OpAdd") { return Data$dOrdering.GT; }
        if (x._1 === "OpDivide") {
          if (y._1 === "OpDivide") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "OpDivide") { return Data$dOrdering.GT; }
        if (x._1 === "OpMultiply") {
          if (y._1 === "OpMultiply") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "OpMultiply") { return Data$dOrdering.GT; }
        if (x._1 === "OpSubtract" && y._1 === "OpSubtract") { return Data$dOrdering.EQ; }
        $runtime.fail();
      }
      return Data$dOrdering.LT;
    }
    if (y.tag === "OpNumberNum") { return Data$dOrdering.GT; }
    if (x.tag === "OpNumberOrd") {
      if (y.tag === "OpNumberOrd") {
        if (x._1 === "OpEq") {
          if (y._1 === "OpEq") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "OpEq") { return Data$dOrdering.GT; }
        if (x._1 === "OpNotEq") {
          if (y._1 === "OpNotEq") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "OpNotEq") { return Data$dOrdering.GT; }
        if (x._1 === "OpGt") {
          if (y._1 === "OpGt") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "OpGt") { return Data$dOrdering.GT; }
        if (x._1 === "OpGte") {
          if (y._1 === "OpGte") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "OpGte") { return Data$dOrdering.GT; }
        if (x._1 === "OpLt") {
          if (y._1 === "OpLt") { return Data$dOrdering.EQ; }
          return Data$dOrdering.LT;
        }
        if (y._1 === "OpLt") { return Data$dOrdering.GT; }
        if (x._1 === "OpLte" && y._1 === "OpLte") { return Data$dOrdering.EQ; }
        $runtime.fail();
      }
      return Data$dOrdering.LT;
    }
    if (y.tag === "OpNumberOrd") { return Data$dOrdering.GT; }
    if (x.tag === "OpStringAppend") {
      if (y.tag === "OpStringAppend") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y.tag === "OpStringAppend") { return Data$dOrdering.GT; }
    if (x.tag === "OpStringOrd" && y.tag === "OpStringOrd") {
      if (x._1 === "OpEq") {
        if (y._1 === "OpEq") { return Data$dOrdering.EQ; }
        return Data$dOrdering.LT;
      }
      if (y._1 === "OpEq") { return Data$dOrdering.GT; }
      if (x._1 === "OpNotEq") {
        if (y._1 === "OpNotEq") { return Data$dOrdering.EQ; }
        return Data$dOrdering.LT;
      }
      if (y._1 === "OpNotEq") { return Data$dOrdering.GT; }
      if (x._1 === "OpGt") {
        if (y._1 === "OpGt") { return Data$dOrdering.EQ; }
        return Data$dOrdering.LT;
      }
      if (y._1 === "OpGt") { return Data$dOrdering.GT; }
      if (x._1 === "OpGte") {
        if (y._1 === "OpGte") { return Data$dOrdering.EQ; }
        return Data$dOrdering.LT;
      }
      if (y._1 === "OpGte") { return Data$dOrdering.GT; }
      if (x._1 === "OpLt") {
        if (y._1 === "OpLt") { return Data$dOrdering.EQ; }
        return Data$dOrdering.LT;
      }
      if (y._1 === "OpLt") { return Data$dOrdering.GT; }
      if (x._1 === "OpLte" && y._1 === "OpLte") { return Data$dOrdering.EQ; }
    }
    $runtime.fail();
  },
  Eq0: () => eqBackendOperator2
};
const eqBackendOperator1 = {
  eq: x => y => {
    if (x.tag === "OpBooleanNot") { return y.tag === "OpBooleanNot"; }
    if (x.tag === "OpIntBitNot") { return y.tag === "OpIntBitNot"; }
    if (x.tag === "OpIntNegate") { return y.tag === "OpIntNegate"; }
    if (x.tag === "OpNumberNegate") { return y.tag === "OpNumberNegate"; }
    if (x.tag === "OpArrayLength") { return y.tag === "OpArrayLength"; }
    return x.tag === "OpIsTag" && y.tag === "OpIsTag" && (x._1._1.tag === "Nothing"
      ? y._1._1.tag === "Nothing"
      : x._1._1.tag === "Just" && y._1._1.tag === "Just" && x._1._1._1 === y._1._1._1) && x._1._2 === y._1._2;
  }
};
const ordBackendOperator1 = {
  compare: x => y => {
    if (x.tag === "OpBooleanNot") {
      if (y.tag === "OpBooleanNot") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y.tag === "OpBooleanNot") { return Data$dOrdering.GT; }
    if (x.tag === "OpIntBitNot") {
      if (y.tag === "OpIntBitNot") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y.tag === "OpIntBitNot") { return Data$dOrdering.GT; }
    if (x.tag === "OpIntNegate") {
      if (y.tag === "OpIntNegate") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y.tag === "OpIntNegate") { return Data$dOrdering.GT; }
    if (x.tag === "OpNumberNegate") {
      if (y.tag === "OpNumberNegate") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y.tag === "OpNumberNegate") { return Data$dOrdering.GT; }
    if (x.tag === "OpArrayLength") {
      if (y.tag === "OpArrayLength") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y.tag === "OpArrayLength") { return Data$dOrdering.GT; }
    if (x.tag === "OpIsTag" && y.tag === "OpIsTag") { return compare(x._1)(y._1); }
    $runtime.fail();
  },
  Eq0: () => eqBackendOperator1
};
const eqBackendOperator = dictEq => (
  {
    eq: x => y => {
      if (x.tag === "Op1") {
        return y.tag === "Op1" && (() => {
          if (x._1.tag === "OpBooleanNot") { return y._1.tag === "OpBooleanNot"; }
          if (x._1.tag === "OpIntBitNot") { return y._1.tag === "OpIntBitNot"; }
          if (x._1.tag === "OpIntNegate") { return y._1.tag === "OpIntNegate"; }
          if (x._1.tag === "OpNumberNegate") { return y._1.tag === "OpNumberNegate"; }
          if (x._1.tag === "OpArrayLength") { return y._1.tag === "OpArrayLength"; }
          return x._1.tag === "OpIsTag" && y._1.tag === "OpIsTag" && (x._1._1._1.tag === "Nothing"
            ? y._1._1._1.tag === "Nothing"
            : x._1._1._1.tag === "Just" && y._1._1._1.tag === "Just" && x._1._1._1._1 === y._1._1._1._1) && x._1._1._2 === y._1._1._2;
        })() && dictEq.eq(x._2)(y._2);
      }
      return x.tag === "Op2" && y.tag === "Op2" && eqBackendOperator2.eq(x._1)(y._1) && dictEq.eq(x._2)(y._2) && dictEq.eq(x._3)(y._3);
    }
  }
);
const eqBackendEffect = dictEq => (
  {
    eq: x => y => {
      if (x.tag === "EffectRefNew") { return y.tag === "EffectRefNew" && dictEq.eq(x._1)(y._1); }
      if (x.tag === "EffectRefRead") { return y.tag === "EffectRefRead" && dictEq.eq(x._1)(y._1); }
      return x.tag === "EffectRefWrite" && y.tag === "EffectRefWrite" && dictEq.eq(x._1)(y._1) && dictEq.eq(x._2)(y._2);
    }
  }
);
const eqBackendAccessor = {
  eq: x => y => {
    if (x.tag === "GetProp") { return y.tag === "GetProp" && x._1 === y._1; }
    if (x.tag === "GetIndex") { return y.tag === "GetIndex" && x._1 === y._1; }
    return x.tag === "GetCtorField" && y.tag === "GetCtorField" && (x._1._1.tag === "Nothing"
      ? y._1._1.tag === "Nothing"
      : x._1._1.tag === "Just" && y._1._1.tag === "Just" && x._1._1._1 === y._1._1._1) && x._1._2 === y._1._2 && (x._2 === "ProductType"
      ? y._2 === "ProductType"
      : x._2 === "SumType" && y._2 === "SumType") && x._3 === y._3 && x._4 === y._4 && x._5 === y._5 && x._6 === y._6;
  }
};
const eqBackendSyntax = dictEq => {
  const $0 = Data$dEq.eqArrayImpl(dictEq.eq);
  const eq20 = Data$dEq.eqArrayImpl(x => y => x._1 === y._1 && dictEq.eq(x._2)(y._2));
  const eq21 = Data$dEq.eqArrayImpl(x => y => x._1 === y._1 && dictEq.eq(x._2)(y._2));
  const eq22 = Data$dEq.eqArrayImpl(x => y => x._1 === y._1 && dictEq.eq(x._2)(y._2));
  const eq23 = Data$dEq.eqArrayImpl(x => y => dictEq.eq(x._1)(y._1) && dictEq.eq(x._2)(y._2));
  return {
    eq: x => y => {
      if (x.tag === "Var") {
        return y.tag === "Var" && (x._1._1.tag === "Nothing" ? y._1._1.tag === "Nothing" : x._1._1.tag === "Just" && y._1._1.tag === "Just" && x._1._1._1 === y._1._1._1) && x._1._2 === y._1._2;
      }
      if (x.tag === "Local") {
        return y.tag === "Local" && (x._1.tag === "Nothing" ? y._1.tag === "Nothing" : x._1.tag === "Just" && y._1.tag === "Just" && x._1._1 === y._1._1) && x._2 === y._2;
      }
      if (x.tag === "Lit") { return y.tag === "Lit" && PureScript$dBackend$dOptimizer$dCoreFn.eqLiteral(dictEq).eq(x._1)(y._1); }
      if (x.tag === "App") { return y.tag === "App" && dictEq.eq(x._1)(y._1) && Data$dEq.eqArrayImpl(dictEq.eq)(x._2)(y._2); }
      if (x.tag === "Abs") { return y.tag === "Abs" && eq9(x._1)(y._1) && dictEq.eq(x._2)(y._2); }
      if (x.tag === "UncurriedApp") { return y.tag === "UncurriedApp" && dictEq.eq(x._1)(y._1) && $0(x._2)(y._2); }
      if (x.tag === "UncurriedAbs") { return y.tag === "UncurriedAbs" && eq10(x._1)(y._1) && dictEq.eq(x._2)(y._2); }
      if (x.tag === "UncurriedEffectApp") { return y.tag === "UncurriedEffectApp" && dictEq.eq(x._1)(y._1) && $0(x._2)(y._2); }
      if (x.tag === "UncurriedEffectAbs") { return y.tag === "UncurriedEffectAbs" && eq10(x._1)(y._1) && dictEq.eq(x._2)(y._2); }
      if (x.tag === "Accessor") { return y.tag === "Accessor" && dictEq.eq(x._1)(y._1) && eqBackendAccessor.eq(x._2)(y._2); }
      if (x.tag === "Update") { return y.tag === "Update" && dictEq.eq(x._1)(y._1) && eq20(x._2)(y._2); }
      if (x.tag === "CtorSaturated") {
        return y.tag === "CtorSaturated" && (x._1._1.tag === "Nothing" ? y._1._1.tag === "Nothing" : x._1._1.tag === "Just" && y._1._1.tag === "Just" && x._1._1._1 === y._1._1._1) && x._1._2 === y._1._2 && (x._2 === "ProductType"
          ? y._2 === "ProductType"
          : x._2 === "SumType" && y._2 === "SumType") && x._3 === y._3 && x._4 === y._4 && eq21(x._5)(y._5);
      }
      if (x.tag === "CtorDef") {
        return y.tag === "CtorDef" && (x._1 === "ProductType" ? y._1 === "ProductType" : x._1 === "SumType" && y._1 === "SumType") && x._2 === y._2 && x._3 === y._3 && eq7(x._4)(y._4);
      }
      if (x.tag === "LetRec") { return y.tag === "LetRec" && x._1 === y._1 && eq22(x._2)(y._2) && dictEq.eq(x._3)(y._3); }
      if (x.tag === "Let") {
        return y.tag === "Let" && (x._1.tag === "Nothing" ? y._1.tag === "Nothing" : x._1.tag === "Just" && y._1.tag === "Just" && x._1._1 === y._1._1) && x._2 === y._2 && dictEq.eq(x._3)(y._3) && dictEq.eq(x._4)(y._4);
      }
      if (x.tag === "EffectBind") {
        return y.tag === "EffectBind" && (x._1.tag === "Nothing" ? y._1.tag === "Nothing" : x._1.tag === "Just" && y._1.tag === "Just" && x._1._1 === y._1._1) && x._2 === y._2 && dictEq.eq(x._3)(y._3) && dictEq.eq(x._4)(y._4);
      }
      if (x.tag === "EffectPure") { return y.tag === "EffectPure" && dictEq.eq(x._1)(y._1); }
      if (x.tag === "EffectDefer") { return y.tag === "EffectDefer" && dictEq.eq(x._1)(y._1); }
      if (x.tag === "Branch") { return y.tag === "Branch" && eq23(x._1)(y._1) && dictEq.eq(x._2)(y._2); }
      if (x.tag === "PrimOp") { return y.tag === "PrimOp" && eqBackendOperator(dictEq).eq(x._1)(y._1); }
      if (x.tag === "PrimEffect") {
        return y.tag === "PrimEffect" && (() => {
          if (x._1.tag === "EffectRefNew") { return y._1.tag === "EffectRefNew" && dictEq.eq(x._1._1)(y._1._1); }
          if (x._1.tag === "EffectRefRead") { return y._1.tag === "EffectRefRead" && dictEq.eq(x._1._1)(y._1._1); }
          return x._1.tag === "EffectRefWrite" && y._1.tag === "EffectRefWrite" && dictEq.eq(x._1._1)(y._1._1) && dictEq.eq(x._1._2)(y._1._2);
        })();
      }
      if (x.tag === "PrimUndefined") { return y.tag === "PrimUndefined"; }
      return x.tag === "Fail" && y.tag === "Fail" && x._1 === y._1;
    }
  };
};
const ordBackendAccessor = {
  compare: x => y => {
    if (x.tag === "GetProp") {
      if (y.tag === "GetProp") { return Data$dOrd.ordString.compare(x._1)(y._1); }
      return Data$dOrdering.LT;
    }
    if (y.tag === "GetProp") { return Data$dOrdering.GT; }
    if (x.tag === "GetIndex") {
      if (y.tag === "GetIndex") { return Data$dOrd.ordInt.compare(x._1)(y._1); }
      return Data$dOrdering.LT;
    }
    if (y.tag === "GetIndex") { return Data$dOrdering.GT; }
    if (x.tag === "GetCtorField" && y.tag === "GetCtorField") {
      const v = compare(x._1)(y._1);
      if (v === "LT") { return Data$dOrdering.LT; }
      if (v === "GT") { return Data$dOrdering.GT; }
      if (x._2 === "ProductType") {
        if (y._2 === "ProductType") {
          const v2 = Data$dOrd.ordString.compare(x._3)(y._3);
          if (v2 === "LT") { return Data$dOrdering.LT; }
          if (v2 === "GT") { return Data$dOrdering.GT; }
          const v3 = Data$dOrd.ordString.compare(x._4)(y._4);
          if (v3 === "LT") { return Data$dOrdering.LT; }
          if (v3 === "GT") { return Data$dOrdering.GT; }
          const v4 = Data$dOrd.ordString.compare(x._5)(y._5);
          if (v4 === "LT") { return Data$dOrdering.LT; }
          if (v4 === "GT") { return Data$dOrdering.GT; }
          return Data$dOrd.ordInt.compare(x._6)(y._6);
        }
        return Data$dOrdering.LT;
      }
      if (y._2 === "ProductType") { return Data$dOrdering.GT; }
      if (x._2 === "SumType" && y._2 === "SumType") {
        const v2 = Data$dOrd.ordString.compare(x._3)(y._3);
        if (v2 === "LT") { return Data$dOrdering.LT; }
        if (v2 === "GT") { return Data$dOrdering.GT; }
        const v3 = Data$dOrd.ordString.compare(x._4)(y._4);
        if (v3 === "LT") { return Data$dOrdering.LT; }
        if (v3 === "GT") { return Data$dOrdering.GT; }
        const v4 = Data$dOrd.ordString.compare(x._5)(y._5);
        if (v4 === "LT") { return Data$dOrdering.LT; }
        if (v4 === "GT") { return Data$dOrdering.GT; }
        return Data$dOrd.ordInt.compare(x._6)(y._6);
      }
    }
    $runtime.fail();
  },
  Eq0: () => eqBackendAccessor
};
const syntaxOf = dict => dict.syntaxOf;
const sndPair = v => v._2;
const fstPair = v => v._1;
export {
  $BackendAccessor,
  $BackendEffect,
  $BackendOperator,
  $BackendOperator1,
  $BackendOperator2,
  $BackendOperatorNum,
  $BackendOperatorOrd,
  $BackendSyntax,
  $Pair,
  Abs,
  Accessor,
  App,
  Branch,
  CtorDef,
  CtorSaturated,
  EffectBind,
  EffectDefer,
  EffectPure,
  EffectRefNew,
  EffectRefRead,
  EffectRefWrite,
  Fail,
  GetCtorField,
  GetIndex,
  GetProp,
  Let,
  LetRec,
  Level,
  Lit,
  Local,
  Op1,
  Op2,
  OpAdd,
  OpArrayIndex,
  OpArrayLength,
  OpBooleanAnd,
  OpBooleanNot,
  OpBooleanOr,
  OpBooleanOrd,
  OpCharOrd,
  OpDivide,
  OpEq,
  OpGt,
  OpGte,
  OpIntBitAnd,
  OpIntBitNot,
  OpIntBitOr,
  OpIntBitShiftLeft,
  OpIntBitShiftRight,
  OpIntBitXor,
  OpIntBitZeroFillShiftRight,
  OpIntNegate,
  OpIntNum,
  OpIntOrd,
  OpIsTag,
  OpLt,
  OpLte,
  OpMultiply,
  OpNotEq,
  OpNumberNegate,
  OpNumberNum,
  OpNumberOrd,
  OpStringAppend,
  OpStringOrd,
  OpSubtract,
  Pair,
  PrimEffect,
  PrimOp,
  PrimUndefined,
  UncurriedAbs,
  UncurriedApp,
  UncurriedEffectAbs,
  UncurriedEffectApp,
  Update,
  Var,
  compare,
  eq10,
  eq7,
  eq9,
  eqBackendAccessor,
  eqBackendEffect,
  eqBackendOperator,
  eqBackendOperator1,
  eqBackendOperator2,
  eqBackendOperatorNum,
  eqBackendOperatorOrd,
  eqBackendSyntax,
  eqLevel,
  eqPair,
  eqTuple2,
  foldableBackendEffect,
  foldableBackendOperator,
  foldableBackendSyntax,
  foldablePair,
  fstPair,
  functorBackendEffect,
  functorBackendOperator,
  functorBackendSyntax,
  functorPair,
  newtypeLevel_,
  ordBackendAccessor,
  ordBackendOperator1,
  ordBackendOperator2,
  ordBackendOperatorNum,
  ordBackendOperatorOrd,
  ordLevel,
  sndPair,
  syntaxOf,
  traversableBackendEffect,
  traversableBackendOperato,
  traversableBackendSyntax,
  traversablePair
};
