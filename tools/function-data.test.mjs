import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { copyFileSync, mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';

test('function metadata preserves ordinary calls, captures and immutable concurrent access', async t => {
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-function-data-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    copyFileSync(new URL('../runtime/runtime.go', import.meta.url),
        join(directory, 'gopurs_runtime/runtime.go'));
    // Reuse the arity/capture builders; run only the new contextual cases.
    for (const name of ['function-data_test.go', 'apply-arity_test.go', 'closure-lifetime_test.go']) {
        copyFileSync(new URL(`./${name}`, import.meta.url), join(directory, 'gopurs_runtime', name));
    }
    for (const race of [false, true]) {
        await t.test(race ? 'race detector' : 'optimized build', () => {
            const result = spawnSync('go', [
                'test', '-json', '-count=1', '-timeout=45s', '-run', '^TestFunctionData',
                ...(race ? ['-race'] : []), './gopurs_runtime',
            ], {
                cwd: directory, encoding: 'utf8', timeout: 120_000,
                env: { ...process.env, GOWORK: 'off' },
            });
            assert.ifError(result.error);
            assert.equal(result.status, 0, result.stdout + result.stderr);
            const passed = new Set(result.stdout.trim().split('\n')
                .map(line => JSON.parse(line))
                .filter(event => event.Action === 'pass' && event.Test)
                .map(event => event.Test));
            for (const name of ['ArityMatrix', 'Uncurried', 'TypesAndReplacement', 'Boxing',
                'Lifetime', 'Concurrent', 'SaturatedAllocations', 'InvalidFunction']) {
                assert.ok(passed.has(`TestFunctionData${name}`), `${name} did not pass`);
            }
            for (let arity = 1; arity <= 11; arity++) {
                for (let supplied = 1; supplied <= 11; supplied++) {
                    assert.ok(passed.has(`TestFunctionDataArityMatrix/arity_${arity}_apply_${supplied}`));
                }
            }
        });
    }
});
