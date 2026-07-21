import * as $runtime from "../runtime.js";
import * as Data$dFoldable from "../Data.Foldable/index.js";
import * as Data$dMonoid from "../Data.Monoid/index.js";
import * as Data$dVoid from "../Data.Void/index.js";
const $TokenOption = tag => tag;
const power = /* #__PURE__ */ Data$dMonoid.power(Data$dMonoid.monoidString);
const foldMap = /* #__PURE__ */ (() => Data$dFoldable.foldableArray.foldMap(Data$dMonoid.monoidString))();
const ShowLayout = /* #__PURE__ */ $TokenOption("ShowLayout");
const HideLayout = /* #__PURE__ */ $TokenOption("HideLayout");
const printQualified = moduleName => name => {
  if (moduleName.tag === "Nothing") { return name; }
  if (moduleName.tag === "Just") { return moduleName._1 + "." + name; }
  $runtime.fail();
};
const printTokenWithOption = option => v => {
  if (v.tag === "TokLeftParen") { return "("; }
  if (v.tag === "TokRightParen") { return ")"; }
  if (v.tag === "TokLeftBrace") { return "{"; }
  if (v.tag === "TokRightBrace") { return "}"; }
  if (v.tag === "TokLeftSquare") { return "["; }
  if (v.tag === "TokRightSquare") { return "]"; }
  if (v.tag === "TokLeftArrow") {
    if (v._1 === "ASCII") { return "<-"; }
    if (v._1 === "Unicode") { return "←"; }
    $runtime.fail();
  }
  if (v.tag === "TokRightArrow") {
    if (v._1 === "ASCII") { return "->"; }
    if (v._1 === "Unicode") { return "→"; }
    $runtime.fail();
  }
  if (v.tag === "TokRightFatArrow") {
    if (v._1 === "ASCII") { return "=>"; }
    if (v._1 === "Unicode") { return "⇒"; }
    $runtime.fail();
  }
  if (v.tag === "TokDoubleColon") {
    if (v._1 === "ASCII") { return "::"; }
    if (v._1 === "Unicode") { return "∷"; }
    $runtime.fail();
  }
  if (v.tag === "TokForall") {
    if (v._1 === "ASCII") { return "forall"; }
    if (v._1 === "Unicode") { return "∀"; }
    $runtime.fail();
  }
  if (v.tag === "TokEquals") { return "="; }
  if (v.tag === "TokPipe") { return "|"; }
  if (v.tag === "TokTick") { return "`"; }
  if (v.tag === "TokDot") { return "."; }
  if (v.tag === "TokComma") { return ","; }
  if (v.tag === "TokUnderscore") { return "_"; }
  if (v.tag === "TokBackslash") { return "\\"; }
  if (v.tag === "TokAt") { return "@"; }
  if (v.tag === "TokLowerName") {
    if (v._1.tag === "Nothing") { return v._2; }
    if (v._1.tag === "Just") { return v._1._1 + "." + v._2; }
    $runtime.fail();
  }
  if (v.tag === "TokUpperName") {
    if (v._1.tag === "Nothing") { return v._2; }
    if (v._1.tag === "Just") { return v._1._1 + "." + v._2; }
    $runtime.fail();
  }
  if (v.tag === "TokOperator") {
    if (v._1.tag === "Nothing") { return v._2; }
    if (v._1.tag === "Just") { return v._1._1 + "." + v._2; }
    $runtime.fail();
  }
  if (v.tag === "TokSymbolName") {
    const $0 = "(" + v._2 + ")";
    if (v._1.tag === "Nothing") { return $0; }
    if (v._1.tag === "Just") { return v._1._1 + "." + $0; }
    $runtime.fail();
  }
  if (v.tag === "TokSymbolArrow") {
    if (v._1 === "ASCII") { return "(->)"; }
    if (v._1 === "Unicode") { return "(→)"; }
    $runtime.fail();
  }
  if (v.tag === "TokHole") { return "?" + v._1; }
  if (v.tag === "TokChar") { return "'" + v._1 + "'"; }
  if (v.tag === "TokString") { return "\"" + v._1 + "\""; }
  if (v.tag === "TokRawString") { return "\"\"\"" + v._1 + "\"\"\""; }
  if (v.tag === "TokInt") { return v._1; }
  if (v.tag === "TokNumber") { return v._1; }
  if (v.tag === "TokLayoutStart") {
    if (option === "ShowLayout") { return "{"; }
    if (option === "HideLayout") { return ""; }
    $runtime.fail();
  }
  if (v.tag === "TokLayoutSep") {
    if (option === "ShowLayout") { return ";"; }
    if (option === "HideLayout") { return ""; }
    $runtime.fail();
  }
  if (v.tag === "TokLayoutEnd") {
    if (option === "ShowLayout") { return "}"; }
    if (option === "HideLayout") { return ""; }
  }
  $runtime.fail();
};
const printToken = /* #__PURE__ */ printTokenWithOption(HideLayout);
const printLineFeed = v => {
  if (v === "LF") { return "\n"; }
  if (v === "CRLF") { return "\r\n"; }
  $runtime.fail();
};
const printComment = k => v => {
  if (v.tag === "Comment") { return v._1; }
  if (v.tag === "Space") { return power(" ")(v._1); }
  if (v.tag === "Line") { return power(k(v._1))(v._2); }
  $runtime.fail();
};
const printSourceTokenWithOption = option => tok => foldMap(printComment(printLineFeed))(tok.leadingComments) + printTokenWithOption(option)(tok.value) + foldMap(printComment(Data$dVoid.absurd))(tok.trailingComments);
const printSourceToken = /* #__PURE__ */ printSourceTokenWithOption(HideLayout);
export {
  $TokenOption,
  HideLayout,
  ShowLayout,
  foldMap,
  power,
  printComment,
  printLineFeed,
  printQualified,
  printSourceToken,
  printSourceTokenWithOption,
  printToken,
  printTokenWithOption
};
