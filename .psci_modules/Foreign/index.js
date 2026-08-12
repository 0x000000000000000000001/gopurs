// | This module defines types and functions for working with _foreign_
// | data.
// |
// | `ExceptT (NonEmptyList ForeignError) m` is used in this library
// | to encode possible failures when dealing with foreign data.
// |
// | The `Alt` instance for `ExceptT` allows us to accumulate errors,
// | unlike `Either`, which preserves only the last error.
import * as $foreign from "./foreign.js";
import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Monad_Error_Class from "../Control.Monad.Error.Class/index.js";
import * as Control_Monad_Except_Trans from "../Control.Monad.Except.Trans/index.js";
import * as Data_Boolean from "../Data.Boolean/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Int from "../Data.Int/index.js";
import * as Data_List_NonEmpty from "../Data.List.NonEmpty/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Ordering from "../Data.Ordering/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Data_String_CodeUnits from "../Data.String.CodeUnits/index.js";
import * as Unsafe_Coerce from "../Unsafe.Coerce/index.js";
var pure = /* #__PURE__ */ Control_Applicative.pure(Data_Either.applicativeEither);
var pure1 = /* #__PURE__ */ Control_Applicative.pure(Data_Either.applicativeEither);

// | A type for foreign type errors
var ForeignError = /* #__PURE__ */ (function () {
    function ForeignError(value0) {
        this.value0 = value0;
    };
    ForeignError.create = function (value0) {
        return new ForeignError(value0);
    };
    return ForeignError;
})();

