import assert from 'node:assert/strict';
import test from 'node:test';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import * as Map from '../output/Data.Map/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { Nothing, Just } from '../output/Data.Maybe/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { NeutralExpr as E } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import { borrowReadOnlyObjects } from '../output/Gopurs.BorrowedObjects/index.js';

const q = (m,n) => new C.Qualified(m === null ? Nothing.value : new Just(m), n);
const global = (m,n) => new E(new S.Var(q(m,n)));
const local = n => new E(new S.Local(Nothing.value,n));
const app = (f,...args) => new E(new S.App(f,args));
const letIn = (n,value,body) => new E(new S.Let(Nothing.value,n,value,body));
const field = (ctor, source) => new E(new S.Accessor(source,new S.GetCtorField(q('Data.Either',ctor),C.SumType.value,'Either',ctor,'value0',0)));
const lambda = body => new E(new S.Abs([new Tuple(Nothing.value,9)],body));
const reader = (obj=local(2),m='Data.Argonaut.Decode.Decoders',n='getField__1234') => app(global(m,n),global('Probe','custom'),obj,new E(new S.Lit(new C.LitString('x'))));
const identity = app(global('Data.Argonaut.Decode.Class','decodeForeignObject'),global('Data.Argonaut.Decode.Class','decodeJsonJson'));
const method = new E(new S.Accessor(global('Probe','objectDecoder'),new S.GetProp('decodeJson')));
const read = reader();
const helper = 'Data.Argonaut.Decode.Internal.Record.borrowObject';
const helperType = new C.Func([C.Any.value,C.Any.value],C.Any.value);
const metadata = {globalTypes: Map.insert(ordString)(helper)(helperType)(Map.empty)};
const moduleOf = bindings => ({
    name:'Probe', bindings, comments:[], imports:emptySet, exports:emptySet, reExports:emptySet,
    dataTypes:Map.empty, dataDecls:[], classDecls:[], foreign:Map.empty, implementations:Map.empty, directives:Map.empty,
});
function fixture(body, producer=identity, recursive=false) {
  return moduleOf([{recursive,bindings:[new Tuple('objectDecoder',producer),new Tuple('probe',letIn(1,app(method,local(0)),body))]}]);
}
function rewrite(mod, meta=metadata) {
  const before=JSON.stringify(mod), result=borrowReadOnlyObjects(meta)(mod);
  assert.equal(JSON.stringify(mod),before);
  return result;
}
const borrows = (mod, meta=metadata) => JSON.stringify(rewrite(mod,meta)).includes('borrowObject');
test('only synchronous read-only object uses permit borrowing', () => {
  assert.ok(borrows(fixture(letIn(2,field('Right',local(1)),read))));
  assert.ok(borrows(fixture(reader(field('Right',local(1))))));
  assert.ok(borrows(fixture(letIn(2,field('Right',local(1)),letIn(3,local(2),reader(local(3)))))));
  const branches = new E(new S.Branch([
    new S.Pair(new E(new S.PrimOp(new S.Op1(new S.OpIsTag(q('Data.Either','Left')),local(1)))),field('Left',local(1)))
  ],letIn(2,field('Right',local(1)),read)));
  assert.ok(borrows(fixture(branches)));
});
test('returned, captured, stored, mutated and opaque consumers keep owned copies', () => {
  const escapes=[local(2),lambda(read),new E(new S.EffectDefer(read)),
    app(global('Foreign.Object','insert'),local(2)),
    reader(local(2),'Other','getField'),reader(local(2),null,'getField'),
    reader(local(2),'Data.Argonaut.Decode.Decoders','getField__fake'),
    new E(new S.Lit(new C.LitArray([local(2)]))),
    app(global('Probe','unknown'),local(2)),
    app(global('Data.Argonaut.Decode.Decoders','getField'),local(2)),
    new E(new S.LetRec(3,[new Tuple('rec',read)],local(3)))];
  for (const escape of escapes) assert.equal(borrows(fixture(letIn(2,field('Right',local(1)),escape))),false,JSON.stringify(escape));
  assert.equal(borrows(fixture(local(1))),false,'Either envelope escape');
});
test('custom/ambiguous dictionaries, absent helper and recursive groups do not match', () => {
  for(const producer of [global('Probe','unknown'), app(global('Data.Argonaut.Decode.Class','decodeForeignObject'),global('Probe','custom')),app(global(null,'decodeForeignObject'),global('Data.Argonaut.Decode.Class','decodeJsonJson'))]) {
    assert.equal(borrows(fixture(letIn(2,field('Right',local(1)),read),producer)),false);
  }
  const mod=fixture(letIn(2,field('Right',local(1)),read));
  assert.equal(borrows(mod,{globalTypes:Map.empty}),false);
  assert.equal(borrows(fixture(letIn(2,field('Right',local(1)),read),identity,true)),false);
});

