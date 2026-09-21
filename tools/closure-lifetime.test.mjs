import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { copyFileSync, mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';

test('boxed closures preserve their captures, partial applications and concurrent callers', async t => {
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-closure-lifetime-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    copyFileSync(new URL('../runtime/runtime.go', import.meta.url),
        join(directory, 'gopurs_runtime/runtime.go'));
    copyFileSync(new URL('./closure-lifetime_test.go', import.meta.url),
        join(directory, 'gopurs_runtime/closure_lifetime_test.go'));

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
            assert.ok(passed.has('TestClosureLifetime'));
            assert.ok(passed.has('TestPartialApplicationLifetime'));
            assert.ok(passed.has('TestConcurrentClosureCreatorsAndCallers'));
            for (let arity = 1; arity <= 11; arity++) {
                assert.ok(passed.has(`TestClosureLifetime/arity_${arity}`));
                for (let prefix = 1; prefix < arity; prefix++) {
                    assert.ok(passed.has(`TestPartialApplicationLifetime/arity_${arity}_apply_${prefix}`));
                }
            }
        });
    }
});
