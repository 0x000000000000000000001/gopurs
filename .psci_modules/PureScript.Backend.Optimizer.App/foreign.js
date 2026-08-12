export const stringify = function(version) {
  return function(obj) {
    return JSON.stringify({ v: version, d: obj });
  };
};

export const parseImpl = function(just) {
  return function(nothing) {
    return function(expectedVersion) {
      return function(str) {
        try {
          const parsed = JSON.parse(str);
          if (parsed && parsed.v === expectedVersion) {
            return just(parsed.d);
          }
          return nothing;
        } catch (e) {
          return nothing;
        }
      };
    };
  };
};