const typed = (type, value) => new E(new S.Typed(type, value));
const wrap = value => typed(C.Any.value, new E(new S.TypeApp(typed(C.Any.value, value), C.String.value)));
const payloadBody = letIn(2, field('Right', local(1)), read);
const probe = mod => mod.bindings.flatMap(group => group.bindings).find(pair => pair.value0 === 'probe').value1;
const withValue = (value, body = payloadBody) => moduleOf([{recursive:false, bindings:[
    new Tuple('objectDecoder', identity), new Tuple('probe', letIn(1, value, body)),
]}]);
const borrowed = (fn, json) => app(typed(helperType, global('Data.Argonaut.Decode.Internal.Record', 'borrowObject')), fn, json);

test('qualified nonrecursive dictionary aliases resolve through wrappers without changing their bindings', () => {
    const source = fixture(payloadBody, wrap(global('Probe', 'alias')));
    source.bindings.push({recursive:false, bindings:[
        new Tuple('alias', wrap(global('Probe', 'base'))), new Tuple('base', wrap(identity)),
    ]});
    const result = rewrite(source);
    assert.deepEqual(result.bindings[1], source.bindings[1]);
    assert.deepEqual(result.bindings[0].bindings[0], source.bindings[0].bindings[0]);
    assert.deepEqual(probe(result), letIn(1, borrowed(method, local(0)), payloadBody));
});

test('dictionary cycles, foreign aliases, unqualified aliases and recursive definitions cannot certify identity', () => {
    for (const [producer, aliases, recursive] of [
        [global('Probe', 'objectDecoder'), [], false],
        [global('Probe', 'alias'), [new Tuple('alias', global('Probe', 'objectDecoder'))], false],
        [global('Other', 'alias'), [new Tuple('alias', identity)], false],
        [global(null, 'alias'), [new Tuple('alias', identity)], false],
        [global('Probe', 'alias'), [new Tuple('alias', identity)], true],
    ]) {
        const source = fixture(payloadBody, producer);
        source.bindings.push({recursive, bindings:aliases});
        assert.deepEqual(rewrite(source), source);
    }
});

test('producer admission requires the exact method, builders and unary curried application', () => {
    const otherMethod = new E(new S.Accessor(global('Probe', 'objectDecoder'), new S.GetProp('other')));
    for (const value of [app(otherMethod, local(0)), app(method, local(0), local(3)),
        app(app(method, local(0)), local(3)), new E(new S.UncurriedApp(method, [local(0)]))]) {
        const source = withValue(value);
        assert.deepEqual(rewrite(source), source);
    }
    for (const producer of [app(global('Other', 'decodeForeignObject'), global('Data.Argonaut.Decode.Class', 'decodeJsonJson')),
        app(global('Data.Argonaut.Decode.Class', 'decodeForeignObject'), global('Other', 'decodeJsonJson')),
        app(global('Data.Argonaut.Decode.Class', 'decodeForeignObject'), global('Data.Argonaut.Decode.Class', 'decodeJsonJson'), local(0))]) {
        const source = fixture(payloadBody, producer);
        assert.deepEqual(rewrite(source), source);
    }
});

test('an admitted replacement preserves annotation layers and the original method and JSON expressions', () => {
    const fn = wrap(method);
    // An admitted call does not recursively rewrite its operands. That separate
    // nested candidate must retain its original stage in the replacement.
    const json = letIn(10, app(method, local(0)), reader(field('Right', local(10))));
    const source = withValue(wrap(app(fn, json)));
    const result = rewrite(source);
    assert.equal(result.bindings.length, source.bindings.length);
    assert.equal(result.bindings[0].bindings.length, source.bindings[0].bindings.length);
    assert.deepEqual(probe(result), letIn(1, wrap(borrowed(fn, json)), payloadBody));
});

test('a rejected outer use still rewrites an independently admissible producer operand', () => {
    const nestedBody = reader(field('Right', local(10)));
    const json = letIn(10, app(method, local(0)), nestedBody);
    const source = withValue(app(method, json), local(1));
    assert.deepEqual(probe(rewrite(source)), letIn(1,
        app(method, letIn(10, borrowed(method, local(0)), nestedBody)), local(1)));
});

test('all standard field readers and only numeric specialization suffixes admit the borrowed object slot', () => {
    for (const base of ['getField', 'getFieldOptional', "getFieldOptional'"]) {
        for (const suffix of ['', '__0', '__00123']) {
            assert.ok(borrows(fixture(letIn(2, field('Right', local(1)), reader(local(2), 'Data.Argonaut.Decode.Decoders', base + suffix)))));
        }
        for (const suffix of ['__', '__-1', '__1x', '__1__2']) {
            assert.equal(borrows(fixture(letIn(2, field('Right', local(1)), reader(local(2), 'Data.Argonaut.Decode.Decoders', base + suffix)))), false);
        }
    }
    const fn = global('Data.Argonaut.Decode.Decoders', 'getField');
    for (const value of [app(fn, local(2), local(2), local(0)), app(fn, local(0), local(2), local(2)),
        app(fn, local(0), local(2)), app(fn, local(0), local(2), local(0), local(0))]) {
        assert.equal(borrows(fixture(letIn(2, field('Right', local(1)), value))), false);
    }
});

