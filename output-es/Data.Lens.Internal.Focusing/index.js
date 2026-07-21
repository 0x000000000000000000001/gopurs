// | This module defines the `Focusing` functor
import * as Data$dTuple from "../Data.Tuple/index.js";
const Focusing = x => x;
const newtypeFocusing = {Coercible0: () => {}};
const focusingFunctor = dictFunctor => ({map: f => v => dictFunctor.map(m => Data$dTuple.$Tuple(m._1, f(m._2)))(v)});
const focusingApply = dictApply => {
  const Functor0 = dictApply.Functor0();
  const focusingFunctor1 = {map: f => v => Functor0.map(m => Data$dTuple.$Tuple(m._1, f(m._2)))(v)};
  return dictSemigroup => {
    const apply1 = Data$dTuple.applyTuple(dictSemigroup).apply;
    return {apply: v => v1 => dictApply.apply(Functor0.map(apply1)(v))(v1), Functor0: () => focusingFunctor1};
  };
};
const focusingApplicative = dictApplicative => {
  const focusingApply1 = focusingApply(dictApplicative.Apply0());
  return dictMonoid => {
    const focusingApply2 = focusingApply1(dictMonoid.Semigroup0());
    return {pure: x => dictApplicative.pure(Data$dTuple.applicativeTuple(dictMonoid).pure(x)), Apply0: () => focusingApply2};
  };
};
export {Focusing, focusingApplicative, focusingApply, focusingFunctor, newtypeFocusing};
