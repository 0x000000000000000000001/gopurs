import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { copyFileSync, mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';

test('batched application preserves arity, order and captures without saturated allocations', async t => {
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-apply-arity-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    copyFileSync(new URL('../runtime/runtime.go', import.meta.url),
        join(directory, 'gopurs_runtime/runtime.go'));
    for (const name of ['apply-arity_test.go', 'closure-lifetime_test.go']) {
        copyFileSync(new URL(`./${name}`, import.meta.url), join(directory, 'gopurs_runtime', name));
    }

    for (const race of [false, true]) {
        await t.test(race ? 'race detector' : 'optimized build', () => {
            const result = spawnSync('go', [
                'test', '-json', '-count=1', '-timeout=45s',
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
            for (const name of ['TestApplyArityMatrix', 'TestSaturatedApplicationAllocations', 'TestApplicationPanics',
                'TestClosureLifetime', 'TestPartialApplicationLifetime', 'TestConcurrentClosureCreatorsAndCallers']) {
                assert.ok(passed.has(name), `${name} did not pass`);
            }
            for (let arity = 1; arity <= 11; arity++) {
                for (let supplied = 2; supplied <= 10; supplied++) {
                    assert.ok(passed.has(`TestApplyArityMatrix/arity_${arity}_apply_${supplied}`));
                }
            }
        });
    }
});
