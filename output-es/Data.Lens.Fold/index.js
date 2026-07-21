// | This module defines functions for working with folds.
import * as $runtime from "../runtime.js";
import * as Control$dApply from "../Control.Apply/index.js";
import * as Data$dArray from "../Data.Array/index.js";
import * as Data$dEither from "../Data.Either/index.js";
import * as Data$dList$dTypes from "../Data.List.Types/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dMonoid$dAdditive from "../Data.Monoid.Additive/index.js";
import * as Data$dMonoid$dMultiplicative from "../Data.Monoid.Multiplicative/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
import * as Unsafe$dCoerce from "../Unsafe.Coerce/index.js";
const identity = x => x;
const unfolded = dictMonoid => {
  const mempty = dictMonoid.mempty;
  return f => p => {
    const go = x => {
      const $0 = f(x);
      if ($0.tag === "Nothing") { return mempty; }
      if ($0.tag === "Just") { return dictMonoid.Semigroup0().append(p($0._1._1))(go($0._1._2)); }
      $runtime.fail();
    };
    return go;
  };
};
const replicated = dictMonoid => {
  const mempty1 = dictMonoid.mempty;
  const $0 = dictMonoid.Semigroup0();
  return i => v => {
    const go = v1 => v2 => {
      if (v1 === 0) { return v$1 => mempty1; }
      const $1 = go(v1 - 1 | 0)(v2);
      return x => $0.append(v2(x))($1(x));
    };
    return go(i)(v);
  };
};
const ifoldMapOf = p => f => p(v => f(v._1)(v._2));
const ifoldlOf = p => f => r => {
  const $0 = p(v => {
    const $0 = f(v._1);
    const $1 = v._2;
    return a => $0(a)($1);
  });
  return x => $0(x)(r);
};
const ifoldrOf = p => f => r => {
  const $0 = p(v => f(v._1)(v._2));
  return x => $0(x)(r);
};
const itoListOf = p => {
  const $0 = p(v => {
    const $0 = v._1;
    const $1 = v._2;
    return xs => Data$dList$dTypes.$List("Cons", Data$dTuple.$Tuple($0, $1), xs);
  });
  return x => $0(x)(Data$dList$dTypes.Nil);
};
const itraverseOf_ = dictApplicative => {
  const Apply0 = dictApplicative.Apply0();
  const $$void = Apply0.Functor0().map(v => {});
  return p => f => {
    const $0 = dictApplicative.pure();
    const $1 = p(v => {
      const $1 = v._1;
      const $2 = v._2;
      return fu => Apply0.apply(Apply0.Functor0().map(v$1 => Control$dApply.identity)($$void(f($1)($2))))(fu);
    });
    return x => $1(x)($0);
  };
};
const iforOf_ = dictApplicative => {
  const $0 = itraverseOf_(dictApplicative);
  return x => {
    const $1 = $0(x);
    return b => a => $1(a)(b);
  };
};
const ifindOf = p => f => {
  const $0 = p(v => {
    const $0 = f(v._1)(v._2) ? Data$dMaybe.$Maybe("Just", v._2) : Data$dMaybe.Nothing;
    return v2 => {
      if (v2.tag === "Nothing") { return $0; }
      if (v2.tag === "Just") { return Data$dMaybe.$Maybe("Just", v2._1); }
      $runtime.fail();
    };
  });
  return x => $0(x)(Data$dMaybe.Nothing);
};
const ianyOf = dictHeytingAlgebra => p => f => p(v => f(v._1)(v._2));
const iallOf = dictHeytingAlgebra => p => f => p(v => f(v._1)(v._2));
const folded = dictMonoid => dictFoldable => dictFoldable.foldMap(dictMonoid);
const foldMapOf = Unsafe$dCoerce.unsafeCoerce;
const foldOf = p => p(identity);
const foldlOf = p => f => r => {
  const $0 = p(x => a => f(a)(x));
  return x => $0(x)(r);
};
const foldrOf = p => f => r => {
  const $0 = p(x => f(x));
  return x => $0(x)(r);
};
const maximumOf = dictOrd => p => {
  const $0 = p(x => x$1 => Data$dMaybe.$Maybe(
    "Just",
    (() => {
      if (x$1.tag === "Nothing") { return x; }
      if (x$1.tag === "Just") {
        if (dictOrd.compare(x)(x$1._1) === "GT") { return x; }
        return x$1._1;
      }
      $runtime.fail();
    })()
  ));
  return x => $0(x)(Data$dMaybe.Nothing);
};
const minimumOf = dictOrd => p => {
  const $0 = p(x => x$1 => Data$dMaybe.$Maybe(
    "Just",
    (() => {
      if (x$1.tag === "Nothing") { return x; }
      if (x$1.tag === "Just") {
        if (dictOrd.compare(x)(x$1._1) === "LT") { return x; }
        return x$1._1;
      }
      $runtime.fail();
    })()
  ));
  return x => $0(x)(Data$dMaybe.Nothing);
};
const toListOf = p => {
  const $0 = p(x => Data$dList$dTypes.Cons(x));
  return x => $0(x)(Data$dList$dTypes.Nil);
};
const toArrayOf = p => {
  const $0 = toListOf(p);
  return x => Data$dArray.fromFoldableImpl(Data$dList$dTypes.foldableList.foldr, $0(x));
};
const toArrayOfOn = s => p => Data$dArray.fromFoldableImpl(Data$dList$dTypes.foldableList.foldr, toListOf(p)(s));
const toListOfOn = s => p => toListOf(p)(s);
const traverseOf_ = dictApplicative => {
  const Apply0 = dictApplicative.Apply0();
  const $$void = Apply0.Functor0().map(v => {});
  return p => f => {
    const $0 = dictApplicative.pure();
    const $1 = p(x => fu => Apply0.apply(Apply0.Functor0().map(v => Control$dApply.identity)($$void(f(x))))(fu));
    return x => $1(x)($0);
  };
};
const has = dictHeytingAlgebra => {
  const tt = dictHeytingAlgebra.tt;
  return p => p(v => tt);
};
const hasn$pt = dictHeytingAlgebra => {
  const ff = dictHeytingAlgebra.ff;
  return p => p(v => ff);
};
const lastOf = p => p(x => Data$dMaybe.$Maybe("Just", x));
const lengthOf = p => p(v => 1);
const preview = p => p(x => Data$dMaybe.$Maybe("Just", x));
const previewOn = s => p => p(x => Data$dMaybe.$Maybe("Just", x))(s);
const productOf = dictSemiring => p => p(Data$dMonoid$dMultiplicative.Multiplicative);
const sequenceOf_ = dictApplicative => {
  const $0 = dictApplicative.Apply0();
  return p => {
    const $1 = dictApplicative.pure();
    const $2 = p(f => v => $0.apply($0.Functor0().map(v$1 => Control$dApply.identity)(f))(v));
    return x => $2(x)($1);
  };
};
const sumOf = dictSemiring => p => p(Data$dMonoid$dAdditive.Additive);
const firstOf = p => p(x => Data$dMaybe.$Maybe("Just", x));
const findOf = p => f => {
  const $0 = p(x => {
    const $0 = f(x) ? Data$dMaybe.$Maybe("Just", x) : Data$dMaybe.Nothing;
    return v2 => {
      if (v2.tag === "Nothing") { return $0; }
      if (v2.tag === "Just") { return Data$dMaybe.$Maybe("Just", v2._1); }
      $runtime.fail();
    };
  });
  return x => $0(x)(Data$dMaybe.Nothing);
};
const filtered = dictChoice => f => {
  const $0 = dictChoice.Profunctor0().dimap(x => {
    if (f(x)) { return Data$dEither.$Either("Right", x); }
    return Data$dEither.$Either("Left", x);
  })(v2 => {
    if (v2.tag === "Left") { return v2._1; }
    if (v2.tag === "Right") { return v2._1; }
    $runtime.fail();
  });
  return x => $0(dictChoice.right(x));
};
const anyOf = dictHeytingAlgebra => p => f => p(x => f(x));
const elemOf = dictEq => p => a => p(x => dictEq.eq(x)(a));
const orOf = dictHeytingAlgebra => p => p(x => x);
const allOf = dictHeytingAlgebra => p => f => p(x => f(x));
const andOf = dictHeytingAlgebra => p => p(x => x);
const notElemOf = dictEq => p => a => p(x => !dictEq.eq(x)(a));
export {
  allOf,
  andOf,
  anyOf,
  elemOf,
  filtered,
  findOf,
  firstOf,
  foldMapOf,
  foldOf,
  folded,
  foldlOf,
  foldrOf,
  has,
  hasn$pt,
  iallOf,
  ianyOf,
  identity,
  ifindOf,
  ifoldMapOf,
  ifoldlOf,
  ifoldrOf,
  iforOf_,
  itoListOf,
  itraverseOf_,
  lastOf,
  lengthOf,
  maximumOf,
  minimumOf,
  notElemOf,
  orOf,
  preview,
  previewOn,
  productOf,
  replicated,
  sequenceOf_,
  sumOf,
  toArrayOf,
  toArrayOfOn,
  toListOf,
  toListOfOn,
  traverseOf_,
  unfolded
};
