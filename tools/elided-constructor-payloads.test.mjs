import { withReboxFields } from './codegen-metadata.mjs';
import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { empty as emptyMap, insert } from '../output/Data.Map/index.js';
import { Just } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet, singleton } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as CodeGen from '../output/Gopurs.CodeGen/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';

const expr = syntax => new NeutralExpr(syntax);
const typed = (type, value) => expr(new S.Typed(type, value));
const local = name => expr(new S.Local(new Just(name), 0));
const lambda = (args, body) => expr(new S.Abs(args.map(name => new Tuple(new Just(name), 0)), body));
const idType = type => new C.ADT('Payload.Id', ['Payload', 'Id'], [type]);
const wrap = (type, value) => typed(idType(type), expr(new S.CtorSaturated(
    new C.Qualified(new Just('Payload'), 'Id'), C.ProductType.value, 'Id', 'Id',
    [new Tuple('value0', value)],
)));

test('elided polymorphic constructors box instantiated fields in workers and closures', t => {
    const metadata = withReboxFields({
        elidedCtors: singleton('Constructor_Payload_Id'),
        ctorTypes: insert(ordString)('Payload.Id')({ vars: ['a'], fields: [new C.TypeVar('a')] })(emptyMap),
        pointerAdtPaths: emptyMap, pointerAdtNodes: emptySet, pointerAdtLeaves: emptyMap,
        enumAdts: emptySet, enumCtors: emptySet, globalTypes: emptyMap,
        globalFunctions: emptyMap, classDeclsFields: emptyMap,
    });
    const add = typed(new C.Func([C.Number.value, C.Number.value], idType(C.Number.value)),
        lambda(['left', 'right'], wrap(C.Number.value, expr(new S.PrimOp(new S.Op2(
            new S.OpNumberNum(S.OpAdd.value), typed(C.Number.value, local('left')), typed(C.Number.value, local('right')),
        ))))));
    const fixtures = [
        ['number', typed(new C.Func([C.Number.value], idType(C.Number.value)),
            lambda(['value'], wrap(C.Number.value, local('value'))))],
        ['string', typed(new C.Func([C.String.value], idType(C.String.value)),
            lambda(['value'], wrap(C.String.value, local('value'))))],
        ['dictionary', expr(new S.Lit(new C.LitRecord([new Tuple('add', add)])))],
    ];
    const code = CodeGen.translate(metadata)({
        name: 'Payload', bindings: fixtures.map(([name, body]) => ({
            recursive: false, bindings: [new Tuple(name, body)],
        })),
        comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
        dataTypes: emptyMap, dataDecls: [], classDecls: [], foreign: emptyMap,
        implementations: emptyMap, directives: emptyMap,
    });
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-elided-payload-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'purescript'));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'purescript/Payload.go'), code);
    writeFileSync(join(directory, 'main.go'), `package main
import (
    "fmt"
    "gopurs/output/gopurs_runtime"
    "gopurs/output/purescript"
)
func main() {
    add := gopurs_runtime.RecordGet(purescript.Get_Payload_dictionary(), "add")
    result := gopurs_runtime.Apply2(add, gopurs_runtime.Float(1.25), gopurs_runtime.Float(2.5))
    fmt.Printf("%.2f %s %.2f", purescript.Call_Payload_number(4.25).FloatVal(),
        purescript.Call_Payload_string("payload").StrVal(), result.FloatVal())
}
`);
    const result = spawnSync('go', ['run', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 30_000,
        env: { ...process.env, GOWORK: 'off' },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stderr);
    assert.equal(result.stdout, '4.25 payload 3.75');
});
