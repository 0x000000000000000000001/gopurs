// | This module defines functions for working with getters.
import * as Control$dSemigroupoid from "../Control.Semigroupoid/index.js";
import * as Data$dProfunctor$dStrong from "../Data.Profunctor.Strong/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
const identity = x => x;
const fanout = /* #__PURE__ */ Data$dProfunctor$dStrong.fanout(Control$dSemigroupoid.semigroupoidFn)(Data$dProfunctor$dStrong.strongFn);
const view = l => l(identity);
const viewOn = s => l => l(identity)(s);
const use = dictMonadState => p => dictMonadState.state(s => Data$dTuple.$Tuple(p(identity)(s), s));
const to = f => p => x => p(f(x));
const takeBoth = l => r => {
  const $0 = fanout(l(identity))(r(identity));
  return p => x => p($0(x));
};
const iview = l => l(identity);
const iuse = dictMonadState => p => {
  const $0 = p(identity);
  return dictMonadState.state(s => Data$dTuple.$Tuple($0(s), s));
};
const cloneGetter = g => {
  const $0 = g(identity);
  return p => x => p($0(x));
};
export {cloneGetter, fanout, identity, iuse, iview, takeBoth, to, use, view, viewOn};
