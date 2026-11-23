// Handwritten FFI for Gopurs.FfiSupport; keep in sync with the .purs foreign imports.
import { execFileSync } from 'node:child_process';
import { closeSync, existsSync, mkdtempSync, openSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { fileURLToPath } from 'node:url';
import { dirname, join } from 'node:path';

function resolveFfiRunner() {
    const currentDir = dirname(fileURLToPath(import.meta.url));
    // The bundle lives in bin; source and Spago modules are one level deeper.
    const candidates = [
        join(currentDir, '..', 'tools', 'ffi-runner.mjs'),
        join(currentDir, '..', '..', 'tools', 'ffi-runner.mjs'),
    ];
    const runner = candidates.find(existsSync);
    if (!runner) {
        throw new Error(`Cannot locate the FFI runner from ${currentDir}`);
    }
    return runner;
}

function runFfiParser(content, prefix, sourcePrefix) {
    let directory, inputDescriptor, outputDescriptor;
    try {
        const args = [resolveFfiRunner()];
        if (prefix !== undefined) args.push(prefix);
        if (sourcePrefix !== undefined) args.push(sourcePrefix);
        // Complete programs use files in both directions. Besides avoiding the
        // response pipe limit, a completed request gives readFileSync(0) an
        // immediate EOF without waiting on Node's synchronous input pipe.
        if (prefix === '--package-source' || prefix === '--bind-source' || prefix === '--share-record-shapes') {
            directory = mkdtempSync(join(tmpdir(), 'gopurs-source-'));
            const request = join(directory, 'request.txt');
            writeFileSync(request, content, { encoding: 'utf8', flag: 'wx', mode: 0o600 });
            inputDescriptor = openSync(request, 'r');
            outputDescriptor = openSync(join(directory, 'response.json'), 'wx+', 0o600);
        }
        const response = execFileSync(process.execPath, args, {
            ...(inputDescriptor === undefined ? { input: content } : {}),
            encoding: 'utf-8',
            maxBuffer: 10 * 1024 * 1024,
            stdio: [inputDescriptor ?? 'pipe', outputDescriptor ?? 'pipe', 'pipe'],
        });
        return directory ? readFileSync(join(directory, 'response.json'), 'utf8') : response;
    } catch (error) {
        const detail = error.stderr?.trim() || error.message;
        throw new Error(`FFI runner failed: ${detail}`);
    } finally {
        try {
            if (inputDescriptor !== undefined) closeSync(inputDescriptor);
        } finally {
            try {
                if (outputDescriptor !== undefined) closeSync(outputDescriptor);
            } finally {
                if (directory) rmSync(directory, { recursive: true, force: true });
            }
        }
    }
}

export const extractFfiAstImpl = content => () => runFfiParser(content);

export const prepareFfiAstImpl = prefix => content => () => runFfiParser(content, prefix);

export const bindGoSourceImpl = content => () => runFfiParser(content, '--bind-source');
export const packageGoSourceImpl = content => () => runFfiParser(content, '--package-source');

export const packageGoFilesImpl = content => () => {
    const root = process.env.GOPURS_COMPILER_SNAPSHOT;
    if (!root) throw new Error('Compiler snapshot is required for Go package file emission');
    let directory, descriptor;
    try {
        // A completed file gives the helper an immediate EOF after the request.
        // Node's synchronous input pipe can otherwise stall with io.ReadAll.
        directory = mkdtempSync(join(tmpdir(), 'gopurs-packages-'));
        const request = join(directory, 'request.json');
        writeFileSync(request, content, { encoding: 'utf8', flag: 'wx', mode: 0o600 });
        descriptor = openSync(request, 'r');
        return execFileSync(join(root, 'tools/go-packages-helper'), [], {
            encoding: 'utf8', maxBuffer: 64 * 1024 * 1024,
            stdio: [descriptor, 'pipe', 'pipe'],
        }).trimEnd();
    } finally {
        try {
            if (descriptor !== undefined) closeSync(descriptor);
        } finally {
            if (directory) rmSync(directory, { recursive: true, force: true });
        }
    }
};

export const shareRecordShapesImpl = prefix => content => () => runFfiParser(content, '--share-record-shapes', prefix);
