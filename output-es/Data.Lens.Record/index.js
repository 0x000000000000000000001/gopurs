import * as Data$dTuple from "../Data.Tuple/index.js";
import * as Record$dUnsafe from "../Record.Unsafe/index.js";
const prop = dictIsSymbol => () => () => l => dictStrong => pab => dictStrong.Profunctor0().dimap(s => Data$dTuple.$Tuple(
  Record$dUnsafe.unsafeGet(dictIsSymbol.reflectSymbol(l))(s),
  b => Record$dUnsafe.unsafeSet(dictIsSymbol.reflectSymbol(l))(b)(s)
))(v => v._2(v._1))(dictStrong.first(pab));
export {prop};
