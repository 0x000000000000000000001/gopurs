import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { opaqueCode, referencedImports } from '../output/Gopurs.GoCode/index.js';
import { collectImports } from '../output/Gopurs.GoImports/index.js';
import { GoFunctionDecl, GoRaw, GoReturn, TypeFloat64, TypeString } from '../output/Gopurs.GoAst/index.js';
import { printGoFile } from '../output/Gopurs.Printer/index.js';

test('opaque Go imports ignore string, rune, raw-string and comment contents', () => {
    for (const source of [
        '"math.Mod sync.Once unsafe.Pointer gopurs_runtime.Value"',
        String.raw`"escaped \" math.Mod"`,
        '`math.Mod\nsync.Once`',
        "'m' // math.Mod\ns := \"unsafe.Pointer\"",
        String.raw`'\'' /* math.Mod sync.Once */`,
        '/* gopurs_runtime.Value */\n"math.Mod"',
        'custommath.Mod math_extra.Value',
    ]) assert.deepEqual(referencedImports(source), [], source);
    assert.deepEqual(referencedImports('"math.Mod" + math.Mod(7, 2)'), ['math']);
    assert.deepEqual(referencedImports('/* math.Mod */ unsafe.Pointer(nil)'), ['unsafe']);
    assert.deepEqual(referencedImports('// math.Mod\nsync.Once{}'), ['sync']);
    assert.deepEqual(referencedImports('math.Mod(1, 2) + math.Mod(3, 4)'), ['math']);
});

test('opaque Go imports retain UTF-16 boundaries through Unicode and long fragments', () => {
    for (const prefix of ['ÿ', '≠', '·', '😀', '\ud800', '\udfff']) {
        assert.deepEqual(referencedImports(`${prefix}math.Mod(1, 2)`), [], prefix);
        assert.deepEqual(referencedImports(`"${prefix}"; math.Mod(1, 2)`), ['math'], prefix);
        assert.deepEqual(referencedImports(`/* ${prefix} unsafe.Pointer */ sync.Once{}`), ['sync'], prefix);
    }
    const fragment = `${'/* 😀 math.Mod */ \"ÿ unsafe.Pointer\"; '.repeat(512)}gopurs_runtime.Value{}; math.Mod(1, 2); sync.Once{}; unsafe.Pointer(nil)`;
    assert.deepEqual(referencedImports(fragment), ['gopurs/output/gopurs_runtime', 'math', 'sync', 'unsafe']);
});

for (const [name, source, result, expectedImports] of [
    ['literal', '"math.Mod sync.Once unsafe.Pointer gopurs_runtime.Value"', TypeString.value, []],
    ['call', 'math.Mod(7, 2)', TypeFloat64.value, ['math']],
]) test(`opaque ${name} produces Go with exactly the imports it uses`, t => {
    const declarationGroups = [[new GoFunctionDecl({
        name: 'probe', params: [], result, body: new GoReturn(new GoRaw(opaqueCode(source))),
    })]];
    const imports = collectImports(declarationGroups);
    assert.deepEqual(imports, expectedImports);
    const code = printGoFile({ packageName: 'probe', imports, declarationGroups });
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-imports-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    writeFileSync(join(directory, 'probe.go'), code);
    const build = spawnSync('go', ['build', '-o', join(directory, 'probe.a'), 'probe.go'], { cwd: directory, encoding: 'utf8' });
    assert.equal(build.status, 0, build.stdout + build.stderr + '\n' + code);
});
