const Grating = x => x;
const profunctorGrating = {dimap: f => g => v => d => g(v(k => d(x => k(f(x)))))};
const newtypeGrating = {Coercible0: () => {}};
const closedGrating = {closed: v => f => x => v(k => f(g => k(g(x)))), Profunctor0: () => profunctorGrating};
export {Grating, closedGrating, newtypeGrating, profunctorGrating};
