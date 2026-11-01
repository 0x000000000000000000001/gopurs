// Handwritten FFI for Gopurs.Cache.Pool; keep in sync with the .purs foreign imports.
export const create = () => ({ index: new Map(), nodes: [], references: [] });
export const intern = pool => text => () => {
  let index = pool.index.get(text);
  if (index === undefined) {
    index = pool.nodes.length;
    pool.index.set(text, index);
    pool.nodes.push(text);
  }
  return index;
};
export const finish = pool => () => pool.nodes.join(",");

export const internRef = pool => text => () => {
  const index = intern(pool)(text)();
  let reference = pool.references[index];
  if (reference === undefined) {
    reference = '["int","' + index + '"]';
    pool.references[index] = reference;
  }
  return reference;
};
