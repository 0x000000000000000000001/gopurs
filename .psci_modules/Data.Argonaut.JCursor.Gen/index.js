import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Lazy from "../Control.Lazy/index.js";
import * as Control_Monad_Gen from "../Control.Monad.Gen/index.js";
import * as Control_Monad_Gen_Class from "../Control.Monad.Gen.Class/index.js";
import * as Data_Argonaut_JCursor from "../Data.Argonaut.JCursor/index.js";
import * as Data_Boolean from "../Data.Boolean/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_String_Gen from "../Data.String.Gen/index.js";
var genJCursor = function (dictMonadGen) {
    var Monad0 = dictMonadGen.Monad0();
    var Bind1 = Monad0.Bind1();
    var Apply0 = Bind1.Apply0();
    var Functor0 = (Bind1.Apply0()).Functor0();
    var Applicative0 = Monad0.Applicative0();
    return function (dictMonadRec) {
        var genUnicodeString = Data_String_Gen.genUnicodeString(dictMonadRec)(dictMonadGen);
        return function (dictLazy) {
            var genIndex = Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(Data_Argonaut_JCursor.JIndex.create)(Control_Monad_Gen_Class.chooseInt(dictMonadGen)(0)(1000)))(Control_Lazy.defer(dictLazy)(function (v) {
                return genJCursor(dictMonadGen)(dictMonadRec)(dictLazy);
            }));
            var genField = Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(Data_Argonaut_JCursor.JField.create)(genUnicodeString))(Control_Lazy.defer(dictLazy)(function (v) {
                return genJCursor(dictMonadGen)(dictMonadRec)(dictLazy);
            }));
            var genJCursor$prime = function (size) {
                if (size > 0) {
                    return Control_Monad_Gen_Class.resize(dictMonadGen)(function (v) {
                        return v - 1 | 0;
                    })(Control_Monad_Gen.choose(dictMonadGen)(genField)(genIndex));
                };
                if (Data_Boolean.otherwise) {
                    return Control_Applicative.pure(Applicative0)(Data_Argonaut_JCursor.JCursorTop.value);
                };
                throw new Error("Failed pattern match at Data.Argonaut.JCursor.Gen (line 16, column 3 - line 18, column 34): " + [ size.constructor.name ]);
            };
            return Control_Monad_Gen_Class.resize(dictMonadGen)(Data_Ord.min(Data_Ord.ordInt)(10))(Control_Monad_Gen_Class.sized(dictMonadGen)(genJCursor$prime));
        };
    };
};
export {
    genJCursor
};
