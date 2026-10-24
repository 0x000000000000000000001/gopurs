import assert from 'node:assert/strict';
import test from 'node:test';
import { Left } from '../output/Data.Either/index.js';
import * as Aff from '../output/Effect.Aff/index.js';
import { mapPreparationJobs, runPreparationJobs } from '../output/Gopurs.Preparation/index.js';

const runAff=aff=>new Promise((resolve,reject)=>Aff.runAff(result=>()=>result instanceof Left?reject(result.value0):resolve(result.value0))(aff)());

test('preparation is deferred and each execution reruns every job exactly once in result order',async()=>{
  for(const configured of [-2,0,1,3,8,999]){
  for(const count of [1,17,129,1025]){
   const calls=Array(count).fill(0);
  const jobs=calls.map((_,index)=>()=>({index,run:++calls[index]}));
  const input=[...jobs],action=runPreparationJobs(configured)(jobs);
   assert.deepEqual(calls,Array(count).fill(0));
  for(const run of [1,2]){
   assert.deepEqual(await runAff(action),calls.map((_,index)=>({index,run})));
    assert.deepEqual(calls,Array(count).fill(run));
  }
   assert.deepEqual(jobs,input,'partitioning does not consume the caller array');
  }
  }
});

test('empty preparation has the same deferred empty result in sequential and parallel modes',async()=>{
 for(const jobs of [0,1,4,99])assert.deepEqual(await runAff(runPreparationJobs(jobs)([])),[]);
});


test('effectful preparation defers execution, bounds workers, orders outputs and joins failed siblings', async () => {
  for (const configured of [1, 8]) {
    const calls = Array(17).fill(0);
    const action = mapPreparationJobs(configured)(index => Aff.monadEffectAff.liftEffect(() => ({ index, run: ++calls[index] })))(calls.map((_, i) => i));
    assert.deepEqual(calls, Array(17).fill(0));
    for (const run of [1, 2]) assert.deepEqual(await runAff(action), calls.map((_, index) => ({ index, run })));
    let active = 0, started = 0, finished = 0, maximum = 0;
    const task = index => Aff.bracket(Aff.monadEffectAff.liftEffect(() => { started++; maximum = Math.max(maximum, ++active); }))
      (_ => Aff.monadEffectAff.liftEffect(() => { active--; finished++; }))
      (_ => Aff.bindAff.bind(Aff.delay(index === 0 ? 10 : 50))(_ => Aff.monadEffectAff.liftEffect(() => {
        if (index === 0) throw new Error('injected encoder failure');
        return index;
      })));
    await assert.rejects(runAff(mapPreparationJobs(configured)(task)(Array.from({ length: 8 }, (_, i) => i))), /injected encoder failure/);
    assert.equal(active, 0);
    assert.equal(started, finished);
    assert.equal(maximum, configured);
  }
});
