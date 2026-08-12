import * as Control_Alt from "../Control.Alt/index.js";
import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Lazy from "../Control.Lazy/index.js";
import * as Data_Array from "../Data.Array/index.js";
import * as Data_Array_NonEmpty from "../Data.Array.NonEmpty/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Set from "../Data.Set/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as PureScript_CST_Errors from "../PureScript.CST.Errors/index.js";
import * as PureScript_CST_Layout from "../PureScript.CST.Layout/index.js";
import * as PureScript_CST_Parser_Monad from "../PureScript.CST.Parser.Monad/index.js";
import * as PureScript_CST_TokenStream from "../PureScript.CST.TokenStream/index.js";
import * as PureScript_CST_Types from "../PureScript.CST.Types/index.js";
var $runtime_lazy = function (name, moduleName, init) {
    var state = 0;
    var val;
    return function (lineNumber) {
        if (state === 2) return val;
        if (state === 1) throw new ReferenceError(name + " was needed before it finished initializing (module " + moduleName + ", line " + lineNumber + ")", moduleName, lineNumber);
        state = 1;
        val = init();
        state = 2;
        return val;
    };
};
var pure = /* #__PURE__ */ Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser);
var pure1 = /* #__PURE__ */ Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser);
var pure2 = /* #__PURE__ */ Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser);
var pure3 = /* #__PURE__ */ Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser);
var pure4 = /* #__PURE__ */ Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser);
var pure5 = /* #__PURE__ */ Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser);
var pure6 = /* #__PURE__ */ Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser);
var pure7 = /* #__PURE__ */ Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser);
var pure8 = /* #__PURE__ */ Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser);
var pure9 = /* #__PURE__ */ Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser);
var pure10 = /* #__PURE__ */ Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser);
var pure11 = /* #__PURE__ */ Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser);
var pure12 = /* #__PURE__ */ Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser);
var pure13 = /* #__PURE__ */ Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser);
var pure14 = /* #__PURE__ */ Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser);
var pure15 = /* #__PURE__ */ Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser);
var pure16 = /* #__PURE__ */ Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser);
var pure17 = /* #__PURE__ */ Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser);
var pure18 = /* #__PURE__ */ Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser);
var wrapped = function (openTok) {
    return function (closeTok) {
        return function (valueParser) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(openTok)(function (open) {
                return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(valueParser)(function (value) {
                    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(closeTok)(function (close) {
                        return pure({
                            open: open,
                            value: value,
                            close: close
                        });
                    });
                });
            });
        };
    };
};
var separated = function (sepParser) {
    return function (valueParser) {
        return Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(function (v) {
            return function (v1) {
                return {
                    head: v,
                    tail: v1
                };
            };
        })(valueParser))(PureScript_CST_Parser_Monad.many(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.Tuple.create)(sepParser))(valueParser)));
    };
};
var reservedKeywords = /* #__PURE__ */ Data_Set.fromFoldable(Data_Foldable.foldableArray)(Data_Ord.ordString)([ "ado", "case", "class", "data", "derive", "do", "else", "false", "foreign", "if", "import", "in", "infix", "infixl", "infixr", "instance", "let", "module", "newtype", "of", "then", "true", "type", "where" ]);
var recoverTokensWhile = function (p) {
    return function (initStream) {
        var indent = Data_Maybe.maybe(0)(function (v) {
            return v.column;
        })(PureScript_CST_Layout.currentIndent(PureScript_CST_TokenStream.layoutStack(initStream)));
        var go = function ($copy_acc) {
            return function ($copy_stream) {
                var $tco_var_acc = $copy_acc;
                var $tco_done = false;
                var $tco_result;
                function $tco_loop(acc, stream) {
                    var v = PureScript_CST_TokenStream.step(stream);
                    if (v instanceof PureScript_CST_TokenStream.TokenError) {
                        $tco_done = true;
                        return new Data_Tuple.Tuple(acc, stream);
                    };
                    if (v instanceof PureScript_CST_TokenStream.TokenEOF) {
                        $tco_done = true;
                        return new Data_Tuple.Tuple(acc, stream);
                    };
                    if (v instanceof PureScript_CST_TokenStream.TokenCons) {
                        var $138 = p(v.value0)(indent);
                        if ($138) {
                            $tco_var_acc = Data_Array.snoc(acc)(v.value0);
                            $copy_stream = v.value2;
                            return;
                        };
                        $tco_done = true;
                        return new Data_Tuple.Tuple(acc, stream);
                    };
                    throw new Error("Failed pattern match at PureScript.CST.Parser (line 1194, column 19 - line 1203, column 25): " + [ v.constructor.name ]);
                };
                while (!$tco_done) {
                    $tco_result = $tco_loop($tco_var_acc, $copy_stream);
                };
                return $tco_result;
            };
        };
        return go([  ])(initStream);
    };
};
var recoverIndent = function (mkNode) {
    return PureScript_CST_Parser_Monad.recover(function (v) {
        return function (stream) {
            var v1 = recoverTokensWhile(function (tok) {
                return function (indent) {
                    if (tok.value instanceof PureScript_CST_Types.TokLayoutEnd) {
                        return tok.value.value0 > indent;
                    };
                    if (tok.value instanceof PureScript_CST_Types.TokLayoutSep) {
                        return tok.value.value0 > indent;
                    };
                    return true;
                };
            })(stream);
            var $148 = Data_Array["null"](v1.value0);
            if ($148) {
                return Data_Maybe.Nothing.value;
            };
            return new Data_Maybe.Just(new Data_Tuple.Tuple(mkNode({
                position: v.position,
                error: v.error,
                tokens: v1.value0
            }), v1.value1));
        };
    });
};
var recoverLetBinding = /* #__PURE__ */ (function () {
    return recoverIndent(PureScript_CST_Types.LetBindingError.create);
})();
var recoverDoStatement = /* #__PURE__ */ (function () {
    return recoverIndent(PureScript_CST_Types.DoError.create);
})();
var recoverDecl = /* #__PURE__ */ (function () {
    return recoverIndent(PureScript_CST_Types.DeclError.create);
})();
var parseSmallInt = /* #__PURE__ */ PureScript_CST_Parser_Monad.take(function (v) {
    if (v.value instanceof PureScript_CST_Types.TokInt && v.value.value1 instanceof PureScript_CST_Types.SmallInt) {
        return new Data_Either.Right(new Data_Tuple.Tuple(v, v.value.value1.value0));
    };
    if (v.value instanceof PureScript_CST_Types.TokInt) {
        return new Data_Either.Left(new PureScript_CST_Errors.LexIntOutOfRange(v.value.value0));
    };
    return new Data_Either.Left(new PureScript_CST_Errors.UnexpectedToken(v.value));
});
var many1 = function (parser) {
    return Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Array_NonEmpty["cons$prime"])(parser))(PureScript_CST_Parser_Monad.many(parser));
};
var expectMap = function (k) {
    return PureScript_CST_Parser_Monad.take(function (tok) {
        var v = k(tok);
        if (v instanceof Data_Maybe.Just) {
            return new Data_Either.Right(v.value0);
        };
        if (v instanceof Data_Maybe.Nothing) {
            return new Data_Either.Left(new PureScript_CST_Errors.UnexpectedToken(tok.value));
        };
        throw new Error("Failed pattern match at PureScript.CST.Parser (line 46, column 3 - line 50, column 39): " + [ v.constructor.name ]);
    });
};
var parseBoolean = /* #__PURE__ */ expectMap(function (v) {
    if (v.value instanceof PureScript_CST_Types.TokLowerName && (v.value.value0 instanceof Data_Maybe.Nothing && v.value.value1 === "true")) {
        return new Data_Maybe.Just(new Data_Tuple.Tuple(v, true));
    };
    if (v.value instanceof PureScript_CST_Types.TokLowerName && (v.value.value0 instanceof Data_Maybe.Nothing && v.value.value1 === "false")) {
        return new Data_Maybe.Just(new Data_Tuple.Tuple(v, false));
    };
    return Data_Maybe.Nothing.value;
});
var parseChar = /* #__PURE__ */ expectMap(function (v) {
    if (v.value instanceof PureScript_CST_Types.TokChar) {
        return new Data_Maybe.Just(new Data_Tuple.Tuple(v, v.value.value1));
    };
    return Data_Maybe.Nothing.value;
});
var parseHole = /* #__PURE__ */ expectMap(function (v) {
    if (v.value instanceof PureScript_CST_Types.TokHole) {
        return new Data_Maybe.Just({
            token: v,
            name: v.value.value0
        });
    };
    return Data_Maybe.Nothing.value;
});
var parseIdent = /* #__PURE__ */ expectMap(function (v) {
    if (v.value instanceof PureScript_CST_Types.TokLowerName && (v.value.value0 instanceof Data_Maybe.Nothing && !Data_Set.member(Data_Ord.ordString)(v.value.value1)(reservedKeywords))) {
        return new Data_Maybe.Just({
            token: v,
            name: v.value.value1
        });
    };
    return Data_Maybe.Nothing.value;
});
var $$parseInt = /* #__PURE__ */ expectMap(function (v) {
    if (v.value instanceof PureScript_CST_Types.TokInt) {
        return new Data_Maybe.Just(new Data_Tuple.Tuple(v, v.value.value1));
    };
    return Data_Maybe.Nothing.value;
});
var parseLabel = /* #__PURE__ */ expectMap(function (v) {
    if (v.value instanceof PureScript_CST_Types.TokRawString) {
        return new Data_Maybe.Just({
            token: v,
            name: v.value.value0
        });
    };
    if (v.value instanceof PureScript_CST_Types.TokString) {
        return new Data_Maybe.Just({
            token: v,
            name: v.value.value1
        });
    };
    if (v.value instanceof PureScript_CST_Types.TokLowerName && v.value.value0 instanceof Data_Maybe.Nothing) {
        return new Data_Maybe.Just({
            token: v,
            name: v.value.value1
        });
    };
    return Data_Maybe.Nothing.value;
});
var parseModuleName = /* #__PURE__ */ expectMap(function (v) {
    if (v.value instanceof PureScript_CST_Types.TokUpperName && v.value.value0 instanceof Data_Maybe.Just) {
        return new Data_Maybe.Just({
            token: v,
            name: v.value.value0.value0 + ("." + v.value.value1)
        });
    };
    if (v.value instanceof PureScript_CST_Types.TokUpperName && v.value.value0 instanceof Data_Maybe.Nothing) {
        return new Data_Maybe.Just({
            token: v,
            name: v.value.value1
        });
    };
    return Data_Maybe.Nothing.value;
});
var parseNumber = /* #__PURE__ */ expectMap(function (v) {
    if (v.value instanceof PureScript_CST_Types.TokNumber) {
        return new Data_Maybe.Just(new Data_Tuple.Tuple(v, v.value.value1));
    };
    return Data_Maybe.Nothing.value;
});
var parseOperator = /* #__PURE__ */ expectMap(function (v) {
    if (v.value instanceof PureScript_CST_Types.TokOperator && v.value.value0 instanceof Data_Maybe.Nothing) {
        return new Data_Maybe.Just({
            token: v,
            name: v.value.value1
        });
    };
    return Data_Maybe.Nothing.value;
});
var parseProper = /* #__PURE__ */ expectMap(function (v) {
    if (v.value instanceof PureScript_CST_Types.TokUpperName && v.value.value0 instanceof Data_Maybe.Nothing) {
        return new Data_Maybe.Just({
            token: v,
            name: v.value.value1
        });
    };
    return Data_Maybe.Nothing.value;
});
var parseQualifiedIdent = /* #__PURE__ */ expectMap(function (v) {
    if (v.value instanceof PureScript_CST_Types.TokLowerName && !Data_Set.member(Data_Ord.ordString)(v.value.value1)(reservedKeywords)) {
        return new Data_Maybe.Just({
            token: v,
            module: v.value.value0,
            name: v.value.value1
        });
    };
    return Data_Maybe.Nothing.value;
});
var parseQualifiedIdentOrProper = /* #__PURE__ */ expectMap(function (v) {
    if (v.value instanceof PureScript_CST_Types.TokLowerName) {
        return new Data_Maybe.Just({
            token: v,
            module: v.value.value0,
            name: new Data_Either.Left(v.value.value1)
        });
    };
    if (v.value instanceof PureScript_CST_Types.TokUpperName) {
        return new Data_Maybe.Just({
            token: v,
            module: v.value.value0,
            name: new Data_Either.Right(v.value.value1)
        });
    };
    return Data_Maybe.Nothing.value;
});
var parseQualifiedOperator = /* #__PURE__ */ expectMap(function (v) {
    if (v.value instanceof PureScript_CST_Types.TokOperator) {
        return new Data_Maybe.Just({
            token: v,
            module: v.value.value0,
            name: v.value.value1
        });
    };
    return Data_Maybe.Nothing.value;
});
var parseQualifiedProper = /* #__PURE__ */ expectMap(function (v) {
    if (v.value instanceof PureScript_CST_Types.TokUpperName) {
        return new Data_Maybe.Just({
            token: v,
            module: v.value.value0,
            name: v.value.value1
        });
    };
    return Data_Maybe.Nothing.value;
});
var parseQualifiedSymbol = /* #__PURE__ */ expectMap(function (v) {
    if (v.value instanceof PureScript_CST_Types.TokSymbolName) {
        return new Data_Maybe.Just({
            token: v,
            module: v.value.value0,
            name: v.value.value1
        });
    };
    return Data_Maybe.Nothing.value;
});
var parseString = /* #__PURE__ */ expectMap(function (v) {
    if (v.value instanceof PureScript_CST_Types.TokString) {
        return new Data_Maybe.Just(new Data_Tuple.Tuple(v, v.value.value1));
    };
    if (v.value instanceof PureScript_CST_Types.TokRawString) {
        return new Data_Maybe.Just(new Data_Tuple.Tuple(v, v.value.value0));
    };
    return Data_Maybe.Nothing.value;
});
var parseSymbol = /* #__PURE__ */ expectMap(function (v) {
    if (v.value instanceof PureScript_CST_Types.TokSymbolName && v.value.value0 instanceof Data_Maybe.Nothing) {
        return new Data_Maybe.Just({
            token: v,
            name: v.value.value1
        });
    };
    return Data_Maybe.Nothing.value;
});
var expect = function (pred) {
    return expectMap(function (tok) {
        var $247 = pred(tok.value);
        if ($247) {
            return new Data_Maybe.Just(tok);
        };
        return Data_Maybe.Nothing.value;
    });
};
var tokAt = /* #__PURE__ */ expect(function (v) {
    if (v instanceof PureScript_CST_Types.TokAt) {
        return true;
    };
    return false;
});
var tokBackslash = /* #__PURE__ */ expect(function (v) {
    if (v instanceof PureScript_CST_Types.TokBackslash) {
        return true;
    };
    return false;
});
var tokComma = /* #__PURE__ */ expect(function (v) {
    if (v instanceof PureScript_CST_Types.TokComma) {
        return true;
    };
    return false;
});
var tokDot = /* #__PURE__ */ expect(function (v) {
    if (v instanceof PureScript_CST_Types.TokDot) {
        return true;
    };
    return false;
});
var parseRecordAccessor = function (expr) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokDot)(function (dot) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(separated(tokDot)(parseLabel))(function (path) {
            return pure1(new PureScript_CST_Types.ExprRecordAccessor({
                expr: expr,
                dot: dot,
                path: path
            }));
        });
    });
};
var tokDoubleColon = /* #__PURE__ */ expect(function (v) {
    if (v instanceof PureScript_CST_Types.TokDoubleColon) {
        return true;
    };
    return false;
});
var parseInstanceName = /* #__PURE__ */ (function () {
    return Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.Tuple.create)(parseIdent))(tokDoubleColon);
})();
var tokEquals = /* #__PURE__ */ expect(function (v) {
    if (v instanceof PureScript_CST_Types.TokEquals) {
        return true;
    };
    return false;
});
var tokForall = /* #__PURE__ */ expect(function (v) {
    if (v instanceof PureScript_CST_Types.TokForall) {
        return true;
    };
    return false;
});
var tokKeyOperator = function (sym) {
    return expect(function (v) {
        if (v instanceof PureScript_CST_Types.TokOperator && v.value0 instanceof Data_Maybe.Nothing) {
            return sym === v.value1;
        };
        return false;
    });
};
var parseBinderNegative = /* #__PURE__ */ Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(/* #__PURE__ */ tokKeyOperator("-"))(function (negative) {
    return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.uncurry(PureScript_CST_Types.BinderInt.create(new Data_Maybe.Just(negative))))($$parseInt))(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.uncurry(PureScript_CST_Types.BinderNumber.create(new Data_Maybe.Just(negative))))(parseNumber));
});
var parseRecordLabeled = function (valueParser) {
    var parseRecordField = Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.uncurry(PureScript_CST_Types.RecordField.create))(PureScript_CST_Parser_Monad["try"](Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.Tuple.create)(parseLabel))(tokKeyOperator(":")))))(valueParser);
    return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseRecordField)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.RecordPun.create)(parseIdent));
};
var parseTypeNegative = /* #__PURE__ */ Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(/* #__PURE__ */ tokKeyOperator("-"))(function (negative) {
    return Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.uncurry(PureScript_CST_Types.TypeInt.create(new Data_Maybe.Just(negative))))($$parseInt);
});
var tokKeySymbol = function (sym) {
    return expect(function (v) {
        if (v instanceof PureScript_CST_Types.TokSymbolName && v.value0 instanceof Data_Maybe.Nothing) {
            return sym === v.value1;
        };
        return false;
    });
};
var tokKeyword = function (kw) {
    return expect(function (v) {
        if (v instanceof PureScript_CST_Types.TokLowerName && v.value0 instanceof Data_Maybe.Nothing) {
            return kw === v.value1;
        };
        return false;
    });
};
var parseFixityKeyword = /* #__PURE__ */ (function () {
    return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Function.flip(Data_Tuple.Tuple.create)(PureScript_CST_Types.Infix.value))(tokKeyword("infix")))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Function.flip(Data_Tuple.Tuple.create)(PureScript_CST_Types.Infixl.value))(tokKeyword("infixl")))(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Function.flip(Data_Tuple.Tuple.create)(PureScript_CST_Types.Infixr.value))(tokKeyword("infixr"))));
})();
var parseFixityOp = /* #__PURE__ */ (function () {
    return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.FixityType.create)(tokKeyword("type")))(parseQualifiedProper))(tokKeyword("as")))(parseOperator))(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.FixityValue.create)(parseQualifiedIdentOrProper))(tokKeyword("as")))(parseOperator));
})();
var parseDeclFixity = /* #__PURE__ */ Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseFixityKeyword)(function (keyword) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseSmallInt)(function (prec) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseFixityOp)(function (operator) {
            return pure2(new PureScript_CST_Types.DeclFixity({
                keyword: keyword,
                prec: prec,
                operator: operator
            }));
        });
    });
});
var parseForeignKind = /* #__PURE__ */ (function () {
    return PureScript_CST_Parser_Monad["try"](Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.ForeignKind.create)(tokKeyword("kind")))(parseProper));
})();
var parseRole = /* #__PURE__ */ (function () {
    return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Function.flip(Data_Tuple.Tuple.create)(PureScript_CST_Types.Representational.value))(tokKeyword("representational")))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Function.flip(Data_Tuple.Tuple.create)(PureScript_CST_Types.Nominal.value))(tokKeyword("nominal")))(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Function.flip(Data_Tuple.Tuple.create)(PureScript_CST_Types.Phantom.value))(tokKeyword("phantom"))));
})();
var parseDeclRole = function (keyword1) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokKeyword("role"))(function (keyword2) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseProper)(function (name) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(many1(parseRole))(function (roles) {
                return pure2(new PureScript_CST_Types.DeclRole(keyword1, keyword2, name, roles));
            });
        });
    });
};
var tokLayoutEnd = /* #__PURE__ */ expect(function (v) {
    if (v instanceof PureScript_CST_Types.TokLayoutEnd) {
        return true;
    };
    return false;
});
var tokLayoutSep = /* #__PURE__ */ expect(function (v) {
    if (v instanceof PureScript_CST_Types.TokLayoutSep) {
        return true;
    };
    return false;
});
var parseInstanceChainSeparator = /* #__PURE__ */ Control_Apply.applyFirst(PureScript_CST_Parser_Monad.applyParser)(/* #__PURE__ */ tokKeyword("else"))(/* #__PURE__ */ PureScript_CST_Parser_Monad.optional(tokLayoutSep));
var tokLayoutStart = /* #__PURE__ */ expect(function (v) {
    if (v instanceof PureScript_CST_Types.TokLayoutStart) {
        return true;
    };
    return false;
});
var layoutNonEmpty = function (valueParser) {
    return Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(function (v) {
        return function (v1) {
            return Data_Array_NonEmpty["cons$prime"](v)(v1);
        };
    })(Control_Apply.applySecond(PureScript_CST_Parser_Monad.applyParser)(tokLayoutStart)(valueParser)))(Control_Apply.applyFirst(PureScript_CST_Parser_Monad.applyParser)(PureScript_CST_Parser_Monad.many(Control_Apply.applySecond(PureScript_CST_Parser_Monad.applyParser)(tokLayoutSep)(valueParser)))(tokLayoutEnd));
};
var tokLeftArrow = /* #__PURE__ */ expect(function (v) {
    if (v instanceof PureScript_CST_Types.TokLeftArrow) {
        return true;
    };
    return false;
});
var tokLeftBrace = /* #__PURE__ */ expect(function (v) {
    if (v instanceof PureScript_CST_Types.TokLeftBrace) {
        return true;
    };
    return false;
});
var tokLeftFatArrow = /* #__PURE__ */ expect(function (v) {
    if (v instanceof PureScript_CST_Types.TokOperator && v.value0 instanceof Data_Maybe.Nothing) {
        return v.value1 === "<=" || v.value1 === "\u21d0";
    };
    return false;
});
var tokLeftParen = /* #__PURE__ */ expect(function (v) {
    if (v instanceof PureScript_CST_Types.TokLeftParen) {
        return true;
    };
    return false;
});
var tokLeftSquare = /* #__PURE__ */ expect(function (v) {
    if (v instanceof PureScript_CST_Types.TokLeftSquare) {
        return true;
    };
    return false;
});
var tokPipe = /* #__PURE__ */ expect(function (v) {
    if (v instanceof PureScript_CST_Types.TokPipe) {
        return true;
    };
    return false;
});
var tokQualifiedKeyword = function (kw) {
    return expect(function (v) {
        if (v instanceof PureScript_CST_Types.TokLowerName) {
            return kw === v.value1;
        };
        return false;
    });
};
var tokRightArrow = /* #__PURE__ */ expect(function (v) {
    if (v instanceof PureScript_CST_Types.TokRightArrow) {
        return true;
    };
    return false;
});
var parseFundep = /* #__PURE__ */ (function () {
    return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.FundepDetermined.create)(tokRightArrow))(many1(parseIdent)))(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.FundepDetermines.create)(many1(parseIdent)))(tokRightArrow))(many1(parseIdent)));
})();
var tokRightBrace = /* #__PURE__ */ expect(function (v) {
    if (v instanceof PureScript_CST_Types.TokRightBrace) {
        return true;
    };
    return false;
});
var tokRightFatArrow = /* #__PURE__ */ expect(function (v) {
    if (v instanceof PureScript_CST_Types.TokRightFatArrow) {
        return true;
    };
    return false;
});
var tokRightParen = /* #__PURE__ */ expect(function (v) {
    if (v instanceof PureScript_CST_Types.TokRightParen) {
        return true;
    };
    return false;
});
var parens = /* #__PURE__ */ wrapped(tokLeftParen)(tokRightParen);
var parseEmptyRow = function (open) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokRightParen)(function (close) {
        return pure3(new PureScript_CST_Types.TypeRow({
            open: open,
            value: {
                labels: Data_Maybe.Nothing.value,
                tail: Data_Maybe.Nothing.value
            },
            close: close
        }));
    });
};
var tokRightSquare = /* #__PURE__ */ expect(function (v) {
    if (v instanceof PureScript_CST_Types.TokRightSquare) {
        return true;
    };
    return false;
});
var tokSymbolArrow = /* #__PURE__ */ expect(function (v) {
    if (v instanceof PureScript_CST_Types.TokSymbolArrow) {
        return true;
    };
    return false;
});
var tokTick = /* #__PURE__ */ expect(function (v) {
    if (v instanceof PureScript_CST_Types.TokTick) {
        return true;
    };
    return false;
});
var tokUnderscore = /* #__PURE__ */ expect(function (v) {
    if (v instanceof PureScript_CST_Types.TokUnderscore) {
        return true;
    };
    return false;
});
var delimited = function (openTok) {
    return function (closeTok) {
        return function (sepTok) {
            return function (valueParser) {
                var parseNonEmpty = function (open) {
                    return Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(function (v) {
                        return function (v1) {
                            return {
                                open: open,
                                value: new Data_Maybe.Just(v),
                                close: v1
                            };
                        };
                    })(separated(sepTok)(valueParser)))(closeTok);
                };
                var parseEmpty = function (open) {
                    return Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(function (v) {
                        return {
                            open: open,
                            value: Data_Maybe.Nothing.value,
                            close: v
                        };
                    })(closeTok);
                };
                return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(openTok)(function (open) {
                    return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseEmpty(open))(parseNonEmpty(open));
                });
            };
        };
    };
};
var parseDataMembers = /* #__PURE__ */ (function () {
    return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.DataAll.create)(tokKeySymbol("..")))(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.DataEnumerated.create)(delimited(tokLeftParen)(tokRightParen)(tokComma)(parseProper)));
})();
var parseExport = /* #__PURE__ */ (function () {
    return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.ExportTypeOp.create)(tokKeyword("type")))(parseSymbol))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.ExportClass.create)(tokKeyword("class")))(parseProper))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.ExportModule.create)(tokKeyword("module")))(parseModuleName))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.ExportOp.create)(parseSymbol))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.ExportValue.create)(parseIdent))(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.ExportType.create)(parseProper))(PureScript_CST_Parser_Monad.optional(parseDataMembers)))))));
})();
var parseImport = /* #__PURE__ */ (function () {
    return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.ImportOp.create)(parseSymbol))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.ImportType.create)(parseProper))(PureScript_CST_Parser_Monad.optional(parseDataMembers)))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.ImportTypeOp.create)(tokKeyword("type")))(parseSymbol))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.ImportClass.create)(tokKeyword("class")))(parseProper))(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.ImportValue.create)(parseIdent)))));
})();
var parseImportDecl = /* #__PURE__ */ Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(/* #__PURE__ */ tokKeyword("import"))(function (keyword) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseModuleName)(function (module_) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.optional(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.Tuple.create)(PureScript_CST_Parser_Monad.optional(tokKeyword("hiding"))))(parens(separated(tokComma)(parseImport)))))(function (names) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.optional(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.Tuple.create)(tokKeyword("as")))(parseModuleName)))(function (qualified) {
                return pure4({
                    keyword: keyword,
                    module: module_,
                    names: names,
                    qualified: qualified
                });
            });
        });
    });
});
var parseModuleImportDecls = /* #__PURE__ */ PureScript_CST_Parser_Monad.many(/* #__PURE__ */ Control_Apply.applyFirst(PureScript_CST_Parser_Monad.applyParser)(parseImportDecl)(/* #__PURE__ */ Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(tokLayoutSep)(/* #__PURE__ */ PureScript_CST_Parser_Monad.lookAhead(tokLayoutEnd))));
var parseModuleHeader = /* #__PURE__ */ Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(/* #__PURE__ */ tokKeyword("module"))(function (keyword) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseModuleName)(function (name) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.optional(parens(separated(tokComma)(parseExport))))(function (exports) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokKeyword("where"))(function (where_) {
                return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(Control_Apply.applySecond(PureScript_CST_Parser_Monad.applyParser)(tokLayoutStart)(parseModuleImportDecls))(function (imports) {
                    return pure5({
                        keyword: keyword,
                        name: name,
                        exports: exports,
                        where: where_,
                        imports: imports
                    });
                });
            });
        });
    });
});
var braces = /* #__PURE__ */ wrapped(tokLeftBrace)(tokRightBrace);
var parseTypeVarKinded = function (parseBindingName) {
    return Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.TypeVarKinded.create)(parens(Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseBindingName)(function (label) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokDoubleColon)(function (separator) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)($lazy_parseType(550))(function (value) {
                return pure6({
                    label: label,
                    separator: separator,
                    value: value
                });
            });
        });
    })));
};
var parseTypeVarBinding = function (parseBindingName) {
    return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseTypeVarKinded(parseBindingName))(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.TypeVarName.create)(parseBindingName));
};
var parseTypeParen = function (open) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)($lazy_parseType(507))(function (value) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokRightParen)(function (close) {
            return pure3(new PureScript_CST_Types.TypeParens({
                open: open,
                value: value,
                close: close
            }));
        });
    });
};
var parseRowTailParen = function (open) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.Tuple.create)(tokPipe))($lazy_parseType(477)))(function (tail) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokRightParen)(function (close) {
            return pure3(new PureScript_CST_Types.TypeRow({
                open: open,
                value: {
                    labels: Data_Maybe.Nothing.value,
                    tail: new Data_Maybe.Just(tail)
                },
                close: close
            }));
        });
    });
};
var parseRowParen = function (open) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad["try"](Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.Tuple.create)(parseLabel))(tokDoubleColon)))(function (v) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)($lazy_parseType(459))(function (value) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.many(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.Tuple.create)(tokComma))($lazy_parseRowLabel(460))))(function (rest) {
                return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.optional(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.Tuple.create)(tokPipe))($lazy_parseType(461))))(function (tail) {
                    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokRightParen)(function (close) {
                        return pure3(new PureScript_CST_Types.TypeRow({
                            open: open,
                            value: {
                                labels: new Data_Maybe.Just({
                                    head: {
                                        label: v.value0,
                                        separator: v.value1,
                                        value: value
                                    },
                                    tail: rest
                                }),
                                tail: tail
                            },
                            close: close
                        }));
                    });
                });
            });
        });
    });
};
var parseKindedVar = function (open) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad["try"](Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.Tuple.create)(parens(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.TypeVar.create)(parseIdent))))(tokDoubleColon)))(function (v) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)($lazy_parseType(497))(function (kind) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokRightParen)(function (close) {
                return pure3(new PureScript_CST_Types.TypeParens({
                    open: open,
                    value: new PureScript_CST_Types.TypeKinded(new PureScript_CST_Types.TypeParens(v.value0), v.value1, kind),
                    close: close
                }));
            });
        });
    });
};
var $lazy_parseForall = /* #__PURE__ */ $runtime_lazy("parseForall", "PureScript.CST.Parser", function () {
    return Control_Lazy.defer(PureScript_CST_Parser_Monad.lazyParser)(function (v) {
        return Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.TypeForall.create)(tokForall))(many1($lazy_parseTypeVarBindingWithVisibility(528))))(tokDot))($lazy_parseType1(530));
    });
});
var $lazy_parseRow = /* #__PURE__ */ $runtime_lazy("parseRow", "PureScript.CST.Parser", function () {
    return Control_Lazy.defer(PureScript_CST_Parser_Monad.lazyParser)(function (v) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.optional(separated(tokComma)($lazy_parseRowLabel(513))))(function (labels) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.optional(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.Tuple.create)(tokPipe))($lazy_parseType(514))))(function (tail) {
                return pure8({
                    labels: labels,
                    tail: tail
                });
            });
        });
    });
});
var $lazy_parseRowLabel = /* #__PURE__ */ $runtime_lazy("parseRowLabel", "PureScript.CST.Parser", function () {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseLabel)(function (label) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokDoubleColon)(function (separator) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)($lazy_parseType(521))(function (value) {
                return pure7({
                    label: label,
                    separator: separator,
                    value: value
                });
            });
        });
    });
});
var $lazy_parseType = /* #__PURE__ */ $runtime_lazy("parseType", "PureScript.CST.Parser", function () {
    return Control_Lazy.defer(PureScript_CST_Parser_Monad.lazyParser)(function (v) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)($lazy_parseType1(393))(function (ty) {
            return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.TypeKinded.create(ty))(tokDoubleColon))($lazy_parseType(394)))(Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser)(ty));
        });
    });
});
var $lazy_parseType1 = /* #__PURE__ */ $runtime_lazy("parseType1", "PureScript.CST.Parser", function () {
    return Control_Lazy.defer(PureScript_CST_Parser_Monad.lazyParser)(function (v) {
        return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)($lazy_parseForall(399))($lazy_parseType2(400));
    });
});
var $lazy_parseType2 = /* #__PURE__ */ $runtime_lazy("parseType2", "PureScript.CST.Parser", function () {
    return Control_Lazy.defer(PureScript_CST_Parser_Monad.lazyParser)(function (v) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)($lazy_parseType3(404))(function (ty) {
            return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.TypeArrow.create(ty))(tokRightArrow))($lazy_parseType1(405)))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.TypeConstrained.create(ty))(tokRightFatArrow))($lazy_parseType1(406)))(Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser)(ty)));
        });
    });
});
var $lazy_parseType3 = /* #__PURE__ */ $runtime_lazy("parseType3", "PureScript.CST.Parser", function () {
    return Control_Lazy.defer(PureScript_CST_Parser_Monad.lazyParser)(function (v) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)($lazy_parseType4(411))(function (ty) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.many(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.Tuple.create)(parseQualifiedOperator))($lazy_parseType4(412))))(function (ops) {
                return Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser)((function () {
                    var v1 = Data_Array_NonEmpty.fromArray(ops);
                    if (v1 instanceof Data_Maybe.Nothing) {
                        return ty;
                    };
                    if (v1 instanceof Data_Maybe.Just) {
                        return new PureScript_CST_Types.TypeOp(ty, v1.value0);
                    };
                    throw new Error("Failed pattern match at PureScript.CST.Parser (line 413, column 8 - line 415, column 28): " + [ v1.constructor.name ]);
                })());
            });
        });
    });
});
var $lazy_parseType4 = /* #__PURE__ */ $runtime_lazy("parseType4", "PureScript.CST.Parser", function () {
    return Control_Lazy.defer(PureScript_CST_Parser_Monad.lazyParser)(function (v) {
        return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseTypeNegative)($lazy_parseType5(419));
    });
});
var $lazy_parseType5 = /* #__PURE__ */ $runtime_lazy("parseType5", "PureScript.CST.Parser", function () {
    return Control_Lazy.defer(PureScript_CST_Parser_Monad.lazyParser)(function (v) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)($lazy_parseTypeAtom(423))(function (ty) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.many($lazy_parseTypeAtom(424)))(function (args) {
                return Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser)((function () {
                    var v1 = Data_Array_NonEmpty.fromArray(args);
                    if (v1 instanceof Data_Maybe.Nothing) {
                        return ty;
                    };
                    if (v1 instanceof Data_Maybe.Just) {
                        return new PureScript_CST_Types.TypeApp(ty, v1.value0);
                    };
                    throw new Error("Failed pattern match at PureScript.CST.Parser (line 425, column 8 - line 427, column 29): " + [ v1.constructor.name ]);
                })());
            });
        });
    });
});
var $lazy_parseTypeAtom = /* #__PURE__ */ $runtime_lazy("parseTypeAtom", "PureScript.CST.Parser", function () {
    return Control_Lazy.defer(PureScript_CST_Parser_Monad.lazyParser)(function (v) {
        return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.TypeVar.create)(parseIdent))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.TypeConstructor.create)(parseQualifiedProper))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.uncurry(PureScript_CST_Types.TypeString.create))(parseString))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.uncurry(PureScript_CST_Types.TypeInt.create(Data_Maybe.Nothing.value)))($$parseInt))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)($lazy_parseTypeParens(435))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.TypeRecord.create)(braces($lazy_parseRow(436))))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.TypeOpName.create)(parseQualifiedSymbol))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.TypeHole.create)(parseHole))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.TypeWildcard.create)(tokUnderscore))(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.TypeArrowName.create)(tokSymbolArrow))))))))));
    });
});
var $lazy_parseTypeParens = /* #__PURE__ */ $runtime_lazy("parseTypeParens", "PureScript.CST.Parser", function () {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokLeftParen)(function (open) {
        return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseRowParen(open))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseRowTailParen(open))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseKindedVar(open))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseTypeParen(open))(parseEmptyRow(open)))));
    });
});
var $lazy_parseTypeVarBindingWithVisibility = /* #__PURE__ */ $runtime_lazy("parseTypeVarBindingWithVisibility", "PureScript.CST.Parser", function () {
    return Control_Lazy.defer(PureScript_CST_Parser_Monad.lazyParser)(function (v) {
        return parseTypeVarBinding(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(function (v1) {
            return function (v2) {
                return {
                    prefix: v1,
                    value: v2
                };
            };
        })(PureScript_CST_Parser_Monad.optional(tokAt)))(parseIdent));
    });
});
var parseForall = /* #__PURE__ */ $lazy_parseForall(524);
var parseRow = /* #__PURE__ */ $lazy_parseRow(511);
var parseRowLabel = /* #__PURE__ */ $lazy_parseRowLabel(517);
var parseType = /* #__PURE__ */ $lazy_parseType(391);
var parseType1 = /* #__PURE__ */ $lazy_parseType1(397);
var parseType2 = /* #__PURE__ */ $lazy_parseType2(402);
var parseType3 = /* #__PURE__ */ $lazy_parseType3(409);
var parseType4 = /* #__PURE__ */ $lazy_parseType4(417);
var parseType5 = /* #__PURE__ */ $lazy_parseType5(421);
var parseTypeAtom = /* #__PURE__ */ $lazy_parseTypeAtom(429);
var parseTypeParens = /* #__PURE__ */ $lazy_parseTypeParens(442);
var parseTypeVarBindingWithVisibility = /* #__PURE__ */ $lazy_parseTypeVarBindingWithVisibility(532);
var parseDataCtor = /* #__PURE__ */ Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(/* #__PURE__ */ Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(function (v) {
    return function (v1) {
        return {
            name: v,
            fields: v1
        };
    };
})(parseProper))(/* #__PURE__ */ PureScript_CST_Parser_Monad.many(parseTypeAtom));
var $lazy_parseBinder = /* #__PURE__ */ $runtime_lazy("parseBinder", "PureScript.CST.Parser", function () {
    return Control_Lazy.defer(PureScript_CST_Parser_Monad.lazyParser)(function (v) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)($lazy_parseBinder1(825))(function (binder) {
            return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.BinderTyped.create(binder))(tokDoubleColon))(parseType))(Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser)(binder));
        });
    });
});
var $lazy_parseBinder1 = /* #__PURE__ */ $runtime_lazy("parseBinder1", "PureScript.CST.Parser", function () {
    return Control_Lazy.defer(PureScript_CST_Parser_Monad.lazyParser)(function (v) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)($lazy_parseBinder2(831))(function (binder) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.many(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.Tuple.create)(parseQualifiedOperator))($lazy_parseBinder2(832))))(function (ops) {
                return Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser)((function () {
                    var v1 = Data_Array_NonEmpty.fromArray(ops);
                    if (v1 instanceof Data_Maybe.Nothing) {
                        return binder;
                    };
                    if (v1 instanceof Data_Maybe.Just) {
                        return new PureScript_CST_Types.BinderOp(binder, v1.value0);
                    };
                    throw new Error("Failed pattern match at PureScript.CST.Parser (line 833, column 8 - line 835, column 34): " + [ v1.constructor.name ]);
                })());
            });
        });
    });
});
var $lazy_parseBinder2 = /* #__PURE__ */ $runtime_lazy("parseBinder2", "PureScript.CST.Parser", function () {
    return Control_Lazy.defer(PureScript_CST_Parser_Monad.lazyParser)(function (v) {
        return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseBinderNegative)(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)($lazy_parseBinderConstructor(840))($lazy_parseBinderAtom(841)));
    });
});
var $lazy_parseBinderAtom = /* #__PURE__ */ $runtime_lazy("parseBinderAtom", "PureScript.CST.Parser", function () {
    return Control_Lazy.defer(PureScript_CST_Parser_Monad.lazyParser)(function (v) {
        return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)($lazy_parseIdentBinder(857))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Function.flip(PureScript_CST_Types.BinderConstructor.create)([  ]))(parseQualifiedProper))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.BinderWildcard.create)(tokUnderscore))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.uncurry(PureScript_CST_Types.BinderString.create))(parseString))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.uncurry(PureScript_CST_Types.BinderChar.create))(parseChar))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.uncurry(PureScript_CST_Types.BinderBoolean.create))(parseBoolean))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.uncurry(PureScript_CST_Types.BinderInt.create(Data_Maybe.Nothing.value)))($$parseInt))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.uncurry(PureScript_CST_Types.BinderNumber.create(Data_Maybe.Nothing.value)))(parseNumber))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.BinderArray.create)(delimited(tokLeftSquare)(tokRightSquare)(tokComma)($lazy_parseBinder(865))))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.BinderRecord.create)(delimited(tokLeftBrace)(tokRightBrace)(tokComma)(parseRecordLabeled($lazy_parseBinder(866)))))(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.BinderParens.create)(parens($lazy_parseBinder(867)))))))))))));
    });
});
var $lazy_parseBinderConstructor = /* #__PURE__ */ $runtime_lazy("parseBinderConstructor", "PureScript.CST.Parser", function () {
    return Control_Lazy.defer(PureScript_CST_Parser_Monad.lazyParser)(function (v) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseQualifiedProper)(function (name) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.many($lazy_parseBinderAtom(852)))(function (apps) {
                return pure9(new PureScript_CST_Types.BinderConstructor(name, apps));
            });
        });
    });
});
var $lazy_parseIdentBinder = /* #__PURE__ */ $runtime_lazy("parseIdentBinder", "PureScript.CST.Parser", function () {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseIdent)(function (ident) {
        return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.BinderNamed.create(ident))(tokAt))($lazy_parseBinderAtom(872)))(Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser)(new PureScript_CST_Types.BinderVar(ident)));
    });
});
var parseBinder = /* #__PURE__ */ $lazy_parseBinder(823);
var parseBinder1 = /* #__PURE__ */ $lazy_parseBinder1(829);
var parseBinder2 = /* #__PURE__ */ $lazy_parseBinder2(837);
var parseBinderAtom = /* #__PURE__ */ $lazy_parseBinderAtom(855);
var parseBinderConstructor = /* #__PURE__ */ $lazy_parseBinderConstructor(849);
var parseIdentBinder = /* #__PURE__ */ $lazy_parseIdentBinder(869);
var parseClassConstraints = function (parseOneConstraint) {
    return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.Many.create)(parens(separated(tokComma)(parseType))))(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.One.create)(parseOneConstraint));
};
var parseDeclDerive = /* #__PURE__ */ Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(/* #__PURE__ */ tokKeyword("derive"))(function (derive_) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.optional(tokKeyword("newtype")))(function (newtype_) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokKeyword("instance"))(function (keyword) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.optional(parseInstanceName))(function (name) {
                return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.optional(PureScript_CST_Parser_Monad["try"](Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.Tuple.create)(parseClassConstraints(parseType3)))(tokRightFatArrow))))(function (constraints) {
                    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseQualifiedProper)(function (className) {
                        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.many(parseTypeAtom))(function (types) {
                            return pure2(new PureScript_CST_Types.DeclDerive(derive_, newtype_, {
                                keyword: keyword,
                                name: name,
                                constraints: constraints,
                                className: className,
                                types: types
                            }));
                        });
                    });
                });
            });
        });
    });
});
var parseClassMember = /* #__PURE__ */ Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseIdent)(function (label) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokDoubleColon)(function (separator) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseType)(function (value) {
            return pure7({
                label: label,
                separator: separator,
                value: value
            });
        });
    });
});
var parseDeclClassSignature = function (keyword) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad["try"](Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.Tuple.create)(parseProper))(tokDoubleColon)))(function (v) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseType)(function (value) {
            return pure2(new PureScript_CST_Types.DeclKindSignature(keyword, {
                label: v.value0,
                separator: v.value1,
                value: value
            }));
        });
    });
};
var parseDeclKindSignature = function (keyword) {
    return function (label) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokDoubleColon)(function (separator) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseType)(function (value) {
                return pure2(new PureScript_CST_Types.DeclKindSignature(keyword, {
                    label: label,
                    separator: separator,
                    value: value
                }));
            });
        });
    };
};
var parseDeclSignature = function (label) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokDoubleColon)(function (separator) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseType)(function (value) {
            return pure2(new PureScript_CST_Types.DeclSignature({
                label: label,
                separator: separator,
                value: value
            }));
        });
    });
};
var parseForeignData = /* #__PURE__ */ Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(/* #__PURE__ */ tokKeyword("data"))(function (keyword) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseProper)(function (label) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokDoubleColon)(function (separator) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseType)(function (value) {
                return pure10(new PureScript_CST_Types.ForeignData(keyword, {
                    label: label,
                    separator: separator,
                    value: value
                }));
            });
        });
    });
});
var parseForeignValue = /* #__PURE__ */ Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseIdent)(function (label) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokDoubleColon)(function (separator) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseType)(function (value) {
            return pure10(new PureScript_CST_Types.ForeignValue({
                label: label,
                separator: separator,
                value: value
            }));
        });
    });
});
var parseDeclForeign = /* #__PURE__ */ Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(/* #__PURE__ */ tokKeyword("foreign"))(function (keyword1) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokKeyword("import"))(function (keyword2) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseForeignData)(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseForeignKind)(parseForeignValue)))(function (foreign_) {
            return pure2(new PureScript_CST_Types.DeclForeign(keyword1, keyword2, foreign_));
        });
    });
});
var parseInstanceBindingSignature = function (label) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokDoubleColon)(function (separator) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseType)(function (value) {
            return pure11(new PureScript_CST_Types.InstanceBindingSignature({
                label: label,
                separator: separator,
                value: value
            }));
        });
    });
};
var parseLetBindingSignature = function (label) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokDoubleColon)(function (separator) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseType)(function (value) {
            return pure12(new PureScript_CST_Types.LetBindingSignature({
                label: label,
                separator: separator,
                value: value
            }));
        });
    });
};
var parseRecordUpdates = function (expr) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad["try"](Control_Apply.applyFirst(PureScript_CST_Parser_Monad.applyParser)(tokLeftBrace)(PureScript_CST_Parser_Monad.lookAhead(Control_Apply.applySecond(PureScript_CST_Parser_Monad.applyParser)(parseLabel)(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(tokEquals)(tokLeftBrace))))))(function (open) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(separated(tokComma)($lazy_parseRecordUpdate(709)))(function (value) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokRightBrace)(function (close) {
                return pure1(new PureScript_CST_Types.ExprRecordUpdate(expr, {
                    open: open,
                    value: value,
                    close: close
                }));
            });
        });
    });
};
var parseRecordUpdateLeaf = function (label) {
    return Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.RecordUpdateLeaf.create(label))(tokEquals))($lazy_parseExpr(723));
};
var parseRecordUpdateBranch = function (label) {
    return Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.RecordUpdateBranch.create(label))(braces(separated(tokComma)($lazy_parseRecordUpdate(728))));
};
var parseLetBindingName = function (name) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.many(parseBinderAtom))(function (binders) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseGuarded(tokEquals))(function (guarded) {
            return pure12(new PureScript_CST_Types.LetBindingName({
                name: name,
                binders: binders,
                guarded: guarded
            }));
        });
    });
};
var parseGuarded = function (sepParser) {
    var parsePatternGuard = Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(function (v) {
        return function (v1) {
            return {
                binder: v,
                expr: v1
            };
        };
    })(PureScript_CST_Parser_Monad.optional(PureScript_CST_Parser_Monad["try"](Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.Tuple.create)(parseBinder))(tokLeftArrow)))))($lazy_parseExpr(814));
    var parseGuardedExpr = Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(function (v) {
        return function (v1) {
            return function (v2) {
                return function (v3) {
                    return {
                        bar: v,
                        patterns: v1,
                        separator: v2,
                        where: v3
                    };
                };
            };
        };
    })(tokPipe))(separated(tokComma)(parsePatternGuard)))(sepParser))($lazy_parseWhere(808));
    return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.Unconditional.create)(sepParser))($lazy_parseWhere(800)))(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.Guarded.create)(many1(parseGuardedExpr)));
};
var parseBadSingleCaseWhere = function (binder) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokRightArrow)(function (arrow) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(Control_Apply.applySecond(PureScript_CST_Parser_Monad.applyParser)(tokLayoutEnd)($lazy_parseWhere(665)))(function (body) {
            return pure15(Data_Array_NonEmpty.singleton(new Data_Tuple.Tuple({
                head: binder,
                tail: [  ]
            }, new PureScript_CST_Types.Unconditional(arrow, body))));
        });
    });
};
var parseBadSingleCaseGuarded = function (binder) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(Control_Apply.applySecond(PureScript_CST_Parser_Monad.applyParser)(tokLayoutEnd)(parseGuarded(tokRightArrow)))(function (body) {
        return pure15(Data_Array_NonEmpty.singleton(new Data_Tuple.Tuple({
            head: binder,
            tail: [  ]
        }, body)));
    });
};
var $lazy_parseAdo = /* #__PURE__ */ $runtime_lazy("parseAdo", "PureScript.CST.Parser", function () {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokQualifiedKeyword("ado"))(function (keyword) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)((function () {
            
            // `recoverDoStatement` recovers too much if it is immediately
            // confronted with `TokLayoutEnd`, since that is associated with a
            // `layoutStack` _of the parent_ as opposed to the stuff we actually
            // want to recover, which we would correctly guess if we saw a statement
            // or two inside the block
var valueParser = recoverDoStatement($lazy_parseDoStatement(690));
            var nonEmptyCase = Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Array.cons)(valueParser))(PureScript_CST_Parser_Monad.many(Control_Apply.applySecond(PureScript_CST_Parser_Monad.applyParser)(tokLayoutSep)(valueParser)));
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokLayoutStart)(function () {
                return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.voidRight(PureScript_CST_Parser_Monad.functorParser)([  ])(tokLayoutEnd))(Control_Apply.applyFirst(PureScript_CST_Parser_Monad.applyParser)(nonEmptyCase)(tokLayoutEnd));
            });
        })())(function (statements) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokKeyword("in"))(function (in_) {
                return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)($lazy_parseExpr(697))(function (result) {
                    return pure1(new PureScript_CST_Types.ExprAdo({
                        keyword: keyword,
                        statements: statements,
                        "in": in_,
                        result: result
                    }));
                });
            });
        });
    });
});
var $lazy_parseBadSingleCaseBranch = /* #__PURE__ */ $runtime_lazy("parseBadSingleCaseBranch", "PureScript.CST.Parser", function () {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(Control_Apply.applySecond(PureScript_CST_Parser_Monad.applyParser)(tokLayoutStart)(parseBinder1))(function (binder) {
        return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseBadSingleCaseWhere(binder))(parseBadSingleCaseGuarded(binder));
    });
});
var $lazy_parseCase = /* #__PURE__ */ $runtime_lazy("parseCase", "PureScript.CST.Parser", function () {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokKeyword("case"))(function (keyword) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(separated(tokComma)($lazy_parseExpr(647)))(function (head) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokKeyword("of"))(function (of_) {
                return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(PureScript_CST_Parser_Monad["try"]($lazy_parseBadSingleCaseBranch(649)))($lazy_parseCaseBranches(649)))(function (branches) {
                    return pure1(new PureScript_CST_Types.ExprCase({
                        keyword: keyword,
                        head: head,
                        of: of_,
                        branches: branches
                    }));
                });
            });
        });
    });
});
var $lazy_parseCaseBranches = /* #__PURE__ */ $runtime_lazy("parseCaseBranches", "PureScript.CST.Parser", function () {
    return Control_Lazy.defer(PureScript_CST_Parser_Monad.lazyParser)(function (v) {
        return layoutNonEmpty(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.Tuple.create)(separated(tokComma)(parseBinder1)))(parseGuarded(tokRightArrow)));
    });
});
var $lazy_parseDo = /* #__PURE__ */ $runtime_lazy("parseDo", "PureScript.CST.Parser", function () {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokQualifiedKeyword("do"))(function (keyword) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(layoutNonEmpty(recoverDoStatement($lazy_parseDoStatement(676))))(function (statements) {
            return pure1(new PureScript_CST_Types.ExprDo({
                keyword: keyword,
                statements: statements
            }));
        });
    });
});
var $lazy_parseDoStatement = /* #__PURE__ */ $runtime_lazy("parseDoStatement", "PureScript.CST.Parser", function () {
    return Control_Lazy.defer(PureScript_CST_Parser_Monad.lazyParser)(function (v) {
        return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.DoLet.create)(tokKeyword("let")))(layoutNonEmpty(recoverLetBinding($lazy_parseLetBinding(771)))))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.uncurry(PureScript_CST_Types.DoBind.create))(PureScript_CST_Parser_Monad["try"](Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.Tuple.create)(parseBinder))(tokLeftArrow))))($lazy_parseExpr(772)))(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.DoDiscard.create)($lazy_parseExpr(773))));
    });
});
var $lazy_parseExpr = /* #__PURE__ */ $runtime_lazy("parseExpr", "PureScript.CST.Parser", function () {
    return Control_Lazy.defer(PureScript_CST_Parser_Monad.lazyParser)(function (v) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)($lazy_parseExpr1(555))(function (expr) {
            return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.ExprTyped.create(expr))(tokDoubleColon))(parseType))(Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser)(expr));
        });
    });
});
var $lazy_parseExpr1 = /* #__PURE__ */ $runtime_lazy("parseExpr1", "PureScript.CST.Parser", function () {
    return Control_Lazy.defer(PureScript_CST_Parser_Monad.lazyParser)(function (v) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)($lazy_parseExpr2(561))(function (expr) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.many(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.Tuple.create)(parseQualifiedOperator))($lazy_parseExpr2(562))))(function (ops) {
                return Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser)((function () {
                    var v1 = Data_Array_NonEmpty.fromArray(ops);
                    if (v1 instanceof Data_Maybe.Nothing) {
                        return expr;
                    };
                    if (v1 instanceof Data_Maybe.Just) {
                        return new PureScript_CST_Types.ExprOp(expr, v1.value0);
                    };
                    throw new Error("Failed pattern match at PureScript.CST.Parser (line 563, column 8 - line 565, column 30): " + [ v1.constructor.name ]);
                })());
            });
        });
    });
});
var $lazy_parseExpr2 = /* #__PURE__ */ $runtime_lazy("parseExpr2", "PureScript.CST.Parser", function () {
    return Control_Lazy.defer(PureScript_CST_Parser_Monad.lazyParser)(function (v) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)($lazy_parseExpr3(569))(function (expr) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.many(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.Tuple.create)($lazy_parseTickExpr(570)))($lazy_parseExpr3(570))))(function (ops) {
                return Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser)((function () {
                    var v1 = Data_Array_NonEmpty.fromArray(ops);
                    if (v1 instanceof Data_Maybe.Nothing) {
                        return expr;
                    };
                    if (v1 instanceof Data_Maybe.Just) {
                        return new PureScript_CST_Types.ExprInfix(expr, v1.value0);
                    };
                    throw new Error("Failed pattern match at PureScript.CST.Parser (line 571, column 8 - line 573, column 33): " + [ v1.constructor.name ]);
                })());
            });
        });
    });
});
var $lazy_parseExpr3 = /* #__PURE__ */ $runtime_lazy("parseExpr3", "PureScript.CST.Parser", function () {
    return Control_Lazy.defer(PureScript_CST_Parser_Monad.lazyParser)(function (v) {
        return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.ExprNegate.create)(tokKeyOperator("-")))($lazy_parseExpr3(592)))($lazy_parseExpr4(593));
    });
});
var $lazy_parseExpr4 = /* #__PURE__ */ $runtime_lazy("parseExpr4", "PureScript.CST.Parser", function () {
    return Control_Lazy.defer(PureScript_CST_Parser_Monad.lazyParser)(function (v) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)($lazy_parseExpr5(597))(function (expr) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.many($lazy_parseExprAppSpine(598)))(function (args) {
                return Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser)((function () {
                    var v1 = Data_Array_NonEmpty.fromArray(args);
                    if (v1 instanceof Data_Maybe.Nothing) {
                        return expr;
                    };
                    if (v1 instanceof Data_Maybe.Just) {
                        return new PureScript_CST_Types.ExprApp(expr, v1.value0);
                    };
                    throw new Error("Failed pattern match at PureScript.CST.Parser (line 599, column 8 - line 601, column 31): " + [ v1.constructor.name ]);
                })());
            });
        });
    });
});
var $lazy_parseExpr5 = /* #__PURE__ */ $runtime_lazy("parseExpr5", "PureScript.CST.Parser", function () {
    return Control_Lazy.defer(PureScript_CST_Parser_Monad.lazyParser)(function (v) {
        return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)($lazy_parseIf(605))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)($lazy_parseLetIn(606))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)($lazy_parseLambda(607))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)($lazy_parseCase(608))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)($lazy_parseDo(609))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)($lazy_parseAdo(610))($lazy_parseExpr6(611)))))));
    });
});
var $lazy_parseExpr6 = /* #__PURE__ */ $runtime_lazy("parseExpr6", "PureScript.CST.Parser", function () {
    return Control_Lazy.defer(PureScript_CST_Parser_Monad.lazyParser)(function (v) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)($lazy_parseExpr7(702))(function (expr) {
            return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseRecordUpdates(expr))(Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser)(expr));
        });
    });
});
var $lazy_parseExpr7 = /* #__PURE__ */ $runtime_lazy("parseExpr7", "PureScript.CST.Parser", function () {
    return Control_Lazy.defer(PureScript_CST_Parser_Monad.lazyParser)(function (v) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)($lazy_parseExprAtom(732))(function (expr) {
            return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseRecordAccessor(expr))(Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser)(expr));
        });
    });
});
var $lazy_parseExprAppSpine = /* #__PURE__ */ $runtime_lazy("parseExprAppSpine", "PureScript.CST.Parser", function () {
    return Control_Lazy.defer(PureScript_CST_Parser_Monad.lazyParser)(function (v) {
        return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.AppType.create)(tokAt))(parseTypeAtom))(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.AppTerm.create)($lazy_parseExpr5(616)));
    });
});
var $lazy_parseExprAtom = /* #__PURE__ */ $runtime_lazy("parseExprAtom", "PureScript.CST.Parser", function () {
    return Control_Lazy.defer(PureScript_CST_Parser_Monad.lazyParser)(function (v) {
        return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.ExprIdent.create)(parseQualifiedIdent))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.ExprConstructor.create)(parseQualifiedProper))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.ExprOpName.create)(parseQualifiedSymbol))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.ExprSection.create)(tokUnderscore))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.ExprHole.create)(parseHole))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.uncurry(PureScript_CST_Types.ExprString.create))(parseString))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.uncurry(PureScript_CST_Types.ExprChar.create))(parseChar))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.uncurry(PureScript_CST_Types.ExprBoolean.create))(parseBoolean))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.uncurry(PureScript_CST_Types.ExprInt.create))($$parseInt))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.uncurry(PureScript_CST_Types.ExprNumber.create))(parseNumber))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.ExprArray.create)(delimited(tokLeftSquare)(tokRightSquare)(tokComma)($lazy_parseExpr(754))))(Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.ExprRecord.create)(delimited(tokLeftBrace)(tokRightBrace)(tokComma)(parseRecordLabeled($lazy_parseExpr(755)))))(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.ExprParens.create)(parens($lazy_parseExpr(756)))))))))))))));
    });
});
var $lazy_parseIdentBinding = /* #__PURE__ */ $runtime_lazy("parseIdentBinding", "PureScript.CST.Parser", function () {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseIdent)(function (ident) {
        return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseLetBindingSignature(ident))(parseLetBindingName(ident));
    });
});
var $lazy_parseIf = /* #__PURE__ */ $runtime_lazy("parseIf", "PureScript.CST.Parser", function () {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokKeyword("if"))(function (keyword) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)($lazy_parseExpr(621))(function (cond) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokKeyword("then"))(function (then_) {
                return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)($lazy_parseExpr(623))(function (true_) {
                    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokKeyword("else"))(function (else_) {
                        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)($lazy_parseExpr(625))(function (false_) {
                            return pure1(new PureScript_CST_Types.ExprIf({
                                keyword: keyword,
                                cond: cond,
                                then: then_,
                                "true": true_,
                                "else": else_,
                                "false": false_
                            }));
                        });
                    });
                });
            });
        });
    });
});
var $lazy_parseLambda = /* #__PURE__ */ $runtime_lazy("parseLambda", "PureScript.CST.Parser", function () {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokBackslash)(function (symbol) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(many1(parseBinderAtom))(function (binders) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokRightArrow)(function (arrow) {
                return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)($lazy_parseExpr(641))(function (body) {
                    return pure1(new PureScript_CST_Types.ExprLambda({
                        symbol: symbol,
                        binders: binders,
                        arrow: arrow,
                        body: body
                    }));
                });
            });
        });
    });
});
var $lazy_parseLetBinding = /* #__PURE__ */ $runtime_lazy("parseLetBinding", "PureScript.CST.Parser", function () {
    return Control_Lazy.defer(PureScript_CST_Parser_Monad.lazyParser)(function (v) {
        return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(PureScript_CST_Parser_Monad["try"]($lazy_parseIdentBinding(777)))(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.LetBindingPattern.create)(parseBinder1))(tokEquals))($lazy_parseWhere(778)));
    });
});
var $lazy_parseLetIn = /* #__PURE__ */ $runtime_lazy("parseLetIn", "PureScript.CST.Parser", function () {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokKeyword("let"))(function (keyword) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(layoutNonEmpty(recoverLetBinding($lazy_parseLetBinding(631))))(function (bindings) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokKeyword("in"))(function (in_) {
                return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)($lazy_parseExpr(633))(function (body) {
                    return pure1(new PureScript_CST_Types.ExprLet({
                        keyword: keyword,
                        bindings: bindings,
                        "in": in_,
                        body: body
                    }));
                });
            });
        });
    });
});
var $lazy_parseRecordUpdate = /* #__PURE__ */ $runtime_lazy("parseRecordUpdate", "PureScript.CST.Parser", function () {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseLabel)(function (label) {
        return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseRecordUpdateLeaf(label))(parseRecordUpdateBranch(label));
    });
});
var $lazy_parseTickExpr = /* #__PURE__ */ $runtime_lazy("parseTickExpr", "PureScript.CST.Parser", function () {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokTick)(function (open) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)($lazy_parseTickExpr1(578))(function (value) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokTick)(function (close) {
                return pure14({
                    open: open,
                    value: value,
                    close: close
                });
            });
        });
    });
});
var $lazy_parseTickExpr1 = /* #__PURE__ */ $runtime_lazy("parseTickExpr1", "PureScript.CST.Parser", function () {
    return Control_Lazy.defer(PureScript_CST_Parser_Monad.lazyParser)(function (v) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)($lazy_parseExpr3(584))(function (expr) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.many(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.Tuple.create)(parseQualifiedOperator))($lazy_parseExpr3(585))))(function (ops) {
                return Control_Applicative.pure(PureScript_CST_Parser_Monad.applicativeParser)((function () {
                    var v1 = Data_Array_NonEmpty.fromArray(ops);
                    if (v1 instanceof Data_Maybe.Nothing) {
                        return expr;
                    };
                    if (v1 instanceof Data_Maybe.Just) {
                        return new PureScript_CST_Types.ExprOp(expr, v1.value0);
                    };
                    throw new Error("Failed pattern match at PureScript.CST.Parser (line 586, column 8 - line 588, column 30): " + [ v1.constructor.name ]);
                })());
            });
        });
    });
});
var $lazy_parseWhere = /* #__PURE__ */ $runtime_lazy("parseWhere", "PureScript.CST.Parser", function () {
    return Control_Lazy.defer(PureScript_CST_Parser_Monad.lazyParser)(function (v) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)($lazy_parseExpr(819))(function (expr) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.optional(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.Tuple.create)(tokKeyword("where")))(layoutNonEmpty(recoverLetBinding($lazy_parseLetBinding(820))))))(function (bindings) {
                return pure13({
                    expr: expr,
                    bindings: bindings
                });
            });
        });
    });
});
var parseAdo = /* #__PURE__ */ $lazy_parseAdo(679);
var parseBadSingleCaseBranch = /* #__PURE__ */ $lazy_parseBadSingleCaseBranch(656);
var parseCase = /* #__PURE__ */ $lazy_parseCase(644);
var parseCaseBranches = /* #__PURE__ */ $lazy_parseCaseBranches(652);
var parseDo = /* #__PURE__ */ $lazy_parseDo(673);
var parseDoStatement = /* #__PURE__ */ $lazy_parseDoStatement(769);
var parseExpr = /* #__PURE__ */ $lazy_parseExpr(553);
var parseExpr1 = /* #__PURE__ */ $lazy_parseExpr1(559);
var parseExpr2 = /* #__PURE__ */ $lazy_parseExpr2(567);
var parseExpr3 = /* #__PURE__ */ $lazy_parseExpr3(590);
var parseExpr4 = /* #__PURE__ */ $lazy_parseExpr4(595);
var parseExpr5 = /* #__PURE__ */ $lazy_parseExpr5(603);
var parseExpr6 = /* #__PURE__ */ $lazy_parseExpr6(700);
var parseExpr7 = /* #__PURE__ */ $lazy_parseExpr7(730);
var parseExprAppSpine = /* #__PURE__ */ $lazy_parseExprAppSpine(613);
var parseExprAtom = /* #__PURE__ */ $lazy_parseExprAtom(742);
var parseIdentBinding = /* #__PURE__ */ $lazy_parseIdentBinding(780);
var parseIf = /* #__PURE__ */ $lazy_parseIf(618);
var parseLambda = /* #__PURE__ */ $lazy_parseLambda(636);
var parseLetBinding = /* #__PURE__ */ $lazy_parseLetBinding(775);
var parseLetIn = /* #__PURE__ */ $lazy_parseLetIn(628);
var parseRecordUpdate = /* #__PURE__ */ $lazy_parseRecordUpdate(713);
var parseTickExpr = /* #__PURE__ */ $lazy_parseTickExpr(575);
var parseTickExpr1 = /* #__PURE__ */ $lazy_parseTickExpr1(582);
var parseWhere = /* #__PURE__ */ $lazy_parseWhere(817);
var parseDeclValue1 = function (name) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.many(parseBinderAtom))(function (binders) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseGuarded(tokEquals))(function (guarded) {
            return pure2(new PureScript_CST_Types.DeclValue({
                name: name,
                binders: binders,
                guarded: guarded
            }));
        });
    });
};
var parseDeclValue = /* #__PURE__ */ Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseIdent)(function (ident) {
    return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseDeclSignature(ident))(parseDeclValue1(ident));
});
var parseInstanceBindingName = function (name) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.many(parseBinderAtom))(function (binders) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseGuarded(tokEquals))(function (guarded) {
            return pure11(new PureScript_CST_Types.InstanceBindingName({
                name: name,
                binders: binders,
                guarded: guarded
            }));
        });
    });
};
var parseInstanceBinding = /* #__PURE__ */ Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseIdent)(function (ident) {
    return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseInstanceBindingSignature(ident))(parseInstanceBindingName(ident));
});
var parseInstance = /* #__PURE__ */ Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(/* #__PURE__ */ tokKeyword("instance"))(function (keyword) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.optional(parseInstanceName))(function (name) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.optional(PureScript_CST_Parser_Monad["try"](Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.Tuple.create)(parseClassConstraints(parseType3)))(tokRightFatArrow))))(function (constraints) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseQualifiedProper)(function (className) {
                return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.many(parseTypeAtom))(function (types) {
                    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.optional(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.Tuple.create)(tokKeyword("where")))(layoutNonEmpty(parseInstanceBinding))))(function (body) {
                        return pure16({
                            head: {
                                keyword: keyword,
                                name: name,
                                constraints: constraints,
                                className: className,
                                types: types
                            },
                            body: body
                        });
                    });
                });
            });
        });
    });
});
var parseDeclInstanceChain = /* #__PURE__ */ (function () {
    return Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(PureScript_CST_Types.DeclInstanceChain.create)(separated(parseInstanceChainSeparator)(parseInstance));
})();
var parseTypeVarBindingPlain = /* #__PURE__ */ parseTypeVarBinding(parseIdent);
var parseDeclClass1 = function (keyword) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.optional(PureScript_CST_Parser_Monad["try"](Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.Tuple.create)(parseClassConstraints(parseType5)))(tokLeftFatArrow))))(function ($$super) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseProper)(function (name) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.many(parseTypeVarBindingPlain))(function (vars) {
                return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.optional(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.Tuple.create)(tokPipe))(separated(tokComma)(parseFundep))))(function (fundeps) {
                    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.optional(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.Tuple.create)(tokKeyword("where")))(layoutNonEmpty(parseClassMember))))(function (members) {
                        return pure2(new PureScript_CST_Types.DeclClass({
                            keyword: keyword,
                            "super": $$super,
                            name: name,
                            vars: vars,
                            fundeps: fundeps
                        }, members));
                    });
                });
            });
        });
    });
};
var parseDeclClass = /* #__PURE__ */ Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(/* #__PURE__ */ tokKeyword("class"))(function (keyword) {
    return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseDeclClassSignature(keyword))(parseDeclClass1(keyword));
});
var parseDeclData1 = function (keyword) {
    return function (name) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.many(parseTypeVarBindingPlain))(function (vars) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.optional(Control_Apply.apply(PureScript_CST_Parser_Monad.applyParser)(Data_Functor.map(PureScript_CST_Parser_Monad.functorParser)(Data_Tuple.Tuple.create)(tokEquals))(separated(tokPipe)(parseDataCtor))))(function (ctors) {
                return pure2(new PureScript_CST_Types.DeclData({
                    keyword: keyword,
                    name: name,
                    vars: vars
                }, ctors));
            });
        });
    };
};
var parseDeclData = /* #__PURE__ */ Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(/* #__PURE__ */ tokKeyword("data"))(function (keyword) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseProper)(function (name) {
        return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseDeclKindSignature(keyword)(name))(parseDeclData1(keyword)(name));
    });
});
var parseDeclNewtype1 = function (keyword) {
    return function (name) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.many(parseTypeVarBindingPlain))(function (vars) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokEquals)(function (tok) {
                return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseProper)(function (wrapper) {
                    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseTypeAtom)(function (body) {
                        return pure2(new PureScript_CST_Types.DeclNewtype({
                            keyword: keyword,
                            name: name,
                            vars: vars
                        }, tok, wrapper, body));
                    });
                });
            });
        });
    };
};
var parseDeclNewtype = /* #__PURE__ */ Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(/* #__PURE__ */ tokKeyword("newtype"))(function (keyword) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseProper)(function (name) {
        return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseDeclKindSignature(keyword)(name))(parseDeclNewtype1(keyword)(name));
    });
});
var parseDeclType2 = function (keyword) {
    return function (name) {
        return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.many(parseTypeVarBindingPlain))(function (vars) {
            return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(tokEquals)(function (tok) {
                return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseType)(function (body) {
                    return pure2(new PureScript_CST_Types.DeclType({
                        keyword: keyword,
                        name: name,
                        vars: vars
                    }, tok, body));
                });
            });
        });
    };
};
var parseDeclType1 = function (keyword) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseProper)(function (name) {
        return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseDeclKindSignature(keyword)(name))(parseDeclType2(keyword)(name));
    });
};
var parseDeclType = /* #__PURE__ */ Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(/* #__PURE__ */ tokKeyword("type"))(function (keyword) {
    return Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseDeclRole(keyword))(parseDeclType1(keyword));
});
var parseDecl = /* #__PURE__ */ Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseDeclData)(/* #__PURE__ */ Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseDeclNewtype)(/* #__PURE__ */ Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseDeclType)(/* #__PURE__ */ Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseDeclClass)(/* #__PURE__ */ Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseDeclInstanceChain)(/* #__PURE__ */ Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseDeclDerive)(/* #__PURE__ */ Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseDeclValue)(/* #__PURE__ */ Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(parseDeclForeign)(parseDeclFixity))))))));
var parseModuleDecls = /* #__PURE__ */ PureScript_CST_Parser_Monad.many(/* #__PURE__ */ Control_Apply.applyFirst(PureScript_CST_Parser_Monad.applyParser)(/* #__PURE__ */ recoverDecl(parseDecl))(/* #__PURE__ */ Control_Alt.alt(PureScript_CST_Parser_Monad.altParser)(tokLayoutSep)(/* #__PURE__ */ PureScript_CST_Parser_Monad.lookAhead(tokLayoutEnd))));
var parseModuleBody = /* #__PURE__ */ Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(/* #__PURE__ */ Control_Apply.applyFirst(PureScript_CST_Parser_Monad.applyParser)(parseModuleDecls)(tokLayoutEnd))(function (decls) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(PureScript_CST_Parser_Monad.eof)(function (v) {
        return pure17({
            decls: decls,
            trailingComments: v.value1,
            end: v.value0
        });
    });
});
var parseModule = /* #__PURE__ */ Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseModuleHeader)(function (header) {
    return Control_Bind.bind(PureScript_CST_Parser_Monad.bindParser)(parseModuleBody)(function (body) {
        return pure18({
            header: header,
            body: body
        });
    });
});
export {
    parseModule,
    parseModuleHeader,
    parseModuleBody,
    parseImportDecl,
    parseDecl,
    parseType,
    parseExpr,
    parseBinder
};
