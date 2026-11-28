import assert from 'node:assert/strict';
import test from 'node:test';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { empty as emptyMap, insert } from '../output/Data.Map/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as CodeGen from '../output/Gopurs.CodeGen/index.js';
import * as CodePointFold from '../output/Gopurs.CodePointFold/index.js';
import * as G from '../output/Gopurs.GoAst/index.js';
import { SymbolReference, generated } from '../output/Gopurs.GoSymbol/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';
import { runtimeSimdFiles } from '../output/Gopurs.Runtime.Simd/index.js';
import { withReboxFields } from './codegen-metadata.mjs';

const n = syntax => new NeutralExpr(syntax);
const typed = (type, value) => n(new S.Typed(type, value));
const intLit = value => typed(C.Int.value, n(new S.Lit(new C.LitInt(value))));
const local = (level, type) => typed(type, n(new S.Local(new Just('v'), level)));
const qual = (owner, name) => new C.Qualified(new Just(owner), name);
const ref = (owner, name) => n(new S.Var(qual(owner, name)));
const apply = (fn, args, type) => typed(type, n(new S.App(fn, args)));
const addInt = (a, b) => typed(C.Int.value, n(new S.PrimOp(new S.Op2(new S.OpIntNum(S.OpAdd.value), a, b))));
const lambda = (levels, types, ret, body) =>
  typed(new C.Func(types, ret), n(new S.Abs(levels.map(level => new Tuple(new Just('v'), level)), body)));

const strArr = new C.Array(C.String.value);
const intArr = new C.Array(C.Int.value);

const cpCallback = lambda([0, 1], [C.Int.value, C.Int.value], C.Int.value, addInt(local(0, C.Int.value), local(1, C.Int.value)));
const checksumBody = lambda([2], [C.String.value], C.Int.value,
  apply(typed(new C.Func([new C.Func([C.Int.value, C.Int.value], C.Int.value), C.Int.value, intArr], C.Int.value),
    ref('Data.Foldable', 'foldlArray')), [
      cpCallback,
      intLit(0),
      apply(typed(new C.Func([C.String.value], intArr), ref('Data.String.CodePoints', 'toCodePointArray')), [local(2, C.String.value)], intArr),
    ], C.Int.value));

const splitCall = (sep, text) => apply(ref('Data.String.Common', 'split'), [sep, text], strArr);
const foldParts = (seed, partsLevel) => apply(ref('Data.Foldable', 'foldlArray'), [
  lambda([20, 21], [C.Int.value, C.String.value], C.Int.value,
    addInt(local(20, C.Int.value), apply(ref('StringSplitFixture', 'checksum'), [local(21, C.String.value)], C.Int.value))),
  seed,
  local(partsLevel, strArr),
], C.Int.value);

const metadata = () => withReboxFields({
  elidedCtors: emptySet, ctorTypes: emptyMap, pointerAdtPaths: emptyMap, pointerAdtNodes: emptySet,
  pointerAdtLeaves: emptyMap, enumAdts: emptySet, enumCtors: emptySet, globalTypes: emptyMap,
  globalFunctions: emptyMap, classDeclsFields: emptyMap,
});

const moduleOf = bindings => ({
  name: 'StringSplitFixture',
  bindings: [{ recursive: false, bindings }],
  comments: [], imports: emptySet, exports: emptySet, reExports: emptySet,
  dataTypes: emptyMap, dataDecls: [], classDecls: [], foreign: emptyMap, implementations: emptyMap, directives: emptyMap,
});

