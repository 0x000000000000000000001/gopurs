import * as $runtime from "../runtime.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
const $Doc = (tag, _1, _2, _3) => ({tag, _1, _2, _3});
const Append = value0 => value1 => $Doc("Append", value0, value1);
const Indent = value0 => $Doc("Indent", value0);
const Align = value0 => value1 => $Doc("Align", value0, value1);
const Annotate = value0 => value1 => $Doc("Annotate", value0, value1);
const FlexSelect = value0 => value1 => value2 => $Doc("FlexSelect", value0, value1, value2);
const FlexAlt = value0 => value1 => $Doc("FlexAlt", value0, value1);
const WithPosition = value0 => $Doc("WithPosition", value0);
const Local = value0 => $Doc("Local", value0);
const Text = value0 => value1 => $Doc("Text", value0, value1);
const Break = /* #__PURE__ */ $Doc("Break");
const Empty = /* #__PURE__ */ $Doc("Empty");
const notEmpty = f => v => {
  if (v.tag === "Empty") { return Empty; }
  return f(v);
};
const isEmpty = v => v.tag === "Empty";
const functorDoc = {
  map: f => m => {
    if (m.tag === "Append") { return $Doc("Append", functorDoc.map(f)(m._1), functorDoc.map(f)(m._2)); }
    if (m.tag === "Indent") { return $Doc("Indent", functorDoc.map(f)(m._1)); }
    if (m.tag === "Align") { return $Doc("Align", m._1, functorDoc.map(f)(m._2)); }
    if (m.tag === "Annotate") { return $Doc("Annotate", f(m._1), functorDoc.map(f)(m._2)); }
    if (m.tag === "FlexSelect") { return $Doc("FlexSelect", functorDoc.map(f)(m._1), functorDoc.map(f)(m._2), functorDoc.map(f)(m._3)); }
    if (m.tag === "FlexAlt") { return $Doc("FlexAlt", functorDoc.map(f)(m._1), functorDoc.map(f)(m._2)); }
    if (m.tag === "WithPosition") {
      return $Doc(
        "WithPosition",
        (() => {
          const $0 = functorDoc.map(f);
          return x => $0(m._1(x));
        })()
      );
    }
    if (m.tag === "Local") {
      return $Doc(
        "Local",
        (() => {
          const $0 = functorDoc.map(f);
          return x => {
            const $1 = m._1(x);
            return Data$dTuple.$Tuple($1._1, $0($1._2));
          };
        })()
      );
    }
    if (m.tag === "Text") { return $Doc("Text", m._1, m._2); }
    if (m.tag === "Break") { return Break; }
    if (m.tag === "Empty") { return Empty; }
    $runtime.fail();
  }
};
const bothNotEmpty = f => v => v1 => {
  if (v.tag === "Empty") { return v1; }
  if (v1.tag === "Empty") { return v; }
  return f(v)(v1);
};
const semigroupDoc = {
  append: v => v1 => {
    if (v.tag === "Empty") { return v1; }
    if (v1.tag === "Empty") { return v; }
    if (v.tag === "Text" && v1.tag === "Text") { return $Doc("Text", v._1 + v1._1 | 0, v._2 + v1._2); }
    return $Doc("Append", v, v1);
  }
};
const monoidDoc = {mempty: Empty, Semigroup0: () => semigroupDoc};
export {$Doc, Align, Annotate, Append, Break, Empty, FlexAlt, FlexSelect, Indent, Local, Text, WithPosition, bothNotEmpty, functorDoc, isEmpty, monoidDoc, notEmpty, semigroupDoc};
