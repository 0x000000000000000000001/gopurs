import { execFileSync } from 'node:child_process';
import { existsSync } from 'node:fs';
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
    try {
        const args = [resolveFfiRunner()];
        if (prefix !== undefined) args.push(prefix);
        return execFileSync(process.execPath, args, {
            input: content,
            encoding: 'utf-8',
            maxBuffer: 10 * 1024 * 1024,
            stdio: ['pipe', 'pipe', 'pipe'],
        });
    } catch (error) {
        const detail = error.stderr?.trim() || error.message;
        throw new Error(`FFI runner failed: ${detail}`);
    }
}

export const extractFfiAstImpl = content => () => runFfiParser(content);

export const prepareFfiAstImpl = prefix => content => () => runFfiParser(content, prefix);
