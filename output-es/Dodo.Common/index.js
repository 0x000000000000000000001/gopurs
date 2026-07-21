import * as Dodo$dInternal from "../Dodo.Internal/index.js";
const trailingComma = /* #__PURE__ */ Dodo$dInternal.$Doc(
  "FlexAlt",
  /* #__PURE__ */ Dodo$dInternal.$Doc("Text", 2, ", "),
  /* #__PURE__ */ Dodo$dInternal.$Doc("Append", /* #__PURE__ */ Dodo$dInternal.$Doc("Text", 1, ","), Dodo$dInternal.Break)
);
const pursSquares = x => {
  const $0 = (() => {
    if (x.tag === "Empty") { return Dodo$dInternal.$Doc("Text", 2, "[]"); }
    const $0 = x.tag === "Empty"
      ? Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 2, " ]"), Dodo$dInternal.$Doc("Append", Dodo$dInternal.Break, Dodo$dInternal.$Doc("Text", 1, "]")))
      : Dodo$dInternal.$Doc(
          "Append",
          x,
          Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 2, " ]"), Dodo$dInternal.$Doc("Append", Dodo$dInternal.Break, Dodo$dInternal.$Doc("Text", 1, "]")))
        );
    if ($0.tag === "Empty") { return Dodo$dInternal.$Doc("Text", 2, "[ "); }
    if ($0.tag === "Text") { return Dodo$dInternal.$Doc("Text", 2 + $0._1 | 0, "[ " + $0._2); }
    return Dodo$dInternal.$Doc("Append", Dodo$dInternal.$Doc("Text", 2, "[ "), $0);
  })();
  if ($0.tag === "Empty") { return Dodo$dInternal.Empty; }
  if ($0.tag === "FlexSelect" && $0._2.tag === "Empty" && $0._3.tag === "Empty") { return $0; }
  return Dodo$dInternal.$Doc("FlexSelect", $0, Dodo$dInternal.Empty, Dodo$dInternal.Empty);
};
const pursParensExpr = x => {
  const $0 = x.tag === "Empty" ? Dodo$dInternal.Empty : Dodo$dInternal.$Doc("Indent", x);
  const $1 = (() => {
    if ($0.tag === "Empty") { return Dodo$dInternal.$Doc("Text", 1, ")"); }
    if ($0.tag === "Text") { return Dodo$dInternal.$Doc("Text", $0._1 + 1 | 0, $0._2 + ")"); }
    return Dodo$dInternal.$Doc("Append", $0, Dodo$dInternal.$Doc("Text", 1, ")"));
  })();
  const $2 = (() => {
    if ($1.tag === "Empty") { return Dodo$dInternal.$Doc("Text", 1, "("); }
    if ($1.tag === "Text") { return Dodo$dInternal.$Doc("Text", 1 + $1._1 | 0, "(" + $1._2); }
    return Dodo$dInternal.$Doc("Append", Dodo$dInternal.$Doc("Text", 1, "("), $1);
  })();
  if ($2.tag === "Empty") { return Dodo$dInternal.Empty; }
  if ($2.tag === "FlexSelect" && $2._2.tag === "Empty" && $2._3.tag === "Empty") { return $2; }
  return Dodo$dInternal.$Doc("FlexSelect", $2, Dodo$dInternal.Empty, Dodo$dInternal.Empty);
};
const pursParens = x => {
  const $0 = (() => {
    if (x.tag === "Empty") { return Dodo$dInternal.$Doc("Text", 2, "()"); }
    const $0 = x.tag === "Empty"
      ? Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 1, ")"), Dodo$dInternal.$Doc("Append", Dodo$dInternal.Break, Dodo$dInternal.$Doc("Text", 1, ")")))
      : Dodo$dInternal.$Doc(
          "Append",
          x,
          Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 1, ")"), Dodo$dInternal.$Doc("Append", Dodo$dInternal.Break, Dodo$dInternal.$Doc("Text", 1, ")")))
        );
    if ($0.tag === "Empty") { return Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 1, "("), Dodo$dInternal.$Doc("Text", 2, "( ")); }
    return Dodo$dInternal.$Doc("Append", Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 1, "("), Dodo$dInternal.$Doc("Text", 2, "( ")), $0);
  })();
  if ($0.tag === "Empty") { return Dodo$dInternal.Empty; }
  if ($0.tag === "FlexSelect" && $0._2.tag === "Empty" && $0._3.tag === "Empty") { return $0; }
  return Dodo$dInternal.$Doc("FlexSelect", $0, Dodo$dInternal.Empty, Dodo$dInternal.Empty);
};
const pursCurlies = x => {
  const $0 = (() => {
    if (x.tag === "Empty") { return Dodo$dInternal.$Doc("Text", 2, "{}"); }
    const $0 = x.tag === "Empty"
      ? Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 2, " }"), Dodo$dInternal.$Doc("Append", Dodo$dInternal.Break, Dodo$dInternal.$Doc("Text", 1, "}")))
      : Dodo$dInternal.$Doc(
          "Append",
          x,
          Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 2, " }"), Dodo$dInternal.$Doc("Append", Dodo$dInternal.Break, Dodo$dInternal.$Doc("Text", 1, "}")))
        );
    if ($0.tag === "Empty") { return Dodo$dInternal.$Doc("Text", 2, "{ "); }
    if ($0.tag === "Text") { return Dodo$dInternal.$Doc("Text", 2 + $0._1 | 0, "{ " + $0._2); }
    return Dodo$dInternal.$Doc("Append", Dodo$dInternal.$Doc("Text", 2, "{ "), $0);
  })();
  if ($0.tag === "Empty") { return Dodo$dInternal.Empty; }
  if ($0.tag === "FlexSelect" && $0._2.tag === "Empty" && $0._3.tag === "Empty") { return $0; }
  return Dodo$dInternal.$Doc("FlexSelect", $0, Dodo$dInternal.Empty, Dodo$dInternal.Empty);
};
const leadingComma = /* #__PURE__ */ Dodo$dInternal.$Doc(
  "FlexAlt",
  /* #__PURE__ */ Dodo$dInternal.$Doc("Text", 2, ", "),
  /* #__PURE__ */ Dodo$dInternal.$Doc("Append", Dodo$dInternal.Break, /* #__PURE__ */ Dodo$dInternal.$Doc("Text", 2, ", "))
);
const jsSquares = x => {
  const $0 = x.tag === "Empty" ? Dodo$dInternal.Empty : Dodo$dInternal.$Doc("Indent", x);
  const $1 = (() => {
    if ($0.tag === "Empty") { return Dodo$dInternal.$Doc("Text", 2, "[]"); }
    const $1 = $0.tag === "Empty"
      ? Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 1, "]"), Dodo$dInternal.$Doc("Append", Dodo$dInternal.Break, Dodo$dInternal.$Doc("Text", 1, "]")))
      : Dodo$dInternal.$Doc(
          "Append",
          $0,
          Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 1, "]"), Dodo$dInternal.$Doc("Append", Dodo$dInternal.Break, Dodo$dInternal.$Doc("Text", 1, "]")))
        );
    if ($1.tag === "Empty") {
      return Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 1, "["), Dodo$dInternal.$Doc("Append", Dodo$dInternal.$Doc("Text", 1, "["), Dodo$dInternal.Break));
    }
    return Dodo$dInternal.$Doc(
      "Append",
      Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 1, "["), Dodo$dInternal.$Doc("Append", Dodo$dInternal.$Doc("Text", 1, "["), Dodo$dInternal.Break)),
      $1
    );
  })();
  if ($1.tag === "Empty") { return Dodo$dInternal.Empty; }
  if ($1.tag === "FlexSelect" && $1._2.tag === "Empty" && $1._3.tag === "Empty") { return $1; }
  return Dodo$dInternal.$Doc("FlexSelect", $1, Dodo$dInternal.Empty, Dodo$dInternal.Empty);
};
const jsParens = x => {
  const $0 = x.tag === "Empty" ? Dodo$dInternal.Empty : Dodo$dInternal.$Doc("Indent", x);
  const $1 = (() => {
    if ($0.tag === "Empty") { return Dodo$dInternal.$Doc("Text", 2, "()"); }
    const $1 = $0.tag === "Empty"
      ? Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 1, ")"), Dodo$dInternal.$Doc("Append", Dodo$dInternal.Break, Dodo$dInternal.$Doc("Text", 1, ")")))
      : Dodo$dInternal.$Doc(
          "Append",
          $0,
          Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 1, ")"), Dodo$dInternal.$Doc("Append", Dodo$dInternal.Break, Dodo$dInternal.$Doc("Text", 1, ")")))
        );
    if ($1.tag === "Empty") {
      return Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 1, "("), Dodo$dInternal.$Doc("Append", Dodo$dInternal.$Doc("Text", 1, "("), Dodo$dInternal.Break));
    }
    return Dodo$dInternal.$Doc(
      "Append",
      Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 1, "("), Dodo$dInternal.$Doc("Append", Dodo$dInternal.$Doc("Text", 1, "("), Dodo$dInternal.Break)),
      $1
    );
  })();
  if ($1.tag === "Empty") { return Dodo$dInternal.Empty; }
  if ($1.tag === "FlexSelect" && $1._2.tag === "Empty" && $1._3.tag === "Empty") { return $1; }
  return Dodo$dInternal.$Doc("FlexSelect", $1, Dodo$dInternal.Empty, Dodo$dInternal.Empty);
};
const jsCurlies = x => {
  const $0 = x.tag === "Empty" ? Dodo$dInternal.Empty : Dodo$dInternal.$Doc("Indent", x);
  const $1 = (() => {
    if ($0.tag === "Empty") { return Dodo$dInternal.$Doc("Text", 2, "{}"); }
    const $1 = $0.tag === "Empty"
      ? Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 1, "}"), Dodo$dInternal.$Doc("Append", Dodo$dInternal.Break, Dodo$dInternal.$Doc("Text", 1, "}")))
      : Dodo$dInternal.$Doc(
          "Append",
          $0,
          Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 1, "}"), Dodo$dInternal.$Doc("Append", Dodo$dInternal.Break, Dodo$dInternal.$Doc("Text", 1, "}")))
        );
    if ($1.tag === "Empty") {
      return Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 1, "{"), Dodo$dInternal.$Doc("Append", Dodo$dInternal.$Doc("Text", 1, "{"), Dodo$dInternal.Break));
    }
    return Dodo$dInternal.$Doc(
      "Append",
      Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 1, "{"), Dodo$dInternal.$Doc("Append", Dodo$dInternal.$Doc("Text", 1, "{"), Dodo$dInternal.Break)),
      $1
    );
  })();
  if ($1.tag === "Empty") { return Dodo$dInternal.Empty; }
  if ($1.tag === "FlexSelect" && $1._2.tag === "Empty" && $1._3.tag === "Empty") { return $1; }
  return Dodo$dInternal.$Doc("FlexSelect", $1, Dodo$dInternal.Empty, Dodo$dInternal.Empty);
};
export {jsCurlies, jsParens, jsSquares, leadingComma, pursCurlies, pursParens, pursParensExpr, pursSquares, trailingComma};
