export const create = () => new Map();
export const lookup = memo => value => () => memo.get(value) ?? -1;
export const insert = memo => value => index => () => { memo.set(value, index); };
