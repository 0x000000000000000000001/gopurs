export const gopursTrace = function() {
  return new Error().stack;
};

const fvsCache = new WeakMap();

export const memoizedFreeVarsImpl = function(calcFreeVars) {
  return function(expr) {
    if (fvsCache.has(expr)) {
      return fvsCache.get(expr);
    }
    const res = calcFreeVars(expr);
    fvsCache.set(expr, res);
    return res;
  };
};
