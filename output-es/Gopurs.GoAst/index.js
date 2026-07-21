import * as Data$dEq from "../Data.Eq/index.js";
const $GoExpr = (tag, _1, _2) => ({tag, _1, _2});
const eq2 = /* #__PURE__ */ Data$dEq.eqArrayImpl(Data$dEq.eqStringImpl);
const GoVar = value0 => $GoExpr("GoVar", value0);
const GoString = value0 => $GoExpr("GoString", value0);
const GoInt = value0 => $GoExpr("GoInt", value0);
const GoCall = value0 => value1 => $GoExpr("GoCall", value0, value1);
const GoSelector = value0 => value1 => $GoExpr("GoSelector", value0, value1);
const GoFunc = value0 => value1 => $GoExpr("GoFunc", value0, value1);
const GoBlock = value0 => $GoExpr("GoBlock", value0);
const GoReturn = value0 => $GoExpr("GoReturn", value0);
const GoAssign = value0 => value1 => $GoExpr("GoAssign", value0, value1);
const GoRaw = value0 => $GoExpr("GoRaw", value0);
const eqGoExpr = {
  eq: x => y => {
    if (x.tag === "GoVar") { return y.tag === "GoVar" && x._1 === y._1; }
    if (x.tag === "GoString") { return y.tag === "GoString" && x._1 === y._1; }
    if (x.tag === "GoInt") { return y.tag === "GoInt" && x._1 === y._1; }
    if (x.tag === "GoCall") { return y.tag === "GoCall" && eqGoExpr.eq(x._1)(y._1) && Data$dEq.eqArrayImpl(eqGoExpr.eq)(x._2)(y._2); }
    if (x.tag === "GoSelector") { return y.tag === "GoSelector" && eqGoExpr.eq(x._1)(y._1) && x._2 === y._2; }
    if (x.tag === "GoFunc") { return y.tag === "GoFunc" && eq2(x._1)(y._1) && eqGoExpr.eq(x._2)(y._2); }
    if (x.tag === "GoBlock") { return y.tag === "GoBlock" && Data$dEq.eqArrayImpl(eqGoExpr.eq)(x._1)(y._1); }
    if (x.tag === "GoReturn") { return y.tag === "GoReturn" && eqGoExpr.eq(x._1)(y._1); }
    if (x.tag === "GoAssign") { return y.tag === "GoAssign" && x._1 === y._1 && eqGoExpr.eq(x._2)(y._2); }
    return x.tag === "GoRaw" && y.tag === "GoRaw" && x._1 === y._1;
  }
};
export {$GoExpr, GoAssign, GoBlock, GoCall, GoFunc, GoInt, GoRaw, GoReturn, GoSelector, GoString, GoVar, eq2, eqGoExpr};
