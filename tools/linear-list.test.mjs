import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { withReboxFields } from './codegen-metadata.mjs';
import { empty as emptyMap, insert } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet, singleton } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as CodeGen from '../output/Gopurs.CodeGen/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import * as LinearList from '../output/Gopurs.LinearList/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';

const map = entries => entries.reduce((acc, [k, v]) => insert(ordString)(k)(v)(acc), emptyMap);
const expr = syntax => new NeutralExpr(syntax);
const typed = (ty, value) => expr(new S.Typed(ty, value));
const local = (name, level, ty) => typed(ty, expr(new S.Local(new Just(name), level)));
const integer = n => typed(C.Int.value, expr(new S.Lit(new C.LitInt(n))));
const qualified = (module, name) => new C.Qualified(new Just(module), name);
const globalVar = (module, name, ty) => typed(ty, expr(new S.Var(qualified(module, name))));
const apply = (ty, fn, ...args) => typed(ty, expr(new S.App(fn, args)));
const lambda = (ty, args, body) => typed(ty,
    expr(new S.Abs(args.map(([name, level]) => new Tuple(new Just(name), level)), body)));
const bin = (ty, op, l, r) => typed(ty, expr(new S.PrimOp(new S.Op2(op, l, r))));
const un = (ty, op, x) => typed(ty, expr(new S.PrimOp(new S.Op1(op, x))));

