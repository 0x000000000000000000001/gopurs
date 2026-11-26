import assert from 'node:assert/strict';
import {spawnSync} from 'node:child_process';
import {mkdirSync,readFileSync,readdirSync,rmSync,symlinkSync,writeFileSync} from 'node:fs';
import {delimiter,join} from 'node:path';
import {fileURLToPath} from 'node:url';
import test from 'node:test';
import {corePackages,createWorkspace,prepareFixture} from './test-workspace.mjs';
import {findTypedCompiler} from './build/native-workspace.mjs';

const root=fileURLToPath(new URL('../',import.meta.url));
const modes=[{name:'js',GOPURS_JS:'1'},{name:'native-parallel',GOPURS_JS:'0'}].filter(x=>!process.env.GOPURS_TEST_MODES||process.env.GOPURS_TEST_MODES.split(',').includes(x.name));
const cases=[[],[0],[97,98,99],[233,8364,65535],[128512,1114111],[55296],[56320],[55296,56320],[56320,55296],[55296,97,56320],[55296,55297,56320],[128512,55296,56320,0,233],Array(3000).fill(1114111)];
const points=input=>Array.from(input,cp=>cp.codePointAt(0));
const checksum=(input,seed=0)=>points(input).reduce((total,cp)=>(total+cp)|0,seed);
const expected=[];
for(const values of cases){
 const input=values.map(cp=>String.fromCodePoint(cp)).join('');
 expected.push([checksum(input),checksum(input),checksum(input),checksum(input),checksum(input,2147483647),points(input).reduce((total,cp)=>((((total*3)|0)+cp)|0),1),points(input).reduce((total,cp)=>(total+cp+3)|0,0),checksum(input)]);
 expected.push(points(input));
 expected.push([0,1,2,10000].map(n=>checksum(Array.from(input).slice(0,n).join(''))));
}
expected.push(301,[97,98,99],'callback failure',[97,98],2147483647,[],301,[1,2],'seed failure',[1]);
const parse=s=>s.trim().split('\n').map(line=>JSON.parse(line));
function generatedText(directory){return readdirSync(directory,{withFileTypes:true}).filter(x=>x.name!=='.gopurs-cache').map(entry=>{const p=join(directory,entry.name);return entry.isDirectory()?generatedText(p):entry.name.endsWith('.go')?readFileSync(p,'utf8'):'';}).join('\n');}
function worker(source,name){const start=source.indexOf(`func ${name}(`);if(start<0)return '';const end=source.indexOf('\nfunc ',start+1);return source.slice(start,end<0?undefined:end);}

