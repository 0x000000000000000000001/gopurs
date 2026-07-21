import * as Data$dLens$dIso from "../Data.Lens.Iso/index.js";
const _Newtype = () => () => dictProfunctor => Data$dLens$dIso.coerced()()(dictProfunctor);
const unto = () => v => dictProfunctor => _Newtype()()(dictProfunctor);
export {_Newtype, unto};
