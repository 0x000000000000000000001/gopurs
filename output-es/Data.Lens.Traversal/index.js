// | `Traversal` is an optic that focuses on zero or more values. An
// | `Array` would be a typical example:
// |
// | ```purescript
// | over    traversed negate [1, 2, 3] == [-1, -2, -3]
// | preview traversed [1, 2, 3] == Just 1
// | firstOf traversed [1, 2, 3] == Just 1  -- same as `preview`
// | lastOf  traversed [1, 2, 3] == Just 3
// | ```
// |
// | `view` might surprise you. It assumes that the wrapped values
// | are a monoid, and `append`s them together:
// |
// | ```purescript
// | view traversed ["D", "a", "w", "n"] == "Dawn"
// | ```
// |
// | Many of the functions you'll use are documented in `Data.Lens.Fold`.
import * as Data$dLens$dIndexed from "../Data.Lens.Indexed/index.js";
import * as Data$dLens$dInternal$dWander from "../Data.Lens.Internal.Wander/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
import * as Unsafe$dCoerce from "../Unsafe.Coerce/index.js";
const identity = x => x;
const traversed = dictTraversable => dictWander => dictWander.wander(dictApplicative => dictTraversable.traverse(dictApplicative));
const traverseOf = Unsafe$dCoerce.unsafeCoerce;
const sequenceOf = t => t(identity);
const itraverseOf = t => x => t(v => x(v._1)(v._2));
const iforOf = x => b => a => x(v => a(v._1)(v._2))(b);
const failover = dictAlternative => {
  const empty = dictAlternative.Plus1().empty;
  return t => f => s => {
    const v = t((() => {
      const $0 = Data$dTuple.Tuple(true);
      return x => $0(f(x));
    })())(s);
    if (v._1) { return dictAlternative.Applicative0().pure(v._2); }
    return empty;
  };
};
const elementsOf = dictWander => tr => pr => dictWander.wander(dictApplicative => {
  const tr1 = tr(Data$dLens$dInternal$dWander.wanderStar(dictApplicative));
  return f => s => tr1(v => {
    if (pr(v._1)) { return f(Data$dTuple.$Tuple(v._1, v._2)); }
    return dictApplicative.pure(v._2);
  })(s);
});
const element = dictWander => {
  const unIndex = Data$dLens$dIndexed.unIndex(dictWander.Choice1().Profunctor0());
  return n => tr => unIndex(elementsOf(dictWander)(dictWander1 => Data$dLens$dIndexed.positions(dictWander2 => tr(dictWander2))(dictWander1))(v => v === n));
};
const cloneTraversal = l => dictWander => dictWander.wander(dictApplicative => l(dictApplicative1 => identity)(dictApplicative));
const both = dictBitraversable => dictWander => dictWander.wander(dictApplicative => {
  const $0 = dictBitraversable.bitraverse(dictApplicative);
  return x => $0(x)(x);
});
export {both, cloneTraversal, element, elementsOf, failover, identity, iforOf, itraverseOf, sequenceOf, traverseOf, traversed};
