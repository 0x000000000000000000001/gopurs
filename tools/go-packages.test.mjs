import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { dirname, join } from 'node:path';
import test from 'node:test';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { printGoFile } from '../output/Gopurs.Printer/index.js';
import { generated } from '../output/Gopurs.GoSymbol/index.js';
import { writeEntryPoints, writePackageOutput } from '../output/Gopurs.Driver.Output/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { Left } from '../output/Data.Either/index.js';
import * as Aff from '../output/Effect.Aff/index.js';

const runAff = action => new Promise((resolve, reject) => {
    Aff.runAff(result => () => result instanceof Left ? reject(result.value0) : resolve(result.value0))(action)();
});
const put = (root, name, text) => {
    const path = join(root, name);
    mkdirSync(dirname(path), { recursive: true });
    writeFileSync(path, text);
};
const run = (root, label, args, expected = 0) => {
    const result = spawnSync('go', args, { cwd: root, encoding: 'utf8', timeout: 120_000,
        maxBuffer: 64 * 1024 * 1024, env: { ...process.env, GOWORK: 'off', GOMAXPROCS: '2' } });
    put(root, label + '.log', result.stdout + result.stderr);
    put(root, label + '.command.json', JSON.stringify(args));
    assert.ifError(result.error);
    assert.equal(result.status, expected, result.stderr);
    return result.stdout;
};
const declaration = (module, name, params, result, body) => new Go.GoFunctionDecl({
    name: generated(module)(name), params, result, body,
});
const source = (module, imports, declarations) => ({
    moduleName: module, owner: 'generated/' + module, path: module + '.go', initializers: [], records: [],
    content: printGoFile({ packageName: 'purescript', imports, declarationGroups: [declarations] }),
});

test('the production AST and output writer preserve annotated records, initializers and selected entries', async t => {
    const root = mkdtempSync(join(process.env.GOPURS_TEST_EVIDENCE ?? tmpdir(), 'go-packages-'));
    if (process.env.GOPURS_TEST_EVIDENCE) console.log('Package writer evidence: ' + root);
    else t.after(() => rmSync(root, { recursive: true, force: true }));
    const fields = [new Tuple('type', Go.TypeString.value), new Tuple('value', Go.TypeInt64.value),
        new Tuple('\ud800', Go.TypeInt64.value), new Tuple('💡', Go.TypeBool.value),
        ...['go__type', 'a-b', 'a_minus_b', '', 'X_empty', '1', 'X_1'].map(label => new Tuple(label, Go.TypeInt64.value))];
    const record = new Go.TypeRecord(fields);
    const provider = source('Provider', [], [declaration('Provider', 'Make', [], record,
        new Go.GoReturn(new Go.GoRecordDict(record, [new Tuple('type', new Go.GoString('tag')),
            new Tuple('value', new Go.GoInt(42)), new Tuple('\ud800', new Go.GoInt(7)), new Tuple('💡', Go.rawGo('true')),
            ...['go__type', 'a-b', 'a_minus_b', '', 'X_empty', '1', 'X_1'].map(label => new Tuple(label, new Go.GoInt(1)))])))]);
    const consumer = source('Consumer', [], [declaration('Consumer', 'Read', [new Tuple('r', record)], Go.TypeInt64.value,
        new Go.GoReturn(new Go.GoRecordFieldAccess(new Go.GoVar('r'), 'value')))]);
    const main = source('Main', ['fmt', 'gopurs/output/gopurs_runtime'], [
        new Go.GoInitDecl(generated('Main')('probe/registration'), Go.rawGo('fmt.Print("init;")')),
        declaration('Main', 'Get_Main_main', [], Go.TypeValue.value, Go.rawGo(`return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
  r := Make()
  fmt.Printf("%d/%s", Read(r), r.go__type)
  return gopurs_runtime.Value{}
})`)),
        // The two fixed identities open boundaries at the delivered 1/8192
        // density; 512 arbitrary declarations alone need not contain an anchor.
        ...[...Array.from({ length: 512 }, (_, i) => i), 20110, 40217].map(i => declaration('Main', 'Leaf_' + i, [], Go.TypeInt64.value,
            new Go.GoReturn(new Go.GoInt(i)))),
    ]);
    const files = [provider, consumer, main];
    const before = join(root, 'before'), after = join(root, 'after');
    for (const directory of [before, after]) {
        put(directory, 'go.mod', 'module gopurs/output\n\ngo 1.22\n');
        put(directory, 'gopurs_runtime/runtime.go', runtimeGoCode);
    }
    for (const file of files) put(before, 'purescript/' + file.path, file.content);
    await runAff(writeEntryPoints(before)(['Main']));
    const expected = run(before, 'execute', ['run', './main']);
    assert.equal(expected, 'init;42/tag');
    const exported = run(before, 'exports', ['list', '-deps', '-export', '-f', '{{.ImportPath}}\t{{.Export}}', './purescript']);
    const exports = {};
    for (const line of exported.trimEnd().split('\n')) {
        const [name, file] = line.split('\t');
        if (file && name !== 'gopurs/output/purescript') exports[name] = readFileSync(file).toString('base64');
    }
    const request = { packagePath: 'gopurs/output/purescript', goVersion: 'go1.22', arch: 'arm64', files, exports,
        entries: [generated('Main')('Get_Main_main')] };
    put(root, 'request.json', JSON.stringify(request));
    const result = await runAff(writePackageOutput(after)(['Main'])(request));
    put(root, 'plan.json', JSON.stringify(result, null, 2));
    assert.ok(result.packages.length > 1);
    assert.equal(run(after, 'execute', ['run', './main']), expected);
    assert.equal(run(after, 'selected-entry', ['run', './Main/main']), expected);
    run(after, 'dag', ['list', '-deps', '-json', './main']);
    for (const file of result.files) {
        assert.equal(readFileSync(join(after, 'purescript', file.path), 'utf8'), file.content);
        assert.ok(!file.content.includes('/*gopurs:record-'));
        assert.ok(!file.content.includes('//gopurs:init-v1:'));
    }
    assert.equal(readFileSync(join(after, 'gopurs_runtime/runtime.go'), 'utf8'), runtimeGoCode);
});
