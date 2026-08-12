import * as Control_Lazy from "../Control.Lazy/index.js";
import * as Data_Array from "../Data.Array/index.js";
import * as Data_Array_NonEmpty from "../Data.Array.NonEmpty/index.js";
import * as Data_Array_NonEmpty_Internal from "../Data.Array.NonEmpty.Internal/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Data_Void from "../Data.Void/index.js";
import * as PureScript_CST_Range_TokenList from "../PureScript.CST.Range.TokenList/index.js";
import * as PureScript_CST_Types from "../PureScript.CST.Types/index.js";
var tokensOfVoid = {
    tokensOf: Data_Void.absurd
};
var tokensOfRecoveredError = {
    tokensOf: function (v) {
        return PureScript_CST_Range_TokenList.fromArray(v.tokens);
    }
};
var tokensOfQualifiedName = {
    tokensOf: function (v) {
        return PureScript_CST_Range_TokenList.singleton(v.token);
    }
};
var tokensOfName = {
    tokensOf: function (v) {
        return PureScript_CST_Range_TokenList.singleton(v.token);
    }
};
var tokensOf = function (dict) {
    return dict.tokensOf;
};
var tokensOfArray = function (dictTokensOf) {
    return {
        tokensOf: Data_Foldable.foldMap(Data_Foldable.foldableArray)(PureScript_CST_Range_TokenList.monoidTokenList)(function (a) {
            return Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v) {
                return tokensOf(dictTokensOf)(a);
            });
        })
    };
};
var tokensOfFixityOp = {
    tokensOf: function (v) {
        if (v instanceof PureScript_CST_Types.FixityValue) {
            return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfQualifiedName)(v.value0))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(PureScript_CST_Range_TokenList.singleton(v.value1))(tokensOf(tokensOfName)(v.value2)));
        };
        if (v instanceof PureScript_CST_Types.FixityType) {
            return PureScript_CST_Range_TokenList.cons(v.value0)(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfQualifiedName)(v.value1))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(PureScript_CST_Range_TokenList.singleton(v.value2))(tokensOf(tokensOfName)(v.value3))));
        };
        throw new Error("Failed pattern match at PureScript.CST.Range (line 643, column 14 - line 647, column 59): " + [ v.constructor.name ]);
    }
};
var tokensOfLabeled = function (dictTokensOf) {
    return function (dictTokensOf1) {
        return {
            tokensOf: function (v) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(dictTokensOf)(v.label))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(PureScript_CST_Range_TokenList.singleton(v.separator))(tokensOf(dictTokensOf1)(v.value)));
            }
        };
    };
};
var tokensOfLabeled1 = /* #__PURE__ */ tokensOfLabeled(tokensOfName);
var tokensOfMaybe = function (dictTokensOf) {
    return {
        tokensOf: Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(PureScript_CST_Range_TokenList.monoidTokenList)(tokensOf(dictTokensOf))
    };
};
var tokensOfNonEmptyArray = function (dictTokensOf) {
    return {
        tokensOf: Data_Foldable.foldMap(Data_Array_NonEmpty_Internal.foldableNonEmptyArray)(PureScript_CST_Range_TokenList.monoidTokenList)(function (a) {
            return Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v) {
                return tokensOf(dictTokensOf)(a);
            });
        })
    };
};
var tokensOfNonEmptyArray1 = /* #__PURE__ */ tokensOfNonEmptyArray(tokensOfName);
var tokensOfClassFundep = {
    tokensOf: function (v) {
        if (v instanceof PureScript_CST_Types.FundepDetermined) {
            return PureScript_CST_Range_TokenList.cons(v.value0)(tokensOf(tokensOfNonEmptyArray1)(v.value1));
        };
        if (v instanceof PureScript_CST_Types.FundepDetermines) {
            return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfNonEmptyArray1)(v.value0))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(PureScript_CST_Range_TokenList.singleton(v.value1))(tokensOf(tokensOfNonEmptyArray1)(v.value2)));
        };
        throw new Error("Failed pattern match at PureScript.CST.Range (line 564, column 14 - line 568, column 50): " + [ v.constructor.name ]);
    }
};
var tokensOfPrefixed = function (dictTokensOf) {
    return {
        tokensOf: function (v) {
            if (v.prefix instanceof Data_Maybe.Just) {
                return PureScript_CST_Range_TokenList.cons(v.prefix.value0)(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return tokensOf(dictTokensOf)(v.value);
                }));
            };
            if (v.prefix instanceof Data_Maybe.Nothing) {
                return tokensOf(dictTokensOf)(v.value);
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 136, column 5 - line 140, column 23): " + [ v.prefix.constructor.name ]);
        }
    };
};
var tokensOfPrefixed1 = /* #__PURE__ */ tokensOfPrefixed(tokensOfName);
var tokensOfRecordLabeled = function (dictTokensOf) {
    return {
        tokensOf: function (v) {
            if (v instanceof PureScript_CST_Types.RecordPun) {
                return tokensOf(tokensOfName)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.RecordField) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfName)(v.value0))(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return PureScript_CST_Range_TokenList.cons(v.value1)(tokensOf(dictTokensOf)(v.value2));
                }));
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 1012, column 14 - line 1016, column 52): " + [ v.constructor.name ]);
        }
    };
};
var tokensOfSeparated = function (dictTokensOf) {
    return {
        tokensOf: function (v) {
            return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(dictTokensOf)(v.head))(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                return Data_Foldable.foldMap(Data_Foldable.foldableArray)(PureScript_CST_Range_TokenList.monoidTokenList)(function (v2) {
                    return PureScript_CST_Range_TokenList.cons(v2.value0)(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v3) {
                        return tokensOf(dictTokensOf)(v2.value1);
                    }));
                })(v.tail);
            }));
        }
    };
};
var tokensOfSeparated1 = /* #__PURE__ */ tokensOfSeparated(tokensOfName);
var tokensOfSeparated2 = /* #__PURE__ */ tokensOfSeparated(tokensOfClassFundep);
var tokensOfTuple = function (dictTokensOf) {
    return function (dictTokensOf1) {
        return {
            tokensOf: function (v) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(dictTokensOf)(v.value0))(tokensOf(dictTokensOf1)(v.value1));
            }
        };
    };
};
var tokensOfTuple1 = /* #__PURE__ */ tokensOfTuple(tokensOfQualifiedName);
var tokensOfWrapped = function (dictTokensOf) {
    return {
        tokensOf: function (v) {
            return PureScript_CST_Range_TokenList.wrap(v.open)(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                return tokensOf(dictTokensOf)(v.value);
            }))(v.close);
        }
    };
};
var tokensOfWrapped1 = /* #__PURE__ */ tokensOfWrapped(/* #__PURE__ */ tokensOfMaybe(/* #__PURE__ */ tokensOfSeparated(tokensOfName)));
var tokensOfDataMembers = {
    tokensOf: function (v) {
        if (v instanceof PureScript_CST_Types.DataAll) {
            return PureScript_CST_Range_TokenList.singleton(v.value0);
        };
        if (v instanceof PureScript_CST_Types.DataEnumerated) {
            return tokensOf(tokensOfWrapped1)(v.value0);
        };
        throw new Error("Failed pattern match at PureScript.CST.Range (line 329, column 14 - line 333, column 17): " + [ v.constructor.name ]);
    }
};
var tokensOf1 = /* #__PURE__ */ tokensOf(tokensOfDataMembers);
var tokensOfExport = function (dictTokensOf) {
    return {
        tokensOf: function (v) {
            if (v instanceof PureScript_CST_Types.ExportValue) {
                return tokensOf(tokensOfName)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.ExportOp) {
                return tokensOf(tokensOfName)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.ExportType) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfName)(v.value0))(Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(PureScript_CST_Range_TokenList.monoidTokenList)(tokensOf1)(v.value1));
            };
            if (v instanceof PureScript_CST_Types.ExportTypeOp) {
                return PureScript_CST_Range_TokenList.cons(v.value0)(tokensOf(tokensOfName)(v.value1));
            };
            if (v instanceof PureScript_CST_Types.ExportClass) {
                return PureScript_CST_Range_TokenList.cons(v.value0)(tokensOf(tokensOfName)(v.value1));
            };
            if (v instanceof PureScript_CST_Types.ExportModule) {
                return PureScript_CST_Range_TokenList.cons(v.value0)(tokensOf(tokensOfName)(v.value1));
            };
            if (v instanceof PureScript_CST_Types.ExportError) {
                return tokensOf(dictTokensOf)(v.value0);
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 305, column 14 - line 319, column 17): " + [ v.constructor.name ]);
        }
    };
};
var tokensOfImport = function (dictTokensOf) {
    return {
        tokensOf: function (v) {
            if (v instanceof PureScript_CST_Types.ImportValue) {
                return tokensOf(tokensOfName)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.ImportOp) {
                return tokensOf(tokensOfName)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.ImportType) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfName)(v.value0))(Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(PureScript_CST_Range_TokenList.monoidTokenList)(tokensOf1)(v.value1));
            };
            if (v instanceof PureScript_CST_Types.ImportTypeOp) {
                return PureScript_CST_Range_TokenList.cons(v.value0)(tokensOf(tokensOfName)(v.value1));
            };
            if (v instanceof PureScript_CST_Types.ImportClass) {
                return PureScript_CST_Range_TokenList.cons(v.value0)(tokensOf(tokensOfName)(v.value1));
            };
            if (v instanceof PureScript_CST_Types.ImportError) {
                return tokensOf(dictTokensOf)(v.value0);
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 384, column 14 - line 396, column 17): " + [ v.constructor.name ]);
        }
    };
};
var tokensOfImportDecl = function (dictTokensOf) {
    var tokensOfWrapped2 = tokensOfWrapped(tokensOfSeparated(tokensOfImport(dictTokensOf)));
    return {
        tokensOf: function (v) {
            return PureScript_CST_Range_TokenList.cons(v.keyword)(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfName)(v.module))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(PureScript_CST_Range_TokenList.monoidTokenList)(function (v2) {
                    return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(PureScript_CST_Range_TokenList.monoidTokenList)(PureScript_CST_Range_TokenList.singleton)(v2.value0))(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v3) {
                        return tokensOf(tokensOfWrapped2)(v2.value1);
                    }));
                })(v.names))(Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(PureScript_CST_Range_TokenList.monoidTokenList)(function (v2) {
                    return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(PureScript_CST_Range_TokenList.singleton(v2.value0))(tokensOf(tokensOfName)(v2.value1));
                })(v.qualified)));
            }));
        }
    };
};
var tokensOfOneOrDelimited = function (dictTokensOf) {
    var tokensOfWrapped2 = tokensOfWrapped(tokensOfSeparated(dictTokensOf));
    return {
        tokensOf: function (v) {
            if (v instanceof PureScript_CST_Types.One) {
                return tokensOf(dictTokensOf)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.Many) {
                return tokensOf(tokensOfWrapped2)(v.value0);
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 148, column 14 - line 150, column 27): " + [ v.constructor.name ]);
        }
    };
};
var tokensOfTypeVarBinding = function (dictTokensOf) {
    var tokensOfLabeled2 = tokensOfLabeled(dictTokensOf);
    return function (dictTokensOf1) {
        return {
            tokensOf: function (v) {
                if (v instanceof PureScript_CST_Types.TypeVarKinded) {
                    return tokensOf(tokensOfWrapped(tokensOfLabeled2(tokensOfType(dictTokensOf1))))(v.value0);
                };
                if (v instanceof PureScript_CST_Types.TypeVarName) {
                    return tokensOf(dictTokensOf)(v.value0);
                };
                throw new Error("Failed pattern match at PureScript.CST.Range (line 269, column 14 - line 273, column 17): " + [ v.constructor.name ]);
            }
        };
    };
};
var tokensOfType = function (dictTokensOf) {
    return {
        tokensOf: function (v) {
            if (v instanceof PureScript_CST_Types.TypeVar) {
                return tokensOf(tokensOfName)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.TypeConstructor) {
                return tokensOf(tokensOfQualifiedName)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.TypeWildcard) {
                return PureScript_CST_Range_TokenList.singleton(v.value0);
            };
            if (v instanceof PureScript_CST_Types.TypeHole) {
                return tokensOf(tokensOfName)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.TypeString) {
                return PureScript_CST_Range_TokenList.singleton(v.value0);
            };
            if (v instanceof PureScript_CST_Types.TypeInt) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(PureScript_CST_Range_TokenList.monoidTokenList)(PureScript_CST_Range_TokenList.singleton)(v.value0))(PureScript_CST_Range_TokenList.singleton(v.value1));
            };
            if (v instanceof PureScript_CST_Types.TypeRow) {
                return tokensOf(tokensOfWrapped(tokensOfRow(dictTokensOf)))(v.value0);
            };
            if (v instanceof PureScript_CST_Types.TypeRecord) {
                return tokensOf(tokensOfWrapped(tokensOfRow(dictTokensOf)))(v.value0);
            };
            if (v instanceof PureScript_CST_Types.TypeForall) {
                return PureScript_CST_Range_TokenList.cons(v.value0)(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfNonEmptyArray(tokensOfTypeVarBinding(tokensOfPrefixed1)(dictTokensOf)))(v.value1))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(PureScript_CST_Range_TokenList.singleton(v.value2))(tokensOf(tokensOfType(dictTokensOf))(v.value3)));
                }));
            };
            if (v instanceof PureScript_CST_Types.TypeKinded) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfType(dictTokensOf))(v.value0))(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(PureScript_CST_Range_TokenList.singleton(v.value1))(tokensOf(tokensOfType(dictTokensOf))(v.value2));
                }));
            };
            if (v instanceof PureScript_CST_Types.TypeApp) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfType(dictTokensOf))(v.value0))(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return tokensOf(tokensOfNonEmptyArray(tokensOfType(dictTokensOf)))(v.value1);
                }));
            };
            if (v instanceof PureScript_CST_Types.TypeOp) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfType(dictTokensOf))(v.value0))(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return Data_Foldable.foldMap(Data_Array_NonEmpty_Internal.foldableNonEmptyArray)(PureScript_CST_Range_TokenList.monoidTokenList)(function (v2) {
                        return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfQualifiedName)(v2.value0))(tokensOf(tokensOfType(dictTokensOf))(v2.value1));
                    })(v.value1);
                }));
            };
            if (v instanceof PureScript_CST_Types.TypeOpName) {
                return tokensOf(tokensOfQualifiedName)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.TypeArrow) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfType(dictTokensOf))(v.value0))(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(PureScript_CST_Range_TokenList.singleton(v.value1))(tokensOf(tokensOfType(dictTokensOf))(v.value2));
                }));
            };
            if (v instanceof PureScript_CST_Types.TypeArrowName) {
                return PureScript_CST_Range_TokenList.singleton(v.value0);
            };
            if (v instanceof PureScript_CST_Types.TypeConstrained) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfType(dictTokensOf))(v.value0))(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(PureScript_CST_Range_TokenList.singleton(v.value1))(tokensOf(tokensOfType(dictTokensOf))(v.value2));
                }));
            };
            if (v instanceof PureScript_CST_Types.TypeParens) {
                return tokensOf(tokensOfWrapped(tokensOfType(dictTokensOf)))(v.value0);
            };
            if (v instanceof PureScript_CST_Types["TypeError"]) {
                return tokensOf(dictTokensOf)(v.value0);
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 210, column 14 - line 254, column 17): " + [ v.constructor.name ]);
        }
    };
};
var tokensOfRow = function (dictTokensOf) {
    return {
        tokensOf: function (v) {
            return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(PureScript_CST_Range_TokenList.monoidTokenList)(tokensOf(tokensOfSeparated(tokensOfLabeled1(tokensOfType(dictTokensOf)))))(v.labels))(Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(PureScript_CST_Range_TokenList.monoidTokenList)(function (v1) {
                return PureScript_CST_Range_TokenList.cons(v1.value0)(tokensOf(tokensOfType(dictTokensOf))(v1.value1));
            })(v.tail));
        }
    };
};
var tokensOfTypeVarBinding1 = /* #__PURE__ */ tokensOfTypeVarBinding(tokensOfName);
var tokensOfAppSpine = function (dictTokensOf) {
    var tokensOfType1 = tokensOfType(dictTokensOf);
    return function (dictTokensOf1) {
        return {
            tokensOf: function (v) {
                if (v instanceof PureScript_CST_Types.AppType) {
                    return PureScript_CST_Range_TokenList.cons(v.value0)(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                        return tokensOf(tokensOfType1)(v.value1);
                    }));
                };
                if (v instanceof PureScript_CST_Types.AppTerm) {
                    return tokensOf(dictTokensOf1)(v.value0);
                };
                throw new Error("Failed pattern match at PureScript.CST.Range (line 856, column 14 - line 860, column 17): " + [ v.constructor.name ]);
            }
        };
    };
};
var tokensOfBinder = function (dictTokensOf) {
    var tokensOfType1 = tokensOfType(dictTokensOf);
    return {
        tokensOf: function (v) {
            if (v instanceof PureScript_CST_Types.BinderWildcard) {
                return PureScript_CST_Range_TokenList.singleton(v.value0);
            };
            if (v instanceof PureScript_CST_Types.BinderVar) {
                return tokensOf(tokensOfName)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.BinderNamed) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfName)(v.value0))(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return PureScript_CST_Range_TokenList.cons(v.value1)(tokensOf(tokensOfBinder(dictTokensOf))(v.value2));
                }));
            };
            if (v instanceof PureScript_CST_Types.BinderConstructor) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfQualifiedName)(v.value0))(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return tokensOf(tokensOfArray(tokensOfBinder(dictTokensOf)))(v.value1);
                }));
            };
            if (v instanceof PureScript_CST_Types.BinderBoolean) {
                return PureScript_CST_Range_TokenList.singleton(v.value0);
            };
            if (v instanceof PureScript_CST_Types.BinderChar) {
                return PureScript_CST_Range_TokenList.singleton(v.value0);
            };
            if (v instanceof PureScript_CST_Types.BinderString) {
                return PureScript_CST_Range_TokenList.singleton(v.value0);
            };
            if (v instanceof PureScript_CST_Types.BinderInt) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(PureScript_CST_Range_TokenList.monoidTokenList)(PureScript_CST_Range_TokenList.singleton)(v.value0))(PureScript_CST_Range_TokenList.singleton(v.value1));
            };
            if (v instanceof PureScript_CST_Types.BinderNumber) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(PureScript_CST_Range_TokenList.monoidTokenList)(PureScript_CST_Range_TokenList.singleton)(v.value0))(PureScript_CST_Range_TokenList.singleton(v.value1));
            };
            if (v instanceof PureScript_CST_Types.BinderArray) {
                return tokensOf(tokensOfWrapped(tokensOfMaybe(tokensOfSeparated(tokensOfBinder(dictTokensOf)))))(v.value0);
            };
            if (v instanceof PureScript_CST_Types.BinderRecord) {
                return tokensOf(tokensOfWrapped(tokensOfMaybe(tokensOfSeparated(tokensOfRecordLabeled(tokensOfBinder(dictTokensOf))))))(v.value0);
            };
            if (v instanceof PureScript_CST_Types.BinderParens) {
                return tokensOf(tokensOfWrapped(tokensOfBinder(dictTokensOf)))(v.value0);
            };
            if (v instanceof PureScript_CST_Types.BinderTyped) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfBinder(dictTokensOf))(v.value0))(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return PureScript_CST_Range_TokenList.cons(v.value1)(tokensOf(tokensOfType1)(v.value2));
                }));
            };
            if (v instanceof PureScript_CST_Types.BinderOp) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfBinder(dictTokensOf))(v.value0))(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return tokensOf(tokensOfNonEmptyArray(tokensOfTuple1(tokensOfBinder(dictTokensOf))))(v.value1);
                }));
            };
            if (v instanceof PureScript_CST_Types.BinderError) {
                return tokensOf(dictTokensOf)(v.value0);
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 979, column 14 - line 1009, column 17): " + [ v.constructor.name ]);
        }
    };
};
var tokensOfDataCtor = function (dictTokensOf) {
    var tokensOfArray1 = tokensOfArray(tokensOfType(dictTokensOf));
    return {
        tokensOf: function (v) {
            return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfName)(v.name))(tokensOf(tokensOfArray1)(v.fields));
        }
    };
};
var tokensOfForeign = function (dictTokensOf) {
    var tokensOfLabeled2 = tokensOfLabeled1(tokensOfType(dictTokensOf));
    return {
        tokensOf: function (v) {
            if (v instanceof PureScript_CST_Types.ForeignValue) {
                return tokensOf(tokensOfLabeled2)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.ForeignData) {
                return PureScript_CST_Range_TokenList.cons(v.value0)(tokensOf(tokensOfLabeled2)(v.value1));
            };
            if (v instanceof PureScript_CST_Types.ForeignKind) {
                return PureScript_CST_Range_TokenList.cons(v.value0)(tokensOf(tokensOfName)(v.value1));
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 663, column 14 - line 669, column 26): " + [ v.constructor.name ]);
        }
    };
};
var tokensOfWhere = function (dictTokensOf) {
    return {
        tokensOf: function (v) {
            return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfExpr(dictTokensOf))(v.expr))(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                return Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(PureScript_CST_Range_TokenList.monoidTokenList)(function (v2) {
                    return PureScript_CST_Range_TokenList.cons(v2.value0)(tokensOf(tokensOfNonEmptyArray(tokensOfLetBinding(dictTokensOf)))(v2.value1));
                })(v.bindings);
            }));
        }
    };
};
var tokensOfRecordUpdate = function (dictTokensOf) {
    return {
        tokensOf: function (v) {
            if (v instanceof PureScript_CST_Types.RecordUpdateLeaf) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfName)(v.value0))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(PureScript_CST_Range_TokenList.singleton(v.value1))(tokensOf(tokensOfExpr(dictTokensOf))(v.value2)));
            };
            if (v instanceof PureScript_CST_Types.RecordUpdateBranch) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfName)(v.value0))(tokensOf(tokensOfWrapped(tokensOfSeparated(tokensOfRecordUpdate(dictTokensOf))))(v.value1));
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 863, column 14 - line 867, column 32): " + [ v.constructor.name ]);
        }
    };
};
var tokensOfPatternGuard = function (dictTokensOf) {
    var tokensOfBinder1 = tokensOfBinder(dictTokensOf);
    return {
        tokensOf: function (v) {
            return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(PureScript_CST_Range_TokenList.monoidTokenList)(function (v1) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfBinder1)(v1.value0))(PureScript_CST_Range_TokenList.singleton(v1.value1));
            })(v.binder))(tokensOf(tokensOfExpr(dictTokensOf))(v.expr));
        }
    };
};
var tokensOfLetBinding = function (dictTokensOf) {
    var tokensOfLabeled2 = tokensOfLabeled1(tokensOfType(dictTokensOf));
    var tokensOfArray1 = tokensOfArray(tokensOfBinder(dictTokensOf));
    var tokensOfBinder1 = tokensOfBinder(dictTokensOf);
    return {
        tokensOf: function (v) {
            if (v instanceof PureScript_CST_Types.LetBindingSignature) {
                return tokensOf(tokensOfLabeled2)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.LetBindingName) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfName)(v.value0.name))(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfArray1)(v.value0.binders))(tokensOf(tokensOfGuarded(dictTokensOf))(v.value0.guarded));
                }));
            };
            if (v instanceof PureScript_CST_Types.LetBindingPattern) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfBinder1)(v.value0))(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return PureScript_CST_Range_TokenList.cons(v.value1)(tokensOf(tokensOfWhere(dictTokensOf))(v.value2));
                }));
            };
            if (v instanceof PureScript_CST_Types.LetBindingError) {
                return tokensOf(dictTokensOf)(v.value0);
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 911, column 14 - line 919, column 17): " + [ v.constructor.name ]);
        }
    };
};
var tokensOfGuardedExpr = function (dictTokensOf) {
    return {
        tokensOf: function (v) {
            return PureScript_CST_Range_TokenList.cons(v.bar)(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfSeparated(tokensOfPatternGuard(dictTokensOf)))(v.patterns))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(PureScript_CST_Range_TokenList.singleton(v.separator))(tokensOf(tokensOfWhere(dictTokensOf))(v.where)));
            }));
        }
    };
};
var tokensOfGuarded = function (dictTokensOf) {
    return {
        tokensOf: function (v) {
            if (v instanceof PureScript_CST_Types.Unconditional) {
                return PureScript_CST_Range_TokenList.cons(v.value0)(tokensOf(tokensOfWhere(dictTokensOf))(v.value1));
            };
            if (v instanceof PureScript_CST_Types.Guarded) {
                return tokensOf(tokensOfNonEmptyArray(tokensOfGuardedExpr(dictTokensOf)))(v.value0);
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 607, column 14 - line 611, column 18): " + [ v.constructor.name ]);
        }
    };
};
var tokensOfExpr = function (dictTokensOf) {
    var tokensOfType1 = tokensOfType(dictTokensOf);
    var tokensOfAppSpine1 = tokensOfAppSpine(dictTokensOf);
    var tokensOfBinder1 = tokensOfBinder(dictTokensOf);
    var tokensOfNonEmptyArray2 = tokensOfNonEmptyArray(tokensOfBinder1);
    var tokensOfTuple2 = tokensOfTuple(tokensOfSeparated(tokensOfBinder1));
    return {
        tokensOf: function (v) {
            if (v instanceof PureScript_CST_Types.ExprHole) {
                return tokensOf(tokensOfName)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.ExprSection) {
                return PureScript_CST_Range_TokenList.singleton(v.value0);
            };
            if (v instanceof PureScript_CST_Types.ExprIdent) {
                return tokensOf(tokensOfQualifiedName)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.ExprConstructor) {
                return tokensOf(tokensOfQualifiedName)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.ExprBoolean) {
                return PureScript_CST_Range_TokenList.singleton(v.value0);
            };
            if (v instanceof PureScript_CST_Types.ExprChar) {
                return PureScript_CST_Range_TokenList.singleton(v.value0);
            };
            if (v instanceof PureScript_CST_Types.ExprString) {
                return PureScript_CST_Range_TokenList.singleton(v.value0);
            };
            if (v instanceof PureScript_CST_Types.ExprInt) {
                return PureScript_CST_Range_TokenList.singleton(v.value0);
            };
            if (v instanceof PureScript_CST_Types.ExprNumber) {
                return PureScript_CST_Range_TokenList.singleton(v.value0);
            };
            if (v instanceof PureScript_CST_Types.ExprArray) {
                return tokensOf(tokensOfWrapped(tokensOfMaybe(tokensOfSeparated(tokensOfExpr(dictTokensOf)))))(v.value0);
            };
            if (v instanceof PureScript_CST_Types.ExprRecord) {
                return tokensOf(tokensOfWrapped(tokensOfMaybe(tokensOfSeparated(tokensOfRecordLabeled(tokensOfExpr(dictTokensOf))))))(v.value0);
            };
            if (v instanceof PureScript_CST_Types.ExprParens) {
                return tokensOf(tokensOfWrapped(tokensOfExpr(dictTokensOf)))(v.value0);
            };
            if (v instanceof PureScript_CST_Types.ExprTyped) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfExpr(dictTokensOf))(v.value0))(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return PureScript_CST_Range_TokenList.cons(v.value1)(tokensOf(tokensOfType1)(v.value2));
                }));
            };
            if (v instanceof PureScript_CST_Types.ExprInfix) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfExpr(dictTokensOf))(v.value0))(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return tokensOf(tokensOfNonEmptyArray(tokensOfTuple(tokensOfWrapped(tokensOfExpr(dictTokensOf)))(tokensOfExpr(dictTokensOf))))(v.value1);
                }));
            };
            if (v instanceof PureScript_CST_Types.ExprOp) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfExpr(dictTokensOf))(v.value0))(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return tokensOf(tokensOfNonEmptyArray(tokensOfTuple1(tokensOfExpr(dictTokensOf))))(v.value1);
                }));
            };
            if (v instanceof PureScript_CST_Types.ExprOpName) {
                return tokensOf(tokensOfQualifiedName)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.ExprNegate) {
                return PureScript_CST_Range_TokenList.cons(v.value0)(tokensOf(tokensOfExpr(dictTokensOf))(v.value1));
            };
            if (v instanceof PureScript_CST_Types.ExprRecordAccessor) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfExpr(dictTokensOf))(v.value0.expr))(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return PureScript_CST_Range_TokenList.cons(v.value0.dot)(tokensOf(tokensOfSeparated1)(v.value0.path));
                }));
            };
            if (v instanceof PureScript_CST_Types.ExprRecordUpdate) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfExpr(dictTokensOf))(v.value0))(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return tokensOf(tokensOfWrapped(tokensOfSeparated(tokensOfRecordUpdate(dictTokensOf))))(v.value1);
                }));
            };
            if (v instanceof PureScript_CST_Types.ExprApp) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfExpr(dictTokensOf))(v.value0))(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return tokensOf(tokensOfNonEmptyArray(tokensOfAppSpine1(tokensOfExpr(dictTokensOf))))(v.value1);
                }));
            };
            if (v instanceof PureScript_CST_Types.ExprLambda) {
                return PureScript_CST_Range_TokenList.cons(v.value0.symbol)(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfNonEmptyArray2)(v.value0.binders))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(PureScript_CST_Range_TokenList.singleton(v.value0.arrow))(tokensOf(tokensOfExpr(dictTokensOf))(v.value0.body)));
                }));
            };
            if (v instanceof PureScript_CST_Types.ExprIf) {
                return PureScript_CST_Range_TokenList.cons(v.value0.keyword)(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfExpr(dictTokensOf))(v.value0.cond))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(PureScript_CST_Range_TokenList.singleton(v.value0.then))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfExpr(dictTokensOf))(v["value0"]["true"]))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(PureScript_CST_Range_TokenList.singleton(v["value0"]["else"]))(tokensOf(tokensOfExpr(dictTokensOf))(v["value0"]["false"])))));
                }));
            };
            if (v instanceof PureScript_CST_Types.ExprCase) {
                return PureScript_CST_Range_TokenList.cons(v.value0.keyword)(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfSeparated(tokensOfExpr(dictTokensOf)))(v.value0.head))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(PureScript_CST_Range_TokenList.singleton(v.value0.of))(tokensOf(tokensOfNonEmptyArray(tokensOfTuple2(tokensOfGuarded(dictTokensOf))))(v.value0.branches)));
                }));
            };
            if (v instanceof PureScript_CST_Types.ExprLet) {
                return PureScript_CST_Range_TokenList.cons(v.value0.keyword)(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfNonEmptyArray(tokensOfLetBinding(dictTokensOf)))(v.value0.bindings))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(PureScript_CST_Range_TokenList.singleton(v["value0"]["in"]))(tokensOf(tokensOfExpr(dictTokensOf))(v.value0.body)));
                }));
            };
            if (v instanceof PureScript_CST_Types.ExprDo) {
                return PureScript_CST_Range_TokenList.cons(v.value0.keyword)(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return tokensOf(tokensOfNonEmptyArray(tokensOfDoStatement(dictTokensOf)))(v.value0.statements);
                }));
            };
            if (v instanceof PureScript_CST_Types.ExprAdo) {
                return PureScript_CST_Range_TokenList.cons(v.value0.keyword)(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfArray(tokensOfDoStatement(dictTokensOf)))(v.value0.statements))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(PureScript_CST_Range_TokenList.singleton(v["value0"]["in"]))(tokensOf(tokensOfExpr(dictTokensOf))(v.value0.result)));
                }));
            };
            if (v instanceof PureScript_CST_Types.ExprError) {
                return tokensOf(dictTokensOf)(v.value0);
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 773, column 14 - line 844, column 17): " + [ v.constructor.name ]);
        }
    };
};
var tokensOfDoStatement = function (dictTokensOf) {
    var tokensOfBinder1 = tokensOfBinder(dictTokensOf);
    return {
        tokensOf: function (v) {
            if (v instanceof PureScript_CST_Types.DoLet) {
                return PureScript_CST_Range_TokenList.cons(v.value0)(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return tokensOf(tokensOfNonEmptyArray(tokensOfLetBinding(dictTokensOf)))(v.value1);
                }));
            };
            if (v instanceof PureScript_CST_Types.DoDiscard) {
                return tokensOf(tokensOfExpr(dictTokensOf))(v.value0);
            };
            if (v instanceof PureScript_CST_Types.DoBind) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfBinder1)(v.value0))(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return PureScript_CST_Range_TokenList.cons(v.value1)(tokensOf(tokensOfExpr(dictTokensOf))(v.value2));
                }));
            };
            if (v instanceof PureScript_CST_Types.DoError) {
                return tokensOf(dictTokensOf)(v.value0);
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 885, column 14 - line 893, column 17): " + [ v.constructor.name ]);
        }
    };
};
var tokensOfInstanceBinding = function (dictTokensOf) {
    var tokensOfLabeled2 = tokensOfLabeled1(tokensOfType(dictTokensOf));
    var tokensOfArray1 = tokensOfArray(tokensOfBinder(dictTokensOf));
    var tokensOfGuarded1 = tokensOfGuarded(dictTokensOf);
    return {
        tokensOf: function (v) {
            if (v instanceof PureScript_CST_Types.InstanceBindingSignature) {
                return tokensOf(tokensOfLabeled2)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.InstanceBindingName) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfName)(v.value0.name))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfArray1)(v.value0.binders))(tokensOf(tokensOfGuarded1)(v.value0.guarded)));
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 681, column 14 - line 687, column 28): " + [ v.constructor.name ]);
        }
    };
};
var tokensOfInstance = function (dictTokensOf) {
    var tokensOfType1 = tokensOfType(dictTokensOf);
    var tokensOfOneOrDelimited1 = tokensOfOneOrDelimited(tokensOfType1);
    var tokensOfArray1 = tokensOfArray(tokensOfType1);
    var tokensOfNonEmptyArray2 = tokensOfNonEmptyArray(tokensOfInstanceBinding(dictTokensOf));
    return {
        tokensOf: function (v) {
            return PureScript_CST_Range_TokenList.cons(v.head.keyword)(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(PureScript_CST_Range_TokenList.monoidTokenList)(function (v2) {
                    return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfName)(v2.value0))(PureScript_CST_Range_TokenList.singleton(v2.value1));
                })(v.head.name))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(PureScript_CST_Range_TokenList.monoidTokenList)(function (v2) {
                    return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfOneOrDelimited1)(v2.value0))(PureScript_CST_Range_TokenList.singleton(v2.value1));
                })(v.head.constraints))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfQualifiedName)(v.head.className))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfArray1)(v.head.types))(Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(PureScript_CST_Range_TokenList.monoidTokenList)(function (v2) {
                    return PureScript_CST_Range_TokenList.cons(v2.value0)(tokensOf(tokensOfNonEmptyArray2)(v2.value1));
                })(v.body)))));
            }));
        }
    };
};
var tokensOfDecl = function (dictTokensOf) {
    var tokensOfArray1 = tokensOfArray(tokensOfTypeVarBinding1(dictTokensOf));
    var tokensOfSeparated3 = tokensOfSeparated(tokensOfDataCtor(dictTokensOf));
    var tokensOfType1 = tokensOfType(dictTokensOf);
    var tokensOfType2 = tokensOfType(dictTokensOf);
    var tokensOfOneOrDelimited1 = tokensOfOneOrDelimited(tokensOfType2);
    var tokensOfNonEmptyArray2 = tokensOfNonEmptyArray(tokensOfLabeled1(tokensOfType2));
    var tokensOfSeparated4 = tokensOfSeparated(tokensOfInstance(dictTokensOf));
    var tokensOfArray2 = tokensOfArray(tokensOfType2);
    var tokensOfLabeled2 = tokensOfLabeled1(tokensOfType2);
    var tokensOfArray3 = tokensOfArray(tokensOfBinder(dictTokensOf));
    var tokensOfGuarded1 = tokensOfGuarded(dictTokensOf);
    var tokensOfForeign1 = tokensOfForeign(dictTokensOf);
    return {
        tokensOf: function (v) {
            if (v instanceof PureScript_CST_Types.DeclData) {
                return PureScript_CST_Range_TokenList.cons(v.value0.keyword)(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfName)(v.value0.name))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfArray1)(v.value0.vars))(Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(PureScript_CST_Range_TokenList.monoidTokenList)(function (v2) {
                        return PureScript_CST_Range_TokenList.cons(v2.value0)(tokensOf(tokensOfSeparated3)(v2.value1));
                    })(v.value1)));
                }));
            };
            if (v instanceof PureScript_CST_Types.DeclType) {
                return PureScript_CST_Range_TokenList.cons(v.value0.keyword)(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfName)(v.value0.name))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfArray1)(v.value0.vars))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(PureScript_CST_Range_TokenList.singleton(v.value1))(tokensOf(tokensOfType1)(v.value2))));
                }));
            };
            if (v instanceof PureScript_CST_Types.DeclNewtype) {
                return PureScript_CST_Range_TokenList.cons(v.value0.keyword)(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfName)(v.value0.name))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfArray1)(v.value0.vars))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(PureScript_CST_Range_TokenList.singleton(v.value1))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfName)(v.value2))(tokensOf(tokensOfType1)(v.value3)))));
                }));
            };
            if (v instanceof PureScript_CST_Types.DeclClass) {
                return PureScript_CST_Range_TokenList.cons(v.value0.keyword)(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(PureScript_CST_Range_TokenList.monoidTokenList)(function (v2) {
                        return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfOneOrDelimited1)(v2.value0))(PureScript_CST_Range_TokenList.singleton(v2.value1));
                    })(v["value0"]["super"]))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfName)(v.value0.name))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfArray1)(v.value0.vars))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(PureScript_CST_Range_TokenList.monoidTokenList)(function (v2) {
                        return PureScript_CST_Range_TokenList.cons(v2.value0)(tokensOf(tokensOfSeparated2)(v2.value1));
                    })(v.value0.fundeps))(Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(PureScript_CST_Range_TokenList.monoidTokenList)(function (v2) {
                        return PureScript_CST_Range_TokenList.cons(v2.value0)(tokensOf(tokensOfNonEmptyArray2)(v2.value1));
                    })(v.value1)))));
                }));
            };
            if (v instanceof PureScript_CST_Types.DeclInstanceChain) {
                return tokensOf(tokensOfSeparated4)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.DeclDerive) {
                return PureScript_CST_Range_TokenList.cons(v.value0)(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(PureScript_CST_Range_TokenList.monoidTokenList)(PureScript_CST_Range_TokenList.singleton)(v.value1))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(PureScript_CST_Range_TokenList.singleton(v.value2.keyword))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(PureScript_CST_Range_TokenList.monoidTokenList)(function (v2) {
                        return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfName)(v2.value0))(PureScript_CST_Range_TokenList.singleton(v2.value1));
                    })(v.value2.name))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(PureScript_CST_Range_TokenList.monoidTokenList)(function (v2) {
                        return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfOneOrDelimited1)(v2.value0))(PureScript_CST_Range_TokenList.singleton(v2.value1));
                    })(v.value2.constraints))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfQualifiedName)(v.value2.className))(tokensOf(tokensOfArray2)(v.value2.types))))));
                }));
            };
            if (v instanceof PureScript_CST_Types.DeclKindSignature) {
                return PureScript_CST_Range_TokenList.cons(v.value0)(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return tokensOf(tokensOfLabeled2)(v.value1);
                }));
            };
            if (v instanceof PureScript_CST_Types.DeclSignature) {
                return tokensOf(tokensOfLabeled2)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.DeclValue) {
                return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfName)(v.value0.name))(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfArray3)(v.value0.binders))(tokensOf(tokensOfGuarded1)(v.value0.guarded));
                }));
            };
            if (v instanceof PureScript_CST_Types.DeclFixity) {
                return PureScript_CST_Range_TokenList.cons(v.value0.keyword.value0)(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return PureScript_CST_Range_TokenList.cons(v.value0.prec.value0)(tokensOf(tokensOfFixityOp)(v.value0.operator));
                }));
            };
            if (v instanceof PureScript_CST_Types.DeclForeign) {
                return PureScript_CST_Range_TokenList.cons(v.value0)(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return PureScript_CST_Range_TokenList.cons(v.value1)(tokensOf(tokensOfForeign1)(v.value2));
                }));
            };
            if (v instanceof PureScript_CST_Types.DeclRole) {
                return PureScript_CST_Range_TokenList.cons(v.value0)(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                    return Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(PureScript_CST_Range_TokenList.singleton(v.value1))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfName)(v.value2))(Data_Foldable.foldMap(Data_Array_NonEmpty_Internal.foldableNonEmptyArray)(PureScript_CST_Range_TokenList.monoidTokenList)(function (v2) {
                        return PureScript_CST_Range_TokenList.singleton(v2.value0);
                    })(v.value3)));
                }));
            };
            if (v instanceof PureScript_CST_Types.DeclError) {
                return tokensOf(dictTokensOf)(v.value0);
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 494, column 14 - line 550, column 17): " + [ v.constructor.name ]);
        }
    };
};
var tokensOfModule = function (dictTokensOf) {
    var tokensOf2 = tokensOf(tokensOfWrapped(tokensOfSeparated(tokensOfExport(dictTokensOf))));
    var tokensOf3 = tokensOf(tokensOfImportDecl(dictTokensOf));
    var tokensOf4 = tokensOf(tokensOfDecl(dictTokensOf));
    return {
        tokensOf: function (v) {
            return PureScript_CST_Range_TokenList.cons(v.header.keyword)(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(tokensOf(tokensOfName)(v.header.name))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                return Data_Foldable.foldMap(Data_Foldable.foldableMaybe)(PureScript_CST_Range_TokenList.monoidTokenList)(tokensOf2)(v.header.exports);
            }))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(PureScript_CST_Range_TokenList.singleton(v.header.where))(Data_Semigroup.append(PureScript_CST_Range_TokenList.semigroupTokenList)(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                return Data_Foldable.foldMap(Data_Foldable.foldableArray)(PureScript_CST_Range_TokenList.monoidTokenList)(tokensOf3)(v.header.imports);
            }))(Control_Lazy.defer(PureScript_CST_Range_TokenList.lazyTokenList)(function (v1) {
                return Data_Foldable.foldMap(Data_Foldable.foldableArray)(PureScript_CST_Range_TokenList.monoidTokenList)(tokensOf4)(v.body.decls);
            }))))));
        }
    };
};
var rangeOfWrapped = {
    rangeOf: function (v) {
        return {
            start: v.open.range.start,
            end: v.close.range.end
        };
    }
};
var rangeOfVoid = {
    rangeOf: Data_Void.absurd
};
var rangeOfRecoveredError = {
    rangeOf: function (v) {
        var v1 = Data_Array_NonEmpty.fromArray(v.tokens);
        if (v1 instanceof Data_Maybe.Just) {
            return {
                start: (Data_Array_NonEmpty.head(v1.value0)).range.start,
                end: (Data_Array_NonEmpty.last(v1.value0)).range.end
            };
        };
        if (v1 instanceof Data_Maybe.Nothing) {
            return {
                start: v.position,
                end: v.position
            };
        };
        throw new Error("Failed pattern match at PureScript.CST.Range (line 49, column 5 - line 57, column 10): " + [ v1.constructor.name ]);
    }
};
var rangeOfQualifiedName = {
    rangeOf: function (v) {
        return v.token.range;
    }
};
var rangeOfName = {
    rangeOf: function (v) {
        return v.token.range;
    }
};
var rangeOfModule = {
    rangeOf: function (v) {
        return {
            start: v.header.keyword.range.start,
            end: v.body.end
        };
    }
};
var rangeOf = function (dict) {
    return dict.rangeOf;
};
var rangeOfClassFundep = {
    rangeOf: function (v) {
        if (v instanceof PureScript_CST_Types.FundepDetermined) {
            return {
                start: v.value0.range.start,
                end: (rangeOf(rangeOfName)(Data_Array_NonEmpty.last(v.value1))).end
            };
        };
        if (v instanceof PureScript_CST_Types.FundepDetermines) {
            return {
                start: (rangeOf(rangeOfName)(Data_Array_NonEmpty.head(v.value0))).start,
                end: (rangeOf(rangeOfName)(Data_Array_NonEmpty.last(v.value2))).end
            };
        };
        throw new Error("Failed pattern match at PureScript.CST.Range (line 553, column 13 - line 561, column 8): " + [ v.constructor.name ]);
    }
};
var rangeOfDataMembers = {
    rangeOf: function (v) {
        if (v instanceof PureScript_CST_Types.DataAll) {
            return v.value0.range;
        };
        if (v instanceof PureScript_CST_Types.DataEnumerated) {
            return rangeOf(rangeOfWrapped)(v.value0);
        };
        throw new Error("Failed pattern match at PureScript.CST.Range (line 322, column 13 - line 326, column 16): " + [ v.constructor.name ]);
    }
};
var rangeOfExport = function (dictRangeOf) {
    return {
        rangeOf: function (v) {
            if (v instanceof PureScript_CST_Types.ExportValue) {
                return rangeOf(rangeOfName)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.ExportOp) {
                return rangeOf(rangeOfName)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.ExportType) {
                if (v.value1 instanceof Data_Maybe.Nothing) {
                    return rangeOf(rangeOfName)(v.value0);
                };
                if (v.value1 instanceof Data_Maybe.Just) {
                    return {
                        start: (rangeOf(rangeOfName)(v.value0)).start,
                        end: (rangeOf(rangeOfDataMembers)(v.value1.value0)).end
                    };
                };
                throw new Error("Failed pattern match at PureScript.CST.Range (line 282, column 7 - line 288, column 12): " + [ v.value1.constructor.name ]);
            };
            if (v instanceof PureScript_CST_Types.ExportTypeOp) {
                return {
                    start: v.value0.range.start,
                    end: (rangeOf(rangeOfName)(v.value1)).end
                };
            };
            if (v instanceof PureScript_CST_Types.ExportClass) {
                return {
                    start: v.value0.range.start,
                    end: (rangeOf(rangeOfName)(v.value1)).end
                };
            };
            if (v instanceof PureScript_CST_Types.ExportModule) {
                return {
                    start: v.value0.range.start,
                    end: (rangeOf(rangeOfName)(v.value1)).end
                };
            };
            if (v instanceof PureScript_CST_Types.ExportError) {
                return rangeOf(dictRangeOf)(v.value0);
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 276, column 13 - line 302, column 16): " + [ v.constructor.name ]);
        }
    };
};
var rangeOfFixityOp = {
    rangeOf: function (v) {
        if (v instanceof PureScript_CST_Types.FixityValue) {
            return {
                start: (rangeOf(rangeOfQualifiedName)(v.value0)).start,
                end: (rangeOf(rangeOfName)(v.value2)).end
            };
        };
        if (v instanceof PureScript_CST_Types.FixityType) {
            return {
                start: v.value0.range.start,
                end: (rangeOf(rangeOfName)(v.value3)).end
            };
        };
        throw new Error("Failed pattern match at PureScript.CST.Range (line 632, column 13 - line 640, column 8): " + [ v.constructor.name ]);
    }
};
var rangeOfImport = function (dictRangeOf) {
    return {
        rangeOf: function (v) {
            if (v instanceof PureScript_CST_Types.ImportValue) {
                return rangeOf(rangeOfName)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.ImportOp) {
                return rangeOf(rangeOfName)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.ImportType) {
                if (v.value1 instanceof Data_Maybe.Nothing) {
                    return rangeOf(rangeOfName)(v.value0);
                };
                if (v.value1 instanceof Data_Maybe.Just) {
                    return {
                        start: (rangeOf(rangeOfName)(v.value0)).start,
                        end: (rangeOf(rangeOfDataMembers)(v.value1.value0)).end
                    };
                };
                throw new Error("Failed pattern match at PureScript.CST.Range (line 365, column 7 - line 371, column 12): " + [ v.value1.constructor.name ]);
            };
            if (v instanceof PureScript_CST_Types.ImportTypeOp) {
                return {
                    start: v.value0.range.start,
                    end: (rangeOf(rangeOfName)(v.value1)).end
                };
            };
            if (v instanceof PureScript_CST_Types.ImportClass) {
                return {
                    start: v.value0.range.start,
                    end: (rangeOf(rangeOfName)(v.value1)).end
                };
            };
            if (v instanceof PureScript_CST_Types.ImportError) {
                return rangeOf(dictRangeOf)(v.value0);
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 359, column 13 - line 381, column 16): " + [ v.constructor.name ]);
        }
    };
};
var rangeOfImportDecl = {
    rangeOf: function (v) {
        var v1 = (function () {
            if (v.qualified instanceof Data_Maybe.Nothing) {
                if (v.names instanceof Data_Maybe.Nothing) {
                    return rangeOf(rangeOfName)(v.module);
                };
                if (v.names instanceof Data_Maybe.Just) {
                    return rangeOf(rangeOfWrapped)(v.names.value0.value1);
                };
                throw new Error("Failed pattern match at PureScript.CST.Range (line 340, column 11 - line 344, column 30): " + [ v.names.constructor.name ]);
            };
            if (v.qualified instanceof Data_Maybe.Just) {
                return rangeOf(rangeOfName)(v.qualified.value0.value1);
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 338, column 17 - line 346, column 20): " + [ v.qualified.constructor.name ]);
        })();
        return {
            start: v.keyword.range.start,
            end: v1.end
        };
    }
};
var rangeOfLabeled = function (dictRangeOf) {
    return function (dictRangeOf1) {
        return {
            rangeOf: function (v) {
                return {
                    start: (rangeOf(dictRangeOf)(v.label)).start,
                    end: (rangeOf(dictRangeOf1)(v.value)).end
                };
            }
        };
    };
};
var rangeOfLabeled1 = /* #__PURE__ */ rangeOfLabeled(rangeOfName);
var rangeOfOneOrDelimited = function (dictRangeOf) {
    return {
        rangeOf: function (v) {
            if (v instanceof PureScript_CST_Types.One) {
                return rangeOf(dictRangeOf)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.Many) {
                return rangeOf(rangeOfWrapped)(v.value0);
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 143, column 13 - line 145, column 26): " + [ v.constructor.name ]);
        }
    };
};
var rangeOfPrefixed = function (dictRangeOf) {
    return {
        rangeOf: function (v) {
            if (v.prefix instanceof Data_Maybe.Just) {
                return {
                    start: v.prefix.value0.range.start,
                    end: (rangeOf(dictRangeOf)(v.value)).end
                };
            };
            if (v.prefix instanceof Data_Maybe.Nothing) {
                return rangeOf(dictRangeOf)(v.value);
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 126, column 5 - line 132, column 22): " + [ v.prefix.constructor.name ]);
        }
    };
};
var rangeOfSeparated = function (dictRangeOf) {
    return {
        rangeOf: function (v) {
            var v1 = Data_Array.last(v.tail);
            if (v1 instanceof Data_Maybe.Just) {
                return {
                    start: (rangeOf(dictRangeOf)(v.head)).start,
                    end: (rangeOf(dictRangeOf)(v1.value0.value1)).end
                };
            };
            if (v1 instanceof Data_Maybe.Nothing) {
                return rangeOf(dictRangeOf)(v.head);
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 101, column 5 - line 107, column 21): " + [ v1.constructor.name ]);
        }
    };
};
var rangeOfSeparated1 = /* #__PURE__ */ rangeOfSeparated(rangeOfName);
var rangeOfSeparated2 = /* #__PURE__ */ rangeOfSeparated(rangeOfClassFundep);
var rangeOfType = function (dictRangeOf) {
    return {
        rangeOf: function (v) {
            if (v instanceof PureScript_CST_Types.TypeVar) {
                return rangeOf(rangeOfName)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.TypeConstructor) {
                return rangeOf(rangeOfQualifiedName)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.TypeWildcard) {
                return v.value0.range;
            };
            if (v instanceof PureScript_CST_Types.TypeHole) {
                return rangeOf(rangeOfName)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.TypeString) {
                return v.value0.range;
            };
            if (v instanceof PureScript_CST_Types.TypeInt) {
                if (v.value0 instanceof Data_Maybe.Nothing) {
                    return v.value1.range;
                };
                if (v.value0 instanceof Data_Maybe.Just) {
                    return {
                        start: v.value0.value0.range.start,
                        end: v.value1.range.end
                    };
                };
                throw new Error("Failed pattern match at PureScript.CST.Range (line 165, column 7 - line 171, column 12): " + [ v.value0.constructor.name ]);
            };
            if (v instanceof PureScript_CST_Types.TypeRow) {
                return rangeOf(rangeOfWrapped)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.TypeRecord) {
                return rangeOf(rangeOfWrapped)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.TypeForall) {
                return {
                    start: v.value0.range.start,
                    end: (rangeOf(rangeOfType(dictRangeOf))(v.value3)).end
                };
            };
            if (v instanceof PureScript_CST_Types.TypeKinded) {
                return {
                    start: (rangeOf(rangeOfType(dictRangeOf))(v.value0)).start,
                    end: (rangeOf(rangeOfType(dictRangeOf))(v.value2)).end
                };
            };
            if (v instanceof PureScript_CST_Types.TypeApp) {
                return {
                    start: (rangeOf(rangeOfType(dictRangeOf))(v.value0)).start,
                    end: (rangeOf(rangeOfType(dictRangeOf))(Data_Array_NonEmpty.last(v.value1))).end
                };
            };
            if (v instanceof PureScript_CST_Types.TypeOp) {
                return {
                    start: (rangeOf(rangeOfType(dictRangeOf))(v.value0)).start,
                    end: (rangeOf(rangeOfType(dictRangeOf))(Data_Tuple.snd(Data_Array_NonEmpty.last(v.value1)))).end
                };
            };
            if (v instanceof PureScript_CST_Types.TypeOpName) {
                return rangeOf(rangeOfQualifiedName)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.TypeArrow) {
                return {
                    start: (rangeOf(rangeOfType(dictRangeOf))(v.value0)).start,
                    end: (rangeOf(rangeOfType(dictRangeOf))(v.value2)).end
                };
            };
            if (v instanceof PureScript_CST_Types.TypeArrowName) {
                return v.value0.range;
            };
            if (v instanceof PureScript_CST_Types.TypeConstrained) {
                return {
                    start: (rangeOf(rangeOfType(dictRangeOf))(v.value0)).start,
                    end: (rangeOf(rangeOfType(dictRangeOf))(v.value2)).end
                };
            };
            if (v instanceof PureScript_CST_Types.TypeParens) {
                return rangeOf(rangeOfWrapped)(v.value0);
            };
            if (v instanceof PureScript_CST_Types["TypeError"]) {
                return rangeOf(dictRangeOf)(v.value0);
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 153, column 13 - line 207, column 16): " + [ v.constructor.name ]);
        }
    };
};
var rangeOfAppSpine = function (dictRangeOf) {
    var rangeOfType1 = rangeOfType(dictRangeOf);
    return function (dictRangeOf1) {
        return {
            rangeOf: function (v) {
                if (v instanceof PureScript_CST_Types.AppType) {
                    return {
                        start: v.value0.range.start,
                        end: (rangeOf(rangeOfType1)(v.value1)).end
                    };
                };
                if (v instanceof PureScript_CST_Types.AppTerm) {
                    return rangeOf(dictRangeOf1)(v.value0);
                };
                throw new Error("Failed pattern match at PureScript.CST.Range (line 847, column 13 - line 853, column 16): " + [ v.constructor.name ]);
            }
        };
    };
};
var rangeOfBinder = function (dictRangeOf) {
    var rangeOfType1 = rangeOfType(dictRangeOf);
    return {
        rangeOf: function (v) {
            if (v instanceof PureScript_CST_Types.BinderWildcard) {
                return v.value0.range;
            };
            if (v instanceof PureScript_CST_Types.BinderVar) {
                return rangeOf(rangeOfName)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.BinderNamed) {
                return {
                    start: (rangeOf(rangeOfName)(v.value0)).start,
                    end: (rangeOf(rangeOfBinder(dictRangeOf))(v.value2)).end
                };
            };
            if (v instanceof PureScript_CST_Types.BinderConstructor) {
                var v1 = Data_Array.last(v.value1);
                if (v1 instanceof Data_Maybe.Nothing) {
                    return rangeOf(rangeOfQualifiedName)(v.value0);
                };
                if (v1 instanceof Data_Maybe.Just) {
                    return {
                        start: (rangeOf(rangeOfQualifiedName)(v.value0)).start,
                        end: (rangeOf(rangeOfBinder(dictRangeOf))(v1.value0)).end
                    };
                };
                throw new Error("Failed pattern match at PureScript.CST.Range (line 932, column 7 - line 938, column 12): " + [ v1.constructor.name ]);
            };
            if (v instanceof PureScript_CST_Types.BinderBoolean) {
                return v.value0.range;
            };
            if (v instanceof PureScript_CST_Types.BinderChar) {
                return v.value0.range;
            };
            if (v instanceof PureScript_CST_Types.BinderString) {
                return v.value0.range;
            };
            if (v instanceof PureScript_CST_Types.BinderInt) {
                if (v.value0 instanceof Data_Maybe.Nothing) {
                    return v.value1.range;
                };
                if (v.value0 instanceof Data_Maybe.Just) {
                    return {
                        start: v.value0.value0.range.start,
                        end: v.value1.range.end
                    };
                };
                throw new Error("Failed pattern match at PureScript.CST.Range (line 946, column 7 - line 952, column 12): " + [ v.value0.constructor.name ]);
            };
            if (v instanceof PureScript_CST_Types.BinderNumber) {
                if (v.value0 instanceof Data_Maybe.Nothing) {
                    return v.value1.range;
                };
                if (v.value0 instanceof Data_Maybe.Just) {
                    return {
                        start: v.value0.value0.range.start,
                        end: v.value1.range.end
                    };
                };
                throw new Error("Failed pattern match at PureScript.CST.Range (line 954, column 7 - line 960, column 12): " + [ v.value0.constructor.name ]);
            };
            if (v instanceof PureScript_CST_Types.BinderArray) {
                return rangeOf(rangeOfWrapped)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.BinderRecord) {
                return rangeOf(rangeOfWrapped)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.BinderParens) {
                return rangeOf(rangeOfWrapped)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.BinderTyped) {
                return {
                    start: (rangeOf(rangeOfBinder(dictRangeOf))(v.value0)).start,
                    end: (rangeOf(rangeOfType1)(v.value2)).end
                };
            };
            if (v instanceof PureScript_CST_Types.BinderOp) {
                return {
                    start: (rangeOf(rangeOfBinder(dictRangeOf))(v.value0)).start,
                    end: (rangeOf(rangeOfBinder(dictRangeOf))(Data_Tuple.snd(Data_Array_NonEmpty.last(v.value1)))).end
                };
            };
            if (v instanceof PureScript_CST_Types.BinderError) {
                return rangeOf(dictRangeOf)(v.value0);
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 922, column 13 - line 976, column 16): " + [ v.constructor.name ]);
        }
    };
};
var rangeOfDataCtor = function (dictRangeOf) {
    var rangeOfType1 = rangeOfType(dictRangeOf);
    return {
        rangeOf: function (v) {
            var v1 = (function () {
                var v2 = Data_Array.last(v.fields);
                if (v2 instanceof Data_Maybe.Nothing) {
                    return rangeOf(rangeOfName)(v.name);
                };
                if (v2 instanceof Data_Maybe.Just) {
                    return rangeOf(rangeOfType1)(v2.value0);
                };
                throw new Error("Failed pattern match at PureScript.CST.Range (line 401, column 17 - line 405, column 21): " + [ v2.constructor.name ]);
            })();
            return {
                start: (rangeOf(rangeOfName)(v.name)).start,
                end: v1.end
            };
        }
    };
};
var rangeOfForeign = function (dictRangeOf) {
    var rangeOfLabeled2 = rangeOfLabeled1(rangeOfType(dictRangeOf));
    return {
        rangeOf: function (v) {
            if (v instanceof PureScript_CST_Types.ForeignValue) {
                return rangeOf(rangeOfLabeled2)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.ForeignData) {
                return {
                    start: v.value0.range.start,
                    end: (rangeOf(rangeOfLabeled2)(v.value1)).end
                };
            };
            if (v instanceof PureScript_CST_Types.ForeignKind) {
                return {
                    start: v.value0.range.start,
                    end: (rangeOf(rangeOfName)(v.value1)).end
                };
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 650, column 13 - line 660, column 8): " + [ v.constructor.name ]);
        }
    };
};
var rangeOfTypeVarBinding = function (dictRangeOf) {
    return {
        rangeOf: function (v) {
            if (v instanceof PureScript_CST_Types.TypeVarKinded) {
                return rangeOf(rangeOfWrapped)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.TypeVarName) {
                return rangeOf(dictRangeOf)(v.value0);
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 262, column 13 - line 266, column 16): " + [ v.constructor.name ]);
        }
    };
};
var rangeOfTypeVarBinding1 = /* #__PURE__ */ rangeOfTypeVarBinding(rangeOfName);
var rangeOfWhere = function (dictRangeOf) {
    return {
        rangeOf: function (v) {
            if (v.bindings instanceof Data_Maybe.Nothing) {
                return rangeOf(rangeOfExpr(dictRangeOf))(v.expr);
            };
            if (v.bindings instanceof Data_Maybe.Just) {
                return {
                    start: (rangeOf(rangeOfExpr(dictRangeOf))(v.expr)).start,
                    end: (rangeOf(rangeOfLetBinding(dictRangeOf))(Data_Array_NonEmpty.last(v.bindings.value0.value1))).end
                };
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 1019, column 40 - line 1025, column 8): " + [ v.bindings.constructor.name ]);
        }
    };
};
var rangeOfLetBinding = function (dictRangeOf) {
    var rangeOfLabeled2 = rangeOfLabeled1(rangeOfType(dictRangeOf));
    var rangeOfBinder1 = rangeOfBinder(dictRangeOf);
    return {
        rangeOf: function (v) {
            if (v instanceof PureScript_CST_Types.LetBindingSignature) {
                return rangeOf(rangeOfLabeled2)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.LetBindingName) {
                return {
                    start: (rangeOf(rangeOfName)(v.value0.name)).start,
                    end: (rangeOf(rangeOfGuarded(dictRangeOf))(v.value0.guarded)).end
                };
            };
            if (v instanceof PureScript_CST_Types.LetBindingPattern) {
                return {
                    start: (rangeOf(rangeOfBinder1)(v.value0)).start,
                    end: (rangeOf(rangeOfWhere(dictRangeOf))(v.value2)).end
                };
            };
            if (v instanceof PureScript_CST_Types.LetBindingError) {
                return rangeOf(dictRangeOf)(v.value0);
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 896, column 13 - line 908, column 16): " + [ v.constructor.name ]);
        }
    };
};
var rangeOfGuardedExpr = function (dictRangeOf) {
    return {
        rangeOf: function (v) {
            return {
                start: v.bar.range.start,
                end: (rangeOf(rangeOfWhere(dictRangeOf))(v.where)).end
            };
        }
    };
};
var rangeOfGuarded = function (dictRangeOf) {
    return {
        rangeOf: function (v) {
            if (v instanceof PureScript_CST_Types.Unconditional) {
                return {
                    start: v.value0.range.start,
                    end: (rangeOf(rangeOfWhere(dictRangeOf))(v.value1)).end
                };
            };
            if (v instanceof PureScript_CST_Types.Guarded) {
                return {
                    start: (rangeOf(rangeOfGuardedExpr(dictRangeOf))(Data_Array_NonEmpty.head(v.value0))).start,
                    end: (rangeOf(rangeOfGuardedExpr(dictRangeOf))(Data_Array_NonEmpty.last(v.value0))).end
                };
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 596, column 13 - line 604, column 8): " + [ v.constructor.name ]);
        }
    };
};
var rangeOfExpr = function (dictRangeOf) {
    var rangeOfType1 = rangeOfType(dictRangeOf);
    var rangeOfAppSpine1 = rangeOfAppSpine(dictRangeOf);
    return {
        rangeOf: function (v) {
            if (v instanceof PureScript_CST_Types.ExprHole) {
                return rangeOf(rangeOfName)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.ExprSection) {
                return v.value0.range;
            };
            if (v instanceof PureScript_CST_Types.ExprIdent) {
                return rangeOf(rangeOfQualifiedName)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.ExprConstructor) {
                return rangeOf(rangeOfQualifiedName)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.ExprBoolean) {
                return v.value0.range;
            };
            if (v instanceof PureScript_CST_Types.ExprChar) {
                return v.value0.range;
            };
            if (v instanceof PureScript_CST_Types.ExprString) {
                return v.value0.range;
            };
            if (v instanceof PureScript_CST_Types.ExprInt) {
                return v.value0.range;
            };
            if (v instanceof PureScript_CST_Types.ExprNumber) {
                return v.value0.range;
            };
            if (v instanceof PureScript_CST_Types.ExprArray) {
                return rangeOf(rangeOfWrapped)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.ExprRecord) {
                return rangeOf(rangeOfWrapped)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.ExprParens) {
                return rangeOf(rangeOfWrapped)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.ExprTyped) {
                return {
                    start: (rangeOf(rangeOfExpr(dictRangeOf))(v.value0)).start,
                    end: (rangeOf(rangeOfType1)(v.value2)).end
                };
            };
            if (v instanceof PureScript_CST_Types.ExprInfix) {
                return {
                    start: (rangeOf(rangeOfExpr(dictRangeOf))(v.value0)).start,
                    end: (rangeOf(rangeOfExpr(dictRangeOf))(Data_Tuple.snd(Data_Array_NonEmpty.last(v.value1)))).end
                };
            };
            if (v instanceof PureScript_CST_Types.ExprOp) {
                return {
                    start: (rangeOf(rangeOfExpr(dictRangeOf))(v.value0)).start,
                    end: (rangeOf(rangeOfExpr(dictRangeOf))(Data_Tuple.snd(Data_Array_NonEmpty.last(v.value1)))).end
                };
            };
            if (v instanceof PureScript_CST_Types.ExprOpName) {
                return rangeOf(rangeOfQualifiedName)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.ExprNegate) {
                return {
                    start: v.value0.range.start,
                    end: (rangeOf(rangeOfExpr(dictRangeOf))(v.value1)).end
                };
            };
            if (v instanceof PureScript_CST_Types.ExprRecordAccessor) {
                return {
                    start: (rangeOf(rangeOfExpr(dictRangeOf))(v.value0.expr)).start,
                    end: (rangeOf(rangeOfSeparated1)(v.value0.path)).end
                };
            };
            if (v instanceof PureScript_CST_Types.ExprRecordUpdate) {
                return {
                    start: (rangeOf(rangeOfExpr(dictRangeOf))(v.value0)).start,
                    end: (rangeOf(rangeOfWrapped)(v.value1)).end
                };
            };
            if (v instanceof PureScript_CST_Types.ExprApp) {
                return {
                    start: (rangeOf(rangeOfExpr(dictRangeOf))(v.value0)).start,
                    end: (rangeOf(rangeOfAppSpine1(rangeOfExpr(dictRangeOf)))(Data_Array_NonEmpty.last(v.value1))).end
                };
            };
            if (v instanceof PureScript_CST_Types.ExprLambda) {
                return {
                    start: v.value0.symbol.range.start,
                    end: (rangeOf(rangeOfExpr(dictRangeOf))(v.value0.body)).end
                };
            };
            if (v instanceof PureScript_CST_Types.ExprIf) {
                return {
                    start: v.value0.keyword.range.start,
                    end: (rangeOf(rangeOfExpr(dictRangeOf))(v["value0"]["false"])).end
                };
            };
            if (v instanceof PureScript_CST_Types.ExprCase) {
                return {
                    start: v.value0.keyword.range.start,
                    end: (rangeOf(rangeOfGuarded(dictRangeOf))(Data_Tuple.snd(Data_Array_NonEmpty.last(v.value0.branches)))).end
                };
            };
            if (v instanceof PureScript_CST_Types.ExprLet) {
                return {
                    start: v.value0.keyword.range.start,
                    end: (rangeOf(rangeOfExpr(dictRangeOf))(v.value0.body)).end
                };
            };
            if (v instanceof PureScript_CST_Types.ExprDo) {
                return {
                    start: v.value0.keyword.range.start,
                    end: (rangeOf(rangeOfDoStatement(dictRangeOf))(Data_Array_NonEmpty.last(v.value0.statements))).end
                };
            };
            if (v instanceof PureScript_CST_Types.ExprAdo) {
                return {
                    start: v.value0.keyword.range.start,
                    end: (rangeOf(rangeOfExpr(dictRangeOf))(v.value0.result)).end
                };
            };
            if (v instanceof PureScript_CST_Types.ExprError) {
                return rangeOf(dictRangeOf)(v.value0);
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 690, column 13 - line 770, column 16): " + [ v.constructor.name ]);
        }
    };
};
var rangeOfDoStatement = function (dictRangeOf) {
    var rangeOfBinder1 = rangeOfBinder(dictRangeOf);
    return {
        rangeOf: function (v) {
            if (v instanceof PureScript_CST_Types.DoLet) {
                return {
                    start: v.value0.range.start,
                    end: (rangeOf(rangeOfLetBinding(dictRangeOf))(Data_Array_NonEmpty.last(v.value1))).end
                };
            };
            if (v instanceof PureScript_CST_Types.DoDiscard) {
                return rangeOf(rangeOfExpr(dictRangeOf))(v.value0);
            };
            if (v instanceof PureScript_CST_Types.DoBind) {
                return {
                    start: (rangeOf(rangeOfBinder1)(v.value0)).start,
                    end: (rangeOf(rangeOfExpr(dictRangeOf))(v.value2)).end
                };
            };
            if (v instanceof PureScript_CST_Types.DoError) {
                return rangeOf(dictRangeOf)(v.value0);
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 870, column 13 - line 882, column 16): " + [ v.constructor.name ]);
        }
    };
};
var rangeOfInstanceBinding = function (dictRangeOf) {
    var rangeOfLabeled2 = rangeOfLabeled1(rangeOfType(dictRangeOf));
    var rangeOfGuarded1 = rangeOfGuarded(dictRangeOf);
    return {
        rangeOf: function (v) {
            if (v instanceof PureScript_CST_Types.InstanceBindingSignature) {
                return rangeOf(rangeOfLabeled2)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.InstanceBindingName) {
                return {
                    start: (rangeOf(rangeOfName)(v.value0.name)).start,
                    end: (rangeOf(rangeOfGuarded1)(v.value0.guarded)).end
                };
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 672, column 13 - line 678, column 8): " + [ v.constructor.name ]);
        }
    };
};
var rangeOfInstance = function (dictRangeOf) {
    var rangeOfType1 = rangeOfType(dictRangeOf);
    var rangeOfInstanceBinding1 = rangeOfInstanceBinding(dictRangeOf);
    return {
        rangeOf: function (v) {
            var v1 = (function () {
                if (v.body instanceof Data_Maybe.Nothing) {
                    var v2 = Data_Array.last(v.head.types);
                    if (v2 instanceof Data_Maybe.Nothing) {
                        return rangeOf(rangeOfQualifiedName)(v.head.className);
                    };
                    if (v2 instanceof Data_Maybe.Just) {
                        return rangeOf(rangeOfType1)(v2.value0);
                    };
                    throw new Error("Failed pattern match at PureScript.CST.Range (line 575, column 11 - line 579, column 25): " + [ v2.constructor.name ]);
                };
                if (v.body instanceof Data_Maybe.Just) {
                    return rangeOf(rangeOfInstanceBinding1)(Data_Array_NonEmpty.last(v.body.value0.value1));
                };
                throw new Error("Failed pattern match at PureScript.CST.Range (line 573, column 17 - line 581, column 42): " + [ v.body.constructor.name ]);
            })();
            return {
                start: v.head.keyword.range.start,
                end: v1.end
            };
        }
    };
};
var rangeOfDecl = function (dictRangeOf) {
    var rangeOf1 = rangeOf(rangeOfDataCtor(dictRangeOf));
    var rangeOfType1 = rangeOfType(dictRangeOf);
    var rangeOfLabeled2 = rangeOfLabeled1(rangeOfType(dictRangeOf));
    var rangeOfSeparated3 = rangeOfSeparated(rangeOfInstance(dictRangeOf));
    var rangeOfGuarded1 = rangeOfGuarded(dictRangeOf);
    var rangeOfForeign1 = rangeOfForeign(dictRangeOf);
    return {
        rangeOf: function (v) {
            if (v instanceof PureScript_CST_Types.DeclData) {
                var v1 = (function () {
                    if (v.value1 instanceof Data_Maybe.Nothing) {
                        var v2 = Data_Array.last(v.value0.vars);
                        if (v2 instanceof Data_Maybe.Nothing) {
                            return rangeOf(rangeOfName)(v.value0.name);
                        };
                        if (v2 instanceof Data_Maybe.Just) {
                            return rangeOf(rangeOfTypeVarBinding1)(v2.value0);
                        };
                        throw new Error("Failed pattern match at PureScript.CST.Range (line 420, column 13 - line 424, column 28): " + [ v2.constructor.name ]);
                    };
                    if (v.value1 instanceof Data_Maybe.Just) {
                        return rangeOf1(Data_Maybe.maybe(v.value1.value0.value1.head)(Data_Tuple.snd)(Data_Array.last(v.value1.value0.value1.tail)));
                    };
                    throw new Error("Failed pattern match at PureScript.CST.Range (line 418, column 19 - line 426, column 55): " + [ v.value1.constructor.name ]);
                })();
                return {
                    start: v.value0.keyword.range.start,
                    end: v1.end
                };
            };
            if (v instanceof PureScript_CST_Types.DeclType) {
                return {
                    start: v.value0.keyword.range.start,
                    end: (rangeOf(rangeOfType1)(v.value2)).end
                };
            };
            if (v instanceof PureScript_CST_Types.DeclNewtype) {
                return {
                    start: v.value0.keyword.range.start,
                    end: (rangeOf(rangeOfType1)(v.value3)).end
                };
            };
            if (v instanceof PureScript_CST_Types.DeclClass) {
                var v1 = (function () {
                    if (v.value1 instanceof Data_Maybe.Nothing) {
                        if (v.value0.fundeps instanceof Data_Maybe.Nothing) {
                            var v2 = Data_Array.last(v.value0.vars);
                            if (v2 instanceof Data_Maybe.Nothing) {
                                return rangeOf(rangeOfName)(v.value0.name);
                            };
                            if (v2 instanceof Data_Maybe.Just) {
                                return rangeOf(rangeOfTypeVarBinding1)(v2.value0);
                            };
                            throw new Error("Failed pattern match at PureScript.CST.Range (line 444, column 17 - line 448, column 32): " + [ v2.constructor.name ]);
                        };
                        if (v.value0.fundeps instanceof Data_Maybe.Just) {
                            return rangeOf(rangeOfSeparated2)(v.value0.fundeps.value0.value1);
                        };
                        throw new Error("Failed pattern match at PureScript.CST.Range (line 442, column 13 - line 450, column 32): " + [ v.value0.fundeps.constructor.name ]);
                    };
                    if (v.value1 instanceof Data_Maybe.Just) {
                        return rangeOf(rangeOfLabeled2)(Data_Array_NonEmpty.last(v.value1.value0.value1));
                    };
                    throw new Error("Failed pattern match at PureScript.CST.Range (line 440, column 19 - line 452, column 44): " + [ v.value1.constructor.name ]);
                })();
                return {
                    start: v.value0.keyword.range.start,
                    end: v1.end
                };
            };
            if (v instanceof PureScript_CST_Types.DeclInstanceChain) {
                return rangeOf(rangeOfSeparated3)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.DeclDerive) {
                var v1 = (function () {
                    var v2 = Data_Array.last(v.value2.types);
                    if (v2 instanceof Data_Maybe.Nothing) {
                        return rangeOf(rangeOfQualifiedName)(v.value2.className);
                    };
                    if (v2 instanceof Data_Maybe.Just) {
                        return rangeOf(rangeOfType1)(v2.value0);
                    };
                    throw new Error("Failed pattern match at PureScript.CST.Range (line 460, column 19 - line 464, column 23): " + [ v2.constructor.name ]);
                })();
                return {
                    start: v.value0.range.start,
                    end: v1.end
                };
            };
            if (v instanceof PureScript_CST_Types.DeclKindSignature) {
                return {
                    start: v.value0.range.start,
                    end: (rangeOf(rangeOfLabeled2)(v.value1)).end
                };
            };
            if (v instanceof PureScript_CST_Types.DeclSignature) {
                return rangeOf(rangeOfLabeled2)(v.value0);
            };
            if (v instanceof PureScript_CST_Types.DeclValue) {
                return {
                    start: (rangeOf(rangeOfName)(v.value0.name)).start,
                    end: (rangeOf(rangeOfGuarded1)(v.value0.guarded)).end
                };
            };
            if (v instanceof PureScript_CST_Types.DeclFixity) {
                return {
                    start: v.value0.keyword.value0.range.start,
                    end: (rangeOf(rangeOfFixityOp)(v.value0.operator)).end
                };
            };
            if (v instanceof PureScript_CST_Types.DeclForeign) {
                return {
                    start: v.value0.range.start,
                    end: (rangeOf(rangeOfForeign1)(v.value2)).end
                };
            };
            if (v instanceof PureScript_CST_Types.DeclRole) {
                return {
                    start: v.value0.range.start,
                    end: (Data_Tuple.fst(Data_Array_NonEmpty.last(v.value3))).range.end
                };
            };
            if (v instanceof PureScript_CST_Types.DeclError) {
                return rangeOf(dictRangeOf)(v.value0);
            };
            throw new Error("Failed pattern match at PureScript.CST.Range (line 415, column 13 - line 491, column 16): " + [ v.constructor.name ]);
        }
    };
};
export {
    rangeOf,
    tokensOf,
    tokensOfTuple,
    tokensOfMaybe,
    tokensOfArray,
    tokensOfNonEmptyArray,
    rangeOfVoid,
    tokensOfVoid,
    rangeOfRecoveredError,
    tokensOfRecoveredError,
    rangeOfModule,
    tokensOfModule,
    rangeOfName,
    tokensOfName,
    rangeOfQualifiedName,
    tokensOfQualifiedName,
    rangeOfWrapped,
    tokensOfWrapped,
    rangeOfSeparated,
    tokensOfSeparated,
    rangeOfLabeled,
    tokensOfLabeled,
    rangeOfPrefixed,
    tokensOfPrefixed,
    rangeOfOneOrDelimited,
    tokensOfOneOrDelimited,
    rangeOfType,
    tokensOfType,
    tokensOfRow,
    rangeOfTypeVarBinding,
    tokensOfTypeVarBinding,
    rangeOfExport,
    tokensOfExport,
    rangeOfDataMembers,
    tokensOfDataMembers,
    rangeOfImportDecl,
    tokensOfImportDecl,
    rangeOfImport,
    tokensOfImport,
    rangeOfDataCtor,
    tokensOfDataCtor,
    rangeOfDecl,
    tokensOfDecl,
    rangeOfClassFundep,
    tokensOfClassFundep,
    rangeOfInstance,
    tokensOfInstance,
    rangeOfGuarded,
    tokensOfGuarded,
    rangeOfGuardedExpr,
    tokensOfGuardedExpr,
    tokensOfPatternGuard,
    rangeOfFixityOp,
    tokensOfFixityOp,
    rangeOfForeign,
    tokensOfForeign,
    rangeOfInstanceBinding,
    tokensOfInstanceBinding,
    rangeOfExpr,
    tokensOfExpr,
    rangeOfAppSpine,
    tokensOfAppSpine,
    tokensOfRecordUpdate,
    rangeOfDoStatement,
    tokensOfDoStatement,
    rangeOfLetBinding,
    tokensOfLetBinding,
    rangeOfBinder,
    tokensOfBinder,
    tokensOfRecordLabeled,
    rangeOfWhere,
    tokensOfWhere
};
