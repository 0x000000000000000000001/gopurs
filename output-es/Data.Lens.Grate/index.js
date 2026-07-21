// | This module defines functions for working with grates.
// |
// | See <http://r6research.livejournal.com/28050.html>.
import * as Data$dFunction from "../Data.Function/index.js";
import * as Data$dFunctor from "../Data.Functor/index.js";
const identity = x => x;
const zipWithOf = g => f => g(f);
const zipFWithOf = g => f => g(f);
const withGrate = g => g(f => f(identity));
const grate = f => dictClosed => pab => dictClosed.Profunctor0().dimap(Data$dFunction.applyFlipped)(f)(dictClosed.closed(pab));
const cotraversed = dictDistributive => {
  const distribute2 = dictDistributive.distribute(Data$dFunctor.functorFn);
  return dictClosed => grate(f => dictDistributive.Functor0().map(f)(distribute2(identity)))(dictClosed);
};
const collectOf = dictFunctor => g => f => {
  const $0 = g(identity);
  const $1 = dictFunctor.map(f);
  return x => $0($1(x));
};
const cloneGrate = g => dictClosed => grate(g(f => f(identity)))(dictClosed);
export {cloneGrate, collectOf, cotraversed, grate, identity, withGrate, zipFWithOf, zipWithOf};
