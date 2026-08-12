import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Data_Array from "../Data.Array/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Ordering from "../Data.Ordering/index.js";
import * as Data_Semiring from "../Data.Semiring/index.js";
import * as Data_String_CodeUnits from "../Data.String.CodeUnits/index.js";
import * as Data_Traversable from "../Data.Traversable/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
var eqArray = /* #__PURE__ */ Data_Eq.eqArray(Data_Eq.eqString);
var eqTuple = /* #__PURE__ */ Data_Tuple.eqTuple(Data_Eq.eqString);
var eqTuple1 = /* #__PURE__ */ Data_Tuple.eqTuple(/* #__PURE__ */ Data_Eq.eqArray(Data_Eq.eqString));
var ordArray = /* #__PURE__ */ Data_Ord.ordArray(Data_Ord.ordString);
var ordTuple = /* #__PURE__ */ Data_Tuple.ordTuple(Data_Ord.ordString);
var ordTuple1 = /* #__PURE__ */ Data_Tuple.ordTuple(/* #__PURE__ */ Data_Ord.ordArray(Data_Ord.ordString));
var zero = /* #__PURE__ */ Data_Semiring.zero(/* #__PURE__ */ Data_Semiring.semiringRecord()(/* #__PURE__ */ Data_Semiring.semiringRecordCons({
    reflectSymbol: function () {
        return "column";
    }
})()(/* #__PURE__ */ Data_Semiring.semiringRecordCons({
    reflectSymbol: function () {
        return "line";
    }
})()(Data_Semiring.semiringRecordNil)(Data_Semiring.semiringInt))(Data_Semiring.semiringInt)));
var ProperName = function (x) {
    return x;
};
var Prop = /* #__PURE__ */ (function () {
    function Prop(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    Prop.create = function (value0) {
        return function (value1) {
            return new Prop(value0, value1);
        };
    };
    return Prop;
})();
var ModuleName = function (x) {
    return x;
};
var Qualified = /* #__PURE__ */ (function () {
    function Qualified(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    Qualified.create = function (value0) {
        return function (value1) {
            return new Qualified(value0, value1);
        };
    };
    return Qualified;
})();
var LitInt = /* #__PURE__ */ (function () {
    function LitInt(value0) {
        this.value0 = value0;
    };
    LitInt.create = function (value0) {
        return new LitInt(value0);
    };
    return LitInt;
})();
var LitNumber = /* #__PURE__ */ (function () {
    function LitNumber(value0) {
        this.value0 = value0;
    };
    LitNumber.create = function (value0) {
        return new LitNumber(value0);
    };
    return LitNumber;
})();
var LitString = /* #__PURE__ */ (function () {
    function LitString(value0) {
        this.value0 = value0;
    };
    LitString.create = function (value0) {
        return new LitString(value0);
    };
    return LitString;
})();
var LitChar = /* #__PURE__ */ (function () {
    function LitChar(value0) {
        this.value0 = value0;
    };
    LitChar.create = function (value0) {
        return new LitChar(value0);
    };
    return LitChar;
})();
var LitBoolean = /* #__PURE__ */ (function () {
    function LitBoolean(value0) {
        this.value0 = value0;
    };
    LitBoolean.create = function (value0) {
        return new LitBoolean(value0);
    };
    return LitBoolean;
})();
var LitArray = /* #__PURE__ */ (function () {
    function LitArray(value0) {
        this.value0 = value0;
    };
    LitArray.create = function (value0) {
        return new LitArray(value0);
    };
    return LitArray;
})();
var LitRecord = /* #__PURE__ */ (function () {
    function LitRecord(value0) {
        this.value0 = value0;
    };
    LitRecord.create = function (value0) {
        return new LitRecord(value0);
    };
    return LitRecord;
})();
var Import = /* #__PURE__ */ (function () {
    function Import(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    Import.create = function (value0) {
        return function (value1) {
            return new Import(value0, value1);
        };
    };
    return Import;
})();
var Ident = function (x) {
    return x;
};
var ReExport = /* #__PURE__ */ (function () {
    function ReExport(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    ReExport.create = function (value0) {
        return function (value1) {
            return new ReExport(value0, value1);
        };
    };
    return ReExport;
})();
var Int = /* #__PURE__ */ (function () {
    function Int() {

    };
    Int.value = new Int();
    return Int;
})();
var $$Number = /* #__PURE__ */ (function () {
    function $$Number() {

    };
    $$Number.value = new $$Number();
    return $$Number;
})();
var $$String = /* #__PURE__ */ (function () {
    function $$String() {

    };
    $$String.value = new $$String();
    return $$String;
})();
var Char = /* #__PURE__ */ (function () {
    function Char() {

    };
    Char.value = new Char();
    return Char;
})();
var $$Boolean = /* #__PURE__ */ (function () {
    function $$Boolean() {

    };
    $$Boolean.value = new $$Boolean();
    return $$Boolean;
})();
var Unit = /* #__PURE__ */ (function () {
    function Unit() {

    };
    Unit.value = new Unit();
    return Unit;
})();
var Any = /* #__PURE__ */ (function () {
    function Any() {

    };
    Any.value = new Any();
    return Any;
})();
var TypeLevelString = /* #__PURE__ */ (function () {
    function TypeLevelString(value0) {
        this.value0 = value0;
    };
    TypeLevelString.create = function (value0) {
        return new TypeLevelString(value0);
    };
    return TypeLevelString;
})();
var $$Array = /* #__PURE__ */ (function () {
    function $$Array(value0) {
        this.value0 = value0;
    };
    $$Array.create = function (value0) {
        return new $$Array(value0);
    };
    return $$Array;
})();
var TypeVar = /* #__PURE__ */ (function () {
    function TypeVar(value0) {
        this.value0 = value0;
    };
    TypeVar.create = function (value0) {
        return new TypeVar(value0);
    };
    return TypeVar;
})();
var ADT = /* #__PURE__ */ (function () {
    function ADT(value0, value1, value2) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
    };
    ADT.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return new ADT(value0, value1, value2);
            };
        };
    };
    return ADT;
})();
var TypeApp = /* #__PURE__ */ (function () {
    function TypeApp(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    TypeApp.create = function (value0) {
        return function (value1) {
            return new TypeApp(value0, value1);
        };
    };
    return TypeApp;
})();
var Func = /* #__PURE__ */ (function () {
    function Func(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    Func.create = function (value0) {
        return function (value1) {
            return new Func(value0, value1);
        };
    };
    return Func;
})();
var Row = /* #__PURE__ */ (function () {
    function Row(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    Row.create = function (value0) {
        return function (value1) {
            return new Row(value0, value1);
        };
    };
    return Row;
})();
var Record = /* #__PURE__ */ (function () {
    function Record(value0) {
        this.value0 = value0;
    };
    Record.create = function (value0) {
        return new Record(value0);
    };
    return Record;
})();
var ForAll = /* #__PURE__ */ (function () {
    function ForAll(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    ForAll.create = function (value0) {
        return function (value1) {
            return new ForAll(value0, value1);
        };
    };
    return ForAll;
})();
var ConstrainedType = /* #__PURE__ */ (function () {
    function ConstrainedType(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    ConstrainedType.create = function (value0) {
        return function (value1) {
            return new ConstrainedType(value0, value1);
        };
    };
    return ConstrainedType;
})();
var ProductType = /* #__PURE__ */ (function () {
    function ProductType() {

    };
    ProductType.value = new ProductType();
    return ProductType;
})();
var SumType = /* #__PURE__ */ (function () {
    function SumType() {

    };
    SumType.value = new SumType();
    return SumType;
})();
var IsConstructor = /* #__PURE__ */ (function () {
    function IsConstructor(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    IsConstructor.create = function (value0) {
        return function (value1) {
            return new IsConstructor(value0, value1);
        };
    };
    return IsConstructor;
})();
var IsNewtype = /* #__PURE__ */ (function () {
    function IsNewtype() {

    };
    IsNewtype.value = new IsNewtype();
    return IsNewtype;
})();
var IsTypeClassConstructor = /* #__PURE__ */ (function () {
    function IsTypeClassConstructor() {

    };
    IsTypeClassConstructor.value = new IsTypeClassConstructor();
    return IsTypeClassConstructor;
})();
var IsForeign = /* #__PURE__ */ (function () {
    function IsForeign() {

    };
    IsForeign.value = new IsForeign();
    return IsForeign;
})();
var IsWhere = /* #__PURE__ */ (function () {
    function IsWhere() {

    };
    IsWhere.value = new IsWhere();
    return IsWhere;
})();
var IsSyntheticApp = /* #__PURE__ */ (function () {
    function IsSyntheticApp() {

    };
    IsSyntheticApp.value = new IsSyntheticApp();
    return IsSyntheticApp;
})();
var LineComment = /* #__PURE__ */ (function () {
    function LineComment(value0) {
        this.value0 = value0;
    };
    LineComment.create = function (value0) {
        return new LineComment(value0);
    };
    return LineComment;
})();
var BlockComment = /* #__PURE__ */ (function () {
    function BlockComment(value0) {
        this.value0 = value0;
    };
    BlockComment.create = function (value0) {
        return new BlockComment(value0);
    };
    return BlockComment;
})();
var BinderNull = /* #__PURE__ */ (function () {
    function BinderNull(value0) {
        this.value0 = value0;
    };
    BinderNull.create = function (value0) {
        return new BinderNull(value0);
    };
    return BinderNull;
})();
var BinderVar = /* #__PURE__ */ (function () {
    function BinderVar(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    BinderVar.create = function (value0) {
        return function (value1) {
            return new BinderVar(value0, value1);
        };
    };
    return BinderVar;
})();
var BinderNamed = /* #__PURE__ */ (function () {
    function BinderNamed(value0, value1, value2) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
    };
    BinderNamed.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return new BinderNamed(value0, value1, value2);
            };
        };
    };
    return BinderNamed;
})();
var BinderLit = /* #__PURE__ */ (function () {
    function BinderLit(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    BinderLit.create = function (value0) {
        return function (value1) {
            return new BinderLit(value0, value1);
        };
    };
    return BinderLit;
})();
var BinderConstructor = /* #__PURE__ */ (function () {
    function BinderConstructor(value0, value1, value2, value3) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
        this.value3 = value3;
    };
    BinderConstructor.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return function (value3) {
                    return new BinderConstructor(value0, value1, value2, value3);
                };
            };
        };
    };
    return BinderConstructor;
})();
var NonRec = /* #__PURE__ */ (function () {
    function NonRec(value0) {
        this.value0 = value0;
    };
    NonRec.create = function (value0) {
        return new NonRec(value0);
    };
    return NonRec;
})();
var Rec = /* #__PURE__ */ (function () {
    function Rec(value0) {
        this.value0 = value0;
    };
    Rec.create = function (value0) {
        return new Rec(value0);
    };
    return Rec;
})();
var Binding = /* #__PURE__ */ (function () {
    function Binding(value0, value1, value2) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
    };
    Binding.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return new Binding(value0, value1, value2);
            };
        };
    };
    return Binding;
})();
var ExprVar = /* #__PURE__ */ (function () {
    function ExprVar(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    ExprVar.create = function (value0) {
        return function (value1) {
            return new ExprVar(value0, value1);
        };
    };
    return ExprVar;
})();
var ExprLit = /* #__PURE__ */ (function () {
    function ExprLit(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    ExprLit.create = function (value0) {
        return function (value1) {
            return new ExprLit(value0, value1);
        };
    };
    return ExprLit;
})();
var ExprConstructor = /* #__PURE__ */ (function () {
    function ExprConstructor(value0, value1, value2, value3) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
        this.value3 = value3;
    };
    ExprConstructor.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return function (value3) {
                    return new ExprConstructor(value0, value1, value2, value3);
                };
            };
        };
    };
    return ExprConstructor;
})();
var ExprAccessor = /* #__PURE__ */ (function () {
    function ExprAccessor(value0, value1, value2) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
    };
    ExprAccessor.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return new ExprAccessor(value0, value1, value2);
            };
        };
    };
    return ExprAccessor;
})();
var ExprUpdate = /* #__PURE__ */ (function () {
    function ExprUpdate(value0, value1, value2) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
    };
    ExprUpdate.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return new ExprUpdate(value0, value1, value2);
            };
        };
    };
    return ExprUpdate;
})();
var ExprAbs = /* #__PURE__ */ (function () {
    function ExprAbs(value0, value1, value2) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
    };
    ExprAbs.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return new ExprAbs(value0, value1, value2);
            };
        };
    };
    return ExprAbs;
})();
var ExprApp = /* #__PURE__ */ (function () {
    function ExprApp(value0, value1, value2) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
    };
    ExprApp.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return new ExprApp(value0, value1, value2);
            };
        };
    };
    return ExprApp;
})();
var ExprCase = /* #__PURE__ */ (function () {
    function ExprCase(value0, value1, value2) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
    };
    ExprCase.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return new ExprCase(value0, value1, value2);
            };
        };
    };
    return ExprCase;
})();
var ExprLet = /* #__PURE__ */ (function () {
    function ExprLet(value0, value1, value2) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
    };
    ExprLet.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return new ExprLet(value0, value1, value2);
            };
        };
    };
    return ExprLet;
})();
var CaseAlternative = /* #__PURE__ */ (function () {
    function CaseAlternative(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    CaseAlternative.create = function (value0) {
        return function (value1) {
            return new CaseAlternative(value0, value1);
        };
    };
    return CaseAlternative;
})();
var Unconditional = /* #__PURE__ */ (function () {
    function Unconditional(value0) {
        this.value0 = value0;
    };
    Unconditional.create = function (value0) {
        return new Unconditional(value0);
    };
    return Unconditional;
})();
var Guarded = /* #__PURE__ */ (function () {
    function Guarded(value0) {
        this.value0 = value0;
    };
    Guarded.create = function (value0) {
        return new Guarded(value0);
    };
    return Guarded;
})();
var Guard = /* #__PURE__ */ (function () {
    function Guard(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    Guard.create = function (value0) {
        return function (value1) {
            return new Guard(value0, value1);
        };
    };
    return Guard;
})();
var Module = function (x) {
    return x;
};
var Ann = function (x) {
    return x;
};
var newtypeModuleName_ = {
    Coercible0: function () {
        return undefined;
    }
};
var newtypeIdent_ = {
    Coercible0: function () {
        return undefined;
    }
};
var functorQualified = {
    map: function (f) {
        return function (m) {
            return new Qualified(m.value0, f(m.value1));
        };
    }
};
var eqProp = function (dictEq) {
    return {
        eq: function (x) {
            return function (y) {
                return x.value0 === y.value0 && Data_Eq.eq(dictEq)(x.value1)(y.value1);
            };
        }
    };
};
var eqLiteral = function (dictEq) {
    var eqArray2 = Data_Eq.eqArray(dictEq);
    var eqArray3 = Data_Eq.eqArray(eqProp(dictEq));
    return {
        eq: function (x) {
            return function (y) {
                if (x instanceof LitInt && y instanceof LitInt) {
                    return x.value0 === y.value0;
                };
                if (x instanceof LitNumber && y instanceof LitNumber) {
                    return x.value0 === y.value0;
                };
                if (x instanceof LitString && y instanceof LitString) {
                    return x.value0 === y.value0;
                };
                if (x instanceof LitChar && y instanceof LitChar) {
                    return x.value0 === y.value0;
                };
                if (x instanceof LitBoolean && y instanceof LitBoolean) {
                    return x.value0 === y.value0;
                };
                if (x instanceof LitArray && y instanceof LitArray) {
                    return Data_Eq.eq(eqArray2)(x.value0)(y.value0);
                };
                if (x instanceof LitRecord && y instanceof LitRecord) {
                    return Data_Eq.eq(eqArray3)(x.value0)(y.value0);
                };
                return false;
            };
        }
    };
};
var unQualified = function (v) {
    return v.value1;
};
var qualifiedModuleName = function (v) {
    return v.value0;
};
var propValue = function (v) {
    return v.value1;
};
var propKey = function (v) {
    return v.value0;
};
var ordProperName = Data_Ord.ordString;
var ordModuleName = Data_Ord.ordString;
var ordMaybe = /* #__PURE__ */ Data_Maybe.ordMaybe(ordModuleName);
var ordIdent = Data_Ord.ordString;
var ordArray1 = /* #__PURE__ */ Data_Ord.ordArray(ordIdent);
var moduleName = function (v) {
    return v.name;
};
var litChildren = function (v) {
    if (v instanceof LitArray) {
        return v.value0;
    };
    if (v instanceof LitRecord) {
        return Data_Functor.map(Data_Functor.functorArray)(propValue)(v.value0);
    };
    return [  ];
};
var isPrimModule = function (v) {
    return v === "Prim" || Data_String_CodeUnits.take(5)(v) === "Prim.";
};
var importName = function (v) {
    return v.value1;
};
var functorProp = {
    map: function (f) {
        return function (m) {
            return new Prop(m.value0, f(m.value1));
        };
    }
};
var functorLiteral = {
    map: function (f) {
        return function (m) {
            if (m instanceof LitInt) {
                return new LitInt(m.value0);
            };
            if (m instanceof LitNumber) {
                return new LitNumber(m.value0);
            };
            if (m instanceof LitString) {
                return new LitString(m.value0);
            };
            if (m instanceof LitChar) {
                return new LitChar(m.value0);
            };
            if (m instanceof LitBoolean) {
                return new LitBoolean(m.value0);
            };
            if (m instanceof LitArray) {
                return new LitArray(Data_Functor.map(Data_Functor.functorArray)(f)(m.value0));
            };
            if (m instanceof LitRecord) {
                return new LitRecord(Data_Functor.map(Data_Functor.functorArray)(Data_Functor.map(functorProp)(f))(m.value0));
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.CoreFn (line 0, column 0 - line 0, column 0): " + [ m.constructor.name ]);
        };
    }
};
var functorImport = {
    map: function (f) {
        return function (m) {
            return new Import(f(m.value0), m.value1);
        };
    }
};
var functorBinder = {
    map: function (f) {
        return function (m) {
            if (m instanceof BinderNull) {
                return new BinderNull(f(m.value0));
            };
            if (m instanceof BinderVar) {
                return new BinderVar(f(m.value0), m.value1);
            };
            if (m instanceof BinderNamed) {
                return new BinderNamed(f(m.value0), m.value1, Data_Functor.map(functorBinder)(f)(m.value2));
            };
            if (m instanceof BinderLit) {
                return new BinderLit(f(m.value0), Data_Functor.map(functorLiteral)(Data_Functor.map(functorBinder)(f))(m.value1));
            };
            if (m instanceof BinderConstructor) {
                return new BinderConstructor(f(m.value0), m.value1, m.value2, Data_Functor.map(Data_Functor.functorArray)(Data_Functor.map(functorBinder)(f))(m.value3));
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.CoreFn (line 0, column 0 - line 0, column 0): " + [ m.constructor.name ]);
        };
    }
};
var functorGuard = {
    map: function (f) {
        return function (m) {
            return new Guard(Data_Functor.map(functorExpr)(f)(m.value0), Data_Functor.map(functorExpr)(f)(m.value1));
        };
    }
};
var functorExpr = {
    map: function (f) {
        return function (m) {
            if (m instanceof ExprVar) {
                return new ExprVar(f(m.value0), m.value1);
            };
            if (m instanceof ExprLit) {
                return new ExprLit(f(m.value0), Data_Functor.map(functorLiteral)(Data_Functor.map(functorExpr)(f))(m.value1));
            };
            if (m instanceof ExprConstructor) {
                return new ExprConstructor(f(m.value0), m.value1, m.value2, m.value3);
            };
            if (m instanceof ExprAccessor) {
                return new ExprAccessor(f(m.value0), Data_Functor.map(functorExpr)(f)(m.value1), m.value2);
            };
            if (m instanceof ExprUpdate) {
                return new ExprUpdate(f(m.value0), Data_Functor.map(functorExpr)(f)(m.value1), Data_Functor.map(Data_Functor.functorArray)(Data_Functor.map(functorProp)(Data_Functor.map(functorExpr)(f)))(m.value2));
            };
            if (m instanceof ExprAbs) {
                return new ExprAbs(f(m.value0), m.value1, Data_Functor.map(functorExpr)(f)(m.value2));
            };
            if (m instanceof ExprApp) {
                return new ExprApp(f(m.value0), Data_Functor.map(functorExpr)(f)(m.value1), Data_Functor.map(functorExpr)(f)(m.value2));
            };
            if (m instanceof ExprCase) {
                return new ExprCase(f(m.value0), Data_Functor.map(Data_Functor.functorArray)(Data_Functor.map(functorExpr)(f))(m.value1), Data_Functor.map(Data_Functor.functorArray)(Data_Functor.map(functorCaseAlternative)(f))(m.value2));
            };
            if (m instanceof ExprLet) {
                return new ExprLet(f(m.value0), Data_Functor.map(Data_Functor.functorArray)(Data_Functor.map(functorBind)(f))(m.value1), Data_Functor.map(functorExpr)(f)(m.value2));
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.CoreFn (line 0, column 0 - line 0, column 0): " + [ m.constructor.name ]);
        };
    }
};
var functorCaseGuard = {
    map: function (f) {
        return function (m) {
            if (m instanceof Unconditional) {
                return new Unconditional(Data_Functor.map(functorExpr)(f)(m.value0));
            };
            if (m instanceof Guarded) {
                return new Guarded(Data_Functor.map(Data_Functor.functorArray)(Data_Functor.map(functorGuard)(f))(m.value0));
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.CoreFn (line 0, column 0 - line 0, column 0): " + [ m.constructor.name ]);
        };
    }
};
var functorCaseAlternative = {
    map: function (f) {
        return function (m) {
            return new CaseAlternative(Data_Functor.map(Data_Functor.functorArray)(Data_Functor.map(functorBinder)(f))(m.value0), Data_Functor.map(functorCaseGuard)(f)(m.value1));
        };
    }
};
var functorBinding = {
    map: function (f) {
        return function (m) {
            return new Binding(f(m.value0), m.value1, Data_Functor.map(functorExpr)(f)(m.value2));
        };
    }
};
var functorBind = {
    map: function (f) {
        return function (m) {
            if (m instanceof NonRec) {
                return new NonRec(Data_Functor.map(functorBinding)(f)(m.value0));
            };
            if (m instanceof Rec) {
                return new Rec(Data_Functor.map(Data_Functor.functorArray)(Data_Functor.map(functorBinding)(f))(m.value0));
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.CoreFn (line 0, column 0 - line 0, column 0): " + [ m.constructor.name ]);
        };
    }
};
var foldableProp = {
    foldl: function (k) {
        return function (a) {
            return function (v) {
                return k(a)(v.value1);
            };
        };
    },
    foldr: function (k) {
        return function (b) {
            return function (v) {
                return k(v.value1)(b);
            };
        };
    },
    foldMap: function (dictMonoid) {
        return function (k) {
            return function (v) {
                return k(v.value1);
            };
        };
    }
};
var traversableProp = {
    traverse: function (dictApplicative) {
        var Functor0 = (dictApplicative.Apply0()).Functor0();
        return function (k) {
            return function (v) {
                return Data_Functor.map(Functor0)(Prop.create(v.value0))(k(v.value1));
            };
        };
    },
    sequence: function (dictApplicative) {
        var Functor0 = (dictApplicative.Apply0()).Functor0();
        return function (v) {
            return Data_Functor.map(Functor0)(Prop.create(v.value0))(v.value1);
        };
    },
    Functor0: function () {
        return functorProp;
    },
    Foldable1: function () {
        return foldableProp;
    }
};
var foldableLiteral = {
    foldl: function (k) {
        return Data_Foldable.foldlDefault(foldableLiteral)(k);
    },
    foldr: function (k) {
        return Data_Foldable.foldrDefault(foldableLiteral)(k);
    },
    foldMap: function (dictMonoid) {
        var mempty = Data_Monoid.mempty(dictMonoid);
        return function (k) {
            return function (v) {
                if (v instanceof LitArray) {
                    return Data_Foldable.foldMap(Data_Foldable.foldableArray)(dictMonoid)(k)(v.value0);
                };
                if (v instanceof LitRecord) {
                    return Data_Foldable.foldMap(Data_Foldable.foldableArray)(dictMonoid)(Data_Foldable.foldMap(foldableProp)(dictMonoid)(k))(v.value0);
                };
                return mempty;
            };
        };
    }
};
var traversableLiteral = {
    traverse: function (dictApplicative) {
        var Functor0 = (dictApplicative.Apply0()).Functor0();
        return function (k) {
            return function (v) {
                if (v instanceof LitArray) {
                    return Data_Functor.map(Functor0)(LitArray.create)(Data_Traversable.traverse(Data_Traversable.traversableArray)(dictApplicative)(k)(v.value0));
                };
                if (v instanceof LitRecord) {
                    return Data_Functor.map(Functor0)(LitRecord.create)(Data_Traversable.traverse(Data_Traversable.traversableArray)(dictApplicative)(Data_Traversable.traverse(traversableProp)(dictApplicative)(k))(v.value0));
                };
                if (v instanceof LitInt) {
                    return Control_Applicative.pure(dictApplicative)(new LitInt(v.value0));
                };
                if (v instanceof LitNumber) {
                    return Control_Applicative.pure(dictApplicative)(new LitNumber(v.value0));
                };
                if (v instanceof LitString) {
                    return Control_Applicative.pure(dictApplicative)(new LitString(v.value0));
                };
                if (v instanceof LitChar) {
                    return Control_Applicative.pure(dictApplicative)(new LitChar(v.value0));
                };
                if (v instanceof LitBoolean) {
                    return Control_Applicative.pure(dictApplicative)(new LitBoolean(v.value0));
                };
                throw new Error("Failed pattern match at PureScript.Backend.Optimizer.CoreFn (line 232, column 16 - line 239, column 40): " + [ v.constructor.name ]);
            };
        };
    },
    sequence: function (dictApplicative) {
        return function (a) {
            return Data_Traversable.sequenceDefault(traversableLiteral)(dictApplicative)(a);
        };
    },
    Functor0: function () {
        return functorLiteral;
    },
    Foldable1: function () {
        return foldableLiteral;
    }
};
var findProp = function (prop) {
    return Data_Array.findMap(function (v) {
        var $446 = prop === v.value0;
        if ($446) {
            return new Data_Maybe.Just(v.value1);
        };
        return Data_Maybe.Nothing.value;
    });
};
var exprAnn = function (v) {
    if (v instanceof ExprVar) {
        return v.value0;
    };
    if (v instanceof ExprLit) {
        return v.value0;
    };
    if (v instanceof ExprConstructor) {
        return v.value0;
    };
    if (v instanceof ExprAccessor) {
        return v.value0;
    };
    if (v instanceof ExprUpdate) {
        return v.value0;
    };
    if (v instanceof ExprAbs) {
        return v.value0;
    };
    if (v instanceof ExprApp) {
        return v.value0;
    };
    if (v instanceof ExprCase) {
        return v.value0;
    };
    if (v instanceof ExprLet) {
        return v.value0;
    };
    throw new Error("Failed pattern match at PureScript.Backend.Optimizer.CoreFn (line 262, column 11 - line 271, column 21): " + [ v.constructor.name ]);
};
var eqProperName = Data_Eq.eqString;
var eqModuleName = Data_Eq.eqString;
var eqMaybe = /* #__PURE__ */ Data_Maybe.eqMaybe(eqModuleName);
var eqQualified = function (dictEq) {
    return {
        eq: function (x) {
            return function (y) {
                return Data_Eq.eq(eqMaybe)(x.value0)(y.value0) && Data_Eq.eq(dictEq)(x.value1)(y.value1);
            };
        }
    };
};
var ordQualified = function (dictOrd) {
    var eqQualified1 = eqQualified(dictOrd.Eq0());
    return {
        compare: function (x) {
            return function (y) {
                var v = Data_Ord.compare(ordMaybe)(x.value0)(y.value0);
                if (v instanceof Data_Ordering.LT) {
                    return Data_Ordering.LT.value;
                };
                if (v instanceof Data_Ordering.GT) {
                    return Data_Ordering.GT.value;
                };
                return Data_Ord.compare(dictOrd)(x.value1)(y.value1);
            };
        },
        Eq0: function () {
            return eqQualified1;
        }
    };
};
var eqIdent = Data_Eq.eqString;
var eqArray1 = /* #__PURE__ */ Data_Eq.eqArray(eqIdent);
var eqReExport = {
    eq: function (x) {
        return function (y) {
            return Data_Eq.eq(eqModuleName)(x.value0)(y.value0) && Data_Eq.eq(eqIdent)(x.value1)(y.value1);
        };
    }
};
var ordReExport = {
    compare: function (x) {
        return function (y) {
            var v = Data_Ord.compare(ordModuleName)(x.value0)(y.value0);
            if (v instanceof Data_Ordering.LT) {
                return Data_Ordering.LT.value;
            };
            if (v instanceof Data_Ordering.GT) {
                return Data_Ordering.GT.value;
            };
            return Data_Ord.compare(ordIdent)(x.value1)(y.value1);
        };
    },
    Eq0: function () {
        return eqReExport;
    }
};
var eqExprType = {
    eq: function (x) {
        return function (y) {
            if (x instanceof Int && y instanceof Int) {
                return true;
            };
            if (x instanceof $$Number && y instanceof $$Number) {
                return true;
            };
            if (x instanceof $$String && y instanceof $$String) {
                return true;
            };
            if (x instanceof Char && y instanceof Char) {
                return true;
            };
            if (x instanceof $$Boolean && y instanceof $$Boolean) {
                return true;
            };
            if (x instanceof Unit && y instanceof Unit) {
                return true;
            };
            if (x instanceof Any && y instanceof Any) {
                return true;
            };
            if (x instanceof TypeLevelString && y instanceof TypeLevelString) {
                return x.value0 === y.value0;
            };
            if (x instanceof $$Array && y instanceof $$Array) {
                return Data_Eq.eq(eqExprType)(x.value0)(y.value0);
            };
            if (x instanceof TypeVar && y instanceof TypeVar) {
                return x.value0 === y.value0;
            };
            if (x instanceof ADT && y instanceof ADT) {
                return x.value0 === y.value0 && Data_Eq.eq(eqArray)(x.value1)(y.value1) && Data_Eq.eq(Data_Eq.eqArray(eqExprType))(x.value2)(y.value2);
            };
            if (x instanceof TypeApp && y instanceof TypeApp) {
                return Data_Eq.eq(eqExprType)(x.value0)(y.value0) && Data_Eq.eq(Data_Eq.eqArray(eqExprType))(x.value1)(y.value1);
            };
            if (x instanceof Func && y instanceof Func) {
                return Data_Eq.eq(Data_Eq.eqArray(eqExprType))(x.value0)(y.value0) && Data_Eq.eq(eqExprType)(x.value1)(y.value1);
            };
            if (x instanceof Row && y instanceof Row) {
                return Data_Eq.eq(Data_Eq.eqArray(eqTuple(eqExprType)))(x.value0)(y.value0) && Data_Eq.eq(Data_Maybe.eqMaybe(eqExprType))(x.value1)(y.value1);
            };
            if (x instanceof Record && y instanceof Record) {
                return Data_Eq.eq(eqExprType)(x.value0)(y.value0);
            };
            if (x instanceof ForAll && y instanceof ForAll) {
                return Data_Eq.eq(eqArray)(x.value0)(y.value0) && Data_Eq.eq(eqExprType)(x.value1)(y.value1);
            };
            if (x instanceof ConstrainedType && y instanceof ConstrainedType) {
                return Data_Eq.eq(Data_Eq.eqArray(eqTuple1(Data_Eq.eqArray(eqExprType))))(x.value0)(y.value0) && Data_Eq.eq(eqExprType)(x.value1)(y.value1);
            };
            return false;
        };
    }
};
var ordExprType = {
    compare: function (x) {
        return function (y) {
            if (x instanceof Int && y instanceof Int) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof Int) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof Int) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof $$Number && y instanceof $$Number) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof $$Number) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof $$Number) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof $$String && y instanceof $$String) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof $$String) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof $$String) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof Char && y instanceof Char) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof Char) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof Char) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof $$Boolean && y instanceof $$Boolean) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof $$Boolean) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof $$Boolean) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof Unit && y instanceof Unit) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof Unit) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof Unit) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof Any && y instanceof Any) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof Any) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof Any) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof TypeLevelString && y instanceof TypeLevelString) {
                return Data_Ord.compare(Data_Ord.ordString)(x.value0)(y.value0);
            };
            if (x instanceof TypeLevelString) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof TypeLevelString) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof $$Array && y instanceof $$Array) {
                return Data_Ord.compare(ordExprType)(x.value0)(y.value0);
            };
            if (x instanceof $$Array) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof $$Array) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof TypeVar && y instanceof TypeVar) {
                return Data_Ord.compare(Data_Ord.ordString)(x.value0)(y.value0);
            };
            if (x instanceof TypeVar) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof TypeVar) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof ADT && y instanceof ADT) {
                var v = Data_Ord.compare(Data_Ord.ordString)(x.value0)(y.value0);
                if (v instanceof Data_Ordering.LT) {
                    return Data_Ordering.LT.value;
                };
                if (v instanceof Data_Ordering.GT) {
                    return Data_Ordering.GT.value;
                };
                var v1 = Data_Ord.compare(ordArray)(x.value1)(y.value1);
                if (v1 instanceof Data_Ordering.LT) {
                    return Data_Ordering.LT.value;
                };
                if (v1 instanceof Data_Ordering.GT) {
                    return Data_Ordering.GT.value;
                };
                return Data_Ord.compare(Data_Ord.ordArray(ordExprType))(x.value2)(y.value2);
            };
            if (x instanceof ADT) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof ADT) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof TypeApp && y instanceof TypeApp) {
                var v = Data_Ord.compare(ordExprType)(x.value0)(y.value0);
                if (v instanceof Data_Ordering.LT) {
                    return Data_Ordering.LT.value;
                };
                if (v instanceof Data_Ordering.GT) {
                    return Data_Ordering.GT.value;
                };
                return Data_Ord.compare(Data_Ord.ordArray(ordExprType))(x.value1)(y.value1);
            };
            if (x instanceof TypeApp) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof TypeApp) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof Func && y instanceof Func) {
                var v = Data_Ord.compare(Data_Ord.ordArray(ordExprType))(x.value0)(y.value0);
                if (v instanceof Data_Ordering.LT) {
                    return Data_Ordering.LT.value;
                };
                if (v instanceof Data_Ordering.GT) {
                    return Data_Ordering.GT.value;
                };
                return Data_Ord.compare(ordExprType)(x.value1)(y.value1);
            };
            if (x instanceof Func) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof Func) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof Row && y instanceof Row) {
                var v = Data_Ord.compare(Data_Ord.ordArray(ordTuple(ordExprType)))(x.value0)(y.value0);
                if (v instanceof Data_Ordering.LT) {
                    return Data_Ordering.LT.value;
                };
                if (v instanceof Data_Ordering.GT) {
                    return Data_Ordering.GT.value;
                };
                return Data_Ord.compare(Data_Maybe.ordMaybe(ordExprType))(x.value1)(y.value1);
            };
            if (x instanceof Row) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof Row) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof Record && y instanceof Record) {
                return Data_Ord.compare(ordExprType)(x.value0)(y.value0);
            };
            if (x instanceof Record) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof Record) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof ForAll && y instanceof ForAll) {
                var v = Data_Ord.compare(ordArray)(x.value0)(y.value0);
                if (v instanceof Data_Ordering.LT) {
                    return Data_Ordering.LT.value;
                };
                if (v instanceof Data_Ordering.GT) {
                    return Data_Ordering.GT.value;
                };
                return Data_Ord.compare(ordExprType)(x.value1)(y.value1);
            };
            if (x instanceof ForAll) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof ForAll) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof ConstrainedType && y instanceof ConstrainedType) {
                var v = Data_Ord.compare(Data_Ord.ordArray(ordTuple1(Data_Ord.ordArray(ordExprType))))(x.value0)(y.value0);
                if (v instanceof Data_Ordering.LT) {
                    return Data_Ordering.LT.value;
                };
                if (v instanceof Data_Ordering.GT) {
                    return Data_Ordering.GT.value;
                };
                return Data_Ord.compare(ordExprType)(x.value1)(y.value1);
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.CoreFn (line 0, column 0 - line 0, column 0): " + [ x.constructor.name, y.constructor.name ]);
        };
    },
    Eq0: function () {
        return eqExprType;
    }
};
var eqConstructorType = {
    eq: function (x) {
        return function (y) {
            if (x instanceof ProductType && y instanceof ProductType) {
                return true;
            };
            if (x instanceof SumType && y instanceof SumType) {
                return true;
            };
            return false;
        };
    }
};
var eqMeta = {
    eq: function (x) {
        return function (y) {
            if (x instanceof IsConstructor && y instanceof IsConstructor) {
                return Data_Eq.eq(eqConstructorType)(x.value0)(y.value0) && Data_Eq.eq(eqArray1)(x.value1)(y.value1);
            };
            if (x instanceof IsNewtype && y instanceof IsNewtype) {
                return true;
            };
            if (x instanceof IsTypeClassConstructor && y instanceof IsTypeClassConstructor) {
                return true;
            };
            if (x instanceof IsForeign && y instanceof IsForeign) {
                return true;
            };
            if (x instanceof IsWhere && y instanceof IsWhere) {
                return true;
            };
            if (x instanceof IsSyntheticApp && y instanceof IsSyntheticApp) {
                return true;
            };
            return false;
        };
    }
};
var ordConstructorType = {
    compare: function (x) {
        return function (y) {
            if (x instanceof ProductType && y instanceof ProductType) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof ProductType) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof ProductType) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof SumType && y instanceof SumType) {
                return Data_Ordering.EQ.value;
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.CoreFn (line 0, column 0 - line 0, column 0): " + [ x.constructor.name, y.constructor.name ]);
        };
    },
    Eq0: function () {
        return eqConstructorType;
    }
};
var ordMeta = {
    compare: function (x) {
        return function (y) {
            if (x instanceof IsConstructor && y instanceof IsConstructor) {
                var v = Data_Ord.compare(ordConstructorType)(x.value0)(y.value0);
                if (v instanceof Data_Ordering.LT) {
                    return Data_Ordering.LT.value;
                };
                if (v instanceof Data_Ordering.GT) {
                    return Data_Ordering.GT.value;
                };
                return Data_Ord.compare(ordArray1)(x.value1)(y.value1);
            };
            if (x instanceof IsConstructor) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof IsConstructor) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof IsNewtype && y instanceof IsNewtype) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof IsNewtype) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof IsNewtype) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof IsTypeClassConstructor && y instanceof IsTypeClassConstructor) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof IsTypeClassConstructor) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof IsTypeClassConstructor) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof IsForeign && y instanceof IsForeign) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof IsForeign) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof IsForeign) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof IsWhere && y instanceof IsWhere) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof IsWhere) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof IsWhere) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof IsSyntheticApp && y instanceof IsSyntheticApp) {
                return Data_Ordering.EQ.value;
            };
            throw new Error("Failed pattern match at PureScript.Backend.Optimizer.CoreFn (line 0, column 0 - line 0, column 0): " + [ x.constructor.name, y.constructor.name ]);
        };
    },
    Eq0: function () {
        return eqMeta;
    }
};
var emptySpan = {
    path: "<internal>",
    start: zero,
    end: zero
};
var emptyAnn = /* #__PURE__ */ (function () {
    return {
        span: emptySpan,
        meta: Data_Maybe.Nothing.value,
        type: Data_Maybe.Nothing.value
    };
})();
export {
    Ident,
    ModuleName,
    ProperName,
    Qualified,
    unQualified,
    qualifiedModuleName,
    Ann,
    Int,
    $$Number as Number,
    $$String as String,
    Char,
    $$Boolean as Boolean,
    Unit,
    Any,
    TypeLevelString,
    $$Array as Array,
    TypeVar,
    ADT,
    TypeApp,
    Func,
    Row,
    Record,
    ForAll,
    ConstrainedType,
    IsConstructor,
    IsNewtype,
    IsTypeClassConstructor,
    IsForeign,
    IsWhere,
    IsSyntheticApp,
    ProductType,
    SumType,
    LineComment,
    BlockComment,
    Module,
    moduleName,
    Import,
    importName,
    ReExport,
    NonRec,
    Rec,
    Binding,
    ExprVar,
    ExprLit,
    ExprConstructor,
    ExprAccessor,
    ExprUpdate,
    ExprAbs,
    ExprApp,
    ExprCase,
    ExprLet,
    CaseAlternative,
    Unconditional,
    Guarded,
    Guard,
    Prop,
    propKey,
    propValue,
    findProp,
    LitInt,
    LitNumber,
    LitString,
    LitChar,
    LitBoolean,
    LitArray,
    LitRecord,
    BinderNull,
    BinderVar,
    BinderNamed,
    BinderLit,
    BinderConstructor,
    emptyAnn,
    emptySpan,
    exprAnn,
    litChildren,
    isPrimModule,
    eqIdent,
    ordIdent,
    newtypeIdent_,
    eqModuleName,
    ordModuleName,
    newtypeModuleName_,
    eqProperName,
    ordProperName,
    eqQualified,
    ordQualified,
    functorQualified,
    eqExprType,
    ordExprType,
    eqMeta,
    ordMeta,
    eqConstructorType,
    ordConstructorType,
    functorImport,
    eqReExport,
    ordReExport,
    functorBind,
    functorBinding,
    functorExpr,
    functorCaseAlternative,
    functorCaseGuard,
    functorGuard,
    eqProp,
    functorProp,
    foldableProp,
    traversableProp,
    eqLiteral,
    functorLiteral,
    foldableLiteral,
    traversableLiteral,
    functorBinder
};
