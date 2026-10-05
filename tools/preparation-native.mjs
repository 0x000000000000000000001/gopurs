import assert from 'node:assert/strict';
import { copyFileSync, existsSync, mkdirSync, writeFileSync } from 'node:fs';
import { join, resolve } from 'node:path';
import { spawnSync } from 'node:child_process';

const [rustArg, outArg] = process.argv.slice(2);
assert(rustArg && outArg, 'preparation-native.mjs GENERATED_COMPILER_RUST NEW_TEST_ARCHIVE');
const rust = resolve(rustArg), out = resolve(outArg);
assert(!existsSync(out)); mkdirSync(join(out, 'src'), { recursive: true });
const modules = ['purust_core', 'Purs_Gopurs_Preparation', 'Purs_Effect_Aff'];
writeFileSync(join(out, 'Cargo.toml'), '[package]\nname = "gopurs_preparation_test"\nversion = "0.0.0"\nedition = "2021"\n' +
  '[profile.release]\nopt-level = 3\ndebug = false\nlto = false\n[dependencies]\n' +
  modules.map(name => `${name} = { path = ${JSON.stringify(join(rust, name))} }\n`).join(''));
copyFileSync(new URL('./preparation-native.rs', import.meta.url), join(out, 'src/main.rs'));
const target = process.env.CARGO_TARGET_DIR ?? join(out, 'target');
const build = spawnSync('cargo', ['build', '--offline', '--release', '--manifest-path', join(out, 'Cargo.toml'), '--target-dir', target],
  { env: { ...process.env, CARGO_BUILD_JOBS: '8', CARGO_INCREMENTAL: '0' }, encoding: 'utf8', timeout: 900000, maxBuffer: 32 * 1024 * 1024 });
writeFileSync(join(out, 'build.stdout'), build.stdout ?? ''); writeFileSync(join(out, 'build.stderr'), build.stderr ?? '');
assert.equal(build.status, 0, build.error?.message ?? build.stderr);
const run = spawnSync(join(target, 'release/gopurs_preparation_test'), [], { encoding: 'utf8', timeout: 120000 });
writeFileSync(join(out, 'test.stdout'), run.stdout ?? ''); writeFileSync(join(out, 'test.stderr'), run.stderr ?? '');
assert.equal(run.status, 0, run.error?.message ?? run.stderr);
console.log(run.stdout.trim());
