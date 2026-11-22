import assert from 'node:assert/strict';
import test from 'node:test';
import { empty as emptyMap, insert } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import * as CodeGen from '../output/Gopurs.CodeGen/index.js';
import { StmtEmpty, StmtLeaf, flattenStmts } from '../output/Gopurs.ExprContext/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { stringLengthComparison } from '../output/Gopurs.PrimitiveExprs/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';
import { TcoExpr } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { NeutralExpr } from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { withReboxFields } from './codegen-metadata.mjs';
import { goModule, runGo, writeGoFile } from './lib/go.mjs';

const signature = new C.Func([C.String.value], C.Int.value);
const metadata = (globalTypes = insert(ordString)('Data.String.CodeUnits.length')(signature)(emptyMap)) => withReboxFields({
    pointerAdtPaths: emptyMap, pointerAdtNodes: emptySet, pointerAdtLeaves: emptyMap,
    enumAdts: emptySet, enumCtors: emptySet, elidedCtors: emptySet,
    ctorTypes: emptyMap, classDeclsFields: emptyMap, globalTypes, globalFunctions: emptyMap,
});
const context = (meta = metadata()) => ({
    metadata: meta, codegenStateRef: Ref.new({ declarations: [], globalId: 0, reboxPairs: emptySet })(),
    depth: 0, modNameStr: 'LengthProbe', recVars: [], moduleFunctions: emptyMap, bound: emptyMap,
    tcoIdent: Nothing.value, loopCtx: [], options: { isTail: false, inEffectBlock: false },
    mbExpectedExprType: Nothing.value, mbExpectedGoType: Nothing.value,
});
const expr = value => new TcoExpr(null, value);
const typed = (type, value) => expr(new S.Typed(type, value));
const ref = (module, name) => expr(new S.Var(new C.Qualified(module === null ? Nothing.value : new Just(module), name)));
const left = typed(C.String.value, ref(null, 'left'));
const right = typed(C.String.value, ref(null, 'right'));
const unannotatedLeft = left.value1.value1;
const unannotatedRight = right.value1.value1;
const length = (value, fn = typed(signature, ref('Data.String.CodeUnits', 'length'))) =>
    typed(C.Int.value, expr(new S.App(fn, [value])));
const equal = new S.OpIntOrd(S.OpEq.value);
const unequal = new S.OpIntOrd(S.OpNotEq.value);

test('length comparison requires two canonical typed calls and only admits integer equality', () => {
    let translations = 0;
    const translate = child => nextId => argument => {
        translations++;
        assert.deepEqual(child.mbExpectedExprType, new Just(C.String.value));
        assert.ok([left, right, unannotatedLeft, unannotatedRight].includes(argument));
        return { stmts: StmtEmpty.value, expr: new Go.GoVar(argument === left || argument === unannotatedLeft ? 'left' : 'right'),
            exprType: Go.TypeString.value, nextId: nextId + 1 };
    };
    const emit = (op, a, b, ctx = context()) => stringLengthComparison(translate)(ctx)(3)(op)(a)(b);
    for (const op of [equal, unequal]) {
        for (const fn of [typed(signature, ref('Data.String.CodeUnits', 'length')), ref('Data.String.CodeUnits', 'length')]) {
            translations = 0;
            const result = emit(op, length(left, fn), length(right, fn));
            assert.ok(result instanceof Just);
            assert.equal(result.value0.nextId, 5);
            assert.equal(translations, 2);
            assert.match(printGoExpr(result.value0.expr), /StringCodeUnitsLengthsEqual/);
        }
    }
    for (const [a, b] of [
        [typed(C.Any.value, length(unannotatedLeft)), typed(C.Any.value, length(unannotatedRight))],
        [expr(new S.UncurriedApp(typed(signature, ref('Data.String.CodeUnits', 'length')), [left])),
            expr(new S.UncurriedApp(typed(signature, ref('Data.String.CodeUnits', 'length')), [right]))],
    ]) {
        translations = 0;
        assert.ok(emit(equal, a, b) instanceof Just);
        assert.equal(translations, 2);
    }
    const invalid = [
        length(left, typed(signature, ref('Other', 'length'))),
        length(left, typed(signature, ref(null, 'length'))),
        length(left, typed(signature, ref('Data.String.CodeUnits', 'other'))),
        length(left, typed(new C.Func([C.String.value], C.Number.value), ref('Data.String.CodeUnits', 'length'))),
        length(left, typed(new C.Func([C.String.value, C.String.value], C.Int.value), ref('Data.String.CodeUnits', 'length'))),
        typed(C.Number.value, length(left)),
        length(typed(C.Number.value, ref(null, 'number'))),
        typed(C.Int.value, expr(new S.App(typed(signature, ref('Data.String.CodeUnits', 'length')), [left, right]))),
        typed(C.Int.value, expr(new S.App(expr(new S.UncurriedApp(typed(signature, ref('Data.String.CodeUnits', 'length')), [])), [left]))),
        typed(C.Int.value, expr(new S.Lit(new C.LitInt(1)))),
    ];
    for (const bad of invalid) {
        translations = 0;
        assert.deepEqual(emit(equal, bad, length(right)), Nothing.value);
        assert.deepEqual(emit(equal, length(left), bad), Nothing.value);
        assert.equal(translations, 0, 'rejected operands must not be translated');
    }
    for (const op of [new S.OpIntOrd(S.OpLt.value), new S.OpIntOrd(S.OpGte.value), new S.OpNumberOrd(S.OpEq.value)]) {
        translations = 0;
        assert.deepEqual(emit(op, length(left), length(right)), Nothing.value);
        assert.equal(translations, 0);
    }
    const unknown = ref('Data.String.CodeUnits', 'length');
    assert.deepEqual(emit(equal, length(left, unknown), length(right), context(metadata(emptyMap))), Nothing.value);
});

