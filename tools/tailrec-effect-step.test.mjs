import assert from 'node:assert/strict';
import test from 'node:test';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { empty as emptyMap, insert } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import { TcoExpr } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as CodeGen from '../output/Gopurs.CodeGen/index.js';
import { recognize } from '../output/Gopurs.TailRecEffectStep/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';
import { withReboxFields } from './codegen-metadata.mjs';

const metadata = withReboxFields({ elidedCtors: emptySet,
    ctorTypes: insert(ordString)('Control.Monad.Rec.Class.Loop')({ vars: ['a','b'], fields: [new C.TypeVar('a')] })(
        insert(ordString)('Control.Monad.Rec.Class.Done')({ vars: ['a','b'], fields: [new C.TypeVar('b')] })(emptyMap)),
    pointerAdtPaths: emptyMap, pointerAdtNodes: emptySet, pointerAdtLeaves: emptyMap,
    enumAdts: emptySet, enumCtors: emptySet, globalTypes: emptyMap, globalFunctions: emptyMap, classDeclsFields: emptyMap });
const stepType = new C.ADT('Control.Monad.Rec.Class.Step', ['Control','Monad','Rec','Class','Step'], [C.Int.value,C.Int.value]);
function syntax(make) {
    const typed = (type, value) => make(new S.Typed(type, value));
    const int = value => typed(C.Int.value, make(new S.Lit(new C.LitInt(value))));
    const local = (name, level) => typed(C.Int.value, make(new S.Local(new Just(name), level)));
    const ref = (module, name) => make(new S.Var(new C.Qualified(new Just(module), name)));
    const step = (name, value, owner='Control.Monad.Rec.Class') => typed(stepType, make(new S.CtorSaturated(
        new C.Qualified(new Just(owner), name), C.SumType.value, 'Step', name, [new Tuple('value0',value)])));
    const pure = value => make(new S.EffectPure(value));
    const lambda = body => typed(new C.Func([C.Int.value], C.Any.value), make(new S.Abs([new Tuple(new Just('state'),0)], body)));
    return { typed, int, local, ref, step, pure, lambda };
}

test('inline step requires a local unary lambda and canonical exposed terminal constructors', () => {
    const make = syntax => new TcoExpr(null, syntax), s=syntax(make);
    const context = { metadata, modNameStr:'Probe' };
    const done=s.pure(s.step('Done',s.int(7)));
    assert.ok(recognize(context)(s.lambda(done)) instanceof Just);
    for(const callback of [s.ref('Probe','callback'), s.lambda(s.ref('Probe','action')),
        s.lambda(s.pure(s.ref('Probe','step'))), s.lambda(s.pure(s.step('Done',s.int(7),'Other'))),
        s.lambda(s.pure(s.step('Other',s.int(7)))),
        s.typed(new C.Func([C.Int.value,C.Int.value],C.Any.value),make(new S.Abs([new Tuple(new Just('x'),0),new Tuple(new Just('y'),1)],done))),
        s.lambda(make(new S.Branch([new S.Pair(make(new S.Lit(new C.LitBoolean(true))),done)],s.ref('Probe','opaqueAction'))))]) {
        assert.deepEqual(recognize(context)(callback),Nothing.value);
    }
});

function generatedFixture() {
    const make = syntax => new NeutralExpr(syntax), s=syntax(make), state=s.local('state',0);
    const op=(operator,a,b,type=C.Int.value)=>s.typed(type,make(new S.PrimOp(new S.Op2(operator,a,b))));
    const next=op(new S.OpIntNum(S.OpAdd.value),state,s.int(1));
    const captured=make(new S.Abs([new Tuple(new Just('ignored'),3)],state));
    const keep=make(new S.App(s.ref('Probe','keep'),[captured]));
    const again=make(new S.EffectBind(new Just('kept'),2,keep,s.pure(s.step('Loop',next))));
    const branches=make(new S.Branch([new S.Pair(op(new S.OpIntOrd(S.OpGte.value),state,s.int(3),C.Boolean.value),s.pure(s.step('Done',state)))],make(new S.EffectDefer(again))));
    const body=make(new S.Let(new Just('marked'),1,s.typed(C.Int.value,make(new S.App(s.ref('Probe','markPure'),[state]))),branches));
    const callback=s.lambda(body);
    const bindings=[['effect','Control.Monad.Rec.Class','monadRecEffect'],['st','Control.Monad.ST.Internal','monadRecST']].map(([name,module,dictionary])=>{
        const call=make(new S.App(s.ref('Control.Monad.Rec.Class','tailRecM'),[s.ref(module,dictionary),callback,s.local('seed',4)]));
        return {recursive:false,bindings:[new Tuple(name,s.typed(new C.Func([C.Int.value],C.Any.value),make(new S.Abs([new Tuple(new Just('seed'),4)],call))))]};
    });
    return CodeGen.translate(metadata)({name:'InlineEffectFixture',bindings,comments:[],imports:emptySet,exports:emptySet,reExports:emptySet,
        dataTypes:emptyMap,dataDecls:[],classDecls:[],foreign:emptyMap,implementations:emptyMap,directives:emptyMap});
}

