import * as $runtime from "../runtime.js";
import * as Data$dArray from "../Data.Array/index.js";
import * as Data$dEither from "../Data.Either/index.js";
import * as Data$dFoldable from "../Data.Foldable/index.js";
import * as Data$dLazy from "../Data.Lazy/index.js";
import * as Data$dMap$dInternal from "../Data.Map.Internal/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dOrd from "../Data.Ord/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
import * as PureScript$dCST$dErrors from "../PureScript.CST.Errors/index.js";
import * as PureScript$dCST$dParser$dMonad from "../PureScript.CST.Parser.Monad/index.js";
import * as PureScript$dCST$dTokenStream from "../PureScript.CST.TokenStream/index.js";
import * as PureScript$dCST$dTypes from "../PureScript.CST.Types/index.js";
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
const separated = sepParser => valueParser => {
  const $0 = PureScript$dCST$dParser$dMonad.many((state1, more, resume, done) => sepParser(
    state1,
    more,
    resume,
    (state2, a) => {
      const $0 = Data$dTuple.Tuple(a);
      return more(v2 => valueParser(state2, more, resume, (state3, a$1) => done(state3, $0(a$1))));
    }
  ));
  return (state1, more, resume, done) => valueParser(state1, more, resume, (state2, a) => more(v2 => $0(state2, more, resume, (state3, a$1) => done(state3, {head: a, tail: a$1}))));
};
const reservedKeywords = /* #__PURE__ */ Data$dFoldable.foldlArray(m => a => Data$dMap$dInternal.insert(Data$dOrd.ordString)(a)()(m))(Data$dMap$dInternal.Leaf)([
  "ado",
  "case",
  "class",
  "data",
  "derive",
  "do",
  "else",
  "false",
  "foreign",
  "if",
  "import",
  "in",
  "infix",
  "infixl",
  "infixr",
  "instance",
  "let",
  "module",
  "newtype",
  "of",
  "then",
  "true",
  "type",
  "where"
]);
const recoverTokensWhile = p => initStream => {
  const go = go$a0$copy => {
    let go$a0 = go$a0$copy, go$c = true, go$r;
    while (go$c) {
      const v = go$a0;
      if (v.tag === "Cons") {
        if (v._1._2 === "LytLet" || v._1._2 === "LytLetStmt" || v._1._2 === "LytWhere" || v._1._2 === "LytOf" || v._1._2 === "LytDo" || v._1._2 === "LytAdo") {
          go$c = false;
          go$r = Data$dMaybe.$Maybe("Just", v._1._1);
          continue;
        }
        go$a0 = v._2;
        continue;
      }
      go$c = false;
      go$r = Data$dMaybe.Nothing;
    }
    return go$r;
  };
  const $0 = go(PureScript$dCST$dTokenStream.layoutStack(initStream));
  const indent = (() => {
    if ($0.tag === "Nothing") { return 0; }
    if ($0.tag === "Just") { return $0._1.column; }
    $runtime.fail();
  })();
  const go$1 = go$1$a0$copy => go$1$a1$copy => {
    let go$1$a0 = go$1$a0$copy, go$1$a1 = go$1$a1$copy, go$1$c = true, go$1$r;
    while (go$1$c) {
      const acc = go$1$a0, stream = go$1$a1;
      const v = Data$dLazy.force(stream);
      if (v.tag === "TokenError") {
        go$1$c = false;
        go$1$r = Data$dTuple.$Tuple(acc, stream);
        continue;
      }
      if (v.tag === "TokenEOF") {
        go$1$c = false;
        go$1$r = Data$dTuple.$Tuple(acc, stream);
        continue;
      }
      if (v.tag === "TokenCons") {
        if (p(v._1)(indent)) {
          go$1$a0 = Data$dArray.snoc(acc)(v._1);
          go$1$a1 = v._3;
          continue;
        }
        go$1$c = false;
        go$1$r = Data$dTuple.$Tuple(acc, stream);
        continue;
      }
      $runtime.fail();
    }
    return go$1$r;
  };
  return go$1([])(initStream);
};
const recoverIndent = mkNode => PureScript$dCST$dParser$dMonad.recover(v => {
  const $0 = v.error;
  const $1 = v.position;
  return stream => {
    const v1 = recoverTokensWhile(tok => indent => {
      if (tok.value.tag === "TokLayoutEnd") { return tok.value._1 > indent; }
      if (tok.value.tag === "TokLayoutSep") { return tok.value._1 > indent; }
      return true;
    })(stream);
    if (v1._1.length === 0) { return Data$dMaybe.Nothing; }
    return Data$dMaybe.$Maybe("Just", Data$dTuple.$Tuple(mkNode({position: $1, error: $0, tokens: v1._1}), v1._2));
  };
});
const parseSmallInt = /* #__PURE__ */ PureScript$dCST$dParser$dMonad.take(v => {
  if (v.value.tag === "TokInt") {
    if (v.value._2.tag === "SmallInt") { return Data$dEither.$Either("Right", Data$dTuple.$Tuple(v, v.value._2._1)); }
    return Data$dEither.$Either("Left", PureScript$dCST$dErrors.$ParseError("LexIntOutOfRange", v.value._1));
  }
  return Data$dEither.$Either("Left", PureScript$dCST$dErrors.$ParseError("UnexpectedToken", v.value));
});
const many1 = parser => (state1, more, resume, done) => parser(
  state1,
  more,
  resume,
  (state2, a) => more(v2 => PureScript$dCST$dParser$dMonad.many(parser)(state2, more, resume, (state3, a$1) => done(state3, [a, ...a$1])))
);
const expectMap = k => PureScript$dCST$dParser$dMonad.take(tok => {
  const v = k(tok);
  if (v.tag === "Just") { return Data$dEither.$Either("Right", v._1); }
  if (v.tag === "Nothing") { return Data$dEither.$Either("Left", PureScript$dCST$dErrors.$ParseError("UnexpectedToken", tok.value)); }
  $runtime.fail();
});
const parseBoolean = /* #__PURE__ */ expectMap(v => {
  if (v.value.tag === "TokLowerName" && v.value._1.tag === "Nothing") {
    if (v.value._2 === "true") { return Data$dMaybe.$Maybe("Just", Data$dTuple.$Tuple(v, true)); }
    if (v.value._2 === "false") { return Data$dMaybe.$Maybe("Just", Data$dTuple.$Tuple(v, false)); }
  }
  return Data$dMaybe.Nothing;
});
const parseChar = /* #__PURE__ */ expectMap(v => {
  if (v.value.tag === "TokChar") { return Data$dMaybe.$Maybe("Just", Data$dTuple.$Tuple(v, v.value._2)); }
  return Data$dMaybe.Nothing;
});
const parseHole = /* #__PURE__ */ expectMap(v => {
  if (v.value.tag === "TokHole") { return Data$dMaybe.$Maybe("Just", {token: v, name: v.value._1}); }
  return Data$dMaybe.Nothing;
});
const parseIdent = /* #__PURE__ */ expectMap(v => {
  if (v.value.tag === "TokLowerName" && v.value._1.tag === "Nothing" && !member(v.value._2)(reservedKeywords)) { return Data$dMaybe.$Maybe("Just", {token: v, name: v.value._2}); }
  return Data$dMaybe.Nothing;
});
const parseInt = /* #__PURE__ */ expectMap(v => {
  if (v.value.tag === "TokInt") { return Data$dMaybe.$Maybe("Just", Data$dTuple.$Tuple(v, v.value._2)); }
  return Data$dMaybe.Nothing;
});
const parseLabel = /* #__PURE__ */ expectMap(v => {
  if (v.value.tag === "TokRawString") { return Data$dMaybe.$Maybe("Just", {token: v, name: v.value._1}); }
  if (v.value.tag === "TokString") { return Data$dMaybe.$Maybe("Just", {token: v, name: v.value._2}); }
  if (v.value.tag === "TokLowerName" && v.value._1.tag === "Nothing") { return Data$dMaybe.$Maybe("Just", {token: v, name: v.value._2}); }
  return Data$dMaybe.Nothing;
});
const parseModuleName = /* #__PURE__ */ expectMap(v => {
  if (v.value.tag === "TokUpperName") {
    if (v.value._1.tag === "Just") { return Data$dMaybe.$Maybe("Just", {token: v, name: v.value._1._1 + "." + v.value._2}); }
    if (v.value._1.tag === "Nothing") { return Data$dMaybe.$Maybe("Just", {token: v, name: v.value._2}); }
  }
  return Data$dMaybe.Nothing;
});
const parseNumber = /* #__PURE__ */ expectMap(v => {
  if (v.value.tag === "TokNumber") { return Data$dMaybe.$Maybe("Just", Data$dTuple.$Tuple(v, v.value._2)); }
  return Data$dMaybe.Nothing;
});
const parseOperator = /* #__PURE__ */ expectMap(v => {
  if (v.value.tag === "TokOperator" && v.value._1.tag === "Nothing") { return Data$dMaybe.$Maybe("Just", {token: v, name: v.value._2}); }
  return Data$dMaybe.Nothing;
});
const parseProper = /* #__PURE__ */ expectMap(v => {
  if (v.value.tag === "TokUpperName" && v.value._1.tag === "Nothing") { return Data$dMaybe.$Maybe("Just", {token: v, name: v.value._2}); }
  return Data$dMaybe.Nothing;
});
const parseQualifiedIdent = /* #__PURE__ */ expectMap(v => {
  if (v.value.tag === "TokLowerName" && !member(v.value._2)(reservedKeywords)) { return Data$dMaybe.$Maybe("Just", {token: v, module: v.value._1, name: v.value._2}); }
  return Data$dMaybe.Nothing;
});
const parseQualifiedIdentOrProper = /* #__PURE__ */ expectMap(v => {
  if (v.value.tag === "TokLowerName") { return Data$dMaybe.$Maybe("Just", {token: v, module: v.value._1, name: Data$dEither.$Either("Left", v.value._2)}); }
  if (v.value.tag === "TokUpperName") { return Data$dMaybe.$Maybe("Just", {token: v, module: v.value._1, name: Data$dEither.$Either("Right", v.value._2)}); }
  return Data$dMaybe.Nothing;
});
const parseQualifiedOperator = /* #__PURE__ */ expectMap(v => {
  if (v.value.tag === "TokOperator") { return Data$dMaybe.$Maybe("Just", {token: v, module: v.value._1, name: v.value._2}); }
  return Data$dMaybe.Nothing;
});
const parseQualifiedProper = /* #__PURE__ */ expectMap(v => {
  if (v.value.tag === "TokUpperName") { return Data$dMaybe.$Maybe("Just", {token: v, module: v.value._1, name: v.value._2}); }
  return Data$dMaybe.Nothing;
});
const parseQualifiedSymbol = /* #__PURE__ */ expectMap(v => {
  if (v.value.tag === "TokSymbolName") { return Data$dMaybe.$Maybe("Just", {token: v, module: v.value._1, name: v.value._2}); }
  return Data$dMaybe.Nothing;
});
const parseString = /* #__PURE__ */ expectMap(v => {
  if (v.value.tag === "TokString") { return Data$dMaybe.$Maybe("Just", Data$dTuple.$Tuple(v, v.value._2)); }
  if (v.value.tag === "TokRawString") { return Data$dMaybe.$Maybe("Just", Data$dTuple.$Tuple(v, v.value._1)); }
  return Data$dMaybe.Nothing;
});
const parseSymbol = /* #__PURE__ */ expectMap(v => {
  if (v.value.tag === "TokSymbolName" && v.value._1.tag === "Nothing") { return Data$dMaybe.$Maybe("Just", {token: v, name: v.value._2}); }
  return Data$dMaybe.Nothing;
});
const expect = pred => expectMap(tok => {
  if (pred(tok.value)) { return Data$dMaybe.$Maybe("Just", tok); }
  return Data$dMaybe.Nothing;
});
const tokAt = /* #__PURE__ */ expect(v => v.tag === "TokAt");
const tokBackslash = /* #__PURE__ */ expect(v => v.tag === "TokBackslash");
const tokComma = /* #__PURE__ */ expect(v => v.tag === "TokComma");
const tokDot = /* #__PURE__ */ expect(v => v.tag === "TokDot");
const parseRecordAccessor = expr => (state1, more, resume, done) => tokDot(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => separated(tokDot)(parseLabel)(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => done(state2$1, PureScript$dCST$dTypes.$Expr("ExprRecordAccessor", {expr, dot: a, path: a$1})))
  ))
);
const tokDoubleColon = /* #__PURE__ */ expect(v => v.tag === "TokDoubleColon");
const parseInstanceName = (state1, more, resume, done) => parseIdent(
  state1,
  more,
  resume,
  (state2, a) => {
    const $0 = Data$dTuple.Tuple(a);
    return more(v2 => tokDoubleColon(state2, more, resume, (state3, a$1) => done(state3, $0(a$1))));
  }
);
const tokEquals = /* #__PURE__ */ expect(v => v.tag === "TokEquals");
const tokForall = /* #__PURE__ */ expect(v => v.tag === "TokForall");
const tokKeyOperator = sym => expect(v => v.tag === "TokOperator" && v._1.tag === "Nothing" && sym === v._2);
const parseBinderNegative = (state1, more, resume, done) => tokKeyOperator("-")(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => {
    const $0 = PureScript$dCST$dTypes.BinderInt(Data$dMaybe.$Maybe("Just", a));
    const $1 = PureScript$dCST$dTypes.BinderNumber(Data$dMaybe.$Maybe("Just", a));
    return parseInt(
      state2.consumed ? {...state2, consumed: false} : state2,
      more,
      (state3, error) => {
        if (state3.consumed) { return resume(state3, error); }
        return parseNumber(state2, more, resume, (state2$1, a$1) => done(state2$1, $1(a$1._1)(a$1._2)));
      },
      (state2$1, a$1) => done(state2$1, $0(a$1._1)(a$1._2))
    );
  })
);
const parseRecordLabeled = valueParser => (state1, more, resume, done) => {
  const $0 = (state3, error) => {
    if (state3.consumed) { return resume(state3, error); }
    return parseIdent(state1, more, resume, (state2, a) => done(state2, PureScript$dCST$dTypes.$RecordLabeled("RecordPun", a)));
  };
  const $1 = state1.consumed ? {...state1, consumed: false} : state1;
  return parseLabel(
    $1,
    more,
    (state2, error) => $0({...state2, consumed: $1.consumed}, error),
    (state2, a) => {
      const $2 = Data$dTuple.Tuple(a);
      return more(v2 => tokKeyOperator(":")(
        state2,
        more,
        (state2$1, error) => $0({...state2$1, consumed: $1.consumed}, error),
        (state3, a$1) => {
          const $3 = $2(a$1);
          const $4 = PureScript$dCST$dTypes.RecordField($3._1)($3._2);
          return more(v2$1 => valueParser(state3, more, $0, (state3$1, a$2) => done(state3$1, $4(a$2))));
        }
      ));
    }
  );
};
const parseTypeNegative = (state1, more, resume, done) => tokKeyOperator("-")(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => {
    const $0 = PureScript$dCST$dTypes.TypeInt(Data$dMaybe.$Maybe("Just", a));
    return parseInt(state2, more, resume, (state2$1, a$1) => done(state2$1, $0(a$1._1)(a$1._2)));
  })
);
const tokKeySymbol = sym => expect(v => v.tag === "TokSymbolName" && v._1.tag === "Nothing" && sym === v._2);
const tokKeyword = kw => expect(v => v.tag === "TokLowerName" && v._1.tag === "Nothing" && kw === v._2);
const parseFixityKeyword = (state1, more, resume, done) => tokKeyword("infix")(
  state1.consumed ? {...state1, consumed: false} : state1,
  more,
  (state3, error) => {
    if (state3.consumed) { return resume(state3, error); }
    return tokKeyword("infixl")(
      state1.consumed ? {...state1, consumed: false} : state1,
      more,
      (state3$1, error$1) => {
        if (state3$1.consumed) { return resume(state3$1, error$1); }
        return tokKeyword("infixr")(state1, more, resume, (state2, a) => done(state2, Data$dTuple.$Tuple(a, PureScript$dCST$dTypes.Infixr)));
      },
      (state2, a) => done(state2, Data$dTuple.$Tuple(a, PureScript$dCST$dTypes.Infixl))
    );
  },
  (state2, a) => done(state2, Data$dTuple.$Tuple(a, PureScript$dCST$dTypes.Infix))
);
const parseFixityOp = (state1, more, resume, done) => {
  const $0 = (state3, error) => {
    if (state3.consumed) { return resume(state3, error); }
    return parseQualifiedIdentOrProper(
      state1,
      more,
      resume,
      (state2, a) => {
        const $0 = PureScript$dCST$dTypes.FixityValue(a);
        return more(v2 => tokKeyword("as")(
          state2,
          more,
          resume,
          (state3$1, a$1) => {
            const $1 = $0(a$1);
            return more(v2$1 => parseOperator(state3$1, more, resume, (state3$2, a$2) => done(state3$2, $1(a$2))));
          }
        ));
      }
    );
  };
  return tokKeyword("type")(
    state1.consumed ? {...state1, consumed: false} : state1,
    more,
    $0,
    (state2, a) => {
      const $1 = PureScript$dCST$dTypes.FixityType(a);
      return more(v2 => parseQualifiedProper(
        state2,
        more,
        $0,
        (state3, a$1) => {
          const $2 = $1(a$1);
          return more(v2$1 => tokKeyword("as")(
            state3,
            more,
            $0,
            (state3$1, a$2) => {
              const $3 = $2(a$2);
              return more(v2$2 => parseOperator(state3$1, more, $0, (state3$2, a$3) => done(state3$2, $3(a$3))));
            }
          ));
        }
      ));
    }
  );
};
const parseDeclFixity = (state1, more, resume, done) => parseFixityKeyword(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => parseSmallInt(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => parseFixityOp(
      state2$1,
      more,
      resume,
      (state2$2, a$2) => more(v1$2 => done(state2$2, PureScript$dCST$dTypes.$Declaration("DeclFixity", {keyword: a, prec: a$1, operator: a$2})))
    ))
  ))
);
const parseForeignKind = (state1, more, resume, done) => tokKeyword("kind")(
  state1,
  more,
  (state2, error) => resume({...state2, consumed: state1.consumed}, error),
  (state2, a) => {
    const $0 = PureScript$dCST$dTypes.ForeignKind(a);
    return more(v2 => parseProper(state2, more, (state2$1, error) => resume({...state2$1, consumed: state1.consumed}, error), (state3, a$1) => done(state3, $0(a$1))));
  }
);
const parseRole = (state1, more, resume, done) => tokKeyword("representational")(
  state1.consumed ? {...state1, consumed: false} : state1,
  more,
  (state3, error) => {
    if (state3.consumed) { return resume(state3, error); }
    return tokKeyword("nominal")(
      state1.consumed ? {...state1, consumed: false} : state1,
      more,
      (state3$1, error$1) => {
        if (state3$1.consumed) { return resume(state3$1, error$1); }
        return tokKeyword("phantom")(state1, more, resume, (state2, a) => done(state2, Data$dTuple.$Tuple(a, PureScript$dCST$dTypes.Phantom)));
      },
      (state2, a) => done(state2, Data$dTuple.$Tuple(a, PureScript$dCST$dTypes.Nominal))
    );
  },
  (state2, a) => done(state2, Data$dTuple.$Tuple(a, PureScript$dCST$dTypes.Representational))
);
const parseDeclRole = keyword1 => (state1, more, resume, done) => tokKeyword("role")(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => parseProper(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => many1(parseRole)(
      state2$1,
      more,
      resume,
      (state2$2, a$2) => more(v1$2 => done(state2$2, PureScript$dCST$dTypes.$Declaration("DeclRole", keyword1, a, a$1, a$2)))
    ))
  ))
);
const tokLayoutEnd = /* #__PURE__ */ expect(v => v.tag === "TokLayoutEnd");
const tokLayoutSep = /* #__PURE__ */ expect(v => v.tag === "TokLayoutSep");
const parseInstanceChainSeparator = (state1, more, resume, done) => tokKeyword("else")(
  state1,
  more,
  resume,
  (state2, a) => more(v2 => tokLayoutSep(
    state2.consumed ? {...state2, consumed: false} : state2,
    more,
    (state3, error) => {
      if (state3.consumed) { return resume(state3, error); }
      return done(state2, a);
    },
    (state2$1, a$1) => done(state2$1, a)
  ))
);
const tokLayoutStart = /* #__PURE__ */ expect(v => v.tag === "TokLayoutStart");
const layoutNonEmpty = valueParser => {
  const $0 = PureScript$dCST$dParser$dMonad.many((state1, more, resume, done) => tokLayoutSep(
    state1,
    more,
    resume,
    (state2, a) => more(v2 => valueParser(state2, more, resume, (state3, a$1) => done(state3, a$1)))
  ));
  return (state1, more, resume, done) => tokLayoutStart(
    state1,
    more,
    resume,
    (state2, a) => more(v2 => valueParser(
      state2,
      more,
      resume,
      (state3, a$1) => more(v2$1 => $0(
        state3,
        more,
        resume,
        (state2$1, a$2) => more(v2$2 => tokLayoutEnd(state2$1, more, resume, (state3$1, a$3) => done(state3$1, [a$1, ...a$2])))
      ))
    ))
  );
};
const tokLeftArrow = /* #__PURE__ */ expect(v => v.tag === "TokLeftArrow");
const tokLeftBrace = /* #__PURE__ */ expect(v => v.tag === "TokLeftBrace");
const tokLeftFatArrow = /* #__PURE__ */ expect(v => v.tag === "TokOperator" && v._1.tag === "Nothing" && (v._2 === "<=" || v._2 === "⇐"));
const tokLeftParen = /* #__PURE__ */ expect(v => v.tag === "TokLeftParen");
const tokLeftSquare = /* #__PURE__ */ expect(v => v.tag === "TokLeftSquare");
const tokPipe = /* #__PURE__ */ expect(v => v.tag === "TokPipe");
const tokQualifiedKeyword = kw => expect(v => v.tag === "TokLowerName" && kw === v._2);
const tokRightArrow = /* #__PURE__ */ expect(v => v.tag === "TokRightArrow");
const parseFundep = (state1, more, resume, done) => tokRightArrow(
  state1.consumed ? {...state1, consumed: false} : state1,
  more,
  (state3, error) => {
    if (state3.consumed) { return resume(state3, error); }
    return many1(parseIdent)(
      state1,
      more,
      resume,
      (state2, a) => {
        const $0 = PureScript$dCST$dTypes.FundepDetermines(a);
        return more(v2 => tokRightArrow(
          state2,
          more,
          resume,
          (state3$1, a$1) => {
            const $1 = $0(a$1);
            return more(v2$1 => many1(parseIdent)(state3$1, more, resume, (state3$2, a$2) => done(state3$2, $1(a$2))));
          }
        ));
      }
    );
  },
  (state2, a) => {
    const $0 = PureScript$dCST$dTypes.FundepDetermined(a);
    return more(v2 => many1(parseIdent)(
      state2,
      more,
      (state3, error) => {
        if (state3.consumed) { return resume(state3, error); }
        return many1(parseIdent)(
          state1,
          more,
          resume,
          (state2$1, a$1) => {
            const $1 = PureScript$dCST$dTypes.FundepDetermines(a$1);
            return more(v2$1 => tokRightArrow(
              state2$1,
              more,
              resume,
              (state3$1, a$2) => {
                const $2 = $1(a$2);
                return more(v2$2 => many1(parseIdent)(state3$1, more, resume, (state3$2, a$3) => done(state3$2, $2(a$3))));
              }
            ));
          }
        );
      },
      (state3, a$1) => done(state3, $0(a$1))
    ));
  }
);
const tokRightBrace = /* #__PURE__ */ expect(v => v.tag === "TokRightBrace");
const tokRightFatArrow = /* #__PURE__ */ expect(v => v.tag === "TokRightFatArrow");
const tokRightParen = /* #__PURE__ */ expect(v => v.tag === "TokRightParen");
const parens = valueParser => (state1, more, resume, done) => tokLeftParen(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => valueParser(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => tokRightParen(state2$1, more, resume, (state2$2, a$2) => more(v1$2 => done(state2$2, {open: a, value: a$1, close: a$2}))))
  ))
);
const parseEmptyRow = open => (state1, more, resume, done) => tokRightParen(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => done(state2, PureScript$dCST$dTypes.$Type("TypeRow", {open, value: {labels: Data$dMaybe.Nothing, tail: Data$dMaybe.Nothing}, close: a})))
);
const tokRightSquare = /* #__PURE__ */ expect(v => v.tag === "TokRightSquare");
const tokSymbolArrow = /* #__PURE__ */ expect(v => v.tag === "TokSymbolArrow");
const tokTick = /* #__PURE__ */ expect(v => v.tag === "TokTick");
const tokUnderscore = /* #__PURE__ */ expect(v => v.tag === "TokUnderscore");
const delimited = openTok => closeTok => sepTok => valueParser => (state1, more, resume, done) => openTok(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => {
    const $0 = separated(sepTok)(valueParser);
    return closeTok(
      state2.consumed ? {...state2, consumed: false} : state2,
      more,
      (state3, error) => {
        if (state3.consumed) { return resume(state3, error); }
        return $0(
          state2,
          more,
          resume,
          (state2$1, a$1) => more(v2 => closeTok(state2$1, more, resume, (state3$1, a$2) => done(state3$1, {open: a, value: Data$dMaybe.$Maybe("Just", a$1), close: a$2})))
        );
      },
      (state2$1, a$1) => done(state2$1, {open: a, value: Data$dMaybe.Nothing, close: a$1})
    );
  })
);
const parseDataMembers = (state1, more, resume, done) => tokKeySymbol("..")(
  state1.consumed ? {...state1, consumed: false} : state1,
  more,
  (state3, error) => {
    if (state3.consumed) { return resume(state3, error); }
    return delimited(tokLeftParen)(tokRightParen)(tokComma)(parseProper)(
      state1,
      more,
      resume,
      (state2, a) => done(state2, PureScript$dCST$dTypes.$DataMembers("DataEnumerated", a))
    );
  },
  (state2, a) => done(state2, PureScript$dCST$dTypes.$DataMembers("DataAll", a))
);
const parseExport = (state1, more, resume, done) => tokKeyword("type")(
  state1.consumed ? {...state1, consumed: false} : state1,
  more,
  (state3, error) => {
    if (state3.consumed) { return resume(state3, error); }
    return tokKeyword("class")(
      state1.consumed ? {...state1, consumed: false} : state1,
      more,
      (state3$1, error$1) => {
        if (state3$1.consumed) { return resume(state3$1, error$1); }
        return tokKeyword("module")(
          state1.consumed ? {...state1, consumed: false} : state1,
          more,
          (state3$2, error$2) => {
            if (state3$2.consumed) { return resume(state3$2, error$2); }
            return parseSymbol(
              state1.consumed ? {...state1, consumed: false} : state1,
              more,
              (state3$3, error$3) => {
                if (state3$3.consumed) { return resume(state3$3, error$3); }
                return parseIdent(
                  state1.consumed ? {...state1, consumed: false} : state1,
                  more,
                  (state3$4, error$4) => {
                    if (state3$4.consumed) { return resume(state3$4, error$4); }
                    return parseProper(
                      state1,
                      more,
                      resume,
                      (state2, a) => {
                        const $0 = PureScript$dCST$dTypes.ExportType(a);
                        return more(v2 => parseDataMembers(
                          state2.consumed ? {...state2, consumed: false} : state2,
                          more,
                          (state3$5, error$5) => {
                            if (state3$5.consumed) { return resume(state3$5, error$5); }
                            return done(state2, $0(Data$dMaybe.Nothing));
                          },
                          (state2$1, a$1) => done(state2$1, $0(Data$dMaybe.$Maybe("Just", a$1)))
                        ));
                      }
                    );
                  },
                  (state2, a) => done(state2, PureScript$dCST$dTypes.$Export("ExportValue", a))
                );
              },
              (state2, a) => done(state2, PureScript$dCST$dTypes.$Export("ExportOp", a))
            );
          },
          (state2, a) => {
            const $0 = PureScript$dCST$dTypes.ExportModule(a);
            return more(v2 => parseModuleName(
              state2,
              more,
              (state3$2, error$2) => {
                if (state3$2.consumed) { return resume(state3$2, error$2); }
                return parseSymbol(
                  state1.consumed ? {...state1, consumed: false} : state1,
                  more,
                  (state3$3, error$3) => {
                    if (state3$3.consumed) { return resume(state3$3, error$3); }
                    return parseIdent(
                      state1.consumed ? {...state1, consumed: false} : state1,
                      more,
                      (state3$4, error$4) => {
                        if (state3$4.consumed) { return resume(state3$4, error$4); }
                        return parseProper(
                          state1,
                          more,
                          resume,
                          (state2$1, a$1) => {
                            const $1 = PureScript$dCST$dTypes.ExportType(a$1);
                            return more(v2$1 => parseDataMembers(
                              state2$1.consumed ? {...state2$1, consumed: false} : state2$1,
                              more,
                              (state3$5, error$5) => {
                                if (state3$5.consumed) { return resume(state3$5, error$5); }
                                return done(state2$1, $1(Data$dMaybe.Nothing));
                              },
                              (state2$2, a$2) => done(state2$2, $1(Data$dMaybe.$Maybe("Just", a$2)))
                            ));
                          }
                        );
                      },
                      (state2$1, a$1) => done(state2$1, PureScript$dCST$dTypes.$Export("ExportValue", a$1))
                    );
                  },
                  (state2$1, a$1) => done(state2$1, PureScript$dCST$dTypes.$Export("ExportOp", a$1))
                );
              },
              (state3$2, a$1) => done(state3$2, $0(a$1))
            ));
          }
        );
      },
      (state2, a) => {
        const $0 = PureScript$dCST$dTypes.ExportClass(a);
        return more(v2 => parseProper(
          state2,
          more,
          (state3$1, error$1) => {
            if (state3$1.consumed) { return resume(state3$1, error$1); }
            return tokKeyword("module")(
              state1.consumed ? {...state1, consumed: false} : state1,
              more,
              (state3$2, error$2) => {
                if (state3$2.consumed) { return resume(state3$2, error$2); }
                return parseSymbol(
                  state1.consumed ? {...state1, consumed: false} : state1,
                  more,
                  (state3$3, error$3) => {
                    if (state3$3.consumed) { return resume(state3$3, error$3); }
                    return parseIdent(
                      state1.consumed ? {...state1, consumed: false} : state1,
                      more,
                      (state3$4, error$4) => {
                        if (state3$4.consumed) { return resume(state3$4, error$4); }
                        return parseProper(
                          state1,
                          more,
                          resume,
                          (state2$1, a$1) => {
                            const $1 = PureScript$dCST$dTypes.ExportType(a$1);
                            return more(v2$1 => parseDataMembers(
                              state2$1.consumed ? {...state2$1, consumed: false} : state2$1,
                              more,
                              (state3$5, error$5) => {
                                if (state3$5.consumed) { return resume(state3$5, error$5); }
                                return done(state2$1, $1(Data$dMaybe.Nothing));
                              },
                              (state2$2, a$2) => done(state2$2, $1(Data$dMaybe.$Maybe("Just", a$2)))
                            ));
                          }
                        );
                      },
                      (state2$1, a$1) => done(state2$1, PureScript$dCST$dTypes.$Export("ExportValue", a$1))
                    );
                  },
                  (state2$1, a$1) => done(state2$1, PureScript$dCST$dTypes.$Export("ExportOp", a$1))
                );
              },
              (state2$1, a$1) => {
                const $1 = PureScript$dCST$dTypes.ExportModule(a$1);
                return more(v2$1 => parseModuleName(
                  state2$1,
                  more,
                  (state3$2, error$2) => {
                    if (state3$2.consumed) { return resume(state3$2, error$2); }
                    return parseSymbol(
                      state1.consumed ? {...state1, consumed: false} : state1,
                      more,
                      (state3$3, error$3) => {
                        if (state3$3.consumed) { return resume(state3$3, error$3); }
                        return parseIdent(
                          state1.consumed ? {...state1, consumed: false} : state1,
                          more,
                          (state3$4, error$4) => {
                            if (state3$4.consumed) { return resume(state3$4, error$4); }
                            return parseProper(
                              state1,
                              more,
                              resume,
                              (state2$2, a$2) => {
                                const $2 = PureScript$dCST$dTypes.ExportType(a$2);
                                return more(v2$2 => parseDataMembers(
                                  state2$2.consumed ? {...state2$2, consumed: false} : state2$2,
                                  more,
                                  (state3$5, error$5) => {
                                    if (state3$5.consumed) { return resume(state3$5, error$5); }
                                    return done(state2$2, $2(Data$dMaybe.Nothing));
                                  },
                                  (state2$3, a$3) => done(state2$3, $2(Data$dMaybe.$Maybe("Just", a$3)))
                                ));
                              }
                            );
                          },
                          (state2$2, a$2) => done(state2$2, PureScript$dCST$dTypes.$Export("ExportValue", a$2))
                        );
                      },
                      (state2$2, a$2) => done(state2$2, PureScript$dCST$dTypes.$Export("ExportOp", a$2))
                    );
                  },
                  (state3$2, a$2) => done(state3$2, $1(a$2))
                ));
              }
            );
          },
          (state3$1, a$1) => done(state3$1, $0(a$1))
        ));
      }
    );
  },
  (state2, a) => {
    const $0 = PureScript$dCST$dTypes.ExportTypeOp(a);
    return more(v2 => parseSymbol(
      state2,
      more,
      (state3, error) => {
        if (state3.consumed) { return resume(state3, error); }
        return tokKeyword("class")(
          state1.consumed ? {...state1, consumed: false} : state1,
          more,
          (state3$1, error$1) => {
            if (state3$1.consumed) { return resume(state3$1, error$1); }
            return tokKeyword("module")(
              state1.consumed ? {...state1, consumed: false} : state1,
              more,
              (state3$2, error$2) => {
                if (state3$2.consumed) { return resume(state3$2, error$2); }
                return parseSymbol(
                  state1.consumed ? {...state1, consumed: false} : state1,
                  more,
                  (state3$3, error$3) => {
                    if (state3$3.consumed) { return resume(state3$3, error$3); }
                    return parseIdent(
                      state1.consumed ? {...state1, consumed: false} : state1,
                      more,
                      (state3$4, error$4) => {
                        if (state3$4.consumed) { return resume(state3$4, error$4); }
                        return parseProper(
                          state1,
                          more,
                          resume,
                          (state2$1, a$1) => {
                            const $1 = PureScript$dCST$dTypes.ExportType(a$1);
                            return more(v2$1 => parseDataMembers(
                              state2$1.consumed ? {...state2$1, consumed: false} : state2$1,
                              more,
                              (state3$5, error$5) => {
                                if (state3$5.consumed) { return resume(state3$5, error$5); }
                                return done(state2$1, $1(Data$dMaybe.Nothing));
                              },
                              (state2$2, a$2) => done(state2$2, $1(Data$dMaybe.$Maybe("Just", a$2)))
                            ));
                          }
                        );
                      },
                      (state2$1, a$1) => done(state2$1, PureScript$dCST$dTypes.$Export("ExportValue", a$1))
                    );
                  },
                  (state2$1, a$1) => done(state2$1, PureScript$dCST$dTypes.$Export("ExportOp", a$1))
                );
              },
              (state2$1, a$1) => {
                const $1 = PureScript$dCST$dTypes.ExportModule(a$1);
                return more(v2$1 => parseModuleName(
                  state2$1,
                  more,
                  (state3$2, error$2) => {
                    if (state3$2.consumed) { return resume(state3$2, error$2); }
                    return parseSymbol(
                      state1.consumed ? {...state1, consumed: false} : state1,
                      more,
                      (state3$3, error$3) => {
                        if (state3$3.consumed) { return resume(state3$3, error$3); }
                        return parseIdent(
                          state1.consumed ? {...state1, consumed: false} : state1,
                          more,
                          (state3$4, error$4) => {
                            if (state3$4.consumed) { return resume(state3$4, error$4); }
                            return parseProper(
                              state1,
                              more,
                              resume,
                              (state2$2, a$2) => {
                                const $2 = PureScript$dCST$dTypes.ExportType(a$2);
                                return more(v2$2 => parseDataMembers(
                                  state2$2.consumed ? {...state2$2, consumed: false} : state2$2,
                                  more,
                                  (state3$5, error$5) => {
                                    if (state3$5.consumed) { return resume(state3$5, error$5); }
                                    return done(state2$2, $2(Data$dMaybe.Nothing));
                                  },
                                  (state2$3, a$3) => done(state2$3, $2(Data$dMaybe.$Maybe("Just", a$3)))
                                ));
                              }
                            );
                          },
                          (state2$2, a$2) => done(state2$2, PureScript$dCST$dTypes.$Export("ExportValue", a$2))
                        );
                      },
                      (state2$2, a$2) => done(state2$2, PureScript$dCST$dTypes.$Export("ExportOp", a$2))
                    );
                  },
                  (state3$2, a$2) => done(state3$2, $1(a$2))
                ));
              }
            );
          },
          (state2$1, a$1) => {
            const $1 = PureScript$dCST$dTypes.ExportClass(a$1);
            return more(v2$1 => parseProper(
              state2$1,
              more,
              (state3$1, error$1) => {
                if (state3$1.consumed) { return resume(state3$1, error$1); }
                return tokKeyword("module")(
                  state1.consumed ? {...state1, consumed: false} : state1,
                  more,
                  (state3$2, error$2) => {
                    if (state3$2.consumed) { return resume(state3$2, error$2); }
                    return parseSymbol(
                      state1.consumed ? {...state1, consumed: false} : state1,
                      more,
                      (state3$3, error$3) => {
                        if (state3$3.consumed) { return resume(state3$3, error$3); }
                        return parseIdent(
                          state1.consumed ? {...state1, consumed: false} : state1,
                          more,
                          (state3$4, error$4) => {
                            if (state3$4.consumed) { return resume(state3$4, error$4); }
                            return parseProper(
                              state1,
                              more,
                              resume,
                              (state2$2, a$2) => {
                                const $2 = PureScript$dCST$dTypes.ExportType(a$2);
                                return more(v2$2 => parseDataMembers(
                                  state2$2.consumed ? {...state2$2, consumed: false} : state2$2,
                                  more,
                                  (state3$5, error$5) => {
                                    if (state3$5.consumed) { return resume(state3$5, error$5); }
                                    return done(state2$2, $2(Data$dMaybe.Nothing));
                                  },
                                  (state2$3, a$3) => done(state2$3, $2(Data$dMaybe.$Maybe("Just", a$3)))
                                ));
                              }
                            );
                          },
                          (state2$2, a$2) => done(state2$2, PureScript$dCST$dTypes.$Export("ExportValue", a$2))
                        );
                      },
                      (state2$2, a$2) => done(state2$2, PureScript$dCST$dTypes.$Export("ExportOp", a$2))
                    );
                  },
                  (state2$2, a$2) => {
                    const $2 = PureScript$dCST$dTypes.ExportModule(a$2);
                    return more(v2$2 => parseModuleName(
                      state2$2,
                      more,
                      (state3$2, error$2) => {
                        if (state3$2.consumed) { return resume(state3$2, error$2); }
                        return parseSymbol(
                          state1.consumed ? {...state1, consumed: false} : state1,
                          more,
                          (state3$3, error$3) => {
                            if (state3$3.consumed) { return resume(state3$3, error$3); }
                            return parseIdent(
                              state1.consumed ? {...state1, consumed: false} : state1,
                              more,
                              (state3$4, error$4) => {
                                if (state3$4.consumed) { return resume(state3$4, error$4); }
                                return parseProper(
                                  state1,
                                  more,
                                  resume,
                                  (state2$3, a$3) => {
                                    const $3 = PureScript$dCST$dTypes.ExportType(a$3);
                                    return more(v2$3 => parseDataMembers(
                                      state2$3.consumed ? {...state2$3, consumed: false} : state2$3,
                                      more,
                                      (state3$5, error$5) => {
                                        if (state3$5.consumed) { return resume(state3$5, error$5); }
                                        return done(state2$3, $3(Data$dMaybe.Nothing));
                                      },
                                      (state2$4, a$4) => done(state2$4, $3(Data$dMaybe.$Maybe("Just", a$4)))
                                    ));
                                  }
                                );
                              },
                              (state2$3, a$3) => done(state2$3, PureScript$dCST$dTypes.$Export("ExportValue", a$3))
                            );
                          },
                          (state2$3, a$3) => done(state2$3, PureScript$dCST$dTypes.$Export("ExportOp", a$3))
                        );
                      },
                      (state3$2, a$3) => done(state3$2, $2(a$3))
                    ));
                  }
                );
              },
              (state3$1, a$2) => done(state3$1, $1(a$2))
            ));
          }
        );
      },
      (state3, a$1) => done(state3, $0(a$1))
    ));
  }
);
const parseImport = (state1, more, resume, done) => parseSymbol(
  state1.consumed ? {...state1, consumed: false} : state1,
  more,
  (state3, error) => {
    if (state3.consumed) { return resume(state3, error); }
    return parseProper(
      state1.consumed ? {...state1, consumed: false} : state1,
      more,
      (state3$1, error$1) => {
        if (state3$1.consumed) { return resume(state3$1, error$1); }
        return tokKeyword("type")(
          state1.consumed ? {...state1, consumed: false} : state1,
          more,
          (state3$2, error$2) => {
            if (state3$2.consumed) { return resume(state3$2, error$2); }
            const $0 = (state3$3, error$3) => {
              if (state3$3.consumed) { return resume(state3$3, error$3); }
              return parseIdent(state1, more, resume, (state2, a) => done(state2, PureScript$dCST$dTypes.$Import("ImportValue", a)));
            };
            return tokKeyword("class")(
              state1.consumed ? {...state1, consumed: false} : state1,
              more,
              $0,
              (state2, a) => {
                const $1 = PureScript$dCST$dTypes.ImportClass(a);
                return more(v2 => parseProper(state2, more, $0, (state3$3, a$1) => done(state3$3, $1(a$1))));
              }
            );
          },
          (state2, a) => {
            const $0 = PureScript$dCST$dTypes.ImportTypeOp(a);
            return more(v2 => parseSymbol(
              state2,
              more,
              (state3$2, error$2) => {
                if (state3$2.consumed) { return resume(state3$2, error$2); }
                const $1 = (state3$3, error$3) => {
                  if (state3$3.consumed) { return resume(state3$3, error$3); }
                  return parseIdent(state1, more, resume, (state2$1, a$1) => done(state2$1, PureScript$dCST$dTypes.$Import("ImportValue", a$1)));
                };
                return tokKeyword("class")(
                  state1.consumed ? {...state1, consumed: false} : state1,
                  more,
                  $1,
                  (state2$1, a$1) => {
                    const $2 = PureScript$dCST$dTypes.ImportClass(a$1);
                    return more(v2$1 => parseProper(state2$1, more, $1, (state3$3, a$2) => done(state3$3, $2(a$2))));
                  }
                );
              },
              (state3$2, a$1) => done(state3$2, $0(a$1))
            ));
          }
        );
      },
      (state2, a) => {
        const $0 = PureScript$dCST$dTypes.ImportType(a);
        return more(v2 => parseDataMembers(
          state2.consumed ? {...state2, consumed: false} : state2,
          more,
          (state3$1, error$1) => {
            if (state3$1.consumed) {
              if (state3$1.consumed) { return resume(state3$1, error$1); }
              return tokKeyword("type")(
                state1.consumed ? {...state1, consumed: false} : state1,
                more,
                (state3$2, error$2) => {
                  if (state3$2.consumed) { return resume(state3$2, error$2); }
                  const $1 = (state3$3, error$3) => {
                    if (state3$3.consumed) { return resume(state3$3, error$3); }
                    return parseIdent(state1, more, resume, (state2$1, a$1) => done(state2$1, PureScript$dCST$dTypes.$Import("ImportValue", a$1)));
                  };
                  return tokKeyword("class")(
                    state1.consumed ? {...state1, consumed: false} : state1,
                    more,
                    $1,
                    (state2$1, a$1) => {
                      const $2 = PureScript$dCST$dTypes.ImportClass(a$1);
                      return more(v2$1 => parseProper(state2$1, more, $1, (state3$3, a$2) => done(state3$3, $2(a$2))));
                    }
                  );
                },
                (state2$1, a$1) => {
                  const $1 = PureScript$dCST$dTypes.ImportTypeOp(a$1);
                  return more(v2$1 => parseSymbol(
                    state2$1,
                    more,
                    (state3$2, error$2) => {
                      if (state3$2.consumed) { return resume(state3$2, error$2); }
                      const $2 = (state3$3, error$3) => {
                        if (state3$3.consumed) { return resume(state3$3, error$3); }
                        return parseIdent(state1, more, resume, (state2$2, a$2) => done(state2$2, PureScript$dCST$dTypes.$Import("ImportValue", a$2)));
                      };
                      return tokKeyword("class")(
                        state1.consumed ? {...state1, consumed: false} : state1,
                        more,
                        $2,
                        (state2$2, a$2) => {
                          const $3 = PureScript$dCST$dTypes.ImportClass(a$2);
                          return more(v2$2 => parseProper(state2$2, more, $2, (state3$3, a$3) => done(state3$3, $3(a$3))));
                        }
                      );
                    },
                    (state3$2, a$2) => done(state3$2, $1(a$2))
                  ));
                }
              );
            }
            return done(state2, $0(Data$dMaybe.Nothing));
          },
          (state2$1, a$1) => done(state2$1, $0(Data$dMaybe.$Maybe("Just", a$1)))
        ));
      }
    );
  },
  (state2, a) => done(state2, PureScript$dCST$dTypes.$Import("ImportOp", a))
);
const parseImportDecl = (state1, more, resume, done) => tokKeyword("import")(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => parseModuleName(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => {
      const $0 = parens(separated(tokComma)(parseImport));
      const $1 = (state2$2, a$2) => more(v1$2 => tokKeyword("as")(
        state2$2.consumed ? {...state2$2, consumed: false} : state2$2,
        more,
        (state3, error) => {
          if (state3.consumed) { return resume(state3, error); }
          return more(v1$3 => done(state2$2, {keyword: a, module: a$1, names: a$2, qualified: Data$dMaybe.Nothing}));
        },
        (state2$3, a$3) => {
          const $1 = Data$dTuple.Tuple(a$3);
          return more(v2 => parseModuleName(
            state2$3,
            more,
            (state3, error) => {
              if (state3.consumed) { return resume(state3, error); }
              return more(v1$3 => done(state2$2, {keyword: a, module: a$1, names: a$2, qualified: Data$dMaybe.Nothing}));
            },
            (state3, a$4) => more(v1$3 => done(state3, {keyword: a, module: a$1, names: a$2, qualified: Data$dMaybe.$Maybe("Just", $1(a$4))}))
          ));
        }
      ));
      const $2 = state2$1.consumed ? {...state2$1, consumed: false} : state2$1;
      const $3 = (state2$2, f) => more(v2 => $0(
        state2$2,
        more,
        (state3, error) => {
          if (state3.consumed) { return resume(state3, error); }
          return $1(state2$1, Data$dMaybe.Nothing);
        },
        (state3, a$2) => $1(state3, Data$dMaybe.$Maybe("Just", f(a$2)))
      ));
      return tokKeyword("hiding")(
        $2.consumed ? {...$2, consumed: false} : $2,
        more,
        (state3, error) => {
          if (state3.consumed) {
            if (state3.consumed) { return resume(state3, error); }
            return $1(state2$1, Data$dMaybe.Nothing);
          }
          return $3($2, Data$dTuple.Tuple(Data$dMaybe.Nothing));
        },
        (state2$2, a$2) => $3(state2$2, Data$dTuple.Tuple(Data$dMaybe.$Maybe("Just", a$2)))
      );
    })
  ))
);
const parseModuleImportDecls = /* #__PURE__ */ PureScript$dCST$dParser$dMonad.many((state1, more, resume, done) => parseImportDecl(
  state1,
  more,
  resume,
  (state2, a) => more(v2 => tokLayoutSep(
    state2.consumed ? {...state2, consumed: false} : state2,
    more,
    (state3, error) => {
      if (state3.consumed) { return resume(state3, error); }
      return tokLayoutEnd(state2, more, (v1, error$1) => resume(state2, error$1), (v1, value) => done(state2, a));
    },
    (state3, a$1) => done(state3, a)
  ))
));
const parseModuleHeader = (state1, more, resume, done) => tokKeyword("module")(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => parseModuleName(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => {
      const $0 = exports => (state1$1, more$1, resume$1, done$1) => tokKeyword("where")(
        state1$1,
        more$1,
        resume$1,
        (state2$2, a$2) => more$1(v1$2 => tokLayoutStart(
          state2$2,
          more$1,
          resume$1,
          (state2$3, a$3) => more$1(v2 => parseModuleImportDecls(
            state2$3,
            more$1,
            resume$1,
            (state3, a$4) => more$1(v1$3 => done$1(state3, {keyword: a, name: a$1, exports, where: a$2, imports: a$4}))
          ))
        ))
      );
      return parens(separated(tokComma)(parseExport))(
        state2$1.consumed ? {...state2$1, consumed: false} : state2$1,
        more,
        (state3, error) => {
          if (state3.consumed) { return resume(state3, error); }
          return more(v1$2 => $0(Data$dMaybe.Nothing)(state2$1, more, resume, done));
        },
        (state2$2, a$2) => more(v1$2 => $0(Data$dMaybe.$Maybe("Just", a$2))(state2$2, more, resume, done))
      );
    })
  ))
);
const braces = valueParser => (state1, more, resume, done) => tokLeftBrace(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => valueParser(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => tokRightBrace(state2$1, more, resume, (state2$2, a$2) => more(v1$2 => done(state2$2, {open: a, value: a$1, close: a$2}))))
  ))
);
const parseTypeVarKinded = parseBindingName => {
  const $0 = parens((state1, more, resume, done) => parseBindingName(
    state1,
    more,
    resume,
    (state2, a) => more(v1 => tokDoubleColon(
      state2,
      more,
      resume,
      (state2$1, a$1) => more(v1$1 => parseType$lazy()(state2$1, more, resume, (state2$2, a$2) => more(v1$2 => done(state2$2, {label: a, separator: a$1, value: a$2}))))
    ))
  ));
  return (state1, more, resume, done) => $0(state1, more, resume, (state2, a) => done(state2, PureScript$dCST$dTypes.$TypeVarBinding("TypeVarKinded", a)));
};
const parseTypeVarBinding = parseBindingName => {
  const $0 = parseTypeVarKinded(parseBindingName);
  return (state1, more, resume, done) => $0(
    state1.consumed ? {...state1, consumed: false} : state1,
    more,
    (state3, error) => {
      if (state3.consumed) { return resume(state3, error); }
      return parseBindingName(state1, more, resume, (state2, a) => done(state2, PureScript$dCST$dTypes.$TypeVarBinding("TypeVarName", a)));
    },
    done
  );
};
const parseTypeParens = (state1, more, resume, done) => tokLeftParen(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => {
    const $0 = parseRowTailParen(a);
    const $1 = parseKindedVar(a);
    const $2 = parseTypeParen(a);
    return parseRowParen(a)(
      state2.consumed ? {...state2, consumed: false} : state2,
      more,
      (state3, error) => {
        if (state3.consumed) { return resume(state3, error); }
        return $0(
          state2.consumed ? {...state2, consumed: false} : state2,
          more,
          (state3$1, error$1) => {
            if (state3$1.consumed) { return resume(state3$1, error$1); }
            return $1(
              state2.consumed ? {...state2, consumed: false} : state2,
              more,
              (state3$2, error$2) => {
                if (state3$2.consumed) { return resume(state3$2, error$2); }
                return $2(
                  state2.consumed ? {...state2, consumed: false} : state2,
                  more,
                  (state3$3, error$3) => {
                    if (state3$3.consumed) { return resume(state3$3, error$3); }
                    return parseEmptyRow(a)(state2, more, resume, done);
                  },
                  done
                );
              },
              done
            );
          },
          done
        );
      },
      done
    );
  })
);
const parseTypeParen = open => (state1, more, resume, done) => parseType$lazy()(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => tokRightParen(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => done(state2$1, PureScript$dCST$dTypes.$Type("TypeParens", {open, value: a, close: a$1})))
  ))
);
const parseRowTailParen = open => (state1, more, resume, done) => tokPipe(
  state1,
  more,
  resume,
  (state2, a) => {
    const $0 = Data$dTuple.Tuple(a);
    return more(v2 => parseType$lazy()(
      state2,
      more,
      resume,
      (state3, a$1) => {
        const $1 = $0(a$1);
        return more(v1 => tokRightParen(
          state3,
          more,
          resume,
          (state2$1, a$2) => more(v1$1 => done(
            state2$1,
            PureScript$dCST$dTypes.$Type("TypeRow", {open, value: {labels: Data$dMaybe.Nothing, tail: Data$dMaybe.$Maybe("Just", $1)}, close: a$2})
          ))
        ));
      }
    ));
  }
);
const parseRowParen = open => (state1, more, resume, done) => parseLabel(
  state1,
  more,
  (state2, error) => resume({...state2, consumed: state1.consumed}, error),
  (state2, a) => {
    const $0 = Data$dTuple.Tuple(a);
    return more(v2 => tokDoubleColon(
      state2,
      more,
      (state2$1, error) => resume({...state2$1, consumed: state1.consumed}, error),
      (state3, a$1) => {
        const $1 = $0(a$1);
        return more(v1 => {
          const $2 = $1._1;
          const $3 = $1._2;
          return parseType$lazy()(
            state3,
            more,
            resume,
            (state2$1, a$2) => more(v1$1 => PureScript$dCST$dParser$dMonad.many((state1$1, more$1, resume$1, done$1) => tokComma(
              state1$1,
              more$1,
              resume$1,
              (state2$2, a$3) => {
                const $4 = Data$dTuple.Tuple(a$3);
                return more$1(v2$1 => parseRowLabel(state2$2, more$1, resume$1, (state3$1, a$4) => done$1(state3$1, $4(a$4))));
              }
            ))(
              state2$1,
              more,
              resume,
              (state2$2, a$3) => more(v1$2 => {
                const $4 = (state2$3, a$4) => more(v1$3 => tokRightParen(
                  state2$3,
                  more,
                  resume,
                  (state2$4, a$5) => more(v1$4 => done(
                    state2$4,
                    PureScript$dCST$dTypes.$Type(
                      "TypeRow",
                      {open, value: {labels: Data$dMaybe.$Maybe("Just", {head: {label: $2, separator: $3, value: a$2}, tail: a$3}), tail: a$4}, close: a$5}
                    )
                  ))
                ));
                return tokPipe(
                  state2$2.consumed ? {...state2$2, consumed: false} : state2$2,
                  more,
                  (state3$1, error) => {
                    if (state3$1.consumed) { return resume(state3$1, error); }
                    return $4(state2$2, Data$dMaybe.Nothing);
                  },
                  (state2$3, a$4) => {
                    const $5 = Data$dTuple.Tuple(a$4);
                    return more(v2$1 => parseType$lazy()(
                      state2$3,
                      more,
                      (state3$1, error) => {
                        if (state3$1.consumed) { return resume(state3$1, error); }
                        return $4(state2$2, Data$dMaybe.Nothing);
                      },
                      (state3$1, a$5) => $4(state3$1, Data$dMaybe.$Maybe("Just", $5(a$5)))
                    ));
                  }
                );
              })
            ))
          );
        });
      }
    ));
  }
);
const parseRowLabel = (state1, more, resume, done) => parseLabel(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => tokDoubleColon(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => parseType$lazy()(state2$1, more, resume, (state2$2, a$2) => more(v1$2 => done(state2$2, {label: a, separator: a$1, value: a$2}))))
  ))
);
const parseKindedVar = open => {
  const $0 = parens((state1, more, resume, done) => parseIdent(state1, more, resume, (state2, a) => done(state2, PureScript$dCST$dTypes.$Type("TypeVar", a))));
  return (state1, more, resume, done) => $0(
    state1,
    more,
    (state2, error) => resume({...state2, consumed: state1.consumed}, error),
    (state2, a) => {
      const $1 = Data$dTuple.Tuple(a);
      return more(v2 => tokDoubleColon(
        state2,
        more,
        (state2$1, error) => resume({...state2$1, consumed: state1.consumed}, error),
        (state3, a$1) => {
          const $2 = $1(a$1);
          return more(v1 => {
            const $3 = $2._2;
            const $4 = $2._1;
            return parseType$lazy()(
              state3,
              more,
              resume,
              (state2$1, a$2) => more(v1$1 => tokRightParen(
                state2$1,
                more,
                resume,
                (state2$2, a$3) => more(v1$2 => done(
                  state2$2,
                  PureScript$dCST$dTypes.$Type(
                    "TypeParens",
                    {open, value: PureScript$dCST$dTypes.$Type("TypeKinded", PureScript$dCST$dTypes.$Type("TypeParens", $4), $3, a$2), close: a$3}
                  )
                ))
              ))
            );
          });
        }
      ));
    }
  );
};
const parseTypeVarBindingWithVisibility$lazy = /* #__PURE__ */ $runtime.binding(() => PureScript$dCST$dParser$dMonad.lazyParser.defer(v => parseTypeVarBinding((
  state1,
  more,
  resume,
  done
) => {
  const $0 = (state2, f) => more(v2 => parseIdent(state2, more, resume, (state3, a) => done(state3, f(a))));
  return tokAt(
    state1.consumed ? {...state1, consumed: false} : state1,
    more,
    (state3, error) => {
      if (state3.consumed) { return resume(state3, error); }
      return $0(state1, v2 => ({prefix: Data$dMaybe.Nothing, value: v2}));
    },
    (state2, a) => $0(state2, v2 => ({prefix: Data$dMaybe.$Maybe("Just", a), value: v2}))
  );
})));
const parseTypeAtom$lazy = /* #__PURE__ */ $runtime.binding(() => PureScript$dCST$dParser$dMonad.lazyParser.defer(v => {
  const $0 = PureScript$dCST$dTypes.TypeInt(Data$dMaybe.Nothing);
  return (state1, more, resume, done) => parseIdent(
    state1.consumed ? {...state1, consumed: false} : state1,
    more,
    (state3, error) => {
      if (state3.consumed) { return resume(state3, error); }
      return parseQualifiedProper(
        state1.consumed ? {...state1, consumed: false} : state1,
        more,
        (state3$1, error$1) => {
          if (state3$1.consumed) { return resume(state3$1, error$1); }
          return parseString(
            state1.consumed ? {...state1, consumed: false} : state1,
            more,
            (state3$2, error$2) => {
              if (state3$2.consumed) { return resume(state3$2, error$2); }
              return parseInt(
                state1.consumed ? {...state1, consumed: false} : state1,
                more,
                (state3$3, error$3) => {
                  if (state3$3.consumed) { return resume(state3$3, error$3); }
                  return parseTypeParens(
                    state1.consumed ? {...state1, consumed: false} : state1,
                    more,
                    (state3$4, error$4) => {
                      if (state3$4.consumed) { return resume(state3$4, error$4); }
                      return braces(parseRow$lazy())(
                        state1.consumed ? {...state1, consumed: false} : state1,
                        more,
                        (state3$5, error$5) => {
                          if (state3$5.consumed) { return resume(state3$5, error$5); }
                          return parseQualifiedSymbol(
                            state1.consumed ? {...state1, consumed: false} : state1,
                            more,
                            (state3$6, error$6) => {
                              if (state3$6.consumed) { return resume(state3$6, error$6); }
                              return parseHole(
                                state1.consumed ? {...state1, consumed: false} : state1,
                                more,
                                (state3$7, error$7) => {
                                  if (state3$7.consumed) { return resume(state3$7, error$7); }
                                  return tokUnderscore(
                                    state1.consumed ? {...state1, consumed: false} : state1,
                                    more,
                                    (state3$8, error$8) => {
                                      if (state3$8.consumed) { return resume(state3$8, error$8); }
                                      return tokSymbolArrow(state1, more, resume, (state2, a) => done(state2, PureScript$dCST$dTypes.$Type("TypeArrowName", a)));
                                    },
                                    (state2, a) => done(state2, PureScript$dCST$dTypes.$Type("TypeWildcard", a))
                                  );
                                },
                                (state2, a) => done(state2, PureScript$dCST$dTypes.$Type("TypeHole", a))
                              );
                            },
                            (state2, a) => done(state2, PureScript$dCST$dTypes.$Type("TypeOpName", a))
                          );
                        },
                        (state2, a) => done(state2, PureScript$dCST$dTypes.$Type("TypeRecord", a))
                      );
                    },
                    done
                  );
                },
                (state2, a) => done(state2, $0(a._1)(a._2))
              );
            },
            (state2, a) => done(state2, PureScript$dCST$dTypes.$Type("TypeString", a._1, a._2))
          );
        },
        (state2, a) => done(state2, PureScript$dCST$dTypes.$Type("TypeConstructor", a))
      );
    },
    (state2, a) => done(state2, PureScript$dCST$dTypes.$Type("TypeVar", a))
  );
}));
const parseType5$lazy = /* #__PURE__ */ $runtime.binding(() => PureScript$dCST$dParser$dMonad.lazyParser.defer(v => (state1, more, resume, done) => parseTypeAtom$lazy()(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => PureScript$dCST$dParser$dMonad.many(parseTypeAtom$lazy())(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => {
      if (a$1.length > 0) { return done(state2$1, PureScript$dCST$dTypes.$Type("TypeApp", a, a$1)); }
      return done(state2$1, a);
    })
  ))
)));
const parseType4$lazy = /* #__PURE__ */ $runtime.binding(() => PureScript$dCST$dParser$dMonad.lazyParser.defer(v => (state1, more, resume, done) => parseTypeNegative(
  state1.consumed ? {...state1, consumed: false} : state1,
  more,
  (state3, error) => {
    if (state3.consumed) { return resume(state3, error); }
    return parseType5$lazy()(state1, more, resume, done);
  },
  done
)));
const parseType3$lazy = /* #__PURE__ */ $runtime.binding(() => PureScript$dCST$dParser$dMonad.lazyParser.defer(v => (state1, more, resume, done) => parseType4$lazy()(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => PureScript$dCST$dParser$dMonad.many((state1$1, more$1, resume$1, done$1) => parseQualifiedOperator(
    state1$1,
    more$1,
    resume$1,
    (state2$1, a$1) => {
      const $0 = Data$dTuple.Tuple(a$1);
      return more$1(v2 => parseType4$lazy()(state2$1, more$1, resume$1, (state3, a$2) => done$1(state3, $0(a$2))));
    }
  ))(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => {
      if (a$1.length > 0) { return done(state2$1, PureScript$dCST$dTypes.$Type("TypeOp", a, a$1)); }
      return done(state2$1, a);
    })
  ))
)));
const parseType2$lazy = /* #__PURE__ */ $runtime.binding(() => PureScript$dCST$dParser$dMonad.lazyParser.defer(v => (state1, more, resume, done) => parseType3$lazy()(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => {
    const $0 = PureScript$dCST$dTypes.TypeArrow(a);
    const $1 = PureScript$dCST$dTypes.TypeConstrained(a);
    const $2 = (state1$1, more$1, resume$1, done$1) => tokRightFatArrow(
      state1$1.consumed ? {...state1$1, consumed: false} : state1$1,
      more$1,
      (state3, error) => {
        if (state3.consumed) { return resume$1(state3, error); }
        return done$1(state1$1, a);
      },
      (state2$1, a$1) => {
        const $2 = $1(a$1);
        return more$1(v2 => parseType1$lazy()(
          state2$1,
          more$1,
          (state3, error) => {
            if (state3.consumed) { return resume$1(state3, error); }
            return done$1(state1$1, a);
          },
          (state3, a$2) => done$1(state3, $2(a$2))
        ));
      }
    );
    return tokRightArrow(
      state2.consumed ? {...state2, consumed: false} : state2,
      more,
      (state3, error) => {
        if (state3.consumed) { return resume(state3, error); }
        return $2(state2, more, resume, done);
      },
      (state2$1, a$1) => {
        const $3 = $0(a$1);
        return more(v2 => parseType1$lazy()(
          state2$1,
          more,
          (state3, error) => {
            if (state3.consumed) { return resume(state3, error); }
            return $2(state2, more, resume, done);
          },
          (state3, a$2) => done(state3, $3(a$2))
        ));
      }
    );
  })
)));
const parseType1$lazy = /* #__PURE__ */ $runtime.binding(() => PureScript$dCST$dParser$dMonad.lazyParser.defer(v => (state1, more, resume, done) => parseForall$lazy()(
  state1.consumed ? {...state1, consumed: false} : state1,
  more,
  (state3, error) => {
    if (state3.consumed) { return resume(state3, error); }
    return parseType2$lazy()(state1, more, resume, done);
  },
  done
)));
const parseType$lazy = /* #__PURE__ */ $runtime.binding(() => PureScript$dCST$dParser$dMonad.lazyParser.defer(v => (state1, more, resume, done) => parseType1$lazy()(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => {
    const $0 = PureScript$dCST$dTypes.TypeKinded(a);
    return tokDoubleColon(
      state2.consumed ? {...state2, consumed: false} : state2,
      more,
      (state3, error) => {
        if (state3.consumed) { return resume(state3, error); }
        return done(state2, a);
      },
      (state2$1, a$1) => {
        const $1 = $0(a$1);
        return more(v2 => parseType$lazy()(
          state2$1,
          more,
          (state3, error) => {
            if (state3.consumed) { return resume(state3, error); }
            return done(state2, a);
          },
          (state3, a$2) => done(state3, $1(a$2))
        ));
      }
    );
  })
)));
const parseRow$lazy = /* #__PURE__ */ $runtime.binding(() => PureScript$dCST$dParser$dMonad.lazyParser.defer(v => {
  const $0 = separated(tokComma)(parseRowLabel);
  return (state1, more, resume, done) => $0(
    state1.consumed ? {...state1, consumed: false} : state1,
    more,
    (state3, error) => {
      if (state3.consumed) { return resume(state3, error); }
      return more(v1 => tokPipe(
        state1.consumed ? {...state1, consumed: false} : state1,
        more,
        (state3$1, error$1) => {
          if (state3$1.consumed) { return resume(state3$1, error$1); }
          return more(v1$1 => done(state1, {labels: Data$dMaybe.Nothing, tail: Data$dMaybe.Nothing}));
        },
        (state2, a) => {
          const $1 = Data$dTuple.Tuple(a);
          return more(v2 => parseType$lazy()(
            state2,
            more,
            (state3$1, error$1) => {
              if (state3$1.consumed) { return resume(state3$1, error$1); }
              return more(v1$1 => done(state1, {labels: Data$dMaybe.Nothing, tail: Data$dMaybe.Nothing}));
            },
            (state3$1, a$1) => more(v1$1 => done(state3$1, {labels: Data$dMaybe.Nothing, tail: Data$dMaybe.$Maybe("Just", $1(a$1))}))
          ));
        }
      ));
    },
    (state2, a) => more(v1 => tokPipe(
      state2.consumed ? {...state2, consumed: false} : state2,
      more,
      (state3, error) => {
        if (state3.consumed) { return resume(state3, error); }
        return more(v1$1 => done(state2, {labels: Data$dMaybe.$Maybe("Just", a), tail: Data$dMaybe.Nothing}));
      },
      (state2$1, a$1) => {
        const $1 = Data$dTuple.Tuple(a$1);
        return more(v2 => parseType$lazy()(
          state2$1,
          more,
          (state3, error) => {
            if (state3.consumed) { return resume(state3, error); }
            return more(v1$1 => done(state2, {labels: Data$dMaybe.$Maybe("Just", a), tail: Data$dMaybe.Nothing}));
          },
          (state3, a$2) => more(v1$1 => done(state3, {labels: Data$dMaybe.$Maybe("Just", a), tail: Data$dMaybe.$Maybe("Just", $1(a$2))}))
        ));
      }
    ))
  );
}));
const parseForall$lazy = /* #__PURE__ */ $runtime.binding(() => PureScript$dCST$dParser$dMonad.lazyParser.defer(v => (state1, more, resume, done) => tokForall(
  state1,
  more,
  resume,
  (state2, a) => {
    const $0 = PureScript$dCST$dTypes.TypeForall(a);
    return more(v2 => many1(parseTypeVarBindingWithVisibility$lazy())(
      state2,
      more,
      resume,
      (state3, a$1) => {
        const $1 = $0(a$1);
        return more(v2$1 => tokDot(
          state3,
          more,
          resume,
          (state3$1, a$2) => {
            const $2 = $1(a$2);
            return more(v2$2 => parseType1$lazy()(state3$1, more, resume, (state3$2, a$3) => done(state3$2, $2(a$3))));
          }
        ));
      }
    ));
  }
)));
const parseTypeVarBindingWithVisibility = /* #__PURE__ */ parseTypeVarBindingWithVisibility$lazy();
const parseTypeAtom = /* #__PURE__ */ parseTypeAtom$lazy();
const parseType5 = /* #__PURE__ */ parseType5$lazy();
const parseType4 = /* #__PURE__ */ parseType4$lazy();
const parseType3 = /* #__PURE__ */ parseType3$lazy();
const parseType2 = /* #__PURE__ */ parseType2$lazy();
const parseType1 = /* #__PURE__ */ parseType1$lazy();
const parseType = /* #__PURE__ */ parseType$lazy();
const parseRow = /* #__PURE__ */ parseRow$lazy();
const parseForall = /* #__PURE__ */ parseForall$lazy();
const parseDataCtor = (state1, more, resume, done) => parseProper(
  state1,
  more,
  resume,
  (state2, a) => more(v2 => PureScript$dCST$dParser$dMonad.many(parseTypeAtom)(state2, more, resume, (state3, a$1) => done(state3, {name: a, fields: a$1})))
);
const parseIdentBinder = (state1, more, resume, done) => parseIdent(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => {
    const $0 = PureScript$dCST$dTypes.BinderNamed(a);
    return tokAt(
      state2.consumed ? {...state2, consumed: false} : state2,
      more,
      (state3, error) => {
        if (state3.consumed) { return resume(state3, error); }
        return done(state2, PureScript$dCST$dTypes.$Binder("BinderVar", a));
      },
      (state2$1, a$1) => {
        const $1 = $0(a$1);
        return more(v2 => parseBinderAtom$lazy()(
          state2$1,
          more,
          (state3, error) => {
            if (state3.consumed) { return resume(state3, error); }
            return done(state2, PureScript$dCST$dTypes.$Binder("BinderVar", a));
          },
          (state3, a$2) => done(state3, $1(a$2))
        ));
      }
    );
  })
);
const parseBinderConstructor$lazy = /* #__PURE__ */ $runtime.binding(() => PureScript$dCST$dParser$dMonad.lazyParser.defer(v => (state1, more, resume, done) => parseQualifiedProper(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => PureScript$dCST$dParser$dMonad.many(parseBinderAtom$lazy())(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => done(state2$1, PureScript$dCST$dTypes.$Binder("BinderConstructor", a, a$1)))
  ))
)));
const parseBinderAtom$lazy = /* #__PURE__ */ $runtime.binding(() => PureScript$dCST$dParser$dMonad.lazyParser.defer(v => {
  const $0 = PureScript$dCST$dTypes.BinderInt(Data$dMaybe.Nothing);
  const $1 = PureScript$dCST$dTypes.BinderNumber(Data$dMaybe.Nothing);
  return (state1, more, resume, done) => parseIdentBinder(
    state1.consumed ? {...state1, consumed: false} : state1,
    more,
    (state3, error) => {
      if (state3.consumed) { return resume(state3, error); }
      return parseQualifiedProper(
        state1.consumed ? {...state1, consumed: false} : state1,
        more,
        (state3$1, error$1) => {
          if (state3$1.consumed) { return resume(state3$1, error$1); }
          return tokUnderscore(
            state1.consumed ? {...state1, consumed: false} : state1,
            more,
            (state3$2, error$2) => {
              if (state3$2.consumed) { return resume(state3$2, error$2); }
              return parseString(
                state1.consumed ? {...state1, consumed: false} : state1,
                more,
                (state3$3, error$3) => {
                  if (state3$3.consumed) { return resume(state3$3, error$3); }
                  return parseChar(
                    state1.consumed ? {...state1, consumed: false} : state1,
                    more,
                    (state3$4, error$4) => {
                      if (state3$4.consumed) { return resume(state3$4, error$4); }
                      return parseBoolean(
                        state1.consumed ? {...state1, consumed: false} : state1,
                        more,
                        (state3$5, error$5) => {
                          if (state3$5.consumed) { return resume(state3$5, error$5); }
                          return parseInt(
                            state1.consumed ? {...state1, consumed: false} : state1,
                            more,
                            (state3$6, error$6) => {
                              if (state3$6.consumed) { return resume(state3$6, error$6); }
                              return parseNumber(
                                state1.consumed ? {...state1, consumed: false} : state1,
                                more,
                                (state3$7, error$7) => {
                                  if (state3$7.consumed) { return resume(state3$7, error$7); }
                                  return delimited(tokLeftSquare)(tokRightSquare)(tokComma)(parseBinder$lazy())(
                                    state1.consumed ? {...state1, consumed: false} : state1,
                                    more,
                                    (state3$8, error$8) => {
                                      if (state3$8.consumed) { return resume(state3$8, error$8); }
                                      return delimited(tokLeftBrace)(tokRightBrace)(tokComma)(parseRecordLabeled(parseBinder$lazy()))(
                                        state1.consumed ? {...state1, consumed: false} : state1,
                                        more,
                                        (state3$9, error$9) => {
                                          if (state3$9.consumed) { return resume(state3$9, error$9); }
                                          return parens(parseBinder$lazy())(state1, more, resume, (state2, a) => done(state2, PureScript$dCST$dTypes.$Binder("BinderParens", a)));
                                        },
                                        (state2, a) => done(state2, PureScript$dCST$dTypes.$Binder("BinderRecord", a))
                                      );
                                    },
                                    (state2, a) => done(state2, PureScript$dCST$dTypes.$Binder("BinderArray", a))
                                  );
                                },
                                (state2, a) => done(state2, $1(a._1)(a._2))
                              );
                            },
                            (state2, a) => done(state2, $0(a._1)(a._2))
                          );
                        },
                        (state2, a) => done(state2, PureScript$dCST$dTypes.$Binder("BinderBoolean", a._1, a._2))
                      );
                    },
                    (state2, a) => done(state2, PureScript$dCST$dTypes.$Binder("BinderChar", a._1, a._2))
                  );
                },
                (state2, a) => done(state2, PureScript$dCST$dTypes.$Binder("BinderString", a._1, a._2))
              );
            },
            (state2, a) => done(state2, PureScript$dCST$dTypes.$Binder("BinderWildcard", a))
          );
        },
        (state2, a) => done(state2, PureScript$dCST$dTypes.$Binder("BinderConstructor", a, []))
      );
    },
    done
  );
}));
const parseBinder2$lazy = /* #__PURE__ */ $runtime.binding(() => PureScript$dCST$dParser$dMonad.lazyParser.defer(v => (state1, more, resume, done) => parseBinderNegative(
  state1.consumed ? {...state1, consumed: false} : state1,
  more,
  (state3, error) => {
    if (state3.consumed) { return resume(state3, error); }
    return parseBinderConstructor$lazy()(
      state1.consumed ? {...state1, consumed: false} : state1,
      more,
      (state3$1, error$1) => {
        if (state3$1.consumed) { return resume(state3$1, error$1); }
        return parseBinderAtom$lazy()(state1, more, resume, done);
      },
      done
    );
  },
  done
)));
const parseBinder1$lazy = /* #__PURE__ */ $runtime.binding(() => PureScript$dCST$dParser$dMonad.lazyParser.defer(v => (state1, more, resume, done) => parseBinder2$lazy()(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => PureScript$dCST$dParser$dMonad.many((state1$1, more$1, resume$1, done$1) => parseQualifiedOperator(
    state1$1,
    more$1,
    resume$1,
    (state2$1, a$1) => {
      const $0 = Data$dTuple.Tuple(a$1);
      return more$1(v2 => parseBinder2$lazy()(state2$1, more$1, resume$1, (state3, a$2) => done$1(state3, $0(a$2))));
    }
  ))(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => {
      if (a$1.length > 0) { return done(state2$1, PureScript$dCST$dTypes.$Binder("BinderOp", a, a$1)); }
      return done(state2$1, a);
    })
  ))
)));
const parseBinder$lazy = /* #__PURE__ */ $runtime.binding(() => PureScript$dCST$dParser$dMonad.lazyParser.defer(v => (state1, more, resume, done) => parseBinder1$lazy()(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => {
    const $0 = PureScript$dCST$dTypes.BinderTyped(a);
    return tokDoubleColon(
      state2.consumed ? {...state2, consumed: false} : state2,
      more,
      (state3, error) => {
        if (state3.consumed) { return resume(state3, error); }
        return done(state2, a);
      },
      (state2$1, a$1) => {
        const $1 = $0(a$1);
        return more(v2 => parseType(
          state2$1,
          more,
          (state3, error) => {
            if (state3.consumed) { return resume(state3, error); }
            return done(state2, a);
          },
          (state3, a$2) => done(state3, $1(a$2))
        ));
      }
    );
  })
)));
const parseBinderConstructor = /* #__PURE__ */ parseBinderConstructor$lazy();
const parseBinderAtom = /* #__PURE__ */ parseBinderAtom$lazy();
const parseBinder2 = /* #__PURE__ */ parseBinder2$lazy();
const parseBinder1 = /* #__PURE__ */ parseBinder1$lazy();
const parseBinder = /* #__PURE__ */ parseBinder$lazy();
const parseClassConstraints = parseOneConstraint => {
  const $0 = parens(separated(tokComma)(parseType));
  return (state1, more, resume, done) => $0(
    state1.consumed ? {...state1, consumed: false} : state1,
    more,
    (state3, error) => {
      if (state3.consumed) { return resume(state3, error); }
      return parseOneConstraint(state1, more, resume, (state2, a) => done(state2, PureScript$dCST$dTypes.$OneOrDelimited("One", a)));
    },
    (state2, a) => done(state2, PureScript$dCST$dTypes.$OneOrDelimited("Many", a))
  );
};
const parseDeclDerive = (state1, more, resume, done) => tokKeyword("derive")(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => {
    const $0 = newtype_ => (state1$1, more$1, resume$1, done$1) => tokKeyword("instance")(
      state1$1,
      more$1,
      resume$1,
      (state2$1, a$1) => more$1(v1$1 => {
        const $0 = name => {
          const $0 = parseClassConstraints(parseType3);
          return (state1$2, more$2, resume$2, done$2) => {
            const $1 = (state2$2, a$2) => more$2(v1$2 => parseQualifiedProper(
              state2$2,
              more$2,
              resume$2,
              (state2$3, a$3) => more$2(v1$3 => PureScript$dCST$dParser$dMonad.many(parseTypeAtom)(
                state2$3,
                more$2,
                resume$2,
                (state2$4, a$4) => more$2(v1$4 => done$2(
                  state2$4,
                  PureScript$dCST$dTypes.$Declaration("DeclDerive", a, newtype_, {keyword: a$1, name, constraints: a$2, className: a$3, types: a$4})
                ))
              ))
            ));
            const $2 = state1$2.consumed ? {...state1$2, consumed: false} : state1$2;
            return $0(
              $2,
              more$2,
              (state2$2, error) => {
                const $3 = {...state2$2, consumed: $2.consumed};
                if ($3.consumed) { return resume$2($3, error); }
                return $1(state1$2, Data$dMaybe.Nothing);
              },
              (state2$2, a$2) => {
                const $3 = Data$dTuple.Tuple(a$2);
                return more$2(v2 => tokRightFatArrow(
                  state2$2,
                  more$2,
                  (state2$3, error) => {
                    const $4 = {...state2$3, consumed: $2.consumed};
                    if ($4.consumed) { return resume$2($4, error); }
                    return $1(state1$2, Data$dMaybe.Nothing);
                  },
                  (state3, a$3) => $1(state3, Data$dMaybe.$Maybe("Just", $3(a$3)))
                ));
              }
            );
          };
        };
        return parseInstanceName(
          state2$1.consumed ? {...state2$1, consumed: false} : state2$1,
          more$1,
          (state3, error) => {
            if (state3.consumed) { return resume$1(state3, error); }
            return more$1(v1$2 => $0(Data$dMaybe.Nothing)(state2$1, more$1, resume$1, done$1));
          },
          (state2$2, a$2) => more$1(v1$2 => $0(Data$dMaybe.$Maybe("Just", a$2))(state2$2, more$1, resume$1, done$1))
        );
      })
    );
    return tokKeyword("newtype")(
      state2.consumed ? {...state2, consumed: false} : state2,
      more,
      (state3, error) => {
        if (state3.consumed) { return resume(state3, error); }
        return more(v1$1 => $0(Data$dMaybe.Nothing)(state2, more, resume, done));
      },
      (state2$1, a$1) => more(v1$1 => $0(Data$dMaybe.$Maybe("Just", a$1))(state2$1, more, resume, done))
    );
  })
);
const parseClassMember = (state1, more, resume, done) => parseIdent(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => tokDoubleColon(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => parseType(state2$1, more, resume, (state2$2, a$2) => more(v1$2 => done(state2$2, {label: a, separator: a$1, value: a$2}))))
  ))
);
const parseDeclClassSignature = keyword => (state1, more, resume, done) => parseProper(
  state1,
  more,
  (state2, error) => resume({...state2, consumed: state1.consumed}, error),
  (state2, a) => {
    const $0 = Data$dTuple.Tuple(a);
    return more(v2 => tokDoubleColon(
      state2,
      more,
      (state2$1, error) => resume({...state2$1, consumed: state1.consumed}, error),
      (state3, a$1) => {
        const $1 = $0(a$1);
        return more(v1 => {
          const $2 = $1._1;
          const $3 = $1._2;
          return parseType(
            state3,
            more,
            resume,
            (state2$1, a$2) => more(v1$1 => done(state2$1, PureScript$dCST$dTypes.$Declaration("DeclKindSignature", keyword, {label: $2, separator: $3, value: a$2})))
          );
        });
      }
    ));
  }
);
const parseDeclKindSignature = keyword => label => (state1, more, resume, done) => tokDoubleColon(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => parseType(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => done(state2$1, PureScript$dCST$dTypes.$Declaration("DeclKindSignature", keyword, {label, separator: a, value: a$1})))
  ))
);
const parseDeclSignature = label => (state1, more, resume, done) => tokDoubleColon(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => parseType(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => done(state2$1, PureScript$dCST$dTypes.$Declaration("DeclSignature", {label, separator: a, value: a$1})))
  ))
);
const parseForeignData = (state1, more, resume, done) => tokKeyword("data")(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => parseProper(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => tokDoubleColon(
      state2$1,
      more,
      resume,
      (state2$2, a$2) => more(v1$2 => parseType(
        state2$2,
        more,
        resume,
        (state2$3, a$3) => more(v1$3 => done(state2$3, PureScript$dCST$dTypes.$Foreign("ForeignData", a, {label: a$1, separator: a$2, value: a$3})))
      ))
    ))
  ))
);
const parseForeignValue = (state1, more, resume, done) => parseIdent(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => tokDoubleColon(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => parseType(
      state2$1,
      more,
      resume,
      (state2$2, a$2) => more(v1$2 => done(state2$2, PureScript$dCST$dTypes.$Foreign("ForeignValue", {label: a, separator: a$1, value: a$2})))
    ))
  ))
);
const parseDeclForeign = (state1, more, resume, done) => tokKeyword("foreign")(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => tokKeyword("import")(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => parseForeignData(
      state2$1.consumed ? {...state2$1, consumed: false} : state2$1,
      more,
      (state3, error) => {
        if (state3.consumed) { return resume(state3, error); }
        return parseForeignKind(
          state2$1.consumed ? {...state2$1, consumed: false} : state2$1,
          more,
          (state3$1, error$1) => {
            if (state3$1.consumed) { return resume(state3$1, error$1); }
            return parseForeignValue(state2$1, more, resume, (state2$2, a$2) => more(v1$2 => done(state2$2, PureScript$dCST$dTypes.$Declaration("DeclForeign", a, a$1, a$2))));
          },
          (state2$2, a$2) => more(v1$2 => done(state2$2, PureScript$dCST$dTypes.$Declaration("DeclForeign", a, a$1, a$2)))
        );
      },
      (state2$2, a$2) => more(v1$2 => done(state2$2, PureScript$dCST$dTypes.$Declaration("DeclForeign", a, a$1, a$2)))
    ))
  ))
);
const parseInstanceBindingSignature = label => (state1, more, resume, done) => tokDoubleColon(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => parseType(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => done(state2$1, PureScript$dCST$dTypes.$InstanceBinding("InstanceBindingSignature", {label, separator: a, value: a$1})))
  ))
);
const parseLetBindingSignature = label => (state1, more, resume, done) => tokDoubleColon(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => parseType(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => done(state2$1, PureScript$dCST$dTypes.$LetBinding("LetBindingSignature", {label, separator: a, value: a$1})))
  ))
);
const parseTickExpr = (state1, more, resume, done) => tokTick(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => parseTickExpr1$lazy()(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => tokTick(state2$1, more, resume, (state2$2, a$2) => more(v1$2 => done(state2$2, {open: a, value: a$1, close: a$2}))))
  ))
);
const parseRecordUpdates = expr => (state1, more, resume, done) => {
  const $0 = (state2, a) => more(v1 => separated(tokComma)(parseRecordUpdate)(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => tokRightBrace(
      state2$1,
      more,
      resume,
      (state2$2, a$2) => more(v1$2 => done(state2$2, PureScript$dCST$dTypes.$Expr("ExprRecordUpdate", expr, {open: a, value: a$1, close: a$2})))
    ))
  ));
  return tokLeftBrace(
    state1,
    more,
    (state2, error) => resume({...state2, consumed: state1.consumed}, error),
    (state2, a) => more(v2 => parseLabel(
      state2,
      more,
      (v1, error) => resume({...state2, consumed: state1.consumed}, error),
      (state2$1, a$1) => more(v2$1 => tokEquals(
        state2$1.consumed ? {...state2$1, consumed: false} : state2$1,
        more,
        (state3, error) => {
          if (state3.consumed) { return resume({...state2, consumed: state1.consumed}, error); }
          return tokLeftBrace(state2$1, more, (v1, error$1) => resume({...state2, consumed: state1.consumed}, error$1), (state3$1, a$2) => $0(state2, a));
        },
        (state3, a$2) => $0(state2, a)
      ))
    ))
  );
};
const parseRecordUpdateLeaf = label => {
  const $0 = PureScript$dCST$dTypes.RecordUpdateLeaf(label);
  return (state1, more, resume, done) => tokEquals(
    state1,
    more,
    resume,
    (state2, a) => {
      const $1 = $0(a);
      return more(v2 => parseExpr$lazy()(state2, more, resume, (state3, a$1) => done(state3, $1(a$1))));
    }
  );
};
const parseRecordUpdateBranch = label => {
  const $0 = PureScript$dCST$dTypes.RecordUpdateBranch(label);
  const $1 = braces(separated(tokComma)(parseRecordUpdate));
  return (state1, more, resume, done) => $1(state1, more, resume, (state2, a) => done(state2, $0(a)));
};
const parseRecordUpdate = (state1, more, resume, done) => parseLabel(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => {
    const $0 = parseRecordUpdateBranch(a);
    return parseRecordUpdateLeaf(a)(
      state2.consumed ? {...state2, consumed: false} : state2,
      more,
      (state3, error) => {
        if (state3.consumed) { return resume(state3, error); }
        return $0(state2, more, resume, done);
      },
      done
    );
  })
);
const parseLetIn = (state1, more, resume, done) => tokKeyword("let")(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => layoutNonEmpty(recoverIndent(PureScript$dCST$dTypes.LetBindingError)(parseLetBinding$lazy()))(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => tokKeyword("in")(
      state2$1,
      more,
      resume,
      (state2$2, a$2) => more(v1$2 => parseExpr$lazy()(
        state2$2,
        more,
        resume,
        (state2$3, a$3) => more(v1$3 => done(state2$3, PureScript$dCST$dTypes.$Expr("ExprLet", {keyword: a, bindings: a$1, in: a$2, body: a$3})))
      ))
    ))
  ))
);
const parseLetBindingName = name => (state1, more, resume, done) => PureScript$dCST$dParser$dMonad.many(parseBinderAtom)(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => parseGuarded(tokEquals)(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => done(state2$1, PureScript$dCST$dTypes.$LetBinding("LetBindingName", {name, binders: a, guarded: a$1})))
  ))
);
const parseLambda = (state1, more, resume, done) => tokBackslash(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => many1(parseBinderAtom)(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => tokRightArrow(
      state2$1,
      more,
      resume,
      (state2$2, a$2) => more(v1$2 => parseExpr$lazy()(
        state2$2,
        more,
        resume,
        (state2$3, a$3) => more(v1$3 => done(state2$3, PureScript$dCST$dTypes.$Expr("ExprLambda", {symbol: a, binders: a$1, arrow: a$2, body: a$3})))
      ))
    ))
  ))
);
const parseIf = (state1, more, resume, done) => tokKeyword("if")(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => parseExpr$lazy()(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => tokKeyword("then")(
      state2$1,
      more,
      resume,
      (state2$2, a$2) => more(v1$2 => parseExpr$lazy()(
        state2$2,
        more,
        resume,
        (state2$3, a$3) => more(v1$3 => tokKeyword("else")(
          state2$3,
          more,
          resume,
          (state2$4, a$4) => more(v1$4 => parseExpr$lazy()(
            state2$4,
            more,
            resume,
            (state2$5, a$5) => more(v1$5 => done(state2$5, PureScript$dCST$dTypes.$Expr("ExprIf", {keyword: a, cond: a$1, then: a$2, true: a$3, else: a$4, false: a$5})))
          ))
        ))
      ))
    ))
  ))
);
const parseIdentBinding = (state1, more, resume, done) => parseIdent(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => parseLetBindingSignature(a)(
    state2.consumed ? {...state2, consumed: false} : state2,
    more,
    (state3, error) => {
      if (state3.consumed) { return resume(state3, error); }
      return parseLetBindingName(a)(state2, more, resume, done);
    },
    done
  ))
);
const parseGuarded = sepParser => {
  const $0 = many1((() => {
    const $0 = separated(tokComma)((state1, more, resume, done) => {
      const $0 = (state2, f) => more(v2 => parseExpr$lazy()(state2, more, resume, (state3, a) => done(state3, f(a))));
      const $1 = state1.consumed ? {...state1, consumed: false} : state1;
      return parseBinder(
        $1,
        more,
        (state2, error) => {
          const $2 = {...state2, consumed: $1.consumed};
          if ($2.consumed) { return resume($2, error); }
          return $0(state1, v1 => ({binder: Data$dMaybe.Nothing, expr: v1}));
        },
        (state2, a) => {
          const $2 = Data$dTuple.Tuple(a);
          return more(v2 => tokLeftArrow(
            state2,
            more,
            (state2$1, error) => {
              const $3 = {...state2$1, consumed: $1.consumed};
              if ($3.consumed) { return resume($3, error); }
              return $0(state1, v1 => ({binder: Data$dMaybe.Nothing, expr: v1}));
            },
            (state3, a$1) => {
              const $3 = Data$dMaybe.$Maybe("Just", $2(a$1));
              return $0(state3, v1 => ({binder: $3, expr: v1}));
            }
          ));
        }
      );
    });
    return (state1, more, resume, done) => tokPipe(
      state1,
      more,
      resume,
      (state2, a) => more(v2 => $0(
        state2,
        more,
        resume,
        (state3, a$1) => more(v2$1 => sepParser(
          state3,
          more,
          resume,
          (state3$1, a$2) => more(v2$2 => parseWhere$lazy()(state3$1, more, resume, (state3$2, a$3) => done(state3$2, {bar: a, patterns: a$1, separator: a$2, where: a$3})))
        ))
      ))
    );
  })());
  return (state1, more, resume, done) => {
    const $1 = (state3, error) => {
      if (state3.consumed) { return resume(state3, error); }
      return $0(state1, more, resume, (state2, a) => done(state2, PureScript$dCST$dTypes.$Guarded("Guarded", a)));
    };
    return sepParser(
      state1.consumed ? {...state1, consumed: false} : state1,
      more,
      $1,
      (state2, a) => {
        const $2 = PureScript$dCST$dTypes.Unconditional(a);
        return more(v2 => parseWhere$lazy()(state2, more, $1, (state3, a$1) => done(state3, $2(a$1))));
      }
    );
  };
};
const parseDo = (state1, more, resume, done) => tokQualifiedKeyword("do")(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => layoutNonEmpty(recoverIndent(PureScript$dCST$dTypes.DoError)(parseDoStatement$lazy()))(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => done(state2$1, PureScript$dCST$dTypes.$Expr("ExprDo", {keyword: a, statements: a$1})))
  ))
);
const parseCase = (state1, more, resume, done) => tokKeyword("case")(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => separated(tokComma)(parseExpr$lazy())(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => tokKeyword("of")(
      state2$1,
      more,
      resume,
      (state2$2, a$2) => more(v1$2 => {
        const $0 = state2$2.consumed ? {...state2$2, consumed: false} : state2$2;
        return parseBadSingleCaseBranch(
          $0,
          more,
          (state2$3, error) => {
            const $1 = {...state2$3, consumed: $0.consumed};
            if ($1.consumed) { return resume($1, error); }
            return parseCaseBranches$lazy()(
              state2$2,
              more,
              resume,
              (state2$4, a$3) => more(v1$3 => done(state2$4, PureScript$dCST$dTypes.$Expr("ExprCase", {keyword: a, head: a$1, of: a$2, branches: a$3})))
            );
          },
          (state2$3, a$3) => more(v1$3 => done(state2$3, PureScript$dCST$dTypes.$Expr("ExprCase", {keyword: a, head: a$1, of: a$2, branches: a$3})))
        );
      })
    ))
  ))
);
const parseBadSingleCaseWhere = binder => (state1, more, resume, done) => tokRightArrow(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => tokLayoutEnd(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v2 => parseWhere$lazy()(
      state2$1,
      more,
      resume,
      (state3, a$2) => more(v1$1 => done(state3, [Data$dTuple.$Tuple({head: binder, tail: []}, PureScript$dCST$dTypes.$Guarded("Unconditional", a, a$2))]))
    ))
  ))
);
const parseBadSingleCaseGuarded = binder => {
  const $0 = parseGuarded(tokRightArrow);
  return (state1, more, resume, done) => tokLayoutEnd(
    state1,
    more,
    resume,
    (state2, a) => more(v2 => $0(state2, more, resume, (state3, a$1) => more(v1 => done(state3, [Data$dTuple.$Tuple({head: binder, tail: []}, a$1)]))))
  );
};
const parseBadSingleCaseBranch = (state1, more, resume, done) => tokLayoutStart(
  state1,
  more,
  resume,
  (state2, a) => more(v2 => parseBinder1(
    state2,
    more,
    resume,
    (state3, a$1) => more(v1 => {
      const $0 = parseBadSingleCaseGuarded(a$1);
      return parseBadSingleCaseWhere(a$1)(
        state3.consumed ? {...state3, consumed: false} : state3,
        more,
        (state3$1, error) => {
          if (state3$1.consumed) { return resume(state3$1, error); }
          return $0(state3, more, resume, done);
        },
        done
      );
    })
  ))
);
const parseAdo = (state1, more, resume, done) => tokQualifiedKeyword("ado")(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => tokLayoutStart(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => {
      const $0 = PureScript$dCST$dParser$dMonad.many((state1$1, more$1, resume$1, done$1) => tokLayoutSep(
        state1$1,
        more$1,
        resume$1,
        (state2$2, a$2) => more$1(v2 => recoverIndent(PureScript$dCST$dTypes.DoError)(parseDoStatement$lazy())(state2$2, more$1, resume$1, (state3, a$3) => done$1(state3, a$3)))
      ));
      const $1 = (state2$2, a$2) => more(v1$2 => tokKeyword("in")(
        state2$2,
        more,
        resume,
        (state2$3, a$3) => more(v1$3 => parseExpr$lazy()(
          state2$3,
          more,
          resume,
          (state2$4, a$4) => more(v1$4 => done(state2$4, PureScript$dCST$dTypes.$Expr("ExprAdo", {keyword: a, statements: a$2, in: a$3, result: a$4})))
        ))
      ));
      return tokLayoutEnd(
        state2$1.consumed ? {...state2$1, consumed: false} : state2$1,
        more,
        (state3, error) => {
          if (state3.consumed) { return resume(state3, error); }
          return recoverIndent(PureScript$dCST$dTypes.DoError)(parseDoStatement$lazy())(
            state2$1,
            more,
            resume,
            (state2$2, a$2) => more(v2 => $0(
              state2$2,
              more,
              resume,
              (state3$1, a$3) => {
                const $2 = [a$2, ...a$3];
                return more(v2$1 => tokLayoutEnd(state3$1, more, resume, (state3$2, a$4) => $1(state3$2, $2)));
              }
            ))
          );
        },
        (state2$2, a$2) => $1(state2$2, [])
      );
    })
  ))
);
const parseWhere$lazy = /* #__PURE__ */ $runtime.binding(() => PureScript$dCST$dParser$dMonad.lazyParser.defer(v => (state1, more, resume, done) => parseExpr$lazy()(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => {
    const $0 = layoutNonEmpty(recoverIndent(PureScript$dCST$dTypes.LetBindingError)(parseLetBinding$lazy()));
    return tokKeyword("where")(
      state2.consumed ? {...state2, consumed: false} : state2,
      more,
      (state3, error) => {
        if (state3.consumed) { return resume(state3, error); }
        return more(v1$1 => done(state2, {expr: a, bindings: Data$dMaybe.Nothing}));
      },
      (state2$1, a$1) => {
        const $1 = Data$dTuple.Tuple(a$1);
        return more(v2 => $0(
          state2$1,
          more,
          (state3, error) => {
            if (state3.consumed) { return resume(state3, error); }
            return more(v1$1 => done(state2, {expr: a, bindings: Data$dMaybe.Nothing}));
          },
          (state3, a$2) => more(v1$1 => done(state3, {expr: a, bindings: Data$dMaybe.$Maybe("Just", $1(a$2))}))
        ));
      }
    );
  })
)));
const parseTickExpr1$lazy = /* #__PURE__ */ $runtime.binding(() => PureScript$dCST$dParser$dMonad.lazyParser.defer(v => (state1, more, resume, done) => parseExpr3$lazy()(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => PureScript$dCST$dParser$dMonad.many((state1$1, more$1, resume$1, done$1) => parseQualifiedOperator(
    state1$1,
    more$1,
    resume$1,
    (state2$1, a$1) => {
      const $0 = Data$dTuple.Tuple(a$1);
      return more$1(v2 => parseExpr3$lazy()(state2$1, more$1, resume$1, (state3, a$2) => done$1(state3, $0(a$2))));
    }
  ))(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => {
      if (a$1.length > 0) { return done(state2$1, PureScript$dCST$dTypes.$Expr("ExprOp", a, a$1)); }
      return done(state2$1, a);
    })
  ))
)));
const parseLetBinding$lazy = /* #__PURE__ */ $runtime.binding(() => PureScript$dCST$dParser$dMonad.lazyParser.defer(v => (state1, more, resume, done) => {
  const $0 = state1.consumed ? {...state1, consumed: false} : state1;
  return parseIdentBinding(
    $0,
    more,
    (state2, error) => {
      const $1 = {...state2, consumed: $0.consumed};
      if ($1.consumed) { return resume($1, error); }
      return parseBinder1(
        state1,
        more,
        resume,
        (state2$1, a) => {
          const $2 = PureScript$dCST$dTypes.LetBindingPattern(a);
          return more(v2 => tokEquals(
            state2$1,
            more,
            resume,
            (state3, a$1) => {
              const $3 = $2(a$1);
              return more(v2$1 => parseWhere$lazy()(state3, more, resume, (state3$1, a$2) => done(state3$1, $3(a$2))));
            }
          ));
        }
      );
    },
    done
  );
}));
const parseExprAtom$lazy = /* #__PURE__ */ $runtime.binding(() => PureScript$dCST$dParser$dMonad.lazyParser.defer(v => (state1, more, resume, done) => parseQualifiedIdent(
  state1.consumed ? {...state1, consumed: false} : state1,
  more,
  (state3, error) => {
    if (state3.consumed) { return resume(state3, error); }
    return parseQualifiedProper(
      state1.consumed ? {...state1, consumed: false} : state1,
      more,
      (state3$1, error$1) => {
        if (state3$1.consumed) { return resume(state3$1, error$1); }
        return parseQualifiedSymbol(
          state1.consumed ? {...state1, consumed: false} : state1,
          more,
          (state3$2, error$2) => {
            if (state3$2.consumed) { return resume(state3$2, error$2); }
            return tokUnderscore(
              state1.consumed ? {...state1, consumed: false} : state1,
              more,
              (state3$3, error$3) => {
                if (state3$3.consumed) { return resume(state3$3, error$3); }
                return parseHole(
                  state1.consumed ? {...state1, consumed: false} : state1,
                  more,
                  (state3$4, error$4) => {
                    if (state3$4.consumed) { return resume(state3$4, error$4); }
                    return parseString(
                      state1.consumed ? {...state1, consumed: false} : state1,
                      more,
                      (state3$5, error$5) => {
                        if (state3$5.consumed) { return resume(state3$5, error$5); }
                        return parseChar(
                          state1.consumed ? {...state1, consumed: false} : state1,
                          more,
                          (state3$6, error$6) => {
                            if (state3$6.consumed) { return resume(state3$6, error$6); }
                            return parseBoolean(
                              state1.consumed ? {...state1, consumed: false} : state1,
                              more,
                              (state3$7, error$7) => {
                                if (state3$7.consumed) { return resume(state3$7, error$7); }
                                return parseInt(
                                  state1.consumed ? {...state1, consumed: false} : state1,
                                  more,
                                  (state3$8, error$8) => {
                                    if (state3$8.consumed) { return resume(state3$8, error$8); }
                                    return parseNumber(
                                      state1.consumed ? {...state1, consumed: false} : state1,
                                      more,
                                      (state3$9, error$9) => {
                                        if (state3$9.consumed) { return resume(state3$9, error$9); }
                                        return delimited(tokLeftSquare)(tokRightSquare)(tokComma)(parseExpr$lazy())(
                                          state1.consumed ? {...state1, consumed: false} : state1,
                                          more,
                                          (state3$10, error$10) => {
                                            if (state3$10.consumed) { return resume(state3$10, error$10); }
                                            return delimited(tokLeftBrace)(tokRightBrace)(tokComma)(parseRecordLabeled(parseExpr$lazy()))(
                                              state1.consumed ? {...state1, consumed: false} : state1,
                                              more,
                                              (state3$11, error$11) => {
                                                if (state3$11.consumed) { return resume(state3$11, error$11); }
                                                return parens(parseExpr$lazy())(state1, more, resume, (state2, a) => done(state2, PureScript$dCST$dTypes.$Expr("ExprParens", a)));
                                              },
                                              (state2, a) => done(state2, PureScript$dCST$dTypes.$Expr("ExprRecord", a))
                                            );
                                          },
                                          (state2, a) => done(state2, PureScript$dCST$dTypes.$Expr("ExprArray", a))
                                        );
                                      },
                                      (state2, a) => done(state2, PureScript$dCST$dTypes.$Expr("ExprNumber", a._1, a._2))
                                    );
                                  },
                                  (state2, a) => done(state2, PureScript$dCST$dTypes.$Expr("ExprInt", a._1, a._2))
                                );
                              },
                              (state2, a) => done(state2, PureScript$dCST$dTypes.$Expr("ExprBoolean", a._1, a._2))
                            );
                          },
                          (state2, a) => done(state2, PureScript$dCST$dTypes.$Expr("ExprChar", a._1, a._2))
                        );
                      },
                      (state2, a) => done(state2, PureScript$dCST$dTypes.$Expr("ExprString", a._1, a._2))
                    );
                  },
                  (state2, a) => done(state2, PureScript$dCST$dTypes.$Expr("ExprHole", a))
                );
              },
              (state2, a) => done(state2, PureScript$dCST$dTypes.$Expr("ExprSection", a))
            );
          },
          (state2, a) => done(state2, PureScript$dCST$dTypes.$Expr("ExprOpName", a))
        );
      },
      (state2, a) => done(state2, PureScript$dCST$dTypes.$Expr("ExprConstructor", a))
    );
  },
  (state2, a) => done(state2, PureScript$dCST$dTypes.$Expr("ExprIdent", a))
)));
const parseExprAppSpine$lazy = /* #__PURE__ */ $runtime.binding(() => PureScript$dCST$dParser$dMonad.lazyParser.defer(v => (state1, more, resume, done) => {
  const $0 = (state3, error) => {
    if (state3.consumed) { return resume(state3, error); }
    return parseExpr5$lazy()(state1, more, resume, (state2, a) => done(state2, PureScript$dCST$dTypes.$AppSpine("AppTerm", a)));
  };
  return tokAt(
    state1.consumed ? {...state1, consumed: false} : state1,
    more,
    $0,
    (state2, a) => {
      const $1 = PureScript$dCST$dTypes.AppType(a);
      return more(v2 => parseTypeAtom(state2, more, $0, (state3, a$1) => done(state3, $1(a$1))));
    }
  );
}));
const parseExpr7$lazy = /* #__PURE__ */ $runtime.binding(() => PureScript$dCST$dParser$dMonad.lazyParser.defer(v => (state1, more, resume, done) => parseExprAtom$lazy()(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => parseRecordAccessor(a)(
    state2.consumed ? {...state2, consumed: false} : state2,
    more,
    (state3, error) => {
      if (state3.consumed) { return resume(state3, error); }
      return done(state2, a);
    },
    done
  ))
)));
const parseExpr6$lazy = /* #__PURE__ */ $runtime.binding(() => PureScript$dCST$dParser$dMonad.lazyParser.defer(v => (state1, more, resume, done) => parseExpr7$lazy()(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => parseRecordUpdates(a)(
    state2.consumed ? {...state2, consumed: false} : state2,
    more,
    (state3, error) => {
      if (state3.consumed) { return resume(state3, error); }
      return done(state2, a);
    },
    done
  ))
)));
const parseExpr5$lazy = /* #__PURE__ */ $runtime.binding(() => PureScript$dCST$dParser$dMonad.lazyParser.defer(v => (state1, more, resume, done) => parseIf(
  state1.consumed ? {...state1, consumed: false} : state1,
  more,
  (state3, error) => {
    if (state3.consumed) { return resume(state3, error); }
    return parseLetIn(
      state1.consumed ? {...state1, consumed: false} : state1,
      more,
      (state3$1, error$1) => {
        if (state3$1.consumed) { return resume(state3$1, error$1); }
        return parseLambda(
          state1.consumed ? {...state1, consumed: false} : state1,
          more,
          (state3$2, error$2) => {
            if (state3$2.consumed) { return resume(state3$2, error$2); }
            return parseCase(
              state1.consumed ? {...state1, consumed: false} : state1,
              more,
              (state3$3, error$3) => {
                if (state3$3.consumed) { return resume(state3$3, error$3); }
                return parseDo(
                  state1.consumed ? {...state1, consumed: false} : state1,
                  more,
                  (state3$4, error$4) => {
                    if (state3$4.consumed) { return resume(state3$4, error$4); }
                    return parseAdo(
                      state1.consumed ? {...state1, consumed: false} : state1,
                      more,
                      (state3$5, error$5) => {
                        if (state3$5.consumed) { return resume(state3$5, error$5); }
                        return parseExpr6$lazy()(state1, more, resume, done);
                      },
                      done
                    );
                  },
                  done
                );
              },
              done
            );
          },
          done
        );
      },
      done
    );
  },
  done
)));
const parseExpr4$lazy = /* #__PURE__ */ $runtime.binding(() => PureScript$dCST$dParser$dMonad.lazyParser.defer(v => (state1, more, resume, done) => parseExpr5$lazy()(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => PureScript$dCST$dParser$dMonad.many(parseExprAppSpine$lazy())(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => {
      if (a$1.length > 0) { return done(state2$1, PureScript$dCST$dTypes.$Expr("ExprApp", a, a$1)); }
      return done(state2$1, a);
    })
  ))
)));
const parseExpr3$lazy = /* #__PURE__ */ $runtime.binding(() => PureScript$dCST$dParser$dMonad.lazyParser.defer(v => (state1, more, resume, done) => tokKeyOperator("-")(
  state1.consumed ? {...state1, consumed: false} : state1,
  more,
  (state3, error) => {
    if (state3.consumed) { return resume(state3, error); }
    return parseExpr4$lazy()(state1, more, resume, done);
  },
  (state2, a) => {
    const $0 = PureScript$dCST$dTypes.ExprNegate(a);
    return more(v2 => parseExpr3$lazy()(
      state2,
      more,
      (state3, error) => {
        if (state3.consumed) { return resume(state3, error); }
        return parseExpr4$lazy()(state1, more, resume, done);
      },
      (state3, a$1) => done(state3, $0(a$1))
    ));
  }
)));
const parseExpr2$lazy = /* #__PURE__ */ $runtime.binding(() => PureScript$dCST$dParser$dMonad.lazyParser.defer(v => (state1, more, resume, done) => parseExpr3$lazy()(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => PureScript$dCST$dParser$dMonad.many((state1$1, more$1, resume$1, done$1) => parseTickExpr(
    state1$1,
    more$1,
    resume$1,
    (state2$1, a$1) => {
      const $0 = Data$dTuple.Tuple(a$1);
      return more$1(v2 => parseExpr3$lazy()(state2$1, more$1, resume$1, (state3, a$2) => done$1(state3, $0(a$2))));
    }
  ))(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => {
      if (a$1.length > 0) { return done(state2$1, PureScript$dCST$dTypes.$Expr("ExprInfix", a, a$1)); }
      return done(state2$1, a);
    })
  ))
)));
const parseExpr1$lazy = /* #__PURE__ */ $runtime.binding(() => PureScript$dCST$dParser$dMonad.lazyParser.defer(v => (state1, more, resume, done) => parseExpr2$lazy()(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => PureScript$dCST$dParser$dMonad.many((state1$1, more$1, resume$1, done$1) => parseQualifiedOperator(
    state1$1,
    more$1,
    resume$1,
    (state2$1, a$1) => {
      const $0 = Data$dTuple.Tuple(a$1);
      return more$1(v2 => parseExpr2$lazy()(state2$1, more$1, resume$1, (state3, a$2) => done$1(state3, $0(a$2))));
    }
  ))(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => {
      if (a$1.length > 0) { return done(state2$1, PureScript$dCST$dTypes.$Expr("ExprOp", a, a$1)); }
      return done(state2$1, a);
    })
  ))
)));
const parseExpr$lazy = /* #__PURE__ */ $runtime.binding(() => PureScript$dCST$dParser$dMonad.lazyParser.defer(v => (state1, more, resume, done) => parseExpr1$lazy()(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => {
    const $0 = PureScript$dCST$dTypes.ExprTyped(a);
    return tokDoubleColon(
      state2.consumed ? {...state2, consumed: false} : state2,
      more,
      (state3, error) => {
        if (state3.consumed) { return resume(state3, error); }
        return done(state2, a);
      },
      (state2$1, a$1) => {
        const $1 = $0(a$1);
        return more(v2 => parseType(
          state2$1,
          more,
          (state3, error) => {
            if (state3.consumed) { return resume(state3, error); }
            return done(state2, a);
          },
          (state3, a$2) => done(state3, $1(a$2))
        ));
      }
    );
  })
)));
const parseDoStatement$lazy = /* #__PURE__ */ $runtime.binding(() => PureScript$dCST$dParser$dMonad.lazyParser.defer(v => {
  const $0 = layoutNonEmpty(recoverIndent(PureScript$dCST$dTypes.LetBindingError)(parseLetBinding$lazy()));
  return (state1, more, resume, done) => tokKeyword("let")(
    state1.consumed ? {...state1, consumed: false} : state1,
    more,
    (state3, error) => {
      if (state3.consumed) { return resume(state3, error); }
      const $1 = (state3$1, error$1) => {
        if (state3$1.consumed) { return resume(state3$1, error$1); }
        return parseExpr$lazy()(state1, more, resume, (state2, a) => done(state2, PureScript$dCST$dTypes.$DoStatement("DoDiscard", a)));
      };
      const $2 = state1.consumed ? {...state1, consumed: false} : state1;
      return parseBinder(
        $2,
        more,
        (state2, error$1) => $1({...state2, consumed: $2.consumed}, error$1),
        (state2, a) => {
          const $3 = Data$dTuple.Tuple(a);
          return more(v2 => tokLeftArrow(
            state2,
            more,
            (state2$1, error$1) => $1({...state2$1, consumed: $2.consumed}, error$1),
            (state3$1, a$1) => {
              const $4 = $3(a$1);
              const $5 = PureScript$dCST$dTypes.DoBind($4._1)($4._2);
              return more(v2$1 => parseExpr$lazy()(state3$1, more, $1, (state3$2, a$2) => done(state3$2, $5(a$2))));
            }
          ));
        }
      );
    },
    (state2, a) => {
      const $1 = PureScript$dCST$dTypes.DoLet(a);
      return more(v2 => $0(
        state2,
        more,
        (state3, error) => {
          if (state3.consumed) { return resume(state3, error); }
          const $2 = (state3$1, error$1) => {
            if (state3$1.consumed) { return resume(state3$1, error$1); }
            return parseExpr$lazy()(state1, more, resume, (state2$1, a$1) => done(state2$1, PureScript$dCST$dTypes.$DoStatement("DoDiscard", a$1)));
          };
          const $3 = state1.consumed ? {...state1, consumed: false} : state1;
          return parseBinder(
            $3,
            more,
            (state2$1, error$1) => $2({...state2$1, consumed: $3.consumed}, error$1),
            (state2$1, a$1) => {
              const $4 = Data$dTuple.Tuple(a$1);
              return more(v2$1 => tokLeftArrow(
                state2$1,
                more,
                (state2$2, error$1) => $2({...state2$2, consumed: $3.consumed}, error$1),
                (state3$1, a$2) => {
                  const $5 = $4(a$2);
                  const $6 = PureScript$dCST$dTypes.DoBind($5._1)($5._2);
                  return more(v2$2 => parseExpr$lazy()(state3$1, more, $2, (state3$2, a$3) => done(state3$2, $6(a$3))));
                }
              ));
            }
          );
        },
        (state3, a$1) => done(state3, $1(a$1))
      ));
    }
  );
}));
const parseCaseBranches$lazy = /* #__PURE__ */ $runtime.binding(() => PureScript$dCST$dParser$dMonad.lazyParser.defer(v => layoutNonEmpty((() => {
  const $0 = separated(tokComma)(parseBinder1);
  const $1 = parseGuarded(tokRightArrow);
  return (state1, more, resume, done) => $0(
    state1,
    more,
    resume,
    (state2, a) => {
      const $2 = Data$dTuple.Tuple(a);
      return more(v2 => $1(state2, more, resume, (state3, a$1) => done(state3, $2(a$1))));
    }
  );
})())));
const parseWhere = /* #__PURE__ */ parseWhere$lazy();
const parseTickExpr1 = /* #__PURE__ */ parseTickExpr1$lazy();
const parseLetBinding = /* #__PURE__ */ parseLetBinding$lazy();
const parseExprAtom = /* #__PURE__ */ parseExprAtom$lazy();
const parseExprAppSpine = /* #__PURE__ */ parseExprAppSpine$lazy();
const parseExpr7 = /* #__PURE__ */ parseExpr7$lazy();
const parseExpr6 = /* #__PURE__ */ parseExpr6$lazy();
const parseExpr5 = /* #__PURE__ */ parseExpr5$lazy();
const parseExpr4 = /* #__PURE__ */ parseExpr4$lazy();
const parseExpr3 = /* #__PURE__ */ parseExpr3$lazy();
const parseExpr2 = /* #__PURE__ */ parseExpr2$lazy();
const parseExpr1 = /* #__PURE__ */ parseExpr1$lazy();
const parseExpr = /* #__PURE__ */ parseExpr$lazy();
const parseDoStatement = /* #__PURE__ */ parseDoStatement$lazy();
const parseCaseBranches = /* #__PURE__ */ parseCaseBranches$lazy();
const parseDeclValue1 = name => (state1, more, resume, done) => PureScript$dCST$dParser$dMonad.many(parseBinderAtom)(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => parseGuarded(tokEquals)(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => done(state2$1, PureScript$dCST$dTypes.$Declaration("DeclValue", {name, binders: a, guarded: a$1})))
  ))
);
const parseDeclValue = (state1, more, resume, done) => parseIdent(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => parseDeclSignature(a)(
    state2.consumed ? {...state2, consumed: false} : state2,
    more,
    (state3, error) => {
      if (state3.consumed) { return resume(state3, error); }
      return parseDeclValue1(a)(state2, more, resume, done);
    },
    done
  ))
);
const parseInstanceBindingName = name => (state1, more, resume, done) => PureScript$dCST$dParser$dMonad.many(parseBinderAtom)(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => parseGuarded(tokEquals)(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => done(state2$1, PureScript$dCST$dTypes.$InstanceBinding("InstanceBindingName", {name, binders: a, guarded: a$1})))
  ))
);
const parseInstanceBinding = (state1, more, resume, done) => parseIdent(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => parseInstanceBindingSignature(a)(
    state2.consumed ? {...state2, consumed: false} : state2,
    more,
    (state3, error) => {
      if (state3.consumed) { return resume(state3, error); }
      return parseInstanceBindingName(a)(state2, more, resume, done);
    },
    done
  ))
);
const parseInstance = (state1, more, resume, done) => tokKeyword("instance")(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => {
    const $0 = name => {
      const $0 = parseClassConstraints(parseType3);
      return (state1$1, more$1, resume$1, done$1) => {
        const $1 = (state2$1, a$1) => more$1(v1$1 => parseQualifiedProper(
          state2$1,
          more$1,
          resume$1,
          (state2$2, a$2) => more$1(v1$2 => PureScript$dCST$dParser$dMonad.many(parseTypeAtom)(
            state2$2,
            more$1,
            resume$1,
            (state2$3, a$3) => more$1(v1$3 => {
              const $1 = layoutNonEmpty(parseInstanceBinding);
              const $2 = (state2$4, a$4) => more$1(v1$4 => done$1(state2$4, {head: {keyword: a, name, constraints: a$1, className: a$2, types: a$3}, body: a$4}));
              return tokKeyword("where")(
                state2$3.consumed ? {...state2$3, consumed: false} : state2$3,
                more$1,
                (state3, error) => {
                  if (state3.consumed) { return resume$1(state3, error); }
                  return $2(state2$3, Data$dMaybe.Nothing);
                },
                (state2$4, a$4) => {
                  const $3 = Data$dTuple.Tuple(a$4);
                  return more$1(v2 => $1(
                    state2$4,
                    more$1,
                    (state3, error) => {
                      if (state3.consumed) { return resume$1(state3, error); }
                      return $2(state2$3, Data$dMaybe.Nothing);
                    },
                    (state3, a$5) => $2(state3, Data$dMaybe.$Maybe("Just", $3(a$5)))
                  ));
                }
              );
            })
          ))
        ));
        const $2 = state1$1.consumed ? {...state1$1, consumed: false} : state1$1;
        return $0(
          $2,
          more$1,
          (state2$1, error) => {
            const $3 = {...state2$1, consumed: $2.consumed};
            if ($3.consumed) { return resume$1($3, error); }
            return $1(state1$1, Data$dMaybe.Nothing);
          },
          (state2$1, a$1) => {
            const $3 = Data$dTuple.Tuple(a$1);
            return more$1(v2 => tokRightFatArrow(
              state2$1,
              more$1,
              (state2$2, error) => {
                const $4 = {...state2$2, consumed: $2.consumed};
                if ($4.consumed) { return resume$1($4, error); }
                return $1(state1$1, Data$dMaybe.Nothing);
              },
              (state3, a$2) => $1(state3, Data$dMaybe.$Maybe("Just", $3(a$2)))
            ));
          }
        );
      };
    };
    return parseInstanceName(
      state2.consumed ? {...state2, consumed: false} : state2,
      more,
      (state3, error) => {
        if (state3.consumed) { return resume(state3, error); }
        return more(v1$1 => $0(Data$dMaybe.Nothing)(state2, more, resume, done));
      },
      (state2$1, a$1) => more(v1$1 => $0(Data$dMaybe.$Maybe("Just", a$1))(state2$1, more, resume, done))
    );
  })
);
const parseDeclInstanceChain = /* #__PURE__ */ (() => {
  const $0 = separated(parseInstanceChainSeparator)(parseInstance);
  return (state1, more, resume, done) => $0(state1, more, resume, (state2, a) => done(state2, PureScript$dCST$dTypes.$Declaration("DeclInstanceChain", a)));
})();
const parseTypeVarBindingPlain = /* #__PURE__ */ parseTypeVarBinding(parseIdent);
const parseDeclClass1 = keyword => {
  const $0 = parseClassConstraints(parseType5);
  return (state1, more, resume, done) => {
    const $1 = (state2, a) => more(v1 => parseProper(
      state2,
      more,
      resume,
      (state2$1, a$1) => more(v1$1 => PureScript$dCST$dParser$dMonad.many(parseTypeVarBindingPlain)(
        state2$1,
        more,
        resume,
        (state2$2, a$2) => more(v1$2 => {
          const $1 = separated(tokComma)(parseFundep);
          const $2 = (state2$3, a$3) => more(v1$3 => {
            const $2 = layoutNonEmpty(parseClassMember);
            const $3 = (state2$4, a$4) => more(v1$4 => done(
              state2$4,
              PureScript$dCST$dTypes.$Declaration("DeclClass", {keyword, super: a, name: a$1, vars: a$2, fundeps: a$3}, a$4)
            ));
            return tokKeyword("where")(
              state2$3.consumed ? {...state2$3, consumed: false} : state2$3,
              more,
              (state3, error) => {
                if (state3.consumed) { return resume(state3, error); }
                return $3(state2$3, Data$dMaybe.Nothing);
              },
              (state2$4, a$4) => {
                const $4 = Data$dTuple.Tuple(a$4);
                return more(v2 => $2(
                  state2$4,
                  more,
                  (state3, error) => {
                    if (state3.consumed) { return resume(state3, error); }
                    return $3(state2$3, Data$dMaybe.Nothing);
                  },
                  (state3, a$5) => $3(state3, Data$dMaybe.$Maybe("Just", $4(a$5)))
                ));
              }
            );
          });
          return tokPipe(
            state2$2.consumed ? {...state2$2, consumed: false} : state2$2,
            more,
            (state3, error) => {
              if (state3.consumed) { return resume(state3, error); }
              return $2(state2$2, Data$dMaybe.Nothing);
            },
            (state2$3, a$3) => {
              const $3 = Data$dTuple.Tuple(a$3);
              return more(v2 => $1(
                state2$3,
                more,
                (state3, error) => {
                  if (state3.consumed) { return resume(state3, error); }
                  return $2(state2$2, Data$dMaybe.Nothing);
                },
                (state3, a$4) => $2(state3, Data$dMaybe.$Maybe("Just", $3(a$4)))
              ));
            }
          );
        })
      ))
    ));
    const $2 = state1.consumed ? {...state1, consumed: false} : state1;
    return $0(
      $2,
      more,
      (state2, error) => {
        const $3 = {...state2, consumed: $2.consumed};
        if ($3.consumed) { return resume($3, error); }
        return $1(state1, Data$dMaybe.Nothing);
      },
      (state2, a) => {
        const $3 = Data$dTuple.Tuple(a);
        return more(v2 => tokLeftFatArrow(
          state2,
          more,
          (state2$1, error) => {
            const $4 = {...state2$1, consumed: $2.consumed};
            if ($4.consumed) { return resume($4, error); }
            return $1(state1, Data$dMaybe.Nothing);
          },
          (state3, a$1) => $1(state3, Data$dMaybe.$Maybe("Just", $3(a$1)))
        ));
      }
    );
  };
};
const parseDeclClass = (state1, more, resume, done) => tokKeyword("class")(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => {
    const $0 = parseDeclClass1(a);
    return parseDeclClassSignature(a)(
      state2.consumed ? {...state2, consumed: false} : state2,
      more,
      (state3, error) => {
        if (state3.consumed) { return resume(state3, error); }
        return $0(state2, more, resume, done);
      },
      done
    );
  })
);
const parseDeclData1 = keyword => name => (state1, more, resume, done) => PureScript$dCST$dParser$dMonad.many(parseTypeVarBindingPlain)(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => {
    const $0 = separated(tokPipe)(parseDataCtor);
    return tokEquals(
      state2.consumed ? {...state2, consumed: false} : state2,
      more,
      (state3, error) => {
        if (state3.consumed) { return resume(state3, error); }
        return more(v1$1 => done(state2, PureScript$dCST$dTypes.$Declaration("DeclData", {keyword, name, vars: a}, Data$dMaybe.Nothing)));
      },
      (state2$1, a$1) => {
        const $1 = Data$dTuple.Tuple(a$1);
        return more(v2 => $0(
          state2$1,
          more,
          (state3, error) => {
            if (state3.consumed) { return resume(state3, error); }
            return more(v1$1 => done(state2, PureScript$dCST$dTypes.$Declaration("DeclData", {keyword, name, vars: a}, Data$dMaybe.Nothing)));
          },
          (state3, a$2) => more(v1$1 => done(state3, PureScript$dCST$dTypes.$Declaration("DeclData", {keyword, name, vars: a}, Data$dMaybe.$Maybe("Just", $1(a$2)))))
        ));
      }
    );
  })
);
const parseDeclData = (state1, more, resume, done) => tokKeyword("data")(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => parseProper(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => parseDeclKindSignature(a)(a$1)(
      state2$1.consumed ? {...state2$1, consumed: false} : state2$1,
      more,
      (state3, error) => {
        if (state3.consumed) { return resume(state3, error); }
        return parseDeclData1(a)(a$1)(state2$1, more, resume, done);
      },
      done
    ))
  ))
);
const parseDeclNewtype1 = keyword => name => (state1, more, resume, done) => PureScript$dCST$dParser$dMonad.many(parseTypeVarBindingPlain)(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => tokEquals(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => parseProper(
      state2$1,
      more,
      resume,
      (state2$2, a$2) => more(v1$2 => parseTypeAtom(
        state2$2,
        more,
        resume,
        (state2$3, a$3) => more(v1$3 => done(state2$3, PureScript$dCST$dTypes.$Declaration("DeclNewtype", {keyword, name, vars: a}, a$1, a$2, a$3)))
      ))
    ))
  ))
);
const parseDeclNewtype = (state1, more, resume, done) => tokKeyword("newtype")(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => parseProper(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => parseDeclKindSignature(a)(a$1)(
      state2$1.consumed ? {...state2$1, consumed: false} : state2$1,
      more,
      (state3, error) => {
        if (state3.consumed) { return resume(state3, error); }
        return parseDeclNewtype1(a)(a$1)(state2$1, more, resume, done);
      },
      done
    ))
  ))
);
const parseDeclType2 = keyword => name => (state1, more, resume, done) => PureScript$dCST$dParser$dMonad.many(parseTypeVarBindingPlain)(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => tokEquals(
    state2,
    more,
    resume,
    (state2$1, a$1) => more(v1$1 => parseType(
      state2$1,
      more,
      resume,
      (state2$2, a$2) => more(v1$2 => done(state2$2, PureScript$dCST$dTypes.$Declaration("DeclType", {keyword, name, vars: a}, a$1, a$2)))
    ))
  ))
);
const parseDeclType1 = keyword => (state1, more, resume, done) => parseProper(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => parseDeclKindSignature(keyword)(a)(
    state2.consumed ? {...state2, consumed: false} : state2,
    more,
    (state3, error) => {
      if (state3.consumed) { return resume(state3, error); }
      return parseDeclType2(keyword)(a)(state2, more, resume, done);
    },
    done
  ))
);
const parseDeclType = (state1, more, resume, done) => tokKeyword("type")(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => parseDeclRole(a)(
    state2.consumed ? {...state2, consumed: false} : state2,
    more,
    (state3, error) => {
      if (state3.consumed) { return resume(state3, error); }
      return parseDeclType1(a)(state2, more, resume, done);
    },
    done
  ))
);
const parseDecl = (state1, more, resume, done) => parseDeclData(
  state1.consumed ? {...state1, consumed: false} : state1,
  more,
  (state3, error) => {
    if (state3.consumed) { return resume(state3, error); }
    return parseDeclNewtype(
      state1.consumed ? {...state1, consumed: false} : state1,
      more,
      (state3$1, error$1) => {
        if (state3$1.consumed) { return resume(state3$1, error$1); }
        return parseDeclType(
          state1.consumed ? {...state1, consumed: false} : state1,
          more,
          (state3$2, error$2) => {
            if (state3$2.consumed) { return resume(state3$2, error$2); }
            return parseDeclClass(
              state1.consumed ? {...state1, consumed: false} : state1,
              more,
              (state3$3, error$3) => {
                if (state3$3.consumed) { return resume(state3$3, error$3); }
                return parseDeclInstanceChain(
                  state1.consumed ? {...state1, consumed: false} : state1,
                  more,
                  (state3$4, error$4) => {
                    if (state3$4.consumed) { return resume(state3$4, error$4); }
                    return parseDeclDerive(
                      state1.consumed ? {...state1, consumed: false} : state1,
                      more,
                      (state3$5, error$5) => {
                        if (state3$5.consumed) { return resume(state3$5, error$5); }
                        return parseDeclValue(
                          state1.consumed ? {...state1, consumed: false} : state1,
                          more,
                          (state3$6, error$6) => {
                            if (state3$6.consumed) { return resume(state3$6, error$6); }
                            return parseDeclForeign(
                              state1.consumed ? {...state1, consumed: false} : state1,
                              more,
                              (state3$7, error$7) => {
                                if (state3$7.consumed) { return resume(state3$7, error$7); }
                                return parseDeclFixity(state1, more, resume, done);
                              },
                              done
                            );
                          },
                          done
                        );
                      },
                      done
                    );
                  },
                  done
                );
              },
              done
            );
          },
          done
        );
      },
      done
    );
  },
  done
);
const parseModuleDecls = /* #__PURE__ */ PureScript$dCST$dParser$dMonad.many((state1, more, resume, done) => recoverIndent(PureScript$dCST$dTypes.DeclError)(parseDecl)(
  state1,
  more,
  resume,
  (state2, a) => more(v2 => tokLayoutSep(
    state2.consumed ? {...state2, consumed: false} : state2,
    more,
    (state3, error) => {
      if (state3.consumed) { return resume(state3, error); }
      return tokLayoutEnd(state2, more, (v1, error$1) => resume(state2, error$1), (v1, value) => done(state2, a));
    },
    (state3, a$1) => done(state3, a)
  ))
));
const parseModuleBody = (state1, more, resume, done) => parseModuleDecls(
  state1,
  more,
  resume,
  (state2, a) => more(v2 => tokLayoutEnd(
    state2,
    more,
    resume,
    (state3, a$1) => more(v1 => PureScript$dCST$dParser$dMonad.eof(
      state3,
      more,
      resume,
      (state2$1, a$2) => more(v1$1 => done(state2$1, {decls: a, trailingComments: a$2._2, end: a$2._1}))
    ))
  ))
);
const parseModule = (state1, more, resume, done) => parseModuleHeader(
  state1,
  more,
  resume,
  (state2, a) => more(v1 => parseModuleBody(state2, more, resume, (state2$1, a$1) => more(v1$1 => done(state2$1, {header: a, body: a$1}))))
);
export {
  braces,
  delimited,
  expect,
  expectMap,
  layoutNonEmpty,
  many1,
  member,
  parens,
  parseAdo,
  parseBadSingleCaseBranch,
  parseBadSingleCaseGuarded,
  parseBadSingleCaseWhere,
  parseBinder,
  parseBinder1,
  parseBinder2,
  parseBinderAtom,
  parseBinderConstructor,
  parseBinderNegative,
  parseBoolean,
  parseCase,
  parseCaseBranches,
  parseChar,
  parseClassConstraints,
  parseClassMember,
  parseDataCtor,
  parseDataMembers,
  parseDecl,
  parseDeclClass,
  parseDeclClass1,
  parseDeclClassSignature,
  parseDeclData,
  parseDeclData1,
  parseDeclDerive,
  parseDeclFixity,
  parseDeclForeign,
  parseDeclInstanceChain,
  parseDeclKindSignature,
  parseDeclNewtype,
  parseDeclNewtype1,
  parseDeclRole,
  parseDeclSignature,
  parseDeclType,
  parseDeclType1,
  parseDeclType2,
  parseDeclValue,
  parseDeclValue1,
  parseDo,
  parseDoStatement,
  parseEmptyRow,
  parseExport,
  parseExpr,
  parseExpr1,
  parseExpr2,
  parseExpr3,
  parseExpr4,
  parseExpr5,
  parseExpr6,
  parseExpr7,
  parseExprAppSpine,
  parseExprAtom,
  parseFixityKeyword,
  parseFixityOp,
  parseForall,
  parseForeignData,
  parseForeignKind,
  parseForeignValue,
  parseFundep,
  parseGuarded,
  parseHole,
  parseIdent,
  parseIdentBinder,
  parseIdentBinding,
  parseIf,
  parseImport,
  parseImportDecl,
  parseInstance,
  parseInstanceBinding,
  parseInstanceBindingName,
  parseInstanceBindingSignature,
  parseInstanceChainSeparator,
  parseInstanceName,
  parseInt,
  parseKindedVar,
  parseLabel,
  parseLambda,
  parseLetBinding,
  parseLetBindingName,
  parseLetBindingSignature,
  parseLetIn,
  parseModule,
  parseModuleBody,
  parseModuleDecls,
  parseModuleHeader,
  parseModuleImportDecls,
  parseModuleName,
  parseNumber,
  parseOperator,
  parseProper,
  parseQualifiedIdent,
  parseQualifiedIdentOrProper,
  parseQualifiedOperator,
  parseQualifiedProper,
  parseQualifiedSymbol,
  parseRecordAccessor,
  parseRecordLabeled,
  parseRecordUpdate,
  parseRecordUpdateBranch,
  parseRecordUpdateLeaf,
  parseRecordUpdates,
  parseRole,
  parseRow,
  parseRowLabel,
  parseRowParen,
  parseRowTailParen,
  parseSmallInt,
  parseString,
  parseSymbol,
  parseTickExpr,
  parseTickExpr1,
  parseType,
  parseType1,
  parseType2,
  parseType3,
  parseType4,
  parseType5,
  parseTypeAtom,
  parseTypeNegative,
  parseTypeParen,
  parseTypeParens,
  parseTypeVarBinding,
  parseTypeVarBindingPlain,
  parseTypeVarBindingWithVisibility,
  parseTypeVarKinded,
  parseWhere,
  recoverIndent,
  recoverTokensWhile,
  reservedKeywords,
  separated,
  tokAt,
  tokBackslash,
  tokComma,
  tokDot,
  tokDoubleColon,
  tokEquals,
  tokForall,
  tokKeyOperator,
  tokKeySymbol,
  tokKeyword,
  tokLayoutEnd,
  tokLayoutSep,
  tokLayoutStart,
  tokLeftArrow,
  tokLeftBrace,
  tokLeftFatArrow,
  tokLeftParen,
  tokLeftSquare,
  tokPipe,
  tokQualifiedKeyword,
  tokRightArrow,
  tokRightBrace,
  tokRightFatArrow,
  tokRightParen,
  tokRightSquare,
  tokSymbolArrow,
  tokTick,
  tokUnderscore
};