test('payload recognition requires the exact Either constructor, slot label and index', () => {
    assert.ok(borrows(fixture(field('Left', local(1)))));
    for (const [owner, ctor, label, index] of [
        ['Other', 'Right', 'value0', 0], ['Data.Either', 'Right', 'value1', 0],
        ['Data.Either', 'Right', 'value0', 1], ['Data.Either', 'Left', 'value1', 0],
    ]) {
        const payload = new E(new S.Accessor(local(1), new S.GetCtorField(q(owner, ctor), C.SumType.value, 'Either', ctor, label, index)));
        assert.equal(borrows(fixture(reader(payload))), false);
    }
});

test('alias tracking uses levels and removes a rebound object alias, while result shadowing stays conservative', () => {
    const renamed = new E(new S.Local(new Just('anotherName'), 2));
    assert.ok(borrows(fixture(letIn(2, field('Right', local(1)), reader(renamed)))));
    const value = new E(new S.Lit(new C.LitInt(7)));
    assert.ok(borrows(fixture(letIn(2, field('Right', local(1)), letIn(2, value, local(2))))));
    assert.equal(borrows(fixture(letIn(1, value, local(1)))), false);
    assert.equal(borrows(fixture(letIn(2, field('Right', local(1)),
        new E(new S.Abs([new Tuple(Nothing.value, 2)], local(2)))))), false);
});

for (const [name, scope] of [
    ['curried closure', value => lambda(value)],
    ['uncurried closure', value => new E(new S.UncurriedAbs([new Tuple(Nothing.value, 9)], value))],
    ['effect closure', value => new E(new S.UncurriedEffectAbs([new Tuple(Nothing.value, 9)], value))],
    ['recursive scope', value => new E(new S.LetRec(3, [new Tuple('rec', value)], local(3)))],
    ['effect bind', value => new E(new S.EffectBind(Nothing.value, 9, value, local(9)))],
    ['deferred effect', value => new E(new S.EffectDefer(value))],
    ['pure effect', value => new E(new S.EffectPure(value))],
    ['primitive effect', value => new E(new S.PrimEffect(new S.EffectRefNew(value)))],
    ['uncurried effect call', value => new E(new S.UncurriedEffectApp(global('Probe', 'effect'), [value]))],
]) {
    test(`${name} requires complete independence from the borrowed result and its aliases`, () => {
        assert.equal(borrows(fixture(letIn(2, field('Right', local(1)), scope(reader())))), false);
        assert.ok(borrows(fixture(letIn(2, field('Right', local(1)), scope(reader(local(0)))))));
    });
}

test('local recursive regions remain unchanged even when they contain an otherwise admissible decode', () => {
    const nested = letIn(1, app(method, local(0)), payloadBody);
    const recursive = new E(new S.LetRec(5, [new Tuple('rec', nested)], local(5)));
    const source = withValue(local(0), recursive);
    assert.deepEqual(rewrite(source), source);
});

test('native identity borrowing and owned-copy fallback preserve containers and exact failure dispatch', t => {
    const work = mkdtempSync(join(tmpdir(), 'gopurs-borrowed-object-native-'));
    t.after(() => rmSync(work, {recursive:true, force:true}));
    mkdirSync(join(work, 'output/gopurs_runtime'), {recursive:true});
    mkdirSync(join(work, 'record'));
    writeFileSync(join(work, 'go.mod'), 'module gopurs\n\ngo 1.22\n');
    writeFileSync(join(work, 'output/gopurs_runtime/runtime.go'), readFileSync(new URL('../runtime/runtime.go', import.meta.url)));
    for (const [name, source] of [
        ['record.go', '../../gopurs-argonaut-codecs/src/Data/Argonaut/Decode/Internal/Record.go'],
        ['parser.go', '../../gopurs-argonaut-core/src/Data/Argonaut/Parser.go'],
        ['text.go', '../../gopurs-argonaut-codecs/src/Data/Argonaut/Decode/Parser.go'],
        ['record_test.go', '../../gopurs-argonaut-codecs/test/record-plan_test.go'],
    ]) writeFileSync(join(work, 'record', name), readFileSync(new URL(source, import.meta.url), 'utf8').replace('package Parser', 'package Record'));
    const checked = spawnSync('go', ['test', '-v', '-race', '-count=1', '-run', '^Test(BorrowObjectIdentityAndFallback|IdentityObjectShortcutOwnsContainer)$', './record'], {
        cwd:work, encoding:'utf8', timeout:60_000, env:{...process.env, GOWORK:'off', GOMAXPROCS:'2'},
    });
    assert.ifError(checked.error);
    assert.equal(checked.status, 0, checked.stdout + checked.stderr);
    assert.match(checked.stdout, /--- PASS: TestBorrowObjectIdentityAndFallback/);
    assert.match(checked.stdout, /--- PASS: TestIdentityObjectShortcutOwnsContainer/);
});
