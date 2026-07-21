import * as $runtime from "../runtime.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
const $TokenList = (tag, _1, _2, _3) => ({tag, _1, _2, _3});
const $UnconsToken = (tag, _1, _2) => ({tag, _1, _2});
const TokenEmpty = /* #__PURE__ */ $TokenList("TokenEmpty");
const TokenCons = value0 => value1 => $TokenList("TokenCons", value0, value1);
const TokenWrap = value0 => value1 => value2 => $TokenList("TokenWrap", value0, value1, value2);
const TokenAppend = value0 => value1 => $TokenList("TokenAppend", value0, value1);
const TokenDefer = value0 => $TokenList("TokenDefer", value0);
const TokenArray = value0 => value1 => value2 => $TokenList("TokenArray", value0, value1, value2);
const UnconsDone = /* #__PURE__ */ $UnconsToken("UnconsDone");
const UnconsMore = value0 => value1 => $UnconsToken("UnconsMore", value0, value1);
const wrap = TokenWrap;
const singleton = a => $TokenList("TokenCons", a, TokenEmpty);
const semigroupTokenList = {
  append: v => v1 => {
    if (v1.tag === "TokenEmpty") { return v; }
    if (v.tag === "TokenEmpty") { return v1; }
    return $TokenList("TokenAppend", v, v1);
  }
};
const $unconuncon = ($unconuncon$b$copy, $unconuncon$a0$copy, $unconuncon$a1$copy, $unconuncon$a2$copy, $unconuncon$a3$copy) => {
  let $unconuncon$b = $unconuncon$b$copy;
  let $unconuncon$a0 = $unconuncon$a0$copy;
  let $unconuncon$a1 = $unconuncon$a1$copy;
  let $unconuncon$a2 = $unconuncon$a2$copy;
  let $unconuncon$a3 = $unconuncon$a3$copy;
  let $unconuncon$c = true;
  let $unconuncon$r;
  while ($unconuncon$c) {
    if ($unconuncon$b === 0) {
      const done = $unconuncon$a0, more = $unconuncon$a1, l = $unconuncon$a2, r = $unconuncon$a3;
      if (l.tag === "TokenEmpty") {
        $unconuncon$b = 1;
        $unconuncon$a0 = done;
        $unconuncon$a1 = more;
        $unconuncon$a2 = r;
        continue;
      }
      if (l.tag === "TokenCons") {
        $unconuncon$c = false;
        $unconuncon$r = more(l._1)((() => {
          if (r.tag === "TokenEmpty") { return l._2; }
          if (l._2.tag === "TokenEmpty") { return r; }
          return $TokenList("TokenAppend", l._2, r);
        })());
        continue;
      }
      if (l.tag === "TokenWrap") {
        $unconuncon$c = false;
        $unconuncon$r = more(l._1)(l._2.tag === "TokenEmpty" ? $TokenList("TokenCons", l._3, r) : $TokenList("TokenAppend", l._2, $TokenList("TokenCons", l._3, r)));
        continue;
      }
      if (l.tag === "TokenAppend") {
        $unconuncon$b = 0;
        $unconuncon$a0 = done;
        $unconuncon$a1 = more;
        $unconuncon$a2 = l._1;
        $unconuncon$a3 = (() => {
          if (r.tag === "TokenEmpty") { return l._2; }
          if (l._2.tag === "TokenEmpty") { return r; }
          return $TokenList("TokenAppend", l._2, r);
        })();
        continue;
      }
      if (l.tag === "TokenDefer") {
        $unconuncon$b = 0;
        $unconuncon$a0 = done;
        $unconuncon$a1 = more;
        $unconuncon$a2 = l._1();
        $unconuncon$a3 = r;
        continue;
      }
      if (l.tag === "TokenArray") {
        $unconuncon$c = false;
        $unconuncon$r = more(l._3[l._1])((() => {
          if (l._1 === l._2) { return r; }
          const $0 = $TokenList("TokenArray", l._1 + 1 | 0, l._2, l._3);
          if (r.tag === "TokenEmpty") { return $0; }
          return $TokenList("TokenAppend", $0, r);
        })());
        continue;
      }
      $runtime.fail();
    }
    if ($unconuncon$b === 1) {
      const done = $unconuncon$a0, more = $unconuncon$a1, v = $unconuncon$a2;
      if (v.tag === "TokenEmpty") {
        $unconuncon$c = false;
        $unconuncon$r = done;
        continue;
      }
      if (v.tag === "TokenCons") {
        $unconuncon$c = false;
        $unconuncon$r = more(v._1)(v._2);
        continue;
      }
      if (v.tag === "TokenWrap") {
        $unconuncon$c = false;
        $unconuncon$r = more(v._1)(v._2.tag === "TokenEmpty"
          ? $TokenList("TokenCons", v._3, TokenEmpty)
          : $TokenList("TokenAppend", v._2, $TokenList("TokenCons", v._3, TokenEmpty)));
        continue;
      }
      if (v.tag === "TokenAppend") {
        $unconuncon$b = 0;
        $unconuncon$a0 = done;
        $unconuncon$a1 = more;
        $unconuncon$a2 = v._1;
        $unconuncon$a3 = v._2;
        continue;
      }
      if (v.tag === "TokenDefer") {
        $unconuncon$b = 1;
        $unconuncon$a0 = done;
        $unconuncon$a1 = more;
        $unconuncon$a2 = v._1();
        continue;
      }
      if (v.tag === "TokenArray") {
        $unconuncon$c = false;
        $unconuncon$r = more(v._3[v._1])(v._1 === v._2 ? TokenEmpty : $TokenList("TokenArray", v._1 + 1 | 0, v._2, v._3));
        continue;
      }
      $runtime.fail();
    }
  }
  return $unconuncon$r;
};
const uncons2 = done => more => l => r => $unconuncon(0, done, more, l, r);
const uncons$p = done => more => v => $unconuncon(1, done, more, v);
const toUnfoldable = dictUnfoldable => dictUnfoldable.unfoldr(uncons$p(Data$dMaybe.Nothing)(a => b => Data$dMaybe.$Maybe("Just", Data$dTuple.$Tuple(a, b))));
const uncons = /* #__PURE__ */ uncons$p(UnconsDone)(UnconsMore);
const toArray = init => {
  const arr = [];
  let cur = init;
  let $$continue = true;
  while ($$continue) {
    const tree = cur;
    const v = uncons$p(UnconsDone)(UnconsMore)(tree);
    if (v.tag === "UnconsDone") {
      $$continue = false;
      continue;
    }
    if (v.tag === "UnconsMore") {
      cur = v._2;
      arr.push(v._1);
      continue;
    }
    $runtime.fail();
  }
  return arr;
};
const monoidTokenList = {mempty: TokenEmpty, Semigroup0: () => semigroupTokenList};
const lazyTokenList = {defer: TokenDefer};
const head = head$a0$copy => {
  let head$a0 = head$a0$copy, head$c = true, head$r;
  while (head$c) {
    const v = head$a0;
    if (v.tag === "TokenEmpty") {
      head$c = false;
      head$r = Data$dMaybe.Nothing;
      continue;
    }
    if (v.tag === "TokenCons") {
      head$c = false;
      head$r = Data$dMaybe.$Maybe("Just", v._1);
      continue;
    }
    if (v.tag === "TokenDefer") {
      head$a0 = v._1();
      continue;
    }
    if (v.tag === "TokenWrap") {
      head$c = false;
      head$r = Data$dMaybe.$Maybe("Just", v._1);
      continue;
    }
    if (v.tag === "TokenAppend") {
      head$a0 = v._1;
      continue;
    }
    if (v.tag === "TokenArray") {
      head$c = false;
      head$r = Data$dMaybe.$Maybe("Just", v._3[v._1]);
      continue;
    }
    $runtime.fail();
  }
  return head$r;
};
const fromArray = arr => {
  const len = arr.length;
  if (len === 0) { return TokenEmpty; }
  return $TokenList("TokenArray", 0, len - 1 | 0, arr);
};
const cons = TokenCons;
export {
  $TokenList,
  $UnconsToken,
  TokenAppend,
  TokenArray,
  TokenCons,
  TokenDefer,
  TokenEmpty,
  TokenWrap,
  UnconsDone,
  UnconsMore,
  cons,
  fromArray,
  head,
  lazyTokenList,
  monoidTokenList,
  semigroupTokenList,
  singleton,
  toArray,
  toUnfoldable,
  uncons,
  uncons$p,
  uncons2,
  wrap
};
