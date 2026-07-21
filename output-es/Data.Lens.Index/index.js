import * as $runtime from "../runtime.js";
import * as Data$dArray from "../Data.Array/index.js";
import * as Data$dEither from "../Data.Either/index.js";
import * as Data$dLens$dAffineTraversal from "../Data.Lens.AffineTraversal/index.js";
import * as Data$dLens$dIso$dNewtype from "../Data.Lens.Iso.Newtype/index.js";
import * as Data$dLens$dPrism$dMaybe from "../Data.Lens.Prism.Maybe/index.js";
import * as Data$dList from "../Data.List/index.js";
import * as Data$dMap$dInternal from "../Data.Map.Internal/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
import * as Foreign$dObject from "../Foreign.Object/index.js";
const ix = dict => dict.ix;
const indexSet = dictOrd => (
  {
    ix: x => dictStrong => dictChoice => Data$dLens$dAffineTraversal.affineTraversal(xs => v => xs)(xs => {
      if (
        (() => {
          const go = go$a0$copy => {
            let go$a0 = go$a0$copy, go$c = true, go$r;
            while (go$c) {
              const v = go$a0;
              if (v.tag === "Leaf") {
                go$c = false;
                go$r = false;
                continue;
              }
              if (v.tag === "Node") {
                const v1 = dictOrd.compare(x)(v._3);
                if (v1 === "LT") {
                  go$a0 = v._5;
                  continue;
                }
                if (v1 === "GT") {
                  go$a0 = v._6;
                  continue;
                }
                if (v1 === "EQ") {
                  go$c = false;
                  go$r = true;
                  continue;
                }
              }
              $runtime.fail();
            }
            return go$r;
          };
          return go(xs);
        })()
      ) {
        return Data$dEither.$Either("Right", undefined);
      }
      return Data$dEither.$Either("Left", xs);
    })(dictStrong)(dictChoice)
  }
);
const indexNonEmptyArray = {
  ix: n => dictStrong => dictChoice => Data$dLens$dAffineTraversal.affineTraversal(s => b => {
    const $0 = Data$dArray._updateAt(Data$dMaybe.Just, Data$dMaybe.Nothing, n, b, s);
    if ($0.tag === "Nothing") { return s; }
    if ($0.tag === "Just") { return $0._1; }
    $runtime.fail();
  })(s => {
    if (n >= 0 && n < s.length) { return Data$dEither.$Either("Right", s[n]); }
    return Data$dEither.$Either("Left", s);
  })(dictStrong)(dictChoice)
};
const indexMaybe = {ix: v => dictStrong => dictChoice => Data$dLens$dPrism$dMaybe._Just(dictChoice)};
const indexMap = dictOrd => (
  {
    ix: k => dictStrong => dictChoice => Data$dLens$dAffineTraversal.affineTraversal(s => b => Data$dMap$dInternal.update(dictOrd)(v => Data$dMaybe.$Maybe("Just", b))(k)(s))(s => {
      const go = go$a0$copy => {
        let go$a0 = go$a0$copy, go$c = true, go$r;
        while (go$c) {
          const v = go$a0;
          if (v.tag === "Leaf") {
            go$c = false;
            go$r = Data$dMaybe.Nothing;
            continue;
          }
          if (v.tag === "Node") {
            const v1 = dictOrd.compare(k)(v._3);
            if (v1 === "LT") {
              go$a0 = v._5;
              continue;
            }
            if (v1 === "GT") {
              go$a0 = v._6;
              continue;
            }
            if (v1 === "EQ") {
              go$c = false;
              go$r = Data$dMaybe.$Maybe("Just", v._4);
              continue;
            }
          }
          $runtime.fail();
        }
        return go$r;
      };
      const $0 = go(s);
      if ($0.tag === "Nothing") { return Data$dEither.$Either("Left", s); }
      if ($0.tag === "Just") { return Data$dEither.$Either("Right", $0._1); }
      $runtime.fail();
    })(dictStrong)(dictChoice)
  }
);
const indexList = {
  ix: n => dictStrong => dictChoice => Data$dLens$dAffineTraversal.affineTraversal(s => b => {
    const $0 = Data$dList.updateAt(n)(b)(s);
    if ($0.tag === "Nothing") { return s; }
    if ($0.tag === "Just") { return $0._1; }
    $runtime.fail();
  })(s => {
    const $0 = Data$dList.index(s)(n);
    if ($0.tag === "Nothing") { return Data$dEither.$Either("Left", s); }
    if ($0.tag === "Just") { return Data$dEither.$Either("Right", $0._1); }
    $runtime.fail();
  })(dictStrong)(dictChoice)
};
const indexIdentity = {ix: v => dictStrong => dictChoice => Data$dLens$dIso$dNewtype._Newtype()()(dictChoice.Profunctor0())};
const indexForeignObject = {
  ix: k => dictStrong => dictChoice => Data$dLens$dAffineTraversal.affineTraversal(s => b => Foreign$dObject.update(v => Data$dMaybe.$Maybe("Just", b))(k)(s))(s => {
    const $0 = Foreign$dObject._lookup(Data$dMaybe.Nothing, Data$dMaybe.Just, k, s);
    if ($0.tag === "Nothing") { return Data$dEither.$Either("Left", s); }
    if ($0.tag === "Just") { return Data$dEither.$Either("Right", $0._1); }
    $runtime.fail();
  })(dictStrong)(dictChoice)
};
const indexFn = dictEq => (
  {
    ix: i => dictStrong => dictChoice => pab => dictStrong.Profunctor0().dimap(s => Data$dTuple.$Tuple(
      s(i),
      b => j => {
        if (dictEq.eq(i)(j)) { return b; }
        return s(j);
      }
    ))(v => v._2(v._1))(dictStrong.first(pab))
  }
);
const indexArray = {
  ix: n => dictStrong => dictChoice => Data$dLens$dAffineTraversal.affineTraversal(s => b => {
    const $0 = Data$dArray._updateAt(Data$dMaybe.Just, Data$dMaybe.Nothing, n, b, s);
    if ($0.tag === "Nothing") { return s; }
    if ($0.tag === "Just") { return $0._1; }
    $runtime.fail();
  })(s => {
    if (n >= 0 && n < s.length) { return Data$dEither.$Either("Right", s[n]); }
    return Data$dEither.$Either("Left", s);
  })(dictStrong)(dictChoice)
};
export {indexArray, indexFn, indexForeignObject, indexIdentity, indexList, indexMap, indexMaybe, indexNonEmptyArray, indexSet, ix};