test('String.split + foldl fuses without allocating intermediate slices and matches strings.Split oracle', t => {
  assert.ok(CodePointFold.describe(checksumBody) instanceof Just);
  const splitWithLength = lambda([10, 11], [C.String.value, C.String.value], C.Int.value,
    typed(C.Int.value, n(new S.Let(new Just('parts'), 12,
      splitCall(local(10, C.String.value), local(11, C.String.value)),
      foldParts(typed(C.Int.value, n(new S.PrimOp(new S.Op1(S.OpArrayLength.value, local(12, strArr))))), 12)))));
  const splitWithSeed = lambda([10, 11, 13], [C.String.value, C.String.value, C.Int.value], C.Int.value,
    typed(C.Int.value, n(new S.Let(new Just('parts'), 12,
      splitCall(local(10, C.String.value), local(11, C.String.value)),
      foldParts(local(13, C.Int.value), 12)))));

  const source = CodeGen.translate(metadata())(moduleOf([
    new Tuple('checksum', checksumBody),
    new Tuple('splitLen', splitWithLength),
    new Tuple('splitSeed', splitWithSeed),
  ]));

  const workerLen = source.slice(source.indexOf('func Call_StringSplitFixture_splitLen('));
  assert.match(workerLen, /stringSplitFold_acc_/);
  assert.doesNotMatch(workerLen, /Data_String_Common_Split|strings\.Split/);

  const directory = mkdtempSync(join(tmpdir(), 'gopurs-string-split-fold-'));
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  mkdirSync(join(directory, 'gopurs_runtime'));
  writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
  writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
  for (const f of runtimeSimdFiles) writeFileSync(join(directory, 'gopurs_runtime', f.name), f.content);
  writeFileSync(join(directory, 'fixture.go'), source.replace('package purescript', 'package main'));
  writeFileSync(join(directory, 'main.go'), `package main
import (
  "fmt"
  "strings"
  "testing"
  r "gopurs/output/gopurs_runtime"
)
var sink int64
func oracleChecksum(s string) int64 {
  acc := int64(0)
  for off := 0; off < len(s); {
    cp, w := r.DecodeStringCodePoint(s[off:])
    acc = r.IntAdd(acc, int64(cp))
    off += w
  }
  return acc
}
func oracle(sep, text string, useLen bool, seed int64) int64 {
  parts := strings.Split(text, sep)
  acc := seed
  if useLen { acc = int64(len(parts)) }
  for _, p := range parts {
    acc = r.IntAdd(acc, oracleChecksum(p))
  }
  return acc
}
func main() {
  seps := []string{"", ",", "::", "—", "\\xff", "abc"}
  texts := []string{
    "", ",", ",,", ",a", "a,", "abc,def,ghi",
    "a😀é,b,c😀", "a::b::::c", "x—y—z",
    "\\xc2,abc", "\\xed\\xa0,def", "\\xed\\xa0\\x80\\xed\\xb0\\x80,z",
    "a\\xffb\\xff\\xffc", "abcabc",
  }
  seeds := []int64{0, 1, -1, 2147483647, -2147483648, 1 << 40}
  checks := 0
  for _, sep := range seps {
    for _, text := range texts {
      got := Call_StringSplitFixture_splitLen(sep, text)
      want := oracle(sep, text, true, 0)
      if got != want { panic(fmt.Sprintf("len sep=%q text=%q got=%d want=%d", sep, text, got, want)) }
      checks++
      for _, seed := range seeds {
        gotS := Call_StringSplitFixture_splitSeed(sep, text, seed)
        wantS := oracle(sep, text, false, seed)
        if gotS != wantS { panic(fmt.Sprintf("seed=%d sep=%q text=%q got=%d want=%d", seed, sep, text, gotS, wantS)) }
        checks++
      }
    }
  }
  big := strings.Repeat("abc,", 500) + "abc"
  allocs := testing.AllocsPerRun(50, func() { sink = Call_StringSplitFixture_splitLen(",", big) })
  if allocs != 0 { panic(fmt.Sprintf("expected 0 allocs, got %v", allocs)) }
  fmt.Printf("%d string split fold controls passed (allocs=%v)\\n", checks, allocs)
}
`);
  const result = spawnSync('go', ['run', '.'], { cwd: directory, encoding: 'utf8', env: { ...process.env, GOWORK: 'off' } });
  assert.equal(result.status, 0, result.stdout + result.stderr);
  assert.match(result.stdout, /string split fold controls passed \(allocs=0\)/);
});