test('code-point folds preserve UTF-16 pairing, lone surrogates, overflow, replay and unknown callback effects',async t=>{
 assert.ok(modes.length);const workspace=createWorkspace();console.log(`String code-point fold fixture workspace: ${workspace}`);let complete=0;t.after(()=>{if(complete===modes.length&&process.env.GOPURS_TEST_KEEP_WORKSPACE!=='1')rmSync(workspace,{recursive:true,force:true});});
 const typedBin=join(workspace,'typed-bin');mkdirSync(typedBin);symlinkSync(findTypedCompiler(root,process.env.GOPURS_PURS),join(typedBin,'purs'));
 const env={...process.env,PATH:[typedBin,join(root,'node_modules/.bin'),process.env.PATH].join(delimiter),GOWORK:'off',GOPURS_DAEMON:'0',GOPURS_RESIDENT:'0',GOPURS_RUST:'0',GOPURS_GO_PACKAGES:'1',GOPURS_GO_PACKAGE_OVERLAY:'1',GOPURS_BUILD_MANIFEST:'1',GOPURS_PREPARE_JOBS:'2',GOPURS_PBO_JOBS:'2',GOPURS_EMIT_JOBS:'2',GOPURS_PIPELINE:'1',GOPURS_FORCE_BUILD:'1'};
 for(const k of ['GOPURS_ALLOC_PROFILE','GOPURS_COMPILER_SNAPSHOT','PPROF'])delete env[k];
 for(const [i,mode] of modes.entries())await t.test(mode.name,()=>{
  const {directory}=prepareFixture(root,workspace,join(root,'tools/fixtures/string-codepoint-fold.purs'),i,corePackages(root));let sequence=0;
  const run=(label,command,args,cwd=directory,overrides={})=>{console.log(`${mode.name}: ${label}`);const r=spawnSync(command,args,{cwd,env:{...env,GOPURS_JS:mode.GOPURS_JS,...overrides},encoding:'utf8',timeout:120000,maxBuffer:16*1024*1024});writeFileSync(join(directory,'logs',`${sequence++}-${label}.log`),(r.stdout??'')+(r.stderr??''));assert.ifError(r.error);assert.equal(r.status,0,`${label}\n${r.stdout}\n${r.stderr}`);return r.stdout;};
  run('frontend','spago',['build','-q','--offline']);const oracle=parse(run('javascript','node',['--input-type=module','-e','import {main} from "./output/Main/index.js";main();']));assert.deepEqual(oracle,expected,'frontend JS agrees with an independent UTF-16/code-point and Int32 oracle');
  if(process.env.GOPURS_TEST_FRONTEND_ONLY!=='1')for(const profile of ['release','dev']){
   run(profile,process.env.GOPURS_TEST_COMPILER??join(root,'bin/gopurs'),[`--${profile}`,'--main','Main']);
   const source=generatedText(join(directory,'output','purescript'));
   const workers=Object.fromEntries(['Call_Main_checksum','Get_Main_checksum','Call_Main_explicit','Call_Main_seeded','Call_Main_weighted','Call_Main_captured','Call_Main_ordered','Call_Main_withCallback','Call_Subject_checksum','Get_Subject_checksum'].map(name=>[name,worker(source,name)]));
   writeFileSync(join(directory,'logs',`${profile}-workers.json`),JSON.stringify(workers,null,2)+'\n');
   if(profile==='release'){
    for(const name of ['Call_Main_checksum','Call_Main_explicit','Call_Main_seeded','Call_Main_weighted','Call_Main_ordered','Call_Subject_checksum'])assert.match(worker(source,name),/DecodeStringCodePoint/,`${name}: the optimized traversal must be active`);
    // This capture lacks a local scalar annotation in the exported IR. The
    // conservative proof does not infer it from the enclosing worker signature.
    for(const name of ['Call_Main_captured','Call_Main_withCallback']){
     assert.doesNotMatch(worker(source,name),/DecodeStringCodePoint/,`${name}: unproved callbacks retain the original array path`);
     assert.match(worker(source,name),/toCodePointArray/,`${name}: the fallback is present, not an absent worker`);
    }
    assert.match(worker(source,'Call_Subject_checksum'),/func Call_Subject_checksum\([^)]*string\) int64/,'imported checksum publishes its native string worker');
   }
   const actual=parse(run(`${profile}-go`,'go',['run','./main'],join(directory,'output')));assert.deepEqual(actual,oracle,`${profile}: code points, complete results, overflow, replay, callback effects and exception priority agree with JavaScript`);
   if(profile==='release'&&process.env.GOPURS_TEST_SKIP_INCREMENTAL!=='1'){
    const subject=join(directory,'src','Subject.purs');const original=readFileSync(subject,'utf8');
    const changed='-- @inline Subject.increment never\n'+original.replace('(\\total cp -> total + fromEnum cp)','increment')+'\nincrement :: Int -> CP.CodePoint -> Int\nincrement total cp = total + fromEnum cp + 1\n';
    writeFileSync(subject,changed);run('incremental-frontend','spago',['build','-q','--offline']);
    const changedExpected=structuredClone(expected);for(let j=0;j<cases.length;j++)changedExpected[j*3][7]=(changedExpected[j*3][7]+changedExpected[j*3+1].length)|0;
    const changedOracle=parse(run('incremental-javascript','node',['--input-type=module','-e','import {main} from "./output/Main/index.js";main();']));assert.deepEqual(changedOracle,changedExpected);
    run('incremental-release',process.env.GOPURS_TEST_COMPILER??join(root,'bin/gopurs'),['--release','--main','Main'],directory,{GOPURS_FORCE_BUILD:'0'});
    const changedSource=generatedText(join(directory,'output','purescript'));assert.doesNotMatch(worker(changedSource,'Call_Subject_checksum'),/DecodeStringCodePoint/,'unknown callback must remove imported traversal specialization');
    const incremental=parse(run('incremental-go','go',['run','./main'],join(directory,'output')));assert.deepEqual(incremental,changedOracle,'cached imported calls must observe the changed callback and signature');
    const clean=prepareFixture(root,workspace,join(root,'tools/fixtures/string-codepoint-fold.purs'),100+i,corePackages(root)).directory;writeFileSync(join(clean,'src','Subject.purs'),changed);
    run('clean-frontend','spago',['build','-q','--offline'],clean);run('clean-release',process.env.GOPURS_TEST_COMPILER??join(root,'bin/gopurs'),['--release','--main','Main'],clean);
    const rebuilt=parse(run('clean-go','go',['run','./main'],join(clean,'output')));assert.deepEqual(rebuilt,incremental,'incremental worker eligibility change equals a fresh build');
    writeFileSync(join(directory,'logs','incremental-validation.json'),JSON.stringify({transition:'closed primitive callback / native string worker -> unknown callback / original composition',clean_equals_incremental:true,original_observations:expected.length,changed_observations:changedOracle.length,changed_imported_results:true},null,2)+'\n');
    writeFileSync(subject,original);run('restore-frontend','spago',['build','-q','--offline']);
   }
  }
  complete++;
 });
});
