import * as Data$dEq from "../Data.Eq/index.js";
import * as Data$dOrd from "../Data.Ord/index.js";
const $AppSpine = (tag, _1, _2) => ({tag, _1, _2});
const $Binder = (tag, _1, _2, _3) => ({tag, _1, _2, _3});
const $ClassFundep = (tag, _1, _2, _3) => ({tag, _1, _2, _3});
const $Comment = (tag, _1, _2) => ({tag, _1, _2});
const $DataMembers = (tag, _1) => ({tag, _1});
const $Declaration = (tag, _1, _2, _3, _4) => ({tag, _1, _2, _3, _4});
const $DoStatement = (tag, _1, _2, _3) => ({tag, _1, _2, _3});
const $Export = (tag, _1, _2) => ({tag, _1, _2});
const $Expr = (tag, _1, _2, _3) => ({tag, _1, _2, _3});
const $Fixity = tag => tag;
const $FixityOp = (tag, _1, _2, _3, _4) => ({tag, _1, _2, _3, _4});
const $Foreign = (tag, _1, _2) => ({tag, _1, _2});
const $Guarded = (tag, _1, _2) => ({tag, _1, _2});
const $Import = (tag, _1, _2) => ({tag, _1, _2});
const $InstanceBinding = (tag, _1) => ({tag, _1});
const $IntValue = (tag, _1) => ({tag, _1});
const $LetBinding = (tag, _1, _2, _3) => ({tag, _1, _2, _3});
const $LineFeed = tag => tag;
const $OneOrDelimited = (tag, _1) => ({tag, _1});
const $RecordLabeled = (tag, _1, _2, _3) => ({tag, _1, _2, _3});
const $RecordUpdate = (tag, _1, _2, _3) => ({tag, _1, _2, _3});
const $Role = tag => tag;
const $SourceStyle = tag => tag;
const $Token = (tag, _1, _2) => ({tag, _1, _2});
const $Type = (tag, _1, _2, _3, _4) => ({tag, _1, _2, _3, _4});
const $TypeVarBinding = (tag, _1) => ({tag, _1});
const ASCII = /* #__PURE__ */ $SourceStyle("ASCII");
const Unicode = /* #__PURE__ */ $SourceStyle("Unicode");
const Nominal = /* #__PURE__ */ $Role("Nominal");
const Representational = /* #__PURE__ */ $Role("Representational");
const Phantom = /* #__PURE__ */ $Role("Phantom");
const Proper = x => x;
const Operator = x => x;
const ModuleName = x => x;
const LF = /* #__PURE__ */ $LineFeed("LF");
const CRLF = /* #__PURE__ */ $LineFeed("CRLF");
const Label = x => x;
const SmallInt = value0 => $IntValue("SmallInt", value0);
const $$BigInt = value0 => $IntValue("BigInt", value0);
const BigHex = value0 => $IntValue("BigHex", value0);
const TokLeftParen = /* #__PURE__ */ $Token("TokLeftParen");
const TokRightParen = /* #__PURE__ */ $Token("TokRightParen");
const TokLeftBrace = /* #__PURE__ */ $Token("TokLeftBrace");
const TokRightBrace = /* #__PURE__ */ $Token("TokRightBrace");
const TokLeftSquare = /* #__PURE__ */ $Token("TokLeftSquare");
const TokRightSquare = /* #__PURE__ */ $Token("TokRightSquare");
const TokLeftArrow = value0 => $Token("TokLeftArrow", value0);
const TokRightArrow = value0 => $Token("TokRightArrow", value0);
const TokRightFatArrow = value0 => $Token("TokRightFatArrow", value0);
const TokDoubleColon = value0 => $Token("TokDoubleColon", value0);
const TokForall = value0 => $Token("TokForall", value0);
const TokEquals = /* #__PURE__ */ $Token("TokEquals");
const TokPipe = /* #__PURE__ */ $Token("TokPipe");
const TokTick = /* #__PURE__ */ $Token("TokTick");
const TokDot = /* #__PURE__ */ $Token("TokDot");
const TokComma = /* #__PURE__ */ $Token("TokComma");
const TokUnderscore = /* #__PURE__ */ $Token("TokUnderscore");
const TokBackslash = /* #__PURE__ */ $Token("TokBackslash");
const TokAt = /* #__PURE__ */ $Token("TokAt");
const TokLowerName = value0 => value1 => $Token("TokLowerName", value0, value1);
const TokUpperName = value0 => value1 => $Token("TokUpperName", value0, value1);
const TokOperator = value0 => value1 => $Token("TokOperator", value0, value1);
const TokSymbolName = value0 => value1 => $Token("TokSymbolName", value0, value1);
const TokSymbolArrow = value0 => $Token("TokSymbolArrow", value0);
const TokHole = value0 => $Token("TokHole", value0);
const TokChar = value0 => value1 => $Token("TokChar", value0, value1);
const TokString = value0 => value1 => $Token("TokString", value0, value1);
const TokRawString = value0 => $Token("TokRawString", value0);
const TokInt = value0 => value1 => $Token("TokInt", value0, value1);
const TokNumber = value0 => value1 => $Token("TokNumber", value0, value1);
const TokLayoutStart = value0 => $Token("TokLayoutStart", value0);
const TokLayoutSep = value0 => $Token("TokLayoutSep", value0);
const TokLayoutEnd = value0 => $Token("TokLayoutEnd", value0);
const Ident = x => x;
const Infix = /* #__PURE__ */ $Fixity("Infix");
const Infixl = /* #__PURE__ */ $Fixity("Infixl");
const Infixr = /* #__PURE__ */ $Fixity("Infixr");
const Comment = value0 => $Comment("Comment", value0);
const Space = value0 => $Comment("Space", value0);
const Line = value0 => value1 => $Comment("Line", value0, value1);
const Labeled = x => x;
const Name = x => x;
const Prefixed = x => x;
const QualifiedName = x => x;
const FixityValue = value0 => value1 => value2 => $FixityOp("FixityValue", value0, value1, value2);
const FixityType = value0 => value1 => value2 => value3 => $FixityOp("FixityType", value0, value1, value2, value3);
const RecordPun = value0 => $RecordLabeled("RecordPun", value0);
const RecordField = value0 => value1 => value2 => $RecordLabeled("RecordField", value0, value1, value2);
const Separated = x => x;
const Wrapped = x => x;
const DataAll = value0 => $DataMembers("DataAll", value0);
const DataEnumerated = value0 => $DataMembers("DataEnumerated", value0);
const ExportValue = value0 => $Export("ExportValue", value0);
const ExportOp = value0 => $Export("ExportOp", value0);
const ExportType = value0 => value1 => $Export("ExportType", value0, value1);
const ExportTypeOp = value0 => value1 => $Export("ExportTypeOp", value0, value1);
const ExportClass = value0 => value1 => $Export("ExportClass", value0, value1);
const ExportModule = value0 => value1 => $Export("ExportModule", value0, value1);
const ExportError = value0 => $Export("ExportError", value0);
const ImportValue = value0 => $Import("ImportValue", value0);
const ImportOp = value0 => $Import("ImportOp", value0);
const ImportType = value0 => value1 => $Import("ImportType", value0, value1);
const ImportTypeOp = value0 => value1 => $Import("ImportTypeOp", value0, value1);
const ImportClass = value0 => value1 => $Import("ImportClass", value0, value1);
const ImportError = value0 => $Import("ImportError", value0);
const ImportDecl = x => x;
const ModuleHeader = x => x;
const One = value0 => $OneOrDelimited("One", value0);
const Many = value0 => $OneOrDelimited("Many", value0);
const TypeVarKinded = value0 => $TypeVarBinding("TypeVarKinded", value0);
const TypeVarName = value0 => $TypeVarBinding("TypeVarName", value0);
const TypeVar = value0 => $Type("TypeVar", value0);
const TypeConstructor = value0 => $Type("TypeConstructor", value0);
const TypeWildcard = value0 => $Type("TypeWildcard", value0);
const TypeHole = value0 => $Type("TypeHole", value0);
const TypeString = value0 => value1 => $Type("TypeString", value0, value1);
const TypeInt = value0 => value1 => value2 => $Type("TypeInt", value0, value1, value2);
const TypeRow = value0 => $Type("TypeRow", value0);
const TypeRecord = value0 => $Type("TypeRecord", value0);
const TypeForall = value0 => value1 => value2 => value3 => $Type("TypeForall", value0, value1, value2, value3);
const TypeKinded = value0 => value1 => value2 => $Type("TypeKinded", value0, value1, value2);
const TypeApp = value0 => value1 => $Type("TypeApp", value0, value1);
const TypeOp = value0 => value1 => $Type("TypeOp", value0, value1);
const TypeOpName = value0 => $Type("TypeOpName", value0);
const TypeArrow = value0 => value1 => value2 => $Type("TypeArrow", value0, value1, value2);
const TypeArrowName = value0 => $Type("TypeArrowName", value0);
const TypeConstrained = value0 => value1 => value2 => $Type("TypeConstrained", value0, value1, value2);
const TypeParens = value0 => $Type("TypeParens", value0);
const $$TypeError = value0 => $Type("TypeError", value0);
const Row = x => x;
const DataCtor = x => x;
const ForeignValue = value0 => $Foreign("ForeignValue", value0);
const ForeignData = value0 => value1 => $Foreign("ForeignData", value0, value1);
const ForeignKind = value0 => value1 => $Foreign("ForeignKind", value0, value1);
const FundepDetermined = value0 => value1 => $ClassFundep("FundepDetermined", value0, value1);
const FundepDetermines = value0 => value1 => value2 => $ClassFundep("FundepDetermines", value0, value1, value2);
const BinderWildcard = value0 => $Binder("BinderWildcard", value0);
const BinderVar = value0 => $Binder("BinderVar", value0);
const BinderNamed = value0 => value1 => value2 => $Binder("BinderNamed", value0, value1, value2);
const BinderConstructor = value0 => value1 => $Binder("BinderConstructor", value0, value1);
const BinderBoolean = value0 => value1 => $Binder("BinderBoolean", value0, value1);
const BinderChar = value0 => value1 => $Binder("BinderChar", value0, value1);
const BinderString = value0 => value1 => $Binder("BinderString", value0, value1);
const BinderInt = value0 => value1 => value2 => $Binder("BinderInt", value0, value1, value2);
const BinderNumber = value0 => value1 => value2 => $Binder("BinderNumber", value0, value1, value2);
const BinderArray = value0 => $Binder("BinderArray", value0);
const BinderRecord = value0 => $Binder("BinderRecord", value0);
const BinderParens = value0 => $Binder("BinderParens", value0);
const BinderTyped = value0 => value1 => value2 => $Binder("BinderTyped", value0, value1, value2);
const BinderOp = value0 => value1 => $Binder("BinderOp", value0, value1);
const BinderError = value0 => $Binder("BinderError", value0);
const AppType = value0 => value1 => $AppSpine("AppType", value0, value1);
const AppTerm = value0 => $AppSpine("AppTerm", value0);
const DoLet = value0 => value1 => $DoStatement("DoLet", value0, value1);
const DoDiscard = value0 => $DoStatement("DoDiscard", value0);
const DoBind = value0 => value1 => value2 => $DoStatement("DoBind", value0, value1, value2);
const DoError = value0 => $DoStatement("DoError", value0);
const LetBindingSignature = value0 => $LetBinding("LetBindingSignature", value0);
const LetBindingName = value0 => $LetBinding("LetBindingName", value0);
const LetBindingPattern = value0 => value1 => value2 => $LetBinding("LetBindingPattern", value0, value1, value2);
const LetBindingError = value0 => $LetBinding("LetBindingError", value0);
const Unconditional = value0 => value1 => $Guarded("Unconditional", value0, value1);
const Guarded = value0 => $Guarded("Guarded", value0);
const Where = x => x;
const ExprHole = value0 => $Expr("ExprHole", value0);
const ExprSection = value0 => $Expr("ExprSection", value0);
const ExprIdent = value0 => $Expr("ExprIdent", value0);
const ExprConstructor = value0 => $Expr("ExprConstructor", value0);
const ExprBoolean = value0 => value1 => $Expr("ExprBoolean", value0, value1);
const ExprChar = value0 => value1 => $Expr("ExprChar", value0, value1);
const ExprString = value0 => value1 => $Expr("ExprString", value0, value1);
const ExprInt = value0 => value1 => $Expr("ExprInt", value0, value1);
const ExprNumber = value0 => value1 => $Expr("ExprNumber", value0, value1);
const ExprArray = value0 => $Expr("ExprArray", value0);
const ExprRecord = value0 => $Expr("ExprRecord", value0);
const ExprParens = value0 => $Expr("ExprParens", value0);
const ExprTyped = value0 => value1 => value2 => $Expr("ExprTyped", value0, value1, value2);
const ExprInfix = value0 => value1 => $Expr("ExprInfix", value0, value1);
const ExprOp = value0 => value1 => $Expr("ExprOp", value0, value1);
const ExprOpName = value0 => $Expr("ExprOpName", value0);
const ExprNegate = value0 => value1 => $Expr("ExprNegate", value0, value1);
const ExprRecordAccessor = value0 => $Expr("ExprRecordAccessor", value0);
const ExprRecordUpdate = value0 => value1 => $Expr("ExprRecordUpdate", value0, value1);
const ExprApp = value0 => value1 => $Expr("ExprApp", value0, value1);
const ExprLambda = value0 => $Expr("ExprLambda", value0);
const ExprIf = value0 => $Expr("ExprIf", value0);
const ExprCase = value0 => $Expr("ExprCase", value0);
const ExprLet = value0 => $Expr("ExprLet", value0);
const ExprDo = value0 => $Expr("ExprDo", value0);
const ExprAdo = value0 => $Expr("ExprAdo", value0);
const ExprError = value0 => $Expr("ExprError", value0);
const RecordUpdateLeaf = value0 => value1 => value2 => $RecordUpdate("RecordUpdateLeaf", value0, value1, value2);
const RecordUpdateBranch = value0 => value1 => $RecordUpdate("RecordUpdateBranch", value0, value1);
const GuardedExpr = x => x;
const PatternGuard = x => x;
const InstanceBindingSignature = value0 => $InstanceBinding("InstanceBindingSignature", value0);
const InstanceBindingName = value0 => $InstanceBinding("InstanceBindingName", value0);
const Instance = x => x;
const DeclData = value0 => value1 => $Declaration("DeclData", value0, value1);
const DeclType = value0 => value1 => value2 => $Declaration("DeclType", value0, value1, value2);
const DeclNewtype = value0 => value1 => value2 => value3 => $Declaration("DeclNewtype", value0, value1, value2, value3);
const DeclClass = value0 => value1 => $Declaration("DeclClass", value0, value1);
const DeclInstanceChain = value0 => $Declaration("DeclInstanceChain", value0);
const DeclDerive = value0 => value1 => value2 => $Declaration("DeclDerive", value0, value1, value2);
const DeclKindSignature = value0 => value1 => $Declaration("DeclKindSignature", value0, value1);
const DeclSignature = value0 => $Declaration("DeclSignature", value0);
const DeclValue = value0 => $Declaration("DeclValue", value0);
const DeclFixity = value0 => $Declaration("DeclFixity", value0);
const DeclForeign = value0 => value1 => value2 => $Declaration("DeclForeign", value0, value1, value2);
const DeclRole = value0 => value1 => value2 => value3 => $Declaration("DeclRole", value0, value1, value2, value3);
const DeclError = value0 => $Declaration("DeclError", value0);
const ModuleBody = x => x;
const Module = x => x;
const ordProper = Data$dOrd.ordString;
const ordOperator = Data$dOrd.ordString;
const ordModuleName = Data$dOrd.ordString;
const ordLabel = Data$dOrd.ordString;
const ordIdent = Data$dOrd.ordString;
const newtypeWrapped = {Coercible0: () => {}};
const newtypeWhere = {Coercible0: () => {}};
const newtypeSeparated = {Coercible0: () => {}};
const newtypeRow = {Coercible0: () => {}};
const newtypeQualifiedName = {Coercible0: () => {}};
const newtypeProper = {Coercible0: () => {}};
const newtypePrefixed = {Coercible0: () => {}};
const newtypePatternGuard = {Coercible0: () => {}};
const newtypeOperator = {Coercible0: () => {}};
const newtypeName = {Coercible0: () => {}};
const newtypeModuleName = {Coercible0: () => {}};
const newtypeModuleHeader = {Coercible0: () => {}};
const newtypeModuleBody = {Coercible0: () => {}};
const newtypeModule = {Coercible0: () => {}};
const newtypeLabeled = {Coercible0: () => {}};
const newtypeLabel = {Coercible0: () => {}};
const newtypeInstance = {Coercible0: () => {}};
const newtypeImportDecl = {Coercible0: () => {}};
const newtypeIdent = {Coercible0: () => {}};
const newtypeGuardedExpr = {Coercible0: () => {}};
const newtypeDataCtor = {Coercible0: () => {}};
const eqSourceStyle = {
  eq: x => y => {
    if (x === "ASCII") { return y === "ASCII"; }
    return x === "Unicode" && y === "Unicode";
  }
};
const eqProper = Data$dEq.eqString;
const eqOperator = Data$dEq.eqString;
const eqModuleName = Data$dEq.eqString;
const eqLabel = Data$dEq.eqString;
const eqIntValue = {
  eq: x => y => {
    if (x.tag === "SmallInt") { return y.tag === "SmallInt" && x._1 === y._1; }
    if (x.tag === "BigInt") { return y.tag === "BigInt" && x._1 === y._1; }
    return x.tag === "BigHex" && y.tag === "BigHex" && x._1 === y._1;
  }
};
const eqToken = {
  eq: x => y => {
    if (x.tag === "TokLeftParen") { return y.tag === "TokLeftParen"; }
    if (x.tag === "TokRightParen") { return y.tag === "TokRightParen"; }
    if (x.tag === "TokLeftBrace") { return y.tag === "TokLeftBrace"; }
    if (x.tag === "TokRightBrace") { return y.tag === "TokRightBrace"; }
    if (x.tag === "TokLeftSquare") { return y.tag === "TokLeftSquare"; }
    if (x.tag === "TokRightSquare") { return y.tag === "TokRightSquare"; }
    if (x.tag === "TokLeftArrow") { return y.tag === "TokLeftArrow" && (x._1 === "ASCII" ? y._1 === "ASCII" : x._1 === "Unicode" && y._1 === "Unicode"); }
    if (x.tag === "TokRightArrow") { return y.tag === "TokRightArrow" && (x._1 === "ASCII" ? y._1 === "ASCII" : x._1 === "Unicode" && y._1 === "Unicode"); }
    if (x.tag === "TokRightFatArrow") { return y.tag === "TokRightFatArrow" && (x._1 === "ASCII" ? y._1 === "ASCII" : x._1 === "Unicode" && y._1 === "Unicode"); }
    if (x.tag === "TokDoubleColon") { return y.tag === "TokDoubleColon" && (x._1 === "ASCII" ? y._1 === "ASCII" : x._1 === "Unicode" && y._1 === "Unicode"); }
    if (x.tag === "TokForall") { return y.tag === "TokForall" && (x._1 === "ASCII" ? y._1 === "ASCII" : x._1 === "Unicode" && y._1 === "Unicode"); }
    if (x.tag === "TokEquals") { return y.tag === "TokEquals"; }
    if (x.tag === "TokPipe") { return y.tag === "TokPipe"; }
    if (x.tag === "TokTick") { return y.tag === "TokTick"; }
    if (x.tag === "TokDot") { return y.tag === "TokDot"; }
    if (x.tag === "TokComma") { return y.tag === "TokComma"; }
    if (x.tag === "TokUnderscore") { return y.tag === "TokUnderscore"; }
    if (x.tag === "TokBackslash") { return y.tag === "TokBackslash"; }
    if (x.tag === "TokAt") { return y.tag === "TokAt"; }
    if (x.tag === "TokLowerName") {
      return y.tag === "TokLowerName" && (x._1.tag === "Nothing" ? y._1.tag === "Nothing" : x._1.tag === "Just" && y._1.tag === "Just" && x._1._1 === y._1._1) && x._2 === y._2;
    }
    if (x.tag === "TokUpperName") {
      return y.tag === "TokUpperName" && (x._1.tag === "Nothing" ? y._1.tag === "Nothing" : x._1.tag === "Just" && y._1.tag === "Just" && x._1._1 === y._1._1) && x._2 === y._2;
    }
    if (x.tag === "TokOperator") {
      return y.tag === "TokOperator" && (x._1.tag === "Nothing" ? y._1.tag === "Nothing" : x._1.tag === "Just" && y._1.tag === "Just" && x._1._1 === y._1._1) && x._2 === y._2;
    }
    if (x.tag === "TokSymbolName") {
      return y.tag === "TokSymbolName" && (x._1.tag === "Nothing" ? y._1.tag === "Nothing" : x._1.tag === "Just" && y._1.tag === "Just" && x._1._1 === y._1._1) && x._2 === y._2;
    }
    if (x.tag === "TokSymbolArrow") { return y.tag === "TokSymbolArrow" && (x._1 === "ASCII" ? y._1 === "ASCII" : x._1 === "Unicode" && y._1 === "Unicode"); }
    if (x.tag === "TokHole") { return y.tag === "TokHole" && x._1 === y._1; }
    if (x.tag === "TokChar") { return y.tag === "TokChar" && x._1 === y._1 && x._2 === y._2; }
    if (x.tag === "TokString") { return y.tag === "TokString" && x._1 === y._1 && x._2 === y._2; }
    if (x.tag === "TokRawString") { return y.tag === "TokRawString" && x._1 === y._1; }
    if (x.tag === "TokInt") {
      return y.tag === "TokInt" && x._1 === y._1 && (() => {
        if (x._2.tag === "SmallInt") { return y._2.tag === "SmallInt" && x._2._1 === y._2._1; }
        if (x._2.tag === "BigInt") { return y._2.tag === "BigInt" && x._2._1 === y._2._1; }
        return x._2.tag === "BigHex" && y._2.tag === "BigHex" && x._2._1 === y._2._1;
      })();
    }
    if (x.tag === "TokNumber") { return y.tag === "TokNumber" && x._1 === y._1 && x._2 === y._2; }
    if (x.tag === "TokLayoutStart") { return y.tag === "TokLayoutStart" && x._1 === y._1; }
    if (x.tag === "TokLayoutSep") { return y.tag === "TokLayoutSep" && x._1 === y._1; }
    return x.tag === "TokLayoutEnd" && y.tag === "TokLayoutEnd" && x._1 === y._1;
  }
};
const eqIdent = Data$dEq.eqString;
export {
  $AppSpine,
  $Binder,
  $ClassFundep,
  $Comment,
  $DataMembers,
  $Declaration,
  $DoStatement,
  $Export,
  $Expr,
  $Fixity,
  $FixityOp,
  $Foreign,
  $Guarded,
  $Import,
  $InstanceBinding,
  $IntValue,
  $LetBinding,
  $LineFeed,
  $OneOrDelimited,
  $RecordLabeled,
  $RecordUpdate,
  $Role,
  $SourceStyle,
  $Token,
  $Type,
  $TypeVarBinding,
  ASCII,
  AppTerm,
  AppType,
  BigHex,
  $$BigInt as BigInt,
  BinderArray,
  BinderBoolean,
  BinderChar,
  BinderConstructor,
  BinderError,
  BinderInt,
  BinderNamed,
  BinderNumber,
  BinderOp,
  BinderParens,
  BinderRecord,
  BinderString,
  BinderTyped,
  BinderVar,
  BinderWildcard,
  CRLF,
  Comment,
  DataAll,
  DataCtor,
  DataEnumerated,
  DeclClass,
  DeclData,
  DeclDerive,
  DeclError,
  DeclFixity,
  DeclForeign,
  DeclInstanceChain,
  DeclKindSignature,
  DeclNewtype,
  DeclRole,
  DeclSignature,
  DeclType,
  DeclValue,
  DoBind,
  DoDiscard,
  DoError,
  DoLet,
  ExportClass,
  ExportError,
  ExportModule,
  ExportOp,
  ExportType,
  ExportTypeOp,
  ExportValue,
  ExprAdo,
  ExprApp,
  ExprArray,
  ExprBoolean,
  ExprCase,
  ExprChar,
  ExprConstructor,
  ExprDo,
  ExprError,
  ExprHole,
  ExprIdent,
  ExprIf,
  ExprInfix,
  ExprInt,
  ExprLambda,
  ExprLet,
  ExprNegate,
  ExprNumber,
  ExprOp,
  ExprOpName,
  ExprParens,
  ExprRecord,
  ExprRecordAccessor,
  ExprRecordUpdate,
  ExprSection,
  ExprString,
  ExprTyped,
  FixityType,
  FixityValue,
  ForeignData,
  ForeignKind,
  ForeignValue,
  FundepDetermined,
  FundepDetermines,
  Guarded,
  GuardedExpr,
  Ident,
  ImportClass,
  ImportDecl,
  ImportError,
  ImportOp,
  ImportType,
  ImportTypeOp,
  ImportValue,
  Infix,
  Infixl,
  Infixr,
  Instance,
  InstanceBindingName,
  InstanceBindingSignature,
  LF,
  Label,
  Labeled,
  LetBindingError,
  LetBindingName,
  LetBindingPattern,
  LetBindingSignature,
  Line,
  Many,
  Module,
  ModuleBody,
  ModuleHeader,
  ModuleName,
  Name,
  Nominal,
  One,
  Operator,
  PatternGuard,
  Phantom,
  Prefixed,
  Proper,
  QualifiedName,
  RecordField,
  RecordPun,
  RecordUpdateBranch,
  RecordUpdateLeaf,
  Representational,
  Row,
  Separated,
  SmallInt,
  Space,
  TokAt,
  TokBackslash,
  TokChar,
  TokComma,
  TokDot,
  TokDoubleColon,
  TokEquals,
  TokForall,
  TokHole,
  TokInt,
  TokLayoutEnd,
  TokLayoutSep,
  TokLayoutStart,
  TokLeftArrow,
  TokLeftBrace,
  TokLeftParen,
  TokLeftSquare,
  TokLowerName,
  TokNumber,
  TokOperator,
  TokPipe,
  TokRawString,
  TokRightArrow,
  TokRightBrace,
  TokRightFatArrow,
  TokRightParen,
  TokRightSquare,
  TokString,
  TokSymbolArrow,
  TokSymbolName,
  TokTick,
  TokUnderscore,
  TokUpperName,
  TypeApp,
  TypeArrow,
  TypeArrowName,
  TypeConstrained,
  TypeConstructor,
  $$TypeError as TypeError,
  TypeForall,
  TypeHole,
  TypeInt,
  TypeKinded,
  TypeOp,
  TypeOpName,
  TypeParens,
  TypeRecord,
  TypeRow,
  TypeString,
  TypeVar,
  TypeVarKinded,
  TypeVarName,
  TypeWildcard,
  Unconditional,
  Unicode,
  Where,
  Wrapped,
  eqIdent,
  eqIntValue,
  eqLabel,
  eqModuleName,
  eqOperator,
  eqProper,
  eqSourceStyle,
  eqToken,
  newtypeDataCtor,
  newtypeGuardedExpr,
  newtypeIdent,
  newtypeImportDecl,
  newtypeInstance,
  newtypeLabel,
  newtypeLabeled,
  newtypeModule,
  newtypeModuleBody,
  newtypeModuleHeader,
  newtypeModuleName,
  newtypeName,
  newtypeOperator,
  newtypePatternGuard,
  newtypePrefixed,
  newtypeProper,
  newtypeQualifiedName,
  newtypeRow,
  newtypeSeparated,
  newtypeWhere,
  newtypeWrapped,
  ordIdent,
  ordLabel,
  ordModuleName,
  ordOperator,
  ordProper
};
