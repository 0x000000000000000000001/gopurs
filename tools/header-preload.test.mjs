import assert from 'node:assert/strict';
import { test } from 'node:test';
import { mkdtempSync, mkdirSync, writeFileSync, rmSync, readFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import * as Aff from '../output/Effect.Aff/index.js';
import { Left } from '../output/Data.Either/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as F from '../output/Data.Foldable/index.js';
import * as U from '../output/Data.Unfoldable/index.js';
import * as Map from '../output/Data.Map/index.js';
import * as Set from '../output/Data.Set/index.js';
import * as List from '../output/Data.List/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import * as Core from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as Sem from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as Wire from '../output/Gopurs.Cache.Wire/index.js';
import * as Header from '../output/Gopurs.Cache.HeaderPreload/index.js';
import * as Cache from '../output/Gopurs.Cache.Modules/index.js';
import { mapPreparationJobs } from '../output/Gopurs.Preparation/index.js';

const run = action => new Promise((resolve, reject) => Aff.runAff(result => () => result instanceof Left ? reject(result.value0) : resolve(result.value0))(action)());
const pure = Aff.applicativeAff.pure, effect = Aff.monadEffectAff.liftEffect, bind = Aff.bindAff.bind;
const table = entries => Map.fromFoldable(ordString)(F.foldableArray)(entries.map(([key, value]) => new Tuple(key, value)));
const set = values => Set.fromFoldable(F.foldableArray)(ordString)(values);
const names = value => Set.toUnfoldable(U.unfoldableArray)(Map.keys(value));
const codec = Wire.codecMap(Sem.ordEvalRef)(Wire.keyOrderEvalRef)(Wire.codecEvalRef)(Wire.codecMap(Sem.ordInlineAccessor)(Wire.keyOrderInlineAccessor)(Wire.codecInlineAccessor)(Wire.codecInlineDirective));
const text = owner => Wire.write(codec)(Map.singleton(new Sem.EvalExtern(new Core.Qualified(new Just(owner), 'value')))(Map.empty))();
const session = { stage: '/private/stage', root: '/private/root', context: 'context', count: 2 };
const entries = ['Alpha', 'Beta'].map(name => ({ name, part: '2', text: text(name) }));
const preload = (jobs, state = session, inventory = ['Alpha', 'Beta'], values = entries) => Header.preloadUsing(jobs)(state)(pure(inventory === null ? Nothing.value : new Just(inventory)))(_ => pure(values));

test('private header reservoir decodes exact owners and only exposes the late safe subset for its own session', async () => {
  for (const jobs of [1, 8]) {
    const result = await run(preload(jobs));
    assert.ok(result instanceof Just);
    assert.deepEqual(names(Header.select(session)(set(['Alpha']))(result.value0).value0), ['Alpha']);
    assert.deepEqual(names(Header.select(session)(set([]))(result.value0).value0), []);
    assert.equal(Header.select(session)(set(['Absent']))(result.value0), Nothing.value);
    for (const key of ['stage', 'root', 'context', 'count']) {
      assert.equal(Header.select({ ...session, [key]: key === 'count' ? 3 : 'different' })(set(['Alpha']))(result.value0), Nothing.value, key);
    }
    for (const [label, inventory, values] of [
      ['absent-inventory', null, entries], ['duplicate-inventory', ['Alpha', 'Alpha'], entries],
      ['missing-inventory', ['Alpha'], entries], ['missing-part', ['Alpha', 'Beta'], entries.slice(0, 1)],
      ['duplicate-part', ['Alpha', 'Beta'], [entries[0], entries[0]]],
      ['outside-owner', ['Alpha', 'Beta'], [entries[0], { ...entries[1], name: 'Outside' }]],
      ['wrong-part', ['Alpha', 'Beta'], [entries[0], { ...entries[1], part: '3' }]],
      ['invalid-codec', ['Alpha', 'Beta'], [entries[0], { ...entries[1], text: '{}' }]],
      ['wrong-owner', ['Alpha', 'Beta'], [entries[0], { ...entries[1], text: entries[0].text }]],
    ]) assert.equal(await run(preload(jobs, session, inventory, values)), Nothing.value, label);
  }
});

test('decoded headers retain prepared, caller, FFI and late headerSafe admission guards', async () => {
  const all = ['Alpha', 'Changed', 'Callers', 'FfiChanged', 'Unsafe'];
  const state = { ...session, count: all.length };
  const value = await run(preload(8, state, all, all.map(name => ({ name, part: '2', text: text(name) }))));
  const keys = name => ({ prepared: name, callerDemands: name + '-callers', optimizerEnvironment: 'prefix' });
  const cache = {
    session: state, lazyContributions: true, headerSafe: set(all.filter(name => name !== 'Unsafe')),
    candidates: table(all.map(name => [name, { ...keys(name), ffiDigest: new Just('no-foreign-v1') }])),
    keys: table(all.map(name => [name, { ...keys(name), prepared: name === 'Changed' ? 'new' : name, callerDemands: name === 'Callers' ? 'new' : name + '-callers' }])),
    ffiDigests: Map.singleton('FfiChanged')('new'), priorFfiDigests: Map.empty,
    priorHeaderDirectives: Map.empty, priorPartTexts: Map.empty,
    restoredPBO: Ref.new(Set.empty)(), invalidations: Ref.new(Map.empty)(),
    priorImplementations: Ref.new(Map.empty)(), priorWorkers: Ref.new(Map.empty)(), backends: Ref.new(Map.empty)(),
  };
  const loaded = await run(Cache.loadHeaderTextsWithPreload(value)(cache));
  assert.deepEqual(names(loaded.priorHeaderDirectives), all.filter(name => name !== 'Unsafe').sort());
  assert.deepEqual(names(loaded.priorPartTexts), [], 'valid reservoir avoids another byte read');
  const restored = await run(Cache.lookupHeaders(8)(loaded)(List.fromFoldable(F.foldableArray)(all.map(name => ({ name, foreign: Map.empty })))));
  assert.deepEqual(names(restored), ['Alpha']);
  for (const field of ['priorImplementations', 'priorWorkers', 'backends']) assert.deepEqual(names(Ref.read(cache[field])()), [], field + ' has no early publication');
});

test('optional failures join all child workers before outer close, including caller failure and cancellation', async () => {
  for (const mode of ['unused', 'error', 'cancel', 'used']) {
    let active = 0, started = 0, finished = 0, closed = 0;
    const task = index => Aff.bracket(effect(() => { started++; active++; }))
      (_ => effect(() => { active--; finished++; }))
      (_ => bind(Aff.delay(index === 0 ? 15 : 45))(_ => effect(() => { if (index === 0 && mode === 'used') throw Error('child failed'); return index; })));
    const preload = bind(mapPreparationJobs(8)(task)([0, 1, 2, 3]))(values => pure(new Just(values)));
    const action = Aff.bracket(pure(null))(_ => effect(() => { closed++; assert.equal(active, 0); assert.equal(started, finished); }))(_ =>
      Header.withOptional(new Just(preload))(pending => mode === 'used' ? pending.value0 : bind(Aff.delay(5))(_ => {
        if (mode === 'error') return effect(() => { throw Error('caller failed'); });
        return mode === 'cancel' ? Aff.delay(200) : pure(Nothing.value);
      })));
    if (mode === 'cancel') {
      const fiber = Aff.launchAff(action)();
      await new Promise(resolve => setTimeout(resolve, 10));
      await run(Aff.killFiber(Error('cancel'))(fiber));
    } else if (mode === 'error') await assert.rejects(run(action), /caller failed/);
    else assert.equal(await run(action), Nothing.value);
    assert.equal(active, 0); assert.equal(started, finished); assert.equal(closed, 1);
    assert.equal(started, 4);
  }
  let used = 0;
  await run(Header.withOptional(Nothing.value)(pending => effect(() => { assert.equal(pending, Nothing.value); used++; })));
  assert.equal(used, 1);
  assert.equal((await run(Header.withOptional(new Just(pure(new Just(42))))(pending => pending.value0))).value0, 42);
});

test('speculative integrity failure is ignored only until the normal required-header read', async t => {
  const directory = mkdtempSync(join(tmpdir(), 'gopurs-header-preload-'));
  const previous = process.env.GOPURS_COMPILER_SNAPSHOT;
  process.env.GOPURS_COMPILER_SNAPSHOT = directory;
  t.after(() => { if (previous === undefined) delete process.env.GOPURS_COMPILER_SNAPSHOT; else process.env.GOPURS_COMPILER_SNAPSHOT = previous; rmSync(directory, { recursive: true, force: true }); });
  mkdirSync(join(directory, 'tools'));
  // This bridge fixture models the existing verified read-parts error contract;
  // native storage integrity itself is covered by the Go and corruption suites.
  writeFileSync(join(directory, 'tools/build-cache-helper'), `#!/usr/bin/env node
const fs=require('node:fs'),path=require('node:path'),r=JSON.parse(fs.readFileSync(0,'utf8'));
fs.appendFileSync(path.join(r.stage,'requests.jsonl'),JSON.stringify(r.owners)+'\\n');
if(r.owners.includes('Beta'))console.log(JSON.stringify({status:'error',reason:'corrupt pinned header'}));
else console.log(JSON.stringify({status:'ready',reason:'',contributions:r.owners.map(name=>({name,part:'2',text:fs.readFileSync(path.join(r.stage,name+'.json'),'utf8')}))}));
`, { mode: 0o755 });
  const stage = join(directory, 'stage'); mkdirSync(stage);
  const current = { ...session, stage };
  writeFileSync(join(stage, 'source-hashes.json'), JSON.stringify(['Alpha', 'Beta'].map(name => ({ name, hash: 'hint' }))));
  writeFileSync(join(stage, 'Alpha.json'), text('Alpha'));
  const base = { session: current, lazyContributions: true, priorHeaderDirectives: Map.empty, priorPartTexts: Map.empty };
  const harmless = await run(Header.withPreload(current)(true)(raw => Header.withDecodedPreload(8)(raw)(pending => bind(pending === Nothing.value ? pure(Nothing.value) : pending.value0)(value => Cache.loadHeaderTextsWithPreload(value)({ ...base, headerSafe: set(['Alpha']) })))));
  assert.deepEqual(names(harmless.priorPartTexts), ['Alpha']);
  await assert.rejects(run(Header.withPreload(current)(true)(raw => Header.withDecodedPreload(8)(raw)(pending => bind(pending === Nothing.value ? pure(Nothing.value) : pending.value0)(value => Cache.loadHeaderTextsWithPreload(value)({ ...base, headerSafe: set(['Alpha', 'Beta']) }))))), /corrupt pinned header/);
  assert.deepEqual(readFileSync(join(stage, 'requests.jsonl'), 'utf8').trim().split('\n').map(JSON.parse), [['Alpha', 'Beta'], ['Alpha'], ['Alpha', 'Beta'], ['Alpha', 'Beta']]);
});


test('raw reads precede the explicit decode boundary and both phases join before close', async () => {
  for (const jobs of [1, 8]) {
    const raw = Header.readUsing(session)(pure(new Just(['Alpha', 'Beta'])))(_ => pure(entries));
    const loaded = await run(Header.withOptional(new Just(raw))(pending => Header.withDecodedPreload(jobs)(pending)(decoded => decoded.value0)));
    const direct = await run(preload(jobs));
    const selected = value => Header.select(session)(set(['Alpha', 'Beta']))(value.value0).value0;
    assert.deepEqual(selected(loaded), selected(direct));
  }
  for (const mode of ['success', 'read-error', 'read-caller-error', 'read-cancel', 'decode-error', 'decode-caller-error', 'decode-cancel']) {
    let boundary = false, decodeCalls = 0, closed = 0;
    const counts = { read: { started: 0, finished: 0, active: 0 }, decode: { started: 0, finished: 0, active: 0 } };
    const child = phase => index => Aff.bracket(effect(() => { counts[phase].started++; counts[phase].active++; }))
      (_ => effect(() => { counts[phase].finished++; counts[phase].active--; }))
      (_ => bind(Aff.delay(index === 0 ? 20 : 45))(_ => effect(() => { if (mode === phase + '-error' && index === 0) throw Error(phase + ' child error'); return index; })));
    const task = phase => bind(mapPreparationJobs(8)(child(phase))([0, 1, 2, 3]))(values => pure(new Just(values)));
    const decoder = _ => bind(effect(() => { assert.equal(boundary, true); decodeCalls++; }))( _ => task('decode'));
    const action = Aff.bracket(pure(null))(_ => effect(() => {
      closed++;
      for (const value of Object.values(counts)) { assert.equal(value.active, 0); assert.equal(value.started, value.finished); }
    }))(_ => Header.withOptional(new Just(task('read')))(raw => bind(Aff.delay(5))(_ => {
      assert.equal(decodeCalls, 0, 'decode must not start during preparation');
      if (mode === 'read-caller-error') return effect(() => { throw Error('caller failed'); });
      if (mode === 'read-cancel') return Aff.delay(200);
      boundary = true;
      return Header.withDecodedUsing(decoder)(raw)(pending => {
        if (mode === 'read-error') return effect(() => assert.equal(pending, Nothing.value));
        if (mode === 'decode-caller-error') return bind(Aff.delay(5))(_ => effect(() => { throw Error('caller failed'); }));
        if (mode === 'decode-cancel') return Aff.delay(200);
        return bind(pending.value0)(value => effect(() => assert.equal(value === Nothing.value, mode === 'decode-error')));
      });
    })));
    if (mode.endsWith('-cancel')) {
      const fiber = Aff.launchAff(action)();
      await new Promise(resolve => setTimeout(resolve, mode === 'read-cancel' ? 10 : 60));
      await run(Aff.killFiber(Error('cancel'))(fiber));
    } else if (mode.endsWith('-caller-error')) await assert.rejects(run(action), /caller failed/);
    else await run(action);
    assert.equal(closed, 1);
    assert.equal(counts.read.started, 4);
    const decoded = mode === 'success' || mode.startsWith('decode-');
    assert.equal(decodeCalls, decoded ? 1 : 0);
    assert.equal(counts.decode.started, decoded ? 4 : 0);
  }
});
