import * as $runtime from "../runtime.js";
import * as Data$dArray from "../Data.Array/index.js";
import * as Data$dList$dTypes from "../Data.List.Types/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dOrdering from "../Data.Ordering/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
import * as PureScript$dCST$dTypes from "../PureScript.CST.Types/index.js";
const $LayoutDelim = tag => tag;
const LytRoot = /* #__PURE__ */ $LayoutDelim("LytRoot");
const LytTopDecl = /* #__PURE__ */ $LayoutDelim("LytTopDecl");
const LytTopDeclHead = /* #__PURE__ */ $LayoutDelim("LytTopDeclHead");
const LytDeclGuard = /* #__PURE__ */ $LayoutDelim("LytDeclGuard");
const LytCase = /* #__PURE__ */ $LayoutDelim("LytCase");
const LytCaseBinders = /* #__PURE__ */ $LayoutDelim("LytCaseBinders");
const LytCaseGuard = /* #__PURE__ */ $LayoutDelim("LytCaseGuard");
const LytLambdaBinders = /* #__PURE__ */ $LayoutDelim("LytLambdaBinders");
const LytParen = /* #__PURE__ */ $LayoutDelim("LytParen");
const LytBrace = /* #__PURE__ */ $LayoutDelim("LytBrace");
const LytSquare = /* #__PURE__ */ $LayoutDelim("LytSquare");
const LytIf = /* #__PURE__ */ $LayoutDelim("LytIf");
const LytThen = /* #__PURE__ */ $LayoutDelim("LytThen");
const LytProperty = /* #__PURE__ */ $LayoutDelim("LytProperty");
const LytForall = /* #__PURE__ */ $LayoutDelim("LytForall");
const LytTick = /* #__PURE__ */ $LayoutDelim("LytTick");
const LytLet = /* #__PURE__ */ $LayoutDelim("LytLet");
const LytLetStmt = /* #__PURE__ */ $LayoutDelim("LytLetStmt");
const LytWhere = /* #__PURE__ */ $LayoutDelim("LytWhere");
const LytOf = /* #__PURE__ */ $LayoutDelim("LytOf");
const LytDo = /* #__PURE__ */ $LayoutDelim("LytDo");
const LytAdo = /* #__PURE__ */ $LayoutDelim("LytAdo");
const lytToken = pos => value => ({range: {start: pos, end: pos}, leadingComments: [], trailingComments: [], value});
const isIndented = v => v === "LytLet" || v === "LytLetStmt" || v === "LytWhere" || v === "LytOf" || v === "LytDo" || v === "LytAdo";
const eqLayoutDelim = {
  eq: x => y => {
    if (x === "LytRoot") { return y === "LytRoot"; }
    if (x === "LytTopDecl") { return y === "LytTopDecl"; }
    if (x === "LytTopDeclHead") { return y === "LytTopDeclHead"; }
    if (x === "LytDeclGuard") { return y === "LytDeclGuard"; }
    if (x === "LytCase") { return y === "LytCase"; }
    if (x === "LytCaseBinders") { return y === "LytCaseBinders"; }
    if (x === "LytCaseGuard") { return y === "LytCaseGuard"; }
    if (x === "LytLambdaBinders") { return y === "LytLambdaBinders"; }
    if (x === "LytParen") { return y === "LytParen"; }
    if (x === "LytBrace") { return y === "LytBrace"; }
    if (x === "LytSquare") { return y === "LytSquare"; }
    if (x === "LytIf") { return y === "LytIf"; }
    if (x === "LytThen") { return y === "LytThen"; }
    if (x === "LytProperty") { return y === "LytProperty"; }
    if (x === "LytForall") { return y === "LytForall"; }
    if (x === "LytTick") { return y === "LytTick"; }
    if (x === "LytLet") { return y === "LytLet"; }
    if (x === "LytLetStmt") { return y === "LytLetStmt"; }
    if (x === "LytWhere") { return y === "LytWhere"; }
    if (x === "LytOf") { return y === "LytOf"; }
    if (x === "LytDo") { return y === "LytDo"; }
    return x === "LytAdo" && y === "LytAdo";
  }
};
const insertLayout = v => nextPos => stack => {
  const tokPos = v.range.start;
  const insertStart = lyt => v1 => {
    const go = go$a0$copy => go$a1$copy => {
      let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
      while (go$c) {
        const b = go$a0, v$1 = go$a1;
        if (v$1.tag === "Nil") {
          go$c = false;
          go$r = b;
          continue;
        }
        if (v$1.tag === "Cons") {
          go$a0 = b.tag === "Nothing" && (
            v$1._1._2 === "LytLet" || v$1._1._2 === "LytLetStmt" || v$1._1._2 === "LytWhere" || v$1._1._2 === "LytOf" || v$1._1._2 === "LytDo" || v$1._1._2 === "LytAdo"
          )
            ? Data$dMaybe.$Maybe("Just", v$1._1)
            : b;
          go$a1 = v$1._2;
          continue;
        }
        $runtime.fail();
      }
      return go$r;
    };
    const v2 = go(Data$dMaybe.Nothing)(v1._1);
    if (v2.tag === "Just" && nextPos.column <= v2._1._1.column) { return v1; }
    const $0 = Data$dList$dTypes.$List("Cons", Data$dTuple.$Tuple(nextPos, lyt), v1._1);
    return Data$dTuple.$Tuple(
      $0,
      Data$dArray.snoc(v1._2)(Data$dTuple.$Tuple(
        {range: {start: nextPos, end: nextPos}, leadingComments: [], trailingComments: [], value: PureScript$dCST$dTypes.$Token("TokLayoutStart", nextPos.column)},
        $0
      ))
    );
  };
  const insertSep = v1 => {
    const sepTok = {range: {start: tokPos, end: tokPos}, leadingComments: [], trailingComments: [], value: PureScript$dCST$dTypes.$Token("TokLayoutSep", tokPos.column)};
    const $0 = (lyt, lytPos) => {
      if (lyt === "LytOf") {
        return Data$dTuple.$Tuple(Data$dList$dTypes.$List("Cons", Data$dTuple.$Tuple(tokPos, LytCaseBinders), v1._1), Data$dArray.snoc(v1._2)(Data$dTuple.$Tuple(sepTok, v1._1)));
      }
      return Data$dTuple.$Tuple(v1._1, Data$dArray.snoc(v1._2)(Data$dTuple.$Tuple(sepTok, v1._1)));
    };
    if (v1._1.tag === "Cons") {
      if (v1._1._1._2 === "LytTopDecl") {
        if (tokPos.column === v1._1._1._1.column && tokPos.line !== v1._1._1._1.line) {
          return Data$dTuple.$Tuple(v1._1._2, Data$dArray.snoc(v1._2)(Data$dTuple.$Tuple(sepTok, v1._1._2)));
        }
        if (
          (v1._1._1._2 === "LytLet" || v1._1._1._2 === "LytLetStmt" || v1._1._1._2 === "LytWhere" || v1._1._1._2 === "LytOf" || v1._1._1._2 === "LytDo" || v1._1._1._2 === "LytAdo") && tokPos.column === v1._1._1._1.column && tokPos.line !== v1._1._1._1.line
        ) {
          return $0(v1._1._1._2, v1._1._1._1);
        }
        return v1;
      }
      if (v1._1._1._2 === "LytTopDeclHead" && tokPos.column === v1._1._1._1.column && tokPos.line !== v1._1._1._1.line) {
        return Data$dTuple.$Tuple(v1._1._2, Data$dArray.snoc(v1._2)(Data$dTuple.$Tuple(sepTok, v1._1._2)));
      }
      if (
        (v1._1._1._2 === "LytLet" || v1._1._1._2 === "LytLetStmt" || v1._1._1._2 === "LytWhere" || v1._1._1._2 === "LytOf" || v1._1._1._2 === "LytDo" || v1._1._1._2 === "LytAdo") && tokPos.column === v1._1._1._1.column && tokPos.line !== v1._1._1._1.line
      ) {
        return $0(v1._1._1._2, v1._1._1._1);
      }
    }
    return v1;
  };
  const collapse = p => {
    const go = go$a0$copy => go$a1$copy => {
      let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
      while (go$c) {
        const v1 = go$a0, v2 = go$a1;
        if (v1.tag === "Cons" && p(v1._1._1)(v1._1._2)) {
          go$a0 = v1._2;
          go$a1 = v1._1._2 === "LytLet" || v1._1._2 === "LytLetStmt" || v1._1._2 === "LytWhere" || v1._1._2 === "LytOf" || v1._1._2 === "LytDo" || v1._1._2 === "LytAdo"
            ? Data$dArray.snoc(v2)(Data$dTuple.$Tuple(
                {range: {start: tokPos, end: tokPos}, leadingComments: [], trailingComments: [], value: PureScript$dCST$dTypes.$Token("TokLayoutEnd", v1._1._1.column)},
                v1._2
              ))
            : v2;
          continue;
        }
        go$c = false;
        go$r = Data$dTuple.$Tuple(v1, v2);
      }
      return go$r;
    };
    return v$1 => go(v$1._1)(v$1._2);
  };
  const insertKwProperty = (k, state) => {
    const $0 = insertSep(collapse(lytPos => lyt => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)(state));
    const v1 = Data$dTuple.$Tuple($0._1, Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1)));
    if (v1._1.tag === "Cons" && v1._1._1._2 === "LytProperty") { return Data$dTuple.$Tuple(v1._1._2, v1._2); }
    return k(v1);
  };
  if (v.value.tag === "TokLowerName") {
    if (v.value._1.tag === "Nothing") {
      if (v.value._2 === "data") {
        const $0 = insertSep(collapse(lytPos => lyt => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)(Data$dTuple.$Tuple(
          stack,
          []
        )));
        const v2 = Data$dTuple.$Tuple($0._1, Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1)));
        if (
          v2._1.tag === "Cons" && v2._1._1._2 === "LytWhere" && v2._1._2.tag === "Cons" && v2._1._2._1._2 === "LytRoot" && v2._1._2._2.tag === "Nil" && tokPos.column === v2._1._1._1.column
        ) {
          return Data$dTuple.$Tuple(Data$dList$dTypes.$List("Cons", Data$dTuple.$Tuple(tokPos, LytTopDecl), v2._1), v2._2);
        }
        if (v2._1.tag === "Cons" && eqLayoutDelim.eq(v2._1._1._2)(LytProperty)) { return Data$dTuple.$Tuple(v2._1._2, v2._2); }
        return v2;
      }
      if (v.value._2 === "class") {
        const $0 = insertSep(collapse(lytPos => lyt => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)(Data$dTuple.$Tuple(
          stack,
          []
        )));
        const v2 = Data$dTuple.$Tuple($0._1, Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1)));
        if (
          v2._1.tag === "Cons" && v2._1._1._2 === "LytWhere" && v2._1._2.tag === "Cons" && v2._1._2._1._2 === "LytRoot" && v2._1._2._2.tag === "Nil" && tokPos.column === v2._1._1._1.column
        ) {
          return Data$dTuple.$Tuple(Data$dList$dTypes.$List("Cons", Data$dTuple.$Tuple(tokPos, LytTopDeclHead), v2._1), v2._2);
        }
        if (v2._1.tag === "Cons" && eqLayoutDelim.eq(v2._1._1._2)(LytProperty)) { return Data$dTuple.$Tuple(v2._1._2, v2._2); }
        return v2;
      }
      if (v.value._2 === "where") {
        const whereP = v2 => v3 => v3 === "LytDo" || (v3 === "LytLet" || v3 === "LytLetStmt" || v3 === "LytWhere" || v3 === "LytOf" || v3 === "LytDo" || v3 === "LytAdo") && tokPos.column <= v2.column;
        if (stack.tag === "Cons") {
          if (stack._1._2 === "LytTopDeclHead") { return insertStart(LytWhere)(Data$dTuple.$Tuple(stack._2, Data$dArray.snoc([])(Data$dTuple.$Tuple(v, stack._2)))); }
          if (stack._1._2 === "LytProperty") { return Data$dTuple.$Tuple(stack._2, Data$dArray.snoc([])(Data$dTuple.$Tuple(v, stack._2))); }
        }
        return insertStart(LytWhere)((() => {
          const $0 = collapse(whereP)(Data$dTuple.$Tuple(stack, []));
          return Data$dTuple.$Tuple($0._1, Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1)));
        })());
      }
      if (v.value._2 === "in") {
        const v2 = collapse(v2 => v3 => {
          if (v3 === "LytLet") { return false; }
          if (v3 === "LytAdo") { return false; }
          return v3 === "LytLet" || v3 === "LytLetStmt" || v3 === "LytWhere" || v3 === "LytOf" || v3 === "LytDo" || v3 === "LytAdo";
        })(Data$dTuple.$Tuple(stack, []));
        if (v2._1.tag === "Cons") {
          if (v2._1._1._2 === "LytLetStmt" && v2._1._2.tag === "Cons" && v2._1._2._1._2 === "LytAdo") {
            return Data$dTuple.$Tuple(
              v2._1._2._2,
              Data$dArray.snoc(Data$dArray.snoc(Data$dArray.snoc(v2._2)(Data$dTuple.$Tuple(
                {range: {start: tokPos, end: tokPos}, leadingComments: [], trailingComments: [], value: PureScript$dCST$dTypes.$Token("TokLayoutEnd", v2._1._1._1.column)},
                v2._1._2._2
              )))(Data$dTuple.$Tuple(
                {range: {start: tokPos, end: tokPos}, leadingComments: [], trailingComments: [], value: PureScript$dCST$dTypes.$Token("TokLayoutEnd", v2._1._2._1._1.column)},
                v2._1._2._2
              )))(Data$dTuple.$Tuple(v, v2._1._2._2))
            );
          }
          if (
            v2._1._1._2 === "LytLet" || v2._1._1._2 === "LytLetStmt" || v2._1._1._2 === "LytWhere" || v2._1._1._2 === "LytOf" || v2._1._1._2 === "LytDo" || v2._1._1._2 === "LytAdo"
          ) {
            return Data$dTuple.$Tuple(
              v2._1._2,
              Data$dArray.snoc(Data$dArray.snoc(v2._2)(Data$dTuple.$Tuple(
                {range: {start: tokPos, end: tokPos}, leadingComments: [], trailingComments: [], value: PureScript$dCST$dTypes.$Token("TokLayoutEnd", v2._1._1._1.column)},
                v2._1._2
              )))(Data$dTuple.$Tuple(v, v2._1._2))
            );
          }
        }
        const $0 = insertSep(collapse(lytPos => lyt => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)(Data$dTuple.$Tuple(
          stack,
          []
        )));
        const $1 = Data$dTuple.$Tuple($0._1, Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1)));
        if ($1._1.tag === "Cons" && eqLayoutDelim.eq($1._1._1._2)(LytProperty)) { return Data$dTuple.$Tuple($1._1._2, $1._2); }
        return $1;
      }
      if (v.value._2 === "let") {
        return insertKwProperty(
          v2 => {
            if (v2._1.tag === "Cons") {
              if (v2._1._1._2 === "LytDo") {
                if (v2._1._1._1.column === tokPos.column) { return insertStart(LytLetStmt)(v2); }
                return insertStart(LytLet)(v2);
              }
              if (v2._1._1._2 === "LytAdo" && v2._1._1._1.column === tokPos.column) { return insertStart(LytLetStmt)(v2); }
            }
            return insertStart(LytLet)(v2);
          },
          Data$dTuple.$Tuple(stack, [])
        );
      }
      if (v.value._2 === "do") { return insertKwProperty(insertStart(LytDo), Data$dTuple.$Tuple(stack, [])); }
      if (v.value._2 === "ado") { return insertKwProperty(insertStart(LytAdo), Data$dTuple.$Tuple(stack, [])); }
      if (v.value._2 === "case") {
        return insertKwProperty(v1 => Data$dTuple.$Tuple(Data$dList$dTypes.$List("Cons", Data$dTuple.$Tuple(tokPos, LytCase), v1._1), v1._2), Data$dTuple.$Tuple(stack, []));
      }
      if (v.value._2 === "of") {
        const v2 = collapse(v$1 => isIndented)(Data$dTuple.$Tuple(stack, []));
        if (v2._1.tag === "Cons" && v2._1._1._2 === "LytCase") {
          const $0 = insertStart(LytOf)(Data$dTuple.$Tuple(v2._1._2, Data$dArray.snoc(v2._2)(Data$dTuple.$Tuple(v, v2._1._2))));
          return Data$dTuple.$Tuple(Data$dList$dTypes.$List("Cons", Data$dTuple.$Tuple(nextPos, LytCaseBinders), $0._1), $0._2);
        }
        const $0 = insertSep(collapse(lytPos => lyt => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)(v2));
        const $1 = Data$dTuple.$Tuple($0._1, Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1)));
        if ($1._1.tag === "Cons" && eqLayoutDelim.eq($1._1._1._2)(LytProperty)) { return Data$dTuple.$Tuple($1._1._2, $1._2); }
        return $1;
      }
      if (v.value._2 === "if") {
        return insertKwProperty(v1 => Data$dTuple.$Tuple(Data$dList$dTypes.$List("Cons", Data$dTuple.$Tuple(tokPos, LytIf), v1._1), v1._2), Data$dTuple.$Tuple(stack, []));
      }
      if (v.value._2 === "then") {
        const v2 = collapse(v$1 => isIndented)(Data$dTuple.$Tuple(stack, []));
        if (v2._1.tag === "Cons" && v2._1._1._2 === "LytIf") {
          return Data$dTuple.$Tuple(Data$dList$dTypes.$List("Cons", Data$dTuple.$Tuple(tokPos, LytThen), v2._1._2), Data$dArray.snoc(v2._2)(Data$dTuple.$Tuple(v, v2._1._2)));
        }
        const $0 = insertSep(collapse(lytPos => lyt => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)(Data$dTuple.$Tuple(
          stack,
          []
        )));
        const $1 = Data$dTuple.$Tuple($0._1, Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1)));
        if ($1._1.tag === "Cons" && eqLayoutDelim.eq($1._1._1._2)(LytProperty)) { return Data$dTuple.$Tuple($1._1._2, $1._2); }
        return $1;
      }
      if (v.value._2 === "else") {
        const v2 = collapse(v$1 => isIndented)(Data$dTuple.$Tuple(stack, []));
        if (v2._1.tag === "Cons" && v2._1._1._2 === "LytThen") { return Data$dTuple.$Tuple(v2._1._2, Data$dArray.snoc(v2._2)(Data$dTuple.$Tuple(v, v2._1._2))); }
        const v3 = collapse(lytPos => lyt => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)(Data$dTuple.$Tuple(
          stack,
          []
        ));
        if (
          v3._1.tag === "Cons" && v3._1._1._2 === "LytWhere" && v3._1._2.tag === "Cons" && v3._1._2._1._2 === "LytRoot" && v3._1._2._2.tag === "Nil" && tokPos.column === v3._1._1._1.column
        ) {
          return Data$dTuple.$Tuple(v3._1, Data$dArray.snoc(v3._2)(Data$dTuple.$Tuple(v, v3._1)));
        }
        const $0 = insertSep(v3);
        const $1 = Data$dTuple.$Tuple($0._1, Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1)));
        if ($1._1.tag === "Cons" && eqLayoutDelim.eq($1._1._1._2)(LytProperty)) { return Data$dTuple.$Tuple($1._1._2, $1._2); }
        return $1;
      }
      const $0 = insertSep(collapse(lytPos => lyt => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)(Data$dTuple.$Tuple(
        stack,
        []
      )));
      const $1 = Data$dTuple.$Tuple($0._1, Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1)));
      if ($1._1.tag === "Cons" && eqLayoutDelim.eq($1._1._1._2)(LytProperty)) { return Data$dTuple.$Tuple($1._1._2, $1._2); }
      return $1;
    }
    if (v.value._2 === "do") { return insertKwProperty(insertStart(LytDo), Data$dTuple.$Tuple(stack, [])); }
    if (v.value._2 === "ado") { return insertKwProperty(insertStart(LytAdo), Data$dTuple.$Tuple(stack, [])); }
    const $0 = insertSep(collapse(lytPos => lyt => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)(Data$dTuple.$Tuple(
      stack,
      []
    )));
    return Data$dTuple.$Tuple($0._1, Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1)));
  }
  if (v.value.tag === "TokForall") {
    return insertKwProperty(v1 => Data$dTuple.$Tuple(Data$dList$dTypes.$List("Cons", Data$dTuple.$Tuple(tokPos, LytForall), v1._1), v1._2), Data$dTuple.$Tuple(stack, []));
  }
  if (v.value.tag === "TokBackslash") {
    const $0 = insertSep(collapse(lytPos => lyt => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)(Data$dTuple.$Tuple(
      stack,
      []
    )));
    return Data$dTuple.$Tuple(Data$dList$dTypes.$List("Cons", Data$dTuple.$Tuple(tokPos, LytLambdaBinders), $0._1), Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1)));
  }
  if (v.value.tag === "TokRightArrow") {
    const $0 = collapse(v2 => v3 => {
      if (v3 === "LytDo") { return true; }
      if (v3 === "LytOf") { return false; }
      return (v3 === "LytLet" || v3 === "LytLetStmt" || v3 === "LytWhere" || v3 === "LytOf" || v3 === "LytDo" || v3 === "LytAdo") && tokPos.column <= v2.column;
    })(Data$dTuple.$Tuple(stack, []));
    if ($0._1.tag === "Cons" && ($0._1._1._2 === "LytCaseBinders" || $0._1._1._2 === "LytCaseGuard" || $0._1._1._2 === "LytLambdaBinders")) {
      return Data$dTuple.$Tuple($0._1._2, Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1._2)));
    }
    return Data$dTuple.$Tuple($0._1, Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1)));
  }
  if (v.value.tag === "TokEquals") {
    const v2 = collapse(v2 => v3 => v3 === "LytWhere" || v3 === "LytLet" || v3 === "LytLetStmt")(Data$dTuple.$Tuple(stack, []));
    if (v2._1.tag === "Cons" && v2._1._1._2 === "LytDeclGuard") { return Data$dTuple.$Tuple(v2._1._2, Data$dArray.snoc(v2._2)(Data$dTuple.$Tuple(v, v2._1._2))); }
    const $0 = insertSep(collapse(lytPos => lyt => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)(Data$dTuple.$Tuple(
      stack,
      []
    )));
    return Data$dTuple.$Tuple($0._1, Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1)));
  }
  if (v.value.tag === "TokPipe") {
    const v2 = collapse(lytPos => lyt => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column <= lytPos.column)(Data$dTuple.$Tuple(
      stack,
      []
    ));
    if (v2._1.tag === "Cons") {
      if (v2._1._1._2 === "LytOf") {
        const $0 = Data$dList$dTypes.$List("Cons", Data$dTuple.$Tuple(tokPos, LytCaseGuard), v2._1);
        return Data$dTuple.$Tuple($0, Data$dArray.snoc(v2._2)(Data$dTuple.$Tuple(v, $0)));
      }
      if (v2._1._1._2 === "LytLet") {
        const $0 = Data$dList$dTypes.$List("Cons", Data$dTuple.$Tuple(tokPos, LytDeclGuard), v2._1);
        return Data$dTuple.$Tuple($0, Data$dArray.snoc(v2._2)(Data$dTuple.$Tuple(v, $0)));
      }
      if (v2._1._1._2 === "LytLetStmt") {
        const $0 = Data$dList$dTypes.$List("Cons", Data$dTuple.$Tuple(tokPos, LytDeclGuard), v2._1);
        return Data$dTuple.$Tuple($0, Data$dArray.snoc(v2._2)(Data$dTuple.$Tuple(v, $0)));
      }
      if (v2._1._1._2 === "LytWhere") {
        const $0 = Data$dList$dTypes.$List("Cons", Data$dTuple.$Tuple(tokPos, LytDeclGuard), v2._1);
        return Data$dTuple.$Tuple($0, Data$dArray.snoc(v2._2)(Data$dTuple.$Tuple(v, $0)));
      }
    }
    const $0 = insertSep(collapse(lytPos => lyt => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)(Data$dTuple.$Tuple(
      stack,
      []
    )));
    return Data$dTuple.$Tuple($0._1, Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1)));
  }
  if (v.value.tag === "TokTick") {
    const v2 = collapse(v$1 => isIndented)(Data$dTuple.$Tuple(stack, []));
    if (v2._1.tag === "Cons" && v2._1._1._2 === "LytTick") { return Data$dTuple.$Tuple(v2._1._2, Data$dArray.snoc(v2._2)(Data$dTuple.$Tuple(v, v2._1._2))); }
    const $0 = insertSep(collapse(lytPos => lyt => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column <= lytPos.column)(Data$dTuple.$Tuple(
      stack,
      []
    )));
    return Data$dTuple.$Tuple(Data$dList$dTypes.$List("Cons", Data$dTuple.$Tuple(tokPos, LytTick), $0._1), Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1)));
  }
  if (v.value.tag === "TokComma") {
    const v2 = collapse(v$1 => isIndented)(Data$dTuple.$Tuple(stack, []));
    if (v2._1.tag === "Cons" && v2._1._1._2 === "LytBrace") {
      return Data$dTuple.$Tuple(Data$dList$dTypes.$List("Cons", Data$dTuple.$Tuple(tokPos, LytProperty), v2._1), Data$dArray.snoc(v2._2)(Data$dTuple.$Tuple(v, v2._1)));
    }
    return Data$dTuple.$Tuple(v2._1, Data$dArray.snoc(v2._2)(Data$dTuple.$Tuple(v, v2._1)));
  }
  if (v.value.tag === "TokDot") {
    const $0 = insertSep(collapse(lytPos => lyt => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)(Data$dTuple.$Tuple(
      stack,
      []
    )));
    const $1 = Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1));
    if ($0._1.tag === "Cons" && $0._1._1._2 === "LytForall") { return Data$dTuple.$Tuple($0._1._2, $1); }
    return Data$dTuple.$Tuple(Data$dList$dTypes.$List("Cons", Data$dTuple.$Tuple(tokPos, LytProperty), $0._1), $1);
  }
  if (v.value.tag === "TokLeftParen") {
    const $0 = insertSep(collapse(lytPos => lyt => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)(Data$dTuple.$Tuple(
      stack,
      []
    )));
    return Data$dTuple.$Tuple(Data$dList$dTypes.$List("Cons", Data$dTuple.$Tuple(tokPos, LytParen), $0._1), Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1)));
  }
  if (v.value.tag === "TokLeftBrace") {
    const $0 = insertSep(collapse(lytPos => lyt => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)(Data$dTuple.$Tuple(
      stack,
      []
    )));
    return Data$dTuple.$Tuple(
      Data$dList$dTypes.$List("Cons", Data$dTuple.$Tuple(tokPos, LytProperty), Data$dList$dTypes.$List("Cons", Data$dTuple.$Tuple(tokPos, LytBrace), $0._1)),
      Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1))
    );
  }
  if (v.value.tag === "TokLeftSquare") {
    const $0 = insertSep(collapse(lytPos => lyt => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)(Data$dTuple.$Tuple(
      stack,
      []
    )));
    return Data$dTuple.$Tuple(Data$dList$dTypes.$List("Cons", Data$dTuple.$Tuple(tokPos, LytSquare), $0._1), Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1)));
  }
  if (v.value.tag === "TokRightParen") {
    const $0 = collapse(v$1 => isIndented)(Data$dTuple.$Tuple(stack, []));
    if ($0._1.tag === "Cons" && eqLayoutDelim.eq($0._1._1._2)(LytParen)) { return Data$dTuple.$Tuple($0._1._2, Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1._2))); }
    return Data$dTuple.$Tuple($0._1, Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1)));
  }
  if (v.value.tag === "TokRightBrace") {
    const $0 = collapse(v$1 => isIndented)(Data$dTuple.$Tuple(stack, []));
    if ($0._1.tag === "Cons" && eqLayoutDelim.eq($0._1._1._2)(LytProperty)) {
      if ($0._1._2.tag === "Cons" && eqLayoutDelim.eq($0._1._2._1._2)(LytBrace)) {
        return Data$dTuple.$Tuple($0._1._2._2, Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1._2._2)));
      }
      return Data$dTuple.$Tuple($0._1._2, Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1._2)));
    }
    if ($0._1.tag === "Cons" && eqLayoutDelim.eq($0._1._1._2)(LytBrace)) { return Data$dTuple.$Tuple($0._1._2, Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1._2))); }
    return Data$dTuple.$Tuple($0._1, Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1)));
  }
  if (v.value.tag === "TokRightSquare") {
    const $0 = collapse(v$1 => isIndented)(Data$dTuple.$Tuple(stack, []));
    if ($0._1.tag === "Cons" && eqLayoutDelim.eq($0._1._1._2)(LytSquare)) { return Data$dTuple.$Tuple($0._1._2, Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1._2))); }
    return Data$dTuple.$Tuple($0._1, Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1)));
  }
  if (v.value.tag === "TokString") {
    const $0 = insertSep(collapse(lytPos => lyt => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)(Data$dTuple.$Tuple(
      stack,
      []
    )));
    const $1 = Data$dTuple.$Tuple($0._1, Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1)));
    if ($1._1.tag === "Cons" && eqLayoutDelim.eq($1._1._1._2)(LytProperty)) { return Data$dTuple.$Tuple($1._1._2, $1._2); }
    return $1;
  }
  if (v.value.tag === "TokOperator") {
    const $0 = insertSep(collapse(lytPos => lyt => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column <= lytPos.column)(Data$dTuple.$Tuple(
      stack,
      []
    )));
    return Data$dTuple.$Tuple($0._1, Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1)));
  }
  const $0 = insertSep(collapse(lytPos => lyt => (lyt === "LytLet" || lyt === "LytLetStmt" || lyt === "LytWhere" || lyt === "LytOf" || lyt === "LytDo" || lyt === "LytAdo") && tokPos.column < lytPos.column)(Data$dTuple.$Tuple(
    stack,
    []
  )));
  return Data$dTuple.$Tuple($0._1, Data$dArray.snoc($0._2)(Data$dTuple.$Tuple(v, $0._1)));
};
const ordLayoutDelim = {
  compare: x => y => {
    if (x === "LytRoot") {
      if (y === "LytRoot") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "LytRoot") { return Data$dOrdering.GT; }
    if (x === "LytTopDecl") {
      if (y === "LytTopDecl") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "LytTopDecl") { return Data$dOrdering.GT; }
    if (x === "LytTopDeclHead") {
      if (y === "LytTopDeclHead") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "LytTopDeclHead") { return Data$dOrdering.GT; }
    if (x === "LytDeclGuard") {
      if (y === "LytDeclGuard") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "LytDeclGuard") { return Data$dOrdering.GT; }
    if (x === "LytCase") {
      if (y === "LytCase") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "LytCase") { return Data$dOrdering.GT; }
    if (x === "LytCaseBinders") {
      if (y === "LytCaseBinders") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "LytCaseBinders") { return Data$dOrdering.GT; }
    if (x === "LytCaseGuard") {
      if (y === "LytCaseGuard") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "LytCaseGuard") { return Data$dOrdering.GT; }
    if (x === "LytLambdaBinders") {
      if (y === "LytLambdaBinders") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "LytLambdaBinders") { return Data$dOrdering.GT; }
    if (x === "LytParen") {
      if (y === "LytParen") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "LytParen") { return Data$dOrdering.GT; }
    if (x === "LytBrace") {
      if (y === "LytBrace") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "LytBrace") { return Data$dOrdering.GT; }
    if (x === "LytSquare") {
      if (y === "LytSquare") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "LytSquare") { return Data$dOrdering.GT; }
    if (x === "LytIf") {
      if (y === "LytIf") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "LytIf") { return Data$dOrdering.GT; }
    if (x === "LytThen") {
      if (y === "LytThen") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "LytThen") { return Data$dOrdering.GT; }
    if (x === "LytProperty") {
      if (y === "LytProperty") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "LytProperty") { return Data$dOrdering.GT; }
    if (x === "LytForall") {
      if (y === "LytForall") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "LytForall") { return Data$dOrdering.GT; }
    if (x === "LytTick") {
      if (y === "LytTick") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "LytTick") { return Data$dOrdering.GT; }
    if (x === "LytLet") {
      if (y === "LytLet") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "LytLet") { return Data$dOrdering.GT; }
    if (x === "LytLetStmt") {
      if (y === "LytLetStmt") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "LytLetStmt") { return Data$dOrdering.GT; }
    if (x === "LytWhere") {
      if (y === "LytWhere") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "LytWhere") { return Data$dOrdering.GT; }
    if (x === "LytOf") {
      if (y === "LytOf") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "LytOf") { return Data$dOrdering.GT; }
    if (x === "LytDo") {
      if (y === "LytDo") { return Data$dOrdering.EQ; }
      return Data$dOrdering.LT;
    }
    if (y === "LytDo") { return Data$dOrdering.GT; }
    if (x === "LytAdo" && y === "LytAdo") { return Data$dOrdering.EQ; }
    $runtime.fail();
  },
  Eq0: () => eqLayoutDelim
};
const currentIndent = /* #__PURE__ */ (() => {
  const go = go$a0$copy => {
    let go$a0 = go$a0$copy, go$c = true, go$r;
    while (go$c) {
      const v = go$a0;
      if (v.tag === "Cons") {
        if (v._1._2 === "LytLet" || v._1._2 === "LytLetStmt" || v._1._2 === "LytWhere" || v._1._2 === "LytOf" || v._1._2 === "LytDo" || v._1._2 === "LytAdo") {
          go$c = false;
          go$r = Data$dMaybe.$Maybe("Just", v._1._1);
          continue;
        }
        go$a0 = v._2;
        continue;
      }
      go$c = false;
      go$r = Data$dMaybe.Nothing;
    }
    return go$r;
  };
  return go;
})();
export {
  $LayoutDelim,
  LytAdo,
  LytBrace,
  LytCase,
  LytCaseBinders,
  LytCaseGuard,
  LytDeclGuard,
  LytDo,
  LytForall,
  LytIf,
  LytLambdaBinders,
  LytLet,
  LytLetStmt,
  LytOf,
  LytParen,
  LytProperty,
  LytRoot,
  LytSquare,
  LytThen,
  LytTick,
  LytTopDecl,
  LytTopDeclHead,
  LytWhere,
  currentIndent,
  eqLayoutDelim,
  insertLayout,
  isIndented,
  lytToken,
  ordLayoutDelim
};
