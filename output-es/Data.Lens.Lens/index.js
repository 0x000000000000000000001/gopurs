// | This module defines functions for working with lenses.
import * as Data$dLens$dInternal$dShop from "../Data.Lens.Internal.Shop/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
const identity = x => x;
const withLens = l => f => {
  const v = l(Data$dLens$dInternal$dShop.$Shop(identity, v1 => b => b));
  return f(v._1)(v._2);
};
const withIndexedLens = l => f => {
  const v = l(Data$dLens$dInternal$dShop.$Shop(identity, v1 => b => b));
  return f(v._1)(v._2);
};
const lensStore = l => withLens(l)(a => b => x => Data$dTuple.$Tuple(a(x), b(x)));
const lens$p = to => dictStrong => pab => dictStrong.Profunctor0().dimap(to)(v => v._2(v._1))(dictStrong.first(pab));
const lens = $$get => $$set => dictStrong => pab => dictStrong.Profunctor0().dimap(s => Data$dTuple.$Tuple($$get(s), b => $$set(s)(b)))(v => v._2(v._1))(dictStrong.first(pab));
const ilens$p = to => dictStrong => pab => dictStrong.Profunctor0().dimap(to)(v => v._2(v._1))(dictStrong.first(pab));
const ilens = $$get => $$set => dictStrong => pab => dictStrong.Profunctor0().dimap(s => Data$dTuple.$Tuple($$get(s), b => $$set(s)(b)))(v => v._2(v._1))(dictStrong.first(pab));
const cloneLens = l => dictStrong => withLens(l)(x => y => p => dictStrong.Profunctor0().dimap(s => Data$dTuple.$Tuple(x(s), b => y(s)(b)))(v => v._2(v._1))(dictStrong.first(p)));
const cloneIndexedLens = l => dictStrong => withIndexedLens(l)(x => y => p => dictStrong.Profunctor0().dimap(s => Data$dTuple.$Tuple(x(s), b => y(s)(b)))(v => v._2(v._1))(dictStrong.first(p)));
export {cloneIndexedLens, cloneLens, identity, ilens, ilens$p, lens, lens$p, lensStore, withIndexedLens, withLens};
