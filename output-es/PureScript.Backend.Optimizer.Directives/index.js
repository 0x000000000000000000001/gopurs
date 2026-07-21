import * as $runtime from "../runtime.js";
import * as Data$dArray from "../Data.Array/index.js";
import * as Data$dEither from "../Data.Either/index.js";
import * as Data$dFoldable from "../Data.Foldable/index.js";
import * as Data$dFoldableWithIndex from "../Data.FoldableWithIndex/index.js";
import * as Data$dMap$dInternal from "../Data.Map.Internal/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dString$dCodeUnits from "../Data.String.CodeUnits/index.js";
import * as Data$dString$dCommon from "../Data.String.Common/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
import * as PureScript$dBackend$dOptimizer$dCoreFn from "../PureScript.Backend.Optimizer.CoreFn/index.js";
import * as PureScript$dBackend$dOptimizer$dSemantics from "../PureScript.Backend.Optimizer.Semantics/index.js";
import * as PureScript$dCST$dErrors from "../PureScript.CST.Errors/index.js";
import * as PureScript$dCST$dLexer from "../PureScript.CST.Lexer/index.js";
import * as PureScript$dCST$dParser$dMonad from "../PureScript.CST.Parser.Monad/index.js";
const expectMap = k => PureScript$dCST$dParser$dMonad.take(tok => {
  const v = k(tok);
  if (v.tag === "Just") { return Data$dEither.$Either("Right", v._1); }
  if (v.tag === "Nothing") { return Data$dEither.$Either("Left", PureScript$dCST$dErrors.$ParseError("UnexpectedToken", tok.value)); }
  $runtime.fail();
});
const keyword = word1 => expectMap(v => {
  if (v.value.tag === "TokLowerName" && v.value._1.tag === "Nothing" && word1 === v.value._2) { return Data$dMaybe.$Maybe("Just", undefined); }
  return Data$dMaybe.Nothing;
});
const label = /* #__PURE__ */ expectMap(v => {
  if (v.value.tag === "TokRawString") { return Data$dMaybe.$Maybe("Just", v.value._1); }
  if (v.value.tag === "TokString") { return Data$dMaybe.$Maybe("Just", v.value._2); }
  if (v.value.tag === "TokLowerName" && v.value._1.tag === "Nothing") { return Data$dMaybe.$Maybe("Just", v.value._2); }
  return Data$dMaybe.Nothing;
});
const natural = /* #__PURE__ */ expectMap(v => {
  if (v.value.tag === "TokInt" && v.value._2.tag === "SmallInt" && v.value._2._1 > 0) { return Data$dMaybe.$Maybe("Just", v.value._2._1); }
  return Data$dMaybe.Nothing;
});
const qualified = /* #__PURE__ */ expectMap(v => {
  if (v.value.tag === "TokLowerName" && v.value._1.tag === "Just") {
    return Data$dMaybe.$Maybe("Just", PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", v.value._1._1), v.value._2));
  }
  return Data$dMaybe.Nothing;
});
const unqualified = /* #__PURE__ */ expectMap(v => {
  if (v.value.tag === "TokLowerName" && v.value._1.tag === "Nothing") { return Data$dMaybe.$Maybe("Just", v.value._2); }
  return Data$dMaybe.Nothing;
});
const equals = /* #__PURE__ */ expectMap(v => {
  if (v.value.tag === "TokEquals") { return Data$dMaybe.$Maybe("Just", undefined); }
  return Data$dMaybe.Nothing;
});
const parseInlineDirective = (state1, more, resume, done) => keyword("default")(
  state1.consumed ? {...state1, consumed: false} : state1,
  more,
  (state3, error) => {
    if (state3.consumed) { return resume(state3, error); }
    return keyword("never")(
      state1.consumed ? {...state1, consumed: false} : state1,
      more,
      (state3$1, error$1) => {
        if (state3$1.consumed) { return resume(state3$1, error$1); }
        return keyword("always")(
          state1.consumed ? {...state1, consumed: false} : state1,
          more,
          (state3$2, error$2) => {
            if (state3$2.consumed) { return resume(state3$2, error$2); }
            return keyword("arity")(
              state1,
              more,
              resume,
              (state2, a) => more(v2 => equals(
                state2,
                more,
                resume,
                (state3$3, a$1) => more(v2$1 => natural(
                  state3$3,
                  more,
                  resume,
                  (state3$4, a$2) => done(state3$4, PureScript$dBackend$dOptimizer$dSemantics.$InlineDirective("InlineArity", a$2))
                ))
              ))
            );
          },
          (state2, a) => done(state2, PureScript$dBackend$dOptimizer$dSemantics.InlineAlways)
        );
      },
      (state2, a) => done(state2, PureScript$dBackend$dOptimizer$dSemantics.InlineNever)
    );
  },
  (state2, a) => done(state2, PureScript$dBackend$dOptimizer$dSemantics.InlineDefault)
);
const dotDot = /* #__PURE__ */ expectMap(v => {
  if (v.value.tag === "TokSymbolName" && v.value._1.tag === "Nothing" && v.value._2 === "..") { return Data$dMaybe.$Maybe("Just", undefined); }
  return Data$dMaybe.Nothing;
});
const dot = /* #__PURE__ */ expectMap(v => {
  if (v.value.tag === "TokDot") { return Data$dMaybe.$Maybe("Just", undefined); }
  return Data$dMaybe.Nothing;
});
const parseInlineAccessor = (state1, more, resume, done) => dot(
  state1.consumed ? {...state1, consumed: false} : state1,
  more,
  (state3, error) => {
    if (state3.consumed) { return resume(state3, error); }
    return dotDot(
      state1.consumed ? {...state1, consumed: false} : state1,
      more,
      (state3$1, error$1) => {
        if (state3$1.consumed) { return resume(state3$1, error$1); }
        return done(state1, PureScript$dBackend$dOptimizer$dSemantics.InlineRef);
      },
      (state2, a) => more(v2 => dot(
        state2,
        more,
        (state3$1, error$1) => {
          if (state3$1.consumed) { return resume(state3$1, error$1); }
          return done(state1, PureScript$dBackend$dOptimizer$dSemantics.InlineRef);
        },
        (state3$1, a$1) => more(v2$1 => label(
          state3$1,
          more,
          (state3$2, error$1) => {
            if (state3$2.consumed) { return resume(state3$2, error$1); }
            return done(state1, PureScript$dBackend$dOptimizer$dSemantics.InlineRef);
          },
          (state3$2, a$2) => done(state3$2, PureScript$dBackend$dOptimizer$dSemantics.$InlineAccessor("InlineSpineProp", a$2))
        ))
      ))
    );
  },
  (state2, a) => more(v2 => label(
    state2,
    more,
    (state3, error) => {
      if (state3.consumed) { return resume(state3, error); }
      return dotDot(
        state1.consumed ? {...state1, consumed: false} : state1,
        more,
        (state3$1, error$1) => {
          if (state3$1.consumed) { return resume(state3$1, error$1); }
          return done(state1, PureScript$dBackend$dOptimizer$dSemantics.InlineRef);
        },
        (state2$1, a$1) => more(v2$1 => dot(
          state2$1,
          more,
          (state3$1, error$1) => {
            if (state3$1.consumed) { return resume(state3$1, error$1); }
            return done(state1, PureScript$dBackend$dOptimizer$dSemantics.InlineRef);
          },
          (state3$1, a$2) => more(v2$2 => label(
            state3$1,
            more,
            (state3$2, error$1) => {
              if (state3$2.consumed) { return resume(state3$2, error$1); }
              return done(state1, PureScript$dBackend$dOptimizer$dSemantics.InlineRef);
            },
            (state3$2, a$3) => done(state3$2, PureScript$dBackend$dOptimizer$dSemantics.$InlineAccessor("InlineSpineProp", a$3))
          ))
        ))
      );
    },
    (state3, a$1) => done(state3, PureScript$dBackend$dOptimizer$dSemantics.$InlineAccessor("InlineProp", a$1))
  ))
);
const parseDirective = (state1, more, resume, done) => qualified(
  state1,
  more,
  resume,
  (state2, a) => more(v2 => parseInlineAccessor(
    state2,
    more,
    resume,
    (state3, a$1) => more(v2$1 => parseInlineDirective(
      state3,
      more,
      resume,
      (state3$1, a$2) => more(v2$2 => PureScript$dCST$dParser$dMonad.eof(
        state3$1,
        more,
        resume,
        (state3$2, a$3) => done(state3$2, Data$dTuple.$Tuple(PureScript$dBackend$dOptimizer$dSemantics.$EvalRef("EvalExtern", a), Data$dTuple.$Tuple(a$1, a$2)))
      ))
    ))
  ))
);
const parseDirectiveMaybe = (state1, more, resume, done) => parseDirective(
  state1.consumed ? {...state1, consumed: false} : state1,
  more,
  (state3, error) => {
    if (state3.consumed) { return resume(state3, error); }
    return PureScript$dCST$dParser$dMonad.eof(state1, more, resume, (state2, a) => done(state2, Data$dMaybe.Nothing));
  },
  (state2, a) => done(state2, Data$dMaybe.$Maybe("Just", a))
);
const parseDirectiveLine = line => {
  const $0 = PureScript$dCST$dParser$dMonad.runParser(PureScript$dCST$dLexer.lex(line))(parseDirectiveMaybe);
  if ($0.tag === "Left") { return Data$dEither.$Either("Left", $0._1); }
  if ($0.tag === "Right") { return Data$dEither.$Either("Right", $0._1._1); }
  $runtime.fail();
};
const parseDirectiveFile = /* #__PURE__ */ (() => {
  const $0 = Data$dFoldableWithIndex.foldableWithIndexArray.foldlWithIndex(line => v => str => {
    const v1 = parseDirectiveLine(str);
    if (v1.tag === "Left") { return {errors: Data$dArray.snoc(v.errors)(Data$dTuple.$Tuple(str, {...v1._1, position: {...v1._1.position, line}})), directives: v.directives}; }
    if (v1.tag === "Right") {
      if (v1._1.tag === "Nothing") { return {errors: v.errors, directives: v.directives}; }
      if (v1._1.tag === "Just") {
        return {errors: v.errors, directives: PureScript$dBackend$dOptimizer$dSemantics.insertDirective(v1._1._1._1)(v1._1._1._2._1)(v1._1._1._2._2)(v.directives)};
      }
    }
    $runtime.fail();
  })({errors: [], directives: Data$dMap$dInternal.Leaf});
  const $1 = Data$dString$dCommon.split("\n");
  return x => $0($1(x));
})();
const parseDirectiveExport = moduleName => (state1, more, resume, done) => keyword("export")(
  state1,
  more,
  resume,
  (state2, a) => more(v2 => unqualified(
    state2,
    more,
    resume,
    (state3, a$1) => more(v2$1 => parseInlineAccessor(
      state3,
      more,
      resume,
      (state3$1, a$2) => more(v2$2 => parseInlineDirective(
        state3$1,
        more,
        resume,
        (state3$2, a$3) => more(v2$3 => PureScript$dCST$dParser$dMonad.eof(
          state3$2,
          more,
          resume,
          (state3$3, a$4) => done(
            state3$3,
            Data$dTuple.$Tuple(
              PureScript$dBackend$dOptimizer$dSemantics.$EvalRef("EvalExtern", PureScript$dBackend$dOptimizer$dCoreFn.$Qualified(Data$dMaybe.$Maybe("Just", moduleName), a$1)),
              Data$dTuple.$Tuple(a$2, a$3)
            )
          )
        ))
      ))
    ))
  ))
);
const parseDirectiveHeader = moduleName => Data$dFoldable.foldlArray(v => {
  const $0 = v.errors;
  const $1 = v.exports;
  const $2 = v.locals;
  return v1 => {
    if (v1.tag === "LineComment") {
      const $3 = Data$dString$dCodeUnits.stripPrefix("@inline")(Data$dString$dCommon.trim(v1._1));
      if ($3.tag === "Just") {
        const line$p = Data$dString$dCommon.trim($3._1);
        const v3 = PureScript$dCST$dParser$dMonad.runParser(PureScript$dCST$dLexer.lex(line$p))((state1, more, resume, done) => parseDirectiveExport(moduleName)(
          state1.consumed ? {...state1, consumed: false} : state1,
          more,
          (state3, error) => {
            if (state3.consumed) { return resume(state3, error); }
            return parseDirective(state1, more, resume, (state2, a) => done(state2, Data$dEither.$Either("Right", a)));
          },
          (state2, a) => done(state2, Data$dEither.$Either("Left", a))
        ));
        if (v3.tag === "Left") { return {errors: Data$dArray.snoc($0)(Data$dTuple.$Tuple(line$p, v3._1)), locals: $2, exports: $1}; }
        if (v3.tag === "Right") {
          if (v3._1._1.tag === "Left") {
            return {errors: $0, locals: $2, exports: PureScript$dBackend$dOptimizer$dSemantics.insertDirective(v3._1._1._1._1)(v3._1._1._1._2._1)(v3._1._1._1._2._2)($1)};
          }
          if (v3._1._1.tag === "Right") {
            return {errors: $0, locals: PureScript$dBackend$dOptimizer$dSemantics.insertDirective(v3._1._1._1._1)(v3._1._1._1._2._1)(v3._1._1._1._2._2)($2), exports: $1};
          }
        }
        $runtime.fail();
      }
    }
    return {errors: $0, locals: $2, exports: $1};
  };
})({errors: [], locals: Data$dMap$dInternal.Leaf, exports: Data$dMap$dInternal.Leaf});
export {
  dot,
  dotDot,
  equals,
  expectMap,
  keyword,
  label,
  natural,
  parseDirective,
  parseDirectiveExport,
  parseDirectiveFile,
  parseDirectiveHeader,
  parseDirectiveLine,
  parseDirectiveMaybe,
  parseInlineAccessor,
  parseInlineDirective,
  qualified,
  unqualified
};
