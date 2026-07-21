import * as $runtime from "../runtime.js";
import * as Data$dFoldable from "../Data.Foldable/index.js";
import * as Data$dFunctor from "../Data.Functor/index.js";
import * as Data$dInt from "../Data.Int/index.js";
import * as Data$dList$dTypes from "../Data.List.Types/index.js";
import * as Data$dMonoid from "../Data.Monoid/index.js";
import * as Data$dNumber from "../Data.Number/index.js";
import * as Data$dOrd from "../Data.Ord/index.js";
import * as Data$dString$dCodePoints from "../Data.String.CodePoints/index.js";
import * as Data$dString$dCommon from "../Data.String.Common/index.js";
import * as Data$dString$dRegex from "../Data.String.Regex/index.js";
import * as Data$dString$dRegex$dFlags from "../Data.String.Regex.Flags/index.js";
import * as Data$dString$dRegex$dUnsafe from "../Data.String.Regex.Unsafe/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
import * as Dodo$dInternal from "../Dodo.Internal/index.js";
import * as Dodo$dInternal$dBuffer from "../Dodo.Internal.Buffer/index.js";
const $DocCmd = (tag, _1, _2) => ({tag, _1, _2});
const $FlexGroupStatus = (tag, _1) => ({tag, _1});
const max = x => y => {
  const v = Data$dOrd.ordInt.compare(x)(y);
  if (v === "LT") { return y; }
  if (v === "EQ") { return x; }
  if (v === "GT") { return x; }
  $runtime.fail();
};
const max1 = x => y => {
  const v = Data$dOrd.ordNumber.compare(x)(y);
  if (v === "LT") { return y; }
  if (v === "EQ") { return x; }
  if (v === "GT") { return x; }
  $runtime.fail();
};
const min = x => y => {
  const v = Data$dOrd.ordNumber.compare(x)(y);
  if (v === "LT") { return x; }
  if (v === "EQ") { return x; }
  if (v === "GT") { return y; }
  $runtime.fail();
};
const power = /* #__PURE__ */ Data$dMonoid.power(Data$dMonoid.monoidString);
const Printer = x => x;
const Doc = value0 => $DocCmd("Doc", value0);
const Dedent = value0 => value1 => $DocCmd("Dedent", value0, value1);
const LeaveAnnotation = value0 => value1 => $DocCmd("LeaveAnnotation", value0, value1);
const LeaveFlexGroup = value0 => value1 => $DocCmd("LeaveFlexGroup", value0, value1);
const LeaveLocal = value0 => $DocCmd("LeaveLocal", value0);
const NoFlexGroup = /* #__PURE__ */ $FlexGroupStatus("NoFlexGroup");
const FlexGroupPending = /* #__PURE__ */ $FlexGroupStatus("FlexGroupPending");
const FlexGroupReset = value0 => $FlexGroupStatus("FlexGroupReset", value0);
const withPosition = Dodo$dInternal.WithPosition;
const withLocalOptions = Dodo$dInternal.Local;
const twoSpaces = {pageWidth: 80, ribbonRatio: 1.0, indentUnit: "  ", indentWidth: 2};
const text = v => {
  if (v === "") { return Dodo$dInternal.Empty; }
  return Dodo$dInternal.$Doc("Text", Data$dString$dCodePoints.toCodePointArray(v).length, v);
};
const tabs = {pageWidth: 120, ribbonRatio: 1.0, indentUnit: "\t", indentWidth: 4};
const space = /* #__PURE__ */ Dodo$dInternal.$Doc("Text", 1, " ");
const plainText = {
  emptyBuffer: "",
  writeText: v => str => buff => buff + str,
  writeIndent: v => str => buff => buff + str,
  writeBreak: buff => buff + "\n",
  enterAnnotation: v => v1 => buff => buff,
  leaveAnnotation: v => v1 => buff => buff,
  flushBuffer: buff => buff
};
const locally = k => doc => Dodo$dInternal.$Doc("Local", options => Data$dTuple.$Tuple(k(options), doc));
const indent = v => {
  if (v.tag === "Empty") { return Dodo$dInternal.Empty; }
  return Dodo$dInternal.$Doc("Indent", v);
};
const fourSpaces = {pageWidth: 120, ribbonRatio: 1.0, indentUnit: "    ", indentWidth: 4};
const foldWith = dictFoldable => f => dictFoldable.foldr(v => v1 => {
  if (v.tag === "Empty") { return v1; }
  if (v1.tag === "Empty") { return v; }
  return f(v)(v1);
})(Dodo$dInternal.Empty);
const foldWithSeparator = dictFoldable => separator => dictFoldable.foldr(v => v1 => {
  if (v.tag === "Empty") { return v1; }
  if (v1.tag === "Empty") { return v; }
  const $0 = (() => {
    if (separator.tag === "Empty") { return v1; }
    if (v1.tag === "Empty") { return separator; }
    if (separator.tag === "Text" && v1.tag === "Text") { return Dodo$dInternal.$Doc("Text", separator._1 + v1._1 | 0, separator._2 + v1._2); }
    return Dodo$dInternal.$Doc("Append", separator, v1);
  })();
  if (v.tag === "Empty") { return $0; }
  if ($0.tag === "Empty") { return v; }
  if (v.tag === "Text" && $0.tag === "Text") { return Dodo$dInternal.$Doc("Text", v._1 + $0._1 | 0, v._2 + $0._2); }
  return Dodo$dInternal.$Doc("Append", v, $0);
})(Dodo$dInternal.Empty);
const flexSelect = doc1 => doc2 => doc3 => {
  if (doc1.tag === "Empty") { return doc2; }
  return Dodo$dInternal.$Doc("FlexSelect", doc1, doc2, doc3);
};
const flexGroup = v => {
  if (v.tag === "Empty") { return Dodo$dInternal.Empty; }
  if (v.tag === "FlexSelect" && v._2.tag === "Empty" && v._3.tag === "Empty") { return v; }
  return Dodo$dInternal.$Doc("FlexSelect", v, Dodo$dInternal.Empty, Dodo$dInternal.Empty);
};
const flexAlt = Dodo$dInternal.FlexAlt;
const encloseWithSeparator = dictFoldable => open => close => separator => inner => {
  const $0 = foldWithSeparator(dictFoldable)(separator)(inner);
  const $1 = (() => {
    if ($0.tag === "Empty") { return close; }
    if (close.tag === "Empty") { return $0; }
    if ($0.tag === "Text" && close.tag === "Text") { return Dodo$dInternal.$Doc("Text", $0._1 + close._1 | 0, $0._2 + close._2); }
    return Dodo$dInternal.$Doc("Append", $0, close);
  })();
  if (open.tag === "Empty") { return $1; }
  if ($1.tag === "Empty") { return open; }
  if (open.tag === "Text" && $1.tag === "Text") { return Dodo$dInternal.$Doc("Text", open._1 + $1._1 | 0, open._2 + $1._2); }
  return Dodo$dInternal.$Doc("Append", open, $1);
};
const encloseEmptyAlt = open => close => $$default => inner => {
  if (inner.tag === "Empty") { return $$default; }
  const $0 = (() => {
    if (inner.tag === "Empty") { return close; }
    if (close.tag === "Empty") { return inner; }
    if (inner.tag === "Text" && close.tag === "Text") { return Dodo$dInternal.$Doc("Text", inner._1 + close._1 | 0, inner._2 + close._2); }
    return Dodo$dInternal.$Doc("Append", inner, close);
  })();
  if (open.tag === "Empty") { return $0; }
  if ($0.tag === "Empty") { return open; }
  if (open.tag === "Text" && $0.tag === "Text") { return Dodo$dInternal.$Doc("Text", open._1 + $0._1 | 0, open._2 + $0._2); }
  return Dodo$dInternal.$Doc("Append", open, $0);
};
const enclose = open => close => inner => {
  const $0 = (() => {
    if (inner.tag === "Empty") { return close; }
    if (close.tag === "Empty") { return inner; }
    if (inner.tag === "Text" && close.tag === "Text") { return Dodo$dInternal.$Doc("Text", inner._1 + close._1 | 0, inner._2 + close._2); }
    return Dodo$dInternal.$Doc("Append", inner, close);
  })();
  if (open.tag === "Empty") { return $0; }
  if ($0.tag === "Empty") { return open; }
  if (open.tag === "Text" && $0.tag === "Text") { return Dodo$dInternal.$Doc("Text", open._1 + $0._1 | 0, open._2 + $0._2); }
  return Dodo$dInternal.$Doc("Append", open, $0);
};
const calcRibbonWidth = v => n => max(0)(Data$dInt.unsafeClamp(Data$dNumber.ceil(v.ribbonRatio * Data$dInt.toNumber(v.pageWidth - n | 0))));
const storeOptions = prevIndent => localOptions => state => {
  const newOptions = {indentUnit: localOptions.indentUnit, indentWidth: localOptions.indentWidth, pageWidth: localOptions.pageWidth, ribbonRatio: localOptions.ribbonRatio};
  return {
    ...state,
    indentSpaces: localOptions.indentSpaces,
    options: newOptions,
    position: {...state.position, pageWidth: newOptions.pageWidth, ribbonWidth: calcRibbonWidth(newOptions)(prevIndent), nextIndent: localOptions.indent}
  };
};
const print = v => opts => {
  const initOptions = {...opts, ribbonRatio: max1(0.0)(min(1.0)(opts.ribbonRatio))};
  const go = go$a0$copy => go$a1$copy => {
    let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
    while (go$c) {
      const stack = go$a0, state = go$a1;
      if (stack.tag === "Nil") {
        go$c = false;
        go$r = v.flushBuffer(Dodo$dInternal$dBuffer.commit(state.buffer).buffer);
        continue;
      }
      if (stack.tag === "Cons") {
        if (stack._1.tag === "Doc") {
          if (stack._1._1.tag === "Append") {
            go$a0 = Data$dList$dTypes.$List("Cons", $DocCmd("Doc", stack._1._1._1), Data$dList$dTypes.$List("Cons", $DocCmd("Doc", stack._1._1._2), stack._2));
            go$a1 = state;
            continue;
          }
          if (stack._1._1.tag === "Text") {
            if (state.position.column === 0 && state.position.indent > 0) {
              go$a0 = stack;
              go$a1 = {
                ...state,
                position: {...state.position, column: state.position.indent},
                buffer: (() => {
                  const $0 = v.writeIndent(state.position.indent)(state.indentSpaces);
                  if (state.buffer.queue.tag === "Cons") {
                    return {...state.buffer, queue: Data$dList$dTypes.$List("Cons", Data$dList$dTypes.$List("Cons", $0, state.buffer.queue._1), state.buffer.queue._2)};
                  }
                  return {...state.buffer, buffer: $0(state.buffer.buffer)};
                })()
              };
              continue;
            }
            if ((state.position.column + stack._1._1._1 | 0) <= (state.position.indent + state.position.ribbonWidth | 0)) {
              go$a0 = stack._2;
              go$a1 = {
                ...state,
                position: {...state.position, column: state.position.column + stack._1._1._1 | 0},
                buffer: (() => {
                  const $0 = v.writeText(stack._1._1._1)(stack._1._1._2);
                  if (state.buffer.queue.tag === "Cons") {
                    return {...state.buffer, queue: Data$dList$dTypes.$List("Cons", Data$dList$dTypes.$List("Cons", $0, state.buffer.queue._1), state.buffer.queue._2)};
                  }
                  return {...state.buffer, buffer: $0(state.buffer.buffer)};
                })()
              };
              continue;
            }
            if (state.flexGroup.tag === "FlexGroupReset") {
              go$a0 = state.flexGroup._1.stack;
              go$a1 = {
                position: state.flexGroup._1.position,
                buffer: state.flexGroup._1.buffer,
                annotations: state.flexGroup._1.annotations,
                indentSpaces: state.flexGroup._1.indentSpaces,
                flexGroup: NoFlexGroup,
                options: state.flexGroup._1.options
              };
              continue;
            }
            go$a0 = stack._2;
            go$a1 = {
              ...state,
              position: {...state.position, column: state.position.column + stack._1._1._1 | 0},
              flexGroup: NoFlexGroup,
              buffer: (() => {
                const $0 = v.writeText(stack._1._1._1)(stack._1._1._2);
                if (state.buffer.queue.tag === "Cons") {
                  return {...state.buffer, queue: Data$dList$dTypes.$List("Cons", Data$dList$dTypes.$List("Cons", $0, state.buffer.queue._1), state.buffer.queue._2)};
                }
                return {...state.buffer, buffer: $0(state.buffer.buffer)};
              })()
            };
            continue;
          }
          if (stack._1._1.tag === "Break") {
            if (state.flexGroup.tag === "FlexGroupReset") {
              go$a0 = state.flexGroup._1.stack;
              go$a1 = {
                position: state.flexGroup._1.position,
                buffer: state.flexGroup._1.buffer,
                annotations: state.flexGroup._1.annotations,
                indentSpaces: state.flexGroup._1.indentSpaces,
                flexGroup: NoFlexGroup,
                options: state.flexGroup._1.options
              };
              continue;
            }
            go$a0 = stack._2;
            go$a1 = {
              ...state,
              position: {
                ...state.position,
                line: state.position.line + 1 | 0,
                column: 0,
                indent: state.position.nextIndent,
                ribbonWidth: calcRibbonWidth(state.options)(state.position.nextIndent)
              },
              buffer: state.buffer.queue.tag === "Cons"
                ? {...state.buffer, queue: Data$dList$dTypes.$List("Cons", Data$dList$dTypes.$List("Cons", v.writeBreak, state.buffer.queue._1), state.buffer.queue._2)}
                : {...state.buffer, buffer: v.writeBreak(state.buffer.buffer)},
              flexGroup: NoFlexGroup
            };
            continue;
          }
          if (stack._1._1.tag === "Indent") {
            if (state.position.column === 0) {
              go$a0 = Data$dList$dTypes.$List(
                "Cons",
                $DocCmd("Doc", stack._1._1._1),
                Data$dList$dTypes.$List("Cons", $DocCmd("Dedent", state.indentSpaces, state.position.nextIndent), stack._2)
              );
              go$a1 = {
                ...state,
                position: {
                  ...state.position,
                  indent: state.position.nextIndent + opts.indentWidth | 0,
                  nextIndent: state.position.nextIndent + opts.indentWidth | 0,
                  ribbonWidth: calcRibbonWidth(state.options)(state.position.nextIndent + opts.indentWidth | 0)
                },
                indentSpaces: state.indentSpaces + opts.indentUnit
              };
              continue;
            }
            go$a0 = Data$dList$dTypes.$List(
              "Cons",
              $DocCmd("Doc", stack._1._1._1),
              Data$dList$dTypes.$List("Cons", $DocCmd("Dedent", state.indentSpaces, state.position.nextIndent), stack._2)
            );
            go$a1 = {...state, position: {...state.position, nextIndent: state.position.nextIndent + opts.indentWidth | 0}, indentSpaces: state.indentSpaces + opts.indentUnit};
            continue;
          }
          if (stack._1._1.tag === "Align") {
            if (state.position.column === 0) {
              go$a0 = Data$dList$dTypes.$List(
                "Cons",
                $DocCmd("Doc", stack._1._1._2),
                Data$dList$dTypes.$List("Cons", $DocCmd("Dedent", state.indentSpaces, state.position.nextIndent), stack._2)
              );
              go$a1 = {
                ...state,
                position: {
                  ...state.position,
                  indent: state.position.nextIndent + stack._1._1._1 | 0,
                  nextIndent: state.position.nextIndent + stack._1._1._1 | 0,
                  ribbonWidth: calcRibbonWidth(state.options)(state.position.nextIndent + stack._1._1._1 | 0)
                },
                indentSpaces: state.indentSpaces + power(" ")(stack._1._1._1)
              };
              continue;
            }
            go$a0 = Data$dList$dTypes.$List(
              "Cons",
              $DocCmd("Doc", stack._1._1._2),
              Data$dList$dTypes.$List("Cons", $DocCmd("Dedent", state.indentSpaces, state.position.nextIndent), stack._2)
            );
            go$a1 = {
              ...state,
              position: {...state.position, nextIndent: state.position.nextIndent + stack._1._1._1 | 0},
              indentSpaces: state.indentSpaces + power(" ")(stack._1._1._1)
            };
            continue;
          }
          if (stack._1._1.tag === "FlexSelect") {
            if (state.flexGroup.tag === "NoFlexGroup") {
              go$a0 = Data$dList$dTypes.$List(
                "Cons",
                $DocCmd("Doc", stack._1._1._1),
                Data$dList$dTypes.$List("Cons", $DocCmd("LeaveFlexGroup", stack._1._1._2, stack._1._1._3), stack._2)
              );
              go$a1 = {...state, flexGroup: FlexGroupPending};
              continue;
            }
            if (state.flexGroup.tag === "FlexGroupPending" && state.position.ribbonWidth > 0) {
              go$a0 = Data$dList$dTypes.$List("Cons", $DocCmd("Doc", stack._1._1._1), Data$dList$dTypes.$List("Cons", $DocCmd("Doc", stack._1._1._2), stack._2));
              go$a1 = {
                ...state,
                flexGroup: $FlexGroupStatus(
                  "FlexGroupReset",
                  {position: state.position, buffer: state.buffer, annotations: state.annotations, indentSpaces: state.indentSpaces, stack, options: state.options}
                ),
                buffer: {...state.buffer, queue: Data$dList$dTypes.$List("Cons", Data$dList$dTypes.Nil, state.buffer.queue)}
              };
              continue;
            }
            go$a0 = Data$dList$dTypes.$List("Cons", $DocCmd("Doc", stack._1._1._1), Data$dList$dTypes.$List("Cons", $DocCmd("Doc", stack._1._1._2), stack._2));
            go$a1 = state;
            continue;
          }
          if (stack._1._1.tag === "FlexAlt") {
            if (state.flexGroup.tag === "FlexGroupReset") {
              go$a0 = Data$dList$dTypes.$List("Cons", $DocCmd("Doc", stack._1._1._1), stack._2);
              go$a1 = state;
              continue;
            }
            if (state.flexGroup.tag === "FlexGroupPending" && state.position.ribbonWidth > 0) {
              go$a0 = Data$dList$dTypes.$List("Cons", $DocCmd("Doc", stack._1._1._1), stack._2);
              go$a1 = {
                ...state,
                flexGroup: $FlexGroupStatus(
                  "FlexGroupReset",
                  {
                    position: state.position,
                    buffer: state.buffer,
                    annotations: state.annotations,
                    indentSpaces: state.indentSpaces,
                    stack: Data$dList$dTypes.$List("Cons", $DocCmd("Doc", stack._1._1._2), stack._2),
                    options: state.options
                  }
                ),
                buffer: {...state.buffer, queue: Data$dList$dTypes.$List("Cons", Data$dList$dTypes.Nil, state.buffer.queue)}
              };
              continue;
            }
            go$a0 = Data$dList$dTypes.$List("Cons", $DocCmd("Doc", stack._1._1._2), stack._2);
            go$a1 = state;
            continue;
          }
          if (stack._1._1.tag === "WithPosition") {
            if (state.position.column === 0 && state.position.nextIndent > 0) {
              go$a0 = Data$dList$dTypes.$List("Cons", $DocCmd("Doc", stack._1._1._1({...state.position, column: state.position.nextIndent})), stack._2);
              go$a1 = state;
              continue;
            }
            go$a0 = Data$dList$dTypes.$List("Cons", $DocCmd("Doc", stack._1._1._1(state.position)), stack._2);
            go$a1 = state;
            continue;
          }
          if (stack._1._1.tag === "Annotate") {
            go$a0 = Data$dList$dTypes.$List(
              "Cons",
              $DocCmd("Doc", stack._1._1._2),
              Data$dList$dTypes.$List("Cons", $DocCmd("LeaveAnnotation", stack._1._1._1, state.annotations), stack._2)
            );
            go$a1 = {
              ...state,
              annotations: Data$dList$dTypes.$List("Cons", stack._1._1._1, state.annotations),
              buffer: (() => {
                const $0 = v.enterAnnotation(stack._1._1._1)(state.annotations);
                if (state.buffer.queue.tag === "Cons") {
                  return {...state.buffer, queue: Data$dList$dTypes.$List("Cons", Data$dList$dTypes.$List("Cons", $0, state.buffer.queue._1), state.buffer.queue._2)};
                }
                return {...state.buffer, buffer: $0(state.buffer.buffer)};
              })()
            };
            continue;
          }
          if (stack._1._1.tag === "Local") {
            const prevOptions = {
              indent: state.position.indent,
              indentSpaces: state.indentSpaces,
              indentUnit: state.options.indentUnit,
              indentWidth: state.options.indentWidth,
              pageWidth: state.options.pageWidth,
              ribbonRatio: state.options.ribbonRatio
            };
            const v1 = stack._1._1._1(prevOptions);
            go$a0 = Data$dList$dTypes.$List("Cons", $DocCmd("Doc", v1._2), Data$dList$dTypes.$List("Cons", $DocCmd("LeaveLocal", prevOptions), stack._2));
            go$a1 = storeOptions(state.position.indent)(v1._1)(state);
            continue;
          }
          if (stack._1._1.tag === "Empty") {
            go$a0 = stack._2;
            go$a1 = state;
            continue;
          }
          $runtime.fail();
        }
        if (stack._1.tag === "LeaveFlexGroup") {
          if (state.flexGroup.tag === "NoFlexGroup") {
            go$a0 = Data$dList$dTypes.$List("Cons", $DocCmd("Doc", stack._1._2), stack._2);
            go$a1 = {...state, buffer: Dodo$dInternal$dBuffer.commit(state.buffer)};
            continue;
          }
          go$a0 = Data$dList$dTypes.$List("Cons", $DocCmd("Doc", stack._1._1), stack._2);
          go$a1 = {...state, flexGroup: NoFlexGroup, buffer: Dodo$dInternal$dBuffer.commit(state.buffer)};
          continue;
        }
        if (stack._1.tag === "Dedent") {
          go$a0 = stack._2;
          go$a1 = {...state, position: {...state.position, nextIndent: stack._1._2}, indentSpaces: stack._1._1};
          continue;
        }
        if (stack._1.tag === "LeaveAnnotation") {
          go$a0 = stack._2;
          go$a1 = {
            ...state,
            annotations: stack._1._2,
            buffer: (() => {
              const $0 = v.leaveAnnotation(stack._1._1)(stack._1._2);
              if (state.buffer.queue.tag === "Cons") {
                return {...state.buffer, queue: Data$dList$dTypes.$List("Cons", Data$dList$dTypes.$List("Cons", $0, state.buffer.queue._1), state.buffer.queue._2)};
              }
              return {...state.buffer, buffer: $0(state.buffer.buffer)};
            })()
          };
          continue;
        }
        if (stack._1.tag === "LeaveLocal") {
          go$a0 = stack._2;
          go$a1 = storeOptions(state.position.indent)(stack._1._1)(state);
          continue;
        }
      }
      $runtime.fail();
    }
    return go$r;
  };
  const $0 = {
    position: {line: 0, column: 0, indent: 0, nextIndent: 0, pageWidth: initOptions.pageWidth, ribbonWidth: calcRibbonWidth(initOptions)(0)},
    buffer: {buffer: v.emptyBuffer, queue: Data$dList$dTypes.Nil},
    annotations: Data$dList$dTypes.Nil,
    indentSpaces: "",
    flexGroup: NoFlexGroup,
    options: initOptions
  };
  return x => go(Data$dList$dTypes.$List("Cons", $DocCmd("Doc", x), Data$dList$dTypes.Nil))($0);
};
const $$break = Dodo$dInternal.Break;
const softBreak = /* #__PURE__ */ Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.Empty, Dodo$dInternal.Break);
const spaceBreak = /* #__PURE__ */ Dodo$dInternal.$Doc("FlexAlt", /* #__PURE__ */ Dodo$dInternal.$Doc("Text", 1, " "), Dodo$dInternal.Break);
const appendSpaceBreak = v => v1 => {
  if (v.tag === "Empty") { return v1; }
  if (v1.tag === "Empty") { return v; }
  const $0 = v1.tag === "Empty"
    ? Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 1, " "), Dodo$dInternal.Break)
    : Dodo$dInternal.$Doc("Append", Dodo$dInternal.$Doc("FlexAlt", Dodo$dInternal.$Doc("Text", 1, " "), Dodo$dInternal.Break), v1);
  const $1 = (() => {
    if ($0.tag === "Empty") { return Dodo$dInternal.Empty; }
    if ($0.tag === "FlexSelect" && $0._2.tag === "Empty" && $0._3.tag === "Empty") { return $0; }
    return Dodo$dInternal.$Doc("FlexSelect", $0, Dodo$dInternal.Empty, Dodo$dInternal.Empty);
  })();
  if (v.tag === "Empty") { return $1; }
  if ($1.tag === "Empty") { return v; }
  if (v.tag === "Text" && $1.tag === "Text") { return Dodo$dInternal.$Doc("Text", v._1 + $1._1 | 0, v._2 + $1._2); }
  return Dodo$dInternal.$Doc("Append", v, $1);
};
const paragraph = dictFoldable => dictFoldable.foldl(appendSpaceBreak)(Dodo$dInternal.Empty);
const textParagraph = /* #__PURE__ */ (() => {
  const $0 = Data$dFoldable.foldlArray(appendSpaceBreak)(Dodo$dInternal.Empty);
  const $1 = Data$dFunctor.arrayMap(text);
  const $2 = Data$dString$dRegex.split(Data$dString$dRegex$dUnsafe.unsafeRegex("[\\s\\n]+")(Data$dString$dRegex$dFlags.global));
  return x => $0($1($2(Data$dString$dCommon.trim(x))));
})();
const appendSpace = v => v1 => {
  if (v.tag === "Empty") { return v1; }
  if (v1.tag === "Empty") { return v; }
  const $0 = (() => {
    if (v1.tag === "Empty") { return Dodo$dInternal.$Doc("Text", 1, " "); }
    if (v1.tag === "Text") { return Dodo$dInternal.$Doc("Text", 1 + v1._1 | 0, " " + v1._2); }
    return Dodo$dInternal.$Doc("Append", Dodo$dInternal.$Doc("Text", 1, " "), v1);
  })();
  if (v.tag === "Empty") { return $0; }
  if ($0.tag === "Empty") { return v; }
  if (v.tag === "Text" && $0.tag === "Text") { return Dodo$dInternal.$Doc("Text", v._1 + $0._1 | 0, v._2 + $0._2); }
  return Dodo$dInternal.$Doc("Append", v, $0);
};
const words = dictFoldable => dictFoldable.foldr(appendSpace)(Dodo$dInternal.Empty);
const appendBreak = v => v1 => {
  if (v.tag === "Empty") { return v1; }
  if (v1.tag === "Empty") { return v; }
  const $0 = v1.tag === "Empty" ? Dodo$dInternal.Break : Dodo$dInternal.$Doc("Append", Dodo$dInternal.Break, v1);
  if (v.tag === "Empty") { return $0; }
  if ($0.tag === "Empty") { return v; }
  if (v.tag === "Text" && $0.tag === "Text") { return Dodo$dInternal.$Doc("Text", v._1 + $0._1 | 0, v._2 + $0._2); }
  return Dodo$dInternal.$Doc("Append", v, $0);
};
const lines = dictFoldable => dictFoldable.foldr(appendBreak)(Dodo$dInternal.Empty);
const annotate = x => {
  const $0 = Dodo$dInternal.Annotate(x);
  return v => {
    if (v.tag === "Empty") { return Dodo$dInternal.Empty; }
    return $0(v);
  };
};
const align = n => doc => {
  if (n > 0) {
    const $0 = Dodo$dInternal.Align(n);
    if (doc.tag === "Empty") { return Dodo$dInternal.Empty; }
    return $0(doc);
  }
  return doc;
};
const alignCurrentColumn = v => {
  if (v.tag === "Empty") { return Dodo$dInternal.Empty; }
  return Dodo$dInternal.$Doc("WithPosition", pos => align(pos.column - pos.nextIndent | 0)(v));
};
export {
  $DocCmd,
  $FlexGroupStatus,
  Dedent,
  Doc,
  FlexGroupPending,
  FlexGroupReset,
  LeaveAnnotation,
  LeaveFlexGroup,
  LeaveLocal,
  NoFlexGroup,
  Printer,
  align,
  alignCurrentColumn,
  annotate,
  appendBreak,
  appendSpace,
  appendSpaceBreak,
  $$break as break,
  calcRibbonWidth,
  enclose,
  encloseEmptyAlt,
  encloseWithSeparator,
  flexAlt,
  flexGroup,
  flexSelect,
  foldWith,
  foldWithSeparator,
  fourSpaces,
  indent,
  lines,
  locally,
  max,
  max1,
  min,
  paragraph,
  plainText,
  power,
  print,
  softBreak,
  space,
  spaceBreak,
  storeOptions,
  tabs,
  text,
  textParagraph,
  twoSpaces,
  withLocalOptions,
  withPosition,
  words
};
