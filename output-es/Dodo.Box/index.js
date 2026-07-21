import * as $runtime from "../runtime.js";
import * as Data$dEither from "../Data.Either/index.js";
import * as Data$dInt from "../Data.Int/index.js";
import * as Data$dList$dTypes from "../Data.List.Types/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dMonoid from "../Data.Monoid/index.js";
import * as Data$dNumber from "../Data.Number/index.js";
import * as Data$dOrd from "../Data.Ord/index.js";
import * as Dodo$dInternal from "../Dodo.Internal/index.js";
import * as Partial from "../Partial/index.js";
const $Align = tag => tag;
const $DocBox = (tag, _1, _2, _3) => ({tag, _1, _2, _3});
const $DocBoxStep = (tag, _1, _2, _3) => ({tag, _1, _2, _3});
const $DocBuildCmd = (tag, _1, _2, _3) => ({tag, _1, _2, _3});
const $DocBuildSize = (tag, _1) => ({tag, _1});
const $DocBuildStk = (tag, _1, _2, _3, _4) => ({tag, _1, _2, _3, _4});
const $DocLine = (tag, _1, _2) => ({tag, _1, _2});
const $DocResumeCmd = (tag, _1) => ({tag, _1});
const $DocResumeStk = (tag, _1, _2, _3) => ({tag, _1, _2, _3});
const max = x => y => {
  const v = Data$dOrd.ordInt.compare(x)(y);
  if (v === "LT") { return y; }
  if (v === "EQ") { return x; }
  if (v === "GT") { return x; }
  $runtime.fail();
};
const power = /* #__PURE__ */ Data$dMonoid.power(Dodo$dInternal.monoidDoc);
const LinePad = value0 => $DocLine("LinePad", value0);
const LineDoc = value0 => $DocLine("LineDoc", value0);
const LineAppend = value0 => value1 => $DocLine("LineAppend", value0, value1);
const FullHeight = value0 => $DocBuildSize("FullHeight", value0);
const FullWidth = value0 => $DocBuildSize("FullWidth", value0);
const AsIs = /* #__PURE__ */ $DocBuildSize("AsIs");
const StpDone = /* #__PURE__ */ $DocBoxStep("StpDone");
const StpLine = value0 => value1 => $DocBoxStep("StpLine", value0, value1);
const StpPad = value0 => value1 => value2 => $DocBoxStep("StpPad", value0, value1, value2);
const StpHorz = value0 => value1 => value2 => $DocBoxStep("StpHorz", value0, value1, value2);
const ResumeEnter = value0 => $DocResumeCmd("ResumeEnter", value0);
const ResumeLeave = value0 => $DocResumeCmd("ResumeLeave", value0);
const ResumeHorzR = value0 => value1 => value2 => $DocResumeStk("ResumeHorzR", value0, value1, value2);
const ResumeHorzH = value0 => value1 => value2 => $DocResumeStk("ResumeHorzH", value0, value1, value2);
const ResumeNil = /* #__PURE__ */ $DocResumeStk("ResumeNil");
const Start = /* #__PURE__ */ $Align("Start");
const Middle = /* #__PURE__ */ $Align("Middle");
const End = /* #__PURE__ */ $Align("End");
const DocLine = value0 => value1 => $DocBox("DocLine", value0, value1);
const DocVApp = value0 => value1 => value2 => $DocBox("DocVApp", value0, value1, value2);
const DocHApp = value0 => value1 => value2 => $DocBox("DocHApp", value0, value1, value2);
const DocAlign = value0 => value1 => value2 => $DocBox("DocAlign", value0, value1, value2);
const DocPad = value0 => $DocBox("DocPad", value0);
const DocEmpty = /* #__PURE__ */ $DocBox("DocEmpty");
const BuildEnter = value0 => value1 => value2 => $DocBuildCmd("BuildEnter", value0, value1, value2);
const BuildLeave = value0 => $DocBuildCmd("BuildLeave", value0);
const BuildVAppR = value0 => value1 => value2 => $DocBuildStk("BuildVAppR", value0, value1, value2);
const BuildHAppR = value0 => value1 => value2 => value3 => $DocBuildStk("BuildHAppR", value0, value1, value2, value3);
const BuildHAppH = value0 => value1 => value2 => $DocBuildStk("BuildHAppH", value0, value1, value2);
const BuildNil = /* #__PURE__ */ $DocBuildStk("BuildNil");
const Horizontal = x => x;
const Vertical = x => x;
const newtypeVertical_ = {Coercible0: () => {}};
const newtypeHorizontal_ = {Coercible0: () => {}};
const functorDocBox = {
  map: f => m => {
    if (m.tag === "DocLine") { return $DocBox("DocLine", Dodo$dInternal.functorDoc.map(f)(m._1), m._2); }
    if (m.tag === "DocVApp") { return $DocBox("DocVApp", functorDocBox.map(f)(m._1), functorDocBox.map(f)(m._2), m._3); }
    if (m.tag === "DocHApp") { return $DocBox("DocHApp", functorDocBox.map(f)(m._1), functorDocBox.map(f)(m._2), m._3); }
    if (m.tag === "DocAlign") { return $DocBox("DocAlign", m._1, m._2, functorDocBox.map(f)(m._3)); }
    if (m.tag === "DocPad") { return $DocBox("DocPad", m._1); }
    if (m.tag === "DocEmpty") { return DocEmpty; }
    $runtime.fail();
  }
};
const functorHorizontal = functorDocBox;
const functorVertical = functorDocBox;
const eqAlign = {
  eq: x => y => {
    if (x === "Start") { return y === "Start"; }
    if (x === "Middle") { return y === "Middle"; }
    return x === "End" && y === "End";
  }
};
const vpadding = height => {
  if (height <= 0) { return DocEmpty; }
  return $DocBox("DocPad", {height, width: 0});
};
const valign = a => v => {
  if (v.tag === "DocAlign" && a === "Start" && v._2 === "Start") { return v._3; }
  if (v.tag === "DocAlign") { return $DocBox("DocAlign", a, v._2, v._3); }
  return $DocBox("DocAlign", a, Start, v);
};
const sizeOf = sizeOf$a0$copy => {
  let sizeOf$a0 = sizeOf$a0$copy, sizeOf$c = true, sizeOf$r;
  while (sizeOf$c) {
    const v = sizeOf$a0;
    if (v.tag === "DocLine") {
      sizeOf$c = false;
      sizeOf$r = {width: v._2, height: 1};
      continue;
    }
    if (v.tag === "DocVApp") {
      sizeOf$c = false;
      sizeOf$r = v._3;
      continue;
    }
    if (v.tag === "DocHApp") {
      sizeOf$c = false;
      sizeOf$r = v._3;
      continue;
    }
    if (v.tag === "DocAlign") {
      sizeOf$a0 = v._3;
      continue;
    }
    if (v.tag === "DocPad") {
      sizeOf$c = false;
      sizeOf$r = v._1;
      continue;
    }
    if (v.tag === "DocEmpty") {
      sizeOf$c = false;
      sizeOf$r = {width: 0, height: 0};
      continue;
    }
    $runtime.fail();
  }
  return sizeOf$r;
};
const vappend = v => v1 => {
  if (v.tag === "DocEmpty") { return v1; }
  if (v1.tag === "DocEmpty") { return v; }
  if (v.tag === "DocPad" && v1.tag === "DocPad") { return $DocBox("DocPad", {width: max(v._1.width)(v1._1.width), height: v._1.height + v1._1.height | 0}); }
  return $DocBox(
    "DocVApp",
    v,
    v1,
    (() => {
      const $0 = sizeOf(v);
      const $1 = sizeOf(v1);
      return {width: max($0.width)($1.width), height: $0.height + $1.height | 0};
    })()
  );
};
const vertical = dictFoldable => dictFoldable.foldr(vappend)(DocEmpty);
const semigroupVertical = {append: vappend};
const monoidVertical = {mempty: DocEmpty, Semigroup0: () => semigroupVertical};
const power1 = /* #__PURE__ */ Data$dMonoid.power(monoidVertical);
const resume = /* #__PURE__ */ (() => {
  const go = go$a0$copy => go$a1$copy => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
    while (go$c) {
      const cmd = go$a0, stack = go$a1;
      if (cmd.tag === "ResumeEnter") {
        if (cmd._1.tag === "StpDone") {
          go$a0 = $DocResumeCmd("ResumeLeave", Data$dMaybe.Nothing);
          go$a1 = stack;
          continue;
        }
        if (cmd._1.tag === "StpLine") {
          go$a0 = $DocResumeCmd("ResumeLeave", Data$dMaybe.$Maybe("Just", {line: $DocLine("LineDoc", cmd._1._1), next: cmd._1._2}));
          go$a1 = stack;
          continue;
        }
        if (cmd._1.tag === "StpPad") {
          if (cmd._1._2 === 0) {
            go$a0 = $DocResumeCmd("ResumeEnter", cmd._1._3);
            go$a1 = stack;
            continue;
          }
          go$a0 = $DocResumeCmd(
            "ResumeLeave",
            Data$dMaybe.$Maybe("Just", {line: $DocLine("LinePad", cmd._1._1), next: $DocBoxStep("StpPad", cmd._1._1, cmd._1._2 - 1 | 0, cmd._1._3)})
          );
          go$a1 = stack;
          continue;
        }
        if (cmd._1.tag === "StpHorz") {
          go$a0 = $DocResumeCmd("ResumeEnter", cmd._1._2);
          go$a1 = $DocResumeStk("ResumeHorzR", cmd._1._1, cmd._1._3, stack);
          continue;
        }
        $runtime.fail();
      }
      if (cmd.tag === "ResumeLeave") {
        if (stack.tag === "ResumeHorzR") {
          go$a0 = $DocResumeCmd("ResumeEnter", stack._1);
          go$a1 = $DocResumeStk("ResumeHorzH", cmd._1, stack._2, stack._3);
          continue;
        }
        if (stack.tag === "ResumeHorzH") {
          if (cmd._1.tag === "Just" && stack._1.tag === "Just") {
            go$a0 = $DocResumeCmd(
              "ResumeLeave",
              Data$dMaybe.$Maybe("Just", {line: $DocLine("LineAppend", cmd._1._1.line, stack._1._1.line), next: $DocBoxStep("StpHorz", cmd._1._1.next, stack._1._1.next, stack._2)})
            );
            go$a1 = stack._3;
            continue;
          }
          go$a0 = $DocResumeCmd("ResumeEnter", stack._2);
          go$a1 = stack._3;
          continue;
        }
        if (stack.tag === "ResumeNil") {
          go$c = false;
          go$r = cmd._1;
          continue;
        }
      }
      $runtime.fail();
    }
    return go$r;
  };
  return x => go($DocResumeCmd("ResumeEnter", x))(ResumeNil);
})();
const padWithAlign = appendFn => paddingFn => padWidth => doc => v => {
  if (v === "Start") { return appendFn(doc)(paddingFn(padWidth)); }
  if (v === "Middle") {
    const mid = Data$dInt.toNumber(padWidth) / 2.0;
    return appendFn(appendFn(paddingFn(Data$dInt.unsafeClamp(Data$dNumber.floor(mid))))(doc))(paddingFn(Data$dInt.unsafeClamp(Data$dNumber.ceil(mid))));
  }
  if (v === "End") { return appendFn(paddingFn(padWidth))(doc); }
  $runtime.fail();
};
const isEmpty = v => v.tag === "DocEmpty";
const hpadding = width => {
  if (width <= 0) { return DocEmpty; }
  return $DocBox("DocPad", {height: 1, width});
};
const happend = v => v1 => {
  if (v.tag === "DocEmpty") { return v1; }
  if (v1.tag === "DocEmpty") { return v; }
  if (v.tag === "DocPad" && v1.tag === "DocPad") { return $DocBox("DocPad", {width: v._1.width + v1._1.width | 0, height: max(v._1.height)(v1._1.height)}); }
  return $DocBox(
    "DocHApp",
    v,
    v1,
    (() => {
      const $0 = sizeOf(v);
      const $1 = sizeOf(v1);
      return {width: $0.width + $1.width | 0, height: max($0.height)($1.height)};
    })()
  );
};
const horizontal = dictFoldable => dictFoldable.foldr(happend)(DocEmpty);
const horizontalWithAlign = dictFoldable => align => dictFoldable.foldr(a => b => happend((() => {
  if (a.tag === "DocAlign" && align === "Start" && a._2 === "Start") { return a._3; }
  if (a.tag === "DocAlign") { return $DocBox("DocAlign", align, a._2, a._3); }
  return $DocBox("DocAlign", align, Start, a);
})())(b))(DocEmpty);
const semigroupHorizontal = {append: happend};
const monoidHorizontal = {mempty: DocEmpty, Semigroup0: () => semigroupHorizontal};
const halign = b => v => {
  if (v.tag === "DocAlign" && v._1 === "Start" && b === "Start") { return v._3; }
  if (v.tag === "DocAlign") { return $DocBox("DocAlign", v._1, b, v._3); }
  return $DocBox("DocAlign", Start, b, v);
};
const resize = newSize => box => {
  const size = sizeOf(box);
  const vpad = newSize.height - size.height | 0;
  const hpad = newSize.width - size.width | 0;
  const box$p = box.tag === "DocAlign" ? box._3 : box;
  const hdoc = (() => {
    if (hpad <= 0) {
      if (box$p.tag === "DocAlign" && box$p._2 === "Start") { return box$p._3; }
      if (box$p.tag === "DocAlign") { return $DocBox("DocAlign", Start, box$p._2, box$p._3); }
      return $DocBox("DocAlign", Start, Start, box$p);
    }
    return padWithAlign(happend)(hpadding)(hpad)(box)(box.tag === "DocAlign" ? box._2 : Start);
  })();
  if (vpad <= 0) {
    if (hdoc.tag === "DocAlign" && hdoc._1 === "Start") { return hdoc._3; }
    if (hdoc.tag === "DocAlign") { return $DocBox("DocAlign", hdoc._1, Start, hdoc._3); }
    return $DocBox("DocAlign", Start, Start, hdoc);
  }
  return padWithAlign(vappend)(vpadding)(vpad)(hdoc)(box.tag === "DocAlign" ? box._1 : Start);
};
const verticalWithAlign = dictFoldable => align => dictFoldable.foldr(a => b => vappend((() => {
  if (a.tag === "DocAlign" && a._1 === "Start" && align === "Start") { return a._3; }
  if (a.tag === "DocAlign") { return $DocBox("DocAlign", a._1, align, a._3); }
  return $DocBox("DocAlign", Start, align, a);
})())(b))(DocEmpty);
const formatLine = /* #__PURE__ */ (() => {
  const go = go$a0$copy => go$a1$copy => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
    while (go$c) {
      const acc = go$a0, v = go$a1;
      if (v.tag === "Nil") {
        go$c = false;
        go$r = acc;
        continue;
      }
      if (v.tag === "Cons") {
        if (v._1.tag === "LinePad") {
          if (acc.tag === "Empty") {
            go$a0 = acc;
            go$a1 = v._2;
            continue;
          }
          go$a0 = (() => {
            const $0 = power(Dodo$dInternal.$Doc("Text", 1, " "))(v._1._1);
            if ($0.tag === "Empty") { return acc; }
            if (acc.tag === "Empty") { return $0; }
            if ($0.tag === "Text" && acc.tag === "Text") { return Dodo$dInternal.$Doc("Text", $0._1 + acc._1 | 0, $0._2 + acc._2); }
            return Dodo$dInternal.$Doc("Append", $0, acc);
          })();
          go$a1 = v._2;
          continue;
        }
        if (v._1.tag === "LineDoc") {
          go$a0 = (() => {
            if (v._1._1.tag === "Empty") { return acc; }
            if (acc.tag === "Empty") { return v._1._1; }
            if (v._1._1.tag === "Text" && acc.tag === "Text") { return Dodo$dInternal.$Doc("Text", v._1._1._1 + acc._1 | 0, v._1._1._2 + acc._2); }
            return Dodo$dInternal.$Doc("Append", v._1._1, acc);
          })();
          go$a1 = v._2;
          continue;
        }
        if (v._1.tag === "LineAppend") {
          go$a0 = acc;
          go$a1 = Data$dList$dTypes.$List("Cons", v._1._2, Data$dList$dTypes.$List("Cons", v._1._1, v._2));
          continue;
        }
      }
      $runtime.fail();
    }
    return go$r;
  };
  const $0 = go(Dodo$dInternal.Empty);
  return x => $0(Data$dList$dTypes.$List("Cons", x, Data$dList$dTypes.Nil));
})();
const fill = ch => v => power1(ch.tag === "Annotate"
  ? $DocBox(
      "DocLine",
      (() => {
        const $0 = Dodo$dInternal.Annotate(ch._1);
        const $1 = power(ch._2)(v.width);
        if ($1.tag === "Empty") { return Dodo$dInternal.Empty; }
        return $0($1);
      })(),
      v.width
    )
  : $DocBox("DocLine", power(ch)(v.width), v.width))(v.height);
