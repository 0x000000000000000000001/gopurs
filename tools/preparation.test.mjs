import assert from 'node:assert/strict';
import test from 'node:test';
import { Left } from '../output/Data.Either/index.js';
import * as Aff from '../output/Effect.Aff/index.js';
import { runPreparationJobs } from '../output/Gopurs.Preparation/index.js';

const runAff=aff=>new Promise((resolve,reject)=>Aff.runAff(result=>()=>result instanceof Left?reject(result.value0):resolve(result.value0))(aff)());

test('preparation is deferred and each execution reruns every job exactly once in result order',async()=>{
 for(const configured of [-2,0,1,3,8,999]){
  const calls=Array(17).fill(0);
  const jobs=calls.map((_,index)=>()=>({index,run:++calls[index]}));
  const input=[...jobs],action=runPreparationJobs(configured)(jobs);
  assert.deepEqual(calls,Array(17).fill(0));
  for(const run of [1,2]){
   assert.deepEqual(await runAff(action),calls.map((_,index)=>({index,run})));
   assert.deepEqual(calls,Array(17).fill(run));
  }
  assert.deepEqual(jobs,input,'partitioning does not consume the caller array');
 }
});

test('empty preparation has the same deferred empty result in sequential and parallel modes',async()=>{
 for(const jobs of [0,1,4,99])assert.deepEqual(await runAff(runPreparationJobs(jobs)([])),[]);
});
