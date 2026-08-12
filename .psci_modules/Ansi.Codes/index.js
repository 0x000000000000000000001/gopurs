// | This module defines a data type representing ANSI escape codes, as well as
// | functions for serialising them as Strings.
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Generic_Rep from "../Data.Generic.Rep/index.js";
import * as Data_List_Types from "../Data.List.Types/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Ordering from "../Data.Ordering/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Data_Show_Generic from "../Data.Show.Generic/index.js";
var genericShowConstructor = /* #__PURE__ */ Data_Show_Generic.genericShowConstructor(Data_Show_Generic.genericShowArgsNoArguments);
var genericShowArgsArgument = /* #__PURE__ */ Data_Show_Generic.genericShowArgsArgument(Data_Show.showInt);
var genericShowConstructor1 = /* #__PURE__ */ Data_Show_Generic.genericShowConstructor(genericShowArgsArgument);
var Bold = /* #__PURE__ */ (function () {
    function Bold() {

    };
    Bold.value = new Bold();
    return Bold;
})();
var Dim = /* #__PURE__ */ (function () {
    function Dim() {

    };
    Dim.value = new Dim();
    return Dim;
})();
var Italic = /* #__PURE__ */ (function () {
    function Italic() {

    };
    Italic.value = new Italic();
    return Italic;
})();
var Underline = /* #__PURE__ */ (function () {
    function Underline() {

    };
    Underline.value = new Underline();
    return Underline;
})();
var Inverse = /* #__PURE__ */ (function () {
    function Inverse() {

    };
    Inverse.value = new Inverse();
    return Inverse;
})();
var Strikethrough = /* #__PURE__ */ (function () {
    function Strikethrough() {

    };
    Strikethrough.value = new Strikethrough();
    return Strikethrough;
})();

// | Specifies how much text to erase.
// |
// | * ToEnd: erase from the cursor to the end of the line or screen.
// | * FromBeginning: erase to the cursor from the beginning of the line or
// |    screen.
// | * Entire: erase the entire line or screen.
var ToEnd = /* #__PURE__ */ (function () {
    function ToEnd() {

    };
    ToEnd.value = new ToEnd();
    return ToEnd;
})();

// | Specifies how much text to erase.
// |
// | * ToEnd: erase from the cursor to the end of the line or screen.
// | * FromBeginning: erase to the cursor from the beginning of the line or
// |    screen.
// | * Entire: erase the entire line or screen.
var FromBeginning = /* #__PURE__ */ (function () {
    function FromBeginning() {

    };
    FromBeginning.value = new FromBeginning();
    return FromBeginning;
})();

// | Specifies how much text to erase.
// |
// | * ToEnd: erase from the cursor to the end of the line or screen.
// | * FromBeginning: erase to the cursor from the beginning of the line or
// |    screen.
// | * Entire: erase the entire line or screen.
var Entire = /* #__PURE__ */ (function () {
    function Entire() {

    };
    Entire.value = new Entire();
    return Entire;
})();

// | The standard set of 16 ANSI colors.
var Black = /* #__PURE__ */ (function () {
    function Black() {

    };
    Black.value = new Black();
    return Black;
})();

// | The standard set of 16 ANSI colors.
var Red = /* #__PURE__ */ (function () {
    function Red() {

    };
    Red.value = new Red();
    return Red;
})();

// | The standard set of 16 ANSI colors.
var Green = /* #__PURE__ */ (function () {
    function Green() {

    };
    Green.value = new Green();
    return Green;
})();

// | The standard set of 16 ANSI colors.
var Yellow = /* #__PURE__ */ (function () {
    function Yellow() {

    };
    Yellow.value = new Yellow();
    return Yellow;
})();

// | The standard set of 16 ANSI colors.
var Blue = /* #__PURE__ */ (function () {
    function Blue() {

    };
    Blue.value = new Blue();
    return Blue;
})();

// | The standard set of 16 ANSI colors.
var Magenta = /* #__PURE__ */ (function () {
    function Magenta() {

    };
    Magenta.value = new Magenta();
    return Magenta;
})();

// | The standard set of 16 ANSI colors.
var Cyan = /* #__PURE__ */ (function () {
    function Cyan() {

    };
    Cyan.value = new Cyan();
    return Cyan;
})();

// | The standard set of 16 ANSI colors.
var White = /* #__PURE__ */ (function () {
    function White() {

    };
    White.value = new White();
    return White;
})();

// | The standard set of 16 ANSI colors.
var BrightBlack = /* #__PURE__ */ (function () {
    function BrightBlack() {

    };
    BrightBlack.value = new BrightBlack();
    return BrightBlack;
})();

// | The standard set of 16 ANSI colors.
var BrightRed = /* #__PURE__ */ (function () {
    function BrightRed() {

    };
    BrightRed.value = new BrightRed();
    return BrightRed;
})();

// | The standard set of 16 ANSI colors.
var BrightGreen = /* #__PURE__ */ (function () {
    function BrightGreen() {

    };
    BrightGreen.value = new BrightGreen();
    return BrightGreen;
})();

// | The standard set of 16 ANSI colors.
var BrightYellow = /* #__PURE__ */ (function () {
    function BrightYellow() {

    };
    BrightYellow.value = new BrightYellow();
    return BrightYellow;
})();

// | The standard set of 16 ANSI colors.
var BrightBlue = /* #__PURE__ */ (function () {
    function BrightBlue() {

    };
    BrightBlue.value = new BrightBlue();
    return BrightBlue;
})();

// | The standard set of 16 ANSI colors.
var BrightMagenta = /* #__PURE__ */ (function () {
    function BrightMagenta() {

    };
    BrightMagenta.value = new BrightMagenta();
    return BrightMagenta;
})();

// | The standard set of 16 ANSI colors.
var BrightCyan = /* #__PURE__ */ (function () {
    function BrightCyan() {

    };
    BrightCyan.value = new BrightCyan();
    return BrightCyan;
})();

