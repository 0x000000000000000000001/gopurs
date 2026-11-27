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
function goFiles(directory){return readdirSync(directory,{withFileTypes:true}).flatMap(e=>{if(e.name==='.gopurs-cache')return [];const p=join(directory,e.name);return e.isDirectory()?goFiles(p):e.name.endsWith('.go')?[p]:[];});}
function worker(directory,name){for(const p of goFiles(directory)){const text=readFileSync(p,'utf8'),a=text.indexOf(`func ${name}(`);if(a>=0){const b=text.indexOf('\nfunc ',a+1);return text.slice(a,b<0?undefined:b);}}return '';}
test('concat folds preserve source Unicode results and imported callback invalidation',async t=>{
 assert.ok(modes.length);const workspace=createWorkspace();console.log('Concat fold fixture workspace:',workspace);let completed=0;
 t.after(()=>{if(completed===modes.length&&process.env.GOPURS_TEST_KEEP_WORKSPACE!=='1')rmSync(workspace,{recursive:true,force:true});});
 const typedBin=join(workspace,'typed-bin');mkdirSync(typedBin);symlinkSync(findTypedCompiler(root,process.env.GOPURS_PURS),join(typedBin,'purs'));
 const env={...process.env,PATH:[typedBin,join(root,'node_modules/.bin'),process.env.PATH].join(delimiter),GOWORK:'off',GOPURS_DAEMON:'0',GOPURS_RESIDENT:'0',GOPURS_RUST:'0',GOPURS_GO_PACKAGES:'1',GOPURS_GO_PACKAGE_OVERLAY:'1',GOPURS_BUILD_MANIFEST:'1',GOPURS_FORCE_BUILD:'0',GOPURS_PREPARE_JOBS:'2',GOPURS_PBO_JOBS:'2',GOPURS_EMIT_JOBS:'2'};
 for(const key of ['GOPURS_ALLOC_PROFILE','GOPURS_COMPILER_SNAPSHOT','PPROF'])delete env[key];
 for(const [index,mode]of modes.entries())await t.test(mode.name,()=>{
  const {directory}=prepareFixture(root,workspace,join(root,'tools/fixtures/string-concat-fold.purs'),index,corePackages(root));const subject=join(directory,'src/Subject.purs');let sequence=0;
  function run(label,command,args,cwd=directory){console.log(mode.name+': '+label);const result=spawnSync(command,args,{cwd,env:{...env,GOPURS_JS:mode.GOPURS_JS},encoding:'utf8',timeout:180000,maxBuffer:16*1024*1024});writeFileSync(join(directory,'logs',`${sequence++}-${label}.log`),(result.stdout??'')+(result.stderr??''));assert.ifError(result.error);assert.equal(result.status,0,label+'\n'+result.stdout+'\n'+result.stderr);return result.stdout.trim();}
  const compile=profile=>run(profile,process.env.GOPURS_TEST_COMPILER??join(root,'bin/gopurs'),['--'+profile,'--main','Main']);
  const go=label=>run(label,'go',['run','./main'],join(directory,'output'));
  const js=label=>run(label,'node',['--input-type=module','-e','import {main} from "./output/Main/index.js";main();']);
  const generated=join(directory,'output');const abi=()=>worker(generated,'Call_Subject_consume').split('\n')[0];
  function assertPaths(){assert.match(worker(generated,'Call_Main_combined'),/concatFold_finish/);for(const name of ['shared','unknown'])assert.doesNotMatch(worker(generated,'Call_Main_'+name),/concatFold_finish/,name+' must preserve the stored/unknown consumer');}
  run('frontend','spago',['build','-q','--offline']);const before=js('source-before');assert.equal(before.split('\n').length,21);let initialAbi;
  for(const profile of ['release','dev']){compile(profile);assertPaths();initialAbi=abi();assert.equal(go(profile+'-values'),before);}
  // Same ABI, new callback descriptor: cached imported callers must re-emit.
  writeFileSync(subject,readFileSync(subject,'utf8').replace('total * 33','total * 37'));
  run('frontend-changed','spago',['build','-q','--offline']);const after=js('source-after');assert.notEqual(after,before);
  compile('release');assertPaths();assert.equal(abi(),initialAbi);const warm=go('warm-values');assert.equal(warm,after);
  for(const p of goFiles(generated))rmSync(p);for(const name of ['.gopurs-cache','.gopurs-ownership.json'])rmSync(join(generated,name),{recursive:true,force:true});
  compile('release');assertPaths();assert.equal(abi(),initialAbi);assert.equal(go('clean-values'),warm);completed++;
 });
});
