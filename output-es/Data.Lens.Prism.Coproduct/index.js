import * as Data$dLens$dIso$dNewtype from "../Data.Lens.Iso.Newtype/index.js";
const _Right = dictChoice => {
  const $0 = Data$dLens$dIso$dNewtype._Newtype()()(dictChoice.Profunctor0());
  return x => $0(dictChoice.right(x));
};
const _Left = dictChoice => {
  const $0 = Data$dLens$dIso$dNewtype._Newtype()()(dictChoice.Profunctor0());
  return x => $0(dictChoice.left(x));
};
export {_Left, _Right};
