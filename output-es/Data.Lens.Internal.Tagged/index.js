// | This module defines the `Tagged` profunctor
import * as Data$dEither from "../Data.Either/index.js";
const Tagged = x => x;
const taggedProfunctor = {dimap: v => g => v1 => g(v1)};
const taggedCostrong = {unfirst: v => v._1, unsecond: v => v._2, Profunctor0: () => taggedProfunctor};
const taggedClosed = {closed: v => v$1 => v, Profunctor0: () => taggedProfunctor};
const taggedChoice = {left: v => Data$dEither.$Either("Left", v), right: v => Data$dEither.$Either("Right", v), Profunctor0: () => taggedProfunctor};
const newtypeTagged = {Coercible0: () => {}};
const functorTagged = {map: f => m => f(m)};
const foldableTagged = {foldMap: dictMonoid => f => v => f(v), foldr: f => b => v => f(v)(b), foldl: f => b => v => f(b)(v)};
const traversableTagged = {
  sequence: dictApplicative => v => dictApplicative.Apply0().Functor0().map(Tagged)(v),
  traverse: dictApplicative => f => v => dictApplicative.Apply0().Functor0().map(Tagged)(f(v)),
  Functor0: () => functorTagged,
  Foldable1: () => foldableTagged
};
const eqTagged = dictEq => ({eq: x => y => dictEq.eq(x)(y)});
const ordTagged = dictOrd => {
  const $0 = dictOrd.Eq0();
  const eqTagged1 = {eq: x => y => $0.eq(x)(y)};
  return {compare: x => y => dictOrd.compare(x)(y), Eq0: () => eqTagged1};
};
const eq1Tagged = {eq1: dictEq => x => y => dictEq.eq(x)(y)};
const ord1Tagged = {compare1: dictOrd => x => y => dictOrd.compare(x)(y), Eq10: () => eq1Tagged};
export {
  Tagged,
  eq1Tagged,
  eqTagged,
  foldableTagged,
  functorTagged,
  newtypeTagged,
  ord1Tagged,
  ordTagged,
  taggedChoice,
  taggedClosed,
  taggedCostrong,
  taggedProfunctor,
  traversableTagged
};
