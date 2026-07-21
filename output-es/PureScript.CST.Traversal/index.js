import * as $runtime from "../runtime.js";
import * as Control$dMonad$dFree from "../Control.Monad.Free/index.js";
import * as Data$dBitraversable from "../Data.Bitraversable/index.js";
import * as Data$dCatList from "../Data.CatList/index.js";
import * as Data$dCatQueue from "../Data.CatQueue/index.js";
import * as Data$dConst from "../Data.Const/index.js";
import * as Data$dFunctor$dCompose from "../Data.Functor.Compose/index.js";
import * as Data$dIdentity from "../Data.Identity/index.js";
import * as Data$dList$dTypes from "../Data.List.Types/index.js";
import * as Data$dTraversable from "../Data.Traversable/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
import * as PureScript$dCST$dTypes from "../PureScript.CST.Types/index.js";
import * as Unsafe$dCoerce from "../Unsafe.Coerce/index.js";
const ltraverse = dictApplicative => {
  const bitraverse2 = Data$dBitraversable.bitraversableTuple.bitraverse(dictApplicative);
  const pure = dictApplicative.pure;
  return f => bitraverse2(f)(pure);
};
const applicativeReaderT = /* #__PURE__ */ (() => {
  const functorReaderT1 = {map: x => v => x$1 => x(v(x$1))};
  const applyReaderT1 = {apply: v => v1 => r => v(r)(v1(r)), Functor0: () => functorReaderT1};
  return {pure: x => v => x, Apply0: () => applyReaderT1};
})();
const applyCompose = dictApply1 => {
  const $0 = dictApply1.Functor0();
  const functorCompose2 = {
    map: f => v => {
      const $1 = $0.map(f);
      return Control$dMonad$dFree.$Free(
        v._1,
        (() => {
          if (v._2.tag === "CatNil") {
            return Data$dCatList.$CatList(
              "CatCons",
              x => Control$dMonad$dFree.$Free(Control$dMonad$dFree.$FreeView("Return", $1(x)), Data$dCatList.CatNil),
              Data$dCatQueue.$CatQueue(Data$dList$dTypes.Nil, Data$dList$dTypes.Nil)
            );
          }
          if (v._2.tag === "CatCons") {
            return Data$dCatList.$CatList(
              "CatCons",
              v._2._1,
              Data$dCatQueue.$CatQueue(
                v._2._2._1,
                Data$dList$dTypes.$List(
                  "Cons",
                  Data$dCatList.$CatList(
                    "CatCons",
                    x => Control$dMonad$dFree.$Free(Control$dMonad$dFree.$FreeView("Return", $1(x)), Data$dCatList.CatNil),
                    Data$dCatQueue.$CatQueue(Data$dList$dTypes.Nil, Data$dList$dTypes.Nil)
                  ),
                  v._2._2._2
                )
              )
            );
          }
          $runtime.fail();
        })()
      );
    }
  };
  return {
    apply: v => v1 => Control$dMonad$dFree.freeApply.apply(Control$dMonad$dFree.$Free(
      v._1,
      (() => {
        if (v._2.tag === "CatNil") {
          return Data$dCatList.$CatList(
            "CatCons",
            x => Control$dMonad$dFree.$Free(Control$dMonad$dFree.$FreeView("Return", dictApply1.apply(x)), Data$dCatList.CatNil),
            Data$dCatQueue.$CatQueue(Data$dList$dTypes.Nil, Data$dList$dTypes.Nil)
          );
        }
        if (v._2.tag === "CatCons") {
          return Data$dCatList.$CatList(
            "CatCons",
            v._2._1,
            Data$dCatQueue.$CatQueue(
              v._2._2._1,
              Data$dList$dTypes.$List(
                "Cons",
                Data$dCatList.$CatList(
                  "CatCons",
                  x => Control$dMonad$dFree.$Free(Control$dMonad$dFree.$FreeView("Return", dictApply1.apply(x)), Data$dCatList.CatNil),
                  Data$dCatQueue.$CatQueue(Data$dList$dTypes.Nil, Data$dList$dTypes.Nil)
                ),
                v._2._2._2
              )
            )
          );
        }
        $runtime.fail();
      })()
    ))(v1),
    Functor0: () => functorCompose2
  };
};
const applicativeCompose = /* #__PURE__ */ Data$dFunctor$dCompose.applicativeCompose(Control$dMonad$dFree.freeApplicative);
const identity = x => x;
const traverseWrapped = dictApplicative => k => v => dictApplicative.Apply0().Functor0().map(value => ({...v, value}))(k(v.value));
const traverseSeparated = dictApplicative => {
  const Apply0 = dictApplicative.Apply0();
  const traverse4 = Data$dTraversable.traversableArray.traverse(dictApplicative);
  return k => v => Apply0.apply(Apply0.Functor0().map(v1 => v2 => ({head: v1, tail: v2}))(k(v.head)))(traverse4(Data$dTraversable.traversableTuple.traverse(dictApplicative)(k))(v.tail));
};
const traverseRecordUpdate = dictApplicative => {
  const $0 = dictApplicative.Apply0().Functor0();
  const traverseSeparated1 = traverseSeparated(dictApplicative);
  return k => v => {
    if (v.tag === "RecordUpdateLeaf") { return $0.map(PureScript$dCST$dTypes.RecordUpdateLeaf(v._1)(v._2))(k.onExpr(v._3)); }
    if (v.tag === "RecordUpdateBranch") {
      const $1 = v._2;
      return $0.map(PureScript$dCST$dTypes.RecordUpdateBranch(v._1))(dictApplicative.Apply0().Functor0().map(value => ({...$1, value}))(traverseSeparated1(traverseRecordUpdate(dictApplicative)(k))($1.value)));
    }
    $runtime.fail();
  };
};
const traverseRecordLabeled = dictApplicative => k => v => {
  if (v.tag === "RecordPun") { return dictApplicative.pure(PureScript$dCST$dTypes.$RecordLabeled("RecordPun", v._1)); }
  if (v.tag === "RecordField") { return dictApplicative.Apply0().Functor0().map(PureScript$dCST$dTypes.RecordField(v._1)(v._2))(k(v._3)); }
  $runtime.fail();
};
const traverseRecordAccessor = dictApplicative => k => r => dictApplicative.Apply0().Functor0().map(v => ({...r, expr: v}))(k.onExpr(r.expr));
const traversePatternGuard = dictApplicative => {
  const Apply0 = dictApplicative.Apply0();
  const ltraverse1 = ltraverse(dictApplicative);
  return k => v => Apply0.apply(Apply0.Functor0().map(binder => expr => ({binder, expr}))(Data$dTraversable.traversableMaybe.traverse(dictApplicative)(ltraverse1(k.onBinder))(v.binder)))(k.onExpr(v.expr));
};
const traverseModuleBody = dictApplicative => {
  const traverse4 = Data$dTraversable.traversableArray.traverse(dictApplicative);
  return k => v => dictApplicative.Apply0().Functor0().map(decls => ({...v, decls}))(traverse4(k.onDecl)(v.decls));
};
const traverseModule = dictApplicative => {
  const traverseModuleBody1 = traverseModuleBody(dictApplicative);
  return k => v => dictApplicative.Apply0().Functor0().map(body => ({...v, header: v.header, body}))(traverseModuleBody1(k)(v.body));
};
const traverseModule1 = /* #__PURE__ */ traverseModule(Control$dMonad$dFree.freeApplicative);
const traverseLambda = dictApplicative => {
  const Apply0 = dictApplicative.Apply0();
  const traverse4 = Data$dTraversable.traversableArray.traverse(dictApplicative);
  return k => l => Apply0.apply(Apply0.Functor0().map(v => v1 => ({...l, binders: v, body: v1}))(traverse4(k.onBinder)(l.binders)))(k.onExpr(l.body));
};
const traverseLabeled = dictApplicative => k => v => dictApplicative.Apply0().Functor0().map(value => ({...v, value}))(k(v.value));
const traverseRow = dictApplicative => {
  const Apply0 = dictApplicative.Apply0();
  const traverseSeparated1 = traverseSeparated(dictApplicative);
  return k => v => Apply0.apply(Apply0.Functor0().map(labels => tail => ({...v, labels, tail}))(Data$dTraversable.traversableMaybe.traverse(dictApplicative)(traverseSeparated1(v$1 => dictApplicative.Apply0().Functor0().map(value => (
    {...v$1, value}
  ))(k.onType(v$1.value))))(v.labels)))(Data$dTraversable.traversableMaybe.traverse(dictApplicative)(Data$dTraversable.traversableTuple.traverse(dictApplicative)(k.onType))(v.tail));
};
const traverseTypeVarBinding = dictApplicative => k => v => {
  if (v.tag === "TypeVarKinded") {
    const $0 = v._1;
    return dictApplicative.Apply0().Functor0().map(PureScript$dCST$dTypes.TypeVarKinded)(dictApplicative.Apply0().Functor0().map(value => ({...$0, value}))((() => {
      const $1 = $0.value;
      return dictApplicative.Apply0().Functor0().map(value => ({...$1, value}))(k.onType($1.value));
    })()));
  }
  if (v.tag === "TypeVarName") { return dictApplicative.pure(PureScript$dCST$dTypes.$TypeVarBinding("TypeVarName", v._1)); }
  $runtime.fail();
};
const traverseType = dictApplicative => {
  const Apply0 = dictApplicative.Apply0();
  const Functor0 = Apply0.Functor0();
  const traverseRow1 = traverseRow(dictApplicative);
  const traverse4 = Data$dTraversable.traversableArray.traverse(dictApplicative);
  return k => v => {
    if (v.tag === "TypeRow") {
      const $0 = v._1;
      return Functor0.map(PureScript$dCST$dTypes.TypeRow)(dictApplicative.Apply0().Functor0().map(value => ({...$0, value}))(traverseRow1(k)($0.value)));
    }
    if (v.tag === "TypeRecord") {
      const $0 = v._1;
      return Functor0.map(PureScript$dCST$dTypes.TypeRecord)(dictApplicative.Apply0().Functor0().map(value => ({...$0, value}))(traverseRow1(k)($0.value)));
    }
    if (v.tag === "TypeForall") {
      const $0 = v._3;
      return Apply0.apply(Functor0.map(f => f($0))(Functor0.map(PureScript$dCST$dTypes.TypeForall(v._1))(traverse4(traverseTypeVarBinding(dictApplicative)(k))(v._2))))(k.onType(v._4));
    }
    if (v.tag === "TypeKinded") {
      const $0 = v._2;
      return Apply0.apply(Functor0.map(f => f($0))(Functor0.map(PureScript$dCST$dTypes.TypeKinded)(k.onType(v._1))))(k.onType(v._3));
    }
    if (v.tag === "TypeApp") { return Apply0.apply(Functor0.map(PureScript$dCST$dTypes.TypeApp)(k.onType(v._1)))(traverse4(k.onType)(v._2)); }
    if (v.tag === "TypeOp") {
      return Apply0.apply(Functor0.map(PureScript$dCST$dTypes.TypeOp)(k.onType(v._1)))(traverse4(Data$dTraversable.traversableTuple.traverse(dictApplicative)(k.onType))(v._2));
    }
    if (v.tag === "TypeArrow") {
      const $0 = v._2;
      return Apply0.apply(Functor0.map(f => f($0))(Functor0.map(PureScript$dCST$dTypes.TypeArrow)(k.onType(v._1))))(k.onType(v._3));
    }
    if (v.tag === "TypeConstrained") {
      const $0 = v._2;
      return Apply0.apply(Functor0.map(f => f($0))(Functor0.map(PureScript$dCST$dTypes.TypeConstrained)(k.onType(v._1))))(k.onType(v._3));
    }
    if (v.tag === "TypeParens") {
      const $0 = v._1;
      return Functor0.map(PureScript$dCST$dTypes.TypeParens)(dictApplicative.Apply0().Functor0().map(value => ({...$0, value}))(k.onType($0.value)));
    }
    return dictApplicative.pure(v);
  };
};
const traverseType1 = /* #__PURE__ */ traverseType(applicativeReaderT);
const traverseType2 = /* #__PURE__ */ traverseType(Control$dMonad$dFree.freeApplicative);
const traverseIfThenElse = dictApplicative => {
  const Apply0 = dictApplicative.Apply0();
  return k => r => Apply0.apply(Apply0.apply(Apply0.Functor0().map(v => v1 => v2 => ({...r, cond: v, true: v1, false: v2}))(k.onExpr(r.cond)))(k.onExpr(r.true)))(k.onExpr(r.false));
};
const traverseWhere = dictApplicative => {
  const Apply0 = dictApplicative.Apply0();
  const traverse6 = Data$dTraversable.traversableArray.traverse(dictApplicative);
  return k => v => Apply0.apply(Apply0.Functor0().map(expr => bindings => ({expr, bindings}))(k.onExpr(v.expr)))(Data$dTraversable.traversableMaybe.traverse(dictApplicative)(Data$dTraversable.traversableTuple.traverse(dictApplicative)(traverse6(traverseLetBinding(dictApplicative)(k))))(v.bindings));
};
const traverseValueBindingFields = dictApplicative => {
  const Apply0 = dictApplicative.Apply0();
  const traverse4 = Data$dTraversable.traversableArray.traverse(dictApplicative);
  return k => v => Apply0.apply(Apply0.Functor0().map(v1 => v2 => ({...v, binders: v1, guarded: v2}))(traverse4(k.onBinder)(v.binders)))(traverseGuarded(dictApplicative)(k)(v.guarded));
};
const traverseLetBinding = dictApplicative => {
  const Apply0 = dictApplicative.Apply0();
  const Functor0 = Apply0.Functor0();
  return k => v => {
    if (v.tag === "LetBindingSignature") {
      const $0 = v._1;
      return Functor0.map(PureScript$dCST$dTypes.LetBindingSignature)(dictApplicative.Apply0().Functor0().map(value => ({...$0, value}))(k.onType($0.value)));
    }
    if (v.tag === "LetBindingName") { return Functor0.map(PureScript$dCST$dTypes.LetBindingName)(traverseValueBindingFields(dictApplicative)(k)(v._1)); }
    if (v.tag === "LetBindingPattern") {
      const $0 = v._2;
      return Apply0.apply(Functor0.map(f => f($0))(Functor0.map(PureScript$dCST$dTypes.LetBindingPattern)(k.onBinder(v._1))))(traverseWhere(dictApplicative)(k)(v._3));
    }
    if (v.tag === "LetBindingError") { return dictApplicative.pure(PureScript$dCST$dTypes.$LetBinding("LetBindingError", v._1)); }
    $runtime.fail();
  };
};
const traverseGuardedExpr = dictApplicative => {
  const Apply0 = dictApplicative.Apply0();
  const traverseSeparated1 = traverseSeparated(dictApplicative);
  const traversePatternGuard1 = traversePatternGuard(dictApplicative);
  return k => v => Apply0.apply(Apply0.Functor0().map(ps => wh => ({...v, patterns: ps, where: wh}))(traverseSeparated1(traversePatternGuard1(k))(v.patterns)))(traverseWhere(dictApplicative)(k)(v.where));
};
const traverseGuarded = dictApplicative => {
  const $0 = dictApplicative.Apply0().Functor0();
  const traverse4 = Data$dTraversable.traversableArray.traverse(dictApplicative);
  return k => v => {
    if (v.tag === "Unconditional") { return $0.map(PureScript$dCST$dTypes.Unconditional(v._1))(traverseWhere(dictApplicative)(k)(v._2)); }
    if (v.tag === "Guarded") { return $0.map(PureScript$dCST$dTypes.Guarded)(traverse4(traverseGuardedExpr(dictApplicative)(k))(v._1)); }
    $runtime.fail();
  };
};
const traverseInstanceBinding = dictApplicative => {
  const $0 = dictApplicative.Apply0().Functor0();
  const traverseValueBindingFields1 = traverseValueBindingFields(dictApplicative);
  return k => v => {
    if (v.tag === "InstanceBindingSignature") {
      const $1 = v._1;
      return $0.map(PureScript$dCST$dTypes.InstanceBindingSignature)(dictApplicative.Apply0().Functor0().map(value => ({...$1, value}))(k.onType($1.value)));
    }
    if (v.tag === "InstanceBindingName") { return $0.map(PureScript$dCST$dTypes.InstanceBindingName)(traverseValueBindingFields1(k)(v._1)); }
    $runtime.fail();
  };
};
const traverseLetIn = dictApplicative => {
  const Apply0 = dictApplicative.Apply0();
  const traverse4 = Data$dTraversable.traversableArray.traverse(dictApplicative);
  const traverseLetBinding1 = traverseLetBinding(dictApplicative);
  return k => l => Apply0.apply(Apply0.Functor0().map(v => v1 => ({...l, bindings: v, body: v1}))(traverse4(traverseLetBinding1(k))(l.bindings)))(k.onExpr(l.body));
};
const traverseForeign = dictApplicative => {
  const $0 = dictApplicative.Apply0().Functor0();
  return k => v => {
    if (v.tag === "ForeignValue") {
      const $1 = v._1;
      return $0.map(PureScript$dCST$dTypes.ForeignValue)(dictApplicative.Apply0().Functor0().map(value => ({...$1, value}))(k.onType($1.value)));
    }
    if (v.tag === "ForeignData") {
      const $1 = v._2;
      return $0.map(PureScript$dCST$dTypes.ForeignData(v._1))(dictApplicative.Apply0().Functor0().map(value => ({...$1, value}))(k.onType($1.value)));
    }
    if (v.tag === "ForeignKind") { return dictApplicative.pure(v); }
    $runtime.fail();
  };
};
const traverseExprAppSpine = dictApplicative => {
  const $0 = dictApplicative.Apply0().Functor0();
  return k => v => {
    if (v.tag === "AppType") { return $0.map(PureScript$dCST$dTypes.AppType(v._1))(k.onType(v._2)); }
    if (v.tag === "AppTerm") { return $0.map(PureScript$dCST$dTypes.AppTerm)(k.onExpr(v._1)); }
    $runtime.fail();
  };
};
const traverseDoStatement = dictApplicative => {
  const Apply0 = dictApplicative.Apply0();
  const Functor0 = Apply0.Functor0();
  const traverse4 = Data$dTraversable.traversableArray.traverse(dictApplicative);
  const traverseLetBinding1 = traverseLetBinding(dictApplicative);
  return k => v => {
    if (v.tag === "DoLet") { return Functor0.map(PureScript$dCST$dTypes.DoLet(v._1))(traverse4(traverseLetBinding1(k))(v._2)); }
    if (v.tag === "DoDiscard") { return Functor0.map(PureScript$dCST$dTypes.DoDiscard)(k.onExpr(v._1)); }
    if (v.tag === "DoBind") {
      const $0 = v._2;
      return Apply0.apply(Functor0.map(f => f($0))(Functor0.map(PureScript$dCST$dTypes.DoBind)(k.onBinder(v._1))))(k.onExpr(v._3));
    }
    if (v.tag === "DoError") { return dictApplicative.pure(PureScript$dCST$dTypes.$DoStatement("DoError", v._1)); }
    $runtime.fail();
  };
};
const traverseDoBlock = dictApplicative => {
  const traverse4 = Data$dTraversable.traversableArray.traverse(dictApplicative);
  const traverseDoStatement1 = traverseDoStatement(dictApplicative);
  return k => d => dictApplicative.Apply0().Functor0().map(v => ({...d, statements: v}))(traverse4(traverseDoStatement1(k))(d.statements));
};
const traverseDelimitedNonEmpty = dictApplicative => {
  const traverseSeparated1 = traverseSeparated(dictApplicative);
  return k => {
    const $0 = traverseSeparated1(k);
    return v => dictApplicative.Apply0().Functor0().map(value => ({...v, value}))($0(v.value));
  };
};
const traverseOneOrDelimited = dictApplicative => {
  const $0 = dictApplicative.Apply0().Functor0();
  const traverseDelimitedNonEmpty1 = traverseDelimitedNonEmpty(dictApplicative);
  return k => v => {
    if (v.tag === "One") { return $0.map(PureScript$dCST$dTypes.One)(k(v._1)); }
    if (v.tag === "Many") { return $0.map(PureScript$dCST$dTypes.Many)(traverseDelimitedNonEmpty1(k)(v._1)); }
    $runtime.fail();
  };
};
const traverseInstanceHead = dictApplicative => {
  const Apply0 = dictApplicative.Apply0();
  const ltraverse1 = ltraverse(dictApplicative);
  const traverseOneOrDelimited1 = traverseOneOrDelimited(dictApplicative);
  const traverse5 = Data$dTraversable.traversableArray.traverse(dictApplicative);
  return k => head => Apply0.apply(Apply0.Functor0().map(v => v1 => ({...head, constraints: v, types: v1}))(Data$dTraversable.traversableMaybe.traverse(dictApplicative)(ltraverse1(traverseOneOrDelimited1(k.onType)))(head.constraints)))(traverse5(k.onType)(head.types));
};
const traverseInstance = dictApplicative => {
  const Apply0 = dictApplicative.Apply0();
  const traverseInstanceHead1 = traverseInstanceHead(dictApplicative);
  const traverse6 = Data$dTraversable.traversableArray.traverse(dictApplicative);
  const traverseInstanceBinding1 = traverseInstanceBinding(dictApplicative);
  return k => v => Apply0.apply(Apply0.Functor0().map(head => body => ({...v, head, body}))(traverseInstanceHead1(k)(v.head)))(Data$dTraversable.traversableMaybe.traverse(dictApplicative)(Data$dTraversable.traversableTuple.traverse(dictApplicative)(traverse6(traverseInstanceBinding1(k))))(v.body));
};
const traverseDelimited = dictApplicative => {
  const traverseSeparated1 = traverseSeparated(dictApplicative);
  return k => {
    const $0 = Data$dTraversable.traversableMaybe.traverse(dictApplicative)(traverseSeparated1(k));
    return v => dictApplicative.Apply0().Functor0().map(value => ({...v, value}))($0(v.value));
  };
};
const traverseDataHead = dictApplicative => {
  const traverse4 = Data$dTraversable.traversableArray.traverse(dictApplicative);
  return k => head => dictApplicative.Apply0().Functor0().map(v => ({...head, vars: v}))(traverse4(traverseTypeVarBinding(dictApplicative)(k))(head.vars));
};
const traverseDataCtor = dictApplicative => {
  const traverse4 = Data$dTraversable.traversableArray.traverse(dictApplicative);
  return k => v => dictApplicative.Apply0().Functor0().map(fields => ({...v, fields}))(traverse4(k.onType)(v.fields));
};
const traverseClassHead = dictApplicative => {
  const Apply0 = dictApplicative.Apply0();
  const ltraverse1 = ltraverse(dictApplicative);
  const traverseOneOrDelimited1 = traverseOneOrDelimited(dictApplicative);
  const traverse5 = Data$dTraversable.traversableArray.traverse(dictApplicative);
  return k => head => Apply0.apply(Apply0.Functor0().map(v => v1 => ({...head, super: v, vars: v1}))(Data$dTraversable.traversableMaybe.traverse(dictApplicative)(ltraverse1(traverseOneOrDelimited1(k.onType)))(head.super)))(traverse5(traverseTypeVarBinding(dictApplicative)(k))(head.vars));
};
const traverseDecl = dictApplicative => {
  const Apply0 = dictApplicative.Apply0();
  const Functor0 = Apply0.Functor0();
  const traverseDataHead1 = traverseDataHead(dictApplicative);
  const traverseSeparated1 = traverseSeparated(dictApplicative);
  const traverseDataCtor1 = traverseDataCtor(dictApplicative);
  const traverseClassHead1 = traverseClassHead(dictApplicative);
  const traverse6 = Data$dTraversable.traversableArray.traverse(dictApplicative);
  const traverseInstance1 = traverseInstance(dictApplicative);
  const traverseInstanceHead1 = traverseInstanceHead(dictApplicative);
  const traverseValueBindingFields1 = traverseValueBindingFields(dictApplicative);
  const traverseForeign1 = traverseForeign(dictApplicative);
  return k => v => {
    if (v.tag === "DeclData") {
      return Apply0.apply(Functor0.map(PureScript$dCST$dTypes.DeclData)(traverseDataHead1(k)(v._1)))(Data$dTraversable.traversableMaybe.traverse(dictApplicative)(Data$dTraversable.traversableTuple.traverse(dictApplicative)(traverseSeparated1(traverseDataCtor1(k))))(v._2));
    }
    if (v.tag === "DeclType") {
      const $0 = v._2;
      return Apply0.apply(Functor0.map(f => f($0))(Functor0.map(PureScript$dCST$dTypes.DeclType)(traverseDataHead1(k)(v._1))))(k.onType(v._3));
    }
    if (v.tag === "DeclNewtype") {
      const $0 = v._3;
      const $1 = v._2;
      return Apply0.apply(Functor0.map(f => f($0))(Functor0.map(f => f($1))(Functor0.map(PureScript$dCST$dTypes.DeclNewtype)(traverseDataHead1(k)(v._1)))))(k.onType(v._4));
    }
    if (v.tag === "DeclClass") {
      return Apply0.apply(Functor0.map(PureScript$dCST$dTypes.DeclClass)(traverseClassHead1(k)(v._1)))(Data$dTraversable.traversableMaybe.traverse(dictApplicative)(Data$dTraversable.traversableTuple.traverse(dictApplicative)(traverse6(v$1 => dictApplicative.Apply0().Functor0().map(value => (
        {...v$1, value}
      ))(k.onType(v$1.value)))))(v._2));
    }
    if (v.tag === "DeclInstanceChain") { return Functor0.map(PureScript$dCST$dTypes.DeclInstanceChain)(traverseSeparated1(traverseInstance1(k))(v._1)); }
    if (v.tag === "DeclDerive") { return Functor0.map(PureScript$dCST$dTypes.DeclDerive(v._1)(v._2))(traverseInstanceHead1(k)(v._3)); }
    if (v.tag === "DeclKindSignature") {
      const $0 = v._2;
      return Functor0.map(PureScript$dCST$dTypes.DeclKindSignature(v._1))(dictApplicative.Apply0().Functor0().map(value => ({...$0, value}))(k.onType($0.value)));
    }
    if (v.tag === "DeclSignature") {
      const $0 = v._1;
      return Functor0.map(PureScript$dCST$dTypes.DeclSignature)(dictApplicative.Apply0().Functor0().map(value => ({...$0, value}))(k.onType($0.value)));
    }
    if (v.tag === "DeclValue") { return Functor0.map(PureScript$dCST$dTypes.DeclValue)(traverseValueBindingFields1(k)(v._1)); }
    if (v.tag === "DeclForeign") { return Functor0.map(PureScript$dCST$dTypes.DeclForeign(v._1)(v._2))(traverseForeign1(k)(v._3)); }
    return dictApplicative.pure(v);
  };
};
const traverseDecl1 = /* #__PURE__ */ traverseDecl(applicativeReaderT);
const traverseDecl2 = /* #__PURE__ */ traverseDecl(Control$dMonad$dFree.freeApplicative);
const traverseCaseOf = dictApplicative => {
  const Apply0 = dictApplicative.Apply0();
  const traverseSeparated1 = traverseSeparated(dictApplicative);
  const traverse4 = Data$dTraversable.traversableArray.traverse(dictApplicative);
  const bitraverse1 = Data$dBitraversable.bitraversableTuple.bitraverse(dictApplicative);
  const traverseGuarded1 = traverseGuarded(dictApplicative);
  return k => r => Apply0.apply(Apply0.Functor0().map(v => v1 => ({...r, head: v, branches: v1}))(traverseSeparated1(k.onExpr)(r.head)))(traverse4(bitraverse1(traverseSeparated1(k.onBinder))(traverseGuarded1(k)))(r.branches));
};
const traverseBinder = dictApplicative => {
  const Apply0 = dictApplicative.Apply0();
  const Functor0 = Apply0.Functor0();
  const traverse4 = Data$dTraversable.traversableArray.traverse(dictApplicative);
  const traverseDelimited1 = traverseDelimited(dictApplicative);
  const traverse5 = Data$dTraversable.traversableArray.traverse(dictApplicative);
  return k => v => {
    if (v.tag === "BinderNamed") { return Functor0.map(PureScript$dCST$dTypes.BinderNamed(v._1)(v._2))(k.onBinder(v._3)); }
    if (v.tag === "BinderConstructor") { return Functor0.map(PureScript$dCST$dTypes.BinderConstructor(v._1))(traverse4(k.onBinder)(v._2)); }
    if (v.tag === "BinderArray") { return Functor0.map(PureScript$dCST$dTypes.BinderArray)(traverseDelimited1(k.onBinder)(v._1)); }
    if (v.tag === "BinderRecord") { return Functor0.map(PureScript$dCST$dTypes.BinderRecord)(traverseDelimited1(traverseRecordLabeled(dictApplicative)(k.onBinder))(v._1)); }
    if (v.tag === "BinderParens") {
      const $0 = v._1;
      return Functor0.map(PureScript$dCST$dTypes.BinderParens)(dictApplicative.Apply0().Functor0().map(value => ({...$0, value}))(k.onBinder($0.value)));
    }
    if (v.tag === "BinderTyped") {
      const $0 = v._2;
      return Apply0.apply(Functor0.map(f => f($0))(Functor0.map(PureScript$dCST$dTypes.BinderTyped)(k.onBinder(v._1))))(k.onType(v._3));
    }
    if (v.tag === "BinderOp") {
      return Apply0.apply(Functor0.map(PureScript$dCST$dTypes.BinderOp)(k.onBinder(v._1)))(traverse5(Data$dTraversable.traversableTuple.traverse(dictApplicative)(k.onBinder))(v._2));
    }
    return dictApplicative.pure(v);
  };
};
const traverseBinder1 = /* #__PURE__ */ traverseBinder(applicativeReaderT);
const traverseBinder2 = /* #__PURE__ */ traverseBinder(Control$dMonad$dFree.freeApplicative);
const traverseAdoBlock = dictApplicative => {
  const Apply0 = dictApplicative.Apply0();
  const traverse4 = Data$dTraversable.traversableArray.traverse(dictApplicative);
  const traverseDoStatement1 = traverseDoStatement(dictApplicative);
  return k => a => Apply0.apply(Apply0.Functor0().map(v => v1 => ({...a, statements: v, result: v1}))(traverse4(traverseDoStatement1(k))(a.statements)))(k.onExpr(a.result));
};
const traverseExpr = dictApplicative => {
  const Apply0 = dictApplicative.Apply0();
  const Functor0 = Apply0.Functor0();
  const traverseDelimited1 = traverseDelimited(dictApplicative);
  const traverse4 = Data$dTraversable.traversableArray.traverse(dictApplicative);
  const bitraverse1 = Data$dBitraversable.bitraversableTuple.bitraverse(dictApplicative);
  const traverseSeparated1 = traverseSeparated(dictApplicative);
  const traverseRecordUpdate1 = traverseRecordUpdate(dictApplicative);
  const traverseExprAppSpine1 = traverseExprAppSpine(dictApplicative);
  const traverseLambda1 = traverseLambda(dictApplicative);
  const Apply0$1 = dictApplicative.Apply0();
  const traverseCaseOf1 = traverseCaseOf(dictApplicative);
  const traverseLetIn1 = traverseLetIn(dictApplicative);
  const traverseDoBlock1 = traverseDoBlock(dictApplicative);
  const traverseAdoBlock1 = traverseAdoBlock(dictApplicative);
  return k => v => {
    if (v.tag === "ExprArray") { return Functor0.map(PureScript$dCST$dTypes.ExprArray)(traverseDelimited1(k.onExpr)(v._1)); }
    if (v.tag === "ExprRecord") { return Functor0.map(PureScript$dCST$dTypes.ExprRecord)(traverseDelimited1(traverseRecordLabeled(dictApplicative)(k.onExpr))(v._1)); }
    if (v.tag === "ExprParens") {
      const $0 = v._1;
      return Functor0.map(PureScript$dCST$dTypes.ExprParens)(dictApplicative.Apply0().Functor0().map(value => ({...$0, value}))(k.onExpr($0.value)));
    }
    if (v.tag === "ExprTyped") {
      const $0 = v._2;
      return Apply0.apply(Functor0.map(f => f($0))(Functor0.map(PureScript$dCST$dTypes.ExprTyped)(k.onExpr(v._1))))(k.onType(v._3));
    }
    if (v.tag === "ExprInfix") {
      return Apply0.apply(Functor0.map(PureScript$dCST$dTypes.ExprInfix)(k.onExpr(v._1)))(traverse4(bitraverse1(v$1 => dictApplicative.Apply0().Functor0().map(value => (
        {...v$1, value}
      ))(k.onExpr(v$1.value)))(k.onExpr))(v._2));
    }
    if (v.tag === "ExprOp") {
      return Apply0.apply(Functor0.map(PureScript$dCST$dTypes.ExprOp)(k.onExpr(v._1)))(traverse4(Data$dTraversable.traversableTuple.traverse(dictApplicative)(k.onExpr))(v._2));
    }
    if (v.tag === "ExprNegate") { return Functor0.map(PureScript$dCST$dTypes.ExprNegate(v._1))(k.onExpr(v._2)); }
    if (v.tag === "ExprRecordAccessor") {
      const $0 = v._1;
      return Functor0.map(PureScript$dCST$dTypes.ExprRecordAccessor)(dictApplicative.Apply0().Functor0().map(v$1 => ({...$0, expr: v$1}))(k.onExpr($0.expr)));
    }
    if (v.tag === "ExprRecordUpdate") {
      const $0 = v._2;
      return Apply0.apply(Functor0.map(PureScript$dCST$dTypes.ExprRecordUpdate)(k.onExpr(v._1)))(dictApplicative.Apply0().Functor0().map(value => ({...$0, value}))(traverseSeparated1(traverseRecordUpdate1(k))($0.value)));
    }
    if (v.tag === "ExprApp") { return Apply0.apply(Functor0.map(PureScript$dCST$dTypes.ExprApp)(k.onExpr(v._1)))(traverse4(traverseExprAppSpine1(k))(v._2)); }
    if (v.tag === "ExprLambda") { return Functor0.map(PureScript$dCST$dTypes.ExprLambda)(traverseLambda1(k)(v._1)); }
    if (v.tag === "ExprIf") {
      return Functor0.map(PureScript$dCST$dTypes.ExprIf)((() => {
        const $0 = v._1;
        return Apply0$1.apply(Apply0$1.apply(Apply0$1.Functor0().map(v$1 => v1 => v2 => ({...$0, cond: v$1, true: v1, false: v2}))(k.onExpr($0.cond)))(k.onExpr($0.true)))(k.onExpr($0.false));
      })());
    }
    if (v.tag === "ExprCase") { return Functor0.map(PureScript$dCST$dTypes.ExprCase)(traverseCaseOf1(k)(v._1)); }
    if (v.tag === "ExprLet") { return Functor0.map(PureScript$dCST$dTypes.ExprLet)(traverseLetIn1(k)(v._1)); }
    if (v.tag === "ExprDo") { return Functor0.map(PureScript$dCST$dTypes.ExprDo)(traverseDoBlock1(k)(v._1)); }
    if (v.tag === "ExprAdo") { return Functor0.map(PureScript$dCST$dTypes.ExprAdo)(traverseAdoBlock1(k)(v._1)); }
    return dictApplicative.pure(v);
  };
};
const traverseExpr1 = /* #__PURE__ */ traverseExpr(applicativeReaderT);
const traverseExpr2 = /* #__PURE__ */ traverseExpr(Control$dMonad$dFree.freeApplicative);
const topDownTraversalWithContextM = dictMonad => {
  const $0 = dictMonad.Bind1();
  const $1 = dictMonad.Applicative0();
  const $2 = $1.Apply0();
  const $3 = $2.Functor0();
  const applicativeReaderT1 = (() => {
    const functorReaderT1 = {
      map: x => {
        const $4 = $3.map(x);
        return v => x$1 => $4(v(x$1));
      }
    };
    const applyReaderT1 = {apply: v => v1 => r => $2.apply(v(r))(v1(r)), Functor0: () => functorReaderT1};
    return {
      pure: x => {
        const $4 = $1.pure(x);
        return v => $4;
      },
      Apply0: () => applyReaderT1
    };
  })();
  const traverseBinder3 = traverseBinder(applicativeReaderT1);
  const traverseExpr3 = traverseExpr(applicativeReaderT1);
  const traverseDecl3 = traverseDecl(applicativeReaderT1);
  const traverseType3 = traverseType(applicativeReaderT1);
  return visitor => {
    const visitor$p = {
      onBinder: a => ctx => $0.bind(visitor.onBinder(ctx)(a))((() => {
        const $4 = traverseBinder3(visitor$p);
        return v => $4(v._2)(v._1);
      })()),
      onExpr: a => ctx => $0.bind(visitor.onExpr(ctx)(a))((() => {
        const $4 = traverseExpr3(visitor$p);
        return v => $4(v._2)(v._1);
      })()),
      onDecl: a => ctx => $0.bind(visitor.onDecl(ctx)(a))((() => {
        const $4 = traverseDecl3(visitor$p);
        return v => $4(v._2)(v._1);
      })()),
      onType: a => ctx => $0.bind(visitor.onType(ctx)(a))((() => {
        const $4 = traverseType3(visitor$p);
        return v => $4(v._2)(v._1);
      })())
    };
    return visitor$p;
  };
};
const topDownTraversalWithContext = visitor => {
  const visitor$p = {
    onBinder: a => ctx => {
      const $0 = visitor.onBinder(ctx)(a);
      return traverseBinder1(visitor$p)($0._2)($0._1);
    },
    onExpr: a => ctx => {
      const $0 = visitor.onExpr(ctx)(a);
      return traverseExpr1(visitor$p)($0._2)($0._1);
    },
    onDecl: a => ctx => {
      const $0 = visitor.onDecl(ctx)(a);
      return traverseDecl1(visitor$p)($0._2)($0._1);
    },
    onType: a => ctx => {
      const $0 = visitor.onType(ctx)(a);
      return traverseType1(visitor$p)($0._2)($0._1);
    }
  };
  return visitor$p;
};
const topDownTraversal = dictMonad => {
  const $0 = dictMonad.Bind1();
  const Applicative0 = dictMonad.Applicative0();
  const traverseBinder3 = traverseBinder(Applicative0);
  const traverseExpr3 = traverseExpr(Applicative0);
  const traverseType3 = traverseType(Applicative0);
  const traverseDecl3 = traverseDecl(Applicative0);
  return visitor => {
    const visitor$p = {
      onBinder: a => $0.bind(visitor.onBinder(a))(traverseBinder3(visitor$p)),
      onExpr: a => $0.bind(visitor.onExpr(a))(traverseExpr3(visitor$p)),
      onType: a => $0.bind(visitor.onType(a))(traverseType3(visitor$p)),
      onDecl: a => $0.bind(visitor.onDecl(a))(traverseDecl3(visitor$p))
    };
    return visitor$p;
  };
};
const topDownPureTraversal = visitor => {
  const visitor$p = {
    onBinder: a => Control$dMonad$dFree.$Free(
      Control$dMonad$dFree.$FreeView("Return", visitor.onBinder(a)),
      Data$dCatList.$CatList("CatCons", traverseBinder2(visitor$p), Data$dCatQueue.$CatQueue(Data$dList$dTypes.Nil, Data$dList$dTypes.Nil))
    ),
    onExpr: a => Control$dMonad$dFree.$Free(
      Control$dMonad$dFree.$FreeView("Return", visitor.onExpr(a)),
      Data$dCatList.$CatList("CatCons", traverseExpr2(visitor$p), Data$dCatQueue.$CatQueue(Data$dList$dTypes.Nil, Data$dList$dTypes.Nil))
    ),
    onType: a => Control$dMonad$dFree.$Free(
      Control$dMonad$dFree.$FreeView("Return", visitor.onType(a)),
      Data$dCatList.$CatList("CatCons", traverseType2(visitor$p), Data$dCatQueue.$CatQueue(Data$dList$dTypes.Nil, Data$dList$dTypes.Nil))
    ),
    onDecl: a => Control$dMonad$dFree.$Free(
      Control$dMonad$dFree.$FreeView("Return", visitor.onDecl(a)),
      Data$dCatList.$CatList("CatCons", traverseDecl2(visitor$p), Data$dCatQueue.$CatQueue(Data$dList$dTypes.Nil, Data$dList$dTypes.Nil))
    )
  };
  return visitor$p;
};
const rewriteWithContextM = dictMonad => {
  const topDownTraversalWithContextM1 = topDownTraversalWithContextM(dictMonad);
  return traversal => visitor => ctx => g => dictMonad.Bind1().Apply0().Functor0().map(Data$dTuple.Tuple(ctx))(traversal(topDownTraversalWithContextM1(visitor))(g)(ctx));
};
const rewriteWithContext = traversal => visitor => ctx => g => Data$dTuple.$Tuple(ctx, traversal(topDownTraversalWithContext(visitor))(g)(ctx));
const rewriteTypeWithContextM = dictMonad => rewriteWithContextM(dictMonad)(v => v.onType);
const rewriteTypeWithContext = visitor => ctx => g => Data$dTuple.$Tuple(ctx, topDownTraversalWithContext(visitor).onType(g)(ctx));
const rewriteTopDownM = dictMonad => {
  const topDownTraversal1 = topDownTraversal(dictMonad);
  return traversal => visitor => traversal(topDownTraversal1(visitor));
};
const rewriteTypeTopDownM = dictMonad => {
  const topDownTraversal1 = topDownTraversal(dictMonad);
  return visitor => topDownTraversal1(visitor).onType;
};
const rewriteTopDown = traversal => visitor => {
  const $0 = Control$dMonad$dFree.runFree(Data$dIdentity.functorIdentity)(Unsafe$dCoerce.unsafeCoerce);
  const $1 = traversal(topDownPureTraversal(visitor));
  return x => $0($1(x));
};
const rewriteTypeTopDown = /* #__PURE__ */ rewriteTopDown(v => v.onType);
const rewriteModuleWithContextM = dictMonad => rewriteWithContextM(dictMonad)(traverseModule((() => {
  const $0 = dictMonad.Applicative0();
  const $1 = $0.Apply0();
  const $2 = $1.Functor0();
  const functorReaderT1 = {
    map: x => {
      const $3 = $2.map(x);
      return v => x$1 => $3(v(x$1));
    }
  };
  const applyReaderT1 = {apply: v => v1 => r => $1.apply(v(r))(v1(r)), Functor0: () => functorReaderT1};
  return {
    pure: x => {
      const $3 = $0.pure(x);
      return v => $3;
    },
    Apply0: () => applyReaderT1
  };
})()));
const rewriteModuleWithContext = /* #__PURE__ */ rewriteWithContext(/* #__PURE__ */ traverseModule(applicativeReaderT));
const rewriteModuleTopDownM = dictMonad => {
  const topDownTraversal1 = topDownTraversal(dictMonad);
  const $0 = traverseModule(dictMonad.Applicative0());
  return visitor => $0(topDownTraversal1(visitor));
};
const rewriteModuleTopDown = /* #__PURE__ */ rewriteTopDown(traverseModule1);
const rewriteExprWithContextM = dictMonad => rewriteWithContextM(dictMonad)(v => v.onExpr);
const rewriteExprWithContext = visitor => ctx => g => Data$dTuple.$Tuple(ctx, topDownTraversalWithContext(visitor).onExpr(g)(ctx));
const rewriteExprTopDownM = dictMonad => {
  const topDownTraversal1 = topDownTraversal(dictMonad);
  return visitor => topDownTraversal1(visitor).onExpr;
};
const rewriteExprTopDown = /* #__PURE__ */ rewriteTopDown(v => v.onExpr);
const rewriteDeclWithContextM = dictMonad => rewriteWithContextM(dictMonad)(v => v.onDecl);
const rewriteDeclWithContext = visitor => ctx => g => Data$dTuple.$Tuple(ctx, topDownTraversalWithContext(visitor).onDecl(g)(ctx));
const rewriteDeclTopDownM = dictMonad => {
  const topDownTraversal1 = topDownTraversal(dictMonad);
  return visitor => topDownTraversal1(visitor).onDecl;
};
const rewriteDeclTopDown = /* #__PURE__ */ rewriteTopDown(v => v.onDecl);
const rewriteBinderWithContextM = dictMonad => rewriteWithContextM(dictMonad)(v => v.onBinder);
const rewriteBinderWithContext = visitor => ctx => g => Data$dTuple.$Tuple(ctx, topDownTraversalWithContext(visitor).onBinder(g)(ctx));
const rewriteBinderTopDownM = dictMonad => {
  const topDownTraversal1 = topDownTraversal(dictMonad);
  return visitor => topDownTraversal1(visitor).onBinder;
};
const rewriteBinderTopDown = /* #__PURE__ */ rewriteTopDown(v => v.onBinder);
const topDownMonoidalTraversal = dictMonoid => {
  const $0 = applyCompose((() => {
    const $0 = dictMonoid.Semigroup0();
    return {apply: v => v1 => $0.append(v)(v1), Functor0: () => Data$dConst.functorConst};
  })());
  const applicativeCompose1 = applicativeCompose(Data$dConst.applicativeConst(dictMonoid));
  const traverseBinder3 = traverseBinder(applicativeCompose1);
  const traverseExpr3 = traverseExpr(applicativeCompose1);
  const traverseDecl3 = traverseDecl(applicativeCompose1);
  const traverseType3 = traverseType(applicativeCompose1);
  return visitor => {
    const visitor$p = {
      onBinder: a => $0.apply(Control$dMonad$dFree.$Free(Control$dMonad$dFree.$FreeView("Return", visitor.onBinder(a)), Data$dCatList.CatNil))(Control$dMonad$dFree.$Free(
        Control$dMonad$dFree.$FreeView("Return", undefined),
        Data$dCatList.$CatList("CatCons", v => traverseBinder3(visitor$p)(a), Data$dCatQueue.$CatQueue(Data$dList$dTypes.Nil, Data$dList$dTypes.Nil))
      )),
      onExpr: a => $0.apply(Control$dMonad$dFree.$Free(Control$dMonad$dFree.$FreeView("Return", visitor.onExpr(a)), Data$dCatList.CatNil))(Control$dMonad$dFree.$Free(
        Control$dMonad$dFree.$FreeView("Return", undefined),
        Data$dCatList.$CatList("CatCons", v => traverseExpr3(visitor$p)(a), Data$dCatQueue.$CatQueue(Data$dList$dTypes.Nil, Data$dList$dTypes.Nil))
      )),
      onDecl: a => $0.apply(Control$dMonad$dFree.$Free(Control$dMonad$dFree.$FreeView("Return", visitor.onDecl(a)), Data$dCatList.CatNil))(Control$dMonad$dFree.$Free(
        Control$dMonad$dFree.$FreeView("Return", undefined),
        Data$dCatList.$CatList("CatCons", v => traverseDecl3(visitor$p)(a), Data$dCatQueue.$CatQueue(Data$dList$dTypes.Nil, Data$dList$dTypes.Nil))
      )),
      onType: a => $0.apply(Control$dMonad$dFree.$Free(Control$dMonad$dFree.$FreeView("Return", visitor.onType(a)), Data$dCatList.CatNil))(Control$dMonad$dFree.$Free(
        Control$dMonad$dFree.$FreeView("Return", undefined),
        Data$dCatList.$CatList("CatCons", v => traverseType3(visitor$p)(a), Data$dCatQueue.$CatQueue(Data$dList$dTypes.Nil, Data$dList$dTypes.Nil))
      ))
    };
    return visitor$p;
  };
};
const monoidalRewrite = dictMonoid => {
  const topDownMonoidalTraversal1 = topDownMonoidalTraversal(dictMonoid);
  return traversal => visitor => g => Control$dMonad$dFree.runFree(Data$dIdentity.functorIdentity)(Unsafe$dCoerce.unsafeCoerce)(traversal(topDownMonoidalTraversal1(visitor))(g));
};
const foldMapBinder = dictMonoid => monoidalRewrite(dictMonoid)(v => v.onBinder);
const foldMapDecl = dictMonoid => monoidalRewrite(dictMonoid)(v => v.onDecl);
const foldMapExpr = dictMonoid => monoidalRewrite(dictMonoid)(v => v.onExpr);
const foldMapModule = dictMonoid => monoidalRewrite(dictMonoid)(traverseModule(applicativeCompose(Data$dConst.applicativeConst(dictMonoid))));
const foldMapType = dictMonoid => monoidalRewrite(dictMonoid)(v => v.onType);
const defaultVisitorWithContextM = dictMonad => {
  const $0 = dictMonad.Applicative0();
  return {
    onBinder: a => b => $0.pure(Data$dTuple.$Tuple(a, b)),
    onDecl: a => b => $0.pure(Data$dTuple.$Tuple(a, b)),
    onExpr: a => b => $0.pure(Data$dTuple.$Tuple(a, b)),
    onType: a => b => $0.pure(Data$dTuple.$Tuple(a, b))
  };
};
const defaultVisitorWithContext = {
  onBinder: a => b => Data$dTuple.$Tuple(a, b),
  onDecl: a => b => Data$dTuple.$Tuple(a, b),
  onExpr: a => b => Data$dTuple.$Tuple(a, b),
  onType: a => b => Data$dTuple.$Tuple(a, b)
};
const defaultVisitorM = dictApplicative => ({onBinder: dictApplicative.pure, onDecl: dictApplicative.pure, onExpr: dictApplicative.pure, onType: dictApplicative.pure});
const defaultVisitor = {onBinder: identity, onDecl: identity, onExpr: identity, onType: identity};
const defaultMonoidalVisitor = dictMonoid => {
  const mempty1 = dictMonoid.mempty;
  return {onBinder: v => mempty1, onDecl: v => mempty1, onExpr: v => mempty1, onType: v => mempty1};
};
const bottomUpTraversal = dictMonad => {
  const $0 = dictMonad.Bind1();
  const Applicative0 = dictMonad.Applicative0();
  const traverseBinder3 = traverseBinder(Applicative0);
  const traverseExpr3 = traverseExpr(Applicative0);
  const traverseType3 = traverseType(Applicative0);
  const traverseDecl3 = traverseDecl(Applicative0);
  return visitor => {
    const visitor$p = {
      onBinder: a => $0.bind(dictMonad.Bind1().bind(dictMonad.Applicative0().pure())(v => traverseBinder3(visitor$p)(a)))(visitor.onBinder),
      onExpr: a => $0.bind(dictMonad.Bind1().bind(dictMonad.Applicative0().pure())(v => traverseExpr3(visitor$p)(a)))(visitor.onExpr),
      onType: a => $0.bind(dictMonad.Bind1().bind(dictMonad.Applicative0().pure())(v => traverseType3(visitor$p)(a)))(visitor.onType),
      onDecl: a => $0.bind(dictMonad.Bind1().bind(dictMonad.Applicative0().pure())(v => traverseDecl3(visitor$p)(a)))(visitor.onDecl)
    };
    return visitor$p;
  };
};
const rewriteBottomUpM = dictMonad => {
  const bottomUpTraversal1 = bottomUpTraversal(dictMonad);
  return traversal => visitor => traversal(bottomUpTraversal1(visitor));
};
const rewriteBinderBottomUpM = dictMonad => {
  const bottomUpTraversal1 = bottomUpTraversal(dictMonad);
  return visitor => bottomUpTraversal1(visitor).onBinder;
};
const rewriteDeclBottomUpM = dictMonad => {
  const bottomUpTraversal1 = bottomUpTraversal(dictMonad);
  return visitor => bottomUpTraversal1(visitor).onDecl;
};
const rewriteExprBottomUpM = dictMonad => {
  const bottomUpTraversal1 = bottomUpTraversal(dictMonad);
  return visitor => bottomUpTraversal1(visitor).onExpr;
};
const rewriteModuleBottomUpM = dictMonad => {
  const bottomUpTraversal1 = bottomUpTraversal(dictMonad);
  const $0 = traverseModule(dictMonad.Applicative0());
  return visitor => $0(bottomUpTraversal1(visitor));
};
const rewriteTypeBottomUpM = dictMonad => {
  const bottomUpTraversal1 = bottomUpTraversal(dictMonad);
  return visitor => bottomUpTraversal1(visitor).onType;
};
const bottomUpPureTraversal = visitor => {
  const visitor$p = {
    onBinder: a => {
      const $0 = traverseBinder2(visitor$p)(a);
      return Control$dMonad$dFree.$Free(
        $0._1,
        (() => {
          if ($0._2.tag === "CatNil") {
            return Data$dCatList.$CatList(
              "CatCons",
              x => Control$dMonad$dFree.$Free(Control$dMonad$dFree.$FreeView("Return", visitor.onBinder(x)), Data$dCatList.CatNil),
              Data$dCatQueue.$CatQueue(Data$dList$dTypes.Nil, Data$dList$dTypes.Nil)
            );
          }
          if ($0._2.tag === "CatCons") {
            return Data$dCatList.$CatList(
              "CatCons",
              $0._2._1,
              Data$dCatQueue.$CatQueue(
                $0._2._2._1,
                Data$dList$dTypes.$List(
                  "Cons",
                  Data$dCatList.$CatList(
                    "CatCons",
                    x => Control$dMonad$dFree.$Free(Control$dMonad$dFree.$FreeView("Return", visitor.onBinder(x)), Data$dCatList.CatNil),
                    Data$dCatQueue.$CatQueue(Data$dList$dTypes.Nil, Data$dList$dTypes.Nil)
                  ),
                  $0._2._2._2
                )
              )
            );
          }
          $runtime.fail();
        })()
      );
    },
    onExpr: a => {
      const $0 = traverseExpr2(visitor$p)(a);
      return Control$dMonad$dFree.$Free(
        $0._1,
        (() => {
          if ($0._2.tag === "CatNil") {
            return Data$dCatList.$CatList(
              "CatCons",
              x => Control$dMonad$dFree.$Free(Control$dMonad$dFree.$FreeView("Return", visitor.onExpr(x)), Data$dCatList.CatNil),
              Data$dCatQueue.$CatQueue(Data$dList$dTypes.Nil, Data$dList$dTypes.Nil)
            );
          }
          if ($0._2.tag === "CatCons") {
            return Data$dCatList.$CatList(
              "CatCons",
              $0._2._1,
              Data$dCatQueue.$CatQueue(
                $0._2._2._1,
                Data$dList$dTypes.$List(
                  "Cons",
                  Data$dCatList.$CatList(
                    "CatCons",
                    x => Control$dMonad$dFree.$Free(Control$dMonad$dFree.$FreeView("Return", visitor.onExpr(x)), Data$dCatList.CatNil),
                    Data$dCatQueue.$CatQueue(Data$dList$dTypes.Nil, Data$dList$dTypes.Nil)
                  ),
                  $0._2._2._2
                )
              )
            );
          }
          $runtime.fail();
        })()
      );
    },
    onType: a => {
      const $0 = traverseType2(visitor$p)(a);
      return Control$dMonad$dFree.$Free(
        $0._1,
        (() => {
          if ($0._2.tag === "CatNil") {
            return Data$dCatList.$CatList(
              "CatCons",
              x => Control$dMonad$dFree.$Free(Control$dMonad$dFree.$FreeView("Return", visitor.onType(x)), Data$dCatList.CatNil),
              Data$dCatQueue.$CatQueue(Data$dList$dTypes.Nil, Data$dList$dTypes.Nil)
            );
          }
          if ($0._2.tag === "CatCons") {
            return Data$dCatList.$CatList(
              "CatCons",
              $0._2._1,
              Data$dCatQueue.$CatQueue(
                $0._2._2._1,
                Data$dList$dTypes.$List(
                  "Cons",
                  Data$dCatList.$CatList(
                    "CatCons",
                    x => Control$dMonad$dFree.$Free(Control$dMonad$dFree.$FreeView("Return", visitor.onType(x)), Data$dCatList.CatNil),
                    Data$dCatQueue.$CatQueue(Data$dList$dTypes.Nil, Data$dList$dTypes.Nil)
                  ),
                  $0._2._2._2
                )
              )
            );
          }
          $runtime.fail();
        })()
      );
    },
    onDecl: a => {
      const $0 = traverseDecl2(visitor$p)(a);
      return Control$dMonad$dFree.$Free(
        $0._1,
        (() => {
          if ($0._2.tag === "CatNil") {
            return Data$dCatList.$CatList(
              "CatCons",
              x => Control$dMonad$dFree.$Free(Control$dMonad$dFree.$FreeView("Return", visitor.onDecl(x)), Data$dCatList.CatNil),
              Data$dCatQueue.$CatQueue(Data$dList$dTypes.Nil, Data$dList$dTypes.Nil)
            );
          }
          if ($0._2.tag === "CatCons") {
            return Data$dCatList.$CatList(
              "CatCons",
              $0._2._1,
              Data$dCatQueue.$CatQueue(
                $0._2._2._1,
                Data$dList$dTypes.$List(
                  "Cons",
                  Data$dCatList.$CatList(
                    "CatCons",
                    x => Control$dMonad$dFree.$Free(Control$dMonad$dFree.$FreeView("Return", visitor.onDecl(x)), Data$dCatList.CatNil),
                    Data$dCatQueue.$CatQueue(Data$dList$dTypes.Nil, Data$dList$dTypes.Nil)
                  ),
                  $0._2._2._2
                )
              )
            );
          }
          $runtime.fail();
        })()
      );
    }
  };
  return visitor$p;
};
const rewriteBottomUp = traversal => visitor => {
  const $0 = Control$dMonad$dFree.runFree(Data$dIdentity.functorIdentity)(Unsafe$dCoerce.unsafeCoerce);
  const $1 = traversal(bottomUpPureTraversal(visitor));
  return x => $0($1(x));
};
const rewriteBinderBottomUp = /* #__PURE__ */ rewriteBottomUp(v => v.onBinder);
const rewriteDeclBottomUp = /* #__PURE__ */ rewriteBottomUp(v => v.onDecl);
const rewriteExprBottomUp = /* #__PURE__ */ rewriteBottomUp(v => v.onExpr);
const rewriteModuleBottomUp = /* #__PURE__ */ rewriteBottomUp(traverseModule1);
const rewriteTypeBottomUp = /* #__PURE__ */ rewriteBottomUp(v => v.onType);
export {
  applicativeCompose,
  applicativeReaderT,
  applyCompose,
  bottomUpPureTraversal,
  bottomUpTraversal,
  defaultMonoidalVisitor,
  defaultVisitor,
  defaultVisitorM,
  defaultVisitorWithContext,
  defaultVisitorWithContextM,
  foldMapBinder,
  foldMapDecl,
  foldMapExpr,
  foldMapModule,
  foldMapType,
  identity,
  ltraverse,
  monoidalRewrite,
  rewriteBinderBottomUp,
  rewriteBinderBottomUpM,
  rewriteBinderTopDown,
  rewriteBinderTopDownM,
  rewriteBinderWithContext,
  rewriteBinderWithContextM,
  rewriteBottomUp,
  rewriteBottomUpM,
  rewriteDeclBottomUp,
  rewriteDeclBottomUpM,
  rewriteDeclTopDown,
  rewriteDeclTopDownM,
  rewriteDeclWithContext,
  rewriteDeclWithContextM,
  rewriteExprBottomUp,
  rewriteExprBottomUpM,
  rewriteExprTopDown,
  rewriteExprTopDownM,
  rewriteExprWithContext,
  rewriteExprWithContextM,
  rewriteModuleBottomUp,
  rewriteModuleBottomUpM,
  rewriteModuleTopDown,
  rewriteModuleTopDownM,
  rewriteModuleWithContext,
  rewriteModuleWithContextM,
  rewriteTopDown,
  rewriteTopDownM,
  rewriteTypeBottomUp,
  rewriteTypeBottomUpM,
  rewriteTypeTopDown,
  rewriteTypeTopDownM,
  rewriteTypeWithContext,
  rewriteTypeWithContextM,
  rewriteWithContext,
  rewriteWithContextM,
  topDownMonoidalTraversal,
  topDownPureTraversal,
  topDownTraversal,
  topDownTraversalWithContext,
  topDownTraversalWithContextM,
  traverseAdoBlock,
  traverseBinder,
  traverseBinder1,
  traverseBinder2,
  traverseCaseOf,
  traverseClassHead,
  traverseDataCtor,
  traverseDataHead,
  traverseDecl,
  traverseDecl1,
  traverseDecl2,
  traverseDelimited,
  traverseDelimitedNonEmpty,
  traverseDoBlock,
  traverseDoStatement,
  traverseExpr,
  traverseExpr1,
  traverseExpr2,
  traverseExprAppSpine,
  traverseForeign,
  traverseGuarded,
  traverseGuardedExpr,
  traverseIfThenElse,
  traverseInstance,
  traverseInstanceBinding,
  traverseInstanceHead,
  traverseLabeled,
  traverseLambda,
  traverseLetBinding,
  traverseLetIn,
  traverseModule,
  traverseModule1,
  traverseModuleBody,
  traverseOneOrDelimited,
  traversePatternGuard,
  traverseRecordAccessor,
  traverseRecordLabeled,
  traverseRecordUpdate,
  traverseRow,
  traverseSeparated,
  traverseType,
  traverseType1,
  traverseType2,
  traverseTypeVarBinding,
  traverseValueBindingFields,
  traverseWhere,
  traverseWrapped
};
