export const memoizeName = sanitize => {
  const names = new Map();
  return name => {
    const cached = names.get(name);
    if (cached !== undefined) return cached;
    const result = sanitize(name);
    names.set(name, result);
    return result;
  };
};
