// | This module defines functions for working with affine traversals.
// | An `AffineTraversal` is a `Traversal` that applies to at most one element.
// |
// | These arise most frequently as the composition of a `Lens` with a `Prism`.
import * as $runtime from "../runtime.js";
import * as Control$dSemigroupoid from "../Control.Semigroupoid/index.js";
import * as Data$dEither from "../Data.Either/index.js";
import * as Data$dLens$dInternal$dStall from "../Data.Lens.Internal.Stall/index.js";
import * as Data$dProfunctor$dStrong from "../Data.Profunctor.Strong/index.js";
const identity = x => x;
const fanout = /* #__PURE__ */ Data$dProfunctor$dStrong.fanout(Control$dSemigroupoid.semigroupoidFn)(Data$dProfunctor$dStrong.strongFn);
const withAffineTraversal = l => f => {
  const v = l(Data$dLens$dInternal$dStall.$Stall(v => identity, Data$dEither.Right));
  return f(v._1)(v._2);
};
const affineTraversal$p = to => dictStrong => dictChoice => pab => dictChoice.Profunctor0().dimap(to)(v => {
  if (v._2.tag === "Left") { return v._2._1; }
  if (v._2.tag === "Right") { return v._1(v._2._1); }
  $runtime.fail();
})(dictStrong.second(dictChoice.right(pab)));
const affineTraversal = $$set => pre => dictStrong => dictChoice => {
  const $0 = fanout($$set)(pre);
  return pab => dictChoice.Profunctor0().dimap($0)(v => {
    if (v._2.tag === "Left") { return v._2._1; }
    if (v._2.tag === "Right") { return v._1(v._2._1); }
    $runtime.fail();
  })(dictStrong.second(dictChoice.right(pab)));
};
const cloneAffineTraversal = l => dictStrong => dictChoice => withAffineTraversal(l)(x => y => p => affineTraversal(x)(y)(dictStrong)(dictChoice)(p));
export {affineTraversal, affineTraversal$p, cloneAffineTraversal, fanout, identity, withAffineTraversal};
