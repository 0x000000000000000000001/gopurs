import assert from 'node:assert/strict';
import test from 'node:test';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { empty as emptyMap, insert } from '../output/Data.Map/index.js';
import { Just } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet, singleton } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import * as Ownership from '../output/Gopurs.Ownership/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import * as CodeGen from '../output/Gopurs.CodeGen/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime/index.js';

// Deliberately unrelated names/layout: acceptance must depend on structure.
const moduleName = 'OwnershipFixture';
const tree = new C.ADT(`${moduleName}.Shape`, [moduleName, 'Shape'], []);
const fields = [C.Int.value, tree, tree];
const map = entries => entries.reduce((result, [key, value]) =>
    insert(ordString)(key)(value)(result), emptyMap);
const metadata = {
    elidedCtors: emptySet,
    ctorTypes: map([[`${moduleName}.Cell`, { vars: [], fields }],
        [`${moduleName}.Vacant`, { vars: [], fields: [] }]]),
    pointerAdtPaths: map([[`${moduleName}.Shape`, { ctorName: 'Cell', arity: 0 }]]),
    pointerAdtNodes: singleton(`Data_${moduleName}_Cell`),
    pointerAdtLeaves: map([[`Data_${moduleName}_Vacant`, {
        nodeBaseStruct: `Data_${moduleName}_Cell`, nodeCtor: 'Cell',
    }]]),
    enumAdts: emptySet, enumCtors: emptySet, globalTypes: emptyMap,
    globalFunctions: emptyMap, classDeclsFields: emptyMap,
};
const expr = syntax => new NeutralExpr(syntax);
const typed = (type, value) => expr(new S.Typed(type, value));
const local = (level = 0) => typed(tree, expr(new S.Local(new Just('input'), level)));
const literal = value => typed(C.Int.value, expr(new S.Lit(new C.LitInt(value))));
const qualified = name => new C.Qualified(new Just(moduleName), name);
const constructor = (name, values) => typed(tree, expr(new S.CtorSaturated(
    qualified(name), C.SumType.value, 'Shape', name,
    values.map((value, index) => new Tuple(`value${index}`, value)),
)));
const empty = () => constructor('Vacant', []);
const cell = (key, left, right) => constructor('Cell', [key, left, right]);
const field = (base, index) => typed(fields[index], expr(new S.Accessor(base,
    new S.GetCtorField(qualified('Cell'), C.SumType.value, 'Shape', 'Cell', `value${index}`, index))));
const branch = body => typed(tree, expr(new S.Branch([
    new S.Pair(typed(C.Boolean.value, expr(new S.PrimOp(new S.Op1(
        new S.OpIsTag(qualified('Cell')), local())))), body),
], empty())));
const functionType = new C.Func([tree], tree);
const call = (name, argument) => typed(tree, expr(new S.App(
    typed(functionType, expr(new S.Var(qualified(name)))), [argument])));
const lambda = body => typed(functionType,
    expr(new S.Abs([new Tuple(new Just('input'), 0)], body)));
const moduleOf = bindings => ({
    name: moduleName,
    bindings: bindings.map(([name, value]) => ({ recursive: false, bindings: [new Tuple(name, value)] })),
    comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
    dataTypes: emptyMap,
    dataDecls: [{ name: 'Shape', vars: [], constructors: [
        { name: 'Vacant', fields: [] }, { name: 'Cell', fields },
    ] }],
    classDecls: [], foreign: emptyMap, implementations: emptyMap, directives: emptyMap,
});
const prepare = bindings => Ownership.prepare(metadata)(moduleOf(bindings));
const resultBindings = result => result.module.bindings.flatMap(group => group.bindings);
const binding = (result, name) => resultBindings(result).find(pair => pair.value0 === name).value1;
const change = () => lambda(branch(cell(literal(9), field(local(), 1), field(local(), 2))));
const freshTree = () => cell(literal(1), cell(literal(2), empty(), empty()), empty());
const freshCaller = (body = call('change', freshTree())) =>
    typed(new C.Func([C.Int.value], tree), expr(new S.Abs([
        new Tuple(new Just('seed'), 0),
    ], body)));
