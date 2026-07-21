import * as $runtime from "../runtime.js";
import * as Data$dArray from "../Data.Array/index.js";
import * as Data$dEq from "../Data.Eq/index.js";
import * as Data$dFoldable from "../Data.Foldable/index.js";
import * as Data$dFunctor from "../Data.Functor/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dOrd from "../Data.Ord/index.js";
import * as Data$dOrdering from "../Data.Ordering/index.js";
import * as Data$dSemiring from "../Data.Semiring/index.js";
import * as Data$dString$dCodeUnits from "../Data.String.CodeUnits/index.js";
import * as Data$dTraversable from "../Data.Traversable/index.js";
import * as Type$dProxy from "../Type.Proxy/index.js";
const $Bind = (tag, _1) => ({tag, _1});
const $Binder = (tag, _1, _2, _3, _4) => ({tag, _1, _2, _3, _4});
const $Binding = (_1, _2, _3) => ({tag: "Binding", _1, _2, _3});
const $CaseAlternative = (_1, _2) => ({tag: "CaseAlternative", _1, _2});
const $CaseGuard = (tag, _1) => ({tag, _1});
const $Comment = (tag, _1) => ({tag, _1});
const $ConstructorType = tag => tag;
const $Expr = (tag, _1, _2, _3, _4) => ({tag, _1, _2, _3, _4});
const $Guard = (_1, _2) => ({tag: "Guard", _1, _2});
const $Import = (_1, _2) => ({tag: "Import", _1, _2});
const $Literal = (tag, _1) => ({tag, _1});
const $Meta = (tag, _1, _2) => ({tag, _1, _2});
const $Prop = (_1, _2) => ({tag: "Prop", _1, _2});
const $Qualified = (_1, _2) => ({tag: "Qualified", _1, _2});
const $ReExport = (_1, _2) => ({tag: "ReExport", _1, _2});
const zero = /* #__PURE__ */ (() => Data$dSemiring.semiringRecordCons({reflectSymbol: () => "column"})()(Data$dSemiring.semiringRecordCons({reflectSymbol: () => "line"})()(Data$dSemiring.semiringRecordNil)(Data$dSemiring.semiringInt))(Data$dSemiring.semiringInt).zeroRecord(Type$dProxy.Proxy)(Type$dProxy.Proxy))();
const ProperName = x => x;
const Prop = value0 => value1 => $Prop(value0, value1);
const ModuleName = x => x;
const Qualified = value0 => value1 => $Qualified(value0, value1);
const LitInt = value0 => $Literal("LitInt", value0);
const LitNumber = value0 => $Literal("LitNumber", value0);
const LitString = value0 => $Literal("LitString", value0);
const LitChar = value0 => $Literal("LitChar", value0);
const LitBoolean = value0 => $Literal("LitBoolean", value0);
const LitArray = value0 => $Literal("LitArray", value0);
const LitRecord = value0 => $Literal("LitRecord", value0);
const Import = value0 => value1 => $Import(value0, value1);
const Ident = x => x;
const ReExport = value0 => value1 => $ReExport(value0, value1);
const ProductType = /* #__PURE__ */ $ConstructorType("ProductType");
const SumType = /* #__PURE__ */ $ConstructorType("SumType");
const IsConstructor = value0 => value1 => $Meta("IsConstructor", value0, value1);
const IsNewtype = /* #__PURE__ */ $Meta("IsNewtype");
const IsTypeClassConstructor = /* #__PURE__ */ $Meta("IsTypeClassConstructor");
const IsForeign = /* #__PURE__ */ $Meta("IsForeign");
const IsWhere = /* #__PURE__ */ $Meta("IsWhere");
const IsSyntheticApp = /* #__PURE__ */ $Meta("IsSyntheticApp");
const LineComment = value0 => $Comment("LineComment", value0);
const BlockComment = value0 => $Comment("BlockComment", value0);
const BinderNull = value0 => $Binder("BinderNull", value0);
const BinderVar = value0 => value1 => $Binder("BinderVar", value0, value1);
const BinderNamed = value0 => value1 => value2 => $Binder("BinderNamed", value0, value1, value2);
const BinderLit = value0 => value1 => $Binder("BinderLit", value0, value1);
const BinderConstructor = value0 => value1 => value2 => value3 => $Binder("BinderConstructor", value0, value1, value2, value3);
const NonRec = value0 => $Bind("NonRec", value0);
const Rec = value0 => $Bind("Rec", value0);
const Binding = value0 => value1 => value2 => $Binding(value0, value1, value2);
const ExprVar = value0 => value1 => $Expr("ExprVar", value0, value1);
const ExprLit = value0 => value1 => $Expr("ExprLit", value0, value1);
const ExprConstructor = value0 => value1 => value2 => value3 => $Expr("ExprConstructor", value0, value1, value2, value3);
const ExprAccessor = value0 => value1 => value2 => $Expr("ExprAccessor", value0, value1, value2);
const ExprUpdate = value0 => value1 => value2 => $Expr("ExprUpdate", value0, value1, value2);
const ExprAbs = value0 => value1 => value2 => $Expr("ExprAbs", value0, value1, value2);
const ExprApp = value0 => value1 => value2 => $Expr("ExprApp", value0, value1, value2);
const ExprCase = value0 => value1 => value2 => $Expr("ExprCase", value0, value1, value2);
const ExprLet = value0 => value1 => value2 => $Expr("ExprLet", value0, value1, value2);
const CaseAlternative = value0 => value1 => $CaseAlternative(value0, value1);
const Unconditional = value0 => $CaseGuard("Unconditional", value0);
const Guarded = value0 => $CaseGuard("Guarded", value0);
const Guard = value0 => value1 => $Guard(value0, value1);
const Module = x => x;
const Ann = x => x;
const newtypeModuleName_ = {Coercible0: () => {}};
const newtypeIdent_ = {Coercible0: () => {}};
const functorQualified = {map: f => m => $Qualified(m._1, f(m._2))};
const eqProp = dictEq => ({eq: x => y => x._1 === y._1 && dictEq.eq(x._2)(y._2)});
const eqLiteral = dictEq => {
  const eq11 = Data$dEq.eqArrayImpl(x => y => x._1 === y._1 && dictEq.eq(x._2)(y._2));
  return {
    eq: x => y => {
      if (x.tag === "LitInt") { return y.tag === "LitInt" && x._1 === y._1; }
      if (x.tag === "LitNumber") { return y.tag === "LitNumber" && x._1 === y._1; }
      if (x.tag === "LitString") { return y.tag === "LitString" && x._1 === y._1; }
      if (x.tag === "LitChar") { return y.tag === "LitChar" && x._1 === y._1; }
      if (x.tag === "LitBoolean") { return y.tag === "LitBoolean" && x._1 === y._1; }
      if (x.tag === "LitArray") { return y.tag === "LitArray" && Data$dEq.eqArrayImpl(dictEq.eq)(x._1)(y._1); }
      return x.tag === "LitRecord" && y.tag === "LitRecord" && eq11(x._1)(y._1);
    }
  };
};
const unQualified = v => v._2;
const qualifiedModuleName = v => v._1;
const propValue = v => v._2;
const propKey = v => v._1;
const ordProperName = Data$dOrd.ordString;
const ordModuleName = Data$dOrd.ordString;
const compare = x => y => {
  if (x.tag === "Nothing") {
    if (y.tag === "Nothing") { return Data$dOrdering.EQ; }
    return Data$dOrdering.LT;
  }
  if (y.tag === "Nothing") { return Data$dOrdering.GT; }
  if (x.tag === "Just" && y.tag === "Just") { return Data$dOrd.ordString.compare(x._1)(y._1); }
  $runtime.fail();
};
const ordIdent = Data$dOrd.ordString;
const compare3 = /* #__PURE__ */ (() => Data$dOrd.ordArray(Data$dOrd.ordString).compare)();
const moduleName = v => v.name;
const litChildren = v => {
  if (v.tag === "LitArray") { return v._1; }
  if (v.tag === "LitRecord") { return Data$dFunctor.arrayMap(propValue)(v._1); }
  return [];
};
const isPrimModule = v => v === "Prim" || Data$dString$dCodeUnits.take(5)(v) === "Prim.";
const importName = v => v._2;
const functorProp = {map: f => m => $Prop(m._1, f(m._2))};
const functorLiteral = {
  map: f => m => {
    if (m.tag === "LitInt") { return $Literal("LitInt", m._1); }
    if (m.tag === "LitNumber") { return $Literal("LitNumber", m._1); }
    if (m.tag === "LitString") { return $Literal("LitString", m._1); }
    if (m.tag === "LitChar") { return $Literal("LitChar", m._1); }
    if (m.tag === "LitBoolean") { return $Literal("LitBoolean", m._1); }
    if (m.tag === "LitArray") { return $Literal("LitArray", Data$dFunctor.arrayMap(f)(m._1)); }
    if (m.tag === "LitRecord") { return $Literal("LitRecord", Data$dFunctor.arrayMap(m$1 => $Prop(m$1._1, f(m$1._2)))(m._1)); }
    $runtime.fail();
  }
};
const functorImport = {map: f => m => $Import(f(m._1), m._2)};
const functorBinder = {
  map: f => m => {
    if (m.tag === "BinderNull") { return $Binder("BinderNull", f(m._1)); }
    if (m.tag === "BinderVar") { return $Binder("BinderVar", f(m._1), m._2); }
    if (m.tag === "BinderNamed") { return $Binder("BinderNamed", f(m._1), m._2, functorBinder.map(f)(m._3)); }
    if (m.tag === "BinderLit") { return $Binder("BinderLit", f(m._1), functorLiteral.map(functorBinder.map(f))(m._2)); }
    if (m.tag === "BinderConstructor") { return $Binder("BinderConstructor", f(m._1), m._2, m._3, Data$dFunctor.arrayMap(functorBinder.map(f))(m._4)); }
    $runtime.fail();
  }
};
const functorGuard = {map: f => m => $Guard(functorExpr.map(f)(m._1), functorExpr.map(f)(m._2))};
const functorExpr = {
  map: f => m => {
    if (m.tag === "ExprVar") { return $Expr("ExprVar", f(m._1), m._2); }
    if (m.tag === "ExprLit") { return $Expr("ExprLit", f(m._1), functorLiteral.map(functorExpr.map(f))(m._2)); }
    if (m.tag === "ExprConstructor") { return $Expr("ExprConstructor", f(m._1), m._2, m._3, m._4); }
    if (m.tag === "ExprAccessor") { return $Expr("ExprAccessor", f(m._1), functorExpr.map(f)(m._2), m._3); }
    if (m.tag === "ExprUpdate") {
      return $Expr(
        "ExprUpdate",
        f(m._1),
        functorExpr.map(f)(m._2),
        Data$dFunctor.arrayMap((() => {
          const $0 = functorExpr.map(f);
          return m$1 => $Prop(m$1._1, $0(m$1._2));
        })())(m._3)
      );
    }
    if (m.tag === "ExprAbs") { return $Expr("ExprAbs", f(m._1), m._2, functorExpr.map(f)(m._3)); }
    if (m.tag === "ExprApp") { return $Expr("ExprApp", f(m._1), functorExpr.map(f)(m._2), functorExpr.map(f)(m._3)); }
    if (m.tag === "ExprCase") { return $Expr("ExprCase", f(m._1), Data$dFunctor.arrayMap(functorExpr.map(f))(m._2), Data$dFunctor.arrayMap(functorCaseAlternative.map(f))(m._3)); }
    if (m.tag === "ExprLet") { return $Expr("ExprLet", f(m._1), Data$dFunctor.arrayMap(functorBind.map(f))(m._2), functorExpr.map(f)(m._3)); }
    $runtime.fail();
  }
};
const functorCaseGuard = {
  map: f => m => {
    if (m.tag === "Unconditional") { return $CaseGuard("Unconditional", functorExpr.map(f)(m._1)); }
    if (m.tag === "Guarded") { return $CaseGuard("Guarded", Data$dFunctor.arrayMap(functorGuard.map(f))(m._1)); }
    $runtime.fail();
  }
};
const functorCaseAlternative = {map: f => m => $CaseAlternative(Data$dFunctor.arrayMap(functorBinder.map(f))(m._1), functorCaseGuard.map(f)(m._2))};
const functorBinding = {map: f => m => $Binding(f(m._1), m._2, functorExpr.map(f)(m._3))};
const functorBind = {
  map: f => m => {
    if (m.tag === "NonRec") { return $Bind("NonRec", functorBinding.map(f)(m._1)); }
    if (m.tag === "Rec") { return $Bind("Rec", Data$dFunctor.arrayMap(functorBinding.map(f))(m._1)); }
    $runtime.fail();
  }
};
const foldableProp = {foldl: k => a => v => k(a)(v._2), foldr: k => b => v => k(v._2)(b), foldMap: dictMonoid => k => v => k(v._2)};
const traversableProp = {
  traverse: dictApplicative => k => v => dictApplicative.Apply0().Functor0().map(Prop(v._1))(k(v._2)),
  sequence: dictApplicative => v => dictApplicative.Apply0().Functor0().map(Prop(v._1))(v._2),
  Functor0: () => functorProp,
  Foldable1: () => foldableProp
};
const foldableLiteral = {
  foldl: k => Data$dFoldable.foldlDefault(foldableLiteral)(k),
  foldr: k => Data$dFoldable.foldrDefault(foldableLiteral)(k),
  foldMap: dictMonoid => {
    const foldMap2 = Data$dFoldable.foldableArray.foldMap(dictMonoid);
    const mempty = dictMonoid.mempty;
    return k => v => {
      if (v.tag === "LitArray") { return foldMap2(k)(v._1); }
      if (v.tag === "LitRecord") { return foldMap2(v$1 => k(v$1._2))(v._1); }
      return mempty;
    };
  }
};
const traversableLiteral = {
  traverse: dictApplicative => {
    const $0 = dictApplicative.Apply0().Functor0();
    const traverse2 = Data$dTraversable.traversableArray.traverse(dictApplicative);
    return k => v => {
      if (v.tag === "LitArray") { return $0.map(LitArray)(traverse2(k)(v._1)); }
      if (v.tag === "LitRecord") { return $0.map(LitRecord)(traverse2(traversableProp.traverse(dictApplicative)(k))(v._1)); }
      if (v.tag === "LitInt") { return dictApplicative.pure($Literal("LitInt", v._1)); }
      if (v.tag === "LitNumber") { return dictApplicative.pure($Literal("LitNumber", v._1)); }
      if (v.tag === "LitString") { return dictApplicative.pure($Literal("LitString", v._1)); }
      if (v.tag === "LitChar") { return dictApplicative.pure($Literal("LitChar", v._1)); }
      if (v.tag === "LitBoolean") { return dictApplicative.pure($Literal("LitBoolean", v._1)); }
      $runtime.fail();
    };
  },
  sequence: dictApplicative => a => traversableLiteral.traverse(dictApplicative)(Data$dTraversable.identity)(a),
  Functor0: () => functorLiteral,
  Foldable1: () => foldableLiteral
};
const findProp = prop => Data$dArray.findMap(v => {
  if (prop === v._1) { return Data$dMaybe.$Maybe("Just", v._2); }
  return Data$dMaybe.Nothing;
});
const exprAnn = v => {
  if (v.tag === "ExprVar") { return v._1; }
  if (v.tag === "ExprLit") { return v._1; }
  if (v.tag === "ExprConstructor") { return v._1; }
  if (v.tag === "ExprAccessor") { return v._1; }
  if (v.tag === "ExprUpdate") { return v._1; }
  if (v.tag === "ExprAbs") { return v._1; }
  if (v.tag === "ExprApp") { return v._1; }
  if (v.tag === "ExprCase") { return v._1; }
  if (v.tag === "ExprLet") { return v._1; }
  $runtime.fail();
};
const eqProperName = Data$dEq.eqString;
const eqModuleName = Data$dEq.eqString;
const eqQualified = dictEq => (
  {eq: x => y => (x._1.tag === "Nothing" ? y._1.tag === "Nothing" : x._1.tag === "Just" && y._1.tag === "Just" && x._1._1 === y._1._1) && dictEq.eq(x._2)(y._2)}
);
const ordQualified = dictOrd => {
  const $0 = dictOrd.Eq0();
  const eqQualified1 = {eq: x => y => (x._1.tag === "Nothing" ? y._1.tag === "Nothing" : x._1.tag === "Just" && y._1.tag === "Just" && x._1._1 === y._1._1) && $0.eq(x._2)(y._2)};
  return {
    compare: x => y => {
      const v = compare(x._1)(y._1);
      if (v === "LT") { return Data$dOrdering.LT; }
      if (v === "GT") { return Data$dOrdering.GT; }
      return dictOrd.compare(x._2)(y._2);
    },
    Eq0: () => eqQualified1
  };
};
const eqIdent = Data$dEq.eqString;
const eq8 = /* #__PURE__ */ Data$dEq.eqArrayImpl(Data$dEq.eqStringImpl);
const eqReExport = {eq: x => y => x._1 === y._1 && x._2 === y._2};
const ordReExport = {
  compare: x => y => {
    const v = Data$dOrd.ordString.compare(x._1)(y._1);
    if (v === "LT") { return Data$dOrdering.LT; }
    if (v === "GT") { return Data$dOrdering.GT; }
    return Data$dOrd.ordString.compare(x._2)(y._2);
  },
  Eq0: () => eqReExport
};
const eqConstructorType = {
  eq: x => y => {
    if (x === "ProductType") { return y === "ProductType"; }
    return x === "SumType" && y === "SumType";
  }
};
const eqMeta = {
  eq: x => y => {
    if (x.tag === "IsConstructor") {
      return y.tag === "IsConstructor" && (x._1 === "ProductType" ? y._1 === "ProductType" : x._1 === "SumType" && y._1 === "SumType") && eq8(x._2)(y._2);
    }
    if (x.tag === "IsNewtype") { return y.tag === "IsNewtype"; }
    if (x.tag === "IsTypeClassConstructor") { return y.tag === "IsTypeClassConstructor"; }
    if (x.tag === "IsForeign") { return y.tag === "IsForeign"; }
    if (x.tag === "IsWhere") { return y.tag === "IsWhere"; }
    return x.tag === "IsSyntheticApp" && y.tag === "IsSyntheticApp";
  }
};
const ordConstructorType = {
  compare: x => y => {
    if (x === "ProductType") {
      if (y === "ProductType") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "ProductType") { return Data$dOrdering.GT; }
    if (x === "SumType" && y === "SumType") { return Data$dOrdering.EQ; }
    $runtime.fail();
  },
  Eq0: () => eqConstructorType
};
const ordMeta = {
  compare: x => y => {
    if (x.tag === "IsConstructor") {
      if (y.tag === "IsConstructor") {
        if (x._1 === "ProductType") {
          if (y._1 === "ProductType") { return compare3(x._2)(y._2); }
          return Data$dOrdering.LT;
        }
        if (y._1 === "ProductType") { return Data$dOrdering.GT; }
        if (x._1 === "SumType" && y._1 === "SumType") { return compare3(x._2)(y._2); }
        $runtime.fail();
      }
      return Data$dOrdering.LT;
    }
    if (y.tag === "IsConstructor") { return Data$dOrdering.GT; }
    if (x.tag === "IsNewtype") {
      if (y.tag === "IsNewtype") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y.tag === "IsNewtype") { return Data$dOrdering.GT; }
    if (x.tag === "IsTypeClassConstructor") {
      if (y.tag === "IsTypeClassConstructor") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y.tag === "IsTypeClassConstructor") { return Data$dOrdering.GT; }
    if (x.tag === "IsForeign") {
      if (y.tag === "IsForeign") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y.tag === "IsForeign") { return Data$dOrdering.GT; }
    if (x.tag === "IsWhere") {
      if (y.tag === "IsWhere") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y.tag === "IsWhere") { return Data$dOrdering.GT; }
    if (x.tag === "IsSyntheticApp" && y.tag === "IsSyntheticApp") { return Data$dOrdering.EQ; }
    $runtime.fail();
  },
  Eq0: () => eqMeta
};
const emptySpan = {path: "<internal>", start: zero, end: zero};
const emptyAnn = {span: emptySpan, meta: Data$dMaybe.Nothing};
export {
  $Bind,
  $Binder,
  $Binding,
  $CaseAlternative,
  $CaseGuard,
  $Comment,
  $ConstructorType,
  $Expr,
  $Guard,
  $Import,
  $Literal,
  $Meta,
  $Prop,
  $Qualified,
  $ReExport,
  Ann,
  BinderConstructor,
  BinderLit,
  BinderNamed,
  BinderNull,
  BinderVar,
  Binding,
  BlockComment,
  CaseAlternative,
  ExprAbs,
  ExprAccessor,
  ExprApp,
  ExprCase,
  ExprConstructor,
  ExprLet,
  ExprLit,
  ExprUpdate,
  ExprVar,
  Guard,
  Guarded,
  Ident,
  Import,
  IsConstructor,
  IsForeign,
  IsNewtype,
  IsSyntheticApp,
  IsTypeClassConstructor,
  IsWhere,
  LineComment,
  LitArray,
  LitBoolean,
  LitChar,
  LitInt,
  LitNumber,
  LitRecord,
  LitString,
  Module,
  ModuleName,
  NonRec,
  ProductType,
  Prop,
  ProperName,
  Qualified,
  ReExport,
  Rec,
  SumType,
  Unconditional,
  compare,
  compare3,
  emptyAnn,
  emptySpan,
  eq8,
  eqConstructorType,
  eqIdent,
  eqLiteral,
  eqMeta,
  eqModuleName,
  eqProp,
  eqProperName,
  eqQualified,
  eqReExport,
  exprAnn,
  findProp,
  foldableLiteral,
  foldableProp,
  functorBind,
  functorBinder,
  functorBinding,
  functorCaseAlternative,
  functorCaseGuard,
  functorExpr,
  functorGuard,
  functorImport,
  functorLiteral,
  functorProp,
  functorQualified,
  importName,
  isPrimModule,
  litChildren,
  moduleName,
  newtypeIdent_,
  newtypeModuleName_,
  ordConstructorType,
  ordIdent,
  ordMeta,
  ordModuleName,
  ordProperName,
  ordQualified,
  ordReExport,
  propKey,
  propValue,
  qualifiedModuleName,
  traversableLiteral,
  traversableProp,
  unQualified,
  zero
};
