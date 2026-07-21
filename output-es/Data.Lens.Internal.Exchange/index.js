// | This module defines the `Exchange` profunctor
const $Exchange = (_1, _2) => ({tag: "Exchange", _1, _2});
const Exchange = value0 => value1 => $Exchange(value0, value1);
const profunctorExchange = {dimap: f => g => v => $Exchange(x => v._1(f(x)), x => g(v._2(x)))};
const functorExchange = {map: f => v => $Exchange(v._1, x => f(v._2(x)))};
export {$Exchange, Exchange, functorExchange, profunctorExchange};