const letTree = (level, value, body, name = 'input') =>
    typed(tree, expr(new S.Let(new Just(name), level, value, body)));
const equalInt = (left, right) => typed(C.Boolean.value, expr(new S.PrimOp(
    new S.Op2(new S.OpIntOrd(S.OpEq.value), left, right))));
const chooseTree = (condition, yes, no) => typed(tree,
    expr(new S.Branch([new S.Pair(condition, yes)], no)));
const assertCallerUnchanged = (caller, name) => {
    const positiveControl = freshCaller();
    const result = prepare([
        ['change', change()], ['freshPositiveControl', positiveControl], [name, caller],
    ]);
    assert.ok(result.declarations.length > 0,
        'the independent fresh caller must enable ownership optimization');
    assert.notDeepEqual(binding(result, 'freshPositiveControl'), positiveControl,
        'the positive control must actually select a consuming worker');
    assert.deepEqual(binding(result, name), caller,
        'this caller must retain its persistent calls despite available workers');
};

test('owned tree workers accept a different constructor name and scalar field position', () => {
    const caller = freshCaller();
    const result = prepare([
        ['change', change()], ['fresh', caller],
    ]);
    assert.ok(result.declarations.length > 0, 'a proven worker must be emitted');
    assert.notDeepEqual(binding(result, 'fresh'), caller, 'the owned call must select the worker');
});

for (const [name, body] of [
    ['duplicate child', () => cell(literal(9), field(local(), 1), field(local(), 1))],
    ['ancestor and descendant', () => cell(literal(9), local(), field(local(), 1))],
    ['unknown call', () => call('unknownForeign', local())],
]) {
    test(`owned tree proof rejects ${name}`, () => {
        const caller = freshCaller();
        const result = prepare([
            ['change', lambda(branch(body()))], ['fresh', caller],
        ]);
        assert.equal(result.declarations.length, 0, 'unsupported sharing must not emit a destructive worker');
        assert.deepEqual(binding(result, 'fresh'), caller,
            'fresh ownership at the call site cannot excuse sharing within a callee');
    });
}

test('a borrowed parameter cannot select a consuming worker', () => {
    const caller = lambda(call('change', local()));
    assertCallerUnchanged(caller, 'borrowed');
});

test('a fresh parent containing a borrowed subtree is not an exclusive tree', () => {
    const caller = lambda(call('change', cell(literal(1), local(), empty())));
    assertCallerUnchanged(caller, 'borrowedChild');
});

test('equal names do not turn an outer borrowed level into an inner owned level', () => {
    // Both locals are called input, but only level 1 is allocated here.
    const caller = lambda(letTree(1, freshTree(),
        cell(literal(3), call('change', local(0)), local(1))));
    assertCallerUnchanged(caller, 'shadowedNames');
});

test('a consuming let cannot overwrite an old root used by the continuation', () => {
    const caller = freshCaller(letTree(1, freshTree(),
        letTree(2, call('change', local(1)),
            cell(literal(3), local(1), local(2)))));
    assertCallerUnchanged(caller, 'retainedRoot');
});

test('an old scalar read in a later branch guard keeps the root observable', () => {
    const caller = freshCaller(letTree(1, freshTree(),
        letTree(2, call('change', local(1)),
            chooseTree(equalInt(field(local(1), 0), literal(1)), local(2), empty()))));
    assertCallerUnchanged(caller, 'guardReadsOldScalar');
});

test('a guard call cannot consume a tree still reachable on its failure path', () => {
    const caller = freshCaller(letTree(1, freshTree(),
        chooseTree(equalInt(field(call('change', local(1)), 0), literal(9)),
            empty(), local(1))));
    assertCallerUnchanged(caller, 'guardFailureKeepsRoot');
});