// | The standard set of 16 ANSI colors.
var BrightWhite = /* #__PURE__ */ (function () {
    function BrightWhite() {

    };
    BrightWhite.value = new BrightWhite();
    return BrightWhite;
})();

// | A graphics parameter, controls how text appears; for example, bold,
// | underlined, foreground color, background color.
var Reset = /* #__PURE__ */ (function () {
    function Reset() {

    };
    Reset.value = new Reset();
    return Reset;
})();

// | A graphics parameter, controls how text appears; for example, bold,
// | underlined, foreground color, background color.
var PMode = /* #__PURE__ */ (function () {
    function PMode(value0) {
        this.value0 = value0;
    };
    PMode.create = function (value0) {
        return new PMode(value0);
    };
    return PMode;
})();

// | A graphics parameter, controls how text appears; for example, bold,
// | underlined, foreground color, background color.
var PForeground = /* #__PURE__ */ (function () {
    function PForeground(value0) {
        this.value0 = value0;
    };
    PForeground.create = function (value0) {
        return new PForeground(value0);
    };
    return PForeground;
})();

// | A graphics parameter, controls how text appears; for example, bold,
// | underlined, foreground color, background color.
var PBackground = /* #__PURE__ */ (function () {
    function PBackground(value0) {
        this.value0 = value0;
    };
    PBackground.create = function (value0) {
        return new PBackground(value0);
    };
    return PBackground;
})();

// | An ANSI escape code. Not all sequences are implemented.
// | See: <https://en.wikipedia.org/wiki/ANSI_escape_code>.
var Up = /* #__PURE__ */ (function () {
    function Up(value0) {
        this.value0 = value0;
    };
    Up.create = function (value0) {
        return new Up(value0);
    };
    return Up;
})();

// | An ANSI escape code. Not all sequences are implemented.
// | See: <https://en.wikipedia.org/wiki/ANSI_escape_code>.
var Down = /* #__PURE__ */ (function () {
    function Down(value0) {
        this.value0 = value0;
    };
    Down.create = function (value0) {
        return new Down(value0);
    };
    return Down;
})();

// | An ANSI escape code. Not all sequences are implemented.
// | See: <https://en.wikipedia.org/wiki/ANSI_escape_code>.
var Forward = /* #__PURE__ */ (function () {
    function Forward(value0) {
        this.value0 = value0;
    };
    Forward.create = function (value0) {
        return new Forward(value0);
    };
    return Forward;
})();

// | An ANSI escape code. Not all sequences are implemented.
// | See: <https://en.wikipedia.org/wiki/ANSI_escape_code>.
var Back = /* #__PURE__ */ (function () {
    function Back(value0) {
        this.value0 = value0;
    };
    Back.create = function (value0) {
        return new Back(value0);
    };
    return Back;
})();

// | An ANSI escape code. Not all sequences are implemented.
// | See: <https://en.wikipedia.org/wiki/ANSI_escape_code>.
var NextLine = /* #__PURE__ */ (function () {
    function NextLine(value0) {
        this.value0 = value0;
    };
    NextLine.create = function (value0) {
        return new NextLine(value0);
    };
    return NextLine;
})();

// | An ANSI escape code. Not all sequences are implemented.
// | See: <https://en.wikipedia.org/wiki/ANSI_escape_code>.
var PreviousLine = /* #__PURE__ */ (function () {
    function PreviousLine(value0) {
        this.value0 = value0;
    };
    PreviousLine.create = function (value0) {
        return new PreviousLine(value0);
    };
    return PreviousLine;
})();

// | An ANSI escape code. Not all sequences are implemented.
// | See: <https://en.wikipedia.org/wiki/ANSI_escape_code>.
var HorizontalAbsolute = /* #__PURE__ */ (function () {
    function HorizontalAbsolute(value0) {
        this.value0 = value0;
    };
    HorizontalAbsolute.create = function (value0) {
        return new HorizontalAbsolute(value0);
    };
    return HorizontalAbsolute;
})();

