import * as $runtime from "../runtime.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
import * as Partial from "../Partial/index.js";
const unsafeView = () => s => l => {
  const $0 = l(x => Data$dMaybe.$Maybe("Just", x))(s);
  if ($0.tag === "Nothing") { return Partial._crashWith("unsafeView: Empty fold"); }
  if ($0.tag === "Just") { return $0._1; }
  $runtime.fail();
};
const unsafeIndexedFold = () => s => l => {
  const $0 = l(v => Data$dMaybe.$Maybe("Just", Data$dTuple.$Tuple(v._1, v._2)))(s);
  if ($0.tag === "Nothing") { return Partial._crashWith("unsafeIndexedFold: empty Fold"); }
  if ($0.tag === "Just") { return $0._1; }
  $runtime.fail();
};
export {unsafeIndexedFold, unsafeView};
