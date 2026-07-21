// @inline Data.Argonaut.Core.caseJson always
import * as $runtime from "../runtime.js";
import * as Control$dBind from "../Control.Bind/index.js";
import * as Data$dArgonaut$dCore from "../Data.Argonaut.Core/index.js";
import * as Data$dArgonaut$dDecode$dDecoders from "../Data.Argonaut.Decode.Decoders/index.js";
import * as Data$dArgonaut$dDecode$dError from "../Data.Argonaut.Decode.Error/index.js";
import * as Data$dEither from "../Data.Either/index.js";
import * as Data$dFoldable from "../Data.Foldable/index.js";
import * as Data$dFunctor from "../Data.Functor/index.js";
import * as Data$dInt from "../Data.Int/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dString$dCodePoints from "../Data.String.CodePoints/index.js";
import * as Data$dString$dCodeUnits from "../Data.String.CodeUnits/index.js";
import * as Data$dTraversable from "../Data.Traversable/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
import * as Foreign$dObject from "../Foreign.Object/index.js";
import * as PureScript$dBackend$dOptimizer$dCoreFn from "../PureScript.Backend.Optimizer.CoreFn/index.js";
const traverse = /* #__PURE__ */ (() => Data$dTraversable.traversableArray.traverse(Data$dEither.applicativeEither))();
const getFieldOptional$p = decode => obj => prop => {
  const v = Foreign$dObject._lookup(Data$dMaybe.Nothing, Data$dMaybe.Just, prop, obj);
  if (v.tag === "Nothing") { return Data$dEither.$Either("Right", Data$dMaybe.Nothing); }
  if (v.tag === "Just") {
    if (Data$dArgonaut$dCore._caseJson(v$1 => true, v$1 => false, v$1 => false, v$1 => false, v$1 => false, v$1 => false, v._1)) {
      return Data$dEither.$Either("Right", Data$dMaybe.Nothing);
    }
    const $0 = decode(v._1);
    if ($0.tag === "Left") { return Data$dEither.$Either("Left", $0._1); }
    if ($0.tag === "Right") { return Data$dEither.$Either("Right", Data$dMaybe.$Maybe("Just", $0._1)); }
  }
  $runtime.fail();
};
const getField = decode => obj => prop => {
  const v = Foreign$dObject._lookup(Data$dMaybe.Nothing, Data$dMaybe.Just, prop, obj);
  if (v.tag === "Nothing") { return Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("AtKey", prop, Data$dArgonaut$dDecode$dError.MissingValue)); }
  if (v.tag === "Just") { return decode(v._1); }
  $runtime.fail();
};
const decodeString = json => Data$dArgonaut$dCore._caseJson(
  v => Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "String")),
  v => Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "String")),
  v => Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "String")),
  Data$dEither.Right,
  v => Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "String")),
  v => Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "String")),
  json
);
const decodeNumber = json => Data$dArgonaut$dCore._caseJson(
  v => Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "Number")),
  v => Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "Number")),
  Data$dEither.Right,
  v => Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "Number")),
  v => Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "Number")),
  v => Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "Number")),
  json
);
const decodeJObject = json => Data$dArgonaut$dCore._caseJson(
  v => Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "Object")),
  v => Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "Object")),
  v => Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "Object")),
  v => Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "Object")),
  v => Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "Object")),
  Data$dEither.Right,
  json
);
const decodeJArray = json => Data$dArgonaut$dCore._caseJson(
  v => Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "Array")),
  v => Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "Array")),
  v => Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "Array")),
  v => Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "Array")),
  Data$dEither.Right,
  v => Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "Array")),
  json
);
const decodeBoolean = json => Data$dArgonaut$dCore._caseJson(
  v => Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "Boolean")),
  Data$dEither.Right,
  v => Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "Boolean")),
  v => Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "Boolean")),
  v => Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "Boolean")),
  v => Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "Boolean")),
  json
);
const decodeArray = decoder => json => {
  const v = decodeJArray(json);
  if (v.tag === "Left") { return Data$dEither.$Either("Left", v._1); }
  if (v.tag === "Right") {
    const $0 = v._1;
    return (() => {
      const out = [];
      let ix = 0;
      let con = true;
      let res = undefined;
      const len = $0.length;
      while (con) {
        const ix$p = ix;
        if (ix$p === len) {
          con = false;
          res = Data$dEither.$Either("Right", out);
          continue;
        }
        const v1 = decoder($0[ix$p]);
        if (v1.tag === "Left") {
          con = false;
          res = Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("AtIndex", ix$p, v1._1));
          continue;
        }
        if (v1.tag === "Right") {
          out.push(v1._1);
          ix = ix$p + 1 | 0;
          continue;
        }
        $runtime.fail();
      }
      return res;
    })();
  }
  $runtime.fail();
};
const decodeModuleName = x => {
  const $0 = decodeArray(decodeString)(x);
  if ($0.tag === "Left") { return Data$dEither.$Either("Left", $0._1); }
  if ($0.tag === "Right") {
    return Data$dEither.$Either(
      "Right",
      Data$dFoldable.foldlArray(v => v1 => {
        if (v.init) { return {init: false, acc: v1}; }
        return {init: false, acc: v.acc + "." + v1};
      })({init: true, acc: ""})($0._1).acc
    );
  }
  $runtime.fail();
};
const decodeConstructorType = json => {
  const $0 = decodeString(json);
  if ($0.tag === "Left") { return Data$dEither.$Either("Left", $0._1); }
  if ($0.tag === "Right") {
    if ($0._1 === "ProductType") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.ProductType); }
    if ($0._1 === "SumType") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.SumType); }
    return Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "ConstructorType"));
  }
  $runtime.fail();
};
const decodeImport = decodeAnn$p => json => {
  const $0 = decodeJObject(json);
  if ($0.tag === "Left") { return Data$dEither.$Either("Left", $0._1); }
  if ($0.tag === "Right") {
    const $1 = getField(decodeAnn$p)($0._1)("annotation");
    if ($1.tag === "Left") { return Data$dEither.$Either("Left", $1._1); }
    if ($1.tag === "Right") {
      const $2 = getField(decodeModuleName)($0._1)("moduleName");
      if ($2.tag === "Left") { return Data$dEither.$Either("Left", $2._1); }
      if ($2.tag === "Right") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.$Import($1._1, $2._1)); }
    }
  }
  $runtime.fail();
};
const decodeInt = json => {
  const $0 = decodeNumber(json);
  if ($0.tag === "Left") { return Data$dEither.$Either("Left", $0._1); }
  if ($0.tag === "Right") {
    const v = Data$dInt.fromNumber($0._1);
    if (v.tag === "Nothing") { return Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "Int")); }
    if (v.tag === "Just") { return Data$dEither.$Either("Right", v._1); }
  }
  $runtime.fail();
};
const decodeCodePoint = a => {
  const $0 = decodeInt(a);
  if ($0.tag === "Left") { return Data$dEither.$Either("Left", $0._1); }
  if ($0.tag === "Right") {
    if ($0._1 >= 0 && $0._1 <= 1114111) { return Data$dEither.$Either("Right", $0._1); }
    return Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "CodePoint"));
  }
  $runtime.fail();
};
const decodeMeta = json => {
  const $0 = decodeJObject(json);
  if ($0.tag === "Left") { return Data$dEither.$Either("Left", $0._1); }
  if ($0.tag === "Right") {
    const $1 = getField(decodeString)($0._1)("metaType");
    if ($1.tag === "Left") { return Data$dEither.$Either("Left", $1._1); }
    if ($1.tag === "Right") {
      if ($1._1 === "IsConstructor") {
        const $2 = getField(decodeConstructorType)($0._1)("constructorType");
        if ($2.tag === "Left") { return Data$dEither.$Either("Left", $2._1); }
        if ($2.tag === "Right") {
          const $3 = getField(decodeArray(decodeString))($0._1)("identifiers");
          if ($3.tag === "Left") { return Data$dEither.$Either("Left", $3._1); }
          if ($3.tag === "Right") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.$Meta("IsConstructor", $2._1, $3._1)); }
        }
        $runtime.fail();
      }
      if ($1._1 === "IsNewtype") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.IsNewtype); }
      if ($1._1 === "IsTypeClassConstructor") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.IsTypeClassConstructor); }
      if ($1._1 === "IsForeign") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.IsForeign); }
      if ($1._1 === "IsWhere") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.IsWhere); }
      if ($1._1 === "IsSyntheticApp") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.IsSyntheticApp); }
      return Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "Meta"));
    }
  }
  $runtime.fail();
};
const decodeAnn = _path => json => {
  const $0 = decodeJObject(json);
  if ($0.tag === "Left") { return Data$dEither.$Either("Left", $0._1); }
  if ($0.tag === "Right") {
    const $1 = getFieldOptional$p(decodeMeta)($0._1)("meta");
    if ($1.tag === "Left") { return Data$dEither.$Either("Left", $1._1); }
    if ($1.tag === "Right") { return Data$dEither.$Either("Right", {span: PureScript$dBackend$dOptimizer$dCoreFn.emptySpan, meta: $1._1}); }
  }
  $runtime.fail();
};
const decodeQualified = k => json => {
  const $0 = decodeJObject(json);
  if ($0.tag === "Left") { return Data$dEither.$Either("Left", $0._1); }
  if ($0.tag === "Right") {
    const $1 = getFieldOptional$p(decodeModuleName)($0._1)("moduleName");
    if ($1.tag === "Left") { return Data$dEither.$Either("Left", $1._1); }
    if ($1.tag === "Right") {
      const $2 = getField(k)($0._1)("identifier");
      if ($2.tag === "Left") { return Data$dEither.$Either("Left", $2._1); }
      if ($2.tag === "Right") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.$Qualified($1._1, $2._1)); }
    }
  }
  $runtime.fail();
};
const decodeReExports = json => {
  const $0 = decodeJObject(json);
  if ($0.tag === "Left") { return Data$dEither.$Either("Left", $0._1); }
  if ($0.tag === "Right") {
    const $1 = traverse(Data$dTraversable.traversableTuple.traverse(Data$dEither.applicativeEither)(decodeArray(decodeString)))(Foreign$dObject.toArrayWithKey(Data$dTuple.Tuple)($0._1));
    if ($1.tag === "Left") { return Data$dEither.$Either("Left", $1._1); }
    if ($1.tag === "Right") {
      return Data$dEither.$Either("Right", Control$dBind.arrayBind($1._1)(v => Data$dFunctor.arrayMap(PureScript$dBackend$dOptimizer$dCoreFn.ReExport(v._1))(v._2)));
    }
  }
  $runtime.fail();
};
const decodeRecord = x => decodeArray(json => {
  const $0 = decodeJArray(json);
  if ($0.tag === "Left") { return Data$dEither.$Either("Left", $0._1); }
  if ($0.tag === "Right") {
    if ($0._1.length === 2) {
      const $1 = decodeString($0._1[0]);
      if ($1.tag === "Left") { return Data$dEither.$Either("Left", $1._1); }
      if ($1.tag === "Right") {
        const $2 = x($0._1[1]);
        if ($2.tag === "Left") { return Data$dEither.$Either("Left", $2._1); }
        if ($2.tag === "Right") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.$Prop($1._1, $2._1)); }
      }
      $runtime.fail();
    }
    return Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "Tuple"));
  }
  $runtime.fail();
});
const decodeSourcePos = json => {
  const $0 = Data$dArgonaut$dDecode$dDecoders.decodeTuple(Data$dArgonaut$dDecode$dDecoders.decodeInt)(Data$dArgonaut$dDecode$dDecoders.decodeInt)(json);
  if ($0.tag === "Left") { return Data$dEither.$Either("Left", $0._1); }
  if ($0.tag === "Right") { return Data$dEither.$Either("Right", {line: $0._1._1, column: $0._1._2}); }
  $runtime.fail();
};
const decodeSourceSpan = path => json => {
  const $0 = decodeJObject(json);
  if ($0.tag === "Left") { return Data$dEither.$Either("Left", $0._1); }
  if ($0.tag === "Right") {
    const $1 = getField(decodeSourcePos)($0._1)("start");
    if ($1.tag === "Left") { return Data$dEither.$Either("Left", $1._1); }
    if ($1.tag === "Right") {
      const $2 = getField(decodeSourcePos)($0._1)("end");
      if ($2.tag === "Left") { return Data$dEither.$Either("Left", $2._1); }
      if ($2.tag === "Right") { return Data$dEither.$Either("Right", {path, start: $1._1, end: $2._1}); }
    }
  }
  $runtime.fail();
};
const decodeComment = json => {
  const $0 = decodeJObject(json);
  if ($0.tag === "Left") { return Data$dEither.$Either("Left", $0._1); }
  if ($0.tag === "Right") {
    const $1 = getField(decodeString)($0._1)("LineComment");
    const $2 = (() => {
      if ($1.tag === "Left") { return Data$dEither.$Either("Left", $1._1); }
      if ($1.tag === "Right") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.$Comment("LineComment", $1._1)); }
      $runtime.fail();
    })();
    if ($2.tag === "Left") {
      const $3 = getField(decodeString)($0._1)("BlockComment");
      if ($3.tag === "Left") { return Data$dEither.$Either("Left", $3._1); }
      if ($3.tag === "Right") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.$Comment("BlockComment", $3._1)); }
      $runtime.fail();
    }
    if ($2.tag === "Right") { return $2; }
  }
  $runtime.fail();
};
const decodeStringLiteral = json => {
  const $0 = decodeString(json);
  if ($0.tag === "Left") {
    const $1 = decodeArray(decodeCodePoint)(json);
    const $2 = (() => {
      if ($1.tag === "Left") { return Data$dEither.$Either("Left", $1._1); }
      if ($1.tag === "Right") { return Data$dEither.$Either("Right", Data$dString$dCodePoints.fromCodePointArray($1._1)); }
      $runtime.fail();
    })();
    if ($2.tag === "Left") { return Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "StringLiteral")); }
    if ($2.tag === "Right") { return $2; }
    $runtime.fail();
  }
  if ($0.tag === "Right") { return $0; }
  $runtime.fail();
};
const decodeLiteral = dec => json => {
  const $0 = decodeJObject(json);
  if ($0.tag === "Left") { return Data$dEither.$Either("Left", $0._1); }
  if ($0.tag === "Right") {
    const $1 = getField(decodeString)($0._1)("literalType");
    if ($1.tag === "Left") { return Data$dEither.$Either("Left", $1._1); }
    if ($1.tag === "Right") {
      if ($1._1 === "IntLiteral") {
        const $2 = getField(decodeInt)($0._1)("value");
        if ($2.tag === "Left") { return Data$dEither.$Either("Left", $2._1); }
        if ($2.tag === "Right") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitInt", $2._1)); }
        $runtime.fail();
      }
      if ($1._1 === "NumberLiteral") {
        const $2 = getField(decodeNumber)($0._1)("value");
        if ($2.tag === "Left") { return Data$dEither.$Either("Left", $2._1); }
        if ($2.tag === "Right") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitNumber", $2._1)); }
        $runtime.fail();
      }
      if ($1._1 === "StringLiteral") {
        const $2 = getField(decodeStringLiteral)($0._1)("value");
        if ($2.tag === "Left") { return Data$dEither.$Either("Left", $2._1); }
        if ($2.tag === "Right") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitString", $2._1)); }
        $runtime.fail();
      }
      if ($1._1 === "CharLiteral") {
        const $2 = getField(decodeString)($0._1)("value");
        if ($2.tag === "Left") { return Data$dEither.$Either("Left", $2._1); }
        if ($2.tag === "Right") {
          if (Data$dString$dCodeUnits.length($2._1) === 1) {
            const $3 = Data$dString$dCodeUnits.toCharArray($2._1);
            if (0 < $3.length) { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitChar", $3[0])); }
          }
          return Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "Char"));
        }
        $runtime.fail();
      }
      if ($1._1 === "BooleanLiteral") {
        const $2 = getField(decodeBoolean)($0._1)("value");
        if ($2.tag === "Left") { return Data$dEither.$Either("Left", $2._1); }
        if ($2.tag === "Right") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitBoolean", $2._1)); }
        $runtime.fail();
      }
      if ($1._1 === "ArrayLiteral") {
        const $2 = getField(decodeArray(dec))($0._1)("value");
        if ($2.tag === "Left") { return Data$dEither.$Either("Left", $2._1); }
        if ($2.tag === "Right") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitArray", $2._1)); }
        $runtime.fail();
      }
      if ($1._1 === "ObjectLiteral") {
        const $2 = getField(decodeRecord(dec))($0._1)("value");
        if ($2.tag === "Left") { return Data$dEither.$Either("Left", $2._1); }
        if ($2.tag === "Right") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.$Literal("LitRecord", $2._1)); }
        $runtime.fail();
      }
      return Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "Literal"));
    }
  }
  $runtime.fail();
};
const decodeBinder = decAnn => json => {
  const $0 = decodeJObject(json);
  if ($0.tag === "Left") { return Data$dEither.$Either("Left", $0._1); }
  if ($0.tag === "Right") {
    const $1 = getField(decAnn)($0._1)("annotation");
    if ($1.tag === "Left") { return Data$dEither.$Either("Left", $1._1); }
    if ($1.tag === "Right") {
      const $2 = getField(decodeString)($0._1)("binderType");
      if ($2.tag === "Left") { return Data$dEither.$Either("Left", $2._1); }
      if ($2.tag === "Right") {
        if ($2._1 === "NullBinder") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.$Binder("BinderNull", $1._1)); }
        if ($2._1 === "VarBinder") {
          const $3 = PureScript$dBackend$dOptimizer$dCoreFn.BinderVar($1._1);
          const $4 = getField(decodeString)($0._1)("identifier");
          if ($4.tag === "Left") { return Data$dEither.$Either("Left", $4._1); }
          if ($4.tag === "Right") { return Data$dEither.$Either("Right", $3($4._1)); }
          $runtime.fail();
        }
        if ($2._1 === "LiteralBinder") {
          const $3 = PureScript$dBackend$dOptimizer$dCoreFn.BinderLit($1._1);
          const $4 = getField(decodeLiteral(decodeBinder(decAnn)))($0._1)("literal");
          if ($4.tag === "Left") { return Data$dEither.$Either("Left", $4._1); }
          if ($4.tag === "Right") { return Data$dEither.$Either("Right", $3($4._1)); }
          $runtime.fail();
        }
        if ($2._1 === "ConstructorBinder") {
          const $3 = getField(decodeQualified(decodeString))($0._1)("typeName");
          if ($3.tag === "Left") { return Data$dEither.$Either("Left", $3._1); }
          if ($3.tag === "Right") {
            const $4 = getField(decodeQualified(decodeString))($0._1)("constructorName");
            if ($4.tag === "Left") { return Data$dEither.$Either("Left", $4._1); }
            if ($4.tag === "Right") {
              const $5 = getField(decodeArray(decodeBinder(decAnn)))($0._1)("binders");
              if ($5.tag === "Left") { return Data$dEither.$Either("Left", $5._1); }
              if ($5.tag === "Right") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.$Binder("BinderConstructor", $1._1, $3._1, $4._1, $5._1)); }
            }
          }
          $runtime.fail();
        }
        if ($2._1 === "NamedBinder") {
          const $3 = getField(decodeString)($0._1)("identifier");
          if ($3.tag === "Left") { return Data$dEither.$Either("Left", $3._1); }
          if ($3.tag === "Right") {
            const $4 = getField(decodeBinder(decAnn))($0._1)("binder");
            if ($4.tag === "Left") { return Data$dEither.$Either("Left", $4._1); }
            if ($4.tag === "Right") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.$Binder("BinderNamed", $1._1, $3._1, $4._1)); }
          }
          $runtime.fail();
        }
        return Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "Binder"));
      }
    }
  }
  $runtime.fail();
};
const decodeGuard = decAnn => json => {
  const $0 = decodeJObject(json);
  if ($0.tag === "Left") { return Data$dEither.$Either("Left", $0._1); }
  if ($0.tag === "Right") {
    const $1 = getField(decodeExpr(decAnn))($0._1)("guard");
    if ($1.tag === "Left") { return Data$dEither.$Either("Left", $1._1); }
    if ($1.tag === "Right") {
      const $2 = getField(decodeExpr(decAnn))($0._1)("expression");
      if ($2.tag === "Left") { return Data$dEither.$Either("Left", $2._1); }
      if ($2.tag === "Right") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.$Guard($1._1, $2._1)); }
    }
  }
  $runtime.fail();
};
const decodeExpr = decAnn => json => {
  const $0 = decodeJObject(json);
  if ($0.tag === "Left") { return Data$dEither.$Either("Left", $0._1); }
  if ($0.tag === "Right") {
    const $1 = getField(decAnn)($0._1)("annotation");
    if ($1.tag === "Left") { return Data$dEither.$Either("Left", $1._1); }
    if ($1.tag === "Right") {
      const $2 = getField(decodeString)($0._1)("type");
      if ($2.tag === "Left") { return Data$dEither.$Either("Left", $2._1); }
      if ($2.tag === "Right") {
        if ($2._1 === "Var") {
          const $3 = PureScript$dBackend$dOptimizer$dCoreFn.ExprVar($1._1);
          const $4 = getField(decodeQualified(decodeString))($0._1)("value");
          if ($4.tag === "Left") { return Data$dEither.$Either("Left", $4._1); }
          if ($4.tag === "Right") { return Data$dEither.$Either("Right", $3($4._1)); }
          $runtime.fail();
        }
        if ($2._1 === "Literal") {
          const $3 = PureScript$dBackend$dOptimizer$dCoreFn.ExprLit($1._1);
          const $4 = getField(decodeLiteral(decodeExpr(decAnn)))($0._1)("value");
          if ($4.tag === "Left") { return Data$dEither.$Either("Left", $4._1); }
          if ($4.tag === "Right") { return Data$dEither.$Either("Right", $3($4._1)); }
          $runtime.fail();
        }
        if ($2._1 === "Constructor") {
          const $3 = getField(decodeString)($0._1)("typeName");
          if ($3.tag === "Left") { return Data$dEither.$Either("Left", $3._1); }
          if ($3.tag === "Right") {
            const $4 = getField(decodeString)($0._1)("constructorName");
            if ($4.tag === "Left") { return Data$dEither.$Either("Left", $4._1); }
            if ($4.tag === "Right") {
              const $5 = getField(decodeArray(decodeString))($0._1)("fieldNames");
              if ($5.tag === "Left") { return Data$dEither.$Either("Left", $5._1); }
              if ($5.tag === "Right") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.$Expr("ExprConstructor", $1._1, $3._1, $4._1, $5._1)); }
            }
          }
          $runtime.fail();
        }
        if ($2._1 === "Accessor") {
          const $3 = getField(decodeExpr(decAnn))($0._1)("expression");
          if ($3.tag === "Left") { return Data$dEither.$Either("Left", $3._1); }
          if ($3.tag === "Right") {
            const $4 = getField(decodeString)($0._1)("fieldName");
            if ($4.tag === "Left") { return Data$dEither.$Either("Left", $4._1); }
            if ($4.tag === "Right") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.$Expr("ExprAccessor", $1._1, $3._1, $4._1)); }
          }
          $runtime.fail();
        }
        if ($2._1 === "ObjectUpdate") {
          const $3 = getField(decodeExpr(decAnn))($0._1)("expression");
          if ($3.tag === "Left") { return Data$dEither.$Either("Left", $3._1); }
          if ($3.tag === "Right") {
            const $4 = getField(decodeRecord(decodeExpr(decAnn)))($0._1)("updates");
            if ($4.tag === "Left") { return Data$dEither.$Either("Left", $4._1); }
            if ($4.tag === "Right") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.$Expr("ExprUpdate", $1._1, $3._1, $4._1)); }
          }
          $runtime.fail();
        }
        if ($2._1 === "Abs") {
          const $3 = getField(decodeString)($0._1)("argument");
          if ($3.tag === "Left") { return Data$dEither.$Either("Left", $3._1); }
          if ($3.tag === "Right") {
            const $4 = getField(decodeExpr(decAnn))($0._1)("body");
            if ($4.tag === "Left") { return Data$dEither.$Either("Left", $4._1); }
            if ($4.tag === "Right") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.$Expr("ExprAbs", $1._1, $3._1, $4._1)); }
          }
          $runtime.fail();
        }
        if ($2._1 === "App") {
          const $3 = getField(decodeExpr(decAnn))($0._1)("abstraction");
          if ($3.tag === "Left") { return Data$dEither.$Either("Left", $3._1); }
          if ($3.tag === "Right") {
            const $4 = getField(decodeExpr(decAnn))($0._1)("argument");
            if ($4.tag === "Left") { return Data$dEither.$Either("Left", $4._1); }
            if ($4.tag === "Right") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.$Expr("ExprApp", $1._1, $3._1, $4._1)); }
          }
          $runtime.fail();
        }
        if ($2._1 === "Case") {
          const $3 = getField(decodeArray(decodeExpr(decAnn)))($0._1)("caseExpressions");
          if ($3.tag === "Left") { return Data$dEither.$Either("Left", $3._1); }
          if ($3.tag === "Right") {
            const $4 = getField(decodeArray(decodeCaseAlternative(decAnn)))($0._1)("caseAlternatives");
            if ($4.tag === "Left") { return Data$dEither.$Either("Left", $4._1); }
            if ($4.tag === "Right") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.$Expr("ExprCase", $1._1, $3._1, $4._1)); }
          }
          $runtime.fail();
        }
        if ($2._1 === "Let") {
          const $3 = getField(decodeArray(decodeBind(decAnn)))($0._1)("binds");
          if ($3.tag === "Left") { return Data$dEither.$Either("Left", $3._1); }
          if ($3.tag === "Right") {
            const $4 = getField(decodeExpr(decAnn))($0._1)("expression");
            if ($4.tag === "Left") { return Data$dEither.$Either("Left", $4._1); }
            if ($4.tag === "Right") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.$Expr("ExprLet", $1._1, $3._1, $4._1)); }
          }
          $runtime.fail();
        }
        return Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "Expr"));
      }
    }
  }
  $runtime.fail();
};
const decodeCaseAlternative = decAnn => json => {
  const $0 = decodeJObject(json);
  if ($0.tag === "Left") { return Data$dEither.$Either("Left", $0._1); }
  if ($0.tag === "Right") {
    const $1 = getField(decodeArray(decodeBinder(decAnn)))($0._1)("binders");
    if ($1.tag === "Left") { return Data$dEither.$Either("Left", $1._1); }
    if ($1.tag === "Right") {
      const $2 = getField(decodeBoolean)($0._1)("isGuarded");
      if ($2.tag === "Left") { return Data$dEither.$Either("Left", $2._1); }
      if ($2.tag === "Right") {
        if ($2._1) {
          const $3 = getField(decodeArray(decodeGuard(decAnn)))($0._1)("expressions");
          if ($3.tag === "Left") { return Data$dEither.$Either("Left", $3._1); }
          if ($3.tag === "Right") {
            return Data$dEither.$Either(
              "Right",
              PureScript$dBackend$dOptimizer$dCoreFn.$CaseAlternative($1._1, PureScript$dBackend$dOptimizer$dCoreFn.$CaseGuard("Guarded", $3._1))
            );
          }
          $runtime.fail();
        }
        const $3 = getField(decodeExpr(decAnn))($0._1)("expression");
        if ($3.tag === "Left") { return Data$dEither.$Either("Left", $3._1); }
        if ($3.tag === "Right") {
          return Data$dEither.$Either(
            "Right",
            PureScript$dBackend$dOptimizer$dCoreFn.$CaseAlternative($1._1, PureScript$dBackend$dOptimizer$dCoreFn.$CaseGuard("Unconditional", $3._1))
          );
        }
      }
    }
  }
  $runtime.fail();
};
const decodeBinding = decAnn => obj => {
  const $0 = getField(decAnn)(obj)("annotation");
  if ($0.tag === "Left") { return Data$dEither.$Either("Left", $0._1); }
  if ($0.tag === "Right") {
    const $1 = getField(decodeString)(obj)("identifier");
    if ($1.tag === "Left") { return Data$dEither.$Either("Left", $1._1); }
    if ($1.tag === "Right") {
      const $2 = getField(decodeExpr(decAnn))(obj)("expression");
      if ($2.tag === "Left") { return Data$dEither.$Either("Left", $2._1); }
      if ($2.tag === "Right") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.$Binding($0._1, $1._1, $2._1)); }
    }
  }
  $runtime.fail();
};
const decodeBind = decAnn => json => {
  const $0 = decodeJObject(json);
  if ($0.tag === "Left") { return Data$dEither.$Either("Left", $0._1); }
  if ($0.tag === "Right") {
    const $1 = getField(decodeString)($0._1)("bindType");
    if ($1.tag === "Left") { return Data$dEither.$Either("Left", $1._1); }
    if ($1.tag === "Right") {
      if ($1._1 === "NonRec") {
        const $2 = decodeBinding(decAnn)($0._1);
        if ($2.tag === "Left") { return Data$dEither.$Either("Left", $2._1); }
        if ($2.tag === "Right") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.$Bind("NonRec", $2._1)); }
        $runtime.fail();
      }
      if ($1._1 === "Rec") {
        const $2 = getField(decodeArray(a => {
          const $2 = decodeJObject(a);
          if ($2.tag === "Left") { return Data$dEither.$Either("Left", $2._1); }
          if ($2.tag === "Right") { return decodeBinding(decAnn)($2._1); }
          $runtime.fail();
        }))($0._1)("binds");
        if ($2.tag === "Left") { return Data$dEither.$Either("Left", $2._1); }
        if ($2.tag === "Right") { return Data$dEither.$Either("Right", PureScript$dBackend$dOptimizer$dCoreFn.$Bind("Rec", $2._1)); }
        $runtime.fail();
      }
      return Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.$JsonDecodeError("TypeMismatch", "Bind"));
    }
  }
  $runtime.fail();
};
const decodeModule$p = decodeAnn$p => json => {
  const $0 = decodeJObject(json);
  if ($0.tag === "Left") { return Data$dEither.$Either("Left", $0._1); }
  if ($0.tag === "Right") {
    const $1 = getField(decodeModuleName)($0._1)("moduleName");
    if ($1.tag === "Left") { return Data$dEither.$Either("Left", $1._1); }
    if ($1.tag === "Right") {
      const $2 = getField(decodeString)($0._1)("modulePath");
      if ($2.tag === "Left") { return Data$dEither.$Either("Left", $2._1); }
      if ($2.tag === "Right") {
        const $3 = getField(decodeSourceSpan($2._1))($0._1)("sourceSpan");
        if ($3.tag === "Left") { return Data$dEither.$Either("Left", $3._1); }
        if ($3.tag === "Right") {
          const $4 = getField(decodeArray(decodeImport(decodeAnn$p($2._1))))($0._1)("imports");
          if ($4.tag === "Left") { return Data$dEither.$Either("Left", $4._1); }
          if ($4.tag === "Right") {
            const $5 = getField(decodeArray(decodeString))($0._1)("exports");
            if ($5.tag === "Left") { return Data$dEither.$Either("Left", $5._1); }
            if ($5.tag === "Right") {
              const $6 = getField(decodeReExports)($0._1)("reExports");
              if ($6.tag === "Left") { return Data$dEither.$Either("Left", $6._1); }
              if ($6.tag === "Right") {
                const $7 = getField(decodeArray(decodeBind(decodeAnn$p($2._1))))($0._1)("decls");
                if ($7.tag === "Left") { return Data$dEither.$Either("Left", $7._1); }
                if ($7.tag === "Right") {
                  const $8 = getField(decodeArray(decodeString))($0._1)("foreign");
                  if ($8.tag === "Left") { return Data$dEither.$Either("Left", $8._1); }
                  if ($8.tag === "Right") {
                    const $9 = getField(decodeArray(decodeComment))($0._1)("comments");
                    if ($9.tag === "Left") { return Data$dEither.$Either("Left", $9._1); }
                    if ($9.tag === "Right") {
                      return Data$dEither.$Either(
                        "Right",
                        {name: $1._1, path: $2._1, span: $3._1, imports: $4._1, exports: $5._1, reExports: $6._1, decls: $7._1, foreign: $8._1, comments: $9._1}
                      );
                    }
                  }
                }
              }
            }
          }
        }
      }
    }
  }
  $runtime.fail();
};
const decodeModule = /* #__PURE__ */ decodeModule$p(decodeAnn);
export {
  decodeAnn,
  decodeArray,
  decodeBind,
  decodeBinder,
  decodeBinding,
  decodeBoolean,
  decodeCaseAlternative,
  decodeCodePoint,
  decodeComment,
  decodeConstructorType,
  decodeExpr,
  decodeGuard,
  decodeImport,
  decodeInt,
  decodeJArray,
  decodeJObject,
  decodeLiteral,
  decodeMeta,
  decodeModule,
  decodeModule$p,
  decodeModuleName,
  decodeNumber,
  decodeQualified,
  decodeReExports,
  decodeRecord,
  decodeSourcePos,
  decodeSourceSpan,
  decodeString,
  decodeStringLiteral,
  getField,
  getFieldOptional$p,
  traverse
};