test('a consumed child remains borrowed while its old parent exposes that path', () => {
    const caller = freshCaller(letTree(1, freshTree(),
        letTree(2, call('change', field(local(1), 1)),
            cell(literal(3), field(local(1), 1), local(2)))));
    assertCallerUnchanged(caller, 'retainedChildPath');
});

test('a newly allocated parent can still duplicate a locally allocated child', () => {
    const caller = freshCaller(letTree(1, freshTree(),
        call('change', cell(literal(3), local(1), local(1)))));
    assertCallerUnchanged(caller, 'aliasedFreshChild');
});

for (const [name, body] of [
    ['root kept after its call', () => letTree(1, call('alter', local()),
        cell(literal(3), local(), local(1)))],
    ['scalar read in a later guard', () => letTree(1, call('alter', local()),
        chooseTree(equalInt(field(local(), 0), literal(1)), local(1), empty()))],
    ['child retained through its parent', () => letTree(1, call('alter', field(local(), 1)),
        cell(literal(3), field(local(), 1), local(1)))],
    ['child retained through a let alias', () => letTree(1, field(local(), 1),
        letTree(2, call('alter', local(1)), cell(literal(3), field(local(), 1), local(2))))],
]) {
    test(`an owned callee rejects a consuming let with ${name}`, () => {
        // No local-let entry shortcut can make this rejection pass trivially:
        // the caller passes a constructor tree directly to the candidate.
        const caller = freshCaller();
        const helperCaller = freshCaller(call('alter', freshTree()));
        const result = prepare([
            ['alter', change()], ['change', lambda(branch(body()))],
            ['fresh', caller], ['helperPositiveControl', helperCaller],
        ]);
        assert.notDeepEqual(binding(result, 'helperPositiveControl'), helperCaller,
            'the mutating dependency must have a usable worker');
        assert.deepEqual(binding(result, 'fresh'), caller,
            'old values observed by the callee continuation prohibit destructive specialization');
    });
}

const referencedLocals = root => {
    if (root === null || typeof root !== 'object') return [];
    if (root instanceof S.Var && root.value0.value0 instanceof Just
        && root.value0.value0.value0 === moduleName) return [root.value0.value1];
    return Object.values(root).flatMap(referencedLocals);
};
const declarationNames = result => result.declarations
    .filter(declaration => declaration instanceof Go.GoFunctionDecl)
    .map(declaration => declaration.value0.name);
const assertUniqueDeclarations = result => {
    const names = declarationNames(result);
    assert.ok(names.length > 0, 'workers must exist for this collision fixture');
    assert.equal(new Set(names).size, names.length, 'every emitted Go function needs a distinct name');
};

test('a user binding with the discovered worker name is preserved and gets a fresh worker neighbour', () => {
    const first = prepare([['change', change()], ['fresh', freshCaller()]]);
    const generatedName = referencedLocals(binding(first, 'fresh')).find(name => name !== 'change');
    assert.ok(generatedName, 'the fresh caller must identify a generated worker');
    const userFunction = lambda(empty());
    const caller = freshCaller();
    const result = prepare([
        ['change', change()], [generatedName, userFunction], ['fresh', caller],
    ]);
    assert.deepEqual(binding(result, generatedName), userFunction, 'the user binding must be preserved');
    assert.notDeepEqual(binding(result, 'fresh'), caller, 'a collision must be resolved with a fresh identity');
    const selected = referencedLocals(binding(result, 'fresh'));
    assert.ok(selected.length > 0);
    assert.ok(selected.every(name => Go.sanitizeName(name) !== Go.sanitizeName(generatedName)),
        'the rewritten call must not accidentally select the user binding');
    assertUniqueDeclarations(result);
});

test('worker and consume helper names cannot collide between separate source functions', () => {
    const result = prepare([
        ['change', change()], ['change_consume', change()],
        ['fresh', freshCaller()], ['freshOther', freshCaller(call('change_consume', freshTree()))],
    ]);
    assertUniqueDeclarations(result);
});

