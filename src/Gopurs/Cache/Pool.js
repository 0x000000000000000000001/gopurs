export const create = () => ({ index: new Map(), nodes: [] });
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
