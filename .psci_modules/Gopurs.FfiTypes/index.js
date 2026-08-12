import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Data_Argonaut_Core from "../Data.Argonaut.Core/index.js";
import * as Data_Argonaut_Decode_Class from "../Data.Argonaut.Decode.Class/index.js";
import * as Data_Argonaut_Decode_Combinators from "../Data.Argonaut.Decode.Combinators/index.js";
import * as Data_Argonaut_Decode_Error from "../Data.Argonaut.Decode.Error/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Ordering from "../Data.Ordering/index.js";
import * as Data_Show from "../Data.Show/index.js";
var TNamed = /* #__PURE__ */ (function () {
    function TNamed(value0) {
        this.value0 = value0;
    };
    TNamed.create = function (value0) {
        return new TNamed(value0);
    };
    return TNamed;
})();
var TFunc = /* #__PURE__ */ (function () {
    function TFunc(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    TFunc.create = function (value0) {
        return function (value1) {
            return new TFunc(value0, value1);
        };
    };
    return TFunc;
})();
var TArray = /* #__PURE__ */ (function () {
    function TArray(value0) {
        this.value0 = value0;
    };
    TArray.create = function (value0) {
        return new TArray(value0);
    };
    return TArray;
})();
var TMap = /* #__PURE__ */ (function () {
    function TMap(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    TMap.create = function (value0) {
        return function (value1) {
            return new TMap(value0, value1);
        };
    };
    return TMap;
})();
var TUnknown = /* #__PURE__ */ (function () {
    function TUnknown(value0) {
        this.value0 = value0;
    };
    TUnknown.create = function (value0) {
        return new TUnknown(value0);
    };
    return TUnknown;
})();
var showTypeNode = {
    show: function (v) {
        if (v instanceof TNamed) {
            return "(TNamed " + (Data_Show.show(Data_Show.showString)(v.value0) + ")");
        };
        if (v instanceof TFunc) {
            return "(TFunc " + (Data_Show.show(Data_Show.showArray(showTypeNode))(v.value0) + (" " + (Data_Show.show(Data_Maybe.showMaybe(showTypeNode))(v.value1) + ")")));
        };
        if (v instanceof TArray) {
            return "(TArray " + (Data_Show.show(showTypeNode)(v.value0) + ")");
        };
        if (v instanceof TMap) {
            return "(TMap " + (Data_Show.show(showTypeNode)(v.value0) + (" " + (Data_Show.show(showTypeNode)(v.value1) + ")")));
        };
        if (v instanceof TUnknown) {
            return "(TUnknown " + (Data_Show.show(Data_Show.showString)(v.value0) + ")");
        };
        throw new Error("Failed pattern match at Gopurs.FfiTypes (line 27, column 1 - line 32, column 52): " + [ v.constructor.name ]);
    }
};
var eqTypeNode = {
    eq: function (x) {
        return function (y) {
            if (x instanceof TNamed && y instanceof TNamed) {
                return x.value0 === y.value0;
            };
            if (x instanceof TFunc && y instanceof TFunc) {
                return Data_Eq.eq(Data_Eq.eqArray(eqTypeNode))(x.value0)(y.value0) && Data_Eq.eq(Data_Maybe.eqMaybe(eqTypeNode))(x.value1)(y.value1);
            };
            if (x instanceof TArray && y instanceof TArray) {
                return Data_Eq.eq(eqTypeNode)(x.value0)(y.value0);
            };
            if (x instanceof TMap && y instanceof TMap) {
                return Data_Eq.eq(eqTypeNode)(x.value0)(y.value0) && Data_Eq.eq(eqTypeNode)(x.value1)(y.value1);
            };
            if (x instanceof TUnknown && y instanceof TUnknown) {
                return x.value0 === y.value0;
            };
            return false;
        };
    }
};
var ordTypeNode = {
    compare: function (x) {
        return function (y) {
            if (x instanceof TNamed && y instanceof TNamed) {
                return Data_Ord.compare(Data_Ord.ordString)(x.value0)(y.value0);
            };
            if (x instanceof TNamed) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof TNamed) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof TFunc && y instanceof TFunc) {
                var v = Data_Ord.compare(Data_Ord.ordArray(ordTypeNode))(x.value0)(y.value0);
                if (v instanceof Data_Ordering.LT) {
                    return Data_Ordering.LT.value;
                };
                if (v instanceof Data_Ordering.GT) {
                    return Data_Ordering.GT.value;
                };
                return Data_Ord.compare(Data_Maybe.ordMaybe(ordTypeNode))(x.value1)(y.value1);
            };
            if (x instanceof TFunc) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof TFunc) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof TArray && y instanceof TArray) {
                return Data_Ord.compare(ordTypeNode)(x.value0)(y.value0);
            };
            if (x instanceof TArray) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof TArray) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof TMap && y instanceof TMap) {
                var v = Data_Ord.compare(ordTypeNode)(x.value0)(y.value0);
                if (v instanceof Data_Ordering.LT) {
                    return Data_Ordering.LT.value;
                };
                if (v instanceof Data_Ordering.GT) {
                    return Data_Ordering.GT.value;
                };
                return Data_Ord.compare(ordTypeNode)(x.value1)(y.value1);
            };
            if (x instanceof TMap) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof TMap) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof TUnknown && y instanceof TUnknown) {
                return Data_Ord.compare(Data_Ord.ordString)(x.value0)(y.value0);
            };
            throw new Error("Failed pattern match at Gopurs.FfiTypes (line 0, column 0 - line 0, column 0): " + [ x.constructor.name, y.constructor.name ]);
        };
    },
    Eq0: function () {
        return eqTypeNode;
    }
};
var decodeJsonTypeNode = {
    decodeJson: function (json) {
        return Control_Bind.bind(Data_Either.bindEither)(Data_Either.note(new Data_Argonaut_Decode_Error.TypeMismatch("Object"))(Data_Argonaut_Core.toObject(json)))(function (obj) {
            return Control_Bind.bind(Data_Either.bindEither)(Data_Argonaut_Decode_Combinators.getField(Data_Argonaut_Decode_Class.decodeJsonString)(obj)("type"))(function (typStr) {
                if (typStr === "Named") {
                    return Data_Functor.map(Data_Either.functorEither)(TNamed.create)(Data_Argonaut_Decode_Combinators.getField(Data_Argonaut_Decode_Class.decodeJsonString)(obj)("name"));
                };
                if (typStr === "Func") {
                    return Control_Bind.bind(Data_Either.bindEither)(Data_Functor.mapFlipped(Data_Either.functorEither)(Data_Argonaut_Decode_Combinators["getFieldOptional$prime"](Data_Argonaut_Decode_Class.decodeArray(decodeJsonTypeNode))(obj)("args"))(Data_Maybe.fromMaybe([  ])))(function (args) {
                        return Control_Bind.bind(Data_Either.bindEither)(Data_Argonaut_Decode_Combinators["getFieldOptional$prime"](decodeJsonTypeNode)(obj)("ret"))(function (ret) {
                            return Control_Applicative.pure(Data_Either.applicativeEither)(new TFunc(args, ret));
                        });
                    });
                };
                if (typStr === "Array") {
                    return Data_Functor.map(Data_Either.functorEither)(TArray.create)(Data_Argonaut_Decode_Combinators.getField(decodeJsonTypeNode)(obj)("elem"));
                };
                if (typStr === "Map") {
                    return Control_Bind.bind(Data_Either.bindEither)(Data_Argonaut_Decode_Combinators.getField(decodeJsonTypeNode)(obj)("key"))(function (k) {
                        return Control_Bind.bind(Data_Either.bindEither)(Data_Argonaut_Decode_Combinators.getField(decodeJsonTypeNode)(obj)("val"))(function (v) {
                            return Control_Applicative.pure(Data_Either.applicativeEither)(new TMap(k, v));
                        });
                    });
                };
                return Data_Functor.map(Data_Either.functorEither)(TUnknown.create)(Data_Argonaut_Decode_Combinators.getField(Data_Argonaut_Decode_Class.decodeJsonString)(obj)("name"));
            });
        });
    }
};
export {
    TNamed,
    TFunc,
    TArray,
    TMap,
    TUnknown,
    eqTypeNode,
    ordTypeNode,
    showTypeNode,
    decodeJsonTypeNode
};