// | An ANSI escape code. Not all sequences are implemented.
// | See: <https://en.wikipedia.org/wiki/ANSI_escape_code>.
var Position = /* #__PURE__ */ (function () {
    function Position(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    Position.create = function (value0) {
        return function (value1) {
            return new Position(value0, value1);
        };
    };
    return Position;
})();

// | An ANSI escape code. Not all sequences are implemented.
// | See: <https://en.wikipedia.org/wiki/ANSI_escape_code>.
var EraseData = /* #__PURE__ */ (function () {
    function EraseData(value0) {
        this.value0 = value0;
    };
    EraseData.create = function (value0) {
        return new EraseData(value0);
    };
    return EraseData;
})();

// | An ANSI escape code. Not all sequences are implemented.
// | See: <https://en.wikipedia.org/wiki/ANSI_escape_code>.
var EraseLine = /* #__PURE__ */ (function () {
    function EraseLine(value0) {
        this.value0 = value0;
    };
    EraseLine.create = function (value0) {
        return new EraseLine(value0);
    };
    return EraseLine;
})();

// | An ANSI escape code. Not all sequences are implemented.
// | See: <https://en.wikipedia.org/wiki/ANSI_escape_code>.
var ScrollUp = /* #__PURE__ */ (function () {
    function ScrollUp(value0) {
        this.value0 = value0;
    };
    ScrollUp.create = function (value0) {
        return new ScrollUp(value0);
    };
    return ScrollUp;
})();

// | An ANSI escape code. Not all sequences are implemented.
// | See: <https://en.wikipedia.org/wiki/ANSI_escape_code>.
var ScrollDown = /* #__PURE__ */ (function () {
    function ScrollDown(value0) {
        this.value0 = value0;
    };
    ScrollDown.create = function (value0) {
        return new ScrollDown(value0);
    };
    return ScrollDown;
})();

// | An ANSI escape code. Not all sequences are implemented.
// | See: <https://en.wikipedia.org/wiki/ANSI_escape_code>.
var Graphics = /* #__PURE__ */ (function () {
    function Graphics(value0) {
        this.value0 = value0;
    };
    Graphics.create = function (value0) {
        return new Graphics(value0);
    };
    return Graphics;
})();

// | An ANSI escape code. Not all sequences are implemented.
// | See: <https://en.wikipedia.org/wiki/ANSI_escape_code>.
var SavePosition = /* #__PURE__ */ (function () {
    function SavePosition() {

    };
    SavePosition.value = new SavePosition();
    return SavePosition;
})();

// | An ANSI escape code. Not all sequences are implemented.
// | See: <https://en.wikipedia.org/wiki/ANSI_escape_code>.
var RestorePosition = /* #__PURE__ */ (function () {
    function RestorePosition() {

    };
    RestorePosition.value = new RestorePosition();
    return RestorePosition;
})();

// | An ANSI escape code. Not all sequences are implemented.
// | See: <https://en.wikipedia.org/wiki/ANSI_escape_code>.
var QueryPosition = /* #__PURE__ */ (function () {
    function QueryPosition() {

    };
    QueryPosition.value = new QueryPosition();
    return QueryPosition;
})();

// | An ANSI escape code. Not all sequences are implemented.
// | See: <https://en.wikipedia.org/wiki/ANSI_escape_code>.
var HideCursor = /* #__PURE__ */ (function () {
    function HideCursor() {

    };
    HideCursor.value = new HideCursor();
    return HideCursor;
})();

// | An ANSI escape code. Not all sequences are implemented.
// | See: <https://en.wikipedia.org/wiki/ANSI_escape_code>.
var ShowCursor = /* #__PURE__ */ (function () {
    function ShowCursor() {

    };
    ShowCursor.value = new ShowCursor();
    return ShowCursor;
})();

// | The prefix for all escape codes.
var prefix = "\x1b[";
var genericRenderingMode = {
    to: function (x) {
        if (x instanceof Data_Generic_Rep.Inl) {
            return Bold.value;
        };
        if (x instanceof Data_Generic_Rep.Inr && x.value0 instanceof Data_Generic_Rep.Inl) {
            return Dim.value;
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0 instanceof Data_Generic_Rep.Inl)) {
            return Italic.value;
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0 instanceof Data_Generic_Rep.Inl))) {
            return Underline.value;
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inl)))) {
            return Inverse.value;
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr)))) {
            return Strikethrough.value;
        };
        throw new Error("Failed pattern match at Ansi.Codes (line 130, column 1 - line 130, column 64): " + [ x.constructor.name ]);
    },
    from: function (x) {
        if (x instanceof Bold) {
            return new Data_Generic_Rep.Inl(Data_Generic_Rep.NoArguments.value);
        };
        if (x instanceof Dim) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(Data_Generic_Rep.NoArguments.value));
        };
        if (x instanceof Italic) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(Data_Generic_Rep.NoArguments.value)));
        };
        if (x instanceof Underline) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(Data_Generic_Rep.NoArguments.value))));
        };
        if (x instanceof Inverse) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(Data_Generic_Rep.NoArguments.value)))));
        };
        if (x instanceof Strikethrough) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(Data_Generic_Rep.NoArguments.value)))));
        };
        throw new Error("Failed pattern match at Ansi.Codes (line 130, column 1 - line 130, column 64): " + [ x.constructor.name ]);
    }
};
var showRenderingMode = {
    show: /* #__PURE__ */ Data_Show_Generic.genericShow(genericRenderingMode)(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "Bold";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "Dim";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "Italic";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "Underline";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "Inverse";
        }
    }))(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "Strikethrough";
        }
    })))))))
};
var genericGraphicsParam = {
    to: function (x) {
        if (x instanceof Data_Generic_Rep.Inl) {
            return Reset.value;
        };
        if (x instanceof Data_Generic_Rep.Inr && x.value0 instanceof Data_Generic_Rep.Inl) {
            return new PMode(x.value0.value0);
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0 instanceof Data_Generic_Rep.Inl)) {
            return new PForeground(x.value0.value0.value0);
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0 instanceof Data_Generic_Rep.Inr)) {
            return new PBackground(x.value0.value0.value0);
        };
        throw new Error("Failed pattern match at Ansi.Codes (line 108, column 1 - line 108, column 64): " + [ x.constructor.name ]);
    },
    from: function (x) {
        if (x instanceof Reset) {
            return new Data_Generic_Rep.Inl(Data_Generic_Rep.NoArguments.value);
        };
        if (x instanceof PMode) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(x.value0));
        };
        if (x instanceof PForeground) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(x.value0)));
        };
        if (x instanceof PBackground) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(x.value0)));
        };
        throw new Error("Failed pattern match at Ansi.Codes (line 108, column 1 - line 108, column 64): " + [ x.constructor.name ]);
    }
};
var genericEscapeCode = {
    to: function (x) {
        if (x instanceof Data_Generic_Rep.Inl) {
            return new Up(x.value0);
        };
        if (x instanceof Data_Generic_Rep.Inr && x.value0 instanceof Data_Generic_Rep.Inl) {
            return new Down(x.value0.value0);
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0 instanceof Data_Generic_Rep.Inl)) {
            return new Forward(x.value0.value0.value0);
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0 instanceof Data_Generic_Rep.Inl))) {
            return new Back(x.value0.value0.value0.value0);
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inl)))) {
            return new NextLine(x.value0.value0.value0.value0.value0);
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inl))))) {
            return new PreviousLine(x.value0.value0.value0.value0.value0.value0);
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inl)))))) {
            return new HorizontalAbsolute(x.value0.value0.value0.value0.value0.value0.value0);
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inl))))))) {
            return new Position(x.value0.value0.value0.value0.value0.value0.value0.value0.value0, x.value0.value0.value0.value0.value0.value0.value0.value0.value1);
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inl)))))))) {
            return new EraseData(x.value0.value0.value0.value0.value0.value0.value0.value0.value0);
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inl))))))))) {
            return new EraseLine(x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0);
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inl)))))))))) {
            return new ScrollUp(x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0);
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inl))))))))))) {
            return new ScrollDown(x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0);
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inl)))))))))))) {
            return new Graphics(x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0);
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inl))))))))))))) {
            return SavePosition.value;
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inl)))))))))))))) {
            return RestorePosition.value;
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inl))))))))))))))) {
            return QueryPosition.value;
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inl)))))))))))))))) {
            return HideCursor.value;
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr)))))))))))))))) {
            return ShowCursor.value;
        };
        throw new Error("Failed pattern match at Ansi.Codes (line 42, column 1 - line 42, column 58): " + [ x.constructor.name ]);
    },
    from: function (x) {
        if (x instanceof Up) {
            return new Data_Generic_Rep.Inl(x.value0);
        };
        if (x instanceof Down) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(x.value0));
        };
        if (x instanceof Forward) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(x.value0)));
        };
        if (x instanceof Back) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(x.value0))));
        };
        if (x instanceof NextLine) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(x.value0)))));
        };
        if (x instanceof PreviousLine) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(x.value0))))));
        };
        if (x instanceof HorizontalAbsolute) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(x.value0)))))));
        };
        if (x instanceof Position) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(new Data_Generic_Rep.Product(x.value0, x.value1)))))))));
        };
        if (x instanceof EraseData) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(x.value0)))))))));
        };
        if (x instanceof EraseLine) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(x.value0))))))))));
        };
        if (x instanceof ScrollUp) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(x.value0)))))))))));
        };
        if (x instanceof ScrollDown) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(x.value0))))))))))));
        };
        if (x instanceof Graphics) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(x.value0)))))))))))));
        };
        if (x instanceof SavePosition) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(Data_Generic_Rep.NoArguments.value))))))))))))));
        };
        if (x instanceof RestorePosition) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(Data_Generic_Rep.NoArguments.value)))))))))))))));
        };
        if (x instanceof QueryPosition) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(Data_Generic_Rep.NoArguments.value))))))))))))))));
        };
        if (x instanceof HideCursor) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(Data_Generic_Rep.NoArguments.value)))))))))))))))));
        };
        if (x instanceof ShowCursor) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(Data_Generic_Rep.NoArguments.value)))))))))))))))));
        };
        throw new Error("Failed pattern match at Ansi.Codes (line 42, column 1 - line 42, column 58): " + [ x.constructor.name ]);
    }
};
var genericEraseParam = {
    to: function (x) {
        if (x instanceof Data_Generic_Rep.Inl) {
            return ToEnd.value;
        };
        if (x instanceof Data_Generic_Rep.Inr && x.value0 instanceof Data_Generic_Rep.Inl) {
            return FromBeginning.value;
        };
        if (x instanceof Data_Generic_Rep.Inr && x.value0 instanceof Data_Generic_Rep.Inr) {
            return Entire.value;
        };
        throw new Error("Failed pattern match at Ansi.Codes (line 87, column 1 - line 87, column 58): " + [ x.constructor.name ]);
    },
    from: function (x) {
        if (x instanceof ToEnd) {
            return new Data_Generic_Rep.Inl(Data_Generic_Rep.NoArguments.value);
        };
        if (x instanceof FromBeginning) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(Data_Generic_Rep.NoArguments.value));
        };
        if (x instanceof Entire) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(Data_Generic_Rep.NoArguments.value));
        };
        throw new Error("Failed pattern match at Ansi.Codes (line 87, column 1 - line 87, column 58): " + [ x.constructor.name ]);
    }
};
var showEraseParam = {
    show: /* #__PURE__ */ Data_Show_Generic.genericShow(genericEraseParam)(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "ToEnd";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "FromBeginning";
        }
    }))(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "Entire";
        }
    }))))
};
var genericShowConstructor2 = /* #__PURE__ */ Data_Show_Generic.genericShowConstructor(/* #__PURE__ */ Data_Show_Generic.genericShowArgsArgument(showEraseParam));
var genericColor = {
    to: function (x) {
        if (x instanceof Data_Generic_Rep.Inl) {
            return Black.value;
        };
        if (x instanceof Data_Generic_Rep.Inr && x.value0 instanceof Data_Generic_Rep.Inl) {
            return Red.value;
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0 instanceof Data_Generic_Rep.Inl)) {
            return Green.value;
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0 instanceof Data_Generic_Rep.Inl))) {
            return Yellow.value;
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inl)))) {
            return Blue.value;
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inl))))) {
            return Magenta.value;
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inl)))))) {
            return Cyan.value;
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inl))))))) {
            return White.value;
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inl)))))))) {
            return BrightBlack.value;
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inl))))))))) {
            return BrightRed.value;
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inl)))))))))) {
            return BrightGreen.value;
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inl))))))))))) {
            return BrightYellow.value;
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inl)))))))))))) {
            return BrightBlue.value;
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inl))))))))))))) {
            return BrightMagenta.value;
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inl)))))))))))))) {
            return BrightCyan.value;
        };
        if (x instanceof Data_Generic_Rep.Inr && (x.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && (x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr && x.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0.value0 instanceof Data_Generic_Rep.Inr)))))))))))))) {
            return BrightWhite.value;
        };
        throw new Error("Failed pattern match at Ansi.Codes (line 165, column 1 - line 165, column 48): " + [ x.constructor.name ]);
    },
    from: function (x) {
        if (x instanceof Black) {
            return new Data_Generic_Rep.Inl(Data_Generic_Rep.NoArguments.value);
        };
        if (x instanceof Red) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(Data_Generic_Rep.NoArguments.value));
        };
        if (x instanceof Green) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(Data_Generic_Rep.NoArguments.value)));
        };
        if (x instanceof Yellow) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(Data_Generic_Rep.NoArguments.value))));
        };
        if (x instanceof Blue) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(Data_Generic_Rep.NoArguments.value)))));
        };
        if (x instanceof Magenta) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(Data_Generic_Rep.NoArguments.value))))));
        };
        if (x instanceof Cyan) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(Data_Generic_Rep.NoArguments.value)))))));
        };
        if (x instanceof White) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(Data_Generic_Rep.NoArguments.value))))))));
        };
        if (x instanceof BrightBlack) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(Data_Generic_Rep.NoArguments.value)))))))));
        };
        if (x instanceof BrightRed) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(Data_Generic_Rep.NoArguments.value))))))))));
        };
        if (x instanceof BrightGreen) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(Data_Generic_Rep.NoArguments.value)))))))))));
        };
        if (x instanceof BrightYellow) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(Data_Generic_Rep.NoArguments.value))))))))))));
        };
        if (x instanceof BrightBlue) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(Data_Generic_Rep.NoArguments.value)))))))))))));
        };
        if (x instanceof BrightMagenta) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(Data_Generic_Rep.NoArguments.value))))))))))))));
        };
        if (x instanceof BrightCyan) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inl(Data_Generic_Rep.NoArguments.value)))))))))))))));
        };
        if (x instanceof BrightWhite) {
            return new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(new Data_Generic_Rep.Inr(Data_Generic_Rep.NoArguments.value)))))))))))))));
        };
        throw new Error("Failed pattern match at Ansi.Codes (line 165, column 1 - line 165, column 48): " + [ x.constructor.name ]);
    }
};
var showColor = {
    show: /* #__PURE__ */ Data_Show_Generic.genericShow(genericColor)(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "Black";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "Red";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "Green";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "Yellow";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "Blue";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "Magenta";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "Cyan";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "White";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "BrightBlack";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "BrightRed";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "BrightGreen";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "BrightYellow";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "BrightBlue";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "BrightMagenta";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "BrightCyan";
        }
    }))(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "BrightWhite";
        }
    })))))))))))))))))
};
var genericShowConstructor3 = /* #__PURE__ */ Data_Show_Generic.genericShowConstructor(/* #__PURE__ */ Data_Show_Generic.genericShowArgsArgument(showColor));
var showGraphicsParam = {
    show: /* #__PURE__ */ Data_Show_Generic.genericShow(genericGraphicsParam)(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "Reset";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ Data_Show_Generic.genericShowConstructor(/* #__PURE__ */ Data_Show_Generic.genericShowArgsArgument(showRenderingMode))({
        reflectSymbol: function () {
            return "PMode";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor3({
        reflectSymbol: function () {
            return "PForeground";
        }
    }))(/* #__PURE__ */ genericShowConstructor3({
        reflectSymbol: function () {
            return "PBackground";
        }
    })))))
};
var showEscapeCode = {
    show: /* #__PURE__ */ Data_Show_Generic.genericShow(genericEscapeCode)(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor1({
        reflectSymbol: function () {
            return "Up";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor1({
        reflectSymbol: function () {
            return "Down";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor1({
        reflectSymbol: function () {
            return "Forward";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor1({
        reflectSymbol: function () {
            return "Back";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor1({
        reflectSymbol: function () {
            return "NextLine";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor1({
        reflectSymbol: function () {
            return "PreviousLine";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor1({
        reflectSymbol: function () {
            return "HorizontalAbsolute";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ Data_Show_Generic.genericShowConstructor(/* #__PURE__ */ Data_Show_Generic.genericShowArgsProduct(genericShowArgsArgument)(genericShowArgsArgument))({
        reflectSymbol: function () {
            return "Position";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor2({
        reflectSymbol: function () {
            return "EraseData";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor2({
        reflectSymbol: function () {
            return "EraseLine";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor1({
        reflectSymbol: function () {
            return "ScrollUp";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor1({
        reflectSymbol: function () {
            return "ScrollDown";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ Data_Show_Generic.genericShowConstructor(/* #__PURE__ */ Data_Show_Generic.genericShowArgsArgument(/* #__PURE__ */ Data_List_Types.showNonEmptyList(showGraphicsParam)))({
        reflectSymbol: function () {
            return "Graphics";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "SavePosition";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "RestorePosition";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "QueryPosition";
        }
    }))(/* #__PURE__ */ Data_Show_Generic.genericShowSum(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "HideCursor";
        }
    }))(/* #__PURE__ */ genericShowConstructor({
        reflectSymbol: function () {
            return "ShowCursor";
        }
    })))))))))))))))))))
};
var eraseParamToString = function (ep) {
    if (ep instanceof ToEnd) {
        return "0";
    };
    if (ep instanceof FromBeginning) {
        return "1";
    };
    if (ep instanceof Entire) {
        return "2";
    };
    throw new Error("Failed pattern match at Ansi.Codes (line 95, column 3 - line 98, column 25): " + [ ep.constructor.name ]);
};
var eqRenderingMode = {
    eq: function (x) {
        return function (y) {
            if (x instanceof Bold && y instanceof Bold) {
                return true;
            };
            if (x instanceof Dim && y instanceof Dim) {
                return true;
            };
            if (x instanceof Italic && y instanceof Italic) {
                return true;
            };
            if (x instanceof Underline && y instanceof Underline) {
                return true;
            };
            if (x instanceof Inverse && y instanceof Inverse) {
                return true;
            };
            if (x instanceof Strikethrough && y instanceof Strikethrough) {
                return true;
            };
            return false;
        };
    }
};
var ordRenderingMode = {
    compare: function (x) {
        return function (y) {
            if (x instanceof Bold && y instanceof Bold) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof Bold) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof Bold) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof Dim && y instanceof Dim) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof Dim) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof Dim) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof Italic && y instanceof Italic) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof Italic) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof Italic) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof Underline && y instanceof Underline) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof Underline) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof Underline) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof Inverse && y instanceof Inverse) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof Inverse) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof Inverse) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof Strikethrough && y instanceof Strikethrough) {
                return Data_Ordering.EQ.value;
            };
            throw new Error("Failed pattern match at Ansi.Codes (line 0, column 0 - line 0, column 0): " + [ x.constructor.name, y.constructor.name ]);
        };
    },
    Eq0: function () {
        return eqRenderingMode;
    }
};
var eqEraseParam = {
    eq: function (x) {
        return function (y) {
            if (x instanceof ToEnd && y instanceof ToEnd) {
                return true;
            };
            if (x instanceof FromBeginning && y instanceof FromBeginning) {
                return true;
            };
            if (x instanceof Entire && y instanceof Entire) {
                return true;
            };
            return false;
        };
    }
};
var ordEraseParam = {
    compare: function (x) {
        return function (y) {
            if (x instanceof ToEnd && y instanceof ToEnd) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof ToEnd) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof ToEnd) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof FromBeginning && y instanceof FromBeginning) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof FromBeginning) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof FromBeginning) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof Entire && y instanceof Entire) {
                return Data_Ordering.EQ.value;
            };
            throw new Error("Failed pattern match at Ansi.Codes (line 0, column 0 - line 0, column 0): " + [ x.constructor.name, y.constructor.name ]);
        };
    },
    Eq0: function () {
        return eqEraseParam;
    }
};
var eqColor = {
    eq: function (x) {
        return function (y) {
            if (x instanceof Black && y instanceof Black) {
                return true;
            };
            if (x instanceof Red && y instanceof Red) {
                return true;
            };
            if (x instanceof Green && y instanceof Green) {
                return true;
            };
            if (x instanceof Yellow && y instanceof Yellow) {
                return true;
            };
            if (x instanceof Blue && y instanceof Blue) {
                return true;
            };
            if (x instanceof Magenta && y instanceof Magenta) {
                return true;
            };
            if (x instanceof Cyan && y instanceof Cyan) {
                return true;
            };
            if (x instanceof White && y instanceof White) {
                return true;
            };
            if (x instanceof BrightBlack && y instanceof BrightBlack) {
                return true;
            };
            if (x instanceof BrightRed && y instanceof BrightRed) {
                return true;
            };
            if (x instanceof BrightGreen && y instanceof BrightGreen) {
                return true;
            };
            if (x instanceof BrightYellow && y instanceof BrightYellow) {
                return true;
            };
            if (x instanceof BrightBlue && y instanceof BrightBlue) {
                return true;
            };
            if (x instanceof BrightMagenta && y instanceof BrightMagenta) {
                return true;
            };
            if (x instanceof BrightCyan && y instanceof BrightCyan) {
                return true;
            };
            if (x instanceof BrightWhite && y instanceof BrightWhite) {
                return true;
            };
            return false;
        };
    }
};
var eqGraphicsParam = {
    eq: function (x) {
        return function (y) {
            if (x instanceof Reset && y instanceof Reset) {
                return true;
            };
            if (x instanceof PMode && y instanceof PMode) {
                return Data_Eq.eq(eqRenderingMode)(x.value0)(y.value0);
            };
            if (x instanceof PForeground && y instanceof PForeground) {
                return Data_Eq.eq(eqColor)(x.value0)(y.value0);
            };
            if (x instanceof PBackground && y instanceof PBackground) {
                return Data_Eq.eq(eqColor)(x.value0)(y.value0);
            };
            return false;
        };
    }
};
var eqNonEmptyList = /* #__PURE__ */ Data_List_Types.eqNonEmptyList(eqGraphicsParam);
var eqEscapeCode = {
    eq: function (x) {
        return function (y) {
            if (x instanceof Up && y instanceof Up) {
                return x.value0 === y.value0;
            };
            if (x instanceof Down && y instanceof Down) {
                return x.value0 === y.value0;
            };
            if (x instanceof Forward && y instanceof Forward) {
                return x.value0 === y.value0;
            };
            if (x instanceof Back && y instanceof Back) {
                return x.value0 === y.value0;
            };
            if (x instanceof NextLine && y instanceof NextLine) {
                return x.value0 === y.value0;
            };
            if (x instanceof PreviousLine && y instanceof PreviousLine) {
                return x.value0 === y.value0;
            };
            if (x instanceof HorizontalAbsolute && y instanceof HorizontalAbsolute) {
                return x.value0 === y.value0;
            };
            if (x instanceof Position && y instanceof Position) {
                return x.value0 === y.value0 && x.value1 === y.value1;
            };
            if (x instanceof EraseData && y instanceof EraseData) {
                return Data_Eq.eq(eqEraseParam)(x.value0)(y.value0);
            };
            if (x instanceof EraseLine && y instanceof EraseLine) {
                return Data_Eq.eq(eqEraseParam)(x.value0)(y.value0);
            };
            if (x instanceof ScrollUp && y instanceof ScrollUp) {
                return x.value0 === y.value0;
            };
            if (x instanceof ScrollDown && y instanceof ScrollDown) {
                return x.value0 === y.value0;
            };
            if (x instanceof Graphics && y instanceof Graphics) {
                return Data_Eq.eq(eqNonEmptyList)(x.value0)(y.value0);
            };
            if (x instanceof SavePosition && y instanceof SavePosition) {
                return true;
            };
            if (x instanceof RestorePosition && y instanceof RestorePosition) {
                return true;
            };
            if (x instanceof QueryPosition && y instanceof QueryPosition) {
                return true;
            };
            if (x instanceof HideCursor && y instanceof HideCursor) {
                return true;
            };
            if (x instanceof ShowCursor && y instanceof ShowCursor) {
                return true;
            };
            return false;
        };
    }
};
var ordColor = {
    compare: function (x) {
        return function (y) {
            if (x instanceof Black && y instanceof Black) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof Black) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof Black) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof Red && y instanceof Red) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof Red) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof Red) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof Green && y instanceof Green) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof Green) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof Green) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof Yellow && y instanceof Yellow) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof Yellow) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof Yellow) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof Blue && y instanceof Blue) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof Blue) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof Blue) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof Magenta && y instanceof Magenta) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof Magenta) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof Magenta) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof Cyan && y instanceof Cyan) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof Cyan) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof Cyan) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof White && y instanceof White) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof White) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof White) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof BrightBlack && y instanceof BrightBlack) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof BrightBlack) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof BrightBlack) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof BrightRed && y instanceof BrightRed) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof BrightRed) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof BrightRed) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof BrightGreen && y instanceof BrightGreen) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof BrightGreen) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof BrightGreen) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof BrightYellow && y instanceof BrightYellow) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof BrightYellow) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof BrightYellow) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof BrightBlue && y instanceof BrightBlue) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof BrightBlue) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof BrightBlue) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof BrightMagenta && y instanceof BrightMagenta) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof BrightMagenta) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof BrightMagenta) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof BrightCyan && y instanceof BrightCyan) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof BrightCyan) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof BrightCyan) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof BrightWhite && y instanceof BrightWhite) {
                return Data_Ordering.EQ.value;
            };
            throw new Error("Failed pattern match at Ansi.Codes (line 0, column 0 - line 0, column 0): " + [ x.constructor.name, y.constructor.name ]);
        };
    },
    Eq0: function () {
        return eqColor;
    }
};
var ordGraphicsParam = {
    compare: function (x) {
        return function (y) {
            if (x instanceof Reset && y instanceof Reset) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof Reset) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof Reset) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof PMode && y instanceof PMode) {
                return Data_Ord.compare(ordRenderingMode)(x.value0)(y.value0);
            };
            if (x instanceof PMode) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof PMode) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof PForeground && y instanceof PForeground) {
                return Data_Ord.compare(ordColor)(x.value0)(y.value0);
            };
            if (x instanceof PForeground) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof PForeground) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof PBackground && y instanceof PBackground) {
                return Data_Ord.compare(ordColor)(x.value0)(y.value0);
            };
            throw new Error("Failed pattern match at Ansi.Codes (line 0, column 0 - line 0, column 0): " + [ x.constructor.name, y.constructor.name ]);
        };
    },
    Eq0: function () {
        return eqGraphicsParam;
    }
};
var ordNonEmptyList = /* #__PURE__ */ Data_List_Types.ordNonEmptyList(ordGraphicsParam);
var ordEscapeCode = {
    compare: function (x) {
        return function (y) {
            if (x instanceof Up && y instanceof Up) {
                return Data_Ord.compare(Data_Ord.ordInt)(x.value0)(y.value0);
            };
            if (x instanceof Up) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof Up) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof Down && y instanceof Down) {
                return Data_Ord.compare(Data_Ord.ordInt)(x.value0)(y.value0);
            };
            if (x instanceof Down) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof Down) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof Forward && y instanceof Forward) {
                return Data_Ord.compare(Data_Ord.ordInt)(x.value0)(y.value0);
            };
            if (x instanceof Forward) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof Forward) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof Back && y instanceof Back) {
                return Data_Ord.compare(Data_Ord.ordInt)(x.value0)(y.value0);
            };
            if (x instanceof Back) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof Back) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof NextLine && y instanceof NextLine) {
                return Data_Ord.compare(Data_Ord.ordInt)(x.value0)(y.value0);
            };
            if (x instanceof NextLine) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof NextLine) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof PreviousLine && y instanceof PreviousLine) {
                return Data_Ord.compare(Data_Ord.ordInt)(x.value0)(y.value0);
            };
            if (x instanceof PreviousLine) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof PreviousLine) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof HorizontalAbsolute && y instanceof HorizontalAbsolute) {
                return Data_Ord.compare(Data_Ord.ordInt)(x.value0)(y.value0);
            };
            if (x instanceof HorizontalAbsolute) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof HorizontalAbsolute) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof Position && y instanceof Position) {
                var v = Data_Ord.compare(Data_Ord.ordInt)(x.value0)(y.value0);
                if (v instanceof Data_Ordering.LT) {
                    return Data_Ordering.LT.value;
                };
                if (v instanceof Data_Ordering.GT) {
                    return Data_Ordering.GT.value;
                };
                return Data_Ord.compare(Data_Ord.ordInt)(x.value1)(y.value1);
            };
            if (x instanceof Position) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof Position) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof EraseData && y instanceof EraseData) {
                return Data_Ord.compare(ordEraseParam)(x.value0)(y.value0);
            };
            if (x instanceof EraseData) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof EraseData) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof EraseLine && y instanceof EraseLine) {
                return Data_Ord.compare(ordEraseParam)(x.value0)(y.value0);
            };
            if (x instanceof EraseLine) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof EraseLine) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof ScrollUp && y instanceof ScrollUp) {
                return Data_Ord.compare(Data_Ord.ordInt)(x.value0)(y.value0);
            };
            if (x instanceof ScrollUp) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof ScrollUp) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof ScrollDown && y instanceof ScrollDown) {
                return Data_Ord.compare(Data_Ord.ordInt)(x.value0)(y.value0);
            };
            if (x instanceof ScrollDown) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof ScrollDown) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof Graphics && y instanceof Graphics) {
                return Data_Ord.compare(ordNonEmptyList)(x.value0)(y.value0);
            };
            if (x instanceof Graphics) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof Graphics) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof SavePosition && y instanceof SavePosition) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof SavePosition) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof SavePosition) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof RestorePosition && y instanceof RestorePosition) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof RestorePosition) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof RestorePosition) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof QueryPosition && y instanceof QueryPosition) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof QueryPosition) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof QueryPosition) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof HideCursor && y instanceof HideCursor) {
                return Data_Ordering.EQ.value;
            };
            if (x instanceof HideCursor) {
                return Data_Ordering.LT.value;
            };
            if (y instanceof HideCursor) {
                return Data_Ordering.GT.value;
            };
            if (x instanceof ShowCursor && y instanceof ShowCursor) {
                return Data_Ordering.EQ.value;
            };
            throw new Error("Failed pattern match at Ansi.Codes (line 0, column 0 - line 0, column 0): " + [ x.constructor.name, y.constructor.name ]);
        };
    },
    Eq0: function () {
        return eqEscapeCode;
    }
};

