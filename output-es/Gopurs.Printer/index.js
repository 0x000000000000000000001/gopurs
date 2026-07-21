import * as $runtime from "../runtime.js";
import * as Data$dFunctor from "../Data.Functor/index.js";
import * as Data$dShow from "../Data.Show/index.js";
import * as Data$dString$dCommon from "../Data.String.Common/index.js";
const printGoExpr = expr => {
  if (expr.tag === "GoVar") { return expr._1; }
  if (expr.tag === "GoString") { return "\"" + expr._1 + "\""; }
  if (expr.tag === "GoInt") { return Data$dShow.showIntImpl(expr._1); }
  if (expr.tag === "GoCall") { return printGoExpr(expr._1) + "(" + Data$dString$dCommon.joinWith(", ")(Data$dFunctor.arrayMap(printGoExpr)(expr._2)) + ")"; }
  if (expr.tag === "GoSelector") { return printGoExpr(expr._1) + "." + expr._2; }
  if (expr.tag === "GoFunc") {
    return "func(" + Data$dString$dCommon.joinWith(", ")(Data$dFunctor.arrayMap(a => a + " gopurs_runtime.Value")(expr._1)) + ") gopurs_runtime.Value {\n" + printGoExpr(expr._2) + "\n}";
  }
  if (expr.tag === "GoBlock") { return Data$dString$dCommon.joinWith("\n")(Data$dFunctor.arrayMap(printGoExpr)(expr._1)); }
  if (expr.tag === "GoReturn") { return "return " + printGoExpr(expr._1); }
  if (expr.tag === "GoAssign") { return expr._1 + " := " + printGoExpr(expr._2); }
  if (expr.tag === "GoRaw") { return expr._1; }
  $runtime.fail();
};
const printGoDecl = v => "var " + v.identifier + " = " + printGoExpr(v.expression);
const printGoFile = v => "package " + v.packageName + "\n\nimport (\n" + Data$dString$dCommon.joinWith("\n")(Data$dFunctor.arrayMap(i => "\t\"" + i + "\"")(v.imports)) + "\n)\n\n" + Data$dString$dCommon.joinWith("\n")(Data$dFunctor.arrayMap(printGoDecl)(v.decls)) + "\n";
export {printGoDecl, printGoExpr, printGoFile};