function buildModule({ module = 'ChainProbe', adt = 'Chain', empty = 'Done', node = 'Link', poly = true } = {}) {
    const intType = C.Int.value;
    const boolType = C.Boolean.value;
    const listType = new C.ADT(`${module}.${adt}`, [module, adt], poly ? [intType] : []);
    const typeVarA = new C.TypeVar('a');
    const declAdt = new C.ADT(`${module}.${adt}`, [module, adt], poly ? [typeVarA] : []);
    const ctorNames = Go.constructorNames(module)(node);
    const leafNames = Go.constructorNames(module)(empty);
    const ctor = (name, ...vals) => typed(listType, expr(new S.CtorSaturated(
        qualified(module, name), C.SumType.value, adt, name,
        vals.map((v, i) => new Tuple(`value${i}`, v)))));
    const field = (v, idx) => typed(idx === 0 ? intType : listType, expr(new S.Accessor(
        v, new S.GetCtorField(qualified(module, node), C.SumType.value, adt, node, `value${idx}`, idx))));
    const tag = (name, v) => typed(boolType, expr(new S.PrimOp(new S.Op1(new S.OpIsTag(qualified(module, name)), v))));
    const cases = (retTy, scrut, onNil, onCons) => typed(retTy, expr(new S.Branch([
        new S.Pair(tag(empty, scrut), onNil),
        new S.Pair(tag(node, scrut), onCons),
    ], expr(new S.Fail('exhausted')))));

    const predType = new C.Func([intType], boolType);
    const filterType = new C.Func([predType, listType], listType);
    const transformType = new C.Func([listType], listType);
    const rangeType = new C.Func([intType, intType], listType);
    const rangeLoopType = new C.Func([intType, listType], listType);
    const foldLoopType = new C.Func([listType, intType], intType);

    // Direct recursive filter: keep p Done = Done; keep p (Link x xs) = if p x then Link x (keep p xs) else keep p xs
    const fp = local('p', 0, predType);
    const fl = local('lst', 1, listType);
    const filterBody = cases(listType, fl, ctor(empty),
        typed(listType, expr(new S.Branch([
            new S.Pair(apply(boolType, fp, field(fl, 0)),
                ctor(node, field(fl, 0), apply(listType, globalVar(module, 'keep', filterType), fp, field(fl, 1)))),
        ], apply(listType, globalVar(module, 'keep', filterType), fp, field(fl, 1))))));
    const keepFn = lambda(filterType, [['p', 0], ['lst', 1]], filterBody);

    // Filtered TRMC transformer: prune Done = Done; prune (Link h t) = Link h (prune (keep (\x -> not (x `mod` h == 0)) t))
    const sl = local('xs', 0, listType);
    const px = local('x', 1, intType);
    const modZero = bin(boolType, new S.OpIntOrd(S.OpEq.value),
        bin(intType, new S.OpIntNum(S.OpMod.value), px, field(sl, 0)), integer(0));
    const predLam = lambda(predType, [['x', 1]], un(boolType, S.OpBooleanNot.value, modZero));
    const pruneFn = lambda(transformType, [['xs', 0]],
        cases(listType, sl, ctor(empty),
            ctor(node, field(sl, 0),
                apply(listType, globalVar(module, 'prune', transformType),
                    apply(listType, globalVar(module, 'keep', filterType), predLam, field(sl, 1))))));

    // Direct TRMC transformer: bump Done = Done; bump (Link h t) = Link (h * 2 + 1) (bump t)
    const bl = local('ys', 0, listType);
    const bumpedHead = bin(intType, new S.OpIntNum(S.OpAdd.value),
        bin(intType, new S.OpIntNum(S.OpMultiply.value), field(bl, 0), integer(2)), integer(1));
    const bumpFn = lambda(transformType, [['ys', 0]],
        cases(listType, bl, ctor(empty),
            ctor(node, bumpedHead, apply(listType, globalVar(module, 'bump', transformType), field(bl, 1)))));

    // Top-level descending range builder: span lo hi = let go c acc = if c < lo then acc else go (c - 1) (Link c acc) in go hi Done
    const rLo = local('lo', 0, intType);
    const rHi = local('hi', 1, intType);
    const rGo = local('go', 2, rangeLoopType);
    const rCurr = local('curr', 3, intType);
    const rAcc = local('acc', 4, listType);
    const rLoop = lambda(rangeLoopType, [['curr', 3], ['acc', 4]],
        typed(listType, expr(new S.Branch([
            new S.Pair(bin(boolType, new S.OpIntOrd(S.OpLt.value), rCurr, rLo), rAcc),
        ], apply(listType, rGo, bin(intType, new S.OpIntNum(S.OpSubtract.value), rCurr, integer(1)), ctor(node, rCurr, rAcc))))));
    const spanFn = lambda(rangeType, [['lo', 0], ['hi', 1]],
        typed(listType, expr(new S.LetRec(2, [new Tuple('go', rLoop)],
            apply(listType, rGo, rHi, ctor(empty))))));

    // Top-level additive folder: sumChain xs = let go Done a = a; go (Link h t) a = go t (a + h) in go xs 0
    const sIn = local('in', 0, listType);
    const sGo = local('go', 1, foldLoopType);
    const sCur = local('cur', 2, listType);
    const sAcc = local('acc', 3, intType);
    const sLoop = lambda(foldLoopType, [['cur', 2], ['acc', 3]],
        cases(intType, sCur, sAcc,
            apply(intType, sGo, field(sCur, 1), bin(intType, new S.OpIntNum(S.OpAdd.value), sAcc, field(sCur, 0)))));
    const sumChainFn = lambda(new C.Func([listType], intType), [['in', 0]],
        typed(intType, expr(new S.LetRec(1, [new Tuple('go', sLoop)],
            apply(intType, sGo, sIn, integer(0))))));

    // Top-level non-additive folder: hashChain seed xs = go xs seed where step = acc * 3 + h
    const hSeed = local('seed', 0, intType);
    const hIn = local('in', 1, listType);
    const hGo = local('go', 2, foldLoopType);
    const hCur = local('cur', 3, listType);
    const hAcc = local('acc', 4, intType);
    const hStep = bin(intType, new S.OpIntNum(S.OpAdd.value),
        bin(intType, new S.OpIntNum(S.OpMultiply.value), hAcc, integer(3)), field(hCur, 0));
    const hLoop = lambda(foldLoopType, [['cur', 3], ['acc', 4]],
        cases(intType, hCur, hAcc, apply(intType, hGo, field(hCur, 1), hStep)));
    const hashChainFn = lambda(new C.Func([intType, listType], intType), [['seed', 0], ['in', 1]],
        typed(intType, expr(new S.LetRec(2, [new Tuple('go', hLoop)],
            apply(intType, hGo, hIn, hSeed)))));

    // Caller 1 (aliased): pruneShared xs = prune xs  --> MUST NOT rewrite to linear worker!
    const shIn = local('shared', 0, listType);
    const pruneSharedFn = lambda(transformType, [['shared', 0]],
        apply(listType, globalVar(module, 'prune', transformType), shIn));

    // Caller 2 (fresh producer): pruneFresh lo hi = prune (span lo hi) --> MUST rewrite to linear worker!
    const fLo = local('lo', 0, intType);
    const fHi = local('hi', 1, intType);
    const pruneFreshFn = lambda(rangeType, [['lo', 0], ['hi', 1]],
        apply(listType, globalVar(module, 'prune', transformType),
            apply(listType, globalVar(module, 'span', rangeType), fLo, fHi)));

    // Caller 3 (fresh composed): bumpFresh lo hi = bump (span lo hi) --> MUST rewrite to linear worker!
    const bLo = local('lo', 0, intType);
    const bHi = local('hi', 1, intType);
    const bumpFreshFn = lambda(rangeType, [['lo', 0], ['hi', 1]],
        apply(listType, globalVar(module, 'bump', transformType),
            apply(listType, globalVar(module, 'span', rangeType), bLo, bHi)));

    // Caller 4 (fused top-level fold + prune + span): sumPruned lo hi = sumChain (prune (span lo hi))
    const pLo = local('lo', 0, intType);
    const pHi = local('hi', 1, intType);
    const sumPrunedFn = lambda(new C.Func([intType, intType], intType), [['lo', 0], ['hi', 1]],
        apply(intType, globalVar(module, 'sumChain', new C.Func([listType], intType)),
            apply(listType, globalVar(module, 'prune', transformType),
                apply(listType, globalVar(module, 'span', rangeType), pLo, pHi))));

    // Caller 5 (fused local LetRec fold + prune + local LetRec range):
    // localFused lo hi seed = letrec goF; letrec goR; goF (prune (goR hi Done)) seed
    const lLo = local('lo', 0, intType);
    const lHi = local('hi', 1, intType);
    const lSeed = local('seed', 2, intType);
    const lFold = local('goF', 3, foldLoopType);
    const lFCur = local('cur', 4, listType);
    const lFAcc = local('acc', 5, intType);
    const lFLoop = lambda(foldLoopType, [['cur', 4], ['acc', 5]],
        cases(intType, lFCur, lFAcc,
            apply(intType, lFold, field(lFCur, 1), bin(intType, new S.OpIntNum(S.OpAdd.value), lFAcc, field(lFCur, 0)))));
    const lRange = local('goR', 6, rangeLoopType);
    const lRCurr = local('curr', 7, intType);
    const lRAcc = local('acc', 8, listType);
    const lRLoop = lambda(rangeLoopType, [['curr', 7], ['acc', 8]],
        typed(listType, expr(new S.Branch([
            new S.Pair(bin(boolType, new S.OpIntOrd(S.OpLt.value), lRCurr, lLo), lRAcc),
        ], apply(listType, lRange, bin(intType, new S.OpIntNum(S.OpSubtract.value), lRCurr, integer(1)), ctor(node, lRCurr, lRAcc))))));
    const localFusedFn = lambda(new C.Func([intType, intType, intType], intType),
        [['lo', 0], ['hi', 1], ['seed', 2]],
        typed(intType, expr(new S.LetRec(3, [new Tuple('goF', lFLoop)],
            typed(intType, expr(new S.LetRec(6, [new Tuple('goR', lRLoop)],
                apply(intType, lFold,
                    apply(listType, globalVar(module, 'prune', transformType),
                        apply(listType, lRange, lHi, ctor(empty))),
                    lSeed))))))));

    // Caller 6 (fused non-additive hash + direct bump + span):
    const hbLo = local('lo', 0, intType);
    const hbHi = local('hi', 1, intType);
    const hbSeed = local('seed', 2, intType);
    const hashBumpedFn = lambda(new C.Func([intType, intType, intType], intType),
        [['lo', 0], ['hi', 1], ['seed', 2]],
        apply(intType, globalVar(module, 'hashChain', new C.Func([intType, listType], intType)),
            hbSeed,
            apply(listType, globalVar(module, 'bump', transformType),
                apply(listType, globalVar(module, 'span', rangeType), hbLo, hbHi))));

    const mod = {
        name: module,
        bindings: [
            { recursive: true, bindings: [new Tuple('keep', keepFn)] },
            { recursive: true, bindings: [new Tuple('prune', pruneFn)] },
            { recursive: true, bindings: [new Tuple('bump', bumpFn)] },
            { recursive: false, bindings: [new Tuple('span', spanFn)] },
            { recursive: false, bindings: [new Tuple('sumChain', sumChainFn)] },
            { recursive: false, bindings: [new Tuple('hashChain', hashChainFn)] },
            { recursive: false, bindings: [new Tuple('pruneShared', pruneSharedFn)] },
            { recursive: false, bindings: [new Tuple('pruneFresh', pruneFreshFn)] },
            { recursive: false, bindings: [new Tuple('bumpFresh', bumpFreshFn)] },
            { recursive: false, bindings: [new Tuple('sumPruned', sumPrunedFn)] },
            { recursive: false, bindings: [new Tuple('localFused', localFusedFn)] },
            { recursive: false, bindings: [new Tuple('hashBumped', hashBumpedFn)] },
        ],
        comments: [],
        imports: emptySet,
        exports: emptySet,
        reExports: emptySet,
        dataTypes: emptyMap,
        dataDecls: [{
            name: adt,
            vars: poly ? ['a'] : [],
            constructors: [
                { name: empty, fields: [] },
                { name: node, fields: [poly ? typeVarA : intType, declAdt] },
            ],
        }],
        classDecls: [],
        foreign: emptyMap,
        implementations: emptyMap,
        directives: emptyMap,
    };

    const metadata = withReboxFields({
        elidedCtors: emptySet,
        ctorTypes: map([
            [`${module}.${node}`, { vars: poly ? ['a'] : [], fields: [poly ? typeVarA : intType, declAdt] }],
            [`${module}.${empty}`, { vars: poly ? ['a'] : [], fields: [] }],
        ]),
        pointerAdtPaths: map([[`${module}.${adt}`, { ctorName: node, arity: poly ? 1 : 0, sharedPairs: [] }]]),
        pointerAdtNodes: singleton(ctorNames.baseStructName),
        pointerAdtLeaves: map([[leafNames.baseStructName, { nodeBaseStruct: ctorNames.baseStructName, nodeCtor: node }]]),
        enumAdts: emptySet,
        enumCtors: emptySet,
        globalTypes: emptyMap,
        globalFunctions: emptyMap,
        classDeclsFields: emptyMap,
    });

    return { mod, metadata, module, node };
}

