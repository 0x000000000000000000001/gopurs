import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, readFileSync, readdirSync, rmSync, symlinkSync, writeFileSync } from 'node:fs';
import { delimiter, join } from 'node:path';
import { fileURLToPath } from 'node:url';
import test from 'node:test';
import { corePackages, createWorkspace, prepareFixture } from './test-workspace.mjs';
import { findTypedCompiler } from './build/native-workspace.mjs';

const root = fileURLToPath(new URL('../', import.meta.url));
const modes = [{ name: 'js', GOPURS_JS: '1' }, { name: 'native-parallel', GOPURS_JS: '0' }]
    .filter(mode => !process.env.GOPURS_TEST_MODES || process.env.GOPURS_TEST_MODES.split(',').includes(mode.name));
const profiles = (process.env.GOPURS_TEST_PROFILES || 'release,dev').split(',');
const fixture = join(root, 'tools/fixtures/codepoint-take-fold.purs');
const inputs = [[], [97,98,99], [233,8364,128512], [55296], [56320], [55296,56320], [56320,55296], [55296,97,56320]].map(points=>Array.from(String.fromCodePoint(...points), c=>c.codePointAt(0)));
const counts = [-2,0,1,2,3,10,2147483647];
const parse = stdout => stdout.trim().split('\n').map(line => JSON.parse(line));
function expected(seed=7, multiplier=3, trace=false) {
    const out=[];
    const consume=values=>values.reduce((total,value)=>{if(trace)out.push(100+value);return(Math.imul(total,multiplier)+value)|0;},seed);
    for(const values of inputs){for(const count of counts){const taken=values.slice(0,Math.max(0,count));
        const direct=consume(taken),indirect=consume(taken),shared=(consume(taken)+taken.length)|0;
        const captured=taken.reduce((total,value)=>(total+value+5)|0,7),unknown=taken.reduce((total,value)=>(total+value)|0,7);
        out.push([direct,indirect,shared,captured,unknown]);
    }out.push(values);}
    for(const count of [-1,0,2])out.push(consume([97,98,99].slice(0,Math.max(0,count))),[1,2]);
    out.push('count failure',[1],'input failure',[1,2]);return out;
}
function generatedText(directory) {
    return readdirSync(directory, { withFileTypes: true }).filter(entry => entry.name !== '.gopurs-cache')
        .map(entry => { const path = join(directory, entry.name); return entry.isDirectory() ? generatedText(path) : entry.name.endsWith('.go') ? readFileSync(path, 'utf8') : ''; }).join('\n');
}
function worker(source, name) {
    const start = source.indexOf(`func ${name}(`); if (start < 0) return '';
    const end = source.indexOf('\nfunc ', start + 1); return source.slice(start, end < 0 ? undefined : end);
}
function snapshots(source) {
    return Object.fromEntries(['bounded','shared','captured','withCallback','ordered'].flatMap(name=>['Get_Main_'+name,'Call_Main_'+name]).concat('Call_Subject_consume').map(name=>[name,worker(source,name)]));
}
function abi(source) { const match=worker(source,'Call_Subject_consume').match(/^func Call_Subject_consume\([^\n]+/);assert.ok(match,'native consumer worker remains present');return match[0]; }
function variants(original){
 const weighted=original.replace('total * 3','total * 5');
 const seeded=weighted.replace(') 7 (CP.',') 19 (CP.');
 const unknown='-- @inline Subject.step never\n'+seeded.replace('import Prelude','import Prelude\nimport Effect.Console (logShow)\nimport Effect.Unsafe (unsafePerformEffect)').replace('Array.foldl (\\total cp -> total * 5 + fromEnum cp) 19','Array.foldl step 19')+'\nstep :: Int -> CP.CodePoint -> Int\nstep total cp = unsafePerformEffect do\n  logShow (100 + fromEnum cp)\n  pure (total * 5 + fromEnum cp)\n';
 return [{name:'callback',source:weighted,expected:expected(7,5),eligible:true},{name:'seed',source:seeded,expected:expected(19,5),eligible:true},{name:'unknown',source:unknown,expected:expected(19,5,true),eligible:false}];
}

test('codepoint take folds preserve source results, phases and imported descriptor cache transitions', async t => {
    assert.ok(modes.length); const workspace = createWorkspace(); console.log(`Codepoint take fixture workspace: ${workspace}`);
    let completed = 0; t.after(() => { if (completed === modes.length && process.env.GOPURS_TEST_KEEP_WORKSPACE !== '1') rmSync(workspace, { recursive: true, force: true }); });
    const typedBin = join(workspace, 'typed-bin'); mkdirSync(typedBin); symlinkSync(findTypedCompiler(root, process.env.GOPURS_PURS), join(typedBin, 'purs'));
    const environment = { ...process.env, PATH: [typedBin, join(root, 'node_modules/.bin'), process.env.PATH].join(delimiter), GOWORK: 'off', GOPURS_DAEMON: '0', GOPURS_RESIDENT: '0',
        GOPURS_RUST: '0', GOPURS_GO_PACKAGES: '1', GOPURS_GO_PACKAGE_OVERLAY: '1', GOPURS_BUILD_MANIFEST: '1', GOPURS_PREPARE_JOBS: '2', GOPURS_PBO_JOBS: '2', GOPURS_EMIT_JOBS: '2', GOPURS_PIPELINE: '1', GOPURS_FORCE_BUILD: '1' };
    for (const key of ['GOPURS_ALLOC_PROFILE', 'GOPURS_COMPILER_SNAPSHOT', 'PPROF']) delete environment[key];
    for (const [hostIndex, mode] of modes.entries()) await t.test(mode.name, () => {
        const directory = prepareFixture(root, workspace, fixture, hostIndex, corePackages(root)).directory;
        const subject = join(directory, 'src/Subject.purs'); const original = readFileSync(subject, 'utf8'); let sequence = 0;
        const run = (label, command, args, cwd = directory, overrides = {}) => {
            console.log(`${mode.name}: ${label}`); const result = spawnSync(command, args, { cwd, env: { ...environment, GOPURS_JS: mode.GOPURS_JS, ...overrides }, encoding: 'utf8', timeout: 240000, maxBuffer: 16 * 1024 * 1024 });
            writeFileSync(join(directory, 'logs', `${sequence++}-${label}.log`), (result.stdout || '') + (result.stderr || '')); assert.ifError(result.error); assert.equal(result.status, 0, `${label}\n${result.stdout}\n${result.stderr}`); return result.stdout;
        };
        const compiler = process.env.GOPURS_TEST_COMPILER || join(root, 'bin/gopurs');
        const javascript = cwd => parse(run('javascript', 'node', ['--input-type=module', '-e', 'import {main} from "./output/Main/index.js";main();'], cwd));
        run('frontend', 'spago', ['build', '-q', '--offline']); const oracle = javascript(directory); assert.deepEqual(oracle, expected());
        if (process.env.GOPURS_TEST_FRONTEND_ONLY === '1') for (const variant of variants(original)) {
            writeFileSync(subject, variant.source); run(`${variant.name}-frontend-only`, 'spago', ['build', '-q', '--offline']);
            assert.deepEqual(javascript(directory), variant.expected, `${variant.name}: source-only oracle`);
        }
        if (process.env.GOPURS_TEST_FRONTEND_ONLY !== '1') for (const [profileIndex, profile] of profiles.entries()) {
            writeFileSync(subject, original); run(`${profile}-frontend`, 'spago', ['build', '-q', '--offline']);
            run(profile, compiler, [`--${profile}`, '--main', 'Main']);
            const baselineSource = generatedText(join(directory, 'output/purescript')); const before = snapshots(baselineSource);
            writeFileSync(join(directory, 'logs', `${profile}-workers.json`), JSON.stringify(before, null, 2) + '\n');
            const baselineAbi = abi(baselineSource);
            if (profile === 'release') for (const name of ['bounded']) assert.match(before[`Get_Main_${name}`] + before[`Call_Main_${name}`], /codePointTake_/, `${name}: actual imported descriptor must activate`);
            for (const name of ['captured', 'withCallback', 'shared']) assert.doesNotMatch(before[`Get_Main_${name}`] + before[`Call_Main_${name}`], /codePointTake_/, `${name}: unknown or captured mapper remains a barrier`);
            assert.deepEqual(parse(run(`${profile}-go`, 'go', ['run', './main'], join(directory, 'output'))), oracle);
            if (process.env.GOPURS_TEST_SKIP_INCREMENTAL === '1') continue;
            for (const [variantIndex, variant] of variants(original).entries()) {
                writeFileSync(subject, variant.source); run(`${profile}-${variant.name}-frontend`, 'spago', ['build', '-q', '--offline']);
                const changedOracle = javascript(directory); assert.deepEqual(changedOracle, variant.expected, `${variant.name}: independent source oracle`);
                run(`${profile}-${variant.name}-warm`, compiler, [`--${profile}`, '--main', 'Main'], directory, { GOPURS_FORCE_BUILD: '0' });
                const warmSource = generatedText(join(directory, 'output/purescript')); assert.equal(abi(warmSource), baselineAbi, 'descriptor transitions retain native ABI');
                if (profile === 'release') {
                    const mapped = worker(warmSource, 'Get_Main_bounded') + worker(warmSource, 'Call_Main_bounded');
                    if (variant.eligible) assert.match(mapped, /codePointTake_/); else assert.doesNotMatch(mapped, /codePointTake_/);
                }
                const warm = parse(run(`${profile}-${variant.name}-warm-go`, 'go', ['run', './main'], join(directory, 'output'))); assert.deepEqual(warm, changedOracle);
                const clean = prepareFixture(root, workspace, fixture, 100 + hostIndex * 20 + profileIndex * 3 + variantIndex, corePackages(root)).directory;
                writeFileSync(join(clean, 'src/Subject.purs'), variant.source); run(`${profile}-${variant.name}-clean-frontend`, 'spago', ['build', '-q', '--offline'], clean);
                run(`${profile}-${variant.name}-clean`, compiler, [`--${profile}`, '--main', 'Main'], clean);
                const rebuilt = parse(run(`${profile}-${variant.name}-clean-go`, 'go', ['run', './main'], join(clean, 'output'))); assert.deepEqual(rebuilt, warm, 'changed descriptor: incremental equals clean');
                const cleanSource = generatedText(join(clean, 'output/purescript')); assert.equal(abi(cleanSource), baselineAbi);
                assert.deepEqual(snapshots(warmSource), snapshots(cleanSource), 'changed consumer metadata must regenerate the same caller workers as a clean build');
                writeFileSync(join(directory, 'logs', `${profile}-${variant.name}-cache.json`), JSON.stringify({ transition: variant.name, eligible: variant.eligible, same_native_abi: baselineAbi, warm_equals_clean_outputs: true, warm_equals_clean_workers: true, observations: warm.length }, null, 2) + '\n');
            }
        }
        completed++;
    });
});