function compiledModule() {
    const e = value => new NeutralExpr(value);
    const annotate = (type, value) => e(new S.Typed(type, value));
    const local = (name, level) => annotate(C.String.value, e(new S.Local(new Just(name), level)));
    const call = (module, value) => annotate(C.Int.value, e(new S.App(
        annotate(signature, e(new S.Var(new C.Qualified(new Just(module), 'length')))), [value])));
    const bindings = [
        ['equal', equal, 'Data.String.CodeUnits'],
        ['unequal', unequal, 'Data.String.CodeUnits'],
        ['less', new S.OpIntOrd(S.OpLt.value), 'Data.String.CodeUnits'],
        ['otherEqual', equal, 'Other'],
    ].map(([name, operator, module]) => {
        const body = annotate(C.Boolean.value, e(new S.PrimOp(new S.Op2(operator,
            call(module, local('left', 0)), call(module, local('right', 1))))));
        const fn = annotate(new C.Func([C.String.value, C.String.value], C.Boolean.value),
            e(new S.Abs([new Tuple(new Just('left'), 0), new Tuple(new Just('right'), 1)], body)));
        return { recursive: false, bindings: [new Tuple(name, fn)] };
    });
    return CodeGen.translate(metadata())({ name: 'LengthProbe', bindings, comments: [], imports: emptySet,
        exports: emptySet, reExports: emptySet, dataTypes: emptyMap, dataDecls: [], classDecls: [],
        foreign: emptyMap, implementations: emptyMap, directives: emptyMap });
}

