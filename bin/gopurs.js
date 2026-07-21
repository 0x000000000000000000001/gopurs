// output-es/runtime.js
function binding(init) {
  let state = 0;
  let value;
  return () => {
    if (state === 2) {
      return value;
    }
    if (state === 1) {
      throw new Error("Binding demanded before initialized");
    }
    state = 1;
    value = init();
    state = 2;
    return value;
  };
}
function fail() {
  throw new Error("Failed pattern match");
}
function intDiv(x, y) {
  if (y > 0) return Math.floor(x / y);
  if (y < 0) return -Math.floor(x / -y);
  return 0;
}

// output-es/Record.Unsafe/foreign.js
var unsafeGet = function(label2) {
  return function(rec) {
    return rec[label2];
  };
};
var unsafeSet = function(label2) {
  return function(value) {
    return function(rec) {
      var copy = {};
      for (var key in rec) {
        if ({}.hasOwnProperty.call(rec, key)) {
          copy[key] = rec[key];
        }
      }
      copy[label2] = value;
      return copy;
    };
  };
};

// output-es/Type.Proxy/index.js
var $$$Proxy = () => ({ tag: "Proxy" });
var $$Proxy = /* @__PURE__ */ $$$Proxy();

// output-es/Data.Show/foreign.js
var showIntImpl = function(n) {
  return n.toString();
};

// output-es/Data.Ordering/index.js
var $Ordering = (tag) => tag;
var LT = /* @__PURE__ */ $Ordering("LT");
var GT = /* @__PURE__ */ $Ordering("GT");
var EQ = /* @__PURE__ */ $Ordering("EQ");

// output-es/Data.Maybe/index.js
var $Maybe = (tag, _1) => ({ tag, _1 });
var Nothing = /* @__PURE__ */ $Maybe("Nothing");
var Just = (value0) => $Maybe("Just", value0);
var isNothing = (v2) => {
  if (v2.tag === "Nothing") {
    return true;
  }
  if (v2.tag === "Just") {
    return false;
  }
  fail();
};
var isJust = (v2) => {
  if (v2.tag === "Nothing") {
    return false;
  }
  if (v2.tag === "Just") {
    return true;
  }
  fail();
};

// output-es/Data.Function/index.js
var $$const = (a) => (v) => a;
var applyFlipped = (x) => (f) => f(x);

// output-es/Control.Semigroupoid/index.js
var semigroupoidFn = { compose: (f) => (g) => (x) => f(g(x)) };

// output-es/Data.Functor/foreign.js
var arrayMap = function(f) {
  return function(arr) {
    var l = arr.length;
    var result = new Array(l);
    for (var i = 0; i < l; i++) {
      result[i] = f(arr[i]);
    }
    return result;
  };
};

// output-es/Data.Functor/index.js
var functorFn = /* @__PURE__ */ (() => ({ map: semigroupoidFn.compose }))();
var functorArray = { map: arrayMap };

// output-es/Control.Apply/index.js
var applyFn = { apply: (f) => (g) => (x) => f(x)(g(x)), Functor0: () => functorFn };

// output-es/Control.Bind/foreign.js
var arrayBind = typeof Array.prototype.flatMap === "function" ? function(arr) {
  return function(f) {
    return arr.flatMap(f);
  };
} : function(arr) {
  return function(f) {
    var result = [];
    var l = arr.length;
    for (var i = 0; i < l; i++) {
      var xs = f(arr[i]);
      var k = xs.length;
      for (var j = 0; j < k; j++) {
        result.push(xs[j]);
      }
    }
    return result;
  };
};

// output-es/Control.Applicative/index.js
var applicativeFn = { pure: (x) => (v) => x, Apply0: () => applyFn };

// output-es/Data.Either/index.js
var $Either = (tag, _1) => ({ tag, _1 });
var Left = (value0) => $Either("Left", value0);
var Right = (value0) => $Either("Right", value0);
var functorEither = {
  map: (f) => (m) => {
    if (m.tag === "Left") {
      return $Either("Left", m._1);
    }
    if (m.tag === "Right") {
      return $Either("Right", f(m._1));
    }
    fail();
  }
};
var applyEither = {
  apply: (v) => (v1) => {
    if (v.tag === "Left") {
      return $Either("Left", v._1);
    }
    if (v.tag === "Right") {
      if (v1.tag === "Left") {
        return $Either("Left", v1._1);
      }
      if (v1.tag === "Right") {
        return $Either("Right", v._1(v1._1));
      }
    }
    fail();
  },
  Functor0: () => functorEither
};
var applicativeEither = { pure: Right, Apply0: () => applyEither };

// output-es/Control.Monad.ST.Uncurried/foreign.js
var runSTFn2 = function runSTFn22(fn) {
  return function(a) {
    return function(b) {
      return function() {
        return fn(a, b);
      };
    };
  };
};

// output-es/Data.Array.ST/foreign.js
var pushImpl = function(a, xs) {
  return xs.push(a);
};

// output-es/Data.Array.ST/index.js
var push = /* @__PURE__ */ runSTFn2(pushImpl);

// output-es/Data.Array.ST.Iterator/index.js
var $Iterator = (_1, _2) => ({ tag: "Iterator", _1, _2 });
var pushWhile = (p) => (iter) => (array) => () => {
  let $$break = false;
  const $0 = iter._2;
  while (/* @__PURE__ */ (() => {
    const $1 = $$break;
    return !$1;
  })()) {
    const i = $0.value;
    const mx = iter._1(i);
    if (mx.tag === "Just" && p(mx._1)) {
      array.push(mx._1);
      iter._2.value;
      const $1 = iter._2.value;
      iter._2.value = $1 + 1 | 0;
      continue;
    }
    $$break = true;
  }
};
var iterate = (iter) => (f) => () => {
  let $$break = false;
  const $0 = iter._2;
  while (/* @__PURE__ */ (() => {
    const $1 = $$break;
    return !$1;
  })()) {
    const i = $0.value;
    const $1 = $0.value;
    $0.value = $1 + 1 | 0;
    const mx = iter._1(i);
    if (mx.tag === "Just") {
      f(mx._1)();
      continue;
    }
    if (mx.tag === "Nothing") {
      $$break = true;
      continue;
    }
    fail();
  }
};

// output-es/Data.Maybe.First/index.js
var semigroupFirst = {
  append: (v) => (v1) => {
    if (v.tag === "Just") {
      return v;
    }
    return v1;
  }
};
var monoidFirst = { mempty: Nothing, Semigroup0: () => semigroupFirst };

// output-es/Data.Foldable/foreign.js
var foldrArray = function(f) {
  return function(init) {
    return function(xs) {
      var acc = init;
      var len = xs.length;
      for (var i = len - 1; i >= 0; i--) {
        acc = f(xs[i])(acc);
      }
      return acc;
    };
  };
};
var foldlArray = function(f) {
  return function(init) {
    return function(xs) {
      var acc = init;
      var len = xs.length;
      for (var i = 0; i < len; i++) {
        acc = f(acc)(xs[i]);
      }
      return acc;
    };
  };
};

// output-es/Data.Foldable/index.js
var identity2 = (x) => x;
var monoidEndo = /* @__PURE__ */ (() => {
  const semigroupEndo1 = { append: (v) => (v1) => (x) => v(v1(x)) };
  return { mempty: (x) => x, Semigroup0: () => semigroupEndo1 };
})();
var monoidDual = /* @__PURE__ */ (() => {
  const $0 = monoidEndo.Semigroup0();
  const semigroupDual1 = { append: (v) => (v1) => $0.append(v1)(v) };
  return { mempty: monoidEndo.mempty, Semigroup0: () => semigroupDual1 };
})();
var foldableTuple = { foldr: (f) => (z) => (v) => f(v._2)(z), foldl: (f) => (z) => (v) => f(z)(v._2), foldMap: (dictMonoid) => (f) => (v) => f(v._2) };
var foldableArray = {
  foldr: foldrArray,
  foldl: foldlArray,
  foldMap: (dictMonoid) => {
    const mempty = dictMonoid.mempty;
    return (f) => foldableArray.foldr((x) => (acc) => dictMonoid.Semigroup0().append(f(x))(acc))(mempty);
  }
};
var foldlDefault = (dictFoldable) => {
  const foldMap23 = dictFoldable.foldMap(monoidDual);
  return (c) => (u) => (xs) => foldMap23((x) => (a) => c(a)(x))(xs)(u);
};
var foldrDefault = (dictFoldable) => {
  const foldMap23 = dictFoldable.foldMap(monoidEndo);
  return (c) => (u) => (xs) => foldMap23((x) => c(x))(xs)(u);
};
var lookup = (dictFoldable) => {
  const foldMap23 = dictFoldable.foldMap(monoidFirst);
  return (dictEq) => (a) => foldMap23((v) => {
    if (dictEq.eq(a)(v._1)) {
      return $Maybe("Just", v._2);
    }
    return Nothing;
  });
};
var or = (dictFoldable) => (dictHeytingAlgebra) => dictFoldable.foldMap((() => {
  const semigroupDisj1 = { append: (v) => (v1) => dictHeytingAlgebra.disj(v)(v1) };
  return { mempty: dictHeytingAlgebra.ff, Semigroup0: () => semigroupDisj1 };
})())(identity2);
var and = (dictFoldable) => (dictHeytingAlgebra) => dictFoldable.foldMap((() => {
  const semigroupConj1 = { append: (v) => (v1) => dictHeytingAlgebra.conj(v)(v1) };
  return { mempty: dictHeytingAlgebra.tt, Semigroup0: () => semigroupConj1 };
})())(identity2);

// output-es/Data.Tuple/index.js
var $Tuple = (_1, _2) => ({ tag: "Tuple", _1, _2 });
var Tuple = (value0) => (value1) => $Tuple(value0, value1);
var snd = (v) => v._2;
var functorTuple = { map: (f) => (m) => $Tuple(m._1, f(m._2)) };
var fst = (v) => v._1;

// output-es/Data.FunctorWithIndex/foreign.js
var mapWithIndexArray = function(f) {
  return function(xs) {
    var l = xs.length;
    var result = Array(l);
    for (var i = 0; i < l; i++) {
      result[i] = f(i)(xs[i]);
    }
    return result;
  };
};

// output-es/Data.FunctorWithIndex/index.js
var functorWithIndexArray = { mapWithIndex: mapWithIndexArray, Functor0: () => functorArray };

// output-es/Data.Eq/foreign.js
var refEq = function(r1) {
  return function(r2) {
    return r1 === r2;
  };
};
var eqBooleanImpl = refEq;
var eqIntImpl = refEq;
var eqNumberImpl = refEq;
var eqCharImpl = refEq;
var eqStringImpl = refEq;
var eqArrayImpl = function(f) {
  return function(xs) {
    return function(ys) {
      if (xs.length !== ys.length) return false;
      for (var i = 0; i < xs.length; i++) {
        if (!f(xs[i])(ys[i])) return false;
      }
      return true;
    };
  };
};

// output-es/Data.Eq/index.js
var eqString = { eq: eqStringImpl };
var eqNumber = { eq: eqNumberImpl };
var eqInt = { eq: eqIntImpl };
var eqChar = { eq: eqCharImpl };
var eqBoolean = { eq: eqBooleanImpl };

// output-es/Data.Ord/foreign.js
var unsafeCompareImpl = function(lt) {
  return function(eq2) {
    return function(gt) {
      return function(x) {
        return function(y) {
          return x < y ? lt : x === y ? eq2 : gt;
        };
      };
    };
  };
};
var ordBooleanImpl = unsafeCompareImpl;
var ordIntImpl = unsafeCompareImpl;
var ordNumberImpl = unsafeCompareImpl;
var ordStringImpl = unsafeCompareImpl;
var ordCharImpl = unsafeCompareImpl;
var ordArrayImpl = function(f) {
  return function(xs) {
    return function(ys) {
      var i = 0;
      var xlen = xs.length;
      var ylen = ys.length;
      while (i < xlen && i < ylen) {
        var x = xs[i];
        var y = ys[i];
        var o = f(x)(y);
        if (o !== 0) {
          return o;
        }
        i++;
      }
      if (xlen === ylen) {
        return 0;
      } else if (xlen > ylen) {
        return -1;
      } else {
        return 1;
      }
    };
  };
};

// output-es/Data.Ord/index.js
var ordString = { compare: /* @__PURE__ */ ordStringImpl(LT)(EQ)(GT), Eq0: () => eqString };
var ordNumber = { compare: /* @__PURE__ */ ordNumberImpl(LT)(EQ)(GT), Eq0: () => eqNumber };
var ordInt = { compare: /* @__PURE__ */ ordIntImpl(LT)(EQ)(GT), Eq0: () => eqInt };
var ordChar = { compare: /* @__PURE__ */ ordCharImpl(LT)(EQ)(GT), Eq0: () => eqChar };
var ordBoolean = { compare: /* @__PURE__ */ ordBooleanImpl(LT)(EQ)(GT), Eq0: () => eqBoolean };
var ordArray = (dictOrd) => {
  const eqArray = { eq: eqArrayImpl(dictOrd.Eq0().eq) };
  return {
    compare: (xs) => (ys) => ordInt.compare(0)(ordArrayImpl((x) => (y) => {
      const v = dictOrd.compare(x)(y);
      if (v === "EQ") {
        return 0;
      }
      if (v === "LT") {
        return 1;
      }
      if (v === "GT") {
        return -1;
      }
      fail();
    })(xs)(ys)),
    Eq0: () => eqArray
  };
};

// output-es/Unsafe.Coerce/foreign.js
var unsafeCoerce = function(x) {
  return x;
};

// output-es/Data.Traversable.Accum.Internal/index.js
var functorStateR = {
  map: (f) => (k) => (s) => {
    const v = k(s);
    return { accum: v.accum, value: f(v.value) };
  }
};
var functorStateL = {
  map: (f) => (k) => (s) => {
    const v = k(s);
    return { accum: v.accum, value: f(v.value) };
  }
};
var applyStateR = {
  apply: (f) => (x) => (s) => {
    const v = x(s);
    const v1 = f(v.accum);
    return { accum: v1.accum, value: v1.value(v.value) };
  },
  Functor0: () => functorStateR
};
var applyStateL = {
  apply: (f) => (x) => (s) => {
    const v = f(s);
    const v1 = x(v.accum);
    return { accum: v1.accum, value: v.value(v1.value) };
  },
  Functor0: () => functorStateL
};
var applicativeStateR = { pure: (a) => (s) => ({ accum: s, value: a }), Apply0: () => applyStateR };
var applicativeStateL = { pure: (a) => (s) => ({ accum: s, value: a }), Apply0: () => applyStateL };

// output-es/Data.Traversable/foreign.js
var traverseArrayImpl = /* @__PURE__ */ (function() {
  function array1(a) {
    return [a];
  }
  function array2(a) {
    return function(b) {
      return [a, b];
    };
  }
  function array3(a) {
    return function(b) {
      return function(c) {
        return [a, b, c];
      };
    };
  }
  function concat22(xs) {
    return function(ys) {
      return xs.concat(ys);
    };
  }
  return function(apply) {
    return function(map) {
      return function(pure) {
        return function(f) {
          return function(array) {
            function go(bot, top) {
              switch (top - bot) {
                case 0:
                  return pure([]);
                case 1:
                  return map(array1)(f(array[bot]));
                case 2:
                  return apply(map(array2)(f(array[bot])))(f(array[bot + 1]));
                case 3:
                  return apply(apply(map(array3)(f(array[bot])))(f(array[bot + 1])))(f(array[bot + 2]));
                default:
                  var pivot = bot + Math.floor((top - bot) / 4) * 2;
                  return apply(map(concat22)(go(bot, pivot)))(go(pivot, top));
              }
            }
            return go(0, array.length);
          };
        };
      };
    };
  };
})();

// output-es/Data.Traversable/index.js
var identity3 = (x) => x;
var traversableTuple = {
  traverse: (dictApplicative) => (f) => (v) => dictApplicative.Apply0().Functor0().map(Tuple(v._1))(f(v._2)),
  sequence: (dictApplicative) => (v) => dictApplicative.Apply0().Functor0().map(Tuple(v._1))(v._2),
  Functor0: () => functorTuple,
  Foldable1: () => foldableTuple
};
var traversableArray = {
  traverse: (dictApplicative) => {
    const Apply0 = dictApplicative.Apply0();
    return traverseArrayImpl(Apply0.apply)(Apply0.Functor0().map)(dictApplicative.pure);
  },
  sequence: (dictApplicative) => traversableArray.traverse(dictApplicative)(identity3),
  Functor0: () => functorArray,
  Foldable1: () => foldableArray
};
var mapAccumR = (dictTraversable) => {
  const traverse22 = dictTraversable.traverse(applicativeStateR);
  return (f) => (s0) => (xs) => traverse22((a) => (s) => f(s)(a))(xs)(s0);
};
var mapAccumL = (dictTraversable) => {
  const traverse22 = dictTraversable.traverse(applicativeStateL);
  return (f) => (s0) => (xs) => traverse22((a) => (s) => f(s)(a))(xs)(s0);
};

// output-es/Data.Array/foreign.js
var replicateFill = function(count, value) {
  if (count < 1) {
    return [];
  }
  var result = new Array(count);
  return result.fill(value);
};
var replicatePolyfill = function(count, value) {
  var result = [];
  var n = 0;
  for (var i = 0; i < count; i++) {
    result[n++] = value;
  }
  return result;
};
var replicateImpl = typeof Array.prototype.fill === "function" ? replicateFill : replicatePolyfill;
var fromFoldableImpl = /* @__PURE__ */ (function() {
  function Cons2(head2, tail) {
    this.head = head2;
    this.tail = tail;
  }
  var emptyList = {};
  function curryCons(head2) {
    return function(tail) {
      return new Cons2(head2, tail);
    };
  }
  function listToArray(list) {
    var result = [];
    var count = 0;
    var xs = list;
    while (xs !== emptyList) {
      result[count++] = xs.head;
      xs = xs.tail;
    }
    return result;
  }
  return function(foldr, xs) {
    return listToArray(foldr(curryCons)(emptyList)(xs));
  };
})();
var unconsImpl = function(empty2, next, xs) {
  return xs.length === 0 ? empty2({}) : next(xs[0])(xs.slice(1));
};
var findMapImpl = function(nothing, isJust2, f, xs) {
  for (var i = 0; i < xs.length; i++) {
    var result = f(xs[i]);
    if (isJust2(result)) return result;
  }
  return nothing;
};
var reverse = function(l) {
  return l.slice().reverse();
};
var filterImpl = function(f, xs) {
  return xs.filter(f);
};
var sortByImpl2 = /* @__PURE__ */ (function() {
  function mergeFromTo(compare4, fromOrdering, xs1, xs2, from, to) {
    var mid;
    var i;
    var j;
    var k;
    var x;
    var y;
    var c;
    mid = from + (to - from >> 1);
    if (mid - from > 1) mergeFromTo(compare4, fromOrdering, xs2, xs1, from, mid);
    if (to - mid > 1) mergeFromTo(compare4, fromOrdering, xs2, xs1, mid, to);
    i = from;
    j = mid;
    k = from;
    while (i < mid && j < to) {
      x = xs2[i];
      y = xs2[j];
      c = fromOrdering(compare4(x)(y));
      if (c > 0) {
        xs1[k++] = y;
        ++j;
      } else {
        xs1[k++] = x;
        ++i;
      }
    }
    while (i < mid) {
      xs1[k++] = xs2[i++];
    }
    while (j < to) {
      xs1[k++] = xs2[j++];
    }
  }
  return function(compare4, fromOrdering, xs) {
    var out;
    if (xs.length < 2) return xs;
    out = xs.slice(0);
    mergeFromTo(compare4, fromOrdering, out, xs.slice(0), 0, xs.length);
    return out;
  };
})();
var sliceImpl = function(s, e, l) {
  return l.slice(s, e);
};
var zipWithImpl = function(f, xs, ys) {
  var l = xs.length < ys.length ? xs.length : ys.length;
  var result = new Array(l);
  for (var i = 0; i < l; i++) {
    result[i] = f(xs[i])(ys[i]);
  }
  return result;
};
var anyImpl = function(p, xs) {
  var len = xs.length;
  for (var i = 0; i < len; i++) {
    if (p(xs[i])) return true;
  }
  return false;
};
var allImpl = function(p, xs) {
  var len = xs.length;
  for (var i = 0; i < len; i++) {
    if (!p(xs[i])) return false;
  }
  return true;
};

// output-es/Data.Array/index.js
var zipWithA = (dictApplicative) => {
  const sequence1 = traversableArray.traverse(dictApplicative)(identity3);
  return (f) => (xs) => (ys) => sequence1(zipWithImpl(f, xs, ys));
};
var sortBy = (comp) => ($0) => sortByImpl2(
  comp,
  (v) => {
    if (v === "GT") {
      return 1;
    }
    if (v === "EQ") {
      return 0;
    }
    if (v === "LT") {
      return -1;
    }
    fail();
  },
  $0
);
var snoc = (xs) => (x) => (() => {
  const $0 = push(x);
  return () => {
    const result = [...xs];
    $0(result)();
    return result;
  };
})()();
var unzip = (xs) => {
  const fsts = [];
  const snds = [];
  const $0 = { value: 0 };
  iterate($Iterator(
    (v) => {
      if (v >= 0 && v < xs.length) {
        return $Maybe("Just", xs[v]);
      }
      return Nothing;
    },
    $0
  ))((v) => {
    const $1 = v._1;
    const $2 = v._2;
    return () => {
      fsts.push($1);
      snds.push($2);
    };
  })();
  return $Tuple(fsts, snds);
};
var groupBy = (op) => (xs) => {
  const result = [];
  const $0 = { value: 0 };
  const iter = $Iterator(
    (v) => {
      if (v >= 0 && v < xs.length) {
        return $Maybe("Just", xs[v]);
      }
      return Nothing;
    },
    $0
  );
  iterate(iter)((x) => () => {
    const sub1 = [];
    sub1.push(x);
    pushWhile(op(x))(iter)(sub1)();
    result.push(sub1);
  })();
  return result;
};
var groupAllBy = (cmp) => {
  const $0 = groupBy((x) => (y) => cmp(x)(y) === "EQ");
  return (x) => $0(sortBy(cmp)(x));
};
var concatMap = (b) => (a) => arrayBind(a)(b);
var mapMaybe = (f) => concatMap((x) => {
  const $0 = f(x);
  if ($0.tag === "Nothing") {
    return [];
  }
  if ($0.tag === "Just") {
    return [$0._1];
  }
  fail();
});
var filterA = (dictApplicative) => {
  const traverse12 = traversableArray.traverse(dictApplicative);
  const $0 = dictApplicative.Apply0().Functor0();
  return (p) => {
    const $1 = traverse12((x) => $0.map(Tuple(x))(p(x)));
    const $2 = $0.map(mapMaybe((v) => {
      if (v._2) {
        return $Maybe("Just", v._1);
      }
      return Nothing;
    }));
    return (x) => $2($1(x));
  };
};
var any = ($0) => ($1) => anyImpl($0, $1);

// output-es/Data.FoldableWithIndex/index.js
var foldableWithIndexArray = {
  foldrWithIndex: (f) => (z) => {
    const $0 = foldrArray((v) => {
      const $02 = v._1;
      const $12 = v._2;
      return (y) => f($02)($12)(y);
    })(z);
    const $1 = mapWithIndexArray(Tuple);
    return (x) => $0($1(x));
  },
  foldlWithIndex: (f) => (z) => {
    const $0 = foldlArray((y) => (v) => f(v._1)(y)(v._2))(z);
    const $1 = mapWithIndexArray(Tuple);
    return (x) => $0($1(x));
  },
  foldMapWithIndex: (dictMonoid) => {
    const mempty = dictMonoid.mempty;
    return (f) => foldableWithIndexArray.foldrWithIndex((i) => (x) => (acc) => dictMonoid.Semigroup0().append(f(i)(x))(acc))(mempty);
  },
  Foldable0: () => foldableArray
};

// output-es/Data.Unfoldable1/foreign.js
var unfoldr1ArrayImpl = function(isNothing2) {
  return function(fromJust3) {
    return function(fst2) {
      return function(snd2) {
        return function(f) {
          return function(b) {
            var result = [];
            var value = b;
            while (true) {
              var tuple = f(value);
              result.push(fst2(tuple));
              var maybe = snd2(tuple);
              if (isNothing2(maybe)) return result;
              value = fromJust3(maybe);
            }
          };
        };
      };
    };
  };
};

// output-es/Data.Unfoldable1/index.js
var fromJust = (v) => {
  if (v.tag === "Just") {
    return v._1;
  }
  fail();
};
var unfoldable1Array = { unfoldr1: /* @__PURE__ */ unfoldr1ArrayImpl(isNothing)(fromJust)(fst)(snd) };

// output-es/Data.Unfoldable/foreign.js
var unfoldrArrayImpl = function(isNothing2) {
  return function(fromJust3) {
    return function(fst2) {
      return function(snd2) {
        return function(f) {
          return function(b) {
            var result = [];
            var value = b;
            while (true) {
              var maybe = f(value);
              if (isNothing2(maybe)) return result;
              var tuple = fromJust3(maybe);
              result.push(fst2(tuple));
              value = snd2(tuple);
            }
          };
        };
      };
    };
  };
};

// output-es/Data.Unfoldable/index.js
var fromJust2 = (v) => {
  if (v.tag === "Just") {
    return v._1;
  }
  fail();
};
var unfoldableArray = {
  unfoldr: /* @__PURE__ */ unfoldrArrayImpl(isNothing)(fromJust2)(fst)(snd),
  Unfoldable10: () => unfoldable1Array
};

// output-es/Foreign.Object/foreign.js
function _lookup(no, yes, k, m) {
  return k in m ? yes(m[k]) : no;
}
function toArrayWithKey(f) {
  return function(m) {
    var r = [];
    for (var k in m) {
      if (hasOwnProperty.call(m, k)) {
        r.push(f(k)(m[k]));
      }
    }
    return r;
  };
}
var keys = Object.keys || toArrayWithKey(function(k) {
  return function() {
    return k;
  };
});

// output-es/Data.Argonaut.Core/foreign.js
function stringify(j) {
  return JSON.stringify(j);
}
function _caseJson(isNull2, isBool, isNum, isStr, isArr, isObj, j) {
  if (j == null) return isNull2();
  else if (typeof j === "boolean") return isBool(j);
  else if (typeof j === "number") return isNum(j);
  else if (typeof j === "string") return isStr(j);
  else if (Object.prototype.toString.call(j) === "[object Array]")
    return isArr(j);
  else return isObj(j);
}

// output-es/Data.Argonaut.Decode.Error/index.js
var $JsonDecodeError = (tag, _1, _2) => ({ tag, _1, _2 });
var AtIndex = (value0) => (value1) => $JsonDecodeError("AtIndex", value0, value1);
var Named = (value0) => (value1) => $JsonDecodeError("Named", value0, value1);
var MissingValue = /* @__PURE__ */ $JsonDecodeError("MissingValue");
var printJsonDecodeError = (err) => {
  const go = (v) => {
    if (v.tag === "TypeMismatch") {
      return "  Expected value of type '" + v._1 + "'.";
    }
    if (v.tag === "UnexpectedValue") {
      return "  Unexpected value " + stringify(v._1) + ".";
    }
    if (v.tag === "AtIndex") {
      return "  At array index " + showIntImpl(v._1) + ":\n" + go(v._2);
    }
    if (v.tag === "AtKey") {
      return "  At object key '" + v._1 + "':\n" + go(v._2);
    }
    if (v.tag === "Named") {
      return "  Under '" + v._1 + "':\n" + go(v._2);
    }
    if (v.tag === "MissingValue") {
      return "  No value was found.";
    }
    fail();
  };
  return "An error occurred while decoding a JSON value:\n" + go(err);
};

// output-es/Data.Argonaut.Parser/foreign.js
function _jsonParser(fail2, succ, s) {
  try {
    return succ(JSON.parse(s));
  } catch (e) {
    return fail2(e.message);
  }
}

// output-es/Data.List.Types/index.js
var $List = (tag, _1, _2) => ({ tag, _1, _2 });
var Nil = /* @__PURE__ */ $List("Nil");
var Cons = (value0) => (value1) => $List("Cons", value0, value1);
var foldableList = {
  foldr: (f) => (b) => {
    const $0 = foldableList.foldl((b$1) => (a) => f(a)(b$1))(b);
    const go = (go$a0$copy) => (go$a1$copy) => {
      let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
      while (go$c) {
        const v = go$a0, v1 = go$a1;
        if (v1.tag === "Nil") {
          go$c = false;
          go$r = v;
          continue;
        }
        if (v1.tag === "Cons") {
          go$a0 = $List("Cons", v1._1, v);
          go$a1 = v1._2;
          continue;
        }
        fail();
      }
      return go$r;
    };
    const $1 = go(Nil);
    return (x) => $0($1(x));
  },
  foldl: (f) => {
    const go = (go$a0$copy) => (go$a1$copy) => {
      let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
      while (go$c) {
        const b = go$a0, v = go$a1;
        if (v.tag === "Nil") {
          go$c = false;
          go$r = b;
          continue;
        }
        if (v.tag === "Cons") {
          go$a0 = f(b)(v._1);
          go$a1 = v._2;
          continue;
        }
        fail();
      }
      return go$r;
    };
    return go;
  },
  foldMap: (dictMonoid) => {
    const mempty = dictMonoid.mempty;
    return (f) => foldableList.foldl((acc) => {
      const $0 = dictMonoid.Semigroup0().append(acc);
      return (x) => $0(f(x));
    })(mempty);
  }
};

// output-es/Data.Map.Internal/index.js
var $$$Map = (tag, _1, _2, _3, _4, _5, _6) => ({ tag, _1, _2, _3, _4, _5, _6 });
var $MapIter = (tag, _1, _2, _3) => ({ tag, _1, _2, _3 });
var $Split = (_1, _2, _3) => ({ tag: "Split", _1, _2, _3 });
var $SplitLast = (_1, _2, _3) => ({ tag: "SplitLast", _1, _2, _3 });
var Leaf = /* @__PURE__ */ $$$Map("Leaf");
var IterLeaf = /* @__PURE__ */ $MapIter("IterLeaf");
var unsafeNode = (k, v, l, r) => {
  if (l.tag === "Leaf") {
    if (r.tag === "Leaf") {
      return $$$Map("Node", 1, 1, k, v, l, r);
    }
    if (r.tag === "Node") {
      return $$$Map("Node", 1 + r._1 | 0, 1 + r._2 | 0, k, v, l, r);
    }
    fail();
  }
  if (l.tag === "Node") {
    if (r.tag === "Leaf") {
      return $$$Map("Node", 1 + l._1 | 0, 1 + l._2 | 0, k, v, l, r);
    }
    if (r.tag === "Node") {
      return $$$Map("Node", l._1 > r._1 ? 1 + l._1 | 0 : 1 + r._1 | 0, (1 + l._2 | 0) + r._2 | 0, k, v, l, r);
    }
  }
  fail();
};
var unsafeBalancedNode = (k, v, l, r) => {
  if (l.tag === "Leaf") {
    if (r.tag === "Leaf") {
      return $$$Map("Node", 1, 1, k, v, Leaf, Leaf);
    }
    if (r.tag === "Node" && r._1 > 1) {
      if (r._5.tag === "Node" && (() => {
        if (r._6.tag === "Leaf") {
          return r._5._1 > 0;
        }
        if (r._6.tag === "Node") {
          return r._5._1 > r._6._1;
        }
        fail();
      })()) {
        return unsafeNode(r._5._3, r._5._4, unsafeNode(k, v, l, r._5._5), unsafeNode(r._3, r._4, r._5._6, r._6));
      }
      return unsafeNode(r._3, r._4, unsafeNode(k, v, l, r._5), r._6);
    }
    return unsafeNode(k, v, l, r);
  }
  if (l.tag === "Node") {
    if (r.tag === "Node") {
      if (r._1 > (l._1 + 1 | 0)) {
        if (r._5.tag === "Node" && (() => {
          if (r._6.tag === "Leaf") {
            return r._5._1 > 0;
          }
          if (r._6.tag === "Node") {
            return r._5._1 > r._6._1;
          }
          fail();
        })()) {
          return unsafeNode(r._5._3, r._5._4, unsafeNode(k, v, l, r._5._5), unsafeNode(r._3, r._4, r._5._6, r._6));
        }
        return unsafeNode(r._3, r._4, unsafeNode(k, v, l, r._5), r._6);
      }
      if (l._1 > (r._1 + 1 | 0)) {
        if (l._6.tag === "Node" && (() => {
          if (l._5.tag === "Leaf") {
            return 0 <= l._6._1;
          }
          if (l._5.tag === "Node") {
            return l._5._1 <= l._6._1;
          }
          fail();
        })()) {
          return unsafeNode(l._6._3, l._6._4, unsafeNode(l._3, l._4, l._5, l._6._5), unsafeNode(k, v, l._6._6, r));
        }
        return unsafeNode(l._3, l._4, l._5, unsafeNode(k, v, l._6, r));
      }
      return unsafeNode(k, v, l, r);
    }
    if (r.tag === "Leaf" && l._1 > 1) {
      if (l._6.tag === "Node" && (() => {
        if (l._5.tag === "Leaf") {
          return 0 <= l._6._1;
        }
        if (l._5.tag === "Node") {
          return l._5._1 <= l._6._1;
        }
        fail();
      })()) {
        return unsafeNode(l._6._3, l._6._4, unsafeNode(l._3, l._4, l._5, l._6._5), unsafeNode(k, v, l._6._6, r));
      }
      return unsafeNode(l._3, l._4, l._5, unsafeNode(k, v, l._6, r));
    }
    return unsafeNode(k, v, l, r);
  }
  fail();
};
var unsafeSplit = (comp, k, m) => {
  if (m.tag === "Leaf") {
    return $Split(Nothing, Leaf, Leaf);
  }
  if (m.tag === "Node") {
    const v = comp(k)(m._3);
    if (v === "LT") {
      const v1 = unsafeSplit(comp, k, m._5);
      return $Split(v1._1, v1._2, unsafeBalancedNode(m._3, m._4, v1._3, m._6));
    }
    if (v === "GT") {
      const v1 = unsafeSplit(comp, k, m._6);
      return $Split(v1._1, unsafeBalancedNode(m._3, m._4, m._5, v1._2), v1._3);
    }
    if (v === "EQ") {
      return $Split($Maybe("Just", m._4), m._5, m._6);
    }
  }
  fail();
};
var unsafeSplitLast = (k, v, l, r) => {
  if (r.tag === "Leaf") {
    return $SplitLast(k, v, l);
  }
  if (r.tag === "Node") {
    const v1 = unsafeSplitLast(r._3, r._4, r._5, r._6);
    return $SplitLast(v1._1, v1._2, unsafeBalancedNode(k, v, l, v1._3));
  }
  fail();
};
var unsafeJoinNodes = (v, v1) => {
  if (v.tag === "Leaf") {
    return v1;
  }
  if (v.tag === "Node") {
    const v2 = unsafeSplitLast(v._3, v._4, v._5, v._6);
    return unsafeBalancedNode(v2._1, v2._2, v2._3, v1);
  }
  fail();
};
var unsafeUnionWith = (comp, app, l, r) => {
  if (l.tag === "Leaf") {
    return r;
  }
  if (r.tag === "Leaf") {
    return l;
  }
  if (r.tag === "Node") {
    const v = unsafeSplit(comp, r._3, l);
    const l$p = unsafeUnionWith(comp, app, v._2, r._5);
    const r$p = unsafeUnionWith(comp, app, v._3, r._6);
    if (v._1.tag === "Just") {
      return unsafeBalancedNode(r._3, app(v._1._1)(r._4), l$p, r$p);
    }
    if (v._1.tag === "Nothing") {
      return unsafeBalancedNode(r._3, r._4, l$p, r$p);
    }
  }
  fail();
};
var update = (dictOrd) => (f) => (k) => {
  const go = (v) => {
    if (v.tag === "Leaf") {
      return Leaf;
    }
    if (v.tag === "Node") {
      const v1 = dictOrd.compare(k)(v._3);
      if (v1 === "LT") {
        return unsafeBalancedNode(v._3, v._4, go(v._5), v._6);
      }
      if (v1 === "GT") {
        return unsafeBalancedNode(v._3, v._4, v._5, go(v._6));
      }
      if (v1 === "EQ") {
        const v2 = f(v._4);
        if (v2.tag === "Nothing") {
          return unsafeJoinNodes(v._5, v._6);
        }
        if (v2.tag === "Just") {
          return $$$Map("Node", v._1, v._2, v._3, v2._1, v._5, v._6);
        }
      }
    }
    fail();
  };
  return go;
};
var pop = (dictOrd) => {
  const compare4 = dictOrd.compare;
  return (k) => (m) => {
    const v = unsafeSplit(compare4, k, m);
    if (v._1.tag === "Just") {
      return $Maybe("Just", $Tuple(v._1._1, unsafeJoinNodes(v._2, v._3)));
    }
    return Nothing;
  };
};
var stepAscCps = (next) => (done) => {
  const go = (go$a0$copy) => {
    let go$a0 = go$a0$copy, go$c = true, go$r;
    while (go$c) {
      const v = go$a0;
      if (v.tag === "IterLeaf") {
        go$c = false;
        go$r = done();
        continue;
      }
      if (v.tag === "IterEmit") {
        go$c = false;
        go$r = next(v._1, v._2, v._3);
        continue;
      }
      if (v.tag === "IterNode") {
        go$a0 = (() => {
          const go$1 = (go$1$a0$copy) => (go$1$a1$copy) => {
            let go$1$a0 = go$1$a0$copy, go$1$a1 = go$1$a1$copy, go$1$c = true, go$1$r;
            while (go$1$c) {
              const iter = go$1$a0, v$1 = go$1$a1;
              if (v$1.tag === "Leaf") {
                go$1$c = false;
                go$1$r = iter;
                continue;
              }
              if (v$1.tag === "Node") {
                if (v$1._6.tag === "Leaf") {
                  go$1$a0 = $MapIter("IterEmit", v$1._3, v$1._4, iter);
                  go$1$a1 = v$1._5;
                  continue;
                }
                go$1$a0 = $MapIter("IterEmit", v$1._3, v$1._4, $MapIter("IterNode", v$1._6, iter));
                go$1$a1 = v$1._5;
                continue;
              }
              fail();
            }
            return go$1$r;
          };
          return go$1(v._2)(v._1);
        })();
        continue;
      }
      fail();
    }
    return go$r;
  };
  return go;
};
var stepUnfoldr = /* @__PURE__ */ stepAscCps((k, v, next) => $Maybe("Just", $Tuple($Tuple(k, v), next)))((v) => Nothing);
var insert = (dictOrd) => (k) => (v) => {
  const go = (v1) => {
    if (v1.tag === "Leaf") {
      return $$$Map("Node", 1, 1, k, v, Leaf, Leaf);
    }
    if (v1.tag === "Node") {
      const v2 = dictOrd.compare(k)(v1._3);
      if (v2 === "LT") {
        return unsafeBalancedNode(v1._3, v1._4, go(v1._5), v1._6);
      }
      if (v2 === "GT") {
        return unsafeBalancedNode(v1._3, v1._4, v1._5, go(v1._6));
      }
      if (v2 === "EQ") {
        return $$$Map("Node", v1._1, v1._2, k, v, v1._5, v1._6);
      }
    }
    fail();
  };
  return go;
};
var filterWithKey = (dictOrd) => (f) => {
  const go = (v) => {
    if (v.tag === "Leaf") {
      return Leaf;
    }
    if (v.tag === "Node") {
      if (f(v._3)(v._4)) {
        return unsafeBalancedNode(v._3, v._4, go(v._5), go(v._6));
      }
      return unsafeJoinNodes(go(v._5), go(v._6));
    }
    fail();
  };
  return go;
};
var fromFoldable = (dictOrd) => (dictFoldable) => dictFoldable.foldl((m) => (v) => insert(dictOrd)(v._1)(v._2)(m))(Leaf);
var $$delete = (dictOrd) => (k) => {
  const go = (v) => {
    if (v.tag === "Leaf") {
      return Leaf;
    }
    if (v.tag === "Node") {
      const v1 = dictOrd.compare(k)(v._3);
      if (v1 === "LT") {
        return unsafeBalancedNode(v._3, v._4, go(v._5), v._6);
      }
      if (v1 === "GT") {
        return unsafeBalancedNode(v._3, v._4, v._5, go(v._6));
      }
      if (v1 === "EQ") {
        return unsafeJoinNodes(v._5, v._6);
      }
    }
    fail();
  };
  return go;
};
var alter = (dictOrd) => {
  const compare4 = dictOrd.compare;
  return (f) => (k) => (m) => {
    const v = unsafeSplit(compare4, k, m);
    const v2 = f(v._1);
    if (v2.tag === "Nothing") {
      return unsafeJoinNodes(v._2, v._3);
    }
    if (v2.tag === "Just") {
      return unsafeBalancedNode(k, v2._1, v._2, v._3);
    }
    fail();
  };
};

// output-es/Data.String.Common/foreign.js
var replaceAll = function(s1) {
  return function(s2) {
    return function(s3) {
      return s3.replace(new RegExp(s1.replace(/[-\/\\^$*+?.()|[\]{}]/g, "\\$&"), "g"), s2);
    };
  };
};
var split = function(sep) {
  return function(s) {
    return s.split(sep);
  };
};
var trim = function(s) {
  return s.trim();
};
var joinWith = function(s) {
  return function(xs) {
    return xs.join(s);
  };
};

// output-es/Control.Monad.Error.Class/index.js
var $$try = (dictMonadError) => {
  const Monad0 = dictMonadError.MonadThrow0().Monad0();
  return (a) => dictMonadError.catchError(Monad0.Bind1().Apply0().Functor0().map(Right)(a))((x) => Monad0.Applicative0().pure($Either("Left", x)));
};

// output-es/Partial/foreign.js
var _crashWith = function(msg) {
  throw new Error(msg);
};

// output-es/Effect.Aff/foreign.js
var Aff = (function() {
  var EMPTY = {};
  var PURE = "Pure";
  var THROW = "Throw";
  var CATCH = "Catch";
  var SYNC = "Sync";
  var ASYNC = "Async";
  var BIND = "Bind";
  var BRACKET = "Bracket";
  var FORK = "Fork";
  var SEQ = "Sequential";
  var MAP = "Map";
  var APPLY = "Apply";
  var ALT = "Alt";
  var CONS = "Cons";
  var RESUME = "Resume";
  var RELEASE = "Release";
  var FINALIZER = "Finalizer";
  var FINALIZED = "Finalized";
  var FORKED = "Forked";
  var FIBER = "Fiber";
  var THUNK = "Thunk";
  function Aff2(tag, _1, _2, _3) {
    this.tag = tag;
    this._1 = _1;
    this._2 = _2;
    this._3 = _3;
  }
  function AffCtr(tag) {
    var fn = function(_1, _2, _3) {
      return new Aff2(tag, _1, _2, _3);
    };
    fn.tag = tag;
    return fn;
  }
  function nonCanceler2(error3) {
    return new Aff2(PURE, void 0);
  }
  function runEff(eff) {
    try {
      eff();
    } catch (error3) {
      setTimeout(function() {
        throw error3;
      }, 0);
    }
  }
  function runSync(left, right, eff) {
    try {
      return right(eff());
    } catch (error3) {
      return left(error3);
    }
  }
  function runAsync(left, eff, k) {
    try {
      return eff(k)();
    } catch (error3) {
      k(left(error3))();
      return nonCanceler2;
    }
  }
  var Scheduler = (function() {
    var limit = 1024;
    var size4 = 0;
    var ix = 0;
    var queue = new Array(limit);
    var draining = false;
    function drain() {
      var thunk;
      draining = true;
      while (size4 !== 0) {
        size4--;
        thunk = queue[ix];
        queue[ix] = void 0;
        ix = (ix + 1) % limit;
        thunk();
      }
      draining = false;
    }
    return {
      isDraining: function() {
        return draining;
      },
      enqueue: function(cb) {
        var i, tmp;
        if (size4 === limit) {
          tmp = draining;
          drain();
          draining = tmp;
        }
        queue[(ix + size4) % limit] = cb;
        size4++;
        if (!draining) {
          drain();
        }
      }
    };
  })();
  function Supervisor(util) {
    var fibers = {};
    var fiberId = 0;
    var count = 0;
    return {
      register: function(fiber) {
        var fid = fiberId++;
        fiber.onComplete({
          rethrow: true,
          handler: function(result) {
            return function() {
              count--;
              delete fibers[fid];
            };
          }
        })();
        fibers[fid] = fiber;
        count++;
      },
      isEmpty: function() {
        return count === 0;
      },
      killAll: function(killError, cb) {
        return function() {
          if (count === 0) {
            return cb();
          }
          var killCount = 0;
          var kills = {};
          function kill(fid) {
            kills[fid] = fibers[fid].kill(killError, function(result) {
              return function() {
                delete kills[fid];
                killCount--;
                if (util.isLeft(result) && util.fromLeft(result)) {
                  setTimeout(function() {
                    throw util.fromLeft(result);
                  }, 0);
                }
                if (killCount === 0) {
                  cb();
                }
              };
            })();
          }
          for (var k in fibers) {
            if (fibers.hasOwnProperty(k)) {
              killCount++;
              kill(k);
            }
          }
          fibers = {};
          fiberId = 0;
          count = 0;
          return function(error3) {
            return new Aff2(SYNC, function() {
              for (var k2 in kills) {
                if (kills.hasOwnProperty(k2)) {
                  kills[k2]();
                }
              }
            });
          };
        };
      }
    };
  }
  var SUSPENDED = 0;
  var CONTINUE = 1;
  var STEP_BIND = 2;
  var STEP_RESULT = 3;
  var PENDING = 4;
  var RETURN = 5;
  var COMPLETED = 6;
  function Fiber(util, supervisor, aff) {
    var runTick = 0;
    var status = SUSPENDED;
    var step = aff;
    var fail2 = null;
    var interrupt = null;
    var bhead = null;
    var btail = null;
    var attempts = null;
    var bracketCount = 0;
    var joinId = 0;
    var joins = null;
    var rethrow = true;
    function run2(localRunTick) {
      var tmp, result, attempt;
      while (true) {
        tmp = null;
        result = null;
        attempt = null;
        switch (status) {
          case STEP_BIND:
            status = CONTINUE;
            try {
              step = bhead(step);
              if (btail === null) {
                bhead = null;
              } else {
                bhead = btail._1;
                btail = btail._2;
              }
            } catch (e) {
              status = RETURN;
              fail2 = util.left(e);
              step = null;
            }
            break;
          case STEP_RESULT:
            if (util.isLeft(step)) {
              status = RETURN;
              fail2 = step;
              step = null;
            } else if (bhead === null) {
              status = RETURN;
            } else {
              status = STEP_BIND;
              step = util.fromRight(step);
            }
            break;
          case CONTINUE:
            switch (step.tag) {
              case BIND:
                if (bhead) {
                  btail = new Aff2(CONS, bhead, btail);
                }
                bhead = step._2;
                status = CONTINUE;
                step = step._1;
                break;
              case PURE:
                if (bhead === null) {
                  status = RETURN;
                  step = util.right(step._1);
                } else {
                  status = STEP_BIND;
                  step = step._1;
                }
                break;
              case SYNC:
                status = STEP_RESULT;
                step = runSync(util.left, util.right, step._1);
                break;
              case ASYNC:
                status = PENDING;
                step = runAsync(util.left, step._1, function(result2) {
                  return function() {
                    if (runTick !== localRunTick) {
                      return;
                    }
                    runTick++;
                    Scheduler.enqueue(function() {
                      if (runTick !== localRunTick + 1) {
                        return;
                      }
                      status = STEP_RESULT;
                      step = result2;
                      run2(runTick);
                    });
                  };
                });
                return;
              case THROW:
                status = RETURN;
                fail2 = util.left(step._1);
                step = null;
                break;
              // Enqueue the Catch so that we can call the error handler later on
              // in case of an exception.
              case CATCH:
                if (bhead === null) {
                  attempts = new Aff2(CONS, step, attempts, interrupt);
                } else {
                  attempts = new Aff2(CONS, step, new Aff2(CONS, new Aff2(RESUME, bhead, btail), attempts, interrupt), interrupt);
                }
                bhead = null;
                btail = null;
                status = CONTINUE;
                step = step._1;
                break;
              // Enqueue the Bracket so that we can call the appropriate handlers
              // after resource acquisition.
              case BRACKET:
                bracketCount++;
                if (bhead === null) {
                  attempts = new Aff2(CONS, step, attempts, interrupt);
                } else {
                  attempts = new Aff2(CONS, step, new Aff2(CONS, new Aff2(RESUME, bhead, btail), attempts, interrupt), interrupt);
                }
                bhead = null;
                btail = null;
                status = CONTINUE;
                step = step._1;
                break;
              case FORK:
                status = STEP_RESULT;
                tmp = Fiber(util, supervisor, step._2);
                if (supervisor) {
                  supervisor.register(tmp);
                }
                if (step._1) {
                  tmp.run();
                }
                step = util.right(tmp);
                break;
              case SEQ:
                status = CONTINUE;
                step = sequential(util, supervisor, step._1);
                break;
            }
            break;
          case RETURN:
            bhead = null;
            btail = null;
            if (attempts === null) {
              status = COMPLETED;
              step = interrupt || fail2 || step;
            } else {
              tmp = attempts._3;
              attempt = attempts._1;
              attempts = attempts._2;
              switch (attempt.tag) {
                // We cannot recover from an unmasked interrupt. Otherwise we should
                // continue stepping, or run the exception handler if an exception
                // was raised.
                case CATCH:
                  if (interrupt && interrupt !== tmp && bracketCount === 0) {
                    status = RETURN;
                  } else if (fail2) {
                    status = CONTINUE;
                    step = attempt._2(util.fromLeft(fail2));
                    fail2 = null;
                  }
                  break;
                // We cannot resume from an unmasked interrupt or exception.
                case RESUME:
                  if (interrupt && interrupt !== tmp && bracketCount === 0 || fail2) {
                    status = RETURN;
                  } else {
                    bhead = attempt._1;
                    btail = attempt._2;
                    status = STEP_BIND;
                    step = util.fromRight(step);
                  }
                  break;
                // If we have a bracket, we should enqueue the handlers,
                // and continue with the success branch only if the fiber has
                // not been interrupted. If the bracket acquisition failed, we
                // should not run either.
                case BRACKET:
                  bracketCount--;
                  if (fail2 === null) {
                    result = util.fromRight(step);
                    attempts = new Aff2(CONS, new Aff2(RELEASE, attempt._2, result), attempts, tmp);
                    if (interrupt === tmp || bracketCount > 0) {
                      status = CONTINUE;
                      step = attempt._3(result);
                    }
                  }
                  break;
                // Enqueue the appropriate handler. We increase the bracket count
                // because it should not be cancelled.
                case RELEASE:
                  attempts = new Aff2(CONS, new Aff2(FINALIZED, step, fail2), attempts, interrupt);
                  status = CONTINUE;
                  if (interrupt && interrupt !== tmp && bracketCount === 0) {
                    step = attempt._1.killed(util.fromLeft(interrupt))(attempt._2);
                  } else if (fail2) {
                    step = attempt._1.failed(util.fromLeft(fail2))(attempt._2);
                  } else {
                    step = attempt._1.completed(util.fromRight(step))(attempt._2);
                  }
                  fail2 = null;
                  bracketCount++;
                  break;
                case FINALIZER:
                  bracketCount++;
                  attempts = new Aff2(CONS, new Aff2(FINALIZED, step, fail2), attempts, interrupt);
                  status = CONTINUE;
                  step = attempt._1;
                  break;
                case FINALIZED:
                  bracketCount--;
                  status = RETURN;
                  step = attempt._1;
                  fail2 = attempt._2;
                  break;
              }
            }
            break;
          case COMPLETED:
            for (var k in joins) {
              if (joins.hasOwnProperty(k)) {
                rethrow = rethrow && joins[k].rethrow;
                runEff(joins[k].handler(step));
              }
            }
            joins = null;
            if (interrupt && fail2) {
              setTimeout(function() {
                throw util.fromLeft(fail2);
              }, 0);
            } else if (util.isLeft(step) && rethrow) {
              setTimeout(function() {
                if (rethrow) {
                  throw util.fromLeft(step);
                }
              }, 0);
            }
            return;
          case SUSPENDED:
            status = CONTINUE;
            break;
          case PENDING:
            return;
        }
      }
    }
    function onComplete(join2) {
      return function() {
        if (status === COMPLETED) {
          rethrow = rethrow && join2.rethrow;
          join2.handler(step)();
          return function() {
          };
        }
        var jid = joinId++;
        joins = joins || {};
        joins[jid] = join2;
        return function() {
          if (joins !== null) {
            delete joins[jid];
          }
        };
      };
    }
    function kill(error3, cb) {
      return function() {
        if (status === COMPLETED) {
          cb(util.right(void 0))();
          return function() {
          };
        }
        var canceler = onComplete({
          rethrow: false,
          handler: function() {
            return cb(util.right(void 0));
          }
        })();
        switch (status) {
          case SUSPENDED:
            interrupt = util.left(error3);
            status = COMPLETED;
            step = interrupt;
            run2(runTick);
            break;
          case PENDING:
            if (interrupt === null) {
              interrupt = util.left(error3);
            }
            if (bracketCount === 0) {
              if (status === PENDING) {
                attempts = new Aff2(CONS, new Aff2(FINALIZER, step(error3)), attempts, interrupt);
              }
              status = RETURN;
              step = null;
              fail2 = null;
              run2(++runTick);
            }
            break;
          default:
            if (interrupt === null) {
              interrupt = util.left(error3);
            }
            if (bracketCount === 0) {
              status = RETURN;
              step = null;
              fail2 = null;
            }
        }
        return canceler;
      };
    }
    function join(cb) {
      return function() {
        var canceler = onComplete({
          rethrow: false,
          handler: cb
        })();
        if (status === SUSPENDED) {
          run2(runTick);
        }
        return canceler;
      };
    }
    return {
      kill,
      join,
      onComplete,
      isSuspended: function() {
        return status === SUSPENDED;
      },
      run: function() {
        if (status === SUSPENDED) {
          if (!Scheduler.isDraining()) {
            Scheduler.enqueue(function() {
              run2(runTick);
            });
          } else {
            run2(runTick);
          }
        }
      }
    };
  }
  function runPar(util, supervisor, par, cb) {
    var fiberId = 0;
    var fibers = {};
    var killId = 0;
    var kills = {};
    var early = new Error("[ParAff] Early exit");
    var interrupt = null;
    var root = EMPTY;
    function kill(error3, par2, cb2) {
      var step = par2;
      var head2 = null;
      var tail = null;
      var count = 0;
      var kills2 = {};
      var tmp, kid;
      loop: while (true) {
        tmp = null;
        switch (step.tag) {
          case FORKED:
            if (step._3 === EMPTY) {
              tmp = fibers[step._1];
              kills2[count++] = tmp.kill(error3, function(result) {
                return function() {
                  count--;
                  if (count === 0) {
                    cb2(result)();
                  }
                };
              });
            }
            if (head2 === null) {
              break loop;
            }
            step = head2._2;
            if (tail === null) {
              head2 = null;
            } else {
              head2 = tail._1;
              tail = tail._2;
            }
            break;
          case MAP:
            step = step._2;
            break;
          case APPLY:
          case ALT:
            if (head2) {
              tail = new Aff2(CONS, head2, tail);
            }
            head2 = step;
            step = step._1;
            break;
        }
      }
      if (count === 0) {
        cb2(util.right(void 0))();
      } else {
        kid = 0;
        tmp = count;
        for (; kid < tmp; kid++) {
          kills2[kid] = kills2[kid]();
        }
      }
      return kills2;
    }
    function join(result, head2, tail) {
      var fail2, step, lhs, rhs, tmp, kid;
      if (util.isLeft(result)) {
        fail2 = result;
        step = null;
      } else {
        step = result;
        fail2 = null;
      }
      loop: while (true) {
        lhs = null;
        rhs = null;
        tmp = null;
        kid = null;
        if (interrupt !== null) {
          return;
        }
        if (head2 === null) {
          cb(fail2 || step)();
          return;
        }
        if (head2._3 !== EMPTY) {
          return;
        }
        switch (head2.tag) {
          case MAP:
            if (fail2 === null) {
              head2._3 = util.right(head2._1(util.fromRight(step)));
              step = head2._3;
            } else {
              head2._3 = fail2;
            }
            break;
          case APPLY:
            lhs = head2._1._3;
            rhs = head2._2._3;
            if (fail2) {
              head2._3 = fail2;
              tmp = true;
              kid = killId++;
              kills[kid] = kill(early, fail2 === lhs ? head2._2 : head2._1, function() {
                return function() {
                  delete kills[kid];
                  if (tmp) {
                    tmp = false;
                  } else if (tail === null) {
                    join(fail2, null, null);
                  } else {
                    join(fail2, tail._1, tail._2);
                  }
                };
              });
              if (tmp) {
                tmp = false;
                return;
              }
            } else if (lhs === EMPTY || rhs === EMPTY) {
              return;
            } else {
              step = util.right(util.fromRight(lhs)(util.fromRight(rhs)));
              head2._3 = step;
            }
            break;
          case ALT:
            lhs = head2._1._3;
            rhs = head2._2._3;
            if (lhs === EMPTY && util.isLeft(rhs) || rhs === EMPTY && util.isLeft(lhs)) {
              return;
            }
            if (lhs !== EMPTY && util.isLeft(lhs) && rhs !== EMPTY && util.isLeft(rhs)) {
              fail2 = step === lhs ? rhs : lhs;
              step = null;
              head2._3 = fail2;
            } else {
              head2._3 = step;
              tmp = true;
              kid = killId++;
              kills[kid] = kill(early, step === lhs ? head2._2 : head2._1, function() {
                return function() {
                  delete kills[kid];
                  if (tmp) {
                    tmp = false;
                  } else if (tail === null) {
                    join(step, null, null);
                  } else {
                    join(step, tail._1, tail._2);
                  }
                };
              });
              if (tmp) {
                tmp = false;
                return;
              }
            }
            break;
        }
        if (tail === null) {
          head2 = null;
        } else {
          head2 = tail._1;
          tail = tail._2;
        }
      }
    }
    function resolve(fiber) {
      return function(result) {
        return function() {
          delete fibers[fiber._1];
          fiber._3 = result;
          join(result, fiber._2._1, fiber._2._2);
        };
      };
    }
    function run2() {
      var status = CONTINUE;
      var step = par;
      var head2 = null;
      var tail = null;
      var tmp, fid;
      loop: while (true) {
        tmp = null;
        fid = null;
        switch (status) {
          case CONTINUE:
            switch (step.tag) {
              case MAP:
                if (head2) {
                  tail = new Aff2(CONS, head2, tail);
                }
                head2 = new Aff2(MAP, step._1, EMPTY, EMPTY);
                step = step._2;
                break;
              case APPLY:
                if (head2) {
                  tail = new Aff2(CONS, head2, tail);
                }
                head2 = new Aff2(APPLY, EMPTY, step._2, EMPTY);
                step = step._1;
                break;
              case ALT:
                if (head2) {
                  tail = new Aff2(CONS, head2, tail);
                }
                head2 = new Aff2(ALT, EMPTY, step._2, EMPTY);
                step = step._1;
                break;
              default:
                fid = fiberId++;
                status = RETURN;
                tmp = step;
                step = new Aff2(FORKED, fid, new Aff2(CONS, head2, tail), EMPTY);
                tmp = Fiber(util, supervisor, tmp);
                tmp.onComplete({
                  rethrow: false,
                  handler: resolve(step)
                })();
                fibers[fid] = tmp;
                if (supervisor) {
                  supervisor.register(tmp);
                }
            }
            break;
          case RETURN:
            if (head2 === null) {
              break loop;
            }
            if (head2._1 === EMPTY) {
              head2._1 = step;
              status = CONTINUE;
              step = head2._2;
              head2._2 = EMPTY;
            } else {
              head2._2 = step;
              step = head2;
              if (tail === null) {
                head2 = null;
              } else {
                head2 = tail._1;
                tail = tail._2;
              }
            }
        }
      }
      root = step;
      for (fid = 0; fid < fiberId; fid++) {
        fibers[fid].run();
      }
    }
    function cancel(error3, cb2) {
      interrupt = util.left(error3);
      var innerKills;
      for (var kid in kills) {
        if (kills.hasOwnProperty(kid)) {
          innerKills = kills[kid];
          for (kid in innerKills) {
            if (innerKills.hasOwnProperty(kid)) {
              innerKills[kid]();
            }
          }
        }
      }
      kills = null;
      var newKills = kill(error3, root, cb2);
      return function(killError) {
        return new Aff2(ASYNC, function(killCb) {
          return function() {
            for (var kid2 in newKills) {
              if (newKills.hasOwnProperty(kid2)) {
                newKills[kid2]();
              }
            }
            return nonCanceler2;
          };
        });
      };
    }
    run2();
    return function(killError) {
      return new Aff2(ASYNC, function(killCb) {
        return function() {
          return cancel(killError, killCb);
        };
      });
    };
  }
  function sequential(util, supervisor, par) {
    return new Aff2(ASYNC, function(cb) {
      return function() {
        return runPar(util, supervisor, par, cb);
      };
    });
  }
  Aff2.EMPTY = EMPTY;
  Aff2.Pure = AffCtr(PURE);
  Aff2.Throw = AffCtr(THROW);
  Aff2.Catch = AffCtr(CATCH);
  Aff2.Sync = AffCtr(SYNC);
  Aff2.Async = AffCtr(ASYNC);
  Aff2.Bind = AffCtr(BIND);
  Aff2.Bracket = AffCtr(BRACKET);
  Aff2.Fork = AffCtr(FORK);
  Aff2.Seq = AffCtr(SEQ);
  Aff2.ParMap = AffCtr(MAP);
  Aff2.ParApply = AffCtr(APPLY);
  Aff2.ParAlt = AffCtr(ALT);
  Aff2.Fiber = Fiber;
  Aff2.Supervisor = Supervisor;
  Aff2.Scheduler = Scheduler;
  Aff2.nonCanceler = nonCanceler2;
  return Aff2;
})();
var _pure = Aff.Pure;
var _throwError = Aff.Throw;
function _catchError(aff) {
  return function(k) {
    return Aff.Catch(aff, k);
  };
}
function _map(f) {
  return function(aff) {
    if (aff.tag === Aff.Pure.tag) {
      return Aff.Pure(f(aff._1));
    } else {
      return Aff.Bind(aff, function(value) {
        return Aff.Pure(f(value));
      });
    }
  };
}
function _bind(aff) {
  return function(k) {
    return Aff.Bind(aff, k);
  };
}
var _liftEffect = Aff.Sync;
var makeAff = Aff.Async;
function _makeFiber(util, aff) {
  return function() {
    return Aff.Fiber(util, null, aff);
  };
}
var _sequential = Aff.Seq;

// output-es/Effect.Aff/index.js
var functorAff = { map: _map };
var ffiUtil = {
  isLeft: (v) => {
    if (v.tag === "Left") {
      return true;
    }
    if (v.tag === "Right") {
      return false;
    }
    fail();
  },
  fromLeft: (v) => {
    if (v.tag === "Left") {
      return v._1;
    }
    if (v.tag === "Right") {
      return _crashWith("unsafeFromLeft: Right");
    }
    fail();
  },
  fromRight: (v) => {
    if (v.tag === "Right") {
      return v._1;
    }
    if (v.tag === "Left") {
      return _crashWith("unsafeFromRight: Left");
    }
    fail();
  },
  left: Left,
  right: Right
};
var monadAff = { Applicative0: () => applicativeAff, Bind1: () => bindAff };
var bindAff = { bind: _bind, Apply0: () => applyAff };
var applyAff = { apply: (f) => (a) => _bind(f)((f$p) => _bind(a)((a$p) => applicativeAff.pure(f$p(a$p)))), Functor0: () => functorAff };
var applicativeAff = { pure: _pure, Apply0: () => applyAff };
var monadThrowAff = { throwError: _throwError, Monad0: () => monadAff };
var monadErrorAff = { catchError: _catchError, MonadThrow0: () => monadThrowAff };
var $$try2 = /* @__PURE__ */ $$try(monadErrorAff);
var nonCanceler = /* @__PURE__ */ (() => {
  const $0 = _pure();
  return (v) => $0;
})();

// output-es/Effect.Console/foreign.js
var error2 = function(s) {
  return function() {
    console.error(s);
  };
};

// output-es/Gopurs.CodeGen/index.js
var translateExpr = (v) => {
  if (v.tag === "Var") {
    if (v._1._2 === "log") {
      return "gopurs_runtime.Func(func(x gopurs_runtime.Value) gopurs_runtime.Value { return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value { fmt.Println(x.StrVal); return gopurs_runtime.Value{} }) })";
    }
    return replaceAll("$")("_")(v._1._2);
  }
  if (v.tag === "Local") {
    if (v._1.tag === "Just") {
      return replaceAll("$")("_")(v._1._1);
    }
    if (v._1.tag === "Nothing") {
      return "_";
    }
    return "gopurs_runtime.Value{}";
  }
  if (v.tag === "Lit") {
    if (v._1.tag === "LitString") {
      return 'gopurs_runtime.Str("' + v._1._1 + '")';
    }
    return "gopurs_runtime.Value{}";
  }
  if (v.tag === "App") {
    return "gopurs_runtime.Apply(" + translateExpr(v._1) + ", " + translateExpr((() => {
      if (0 < v._2.length) {
        return v._2[0];
      }
      fail();
    })()) + ")";
  }
  return "gopurs_runtime.Value{}";
};
var translateBinding = (v) => {
  const safeName = replaceAll("$")("_")(v._1);
  return $Maybe("Just", (safeName === "main" ? "var Main = " : "var " + safeName + " = ") + translateExpr(v._2));
};
var translateBindingGroup = (bg) => $Maybe("Just", joinWith("\n")(mapMaybe(translateBinding)(bg.bindings)));
var translate = (imports) => (backendMod) => "package " + replaceAll(".")("_")(backendMod.name) + '\n\nimport (\n	"gopurs/output/gopurs_runtime"\n	"fmt"\n)\n\nvar _ = fmt.Println\nvar _ = gopurs_runtime.TypeInt\n\n' + joinWith("\n\n")(mapMaybe(translateBindingGroup)(fromFoldableImpl(
  foldrArray,
  backendMod.bindings
)));

// output-es/Gopurs.Runtime/index.js
var runtimeGoCode = 'package gopurs_runtime\n\nconst (\n	TypeInt = 1\n	TypeString = 2\n	TypeRecord = 3\n	TypeFunc = 4\n	TypeConstructor = 5\n)\n\ntype Value struct {\n	Type   uint8\n	IntVal int64\n	StrVal string\n	PtrVal any\n}\n\nfunc Str(v string) Value {\n	return Value{Type: TypeString, StrVal: v}\n}\n\nfunc Record(m map[string]Value) Value {\n	return Value{Type: TypeRecord, PtrVal: m}\n}\n\nfunc Cons(tag string, args []Value) Value {\n	return Value{Type: TypeConstructor, StrVal: tag, PtrVal: args}\n}\n\n// Function with 1 arg (curried)\nfunc Func(f func(Value) Value) Value {\n	return Value{Type: TypeFunc, PtrVal: f}\n}\n\n// Uncurried application helper\nfunc Apply(f Value, arg Value) Value {\n	if f.Type != TypeFunc {\n		panic("Attempted to apply a non-function")\n	}\n	fn := f.PtrVal.(func(Value) Value)\n	return fn(arg)\n}\n';

// output-es/Node.Encoding/index.js
var $Encoding = (tag) => tag;
var UTF8 = /* @__PURE__ */ $Encoding("UTF8");

// output-es/Data.Nullable/foreign.js
function nullable(a, r, f) {
  return a == null ? r : f(a);
}

// output-es/Data.Semiring/foreign.js
var intAdd = function(x) {
  return function(y) {
    return x + y | 0;
  };
};
var intMul = function(x) {
  return function(y) {
    return x * y | 0;
  };
};

// output-es/Data.Semiring/index.js
var semiringRecordNil = { addRecord: (v) => (v1) => (v2) => ({}), mulRecord: (v) => (v1) => (v2) => ({}), oneRecord: (v) => (v1) => ({}), zeroRecord: (v) => (v1) => ({}) };
var semiringInt = { add: intAdd, zero: 0, mul: intMul, one: 1 };
var semiringRecordCons = (dictIsSymbol) => () => (dictSemiringRecord) => (dictSemiring) => {
  const one1 = dictSemiring.one;
  const zero1 = dictSemiring.zero;
  return {
    addRecord: (v) => (ra) => (rb) => {
      const key = dictIsSymbol.reflectSymbol($$Proxy);
      const $$get = unsafeGet(key);
      return unsafeSet(key)(dictSemiring.add($$get(ra))($$get(rb)))(dictSemiringRecord.addRecord($$Proxy)(ra)(rb));
    },
    mulRecord: (v) => (ra) => (rb) => {
      const key = dictIsSymbol.reflectSymbol($$Proxy);
      const $$get = unsafeGet(key);
      return unsafeSet(key)(dictSemiring.mul($$get(ra))($$get(rb)))(dictSemiringRecord.mulRecord($$Proxy)(ra)(rb));
    },
    oneRecord: (v) => (v1) => unsafeSet(dictIsSymbol.reflectSymbol($$Proxy))(one1)(dictSemiringRecord.oneRecord($$Proxy)($$Proxy)),
    zeroRecord: (v) => (v1) => unsafeSet(dictIsSymbol.reflectSymbol($$Proxy))(zero1)(dictSemiringRecord.zeroRecord($$Proxy)($$Proxy))
  };
};

// output-es/Data.EuclideanRing/foreign.js
var intMod = function(x) {
  return function(y) {
    if (y === 0) return 0;
    var yy = Math.abs(y);
    return (x % yy + yy) % yy;
  };
};

// output-es/Data.Number/foreign.js
var isFiniteImpl = isFinite;
function fromStringImpl(str, isFinite2, just, nothing) {
  var num = parseFloat(str);
  if (isFinite2(num)) {
    return just(num);
  } else {
    return nothing;
  }
}

// output-es/Data.Int/foreign.js
var fromNumberImpl = function(just) {
  return function(nothing) {
    return function(n) {
      return (n | 0) === n ? just(n) : nothing;
    };
  };
};
var fromStringAsImpl = function(just) {
  return function(nothing) {
    return function(radix) {
      var digits;
      if (radix < 11) {
        digits = "[0-" + (radix - 1).toString() + "]";
      } else if (radix === 11) {
        digits = "[0-9a]";
      } else {
        digits = "[0-9a-" + String.fromCharCode(86 + radix) + "]";
      }
      var pattern = new RegExp("^[\\+\\-]?" + digits + "+$", "i");
      return function(s) {
        if (pattern.test(s)) {
          var i = parseInt(s, radix);
          return (i | 0) === i ? just(i) : nothing;
        } else {
          return nothing;
        }
      };
    };
  };
};

// output-es/Data.Int/index.js
var fromStringAs = /* @__PURE__ */ fromStringAsImpl(Just)(Nothing);
var fromString = /* @__PURE__ */ fromStringAs(10);
var fromNumber = /* @__PURE__ */ fromNumberImpl(Just)(Nothing);

// output-es/Node.FS.Constants/foreign.js
import { constants } from "node:fs";
var f_OK = constants.F_OK;
var r_OK = constants.R_OK;
var w_OK = constants.W_OK;
var x_OK = constants.X_OK;
var copyFile_EXCL = constants.COPYFILE_EXCL;
var copyFile_FICLONE = constants.COPYFILE_FICLONE;
var copyFile_FICLONE_FORCE = constants.COPYFILE_FICLONE_FORCE;

// output-es/Data.Bounded/foreign.js
var topChar = String.fromCharCode(65535);
var bottomChar = String.fromCharCode(0);
var topNumber = Number.POSITIVE_INFINITY;
var bottomNumber = Number.NEGATIVE_INFINITY;

// output-es/Data.Enum/foreign.js
function toCharCode(c) {
  return c.charCodeAt(0);
}
function fromCharCode(c) {
  return String.fromCharCode(c);
}

// output-es/Data.Enum/index.js
var charToEnum = (v) => {
  if (v >= 0 && v <= 65535) {
    return $Maybe("Just", fromCharCode(v));
  }
  return Nothing;
};

// output-es/Data.String.Unsafe/foreign.js
var charAt = function(i) {
  return function(s) {
    if (i >= 0 && i < s.length) return s.charAt(i);
    throw new Error("Data.String.Unsafe.charAt: Invalid index.");
  };
};

// output-es/Data.String.CodeUnits/foreign.js
var toCharArray = function(s) {
  return s.split("");
};
var singleton = function(c) {
  return c;
};
var _charAt = function(just) {
  return function(nothing) {
    return function(i) {
      return function(s) {
        return i >= 0 && i < s.length ? just(s.charAt(i)) : nothing;
      };
    };
  };
};
var length2 = function(s) {
  return s.length;
};
var _indexOfStartingAt = function(just) {
  return function(nothing) {
    return function(x) {
      return function(startAt) {
        return function(s) {
          if (startAt < 0 || startAt > s.length) return nothing;
          var i = s.indexOf(x, startAt);
          return i === -1 ? nothing : just(i);
        };
      };
    };
  };
};
var take = function(n) {
  return function(s) {
    return s.substr(0, n);
  };
};
var drop = function(n) {
  return function(s) {
    return s.substring(n);
  };
};
var splitAt = function(i) {
  return function(s) {
    return { before: s.substring(0, i), after: s.substring(i) };
  };
};

// output-es/Data.String.CodeUnits/index.js
var stripPrefix = (v) => (str) => {
  const v1 = splitAt(length2(v))(str);
  if (v1.before === v) {
    return $Maybe("Just", v1.after);
  }
  return Nothing;
};
var indexOf$p = /* @__PURE__ */ _indexOfStartingAt(Just)(Nothing);
var charAt2 = /* @__PURE__ */ _charAt(Just)(Nothing);

// output-es/Data.String.CodePoints/foreign.js
var hasArrayFrom = typeof Array.from === "function";
var hasStringIterator = typeof Symbol !== "undefined" && Symbol != null && typeof Symbol.iterator !== "undefined" && typeof String.prototype[Symbol.iterator] === "function";
var hasFromCodePoint = typeof String.prototype.fromCodePoint === "function";
var hasCodePointAt = typeof String.prototype.codePointAt === "function";
var _unsafeCodePointAt0 = function(fallback) {
  return hasCodePointAt ? function(str) {
    return str.codePointAt(0);
  } : fallback;
};
var _fromCodePointArray = function(singleton3) {
  return hasFromCodePoint ? function(cps) {
    if (cps.length < 1e4) {
      return String.fromCodePoint.apply(String, cps);
    }
    return cps.map(singleton3).join("");
  } : function(cps) {
    return cps.map(singleton3).join("");
  };
};
var _singleton = function(fallback) {
  return hasFromCodePoint ? String.fromCodePoint : fallback;
};
var _take = function(fallback) {
  return function(n) {
    if (hasStringIterator) {
      return function(str) {
        var accum = "";
        var iter = str[Symbol.iterator]();
        for (var i = 0; i < n; ++i) {
          var o = iter.next();
          if (o.done) return accum;
          accum += o.value;
        }
        return accum;
      };
    }
    return fallback(n);
  };
};
var _toCodePointArray = function(fallback) {
  return function(unsafeCodePointAt02) {
    if (hasArrayFrom) {
      return function(str) {
        return Array.from(str, unsafeCodePointAt02);
      };
    }
    return fallback;
  };
};

// output-es/Data.String.CodePoints/index.js
var uncons = (s) => {
  const v = length2(s);
  if (v === 0) {
    return Nothing;
  }
  if (v === 1) {
    return $Maybe("Just", { head: toCharCode(charAt(0)(s)), tail: "" });
  }
  const cu1 = toCharCode(charAt(1)(s));
  const cu0 = toCharCode(charAt(0)(s));
  if (55296 <= cu0 && cu0 <= 56319 && 56320 <= cu1 && cu1 <= 57343) {
    return $Maybe("Just", { head: (((cu0 - 55296 | 0) * 1024 | 0) + (cu1 - 56320 | 0) | 0) + 65536 | 0, tail: drop(2)(s) });
  }
  return $Maybe("Just", { head: cu0, tail: drop(1)(s) });
};
var unconsButWithTuple = (s) => {
  const $0 = uncons(s);
  if ($0.tag === "Just") {
    return $Maybe("Just", $Tuple($0._1.head, $0._1.tail));
  }
  return Nothing;
};
var toCodePointArrayFallback = (s) => unfoldableArray.unfoldr(unconsButWithTuple)(s);
var unsafeCodePointAt0Fallback = (s) => {
  const cu0 = toCharCode(charAt(0)(s));
  if (55296 <= cu0 && cu0 <= 56319 && length2(s) > 1) {
    const cu1 = toCharCode(charAt(1)(s));
    if (56320 <= cu1 && cu1 <= 57343) {
      return (((cu0 - 55296 | 0) * 1024 | 0) + (cu1 - 56320 | 0) | 0) + 65536 | 0;
    }
  }
  return cu0;
};
var unsafeCodePointAt0 = /* @__PURE__ */ _unsafeCodePointAt0(unsafeCodePointAt0Fallback);
var toCodePointArray = /* @__PURE__ */ _toCodePointArray(toCodePointArrayFallback)(unsafeCodePointAt0);
var fromCharCode2 = (x) => singleton((() => {
  if (x >= 0 && x <= 65535) {
    return fromCharCode(x);
  }
  if (x < 0) {
    return "\0";
  }
  return "\uFFFF";
})());
var singletonFallback = (v) => {
  if (v <= 65535) {
    return fromCharCode2(v);
  }
  return fromCharCode2(intDiv(v - 65536 | 0, 1024) + 55296 | 0) + fromCharCode2(intMod(v - 65536 | 0)(1024) + 56320 | 0);
};
var fromCodePointArray = /* @__PURE__ */ _fromCodePointArray(singletonFallback);
var singleton2 = /* @__PURE__ */ _singleton(singletonFallback);
var takeFallback = (v) => (v1) => {
  if (v < 1) {
    return "";
  }
  const v2 = uncons(v1);
  if (v2.tag === "Just") {
    return singleton2(v2._1.head) + takeFallback(v - 1 | 0)(v2._1.tail);
  }
  return v1;
};
var take2 = /* @__PURE__ */ _take(takeFallback);
var eqCodePoint = { eq: (x) => (y) => x === y };
var ordCodePoint = { compare: (x) => (y) => ordInt.compare(x)(y), Eq0: () => eqCodePoint };
var codePointFromChar = (x) => toCharCode(x);
var boundedCodePoint = { bottom: 0, top: 1114111, Ord0: () => ordCodePoint };
var boundedEnumCodePoint = {
  cardinality: 1114112,
  fromEnum: (v) => v,
  toEnum: (n) => {
    if (n >= 0 && n <= 1114111) {
      return $Maybe("Just", n);
    }
    return Nothing;
  },
  Bounded0: () => boundedCodePoint,
  Enum1: () => enumCodePoint
};
var enumCodePoint = {
  succ: (a) => {
    const $0 = a + 1 | 0;
    if ($0 >= 0 && $0 <= 1114111) {
      return $Maybe("Just", $0);
    }
    return Nothing;
  },
  pred: (a) => {
    const $0 = a - 1 | 0;
    if ($0 >= 0 && $0 <= 1114111) {
      return $Maybe("Just", $0);
    }
    return Nothing;
  },
  Ord0: () => ordCodePoint
};

// output-es/Node.FS.Perms/index.js
var semiringPerm = {
  add: (v) => (v1) => ({ r: v.r || v1.r, w: v.w || v1.w, x: v.x || v1.x }),
  zero: { r: false, w: false, x: false },
  mul: (v) => (v1) => ({ r: v.r && v1.r, w: v.w && v1.w, x: v.x && v1.x }),
  one: { r: true, w: true, x: true }
};
var permToString = (x) => showIntImpl(((x.r ? 4 : 0) + (x.w ? 2 : 0) | 0) + (x.x ? 1 : 0) | 0);
var permsToString = (v) => "0" + permToString(v.u) + permToString(v.g) + permToString(v.o);

// output-es/Node.FS.Async/foreign.js
import {
  access,
  copyFile,
  mkdtemp,
  rename,
  truncate,
  chown,
  chmod,
  stat,
  lstat,
  link,
  symlink,
  readlink,
  realpath,
  unlink,
  rmdir,
  rm,
  mkdir,
  readdir,
  utimes,
  readFile,
  writeFile,
  appendFile,
  open,
  read as read2,
  write as write2,
  close
} from "node:fs";

// output-es/Node.FS.Async/index.js
var handleCallback = (cb) => (err, a) => {
  const v = nullable(err, Nothing, Just);
  if (v.tag === "Nothing") {
    return cb($Either("Right", a))();
  }
  if (v.tag === "Just") {
    return cb($Either("Left", v._1))();
  }
  fail();
};
var mkdir$p = (file) => (v) => (cb) => {
  const $0 = { recursive: v.recursive, mode: permsToString(v.mode) };
  return () => mkdir(file, $0, handleCallback(cb));
};
var mkdir2 = (path) => mkdir$p(path)({ recursive: false, mode: { u: semiringPerm.one, g: semiringPerm.one, o: semiringPerm.one } });
var readTextFile = (encoding) => (file) => (cb) => {
  const $0 = {
    encoding: (() => {
      if (encoding === "ASCII") {
        return "ASCII";
      }
      if (encoding === "UTF8") {
        return "UTF8";
      }
      if (encoding === "UTF16LE") {
        return "UTF16LE";
      }
      if (encoding === "UCS2") {
        return "UCS2";
      }
      if (encoding === "Base64") {
        return "Base64";
      }
      if (encoding === "Base64Url") {
        return "Base64Url";
      }
      if (encoding === "Latin1") {
        return "Latin1";
      }
      if (encoding === "Binary") {
        return "Binary";
      }
      if (encoding === "Hex") {
        return "Hex";
      }
      fail();
    })()
  };
  return () => readFile(file, $0, handleCallback(cb));
};
var readdir2 = (file) => (cb) => () => readdir(file, handleCallback(cb));
var stat2 = (file) => (cb) => () => stat(file, handleCallback(cb));
var writeTextFile = (encoding) => (file) => (buff) => (cb) => {
  const $0 = {
    encoding: (() => {
      if (encoding === "ASCII") {
        return "ASCII";
      }
      if (encoding === "UTF8") {
        return "UTF8";
      }
      if (encoding === "UTF16LE") {
        return "UTF16LE";
      }
      if (encoding === "UCS2") {
        return "UCS2";
      }
      if (encoding === "Base64") {
        return "Base64";
      }
      if (encoding === "Base64Url") {
        return "Base64Url";
      }
      if (encoding === "Latin1") {
        return "Latin1";
      }
      if (encoding === "Binary") {
        return "Binary";
      }
      if (encoding === "Hex") {
        return "Hex";
      }
      fail();
    })()
  };
  return () => writeFile(file, buff, $0, handleCallback(cb));
};

// output-es/Node.FS.Aff/index.js
var toAff1 = (f) => (a) => {
  const $0 = f(a);
  return makeAff((k) => {
    const $1 = $0(k);
    return () => {
      $1();
      return nonCanceler;
    };
  });
};
var toAff2 = (f) => (a) => (b) => {
  const $0 = f(a)(b);
  return makeAff((k) => {
    const $1 = $0(k);
    return () => {
      $1();
      return nonCanceler;
    };
  });
};
var toAff3 = (f) => (a) => (b) => (c) => {
  const $0 = f(a)(b)(c);
  return makeAff((k) => {
    const $1 = $0(k);
    return () => {
      $1();
      return nonCanceler;
    };
  });
};

// output-es/Node.FS.Stats/foreign.js
var isDirectoryImpl = (s) => s.isDirectory();

// output-es/Data.List/index.js
var foldM = (dictMonad) => (v) => (v1) => (v2) => {
  if (v2.tag === "Nil") {
    return dictMonad.Applicative0().pure(v1);
  }
  if (v2.tag === "Cons") {
    const $0 = v2._2;
    return dictMonad.Bind1().bind(v(v1)(v2._1))((b$p) => foldM(dictMonad)(v)(b$p)($0));
  }
  fail();
};

// output-es/Data.Semigroup/foreign.js
var concatArray = function(xs) {
  return function(ys) {
    if (xs.length === 0) return ys;
    if (ys.length === 0) return xs;
    return xs.concat(ys);
  };
};

// output-es/Data.Semigroup/index.js
var semigroupArray = { append: concatArray };

// output-es/Data.Semigroup.Foldable/index.js
var maximum = (dictOrd) => {
  const semigroupMax = {
    append: (v) => (v1) => {
      const v$1 = dictOrd.compare(v)(v1);
      if (v$1 === "LT") {
        return v1;
      }
      if (v$1 === "EQ") {
        return v;
      }
      if (v$1 === "GT") {
        return v;
      }
      fail();
    }
  };
  return (dictFoldable1) => dictFoldable1.foldMap1(semigroupMax)(unsafeCoerce);
};

// output-es/Data.TraversableWithIndex/index.js
var traversableWithIndexArray = {
  traverseWithIndex: (dictApplicative) => {
    const sequence1 = traversableWithIndexArray.Traversable2().sequence(dictApplicative);
    return (f) => {
      const $0 = traversableWithIndexArray.FunctorWithIndex0().mapWithIndex(f);
      return (x) => sequence1($0(x));
    };
  },
  FunctorWithIndex0: () => functorWithIndexArray,
  FoldableWithIndex1: () => foldableWithIndexArray,
  Traversable2: () => traversableArray
};

// output-es/Data.Array.NonEmpty.Internal/foreign.js
var foldr1Impl = function(f, xs) {
  var acc = xs[xs.length - 1];
  for (var i = xs.length - 2; i >= 0; i--) {
    acc = f(xs[i])(acc);
  }
  return acc;
};
var foldl1Impl = function(f, xs) {
  var acc = xs[0];
  var len = xs.length;
  for (var i = 1; i < len; i++) {
    acc = f(acc)(xs[i]);
  }
  return acc;
};

// output-es/Data.Array.NonEmpty.Internal/index.js
var foldable1NonEmptyArray = {
  foldMap1: (dictSemigroup) => {
    const append = dictSemigroup.append;
    return (f) => {
      const $0 = arrayMap(f);
      const $1 = foldable1NonEmptyArray.foldl1(append);
      return (x) => $1($0(x));
    };
  },
  foldr1: ($0) => ($1) => foldr1Impl($0, $1),
  foldl1: ($0) => ($1) => foldl1Impl($0, $1),
  Foldable0: () => foldableArray
};

// output-es/Data.Array.NonEmpty/index.js
var head = (x) => {
  if (0 < x.length) {
    return x[0];
  }
  fail();
};
var uncons2 = (x) => {
  const $0 = unconsImpl((v) => Nothing, (x$1) => (xs) => $Maybe("Just", { head: x$1, tail: xs }), x);
  if ($0.tag === "Just") {
    return $0._1;
  }
  fail();
};

// output-es/Data.Map/index.js
var semigroupSemigroupMap = (dictOrd) => {
  const compare4 = dictOrd.compare;
  return (dictSemigroup) => {
    const append = dictSemigroup.append;
    return { append: (v) => (v1) => unsafeUnionWith(compare4, append, v, v1) };
  };
};
var monoidSemigroupMap = (dictOrd) => {
  const semigroupSemigroupMap1 = semigroupSemigroupMap(dictOrd);
  return (dictSemigroup) => {
    const semigroupSemigroupMap2 = semigroupSemigroupMap1(dictSemigroup);
    return { mempty: Leaf, Semigroup0: () => semigroupSemigroupMap2 };
  };
};

// output-es/Data.Monoid/index.js
var monoidArray = { mempty: [], Semigroup0: () => semigroupArray };
var power = (dictMonoid) => {
  const mempty1 = dictMonoid.mempty;
  const $0 = dictMonoid.Semigroup0();
  return (x) => {
    const go = (p) => {
      if (p <= 0) {
        return mempty1;
      }
      if (p === 1) {
        return x;
      }
      if (intMod(p)(2) === 0) {
        const x$p2 = go(intDiv(p, 2));
        return $0.append(x$p2)(x$p2);
      }
      const x$p = go(intDiv(p, 2));
      return $0.append(x$p)($0.append(x$p)(x));
    };
    return go;
  };
};

// output-es/Data.Semigroup.First/index.js
var semigroupFirst2 = { append: (x) => (v) => x };

// output-es/Data.Set/index.js
var foldableSet = {
  foldMap: (dictMonoid) => {
    const foldMap13 = foldableList.foldMap(dictMonoid);
    return (f) => {
      const $0 = foldMap13(f);
      return (x) => $0((() => {
        const go = (m$p, z$p) => {
          if (m$p.tag === "Leaf") {
            return z$p;
          }
          if (m$p.tag === "Node") {
            return go(m$p._5, $List("Cons", m$p._3, go(m$p._6, z$p)));
          }
          fail();
        };
        return go(x, Nil);
      })());
    };
  },
  foldl: (f) => (x) => {
    const go = (go$a0$copy) => (go$a1$copy) => {
      let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
      while (go$c) {
        const b = go$a0, v = go$a1;
        if (v.tag === "Nil") {
          go$c = false;
          go$r = b;
          continue;
        }
        if (v.tag === "Cons") {
          go$a0 = f(b)(v._1);
          go$a1 = v._2;
          continue;
        }
        fail();
      }
      return go$r;
    };
    const $0 = go(x);
    return (x$1) => $0((() => {
      const go$1 = (m$p, z$p) => {
        if (m$p.tag === "Leaf") {
          return z$p;
        }
        if (m$p.tag === "Node") {
          return go$1(m$p._5, $List("Cons", m$p._3, go$1(m$p._6, z$p)));
        }
        fail();
      };
      return go$1(x$1, Nil);
    })());
  },
  foldr: (f) => (x) => {
    const $0 = foldableList.foldr(f)(x);
    return (x$1) => $0((() => {
      const go = (m$p, z$p) => {
        if (m$p.tag === "Leaf") {
          return z$p;
        }
        if (m$p.tag === "Node") {
          return go(m$p._5, $List("Cons", m$p._3, go(m$p._6, z$p)));
        }
        fail();
      };
      return go(x$1, Nil);
    })());
  }
};
var mapMaybe2 = (dictOrd) => (f) => foldableSet.foldr((a) => (acc) => {
  const $0 = f(a);
  if ($0.tag === "Nothing") {
    return acc;
  }
  if ($0.tag === "Just") {
    return insert(dictOrd)($0._1)()(acc);
  }
  fail();
})(Leaf);
var monoidSet = (dictOrd) => {
  const semigroupSet1 = {
    append: (() => {
      const compare4 = dictOrd.compare;
      return (m1) => (m2) => unsafeUnionWith(compare4, $$const, m1, m2);
    })()
  };
  return { mempty: Leaf, Semigroup0: () => semigroupSet1 };
};

// output-es/PureScript.Backend.Optimizer.CoreFn/index.js
var $Bind = (tag, _1) => ({ tag, _1 });
var $Binder = (tag, _1, _2, _3, _4) => ({ tag, _1, _2, _3, _4 });
var $Binding = (_1, _2, _3) => ({ tag: "Binding", _1, _2, _3 });
var $CaseAlternative = (_1, _2) => ({ tag: "CaseAlternative", _1, _2 });
var $CaseGuard = (tag, _1) => ({ tag, _1 });
var $Comment = (tag, _1) => ({ tag, _1 });
var $ConstructorType = (tag) => tag;
var $Expr = (tag, _1, _2, _3, _4) => ({ tag, _1, _2, _3, _4 });
var $Guard = (_1, _2) => ({ tag: "Guard", _1, _2 });
var $Import = (_1, _2) => ({ tag: "Import", _1, _2 });
var $Literal = (tag, _1) => ({ tag, _1 });
var $Meta = (tag, _1, _2) => ({ tag, _1, _2 });
var $Prop = (_1, _2) => ({ tag: "Prop", _1, _2 });
var $Qualified = (_1, _2) => ({ tag: "Qualified", _1, _2 });
var $ReExport = (_1, _2) => ({ tag: "ReExport", _1, _2 });
var zero = /* @__PURE__ */ (() => semiringRecordCons({ reflectSymbol: () => "column" })()(semiringRecordCons({ reflectSymbol: () => "line" })()(semiringRecordNil)(semiringInt))(semiringInt).zeroRecord($$Proxy)($$Proxy))();
var Prop = (value0) => (value1) => $Prop(value0, value1);
var LitArray = (value0) => $Literal("LitArray", value0);
var LitRecord = (value0) => $Literal("LitRecord", value0);
var ReExport = (value0) => (value1) => $ReExport(value0, value1);
var ProductType = /* @__PURE__ */ $ConstructorType("ProductType");
var SumType = /* @__PURE__ */ $ConstructorType("SumType");
var IsNewtype = /* @__PURE__ */ $Meta("IsNewtype");
var IsTypeClassConstructor = /* @__PURE__ */ $Meta("IsTypeClassConstructor");
var IsForeign = /* @__PURE__ */ $Meta("IsForeign");
var IsWhere = /* @__PURE__ */ $Meta("IsWhere");
var IsSyntheticApp = /* @__PURE__ */ $Meta("IsSyntheticApp");
var BinderVar = (value0) => (value1) => $Binder("BinderVar", value0, value1);
var BinderLit = (value0) => (value1) => $Binder("BinderLit", value0, value1);
var ExprVar = (value0) => (value1) => $Expr("ExprVar", value0, value1);
var ExprLit = (value0) => (value1) => $Expr("ExprLit", value0, value1);
var eqLiteral = (dictEq) => {
  const eq11 = eqArrayImpl((x) => (y) => x._1 === y._1 && dictEq.eq(x._2)(y._2));
  return {
    eq: (x) => (y) => {
      if (x.tag === "LitInt") {
        return y.tag === "LitInt" && x._1 === y._1;
      }
      if (x.tag === "LitNumber") {
        return y.tag === "LitNumber" && x._1 === y._1;
      }
      if (x.tag === "LitString") {
        return y.tag === "LitString" && x._1 === y._1;
      }
      if (x.tag === "LitChar") {
        return y.tag === "LitChar" && x._1 === y._1;
      }
      if (x.tag === "LitBoolean") {
        return y.tag === "LitBoolean" && x._1 === y._1;
      }
      if (x.tag === "LitArray") {
        return y.tag === "LitArray" && eqArrayImpl(dictEq.eq)(x._1)(y._1);
      }
      return x.tag === "LitRecord" && y.tag === "LitRecord" && eq11(x._1)(y._1);
    }
  };
};
var propKey = (v) => v._1;
var compare = (x) => (y) => {
  if (x.tag === "Nothing") {
    if (y.tag === "Nothing") {
      return EQ;
    }
    return LT;
  }
  if (y.tag === "Nothing") {
    return GT;
  }
  if (x.tag === "Just" && y.tag === "Just") {
    return ordString.compare(x._1)(y._1);
  }
  fail();
};
var functorProp = { map: (f) => (m) => $Prop(m._1, f(m._2)) };
var functorLiteral = {
  map: (f) => (m) => {
    if (m.tag === "LitInt") {
      return $Literal("LitInt", m._1);
    }
    if (m.tag === "LitNumber") {
      return $Literal("LitNumber", m._1);
    }
    if (m.tag === "LitString") {
      return $Literal("LitString", m._1);
    }
    if (m.tag === "LitChar") {
      return $Literal("LitChar", m._1);
    }
    if (m.tag === "LitBoolean") {
      return $Literal("LitBoolean", m._1);
    }
    if (m.tag === "LitArray") {
      return $Literal("LitArray", arrayMap(f)(m._1));
    }
    if (m.tag === "LitRecord") {
      return $Literal("LitRecord", arrayMap((m$1) => $Prop(m$1._1, f(m$1._2)))(m._1));
    }
    fail();
  }
};
var foldableProp = { foldl: (k) => (a) => (v) => k(a)(v._2), foldr: (k) => (b) => (v) => k(v._2)(b), foldMap: (dictMonoid) => (k) => (v) => k(v._2) };
var traversableProp = {
  traverse: (dictApplicative) => (k) => (v) => dictApplicative.Apply0().Functor0().map(Prop(v._1))(k(v._2)),
  sequence: (dictApplicative) => (v) => dictApplicative.Apply0().Functor0().map(Prop(v._1))(v._2),
  Functor0: () => functorProp,
  Foldable1: () => foldableProp
};
var foldableLiteral = {
  foldl: (k) => foldlDefault(foldableLiteral)(k),
  foldr: (k) => foldrDefault(foldableLiteral)(k),
  foldMap: (dictMonoid) => {
    const foldMap23 = foldableArray.foldMap(dictMonoid);
    const mempty = dictMonoid.mempty;
    return (k) => (v) => {
      if (v.tag === "LitArray") {
        return foldMap23(k)(v._1);
      }
      if (v.tag === "LitRecord") {
        return foldMap23((v$1) => k(v$1._2))(v._1);
      }
      return mempty;
    };
  }
};
var traversableLiteral = {
  traverse: (dictApplicative) => {
    const $0 = dictApplicative.Apply0().Functor0();
    const traverse22 = traversableArray.traverse(dictApplicative);
    return (k) => (v) => {
      if (v.tag === "LitArray") {
        return $0.map(LitArray)(traverse22(k)(v._1));
      }
      if (v.tag === "LitRecord") {
        return $0.map(LitRecord)(traverse22(traversableProp.traverse(dictApplicative)(k))(v._1));
      }
      if (v.tag === "LitInt") {
        return dictApplicative.pure($Literal("LitInt", v._1));
      }
      if (v.tag === "LitNumber") {
        return dictApplicative.pure($Literal("LitNumber", v._1));
      }
      if (v.tag === "LitString") {
        return dictApplicative.pure($Literal("LitString", v._1));
      }
      if (v.tag === "LitChar") {
        return dictApplicative.pure($Literal("LitChar", v._1));
      }
      if (v.tag === "LitBoolean") {
        return dictApplicative.pure($Literal("LitBoolean", v._1));
      }
      fail();
    };
  },
  sequence: (dictApplicative) => (a) => traversableLiteral.traverse(dictApplicative)(identity3)(a),
  Functor0: () => functorLiteral,
  Foldable1: () => foldableLiteral
};
var ordQualified = (dictOrd) => {
  const $0 = dictOrd.Eq0();
  const eqQualified1 = { eq: (x) => (y) => (x._1.tag === "Nothing" ? y._1.tag === "Nothing" : x._1.tag === "Just" && y._1.tag === "Just" && x._1._1 === y._1._1) && $0.eq(x._2)(y._2) };
  return {
    compare: (x) => (y) => {
      const v = compare(x._1)(y._1);
      if (v === "LT") {
        return LT;
      }
      if (v === "GT") {
        return GT;
      }
      return dictOrd.compare(x._2)(y._2);
    },
    Eq0: () => eqQualified1
  };
};
var eqReExport = { eq: (x) => (y) => x._1 === y._1 && x._2 === y._2 };
var ordReExport = {
  compare: (x) => (y) => {
    const v = ordString.compare(x._1)(y._1);
    if (v === "LT") {
      return LT;
    }
    if (v === "GT") {
      return GT;
    }
    return ordString.compare(x._2)(y._2);
  },
  Eq0: () => eqReExport
};
var emptySpan = { path: "<internal>", start: zero, end: zero };

// output-es/PureScript.Backend.Optimizer.Syntax/index.js
var $BackendAccessor = (tag, _1, _2, _3, _4, _5, _6) => ({ tag, _1, _2, _3, _4, _5, _6 });
var $BackendEffect = (tag, _1, _2) => ({ tag, _1, _2 });
var $BackendOperator = (tag, _1, _2, _3) => ({ tag, _1, _2, _3 });
var $BackendOperator1 = (tag, _1) => ({ tag, _1 });
var $BackendOperator2 = (tag, _1) => ({ tag, _1 });
var $BackendOperatorOrd = (tag) => tag;
var $BackendSyntax = (tag, _1, _2, _3, _4, _5) => ({ tag, _1, _2, _3, _4, _5 });
var $Pair = (_1, _2) => ({ tag: "Pair", _1, _2 });
var eq7 = /* @__PURE__ */ eqArrayImpl(eqStringImpl);
var Pair = (value0) => (value1) => $Pair(value0, value1);
var OpEq = /* @__PURE__ */ $BackendOperatorOrd("OpEq");
var OpNotEq = /* @__PURE__ */ $BackendOperatorOrd("OpNotEq");
var OpGt = /* @__PURE__ */ $BackendOperatorOrd("OpGt");
var OpGte = /* @__PURE__ */ $BackendOperatorOrd("OpGte");
var OpLt = /* @__PURE__ */ $BackendOperatorOrd("OpLt");
var OpLte = /* @__PURE__ */ $BackendOperatorOrd("OpLte");
var OpBooleanAnd = /* @__PURE__ */ $BackendOperator2("OpBooleanAnd");
var OpBooleanOr = /* @__PURE__ */ $BackendOperator2("OpBooleanOr");
var OpBooleanNot = /* @__PURE__ */ $BackendOperator1("OpBooleanNot");
var OpIntNegate = /* @__PURE__ */ $BackendOperator1("OpIntNegate");
var OpNumberNegate = /* @__PURE__ */ $BackendOperator1("OpNumberNegate");
var OpArrayLength = /* @__PURE__ */ $BackendOperator1("OpArrayLength");
var Op1 = (value0) => (value1) => $BackendOperator("Op1", value0, value1);
var Op2 = (value0) => (value1) => (value2) => $BackendOperator("Op2", value0, value1, value2);
var EffectRefNew = (value0) => $BackendEffect("EffectRefNew", value0);
var EffectRefRead = (value0) => $BackendEffect("EffectRefRead", value0);
var EffectRefWrite = (value0) => (value1) => $BackendEffect("EffectRefWrite", value0, value1);
var App2 = (value0) => (value1) => $BackendSyntax("App", value0, value1);
var Abs = (value0) => (value1) => $BackendSyntax("Abs", value0, value1);
var UncurriedApp = (value0) => (value1) => $BackendSyntax("UncurriedApp", value0, value1);
var UncurriedAbs = (value0) => (value1) => $BackendSyntax("UncurriedAbs", value0, value1);
var UncurriedEffectApp = (value0) => (value1) => $BackendSyntax("UncurriedEffectApp", value0, value1);
var UncurriedEffectAbs = (value0) => (value1) => $BackendSyntax("UncurriedEffectAbs", value0, value1);
var Update = (value0) => (value1) => $BackendSyntax("Update", value0, value1);
var CtorSaturated = (value0) => (value1) => (value2) => (value3) => (value4) => $BackendSyntax("CtorSaturated", value0, value1, value2, value3, value4);
var LetRec = (value0) => (value1) => (value2) => $BackendSyntax("LetRec", value0, value1, value2);
var Let = (value0) => (value1) => (value2) => (value3) => $BackendSyntax("Let", value0, value1, value2, value3);
var EffectBind = (value0) => (value1) => (value2) => (value3) => $BackendSyntax("EffectBind", value0, value1, value2, value3);
var EffectPure = (value0) => $BackendSyntax("EffectPure", value0);
var EffectDefer = (value0) => $BackendSyntax("EffectDefer", value0);
var Branch = (value0) => (value1) => $BackendSyntax("Branch", value0, value1);
var PrimOp = (value0) => $BackendSyntax("PrimOp", value0);
var PrimEffect = (value0) => $BackendSyntax("PrimEffect", value0);
var PrimUndefined = /* @__PURE__ */ $BackendSyntax("PrimUndefined");
var functorPair = { map: (f) => (m) => $Pair(f(m._1), f(m._2)) };
var functorBackendOperator = {
  map: (f) => (m) => {
    if (m.tag === "Op1") {
      return $BackendOperator("Op1", m._1, f(m._2));
    }
    if (m.tag === "Op2") {
      return $BackendOperator("Op2", m._1, f(m._2), f(m._3));
    }
    fail();
  }
};
var functorBackendEffect = {
  map: (f) => (m) => {
    if (m.tag === "EffectRefNew") {
      return $BackendEffect("EffectRefNew", f(m._1));
    }
    if (m.tag === "EffectRefRead") {
      return $BackendEffect("EffectRefRead", f(m._1));
    }
    if (m.tag === "EffectRefWrite") {
      return $BackendEffect("EffectRefWrite", f(m._1), f(m._2));
    }
    fail();
  }
};
var functorBackendSyntax = {
  map: (f) => (m) => {
    if (m.tag === "Var") {
      return $BackendSyntax("Var", m._1);
    }
    if (m.tag === "Local") {
      return $BackendSyntax("Local", m._1, m._2);
    }
    if (m.tag === "Lit") {
      return $BackendSyntax("Lit", functorLiteral.map(f)(m._1));
    }
    if (m.tag === "App") {
      return $BackendSyntax("App", f(m._1), arrayMap(f)(m._2));
    }
    if (m.tag === "Abs") {
      return $BackendSyntax("Abs", m._1, f(m._2));
    }
    if (m.tag === "UncurriedApp") {
      return $BackendSyntax("UncurriedApp", f(m._1), arrayMap(f)(m._2));
    }
    if (m.tag === "UncurriedAbs") {
      return $BackendSyntax("UncurriedAbs", m._1, f(m._2));
    }
    if (m.tag === "UncurriedEffectApp") {
      return $BackendSyntax("UncurriedEffectApp", f(m._1), arrayMap(f)(m._2));
    }
    if (m.tag === "UncurriedEffectAbs") {
      return $BackendSyntax("UncurriedEffectAbs", m._1, f(m._2));
    }
    if (m.tag === "Accessor") {
      return $BackendSyntax("Accessor", f(m._1), m._2);
    }
    if (m.tag === "Update") {
      return $BackendSyntax("Update", f(m._1), arrayMap((m$1) => $Prop(m$1._1, f(m$1._2)))(m._2));
    }
    if (m.tag === "CtorSaturated") {
      return $BackendSyntax("CtorSaturated", m._1, m._2, m._3, m._4, arrayMap((m$1) => $Tuple(m$1._1, f(m$1._2)))(m._5));
    }
    if (m.tag === "CtorDef") {
      return $BackendSyntax("CtorDef", m._1, m._2, m._3, m._4);
    }
    if (m.tag === "LetRec") {
      return $BackendSyntax("LetRec", m._1, arrayMap((m$1) => $Tuple(m$1._1, f(m$1._2)))(m._2), f(m._3));
    }
    if (m.tag === "Let") {
      return $BackendSyntax("Let", m._1, m._2, f(m._3), f(m._4));
    }
    if (m.tag === "EffectBind") {
      return $BackendSyntax("EffectBind", m._1, m._2, f(m._3), f(m._4));
    }
    if (m.tag === "EffectPure") {
      return $BackendSyntax("EffectPure", f(m._1));
    }
    if (m.tag === "EffectDefer") {
      return $BackendSyntax("EffectDefer", f(m._1));
    }
    if (m.tag === "Branch") {
      return $BackendSyntax("Branch", arrayMap((m$1) => $Pair(f(m$1._1), f(m$1._2)))(m._1), f(m._2));
    }
    if (m.tag === "PrimOp") {
      return $BackendSyntax(
        "PrimOp",
        (() => {
          if (m._1.tag === "Op1") {
            return $BackendOperator("Op1", m._1._1, f(m._1._2));
          }
          if (m._1.tag === "Op2") {
            return $BackendOperator("Op2", m._1._1, f(m._1._2), f(m._1._3));
          }
          fail();
        })()
      );
    }
    if (m.tag === "PrimEffect") {
      return $BackendSyntax(
        "PrimEffect",
        (() => {
          if (m._1.tag === "EffectRefNew") {
            return $BackendEffect("EffectRefNew", f(m._1._1));
          }
          if (m._1.tag === "EffectRefRead") {
            return $BackendEffect("EffectRefRead", f(m._1._1));
          }
          if (m._1.tag === "EffectRefWrite") {
            return $BackendEffect("EffectRefWrite", f(m._1._1), f(m._1._2));
          }
          fail();
        })()
      );
    }
    if (m.tag === "PrimUndefined") {
      return PrimUndefined;
    }
    if (m.tag === "Fail") {
      return $BackendSyntax("Fail", m._1);
    }
    fail();
  }
};
var foldablePair = {
  foldl: (f) => (acc) => (v) => f(f(acc)(v._1))(v._2),
  foldr: (f) => (acc) => (v) => f(v._1)(f(v._2)(acc)),
  foldMap: (dictMonoid) => (f) => (v) => dictMonoid.Semigroup0().append(f(v._1))(f(v._2))
};
var traversablePair = {
  sequence: (dictApplicative) => (a) => traversablePair.traverse(dictApplicative)(identity3)(a),
  traverse: (dictApplicative) => {
    const Apply0 = dictApplicative.Apply0();
    return (f) => (v) => Apply0.apply(Apply0.Functor0().map(Pair)(f(v._1)))(f(v._2));
  },
  Functor0: () => functorPair,
  Foldable1: () => foldablePair
};
var foldableBackendOperator = {
  foldr: (a) => foldrDefault(foldableBackendOperator)(a),
  foldl: (a) => foldlDefault(foldableBackendOperator)(a),
  foldMap: (dictMonoid) => (f) => (v) => {
    if (v.tag === "Op1") {
      return f(v._2);
    }
    if (v.tag === "Op2") {
      return dictMonoid.Semigroup0().append(f(v._2))(f(v._3));
    }
    fail();
  }
};
var traversableBackendOperato = {
  sequence: (dictApplicative) => (a) => traversableBackendOperato.traverse(dictApplicative)(identity3)(a),
  traverse: (dictApplicative) => {
    const Apply0 = dictApplicative.Apply0();
    const $0 = Apply0.Functor0();
    return (f) => (v) => {
      if (v.tag === "Op1") {
        return $0.map(Op1(v._1))(f(v._2));
      }
      if (v.tag === "Op2") {
        return Apply0.apply($0.map(Op2(v._1))(f(v._2)))(f(v._3));
      }
      fail();
    };
  },
  Functor0: () => functorBackendOperator,
  Foldable1: () => foldableBackendOperator
};
var foldableBackendEffect = {
  foldr: (a) => foldrDefault(foldableBackendEffect)(a),
  foldl: (a) => foldlDefault(foldableBackendEffect)(a),
  foldMap: (dictMonoid) => (f) => (v) => {
    if (v.tag === "EffectRefNew") {
      return f(v._1);
    }
    if (v.tag === "EffectRefRead") {
      return f(v._1);
    }
    if (v.tag === "EffectRefWrite") {
      return dictMonoid.Semigroup0().append(f(v._1))(f(v._2));
    }
    fail();
  }
};
var foldableBackendSyntax = {
  foldr: (a) => foldrDefault(foldableBackendSyntax)(a),
  foldl: (a) => foldlDefault(foldableBackendSyntax)(a),
  foldMap: (dictMonoid) => {
    const mempty = dictMonoid.mempty;
    const foldMap7 = foldableArray.foldMap(dictMonoid);
    const $0 = dictMonoid.Semigroup0();
    const foldMap9 = foldableArray.foldMap(dictMonoid);
    return (f) => (v) => {
      if (v.tag === "Var") {
        return mempty;
      }
      if (v.tag === "Local") {
        return mempty;
      }
      if (v.tag === "Lit") {
        if (v._1.tag === "LitArray") {
          return foldMap7(f)(v._1._1);
        }
        if (v._1.tag === "LitRecord") {
          return foldMap7((v$1) => f(v$1._2))(v._1._1);
        }
        return mempty;
      }
      if (v.tag === "App") {
        return $0.append(f(v._1))(foldMap9(f)(v._2));
      }
      if (v.tag === "Abs") {
        return f(v._2);
      }
      if (v.tag === "UncurriedApp") {
        return $0.append(f(v._1))(foldMap7(f)(v._2));
      }
      if (v.tag === "UncurriedAbs") {
        return f(v._2);
      }
      if (v.tag === "UncurriedEffectApp") {
        return $0.append(f(v._1))(foldMap7(f)(v._2));
      }
      if (v.tag === "UncurriedEffectAbs") {
        return f(v._2);
      }
      if (v.tag === "Accessor") {
        return f(v._1);
      }
      if (v.tag === "Update") {
        return $0.append(f(v._1))(foldMap7((v$1) => f(v$1._2))(v._2));
      }
      if (v.tag === "LetRec") {
        return $0.append(foldMap9((v$1) => f(v$1._2))(v._2))(f(v._3));
      }
      if (v.tag === "Let") {
        return $0.append(f(v._3))(f(v._4));
      }
      if (v.tag === "EffectBind") {
        return $0.append(f(v._3))(f(v._4));
      }
      if (v.tag === "EffectPure") {
        return f(v._1);
      }
      if (v.tag === "EffectDefer") {
        return f(v._1);
      }
      if (v.tag === "Branch") {
        return $0.append(foldMap9((v$1) => dictMonoid.Semigroup0().append(f(v$1._1))(f(v$1._2)))(v._1))(f(v._2));
      }
      if (v.tag === "PrimOp") {
        if (v._1.tag === "Op1") {
          return f(v._1._2);
        }
        if (v._1.tag === "Op2") {
          return dictMonoid.Semigroup0().append(f(v._1._2))(f(v._1._3));
        }
        fail();
      }
      if (v.tag === "PrimEffect") {
        if (v._1.tag === "EffectRefNew") {
          return f(v._1._1);
        }
        if (v._1.tag === "EffectRefRead") {
          return f(v._1._1);
        }
        if (v._1.tag === "EffectRefWrite") {
          return dictMonoid.Semigroup0().append(f(v._1._1))(f(v._1._2));
        }
        fail();
      }
      if (v.tag === "PrimUndefined") {
        return mempty;
      }
      if (v.tag === "CtorSaturated") {
        return foldMap7((v$1) => f(v$1._2))(v._5);
      }
      if (v.tag === "CtorDef") {
        return mempty;
      }
      if (v.tag === "Fail") {
        return mempty;
      }
      fail();
    };
  }
};
var traversableBackendEffect = {
  sequence: (dictApplicative) => (a) => traversableBackendEffect.traverse(dictApplicative)(identity3)(a),
  traverse: (dictApplicative) => {
    const Apply0 = dictApplicative.Apply0();
    const $0 = Apply0.Functor0();
    return (f) => (v) => {
      if (v.tag === "EffectRefNew") {
        return $0.map(EffectRefNew)(f(v._1));
      }
      if (v.tag === "EffectRefRead") {
        return $0.map(EffectRefRead)(f(v._1));
      }
      if (v.tag === "EffectRefWrite") {
        return Apply0.apply($0.map(EffectRefWrite)(f(v._1)))(f(v._2));
      }
      fail();
    };
  },
  Functor0: () => functorBackendEffect,
  Foldable1: () => foldableBackendEffect
};
var traversableBackendSyntax = {
  sequence: (dictApplicative) => (a) => traversableBackendSyntax.traverse(dictApplicative)(identity3)(a),
  traverse: (dictApplicative) => {
    const Apply0 = dictApplicative.Apply0();
    const $0 = Apply0.Functor0();
    const traverse7 = traversableArray.traverse(dictApplicative);
    const traverse9 = traversableArray.traverse(dictApplicative);
    const traverse11 = traversablePair.traverse(dictApplicative);
    const traverse12 = traversableBackendOperato.traverse(dictApplicative);
    const traverse13 = traversableBackendEffect.traverse(dictApplicative);
    return (f) => (v) => {
      if (v.tag === "Var") {
        return dictApplicative.pure($BackendSyntax("Var", v._1));
      }
      if (v.tag === "Local") {
        return dictApplicative.pure($BackendSyntax("Local", v._1, v._2));
      }
      if (v.tag === "Lit") {
        if (v._1.tag === "LitInt") {
          return dictApplicative.pure($BackendSyntax("Lit", $Literal("LitInt", v._1._1)));
        }
        if (v._1.tag === "LitNumber") {
          return dictApplicative.pure($BackendSyntax("Lit", $Literal("LitNumber", v._1._1)));
        }
        if (v._1.tag === "LitString") {
          return dictApplicative.pure($BackendSyntax("Lit", $Literal("LitString", v._1._1)));
        }
        if (v._1.tag === "LitChar") {
          return dictApplicative.pure($BackendSyntax("Lit", $Literal("LitChar", v._1._1)));
        }
        if (v._1.tag === "LitBoolean") {
          return dictApplicative.pure($BackendSyntax("Lit", $Literal("LitBoolean", v._1._1)));
        }
        if (v._1.tag === "LitArray") {
          return $0.map((x) => $BackendSyntax("Lit", $Literal("LitArray", x)))(traverse7(f)(v._1._1));
        }
        if (v._1.tag === "LitRecord") {
          return $0.map((x) => $BackendSyntax("Lit", $Literal("LitRecord", x)))(traverse7(traversableProp.traverse(dictApplicative)(f))(v._1._1));
        }
        fail();
      }
      if (v.tag === "App") {
        return Apply0.apply($0.map(App2)(f(v._1)))(traverse9(f)(v._2));
      }
      if (v.tag === "Abs") {
        return $0.map(Abs(v._1))(f(v._2));
      }
      if (v.tag === "UncurriedApp") {
        return Apply0.apply($0.map(UncurriedApp)(f(v._1)))(traverse7(f)(v._2));
      }
      if (v.tag === "UncurriedAbs") {
        return $0.map(UncurriedAbs(v._1))(f(v._2));
      }
      if (v.tag === "UncurriedEffectApp") {
        return Apply0.apply($0.map(UncurriedEffectApp)(f(v._1)))(traverse7(f)(v._2));
      }
      if (v.tag === "UncurriedEffectAbs") {
        return $0.map(UncurriedEffectAbs(v._1))(f(v._2));
      }
      if (v.tag === "Accessor") {
        const $1 = v._2;
        return $0.map((a) => $BackendSyntax("Accessor", a, $1))(f(v._1));
      }
      if (v.tag === "Update") {
        return Apply0.apply($0.map(Update)(f(v._1)))(traverse7(traversableProp.traverse(dictApplicative)(f))(v._2));
      }
      if (v.tag === "CtorDef") {
        return dictApplicative.pure($BackendSyntax("CtorDef", v._1, v._2, v._3, v._4));
      }
      if (v.tag === "CtorSaturated") {
        return $0.map(CtorSaturated(v._1)(v._2)(v._3)(v._4))(traverse7(traversableTuple.traverse(dictApplicative)(f))(v._5));
      }
      if (v.tag === "LetRec") {
        return Apply0.apply($0.map(LetRec(v._1))(traverse9(traversableTuple.traverse(dictApplicative)(f))(v._2)))(f(v._3));
      }
      if (v.tag === "Let") {
        return Apply0.apply($0.map(Let(v._1)(v._2))(f(v._3)))(f(v._4));
      }
      if (v.tag === "EffectBind") {
        return Apply0.apply($0.map(EffectBind(v._1)(v._2))(f(v._3)))(f(v._4));
      }
      if (v.tag === "EffectPure") {
        return $0.map(EffectPure)(f(v._1));
      }
      if (v.tag === "EffectDefer") {
        return $0.map(EffectDefer)(f(v._1));
      }
      if (v.tag === "Branch") {
        return Apply0.apply($0.map(Branch)(traverse9(traverse11(f))(v._1)))(f(v._2));
      }
      if (v.tag === "PrimOp") {
        return $0.map(PrimOp)(traverse12(f)(v._1));
      }
      if (v.tag === "PrimEffect") {
        return $0.map(PrimEffect)(traverse13(f)(v._1));
      }
      if (v.tag === "PrimUndefined") {
        return dictApplicative.pure(PrimUndefined);
      }
      if (v.tag === "Fail") {
        return dictApplicative.pure($BackendSyntax("Fail", v._1));
      }
      fail();
    };
  },
  Functor0: () => functorBackendSyntax,
  Foldable1: () => foldableBackendSyntax
};
var eqTuple2 = { eq: (x) => (y) => (x._1.tag === "Nothing" ? y._1.tag === "Nothing" : x._1.tag === "Just" && y._1.tag === "Just" && x._1._1 === y._1._1) && x._2 === y._2 };
var eq9 = /* @__PURE__ */ (() => eqArrayImpl(eqTuple2.eq))();
var eq10 = /* @__PURE__ */ (() => eqArrayImpl(eqTuple2.eq))();
var eqBackendOperator2 = {
  eq: (x) => (y) => {
    if (x.tag === "OpArrayIndex") {
      return y.tag === "OpArrayIndex";
    }
    if (x.tag === "OpBooleanAnd") {
      return y.tag === "OpBooleanAnd";
    }
    if (x.tag === "OpBooleanOr") {
      return y.tag === "OpBooleanOr";
    }
    if (x.tag === "OpBooleanOrd") {
      return y.tag === "OpBooleanOrd" && (() => {
        if (x._1 === "OpEq") {
          return y._1 === "OpEq";
        }
        if (x._1 === "OpNotEq") {
          return y._1 === "OpNotEq";
        }
        if (x._1 === "OpGt") {
          return y._1 === "OpGt";
        }
        if (x._1 === "OpGte") {
          return y._1 === "OpGte";
        }
        if (x._1 === "OpLt") {
          return y._1 === "OpLt";
        }
        return x._1 === "OpLte" && y._1 === "OpLte";
      })();
    }
    if (x.tag === "OpCharOrd") {
      return y.tag === "OpCharOrd" && (() => {
        if (x._1 === "OpEq") {
          return y._1 === "OpEq";
        }
        if (x._1 === "OpNotEq") {
          return y._1 === "OpNotEq";
        }
        if (x._1 === "OpGt") {
          return y._1 === "OpGt";
        }
        if (x._1 === "OpGte") {
          return y._1 === "OpGte";
        }
        if (x._1 === "OpLt") {
          return y._1 === "OpLt";
        }
        return x._1 === "OpLte" && y._1 === "OpLte";
      })();
    }
    if (x.tag === "OpIntBitAnd") {
      return y.tag === "OpIntBitAnd";
    }
    if (x.tag === "OpIntBitOr") {
      return y.tag === "OpIntBitOr";
    }
    if (x.tag === "OpIntBitShiftLeft") {
      return y.tag === "OpIntBitShiftLeft";
    }
    if (x.tag === "OpIntBitShiftRight") {
      return y.tag === "OpIntBitShiftRight";
    }
    if (x.tag === "OpIntBitXor") {
      return y.tag === "OpIntBitXor";
    }
    if (x.tag === "OpIntBitZeroFillShiftRight") {
      return y.tag === "OpIntBitZeroFillShiftRight";
    }
    if (x.tag === "OpIntNum") {
      return y.tag === "OpIntNum" && (() => {
        if (x._1 === "OpAdd") {
          return y._1 === "OpAdd";
        }
        if (x._1 === "OpDivide") {
          return y._1 === "OpDivide";
        }
        if (x._1 === "OpMultiply") {
          return y._1 === "OpMultiply";
        }
        return x._1 === "OpSubtract" && y._1 === "OpSubtract";
      })();
    }
    if (x.tag === "OpIntOrd") {
      return y.tag === "OpIntOrd" && (() => {
        if (x._1 === "OpEq") {
          return y._1 === "OpEq";
        }
        if (x._1 === "OpNotEq") {
          return y._1 === "OpNotEq";
        }
        if (x._1 === "OpGt") {
          return y._1 === "OpGt";
        }
        if (x._1 === "OpGte") {
          return y._1 === "OpGte";
        }
        if (x._1 === "OpLt") {
          return y._1 === "OpLt";
        }
        return x._1 === "OpLte" && y._1 === "OpLte";
      })();
    }
    if (x.tag === "OpNumberNum") {
      return y.tag === "OpNumberNum" && (() => {
        if (x._1 === "OpAdd") {
          return y._1 === "OpAdd";
        }
        if (x._1 === "OpDivide") {
          return y._1 === "OpDivide";
        }
        if (x._1 === "OpMultiply") {
          return y._1 === "OpMultiply";
        }
        return x._1 === "OpSubtract" && y._1 === "OpSubtract";
      })();
    }
    if (x.tag === "OpNumberOrd") {
      return y.tag === "OpNumberOrd" && (() => {
        if (x._1 === "OpEq") {
          return y._1 === "OpEq";
        }
        if (x._1 === "OpNotEq") {
          return y._1 === "OpNotEq";
        }
        if (x._1 === "OpGt") {
          return y._1 === "OpGt";
        }
        if (x._1 === "OpGte") {
          return y._1 === "OpGte";
        }
        if (x._1 === "OpLt") {
          return y._1 === "OpLt";
        }
        return x._1 === "OpLte" && y._1 === "OpLte";
      })();
    }
    if (x.tag === "OpStringAppend") {
      return y.tag === "OpStringAppend";
    }
    return x.tag === "OpStringOrd" && y.tag === "OpStringOrd" && (() => {
      if (x._1 === "OpEq") {
        return y._1 === "OpEq";
      }
      if (x._1 === "OpNotEq") {
        return y._1 === "OpNotEq";
      }
      if (x._1 === "OpGt") {
        return y._1 === "OpGt";
      }
      if (x._1 === "OpGte") {
        return y._1 === "OpGte";
      }
      if (x._1 === "OpLt") {
        return y._1 === "OpLt";
      }
      return x._1 === "OpLte" && y._1 === "OpLte";
    })();
  }
};
var eqBackendOperator = (dictEq) => ({
  eq: (x) => (y) => {
    if (x.tag === "Op1") {
      return y.tag === "Op1" && (() => {
        if (x._1.tag === "OpBooleanNot") {
          return y._1.tag === "OpBooleanNot";
        }
        if (x._1.tag === "OpIntBitNot") {
          return y._1.tag === "OpIntBitNot";
        }
        if (x._1.tag === "OpIntNegate") {
          return y._1.tag === "OpIntNegate";
        }
        if (x._1.tag === "OpNumberNegate") {
          return y._1.tag === "OpNumberNegate";
        }
        if (x._1.tag === "OpArrayLength") {
          return y._1.tag === "OpArrayLength";
        }
        return x._1.tag === "OpIsTag" && y._1.tag === "OpIsTag" && (x._1._1._1.tag === "Nothing" ? y._1._1._1.tag === "Nothing" : x._1._1._1.tag === "Just" && y._1._1._1.tag === "Just" && x._1._1._1._1 === y._1._1._1._1) && x._1._1._2 === y._1._1._2;
      })() && dictEq.eq(x._2)(y._2);
    }
    return x.tag === "Op2" && y.tag === "Op2" && eqBackendOperator2.eq(x._1)(y._1) && dictEq.eq(x._2)(y._2) && dictEq.eq(x._3)(y._3);
  }
});
var eqBackendAccessor = {
  eq: (x) => (y) => {
    if (x.tag === "GetProp") {
      return y.tag === "GetProp" && x._1 === y._1;
    }
    if (x.tag === "GetIndex") {
      return y.tag === "GetIndex" && x._1 === y._1;
    }
    return x.tag === "GetCtorField" && y.tag === "GetCtorField" && (x._1._1.tag === "Nothing" ? y._1._1.tag === "Nothing" : x._1._1.tag === "Just" && y._1._1.tag === "Just" && x._1._1._1 === y._1._1._1) && x._1._2 === y._1._2 && (x._2 === "ProductType" ? y._2 === "ProductType" : x._2 === "SumType" && y._2 === "SumType") && x._3 === y._3 && x._4 === y._4 && x._5 === y._5 && x._6 === y._6;
  }
};
var eqBackendSyntax = (dictEq) => {
  const $0 = eqArrayImpl(dictEq.eq);
  const eq20 = eqArrayImpl((x) => (y) => x._1 === y._1 && dictEq.eq(x._2)(y._2));
  const eq21 = eqArrayImpl((x) => (y) => x._1 === y._1 && dictEq.eq(x._2)(y._2));
  const eq22 = eqArrayImpl((x) => (y) => x._1 === y._1 && dictEq.eq(x._2)(y._2));
  const eq23 = eqArrayImpl((x) => (y) => dictEq.eq(x._1)(y._1) && dictEq.eq(x._2)(y._2));
  return {
    eq: (x) => (y) => {
      if (x.tag === "Var") {
        return y.tag === "Var" && (x._1._1.tag === "Nothing" ? y._1._1.tag === "Nothing" : x._1._1.tag === "Just" && y._1._1.tag === "Just" && x._1._1._1 === y._1._1._1) && x._1._2 === y._1._2;
      }
      if (x.tag === "Local") {
        return y.tag === "Local" && (x._1.tag === "Nothing" ? y._1.tag === "Nothing" : x._1.tag === "Just" && y._1.tag === "Just" && x._1._1 === y._1._1) && x._2 === y._2;
      }
      if (x.tag === "Lit") {
        return y.tag === "Lit" && eqLiteral(dictEq).eq(x._1)(y._1);
      }
      if (x.tag === "App") {
        return y.tag === "App" && dictEq.eq(x._1)(y._1) && eqArrayImpl(dictEq.eq)(x._2)(y._2);
      }
      if (x.tag === "Abs") {
        return y.tag === "Abs" && eq9(x._1)(y._1) && dictEq.eq(x._2)(y._2);
      }
      if (x.tag === "UncurriedApp") {
        return y.tag === "UncurriedApp" && dictEq.eq(x._1)(y._1) && $0(x._2)(y._2);
      }
      if (x.tag === "UncurriedAbs") {
        return y.tag === "UncurriedAbs" && eq10(x._1)(y._1) && dictEq.eq(x._2)(y._2);
      }
      if (x.tag === "UncurriedEffectApp") {
        return y.tag === "UncurriedEffectApp" && dictEq.eq(x._1)(y._1) && $0(x._2)(y._2);
      }
      if (x.tag === "UncurriedEffectAbs") {
        return y.tag === "UncurriedEffectAbs" && eq10(x._1)(y._1) && dictEq.eq(x._2)(y._2);
      }
      if (x.tag === "Accessor") {
        return y.tag === "Accessor" && dictEq.eq(x._1)(y._1) && eqBackendAccessor.eq(x._2)(y._2);
      }
      if (x.tag === "Update") {
        return y.tag === "Update" && dictEq.eq(x._1)(y._1) && eq20(x._2)(y._2);
      }
      if (x.tag === "CtorSaturated") {
        return y.tag === "CtorSaturated" && (x._1._1.tag === "Nothing" ? y._1._1.tag === "Nothing" : x._1._1.tag === "Just" && y._1._1.tag === "Just" && x._1._1._1 === y._1._1._1) && x._1._2 === y._1._2 && (x._2 === "ProductType" ? y._2 === "ProductType" : x._2 === "SumType" && y._2 === "SumType") && x._3 === y._3 && x._4 === y._4 && eq21(x._5)(y._5);
      }
      if (x.tag === "CtorDef") {
        return y.tag === "CtorDef" && (x._1 === "ProductType" ? y._1 === "ProductType" : x._1 === "SumType" && y._1 === "SumType") && x._2 === y._2 && x._3 === y._3 && eq7(x._4)(y._4);
      }
      if (x.tag === "LetRec") {
        return y.tag === "LetRec" && x._1 === y._1 && eq22(x._2)(y._2) && dictEq.eq(x._3)(y._3);
      }
      if (x.tag === "Let") {
        return y.tag === "Let" && (x._1.tag === "Nothing" ? y._1.tag === "Nothing" : x._1.tag === "Just" && y._1.tag === "Just" && x._1._1 === y._1._1) && x._2 === y._2 && dictEq.eq(x._3)(y._3) && dictEq.eq(x._4)(y._4);
      }
      if (x.tag === "EffectBind") {
        return y.tag === "EffectBind" && (x._1.tag === "Nothing" ? y._1.tag === "Nothing" : x._1.tag === "Just" && y._1.tag === "Just" && x._1._1 === y._1._1) && x._2 === y._2 && dictEq.eq(x._3)(y._3) && dictEq.eq(x._4)(y._4);
      }
      if (x.tag === "EffectPure") {
        return y.tag === "EffectPure" && dictEq.eq(x._1)(y._1);
      }
      if (x.tag === "EffectDefer") {
        return y.tag === "EffectDefer" && dictEq.eq(x._1)(y._1);
      }
      if (x.tag === "Branch") {
        return y.tag === "Branch" && eq23(x._1)(y._1) && dictEq.eq(x._2)(y._2);
      }
      if (x.tag === "PrimOp") {
        return y.tag === "PrimOp" && eqBackendOperator(dictEq).eq(x._1)(y._1);
      }
      if (x.tag === "PrimEffect") {
        return y.tag === "PrimEffect" && (() => {
          if (x._1.tag === "EffectRefNew") {
            return y._1.tag === "EffectRefNew" && dictEq.eq(x._1._1)(y._1._1);
          }
          if (x._1.tag === "EffectRefRead") {
            return y._1.tag === "EffectRefRead" && dictEq.eq(x._1._1)(y._1._1);
          }
          return x._1.tag === "EffectRefWrite" && y._1.tag === "EffectRefWrite" && dictEq.eq(x._1._1)(y._1._1) && dictEq.eq(x._1._2)(y._1._2);
        })();
      }
      if (x.tag === "PrimUndefined") {
        return y.tag === "PrimUndefined";
      }
      return x.tag === "Fail" && y.tag === "Fail" && x._1 === y._1;
    }
  };
};

// output-es/PureScript.Backend.Optimizer.Analysis/index.js
var $Capture = (tag) => tag;
var $Complexity = (tag) => tag;
var $ResultTerm = (tag) => tag;
var ordQualified2 = /* @__PURE__ */ ordQualified(ordString);
var pop2 = /* @__PURE__ */ pop(ordInt);
var KnownNeutral = /* @__PURE__ */ $ResultTerm("KnownNeutral");
var Unknown = /* @__PURE__ */ $ResultTerm("Unknown");
var Trivial = /* @__PURE__ */ $Complexity("Trivial");
var Deref = /* @__PURE__ */ $Complexity("Deref");
var KnownSize = /* @__PURE__ */ $Complexity("KnownSize");
var NonTrivial = /* @__PURE__ */ $Complexity("NonTrivial");
var CaptureNone = /* @__PURE__ */ $Capture("CaptureNone");
var CaptureBranch = /* @__PURE__ */ $Capture("CaptureBranch");
var CaptureClosure = /* @__PURE__ */ $Capture("CaptureClosure");
var semigroupResultTerm = {
  append: (v) => (v1) => {
    if (v === "Unknown") {
      return Unknown;
    }
    if (v1 === "Unknown") {
      return Unknown;
    }
    return KnownNeutral;
  }
};
var monoidResultTerm = { mempty: KnownNeutral, Semigroup0: () => semigroupResultTerm };
var foldMap1 = /* @__PURE__ */ (() => foldableArray.foldMap(monoidResultTerm))();
var semigroupUsage = {
  append: (v) => (v1) => ({
    total: v.total + v1.total | 0,
    captured: (() => {
      if (v.captured === "CaptureNone") {
        if (v1.captured === "CaptureNone") {
          return v.captured;
        }
        return v1.captured;
      }
      if (v1.captured === "CaptureNone") {
        return v.captured;
      }
      if (v.captured === "CaptureBranch") {
        if (v1.captured === "CaptureBranch") {
          return v.captured;
        }
        return v1.captured;
      }
      if (v1.captured === "CaptureBranch") {
        return v.captured;
      }
      if (v.captured === "CaptureClosure" && v1.captured === "CaptureClosure") {
        return v.captured;
      }
      fail();
    })(),
    arities: unsafeUnionWith(ordInt.compare, $$const, v.arities, v1.arities),
    call: v.call + v1.call | 0,
    access: v.access + v1.access | 0,
    case: v.case + v1.case | 0,
    update: v.update + v1.update | 0
  })
};
var monoidUsage = { mempty: { total: 0, captured: CaptureNone, arities: Leaf, call: 0, access: 0, case: 0, update: 0 }, Semigroup0: () => semigroupUsage };
var semigroupBackendAnalysis = {
  append: (v) => (v1) => ({
    usages: unsafeUnionWith(ordInt.compare, semigroupUsage.append, v.usages, v1.usages),
    size: v.size + v1.size | 0,
    complexity: (() => {
      if (v.complexity === "Trivial") {
        if (v1.complexity === "Trivial") {
          return v.complexity;
        }
        return v1.complexity;
      }
      if (v1.complexity === "Trivial") {
        return v.complexity;
      }
      if (v.complexity === "Deref") {
        if (v1.complexity === "Deref") {
          return v.complexity;
        }
        return v1.complexity;
      }
      if (v1.complexity === "Deref") {
        return v.complexity;
      }
      if (v.complexity === "KnownSize") {
        if (v1.complexity === "KnownSize") {
          return v.complexity;
        }
        return v1.complexity;
      }
      if (v1.complexity === "KnownSize") {
        return v.complexity;
      }
      if (v.complexity === "NonTrivial" && v1.complexity === "NonTrivial") {
        return v.complexity;
      }
      fail();
    })(),
    args: [],
    rewrite: v.rewrite || v1.rewrite,
    deps: unsafeUnionWith(ordQualified2.compare, $$const, v.deps, v1.deps),
    result: (() => {
      if (v.result === "Unknown") {
        return Unknown;
      }
      if (v1.result === "Unknown") {
        return Unknown;
      }
      return KnownNeutral;
    })(),
    externs: v.externs || v1.externs
  })
};
var monoidBackendAnalysis = {
  mempty: { usages: Leaf, size: 0, complexity: Trivial, args: [], rewrite: false, deps: Leaf, result: KnownNeutral, externs: false },
  Semigroup0: () => semigroupBackendAnalysis
};
var foldMap2 = /* @__PURE__ */ (() => foldableBackendSyntax.foldMap(monoidBackendAnalysis))();
var foldMap3 = /* @__PURE__ */ (() => foldableArray.foldMap(monoidBackendAnalysis))();
var foldMap4 = /* @__PURE__ */ (() => foldableArray.foldMap(monoidBackendAnalysis))();
var foldMap6 = (f) => (v) => semigroupBackendAnalysis.append(f(v._1))(f(v._2));
var used = (level) => ({
  ...monoidBackendAnalysis.mempty,
  usages: $$$Map(
    "Node",
    1,
    1,
    level,
    { total: 1, captured: CaptureNone, arities: Leaf, call: 0, access: 0, case: 0, update: 0 },
    Leaf,
    Leaf
  )
});
var updated = (level) => (v) => ({ ...v, usages: update(ordInt)((x) => $Maybe("Just", { ...x, update: x.update + 1 | 0 }))(level)(v.usages) });
var cased = (level) => (v) => ({ ...v, usages: update(ordInt)((x) => $Maybe("Just", { ...x, case: x.case + 1 | 0 }))(level)(v.usages) });
var callArity = (lvl) => (arity) => (v) => ({
  ...v,
  usages: update(ordInt)((x) => $Maybe(
    "Just",
    { ...x, arities: insert(ordInt)(arity)()(x.arities), call: x.call + 1 | 0 }
  ))(lvl)(v.usages)
});
var boundArg = (level) => (v) => {
  const v1 = pop2(level)(v.usages);
  if (v1.tag === "Nothing") {
    return { ...v, args: [monoidUsage.mempty, ...v.args] };
  }
  if (v1.tag === "Just") {
    return { ...v, usages: v1._1._2, args: [v1._1._1, ...v.args] };
  }
  fail();
};
var analyzeDefault = (dictHasAnalysis) => {
  const $0 = foldMap2(dictHasAnalysis.analysisOf);
  return (x) => {
    const $1 = $0(x);
    return { ...$1, size: $1.size + 1 | 0 };
  };
};
var accessed = (level) => (v) => ({ ...v, usages: update(ordInt)((x) => $Maybe("Just", { ...x, access: x.access + 1 | 0 }))(level)(v.usages) });
var analyze = (dictHasAnalysis) => {
  const analysisOf1 = dictHasAnalysis.analysisOf;
  const analyzeDefault1 = analyzeDefault(dictHasAnalysis);
  return (dictHasSyntax) => (externAnalysis) => (expr) => {
    if (expr.tag === "Var") {
      const analysis = { ...monoidBackendAnalysis.mempty, deps: insert(ordQualified2)(expr._1)()(monoidBackendAnalysis.mempty.deps), externs: true, size: 1 };
      const v = externAnalysis(expr._1)(Nothing);
      if (v.tag === "Just") {
        return { ...analysis, args: v._1.args };
      }
      if (v.tag === "Nothing") {
        return analysis;
      }
      fail();
    }
    if (expr.tag === "Local") {
      const $0 = used(expr._2);
      return { ...$0, size: $0.size + 1 | 0 };
    }
    if (expr.tag === "Let") {
      const $0 = semigroupBackendAnalysis.append(analysisOf1(expr._3))((() => {
        const $02 = analysisOf1(expr._4);
        return { ...$02, usages: $$delete(ordInt)(expr._2)($02.usages) };
      })());
      return {
        ...$0,
        complexity: (() => {
          if ($0.complexity === "Trivial") {
            return NonTrivial;
          }
          if ($0.complexity === "Deref") {
            return NonTrivial;
          }
          if ($0.complexity === "KnownSize") {
            return NonTrivial;
          }
          if ($0.complexity === "NonTrivial") {
            return $0.complexity;
          }
          fail();
        })(),
        result: dictHasAnalysis.analysisOf(expr._4).result,
        size: $0.size + 1 | 0
      };
    }
    if (expr.tag === "LetRec") {
      const $0 = semigroupBackendAnalysis.append(foldMap3((x) => analysisOf1(x._2))(expr._2))(analysisOf1(expr._3));
      return {
        ...$0,
        complexity: (() => {
          if ($0.complexity === "Trivial") {
            return NonTrivial;
          }
          if ($0.complexity === "Deref") {
            return NonTrivial;
          }
          if ($0.complexity === "KnownSize") {
            return NonTrivial;
          }
          if ($0.complexity === "NonTrivial") {
            return $0.complexity;
          }
          fail();
        })(),
        result: dictHasAnalysis.analysisOf(expr._3).result,
        size: $0.size + 1 | 0,
        usages: $$delete(ordInt)(expr._1)($0.usages)
      };
    }
    if (expr.tag === "EffectBind") {
      const $0 = semigroupBackendAnalysis.append(analysisOf1(expr._3))((() => {
        const $02 = analysisOf1(expr._4);
        return { ...$02, usages: $$delete(ordInt)(expr._2)($02.usages) };
      })());
      const go = (v) => {
        if (v.tag === "Leaf") {
          return Leaf;
        }
        if (v.tag === "Node") {
          return $$$Map("Node", v._1, v._2, v._3, { ...v._4, captured: CaptureClosure }, go(v._5), go(v._6));
        }
        fail();
      };
      return {
        ...$0,
        complexity: (() => {
          if ($0.complexity === "Trivial") {
            return NonTrivial;
          }
          if ($0.complexity === "Deref") {
            return NonTrivial;
          }
          if ($0.complexity === "KnownSize") {
            return NonTrivial;
          }
          if ($0.complexity === "NonTrivial") {
            return $0.complexity;
          }
          fail();
        })(),
        result: Unknown,
        size: $0.size + 1 | 0,
        usages: go($0.usages)
      };
    }
    if (expr.tag === "EffectPure") {
      const $0 = analysisOf1(expr._1);
      return {
        ...$0,
        result: Unknown,
        size: $0.size + 1 | 0,
        usages: (() => {
          const go = (v) => {
            if (v.tag === "Leaf") {
              return Leaf;
            }
            if (v.tag === "Node") {
              return $$$Map("Node", v._1, v._2, v._3, { ...v._4, captured: CaptureClosure }, go(v._5), go(v._6));
            }
            fail();
          };
          return go($0.usages);
        })()
      };
    }
    if (expr.tag === "EffectDefer") {
      const $0 = analysisOf1(expr._1);
      return {
        ...$0,
        result: Unknown,
        size: $0.size + 1 | 0,
        usages: (() => {
          const go = (v) => {
            if (v.tag === "Leaf") {
              return Leaf;
            }
            if (v.tag === "Node") {
              return $$$Map("Node", v._1, v._2, v._3, { ...v._4, captured: CaptureClosure }, go(v._5), go(v._6));
            }
            fail();
          };
          return go($0.usages);
        })()
      };
    }
    if (expr.tag === "Abs") {
      const $0 = foldrArray((x) => boundArg(x._2))(analyzeDefault1(expr))(expr._1);
      const go = (v) => {
        if (v.tag === "Leaf") {
          return Leaf;
        }
        if (v.tag === "Node") {
          return $$$Map("Node", v._1, v._2, v._3, { ...v._4, captured: CaptureClosure }, go(v._5), go(v._6));
        }
        fail();
      };
      return {
        ...$0,
        complexity: (() => {
          if ($0.complexity === "Trivial") {
            return KnownSize;
          }
          if ($0.complexity === "Deref") {
            return KnownSize;
          }
          return $0.complexity;
        })(),
        result: KnownNeutral,
        usages: go($0.usages)
      };
    }
    if (expr.tag === "UncurriedAbs") {
      const $0 = foldrArray((x) => boundArg(x._2))(analyzeDefault1(expr))(expr._1);
      const go = (v) => {
        if (v.tag === "Leaf") {
          return Leaf;
        }
        if (v.tag === "Node") {
          return $$$Map("Node", v._1, v._2, v._3, { ...v._4, captured: CaptureClosure }, go(v._5), go(v._6));
        }
        fail();
      };
      return {
        ...$0,
        complexity: (() => {
          if ($0.complexity === "Trivial") {
            return KnownSize;
          }
          if ($0.complexity === "Deref") {
            return KnownSize;
          }
          return $0.complexity;
        })(),
        result: KnownNeutral,
        usages: go($0.usages)
      };
    }
    if (expr.tag === "UncurriedApp") {
      const $0 = analyzeDefault1(expr);
      const analysis = {
        ...$0,
        complexity: (() => {
          if ($0.complexity === "Trivial") {
            return NonTrivial;
          }
          if ($0.complexity === "Deref") {
            return NonTrivial;
          }
          if ($0.complexity === "KnownSize") {
            return NonTrivial;
          }
          if ($0.complexity === "NonTrivial") {
            return $0.complexity;
          }
          fail();
        })(),
        result: Unknown
      };
      const v = dictHasSyntax.syntaxOf(expr._1);
      if (v.tag === "Just" && v._1.tag === "Local") {
        return callArity(v._1._2)(expr._2.length)(analysis);
      }
      return analysis;
    }
    if (expr.tag === "UncurriedEffectAbs") {
      const $0 = foldrArray((x) => boundArg(x._2))(analyzeDefault1(expr))(expr._1);
      const go = (v) => {
        if (v.tag === "Leaf") {
          return Leaf;
        }
        if (v.tag === "Node") {
          return $$$Map("Node", v._1, v._2, v._3, { ...v._4, captured: CaptureClosure }, go(v._5), go(v._6));
        }
        fail();
      };
      return {
        ...$0,
        complexity: (() => {
          if ($0.complexity === "Trivial") {
            return KnownSize;
          }
          if ($0.complexity === "Deref") {
            return KnownSize;
          }
          return $0.complexity;
        })(),
        result: KnownNeutral,
        usages: go($0.usages)
      };
    }
    if (expr.tag === "UncurriedEffectApp") {
      const $0 = analyzeDefault1(expr);
      const go = (v2) => {
        if (v2.tag === "Leaf") {
          return Leaf;
        }
        if (v2.tag === "Node") {
          return $$$Map("Node", v2._1, v2._2, v2._3, { ...v2._4, captured: CaptureClosure }, go(v2._5), go(v2._6));
        }
        fail();
      };
      const analysis = {
        ...$0,
        complexity: (() => {
          if ($0.complexity === "Trivial") {
            return NonTrivial;
          }
          if ($0.complexity === "Deref") {
            return NonTrivial;
          }
          if ($0.complexity === "KnownSize") {
            return NonTrivial;
          }
          if ($0.complexity === "NonTrivial") {
            return $0.complexity;
          }
          fail();
        })(),
        result: Unknown,
        usages: go($0.usages)
      };
      const v = dictHasSyntax.syntaxOf(expr._1);
      if (v.tag === "Just" && v._1.tag === "Local") {
        return callArity(v._1._2)(expr._2.length)(analysis);
      }
      return analysis;
    }
    if (expr.tag === "App") {
      const $0 = analysisOf1(expr._1).args;
      const $1 = expr._2.length;
      const remainingArgs = $1 < 1 ? $0 : sliceImpl($1, $0.length, $0);
      const analysis = (() => {
        if (remainingArgs.length === 0) {
          const $2 = analyzeDefault1(expr);
          return {
            ...$2,
            complexity: (() => {
              if ($2.complexity === "Trivial") {
                return NonTrivial;
              }
              if ($2.complexity === "Deref") {
                return NonTrivial;
              }
              if ($2.complexity === "KnownSize") {
                return NonTrivial;
              }
              if ($2.complexity === "NonTrivial") {
                return $2.complexity;
              }
              fail();
            })()
          };
        }
        return analyzeDefault1(expr);
      })();
      const v1 = dictHasSyntax.syntaxOf(expr._1);
      return {
        ...v1.tag === "Just" && v1._1.tag === "Local" ? { ...callArity(v1._1._2)(expr._2.length)({ ...analysis, size: analysis.size + 1 | 0 }), result: Unknown } : { ...analysis, result: Unknown, size: analysis.size + 1 | 0 },
        args: remainingArgs
      };
    }
    if (expr.tag === "Update") {
      const $0 = analyzeDefault1(expr);
      const analysis = {
        ...$0,
        complexity: (() => {
          if ($0.complexity === "Trivial") {
            return NonTrivial;
          }
          if ($0.complexity === "Deref") {
            return NonTrivial;
          }
          if ($0.complexity === "KnownSize") {
            return NonTrivial;
          }
          if ($0.complexity === "NonTrivial") {
            return $0.complexity;
          }
          fail();
        })(),
        result: Unknown
      };
      const v2 = dictHasSyntax.syntaxOf(expr._1);
      if (v2.tag === "Just" && v2._1.tag === "Local") {
        return updated(v2._1._2)(analysis);
      }
      return analysis;
    }
    if (expr.tag === "CtorSaturated") {
      const $0 = foldMap4((v) => analysisOf1(v._2))(expr._5);
      return { ...$0, deps: insert(ordQualified2)(expr._1)()($0.deps), result: KnownNeutral, size: $0.size + 1 | 0 };
    }
    if (expr.tag === "CtorDef") {
      const $0 = analyzeDefault1(expr);
      return {
        ...$0,
        complexity: (() => {
          if ($0.complexity === "Trivial") {
            return NonTrivial;
          }
          if ($0.complexity === "Deref") {
            return NonTrivial;
          }
          if ($0.complexity === "KnownSize") {
            return NonTrivial;
          }
          if ($0.complexity === "NonTrivial") {
            return $0.complexity;
          }
          fail();
        })()
      };
    }
    if (expr.tag === "Branch") {
      const v2 = (() => {
        if (0 < expr._1.length) {
          return expr._1[0];
        }
        fail();
      })();
      const $0 = semigroupBackendAnalysis.append(analysisOf1(v2._1))(semigroupBackendAnalysis.append((() => {
        const $02 = analysisOf1(v2._2);
        return {
          ...$02,
          usages: (() => {
            const go = (v) => {
              if (v.tag === "Leaf") {
                return Leaf;
              }
              if (v.tag === "Node") {
                return $$$Map("Node", v._1, v._2, v._3, { ...v._4, captured: CaptureBranch }, go(v._5), go(v._6));
              }
              fail();
            };
            return go($02.usages);
          })()
        };
      })())(semigroupBackendAnalysis.append((() => {
        const $02 = foldMap4(foldMap6(analysisOf1))((() => {
          const $03 = unconsImpl((v) => Nothing, (v) => (xs) => $Maybe("Just", xs), expr._1);
          if ($03.tag === "Just") {
            return $03._1;
          }
          fail();
        })());
        return {
          ...$02,
          usages: (() => {
            const go = (v) => {
              if (v.tag === "Leaf") {
                return Leaf;
              }
              if (v.tag === "Node") {
                return $$$Map("Node", v._1, v._2, v._3, { ...v._4, captured: CaptureBranch }, go(v._5), go(v._6));
              }
              fail();
            };
            return go($02.usages);
          })()
        };
      })())((() => {
        const $02 = analysisOf1(expr._2);
        return {
          ...$02,
          usages: (() => {
            const go = (v) => {
              if (v.tag === "Leaf") {
                return Leaf;
              }
              if (v.tag === "Node") {
                return $$$Map("Node", v._1, v._2, v._3, { ...v._4, captured: CaptureBranch }, go(v._5), go(v._6));
              }
              fail();
            };
            return go($02.usages);
          })()
        };
      })())));
      return {
        ...$0,
        complexity: (() => {
          if ($0.complexity === "Trivial") {
            return NonTrivial;
          }
          if ($0.complexity === "Deref") {
            return NonTrivial;
          }
          if ($0.complexity === "KnownSize") {
            return NonTrivial;
          }
          if ($0.complexity === "NonTrivial") {
            return $0.complexity;
          }
          fail();
        })(),
        result: foldMap1((x) => dictHasAnalysis.analysisOf(x._2).result)(expr._1)
      };
    }
    if (expr.tag === "Fail") {
      const $0 = analyzeDefault1(expr);
      return {
        ...$0,
        complexity: (() => {
          if ($0.complexity === "Trivial") {
            return NonTrivial;
          }
          if ($0.complexity === "Deref") {
            return NonTrivial;
          }
          if ($0.complexity === "KnownSize") {
            return NonTrivial;
          }
          if ($0.complexity === "NonTrivial") {
            return $0.complexity;
          }
          fail();
        })()
      };
    }
    if (expr.tag === "PrimOp") {
      const $0 = expr._1;
      const $1 = analyzeDefault1(expr);
      const analysis = {
        ...$1,
        complexity: (() => {
          if ($1.complexity === "Trivial") {
            return NonTrivial;
          }
          if ($1.complexity === "Deref") {
            return NonTrivial;
          }
          if ($1.complexity === "KnownSize") {
            return NonTrivial;
          }
          if ($1.complexity === "NonTrivial") {
            return $1.complexity;
          }
          fail();
        })(),
        result: Unknown
      };
      const v2 = (v3) => {
        if ($0.tag === "Op1" && $0._1.tag === "OpIsTag") {
          return { ...analysis, deps: insert(ordQualified2)($0._1._1)()(analysis.deps) };
        }
        return analysis;
      };
      if ($0.tag === "Op1" && $0._1.tag === "OpIsTag") {
        const $2 = dictHasSyntax.syntaxOf($0._2);
        if ($2.tag === "Just" && $2._1.tag === "Local") {
          return cased($2._1._2)({ ...analysis, deps: insert(ordQualified2)($0._1._1)()(analysis.deps) });
        }
      }
      return v2(true);
    }
    if (expr.tag === "PrimEffect") {
      const $0 = analyzeDefault1(expr);
      const go = (v) => {
        if (v.tag === "Leaf") {
          return Leaf;
        }
        if (v.tag === "Node") {
          return $$$Map("Node", v._1, v._2, v._3, { ...v._4, captured: CaptureClosure }, go(v._5), go(v._6));
        }
        fail();
      };
      return {
        ...$0,
        complexity: (() => {
          if ($0.complexity === "Trivial") {
            return NonTrivial;
          }
          if ($0.complexity === "Deref") {
            return NonTrivial;
          }
          if ($0.complexity === "KnownSize") {
            return NonTrivial;
          }
          if ($0.complexity === "NonTrivial") {
            return $0.complexity;
          }
          fail();
        })(),
        result: Unknown,
        usages: go($0.usages)
      };
    }
    if (expr.tag === "PrimUndefined") {
      return analyzeDefault1(expr);
    }
    if (expr.tag === "Accessor") {
      const analysis = (() => {
        if (expr._2.tag === "GetCtorField") {
          const $0 = analyzeDefault1(expr);
          return { ...$0, deps: insert(ordQualified2)(expr._2._1)()($0.deps), result: Unknown };
        }
        return { ...analyzeDefault1(expr), result: Unknown };
      })();
      const v2 = dictHasSyntax.syntaxOf(expr._1);
      if (v2.tag === "Just") {
        if (v2._1.tag === "Accessor") {
          return analysis;
        }
        if (v2._1.tag === "Local") {
          return accessed(v2._1._2)({ ...analysis, complexity: analysis.complexity === "Trivial" ? Deref : analysis.complexity });
        }
        if (v2._1.tag === "Var") {
          if (expr._2.tag === "GetProp") {
            const $0 = externAnalysis(v2._1._1)($Maybe("Just", expr._2._1));
            if ($0.tag === "Just") {
              return { ...analysis, args: $0._1.args, complexity: analysis.complexity };
            }
          }
          return { ...analysis, complexity: analysis.complexity };
        }
      }
      return { ...analysis, complexity: analysis.complexity === "Trivial" ? Deref : analysis.complexity };
    }
    if (expr.tag === "Lit") {
      const analysis = { ...analyzeDefault1(expr), result: KnownNeutral };
      if (expr._1.tag === "LitArray") {
        if (expr._1._1.length > 0) {
          return {
            ...analysis,
            complexity: (() => {
              if (analysis.complexity === "Trivial") {
                return KnownSize;
              }
              if (analysis.complexity === "Deref") {
                return KnownSize;
              }
              return analysis.complexity;
            })()
          };
        }
        return analysis;
      }
      if (expr._1.tag === "LitRecord") {
        if (expr._1._1.length > 0) {
          return {
            ...analysis,
            complexity: (() => {
              if (analysis.complexity === "Trivial") {
                return KnownSize;
              }
              if (analysis.complexity === "Deref") {
                return KnownSize;
              }
              return analysis.complexity;
            })()
          };
        }
        return analysis;
      }
      if (expr._1.tag === "LitString" && length2(expr._1._1) > 128) {
        return {
          ...analysis,
          complexity: (() => {
            if (analysis.complexity === "Trivial") {
              return KnownSize;
            }
            if (analysis.complexity === "Deref") {
              return KnownSize;
            }
            return analysis.complexity;
          })()
        };
      }
      return analysis;
    }
    fail();
  };
};
var analyzeEffectBlock = (dictHasAnalysis) => {
  const analyzeDefault1 = analyzeDefault(dictHasAnalysis);
  const analyze1 = analyze(dictHasAnalysis);
  return (dictHasSyntax) => {
    const analyze22 = analyze1(dictHasSyntax);
    return (externAnalysis) => (expr) => {
      if (expr.tag === "Let") {
        const $0 = semigroupBackendAnalysis.append(dictHasAnalysis.analysisOf(expr._3))((() => {
          const $02 = dictHasAnalysis.analysisOf(expr._4);
          return { ...$02, usages: $$delete(ordInt)(expr._2)($02.usages) };
        })());
        return {
          ...$0,
          complexity: (() => {
            if ($0.complexity === "Trivial") {
              return NonTrivial;
            }
            if ($0.complexity === "Deref") {
              return NonTrivial;
            }
            if ($0.complexity === "KnownSize") {
              return NonTrivial;
            }
            if ($0.complexity === "NonTrivial") {
              return $0.complexity;
            }
            fail();
          })(),
          result: dictHasAnalysis.analysisOf(expr._4).result,
          size: $0.size + 1 | 0
        };
      }
      if (expr.tag === "LetRec") {
        const $0 = semigroupBackendAnalysis.append(foldMap3((x) => dictHasAnalysis.analysisOf(x._2))(expr._2))(dictHasAnalysis.analysisOf(expr._3));
        return {
          ...$0,
          complexity: (() => {
            if ($0.complexity === "Trivial") {
              return NonTrivial;
            }
            if ($0.complexity === "Deref") {
              return NonTrivial;
            }
            if ($0.complexity === "KnownSize") {
              return NonTrivial;
            }
            if ($0.complexity === "NonTrivial") {
              return $0.complexity;
            }
            fail();
          })(),
          result: dictHasAnalysis.analysisOf(expr._3).result,
          size: $0.size + 1 | 0,
          usages: $$delete(ordInt)(expr._1)($0.usages)
        };
      }
      if (expr.tag === "EffectBind") {
        const $0 = semigroupBackendAnalysis.append(dictHasAnalysis.analysisOf(expr._3))((() => {
          const $02 = dictHasAnalysis.analysisOf(expr._4);
          return { ...$02, usages: $$delete(ordInt)(expr._2)($02.usages) };
        })());
        return {
          ...$0,
          complexity: (() => {
            if ($0.complexity === "Trivial") {
              return NonTrivial;
            }
            if ($0.complexity === "Deref") {
              return NonTrivial;
            }
            if ($0.complexity === "KnownSize") {
              return NonTrivial;
            }
            if ($0.complexity === "NonTrivial") {
              return $0.complexity;
            }
            fail();
          })(),
          result: Unknown,
          size: $0.size + 1 | 0
        };
      }
      if (expr.tag === "EffectPure") {
        const $0 = dictHasAnalysis.analysisOf(expr._1);
        return { ...$0, result: Unknown, size: $0.size + 1 | 0 };
      }
      if (expr.tag === "EffectDefer") {
        const $0 = dictHasAnalysis.analysisOf(expr._1);
        return { ...$0, result: Unknown, size: $0.size + 1 | 0 };
      }
      if (expr.tag === "UncurriedEffectApp") {
        const $0 = analyzeDefault1(expr);
        const analysis = {
          ...$0,
          complexity: (() => {
            if ($0.complexity === "Trivial") {
              return NonTrivial;
            }
            if ($0.complexity === "Deref") {
              return NonTrivial;
            }
            if ($0.complexity === "KnownSize") {
              return NonTrivial;
            }
            if ($0.complexity === "NonTrivial") {
              return $0.complexity;
            }
            fail();
          })(),
          result: Unknown
        };
        const v = dictHasSyntax.syntaxOf(expr._1);
        if (v.tag === "Just" && v._1.tag === "Local") {
          return callArity(v._1._2)(expr._2.length)(analysis);
        }
        return analysis;
      }
      if (expr.tag === "PrimEffect") {
        const $0 = analyzeDefault1(expr);
        return {
          ...$0,
          complexity: (() => {
            if ($0.complexity === "Trivial") {
              return NonTrivial;
            }
            if ($0.complexity === "Deref") {
              return NonTrivial;
            }
            if ($0.complexity === "KnownSize") {
              return NonTrivial;
            }
            if ($0.complexity === "NonTrivial") {
              return $0.complexity;
            }
            fail();
          })(),
          result: Unknown
        };
      }
      return analyze22(externAnalysis)(expr);
    };
  };
};

// output-es/Data.HeytingAlgebra/foreign.js
var boolConj = function(b1) {
  return function(b2) {
    return b1 && b2;
  };
};
var boolDisj = function(b1) {
  return function(b2) {
    return b1 || b2;
  };
};
var boolNot = function(b) {
  return !b;
};

// output-es/Data.HeytingAlgebra/index.js
var heytingAlgebraBoolean = { ff: false, tt: true, implies: (a) => (b) => heytingAlgebraBoolean.disj(heytingAlgebraBoolean.not(a))(b), conj: boolConj, disj: boolDisj, not: boolNot };

// output-es/Data.Lazy/foreign.js
var defer = function(thunk) {
  var v = null;
  return function() {
    if (thunk === void 0) return v;
    v = thunk();
    thunk = void 0;
    return v;
  };
};
var force = function(l) {
  return l();
};

// output-es/PureScript.Backend.Optimizer.Semantics/index.js
var $BackendExpr = (tag, _1, _2) => ({ tag, _1, _2 });
var $BackendRewrite = (tag, _1, _2, _3, _4, _5) => ({ tag, _1, _2, _3, _4, _5 });
var $BackendSemantics = (tag, _1, _2, _3, _4, _5) => ({ tag, _1, _2, _3, _4, _5 });
var $DistOp = (tag, _1, _2) => ({ tag, _1, _2 });
var $EvalRef = (tag, _1, _2) => ({ tag, _1, _2 });
var $ExternImpl = (tag, _1, _2, _3, _4, _5) => ({ tag, _1, _2, _3, _4, _5 });
var $ExternSpine = (tag, _1) => ({ tag, _1 });
var $InlineAccessor = (tag, _1) => ({ tag, _1 });
var $InlineDirective = (tag, _1) => ({ tag, _1 });
var $LocalBinding = (tag, _1) => ({ tag, _1 });
var $MkFn = (tag, _1, _2) => ({ tag, _1, _2 });
var $SemConditional = (_1, _2) => ({ tag: "SemConditional", _1, _2 });
var $UnpackOp = (tag, _1, _2, _3, _4, _5) => ({ tag, _1, _2, _3, _4, _5 });
var compare1 = /* @__PURE__ */ (() => ordQualified(ordString).compare)();
var compare2 = (x) => (y) => {
  if (x.tag === "Nothing") {
    if (y.tag === "Nothing") {
      return EQ;
    }
    return LT;
  }
  if (y.tag === "Nothing") {
    return GT;
  }
  if (x.tag === "Just" && y.tag === "Just") {
    return ordString.compare(x._1)(y._1);
  }
  fail();
};
var eq102 = /* @__PURE__ */ eqArrayImpl((x) => (y) => (x._1.tag === "Nothing" ? y._1.tag === "Nothing" : x._1.tag === "Just" && y._1.tag === "Just" && x._1._1 === y._1._1) && x._2 === y._2);
var lookup2 = (k) => {
  const go = (go$a0$copy) => {
    let go$a0 = go$a0$copy, go$c = true, go$r;
    while (go$c) {
      const v = go$a0;
      if (v.tag === "Leaf") {
        go$c = false;
        go$r = Nothing;
        continue;
      }
      if (v.tag === "Node") {
        const v1 = ordInt.compare(k)(v._3);
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
          go$r = $Maybe("Just", v._4);
          continue;
        }
      }
      fail();
    }
    return go$r;
  };
  return go;
};
var toUnfoldable = /* @__PURE__ */ (() => {
  const $0 = unfoldableArray.unfoldr((xs) => {
    if (xs.tag === "Nil") {
      return Nothing;
    }
    if (xs.tag === "Cons") {
      return $Maybe("Just", $Tuple(xs._1, xs._2));
    }
    fail();
  });
  return (x) => $0((() => {
    const go = (m$p, z$p) => {
      if (m$p.tag === "Leaf") {
        return z$p;
      }
      if (m$p.tag === "Node") {
        return go(m$p._5, $List("Cons", m$p._3, go(m$p._6, z$p)));
      }
      fail();
    };
    return go(x, Nil);
  })());
})();
var or2 = /* @__PURE__ */ or(foldableArray)(heytingAlgebraBoolean);
var and2 = /* @__PURE__ */ and(foldableArray)(heytingAlgebraBoolean);
var foldMap = /* @__PURE__ */ (() => foldableArray.foldMap(monoidBackendAnalysis))();
var foldMap12 = /* @__PURE__ */ (() => foldableArray.foldMap(monoidBackendAnalysis))();
var power2 = /* @__PURE__ */ power(monoidBackendAnalysis);
var toUnfoldable1 = /* @__PURE__ */ (() => unfoldableArray.unfoldr((xs) => {
  if (xs.tag === "Nil") {
    return Nothing;
  }
  if (xs.tag === "Cons") {
    return $Maybe("Just", $Tuple(xs._1, xs._2));
  }
  fail();
}))();
var fromFoldable2 = /* @__PURE__ */ foldrArray(Cons)(Nil);
var eq16 = (x) => (y) => {
  if (x.tag === "Left") {
    return y.tag === "Left" && (x._1._1.tag === "Nothing" ? y._1._1.tag === "Nothing" : x._1._1.tag === "Just" && y._1._1.tag === "Just" && x._1._1._1 === y._1._1._1) && x._1._2 === y._1._2;
  }
  return x.tag === "Right" && y.tag === "Right" && eqBackendOperator2.eq(x._1)(y._1);
};
var identity9 = (x) => x;
var lookup1 = /* @__PURE__ */ lookup(foldableArray)(eqString);
var NeutralExpr = (x) => x;
var InlineDefault = /* @__PURE__ */ $InlineDirective("InlineDefault");
var InlineNever = /* @__PURE__ */ $InlineDirective("InlineNever");
var InlineAlways = /* @__PURE__ */ $InlineDirective("InlineAlways");
var InlineRef = /* @__PURE__ */ $InlineAccessor("InlineRef");
var NeutData = (value0) => (value1) => (value2) => (value3) => (value4) => $BackendSemantics("NeutData", value0, value1, value2, value3, value4);
var NeutUpdate = (value0) => (value1) => $BackendSemantics("NeutUpdate", value0, value1);
var NeutLit = (value0) => $BackendSemantics("NeutLit", value0);
var NeutUncurriedApp = (value0) => (value1) => $BackendSemantics("NeutUncurriedApp", value0, value1);
var NeutUncurriedEffectApp = (value0) => (value1) => $BackendSemantics("NeutUncurriedEffectApp", value0, value1);
var NeutPrimEffect = (value0) => $BackendSemantics("NeutPrimEffect", value0);
var NeutPrimUndefined = /* @__PURE__ */ $BackendSemantics("NeutPrimUndefined");
var hasSyntaxBackendExpr = {
  syntaxOf: (v) => {
    if (v.tag === "ExprSyntax") {
      return $Maybe("Just", v._2);
    }
    return Nothing;
  }
};
var hasAnalysisBackendExpr = {
  analysisOf: (v) => {
    if (v.tag === "ExprSyntax") {
      return v._1;
    }
    if (v.tag === "ExprRewrite") {
      return v._1;
    }
    fail();
  }
};
var eqUnpackOp = (dictEq) => {
  const eq19 = eqArrayImpl((x) => (y) => x._1 === y._1 && dictEq.eq(x._2)(y._2));
  const eq22 = eqArrayImpl((x) => (y) => x._1 === y._1 && dictEq.eq(x._2)(y._2));
  return {
    eq: (x) => (y) => {
      if (x.tag === "UnpackRecord") {
        return y.tag === "UnpackRecord" && eq19(x._1)(y._1);
      }
      if (x.tag === "UnpackUpdate") {
        return y.tag === "UnpackUpdate" && dictEq.eq(x._1)(y._1) && eq19(x._2)(y._2);
      }
      if (x.tag === "UnpackArray") {
        return y.tag === "UnpackArray" && eqArrayImpl(dictEq.eq)(x._1)(y._1);
      }
      return x.tag === "UnpackData" && y.tag === "UnpackData" && (x._1._1.tag === "Nothing" ? y._1._1.tag === "Nothing" : x._1._1.tag === "Just" && y._1._1.tag === "Just" && x._1._1._1 === y._1._1._1) && x._1._2 === y._1._2 && (x._2 === "ProductType" ? y._2 === "ProductType" : x._2 === "SumType" && y._2 === "SumType") && x._3 === y._3 && x._4 === y._4 && eq22(x._5)(y._5);
    }
  };
};
var eqInlineAccessor = {
  eq: (x) => (y) => {
    if (x.tag === "InlineProp") {
      return y.tag === "InlineProp" && x._1 === y._1;
    }
    if (x.tag === "InlineSpineProp") {
      return y.tag === "InlineSpineProp" && x._1 === y._1;
    }
    return x.tag === "InlineRef" && y.tag === "InlineRef";
  }
};
var ordInlineAccessor = {
  compare: (x) => (y) => {
    if (x.tag === "InlineProp") {
      if (y.tag === "InlineProp") {
        return ordString.compare(x._1)(y._1);
      }
      return LT;
    }
    if (y.tag === "InlineProp") {
      return GT;
    }
    if (x.tag === "InlineSpineProp") {
      if (y.tag === "InlineSpineProp") {
        return ordString.compare(x._1)(y._1);
      }
      return LT;
    }
    if (y.tag === "InlineSpineProp") {
      return GT;
    }
    if (x.tag === "InlineRef" && y.tag === "InlineRef") {
      return EQ;
    }
    fail();
  },
  Eq0: () => eqInlineAccessor
};
var lookup22 = (k) => {
  const go = (go$a0$copy) => {
    let go$a0 = go$a0$copy, go$c = true, go$r;
    while (go$c) {
      const v = go$a0;
      if (v.tag === "Leaf") {
        go$c = false;
        go$r = Nothing;
        continue;
      }
      if (v.tag === "Node") {
        const v1 = ordInlineAccessor.compare(k)(v._3);
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
          go$r = $Maybe("Just", v._4);
          continue;
        }
      }
      fail();
    }
    return go$r;
  };
  return go;
};
var eqEvalRef = {
  eq: (x) => (y) => {
    if (x.tag === "EvalExtern") {
      return y.tag === "EvalExtern" && (x._1._1.tag === "Nothing" ? y._1._1.tag === "Nothing" : x._1._1.tag === "Just" && y._1._1.tag === "Just" && x._1._1._1 === y._1._1._1) && x._1._2 === y._1._2;
    }
    return x.tag === "EvalLocal" && y.tag === "EvalLocal" && (x._1.tag === "Nothing" ? y._1.tag === "Nothing" : x._1.tag === "Just" && y._1.tag === "Just" && x._1._1 === y._1._1) && x._2 === y._2;
  }
};
var ordEvalRef = {
  compare: (x) => (y) => {
    if (x.tag === "EvalExtern") {
      if (y.tag === "EvalExtern") {
        return compare1(x._1)(y._1);
      }
      return LT;
    }
    if (y.tag === "EvalExtern") {
      return GT;
    }
    if (x.tag === "EvalLocal" && y.tag === "EvalLocal") {
      const v = compare2(x._1)(y._1);
      if (v === "LT") {
        return LT;
      }
      if (v === "GT") {
        return GT;
      }
      return ordInt.compare(x._2)(y._2);
    }
    fail();
  },
  Eq0: () => eqEvalRef
};
var alter2 = /* @__PURE__ */ alter(ordEvalRef);
var lookup3 = (k) => {
  const go = (go$a0$copy) => {
    let go$a0 = go$a0$copy, go$c = true, go$r;
    while (go$c) {
      const v = go$a0;
      if (v.tag === "Leaf") {
        go$c = false;
        go$r = Nothing;
        continue;
      }
      if (v.tag === "Node") {
        const v1 = ordEvalRef.compare(k)(v._3);
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
          go$r = $Maybe("Just", v._4);
          continue;
        }
      }
      fail();
    }
    return go$r;
  };
  return go;
};
var eqDistOp = (dictEq) => ({
  eq: (x) => (y) => {
    if (x.tag === "DistApp") {
      return y.tag === "DistApp" && eqArrayImpl(dictEq.eq)(x._1)(y._1);
    }
    if (x.tag === "DistUncurriedApp") {
      return y.tag === "DistUncurriedApp" && eqArrayImpl(dictEq.eq)(x._1)(y._1);
    }
    if (x.tag === "DistAccessor") {
      return y.tag === "DistAccessor" && eqBackendAccessor.eq(x._1)(y._1);
    }
    if (x.tag === "DistPrimOp1") {
      return y.tag === "DistPrimOp1" && (() => {
        if (x._1.tag === "OpBooleanNot") {
          return y._1.tag === "OpBooleanNot";
        }
        if (x._1.tag === "OpIntBitNot") {
          return y._1.tag === "OpIntBitNot";
        }
        if (x._1.tag === "OpIntNegate") {
          return y._1.tag === "OpIntNegate";
        }
        if (x._1.tag === "OpNumberNegate") {
          return y._1.tag === "OpNumberNegate";
        }
        if (x._1.tag === "OpArrayLength") {
          return y._1.tag === "OpArrayLength";
        }
        return x._1.tag === "OpIsTag" && y._1.tag === "OpIsTag" && (x._1._1._1.tag === "Nothing" ? y._1._1._1.tag === "Nothing" : x._1._1._1.tag === "Just" && y._1._1._1.tag === "Just" && x._1._1._1._1 === y._1._1._1._1) && x._1._1._2 === y._1._1._2;
      })();
    }
    if (x.tag === "DistPrimOp2L") {
      return y.tag === "DistPrimOp2L" && eqBackendOperator2.eq(x._1)(y._1) && dictEq.eq(x._2)(y._2);
    }
    return x.tag === "DistPrimOp2R" && y.tag === "DistPrimOp2R" && dictEq.eq(x._1)(y._1) && eqBackendOperator2.eq(x._2)(y._2);
  }
});
var eqBackendRewrite = (dictEq) => {
  const eq21 = eqArrayImpl((x) => (y) => dictEq.eq(x._1)(y._1) && dictEq.eq(x._2)(y._2));
  return {
    eq: (x) => (y) => {
      if (x.tag === "RewriteInline") {
        return y.tag === "RewriteInline" && (x._1.tag === "Nothing" ? y._1.tag === "Nothing" : x._1.tag === "Just" && y._1.tag === "Just" && x._1._1 === y._1._1) && x._2 === y._2 && dictEq.eq(x._3)(y._3) && dictEq.eq(x._4)(y._4);
      }
      if (x.tag === "RewriteUncurry") {
        return y.tag === "RewriteUncurry" && (x._1.tag === "Nothing" ? y._1.tag === "Nothing" : x._1.tag === "Just" && y._1.tag === "Just" && x._1._1 === y._1._1) && x._2 === y._2 && eq102(x._3)(y._3) && dictEq.eq(x._4)(y._4) && dictEq.eq(x._5)(y._5);
      }
      if (x.tag === "RewriteStop") {
        return y.tag === "RewriteStop" && (x._1._1.tag === "Nothing" ? y._1._1.tag === "Nothing" : x._1._1.tag === "Just" && y._1._1.tag === "Just" && x._1._1._1 === y._1._1._1) && x._1._2 === y._1._2;
      }
      if (x.tag === "RewriteUnpackOp") {
        return y.tag === "RewriteUnpackOp" && (x._1.tag === "Nothing" ? y._1.tag === "Nothing" : x._1.tag === "Just" && y._1.tag === "Just" && x._1._1 === y._1._1) && x._2 === y._2 && eqUnpackOp(dictEq).eq(x._3)(y._3) && dictEq.eq(x._4)(y._4);
      }
      if (x.tag === "RewriteDistBranchesLet") {
        return y.tag === "RewriteDistBranchesLet" && (x._1.tag === "Nothing" ? y._1.tag === "Nothing" : x._1.tag === "Just" && y._1.tag === "Just" && x._1._1 === y._1._1) && x._2 === y._2 && eq21(x._3)(y._3) && dictEq.eq(x._4)(y._4) && dictEq.eq(x._5)(y._5);
      }
      return x.tag === "RewriteDistBranchesOp" && y.tag === "RewriteDistBranchesOp" && eq21(x._1)(y._1) && dictEq.eq(x._2)(y._2) && eqDistOp(dictEq).eq(x._3)(y._3);
    }
  };
};
var eqBackendExpr = {
  eq: (v) => (v1) => {
    if (v.tag === "ExprSyntax") {
      return v1.tag === "ExprSyntax" && v._1.size === v1._1.size && eqBackendSyntax(eqBackendExpr).eq(v._2)(v1._2);
    }
    return v.tag === "ExprRewrite" && v1.tag === "ExprRewrite" && v._1.size === v1._1.size && eqBackendRewrite(eqBackendExpr).eq(v._2)(v1._2);
  }
};
var snocApp = (prev) => (next) => {
  const $0 = prev.length - 1 | 0;
  if ($0 >= 0 && $0 < prev.length && prev[$0].tag === "ExternApp") {
    return snoc((() => {
      const $1 = prev.length - 1 | 0;
      if ($1 < 1) {
        return [];
      }
      return sliceImpl(0, $1, prev);
    })())($ExternSpine("ExternApp", snoc(prev[$0]._1)(next)));
  }
  return snoc(prev)($ExternSpine("ExternApp", [next]));
};
var simplifyCondIsTag = (v) => (v1) => (v2) => {
  if (v1._1.tag === "ExprSyntax" && v1._1._2.tag === "PrimOp" && v1._1._2._1.tag === "Op1" && v1._1._2._1._1.tag === "OpIsTag" && v1._2.tag === "ExprSyntax" && v1._2._2.tag === "Lit" && v1._2._2._1.tag === "LitBoolean" && !v1._2._2._1._1 && v2.tag === "ExprSyntax" && v2._2.tag === "PrimOp" && v2._2._1.tag === "Op1" && v2._2._1._1.tag === "OpIsTag" && eqBackendExpr.eq(v1._1._2._1._2)(v2._2._1._2)) {
    return $Maybe("Just", v2);
  }
  return Nothing;
};
var shouldUnpackUpdate = (ident) => (level) => (binding2) => (body) => {
  const $0 = (() => {
    if (body.tag === "ExprSyntax") {
      return body._1;
    }
    if (body.tag === "ExprRewrite") {
      return body._1;
    }
    fail();
  })();
  if (binding2.tag === "ExprSyntax" && binding2._2.tag === "Update") {
    const $1 = lookup2(level)($0.usages);
    if ($1.tag === "Just" && $1._1.total === ($1._1.access + $1._1.update | 0)) {
      return $Maybe(
        "Just",
        $BackendExpr(
          "ExprRewrite",
          {
            ...updated(level)(semigroupBackendAnalysis.append((() => {
              if (binding2._2._1.tag === "ExprSyntax") {
                return binding2._2._1._1;
              }
              if (binding2._2._1.tag === "ExprRewrite") {
                return binding2._2._1._1;
              }
              fail();
            })())(foldrArray((x) => semigroupBackendAnalysis.append((() => {
              if (x._2.tag === "ExprSyntax") {
                return x._2._1;
              }
              if (x._2.tag === "ExprRewrite") {
                return x._2._1;
              }
              fail();
            })()))({
              ...$0,
              complexity: (() => {
                if ($0.complexity === "Trivial") {
                  return NonTrivial;
                }
                if ($0.complexity === "Deref") {
                  return NonTrivial;
                }
                if ($0.complexity === "KnownSize") {
                  return NonTrivial;
                }
                if ($0.complexity === "NonTrivial") {
                  return $0.complexity;
                }
                fail();
              })(),
              usages: $$delete(ordInt)(level)($0.usages)
            })(binding2._2._2))),
            rewrite: true
          },
          $BackendRewrite("RewriteUnpackOp", ident, level, $UnpackOp("UnpackUpdate", binding2._2._1, binding2._2._2), body)
        )
      );
    }
  }
  return Nothing;
};
var shouldUnpackRecord = (ident) => (level) => (binding2) => (body) => {
  const $0 = (() => {
    if (body.tag === "ExprSyntax") {
      return body._1;
    }
    if (body.tag === "ExprRewrite") {
      return body._1;
    }
    fail();
  })();
  if (binding2.tag === "ExprSyntax" && binding2._2.tag === "Lit" && binding2._2._1.tag === "LitRecord") {
    const $1 = lookup2(level)($0.usages);
    if ($1.tag === "Just" && $1._1.total === ($1._1.access + $1._1.update | 0)) {
      return $Maybe(
        "Just",
        $BackendExpr(
          "ExprRewrite",
          {
            ...foldrArray((x) => semigroupBackendAnalysis.append((() => {
              if (x._2.tag === "ExprSyntax") {
                return x._2._1;
              }
              if (x._2.tag === "ExprRewrite") {
                return x._2._1;
              }
              fail();
            })()))({
              ...$0,
              complexity: (() => {
                if ($0.complexity === "Trivial") {
                  return NonTrivial;
                }
                if ($0.complexity === "Deref") {
                  return NonTrivial;
                }
                if ($0.complexity === "KnownSize") {
                  return NonTrivial;
                }
                if ($0.complexity === "NonTrivial") {
                  return $0.complexity;
                }
                fail();
              })(),
              usages: $$delete(ordInt)(level)($0.usages)
            })(binding2._2._1._1),
            rewrite: true
          },
          $BackendRewrite("RewriteUnpackOp", ident, level, $UnpackOp("UnpackRecord", binding2._2._1._1), body)
        )
      );
    }
  }
  return Nothing;
};
var shouldUnpackCtor = (ident) => (level) => (a) => (body) => {
  const $0 = (() => {
    if (body.tag === "ExprSyntax") {
      return body._1;
    }
    if (body.tag === "ExprRewrite") {
      return body._1;
    }
    fail();
  })();
  if (a.tag === "ExprSyntax" && a._2.tag === "CtorSaturated") {
    const $1 = lookup2(level)($0.usages);
    if ($1.tag === "Just" && $1._1.total === ($1._1.access + $1._1.case | 0)) {
      return $Maybe(
        "Just",
        $BackendExpr(
          "ExprRewrite",
          {
            ...foldrArray((x) => semigroupBackendAnalysis.append((() => {
              if (x._2.tag === "ExprSyntax") {
                return x._2._1;
              }
              if (x._2.tag === "ExprRewrite") {
                return x._2._1;
              }
              fail();
            })()))({
              ...$0,
              complexity: (() => {
                if ($0.complexity === "Trivial") {
                  return NonTrivial;
                }
                if ($0.complexity === "Deref") {
                  return NonTrivial;
                }
                if ($0.complexity === "KnownSize") {
                  return NonTrivial;
                }
                if ($0.complexity === "NonTrivial") {
                  return $0.complexity;
                }
                fail();
              })(),
              usages: $$delete(ordInt)(level)($0.usages)
            })(a._2._5),
            rewrite: true
          },
          $BackendRewrite("RewriteUnpackOp", ident, level, $UnpackOp("UnpackData", a._2._1, a._2._2, a._2._3, a._2._4, a._2._5), body)
        )
      );
    }
  }
  return Nothing;
};
var shouldUnpackArray = (ident) => (level) => (binding2) => (body) => {
  const $0 = (() => {
    if (body.tag === "ExprSyntax") {
      return body._1;
    }
    if (body.tag === "ExprRewrite") {
      return body._1;
    }
    fail();
  })();
  if (binding2.tag === "ExprSyntax" && binding2._2.tag === "Lit" && binding2._2._1.tag === "LitArray") {
    const $1 = lookup2(level)($0.usages);
    if ($1.tag === "Just" && $1._1.total === $1._1.access) {
      return $Maybe(
        "Just",
        $BackendExpr(
          "ExprRewrite",
          {
            ...foldrArray((x) => semigroupBackendAnalysis.append((() => {
              if (x.tag === "ExprSyntax") {
                return x._1;
              }
              if (x.tag === "ExprRewrite") {
                return x._1;
              }
              fail();
            })()))({
              ...$0,
              complexity: (() => {
                if ($0.complexity === "Trivial") {
                  return NonTrivial;
                }
                if ($0.complexity === "Deref") {
                  return NonTrivial;
                }
                if ($0.complexity === "KnownSize") {
                  return NonTrivial;
                }
                if ($0.complexity === "NonTrivial") {
                  return $0.complexity;
                }
                fail();
              })(),
              usages: $$delete(ordInt)(level)($0.usages)
            })(binding2._2._1._1),
            rewrite: true
          },
          $BackendRewrite("RewriteUnpackOp", ident, level, $UnpackOp("UnpackArray", binding2._2._1._1), body)
        )
      );
    }
  }
  return Nothing;
};
var shouldUncurryAbs = (ident) => (level) => (a) => (b) => {
  const $0 = (() => {
    if (b.tag === "ExprSyntax") {
      return b._1;
    }
    if (b.tag === "ExprRewrite") {
      return b._1;
    }
    fail();
  })();
  if (a.tag === "ExprSyntax" && a._2.tag === "Abs") {
    const $1 = lookup2(level)($0.usages);
    if ($1.tag === "Just") {
      const $2 = toUnfoldable($1._1.arities);
      if ($2.length === 1 && $2[0] === a._2._1.length) {
        return $Maybe(
          "Just",
          $BackendExpr(
            "ExprRewrite",
            (() => {
              const $3 = semigroupBackendAnalysis.append((() => {
                if (a.tag === "ExprSyntax") {
                  return a._1;
                }
                if (a.tag === "ExprRewrite") {
                  return a._1;
                }
                fail();
              })())((() => {
                const $32 = (() => {
                  if (b.tag === "ExprSyntax") {
                    return b._1;
                  }
                  if (b.tag === "ExprRewrite") {
                    return b._1;
                  }
                  fail();
                })();
                return { ...$32, usages: $$delete(ordInt)(level)($32.usages) };
              })());
              return {
                ...$3,
                complexity: (() => {
                  if ($3.complexity === "Trivial") {
                    return NonTrivial;
                  }
                  if ($3.complexity === "Deref") {
                    return NonTrivial;
                  }
                  if ($3.complexity === "KnownSize") {
                    return NonTrivial;
                  }
                  if ($3.complexity === "NonTrivial") {
                    return $3.complexity;
                  }
                  fail();
                })(),
                result: (() => {
                  if (b.tag === "ExprSyntax") {
                    return b._1.result;
                  }
                  if (b.tag === "ExprRewrite") {
                    return b._1.result;
                  }
                  fail();
                })(),
                rewrite: true,
                size: $3.size + 1 | 0
              };
            })(),
            $BackendRewrite("RewriteUncurry", ident, level, a._2._1, a._2._2, b)
          )
        );
      }
    }
  }
  return Nothing;
};
var shouldInlineExternLiteral = (v) => {
  if (v.tag === "LitInt") {
    return true;
  }
  if (v.tag === "LitNumber") {
    return true;
  }
  if (v.tag === "LitString") {
    return toCodePointArray(v._1).length <= 32;
  }
  if (v.tag === "LitChar") {
    return true;
  }
  if (v.tag === "LitBoolean") {
    return true;
  }
  if (v.tag === "LitArray") {
    return v._1.length === 0;
  }
  if (v.tag === "LitRecord") {
    return v._1.length === 0;
  }
  fail();
};
var shouldInlineExternAppArg = (v) => (v1) => v1.tag === "SemLam" && (v.captured === "CaptureNone" || v.captured === "CaptureBranch") && v.total > 0 && v.call === v.total;
var shouldInlineExternApp = (v) => (v1) => (v2) => (args) => {
  const delayed = v1.args.length > 0;
  return (v1.complexity === "Trivial" || v1.complexity === "Deref") && v1.size < 16 || v1.usages.tag === "Leaf" && !v1.externs && v1.size < 64 || delayed && v1.args.length <= args.length && v1.size < 16 || delayed && or2(zipWithImpl(
    shouldInlineExternAppArg,
    v1.args,
    args
  )) && v1.size < 16;
};
var shouldEtaReduce = (level1) => (binding2) => (v) => {
  if (v.tag === "ExprSyntax" && v._2.tag === "Abs" && v._2._2.tag === "ExprSyntax" && v._2._2._2.tag === "App" && v._2._2._2._1.tag === "ExprSyntax" && v._2._2._2._1._2.tag === "Local") {
    const $0 = v._2._2._2._2;
    if (level1 === v._2._2._2._1._2._2 && v._2._1.length === $0.length && and2(zipWithImpl(
      (v$1) => {
        const $1 = v$1._2;
        return (v1) => v1.tag === "ExprSyntax" && v1._2.tag === "Local" && $1 === v1._2._2;
      },
      v._2._1,
      $0
    ))) {
      return $Maybe("Just", binding2);
    }
  }
  return Nothing;
};
var shouldDistributeBranches = (ident) => (level) => (a) => (body) => {
  const $0 = (() => {
    if (body.tag === "ExprSyntax") {
      return body._1;
    }
    if (body.tag === "ExprRewrite") {
      return body._1;
    }
    fail();
  })();
  if (a.tag === "ExprSyntax" && a._2.tag === "Branch" && $0.size <= 128 && a._1.result === "KnownNeutral") {
    const $1 = lookup2(level)($0.usages);
    if ($1.tag === "Just" && $1._1.total === ($1._1.access + $1._1.case | 0)) {
      return $Maybe(
        "Just",
        $BackendExpr(
          "ExprRewrite",
          {
            ...semigroupBackendAnalysis.append((() => {
              if (a.tag === "ExprSyntax") {
                return a._1;
              }
              if (a.tag === "ExprRewrite") {
                return a._1;
              }
              fail();
            })())((() => {
              const $2 = (() => {
                if (body.tag === "ExprSyntax") {
                  return body._1;
                }
                if (body.tag === "ExprRewrite") {
                  return body._1;
                }
                fail();
              })();
              return { ...$2, usages: $$delete(ordInt)(level)($2.usages) };
            })()),
            rewrite: true
          },
          $BackendRewrite("RewriteDistBranchesLet", ident, level, a._2._1, a._2._2, body)
        )
      );
    }
  }
  return Nothing;
};
var shouldDistributeBranchUncurriedApps = (analysis1) => (branches) => (def) => (spine) => {
  if (allImpl(
    (x) => {
      const $0 = (() => {
        if (x.tag === "ExprSyntax") {
          return x._1;
        }
        if (x.tag === "ExprRewrite") {
          return x._1;
        }
        fail();
      })();
      return $0.complexity === "Trivial" || $0.complexity === "Deref";
    },
    spine
  )) {
    return $Maybe(
      "Just",
      $BackendExpr(
        "ExprRewrite",
        { ...semigroupBackendAnalysis.append(analysis1)(foldMap(hasAnalysisBackendExpr.analysisOf)(spine)), rewrite: true },
        $BackendRewrite("RewriteDistBranchesOp", branches, def, $DistOp("DistUncurriedApp", spine))
      )
    );
  }
  return Nothing;
};
var shouldDistributeBranchPrimOp2R = (analysis1) => (branches) => (def) => (lhs) => (op2) => {
  if ((() => {
    const $0 = (() => {
      if (lhs.tag === "ExprSyntax") {
        return lhs._1.complexity;
      }
      if (lhs.tag === "ExprRewrite") {
        return lhs._1.complexity;
      }
      fail();
    })();
    return $0 === "Trivial" || $0 === "Deref";
  })()) {
    return $Maybe(
      "Just",
      $BackendExpr(
        "ExprRewrite",
        (() => {
          const $0 = semigroupBackendAnalysis.append(analysis1)((() => {
            if (lhs.tag === "ExprSyntax") {
              return lhs._1;
            }
            if (lhs.tag === "ExprRewrite") {
              return lhs._1;
            }
            fail();
          })());
          return { ...$0, rewrite: true, size: $0.size + 1 | 0 };
        })(),
        $BackendRewrite("RewriteDistBranchesOp", branches, def, $DistOp("DistPrimOp2R", lhs, op2))
      )
    );
  }
  return Nothing;
};
var shouldDistributeBranchPrimOp2L = (analysis1) => (branches) => (def) => (op2) => (rhs) => {
  if ((() => {
    const $0 = (() => {
      if (rhs.tag === "ExprSyntax") {
        return rhs._1.complexity;
      }
      if (rhs.tag === "ExprRewrite") {
        return rhs._1.complexity;
      }
      fail();
    })();
    return $0 === "Trivial" || $0 === "Deref";
  })()) {
    return $Maybe(
      "Just",
      $BackendExpr(
        "ExprRewrite",
        (() => {
          const $0 = semigroupBackendAnalysis.append(analysis1)((() => {
            if (rhs.tag === "ExprSyntax") {
              return rhs._1;
            }
            if (rhs.tag === "ExprRewrite") {
              return rhs._1;
            }
            fail();
          })());
          return { ...$0, rewrite: true, size: $0.size + 1 | 0 };
        })(),
        $BackendRewrite("RewriteDistBranchesOp", branches, def, $DistOp("DistPrimOp2L", op2, rhs))
      )
    );
  }
  return Nothing;
};
var shouldDistributeBranchApps = (analysis1) => (branches) => (def) => (spine) => {
  if (allImpl(
    (x) => {
      const $0 = (() => {
        if (x.tag === "ExprSyntax") {
          return x._1;
        }
        if (x.tag === "ExprRewrite") {
          return x._1;
        }
        fail();
      })();
      return $0.complexity === "Trivial" || $0.complexity === "Deref";
    },
    spine
  )) {
    return $Maybe(
      "Just",
      $BackendExpr(
        "ExprRewrite",
        { ...semigroupBackendAnalysis.append(analysis1)(foldMap12(hasAnalysisBackendExpr.analysisOf)(spine)), rewrite: true },
        $BackendRewrite("RewriteDistBranchesOp", branches, def, $DistOp("DistApp", spine))
      )
    );
  }
  return Nothing;
};
var rewriteInline = (ident) => (level) => (binding2) => (body) => {
  const s2 = (() => {
    if (body.tag === "ExprSyntax") {
      return body._1;
    }
    if (body.tag === "ExprRewrite") {
      return body._1;
    }
    fail();
  })();
  return $BackendExpr(
    "ExprRewrite",
    (() => {
      const v = lookup2(level)(s2.usages);
      const $0 = (() => {
        if (v.tag === "Just") {
          return semigroupBackendAnalysis.append(s2)(power2((() => {
            if (binding2.tag === "ExprSyntax") {
              return binding2._1;
            }
            if (binding2.tag === "ExprRewrite") {
              return binding2._1;
            }
            fail();
          })())(v._1.total));
        }
        if (v.tag === "Nothing") {
          return s2;
        }
        fail();
      })();
      return { ...$0, rewrite: true, usages: $$delete(ordInt)(level)($0.usages) };
    })(),
    $BackendRewrite("RewriteInline", ident, level, binding2, body)
  );
};
var rewriteBranches = (k) => {
  const go = (v) => {
    if (v.tag === "SemLet") {
      return $BackendSemantics("SemLet", v._1, v._2, (x) => go(v._3(x)));
    }
    if (v.tag === "SemLetRec") {
      return $BackendSemantics("SemLetRec", v._1, (x) => go(v._2(x)));
    }
    if (v.tag === "SemBranch") {
      const $0 = v._2;
      return $BackendSemantics(
        "SemBranch",
        arrayMap((v1) => {
          const $1 = v1._2;
          return $SemConditional(v1._1, defer((v$1) => go(force($1))));
        })(v._1),
        defer((v$1) => go(force($0)))
      );
    }
    return k(v);
  };
  return go;
};
var neutralSpine = /* @__PURE__ */ foldlArray((hd) => (v) => {
  if (v.tag === "ExternApp") {
    return $BackendSemantics("NeutApp", hd, v._1);
  }
  if (v.tag === "ExternUncurriedApp") {
    return $BackendSemantics("NeutUncurriedApp", hd, v._1);
  }
  if (v.tag === "ExternAccessor") {
    return $BackendSemantics("NeutAccessor", hd, v._1);
  }
  if (v.tag === "ExternPrimOp") {
    return $BackendSemantics("NeutPrimOp", $BackendOperator("Op1", v._1, hd));
  }
  fail();
});
var shouldInlineLet = (level) => (a) => (b) => {
  const $0 = (() => {
    if (a.tag === "ExprSyntax") {
      return a._1;
    }
    if (a.tag === "ExprRewrite") {
      return a._1;
    }
    fail();
  })();
  const v2 = lookup2(level)((() => {
    if (b.tag === "ExprSyntax") {
      return b._1.usages;
    }
    if (b.tag === "ExprRewrite") {
      return b._1.usages;
    }
    fail();
  })());
  if (v2.tag === "Nothing") {
    return true;
  }
  if (v2.tag === "Just") {
    return $0.complexity === "Trivial" || v2._1.captured === "CaptureNone" && v2._1.total === 1 || (v2._1.captured === "CaptureNone" || v2._1.captured === "CaptureBranch") && ($0.complexity === "Trivial" || $0.complexity === "Deref") && $0.size < 5 || $0.complexity === "Deref" && v2._1.call === v2._1.total || $0.complexity === "KnownSize" && v2._1.total === 1 || a.tag === "ExprSyntax" && (a._2.tag === "Abs" || a._2.tag === "UncurriedAbs" || a._2.tag === "UncurriedEffectAbs" || a._2.tag === "EffectDefer") && (v2._1.total === 1 || $0.usages.tag === "Leaf" || $0.size < 16) || a.tag === "ExprSyntax" && (a._2.tag === "PrimEffect" || a._2.tag === "UncurriedEffectApp" || a._2.tag === "EffectBind" || a._2.tag === "EffectDefer") && v2._1.total === 1;
  }
  fail();
};
var insertDirective = (ref) => (acc) => (dir) => alter2((v) => {
  if (v.tag === "Just") {
    return $Maybe("Just", insert(ordInlineAccessor)(acc)(dir)(v._1));
  }
  if (v.tag === "Nothing") {
    return $Maybe("Just", $$$Map("Node", 1, 1, acc, dir, Leaf, Leaf));
  }
  fail();
})(ref);
var guardFailOver1 = (f) => (as) => (k) => {
  const v = foldlArray((v2) => (v1) => {
    if (v2.tag === "Nothing") {
      const $0 = f(v1);
      if ($0.tag === "NeutFail") {
        return $Maybe("Just", $0);
      }
      return Nothing;
    }
    return v2;
  })(Nothing)(as);
  if (v.tag === "Just") {
    return v._1;
  }
  if (v.tag === "Nothing") {
    return k(as);
  }
  fail();
};
var guardFailOver2 = (f) => (as) => (k) => {
  const v = foldlDefault(foldableBackendEffect)((v2) => (v1) => {
    if (v2.tag === "Nothing") {
      const $0 = f(v1);
      if ($0.tag === "NeutFail") {
        return $Maybe("Just", $0);
      }
      return Nothing;
    }
    return v2;
  })(Nothing)(as);
  if (v.tag === "Just") {
    return v._1;
  }
  if (v.tag === "Nothing") {
    return k(as);
  }
  fail();
};
var guardFailOver3 = (f) => (as) => (k) => {
  const v = foldlDefault(foldableLiteral)((v2) => (v1) => {
    if (v2.tag === "Nothing") {
      const $0 = f(v1);
      if ($0.tag === "NeutFail") {
        return $Maybe("Just", $0);
      }
      return Nothing;
    }
    return v2;
  })(Nothing)(as);
  if (v.tag === "Just") {
    return v._1;
  }
  if (v.tag === "Nothing") {
    return k(as);
  }
  fail();
};
var foldBackendExpr = (foldSyntax) => (foldRewrite) => {
  const go = (v) => {
    if (v.tag === "ExprSyntax") {
      return foldSyntax(functorBackendSyntax.map(go)(v._2));
    }
    if (v.tag === "ExprRewrite") {
      return foldRewrite(v._2)((() => {
        if (v._2.tag === "RewriteInline") {
          return foldSyntax($BackendSyntax("Let", v._2._1, v._2._2, go(v._2._3), go(v._2._4)));
        }
        if (v._2.tag === "RewriteUncurry") {
          return foldSyntax($BackendSyntax(
            "Let",
            v._2._1,
            v._2._2,
            foldSyntax($BackendSyntax("Abs", v._2._3, go(v._2._4))),
            go(v._2._5)
          ));
        }
        if (v._2.tag === "RewriteStop") {
          return foldSyntax($BackendSyntax("Var", v._2._1));
        }
        if (v._2.tag === "RewriteUnpackOp") {
          if (v._2._3.tag === "UnpackRecord") {
            return foldSyntax($BackendSyntax(
              "Let",
              v._2._1,
              v._2._2,
              foldSyntax($BackendSyntax(
                "Lit",
                $Literal("LitRecord", arrayMap((m) => $Prop(m._1, go(m._2)))(v._2._3._1))
              )),
              go(v._2._4)
            ));
          }
          if (v._2._3.tag === "UnpackUpdate") {
            return foldSyntax($BackendSyntax(
              "Let",
              v._2._1,
              v._2._2,
              foldSyntax($BackendSyntax(
                "Update",
                go(v._2._3._1),
                arrayMap((m) => $Prop(m._1, go(m._2)))(v._2._3._2)
              )),
              go(v._2._4)
            ));
          }
          if (v._2._3.tag === "UnpackArray") {
            return foldSyntax($BackendSyntax(
              "Let",
              v._2._1,
              v._2._2,
              foldSyntax($BackendSyntax(
                "Lit",
                $Literal("LitArray", arrayMap(go)(v._2._3._1))
              )),
              go(v._2._4)
            ));
          }
          if (v._2._3.tag === "UnpackData") {
            return foldSyntax($BackendSyntax(
              "Let",
              v._2._1,
              v._2._2,
              foldSyntax($BackendSyntax(
                "CtorSaturated",
                v._2._3._1,
                v._2._3._2,
                v._2._3._3,
                v._2._3._4,
                arrayMap((m) => $Tuple(m._1, go(m._2)))(v._2._3._5)
              )),
              go(v._2._4)
            ));
          }
          fail();
        }
        if (v._2.tag === "RewriteDistBranchesLet") {
          return foldSyntax($BackendSyntax(
            "Let",
            v._2._1,
            v._2._2,
            foldSyntax($BackendSyntax(
              "Branch",
              arrayMap((m) => $Pair(go(m._1), go(m._2)))(v._2._3),
              go(v._2._4)
            )),
            go(v._2._5)
          ));
        }
        if (v._2.tag === "RewriteDistBranchesOp") {
          const branches$p = foldSyntax($BackendSyntax(
            "Branch",
            arrayMap((m) => $Pair(go(m._1), go(m._2)))(v._2._1),
            go(v._2._2)
          ));
          if (v._2._3.tag === "DistApp") {
            return foldSyntax($BackendSyntax("App", branches$p, arrayMap(go)(v._2._3._1)));
          }
          if (v._2._3.tag === "DistUncurriedApp") {
            return foldSyntax($BackendSyntax("UncurriedApp", branches$p, arrayMap(go)(v._2._3._1)));
          }
          if (v._2._3.tag === "DistAccessor") {
            return foldSyntax($BackendSyntax("Accessor", branches$p, v._2._3._1));
          }
          if (v._2._3.tag === "DistPrimOp1") {
            return foldSyntax($BackendSyntax(
              "PrimOp",
              $BackendOperator("Op1", v._2._3._1, branches$p)
            ));
          }
          if (v._2._3.tag === "DistPrimOp2L") {
            return foldSyntax($BackendSyntax(
              "PrimOp",
              $BackendOperator("Op2", v._2._3._1, branches$p, go(v._2._3._2))
            ));
          }
          if (v._2._3.tag === "DistPrimOp2R") {
            return foldSyntax($BackendSyntax(
              "PrimOp",
              $BackendOperator("Op2", v._2._3._2, go(v._2._3._1), branches$p)
            ));
          }
        }
        fail();
      })());
    }
    fail();
  };
  return go;
};
var freeze = (init) => $Tuple(
  (() => {
    if (init.tag === "ExprSyntax") {
      return init._1;
    }
    if (init.tag === "ExprRewrite") {
      return init._1;
    }
    fail();
  })(),
  foldBackendExpr(NeutralExpr)((v) => (neutExpr) => neutExpr)(init)
);
var floatLetWith$lazy = /* @__PURE__ */ binding(() => {
  const go = (go$a0$copy) => (go$a1$copy) => (go$a2$copy) => (go$a3$copy) => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$a2 = go$a2$copy, go$a3 = go$a3$copy, go$c = true, go$r;
    while (go$c) {
      const f = go$a0, ident1 = go$a1, binding1 = go$a2, k1 = go$a3;
      if (binding1.tag === "SemLet") {
        go$a0 = makeLet$lazy();
        go$a1 = binding1._1;
        go$a2 = binding1._2;
        go$a3 = (nextBinding2) => f(ident1)(binding1._3(nextBinding2))(k1);
        continue;
      }
      if (binding1.tag === "SemLetRec") {
        go$c = false;
        go$r = $BackendSemantics("SemLetRec", binding1._1, (nextBindings) => makeLet$lazy()(ident1)(binding1._2(nextBindings))(k1));
        continue;
      }
      if (binding1.tag === "NeutFail") {
        go$c = false;
        go$r = binding1;
        continue;
      }
      go$c = false;
      go$r = f(ident1)(binding1)(k1);
    }
    return go$r;
  };
  return go;
});
var makeLet$lazy = /* @__PURE__ */ binding(() => floatLetWith$lazy()((ident) => (binding2) => (k) => {
  if (binding2.tag === "SemRef") {
    if (binding2._2.length === 0) {
      return k(binding2);
    }
    return $BackendSemantics("SemLet", ident, binding2, k);
  }
  if (binding2.tag === "NeutLocal") {
    return k(binding2);
  }
  if (binding2.tag === "NeutStop") {
    return k(binding2);
  }
  if (binding2.tag === "NeutVar") {
    return k(binding2);
  }
  return $BackendSemantics("SemLet", ident, binding2, k);
}));
var floatLetWith = /* @__PURE__ */ floatLetWith$lazy();
var makeLet = /* @__PURE__ */ makeLet$lazy();
var floatLet = /* @__PURE__ */ floatLetWith((v) => applyFlipped)(Nothing);
var makeEffectBind$lazy = /* @__PURE__ */ binding(() => {
  const go = (go$a0$copy) => (go$a1$copy) => (go$a2$copy) => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$a2 = go$a2$copy, go$c = true, go$r;
    while (go$c) {
      const ident1 = go$a0, binding1 = go$a1, k1 = go$a2;
      if (binding1.tag === "SemLet") {
        go$c = false;
        go$r = makeLet(binding1._1)(binding1._2)((nextBinding2) => makeEffectBind$lazy()(ident1)(binding1._3(nextBinding2))(k1));
        continue;
      }
      if (binding1.tag === "SemEffectBind") {
        go$a0 = binding1._1;
        go$a1 = binding1._2;
        go$a2 = (nextBinding2) => makeEffectBind$lazy()(ident1)(binding1._3(nextBinding2))(k1);
        continue;
      }
      if (binding1.tag === "SemEffectDefer") {
        go$c = false;
        go$r = $BackendSemantics("SemEffectDefer", floatLet(binding1._1)((nextBinding2) => makeEffectBind$lazy()(ident1)(nextBinding2)(k1)));
        continue;
      }
      go$c = false;
      go$r = floatLet(binding1)((nextBinding2) => $BackendSemantics("SemEffectBind", ident1, nextBinding2, k1));
    }
    return go$r;
  };
  return go;
});
var makeEffectBind = /* @__PURE__ */ makeEffectBind$lazy();
var evalUpdate = (lhs) => (props) => floatLet(lhs)((v) => {
  if (v.tag === "NeutLit") {
    if (v._1.tag === "LitRecord") {
      return $BackendSemantics(
        "NeutLit",
        $Literal(
          "LitRecord",
          arrayMap(head)(groupAllBy((x) => (y) => ordString.compare(x._1)(y._1))([...props, ...v._1._1]))
        )
      );
    }
    return $BackendSemantics("NeutUpdate", v, props);
  }
  if (v.tag === "NeutUpdate") {
    return $BackendSemantics(
      "NeutUpdate",
      v._1,
      arrayMap(head)(groupAllBy((x) => (y) => ordString.compare(x._1)(y._1))([...props, ...v._2]))
    );
  }
  return $BackendSemantics("NeutUpdate", v, props);
});
var evalUncurriedBeta = (fn) => (mk) => (spine) => {
  const go = (v) => (v1) => {
    if (v.tag === "MkFnNext") {
      if (v1.tag === "Cons") {
        if (v1._1.tag === "NeutFail") {
          return $BackendSemantics("NeutFail", v1._1._1);
        }
        const $0 = v1._2;
        return makeLet(Nothing)(v1._1)((nextArg) => go(v._2(nextArg))($0));
      }
      return _crashWith("Uncurried function applied to too few arguments");
    }
    if (v.tag === "MkFnApplied") {
      if (v1.tag === "Nil") {
        return v._1;
      }
      return fn(v._1)(toUnfoldable1(v1));
    }
    fail();
  };
  return go(mk)(fromFoldable2(spine));
};
var evalPrimOpOrd = (dictOrd) => {
  const Eq0 = dictOrd.Eq0();
  return (op) => (x) => (y) => {
    if (op === "OpEq") {
      return Eq0.eq(x)(y);
    }
    if (op === "OpNotEq") {
      return !Eq0.eq(x)(y);
    }
    if (op === "OpGt") {
      return dictOrd.compare(x)(y) === "GT";
    }
    if (op === "OpGte") {
      return dictOrd.compare(x)(y) !== "LT";
    }
    if (op === "OpLt") {
      return dictOrd.compare(x)(y) === "LT";
    }
    if (op === "OpLte") {
      return dictOrd.compare(x)(y) !== "GT";
    }
    fail();
  };
};
var evalPrimOpOrd1 = /* @__PURE__ */ evalPrimOpOrd(ordString);
var evalPrimOpOrd2 = /* @__PURE__ */ evalPrimOpOrd(ordInt);
var evalPrimOpOrd3 = /* @__PURE__ */ evalPrimOpOrd(ordChar);
var evalPrimOpOrd4 = /* @__PURE__ */ evalPrimOpOrd(ordBoolean);
var evalPrimOpNot = (v) => {
  if (v.tag === "Op1") {
    if (v._1.tag === "OpBooleanNot") {
      return v._2;
    }
    return $BackendSemantics(
      "NeutPrimOp",
      $BackendOperator(
        "Op1",
        OpBooleanNot,
        $BackendSemantics("NeutPrimOp", $BackendOperator("Op1", v._1, v._2))
      )
    );
  }
  if (v.tag === "Op2") {
    if (v._1.tag === "OpIntOrd") {
      return $BackendSemantics(
        "NeutPrimOp",
        $BackendOperator(
          "Op2",
          $BackendOperator2(
            "OpIntOrd",
            (() => {
              if (v._1._1 === "OpEq") {
                return OpNotEq;
              }
              if (v._1._1 === "OpNotEq") {
                return OpEq;
              }
              if (v._1._1 === "OpLt") {
                return OpGte;
              }
              if (v._1._1 === "OpLte") {
                return OpGt;
              }
              if (v._1._1 === "OpGt") {
                return OpLte;
              }
              if (v._1._1 === "OpGte") {
                return OpLt;
              }
              fail();
            })()
          ),
          v._2,
          v._3
        )
      );
    }
    if (v._1.tag === "OpNumberOrd") {
      return $BackendSemantics(
        "NeutPrimOp",
        $BackendOperator(
          "Op2",
          $BackendOperator2(
            "OpNumberOrd",
            (() => {
              if (v._1._1 === "OpEq") {
                return OpNotEq;
              }
              if (v._1._1 === "OpNotEq") {
                return OpEq;
              }
              if (v._1._1 === "OpLt") {
                return OpGte;
              }
              if (v._1._1 === "OpLte") {
                return OpGt;
              }
              if (v._1._1 === "OpGt") {
                return OpLte;
              }
              if (v._1._1 === "OpGte") {
                return OpLt;
              }
              fail();
            })()
          ),
          v._2,
          v._3
        )
      );
    }
    if (v._1.tag === "OpStringOrd") {
      return $BackendSemantics(
        "NeutPrimOp",
        $BackendOperator(
          "Op2",
          $BackendOperator2(
            "OpStringOrd",
            (() => {
              if (v._1._1 === "OpEq") {
                return OpNotEq;
              }
              if (v._1._1 === "OpNotEq") {
                return OpEq;
              }
              if (v._1._1 === "OpLt") {
                return OpGte;
              }
              if (v._1._1 === "OpLte") {
                return OpGt;
              }
              if (v._1._1 === "OpGt") {
                return OpLte;
              }
              if (v._1._1 === "OpGte") {
                return OpLt;
              }
              fail();
            })()
          ),
          v._2,
          v._3
        )
      );
    }
    if (v._1.tag === "OpCharOrd") {
      return $BackendSemantics(
        "NeutPrimOp",
        $BackendOperator(
          "Op2",
          $BackendOperator2(
            "OpCharOrd",
            (() => {
              if (v._1._1 === "OpEq") {
                return OpNotEq;
              }
              if (v._1._1 === "OpNotEq") {
                return OpEq;
              }
              if (v._1._1 === "OpLt") {
                return OpGte;
              }
              if (v._1._1 === "OpLte") {
                return OpGt;
              }
              if (v._1._1 === "OpGt") {
                return OpLte;
              }
              if (v._1._1 === "OpGte") {
                return OpLt;
              }
              fail();
            })()
          ),
          v._2,
          v._3
        )
      );
    }
    if (v._1.tag === "OpBooleanOrd") {
      return $BackendSemantics(
        "NeutPrimOp",
        $BackendOperator(
          "Op2",
          $BackendOperator2(
            "OpBooleanOrd",
            (() => {
              if (v._1._1 === "OpEq") {
                return OpNotEq;
              }
              if (v._1._1 === "OpNotEq") {
                return OpEq;
              }
              if (v._1._1 === "OpLt") {
                return OpGte;
              }
              if (v._1._1 === "OpLte") {
                return OpGt;
              }
              if (v._1._1 === "OpGt") {
                return OpLte;
              }
              if (v._1._1 === "OpGte") {
                return OpLt;
              }
              fail();
            })()
          ),
          v._2,
          v._3
        )
      );
    }
    return $BackendSemantics(
      "NeutPrimOp",
      $BackendOperator(
        "Op1",
        OpBooleanNot,
        $BackendSemantics("NeutPrimOp", $BackendOperator("Op2", v._1, v._2, v._3))
      )
    );
  }
  fail();
};
var evalPair = (dictEval) => (env) => (v) => {
  const $0 = v._1;
  const $1 = v._2;
  return $SemConditional(defer((v1) => dictEval.eval(env)($0)), defer((v1) => dictEval.eval(env)($1)));
};
var evalBranches = (v) => (initConds) => (initDef) => {
  const go = (go$a0$copy) => (go$a1$copy) => (go$a2$copy) => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$a2 = go$a2$copy, go$c = true, go$r;
    while (go$c) {
      const acc = go$a0, conds = go$a1, def = go$a2;
      const v1 = unconsImpl((v$1) => Nothing, (x) => (xs) => $Maybe("Just", { head: x, tail: xs }), conds);
      if (v1.tag === "Just") {
        const $0 = force(v1._1.head._1);
        const v2 = $0.tag === "SemRef" ? force($0._3) : $0;
        if (v2.tag === "NeutLit") {
          if (v2._1.tag === "LitBoolean") {
            if (v2._1._1) {
              go$a0 = acc;
              go$a1 = [];
              go$a2 = v1._1.head._2;
              continue;
            }
            go$a0 = acc;
            go$a1 = v1._1.tail;
            go$a2 = def;
            continue;
          }
          go$a0 = snoc(acc)(v1._1.head);
          go$a1 = v1._1.tail;
          go$a2 = def;
          continue;
        }
        if (v2.tag === "NeutFail") {
          const $1 = v2._1;
          go$a0 = acc;
          go$a1 = [];
          go$a2 = defer((v3) => $BackendSemantics("NeutFail", $1));
          continue;
        }
        go$a0 = snoc(acc)(v1._1.head);
        go$a1 = v1._1.tail;
        go$a2 = def;
        continue;
      }
      if (v1.tag === "Nothing") {
        if (acc.length > 0) {
          go$c = false;
          go$r = $BackendSemantics("SemBranch", acc, def);
          continue;
        }
        go$c = false;
        go$r = force(def);
        continue;
      }
      fail();
    }
    return go$r;
  };
  return go([])(initConds)(initDef);
};
var evalPrimOpNumInt = (op) => (x) => (y) => {
  const $0 = x.tag === "SemRef" ? force(x._3) : x;
  if ($0.tag === "NeutLit" && $0._1.tag === "LitInt") {
    const $1 = y.tag === "SemRef" ? force(y._3) : y;
    if ($1.tag === "NeutLit" && $1._1.tag === "LitInt") {
      if (op === "OpAdd") {
        const res = $0._1._1 + $1._1._1 | 0;
        if ($1._1._1 > 0 && res < $0._1._1 || $1._1._1 < 0 && res > $0._1._1) {
          return Nothing;
        }
        return $Maybe("Just", $BackendSemantics("NeutLit", $Literal("LitInt", res)));
      }
      if (op === "OpMultiply") {
        const res = $0._1._1 * $1._1._1 | 0;
        if ($0._1._1 !== intDiv(res, $1._1._1)) {
          return Nothing;
        }
        return $Maybe("Just", $BackendSemantics("NeutLit", $Literal("LitInt", res)));
      }
      if (op === "OpSubtract") {
        const res = $0._1._1 - $1._1._1 | 0;
        if ($1._1._1 > 0 && res > $0._1._1 || $1._1._1 < 0 && res < $0._1._1) {
          return Nothing;
        }
        return $Maybe("Just", $BackendSemantics("NeutLit", $Literal("LitInt", res)));
      }
      if (op === "OpDivide") {
        return $Maybe("Just", $BackendSemantics("NeutLit", $Literal("LitInt", intDiv($0._1._1, $1._1._1))));
      }
      fail();
    }
  }
  return Nothing;
};
var evalPrimOpNumNumber = (op) => (x) => (y) => {
  const $0 = x.tag === "SemRef" ? force(x._3) : x;
  if ($0.tag === "NeutLit" && $0._1.tag === "LitNumber") {
    const $1 = y.tag === "SemRef" ? force(y._3) : y;
    if ($1.tag === "NeutLit" && $1._1.tag === "LitNumber") {
      return $Maybe(
        "Just",
        $BackendSemantics(
          "NeutLit",
          $Literal(
            "LitNumber",
            (() => {
              if (op === "OpAdd") {
                return $0._1._1 + $1._1._1;
              }
              if (op === "OpMultiply") {
                return $0._1._1 * $1._1._1;
              }
              if (op === "OpSubtract") {
                return $0._1._1 - $1._1._1;
              }
              if (op === "OpDivide") {
                return $0._1._1 / $1._1._1;
              }
              fail();
            })()
          )
        )
      );
    }
  }
  return Nothing;
};
var evalRefSpine = (env) => (ref) => (spine) => (sem) => (v) => {
  if (v.tag === "ExternApp") {
    return neutralSpine((() => {
      if (ref.tag === "EvalExtern") {
        return $BackendSemantics("NeutVar", ref._1);
      }
      if (ref.tag === "EvalLocal") {
        return $BackendSemantics("NeutLocal", ref._1, ref._2);
      }
      fail();
    })())(spine);
  }
  if (v.tag === "ExternUncurriedApp") {
    return neutralSpine((() => {
      if (ref.tag === "EvalExtern") {
        return $BackendSemantics("NeutVar", ref._1);
      }
      if (ref.tag === "EvalLocal") {
        return $BackendSemantics("NeutLocal", ref._1, ref._2);
      }
      fail();
    })())(spine);
  }
  if (v.tag === "ExternAccessor") {
    return evalAccessor(env)(force(sem))(v._1);
  }
  if (v.tag === "ExternPrimOp") {
    return evalPrimOp(env)($BackendOperator("Op1", v._1, force(sem)));
  }
  fail();
};
var evalRef = (v) => (ref) => (spine) => (last) => (sem) => {
  const spine$p = last.tag === "ExternApp" ? foldlArray(snocApp)(spine)(last._1) : snoc(spine)(last);
  const v1 = (v2) => $BackendSemantics(
    "SemRef",
    ref,
    spine$p,
    defer((v3) => {
      const $0 = evalRefSpine(v)(ref)(spine$p)(sem)(last);
      if ($0.tag === "SemRef") {
        return force($0._3);
      }
      return $0;
    })
  );
  if (ref.tag === "EvalExtern") {
    const $0 = v.evalExternSpine(v)(ref._1)(spine$p);
    if ($0.tag === "Just") {
      return $0._1;
    }
  }
  return v1(true);
};
var evalPrimOp = (env) => (v) => {
  if (v.tag === "Op1") {
    const $0 = v._1;
    const $1 = v._2;
    const v1 = (v2) => {
      if ($0.tag === "OpBooleanNot" && $1.tag === "NeutPrimOp") {
        return evalPrimOpNot($1._1);
      }
      const v5 = (v6) => {
        const v7 = (v8) => {
          const v9 = (v10) => {
            const v11 = (v12) => {
              const v13 = (v14) => {
                if ($1.tag === "SemRef") {
                  return evalRef(env)($1._1)($1._2)($ExternSpine("ExternPrimOp", $0))($1._3);
                }
                if ($1.tag === "NeutFail") {
                  return $BackendSemantics("NeutFail", $1._1);
                }
                return floatLet($1)((() => {
                  const $2 = Op1($0);
                  return (x) => $BackendSemantics("NeutPrimOp", $2(x));
                })());
              };
              if ($0.tag === "OpNumberNegate") {
                const $2 = $1.tag === "SemRef" ? force($1._3) : $1;
                if ($2.tag === "NeutLit" && $2._1.tag === "LitNumber") {
                  return $BackendSemantics("NeutLit", $Literal("LitNumber", -$2._1._1));
                }
              }
              return v13(true);
            };
            if ($0.tag === "OpIntNegate") {
              const $2 = $1.tag === "SemRef" ? force($1._3) : $1;
              if ($2.tag === "NeutLit" && $2._1.tag === "LitInt") {
                return $BackendSemantics("NeutLit", $Literal("LitInt", -$2._1._1));
              }
            }
            return v11(true);
          };
          if ($0.tag === "OpArrayLength") {
            const $2 = $1.tag === "SemRef" ? force($1._3) : $1;
            if ($2.tag === "NeutLit" && $2._1.tag === "LitArray") {
              return $BackendSemantics("NeutLit", $Literal("LitInt", $2._1._1.length));
            }
          }
          return v9(true);
        };
        if ($0.tag === "OpIsTag") {
          const $2 = $1.tag === "SemRef" ? force($1._3) : $1;
          if ($2.tag === "NeutData") {
            return $BackendSemantics(
              "NeutLit",
              $Literal(
                "LitBoolean",
                ($0._1._1.tag === "Nothing" ? $2._1._1.tag === "Nothing" : $0._1._1.tag === "Just" && $2._1._1.tag === "Just" && $0._1._1._1 === $2._1._1._1) && $0._1._2 === $2._1._2
              )
            );
          }
        }
        return v7(true);
      };
      if ($0.tag === "OpIntBitNot") {
        const $2 = $1.tag === "SemRef" ? force($1._3) : $1;
        if ($2.tag === "NeutLit" && $2._1.tag === "LitInt") {
          return $BackendSemantics("NeutLit", $Literal("LitInt", ~$2._1._1));
        }
      }
      return v5(true);
    };
    if ($0.tag === "OpBooleanNot") {
      const $2 = $1.tag === "SemRef" ? force($1._3) : $1;
      if ($2.tag === "NeutLit" && $2._1.tag === "LitBoolean") {
        return $BackendSemantics("NeutLit", $Literal("LitBoolean", !$2._1._1));
      }
    }
    return v1(true);
  }
  if (v.tag === "Op2") {
    const $0 = v._1;
    const $1 = v._2;
    const $2 = v._3;
    const v1 = (v2) => {
      const v3 = (v4) => {
        const v5 = (v6) => {
          const v7 = (v8) => {
            const v9 = (v10) => {
              const v11 = (v12) => {
                const v13 = (v14) => {
                  const v15 = (v16) => {
                    const v17 = (v18) => {
                      const v19 = (v20) => {
                        const v21 = (v22) => {
                          const v23 = (v24) => {
                            const v25 = (v26) => {
                              const v27 = (v28) => {
                                const v29 = (v30) => {
                                  const v31 = (v32) => {
                                    const v33 = (v34) => {
                                      const v35 = (v36) => {
                                        const v37 = (v38) => {
                                          const v39 = (v40) => {
                                            const v41 = (v42) => {
                                              const v43 = (v44) => {
                                                const v45 = (v46) => {
                                                  const v47 = (v48) => {
                                                    const v49 = (v50) => {
                                                      if ($0.tag === "OpStringAppend" && $1.tag === "NeutLit" && $1._1.tag === "LitString" && $2.tag === "NeutLit" && $2._1.tag === "LitString") {
                                                        return $BackendSemantics("NeutLit", $Literal("LitString", $1._1._1 + $2._1._1));
                                                      }
                                                      if ($0.tag === "OpArrayIndex" && $2.tag === "NeutLit" && $2._1.tag === "LitInt") {
                                                        return evalAccessor(env)($1)($BackendAccessor("GetIndex", $2._1._1));
                                                      }
                                                      if ($0.tag === "OpBooleanAnd") {
                                                        if ($1.tag === "NeutFail") {
                                                          return $BackendSemantics("NeutFail", $1._1);
                                                        }
                                                        if ($2.tag === "NeutFail") {
                                                          return $BackendSemantics("NeutFail", $2._1);
                                                        }
                                                        return $BackendSemantics("NeutPrimOp", $BackendOperator("Op2", $0, $1, $2));
                                                      }
                                                      if ($0.tag === "OpBooleanOr") {
                                                        if ($1.tag === "NeutFail") {
                                                          return $BackendSemantics("NeutFail", $1._1);
                                                        }
                                                        if ($2.tag === "NeutFail") {
                                                          return $BackendSemantics("NeutFail", $2._1);
                                                        }
                                                        return $BackendSemantics("NeutPrimOp", $BackendOperator("Op2", $0, $1, $2));
                                                      }
                                                      if ($1.tag === "NeutFail") {
                                                        return $BackendSemantics("NeutFail", $1._1);
                                                      }
                                                      if ($2.tag === "NeutFail") {
                                                        return $BackendSemantics("NeutFail", $2._1);
                                                      }
                                                      return floatLet($1)((x$p) => floatLet($2)((y$p) => {
                                                        if ((() => {
                                                          if ($0.tag === "OpIntNum") {
                                                            return $0._1 === "OpAdd" || $0._1 === "OpMultiply";
                                                          }
                                                          if ($0.tag === "OpNumberNum") {
                                                            return $0._1 === "OpAdd" || $0._1 === "OpMultiply";
                                                          }
                                                          return $0.tag === "OpStringAppend";
                                                        })()) {
                                                          return evalAssocOp(env)($Either("Right", $0))(x$p)(y$p);
                                                        }
                                                        return $BackendSemantics("NeutPrimOp", $BackendOperator("Op2", $0, x$p, y$p));
                                                      }));
                                                    };
                                                    if ($0.tag === "OpStringOrd") {
                                                      const $3 = $1.tag === "SemRef" ? force($1._3) : $1;
                                                      if ($3.tag === "NeutLit" && $3._1.tag === "LitString") {
                                                        const $4 = $2.tag === "SemRef" ? force($2._3) : $2;
                                                        if ($4.tag === "NeutLit" && $4._1.tag === "LitString") {
                                                          return $BackendSemantics(
                                                            "NeutLit",
                                                            $Literal("LitBoolean", evalPrimOpOrd1($0._1)($3._1._1)($4._1._1))
                                                          );
                                                        }
                                                      }
                                                    }
                                                    return v49(true);
                                                  };
                                                  if ($0.tag === "OpNumberOrd") {
                                                    const $3 = $1.tag === "SemRef" ? force($1._3) : $1;
                                                    if ($3.tag === "NeutLit" && $3._1.tag === "LitNumber") {
                                                      const $4 = $2.tag === "SemRef" ? force($2._3) : $2;
                                                      if ($4.tag === "NeutLit" && $4._1.tag === "LitNumber") {
                                                        return $BackendSemantics(
                                                          "NeutLit",
                                                          $Literal(
                                                            "LitBoolean",
                                                            (() => {
                                                              if ($0._1 === "OpEq") {
                                                                return $3._1._1 === $4._1._1;
                                                              }
                                                              if ($0._1 === "OpNotEq") {
                                                                return $3._1._1 !== $4._1._1;
                                                              }
                                                              if ($0._1 === "OpGt") {
                                                                return $3._1._1 > $4._1._1;
                                                              }
                                                              if ($0._1 === "OpGte") {
                                                                return $3._1._1 >= $4._1._1;
                                                              }
                                                              if ($0._1 === "OpLt") {
                                                                return $3._1._1 < $4._1._1;
                                                              }
                                                              if ($0._1 === "OpLte") {
                                                                return $3._1._1 <= $4._1._1;
                                                              }
                                                              fail();
                                                            })()
                                                          )
                                                        );
                                                      }
                                                    }
                                                  }
                                                  return v47(true);
                                                };
                                                if ($0.tag === "OpNumberNum") {
                                                  const $3 = evalPrimOpNumNumber($0._1)($1)($2);
                                                  if ($3.tag === "Just") {
                                                    return $3._1;
                                                  }
                                                }
                                                return v45(true);
                                              };
                                              if ($0.tag === "OpNumberNum" && $0._1 === "OpSubtract") {
                                                const $3 = $1.tag === "SemRef" ? force($1._3) : $1;
                                                if ($3.tag === "NeutLit" && $3._1.tag === "LitNumber" && $3._1._1 === 0) {
                                                  return evalPrimOp(env)($BackendOperator(
                                                    "Op1",
                                                    OpNumberNegate,
                                                    $2
                                                  ));
                                                }
                                              }
                                              return v43(true);
                                            };
                                            if ($0.tag === "OpIntOrd") {
                                              const $3 = $1.tag === "SemRef" ? force($1._3) : $1;
                                              if ($3.tag === "NeutLit" && $3._1.tag === "LitInt") {
                                                const $4 = $2.tag === "SemRef" ? force($2._3) : $2;
                                                if ($4.tag === "NeutLit" && $4._1.tag === "LitInt") {
                                                  return $BackendSemantics(
                                                    "NeutLit",
                                                    $Literal("LitBoolean", evalPrimOpOrd2($0._1)($3._1._1)($4._1._1))
                                                  );
                                                }
                                              }
                                            }
                                            return v41(true);
                                          };
                                          if ($0.tag === "OpIntNum") {
                                            const $3 = evalPrimOpNumInt($0._1)($1)($2);
                                            if ($3.tag === "Just") {
                                              return $3._1;
                                            }
                                          }
                                          return v39(true);
                                        };
                                        if ($0.tag === "OpIntNum" && $0._1 === "OpSubtract") {
                                          const $3 = $1.tag === "SemRef" ? force($1._3) : $1;
                                          if ($3.tag === "NeutLit" && $3._1.tag === "LitInt" && $3._1._1 === 0) {
                                            return evalPrimOp(env)($BackendOperator(
                                              "Op1",
                                              OpIntNegate,
                                              $2
                                            ));
                                          }
                                        }
                                        return v37(true);
                                      };
                                      if ($0.tag === "OpIntBitZeroFillShiftRight") {
                                        const $3 = $1.tag === "SemRef" ? force($1._3) : $1;
                                        if ($3.tag === "NeutLit" && $3._1.tag === "LitInt") {
                                          const $4 = $2.tag === "SemRef" ? force($2._3) : $2;
                                          if ($4.tag === "NeutLit" && $4._1.tag === "LitInt") {
                                            return $BackendSemantics("NeutLit", $Literal("LitInt", $3._1._1 >>> $4._1._1));
                                          }
                                        }
                                      }
                                      return v35(true);
                                    };
                                    if ($0.tag === "OpIntBitXor") {
                                      const $3 = $1.tag === "SemRef" ? force($1._3) : $1;
                                      if ($3.tag === "NeutLit" && $3._1.tag === "LitInt") {
                                        const $4 = $2.tag === "SemRef" ? force($2._3) : $2;
                                        if ($4.tag === "NeutLit" && $4._1.tag === "LitInt") {
                                          return $BackendSemantics("NeutLit", $Literal("LitInt", $3._1._1 ^ $4._1._1));
                                        }
                                      }
                                    }
                                    return v33(true);
                                  };
                                  if ($0.tag === "OpIntBitShiftRight") {
                                    const $3 = $1.tag === "SemRef" ? force($1._3) : $1;
                                    if ($3.tag === "NeutLit" && $3._1.tag === "LitInt") {
                                      const $4 = $2.tag === "SemRef" ? force($2._3) : $2;
                                      if ($4.tag === "NeutLit" && $4._1.tag === "LitInt") {
                                        return $BackendSemantics("NeutLit", $Literal("LitInt", $3._1._1 >> $4._1._1));
                                      }
                                    }
                                  }
                                  return v31(true);
                                };
                                if ($0.tag === "OpIntBitShiftLeft") {
                                  const $3 = $1.tag === "SemRef" ? force($1._3) : $1;
                                  if ($3.tag === "NeutLit" && $3._1.tag === "LitInt") {
                                    const $4 = $2.tag === "SemRef" ? force($2._3) : $2;
                                    if ($4.tag === "NeutLit" && $4._1.tag === "LitInt") {
                                      return $BackendSemantics("NeutLit", $Literal("LitInt", $3._1._1 << $4._1._1));
                                    }
                                  }
                                }
                                return v29(true);
                              };
                              if ($0.tag === "OpIntBitOr") {
                                const $3 = $1.tag === "SemRef" ? force($1._3) : $1;
                                if ($3.tag === "NeutLit" && $3._1.tag === "LitInt") {
                                  const $4 = $2.tag === "SemRef" ? force($2._3) : $2;
                                  if ($4.tag === "NeutLit" && $4._1.tag === "LitInt") {
                                    return $BackendSemantics("NeutLit", $Literal("LitInt", $3._1._1 | $4._1._1));
                                  }
                                }
                              }
                              return v27(true);
                            };
                            if ($0.tag === "OpIntBitAnd") {
                              const $3 = $1.tag === "SemRef" ? force($1._3) : $1;
                              if ($3.tag === "NeutLit" && $3._1.tag === "LitInt") {
                                const $4 = $2.tag === "SemRef" ? force($2._3) : $2;
                                if ($4.tag === "NeutLit" && $4._1.tag === "LitInt") {
                                  return $BackendSemantics("NeutLit", $Literal("LitInt", $3._1._1 & $4._1._1));
                                }
                              }
                            }
                            return v25(true);
                          };
                          if ($0.tag === "OpCharOrd") {
                            const $3 = $1.tag === "SemRef" ? force($1._3) : $1;
                            if ($3.tag === "NeutLit" && $3._1.tag === "LitChar") {
                              const $4 = $2.tag === "SemRef" ? force($2._3) : $2;
                              if ($4.tag === "NeutLit" && $4._1.tag === "LitChar") {
                                return $BackendSemantics("NeutLit", $Literal("LitBoolean", evalPrimOpOrd3($0._1)($3._1._1)($4._1._1)));
                              }
                            }
                          }
                          return v23(true);
                        };
                        if ($0.tag === "OpBooleanOrd") {
                          const $3 = $1.tag === "SemRef" ? force($1._3) : $1;
                          if ($3.tag === "NeutLit" && $3._1.tag === "LitBoolean") {
                            const $4 = $2.tag === "SemRef" ? force($2._3) : $2;
                            if ($4.tag === "NeutLit" && $4._1.tag === "LitBoolean") {
                              return $BackendSemantics("NeutLit", $Literal("LitBoolean", evalPrimOpOrd4($0._1)($3._1._1)($4._1._1)));
                            }
                          }
                        }
                        return v21(true);
                      };
                      if ($0.tag === "OpBooleanOrd" && $0._1 === "OpEq") {
                        const $3 = $2.tag === "SemRef" ? force($2._3) : $2;
                        if ($3.tag === "NeutLit" && $3._1.tag === "LitBoolean") {
                          if ($3._1._1) {
                            return $1;
                          }
                          return evalPrimOp(env)($BackendOperator("Op1", OpBooleanNot, $1));
                        }
                      }
                      return v19(true);
                    };
                    if ($0.tag === "OpBooleanOrd" && $0._1 === "OpEq") {
                      const $3 = $1.tag === "SemRef" ? force($1._3) : $1;
                      if ($3.tag === "NeutLit" && $3._1.tag === "LitBoolean") {
                        if ($3._1._1) {
                          return $2;
                        }
                        return evalPrimOp(env)($BackendOperator("Op1", OpBooleanNot, $2));
                      }
                    }
                    return v17(true);
                  };
                  if ($0.tag === "OpBooleanOr") {
                    const $3 = $2.tag === "SemRef" ? force($2._3) : $2;
                    if ($3.tag === "NeutLit" && $3._1.tag === "LitBoolean" && $3._1._1) {
                      return $2;
                    }
                  }
                  return v15(true);
                };
                if ($0.tag === "OpBooleanOr") {
                  const $3 = $1.tag === "SemRef" ? force($1._3) : $1;
                  if ($3.tag === "NeutLit" && $3._1.tag === "LitBoolean" && $3._1._1) {
                    return $1;
                  }
                }
                return v13(true);
              };
              if ($0.tag === "OpBooleanOr") {
                const $3 = $2.tag === "SemRef" ? force($2._3) : $2;
                if ($3.tag === "NeutLit" && $3._1.tag === "LitBoolean" && !$3._1._1) {
                  return $1;
                }
              }
              return v11(true);
            };
            if ($0.tag === "OpBooleanOr") {
              const $3 = $1.tag === "SemRef" ? force($1._3) : $1;
              if ($3.tag === "NeutLit" && $3._1.tag === "LitBoolean" && !$3._1._1) {
                return $2;
              }
            }
            return v9(true);
          };
          if ($0.tag === "OpBooleanAnd") {
            const $3 = $2.tag === "SemRef" ? force($2._3) : $2;
            if ($3.tag === "NeutLit" && $3._1.tag === "LitBoolean" && $3._1._1) {
              return $1;
            }
          }
          return v7(true);
        };
        if ($0.tag === "OpBooleanAnd") {
          const $3 = $1.tag === "SemRef" ? force($1._3) : $1;
          if ($3.tag === "NeutLit" && $3._1.tag === "LitBoolean" && $3._1._1) {
            return $2;
          }
        }
        return v5(true);
      };
      if ($0.tag === "OpBooleanAnd") {
        const $3 = $2.tag === "SemRef" ? force($2._3) : $2;
        if ($3.tag === "NeutLit" && $3._1.tag === "LitBoolean" && !$3._1._1) {
          return $2;
        }
      }
      return v3(true);
    };
    if ($0.tag === "OpBooleanAnd") {
      const $3 = $1.tag === "SemRef" ? force($1._3) : $1;
      if ($3.tag === "NeutLit" && $3._1.tag === "LitBoolean" && !$3._1._1) {
        return $1;
      }
    }
    return v1(true);
  }
  fail();
};
var evalAssocOp$p = (v) => (op) => (a) => (b) => {
  if (op.tag === "Left") {
    const v1 = v.evalExternSpine(v)(op._1)([$ExternSpine("ExternApp", [a, b])]);
    if (v1.tag === "Just") {
      return v1._1;
    }
    if (v1.tag === "Nothing") {
      return $BackendSemantics("SemAssocOp", op, [a, b]);
    }
    fail();
  }
  if (op.tag === "Right") {
    return evalPrimOp(v)($BackendOperator("Op2", op._1, a, b));
  }
  fail();
};
var evalAssocOp = (env) => (op1) => (v) => (v1) => {
  if (v.tag === "SemAssocOp" && v1.tag === "SemAssocOp" && eq16(op1)(v._1) && eq16(v._1)(v1._1)) {
    const v3 = evalAssocOp$p(env)(op1)((() => {
      const $02 = v._2.length - 1 | 0;
      if ($02 >= 0 && $02 < v._2.length) {
        return v._2[$02];
      }
      fail();
    })())((() => {
      if (0 < v1._2.length) {
        return v1._2[0];
      }
      fail();
    })());
    if (v3.tag === "SemAssocOp" && eq16(v1._1)(v3._1)) {
      return $BackendSemantics(
        "SemAssocOp",
        op1,
        (() => {
          const $02 = unconsImpl((v$1) => Nothing, (v$1) => (xs) => $Maybe("Just", xs), v1._2);
          return [
            ...(() => {
              if (v._2.length === 0) {
                fail();
              }
              return sliceImpl(0, v._2.length - 1 | 0, v._2);
            })(),
            ...v3._2,
            ...(() => {
              if ($02.tag === "Just") {
                return $02._1;
              }
              fail();
            })()
          ];
        })()
      );
    }
    return $BackendSemantics(
      "SemAssocOp",
      op1,
      (() => {
        const $02 = unconsImpl((v$1) => Nothing, (v$1) => (xs) => $Maybe("Just", xs), v1._2);
        return [
          ...(() => {
            if (v._2.length === 0) {
              fail();
            }
            return sliceImpl(0, v._2.length - 1 | 0, v._2);
          })(),
          v3,
          ...(() => {
            if ($02.tag === "Just") {
              return $02._1;
            }
            fail();
          })()
        ];
      })()
    );
  }
  const $0 = (as, b, op2) => {
    const v4 = evalAssocOp$p(env)(op1)((() => {
      const $02 = as.length - 1 | 0;
      if ($02 >= 0 && $02 < as.length) {
        return as[$02];
      }
      fail();
    })())(b);
    if (v4.tag === "SemAssocOp" && eq16(op2)(v4._1)) {
      return $BackendSemantics(
        "SemAssocOp",
        op1,
        [
          ...(() => {
            if (as.length === 0) {
              fail();
            }
            return sliceImpl(0, as.length - 1 | 0, as);
          })(),
          ...v4._2
        ]
      );
    }
    return $BackendSemantics(
      "SemAssocOp",
      op1,
      snoc((() => {
        if (as.length === 0) {
          fail();
        }
        return sliceImpl(0, as.length - 1 | 0, as);
      })())(v4)
    );
  };
  if (v1.tag === "SemAssocOp" && eq16(op1)(v1._1)) {
    const v4 = evalAssocOp$p(env)(op1)(v)((() => {
      if (0 < v1._2.length) {
        return v1._2[0];
      }
      fail();
    })());
    if (v4.tag === "SemAssocOp" && eq16(v1._1)(v4._1)) {
      return $BackendSemantics(
        "SemAssocOp",
        op1,
        (() => {
          const $1 = unconsImpl((v$1) => Nothing, (v$1) => (xs) => $Maybe("Just", xs), v1._2);
          return [
            ...v4._2,
            ...(() => {
              if ($1.tag === "Just") {
                return $1._1;
              }
              fail();
            })()
          ];
        })()
      );
    }
    return $BackendSemantics(
      "SemAssocOp",
      op1,
      (() => {
        const $1 = unconsImpl((v$1) => Nothing, (v$1) => (xs) => $Maybe("Just", xs), v1._2);
        return [
          v4,
          ...(() => {
            if ($1.tag === "Just") {
              return $1._1;
            }
            fail();
          })()
        ];
      })()
    );
  }
  if (v.tag === "SemAssocOp" && eq16(op1)(v._1)) {
    return $0(v._2, v1, v._1);
  }
  return $BackendSemantics("SemAssocOp", op1, [v, v1]);
};
var evalAccessor = (env) => (lhs) => (accessor) => floatLet(lhs)((v) => {
  if (v.tag === "SemRef") {
    return evalRef(env)(v._1)(v._2)($ExternSpine("ExternAccessor", accessor))(v._3);
  }
  const v1 = (v2) => {
    if (v.tag === "NeutUpdate" && accessor.tag === "GetProp") {
      const $0 = accessor._1;
      const v4 = findMapImpl(
        Nothing,
        isJust,
        (v5) => {
          if (v5._1 === $0) {
            return $Maybe("Just", v5._2);
          }
          return Nothing;
        },
        v._2
      );
      if (v4.tag === "Just") {
        return v4._1;
      }
      if (v4.tag === "Nothing") {
        return evalAccessor(env)(v._1)(accessor);
      }
      fail();
    }
    if (v.tag === "NeutLit" && v._1.tag === "LitArray" && accessor.tag === "GetIndex" && accessor._1 >= 0 && accessor._1 < v._1._1.length) {
      return v._1._1[accessor._1];
    }
    if (v.tag === "NeutData" && accessor.tag === "GetCtorField" && accessor._6 >= 0 && accessor._6 < v._5.length) {
      return v._5[accessor._6]._2;
    }
    if (v.tag === "NeutFail") {
      return $BackendSemantics("NeutFail", v._1);
    }
    return $BackendSemantics("NeutAccessor", v, accessor);
  };
  if (v.tag === "NeutLit" && v._1.tag === "LitRecord" && accessor.tag === "GetProp") {
    const $0 = accessor._1;
    const $1 = findMapImpl(
      Nothing,
      isJust,
      (v2) => {
        if (v2._1 === $0) {
          return $Maybe("Just", v2._2);
        }
        return Nothing;
      },
      v._1._1
    );
    if ($1.tag === "Just") {
      return $1._1;
    }
  }
  return v1(true);
});
var build = (ctx) => (v) => {
  const $0 = () => {
    const v1 = (v2) => {
      const v3 = (v4) => {
        const v5 = (v6) => {
          const v7 = (v8) => {
            const v9 = (v10) => {
              const v11 = (v12) => {
                const v13 = (v14) => {
                  const v15 = (v16) => {
                    const v17 = (v18) => {
                      if (v.tag === "Accessor" && v._1.tag === "ExprSyntax" && v._1._2.tag === "Branch") {
                        return $BackendExpr(
                          "ExprRewrite",
                          { ...v._1._1, rewrite: true, size: v._1._1.size + 1 | 0 },
                          $BackendRewrite("RewriteDistBranchesOp", v._1._2._1, v._1._2._2, $DistOp("DistAccessor", v._2))
                        );
                      }
                      if (v.tag === "PrimOp" && v._1.tag === "Op1" && v._1._2.tag === "ExprSyntax" && v._1._2._2.tag === "Branch") {
                        return $BackendExpr(
                          "ExprRewrite",
                          { ...v._1._2._1, rewrite: true, size: v._1._2._1.size + 1 | 0 },
                          $BackendRewrite("RewriteDistBranchesOp", v._1._2._2._1, v._1._2._2._2, $DistOp("DistPrimOp1", v._1._1))
                        );
                      }
                      const v23 = (v24) => {
                        const v25 = (v26) => {
                          if (v.tag === "EffectBind") {
                            if (v._3.tag === "ExprSyntax") {
                              if (v._3._2.tag === "EffectPure") {
                                return build(ctx)($BackendSyntax(
                                  "EffectDefer",
                                  build(ctx)($BackendSyntax("Let", v._1, v._2, v._3._2._1, v._4))
                                ));
                              }
                              if (v._3._2.tag === "EffectDefer") {
                                return build(ctx)($BackendSyntax("EffectBind", v._1, v._2, v._3._2._1, v._4));
                              }
                              if (v._4.tag === "ExprSyntax") {
                                if (v._4._2.tag === "EffectDefer") {
                                  return build(ctx)($BackendSyntax("EffectBind", v._1, v._2, v._3, v._4._2._1));
                                }
                                if (v._4._2.tag === "EffectPure" && v._4._2._1.tag === "ExprSyntax" && v._4._2._1._2.tag === "Local" && v._2 === v._4._2._1._2._2) {
                                  return v._3;
                                }
                              }
                              return $BackendExpr("ExprSyntax", ctx.analyze(ctx)(v), v);
                            }
                            if (v._4.tag === "ExprSyntax") {
                              if (v._4._2.tag === "EffectDefer") {
                                return build(ctx)($BackendSyntax("EffectBind", v._1, v._2, v._3, v._4._2._1));
                              }
                              if (v._4._2.tag === "EffectPure" && v._4._2._1.tag === "ExprSyntax" && v._4._2._1._2.tag === "Local" && v._2 === v._4._2._1._2._2) {
                                return v._3;
                              }
                            }
                            return $BackendExpr("ExprSyntax", ctx.analyze(ctx)(v), v);
                          }
                          if (v.tag === "EffectDefer") {
                            if (v._1.tag === "ExprSyntax" && v._1._2.tag === "EffectDefer") {
                              return v._1;
                            }
                            return $BackendExpr("ExprSyntax", ctx.analyze(ctx)(v), v);
                          }
                          if (v.tag === "PrimOp" && v._1.tag === "Op1" && v._1._1.tag === "OpBooleanNot" && v._1._2.tag === "ExprSyntax" && v._1._2._2.tag === "PrimOp" && v._1._2._2._1.tag === "Op1" && v._1._2._2._1._1.tag === "OpBooleanNot") {
                            return v._1._2._2._1._2;
                          }
                          return $BackendExpr("ExprSyntax", ctx.analyze(ctx)(v), v);
                        };
                        if (v.tag === "PrimOp" && v._1.tag === "Op2" && v._1._3.tag === "ExprSyntax" && v._1._3._2.tag === "Branch") {
                          const $02 = shouldDistributeBranchPrimOp2R(v._1._3._1)(v._1._3._2._1)(v._1._3._2._2)(v._1._2)(v._1._1);
                          if ($02.tag === "Just") {
                            return $02._1;
                          }
                        }
                        return v25(true);
                      };
                      if (v.tag === "PrimOp" && v._1.tag === "Op2" && v._1._2.tag === "ExprSyntax" && v._1._2._2.tag === "Branch") {
                        const $02 = shouldDistributeBranchPrimOp2L(v._1._2._1)(v._1._2._2._1)(v._1._2._2._2)(v._1._1)(v._1._3);
                        if ($02.tag === "Just") {
                          return $02._1;
                        }
                      }
                      return v23(true);
                    };
                    if (v.tag === "UncurriedApp" && v._1.tag === "ExprSyntax" && v._1._2.tag === "Branch") {
                      const $02 = shouldDistributeBranchUncurriedApps(v._1._1)(v._1._2._1)(v._1._2._2)(v._2);
                      if ($02.tag === "Just") {
                        return $02._1;
                      }
                    }
                    return v17(true);
                  };
                  if (v.tag === "App" && v._1.tag === "ExprSyntax" && v._1._2.tag === "Branch") {
                    const $02 = shouldDistributeBranchApps(v._1._1)(v._1._2._1)(v._1._2._2)(v._2);
                    if ($02.tag === "Just") {
                      return $02._1;
                    }
                  }
                  return v15(true);
                };
                if (v.tag === "Let") {
                  const $02 = shouldEtaReduce(v._2)(v._3)(v._4);
                  if ($02.tag === "Just") {
                    return $02._1;
                  }
                }
                return v13(true);
              };
              if (v.tag === "Let") {
                const $02 = shouldDistributeBranches(v._1)(v._2)(v._3)(v._4);
                if ($02.tag === "Just") {
                  return $02._1;
                }
              }
              return v11(true);
            };
            if (v.tag === "Let") {
              const $02 = shouldUnpackArray(v._1)(v._2)(v._3)(v._4);
              if ($02.tag === "Just") {
                return $02._1;
              }
            }
            return v9(true);
          };
          if (v.tag === "Let") {
            const $02 = shouldUnpackCtor(v._1)(v._2)(v._3)(v._4);
            if ($02.tag === "Just") {
              return $02._1;
            }
          }
          return v7(true);
        };
        if (v.tag === "Let") {
          const $02 = shouldUnpackUpdate(v._1)(v._2)(v._3)(v._4);
          if ($02.tag === "Just") {
            return $02._1;
          }
        }
        return v5(true);
      };
      if (v.tag === "Let") {
        const $02 = shouldUnpackRecord(v._1)(v._2)(v._3)(v._4);
        if ($02.tag === "Just") {
          return $02._1;
        }
      }
      return v3(true);
    };
    if (v.tag === "Let") {
      const $02 = shouldUncurryAbs(v._1)(v._2)(v._3)(v._4);
      if ($02.tag === "Just") {
        return $02._1;
      }
    }
    return v1(true);
  };
  if (v.tag === "App") {
    if (v._1.tag === "ExprSyntax" && v._1._2.tag === "App") {
      return build(ctx)($BackendSyntax("App", v._1._2._1, [...v._1._2._2, ...v._2]));
    }
    return $0();
  }
  if (v.tag === "Abs") {
    if (v._2.tag === "ExprSyntax" && v._2._2.tag === "Abs") {
      return build(ctx)($BackendSyntax("Abs", [...v._1, ...v._2._2._1], v._2._2._2));
    }
    return $0();
  }
  if (v.tag === "Let" && shouldInlineLet(v._2)(v._3)(v._4)) {
    return rewriteInline(v._1)(v._2)(v._3)(v._4);
  }
  return $0();
};
var simplifyCondBoolean = (ctx) => (v) => (v1) => {
  if (v._2.tag === "ExprSyntax" && v._2._2.tag === "Lit" && v._2._2._1.tag === "LitBoolean") {
    if (v1.tag === "ExprSyntax" && v1._2.tag === "Lit" && v1._2._1.tag === "LitBoolean") {
      if (v._2._2._1._1 === v1._2._1._1) {
        return $Maybe("Just", v._2);
      }
      if (v._2._2._1._1 && !v1._2._1._1) {
        return $Maybe("Just", v._1);
      }
      if (!v._2._2._1._1 && v1._2._1._1) {
        return $Maybe(
          "Just",
          build(ctx)($BackendSyntax(
            "PrimOp",
            $BackendOperator("Op1", OpBooleanNot, v._1)
          ))
        );
      }
      if (v._2._2._1._1 && v1.tag === "ExprSyntax" && (v1._2.tag === "Lit" || v1._2.tag === "Var" || v1._2.tag === "Local" || v1._2.tag === "PrimOp")) {
        return $Maybe(
          "Just",
          build(ctx)($BackendSyntax(
            "PrimOp",
            $BackendOperator("Op2", OpBooleanOr, v._1, v1)
          ))
        );
      }
      if (!v1._2._1._1) {
        return $Maybe(
          "Just",
          build(ctx)($BackendSyntax(
            "PrimOp",
            $BackendOperator("Op2", OpBooleanAnd, v._1, v._2)
          ))
        );
      }
      return Nothing;
    }
    if (v._2._2._1._1 && v1.tag === "ExprSyntax" && (v1._2.tag === "Lit" || v1._2.tag === "Var" || v1._2.tag === "Local" || v1._2.tag === "PrimOp")) {
      return $Maybe(
        "Just",
        build(ctx)($BackendSyntax(
          "PrimOp",
          $BackendOperator("Op2", OpBooleanOr, v._1, v1)
        ))
      );
    }
    return Nothing;
  }
  if (v1.tag === "ExprSyntax" && v1._2.tag === "Lit" && v1._2._1.tag === "LitBoolean" && !v1._2._1._1) {
    return $Maybe(
      "Just",
      build(ctx)($BackendSyntax(
        "PrimOp",
        $BackendOperator("Op2", OpBooleanAnd, v._1, v._2)
      ))
    );
  }
  return Nothing;
};
var simplifyCondRedundantElse = (ctx) => (v) => (v1) => {
  if (v1.tag === "ExprSyntax" && v1._2.tag === "Branch") {
    const $0 = (() => {
      if (0 < v1._2._1.length) {
        return v1._2._1[0];
      }
      fail();
    })();
    if ($0._1.tag === "ExprSyntax" && $0._1._2.tag === "PrimOp" && $0._1._2._1.tag === "Op1" && $0._1._2._1._1.tag === "OpBooleanNot" && eqBackendExpr.eq(v._1)($0._1._2._1._2)) {
      return $Maybe("Just", buildBranchCond(ctx)($Pair(v._1, v._2))($0._2));
    }
  }
  return Nothing;
};
var simplifyCondLiftAnd = (ctx) => (pair) => (def1) => {
  if (pair._2.tag === "ExprSyntax" && pair._2._2.tag === "Branch" && pair._2._2._1.length === 1 && eqBackendExpr.eq(def1)(pair._2._2._2)) {
    return $Maybe(
      "Just",
      buildBranchCond(ctx)($Pair(
        build(ctx)($BackendSyntax(
          "PrimOp",
          $BackendOperator("Op2", OpBooleanAnd, pair._1, pair._2._2._1[0]._1)
        )),
        pair._2._2._1[0]._2
      ))(def1)
    );
  }
  return Nothing;
};
var buildBranchCond = (ctx) => (pair) => (def) => {
  const $0 = simplifyCondIsTag(ctx)(pair)(def);
  if ($0.tag === "Just") {
    return $0._1;
  }
  const $1 = simplifyCondBoolean(ctx)(pair)(def);
  if ($1.tag === "Just") {
    return $1._1;
  }
  const $2 = simplifyCondLiftAnd(ctx)(pair)(def);
  if ($2.tag === "Just") {
    return $2._1;
  }
  const $3 = simplifyCondRedundantElse(ctx)(pair)(def);
  if ($3.tag === "Just") {
    return $3._1;
  }
  if (def.tag === "ExprSyntax" && def._2.tag === "Branch") {
    return build(ctx)($BackendSyntax("Branch", [pair, ...def._2._1], def._2._2));
  }
  return build(ctx)($BackendSyntax("Branch", [pair], def));
};
var quote$lazy = /* @__PURE__ */ binding(() => {
  const go = (go$a0$copy) => (go$a1$copy) => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
    while (go$c) {
      const ctx = go$a0, v = go$a1;
      if (v.tag === "SemLet") {
        const $0 = v._2;
        go$c = false;
        go$r = build(ctx)($BackendSyntax(
          "Let",
          v._1,
          ctx.currentLevel,
          quote$lazy()(ctx.effect ? { ...ctx, effect: false } : ctx)($0),
          quote$lazy()({ ...ctx, currentLevel: ctx.currentLevel + 1 | 0 })(v._3($BackendSemantics(
            "SemRef",
            $EvalRef("EvalLocal", v._1, ctx.currentLevel),
            [],
            defer((v2) => {
              if ($0.tag === "SemRef") {
                return force($0._3);
              }
              return $0;
            })
          )))
        ));
        continue;
      }
      if (v.tag === "SemLetRec") {
        const $0 = ctx.currentLevel;
        const $1 = { ...ctx, currentLevel: ctx.currentLevel + 1 | 0 };
        const neutBindings = arrayMap((v2) => {
          const $2 = v2._1;
          return $Tuple($2, defer((v3) => $BackendSemantics("NeutLocal", $Maybe("Just", $2), $0)));
        })(v._1);
        go$c = false;
        go$r = build(ctx)($BackendSyntax(
          "LetRec",
          $0,
          arrayMap((m) => $Tuple(m._1, quote$lazy()($1.effect ? { ...$1, effect: false } : $1)(m._2(neutBindings))))(v._1),
          quote$lazy()($1)(v._2(neutBindings))
        ));
        continue;
      }
      if (v.tag === "SemEffectBind") {
        const ctx$p = ctx.effect ? ctx : { ...ctx, effect: true };
        go$c = false;
        go$r = build(ctx)($BackendSyntax(
          "EffectBind",
          v._1,
          ctx$p.currentLevel,
          quote$lazy()(ctx$p)(v._2),
          quote$lazy()({ ...ctx$p, currentLevel: ctx$p.currentLevel + 1 | 0 })(v._3($BackendSemantics("NeutLocal", v._1, ctx$p.currentLevel)))
        ));
        continue;
      }
      if (v.tag === "SemEffectPure") {
        go$c = false;
        go$r = build(ctx)($BackendSyntax("EffectPure", quote$lazy()(ctx.effect ? { ...ctx, effect: false } : ctx)(v._1)));
        continue;
      }
      if (v.tag === "SemEffectDefer") {
        go$c = false;
        go$r = build(ctx)($BackendSyntax("EffectDefer", quote$lazy()(ctx.effect ? ctx : { ...ctx, effect: true })(v._1)));
        continue;
      }
      if (v.tag === "SemBranch") {
        const ctx$p = ctx.effect ? { ...ctx, effect: false } : ctx;
        go$c = false;
        go$r = foldrArray(buildBranchCond(ctx))(quote$lazy()(ctx)(force(v._2)))(arrayMap((v1) => $Pair(
          quote$lazy()(ctx$p)(force(v1._1)),
          quote$lazy()(ctx)(force(v1._2))
        ))(v._1));
        continue;
      }
      if (v.tag === "SemRef") {
        if (v._1.tag === "EvalExtern") {
          go$a0 = ctx;
          go$a1 = neutralSpine($BackendSemantics("NeutVar", v._1._1))(v._2);
          continue;
        }
        if (v._1.tag === "EvalLocal") {
          go$a0 = ctx;
          go$a1 = neutralSpine($BackendSemantics("NeutLocal", v._1._1, v._1._2))(v._2);
          continue;
        }
        fail();
      }
      if (v.tag === "SemLam") {
        const $0 = { ...ctx, currentLevel: ctx.currentLevel + 1 | 0 };
        go$c = false;
        go$r = build(ctx)($BackendSyntax(
          "Abs",
          [$Tuple(v._1, ctx.currentLevel)],
          quote$lazy()($0.effect ? { ...$0, effect: false } : $0)(v._2($BackendSemantics("NeutLocal", v._1, ctx.currentLevel)))
        ));
        continue;
      }
      if (v.tag === "SemMkFn") {
        const loop = (loop$a0$copy) => (loop$a1$copy) => (loop$a2$copy) => {
          let loop$a0 = loop$a0$copy, loop$a1 = loop$a1$copy, loop$a2 = loop$a2$copy, loop$c = true, loop$r;
          while (loop$c) {
            const ctx$p = loop$a0, idents = loop$a1, v1 = loop$a2;
            if (v1.tag === "MkFnNext") {
              loop$a0 = { ...ctx$p, currentLevel: ctx$p.currentLevel + 1 | 0 };
              loop$a1 = snoc(idents)($Tuple(v1._1, ctx$p.currentLevel));
              loop$a2 = v1._2($BackendSemantics("NeutLocal", v1._1, ctx$p.currentLevel));
              continue;
            }
            if (v1.tag === "MkFnApplied") {
              loop$c = false;
              loop$r = build(ctx$p)($BackendSyntax(
                "UncurriedAbs",
                idents,
                quote$lazy()(ctx$p.effect ? { ...ctx$p, effect: false } : ctx$p)(v1._1)
              ));
              continue;
            }
            fail();
          }
          return loop$r;
        };
        go$c = false;
        go$r = loop(ctx)([])(v._1);
        continue;
      }
      if (v.tag === "SemMkEffectFn") {
        const loop = (loop$a0$copy) => (loop$a1$copy) => (loop$a2$copy) => {
          let loop$a0 = loop$a0$copy, loop$a1 = loop$a1$copy, loop$a2 = loop$a2$copy, loop$c = true, loop$r;
          while (loop$c) {
            const ctx$p = loop$a0, idents = loop$a1, v1 = loop$a2;
            if (v1.tag === "MkFnNext") {
              loop$a0 = { ...ctx$p, currentLevel: ctx$p.currentLevel + 1 | 0 };
              loop$a1 = snoc(idents)($Tuple(v1._1, ctx$p.currentLevel));
              loop$a2 = v1._2($BackendSemantics("NeutLocal", v1._1, ctx$p.currentLevel));
              continue;
            }
            if (v1.tag === "MkFnApplied") {
              loop$c = false;
              loop$r = build(ctx$p)($BackendSyntax(
                "UncurriedEffectAbs",
                idents,
                quote$lazy()(ctx$p.effect ? { ...ctx$p, effect: false } : ctx$p)(v1._1)
              ));
              continue;
            }
            fail();
          }
          return loop$r;
        };
        go$c = false;
        go$r = loop(ctx)([])(v._1);
        continue;
      }
      if (v.tag === "SemAssocOp") {
        const $0 = v._1;
        const $1 = v._2;
        const len = $1.length;
        const go$1 = (go$1$a0$copy) => (go$1$a1$copy) => {
          let go$1$a0 = go$1$a0$copy, go$1$a1 = go$1$a1$copy, go$1$c = true, go$1$r;
          while (go$1$c) {
            const ix = go$1$a0, acc = go$1$a1;
            if (ix === len) {
              go$1$c = false;
              go$1$r = acc;
              continue;
            }
            go$1$a0 = ix + 1 | 0;
            go$1$a1 = (() => {
              const $2 = $1[ix];
              if ($0.tag === "Left") {
                return build(ctx)($BackendSyntax(
                  "App",
                  build(ctx)($BackendSyntax("Var", $0._1)),
                  [acc, quote$lazy()(ctx)($2)]
                ));
              }
              if ($0.tag === "Right") {
                return build(ctx)($BackendSyntax(
                  "PrimOp",
                  $BackendOperator("Op2", $0._1, acc, quote$lazy()(ctx)($2))
                ));
              }
              fail();
            })();
          }
          return go$1$r;
        };
        go$c = false;
        go$r = go$1(1)(quote$lazy()(ctx)((() => {
          if (0 < $1.length) {
            return $1[0];
          }
          fail();
        })()));
        continue;
      }
      if (v.tag === "NeutLocal") {
        go$c = false;
        go$r = build(ctx)($BackendSyntax("Local", v._1, v._2));
        continue;
      }
      if (v.tag === "NeutVar") {
        go$c = false;
        go$r = build(ctx)($BackendSyntax("Var", v._1));
        continue;
      }
      if (v.tag === "NeutStop") {
        go$c = false;
        go$r = $BackendExpr("ExprRewrite", ctx.analyze(ctx)($BackendSyntax("Var", v._1)), $BackendRewrite("RewriteStop", v._1));
        continue;
      }
      if (v.tag === "NeutData") {
        go$c = false;
        go$r = build(ctx)($BackendSyntax(
          "CtorSaturated",
          v._1,
          v._2,
          v._3,
          v._4,
          arrayMap((() => {
            const $0 = quote$lazy()(ctx);
            return (m) => $Tuple(m._1, $0(m._2));
          })())(v._5)
        ));
        continue;
      }
      if (v.tag === "NeutCtorDef") {
        go$c = false;
        go$r = build(ctx)($BackendSyntax("CtorDef", v._2, v._3, v._4, v._5));
        continue;
      }
      if (v.tag === "NeutUncurriedApp") {
        const ctx$p = ctx.effect ? { ...ctx, effect: false } : ctx;
        go$c = false;
        go$r = build(ctx)($BackendSyntax("UncurriedApp", quote$lazy()(ctx$p)(v._1), arrayMap(quote$lazy()(ctx$p))(v._2)));
        continue;
      }
      if (v.tag === "NeutUncurriedEffectApp") {
        const ctx$p = ctx.effect ? { ...ctx, effect: false } : ctx;
        go$c = false;
        go$r = build(ctx)($BackendSyntax("UncurriedEffectApp", quote$lazy()(ctx$p)(v._1), arrayMap(quote$lazy()(ctx$p))(v._2)));
        continue;
      }
      if (v.tag === "NeutApp") {
        const ctx$p = ctx.effect ? { ...ctx, effect: false } : ctx;
        const hd$p = quote$lazy()(ctx$p)(v._1);
        const $0 = arrayMap(quote$lazy()(ctx$p))(v._2);
        if ($0.length > 0) {
          go$c = false;
          go$r = build(ctx)($BackendSyntax("App", hd$p, $0));
          continue;
        }
        go$c = false;
        go$r = hd$p;
        continue;
      }
      if (v.tag === "NeutAccessor") {
        go$c = false;
        go$r = build(ctx)($BackendSyntax("Accessor", quote$lazy()(ctx)(v._1), v._2));
        continue;
      }
      if (v.tag === "NeutUpdate") {
        go$c = false;
        go$r = build(ctx)($BackendSyntax(
          "Update",
          quote$lazy()(ctx)(v._1),
          arrayMap((() => {
            const $0 = quote$lazy()(ctx);
            return (m) => $Prop(m._1, $0(m._2));
          })())(v._2)
        ));
        continue;
      }
      if (v.tag === "NeutLit") {
        go$c = false;
        go$r = build(ctx)($BackendSyntax("Lit", functorLiteral.map(quote$lazy()(ctx))(v._1)));
        continue;
      }
      if (v.tag === "NeutPrimOp") {
        go$c = false;
        go$r = build(ctx)($BackendSyntax(
          "PrimOp",
          (() => {
            const $0 = quote$lazy()(ctx);
            if (v._1.tag === "Op1") {
              return $BackendOperator("Op1", v._1._1, $0(v._1._2));
            }
            if (v._1.tag === "Op2") {
              return $BackendOperator("Op2", v._1._1, $0(v._1._2), $0(v._1._3));
            }
            fail();
          })()
        ));
        continue;
      }
      if (v.tag === "NeutPrimEffect") {
        go$c = false;
        go$r = build(ctx)($BackendSyntax(
          "PrimEffect",
          (() => {
            const $0 = quote$lazy()(ctx.effect ? { ...ctx, effect: false } : ctx);
            if (v._1.tag === "EffectRefNew") {
              return $BackendEffect("EffectRefNew", $0(v._1._1));
            }
            if (v._1.tag === "EffectRefRead") {
              return $BackendEffect("EffectRefRead", $0(v._1._1));
            }
            if (v._1.tag === "EffectRefWrite") {
              return $BackendEffect("EffectRefWrite", $0(v._1._1), $0(v._1._2));
            }
            fail();
          })()
        ));
        continue;
      }
      if (v.tag === "NeutPrimUndefined") {
        go$c = false;
        go$r = build(ctx)(PrimUndefined);
        continue;
      }
      if (v.tag === "NeutFail") {
        go$c = false;
        go$r = build(ctx)($BackendSyntax("Fail", v._1));
        continue;
      }
      fail();
    }
    return go$r;
  };
  return go;
});
var quote = /* @__PURE__ */ quote$lazy();
var evalApp = (env) => (hd) => (spine) => {
  const go = (env$p) => (v) => (v1) => {
    const $0 = (args, ident, k, val) => $BackendSemantics(
      "SemLet",
      ident,
      val,
      (nextVal) => makeLet(Nothing)(k(nextVal))((nextFn) => go({
        ...env$p,
        locals: snoc(snoc(env$p.locals)($LocalBinding("One", nextVal)))($LocalBinding("One", nextFn))
      })(nextFn)(args))
    );
    const $1 = (args, k, vals) => $BackendSemantics(
      "SemLetRec",
      vals,
      (nextVals) => makeLet(Nothing)(k(nextVals))((nextFn) => go({
        ...env$p,
        locals: snoc(snoc(env$p.locals)($LocalBinding("Group", nextVals)))($LocalBinding("One", nextFn))
      })(nextFn)(args))
    );
    if (v1.tag === "Cons") {
      if (v1._1.tag === "NeutFail") {
        return $BackendSemantics("NeutFail", v1._1._1);
      }
      if (v.tag === "NeutFail") {
        return $BackendSemantics("NeutFail", v._1);
      }
      if (v.tag === "SemLam") {
        const $2 = v1._2;
        return makeLet(Nothing)(v1._1)((nextArg) => go(env$p)(v._2(nextArg))($2));
      }
      if (v.tag === "SemRef") {
        return go(env$p)(evalRef(env$p)(v._1)(v._2)($ExternSpine("ExternApp", [v1._1]))(v._3))(v1._2);
      }
      if (v.tag === "SemLet") {
        return $0(v1, v._1, v._3, v._2);
      }
      if (v.tag === "SemLetRec") {
        return $1(v1, v._2, v._1);
      }
      return $BackendSemantics("NeutApp", v, toUnfoldable1(v1));
    }
    if (v.tag === "NeutFail") {
      return $BackendSemantics("NeutFail", v._1);
    }
    if (v.tag === "SemLet") {
      return $0(v1, v._1, v._3, v._2);
    }
    if (v.tag === "SemLetRec") {
      return $1(v1, v._2, v._1);
    }
    if (v1.tag === "Nil") {
      return v;
    }
    return $BackendSemantics("NeutApp", v, toUnfoldable1(v1));
  };
  return go(env)(hd)(fromFoldable2(spine));
};
var evalUncurriedApp = (env) => (hd) => (spine) => {
  if (hd.tag === "SemMkFn") {
    return evalUncurriedBeta(NeutUncurriedApp)(hd._1)(spine);
  }
  if (hd.tag === "SemRef") {
    const $0 = hd._1;
    const $1 = hd._3;
    const $2 = hd._2;
    return guardFailOver1(identity9)(spine)((spine$p) => evalRef(env)($0)($2)($ExternSpine("ExternUncurriedApp", spine$p))($1));
  }
  if (hd.tag === "SemLet") {
    return $BackendSemantics(
      "SemLet",
      hd._1,
      hd._2,
      (nextVal) => makeLet(Nothing)(hd._3(nextVal))((nextFn) => evalUncurriedApp({
        ...env,
        locals: snoc(snoc(env.locals)($LocalBinding("One", nextVal)))($LocalBinding("One", nextFn))
      })(nextFn)(spine))
    );
  }
  if (hd.tag === "NeutFail") {
    return $BackendSemantics("NeutFail", hd._1);
  }
  return guardFailOver1(identity9)(spine)(NeutUncurriedApp(hd));
};
var evalSpine = (env) => foldlArray((hd) => (v) => {
  if (v.tag === "ExternApp") {
    return evalApp(env)(hd)(v._1);
  }
  if (v.tag === "ExternUncurriedApp") {
    return evalUncurriedApp(env)(hd)(v._1);
  }
  if (v.tag === "ExternAccessor") {
    return evalAccessor(env)(hd)(v._1);
  }
  if (v.tag === "ExternPrimOp") {
    return evalPrimOp(env)($BackendOperator("Op1", v._1, hd));
  }
  fail();
});
var mkUncurriedAppRewrite = (env) => (hd) => {
  const go = (acc) => (n) => {
    if (n === 0) {
      return evalUncurriedApp(env)(hd)(acc);
    }
    return $BackendSemantics("SemLam", Nothing, (arg) => go(snoc(acc)(arg))(n - 1 | 0));
  };
  return go([]);
};
var evalUncurriedEffectApp = (env) => (hd) => (spine) => {
  if (hd.tag === "SemMkEffectFn") {
    return evalUncurriedBeta(NeutUncurriedEffectApp)(hd._1)(spine);
  }
  if (hd.tag === "SemLet") {
    return $BackendSemantics(
      "SemLet",
      hd._1,
      hd._2,
      (nextVal) => makeLet(Nothing)(hd._3(nextVal))((nextFn) => evalUncurriedEffectApp({
        ...env,
        locals: snoc(snoc(env.locals)($LocalBinding("One", nextVal)))($LocalBinding("One", nextFn))
      })(nextFn)(spine))
    );
  }
  if (hd.tag === "NeutFail") {
    return $BackendSemantics("NeutFail", hd._1);
  }
  return guardFailOver1(identity9)(spine)(NeutUncurriedEffectApp(hd));
};
var mkFnFromArgs = (dictEval) => (env) => (args) => (body) => $BackendSemantics(
  "SemMkFn",
  foldrArray((v) => {
    const $0 = v._1;
    return (next) => (env$p) => $MkFn("MkFnNext", $0, (x) => next({ ...env$p, locals: snoc(env$p.locals)($LocalBinding("One", x)) }));
  })((x) => $MkFn("MkFnApplied", dictEval.eval(x)(body)))(args)(env)
);
var evalBackendSyntax = (dictEval) => ({
  eval: (v) => (v1) => {
    if (v1.tag === "Var") {
      const v2 = v.evalExternSpine(v)(v1._1)([]);
      if (v2.tag === "Just") {
        return v2._1;
      }
      if (v2.tag === "Nothing") {
        return $BackendSemantics(
          "SemRef",
          $EvalRef("EvalExtern", v1._1),
          [],
          defer((v3) => {
            const v4 = v.evalExternRef(v)(v1._1);
            if (v4.tag === "Just") {
              if (v4._1.tag === "SemRef") {
                return force(v4._1._3);
              }
              return v4._1;
            }
            if (v4.tag === "Nothing") {
              return $BackendSemantics("NeutVar", v1._1);
            }
            fail();
          })
        );
      }
      fail();
    }
    if (v1.tag === "Local") {
      if (v1._2 >= 0 && v1._2 < v.locals.length) {
        if (v.locals[v1._2].tag === "One") {
          return v.locals[v1._2]._1;
        }
        if (v.locals[v1._2].tag === "Group") {
          const $0 = v.locals[v1._2]._1;
          const $1 = (() => {
            if (v1._1.tag === "Just") {
              return lookup1(v1._1._1)($0);
            }
            if (v1._1.tag === "Nothing") {
              return Nothing;
            }
            fail();
          })();
          if ($1.tag === "Just") {
            return force($1._1);
          }
        }
      }
      return _crashWith("Unbound local at level " + showIntImpl(v1._2));
    }
    if (v1.tag === "App") {
      return evalApp(v)(dictEval.eval(v)(v1._1))(arrayMap(dictEval.eval(v))(v1._2));
    }
    if (v1.tag === "UncurriedApp") {
      return evalUncurriedApp(v)(dictEval.eval(v)(v1._1))(arrayMap(dictEval.eval(v))(v1._2));
    }
    if (v1.tag === "UncurriedAbs") {
      const $0 = v1._2;
      const loop = (env$p) => (v2) => {
        if (v2.tag === "Nil") {
          return $MkFn("MkFnApplied", dictEval.eval(env$p)($0));
        }
        if (v2.tag === "Cons") {
          const $1 = v2._2;
          return $MkFn("MkFnNext", v2._1, (nextArg) => loop({ ...env$p, locals: snoc(env$p.locals)($LocalBinding("One", nextArg)) })($1));
        }
        fail();
      };
      return $BackendSemantics(
        "SemMkFn",
        loop(v)((() => {
          const $1 = arrayMap(fst)(v1._1);
          const len = $1.length;
          const go = (go$a0$copy) => (go$a1$copy) => {
            let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
            while (go$c) {
              const source2 = go$a0, memo = go$a1;
              if (source2 < len) {
                go$a0 = source2 + 1 | 0;
                go$a1 = $List("Cons", $1[source2], memo);
                continue;
              }
              const go$1 = (go$1$a0$copy) => (go$1$a1$copy) => {
                let go$1$a0 = go$1$a0$copy, go$1$a1 = go$1$a1$copy, go$1$c = true, go$1$r;
                while (go$1$c) {
                  const b = go$1$a0, v$1 = go$1$a1;
                  if (v$1.tag === "Nil") {
                    go$1$c = false;
                    go$1$r = b;
                    continue;
                  }
                  if (v$1.tag === "Cons") {
                    go$1$a0 = $List("Cons", v$1._1, b);
                    go$1$a1 = v$1._2;
                    continue;
                  }
                  fail();
                }
                return go$1$r;
              };
              go$c = false;
              go$r = go$1(Nil)(memo);
            }
            return go$r;
          };
          return go(0)(Nil);
        })())
      );
    }
    if (v1.tag === "UncurriedEffectApp") {
      return evalUncurriedEffectApp(v)(dictEval.eval(v)(v1._1))(arrayMap(dictEval.eval(v))(v1._2));
    }
    if (v1.tag === "UncurriedEffectAbs") {
      const $0 = v1._2;
      const loop = (env$p) => (v2) => {
        if (v2.tag === "Nil") {
          return $MkFn("MkFnApplied", dictEval.eval(env$p)($0));
        }
        if (v2.tag === "Cons") {
          const $1 = v2._2;
          return $MkFn("MkFnNext", v2._1, (nextArg) => loop({ ...env$p, locals: snoc(env$p.locals)($LocalBinding("One", nextArg)) })($1));
        }
        fail();
      };
      return $BackendSemantics(
        "SemMkEffectFn",
        loop(v)((() => {
          const $1 = arrayMap(fst)(v1._1);
          const len = $1.length;
          const go = (go$a0$copy) => (go$a1$copy) => {
            let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
            while (go$c) {
              const source2 = go$a0, memo = go$a1;
              if (source2 < len) {
                go$a0 = source2 + 1 | 0;
                go$a1 = $List("Cons", $1[source2], memo);
                continue;
              }
              const go$1 = (go$1$a0$copy) => (go$1$a1$copy) => {
                let go$1$a0 = go$1$a0$copy, go$1$a1 = go$1$a1$copy, go$1$c = true, go$1$r;
                while (go$1$c) {
                  const b = go$1$a0, v$1 = go$1$a1;
                  if (v$1.tag === "Nil") {
                    go$1$c = false;
                    go$1$r = b;
                    continue;
                  }
                  if (v$1.tag === "Cons") {
                    go$1$a0 = $List("Cons", v$1._1, b);
                    go$1$a1 = v$1._2;
                    continue;
                  }
                  fail();
                }
                return go$1$r;
              };
              go$c = false;
              go$r = go$1(Nil)(memo);
            }
            return go$r;
          };
          return go(0)(Nil);
        })())
      );
    }
    if (v1.tag === "Abs") {
      const $0 = v1._2;
      const $1 = v1._1;
      const go = (go$a0$copy) => (go$a1$copy) => {
        let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
        while (go$c) {
          const ix = go$a0, acc = go$a1;
          if (ix < 0) {
            go$c = false;
            go$r = acc;
            continue;
          }
          go$a0 = ix - 1 | 0;
          go$a1 = (() => {
            const $2 = $1[ix]._1;
            return (env$p) => $BackendSemantics("SemLam", $2, (x) => acc({ ...env$p, locals: snoc(env$p.locals)($LocalBinding("One", x)) }));
          })();
        }
        return go$r;
      };
      return go($1.length - 2 | 0)((() => {
        const $2 = $1.length - 1 | 0;
        const $3 = (() => {
          if ($2 >= 0 && $2 < $1.length) {
            return $1[$2]._1;
          }
          fail();
        })();
        return (env$p) => $BackendSemantics("SemLam", $3, (x) => dictEval.eval({ ...env$p, locals: snoc(env$p.locals)($LocalBinding("One", x)) })($0));
      })())(v);
    }
    if (v1.tag === "Let") {
      const $0 = v1._4;
      return makeLet(v1._1)(dictEval.eval(v)(v1._3))((x) => dictEval.eval({ ...v, locals: snoc(v.locals)($LocalBinding("One", x)) })($0));
    }
    if (v1.tag === "LetRec") {
      const bindGroup = (sem) => (x) => dictEval.eval({ ...v, locals: snoc(v.locals)($LocalBinding("Group", x)) })(sem);
      return $BackendSemantics("SemLetRec", arrayMap((m) => $Tuple(m._1, bindGroup(m._2)))(v1._2), bindGroup(v1._3));
    }
    if (v1.tag === "EffectBind") {
      const $0 = v1._4;
      return makeEffectBind(v1._1)(dictEval.eval(v)(v1._3))((x) => dictEval.eval({ ...v, locals: snoc(v.locals)($LocalBinding("One", x)) })($0));
    }
    if (v1.tag === "EffectPure") {
      const $0 = dictEval.eval(v)(v1._1);
      if ($0.tag === "NeutFail") {
        return $BackendSemantics("NeutFail", $0._1);
      }
      return $BackendSemantics("SemEffectPure", $0);
    }
    if (v1.tag === "EffectDefer") {
      const $0 = dictEval.eval(v)(v1._1);
      if ($0.tag === "NeutFail") {
        return $BackendSemantics("NeutFail", $0._1);
      }
      return $BackendSemantics("SemEffectDefer", $0);
    }
    if (v1.tag === "Accessor") {
      return evalAccessor(v)(dictEval.eval(v)(v1._1))(v1._2);
    }
    if (v1.tag === "Update") {
      return evalUpdate(dictEval.eval(v)(v1._1))(arrayMap((() => {
        const $0 = dictEval.eval(v);
        return (m) => $Prop(m._1, $0(m._2));
      })())(v1._2));
    }
    if (v1.tag === "Branch") {
      const $0 = v1._2;
      return evalBranches(v)(arrayMap(evalPair(dictEval)(v))(v1._1))(defer((v2) => dictEval.eval(v)($0)));
    }
    if (v1.tag === "PrimOp") {
      return evalPrimOp(v)((() => {
        const $0 = dictEval.eval(v);
        if (v1._1.tag === "Op1") {
          return $BackendOperator("Op1", v1._1._1, $0(v1._1._2));
        }
        if (v1._1.tag === "Op2") {
          return $BackendOperator("Op2", v1._1._1, $0(v1._1._2), $0(v1._1._3));
        }
        fail();
      })());
    }
    if (v1.tag === "PrimEffect") {
      return guardFailOver2(identity9)((() => {
        const $0 = dictEval.eval(v);
        if (v1._1.tag === "EffectRefNew") {
          return $BackendEffect("EffectRefNew", $0(v1._1._1));
        }
        if (v1._1.tag === "EffectRefRead") {
          return $BackendEffect("EffectRefRead", $0(v1._1._1));
        }
        if (v1._1.tag === "EffectRefWrite") {
          return $BackendEffect("EffectRefWrite", $0(v1._1._1), $0(v1._1._2));
        }
        fail();
      })())(NeutPrimEffect);
    }
    if (v1.tag === "PrimUndefined") {
      return NeutPrimUndefined;
    }
    if (v1.tag === "Lit") {
      return guardFailOver3(identity9)(functorLiteral.map(dictEval.eval(v))(v1._1))(NeutLit);
    }
    if (v1.tag === "Fail") {
      return $BackendSemantics("NeutFail", v1._1);
    }
    if (v1.tag === "CtorDef") {
      return $BackendSemantics("NeutCtorDef", $Qualified($Maybe("Just", v.currentModule), v1._3), v1._1, v1._2, v1._3, v1._4);
    }
    if (v1.tag === "CtorSaturated") {
      return guardFailOver1(snd)(arrayMap((() => {
        const $0 = dictEval.eval(v);
        return (m) => $Tuple(m._1, $0(m._2));
      })())(v1._5))(NeutData(v1._1)(v1._2)(v1._3)(v1._4));
    }
    fail();
  }
});
var evalBackendExpr$lazy = /* @__PURE__ */ binding(() => ({
  eval: /* @__PURE__ */ (() => {
    const go = (go$a0$copy) => (go$a1$copy) => {
      let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
      while (go$c) {
        const env = go$a0, v = go$a1;
        if (v.tag === "ExprRewrite") {
          if (v._2.tag === "RewriteInline") {
            go$a0 = { ...env, locals: snoc(env.locals)($LocalBinding("One", evalBackendExpr$lazy().eval(env)(v._2._3))) };
            go$a1 = v._2._4;
            continue;
          }
          if (v._2.tag === "RewriteUncurry") {
            const $0 = v._2._3;
            const $1 = v._2._5;
            go$c = false;
            go$r = $BackendSemantics(
              "SemLet",
              v._2._1,
              mkFnFromArgs(evalBackendExpr$lazy())(env)($0)(v._2._4),
              (newFn) => evalBackendExpr$lazy().eval({ ...env, locals: snoc(env.locals)($LocalBinding("One", mkUncurriedAppRewrite(env)(newFn)($0.length))) })($1)
            );
            continue;
          }
          if (v._2.tag === "RewriteStop") {
            go$c = false;
            go$r = $BackendSemantics("NeutStop", v._2._1);
            continue;
          }
          if (v._2.tag === "RewriteUnpackOp") {
            if (v._2._3.tag === "UnpackRecord") {
              go$c = false;
              go$r = foldrArray((v1) => {
                const $0 = v1._2;
                const $1 = v1._1;
                return (next) => (props$p) => makeLet(Nothing)(evalBackendExpr$lazy().eval(env)($0))((val) => next(snoc(props$p)($Prop(
                  $1,
                  val
                ))));
              })((() => {
                const $0 = v._2._4;
                return (x) => evalBackendExpr$lazy().eval({
                  ...env,
                  locals: snoc(env.locals)($LocalBinding("One", $BackendSemantics("NeutLit", $Literal("LitRecord", x))))
                })($0);
              })())(v._2._3._1)([]);
              continue;
            }
            if (v._2._3.tag === "UnpackUpdate") {
              const $0 = v._2._3._2;
              go$c = false;
              go$r = makeLet(Nothing)(evalBackendExpr$lazy().eval(env)(v._2._3._1))((hd$p) => foldrArray((v1) => {
                const $1 = v1._2;
                const $2 = v1._1;
                return (next) => (props$p) => makeLet(Nothing)(evalBackendExpr$lazy().eval(env)($1))((val) => next(snoc(props$p)($Prop(
                  $2,
                  val
                ))));
              })((() => {
                const $1 = v._2._4;
                const $2 = NeutUpdate(hd$p);
                return (x) => evalBackendExpr$lazy().eval({ ...env, locals: snoc(env.locals)($LocalBinding("One", $2(x))) })($1);
              })())($0)([]));
              continue;
            }
            if (v._2._3.tag === "UnpackArray") {
              go$c = false;
              go$r = foldrArray((expr) => (next) => (exprs$p) => makeLet(Nothing)(evalBackendExpr$lazy().eval(env)(expr))((val) => next(snoc(exprs$p)(val))))((() => {
                const $0 = v._2._4;
                return (x) => evalBackendExpr$lazy().eval({
                  ...env,
                  locals: snoc(env.locals)($LocalBinding("One", $BackendSemantics("NeutLit", $Literal("LitArray", x))))
                })($0);
              })())(v._2._3._1)([]);
              continue;
            }
            if (v._2._3.tag === "UnpackData") {
              go$c = false;
              go$r = foldrArray((v1) => {
                const $0 = v1._2;
                const $1 = v1._1;
                return (next) => (props$p) => makeLet(Nothing)(evalBackendExpr$lazy().eval(env)($0))((val) => next(snoc(props$p)($Tuple($1, val))));
              })((() => {
                const $0 = v._2._4;
                const $1 = NeutData(v._2._3._1)(v._2._3._2)(v._2._3._3)(v._2._3._4);
                return (x) => evalBackendExpr$lazy().eval({ ...env, locals: snoc(env.locals)($LocalBinding("One", $1(x))) })($0);
              })())(v._2._3._5)([]);
              continue;
            }
            fail();
          }
          if (v._2.tag === "RewriteDistBranchesLet") {
            const $0 = v._2._4;
            go$c = false;
            go$r = rewriteBranches((() => {
              const $1 = v._2._5;
              return (x) => evalBackendExpr$lazy().eval({ ...env, locals: snoc(env.locals)($LocalBinding("One", x)) })($1);
            })())(evalBranches(env)(arrayMap(evalPair(evalBackendExpr$lazy())(env))(v._2._3))(defer((v1) => evalBackendExpr$lazy().eval(env)($0))));
            continue;
          }
          if (v._2.tag === "RewriteDistBranchesOp") {
            const $0 = v._2._2;
            go$c = false;
            go$r = rewriteBranches((() => {
              if (v._2._3.tag === "DistApp") {
                const $1 = arrayMap(evalBackendExpr$lazy().eval(env))(v._2._3._1);
                return (a) => evalApp(env)(a)($1);
              }
              if (v._2._3.tag === "DistUncurriedApp") {
                const $1 = arrayMap(evalBackendExpr$lazy().eval(env))(v._2._3._1);
                return (a) => evalUncurriedApp(env)(a)($1);
              }
              if (v._2._3.tag === "DistAccessor") {
                const $1 = v._2._3._1;
                return (a) => evalAccessor(env)(a)($1);
              }
              if (v._2._3.tag === "DistPrimOp1") {
                const $1 = Op1(v._2._3._1);
                return (x) => evalPrimOp(env)($1(x));
              }
              if (v._2._3.tag === "DistPrimOp2L") {
                const $1 = Op2(v._2._3._1);
                const $2 = evalBackendExpr$lazy().eval(env)(v._2._3._2);
                return (x) => evalPrimOp(env)($1(x)($2));
              }
              if (v._2._3.tag === "DistPrimOp2R") {
                const $1 = Op2(v._2._3._2)(evalBackendExpr$lazy().eval(env)(v._2._3._1));
                return (x) => evalPrimOp(env)($1(x));
              }
              fail();
            })())(evalBranches(env)(arrayMap(evalPair(evalBackendExpr$lazy())(env))(v._2._1))(defer((v1) => evalBackendExpr$lazy().eval(env)($0))));
            continue;
          }
          fail();
        }
        if (v.tag === "ExprSyntax") {
          go$c = false;
          go$r = evalBackendSyntax(evalBackendExpr$lazy()).eval(env)(v._2);
          continue;
        }
        fail();
      }
      return go$r;
    };
    return go;
  })()
}));
var evalBackendExpr = /* @__PURE__ */ evalBackendExpr$lazy();
var optimize = (traceSteps) => (ctx) => (env) => (v) => (initN) => (originalExpr) => {
  const $0 = v._2;
  const $1 = v._1;
  const go = (go$a0$copy) => (go$a1$copy) => (go$a2$copy) => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$a2 = go$a2$copy, go$c = true, go$r;
    while (go$c) {
      const steps = go$a0, n = go$a1, expr1 = go$a2;
      const v1 = (() => {
        if (n === 0) {
          return _crashWith((() => {
            if ($1.tag === "Nothing") {
              return "" + $0 + ": Possible infinite optimization loop.";
            }
            if ($1.tag === "Just") {
              return $1._1 + "." + $0 + ": Possible infinite optimization loop.";
            }
            fail();
          })());
        }
        const expr2 = quote(ctx)(evalBackendExpr.eval(env)(expr1));
        return $Tuple(
          (() => {
            if (expr2.tag === "ExprSyntax") {
              return expr2._1.rewrite;
            }
            if (expr2.tag === "ExprRewrite") {
              return expr2._1.rewrite;
            }
            fail();
          })(),
          expr2
        );
      })();
      const newSteps = traceSteps ? $List("Cons", v1._2, steps) : steps;
      if (v1._1) {
        go$a0 = newSteps;
        go$a1 = n - 1 | 0;
        go$a2 = v1._2;
        continue;
      }
      go$c = false;
      go$r = $Tuple(reverse(toUnfoldable1(newSteps)), v1._2);
    }
    return go$r;
  };
  return go(traceSteps ? $List("Cons", originalExpr, Nil) : Nil)(initN)(originalExpr);
};
var evalNeutralExpr = { eval: (env) => (v) => evalBackendSyntax(evalNeutralExpr).eval(env)(v) };
var eval3 = /* @__PURE__ */ (() => evalBackendSyntax(evalNeutralExpr).eval)();
var addStop = (v) => (ref) => (acc) => ({
  ...v,
  directives: alter2((v2) => {
    if (v2.tag === "Just") {
      return $Maybe("Just", insert(ordInlineAccessor)(acc)(InlineNever)(v2._1));
    }
    return $Maybe("Just", $$$Map("Node", 1, 1, acc, InlineNever, Leaf, Leaf));
  })(ref)(v.directives)
});
var evalExternFromImpl = (v) => (qual) => (v1) => (spine) => {
  if (spine.length === 0) {
    if (v1._2.tag === "ExternExpr") {
      const $0 = v1._2._2;
      const $1 = v1._2._1;
      const $2 = lookup3($EvalRef("EvalExtern", qual))(v.directives);
      const $3 = lookup22(InlineRef);
      const v2 = (() => {
        if ($2.tag === "Just") {
          return $3($2._1);
        }
        if ($2.tag === "Nothing") {
          return Nothing;
        }
        fail();
      })();
      const $4 = () => {
        if ($0.tag === "Lit" && shouldInlineExternLiteral($0._1)) {
          return $Maybe("Just", evalBackendSyntax(evalNeutralExpr).eval($1.length === 0 ? v : addStop(v)($EvalRef("EvalExtern", qual))(InlineRef))($0));
        }
        if ((v1._1.complexity === "Trivial" || v1._1.complexity === "Deref") && v1._1.size < 16) {
          return $Maybe("Just", evalBackendSyntax(evalNeutralExpr).eval($1.length === 0 ? v : addStop(v)($EvalRef("EvalExtern", qual))(InlineRef))($0));
        }
        return Nothing;
      };
      if (v2.tag === "Just") {
        if (v2._1.tag === "InlineNever") {
          return $Maybe("Just", $BackendSemantics("NeutStop", qual));
        }
        if (v2._1.tag === "InlineAlways") {
          return $Maybe("Just", evalBackendSyntax(evalNeutralExpr).eval($1.length === 0 ? v : addStop(v)($EvalRef("EvalExtern", qual))(InlineRef))($0));
        }
        if (v2._1.tag === "InlineArity") {
          return Nothing;
        }
      }
      return $4();
    }
    if (v1._2.tag === "ExternCtor" && v1._2._5.length === 0) {
      return $Maybe("Just", $BackendSemantics("NeutData", qual, v1._2._2, v1._2._3, v1._2._4, []));
    }
    return Nothing;
  }
  if (spine.length === 1) {
    if (spine[0].tag === "ExternAccessor") {
      if (spine[0]._1.tag === "GetProp") {
        if (v1._2.tag === "ExternExpr") {
          const $0 = lookup3($EvalRef("EvalExtern", qual))(v.directives);
          const $1 = lookup22($InlineAccessor("InlineProp", spine[0]._1._1));
          const v2 = (() => {
            if ($0.tag === "Just") {
              return $1($0._1);
            }
            if ($0.tag === "Nothing") {
              return Nothing;
            }
            fail();
          })();
          if (v2.tag === "Just") {
            if (v2._1.tag === "InlineNever") {
              return $Maybe("Just", neutralSpine($BackendSemantics("NeutStop", qual))(spine));
            }
            if (v2._1.tag === "InlineAlways") {
              return $Maybe(
                "Just",
                evalSpine(v)(evalBackendSyntax(evalNeutralExpr).eval(v1._2._1.length === 0 ? v : addStop(v)($EvalRef("EvalExtern", qual))($InlineAccessor("InlineProp", spine[0]._1._1)))(v1._2._2))(spine)
              );
            }
          }
          return Nothing;
        }
        if (v1._2.tag === "ExternDict") {
          const $0 = findMapImpl(
            Nothing,
            isJust,
            (v$1) => {
              if (spine[0]._1._1 === v$1._1) {
                return $Maybe("Just", v$1._2);
              }
              return Nothing;
            },
            v1._2._2
          );
          if ($0.tag === "Just") {
            const $1 = $0._1._2;
            const $2 = lookup3($EvalRef("EvalExtern", qual))(v.directives);
            const $3 = lookup22($InlineAccessor("InlineProp", spine[0]._1._1));
            const v3 = (() => {
              if ($2.tag === "Just") {
                return $3($2._1);
              }
              if ($2.tag === "Nothing") {
                return Nothing;
              }
              fail();
            })();
            const $4 = () => $Maybe(
              "Just",
              evalBackendSyntax(evalNeutralExpr).eval(v1._2._1.length === 0 ? v : addStop(v)($EvalRef("EvalExtern", qual))($InlineAccessor("InlineProp", spine[0]._1._1)))($1)
            );
            if (v3.tag === "Just") {
              if (v3._1.tag === "InlineNever") {
                return $Maybe("Just", neutralSpine($BackendSemantics("NeutStop", qual))(spine));
              }
              if (v3._1.tag === "InlineAlways") {
                return $Maybe(
                  "Just",
                  evalBackendSyntax(evalNeutralExpr).eval(v1._2._1.length === 0 ? v : addStop(v)($EvalRef("EvalExtern", qual))($InlineAccessor("InlineProp", spine[0]._1._1)))($1)
                );
              }
              if (v3._1.tag === "InlineArity") {
                return Nothing;
              }
            }
            if (($0._1._1.complexity === "Trivial" || $0._1._1.complexity === "Deref") && $0._1._1.size < 16) {
              return $4();
            }
          }
        }
      }
      return Nothing;
    }
    if (spine[0].tag === "ExternApp") {
      if (v1._2.tag === "ExternExpr") {
        const $0 = v1._2._2;
        const $1 = v1._2._1;
        const $2 = lookup3($EvalRef("EvalExtern", qual))(v.directives);
        const $3 = lookup22(InlineRef);
        const v2 = (() => {
          if ($2.tag === "Just") {
            return $3($2._1);
          }
          if ($2.tag === "Nothing") {
            return Nothing;
          }
          fail();
        })();
        const $4 = () => $Maybe(
          "Just",
          evalApp(v)(evalBackendSyntax(evalNeutralExpr).eval($1.length === 0 ? v : addStop(v)($EvalRef("EvalExtern", qual))(InlineRef))($0))(spine[0]._1)
        );
        if (v2.tag === "Just") {
          if (v2._1.tag === "InlineNever") {
            return $Maybe("Just", neutralSpine($BackendSemantics("NeutStop", qual))(spine));
          }
          if (v2._1.tag === "InlineAlways") {
            return $Maybe(
              "Just",
              evalApp(v)(evalBackendSyntax(evalNeutralExpr).eval($1.length === 0 ? v : addStop(v)($EvalRef("EvalExtern", qual))(InlineRef))($0))(spine[0]._1)
            );
          }
          if (v2._1.tag === "InlineArity") {
            if (spine[0]._1.length >= v2._1._1) {
              return $Maybe(
                "Just",
                evalApp(v)(evalBackendSyntax(evalNeutralExpr).eval($1.length === 0 ? v : addStop(v)($EvalRef("EvalExtern", qual))(InlineRef))($0))(spine[0]._1)
              );
            }
            return Nothing;
          }
        }
        if (shouldInlineExternApp(qual)(v1._1)($0)(spine[0]._1)) {
          return $4();
        }
        return Nothing;
      }
      if (v1._2.tag === "ExternCtor" && v1._2._5.length === spine[0]._1.length) {
        return $Maybe("Just", $BackendSemantics("NeutData", qual, v1._2._2, v1._2._3, v1._2._4, zipWithImpl(Tuple, v1._2._5, spine[0]._1)));
      }
    }
    return Nothing;
  }
  if (spine.length === 2) {
    if (spine[0].tag === "ExternAccessor") {
      if (spine[0]._1.tag === "GetProp" && spine[1].tag === "ExternApp") {
        if (v1._2.tag === "ExternExpr") {
          const $0 = lookup3($EvalRef("EvalExtern", qual))(v.directives);
          const $1 = lookup22($InlineAccessor("InlineProp", spine[0]._1._1));
          const v2 = (() => {
            if ($0.tag === "Just") {
              return $1($0._1);
            }
            if ($0.tag === "Nothing") {
              return Nothing;
            }
            fail();
          })();
          if (v2.tag === "Just") {
            if (v2._1.tag === "InlineNever") {
              return $Maybe("Just", neutralSpine($BackendSemantics("NeutStop", qual))(spine));
            }
            if (v2._1.tag === "InlineAlways") {
              return $Maybe(
                "Just",
                evalSpine(v)(evalBackendSyntax(evalNeutralExpr).eval(v1._2._1.length === 0 ? v : addStop(v)($EvalRef("EvalExtern", qual))($InlineAccessor("InlineProp", spine[0]._1._1)))(v1._2._2))(spine)
              );
            }
            if (v2._1.tag === "InlineArity" && spine[1]._1.length >= v2._1._1) {
              return $Maybe(
                "Just",
                evalSpine(v)(evalBackendSyntax(evalNeutralExpr).eval(v1._2._1.length === 0 ? v : addStop(v)($EvalRef("EvalExtern", qual))($InlineAccessor("InlineProp", spine[0]._1._1)))(v1._2._2))(spine)
              );
            }
          }
          return Nothing;
        }
        if (v1._2.tag === "ExternDict") {
          const $0 = findMapImpl(
            Nothing,
            isJust,
            (v$1) => {
              if (spine[0]._1._1 === v$1._1) {
                return $Maybe("Just", v$1._2);
              }
              return Nothing;
            },
            v1._2._2
          );
          if ($0.tag === "Just") {
            const $1 = $0._1._2;
            const $2 = lookup3($EvalRef("EvalExtern", qual))(v.directives);
            const $3 = lookup22($InlineAccessor("InlineProp", spine[0]._1._1));
            const v3 = (() => {
              if ($2.tag === "Just") {
                return $3($2._1);
              }
              if ($2.tag === "Nothing") {
                return Nothing;
              }
              fail();
            })();
            const $4 = () => $Maybe(
              "Just",
              evalApp(v)(evalBackendSyntax(evalNeutralExpr).eval(v1._2._1.length === 0 ? v : addStop(v)($EvalRef("EvalExtern", qual))($InlineAccessor("InlineProp", spine[0]._1._1)))($1))(spine[1]._1)
            );
            if (v3.tag === "Just") {
              if (v3._1.tag === "InlineNever") {
                return $Maybe("Just", neutralSpine($BackendSemantics("NeutStop", qual))(spine));
              }
              if (v3._1.tag === "InlineAlways") {
                return $Maybe(
                  "Just",
                  evalApp(v)(evalBackendSyntax(evalNeutralExpr).eval(v1._2._1.length === 0 ? v : addStop(v)($EvalRef("EvalExtern", qual))($InlineAccessor("InlineProp", spine[0]._1._1)))($1))(spine[1]._1)
                );
              }
              if (v3._1.tag === "InlineArity") {
                if (spine[1]._1.length >= v3._1._1) {
                  return $Maybe(
                    "Just",
                    evalApp(v)(evalBackendSyntax(evalNeutralExpr).eval(v1._2._1.length === 0 ? v : addStop(v)($EvalRef("EvalExtern", qual))($InlineAccessor("InlineProp", spine[0]._1._1)))($1))(spine[1]._1)
                  );
                }
                return Nothing;
              }
            }
            if (shouldInlineExternApp(qual)($0._1._1)($1)(spine[1]._1)) {
              return $4();
            }
          }
        }
      }
      return Nothing;
    }
    if (spine[0].tag === "ExternApp" && spine[1].tag === "ExternAccessor" && spine[1]._1.tag === "GetProp" && v1._2.tag === "ExternExpr") {
      const $0 = lookup3($EvalRef("EvalExtern", qual))(v.directives);
      const $1 = lookup22($InlineAccessor("InlineSpineProp", spine[1]._1._1));
      const v2 = (() => {
        if ($0.tag === "Just") {
          return $1($0._1);
        }
        if ($0.tag === "Nothing") {
          return Nothing;
        }
        fail();
      })();
      if (v2.tag === "Just") {
        if (v2._1.tag === "InlineNever") {
          return $Maybe("Just", neutralSpine($BackendSemantics("NeutStop", qual))(spine));
        }
        if (v2._1.tag === "InlineAlways") {
          return $Maybe(
            "Just",
            evalSpine(v)(evalBackendSyntax(evalNeutralExpr).eval(v1._2._1.length === 0 ? v : addStop(v)($EvalRef("EvalExtern", qual))($InlineAccessor("InlineSpineProp", spine[1]._1._1)))(v1._2._2))(spine)
          );
        }
      }
    }
    return Nothing;
  }
  if (spine.length === 3 && spine[0].tag === "ExternApp" && spine[1].tag === "ExternAccessor" && spine[1]._1.tag === "GetProp" && spine[2].tag === "ExternApp" && v1._2.tag === "ExternExpr") {
    const $0 = lookup3($EvalRef("EvalExtern", qual))(v.directives);
    const $1 = lookup22($InlineAccessor("InlineSpineProp", spine[1]._1._1));
    const v2 = (() => {
      if ($0.tag === "Just") {
        return $1($0._1);
      }
      if ($0.tag === "Nothing") {
        return Nothing;
      }
      fail();
    })();
    if (v2.tag === "Just") {
      if (v2._1.tag === "InlineNever") {
        return $Maybe("Just", neutralSpine($BackendSemantics("NeutStop", qual))(spine));
      }
      if (v2._1.tag === "InlineAlways") {
        return $Maybe(
          "Just",
          evalSpine(v)(evalBackendSyntax(evalNeutralExpr).eval(v1._2._1.length === 0 ? v : addStop(v)($EvalRef("EvalExtern", qual))($InlineAccessor("InlineSpineProp", spine[1]._1._1)))(v1._2._2))(spine)
        );
      }
      if (v2._1.tag === "InlineArity" && spine[2]._1.length >= v2._1._1) {
        return $Maybe(
          "Just",
          evalSpine(v)(evalBackendSyntax(evalNeutralExpr).eval(v1._2._1.length === 0 ? v : addStop(v)($EvalRef("EvalExtern", qual))($InlineAccessor("InlineSpineProp", spine[1]._1._1)))(v1._2._2))(spine)
        );
      }
    }
  }
  return Nothing;
};
var evalExternRefFromImpl = (env) => (qual) => (v) => {
  if (v._2.tag === "ExternExpr") {
    if (v._2._2.tag === "Var" || v._2._2.tag === "Lit" || v._2._2.tag === "CtorSaturated" || v._2._2.tag === "Accessor" || v._2._2.tag === "Update" || v._2._2.tag === "PrimOp") {
      return eval3(v._2._1.length === 0 ? env : addStop(env)($EvalRef("EvalExtern", qual))(InlineRef))(v._2._2);
    }
    return $BackendSemantics("NeutVar", qual);
  }
  if (v._2.tag === "ExternDict") {
    const $0 = v._2._1;
    return $BackendSemantics(
      "NeutLit",
      $Literal(
        "LitRecord",
        arrayMap((v1) => $Prop(
          v1._1,
          eval3($0.length === 0 ? env : addStop(env)($EvalRef("EvalExtern", qual))($InlineAccessor("InlineProp", v1._1)))(v1._2._2)
        ))(v._2._2)
      )
    );
  }
  return $BackendSemantics("NeutVar", qual);
};

// output-es/PureScript.CST.Errors/index.js
var $ParseError = (tag, _1, _2) => ({ tag, _1, _2 });
var UnexpectedEof = /* @__PURE__ */ $ParseError("UnexpectedEof");
var LexExpected = (value0) => (value1) => $ParseError("LexExpected", value0, value1);

// output-es/Data.String.Regex/foreign.js
var regexImpl = function(left) {
  return function(right) {
    return function(s1) {
      return function(s2) {
        try {
          return right(new RegExp(s1, s2));
        } catch (e) {
          return left(e.message);
        }
      };
    };
  };
};
var _match = function(just) {
  return function(nothing) {
    return function(r) {
      return function(s) {
        var m = s.match(r);
        if (m == null || m.length === 0) {
          return nothing;
        } else {
          for (var i = 0; i < m.length; i++) {
            m[i] = m[i] == null ? nothing : just(m[i]);
          }
          return just(m);
        }
      };
    };
  };
};

// output-es/Data.String.Regex/index.js
var regex = (s) => (f) => regexImpl(Left)(Right)(s)((f.global ? "g" : "") + (f.ignoreCase ? "i" : "") + (f.multiline ? "m" : "") + (f.dotAll ? "s" : "") + (f.sticky ? "y" : "") + (f.unicode ? "u" : ""));
var match = /* @__PURE__ */ _match(Just)(Nothing);

// output-es/Data.String.Regex.Flags/index.js
var unicode = { global: false, ignoreCase: false, multiline: false, dotAll: false, sticky: false, unicode: true };

// output-es/Data.String.Regex.Unsafe/index.js
var unsafeRegex = (s) => (f) => {
  const $0 = regex(s)(f);
  if ($0.tag === "Left") {
    return _crashWith($0._1);
  }
  if ($0.tag === "Right") {
    return $0._1;
  }
  fail();
};

// output-es/PureScript.CST.Types/index.js
var $Comment2 = (tag, _1, _2) => ({ tag, _1, _2 });
var $IntValue = (tag, _1) => ({ tag, _1 });
var $LineFeed = (tag) => tag;
var $SourceStyle = (tag) => tag;
var $Token = (tag, _1, _2) => ({ tag, _1, _2 });
var ASCII = /* @__PURE__ */ $SourceStyle("ASCII");
var Unicode = /* @__PURE__ */ $SourceStyle("Unicode");
var LF = /* @__PURE__ */ $LineFeed("LF");
var CRLF = /* @__PURE__ */ $LineFeed("CRLF");
var TokLeftParen = /* @__PURE__ */ $Token("TokLeftParen");
var TokRightParen = /* @__PURE__ */ $Token("TokRightParen");
var TokLeftBrace = /* @__PURE__ */ $Token("TokLeftBrace");
var TokRightBrace = /* @__PURE__ */ $Token("TokRightBrace");
var TokLeftSquare = /* @__PURE__ */ $Token("TokLeftSquare");
var TokRightSquare = /* @__PURE__ */ $Token("TokRightSquare");
var TokEquals = /* @__PURE__ */ $Token("TokEquals");
var TokPipe = /* @__PURE__ */ $Token("TokPipe");
var TokTick = /* @__PURE__ */ $Token("TokTick");
var TokDot = /* @__PURE__ */ $Token("TokDot");
var TokComma = /* @__PURE__ */ $Token("TokComma");
var TokUnderscore = /* @__PURE__ */ $Token("TokUnderscore");
var TokBackslash = /* @__PURE__ */ $Token("TokBackslash");
var TokAt = /* @__PURE__ */ $Token("TokAt");
var Line = (value0) => (value1) => $Comment2("Line", value0, value1);

// output-es/PureScript.CST.Layout/index.js
var $LayoutDelim = (tag) => tag;
var LytRoot = /* @__PURE__ */ $LayoutDelim("LytRoot");
var LytTopDecl = /* @__PURE__ */ $LayoutDelim("LytTopDecl");
var LytTopDeclHead = /* @__PURE__ */ $LayoutDelim("LytTopDeclHead");
var LytDeclGuard = /* @__PURE__ */ $LayoutDelim("LytDeclGuard");
var LytCase = /* @__PURE__ */ $LayoutDelim("LytCase");
var LytCaseBinders = /* @__PURE__ */ $LayoutDelim("LytCaseBinders");
var LytCaseGuard = /* @__PURE__ */ $LayoutDelim("LytCaseGuard");
var LytLambdaBinders = /* @__PURE__ */ $LayoutDelim("LytLambdaBinders");
var LytParen = /* @__PURE__ */ $LayoutDelim("LytParen");
var LytBrace = /* @__PURE__ */ $LayoutDelim("LytBrace");
var LytSquare = /* @__PURE__ */ $LayoutDelim("LytSquare");
var LytIf = /* @__PURE__ */ $LayoutDelim("LytIf");
var LytThen = /* @__PURE__ */ $LayoutDelim("LytThen");
var LytProperty = /* @__PURE__ */ $LayoutDelim("LytProperty");
var LytForall = /* @__PURE__ */ $LayoutDelim("LytForall");
var LytTick = /* @__PURE__ */ $LayoutDelim("LytTick");
var LytLet = /* @__PURE__ */ $LayoutDelim("LytLet");
var LytLetStmt = /* @__PURE__ */ $LayoutDelim("LytLetStmt");
var LytWhere = /* @__PURE__ */ $LayoutDelim("LytWhere");
var LytOf = /* @__PURE__ */ $LayoutDelim("LytOf");
var LytDo = /* @__PURE__ */ $LayoutDelim("LytDo");
var LytAdo = /* @__PURE__ */ $LayoutDelim("LytAdo");
var isIndented = (v) => v === "LytLet" || v === "LytLetStmt" || v === "LytWhere" || v === "LytOf" || v === "LytDo" || v === "LytAdo";
var eqLayoutDelim = {
  eq: (x) => (y) => {
    if (x === "LytRoot") {
      return y === "LytRoot";
    }
    if (x === "LytTopDecl") {
      return y === "LytTopDecl";
    }
    if (x === "LytTopDeclHead") {
      return y === "LytTopDeclHead";
    }
    if (x === "LytDeclGuard") {
      return y === "LytDeclGuard";
    }
    if (x === "LytCase") {
      return y === "LytCase";
    }
    if (x === "LytCaseBinders") {
      return y === "LytCaseBinders";
    }
    if (x === "LytCaseGuard") {
      return y === "LytCaseGuard";
    }
    if (x === "LytLambdaBinders") {
      return y === "LytLambdaBinders";
    }
    if (x === "LytParen") {
      return y === "LytParen";
    }
    if (x === "LytBrace") {
      return y === "LytBrace";
    }
    if (x === "LytSquare") {
      return y === "LytSquare";
    }
    if (x === "LytIf") {
      return y === "LytIf";
    }
    if (x === "LytThen") {
      return y === "LytThen";
    }
    if (x === "LytProperty") {
      return y === "LytProperty";
    }
    if (x === "LytForall") {
      return y === "LytForall";
    }
    if (x === "LytTick") {
      return y === "LytTick";
    }
    if (x === "LytLet") {
      return y === "LytLet";
    }
    if (x === "LytLetStmt") {
      return y === "LytLetStmt";
    }
    if (x === "LytWhere") {
      return y === "LytWhere";
    }
    if (x === "LytOf") {
      return y === "LytOf";
    }
    if (x === "LytDo") {
      return y === "LytDo";
    }
    return x === "LytAdo" && y === "LytAdo";
  }
};
var insertLayout = (v) => (nextPos) => (stack) => {
  const tokPos = v.range.start;
  const insertStart = (lyt) => (v1) => {
    const go = (go$a0$copy) => (go$a1$copy) => {
      let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
      while (go$c) {
        const b = go$a0, v$1 = go$a1;
        if (v$1.tag === "Nil") {
          go$c = false;
          go$r = b;
          continue;
        }
        if (v$1.tag === "Cons") {
          go$a0 = b.tag === "Nothing" && (v$1._1._2 === "LytLet" || v$1._1._2 === "LytLetStmt" || v$1._1._2 === "LytWhere" || v$1._1._2 === "LytOf" || v$1._1._2 === "LytDo" || v$1._1._2 === "LytAdo") ? $Maybe("Just", v$1._1) : b;
          go$a1 = v$1._2;
          continue;
        }
        fail();
      }
      return go$r;
    };
    const v2 = go(Nothing)(v1._1);
    if (v2.tag === "Just" && nextPos.column <= v2._1._1.column) {
      return v1;
    }
    const $02 = $List("Cons", $Tuple(nextPos, lyt), v1._1);
    return $Tuple(
      $02,
      snoc(v1._2)($Tuple(
        { range: { start: nextPos, end: nextPos }, leadingComments: [], trailingComments: [], value: $Token("TokLayoutStart", nextPos.column) },
        $02
      ))
    );
  };
  const insertSep = (v1) => {
    const sepTok = { range: { start: tokPos, end: tokPos }, leadingComments: [], trailingComments: [], value: $Token("TokLayoutSep", tokPos.column) };
    const $02 = (lyt, lytPos) => {
      if (lyt === "LytOf") {
        return $Tuple($List("Cons", $Tuple(tokPos, LytCaseBinders), v1._1), snoc(v1._2)($Tuple(sepTok, v1._1)));
      }
      return $Tuple(v1._1, snoc(v1._2)($Tuple(sepTok, v1._1)));
    };
    if (v1._1.tag === "Cons") {
      if (v1._1._1._2 === "LytTopDecl") {
        if (tokPos.column === v1._1._1._1.column && tokPos.line !== v1._1._1._1.line) {
          return $Tuple(v1._1._2, snoc(v1._2)($Tuple(sepTok, v1._1._2)));
        }
        if ((v1._1._1._2 === "LytLet" || v1._1._1._2 === "LytLetStmt" || v1._1._1._2 === "LytWhere" || v1._1._1._2 === "LytOf" || v1._1._1._2 === "LytDo" || v1._1._1._2 === "LytAdo") && tokPos.column === v1._1._1._1.column && tokPos.line !== v1._1._1._1.line) {
          return $02(v1._1._1._2, v1._1._1._1);
        }
        return v1;
      }
      if (v1._1._1._2 === "LytTopDeclHead" && tokPos.column === v1._1._1._1.column && tokPos.line !== v1._1._1._1.line) {
        return $Tuple(v1._1._2, snoc(v1._2)($Tuple(sepTok, v1._1._2)));
      }
      if ((v1._1._1._2 === "LytLet" || v1._1._1._2 === "LytLetStmt" || v1._1._1._2 === "LytWhere" || v1._1._1._2 === "LytOf" || v1._1._1._2 === "LytDo" || v1._1._1._2 === "LytAdo") && tokPos.column === v1._1._1._1.column && tokPos.line !== v1._1._1._1.line) {
        return $02(v1._1._1._2, v1._1._1._1);
      }
    }
    return v1;
  };
  const collapse = (p) => {
    const go = (go$a0$copy) => (go$a1$copy) => {
      let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
      while (go$c) {
        const v1 = go$a0, v2 = go$a1;
        if (v1.tag === "Cons" && p(v1._1._1)(v1._1._2)) {
          go$a0 = v1._2;
          go$a1 = v1._1._2 === "LytLet" || v1._1._2 === "LytLetStmt" || v1._1._2 === "LytWhere" || v1._1._2 === "LytOf" || v1._1._2 === "LytDo" || v1._1._2 === "LytAdo" ? snoc(v2)($Tuple(
            { range: { start: tokPos, end: tokPos }, leadingComments: [], trailingComments: [], value: $Token("TokLayoutEnd", v1._1._1.column) },
            v1._2
          )) : v2;
          continue;
        }
        go$c = false;
        go$r = $Tuple(v1, v2);
      }
      return go$r;
    };
    return (v$1) => go(v$1._1)(v$1._2);
  };
  const insertKwProperty = (k, state) => {
    const $02 = insertSep(collapse((lytPos) => (lyt) => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)(state));
    const v1 = $Tuple($02._1, snoc($02._2)($Tuple(v, $02._1)));
    if (v1._1.tag === "Cons" && v1._1._1._2 === "LytProperty") {
      return $Tuple(v1._1._2, v1._2);
    }
    return k(v1);
  };
  if (v.value.tag === "TokLowerName") {
    if (v.value._1.tag === "Nothing") {
      if (v.value._2 === "data") {
        const $04 = insertSep(collapse((lytPos) => (lyt) => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)($Tuple(
          stack,
          []
        )));
        const v2 = $Tuple($04._1, snoc($04._2)($Tuple(v, $04._1)));
        if (v2._1.tag === "Cons" && v2._1._1._2 === "LytWhere" && v2._1._2.tag === "Cons" && v2._1._2._1._2 === "LytRoot" && v2._1._2._2.tag === "Nil" && tokPos.column === v2._1._1._1.column) {
          return $Tuple($List("Cons", $Tuple(tokPos, LytTopDecl), v2._1), v2._2);
        }
        if (v2._1.tag === "Cons" && eqLayoutDelim.eq(v2._1._1._2)(LytProperty)) {
          return $Tuple(v2._1._2, v2._2);
        }
        return v2;
      }
      if (v.value._2 === "class") {
        const $04 = insertSep(collapse((lytPos) => (lyt) => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)($Tuple(
          stack,
          []
        )));
        const v2 = $Tuple($04._1, snoc($04._2)($Tuple(v, $04._1)));
        if (v2._1.tag === "Cons" && v2._1._1._2 === "LytWhere" && v2._1._2.tag === "Cons" && v2._1._2._1._2 === "LytRoot" && v2._1._2._2.tag === "Nil" && tokPos.column === v2._1._1._1.column) {
          return $Tuple($List("Cons", $Tuple(tokPos, LytTopDeclHead), v2._1), v2._2);
        }
        if (v2._1.tag === "Cons" && eqLayoutDelim.eq(v2._1._1._2)(LytProperty)) {
          return $Tuple(v2._1._2, v2._2);
        }
        return v2;
      }
      if (v.value._2 === "where") {
        const whereP = (v2) => (v3) => v3 === "LytDo" || (v3 === "LytLet" || v3 === "LytLetStmt" || v3 === "LytWhere" || v3 === "LytOf" || v3 === "LytDo" || v3 === "LytAdo") && tokPos.column <= v2.column;
        if (stack.tag === "Cons") {
          if (stack._1._2 === "LytTopDeclHead") {
            return insertStart(LytWhere)($Tuple(stack._2, snoc([])($Tuple(v, stack._2))));
          }
          if (stack._1._2 === "LytProperty") {
            return $Tuple(stack._2, snoc([])($Tuple(v, stack._2)));
          }
        }
        return insertStart(LytWhere)((() => {
          const $04 = collapse(whereP)($Tuple(stack, []));
          return $Tuple($04._1, snoc($04._2)($Tuple(v, $04._1)));
        })());
      }
      if (v.value._2 === "in") {
        const v2 = collapse((v22) => (v3) => {
          if (v3 === "LytLet") {
            return false;
          }
          if (v3 === "LytAdo") {
            return false;
          }
          return v3 === "LytLet" || v3 === "LytLetStmt" || v3 === "LytWhere" || v3 === "LytOf" || v3 === "LytDo" || v3 === "LytAdo";
        })($Tuple(stack, []));
        if (v2._1.tag === "Cons") {
          if (v2._1._1._2 === "LytLetStmt" && v2._1._2.tag === "Cons" && v2._1._2._1._2 === "LytAdo") {
            return $Tuple(
              v2._1._2._2,
              snoc(snoc(snoc(v2._2)($Tuple(
                { range: { start: tokPos, end: tokPos }, leadingComments: [], trailingComments: [], value: $Token("TokLayoutEnd", v2._1._1._1.column) },
                v2._1._2._2
              )))($Tuple(
                { range: { start: tokPos, end: tokPos }, leadingComments: [], trailingComments: [], value: $Token("TokLayoutEnd", v2._1._2._1._1.column) },
                v2._1._2._2
              )))($Tuple(v, v2._1._2._2))
            );
          }
          if (v2._1._1._2 === "LytLet" || v2._1._1._2 === "LytLetStmt" || v2._1._1._2 === "LytWhere" || v2._1._1._2 === "LytOf" || v2._1._1._2 === "LytDo" || v2._1._1._2 === "LytAdo") {
            return $Tuple(
              v2._1._2,
              snoc(snoc(v2._2)($Tuple(
                { range: { start: tokPos, end: tokPos }, leadingComments: [], trailingComments: [], value: $Token("TokLayoutEnd", v2._1._1._1.column) },
                v2._1._2
              )))($Tuple(v, v2._1._2))
            );
          }
        }
        const $04 = insertSep(collapse((lytPos) => (lyt) => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)($Tuple(
          stack,
          []
        )));
        const $12 = $Tuple($04._1, snoc($04._2)($Tuple(v, $04._1)));
        if ($12._1.tag === "Cons" && eqLayoutDelim.eq($12._1._1._2)(LytProperty)) {
          return $Tuple($12._1._2, $12._2);
        }
        return $12;
      }
      if (v.value._2 === "let") {
        return insertKwProperty(
          (v2) => {
            if (v2._1.tag === "Cons") {
              if (v2._1._1._2 === "LytDo") {
                if (v2._1._1._1.column === tokPos.column) {
                  return insertStart(LytLetStmt)(v2);
                }
                return insertStart(LytLet)(v2);
              }
              if (v2._1._1._2 === "LytAdo" && v2._1._1._1.column === tokPos.column) {
                return insertStart(LytLetStmt)(v2);
              }
            }
            return insertStart(LytLet)(v2);
          },
          $Tuple(stack, [])
        );
      }
      if (v.value._2 === "do") {
        return insertKwProperty(insertStart(LytDo), $Tuple(stack, []));
      }
      if (v.value._2 === "ado") {
        return insertKwProperty(insertStart(LytAdo), $Tuple(stack, []));
      }
      if (v.value._2 === "case") {
        return insertKwProperty((v1) => $Tuple($List("Cons", $Tuple(tokPos, LytCase), v1._1), v1._2), $Tuple(stack, []));
      }
      if (v.value._2 === "of") {
        const v2 = collapse((v$1) => isIndented)($Tuple(stack, []));
        if (v2._1.tag === "Cons" && v2._1._1._2 === "LytCase") {
          const $05 = insertStart(LytOf)($Tuple(v2._1._2, snoc(v2._2)($Tuple(v, v2._1._2))));
          return $Tuple($List("Cons", $Tuple(nextPos, LytCaseBinders), $05._1), $05._2);
        }
        const $04 = insertSep(collapse((lytPos) => (lyt) => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)(v2));
        const $12 = $Tuple($04._1, snoc($04._2)($Tuple(v, $04._1)));
        if ($12._1.tag === "Cons" && eqLayoutDelim.eq($12._1._1._2)(LytProperty)) {
          return $Tuple($12._1._2, $12._2);
        }
        return $12;
      }
      if (v.value._2 === "if") {
        return insertKwProperty((v1) => $Tuple($List("Cons", $Tuple(tokPos, LytIf), v1._1), v1._2), $Tuple(stack, []));
      }
      if (v.value._2 === "then") {
        const v2 = collapse((v$1) => isIndented)($Tuple(stack, []));
        if (v2._1.tag === "Cons" && v2._1._1._2 === "LytIf") {
          return $Tuple($List("Cons", $Tuple(tokPos, LytThen), v2._1._2), snoc(v2._2)($Tuple(v, v2._1._2)));
        }
        const $04 = insertSep(collapse((lytPos) => (lyt) => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)($Tuple(
          stack,
          []
        )));
        const $12 = $Tuple($04._1, snoc($04._2)($Tuple(v, $04._1)));
        if ($12._1.tag === "Cons" && eqLayoutDelim.eq($12._1._1._2)(LytProperty)) {
          return $Tuple($12._1._2, $12._2);
        }
        return $12;
      }
      if (v.value._2 === "else") {
        const v2 = collapse((v$1) => isIndented)($Tuple(stack, []));
        if (v2._1.tag === "Cons" && v2._1._1._2 === "LytThen") {
          return $Tuple(v2._1._2, snoc(v2._2)($Tuple(v, v2._1._2)));
        }
        const v3 = collapse((lytPos) => (lyt) => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)($Tuple(
          stack,
          []
        ));
        if (v3._1.tag === "Cons" && v3._1._1._2 === "LytWhere" && v3._1._2.tag === "Cons" && v3._1._2._1._2 === "LytRoot" && v3._1._2._2.tag === "Nil" && tokPos.column === v3._1._1._1.column) {
          return $Tuple(v3._1, snoc(v3._2)($Tuple(v, v3._1)));
        }
        const $04 = insertSep(v3);
        const $12 = $Tuple($04._1, snoc($04._2)($Tuple(v, $04._1)));
        if ($12._1.tag === "Cons" && eqLayoutDelim.eq($12._1._1._2)(LytProperty)) {
          return $Tuple($12._1._2, $12._2);
        }
        return $12;
      }
      const $03 = insertSep(collapse((lytPos) => (lyt) => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)($Tuple(
        stack,
        []
      )));
      const $1 = $Tuple($03._1, snoc($03._2)($Tuple(v, $03._1)));
      if ($1._1.tag === "Cons" && eqLayoutDelim.eq($1._1._1._2)(LytProperty)) {
        return $Tuple($1._1._2, $1._2);
      }
      return $1;
    }
    if (v.value._2 === "do") {
      return insertKwProperty(insertStart(LytDo), $Tuple(stack, []));
    }
    if (v.value._2 === "ado") {
      return insertKwProperty(insertStart(LytAdo), $Tuple(stack, []));
    }
    const $02 = insertSep(collapse((lytPos) => (lyt) => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)($Tuple(
      stack,
      []
    )));
    return $Tuple($02._1, snoc($02._2)($Tuple(v, $02._1)));
  }
  if (v.value.tag === "TokForall") {
    return insertKwProperty((v1) => $Tuple($List("Cons", $Tuple(tokPos, LytForall), v1._1), v1._2), $Tuple(stack, []));
  }
  if (v.value.tag === "TokBackslash") {
    const $02 = insertSep(collapse((lytPos) => (lyt) => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)($Tuple(
      stack,
      []
    )));
    return $Tuple($List("Cons", $Tuple(tokPos, LytLambdaBinders), $02._1), snoc($02._2)($Tuple(v, $02._1)));
  }
  if (v.value.tag === "TokRightArrow") {
    const $02 = collapse((v2) => (v3) => {
      if (v3 === "LytDo") {
        return true;
      }
      if (v3 === "LytOf") {
        return false;
      }
      return (v3 === "LytLet" || v3 === "LytLetStmt" || v3 === "LytWhere" || v3 === "LytOf" || v3 === "LytDo" || v3 === "LytAdo") && tokPos.column <= v2.column;
    })($Tuple(stack, []));
    if ($02._1.tag === "Cons" && ($02._1._1._2 === "LytCaseBinders" || $02._1._1._2 === "LytCaseGuard" || $02._1._1._2 === "LytLambdaBinders")) {
      return $Tuple($02._1._2, snoc($02._2)($Tuple(v, $02._1._2)));
    }
    return $Tuple($02._1, snoc($02._2)($Tuple(v, $02._1)));
  }
  if (v.value.tag === "TokEquals") {
    const v2 = collapse((v22) => (v3) => v3 === "LytWhere" || v3 === "LytLet" || v3 === "LytLetStmt")($Tuple(stack, []));
    if (v2._1.tag === "Cons" && v2._1._1._2 === "LytDeclGuard") {
      return $Tuple(v2._1._2, snoc(v2._2)($Tuple(v, v2._1._2)));
    }
    const $02 = insertSep(collapse((lytPos) => (lyt) => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)($Tuple(
      stack,
      []
    )));
    return $Tuple($02._1, snoc($02._2)($Tuple(v, $02._1)));
  }
  if (v.value.tag === "TokPipe") {
    const v2 = collapse((lytPos) => (lyt) => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column <= lytPos.column)($Tuple(
      stack,
      []
    ));
    if (v2._1.tag === "Cons") {
      if (v2._1._1._2 === "LytOf") {
        const $03 = $List("Cons", $Tuple(tokPos, LytCaseGuard), v2._1);
        return $Tuple($03, snoc(v2._2)($Tuple(v, $03)));
      }
      if (v2._1._1._2 === "LytLet") {
        const $03 = $List("Cons", $Tuple(tokPos, LytDeclGuard), v2._1);
        return $Tuple($03, snoc(v2._2)($Tuple(v, $03)));
      }
      if (v2._1._1._2 === "LytLetStmt") {
        const $03 = $List("Cons", $Tuple(tokPos, LytDeclGuard), v2._1);
        return $Tuple($03, snoc(v2._2)($Tuple(v, $03)));
      }
      if (v2._1._1._2 === "LytWhere") {
        const $03 = $List("Cons", $Tuple(tokPos, LytDeclGuard), v2._1);
        return $Tuple($03, snoc(v2._2)($Tuple(v, $03)));
      }
    }
    const $02 = insertSep(collapse((lytPos) => (lyt) => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)($Tuple(
      stack,
      []
    )));
    return $Tuple($02._1, snoc($02._2)($Tuple(v, $02._1)));
  }
  if (v.value.tag === "TokTick") {
    const v2 = collapse((v$1) => isIndented)($Tuple(stack, []));
    if (v2._1.tag === "Cons" && v2._1._1._2 === "LytTick") {
      return $Tuple(v2._1._2, snoc(v2._2)($Tuple(v, v2._1._2)));
    }
    const $02 = insertSep(collapse((lytPos) => (lyt) => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column <= lytPos.column)($Tuple(
      stack,
      []
    )));
    return $Tuple($List("Cons", $Tuple(tokPos, LytTick), $02._1), snoc($02._2)($Tuple(v, $02._1)));
  }
  if (v.value.tag === "TokComma") {
    const v2 = collapse((v$1) => isIndented)($Tuple(stack, []));
    if (v2._1.tag === "Cons" && v2._1._1._2 === "LytBrace") {
      return $Tuple($List("Cons", $Tuple(tokPos, LytProperty), v2._1), snoc(v2._2)($Tuple(v, v2._1)));
    }
    return $Tuple(v2._1, snoc(v2._2)($Tuple(v, v2._1)));
  }
  if (v.value.tag === "TokDot") {
    const $02 = insertSep(collapse((lytPos) => (lyt) => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)($Tuple(
      stack,
      []
    )));
    const $1 = snoc($02._2)($Tuple(v, $02._1));
    if ($02._1.tag === "Cons" && $02._1._1._2 === "LytForall") {
      return $Tuple($02._1._2, $1);
    }
    return $Tuple($List("Cons", $Tuple(tokPos, LytProperty), $02._1), $1);
  }
  if (v.value.tag === "TokLeftParen") {
    const $02 = insertSep(collapse((lytPos) => (lyt) => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)($Tuple(
      stack,
      []
    )));
    return $Tuple($List("Cons", $Tuple(tokPos, LytParen), $02._1), snoc($02._2)($Tuple(v, $02._1)));
  }
  if (v.value.tag === "TokLeftBrace") {
    const $02 = insertSep(collapse((lytPos) => (lyt) => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)($Tuple(
      stack,
      []
    )));
    return $Tuple(
      $List("Cons", $Tuple(tokPos, LytProperty), $List("Cons", $Tuple(tokPos, LytBrace), $02._1)),
      snoc($02._2)($Tuple(v, $02._1))
    );
  }
  if (v.value.tag === "TokLeftSquare") {
    const $02 = insertSep(collapse((lytPos) => (lyt) => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)($Tuple(
      stack,
      []
    )));
    return $Tuple($List("Cons", $Tuple(tokPos, LytSquare), $02._1), snoc($02._2)($Tuple(v, $02._1)));
  }
  if (v.value.tag === "TokRightParen") {
    const $02 = collapse((v$1) => isIndented)($Tuple(stack, []));
    if ($02._1.tag === "Cons" && eqLayoutDelim.eq($02._1._1._2)(LytParen)) {
      return $Tuple($02._1._2, snoc($02._2)($Tuple(v, $02._1._2)));
    }
    return $Tuple($02._1, snoc($02._2)($Tuple(v, $02._1)));
  }
  if (v.value.tag === "TokRightBrace") {
    const $02 = collapse((v$1) => isIndented)($Tuple(stack, []));
    if ($02._1.tag === "Cons" && eqLayoutDelim.eq($02._1._1._2)(LytProperty)) {
      if ($02._1._2.tag === "Cons" && eqLayoutDelim.eq($02._1._2._1._2)(LytBrace)) {
        return $Tuple($02._1._2._2, snoc($02._2)($Tuple(v, $02._1._2._2)));
      }
      return $Tuple($02._1._2, snoc($02._2)($Tuple(v, $02._1._2)));
    }
    if ($02._1.tag === "Cons" && eqLayoutDelim.eq($02._1._1._2)(LytBrace)) {
      return $Tuple($02._1._2, snoc($02._2)($Tuple(v, $02._1._2)));
    }
    return $Tuple($02._1, snoc($02._2)($Tuple(v, $02._1)));
  }
  if (v.value.tag === "TokRightSquare") {
    const $02 = collapse((v$1) => isIndented)($Tuple(stack, []));
    if ($02._1.tag === "Cons" && eqLayoutDelim.eq($02._1._1._2)(LytSquare)) {
      return $Tuple($02._1._2, snoc($02._2)($Tuple(v, $02._1._2)));
    }
    return $Tuple($02._1, snoc($02._2)($Tuple(v, $02._1)));
  }
  if (v.value.tag === "TokString") {
    const $02 = insertSep(collapse((lytPos) => (lyt) => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)($Tuple(
      stack,
      []
    )));
    const $1 = $Tuple($02._1, snoc($02._2)($Tuple(v, $02._1)));
    if ($1._1.tag === "Cons" && eqLayoutDelim.eq($1._1._1._2)(LytProperty)) {
      return $Tuple($1._1._2, $1._2);
    }
    return $1;
  }
  if (v.value.tag === "TokOperator") {
    const $02 = insertSep(collapse((lytPos) => (lyt) => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column <= lytPos.column)($Tuple(
      stack,
      []
    )));
    return $Tuple($02._1, snoc($02._2)($Tuple(v, $02._1)));
  }
  const $0 = insertSep(collapse((lytPos) => (lyt) => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)($Tuple(
    stack,
    []
  )));
  return $Tuple($0._1, snoc($0._2)($Tuple(v, $0._1)));
};

// output-es/PureScript.CST.TokenStream/index.js
var $TokenStep = (tag, _1, _2, _3, _4) => ({ tag, _1, _2, _3, _4 });
var unwindLayout = (pos) => (eof2) => {
  const go = (stk) => defer((v) => {
    if (stk.tag === "Nil") {
      return force(eof2);
    }
    if (stk.tag === "Cons") {
      if (stk._1._2 === "LytRoot") {
        return force(eof2);
      }
      if (stk._1._2 === "LytLet" || stk._1._2 === "LytLetStmt" || stk._1._2 === "LytWhere" || stk._1._2 === "LytOf" || stk._1._2 === "LytDo" || stk._1._2 === "LytAdo") {
        return $TokenStep(
          "TokenCons",
          { range: { start: pos, end: pos }, leadingComments: [], trailingComments: [], value: $Token("TokLayoutEnd", stk._1._1.column) },
          pos,
          go(stk._2),
          stk._2
        );
      }
      return force(go(stk._2));
    }
    fail();
  });
  return go;
};
var consTokens = (dictFoldable) => {
  const $0 = dictFoldable.foldr((v) => (v1) => {
    const $02 = v1._2;
    const $1 = v1._1;
    const $2 = v._2;
    const $3 = v._1;
    return $Tuple($3.range.start, defer((v2) => $TokenStep("TokenCons", $3, $1, $02, $2)));
  });
  return (b) => (a) => $0(a)(b);
};

// output-es/PureScript.CST.Lexer/index.js
var $LexResult = (tag, _1, _2) => ({ tag, _1, _2 });
var fold1 = /* @__PURE__ */ (() => foldableArray.foldMap(/* @__PURE__ */ (() => {
  const semigroupRecord1 = { append: (ra) => (rb) => ({ raw: ra.raw + rb.raw, string: ra.string + rb.string }) };
  return { mempty: { raw: "", string: "" }, Semigroup0: () => semigroupRecord1 };
})())(identity2))();
var consTokens2 = /* @__PURE__ */ consTokens(foldableArray);
var isCharCodePoint = /* @__PURE__ */ (() => ({ fromChar: codePointFromChar, fromCharCode: boundedEnumCodePoint.toEnum }))();
var isCharChar = { fromChar: (x) => x, fromCharCode: charToEnum };
var toModuleName = (v) => {
  if (v === "") {
    return Nothing;
  }
  return $Maybe("Just", take(length2(v) - 1 | 0)(v));
};
var optional = (v) => (str) => {
  const v1 = v(str);
  if (v1.tag === "LexFail") {
    if (length2(str) === length2(v1._2)) {
      return $LexResult("LexSucc", Nothing, str);
    }
    return $LexResult("LexFail", v1._1, v1._2);
  }
  if (v1.tag === "LexSucc") {
    return $LexResult("LexSucc", $Maybe("Just", v1._1), v1._2);
  }
  fail();
};
var mkUnexpected = (str) => {
  const start = take2(6)(str);
  const len = toCodePointArray(start).length;
  if (len === 0) {
    return "end of file";
  }
  if (len < 6) {
    return start;
  }
  return start + "...";
};
var regex2 = (mkErr) => (regexStr) => {
  const matchRegex = unsafeRegex("^(?:" + regexStr + ")")(unicode);
  return (str) => {
    const v = match(matchRegex)(str);
    if (v.tag === "Just") {
      const $0 = (() => {
        if (0 < v._1.length) {
          return v._1[0];
        }
        fail();
      })();
      if ($0.tag === "Just") {
        return $LexResult("LexSucc", $0._1, drop(length2($0._1))(str));
      }
    }
    return $LexResult("LexFail", (v3) => mkErr(mkUnexpected(str)), str);
  };
};
var satisfy = (mkErr) => (p) => (str) => {
  const v = charAt2(0)(str);
  if (v.tag === "Just" && p(v._1)) {
    return $LexResult("LexSucc", v._1, drop(1)(str));
  }
  return $LexResult("LexFail", (v1) => mkErr(mkUnexpected(str)), str);
};
var string = (mkErr) => (match2) => (str) => {
  if (take(length2(match2))(str) === match2) {
    return $LexResult("LexSucc", match2, drop(length2(match2))(str));
  }
  return $LexResult("LexFail", (v) => mkErr(mkUnexpected(str)), str);
};
var many = (v) => (str) => {
  const valuesRef = [];
  let strRef = str;
  let contRef = true;
  let resRef = $LexResult("LexSucc", [], str);
  while (contRef) {
    const str$p = strRef;
    const v1 = v(str$p);
    if (v1.tag === "LexFail") {
      if (length2(str$p) === length2(v1._2)) {
        resRef = $LexResult("LexSucc", valuesRef, v1._2);
        contRef = false;
        continue;
      }
      resRef = $LexResult("LexFail", v1._1, v1._2);
      contRef = false;
      continue;
    }
    if (v1.tag === "LexSucc") {
      valuesRef.push(v1._1);
      strRef = v1._2;
      continue;
    }
    fail();
  }
  return resRef;
};
var functorLex = {
  map: (f) => (v) => (str) => {
    const v1 = v(str);
    if (v1.tag === "LexFail") {
      return $LexResult("LexFail", v1._1, v1._2);
    }
    if (v1.tag === "LexSucc") {
      return $LexResult("LexSucc", f(v1._1), v1._2);
    }
    fail();
  }
};
var spaceComment = /* @__PURE__ */ (() => {
  const $0 = regex2(LexExpected("spaces"))(" +");
  return (str) => {
    const v1 = $0(str);
    if (v1.tag === "LexFail") {
      return $LexResult("LexFail", v1._1, v1._2);
    }
    if (v1.tag === "LexSucc") {
      return $LexResult("LexSucc", length2(v1._1), v1._2);
    }
    fail();
  };
})();
var char$p = (mkErr) => (res) => (match2) => (str) => {
  if (singleton(match2) === take(1)(str)) {
    return $LexResult("LexSucc", res, drop(1)(str));
  }
  return $LexResult("LexFail", (v) => mkErr(mkUnexpected(str)), str);
};
var $$char = (mkErr) => (match2) => (str) => {
  if (singleton(match2) === take(1)(str)) {
    return $LexResult("LexSucc", match2, drop(1)(str));
  }
  return $LexResult("LexFail", (v) => mkErr(mkUnexpected(str)), str);
};
var bumpText = (v) => (colOffset) => (str) => {
  const $0 = v.column;
  const $1 = v.line;
  const go = (go$a0$copy) => (go$a1$copy) => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
    while (go$c) {
      const n = go$a0, ix = go$a1;
      const v1 = indexOf$p("\n")(ix)(str);
      if (v1.tag === "Just") {
        go$a0 = n + 1 | 0;
        go$a1 = v1._1 + 1 | 0;
        continue;
      }
      if (v1.tag === "Nothing") {
        if (n === 0) {
          go$c = false;
          go$r = { line: $1, column: ($0 + toCodePointArray(str).length | 0) + (colOffset * 2 | 0) | 0 };
          continue;
        }
        go$c = false;
        go$r = { line: $1 + n | 0, column: toCodePointArray(drop(ix)(str)).length + colOffset | 0 };
        continue;
      }
      fail();
    }
    return go$r;
  };
  return go(0)(0);
};
var bumpToken = (v) => {
  const $0 = v.column;
  const $1 = v.line;
  return (v1) => {
    if (v1.tag === "TokLeftParen") {
      return { line: $1, column: $0 + 1 | 0 };
    }
    if (v1.tag === "TokRightParen") {
      return { line: $1, column: $0 + 1 | 0 };
    }
    if (v1.tag === "TokLeftBrace") {
      return { line: $1, column: $0 + 1 | 0 };
    }
    if (v1.tag === "TokRightBrace") {
      return { line: $1, column: $0 + 1 | 0 };
    }
    if (v1.tag === "TokLeftSquare") {
      return { line: $1, column: $0 + 1 | 0 };
    }
    if (v1.tag === "TokRightSquare") {
      return { line: $1, column: $0 + 1 | 0 };
    }
    if (v1.tag === "TokLeftArrow") {
      if (v1._1 === "ASCII") {
        return { line: $1, column: $0 + 2 | 0 };
      }
      if (v1._1 === "Unicode") {
        return { line: $1, column: $0 + 1 | 0 };
      }
      fail();
    }
    if (v1.tag === "TokRightArrow") {
      if (v1._1 === "ASCII") {
        return { line: $1, column: $0 + 2 | 0 };
      }
      if (v1._1 === "Unicode") {
        return { line: $1, column: $0 + 1 | 0 };
      }
      fail();
    }
    if (v1.tag === "TokRightFatArrow") {
      if (v1._1 === "ASCII") {
        return { line: $1, column: $0 + 2 | 0 };
      }
      if (v1._1 === "Unicode") {
        return { line: $1, column: $0 + 1 | 0 };
      }
      fail();
    }
    if (v1.tag === "TokDoubleColon") {
      if (v1._1 === "ASCII") {
        return { line: $1, column: $0 + 2 | 0 };
      }
      if (v1._1 === "Unicode") {
        return { line: $1, column: $0 + 1 | 0 };
      }
      fail();
    }
    if (v1.tag === "TokForall") {
      if (v1._1 === "ASCII") {
        return { line: $1, column: $0 + 6 | 0 };
      }
      if (v1._1 === "Unicode") {
        return { line: $1, column: $0 + 1 | 0 };
      }
      fail();
    }
    if (v1.tag === "TokEquals") {
      return { line: $1, column: $0 + 1 | 0 };
    }
    if (v1.tag === "TokPipe") {
      return { line: $1, column: $0 + 1 | 0 };
    }
    if (v1.tag === "TokTick") {
      return { line: $1, column: $0 + 1 | 0 };
    }
    if (v1.tag === "TokDot") {
      return { line: $1, column: $0 + 1 | 0 };
    }
    if (v1.tag === "TokComma") {
      return { line: $1, column: $0 + 1 | 0 };
    }
    if (v1.tag === "TokUnderscore") {
      return { line: $1, column: $0 + 1 | 0 };
    }
    if (v1.tag === "TokBackslash") {
      return { line: $1, column: $0 + 1 | 0 };
    }
    if (v1.tag === "TokAt") {
      return { line: $1, column: $0 + 1 | 0 };
    }
    if (v1.tag === "TokLowerName") {
      return {
        line: $1,
        column: (() => {
          if (v1._1.tag === "Nothing") {
            return $0 + 0 | 0;
          }
          if (v1._1.tag === "Just") {
            return ($0 + 1 | 0) + toCodePointArray(v1._1._1).length | 0;
          }
          fail();
        })() + toCodePointArray(v1._2).length | 0
      };
    }
    if (v1.tag === "TokUpperName") {
      return {
        line: $1,
        column: (() => {
          if (v1._1.tag === "Nothing") {
            return $0 + 0 | 0;
          }
          if (v1._1.tag === "Just") {
            return ($0 + 1 | 0) + toCodePointArray(v1._1._1).length | 0;
          }
          fail();
        })() + toCodePointArray(v1._2).length | 0
      };
    }
    if (v1.tag === "TokOperator") {
      return {
        line: $1,
        column: (() => {
          if (v1._1.tag === "Nothing") {
            return $0 + 0 | 0;
          }
          if (v1._1.tag === "Just") {
            return ($0 + 1 | 0) + toCodePointArray(v1._1._1).length | 0;
          }
          fail();
        })() + toCodePointArray(v1._2).length | 0
      };
    }
    if (v1.tag === "TokSymbolName") {
      return {
        line: $1,
        column: ((() => {
          if (v1._1.tag === "Nothing") {
            return $0 + 0 | 0;
          }
          if (v1._1.tag === "Just") {
            return ($0 + 1 | 0) + toCodePointArray(v1._1._1).length | 0;
          }
          fail();
        })() + toCodePointArray(v1._2).length | 0) + 2 | 0
      };
    }
    if (v1.tag === "TokSymbolArrow") {
      if (v1._1 === "Unicode") {
        return { line: $1, column: $0 + 3 | 0 };
      }
      if (v1._1 === "ASCII") {
        return { line: $1, column: $0 + 4 | 0 };
      }
      fail();
    }
    if (v1.tag === "TokHole") {
      return { line: $1, column: ($0 + toCodePointArray(v1._1).length | 0) + 1 | 0 };
    }
    if (v1.tag === "TokChar") {
      return { line: $1, column: ($0 + toCodePointArray(v1._1).length | 0) + 2 | 0 };
    }
    if (v1.tag === "TokInt") {
      return { line: $1, column: $0 + toCodePointArray(v1._1).length | 0 };
    }
    if (v1.tag === "TokNumber") {
      return { line: $1, column: $0 + toCodePointArray(v1._1).length | 0 };
    }
    if (v1.tag === "TokString") {
      return bumpText(v)(1)(v1._1);
    }
    if (v1.tag === "TokRawString") {
      return bumpText(v)(3)(v1._1);
    }
    if (v1.tag === "TokLayoutStart") {
      return v;
    }
    if (v1.tag === "TokLayoutSep") {
      return v;
    }
    if (v1.tag === "TokLayoutEnd") {
      return v;
    }
    fail();
  };
};
var bumpComment = (v) => {
  const $0 = v.column;
  const $1 = v.line;
  return (v1) => {
    if (v1.tag === "Comment") {
      return bumpText(v)(0)(v1._1);
    }
    if (v1.tag === "Space") {
      return { line: $1, column: $0 + v1._1 | 0 };
    }
    if (v1.tag === "Line") {
      return { line: $1 + v1._2 | 0, column: 0 };
    }
    fail();
  };
};
var altLex = {
  alt: (v) => (v1) => (str) => {
    const v2 = v(str);
    if (v2.tag === "LexFail") {
      if (length2(str) === length2(v2._2)) {
        return v1(str);
      }
      return $LexResult("LexFail", v2._1, v2._2);
    }
    if (v2.tag === "LexSucc") {
      return $LexResult("LexSucc", v2._1, v2._2);
    }
    fail();
  },
  Functor0: () => functorLex
};
var comment = /* @__PURE__ */ (() => altLex.alt(regex2(LexExpected("block comment"))("\\{-(-(?!\\})|[^-]+)*(-\\}|$)"))(regex2(LexExpected("line comment"))("--[^\\r\\n]*")))();
var lineComment = /* @__PURE__ */ (() => altLex.alt((() => {
  const $0 = Line(LF);
  const $1 = regex2(LexExpected("newline"))("\n+");
  return (str) => {
    const v1 = $1(str);
    if (v1.tag === "LexFail") {
      return $LexResult("LexFail", v1._1, v1._2);
    }
    if (v1.tag === "LexSucc") {
      return $LexResult("LexSucc", $0(toCodePointArray(v1._1).length), v1._2);
    }
    fail();
  };
})())((() => {
  const $0 = Line(CRLF);
  const $1 = regex2(LexExpected("newline"))("(?:\r\n)+");
  return (str) => {
    const v1 = $1(str);
    if (v1.tag === "LexFail") {
      return $LexResult("LexFail", v1._1, v1._2);
    }
    if (v1.tag === "LexSucc") {
      return $LexResult("LexSucc", $0(intDiv(toCodePointArray(v1._1).length, 2)), v1._2);
    }
    fail();
  };
})()))();
var leadingComments = /* @__PURE__ */ (() => many(altLex.alt((str) => {
  const v1 = comment(str);
  if (v1.tag === "LexFail") {
    return $LexResult("LexFail", v1._1, v1._2);
  }
  if (v1.tag === "LexSucc") {
    return $LexResult("LexSucc", $Comment2("Comment", v1._1), v1._2);
  }
  fail();
})(altLex.alt((str) => {
  const v1 = spaceComment(str);
  if (v1.tag === "LexFail") {
    return $LexResult("LexFail", v1._1, v1._2);
  }
  if (v1.tag === "LexSucc") {
    return $LexResult("LexSucc", $Comment2("Space", v1._1), v1._2);
  }
  fail();
})(lineComment))))();
var token = /* @__PURE__ */ (() => {
  const tokenRightParen = char$p(LexExpected("right paren"))(TokRightParen)(")");
  const tokenLeftParen = char$p(LexExpected("left paren"))(TokLeftParen)("(");
  const stripUnderscores = replaceAll("_")("");
  const parseSymbolIdent = regex2(LexExpected("symbol"))("(?:[:!#$%&*+./<=>?@\\\\^|~-]|(?!\\p{P})\\p{S})+");
  const parseProper = regex2(LexExpected("proper name"))("\\p{Lu}[\\p{L}0-9_']*");
  const parseIdent = regex2(LexExpected("ident"))("[\\p{Ll}_][\\p{L}0-9_']*");
  const intPartRegex = regex2(LexExpected("int part"))("(0|[1-9][0-9_]*)");
  const hexEscapeRegex = regex2(LexExpected("hex"))("[a-fA-F0-9]{1,6}");
  const charSingleQuote = $$char(LexExpected("single quote"))("'");
  const charQuote = $$char(LexExpected("quote"))('"');
  const charAny = satisfy(LexExpected("char"))((v) => true);
  const parseEscape = (dictIsChar) => (str) => {
    const v1 = charAny(str);
    if (v1.tag === "LexFail") {
      return $LexResult("LexFail", v1._1, v1._2);
    }
    if (v1.tag === "LexSucc") {
      if (v1._1 === "t") {
        return $LexResult("LexSucc", { raw: "\\t", char: dictIsChar.fromChar("	") }, v1._2);
      }
      if (v1._1 === "r") {
        return $LexResult("LexSucc", { raw: "\\r", char: dictIsChar.fromChar("\r") }, v1._2);
      }
      if (v1._1 === "n") {
        return $LexResult("LexSucc", { raw: "\\n", char: dictIsChar.fromChar("\n") }, v1._2);
      }
      if (v1._1 === '"') {
        return $LexResult("LexSucc", { raw: '\\"', char: dictIsChar.fromChar('"') }, v1._2);
      }
      if (v1._1 === "'") {
        return $LexResult("LexSucc", { raw: "\\'", char: dictIsChar.fromChar("'") }, v1._2);
      }
      if (v1._1 === "\\") {
        return $LexResult("LexSucc", { raw: "\\\\", char: dictIsChar.fromChar("\\") }, v1._2);
      }
      if (v1._1 === "x") {
        const v1$1 = hexEscapeRegex(v1._2);
        if (v1$1.tag === "LexFail") {
          return $LexResult("LexFail", v1$1._1, v1$1._2);
        }
        if (v1$1.tag === "LexSucc") {
          const $02 = fromStringAs(16)(v1$1._1);
          const v = (() => {
            if ($02.tag === "Just") {
              return dictIsChar.fromCharCode($02._1);
            }
            if ($02.tag === "Nothing") {
              return Nothing;
            }
            fail();
          })();
          if (v.tag === "Just") {
            return $LexResult("LexSucc", { raw: "\\x" + v1$1._1, char: v._1 }, v1$1._2);
          }
          if (v.tag === "Nothing") {
            return $LexResult("LexFail", (v$1) => $ParseError("LexCharEscapeOutOfRange", v1$1._1), v1$1._2);
          }
        }
        fail();
      }
      const $0 = $ParseError("LexInvalidCharEscape", singleton(v1._1));
      return $LexResult("LexFail", (v) => $0, v1._2);
    }
    fail();
  };
  const parseEscape1 = parseEscape(isCharChar);
  return altLex.alt((() => {
    const $0 = $$char(LexExpected("question mark"))("?");
    return (str) => {
      const v1 = $0(str);
      if (v1.tag === "LexFail") {
        return $LexResult("LexFail", v1._1, str);
      }
      if (v1.tag === "LexSucc") {
        const v3 = altLex.alt(parseIdent)(parseProper)(v1._2);
        if (v3.tag === "LexFail") {
          return $LexResult("LexFail", v3._1, str);
        }
        if (v3.tag === "LexSucc") {
          return $LexResult("LexSucc", $Token("TokHole", v3._1), v3._2);
        }
      }
      fail();
    };
  })())(altLex.alt((() => {
    const $0 = regex2(LexExpected("module name"))("(?:(?:\\p{Lu}[\\p{L}0-9_']*)\\.)*");
    const $1 = altLex.alt((str) => {
      const v1 = parseIdent(str);
      if (v1.tag === "LexFail") {
        return $LexResult("LexFail", v1._1, v1._2);
      }
      if (v1.tag === "LexSucc") {
        return $LexResult(
          "LexSucc",
          (() => {
            const $12 = v1._1;
            return (v1$1) => {
              if (v1$1.tag === "Nothing") {
                if ($12 === "forall") {
                  return $Token("TokForall", ASCII);
                }
                if ($12 === "_") {
                  return TokUnderscore;
                }
                return $Token("TokLowerName", Nothing, $12);
              }
              return $Token("TokLowerName", v1$1, $12);
            };
          })(),
          v1._2
        );
      }
      fail();
    })(altLex.alt((str) => {
      const v1 = parseProper(str);
      if (v1.tag === "LexFail") {
        return $LexResult("LexFail", v1._1, v1._2);
      }
      if (v1.tag === "LexSucc") {
        return $LexResult(
          "LexSucc",
          (() => {
            const $12 = v1._1;
            return (a) => $Token("TokUpperName", a, $12);
          })(),
          v1._2
        );
      }
      fail();
    })(altLex.alt((str) => {
      const v1 = parseSymbolIdent(str);
      if (v1.tag === "LexFail") {
        return $LexResult("LexFail", v1._1, v1._2);
      }
      if (v1.tag === "LexSucc") {
        return $LexResult(
          "LexSucc",
          (() => {
            const $12 = v1._1;
            return (v1$1) => {
              if (v1$1.tag === "Nothing") {
                if ($12 === "<-") {
                  return $Token("TokLeftArrow", ASCII);
                }
                if ($12 === "\u2190") {
                  return $Token("TokLeftArrow", Unicode);
                }
                if ($12 === "->") {
                  return $Token("TokRightArrow", ASCII);
                }
                if ($12 === "\u2192") {
                  return $Token("TokRightArrow", Unicode);
                }
                if ($12 === "=>") {
                  return $Token("TokRightFatArrow", ASCII);
                }
                if ($12 === "\u21D2") {
                  return $Token("TokRightFatArrow", Unicode);
                }
                if ($12 === "::") {
                  return $Token("TokDoubleColon", ASCII);
                }
                if ($12 === "\u2237") {
                  return $Token("TokDoubleColon", Unicode);
                }
                if ($12 === "\u2200") {
                  return $Token("TokForall", Unicode);
                }
                if ($12 === "=") {
                  return TokEquals;
                }
                if ($12 === ".") {
                  return TokDot;
                }
                if ($12 === "\\") {
                  return TokBackslash;
                }
                if ($12 === "|") {
                  return TokPipe;
                }
                if ($12 === "@") {
                  return TokAt;
                }
                if ($12 === "`") {
                  return TokTick;
                }
                return $Token("TokOperator", Nothing, $12);
              }
              return $Token("TokOperator", v1$1, $12);
            };
          })(),
          v1._2
        );
      }
      fail();
    })((str) => {
      const v1 = tokenLeftParen(str);
      if (v1.tag === "LexFail") {
        return $LexResult("LexFail", v1._1, str);
      }
      if (v1.tag === "LexSucc") {
        const v3 = parseSymbolIdent(v1._2);
        if (v3.tag === "LexFail") {
          return $LexResult("LexFail", v3._1, str);
        }
        if (v3.tag === "LexSucc") {
          const v3$1 = tokenRightParen(v3._2);
          if (v3$1.tag === "LexFail") {
            return $LexResult("LexFail", v3$1._1, str);
          }
          if (v3$1.tag === "LexSucc") {
            return $LexResult(
              "LexSucc",
              (() => {
                const $12 = v3._1;
                return (v1$1) => {
                  if (v1$1.tag === "Nothing") {
                    if ($12 === "->") {
                      return $Token("TokSymbolArrow", ASCII);
                    }
                    if ($12 === "\u2192") {
                      return $Token("TokSymbolArrow", Unicode);
                    }
                    return $Token("TokSymbolName", Nothing, $12);
                  }
                  return $Token("TokSymbolName", v1$1, $12);
                };
              })(),
              v3$1._2
            );
          }
        }
      }
      fail();
    })));
    return (str) => {
      const v1 = $0(str);
      if (v1.tag === "LexFail") {
        return $LexResult("LexFail", v1._1, v1._2);
      }
      if (v1.tag === "LexSucc") {
        const v3 = $1(v1._2);
        if (v3.tag === "LexFail") {
          return $LexResult("LexFail", v3._1, v3._2);
        }
        if (v3.tag === "LexSucc") {
          return $LexResult("LexSucc", v3._1(toModuleName(v1._1)), v3._2);
        }
      }
      fail();
    };
  })())(altLex.alt((str) => {
    const v1 = charSingleQuote(str);
    if (v1.tag === "LexFail") {
      return $LexResult("LexFail", v1._1, v1._2);
    }
    if (v1.tag === "LexSucc") {
      const v1$1 = charAny(v1._2);
      const v3 = (() => {
        if (v1$1.tag === "LexFail") {
          return $LexResult("LexFail", v1$1._1, v1$1._2);
        }
        if (v1$1.tag === "LexSucc") {
          if (v1$1._1 === "\\") {
            return parseEscape1(v1$1._2);
          }
          if (v1$1._1 === "'") {
            return $LexResult("LexFail", (v) => $ParseError("LexExpected", "character", "empty character literal"), v1$1._2);
          }
          return $LexResult("LexSucc", { raw: singleton(v1$1._1), char: v1$1._1 }, v1$1._2);
        }
        fail();
      })();
      if (v3.tag === "LexFail") {
        return $LexResult("LexFail", v3._1, v3._2);
      }
      if (v3.tag === "LexSucc") {
        const v3$1 = charSingleQuote(v3._2);
        if (v3$1.tag === "LexFail") {
          return $LexResult("LexFail", v3$1._1, v3$1._2);
        }
        if (v3$1.tag === "LexSucc") {
          return $LexResult("LexSucc", $Token("TokChar", v3._1.raw, v3._1.char), v3$1._2);
        }
      }
    }
    fail();
  })(altLex.alt(altLex.alt((() => {
    const $0 = regex2(LexExpected("raw string characters"))('""""{0,2}([^"]+"{1,2})*[^"]*"""');
    return (str) => {
      const v1 = $0(str);
      if (v1.tag === "LexFail") {
        return $LexResult("LexFail", v1._1, v1._2);
      }
      if (v1.tag === "LexSucc") {
        return $LexResult(
          "LexSucc",
          $Token(
            "TokRawString",
            (() => {
              const $1 = drop(3)(v1._1);
              return take(length2($1) - 3 | 0)($1);
            })()
          ),
          v1._2
        );
      }
      fail();
    };
  })())((() => {
    const $0 = many(altLex.alt((() => {
      const $02 = regex2(LexExpected("string characters"))('[^"\\\\]+');
      return (str) => {
        const v1 = $02(str);
        if (v1.tag === "LexFail") {
          return $LexResult("LexFail", v1._1, v1._2);
        }
        if (v1.tag === "LexSucc") {
          return $LexResult("LexSucc", { raw: v1._1, string: v1._1 }, v1._2);
        }
        fail();
      };
    })())(altLex.alt((() => {
      const $02 = regex2(LexExpected("whitespace escape"))("\\\\[ \\r\\n]+\\\\");
      return (str) => {
        const v1 = $02(str);
        if (v1.tag === "LexFail") {
          return $LexResult("LexFail", v1._1, v1._2);
        }
        if (v1.tag === "LexSucc") {
          return $LexResult("LexSucc", { raw: v1._1, string: "" }, v1._2);
        }
        fail();
      };
    })())((() => {
      const $02 = $$char(LexExpected("backslash"))("\\");
      const $1 = parseEscape(isCharCodePoint);
      return (str) => {
        const v1 = $02(str);
        if (v1.tag === "LexFail") {
          return $LexResult("LexFail", v1._1, v1._2);
        }
        if (v1.tag === "LexSucc") {
          const v3 = $1(v1._2);
          if (v3.tag === "LexFail") {
            return $LexResult("LexFail", v3._1, v3._2);
          }
          if (v3.tag === "LexSucc") {
            return $LexResult("LexSucc", { raw: v3._1.raw, string: singleton2(v3._1.char) }, v3._2);
          }
        }
        fail();
      };
    })())));
    return (str) => {
      const v1 = charQuote(str);
      if (v1.tag === "LexFail") {
        return $LexResult("LexFail", v1._1, v1._2);
      }
      if (v1.tag === "LexSucc") {
        const v3 = $0(v1._2);
        if (v3.tag === "LexFail") {
          return $LexResult("LexFail", v3._1, v3._2);
        }
        if (v3.tag === "LexSucc") {
          const v3$1 = charQuote(v3._2);
          if (v3$1.tag === "LexFail") {
            return $LexResult("LexFail", v3$1._1, v3$1._2);
          }
          if (v3$1.tag === "LexSucc") {
            return $LexResult(
              "LexSucc",
              (() => {
                const v1$1 = fold1(v3._1);
                return $Token("TokString", v1$1.raw, v1$1.string);
              })(),
              v3$1._2
            );
          }
        }
      }
      fail();
    };
  })()))(altLex.alt(altLex.alt((() => {
    const $0 = string(LexExpected("hex int prefix"))("0x");
    const $1 = regex2(LexExpected("hex int"))("[a-fA-F0-9]+");
    return (str) => {
      const v1 = $0(str);
      if (v1.tag === "LexFail") {
        return $LexResult("LexFail", v1._1, v1._2);
      }
      if (v1.tag === "LexSucc") {
        const v3 = $1(v1._2);
        if (v3.tag === "LexFail") {
          return $LexResult("LexFail", v3._1, v3._2);
        }
        if (v3.tag === "LexSucc") {
          const v = fromStringAs(16)(v3._1);
          if (v.tag === "Just") {
            return $LexResult("LexSucc", $Token("TokInt", "0x" + v3._1, $IntValue("SmallInt", v._1)), v3._2);
          }
          if (v.tag === "Nothing") {
            return $LexResult("LexSucc", $Token("TokInt", "0x" + v3._1, $IntValue("BigHex", v3._1)), v3._2);
          }
        }
      }
      fail();
    };
  })())((str) => {
    const v1 = intPartRegex(str);
    if (v1.tag === "LexFail") {
      return $LexResult("LexFail", v1._1, v1._2);
    }
    if (v1.tag === "LexSucc") {
      const v1$1 = optional((() => {
        const $0 = $$char(LexExpected("dot"))(".");
        const $1 = regex2(LexExpected("fraction part"))("[0-9_]+");
        return (str$1) => {
          const v1$12 = $0(str$1);
          if (v1$12.tag === "LexFail") {
            return $LexResult("LexFail", v1$12._1, str$1);
          }
          if (v1$12.tag === "LexSucc") {
            const v3 = $1(v1$12._2);
            if (v3.tag === "LexFail") {
              return $LexResult("LexFail", v3._1, str$1);
            }
            if (v3.tag === "LexSucc") {
              return $LexResult("LexSucc", v3._1, v3._2);
            }
          }
          fail();
        };
      })())(v1._2);
      if (v1$1.tag === "LexFail") {
        return $LexResult("LexFail", v1$1._1, v1$1._2);
      }
      if (v1$1.tag === "LexSucc") {
        const v1$2 = optional((() => {
          const $0 = $$char(LexExpected("exponent"))("e");
          const $1 = optional(altLex.alt(string(LexExpected("negative"))("-"))(string(LexExpected("positive"))("+")));
          return (str$1) => {
            const v1$22 = $0(str$1);
            if (v1$22.tag === "LexFail") {
              return $LexResult("LexFail", v1$22._1, v1$22._2);
            }
            if (v1$22.tag === "LexSucc") {
              const v1$3 = $1(v1$22._2);
              if (v1$3.tag === "LexFail") {
                return $LexResult("LexFail", v1$3._1, v1$3._2);
              }
              if (v1$3.tag === "LexSucc") {
                const v3 = intPartRegex(v1$3._2);
                if (v3.tag === "LexFail") {
                  return $LexResult("LexFail", v3._1, v3._2);
                }
                if (v3.tag === "LexSucc") {
                  return $LexResult("LexSucc", { sign: v1$3._1, exponent: v3._1 }, v3._2);
                }
              }
            }
            fail();
          };
        })())(v1$1._2);
        if (v1$2.tag === "LexFail") {
          return $LexResult("LexFail", v1$2._1, v1$2._2);
        }
        if (v1$2.tag === "LexSucc") {
          if ((() => {
            if (v1$1._1.tag === "Nothing") {
              return true;
            }
            if (v1$1._1.tag === "Just") {
              return false;
            }
            fail();
          })() && (() => {
            if (v1$2._1.tag === "Nothing") {
              return true;
            }
            if (v1$2._1.tag === "Just") {
              return false;
            }
            fail();
          })()) {
            const intVal = stripUnderscores(v1._1);
            const v2 = fromString(intVal);
            if (v2.tag === "Just") {
              return $LexResult("LexSucc", $Token("TokInt", v1._1, $IntValue("SmallInt", v2._1)), v1$2._2);
            }
            if (v2.tag === "Nothing") {
              return $LexResult("LexSucc", $Token("TokInt", v1._1, $IntValue("BigInt", intVal)), v1$2._2);
            }
            fail();
          }
          const raw = (() => {
            if (v1$1._1.tag === "Nothing") {
              return v1._1 + "";
            }
            if (v1$1._1.tag === "Just") {
              return v1._1 + "." + v1$1._1._1;
            }
            fail();
          })() + (() => {
            if (v1$2._1.tag === "Nothing") {
              return "";
            }
            if (v1$2._1.tag === "Just") {
              if (v1$2._1._1.sign.tag === "Nothing") {
                return "e" + v1$2._1._1.exponent;
              }
              if (v1$2._1._1.sign.tag === "Just") {
                return "e" + v1$2._1._1.sign._1 + v1$2._1._1.exponent;
              }
            }
            fail();
          })();
          const v = fromStringImpl(stripUnderscores(raw), isFiniteImpl, Just, Nothing);
          if (v.tag === "Just") {
            return $LexResult("LexSucc", $Token("TokNumber", raw, v._1), v1$2._2);
          }
          if (v.tag === "Nothing") {
            return $LexResult("LexFail", (v$1) => $ParseError("LexNumberOutOfRange", raw), v1$2._2);
          }
        }
      }
    }
    fail();
  }))(altLex.alt(tokenLeftParen)(altLex.alt(tokenRightParen)(altLex.alt(char$p(LexExpected("left brace"))(TokLeftBrace)("{"))(altLex.alt(char$p(LexExpected("right brace"))(TokRightBrace)("}"))(altLex.alt(char$p(LexExpected("left square"))(TokLeftSquare)("["))(altLex.alt(char$p(LexExpected("right square"))(TokRightSquare)("]"))(altLex.alt(char$p(LexExpected("backtick"))(TokTick)("`"))(char$p(LexExpected("comma"))(TokComma)(",")))))))))))));
})();
var trailingComments = /* @__PURE__ */ (() => many(altLex.alt((str) => {
  const v1 = comment(str);
  if (v1.tag === "LexFail") {
    return $LexResult("LexFail", v1._1, v1._2);
  }
  if (v1.tag === "LexSucc") {
    return $LexResult("LexSucc", $Comment2("Comment", v1._1), v1._2);
  }
  fail();
})((str) => {
  const v1 = spaceComment(str);
  if (v1.tag === "LexFail") {
    return $LexResult("LexFail", v1._1, v1._2);
  }
  if (v1.tag === "LexSucc") {
    return $LexResult("LexSucc", $Comment2("Space", v1._1), v1._2);
  }
  fail();
})))();
var lexWithState$p = (lexLeadingComments) => {
  const go = (stack) => (startPos) => (leading) => (str) => defer((v) => {
    if (str === "") {
      return force(unwindLayout(startPos)(defer((v12) => $TokenStep("TokenEOF", startPos, leading)))(stack));
    }
    const v1 = token(str);
    if (v1.tag === "LexFail") {
      return $TokenStep(
        "TokenError",
        bumpText(startPos)(0)(take(length2(str) - length2(v1._2) | 0)(str)),
        v1._1(),
        Nothing,
        stack
      );
    }
    if (v1.tag === "LexSucc") {
      const v3 = trailingComments(v1._2);
      if (v3.tag === "LexFail") {
        return $TokenStep(
          "TokenError",
          bumpText(startPos)(0)(take(length2(str) - length2(v3._2) | 0)(str)),
          v3._1(),
          Nothing,
          stack
        );
      }
      if (v3.tag === "LexSucc") {
        const v3$1 = leadingComments(v3._2);
        if (v3$1.tag === "LexFail") {
          return $TokenStep(
            "TokenError",
            bumpText(startPos)(0)(take(length2(str) - length2(v3$1._2) | 0)(str)),
            v3$1._1(),
            Nothing,
            stack
          );
        }
        if (v3$1.tag === "LexSucc") {
          const endPos = bumpToken(startPos)(v1._1);
          const nextStart = foldlArray(bumpComment)(foldlArray(bumpComment)(endPos)(v3._1))(v3$1._1);
          const v2 = insertLayout({ range: { start: startPos, end: endPos }, leadingComments: leading, trailingComments: v3._1, value: v1._1 })(nextStart)(stack);
          return force(consTokens2(v2._2)($Tuple(nextStart, go(v2._1)(nextStart)(v3$1._1)(v3$1._2)))._2);
        }
      }
    }
    fail();
  });
  return (initStack) => (initPos) => (str) => defer((v) => {
    const v1 = lexLeadingComments(str);
    if (v1.tag === "LexFail") {
      return _crashWith("Leading comments can't fail.");
    }
    if (v1.tag === "LexSucc") {
      return force(go(initStack)(foldlArray(bumpComment)(initPos)(v1._1))(v1._1)(v1._2));
    }
    fail();
  });
};
var lexWithState = /* @__PURE__ */ lexWithState$p(leadingComments);
var lex = /* @__PURE__ */ lexWithState(/* @__PURE__ */ $List(
  "Cons",
  /* @__PURE__ */ $Tuple({ line: 0, column: 0 }, LytRoot),
  Nil
))({ line: 0, column: 0 });

// output-es/PureScript.CST.Parser.Monad/index.js
var $ParserResult = (tag, _1, _2) => ({ tag, _1, _2 });
var $Trampoline = (tag, _1) => ({ tag, _1 });
var More = (value0) => $Trampoline("More", value0);
var take3 = (k) => (state, v, resume, done) => {
  const v1 = force(state.stream);
  if (v1.tag === "TokenError") {
    return resume(state, { error: v1._2, position: v1._1 });
  }
  if (v1.tag === "TokenEOF") {
    return resume(state, { error: UnexpectedEof, position: v1._1 });
  }
  if (v1.tag === "TokenCons") {
    const v2 = k(v1._1);
    if (v2.tag === "Left") {
      return resume(state, { error: v2._1, position: v1._1.range.start });
    }
    if (v2.tag === "Right") {
      return done({ ...state, consumed: true, stream: v1._3 }, v2._1);
    }
  }
  fail();
};
var runParser$p = (state1) => (v) => {
  const run2 = (run$a0$copy) => {
    let run$a0 = run$a0$copy, run$c = true, run$r;
    while (run$c) {
      const v1 = run$a0;
      if (v1.tag === "More") {
        run$a0 = v1._1();
        continue;
      }
      if (v1.tag === "Done") {
        run$c = false;
        run$r = v1._1;
        continue;
      }
      fail();
    }
    return run$r;
  };
  return run2(v(
    state1,
    More,
    (state2, error3) => $Trampoline("Done", $ParserResult("ParseFail", error3, state2)),
    (state2, value) => $Trampoline("Done", $ParserResult("ParseSucc", value, state2))
  ));
};
var runParser = (stream) => (x) => {
  const $0 = runParser$p({ consumed: false, errors: [], stream })(x);
  if ($0.tag === "ParseFail") {
    return $Either("Left", $0._1);
  }
  if ($0.tag === "ParseSucc") {
    return $Either("Right", $Tuple($0._1, $0._2.errors));
  }
  fail();
};
var eof = (state, v, resume, done) => {
  const v1 = force(state.stream);
  if (v1.tag === "TokenError") {
    return resume(state, { error: v1._2, position: v1._1 });
  }
  if (v1.tag === "TokenEOF") {
    return done({ ...state, consumed: true }, $Tuple(v1._1, v1._2));
  }
  if (v1.tag === "TokenCons") {
    return resume(state, { error: $ParseError("ExpectedEof", v1._1.value), position: v1._1.range.start });
  }
  fail();
};

// output-es/PureScript.Backend.Optimizer.Directives/index.js
var expectMap = (k) => take3((tok) => {
  const v = k(tok);
  if (v.tag === "Just") {
    return $Either("Right", v._1);
  }
  if (v.tag === "Nothing") {
    return $Either("Left", $ParseError("UnexpectedToken", tok.value));
  }
  fail();
});
var keyword = (word1) => expectMap((v) => {
  if (v.value.tag === "TokLowerName" && v.value._1.tag === "Nothing" && word1 === v.value._2) {
    return $Maybe("Just", void 0);
  }
  return Nothing;
});
var label = /* @__PURE__ */ expectMap((v) => {
  if (v.value.tag === "TokRawString") {
    return $Maybe("Just", v.value._1);
  }
  if (v.value.tag === "TokString") {
    return $Maybe("Just", v.value._2);
  }
  if (v.value.tag === "TokLowerName" && v.value._1.tag === "Nothing") {
    return $Maybe("Just", v.value._2);
  }
  return Nothing;
});
var natural = /* @__PURE__ */ expectMap((v) => {
  if (v.value.tag === "TokInt" && v.value._2.tag === "SmallInt" && v.value._2._1 > 0) {
    return $Maybe("Just", v.value._2._1);
  }
  return Nothing;
});
var qualified = /* @__PURE__ */ expectMap((v) => {
  if (v.value.tag === "TokLowerName" && v.value._1.tag === "Just") {
    return $Maybe("Just", $Qualified($Maybe("Just", v.value._1._1), v.value._2));
  }
  return Nothing;
});
var unqualified = /* @__PURE__ */ expectMap((v) => {
  if (v.value.tag === "TokLowerName" && v.value._1.tag === "Nothing") {
    return $Maybe("Just", v.value._2);
  }
  return Nothing;
});
var equals = /* @__PURE__ */ expectMap((v) => {
  if (v.value.tag === "TokEquals") {
    return $Maybe("Just", void 0);
  }
  return Nothing;
});
var parseInlineDirective = (state1, more, resume, done) => keyword("default")(
  state1.consumed ? { ...state1, consumed: false } : state1,
  more,
  (state3, error3) => {
    if (state3.consumed) {
      return resume(state3, error3);
    }
    return keyword("never")(
      state1.consumed ? { ...state1, consumed: false } : state1,
      more,
      (state3$1, error$1) => {
        if (state3$1.consumed) {
          return resume(state3$1, error$1);
        }
        return keyword("always")(
          state1.consumed ? { ...state1, consumed: false } : state1,
          more,
          (state3$2, error$2) => {
            if (state3$2.consumed) {
              return resume(state3$2, error$2);
            }
            return keyword("arity")(
              state1,
              more,
              resume,
              (state2, a) => more((v2) => equals(
                state2,
                more,
                resume,
                (state3$3, a$1) => more((v2$1) => natural(
                  state3$3,
                  more,
                  resume,
                  (state3$4, a$2) => done(state3$4, $InlineDirective("InlineArity", a$2))
                ))
              ))
            );
          },
          (state2, a) => done(state2, InlineAlways)
        );
      },
      (state2, a) => done(state2, InlineNever)
    );
  },
  (state2, a) => done(state2, InlineDefault)
);
var dotDot = /* @__PURE__ */ expectMap((v) => {
  if (v.value.tag === "TokSymbolName" && v.value._1.tag === "Nothing" && v.value._2 === "..") {
    return $Maybe("Just", void 0);
  }
  return Nothing;
});
var dot = /* @__PURE__ */ expectMap((v) => {
  if (v.value.tag === "TokDot") {
    return $Maybe("Just", void 0);
  }
  return Nothing;
});
var parseInlineAccessor = (state1, more, resume, done) => dot(
  state1.consumed ? { ...state1, consumed: false } : state1,
  more,
  (state3, error3) => {
    if (state3.consumed) {
      return resume(state3, error3);
    }
    return dotDot(
      state1.consumed ? { ...state1, consumed: false } : state1,
      more,
      (state3$1, error$1) => {
        if (state3$1.consumed) {
          return resume(state3$1, error$1);
        }
        return done(state1, InlineRef);
      },
      (state2, a) => more((v2) => dot(
        state2,
        more,
        (state3$1, error$1) => {
          if (state3$1.consumed) {
            return resume(state3$1, error$1);
          }
          return done(state1, InlineRef);
        },
        (state3$1, a$1) => more((v2$1) => label(
          state3$1,
          more,
          (state3$2, error$1) => {
            if (state3$2.consumed) {
              return resume(state3$2, error$1);
            }
            return done(state1, InlineRef);
          },
          (state3$2, a$2) => done(state3$2, $InlineAccessor("InlineSpineProp", a$2))
        ))
      ))
    );
  },
  (state2, a) => more((v2) => label(
    state2,
    more,
    (state3, error3) => {
      if (state3.consumed) {
        return resume(state3, error3);
      }
      return dotDot(
        state1.consumed ? { ...state1, consumed: false } : state1,
        more,
        (state3$1, error$1) => {
          if (state3$1.consumed) {
            return resume(state3$1, error$1);
          }
          return done(state1, InlineRef);
        },
        (state2$1, a$1) => more((v2$1) => dot(
          state2$1,
          more,
          (state3$1, error$1) => {
            if (state3$1.consumed) {
              return resume(state3$1, error$1);
            }
            return done(state1, InlineRef);
          },
          (state3$1, a$2) => more((v2$2) => label(
            state3$1,
            more,
            (state3$2, error$1) => {
              if (state3$2.consumed) {
                return resume(state3$2, error$1);
              }
              return done(state1, InlineRef);
            },
            (state3$2, a$3) => done(state3$2, $InlineAccessor("InlineSpineProp", a$3))
          ))
        ))
      );
    },
    (state3, a$1) => done(state3, $InlineAccessor("InlineProp", a$1))
  ))
);
var parseDirective = (state1, more, resume, done) => qualified(
  state1,
  more,
  resume,
  (state2, a) => more((v2) => parseInlineAccessor(
    state2,
    more,
    resume,
    (state3, a$1) => more((v2$1) => parseInlineDirective(
      state3,
      more,
      resume,
      (state3$1, a$2) => more((v2$2) => eof(
        state3$1,
        more,
        resume,
        (state3$2, a$3) => done(state3$2, $Tuple($EvalRef("EvalExtern", a), $Tuple(a$1, a$2)))
      ))
    ))
  ))
);
var parseDirectiveExport = (moduleName) => (state1, more, resume, done) => keyword("export")(
  state1,
  more,
  resume,
  (state2, a) => more((v2) => unqualified(
    state2,
    more,
    resume,
    (state3, a$1) => more((v2$1) => parseInlineAccessor(
      state3,
      more,
      resume,
      (state3$1, a$2) => more((v2$2) => parseInlineDirective(
        state3$1,
        more,
        resume,
        (state3$2, a$3) => more((v2$3) => eof(
          state3$2,
          more,
          resume,
          (state3$3, a$4) => done(
            state3$3,
            $Tuple(
              $EvalRef("EvalExtern", $Qualified($Maybe("Just", moduleName), a$1)),
              $Tuple(a$2, a$3)
            )
          )
        ))
      ))
    ))
  ))
);
var parseDirectiveHeader = (moduleName) => foldlArray((v) => {
  const $0 = v.errors;
  const $1 = v.exports;
  const $2 = v.locals;
  return (v1) => {
    if (v1.tag === "LineComment") {
      const $3 = stripPrefix("@inline")(trim(v1._1));
      if ($3.tag === "Just") {
        const line$p = trim($3._1);
        const v3 = runParser(lex(line$p))((state1, more, resume, done) => parseDirectiveExport(moduleName)(
          state1.consumed ? { ...state1, consumed: false } : state1,
          more,
          (state3, error3) => {
            if (state3.consumed) {
              return resume(state3, error3);
            }
            return parseDirective(state1, more, resume, (state2, a) => done(state2, $Either("Right", a)));
          },
          (state2, a) => done(state2, $Either("Left", a))
        ));
        if (v3.tag === "Left") {
          return { errors: snoc($0)($Tuple(line$p, v3._1)), locals: $2, exports: $1 };
        }
        if (v3.tag === "Right") {
          if (v3._1._1.tag === "Left") {
            return { errors: $0, locals: $2, exports: insertDirective(v3._1._1._1._1)(v3._1._1._1._2._1)(v3._1._1._1._2._2)($1) };
          }
          if (v3._1._1.tag === "Right") {
            return { errors: $0, locals: insertDirective(v3._1._1._1._1)(v3._1._1._1._2._1)(v3._1._1._1._2._2)($2), exports: $1 };
          }
        }
        fail();
      }
    }
    return { errors: $0, locals: $2, exports: $1 };
  };
})({ errors: [], locals: Leaf, exports: Leaf });

// output-es/PureScript.Backend.Optimizer.Convert/index.js
var $CaseRowGuardedExpr = (tag, _1) => ({ tag, _1 });
var $PatternCase = (tag, _1, _2) => ({ tag, _1, _2 });
var eq = /* @__PURE__ */ eqArrayImpl(eqStringImpl);
var compare3 = /* @__PURE__ */ (() => ordArray(ordString).compare)();
var compare12 = /* @__PURE__ */ (() => ordQualified(ordString).compare)();
var ordQualified3 = /* @__PURE__ */ ordQualified(ordString);
var lookup4 = (k) => {
  const go = (go$a0$copy) => {
    let go$a0 = go$a0$copy, go$c = true, go$r;
    while (go$c) {
      const v = go$a0;
      if (v.tag === "Leaf") {
        go$c = false;
        go$r = Nothing;
        continue;
      }
      if (v.tag === "Node") {
        const v1 = ordString.compare(k)(v._3);
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
          go$r = $Maybe("Just", v._4);
          continue;
        }
      }
      fail();
    }
    return go$r;
  };
  return go;
};
var monoidSemigroupMap2 = /* @__PURE__ */ monoidSemigroupMap(ordString)(semigroupFirst2);
var foldMap5 = /* @__PURE__ */ (() => foldableSet.foldMap(monoidSemigroupMap2))();
var toUnfoldable2 = /* @__PURE__ */ (() => {
  const $0 = unfoldableArray.unfoldr((xs) => {
    if (xs.tag === "Nil") {
      return Nothing;
    }
    if (xs.tag === "Cons") {
      return $Maybe("Just", $Tuple(xs._1, xs._2));
    }
    fail();
  });
  return (x) => $0((() => {
    const go = (m$p, z$p) => {
      if (m$p.tag === "Leaf") {
        return z$p;
      }
      if (m$p.tag === "Node") {
        return go(m$p._5, $List("Cons", m$p._3, go(m$p._6, z$p)));
      }
      fail();
    };
    return go(x, Nil);
  })());
})();
var foldMap22 = /* @__PURE__ */ (() => foldableArray.foldMap(monoidArray))();
var fromFoldable1 = /* @__PURE__ */ foldlArray((m) => (a) => insert(ordString)(a)()(m))(Leaf);
var lookup12 = (k) => {
  const go = (go$a0$copy) => {
    let go$a0 = go$a0$copy, go$c = true, go$r;
    while (go$c) {
      const v = go$a0;
      if (v.tag === "Leaf") {
        go$c = false;
        go$r = Nothing;
        continue;
      }
      if (v.tag === "Node") {
        const v1 = ordQualified3.compare(k)(v._3);
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
          go$r = $Maybe("Just", v._4);
          continue;
        }
      }
      fail();
    }
    return go$r;
  };
  return go;
};
var lookup23 = (k) => {
  const go = (go$a0$copy) => {
    let go$a0 = go$a0$copy, go$c = true, go$r;
    while (go$c) {
      const v = go$a0;
      if (v.tag === "Leaf") {
        go$c = false;
        go$r = Nothing;
        continue;
      }
      if (v.tag === "Node") {
        const v1 = ordEvalRef.compare(k)(v._3);
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
          go$r = $Maybe("Just", v._4);
          continue;
        }
      }
      fail();
    }
    return go$r;
  };
  return go;
};
var lookup32 = (k) => {
  const go = (go$a0$copy) => {
    let go$a0 = go$a0$copy, go$c = true, go$r;
    while (go$c) {
      const v = go$a0;
      if (v.tag === "Leaf") {
        go$c = false;
        go$r = Nothing;
        continue;
      }
      if (v.tag === "Node") {
        const v1 = ordInlineAccessor.compare(k)(v._3);
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
          go$r = $Maybe("Just", v._4);
          continue;
        }
      }
      fail();
    }
    return go$r;
  };
  return go;
};
var analyzeEffectBlock2 = /* @__PURE__ */ analyzeEffectBlock(hasAnalysisBackendExpr)(hasSyntaxBackendExpr);
var analyze2 = /* @__PURE__ */ analyze(hasAnalysisBackendExpr)(hasSyntaxBackendExpr);
var foldMap32 = /* @__PURE__ */ (() => foldableArray.foldMap(/* @__PURE__ */ (() => {
  const semigroupRecord1 = { append: (ra) => (rb) => ({ rowsNoMatch: [...ra.rowsNoMatch, ...rb.rowsNoMatch], rowsWithMatch: [...ra.rowsWithMatch, ...rb.rowsWithMatch] }) };
  return { mempty: { rowsNoMatch: [], rowsWithMatch: [] }, Semigroup0: () => semigroupRecord1 };
})()))();
var lookup42 = (k) => {
  const go = (go$a0$copy) => {
    let go$a0 = go$a0$copy, go$c = true, go$r;
    while (go$c) {
      const v = go$a0;
      if (v.tag === "Leaf") {
        go$c = false;
        go$r = Nothing;
        continue;
      }
      if (v.tag === "Node") {
        const v1 = ordString.compare(k)(v._3);
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
          go$r = $Maybe("Just", v._4);
          continue;
        }
      }
      fail();
    }
    return go$r;
  };
  return go;
};
var forWithIndex = /* @__PURE__ */ (() => {
  const $0 = traversableWithIndexArray.traverseWithIndex(applicativeFn);
  return (b) => (a) => $0(a)(b);
})();
var zipWithA2 = /* @__PURE__ */ zipWithA(applicativeFn);
var traverse = /* @__PURE__ */ (() => traversableLiteral.traverse(applicativeFn))();
var traverse1 = /* @__PURE__ */ (() => traversableArray.traverse(applicativeFn))();
var traverse3 = /* @__PURE__ */ (() => traversableArray.traverse(applicativeFn))();
var append2 = /* @__PURE__ */ (() => semigroupSemigroupMap(ordString)(semigroupFirst2).append)();
var toUnfoldable12 = /* @__PURE__ */ (() => {
  const $0 = unfoldableArray.unfoldr(stepUnfoldr);
  return (x) => $0($MapIter("IterNode", x, IterLeaf));
})();
var foldMap42 = /* @__PURE__ */ (() => foldableArray.foldMap(monoidSemigroupMap2))();
var $$for = /* @__PURE__ */ (() => {
  const traverse22 = traversableArray.traverse(applicativeFn);
  return (x) => (f) => traverse22(f)(x);
})();
var member = (k) => {
  const go = (go$a0$copy) => {
    let go$a0 = go$a0$copy, go$c = true, go$r;
    while (go$c) {
      const v = go$a0;
      if (v.tag === "Leaf") {
        go$c = false;
        go$r = false;
        continue;
      }
      if (v.tag === "Node") {
        const v1 = ordQualified3.compare(k)(v._3);
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
      fail();
    }
    return go$r;
  };
  return go;
};
var alter3 = /* @__PURE__ */ alter(ordEvalRef);
var mapAccumL2 = /* @__PURE__ */ mapAccumL(traversableArray);
var fromFoldable22 = /* @__PURE__ */ foldlArray((m) => (a) => insert(ordString)(a)()(m))(Leaf);
var member1 = (k) => {
  const go = (go$a0$copy) => {
    let go$a0 = go$a0$copy, go$c = true, go$r;
    while (go$c) {
      const v = go$a0;
      if (v.tag === "Leaf") {
        go$c = false;
        go$r = false;
        continue;
      }
      if (v.tag === "Node") {
        const v1 = ordString.compare(k)(v._3);
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
      fail();
    }
    return go$r;
  };
  return go;
};
var fromFoldable3 = /* @__PURE__ */ fromFoldable(ordString)(foldableArray);
var fromFoldable4 = /* @__PURE__ */ fromFoldable(ordString)(foldableArray);
var maximum2 = /* @__PURE__ */ maximum(ordInt)(foldable1NonEmptyArray);
var mapAccumR2 = /* @__PURE__ */ mapAccumR(traversableArray);
var foldMap52 = /* @__PURE__ */ (() => foldableArray.foldMap(monoidSet(ordQualified3)))();
var fromFoldable5 = /* @__PURE__ */ foldlArray((m) => (a) => insert(ordReExport)(a)()(m))(Leaf);
var PatWild = /* @__PURE__ */ $PatternCase("PatWild");
var eqPatternCase = {
  eq: (x) => (y) => {
    if (x.tag === "PatWild") {
      return y.tag === "PatWild";
    }
    if (x.tag === "PatRecord") {
      return y.tag === "PatRecord" && eq(x._1)(y._1);
    }
    if (x.tag === "PatProduct") {
      return y.tag === "PatProduct" && (x._1._1.tag === "Nothing" ? y._1._1.tag === "Nothing" : x._1._1.tag === "Just" && y._1._1.tag === "Just" && x._1._1._1 === y._1._1._1) && x._1._2 === y._1._2 && (x._2._1.tag === "Nothing" ? y._2._1.tag === "Nothing" : x._2._1.tag === "Just" && y._2._1.tag === "Just" && x._2._1._1 === y._2._1._1) && x._2._2 === y._2._2;
    }
    if (x.tag === "PatArray") {
      return y.tag === "PatArray" && x._1 === y._1;
    }
    if (x.tag === "PatSum") {
      return y.tag === "PatSum" && (x._1._1.tag === "Nothing" ? y._1._1.tag === "Nothing" : x._1._1.tag === "Just" && y._1._1.tag === "Just" && x._1._1._1 === y._1._1._1) && x._1._2 === y._1._2 && (x._2._1.tag === "Nothing" ? y._2._1.tag === "Nothing" : x._2._1.tag === "Just" && y._2._1.tag === "Just" && x._2._1._1 === y._2._1._1) && x._2._2 === y._2._2;
    }
    if (x.tag === "PatInt") {
      return y.tag === "PatInt" && x._1 === y._1;
    }
    if (x.tag === "PatNumber") {
      return y.tag === "PatNumber" && x._1 === y._1;
    }
    if (x.tag === "PatString") {
      return y.tag === "PatString" && x._1 === y._1;
    }
    if (x.tag === "PatChar") {
      return y.tag === "PatChar" && x._1 === y._1;
    }
    return x.tag === "PatBoolean" && y.tag === "PatBoolean" && x._1 === y._1;
  }
};
var ordPatternCase = {
  compare: (x) => (y) => {
    if (x.tag === "PatWild") {
      if (y.tag === "PatWild") {
        return EQ;
      }
      return LT;
    }
    if (y.tag === "PatWild") {
      return GT;
    }
    if (x.tag === "PatRecord") {
      if (y.tag === "PatRecord") {
        return compare3(x._1)(y._1);
      }
      return LT;
    }
    if (y.tag === "PatRecord") {
      return GT;
    }
    if (x.tag === "PatProduct") {
      if (y.tag === "PatProduct") {
        const v = compare12(x._1)(y._1);
        if (v === "LT") {
          return LT;
        }
        if (v === "GT") {
          return GT;
        }
        return ordQualified3.compare(x._2)(y._2);
      }
      return LT;
    }
    if (y.tag === "PatProduct") {
      return GT;
    }
    if (x.tag === "PatArray") {
      if (y.tag === "PatArray") {
        return ordInt.compare(x._1)(y._1);
      }
      return LT;
    }
    if (y.tag === "PatArray") {
      return GT;
    }
    if (x.tag === "PatSum") {
      if (y.tag === "PatSum") {
        const v = compare12(x._1)(y._1);
        if (v === "LT") {
          return LT;
        }
        if (v === "GT") {
          return GT;
        }
        return ordQualified3.compare(x._2)(y._2);
      }
      return LT;
    }
    if (y.tag === "PatSum") {
      return GT;
    }
    if (x.tag === "PatInt") {
      if (y.tag === "PatInt") {
        return ordInt.compare(x._1)(y._1);
      }
      return LT;
    }
    if (y.tag === "PatInt") {
      return GT;
    }
    if (x.tag === "PatNumber") {
      if (y.tag === "PatNumber") {
        return ordNumber.compare(x._1)(y._1);
      }
      return LT;
    }
    if (y.tag === "PatNumber") {
      return GT;
    }
    if (x.tag === "PatString") {
      if (y.tag === "PatString") {
        return ordString.compare(x._1)(y._1);
      }
      return LT;
    }
    if (y.tag === "PatString") {
      return GT;
    }
    if (x.tag === "PatChar") {
      if (y.tag === "PatChar") {
        return ordChar.compare(x._1)(y._1);
      }
      return LT;
    }
    if (y.tag === "PatChar") {
      return GT;
    }
    if (x.tag === "PatBoolean" && y.tag === "PatBoolean") {
      return ordBoolean.compare(x._1)(y._1);
    }
    fail();
  },
  Eq0: () => eqPatternCase
};
var monoidSet2 = /* @__PURE__ */ monoidSet(ordPatternCase);
var foldMapWithIndex = /* @__PURE__ */ (() => foldableWithIndexArray.foldMapWithIndex((() => {
  const Semigroup0 = monoidSet2.Semigroup0();
  const semigroupRecord1 = {
    append: (ra) => (rb) => ({ aScore: ra.aScore + rb.aScore | 0, ctors: Semigroup0.append(ra.ctors)(rb.ctors), tailRowIndices: [...ra.tailRowIndices, ...rb.tailRowIndices] })
  };
  return { mempty: { aScore: 0, ctors: monoidSet2.mempty, tailRowIndices: [] }, Semigroup0: () => semigroupRecord1 };
})()))();
var toExternImpl = (env) => (group2) => (expr) => {
  const $0 = () => {
    const v = freeze(expr);
    return $Tuple($Tuple(v._1, $ExternImpl("ExternExpr", group2, v._2)), v._2);
  };
  if (expr.tag === "ExprSyntax") {
    if (expr._2.tag === "Lit") {
      if (expr._2._1.tag === "LitRecord") {
        const propsWithAnalysis = arrayMap((m) => $Prop(m._1, freeze(m._2)))(expr._2._1._1);
        return $Tuple(
          $Tuple(expr._1, $ExternImpl("ExternDict", group2, propsWithAnalysis)),
          $BackendSyntax(
            "Lit",
            $Literal(
              "LitRecord",
              arrayMap((m) => $Prop(m._1, m._2._2))(propsWithAnalysis)
            )
          )
        );
      }
      return $0();
    }
    if (expr._2.tag === "CtorDef") {
      const v = freeze(expr);
      return $Tuple(
        $Tuple(
          v._1,
          $ExternImpl(
            "ExternCtor",
            (() => {
              const $1 = lookup4(expr._2._2)(env.dataTypes);
              if ($1.tag === "Just") {
                return $1._1;
              }
              fail();
            })(),
            expr._2._1,
            expr._2._2,
            expr._2._3,
            expr._2._4
          )
        ),
        v._2
      );
    }
  }
  return $0();
};
var toCaseRowVars = (v) => {
  const $0 = v.column;
  return foldMap5((x) => $$$Map("Node", 1, 1, x, $0, Leaf, Leaf))(v.pattern.vars);
};
var patternVars = (v) => [...toUnfoldable2(v.pattern.vars), ...foldMap22(patternVars)(v.pattern.subterms)];
var normalizeCaseRows = (x) => {
  const go = (go$a0$copy) => (go$a1$copy) => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
    while (go$c) {
      const columnIdx = go$a0, columnsAcc = go$a1;
      const nextColumnFields = foldlArray((acc) => (next) => {
        if (columnIdx >= 0 && columnIdx < next.patterns.length) {
          const $02 = next.patterns[columnIdx];
          return $Maybe(
            "Just",
            (() => {
              if ($02.pattern.patternCase.tag === "PatRecord") {
                const keys2 = fromFoldable1($02.pattern.patternCase._1);
                if (acc.tag === "Nothing") {
                  return keys2;
                }
                if (acc.tag === "Just") {
                  return unsafeUnionWith(ordString.compare, $$const, keys2, acc._1);
                }
                fail();
              }
              if (acc.tag === "Nothing") {
                return Leaf;
              }
              if (acc.tag === "Just") {
                return acc._1;
              }
              fail();
            })()
          );
        }
        return Nothing;
      })(Nothing)(x);
      if (nextColumnFields.tag === "Nothing") {
        go$c = false;
        go$r = columnsAcc;
        continue;
      }
      if (nextColumnFields.tag === "Just") {
        go$a0 = columnIdx + 1 | 0;
        go$a1 = snoc(columnsAcc)(nextColumnFields._1);
        continue;
      }
      fail();
    }
    return go$r;
  };
  const $0 = go(0)([]);
  return arrayMap((nextRow) => ({
    ...nextRow,
    patterns: zipWithImpl(
      (allFieldsSet) => (pat) => {
        if (pat.pattern.patternCase.tag === "PatRecord") {
          const v1 = unzip(arrayMap(head)(groupAllBy((x$1) => (y) => ordString.compare(x$1._1)(y._1))([
            ...zipWithImpl(Tuple, pat.pattern.patternCase._1, pat.pattern.subterms),
            ...arrayMap((fieldName) => $Tuple(
              fieldName,
              {
                accessor: $BackendAccessor("GetProp", fieldName),
                pattern: { vars: Leaf, patternCase: PatWild, subterms: [] }
              }
            ))(toUnfoldable2(allFieldsSet))
          ])));
          return { ...pat, pattern: { ...pat.pattern, patternCase: $PatternCase("PatRecord", v1._1), subterms: v1._2 } };
        }
        return pat;
      },
      $0,
      nextRow.patterns
    )
  }))(x);
};
var makeExternEvalSpine = (conv) => (env) => (qual) => (spine) => {
  const $0 = lookup12(qual)(conv.foreignSemantics);
  const result = (() => {
    if ($0.tag === "Just") {
      return $0._1(env)(qual)(spine);
    }
    if ($0.tag === "Nothing") {
      return Nothing;
    }
    fail();
  })();
  if (result.tag === "Nothing") {
    const $1 = lookup12(qual)(conv.implementations);
    if ($1.tag === "Just") {
      return evalExternFromImpl({ ...env, locals: [] })(qual)($1._1)(spine);
    }
    if ($1.tag === "Nothing") {
      return Nothing;
    }
    fail();
  }
  return result;
};
var makeExternEvalRef = (conv) => (env) => (qual) => {
  const $0 = lookup12(qual)(conv.implementations);
  if ($0.tag === "Just") {
    return $Maybe("Just", evalExternRefFromImpl(env)(qual)($0._1));
  }
  return Nothing;
};
var intro = (dictFoldable) => (ident) => (lvl) => (f) => (env) => f({
  ...env,
  currentLevel: env.currentLevel + 1 | 0,
  toLevel: dictFoldable.foldr((a) => insert(ordString)(a)(lvl))(env.toLevel)(ident)
});
var inferTransitiveDirective = (directives) => (impl) => (backendExpr) => (cfn) => {
  const $0 = (() => {
    if (impl.tag === "ExternExpr") {
      if (impl._2.tag === "App") {
        if (impl._2._1.tag === "Var") {
          const v = lookup23($EvalRef("EvalExtern", impl._2._1._1))(directives);
          if (v.tag === "Just") {
            const go = (m$p, z$p) => {
              if (m$p.tag === "Leaf") {
                return z$p;
              }
              if (m$p.tag === "Node") {
                return go(
                  m$p._5,
                  (() => {
                    const $02 = m$p._4;
                    const $12 = go(m$p._6, z$p);
                    const $2 = (prop) => insert(ordInlineAccessor)($InlineAccessor(
                      "InlineSpineProp",
                      prop
                    ))($02)(insert(ordInlineAccessor)($InlineAccessor(
                      "InlineProp",
                      prop
                    ))($02)($12));
                    if ($02.tag === "InlineArity" && m$p._3.tag === "InlineRef") {
                      return insert(ordInlineAccessor)(InlineRef)($InlineDirective(
                        "InlineArity",
                        $02._1 - impl._2._2.length | 0
                      ))($12);
                    }
                    if (m$p._3.tag === "InlineSpineProp") {
                      return $2(m$p._3._1);
                    }
                    return $12;
                  })()
                );
              }
              fail();
            };
            const newDirs = go(v._1, Leaf);
            if (newDirs.tag === "Leaf") {
              return Nothing;
            }
            return $Maybe("Just", newDirs);
          }
        }
        return Nothing;
      }
      if (impl._2.tag === "Accessor" && impl._2._1.tag === "App" && impl._2._1._1.tag === "Var" && impl._2._2.tag === "GetProp") {
        const $02 = lookup23($EvalRef("EvalExtern", impl._2._1._1._1))(directives);
        const $12 = lookup32($InlineAccessor("InlineSpineProp", impl._2._2._1));
        const v = (() => {
          if ($02.tag === "Just") {
            return $12($02._1);
          }
          if ($02.tag === "Nothing") {
            return Nothing;
          }
          fail();
        })();
        if (v.tag === "Just" && v._1.tag === "InlineArity") {
          return $Maybe(
            "Just",
            $$$Map(
              "Node",
              1,
              1,
              InlineRef,
              $InlineDirective("InlineArity", v._1._1),
              Leaf,
              Leaf
            )
          );
        }
      }
    }
    return Nothing;
  })();
  const $1 = (() => {
    if (backendExpr.tag === "ExprSyntax" && backendExpr._2.tag === "App" && backendExpr._2._1.tag === "ExprSyntax" && backendExpr._2._1._2.tag === "Var") {
      const $12 = lookup23($EvalRef("EvalExtern", backendExpr._2._1._2._1))(directives);
      const $2 = lookup32(InlineRef);
      const v = (() => {
        if ($12.tag === "Just") {
          return $2($12._1);
        }
        if ($12.tag === "Nothing") {
          return Nothing;
        }
        fail();
      })();
      if (v.tag === "Just" && v._1.tag === "InlineArity" && cfn.tag === "ExprApp" && cfn._1.meta.tag === "Just" && cfn._1.meta._1.tag === "IsSyntheticApp" && backendExpr._2._2.length >= v._1._1) {
        return $Maybe(
          "Just",
          $$$Map(
            "Node",
            1,
            1,
            InlineRef,
            InlineAlways,
            Leaf,
            Leaf
          )
        );
      }
    }
    return Nothing;
  })();
  if ($0.tag === "Nothing") {
    return $1;
  }
  return $0;
};
var getCtx = (env) => {
  const lookupExtern = (qual) => (acc) => {
    const $0 = lookup12(qual)(env.implementations);
    if ($0.tag === "Just") {
      if ($0._1._2.tag === "ExternExpr") {
        if (acc.tag === "Nothing") {
          return $Maybe("Just", $0._1._1);
        }
        return Nothing;
      }
      if ($0._1._2.tag === "ExternDict") {
        if (acc.tag === "Just") {
          const $1 = acc._1;
          const $2 = findMapImpl(
            Nothing,
            isJust,
            (v) => {
              if ($1 === v._1) {
                return $Maybe("Just", v._2);
              }
              return Nothing;
            },
            $0._1._2._2
          );
          if ($2.tag === "Just") {
            return $Maybe("Just", $2._1._1);
          }
          return Nothing;
        }
        if (acc.tag === "Nothing") {
          return $Maybe("Just", $0._1._1);
        }
        fail();
      }
      if ($0._1._2.tag === "ExternCtor") {
        return Nothing;
      }
      fail();
    }
    if ($0.tag === "Nothing") {
      return Nothing;
    }
    fail();
  };
  return {
    currentLevel: env.currentLevel,
    lookupExtern,
    analyze: (v) => {
      const $0 = v.effect;
      return (expr) => {
        const v1 = env.analyzeCustom(v)(expr);
        if (v1.tag === "Just") {
          return v1._1;
        }
        if (v1.tag === "Nothing") {
          if ($0) {
            return analyzeEffectBlock2(lookupExtern)(expr);
          }
          return analyze2(lookupExtern)(expr);
        }
        fail();
      };
    },
    effect: false
  };
};
var decompose = (chosenColumn) => {
  const checkMatch = (p) => {
    if (p.column === chosenColumn.column && (eqPatternCase.eq(p.pattern.patternCase)(PatWild) || eqPatternCase.eq(chosenColumn.pattern.patternCase)(p.pattern.patternCase))) {
      return { nonMatchesBefore: [], match: $Maybe("Just", { match: p, nonMatchesAfter: [] }) };
    }
    return { nonMatchesBefore: [p], match: Nothing };
  };
  return foldMap32((row) => {
    const v = row.patterns.length > 0 ? $Maybe("Just", row.patterns) : Nothing;
    if (v.tag === "Nothing") {
      return _crashWith("decompose - nextRow.patterns cannot be empty since the first row contains at least one `PatCtor` patternCase");
    }
    if (v.tag === "Just") {
      const $0 = v._1;
      const len = $0.length;
      const v1 = (() => {
        const go = (go$a0$copy) => (go$a1$copy) => {
          let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
          while (go$c) {
            const ix = go$a0, acc = go$a1;
            if (ix === len) {
              go$c = false;
              go$r = acc;
              continue;
            }
            go$a0 = ix + 1 | 0;
            go$a1 = (() => {
              const $1 = checkMatch($0[ix]);
              if (acc.match.tag === "Just") {
                if ($1.match.tag === "Just") {
                  return _crashWith("mergeResults - impossible: cannot match the same column twice in the same row");
                }
                if ($1.match.tag === "Nothing") {
                  return { ...acc, match: $Maybe("Just", { ...acc.match._1, nonMatchesAfter: [...acc.match._1.nonMatchesAfter, ...$1.nonMatchesBefore] }) };
                }
                fail();
              }
              if (acc.match.tag === "Nothing") {
                if ($1.match.tag === "Nothing") {
                  return { ...$1, nonMatchesBefore: [...acc.nonMatchesBefore, ...$1.nonMatchesBefore] };
                }
                if ($1.match.tag === "Just") {
                  return { ...$1, nonMatchesBefore: [...acc.nonMatchesBefore, ...$1.nonMatchesBefore] };
                }
              }
              fail();
            })();
          }
          return go$r;
        };
        return go(1)(checkMatch((() => {
          if (0 < $0.length) {
            return $0[0];
          }
          fail();
        })()));
      })();
      if (v1.match.tag === "Just") {
        return {
          rowsWithMatch: [{ guardFn: row.guardFn, vars: row.vars, nonMatchesBefore: v1.nonMatchesBefore, match: v1.match._1.match, nonMatchesAfter: v1.match._1.nonMatchesAfter }],
          rowsNoMatch: eqPatternCase.eq(v1.match._1.match.pattern.patternCase)(PatWild) ? [row] : []
        };
      }
      if (v1.match.tag === "Nothing") {
        return { rowsWithMatch: [], rowsNoMatch: [row] };
      }
    }
    fail();
  });
};
var chooseNextPattern = (row0Patterns) => (tailRows) => {
  const expandIfPossible = findMapImpl(
    Nothing,
    isJust,
    (v) => {
      if (v._2.pattern.patternCase.tag === "PatRecord") {
        return $Maybe("Just", v._2);
      }
      if (v._2.pattern.patternCase.tag === "PatProduct") {
        return $Maybe("Just", v._2);
      }
      return Nothing;
    },
    row0Patterns
  );
  if (expandIfPossible.tag === "Just") {
    return expandIfPossible._1;
  }
  if (expandIfPossible.tag === "Nothing") {
    const $0 = foldlArray((acc) => (next) => {
      if (acc.tag === "Nothing") {
        return $Maybe("Just", [next]);
      }
      if (acc.tag === "Just") {
        const v = ordInt.compare((() => {
          if (0 < acc._1.length) {
            return acc._1[0].bScore;
          }
          fail();
        })())(next.bScore);
        if (v === "GT") {
          return acc;
        }
        if (v === "EQ") {
          return $Maybe("Just", snoc(acc._1)(next));
        }
        if (v === "LT") {
          return $Maybe("Just", [next]);
        }
      }
      fail();
    })(Nothing);
    const $1 = foldlArray((acc) => (next) => {
      if (acc.tag === "Nothing") {
        return $Maybe("Just", [next]);
      }
      if (acc.tag === "Just") {
        const v = ordInt.compare((() => {
          if (0 < acc._1.length) {
            return acc._1[0].aScore;
          }
          fail();
        })())(next.aScore);
        if (v === "GT") {
          return acc;
        }
        if (v === "EQ") {
          return $Maybe("Just", snoc(acc._1)(next));
        }
        if (v === "LT") {
          return $Maybe("Just", [next]);
        }
      }
      fail();
    })(Nothing);
    const $2 = foldlArray((acc) => (next) => {
      if (acc.tag === "Nothing") {
        return $Maybe("Just", [next]);
      }
      if (acc.tag === "Just") {
        const v = ordInt.compare((() => {
          if (0 < acc._1.length) {
            return acc._1[0].pScore;
          }
          fail();
        })())(next.pScore);
        if (v === "GT") {
          return acc;
        }
        if (v === "EQ") {
          return $Maybe("Just", snoc(acc._1)(next));
        }
        if (v === "LT") {
          return $Maybe("Just", [next]);
        }
      }
      fail();
    })(Nothing)(arrayMap((v) => {
      const $22 = v._1;
      const $3 = v._2;
      const matchingCols = foldMapWithIndex((rowIdx) => (row) => {
        if ($22 >= 0 && $22 < row.patterns.length) {
          return {
            tailRowIndices: (() => {
              const $4 = [rowIdx + 1 | 0];
              if (eqPatternCase.eq($3.pattern.patternCase)(row.patterns[$22].pattern.patternCase)) {
                return $4;
              }
              return [];
            })(),
            ctors: (() => {
              const $4 = $$$Map("Node", 1, 1, row.patterns[$22].pattern.patternCase, void 0, Leaf, Leaf);
              if (!eqPatternCase.eq(row.patterns[$22].pattern.patternCase)(PatWild)) {
                return $4;
              }
              return monoidSet2.mempty;
            })(),
            aScore: -filterImpl((x) => !eqPatternCase.eq(PatWild)(x.pattern.patternCase), row.patterns[$22].pattern.subterms).length
          };
        }
        return _crashWith("Impossible: rows' column lengths differ in pattern match");
      })(tailRows);
      return {
        pattern: $3,
        pScore: foldlArray((l) => (r) => {
          if ((l + 1 | 0) === r) {
            return r;
          }
          return l;
        })(0)(matchingCols.tailRowIndices),
        bScore: (() => {
          const $4 = insert(ordPatternCase)($3.pattern.patternCase)()(matchingCols.ctors);
          if ($4.tag === "Leaf") {
            return 0;
          }
          if ($4.tag === "Node") {
            return -$4._2;
          }
          fail();
        })(),
        aScore: matchingCols.aScore
      };
    })(row0Patterns));
    if ($2.tag === "Just") {
      const $3 = $0($2._1);
      if ($3.tag === "Just") {
        const $4 = $1($3._1);
        if ($4.tag === "Just") {
          if (0 < $4._1.length) {
            return $4._1[0].pattern;
          }
          fail();
        }
        if (0 < row0Patterns.length) {
          return row0Patterns[0]._2;
        }
        fail();
      }
      if ($3.tag === "Nothing" && 0 < row0Patterns.length) {
        return row0Patterns[0]._2;
      }
      fail();
    }
    if ($2.tag === "Nothing" && 0 < row0Patterns.length) {
      return row0Patterns[0]._2;
    }
  }
  fail();
};
var buildM = (a) => (env) => build(getCtx(env))(a);
var make = (a) => {
  const $0 = traversableBackendSyntax.traverse(applicativeFn)(identity3)(a);
  return (x) => build(getCtx(x))($0(x));
};
var guardBoolean = (n) => (lhs) => $BackendSyntax(
  "PrimOp",
  $BackendOperator(
    "Op2",
    $BackendOperator2("OpBooleanOrd", OpEq),
    lhs,
    make($BackendSyntax("Lit", $Literal("LitBoolean", n)))
  )
);
var guardChar = (n) => (lhs) => $BackendSyntax(
  "PrimOp",
  $BackendOperator(
    "Op2",
    $BackendOperator2("OpCharOrd", OpEq),
    lhs,
    make($BackendSyntax("Lit", $Literal("LitChar", n)))
  )
);
var guardInt = (n) => (lhs) => $BackendSyntax(
  "PrimOp",
  $BackendOperator(
    "Op2",
    $BackendOperator2("OpIntOrd", OpEq),
    lhs,
    make($BackendSyntax("Lit", $Literal("LitInt", n)))
  )
);
var guardArrayLength = (n) => (lhs) => $BackendSyntax(
  "PrimOp",
  $BackendOperator(
    "Op2",
    $BackendOperator2("OpIntOrd", OpEq),
    make($BackendSyntax(
      "PrimOp",
      $BackendOperator("Op1", OpArrayLength, lhs)
    )),
    make($BackendSyntax("Lit", $Literal("LitInt", n)))
  )
);
var guardNumber = (n) => (lhs) => $BackendSyntax(
  "PrimOp",
  $BackendOperator(
    "Op2",
    $BackendOperator2("OpNumberOrd", OpEq),
    lhs,
    make($BackendSyntax("Lit", $Literal("LitNumber", n)))
  )
);
var guardString = (n) => (lhs) => $BackendSyntax(
  "PrimOp",
  $BackendOperator(
    "Op2",
    $BackendOperator2("OpStringOrd", OpEq),
    lhs,
    make($BackendSyntax("Lit", $Literal("LitString", n)))
  )
);
var makeGuard = (lvl) => (g) => (inner) => (def) => make($BackendSyntax(
  "Branch",
  [$Pair(make(g(make($BackendSyntax("Local", Nothing, lvl)))), inner)],
  def
));
var makeLet2 = (id2) => (a) => (k) => (x) => {
  if (id2.tag === "Nothing") {
    return make($BackendSyntax(
      "Let",
      id2,
      x.currentLevel,
      a,
      (() => {
        const $0 = k(x.currentLevel);
        return (env) => $0({ ...env, currentLevel: env.currentLevel + 1 | 0 });
      })()
    ))(x);
  }
  if (id2.tag === "Just") {
    return make($BackendSyntax("Let", id2, x.currentLevel, a, intro(foldableArray)([id2._1])(x.currentLevel)(k(x.currentLevel))))(x);
  }
  fail();
};
var makeUncurriedAbs = (args) => (cb) => foldrArray((ident) => (next) => (tmps) => (x) => intro(foldableArray)([ident])(x.currentLevel)(next(snoc(tmps)($Tuple(
  $Maybe("Just", ident),
  x.currentLevel
))))(x))((tmps) => make($BackendSyntax("UncurriedAbs", tmps, cb(tmps))))(args)([]);
var patternFail = /* @__PURE__ */ make(/* @__PURE__ */ $BackendSyntax("Fail", "Failed pattern match"));
var binderToPattern = (v) => {
  if (v.tag === "BinderNull") {
    return (v$1) => ({ vars: Leaf, patternCase: PatWild, subterms: [] });
  }
  if (v.tag === "BinderVar") {
    return (v$1) => ({ vars: $$$Map("Node", 1, 1, v._2, void 0, Leaf, Leaf), patternCase: PatWild, subterms: [] });
  }
  if (v.tag === "BinderNamed") {
    const $0 = v._2;
    const $1 = binderToPattern(v._3);
    return (x) => {
      const $2 = $1(x);
      return { ...$2, vars: insert(ordString)($0)()($2.vars) };
    };
  }
  if (v.tag === "BinderLit") {
    if (v._2.tag === "LitInt") {
      return (v$1) => ({ vars: Leaf, patternCase: $PatternCase("PatInt", v._2._1), subterms: [] });
    }
    if (v._2.tag === "LitNumber") {
      return (v$1) => ({ vars: Leaf, patternCase: $PatternCase("PatNumber", v._2._1), subterms: [] });
    }
    if (v._2.tag === "LitString") {
      return (v$1) => ({ vars: Leaf, patternCase: $PatternCase("PatString", v._2._1), subterms: [] });
    }
    if (v._2.tag === "LitChar") {
      return (v$1) => ({ vars: Leaf, patternCase: $PatternCase("PatChar", v._2._1), subterms: [] });
    }
    if (v._2.tag === "LitBoolean") {
      return (v$1) => ({ vars: Leaf, patternCase: $PatternCase("PatBoolean", v._2._1), subterms: [] });
    }
    if (v._2.tag === "LitArray") {
      const $0 = $PatternCase("PatArray", v._2._1.length);
      const $1 = forWithIndex(v._2._1)((idx) => (nextArg) => {
        const $12 = binderToPattern(nextArg);
        return (x) => ({ accessor: $BackendAccessor("GetIndex", idx), pattern: $12(x) });
      });
      return (x) => ({ vars: Leaf, patternCase: $0, subterms: $1(x) });
    }
    if (v._2.tag === "LitRecord") {
      const $0 = $PatternCase("PatRecord", arrayMap(propKey)(v._2._1));
      const $1 = forWithIndex(v._2._1)((idx) => (nextArg) => {
        const $12 = binderToPattern(nextArg._2);
        return (x) => ({ accessor: $BackendAccessor("GetProp", nextArg._1), pattern: $12(x) });
      });
      return (x) => ({ vars: Leaf, patternCase: $0, subterms: $1(x) });
    }
    fail();
  }
  if (v.tag === "BinderConstructor") {
    if (v._1.meta.tag === "Just" && v._1.meta._1.tag === "IsNewtype" && v._4.length === 1) {
      return binderToPattern(v._4[0]);
    }
    if (v._1.meta.tag === "Just") {
      if (v._1.meta._1.tag === "IsNewtype") {
        return _crashWith("Newtype binder didn't wrap 1 arg");
      }
      if (v._1.meta._1.tag === "IsConstructor") {
        if (v._1.meta._1._1 === "ProductType") {
          return (x) => {
            const v$1 = lookup12(v._3)(x.implementations);
            return {
              vars: Leaf,
              patternCase: $PatternCase("PatProduct", v._2, v._3),
              subterms: forWithIndex(zipWithImpl(
                Tuple,
                v._4,
                (() => {
                  if (v$1.tag === "Just" && v$1._1._2.tag === "ExternCtor") {
                    const $02 = lookup4(v._2._2)(x.dataTypes);
                    if ($02.tag === "Just") {
                      return v$1._1._2._5;
                    }
                    return v$1._1._2._5;
                  }
                  const $0 = lookup4(v._2._2)(x.dataTypes);
                  if ($0.tag === "Just") {
                    const $1 = lookup42(v._3._2)($0._1.constructors);
                    if ($1.tag === "Just") {
                      return $1._1.fields;
                    }
                    return _crashWith("Invariant broken: could not determine pattern matched constructor's fields during conversion.")(x);
                  }
                  if ($0.tag === "Nothing") {
                    return _crashWith("Invariant broken: could not determine pattern matched constructor's fields during conversion.")(x);
                  }
                  fail();
                })()
              ))((idx) => (nextArg) => {
                const $0 = binderToPattern(nextArg._1);
                return (x$1) => ({
                  accessor: $BackendAccessor(
                    "GetCtorField",
                    v._3,
                    ProductType,
                    v._2._2,
                    v._3._2,
                    nextArg._2,
                    idx
                  ),
                  pattern: $0(x$1)
                });
              })(x)
            };
          };
        }
        if (v._1.meta._1._1 === "SumType") {
          return (x) => {
            const v$1 = lookup12(v._3)(x.implementations);
            return {
              vars: Leaf,
              patternCase: $PatternCase("PatSum", v._2, v._3),
              subterms: forWithIndex(zipWithImpl(
                Tuple,
                v._4,
                (() => {
                  if (v$1.tag === "Just" && v$1._1._2.tag === "ExternCtor") {
                    const $02 = lookup4(v._2._2)(x.dataTypes);
                    if ($02.tag === "Just") {
                      return v$1._1._2._5;
                    }
                    return v$1._1._2._5;
                  }
                  const $0 = lookup4(v._2._2)(x.dataTypes);
                  if ($0.tag === "Just") {
                    const $1 = lookup42(v._3._2)($0._1.constructors);
                    if ($1.tag === "Just") {
                      return $1._1.fields;
                    }
                    return _crashWith("Invariant broken: could not determine pattern matched constructor's fields during conversion.")(x);
                  }
                  if ($0.tag === "Nothing") {
                    return _crashWith("Invariant broken: could not determine pattern matched constructor's fields during conversion.")(x);
                  }
                  fail();
                })()
              ))((idx) => (nextArg) => {
                const $0 = binderToPattern(nextArg._1);
                return (x$1) => ({
                  accessor: $BackendAccessor(
                    "GetCtorField",
                    v._3,
                    SumType,
                    v._2._2,
                    v._3._2,
                    nextArg._2,
                    idx
                  ),
                  pattern: $0(x$1)
                });
              })(x)
            };
          };
        }
      }
    }
    return _crashWith("binderToPattern - invalid meta");
  }
  fail();
};
var toBackendExpr = (v) => {
  if (v.tag === "ExprVar") {
    const $0 = v._2;
    return (x) => {
      const $1 = x.currentModule;
      const $2 = x.toLevel;
      const v2 = (v3) => {
        const v4 = (v5) => {
          if ($0._2 === "undefined" && $0._1.tag === "Just") {
            if ($0._1._1 === "Prim") {
              return buildM(PrimUndefined);
            }
            return buildM($BackendSyntax("Var", $0));
          }
          if ($0._1.tag === "Nothing") {
            return buildM($BackendSyntax("Var", $Qualified($Maybe("Just", $1), $0._2)));
          }
          return buildM($BackendSyntax("Var", $0));
        };
        if ($0._1.tag === "Just" && $0._1._1 === $1) {
          const $3 = lookup42($0._2)($2);
          if ($3.tag === "Just") {
            return buildM($BackendSyntax("Local", $Maybe("Just", $0._2), $3._1));
          }
        }
        return v4(true);
      };
      if ($0._1.tag === "Nothing") {
        const $3 = lookup42($0._2)($2);
        if ($3.tag === "Just") {
          return build(getCtx(x))($BackendSyntax("Local", $Maybe("Just", $0._2), $3._1));
        }
      }
      return v2(true)(x);
    };
  }
  if (v.tag === "ExprLit") {
    const $0 = traverse(toBackendExpr)(v._2);
    return (x) => build(getCtx(x))($BackendSyntax("Lit", $0(x)));
  }
  if (v.tag === "ExprConstructor") {
    const $0 = v._4;
    const $1 = v._3;
    const $2 = v._2;
    return (x) => build(getCtx(x))($BackendSyntax(
      "CtorDef",
      (() => {
        const v2 = lookup4($2)(x.dataTypes);
        if (v2.tag === "Just" && (() => {
          if (v2._1.constructors.tag === "Leaf") {
            return false;
          }
          if (v2._1.constructors.tag === "Node") {
            return v2._1.constructors._2 === 1;
          }
          fail();
        })()) {
          return ProductType;
        }
        return SumType;
      })(),
      $2,
      $1,
      $0
    ));
  }
  if (v.tag === "ExprAccessor") {
    const $0 = toBackendExpr(v._2);
    return (x) => build(getCtx(x))($BackendSyntax(
      "Accessor",
      $0(x),
      $BackendAccessor("GetProp", v._3)
    ));
  }
  if (v.tag === "ExprUpdate") {
    const $0 = toBackendExpr(v._2);
    const $1 = traverse1(traversableProp.traverse(applicativeFn)(toBackendExpr))(v._3);
    return (x) => build(getCtx(x))($BackendSyntax("Update", $0(x), $1(x)));
  }
  if (v.tag === "ExprAbs") {
    const $0 = v._2;
    const $1 = v._3;
    return (x) => make($BackendSyntax(
      "Abs",
      [$Tuple($Maybe("Just", $0), x.currentLevel)],
      intro(foldableArray)([$0])(x.currentLevel)(toBackendExpr($1))
    ))(x);
  }
  if (v.tag === "ExprApp" && v._2.tag === "ExprVar" && v._2._1.meta.tag === "Just" && v._2._1.meta._1.tag === "IsNewtype") {
    return toBackendExpr(v._3);
  }
  if (v.tag === "ExprApp") {
    return make($BackendSyntax("App", toBackendExpr(v._2), [toBackendExpr(v._3)]));
  }
  if (v.tag === "ExprLet") {
    return foldrArray((bind$p) => (next) => {
      if (bind$p.tag === "NonRec") {
        return makeLet2($Maybe("Just", bind$p._1._2))(toBackendExpr(bind$p._1._3))((v3) => next);
      }
      if (bind$p.tag === "Rec" && bind$p._1.length > 0) {
        const $0 = bind$p._1;
        return (x) => {
          const idents = arrayMap((v4) => v4._2)($0);
          return build(getCtx(x))($BackendSyntax(
            "LetRec",
            x.currentLevel,
            intro(foldableArray)(idents)(x.currentLevel)(traverse3(toBackendBinding)($0))(x),
            intro(foldableArray)(idents)(x.currentLevel)(next)(x)
          ));
        };
      }
      if (bind$p.tag === "Rec") {
        return _crashWith("CoreFn empty Rec binding group");
      }
      fail();
    })(toBackendExpr(v._3))(v._2);
  }
  if (v.tag === "ExprCase") {
    const $0 = v._3;
    return foldrArray((expr) => (next) => (idents) => makeLet2(Nothing)(toBackendExpr(expr))((tmp) => next(snoc(idents)(tmp))))((idents) => foldrArray((v$1) => {
      const $1 = v$1._1;
      const $2 = v$1._2;
      return (mainCb) => (caseRows) => {
        const $3 = zipWithA2((ident) => (b) => {
          const $32 = binderToPattern(b);
          return (x) => ({ column: ident, pattern: $32(x) });
        })(idents)($1);
        return (x) => {
          const $4 = $3(x);
          const args = sortBy(ordString.compare)(foldMap22(patternVars)($4));
          if ($2.tag === "Unconditional") {
            const $5 = $2._1;
            return makeLet2(Nothing)(makeUncurriedAbs(args)((v1) => toBackendExpr($5)))((tmp) => mainCb(snoc(caseRows)({
              patterns: $4,
              guardFn: $CaseRowGuardedExpr("UnconditionalFn", tmp),
              vars: Leaf
            })))(x);
          }
          if ($2.tag === "Guarded") {
            return foldrArray((v1) => {
              const $5 = v1._2;
              const $6 = v1._1;
              return (cb) => (xs) => makeLet2(Nothing)(makeUncurriedAbs(args)((v2) => toBackendExpr($5)))((tmp) => cb(snoc(xs)($Tuple($6, tmp))));
            })((xs) => {
              if (xs.length > 0) {
                return mainCb(snoc(caseRows)({ patterns: $4, guardFn: $CaseRowGuardedExpr("GuardedFn", xs), vars: Leaf }));
              }
              return _crashWith("CoreFn empty Guarded");
            })($2._1)([])(x);
          }
          fail();
        };
      };
    })((caseRows) => buildCaseTreeFromRows(caseRows))($0)([]))(v._2)([]);
  }
  fail();
};
var toBackendBinding = (v) => {
  const $0 = Tuple(v._2);
  const $1 = toBackendExpr(v._3);
  return (x) => $0($1(x));
};
var buildCaseTreeFromRows = (denormalizedRows) => {
  const $0 = normalizeCaseRows(denormalizedRows);
  if ($0.length > 0) {
    const v1 = uncons2($0);
    const $1 = foldableWithIndexArray.foldlWithIndex((idx) => (acc) => (p) => {
      if (!eqPatternCase.eq(p.pattern.patternCase)(PatWild)) {
        return snoc(acc)($Tuple(idx, p));
      }
      return acc;
    })([])(v1.head.patterns);
    if ($1.length > 0) {
      return buildCasePattern(chooseNextPattern($1)(v1.tail))($0);
    }
    return buildCaseLeaf(v1.head)(v1.tail);
  }
  return patternFail;
};
var buildCasePattern = (chosenColumn) => (rows) => {
  const letBindSubterm = (v) => (nextCb) => (idents) => makeLet2(Nothing)(make($BackendSyntax(
    "Accessor",
    make($BackendSyntax("Local", Nothing, chosenColumn.column)),
    v.accessor
  )))((tmp) => nextCb(snoc(idents)(tmp)));
  const $0 = decompose(chosenColumn)(rows).rowsWithMatch;
  const expandSubterms = foldrArray(letBindSubterm)((idents) => buildCaseTreeFromRows((() => {
    const inlineWildSubterms = arrayMap((column) => ({ column, pattern: { vars: Leaf, patternCase: PatWild, subterms: [] } }))(idents);
    return arrayMap((v) => ({
      guardFn: v.guardFn,
      vars: append2(v.vars)(toCaseRowVars(v.match)),
      patterns: [
        ...v.nonMatchesBefore,
        ...v.match.pattern.patternCase.tag === "PatWild" ? inlineWildSubterms : zipWithImpl((column) => (v$1) => ({ column, pattern: v$1.pattern }), idents, v.match.pattern.subterms),
        ...v.nonMatchesAfter
      ]
    }))($0);
  })()))(chosenColumn.pattern.subterms)([]);
  const buildCaseBranch = (guardExpr) => {
    const v = decompose(chosenColumn)(rows);
    const $1 = v.rowsWithMatch;
    return makeGuard(chosenColumn.column)(guardExpr)(foldrArray(letBindSubterm)((idents) => buildCaseTreeFromRows((() => {
      const inlineWildSubterms = arrayMap((column) => ({ column, pattern: { vars: Leaf, patternCase: PatWild, subterms: [] } }))(idents);
      return arrayMap((v$1) => ({
        guardFn: v$1.guardFn,
        vars: append2(v$1.vars)(toCaseRowVars(v$1.match)),
        patterns: [
          ...v$1.nonMatchesBefore,
          ...v$1.match.pattern.patternCase.tag === "PatWild" ? inlineWildSubterms : zipWithImpl((column) => (v$2) => ({ column, pattern: v$2.pattern }), idents, v$1.match.pattern.subterms),
          ...v$1.nonMatchesAfter
        ]
      }))($1);
    })()))(chosenColumn.pattern.subterms)([]))(buildCaseTreeFromRows(v.rowsNoMatch));
  };
  if (chosenColumn.pattern.patternCase.tag === "PatWild") {
    return _crashWith("Impossible: chosen column cannot be wild pattern");
  }
  if (chosenColumn.pattern.patternCase.tag === "PatRecord") {
    return expandSubterms;
  }
  if (chosenColumn.pattern.patternCase.tag === "PatProduct") {
    return expandSubterms;
  }
  if (chosenColumn.pattern.patternCase.tag === "PatSum") {
    const $1 = chosenColumn.pattern.patternCase._2;
    return buildCaseBranch((lhs) => $BackendSyntax(
      "PrimOp",
      $BackendOperator("Op1", $BackendOperator1("OpIsTag", $1), lhs)
    ));
  }
  if (chosenColumn.pattern.patternCase.tag === "PatArray") {
    return buildCaseBranch(guardArrayLength(chosenColumn.pattern.patternCase._1));
  }
  if (chosenColumn.pattern.patternCase.tag === "PatInt") {
    return buildCaseBranch(guardInt(chosenColumn.pattern.patternCase._1));
  }
  if (chosenColumn.pattern.patternCase.tag === "PatNumber") {
    return buildCaseBranch(guardNumber(chosenColumn.pattern.patternCase._1));
  }
  if (chosenColumn.pattern.patternCase.tag === "PatString") {
    return buildCaseBranch(guardString(chosenColumn.pattern.patternCase._1));
  }
  if (chosenColumn.pattern.patternCase.tag === "PatChar") {
    return buildCaseBranch(guardChar(chosenColumn.pattern.patternCase._1));
  }
  if (chosenColumn.pattern.patternCase.tag === "PatBoolean") {
    return buildCaseBranch(guardBoolean(chosenColumn.pattern.patternCase._1));
  }
  fail();
};
var buildCaseLeaf = (row0) => (tailRows) => {
  const orderedArgs = toUnfoldable12(append2(row0.vars)(foldMap42(toCaseRowVars)(row0.patterns)));
  if (row0.guardFn.tag === "UnconditionalFn") {
    return make($BackendSyntax(
      "UncurriedApp",
      make($BackendSyntax("Local", Nothing, row0.guardFn._1)),
      arrayMap((v) => make($BackendSyntax("Local", $Maybe("Just", v._1), v._2)))(orderedArgs)
    ));
  }
  if (row0.guardFn.tag === "GuardedFn") {
    const $0 = row0.guardFn._1;
    return foldrArray((v) => {
      const $1 = v._1;
      const $2 = v._2;
      return (cb) => (args) => makeLet2($Maybe("Just", $1))(make($BackendSyntax("Local", Nothing, $2)))((tmp) => cb(snoc(args)($Tuple(
        $1,
        tmp
      ))));
    })((args) => {
      const $1 = $$for($0)((v) => {
        const $12 = toBackendExpr(v._1);
        const $2 = make($BackendSyntax(
          "UncurriedApp",
          make($BackendSyntax("Local", Nothing, v._2)),
          arrayMap((v$1) => make($BackendSyntax("Local", $Maybe("Just", v$1._1), v$1._2)))(args)
        ));
        return (x) => $Pair($12(x), $2(x));
      });
      return (x) => build(getCtx(x))($BackendSyntax(
        "Branch",
        $1(x),
        buildCaseTreeFromRows(tailRows)(x)
      ));
    })(orderedArgs)([]);
  }
  fail();
};
var toTopLevelBackendBinding = (group2) => (env) => (v) => {
  const qualifiedIdent = $Qualified($Maybe("Just", env.currentModule), v._2);
  const backendExpr = toBackendExpr(v._3)(env);
  const v1 = optimize(member(qualifiedIdent)(env.traceIdents))(getCtx(env))({
    currentModule: env.currentModule,
    evalExternRef: makeExternEvalRef(env),
    evalExternSpine: makeExternEvalSpine(env),
    locals: [],
    directives: env.directives
  })(qualifiedIdent)(env.rewriteLimit)(backendExpr);
  const v2 = toExternImpl(env)(group2)(v1._2);
  return {
    accum: {
      ...env,
      implementations: insert(ordQualified3)(qualifiedIdent)(v2._1)(env.implementations),
      moduleImplementations: insert(ordQualified3)(qualifiedIdent)(v2._1)(env.moduleImplementations),
      optimizationSteps: (() => {
        const $0 = Tuple(qualifiedIdent);
        if (v1._1.length > 0) {
          return snoc(env.optimizationSteps)($0(v1._1));
        }
        return env.optimizationSteps;
      })(),
      directives: (() => {
        const v4 = inferTransitiveDirective(env.directives)(v2._1._2)(backendExpr)(v._3);
        if (v4.tag === "Just") {
          const $0 = v4._1;
          return alter3((v5) => {
            if (v5.tag === "Just") {
              return $Maybe(
                "Just",
                unsafeUnionWith(ordInlineAccessor.compare, $$const, v5._1, $0)
              );
            }
            if (v5.tag === "Nothing") {
              return $Maybe("Just", $0);
            }
            fail();
          })($EvalRef(
            "EvalExtern",
            $Qualified($Maybe("Just", env.currentModule), v._2)
          ))(env.directives);
        }
        if (v4.tag === "Nothing") {
          return env.directives;
        }
        fail();
      })()
    },
    value: $Tuple(v._2, $Tuple(v2._1._1.deps, v2._2))
  };
};
var toBackendTopLevelBindingGroup = (env) => (v) => {
  if (v.tag === "Rec") {
    const $0 = mapAccumL2(toTopLevelBackendBinding(arrayMap((v1) => $Qualified(
      $Maybe("Just", env.currentModule),
      v1._2
    ))(v._1)))(env)(v._1);
    return { ...$0, value: { recursive: true, bindings: $0.value } };
  }
  if (v.tag === "NonRec") {
    const $0 = mapAccumL2(toTopLevelBackendBinding([]))(env)([v._1]);
    return { ...$0, value: { recursive: false, bindings: $0.value } };
  }
  fail();
};
var toBackendTopLevelBindingGroups = (binds) => (env) => {
  const result = mapAccumL2(toBackendTopLevelBindingGroup)(env)(binds);
  return {
    ...result,
    value: arrayMap((as) => ({
      recursive: (() => {
        if (0 < as.length) {
          return as[0].recursive;
        }
        fail();
      })(),
      bindings: arrayBind(as)((v1) => v1.bindings)
    }))(groupBy((x) => (y) => !x.recursive && !y.recursive)(result.value))
  };
};
var toBackendModule = (v) => (env) => {
  const localExports = fromFoldable22(v.exports);
  const isBindingUsed = (deps) => (v1) => member1(v1._1)(localExports) || member($Qualified($Maybe("Just", v.name), v1._1))(deps);
  const directives = parseDirectiveHeader(v.name)(v.comments);
  const dataTypes = fromFoldable3(arrayMap((group2) => $Tuple(
    (() => {
      if (0 < group2.length) {
        return group2[0]._1;
      }
      fail();
    })(),
    {
      constructors: fromFoldable4(mapWithIndexArray((tag) => (v1) => $Tuple(v1._2._1, { fields: v1._2._2, tag }))(group2)),
      size: maximum2(arrayMap((x) => x._2._2.length)(group2))
    }
  ))(groupAllBy((x) => (y) => ordString.compare(x._1)(y._1))(arrayBind(arrayBind(v.decls)((v1) => {
    if (v1.tag === "Rec") {
      return v1._1;
    }
    if (v1.tag === "NonRec") {
      return [v1._1];
    }
    fail();
  }))((v1) => {
    if (v1._3.tag === "ExprConstructor") {
      return [$Tuple(v1._3._2, $Tuple(v1._3._3, v1._3._4))];
    }
    return [];
  }))));
  const moduleBindings = toBackendTopLevelBindingGroups(v.decls)({
    ...env,
    dataTypes,
    directives: (() => {
      const go = (z$p, m$p) => {
        if (m$p.tag === "Leaf") {
          return z$p;
        }
        if (m$p.tag === "Node") {
          return go(
            (() => {
              const $0 = m$p._4;
              return alter3((v2) => {
                if (v2.tag === "Nothing") {
                  return $Maybe("Just", $0);
                }
                if (v2.tag === "Just") {
                  return $Maybe("Just", v2._1);
                }
                fail();
              })(m$p._3)(go(z$p, m$p._5));
            })(),
            m$p._6
          );
        }
        fail();
      };
      return go(
        unsafeUnionWith(ordEvalRef.compare, $$const, directives.locals, env.directives),
        directives.exports
      );
    })(),
    moduleImplementations: Leaf
  });
  const usedBindings = mapAccumR2((deps) => (group2) => {
    if (group2.recursive) {
      if (anyImpl(isBindingUsed(deps), group2.bindings)) {
        return {
          accum: unsafeUnionWith(ordQualified3.compare, $$const, foldMap52((x) => x._2._1)(group2.bindings), deps),
          value: { ...group2, bindings: mapMaybe((x) => x)(arrayMap((x) => $Maybe("Just", $Tuple(x._1, x._2._2)))(group2.bindings)) }
        };
      }
      return { accum: deps, value: { ...group2, bindings: mapMaybe((x) => x)([]) } };
    }
    return {
      accum: mapAccumR2((deps$p) => (v2) => {
        if (isBindingUsed(deps$p)(v2)) {
          return {
            accum: unsafeUnionWith(ordQualified3.compare, $$const, v2._2._1, deps$p),
            value: $Maybe("Just", $Tuple(v2._1, v2._2._2))
          };
        }
        return { accum: deps$p, value: Nothing };
      })(deps)(group2.bindings).accum,
      value: {
        ...group2,
        bindings: mapMaybe((x) => x)(mapAccumR2((deps$p) => (v2) => {
          if (isBindingUsed(deps$p)(v2)) {
            return {
              accum: unsafeUnionWith(ordQualified3.compare, $$const, v2._2._1, deps$p),
              value: $Maybe("Just", $Tuple(v2._1, v2._2._2))
            };
          }
          return { accum: deps$p, value: Nothing };
        })(deps)(group2.bindings).value)
      }
    };
  })(Leaf)(moduleBindings.value);
  return $Tuple(
    moduleBindings.accum.optimizationSteps,
    {
      name: v.name,
      comments: v.comments,
      imports: mapMaybe2(ordString)((qi) => {
        if (qi._1.tag === "Just") {
          if (qi._1._1 !== v.name && qi._1._1 !== "Prim") {
            return $Maybe("Just", qi._1._1);
          }
          return Nothing;
        }
        if (qi._1.tag === "Nothing") {
          return Nothing;
        }
        fail();
      })(usedBindings.accum),
      dataTypes: (() => {
        const $0 = any(isBindingUsed(usedBindings.accum));
        return filterWithKey(ordString)((v$1) => (x) => $0(toUnfoldable12(x.constructors)))(dataTypes);
      })(),
      bindings: usedBindings.value,
      exports: localExports,
      reExports: fromFoldable5(v.reExports),
      implementations: moduleBindings.accum.moduleImplementations,
      directives: directives.exports,
      foreign: fromFoldable22(v.foreign)
    }
  );
};

// output-es/PureScript.Backend.Optimizer.Builder/index.js
var buildModules = (dictMonad) => {
  const Bind1 = dictMonad.Bind1();
  const $$void = Bind1.Apply0().Functor0().map((v) => {
  });
  return (options) => (coreFnModules) => {
    const go = (go$a0$copy) => (go$a1$copy) => {
      let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
      while (go$c) {
        const b = go$a0, v = go$a1;
        if (v.tag === "Nil") {
          go$c = false;
          go$r = b;
          continue;
        }
        if (v.tag === "Cons") {
          go$a0 = b + 1 | 0;
          go$a1 = v._2;
          continue;
        }
        fail();
      }
      return go$r;
    };
    const moduleCount = go(0)(coreFnModules);
    return $$void(foldM(dictMonad)((v) => (coreFnModule) => {
      const $0 = v.directives;
      const $1 = v.implementations;
      const $2 = v.moduleIndex;
      const buildEnv = { implementations: $1, moduleCount, moduleIndex: $2 };
      return Bind1.bind(options.onPrepareModule(buildEnv)(coreFnModule))((v1) => {
        const v2 = toBackendModule(v1)({
          analyzeCustom: options.analyzeCustom,
          currentModule: v1.name,
          currentLevel: 0,
          toLevel: Leaf,
          implementations: $1,
          moduleImplementations: Leaf,
          directives: $0,
          dataTypes: Leaf,
          foreignSemantics: options.foreignSemantics,
          rewriteLimit: 1e4,
          traceIdents: options.traceIdents,
          optimizationSteps: []
        });
        const $3 = v2._2;
        const go$1 = (m$p, z$p) => {
          if (m$p.tag === "Leaf") {
            return z$p;
          }
          if (m$p.tag === "Node") {
            return go$1(m$p._5, insert(ordQualified(ordString))(m$p._3)(m$p._4)(go$1(m$p._6, z$p)));
          }
          fail();
        };
        const newImplementations = go$1($3.implementations, $1);
        return Bind1.bind(options.onCodegenModule({ ...buildEnv, implementations: newImplementations })(v1)($3)(v2._1))(() => dictMonad.Applicative0().pure({
          directives: (() => {
            const go$2 = (m$p, z$p) => {
              if (m$p.tag === "Leaf") {
                return z$p;
              }
              if (m$p.tag === "Node") {
                return go$2(m$p._5, insert(ordEvalRef)(m$p._3)(m$p._4)(go$2(m$p._6, z$p)));
              }
              fail();
            };
            return go$2($3.directives, $0);
          })(),
          implementations: newImplementations,
          moduleIndex: $2 + 1 | 0
        }));
      });
    })({ directives: options.directives, implementations: Leaf, moduleIndex: 0 })(coreFnModules));
  };
};

// output-es/Data.Argonaut.Decode.Decoders/index.js
var traverseWithIndex = /* @__PURE__ */ (() => traversableWithIndexArray.traverseWithIndex(applicativeEither))();
var decodeJArray = (x) => {
  const $0 = _caseJson(
    (v) => Nothing,
    (v) => Nothing,
    (v) => Nothing,
    (v) => Nothing,
    Just,
    (v) => Nothing,
    x
  );
  if ($0.tag === "Nothing") {
    return $Either("Left", $JsonDecodeError("TypeMismatch", "Array"));
  }
  if ($0.tag === "Just") {
    return $Either("Right", $0._1);
  }
  fail();
};
var decodeInt = (a) => {
  const $0 = _caseJson(
    (v) => $Either("Left", $JsonDecodeError("TypeMismatch", "Number")),
    (v) => $Either("Left", $JsonDecodeError("TypeMismatch", "Number")),
    Right,
    (v) => $Either("Left", $JsonDecodeError("TypeMismatch", "Number")),
    (v) => $Either("Left", $JsonDecodeError("TypeMismatch", "Number")),
    (v) => $Either("Left", $JsonDecodeError("TypeMismatch", "Number")),
    a
  );
  if ($0.tag === "Left") {
    return $Either("Left", $0._1);
  }
  if ($0.tag === "Right") {
    const $1 = fromNumber($0._1);
    if ($1.tag === "Nothing") {
      return $Either("Left", $JsonDecodeError("TypeMismatch", "Integer"));
    }
    if ($1.tag === "Just") {
      return $Either("Right", $1._1);
    }
  }
  fail();
};
var decodeArray = (decoder) => {
  const $0 = Named("Array");
  const $1 = traverseWithIndex((i) => {
    const $12 = AtIndex(i);
    return (x) => {
      const $2 = decoder(x);
      if ($2.tag === "Left") {
        return $Either("Left", $12($2._1));
      }
      if ($2.tag === "Right") {
        return $Either("Right", $2._1);
      }
      fail();
    };
  });
  return (a) => {
    const $2 = decodeJArray(a);
    if ($2.tag === "Left") {
      return $Either("Left", $2._1);
    }
    if ($2.tag === "Right") {
      const $3 = $1($2._1);
      if ($3.tag === "Left") {
        return $Either("Left", $0($3._1));
      }
      if ($3.tag === "Right") {
        return $Either("Right", $3._1);
      }
    }
    fail();
  };
};
var decodeTuple = (decoderA) => (decoderB) => (json) => {
  const $0 = decodeArray(Right)(json);
  return (() => {
    if ($0.tag === "Left") {
      const $1 = $0._1;
      return (v) => $Either("Left", $1);
    }
    if ($0.tag === "Right") {
      const $1 = $0._1;
      return (f) => f($1);
    }
    fail();
  })()((v) => {
    if (v.length === 2) {
      const $1 = decoderA(v[0]);
      if ($1.tag === "Left") {
        return $Either("Left", $1._1);
      }
      if ($1.tag === "Right") {
        const $2 = decoderB(v[1]);
        if ($2.tag === "Left") {
          return $Either("Left", $2._1);
        }
        if ($2.tag === "Right") {
          return $Either("Right", $Tuple($1._1, $2._1));
        }
      }
      fail();
    }
    return $Either("Left", $JsonDecodeError("TypeMismatch", "Tuple"));
  });
};

// output-es/PureScript.Backend.Optimizer.CoreFn.Json/index.js
var traverse2 = /* @__PURE__ */ (() => traversableArray.traverse(applicativeEither))();
var getFieldOptional$p = (decode) => (obj) => (prop) => {
  const v = _lookup(Nothing, Just, prop, obj);
  if (v.tag === "Nothing") {
    return $Either("Right", Nothing);
  }
  if (v.tag === "Just") {
    if (_caseJson((v$1) => true, (v$1) => false, (v$1) => false, (v$1) => false, (v$1) => false, (v$1) => false, v._1)) {
      return $Either("Right", Nothing);
    }
    const $0 = decode(v._1);
    if ($0.tag === "Left") {
      return $Either("Left", $0._1);
    }
    if ($0.tag === "Right") {
      return $Either("Right", $Maybe("Just", $0._1));
    }
  }
  fail();
};
var getField = (decode) => (obj) => (prop) => {
  const v = _lookup(Nothing, Just, prop, obj);
  if (v.tag === "Nothing") {
    return $Either("Left", $JsonDecodeError("AtKey", prop, MissingValue));
  }
  if (v.tag === "Just") {
    return decode(v._1);
  }
  fail();
};
var decodeString = (json) => _caseJson(
  (v) => $Either("Left", $JsonDecodeError("TypeMismatch", "String")),
  (v) => $Either("Left", $JsonDecodeError("TypeMismatch", "String")),
  (v) => $Either("Left", $JsonDecodeError("TypeMismatch", "String")),
  Right,
  (v) => $Either("Left", $JsonDecodeError("TypeMismatch", "String")),
  (v) => $Either("Left", $JsonDecodeError("TypeMismatch", "String")),
  json
);
var decodeNumber = (json) => _caseJson(
  (v) => $Either("Left", $JsonDecodeError("TypeMismatch", "Number")),
  (v) => $Either("Left", $JsonDecodeError("TypeMismatch", "Number")),
  Right,
  (v) => $Either("Left", $JsonDecodeError("TypeMismatch", "Number")),
  (v) => $Either("Left", $JsonDecodeError("TypeMismatch", "Number")),
  (v) => $Either("Left", $JsonDecodeError("TypeMismatch", "Number")),
  json
);
var decodeJObject = (json) => _caseJson(
  (v) => $Either("Left", $JsonDecodeError("TypeMismatch", "Object")),
  (v) => $Either("Left", $JsonDecodeError("TypeMismatch", "Object")),
  (v) => $Either("Left", $JsonDecodeError("TypeMismatch", "Object")),
  (v) => $Either("Left", $JsonDecodeError("TypeMismatch", "Object")),
  (v) => $Either("Left", $JsonDecodeError("TypeMismatch", "Object")),
  Right,
  json
);
var decodeJArray2 = (json) => _caseJson(
  (v) => $Either("Left", $JsonDecodeError("TypeMismatch", "Array")),
  (v) => $Either("Left", $JsonDecodeError("TypeMismatch", "Array")),
  (v) => $Either("Left", $JsonDecodeError("TypeMismatch", "Array")),
  (v) => $Either("Left", $JsonDecodeError("TypeMismatch", "Array")),
  Right,
  (v) => $Either("Left", $JsonDecodeError("TypeMismatch", "Array")),
  json
);
var decodeBoolean = (json) => _caseJson(
  (v) => $Either("Left", $JsonDecodeError("TypeMismatch", "Boolean")),
  Right,
  (v) => $Either("Left", $JsonDecodeError("TypeMismatch", "Boolean")),
  (v) => $Either("Left", $JsonDecodeError("TypeMismatch", "Boolean")),
  (v) => $Either("Left", $JsonDecodeError("TypeMismatch", "Boolean")),
  (v) => $Either("Left", $JsonDecodeError("TypeMismatch", "Boolean")),
  json
);
var decodeArray2 = (decoder) => (json) => {
  const v = decodeJArray2(json);
  if (v.tag === "Left") {
    return $Either("Left", v._1);
  }
  if (v.tag === "Right") {
    const $0 = v._1;
    return (() => {
      const out = [];
      let ix = 0;
      let con = true;
      let res = void 0;
      const len = $0.length;
      while (con) {
        const ix$p = ix;
        if (ix$p === len) {
          con = false;
          res = $Either("Right", out);
          continue;
        }
        const v1 = decoder($0[ix$p]);
        if (v1.tag === "Left") {
          con = false;
          res = $Either("Left", $JsonDecodeError("AtIndex", ix$p, v1._1));
          continue;
        }
        if (v1.tag === "Right") {
          out.push(v1._1);
          ix = ix$p + 1 | 0;
          continue;
        }
        fail();
      }
      return res;
    })();
  }
  fail();
};
var decodeModuleName = (x) => {
  const $0 = decodeArray2(decodeString)(x);
  if ($0.tag === "Left") {
    return $Either("Left", $0._1);
  }
  if ($0.tag === "Right") {
    return $Either(
      "Right",
      foldlArray((v) => (v1) => {
        if (v.init) {
          return { init: false, acc: v1 };
        }
        return { init: false, acc: v.acc + "." + v1 };
      })({ init: true, acc: "" })($0._1).acc
    );
  }
  fail();
};
var decodeConstructorType = (json) => {
  const $0 = decodeString(json);
  if ($0.tag === "Left") {
    return $Either("Left", $0._1);
  }
  if ($0.tag === "Right") {
    if ($0._1 === "ProductType") {
      return $Either("Right", ProductType);
    }
    if ($0._1 === "SumType") {
      return $Either("Right", SumType);
    }
    return $Either("Left", $JsonDecodeError("TypeMismatch", "ConstructorType"));
  }
  fail();
};
var decodeImport = (decodeAnn$p) => (json) => {
  const $0 = decodeJObject(json);
  if ($0.tag === "Left") {
    return $Either("Left", $0._1);
  }
  if ($0.tag === "Right") {
    const $1 = getField(decodeAnn$p)($0._1)("annotation");
    if ($1.tag === "Left") {
      return $Either("Left", $1._1);
    }
    if ($1.tag === "Right") {
      const $2 = getField(decodeModuleName)($0._1)("moduleName");
      if ($2.tag === "Left") {
        return $Either("Left", $2._1);
      }
      if ($2.tag === "Right") {
        return $Either("Right", $Import($1._1, $2._1));
      }
    }
  }
  fail();
};
var decodeInt2 = (json) => {
  const $0 = decodeNumber(json);
  if ($0.tag === "Left") {
    return $Either("Left", $0._1);
  }
  if ($0.tag === "Right") {
    const v = fromNumber($0._1);
    if (v.tag === "Nothing") {
      return $Either("Left", $JsonDecodeError("TypeMismatch", "Int"));
    }
    if (v.tag === "Just") {
      return $Either("Right", v._1);
    }
  }
  fail();
};
var decodeCodePoint = (a) => {
  const $0 = decodeInt2(a);
  if ($0.tag === "Left") {
    return $Either("Left", $0._1);
  }
  if ($0.tag === "Right") {
    if ($0._1 >= 0 && $0._1 <= 1114111) {
      return $Either("Right", $0._1);
    }
    return $Either("Left", $JsonDecodeError("TypeMismatch", "CodePoint"));
  }
  fail();
};
var decodeMeta = (json) => {
  const $0 = decodeJObject(json);
  if ($0.tag === "Left") {
    return $Either("Left", $0._1);
  }
  if ($0.tag === "Right") {
    const $1 = getField(decodeString)($0._1)("metaType");
    if ($1.tag === "Left") {
      return $Either("Left", $1._1);
    }
    if ($1.tag === "Right") {
      if ($1._1 === "IsConstructor") {
        const $2 = getField(decodeConstructorType)($0._1)("constructorType");
        if ($2.tag === "Left") {
          return $Either("Left", $2._1);
        }
        if ($2.tag === "Right") {
          const $3 = getField(decodeArray2(decodeString))($0._1)("identifiers");
          if ($3.tag === "Left") {
            return $Either("Left", $3._1);
          }
          if ($3.tag === "Right") {
            return $Either("Right", $Meta("IsConstructor", $2._1, $3._1));
          }
        }
        fail();
      }
      if ($1._1 === "IsNewtype") {
        return $Either("Right", IsNewtype);
      }
      if ($1._1 === "IsTypeClassConstructor") {
        return $Either("Right", IsTypeClassConstructor);
      }
      if ($1._1 === "IsForeign") {
        return $Either("Right", IsForeign);
      }
      if ($1._1 === "IsWhere") {
        return $Either("Right", IsWhere);
      }
      if ($1._1 === "IsSyntheticApp") {
        return $Either("Right", IsSyntheticApp);
      }
      return $Either("Left", $JsonDecodeError("TypeMismatch", "Meta"));
    }
  }
  fail();
};
var decodeAnn = (_path) => (json) => {
  const $0 = decodeJObject(json);
  if ($0.tag === "Left") {
    return $Either("Left", $0._1);
  }
  if ($0.tag === "Right") {
    const $1 = getFieldOptional$p(decodeMeta)($0._1)("meta");
    if ($1.tag === "Left") {
      return $Either("Left", $1._1);
    }
    if ($1.tag === "Right") {
      return $Either("Right", { span: emptySpan, meta: $1._1 });
    }
  }
  fail();
};
var decodeQualified = (k) => (json) => {
  const $0 = decodeJObject(json);
  if ($0.tag === "Left") {
    return $Either("Left", $0._1);
  }
  if ($0.tag === "Right") {
    const $1 = getFieldOptional$p(decodeModuleName)($0._1)("moduleName");
    if ($1.tag === "Left") {
      return $Either("Left", $1._1);
    }
    if ($1.tag === "Right") {
      const $2 = getField(k)($0._1)("identifier");
      if ($2.tag === "Left") {
        return $Either("Left", $2._1);
      }
      if ($2.tag === "Right") {
        return $Either("Right", $Qualified($1._1, $2._1));
      }
    }
  }
  fail();
};
var decodeReExports = (json) => {
  const $0 = decodeJObject(json);
  if ($0.tag === "Left") {
    return $Either("Left", $0._1);
  }
  if ($0.tag === "Right") {
    const $1 = traverse2(traversableTuple.traverse(applicativeEither)(decodeArray2(decodeString)))(toArrayWithKey(Tuple)($0._1));
    if ($1.tag === "Left") {
      return $Either("Left", $1._1);
    }
    if ($1.tag === "Right") {
      return $Either("Right", arrayBind($1._1)((v) => arrayMap(ReExport(v._1))(v._2)));
    }
  }
  fail();
};
var decodeRecord = (x) => decodeArray2((json) => {
  const $0 = decodeJArray2(json);
  if ($0.tag === "Left") {
    return $Either("Left", $0._1);
  }
  if ($0.tag === "Right") {
    if ($0._1.length === 2) {
      const $1 = decodeString($0._1[0]);
      if ($1.tag === "Left") {
        return $Either("Left", $1._1);
      }
      if ($1.tag === "Right") {
        const $2 = x($0._1[1]);
        if ($2.tag === "Left") {
          return $Either("Left", $2._1);
        }
        if ($2.tag === "Right") {
          return $Either("Right", $Prop($1._1, $2._1));
        }
      }
      fail();
    }
    return $Either("Left", $JsonDecodeError("TypeMismatch", "Tuple"));
  }
  fail();
});
var decodeSourcePos = (json) => {
  const $0 = decodeTuple(decodeInt)(decodeInt)(json);
  if ($0.tag === "Left") {
    return $Either("Left", $0._1);
  }
  if ($0.tag === "Right") {
    return $Either("Right", { line: $0._1._1, column: $0._1._2 });
  }
  fail();
};
var decodeSourceSpan = (path) => (json) => {
  const $0 = decodeJObject(json);
  if ($0.tag === "Left") {
    return $Either("Left", $0._1);
  }
  if ($0.tag === "Right") {
    const $1 = getField(decodeSourcePos)($0._1)("start");
    if ($1.tag === "Left") {
      return $Either("Left", $1._1);
    }
    if ($1.tag === "Right") {
      const $2 = getField(decodeSourcePos)($0._1)("end");
      if ($2.tag === "Left") {
        return $Either("Left", $2._1);
      }
      if ($2.tag === "Right") {
        return $Either("Right", { path, start: $1._1, end: $2._1 });
      }
    }
  }
  fail();
};
var decodeComment = (json) => {
  const $0 = decodeJObject(json);
  if ($0.tag === "Left") {
    return $Either("Left", $0._1);
  }
  if ($0.tag === "Right") {
    const $1 = getField(decodeString)($0._1)("LineComment");
    const $2 = (() => {
      if ($1.tag === "Left") {
        return $Either("Left", $1._1);
      }
      if ($1.tag === "Right") {
        return $Either("Right", $Comment("LineComment", $1._1));
      }
      fail();
    })();
    if ($2.tag === "Left") {
      const $3 = getField(decodeString)($0._1)("BlockComment");
      if ($3.tag === "Left") {
        return $Either("Left", $3._1);
      }
      if ($3.tag === "Right") {
        return $Either("Right", $Comment("BlockComment", $3._1));
      }
      fail();
    }
    if ($2.tag === "Right") {
      return $2;
    }
  }
  fail();
};
var decodeStringLiteral = (json) => {
  const $0 = decodeString(json);
  if ($0.tag === "Left") {
    const $1 = decodeArray2(decodeCodePoint)(json);
    const $2 = (() => {
      if ($1.tag === "Left") {
        return $Either("Left", $1._1);
      }
      if ($1.tag === "Right") {
        return $Either("Right", fromCodePointArray($1._1));
      }
      fail();
    })();
    if ($2.tag === "Left") {
      return $Either("Left", $JsonDecodeError("TypeMismatch", "StringLiteral"));
    }
    if ($2.tag === "Right") {
      return $2;
    }
    fail();
  }
  if ($0.tag === "Right") {
    return $0;
  }
  fail();
};
var decodeLiteral = (dec) => (json) => {
  const $0 = decodeJObject(json);
  if ($0.tag === "Left") {
    return $Either("Left", $0._1);
  }
  if ($0.tag === "Right") {
    const $1 = getField(decodeString)($0._1)("literalType");
    if ($1.tag === "Left") {
      return $Either("Left", $1._1);
    }
    if ($1.tag === "Right") {
      if ($1._1 === "IntLiteral") {
        const $2 = getField(decodeInt2)($0._1)("value");
        if ($2.tag === "Left") {
          return $Either("Left", $2._1);
        }
        if ($2.tag === "Right") {
          return $Either("Right", $Literal("LitInt", $2._1));
        }
        fail();
      }
      if ($1._1 === "NumberLiteral") {
        const $2 = getField(decodeNumber)($0._1)("value");
        if ($2.tag === "Left") {
          return $Either("Left", $2._1);
        }
        if ($2.tag === "Right") {
          return $Either("Right", $Literal("LitNumber", $2._1));
        }
        fail();
      }
      if ($1._1 === "StringLiteral") {
        const $2 = getField(decodeStringLiteral)($0._1)("value");
        if ($2.tag === "Left") {
          return $Either("Left", $2._1);
        }
        if ($2.tag === "Right") {
          return $Either("Right", $Literal("LitString", $2._1));
        }
        fail();
      }
      if ($1._1 === "CharLiteral") {
        const $2 = getField(decodeString)($0._1)("value");
        if ($2.tag === "Left") {
          return $Either("Left", $2._1);
        }
        if ($2.tag === "Right") {
          if (length2($2._1) === 1) {
            const $3 = toCharArray($2._1);
            if (0 < $3.length) {
              return $Either("Right", $Literal("LitChar", $3[0]));
            }
          }
          return $Either("Left", $JsonDecodeError("TypeMismatch", "Char"));
        }
        fail();
      }
      if ($1._1 === "BooleanLiteral") {
        const $2 = getField(decodeBoolean)($0._1)("value");
        if ($2.tag === "Left") {
          return $Either("Left", $2._1);
        }
        if ($2.tag === "Right") {
          return $Either("Right", $Literal("LitBoolean", $2._1));
        }
        fail();
      }
      if ($1._1 === "ArrayLiteral") {
        const $2 = getField(decodeArray2(dec))($0._1)("value");
        if ($2.tag === "Left") {
          return $Either("Left", $2._1);
        }
        if ($2.tag === "Right") {
          return $Either("Right", $Literal("LitArray", $2._1));
        }
        fail();
      }
      if ($1._1 === "ObjectLiteral") {
        const $2 = getField(decodeRecord(dec))($0._1)("value");
        if ($2.tag === "Left") {
          return $Either("Left", $2._1);
        }
        if ($2.tag === "Right") {
          return $Either("Right", $Literal("LitRecord", $2._1));
        }
        fail();
      }
      return $Either("Left", $JsonDecodeError("TypeMismatch", "Literal"));
    }
  }
  fail();
};
var decodeBinder = (decAnn) => (json) => {
  const $0 = decodeJObject(json);
  if ($0.tag === "Left") {
    return $Either("Left", $0._1);
  }
  if ($0.tag === "Right") {
    const $1 = getField(decAnn)($0._1)("annotation");
    if ($1.tag === "Left") {
      return $Either("Left", $1._1);
    }
    if ($1.tag === "Right") {
      const $2 = getField(decodeString)($0._1)("binderType");
      if ($2.tag === "Left") {
        return $Either("Left", $2._1);
      }
      if ($2.tag === "Right") {
        if ($2._1 === "NullBinder") {
          return $Either("Right", $Binder("BinderNull", $1._1));
        }
        if ($2._1 === "VarBinder") {
          const $3 = BinderVar($1._1);
          const $4 = getField(decodeString)($0._1)("identifier");
          if ($4.tag === "Left") {
            return $Either("Left", $4._1);
          }
          if ($4.tag === "Right") {
            return $Either("Right", $3($4._1));
          }
          fail();
        }
        if ($2._1 === "LiteralBinder") {
          const $3 = BinderLit($1._1);
          const $4 = getField(decodeLiteral(decodeBinder(decAnn)))($0._1)("literal");
          if ($4.tag === "Left") {
            return $Either("Left", $4._1);
          }
          if ($4.tag === "Right") {
            return $Either("Right", $3($4._1));
          }
          fail();
        }
        if ($2._1 === "ConstructorBinder") {
          const $3 = getField(decodeQualified(decodeString))($0._1)("typeName");
          if ($3.tag === "Left") {
            return $Either("Left", $3._1);
          }
          if ($3.tag === "Right") {
            const $4 = getField(decodeQualified(decodeString))($0._1)("constructorName");
            if ($4.tag === "Left") {
              return $Either("Left", $4._1);
            }
            if ($4.tag === "Right") {
              const $5 = getField(decodeArray2(decodeBinder(decAnn)))($0._1)("binders");
              if ($5.tag === "Left") {
                return $Either("Left", $5._1);
              }
              if ($5.tag === "Right") {
                return $Either("Right", $Binder("BinderConstructor", $1._1, $3._1, $4._1, $5._1));
              }
            }
          }
          fail();
        }
        if ($2._1 === "NamedBinder") {
          const $3 = getField(decodeString)($0._1)("identifier");
          if ($3.tag === "Left") {
            return $Either("Left", $3._1);
          }
          if ($3.tag === "Right") {
            const $4 = getField(decodeBinder(decAnn))($0._1)("binder");
            if ($4.tag === "Left") {
              return $Either("Left", $4._1);
            }
            if ($4.tag === "Right") {
              return $Either("Right", $Binder("BinderNamed", $1._1, $3._1, $4._1));
            }
          }
          fail();
        }
        return $Either("Left", $JsonDecodeError("TypeMismatch", "Binder"));
      }
    }
  }
  fail();
};
var decodeGuard = (decAnn) => (json) => {
  const $0 = decodeJObject(json);
  if ($0.tag === "Left") {
    return $Either("Left", $0._1);
  }
  if ($0.tag === "Right") {
    const $1 = getField(decodeExpr(decAnn))($0._1)("guard");
    if ($1.tag === "Left") {
      return $Either("Left", $1._1);
    }
    if ($1.tag === "Right") {
      const $2 = getField(decodeExpr(decAnn))($0._1)("expression");
      if ($2.tag === "Left") {
        return $Either("Left", $2._1);
      }
      if ($2.tag === "Right") {
        return $Either("Right", $Guard($1._1, $2._1));
      }
    }
  }
  fail();
};
var decodeExpr = (decAnn) => (json) => {
  const $0 = decodeJObject(json);
  if ($0.tag === "Left") {
    return $Either("Left", $0._1);
  }
  if ($0.tag === "Right") {
    const $1 = getField(decAnn)($0._1)("annotation");
    if ($1.tag === "Left") {
      return $Either("Left", $1._1);
    }
    if ($1.tag === "Right") {
      const $2 = getField(decodeString)($0._1)("type");
      if ($2.tag === "Left") {
        return $Either("Left", $2._1);
      }
      if ($2.tag === "Right") {
        if ($2._1 === "Var") {
          const $3 = ExprVar($1._1);
          const $4 = getField(decodeQualified(decodeString))($0._1)("value");
          if ($4.tag === "Left") {
            return $Either("Left", $4._1);
          }
          if ($4.tag === "Right") {
            return $Either("Right", $3($4._1));
          }
          fail();
        }
        if ($2._1 === "Literal") {
          const $3 = ExprLit($1._1);
          const $4 = getField(decodeLiteral(decodeExpr(decAnn)))($0._1)("value");
          if ($4.tag === "Left") {
            return $Either("Left", $4._1);
          }
          if ($4.tag === "Right") {
            return $Either("Right", $3($4._1));
          }
          fail();
        }
        if ($2._1 === "Constructor") {
          const $3 = getField(decodeString)($0._1)("typeName");
          if ($3.tag === "Left") {
            return $Either("Left", $3._1);
          }
          if ($3.tag === "Right") {
            const $4 = getField(decodeString)($0._1)("constructorName");
            if ($4.tag === "Left") {
              return $Either("Left", $4._1);
            }
            if ($4.tag === "Right") {
              const $5 = getField(decodeArray2(decodeString))($0._1)("fieldNames");
              if ($5.tag === "Left") {
                return $Either("Left", $5._1);
              }
              if ($5.tag === "Right") {
                return $Either("Right", $Expr("ExprConstructor", $1._1, $3._1, $4._1, $5._1));
              }
            }
          }
          fail();
        }
        if ($2._1 === "Accessor") {
          const $3 = getField(decodeExpr(decAnn))($0._1)("expression");
          if ($3.tag === "Left") {
            return $Either("Left", $3._1);
          }
          if ($3.tag === "Right") {
            const $4 = getField(decodeString)($0._1)("fieldName");
            if ($4.tag === "Left") {
              return $Either("Left", $4._1);
            }
            if ($4.tag === "Right") {
              return $Either("Right", $Expr("ExprAccessor", $1._1, $3._1, $4._1));
            }
          }
          fail();
        }
        if ($2._1 === "ObjectUpdate") {
          const $3 = getField(decodeExpr(decAnn))($0._1)("expression");
          if ($3.tag === "Left") {
            return $Either("Left", $3._1);
          }
          if ($3.tag === "Right") {
            const $4 = getField(decodeRecord(decodeExpr(decAnn)))($0._1)("updates");
            if ($4.tag === "Left") {
              return $Either("Left", $4._1);
            }
            if ($4.tag === "Right") {
              return $Either("Right", $Expr("ExprUpdate", $1._1, $3._1, $4._1));
            }
          }
          fail();
        }
        if ($2._1 === "Abs") {
          const $3 = getField(decodeString)($0._1)("argument");
          if ($3.tag === "Left") {
            return $Either("Left", $3._1);
          }
          if ($3.tag === "Right") {
            const $4 = getField(decodeExpr(decAnn))($0._1)("body");
            if ($4.tag === "Left") {
              return $Either("Left", $4._1);
            }
            if ($4.tag === "Right") {
              return $Either("Right", $Expr("ExprAbs", $1._1, $3._1, $4._1));
            }
          }
          fail();
        }
        if ($2._1 === "App") {
          const $3 = getField(decodeExpr(decAnn))($0._1)("abstraction");
          if ($3.tag === "Left") {
            return $Either("Left", $3._1);
          }
          if ($3.tag === "Right") {
            const $4 = getField(decodeExpr(decAnn))($0._1)("argument");
            if ($4.tag === "Left") {
              return $Either("Left", $4._1);
            }
            if ($4.tag === "Right") {
              return $Either("Right", $Expr("ExprApp", $1._1, $3._1, $4._1));
            }
          }
          fail();
        }
        if ($2._1 === "Case") {
          const $3 = getField(decodeArray2(decodeExpr(decAnn)))($0._1)("caseExpressions");
          if ($3.tag === "Left") {
            return $Either("Left", $3._1);
          }
          if ($3.tag === "Right") {
            const $4 = getField(decodeArray2(decodeCaseAlternative(decAnn)))($0._1)("caseAlternatives");
            if ($4.tag === "Left") {
              return $Either("Left", $4._1);
            }
            if ($4.tag === "Right") {
              return $Either("Right", $Expr("ExprCase", $1._1, $3._1, $4._1));
            }
          }
          fail();
        }
        if ($2._1 === "Let") {
          const $3 = getField(decodeArray2(decodeBind(decAnn)))($0._1)("binds");
          if ($3.tag === "Left") {
            return $Either("Left", $3._1);
          }
          if ($3.tag === "Right") {
            const $4 = getField(decodeExpr(decAnn))($0._1)("expression");
            if ($4.tag === "Left") {
              return $Either("Left", $4._1);
            }
            if ($4.tag === "Right") {
              return $Either("Right", $Expr("ExprLet", $1._1, $3._1, $4._1));
            }
          }
          fail();
        }
        return $Either("Left", $JsonDecodeError("TypeMismatch", "Expr"));
      }
    }
  }
  fail();
};
var decodeCaseAlternative = (decAnn) => (json) => {
  const $0 = decodeJObject(json);
  if ($0.tag === "Left") {
    return $Either("Left", $0._1);
  }
  if ($0.tag === "Right") {
    const $1 = getField(decodeArray2(decodeBinder(decAnn)))($0._1)("binders");
    if ($1.tag === "Left") {
      return $Either("Left", $1._1);
    }
    if ($1.tag === "Right") {
      const $2 = getField(decodeBoolean)($0._1)("isGuarded");
      if ($2.tag === "Left") {
        return $Either("Left", $2._1);
      }
      if ($2.tag === "Right") {
        if ($2._1) {
          const $32 = getField(decodeArray2(decodeGuard(decAnn)))($0._1)("expressions");
          if ($32.tag === "Left") {
            return $Either("Left", $32._1);
          }
          if ($32.tag === "Right") {
            return $Either(
              "Right",
              $CaseAlternative($1._1, $CaseGuard("Guarded", $32._1))
            );
          }
          fail();
        }
        const $3 = getField(decodeExpr(decAnn))($0._1)("expression");
        if ($3.tag === "Left") {
          return $Either("Left", $3._1);
        }
        if ($3.tag === "Right") {
          return $Either(
            "Right",
            $CaseAlternative($1._1, $CaseGuard("Unconditional", $3._1))
          );
        }
      }
    }
  }
  fail();
};
var decodeBinding = (decAnn) => (obj) => {
  const $0 = getField(decAnn)(obj)("annotation");
  if ($0.tag === "Left") {
    return $Either("Left", $0._1);
  }
  if ($0.tag === "Right") {
    const $1 = getField(decodeString)(obj)("identifier");
    if ($1.tag === "Left") {
      return $Either("Left", $1._1);
    }
    if ($1.tag === "Right") {
      const $2 = getField(decodeExpr(decAnn))(obj)("expression");
      if ($2.tag === "Left") {
        return $Either("Left", $2._1);
      }
      if ($2.tag === "Right") {
        return $Either("Right", $Binding($0._1, $1._1, $2._1));
      }
    }
  }
  fail();
};
var decodeBind = (decAnn) => (json) => {
  const $0 = decodeJObject(json);
  if ($0.tag === "Left") {
    return $Either("Left", $0._1);
  }
  if ($0.tag === "Right") {
    const $1 = getField(decodeString)($0._1)("bindType");
    if ($1.tag === "Left") {
      return $Either("Left", $1._1);
    }
    if ($1.tag === "Right") {
      if ($1._1 === "NonRec") {
        const $2 = decodeBinding(decAnn)($0._1);
        if ($2.tag === "Left") {
          return $Either("Left", $2._1);
        }
        if ($2.tag === "Right") {
          return $Either("Right", $Bind("NonRec", $2._1));
        }
        fail();
      }
      if ($1._1 === "Rec") {
        const $2 = getField(decodeArray2((a) => {
          const $22 = decodeJObject(a);
          if ($22.tag === "Left") {
            return $Either("Left", $22._1);
          }
          if ($22.tag === "Right") {
            return decodeBinding(decAnn)($22._1);
          }
          fail();
        }))($0._1)("binds");
        if ($2.tag === "Left") {
          return $Either("Left", $2._1);
        }
        if ($2.tag === "Right") {
          return $Either("Right", $Bind("Rec", $2._1));
        }
        fail();
      }
      return $Either("Left", $JsonDecodeError("TypeMismatch", "Bind"));
    }
  }
  fail();
};
var decodeModule$p = (decodeAnn$p) => (json) => {
  const $0 = decodeJObject(json);
  if ($0.tag === "Left") {
    return $Either("Left", $0._1);
  }
  if ($0.tag === "Right") {
    const $1 = getField(decodeModuleName)($0._1)("moduleName");
    if ($1.tag === "Left") {
      return $Either("Left", $1._1);
    }
    if ($1.tag === "Right") {
      const $2 = getField(decodeString)($0._1)("modulePath");
      if ($2.tag === "Left") {
        return $Either("Left", $2._1);
      }
      if ($2.tag === "Right") {
        const $3 = getField(decodeSourceSpan($2._1))($0._1)("sourceSpan");
        if ($3.tag === "Left") {
          return $Either("Left", $3._1);
        }
        if ($3.tag === "Right") {
          const $4 = getField(decodeArray2(decodeImport(decodeAnn$p($2._1))))($0._1)("imports");
          if ($4.tag === "Left") {
            return $Either("Left", $4._1);
          }
          if ($4.tag === "Right") {
            const $5 = getField(decodeArray2(decodeString))($0._1)("exports");
            if ($5.tag === "Left") {
              return $Either("Left", $5._1);
            }
            if ($5.tag === "Right") {
              const $6 = getField(decodeReExports)($0._1)("reExports");
              if ($6.tag === "Left") {
                return $Either("Left", $6._1);
              }
              if ($6.tag === "Right") {
                const $7 = getField(decodeArray2(decodeBind(decodeAnn$p($2._1))))($0._1)("decls");
                if ($7.tag === "Left") {
                  return $Either("Left", $7._1);
                }
                if ($7.tag === "Right") {
                  const $8 = getField(decodeArray2(decodeString))($0._1)("foreign");
                  if ($8.tag === "Left") {
                    return $Either("Left", $8._1);
                  }
                  if ($8.tag === "Right") {
                    const $9 = getField(decodeArray2(decodeComment))($0._1)("comments");
                    if ($9.tag === "Left") {
                      return $Either("Left", $9._1);
                    }
                    if ($9.tag === "Right") {
                      return $Either(
                        "Right",
                        { name: $1._1, path: $2._1, span: $3._1, imports: $4._1, exports: $5._1, reExports: $6._1, decls: $7._1, foreign: $8._1, comments: $9._1 }
                      );
                    }
                  }
                }
              }
            }
          }
        }
      }
    }
  }
  fail();
};

// output-es/PureScript.Backend.Optimizer.CoreFn.Sort/index.js
var lookup5 = (k) => {
  const go = (go$a0$copy) => {
    let go$a0 = go$a0$copy, go$c = true, go$r;
    while (go$c) {
      const v = go$a0;
      if (v.tag === "Leaf") {
        go$c = false;
        go$r = Nothing;
        continue;
      }
      if (v.tag === "Node") {
        const v1 = ordString.compare(k)(v._3);
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
          go$r = $Maybe("Just", v._4);
          continue;
        }
      }
      fail();
    }
    return go$r;
  };
  return go;
};
var runSort = /* @__PURE__ */ (() => {
  const go = (go$a0$copy) => (go$a1$copy) => (go$a2$copy) => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$a2 = go$a2$copy, go$c = true, go$r;
    while (go$c) {
      const acc = go$a0, modIndex = go$a1, v = go$a2;
      if (v.tag === "Cons" && v._1.tag === "Left") {
        go$a0 = $List("Cons", v._1._1, acc);
        go$a1 = modIndex;
        go$a2 = v._2;
        continue;
      }
      const v1 = (v2) => {
        if (v.tag === "Cons") {
          go$a0 = acc;
          go$a1 = modIndex;
          go$a2 = v._2;
          return;
        }
        if (v.tag === "Nil") {
          const go$1 = (go$1$a0$copy) => (go$1$a1$copy) => {
            let go$1$a0 = go$1$a0$copy, go$1$a1 = go$1$a1$copy, go$1$c = true, go$1$r;
            while (go$1$c) {
              const v$1 = go$1$a0, v12 = go$1$a1;
              if (v12.tag === "Nil") {
                go$1$c = false;
                go$1$r = v$1;
                continue;
              }
              if (v12.tag === "Cons") {
                go$1$a0 = $List("Cons", v12._1, v$1);
                go$1$a1 = v12._2;
                continue;
              }
              fail();
            }
            return go$1$r;
          };
          go$c = false;
          go$r = go$1(Nil)(acc);
          return;
        }
        fail();
      };
      if (v.tag === "Cons" && v._1.tag === "Right") {
        const $0 = lookup5(v._1._1)(modIndex);
        if ($0.tag === "Just" && !$0._1._1) {
          go$a0 = acc;
          go$a1 = insert(ordString)(v._1._1)($Tuple(true, $0._1._2))(modIndex);
          go$a2 = foldrArray((x) => Cons($Either("Right", x._2)))($List(
            "Cons",
            $Either("Left", $0._1._2),
            v._2
          ))($0._1._2.imports);
          continue;
        }
      }
      v1(true);
    }
    return go$r;
  };
  return go(Nil);
})();
var sortModules = (dictFoldable) => (init) => runSort(dictFoldable.foldr((m) => insert(ordString)(m.name)($Tuple(false, m)))(Leaf)(init))(dictFoldable.foldr((x) => Cons($Either(
  "Right",
  x.name
)))(Nil)(init));

// output-es/Main/index.js
var filterA2 = /* @__PURE__ */ filterA(applicativeAff);
var traverse4 = /* @__PURE__ */ (() => traversableArray.traverse(applicativeAff))();
var fromFoldable6 = /* @__PURE__ */ foldrArray(Cons)(Nil);
var buildModules2 = /* @__PURE__ */ buildModules(monadAff);
var readCoreFnModule = (filePath) => _bind($$try2(toAff1(stat2)(filePath)))((statRes) => {
  if ((() => {
    if (statRes.tag === "Left") {
      return false;
    }
    if (statRes.tag === "Right") {
      return true;
    }
    fail();
  })()) {
    return _bind(toAff2(readTextFile)(UTF8)(filePath))((contents) => {
      const $0 = _jsonParser(Left, Right, contents);
      const v = (() => {
        if ($0.tag === "Left") {
          const $1 = $0._1;
          return (v2) => $Either("Left", $1);
        }
        if ($0.tag === "Right") {
          const $1 = $0._1;
          return (f) => f($1);
        }
        fail();
      })()((x) => {
        const $1 = decodeModule$p(decodeAnn)(x);
        if ($1.tag === "Left") {
          return $Either("Left", printJsonDecodeError($1._1));
        }
        if ($1.tag === "Right") {
          return $Either("Right", $1._1);
        }
        fail();
      });
      if (v.tag === "Left") {
        return _bind(_liftEffect(error2("Failed to decode " + filePath + ": " + v._1)))(() => _pure(Nothing));
      }
      if (v.tag === "Right") {
        return _pure($Maybe("Just", v._1));
      }
      fail();
    });
  }
  return _pure(Nothing);
});
var main = /* @__PURE__ */ (() => {
  const $0 = _makeFiber(
    ffiUtil,
    _bind(toAff1(readdir2)("output"))((files) => _bind(filterA2((f) => _bind(toAff1(stat2)("output/" + f))((stat3) => _pure(isDirectoryImpl(stat3))))(files))((validDirs) => _bind(traverse4((dir) => readCoreFnModule("output/" + dir + "/corefn.json"))(validDirs))((mbModules) => {
      const finalModules = sortModules(foldableList)(fromFoldable6(mapMaybe((x) => x)(mbModules)));
      return _bind($$try2(toAff1(mkdir2)("output/gopurs_runtime")))(() => _bind(toAff3(writeTextFile)(UTF8)("output/gopurs_runtime/runtime.go")(runtimeGoCode))(() => _bind(toAff3(writeTextFile)(UTF8)("output/go.mod")("module gopurs/output\n\ngo 1.22\n"))(() => _bind(buildModules2({
        directives: Leaf,
        analyzeCustom: (v) => (v1) => Nothing,
        foreignSemantics: Leaf,
        traceIdents: Leaf,
        onPrepareModule: (v) => (m) => _pure(m),
        onCodegenModule: (v) => (v1) => (backendMod) => (v2) => toAff3(writeTextFile)(UTF8)("output/" + backendMod.name + "/" + replaceAll(".")("_")(backendMod.name) + ".go")(translate(arrayMap((i) => split(".")(i._2))(v1.imports))(backendMod))
      })(finalModules))(() => toAff3(writeTextFile)(UTF8)("output/main.go")('package main\n\nimport (\n	"gopurs/output/Test1110"\n	"gopurs/output/gopurs_runtime"\n)\n\nfunc main() {\n	gopurs_runtime.Apply(Test1110.Main, gopurs_runtime.Value{})\n}\n')))));
    })))
  );
  return () => {
    const fiber = $0();
    fiber.run();
  };
})();

// <stdin>
main();