test('LinearList emits consuming TRMC & fused workers, protects shared callers, and eliminates dead LetRecs', () => {
    for (const poly of [true, false]) {
        const { mod, metadata, module } = buildModule({ poly, module: poly ? 'PolyChain' : 'MonoChain' });
        const prepared = LinearList.prepare(metadata)(mod);
        assert.ok(prepared.declarations.length >= 4, 'expected linear + fused worker declarations');

        const code = CodeGen.translate(metadata)(mod);
        assert.match(code, new RegExp(`func Call_${module}_prune__gopurs_linear_0\\(`));
        assert.match(code, new RegExp(`func Call_${module}_bump__gopurs_linear_0\\(`));
        assert.match(code, new RegExp(`func Call_${module}_prune__gopurs_linear_fold_0\\(`));
        assert.match(code, new RegExp(`func Call_${module}_bump__gopurs_linear_fold_0\\(`));

        // Shared caller MUST call persistent Call_<mod>_prune, never __gopurs_linear_0.
        const sharedMatch = code.match(new RegExp(`func Call_${module}_pruneShared\\([\\s\\S]*?\\n\\}`));
        assert.ok(sharedMatch, 'pruneShared must be emitted');
        assert.match(sharedMatch[0], new RegExp(`Call_${module}_prune\\(`));
        assert.doesNotMatch(sharedMatch[0], /__gopurs_linear_/);

        // Fresh callers MUST use the consuming linear workers.
        const freshMatch = code.match(new RegExp(`func Call_${module}_pruneFresh\\([\\s\\S]*?\\n\\}`));
        assert.ok(freshMatch);
        assert.match(freshMatch[0], new RegExp(`Call_${module}_prune__gopurs_linear_0\\(`));

        const bumpFreshMatch = code.match(new RegExp(`func Call_${module}_bumpFresh\\([\\s\\S]*?\\n\\}`));
        assert.ok(bumpFreshMatch);
        assert.match(bumpFreshMatch[0], new RegExp(`Call_${module}_bump__gopurs_linear_0\\(`));

        // Local LetRec fused caller MUST drop dead local loops and call the fused worker directly.
        const localMatch = code.match(new RegExp(`func Call_${module}_localFused\\([\\s\\S]*?\\n\\}`));
        assert.ok(localMatch);
        assert.match(localMatch[0], new RegExp(`Call_${module}_prune__gopurs_linear_fold_0\\(`));
        assert.doesNotMatch(localMatch[0], /Call_local_/);
    }
});

