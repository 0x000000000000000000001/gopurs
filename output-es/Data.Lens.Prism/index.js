// | Prisms are used for selecting cases of a type, most often a sum
// | type. Consider this:
// |
// | ```purescript
// | data Fill -- think of a paint program filling a shape
// |   = NoFill
// |   | Solid Color
// |   | ...
// | ```
// |
// | A prism that focuses on `Solid` fills could be written like this:
// |
// | ```purescript
// | solidFocus :: Prism' Fill Color
// | solidFocus = prism' Solid case _ of
// |   Solid color -> Just color
// |   _ -> Nothing
// | ```
// |
// | ... and used like this:
// |
// | ```purescript
// | preview solidFocus (Solid Color.white) == Just Color.white
// | preview solidFocus NoFill == Nothing
// |
// | is solidFocus (Solid Color.white) == true
// | ```
// |
// | `review` can be used to go from a `Color` to a `Fill`:
// |
// | ```purescript
// | review solidFocus Color.white == Solid Color.white
// | ```
// | 
// | For more information, see `PrismsForSumTypes.purs` in the
// | `examples/src` directory.
// |
// | ---------------
// |
// | A well-behaved `Prism` will follow these laws:
// |
// | **review-preview**: `preview` retrieves what `review` creates. Equationally:
// |   
// | ```purescript
// | review prism >>> preview prism ≡ Just
// | ```
// |
// | An example:
// | 
// | ```purescript
// | Color.white # review solidFocus # preview solidFocus
// |   == Just Color.white
// | ```
// | 
// | **preview-review**: If `preview` retrieves something, `review` can create
// | the original from that something. Equationally:
// | 
// | ```purescript
// | if preview prism s ≡ Just a then review prism a ≡ s
// | ```
// |
// | An example:
// |
// | ```purescript
// | Solid Color.white # preview solidFocus <#> review solidFocus
// |   == Solid Color.white
// | ```
import * as $runtime from "../runtime.js";
import * as Data$dEither from "../Data.Either/index.js";
import * as Data$dLens$dInternal$dMarket from "../Data.Lens.Internal.Market/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dProfunctor from "../Data.Profunctor/index.js";
import * as Unsafe$dCoerce from "../Unsafe.Coerce/index.js";
const identity = x => x;
const withPrism = l => f => {
  const v = l(Data$dLens$dInternal$dMarket.$Market(identity, Data$dEither.Right));
  return f(v._1)(v._2);
};
const review = Unsafe$dCoerce.unsafeCoerce;
const prism = to => fro => dictChoice => {
  const Profunctor0 = dictChoice.Profunctor0();
  return pab => Profunctor0.dimap(fro)(v2 => {
    if (v2.tag === "Left") { return v2._1; }
    if (v2.tag === "Right") { return v2._1; }
    $runtime.fail();
  })(dictChoice.right(Profunctor0.dimap(Data$dProfunctor.identity)(to)(pab)));
};
const prism$p = to => fro => dictChoice => prism(to)(s => {
  const $0 = fro(s);
  if ($0.tag === "Nothing") { return Data$dEither.$Either("Left", s); }
  if ($0.tag === "Just") { return Data$dEither.$Either("Right", $0._1); }
  $runtime.fail();
})(dictChoice);
const nearly = x => f => dictChoice => prism$p(v => x)(x$1 => {
  if (f(x$1)) { return Data$dMaybe.$Maybe("Just", undefined); }
  return Data$dMaybe.Nothing;
})(dictChoice);
const only = dictEq => x => dictChoice => prism$p(v => x)(x$1 => {
  if (dictEq.eq(x$1)(x)) { return Data$dMaybe.$Maybe("Just", undefined); }
  return Data$dMaybe.Nothing;
})(dictChoice);
const matching = l => l(Data$dLens$dInternal$dMarket.$Market(identity, Data$dEither.Right))._2;
const is = dictHeytingAlgebra => {
  const ff = dictHeytingAlgebra.ff;
  const tt = dictHeytingAlgebra.tt;
  return l => {
    const $0 = l(Data$dLens$dInternal$dMarket.$Market(identity, Data$dEither.Right))._2;
    return x => {
      const $1 = $0(x);
      if ($1.tag === "Left") { return ff; }
      if ($1.tag === "Right") { return tt; }
      $runtime.fail();
    };
  };
};
const isn$pt = dictHeytingAlgebra => {
  const is1 = is(dictHeytingAlgebra);
  return l => {
    const $0 = is1(l);
    return x => dictHeytingAlgebra.not($0(x));
  };
};
const clonePrism = l => dictChoice => {
  const v = l(Data$dLens$dInternal$dMarket.$Market(identity, Data$dEither.Right));
  const $0 = v._1;
  const $1 = v._2;
  return p => prism($0)($1)(dictChoice)(p);
};
const below = dictTraversable => {
  const traverse = dictTraversable.traverse(Data$dEither.applicativeEither);
  return k => dictChoice => {
    const v = k(Data$dLens$dInternal$dMarket.$Market(identity, Data$dEither.Right));
    const $0 = v._2;
    return prism(dictTraversable.Functor0().map(v._1))(s => {
      const v$1 = traverse($0)(s);
      if (v$1.tag === "Left") { return Data$dEither.$Either("Left", s); }
      if (v$1.tag === "Right") { return Data$dEither.$Either("Right", v$1._1); }
      $runtime.fail();
    })(dictChoice);
  };
};
export {below, clonePrism, identity, is, isn$pt, matching, nearly, only, prism, prism$p, review, withPrism};
