import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Data_Boolean from "../Data.Boolean/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Int from "../Data.Int/index.js";
import * as Data_List_Types from "../Data.List.Types/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_String_CodePoints from "../Data.String.CodePoints/index.js";
import * as Data_String_Common from "../Data.String.Common/index.js";
import * as Data_String_Regex from "../Data.String.Regex/index.js";
import * as Data_String_Regex_Flags from "../Data.String.Regex.Flags/index.js";
import * as Data_String_Regex_Unsafe from "../Data.String.Regex.Unsafe/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Dodo_Internal from "../Dodo.Internal/index.js";
import * as Dodo_Internal_Buffer from "../Dodo.Internal.Buffer/index.js";
var mempty = /* #__PURE__ */ Data_Monoid.mempty(Dodo_Internal.monoidDoc);
var pure = /* #__PURE__ */ Control_Applicative.pure(Data_List_Types.applicativeList);

// | Custom printers can be used to render richer documents than just plain
// | text.
// | * `emptyBuffer` - The initial buffer.
// | * `writeText` - Should write a string with the given width to the buffer.
// | * `writeIndent` - Should write indentation with the given width to the buffer.
// | * `writeBreak` - Should write a line break to the buffer.
// | * `enterAnnotation` - Called when entering a new annotated region. Provides the full annotation stack.
// | * `leaveAnnotation` - Called when leaving an annotated region. Provides the full annotation stack.
// | * `flushBuffer` - Called at the end of the document to get the final result.
var Printer = function (x) {
    return x;
};
var Doc = /* #__PURE__ */ (function () {
    function Doc(value0) {
        this.value0 = value0;
    };
    Doc.create = function (value0) {
        return new Doc(value0);
    };
    return Doc;
})();
var Dedent = /* #__PURE__ */ (function () {
    function Dedent(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    Dedent.create = function (value0) {
        return function (value1) {
            return new Dedent(value0, value1);
        };
    };
    return Dedent;
})();
var LeaveAnnotation = /* #__PURE__ */ (function () {
    function LeaveAnnotation(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    LeaveAnnotation.create = function (value0) {
        return function (value1) {
            return new LeaveAnnotation(value0, value1);
        };
    };
    return LeaveAnnotation;
})();
var LeaveFlexGroup = /* #__PURE__ */ (function () {
    function LeaveFlexGroup(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    LeaveFlexGroup.create = function (value0) {
        return function (value1) {
            return new LeaveFlexGroup(value0, value1);
        };
    };
    return LeaveFlexGroup;
})();
var LeaveLocal = /* #__PURE__ */ (function () {
    function LeaveLocal(value0) {
        this.value0 = value0;
    };
    LeaveLocal.create = function (value0) {
        return new LeaveLocal(value0);
    };
    return LeaveLocal;
})();
var NoFlexGroup = /* #__PURE__ */ (function () {
    function NoFlexGroup() {

    };
    NoFlexGroup.value = new NoFlexGroup();
    return NoFlexGroup;
})();
var FlexGroupPending = /* #__PURE__ */ (function () {
    function FlexGroupPending() {

    };
    FlexGroupPending.value = new FlexGroupPending();
    return FlexGroupPending;
})();
var FlexGroupReset = /* #__PURE__ */ (function () {
    function FlexGroupReset(value0) {
        this.value0 = value0;
    };
    FlexGroupReset.create = function (value0) {
        return new FlexGroupReset(value0);
    };
    return FlexGroupReset;
})();

// | Build a document based on the current layout position.
var withPosition = /* #__PURE__ */ (function () {
    return Dodo_Internal.WithPosition.create;
})();

// | *EXPERIMENTAL:* modifies printing state and options locally for a document.
// | This may change or be removed at any time. Differs from `locally` in that the
// | document can be responsive to options.
var withLocalOptions = /* #__PURE__ */ (function () {
    return Dodo_Internal.Local.create;
})();

// | Prints 2-space indents, with a default 80-column page width.
var twoSpaces = {
    pageWidth: 80,
    ribbonRatio: 1.0,
    indentUnit: "  ",
    indentWidth: 2
};

// | The most basic document leaf. This should not contain newlines. If it does
// | your document will look very funny.
var text = function (v) {
    if (v === "") {
        return Dodo_Internal.Empty.value;
    };
    return new Dodo_Internal.Text(Data_String_CodePoints.length(v), v);
};

// | Prints tab indents (4-wide), with a default 120-column page width.
var tabs = {
    pageWidth: 120,
    ribbonRatio: 1.0,
    indentUnit: "\x09",
    indentWidth: 4
};
var storeState = function (stack) {
    return function (v) {
        return {
            position: v.position,
            buffer: v.buffer,
            annotations: v.annotations,
            indentSpaces: v.indentSpaces,
            stack: stack,
            options: v.options
        };
    };
};

// | A singe space character.
var space = /* #__PURE__ */ text(" ");
var resetState = function (v) {
    return {
        position: v.position,
        buffer: v.buffer,
        annotations: v.annotations,
        indentSpaces: v.indentSpaces,
        flexGroup: NoFlexGroup.value,
        options: v.options
    };
};

// | A plain text printer. Can be used with any document.
var plainText = {
    emptyBuffer: "",
    writeText: function (v) {
        return function (str) {
            return function (buff) {
                return buff + str;
            };
        };
    },
    writeIndent: function (v) {
        return function (str) {
            return function (buff) {
                return buff + str;
            };
        };
    },
    writeBreak: function (buff) {
        return buff + "\x0a";
    },
    enterAnnotation: function (v) {
        return function (v1) {
            return function (buff) {
                return buff;
            };
        };
    },
    leaveAnnotation: function (v) {
        return function (v1) {
            return function (buff) {
                return buff;
            };
        };
    },
    flushBuffer: function (buff) {
        return buff;
    }
};

// | *EXPERIMENTAL:* modifies printing state and options locally for a document.
// | This may change or be removed at any time.
var locally = function (k) {
    return function (doc) {
        return new Dodo_Internal.Local(function (options) {
            return new Data_Tuple.Tuple(k(options), doc);
        });
    };
};

// | Increases the indentation level by one indent.
var indent = /* #__PURE__ */ (function () {
    return Dodo_Internal.notEmpty(Dodo_Internal.Indent.create);
})();

// | Prints 4-space indents, with a default 120-column page width.
var fourSpaces = {
    pageWidth: 120,
    ribbonRatio: 1.0,
    indentUnit: "    ",
    indentWidth: 4
};

// | Appends a series of documents together with a given append function. This
// | is notable because it ignores empty documents.
var foldWith = function (dictFoldable) {
    return function (f) {
        return Data_Foldable.foldr(dictFoldable)(Dodo_Internal.bothNotEmpty(f))(mempty);
    };
};

// | Appends a series of documents together with a separator in between them.
var foldWithSeparator = function (dictFoldable) {
    return function (separator) {
        return foldWith(dictFoldable)(function (a) {
            return function (b) {
                return Data_Semigroup.append(Dodo_Internal.semigroupDoc)(a)(Data_Semigroup.append(Dodo_Internal.semigroupDoc)(separator)(b));
            };
        });
    };
};

// | Attempts to layout the first document with flex alternatives, falling
// | back to defaults if it doesn't fit the page width. If the flex alternatives
// | are used then the second document will be appended, otherwise the third
// | document will be appended.
var flexSelect = function (doc1) {
    return function (doc2) {
        return function (doc3) {
            if (Dodo_Internal.isEmpty(doc1)) {
                return doc2;
            };
            if (Data_Boolean.otherwise) {
                return new Dodo_Internal.FlexSelect(doc1, doc2, doc3);
            };
            throw new Error("Failed pattern match at Dodo (line 89, column 1 - line 89, column 57): " + [ doc1.constructor.name, doc2.constructor.name, doc3.constructor.name ]);
        };
    };
};

// | Attempts to layout the document with flex alternatives, falling back
// | to defaults if it doesn't fit the page width.
var flexGroup = /* #__PURE__ */ Dodo_Internal.notEmpty(function (v) {
    if (v instanceof Dodo_Internal.FlexSelect && (Dodo_Internal.isEmpty(v.value1) && Dodo_Internal.isEmpty(v.value2))) {
        return v;
    };
    return new Dodo_Internal.FlexSelect(v, Dodo_Internal.Empty.value, Dodo_Internal.Empty.value);
});

// | Attempts to layout the first document when in a flex group, falling back
// | to the second as a default.
var flexAlt = /* #__PURE__ */ (function () {
    return Dodo_Internal.FlexAlt.create;
})();

// | Uses an opening and closing document, as a well as a separator, to render
// | a series of documents.
// | ```purescript`
// | example = encloseWithSeparator (text "[") (text "]") (",") [ text "one", text "two" ]
// | ```
var encloseWithSeparator = function (dictFoldable) {
    return function (open) {
        return function (close) {
            return function (separator) {
                return function (inner) {
                    return Data_Semigroup.append(Dodo_Internal.semigroupDoc)(open)(Data_Semigroup.append(Dodo_Internal.semigroupDoc)(foldWithSeparator(dictFoldable)(separator)(inner))(close));
                };
            };
        };
    };
};

// | Uses an opening and closing document to wrap another document, falling
// | back when the inner document is empty.
// | ```purescript
// | example = encloseEmptyAlt (text "[ ") (text " ]") (text "[]") mempty
// | ```
var encloseEmptyAlt = function (open) {
    return function (close) {
        return function ($$default) {
            return function (inner) {
                if (Dodo_Internal.isEmpty(inner)) {
                    return $$default;
                };
                if (Data_Boolean.otherwise) {
                    return Data_Semigroup.append(Dodo_Internal.semigroupDoc)(open)(Data_Semigroup.append(Dodo_Internal.semigroupDoc)(inner)(close));
                };
                throw new Error("Failed pattern match at Dodo (line 176, column 1 - line 176, column 71): " + [ open.constructor.name, close.constructor.name, $$default.constructor.name, inner.constructor.name ]);
            };
        };
    };
};

// | Uses an opening and closing document to wrap another document.
// | ```purescript
// | example = enclose (text "(") (text ")") (text "inner")
// | ```
var enclose = function (open) {
    return function (close) {
        return function (inner) {
            return Data_Semigroup.append(Dodo_Internal.semigroupDoc)(open)(Data_Semigroup.append(Dodo_Internal.semigroupDoc)(inner)(close));
        };
    };
};
var calcRibbonWidth = function (v) {
    return function (n) {
        return Data_Ord.max(Data_Ord.ordInt)(0)(Data_Int.ceil(v.ribbonRatio * Data_Int.toNumber(v.pageWidth - n | 0)));
    };
};
var storeOptions = function (prevIndent) {
    return function (localOptions) {
        return function (state) {
            var newOptions = {
                indentUnit: localOptions.indentUnit,
                indentWidth: localOptions.indentWidth,
                pageWidth: localOptions.pageWidth,
                ribbonRatio: localOptions.ribbonRatio
            };
            return {
                buffer: state.buffer,
                annotations: state.annotations,
                flexGroup: state.flexGroup,
                indentSpaces: localOptions.indentSpaces,
                options: newOptions,
                position: {
                    line: state.position.line,
                    column: state.position.column,
                    indent: state.position.indent,
                    pageWidth: newOptions.pageWidth,
                    ribbonWidth: calcRibbonWidth(newOptions)(prevIndent),
                    nextIndent: localOptions.indent
                }
            };
        };
    };
};

// | Prints a documents given a printer and print options.
// |
// | ```purescript
// | print plainText twoSpaces myDoc
// | ```
// |
// | This will use full line-lookahead from the start of a flex group. If it
// | encounters a break or content overflows the page-width, it will layout
// | the group using flex alternative defaults instead.
var print = function (v) {
    return function (opts) {
        var initOptions = {
            pageWidth: opts.pageWidth,
            indentUnit: opts.indentUnit,
            indentWidth: opts.indentWidth,
            ribbonRatio: Data_Ord.max(Data_Ord.ordNumber)(0.0)(Data_Ord.min(Data_Ord.ordNumber)(1.0)(opts.ribbonRatio))
        };
        var initState = {
            position: {
                line: 0,
                column: 0,
                indent: 0,
                nextIndent: 0,
                pageWidth: initOptions.pageWidth,
                ribbonWidth: calcRibbonWidth(initOptions)(0)
            },
            buffer: Dodo_Internal_Buffer["new"](v.emptyBuffer),
            annotations: Data_List_Types.Nil.value,
            indentSpaces: "",
            flexGroup: NoFlexGroup.value,
            options: initOptions
        };
        var go = function ($copy_stack) {
            return function ($copy_state) {
                var $tco_var_stack = $copy_stack;
                var $tco_done = false;
                var $tco_result;
                function $tco_loop(stack, state) {
                    if (stack instanceof Data_List_Types.Nil) {
                        $tco_done = true;
                        return v.flushBuffer(Dodo_Internal_Buffer.get(state.buffer));
                    };
                    if (stack instanceof Data_List_Types.Cons) {
                        if (stack.value0 instanceof Doc) {
                            if (stack.value0.value0 instanceof Dodo_Internal.Append) {
                                $tco_var_stack = new Data_List_Types.Cons(new Doc(stack.value0.value0.value0), new Data_List_Types.Cons(new Doc(stack.value0.value0.value1), stack.value1));
                                $copy_state = state;
                                return;
                            };
                            if (stack.value0.value0 instanceof Dodo_Internal.Text) {
                                if (state.position.column === 0 && state.position.indent > 0) {
                                    $tco_var_stack = stack;
                                    $copy_state = {
                                        annotations: state.annotations,
                                        indentSpaces: state.indentSpaces,
                                        flexGroup: state.flexGroup,
                                        options: state.options,
                                        position: {
                                            line: state.position.line,
                                            indent: state.position.indent,
                                            nextIndent: state.position.nextIndent,
                                            pageWidth: state.position.pageWidth,
                                            ribbonWidth: state.position.ribbonWidth,
                                            column: state.position.indent
                                        },
                                        buffer: Dodo_Internal_Buffer.modify(v.writeIndent(state.position.indent)(state.indentSpaces))(state.buffer)
                                    };
                                    return;
                                };
                                if ((state.position.column + stack.value0.value0.value0 | 0) <= (state.position.indent + state.position.ribbonWidth | 0)) {
                                    $tco_var_stack = stack.value1;
                                    $copy_state = {
                                        annotations: state.annotations,
                                        indentSpaces: state.indentSpaces,
                                        flexGroup: state.flexGroup,
                                        options: state.options,
                                        position: {
                                            line: state.position.line,
                                            indent: state.position.indent,
                                            nextIndent: state.position.nextIndent,
                                            pageWidth: state.position.pageWidth,
                                            ribbonWidth: state.position.ribbonWidth,
                                            column: state.position.column + stack.value0.value0.value0 | 0
                                        },
                                        buffer: Dodo_Internal_Buffer.modify(v.writeText(stack.value0.value0.value0)(stack.value0.value0.value1))(state.buffer)
                                    };
                                    return;
                                };
                                if (Data_Boolean.otherwise) {
                                    if (state.flexGroup instanceof FlexGroupReset) {
                                        $tco_var_stack = state.flexGroup.value0.stack;
                                        $copy_state = resetState(state.flexGroup.value0);
                                        return;
                                    };
                                    $tco_var_stack = stack.value1;
                                    $copy_state = {
                                        annotations: state.annotations,
                                        indentSpaces: state.indentSpaces,
                                        options: state.options,
                                        position: {
                                            line: state.position.line,
                                            indent: state.position.indent,
                                            nextIndent: state.position.nextIndent,
                                            pageWidth: state.position.pageWidth,
                                            ribbonWidth: state.position.ribbonWidth,
                                            column: state.position.column + stack.value0.value0.value0 | 0
                                        },
                                        flexGroup: NoFlexGroup.value,
                                        buffer: Dodo_Internal_Buffer.modify(v.writeText(stack.value0.value0.value0)(stack.value0.value0.value1))(state.buffer)
                                    };
                                    return;
                                };
                            };
                            if (stack.value0.value0 instanceof Dodo_Internal.Break) {
                                if (state.flexGroup instanceof FlexGroupReset) {
                                    $tco_var_stack = state.flexGroup.value0.stack;
                                    $copy_state = resetState(state.flexGroup.value0);
                                    return;
                                };
                                $tco_var_stack = stack.value1;
                                $copy_state = {
                                    annotations: state.annotations,
                                    indentSpaces: state.indentSpaces,
                                    options: state.options,
                                    position: {
                                        nextIndent: state.position.nextIndent,
                                        pageWidth: state.position.pageWidth,
                                        line: state.position.line + 1 | 0,
                                        column: 0,
                                        indent: state.position.nextIndent,
                                        ribbonWidth: calcRibbonWidth(state.options)(state.position.nextIndent)
                                    },
                                    buffer: Dodo_Internal_Buffer.modify(v.writeBreak)(state.buffer),
                                    flexGroup: NoFlexGroup.value
                                };
                                return;
                            };
                            if (stack.value0.value0 instanceof Dodo_Internal.Indent) {
                                if (state.position.column === 0) {
                                    $tco_var_stack = new Data_List_Types.Cons(new Doc(stack.value0.value0.value0), new Data_List_Types.Cons(new Dedent(state.indentSpaces, state.position.nextIndent), stack.value1));
                                    $copy_state = {
                                        buffer: state.buffer,
                                        annotations: state.annotations,
                                        flexGroup: state.flexGroup,
                                        options: state.options,
                                        position: {
                                            line: state.position.line,
                                            column: state.position.column,
                                            pageWidth: state.position.pageWidth,
                                            indent: state.position.nextIndent + opts.indentWidth | 0,
                                            nextIndent: state.position.nextIndent + opts.indentWidth | 0,
                                            ribbonWidth: calcRibbonWidth(state.options)(state.position.nextIndent + opts.indentWidth | 0)
                                        },
                                        indentSpaces: state.indentSpaces + opts.indentUnit
                                    };
                                    return;
                                };
                                if (Data_Boolean.otherwise) {
                                    $tco_var_stack = new Data_List_Types.Cons(new Doc(stack.value0.value0.value0), new Data_List_Types.Cons(new Dedent(state.indentSpaces, state.position.nextIndent), stack.value1));
                                    $copy_state = {
                                        buffer: state.buffer,
                                        annotations: state.annotations,
                                        flexGroup: state.flexGroup,
                                        options: state.options,
                                        position: {
                                            line: state.position.line,
                                            column: state.position.column,
                                            indent: state.position.indent,
                                            pageWidth: state.position.pageWidth,
                                            ribbonWidth: state.position.ribbonWidth,
                                            nextIndent: state.position.nextIndent + opts.indentWidth | 0
                                        },
                                        indentSpaces: state.indentSpaces + opts.indentUnit
                                    };
                                    return;
                                };
                            };
                            if (stack.value0.value0 instanceof Dodo_Internal.Align) {
                                if (state.position.column === 0) {
                                    $tco_var_stack = new Data_List_Types.Cons(new Doc(stack.value0.value0.value1), new Data_List_Types.Cons(new Dedent(state.indentSpaces, state.position.nextIndent), stack.value1));
                                    $copy_state = {
                                        buffer: state.buffer,
                                        annotations: state.annotations,
                                        flexGroup: state.flexGroup,
                                        options: state.options,
                                        position: {
                                            line: state.position.line,
                                            column: state.position.column,
                                            pageWidth: state.position.pageWidth,
                                            indent: state.position.nextIndent + stack.value0.value0.value0 | 0,
                                            nextIndent: state.position.nextIndent + stack.value0.value0.value0 | 0,
                                            ribbonWidth: calcRibbonWidth(state.options)(state.position.nextIndent + stack.value0.value0.value0 | 0)
                                        },
                                        indentSpaces: state.indentSpaces + Data_Monoid.power(Data_Monoid.monoidString)(" ")(stack.value0.value0.value0)
                                    };
                                    return;
                                };
                                if (Data_Boolean.otherwise) {
                                    $tco_var_stack = new Data_List_Types.Cons(new Doc(stack.value0.value0.value1), new Data_List_Types.Cons(new Dedent(state.indentSpaces, state.position.nextIndent), stack.value1));
                                    $copy_state = {
                                        buffer: state.buffer,
                                        annotations: state.annotations,
                                        flexGroup: state.flexGroup,
                                        options: state.options,
                                        position: {
                                            line: state.position.line,
                                            column: state.position.column,
                                            indent: state.position.indent,
                                            pageWidth: state.position.pageWidth,
                                            ribbonWidth: state.position.ribbonWidth,
                                            nextIndent: state.position.nextIndent + stack.value0.value0.value0 | 0
                                        },
                                        indentSpaces: state.indentSpaces + Data_Monoid.power(Data_Monoid.monoidString)(" ")(stack.value0.value0.value0)
                                    };
                                    return;
                                };
                            };
                            if (stack.value0.value0 instanceof Dodo_Internal.FlexSelect) {
                                if (state.flexGroup instanceof NoFlexGroup) {
                                    $tco_var_stack = new Data_List_Types.Cons(new Doc(stack.value0.value0.value0), new Data_List_Types.Cons(new LeaveFlexGroup(stack.value0.value0.value1, stack.value0.value0.value2), stack.value1));
                                    $copy_state = {
                                        position: state.position,
                                        buffer: state.buffer,
                                        annotations: state.annotations,
                                        indentSpaces: state.indentSpaces,
                                        options: state.options,
                                        flexGroup: FlexGroupPending.value
                                    };
                                    return;
                                };
                                if (state.flexGroup instanceof FlexGroupPending && state.position.ribbonWidth > 0) {
                                    $tco_var_stack = new Data_List_Types.Cons(new Doc(stack.value0.value0.value0), new Data_List_Types.Cons(new Doc(stack.value0.value0.value1), stack.value1));
                                    $copy_state = {
                                        position: state.position,
                                        annotations: state.annotations,
                                        indentSpaces: state.indentSpaces,
                                        options: state.options,
                                        flexGroup: new FlexGroupReset(storeState(stack)(state)),
                                        buffer: Dodo_Internal_Buffer.branch(state.buffer)
                                    };
                                    return;
                                };
                                $tco_var_stack = new Data_List_Types.Cons(new Doc(stack.value0.value0.value0), new Data_List_Types.Cons(new Doc(stack.value0.value0.value1), stack.value1));
                                $copy_state = state;
                                return;
                            };
                            if (stack.value0.value0 instanceof Dodo_Internal.FlexAlt) {
                                if (state.flexGroup instanceof FlexGroupReset) {
                                    $tco_var_stack = new Data_List_Types.Cons(new Doc(stack.value0.value0.value0), stack.value1);
                                    $copy_state = state;
                                    return;
                                };
                                if (state.flexGroup instanceof FlexGroupPending && state.position.ribbonWidth > 0) {
                                    $tco_var_stack = new Data_List_Types.Cons(new Doc(stack.value0.value0.value0), stack.value1);
                                    $copy_state = {
                                        position: state.position,
                                        annotations: state.annotations,
                                        indentSpaces: state.indentSpaces,
                                        options: state.options,
                                        flexGroup: new FlexGroupReset(storeState(new Data_List_Types.Cons(new Doc(stack.value0.value0.value1), stack.value1))(state)),
                                        buffer: Dodo_Internal_Buffer.branch(state.buffer)
                                    };
                                    return;
                                };
                                $tco_var_stack = new Data_List_Types.Cons(new Doc(stack.value0.value0.value1), stack.value1);
                                $copy_state = state;
                                return;
                            };
                            if (stack.value0.value0 instanceof Dodo_Internal.WithPosition) {
                                if (state.position.column === 0 && state.position.nextIndent > 0) {
                                    $tco_var_stack = new Data_List_Types.Cons(new Doc(stack.value0.value0.value0({
                                        line: state.position.line,
                                        indent: state.position.indent,
                                        nextIndent: state.position.nextIndent,
                                        pageWidth: state.position.pageWidth,
                                        ribbonWidth: state.position.ribbonWidth,
                                        column: state.position.nextIndent
                                    })), stack.value1);
                                    $copy_state = state;
                                    return;
                                };
                                if (Data_Boolean.otherwise) {
                                    $tco_var_stack = new Data_List_Types.Cons(new Doc(stack.value0.value0.value0(state.position)), stack.value1);
                                    $copy_state = state;
                                    return;
                                };
                            };
                            if (stack.value0.value0 instanceof Dodo_Internal.Annotate) {
                                $tco_var_stack = new Data_List_Types.Cons(new Doc(stack.value0.value0.value1), new Data_List_Types.Cons(new LeaveAnnotation(stack.value0.value0.value0, state.annotations), stack.value1));
                                $copy_state = {
                                    position: state.position,
                                    indentSpaces: state.indentSpaces,
                                    flexGroup: state.flexGroup,
                                    options: state.options,
                                    annotations: new Data_List_Types.Cons(stack.value0.value0.value0, state.annotations),
                                    buffer: Dodo_Internal_Buffer.modify(v.enterAnnotation(stack.value0.value0.value0)(state.annotations))(state.buffer)
                                };
                                return;
                            };
                            if (stack.value0.value0 instanceof Dodo_Internal.Local) {
                                var prevOptions = {
                                    indent: state.position.indent,
                                    indentSpaces: state.indentSpaces,
                                    indentUnit: state.options.indentUnit,
                                    indentWidth: state.options.indentWidth,
                                    pageWidth: state.options.pageWidth,
                                    ribbonRatio: state.options.ribbonRatio
                                };
                                var v1 = stack.value0.value0.value0(prevOptions);
                                $tco_var_stack = new Data_List_Types.Cons(new Doc(v1.value1), new Data_List_Types.Cons(new LeaveLocal(prevOptions), stack.value1));
                                $copy_state = storeOptions(state.position.indent)(v1.value0)(state);
                                return;
                            };
                            if (stack.value0.value0 instanceof Dodo_Internal.Empty) {
                                $tco_var_stack = stack.value1;
                                $copy_state = state;
                                return;
                            };
                            throw new Error("Failed pattern match at Dodo (line 365, column 18 - line 477, column 23): " + [ stack.value0.value0.constructor.name ]);
                        };
                        if (stack.value0 instanceof LeaveFlexGroup) {
                            if (state.flexGroup instanceof NoFlexGroup) {
                                $tco_var_stack = new Data_List_Types.Cons(new Doc(stack.value0.value1), stack.value1);
                                $copy_state = {
                                    position: state.position,
                                    annotations: state.annotations,
                                    indentSpaces: state.indentSpaces,
                                    flexGroup: state.flexGroup,
                                    options: state.options,
                                    buffer: Dodo_Internal_Buffer.commit(state.buffer)
                                };
                                return;
                            };
                            $tco_var_stack = new Data_List_Types.Cons(new Doc(stack.value0.value0), stack.value1);
                            $copy_state = {
                                position: state.position,
                                annotations: state.annotations,
                                indentSpaces: state.indentSpaces,
                                options: state.options,
                                flexGroup: NoFlexGroup.value,
                                buffer: Dodo_Internal_Buffer.commit(state.buffer)
                            };
                            return;
                        };
                        if (stack.value0 instanceof Dedent) {
                            $tco_var_stack = stack.value1;
                            $copy_state = {
                                buffer: state.buffer,
                                annotations: state.annotations,
                                flexGroup: state.flexGroup,
                                options: state.options,
                                position: {
                                    line: state.position.line,
                                    column: state.position.column,
                                    indent: state.position.indent,
                                    pageWidth: state.position.pageWidth,
                                    ribbonWidth: state.position.ribbonWidth,
                                    nextIndent: stack.value0.value1
                                },
                                indentSpaces: stack.value0.value0
                            };
                            return;
                        };
                        if (stack.value0 instanceof LeaveAnnotation) {
                            $tco_var_stack = stack.value1;
                            $copy_state = {
                                position: state.position,
                                indentSpaces: state.indentSpaces,
                                flexGroup: state.flexGroup,
                                options: state.options,
                                annotations: stack.value0.value1,
                                buffer: Dodo_Internal_Buffer.modify(v.leaveAnnotation(stack.value0.value0)(stack.value0.value1))(state.buffer)
                            };
                            return;
                        };
                        if (stack.value0 instanceof LeaveLocal) {
                            $tco_var_stack = stack.value1;
                            $copy_state = storeOptions(state.position.indent)(stack.value0.value0)(state);
                            return;
                        };
                        throw new Error("Failed pattern match at Dodo (line 364, column 18 - line 499, column 70): " + [ stack.value0.constructor.name ]);
                    };
                    throw new Error("Failed pattern match at Dodo (line 361, column 20 - line 499, column 70): " + [ stack.constructor.name ]);
                };
                while (!$tco_done) {
                    $tco_result = $tco_loop($tco_var_stack, $copy_state);
                };
                return $tco_result;
            };
        };
        var $114 = Data_Function.flip(go)(initState);
        return function ($115) {
            return $114(pure(Doc.create($115)));
        };
    };
};

// | Inserts a hard line break.
var $$break = /* #__PURE__ */ (function () {
    return Dodo_Internal.Break.value;
})();

// | Inserts nothing when in a flex group, otherwise inserts a break.
var softBreak = /* #__PURE__ */ flexAlt(mempty)($$break);

// | Inserts a space when in a flex group, otherwise inserts a break.
var spaceBreak = /* #__PURE__ */ flexAlt(space)($$break);

// | Appends two documents with a space between them, falling back to a
// | break if that does not fit.
var appendSpaceBreak = /* #__PURE__ */ Dodo_Internal.bothNotEmpty(function (a) {
    return function (b) {
        return Data_Semigroup.append(Dodo_Internal.semigroupDoc)(a)(flexGroup(Data_Semigroup.append(Dodo_Internal.semigroupDoc)(spaceBreak)(b)));
    };
});

// | Appends documents with a space-break in between them.
var paragraph = function (dictFoldable) {
    return Data_Foldable.foldl(dictFoldable)(appendSpaceBreak)(Dodo_Internal.Empty.value);
};

// | Constructs a wrapping paragraph from a blob of text. Ignores newlines and
// | multiple spaces.
var textParagraph = /* #__PURE__ */ (function () {
    var spaceRegex = Data_String_Regex_Unsafe.unsafeRegex("[\\s\\n]+")(Data_String_Regex_Flags.global);
    var $116 = paragraph(Data_Foldable.foldableArray);
    var $117 = Data_Functor.map(Data_Functor.functorArray)(text);
    var $118 = Data_String_Regex.split(spaceRegex);
    return function ($119) {
        return $116($117($118(Data_String_Common.trim($119))));
    };
})();

// | Appends two documents with a space between them.
var appendSpace = /* #__PURE__ */ Dodo_Internal.bothNotEmpty(function (a) {
    return function (b) {
        return Data_Semigroup.append(Dodo_Internal.semigroupDoc)(a)(Data_Semigroup.append(Dodo_Internal.semigroupDoc)(space)(b));
    };
});

// | Appends documents with a space in between them.
var words = function (dictFoldable) {
    return Data_Foldable.foldr(dictFoldable)(appendSpace)(Dodo_Internal.Empty.value);
};

// | Appends two documents with a break between them.
var appendBreak = /* #__PURE__ */ Dodo_Internal.bothNotEmpty(function (a) {
    return function (b) {
        return Data_Semigroup.append(Dodo_Internal.semigroupDoc)(a)(Data_Semigroup.append(Dodo_Internal.semigroupDoc)($$break)(b));
    };
});

// | Appends documents with a break in between them.
var lines = function (dictFoldable) {
    return Data_Foldable.foldr(dictFoldable)(appendBreak)(Dodo_Internal.Empty.value);
};

// | Adds an annotation to a document. Printers can interpret annotations to style
// | their output, eg. ANSI colors.
var annotate = function ($120) {
    return Dodo_Internal.notEmpty(Dodo_Internal.Annotate.create($120));
};

// | Increases the indentation level by the number of spaces (for alignment purposes).
var align = function (n) {
    return function (doc) {
        if (n > 0) {
            return Dodo_Internal.notEmpty(Dodo_Internal.Align.create(n))(doc);
        };
        if (Data_Boolean.otherwise) {
            return doc;
        };
        throw new Error("Failed pattern match at Dodo (line 64, column 1 - line 64, column 41): " + [ n.constructor.name, doc.constructor.name ]);
    };
};

// | Increases the indentation level so that it aligns to the current column.
var alignCurrentColumn = /* #__PURE__ */ Dodo_Internal.notEmpty(function (doc) {
    return withPosition(function (pos) {
        return align(pos.column - pos.nextIndent | 0)(doc);
    });
});
export {
    indent,
    align,
    alignCurrentColumn,
    annotate,
    withPosition,
    text,
    $$break as break,
    spaceBreak,
    softBreak,
    space,
    lines,
    words,
    appendBreak,
    appendSpace,
    appendSpaceBreak,
    flexAlt,
    flexGroup,
    flexSelect,
    paragraph,
    textParagraph,
    enclose,
    encloseEmptyAlt,
    encloseWithSeparator,
    foldWithSeparator,
    foldWith,
    locally,
    withLocalOptions,
    print,
    Printer,
    plainText,
    twoSpaces,
    fourSpaces,
    tabs
};
export {
    bothNotEmpty,
    isEmpty,
    notEmpty
} from "../Dodo.Internal/index.js";
