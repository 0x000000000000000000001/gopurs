// GOPURS_NATIVE_OUTPUT=/path/to/bootstrap/output node --test tools/preparation-native.test.mjs
import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { existsSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join, resolve } from 'node:path';
import test from 'node:test';

const nativeOutput = process.env.GOPURS_NATIVE_OUTPUT;

test('native preparation defers work, overlaps bounded workers, and preserves result order', {
    skip: nativeOutput ? false : 'requires GOPURS_NATIVE_OUTPUT from a native bootstrap',
}, t => {
    const output = resolve(nativeOutput);
    assert.ok(existsSync(join(output, 'purescript/Gopurs_Preparation.go')),
        'bootstrap must include Gopurs.Preparation');
    // Materialize the actual test source for Go's test discovery. This is a
    // retained, generated bootstrap; never overwrite an existing file.
    const nativeTest = join(output, 'purescript', 'preparation_native_test.go');
    writeFileSync(nativeTest, readFileSync(new URL('./preparation-native_test.go', import.meta.url)), { flag: 'wx' });
    t.after(() => rmSync(nativeTest, { force: true }));
    // Compilation includes the whole native compiler; keep its cold-cache
    // budget separate from the 30-second execution limit of these contracts.
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-preparation-native-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    const executable = join(directory, 'preparation.test');
    const build = spawnSync('go', [
        'test', '-race', '-c', '-o', executable, './purescript',
    ], {
        cwd: output, encoding: 'utf8', timeout: 900_000,
        env: { ...process.env, GOWORK: 'off' }, maxBuffer: 1024 * 1024,
    });
    assert.ifError(build.error);
    assert.equal(build.status, 0, build.stdout + build.stderr);
    const result = spawnSync(executable, [
        '-test.count=1', '-test.timeout=30s', '-test.run=^TestPreparationNative', '-test.v',
    ], {
        cwd: output, encoding: 'utf8', timeout: 60_000,
        env: { ...process.env, GOWORK: 'off' }, maxBuffer: 1024 * 1024,
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stdout + result.stderr);
    assert.ok(!(result.stdout + result.stderr).includes('[no test files]'), result.stdout + result.stderr);
    for (const name of [
        'TestPreparationNativeDefersAndReruns',
        'TestPreparationNativeParallelBoundAndOrder',
        'TestPreparationNativeParallelBoundAndOrder/jobs=3',
        'TestPreparationNativeParallelBoundAndOrder/jobs-clamped-to-8',
        'TestPreparationNativeSequential',
        'TestPreparationNativeSequential/jobs=-2',
        'TestPreparationNativeSequential/jobs=0',
        'TestPreparationNativeSequential/jobs=1',
    ]) {
        assert.ok(result.stdout.includes(`--- PASS: ${name} (`),
            `expected native test did not pass: ${name}\n${result.stdout}${result.stderr}`);
    }
    t.diagnostic(result.stdout.trim());
});
