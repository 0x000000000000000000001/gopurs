import * as $runtime from "../runtime.js";
import * as Data$dFoldable from "../Data.Foldable/index.js";
import * as Data$dLazy from "../Data.Lazy/index.js";
import * as Data$dMonoid from "../Data.Monoid/index.js";
import * as PureScript$dCST$dLexer from "../PureScript.CST.Lexer/index.js";
import * as PureScript$dCST$dParser from "../PureScript.CST.Parser/index.js";
import * as PureScript$dCST$dParser$dMonad from "../PureScript.CST.Parser.Monad/index.js";
import * as PureScript$dCST$dPrint from "../PureScript.CST.Print/index.js";
import * as PureScript$dCST$dRange from "../PureScript.CST.Range/index.js";
import * as PureScript$dCST$dRange$dTokenList from "../PureScript.CST.Range.TokenList/index.js";
import * as Unsafe$dCoerce from "../Unsafe.Coerce/index.js";
const $RecoveredParserResult = (tag, _1, _2) => ({tag, _1, _2});
const foldMap = /* #__PURE__ */ (() => Data$dFoldable.foldableArray.foldMap(Data$dMonoid.monoidString))();
const ParseSucceeded = value0 => $RecoveredParserResult("ParseSucceeded", value0);
const ParseSucceededWithErrors = value0 => value1 => $RecoveredParserResult("ParseSucceededWithErrors", value0, value1);
const ParseFailed = value0 => $RecoveredParserResult("ParseFailed", value0);
const PartialModule = x => x;
const toRecovered = Unsafe$dCoerce.unsafeCoerce;
const runRecoveredParser = p => x => {
  const $0 = PureScript$dCST$dParser$dMonad.runParser(x)(p);
  if ($0.tag === "Right" && $0._1._2.length > 0) { return $RecoveredParserResult("ParseSucceededWithErrors", $0._1._1, $0._1._2); }
  if ($0.tag === "Right") { return $RecoveredParserResult("ParseSucceeded", $0._1._1); }
  if ($0.tag === "Left") { return $RecoveredParserResult("ParseFailed", $0._1); }
  $runtime.fail();
};
const printModule = dictTokensOf => mod => foldMap(PureScript$dCST$dPrint.printSourceTokenWithOption(PureScript$dCST$dPrint.HideLayout))(PureScript$dCST$dRange$dTokenList.toArray(PureScript$dCST$dRange.tokensOfModule(dictTokensOf).tokensOf(mod))) + foldMap(PureScript$dCST$dPrint.printComment(PureScript$dCST$dPrint.printLineFeed))(mod.body.trailingComments);
const parseType = x => runRecoveredParser(PureScript$dCST$dParser.parseType)(PureScript$dCST$dLexer.lex(x));
const parsePartialModule = src => {
  const v = PureScript$dCST$dParser$dMonad.runParser$p({consumed: false, errors: [], stream: PureScript$dCST$dLexer.lexModule(src)})(PureScript$dCST$dParser.parseModuleHeader);
  if (v.tag === "ParseSucc") {
    if (v._2.errors.length > 0) {
      return $RecoveredParserResult(
        "ParseSucceededWithErrors",
        {
          header: v._1,
          full: Data$dLazy.defer(v1 => {
            const $0 = PureScript$dCST$dParser$dMonad.runParser$p(v._2)((state1, more, resume, done) => PureScript$dCST$dParser.parseModuleBody(
              state1,
              more,
              resume,
              (state2, a) => more(v1$1 => done(state2, {header: v._1, body: a}))
            ));
            if ($0.tag === "ParseFail") { return $RecoveredParserResult("ParseFailed", $0._1); }
            if ($0.tag === "ParseSucc") {
              if ($0._2.errors.length > 0) { return $RecoveredParserResult("ParseSucceededWithErrors", $0._1, $0._2.errors); }
              return $RecoveredParserResult("ParseSucceeded", $0._1);
            }
            $runtime.fail();
          })
        },
        v._2.errors
      );
    }
    return $RecoveredParserResult(
      "ParseSucceeded",
      {
        header: v._1,
        full: Data$dLazy.defer(v1 => {
          const $0 = PureScript$dCST$dParser$dMonad.runParser$p(v._2)((state1, more, resume, done) => PureScript$dCST$dParser.parseModuleBody(
            state1,
            more,
            resume,
            (state2, a) => more(v1$1 => done(state2, {header: v._1, body: a}))
          ));
          if ($0.tag === "ParseFail") { return $RecoveredParserResult("ParseFailed", $0._1); }
          if ($0.tag === "ParseSucc") {
            if ($0._2.errors.length > 0) { return $RecoveredParserResult("ParseSucceededWithErrors", $0._1, $0._2.errors); }
            return $RecoveredParserResult("ParseSucceeded", $0._1);
          }
          $runtime.fail();
        })
      }
    );
  }
  if (v.tag === "ParseFail") { return $RecoveredParserResult("ParseFailed", v._1); }
  $runtime.fail();
};
const parseModule = x => runRecoveredParser(PureScript$dCST$dParser.parseModule)(PureScript$dCST$dLexer.lexModule(x));
const parseImportDecl = x => runRecoveredParser(PureScript$dCST$dParser.parseImportDecl)(PureScript$dCST$dLexer.lex(x));
const parseExpr = x => runRecoveredParser(PureScript$dCST$dParser.parseExpr)(PureScript$dCST$dLexer.lex(x));
const parseDecl = x => runRecoveredParser(PureScript$dCST$dParser.parseDecl)(PureScript$dCST$dLexer.lex(x));
const parseBinder = x => runRecoveredParser(PureScript$dCST$dParser.parseBinder)(PureScript$dCST$dLexer.lex(x));
export {
  $RecoveredParserResult,
  ParseFailed,
  ParseSucceeded,
  ParseSucceededWithErrors,
  PartialModule,
  foldMap,
  parseBinder,
  parseDecl,
  parseExpr,
  parseImportDecl,
  parseModule,
  parsePartialModule,
  parseType,
  printModule,
  runRecoveredParser,
  toRecovered
};
