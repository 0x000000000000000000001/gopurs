// | This module defines the `Stall` profunctor
import * as $runtime from "../runtime.js";
import * as Data$dEither from "../Data.Either/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
const $Stall = (_1, _2) => ({tag: "Stall", _1, _2});
const Stall = value0 => value1 => $Stall(value0, value1);
const profunctorStall = {
  dimap: f => g => v => $Stall(
    x => {
      const $0 = v._1(f(x));
      return x$1 => g($0(x$1));
    },
    x => {
      const $0 = v._2(f(x));
      if ($0.tag === "Left") { return Data$dEither.$Either("Left", g($0._1)); }
      if ($0.tag === "Right") { return Data$dEither.$Either("Right", $0._1); }
      $runtime.fail();
    }
  )
};
const strongStall = {
  first: v => $Stall(
    v1 => {
      const $0 = v1._1;
      const $1 = v1._2;
      return b => Data$dTuple.$Tuple(v._1($0)(b), $1);
    },
    v1 => {
      const $0 = v._2(v1._1);
      if ($0.tag === "Left") { return Data$dEither.$Either("Left", Data$dTuple.$Tuple($0._1, v1._2)); }
      if ($0.tag === "Right") { return Data$dEither.$Either("Right", $0._1); }
      $runtime.fail();
    }
  ),
  second: v => $Stall(
    v1 => {
      const $0 = v1._2;
      const $1 = v1._1;
      return b => Data$dTuple.$Tuple($1, v._1($0)(b));
    },
    v1 => {
      const $0 = Data$dTuple.Tuple(v1._1);
      const $1 = v._2(v1._2);
      if ($1.tag === "Left") { return Data$dEither.$Either("Left", $0($1._1)); }
      if ($1.tag === "Right") { return Data$dEither.$Either("Right", $1._1); }
      $runtime.fail();
    }
  ),
  Profunctor0: () => profunctorStall
};
const functorStall = {
  map: f => v => $Stall(
    x => {
      const $0 = v._1(x);
      return x$1 => f($0(x$1));
    },
    x => {
      const $0 = v._2(x);
      if ($0.tag === "Left") { return Data$dEither.$Either("Left", f($0._1)); }
      if ($0.tag === "Right") { return Data$dEither.$Either("Right", $0._1); }
      $runtime.fail();
    }
  )
};
const choiceStall = {
  left: v => $Stall(
    v1 => {
      if (v1.tag === "Left") {
        const $0 = v1._1;
        return b => Data$dEither.$Either("Left", v._1($0)(b));
      }
      if (v1.tag === "Right") {
        const $0 = v1._1;
        return v2 => Data$dEither.$Either("Right", $0);
      }
      $runtime.fail();
    },
    v1 => {
      if (v1.tag === "Left") {
        const $0 = v._2(v1._1);
        if ($0.tag === "Left") { return Data$dEither.$Either("Left", Data$dEither.$Either("Left", $0._1)); }
        if ($0.tag === "Right") { return Data$dEither.$Either("Right", $0._1); }
        $runtime.fail();
      }
      if (v1.tag === "Right") { return Data$dEither.$Either("Left", Data$dEither.$Either("Right", v1._1)); }
      $runtime.fail();
    }
  ),
  right: v => $Stall(
    v1 => {
      if (v1.tag === "Left") {
        const $0 = v1._1;
        return v2 => Data$dEither.$Either("Left", $0);
      }
      if (v1.tag === "Right") {
        const $0 = v1._1;
        return b => Data$dEither.$Either("Right", v._1($0)(b));
      }
      $runtime.fail();
    },
    v1 => {
      if (v1.tag === "Left") { return Data$dEither.$Either("Left", Data$dEither.$Either("Left", v1._1)); }
      if (v1.tag === "Right") {
        const $0 = v._2(v1._1);
        if ($0.tag === "Left") { return Data$dEither.$Either("Left", Data$dEither.$Either("Right", $0._1)); }
        if ($0.tag === "Right") { return Data$dEither.$Either("Right", $0._1); }
      }
      $runtime.fail();
    }
  ),
  Profunctor0: () => profunctorStall
};
export {$Stall, Stall, choiceStall, functorStall, profunctorStall, strongStall};
