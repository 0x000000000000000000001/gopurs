// | This module defines functions for working with setters.
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
const over = l => l;
const $$set = l => b => l(v => b);
const setJust = p => x => p(v => Data$dMaybe.$Maybe("Just", x));
const subOver = dictRing => p => x => p(a => dictRing.sub(a)(x));
const mulOver = dictSemiring => p => x => p(a => dictSemiring.mul(a)(x));
const modifying = dictMonadState => {
  const $$void = dictMonadState.Monad0().Bind1().Apply0().Functor0().map(v => {});
  return p => f => $$void((() => {
    const $0 = p(f);
    return dictMonadState.state(s => {
      const s$p = $0(s);
      return Data$dTuple.$Tuple(s$p, s$p);
    });
  })());
};
const mulModifying = dictMonadState => {
  const $$void = dictMonadState.Monad0().Bind1().Apply0().Functor0().map(v => {});
  return dictSemiring => p => x => $$void((() => {
    const $0 = p(a => dictSemiring.mul(a)(x));
    return dictMonadState.state(s => {
      const s$p = $0(s);
      return Data$dTuple.$Tuple(s$p, s$p);
    });
  })());
};
const subModifying = dictMonadState => {
  const $$void = dictMonadState.Monad0().Bind1().Apply0().Functor0().map(v => {});
  return dictRing => p => x => $$void((() => {
    const $0 = p(a => dictRing.sub(a)(x));
    return dictMonadState.state(s => {
      const s$p = $0(s);
      return Data$dTuple.$Tuple(s$p, s$p);
    });
  })());
};
const iover = l => f => l(v => f(v._1)(v._2));
const divOver = dictEuclideanRing => p => x => p(a => dictEuclideanRing.div(a)(x));
const divModifying = dictMonadState => {
  const $$void = dictMonadState.Monad0().Bind1().Apply0().Functor0().map(v => {});
  return dictEuclideanRing => p => x => $$void((() => {
    const $0 = p(a => dictEuclideanRing.div(a)(x));
    return dictMonadState.state(s => {
      const s$p = $0(s);
      return Data$dTuple.$Tuple(s$p, s$p);
    });
  })());
};
const disjOver = dictHeytingAlgebra => p => x => p(a => dictHeytingAlgebra.disj(a)(x));
const disjModifying = dictMonadState => {
  const $$void = dictMonadState.Monad0().Bind1().Apply0().Functor0().map(v => {});
  return dictHeytingAlgebra => p => x => $$void((() => {
    const $0 = p(a => dictHeytingAlgebra.disj(a)(x));
    return dictMonadState.state(s => {
      const s$p = $0(s);
      return Data$dTuple.$Tuple(s$p, s$p);
    });
  })());
};
const conjOver = dictHeytingAlgebra => p => x => p(a => dictHeytingAlgebra.conj(a)(x));
const conjModifying = dictMonadState => {
  const $$void = dictMonadState.Monad0().Bind1().Apply0().Functor0().map(v => {});
  return dictHeytingAlgebra => p => x => $$void((() => {
    const $0 = p(a => dictHeytingAlgebra.conj(a)(x));
    return dictMonadState.state(s => {
      const s$p = $0(s);
      return Data$dTuple.$Tuple(s$p, s$p);
    });
  })());
};
const assign = dictMonadState => {
  const $$void = dictMonadState.Monad0().Bind1().Apply0().Functor0().map(v => {});
  return p => b => $$void((() => {
    const $0 = p(v => b);
    return dictMonadState.state(s => {
      const s$p = $0(s);
      return Data$dTuple.$Tuple(s$p, s$p);
    });
  })());
};
const assignJust = dictMonadState => {
  const $$void = dictMonadState.Monad0().Bind1().Apply0().Functor0().map(v => {});
  return p => x => $$void((() => {
    const $0 = p(v => Data$dMaybe.$Maybe("Just", x));
    return dictMonadState.state(s => {
      const s$p = $0(s);
      return Data$dTuple.$Tuple(s$p, s$p);
    });
  })());
};
const appendOver = dictSemigroup => p => x => p(a => dictSemigroup.append(a)(x));
const appendModifying = dictMonadState => {
  const $$void = dictMonadState.Monad0().Bind1().Apply0().Functor0().map(v => {});
  return dictSemigroup => p => x => $$void((() => {
    const $0 = p(a => dictSemigroup.append(a)(x));
    return dictMonadState.state(s => {
      const s$p = $0(s);
      return Data$dTuple.$Tuple(s$p, s$p);
    });
  })());
};
const addOver = dictSemiring => p => x => p(dictSemiring.add(x));
const addModifying = dictMonadState => {
  const $$void = dictMonadState.Monad0().Bind1().Apply0().Functor0().map(v => {});
  return dictSemiring => p => x => $$void((() => {
    const $0 = p(dictSemiring.add(x));
    return dictMonadState.state(s => {
      const s$p = $0(s);
      return Data$dTuple.$Tuple(s$p, s$p);
    });
  })());
};
export {
  addModifying,
  addOver,
  appendModifying,
  appendOver,
  assign,
  assignJust,
  conjModifying,
  conjOver,
  disjModifying,
  disjOver,
  divModifying,
  divOver,
  iover,
  modifying,
  mulModifying,
  mulOver,
  over,
  $$set as set,
  setJust,
  subModifying,
  subOver
};
