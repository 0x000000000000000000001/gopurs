import * as $runtime from "../runtime.js";
import * as Data$dArgonaut$dCore from "../Data.Argonaut.Core/index.js";
import * as Data$dArgonaut$dDecode$dDecoders from "../Data.Argonaut.Decode.Decoders/index.js";
import * as Data$dArgonaut$dDecode$dError from "../Data.Argonaut.Decode.Error/index.js";
import * as Data$dArgonaut$dEncode$dClass from "../Data.Argonaut.Encode.Class/index.js";
import * as Data$dArray from "../Data.Array/index.js";
import * as Data$dEither from "../Data.Either/index.js";
import * as Data$dFoldable from "../Data.Foldable/index.js";
import * as Data$dInt from "../Data.Int/index.js";
import * as Data$dList$dTypes from "../Data.List.Types/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dOrd from "../Data.Ord/index.js";
import * as Data$dOrdering from "../Data.Ordering/index.js";
import * as Data$dShow from "../Data.Show/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
import * as Data$dUnfoldable from "../Data.Unfoldable/index.js";
import * as Foreign$dObject from "../Foreign.Object/index.js";
const $JCursor = (tag, _1, _2) => ({tag, _1, _2});
const toUnfoldable = /* #__PURE__ */ Foreign$dObject.toUnfoldable(Data$dList$dTypes.unfoldableList);
const fromFoldable = /* #__PURE__ */ (() => Data$dList$dTypes.foldableList.foldr(Data$dList$dTypes.Cons)(Data$dList$dTypes.Nil))();
const fromFoldable1 = /* #__PURE__ */ Data$dFoldable.foldrArray(Data$dList$dTypes.Cons)(Data$dList$dTypes.Nil);
const decodeJson = /* #__PURE__ */ Data$dArgonaut$dDecode$dDecoders.decodeArray(Data$dEither.Right);
const replicate = n => v => Data$dUnfoldable.unfoldableArray.unfoldr(i => {
  if (i <= 0) { return Data$dMaybe.Nothing; }
  return Data$dMaybe.$Maybe("Just", Data$dTuple.$Tuple(v, i - 1 | 0));
})(n);
const JsonPrim = x => x;
const JCursorTop = /* #__PURE__ */ $JCursor("JCursorTop");
const JField = value0 => value1 => $JCursor("JField", value0, value1);
const JIndex = value0 => value1 => $JCursor("JIndex", value0, value1);
const showJCursor = {
  show: v => {
    if (v.tag === "JCursorTop") { return "JCursorTop"; }
    if (v.tag === "JField") { return "(JField " + Data$dShow.showStringImpl(v._1) + " " + showJCursor.show(v._2) + ")"; }
    if (v.tag === "JIndex") { return "(JIndex " + Data$dShow.showIntImpl(v._1) + " " + showJCursor.show(v._2) + ")"; }
    $runtime.fail();
  }
};
const semigroupJCursor = {
  append: v => v1 => {
    if (v1.tag === "JCursorTop") { return v; }
    if (v.tag === "JCursorTop") { return v1; }
    if (v.tag === "JField") { return $JCursor("JField", v._1, semigroupJCursor.append(v._2)(v1)); }
    if (v.tag === "JIndex") { return $JCursor("JIndex", v._1, semigroupJCursor.append(v._2)(v1)); }
    $runtime.fail();
  }
};
const runJsonPrim = v => v;
const showJsonPrim = {show: p => p(Data$dShow.showUnit.show)(Data$dShow.showBoolean.show)(Data$dShow.showNumberImpl)(Data$dShow.showStringImpl)};
const print = v => {
  if (v.tag === "JCursorTop") { return ""; }
  if (v.tag === "JField") { return "." + v._1 + showJCursor.show(v._2); }
  if (v.tag === "JIndex") { return "[" + Data$dShow.showIntImpl(v._1) + "]" + showJCursor.show(v._2); }
  $runtime.fail();
};
const primToJson = p => p(v => Data$dArgonaut$dCore.jsonNull)(Data$dArgonaut$dCore.fromBoolean)(Data$dArgonaut$dCore.fromNumber)(Data$dArgonaut$dCore.fromString);
const primStr = v => v1 => v2 => v3 => f => f(v);
const primNum = v => v1 => v2 => f => v3 => f(v);
const primNull = f => v => v1 => v2 => f();
const primBool = v => v1 => f => v2 => v3 => f(v);
const toPrims$lazy = /* #__PURE__ */ $runtime.binding(() => Data$dArgonaut$dCore.caseJson(v => Data$dList$dTypes.$List(
  "Cons",
  Data$dTuple.$Tuple(JCursorTop, primNull),
  Data$dList$dTypes.Nil
))(b => Data$dList$dTypes.$List("Cons", Data$dTuple.$Tuple(JCursorTop, v1 => f => v2 => v3 => f(b)), Data$dList$dTypes.Nil))(n => Data$dList$dTypes.$List(
  "Cons",
  Data$dTuple.$Tuple(JCursorTop, v1 => v2 => f => v3 => f(n)),
  Data$dList$dTypes.Nil
))(s => Data$dList$dTypes.$List("Cons", Data$dTuple.$Tuple(JCursorTop, v1 => v2 => v3 => f => f(s)), Data$dList$dTypes.Nil))(arr => Data$dList$dTypes.bindList.bind((() => {
  const go = go$a0$copy => go$a1$copy => go$a2$copy => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$a2 = go$a2$copy, go$c = true, go$r;
    while (go$c) {
      const v = go$a0, v1 = go$a1, v2 = go$a2;
      if (v.tag === "Nil") {
        go$c = false;
        go$r = v2;
        continue;
      }
      if (v1.tag === "Nil") {
        go$c = false;
        go$r = v2;
        continue;
      }
      if (v.tag === "Cons" && v1.tag === "Cons") {
        go$a0 = v._2;
        go$a1 = v1._2;
        go$a2 = Data$dList$dTypes.$List("Cons", Data$dTuple.$Tuple(v._1, v1._1), v2);
        continue;
      }
      $runtime.fail();
    }
    return go$r;
  };
  const go$1 = go$1$a0$copy => go$1$a1$copy => {
    let go$1$a0 = go$1$a0$copy, go$1$a1 = go$1$a1$copy, go$1$c = true, go$1$r;
    while (go$1$c) {
      const v = go$1$a0, v1 = go$1$a1;
      if (v1.tag === "Nil") {
        go$1$c = false;
        go$1$r = v;
        continue;
      }
      if (v1.tag === "Cons") {
        go$1$a0 = Data$dList$dTypes.$List("Cons", v1._1, v);
        go$1$a1 = v1._2;
        continue;
      }
      $runtime.fail();
    }
    return go$1$r;
  };
  return go$1(Data$dList$dTypes.Nil)(go((() => {
    const $0 = arr.length - 1 | 0;
    if (0 === $0) { return Data$dList$dTypes.$List("Cons", 0, Data$dList$dTypes.Nil); }
    const go$2 = go$2$a0$copy => go$2$a1$copy => go$2$a2$copy => go$2$a3$copy => {
      let go$2$a0 = go$2$a0$copy, go$2$a1 = go$2$a1$copy, go$2$a2 = go$2$a2$copy, go$2$a3 = go$2$a3$copy, go$2$c = true, go$2$r;
      while (go$2$c) {
        const s = go$2$a0, e = go$2$a1, step = go$2$a2, rest = go$2$a3;
        if (s === e) {
          go$2$c = false;
          go$2$r = Data$dList$dTypes.$List("Cons", s, rest);
          continue;
        }
        go$2$a0 = s + step | 0;
        go$2$a1 = e;
        go$2$a2 = step;
        go$2$a3 = Data$dList$dTypes.$List("Cons", s, rest);
      }
      return go$2$r;
    };
    return go$2($0)(0)(0 > $0 ? 1 : -1)(Data$dList$dTypes.Nil);
  })())(fromFoldable1(arr))(Data$dList$dTypes.Nil));
})())(v => {
  const $0 = v._1;
  return fromFoldable(Data$dList$dTypes.listMap(t => Data$dTuple.$Tuple($JCursor("JIndex", $0, t._1), t._2))(toPrims$lazy()(v._2)));
}))(obj => Data$dList$dTypes.bindList.bind(toUnfoldable(obj))(v => {
  const $0 = v._1;
  return Data$dList$dTypes.listMap(t => Data$dTuple.$Tuple($JCursor("JField", $0, t._1), t._2))(toPrims$lazy()(v._2));
})));
const toPrims = /* #__PURE__ */ toPrims$lazy();
const monoidJCursor = {mempty: JCursorTop, Semigroup0: () => semigroupJCursor};
const inferEmpty = v => {
  if (v.tag === "JCursorTop") { return Data$dArgonaut$dCore.jsonNull; }
  if (v.tag === "JField") { return Data$dArgonaut$dCore.jsonEmptyObject; }
  if (v.tag === "JIndex") { return Data$dArgonaut$dCore.jsonEmptyArray; }
  $runtime.fail();
};
const eqJCursor = {
  eq: x => y => {
    if (x.tag === "JCursorTop") { return y.tag === "JCursorTop"; }
    if (x.tag === "JField") { return y.tag === "JField" && x._1 === y._1 && eqJCursor.eq(x._2)(y._2); }
    return x.tag === "JIndex" && y.tag === "JIndex" && x._1 === y._1 && eqJCursor.eq(x._2)(y._2);
  }
};
const ordJCursor = {
  compare: x => y => {
    if (x.tag === "JCursorTop") {
      if (y.tag === "JCursorTop") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y.tag === "JCursorTop") { return Data$dOrdering.GT; }
    if (x.tag === "JField") {
      if (y.tag === "JField") {
        const v = Data$dOrd.ordString.compare(x._1)(y._1);
        if (v === "LT") { return Data$dOrdering.LT; }
        if (v === "GT") { return Data$dOrdering.GT; }
        return ordJCursor.compare(x._2)(y._2);
      }
      return Data$dOrdering.LT;
    }
    if (y.tag === "JField") { return Data$dOrdering.GT; }
    if (x.tag === "JIndex" && y.tag === "JIndex") {
      const v = Data$dOrd.ordInt.compare(x._1)(y._1);
      if (v === "LT") { return Data$dOrdering.LT; }
      if (v === "GT") { return Data$dOrdering.GT; }
      return ordJCursor.compare(x._2)(y._2);
    }
    $runtime.fail();
  },
  Eq0: () => eqJCursor
};
const encodeJsonJCursor = {
  encodeJson: /* #__PURE__ */ (() => {
    const loop = v => {
      if (v.tag === "JCursorTop") { return []; }
      if (v.tag === "JField") { return [Data$dArgonaut$dCore.fromString(v._1), ...loop(v._2)]; }
      if (v.tag === "JIndex") { return [Data$dArgonaut$dCore.fromNumber(Data$dInt.toNumber(v._1)), ...loop(v._2)]; }
      $runtime.fail();
    };
    return x => Data$dArgonaut$dEncode$dClass.encodeJsonArray(Data$dArgonaut$dEncode$dClass.encodeJsonJson).encodeJson(loop(x));
  })()
};
const downIndex = i => {
  const downIndex$p = v => {
    if (v.tag === "JCursorTop") { return $JCursor("JIndex", i, JCursorTop); }
    if (v.tag === "JField") { return $JCursor("JField", v._1, downIndex$p(v._2)); }
    if (v.tag === "JIndex") { return $JCursor("JIndex", v._1, downIndex$p(v._2)); }
    $runtime.fail();
  };
  return downIndex$p;
};
const downField = i => {
  const downField$p = v => {
    if (v.tag === "JCursorTop") { return $JCursor("JField", i, JCursorTop); }
    if (v.tag === "JField") { return $JCursor("JField", v._1, downField$p(v._2)); }
    if (v.tag === "JIndex") { return $JCursor("JIndex", v._1, downField$p(v._2)); }
    $runtime.fail();
  };
  return downField$p;
};
const insideOut = v => {
  if (v.tag === "JCursorTop") { return JCursorTop; }
  if (v.tag === "JField") {
    const $0 = v._1;
    const downField$p = v$1 => {
      if (v$1.tag === "JCursorTop") { return $JCursor("JField", $0, JCursorTop); }
      if (v$1.tag === "JField") { return $JCursor("JField", v$1._1, downField$p(v$1._2)); }
      if (v$1.tag === "JIndex") { return $JCursor("JIndex", v$1._1, downField$p(v$1._2)); }
      $runtime.fail();
    };
    return downField$p(insideOut(v._2));
  }
  if (v.tag === "JIndex") {
    const $0 = v._1;
    const downIndex$p = v$1 => {
      if (v$1.tag === "JCursorTop") { return $JCursor("JIndex", $0, JCursorTop); }
      if (v$1.tag === "JField") { return $JCursor("JField", v$1._1, downIndex$p(v$1._2)); }
      if (v$1.tag === "JIndex") { return $JCursor("JIndex", v$1._1, downIndex$p(v$1._2)); }
      $runtime.fail();
    };
    return downIndex$p(insideOut(v._2));
  }
  $runtime.fail();
};
const decodeJsonJCursor = {
  decodeJson: j => {
    const loop = arr => {
      if (0 < arr.length) {
        const $0 = Data$dArray.unconsImpl(v => Data$dMaybe.Nothing, v => xs => Data$dMaybe.$Maybe("Just", xs), arr);
        if ($0.tag === "Just") { return goLoop(Data$dTuple.$Tuple(arr[0], $0._1)); }
      }
      return Data$dEither.$Either("Right", JCursorTop);
    };
    const goLoop = v => {
      const $0 = v._1;
      const $1 = loop(v._2);
      return (() => {
        if ($1.tag === "Left") {
          const $2 = $1._1;
          return v$1 => Data$dEither.$Either("Left", $2);
        }
        if ($1.tag === "Right") {
          const $2 = $1._1;
          return f => f($2);
        }
        $runtime.fail();
      })()(c => Data$dArgonaut$dCore._caseJson(
        v$1 => Data$dEither.$Either(
          "Left",
          Data$dArgonaut$dDecode$dError.$JsonDecodeError("Named", "Int or String", Data$dArgonaut$dDecode$dError.$JsonDecodeError("UnexpectedValue", $0))
        ),
        v$1 => Data$dEither.$Either(
          "Left",
          Data$dArgonaut$dDecode$dError.$JsonDecodeError("Named", "Int or String", Data$dArgonaut$dDecode$dError.$JsonDecodeError("UnexpectedValue", $0))
        ),
        x => {
          const $2 = Data$dInt.fromNumber(x);
          if ($2.tag === "Nothing") { return Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "Int")); }
          if ($2.tag === "Just") { return Data$dEither.$Either("Right", $JCursor("JIndex", $2._1, c)); }
          $runtime.fail();
        },
        x => Data$dEither.$Either("Right", $JCursor("JField", x, c)),
        v$1 => Data$dEither.$Either(
          "Left",
          Data$dArgonaut$dDecode$dError.$JsonDecodeError("Named", "Int or String", Data$dArgonaut$dDecode$dError.$JsonDecodeError("UnexpectedValue", $0))
        ),
        v$1 => Data$dEither.$Either(
          "Left",
          Data$dArgonaut$dDecode$dError.$JsonDecodeError("Named", "Int or String", Data$dArgonaut$dDecode$dError.$JsonDecodeError("UnexpectedValue", $0))
        ),
        $0
      ));
    };
    const $0 = decodeJson(j);
    if ($0.tag === "Left") { return Data$dEither.$Either("Left", $0._1); }
    if ($0.tag === "Right") { return loop($0._1); }
    $runtime.fail();
  }
};
const cursorSet = v => v1 => {
  if (v.tag === "JCursorTop") { return x => Data$dMaybe.$Maybe("Just", v1); }
  if (v.tag === "JField") {
    const $0 = v._2;
    const $1 = v._1;
    return Data$dArgonaut$dCore.caseJsonObject((() => {
      const $2 = cursorSet($0)(v1)((() => {
        if ($0.tag === "JCursorTop") { return Data$dArgonaut$dCore.jsonNull; }
        if ($0.tag === "JField") { return Data$dArgonaut$dCore.jsonEmptyObject; }
        if ($0.tag === "JIndex") { return Data$dArgonaut$dCore.jsonEmptyArray; }
        $runtime.fail();
      })());
      if ($2.tag === "Just") {
        return Data$dMaybe.$Maybe(
          "Just",
          (() => {
            const $3 = $2._1;
            return Data$dArgonaut$dCore.fromObject((() => {
              const $4 = {};
              $4[$1] = $3;
              return $4;
            })());
          })()
        );
      }
      return Data$dMaybe.Nothing;
    })())(m => {
      const $2 = cursorSet($0)(v1)((() => {
        const $2 = (() => {
          if ($0.tag === "JCursorTop") { return Data$dArgonaut$dCore.jsonNull; }
          if ($0.tag === "JField") { return Data$dArgonaut$dCore.jsonEmptyObject; }
          if ($0.tag === "JIndex") { return Data$dArgonaut$dCore.jsonEmptyArray; }
          $runtime.fail();
        })();
        const $3 = Foreign$dObject._lookup(Data$dMaybe.Nothing, Data$dMaybe.Just, $1, m);
        if ($3.tag === "Nothing") { return $2; }
        if ($3.tag === "Just") { return $3._1; }
        $runtime.fail();
      })());
      if ($2.tag === "Just") {
        return Data$dMaybe.$Maybe(
          "Just",
          (() => {
            const $3 = $2._1;
            return Data$dArgonaut$dCore.fromObject(Foreign$dObject.mutate($4 => () => {
              $4[$1] = $3;
              return $4;
            })(m));
          })()
        );
      }
      return Data$dMaybe.Nothing;
    });
  }
  if (v.tag === "JIndex") {
    const $0 = v._2;
    const $1 = v._1;
    const setArr = setArr$a0$copy => setArr$a1$copy => setArr$a2$copy => {
      let setArr$a0 = setArr$a0$copy, setArr$a1 = setArr$a1$copy, setArr$a2 = setArr$a2$copy, setArr$c = true, setArr$r;
      while (setArr$c) {
        const xs = setArr$a0, i$p = setArr$a1, v$p = setArr$a2;
        const len = xs.length;
        if (i$p < 0) {
          setArr$c = false;
          setArr$r = Data$dMaybe.Nothing;
          continue;
        }
        if (i$p >= len) {
          setArr$a0 = [...xs, ...replicate((i$p - len | 0) + 1 | 0)(Data$dArgonaut$dCore.jsonNull)];
          setArr$a1 = i$p;
          setArr$a2 = v$p;
          continue;
        }
        const $2 = Data$dArray._updateAt(Data$dMaybe.Just, Data$dMaybe.Nothing, i$p, v$p, xs);
        if ($2.tag === "Just") {
          setArr$c = false;
          setArr$r = Data$dMaybe.$Maybe("Just", Data$dArgonaut$dCore.fromArray($2._1));
          continue;
        }
        setArr$c = false;
        setArr$r = Data$dMaybe.Nothing;
      }
      return setArr$r;
    };
    return Data$dArgonaut$dCore.caseJsonArray((() => {
      const $2 = replicate($1 + 1 | 0)(Data$dArgonaut$dCore.jsonNull);
      const $3 = cursorSet($0)(v1)((() => {
        if ($0.tag === "JCursorTop") { return Data$dArgonaut$dCore.jsonNull; }
        if ($0.tag === "JField") { return Data$dArgonaut$dCore.jsonEmptyObject; }
        if ($0.tag === "JIndex") { return Data$dArgonaut$dCore.jsonEmptyArray; }
        $runtime.fail();
      })());
      const $4 = (() => {
        if ($3.tag === "Just") { return Data$dArray._updateAt(Data$dMaybe.Just, Data$dMaybe.Nothing, $1, $3._1, $2); }
        if ($3.tag === "Nothing") { return Data$dMaybe.Nothing; }
        $runtime.fail();
      })();
      if ($4.tag === "Just") { return Data$dMaybe.$Maybe("Just", Data$dArgonaut$dCore.fromArray($4._1)); }
      return Data$dMaybe.Nothing;
    })())(a => {
      const $2 = setArr(a)($1);
      const $3 = cursorSet($0)(v1)((() => {
        const $3 = (() => {
          if ($0.tag === "JCursorTop") { return Data$dArgonaut$dCore.jsonNull; }
          if ($0.tag === "JField") { return Data$dArgonaut$dCore.jsonEmptyObject; }
          if ($0.tag === "JIndex") { return Data$dArgonaut$dCore.jsonEmptyArray; }
          $runtime.fail();
        })();
        if ($1 >= 0 && $1 < a.length) { return a[$1]; }
        return $3;
      })());
      if ($3.tag === "Just") { return $2($3._1); }
      if ($3.tag === "Nothing") { return Data$dMaybe.Nothing; }
      $runtime.fail();
    });
  }
  $runtime.fail();
};
const fromPrims = lst => {
  const go = go$a0$copy => go$a1$copy => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
    while (go$c) {
      const b = go$a0, v = go$a1;
      if (v.tag === "Nil") {
        go$c = false;
        go$r = b;
        continue;
      }
      if (v.tag === "Cons") {
        go$a0 = (() => {
          const $0 = cursorSet(v._1._1)(v._1._2(v$1 => Data$dArgonaut$dCore.jsonNull)(Data$dArgonaut$dCore.fromBoolean)(Data$dArgonaut$dCore.fromNumber)(Data$dArgonaut$dCore.fromString));
          if (b.tag === "Just") { return $0(b._1); }
          if (b.tag === "Nothing") { return Data$dMaybe.Nothing; }
          $runtime.fail();
        })();
        go$a1 = v._2;
        continue;
      }
      $runtime.fail();
    }
    return go$r;
  };
  return go((() => {
    if (lst.tag === "Nil") { return Data$dMaybe.Nothing; }
    if (lst.tag === "Cons") {
      return Data$dMaybe.$Maybe(
        "Just",
        (() => {
          if (lst._1._1.tag === "JCursorTop") { return Data$dArgonaut$dCore.jsonNull; }
          if (lst._1._1.tag === "JField") { return Data$dArgonaut$dCore.jsonEmptyObject; }
          if (lst._1._1.tag === "JIndex") { return Data$dArgonaut$dCore.jsonEmptyArray; }
          $runtime.fail();
        })()
      );
    }
    $runtime.fail();
  })())(lst);
};
const cursorGet = v => {
  if (v.tag === "JCursorTop") { return Data$dMaybe.Just; }
  if (v.tag === "JField") {
    return Data$dArgonaut$dCore.caseJsonObject(Data$dMaybe.Nothing)((() => {
      const $0 = cursorGet(v._2);
      return a => {
        const $1 = Foreign$dObject._lookup(Data$dMaybe.Nothing, Data$dMaybe.Just, v._1, a);
        if ($1.tag === "Just") { return $0($1._1); }
        if ($1.tag === "Nothing") { return Data$dMaybe.Nothing; }
        $runtime.fail();
      };
    })());
  }
  if (v.tag === "JIndex") {
    const $0 = v._1;
    return Data$dArgonaut$dCore.caseJsonArray(Data$dMaybe.Nothing)((() => {
      const $1 = cursorGet(v._2);
      return a => {
        if ($0 >= 0 && $0 < a.length) { return $1(a[$0]); }
        return Data$dMaybe.Nothing;
      };
    })());
  }
  $runtime.fail();
};
export {
  $JCursor,
  JCursorTop,
  JField,
  JIndex,
  JsonPrim,
  cursorGet,
  cursorSet,
  decodeJson,
  decodeJsonJCursor,
  downField,
  downIndex,
  encodeJsonJCursor,
  eqJCursor,
  fromFoldable,
  fromFoldable1,
  fromPrims,
  inferEmpty,
  insideOut,
  monoidJCursor,
  ordJCursor,
  primBool,
  primNull,
  primNum,
  primStr,
  primToJson,
  print,
  replicate,
  runJsonPrim,
  semigroupJCursor,
  showJCursor,
  showJsonPrim,
  toPrims,
  toUnfoldable
};
