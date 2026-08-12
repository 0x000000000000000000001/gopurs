import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Monad_Gen from "../Control.Monad.Gen/index.js";
import * as Control_Monad_Gen_Class from "../Control.Monad.Gen.Class/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_List_Types from "../Data.List.Types/index.js";
import * as Data_Map_Internal from "../Data.Map.Internal/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";

// | Generates a `Map` using the specified key and value generators.
var genMap = function (dictMonadRec) {
    return function (dictMonadGen) {
        var Monad0 = dictMonadGen.Monad0();
        var Bind1 = Monad0.Bind1();
        var Bind11 = Monad0.Bind1();
        var Functor0 = (Bind11.Apply0()).Functor0();
        var Apply0 = Bind11.Apply0();
        return function (dictOrd) {
            var fromFoldable = Data_Map_Internal.fromFoldable(dictOrd)(Data_List_Types.foldableList);
            return function (genKey) {
                return function (genValue) {
                    return Control_Monad_Gen_Class.sized(dictMonadGen)(function (size) {
                        return Control_Bind.bind(Bind1)(Control_Monad_Gen_Class.chooseInt(dictMonadGen)(0)(size))(function (newSize) {
                            return Control_Monad_Gen_Class.resize(dictMonadGen)(Data_Function["const"](newSize))(Data_Functor.map(Functor0)(fromFoldable)(Control_Monad_Gen.unfoldable(dictMonadRec)(dictMonadGen)(Data_List_Types.unfoldableList)(Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(Data_Tuple.Tuple.create)(genKey))(genValue))));
                        });
                    });
                };
            };
        };
    };
};
export {
    genMap
};
