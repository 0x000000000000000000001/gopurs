// | This module defines functions for working with isomorphisms.
import * as $runtime from "../runtime.js";
import * as Data$dFunction from "../Data.Function/index.js";
import * as Data$dLens$dInternal$dExchange from "../Data.Lens.Internal.Exchange/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dProfunctor from "../Data.Profunctor/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
import * as Unsafe$dCoerce from "../Unsafe.Coerce/index.js";
const identity = x => x;
const withIso = l => f => {
  const v = l(Data$dLens$dInternal$dExchange.$Exchange(identity, identity));
  return f(v._1)(v._2);
};
const under = l => {
  const v = l(Data$dLens$dInternal$dExchange.$Exchange(identity, identity));
  return ts => x => v._1(ts(v._2(x)));
};
const re = t => t(identity);
const iso = f => g => dictProfunctor => pab => dictProfunctor.dimap(f)(g)(pab);
const mapping = dictFunctor => dictFunctor1 => l => dictProfunctor => {
  const v = l(Data$dLens$dInternal$dExchange.$Exchange(identity, identity));
  const $0 = dictFunctor.map(v._1);
  const $1 = dictFunctor1.map(v._2);
  return pab => dictProfunctor.dimap($0)($1)(pab);
};
const non = dictEq => def => dictProfunctor => pab => dictProfunctor.dimap(v2 => {
  if (v2.tag === "Nothing") { return def; }
  if (v2.tag === "Just") { return v2._1; }
  $runtime.fail();
})(a => {
  if (dictEq.eq(a)(def)) { return Data$dMaybe.Nothing; }
  return Data$dMaybe.$Maybe("Just", a);
})(pab);
const uncurried = dictProfunctor => pab => dictProfunctor.dimap(Data$dTuple.uncurry)(Data$dTuple.curry)(pab);
const flipped = dictProfunctor => pab => dictProfunctor.dimap(Data$dFunction.flip)(Data$dFunction.flip)(pab);
const dimapping = dictProfunctor => dictProfunctor1 => f => g => dictProfunctor2 => {
  const v = f(Data$dLens$dInternal$dExchange.$Exchange(identity, identity));
  const v$1 = g(Data$dLens$dInternal$dExchange.$Exchange(identity, identity));
  const $0 = dictProfunctor.dimap(v._1)(v$1._1);
  const $1 = dictProfunctor1.dimap(v._2)(v$1._2);
  return pab => dictProfunctor2.dimap($0)($1)(pab);
};
const curried = dictProfunctor => pab => dictProfunctor.dimap(Data$dTuple.curry)(Data$dTuple.uncurry)(pab);
const coerced = () => () => dictProfunctor => pab => dictProfunctor.dimap(Unsafe$dCoerce.unsafeCoerce)(Unsafe$dCoerce.unsafeCoerce)(pab);
const cloneIso = l => dictProfunctor => {
  const v = l(Data$dLens$dInternal$dExchange.$Exchange(identity, identity));
  const $0 = v._1;
  const $1 = v._2;
  return p => dictProfunctor.dimap($0)($1)(p);
};
const auf = dictProfunctor => l => {
  const v = l(Data$dLens$dInternal$dExchange.$Exchange(identity, identity));
  const $0 = v._1;
  return f => g => e => v._2(f(dictProfunctor.dimap(Data$dProfunctor.identity)($0)(g))(e));
};
const au = l => {
  const v = l(Data$dLens$dInternal$dExchange.$Exchange(identity, identity));
  const $0 = v._2;
  return f => e => v._1(f($0)(e));
};
export {au, auf, cloneIso, coerced, curried, dimapping, flipped, identity, iso, mapping, non, re, uncurried, under, withIso};
