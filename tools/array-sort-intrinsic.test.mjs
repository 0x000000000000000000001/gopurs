import assert from 'node:assert/strict';
import test from 'node:test';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { empty as emptyMap, insert } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { empty as emptySet, insert as insertSet } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import { emit } from '../output/Gopurs.ArraySort/index.js';
import { StmtLeaf, flattenStmts } from '../output/Gopurs.ExprContext/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { TcoExpr } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { hashString } from '../output/PureScript.Backend.Optimizer.FfiSupport/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';
import { withReboxFields } from './codegen-metadata.mjs';

const metadata = withReboxFields({elidedCtors:emptySet,
    ctorTypes:['LT','EQ','GT'].reduce((m,n)=>insert(ordString)(`Data.Ordering.${n}`)({vars:[],fields:[]})(m),emptyMap),
    pointerAdtPaths:emptyMap,pointerAdtNodes:emptySet,pointerAdtLeaves:emptyMap,
    enumAdts:insertSet(ordString)('Data.Ordering.Ordering')(emptySet),
    enumCtors:['LT','EQ','GT'].reduce((s,n)=>insertSet(ordString)(`Data_Data_Ordering_${n}`)(s),emptySet),
    globalTypes:emptyMap,globalFunctions:emptyMap,classDeclsFields:emptyMap});
const context=()=>({metadata,codegenStateRef:Ref.new({declarations:[],globalId:0,reboxPairs:emptySet})(),
    depth:0,modNameStr:'Probe',recVars:[],moduleFunctions:emptyMap,bound:emptyMap,tcoIdent:Nothing.value,loopCtx:[],
    options:{isTail:false,inEffectBlock:false},mbExpectedExprType:Nothing.value,mbExpectedGoType:Nothing.value});
const variable=(module,name)=>new TcoExpr(null,new S.Var(new C.Qualified(new Just(module),name)));
const comparator=(module='Data.Ord',name='ordInt')=>new TcoExpr(null,new S.Accessor(variable(module,name),new S.GetProp('compare')));
const fn=variable('Data.Array','sortByImpl');
const args=[comparator(),variable('Probe','conversion'),variable('Probe','array')];

test('native sort requires the standard Int comparator and exact FFI application',()=>{
    const forbidden=()=>{throw Error('unexpected translation')};
    for(const [target,supplied] of [[variable('Other','sortByImpl'),args],[fn,args.slice(0,2)],[fn,[...args,args[2]]],
        [fn,[comparator('Other'),...args.slice(1)]],[fn,[comparator('Data.Ord','ordNumber'),...args.slice(1)]],
        [fn,[variable('Probe','customCompare'),...args.slice(1)]]]) assert.deepEqual(emit(forbidden)(context())(0)(target)(supplied),Nothing.value);
});

