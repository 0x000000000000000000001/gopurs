import * as $runtime from "../runtime.js";
import * as Data$dEither from "../Data.Either/index.js";
import * as Data$dLens$dPrism from "../Data.Lens.Prism/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
const _Nothing = dictChoice => Data$dLens$dPrism.prism(v => Data$dMaybe.Nothing)(v2 => {
  if (v2.tag === "Nothing") { return Data$dEither.$Either("Right", undefined); }
  if (v2.tag === "Just") { return Data$dEither.$Either("Left", Data$dMaybe.Nothing); }
  $runtime.fail();
})(dictChoice);
const _Just = dictChoice => Data$dLens$dPrism.prism(Data$dMaybe.Just)(v2 => {
  if (v2.tag === "Nothing") { return Data$dEither.$Either("Left", Data$dMaybe.Nothing); }
  if (v2.tag === "Just") { return Data$dEither.$Either("Right", v2._1); }
  $runtime.fail();
})(dictChoice);
export {_Just, _Nothing};
