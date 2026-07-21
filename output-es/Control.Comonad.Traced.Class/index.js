// | This module defines the `ComonadTraced` type class and its instances.
import * as Control$dComonad$dEnv$dTrans from "../Control.Comonad.Env.Trans/index.js";
import * as Control$dComonad$dStore$dTrans from "../Control.Comonad.Store.Trans/index.js";
import * as Control$dComonad$dTraced$dTrans from "../Control.Comonad.Traced.Trans/index.js";
import * as Control$dComonad$dTrans$dClass from "../Control.Comonad.Trans.Class/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
const track = dict => dict.track;
const tracks = dictComonadTraced => f => w => dictComonadTraced.track(f(dictComonadTraced.Comonad0().extract(w)))(w);
const lowerTrack2 = dictComonadTraced => {
  const lower1 = Control$dComonad$dTrans$dClass.comonadTransIdentityT.lower(dictComonadTraced.Comonad0());
  return m => {
    const $0 = dictComonadTraced.track(m);
    return x => $0(lower1(x));
  };
};
const listens = dictFunctor => f => v => dictFunctor.map(g => t => Data$dTuple.$Tuple(g(t), f(t)))(v);
const listen = dictFunctor => v => dictFunctor.map(f => t => Data$dTuple.$Tuple(f(t), t))(v);
const comonadTracedTracedT = dictComonad => {
  const comonadTracedT = Control$dComonad$dTraced$dTrans.comonadTracedT(dictComonad);
  return dictMonoid => {
    const comonadTracedT1 = comonadTracedT(dictMonoid);
    return {track: t => v => dictComonad.extract(v)(t), Comonad0: () => comonadTracedT1};
  };
};
const comonadTracedStoreT = dictComonadTraced => {
  const comonadStoreT = Control$dComonad$dStore$dTrans.comonadStoreT(dictComonadTraced.Comonad0());
  return {
    track: (() => {
      const $0 = dictComonadTraced.Comonad0();
      return m => {
        const $1 = dictComonadTraced.track(m);
        return x => $1((() => {
          const $2 = x._2;
          return $0.Extend0().Functor0().map(v1 => v1($2))(x._1);
        })());
      };
    })(),
    Comonad0: () => comonadStoreT
  };
};
const comonadTracedIdentityT = dictComonadTraced => {
  const $0 = dictComonadTraced.Comonad0();
  const $1 = $0.Extend0();
  const comonadIdentityT = (() => {
    const $2 = $1.Functor0();
    const extendIdentityI1 = {extend: f => v => $1.extend(x => f(x))(v), Functor0: () => $2};
    return {extract: x => $0.extract(x), Extend0: () => extendIdentityI1};
  })();
  return {track: lowerTrack2(dictComonadTraced), Comonad0: () => comonadIdentityT};
};
const comonadTracedEnvT = dictComonadTraced => {
  const comonadEnvT = Control$dComonad$dEnv$dTrans.comonadEnvT(dictComonadTraced.Comonad0());
  return {
    track: m => {
      const $0 = dictComonadTraced.track(m);
      return x => $0(x._2);
    },
    Comonad0: () => comonadEnvT
  };
};
const censor = dictFunctor => f => v => dictFunctor.map(v1 => x => v1(f(x)))(v);
export {censor, comonadTracedEnvT, comonadTracedIdentityT, comonadTracedStoreT, comonadTracedTracedT, listen, listens, lowerTrack2, track, tracks};
