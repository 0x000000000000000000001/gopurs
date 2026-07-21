// | This module defines the `Wander` type class, which is used to define
// | `Traversal`s.
import * as Data$dIdentity from "../Data.Identity/index.js";
import * as Data$dProfunctor$dChoice from "../Data.Profunctor.Choice/index.js";
import * as Data$dProfunctor$dStar from "../Data.Profunctor.Star/index.js";
import * as Data$dProfunctor$dStrong from "../Data.Profunctor.Strong/index.js";
const wanderStar = dictApplicative => {
  const strongStar = Data$dProfunctor$dStar.strongStar(dictApplicative.Apply0().Functor0());
  const choiceStar = Data$dProfunctor$dStar.choiceStar(dictApplicative);
  return {wander: t => v => t(dictApplicative)(v), Strong0: () => strongStar, Choice1: () => choiceStar};
};
const wanderFunction = {wander: t => t(Data$dIdentity.applicativeIdentity), Strong0: () => Data$dProfunctor$dStrong.strongFn, Choice1: () => Data$dProfunctor$dChoice.choiceFn};
const wander = dict => dict.wander;
export {wander, wanderFunction, wanderStar};
