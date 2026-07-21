// | This module defines the `Shop` profunctor
import * as Data$dTuple from "../Data.Tuple/index.js";
const $Shop = (_1, _2) => ({tag: "Shop", _1, _2});
const Shop = value0 => value1 => $Shop(value0, value1);
const profunctorShop = {
  dimap: f => g => v => $Shop(
    x => v._1(f(x)),
    s => {
      const $0 = v._2(f(s));
      return x => g($0(x));
    }
  )
};
const strongShop = {
  first: v => $Shop(
    v1 => v._1(v1._1),
    v1 => {
      const $0 = v1._2;
      const $1 = v1._1;
      return b => Data$dTuple.$Tuple(v._2($1)(b), $0);
    }
  ),
  second: v => $Shop(
    v1 => v._1(v1._2),
    v1 => {
      const $0 = v1._1;
      const $1 = v1._2;
      return b => Data$dTuple.$Tuple($0, v._2($1)(b));
    }
  ),
  Profunctor0: () => profunctorShop
};
export {$Shop, Shop, profunctorShop, strongShop};
