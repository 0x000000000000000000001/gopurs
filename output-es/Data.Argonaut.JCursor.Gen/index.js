import * as $runtime from "../runtime.js";
import * as Data$dArgonaut$dJCursor from "../Data.Argonaut.JCursor/index.js";
import * as Data$dChar$dGen from "../Data.Char.Gen/index.js";
import * as Data$dOrd from "../Data.Ord/index.js";
import * as Data$dString$dGen from "../Data.String.Gen/index.js";
const min = x => y => {
  const v = Data$dOrd.ordInt.compare(x)(y);
  if (v === "LT") { return x; }
  if (v === "EQ") { return x; }
  if (v === "GT") { return y; }
  $runtime.fail();
};
const genJCursor = dictMonadGen => {
  const Monad0 = dictMonadGen.Monad0();
  const Apply0 = Monad0.Bind1().Apply0();
  const $0 = Apply0.Functor0();
  const chooseBool = dictMonadGen.chooseBool;
  return dictMonadRec => {
    const genUnicodeString = Data$dString$dGen.genString(dictMonadRec)(dictMonadGen)(Data$dChar$dGen.genUnicodeChar(dictMonadGen));
    return dictLazy => {
      const genIndex = Apply0.apply($0.map(Data$dArgonaut$dJCursor.JIndex)(dictMonadGen.chooseInt(0)(1000)))(dictLazy.defer(v => genJCursor(dictMonadGen)(dictMonadRec)(dictLazy)));
      const genField = Apply0.apply($0.map(Data$dArgonaut$dJCursor.JField)(genUnicodeString))(dictLazy.defer(v => genJCursor(dictMonadGen)(dictMonadRec)(dictLazy)));
      return dictMonadGen.resize(min(10))(dictMonadGen.sized(size => {
        if (size > 0) {
          return dictMonadGen.resize(v => v - 1 | 0)(dictMonadGen.Monad0().Bind1().bind(chooseBool)(v => {
            if (v) { return genField; }
            return genIndex;
          }));
        }
        return Monad0.Applicative0().pure(Data$dArgonaut$dJCursor.JCursorTop);
      }));
    };
  };
};
export {genJCursor, min};