// | The suffix for escape codes; note that this is only required for colors.
var colorSuffix = "m";
var colorCode = function (c) {
    if (c instanceof Black) {
        return 30;
    };
    if (c instanceof Red) {
        return 31;
    };
    if (c instanceof Green) {
        return 32;
    };
    if (c instanceof Yellow) {
        return 33;
    };
    if (c instanceof Blue) {
        return 34;
    };
    if (c instanceof Magenta) {
        return 35;
    };
    if (c instanceof Cyan) {
        return 36;
    };
    if (c instanceof White) {
        return 37;
    };
    if (c instanceof BrightBlack) {
        return 90;
    };
    if (c instanceof BrightRed) {
        return 91;
    };
    if (c instanceof BrightGreen) {
        return 92;
    };
    if (c instanceof BrightYellow) {
        return 93;
    };
    if (c instanceof BrightBlue) {
        return 94;
    };
    if (c instanceof BrightMagenta) {
        return 95;
    };
    if (c instanceof BrightCyan) {
        return 96;
    };
    if (c instanceof BrightWhite) {
        return 97;
    };
    throw new Error("Failed pattern match at Ansi.Codes (line 173, column 3 - line 189, column 22): " + [ c.constructor.name ]);
};
var codeForRenderingMode = function (m) {
    if (m instanceof Bold) {
        return 1;
    };
    if (m instanceof Dim) {
        return 2;
    };
    if (m instanceof Italic) {
        return 3;
    };
    if (m instanceof Underline) {
        return 4;
    };
    if (m instanceof Inverse) {
        return 7;
    };
    if (m instanceof Strikethrough) {
        return 9;
    };
    throw new Error("Failed pattern match at Ansi.Codes (line 138, column 3 - line 144, column 23): " + [ m.constructor.name ]);
};
var graphicsParamToString = function (gp) {
    if (gp instanceof Reset) {
        return "0";
    };
    if (gp instanceof PMode) {
        return Data_Show.show(Data_Show.showInt)(codeForRenderingMode(gp.value0));
    };
    if (gp instanceof PForeground) {
        return Data_Show.show(Data_Show.showInt)(colorCode(gp.value0));
    };
    if (gp instanceof PBackground) {
        return Data_Show.show(Data_Show.showInt)(colorCode(gp.value0) + 10 | 0);
    };
    throw new Error("Failed pattern match at Ansi.Codes (line 116, column 3 - line 120, column 45): " + [ gp.constructor.name ]);
};

