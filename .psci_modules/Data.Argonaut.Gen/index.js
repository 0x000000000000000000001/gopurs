import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Lazy from "../Control.Lazy/index.js";
import * as Control_Monad_Gen from "../Control.Monad.Gen/index.js";
import * as Control_Monad_Gen_Class from "../Control.Monad.Gen.Class/index.js";
import * as Data_Argonaut_Core from "../Data.Argonaut.Core/index.js";
import * as Data_Array from "../Data.Array/index.js";
import * as Data_Boolean from "../Data.Boolean/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_NonEmpty from "../Data.NonEmpty/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_String_Gen from "../Data.String.Gen/index.js";
import * as Data_Unfoldable from "../Data.Unfoldable/index.js";
import * as Foreign_Object from "../Foreign.Object/index.js";
var foldable1NonEmpty = /* #__PURE__ */ Data_NonEmpty.foldable1NonEmpty(Data_Foldable.foldableArray);

// | A generator for `Json` values. Especially useful for writing property-based
// | tests.
var genJson = function (dictMonadGen) {
    var Monad0 = dictMonadGen.Monad0();
    var Functor0 = ((Monad0.Bind1()).Apply0()).Functor0();
    var chooseBool = Control_Monad_Gen_Class.chooseBool(dictMonadGen);
    var oneOf = Control_Monad_Gen.oneOf(dictMonadGen)(foldable1NonEmpty);
    var Applicative0 = Monad0.Applicative0();
    var Bind1 = Monad0.Bind1();
    var pure = Control_Applicative.pure(Monad0.Applicative0());
    var Monad01 = dictMonadGen.Monad0();
    return function (dictMonadRec) {
        var genUnicodeString = Data_String_Gen.genUnicodeString(dictMonadRec)(dictMonadGen);
        return function (dictLazy) {
            var genJString = Data_Functor.map(Functor0)(Data_Argonaut_Core.fromString)(genUnicodeString);
            var genJNumber = Data_Functor.map(Functor0)(Data_Argonaut_Core.fromNumber)(Control_Monad_Gen_Class.chooseFloat(dictMonadGen)(-1000000.0)(1000000.0));
            var genJBoolean = Data_Functor.map(Functor0)(Data_Argonaut_Core.fromBoolean)(chooseBool);
            var genLeaf = oneOf(new Data_NonEmpty.NonEmpty(Control_Applicative.pure(Applicative0)(Data_Argonaut_Core.jsonNull), [ genJBoolean, genJNumber, genJString ]));
            var genJArray = Data_Functor.map(Functor0)(Data_Argonaut_Core.fromArray)(Control_Monad_Gen.unfoldable(dictMonadRec)(dictMonadGen)(Data_Unfoldable.unfoldableArray)(Control_Lazy.defer(dictLazy)(function (v) {
                return genJson(dictMonadGen)(dictMonadRec)(dictLazy);
            })));
            var extendJObj = function (obj) {
                return function (k) {
                    return Control_Bind.bind(Bind1)(genJson(dictMonadGen)(dictMonadRec)(dictLazy))(function (v) {
                        return pure(Data_Argonaut_Core.caseJsonObject(Data_Argonaut_Core.jsonSingletonObject(k)(v))((function () {
                            var $21 = Foreign_Object.insert(k)(v);
                            return function ($22) {
                                return Data_Argonaut_Core.fromObject($21($22));
                            };
                        })())(obj));
                    });
                };
            };
            var genJObject = Control_Bind.bindFlipped(Bind1)(Data_Array.foldM(Monad01)(extendJObj)(Data_Argonaut_Core.jsonEmptyObject))(Control_Monad_Gen.unfoldable(dictMonadRec)(dictMonadGen)(Data_Unfoldable.unfoldableArray)(genUnicodeString));
            var genJson$prime = function (size) {
                if (size > 1) {
                    return Control_Monad_Gen_Class.resize(dictMonadGen)(function (v) {
                        return v - 1 | 0;
                    })(Control_Monad_Gen.choose(dictMonadGen)(genJArray)(genJObject));
                };
                if (Data_Boolean.otherwise) {
                    return genLeaf;
                };
                throw new Error("Failed pattern match at Data.Argonaut.Gen (line 20, column 3 - line 20, column 30): " + [ size.constructor.name ]);
            };
            return Control_Monad_Gen_Class.resize(dictMonadGen)(Data_Ord.min(Data_Ord.ordInt)(5))(Control_Monad_Gen_Class.sized(dictMonadGen)(genJson$prime));
        };
    };
};
export {
    genJson
};
