import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Category from "../Control.Category/index.js";
import * as Control_Monad_Free from "../Control.Monad.Free/index.js";
import * as Control_Monad_Reader_Trans from "../Control.Monad.Reader.Trans/index.js";
import * as Data_Array_NonEmpty_Internal from "../Data.Array.NonEmpty.Internal/index.js";
import * as Data_Bitraversable from "../Data.Bitraversable/index.js";
import * as Data_Const from "../Data.Const/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Functor_Compose from "../Data.Functor.Compose/index.js";
import * as Data_Identity from "../Data.Identity/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Newtype from "../Data.Newtype/index.js";
import * as Data_Traversable from "../Data.Traversable/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Data_Unit from "../Data.Unit/index.js";
import * as PureScript_CST_Types from "../PureScript.CST.Types/index.js";
var applicativeReaderT = /* #__PURE__ */ Control_Monad_Reader_Trans.applicativeReaderT(Data_Identity.applicativeIdentity);
var applyCompose = /* #__PURE__ */ Data_Functor_Compose.applyCompose(Control_Monad_Free.freeApply);
var applicativeCompose = /* #__PURE__ */ Data_Functor_Compose.applicativeCompose(Control_Monad_Free.freeApplicative);
var pure = /* #__PURE__ */ Control_Applicative.pure(Control_Monad_Free.freeApplicative);
var pure1 = /* #__PURE__ */ Control_Applicative.pure(Control_Monad_Free.freeApplicative);
var pure2 = /* #__PURE__ */ Control_Applicative.pure(Control_Monad_Free.freeApplicative);
var pure3 = /* #__PURE__ */ Control_Applicative.pure(Control_Monad_Free.freeApplicative);
var traverseWrapped = function (dictApplicative) {
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (v) {
            return Data_Functor.map(Functor0)(function (value) {
                return {
                    open: v.open,
                    close: v.close,
                    value: value
                };
            })(k(v.value));
        };
    };
};
var traverseSeparated = function (dictApplicative) {
    var Apply0 = dictApplicative.Apply0();
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (v) {
            return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(function (v1) {
                return function (v2) {
                    return {
                        head: v1,
                        tail: v2
                    };
                };
            })(k(v.head)))(Data_Traversable.traverse(Data_Traversable.traversableArray)(dictApplicative)(Data_Traversable.traverse(Data_Traversable.traversableTuple)(dictApplicative)(k))(v.tail));
        };
    };
};
var traverseRecordUpdate = function (dictApplicative) {
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (v) {
            if (v instanceof PureScript_CST_Types.RecordUpdateLeaf) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.RecordUpdateLeaf.create(v.value0)(v.value1))(k.onExpr(v.value2));
            };
            if (v instanceof PureScript_CST_Types.RecordUpdateBranch) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.RecordUpdateBranch.create(v.value0))(traverseWrapped(dictApplicative)(traverseSeparated(dictApplicative)(traverseRecordUpdate(dictApplicative)(k)))(v.value1));
            };
            throw new Error("Failed pattern match at PureScript.CST.Traversal (line 419, column 26 - line 421, column 146): " + [ v.constructor.name ]);
        };
    };
};
var traverseRecordLabeled = function (dictApplicative) {
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (v) {
            if (v instanceof PureScript_CST_Types.RecordPun) {
                return Control_Applicative.pure(dictApplicative)(new PureScript_CST_Types.RecordPun(v.value0));
            };
            if (v instanceof PureScript_CST_Types.RecordField) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.RecordField.create(v.value0)(v.value1))(k(v.value2));
            };
            throw new Error("Failed pattern match at PureScript.CST.Traversal (line 394, column 27 - line 396, column 57): " + [ v.constructor.name ]);
        };
    };
};
var traverseRecordAccessor = function (dictApplicative) {
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (r) {
            return Data_Functor.map(Functor0)(function (v) {
                return {
                    dot: r.dot,
                    path: r.path,
                    expr: v
                };
            })(k.onExpr(r.expr));
        };
    };
};
var traversePatternGuard = function (dictApplicative) {
    var Apply0 = dictApplicative.Apply0();
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (v) {
            return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(function (binder) {
                return function (expr) {
                    return {
                        binder: binder,
                        expr: expr
                    };
                };
            })(Data_Traversable.traverse(Data_Traversable.traversableMaybe)(dictApplicative)(Data_Bitraversable.ltraverse(Data_Bitraversable.bitraversableTuple)(dictApplicative)(k.onBinder))(v.binder)))(k.onExpr(v.expr));
        };
    };
};
var traverseModuleBody = function (dictApplicative) {
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (v) {
            return Data_Functor.map(Functor0)(function (decls) {
                return {
                    trailingComments: v.trailingComments,
                    end: v.end,
                    decls: decls
                };
            })(Data_Traversable.traverse(Data_Traversable.traversableArray)(dictApplicative)(k.onDecl)(v.decls));
        };
    };
};
var traverseModule = function (dictApplicative) {
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (v) {
            return Data_Functor.map(Functor0)(function (body) {
                return {
                    header: v.header,
                    body: body
                };
            })(traverseModuleBody(dictApplicative)(k)(v.body));
        };
    };
};
var traverseModule1 = /* #__PURE__ */ traverseModule(Control_Monad_Free.freeApplicative);
var traverseLambda = function (dictApplicative) {
    var Apply0 = dictApplicative.Apply0();
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (l) {
            return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(function (v) {
                return function (v1) {
                    return {
                        arrow: l.arrow,
                        symbol: l.symbol,
                        binders: v,
                        body: v1
                    };
                };
            })(Data_Traversable.traverse(Data_Array_NonEmpty_Internal.traversableNonEmptyArray)(dictApplicative)(k.onBinder)(l.binders)))(k.onExpr(l.body));
        };
    };
};
var traverseLabeled = function (dictApplicative) {
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (v) {
            return Data_Functor.map(Functor0)(function (value) {
                return {
                    label: v.label,
                    separator: v.separator,
                    value: value
                };
            })(k(v.value));
        };
    };
};
var traverseRow = function (dictApplicative) {
    var Apply0 = dictApplicative.Apply0();
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (v) {
            return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(function (labels) {
                return function (tail) {
                    return {
                        labels: labels,
                        tail: tail
                    };
                };
            })(Data_Traversable.traverse(Data_Traversable.traversableMaybe)(dictApplicative)(traverseSeparated(dictApplicative)(traverseLabeled(dictApplicative)(k.onType)))(v.labels)))(Data_Traversable.traverse(Data_Traversable.traversableMaybe)(dictApplicative)(Data_Traversable.traverse(Data_Traversable.traversableTuple)(dictApplicative)(k.onType))(v.tail));
        };
    };
};
var traverseTypeVarBinding = function (dictApplicative) {
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (v) {
            if (v instanceof PureScript_CST_Types.TypeVarKinded) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.TypeVarKinded.create)(traverseWrapped(dictApplicative)(traverseLabeled(dictApplicative)(k.onType))(v.value0));
            };
            if (v instanceof PureScript_CST_Types.TypeVarName) {
                return Control_Applicative.pure(dictApplicative)(new PureScript_CST_Types.TypeVarName(v.value0));
            };
            throw new Error("Failed pattern match at PureScript.CST.Traversal (line 320, column 28 - line 322, column 46): " + [ v.constructor.name ]);
        };
    };
};
var traverseType = function (dictApplicative) {
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    var Apply0 = dictApplicative.Apply0();
    return function (k) {
        return function (v) {
            if (v instanceof PureScript_CST_Types.TypeRow) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.TypeRow.create)(traverseWrapped(dictApplicative)(traverseRow(dictApplicative)(k))(v.value0));
            };
            if (v instanceof PureScript_CST_Types.TypeRecord) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.TypeRecord.create)(traverseWrapped(dictApplicative)(traverseRow(dictApplicative)(k))(v.value0));
            };
            if (v instanceof PureScript_CST_Types.TypeForall) {
                return Control_Apply.apply(Apply0)(Data_Functor.flap(Functor0)(Data_Functor.map(Functor0)(PureScript_CST_Types.TypeForall.create(v.value0))(Data_Traversable.traverse(Data_Array_NonEmpty_Internal.traversableNonEmptyArray)(dictApplicative)(traverseTypeVarBinding(dictApplicative)(k))(v.value1)))(v.value2))(k.onType(v.value3));
            };
            if (v instanceof PureScript_CST_Types.TypeKinded) {
                return Control_Apply.apply(Apply0)(Data_Functor.flap(Functor0)(Data_Functor.map(Functor0)(PureScript_CST_Types.TypeKinded.create)(k.onType(v.value0)))(v.value1))(k.onType(v.value2));
            };
            if (v instanceof PureScript_CST_Types.TypeApp) {
                return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(PureScript_CST_Types.TypeApp.create)(k.onType(v.value0)))(Data_Traversable.traverse(Data_Array_NonEmpty_Internal.traversableNonEmptyArray)(dictApplicative)(k.onType)(v.value1));
            };
            if (v instanceof PureScript_CST_Types.TypeOp) {
                return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(PureScript_CST_Types.TypeOp.create)(k.onType(v.value0)))(Data_Traversable.traverse(Data_Array_NonEmpty_Internal.traversableNonEmptyArray)(dictApplicative)(Data_Traversable.traverse(Data_Traversable.traversableTuple)(dictApplicative)(k.onType))(v.value1));
            };
            if (v instanceof PureScript_CST_Types.TypeArrow) {
                return Control_Apply.apply(Apply0)(Data_Functor.flap(Functor0)(Data_Functor.map(Functor0)(PureScript_CST_Types.TypeArrow.create)(k.onType(v.value0)))(v.value1))(k.onType(v.value2));
            };
            if (v instanceof PureScript_CST_Types.TypeConstrained) {
                return Control_Apply.apply(Apply0)(Data_Functor.flap(Functor0)(Data_Functor.map(Functor0)(PureScript_CST_Types.TypeConstrained.create)(k.onType(v.value0)))(v.value1))(k.onType(v.value2));
            };
            if (v instanceof PureScript_CST_Types.TypeParens) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.TypeParens.create)(traverseWrapped(dictApplicative)(k.onType)(v.value0));
            };
            return Control_Applicative.pure(dictApplicative)(v);
        };
    };
};
var traverseIfThenElse = function (dictApplicative) {
    var Apply0 = dictApplicative.Apply0();
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (r) {
            return Control_Apply.apply(Apply0)(Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(function (v) {
                return function (v1) {
                    return function (v2) {
                        return {
                            "else": r["else"],
                            keyword: r.keyword,
                            then: r.then,
                            cond: v,
                            "true": v1,
                            "false": v2
                        };
                    };
                };
            })(k.onExpr(r.cond)))(k.onExpr(r["true"])))(k.onExpr(r["false"]));
        };
    };
};
var traverseWhere = function (dictApplicative) {
    var Apply0 = dictApplicative.Apply0();
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (v) {
            return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(function (expr) {
                return function (bindings) {
                    return {
                        expr: expr,
                        bindings: bindings
                    };
                };
            })(k.onExpr(v.expr)))(Data_Traversable.traverse(Data_Traversable.traversableMaybe)(dictApplicative)(Data_Traversable.traverse(Data_Traversable.traversableTuple)(dictApplicative)(Data_Traversable.traverse(Data_Array_NonEmpty_Internal.traversableNonEmptyArray)(dictApplicative)(traverseLetBinding(dictApplicative)(k))))(v.bindings));
        };
    };
};
var traverseValueBindingFields = function (dictApplicative) {
    var Apply0 = dictApplicative.Apply0();
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (v) {
            return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(function (v1) {
                return function (v2) {
                    return {
                        name: v.name,
                        binders: v1,
                        guarded: v2
                    };
                };
            })(Data_Traversable.traverse(Data_Traversable.traversableArray)(dictApplicative)(k.onBinder)(v.binders)))(traverseGuarded(dictApplicative)(k)(v.guarded));
        };
    };
};
var traverseLetBinding = function (dictApplicative) {
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    var Apply0 = dictApplicative.Apply0();
    return function (k) {
        return function (v) {
            if (v instanceof PureScript_CST_Types.LetBindingSignature) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.LetBindingSignature.create)(traverseLabeled(dictApplicative)(k.onType)(v.value0));
            };
            if (v instanceof PureScript_CST_Types.LetBindingName) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.LetBindingName.create)(traverseValueBindingFields(dictApplicative)(k)(v.value0));
            };
            if (v instanceof PureScript_CST_Types.LetBindingPattern) {
                return Control_Apply.apply(Apply0)(Data_Functor.flap(Functor0)(Data_Functor.map(Functor0)(PureScript_CST_Types.LetBindingPattern.create)(k.onBinder(v.value0)))(v.value1))(traverseWhere(dictApplicative)(k)(v.value2));
            };
            if (v instanceof PureScript_CST_Types.LetBindingError) {
                return Control_Applicative.pure(dictApplicative)(new PureScript_CST_Types.LetBindingError(v.value0));
            };
            throw new Error("Failed pattern match at PureScript.CST.Traversal (line 498, column 24 - line 502, column 48): " + [ v.constructor.name ]);
        };
    };
};
var traverseGuardedExpr = function (dictApplicative) {
    var Apply0 = dictApplicative.Apply0();
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (v) {
            return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(function (ps) {
                return function (wh) {
                    return {
                        bar: v.bar,
                        separator: v.separator,
                        patterns: ps,
                        where: wh
                    };
                };
            })(traverseSeparated(dictApplicative)(traversePatternGuard(dictApplicative)(k))(v.patterns)))(traverseWhere(dictApplicative)(k)(v.where));
        };
    };
};
var traverseGuarded = function (dictApplicative) {
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (v) {
            if (v instanceof PureScript_CST_Types.Unconditional) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.Unconditional.create(v.value0))(traverseWhere(dictApplicative)(k)(v.value1));
            };
            if (v instanceof PureScript_CST_Types.Guarded) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.Guarded.create)(Data_Traversable.traverse(Data_Array_NonEmpty_Internal.traversableNonEmptyArray)(dictApplicative)(traverseGuardedExpr(dictApplicative)(k))(v.value0));
            };
            throw new Error("Failed pattern match at PureScript.CST.Traversal (line 459, column 21 - line 461, column 72): " + [ v.constructor.name ]);
        };
    };
};
var traverseInstanceBinding = function (dictApplicative) {
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (v) {
            if (v instanceof PureScript_CST_Types.InstanceBindingSignature) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.InstanceBindingSignature.create)(traverseLabeled(dictApplicative)(k.onType)(v.value0));
            };
            if (v instanceof PureScript_CST_Types.InstanceBindingName) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.InstanceBindingName.create)(traverseValueBindingFields(dictApplicative)(k)(v.value0));
            };
            throw new Error("Failed pattern match at PureScript.CST.Traversal (line 247, column 29 - line 249, column 92): " + [ v.constructor.name ]);
        };
    };
};
var traverseLetIn = function (dictApplicative) {
    var Apply0 = dictApplicative.Apply0();
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (l) {
            return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(function (v) {
                return function (v1) {
                    return {
                        "in": l["in"],
                        keyword: l.keyword,
                        bindings: v,
                        body: v1
                    };
                };
            })(Data_Traversable.traverse(Data_Array_NonEmpty_Internal.traversableNonEmptyArray)(dictApplicative)(traverseLetBinding(dictApplicative)(k))(l.bindings)))(k.onExpr(l.body));
        };
    };
};
var traverseForeign = function (dictApplicative) {
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (v) {
            if (v instanceof PureScript_CST_Types.ForeignValue) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.ForeignValue.create)(traverseLabeled(dictApplicative)(k.onType)(v.value0));
            };
            if (v instanceof PureScript_CST_Types.ForeignData) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.ForeignData.create(v.value0))(traverseLabeled(dictApplicative)(k.onType)(v.value1));
            };
            if (v instanceof PureScript_CST_Types.ForeignKind) {
                return Control_Applicative.pure(dictApplicative)(v);
            };
            throw new Error("Failed pattern match at PureScript.CST.Traversal (line 217, column 21 - line 220, column 40): " + [ v.constructor.name ]);
        };
    };
};
var traverseExprAppSpine = function (dictApplicative) {
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (v) {
            if (v instanceof PureScript_CST_Types.AppType) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.AppType.create(v.value0))(k.onType(v.value1));
            };
            if (v instanceof PureScript_CST_Types.AppTerm) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.AppTerm.create)(k.onExpr(v.value0));
            };
            throw new Error("Failed pattern match at PureScript.CST.Traversal (line 353, column 26 - line 355, column 44): " + [ v.constructor.name ]);
        };
    };
};
var traverseDoStatement = function (dictApplicative) {
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    var Apply0 = dictApplicative.Apply0();
    return function (k) {
        return function (v) {
            if (v instanceof PureScript_CST_Types.DoLet) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.DoLet.create(v.value0))(Data_Traversable.traverse(Data_Array_NonEmpty_Internal.traversableNonEmptyArray)(dictApplicative)(traverseLetBinding(dictApplicative)(k))(v.value1));
            };
            if (v instanceof PureScript_CST_Types.DoDiscard) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.DoDiscard.create)(k.onExpr(v.value0));
            };
            if (v instanceof PureScript_CST_Types.DoBind) {
                return Control_Apply.apply(Apply0)(Data_Functor.flap(Functor0)(Data_Functor.map(Functor0)(PureScript_CST_Types.DoBind.create)(k.onBinder(v.value0)))(v.value1))(k.onExpr(v.value2));
            };
            if (v instanceof PureScript_CST_Types.DoError) {
                return Control_Applicative.pure(dictApplicative)(new PureScript_CST_Types.DoError(v.value0));
            };
            throw new Error("Failed pattern match at PureScript.CST.Traversal (line 529, column 25 - line 533, column 32): " + [ v.constructor.name ]);
        };
    };
};
var traverseDoBlock = function (dictApplicative) {
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (d) {
            return Data_Functor.map(Functor0)(function (v) {
                return {
                    keyword: d.keyword,
                    statements: v
                };
            })(Data_Traversable.traverse(Data_Array_NonEmpty_Internal.traversableNonEmptyArray)(dictApplicative)(traverseDoStatement(dictApplicative)(k))(d.statements));
        };
    };
};
var traverseDelimitedNonEmpty = function (dictApplicative) {
    return function (k) {
        return traverseWrapped(dictApplicative)(traverseSeparated(dictApplicative)(k));
    };
};
var traverseOneOrDelimited = function (dictApplicative) {
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (v) {
            if (v instanceof PureScript_CST_Types.One) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.One.create)(k(v.value0));
            };
            if (v instanceof PureScript_CST_Types.Many) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.Many.create)(traverseDelimitedNonEmpty(dictApplicative)(k)(v.value0));
            };
            throw new Error("Failed pattern match at PureScript.CST.Traversal (line 266, column 28 - line 268, column 55): " + [ v.constructor.name ]);
        };
    };
};
var traverseInstanceHead = function (dictApplicative) {
    var Apply0 = dictApplicative.Apply0();
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (head) {
            return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(function (v) {
                return function (v1) {
                    return {
                        className: head.className,
                        keyword: head.keyword,
                        name: head.name,
                        constraints: v,
                        types: v1
                    };
                };
            })(Data_Traversable.traverse(Data_Traversable.traversableMaybe)(dictApplicative)(Data_Bitraversable.ltraverse(Data_Bitraversable.bitraversableTuple)(dictApplicative)(traverseOneOrDelimited(dictApplicative)(k.onType)))(head.constraints)))(Data_Traversable.traverse(Data_Traversable.traversableArray)(dictApplicative)(k.onType)(head.types));
        };
    };
};
var traverseInstance = function (dictApplicative) {
    var Apply0 = dictApplicative.Apply0();
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (v) {
            return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(function (head) {
                return function (body) {
                    return {
                        head: head,
                        body: body
                    };
                };
            })(traverseInstanceHead(dictApplicative)(k)(v.head)))(Data_Traversable.traverse(Data_Traversable.traversableMaybe)(dictApplicative)(Data_Traversable.traverse(Data_Traversable.traversableTuple)(dictApplicative)(Data_Traversable.traverse(Data_Array_NonEmpty_Internal.traversableNonEmptyArray)(dictApplicative)(traverseInstanceBinding(dictApplicative)(k))))(v.body));
        };
    };
};
var traverseDelimited = function (dictApplicative) {
    return function (k) {
        return traverseWrapped(dictApplicative)(Data_Traversable.traverse(Data_Traversable.traversableMaybe)(dictApplicative)(traverseSeparated(dictApplicative)(k)));
    };
};
var traverseDataHead = function (dictApplicative) {
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (head) {
            return Data_Functor.map(Functor0)(function (v) {
                return {
                    keyword: head.keyword,
                    name: head.name,
                    vars: v
                };
            })(Data_Traversable.traverse(Data_Traversable.traversableArray)(dictApplicative)(traverseTypeVarBinding(dictApplicative)(k))(head.vars));
        };
    };
};
var traverseDataCtor = function (dictApplicative) {
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (v) {
            return Data_Functor.map(Functor0)(function (fields) {
                return {
                    name: v.name,
                    fields: fields
                };
            })(Data_Traversable.traverse(Data_Traversable.traversableArray)(dictApplicative)(k.onType)(v.fields));
        };
    };
};
var traverseClassHead = function (dictApplicative) {
    var Apply0 = dictApplicative.Apply0();
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (head) {
            return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(function (v) {
                return function (v1) {
                    return {
                        fundeps: head.fundeps,
                        keyword: head.keyword,
                        name: head.name,
                        "super": v,
                        vars: v1
                    };
                };
            })(Data_Traversable.traverse(Data_Traversable.traversableMaybe)(dictApplicative)(Data_Bitraversable.ltraverse(Data_Bitraversable.bitraversableTuple)(dictApplicative)(traverseOneOrDelimited(dictApplicative)(k.onType)))(head["super"])))(Data_Traversable.traverse(Data_Traversable.traversableArray)(dictApplicative)(traverseTypeVarBinding(dictApplicative)(k))(head.vars));
        };
    };
};
var traverseDecl = function (dictApplicative) {
    var Apply0 = dictApplicative.Apply0();
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (v) {
            if (v instanceof PureScript_CST_Types.DeclData) {
                return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(PureScript_CST_Types.DeclData.create)(traverseDataHead(dictApplicative)(k)(v.value0)))(Data_Traversable.traverse(Data_Traversable.traversableMaybe)(dictApplicative)(Data_Traversable.traverse(Data_Traversable.traversableTuple)(dictApplicative)(traverseSeparated(dictApplicative)(traverseDataCtor(dictApplicative)(k))))(v.value1));
            };
            if (v instanceof PureScript_CST_Types.DeclType) {
                return Control_Apply.apply(Apply0)(Data_Functor.flap(Functor0)(Data_Functor.map(Functor0)(PureScript_CST_Types.DeclType.create)(traverseDataHead(dictApplicative)(k)(v.value0)))(v.value1))(k.onType(v.value2));
            };
            if (v instanceof PureScript_CST_Types.DeclNewtype) {
                return Control_Apply.apply(Apply0)(Data_Functor.flap(Functor0)(Data_Functor.flap(Functor0)(Data_Functor.map(Functor0)(PureScript_CST_Types.DeclNewtype.create)(traverseDataHead(dictApplicative)(k)(v.value0)))(v.value1))(v.value2))(k.onType(v.value3));
            };
            if (v instanceof PureScript_CST_Types.DeclClass) {
                return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(PureScript_CST_Types.DeclClass.create)(traverseClassHead(dictApplicative)(k)(v.value0)))(Data_Traversable.traverse(Data_Traversable.traversableMaybe)(dictApplicative)(Data_Traversable.traverse(Data_Traversable.traversableTuple)(dictApplicative)(Data_Traversable.traverse(Data_Array_NonEmpty_Internal.traversableNonEmptyArray)(dictApplicative)(traverseLabeled(dictApplicative)(k.onType))))(v.value1));
            };
            if (v instanceof PureScript_CST_Types.DeclInstanceChain) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.DeclInstanceChain.create)(traverseSeparated(dictApplicative)(traverseInstance(dictApplicative)(k))(v.value0));
            };
            if (v instanceof PureScript_CST_Types.DeclDerive) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.DeclDerive.create(v.value0)(v.value1))(traverseInstanceHead(dictApplicative)(k)(v.value2));
            };
            if (v instanceof PureScript_CST_Types.DeclKindSignature) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.DeclKindSignature.create(v.value0))(traverseLabeled(dictApplicative)(k.onType)(v.value1));
            };
            if (v instanceof PureScript_CST_Types.DeclSignature) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.DeclSignature.create)(traverseLabeled(dictApplicative)(k.onType)(v.value0));
            };
            if (v instanceof PureScript_CST_Types.DeclValue) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.DeclValue.create)(traverseValueBindingFields(dictApplicative)(k)(v.value0));
            };
            if (v instanceof PureScript_CST_Types.DeclForeign) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.DeclForeign.create(v.value0)(v.value1))(traverseForeign(dictApplicative)(k)(v.value2));
            };
            return Control_Applicative.pure(dictApplicative)(v);
        };
    };
};
var traverseCaseOf = function (dictApplicative) {
    var Apply0 = dictApplicative.Apply0();
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (r) {
            return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(function (v) {
                return function (v1) {
                    return {
                        keyword: r.keyword,
                        of: r.of,
                        head: v,
                        branches: v1
                    };
                };
            })(traverseSeparated(dictApplicative)(k.onExpr)(r.head)))(Data_Traversable.traverse(Data_Array_NonEmpty_Internal.traversableNonEmptyArray)(dictApplicative)(Data_Bitraversable.bitraverse(Data_Bitraversable.bitraversableTuple)(dictApplicative)(traverseSeparated(dictApplicative)(k.onBinder))(traverseGuarded(dictApplicative)(k)))(r.branches));
        };
    };
};
var traverseBinder = function (dictApplicative) {
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    var Apply0 = dictApplicative.Apply0();
    return function (k) {
        return function (v) {
            if (v instanceof PureScript_CST_Types.BinderNamed) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.BinderNamed.create(v.value0)(v.value1))(k.onBinder(v.value2));
            };
            if (v instanceof PureScript_CST_Types.BinderConstructor) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.BinderConstructor.create(v.value0))(Data_Traversable.traverse(Data_Traversable.traversableArray)(dictApplicative)(k.onBinder)(v.value1));
            };
            if (v instanceof PureScript_CST_Types.BinderArray) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.BinderArray.create)(traverseDelimited(dictApplicative)(k.onBinder)(v.value0));
            };
            if (v instanceof PureScript_CST_Types.BinderRecord) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.BinderRecord.create)(traverseDelimited(dictApplicative)(traverseRecordLabeled(dictApplicative)(k.onBinder))(v.value0));
            };
            if (v instanceof PureScript_CST_Types.BinderParens) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.BinderParens.create)(traverseWrapped(dictApplicative)(k.onBinder)(v.value0));
            };
            if (v instanceof PureScript_CST_Types.BinderTyped) {
                return Control_Apply.apply(Apply0)(Data_Functor.flap(Functor0)(Data_Functor.map(Functor0)(PureScript_CST_Types.BinderTyped.create)(k.onBinder(v.value0)))(v.value1))(k.onType(v.value2));
            };
            if (v instanceof PureScript_CST_Types.BinderOp) {
                return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(PureScript_CST_Types.BinderOp.create)(k.onBinder(v.value0)))(Data_Traversable.traverse(Data_Array_NonEmpty_Internal.traversableNonEmptyArray)(dictApplicative)(Data_Traversable.traverse(Data_Traversable.traversableTuple)(dictApplicative)(k.onBinder))(v.value1));
            };
            return Control_Applicative.pure(dictApplicative)(v);
        };
    };
};
var traverseAdoBlock = function (dictApplicative) {
    var Apply0 = dictApplicative.Apply0();
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    return function (k) {
        return function (a) {
            return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(function (v) {
                return function (v1) {
                    return {
                        "in": a["in"],
                        keyword: a.keyword,
                        statements: v,
                        result: v1
                    };
                };
            })(Data_Traversable.traverse(Data_Traversable.traversableArray)(dictApplicative)(traverseDoStatement(dictApplicative)(k))(a.statements)))(k.onExpr(a.result));
        };
    };
};
var traverseExpr = function (dictApplicative) {
    var Functor0 = (dictApplicative.Apply0()).Functor0();
    var Apply0 = dictApplicative.Apply0();
    return function (k) {
        return function (v) {
            if (v instanceof PureScript_CST_Types.ExprArray) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.ExprArray.create)(traverseDelimited(dictApplicative)(k.onExpr)(v.value0));
            };
            if (v instanceof PureScript_CST_Types.ExprRecord) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.ExprRecord.create)(traverseDelimited(dictApplicative)(traverseRecordLabeled(dictApplicative)(k.onExpr))(v.value0));
            };
            if (v instanceof PureScript_CST_Types.ExprParens) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.ExprParens.create)(traverseWrapped(dictApplicative)(k.onExpr)(v.value0));
            };
            if (v instanceof PureScript_CST_Types.ExprTyped) {
                return Control_Apply.apply(Apply0)(Data_Functor.flap(Functor0)(Data_Functor.map(Functor0)(PureScript_CST_Types.ExprTyped.create)(k.onExpr(v.value0)))(v.value1))(k.onType(v.value2));
            };
            if (v instanceof PureScript_CST_Types.ExprInfix) {
                return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(PureScript_CST_Types.ExprInfix.create)(k.onExpr(v.value0)))(Data_Traversable.traverse(Data_Array_NonEmpty_Internal.traversableNonEmptyArray)(dictApplicative)(Data_Bitraversable.bitraverse(Data_Bitraversable.bitraversableTuple)(dictApplicative)(traverseWrapped(dictApplicative)(k.onExpr))(k.onExpr))(v.value1));
            };
            if (v instanceof PureScript_CST_Types.ExprOp) {
                return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(PureScript_CST_Types.ExprOp.create)(k.onExpr(v.value0)))(Data_Traversable.traverse(Data_Array_NonEmpty_Internal.traversableNonEmptyArray)(dictApplicative)(Data_Traversable.traverse(Data_Traversable.traversableTuple)(dictApplicative)(k.onExpr))(v.value1));
            };
            if (v instanceof PureScript_CST_Types.ExprNegate) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.ExprNegate.create(v.value0))(k.onExpr(v.value1));
            };
            if (v instanceof PureScript_CST_Types.ExprRecordAccessor) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.ExprRecordAccessor.create)(traverseRecordAccessor(dictApplicative)(k)(v.value0));
            };
            if (v instanceof PureScript_CST_Types.ExprRecordUpdate) {
                return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(PureScript_CST_Types.ExprRecordUpdate.create)(k.onExpr(v.value0)))(traverseWrapped(dictApplicative)(traverseSeparated(dictApplicative)(traverseRecordUpdate(dictApplicative)(k)))(v.value1));
            };
            if (v instanceof PureScript_CST_Types.ExprApp) {
                return Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(PureScript_CST_Types.ExprApp.create)(k.onExpr(v.value0)))(Data_Traversable.traverse(Data_Array_NonEmpty_Internal.traversableNonEmptyArray)(dictApplicative)(traverseExprAppSpine(dictApplicative)(k))(v.value1));
            };
            if (v instanceof PureScript_CST_Types.ExprLambda) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.ExprLambda.create)(traverseLambda(dictApplicative)(k)(v.value0));
            };
            if (v instanceof PureScript_CST_Types.ExprIf) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.ExprIf.create)(traverseIfThenElse(dictApplicative)(k)(v.value0));
            };
            if (v instanceof PureScript_CST_Types.ExprCase) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.ExprCase.create)(traverseCaseOf(dictApplicative)(k)(v.value0));
            };
            if (v instanceof PureScript_CST_Types.ExprLet) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.ExprLet.create)(traverseLetIn(dictApplicative)(k)(v.value0));
            };
            if (v instanceof PureScript_CST_Types.ExprDo) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.ExprDo.create)(traverseDoBlock(dictApplicative)(k)(v.value0));
            };
            if (v instanceof PureScript_CST_Types.ExprAdo) {
                return Data_Functor.map(Functor0)(PureScript_CST_Types.ExprAdo.create)(traverseAdoBlock(dictApplicative)(k)(v.value0));
            };
            return Control_Applicative.pure(dictApplicative)(v);
        };
    };
};
var topDownTraversalWithContextM = function (dictMonad) {
    var Bind1 = dictMonad.Bind1();
    var applicativeReaderT1 = Control_Monad_Reader_Trans.applicativeReaderT(dictMonad.Applicative0());
    return function (visitor) {
        var visitor$prime = {
            onBinder: function (a) {
                return function (ctx) {
                    return Control_Bind.bind(Bind1)(visitor.onBinder(ctx)(a))(Data_Tuple.uncurry(Data_Function.flip((function () {
                        var $479 = traverseBinder(applicativeReaderT1)(visitor$prime);
                        return function ($480) {
                            return Control_Monad_Reader_Trans.runReaderT($479($480));
                        };
                    })())));
                };
            },
            onExpr: function (a) {
                return function (ctx) {
                    return Control_Bind.bind(Bind1)(visitor.onExpr(ctx)(a))(Data_Tuple.uncurry(Data_Function.flip((function () {
                        var $481 = traverseExpr(applicativeReaderT1)(visitor$prime);
                        return function ($482) {
                            return Control_Monad_Reader_Trans.runReaderT($481($482));
                        };
                    })())));
                };
            },
            onDecl: function (a) {
                return function (ctx) {
                    return Control_Bind.bind(Bind1)(visitor.onDecl(ctx)(a))(Data_Tuple.uncurry(Data_Function.flip((function () {
                        var $483 = traverseDecl(applicativeReaderT1)(visitor$prime);
                        return function ($484) {
                            return Control_Monad_Reader_Trans.runReaderT($483($484));
                        };
                    })())));
                };
            },
            onType: function (a) {
                return function (ctx) {
                    return Control_Bind.bind(Bind1)(visitor.onType(ctx)(a))(Data_Tuple.uncurry(Data_Function.flip((function () {
                        var $485 = traverseType(applicativeReaderT1)(visitor$prime);
                        return function ($486) {
                            return Control_Monad_Reader_Trans.runReaderT($485($486));
                        };
                    })())));
                };
            }
        };
        return visitor$prime;
    };
};
var topDownTraversalWithContext = function (visitor) {
    var visitor$prime = {
        onBinder: function (a) {
            return function (ctx) {
                return Control_Bind.bind(Data_Identity.bindIdentity)(Control_Applicative.pure(Data_Identity.applicativeIdentity)(visitor.onBinder(ctx)(a)))(Data_Tuple.uncurry(Data_Function.flip((function () {
                    var $487 = traverseBinder(applicativeReaderT)(visitor$prime);
                    return function ($488) {
                        return Control_Monad_Reader_Trans.runReaderT($487($488));
                    };
                })())));
            };
        },
        onExpr: function (a) {
            return function (ctx) {
                return Control_Bind.bind(Data_Identity.bindIdentity)(Control_Applicative.pure(Data_Identity.applicativeIdentity)(visitor.onExpr(ctx)(a)))(Data_Tuple.uncurry(Data_Function.flip((function () {
                    var $489 = traverseExpr(applicativeReaderT)(visitor$prime);
                    return function ($490) {
                        return Control_Monad_Reader_Trans.runReaderT($489($490));
                    };
                })())));
            };
        },
        onDecl: function (a) {
            return function (ctx) {
                return Control_Bind.bind(Data_Identity.bindIdentity)(Control_Applicative.pure(Data_Identity.applicativeIdentity)(visitor.onDecl(ctx)(a)))(Data_Tuple.uncurry(Data_Function.flip((function () {
                    var $491 = traverseDecl(applicativeReaderT)(visitor$prime);
                    return function ($492) {
                        return Control_Monad_Reader_Trans.runReaderT($491($492));
                    };
                })())));
            };
        },
        onType: function (a) {
            return function (ctx) {
                return Control_Bind.bind(Data_Identity.bindIdentity)(Control_Applicative.pure(Data_Identity.applicativeIdentity)(visitor.onType(ctx)(a)))(Data_Tuple.uncurry(Data_Function.flip((function () {
                    var $493 = traverseType(applicativeReaderT)(visitor$prime);
                    return function ($494) {
                        return Control_Monad_Reader_Trans.runReaderT($493($494));
                    };
                })())));
            };
        }
    };
    return visitor$prime;
};
var topDownTraversal = function (dictMonad) {
    var Bind1 = dictMonad.Bind1();
    var Applicative0 = dictMonad.Applicative0();
    return function (visitor) {
        var visitor$prime = {
            onBinder: function (a) {
                return Control_Bind.bind(Bind1)(visitor.onBinder(a))(traverseBinder(Applicative0)(visitor$prime));
            },
            onExpr: function (a) {
                return Control_Bind.bind(Bind1)(visitor.onExpr(a))(traverseExpr(Applicative0)(visitor$prime));
            },
            onType: function (a) {
                return Control_Bind.bind(Bind1)(visitor.onType(a))(traverseType(Applicative0)(visitor$prime));
            },
            onDecl: function (a) {
                return Control_Bind.bind(Bind1)(visitor.onDecl(a))(traverseDecl(Applicative0)(visitor$prime));
            }
        };
        return visitor$prime;
    };
};
var topDownPureTraversal = function (visitor) {
    var visitor$prime = {
        onBinder: function (a) {
            return Control_Bind.bind(Control_Monad_Free.freeBind)(Control_Applicative.pure(Control_Monad_Free.freeApplicative)(visitor.onBinder(a)))(traverseBinder(Control_Monad_Free.freeApplicative)(visitor$prime));
        },
        onExpr: function (a) {
            return Control_Bind.bind(Control_Monad_Free.freeBind)(Control_Applicative.pure(Control_Monad_Free.freeApplicative)(visitor.onExpr(a)))(traverseExpr(Control_Monad_Free.freeApplicative)(visitor$prime));
        },
        onType: function (a) {
            return Control_Bind.bind(Control_Monad_Free.freeBind)(Control_Applicative.pure(Control_Monad_Free.freeApplicative)(visitor.onType(a)))(traverseType(Control_Monad_Free.freeApplicative)(visitor$prime));
        },
        onDecl: function (a) {
            return Control_Bind.bind(Control_Monad_Free.freeBind)(Control_Applicative.pure(Control_Monad_Free.freeApplicative)(visitor.onDecl(a)))(traverseDecl(Control_Monad_Free.freeApplicative)(visitor$prime));
        }
    };
    return visitor$prime;
};
var rewriteWithContextM = function (dictMonad) {
    var Functor0 = ((dictMonad.Bind1()).Apply0()).Functor0();
    return function (traversal) {
        return function (visitor) {
            return function (ctx) {
                return function (g) {
                    var visitor$prime = topDownTraversalWithContextM(dictMonad)(visitor);
                    return Data_Functor.map(Functor0)(Data_Tuple.Tuple.create(ctx))(Control_Monad_Reader_Trans.runReaderT(traversal(visitor$prime)(g))(ctx));
                };
            };
        };
    };
};
var rewriteWithContext = function (traversal) {
    return function (visitor) {
        return function (ctx) {
            return function (g) {
                var visitor$prime = topDownTraversalWithContext(visitor);
                return new Data_Tuple.Tuple(ctx, Data_Newtype.un()(Data_Identity.Identity)(Control_Monad_Reader_Trans.runReaderT(traversal(visitor$prime)(g))(ctx)));
            };
        };
    };
};
var rewriteTypeWithContextM = function (dictMonad) {
    return rewriteWithContextM(dictMonad)(function (v) {
        return v.onType;
    });
};
var rewriteTypeWithContext = /* #__PURE__ */ rewriteWithContext(function (v) {
    return v.onType;
});
var rewriteTopDownM = function (dictMonad) {
    return function (traversal) {
        return function (visitor) {
            var visitor$prime = topDownTraversal(dictMonad)(visitor);
            return traversal(visitor$prime);
        };
    };
};
var rewriteTypeTopDownM = function (dictMonad) {
    return rewriteTopDownM(dictMonad)(function (v) {
        return v.onType;
    });
};
var rewriteTopDown = function (traversal) {
    return function (visitor) {
        var visitor$prime = topDownPureTraversal(visitor);
        var $495 = Control_Monad_Free.runFree(Data_Identity.functorIdentity)(Data_Newtype.un()(Data_Identity.Identity));
        var $496 = traversal(visitor$prime);
        return function ($497) {
            return $495($496($497));
        };
    };
};
var rewriteTypeTopDown = /* #__PURE__ */ rewriteTopDown(function (v) {
    return v.onType;
});
var rewriteModuleWithContextM = function (dictMonad) {
    return rewriteWithContextM(dictMonad)(traverseModule(Control_Monad_Reader_Trans.applicativeReaderT(dictMonad.Applicative0())));
};
var rewriteModuleWithContext = /* #__PURE__ */ rewriteWithContext(/* #__PURE__ */ traverseModule(/* #__PURE__ */ Control_Monad_Reader_Trans.applicativeReaderT(Data_Identity.applicativeIdentity)));
var rewriteModuleTopDownM = function (dictMonad) {
    return rewriteTopDownM(dictMonad)(traverseModule(dictMonad.Applicative0()));
};
var rewriteModuleTopDown = /* #__PURE__ */ rewriteTopDown(traverseModule1);
var rewriteExprWithContextM = function (dictMonad) {
    return rewriteWithContextM(dictMonad)(function (v) {
        return v.onExpr;
    });
};
var rewriteExprWithContext = /* #__PURE__ */ rewriteWithContext(function (v) {
    return v.onExpr;
});
var rewriteExprTopDownM = function (dictMonad) {
    return rewriteTopDownM(dictMonad)(function (v) {
        return v.onExpr;
    });
};
var rewriteExprTopDown = /* #__PURE__ */ rewriteTopDown(function (v) {
    return v.onExpr;
});
var rewriteDeclWithContextM = function (dictMonad) {
    return rewriteWithContextM(dictMonad)(function (v) {
        return v.onDecl;
    });
};
var rewriteDeclWithContext = /* #__PURE__ */ rewriteWithContext(function (v) {
    return v.onDecl;
});
var rewriteDeclTopDownM = function (dictMonad) {
    return rewriteTopDownM(dictMonad)(function (v) {
        return v.onDecl;
    });
};
var rewriteDeclTopDown = /* #__PURE__ */ rewriteTopDown(function (v) {
    return v.onDecl;
});
var rewriteBinderWithContextM = function (dictMonad) {
    return rewriteWithContextM(dictMonad)(function (v) {
        return v.onBinder;
    });
};
var rewriteBinderWithContext = /* #__PURE__ */ rewriteWithContext(function (v) {
    return v.onBinder;
});
var rewriteBinderTopDownM = function (dictMonad) {
    return rewriteTopDownM(dictMonad)(function (v) {
        return v.onBinder;
    });
};
var rewriteBinderTopDown = /* #__PURE__ */ rewriteTopDown(function (v) {
    return v.onBinder;
});
var defer = function (dictMonad) {
    var Bind1 = dictMonad.Bind1();
    var Applicative0 = dictMonad.Applicative0();
    return function (v) {
        return Control_Bind.bind(Bind1)(Control_Applicative.pure(Applicative0)(Data_Unit.unit))(v);
    };
};
var topDownMonoidalTraversal = function (dictMonoid) {
    var applyCompose1 = applyCompose(Data_Const.applyConst(dictMonoid.Semigroup0()));
    var applicativeCompose1 = applicativeCompose(Data_Const.applicativeConst(dictMonoid));
    return function (visitor) {
        var visitor$prime = {
            onBinder: function (a) {
                return Control_Apply.apply(applyCompose1)(Control_Applicative.pure(Control_Monad_Free.freeApplicative)(visitor.onBinder(a)))(defer(Control_Monad_Free.freeMonad)(function (v) {
                    return Data_Newtype.un()(Data_Functor_Compose.Compose)(traverseBinder(applicativeCompose1)(visitor$prime)(a));
                }));
            },
            onExpr: function (a) {
                return Control_Apply.apply(applyCompose1)(Control_Applicative.pure(Control_Monad_Free.freeApplicative)(visitor.onExpr(a)))(defer(Control_Monad_Free.freeMonad)(function (v) {
                    return Data_Newtype.un()(Data_Functor_Compose.Compose)(traverseExpr(applicativeCompose1)(visitor$prime)(a));
                }));
            },
            onDecl: function (a) {
                return Control_Apply.apply(applyCompose1)(Control_Applicative.pure(Control_Monad_Free.freeApplicative)(visitor.onDecl(a)))(defer(Control_Monad_Free.freeMonad)(function (v) {
                    return Data_Newtype.un()(Data_Functor_Compose.Compose)(traverseDecl(applicativeCompose1)(visitor$prime)(a));
                }));
            },
            onType: function (a) {
                return Control_Apply.apply(applyCompose1)(Control_Applicative.pure(Control_Monad_Free.freeApplicative)(visitor.onType(a)))(defer(Control_Monad_Free.freeMonad)(function (v) {
                    return Data_Newtype.un()(Data_Functor_Compose.Compose)(traverseType(applicativeCompose1)(visitor$prime)(a));
                }));
            }
        };
        return visitor$prime;
    };
};
var monoidalRewrite = function (dictMonoid) {
    return function (traversal) {
        return function (visitor) {
            return function (g) {
                var visitor$prime = topDownMonoidalTraversal(dictMonoid)(visitor);
                return Data_Newtype.un()(Data_Const.Const)(Control_Monad_Free.runFree(Data_Identity.functorIdentity)(Data_Newtype.un()(Data_Identity.Identity))(Data_Newtype.un()(Data_Functor_Compose.Compose)(traversal(visitor$prime)(g))));
            };
        };
    };
};
var foldMapBinder = function (dictMonoid) {
    return monoidalRewrite(dictMonoid)(function (v) {
        return v.onBinder;
    });
};
var foldMapDecl = function (dictMonoid) {
    return monoidalRewrite(dictMonoid)(function (v) {
        return v.onDecl;
    });
};
var foldMapExpr = function (dictMonoid) {
    return monoidalRewrite(dictMonoid)(function (v) {
        return v.onExpr;
    });
};
var foldMapModule = function (dictMonoid) {
    return monoidalRewrite(dictMonoid)(traverseModule(applicativeCompose(Data_Const.applicativeConst(dictMonoid))));
};
var foldMapType = function (dictMonoid) {
    return monoidalRewrite(dictMonoid)(function (v) {
        return v.onType;
    });
};
var defaultVisitorWithContextM = function (dictMonad) {
    var Applicative0 = dictMonad.Applicative0();
    return {
        onBinder: Data_Tuple.curry(Control_Applicative.pure(Applicative0)),
        onDecl: Data_Tuple.curry(Control_Applicative.pure(Applicative0)),
        onExpr: Data_Tuple.curry(Control_Applicative.pure(Applicative0)),
        onType: Data_Tuple.curry(Control_Applicative.pure(Applicative0))
    };
};
var defaultVisitorWithContext = {
    onBinder: /* #__PURE__ */ Data_Tuple.curry(/* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn)),
    onDecl: /* #__PURE__ */ Data_Tuple.curry(/* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn)),
    onExpr: /* #__PURE__ */ Data_Tuple.curry(/* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn)),
    onType: /* #__PURE__ */ Data_Tuple.curry(/* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn))
};
var defaultVisitorM = function (dictApplicative) {
    return {
        onBinder: Control_Applicative.pure(dictApplicative),
        onDecl: Control_Applicative.pure(dictApplicative),
        onExpr: Control_Applicative.pure(dictApplicative),
        onType: Control_Applicative.pure(dictApplicative)
    };
};
var defaultVisitor = {
    onBinder: /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn),
    onDecl: /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn),
    onExpr: /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn),
    onType: /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn)
};
var defaultMonoidalVisitor = function (dictMonoid) {
    var monoidFn = Data_Monoid.monoidFn(dictMonoid);
    return {
        onBinder: Data_Monoid.mempty(monoidFn),
        onDecl: Data_Monoid.mempty(monoidFn),
        onExpr: Data_Monoid.mempty(monoidFn),
        onType: Data_Monoid.mempty(monoidFn)
    };
};
var bottomUpTraversal = function (dictMonad) {
    var Bind1 = dictMonad.Bind1();
    var Applicative0 = dictMonad.Applicative0();
    return function (visitor) {
        var visitor$prime = {
            onBinder: function (a) {
                return Control_Bind.bindFlipped(Bind1)(visitor.onBinder)(defer(dictMonad)(function (v) {
                    return traverseBinder(Applicative0)(visitor$prime)(a);
                }));
            },
            onExpr: function (a) {
                return Control_Bind.bindFlipped(Bind1)(visitor.onExpr)(defer(dictMonad)(function (v) {
                    return traverseExpr(Applicative0)(visitor$prime)(a);
                }));
            },
            onType: function (a) {
                return Control_Bind.bindFlipped(Bind1)(visitor.onType)(defer(dictMonad)(function (v) {
                    return traverseType(Applicative0)(visitor$prime)(a);
                }));
            },
            onDecl: function (a) {
                return Control_Bind.bindFlipped(Bind1)(visitor.onDecl)(defer(dictMonad)(function (v) {
                    return traverseDecl(Applicative0)(visitor$prime)(a);
                }));
            }
        };
        return visitor$prime;
    };
};
var rewriteBottomUpM = function (dictMonad) {
    return function (traversal) {
        return function (visitor) {
            var visitor$prime = bottomUpTraversal(dictMonad)(visitor);
            return traversal(visitor$prime);
        };
    };
};
var rewriteBinderBottomUpM = function (dictMonad) {
    return rewriteBottomUpM(dictMonad)(function (v) {
        return v.onBinder;
    });
};
var rewriteDeclBottomUpM = function (dictMonad) {
    return rewriteBottomUpM(dictMonad)(function (v) {
        return v.onDecl;
    });
};
var rewriteExprBottomUpM = function (dictMonad) {
    return rewriteBottomUpM(dictMonad)(function (v) {
        return v.onExpr;
    });
};
var rewriteModuleBottomUpM = function (dictMonad) {
    return rewriteBottomUpM(dictMonad)(traverseModule(dictMonad.Applicative0()));
};
var rewriteTypeBottomUpM = function (dictMonad) {
    return rewriteBottomUpM(dictMonad)(function (v) {
        return v.onType;
    });
};
var bottomUpPureTraversal = function (visitor) {
    var visitor$prime = {
        onBinder: function (a) {
            return Control_Bind.bindFlipped(Control_Monad_Free.freeBind)(function ($498) {
                return pure(visitor.onBinder($498));
            })(traverseBinder(Control_Monad_Free.freeApplicative)(visitor$prime)(a));
        },
        onExpr: function (a) {
            return Control_Bind.bindFlipped(Control_Monad_Free.freeBind)(function ($499) {
                return pure1(visitor.onExpr($499));
            })(traverseExpr(Control_Monad_Free.freeApplicative)(visitor$prime)(a));
        },
        onType: function (a) {
            return Control_Bind.bindFlipped(Control_Monad_Free.freeBind)(function ($500) {
                return pure2(visitor.onType($500));
            })(traverseType(Control_Monad_Free.freeApplicative)(visitor$prime)(a));
        },
        onDecl: function (a) {
            return Control_Bind.bindFlipped(Control_Monad_Free.freeBind)(function ($501) {
                return pure3(visitor.onDecl($501));
            })(traverseDecl(Control_Monad_Free.freeApplicative)(visitor$prime)(a));
        }
    };
    return visitor$prime;
};
var rewriteBottomUp = function (traversal) {
    return function (visitor) {
        var visitor$prime = bottomUpPureTraversal(visitor);
        var $502 = Control_Monad_Free.runFree(Data_Identity.functorIdentity)(Data_Newtype.un()(Data_Identity.Identity));
        var $503 = traversal(visitor$prime);
        return function ($504) {
            return $502($503($504));
        };
    };
};
var rewriteBinderBottomUp = /* #__PURE__ */ rewriteBottomUp(function (v) {
    return v.onBinder;
});
var rewriteDeclBottomUp = /* #__PURE__ */ rewriteBottomUp(function (v) {
    return v.onDecl;
});
var rewriteExprBottomUp = /* #__PURE__ */ rewriteBottomUp(function (v) {
    return v.onExpr;
});
var rewriteModuleBottomUp = /* #__PURE__ */ rewriteBottomUp(traverseModule1);
var rewriteTypeBottomUp = /* #__PURE__ */ rewriteBottomUp(function (v) {
    return v.onType;
});
export {
    defaultVisitorM,
    rewriteModuleBottomUpM,
    rewriteBinderBottomUpM,
    rewriteExprBottomUpM,
    rewriteDeclBottomUpM,
    rewriteTypeBottomUpM,
    rewriteModuleTopDownM,
    rewriteBinderTopDownM,
    rewriteExprTopDownM,
    rewriteDeclTopDownM,
    rewriteTypeTopDownM,
    defaultVisitorWithContextM,
    rewriteModuleWithContextM,
    rewriteBinderWithContextM,
    rewriteExprWithContextM,
    rewriteDeclWithContextM,
    rewriteTypeWithContextM,
    defaultMonoidalVisitor,
    foldMapModule,
    foldMapBinder,
    foldMapDecl,
    foldMapExpr,
    foldMapType,
    defaultVisitor,
    rewriteModuleBottomUp,
    rewriteBinderBottomUp,
    rewriteExprBottomUp,
    rewriteDeclBottomUp,
    rewriteTypeBottomUp,
    rewriteModuleTopDown,
    rewriteBinderTopDown,
    rewriteExprTopDown,
    rewriteDeclTopDown,
    rewriteTypeTopDown,
    defaultVisitorWithContext,
    rewriteModuleWithContext,
    rewriteBinderWithContext,
    rewriteExprWithContext,
    rewriteDeclWithContext,
    rewriteTypeWithContext,
    traverseModule,
    traverseModuleBody,
    traverseDecl,
    traverseForeign,
    traverseInstance,
    traverseInstanceHead,
    traverseInstanceBinding,
    traverseClassHead,
    traverseOneOrDelimited,
    traverseDataHead,
    traverseDataCtor,
    traverseType,
    traverseRow,
    traverseTypeVarBinding,
    traverseExpr,
    traverseExprAppSpine,
    traverseDelimited,
    traverseDelimitedNonEmpty,
    traverseSeparated,
    traverseWrapped,
    traverseRecordLabeled,
    traverseLabeled,
    traverseRecordAccessor,
    traverseRecordUpdate,
    traverseLambda,
    traverseIfThenElse,
    traverseCaseOf,
    traverseGuarded,
    traverseGuardedExpr,
    traversePatternGuard,
    traverseWhere,
    traverseLetBinding,
    traverseValueBindingFields,
    traverseLetIn,
    traverseDoStatement,
    traverseDoBlock,
    traverseAdoBlock,
    traverseBinder,
    bottomUpTraversal,
    rewriteBottomUpM,
    topDownTraversal,
    rewriteTopDownM,
    topDownTraversalWithContextM,
    rewriteWithContextM,
    topDownMonoidalTraversal,
    monoidalRewrite,
    bottomUpPureTraversal,
    rewriteBottomUp,
    topDownPureTraversal,
    rewriteTopDown,
    topDownTraversalWithContext,
    rewriteWithContext
};
