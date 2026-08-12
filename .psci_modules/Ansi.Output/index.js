// | Convenience functions to simplify outputting ANSI escape codes to
// | terminals.
import * as Ansi_Codes from "../Ansi.Codes/index.js";
import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Data_List_Types from "../Data.List.Types/index.js";

// | Wrap the given text in escape codes corresponding to the given parameters.
// | For example:
// |
// | ```purescript
// | Console.log $ withGraphics (bold <> underline <> foreground BrightRed) "hello world"
// | ```
// |
// | would print "hello world" to the terminal, bold, underlined, and in bright
// | red, and then reset (so that further logging to the console uses the
// | normal color and style).
// |
// | This function simply wraps the given text in an escape code and a reset
// | code, so that it is a little more comfortable to use than the functions
// | in `Ansi.Codes`.
var withGraphics = function (params) {
    return function (text) {
        return Ansi_Codes.escapeCodeToString(new Ansi_Codes.Graphics(params)) + (text + Ansi_Codes.escapeCodeToString(new Ansi_Codes.Graphics(Control_Applicative.pure(Data_List_Types.applicativeNonEmptyList)(Ansi_Codes.Reset.value))));
    };
};
var underline = /* #__PURE__ */ (function () {
    return Control_Applicative.pure(Data_List_Types.applicativeNonEmptyList)(new Ansi_Codes.PMode(Ansi_Codes.Underline.value));
})();
var strikethrough = /* #__PURE__ */ (function () {
    return Control_Applicative.pure(Data_List_Types.applicativeNonEmptyList)(new Ansi_Codes.PMode(Ansi_Codes.Strikethrough.value));
})();
var italic = /* #__PURE__ */ (function () {
    return Control_Applicative.pure(Data_List_Types.applicativeNonEmptyList)(new Ansi_Codes.PMode(Ansi_Codes.Italic.value));
})();
var inverse = /* #__PURE__ */ (function () {
    return Control_Applicative.pure(Data_List_Types.applicativeNonEmptyList)(new Ansi_Codes.PMode(Ansi_Codes.Inverse.value));
})();
var foreground = function (c) {
    return Control_Applicative.pure(Data_List_Types.applicativeNonEmptyList)(new Ansi_Codes.PForeground(c));
};
var dim = /* #__PURE__ */ (function () {
    return Control_Applicative.pure(Data_List_Types.applicativeNonEmptyList)(new Ansi_Codes.PMode(Ansi_Codes.Dim.value));
})();
var bold = /* #__PURE__ */ (function () {
    return Control_Applicative.pure(Data_List_Types.applicativeNonEmptyList)(new Ansi_Codes.PMode(Ansi_Codes.Bold.value));
})();
var background = function (c) {
    return Control_Applicative.pure(Data_List_Types.applicativeNonEmptyList)(new Ansi_Codes.PBackground(c));
};
export {
    withGraphics,
    bold,
    dim,
    italic,
    underline,
    inverse,
    strikethrough,
    foreground,
    background
};
