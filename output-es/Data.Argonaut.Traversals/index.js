import * as $runtime from "../runtime.js";
import * as Data$dArgonaut$dCore from "../Data.Argonaut.Core/index.js";
import * as Data$dEither from "../Data.Either/index.js";
const _JsonString = dictWander => {
  const $0 = dictWander.Choice1();
  const $1 = $0.Profunctor0().dimap(x => {
    if (Data$dArgonaut$dCore._caseJson(v => false, v => false, v => false, v => true, v => false, v => false, x)) { return Data$dEither.$Either("Right", x); }
    return Data$dEither.$Either("Left", x);
  })(v2 => {
    if (v2.tag === "Left") { return v2._1; }
    if (v2.tag === "Right") { return v2._1; }
    $runtime.fail();
  });
  return x => $1($0.right(x));
};
const _JsonObject = dictWander => {
  const $0 = dictWander.Choice1();
  const $1 = $0.Profunctor0().dimap(x => {
    if (Data$dArgonaut$dCore._caseJson(v => false, v => false, v => false, v => false, v => false, v => true, x)) { return Data$dEither.$Either("Right", x); }
    return Data$dEither.$Either("Left", x);
  })(v2 => {
    if (v2.tag === "Left") { return v2._1; }
    if (v2.tag === "Right") { return v2._1; }
    $runtime.fail();
  });
  return x => $1($0.right(x));
};
const _JsonNumber = dictWander => {
  const $0 = dictWander.Choice1();
  const $1 = $0.Profunctor0().dimap(x => {
    if (Data$dArgonaut$dCore._caseJson(v => false, v => false, v => true, v => false, v => false, v => false, x)) { return Data$dEither.$Either("Right", x); }
    return Data$dEither.$Either("Left", x);
  })(v2 => {
    if (v2.tag === "Left") { return v2._1; }
    if (v2.tag === "Right") { return v2._1; }
    $runtime.fail();
  });
  return x => $1($0.right(x));
};
const _JsonNull = dictWander => {
  const $0 = dictWander.Choice1();
  const $1 = $0.Profunctor0().dimap(x => {
    if (Data$dArgonaut$dCore._caseJson(v => true, v => false, v => false, v => false, v => false, v => false, x)) { return Data$dEither.$Either("Right", x); }
    return Data$dEither.$Either("Left", x);
  })(v2 => {
    if (v2.tag === "Left") { return v2._1; }
    if (v2.tag === "Right") { return v2._1; }
    $runtime.fail();
  });
  return x => $1($0.right(x));
};
const _JsonBoolean = dictWander => {
  const $0 = dictWander.Choice1();
  const $1 = $0.Profunctor0().dimap(x => {
    if (Data$dArgonaut$dCore._caseJson(v => false, v => true, v => false, v => false, v => false, v => false, x)) { return Data$dEither.$Either("Right", x); }
    return Data$dEither.$Either("Left", x);
  })(v2 => {
    if (v2.tag === "Left") { return v2._1; }
    if (v2.tag === "Right") { return v2._1; }
    $runtime.fail();
  });
  return x => $1($0.right(x));
};
const _JsonArray = dictWander => {
  const $0 = dictWander.Choice1();
  const $1 = $0.Profunctor0().dimap(x => {
    if (Data$dArgonaut$dCore._caseJson(v => false, v => false, v => false, v => false, v => true, v => false, x)) { return Data$dEither.$Either("Right", x); }
    return Data$dEither.$Either("Left", x);
  })(v2 => {
    if (v2.tag === "Left") { return v2._1; }
    if (v2.tag === "Right") { return v2._1; }
    $runtime.fail();
  });
  return x => $1($0.right(x));
};
export {_JsonArray, _JsonBoolean, _JsonNull, _JsonNumber, _JsonObject, _JsonString};
