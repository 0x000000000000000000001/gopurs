import * as $runtime from "../runtime.js";
import * as Data$dArray from "../Data.Array/index.js";
import * as Data$dFoldable from "../Data.Foldable/index.js";
import * as Data$dFunctor from "../Data.Functor/index.js";
import * as Data$dFunctorWithIndex from "../Data.FunctorWithIndex/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dMonoid from "../Data.Monoid/index.js";
import * as Data$dShow from "../Data.Show/index.js";
import * as Data$dString$dCodePoints from "../Data.String.CodePoints/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
import * as Dodo from "../Dodo/index.js";
import * as Dodo$dCommon from "../Dodo.Common/index.js";
import * as Dodo$dInternal from "../Dodo.Internal/index.js";
import * as PureScript$dBackend$dOptimizer$dCoreFn from "../PureScript.Backend.Optimizer.CoreFn/index.js";
import * as PureScript$dBackend$dOptimizer$dSemantics from "../PureScript.Backend.Optimizer.Semantics/index.js";
import * as PureScript$dBackend$dOptimizer$dSyntax from "../PureScript.Backend.Optimizer.Syntax/index.js";
const $Prec = tag => tag;
const fold = /* #__PURE__ */ (() => Data$dFoldable.foldableArray.foldMap(Dodo$dInternal.monoidDoc)(Data$dFoldable.identity))();
const lines = /* #__PURE__ */ Data$dFoldable.foldrArray(Dodo.appendBreak)(Dodo$dInternal.Empty);
const words = /* #__PURE__ */ Data$dFoldable.foldrArray(Dodo.appendSpace)(Dodo$dInternal.Empty);
const power = /* #__PURE__ */ Data$dMonoid.power(Data$dMonoid.monoidString);
const PrecBlock = /* #__PURE__ */ $Prec("PrecBlock");
const PrecApply = /* #__PURE__ */ $Prec("PrecApply");
const PrecAccess = /* #__PURE__ */ $Prec("PrecAccess");
const PrecAtom = /* #__PURE__ */ $Prec("PrecAtom");
const wrapPrec = prec1 => v => {
  if (
    (() => {
      if (v._1 === "PrecBlock") { return prec1 !== "PrecBlock"; }
      if (prec1 === "PrecBlock") { return false; }
      if (v._1 === "PrecApply") { return prec1 !== "PrecApply"; }
      if (prec1 === "PrecApply") { return false; }
      if (v._1 === "PrecAccess") { return prec1 !== "PrecAccess"; }
      if (prec1 === "PrecAccess") { return false; }
      if (v._1 === "PrecAtom" && prec1 === "PrecAtom") { return false; }
      $runtime.fail();
    })()
  ) {
    return Dodo$dCommon.pursParens(v._2);
  }
  return v._2;
};
const printUncurriedApp = isEffectful => fn => args => Data$dTuple.$Tuple(
  PrecApply,
  (() => {
    const $0 = fold([
      wrapPrec(PrecAccess)(fn),
      Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.Empty, Dodo$dInternal.Break),
      (() => {
        const $0 = fold([
          Dodo$dCommon.pursParens(Dodo.foldWithSeparator(Data$dFoldable.foldableArray)(Dodo$dInternal.$Doc(
            "FlexAlt",
            Dodo$dInternal.$Doc("Text", 2, ", "),
            Dodo$dInternal.$Doc("Append", Dodo$dInternal.Break, Dodo$dInternal.$Doc("Text", 2, ", "))
          ))(Data$dFunctor.arrayMap(Data$dTuple.snd)(args))),
          isEffectful ? Dodo$dInternal.$Doc("Text", 1, "!") : Dodo$dInternal.Empty
        ]);
        if ($0.tag === "Empty") { return Dodo$dInternal.Empty; }
        return Dodo$dInternal.$Doc("Indent", $0);
      })()
    ]);
    if ($0.tag === "Empty") { return Dodo$dInternal.Empty; }
    if ($0.tag === "FlexSelect" && $0._2.tag === "Empty" && $0._3.tag === "Empty") { return $0; }
    return Dodo$dInternal.$Doc("FlexSelect", $0, Dodo$dInternal.Empty, Dodo$dInternal.Empty);
  })()
);
const printUncurriedAbs = isEffectful => args => body => Data$dTuple.$Tuple(
  PrecBlock,
  (() => {
    const $0 = fold([
      Dodo$dInternal.$Doc("Text", 1, "\\"),
      (() => {
        const $0 = Dodo$dCommon.pursParens(Dodo.foldWithSeparator(Data$dFoldable.foldableArray)(Dodo$dInternal.$Doc(
          "FlexAlt",
          Dodo$dInternal.$Doc("Text", 2, ", "),
          Dodo$dInternal.$Doc("Append", Dodo$dInternal.Break, Dodo$dInternal.$Doc("Text", 2, ", "))
        ))(args));
        const $1 = $0.tag === "Empty" ? Dodo$dInternal.Empty : Dodo$dInternal.$Doc("WithPosition", pos => Dodo.align(pos.column - pos.nextIndent | 0)($0));
        if ($1.tag === "Empty") { return Dodo$dInternal.Empty; }
        if ($1.tag === "FlexSelect" && $1._2.tag === "Empty" && $1._3.tag === "Empty") { return $1; }
        return Dodo$dInternal.$Doc("FlexSelect", $1, Dodo$dInternal.Empty, Dodo$dInternal.Empty);
      })(),
      isEffectful ? Dodo$dInternal.$Doc("Text", 1, "!") : Dodo$dInternal.Empty,
      Dodo$dInternal.$Doc("Text", 1, " "),
      Dodo$dInternal.$Doc("Text", 2, "->"),
      Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 1, " "), Dodo$dInternal.Break),
      body._2.tag === "Empty" ? Dodo$dInternal.Empty : Dodo$dInternal.$Doc("Indent", body._2)
    ]);
    if ($0.tag === "Empty") { return Dodo$dInternal.Empty; }
    if ($0.tag === "FlexSelect" && $0._2.tag === "Empty" && $0._3.tag === "Empty") { return $0; }
    return Dodo$dInternal.$Doc("FlexSelect", $0, Dodo$dInternal.Empty, Dodo$dInternal.Empty);
  })()
);
const printRewrite = str => {
  const $0 = "{- " + str + " -}";
  if ($0 === "") { return Dodo$dInternal.Empty; }
  return Dodo$dInternal.$Doc("Text", Data$dString$dCodePoints.toCodePointArray($0).length, $0);
};
const printUnpackOpCase = v => {
  if (v.tag === "UnpackRecord") { return printRewrite("UnpackRecord"); }
  if (v.tag === "UnpackUpdate") { return printRewrite("UnpackUpdate"); }
  if (v.tag === "UnpackArray") { return printRewrite("UnpackArray"); }
  if (v.tag === "UnpackData") { return printRewrite("UnpackData"); }
  $runtime.fail();
};
const printRecord = sep => {
  const $0 = Dodo.foldWithSeparator(Data$dFoldable.foldableArray)(Dodo$dInternal.$Doc(
    "FlexAlt",
    Dodo$dInternal.$Doc("Text", 2, ", "),
    Dodo$dInternal.$Doc("Append", Dodo$dInternal.Break, Dodo$dInternal.$Doc("Text", 2, ", "))
  ));
  const $1 = Data$dFunctor.arrayMap(v => {
    const $1 = fold([
      (() => {
        const $1 = Data$dShow.showStringImpl(v._1);
        if ($1 === "") { return Dodo$dInternal.Empty; }
        return Dodo$dInternal.$Doc("Text", Data$dString$dCodePoints.toCodePointArray($1).length, $1);
      })(),
      sep,
      Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 1, " "), Dodo$dInternal.Break),
      Dodo.align(2)(v._2.tag === "Empty" ? Dodo$dInternal.Empty : Dodo$dInternal.$Doc("Indent", v._2))
    ]);
    if ($1.tag === "Empty") { return Dodo$dInternal.Empty; }
    if ($1.tag === "FlexSelect" && $1._2.tag === "Empty" && $1._3.tag === "Empty") { return $1; }
    return Dodo$dInternal.$Doc("FlexSelect", $1, Dodo$dInternal.Empty, Dodo$dInternal.Empty);
  });
  return x => Dodo$dCommon.pursCurlies($0($1(x)));
};
const printQualified = printA => v => {
  const $0 = printA(v._2);
  if (v._1.tag === "Nothing") { return $0; }
  if (v._1.tag === "Just") {
    const $1 = v._1._1 === "" ? Dodo$dInternal.Empty : Dodo$dInternal.$Doc("Text", Data$dString$dCodePoints.toCodePointArray(v._1._1).length, v._1._1);
    const $2 = printA(v._2);
    const $3 = (() => {
      if ($2.tag === "Empty") { return Dodo$dInternal.$Doc("Text", 1, "."); }
      if ($2.tag === "Text") { return Dodo$dInternal.$Doc("Text", 1 + $2._1 | 0, "." + $2._2); }
      return Dodo$dInternal.$Doc("Append", Dodo$dInternal.$Doc("Text", 1, "."), $2);
    })();
    if ($1.tag === "Empty") { return $3; }
    if ($3.tag === "Empty") { return $1; }
    if ($1.tag === "Text" && $3.tag === "Text") { return Dodo$dInternal.$Doc("Text", $1._1 + $3._1 | 0, $1._2 + $3._2); }
    return Dodo$dInternal.$Doc("Append", $1, $3);
  }
  $runtime.fail();
};
const printLevel = v => {
  const $0 = Data$dShow.showIntImpl(v);
  if ($0 === "") { return Dodo$dInternal.Empty; }
  return Dodo$dInternal.$Doc("Text", Data$dString$dCodePoints.toCodePointArray($0).length, $0);
};
const printLet$p = letKeyword => bindings => {
  const $0 = fold([
    letKeyword === "" ? Dodo$dInternal.Empty : Dodo$dInternal.$Doc("Text", Data$dString$dCodePoints.toCodePointArray(letKeyword).length, letKeyword),
    Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 1, " "), Dodo$dInternal.Break),
    (() => {
      const $0 = lines(Data$dFunctor.arrayMap(v => {
        const $0 = fold([
          v._1,
          Dodo$dInternal.$Doc("Text", 1, " "),
          Dodo$dInternal.$Doc("Text", 1, "="),
          Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 1, " "), Dodo$dInternal.Break),
          v._2.tag === "Empty" ? Dodo$dInternal.Empty : Dodo$dInternal.$Doc("Indent", v._2)
        ]);
        if ($0.tag === "Empty") { return Dodo$dInternal.Empty; }
        if ($0.tag === "FlexSelect" && $0._2.tag === "Empty" && $0._3.tag === "Empty") { return $0; }
        return Dodo$dInternal.$Doc("FlexSelect", $0, Dodo$dInternal.Empty, Dodo$dInternal.Empty);
      })(bindings));
      if ($0.tag === "Empty") { return Dodo$dInternal.Empty; }
      return Dodo$dInternal.$Doc("Indent", $0);
    })(),
    Dodo$dInternal.$Doc("Text", 1, " "),
    Dodo$dInternal.$Doc("Text", 2, "in")
  ]);
  if ($0.tag === "Empty") { return Dodo$dInternal.Empty; }
  if ($0.tag === "FlexSelect" && $0._2.tag === "Empty" && $0._3.tag === "Empty") { return $0; }
  return Dodo$dInternal.$Doc("FlexSelect", $0, Dodo$dInternal.Empty, Dodo$dInternal.Empty);
};
const printIdent = v => {
  if (v === "") { return Dodo$dInternal.Empty; }
  return Dodo$dInternal.$Doc("Text", Data$dString$dCodePoints.toCodePointArray(v).length, v);
};
const printLocal = mbIdent => lvl => {
  if (mbIdent.tag === "Just" && mbIdent._1 === "$__unused") { return Dodo$dInternal.$Doc("Text", 1, "_"); }
  const $0 = (() => {
    if (mbIdent.tag === "Nothing") { return Dodo$dInternal.Empty; }
    if (mbIdent.tag === "Just") {
      if (mbIdent._1 === "") { return Dodo$dInternal.Empty; }
      return Dodo$dInternal.$Doc("Text", Data$dString$dCodePoints.toCodePointArray(mbIdent._1).length, mbIdent._1);
    }
    $runtime.fail();
  })();
  const $1 = printLevel(lvl);
  const $2 = (() => {
    if ($1.tag === "Empty") { return Dodo$dInternal.$Doc("Text", 1, "@"); }
    if ($1.tag === "Text") { return Dodo$dInternal.$Doc("Text", 1 + $1._1 | 0, "@" + $1._2); }
    return Dodo$dInternal.$Doc("Append", Dodo$dInternal.$Doc("Text", 1, "@"), $1);
  })();
  if ($0.tag === "Empty") { return $2; }
  if ($2.tag === "Empty") { return $0; }
  if ($0.tag === "Text" && $2.tag === "Text") { return Dodo$dInternal.$Doc("Text", $0._1 + $2._1 | 0, $0._2 + $2._2); }
  return Dodo$dInternal.$Doc("Append", $0, $2);
};
const printDistOpCase = v => {
  if (v.tag === "DistApp") { return printRewrite("DistApp"); }
  if (v.tag === "DistUncurriedApp") { return printRewrite("DistUncurriedApp"); }
  if (v.tag === "DistAccessor") { return printRewrite("DistAccessor"); }
  if (v.tag === "DistPrimOp1") { return printRewrite("DistPrimOp1"); }
  if (v.tag === "DistPrimOp2L") { return printRewrite("DistPrimOp2L"); }
  if (v.tag === "DistPrimOp2R") { return printRewrite("DistPrimOp2R"); }
  $runtime.fail();
};
const printCurriedApp = fn => args => Data$dTuple.$Tuple(
  PrecApply,
  (() => {
    const $0 = Dodo.foldWithSeparator(Data$dFoldable.foldableArray)(Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 1, " "), Dodo$dInternal.Break))([
      wrapPrec(PrecApply)(fn),
      ...Data$dFunctor.arrayMap(x => {
        const $0 = wrapPrec(PrecAccess)(x);
        if ($0.tag === "Empty") { return Dodo$dInternal.Empty; }
        return Dodo$dInternal.$Doc("Indent", $0);
      })(args)
    ]);
    if ($0.tag === "Empty") { return Dodo$dInternal.Empty; }
    if ($0.tag === "FlexSelect" && $0._2.tag === "Empty" && $0._3.tag === "Empty") { return $0; }
    return Dodo$dInternal.$Doc("FlexSelect", $0, Dodo$dInternal.Empty, Dodo$dInternal.Empty);
  })()
);
const printCurriedAbs = args => body => Data$dTuple.$Tuple(
  PrecBlock,
  (() => {
    const $0 = fold([
      Dodo$dInternal.$Doc("Text", 1, "\\"),
      (() => {
        const $0 = Dodo.foldWithSeparator(Data$dFoldable.foldableArray)(Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 1, " "), Dodo$dInternal.Break))(args);
        const $1 = $0.tag === "Empty" ? Dodo$dInternal.Empty : Dodo$dInternal.$Doc("WithPosition", pos => Dodo.align(pos.column - pos.nextIndent | 0)($0));
        if ($1.tag === "Empty") { return Dodo$dInternal.Empty; }
        if ($1.tag === "FlexSelect" && $1._2.tag === "Empty" && $1._3.tag === "Empty") { return $1; }
        return Dodo$dInternal.$Doc("FlexSelect", $1, Dodo$dInternal.Empty, Dodo$dInternal.Empty);
      })(),
      Dodo$dInternal.$Doc("Text", 1, " "),
      Dodo$dInternal.$Doc("Text", 2, "->"),
      Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 1, " "), Dodo$dInternal.Break),
      body._2.tag === "Empty" ? Dodo$dInternal.Empty : Dodo$dInternal.$Doc("Indent", body._2)
    ]);
    if ($0.tag === "Empty") { return Dodo$dInternal.Empty; }
    if ($0.tag === "FlexSelect" && $0._2.tag === "Empty" && $0._3.tag === "Empty") { return $0; }
    return Dodo$dInternal.$Doc("FlexSelect", $0, Dodo$dInternal.Empty, Dodo$dInternal.Empty);
  })()
);
const printBranch = conds => fallback => Data$dFoldable.foldrArray(v => {
  const $0 = v._1;
  const $1 = v._2;
  return other => fold([
    (() => {
      const $2 = fold([
        Dodo$dInternal.$Doc("Text", 2, "if"),
        Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 1, " "), Dodo$dInternal.Break),
        $0.tag === "Empty" ? Dodo$dInternal.Empty : Dodo$dInternal.$Doc("Indent", $0),
        Dodo$dInternal.$Doc("Text", 1, " "),
        Dodo$dInternal.$Doc("Text", 4, "then")
      ]);
      if ($2.tag === "Empty") { return Dodo$dInternal.Empty; }
      if ($2.tag === "FlexSelect" && $2._2.tag === "Empty" && $2._3.tag === "Empty") { return $2; }
      return Dodo$dInternal.$Doc("FlexSelect", $2, Dodo$dInternal.Empty, Dodo$dInternal.Empty);
    })(),
    $1.tag === "Empty" ? Dodo$dInternal.Empty : Dodo$dInternal.$Doc("Indent", $1),
    Dodo$dInternal.Break,
    Dodo$dInternal.$Doc("Text", 4, "else"),
    Dodo$dInternal.$Doc("Text", 1, " "),
    other
  ]);
})((() => {
  const $0 = fallback.tag === "Empty" ? Dodo$dInternal.Empty : Dodo$dInternal.$Doc("Indent", fallback);
  if ($0.tag === "Empty") { return Dodo$dInternal.Break; }
  return Dodo$dInternal.$Doc("Append", Dodo$dInternal.Break, $0);
})())(conds);
const printBackendRewriteCase = v => {
  if (v.tag === "RewriteInline") { return printRewrite("Inline"); }
  if (v.tag === "RewriteUncurry") { return printRewrite("Uncurry"); }
  if (v.tag === "RewriteStop") { return printRewrite("Stop"); }
  if (v.tag === "RewriteUnpackOp") { return printUnpackOpCase(v._3); }
  if (v.tag === "RewriteDistBranchesLet") { return printRewrite("DistLet"); }
  if (v.tag === "RewriteDistBranchesOp") { return printDistOpCase(v._3); }
  $runtime.fail();
};
const printBackendAccessor = v => {
  if (v.tag === "GetProp") {
    const $0 = "." + Data$dShow.showStringImpl(v._1);
    if ($0 === "") { return Dodo$dInternal.Empty; }
    return Dodo$dInternal.$Doc("Text", Data$dString$dCodePoints.toCodePointArray($0).length, $0);
  }
  if (v.tag === "GetIndex") {
    const $0 = "[" + Data$dShow.showIntImpl(v._1) + "]";
    if ($0 === "") { return Dodo$dInternal.Empty; }
    return Dodo$dInternal.$Doc("Text", Data$dString$dCodePoints.toCodePointArray($0).length, $0);
  }
  if (v.tag === "GetCtorField") {
    const $0 = "#" + Data$dShow.showIntImpl(v._6);
    if ($0 === "") { return Dodo$dInternal.Empty; }
    return Dodo$dInternal.$Doc("Text", Data$dString$dCodePoints.toCodePointArray($0).length, $0);
  }
  $runtime.fail();
};
const printArray = /* #__PURE__ */ (() => {
  const $0 = Dodo.foldWithSeparator(Data$dFoldable.foldableArray)(Dodo$dInternal.$Doc(
    "FlexAlt",
    Dodo$dInternal.$Doc("Text", 2, ", "),
    Dodo$dInternal.$Doc("Append", Dodo$dInternal.Break, Dodo$dInternal.$Doc("Text", 2, ", "))
  ));
  return x => Dodo$dCommon.pursSquares($0(x));
})();
const printLiteral = /* #__PURE__ */ (() => {
  const $0 = Data$dTuple.Tuple(PrecAtom);
  return x => $0((() => {
    if (x.tag === "LitInt") {
      const $1 = Data$dShow.showIntImpl(x._1);
      if ($1 === "") { return Dodo$dInternal.Empty; }
      return Dodo$dInternal.$Doc("Text", Data$dString$dCodePoints.toCodePointArray($1).length, $1);
    }
    if (x.tag === "LitNumber") {
      const $1 = Data$dShow.showNumberImpl(x._1);
      if ($1 === "") { return Dodo$dInternal.Empty; }
      return Dodo$dInternal.$Doc("Text", Data$dString$dCodePoints.toCodePointArray($1).length, $1);
    }
    if (x.tag === "LitString") {
      const $1 = Data$dShow.showStringImpl(x._1);
      if ($1 === "") { return Dodo$dInternal.Empty; }
      return Dodo$dInternal.$Doc("Text", Data$dString$dCodePoints.toCodePointArray($1).length, $1);
    }
    if (x.tag === "LitChar") {
      const $1 = Data$dShow.showCharImpl(x._1);
      if ($1 === "") { return Dodo$dInternal.Empty; }
      return Dodo$dInternal.$Doc("Text", Data$dString$dCodePoints.toCodePointArray($1).length, $1);
    }
    if (x.tag === "LitBoolean") {
      const $1 = x._1 ? "true" : "false";
      if ($1 === "") { return Dodo$dInternal.Empty; }
      return Dodo$dInternal.$Doc("Text", Data$dString$dCodePoints.toCodePointArray($1).length, $1);
    }
    if (x.tag === "LitArray") { return printArray(Data$dFunctor.arrayMap(Data$dTuple.snd)(x._1)); }
    if (x.tag === "LitRecord") {
      return printRecord(Dodo$dInternal.$Doc("Text", 1, ":"))(Data$dFunctor.arrayMap(m => PureScript$dBackend$dOptimizer$dCoreFn.$Prop(m._1, m._2._2))(x._1));
    }
    $runtime.fail();
  })());
})();
const primOp = ns => op => {
  if (ns === "") {
    return Data$dTuple.$Tuple(
      PrecAtom,
      (() => {
        const $0 = "#[prim." + op + "]";
        if ($0 === "") { return Dodo$dInternal.Empty; }
        return Dodo$dInternal.$Doc("Text", Data$dString$dCodePoints.toCodePointArray($0).length, $0);
      })()
    );
  }
  return Data$dTuple.$Tuple(
    PrecAtom,
    (() => {
      const $0 = "#[prim." + ns + "." + op + "]";
      if ($0 === "") { return Dodo$dInternal.Empty; }
      return Dodo$dInternal.$Doc("Text", Data$dString$dCodePoints.toCodePointArray($0).length, $0);
    })()
  );
};
const printBackendEffect = v => {
  if (v.tag === "EffectRefNew") { return printUncurriedApp(true)(primOp("ref")("new"))([v._1]); }
  if (v.tag === "EffectRefRead") { return printUncurriedApp(true)(primOp("ref")("read"))([v._1]); }
  if (v.tag === "EffectRefWrite") { return printUncurriedApp(true)(primOp("ref")("write"))([v._1, v._2]); }
  $runtime.fail();
};
const printBackendOperator1 = v => {
  if (v.tag === "OpBooleanNot") { return primOp("boolean")("not"); }
  if (v.tag === "OpIntBitNot") { return primOp("int")("bitnot"); }
  if (v.tag === "OpIntNegate") { return primOp("int")("negate"); }
  if (v.tag === "OpNumberNegate") { return primOp("number")("negate"); }
  if (v.tag === "OpArrayLength") { return primOp("array")("length"); }
  if (v.tag === "OpIsTag") {
    return Data$dTuple.$Tuple(
      PrecAtom,
      (() => {
        const $0 = printQualified(printIdent)(v._1);
        const $1 = (() => {
          if ($0.tag === "Empty") { return Dodo$dInternal.$Doc("Text", 1, "]"); }
          if ($0.tag === "Text") { return Dodo$dInternal.$Doc("Text", $0._1 + 1 | 0, $0._2 + "]"); }
          return Dodo$dInternal.$Doc("Append", $0, Dodo$dInternal.$Doc("Text", 1, "]"));
        })();
        if ($1.tag === "Empty") { return Dodo$dInternal.$Doc("Text", 13, "#[prim.istag "); }
        if ($1.tag === "Text") { return Dodo$dInternal.$Doc("Text", 13 + $1._1 | 0, "#[prim.istag " + $1._2); }
        return Dodo$dInternal.$Doc("Append", Dodo$dInternal.$Doc("Text", 13, "#[prim.istag "), $1);
      })()
    );
  }
  $runtime.fail();
};
const printBackendOperatorNum = ns => x => primOp(ns)((() => {
  if (x === "OpAdd") { return "add"; }
  if (x === "OpDivide") { return "div"; }
  if (x === "OpMultiply") { return "mul"; }
  if (x === "OpSubtract") { return "sub"; }
  $runtime.fail();
})());
const printBackendOperatorOrd = ns => x => primOp(ns)((() => {
  if (x === "OpEq") { return "eq"; }
  if (x === "OpNotEq") { return "neq"; }
  if (x === "OpGt") { return "gt"; }
  if (x === "OpGte") { return "gte"; }
  if (x === "OpLt") { return "lt"; }
  if (x === "OpLte") { return "lte"; }
  $runtime.fail();
})());
const printBackendOperator2 = v => {
  if (v.tag === "OpArrayIndex") { return primOp("array")("index"); }
  if (v.tag === "OpBooleanAnd") { return primOp("boolean")("and"); }
  if (v.tag === "OpBooleanOr") { return primOp("boolean")("or"); }
  if (v.tag === "OpBooleanOrd") { return printBackendOperatorOrd("boolean")(v._1); }
  if (v.tag === "OpCharOrd") { return printBackendOperatorOrd("char")(v._1); }
  if (v.tag === "OpIntBitAnd") { return primOp("int")("bitand"); }
  if (v.tag === "OpIntBitOr") { return primOp("int")("bitor"); }
  if (v.tag === "OpIntBitShiftLeft") { return primOp("int")("bitshl"); }
  if (v.tag === "OpIntBitShiftRight") { return primOp("int")("bitshr"); }
  if (v.tag === "OpIntBitXor") { return primOp("int")("bitxor"); }
  if (v.tag === "OpIntBitZeroFillShiftRight") { return primOp("int")("bitzshr"); }
  if (v.tag === "OpIntNum") { return printBackendOperatorNum("int")(v._1); }
  if (v.tag === "OpIntOrd") { return printBackendOperatorOrd("int")(v._1); }
  if (v.tag === "OpNumberNum") { return printBackendOperatorNum("number")(v._1); }
  if (v.tag === "OpNumberOrd") { return printBackendOperatorOrd("number")(v._1); }
  if (v.tag === "OpStringAppend") { return primOp("string")("append"); }
  if (v.tag === "OpStringOrd") { return printBackendOperatorOrd("string")(v._1); }
  $runtime.fail();
};
const printBackendOperator = v => {
  if (v.tag === "Op1") { return printUncurriedApp(false)(printBackendOperator1(v._1))([v._2]); }
  if (v.tag === "Op2") { return printUncurriedApp(false)(printBackendOperator2(v._1))([v._2, v._3]); }
  $runtime.fail();
};
const printBackendSyntax = v => {
  if (v.tag === "Var") { return Data$dTuple.$Tuple(PrecAtom, printQualified(printIdent)(v._1)); }
  if (v.tag === "Local") { return Data$dTuple.$Tuple(PrecAtom, printLocal(v._1)(v._2)); }
  if (v.tag === "Lit") { return printLiteral(v._1); }
  if (v.tag === "App") { return printCurriedApp(v._1)(v._2); }
  if (v.tag === "Abs") { return printCurriedAbs(Data$dFunctor.arrayMap(v$1 => printLocal(v$1._1)(v$1._2))(v._1))(v._2); }
  if (v.tag === "UncurriedApp") { return printUncurriedApp(false)(v._1)(v._2); }
  if (v.tag === "UncurriedAbs") { return printUncurriedAbs(false)(Data$dFunctor.arrayMap(v$1 => printLocal(v$1._1)(v$1._2))(v._1))(v._2); }
  if (v.tag === "UncurriedEffectApp") { return printUncurriedApp(true)(v._1)(v._2); }
  if (v.tag === "UncurriedEffectAbs") { return printUncurriedAbs(true)(Data$dFunctor.arrayMap(v$1 => printLocal(v$1._1)(v$1._2))(v._1))(v._2); }
  if (v.tag === "Accessor") { return Data$dTuple.$Tuple(PrecAccess, fold([wrapPrec(PrecAccess)(v._1), printBackendAccessor(v._2)])); }
  if (v.tag === "Update") {
    return Data$dTuple.$Tuple(
      PrecApply,
      (() => {
        const $0 = fold([
          wrapPrec(PrecAccess)(v._1),
          Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 1, " "), Dodo$dInternal.Break),
          (() => {
            const $0 = printRecord(Dodo$dInternal.$Doc("Text", 2, " ="))(Data$dFunctor.arrayMap(m => PureScript$dBackend$dOptimizer$dCoreFn.$Prop(m._1, m._2._2))(v._2));
            if ($0.tag === "Empty") { return Dodo$dInternal.Empty; }
            return Dodo$dInternal.$Doc("Indent", $0);
          })()
        ]);
        if ($0.tag === "Empty") { return Dodo$dInternal.Empty; }
        if ($0.tag === "FlexSelect" && $0._2.tag === "Empty" && $0._3.tag === "Empty") { return $0; }
        return Dodo$dInternal.$Doc("FlexSelect", $0, Dodo$dInternal.Empty, Dodo$dInternal.Empty);
      })()
    );
  }
  if (v.tag === "CtorSaturated") { return printUncurriedApp(false)(Data$dTuple.$Tuple(PrecAtom, printQualified(printIdent)(v._1)))(Data$dFunctor.arrayMap(Data$dTuple.snd)(v._5)); }
  if (v.tag === "CtorDef") {
    return Data$dTuple.$Tuple(
      PrecBlock,
      words([
        Dodo$dInternal.$Doc("Text", 11, "constructor"),
        (() => {
          const $0 = v._3 === "" ? Dodo$dInternal.Empty : Dodo$dInternal.$Doc("Text", Data$dString$dCodePoints.toCodePointArray(v._3).length, v._3);
          const $1 = Dodo.encloseWithSeparator(Data$dFoldable.foldableArray)(Dodo$dInternal.$Doc("Text", 1, "("))(Dodo$dInternal.$Doc("Text", 1, ")"))(Dodo$dInternal.$Doc(
            "Text",
            2,
            ", "
          ))(Data$dArray.replicateImpl(v._4.length, Dodo$dInternal.$Doc("Text", 1, "_")));
          if ($0.tag === "Empty") { return $1; }
          if ($1.tag === "Empty") { return $0; }
          if ($0.tag === "Text" && $1.tag === "Text") { return Dodo$dInternal.$Doc("Text", $0._1 + $1._1 | 0, $0._2 + $1._2); }
          return Dodo$dInternal.$Doc("Append", $0, $1);
        })(),
        Dodo$dInternal.$Doc("Text", 2, "of"),
        v._2 === "" ? Dodo$dInternal.Empty : Dodo$dInternal.$Doc("Text", Data$dString$dCodePoints.toCodePointArray(v._2).length, v._2)
      ])
    );
  }
  if (v.tag === "LetRec") {
    const $0 = v._1;
    return Data$dTuple.$Tuple(
      PrecBlock,
      lines([printLet$p("letrec")(Data$dFunctor.arrayMap(v1 => Data$dTuple.$Tuple(printLocal(Data$dMaybe.$Maybe("Just", v1._1))($0), v1._2._2))(v._2)), v._3._2])
    );
  }
  if (v.tag === "Let") { return Data$dTuple.$Tuple(PrecBlock, lines([printLet$p("let")([Data$dTuple.$Tuple(printLocal(v._1)(v._2), v._3._2)]), v._4._2])); }
  if (v.tag === "EffectBind") { return Data$dTuple.$Tuple(PrecBlock, lines([printLet$p("let!")([Data$dTuple.$Tuple(printLocal(v._1)(v._2), v._3._2)]), v._4._2])); }
  if (v.tag === "EffectPure") { return printUncurriedApp(true)(primOp("effect")("pure"))([v._1]); }
  if (v.tag === "EffectDefer") { return printUncurriedApp(true)(primOp("effect")("defer"))([v._1]); }
  if (v.tag === "Branch") {
    return Data$dTuple.$Tuple(PrecBlock, printBranch(Data$dFunctor.arrayMap(m => PureScript$dBackend$dOptimizer$dSyntax.$Pair(m._1._2, m._2._2))(v._1))(v._2._2));
  }
  if (v.tag === "PrimOp") { return printBackendOperator(v._1); }
  if (v.tag === "PrimEffect") { return printBackendEffect(v._1); }
  if (v.tag === "PrimUndefined") { return primOp("")("undefined"); }
  if (v.tag === "Fail") { return printUncurriedApp(false)(primOp("")("fail"))([]); }
  $runtime.fail();
};
const printBackendExpr = /* #__PURE__ */ PureScript$dBackend$dOptimizer$dSemantics.foldBackendExpr(printBackendSyntax)(rewrite => expr => Data$dTuple.$Tuple(
  expr._1,
  lines([printBackendRewriteCase(rewrite), expr._2])
));
const heading = repeat => hd => words([
  (() => {
    const $0 = power(repeat)(2);
    if ($0 === "") { return Dodo$dInternal.Empty; }
    return Dodo$dInternal.$Doc("Text", Data$dString$dCodePoints.toCodePointArray($0).length, $0);
  })(),
  hd,
  Dodo$dInternal.$Doc(
    "WithPosition",
    v => {
      const $0 = power(repeat)(v.pageWidth - v.column | 0);
      if ($0 === "") { return Dodo$dInternal.Empty; }
      return Dodo$dInternal.$Doc("Text", Data$dString$dCodePoints.toCodePointArray($0).length, $0);
    }
  )
]);
const printSteps = qual => steps => {
  const ident = printQualified(printIdent)(qual);
  return Dodo.foldWithSeparator(Data$dFoldable.foldableArray)(Dodo$dInternal.$Doc("Append", Dodo$dInternal.Break, Dodo$dInternal.Break))(Data$dFunctorWithIndex.mapWithIndexArray(idx => step => lines([
    heading(idx === 0 ? "+" : "-")(words([
      ident,
      Dodo$dInternal.$Doc("Text", 4, "Step"),
      (() => {
        const $0 = Data$dShow.showIntImpl(idx + 1 | 0);
        if ($0 === "") { return Dodo$dInternal.Empty; }
        return Dodo$dInternal.$Doc("Text", Data$dString$dCodePoints.toCodePointArray($0).length, $0);
      })(),
      (() => {
        const v = idx === (steps.length - 1 | 0);
        if (idx === 0) {
          if (v) { return Dodo$dInternal.$Doc("Text", 16, "(Original/Final)"); }
          return Dodo$dInternal.$Doc("Text", 10, "(Original)");
        }
        if (v) { return Dodo$dInternal.$Doc("Text", 7, "(Final)"); }
        return Dodo$dInternal.Empty;
      })()
    ])),
    (() => {
      const $0 = printBackendExpr(step);
      if ($0._2.tag === "Empty") { return Dodo$dInternal.Empty; }
      return Dodo$dInternal.$Doc("Indent", $0._2);
    })()
  ]))(steps));
};
const printModuleSteps = v => steps => fold([
  heading("=")(v === "" ? Dodo$dInternal.Empty : Dodo$dInternal.$Doc("Text", Data$dString$dCodePoints.toCodePointArray(v).length, v)),
  Dodo$dInternal.Break,
  Dodo$dInternal.$Doc(
    "WithPosition",
    v1 => {
      const $0 = power("=")(v1.pageWidth);
      if ($0 === "") { return Dodo$dInternal.Empty; }
      return Dodo$dInternal.$Doc("Text", Data$dString$dCodePoints.toCodePointArray($0).length, $0);
    }
  ),
  Dodo$dInternal.Break,
  Dodo$dInternal.Break,
  Dodo.foldWithSeparator(Data$dFoldable.foldableArray)(Dodo$dInternal.$Doc("Append", Dodo$dInternal.Break, Dodo$dInternal.Break))(Data$dFunctor.arrayMap(v$1 => printSteps(v$1._1)(v$1._2))(steps))
]);
export {
  $Prec,
  PrecAccess,
  PrecApply,
  PrecAtom,
  PrecBlock,
  fold,
  heading,
  lines,
  power,
  primOp,
  printArray,
  printBackendAccessor,
  printBackendEffect,
  printBackendExpr,
  printBackendOperator,
  printBackendOperator1,
  printBackendOperator2,
  printBackendOperatorNum,
  printBackendOperatorOrd,
  printBackendRewriteCase,
  printBackendSyntax,
  printBranch,
  printCurriedAbs,
  printCurriedApp,
  printDistOpCase,
  printIdent,
  printLet$p,
  printLevel,
  printLiteral,
  printLocal,
  printModuleSteps,
  printQualified,
  printRecord,
  printRewrite,
  printSteps,
  printUncurriedAbs,
  printUncurriedApp,
  printUnpackOpCase,
  words,
  wrapPrec
};
