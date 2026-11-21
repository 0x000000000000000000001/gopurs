import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { empty as emptyMap, insert } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import { emit } from '../output/Gopurs.TailRecEffect/index.js';
import { StmtLeaf, flattenStmts } from '../output/Gopurs.ExprContext/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';
import { TcoExpr } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { hashString } from '../output/PureScript.Backend.Optimizer.FfiSupport/index.js';
import { withReboxFields } from './codegen-metadata.mjs';

const constructors = insert(ordString)('Control.Monad.Rec.Class.Loop')({ vars: ['a', 'b'], fields: [new C.TypeVar('a')] })(
    insert(ordString)('Control.Monad.Rec.Class.Done')({ vars: ['a', 'b'], fields: [new C.TypeVar('b')] })(emptyMap));
const metadata = withReboxFields({
    elidedCtors: emptySet, ctorTypes: constructors, pointerAdtPaths: emptyMap,
    pointerAdtNodes: emptySet, pointerAdtLeaves: emptyMap, enumAdts: emptySet,
    enumCtors: emptySet, globalTypes: emptyMap, globalFunctions: emptyMap, classDeclsFields: emptyMap,
});
const context = () => ({
    metadata, codegenStateRef: Ref.new({ declarations: [], globalId: 0, reboxPairs: emptySet })(),
    depth: 0, modNameStr: 'Probe', recVars: [], moduleFunctions: emptyMap, bound: emptyMap,
    tcoIdent: Nothing.value, loopCtx: [], options: { isTail: false, inEffectBlock: false },
    mbExpectedExprType: Nothing.value, mbExpectedGoType: Nothing.value,
});
const variable = (module, name) => new TcoExpr(null, new S.Var(new C.Qualified(new Just(module), name)));
const target = variable('Control.Monad.Rec.Class', 'tailRecM');
const dictionaries = [variable('Control.Monad.Rec.Class', 'monadRecEffect'), variable('Control.Monad.ST.Internal', 'monadRecST')];
const args = [dictionaries[0], variable('Probe', 'callback'), variable('Probe', 'seed')];

test('tailRecM requires a saturated call and the exact standard synchronous instance', () => {
    const forbidden = () => { throw new Error('unexpected translation'); };
    const reject = (fn, supplied) => assert.deepEqual(emit(forbidden)(context())(0)(fn)(supplied), Nothing.value);
    reject(variable('Other', 'tailRecM'), args);
    reject(target, args.slice(0, 2));
    reject(target, [...args, args[2]]);
    for (const dictionary of [variable('Other', 'monadRecEffect'), variable('Control.Monad.Rec.Class', 'monadRecFunction'),
        new TcoExpr(null, new S.Local(new Just('dictionary'), 0))]) {
        reject(target, [dictionary, ...args.slice(1)]);
        reject(new TcoExpr(null, new S.Accessor(dictionary, new S.GetProp('tailRecM'))), args.slice(1));
    }
    reject(new TcoExpr(null, new S.Accessor(dictionaries[0], new S.GetProp('other'))), args.slice(1));
});

