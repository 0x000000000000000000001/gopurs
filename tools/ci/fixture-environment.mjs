import assert from 'node:assert/strict';
import { createHash } from 'node:crypto';
import { execFileSync, spawnSync } from 'node:child_process';
import { closeSync, existsSync, lstatSync, mkdirSync, mkdtempSync, openSync, readFileSync, readSync, readdirSync, realpathSync, rmSync, writeFileSync } from 'node:fs';
import { arch, platform } from 'node:os';
import { dirname, isAbsolute, join, relative, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';

const root = fileURLToPath(new URL('../../', import.meta.url));
const storePath = /^\/nix\/store\/[0-9abcdfghijklmnpqrsvwxyz]{32}-[A-Za-z0-9+._?=-]+$/;
const sha256 = value => createHash('sha256').update(value).digest('hex');
const fileHash = file => {
  const hash = createHash('sha256'), block = Buffer.alloc(1024 ** 2), fd = openSync(file, 'r');
  try { for (let size; (size = readSync(fd, block, 0, block.length, null)) > 0;) hash.update(block.subarray(0, size)); }
  finally { closeSync(fd); }
  return hash.digest('hex');
};
const digest = value => assert.match(value, /^[a-f0-9]{64}$/);
const positive = value => assert.match(value, /^[1-9][0-9]*$/);
const canonical = value => JSON.stringify(value);
const offline = ['--offline', '--no-substitute', '--max-jobs', '0', '--builders', ''];
// The downstream acquisition has a shared 1 GiB extraction budget. Reserve
// 128 MiB for the tested candidate beside the compressed runtime cache.
export const environmentBudget = 896 * 1024 ** 2;

function command(program, args, { timeout = 600_000, cwd = root } = {}) {
  console.error(`[fixture environment] ${program} ${args.join(' ')}`);
  return execFileSync(program, args, { cwd, encoding: 'utf8', timeout, maxBuffer: 16 * 1024 ** 2, stdio: ['ignore', 'pipe', 'inherit'] }).trim();
}

export function closureIdentity(value) {
  const entries = Array.isArray(value) ? value : Object.entries(value).map(([path, details]) => ({ path, ...details }));
  const result = entries.map(({ path, narHash, narSize, references }) => {
    assert.match(path, storePath); assert.ok(!path.endsWith('.drv'), 'Only runtime outputs may be transferred');
    assert.match(narHash, /^sha256-[A-Za-z0-9+/]{43}=$/);
    assert.ok(Number.isSafeInteger(narSize) && narSize > 0);
    assert.ok(Array.isArray(references)); references.forEach(ref => assert.match(ref, storePath));
    assert.equal(new Set(references).size, references.length);
    return { path, narHash, narSize, references: [...references].sort() };
  }).sort((a, b) => a.path.localeCompare(b.path));
  assert.ok(result.length > 0); assert.equal(new Set(result.map(entry => entry.path)).size, result.length);
  const paths = new Set(result.map(entry => entry.path));
  for (const entry of result) for (const ref of entry.references) assert.ok(paths.has(ref), 'Incomplete runtime closure');
  return result;
}

export function cacheInventory(directory) {
  const result = [];
  function walk(current) {
    for (const name of readdirSync(current).sort()) {
      const file = join(current, name), stat = lstatSync(file), path = relative(directory, file);
      assert.ok(!stat.isSymbolicLink(), 'Cache links are not transferable');
      if (stat.isDirectory()) {
        // Nix versions create different auxiliary cache directories. Accept
        // them only when empty; metadata and build logs must never be ignored.
        if (['log', 'build-trace-v2', 'realisations'].includes(path)) assert.equal(readdirSync(file).length, 0, `Unexpected contents in runtime cache directory ${path}`);
        else { assert.equal(path, 'nar'); walk(file); }
      }
      else {
        assert.ok(stat.isFile(), 'Nonregular cache entry');
        assert.match(path, /^(?:nix-cache-info|[0-9abcdfghijklmnpqrsvwxyz]{32}\.narinfo|nar\/[A-Za-z0-9+._=-]+\.nar(?:\.(?:xz|zst|bz2|gz))?)$/);
        result.push({ path, bytes: stat.size, sha256: fileHash(file) });
      }
    }
  }
  walk(directory); assert.ok(result.some(entry => entry.path.endsWith('.narinfo')));
  return result;
}

export function validateManifest(manifest, context) {
  assert.equal(manifest.schema, 1);
  assert.deepEqual(manifest.source, context.source); assert.equal(manifest.target, context.target);
  assert.deepEqual(manifest.compiler, context.compiler);
  for (const value of Object.values(manifest.compiler)) digest(value);
  assert.deepEqual(manifest.github, context.github);
  positive(manifest.github.runId); positive(manifest.github.runAttempt);
  assert.match(manifest.environment?.profile, storePath);
  assert.ok(manifest.environment.profile.endsWith('-env'), 'Expected a recorded Nix build environment');
  digest(manifest.environment.profileSha256); digest(manifest.environment.closureSha256);
  const closure = closureIdentity(manifest.closure);
  assert.deepEqual(manifest.closure, closure);
  assert.ok(closure.some(entry => entry.path === manifest.environment.profile));
  assert.match(manifest.environment.shell, /^\/nix\/store\/[0-9abcdfghijklmnpqrsvwxyz]{32}-[A-Za-z0-9+._?=-]+\/bin\/bash$/);
  assert.ok(closure.some(entry => entry.path === dirname(dirname(manifest.environment.shell))), 'Recorded Bash is absent from runtime closure');
  assert.equal(sha256(canonical(closure)), manifest.environment.closureSha256, 'Runtime closure identity differs');
  assert.ok(Array.isArray(manifest.cache) && manifest.cache.length > 0);
  return manifest;
}

export function validateCoreEnvironment(manifest, coreIdentity) {
  assert.deepEqual(coreIdentity, { source: { commit: manifest.source.commit, dirty: false },
    compiler: manifest.compiler, target: manifest.target, github: manifest.github,
    environment: manifest.environment }, 'Core qualification and transferred environment differ');
}

export function printEnvironmentCommand(profile) {
  assert.match(profile, storePath); assert.ok(profile.endsWith('-env'));
  // `nix develop <profile>` still looks up nixpkgs#bashInteractive. Rendering
  // the recorded environment avoids that hidden registry/network dependency.
  return ['print-dev-env', profile, ...offline];
}

export function environmentScript(script) {
  // print-dev-env already evaluates shellHook exactly once. Stop before the
  // requested command if any environment preparation or hook command fails.
  return 'set -e\n' + script + '\nexec "$@"\n';
}

export function runProfile(environment, args, env = process.env) {
  assert.ok(Array.isArray(args) && args.length > 0);
  assert.match(environment.shell, /^\/nix\/store\/[0-9abcdfghijklmnpqrsvwxyz]{32}-[A-Za-z0-9+._?=-]+\/bin\/bash$/);
  const script = command('nix', printEnvironmentCommand(environment.profile), { timeout: 60_000 });
  const directory = mkdtempSync(join(realpathSync(env.RUNNER_TEMP), 'gopurs-fixture-command-'));
  try {
    const file = join(directory, 'environment.sh');
    writeFileSync(file, environmentScript(script), { flag: 'wx', mode: 0o700 });
    return spawnSync(environment.shell, [file, ...args], { cwd: root, env: { ...env, NIX_GCROOT: environment.profile }, stdio: 'inherit' });
  } finally { rmSync(directory, { recursive: true, force: true }); }
}

function sourceContext(env) {
  const git = (...args) => command('git', args, { timeout: 10_000 });
  const commit = git('rev-parse', 'HEAD'); assert.equal(commit, env.GITHUB_SHA);
  assert.equal(git('status', '--porcelain', '--untracked-files=all'), '');
  const target = `${platform()}-${{ x64: 'amd64', arm64: 'arm64' }[arch()]}`;
  assert.equal(target, env.GOPURS_RELEASE_TARGET);
  assert.equal(env.CI, 'true'); assert.equal(env.GITHUB_ACTIONS, 'true');
  return { source: { commit, tree: git('rev-parse', 'HEAD^{tree}') }, target,
    compiler: { javascriptSha256: fileHash(join(root, 'bin/gopurs.js')), nativeSha256: fileHash(join(root, 'bin/gopurs-native')) },
    github: { repository: env.GITHUB_REPOSITORY, runId: env.GITHUB_RUN_ID, runAttempt: env.GITHUB_RUN_ATTEMPT } };
}

function exportEnvironment(profile, directory, env) {
  assert.ok(!existsSync(directory)); mkdirSync(directory, { mode: 0o700 });
  const original = realpathSync(profile); assert.match(original, storePath); assert.ok(original.endsWith('-env'));
  const context = sourceContext(env);
  const closure = closureIdentity(JSON.parse(command('nix', ['path-info', '--json', '--recursive', ...offline, original])));
  const cache = join(directory, 'cache');
  command('nix', ['copy', '--to', `file://${cache}?compression=zstd`, ...offline, original]);
  const shell = JSON.parse(readFileSync(original, 'utf8')).variables?.SHELL?.value;
  const manifest = { schema: 1, ...context,
    environment: { profile: original, profileSha256: fileHash(original), closureSha256: sha256(canonical(closure)), shell },
    closure, cache: cacheInventory(cache) };
  validateManifest(manifest, context);
  const bytes = manifest.cache.reduce((sum, entry) => sum + entry.bytes, 0);
  assert.ok(bytes < environmentBudget, 'Prepared runtime cache exceeds the reserved 896 MiB acquisition budget');
  const manifestPath = join(directory, 'manifest.json');
  writeFileSync(manifestPath, JSON.stringify(manifest, null, 2) + '\n', { flag: 'wx', mode: 0o600 });
  console.log(`Exported ${closure.length} runtime paths (${bytes} compressed bytes)`);
  writeFileSync(env.GITHUB_ENV, `GOPURS_QUALITY_ENVIRONMENT=${manifestPath}\n`, { flag: 'a' });
}

function importEnvironment(directory, contextPath, profile, env) {
  const selected = JSON.parse(readFileSync(contextPath, 'utf8'));
  const manifestPath = join(directory, 'manifest.json'), manifest = JSON.parse(readFileSync(manifestPath, 'utf8'));
  const context = sourceContext(env), artifact = selected.artifacts.find(value => value.logicalName === `fixture-environment-${context.target}`);
  assert.equal(selected.source.commit, context.source.commit); assert.equal(selected.target, context.target);
  assert.equal(selected.github.repository, context.github.repository); assert.equal(selected.github.runId, context.github.runId);
  assert.equal(selected.github.finalizationAttempt, context.github.runAttempt);
  assert.ok(artifact); positive(artifact.attempt); assert.ok(Number(artifact.attempt) <= Number(context.github.runAttempt));
  assert.equal(selected.jobs.length, 1);
  assert.deepEqual(selected.jobs[0], { id: artifact.producerJobId, name: `build and test (${context.target})`, attempt: artifact.attempt });
  context.github.runAttempt = artifact.attempt;
  validateManifest(manifest, context);
  validateCoreEnvironment(manifest, selected.qualificationIdentity);
  assert.deepEqual(cacheInventory(join(directory, 'cache')), manifest.cache, 'Transferred cache bytes differ');
  command('nix', ['copy', '--from', `file://${join(directory, 'cache')}`, '--no-check-sigs', ...offline, manifest.environment.profile]);
  assert.equal(fileHash(manifest.environment.profile), manifest.environment.profileSha256, 'Imported profile bytes differ');
  const imported = closureIdentity(JSON.parse(command('nix', ['path-info', '--json', '--recursive', ...offline, manifest.environment.profile])));
  assert.deepEqual(imported, manifest.closure, 'Imported runtime closure differs');
  assert.ok(!existsSync(profile));
  command('nix-store', ['--add-root', profile, '--realise', manifest.environment.profile,
    '--option', 'substitute', 'false', '--option', 'max-jobs', '0', '--option', 'builders', '']);
  console.log(`Imported ${imported.length} runtime paths from original build attempt ${artifact.attempt}`);
  writeFileSync(env.GITHUB_ENV, `GOPURS_QUALITY_ENVIRONMENT=${manifestPath}\nGOPURS_FIXTURE_PROFILE=${manifest.environment.profile}\n`, { flag: 'a' });
}

function main(args, env = process.env) {
  const mode = args.shift();
  if (mode === 'run') {
    assert.equal(args.shift(), '--');
    const started = Date.now(); console.error(`[fixture environment] ${new Date(started).toISOString()} start ${args.join(' ')}`);
    const manifest = JSON.parse(readFileSync(env.GOPURS_QUALITY_ENVIRONMENT, 'utf8'));
    assert.equal(manifest.environment.profile, env.GOPURS_FIXTURE_PROFILE);
    assert.equal(fileHash(manifest.environment.profile), manifest.environment.profileSha256);
    const result = runProfile(manifest.environment, args, env);
    console.error(`[fixture environment] ${new Date().toISOString()} end ${args[0]} after ${Date.now() - started}ms: exit=${result.status} signal=${result.signal}`);
    if (result.error) throw result.error;
    process.exitCode = result.status ?? 1;
    return;
  }
  const options = {};
  for (let index = 0; index < args.length; index += 2) {
    assert.ok(['--profile', '--directory', '--context'].includes(args[index]) && args[index + 1] && !options[args[index]]);
    options[args[index]] = resolve(args[index + 1]);
  }
  assert.ok(['export', 'import'].includes(mode) && options['--profile'] && options['--directory']);
  const temp = realpathSync(env.RUNNER_TEMP);
  for (const path of Object.values(options)) {
    const rel = relative(temp, path); assert.ok(rel && !rel.startsWith('..') && !isAbsolute(rel));
    const parent = realpathSync(dirname(path)); assert.ok(parent === temp || parent.startsWith(temp + '/'));
  }
  if (mode === 'export') { assert.ok(!options['--context']); exportEnvironment(options['--profile'], options['--directory'], env); }
  else { assert.ok(options['--context']); importEnvironment(options['--directory'], options['--context'], options['--profile'], env); }
}

if (process.argv[1] && resolve(process.argv[1]) === fileURLToPath(import.meta.url)) main(process.argv.slice(2));
