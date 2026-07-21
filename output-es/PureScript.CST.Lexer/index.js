import * as $runtime from "../runtime.js";
import * as Data$dEither from "../Data.Either/index.js";
import * as Data$dEnum from "../Data.Enum/index.js";
import * as Data$dFoldable from "../Data.Foldable/index.js";
import * as Data$dInt from "../Data.Int/index.js";
import * as Data$dLazy from "../Data.Lazy/index.js";
import * as Data$dList$dTypes from "../Data.List.Types/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dMonoid from "../Data.Monoid/index.js";
import * as Data$dNumber from "../Data.Number/index.js";
import * as Data$dString$dCodePoints from "../Data.String.CodePoints/index.js";
import * as Data$dString$dCodeUnits from "../Data.String.CodeUnits/index.js";
import * as Data$dString$dCommon from "../Data.String.Common/index.js";
import * as Data$dString$dRegex from "../Data.String.Regex/index.js";
import * as Data$dString$dRegex$dFlags from "../Data.String.Regex.Flags/index.js";
import * as Data$dString$dRegex$dUnsafe from "../Data.String.Regex.Unsafe/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
import * as Partial from "../Partial/index.js";
import * as PureScript$dCST$dErrors from "../PureScript.CST.Errors/index.js";
import * as PureScript$dCST$dLayout from "../PureScript.CST.Layout/index.js";
import * as PureScript$dCST$dTokenStream from "../PureScript.CST.TokenStream/index.js";
import * as PureScript$dCST$dTypes from "../PureScript.CST.Types/index.js";
const $LexResult = (tag, _1, _2) => ({tag, _1, _2});
const foldMap = /* #__PURE__ */ (() => Data$dFoldable.foldableArray.foldMap(Data$dMonoid.monoidArray))();
const fold1 = /* #__PURE__ */ (() => Data$dFoldable.foldableArray.foldMap((() => {
  const semigroupRecord1 = {append: ra => rb => ({raw: ra.raw + rb.raw, string: ra.string + rb.string})};
  return {mempty: {raw: "", string: ""}, Semigroup0: () => semigroupRecord1};
})())(Data$dFoldable.identity))();
const consTokens = /* #__PURE__ */ PureScript$dCST$dTokenStream.consTokens(Data$dFoldable.foldableArray);
const LexFail = value0 => value1 => $LexResult("LexFail", value0, value1);
const LexSucc = value0 => value1 => $LexResult("LexSucc", value0, value1);
const isCharCodePoint = /* #__PURE__ */ (() => ({fromChar: Data$dString$dCodePoints.codePointFromChar, fromCharCode: Data$dString$dCodePoints.boundedEnumCodePoint.toEnum}))();
const isCharChar = {fromChar: x => x, fromCharCode: Data$dEnum.charToEnum};
const toModuleName = v => {
  if (v === "") { return Data$dMaybe.Nothing; }
  return Data$dMaybe.$Maybe("Just", Data$dString$dCodeUnits.take(Data$dString$dCodeUnits.length(v) - 1 | 0)(v));
};
const optional = v => str => {
  const v1 = v(str);
  if (v1.tag === "LexFail") {
    if (Data$dString$dCodeUnits.length(str) === Data$dString$dCodeUnits.length(v1._2)) { return $LexResult("LexSucc", Data$dMaybe.Nothing, str); }
    return $LexResult("LexFail", v1._1, v1._2);
  }
  if (v1.tag === "LexSucc") { return $LexResult("LexSucc", Data$dMaybe.$Maybe("Just", v1._1), v1._2); }
  $runtime.fail();
};
const mkUnexpected = str => {
  const start = Data$dString$dCodePoints.take(6)(str);
  const len = Data$dString$dCodePoints.toCodePointArray(start).length;
  if (len === 0) { return "end of file"; }
  if (len < 6) { return start; }
  return start + "...";
};
const regex = mkErr => regexStr => {
  const matchRegex = Data$dString$dRegex$dUnsafe.unsafeRegex("^(?:" + regexStr + ")")(Data$dString$dRegex$dFlags.unicode);
  return str => {
    const v = Data$dString$dRegex.match(matchRegex)(str);
    if (v.tag === "Just") {
      const $0 = (() => {
        if (0 < v._1.length) { return v._1[0]; }
        $runtime.fail();
      })();
      if ($0.tag === "Just") { return $LexResult("LexSucc", $0._1, Data$dString$dCodeUnits.drop(Data$dString$dCodeUnits.length($0._1))(str)); }
    }
    return $LexResult("LexFail", v3 => mkErr(mkUnexpected(str)), str);
  };
};
const shebangComment = /* #__PURE__ */ regex(/* #__PURE__ */ PureScript$dCST$dErrors.LexExpected("shebang"))("#![^\\r\\n]*");
const satisfy = mkErr => p => str => {
  const v = Data$dString$dCodeUnits.charAt(0)(str);
  if (v.tag === "Just" && p(v._1)) { return $LexResult("LexSucc", v._1, Data$dString$dCodeUnits.drop(1)(str)); }
  return $LexResult("LexFail", v1 => mkErr(mkUnexpected(str)), str);
};
const string = mkErr => match => str => {
  if (Data$dString$dCodeUnits.take(Data$dString$dCodeUnits.length(match))(str) === match) {
    return $LexResult("LexSucc", match, Data$dString$dCodeUnits.drop(Data$dString$dCodeUnits.length(match))(str));
  }
  return $LexResult("LexFail", v => mkErr(mkUnexpected(str)), str);
};
const many = v => str => {
  const valuesRef = [];
  let strRef = str;
  let contRef = true;
  let resRef = $LexResult("LexSucc", [], str);
  while (contRef) {
    const str$p = strRef;
    const v1 = v(str$p);
    if (v1.tag === "LexFail") {
      if (Data$dString$dCodeUnits.length(str$p) === Data$dString$dCodeUnits.length(v1._2)) {
        resRef = $LexResult("LexSucc", valuesRef, v1._2);
        contRef = false;
        continue;
      }
      resRef = $LexResult("LexFail", v1._1, v1._2);
      contRef = false;
      continue;
    }
    if (v1.tag === "LexSucc") {
      valuesRef.push(v1._1);
      strRef = v1._2;
      continue;
    }
    $runtime.fail();
  }
  return resRef;
};
const functorLex = {
  map: f => v => str => {
    const v1 = v(str);
    if (v1.tag === "LexFail") { return $LexResult("LexFail", v1._1, v1._2); }
    if (v1.tag === "LexSucc") { return $LexResult("LexSucc", f(v1._1), v1._2); }
    $runtime.fail();
  }
};
const spaceComment = /* #__PURE__ */ (() => {
  const $0 = regex(PureScript$dCST$dErrors.LexExpected("spaces"))(" +");
  return str => {
    const v1 = $0(str);
    if (v1.tag === "LexFail") { return $LexResult("LexFail", v1._1, v1._2); }
    if (v1.tag === "LexSucc") { return $LexResult("LexSucc", Data$dString$dCodeUnits.length(v1._1), v1._2); }
    $runtime.fail();
  };
})();
const char$p = mkErr => res => match => str => {
  if (Data$dString$dCodeUnits.singleton(match) === Data$dString$dCodeUnits.take(1)(str)) { return $LexResult("LexSucc", res, Data$dString$dCodeUnits.drop(1)(str)); }
  return $LexResult("LexFail", v => mkErr(mkUnexpected(str)), str);
};
const $$char = mkErr => match => str => {
  if (Data$dString$dCodeUnits.singleton(match) === Data$dString$dCodeUnits.take(1)(str)) { return $LexResult("LexSucc", match, Data$dString$dCodeUnits.drop(1)(str)); }
  return $LexResult("LexFail", v => mkErr(mkUnexpected(str)), str);
};
const bumpText = v => colOffset => str => {
  const $0 = v.column;
  const $1 = v.line;
  const go = go$a0$copy => go$a1$copy => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
    while (go$c) {
      const n = go$a0, ix = go$a1;
      const v1 = Data$dString$dCodeUnits.indexOf$p("\n")(ix)(str);
      if (v1.tag === "Just") {
        go$a0 = n + 1 | 0;
        go$a1 = v1._1 + 1 | 0;
        continue;
      }
      if (v1.tag === "Nothing") {
        if (n === 0) {
          go$c = false;
          go$r = {line: $1, column: ($0 + Data$dString$dCodePoints.toCodePointArray(str).length | 0) + (colOffset * 2 | 0) | 0};
          continue;
        }
        go$c = false;
        go$r = {line: $1 + n | 0, column: Data$dString$dCodePoints.toCodePointArray(Data$dString$dCodeUnits.drop(ix)(str)).length + colOffset | 0};
        continue;
      }
      $runtime.fail();
    }
    return go$r;
  };
  return go(0)(0);
};
const bumpToken = v => {
  const $0 = v.column;
  const $1 = v.line;
  return v1 => {
    if (v1.tag === "TokLeftParen") { return {line: $1, column: $0 + 1 | 0}; }
    if (v1.tag === "TokRightParen") { return {line: $1, column: $0 + 1 | 0}; }
    if (v1.tag === "TokLeftBrace") { return {line: $1, column: $0 + 1 | 0}; }
    if (v1.tag === "TokRightBrace") { return {line: $1, column: $0 + 1 | 0}; }
    if (v1.tag === "TokLeftSquare") { return {line: $1, column: $0 + 1 | 0}; }
    if (v1.tag === "TokRightSquare") { return {line: $1, column: $0 + 1 | 0}; }
    if (v1.tag === "TokLeftArrow") {
      if (v1._1 === "ASCII") { return {line: $1, column: $0 + 2 | 0}; }
      if (v1._1 === "Unicode") { return {line: $1, column: $0 + 1 | 0}; }
      $runtime.fail();
    }
    if (v1.tag === "TokRightArrow") {
      if (v1._1 === "ASCII") { return {line: $1, column: $0 + 2 | 0}; }
      if (v1._1 === "Unicode") { return {line: $1, column: $0 + 1 | 0}; }
      $runtime.fail();
    }
    if (v1.tag === "TokRightFatArrow") {
      if (v1._1 === "ASCII") { return {line: $1, column: $0 + 2 | 0}; }
      if (v1._1 === "Unicode") { return {line: $1, column: $0 + 1 | 0}; }
      $runtime.fail();
    }
    if (v1.tag === "TokDoubleColon") {
      if (v1._1 === "ASCII") { return {line: $1, column: $0 + 2 | 0}; }
      if (v1._1 === "Unicode") { return {line: $1, column: $0 + 1 | 0}; }
      $runtime.fail();
    }
    if (v1.tag === "TokForall") {
      if (v1._1 === "ASCII") { return {line: $1, column: $0 + 6 | 0}; }
      if (v1._1 === "Unicode") { return {line: $1, column: $0 + 1 | 0}; }
      $runtime.fail();
    }
    if (v1.tag === "TokEquals") { return {line: $1, column: $0 + 1 | 0}; }
    if (v1.tag === "TokPipe") { return {line: $1, column: $0 + 1 | 0}; }
    if (v1.tag === "TokTick") { return {line: $1, column: $0 + 1 | 0}; }
    if (v1.tag === "TokDot") { return {line: $1, column: $0 + 1 | 0}; }
    if (v1.tag === "TokComma") { return {line: $1, column: $0 + 1 | 0}; }
    if (v1.tag === "TokUnderscore") { return {line: $1, column: $0 + 1 | 0}; }
    if (v1.tag === "TokBackslash") { return {line: $1, column: $0 + 1 | 0}; }
    if (v1.tag === "TokAt") { return {line: $1, column: $0 + 1 | 0}; }
    if (v1.tag === "TokLowerName") {
      return {
        line: $1,
        column: (() => {
          if (v1._1.tag === "Nothing") { return $0 + 0 | 0; }
          if (v1._1.tag === "Just") { return ($0 + 1 | 0) + Data$dString$dCodePoints.toCodePointArray(v1._1._1).length | 0; }
          $runtime.fail();
        })() + Data$dString$dCodePoints.toCodePointArray(v1._2).length | 0
      };
    }
    if (v1.tag === "TokUpperName") {
      return {
        line: $1,
        column: (() => {
          if (v1._1.tag === "Nothing") { return $0 + 0 | 0; }
          if (v1._1.tag === "Just") { return ($0 + 1 | 0) + Data$dString$dCodePoints.toCodePointArray(v1._1._1).length | 0; }
          $runtime.fail();
        })() + Data$dString$dCodePoints.toCodePointArray(v1._2).length | 0
      };
    }
    if (v1.tag === "TokOperator") {
      return {
        line: $1,
        column: (() => {
          if (v1._1.tag === "Nothing") { return $0 + 0 | 0; }
          if (v1._1.tag === "Just") { return ($0 + 1 | 0) + Data$dString$dCodePoints.toCodePointArray(v1._1._1).length | 0; }
          $runtime.fail();
        })() + Data$dString$dCodePoints.toCodePointArray(v1._2).length | 0
      };
    }
    if (v1.tag === "TokSymbolName") {
      return {
        line: $1,
        column: (
          (() => {
            if (v1._1.tag === "Nothing") { return $0 + 0 | 0; }
            if (v1._1.tag === "Just") { return ($0 + 1 | 0) + Data$dString$dCodePoints.toCodePointArray(v1._1._1).length | 0; }
            $runtime.fail();
          })() + Data$dString$dCodePoints.toCodePointArray(v1._2).length | 0
        ) + 2 | 0
      };
    }
    if (v1.tag === "TokSymbolArrow") {
      if (v1._1 === "Unicode") { return {line: $1, column: $0 + 3 | 0}; }
      if (v1._1 === "ASCII") { return {line: $1, column: $0 + 4 | 0}; }
      $runtime.fail();
    }
    if (v1.tag === "TokHole") { return {line: $1, column: ($0 + Data$dString$dCodePoints.toCodePointArray(v1._1).length | 0) + 1 | 0}; }
    if (v1.tag === "TokChar") { return {line: $1, column: ($0 + Data$dString$dCodePoints.toCodePointArray(v1._1).length | 0) + 2 | 0}; }
    if (v1.tag === "TokInt") { return {line: $1, column: $0 + Data$dString$dCodePoints.toCodePointArray(v1._1).length | 0}; }
    if (v1.tag === "TokNumber") { return {line: $1, column: $0 + Data$dString$dCodePoints.toCodePointArray(v1._1).length | 0}; }
    if (v1.tag === "TokString") { return bumpText(v)(1)(v1._1); }
    if (v1.tag === "TokRawString") { return bumpText(v)(3)(v1._1); }
    if (v1.tag === "TokLayoutStart") { return v; }
    if (v1.tag === "TokLayoutSep") { return v; }
    if (v1.tag === "TokLayoutEnd") { return v; }
    $runtime.fail();
  };
};
const bumpComment = v => {
  const $0 = v.column;
  const $1 = v.line;
  return v1 => {
    if (v1.tag === "Comment") { return bumpText(v)(0)(v1._1); }
    if (v1.tag === "Space") { return {line: $1, column: $0 + v1._1 | 0}; }
    if (v1.tag === "Line") { return {line: $1 + v1._2 | 0, column: 0}; }
    $runtime.fail();
  };
};
const altLex = {
  alt: v => v1 => str => {
    const v2 = v(str);
    if (v2.tag === "LexFail") {
      if (Data$dString$dCodeUnits.length(str) === Data$dString$dCodeUnits.length(v2._2)) { return v1(str); }
      return $LexResult("LexFail", v2._1, v2._2);
    }
    if (v2.tag === "LexSucc") { return $LexResult("LexSucc", v2._1, v2._2); }
    $runtime.fail();
  },
  Functor0: () => functorLex
};
const comment = /* #__PURE__ */ (() => altLex.alt(regex(PureScript$dCST$dErrors.LexExpected("block comment"))("\\{-(-(?!\\})|[^-]+)*(-\\}|$)"))(regex(PureScript$dCST$dErrors.LexExpected("line comment"))("--[^\\r\\n]*")))();
const lineComment = /* #__PURE__ */ (() => altLex.alt((() => {
  const $0 = PureScript$dCST$dTypes.Line(PureScript$dCST$dTypes.LF);
  const $1 = regex(PureScript$dCST$dErrors.LexExpected("newline"))("\n+");
  return str => {
    const v1 = $1(str);
    if (v1.tag === "LexFail") { return $LexResult("LexFail", v1._1, v1._2); }
    if (v1.tag === "LexSucc") { return $LexResult("LexSucc", $0(Data$dString$dCodePoints.toCodePointArray(v1._1).length), v1._2); }
    $runtime.fail();
  };
})())((() => {
  const $0 = PureScript$dCST$dTypes.Line(PureScript$dCST$dTypes.CRLF);
  const $1 = regex(PureScript$dCST$dErrors.LexExpected("newline"))("(?:\r\n)+");
  return str => {
    const v1 = $1(str);
    if (v1.tag === "LexFail") { return $LexResult("LexFail", v1._1, v1._2); }
    if (v1.tag === "LexSucc") { return $LexResult("LexSucc", $0($runtime.intDiv(Data$dString$dCodePoints.toCodePointArray(v1._1).length, 2)), v1._2); }
    $runtime.fail();
  };
})()))();
const leadingComments = /* #__PURE__ */ (() => many(altLex.alt(str => {
  const v1 = comment(str);
  if (v1.tag === "LexFail") { return $LexResult("LexFail", v1._1, v1._2); }
  if (v1.tag === "LexSucc") { return $LexResult("LexSucc", PureScript$dCST$dTypes.$Comment("Comment", v1._1), v1._2); }
  $runtime.fail();
})(altLex.alt(str => {
  const v1 = spaceComment(str);
  if (v1.tag === "LexFail") { return $LexResult("LexFail", v1._1, v1._2); }
  if (v1.tag === "LexSucc") { return $LexResult("LexSucc", PureScript$dCST$dTypes.$Comment("Space", v1._1), v1._2); }
  $runtime.fail();
})(lineComment))))();
const oneLineComment = str => {
  const v1 = lineComment(str);
  if (v1.tag === "LexFail") { return $LexResult("LexFail", v1._1, v1._2); }
  if (v1.tag === "LexSucc") {
    if (v1._1.tag === "Line" && v1._1._2 === 1) { return $LexResult("LexSucc", v1._1, v1._2); }
    return $LexResult("LexFail", v => PureScript$dCST$dErrors.$ParseError("LexExpected", "one newline", "multiple newlines"), v1._2);
  }
  $runtime.fail();
};
const leadingShebangs = /* #__PURE__ */ (() => {
  const $0 = many(str => {
    const v1 = oneLineComment(str);
    if (v1.tag === "LexFail") { return $LexResult("LexFail", v1._1, str); }
    if (v1.tag === "LexSucc") {
      const v3 = shebangComment(v1._2);
      if (v3.tag === "LexFail") { return $LexResult("LexFail", v3._1, str); }
      if (v3.tag === "LexSucc") { return $LexResult("LexSucc", Data$dTuple.$Tuple(v1._1, v3._1), v3._2); }
    }
    $runtime.fail();
  });
  return str => {
    const v1 = shebangComment(str);
    if (v1.tag === "LexFail") { return $LexResult("LexFail", v1._1, v1._2); }
    if (v1.tag === "LexSucc") {
      const v3 = $0(v1._2);
      if (v3.tag === "LexFail") { return $LexResult("LexFail", v3._1, v3._2); }
      if (v3.tag === "LexSucc") {
        return $LexResult(
          "LexSucc",
          [PureScript$dCST$dTypes.$Comment("Comment", v1._1), ...foldMap(v2 => [v2._1, PureScript$dCST$dTypes.$Comment("Comment", v2._2)])(v3._1)],
          v3._2
        );
      }
    }
    $runtime.fail();
  };
})();
const leadingModuleComments = /* #__PURE__ */ (() => {
  const $0 = altLex.alt(leadingShebangs)(LexSucc([]));
  return str => {
    const v1 = $0(str);
    if (v1.tag === "LexFail") { return $LexResult("LexFail", v1._1, v1._2); }
    if (v1.tag === "LexSucc") {
      const v3 = leadingComments(v1._2);
      if (v3.tag === "LexFail") { return $LexResult("LexFail", v3._1, v3._2); }
      if (v3.tag === "LexSucc") { return $LexResult("LexSucc", [...v1._1, ...v3._1], v3._2); }
    }
    $runtime.fail();
  };
})();
const token = /* #__PURE__ */ (() => {
  const tokenRightParen = char$p(PureScript$dCST$dErrors.LexExpected("right paren"))(PureScript$dCST$dTypes.TokRightParen)(")");
  const tokenLeftParen = char$p(PureScript$dCST$dErrors.LexExpected("left paren"))(PureScript$dCST$dTypes.TokLeftParen)("(");
  const stripUnderscores = Data$dString$dCommon.replaceAll("_")("");
  const parseSymbolIdent = regex(PureScript$dCST$dErrors.LexExpected("symbol"))("(?:[:!#$%&*+./<=>?@\\\\^|~-]|(?!\\p{P})\\p{S})+");
  const parseProper = regex(PureScript$dCST$dErrors.LexExpected("proper name"))("\\p{Lu}[\\p{L}0-9_']*");
  const parseIdent = regex(PureScript$dCST$dErrors.LexExpected("ident"))("[\\p{Ll}_][\\p{L}0-9_']*");
  const intPartRegex = regex(PureScript$dCST$dErrors.LexExpected("int part"))("(0|[1-9][0-9_]*)");
  const hexEscapeRegex = regex(PureScript$dCST$dErrors.LexExpected("hex"))("[a-fA-F0-9]{1,6}");
  const charSingleQuote = $$char(PureScript$dCST$dErrors.LexExpected("single quote"))("'");
  const charQuote = $$char(PureScript$dCST$dErrors.LexExpected("quote"))("\"");
  const charAny = satisfy(PureScript$dCST$dErrors.LexExpected("char"))(v => true);
  const parseEscape = dictIsChar => str => {
    const v1 = charAny(str);
    if (v1.tag === "LexFail") { return $LexResult("LexFail", v1._1, v1._2); }
    if (v1.tag === "LexSucc") {
      if (v1._1 === "t") { return $LexResult("LexSucc", {raw: "\\t", char: dictIsChar.fromChar("\t")}, v1._2); }
      if (v1._1 === "r") { return $LexResult("LexSucc", {raw: "\\r", char: dictIsChar.fromChar("\r")}, v1._2); }
      if (v1._1 === "n") { return $LexResult("LexSucc", {raw: "\\n", char: dictIsChar.fromChar("\n")}, v1._2); }
      if (v1._1 === "\"") { return $LexResult("LexSucc", {raw: "\\\"", char: dictIsChar.fromChar("\"")}, v1._2); }
      if (v1._1 === "'") { return $LexResult("LexSucc", {raw: "\\'", char: dictIsChar.fromChar("'")}, v1._2); }
      if (v1._1 === "\\") { return $LexResult("LexSucc", {raw: "\\\\", char: dictIsChar.fromChar("\\")}, v1._2); }
      if (v1._1 === "x") {
        const v1$1 = hexEscapeRegex(v1._2);
        if (v1$1.tag === "LexFail") { return $LexResult("LexFail", v1$1._1, v1$1._2); }
        if (v1$1.tag === "LexSucc") {
          const $0 = Data$dInt.fromStringAs(16)(v1$1._1);
          const v = (() => {
            if ($0.tag === "Just") { return dictIsChar.fromCharCode($0._1); }
            if ($0.tag === "Nothing") { return Data$dMaybe.Nothing; }
            $runtime.fail();
          })();
          if (v.tag === "Just") { return $LexResult("LexSucc", {raw: "\\x" + v1$1._1, char: v._1}, v1$1._2); }
          if (v.tag === "Nothing") { return $LexResult("LexFail", v$1 => PureScript$dCST$dErrors.$ParseError("LexCharEscapeOutOfRange", v1$1._1), v1$1._2); }
        }
        $runtime.fail();
      }
      const $0 = PureScript$dCST$dErrors.$ParseError("LexInvalidCharEscape", Data$dString$dCodeUnits.singleton(v1._1));
      return $LexResult("LexFail", v => $0, v1._2);
    }
    $runtime.fail();
  };
  const parseEscape1 = parseEscape(isCharChar);
  return altLex.alt((() => {
    const $0 = $$char(PureScript$dCST$dErrors.LexExpected("question mark"))("?");
    return str => {
      const v1 = $0(str);
      if (v1.tag === "LexFail") { return $LexResult("LexFail", v1._1, str); }
      if (v1.tag === "LexSucc") {
        const v3 = altLex.alt(parseIdent)(parseProper)(v1._2);
        if (v3.tag === "LexFail") { return $LexResult("LexFail", v3._1, str); }
        if (v3.tag === "LexSucc") { return $LexResult("LexSucc", PureScript$dCST$dTypes.$Token("TokHole", v3._1), v3._2); }
      }
      $runtime.fail();
    };
  })())(altLex.alt((() => {
    const $0 = regex(PureScript$dCST$dErrors.LexExpected("module name"))("(?:(?:\\p{Lu}[\\p{L}0-9_']*)\\.)*");
    const $1 = altLex.alt(str => {
      const v1 = parseIdent(str);
      if (v1.tag === "LexFail") { return $LexResult("LexFail", v1._1, v1._2); }
      if (v1.tag === "LexSucc") {
        return $LexResult(
          "LexSucc",
          (() => {
            const $1 = v1._1;
            return v1$1 => {
              if (v1$1.tag === "Nothing") {
                if ($1 === "forall") { return PureScript$dCST$dTypes.$Token("TokForall", PureScript$dCST$dTypes.ASCII); }
                if ($1 === "_") { return PureScript$dCST$dTypes.TokUnderscore; }
                return PureScript$dCST$dTypes.$Token("TokLowerName", Data$dMaybe.Nothing, $1);
              }
              return PureScript$dCST$dTypes.$Token("TokLowerName", v1$1, $1);
            };
          })(),
          v1._2
        );
      }
      $runtime.fail();
    })(altLex.alt(str => {
      const v1 = parseProper(str);
      if (v1.tag === "LexFail") { return $LexResult("LexFail", v1._1, v1._2); }
      if (v1.tag === "LexSucc") {
        return $LexResult(
          "LexSucc",
          (() => {
            const $1 = v1._1;
            return a => PureScript$dCST$dTypes.$Token("TokUpperName", a, $1);
          })(),
          v1._2
        );
      }
      $runtime.fail();
    })(altLex.alt(str => {
      const v1 = parseSymbolIdent(str);
      if (v1.tag === "LexFail") { return $LexResult("LexFail", v1._1, v1._2); }
      if (v1.tag === "LexSucc") {
        return $LexResult(
          "LexSucc",
          (() => {
            const $1 = v1._1;
            return v1$1 => {
              if (v1$1.tag === "Nothing") {
                if ($1 === "<-") { return PureScript$dCST$dTypes.$Token("TokLeftArrow", PureScript$dCST$dTypes.ASCII); }
                if ($1 === "←") { return PureScript$dCST$dTypes.$Token("TokLeftArrow", PureScript$dCST$dTypes.Unicode); }
                if ($1 === "->") { return PureScript$dCST$dTypes.$Token("TokRightArrow", PureScript$dCST$dTypes.ASCII); }
                if ($1 === "→") { return PureScript$dCST$dTypes.$Token("TokRightArrow", PureScript$dCST$dTypes.Unicode); }
                if ($1 === "=>") { return PureScript$dCST$dTypes.$Token("TokRightFatArrow", PureScript$dCST$dTypes.ASCII); }
                if ($1 === "⇒") { return PureScript$dCST$dTypes.$Token("TokRightFatArrow", PureScript$dCST$dTypes.Unicode); }
                if ($1 === "::") { return PureScript$dCST$dTypes.$Token("TokDoubleColon", PureScript$dCST$dTypes.ASCII); }
                if ($1 === "∷") { return PureScript$dCST$dTypes.$Token("TokDoubleColon", PureScript$dCST$dTypes.Unicode); }
                if ($1 === "∀") { return PureScript$dCST$dTypes.$Token("TokForall", PureScript$dCST$dTypes.Unicode); }
                if ($1 === "=") { return PureScript$dCST$dTypes.TokEquals; }
                if ($1 === ".") { return PureScript$dCST$dTypes.TokDot; }
                if ($1 === "\\") { return PureScript$dCST$dTypes.TokBackslash; }
                if ($1 === "|") { return PureScript$dCST$dTypes.TokPipe; }
                if ($1 === "@") { return PureScript$dCST$dTypes.TokAt; }
                if ($1 === "`") { return PureScript$dCST$dTypes.TokTick; }
                return PureScript$dCST$dTypes.$Token("TokOperator", Data$dMaybe.Nothing, $1);
              }
              return PureScript$dCST$dTypes.$Token("TokOperator", v1$1, $1);
            };
          })(),
          v1._2
        );
      }
      $runtime.fail();
    })(str => {
      const v1 = tokenLeftParen(str);
      if (v1.tag === "LexFail") { return $LexResult("LexFail", v1._1, str); }
      if (v1.tag === "LexSucc") {
        const v3 = parseSymbolIdent(v1._2);
        if (v3.tag === "LexFail") { return $LexResult("LexFail", v3._1, str); }
        if (v3.tag === "LexSucc") {
          const v3$1 = tokenRightParen(v3._2);
          if (v3$1.tag === "LexFail") { return $LexResult("LexFail", v3$1._1, str); }
          if (v3$1.tag === "LexSucc") {
            return $LexResult(
              "LexSucc",
              (() => {
                const $1 = v3._1;
                return v1$1 => {
                  if (v1$1.tag === "Nothing") {
                    if ($1 === "->") { return PureScript$dCST$dTypes.$Token("TokSymbolArrow", PureScript$dCST$dTypes.ASCII); }
                    if ($1 === "→") { return PureScript$dCST$dTypes.$Token("TokSymbolArrow", PureScript$dCST$dTypes.Unicode); }
                    return PureScript$dCST$dTypes.$Token("TokSymbolName", Data$dMaybe.Nothing, $1);
                  }
                  return PureScript$dCST$dTypes.$Token("TokSymbolName", v1$1, $1);
                };
              })(),
              v3$1._2
            );
          }
        }
      }
      $runtime.fail();
    })));
    return str => {
      const v1 = $0(str);
      if (v1.tag === "LexFail") { return $LexResult("LexFail", v1._1, v1._2); }
      if (v1.tag === "LexSucc") {
        const v3 = $1(v1._2);
        if (v3.tag === "LexFail") { return $LexResult("LexFail", v3._1, v3._2); }
        if (v3.tag === "LexSucc") { return $LexResult("LexSucc", v3._1(toModuleName(v1._1)), v3._2); }
      }
      $runtime.fail();
    };
  })())(altLex.alt(str => {
    const v1 = charSingleQuote(str);
    if (v1.tag === "LexFail") { return $LexResult("LexFail", v1._1, v1._2); }
    if (v1.tag === "LexSucc") {
      const v1$1 = charAny(v1._2);
      const v3 = (() => {
        if (v1$1.tag === "LexFail") { return $LexResult("LexFail", v1$1._1, v1$1._2); }
        if (v1$1.tag === "LexSucc") {
          if (v1$1._1 === "\\") { return parseEscape1(v1$1._2); }
          if (v1$1._1 === "'") { return $LexResult("LexFail", v => PureScript$dCST$dErrors.$ParseError("LexExpected", "character", "empty character literal"), v1$1._2); }
          return $LexResult("LexSucc", {raw: Data$dString$dCodeUnits.singleton(v1$1._1), char: v1$1._1}, v1$1._2);
        }
        $runtime.fail();
      })();
      if (v3.tag === "LexFail") { return $LexResult("LexFail", v3._1, v3._2); }
      if (v3.tag === "LexSucc") {
        const v3$1 = charSingleQuote(v3._2);
        if (v3$1.tag === "LexFail") { return $LexResult("LexFail", v3$1._1, v3$1._2); }
        if (v3$1.tag === "LexSucc") { return $LexResult("LexSucc", PureScript$dCST$dTypes.$Token("TokChar", v3._1.raw, v3._1.char), v3$1._2); }
      }
    }
    $runtime.fail();
  })(altLex.alt(altLex.alt((() => {
    const $0 = regex(PureScript$dCST$dErrors.LexExpected("raw string characters"))("\"\"\"\"{0,2}([^\"]+\"{1,2})*[^\"]*\"\"\"");
    return str => {
      const v1 = $0(str);
      if (v1.tag === "LexFail") { return $LexResult("LexFail", v1._1, v1._2); }
      if (v1.tag === "LexSucc") {
        return $LexResult(
          "LexSucc",
          PureScript$dCST$dTypes.$Token(
            "TokRawString",
            (() => {
              const $1 = Data$dString$dCodeUnits.drop(3)(v1._1);
              return Data$dString$dCodeUnits.take(Data$dString$dCodeUnits.length($1) - 3 | 0)($1);
            })()
          ),
          v1._2
        );
      }
      $runtime.fail();
    };
  })())((() => {
    const $0 = many(altLex.alt((() => {
      const $0 = regex(PureScript$dCST$dErrors.LexExpected("string characters"))("[^\"\\\\]+");
      return str => {
        const v1 = $0(str);
        if (v1.tag === "LexFail") { return $LexResult("LexFail", v1._1, v1._2); }
        if (v1.tag === "LexSucc") { return $LexResult("LexSucc", {raw: v1._1, string: v1._1}, v1._2); }
        $runtime.fail();
      };
    })())(altLex.alt((() => {
      const $0 = regex(PureScript$dCST$dErrors.LexExpected("whitespace escape"))("\\\\[ \\r\\n]+\\\\");
      return str => {
        const v1 = $0(str);
        if (v1.tag === "LexFail") { return $LexResult("LexFail", v1._1, v1._2); }
        if (v1.tag === "LexSucc") { return $LexResult("LexSucc", {raw: v1._1, string: ""}, v1._2); }
        $runtime.fail();
      };
    })())((() => {
      const $0 = $$char(PureScript$dCST$dErrors.LexExpected("backslash"))("\\");
      const $1 = parseEscape(isCharCodePoint);
      return str => {
        const v1 = $0(str);
        if (v1.tag === "LexFail") { return $LexResult("LexFail", v1._1, v1._2); }
        if (v1.tag === "LexSucc") {
          const v3 = $1(v1._2);
          if (v3.tag === "LexFail") { return $LexResult("LexFail", v3._1, v3._2); }
          if (v3.tag === "LexSucc") { return $LexResult("LexSucc", {raw: v3._1.raw, string: Data$dString$dCodePoints.singleton(v3._1.char)}, v3._2); }
        }
        $runtime.fail();
      };
    })())));
    return str => {
      const v1 = charQuote(str);
      if (v1.tag === "LexFail") { return $LexResult("LexFail", v1._1, v1._2); }
      if (v1.tag === "LexSucc") {
        const v3 = $0(v1._2);
        if (v3.tag === "LexFail") { return $LexResult("LexFail", v3._1, v3._2); }
        if (v3.tag === "LexSucc") {
          const v3$1 = charQuote(v3._2);
          if (v3$1.tag === "LexFail") { return $LexResult("LexFail", v3$1._1, v3$1._2); }
          if (v3$1.tag === "LexSucc") {
            return $LexResult(
              "LexSucc",
              (() => {
                const v1$1 = fold1(v3._1);
                return PureScript$dCST$dTypes.$Token("TokString", v1$1.raw, v1$1.string);
              })(),
              v3$1._2
            );
          }
        }
      }
      $runtime.fail();
    };
  })()))(altLex.alt(altLex.alt((() => {
    const $0 = string(PureScript$dCST$dErrors.LexExpected("hex int prefix"))("0x");
    const $1 = regex(PureScript$dCST$dErrors.LexExpected("hex int"))("[a-fA-F0-9]+");
    return str => {
      const v1 = $0(str);
      if (v1.tag === "LexFail") { return $LexResult("LexFail", v1._1, v1._2); }
      if (v1.tag === "LexSucc") {
        const v3 = $1(v1._2);
        if (v3.tag === "LexFail") { return $LexResult("LexFail", v3._1, v3._2); }
        if (v3.tag === "LexSucc") {
          const v = Data$dInt.fromStringAs(16)(v3._1);
          if (v.tag === "Just") { return $LexResult("LexSucc", PureScript$dCST$dTypes.$Token("TokInt", "0x" + v3._1, PureScript$dCST$dTypes.$IntValue("SmallInt", v._1)), v3._2); }
          if (v.tag === "Nothing") {
            return $LexResult("LexSucc", PureScript$dCST$dTypes.$Token("TokInt", "0x" + v3._1, PureScript$dCST$dTypes.$IntValue("BigHex", v3._1)), v3._2);
          }
        }
      }
      $runtime.fail();
    };
  })())(str => {
    const v1 = intPartRegex(str);
    if (v1.tag === "LexFail") { return $LexResult("LexFail", v1._1, v1._2); }
    if (v1.tag === "LexSucc") {
      const v1$1 = optional((() => {
        const $0 = $$char(PureScript$dCST$dErrors.LexExpected("dot"))(".");
        const $1 = regex(PureScript$dCST$dErrors.LexExpected("fraction part"))("[0-9_]+");
        return str$1 => {
          const v1$1 = $0(str$1);
          if (v1$1.tag === "LexFail") { return $LexResult("LexFail", v1$1._1, str$1); }
          if (v1$1.tag === "LexSucc") {
            const v3 = $1(v1$1._2);
            if (v3.tag === "LexFail") { return $LexResult("LexFail", v3._1, str$1); }
            if (v3.tag === "LexSucc") { return $LexResult("LexSucc", v3._1, v3._2); }
          }
          $runtime.fail();
        };
      })())(v1._2);
      if (v1$1.tag === "LexFail") { return $LexResult("LexFail", v1$1._1, v1$1._2); }
      if (v1$1.tag === "LexSucc") {
        const v1$2 = optional((() => {
          const $0 = $$char(PureScript$dCST$dErrors.LexExpected("exponent"))("e");
          const $1 = optional(altLex.alt(string(PureScript$dCST$dErrors.LexExpected("negative"))("-"))(string(PureScript$dCST$dErrors.LexExpected("positive"))("+")));
          return str$1 => {
            const v1$2 = $0(str$1);
            if (v1$2.tag === "LexFail") { return $LexResult("LexFail", v1$2._1, v1$2._2); }
            if (v1$2.tag === "LexSucc") {
              const v1$3 = $1(v1$2._2);
              if (v1$3.tag === "LexFail") { return $LexResult("LexFail", v1$3._1, v1$3._2); }
              if (v1$3.tag === "LexSucc") {
                const v3 = intPartRegex(v1$3._2);
                if (v3.tag === "LexFail") { return $LexResult("LexFail", v3._1, v3._2); }
                if (v3.tag === "LexSucc") { return $LexResult("LexSucc", {sign: v1$3._1, exponent: v3._1}, v3._2); }
              }
            }
            $runtime.fail();
          };
        })())(v1$1._2);
        if (v1$2.tag === "LexFail") { return $LexResult("LexFail", v1$2._1, v1$2._2); }
        if (v1$2.tag === "LexSucc") {
          if (
            (() => {
              if (v1$1._1.tag === "Nothing") { return true; }
              if (v1$1._1.tag === "Just") { return false; }
              $runtime.fail();
            })() && (() => {
              if (v1$2._1.tag === "Nothing") { return true; }
              if (v1$2._1.tag === "Just") { return false; }
              $runtime.fail();
            })()
          ) {
            const intVal = stripUnderscores(v1._1);
            const v = Data$dInt.fromString(intVal);
            if (v.tag === "Just") { return $LexResult("LexSucc", PureScript$dCST$dTypes.$Token("TokInt", v1._1, PureScript$dCST$dTypes.$IntValue("SmallInt", v._1)), v1$2._2); }
            if (v.tag === "Nothing") { return $LexResult("LexSucc", PureScript$dCST$dTypes.$Token("TokInt", v1._1, PureScript$dCST$dTypes.$IntValue("BigInt", intVal)), v1$2._2); }
            $runtime.fail();
          }
          const raw = (() => {
            if (v1$1._1.tag === "Nothing") { return v1._1 + ""; }
            if (v1$1._1.tag === "Just") { return v1._1 + "." + v1$1._1._1; }
            $runtime.fail();
          })() + (() => {
            if (v1$2._1.tag === "Nothing") { return ""; }
            if (v1$2._1.tag === "Just") {
              if (v1$2._1._1.sign.tag === "Nothing") { return "e" + v1$2._1._1.exponent; }
              if (v1$2._1._1.sign.tag === "Just") { return "e" + v1$2._1._1.sign._1 + v1$2._1._1.exponent; }
            }
            $runtime.fail();
          })();
          const v = Data$dNumber.fromStringImpl(stripUnderscores(raw), Data$dNumber.isFinite, Data$dMaybe.Just, Data$dMaybe.Nothing);
          if (v.tag === "Just") { return $LexResult("LexSucc", PureScript$dCST$dTypes.$Token("TokNumber", raw, v._1), v1$2._2); }
          if (v.tag === "Nothing") { return $LexResult("LexFail", v$1 => PureScript$dCST$dErrors.$ParseError("LexNumberOutOfRange", raw), v1$2._2); }
        }
      }
    }
    $runtime.fail();
  }))(altLex.alt(tokenLeftParen)(altLex.alt(tokenRightParen)(altLex.alt(char$p(PureScript$dCST$dErrors.LexExpected("left brace"))(PureScript$dCST$dTypes.TokLeftBrace)("{"))(altLex.alt(char$p(PureScript$dCST$dErrors.LexExpected("right brace"))(PureScript$dCST$dTypes.TokRightBrace)("}"))(altLex.alt(char$p(PureScript$dCST$dErrors.LexExpected("left square"))(PureScript$dCST$dTypes.TokLeftSquare)("["))(altLex.alt(char$p(PureScript$dCST$dErrors.LexExpected("right square"))(PureScript$dCST$dTypes.TokRightSquare)("]"))(altLex.alt(char$p(PureScript$dCST$dErrors.LexExpected("backtick"))(PureScript$dCST$dTypes.TokTick)("`"))(char$p(PureScript$dCST$dErrors.LexExpected("comma"))(PureScript$dCST$dTypes.TokComma)(",")))))))))))));
})();
const lexToken = x => {
  const $0 = token(x);
  if ($0.tag === "LexSucc") {
    if ($0._2 === "") { return Data$dEither.$Either("Right", $0._1); }
    const $1 = $0._1;
    return Data$dEither.$Either("Left", v1 => PureScript$dCST$dErrors.$ParseError("ExpectedEof", $1));
  }
  if ($0.tag === "LexFail") { return Data$dEither.$Either("Left", $0._1); }
  $runtime.fail();
};
const trailingComments = /* #__PURE__ */ (() => many(altLex.alt(str => {
  const v1 = comment(str);
  if (v1.tag === "LexFail") { return $LexResult("LexFail", v1._1, v1._2); }
  if (v1.tag === "LexSucc") { return $LexResult("LexSucc", PureScript$dCST$dTypes.$Comment("Comment", v1._1), v1._2); }
  $runtime.fail();
})(str => {
  const v1 = spaceComment(str);
  if (v1.tag === "LexFail") { return $LexResult("LexFail", v1._1, v1._2); }
  if (v1.tag === "LexSucc") { return $LexResult("LexSucc", PureScript$dCST$dTypes.$Comment("Space", v1._1), v1._2); }
  $runtime.fail();
})))();
const lexWithState$p = lexLeadingComments => {
  const go = stack => startPos => leading => str => Data$dLazy.defer(v => {
    if (str === "") {
      return Data$dLazy.force(PureScript$dCST$dTokenStream.unwindLayout(startPos)(Data$dLazy.defer(v1 => PureScript$dCST$dTokenStream.$TokenStep("TokenEOF", startPos, leading)))(stack));
    }
    const v1 = token(str);
    if (v1.tag === "LexFail") {
      return PureScript$dCST$dTokenStream.$TokenStep(
        "TokenError",
        bumpText(startPos)(0)(Data$dString$dCodeUnits.take(Data$dString$dCodeUnits.length(str) - Data$dString$dCodeUnits.length(v1._2) | 0)(str)),
        v1._1(),
        Data$dMaybe.Nothing,
        stack
      );
    }
    if (v1.tag === "LexSucc") {
      const v3 = trailingComments(v1._2);
      if (v3.tag === "LexFail") {
        return PureScript$dCST$dTokenStream.$TokenStep(
          "TokenError",
          bumpText(startPos)(0)(Data$dString$dCodeUnits.take(Data$dString$dCodeUnits.length(str) - Data$dString$dCodeUnits.length(v3._2) | 0)(str)),
          v3._1(),
          Data$dMaybe.Nothing,
          stack
        );
      }
      if (v3.tag === "LexSucc") {
        const v3$1 = leadingComments(v3._2);
        if (v3$1.tag === "LexFail") {
          return PureScript$dCST$dTokenStream.$TokenStep(
            "TokenError",
            bumpText(startPos)(0)(Data$dString$dCodeUnits.take(Data$dString$dCodeUnits.length(str) - Data$dString$dCodeUnits.length(v3$1._2) | 0)(str)),
            v3$1._1(),
            Data$dMaybe.Nothing,
            stack
          );
        }
        if (v3$1.tag === "LexSucc") {
          const endPos = bumpToken(startPos)(v1._1);
          const nextStart = Data$dFoldable.foldlArray(bumpComment)(Data$dFoldable.foldlArray(bumpComment)(endPos)(v3._1))(v3$1._1);
          const v2 = PureScript$dCST$dLayout.insertLayout({range: {start: startPos, end: endPos}, leadingComments: leading, trailingComments: v3._1, value: v1._1})(nextStart)(stack);
          return Data$dLazy.force(consTokens(v2._2)(Data$dTuple.$Tuple(nextStart, go(v2._1)(nextStart)(v3$1._1)(v3$1._2)))._2);
        }
      }
    }
    $runtime.fail();
  });
  return initStack => initPos => str => Data$dLazy.defer(v => {
    const v1 = lexLeadingComments(str);
    if (v1.tag === "LexFail") { return Partial._crashWith("Leading comments can't fail."); }
    if (v1.tag === "LexSucc") { return Data$dLazy.force(go(initStack)(Data$dFoldable.foldlArray(bumpComment)(initPos)(v1._1))(v1._1)(v1._2)); }
    $runtime.fail();
  });
};
const lexModule = /* #__PURE__ */ lexWithState$p(leadingModuleComments)(/* #__PURE__ */ Data$dList$dTypes.$List(
  "Cons",
  /* #__PURE__ */ Data$dTuple.$Tuple({line: 0, column: 0}, PureScript$dCST$dLayout.LytRoot),
  Data$dList$dTypes.Nil
))({line: 0, column: 0});
const lexWithState = /* #__PURE__ */ lexWithState$p(leadingComments);
const lex = /* #__PURE__ */ lexWithState(/* #__PURE__ */ Data$dList$dTypes.$List(
  "Cons",
  /* #__PURE__ */ Data$dTuple.$Tuple({line: 0, column: 0}, PureScript$dCST$dLayout.LytRoot),
  Data$dList$dTypes.Nil
))({line: 0, column: 0});
export {
  $LexResult,
  LexFail,
  LexSucc,
  altLex,
  bumpComment,
  bumpText,
  bumpToken,
  $$char as char,
  char$p,
  comment,
  consTokens,
  fold1,
  foldMap,
  functorLex,
  isCharChar,
  isCharCodePoint,
  leadingComments,
  leadingModuleComments,
  leadingShebangs,
  lex,
  lexModule,
  lexToken,
  lexWithState,
  lexWithState$p,
  lineComment,
  many,
  mkUnexpected,
  oneLineComment,
  optional,
  regex,
  satisfy,
  shebangComment,
  spaceComment,
  string,
  toModuleName,
  token,
  trailingComments
};
