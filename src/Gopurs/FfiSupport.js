import { execFileSync } from 'node:child_process';
import { closeSync, existsSync, mkdtempSync, openSync, readFileSync, rmSync } from 'node:fs';
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

function runFfiParser(content, prefix) {
    let directory, descriptor;
    try {
        const args = [resolveFfiRunner()];
        if (prefix !== undefined) args.push(prefix);
        // Complete-program results can exceed the per-FFI pipe limit. Spool
        // those bytes directly, then decode once, without a second pipe copy.
        if (prefix === '--package-source' || prefix === '--bind-source') {
            directory = mkdtempSync(join(tmpdir(), 'gopurs-source-'));
            descriptor = openSync(join(directory, 'response.json'), 'wx+', 0o600);
        }
        const response = execFileSync(process.execPath, args, {
            input: content,
            encoding: 'utf-8',
            maxBuffer: 10 * 1024 * 1024,
            stdio: ['pipe', descriptor ?? 'pipe', 'pipe'],
        });
        return directory ? readFileSync(join(directory, 'response.json'), 'utf8') : response;
    } catch (error) {
        const detail = error.stderr?.trim() || error.message;
        throw new Error(`FFI runner failed: ${detail}`);
    } finally {
        if (descriptor !== undefined) closeSync(descriptor);
        if (directory) rmSync(directory, { recursive: true, force: true });
    }
}

export const extractFfiAstImpl = content => () => runFfiParser(content);

export const prepareFfiAstImpl = prefix => content => () => runFfiParser(content, prefix);

export const bindGoSourceImpl = content => () => runFfiParser(content, '--bind-source');
export const packageGoSourceImpl = content => () => runFfiParser(content, '--package-source');

export const packageGoFilesImpl = content => () => {
    const root = process.env.GOPURS_COMPILER_SNAPSHOT;
    if (!root) throw new Error('Compiler snapshot is required for Go package file emission');
    return execFileSync(join(root, 'tools/go-packages-helper'), [], {
        input: content, encoding: 'utf8', maxBuffer: 64 * 1024 * 1024,
        stdio: ['pipe', 'pipe', 'pipe'],
    }).trimEnd();
};
