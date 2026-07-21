// | This module defines the `Market` profunctor
import * as $runtime from "../runtime.js";
import * as Data$dEither from "../Data.Either/index.js";
const $Market = (_1, _2) => ({tag: "Market", _1, _2});
const Market = value0 => value1 => $Market(value0, value1);
const profunctorMarket = {
  dimap: f => g => v => $Market(
    x => g(v._1(x)),
    x => {
      const $0 = v._2(f(x));
      if ($0.tag === "Left") { return Data$dEither.$Either("Left", g($0._1)); }
      if ($0.tag === "Right") { return Data$dEither.$Either("Right", $0._1); }
      $runtime.fail();
    }
  )
};
const functorMarket = {
  map: f => v => $Market(
    x => f(v._1(x)),
    x => {
      const $0 = v._2(x);
      if ($0.tag === "Left") { return Data$dEither.$Either("Left", f($0._1)); }
      if ($0.tag === "Right") { return Data$dEither.$Either("Right", $0._1); }
      $runtime.fail();
    }
  )
};
const choiceMarket = {
  left: v => $Market(
    x => Data$dEither.$Either("Left", v._1(x)),
    v2 => {
      if (v2.tag === "Left") {
        const $0 = v._2(v2._1);
        if ($0.tag === "Left") { return Data$dEither.$Either("Left", Data$dEither.$Either("Left", $0._1)); }
        if ($0.tag === "Right") { return Data$dEither.$Either("Right", $0._1); }
        $runtime.fail();
      }
      if (v2.tag === "Right") { return Data$dEither.$Either("Left", Data$dEither.$Either("Right", v2._1)); }
      $runtime.fail();
    }
  ),
  right: v => $Market(
    x => Data$dEither.$Either("Right", v._1(x)),
    v2 => {
      if (v2.tag === "Left") { return Data$dEither.$Either("Left", Data$dEither.$Either("Left", v2._1)); }
      if (v2.tag === "Right") {
        const $0 = v._2(v2._1);
        if ($0.tag === "Left") { return Data$dEither.$Either("Left", Data$dEither.$Either("Right", $0._1)); }
        if ($0.tag === "Right") { return Data$dEither.$Either("Right", $0._1); }
      }
      $runtime.fail();
    }
  ),
  Profunctor0: () => profunctorMarket
};
export {$Market, Market, choiceMarket, functorMarket, profunctorMarket};