// | A type for foreign type errors
var TypeMismatch = /* #__PURE__ */ (function () {
    function TypeMismatch(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    TypeMismatch.create = function (value0) {
        return function (value1) {
            return new TypeMismatch(value0, value1);
        };
    };
    return TypeMismatch;
})();

// | A type for foreign type errors
var ErrorAtIndex = /* #__PURE__ */ (function () {
    function ErrorAtIndex(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    ErrorAtIndex.create = function (value0) {
        return function (value1) {
            return new ErrorAtIndex(value0, value1);
        };
    };
    return ErrorAtIndex;
})();

// | A type for foreign type errors
var ErrorAtProperty = /* #__PURE__ */ (function () {
    function ErrorAtProperty(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    ErrorAtProperty.create = function (value0) {
        return function (value1) {
            return new ErrorAtProperty(value0, value1);
        };
    };
    return ErrorAtProperty;
})();

// | Coerce any value to the a `Foreign` value.
// |
// | This is considered unsafe as it's only intended to be used on primitive
// | JavaScript types, rather than PureScript types. Exporting PureScript values
// | via the FFI can be dangerous as they can be mutated by code outside the
// | PureScript program, resulting in difficult to diagnose problems elsewhere.
var unsafeToForeign = Unsafe_Coerce.unsafeCoerce;

// | Unsafely coerce a `Foreign` value.
var unsafeFromForeign = Unsafe_Coerce.unsafeCoerce;
var showForeignError = {
    show: function (v) {
        if (v instanceof ForeignError) {
            return "(ForeignError " + (Data_Show.show(Data_Show.showString)(v.value0) + ")");
        };
        if (v instanceof ErrorAtIndex) {
            return "(ErrorAtIndex " + (Data_Show.show(Data_Show.showInt)(v.value0) + (" " + (Data_Show.show(showForeignError)(v.value1) + ")")));
        };
        if (v instanceof ErrorAtProperty) {
            return "(ErrorAtProperty " + (Data_Show.show(Data_Show.showString)(v.value0) + (" " + (Data_Show.show(showForeignError)(v.value1) + ")")));
        };
        if (v instanceof TypeMismatch) {
            return "(TypeMismatch " + (Data_Show.show(Data_Show.showString)(v.value0) + (" " + (Data_Show.show(Data_Show.showString)(v.value1) + ")")));
        };
        throw new Error("Failed pattern match at Foreign (line 69, column 1 - line 73, column 89): " + [ v.constructor.name ]);
    }
};
var renderForeignError = function (v) {
    if (v instanceof ForeignError) {
        return v.value0;
    };
    if (v instanceof ErrorAtIndex) {
        return "Error at array index " + (Data_Show.show(Data_Show.showInt)(v.value0) + (": " + renderForeignError(v.value1)));
    };
    if (v instanceof ErrorAtProperty) {
        return "Error at property " + (Data_Show.show(Data_Show.showString)(v.value0) + (": " + renderForeignError(v.value1)));
    };
    if (v instanceof TypeMismatch) {
        return "Type mismatch: expected " + (v.value0 + (", found " + v.value1));
    };
    throw new Error("Failed pattern match at Foreign (line 78, column 1 - line 78, column 45): " + [ v.constructor.name ]);
};
var readUndefined = function (dictMonad) {
    var applicativeExceptT = Control_Monad_Except_Trans.applicativeExceptT(dictMonad);
    return function (value) {
        if ($foreign.isUndefined(value)) {
            return Control_Applicative.pure(applicativeExceptT)(Data_Maybe.Nothing.value);
        };
        if (Data_Boolean.otherwise) {
            return Control_Applicative.pure(applicativeExceptT)(new Data_Maybe.Just(value));
        };
        throw new Error("Failed pattern match at Foreign (line 174, column 1 - line 174, column 103): " + [ value.constructor.name ]);
    };
};
var readNullOrUndefined = function (dictMonad) {
    var applicativeExceptT = Control_Monad_Except_Trans.applicativeExceptT(dictMonad);
    return function (value) {
        if ($foreign.isNull(value) || $foreign.isUndefined(value)) {
            return Control_Applicative.pure(applicativeExceptT)(Data_Maybe.Nothing.value);
        };
        if (Data_Boolean.otherwise) {
            return Control_Applicative.pure(applicativeExceptT)(new Data_Maybe.Just(value));
        };
        throw new Error("Failed pattern match at Foreign (line 179, column 1 - line 179, column 109): " + [ value.constructor.name ]);
    };
};
var readNull = function (dictMonad) {
    var applicativeExceptT = Control_Monad_Except_Trans.applicativeExceptT(dictMonad);
    return function (value) {
        if ($foreign.isNull(value)) {
            return Control_Applicative.pure(applicativeExceptT)(Data_Maybe.Nothing.value);
        };
        if (Data_Boolean.otherwise) {
            return Control_Applicative.pure(applicativeExceptT)(new Data_Maybe.Just(value));
        };
        throw new Error("Failed pattern match at Foreign (line 169, column 1 - line 169, column 98): " + [ value.constructor.name ]);
    };
};

// | Throws a failure error in `ExceptT (NonEmptyList ForeignError) m`.
var fail = function (dictMonad) {
    var $135 = Control_Monad_Error_Class.throwError(Control_Monad_Except_Trans.monadThrowExceptT(dictMonad));
    return function ($136) {
        return $135(Data_List_NonEmpty.singleton($136));
    };
};

// | Attempt to coerce a foreign value to an array.
var readArray = function (dictMonad) {
    var pure2 = Control_Applicative.pure(Control_Monad_Except_Trans.applicativeExceptT(dictMonad));
    var fail1 = fail(dictMonad);
    return function (value) {
        if ($foreign.isArray(value)) {
            return pure2(unsafeFromForeign(value));
        };
        if (Data_Boolean.otherwise) {
            return fail1(new TypeMismatch("array", $foreign.tagOf(value)));
        };
        throw new Error("Failed pattern match at Foreign (line 164, column 1 - line 164, column 99): " + [ value.constructor.name ]);
    };
};

// | Unsafely coerce a `Foreign` value when the value has a particular `tagOf`
// | value.
var unsafeReadTagged = function (dictMonad) {
    var applicativeExceptT = Control_Monad_Except_Trans.applicativeExceptT(dictMonad);
    var fail1 = fail(dictMonad);
    return function (tag) {
        return function (value) {
            if ($foreign.tagOf(value) === tag) {
                return Control_Applicative.pure(applicativeExceptT)(unsafeFromForeign(value));
            };
            if (Data_Boolean.otherwise) {
                return fail1(new TypeMismatch(tag, $foreign.tagOf(value)));
            };
            throw new Error("Failed pattern match at Foreign (line 123, column 1 - line 123, column 104): " + [ tag.constructor.name, value.constructor.name ]);
        };
    };
};

// | Attempt to coerce a foreign value to a `Boolean`.
var readBoolean = function (dictMonad) {
    return unsafeReadTagged(dictMonad)("Boolean");
};

// | Attempt to coerce a foreign value to a `Number`.
var readNumber = function (dictMonad) {
    return unsafeReadTagged(dictMonad)("Number");
};

// | Attempt to coerce a foreign value to an `Int`.
var readInt = function (dictMonad) {
    var map = Data_Functor.map(((dictMonad.Bind1()).Apply0()).Functor0());
    return function (value) {
        var error = new Data_Either.Left(Data_List_NonEmpty.singleton(new TypeMismatch("Int", $foreign.tagOf(value))));
        var fromNumber = (function () {
            var $137 = Data_Maybe.maybe(error)(pure);
            return function ($138) {
                return $137(Data_Int.fromNumber($138));
            };
        })();
        return Control_Monad_Except_Trans.mapExceptT(map(Data_Either.either(Data_Function["const"](error))(fromNumber)))(readNumber(dictMonad)(value));
    };
};

// | Attempt to coerce a foreign value to a `String`.
var readString = function (dictMonad) {
    return unsafeReadTagged(dictMonad)("String");
};

// | Attempt to coerce a foreign value to a `Char`.
var readChar = function (dictMonad) {
    var map = Data_Functor.map(((dictMonad.Bind1()).Apply0()).Functor0());
    return function (value) {
        var error = new Data_Either.Left(Data_List_NonEmpty.singleton(new TypeMismatch("Char", $foreign.tagOf(value))));
        var fromString = (function () {
            var $139 = Data_Maybe.maybe(error)(pure1);
            return function ($140) {
                return $139(Data_String_CodeUnits.toChar($140));
            };
        })();
        return Control_Monad_Except_Trans.mapExceptT(map(Data_Either.either(Data_Function["const"](error))(fromString)))(readString(dictMonad)(value));
    };
};
var eqForeignError = {
    eq: function (x) {
        return function (y) {
            if (x instanceof ForeignError && y instanceof ForeignError) {
                return x.value0 === y.value0;
            };
            if (x instanceof TypeMismatch && y instanceof TypeMismatch) {
                return x.value0 === y.value0 && x.value1 === y.value1;
            };
            if (x instanceof ErrorAtIndex && y instanceof ErrorAtIndex) {
                return x.value0 === y.value0 && Data_Eq.eq(eqForeignError)(x.value1)(y.value1);
            };
            if (x instanceof ErrorAtProperty && y instanceof ErrorAtProperty) {
                return x.value0 === y.value0 && Data_Eq.eq(eqForeignError)(x.value1)(y.value1);
            };
            return false;
        };
    }
};
var ordForeignError = {
    compare: function (x) {
        return function (y) {
            if (x instanceof ForeignError && y instanceof ForeignError) {
                return Data_Ord.compare(Data_Ord.ordString)(x.value0)(y.value0);
            };
            if (x instanceof ForeignError) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof ForeignError) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof TypeMismatch && y instanceof TypeMismatch) {
                var v = Data_Ord.compare(Data_Ord.ordString)(x.value0)(y.value0);
                if (v instanceof Data_Ordering.LT) {
                    return Data_Ordering.LT.value;
                };
                if (v instanceof Data_Ordering.GT) {
                    return Data_Ordering.GT.value;
                };
                return Data_Ord.compare(Data_Ord.ordString)(x.value1)(y.value1);
            };
            if (x instanceof TypeMismatch) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof TypeMismatch) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof ErrorAtIndex && y instanceof ErrorAtIndex) {
                var v = Data_Ord.compare(Data_Ord.ordInt)(x.value0)(y.value0);
                if (v instanceof Data_Ordering.LT) {
                    return Data_Ordering.LT.value;
                };
                if (v instanceof Data_Ordering.GT) {
                    return Data_Ordering.GT.value;
                };
                return Data_Ord.compare(ordForeignError)(x.value1)(y.value1);
            };
            if (x instanceof ErrorAtIndex) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof ErrorAtIndex) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof ErrorAtProperty && y instanceof ErrorAtProperty) {
                var v = Data_Ord.compare(Data_Ord.ordString)(x.value0)(y.value0);
                if (v instanceof Data_Ordering.LT) {
                    return Data_Ordering.LT.value;
                };
                if (v instanceof Data_Ordering.GT) {
                    return Data_Ordering.GT.value;
                };
                return Data_Ord.compare(ordForeignError)(x.value1)(y.value1);
            };
            throw new Error("Failed pattern match at Foreign (line 0, column 0 - line 0, column 0): " + [ x.constructor.name, y.constructor.name ]);
        };
    },
    Eq0: function () {
        return eqForeignError;
    }
};
export {
    typeOf,
    tagOf,
    isNull,
    isUndefined,
    isArray
} from "./foreign.js";
export {
    ForeignError,
    TypeMismatch,
    ErrorAtIndex,
    ErrorAtProperty,
    renderForeignError,
    unsafeToForeign,
    unsafeFromForeign,
    unsafeReadTagged,
    readString,
    readChar,
    readBoolean,
    readNumber,
    readInt,
    readArray,
    readNull,
    readUndefined,
    readNullOrUndefined,
    fail,
    eqForeignError,
    ordForeignError,
    showForeignError
};
