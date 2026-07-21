// | This module provides functions printing with cascading ANSI styles.
// | ANSI annotations closer to the root will cascade down to child nodes,
// | where styles closer to the leaves take precedence. Indentation is
// | never printed with ANSI styles, only the text elements of the document.
import * as $runtime from "../runtime.js";
import * as Ansi$dCodes from "../Ansi.Codes/index.js";
import * as Data$dList from "../Data.List/index.js";
import * as Data$dList$dTypes from "../Data.List.Types/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dNonEmpty from "../Data.NonEmpty/index.js";
import * as Dodo$dInternal from "../Dodo.Internal/index.js";
const AnsiBuffer = x => x;
const underline = /* #__PURE__ */ (() => {
  const $0 = Dodo$dInternal.Annotate(Ansi$dCodes.$GraphicsParam("PMode", Ansi$dCodes.Underline));
  return v => {
    if (v.tag === "Empty") { return Dodo$dInternal.Empty; }
    return $0(v);
  };
})();
const strikethrough = /* #__PURE__ */ (() => {
  const $0 = Dodo$dInternal.Annotate(Ansi$dCodes.$GraphicsParam("PMode", Ansi$dCodes.Strikethrough));
  return v => {
    if (v.tag === "Empty") { return Dodo$dInternal.Empty; }
    return $0(v);
  };
})();
const reset = /* #__PURE__ */ (() => {
  const $0 = Dodo$dInternal.Annotate(Ansi$dCodes.Reset);
  return v => {
    if (v.tag === "Empty") { return Dodo$dInternal.Empty; }
    return $0(v);
  };
})();
const italic = /* #__PURE__ */ (() => {
  const $0 = Dodo$dInternal.Annotate(Ansi$dCodes.$GraphicsParam("PMode", Ansi$dCodes.Italic));
  return v => {
    if (v.tag === "Empty") { return Dodo$dInternal.Empty; }
    return $0(v);
  };
})();
const inverse = /* #__PURE__ */ (() => {
  const $0 = Dodo$dInternal.Annotate(Ansi$dCodes.$GraphicsParam("PMode", Ansi$dCodes.Inverse));
  return v => {
    if (v.tag === "Empty") { return Dodo$dInternal.Empty; }
    return $0(v);
  };
})();
const foreground = color => {
  const $0 = Dodo$dInternal.Annotate(Ansi$dCodes.$GraphicsParam("PForeground", color));
  return v => {
    if (v.tag === "Empty") { return Dodo$dInternal.Empty; }
    return $0(v);
  };
};
const dim = /* #__PURE__ */ (() => {
  const $0 = Dodo$dInternal.Annotate(Ansi$dCodes.$GraphicsParam("PMode", Ansi$dCodes.Dim));
  return v => {
    if (v.tag === "Empty") { return Dodo$dInternal.Empty; }
    return $0(v);
  };
})();
const bold = /* #__PURE__ */ (() => {
  const $0 = Dodo$dInternal.Annotate(Ansi$dCodes.$GraphicsParam("PMode", Ansi$dCodes.Bold));
  return v => {
    if (v.tag === "Empty") { return Dodo$dInternal.Empty; }
    return $0(v);
  };
})();
const background = color => {
  const $0 = Dodo$dInternal.Annotate(Ansi$dCodes.$GraphicsParam("PBackground", color));
  return v => {
    if (v.tag === "Empty") { return Dodo$dInternal.Empty; }
    return $0(v);
  };
};
const ansiGraphics = /* #__PURE__ */ (() => {
  const resetCode = Ansi$dCodes.escapeCodeToString(Ansi$dCodes.$EscapeCode("Graphics", Data$dNonEmpty.$NonEmpty(Ansi$dCodes.Reset, Data$dList$dTypes.Nil)));
  const go = go$a0$copy => go$a1$copy => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
    while (go$c) {
      const v = go$a0, v1 = go$a1;
      if (v1.tag === "Cons" && !Ansi$dCodes.eqGraphicsParam.eq(v1._1)(Ansi$dCodes.Reset)) {
        go$a0 = Data$dList$dTypes.$List("Cons", v1._1, v);
        go$a1 = v1._2;
        continue;
      }
      const go$1 = go$1$a0$copy => go$1$a1$copy => {
        let go$1$a0 = go$1$a0$copy, go$1$a1 = go$1$a1$copy, go$1$c = true, go$1$r;
        while (go$1$c) {
          const v$1 = go$1$a0, v1$1 = go$1$a1;
          if (v1$1.tag === "Nil") {
            go$1$c = false;
            go$1$r = v$1;
            continue;
          }
          if (v1$1.tag === "Cons") {
            go$1$a0 = Data$dList$dTypes.$List("Cons", v1$1._1, v$1);
            go$1$a1 = v1$1._2;
            continue;
          }
          $runtime.fail();
        }
        return go$1$r;
      };
      go$c = false;
      go$r = go$1(Data$dList$dTypes.Nil)(v);
    }
    return go$r;
  };
  const $0 = go(Data$dList$dTypes.Nil);
  const $1 = Data$dList.nubByEq(v => v1 => {
    if (v.tag === "Reset") { return v1.tag === "Reset"; }
    if (v.tag === "PForeground") { return v1.tag === "PForeground"; }
    if (v.tag === "PBackground") { return v1.tag === "PBackground"; }
    return v.tag === "PMode" && v1.tag === "PMode" && (() => {
      if (v._1 === "Bold") { return v1._1 === "Bold"; }
      if (v._1 === "Dim") { return v1._1 === "Dim"; }
      if (v._1 === "Italic") { return v1._1 === "Italic"; }
      if (v._1 === "Underline") { return v1._1 === "Underline"; }
      if (v._1 === "Inverse") { return v1._1 === "Inverse"; }
      return v._1 === "Strikethrough" && v1._1 === "Strikethrough";
    })();
  });
  const getCurrentGraphics = x => {
    const go$1 = go$1$a0$copy => go$1$a1$copy => {
      let go$1$a0 = go$1$a0$copy, go$1$a1 = go$1$a1$copy, go$1$c = true, go$1$r;
      while (go$1$c) {
        const v = go$1$a0, v1 = go$1$a1;
        if (v1.tag === "Nil") {
          go$1$c = false;
          go$1$r = v;
          continue;
        }
        if (v1.tag === "Cons") {
          go$1$a0 = Data$dList$dTypes.$List("Cons", v1._1, v);
          go$1$a1 = v1._2;
          continue;
        }
        $runtime.fail();
      }
      return go$1$r;
    };
    return go$1(Data$dList$dTypes.Nil)($1($0(x)));
  };
  return {
    emptyBuffer: {output: "", pending: Data$dMaybe.Nothing, current: Data$dList$dTypes.Nil, previous: Data$dList$dTypes.Nil},
    writeText: v => text => output => {
      const $2 = (() => {
        if (output.pending.tag === "Nothing") { return output; }
        if (output.pending.tag === "Just") {
          return {...output, output: output.output + Ansi$dCodes.escapeCodeToString(Ansi$dCodes.$EscapeCode("Graphics", output.pending._1)), pending: Data$dMaybe.Nothing};
        }
        $runtime.fail();
      })();
      return {...$2, output: $2.output + text, previous: Data$dList$dTypes.Nil};
    },
    writeIndent: v => text => v1 => ({...v1, output: v1.output + text}),
    writeBreak: v => {
      const pending = (() => {
        if (v.current.tag === "Nil") { return Data$dMaybe.Nothing; }
        if (v.current.tag === "Cons") {
          return Data$dMaybe.$Maybe("Just", Data$dNonEmpty.$NonEmpty(Ansi$dCodes.Reset, Data$dList$dTypes.$List("Cons", v.current._1, v.current._2)));
        }
        $runtime.fail();
      })();
      return {
        ...v,
        output: (() => {
          if (pending.tag === "Nothing") { return true; }
          if (pending.tag === "Just") { return false; }
          $runtime.fail();
        })() && v.previous.tag === "Nil"
          ? v.output + "\n"
          : v.output + resetCode + "\n",
        pending
      };
    },
    enterAnnotation: a => as => v => {
      const current = getCurrentGraphics(Data$dList$dTypes.$List("Cons", a, as));
      return {
        ...v,
        pending: (() => {
          if (current.tag === "Nil") { return Data$dMaybe.Nothing; }
          if (current.tag === "Cons") { return Data$dMaybe.$Maybe("Just", Data$dNonEmpty.$NonEmpty(Ansi$dCodes.Reset, Data$dList$dTypes.$List("Cons", current._1, current._2))); }
          $runtime.fail();
        })(),
        current,
        previous: v.current
      };
    },
    leaveAnnotation: v => as => v1 => {
      const current = getCurrentGraphics(as);
      return {
        ...v1,
        pending: Data$dMaybe.$Maybe(
          "Just",
          (() => {
            if (current.tag === "Nil") { return Data$dNonEmpty.$NonEmpty(Ansi$dCodes.Reset, Data$dList$dTypes.Nil); }
            if (current.tag === "Cons") { return Data$dNonEmpty.$NonEmpty(Ansi$dCodes.Reset, Data$dList$dTypes.$List("Cons", current._1, current._2)); }
            $runtime.fail();
          })()
        ),
        current,
        previous: v1.current
      };
    },
    flushBuffer: ansiBuffer => {
      if (ansiBuffer.pending.tag === "Nothing") { return ansiBuffer.output; }
      if (ansiBuffer.pending.tag === "Just") { return ansiBuffer.output + Ansi$dCodes.escapeCodeToString(Ansi$dCodes.$EscapeCode("Graphics", ansiBuffer.pending._1)); }
      $runtime.fail();
    }
  };
})();
export {AnsiBuffer, ansiGraphics, background, bold, dim, foreground, inverse, italic, reset, strikethrough, underline};
