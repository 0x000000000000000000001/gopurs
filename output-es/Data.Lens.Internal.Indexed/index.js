// | This module defines the `Indexed` profunctor.
import * as $runtime from "../runtime.js";
import * as Data$dEither from "../Data.Either/index.js";
import * as Data$dProfunctor from "../Data.Profunctor/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
const Indexed = x => x;
const profunctorIndexed = dictProfunctor => ({dimap: f => g => v => dictProfunctor.dimap(m => Data$dTuple.$Tuple(m._1, f(m._2)))(g)(v)});
const strongIndexed = dictStrong => {
  const Profunctor0 = dictStrong.Profunctor0();
  const profunctorIndexed1 = {dimap: f => g => v => Profunctor0.dimap(m => Data$dTuple.$Tuple(m._1, f(m._2)))(g)(v)};
  return {
    first: v => Profunctor0.dimap(v1 => Data$dTuple.$Tuple(Data$dTuple.$Tuple(v1._1, v1._2._1), v1._2._2))(Data$dProfunctor.identity)(dictStrong.first(v)),
    second: v => Profunctor0.dimap(v1 => Data$dTuple.$Tuple(v1._2._1, Data$dTuple.$Tuple(v1._1, v1._2._2)))(Data$dProfunctor.identity)(dictStrong.second(v)),
    Profunctor0: () => profunctorIndexed1
  };
};
const newtypeIndexed = {Coercible0: () => {}};
const choiceIndexed = dictChoice => {
  const Profunctor0 = dictChoice.Profunctor0();
  const profunctorIndexed1 = {dimap: f => g => v => Profunctor0.dimap(m => Data$dTuple.$Tuple(m._1, f(m._2)))(g)(v)};
  return {
    left: v => Profunctor0.dimap(v1 => {
      const $0 = Data$dTuple.Tuple(v1._1);
      if (v1._2.tag === "Left") { return Data$dEither.$Either("Left", $0(v1._2._1)); }
      if (v1._2.tag === "Right") { return Data$dEither.$Either("Right", v1._2._1); }
      $runtime.fail();
    })(Data$dProfunctor.identity)(dictChoice.left(v)),
    right: v => Profunctor0.dimap(v1 => {
      const $0 = Data$dTuple.Tuple(v1._1);
      if (v1._2.tag === "Left") { return Data$dEither.$Either("Left", v1._2._1); }
      if (v1._2.tag === "Right") { return Data$dEither.$Either("Right", $0(v1._2._1)); }
      $runtime.fail();
    })(Data$dProfunctor.identity)(dictChoice.right(v)),
    Profunctor0: () => profunctorIndexed1
  };
};
const wanderIndexed = dictWander => {
  const strongIndexed1 = strongIndexed(dictWander.Strong0());
  const choiceIndexed1 = choiceIndexed(dictWander.Choice1());
  return {
    wander: trav => v => dictWander.wander(dictApplicative => {
      const trav2 = trav(dictApplicative);
      return ia2fb => v1 => trav2((() => {
        const $0 = Data$dTuple.Tuple(v1._1);
        return x => ia2fb($0(x));
      })())(v1._2);
    })(v),
    Strong0: () => strongIndexed1,
    Choice1: () => choiceIndexed1
  };
};
export {Indexed, choiceIndexed, newtypeIndexed, profunctorIndexed, strongIndexed, wanderIndexed};
