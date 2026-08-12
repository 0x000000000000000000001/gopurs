import * as Control_Alt from "../Control.Alt/index.js";
import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Category from "../Control.Category/index.js";
import * as Data_Array from "../Data.Array/index.js";
import * as Data_Array_NonEmpty from "../Data.Array.NonEmpty/index.js";
import * as Data_Array_ST from "../Data.Array.ST/index.js";
import * as Data_Boolean from "../Data.Boolean/index.js";
import * as Data_Char from "../Data.Char/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Enum from "../Data.Enum/index.js";
import * as Data_EuclideanRing from "../Data.EuclideanRing/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Int from "../Data.Int/index.js";
import * as Data_Lazy from "../Data.Lazy/index.js";
import * as Data_List_Types from "../Data.List.Types/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Newtype from "../Data.Newtype/index.js";
import * as Data_Number from "../Data.Number/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Semiring from "../Data.Semiring/index.js";
import * as Data_String_CodePoints from "../Data.String.CodePoints/index.js";
import * as Data_String_CodeUnits from "../Data.String.CodeUnits/index.js";
import * as Data_String_Common from "../Data.String.Common/index.js";
import * as Data_String_Regex from "../Data.String.Regex/index.js";
import * as Data_String_Regex_Flags from "../Data.String.Regex.Flags/index.js";
import * as Data_String_Regex_Unsafe from "../Data.String.Regex.Unsafe/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Data_Unit from "../Data.Unit/index.js";
import * as Partial_Unsafe from "../Partial.Unsafe/index.js";
import * as PureScript_CST_Errors from "../PureScript.CST.Errors/index.js";
import * as PureScript_CST_Layout from "../PureScript.CST.Layout/index.js";
import * as PureScript_CST_TokenStream from "../PureScript.CST.TokenStream/index.js";
import * as PureScript_CST_Types from "../PureScript.CST.Types/index.js";
var monoidRecord = /* #__PURE__ */ Data_Monoid.monoidRecord()(/* #__PURE__ */ Data_Monoid.monoidRecordCons({
    reflectSymbol: function () {
        return "raw";
    }
})(Data_Monoid.monoidString)()(/* #__PURE__ */ Data_Monoid.monoidRecordCons({
    reflectSymbol: function () {
        return "string";
    }
})(Data_Monoid.monoidString)()(Data_Monoid.monoidRecordNil)));
var LexFail = /* #__PURE__ */ (function () {
    function LexFail(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    LexFail.create = function (value0) {
        return function (value1) {
            return new LexFail(value0, value1);
        };
    };
    return LexFail;
})();
var LexSucc = /* #__PURE__ */ (function () {
    function LexSucc(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    LexSucc.create = function (value0) {
        return function (value1) {
            return new LexSucc(value0, value1);
        };
    };
    return LexSucc;
})();
var Lex = function (x) {
    return x;
};
var isCharCodePoint = {
    fromChar: Data_String_CodePoints.codePointFromChar,
    fromCharCode: /* #__PURE__ */ Data_Enum.toEnum(Data_String_CodePoints.boundedEnumCodePoint)
};
var isCharChar = {
    fromChar: /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn),
    fromCharCode: Data_Char.fromCharCode
};
var $$try = function (v) {
    return function (str) {
        var v1 = v(str);
        if (v1 instanceof LexFail) {
            return new LexFail(v1.value0, str);
        };
        if (v1 instanceof LexSucc) {
            return new LexSucc(v1.value0, v1.value1);
        };
        throw new Error("Failed pattern match at PureScript.CST.Lexer (line 104, column 3 - line 106, column 31): " + [ v1.constructor.name ]);
    };
};
var toModuleName = function (v) {
    if (v === "") {
        return Data_Maybe.Nothing.value;
    };
    return new Data_Maybe.Just(Data_String_CodeUnits.dropRight(1)(v));
};
var qualLength = /* #__PURE__ */ Data_Maybe.maybe(0)(/* #__PURE__ */ (function () {
    var $286 = Data_Semiring.add(Data_Semiring.semiringInt)(1);
    var $287 = Data_Newtype.unwrap();
    return function ($288) {
        return $286(Data_String_CodePoints.length($287($288)));
    };
})());
var optional = function (v) {
    return function (str) {
        var v1 = v(str);
        if (v1 instanceof LexFail) {
            if (Data_String_CodeUnits.length(str) === Data_String_CodeUnits.length(v1.value1)) {
                return new LexSucc(Data_Maybe.Nothing.value, str);
            };
            if (Data_Boolean.otherwise) {
                return new LexFail(v1.value0, v1.value1);
            };
        };
        if (v1 instanceof LexSucc) {
            return new LexSucc(new Data_Maybe.Just(v1.value0), v1.value1);
        };
        throw new Error("Failed pattern match at PureScript.CST.Lexer (line 153, column 3 - line 160, column 25): " + [ v1.constructor.name ]);
    };
};
var mkUnexpected = function (str) {
    var start = Data_String_CodePoints.take(6)(str);
    var len = Data_String_CodePoints.length(start);
    var $115 = len === 0;
    if ($115) {
        return "end of file";
    };
    var $116 = len < 6;
    if ($116) {
        return start;
    };
    return start + "...";
};
var regex = function (mkErr) {
    return function (regexStr) {
        var matchRegex = Data_String_Regex_Unsafe.unsafeRegex("^(?:" + (regexStr + ")"))(Data_String_Regex_Flags.unicode);
        return function (str) {
            var v = Data_String_Regex.match(matchRegex)(str);
            var v1 = function (v2) {
                return new LexFail(function (v3) {
                    return mkErr(mkUnexpected(str));
                }, str);
            };
            if (v instanceof Data_Maybe.Just) {
                var $118 = Data_Array_NonEmpty.head(v.value0);
                if ($118 instanceof Data_Maybe.Just) {
                    return new LexSucc($118.value0, Data_String_CodeUnits.drop(Data_String_CodeUnits.length($118.value0))(str));
                };
                return v1(true);
            };
            return v1(true);
        };
    };
};
var shebangComment = /* #__PURE__ */ (function () {
    return regex(PureScript_CST_Errors.LexExpected.create("shebang"))("#![^\\r\\n]*");
})();
var satisfy = function (mkErr) {
    return function (p) {
        return function (str) {
            var v = Data_String_CodeUnits.charAt(0)(str);
            if (v instanceof Data_Maybe.Just && p(v.value0)) {
                return new LexSucc(v.value0, Data_String_CodeUnits.drop(1)(str));
            };
            return new LexFail(function (v1) {
                return mkErr(mkUnexpected(str));
            }, str);
        };
    };
};
var string = function (mkErr) {
    return function (match) {
        return function (str) {
            var $123 = Data_String_CodeUnits.take(Data_String_CodeUnits.length(match))(str) === match;
            if ($123) {
                return new LexSucc(match, Data_String_CodeUnits.drop(Data_String_CodeUnits.length(match))(str));
            };
            return new LexFail(function (v) {
                return mkErr(mkUnexpected(str));
            }, str);
        };
    };
};
var many = function (v) {
    return function (str) {
        return (function __do() {
            var valuesRef = Data_Array_ST["new"]();
            var strRef = str;
            var contRef = true;
            var resRef = new LexSucc([  ], str);
            (function () {
                while (contRef) {
                    (function __do() {
                        var str$prime = strRef;
                        var v1 = v(str$prime);
                        if (v1 instanceof LexFail) {
                            if (Data_String_CodeUnits.length(str$prime) === Data_String_CodeUnits.length(v1.value1)) {
                                var values = Data_Array_ST.unsafeFreeze(valuesRef)();
                                resRef = new LexSucc(values, v1.value1);
                                contRef = false;
                                return Data_Unit.unit;
                            };
                            if (Data_Boolean.otherwise) {
                                resRef = new LexFail(v1.value0, v1.value1);
                                contRef = false;
                                return Data_Unit.unit;
                            };
                        };
                        if (v1 instanceof LexSucc) {
                            Data_Array_ST.push(v1.value0)(valuesRef)();
                            strRef = v1.value1;
                            return Data_Unit.unit;
                        };
                        throw new Error("Failed pattern match at PureScript.CST.Lexer (line 178, column 5 - line 192, column 18): " + [ v1.constructor.name ]);
                    })();
                };
                return {};
            })();
            return resRef;
        })();
    };
};
var functorLex = {
    map: function (f) {
        return function (v) {
            return function (str) {
                var v1 = v(str);
                if (v1 instanceof LexFail) {
                    return new LexFail(v1.value0, v1.value1);
                };
                if (v1 instanceof LexSucc) {
                    return new LexSucc(f(v1.value0), v1.value1);
                };
                throw new Error("Failed pattern match at PureScript.CST.Lexer (line 66, column 5 - line 68, column 37): " + [ v1.constructor.name ]);
            };
        };
    }
};
var spaceComment = /* #__PURE__ */ (function () {
    return Data_Functor.map(functorLex)(Data_String_CodeUnits.length)(regex(PureScript_CST_Errors.LexExpected.create("spaces"))(" +"));
})();
var fromCharCode = function (dict) {
    return dict.fromCharCode;
};
var fromChar = function (dict) {
    return dict.fromChar;
};
var fail = function ($289) {
    return Lex(LexFail.create(Data_Function["const"]($289)));
};
var char$prime = function (mkErr) {
    return function (res) {
        return function (match) {
            return function (str) {
                var $139 = Data_String_CodeUnits.singleton(match) === Data_String_CodeUnits.take(1)(str);
                if ($139) {
                    return new LexSucc(res, Data_String_CodeUnits.drop(1)(str));
                };
                return new LexFail(function (v) {
                    return mkErr(mkUnexpected(str));
                }, str);
            };
        };
    };
};
var $$char = function (mkErr) {
    return function (match) {
        return function (str) {
            var $140 = Data_String_CodeUnits.singleton(match) === Data_String_CodeUnits.take(1)(str);
            if ($140) {
                return new LexSucc(match, Data_String_CodeUnits.drop(1)(str));
            };
            return new LexFail(function (v) {
                return mkErr(mkUnexpected(str));
            }, str);
        };
    };
};
var bumpText = function (v) {
    return function (colOffset) {
        return function (str) {
            var go = function ($copy_n) {
                return function ($copy_ix) {
                    var $tco_var_n = $copy_n;
                    var $tco_done = false;
                    var $tco_result;
                    function $tco_loop(n, ix) {
                        var v1 = Data_String_CodeUnits["indexOf$prime"]("\x0a")(ix)(str);
                        if (v1 instanceof Data_Maybe.Just) {
                            $tco_var_n = n + 1 | 0;
                            $copy_ix = v1.value0 + 1 | 0;
                            return;
                        };
                        if (v1 instanceof Data_Maybe.Nothing) {
                            if (n === 0) {
                                $tco_done = true;
                                return {
                                    line: v.line,
                                    column: (v.column + Data_String_CodePoints.length(str) | 0) + (colOffset * 2 | 0) | 0
                                };
                            };
                            if (Data_Boolean.otherwise) {
                                $tco_done = true;
                                return {
                                    line: v.line + n | 0,
                                    column: Data_String_CodePoints.length(Data_String_CodeUnits.drop(ix)(str)) + colOffset | 0
                                };
                            };
                        };
                        throw new Error("Failed pattern match at PureScript.CST.Lexer (line 348, column 13 - line 357, column 12): " + [ v1.constructor.name ]);
                    };
                    while (!$tco_done) {
                        $tco_result = $tco_loop($tco_var_n, $copy_ix);
                    };
                    return $tco_result;
                };
            };
            return go(0)(0);
        };
    };
};
var bumpToken = function (v) {
    return function (v1) {
        if (v1 instanceof PureScript_CST_Types.TokLeftParen) {
            return {
                line: v.line,
                column: v.column + 1 | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokRightParen) {
            return {
                line: v.line,
                column: v.column + 1 | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokLeftBrace) {
            return {
                line: v.line,
                column: v.column + 1 | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokRightBrace) {
            return {
                line: v.line,
                column: v.column + 1 | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokLeftSquare) {
            return {
                line: v.line,
                column: v.column + 1 | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokRightSquare) {
            return {
                line: v.line,
                column: v.column + 1 | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokLeftArrow && v1.value0 instanceof PureScript_CST_Types.ASCII) {
            return {
                line: v.line,
                column: v.column + 2 | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokLeftArrow && v1.value0 instanceof PureScript_CST_Types.Unicode) {
            return {
                line: v.line,
                column: v.column + 1 | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokRightArrow && v1.value0 instanceof PureScript_CST_Types.ASCII) {
            return {
                line: v.line,
                column: v.column + 2 | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokRightArrow && v1.value0 instanceof PureScript_CST_Types.Unicode) {
            return {
                line: v.line,
                column: v.column + 1 | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokRightFatArrow && v1.value0 instanceof PureScript_CST_Types.ASCII) {
            return {
                line: v.line,
                column: v.column + 2 | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokRightFatArrow && v1.value0 instanceof PureScript_CST_Types.Unicode) {
            return {
                line: v.line,
                column: v.column + 1 | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokDoubleColon && v1.value0 instanceof PureScript_CST_Types.ASCII) {
            return {
                line: v.line,
                column: v.column + 2 | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokDoubleColon && v1.value0 instanceof PureScript_CST_Types.Unicode) {
            return {
                line: v.line,
                column: v.column + 1 | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokForall && v1.value0 instanceof PureScript_CST_Types.ASCII) {
            return {
                line: v.line,
                column: v.column + 6 | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokForall && v1.value0 instanceof PureScript_CST_Types.Unicode) {
            return {
                line: v.line,
                column: v.column + 1 | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokEquals) {
            return {
                line: v.line,
                column: v.column + 1 | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokPipe) {
            return {
                line: v.line,
                column: v.column + 1 | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokTick) {
            return {
                line: v.line,
                column: v.column + 1 | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokDot) {
            return {
                line: v.line,
                column: v.column + 1 | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokComma) {
            return {
                line: v.line,
                column: v.column + 1 | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokUnderscore) {
            return {
                line: v.line,
                column: v.column + 1 | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokBackslash) {
            return {
                line: v.line,
                column: v.column + 1 | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokAt) {
            return {
                line: v.line,
                column: v.column + 1 | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokLowerName) {
            return {
                line: v.line,
                column: (v.column + qualLength(v1.value0) | 0) + Data_String_CodePoints.length(v1.value1) | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokUpperName) {
            return {
                line: v.line,
                column: (v.column + qualLength(v1.value0) | 0) + Data_String_CodePoints.length(v1.value1) | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokOperator) {
            return {
                line: v.line,
                column: (v.column + qualLength(v1.value0) | 0) + Data_String_CodePoints.length(v1.value1) | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokSymbolName) {
            return {
                line: v.line,
                column: ((v.column + qualLength(v1.value0) | 0) + Data_String_CodePoints.length(v1.value1) | 0) + 2 | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokSymbolArrow && v1.value0 instanceof PureScript_CST_Types.Unicode) {
            return {
                line: v.line,
                column: v.column + 3 | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokSymbolArrow && v1.value0 instanceof PureScript_CST_Types.ASCII) {
            return {
                line: v.line,
                column: v.column + 4 | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokHole) {
            return {
                line: v.line,
                column: (v.column + Data_String_CodePoints.length(v1.value0) | 0) + 1 | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokChar) {
            return {
                line: v.line,
                column: (v.column + Data_String_CodePoints.length(v1.value0) | 0) + 2 | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokInt) {
            return {
                line: v.line,
                column: v.column + Data_String_CodePoints.length(v1.value0) | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokNumber) {
            return {
                line: v.line,
                column: v.column + Data_String_CodePoints.length(v1.value0) | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.TokString) {
            return bumpText(v)(1)(v1.value0);
        };
        if (v1 instanceof PureScript_CST_Types.TokRawString) {
            return bumpText(v)(3)(v1.value0);
        };
        if (v1 instanceof PureScript_CST_Types.TokLayoutStart) {
            return v;
        };
        if (v1 instanceof PureScript_CST_Types.TokLayoutSep) {
            return v;
        };
        if (v1 instanceof PureScript_CST_Types.TokLayoutEnd) {
            return v;
        };
        throw new Error("Failed pattern match at PureScript.CST.Lexer (line 265, column 34 - line 343, column 8): " + [ v1.constructor.name ]);
    };
};
var bumpComment = function (v) {
    return function (v1) {
        if (v1 instanceof PureScript_CST_Types.Comment) {
            return bumpText(v)(0)(v1.value0);
        };
        if (v1 instanceof PureScript_CST_Types.Space) {
            return {
                line: v.line,
                column: v.column + v1.value0 | 0
            };
        };
        if (v1 instanceof PureScript_CST_Types.Line) {
            return {
                line: v.line + v1.value1 | 0,
                column: 0
            };
        };
        throw new Error("Failed pattern match at PureScript.CST.Lexer (line 360, column 36 - line 366, column 34): " + [ v1.constructor.name ]);
    };
};
var applyLex = {
    apply: function (v) {
        return function (v1) {
            return function (str) {
                var v2 = v(str);
                if (v2 instanceof LexFail) {
                    return new LexFail(v2.value0, v2.value1);
                };
                if (v2 instanceof LexSucc) {
                    var v3 = v1(v2.value1);
                    if (v3 instanceof LexFail) {
                        return new LexFail(v3.value0, v3.value1);
                    };
                    if (v3 instanceof LexSucc) {
                        return new LexSucc(v2.value0(v3.value0), v3.value1);
                    };
                    throw new Error("Failed pattern match at PureScript.CST.Lexer (line 75, column 9 - line 78, column 32): " + [ v3.constructor.name ]);
                };
                throw new Error("Failed pattern match at PureScript.CST.Lexer (line 72, column 5 - line 78, column 32): " + [ v2.constructor.name ]);
            };
        };
    },
    Functor0: function () {
        return functorLex;
    }
};
var bindLex = {
    bind: function (v) {
        return function (k) {
            return function (str) {
                var v1 = v(str);
                if (v1 instanceof LexFail) {
                    return new LexFail(v1.value0, v1.value1);
                };
                if (v1 instanceof LexSucc) {
                    var v2 = k(v1.value0);
                    return v2(v1.value1);
                };
                throw new Error("Failed pattern match at PureScript.CST.Lexer (line 85, column 5 - line 89, column 16): " + [ v1.constructor.name ]);
            };
        };
    },
    Apply0: function () {
        return applyLex;
    }
};
var applicativeLex = {
    pure: function ($290) {
        return Lex(LexSucc.create($290));
    },
    Apply0: function () {
        return applyLex;
    }
};
var pure = /* #__PURE__ */ Control_Applicative.pure(applicativeLex);
var altLex = {
    alt: function (v) {
        return function (v1) {
            return function (str) {
                var v2 = v(str);
                if (v2 instanceof LexFail) {
                    if (Data_String_CodeUnits.length(str) === Data_String_CodeUnits.length(v2.value1)) {
                        return v1(str);
                    };
                    if (Data_Boolean.otherwise) {
                        return new LexFail(v2.value0, v2.value1);
                    };
                };
                if (v2 instanceof LexSucc) {
                    return new LexSucc(v2.value0, v2.value1);
                };
                throw new Error("Failed pattern match at PureScript.CST.Lexer (line 93, column 5 - line 100, column 20): " + [ v2.constructor.name ]);
            };
        };
    },
    Functor0: function () {
        return functorLex;
    }
};
var comment = /* #__PURE__ */ (function () {
    return Control_Alt.alt(altLex)(regex(PureScript_CST_Errors.LexExpected.create("block comment"))("\\{-(-(?!\\})|[^-]+)*(-\\}|$)"))(regex(PureScript_CST_Errors.LexExpected.create("line comment"))("--[^\\r\\n]*"));
})();
var lineComment = /* #__PURE__ */ (function () {
    return Control_Alt.alt(altLex)(Data_Functor.map(functorLex)((function () {
        var $291 = PureScript_CST_Types.Line.create(PureScript_CST_Types.LF.value);
        return function ($292) {
            return $291(Data_String_CodePoints.length($292));
        };
    })())(regex(PureScript_CST_Errors.LexExpected.create("newline"))("\x0a+")))(Data_Functor.map(functorLex)((function () {
        var $293 = PureScript_CST_Types.Line.create(PureScript_CST_Types.CRLF.value);
        return function ($294) {
            return $293((function (v) {
                return Data_EuclideanRing.div(Data_EuclideanRing.euclideanRingInt)(v)(2);
            })(Data_String_CodePoints.length($294)));
        };
    })())(regex(PureScript_CST_Errors.LexExpected.create("newline"))("(?:\x0d\x0a)+")));
})();
var leadingComments = /* #__PURE__ */ (function () {
    return many(Control_Alt.alt(altLex)(Data_Functor.map(functorLex)(PureScript_CST_Types.Comment.create)(comment))(Control_Alt.alt(altLex)(Data_Functor.map(functorLex)(PureScript_CST_Types.Space.create)(spaceComment))(lineComment)));
})();
var oneLineComment = /* #__PURE__ */ Control_Bind.bind(bindLex)(lineComment)(function (line) {
    if (line instanceof PureScript_CST_Types.Line && line.value1 === 1) {
        return Control_Applicative.pure(applicativeLex)(line);
    };
    return fail(new PureScript_CST_Errors.LexExpected("one newline", "multiple newlines"));
});
var leadingShebangs = /* #__PURE__ */ (function () {
    return Control_Apply.apply(applyLex)(Data_Functor.map(functorLex)(function (v) {
        return function (v1) {
            return Data_Array.cons(new PureScript_CST_Types.Comment(v))(Data_Foldable.foldMap(Data_Foldable.foldableArray)(Data_Monoid.monoidArray)(function (v2) {
                return [ v2.value0, new PureScript_CST_Types.Comment(v2.value1) ];
            })(v1));
        };
    })(shebangComment))(many($$try(Control_Apply.apply(applyLex)(Data_Functor.map(functorLex)(Data_Tuple.Tuple.create)(oneLineComment))(shebangComment))));
})();
var leadingModuleComments = /* #__PURE__ */ Control_Apply.apply(applyLex)(/* #__PURE__ */ Data_Functor.map(functorLex)(/* #__PURE__ */ Data_Semigroup.append(Data_Semigroup.semigroupArray))(/* #__PURE__ */ Control_Alt.alt(altLex)(leadingShebangs)(/* #__PURE__ */ Control_Applicative.pure(applicativeLex)([  ]))))(leadingComments);
var token = /* #__PURE__ */ (function () {
    var tokenTick = char$prime(PureScript_CST_Errors.LexExpected.create("backtick"))(PureScript_CST_Types.TokTick.value)("`");
    var tokenRightSquare = char$prime(PureScript_CST_Errors.LexExpected.create("right square"))(PureScript_CST_Types.TokRightSquare.value)("]");
    var tokenRightParen = char$prime(PureScript_CST_Errors.LexExpected.create("right paren"))(PureScript_CST_Types.TokRightParen.value)(")");
    var tokenRightBrace = char$prime(PureScript_CST_Errors.LexExpected.create("right brace"))(PureScript_CST_Types.TokRightBrace.value)("}");
    var tokenLeftSquare = char$prime(PureScript_CST_Errors.LexExpected.create("left square"))(PureScript_CST_Types.TokLeftSquare.value)("[");
    var tokenLeftParen = char$prime(PureScript_CST_Errors.LexExpected.create("left paren"))(PureScript_CST_Types.TokLeftParen.value)("(");
    var tokenLeftBrace = char$prime(PureScript_CST_Errors.LexExpected.create("left brace"))(PureScript_CST_Types.TokLeftBrace.value)("{");
    var tokenComma = char$prime(PureScript_CST_Errors.LexExpected.create("comma"))(PureScript_CST_Types.TokComma.value)(",");
    var stripUnderscores = Data_String_Common.replaceAll("_")("");
    var stringSpaceEscapeRegex = regex(PureScript_CST_Errors.LexExpected.create("whitespace escape"))("\\\\[ \\r\\n]+\\\\");
    var stringCharsRegex = regex(PureScript_CST_Errors.LexExpected.create("string characters"))("[^\"\\\\]+");
    var rawStringCharsRegex = regex(PureScript_CST_Errors.LexExpected.create("raw string characters"))("\"\"\"\"{0,2}([^\"]+\"{1,2})*[^\"]*\"\"\"");
    var parseSymbolIdent = regex(PureScript_CST_Errors.LexExpected.create("symbol"))("(?:[:!#$%&*+./<=>?@\\\\^|~-]|(?!\\p{P})\\p{S})+");
    var parseSymbol = Data_Functor.map(functorLex)(function (v) {
        return function (v1) {
            if (v1 instanceof Data_Maybe.Nothing) {
                if (v === "->") {
                    return new PureScript_CST_Types.TokSymbolArrow(PureScript_CST_Types.ASCII.value);
                };
                if (v === "\u2192") {
                    return new PureScript_CST_Types.TokSymbolArrow(PureScript_CST_Types.Unicode.value);
                };
                return new PureScript_CST_Types.TokSymbolName(Data_Maybe.Nothing.value, v);
            };
            return new PureScript_CST_Types.TokSymbolName(v1, v);
        };
    })($$try(Control_Apply.applyFirst(applyLex)(Control_Apply.applySecond(applyLex)(tokenLeftParen)(parseSymbolIdent))(tokenRightParen)));
    var parseStringSpaceEscape = Data_Functor.map(functorLex)(function (v) {
        return {
            raw: v,
            string: ""
        };
    })(stringSpaceEscapeRegex);
    var parseStringChars = Data_Functor.map(functorLex)(function (v) {
        return {
            raw: v,
            string: v
        };
    })(stringCharsRegex);
    var parseRawString = Data_Functor.map(functorLex)(function (v) {
        return new PureScript_CST_Types.TokRawString(Data_String_CodeUnits.dropRight(3)(Data_String_CodeUnits.drop(3)(v)));
    })(rawStringCharsRegex);
    var parseProper = regex(PureScript_CST_Errors.LexExpected.create("proper name"))("\\p{Lu}[\\p{L}0-9_']*");
    var parseUpper = Data_Functor.map(functorLex)(Data_Function.flip(PureScript_CST_Types.TokUpperName.create))(parseProper);
    var parseOperator = Data_Functor.map(functorLex)(function (v) {
        return function (v1) {
            if (v1 instanceof Data_Maybe.Nothing) {
                if (v === "<-") {
                    return new PureScript_CST_Types.TokLeftArrow(PureScript_CST_Types.ASCII.value);
                };
                if (v === "\u2190") {
                    return new PureScript_CST_Types.TokLeftArrow(PureScript_CST_Types.Unicode.value);
                };
                if (v === "->") {
                    return new PureScript_CST_Types.TokRightArrow(PureScript_CST_Types.ASCII.value);
                };
                if (v === "\u2192") {
                    return new PureScript_CST_Types.TokRightArrow(PureScript_CST_Types.Unicode.value);
                };
                if (v === "=>") {
                    return new PureScript_CST_Types.TokRightFatArrow(PureScript_CST_Types.ASCII.value);
                };
                if (v === "\u21d2") {
                    return new PureScript_CST_Types.TokRightFatArrow(PureScript_CST_Types.Unicode.value);
                };
                if (v === "::") {
                    return new PureScript_CST_Types.TokDoubleColon(PureScript_CST_Types.ASCII.value);
                };
                if (v === "\u2237") {
                    return new PureScript_CST_Types.TokDoubleColon(PureScript_CST_Types.Unicode.value);
                };
                if (v === "\u2200") {
                    return new PureScript_CST_Types.TokForall(PureScript_CST_Types.Unicode.value);
                };
                if (v === "=") {
                    return PureScript_CST_Types.TokEquals.value;
                };
                if (v === ".") {
                    return PureScript_CST_Types.TokDot.value;
                };
                if (v === "\\") {
                    return PureScript_CST_Types.TokBackslash.value;
                };
                if (v === "|") {
                    return PureScript_CST_Types.TokPipe.value;
                };
                if (v === "@") {
                    return PureScript_CST_Types.TokAt.value;
                };
                if (v === "`") {
                    return PureScript_CST_Types.TokTick.value;
                };
                return new PureScript_CST_Types.TokOperator(Data_Maybe.Nothing.value, v);
            };
            return new PureScript_CST_Types.TokOperator(v1, v);
        };
    })(parseSymbolIdent);
    var parseModuleNamePrefix = regex(PureScript_CST_Errors.LexExpected.create("module name"))("(?:(?:\\p{Lu}[\\p{L}0-9_']*)\\.)*");
    var parseIdent = regex(PureScript_CST_Errors.LexExpected.create("ident"))("[\\p{Ll}_][\\p{L}0-9_']*");
    var parseLower = Data_Functor.map(functorLex)(function (v) {
        return function (v1) {
            if (v1 instanceof Data_Maybe.Nothing) {
                if (v === "forall") {
                    return new PureScript_CST_Types.TokForall(PureScript_CST_Types.ASCII.value);
                };
                if (v === "_") {
                    return PureScript_CST_Types.TokUnderscore.value;
                };
                return new PureScript_CST_Types.TokLowerName(Data_Maybe.Nothing.value, v);
            };
            return new PureScript_CST_Types.TokLowerName(v1, v);
        };
    })(parseIdent);
    var parseName = Control_Alt.alt(altLex)(parseLower)(Control_Alt.alt(altLex)(parseUpper)(Control_Alt.alt(altLex)(parseOperator)(parseSymbol)));
    var parseModuleName = Control_Apply.apply(applyLex)(Data_Functor.map(functorLex)(function (v) {
        return function (v1) {
            return v1(toModuleName(v));
        };
    })(parseModuleNamePrefix))(parseName);
    var parseExponentSign = Control_Alt.alt(altLex)(string(PureScript_CST_Errors.LexExpected.create("negative"))("-"))(string(PureScript_CST_Errors.LexExpected.create("positive"))("+"));
    var intPartRegex = regex(PureScript_CST_Errors.LexExpected.create("int part"))("(0|[1-9][0-9_]*)");
    var parseExponentPart = Control_Apply.apply(applyLex)(Data_Functor.map(functorLex)(function (v) {
        return function (v1) {
            return {
                sign: v,
                exponent: v1
            };
        };
    })(optional(parseExponentSign)))(intPartRegex);
    var hexIntRegex = regex(PureScript_CST_Errors.LexExpected.create("hex int"))("[a-fA-F0-9]+");
    var hexIntPrefix = string(PureScript_CST_Errors.LexExpected.create("hex int prefix"))("0x");
    var parseHexInt = Control_Bind.bind(bindLex)(Control_Apply.applySecond(applyLex)(hexIntPrefix)(hexIntRegex))(function (raw) {
        var v = Data_Int.fromStringAs(Data_Int.hexadecimal)(raw);
        if (v instanceof Data_Maybe.Just) {
            return pure(new PureScript_CST_Types.TokInt("0x" + raw, new PureScript_CST_Types.SmallInt(v.value0)));
        };
        if (v instanceof Data_Maybe.Nothing) {
            return pure(new PureScript_CST_Types.TokInt("0x" + raw, new PureScript_CST_Types.BigHex(raw)));
        };
        throw new Error("Failed pattern match at PureScript.CST.Lexer (line 631, column 5 - line 635, column 49): " + [ v.constructor.name ]);
    });
    var hexEscapeRegex = regex(PureScript_CST_Errors.LexExpected.create("hex"))("[a-fA-F0-9]{1,6}");
    var parseHexEscape = function (dictIsChar) {
        var fromCharCode1 = fromCharCode(dictIsChar);
        return Control_Bind.bind(bindLex)(hexEscapeRegex)(function (esc) {
            var v = Control_Bind.bindFlipped(Data_Maybe.bindMaybe)(fromCharCode1)(Data_Int.fromStringAs(Data_Int.hexadecimal)(esc));
            if (v instanceof Data_Maybe.Just) {
                return Control_Applicative.pure(applicativeLex)({
                    raw: "\\x" + esc,
                    "char": v.value0
                });
            };
            if (v instanceof Data_Maybe.Nothing) {
                return fail(new PureScript_CST_Errors.LexCharEscapeOutOfRange(esc));
            };
            throw new Error("Failed pattern match at PureScript.CST.Lexer (line 579, column 5 - line 583, column 43): " + [ v.constructor.name ]);
        });
    };
    var fractionPartRegex = regex(PureScript_CST_Errors.LexExpected.create("fraction part"))("[0-9_]+");
    var charSingleQuote = $$char(PureScript_CST_Errors.LexExpected.create("single quote"))("'");
    var charQuote = $$char(PureScript_CST_Errors.LexExpected.create("quote"))("\"");
    var charQuestionMark = $$char(PureScript_CST_Errors.LexExpected.create("question mark"))("?");
    var parseHole = Data_Functor.map(functorLex)(function (v) {
        return new PureScript_CST_Types.TokHole(v);
    })($$try(Control_Apply.applySecond(applyLex)(charQuestionMark)(Control_Alt.alt(altLex)(parseIdent)(parseProper))));
    var charExponent = $$char(PureScript_CST_Errors.LexExpected.create("exponent"))("e");
    var parseNumberExponentPart = optional(Control_Apply.applySecond(applyLex)(charExponent)(parseExponentPart));
    var charDot = $$char(PureScript_CST_Errors.LexExpected.create("dot"))(".");
    var parseNumberFractionPart = optional($$try(Control_Apply.applySecond(applyLex)(charDot)(fractionPartRegex)));
    var parseNumber = Control_Bind.bind(bindLex)(intPartRegex)(function (intPart) {
        return Control_Bind.bind(bindLex)(parseNumberFractionPart)(function (fractionPart) {
            return Control_Bind.bind(bindLex)(parseNumberExponentPart)(function (exponentPart) {
                var $249 = Data_Maybe.isNothing(fractionPart) && Data_Maybe.isNothing(exponentPart);
                if ($249) {
                    var intVal = stripUnderscores(intPart);
                    var v = Data_Int.fromString(intVal);
                    if (v instanceof Data_Maybe.Just) {
                        return pure(new PureScript_CST_Types.TokInt(intPart, new PureScript_CST_Types.SmallInt(v.value0)));
                    };
                    if (v instanceof Data_Maybe.Nothing) {
                        return pure(new PureScript_CST_Types.TokInt(intPart, new PureScript_CST_Types.BigInt(intVal)));
                    };
                    throw new Error("Failed pattern match at PureScript.CST.Lexer (line 643, column 7 - line 647, column 48): " + [ v.constructor.name ]);
                };
                var raw = intPart + (Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(Data_Monoid.monoidString)(function (fr) {
                    return "." + fr;
                })(fractionPart) + Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(Data_Monoid.monoidString)(function (ex) {
                    return "e" + (Data_Foldable.fold(Data_Foldable.foldableMaybe)(Data_Monoid.monoidString)(ex.sign) + ex.exponent);
                })(exponentPart));
                var v = Data_Number.fromString(stripUnderscores(raw));
                if (v instanceof Data_Maybe.Just) {
                    return pure(new PureScript_CST_Types.TokNumber(raw, v.value0));
                };
                if (v instanceof Data_Maybe.Nothing) {
                    return fail(new PureScript_CST_Errors.LexNumberOutOfRange(raw));
                };
                throw new Error("Failed pattern match at PureScript.CST.Lexer (line 654, column 7 - line 658, column 41): " + [ v.constructor.name ]);
            });
        });
    });
    var parseNumericLiteral = Control_Alt.alt(altLex)(parseHexInt)(parseNumber);
    var charBackslash = $$char(PureScript_CST_Errors.LexExpected.create("backslash"))("\\");
    var charAny = satisfy(PureScript_CST_Errors.LexExpected.create("char"))(Data_Function["const"](true));
    var parseEscape = function (dictIsChar) {
        var parseHexEscape1 = parseHexEscape(dictIsChar);
        return Control_Bind.bind(bindLex)(charAny)(function (ch) {
            if (ch === "t") {
                return Control_Applicative.pure(applicativeLex)({
                    raw: "\\t",
                    "char": fromChar(dictIsChar)("\x09")
                });
            };
            if (ch === "r") {
                return Control_Applicative.pure(applicativeLex)({
                    raw: "\\r",
                    "char": fromChar(dictIsChar)("\x0d")
                });
            };
            if (ch === "n") {
                return Control_Applicative.pure(applicativeLex)({
                    raw: "\\n",
                    "char": fromChar(dictIsChar)("\x0a")
                });
            };
            if (ch === "\"") {
                return Control_Applicative.pure(applicativeLex)({
                    raw: "\\\"",
                    "char": fromChar(dictIsChar)("\"")
                });
            };
            if (ch === "'") {
                return Control_Applicative.pure(applicativeLex)({
                    raw: "\\'",
                    "char": fromChar(dictIsChar)("'")
                });
            };
            if (ch === "\\") {
                return Control_Applicative.pure(applicativeLex)({
                    raw: "\\\\",
                    "char": fromChar(dictIsChar)("\\")
                });
            };
            if (ch === "x") {
                return parseHexEscape1;
            };
            return fail(new PureScript_CST_Errors.LexInvalidCharEscape(Data_String_CodeUnits.singleton(ch)));
        });
    };
    var parseEscape1 = parseEscape(isCharChar);
    var parseChar = Control_Bind.bind(bindLex)(charAny)(function (ch) {
        if (ch === "\\") {
            return parseEscape1;
        };
        if (ch === "'") {
            return fail(new PureScript_CST_Errors.LexExpected("character", "empty character literal"));
        };
        return Control_Applicative.pure(applicativeLex)({
            raw: Data_String_CodeUnits.singleton(ch),
            "char": ch
        });
    });
    var parseCharLiteral = Data_Functor.map(functorLex)(function (v) {
        return new PureScript_CST_Types.TokChar(v.raw, v["char"]);
    })(Control_Apply.applyFirst(applyLex)(Control_Apply.applySecond(applyLex)(charSingleQuote)(parseChar))(charSingleQuote));
    var parseStringEscape = Data_Functor.map(functorLex)(function (v) {
        return {
            raw: v.raw,
            string: Data_String_CodePoints.singleton(v["char"])
        };
    })(Control_Apply.applySecond(applyLex)(charBackslash)(parseEscape(isCharCodePoint)));
    var parseStringPart = Control_Alt.alt(altLex)(parseStringChars)(Control_Alt.alt(altLex)(parseStringSpaceEscape)(parseStringEscape));
    var parseString = Data_Functor.map(functorLex)(function (v) {
        var v1 = Data_Foldable.fold(Data_Foldable.foldableArray)(monoidRecord)(v);
        return new PureScript_CST_Types.TokString(v1.raw, v1.string);
    })(Control_Apply.applyFirst(applyLex)(Control_Apply.applySecond(applyLex)(charQuote)(many(parseStringPart)))(charQuote));
    var parseStringLiteral = Control_Alt.alt(altLex)(parseRawString)(parseString);
    return Control_Alt.alt(altLex)(parseHole)(Control_Alt.alt(altLex)(parseModuleName)(Control_Alt.alt(altLex)(parseCharLiteral)(Control_Alt.alt(altLex)(parseStringLiteral)(Control_Alt.alt(altLex)(parseNumericLiteral)(Control_Alt.alt(altLex)(tokenLeftParen)(Control_Alt.alt(altLex)(tokenRightParen)(Control_Alt.alt(altLex)(tokenLeftBrace)(Control_Alt.alt(altLex)(tokenRightBrace)(Control_Alt.alt(altLex)(tokenLeftSquare)(Control_Alt.alt(altLex)(tokenRightSquare)(Control_Alt.alt(altLex)(tokenTick)(tokenComma))))))))))));
})();
var lexToken = function ($295) {
    return (function (v) {
        if (v instanceof LexSucc && v.value1 === "") {
            return new Data_Either.Right(v.value0);
        };
        if (v instanceof LexSucc) {
            return new Data_Either.Left(function (v1) {
                return new PureScript_CST_Errors.ExpectedEof(v.value0);
            });
        };
        if (v instanceof LexFail) {
            return new Data_Either.Left(v.value0);
        };
        throw new Error("Failed pattern match at PureScript.CST.Lexer (line 257, column 18 - line 260, column 28): " + [ v.constructor.name ]);
    })(token($295));
};
var trailingComments = /* #__PURE__ */ (function () {
    return many(Control_Alt.alt(altLex)(Data_Functor.map(functorLex)(PureScript_CST_Types.Comment.create)(comment))(Data_Functor.map(functorLex)(PureScript_CST_Types.Space.create)(spaceComment)));
})();
var lexWithState$prime = function (lexLeadingComments) {
    var token$prime = Control_Apply.apply(applyLex)(Control_Apply.apply(applyLex)(Data_Functor.map(functorLex)(function (v) {
        return function (v1) {
            return function (v2) {
                return {
                    token: v,
                    trailing: v1,
                    nextLeading: v2
                };
            };
        };
    })(token))(trailingComments))(leadingComments);
    var go = function (stack) {
        return function (startPos) {
            return function (leading) {
                return function (str) {
                    return Data_Lazy.defer(function (v) {
                        var $270 = str === "";
                        if ($270) {
                            return PureScript_CST_TokenStream.step(PureScript_CST_TokenStream.unwindLayout(startPos)(Data_Lazy.defer(function (v1) {
                                return new PureScript_CST_TokenStream.TokenEOF(startPos, leading);
                            }))(stack));
                        };
                        var v1 = token$prime(str);
                        if (v1 instanceof LexFail) {
                            var errPos = bumpText(startPos)(0)(Data_String_CodeUnits.take(Data_String_CodeUnits.length(str) - Data_String_CodeUnits.length(v1.value1) | 0)(str));
                            return new PureScript_CST_TokenStream.TokenError(errPos, v1.value0(Data_Unit.unit), Data_Maybe.Nothing.value, stack);
                        };
                        if (v1 instanceof LexSucc) {
                            var endPos = bumpToken(startPos)(v1.value0.token);
                            var nextStart = Data_Foldable.foldl(Data_Foldable.foldableArray)(bumpComment)(Data_Foldable.foldl(Data_Foldable.foldableArray)(bumpComment)(endPos)(v1.value0.trailing))(v1.value0.nextLeading);
                            var posToken = {
                                range: {
                                    start: startPos,
                                    end: endPos
                                },
                                leadingComments: leading,
                                trailingComments: v1.value0.trailing,
                                value: v1.value0.token
                            };
                            var v2 = PureScript_CST_Layout.insertLayout(posToken)(nextStart)(stack);
                            return PureScript_CST_TokenStream.step(Data_Tuple.snd(PureScript_CST_TokenStream.consTokens(Data_Foldable.foldableArray)(v2.value1)(new Data_Tuple.Tuple(nextStart, go(v2.value0)(nextStart)(v1.value0.nextLeading)(v1.value1)))));
                        };
                        throw new Error("Failed pattern match at PureScript.CST.Lexer (line 228, column 7 - line 247, column 63): " + [ v1.constructor.name ]);
                    });
                };
            };
        };
    };
    var init = function (initStack) {
        return function (initPos) {
            return function (str) {
                return Data_Lazy.defer(function (v) {
                    var v1 = lexLeadingComments(str);
                    if (v1 instanceof LexFail) {
                        return Partial_Unsafe.unsafeCrashWith("Leading comments can't fail.");
                    };
                    if (v1 instanceof LexSucc) {
                        var nextPos = Data_Foldable.foldl(Data_Foldable.foldableArray)(bumpComment)(initPos)(v1.value0);
                        return PureScript_CST_TokenStream.step(go(initStack)(nextPos)(v1.value0)(v1.value1));
                    };
                    throw new Error("Failed pattern match at PureScript.CST.Lexer (line 215, column 5 - line 220, column 51): " + [ v1.constructor.name ]);
                });
            };
        };
    };
    return init;
};

// | Lexes according to root layout rules as well as supporting leading shebang comments.
var lexModule = /* #__PURE__ */ (function () {
    return lexWithState$prime(leadingModuleComments)(new Data_List_Types.Cons(new Data_Tuple.Tuple({
        line: 0,
        column: 0
    }, PureScript_CST_Layout.LytRoot.value), Data_List_Types.Nil.value))({
        line: 0,
        column: 0
    });
})();
var lexWithState = /* #__PURE__ */ lexWithState$prime(leadingComments);

// | Lexes according to root layout rules and standard language comments.
var lex = /* #__PURE__ */ (function () {
    return lexWithState(new Data_List_Types.Cons(new Data_Tuple.Tuple({
        line: 0,
        column: 0
    }, PureScript_CST_Layout.LytRoot.value), Data_List_Types.Nil.value))({
        line: 0,
        column: 0
    });
})();
export {
    lex,
    lexModule,
    lexWithState,
    lexToken
};
