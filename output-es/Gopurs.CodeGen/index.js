import * as $runtime from "../runtime.js";
import * as Data$dArray from "../Data.Array/index.js";
import * as Data$dFoldable from "../Data.Foldable/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dString$dCommon from "../Data.String.Common/index.js";
const translateExpr = v => {
  if (v.tag === "Var") {
    if (v._1._2 === "log") {
      return "gopurs_runtime.Func(func(x gopurs_runtime.Value) gopurs_runtime.Value { return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value { fmt.Println(x.StrVal); return gopurs_runtime.Value{} }) })";
    }
    return Data$dString$dCommon.replaceAll("$")("_")(v._1._2);
  }
  if (v.tag === "Local") {
    if (v._1.tag === "Just") { return Data$dString$dCommon.replaceAll("$")("_")(v._1._1); }
    if (v._1.tag === "Nothing") { return "_"; }
    return "gopurs_runtime.Value{}";
  }
  if (v.tag === "Lit") {
    if (v._1.tag === "LitString") { return "gopurs_runtime.Str(\"" + v._1._1 + "\")"; }
    return "gopurs_runtime.Value{}";
  }
  if (v.tag === "App") {
    return "gopurs_runtime.Apply(" + translateExpr(v._1) + ", " + translateExpr((() => {
      if (0 < v._2.length) { return v._2[0]; }
      $runtime.fail();
    })()) + ")";
  }
  return "gopurs_runtime.Value{}";
};
const translateBinding = v => {
  const safeName = Data$dString$dCommon.replaceAll("$")("_")(v._1);
  return Data$dMaybe.$Maybe("Just", (safeName === "main" ? "var Main = " : "var " + safeName + " = ") + translateExpr(v._2));
};
const translateBindingGroup = bg => Data$dMaybe.$Maybe("Just", Data$dString$dCommon.joinWith("\n")(Data$dArray.mapMaybe(translateBinding)(bg.bindings)));
const translate = imports => backendMod => "package " + Data$dString$dCommon.replaceAll(".")("_")(backendMod.name) + "\n\nimport (\n\t\"gopurs/output/gopurs_runtime\"\n\t\"fmt\"\n)\n\nvar _ = fmt.Println\nvar _ = gopurs_runtime.TypeInt\n\n" + Data$dString$dCommon.joinWith("\n\n")(Data$dArray.mapMaybe(translateBindingGroup)(Data$dArray.fromFoldableImpl(
  Data$dFoldable.foldrArray,
  backendMod.bindings
)));
export {translate, translateBinding, translateBindingGroup, translateExpr};
