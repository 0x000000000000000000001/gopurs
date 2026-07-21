import * as $runtime from "../runtime.js";
import * as Data$dFoldable from "../Data.Foldable/index.js";
import * as Data$dVoid from "../Data.Void/index.js";
import * as PureScript$dCST$dRange$dTokenList from "../PureScript.CST.Range.TokenList/index.js";
const foldMap = /* #__PURE__ */ (() => Data$dFoldable.foldableArray.foldMap(PureScript$dCST$dRange$dTokenList.monoidTokenList))();
const foldMap2 = /* #__PURE__ */ (() => Data$dFoldable.foldableArray.foldMap(PureScript$dCST$dRange$dTokenList.monoidTokenList))();
const tokensOfVoid = {tokensOf: Data$dVoid.absurd};
const tokensOfRecoveredError = {
  tokensOf: v => {
    const len = v.tokens.length;
    if (len === 0) { return PureScript$dCST$dRange$dTokenList.TokenEmpty; }
    return PureScript$dCST$dRange$dTokenList.$TokenList("TokenArray", 0, len - 1 | 0, v.tokens);
  }
};
const tokensOfQualifiedName = {tokensOf: v => PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v.token, PureScript$dCST$dRange$dTokenList.TokenEmpty)};
const tokensOfName = {tokensOf: v => PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v.token, PureScript$dCST$dRange$dTokenList.TokenEmpty)};
const tokensOf = dict => dict.tokensOf;
const tokensOfArray = dictTokensOf => ({tokensOf: foldMap(a => PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v => dictTokensOf.tokensOf(a)))});
const tokensOfFixityOp = {
  tokensOf: v => {
    if (v.tag === "FixityValue") {
      return PureScript$dCST$dRange$dTokenList.$TokenList(
        "TokenAppend",
        PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty),
        PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenAppend",
          PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._2, PureScript$dCST$dRange$dTokenList.TokenEmpty),
          PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._3.token, PureScript$dCST$dRange$dTokenList.TokenEmpty)
        )
      );
    }
    if (v.tag === "FixityType") {
      return PureScript$dCST$dRange$dTokenList.$TokenList(
        "TokenCons",
        v._1,
        PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenAppend",
          PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._2.token, PureScript$dCST$dRange$dTokenList.TokenEmpty),
          PureScript$dCST$dRange$dTokenList.$TokenList(
            "TokenAppend",
            PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._3, PureScript$dCST$dRange$dTokenList.TokenEmpty),
            PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._4.token, PureScript$dCST$dRange$dTokenList.TokenEmpty)
          )
        )
      );
    }
    $runtime.fail();
  }
};
const tokensOfLabeled = dictTokensOf => dictTokensOf1 => (
  {
    tokensOf: v => {
      const $0 = dictTokensOf.tokensOf(v.label);
      const $1 = dictTokensOf1.tokensOf(v.value);
      const $2 = $1.tag === "TokenEmpty"
        ? PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v.separator, PureScript$dCST$dRange$dTokenList.TokenEmpty)
        : PureScript$dCST$dRange$dTokenList.$TokenList(
            "TokenAppend",
            PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v.separator, PureScript$dCST$dRange$dTokenList.TokenEmpty),
            $1
          );
      if ($2.tag === "TokenEmpty") { return $0; }
      if ($0.tag === "TokenEmpty") { return $2; }
      return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $0, $2);
    }
  }
);
const tokensOfMaybe = dictTokensOf => (
  {
    tokensOf: v1 => {
      if (v1.tag === "Nothing") { return PureScript$dCST$dRange$dTokenList.TokenEmpty; }
      if (v1.tag === "Just") { return dictTokensOf.tokensOf(v1._1); }
      $runtime.fail();
    }
  }
);
const tokensOfNonEmptyArray = dictTokensOf => ({tokensOf: foldMap2(a => PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v => dictTokensOf.tokensOf(a)))});
const tokensOf4 = /* #__PURE__ */ foldMap2(a => PureScript$dCST$dRange$dTokenList.$TokenList(
  "TokenDefer",
  v => PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", a.token, PureScript$dCST$dRange$dTokenList.TokenEmpty)
));
const tokensOfClassFundep = {
  tokensOf: v => {
    if (v.tag === "FundepDetermined") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1, tokensOf4(v._2)); }
    if (v.tag === "FundepDetermines") {
      const $0 = tokensOf4(v._1);
      const $1 = tokensOf4(v._3);
      const $2 = $1.tag === "TokenEmpty"
        ? PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._2, PureScript$dCST$dRange$dTokenList.TokenEmpty)
        : PureScript$dCST$dRange$dTokenList.$TokenList(
            "TokenAppend",
            PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._2, PureScript$dCST$dRange$dTokenList.TokenEmpty),
            $1
          );
      if ($2.tag === "TokenEmpty") { return $0; }
      if ($0.tag === "TokenEmpty") { return $2; }
      return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $0, $2);
    }
    $runtime.fail();
  }
};
const tokensOfPrefixed = dictTokensOf => (
  {
    tokensOf: v => {
      if (v.prefix.tag === "Just") {
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenCons",
          v.prefix._1,
          PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v1 => dictTokensOf.tokensOf(v.value))
        );
      }
      if (v.prefix.tag === "Nothing") { return dictTokensOf.tokensOf(v.value); }
      $runtime.fail();
    }
  }
);
const tokensOfPrefixed1 = {
  tokensOf: v => {
    if (v.prefix.tag === "Just") {
      return PureScript$dCST$dRange$dTokenList.$TokenList(
        "TokenCons",
        v.prefix._1,
        PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenDefer",
          v1 => PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v.value.token, PureScript$dCST$dRange$dTokenList.TokenEmpty)
        )
      );
    }
    if (v.prefix.tag === "Nothing") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v.value.token, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
    $runtime.fail();
  }
};
const tokensOfRecordLabeled = dictTokensOf => (
  {
    tokensOf: v => {
      if (v.tag === "RecordPun") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
      if (v.tag === "RecordField") {
        const $0 = v._3;
        const $1 = v._2;
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenAppend",
          PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty),
          PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v1 => PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $1, dictTokensOf.tokensOf($0)))
        );
      }
      $runtime.fail();
    }
  }
);
const tokensOfSeparated = dictTokensOf => (
  {
    tokensOf: v => {
      const $0 = v.tail;
      const $1 = dictTokensOf.tokensOf(v.head);
      const $2 = PureScript$dCST$dRange$dTokenList.$TokenList(
        "TokenDefer",
        v1 => foldMap(v2 => {
          const $2 = v2._2;
          return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v2._1, PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v3 => dictTokensOf.tokensOf($2)));
        })($0)
      );
      if ($1.tag === "TokenEmpty") { return $2; }
      return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $1, $2);
    }
  }
);
const tokensOfSeparated1 = /* #__PURE__ */ tokensOfSeparated(tokensOfName);
const tokensOf6 = /* #__PURE__ */ (() => tokensOfSeparated(tokensOfClassFundep).tokensOf)();
const tokensOfTuple = dictTokensOf => dictTokensOf1 => (
  {
    tokensOf: v => {
      const $0 = dictTokensOf.tokensOf(v._1);
      const $1 = dictTokensOf1.tokensOf(v._2);
      if ($1.tag === "TokenEmpty") { return $0; }
      if ($0.tag === "TokenEmpty") { return $1; }
      return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $0, $1);
    }
  }
);
const tokensOfWrapped = dictTokensOf => (
  {
    tokensOf: v => {
      const $0 = v.value;
      return PureScript$dCST$dRange$dTokenList.$TokenList("TokenWrap", v.open, PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v1 => dictTokensOf.tokensOf($0)), v.close);
    }
  }
);
const tokensOf7 = v => {
  const $0 = v.value;
  return PureScript$dCST$dRange$dTokenList.$TokenList(
    "TokenWrap",
    v.open,
    PureScript$dCST$dRange$dTokenList.$TokenList(
      "TokenDefer",
      v1 => {
        if ($0.tag === "Nothing") { return PureScript$dCST$dRange$dTokenList.TokenEmpty; }
        if ($0.tag === "Just") { return tokensOfSeparated1.tokensOf($0._1); }
        $runtime.fail();
      }
    ),
    v.close
  );
};
const tokensOfDataMembers = {
  tokensOf: v => {
    if (v.tag === "DataAll") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
    if (v.tag === "DataEnumerated") { return tokensOf7(v._1); }
    $runtime.fail();
  }
};
const tokensOfExport = dictTokensOf => (
  {
    tokensOf: v => {
      if (v.tag === "ExportValue") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
      if (v.tag === "ExportOp") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
      if (v.tag === "ExportType") {
        const $0 = PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty);
        const $1 = (() => {
          if (v._2.tag === "Nothing") { return PureScript$dCST$dRange$dTokenList.TokenEmpty; }
          if (v._2.tag === "Just") {
            if (v._2._1.tag === "DataAll") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._2._1._1, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
            if (v._2._1.tag === "DataEnumerated") { return tokensOf7(v._2._1._1); }
          }
          $runtime.fail();
        })();
        if ($1.tag === "TokenEmpty") { return $0; }
        return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $0, $1);
      }
      if (v.tag === "ExportTypeOp") {
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenCons",
          v._1,
          PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._2.token, PureScript$dCST$dRange$dTokenList.TokenEmpty)
        );
      }
      if (v.tag === "ExportClass") {
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenCons",
          v._1,
          PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._2.token, PureScript$dCST$dRange$dTokenList.TokenEmpty)
        );
      }
      if (v.tag === "ExportModule") {
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenCons",
          v._1,
          PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._2.token, PureScript$dCST$dRange$dTokenList.TokenEmpty)
        );
      }
      if (v.tag === "ExportError") { return dictTokensOf.tokensOf(v._1); }
      $runtime.fail();
    }
  }
);
const tokensOfImport = dictTokensOf => (
  {
    tokensOf: v => {
      if (v.tag === "ImportValue") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
      if (v.tag === "ImportOp") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
      if (v.tag === "ImportType") {
        const $0 = PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty);
        const $1 = (() => {
          if (v._2.tag === "Nothing") { return PureScript$dCST$dRange$dTokenList.TokenEmpty; }
          if (v._2.tag === "Just") {
            if (v._2._1.tag === "DataAll") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._2._1._1, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
            if (v._2._1.tag === "DataEnumerated") { return tokensOf7(v._2._1._1); }
          }
          $runtime.fail();
        })();
        if ($1.tag === "TokenEmpty") { return $0; }
        return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $0, $1);
      }
      if (v.tag === "ImportTypeOp") {
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenCons",
          v._1,
          PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._2.token, PureScript$dCST$dRange$dTokenList.TokenEmpty)
        );
      }
      if (v.tag === "ImportClass") {
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenCons",
          v._1,
          PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._2.token, PureScript$dCST$dRange$dTokenList.TokenEmpty)
        );
      }
      if (v.tag === "ImportError") { return dictTokensOf.tokensOf(v._1); }
      $runtime.fail();
    }
  }
);
const tokensOfImportDecl = dictTokensOf => {
  const $0 = tokensOfSeparated(tokensOfImport(dictTokensOf));
  return {
    tokensOf: v => {
      const $1 = v.module;
      const $2 = v.names;
      const $3 = v.qualified;
      return PureScript$dCST$dRange$dTokenList.$TokenList(
        "TokenCons",
        v.keyword,
        PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenDefer",
          v1 => {
            const $4 = (() => {
              if ($2.tag === "Nothing") { return PureScript$dCST$dRange$dTokenList.TokenEmpty; }
              if ($2.tag === "Just") {
                const $4 = $2._1._2;
                const $5 = (() => {
                  if ($2._1._1.tag === "Nothing") { return PureScript$dCST$dRange$dTokenList.TokenEmpty; }
                  if ($2._1._1.tag === "Just") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $2._1._1._1, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
                  $runtime.fail();
                })();
                const $6 = PureScript$dCST$dRange$dTokenList.$TokenList(
                  "TokenDefer",
                  v3 => {
                    const $6 = $4.value;
                    return PureScript$dCST$dRange$dTokenList.$TokenList(
                      "TokenWrap",
                      $4.open,
                      PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v1$1 => $0.tokensOf($6)),
                      $4.close
                    );
                  }
                );
                if ($5.tag === "TokenEmpty") { return $6; }
                return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $5, $6);
              }
              $runtime.fail();
            })();
            const $5 = (() => {
              if ($3.tag === "Nothing") { return PureScript$dCST$dRange$dTokenList.TokenEmpty; }
              if ($3.tag === "Just") {
                return PureScript$dCST$dRange$dTokenList.$TokenList(
                  "TokenAppend",
                  PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $3._1._1, PureScript$dCST$dRange$dTokenList.TokenEmpty),
                  PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $3._1._2.token, PureScript$dCST$dRange$dTokenList.TokenEmpty)
                );
              }
              $runtime.fail();
            })();
            const $6 = (() => {
              if ($5.tag === "TokenEmpty") { return $4; }
              if ($4.tag === "TokenEmpty") { return $5; }
              return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $4, $5);
            })();
            if ($6.tag === "TokenEmpty") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
            return PureScript$dCST$dRange$dTokenList.$TokenList(
              "TokenAppend",
              PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty),
              $6
            );
          }
        )
      );
    }
  };
};
const tokensOfOneOrDelimited = dictTokensOf => {
  const $0 = tokensOfSeparated(dictTokensOf);
  return {
    tokensOf: v => {
      if (v.tag === "One") { return dictTokensOf.tokensOf(v._1); }
      if (v.tag === "Many") {
        const $1 = v._1.value;
        return PureScript$dCST$dRange$dTokenList.$TokenList("TokenWrap", v._1.open, PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v1 => $0.tokensOf($1)), v._1.close);
      }
      $runtime.fail();
    }
  };
};
const tokensOfTypeVarBinding = dictTokensOf => dictTokensOf1 => (
  {
    tokensOf: v => {
      if (v.tag === "TypeVarKinded") {
        const $0 = tokensOfType(dictTokensOf1);
        const $1 = v._1.value;
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenWrap",
          v._1.open,
          PureScript$dCST$dRange$dTokenList.$TokenList(
            "TokenDefer",
            v1 => {
              const $2 = dictTokensOf.tokensOf($1.label);
              const $3 = $0.tokensOf($1.value);
              const $4 = $3.tag === "TokenEmpty"
                ? PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $1.separator, PureScript$dCST$dRange$dTokenList.TokenEmpty)
                : PureScript$dCST$dRange$dTokenList.$TokenList(
                    "TokenAppend",
                    PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $1.separator, PureScript$dCST$dRange$dTokenList.TokenEmpty),
                    $3
                  );
              if ($4.tag === "TokenEmpty") { return $2; }
              if ($2.tag === "TokenEmpty") { return $4; }
              return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $2, $4);
            }
          ),
          v._1.close
        );
      }
      if (v.tag === "TypeVarName") { return dictTokensOf.tokensOf(v._1); }
      $runtime.fail();
    }
  }
);
const tokensOfType = dictTokensOf => (
  {
    tokensOf: v => {
      if (v.tag === "TypeVar") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
      if (v.tag === "TypeConstructor") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
      if (v.tag === "TypeWildcard") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
      if (v.tag === "TypeHole") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
      if (v.tag === "TypeString") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
      if (v.tag === "TypeInt") {
        const $0 = (() => {
          if (v._1.tag === "Nothing") { return PureScript$dCST$dRange$dTokenList.TokenEmpty; }
          if (v._1.tag === "Just") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1._1, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
          $runtime.fail();
        })();
        if ($0.tag === "TokenEmpty") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._2, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenAppend",
          $0,
          PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._2, PureScript$dCST$dRange$dTokenList.TokenEmpty)
        );
      }
      if (v.tag === "TypeRow") {
        const $0 = tokensOfRow(dictTokensOf);
        const $1 = v._1.value;
        return PureScript$dCST$dRange$dTokenList.$TokenList("TokenWrap", v._1.open, PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v1 => $0.tokensOf($1)), v._1.close);
      }
      if (v.tag === "TypeRecord") {
        const $0 = tokensOfRow(dictTokensOf);
        const $1 = v._1.value;
        return PureScript$dCST$dRange$dTokenList.$TokenList("TokenWrap", v._1.open, PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v1 => $0.tokensOf($1)), v._1.close);
      }
      if (v.tag === "TypeForall") {
        const $0 = v._3;
        const $1 = v._4;
        const $2 = v._2;
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenCons",
          v._1,
          PureScript$dCST$dRange$dTokenList.$TokenList(
            "TokenDefer",
            v1 => {
              const $3 = tokensOfTypeVarBinding(tokensOfPrefixed1)(dictTokensOf);
              const $4 = foldMap2(a => PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v$1 => $3.tokensOf(a)))($2);
              const $5 = tokensOfType(dictTokensOf).tokensOf($1);
              const $6 = $5.tag === "TokenEmpty"
                ? PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $0, PureScript$dCST$dRange$dTokenList.TokenEmpty)
                : PureScript$dCST$dRange$dTokenList.$TokenList(
                    "TokenAppend",
                    PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $0, PureScript$dCST$dRange$dTokenList.TokenEmpty),
                    $5
                  );
              if ($6.tag === "TokenEmpty") { return $4; }
              if ($4.tag === "TokenEmpty") { return $6; }
              return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $4, $6);
            }
          )
        );
      }
      if (v.tag === "TypeKinded") {
        const $0 = v._2;
        const $1 = v._3;
        const $2 = tokensOfType(dictTokensOf).tokensOf(v._1);
        const $3 = PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenDefer",
          v1 => {
            const $3 = tokensOfType(dictTokensOf).tokensOf($1);
            if ($3.tag === "TokenEmpty") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $0, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
            return PureScript$dCST$dRange$dTokenList.$TokenList(
              "TokenAppend",
              PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $0, PureScript$dCST$dRange$dTokenList.TokenEmpty),
              $3
            );
          }
        );
        if ($2.tag === "TokenEmpty") { return $3; }
        return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $2, $3);
      }
      if (v.tag === "TypeApp") {
        const $0 = v._2;
        const $1 = tokensOfType(dictTokensOf).tokensOf(v._1);
        const $2 = PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenDefer",
          v1 => {
            const $2 = tokensOfType(dictTokensOf);
            return foldMap2(a => PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v$1 => $2.tokensOf(a)))($0);
          }
        );
        if ($1.tag === "TokenEmpty") { return $2; }
        return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $1, $2);
      }
      if (v.tag === "TypeOp") {
        const $0 = v._2;
        const $1 = tokensOfType(dictTokensOf).tokensOf(v._1);
        const $2 = PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenDefer",
          v1 => foldMap2(v2 => {
            const $2 = PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v2._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty);
            const $3 = tokensOfType(dictTokensOf).tokensOf(v2._2);
            if ($3.tag === "TokenEmpty") { return $2; }
            return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $2, $3);
          })($0)
        );
        if ($1.tag === "TokenEmpty") { return $2; }
        return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $1, $2);
      }
      if (v.tag === "TypeOpName") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
      if (v.tag === "TypeArrow") {
        const $0 = v._2;
        const $1 = v._3;
        const $2 = tokensOfType(dictTokensOf).tokensOf(v._1);
        const $3 = PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenDefer",
          v1 => {
            const $3 = tokensOfType(dictTokensOf).tokensOf($1);
            if ($3.tag === "TokenEmpty") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $0, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
            return PureScript$dCST$dRange$dTokenList.$TokenList(
              "TokenAppend",
              PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $0, PureScript$dCST$dRange$dTokenList.TokenEmpty),
              $3
            );
          }
        );
        if ($2.tag === "TokenEmpty") { return $3; }
        return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $2, $3);
      }
      if (v.tag === "TypeArrowName") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
      if (v.tag === "TypeConstrained") {
        const $0 = v._2;
        const $1 = v._3;
        const $2 = tokensOfType(dictTokensOf).tokensOf(v._1);
        const $3 = PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenDefer",
          v1 => {
            const $3 = tokensOfType(dictTokensOf).tokensOf($1);
            if ($3.tag === "TokenEmpty") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $0, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
            return PureScript$dCST$dRange$dTokenList.$TokenList(
              "TokenAppend",
              PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $0, PureScript$dCST$dRange$dTokenList.TokenEmpty),
              $3
            );
          }
        );
        if ($2.tag === "TokenEmpty") { return $3; }
        return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $2, $3);
      }
      if (v.tag === "TypeParens") {
        const $0 = tokensOfType(dictTokensOf);
        const $1 = v._1.value;
        return PureScript$dCST$dRange$dTokenList.$TokenList("TokenWrap", v._1.open, PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v1 => $0.tokensOf($1)), v._1.close);
      }
      if (v.tag === "TypeError") { return dictTokensOf.tokensOf(v._1); }
      $runtime.fail();
    }
  }
);
const tokensOfRow = dictTokensOf => (
  {
    tokensOf: v => {
      const $0 = (() => {
        if (v.labels.tag === "Nothing") { return PureScript$dCST$dRange$dTokenList.TokenEmpty; }
        if (v.labels.tag === "Just") {
          return tokensOfSeparated((() => {
            const $0 = tokensOfType(dictTokensOf);
            return {
              tokensOf: v$1 => {
                const $1 = PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v$1.label.token, PureScript$dCST$dRange$dTokenList.TokenEmpty);
                const $2 = $0.tokensOf(v$1.value);
                const $3 = $2.tag === "TokenEmpty"
                  ? PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v$1.separator, PureScript$dCST$dRange$dTokenList.TokenEmpty)
                  : PureScript$dCST$dRange$dTokenList.$TokenList(
                      "TokenAppend",
                      PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v$1.separator, PureScript$dCST$dRange$dTokenList.TokenEmpty),
                      $2
                    );
                if ($3.tag === "TokenEmpty") { return $1; }
                return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $1, $3);
              }
            };
          })()).tokensOf(v.labels._1);
        }
        $runtime.fail();
      })();
      const $1 = (() => {
        if (v.tail.tag === "Nothing") { return PureScript$dCST$dRange$dTokenList.TokenEmpty; }
        if (v.tail.tag === "Just") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v.tail._1._1, tokensOfType(dictTokensOf).tokensOf(v.tail._1._2)); }
        $runtime.fail();
      })();
      if ($1.tag === "TokenEmpty") { return $0; }
      if ($0.tag === "TokenEmpty") { return $1; }
      return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $0, $1);
    }
  }
);
const tokensOfAppSpine = dictTokensOf => dictTokensOf1 => (
  {
    tokensOf: v => {
      if (v.tag === "AppType") {
        const $0 = v._2;
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenCons",
          v._1,
          PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v1 => tokensOfType(dictTokensOf).tokensOf($0))
        );
      }
      if (v.tag === "AppTerm") { return dictTokensOf1.tokensOf(v._1); }
      $runtime.fail();
    }
  }
);
const tokensOfBinder = dictTokensOf => (
  {
    tokensOf: v => {
      if (v.tag === "BinderWildcard") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
      if (v.tag === "BinderVar") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
      if (v.tag === "BinderNamed") {
        const $0 = v._3;
        const $1 = v._2;
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenAppend",
          PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty),
          PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v1 => PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $1, tokensOfBinder(dictTokensOf).tokensOf($0)))
        );
      }
      if (v.tag === "BinderConstructor") {
        const $0 = v._2;
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenAppend",
          PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty),
          PureScript$dCST$dRange$dTokenList.$TokenList(
            "TokenDefer",
            v1 => {
              const $1 = tokensOfBinder(dictTokensOf);
              return foldMap(a => PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v$1 => $1.tokensOf(a)))($0);
            }
          )
        );
      }
      if (v.tag === "BinderBoolean") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
      if (v.tag === "BinderChar") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
      if (v.tag === "BinderString") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
      if (v.tag === "BinderInt") {
        const $0 = (() => {
          if (v._1.tag === "Nothing") { return PureScript$dCST$dRange$dTokenList.TokenEmpty; }
          if (v._1.tag === "Just") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1._1, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
          $runtime.fail();
        })();
        if ($0.tag === "TokenEmpty") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._2, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenAppend",
          $0,
          PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._2, PureScript$dCST$dRange$dTokenList.TokenEmpty)
        );
      }
      if (v.tag === "BinderNumber") {
        const $0 = (() => {
          if (v._1.tag === "Nothing") { return PureScript$dCST$dRange$dTokenList.TokenEmpty; }
          if (v._1.tag === "Just") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1._1, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
          $runtime.fail();
        })();
        if ($0.tag === "TokenEmpty") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._2, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenAppend",
          $0,
          PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._2, PureScript$dCST$dRange$dTokenList.TokenEmpty)
        );
      }
      if (v.tag === "BinderArray") {
        const $0 = tokensOfSeparated(tokensOfBinder(dictTokensOf));
        const $1 = v._1.value;
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenWrap",
          v._1.open,
          PureScript$dCST$dRange$dTokenList.$TokenList(
            "TokenDefer",
            v1 => {
              if ($1.tag === "Nothing") { return PureScript$dCST$dRange$dTokenList.TokenEmpty; }
              if ($1.tag === "Just") { return $0.tokensOf($1._1); }
              $runtime.fail();
            }
          ),
          v._1.close
        );
      }
      if (v.tag === "BinderRecord") {
        const $0 = tokensOfSeparated((() => {
          const $0 = tokensOfBinder(dictTokensOf);
          return {
            tokensOf: v$1 => {
              if (v$1.tag === "RecordPun") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v$1._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
              if (v$1.tag === "RecordField") {
                const $1 = v$1._3;
                const $2 = v$1._2;
                return PureScript$dCST$dRange$dTokenList.$TokenList(
                  "TokenAppend",
                  PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v$1._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty),
                  PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v1 => PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $2, $0.tokensOf($1)))
                );
              }
              $runtime.fail();
            }
          };
        })());
        const $1 = v._1.value;
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenWrap",
          v._1.open,
          PureScript$dCST$dRange$dTokenList.$TokenList(
            "TokenDefer",
            v1 => {
              if ($1.tag === "Nothing") { return PureScript$dCST$dRange$dTokenList.TokenEmpty; }
              if ($1.tag === "Just") { return $0.tokensOf($1._1); }
              $runtime.fail();
            }
          ),
          v._1.close
        );
      }
      if (v.tag === "BinderParens") {
        const $0 = tokensOfBinder(dictTokensOf);
        const $1 = v._1.value;
        return PureScript$dCST$dRange$dTokenList.$TokenList("TokenWrap", v._1.open, PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v1 => $0.tokensOf($1)), v._1.close);
      }
      if (v.tag === "BinderTyped") {
        const $0 = v._2;
        const $1 = v._3;
        const $2 = tokensOfBinder(dictTokensOf).tokensOf(v._1);
        const $3 = PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenDefer",
          v1 => PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $0, tokensOfType(dictTokensOf).tokensOf($1))
        );
        if ($2.tag === "TokenEmpty") { return $3; }
        return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $2, $3);
      }
      if (v.tag === "BinderOp") {
        const $0 = v._2;
        const $1 = tokensOfBinder(dictTokensOf).tokensOf(v._1);
        const $2 = PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenDefer",
          v1 => {
            const $2 = tokensOfBinder(dictTokensOf);
            return foldMap2(a => PureScript$dCST$dRange$dTokenList.$TokenList(
              "TokenDefer",
              v$1 => {
                const $3 = PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", a._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty);
                const $4 = $2.tokensOf(a._2);
                if ($4.tag === "TokenEmpty") { return $3; }
                return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $3, $4);
              }
            ))($0);
          }
        );
        if ($1.tag === "TokenEmpty") { return $2; }
        return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $1, $2);
      }
      if (v.tag === "BinderError") { return dictTokensOf.tokensOf(v._1); }
      $runtime.fail();
    }
  }
);
const tokensOfDataCtor = dictTokensOf => {
  const $0 = tokensOfType(dictTokensOf);
  const tokensOf9 = foldMap(a => PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v => $0.tokensOf(a)));
  return {
    tokensOf: v => {
      const $1 = PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v.name.token, PureScript$dCST$dRange$dTokenList.TokenEmpty);
      const $2 = tokensOf9(v.fields);
      if ($2.tag === "TokenEmpty") { return $1; }
      return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $1, $2);
    }
  };
};
const tokensOfForeign = dictTokensOf => {
  const $0 = tokensOfType(dictTokensOf);
  const tokensOf9 = v => {
    const $1 = PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v.label.token, PureScript$dCST$dRange$dTokenList.TokenEmpty);
    const $2 = $0.tokensOf(v.value);
    const $3 = $2.tag === "TokenEmpty"
      ? PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v.separator, PureScript$dCST$dRange$dTokenList.TokenEmpty)
      : PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenAppend",
          PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v.separator, PureScript$dCST$dRange$dTokenList.TokenEmpty),
          $2
        );
    if ($3.tag === "TokenEmpty") { return $1; }
    return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $1, $3);
  };
  return {
    tokensOf: v => {
      if (v.tag === "ForeignValue") { return tokensOf9(v._1); }
      if (v.tag === "ForeignData") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1, tokensOf9(v._2)); }
      if (v.tag === "ForeignKind") {
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenCons",
          v._1,
          PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._2.token, PureScript$dCST$dRange$dTokenList.TokenEmpty)
        );
      }
      $runtime.fail();
    }
  };
};
const tokensOfWhere = dictTokensOf => (
  {
    tokensOf: v => {
      const $0 = v.bindings;
      const $1 = tokensOfExpr(dictTokensOf).tokensOf(v.expr);
      const $2 = PureScript$dCST$dRange$dTokenList.$TokenList(
        "TokenDefer",
        v1 => {
          if ($0.tag === "Nothing") { return PureScript$dCST$dRange$dTokenList.TokenEmpty; }
          if ($0.tag === "Just") {
            return PureScript$dCST$dRange$dTokenList.$TokenList(
              "TokenCons",
              $0._1._1,
              (() => {
                const $2 = tokensOfLetBinding(dictTokensOf);
                return foldMap2(a => PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v$1 => $2.tokensOf(a)))($0._1._2);
              })()
            );
          }
          $runtime.fail();
        }
      );
      if ($1.tag === "TokenEmpty") { return $2; }
      return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $1, $2);
    }
  }
);
const tokensOfRecordUpdate = dictTokensOf => (
  {
    tokensOf: v => {
      if (v.tag === "RecordUpdateLeaf") {
        const $0 = PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty);
        const $1 = tokensOfExpr(dictTokensOf).tokensOf(v._3);
        const $2 = $1.tag === "TokenEmpty"
          ? PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._2, PureScript$dCST$dRange$dTokenList.TokenEmpty)
          : PureScript$dCST$dRange$dTokenList.$TokenList(
              "TokenAppend",
              PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._2, PureScript$dCST$dRange$dTokenList.TokenEmpty),
              $1
            );
        if ($2.tag === "TokenEmpty") { return $0; }
        return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $0, $2);
      }
      if (v.tag === "RecordUpdateBranch") {
        const $0 = tokensOfSeparated(tokensOfRecordUpdate(dictTokensOf));
        const $1 = v._2.value;
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenAppend",
          PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty),
          PureScript$dCST$dRange$dTokenList.$TokenList("TokenWrap", v._2.open, PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v1 => $0.tokensOf($1)), v._2.close)
        );
      }
      $runtime.fail();
    }
  }
);
const tokensOfPatternGuard = dictTokensOf => (
  {
    tokensOf: v => {
      const $0 = (() => {
        if (v.binder.tag === "Nothing") { return PureScript$dCST$dRange$dTokenList.TokenEmpty; }
        if (v.binder.tag === "Just") {
          const $0 = tokensOfBinder(dictTokensOf).tokensOf(v.binder._1._1);
          if ($0.tag === "TokenEmpty") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v.binder._1._2, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
          return PureScript$dCST$dRange$dTokenList.$TokenList(
            "TokenAppend",
            $0,
            PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v.binder._1._2, PureScript$dCST$dRange$dTokenList.TokenEmpty)
          );
        }
        $runtime.fail();
      })();
      const $1 = tokensOfExpr(dictTokensOf).tokensOf(v.expr);
      if ($1.tag === "TokenEmpty") { return $0; }
      if ($0.tag === "TokenEmpty") { return $1; }
      return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $0, $1);
    }
  }
);
const tokensOfLetBinding = dictTokensOf => {
  const $0 = tokensOfType(dictTokensOf);
  const tokensOfBinder1 = tokensOfBinder(dictTokensOf);
  return {
    tokensOf: v => {
      if (v.tag === "LetBindingSignature") {
        const $1 = PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.label.token, PureScript$dCST$dRange$dTokenList.TokenEmpty);
        const $2 = $0.tokensOf(v._1.value);
        const $3 = $2.tag === "TokenEmpty"
          ? PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.separator, PureScript$dCST$dRange$dTokenList.TokenEmpty)
          : PureScript$dCST$dRange$dTokenList.$TokenList(
              "TokenAppend",
              PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.separator, PureScript$dCST$dRange$dTokenList.TokenEmpty),
              $2
            );
        if ($3.tag === "TokenEmpty") { return $1; }
        return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $1, $3);
      }
      if (v.tag === "LetBindingName") {
        const $1 = v._1.binders;
        const $2 = v._1.guarded;
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenAppend",
          PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.name.token, PureScript$dCST$dRange$dTokenList.TokenEmpty),
          PureScript$dCST$dRange$dTokenList.$TokenList(
            "TokenDefer",
            v1 => {
              const $3 = foldMap(a => PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v$1 => tokensOfBinder1.tokensOf(a)))($1);
              const $4 = tokensOfGuarded(dictTokensOf).tokensOf($2);
              if ($4.tag === "TokenEmpty") { return $3; }
              if ($3.tag === "TokenEmpty") { return $4; }
              return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $3, $4);
            }
          )
        );
      }
      if (v.tag === "LetBindingPattern") {
        const $1 = v._2;
        const $2 = v._3;
        const $3 = tokensOfBinder1.tokensOf(v._1);
        const $4 = PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenDefer",
          v1 => PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $1, tokensOfWhere(dictTokensOf).tokensOf($2))
        );
        if ($3.tag === "TokenEmpty") { return $4; }
        return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $3, $4);
      }
      if (v.tag === "LetBindingError") { return dictTokensOf.tokensOf(v._1); }
      $runtime.fail();
    }
  };
};
const tokensOfGuardedExpr = dictTokensOf => (
  {
    tokensOf: v => PureScript$dCST$dRange$dTokenList.$TokenList(
      "TokenCons",
      v.bar,
      PureScript$dCST$dRange$dTokenList.$TokenList(
        "TokenDefer",
        v1 => {
          const $0 = tokensOfSeparated(tokensOfPatternGuard(dictTokensOf)).tokensOf(v.patterns);
          const $1 = tokensOfWhere(dictTokensOf).tokensOf(v.where);
          const $2 = $1.tag === "TokenEmpty"
            ? PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v.separator, PureScript$dCST$dRange$dTokenList.TokenEmpty)
            : PureScript$dCST$dRange$dTokenList.$TokenList(
                "TokenAppend",
                PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v.separator, PureScript$dCST$dRange$dTokenList.TokenEmpty),
                $1
              );
          if ($2.tag === "TokenEmpty") { return $0; }
          if ($0.tag === "TokenEmpty") { return $2; }
          return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $0, $2);
        }
      )
    )
  }
);
const tokensOfGuarded = dictTokensOf => (
  {
    tokensOf: v => {
      if (v.tag === "Unconditional") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1, tokensOfWhere(dictTokensOf).tokensOf(v._2)); }
      if (v.tag === "Guarded") {
        const $0 = tokensOfGuardedExpr(dictTokensOf);
        return foldMap2(a => PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v$1 => $0.tokensOf(a)))(v._1);
      }
      $runtime.fail();
    }
  }
);
const tokensOfExpr = dictTokensOf => {
  const tokensOfBinder1 = tokensOfBinder(dictTokensOf);
  const $0 = tokensOfSeparated(tokensOfBinder1);
  return {
    tokensOf: v => {
      if (v.tag === "ExprHole") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
      if (v.tag === "ExprSection") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
      if (v.tag === "ExprIdent") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
      if (v.tag === "ExprConstructor") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
      if (v.tag === "ExprBoolean") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
      if (v.tag === "ExprChar") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
      if (v.tag === "ExprString") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
      if (v.tag === "ExprInt") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
      if (v.tag === "ExprNumber") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
      if (v.tag === "ExprArray") {
        const $1 = tokensOfSeparated(tokensOfExpr(dictTokensOf));
        const $2 = v._1.value;
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenWrap",
          v._1.open,
          PureScript$dCST$dRange$dTokenList.$TokenList(
            "TokenDefer",
            v1 => {
              if ($2.tag === "Nothing") { return PureScript$dCST$dRange$dTokenList.TokenEmpty; }
              if ($2.tag === "Just") { return $1.tokensOf($2._1); }
              $runtime.fail();
            }
          ),
          v._1.close
        );
      }
      if (v.tag === "ExprRecord") {
        const $1 = tokensOfSeparated((() => {
          const $1 = tokensOfExpr(dictTokensOf);
          return {
            tokensOf: v$1 => {
              if (v$1.tag === "RecordPun") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v$1._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
              if (v$1.tag === "RecordField") {
                const $2 = v$1._3;
                const $3 = v$1._2;
                return PureScript$dCST$dRange$dTokenList.$TokenList(
                  "TokenAppend",
                  PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v$1._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty),
                  PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v1 => PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $3, $1.tokensOf($2)))
                );
              }
              $runtime.fail();
            }
          };
        })());
        const $2 = v._1.value;
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenWrap",
          v._1.open,
          PureScript$dCST$dRange$dTokenList.$TokenList(
            "TokenDefer",
            v1 => {
              if ($2.tag === "Nothing") { return PureScript$dCST$dRange$dTokenList.TokenEmpty; }
              if ($2.tag === "Just") { return $1.tokensOf($2._1); }
              $runtime.fail();
            }
          ),
          v._1.close
        );
      }
      if (v.tag === "ExprParens") {
        const $1 = tokensOfExpr(dictTokensOf);
        const $2 = v._1.value;
        return PureScript$dCST$dRange$dTokenList.$TokenList("TokenWrap", v._1.open, PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v1 => $1.tokensOf($2)), v._1.close);
      }
      if (v.tag === "ExprTyped") {
        const $1 = v._2;
        const $2 = v._3;
        const $3 = tokensOfExpr(dictTokensOf).tokensOf(v._1);
        const $4 = PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenDefer",
          v1 => PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $1, tokensOfType(dictTokensOf).tokensOf($2))
        );
        if ($3.tag === "TokenEmpty") { return $4; }
        return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $3, $4);
      }
      if (v.tag === "ExprInfix") {
        const $1 = v._2;
        const $2 = tokensOfExpr(dictTokensOf).tokensOf(v._1);
        const $3 = PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenDefer",
          v1 => {
            const $3 = tokensOfExpr(dictTokensOf);
            const $4 = tokensOfExpr(dictTokensOf);
            return foldMap2(a => PureScript$dCST$dRange$dTokenList.$TokenList(
              "TokenDefer",
              v$1 => {
                const $5 = a._1.value;
                const $6 = PureScript$dCST$dRange$dTokenList.$TokenList(
                  "TokenWrap",
                  a._1.open,
                  PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v1$1 => $3.tokensOf($5)),
                  a._1.close
                );
                const $7 = $4.tokensOf(a._2);
                if ($7.tag === "TokenEmpty") { return $6; }
                return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $6, $7);
              }
            ))($1);
          }
        );
        if ($2.tag === "TokenEmpty") { return $3; }
        return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $2, $3);
      }
      if (v.tag === "ExprOp") {
        const $1 = v._2;
        const $2 = tokensOfExpr(dictTokensOf).tokensOf(v._1);
        const $3 = PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenDefer",
          v1 => {
            const $3 = tokensOfExpr(dictTokensOf);
            return foldMap2(a => PureScript$dCST$dRange$dTokenList.$TokenList(
              "TokenDefer",
              v$1 => {
                const $4 = PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", a._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty);
                const $5 = $3.tokensOf(a._2);
                if ($5.tag === "TokenEmpty") { return $4; }
                return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $4, $5);
              }
            ))($1);
          }
        );
        if ($2.tag === "TokenEmpty") { return $3; }
        return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $2, $3);
      }
      if (v.tag === "ExprOpName") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
      if (v.tag === "ExprNegate") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1, tokensOfExpr(dictTokensOf).tokensOf(v._2)); }
      if (v.tag === "ExprRecordAccessor") {
        const $1 = v._1.dot;
        const $2 = v._1.path;
        const $3 = tokensOfExpr(dictTokensOf).tokensOf(v._1.expr);
        const $4 = PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v1 => PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $1, tokensOfSeparated1.tokensOf($2)));
        if ($3.tag === "TokenEmpty") { return $4; }
        return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $3, $4);
      }
      if (v.tag === "ExprRecordUpdate") {
        const $1 = v._2;
        const $2 = tokensOfExpr(dictTokensOf).tokensOf(v._1);
        const $3 = PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenDefer",
          v1 => {
            const $3 = tokensOfSeparated(tokensOfRecordUpdate(dictTokensOf));
            const $4 = $1.value;
            return PureScript$dCST$dRange$dTokenList.$TokenList("TokenWrap", $1.open, PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v1$1 => $3.tokensOf($4)), $1.close);
          }
        );
        if ($2.tag === "TokenEmpty") { return $3; }
        return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $2, $3);
      }
      if (v.tag === "ExprApp") {
        const $1 = v._2;
        const $2 = tokensOfExpr(dictTokensOf).tokensOf(v._1);
        const $3 = PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenDefer",
          v1 => {
            const $3 = tokensOfAppSpine(dictTokensOf)(tokensOfExpr(dictTokensOf));
            return foldMap2(a => PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v$1 => $3.tokensOf(a)))($1);
          }
        );
        if ($2.tag === "TokenEmpty") { return $3; }
        return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $2, $3);
      }
      if (v.tag === "ExprLambda") {
        const $1 = v._1.arrow;
        const $2 = v._1.binders;
        const $3 = v._1.body;
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenCons",
          v._1.symbol,
          PureScript$dCST$dRange$dTokenList.$TokenList(
            "TokenDefer",
            v1 => {
              const $4 = foldMap2(a => PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v$1 => tokensOfBinder1.tokensOf(a)))($2);
              const $5 = tokensOfExpr(dictTokensOf).tokensOf($3);
              const $6 = $5.tag === "TokenEmpty"
                ? PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $1, PureScript$dCST$dRange$dTokenList.TokenEmpty)
                : PureScript$dCST$dRange$dTokenList.$TokenList(
                    "TokenAppend",
                    PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $1, PureScript$dCST$dRange$dTokenList.TokenEmpty),
                    $5
                  );
              if ($6.tag === "TokenEmpty") { return $4; }
              if ($4.tag === "TokenEmpty") { return $6; }
              return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $4, $6);
            }
          )
        );
      }
      if (v.tag === "ExprIf") {
        const $1 = v._1;
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenCons",
          $1.keyword,
          PureScript$dCST$dRange$dTokenList.$TokenList(
            "TokenDefer",
            v1 => {
              const $2 = tokensOfExpr(dictTokensOf).tokensOf($1.cond);
              const $3 = tokensOfExpr(dictTokensOf).tokensOf($1.true);
              const $4 = tokensOfExpr(dictTokensOf).tokensOf($1.false);
              const $5 = $4.tag === "TokenEmpty"
                ? PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $1.else, PureScript$dCST$dRange$dTokenList.TokenEmpty)
                : PureScript$dCST$dRange$dTokenList.$TokenList(
                    "TokenAppend",
                    PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $1.else, PureScript$dCST$dRange$dTokenList.TokenEmpty),
                    $4
                  );
              const $6 = (() => {
                if ($5.tag === "TokenEmpty") { return $3; }
                if ($3.tag === "TokenEmpty") { return $5; }
                return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $3, $5);
              })();
              const $7 = $6.tag === "TokenEmpty"
                ? PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $1.then, PureScript$dCST$dRange$dTokenList.TokenEmpty)
                : PureScript$dCST$dRange$dTokenList.$TokenList(
                    "TokenAppend",
                    PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $1.then, PureScript$dCST$dRange$dTokenList.TokenEmpty),
                    $6
                  );
              if ($7.tag === "TokenEmpty") { return $2; }
              if ($2.tag === "TokenEmpty") { return $7; }
              return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $2, $7);
            }
          )
        );
      }
      if (v.tag === "ExprCase") {
        const $1 = v._1;
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenCons",
          $1.keyword,
          PureScript$dCST$dRange$dTokenList.$TokenList(
            "TokenDefer",
            v1 => {
              const $2 = tokensOfSeparated(tokensOfExpr(dictTokensOf)).tokensOf($1.head);
              const $3 = tokensOfGuarded(dictTokensOf);
              const $4 = foldMap2(a => PureScript$dCST$dRange$dTokenList.$TokenList(
                "TokenDefer",
                v$1 => {
                  const $4 = $0.tokensOf(a._1);
                  const $5 = $3.tokensOf(a._2);
                  if ($5.tag === "TokenEmpty") { return $4; }
                  if ($4.tag === "TokenEmpty") { return $5; }
                  return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $4, $5);
                }
              ))($1.branches);
              const $5 = $4.tag === "TokenEmpty"
                ? PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $1.of, PureScript$dCST$dRange$dTokenList.TokenEmpty)
                : PureScript$dCST$dRange$dTokenList.$TokenList(
                    "TokenAppend",
                    PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $1.of, PureScript$dCST$dRange$dTokenList.TokenEmpty),
                    $4
                  );
              if ($5.tag === "TokenEmpty") { return $2; }
              if ($2.tag === "TokenEmpty") { return $5; }
              return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $2, $5);
            }
          )
        );
      }
      if (v.tag === "ExprLet") {
        const $1 = v._1;
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenCons",
          $1.keyword,
          PureScript$dCST$dRange$dTokenList.$TokenList(
            "TokenDefer",
            v1 => {
              const $2 = tokensOfLetBinding(dictTokensOf);
              const $3 = foldMap2(a => PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v$1 => $2.tokensOf(a)))($1.bindings);
              const $4 = tokensOfExpr(dictTokensOf).tokensOf($1.body);
              const $5 = $4.tag === "TokenEmpty"
                ? PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $1.in, PureScript$dCST$dRange$dTokenList.TokenEmpty)
                : PureScript$dCST$dRange$dTokenList.$TokenList(
                    "TokenAppend",
                    PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $1.in, PureScript$dCST$dRange$dTokenList.TokenEmpty),
                    $4
                  );
              if ($5.tag === "TokenEmpty") { return $3; }
              if ($3.tag === "TokenEmpty") { return $5; }
              return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $3, $5);
            }
          )
        );
      }
      if (v.tag === "ExprDo") {
        const $1 = v._1.statements;
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenCons",
          v._1.keyword,
          PureScript$dCST$dRange$dTokenList.$TokenList(
            "TokenDefer",
            v1 => {
              const $2 = tokensOfDoStatement(dictTokensOf);
              return foldMap2(a => PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v$1 => $2.tokensOf(a)))($1);
            }
          )
        );
      }
      if (v.tag === "ExprAdo") {
        const $1 = v._1;
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenCons",
          $1.keyword,
          PureScript$dCST$dRange$dTokenList.$TokenList(
            "TokenDefer",
            v1 => {
              const $2 = tokensOfDoStatement(dictTokensOf);
              const $3 = foldMap(a => PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v$1 => $2.tokensOf(a)))($1.statements);
              const $4 = tokensOfExpr(dictTokensOf).tokensOf($1.result);
              const $5 = $4.tag === "TokenEmpty"
                ? PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $1.in, PureScript$dCST$dRange$dTokenList.TokenEmpty)
                : PureScript$dCST$dRange$dTokenList.$TokenList(
                    "TokenAppend",
                    PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $1.in, PureScript$dCST$dRange$dTokenList.TokenEmpty),
                    $4
                  );
              if ($5.tag === "TokenEmpty") { return $3; }
              if ($3.tag === "TokenEmpty") { return $5; }
              return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $3, $5);
            }
          )
        );
      }
      if (v.tag === "ExprError") { return dictTokensOf.tokensOf(v._1); }
      $runtime.fail();
    }
  };
};
const tokensOfDoStatement = dictTokensOf => (
  {
    tokensOf: v => {
      if (v.tag === "DoLet") {
        const $0 = v._2;
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenCons",
          v._1,
          PureScript$dCST$dRange$dTokenList.$TokenList(
            "TokenDefer",
            v1 => {
              const $1 = tokensOfLetBinding(dictTokensOf);
              return foldMap2(a => PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v$1 => $1.tokensOf(a)))($0);
            }
          )
        );
      }
      if (v.tag === "DoDiscard") { return tokensOfExpr(dictTokensOf).tokensOf(v._1); }
      if (v.tag === "DoBind") {
        const $0 = v._3;
        const $1 = v._2;
        const $2 = tokensOfBinder(dictTokensOf).tokensOf(v._1);
        const $3 = PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenDefer",
          v1 => PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $1, tokensOfExpr(dictTokensOf).tokensOf($0))
        );
        if ($2.tag === "TokenEmpty") { return $3; }
        return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $2, $3);
      }
      if (v.tag === "DoError") { return dictTokensOf.tokensOf(v._1); }
      $runtime.fail();
    }
  }
);
const tokensOfInstanceBinding = dictTokensOf => {
  const $0 = tokensOfType(dictTokensOf);
  const $1 = tokensOfBinder(dictTokensOf);
  const tokensOf10 = foldMap(a => PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v => $1.tokensOf(a)));
  return {
    tokensOf: v => {
      if (v.tag === "InstanceBindingSignature") {
        const $2 = PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.label.token, PureScript$dCST$dRange$dTokenList.TokenEmpty);
        const $3 = $0.tokensOf(v._1.value);
        const $4 = $3.tag === "TokenEmpty"
          ? PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.separator, PureScript$dCST$dRange$dTokenList.TokenEmpty)
          : PureScript$dCST$dRange$dTokenList.$TokenList(
              "TokenAppend",
              PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.separator, PureScript$dCST$dRange$dTokenList.TokenEmpty),
              $3
            );
        if ($4.tag === "TokenEmpty") { return $2; }
        return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $2, $4);
      }
      if (v.tag === "InstanceBindingName") {
        const $2 = PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.name.token, PureScript$dCST$dRange$dTokenList.TokenEmpty);
        const $3 = tokensOf10(v._1.binders);
        const $4 = tokensOfGuarded(dictTokensOf).tokensOf(v._1.guarded);
        const $5 = (() => {
          if ($4.tag === "TokenEmpty") { return $3; }
          if ($3.tag === "TokenEmpty") { return $4; }
          return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $3, $4);
        })();
        if ($5.tag === "TokenEmpty") { return $2; }
        return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $2, $5);
      }
      $runtime.fail();
    }
  };
};
const tokensOfInstance = dictTokensOf => {
  const tokensOfType1 = tokensOfType(dictTokensOf);
  const $0 = tokensOfInstanceBinding(dictTokensOf);
  const tokensOf11 = foldMap2(a => PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v => $0.tokensOf(a)));
  return {
    tokensOf: v => {
      const $1 = v.body;
      const $2 = v.head;
      return PureScript$dCST$dRange$dTokenList.$TokenList(
        "TokenCons",
        $2.keyword,
        PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenDefer",
          v1 => {
            const $3 = (() => {
              if ($2.name.tag === "Nothing") { return PureScript$dCST$dRange$dTokenList.TokenEmpty; }
              if ($2.name.tag === "Just") {
                return PureScript$dCST$dRange$dTokenList.$TokenList(
                  "TokenAppend",
                  PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $2.name._1._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty),
                  PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $2.name._1._2, PureScript$dCST$dRange$dTokenList.TokenEmpty)
                );
              }
              $runtime.fail();
            })();
            const $4 = (() => {
              if ($2.constraints.tag === "Nothing") { return PureScript$dCST$dRange$dTokenList.TokenEmpty; }
              if ($2.constraints.tag === "Just") {
                const $4 = tokensOfOneOrDelimited(tokensOfType1).tokensOf($2.constraints._1._1);
                if ($4.tag === "TokenEmpty") {
                  return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $2.constraints._1._2, PureScript$dCST$dRange$dTokenList.TokenEmpty);
                }
                return PureScript$dCST$dRange$dTokenList.$TokenList(
                  "TokenAppend",
                  $4,
                  PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $2.constraints._1._2, PureScript$dCST$dRange$dTokenList.TokenEmpty)
                );
              }
              $runtime.fail();
            })();
            const $5 = (() => {
              const $5 = PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $2.className.token, PureScript$dCST$dRange$dTokenList.TokenEmpty);
              const $6 = foldMap(a => PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v$1 => tokensOfType1.tokensOf(a)))($2.types);
              const $7 = (() => {
                const $7 = (() => {
                  if ($1.tag === "Nothing") { return PureScript$dCST$dRange$dTokenList.TokenEmpty; }
                  if ($1.tag === "Just") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $1._1._1, tokensOf11($1._1._2)); }
                  $runtime.fail();
                })();
                const $8 = (() => {
                  if ($7.tag === "TokenEmpty") { return $6; }
                  if ($6.tag === "TokenEmpty") { return $7; }
                  return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $6, $7);
                })();
                if ($8.tag === "TokenEmpty") { return $5; }
                return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $5, $8);
              })();
              if ($7.tag === "TokenEmpty") { return $4; }
              if ($4.tag === "TokenEmpty") { return $7; }
              return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $4, $7);
            })();
            if ($5.tag === "TokenEmpty") { return $3; }
            if ($3.tag === "TokenEmpty") { return $5; }
            return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $3, $5);
          }
        )
      );
    }
  };
};
const tokensOfDecl = dictTokensOf => {
  const $0 = tokensOfTypeVarBinding(tokensOfName)(dictTokensOf);
  const tokensOf9 = foldMap(a => PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v => $0.tokensOf(a)));
  const tokensOfType1 = tokensOfType(dictTokensOf);
  const $1 = tokensOfOneOrDelimited(tokensOfType1);
  const $2 = v => {
    const $2 = PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v.label.token, PureScript$dCST$dRange$dTokenList.TokenEmpty);
    const $3 = tokensOfType1.tokensOf(v.value);
    const $4 = $3.tag === "TokenEmpty"
      ? PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v.separator, PureScript$dCST$dRange$dTokenList.TokenEmpty)
      : PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenAppend",
          PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v.separator, PureScript$dCST$dRange$dTokenList.TokenEmpty),
          $3
        );
    if ($4.tag === "TokenEmpty") { return $2; }
    return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $2, $4);
  };
  const $3 = tokensOfBinder(dictTokensOf);
  const tokensOf17 = foldMap(a => PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v => $3.tokensOf(a)));
  return {
    tokensOf: v => {
      if (v.tag === "DeclData") {
        const $4 = v._2;
        const $5 = v._1.name;
        const $6 = v._1.vars;
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenCons",
          v._1.keyword,
          PureScript$dCST$dRange$dTokenList.$TokenList(
            "TokenDefer",
            v1 => {
              const $7 = tokensOf9($6);
              const $8 = (() => {
                if ($4.tag === "Nothing") { return PureScript$dCST$dRange$dTokenList.TokenEmpty; }
                if ($4.tag === "Just") {
                  return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $4._1._1, tokensOfSeparated(tokensOfDataCtor(dictTokensOf)).tokensOf($4._1._2));
                }
                $runtime.fail();
              })();
              const $9 = (() => {
                if ($8.tag === "TokenEmpty") { return $7; }
                if ($7.tag === "TokenEmpty") { return $8; }
                return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $7, $8);
              })();
              if ($9.tag === "TokenEmpty") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $5.token, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
              return PureScript$dCST$dRange$dTokenList.$TokenList(
                "TokenAppend",
                PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $5.token, PureScript$dCST$dRange$dTokenList.TokenEmpty),
                $9
              );
            }
          )
        );
      }
      if (v.tag === "DeclType") {
        const $4 = v._1.name;
        const $5 = v._2;
        const $6 = v._3;
        const $7 = v._1.vars;
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenCons",
          v._1.keyword,
          PureScript$dCST$dRange$dTokenList.$TokenList(
            "TokenDefer",
            v1 => {
              const $8 = tokensOf9($7);
              const $9 = tokensOfType1.tokensOf($6);
              const $10 = $9.tag === "TokenEmpty"
                ? PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $5, PureScript$dCST$dRange$dTokenList.TokenEmpty)
                : PureScript$dCST$dRange$dTokenList.$TokenList(
                    "TokenAppend",
                    PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $5, PureScript$dCST$dRange$dTokenList.TokenEmpty),
                    $9
                  );
              const $11 = (() => {
                if ($10.tag === "TokenEmpty") { return $8; }
                if ($8.tag === "TokenEmpty") { return $10; }
                return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $8, $10);
              })();
              if ($11.tag === "TokenEmpty") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $4.token, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
              return PureScript$dCST$dRange$dTokenList.$TokenList(
                "TokenAppend",
                PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $4.token, PureScript$dCST$dRange$dTokenList.TokenEmpty),
                $11
              );
            }
          )
        );
      }
      if (v.tag === "DeclNewtype") {
        const $4 = v._3;
        const $5 = v._1.name;
        const $6 = v._2;
        const $7 = v._4;
        const $8 = v._1.vars;
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenCons",
          v._1.keyword,
          PureScript$dCST$dRange$dTokenList.$TokenList(
            "TokenDefer",
            v1 => {
              const $9 = tokensOf9($8);
              const $10 = tokensOfType1.tokensOf($7);
              const $11 = $10.tag === "TokenEmpty"
                ? PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $4.token, PureScript$dCST$dRange$dTokenList.TokenEmpty)
                : PureScript$dCST$dRange$dTokenList.$TokenList(
                    "TokenAppend",
                    PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $4.token, PureScript$dCST$dRange$dTokenList.TokenEmpty),
                    $10
                  );
              const $12 = (() => {
                const $12 = $11.tag === "TokenEmpty"
                  ? PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $6, PureScript$dCST$dRange$dTokenList.TokenEmpty)
                  : PureScript$dCST$dRange$dTokenList.$TokenList(
                      "TokenAppend",
                      PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $6, PureScript$dCST$dRange$dTokenList.TokenEmpty),
                      $11
                    );
                if ($12.tag === "TokenEmpty") { return $9; }
                if ($9.tag === "TokenEmpty") { return $12; }
                return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $9, $12);
              })();
              if ($12.tag === "TokenEmpty") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $5.token, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
              return PureScript$dCST$dRange$dTokenList.$TokenList(
                "TokenAppend",
                PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $5.token, PureScript$dCST$dRange$dTokenList.TokenEmpty),
                $12
              );
            }
          )
        );
      }
      if (v.tag === "DeclClass") {
        const $4 = v._1.fundeps;
        const $5 = v._2;
        const $6 = v._1.name;
        const $7 = v._1.super;
        const $8 = v._1.vars;
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenCons",
          v._1.keyword,
          PureScript$dCST$dRange$dTokenList.$TokenList(
            "TokenDefer",
            v1 => {
              const $9 = (() => {
                if ($7.tag === "Nothing") { return PureScript$dCST$dRange$dTokenList.TokenEmpty; }
                if ($7.tag === "Just") {
                  const $9 = $1.tokensOf($7._1._1);
                  if ($9.tag === "TokenEmpty") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $7._1._2, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
                  return PureScript$dCST$dRange$dTokenList.$TokenList(
                    "TokenAppend",
                    $9,
                    PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $7._1._2, PureScript$dCST$dRange$dTokenList.TokenEmpty)
                  );
                }
                $runtime.fail();
              })();
              const $10 = tokensOf9($8);
              const $11 = (() => {
                if ($4.tag === "Nothing") { return PureScript$dCST$dRange$dTokenList.TokenEmpty; }
                if ($4.tag === "Just") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $4._1._1, tokensOf6($4._1._2)); }
                $runtime.fail();
              })();
              const $12 = (() => {
                const $12 = (() => {
                  if ($5.tag === "Nothing") { return PureScript$dCST$dRange$dTokenList.TokenEmpty; }
                  if ($5.tag === "Just") {
                    return PureScript$dCST$dRange$dTokenList.$TokenList(
                      "TokenCons",
                      $5._1._1,
                      foldMap2(a => PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v$1 => $2(a)))($5._1._2)
                    );
                  }
                  $runtime.fail();
                })();
                const $13 = (() => {
                  const $13 = (() => {
                    if ($12.tag === "TokenEmpty") { return $11; }
                    if ($11.tag === "TokenEmpty") { return $12; }
                    return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $11, $12);
                  })();
                  if ($13.tag === "TokenEmpty") { return $10; }
                  if ($10.tag === "TokenEmpty") { return $13; }
                  return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $10, $13);
                })();
                if ($13.tag === "TokenEmpty") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $6.token, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
                return PureScript$dCST$dRange$dTokenList.$TokenList(
                  "TokenAppend",
                  PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $6.token, PureScript$dCST$dRange$dTokenList.TokenEmpty),
                  $13
                );
              })();
              if ($12.tag === "TokenEmpty") { return $9; }
              if ($9.tag === "TokenEmpty") { return $12; }
              return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $9, $12);
            }
          )
        );
      }
      if (v.tag === "DeclInstanceChain") { return tokensOfSeparated(tokensOfInstance(dictTokensOf)).tokensOf(v._1); }
      if (v.tag === "DeclDerive") {
        const $4 = v._3;
        const $5 = v._2;
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenCons",
          v._1,
          PureScript$dCST$dRange$dTokenList.$TokenList(
            "TokenDefer",
            v1 => {
              const $6 = (() => {
                if ($5.tag === "Nothing") { return PureScript$dCST$dRange$dTokenList.TokenEmpty; }
                if ($5.tag === "Just") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $5._1, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
                $runtime.fail();
              })();
              const $7 = (() => {
                if ($4.name.tag === "Nothing") { return PureScript$dCST$dRange$dTokenList.TokenEmpty; }
                if ($4.name.tag === "Just") {
                  return PureScript$dCST$dRange$dTokenList.$TokenList(
                    "TokenAppend",
                    PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $4.name._1._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty),
                    PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $4.name._1._2, PureScript$dCST$dRange$dTokenList.TokenEmpty)
                  );
                }
                $runtime.fail();
              })();
              const $8 = (() => {
                if ($4.constraints.tag === "Nothing") { return PureScript$dCST$dRange$dTokenList.TokenEmpty; }
                if ($4.constraints.tag === "Just") {
                  const $8 = $1.tokensOf($4.constraints._1._1);
                  if ($8.tag === "TokenEmpty") {
                    return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $4.constraints._1._2, PureScript$dCST$dRange$dTokenList.TokenEmpty);
                  }
                  return PureScript$dCST$dRange$dTokenList.$TokenList(
                    "TokenAppend",
                    $8,
                    PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $4.constraints._1._2, PureScript$dCST$dRange$dTokenList.TokenEmpty)
                  );
                }
                $runtime.fail();
              })();
              const $9 = (() => {
                const $9 = (() => {
                  const $9 = PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $4.className.token, PureScript$dCST$dRange$dTokenList.TokenEmpty);
                  const $10 = (() => {
                    const $10 = foldMap(a => PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v$1 => tokensOfType1.tokensOf(a)))($4.types);
                    const $11 = $10.tag === "TokenEmpty" ? $9 : PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $9, $10);
                    if ($11.tag === "TokenEmpty") { return $8; }
                    if ($8.tag === "TokenEmpty") { return $11; }
                    return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $8, $11);
                  })();
                  if ($10.tag === "TokenEmpty") { return $7; }
                  if ($7.tag === "TokenEmpty") { return $10; }
                  return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $7, $10);
                })();
                if ($9.tag === "TokenEmpty") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $4.keyword, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
                return PureScript$dCST$dRange$dTokenList.$TokenList(
                  "TokenAppend",
                  PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $4.keyword, PureScript$dCST$dRange$dTokenList.TokenEmpty),
                  $9
                );
              })();
              if ($9.tag === "TokenEmpty") { return $6; }
              if ($6.tag === "TokenEmpty") { return $9; }
              return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $6, $9);
            }
          )
        );
      }
      if (v.tag === "DeclKindSignature") {
        const $4 = v._2;
        return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1, PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v1 => $2($4)));
      }
      if (v.tag === "DeclSignature") { return $2(v._1); }
      if (v.tag === "DeclValue") {
        const $4 = v._1.binders;
        const $5 = v._1.guarded;
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenAppend",
          PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v._1.name.token, PureScript$dCST$dRange$dTokenList.TokenEmpty),
          PureScript$dCST$dRange$dTokenList.$TokenList(
            "TokenDefer",
            v1 => {
              const $6 = tokensOf17($4);
              const $7 = tokensOfGuarded(dictTokensOf).tokensOf($5);
              if ($7.tag === "TokenEmpty") { return $6; }
              if ($6.tag === "TokenEmpty") { return $7; }
              return PureScript$dCST$dRange$dTokenList.$TokenList("TokenAppend", $6, $7);
            }
          )
        );
      }
      if (v.tag === "DeclFixity") {
        const $4 = v._1.operator;
        const $5 = v._1.prec._1;
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenCons",
          v._1.keyword._1,
          PureScript$dCST$dRange$dTokenList.$TokenList(
            "TokenDefer",
            v1 => PureScript$dCST$dRange$dTokenList.$TokenList(
              "TokenCons",
              $5,
              (() => {
                if ($4.tag === "FixityValue") {
                  return PureScript$dCST$dRange$dTokenList.$TokenList(
                    "TokenAppend",
                    PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $4._1.token, PureScript$dCST$dRange$dTokenList.TokenEmpty),
                    PureScript$dCST$dRange$dTokenList.$TokenList(
                      "TokenAppend",
                      PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $4._2, PureScript$dCST$dRange$dTokenList.TokenEmpty),
                      PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $4._3.token, PureScript$dCST$dRange$dTokenList.TokenEmpty)
                    )
                  );
                }
                if ($4.tag === "FixityType") {
                  return PureScript$dCST$dRange$dTokenList.$TokenList(
                    "TokenCons",
                    $4._1,
                    PureScript$dCST$dRange$dTokenList.$TokenList(
                      "TokenAppend",
                      PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $4._2.token, PureScript$dCST$dRange$dTokenList.TokenEmpty),
                      PureScript$dCST$dRange$dTokenList.$TokenList(
                        "TokenAppend",
                        PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $4._3, PureScript$dCST$dRange$dTokenList.TokenEmpty),
                        PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $4._4.token, PureScript$dCST$dRange$dTokenList.TokenEmpty)
                      )
                    )
                  );
                }
                $runtime.fail();
              })()
            )
          )
        );
      }
      if (v.tag === "DeclForeign") {
        const $4 = v._3;
        const $5 = v._2;
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenCons",
          v._1,
          PureScript$dCST$dRange$dTokenList.$TokenList(
            "TokenDefer",
            v1 => PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $5, tokensOfForeign(dictTokensOf).tokensOf($4))
          )
        );
      }
      if (v.tag === "DeclRole") {
        const $4 = v._3;
        const $5 = v._2;
        const $6 = v._4;
        return PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenCons",
          v._1,
          PureScript$dCST$dRange$dTokenList.$TokenList(
            "TokenDefer",
            v1 => {
              const $7 = foldMap2(v2 => PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", v2._1, PureScript$dCST$dRange$dTokenList.TokenEmpty))($6);
              const $8 = $7.tag === "TokenEmpty"
                ? PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $4.token, PureScript$dCST$dRange$dTokenList.TokenEmpty)
                : PureScript$dCST$dRange$dTokenList.$TokenList(
                    "TokenAppend",
                    PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $4.token, PureScript$dCST$dRange$dTokenList.TokenEmpty),
                    $7
                  );
              if ($8.tag === "TokenEmpty") { return PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $5, PureScript$dCST$dRange$dTokenList.TokenEmpty); }
              return PureScript$dCST$dRange$dTokenList.$TokenList(
                "TokenAppend",
                PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $5, PureScript$dCST$dRange$dTokenList.TokenEmpty),
                $8
              );
            }
          )
        );
      }
      if (v.tag === "DeclError") { return dictTokensOf.tokensOf(v._1); }
      $runtime.fail();
    }
  };
};
const tokensOfModule = dictTokensOf => {
  const $0 = tokensOfSeparated(tokensOfExport(dictTokensOf));
  const tokensOf10 = tokensOfImportDecl(dictTokensOf).tokensOf;
  const tokensOf11 = tokensOfDecl(dictTokensOf).tokensOf;
  return {
    tokensOf: v => {
      const $1 = v.body;
      const $2 = v.header;
      return PureScript$dCST$dRange$dTokenList.$TokenList(
        "TokenCons",
        $2.keyword,
        PureScript$dCST$dRange$dTokenList.$TokenList(
          "TokenAppend",
          PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $2.name.token, PureScript$dCST$dRange$dTokenList.TokenEmpty),
          PureScript$dCST$dRange$dTokenList.$TokenList(
            "TokenAppend",
            PureScript$dCST$dRange$dTokenList.$TokenList(
              "TokenDefer",
              v1 => {
                if ($2.exports.tag === "Nothing") { return PureScript$dCST$dRange$dTokenList.TokenEmpty; }
                if ($2.exports.tag === "Just") {
                  const $3 = $2.exports._1.value;
                  return PureScript$dCST$dRange$dTokenList.$TokenList(
                    "TokenWrap",
                    $2.exports._1.open,
                    PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v1$1 => $0.tokensOf($3)),
                    $2.exports._1.close
                  );
                }
                $runtime.fail();
              }
            ),
            PureScript$dCST$dRange$dTokenList.$TokenList(
              "TokenAppend",
              PureScript$dCST$dRange$dTokenList.$TokenList("TokenCons", $2.where, PureScript$dCST$dRange$dTokenList.TokenEmpty),
              PureScript$dCST$dRange$dTokenList.$TokenList(
                "TokenAppend",
                PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v1 => foldMap(tokensOf10)($2.imports)),
                PureScript$dCST$dRange$dTokenList.$TokenList("TokenDefer", v1 => foldMap(tokensOf11)($1.decls))
              )
            )
          )
        )
      );
    }
  };
};
const rangeOfWrapped = {rangeOf: v => ({start: v.open.range.start, end: v.close.range.end})};
const rangeOfVoid = {rangeOf: Data$dVoid.absurd};
const rangeOfRecoveredError = {
  rangeOf: v => {
    if (v.tokens.length > 0) {
      return {
        start: (() => {
          if (0 < v.tokens.length) { return v.tokens[0].range.start; }
          $runtime.fail();
        })(),
        end: (() => {
          const $0 = v.tokens.length - 1 | 0;
          if ($0 >= 0 && $0 < v.tokens.length) { return v.tokens[$0].range.end; }
          $runtime.fail();
        })()
      };
    }
    return {start: v.position, end: v.position};
  }
};
const rangeOfQualifiedName = {rangeOf: v => v.token.range};
const rangeOfName = {rangeOf: v => v.token.range};
const rangeOfModule = {rangeOf: v => ({start: v.header.keyword.range.start, end: v.body.end})};
const rangeOf = dict => dict.rangeOf;
const rangeOfClassFundep = {
  rangeOf: v => {
    if (v.tag === "FundepDetermined") {
      return {
        start: v._1.range.start,
        end: (() => {
          const $0 = v._2.length - 1 | 0;
          if ($0 >= 0 && $0 < v._2.length) { return v._2[$0].token.range.end; }
          $runtime.fail();
        })()
      };
    }
    if (v.tag === "FundepDetermines") {
      return {
        start: (() => {
          if (0 < v._1.length) { return v._1[0].token.range.start; }
          $runtime.fail();
        })(),
        end: (() => {
          const $0 = v._3.length - 1 | 0;
          if ($0 >= 0 && $0 < v._3.length) { return v._3[$0].token.range.end; }
          $runtime.fail();
        })()
      };
    }
    $runtime.fail();
  }
};
const rangeOfDataMembers = {
  rangeOf: v => {
    if (v.tag === "DataAll") { return v._1.range; }
    if (v.tag === "DataEnumerated") { return {start: v._1.open.range.start, end: v._1.close.range.end}; }
    $runtime.fail();
  }
};
const rangeOfExport = dictRangeOf => (
  {
    rangeOf: v => {
      if (v.tag === "ExportValue") { return v._1.token.range; }
      if (v.tag === "ExportOp") { return v._1.token.range; }
      if (v.tag === "ExportType") {
        if (v._2.tag === "Nothing") { return v._1.token.range; }
        if (v._2.tag === "Just") {
          return {
            start: v._1.token.range.start,
            end: (() => {
              if (v._2._1.tag === "DataAll") { return v._2._1._1.range.end; }
              if (v._2._1.tag === "DataEnumerated") { return v._2._1._1.close.range.end; }
              $runtime.fail();
            })()
          };
        }
        $runtime.fail();
      }
      if (v.tag === "ExportTypeOp") { return {start: v._1.range.start, end: v._2.token.range.end}; }
      if (v.tag === "ExportClass") { return {start: v._1.range.start, end: v._2.token.range.end}; }
      if (v.tag === "ExportModule") { return {start: v._1.range.start, end: v._2.token.range.end}; }
      if (v.tag === "ExportError") { return dictRangeOf.rangeOf(v._1); }
      $runtime.fail();
    }
  }
);
const rangeOfFixityOp = {
  rangeOf: v => {
    if (v.tag === "FixityValue") { return {start: v._1.token.range.start, end: v._3.token.range.end}; }
    if (v.tag === "FixityType") { return {start: v._1.range.start, end: v._4.token.range.end}; }
    $runtime.fail();
  }
};
const rangeOfImport = dictRangeOf => (
  {
    rangeOf: v => {
      if (v.tag === "ImportValue") { return v._1.token.range; }
      if (v.tag === "ImportOp") { return v._1.token.range; }
      if (v.tag === "ImportType") {
        if (v._2.tag === "Nothing") { return v._1.token.range; }
        if (v._2.tag === "Just") {
          return {
            start: v._1.token.range.start,
            end: (() => {
              if (v._2._1.tag === "DataAll") { return v._2._1._1.range.end; }
              if (v._2._1.tag === "DataEnumerated") { return v._2._1._1.close.range.end; }
              $runtime.fail();
            })()
          };
        }
        $runtime.fail();
      }
      if (v.tag === "ImportTypeOp") { return {start: v._1.range.start, end: v._2.token.range.end}; }
      if (v.tag === "ImportClass") { return {start: v._1.range.start, end: v._2.token.range.end}; }
      if (v.tag === "ImportError") { return dictRangeOf.rangeOf(v._1); }
      $runtime.fail();
    }
  }
);
const rangeOfImportDecl = {
  rangeOf: v => (
    {
      start: v.keyword.range.start,
      end: (() => {
        if (v.qualified.tag === "Nothing") {
          if (v.names.tag === "Nothing") { return v.module.token.range.end; }
          if (v.names.tag === "Just") { return v.names._1._2.close.range.end; }
          $runtime.fail();
        }
        if (v.qualified.tag === "Just") { return v.qualified._1._2.token.range.end; }
        $runtime.fail();
      })()
    }
  )
};
const rangeOfLabeled = dictRangeOf => dictRangeOf1 => ({rangeOf: v => ({start: dictRangeOf.rangeOf(v.label).start, end: dictRangeOf1.rangeOf(v.value).end})});
const rangeOfOneOrDelimited = dictRangeOf => (
  {
    rangeOf: v => {
      if (v.tag === "One") { return dictRangeOf.rangeOf(v._1); }
      if (v.tag === "Many") { return {start: v._1.open.range.start, end: v._1.close.range.end}; }
      $runtime.fail();
    }
  }
);
const rangeOfPrefixed = dictRangeOf => (
  {
    rangeOf: v => {
      if (v.prefix.tag === "Just") { return {start: v.prefix._1.range.start, end: dictRangeOf.rangeOf(v.value).end}; }
      if (v.prefix.tag === "Nothing") { return dictRangeOf.rangeOf(v.value); }
      $runtime.fail();
    }
  }
);
const rangeOfSeparated = dictRangeOf => (
  {
    rangeOf: v => {
      const $0 = v.tail.length - 1 | 0;
      if ($0 >= 0 && $0 < v.tail.length) { return {start: dictRangeOf.rangeOf(v.head).start, end: dictRangeOf.rangeOf(v.tail[$0]._2).end}; }
      return dictRangeOf.rangeOf(v.head);
    }
  }
);
const rangeOf7 = v => {
  const $0 = v.tail.length - 1 | 0;
  if ($0 >= 0 && $0 < v.tail.length) { return {start: rangeOfClassFundep.rangeOf(v.head).start, end: rangeOfClassFundep.rangeOf(v.tail[$0]._2).end}; }
  return rangeOfClassFundep.rangeOf(v.head);
};
const rangeOfType = dictRangeOf => (
  {
    rangeOf: v => {
      if (v.tag === "TypeVar") { return v._1.token.range; }
      if (v.tag === "TypeConstructor") { return v._1.token.range; }
      if (v.tag === "TypeWildcard") { return v._1.range; }
      if (v.tag === "TypeHole") { return v._1.token.range; }
      if (v.tag === "TypeString") { return v._1.range; }
      if (v.tag === "TypeInt") {
        if (v._1.tag === "Nothing") { return v._2.range; }
        if (v._1.tag === "Just") { return {start: v._1._1.range.start, end: v._2.range.end}; }
        $runtime.fail();
      }
      if (v.tag === "TypeRow") { return {start: v._1.open.range.start, end: v._1.close.range.end}; }
      if (v.tag === "TypeRecord") { return {start: v._1.open.range.start, end: v._1.close.range.end}; }
      if (v.tag === "TypeForall") { return {start: v._1.range.start, end: rangeOfType(dictRangeOf).rangeOf(v._4).end}; }
      if (v.tag === "TypeKinded") { return {start: rangeOfType(dictRangeOf).rangeOf(v._1).start, end: rangeOfType(dictRangeOf).rangeOf(v._3).end}; }
      if (v.tag === "TypeApp") {
        return {
          start: rangeOfType(dictRangeOf).rangeOf(v._1).start,
          end: rangeOfType(dictRangeOf).rangeOf((() => {
            const $0 = v._2.length - 1 | 0;
            if ($0 >= 0 && $0 < v._2.length) { return v._2[$0]; }
            $runtime.fail();
          })()).end
        };
      }
      if (v.tag === "TypeOp") {
        return {
          start: rangeOfType(dictRangeOf).rangeOf(v._1).start,
          end: rangeOfType(dictRangeOf).rangeOf((() => {
            const $0 = v._2.length - 1 | 0;
            if ($0 >= 0 && $0 < v._2.length) { return v._2[$0]._2; }
            $runtime.fail();
          })()).end
        };
      }
      if (v.tag === "TypeOpName") { return v._1.token.range; }
      if (v.tag === "TypeArrow") { return {start: rangeOfType(dictRangeOf).rangeOf(v._1).start, end: rangeOfType(dictRangeOf).rangeOf(v._3).end}; }
      if (v.tag === "TypeArrowName") { return v._1.range; }
      if (v.tag === "TypeConstrained") { return {start: rangeOfType(dictRangeOf).rangeOf(v._1).start, end: rangeOfType(dictRangeOf).rangeOf(v._3).end}; }
      if (v.tag === "TypeParens") { return {start: v._1.open.range.start, end: v._1.close.range.end}; }
      if (v.tag === "TypeError") { return dictRangeOf.rangeOf(v._1); }
      $runtime.fail();
    }
  }
);
const rangeOfAppSpine = dictRangeOf => dictRangeOf1 => (
  {
    rangeOf: v => {
      if (v.tag === "AppType") { return {start: v._1.range.start, end: rangeOfType(dictRangeOf).rangeOf(v._2).end}; }
      if (v.tag === "AppTerm") { return dictRangeOf1.rangeOf(v._1); }
      $runtime.fail();
    }
  }
);
const rangeOfBinder = dictRangeOf => (
  {
    rangeOf: v => {
      if (v.tag === "BinderWildcard") { return v._1.range; }
      if (v.tag === "BinderVar") { return v._1.token.range; }
      if (v.tag === "BinderNamed") { return {start: v._1.token.range.start, end: rangeOfBinder(dictRangeOf).rangeOf(v._3).end}; }
      if (v.tag === "BinderConstructor") {
        const $0 = v._2.length - 1 | 0;
        if ($0 >= 0 && $0 < v._2.length) { return {start: v._1.token.range.start, end: rangeOfBinder(dictRangeOf).rangeOf(v._2[$0]).end}; }
        return v._1.token.range;
      }
      if (v.tag === "BinderBoolean") { return v._1.range; }
      if (v.tag === "BinderChar") { return v._1.range; }
      if (v.tag === "BinderString") { return v._1.range; }
      if (v.tag === "BinderInt") {
        if (v._1.tag === "Nothing") { return v._2.range; }
        if (v._1.tag === "Just") { return {start: v._1._1.range.start, end: v._2.range.end}; }
        $runtime.fail();
      }
      if (v.tag === "BinderNumber") {
        if (v._1.tag === "Nothing") { return v._2.range; }
        if (v._1.tag === "Just") { return {start: v._1._1.range.start, end: v._2.range.end}; }
        $runtime.fail();
      }
      if (v.tag === "BinderArray") { return {start: v._1.open.range.start, end: v._1.close.range.end}; }
      if (v.tag === "BinderRecord") { return {start: v._1.open.range.start, end: v._1.close.range.end}; }
      if (v.tag === "BinderParens") { return {start: v._1.open.range.start, end: v._1.close.range.end}; }
      if (v.tag === "BinderTyped") { return {start: rangeOfBinder(dictRangeOf).rangeOf(v._1).start, end: rangeOfType(dictRangeOf).rangeOf(v._3).end}; }
      if (v.tag === "BinderOp") {
        return {
          start: rangeOfBinder(dictRangeOf).rangeOf(v._1).start,
          end: rangeOfBinder(dictRangeOf).rangeOf((() => {
            const $0 = v._2.length - 1 | 0;
            if ($0 >= 0 && $0 < v._2.length) { return v._2[$0]._2; }
            $runtime.fail();
          })()).end
        };
      }
      if (v.tag === "BinderError") { return dictRangeOf.rangeOf(v._1); }
      $runtime.fail();
    }
  }
);
const rangeOfDataCtor = dictRangeOf => (
  {
    rangeOf: v => {
      const $0 = v.fields.length - 1 | 0;
      return {start: v.name.token.range.start, end: $0 >= 0 && $0 < v.fields.length ? rangeOfType(dictRangeOf).rangeOf(v.fields[$0]).end : v.name.token.range.end};
    }
  }
);
const rangeOfForeign = dictRangeOf => {
  const $0 = rangeOfType(dictRangeOf);
  return {
    rangeOf: v => {
      if (v.tag === "ForeignValue") { return {start: v._1.label.token.range.start, end: $0.rangeOf(v._1.value).end}; }
      if (v.tag === "ForeignData") { return {start: v._1.range.start, end: $0.rangeOf(v._2.value).end}; }
      if (v.tag === "ForeignKind") { return {start: v._1.range.start, end: v._2.token.range.end}; }
      $runtime.fail();
    }
  };
};
const rangeOfTypeVarBinding = dictRangeOf => (
  {
    rangeOf: v => {
      if (v.tag === "TypeVarKinded") { return {start: v._1.open.range.start, end: v._1.close.range.end}; }
      if (v.tag === "TypeVarName") { return dictRangeOf.rangeOf(v._1); }
      $runtime.fail();
    }
  }
);
const rangeOfWhere = dictRangeOf => (
  {
    rangeOf: v => {
      if (v.bindings.tag === "Nothing") { return rangeOfExpr(dictRangeOf).rangeOf(v.expr); }
      if (v.bindings.tag === "Just") {
        return {
          start: rangeOfExpr(dictRangeOf).rangeOf(v.expr).start,
          end: rangeOfLetBinding(dictRangeOf).rangeOf((() => {
            const $0 = v.bindings._1._2.length - 1 | 0;
            if ($0 >= 0 && $0 < v.bindings._1._2.length) { return v.bindings._1._2[$0]; }
            $runtime.fail();
          })()).end
        };
      }
      $runtime.fail();
    }
  }
);
const rangeOfLetBinding = dictRangeOf => {
  const $0 = rangeOfType(dictRangeOf);
  return {
    rangeOf: v => {
      if (v.tag === "LetBindingSignature") { return {start: v._1.label.token.range.start, end: $0.rangeOf(v._1.value).end}; }
      if (v.tag === "LetBindingName") { return {start: v._1.name.token.range.start, end: rangeOfGuarded(dictRangeOf).rangeOf(v._1.guarded).end}; }
      if (v.tag === "LetBindingPattern") { return {start: rangeOfBinder(dictRangeOf).rangeOf(v._1).start, end: rangeOfWhere(dictRangeOf).rangeOf(v._3).end}; }
      if (v.tag === "LetBindingError") { return dictRangeOf.rangeOf(v._1); }
      $runtime.fail();
    }
  };
};
const rangeOfGuardedExpr = dictRangeOf => ({rangeOf: v => ({start: v.bar.range.start, end: rangeOfWhere(dictRangeOf).rangeOf(v.where).end})});
const rangeOfGuarded = dictRangeOf => (
  {
    rangeOf: v => {
      if (v.tag === "Unconditional") { return {start: v._1.range.start, end: rangeOfWhere(dictRangeOf).rangeOf(v._2).end}; }
      if (v.tag === "Guarded") {
        return {
          start: rangeOfGuardedExpr(dictRangeOf).rangeOf((() => {
            if (0 < v._1.length) { return v._1[0]; }
            $runtime.fail();
          })()).start,
          end: rangeOfGuardedExpr(dictRangeOf).rangeOf((() => {
            const $0 = v._1.length - 1 | 0;
            if ($0 >= 0 && $0 < v._1.length) { return v._1[$0]; }
            $runtime.fail();
          })()).end
        };
      }
      $runtime.fail();
    }
  }
);
const rangeOfExpr = dictRangeOf => (
  {
    rangeOf: v => {
      if (v.tag === "ExprHole") { return v._1.token.range; }
      if (v.tag === "ExprSection") { return v._1.range; }
      if (v.tag === "ExprIdent") { return v._1.token.range; }
      if (v.tag === "ExprConstructor") { return v._1.token.range; }
      if (v.tag === "ExprBoolean") { return v._1.range; }
      if (v.tag === "ExprChar") { return v._1.range; }
      if (v.tag === "ExprString") { return v._1.range; }
      if (v.tag === "ExprInt") { return v._1.range; }
      if (v.tag === "ExprNumber") { return v._1.range; }
      if (v.tag === "ExprArray") { return {start: v._1.open.range.start, end: v._1.close.range.end}; }
      if (v.tag === "ExprRecord") { return {start: v._1.open.range.start, end: v._1.close.range.end}; }
      if (v.tag === "ExprParens") { return {start: v._1.open.range.start, end: v._1.close.range.end}; }
      if (v.tag === "ExprTyped") { return {start: rangeOfExpr(dictRangeOf).rangeOf(v._1).start, end: rangeOfType(dictRangeOf).rangeOf(v._3).end}; }
      if (v.tag === "ExprInfix") {
        return {
          start: rangeOfExpr(dictRangeOf).rangeOf(v._1).start,
          end: rangeOfExpr(dictRangeOf).rangeOf((() => {
            const $0 = v._2.length - 1 | 0;
            if ($0 >= 0 && $0 < v._2.length) { return v._2[$0]._2; }
            $runtime.fail();
          })()).end
        };
      }
      if (v.tag === "ExprOp") {
        return {
          start: rangeOfExpr(dictRangeOf).rangeOf(v._1).start,
          end: rangeOfExpr(dictRangeOf).rangeOf((() => {
            const $0 = v._2.length - 1 | 0;
            if ($0 >= 0 && $0 < v._2.length) { return v._2[$0]._2; }
            $runtime.fail();
          })()).end
        };
      }
      if (v.tag === "ExprOpName") { return v._1.token.range; }
      if (v.tag === "ExprNegate") { return {start: v._1.range.start, end: rangeOfExpr(dictRangeOf).rangeOf(v._2).end}; }
      if (v.tag === "ExprRecordAccessor") {
        return {
          start: rangeOfExpr(dictRangeOf).rangeOf(v._1.expr).start,
          end: (() => {
            const $0 = v._1.path.tail.length - 1 | 0;
            if ($0 >= 0 && $0 < v._1.path.tail.length) { return v._1.path.tail[$0]._2.token.range.end; }
            return v._1.path.head.token.range.end;
          })()
        };
      }
      if (v.tag === "ExprRecordUpdate") { return {start: rangeOfExpr(dictRangeOf).rangeOf(v._1).start, end: v._2.close.range.end}; }
      if (v.tag === "ExprApp") {
        return {
          start: rangeOfExpr(dictRangeOf).rangeOf(v._1).start,
          end: rangeOfAppSpine(dictRangeOf)(rangeOfExpr(dictRangeOf)).rangeOf((() => {
            const $0 = v._2.length - 1 | 0;
            if ($0 >= 0 && $0 < v._2.length) { return v._2[$0]; }
            $runtime.fail();
          })()).end
        };
      }
      if (v.tag === "ExprLambda") { return {start: v._1.symbol.range.start, end: rangeOfExpr(dictRangeOf).rangeOf(v._1.body).end}; }
      if (v.tag === "ExprIf") { return {start: v._1.keyword.range.start, end: rangeOfExpr(dictRangeOf).rangeOf(v._1.false).end}; }
      if (v.tag === "ExprCase") {
        return {
          start: v._1.keyword.range.start,
          end: rangeOfGuarded(dictRangeOf).rangeOf((() => {
            const $0 = v._1.branches.length - 1 | 0;
            if ($0 >= 0 && $0 < v._1.branches.length) { return v._1.branches[$0]._2; }
            $runtime.fail();
          })()).end
        };
      }
      if (v.tag === "ExprLet") { return {start: v._1.keyword.range.start, end: rangeOfExpr(dictRangeOf).rangeOf(v._1.body).end}; }
      if (v.tag === "ExprDo") {
        return {
          start: v._1.keyword.range.start,
          end: rangeOfDoStatement(dictRangeOf).rangeOf((() => {
            const $0 = v._1.statements.length - 1 | 0;
            if ($0 >= 0 && $0 < v._1.statements.length) { return v._1.statements[$0]; }
            $runtime.fail();
          })()).end
        };
      }
      if (v.tag === "ExprAdo") { return {start: v._1.keyword.range.start, end: rangeOfExpr(dictRangeOf).rangeOf(v._1.result).end}; }
      if (v.tag === "ExprError") { return dictRangeOf.rangeOf(v._1); }
      $runtime.fail();
    }
  }
);
const rangeOfDoStatement = dictRangeOf => (
  {
    rangeOf: v => {
      if (v.tag === "DoLet") {
        return {
          start: v._1.range.start,
          end: rangeOfLetBinding(dictRangeOf).rangeOf((() => {
            const $0 = v._2.length - 1 | 0;
            if ($0 >= 0 && $0 < v._2.length) { return v._2[$0]; }
            $runtime.fail();
          })()).end
        };
      }
      if (v.tag === "DoDiscard") { return rangeOfExpr(dictRangeOf).rangeOf(v._1); }
      if (v.tag === "DoBind") { return {start: rangeOfBinder(dictRangeOf).rangeOf(v._1).start, end: rangeOfExpr(dictRangeOf).rangeOf(v._3).end}; }
      if (v.tag === "DoError") { return dictRangeOf.rangeOf(v._1); }
      $runtime.fail();
    }
  }
);
const rangeOfInstanceBinding = dictRangeOf => {
  const $0 = rangeOfType(dictRangeOf);
  return {
    rangeOf: v => {
      if (v.tag === "InstanceBindingSignature") { return {start: v._1.label.token.range.start, end: $0.rangeOf(v._1.value).end}; }
      if (v.tag === "InstanceBindingName") { return {start: v._1.name.token.range.start, end: rangeOfGuarded(dictRangeOf).rangeOf(v._1.guarded).end}; }
      $runtime.fail();
    }
  };
};
const rangeOfInstance = dictRangeOf => (
  {
    rangeOf: v => (
      {
        start: v.head.keyword.range.start,
        end: (() => {
          if (v.body.tag === "Nothing") {
            const $0 = v.head.types.length - 1 | 0;
            if ($0 >= 0 && $0 < v.head.types.length) { return rangeOfType(dictRangeOf).rangeOf(v.head.types[$0]).end; }
            return v.head.className.token.range.end;
          }
          if (v.body.tag === "Just") {
            return rangeOfInstanceBinding(dictRangeOf).rangeOf((() => {
              const $0 = v.body._1._2.length - 1 | 0;
              if ($0 >= 0 && $0 < v.body._1._2.length) { return v.body._1._2[$0]; }
              $runtime.fail();
            })()).end;
          }
          $runtime.fail();
        })()
      }
    )
  }
);
const rangeOfDecl = dictRangeOf => {
  const rangeOfType1 = rangeOfType(dictRangeOf);
  const $0 = rangeOfInstance(dictRangeOf);
  return {
    rangeOf: v => {
      if (v.tag === "DeclData") {
        return {
          start: v._1.keyword.range.start,
          end: (() => {
            if (v._2.tag === "Nothing") {
              const $1 = v._1.vars.length - 1 | 0;
              if ($1 >= 0 && $1 < v._1.vars.length) {
                if (v._1.vars[$1].tag === "TypeVarKinded") { return v._1.vars[$1]._1.close.range.end; }
                if (v._1.vars[$1].tag === "TypeVarName") { return v._1.vars[$1]._1.token.range.end; }
                $runtime.fail();
              }
              return v._1.name.token.range.end;
            }
            if (v._2.tag === "Just") {
              return rangeOfDataCtor(dictRangeOf).rangeOf((() => {
                const $1 = v._2._1._2.tail.length - 1 | 0;
                if ($1 >= 0 && $1 < v._2._1._2.tail.length) { return v._2._1._2.tail[$1]._2; }
                return v._2._1._2.head;
              })()).end;
            }
            $runtime.fail();
          })()
        };
      }
      if (v.tag === "DeclType") { return {start: v._1.keyword.range.start, end: rangeOfType1.rangeOf(v._3).end}; }
      if (v.tag === "DeclNewtype") { return {start: v._1.keyword.range.start, end: rangeOfType1.rangeOf(v._4).end}; }
      if (v.tag === "DeclClass") {
        return {
          start: v._1.keyword.range.start,
          end: (() => {
            if (v._2.tag === "Nothing") {
              if (v._1.fundeps.tag === "Nothing") {
                const $1 = v._1.vars.length - 1 | 0;
                if ($1 >= 0 && $1 < v._1.vars.length) {
                  if (v._1.vars[$1].tag === "TypeVarKinded") { return v._1.vars[$1]._1.close.range.end; }
                  if (v._1.vars[$1].tag === "TypeVarName") { return v._1.vars[$1]._1.token.range.end; }
                  $runtime.fail();
                }
                return v._1.name.token.range.end;
              }
              if (v._1.fundeps.tag === "Just") { return rangeOf7(v._1.fundeps._1._2).end; }
              $runtime.fail();
            }
            if (v._2.tag === "Just") {
              const $1 = v._2._1._2.length - 1 | 0;
              return rangeOfType1.rangeOf((() => {
                if ($1 >= 0 && $1 < v._2._1._2.length) { return v._2._1._2[$1].value; }
                $runtime.fail();
              })()).end;
            }
            $runtime.fail();
          })()
        };
      }
      if (v.tag === "DeclInstanceChain") {
        const $1 = v._1.tail.length - 1 | 0;
        if ($1 >= 0 && $1 < v._1.tail.length) { return {start: $0.rangeOf(v._1.head).start, end: $0.rangeOf(v._1.tail[$1]._2).end}; }
        return $0.rangeOf(v._1.head);
      }
      if (v.tag === "DeclDerive") {
        const $1 = v._3.types.length - 1 | 0;
        return {start: v._1.range.start, end: $1 >= 0 && $1 < v._3.types.length ? rangeOfType1.rangeOf(v._3.types[$1]).end : v._3.className.token.range.end};
      }
      if (v.tag === "DeclKindSignature") { return {start: v._1.range.start, end: rangeOfType1.rangeOf(v._2.value).end}; }
      if (v.tag === "DeclSignature") { return {start: v._1.label.token.range.start, end: rangeOfType1.rangeOf(v._1.value).end}; }
      if (v.tag === "DeclValue") { return {start: v._1.name.token.range.start, end: rangeOfGuarded(dictRangeOf).rangeOf(v._1.guarded).end}; }
      if (v.tag === "DeclFixity") {
        return {
          start: v._1.keyword._1.range.start,
          end: (() => {
            if (v._1.operator.tag === "FixityValue") { return v._1.operator._3.token.range.end; }
            if (v._1.operator.tag === "FixityType") { return v._1.operator._4.token.range.end; }
            $runtime.fail();
          })()
        };
      }
      if (v.tag === "DeclForeign") { return {start: v._1.range.start, end: rangeOfForeign(dictRangeOf).rangeOf(v._3).end}; }
      if (v.tag === "DeclRole") {
        return {
          start: v._1.range.start,
          end: (() => {
            const $1 = v._4.length - 1 | 0;
            if ($1 >= 0 && $1 < v._4.length) { return v._4[$1]._1.range.end; }
            $runtime.fail();
          })()
        };
      }
      if (v.tag === "DeclError") { return dictRangeOf.rangeOf(v._1); }
      $runtime.fail();
    }
  };
};
export {
  foldMap,
  foldMap2,
  rangeOf,
  rangeOf7,
  rangeOfAppSpine,
  rangeOfBinder,
  rangeOfClassFundep,
  rangeOfDataCtor,
  rangeOfDataMembers,
  rangeOfDecl,
  rangeOfDoStatement,
  rangeOfExport,
  rangeOfExpr,
  rangeOfFixityOp,
  rangeOfForeign,
  rangeOfGuarded,
  rangeOfGuardedExpr,
  rangeOfImport,
  rangeOfImportDecl,
  rangeOfInstance,
  rangeOfInstanceBinding,
  rangeOfLabeled,
  rangeOfLetBinding,
  rangeOfModule,
  rangeOfName,
  rangeOfOneOrDelimited,
  rangeOfPrefixed,
  rangeOfQualifiedName,
  rangeOfRecoveredError,
  rangeOfSeparated,
  rangeOfType,
  rangeOfTypeVarBinding,
  rangeOfVoid,
  rangeOfWhere,
  rangeOfWrapped,
  tokensOf,
  tokensOf4,
  tokensOf6,
  tokensOf7,
  tokensOfAppSpine,
  tokensOfArray,
  tokensOfBinder,
  tokensOfClassFundep,
  tokensOfDataCtor,
  tokensOfDataMembers,
  tokensOfDecl,
  tokensOfDoStatement,
  tokensOfExport,
  tokensOfExpr,
  tokensOfFixityOp,
  tokensOfForeign,
  tokensOfGuarded,
  tokensOfGuardedExpr,
  tokensOfImport,
  tokensOfImportDecl,
  tokensOfInstance,
  tokensOfInstanceBinding,
  tokensOfLabeled,
  tokensOfLetBinding,
  tokensOfMaybe,
  tokensOfModule,
  tokensOfName,
  tokensOfNonEmptyArray,
  tokensOfOneOrDelimited,
  tokensOfPatternGuard,
  tokensOfPrefixed,
  tokensOfPrefixed1,
  tokensOfQualifiedName,
  tokensOfRecordLabeled,
  tokensOfRecordUpdate,
  tokensOfRecoveredError,
  tokensOfRow,
  tokensOfSeparated,
  tokensOfSeparated1,
  tokensOfTuple,
  tokensOfType,
  tokensOfTypeVarBinding,
  tokensOfVoid,
  tokensOfWhere,
  tokensOfWrapped
};