test('native sort preserves the library algorithm, conversion calls, strict capture and fresh input',t=>{
    const declarations=[];
    for(const boxed of [false,true])for(const native of [false,true]){
        const translate=_context=>nextId=>expression=>{
            const index=args.indexOf(expression),label=['comparator','conversion','array'][index];
            const type=index===0?Go.TypeValue.value:index===1?(native?new Go.TypeFunc([Go.TypeUint32.value],Go.TypeInt64.value):Go.TypeValue.value)
                :boxed?Go.TypeValue.value:new Go.TypeNativeArray(Go.TypeInt64.value);
            return {stmts:new StmtLeaf(Go.rawGo(`trace=append(trace,"${label}:statement")`)),expr:Go.rawGo(`capture("${label}:expression",${label})`),exprType:type,nextId};
        };
        const result=emit(translate)(context())(0)(fn)(args);assert.ok(result instanceof Just);
        const body=[...flattenStmts(result.value0.stmts),new Go.GoReturn(result.value0.expr)].map(printGoExpr).join('\n');
        declarations.push(`func sort_${boxed}_${native}(array ${boxed?'r.Value':'[]int64'}, conversion ${native?'func(uint32)int64':'r.Value'})[]int64{comparator:=r.Value{}\n${body}\n}`);
    }
    const directory=mkdtempSync(join(tmpdir(),'gopurs-sort-intrinsic-'));t.after(()=>rmSync(directory,{recursive:true,force:true}));
    mkdirSync(join(directory,'gopurs_runtime'));writeFileSync(join(directory,'go.mod'),'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory,'gopurs_runtime/runtime.go'),runtimeGoCode);
    writeFileSync(join(directory,'array.go'),'package main\nimport "gopurs/output/gopurs_runtime"\n'+readFileSync(new URL('../../gopurs-arrays/src/Data/Array.go',import.meta.url),'utf8'));
    writeFileSync(join(directory,'main.go'),`package main
import("fmt";"slices";"reflect";r "gopurs/output/gopurs_runtime";"gopurs/output/gopurs_runtime")
var trace []string
func capture[T any](s string,v T)T{trace=append(trace,s);return v}
const lt=${hashString('Data_Data_Ordering_LT')};const eq=${hashString('Data_Data_Ordering_EQ')};const gt=${hashString('Data_Data_Ordering_GT')}
func ordering(a,b int64)uint32{if a<b{return lt};if a==b{return eq};return gt}
${declarations.join('\n')}
func main(){
    cases:=[][]int64{{},{1},{2,1},{1,2},{2,2,1},{2,1,2,0,1,2},{0,0,0,0}}
    for seed:=int64(1);seed<20;seed++{input:=make([]int64,seed*11);x:=seed;for i:=range input{x=(x*1664525+1013904223)&0xffffffff;input[i]=x%7};cases=append(cases,input)}
    for _,input:=range cases{for _,mode:=range []int64{0,1,2}{
        original:=append([]int64{},input...);values:=make([]r.Value,len(input));for i,v:=range input{values[i]=r.Int(v)}
        convert:=func(tag uint32)int64{trace=append(trace,fmt.Sprint("convert:",tag));if mode==2{return 0};answer:=int64(0);if tag==lt{answer=-1};if tag==gt{answer=1};if mode==1{return -answer};return answer}
        dynamic:=r.Func(func(v r.Value)r.Value{return r.Int(convert(uint32(v.IntVal)))})
        trace=nil;expectedValues:=SortByImpl(func(a,b r.Value)r.Value{return r.Value{Type:r.TypeConstructor,IntVal:int64(ordering(a.IntVal,b.IntVal))}},func(v r.Value)int64{return convert(uint32(v.IntVal))},values)
        expected:=make([]int64,len(expectedValues));for i,v:=range expectedValues{expected[i]=v.IntVal};calls:=append([]string{},trace...)
        probes:=[]func()[]int64{func()[]int64{return sort_false_false(input,dynamic)},func()[]int64{return sort_false_true(input,convert)},func()[]int64{return sort_true_false(r.Array(values),dynamic)},func()[]int64{return sort_true_true(r.Array(values),convert)}}
        for _,probe:=range probes{trace=[]string{};actual:=probe();prefix:=[]string{"comparator:statement","comparator:expression","conversion:statement","conversion:expression","array:statement","array:expression"};
            if !slices.Equal(actual,expected)||!reflect.DeepEqual(trace,append(prefix,calls...)){panic(fmt.Sprint("sort/conversion mismatch ",input,mode,actual,expected,trace,calls))}
            if len(actual)>0{actual[0]=99};if !slices.Equal(input,original){panic("native input mutated")};for i,v:=range values{if v.IntVal!=input[i]{panic("boxed input mutated")}}
        }
    }}
    fmt.Println("stable Int sort, custom conversion order and immutable inputs: ok")
}
`);
    const run=spawnSync('go',['run','.'],{cwd:directory,encoding:'utf8',timeout:60000,env:{...process.env,GOWORK:'off'}});
    assert.ifError(run.error);assert.equal(run.status,0,run.stdout+run.stderr);assert.match(run.stdout,/stable Int sort, custom conversion order/);
});

test('static Ordering -> Int lambdas specialize without per-element comparison closures', t => {
    const tagCond = (name, level) => new TcoExpr(null, new S.PrimOp(new S.Op1(
        new S.OpIsTag(new C.Qualified(new Just('Data.Ordering'), name)),
        new TcoExpr(null, new S.Local(new Just('o'), level)))));
    const intLit = n => new TcoExpr(null, new S.Lit(new C.LitInt(n)));
    const staticLambda = (lt, eq, gt) => new TcoExpr(null, new S.Abs([new Tuple(new Just('o'), 7)],
        new TcoExpr(null, new S.Branch([
            new S.Pair(tagCond('GT', 7), intLit(gt)),
            new S.Pair(tagCond('EQ', 7), intLit(eq)),
            new S.Pair(tagCond('LT', 7), intLit(lt)),
        ], new TcoExpr(null, new S.Fail('Failed pattern match'))))));
    const specs = [
        { name: 'asc', lt: -1, eq: 0, gt: 1, expectNoCompare: true },
        { name: 'desc', lt: 1, eq: 0, gt: -1, expectNoCompare: true },
        { name: 'zero', lt: 0, eq: 0, gt: 0, expectNoCompare: false },
    ];
    const declarations = [];
    for (const spec of specs) {
        const supplied = [comparator(), staticLambda(spec.lt, spec.eq, spec.gt), variable('Probe', 'array')];
        const translate = _ctx => nextId => expr => {
            if (expr === supplied[1]) throw Error('static conversion lambda must not be translated');
            const isArray = expr === supplied[2];
            return { stmts: new StmtLeaf(Go.rawGo(`trace=append(trace,"${isArray ? 'array' : 'comparator'}")`)),
                expr: Go.rawGo(isArray ? 'array' : 'r.Value{}'),
                exprType: isArray ? new Go.TypeNativeArray(Go.TypeInt64.value) : Go.TypeValue.value, nextId };
        };
        const result = emit(translate)(context())(0)(fn)(supplied);
        assert.ok(result instanceof Just);
        const body = [...flattenStmts(result.value0.stmts), new Go.GoReturn(result.value0.expr)].map(printGoExpr).join('\n');
        if (spec.expectNoCompare) assert.doesNotMatch(body, /arraySort_compare/);
        assert.match(body, /copy\(arraySort_output/);
        declarations.push(`func sort_${spec.name}(array []int64) []int64 {\n${body}\n}`);
    }
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-sort-static-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'main.go'), `package main
import("fmt";"slices";r "gopurs/output/gopurs_runtime")
var trace []string
${declarations.join('\n')}
func main() {
    cases := [][]int64{{}, {1}, {2,1}, {1,2}, {3,2,2}, {2,2,1}, {1,2,2,3}, {5,4,3,2,1}, {2,1,2,0,1,2}, {0,0,0,0}}
    for _, input := range cases {
        orig := append([]int64{}, input...)
        trace = nil; gotAsc := sort_asc(input)
        wantAsc := append([]int64{}, input...); slices.Sort(wantAsc)
        if !slices.Equal(gotAsc, wantAsc) { panic(fmt.Sprint("asc mismatch", input, gotAsc, wantAsc)) }
        if len(trace) != 2 || trace[0] != "comparator" || trace[1] != "array" { panic(fmt.Sprint("trace", trace)) }
        gotDesc := sort_desc(input)
        wantDesc := append([]int64{}, wantAsc...); slices.Reverse(wantDesc)
        if !slices.Equal(gotDesc, wantDesc) { panic(fmt.Sprint("desc mismatch", input, gotDesc, wantDesc)) }
        gotZero := sort_zero(input)
        if !slices.Equal(gotZero, orig) { panic(fmt.Sprint("zero mismatch", input, gotZero, orig)) }
        if !slices.Equal(input, orig) { panic("mutated") }
    }
    fmt.Println("static ordering specialization: ok")
}
`);
    const run = spawnSync('go', ['run', '.'], { cwd: directory, encoding: 'utf8', timeout: 60000, env: { ...process.env, GOWORK: 'off' } });
    assert.ifError(run.error); assert.equal(run.status, 0, run.stdout + run.stderr);
    assert.match(run.stdout, /static ordering specialization: ok/);
});
