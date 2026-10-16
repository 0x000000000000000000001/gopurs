import { createHash } from 'node:crypto';
import { execFileSync } from 'node:child_process';
import { existsSync, mkdirSync, mkdtempSync, readFileSync, renameSync, rmSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';

const script = fileURLToPath(import.meta.url);
const tools = dirname(script);
const sources = join(tools, 'ffi-gen');

export function resolveGoPackagesHelper() {
    const hash = createHash('sha256');
    for (const name of ['go.mod', 'bindings.go', 'packages.go', 'workspace.go', 'main_workspace.go']) {
        const content = readFileSync(join(sources, name));
        hash.update(`${name}:${content.length}:`).update(content);
    }
    const directory = join(tools, '../node_modules/.cache/gopurs-go-packages');
    const binary = join(directory, `${process.platform}-${process.arch}-${hash.digest('hex')}`);
    if (!existsSync(binary)) {
        mkdirSync(directory, { recursive: true });
        const staging = mkdtempSync(join(directory, '.build-'));
        try {
            execFileSync('go', ['build', '-trimpath', '-buildvcs=false', '-tags=packagecli', '-o', join(staging, 'helper'),
                'bindings.go', 'packages.go', 'workspace.go', 'main_workspace.go'], {
                cwd: sources,
                env: { ...process.env, GOWORK: 'off', GOFLAGS: '', CGO_ENABLED: '0',
                    GOOS: process.platform, GOARCH: { x64: 'amd64', arm64: 'arm64' }[process.arch] ?? process.arch },
                stdio: ['ignore', 'pipe', 'pipe'],
            });
            renameSync(join(staging, 'helper'), binary);
        } finally {
            rmSync(staging, { recursive: true, force: true });
        }
    }
    return binary;
}

if (process.argv[1] && resolve(process.argv[1]) === script) {
    process.stdout.write(resolveGoPackagesHelper() + '\n');
}
