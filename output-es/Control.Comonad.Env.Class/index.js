// | This module defines the `ComonadEnv` type class and its instances.
import * as Control$dComonad$dEnv$dTrans from "../Control.Comonad.Env.Trans/index.js";
import * as Control$dComonad$dStore$dTrans from "../Control.Comonad.Store.Trans/index.js";
import * as Control$dComonad$dTraced$dTrans from "../Control.Comonad.Traced.Trans/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
const local = dict => dict.local;
const comonadAskTuple = {ask: Data$dTuple.fst, Comonad0: () => Data$dTuple.comonadTuple};
const comonadEnvTuple = {local: f => v => Data$dTuple.$Tuple(f(v._1), v._2), ComonadAsk0: () => comonadAskTuple};
const comonadAskEnvT = dictComonad => {
  const comonadEnvT = Control$dComonad$dEnv$dTrans.comonadEnvT(dictComonad);
  return {ask: v => v._1, Comonad0: () => comonadEnvT};
};
const comonadEnvEnvT = dictComonad => {
  const comonadEnvT = Control$dComonad$dEnv$dTrans.comonadEnvT(dictComonad);
  return {local: f => v => Data$dTuple.$Tuple(f(v._1), v._2), ComonadAsk0: () => ({ask: v => v._1, Comonad0: () => comonadEnvT})};
};
const ask = dict => dict.ask;
const asks = dictComonadAsk => f => x => f(dictComonadAsk.ask(x));
const comonadAskStoreT = dictComonadAsk => {
  const Comonad0 = dictComonadAsk.Comonad0();
  const comonadStoreT = Control$dComonad$dStore$dTrans.comonadStoreT(Comonad0);
  return {
    ask: x => dictComonadAsk.ask((() => {
      const $0 = x._2;
      return Comonad0.Extend0().Functor0().map(v1 => v1($0))(x._1);
    })()),
    Comonad0: () => comonadStoreT
  };
};
const comonadEnvStoreT = dictComonadEnv => {
  const comonadAskStoreT1 = comonadAskStoreT(dictComonadEnv.ComonadAsk0());
  return {local: f => v => Data$dTuple.$Tuple(dictComonadEnv.local(f)(v._1), v._2), ComonadAsk0: () => comonadAskStoreT1};
};
const comonadAskTracedT = dictComonadAsk => {
  const Comonad0 = dictComonadAsk.Comonad0();
  const comonadTracedT = Control$dComonad$dTraced$dTrans.comonadTracedT(Comonad0);
  return dictMonoid => {
    const comonadTracedT1 = comonadTracedT(dictMonoid);
    return {
      ask: (() => {
        const mempty = dictMonoid.mempty;
        return x => dictComonadAsk.ask(Comonad0.Extend0().Functor0().map(f => f(mempty))(x));
      })(),
      Comonad0: () => comonadTracedT1
    };
  };
};
const comonadEnvTracedT = dictComonadEnv => {
  const comonadAskTracedT1 = comonadAskTracedT(dictComonadEnv.ComonadAsk0());
  return dictMonoid => {
    const comonadAskTracedT2 = comonadAskTracedT1(dictMonoid);
    return {local: f => v => dictComonadEnv.local(f)(v), ComonadAsk0: () => comonadAskTracedT2};
  };
};
export {ask, asks, comonadAskEnvT, comonadAskStoreT, comonadAskTracedT, comonadAskTuple, comonadEnvEnvT, comonadEnvStoreT, comonadEnvTracedT, comonadEnvTuple, local};
