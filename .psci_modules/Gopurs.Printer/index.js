import * as $foreign from "./foreign.js";
import * as Data_Array from "../Data.Array/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Data_String_CodeUnits from "../Data.String.CodeUnits/index.js";
import * as Data_String_Common from "../Data.String.Common/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Gopurs_GoAst from "../Gopurs.GoAst/index.js";
var escapeGoString = $foreign.escapeGoStringImpl;
var printGoExpr = function (expr) {
    if (expr instanceof Gopurs_GoAst.GoVar) {
        return expr.value0;
    };
    if (expr instanceof Gopurs_GoAst.GoString) {
        return "\"" + (escapeGoString(expr.value0) + "\"");
    };
    if (expr instanceof Gopurs_GoAst.GoInt) {
        return Data_Show.show(Data_Show.showInt)(expr.value0);
    };
    if (expr instanceof Gopurs_GoAst.GoCall) {
        return printGoExpr(expr.value0) + ("(" + (Data_String_Common.joinWith(", ")(Data_Functor.map(Data_Functor.functorArray)(printGoExpr)(expr.value1)) + ")"));
    };
    if (expr instanceof Gopurs_GoAst.GoSelector) {
        return printGoExpr(expr.value0) + ("." + expr.value1);
    };
    if (expr instanceof Gopurs_GoAst.GoFunc) {
        var flattenGoFunc = function ($copy_v) {
            return function ($copy_v1) {
                var $tco_var_v = $copy_v;
                var $tco_done = false;
                var $tco_result;
                function $tco_loop(v, v1) {
                    if (v instanceof Gopurs_GoAst.GoFunc) {
                        $tco_var_v = v.value3;
                        $copy_v1 = Data_Array.snoc(v1)(new Data_Tuple.Tuple(v.value0, v.value1));
                        return;
                    };
                    $tco_done = true;
                    return new Data_Tuple.Tuple(v1, v);
                };
                while (!$tco_done) {
                    $tco_result = $tco_loop($tco_var_v, $copy_v1);
                };
                return $tco_result;
            };
        };
        var v = flattenGoFunc(expr.value3)([ new Data_Tuple.Tuple(expr.value0, expr.value1) ]);
        var len = Data_Array.length(v.value0);
        var funcName = (function () {
            var $40 = len === 1;
            if ($40) {
                return "Func";
            };
            return "Func" + Data_Show.show(Data_Show.showInt)(len);
        })();
        var argStr = Data_String_Common.joinWith(", ")(Data_Functor.map(Data_Functor.functorArray)(function (v1) {
            return v1.value0 + (" " + Gopurs_GoAst.goTypeToStr(v1.value1));
        })(v.value0));
        var $44 = len > 10;
        if ($44) {
            if (expr.value3 instanceof Gopurs_GoAst.GoBlock) {
                return "gopurs_runtime.Func(func(" + (expr.value0 + (" " + (Gopurs_GoAst.goTypeToStr(expr.value1) + (") " + (Gopurs_GoAst.goTypeToStr(expr.value2) + (" {\x0a" + (printGoExpr(expr.value3) + "\x0a})")))))));
            };
            return "gopurs_runtime.Func(func(" + (expr.value0 + (" " + (Gopurs_GoAst.goTypeToStr(expr.value1) + (") " + (Gopurs_GoAst.goTypeToStr(expr.value2) + (" {\x0areturn " + (printGoExpr(expr.value3) + "\x0a})")))))));
        };
        if (v.value1 instanceof Gopurs_GoAst.GoBlock) {
            return "gopurs_runtime." + (funcName + ("(func(" + (argStr + (") " + (Gopurs_GoAst.goTypeToStr(expr.value2) + (" {\x0a" + (printGoExpr(v.value1) + "\x0a})")))))));
        };
        return "gopurs_runtime." + (funcName + ("(func(" + (argStr + (") " + (Gopurs_GoAst.goTypeToStr(expr.value2) + (" {\x0areturn " + (printGoExpr(v.value1) + "\x0a})")))))));
    };
    if (expr instanceof Gopurs_GoAst.GoBlock) {
        return Data_String_Common.joinWith("\x0a")(Data_Functor.map(Data_Functor.functorArray)(printGoExpr)(expr.value0));
    };
    if (expr instanceof Gopurs_GoAst.GoReturn) {
        return "return " + printGoExpr(expr.value0);
    };
    if (expr instanceof Gopurs_GoAst.GoAssign) {
        return expr.value0 + (" := " + (printGoExpr(expr.value1) + ("\x0a_ = " + expr.value0)));
    };
    if (expr instanceof Gopurs_GoAst.GoRecordDict) {
        var v = Data_Array.length(expr.value0);
        if (v === 0) {
            return "gopurs_runtime.RecordDict0()";
        };
        if (v === 1) {
            var v1 = Data_Array.index(expr.value0)(0);
            if (v1 instanceof Data_Maybe.Just) {
                return "gopurs_runtime.RecordDict1(\"" + (v1.value0.value0 + ("\", " + (printGoExpr(v1.value0.value1) + ")")));
            };
            if (v1 instanceof Data_Maybe.Nothing) {
                return "";
            };
            throw new Error("Failed pattern match at Gopurs.Printer (line 52, column 12 - line 54, column 22): " + [ v1.constructor.name ]);
        };
        if (v === 2) {
            var v1 = new Data_Tuple.Tuple(Data_Array.index(expr.value0)(0), Data_Array.index(expr.value0)(1));
            if (v1.value0 instanceof Data_Maybe.Just && v1.value1 instanceof Data_Maybe.Just) {
                return "gopurs_runtime.RecordDict2(\"" + (v1.value0.value0.value0 + ("\", \"" + (v1.value1.value0.value0 + ("\", " + (printGoExpr(v1.value0.value0.value1) + (", " + (printGoExpr(v1.value1.value0.value1) + ")")))))));
            };
            return "";
        };
        if (v === 3) {
            var v1 = new Data_Tuple.Tuple(new Data_Tuple.Tuple(Data_Array.index(expr.value0)(0), Data_Array.index(expr.value0)(1)), Data_Array.index(expr.value0)(2));
            if (v1.value0.value0 instanceof Data_Maybe.Just && (v1.value0.value1 instanceof Data_Maybe.Just && v1.value1 instanceof Data_Maybe.Just)) {
                return "gopurs_runtime.RecordDict3(\"" + (v1.value0.value0.value0.value0 + ("\", \"" + (v1.value0.value1.value0.value0 + ("\", \"" + (v1.value1.value0.value0 + ("\", " + (printGoExpr(v1.value0.value0.value0.value1) + (", " + (printGoExpr(v1.value0.value1.value0.value1) + (", " + (printGoExpr(v1.value1.value0.value1) + ")")))))))))));
            };
            return "";
        };
        if (v === 4) {
            var v1 = new Data_Tuple.Tuple(new Data_Tuple.Tuple(Data_Array.index(expr.value0)(0), Data_Array.index(expr.value0)(1)), new Data_Tuple.Tuple(Data_Array.index(expr.value0)(2), Data_Array.index(expr.value0)(3)));
            if (v1.value0.value0 instanceof Data_Maybe.Just && (v1.value0.value1 instanceof Data_Maybe.Just && (v1.value1.value0 instanceof Data_Maybe.Just && v1.value1.value1 instanceof Data_Maybe.Just))) {
                return "gopurs_runtime.RecordDict4(\"" + (v1.value0.value0.value0.value0 + ("\", \"" + (v1.value0.value1.value0.value0 + ("\", \"" + (v1.value1.value0.value0.value0 + ("\", \"" + (v1.value1.value1.value0.value0 + ("\", " + (printGoExpr(v1.value0.value0.value0.value1) + (", " + (printGoExpr(v1.value0.value1.value0.value1) + (", " + (printGoExpr(v1.value1.value0.value0.value1) + (", " + (printGoExpr(v1.value1.value1.value0.value1) + ")")))))))))))))));
            };
            return "";
        };
        if (v === 5) {
            var v1 = new Data_Tuple.Tuple(new Data_Tuple.Tuple(Data_Array.index(expr.value0)(0), Data_Array.index(expr.value0)(1)), new Data_Tuple.Tuple(new Data_Tuple.Tuple(Data_Array.index(expr.value0)(2), Data_Array.index(expr.value0)(3)), Data_Array.index(expr.value0)(4)));
            if (v1.value0.value0 instanceof Data_Maybe.Just && (v1.value0.value1 instanceof Data_Maybe.Just && (v1.value1.value0.value0 instanceof Data_Maybe.Just && (v1.value1.value0.value1 instanceof Data_Maybe.Just && v1.value1.value1 instanceof Data_Maybe.Just)))) {
                return "gopurs_runtime.RecordDict5(\"" + (v1.value0.value0.value0.value0 + ("\", \"" + (v1.value0.value1.value0.value0 + ("\", \"" + (v1.value1.value0.value0.value0.value0 + ("\", \"" + (v1.value1.value0.value1.value0.value0 + ("\", \"" + (v1.value1.value1.value0.value0 + ("\", " + (printGoExpr(v1.value0.value0.value0.value1) + (", " + (printGoExpr(v1.value0.value1.value0.value1) + (", " + (printGoExpr(v1.value1.value0.value0.value0.value1) + (", " + (printGoExpr(v1.value1.value0.value1.value0.value1) + (", " + (printGoExpr(v1.value1.value1.value0.value1) + ")")))))))))))))))))));
            };
            return "";
        };
        var valsStr = Data_String_Common.joinWith(", ")(Data_Functor.map(Data_Functor.functorArray)(function (v1) {
            return printGoExpr(v1.value1);
        })(expr.value0));
        var keysStr = Data_String_Common.joinWith(", ")(Data_Functor.map(Data_Functor.functorArray)(function (v1) {
            return "\"" + (v1.value0 + "\"");
        })(expr.value0));
        return "gopurs_runtime.RecordDict([]string{" + (keysStr + ("}, []gopurs_runtime.Value{" + (valsStr + "})")));
    };
    if (expr instanceof Gopurs_GoAst.GoRecordUpdateDict) {
        var valsStr = Data_String_Common.joinWith(", ")(Data_Functor.map(Data_Functor.functorArray)(function (v) {
            return printGoExpr(v.value1);
        })(expr.value1));
        var len = Data_Array.length(expr.value1);
        var keysStr = Data_String_Common.joinWith(", ")(Data_Functor.map(Data_Functor.functorArray)(function (v) {
            return "\"" + (v.value0 + "\"");
        })(expr.value1));
        var $143 = len === 1;
        if ($143) {
            if (expr.value1.length === 1) {
                return "gopurs_runtime.RecordUpdate1(" + (printGoExpr(expr.value0) + (", \"" + (expr["value1"][0].value0 + ("\", " + (printGoExpr(expr["value1"][0].value1) + ")")))));
            };
            return "gopurs_runtime.RecordUpdateDict(" + (printGoExpr(expr.value0) + (", []string{" + (keysStr + ("}, []gopurs_runtime.Value{" + (valsStr + "})")))));
        };
        var $148 = len === 2;
        if ($148) {
            if (expr.value1.length === 2) {
                return "gopurs_runtime.RecordUpdate2(" + (printGoExpr(expr.value0) + (", \"" + (expr["value1"][0].value0 + ("\", " + (printGoExpr(expr["value1"][0].value1) + (", \"" + (expr["value1"][1].value0 + ("\", " + (printGoExpr(expr["value1"][1].value1) + ")")))))))));
            };
            return "gopurs_runtime.RecordUpdateDict(" + (printGoExpr(expr.value0) + (", []string{" + (keysStr + ("}, []gopurs_runtime.Value{" + (valsStr + "})")))));
        };
        var $156 = len === 3;
        if ($156) {
            if (expr.value1.length === 3) {
                return "gopurs_runtime.RecordUpdate3(" + (printGoExpr(expr.value0) + (", \"" + (expr["value1"][0].value0 + ("\", " + (printGoExpr(expr["value1"][0].value1) + (", \"" + (expr["value1"][1].value0 + ("\", " + (printGoExpr(expr["value1"][1].value1) + (", \"" + (expr["value1"][2].value0 + ("\", " + (printGoExpr(expr["value1"][2].value1) + ")")))))))))))));
            };
            return "gopurs_runtime.RecordUpdateDict(" + (printGoExpr(expr.value0) + (", []string{" + (keysStr + ("}, []gopurs_runtime.Value{" + (valsStr + "})")))));
        };
        return "gopurs_runtime.RecordUpdateDict(" + (printGoExpr(expr.value0) + (", []string{" + (keysStr + ("}, []gopurs_runtime.Value{" + (valsStr + "})")))));
    };
    if (expr instanceof Gopurs_GoAst.GoRecordUpdateStatic) {
        var typeVal = (function () {
            var $169 = expr.value1 >= 6;
            if ($169) {
                return "gopurs_runtime.TypeRecordData";
            };
            return "gopurs_runtime.TypeRecord" + Data_Show.show(Data_Show.showInt)(expr.value1);
        })();
        var structName = (function () {
            var $170 = expr.value1 >= 6;
            if ($170) {
                return "gopurs_runtime.RecordData";
            };
            return "gopurs_runtime.RecordData" + Data_Show.show(Data_Show.showInt)(expr.value1);
        })();
        var fallbackVals = Data_String_Common.joinWith(", ")(Data_Functor.map(Data_Functor.functorArray)(function (v) {
            return printGoExpr(v.value1);
        })(expr.value3));
        var fallbackKeys = Data_String_Common.joinWith(", ")(Data_Functor.map(Data_Functor.functorArray)(function (v) {
            return "\"" + (v.value0 + "\"");
        })(expr.value3));
        var fallbackCall = "gopurs_runtime.RecordUpdateDict(origVal, []string{" + (fallbackKeys + ("}, []gopurs_runtime.Value{" + (fallbackVals + "})")));
        var $177 = expr.value1 >= 6;
        if ($177) {
            var assignments = Data_String_Common.joinWith("\x0a")(Data_Functor.map(Data_Functor.functorArray)(function (v) {
                return "newVals[" + (Data_Show.show(Data_Show.showInt)(v.value0) + ("] = " + printGoExpr(v.value1)));
            })(expr.value2));
            return "func() gopurs_runtime.Value {\x0aorigVal := " + (printGoExpr(expr.value0) + ("\x0aif origVal.Type != " + (typeVal + (" {\x0areturn " + (fallbackCall + ("\x0a}\x0ar := (*" + (structName + (")(origVal.UnsafePtr)\x0anewVals := make([]gopurs_runtime.Value, len(r.Vals))\x0acopy(newVals, r.Vals)\x0a" + (assignments + ("\x0anewR := gopurs_runtime.RecordData{Keys: r.Keys, Vals: newVals}\x0areturn gopurs_runtime.Value{Type: " + (typeVal + ", UnsafePtr: unsafe.Pointer(&newR)}\x0a}()")))))))))));
        };
        var assignments = Data_String_Common.joinWith("\x0a")(Data_Functor.map(Data_Functor.functorArray)(function (v) {
            return "clone.V" + (Data_Show.show(Data_Show.showInt)(v.value0) + (" = " + printGoExpr(v.value1)));
        })(expr.value2));
        return "func() gopurs_runtime.Value {\x0aorigVal := " + (printGoExpr(expr.value0) + ("\x0aif origVal.Type != " + (typeVal + (" {\x0areturn " + (fallbackCall + ("\x0a}\x0aclone := *((*" + (structName + (")(origVal.UnsafePtr))\x0a" + (assignments + ("\x0areturn gopurs_runtime.Value{Type: " + (typeVal + ", UnsafePtr: unsafe.Pointer(&clone)}\x0a}()")))))))))));
    };
    if (expr instanceof Gopurs_GoAst.GoIIFE) {
        var assignment = (function () {
            var $188 = expr.value0 === "_";
            if ($188) {
                return expr.value0 + (" = " + printGoExpr(expr.value1));
            };
            return expr.value0 + (" := " + (printGoExpr(expr.value1) + ("\x0a_ = " + expr.value0)));
        })();
        if (expr.value2 instanceof Gopurs_GoAst.GoBlock) {
            return "func() gopurs_runtime.Value {\x0a" + (assignment + ("\x0a" + (printGoExpr(expr.value2) + "\x0a}()")));
        };
        return "func() gopurs_runtime.Value {\x0a" + (assignment + ("\x0areturn " + (printGoExpr(expr.value2) + "\x0a}()")));
    };
    if (expr instanceof Gopurs_GoAst.GoLetRec) {
        return "func() gopurs_runtime.Value {\x0a" + (Data_String_Common.joinWith("\x0a")(Data_Functor.map(Data_Functor.functorArray)(function (v) {
            return "var " + (v.value0 + " gopurs_runtime.Value");
        })(expr.value0)) + ("\x0a" + (Data_String_Common.joinWith("\x0a")(Data_Functor.map(Data_Functor.functorArray)(function (v) {
            return "_ = " + v.value0;
        })(expr.value0)) + ("\x0a" + (Data_String_Common.joinWith("\x0a")(Data_Functor.map(Data_Functor.functorArray)(function (v) {
            return v.value0 + (" = " + printGoExpr(v.value1));
        })(expr.value0)) + ("\x0a" + ("return " + (printGoExpr(expr.value1) + "\x0a}()"))))))));
    };
    if (expr instanceof Gopurs_GoAst.GoRecordAccess) {
        return "gopurs_runtime.RecordGet(" + (printGoExpr(expr.value0) + (", \"" + (expr.value1 + "\")")));
    };
    if (expr instanceof Gopurs_GoAst.GoStructAccess) {
        return printGoExpr(expr.value0) + ("." + expr.value1);
    };
    if (expr instanceof Gopurs_GoAst.GoRecordAccessStatic) {
        var $209 = expr.value1 >= 6;
        if ($209) {
            return "((*gopurs_runtime.RecordData)(" + (printGoExpr(expr.value0) + (".UnsafePtr)).Vals[" + (Data_Show.show(Data_Show.showInt)(expr.value2) + "]")));
        };
        return "((*gopurs_runtime.RecordData" + (Data_Show.show(Data_Show.showInt)(expr.value1) + (")(" + (printGoExpr(expr.value0) + (".UnsafePtr)).V" + Data_Show.show(Data_Show.showInt)(expr.value2)))));
    };
    if (expr instanceof Gopurs_GoAst.GoConstructor) {
        var typeArgsStr = (function () {
            var $213 = Data_Array.length(expr.value2) > 0;
            if ($213) {
                return "[" + (Data_String_Common.joinWith(", ")(Data_Functor.map(Data_Functor.functorArray)(Gopurs_GoAst.goTypeToStr)(expr.value2)) + "]");
            };
            return "";
        })();
        var $214 = Data_Array["null"](expr.value3);
        if ($214) {
            return "nil";
        };
        return "&" + (expr.value1 + (typeArgsStr + ("{1, " + (Data_String_Common.joinWith(", ")(Data_Functor.map(Data_Functor.functorArray)(printGoExpr)(expr.value3)) + "}"))));
    };
    if (expr instanceof Gopurs_GoAst.GoConstructorAccess) {
        var typeArgsStr = (function () {
            var $219 = Data_Array.length(expr.value2) > 0;
            if ($219) {
                return "[" + (Data_String_Common.joinWith(", ")(Data_Functor.map(Data_Functor.functorArray)(Gopurs_GoAst.goTypeToStr)(expr.value2)) + "]");
            };
            return "";
        })();
        return "(*" + (expr.value1 + (typeArgsStr + (")(" + (printGoExpr(expr.value0) + (".UnsafePtr).V" + Data_Show.show(Data_Show.showInt)(expr.value3))))));
    };
    if (expr instanceof Gopurs_GoAst.GoBranch) {
        return "func() gopurs_runtime.Value {\x0a" + (Data_String_Common.joinWith("\x0a")(Data_Functor.map(Data_Functor.functorArray)(function (v) {
            return "if (" + (printGoExpr(v.value0) + (").IntVal != 0 {\x0areturn " + (printGoExpr(v.value1) + "\x0a}")));
        })(expr.value0)) + ("\x0areturn " + (printGoExpr(expr.value1) + "\x0a}()")));
    };
    if (expr instanceof Gopurs_GoAst.GoBinOp) {
        return "(" + (printGoExpr(expr.value1) + (") " + (expr.value0 + (" (" + (printGoExpr(expr.value2) + ")")))));
    };
    if (expr instanceof Gopurs_GoAst.GoPrefixOp) {
        return expr.value0 + ("(" + (printGoExpr(expr.value1) + ")"));
    };
    if (expr instanceof Gopurs_GoAst.GoTypeAssertion) {
        return printGoExpr(expr.value0) + (".(" + (expr.value1 + ")"));
    };
    if (expr instanceof Gopurs_GoAst.GoIndex) {
        return "(" + (printGoExpr(expr.value0) + (")[" + (printGoExpr(expr.value1) + "]")));
    };
    if (expr instanceof Gopurs_GoAst.GoRaw) {
        return expr.value0;
    };
    if (expr instanceof Gopurs_GoAst.GoFor) {
        return expr.value0 + (":\x0afor {\x0aif false { continue " + (expr.value0 + (" }\x0a" + (Data_String_Common.joinWith("\x0a")(Data_Functor.map(Data_Functor.functorArray)(printGoExpr)(expr.value1)) + "\x0a}"))));
    };
    if (expr instanceof Gopurs_GoAst.GoForRange) {
        return "for " + (expr.value0 + (" {\x0a" + (Data_String_Common.joinWith("\x0a")(Data_Functor.map(Data_Functor.functorArray)(printGoExpr)(expr.value1)) + "\x0a}")));
    };
    if (expr instanceof Gopurs_GoAst.GoContinue) {
        return "continue " + expr.value0;
    };
    if (expr instanceof Gopurs_GoAst.GoIfElse) {
        return "if " + (printGoExpr(expr.value0) + (" {\x0a" + (Data_String_Common.joinWith("\x0a")(Data_Functor.map(Data_Functor.functorArray)(printGoExpr)(expr.value1)) + ("\x0a} else {\x0a" + (Data_String_Common.joinWith("\x0a")(Data_Functor.map(Data_Functor.functorArray)(printGoExpr)(expr.value2)) + "\x0a}")))));
    };
    if (expr instanceof Gopurs_GoAst.GoMutate) {
        return expr.value0 + (" = " + printGoExpr(expr.value1));
    };
    if (expr instanceof Gopurs_GoAst.GoFuncLit) {
        return "func(" + (Data_String_Common.joinWith(", ")(Data_Functor.map(Data_Functor.functorArray)(function (v) {
            return v.value0 + (" " + Gopurs_GoAst.goTypeToStr(v.value1));
        })(expr.value0)) + (") " + (Gopurs_GoAst.goTypeToStr(expr.value3) + (" {\x0a" + (printGoExpr(new Gopurs_GoAst.GoBlock(Data_Semigroup.append(Data_Semigroup.semigroupArray)(expr.value1)([ new Gopurs_GoAst.GoReturn(expr.value2) ]))) + "\x0a}")))));
    };
    throw new Error("Failed pattern match at Gopurs.Printer (line 15, column 20 - line 174, column 204): " + [ expr.constructor.name ]);
};
var printGoDeclVar = function (v) {
    var typeStr = Gopurs_GoAst.goTypeToStr(v.goType);
    return "var cache_" + (v.identifier + (" " + (typeStr + ("\x0a" + ("var once_" + (v.identifier + (" sync.Once\x0a" + ("func Get_" + (v.identifier + ("() " + (typeStr + (" {\x0a" + ("\x09once_" + (v.identifier + (".Do(func() {\x0a" + ("\x09\x09cache_" + (v.identifier + (" = " + (printGoExpr(v.expression) + ("\x0a" + ("\x09})\x0a" + ("\x09return cache_" + (v.identifier + ("\x0a" + "}"))))))))))))))))))))))));
};
var printGoFile = function (v) {
    var missingDeps = [ "Unsafe.Coerce", "Data.Unit", "Control.Monad.ST.Internal", "Data.Eq", "Data.Function.Uncurried", "Control.Category" ];
    var declsStr = Data_String_Common.joinWith("\x0a\x0a")(Data_Functor.map(Data_Functor.functorArray)(printGoDeclVar)(v.decls)) + ("\x0a\x0a" + (Data_String_Common.joinWith("\x0a\x0a")(v.rawDecls) + ("\x0a\x0a" + (Data_String_Common.joinWith("\x0a\x0a")(Data_Functor.map(Data_Functor.functorArray)(function (f) {
        return "func Get_" + (f.pursName + ("() gopurs_runtime.Value {\x0a\x09return " + (f.goName + "\x0a}")));
    })(v.foreigns)) + "\x0a"))));
    var unsafeImport = (function () {
        var $261 = Data_String_CodeUnits.contains("unsafe.")(declsStr);
        if ($261) {
            return [ "unsafe" ];
        };
        return [  ];
    })();
    var usedImports1 = Data_Array.filter(function (i) {
        var $262 = i === "gopurs/output/gopurs_runtime" || i === "sync";
        if ($262) {
            return true;
        };
        var pkg = Data_Array.last(Data_String_Common.split("/")(i));
        var pkgAlias = "pkg_" + Data_String_Common.replaceAll(".")("_")(Data_Maybe.fromMaybe("")(pkg));
        return Data_String_CodeUnits.contains(pkgAlias + ".")(declsStr);
    })(v.imports);
    var injectedDeps = Data_Array.filter(function (dep) {
        var pkgAlias = "pkg_" + Data_String_Common.replaceAll(".")("_")(dep);
        return Data_String_CodeUnits.contains(pkgAlias + ".")(declsStr) && !Data_Array.elem(Data_Eq.eqString)("gopurs/output/" + dep)(usedImports1);
    })(missingDeps);
    var usedImports = Data_Semigroup.append(Data_Semigroup.semigroupArray)(usedImports1)(Data_Semigroup.append(Data_Semigroup.semigroupArray)(Data_Functor.map(Data_Functor.functorArray)(function (dep) {
        return "gopurs/output/" + dep;
    })(injectedDeps))(unsafeImport));
    return "package " + (v.packageName + ("\x0a\x0a" + ("import (\x0a" + (Data_String_Common.joinWith("\x0a")(Data_Functor.map(Data_Functor.functorArray)(function (i) {
        var pkg = Data_Array.last(Data_String_Common.split("/")(i));
        var pkgAlias = (function () {
            var $263 = i === "gopurs/output/gopurs_runtime" || (i === "sync" || i === "unsafe");
            if ($263) {
                return Data_Maybe.fromMaybe("")(pkg);
            };
            return "pkg_" + Data_String_Common.replaceAll(".")("_")(Data_Maybe.fromMaybe("")(pkg));
        })();
        return "\x09" + (pkgAlias + (" \"" + (i + "\"")));
    })(usedImports)) + ("\x0a" + (")\x0a\x0a" + declsStr))))));
};
export {
    escapeGoStringImpl
} from "./foreign.js";
export {
    escapeGoString,
    printGoExpr,
    printGoDeclVar,
    printGoFile
};