const empty = DocEmpty;
const docBox = /* #__PURE__ */ (() => {
  const go = go$a0$copy => go$a1$copy => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
    while (go$c) {
      const b = go$a0, v = go$a1;
      if (v.tag === "Nil") {
        go$c = false;
        go$r = b;
        continue;
      }
      if (v.tag === "Cons") {
        go$a0 = (() => {
          if (v._1.tag === "Left") {
            const $0 = Dodo$dInternal.Annotate(v._1._1);
            if (b.tag === "Empty") { return Dodo$dInternal.Empty; }
            return $0(b);
          }
          if (v._1.tag === "Right") {
            if (v._1._1.tag === "Empty") { return b; }
            if (b.tag === "Empty") { return v._1._1; }
            if (v._1._1.tag === "Text" && b.tag === "Text") { return Dodo$dInternal.$Doc("Text", v._1._1._1 + b._1 | 0, v._1._1._2 + b._2); }
            return Dodo$dInternal.$Doc("Append", v._1._1, b);
          }
          $runtime.fail();
        })();
        go$a1 = v._2;
        continue;
      }
      $runtime.fail();
    }
    return go$r;
  };
  const stkToDoc = go(Dodo$dInternal.Empty);
  return {
    emptyBuffer: {currentIndent: Dodo$dInternal.Empty, currentLine: Data$dList$dTypes.Nil, currentWidth: 0, lines: DocEmpty},
    writeText: width => text => v => (
      {
        ...v,
        currentLine: v.currentLine.tag === "Cons" && v.currentLine._1.tag === "Right"
          ? Data$dList$dTypes.$List(
              "Cons",
              Data$dEither.$Either(
                "Right",
                (() => {
                  if (v.currentLine._1._1.tag === "Empty") { return Dodo$dInternal.$Doc("Text", width, text); }
                  if (v.currentLine._1._1.tag === "Text") { return Dodo$dInternal.$Doc("Text", v.currentLine._1._1._1 + width | 0, v.currentLine._1._1._2 + text); }
                  return Dodo$dInternal.$Doc("Append", v.currentLine._1._1, Dodo$dInternal.$Doc("Text", width, text));
                })()
              ),
              v.currentLine._2
            )
          : Data$dList$dTypes.$List("Cons", Data$dEither.$Either("Right", Dodo$dInternal.$Doc("Text", width, text)), v.currentLine),
        currentWidth: v.currentWidth + width | 0
      }
    ),
    writeIndent: width => text => v => (
      {
        ...v,
        currentIndent: (() => {
          if (v.currentIndent.tag === "Empty") { return Dodo$dInternal.$Doc("Text", width, text); }
          if (v.currentIndent.tag === "Text") { return Dodo$dInternal.$Doc("Text", v.currentIndent._1 + width | 0, v.currentIndent._2 + text); }
          return Dodo$dInternal.$Doc("Append", v.currentIndent, Dodo$dInternal.$Doc("Text", width, text));
        })(),
        currentWidth: v.currentWidth + width | 0
      }
    ),
    writeBreak: v => (
      {
        ...v,
        currentIndent: Dodo$dInternal.Empty,
        currentLine: (() => {
          const go$1 = go$1$a0$copy => go$1$a1$copy => {
            let go$1$a0 = go$1$a0$copy, go$1$a1 = go$1$a1$copy, go$1$c = true, go$1$r;
            while (go$1$c) {
              const v$1 = go$1$a0, v1 = go$1$a1;
              if (v1.tag === "Nil") {
                const go$2 = go$2$a0$copy => go$2$a1$copy => {
                  let go$2$a0 = go$2$a0$copy, go$2$a1 = go$2$a1$copy, go$2$c = true, go$2$r;
                  while (go$2$c) {
                    const v$2 = go$2$a0, v1$1 = go$2$a1;
                    if (v1$1.tag === "Nil") {
                      go$2$c = false;
                      go$2$r = v$2;
                      continue;
                    }
                    if (v1$1.tag === "Cons") {
                      go$2$a0 = Data$dList$dTypes.$List("Cons", v1$1._1, v$2);
                      go$2$a1 = v1$1._2;
                      continue;
                    }
                    $runtime.fail();
                  }
                  return go$2$r;
                };
                go$1$c = false;
                go$1$r = go$2(Data$dList$dTypes.Nil)(v$1);
                continue;
              }
              if (v1.tag === "Cons") {
                if (
                  (() => {
                    if (v1._1.tag === "Left") { return true; }
                    if (v1._1.tag === "Right") { return false; }
                    $runtime.fail();
                  })()
                ) {
                  go$1$a0 = Data$dList$dTypes.$List("Cons", v1._1, v$1);
                  go$1$a1 = v1._2;
                  continue;
                }
                go$1$a0 = v$1;
                go$1$a1 = v1._2;
                continue;
              }
              $runtime.fail();
            }
            return go$1$r;
          };
          return go$1(Data$dList$dTypes.Nil)(v.currentLine);
        })(),
        currentWidth: 0,
        lines: vappend(v.lines)($DocBox(
          "DocLine",
          (() => {
            const $0 = stkToDoc(v.currentLine);
            if (v.currentIndent.tag === "Empty") { return $0; }
            if ($0.tag === "Empty") { return v.currentIndent; }
            if (v.currentIndent.tag === "Text" && $0.tag === "Text") { return Dodo$dInternal.$Doc("Text", v.currentIndent._1 + $0._1 | 0, v.currentIndent._2 + $0._2); }
            return Dodo$dInternal.$Doc("Append", v.currentIndent, $0);
          })(),
          v.currentWidth
        ))
      }
    ),
    enterAnnotation: ann => v => v1 => ({...v1, currentLine: Data$dList$dTypes.$List("Cons", Data$dEither.$Either("Left", ann), v1.currentLine)}),
    leaveAnnotation: v => v1 => v2 => (
      {
        ...v2,
        currentLine: (() => {
          if (v2.currentLine.tag === "Cons") {
            if (v2.currentLine._2.tag === "Cons" && v2.currentLine._2._1.tag === "Left" && v2.currentLine._1.tag === "Right") {
              return Data$dList$dTypes.$List(
                "Cons",
                Data$dEither.$Either(
                  "Right",
                  (() => {
                    const $0 = Dodo$dInternal.Annotate(v2.currentLine._2._1._1);
                    if (v2.currentLine._1._1.tag === "Empty") { return Dodo$dInternal.Empty; }
                    return $0(v2.currentLine._1._1);
                  })()
                ),
                v2.currentLine._2._2
              );
            }
            if (v2.currentLine._1.tag === "Left") { return v2.currentLine._2; }
          }
          return Partial._crashWith("leaveAnnotation: docs and annotations must be interleaved");
        })()
      }
    ),
    flushBuffer: v => {
      if (v.lines.tag === "DocEmpty" && v.currentLine.tag === "Nil") { return DocEmpty; }
      return vappend(v.lines)($DocBox(
        "DocLine",
        (() => {
          const $0 = stkToDoc(v.currentLine);
          if (v.currentIndent.tag === "Empty") { return $0; }
          if ($0.tag === "Empty") { return v.currentIndent; }
          if (v.currentIndent.tag === "Text" && $0.tag === "Text") { return Dodo$dInternal.$Doc("Text", v.currentIndent._1 + $0._1 | 0, v.currentIndent._2 + $0._2); }
          return Dodo$dInternal.$Doc("Append", v.currentIndent, $0);
        })(),
        v.currentWidth
      ));
    }
  };
})();
const build = /* #__PURE__ */ (() => {
  const go = go$a0$copy => go$a1$copy => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
    while (go$c) {
      const cmd = go$a0, stack = go$a1;
      if (cmd.tag === "BuildEnter") {
        if (cmd._1.tag === "FullHeight") {
          if (cmd._3.tag === "DocHApp") {
            go$a0 = $DocBuildCmd("BuildEnter", cmd._1, StpDone, cmd._3._2);
            go$a1 = $DocBuildStk("BuildHAppR", cmd._1._1, cmd._3._1, cmd._2, stack);
            continue;
          }
          go$a0 = $DocBuildCmd("BuildEnter", AsIs, cmd._2, resize({width: 0, height: cmd._1._1})(cmd._3));
          go$a1 = stack;
          continue;
        }
        if (cmd._1.tag === "FullWidth") {
          if (cmd._3.tag === "DocVApp") {
            go$a0 = $DocBuildCmd("BuildEnter", cmd._1, cmd._2, cmd._3._2);
            go$a1 = $DocBuildStk("BuildVAppR", cmd._1._1, cmd._3._1, stack);
            continue;
          }
          go$a0 = $DocBuildCmd("BuildEnter", AsIs, cmd._2, resize({width: cmd._1._1, height: 0})(cmd._3));
          go$a1 = stack;
          continue;
        }
        if (cmd._1.tag === "AsIs") {
          if (cmd._3.tag === "DocVApp") {
            go$a0 = $DocBuildCmd("BuildEnter", $DocBuildSize("FullWidth", cmd._3._3.width), cmd._2, cmd._3._2);
            go$a1 = $DocBuildStk("BuildVAppR", cmd._3._3.width, cmd._3._1, stack);
            continue;
          }
          if (cmd._3.tag === "DocHApp") {
            go$a0 = $DocBuildCmd("BuildEnter", $DocBuildSize("FullHeight", cmd._3._3.height), StpDone, cmd._3._2);
            go$a1 = $DocBuildStk("BuildHAppR", cmd._3._3.height, cmd._3._1, cmd._2, stack);
            continue;
          }
          if (cmd._3.tag === "DocAlign") {
            go$a0 = $DocBuildCmd("BuildEnter", cmd._1, cmd._2, cmd._3._3);
            go$a1 = stack;
            continue;
          }
          if (cmd._3.tag === "DocLine") {
            go$a0 = $DocBuildCmd("BuildLeave", $DocBoxStep("StpLine", cmd._3._1, cmd._2));
            go$a1 = stack;
            continue;
          }
          if (cmd._3.tag === "DocPad") {
            go$a0 = $DocBuildCmd("BuildLeave", $DocBoxStep("StpPad", cmd._3._1.width, cmd._3._1.height, cmd._2));
            go$a1 = stack;
            continue;
          }
          if (cmd._3.tag === "DocEmpty") {
            go$a0 = $DocBuildCmd("BuildLeave", StpDone);
            go$a1 = stack;
            continue;
          }
        }
        $runtime.fail();
      }
      if (cmd.tag === "BuildLeave") {
        if (stack.tag === "BuildVAppR") {
          go$a0 = $DocBuildCmd("BuildEnter", $DocBuildSize("FullWidth", stack._1), cmd._1, stack._2);
          go$a1 = stack._3;
          continue;
        }
        if (stack.tag === "BuildHAppR") {
          go$a0 = $DocBuildCmd("BuildEnter", $DocBuildSize("FullHeight", stack._1), StpDone, stack._2);
          go$a1 = $DocBuildStk("BuildHAppH", cmd._1, stack._3, stack._4);
          continue;
        }
        if (stack.tag === "BuildHAppH") {
          go$a0 = $DocBuildCmd("BuildLeave", $DocBoxStep("StpHorz", cmd._1, stack._1, stack._2));
          go$a1 = stack._3;
          continue;
        }
        if (stack.tag === "BuildNil") {
          go$c = false;
          go$r = cmd._1;
          continue;
        }
      }
      $runtime.fail();
    }
    return go$r;
  };
  return size => next => box => go($DocBuildCmd("BuildEnter", size, next, box))(BuildNil);
})();
const toDoc = /* #__PURE__ */ (() => {
  const go2 = go2$a0$copy => go2$a1$copy => {
    let go2$a0 = go2$a0$copy, go2$a1 = go2$a1$copy, go2$c = true, go2$r;
    while (go2$c) {
      const acc = go2$a0, v = go2$a1;
      if (v.tag === "Nothing") {
        go2$c = false;
        go2$r = acc;
        continue;
      }
      if (v.tag === "Just") {
        go2$a0 = (() => {
          const $0 = formatLine(v._1.line);
          const $1 = $0.tag === "Empty" ? Dodo$dInternal.Break : Dodo$dInternal.$Doc("Append", Dodo$dInternal.Break, $0);
          if (acc.tag === "Empty") { return $1; }
          if ($1.tag === "Empty") { return acc; }
          if (acc.tag === "Text" && $1.tag === "Text") { return Dodo$dInternal.$Doc("Text", acc._1 + $1._1 | 0, acc._2 + $1._2); }
          return Dodo$dInternal.$Doc("Append", acc, $1);
        })();
        go2$a1 = resume(v._1.next);
        continue;
      }
      $runtime.fail();
    }
    return go2$r;
  };
  const $0 = build(AsIs)(StpDone);
  return x => {
    const $1 = resume($0(x));
    if ($1.tag === "Nothing") { return Dodo$dInternal.Empty; }
    if ($1.tag === "Just") { return go2(formatLine($1._1.line))(resume($1._1.next)); }
    $runtime.fail();
  };
})();
export {
  $Align,
  $DocBox,
  $DocBoxStep,
  $DocBuildCmd,
  $DocBuildSize,
  $DocBuildStk,
  $DocLine,
  $DocResumeCmd,
  $DocResumeStk,
  AsIs,
  BuildEnter,
  BuildHAppH,
  BuildHAppR,
  BuildLeave,
  BuildNil,
  BuildVAppR,
  DocAlign,
  DocEmpty,
  DocHApp,
  DocLine,
  DocPad,
  DocVApp,
  End,
  FullHeight,
  FullWidth,
  Horizontal,
  LineAppend,
  LineDoc,
  LinePad,
  Middle,
  ResumeEnter,
  ResumeHorzH,
  ResumeHorzR,
  ResumeLeave,
  ResumeNil,
  Start,
  StpDone,
  StpHorz,
  StpLine,
  StpPad,
  Vertical,
  build,
  docBox,
  empty,
  eqAlign,
  fill,
  formatLine,
  functorDocBox,
  functorHorizontal,
  functorVertical,
  halign,
  happend,
  horizontal,
  horizontalWithAlign,
  hpadding,
  isEmpty,
  max,
  monoidHorizontal,
  monoidVertical,
  newtypeHorizontal_,
  newtypeVertical_,
  padWithAlign,
  power,
  power1,
  resize,
  resume,
  semigroupHorizontal,
  semigroupVertical,
  sizeOf,
  toDoc,
  valign,
  vappend,
  vertical,
  verticalWithAlign,
  vpadding
};
