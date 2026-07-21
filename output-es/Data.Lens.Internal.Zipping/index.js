const Zipping = x => x;
const profunctorZipping = {dimap: f => g => v => a1 => a2 => g(v(f(a1))(f(a2)))};
const newtypeZipping = {Coercible0: () => {}};
const closedZipping = {closed: v => f1 => f2 => x => v(f1(x))(f2(x)), Profunctor0: () => profunctorZipping};
export {Zipping, closedZipping, newtypeZipping, profunctorZipping};