test('public codegen emits equality and inequality helpers while retaining other length semantics', t => {
    const code = compiledModule();
    assert.equal((code.match(/StringCodeUnitsLengthsEqual/g) ?? []).length, 2);
    assert.match(code, /Get_Data_String_CodeUnits_length/);
    assert.match(code, /Get_Other_length/);
    const directory = goModule(t, 'gopurs-length-codegen-', { runtime: runtimeGoCode });
    writeGoFile(directory, 'purescript/LengthProbe.go', code);
    writeGoFile(directory, 'purescript/length_test.go', `package purescript
import ("testing"; "gopurs/output/gopurs_runtime")
func Get_Data_String_CodeUnits_length() gopurs_runtime.Value {
    return gopurs_runtime.Func(func(v gopurs_runtime.Value) gopurs_runtime.Value {
        return gopurs_runtime.Int(int64(gopurs_runtime.StringCodeUnitsLength(v)))
    })
}
func Get_Other_length() gopurs_runtime.Value {
    return gopurs_runtime.Func(func(v gopurs_runtime.Value) gopurs_runtime.Value {
        if v.StrVal()=="a" { return gopurs_runtime.Int(1) }; return gopurs_runtime.Int(2)
    })
}
func TestCompiledComparisons(t *testing.T) {
    for _, pair := range []struct { left,right string; equal bool } {
        {"", "", true}, {"é", "a", true}, {"é", "ab", false}, {"😀", "ab", true}, {"a", "abc", false},
    } {
        if Call_LengthProbe_equal(pair.left,pair.right)!=pair.equal { t.Fatal("compiled equality",pair) }
        if Call_LengthProbe_unequal(pair.left,pair.right)==pair.equal { t.Fatal("compiled inequality",pair) }
    }
    if !Call_LengthProbe_less("a","😀") || Call_LengthProbe_less("😀","ab") { t.Fatal("ordered length comparison changed") }
    if Call_LengthProbe_otherEqual("a","b") { t.Fatal("unrelated length function optimized") }
}
`);
    runGo(directory, ['test', '-count=1', './...']);
});

test('operands preserve statement/expression order, evaluate once and retain early panic', t => {
    const declarations = [];
    const names = [];
    for (const [operation, operator] of [['eq', equal], ['ne', unequal]]) {
        for (const leftNative of [false, true]) for (const rightNative of [false, true]) {
            const name = `${operation}_${leftNative ? 'native' : 'boxed'}_${rightNative ? 'native' : 'boxed'}`;
            const translate = _context => nextId => argument => {
                const label = argument === left ? 'left' : 'right';
                const native = label === 'left' ? leftNative : rightNative;
                return { stmts: new StmtLeaf(Go.rawGo(`mark("${label}:statement")`)),
                    expr: Go.rawGo(`observed("${label}:expression", ${label})`),
                    exprType: native ? Go.TypeString.value : Go.TypeValue.value, nextId: nextId + 1 };
            };
            const result = stringLengthComparison(translate)(context())(0)(operator)(length(left))(length(right)).value0;
            assert.equal(result.nextId, 2);
            const body = [...flattenStmts(result.stmts), new Go.GoReturn(result.expr)].map(printGoExpr).join('\n');
            declarations.push(`func ${name}(left ${leftNative ? 'string' : 'gopurs_runtime.Value'}, right ${rightNative ? 'string' : 'gopurs_runtime.Value'}) bool { ${body} }`);
            names.push({ name, operation, leftNative, rightNative });
        }
    }
    const directory = goModule(t, 'gopurs-length-evaluation-', { runtime: runtimeGoCode });
    writeGoFile(directory, 'evaluation_test.go', `package main
import ("reflect"; "testing"; "gopurs/output/gopurs_runtime")
var trace []string
var panicLeft bool
func mark(label string) { trace=append(trace,label) }
func observed[T any](label string,value T) T { mark(label); if panicLeft && label=="left:expression" { panic("left") }; return value }
${declarations.join('\n')}
func TestEvaluation(t *testing.T) {
    left,right:="é","a"
    probes:=[]struct{name string; want bool; run func()bool}{
        ${names.map(({ name, operation, leftNative, rightNative }) => `{"${name}",${operation === 'eq'},func()bool{return ${name}(${leftNative ? 'left' : 'gopurs_runtime.StrWithCodeUnits(left,1)'},${rightNative ? 'right' : 'gopurs_runtime.StrWithCodeUnits(right,1)'})}}`).join(',\n')},
    }
    for _,probe:=range probes {
        panicLeft=false;trace=nil
        if probe.run()!=probe.want { t.Fatal(probe.name,"value") }
        if !reflect.DeepEqual(trace,[]string{"left:statement","left:expression","right:statement","right:expression"}) { t.Fatal(probe.name,"order/count",trace) }
        panicLeft=true;trace=nil
        func(){defer func(){if recover()!="left" {t.Fatal(probe.name,"panic")}}();probe.run()}()
        if !reflect.DeepEqual(trace,[]string{"left:statement","left:expression"}) { t.Fatal(probe.name,"right operand ran after left panic",trace) }
    }
}
`);
    runGo(directory, ['test', '-count=1', '.']);
});
