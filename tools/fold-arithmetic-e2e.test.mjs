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
const euclidean=(a,b)=>b===0?0:((a%Math.abs(b))+Math.abs(b))%Math.abs(b);
const product=(a,b)=>(a*b)|0;
function weighted(values){let result=0;for(let i=0;i<values.length;i++)result=(result+product(1+i%2,values[i]))|0;return result;}
const minimum=-2147483648,maximum=2147483647,unsigned=4294967295,quotient=2147483648;
const cases=[ [[],[]],[[1],[]],[[1,2,3,4,5],[-1,0,2]],[[-7,0,-7,4,-2],[1,1]],[[minimum,maximum,unsigned,quotient,0,-1],[unsigned,minimum]],[Array.from({length:35},(_,i)=>i-17),[7,6,5,4,3,2,1]] ];
const expected=[];
for(const [left,right] of cases){const mapped=left.map(x=>(x+1)|0);const scores=[weighted(mapped),weighted([...left,...right]),weighted(mapped),left.length?-1:0,left.length?-1:0,weighted(left.filter(x=>euclidean(x,2)===0)),weighted([...left].sort((a,b)=>a-b))];expected.push(scores,scores,[left,right]);}
for(const v of [minimum,minimum+1,-65537,-9,-2,-1,0,1,2,3,7,65535,maximum,unsigned,quotient])expected.push([euclidean(v,2),euclidean(v,4),euclidean(v,8),euclidean(v,-2),0,product(v,v),product(1+euclidean(v,2),v),product(v,1+euclidean(v,2)),product(euclidean(v,2),v),product(euclidean(v,4)-2,v),product(v,euclidean(v,4)-2)]);
expected.push(-2,[1,2],'left operand failure',[1,2,3]);
const parse=s=>s.trim().split('\n').map(line=>JSON.parse(line));
function emittedWorker(directory,name){for(const entry of readdirSync(directory,{withFileTypes:true})){if(entry.name==='.gopurs-cache')continue;const p=join(directory,entry.name);const s=entry.isDirectory()?emittedWorker(p,name):entry.name.endsWith('.go')?readFileSync(p,'utf8'):'';const start=s.indexOf(`func ${name}(`);if(start>=0){const end=s.indexOf('\nfunc ',start+1);return s.slice(start,end<0?undefined:end);}}return '';}
test('fold arithmetic preserves Number multiplication, Euclidean modulo, noncanonical Ints and seven array consumers',async t=>{
 assert.ok(modes.length);const workspace=createWorkspace();console.log(`Fold arithmetic fixture workspace: ${workspace}`);let complete=0;t.after(()=>{if(complete===modes.length&&process.env.GOPURS_TEST_KEEP_WORKSPACE!=='1')rmSync(workspace,{recursive:true,force:true});});
 const typedBin=join(workspace,'typed-bin');mkdirSync(typedBin);symlinkSync(findTypedCompiler(root,process.env.GOPURS_PURS),join(typedBin,'purs'));
 const env={...process.env,PATH:[typedBin,join(root,'node_modules/.bin'),process.env.PATH].join(delimiter),GOWORK:'off',GOPURS_DAEMON:'0',GOPURS_RESIDENT:'0',GOPURS_RUST:'0',GOPURS_GO_PACKAGES:'1',GOPURS_GO_PACKAGE_OVERLAY:'1',GOPURS_BUILD_MANIFEST:'1',GOPURS_PREPARE_JOBS:'2',GOPURS_PBO_JOBS:'2',GOPURS_EMIT_JOBS:'2',GOPURS_PIPELINE:'1',GOPURS_FORCE_BUILD:'1'};
 for(const k of ['GOPURS_ALLOC_PROFILE','GOPURS_COMPILER_SNAPSHOT','PPROF'])delete env[k];
 for(const [i,mode] of modes.entries())await t.test(mode.name,()=>{
  const {directory}=prepareFixture(root,workspace,join(root,'tools/fixtures/fold-arithmetic.purs'),i,corePackages(root));let sequence=0;
  const run=(label,command,args,cwd=directory)=>{console.log(`${mode.name}: ${label}`);const r=spawnSync(command,args,{cwd,env:{...env,GOPURS_JS:mode.GOPURS_JS},encoding:'utf8',timeout:120000,maxBuffer:16*1024*1024});writeFileSync(join(directory,'logs',`${sequence++}-${label}.log`),(r.stdout??'')+(r.stderr??''));assert.ifError(r.error);assert.equal(r.status,0,`${label}\n${r.stdout}\n${r.stderr}`);return r.stdout;};
  run('frontend','spago',['build','-q','--offline']);const oracle=parse(run('javascript','node',['--input-type=module','-e','import {main} from "./output/Main/index.js";main();']));assert.deepEqual(oracle,expected,'frontend JS agrees with independent Number/ToInt32 and Euclidean reference');
  if(process.env.GOPURS_TEST_FRONTEND_ONLY!=='1')for(const profile of ['release','dev']){
   run(profile,process.env.GOPURS_TEST_COMPILER??join(root,'bin/gopurs'),[`--${profile}`,'--main','Main']);
   const worker=emittedWorker(join(directory,'output'),'Call_Main_weighted');assert.match(worker,/\[\]int64/,'the canonical fold must retain its native array worker');if(profile==='release'){assert.match(worker,/guardedLeft/,'the guarded product must be active in the actual fold');assert.match(worker,/guardedRight/);assert.match(worker,/IntMul/,'the out-of-range path must retain historical Number multiplication');}writeFileSync(join(directory,'logs',`${profile}-weighted.go.txt`),worker);
   const actual=parse(run(`${profile}-go`,'go',['run','./main'],join(directory,'output')));assert.deepEqual(actual,oracle,`${profile}: all complete results, replay, input arrays and operand traces agree with JavaScript`);
  }
  complete++;
 });
});
