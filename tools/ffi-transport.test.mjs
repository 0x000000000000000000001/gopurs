import assert from 'node:assert/strict';
import { spawn } from 'node:child_process';
import { copyFileSync, mkdirSync, mkdtempSync, rmSync, symlinkSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { dirname, join } from 'node:path';
import { fileURLToPath } from 'node:url';
import test from 'node:test';

const root = dirname(dirname(fileURLToPath(import.meta.url)));

async function runTransport(t, body, failingRunner = false) {
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-ffi-transport-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    const support = join(directory, 'src/Gopurs');
    const temporary = join(directory, 'temporary');
    mkdirSync(support, { recursive: true });
    mkdirSync(temporary);
    writeFileSync(join(directory, 'package.json'), '{"type":"module"}');
    copyFileSync(join(root, 'src/Gopurs/FfiSupport.js'), join(support, 'FfiSupport.js'));
    if (failingRunner) {
        mkdirSync(join(directory, 'tools'));
        writeFileSync(join(directory, 'tools/ffi-runner.mjs'),
            'process.stderr.write("transport failure"); process.exitCode = 1;');
    } else {
        symlinkSync(join(root, 'tools'), join(directory, 'tools'), 'dir');
    }
    writeFileSync(join(directory, 'probe.mjs'), `
        import assert from 'node:assert/strict';
        import { readdirSync } from 'node:fs';
        import { tmpdir } from 'node:os';
        import * as api from './src/Gopurs/FfiSupport.js';
        ${body}
        assert.deepEqual(readdirSync(tmpdir()), [], 'temporary request/response files leaked');
        process.stdout.write('ok');
    `);
    const child = spawn(process.execPath, [join(directory, 'probe.mjs')], {
        env: { ...process.env, TMPDIR: temporary, TMP: temporary, TEMP: temporary },
        detached: process.platform !== 'win32', stdio: ['ignore', 'pipe', 'pipe'],
    });
    let stdout = '', stderr = '', timedOut = false;
    child.stdout.setEncoding('utf8').on('data', data => { stdout += data; });
    child.stderr.setEncoding('utf8').on('data', data => { stderr += data; });
    const timer = setTimeout(() => {
        timedOut = true;
        if (process.platform === 'win32') child.kill('SIGKILL');
        else process.kill(-child.pid, 'SIGKILL');
    }, 15_000);
    const code = await new Promise((resolve, reject) => {
        child.once('error', reject);
        child.once('close', resolve);
    }).finally(() => clearTimeout(timer));
    assert.equal(timedOut, false, 'complete-program FFI transport did not finish before its deadline');
    assert.equal(code, 0, stderr);
    assert.equal(stderr, '');
    assert.equal(stdout, 'ok');
}

test('record shape sharing preserves programs larger than the response pipe limit', async t => {
    await runTransport(t, `
        const comment = '// ' + 'x'.repeat(12 * 1024 * 1024) + ' é λ 😀\\n';
        const source = 'package p\\nimport r "gopurs/output/gopurs_runtime"\\n' + comment
            + 'var value = r.RecordDict2("line", "column", r.Int(1), r.Int(2))\\n';
        const result = api.shareRecordShapesImpl('transport')(source)();
        assert.ok(result.includes(comment), 'source bytes were truncated or altered');
        assert.match(result, /r\\.RecordShape2\\(&GopursRecordKeys_/);
        assert.doesNotMatch(result, /r\\.RecordDict2/);
    `);
});

test('binding and packaging preserve large requests and responses', async t => {
    await runTransport(t, `
        const payload = 'x'.repeat(12 * 1024 * 1024) + ' é λ 😀';
        const content = 'package purescript\\nconst Payload = ' + JSON.stringify(payload)
            + '\\nfunc Read() string { return Payload }\\n';
        const request = { packagePath: 'example/purescript', goVersion: 'go1.26', arch: 'arm64', exports: {},
            files: [{ owner: 'generated/Probe', moduleName: 'Probe', path: 'Probe.go', content, initializers: [] }],
            entries: [{ owner: 'generated/Probe', name: 'Read' }] };
        const bound = JSON.parse(api.bindGoSourceImpl(JSON.stringify(request))());
        assert.equal(bound.files[0].text, content);
        const packaged = JSON.parse(api.packageGoSourceImpl(JSON.stringify(request))());
        assert.ok(packaged.files.some(file => file.content.includes(payload)), 'packaged source was truncated');
    `);
});

test('complete-program runner failures preserve diagnostics and remove temporary files', async t => {
    await runTransport(t, `
        for (const call of [api.bindGoSourceImpl, api.packageGoSourceImpl, api.shareRecordShapesImpl('transport')]) {
            assert.throws(() => call('x'.repeat(2 * 1024 * 1024))(), /FFI runner failed: transport failure/);
            assert.deepEqual(readdirSync(tmpdir()), [], 'failed request left temporary files');
        }
    `, true);
});
