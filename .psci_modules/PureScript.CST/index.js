import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Data_Array_NonEmpty from "../Data.Array.NonEmpty/index.js";
import * as Data_Boolean from "../Data.Boolean/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Lazy from "../Data.Lazy/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Newtype from "../Data.Newtype/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as PureScript_CST_Lexer from "../PureScript.CST.Lexer/index.js";
import * as PureScript_CST_Parser from "../PureScript.CST.Parser/index.js";
import * as PureScript_CST_Parser_Monad from "../PureScript.CST.Parser.Monad/index.js";
import * as PureScript_CST_Print from "../PureScript.CST.Print/index.js";
import * as PureScript_CST_Range from "../PureScript.CST.Range/index.js";
import * as PureScript_CST_Range_TokenList from "../PureScript.CST.Range.TokenList/index.js";
import * as Unsafe_Coerce from "../Unsafe.Coerce/index.js";
var pure = /* #__PURE__ */ Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser);
var ParseSucceeded = /* #__PURE__ */ (function () {
    function ParseSucceeded(value0) {
        this.value0 = value0;
    };
    ParseSucceeded.create = function (value0) {
        return new ParseSucceeded(value0);
    };
    return ParseSucceeded;
})();
var ParseSucceededWithErrors = /* #__PURE__ */ (function () {
    function ParseSucceededWithErrors(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    ParseSucceededWithErrors.create = function (value0) {
        return function (value1) {
            return new ParseSucceededWithErrors(value0, value1);
        };
    };
    return ParseSucceededWithErrors;
})();
var ParseFailed = /* #__PURE__ */ (function () {
    function ParseFailed(value0) {
        this.value0 = value0;
    };
    ParseFailed.create = function (value0) {
        return new ParseFailed(value0);
    };
    return ParseFailed;
})();
var PartialModule = function (x) {
    return x;
};
var toRecoveredParserResult = function (v) {
    var v1 = function (v2) {
        if (v instanceof Data_Either.Right && Data_Boolean.otherwise) {
            return new ParseSucceeded(v.value0.value0);
        };
        if (v instanceof Data_Either.Left) {
            return new ParseFailed(v.value0);
        };
        throw new Error("Failed pattern match at PureScript.CST (line 42, column 1 - line 45, column 29): " + [ v.constructor.name ]);
    };
    if (v instanceof Data_Either.Right) {
        var $14 = Data_Array_NonEmpty.fromArray(v.value0.value1);
        if ($14 instanceof Data_Maybe.Just) {
            return new ParseSucceededWithErrors(v.value0.value0, $14.value0);
        };
        return v1(true);
    };
    return v1(true);
};
var toRecovered = Unsafe_Coerce.unsafeCoerce;
var runRecoveredParser = function (p) {
    var $24 = Data_Function.flip(PureScript_CST_Parser_Monad.runParser)(p);
    return function ($25) {
        return toRecoveredParserResult($24($25));
    };
};
var printModule = function (dictTokensOf) {
    var tokensOfModule = PureScript_CST_Range.tokensOfModule(dictTokensOf);
    return function (mod) {
        return Data_Foldable.foldMap(Data_Foldable.foldableArray)(Data_Monoid.monoidString)(PureScript_CST_Print.printSourceToken)(PureScript_CST_Range_TokenList.toArray(PureScript_CST_Range.tokensOf(tokensOfModule)(mod))) + Data_Foldable.foldMap(Data_Foldable.foldableArray)(Data_Monoid.monoidString)(PureScript_CST_Print.printComment(PureScript_CST_Print.printLineFeed))((Data_Newtype.unwrap()((Data_Newtype.unwrap()(mod)).body)).trailingComments);
    };
};
var parseType = /* #__PURE__ */ (function () {
    var $26 = runRecoveredParser(PureScript_CST_Parser.parseType);
    return function ($27) {
        return $26(PureScript_CST_Lexer.lex($27));
    };
})();
var parsePartialModule = function (src) {
    return toRecoveredParserResult((function () {
        var v = PureScript_CST_Parser_Monad["runParser$prime"](PureScript_CST_Parser_Monad.initialParserState(PureScript_CST_Lexer.lexModule(src)))(PureScript_CST_Parser.parseModuleHeader);
        if (v instanceof PureScript_CST_Parser_Monad.ParseSucc) {
            var res = {
                header: v.value0,
                full: Data_Lazy.defer(function (v1) {
                    return toRecoveredParserResult(PureScript_CST_Parser_Monad.fromParserResult(PureScript_CST_Parser_Monad["runParser$prime"](v.value1)(Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser.parseModuleBody)(function (body) {
                        return pure({
                            header: v.value0,
                            body: body
                        });
                    }))));
                })
            };
            return new Data_Either.Right(new Data_Tuple.Tuple(res, v.value1.errors));
        };
        if (v instanceof PureScript_CST_Parser_Monad.ParseFail) {
            return new Data_Either.Left(v.value0);
        };
        throw new Error("Failed pattern match at PureScript.CST (line 86, column 29 - line 98, column 17): " + [ v.constructor.name ]);
    })());
};
var parseModule = /* #__PURE__ */ (function () {
    var $28 = runRecoveredParser(PureScript_CST_Parser.parseModule);
    return function ($29) {
        return $28(PureScript_CST_Lexer.lexModule($29));
    };
})();
var parseImportDecl = /* #__PURE__ */ (function () {
    var $30 = runRecoveredParser(PureScript_CST_Parser.parseImportDecl);
    return function ($31) {
        return $30(PureScript_CST_Lexer.lex($31));
    };
})();
var parseExpr = /* #__PURE__ */ (function () {
    var $32 = runRecoveredParser(PureScript_CST_Parser.parseExpr);
    return function ($33) {
        return $32(PureScript_CST_Lexer.lex($33));
    };
})();
var parseDecl = /* #__PURE__ */ (function () {
    var $34 = runRecoveredParser(PureScript_CST_Parser.parseDecl);
    return function ($35) {
        return $34(PureScript_CST_Lexer.lex($35));
    };
})();
var parseBinder = /* #__PURE__ */ (function () {
    var $36 = runRecoveredParser(PureScript_CST_Parser.parseBinder);
    return function ($37) {
        return $36(PureScript_CST_Lexer.lex($37));
    };
})();
export {
    ParseSucceeded,
    ParseSucceededWithErrors,
    ParseFailed,
    PartialModule,
    parseModule,
    parsePartialModule,
    parseImportDecl,
    parseDecl,
    parseExpr,
    parseType,
    parseBinder,
    printModule,
    toRecovered
};
