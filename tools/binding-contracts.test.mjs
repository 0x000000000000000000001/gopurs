import assert from 'node:assert/strict';
import test from 'node:test';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { empty as emptyMap, lookup } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import * as Bindings from '../output/Gopurs.BindingExprs/index.js';
import * as CodeGen from '../output/Gopurs.CodeGen/index.js';
import { StmtEmpty, flattenStmts } from '../output/Gopurs.ExprContext/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime/index.js';
import { TcoExpr } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { localId } from '../output/PureScript.Backend.Optimizer.FreeVars/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { withReboxFields } from './codegen-metadata.mjs';

const metadata = withReboxFields({
    elidedCtors: emptySet, ctorTypes: emptyMap, pointerAdtPaths: emptyMap,
    pointerAdtNodes: emptySet, pointerAdtLeaves: emptyMap, enumAdts: emptySet,
    enumCtors: emptySet, globalTypes: emptyMap, globalFunctions: emptyMap, classDeclsFields: emptyMap,
});
const get = (map, key) => lookup(ordString)(key)(map).value0;

test('recursive workers publish peer signatures first, refine results in order, and declare before assigning', () => {
    const context = {
        metadata, codegenStateRef: Ref.new({ declarations: [], globalId: 11, reboxPairs: emptySet })(),
        depth: 0, modNameStr: 'Bindings', recVars: [], moduleFunctions: emptyMap, bound: emptyMap,
        tcoIdent: Nothing.value, loopCtx: [], options: { isTail: false, inEffectBlock: false },
        mbExpectedExprType: Nothing.value,
    };
    const node = syntax => new TcoExpr(null, syntax);
    const firstBody = node(new S.Lit(new C.LitInt(1)));
    const secondBody = node(new S.Lit(new C.LitInt(2)));
    const outerBody = node(new S.Lit(new C.LitInt(3)));
    const abstraction = (param, body) => node(new S.Typed(new C.Func([C.Int.value], C.Any.value),
        node(new S.Abs([new Tuple(new Just(param), 2)], body))));
    const first = localId(new Just('first'))(1), second = localId(new Just('second'))(1);
    const calls = [];
    const translate = child => nextId => value => {
        const a = get(child.bound, first), b = get(child.bound, second);
        assert.ok(a && b, 'the whole recursive group must be in scope');
        const aInfo = get(child.moduleFunctions, a.name), bInfo = get(child.moduleFunctions, b.name);
        assert.deepEqual(aInfo.fArgs, [Go.TypeInt64.value]);
        assert.deepEqual(bInfo.fArgs, [Go.TypeInt64.value]);
        calls.push(value);
        if (value === firstBody) {
            assert.deepEqual(aInfo.fRet, Go.TypeValue.value);
            assert.deepEqual(bInfo.fRet, Go.TypeValue.value);
        } else if (value === secondBody) {
            assert.deepEqual(aInfo.fRet, Go.TypeInt64.value);
            assert.deepEqual(bInfo.fRet, Go.TypeValue.value);
            // Peer bindings remain the predeclared view while bodies are translated.
            assert.deepEqual(a.goType, new Go.TypeFunc([Go.TypeInt64.value], Go.TypeValue.value));
        } else {
            assert.equal(value, outerBody);
            assert.deepEqual(aInfo.fRet, Go.TypeInt64.value);
            assert.deepEqual(bInfo.fRet, Go.TypeInt64.value);
            assert.deepEqual(a.goType, new Go.TypeFunc([Go.TypeInt64.value], Go.TypeInt64.value));
            assert.deepEqual(b.goType, new Go.TypeFunc([Go.TypeInt64.value], Go.TypeInt64.value));
        }
        if (value !== outerBody) {
            assert.equal(child.loopCtx.length, 1);
            assert.equal(child.loopCtx[0].ident, value === firstBody ? a.name : b.name);
            assert.equal(child.options.isTail, true);
        }
        return { stmts: StmtEmpty.value, expr: new Go.GoInt(42), exprType: Go.TypeInt64.value, nextId: nextId + 1 };
    };
    const result = Bindings.recursive(translate)(context)(5)(new TcoExpr({ role: { isLoop: true } }, null))(1)
        ([new Tuple('first', abstraction('x', firstBody)), new Tuple('second', abstraction('y', secondBody))])(outerBody);
    assert.deepEqual(calls, [firstBody, secondBody, outerBody]);
    assert.equal(result.nextId, 10);
    assert.equal(Ref.read(context.codegenStateRef)().globalId, 13);
    const statements = flattenStmts(result.stmts);
    assert.equal(statements.length, 12);
    for (const statement of statements.slice(0, 8)) assert.ok(!(statement instanceof Go.GoMutate));
    for (const statement of statements.slice(8)) assert.ok(statement instanceof Go.GoMutate);
    assert.match(printGoExpr(statements[0]), new RegExp(`^var Call_local_Bindings_${first}_5_11 func\\(int64\\) int64$`));
    assert.match(printGoExpr(statements[4]), new RegExp(`^var Call_local_Bindings_${second}_6_12 func\\(int64\\) int64$`));
});

