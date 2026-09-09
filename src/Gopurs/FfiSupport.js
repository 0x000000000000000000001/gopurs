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

export const extractFfiAstImpl = function(moduleName) {
    return function(content) {
        return function() {
            try {
                return execFileSync(process.execPath, [resolveFfiRunner()], {
                    input: content,
                    encoding: 'utf-8',
                    maxBuffer: 10 * 1024 * 1024,
                });
            } catch (e) {
                console.error("FFI AST Extraction Error in module", moduleName);
                if (e.stdout) console.error(e.stdout.toString());
                if (e.stderr) console.error(e.stderr.toString());
                throw e;
            }
        };
    };
};
