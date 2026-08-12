import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_String_Common from "../Data.String.Common/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
var eqTuple = /* #__PURE__ */ Data_Tuple.eqTuple(Data_Eq.eqString);
var eqTuple1 = /* #__PURE__ */ Data_Tuple.eqTuple(Data_Eq.eqInt);
var TypeValue = /* #__PURE__ */ (function () {
    function TypeValue() {

    };
    TypeValue.value = new TypeValue();
    return TypeValue;
})();
var TypeInt64 = /* #__PURE__ */ (function () {
    function TypeInt64() {

    };
    TypeInt64.value = new TypeInt64();
    return TypeInt64;
})();
var TypeFloat64 = /* #__PURE__ */ (function () {
    function TypeFloat64() {

    };
    TypeFloat64.value = new TypeFloat64();
    return TypeFloat64;
})();
var TypeString = /* #__PURE__ */ (function () {
    function TypeString() {

    };
    TypeString.value = new TypeString();
    return TypeString;
})();
var TypeBool = /* #__PURE__ */ (function () {
    function TypeBool() {

    };
    TypeBool.value = new TypeBool();
    return TypeBool;
})();
var TypeStructPointer = /* #__PURE__ */ (function () {
    function TypeStructPointer(value0, value1, value2, value3) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
        this.value3 = value3;
    };
    TypeStructPointer.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return function (value3) {
                    return new TypeStructPointer(value0, value1, value2, value3);
                };
            };
        };
    };
    return TypeStructPointer;
})();
var TypeRecord = /* #__PURE__ */ (function () {
    function TypeRecord(value0) {
        this.value0 = value0;
    };
    TypeRecord.create = function (value0) {
        return new TypeRecord(value0);
    };
    return TypeRecord;
})();
var TypeInterface = /* #__PURE__ */ (function () {
    function TypeInterface(value0) {
        this.value0 = value0;
    };
    TypeInterface.create = function (value0) {
        return new TypeInterface(value0);
    };
    return TypeInterface;
})();
var TypeNativeArray = /* #__PURE__ */ (function () {
    function TypeNativeArray(value0) {
        this.value0 = value0;
    };
    TypeNativeArray.create = function (value0) {
        return new TypeNativeArray(value0);
    };
    return TypeNativeArray;
})();
var TypeGenericParam = /* #__PURE__ */ (function () {
    function TypeGenericParam(value0) {
        this.value0 = value0;
    };
    TypeGenericParam.create = function (value0) {
        return new TypeGenericParam(value0);
    };
    return TypeGenericParam;
})();
var TypeFunc = /* #__PURE__ */ (function () {
    function TypeFunc(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    TypeFunc.create = function (value0) {
        return function (value1) {
            return new TypeFunc(value0, value1);
        };
    };
    return TypeFunc;
})();
var GoVar = /* #__PURE__ */ (function () {
    function GoVar(value0) {
        this.value0 = value0;
    };
    GoVar.create = function (value0) {
        return new GoVar(value0);
    };
    return GoVar;
})();
var GoString = /* #__PURE__ */ (function () {
    function GoString(value0) {
        this.value0 = value0;
    };
    GoString.create = function (value0) {
        return new GoString(value0);
    };
    return GoString;
})();
var GoInt = /* #__PURE__ */ (function () {
    function GoInt(value0) {
        this.value0 = value0;
    };
    GoInt.create = function (value0) {
        return new GoInt(value0);
    };
    return GoInt;
})();
var GoCall = /* #__PURE__ */ (function () {
    function GoCall(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    GoCall.create = function (value0) {
        return function (value1) {
            return new GoCall(value0, value1);
        };
    };
    return GoCall;
})();
var GoSelector = /* #__PURE__ */ (function () {
    function GoSelector(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    GoSelector.create = function (value0) {
        return function (value1) {
            return new GoSelector(value0, value1);
        };
    };
    return GoSelector;
})();
var GoFunc = /* #__PURE__ */ (function () {
    function GoFunc(value0, value1, value2, value3) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
        this.value3 = value3;
    };
    GoFunc.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return function (value3) {
                    return new GoFunc(value0, value1, value2, value3);
                };
            };
        };
    };
    return GoFunc;
})();
var GoBlock = /* #__PURE__ */ (function () {
    function GoBlock(value0) {
        this.value0 = value0;
    };
    GoBlock.create = function (value0) {
        return new GoBlock(value0);
    };
    return GoBlock;
})();
var GoReturn = /* #__PURE__ */ (function () {
    function GoReturn(value0) {
        this.value0 = value0;
    };
    GoReturn.create = function (value0) {
        return new GoReturn(value0);
    };
    return GoReturn;
})();
var GoAssign = /* #__PURE__ */ (function () {
    function GoAssign(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    GoAssign.create = function (value0) {
        return function (value1) {
            return new GoAssign(value0, value1);
        };
    };
    return GoAssign;
})();
var GoRecordDict = /* #__PURE__ */ (function () {
    function GoRecordDict(value0) {
        this.value0 = value0;
    };
    GoRecordDict.create = function (value0) {
        return new GoRecordDict(value0);
    };
    return GoRecordDict;
})();
var GoRecordUpdateDict = /* #__PURE__ */ (function () {
    function GoRecordUpdateDict(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    GoRecordUpdateDict.create = function (value0) {
        return function (value1) {
            return new GoRecordUpdateDict(value0, value1);
        };
    };
    return GoRecordUpdateDict;
})();
var GoRecordUpdateStatic = /* #__PURE__ */ (function () {
    function GoRecordUpdateStatic(value0, value1, value2, value3) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
        this.value3 = value3;
    };
    GoRecordUpdateStatic.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return function (value3) {
                    return new GoRecordUpdateStatic(value0, value1, value2, value3);
                };
            };
        };
    };
    return GoRecordUpdateStatic;
})();
var GoIIFE = /* #__PURE__ */ (function () {
    function GoIIFE(value0, value1, value2) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
    };
    GoIIFE.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return new GoIIFE(value0, value1, value2);
            };
        };
    };
    return GoIIFE;
})();
var GoLetRec = /* #__PURE__ */ (function () {
    function GoLetRec(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    GoLetRec.create = function (value0) {
        return function (value1) {
            return new GoLetRec(value0, value1);
        };
    };
    return GoLetRec;
})();
var GoRecordAccess = /* #__PURE__ */ (function () {
    function GoRecordAccess(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    GoRecordAccess.create = function (value0) {
        return function (value1) {
            return new GoRecordAccess(value0, value1);
        };
    };
    return GoRecordAccess;
})();
var GoStructAccess = /* #__PURE__ */ (function () {
    function GoStructAccess(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    GoStructAccess.create = function (value0) {
        return function (value1) {
            return new GoStructAccess(value0, value1);
        };
    };
    return GoStructAccess;
})();
var GoRecordAccessStatic = /* #__PURE__ */ (function () {
    function GoRecordAccessStatic(value0, value1, value2) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
    };
    GoRecordAccessStatic.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return new GoRecordAccessStatic(value0, value1, value2);
            };
        };
    };
    return GoRecordAccessStatic;
})();
var GoConstructor = /* #__PURE__ */ (function () {
    function GoConstructor(value0, value1, value2, value3) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
        this.value3 = value3;
    };
    GoConstructor.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return function (value3) {
                    return new GoConstructor(value0, value1, value2, value3);
                };
            };
        };
    };
    return GoConstructor;
})();
var GoConstructorAccess = /* #__PURE__ */ (function () {
    function GoConstructorAccess(value0, value1, value2, value3) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
        this.value3 = value3;
    };
    GoConstructorAccess.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return function (value3) {
                    return new GoConstructorAccess(value0, value1, value2, value3);
                };
            };
        };
    };
    return GoConstructorAccess;
})();
var GoBranch = /* #__PURE__ */ (function () {
    function GoBranch(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    GoBranch.create = function (value0) {
        return function (value1) {
            return new GoBranch(value0, value1);
        };
    };
    return GoBranch;
})();
var GoBinOp = /* #__PURE__ */ (function () {
    function GoBinOp(value0, value1, value2) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
    };
    GoBinOp.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return new GoBinOp(value0, value1, value2);
            };
        };
    };
    return GoBinOp;
})();
var GoPrefixOp = /* #__PURE__ */ (function () {
    function GoPrefixOp(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    GoPrefixOp.create = function (value0) {
        return function (value1) {
            return new GoPrefixOp(value0, value1);
        };
    };
    return GoPrefixOp;
})();
var GoTypeAssertion = /* #__PURE__ */ (function () {
    function GoTypeAssertion(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    GoTypeAssertion.create = function (value0) {
        return function (value1) {
            return new GoTypeAssertion(value0, value1);
        };
    };
    return GoTypeAssertion;
})();
var GoIndex = /* #__PURE__ */ (function () {
    function GoIndex(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    GoIndex.create = function (value0) {
        return function (value1) {
            return new GoIndex(value0, value1);
        };
    };
    return GoIndex;
})();
var GoRaw = /* #__PURE__ */ (function () {
    function GoRaw(value0) {
        this.value0 = value0;
    };
    GoRaw.create = function (value0) {
        return new GoRaw(value0);
    };
    return GoRaw;
})();
var GoFor = /* #__PURE__ */ (function () {
    function GoFor(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    GoFor.create = function (value0) {
        return function (value1) {
            return new GoFor(value0, value1);
        };
    };
    return GoFor;
})();
var GoForRange = /* #__PURE__ */ (function () {
    function GoForRange(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    GoForRange.create = function (value0) {
        return function (value1) {
            return new GoForRange(value0, value1);
        };
    };
    return GoForRange;
})();
var GoContinue = /* #__PURE__ */ (function () {
    function GoContinue(value0) {
        this.value0 = value0;
    };
    GoContinue.create = function (value0) {
        return new GoContinue(value0);
    };
    return GoContinue;
})();
var GoMutate = /* #__PURE__ */ (function () {
    function GoMutate(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    GoMutate.create = function (value0) {
        return function (value1) {
            return new GoMutate(value0, value1);
        };
    };
    return GoMutate;
})();
var GoIfElse = /* #__PURE__ */ (function () {
    function GoIfElse(value0, value1, value2) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
    };
    GoIfElse.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return new GoIfElse(value0, value1, value2);
            };
        };
    };
    return GoIfElse;
})();
var GoFuncLit = /* #__PURE__ */ (function () {
    function GoFuncLit(value0, value1, value2, value3) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
        this.value3 = value3;
    };
    GoFuncLit.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return function (value3) {
                    return new GoFuncLit(value0, value1, value2, value3);
                };
            };
        };
    };
    return GoFuncLit;
})();
var sanitizeName = function (name) {
    var s1 = Data_String_Common.replaceAll("\"")("_quote_")(Data_String_Common.replaceAll(".")("_dot_")(Data_String_Common.replaceAll("'")("_prime")(Data_String_Common.replaceAll("$")("_dollar")(name))));
    var $166 = s1 === "break" || (s1 === "default" || (s1 === "func" || (s1 === "interface" || (s1 === "select" || (s1 === "case" || (s1 === "defer" || (s1 === "go" || (s1 === "map" || (s1 === "struct" || (s1 === "chan" || (s1 === "else" || (s1 === "goto" || (s1 === "package" || (s1 === "switch" || (s1 === "const" || (s1 === "fallthrough" || (s1 === "if" || (s1 === "range" || (s1 === "type" || (s1 === "continue" || (s1 === "for" || (s1 === "import" || (s1 === "return" || (s1 === "var" || (s1 === "init" || (s1 === "append" || (s1 === "make" || (s1 === "len" || (s1 === "cap" || (s1 === "new" || (s1 === "close" || (s1 === "delete" || (s1 === "complex" || (s1 === "real" || (s1 === "imag" || (s1 === "panic" || (s1 === "recover" || (s1 === "print" || s1 === "println"))))))))))))))))))))))))))))))))))))));
    if ($166) {
        return "go__" + s1;
    };
    return s1;
};
var goTypeToStr = function (v) {
    if (v instanceof TypeInt64) {
        return "int64";
    };
    if (v instanceof TypeFloat64) {
        return "float64";
    };
    if (v instanceof TypeString) {
        return "string";
    };
    if (v instanceof TypeBool) {
        return "bool";
    };
    if (v instanceof TypeStructPointer) {
        return "*" + v.value2;
    };
    if (v instanceof TypeInterface) {
        return v.value0;
    };
    if (v instanceof TypeNativeArray) {
        return "[]" + goTypeToStr(v.value0);
    };
    if (v instanceof TypeGenericParam) {
        return "T_" + sanitizeName(v.value0);
    };
    if (v instanceof TypeFunc) {
        return "func(" + (Data_String_Common.joinWith(", ")(Data_Functor.map(Data_Functor.functorArray)(goTypeToStr)(v.value0)) + (") " + goTypeToStr(v.value1)));
    };
    return "gopurs_runtime.Value";
};
var eqGoType = {
    eq: function (x) {
        return function (y) {
            if (x instanceof TypeValue && y instanceof TypeValue) {
                return true;
            };
            if (x instanceof TypeInt64 && y instanceof TypeInt64) {
                return true;
            };
            if (x instanceof TypeFloat64 && y instanceof TypeFloat64) {
                return true;
            };
            if (x instanceof TypeString && y instanceof TypeString) {
                return true;
            };
            if (x instanceof TypeBool && y instanceof TypeBool) {
                return true;
            };
            if (x instanceof TypeStructPointer && y instanceof TypeStructPointer) {
                return x.value0 === y.value0 && x.value1 === y.value1 && x.value2 === y.value2 && Data_Eq.eq(Data_Eq.eqArray(eqGoType))(x.value3)(y.value3);
            };
            if (x instanceof TypeRecord && y instanceof TypeRecord) {
                return Data_Eq.eq(Data_Eq.eqArray(eqTuple(eqGoType)))(x.value0)(y.value0);
            };
            if (x instanceof TypeInterface && y instanceof TypeInterface) {
                return x.value0 === y.value0;
            };
            if (x instanceof TypeNativeArray && y instanceof TypeNativeArray) {
                return Data_Eq.eq(eqGoType)(x.value0)(y.value0);
            };
            if (x instanceof TypeGenericParam && y instanceof TypeGenericParam) {
                return x.value0 === y.value0;
            };
            if (x instanceof TypeFunc && y instanceof TypeFunc) {
                return Data_Eq.eq(Data_Eq.eqArray(eqGoType))(x.value0)(y.value0) && Data_Eq.eq(eqGoType)(x.value1)(y.value1);
            };
            return false;
        };
    }
};
var eqArray = /* #__PURE__ */ Data_Eq.eqArray(eqGoType);
var eqArray1 = /* #__PURE__ */ Data_Eq.eqArray(/* #__PURE__ */ eqTuple(eqGoType));
var eqGoExpr = {
    eq: function (x) {
        return function (y) {
            if (x instanceof GoVar && y instanceof GoVar) {
                return x.value0 === y.value0;
            };
            if (x instanceof GoString && y instanceof GoString) {
                return x.value0 === y.value0;
            };
            if (x instanceof GoInt && y instanceof GoInt) {
                return x.value0 === y.value0;
            };
            if (x instanceof GoCall && y instanceof GoCall) {
                return Data_Eq.eq(eqGoExpr)(x.value0)(y.value0) && Data_Eq.eq(Data_Eq.eqArray(eqGoExpr))(x.value1)(y.value1);
            };
            if (x instanceof GoSelector && y instanceof GoSelector) {
                return Data_Eq.eq(eqGoExpr)(x.value0)(y.value0) && x.value1 === y.value1;
            };
            if (x instanceof GoFunc && y instanceof GoFunc) {
                return x.value0 === y.value0 && Data_Eq.eq(eqGoType)(x.value1)(y.value1) && Data_Eq.eq(eqGoType)(x.value2)(y.value2) && Data_Eq.eq(eqGoExpr)(x.value3)(y.value3);
            };
            if (x instanceof GoBlock && y instanceof GoBlock) {
                return Data_Eq.eq(Data_Eq.eqArray(eqGoExpr))(x.value0)(y.value0);
            };
            if (x instanceof GoReturn && y instanceof GoReturn) {
                return Data_Eq.eq(eqGoExpr)(x.value0)(y.value0);
            };
            if (x instanceof GoAssign && y instanceof GoAssign) {
                return x.value0 === y.value0 && Data_Eq.eq(eqGoExpr)(x.value1)(y.value1);
            };
            if (x instanceof GoRecordDict && y instanceof GoRecordDict) {
                return Data_Eq.eq(Data_Eq.eqArray(eqTuple(eqGoExpr)))(x.value0)(y.value0);
            };
            if (x instanceof GoRecordUpdateDict && y instanceof GoRecordUpdateDict) {
                return Data_Eq.eq(eqGoExpr)(x.value0)(y.value0) && Data_Eq.eq(Data_Eq.eqArray(eqTuple(eqGoExpr)))(x.value1)(y.value1);
            };
            if (x instanceof GoRecordUpdateStatic && y instanceof GoRecordUpdateStatic) {
                return Data_Eq.eq(eqGoExpr)(x.value0)(y.value0) && x.value1 === y.value1 && Data_Eq.eq(Data_Eq.eqArray(eqTuple1(eqGoExpr)))(x.value2)(y.value2) && Data_Eq.eq(Data_Eq.eqArray(eqTuple(eqGoExpr)))(x.value3)(y.value3);
            };
            if (x instanceof GoIIFE && y instanceof GoIIFE) {
                return x.value0 === y.value0 && Data_Eq.eq(eqGoExpr)(x.value1)(y.value1) && Data_Eq.eq(eqGoExpr)(x.value2)(y.value2);
            };
            if (x instanceof GoLetRec && y instanceof GoLetRec) {
                return Data_Eq.eq(Data_Eq.eqArray(eqTuple(eqGoExpr)))(x.value0)(y.value0) && Data_Eq.eq(eqGoExpr)(x.value1)(y.value1);
            };
            if (x instanceof GoRecordAccess && y instanceof GoRecordAccess) {
                return Data_Eq.eq(eqGoExpr)(x.value0)(y.value0) && x.value1 === y.value1;
            };
            if (x instanceof GoStructAccess && y instanceof GoStructAccess) {
                return Data_Eq.eq(eqGoExpr)(x.value0)(y.value0) && x.value1 === y.value1;
            };
            if (x instanceof GoRecordAccessStatic && y instanceof GoRecordAccessStatic) {
                return Data_Eq.eq(eqGoExpr)(x.value0)(y.value0) && x.value1 === y.value1 && x.value2 === y.value2;
            };
            if (x instanceof GoConstructor && y instanceof GoConstructor) {
                return x.value0 === y.value0 && x.value1 === y.value1 && Data_Eq.eq(eqArray)(x.value2)(y.value2) && Data_Eq.eq(Data_Eq.eqArray(eqGoExpr))(x.value3)(y.value3);
            };
            if (x instanceof GoConstructorAccess && y instanceof GoConstructorAccess) {
                return Data_Eq.eq(eqGoExpr)(x.value0)(y.value0) && x.value1 === y.value1 && Data_Eq.eq(eqArray)(x.value2)(y.value2) && x.value3 === y.value3;
            };
            if (x instanceof GoBranch && y instanceof GoBranch) {
                return Data_Eq.eq(Data_Eq.eqArray(Data_Tuple.eqTuple(eqGoExpr)(eqGoExpr)))(x.value0)(y.value0) && Data_Eq.eq(eqGoExpr)(x.value1)(y.value1);
            };
            if (x instanceof GoBinOp && y instanceof GoBinOp) {
                return x.value0 === y.value0 && Data_Eq.eq(eqGoExpr)(x.value1)(y.value1) && Data_Eq.eq(eqGoExpr)(x.value2)(y.value2);
            };
            if (x instanceof GoPrefixOp && y instanceof GoPrefixOp) {
                return x.value0 === y.value0 && Data_Eq.eq(eqGoExpr)(x.value1)(y.value1);
            };
            if (x instanceof GoTypeAssertion && y instanceof GoTypeAssertion) {
                return Data_Eq.eq(eqGoExpr)(x.value0)(y.value0) && x.value1 === y.value1;
            };
            if (x instanceof GoIndex && y instanceof GoIndex) {
                return Data_Eq.eq(eqGoExpr)(x.value0)(y.value0) && Data_Eq.eq(eqGoExpr)(x.value1)(y.value1);
            };
            if (x instanceof GoRaw && y instanceof GoRaw) {
                return x.value0 === y.value0;
            };
            if (x instanceof GoFor && y instanceof GoFor) {
                return x.value0 === y.value0 && Data_Eq.eq(Data_Eq.eqArray(eqGoExpr))(x.value1)(y.value1);
            };
            if (x instanceof GoForRange && y instanceof GoForRange) {
                return x.value0 === y.value0 && Data_Eq.eq(Data_Eq.eqArray(eqGoExpr))(x.value1)(y.value1);
            };
            if (x instanceof GoContinue && y instanceof GoContinue) {
                return x.value0 === y.value0;
            };
            if (x instanceof GoMutate && y instanceof GoMutate) {
                return x.value0 === y.value0 && Data_Eq.eq(eqGoExpr)(x.value1)(y.value1);
            };
            if (x instanceof GoIfElse && y instanceof GoIfElse) {
                return Data_Eq.eq(eqGoExpr)(x.value0)(y.value0) && Data_Eq.eq(Data_Eq.eqArray(eqGoExpr))(x.value1)(y.value1) && Data_Eq.eq(Data_Eq.eqArray(eqGoExpr))(x.value2)(y.value2);
            };
            if (x instanceof GoFuncLit && y instanceof GoFuncLit) {
                return Data_Eq.eq(eqArray1)(x.value0)(y.value0) && Data_Eq.eq(Data_Eq.eqArray(eqGoExpr))(x.value1)(y.value1) && Data_Eq.eq(eqGoExpr)(x.value2)(y.value2) && Data_Eq.eq(eqGoType)(x.value3)(y.value3);
            };
            return false;
        };
    }
};
export {
    GoVar,
    GoString,
    GoInt,
    GoCall,
    GoSelector,
    GoFunc,
    GoBlock,
    GoReturn,
    GoAssign,
    GoRecordDict,
    GoRecordUpdateDict,
    GoRecordUpdateStatic,
    GoIIFE,
    GoLetRec,
    GoRecordAccess,
    GoStructAccess,
    GoRecordAccessStatic,
    GoConstructor,
    GoConstructorAccess,
    GoBranch,
    GoBinOp,
    GoPrefixOp,
    GoTypeAssertion,
    GoIndex,
    GoRaw,
    GoFor,
    GoForRange,
    GoContinue,
    GoMutate,
    GoIfElse,
    GoFuncLit,
    TypeValue,
    TypeInt64,
    TypeFloat64,
    TypeString,
    TypeBool,
    TypeStructPointer,
    TypeRecord,
    TypeInterface,
    TypeNativeArray,
    TypeGenericParam,
    TypeFunc,
    goTypeToStr,
    sanitizeName,
    eqGoExpr,
    eqGoType
};