test('tail loops keep simultaneous argument values and per-iteration closure captures', t => {
    const expr = syntax => new NeutralExpr(syntax);
    const typed = (type, value) => expr(new S.Typed(type, value));
    const local = (name, level) => typed(C.Int.value, expr(new S.Local(new Just(name), level)));
    const integer = value => typed(C.Int.value, expr(new S.Lit(new C.LitInt(value))));
    const app = (fn, args) => expr(new S.App(fn, args));
    const variable = (module, name) => expr(new S.Var(new C.Qualified(new Just(module), name)));
    const binary = (type, op, a, b) => typed(type, expr(new S.PrimOp(new S.Op2(op, a, b))));
    const n = local('n', 0), a = local('a', 1), b = local('b', 2);
    const decrement = binary(C.Int.value, new S.OpIntNum(S.OpSubtract.value), n, integer(1));
    const callback = expr(new S.Abs([new Tuple(new Just('ignored'), 4)], a));
    const save = app(variable('Probe', 'save'), [callback]);
    const step = expr(new S.Let(new Just('saved'), 3, save,
        app(variable('Bindings', 'loop'), [decrement, b, a])));
    const body = expr(new S.Branch([new S.Pair(
        binary(C.Boolean.value, new S.OpIntOrd(S.OpEq.value), n, integer(0)), a)], step));
    const fn = typed(new C.Func([C.Int.value, C.Int.value, C.Int.value], C.Int.value),
        expr(new S.Abs(['n', 'a', 'b'].map((name, level) => new Tuple(new Just(name), level)), body)));
    const code = CodeGen.translate(metadata)({
        name: 'Bindings', bindings: [{ recursive: true, bindings: [new Tuple('loop', fn)] }],
        comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
        dataTypes: emptyMap, dataDecls: [], classDecls: [], foreign: emptyMap,
        implementations: emptyMap, directives: emptyMap,
    });
    assert.match(code, /continue loop/);
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-binding-contracts-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'purescript'));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'purescript/Bindings.go'), code);
    writeFileSync(join(directory, 'purescript/Probe.go'), `package purescript
import "gopurs/output/gopurs_runtime"
var Saved []gopurs_runtime.Value
func Get_Probe_save() gopurs_runtime.Value {
    return gopurs_runtime.Func(func(callback gopurs_runtime.Value) gopurs_runtime.Value {
        Saved = append(Saved, callback)
        return gopurs_runtime.Value{}
    })
}
`);
    writeFileSync(join(directory, 'main.go'), `package main
import (
    "fmt"
    "gopurs/output/gopurs_runtime"
    "gopurs/output/purescript"
)
func main() {
    fmt.Println(purescript.Call_Bindings_loop(3, 10, 20))
    for _, callback := range purescript.Saved {
        fmt.Println(gopurs_runtime.Apply(callback, gopurs_runtime.Value{}).IntVal)
    }
}
`);
    const result = spawnSync('go', ['run', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 30_000, env: { ...process.env, GOWORK: 'off' },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stdout + result.stderr);
    assert.equal(result.stdout, '20\n10\n20\n10\n');
});