test('LinearList generated Go code passes -race exactness, immutability, wide int64, >512 growth, and 0-alloc checks', t => {
    const { mod, metadata, module, node } = buildModule({ poly: true, module: 'ChainProbe' });
    const code = CodeGen.translate(metadata)(mod);
    const dir = mkdtempSync(join(tmpdir(), 'gopurs-linear-list-'));
    t.after(() => rmSync(dir, { recursive: true, force: true }));
    mkdirSync(join(dir, 'purescript'));
    mkdirSync(join(dir, 'gopurs_runtime'));
    writeFileSync(join(dir, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(dir, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(dir, `purescript/${module}.go`), code);
    writeFileSync(join(dir, 'purescript/linear_list_test.go'), `package purescript

import (
	"math"
	"reflect"
	"testing"
)

func toSlice(xs *Constructor_${module}_${node}[int64]) []int64 {
	var out []int64
	for cur := xs; cur != nil; cur = cur.V1 {
		out = append(out, cur.V0)
	}
	return out
}

func fromSlice(vals []int64) *Constructor_${module}_${node}[int64] {
	var acc *Constructor_${module}_${node}[int64]
	for i := len(vals) - 1; i >= 0; i-- {
		acc = &Constructor_${module}_${node}[int64]{Rc: 1, V0: vals[i], V1: acc}
	}
	return acc
}

func TestSharedInputIsNeverMutated(t *testing.T) {
	shared := Call_${module}_span(2, 30)
	before := toSlice(shared)
	got := toSlice(Call_${module}_pruneShared(shared))
	after := toSlice(shared)
	if !reflect.DeepEqual(before, after) {
		t.Fatalf("shared list mutated: before=%v after=%v", before, after)
	}
	want := []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("pruneShared mismatch: got=%v want=%v", got, want)
	}
}

func TestFreshLinearWorkersMatchPersistent(t *testing.T) {
	cases := [][2]int64{
		{5, 2},     // empty
		{2, 2},     // singleton
		{2, 50},    // small
		{2, 500},   // boundary <= 512
		{2, 900},   // > 512 slice growth
		{-25, 25},  // negatives, zero, -1, 1
	}
	for _, tc := range cases {
		lo, hi := tc[0], tc[1]
		wantPrune := toSlice(Call_${module}_prune(Call_${module}_span(lo, hi)))
		gotPrune := toSlice(Call_${module}_pruneFresh(lo, hi))
		if !reflect.DeepEqual(gotPrune, wantPrune) {
			t.Fatalf("pruneFresh(%d,%d): got=%v want=%v", lo, hi, gotPrune, wantPrune)
		}
		wantBump := toSlice(Call_${module}_bump(Call_${module}_span(lo, hi)))
		gotBump := toSlice(Call_${module}_bumpFresh(lo, hi))
		if !reflect.DeepEqual(gotBump, wantBump) {
			t.Fatalf("bumpFresh(%d,%d): got=%v want=%v", lo, hi, gotBump, wantBump)
		}
	}

	// Wide int64 values outside int32 range must hit the int64 fallback accurately.
	wideVals := []int64{int64(math.MaxInt32) + 7, 10, int64(math.MinInt32) - 5, 20, 30, 0, -1, 7}
	wantWide := toSlice(Call_${module}_prune(fromSlice(wideVals)))
	gotWide := toSlice(Call_${module}_prune__gopurs_linear_0(fromSlice(wideVals)))
	if !reflect.DeepEqual(gotWide, wantWide) {
		t.Fatalf("wide int64 linear worker: got=%v want=%v", gotWide, wantWide)
	}
}

func TestFusedRangeFoldWorkersExactAndZeroAlloc(t *testing.T) {
	cases := [][3]int64{
		{10, 2, 42},                   // empty returns exact seed
		{10, 2, int64(1) << 40},       // empty returns exact wide int64 seed
		{2, 2, 0},                     // singleton
		{2, 500, 0},                   // primes up to 500 -> 21536
		{2, 500, 100},                 // non-zero seed
		{2, 1200, 0},                  // > 512 slice growth
		{-40, 40, 7},                  // negatives, 0, -1, 1
		{math.MaxInt32 - 3, math.MaxInt32, 0}, // MaxInt32 boundary
	}
	for _, tc := range cases {
		lo, hi, seed := tc[0], tc[1], tc[2]
		wantSum := Call_${module}_sumChain(Call_${module}_prune(Call_${module}_span(lo, hi)))
		if seed == 0 {
			if got := Call_${module}_sumPruned(lo, hi); got != wantSum {
				t.Fatalf("sumPruned(%d,%d): got=%d want=%d", lo, hi, got, wantSum)
			}
		}
		// localFused adds seed + sum(prune(span(lo, hi)))
		gotLocal := Call_${module}_localFused(lo, hi, seed)
		var wantLocal int64
		if hi < lo {
			wantLocal = seed
		} else {
			wantLocal = int64(int32(seed) + int32(wantSum))
		}
		if gotLocal != wantLocal {
			t.Fatalf("localFused(%d,%d,%d): got=%d want=%d", lo, hi, seed, gotLocal, wantLocal)
		}

		// Non-additive hash fold over direct map (bump)
		wantHash := Call_${module}_hashChain(seed, Call_${module}_bump(Call_${module}_span(lo, hi)))
		gotHash := Call_${module}_hashBumped(lo, hi, seed)
		if gotHash != wantHash {
			t.Fatalf("hashBumped(%d,%d,%d): got=%d want=%d", lo, hi, seed, gotHash, wantHash)
		}
	}
	if got := Call_${module}_sumPruned(2, 500); got != 21536 {
		t.Fatalf("sumPruned(2,500) = %d; want 21536", got)
	}
}

func TestFusedRangeFoldZeroAllocations(t *testing.T) {
	allocs := testing.AllocsPerRun(50, func() {
		if Call_${module}_sumPruned(2, 500) != 21536 {
			panic("bad sum")
		}
	})
	if allocs != 0 {
		t.Fatalf("expected 0 allocs/op for fused range+prune+sum (len<=512), got %v", allocs)
	}
}
`);

    const raceRes = spawnSync('go', ['test', '-race', '-count=1', './purescript'], {
        cwd: dir,
        encoding: 'utf8',
        timeout: 60_000,
        env: { ...process.env, GOWORK: 'off' },
    });
    assert.equal(raceRes.status, 0, `go test -race failed:\n${raceRes.stdout}\n${raceRes.stderr}`);

    // AllocsPerRun must be checked without -race (race detector adds instrumentation overhead).
    const allocRes = spawnSync('go', ['test', '-run', 'TestFusedRangeFoldZeroAllocations', '-count=1', './purescript'], {
        cwd: dir,
        encoding: 'utf8',
        timeout: 60_000,
        env: { ...process.env, GOWORK: 'off' },
    });
    assert.equal(allocRes.status, 0, `0-alloc test failed:\n${allocRes.stdout}\n${allocRes.stderr}`);
});
