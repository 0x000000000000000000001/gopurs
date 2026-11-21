import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { cpSync, mkdirSync, mkdtempSync, readFileSync, rmSync, symlinkSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { basename, delimiter, join, resolve } from 'node:path';
import { fileURLToPath, pathToFileURL } from 'node:url';

const root = fileURLToPath(new URL('../', import.meta.url));
const compiler = resolve(process.env.GOPURS_COMPILER_ROOT || join(root, '../..'));
const libraries = resolve(process.env.GOPURS_LIBRARY_ROOT || join(compiler, '..'));
const { corePackages } = await import(pathToFileURL(join(compiler, 'tools/test-workspace.mjs')));
const { findTypedCompiler } = await import(pathToFileURL(join(compiler, 'tools/build/native-workspace.mjs')));
const profile = process.argv.includes('--dev') ? 'dev' : 'release';
const race = process.argv.includes('--race');
const workspace = mkdtempSync(join(tmpdir(), 'gopurs-task-build-'));
const destination = join(root, '.build');
mkdirSync(join(workspace, 'bin'));
mkdirSync(join(destination, 'logs'), { recursive: true });
symlinkSync(findTypedCompiler(compiler, process.env.GOPURS_PURS), join(workspace, 'bin/purs'));
cpSync(join(root, 'src'), join(workspace, 'src'), { recursive: true });

const dependencies = ['aff', 'argonaut-codecs', 'arrays', 'console', 'effect', 'either', 'exceptions',
  'foldable-traversable', 'foreign', 'maybe', 'node-process', 'prelude', 'sqlite3', 'strings', 'transformers'];
const overrides = [...new Set([...corePackages(compiler), 'argonaut-core', 'argonaut-codecs',
  'foreign', 'foreign-object', 'record', 'ordered-collections', 'node-process', 'sqlite3'])];
// Freeze library inputs so an unrelated edit cannot change the program midway
// through generation. Keep private copies outside Spago's source discovery.
const snapshots = join(workspace, '.spago', 'ports');
const excluded = new Set(['.git', '.spago', '.cache', '.build', 'node_modules', 'output', 'output.bak']);
mkdirSync(snapshots, { recursive: true });
for (const name of overrides) cpSync(join(libraries, 'gopurs-' + name), join(snapshots, 'gopurs-' + name), {
  recursive: true, verbatimSymlinks: true, filter: path => !excluded.has(basename(path)),
});
writeFileSync(join(workspace, 'spago.yaml'),
  'package:\n  name: task-list-api\n  dependencies:\n' + dependencies.map(name => `    - ${name}\n`).join('') +
  'workspace:\n  packageSet:\n    registry: 77.10.1\n  extraPackages:\n' +
  overrides.map(name => `    ${name}:\n      path: ${JSON.stringify(join(snapshots, 'gopurs-' + name))}\n`).join(''));

const env = { ...process.env, PATH: [join(workspace, 'bin'), process.env.PATH].join(delimiter),
  GOPURS_JS: process.env.GOPURS_JS || '1', GOPURS_RUST: '0', GOPURS_DAEMON: '0', GOPURS_RESIDENT: '0',
  GOWORK: 'off', GOPURS_PBO_JOBS: '2', GOPURS_PREPARE_JOBS: '2' };
function run(label, command, args, cwd = workspace) {
  console.log(label);
  const result = spawnSync(command, args, { cwd, env, encoding: 'utf8', timeout: 180000, maxBuffer: 16 * 1024 * 1024 });
  const log = (result.stdout || '') + (result.stderr || '');
  writeFileSync(join(destination, 'logs', label + '.log'), log);
  assert.ifError(result.error);
  assert.equal(result.status, 0, `${label} failed in ${workspace}\n${log}`);
}
console.log(`Workspace: ${workspace}`);
run('frontend', 'spago', ['build', ...(process.env.GOPURS_OFFLINE === '1' ? ['--offline'] : [])]);
run('go-dependencies', process.execPath, [join(compiler, 'tools/library-go-dependencies.mjs'),
  join(workspace, 'output'), ...overrides.map(name => join(snapshots, 'gopurs-' + name))]);
run('generate', join(compiler, 'bin/gopurs'), ['--' + profile, '--main', 'Main']);
run('build', 'go', ['build', '-mod=readonly', ...(race ? ['-race'] : []),
  ...(profile === 'dev' ? ['-gcflags=all=-N -l'] : []), '-o', join(destination, 'task-list-api'), './main'], join(workspace, 'output'));
writeFileSync(join(destination, 'build.json'), JSON.stringify({ profile, race, compiler, libraries,
  host: env.GOPURS_JS === '1' ? 'javascript' : 'native',
  goModule: readFileSync(join(workspace, 'output/go.mod'), 'utf8') }, null, 2) + '\n');
if (process.env.GOPURS_TEST_KEEP_WORKSPACE !== '1') rmSync(workspace, { recursive: true, force: true });
console.log(`Built ${join(destination, 'task-list-api')}`);