test('native effect loops preserve strict capture, deferred execution, replay, state, result and exceptions', t => {
    const declarations = [];
    for (const [index, dictionary] of dictionaries.entries()) {
        for (const accessor of [false, true]) {
            const translate = _context => nextId => expression => {
                const name = expression.value1.value0.value1;
                const label = name.startsWith('monadRec') ? 'dictionary' : name;
                const value = label === 'dictionary' ? 'gopurs_runtime.Value{}' : label;
                return { stmts: new StmtLeaf(Go.rawGo(`mark("${label}:statement")`)),
                    expr: Go.rawGo(`capture("${label}:expression", ${value})`), exprType: Go.TypeValue.value, nextId };
            };
            const fn = accessor ? new TcoExpr(null, new S.Accessor(dictionary, new S.GetProp('tailRecM'))) : target;
            const supplied = accessor ? args.slice(1) : [dictionary, ...args.slice(1)];
            const result = emit(translate)(context())(0)(fn)(supplied);
            assert.ok(result instanceof Just);
            const body = [...flattenStmts(result.value0.stmts), new Go.GoReturn(result.value0.expr)].map(printGoExpr).join('\n');
            declarations.push(`func loop${index}_${accessor}(callback, seed gopurs_runtime.Value) gopurs_runtime.Value {\n${body}\n}`);
        }
    }
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-tailrec-effect-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'main.go'), `package main
import ("fmt"; "reflect"; "unsafe"; "gopurs/output/gopurs_runtime")
type Constructor_Control_Monad_Rec_Class_Loop[A,B any] struct { Rc int32; V0 A }
type Constructor_Control_Monad_Rec_Class_Done[A,B any] struct { Rc int32; V0 B }
var trace []string
func mark(s string) { trace=append(trace,s) }
func capture(s string, v gopurs_runtime.Value) gopurs_runtime.Value { mark(s); return v }
func step(v gopurs_runtime.Value, done bool) gopurs_runtime.Value {
    if done { return gopurs_runtime.Value{Type:9,IntVal:${hashString('Data_Control_Monad_Rec_Class_Done')},UnsafePtr:unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value,gopurs_runtime.Value]{1,v})} }
    return gopurs_runtime.Value{Type:9,IntVal:${hashString('Data_Control_Monad_Rec_Class_Loop')},UnsafePtr:unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value,gopurs_runtime.Value]{1,v})}
}
func check(ok bool) { if !ok { panic(fmt.Sprint("bad result: ",trace)) } }
${declarations.join('\n')}
func main() {
    loops:=[]func(gopurs_runtime.Value,gopurs_runtime.Value)gopurs_runtime.Value{loop0_false,loop0_true,loop1_false,loop1_true}
    for _,loop:=range loops {
        trace=nil
        callback:=gopurs_runtime.Func(func(state gopurs_runtime.Value)gopurs_runtime.Value {
            mark(fmt.Sprint("callback:",state.IntVal))
            return gopurs_runtime.Func(func(_ gopurs_runtime.Value)gopurs_runtime.Value {
                mark(fmt.Sprint("effect:",state.IntVal))
                if state.IntVal==2 { return step(gopurs_runtime.Str("done"),true) }
                return step(gopurs_runtime.Int(state.IntVal+1),false)
            })
        })
        action:=loop(callback,gopurs_runtime.Int(0))
        captured:=[]string{"dictionary:statement","dictionary:expression","callback:statement","callback:expression","seed:statement","seed:expression"}
        check(reflect.DeepEqual(trace,captured))
        for repeat:=0;repeat<2;repeat++ {
            trace=nil
            value:=gopurs_runtime.Apply(action,gopurs_runtime.Value{})
            check(value.StrVal()=="done")
            check(reflect.DeepEqual(trace,[]string{"callback:0","effect:0","callback:1","effect:1","callback:2","effect:2"}))
        }
        trace=nil
        immediate:=loop(callback,gopurs_runtime.Int(2));trace=nil
        check(gopurs_runtime.Apply(immediate,gopurs_runtime.Value{}).StrVal()=="done")
        check(reflect.DeepEqual(trace,[]string{"callback:2","effect:2"}))
        sentinel:=&struct{}{}
        failing:=loop(gopurs_runtime.Func(func(_ gopurs_runtime.Value)gopurs_runtime.Value {
            return gopurs_runtime.Func(func(_ gopurs_runtime.Value)gopurs_runtime.Value { panic(sentinel) })
        }),gopurs_runtime.Value{})
        func(){ defer func(){check(recover()==sentinel)}();gopurs_runtime.Apply(failing,gopurs_runtime.Value{});panic("exception lost") }()
    }
    fmt.Println("tailRecM effect semantics passed")
}
`);
    const run = spawnSync('go', ['run', '.'], { cwd: directory, encoding: 'utf8', timeout: 120000,
        env: { ...process.env, GOWORK: 'off' } });
    assert.ifError(run.error);
    assert.equal(run.status, 0, run.stdout + run.stderr);
    assert.match(run.stdout, /tailRecM effect semantics passed/);
});
