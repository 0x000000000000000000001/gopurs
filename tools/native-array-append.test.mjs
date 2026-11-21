import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';

test('native array append preserves order, values and independent nonempty output', t => {
    const dir = mkdtempSync(join(tmpdir(), 'gopurs-array-append-'));
    t.after(() => rmSync(dir, { recursive: true, force: true }));
    mkdirSync(join(dir, 'gopurs_runtime'));
    writeFileSync(join(dir, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(dir, 'gopurs_runtime/runtime.go'), readFileSync(new URL('../runtime/runtime.go', import.meta.url)));
    writeFileSync(join(dir, 'append.go'), 'package appendtest\n' + readFileSync(new URL('../../gopurs-prelude/src/Data/Semigroup.go', import.meta.url)));
    writeFileSync(join(dir, 'append_test.go'), `package appendtest
import ("testing"; r "gopurs/output/gopurs_runtime")
func TestAppend(t *testing.T) {
 samples:=[][]r.Value{nil,{}, {r.Int(7)}, {r.Str("a"),r.Bool(true),r.Float(-1.5)}}
 for _,xs:=range samples {for _,ys:=range samples {
  beforeX:=append([]r.Value(nil),xs...);beforeY:=append([]r.Value(nil),ys...)
  got:=ConcatArray(xs,ys);if len(got)!=len(xs)+len(ys){t.Fatal("length")}
  for i,v:=range beforeX {if got[i]!=v {t.Fatal("left order/value")}}
  for i,v:=range beforeY {if got[len(xs)+i]!=v {t.Fatal("right order/value")}}
  if len(xs)>0 && len(ys)>0 {got[0]=r.Int(999);got[len(xs)]=r.Int(888)}
  for i,v:=range beforeX {if xs[i]!=v {t.Fatal("mutated left")}}
  for i,v:=range beforeY {if ys[i]!=v {t.Fatal("mutated right")}}
 }}
}
`);
    const run = spawnSync('go', ['test', '.'], { cwd: dir, encoding: 'utf8', env: { ...process.env, GOWORK: 'off' } });
    assert.equal(run.status, 0, run.stdout + run.stderr);
});
