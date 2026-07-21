import * as Control$dMonad$dState$dTrans from "../Control.Monad.State.Trans/index.js";
import * as Data$dFunction from "../Data.Function/index.js";
import * as Data$dFunctor$dCompose from "../Data.Functor.Compose/index.js";
import * as Data$dIdentity from "../Data.Identity/index.js";
import * as Data$dLens$dInternal$dWander from "../Data.Lens.Internal.Wander/index.js";
import * as Data$dProfunctor from "../Data.Profunctor/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
const identity = x => x;
const applicativeStateT = /* #__PURE__ */ Control$dMonad$dState$dTrans.applicativeStateT(Data$dIdentity.monadIdentity);
const applicativeCompose = /* #__PURE__ */ Data$dFunctor$dCompose.applicativeCompose(applicativeStateT);
const applyStateT = /* #__PURE__ */ Control$dMonad$dState$dTrans.applyStateT(Data$dIdentity.monadIdentity);
const monadStateStateT = /* #__PURE__ */ Control$dMonad$dState$dTrans.monadStateStateT(Data$dIdentity.monadIdentity);
const $$get = /* #__PURE__ */ (() => monadStateStateT.state(s => Data$dTuple.$Tuple(s, s)))();
const unIndex = dictProfunctor => l => {
  const $0 = dictProfunctor.dimap(Data$dTuple.snd)(identity);
  return x => l($0(x));
};
const reindexed = dictProfunctor => ij => v => {
  const $0 = dictProfunctor.dimap(v$1 => Data$dTuple.$Tuple(ij(v$1._1), v$1._2))(Data$dProfunctor.identity);
  return x => v($0(x));
};
const iwander = itr => dictWander => dictWander.wander(dictApplicative => {
  const itr1 = itr(dictApplicative);
  return f => s => itr1(a => b => f(Data$dTuple.$Tuple(a, b)))(s);
});
const positions = tr => dictWander => dictWander.wander(dictApplicative => {
  const tr1 = tr(Data$dLens$dInternal$dWander.wanderStar(applicativeCompose(dictApplicative)));
  return f => s => tr1(a => applyStateT.apply(applyStateT.Functor0().map(Data$dFunction.const)(applyStateT.apply(s$1 => {
    const $0 = $$get(s$1);
    return Data$dTuple.$Tuple(
      (() => {
        const $1 = $0._1;
        return b => f(Data$dTuple.$Tuple($1, b));
      })(),
      $0._2
    );
  })(applicativeStateT.pure(a))))(monadStateStateT.state(s$1 => {
    const s$p = s$1 + 1 | 0;
    return Data$dTuple.$Tuple(s$p, s$p);
  })))(s)(0)._1;
});
const itraversed = dictTraversableWithIndex => dictWander => dictWander.wander(dictApplicative => {
  const itr1 = dictTraversableWithIndex.traverseWithIndex(dictApplicative);
  return f => s => itr1(a => b => f(Data$dTuple.$Tuple(a, b)))(s);
});
const asIndex = dictProfunctor => l => {
  const $0 = dictProfunctor.dimap(Data$dTuple.fst)(identity);
  return x => l($0(x));
};
export {applicativeCompose, applicativeStateT, applyStateT, asIndex, $$get as get, identity, itraversed, iwander, monadStateStateT, positions, reindexed, unIndex};