test('a distinct user name that sanitizes like a worker forces a fresh worker name', () => {
    const caller = freshCaller(call('change$', freshTree()));
    const first = prepare([['change$', change()], ['fresh', caller]]);
    const generatedName = referencedLocals(binding(first, 'fresh')).find(name => name !== 'change$');
    assert.ok(generatedName);
    const userName = generatedName.includes('$')
        ? generatedName.replaceAll('$', '_dollar_')
        : generatedName.replace('_dollar_', '$');
    assert.notEqual(userName, generatedName, 'the source spellings must be distinct');
    assert.equal(Go.sanitizeName(userName), Go.sanitizeName(generatedName),
        'the adversarial name must collide in the emitted Go namespace');
    const userFunction = lambda(empty());
    const result = prepare([
        ['change$', change()], [userName, userFunction], ['fresh', caller],
    ]);
    assert.deepEqual(binding(result, userName), userFunction);
    assert.notDeepEqual(binding(result, 'fresh'), caller);
    assert.ok(referencedLocals(binding(result, 'fresh'))
        .every(name => Go.sanitizeName(name) !== Go.sanitizeName(userName)));
    assertUniqueDeclarations(result);
});

test('foreign bindings also reserve the generated worker namespace', () => {
    const caller = freshCaller();
    const first = prepare([['change', change()], ['fresh', caller]]);
    const generatedName = referencedLocals(binding(first, 'fresh')).find(name => name !== 'change');
    assert.ok(generatedName);
    const source = moduleOf([['change', change()], ['fresh', caller]]);
    const result = Ownership.prepare(metadata)({
        ...source, foreign: map([[generatedName, new Just(functionType)]]),
    });
    assert.notDeepEqual(binding(result, 'fresh'), caller);
    assert.ok(referencedLocals(binding(result, 'fresh'))
        .every(name => Go.sanitizeName(name) !== Go.sanitizeName(generatedName)));
    assertUniqueDeclarations(result);
});

test('a dead root donated through a computed let cannot be recycled again as its own parent', t => {
    const caller = freshCaller();
    const source = moduleOf([
        // The creator ignores its tree argument, so the caller may supply its
        // otherwise-dead root as the allocation donor for this new leaf.
        ['create', lambda(cell(literal(1), empty(), empty()))],
        ['change', lambda(letTree(1, call('create', empty()),
            cell(literal(0), local(1), empty())))],
        ['fresh', caller],
    ]);
    const prepared = Ownership.prepare(metadata)(source);
    assert.notDeepEqual(binding(prepared, 'fresh'), caller,
        'this positive donation case must actually use the consuming implementation');
    const code = CodeGen.translate(metadata)(source);
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-owned-donor-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'purescript'));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'purescript/OwnershipFixture.go'), code);
    writeFileSync(join(directory, 'main.go'), `package main
import (
    "fmt"
    "gopurs/output/purescript"
)
type tree = purescript.Constructor_OwnershipFixture_Cell
func render(node *tree, seen map[*tree]bool) string {
    if node == nil { return "E" }
    if seen[node] { panic("a donor cell was used twice: aliased or cyclic result") }
    seen[node] = true
    return fmt.Sprintf("%d(%s)(%s)", node.V0, render(node.V1, seen), render(node.V2, seen))
}
func main() {
    for i := int64(0); i < 10; i++ {
        result := purescript.Call_OwnershipFixture_fresh(i)
        got := render(result, map[*tree]bool{})
        if got != "0(1(E)(E))(E)" { panic(got) }
    }
    fmt.Println("acyclic owned result")
}
`);
    const result = spawnSync('go', ['run', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 30_000,
        env: { ...process.env, GOWORK: 'off', GOCACHE: join(tmpdir(), 'gopurs-owned-tests-go-cache') },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stderr);
    assert.equal(result.stdout, 'acyclic owned result\n');
});
