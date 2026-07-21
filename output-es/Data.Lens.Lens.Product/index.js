import * as Data$dLens$dIso$dNewtype from "../Data.Lens.Iso.Newtype/index.js";
const _2 = dictStrong => {
  const $0 = Data$dLens$dIso$dNewtype._Newtype()()(dictStrong.Profunctor0());
  return x => $0(dictStrong.second(x));
};
const _1 = dictStrong => {
  const $0 = Data$dLens$dIso$dNewtype._Newtype()()(dictStrong.Profunctor0());
  return x => $0(dictStrong.first(x));
};
export {_1, _2};
