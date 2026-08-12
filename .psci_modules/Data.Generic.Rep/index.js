import * as Data_Show from "../Data.Show/index.js";
import * as Data_Symbol from "../Data.Symbol/index.js";
import * as Type_Proxy from "../Type.Proxy/index.js";

// | A representation for types with multiple constructors.
var Inl = /* #__PURE__ */ (function () {
    function Inl(value0) {
        this.value0 = value0;
    };
    Inl.create = function (value0) {
        return new Inl(value0);
    };
    return Inl;
})();

// | A representation for types with multiple constructors.
var Inr = /* #__PURE__ */ (function () {
    function Inr(value0) {
        this.value0 = value0;
    };
    Inr.create = function (value0) {
        return new Inr(value0);
    };
    return Inr;
})();

// | A representation for constructors with multiple fields.
var Product = /* #__PURE__ */ (function () {
    function Product(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    Product.create = function (value0) {
        return function (value1) {
            return new Product(value0, value1);
        };
    };
    return Product;
})();

// | A representation for types with no constructors.
var NoConstructors = function (x) {
    return x;
};

// | A representation for constructors with no arguments.
var NoArguments = /* #__PURE__ */ (function () {
    function NoArguments() {

    };
    NoArguments.value = new NoArguments();
    return NoArguments;
})();

// | A representation for constructors which includes the data constructor name
// | as a type-level string.
var Constructor = function (x) {
    return x;
};

// | A representation for an argument in a data constructor.
var Argument = function (x) {
    return x;
};
var to = function (dict) {
    return dict.to;
};
var showSum = function (dictShow) {
    return function (dictShow1) {
        return {
            show: function (v) {
                if (v instanceof Inl) {
                    return "(Inl " + (Data_Show.show(dictShow)(v.value0) + ")");
                };
                if (v instanceof Inr) {
                    return "(Inr " + (Data_Show.show(dictShow1)(v.value0) + ")");
                };
                throw new Error("Failed pattern match at Data.Generic.Rep (line 32, column 1 - line 34, column 42): " + [ v.constructor.name ]);
            }
        };
    };
};
var showProduct = function (dictShow) {
    return function (dictShow1) {
        return {
            show: function (v) {
                return "(Product " + (Data_Show.show(dictShow)(v.value0) + (" " + (Data_Show.show(dictShow1)(v.value1) + ")")));
            }
        };
    };
};
var showNoArguments = {
    show: function (v) {
        return "NoArguments";
    }
};
var showConstructor = function (dictIsSymbol) {
    return function (dictShow) {
        return {
            show: function (v) {
                return "(Constructor @" + (Data_Show.show(Data_Show.showString)(Data_Symbol.reflectSymbol(dictIsSymbol)(Type_Proxy["Proxy"].value)) + (" " + (Data_Show.show(dictShow)(v) + ")")));
            }
        };
    };
};
var showArgument = function (dictShow) {
    return {
        show: function (v) {
            return "(Argument " + (Data_Show.show(dictShow)(v) + ")");
        }
    };
};
var repOf = function (dictGeneric) {
    return function (v) {
        return Type_Proxy["Proxy"].value;
    };
};
var from = function (dict) {
    return dict.from;
};
export {
    to,
    from,
    repOf,
    NoArguments,
    Inl,
    Inr,
    Product,
    Constructor,
    Argument,
    showNoArguments,
    showSum,
    showProduct,
    showConstructor,
    showArgument
};
