import assert from 'node:assert/strict';
import { createHash } from 'node:crypto';
import fs from 'node:fs';
import { syncBuiltinESMExports } from 'node:module';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { packageGoFilesImpl } from '../src/Gopurs/FfiSupport.js';

// Exercise the production source without rebuilding the compiler or Go helper.
// A real child process checks the transport, bytes, EOF and error propagation.
function withHelper(helper, action, { failOpen, failWrite } = {}) {
    const directory = fs.mkdtempSync(join(tmpdir(), 'gopurs-package-transport-test-'));
    const snapshot = join(directory, 'snapshot with spaces');
    const spool = join(directory, 'spool');
    fs.mkdirSync(join(snapshot, 'tools'), { recursive: true });
    fs.mkdirSync(spool);
    fs.writeFileSync(join(snapshot, 'tools/go-packages-helper'),
        `#!${process.execPath}\nconst fs = require('node:fs');\n${helper}\n`, { mode: 0o700 });
    const oldSnapshot = process.env.GOPURS_COMPILER_SNAPSHOT;
    const oldTmp = process.env.TMPDIR;
    const originalOpen = fs.openSync;
    const originalWrite = fs.writeFileSync;
    const descriptors = [];
    process.env.GOPURS_COMPILER_SNAPSHOT = snapshot;
    process.env.TMPDIR = spool;
    fs.openSync = function (path, flags, ...args) {
        const request = typeof path === 'string' && path.startsWith(spool + '/') && flags === 'r';
        if (request && failOpen) throw failOpen;
        const descriptor = originalOpen(path, flags, ...args);
        if (request) descriptors.push(descriptor);
        return descriptor;
    };
    fs.writeFileSync = function (path, data, ...args) {
        if (typeof path === 'string' && path.startsWith(spool + '/') && failWrite) {
            originalWrite(path, 'partial request', ...args);
            throw failWrite;
        }
        return originalWrite(path, data, ...args);
    };
    syncBuiltinESMExports();
    try {
        return action({ snapshot, spool, descriptors,
            assertClean(expectedDescriptors = 1) {
                assert.equal(descriptors.length, expectedDescriptors);
                for (const descriptor of descriptors) {
                    assert.throws(() => fs.fstatSync(descriptor), { code: 'EBADF' }, 'Request descriptor must be closed');
                }
                assert.deepEqual(fs.readdirSync(spool), [], 'Private request directory must be removed');
            } });
    } finally {
        fs.openSync = originalOpen;
        fs.writeFileSync = originalWrite;
        syncBuiltinESMExports();
        if (oldSnapshot === undefined) delete process.env.GOPURS_COMPILER_SNAPSHOT;
        else process.env.GOPURS_COMPILER_SNAPSHOT = oldSnapshot;
        if (oldTmp === undefined) delete process.env.TMPDIR;
        else process.env.TMPDIR = oldTmp;
        fs.rmSync(directory, { recursive: true, force: true });
    }
}

test('Go package transport uses a private read-only file and preserves large UTF-8 bytes and trimEnd', () => {
    const payload = '  ' + JSON.stringify({ text: 'é😀λ\u0000'.repeat(256 * 1024), files: [] });
    assert.ok(Buffer.byteLength(payload) > 1024 * 1024);
    withHelper(`
const stat = fs.fstatSync(0);
if (!stat.isFile()) throw new Error('stdin is not a regular file');
if (stat.mode & 0o077) throw new Error('request is not private');
try { fs.writeSync(0, 'x'); throw new Error('stdin is writable'); }
catch (error) { if (error.code !== 'EBADF') throw error; }
const request = fs.readFileSync(0);
if (fs.readSync(0, Buffer.alloc(1), 0, 1, null) !== 0) throw new Error('missing EOF');
if (!JSON.parse(request.toString('utf8')).text.includes('é😀λ')) throw new Error('corrupt UTF-8');
process.stdout.write(request);
process.stdout.write('\\n \\t');
`, ({ assertClean }) => {
        const response = packageGoFilesImpl(payload)();
        assert.equal(Buffer.byteLength(response), Buffer.byteLength(payload));
        const hash = value => createHash('sha256').update(value).digest('hex');
        assert.equal(hash(response), hash(payload));
        assert.ok(response.startsWith('  {'), 'Leading whitespace must remain intact');
        assertClean();
    });
});

test('Go package helper errors preserve status, stdout and stderr and close/remove the request', () => {
    withHelper(`
if (!fs.fstatSync(0).isFile()) throw new Error('stdin is not a file');
if (fs.readFileSync(0, 'utf8') !== '{"files":[]}') throw new Error('wrong request');
process.stdout.write('partial output\\n');
process.stderr.write('package failure: é😀\\n');
process.exitCode = 23;
`, ({ assertClean }) => {
        assert.throws(() => packageGoFilesImpl('{"files":[]}')(), error => {
            assert.equal(error.status, 23);
            assert.equal(error.signal, null);
            assert.equal(error.stdout, 'partial output\n');
            assert.equal(error.stderr, 'package failure: é😀\n');
            return true;
        });
        assertClean();
    });
});

test('Go package transport closes and removes the request when the frozen helper is missing', () => {
    withHelper('', ({ snapshot, assertClean }) => {
        fs.rmSync(join(snapshot, 'tools/go-packages-helper'));
        assert.throws(() => packageGoFilesImpl('{"files":[]}')(), { code: 'ENOENT' });
        assertClean();
    });
});

test('Go package transport removes a written request when opening it for reading fails', () => {
    const failure = Object.assign(new Error('denied request descriptor'), { code: 'EACCES' });
    withHelper('', ({ assertClean }) => {
        assert.throws(() => packageGoFilesImpl('{"files":[]}')(), error => error === failure);
        assertClean(0);
    }, { failOpen: failure });
});

test('Go package transport removes a partially written request when writing fails', () => {
    const failure = Object.assign(new Error('request disk full'), { code: 'ENOSPC' });
    withHelper('', ({ assertClean }) => {
        assert.throws(() => packageGoFilesImpl('{"files":[]}')(), error => error === failure);
        assertClean(0);
    }, { failWrite: failure });
});

test('Go package transport requires its frozen compiler before creating a request', () => {
    withHelper('', ({ assertClean }) => {
        delete process.env.GOPURS_COMPILER_SNAPSHOT;
        assert.throws(() => packageGoFilesImpl('{}')(), /Compiler snapshot is required/);
        assertClean(0);
    });
});
