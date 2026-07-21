import * as $runtime from "../runtime.js";
import * as Data$dConst from "../Data.Const/index.js";
import * as Data$dEither from "../Data.Either/index.js";
const Forget = x => x;
const semigroupForget = dictSemigroup => ({append: f => g => x => dictSemigroup.append(f(x))(g(x))});
const profunctorForget = {dimap: f => v => v1 => x => v1(f(x))};
const strongForget = {first: v => x => v(x._1), second: v => x => v(x._2), Profunctor0: () => profunctorForget};
const newtypeForget = {Coercible0: () => {}};
const monoidForget = dictMonoid => {
  const mempty1 = dictMonoid.mempty;
  const $0 = dictMonoid.Semigroup0();
  const semigroupFn = {append: f => g => x => $0.append(f(x))(g(x))};
  return {mempty: v => mempty1, Semigroup0: () => semigroupFn};
};
const cochoiceForget = {unleft: v => x => v(Data$dEither.$Either("Left", x)), unright: v => x => v(Data$dEither.$Either("Right", x)), Profunctor0: () => profunctorForget};
const choiceForget = dictMonoid => {
  const mempty1 = dictMonoid.mempty;
  return {
    left: v => v2 => {
      if (v2.tag === "Left") { return v(v2._1); }
      if (v2.tag === "Right") { return mempty1; }
      $runtime.fail();
    },
    right: v => v2 => {
      if (v2.tag === "Left") { return mempty1; }
      if (v2.tag === "Right") { return v(v2._1); }
      $runtime.fail();
    },
    Profunctor0: () => profunctorForget
  };
};
const wanderForget = dictMonoid => {
  const applicativeConst = Data$dConst.applicativeConst(dictMonoid);
  const choiceForget1 = choiceForget(dictMonoid);
  return {wander: f => v => f(applicativeConst)(v), Strong0: () => strongForget, Choice1: () => choiceForget1};
};
export {Forget, choiceForget, cochoiceForget, monoidForget, newtypeForget, profunctorForget, semigroupForget, strongForget, wanderForget};
