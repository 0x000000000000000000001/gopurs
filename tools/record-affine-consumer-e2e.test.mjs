import assert from 'node:assert/strict';
import {spawnSync} from 'node:child_process';
import {mkdirSync,readdirSync,readFileSync,rmSync,symlinkSync,writeFileSync} from 'node:fs';
import {delimiter,join} from 'node:path';
import {fileURLToPath} from 'node:url';
import test from 'node:test';
import {corePackages,createWorkspace,prepareFixture} from './test-workspace.mjs';
import {findTypedCompiler} from './build/native-workspace.mjs';
const root=fileURLToPath(new URL('../',import.meta.url));
const modes=[{name:'js',GOPURS_JS:'1'},{name:'native-parallel',GOPURS_JS:'0'}].filter(x=>!process.env.GOPURS_TEST_MODES||process.env.GOPURS_TEST_MODES.split(',').includes(x.name));
function goFiles(dir){return readdirSync(dir,{withFileTypes:true}).flatMap(e=>{if(e.name==='.gopurs-cache')return [];const p=join(dir,e.name);return e.isDirectory()?goFiles(p):e.name.endsWith('.go')?[p]:[];});}
function worker(dir,name){for(const p of goFiles(dir)){const s=readFileSync(p,'utf8'),a=s.indexOf(`func ${name}(`);if(a>=0){const b=s.indexOf('\nfunc ',a+1);return s.slice(a,b<0?undefined:b);}}return '';}
const parse=s=>s.trim().split('\n').map(x=>JSON.parse(x));
function expected(delta){const out=[];for(const seed of [0,-2147483648,2147483647,-9])for(const count of [0,1,2,7,31]){const sum=(8*seed+28)|0,base=(sum+sum)|0,up=(base+delta*count)|0,theta=(seed+6+delta*count)|0;out.push([up,(base-count)|0,base,(base+count)|0,(up+theta+11)|0,sum]);out.push([up,theta]);}out.push(136+7*delta,360+7*delta,5988+2*delta,5628);return out;}
test('record consumer preserves source values, replay and intramodule proof invalidation',async t=>{
 assert.ok(modes.length);const workspace=createWorkspace();console.log('Record consumer fixture workspace:',workspace);let completed=0;
 t.after(()=>{if(completed===modes.length&&process.env.GOPURS_TEST_KEEP_WORKSPACE!=='1')rmSync(workspace,{recursive:true,force:true});});
 const typedBin=join(workspace,'typed-bin');mkdirSync(typedBin);symlinkSync(findTypedCompiler(root,process.env.GOPURS_PURS),join(typedBin,'purs'));
 const env={...process.env,PATH:[typedBin,join(root,'node_modules/.bin'),process.env.PATH].join(delimiter),GOWORK:'off',GOPURS_DAEMON:'0',GOPURS_RESIDENT:'0',GOPURS_RUST:'0',GOPURS_GO_PACKAGES:'1',GOPURS_GO_PACKAGE_OVERLAY:'1',GOPURS_BUILD_MANIFEST:'1',GOPURS_FORCE_BUILD:'0',GOPURS_PREPARE_JOBS:'2',GOPURS_PBO_JOBS:'2',GOPURS_EMIT_JOBS:'2'};
 for(const key of ['GOPURS_ALLOC_PROFILE','GOPURS_COMPILER_SNAPSHOT','PPROF'])delete env[key];
 for(const[index,mode]of modes.entries())await t.test(mode.name,()=>{
  const {directory}=prepareFixture(root,workspace,join(root,'tools/fixtures/record-affine-consumer/Main.purs'),index,corePackages(root));let sequence=0;
  function run(label,command,args,cwd=directory){console.log(mode.name+': '+label);const r=spawnSync(command,args,{cwd,env:{...env,GOPURS_JS:mode.GOPURS_JS},encoding:'utf8',timeout:180000,maxBuffer:16*1024*1024});writeFileSync(join(directory,'logs',`${sequence++}-${label}.log`),(r.stdout??'')+(r.stderr??''));assert.ifError(r.error);assert.equal(r.status,0,label+'\n'+r.stdout+'\n'+r.stderr);return r.stdout.trim();}
  const compile=profile=>run(profile,process.env.GOPURS_TEST_COMPILER??join(root,'bin/gopurs'),['--'+profile,'--main','Main']);
  const go=label=>parse(run(label,'go',['run','./main'],join(directory,'output')));
  const js=label=>parse(run(label,'node',['--input-type=module','-e','import {main} from "./output/Main/index.js";main();']));
  const output=join(directory,'output');const paths=()=>{const yes=worker(output,'Get_Main_consume')+worker(output,'Call_Main_consume');assert.match(yes,/recordConsumer_/,'canonical source must activate the generic consumer');for(const name of ['withUnknown','withExtraUse','withInterveningEffect'])assert.doesNotMatch(worker(output,'Get_Main_'+name)+worker(output,'Call_Main_'+name),/recordConsumer_/,name+' retains its original boundary');};
  run('frontend','spago',['build','-q','--offline']);const before=js('source-before');assert.deepEqual(before,expected(1));
  if(process.env.GOPURS_TEST_FRONTEND_ONLY==='1'){completed++;return;}
  for(const profile of ['release','dev']){compile(profile);paths();assert.deepEqual(go(profile+'-values'),before);}
  const source=join(directory,'src/Main.purs');writeFileSync(source,readFileSync(source,'utf8').replace('theta = row.theta + 1','theta = row.theta + 2'));
  run('frontend-changed','spago',['build','-q','--offline']);const after=js('source-after');assert.deepEqual(after,expected(2));assert.notDeepEqual(after,before);
  compile('release');const warm=go('warm-values');assert.deepEqual(warm,after);
  for(const p of goFiles(output))rmSync(p);for(const name of ['.gopurs-cache','.gopurs-ownership.json'])rmSync(join(output,name),{recursive:true,force:true});
  compile('release');assert.deepEqual(go('clean-values'),warm);completed++;
 });
});
