import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Bind from "../Control.Bind/index.js";
import * as Control_Monad_State from "../Control.Monad.State/index.js";
import * as Control_Monad_State_Class from "../Control.Monad.State.Class/index.js";
import * as Control_Monad_State_Trans from "../Control.Monad.State.Trans/index.js";
import * as Data_Array from "../Data.Array/index.js";
import * as Data_Array_NonEmpty_Internal from "../Data.Array.NonEmpty.Internal/index.js";
import * as Data_Identity from "../Data.Identity/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Data_Traversable from "../Data.Traversable/index.js";
import * as PureScript_Backend_Optimizer_Codegen_Tco from "../PureScript.Backend.Optimizer.Codegen.Tco/index.js";
import * as PureScript_Backend_Optimizer_Syntax from "../PureScript.Backend.Optimizer.Syntax/index.js";
var bindStateT = /* #__PURE__ */ Control_Monad_State_Trans.bindStateT(Data_Identity.monadIdentity);
var monadStateStateT = /* #__PURE__ */ Control_Monad_State_Trans.monadStateStateT(Data_Identity.monadIdentity);
var get = /* #__PURE__ */ Control_Monad_State_Class.get(monadStateStateT);
var monadStateStateT1 = /* #__PURE__ */ Control_Monad_State_Trans.monadStateStateT(Data_Identity.monadIdentity);
var applicativeStateT = /* #__PURE__ */ Control_Monad_State_Trans.applicativeStateT(Data_Identity.monadIdentity);
var mempty = /* #__PURE__ */ Data_Monoid.mempty(PureScript_Backend_Optimizer_Codegen_Tco.monoidTcoAnalysis);
var bindStateT1 = /* #__PURE__ */ Control_Monad_State_Trans.bindStateT(Data_Identity.monadIdentity);
var applicativeStateT1 = /* #__PURE__ */ Control_Monad_State_Trans.applicativeStateT(Data_Identity.monadIdentity);
var get1 = /* #__PURE__ */ Control_Monad_State_Class.get(monadStateStateT);
var monadStateStateT2 = /* #__PURE__ */ Control_Monad_State_Trans.monadStateStateT(Data_Identity.monadIdentity);
var extractAccessorsTraverser = function (v) {
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Accessor) {
        return Control_Bind.bind(bindStateT)(extractAccessorsTraverser(v.value1.value0))(function (obj$prime) {
            return Control_Bind.bind(bindStateT)(get)(function (st) {
                var varName = "__fbip_proj_" + Data_Show.show(Data_Show.showInt)(st.nextId);
                var newBinding = {
                    ident: varName,
                    expr: new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.Accessor(obj$prime, v.value1.value1))
                };
                return Control_Bind.discard(Control_Bind.discardUnit)(bindStateT)(Control_Monad_State_Class.put(monadStateStateT1)({
                    nextId: st.nextId + 1 | 0,
                    bindings: Data_Array.snoc(st.bindings)(newBinding)
                }))(function () {
                    return Control_Applicative.pure(applicativeStateT)(new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(mempty, new PureScript_Backend_Optimizer_Syntax.Local(new Data_Maybe.Just(varName), 0)));
                });
            });
        });
    };
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Branch) {
        return Control_Applicative.pure(applicativeStateT)(v);
    };
    return Control_Bind.bind(bindStateT)(Data_Traversable.traverse(PureScript_Backend_Optimizer_Syntax.traversableBackendSyntax)(applicativeStateT)(extractAccessorsTraverser)(v.value1))(function (syn$prime) {
        return Control_Applicative.pure(applicativeStateT)(new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, syn$prime));
    });
};
var processExpr = function (v) {
    if (v.value1 instanceof PureScript_Backend_Optimizer_Syntax.Branch) {
        return Control_Bind.bind(bindStateT1)(Data_Traversable.traverse(Data_Array_NonEmpty_Internal.traversableNonEmptyArray)(applicativeStateT1)(function (v1) {
            return Control_Bind.bind(bindStateT1)(processExpr(v1.value0))(function (cond$prime) {
                return Control_Bind.bind(bindStateT1)(processBranchBody(v1.value1))(function (body$prime) {
                    return Control_Applicative.pure(applicativeStateT1)(new PureScript_Backend_Optimizer_Syntax.Pair(cond$prime, body$prime));
                });
            });
        })(v.value1.value0))(function (newBranches) {
            return Control_Bind.bind(bindStateT1)(processBranchBody(v.value1.value1))(function (def$prime) {
                return Control_Applicative.pure(applicativeStateT1)(new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, new PureScript_Backend_Optimizer_Syntax.Branch(newBranches, def$prime)));
            });
        });
    };
    return Control_Bind.bind(bindStateT1)(Data_Traversable.traverse(PureScript_Backend_Optimizer_Syntax.traversableBackendSyntax)(applicativeStateT1)(processExpr)(v.value1))(function (syn$prime) {
        return Control_Applicative.pure(applicativeStateT1)(new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(v.value0, syn$prime));
    });
};
var processBranchBody = function (body) {
    return Control_Bind.bind(bindStateT1)(get1)(function (startId) {
        var v = Control_Monad_State.runState(extractAccessorsTraverser(body))({
            nextId: startId,
            bindings: [  ]
        });
        return Control_Bind.discard(Control_Bind.discardUnit)(bindStateT1)(Control_Monad_State_Class.put(monadStateStateT2)(v.value1.nextId))(function () {
            var wrapped = Data_Array.foldr(function (b) {
                return function (acc) {
                    return new PureScript_Backend_Optimizer_Codegen_Tco.TcoExpr(mempty, new PureScript_Backend_Optimizer_Syntax.Let(new Data_Maybe.Just(b.ident), 0, b.expr, acc));
                };
            })(v.value0)(v.value1.bindings);
            return processExpr(wrapped);
        });
    });
};
var extractProjections = function (expr) {
    return Control_Monad_State.evalState(processExpr(expr))(0);
};
export {
    extractProjections
};