test('generated inline loops preserve deferred execution, replay, escaping captures and exceptions', t => {
    const code=generatedFixture();
    assert.match(code,/tailRecEffect_parameter_/);
    assert.doesNotMatch(code,/Constructor_Control_Monad_Rec_Class_(Loop|Done)/,'terminal Step values must not be allocated');
    assert.doesNotMatch(code,/tailRecEffect_step_/,'recognized callbacks run inside the loop');
    const directory=mkdtempSync(join(tmpdir(),'gopurs-inline-effect-'));t.after(()=>rmSync(directory,{recursive:true,force:true}));
    mkdirSync(join(directory,'gopurs_runtime'));
    writeFileSync(join(directory,'go.mod'),'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory,'gopurs_runtime/runtime.go'),runtimeGoCode);
    writeFileSync(join(directory,'fixture.go'),code.replace('package purescript','package main'));
    writeFileSync(join(directory,'main.go'),`package main
import("fmt";"reflect";r "gopurs/output/gopurs_runtime")
var trace []string
var saved []r.Value
var failing bool
var sentinel = &struct{}{}
func Get_Control_Monad_Rec_Class_monadRecEffect()r.Value{trace=append(trace,"dictionary");return r.Value{}}
func Get_Control_Monad_ST_Internal_monadRecST()r.Value{trace=append(trace,"dictionary");return r.Value{}}
func Get_Probe_markPure()r.Value{return r.Func(func(v r.Value)r.Value{trace=append(trace,fmt.Sprint("pure:",v.IntVal));return v})}
func Get_Probe_keep()r.Value{return r.Func(func(callback r.Value)r.Value{
    trace=append(trace,"construct")
    return r.Func(func(_ r.Value)r.Value{value:=r.Apply(callback,r.Value{});trace=append(trace,fmt.Sprint("effect:",value.IntVal));saved=append(saved,callback);if failing&&value.IntVal==1{panic(sentinel)};return r.Value{}})
})}
func check(ok bool){if !ok{panic(fmt.Sprint("unexpected trace: ",trace))}}
func main(){
    for _,makeAction:=range []func(int64)r.Value{Call_InlineEffectFixture_effect,Call_InlineEffectFixture_st}{
        trace=nil;action:=makeAction(0);check(reflect.DeepEqual(trace,[]string{"dictionary"}))
        for repeat:=0;repeat<2;repeat++{
            trace=nil;saved=nil;check(r.Apply(action,r.Value{}).IntVal==3)
            check(reflect.DeepEqual(trace,[]string{"pure:0","construct","effect:0","pure:1","construct","effect:1","pure:2","construct","effect:2","pure:3"}))
            for i,f:=range saved{check(r.Apply(f,r.Value{}).IntVal==int64(i))}
        }
        trace=nil;immediate:=makeAction(9);trace=nil;check(r.Apply(immediate,r.Value{}).IntVal==9);check(reflect.DeepEqual(trace,[]string{"pure:9"}))
        failing=true;trace=nil;func(){defer func(){check(recover()==sentinel)}();r.Apply(action,r.Value{});panic("exception lost")}();failing=false
        check(reflect.DeepEqual(trace,[]string{"pure:0","construct","effect:0","pure:1","construct","effect:1"}))
    }
    fmt.Println("inline effect semantics passed")
}
`);
    const run=spawnSync('go',['run','.'],{cwd:directory,encoding:'utf8',timeout:60000,env:{...process.env,GOWORK:'off'}});
    assert.ifError(run.error);assert.equal(run.status,0,run.stdout+run.stderr);assert.match(run.stdout,/inline effect semantics passed/);
});
