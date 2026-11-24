import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, readFileSync, readdirSync, rmSync, symlinkSync, writeFileSync } from 'node:fs';
import { delimiter, join } from 'node:path';
import { fileURLToPath } from 'node:url';
import test from 'node:test';
import { corePackages, createWorkspace, prepareFixture } from './test-workspace.mjs';
import { findTypedCompiler } from './build/native-workspace.mjs';

const root=fileURLToPath(new URL('../',import.meta.url));
const modes=[{name:'js',GOPURS_JS:'1'},{name:'native-parallel',GOPURS_JS:'0'}]
  .filter(mode=>!process.env.GOPURS_TEST_MODES||process.env.GOPURS_TEST_MODES.split(',').includes(mode.name));
const values=[0,1,2,3,8,12];
const add=(a,b)=>(a+b)|0;
const fib=n=>n<2?n:add(fib(n-1),fib(n-2));
const weighted=(n,seed)=>n<=0?seed:add(add(n,seed),weighted(n-1,add(seed,7)));
const permuted=(n,a,b)=>n===0?a:add(b,permuted(n-1,b,add(a,1)));
const expected=[values.map(fib)];
for(const seed of [0,7,2147483647,-2147483648,2147483648]) {
  expected.push(values.map(n=>weighted(n,seed)),values.map(n=>permuted(n,seed,-17)));
  for(const n of [0,1,3]) {
    let value=seed;
    for(let i=0;i<n;i++)value=add(2147483647,value);
    expected.push(value,Array.from({length:n},(_,i)=>[100+n-i,200+n-i-1]).flat().concat(300));
  }
}
const parse=text=>text.trim().split('\n').map(JSON.parse);

test('tail addition preserves renamed recurrences, wide bases, overflow and opaque operand order',async t=>{
  assert.ok(modes.length>0);
  const workspace=createWorkspace();
  console.log(`Tail addition fixture workspace: ${workspace}`);
  let completed=0;
  t.after(()=>{if(completed===modes.length&&process.env.GOPURS_TEST_KEEP_WORKSPACE!=='1')rmSync(workspace,{recursive:true,force:true});});
  const typedBin=join(workspace,'typed-bin');mkdirSync(typedBin);
  symlinkSync(findTypedCompiler(root,process.env.GOPURS_PURS),join(typedBin,'purs'));
  const env={...process.env,PATH:[typedBin,join(root,'node_modules/.bin'),process.env.PATH].join(delimiter),
    GOWORK:'off',GOPURS_DAEMON:'0',GOPURS_RESIDENT:'0',GOPURS_RUST:'0',GOPURS_FORCE_BUILD:'1',
    GOPURS_GO_PACKAGES:'1',GOPURS_GO_PACKAGE_OVERLAY:'1',GOPURS_BUILD_MANIFEST:'1',
    GOPURS_PREPARE_JOBS:'2',GOPURS_PBO_JOBS:'2',GOPURS_EMIT_JOBS:'2',GOPURS_PIPELINE:'1'};
  for(const key of ['GOPURS_ALLOC_PROFILE','GOPURS_COMPILER_SNAPSHOT','PPROF'])delete env[key];
  const packages=corePackages(root);
  for(const [index,mode] of modes.entries())await t.test(mode.name,()=>{
    const {directory}=prepareFixture(root,workspace,join(root,'tools/fixtures/tail-add-loop.purs'),index,packages);
    let sequence=0;
    const run=(label,command,args,cwd=directory)=>{
      console.log(`${mode.name}: ${label}`);
      const result=spawnSync(command,args,{cwd,env:{...env,GOPURS_JS:mode.GOPURS_JS},encoding:'utf8',timeout:120000,maxBuffer:16*1024*1024});
      writeFileSync(join(directory,'logs',`${sequence++}-${label}.log`),(result.stdout??'')+(result.stderr??''));
      assert.ifError(result.error);assert.equal(result.status,0,`${label}\n${result.stdout}\n${result.stderr}`);
      return result.stdout;
    };
    run('frontend','spago',['build','-q','--offline']);
    assert.deepEqual(parse(run('javascript','node',['--input-type=module','-e','import { main } from "./output/Main/index.js"; main();'])),expected);
    for(const profile of ['release','dev']) {
      run(profile,process.env.GOPURS_TEST_COMPILER??join(root,'bin/gopurs'),[`--${profile}`,'--main','Main']);
      const emitted=readdirSync(join(directory,'output'),{recursive:true})
        .filter(name=>name.endsWith('.go')&&!name.includes('.gopurs-cache'))
        .map(name=>readFileSync(join(directory,'output',name),'utf8')).join('\n');
      if(profile==='release'&&process.env.GOPURS_TEST_EXPECT_TAIL_ADD!=='0')
        assert.match(emitted,/__gopurs_tail_add_\d+_used/,`${profile}: native tail addition must be emitted`);
      assert.deepEqual(parse(run(`${profile}-go`,'go',['run','./main'],join(directory,'output'))),expected,`${profile}: results and opaque operand order`);
    }
    completed++;
  });
});
