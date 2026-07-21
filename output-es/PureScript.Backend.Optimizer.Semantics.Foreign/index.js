import * as $runtime from "../runtime.js";
import * as Data$dArray from "../Data.Array/index.js";
import * as Data$dArray$dNonEmpty from "../Data.Array.NonEmpty/index.js";
import * as Data$dEither from "../Data.Either/index.js";
import * as Data$dFoldable from "../Data.Foldable/index.js";
import * as Data$dFunctor from "../Data.Functor/index.js";
import * as Data$dList$dTypes from "../Data.List.Types/index.js";
import * as Data$dMap$dInternal from "../Data.Map.Internal/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dOrd from "../Data.Ord/index.js";
import * as Data$dShow from "../Data.Show/index.js";
import * as Data$dString$dCodePoints from "../Data.String.CodePoints/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
import * as PureScript$dBackend$dOptimizer$dCoreFn from "../PureScript.Backend.Optimizer.CoreFn/index.js";
import * as PureScript$dBackend$dOptimizer$dSemantics from "../PureScript.Backend.Optimizer.Semantics/index.js";
import * as PureScript$dBackend$dOptimizer$dSyntax from "../PureScript.Backend.Optimizer.Syntax/index.js";
const fromFoldable = /* #__PURE__ */ Data$dFoldable.foldrArray(Data$dList$dTypes.Cons)(Data$dList$dTypes.Nil);
const viewCopyRecord = v => {
  if (
    v.tag === "SemRef" && v._1.tag === "EvalExtern" && v._1._1._1.tag === "Just" && v._1._1._1._1 === "Record.Builder" && v._1._1._2 === "copyRecord" && v._2.length === 1 && v._2[0].tag === "ExternApp" && v._2[0]._1.length === 1
  ) {
    return Data$dMaybe.$Maybe("Just", v._2[0]._1[0]);
  }
  return Data$dMaybe.Nothing;
};
const qualified = mod => id => PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", mod), id);
const record_builder_copyRecord = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Record.Builder"), "copyRecord"),
  v => v1 => v2 => {
    if (v2.length === 1 && v2[0].tag === "ExternApp" && v2[0]._1.length === 1) {
      if (v2[0]._1[0].tag === "NeutLit") {
        if (v2[0]._1[0]._1.tag === "LitRecord") { return Data$dMaybe.$Maybe("Just", v2[0]._1[0]); }
        return Data$dMaybe.Nothing;
      }
      if (v2[0]._1[0].tag === "NeutUpdate") { return Data$dMaybe.$Maybe("Just", v2[0]._1[0]); }
    }
    return Data$dMaybe.Nothing;
  }
);
const record_builder_unsafeDelete = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Record.Builder"), "unsafeDelete"),
  v => v1 => v2 => {
    if (
      v2.length === 1 && v2[0].tag === "ExternApp" && v2[0]._1.length === 2 && v2[0]._1[0].tag === "NeutLit" && v2[0]._1[0]._1.tag === "LitString" && v2[0]._1[1].tag === "NeutLit" && v2[0]._1[1]._1.tag === "LitRecord"
    ) {
      const $0 = v2[0]._1[0]._1._1;
      return Data$dMaybe.$Maybe(
        "Just",
        PureScript$dBackend$dOptimizer$dSemantics.$BackendSemantics(
          "NeutLit",
          PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitRecord", Data$dArray.filterImpl(x => $0 !== x._1, v2[0]._1[1]._1._1))
        )
      );
    }
    return Data$dMaybe.Nothing;
  }
);
const record_builder_unsafeInsert = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Record.Builder"), "unsafeInsert"),
  v => v1 => v2 => {
    const $0 = () => {
      if (v2.length === 1 && v2[0].tag === "ExternApp" && v2[0]._1.length === 3 && v2[0]._1[0].tag === "NeutLit" && v2[0]._1[0]._1.tag === "LitString") {
        const $0 = v2[0]._1[0]._1._1;
        if (
          v2[0]._1[2].tag === "SemRef" && v2[0]._1[2]._1.tag === "EvalExtern" && v2[0]._1[2]._1._1._1.tag === "Just" && v2[0]._1[2]._1._1._1._1 === "Record.Builder" && v2[0]._1[2]._1._1._2 === "copyRecord" && v2[0]._1[2]._2.length === 1 && v2[0]._1[2]._2[0].tag === "ExternApp" && v2[0]._1[2]._2[0]._1.length === 1
        ) {
          return Data$dMaybe.$Maybe(
            "Just",
            PureScript$dBackend$dOptimizer$dSemantics.evalUpdate(v2[0]._1[2]._2[0]._1[0])([PureScript$dBackend$dOptimizer$dCoreFn.$Prop($0, v2[0]._1[1])])
          );
        }
      }
      return Data$dMaybe.Nothing;
    };
    if (v2.length === 1 && v2[0].tag === "ExternApp" && v2[0]._1.length === 3 && v2[0]._1[0].tag === "NeutLit" && v2[0]._1[0]._1.tag === "LitString") {
      if (v2[0]._1[2].tag === "NeutLit") {
        if (v2[0]._1[2]._1.tag === "LitRecord") {
          return Data$dMaybe.$Maybe(
            "Just",
            PureScript$dBackend$dOptimizer$dSemantics.$BackendSemantics(
              "NeutLit",
              PureScript$dBackend$dOptimizer$dCoreFn.$Literal(
                "LitRecord",
                Data$dArray.snoc(v2[0]._1[2]._1._1)(PureScript$dBackend$dOptimizer$dCoreFn.$Prop(v2[0]._1[0]._1._1, v2[0]._1[1]))
              )
            )
          );
        }
        return $0();
      }
      if (v2[0]._1[2].tag === "NeutUpdate") {
        return Data$dMaybe.$Maybe(
          "Just",
          PureScript$dBackend$dOptimizer$dSemantics.evalUpdate(v2[0]._1[2])([PureScript$dBackend$dOptimizer$dCoreFn.$Prop(v2[0]._1[0]._1._1, v2[0]._1[1])])
        );
      }
    }
    return $0();
  }
);
const record_builder_unsafeModify = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Record.Builder"), "unsafeModify"),
  env => v => v1 => {
    const $0 = () => {
      if (v1.length === 1 && v1[0].tag === "ExternApp" && v1[0]._1.length === 3 && v1[0]._1[0].tag === "NeutLit" && v1[0]._1[0]._1.tag === "LitString") {
        const $0 = v1[0]._1[0]._1._1;
        if (
          v1[0]._1[2].tag === "SemRef" && v1[0]._1[2]._1.tag === "EvalExtern" && v1[0]._1[2]._1._1._1.tag === "Just" && v1[0]._1[2]._1._1._1._1 === "Record.Builder" && v1[0]._1[2]._1._1._2 === "copyRecord" && v1[0]._1[2]._2.length === 1 && v1[0]._1[2]._2[0].tag === "ExternApp" && v1[0]._1[2]._2[0]._1.length === 1
        ) {
          const $1 = v1[0]._1[2]._2[0]._1[0];
          return Data$dMaybe.$Maybe(
            "Just",
            PureScript$dBackend$dOptimizer$dSemantics.makeLet(Data$dMaybe.Nothing)($1)(r$p => PureScript$dBackend$dOptimizer$dSemantics.evalUpdate($1)([
              PureScript$dBackend$dOptimizer$dCoreFn.$Prop(
                $0,
                PureScript$dBackend$dOptimizer$dSemantics.evalApp(env)(v1[0]._1[1])([
                  PureScript$dBackend$dOptimizer$dSemantics.evalAccessor(env)(r$p)(PureScript$dBackend$dOptimizer$dSyntax.$BackendAccessor("GetProp", $0))
                ])
              )
            ]))
          );
        }
      }
      return Data$dMaybe.Nothing;
    };
    if (v1.length === 1 && v1[0].tag === "ExternApp" && v1[0]._1.length === 3 && v1[0]._1[0].tag === "NeutLit" && v1[0]._1[0]._1.tag === "LitString") {
      if (v1[0]._1[2].tag === "NeutLit") {
        if (v1[0]._1[2]._1.tag === "LitRecord") {
          const $1 = v1[0]._1[1];
          const $2 = v1[0]._1[0]._1._1;
          return Data$dMaybe.$Maybe(
            "Just",
            PureScript$dBackend$dOptimizer$dSemantics.$BackendSemantics(
              "NeutLit",
              PureScript$dBackend$dOptimizer$dCoreFn.$Literal(
                "LitRecord",
                Data$dFunctor.arrayMap(v2 => {
                  if ($2 === v2._1) { return PureScript$dBackend$dOptimizer$dCoreFn.$Prop(v2._1, PureScript$dBackend$dOptimizer$dSemantics.evalApp(env)($1)([v2._2])); }
                  return PureScript$dBackend$dOptimizer$dCoreFn.$Prop(v2._1, v2._2);
                })(v1[0]._1[2]._1._1)
              )
            )
          );
        }
        return $0();
      }
      if (v1[0]._1[2].tag === "NeutUpdate" && v1[0]._1[2]._1.tag === "NeutLocal") {
        const $1 = v1[0]._1[0]._1._1;
        return Data$dMaybe.$Maybe(
          "Just",
          PureScript$dBackend$dOptimizer$dSemantics.evalUpdate(v1[0]._1[2])([
            PureScript$dBackend$dOptimizer$dCoreFn.$Prop(
              $1,
              PureScript$dBackend$dOptimizer$dSemantics.evalApp(env)(v1[0]._1[1])([
                PureScript$dBackend$dOptimizer$dSemantics.evalAccessor(env)(v1[0]._1[2]._1)(PureScript$dBackend$dOptimizer$dSyntax.$BackendAccessor("GetProp", $1))
              ])
            )
          ])
        );
      }
    }
    return $0();
  }
);
const record_builder_unsafeRename = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Record.Builder"), "unsafeRename"),
  v => v1 => v2 => {
    if (
      v2.length === 1 && v2[0].tag === "ExternApp" && v2[0]._1.length === 3 && v2[0]._1[0].tag === "NeutLit" && v2[0]._1[0]._1.tag === "LitString" && v2[0]._1[1].tag === "NeutLit" && v2[0]._1[1]._1.tag === "LitString" && v2[0]._1[2].tag === "NeutLit" && v2[0]._1[2]._1.tag === "LitRecord"
    ) {
      const $0 = v2[0]._1[0]._1._1;
      const $1 = v2[0]._1[1]._1._1;
      return Data$dMaybe.$Maybe(
        "Just",
        PureScript$dBackend$dOptimizer$dSemantics.$BackendSemantics(
          "NeutLit",
          PureScript$dBackend$dOptimizer$dCoreFn.$Literal(
            "LitRecord",
            Data$dFunctor.arrayMap(v3 => {
              if ($0 === v3._1) { return PureScript$dBackend$dOptimizer$dCoreFn.$Prop($1, v3._2); }
              return PureScript$dBackend$dOptimizer$dCoreFn.$Prop(v3._1, v3._2);
            })(v2[0]._1[2]._1._1)
          )
        )
      );
    }
    return Data$dMaybe.Nothing;
  }
);
const record_unsafe_union_unsafeUnionFn = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Record.Unsafe.Union"), "unsafeUnionFn"),
  v => v1 => v2 => {
    if (
      v2.length === 1 && v2[0].tag === "ExternUncurriedApp" && v2[0]._1.length === 2 && v2[0]._1[0].tag === "NeutLit" && v2[0]._1[0]._1.tag === "LitRecord" && v2[0]._1[1].tag === "NeutLit" && v2[0]._1[1]._1.tag === "LitRecord"
    ) {
      return Data$dMaybe.$Maybe(
        "Just",
        PureScript$dBackend$dOptimizer$dSemantics.$BackendSemantics(
          "NeutLit",
          PureScript$dBackend$dOptimizer$dCoreFn.$Literal(
            "LitRecord",
            Data$dFunctor.arrayMap(Data$dArray$dNonEmpty.head)(Data$dArray.groupAllBy(x => y => Data$dOrd.ordString.compare(x._1)(y._1))([
              ...v2[0]._1[0]._1._1,
              ...v2[0]._1[1]._1._1
            ]))
          )
        )
      );
    }
    return Data$dMaybe.Nothing;
  }
);
const record_unsafe_unsafeDelete = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Record.Unsafe"), "unsafeDelete"),
  v => v1 => v2 => {
    if (
      v2.length === 1 && v2[0].tag === "ExternApp" && v2[0]._1.length === 2 && v2[0]._1[0].tag === "NeutLit" && v2[0]._1[0]._1.tag === "LitString" && v2[0]._1[1].tag === "NeutLit" && v2[0]._1[1]._1.tag === "LitRecord"
    ) {
      const $0 = v2[0]._1[0]._1._1;
      return Data$dMaybe.$Maybe(
        "Just",
        PureScript$dBackend$dOptimizer$dSemantics.$BackendSemantics(
          "NeutLit",
          PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitRecord", Data$dArray.filterImpl(x => $0 !== x._1, v2[0]._1[1]._1._1))
        )
      );
    }
    return Data$dMaybe.Nothing;
  }
);
const record_unsafe_unsafeGet = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Record.Unsafe"), "unsafeGet"),
  env => v => v1 => {
    if (v1.length === 1 && v1[0].tag === "ExternApp" && v1[0]._1.length === 1 && v1[0]._1[0].tag === "NeutLit" && v1[0]._1[0]._1.tag === "LitString") {
      const $0 = v1[0]._1[0]._1._1;
      return Data$dMaybe.$Maybe(
        "Just",
        PureScript$dBackend$dOptimizer$dSemantics.$BackendSemantics(
          "SemLam",
          Data$dMaybe.Nothing,
          r => PureScript$dBackend$dOptimizer$dSemantics.evalAccessor(env)(r)(PureScript$dBackend$dOptimizer$dSyntax.$BackendAccessor("GetProp", $0))
        )
      );
    }
    return Data$dMaybe.Nothing;
  }
);
const record_unsafe_unsafeHas = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Record.Unsafe"), "unsafeHas"),
  v => v1 => v2 => {
    if (
      v2.length === 1 && v2[0].tag === "ExternApp" && v2[0]._1.length === 2 && v2[0]._1[0].tag === "NeutLit" && v2[0]._1[0]._1.tag === "LitString" && v2[0]._1[1].tag === "NeutLit" && v2[0]._1[1]._1.tag === "LitRecord"
    ) {
      const $0 = v2[0]._1[0]._1._1;
      return Data$dMaybe.$Maybe(
        "Just",
        PureScript$dBackend$dOptimizer$dSemantics.$BackendSemantics(
          "NeutLit",
          PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitBoolean", Data$dArray.anyImpl(x => $0 === x._1, v2[0]._1[1]._1._1))
        )
      );
    }
    return Data$dMaybe.Nothing;
  }
);
const record_unsafe_unsafeSet = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Record.Unsafe"), "unsafeSet"),
  v => v1 => v2 => {
    if (v2.length === 1 && v2[0].tag === "ExternApp" && v2[0]._1.length === 3 && v2[0]._1[0].tag === "NeutLit" && v2[0]._1[0]._1.tag === "LitString") {
      return Data$dMaybe.$Maybe(
        "Just",
        PureScript$dBackend$dOptimizer$dSemantics.evalUpdate(v2[0]._1[2])([PureScript$dBackend$dOptimizer$dCoreFn.$Prop(v2[0]._1[0]._1._1, v2[0]._1[1])])
      );
    }
    return Data$dMaybe.Nothing;
  }
);
const runEffectFn = mod => name => n => {
  const goRunEffectFn = env => acc => head => v => {
    if (v.tag === "Nil") { return PureScript$dBackend$dOptimizer$dSemantics.evalUncurriedEffectApp(env)(head)(acc); }
    if (v.tag === "Cons") {
      const $0 = v._2;
      return PureScript$dBackend$dOptimizer$dSemantics.makeLet(Data$dMaybe.Nothing)(v._1)(nextArg => goRunEffectFn(env)(Data$dArray.snoc(acc)(nextArg))(head)($0));
    }
    $runtime.fail();
  };
  return Data$dTuple.$Tuple(
    PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", mod), name + Data$dShow.showIntImpl(n)),
    env => v => v1 => {
      if (v1.length === 1 && v1[0].tag === "ExternApp") {
        const $0 = Data$dArray.unconsImpl(v$1 => Data$dMaybe.Nothing, x => xs => Data$dMaybe.$Maybe("Just", {head: x, tail: xs}), v1[0]._1);
        if ($0.tag === "Just" && $0._1.tail.length === n) { return Data$dMaybe.$Maybe("Just", goRunEffectFn(env)([])($0._1.head)(fromFoldable($0._1.tail))); }
      }
      return Data$dMaybe.Nothing;
    }
  );
};
const unsafe_coerce_unsafeCoerce = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Unsafe.Coerce"), "unsafeCoerce"),
  v => v1 => v2 => {
    if (v2.length === 1 && v2[0].tag === "ExternApp" && v2[0]._1.length === 1) { return Data$dMaybe.$Maybe("Just", v2[0]._1[0]); }
    return Data$dMaybe.Nothing;
  }
);
const primUnaryOperator = op => env => v => v1 => {
  if (v1.length === 1 && v1[0].tag === "ExternApp" && v1[0]._1.length === 1) {
    return Data$dMaybe.$Maybe("Just", PureScript$dBackend$dOptimizer$dSemantics.evalPrimOp(env)(PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op1", op, v1[0]._1[0])));
  }
  return Data$dMaybe.Nothing;
};
const primBinaryOperator = op => env => v => v1 => {
  if (v1.length === 1 && v1[0].tag === "ExternApp" && v1[0]._1.length === 1) {
    return Data$dMaybe.$Maybe(
      "Just",
      PureScript$dBackend$dOptimizer$dSemantics.makeLet(Data$dMaybe.Nothing)(v1[0]._1[0])(a$p => PureScript$dBackend$dOptimizer$dSemantics.$BackendSemantics(
        "SemLam",
        Data$dMaybe.Nothing,
        b$p => PureScript$dBackend$dOptimizer$dSemantics.evalPrimOp(env)(PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator("Op2", op, a$p, b$p))
      ))
    );
  }
  return Data$dMaybe.Nothing;
};
const partial_unsafe_unsafePartial = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Partial.Unsafe"), "_unsafePartial"),
  v => v1 => v2 => {
    if (v2.length === 1 && v2[0].tag === "ExternApp" && v2[0]._1.length === 1 && v2[0]._1[0].tag === "SemLam") {
      return Data$dMaybe.$Maybe(
        "Just",
        v2[0]._1[0]._2(PureScript$dBackend$dOptimizer$dSemantics.$BackendSemantics("NeutLit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitRecord", [])))
      );
    }
    return Data$dMaybe.Nothing;
  }
);
const mkEffectFn = mod => name => n => Data$dTuple.$Tuple(
  PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", mod), name + Data$dShow.showIntImpl(n)),
  env => v => v1 => {
    if (v1.length === 1 && v1[0].tag === "ExternApp" && v1[0]._1.length === 1) {
      return Data$dMaybe.$Maybe(
        "Just",
        PureScript$dBackend$dOptimizer$dSemantics.$BackendSemantics("SemMkEffectFn", PureScript$dBackend$dOptimizer$dSemantics.evalMkFn(env)(n)(v1[0]._1[0]))
      );
    }
    return Data$dMaybe.Nothing;
  }
);
const isQualified = mod => tag => v => v._1.tag === "Just" && mod === v._1._1 && tag === v._2;
const primOrdOperator = op => env => v => v1 => {
  if (
    v1.length === 3 && v1[0].tag === "ExternAccessor" && v1[0]._1.tag === "GetProp" && v1[0]._1._1 === "compare" && v1[1].tag === "ExternApp" && v1[1]._1.length === 2 && v1[2].tag === "ExternPrimOp" && v1[2]._1.tag === "OpIsTag"
  ) {
    if (v1[2]._1._1._1.tag === "Just" && "Data.Ordering" === v1[2]._1._1._1._1 && "LT" === v1[2]._1._1._2) {
      return Data$dMaybe.$Maybe(
        "Just",
        PureScript$dBackend$dOptimizer$dSemantics.evalPrimOp(env)(PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator(
          "Op2",
          op(PureScript$dBackend$dOptimizer$dSyntax.OpLt),
          v1[1]._1[0],
          v1[1]._1[1]
        ))
      );
    }
    if (v1[2]._1._1._1.tag === "Just" && "Data.Ordering" === v1[2]._1._1._1._1 && "GT" === v1[2]._1._1._2) {
      return Data$dMaybe.$Maybe(
        "Just",
        PureScript$dBackend$dOptimizer$dSemantics.evalPrimOp(env)(PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator(
          "Op2",
          op(PureScript$dBackend$dOptimizer$dSyntax.OpGt),
          v1[1]._1[0],
          v1[1]._1[1]
        ))
      );
    }
    if (v1[2]._1._1._1.tag === "Just" && "Data.Ordering" === v1[2]._1._1._1._1 && "EQ" === v1[2]._1._1._2) {
      return Data$dMaybe.$Maybe(
        "Just",
        PureScript$dBackend$dOptimizer$dSemantics.evalPrimOp(env)(PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator(
          "Op2",
          op(PureScript$dBackend$dOptimizer$dSyntax.OpEq),
          v1[1]._1[0],
          v1[1]._1[1]
        ))
      );
    }
  }
  return Data$dMaybe.Nothing;
};
const effect_uncurried_runEffectFn = /* #__PURE__ */ runEffectFn("Effect.Uncurried")("runEffectFn");
const effect_uncurried_mkEffectFn = /* #__PURE__ */ mkEffectFn("Effect.Uncurried")("mkEffectFn");
const effectUnsafePerform = v => v1 => v2 => {
  if (v2.length === 1 && v2[0].tag === "ExternApp" && v2[0]._1.length === 1 && v2[0]._1[0].tag === "SemEffectPure") { return Data$dMaybe.$Maybe("Just", v2[0]._1[0]._1); }
  return Data$dMaybe.Nothing;
};
const effect_unsafe_unsafePerformEffect = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Effect.Unsafe"), "unsafePerformEffect"),
  effectUnsafePerform
);
const effectRefWrite = v => v1 => v2 => {
  if (v2.length === 1 && v2[0].tag === "ExternApp" && v2[0]._1.length === 2) {
    const $0 = v2[0]._1[1];
    return Data$dMaybe.$Maybe(
      "Just",
      PureScript$dBackend$dOptimizer$dSemantics.makeLet(Data$dMaybe.Nothing)(v2[0]._1[0])(val$p => PureScript$dBackend$dOptimizer$dSemantics.makeLet(Data$dMaybe.Nothing)($0)(ref$p => PureScript$dBackend$dOptimizer$dSemantics.$BackendSemantics(
        "NeutPrimEffect",
        PureScript$dBackend$dOptimizer$dSyntax.$BackendEffect("EffectRefWrite", ref$p, val$p)
      )))
    );
  }
  return Data$dMaybe.Nothing;
};
const effect_ref_write = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Effect.Ref"), "write"),
  effectRefWrite
);
const effectRefRead = v => v1 => v2 => {
  if (v2.length === 1 && v2[0].tag === "ExternApp" && v2[0]._1.length === 1) {
    return Data$dMaybe.$Maybe(
      "Just",
      PureScript$dBackend$dOptimizer$dSemantics.makeLet(Data$dMaybe.Nothing)(v2[0]._1[0])(val$p => PureScript$dBackend$dOptimizer$dSemantics.$BackendSemantics(
        "NeutPrimEffect",
        PureScript$dBackend$dOptimizer$dSyntax.$BackendEffect("EffectRefRead", val$p)
      ))
    );
  }
  return Data$dMaybe.Nothing;
};
const effect_ref_read = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Effect.Ref"), "read"),
  effectRefRead
);
const effectRefNew = v => v1 => v2 => {
  if (v2.length === 1 && v2[0].tag === "ExternApp" && v2[0]._1.length === 1) {
    return Data$dMaybe.$Maybe(
      "Just",
      PureScript$dBackend$dOptimizer$dSemantics.makeLet(Data$dMaybe.Nothing)(v2[0]._1[0])(val$p => PureScript$dBackend$dOptimizer$dSemantics.$BackendSemantics(
        "NeutPrimEffect",
        PureScript$dBackend$dOptimizer$dSyntax.$BackendEffect("EffectRefNew", val$p)
      ))
    );
  }
  return Data$dMaybe.Nothing;
};
const effect_ref_new = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Effect.Ref"), "_new"),
  effectRefNew
);
const effectRefModify = env => v => v1 => {
  if (v1.length === 1 && v1[0].tag === "ExternApp" && v1[0]._1.length === 2) {
    const $0 = v1[0]._1[1];
    return Data$dMaybe.$Maybe(
      "Just",
      PureScript$dBackend$dOptimizer$dSemantics.makeLet(Data$dMaybe.Nothing)(v1[0]._1[0])(fn$p => PureScript$dBackend$dOptimizer$dSemantics.makeLet(Data$dMaybe.Nothing)($0)(ref$p => PureScript$dBackend$dOptimizer$dSemantics.makeEffectBind(Data$dMaybe.Nothing)(PureScript$dBackend$dOptimizer$dSemantics.$BackendSemantics(
        "NeutPrimEffect",
        PureScript$dBackend$dOptimizer$dSyntax.$BackendEffect("EffectRefRead", ref$p)
      ))(val => PureScript$dBackend$dOptimizer$dSemantics.$BackendSemantics(
        "NeutPrimEffect",
        PureScript$dBackend$dOptimizer$dSyntax.$BackendEffect("EffectRefWrite", ref$p, PureScript$dBackend$dOptimizer$dSemantics.evalApp(env)(fn$p)([val]))
      ))))
    );
  }
  return Data$dMaybe.Nothing;
};
const effect_ref_modify = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Effect.Ref"), "modify"),
  effectRefModify
);
const effectPure = v => v1 => v2 => {
  if (v2.length === 1 && v2[0].tag === "ExternApp" && v2[0]._1.length === 1) {
    return Data$dMaybe.$Maybe("Just", PureScript$dBackend$dOptimizer$dSemantics.makeLet(Data$dMaybe.Nothing)(v2[0]._1[0])(PureScript$dBackend$dOptimizer$dSemantics.SemEffectPure));
  }
  return Data$dMaybe.Nothing;
};
const effect_pureE = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Effect"), "pureE"),
  effectPure
);
const effectMap = env => v => v1 => {
  if (v1.length === 1 && v1[0].tag === "ExternApp" && v1[0]._1.length === 1) {
    return Data$dMaybe.$Maybe(
      "Just",
      PureScript$dBackend$dOptimizer$dSemantics.makeLet(Data$dMaybe.Nothing)(v1[0]._1[0])(fn$p => PureScript$dBackend$dOptimizer$dSemantics.$BackendSemantics(
        "SemLam",
        Data$dMaybe.Nothing,
        val => PureScript$dBackend$dOptimizer$dSemantics.makeEffectBind(Data$dMaybe.Nothing)(val)(nextVal => PureScript$dBackend$dOptimizer$dSemantics.$BackendSemantics(
          "SemEffectPure",
          PureScript$dBackend$dOptimizer$dSemantics.evalApp(env)(fn$p)([nextVal])
        ))
      ))
    );
  }
  return Data$dMaybe.Nothing;
};
const effectBind = env => v => v1 => {
  if (v1.length === 1 && v1[0].tag === "ExternApp" && v1[0]._1.length === 2) {
    if (v1[0]._1[1].tag === "SemLam") {
      const $0 = v1[0]._1[1]._1;
      const $1 = v1[0]._1[1]._2;
      return Data$dMaybe.$Maybe(
        "Just",
        PureScript$dBackend$dOptimizer$dSemantics.makeLet(Data$dMaybe.Nothing)(v1[0]._1[0])(nextEff => PureScript$dBackend$dOptimizer$dSemantics.makeEffectBind($0)(nextEff)($1))
      );
    }
    const $0 = v1[0]._1[1];
    return Data$dMaybe.$Maybe(
      "Just",
      PureScript$dBackend$dOptimizer$dSemantics.makeLet(Data$dMaybe.Nothing)(v1[0]._1[0])(nextEff => PureScript$dBackend$dOptimizer$dSemantics.makeLet(Data$dMaybe.Nothing)($0)(nextK => PureScript$dBackend$dOptimizer$dSemantics.makeEffectBind(Data$dMaybe.Nothing)(nextEff)(a => PureScript$dBackend$dOptimizer$dSemantics.evalApp(env)(nextK)([
        a
      ]))))
    );
  }
  return Data$dMaybe.Nothing;
};
const effect_bindE = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Effect"), "bindE"),
  effectBind
);
const data_string_codePoints_toCodePointArray = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.String.CodePoints"), "toCodePointArray"),
  v => v1 => v2 => {
    if (v2.length === 1 && v2[0].tag === "ExternApp" && v2[0]._1.length === 1 && v2[0]._1[0].tag === "NeutLit" && v2[0]._1[0]._1.tag === "LitString") {
      return Data$dMaybe.$Maybe(
        "Just",
        PureScript$dBackend$dOptimizer$dSemantics.$BackendSemantics(
          "NeutLit",
          PureScript$dBackend$dOptimizer$dCoreFn.$Literal(
            "LitArray",
            Data$dFunctor.arrayMap(x => PureScript$dBackend$dOptimizer$dSemantics.$BackendSemantics("NeutLit", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitInt", x)))(Data$dString$dCodePoints.toCodePointArray(v2[0]._1[0]._1._1))
          )
        )
      );
    }
    return Data$dMaybe.Nothing;
  }
);
const data_semiring_numMul = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.Semiring"), "numMul"),
  /* #__PURE__ */ primBinaryOperator(/* #__PURE__ */ PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpNumberNum", PureScript$dBackend$dOptimizer$dSyntax.OpMultiply))
);
const data_semiring_numAdd = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.Semiring"), "numAdd"),
  /* #__PURE__ */ primBinaryOperator(/* #__PURE__ */ PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpNumberNum", PureScript$dBackend$dOptimizer$dSyntax.OpAdd))
);
const data_semiring_intMul = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.Semiring"), "intMul"),
  /* #__PURE__ */ primBinaryOperator(/* #__PURE__ */ PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpIntNum", PureScript$dBackend$dOptimizer$dSyntax.OpMultiply))
);
const data_semiring_intAdd = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.Semiring"), "intAdd"),
  /* #__PURE__ */ primBinaryOperator(/* #__PURE__ */ PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpIntNum", PureScript$dBackend$dOptimizer$dSyntax.OpAdd))
);
const data_semigroup_concatString = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.Semigroup"), "concatString"),
  /* #__PURE__ */ primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.OpStringAppend)
);
const data_semigroup_concatArray = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.Semigroup"), "concatArray"),
  env => qual => v => {
    if (v.length === 1 && v[0].tag === "ExternApp" && v[0]._1.length === 2) {
      if (v[0]._1[0].tag === "NeutLit" && v[0]._1[0]._1.tag === "LitArray" && v[0]._1[1].tag === "NeutLit" && v[0]._1[1]._1.tag === "LitArray") {
        return Data$dMaybe.$Maybe(
          "Just",
          PureScript$dBackend$dOptimizer$dSemantics.$BackendSemantics(
            "NeutLit",
            PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitArray", [...v[0]._1[0]._1._1, ...v[0]._1[1]._1._1])
          )
        );
      }
      return Data$dMaybe.$Maybe("Just", PureScript$dBackend$dOptimizer$dSemantics.evalAssocOp(env)(Data$dEither.$Either("Left", qual))(v[0]._1[0])(v[0]._1[1]));
    }
    return Data$dMaybe.Nothing;
  }
);
const data_ring_numSub = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.Ring"), "numSub"),
  /* #__PURE__ */ primBinaryOperator(/* #__PURE__ */ PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpNumberNum", PureScript$dBackend$dOptimizer$dSyntax.OpSubtract))
);
const data_ring_intSub = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.Ring"), "intSub"),
  /* #__PURE__ */ primBinaryOperator(/* #__PURE__ */ PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpIntNum", PureScript$dBackend$dOptimizer$dSyntax.OpSubtract))
);
const data_ord_ordString = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.Ord"), "ordString"),
  /* #__PURE__ */ primOrdOperator(PureScript$dBackend$dOptimizer$dSyntax.OpStringOrd)
);
const data_ord_ordNumber = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.Ord"), "ordNumber"),
  /* #__PURE__ */ primOrdOperator(PureScript$dBackend$dOptimizer$dSyntax.OpNumberOrd)
);
const data_ord_ordInt = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.Ord"), "ordInt"),
  /* #__PURE__ */ primOrdOperator(PureScript$dBackend$dOptimizer$dSyntax.OpIntOrd)
);
const data_ord_ordChar = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.Ord"), "ordChar"),
  /* #__PURE__ */ primOrdOperator(PureScript$dBackend$dOptimizer$dSyntax.OpCharOrd)
);
const data_ord_ordBoolean = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.Ord"), "ordBoolean"),
  /* #__PURE__ */ primOrdOperator(PureScript$dBackend$dOptimizer$dSyntax.OpBooleanOrd)
);
const data_int_bits_zshr = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.Int.Bits"), "zshr"),
  /* #__PURE__ */ primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.OpIntBitZeroFillShiftRight)
);
const data_int_bits_xor = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.Int.Bits"), "xor"),
  /* #__PURE__ */ primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.OpIntBitXor)
);
const data_int_bits_shr = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.Int.Bits"), "shr"),
  /* #__PURE__ */ primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.OpIntBitShiftRight)
);
const data_int_bits_shl = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.Int.Bits"), "shl"),
  /* #__PURE__ */ primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.OpIntBitShiftLeft)
);
const data_int_bits_or = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.Int.Bits"), "or"),
  /* #__PURE__ */ primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.OpIntBitOr)
);
const data_int_bits_complement = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.Int.Bits"), "complement"),
  /* #__PURE__ */ primUnaryOperator(PureScript$dBackend$dOptimizer$dSyntax.OpIntBitNot)
);
const data_int_bits_and = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.Int.Bits"), "and"),
  /* #__PURE__ */ primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.OpIntBitAnd)
);
const data_heytingAlgebra_boolNot = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.HeytingAlgebra"), "boolNot"),
  /* #__PURE__ */ primUnaryOperator(PureScript$dBackend$dOptimizer$dSyntax.OpBooleanNot)
);
const data_heytingAlgebra_boolDisj = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.HeytingAlgebra"), "boolDisj"),
  /* #__PURE__ */ primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.OpBooleanOr)
);
const data_heytingAlgebra_boolConj = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.HeytingAlgebra"), "boolConj"),
  /* #__PURE__ */ primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.OpBooleanAnd)
);
const data_function_uncurried_runFn = n => {
  const goRunFn = env => n$p => head => tail => {
    if (n$p <= 0) { return PureScript$dBackend$dOptimizer$dSemantics.evalUncurriedApp(env)(head)(tail); }
    return PureScript$dBackend$dOptimizer$dSemantics.$BackendSemantics("SemLam", Data$dMaybe.Nothing, val => goRunFn(env)(n$p - 1 | 0)(head)(Data$dArray.snoc(tail)(val)));
  };
  return Data$dTuple.$Tuple(
    PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.Function.Uncurried"), "runFn" + Data$dShow.showIntImpl(n)),
    env => v => v1 => {
      if (v1.length === 1 && v1[0].tag === "ExternApp") {
        const $0 = Data$dArray.unconsImpl(v$1 => Data$dMaybe.Nothing, x => xs => Data$dMaybe.$Maybe("Just", {head: x, tail: xs}), v1[0]._1);
        if ($0.tag === "Just") { return Data$dMaybe.$Maybe("Just", goRunFn(env)(n - $0._1.tail.length | 0)($0._1.head)($0._1.tail)); }
      }
      return Data$dMaybe.Nothing;
    }
  );
};
const data_function_uncurried_mkFn = n => Data$dTuple.$Tuple(
  PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.Function.Uncurried"), "mkFn" + Data$dShow.showIntImpl(n)),
  env => v => v1 => {
    if (v1.length === 1 && v1[0].tag === "ExternApp" && v1[0]._1.length === 1) {
      return Data$dMaybe.$Maybe(
        "Just",
        PureScript$dBackend$dOptimizer$dSemantics.$BackendSemantics("SemMkFn", PureScript$dBackend$dOptimizer$dSemantics.evalMkFn(env)(n)(v1[0]._1[0]))
      );
    }
    return Data$dMaybe.Nothing;
  }
);
const data_euclideanRing_numDiv = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.EuclideanRing"), "numDiv"),
  /* #__PURE__ */ primBinaryOperator(/* #__PURE__ */ PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpNumberNum", PureScript$dBackend$dOptimizer$dSyntax.OpDivide))
);
const data_euclideanRing_intDiv = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.EuclideanRing"), "intDiv"),
  /* #__PURE__ */ primBinaryOperator(/* #__PURE__ */ PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpIntNum", PureScript$dBackend$dOptimizer$dSyntax.OpDivide))
);
const data_eq_eqStringImpl = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.Eq"), "eqStringImpl"),
  /* #__PURE__ */ primBinaryOperator(/* #__PURE__ */ PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpStringOrd", PureScript$dBackend$dOptimizer$dSyntax.OpEq))
);
const data_eq_eqNumberImpl = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.Eq"), "eqNumberImpl"),
  /* #__PURE__ */ primBinaryOperator(/* #__PURE__ */ PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpNumberOrd", PureScript$dBackend$dOptimizer$dSyntax.OpEq))
);
const data_eq_eqIntImpl = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.Eq"), "eqIntImpl"),
  /* #__PURE__ */ primBinaryOperator(/* #__PURE__ */ PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpIntOrd", PureScript$dBackend$dOptimizer$dSyntax.OpEq))
);
const data_eq_eqCharImpl = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.Eq"), "eqCharImpl"),
  /* #__PURE__ */ primBinaryOperator(/* #__PURE__ */ PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpCharOrd", PureScript$dBackend$dOptimizer$dSyntax.OpEq))
);
const data_eq_eqBooleanImpl = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.Eq"), "eqBooleanImpl"),
  /* #__PURE__ */ primBinaryOperator(/* #__PURE__ */ PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpBooleanOrd", PureScript$dBackend$dOptimizer$dSyntax.OpEq))
);
const data_array_unsafeIndexImpl = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.Array"), "unsafeIndexImpl"),
  env => v => v1 => {
    if (v1.length === 1) {
      if (v1[0].tag === "ExternUncurriedApp") {
        if (v1[0]._1.length === 2) {
          const $0 = v1[0]._1[1];
          return Data$dMaybe.$Maybe(
            "Just",
            PureScript$dBackend$dOptimizer$dSemantics.makeLet(Data$dMaybe.Nothing)(v1[0]._1[0])(a$p => PureScript$dBackend$dOptimizer$dSemantics.makeLet(Data$dMaybe.Nothing)($0)(b$p => PureScript$dBackend$dOptimizer$dSemantics.evalPrimOp(env)(PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator(
              "Op2",
              PureScript$dBackend$dOptimizer$dSyntax.OpArrayIndex,
              a$p,
              b$p
            ))))
          );
        }
        return Data$dMaybe.Nothing;
      }
      if (v1[0].tag === "ExternApp" && v1[0]._1.length === 1) {
        return Data$dMaybe.$Maybe(
          "Just",
          PureScript$dBackend$dOptimizer$dSemantics.makeLet(Data$dMaybe.Nothing)(v1[0]._1[0])(a$p => PureScript$dBackend$dOptimizer$dSemantics.$BackendSemantics(
            "SemLam",
            Data$dMaybe.Nothing,
            b$p => PureScript$dBackend$dOptimizer$dSemantics.evalPrimOp(env)(PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator(
              "Op2",
              PureScript$dBackend$dOptimizer$dSyntax.OpArrayIndex,
              a$p,
              b$p
            ))
          ))
        );
      }
    }
    return Data$dMaybe.Nothing;
  }
);
const data_array_length = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Data.Array"), "length"),
  /* #__PURE__ */ primUnaryOperator(PureScript$dBackend$dOptimizer$dSyntax.OpArrayLength)
);
const control_monad_st_uncurried_runSTFn = /* #__PURE__ */ runEffectFn("Control.Monad.ST.Uncurried")("runSTFn");
const control_monad_st_uncurried_mkSTFn = /* #__PURE__ */ mkEffectFn("Control.Monad.ST.Uncurried")("mkSTFn");
const control_monad_st_internal_write = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Control.Monad.ST.Internal"), "write"),
  effectRefWrite
);
const control_monad_st_internal_run = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Control.Monad.ST.Internal"), "run"),
  effectUnsafePerform
);
const control_monad_st_internal_read = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Control.Monad.ST.Internal"), "read"),
  effectRefRead
);
const control_monad_st_internal_pure = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Control.Monad.ST.Internal"), "pure_"),
  effectPure
);
const control_monad_st_internal_new = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Control.Monad.ST.Internal"), "new"),
  effectRefNew
);
const control_monad_st_internal_modify = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Control.Monad.ST.Internal"), "modify"),
  effectRefModify
);
const control_monad_st_internal_map = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Control.Monad.ST.Internal"), "map_"),
  effectMap
);
const control_monad_st_internal_bind = /* #__PURE__ */ Data$dTuple.$Tuple(
  /* #__PURE__ */ PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(/* #__PURE__ */ Data$dMaybe.$Maybe("Just", "Control.Monad.ST.Internal"), "bind_"),
  effectBind
);
const coreForeignSemantics = /* #__PURE__ */ (() => {
  const oneToTen = Data$dArray.rangeImpl(1, 10);
  return Data$dMap$dInternal.fromFoldable(PureScript$dBackend$dOptimizer$dCoreFn.ordQualified(Data$dOrd.ordString))(Data$dFoldable.foldableArray)([
    Data$dTuple.$Tuple(PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Control.Monad.ST.Internal"), "bind_"), effectBind),
    Data$dTuple.$Tuple(PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Control.Monad.ST.Internal"), "map_"), effectMap),
    Data$dTuple.$Tuple(PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Control.Monad.ST.Internal"), "modify"), effectRefModify),
    Data$dTuple.$Tuple(PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Control.Monad.ST.Internal"), "new"), effectRefNew),
    Data$dTuple.$Tuple(PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Control.Monad.ST.Internal"), "pure_"), effectPure),
    Data$dTuple.$Tuple(PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Control.Monad.ST.Internal"), "read"), effectRefRead),
    Data$dTuple.$Tuple(PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Control.Monad.ST.Internal"), "run"), effectUnsafePerform),
    Data$dTuple.$Tuple(PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Control.Monad.ST.Internal"), "write"), effectRefWrite),
    Data$dTuple.$Tuple(
      PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.Array"), "length"),
      primUnaryOperator(PureScript$dBackend$dOptimizer$dSyntax.OpArrayLength)
    ),
    data_array_unsafeIndexImpl,
    Data$dTuple.$Tuple(
      PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.Eq"), "eqBooleanImpl"),
      primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpBooleanOrd", PureScript$dBackend$dOptimizer$dSyntax.OpEq))
    ),
    Data$dTuple.$Tuple(
      PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.Eq"), "eqCharImpl"),
      primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpCharOrd", PureScript$dBackend$dOptimizer$dSyntax.OpEq))
    ),
    Data$dTuple.$Tuple(
      PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.Eq"), "eqIntImpl"),
      primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpIntOrd", PureScript$dBackend$dOptimizer$dSyntax.OpEq))
    ),
    Data$dTuple.$Tuple(
      PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.Eq"), "eqNumberImpl"),
      primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpNumberOrd", PureScript$dBackend$dOptimizer$dSyntax.OpEq))
    ),
    Data$dTuple.$Tuple(
      PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.Eq"), "eqStringImpl"),
      primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpStringOrd", PureScript$dBackend$dOptimizer$dSyntax.OpEq))
    ),
    Data$dTuple.$Tuple(
      PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.EuclideanRing"), "intDiv"),
      primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpIntNum", PureScript$dBackend$dOptimizer$dSyntax.OpDivide))
    ),
    Data$dTuple.$Tuple(
      PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.EuclideanRing"), "numDiv"),
      primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpNumberNum", PureScript$dBackend$dOptimizer$dSyntax.OpDivide))
    ),
    Data$dTuple.$Tuple(
      PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.HeytingAlgebra"), "boolConj"),
      primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.OpBooleanAnd)
    ),
    Data$dTuple.$Tuple(
      PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.HeytingAlgebra"), "boolDisj"),
      primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.OpBooleanOr)
    ),
    Data$dTuple.$Tuple(
      PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.HeytingAlgebra"), "boolNot"),
      primUnaryOperator(PureScript$dBackend$dOptimizer$dSyntax.OpBooleanNot)
    ),
    Data$dTuple.$Tuple(
      PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.Int.Bits"), "and"),
      primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.OpIntBitAnd)
    ),
    Data$dTuple.$Tuple(
      PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.Int.Bits"), "complement"),
      primUnaryOperator(PureScript$dBackend$dOptimizer$dSyntax.OpIntBitNot)
    ),
    Data$dTuple.$Tuple(
      PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.Int.Bits"), "or"),
      primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.OpIntBitOr)
    ),
    Data$dTuple.$Tuple(
      PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.Int.Bits"), "shl"),
      primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.OpIntBitShiftLeft)
    ),
    Data$dTuple.$Tuple(
      PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.Int.Bits"), "shr"),
      primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.OpIntBitShiftRight)
    ),
    Data$dTuple.$Tuple(
      PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.Int.Bits"), "xor"),
      primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.OpIntBitXor)
    ),
    Data$dTuple.$Tuple(
      PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.Int.Bits"), "zshr"),
      primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.OpIntBitZeroFillShiftRight)
    ),
    Data$dTuple.$Tuple(
      PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.Ord"), "ordBoolean"),
      primOrdOperator(PureScript$dBackend$dOptimizer$dSyntax.OpBooleanOrd)
    ),
    Data$dTuple.$Tuple(
      PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.Ord"), "ordChar"),
      primOrdOperator(PureScript$dBackend$dOptimizer$dSyntax.OpCharOrd)
    ),
    Data$dTuple.$Tuple(
      PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.Ord"), "ordInt"),
      primOrdOperator(PureScript$dBackend$dOptimizer$dSyntax.OpIntOrd)
    ),
    Data$dTuple.$Tuple(
      PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.Ord"), "ordNumber"),
      primOrdOperator(PureScript$dBackend$dOptimizer$dSyntax.OpNumberOrd)
    ),
    Data$dTuple.$Tuple(
      PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.Ord"), "ordString"),
      primOrdOperator(PureScript$dBackend$dOptimizer$dSyntax.OpStringOrd)
    ),
    Data$dTuple.$Tuple(
      PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.Ring"), "intSub"),
      primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpIntNum", PureScript$dBackend$dOptimizer$dSyntax.OpSubtract))
    ),
    Data$dTuple.$Tuple(
      PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.Ring"), "numSub"),
      primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpNumberNum", PureScript$dBackend$dOptimizer$dSyntax.OpSubtract))
    ),
    data_semigroup_concatArray,
    Data$dTuple.$Tuple(
      PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.Semigroup"), "concatString"),
      primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.OpStringAppend)
    ),
    Data$dTuple.$Tuple(
      PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.Semiring"), "intAdd"),
      primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpIntNum", PureScript$dBackend$dOptimizer$dSyntax.OpAdd))
    ),
    Data$dTuple.$Tuple(
      PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.Semiring"), "intMul"),
      primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpIntNum", PureScript$dBackend$dOptimizer$dSyntax.OpMultiply))
    ),
    Data$dTuple.$Tuple(
      PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.Semiring"), "numAdd"),
      primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpNumberNum", PureScript$dBackend$dOptimizer$dSyntax.OpAdd))
    ),
    Data$dTuple.$Tuple(
      PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Data.Semiring"), "numMul"),
      primBinaryOperator(PureScript$dBackend$dOptimizer$dSyntax.$BackendOperator2("OpNumberNum", PureScript$dBackend$dOptimizer$dSyntax.OpMultiply))
    ),
    data_string_codePoints_toCodePointArray,
    Data$dTuple.$Tuple(PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Effect"), "bindE"), effectBind),
    Data$dTuple.$Tuple(PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Effect"), "pureE"), effectPure),
    Data$dTuple.$Tuple(PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Effect.Ref"), "modify"), effectRefModify),
    Data$dTuple.$Tuple(PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Effect.Ref"), "_new"), effectRefNew),
    Data$dTuple.$Tuple(PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Effect.Ref"), "read"), effectRefRead),
    Data$dTuple.$Tuple(PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Effect.Ref"), "write"), effectRefWrite),
    Data$dTuple.$Tuple(PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", "Effect.Unsafe"), "unsafePerformEffect"), effectUnsafePerform),
    partial_unsafe_unsafePartial,
    record_builder_copyRecord,
    record_builder_unsafeDelete,
    record_builder_unsafeInsert,
    record_builder_unsafeModify,
    record_builder_unsafeRename,
    record_unsafe_union_unsafeUnionFn,
    record_unsafe_unsafeDelete,
    record_unsafe_unsafeGet,
    record_unsafe_unsafeHas,
    record_unsafe_unsafeSet,
    unsafe_coerce_unsafeCoerce,
    ...Data$dFunctor.arrayMap(data_function_uncurried_mkFn)(oneToTen),
    ...Data$dFunctor.arrayMap(data_function_uncurried_runFn)(oneToTen),
    ...Data$dFunctor.arrayMap(mkEffectFn("Effect.Uncurried")("mkEffectFn"))(oneToTen),
    ...Data$dFunctor.arrayMap(runEffectFn("Effect.Uncurried")("runEffectFn"))(oneToTen),
    ...Data$dFunctor.arrayMap(mkEffectFn("Control.Monad.ST.Uncurried")("mkSTFn"))(oneToTen),
    ...Data$dFunctor.arrayMap(runEffectFn("Control.Monad.ST.Uncurried")("runSTFn"))(oneToTen)
  ]);
})();
export {
  control_monad_st_internal_bind,
  control_monad_st_internal_map,
  control_monad_st_internal_modify,
  control_monad_st_internal_new,
  control_monad_st_internal_pure,
  control_monad_st_internal_read,
  control_monad_st_internal_run,
  control_monad_st_internal_write,
  control_monad_st_uncurried_mkSTFn,
  control_monad_st_uncurried_runSTFn,
  coreForeignSemantics,
  data_array_length,
  data_array_unsafeIndexImpl,
  data_eq_eqBooleanImpl,
  data_eq_eqCharImpl,
  data_eq_eqIntImpl,
  data_eq_eqNumberImpl,
  data_eq_eqStringImpl,
  data_euclideanRing_intDiv,
  data_euclideanRing_numDiv,
  data_function_uncurried_mkFn,
  data_function_uncurried_runFn,
  data_heytingAlgebra_boolConj,
  data_heytingAlgebra_boolDisj,
  data_heytingAlgebra_boolNot,
  data_int_bits_and,
  data_int_bits_complement,
  data_int_bits_or,
  data_int_bits_shl,
  data_int_bits_shr,
  data_int_bits_xor,
  data_int_bits_zshr,
  data_ord_ordBoolean,
  data_ord_ordChar,
  data_ord_ordInt,
  data_ord_ordNumber,
  data_ord_ordString,
  data_ring_intSub,
  data_ring_numSub,
  data_semigroup_concatArray,
  data_semigroup_concatString,
  data_semiring_intAdd,
  data_semiring_intMul,
  data_semiring_numAdd,
  data_semiring_numMul,
  data_string_codePoints_toCodePointArray,
  effectBind,
  effectMap,
  effectPure,
  effectRefModify,
  effectRefNew,
  effectRefRead,
  effectRefWrite,
  effectUnsafePerform,
  effect_bindE,
  effect_pureE,
  effect_ref_modify,
  effect_ref_new,
  effect_ref_read,
  effect_ref_write,
  effect_uncurried_mkEffectFn,
  effect_uncurried_runEffectFn,
  effect_unsafe_unsafePerformEffect,
  fromFoldable,
  isQualified,
  mkEffectFn,
  partial_unsafe_unsafePartial,
  primBinaryOperator,
  primOrdOperator,
  primUnaryOperator,
  qualified,
  record_builder_copyRecord,
  record_builder_unsafeDelete,
  record_builder_unsafeInsert,
  record_builder_unsafeModify,
  record_builder_unsafeRename,
  record_unsafe_union_unsafeUnionFn,
  record_unsafe_unsafeDelete,
  record_unsafe_unsafeGet,
  record_unsafe_unsafeHas,
  record_unsafe_unsafeSet,
  runEffectFn,
  unsafe_coerce_unsafeCoerce,
  viewCopyRecord
};
