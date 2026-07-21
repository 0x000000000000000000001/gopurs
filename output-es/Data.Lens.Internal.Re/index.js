// | This module defines the `Re` profunctor
const Re = x => x;
const profunctorRe = dictProfunctor => (
  {
    dimap: f => g => v => {
      const $0 = dictProfunctor.dimap(g)(f);
      return x => v($0(x));
    }
  }
);
const strongRe = dictStrong => {
  const $0 = dictStrong.Profunctor0();
  const profunctorRe1 = {
    dimap: f => g => v => {
      const $1 = $0.dimap(g)(f);
      return x => v($1(x));
    }
  };
  return {unfirst: v => x => v(dictStrong.first(x)), unsecond: v => x => v(dictStrong.second(x)), Profunctor0: () => profunctorRe1};
};
const newtypeRe = {Coercible0: () => {}};
const costrongRe = dictCostrong => {
  const $0 = dictCostrong.Profunctor0();
  const profunctorRe1 = {
    dimap: f => g => v => {
      const $1 = $0.dimap(g)(f);
      return x => v($1(x));
    }
  };
  return {first: v => x => v(dictCostrong.unfirst(x)), second: v => x => v(dictCostrong.unsecond(x)), Profunctor0: () => profunctorRe1};
};
const cochoiceRe = dictCochoice => {
  const $0 = dictCochoice.Profunctor0();
  const profunctorRe1 = {
    dimap: f => g => v => {
      const $1 = $0.dimap(g)(f);
      return x => v($1(x));
    }
  };
  return {left: v => x => v(dictCochoice.unleft(x)), right: v => x => v(dictCochoice.unright(x)), Profunctor0: () => profunctorRe1};
};
const choiceRe = dictChoice => {
  const $0 = dictChoice.Profunctor0();
  const profunctorRe1 = {
    dimap: f => g => v => {
      const $1 = $0.dimap(g)(f);
      return x => v($1(x));
    }
  };
  return {unleft: v => x => v(dictChoice.left(x)), unright: v => x => v(dictChoice.right(x)), Profunctor0: () => profunctorRe1};
};
export {Re, choiceRe, cochoiceRe, costrongRe, newtypeRe, profunctorRe, strongRe};
