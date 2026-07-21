import * as $runtime from "../runtime.js";
import * as Data$dArray from "../Data.Array/index.js";
import * as Data$dEither from "../Data.Either/index.js";
import * as Data$dFoldable from "../Data.Foldable/index.js";
import * as Data$dFunction from "../Data.Function/index.js";
import * as Data$dLazy from "../Data.Lazy/index.js";
import * as Data$dList$dTypes from "../Data.List.Types/index.js";
import * as Data$dMap$dInternal from "../Data.Map.Internal/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dOrd from "../Data.Ord/index.js";
import * as Data$dString$dCodeUnits from "../Data.String.CodeUnits/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
const lookup = k => {
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
        const v1 = Data$dOrd.ordString.compare(k)(v._3);
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
  return go;
};
const fromFoldable = /* #__PURE__ */ Data$dFoldable.foldlArray(m => a => Data$dMap$dInternal.insert(Data$dOrd.ordString)(a)()(m))(Data$dMap$dInternal.Leaf);
const member = k => {
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
        const v1 = Data$dOrd.ordString.compare(k)(v._3);
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
  return go;
};
const runSort = /* #__PURE__ */ (() => {
  const go = go$a0$copy => go$a1$copy => go$a2$copy => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$a2 = go$a2$copy, go$c = true, go$r;
    while (go$c) {
      const acc = go$a0, modIndex = go$a1, v = go$a2;
      if (v.tag === "Cons" && v._1.tag === "Left") {
        go$a0 = Data$dList$dTypes.$List("Cons", v._1._1, acc);
        go$a1 = modIndex;
        go$a2 = v._2;
        continue;
      }
      const v1 = v2 => {
        if (v.tag === "Cons") {
          go$a0 = acc;
          go$a1 = modIndex;
          go$a2 = v._2;
          return;
        }
        if (v.tag === "Nil") {
          const go$1 = go$1$a0$copy => go$1$a1$copy => {
            let go$1$a0 = go$1$a0$copy, go$1$a1 = go$1$a1$copy, go$1$c = true, go$1$r;
            while (go$1$c) {
              const v$1 = go$1$a0, v1 = go$1$a1;
              if (v1.tag === "Nil") {
                go$1$c = false;
                go$1$r = v$1;
                continue;
              }
              if (v1.tag === "Cons") {
                go$1$a0 = Data$dList$dTypes.$List("Cons", v1._1, v$1);
                go$1$a1 = v1._2;
                continue;
              }
              $runtime.fail();
            }
            return go$1$r;
          };
          go$c = false;
          go$r = go$1(Data$dList$dTypes.Nil)(acc);
          return;
        }
        $runtime.fail();
      };
      if (v.tag === "Cons" && v._1.tag === "Right") {
        const $0 = lookup(v._1._1)(modIndex);
        if ($0.tag === "Just" && !$0._1._1) {
          go$a0 = acc;
          go$a1 = Data$dMap$dInternal.insert(Data$dOrd.ordString)(v._1._1)(Data$dTuple.$Tuple(true, $0._1._2))(modIndex);
          go$a2 = Data$dFoldable.foldrArray(x => Data$dList$dTypes.Cons(Data$dEither.$Either("Right", x._2)))(Data$dList$dTypes.$List(
            "Cons",
            Data$dEither.$Either("Left", $0._1._2),
            v._2
          ))($0._1._2.imports);
          continue;
        }
      }
      v1(true);
    }
    return go$r;
  };
  return go(Data$dList$dTypes.Nil);
})();
const sortModules = dictFoldable => init => runSort(dictFoldable.foldr(m => Data$dMap$dInternal.insert(Data$dOrd.ordString)(m.name)(Data$dTuple.$Tuple(false, m)))(Data$dMap$dInternal.Leaf)(init))(dictFoldable.foldr(x => Data$dList$dTypes.Cons(Data$dEither.$Either(
  "Right",
  x.name
)))(Data$dList$dTypes.Nil)(init));
const resumePull = v => v.resume;
const pullResult = v => v.result;
const emptyPull = /* #__PURE__ */ (() => {
  const resume = index => needed => v => {
    const index$p = Data$dMap$dInternal.insert(Data$dOrd.ordString)(v.name)(Data$dTuple.$Tuple(false, v))(index);
    const needed$p = Data$dMap$dInternal.unsafeUnionWith(
      Data$dOrd.ordString.compare,
      Data$dFunction.const,
      Data$dMap$dInternal.delete(Data$dOrd.ordString)(v.name)(needed),
      fromFoldable(Data$dArray.mapMaybe(v1 => {
        if (!(member(v1._2)(index$p) || v1._2 === "Prim" || Data$dString$dCodeUnits.take(5)(v1._2) === "Prim.")) { return Data$dMaybe.$Maybe("Just", v1._2); }
        return Data$dMaybe.Nothing;
      })(v.imports))
    );
    if (needed$p.tag === "Leaf") {
      return {
        result: Data$dEither.$Either(
          "Right",
          Data$dLazy.defer(v2 => runSort(index$p)((() => {
            const go = (m$p, z$p) => {
              if (m$p.tag === "Leaf") { return z$p; }
              if (m$p.tag === "Node") { return go(m$p._5, Data$dList$dTypes.$List("Cons", Data$dEither.$Either("Right", m$p._4._2.name), go(m$p._6, z$p))); }
              $runtime.fail();
            };
            return go(index$p, Data$dList$dTypes.Nil);
          })()))
        ),
        resume: resume(index$p)(needed$p)
      };
    }
    return {result: Data$dEither.$Either("Left", needed$p), resume: resume(index$p)(needed$p)};
  };
  return {result: Data$dEither.$Either("Right", Data$dLazy.defer(v => Data$dList$dTypes.Nil)), resume: resume(Data$dMap$dInternal.Leaf)(Data$dMap$dInternal.Leaf)};
})();
export {emptyPull, fromFoldable, lookup, member, pullResult, resumePull, runSort, sortModules};
