import * as Data$dBitraversable from "../Data.Bitraversable/index.js";
import * as Data$dTraversable from "../Data.Traversable/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
const Bazaar = x => x;
const runBazaar = v => dictApplicative => v(dictApplicative);
const profunctorBazaar = {
  dimap: f => g => v => dictApplicative => {
    const b1 = v(dictApplicative);
    return pafb => s => dictApplicative.Apply0().Functor0().map(g)(b1(pafb)(f(s)));
  }
};
const strongBazaar = {
  first: v => dictApplicative => {
    const b1 = v(dictApplicative);
    return pafb => v1 => {
      const $0 = v1._2;
      return dictApplicative.Apply0().Functor0().map(a => Data$dTuple.$Tuple(a, $0))(b1(pafb)(v1._1));
    };
  },
  second: v => dictApplicative => {
    const b1 = v(dictApplicative);
    return pafb => v1 => dictApplicative.Apply0().Functor0().map(Data$dTuple.Tuple(v1._1))(b1(pafb)(v1._2));
  },
  Profunctor0: () => profunctorBazaar
};
const choiceBazaar = {
  left: v => dictApplicative => {
    const bitraverse1 = Data$dBitraversable.bitraversableEither.bitraverse(dictApplicative);
    const b1 = v(dictApplicative);
    const pure = dictApplicative.pure;
    return pafb => e => bitraverse1(b1(pafb))(pure)(e);
  },
  right: v => dictApplicative => {
    const b1 = v(dictApplicative);
    return pafb => e => Data$dTraversable.traversableEither.traverse(dictApplicative)(b1(pafb))(e);
  },
  Profunctor0: () => profunctorBazaar
};
const wanderBazaar = {
  wander: w => v => dictApplicative => {
    const w2 = w(dictApplicative);
    const f1 = v(dictApplicative);
    return pafb => s => w2(f1(pafb))(s);
  },
  Strong0: () => strongBazaar,
  Choice1: () => choiceBazaar
};
export {Bazaar, choiceBazaar, profunctorBazaar, runBazaar, strongBazaar, wanderBazaar};
