import * as $runtime from "../runtime.js";
const $ParseError = (tag, _1, _2) => ({tag, _1, _2});
const UnexpectedEof = /* #__PURE__ */ $ParseError("UnexpectedEof");
const ExpectedEof = value0 => $ParseError("ExpectedEof", value0);
const UnexpectedToken = value0 => $ParseError("UnexpectedToken", value0);
const ExpectedToken = value0 => value1 => $ParseError("ExpectedToken", value0, value1);
const ExpectedClass = value0 => value1 => $ParseError("ExpectedClass", value0, value1);
const LexExpected = value0 => value1 => $ParseError("LexExpected", value0, value1);
const LexInvalidCharEscape = value0 => $ParseError("LexInvalidCharEscape", value0);
const LexCharEscapeOutOfRange = value0 => $ParseError("LexCharEscapeOutOfRange", value0);
const LexHexOutOfRange = value0 => $ParseError("LexHexOutOfRange", value0);
const LexIntOutOfRange = value0 => $ParseError("LexIntOutOfRange", value0);
const LexNumberOutOfRange = value0 => $ParseError("LexNumberOutOfRange", value0);
const RecoveredError = x => x;
const printTokenError = v => {
  if (v.tag === "TokLeftParen") { return "'('"; }
  if (v.tag === "TokRightParen") { return "')'"; }
  if (v.tag === "TokLeftBrace") { return "'{'"; }
  if (v.tag === "TokRightBrace") { return "'}'"; }
  if (v.tag === "TokLeftSquare") { return "'['"; }
  if (v.tag === "TokRightSquare") { return "']'"; }
  if (v.tag === "TokLeftArrow") {
    if (v._1 === "ASCII") { return "'<-'"; }
    if (v._1 === "Unicode") { return "'←'"; }
    $runtime.fail();
  }
  if (v.tag === "TokRightArrow") {
    if (v._1 === "ASCII") { return "'->'"; }
    if (v._1 === "Unicode") { return "'→'"; }
    $runtime.fail();
  }
  if (v.tag === "TokRightFatArrow") {
    if (v._1 === "ASCII") { return "'=>'"; }
    if (v._1 === "Unicode") { return "'⇒'"; }
    $runtime.fail();
  }
  if (v.tag === "TokDoubleColon") {
    if (v._1 === "ASCII") { return "'::'"; }
    if (v._1 === "Unicode") { return "'∷'"; }
    $runtime.fail();
  }
  if (v.tag === "TokForall") {
    if (v._1 === "ASCII") { return "forall"; }
    if (v._1 === "Unicode") { return "'∀'"; }
    $runtime.fail();
  }
  if (v.tag === "TokEquals") { return "'='"; }
  if (v.tag === "TokPipe") { return "'|'"; }
  if (v.tag === "TokTick") { return "`"; }
  if (v.tag === "TokDot") { return "."; }
  if (v.tag === "TokComma") { return "','"; }
  if (v.tag === "TokUnderscore") { return "'_'"; }
  if (v.tag === "TokBackslash") { return "'\\'"; }
  if (v.tag === "TokAt") { return "'@'"; }
  if (v.tag === "TokLowerName") {
    if (v._1.tag === "Nothing") { return "identifier " + v._2; }
    if (v._1.tag === "Just") { return "identifier " + v._1._1 + "." + v._2; }
    $runtime.fail();
  }
  if (v.tag === "TokUpperName") {
    if (v._1.tag === "Nothing") { return "proper identifier " + v._2; }
    if (v._1.tag === "Just") { return "proper identifier " + v._1._1 + "." + v._2; }
    $runtime.fail();
  }
  if (v.tag === "TokOperator") {
    if (v._1.tag === "Nothing") { return "operator " + v._2; }
    if (v._1.tag === "Just") { return "operator " + v._1._1 + "." + v._2; }
    $runtime.fail();
  }
  if (v.tag === "TokSymbolName") {
    if (v._1.tag === "Nothing") { return "symbol " + v._2; }
    if (v._1.tag === "Just") { return "symbol " + v._1._1 + "." + v._2; }
    $runtime.fail();
  }
  if (v.tag === "TokSymbolArrow") {
    if (v._1 === "ASCII") { return "(->)"; }
    if (v._1 === "Unicode") { return "(→)"; }
    $runtime.fail();
  }
  if (v.tag === "TokHole") { return "hole ?" + v._1; }
  if (v.tag === "TokChar") { return "char literal '" + v._1 + "'"; }
  if (v.tag === "TokString") { return "string literal"; }
  if (v.tag === "TokRawString") { return "raw string literal"; }
  if (v.tag === "TokInt") { return "int literal " + v._1; }
  if (v.tag === "TokNumber") { return "number literal " + v._1; }
  if (v.tag === "TokLayoutStart") { return "start of indented block"; }
  if (v.tag === "TokLayoutSep") { return "new indented block item"; }
  if (v.tag === "TokLayoutEnd") { return "end of indented block"; }
  $runtime.fail();
};
const printParseError = v => {
  if (v.tag === "UnexpectedEof") { return "Unexpected end of file"; }
  if (v.tag === "ExpectedEof") { return "Expected end of file, saw " + printTokenError(v._1); }
  if (v.tag === "UnexpectedToken") { return "Unexpected " + printTokenError(v._1); }
  if (v.tag === "ExpectedToken") { return "Expected " + printTokenError(v._1) + ", saw " + printTokenError(v._2); }
  if (v.tag === "ExpectedClass") { return "Expected " + v._1 + ", saw " + printTokenError(v._2); }
  if (v.tag === "LexExpected") { return "Expected " + v._1 + ", saw " + v._2; }
  if (v.tag === "LexInvalidCharEscape") { return "Invalid character escape \\" + v._1; }
  if (v.tag === "LexCharEscapeOutOfRange") { return "Character escape out of range \\" + v._1; }
  if (v.tag === "LexHexOutOfRange") { return "Hex integer out of range 0x" + v._1; }
  if (v.tag === "LexIntOutOfRange") { return "Int out of range " + v._1; }
  if (v.tag === "LexNumberOutOfRange") { return "Number out of range " + v._1; }
  $runtime.fail();
};
export {
  $ParseError,
  ExpectedClass,
  ExpectedEof,
  ExpectedToken,
  LexCharEscapeOutOfRange,
  LexExpected,
  LexHexOutOfRange,
  LexIntOutOfRange,
  LexInvalidCharEscape,
  LexNumberOutOfRange,
  RecoveredError,
  UnexpectedEof,
  UnexpectedToken,
  printParseError,
  printTokenError
};