// | Convert an escape code to the form recognised by terminals.
var escapeCodeToString = /* #__PURE__ */ (function () {
    var go = function (c) {
        if (c instanceof Up) {
            return Data_Show.show(Data_Show.showInt)(c.value0) + "A";
        };
        if (c instanceof Down) {
            return Data_Show.show(Data_Show.showInt)(c.value0) + "B";
        };
        if (c instanceof Forward) {
            return Data_Show.show(Data_Show.showInt)(c.value0) + "C";
        };
        if (c instanceof Back) {
            return Data_Show.show(Data_Show.showInt)(c.value0) + "D";
        };
        if (c instanceof NextLine) {
            return Data_Show.show(Data_Show.showInt)(c.value0) + "E";
        };
        if (c instanceof PreviousLine) {
            return Data_Show.show(Data_Show.showInt)(c.value0) + "F";
        };
        if (c instanceof HorizontalAbsolute) {
            return Data_Show.show(Data_Show.showInt)(c.value0) + "G";
        };
        if (c instanceof Position) {
            return Data_Show.show(Data_Show.showInt)(c.value0) + (";" + (Data_Show.show(Data_Show.showInt)(c.value1) + "H"));
        };
        if (c instanceof EraseData) {
            return eraseParamToString(c.value0) + "J";
        };
        if (c instanceof EraseLine) {
            return eraseParamToString(c.value0) + "K";
        };
        if (c instanceof ScrollUp) {
            return Data_Show.show(Data_Show.showInt)(c.value0) + "S";
        };
        if (c instanceof ScrollDown) {
            return Data_Show.show(Data_Show.showInt)(c.value0) + "T";
        };
        if (c instanceof Graphics) {
            return Data_Foldable.intercalate(Data_List_Types.foldableNonEmptyList)(Data_Monoid.monoidString)(";")(Data_Functor.map(Data_List_Types.functorNonEmptyList)(graphicsParamToString)(c.value0)) + colorSuffix;
        };
        if (c instanceof SavePosition) {
            return "s";
        };
        if (c instanceof RestorePosition) {
            return "u";
        };
        if (c instanceof QueryPosition) {
            return "6n";
        };
        if (c instanceof HideCursor) {
            return "?25l";
        };
        if (c instanceof ShowCursor) {
            return "?25h";
        };
        throw new Error("Failed pattern match at Ansi.Codes (line 53, column 5 - line 71, column 37): " + [ c.constructor.name ]);
    };
    return function ($828) {
        return (function (v) {
            return prefix + v;
        })(go($828));
    };
})();
export {
    prefix,
    colorSuffix,
    Up,
    Down,
    Forward,
    Back,
    NextLine,
    PreviousLine,
    HorizontalAbsolute,
    Position,
    EraseData,
    EraseLine,
    ScrollUp,
    ScrollDown,
    Graphics,
    SavePosition,
    RestorePosition,
    QueryPosition,
    HideCursor,
    ShowCursor,
    escapeCodeToString,
    ToEnd,
    FromBeginning,
    Entire,
    eraseParamToString,
    Reset,
    PMode,
    PForeground,
    PBackground,
    graphicsParamToString,
    Bold,
    Dim,
    Italic,
    Underline,
    Inverse,
    Strikethrough,
    codeForRenderingMode,
    Black,
    Red,
    Green,
    Yellow,
    Blue,
    Magenta,
    Cyan,
    White,
    BrightBlack,
    BrightRed,
    BrightGreen,
    BrightYellow,
    BrightBlue,
    BrightMagenta,
    BrightCyan,
    BrightWhite,
    colorCode,
    genericEscapeCode,
    eqEscapeCode,
    ordEscapeCode,
    showEscapeCode,
    genericEraseParam,
    eqEraseParam,
    ordEraseParam,
    showEraseParam,
    genericGraphicsParam,
    eqGraphicsParam,
    ordGraphicsParam,
    showGraphicsParam,
    genericRenderingMode,
    eqRenderingMode,
    ordRenderingMode,
    showRenderingMode,
    genericColor,
    eqColor,
    ordColor,
    showColor
};
