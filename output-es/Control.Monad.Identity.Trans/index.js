const IdentityT = x => x;
const monadSTIdentityT = dictMonadST => dictMonadST;
const traversableIdentityT = dictTraversable => dictTraversable;
const runIdentityT = v => v;
const plusIdentityT = dictPlus => dictPlus;
const newtypeIdentityT = {Coercible0: () => {}};
const monadWriterIdentityT = dictMonadWriter => dictMonadWriter;
const monadTransIdentityT = {lift: dictMonad => IdentityT};
const monadThrowIdentityT = dictMonadThrow => dictMonadThrow;
const monadTellIdentityT = dictMonadTell => dictMonadTell;
const monadStateIdentityT = dictMonadState => dictMonadState;
const monadRecIdentityT = dictMonadRec => dictMonadRec;
const monadReaderIdentityT = dictMonadReader => dictMonadReader;
const monadPlusIdentityT = dictMonadPlus => dictMonadPlus;
const monadIdentityT = dictMonad => dictMonad;
const monadErrorIdentityT = dictMonadError => dictMonadError;
const monadEffectIdentityT = dictMonadEffect => dictMonadEffect;
const monadContIdentityT = dictMonadCont => dictMonadCont;
const monadAskIdentityT = dictMonadAsk => dictMonadAsk;
const mapIdentityT = f => v => f(v);
const functorIdentityT = dictFunctor => dictFunctor;
const foldableIdentityT = dictFoldable => dictFoldable;
const extendIdentityI = dictExtend => {
  const $0 = dictExtend.Functor0();
  return {extend: f => v => dictExtend.extend(x => f(x))(v), Functor0: () => $0};
};
const eqIdentityT = dictEq1 => dictEq => {
  const eq11 = dictEq1.eq1(dictEq);
  return {eq: x => y => eq11(x)(y)};
};
const ordIdentityT = dictOrd1 => {
  const $0 = dictOrd1.Eq10();
  return dictOrd => {
    const compare11 = dictOrd1.compare1(dictOrd);
    const eq11 = $0.eq1(dictOrd.Eq0());
    const eqIdentityT2 = {eq: x => y => eq11(x)(y)};
    return {compare: x => y => compare11(x)(y), Eq0: () => eqIdentityT2};
  };
};
const eq1IdentityT = dictEq1 => ({eq1: dictEq => dictEq1.eq1(dictEq)});
const ord1IdentityT = dictOrd1 => {
  const $0 = dictOrd1.Eq10();
  const $1 = dictOrd1.Eq10();
  const eq1IdentityT1 = {eq1: dictEq => $1.eq1(dictEq)};
  return {
    compare1: dictOrd => {
      const compare11 = dictOrd1.compare1(dictOrd);
      const eq11 = $0.eq1(dictOrd.Eq0());
      return (() => {
        const eqIdentityT2 = {eq: x => y => eq11(x)(y)};
        return {compare: x => y => compare11(x)(y), Eq0: () => eqIdentityT2};
      })().compare;
    },
    Eq10: () => eq1IdentityT1
  };
};
const comonadIdentityT = dictComonad => {
  const $0 = dictComonad.Extend0();
  const $1 = $0.Functor0();
  const extendIdentityI1 = {extend: f => v => $0.extend(x => f(x))(v), Functor0: () => $1};
  return {extract: x => dictComonad.extract(x), Extend0: () => extendIdentityI1};
};
const bindIdentityT = dictBind => dictBind;
const applyIdentityT = dictApply => dictApply;
const applicativeIdentityT = dictApplicative => dictApplicative;
const alternativeIdentityT = dictAlternative => dictAlternative;
const altIdentityT = dictAlt => dictAlt;
export {
  IdentityT,
  altIdentityT,
  alternativeIdentityT,
  applicativeIdentityT,
  applyIdentityT,
  bindIdentityT,
  comonadIdentityT,
  eq1IdentityT,
  eqIdentityT,
  extendIdentityI,
  foldableIdentityT,
  functorIdentityT,
  mapIdentityT,
  monadAskIdentityT,
  monadContIdentityT,
  monadEffectIdentityT,
  monadErrorIdentityT,
  monadIdentityT,
  monadPlusIdentityT,
  monadReaderIdentityT,
  monadRecIdentityT,
  monadSTIdentityT,
  monadStateIdentityT,
  monadTellIdentityT,
  monadThrowIdentityT,
  monadTransIdentityT,
  monadWriterIdentityT,
  newtypeIdentityT,
  ord1IdentityT,
  ordIdentityT,
  plusIdentityT,
  runIdentityT,
  traversableIdentityT
};
