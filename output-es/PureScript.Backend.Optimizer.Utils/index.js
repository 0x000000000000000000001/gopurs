import * as $runtime from "../runtime.js";
const foldr1Array = f => g => arr => {
  const go = go$a0$copy => go$a1$copy => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
    while (go$c) {
      const ix = go$a0, acc = go$a1;
      if (ix < 0) {
        go$c = false;
        go$r = acc;
        continue;
      }
      go$a0 = ix - 1 | 0;
      go$a1 = f(arr[ix])(acc);
    }
    return go$r;
  };
  return go(arr.length - 2 | 0)(g((() => {
    const $0 = arr.length - 1 | 0;
    if ($0 >= 0 && $0 < arr.length) { return arr[$0]; }
    $runtime.fail();
  })()));
};
const foldl1Array = f => g => arr => {
  const len = arr.length;
  const go = go$a0$copy => go$a1$copy => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
    while (go$c) {
      const ix = go$a0, acc = go$a1;
      if (ix === len) {
        go$c = false;
        go$r = acc;
        continue;
      }
      go$a0 = ix + 1 | 0;
      go$a1 = f(acc)(arr[ix]);
    }
    return go$r;
  };
  return go(1)(g((() => {
    if (0 < arr.length) { return arr[0]; }
    $runtime.fail();
  })()));
};
export {foldl1Array, foldr1Array};
