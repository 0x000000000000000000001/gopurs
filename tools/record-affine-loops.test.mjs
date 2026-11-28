import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import * as Go from '../output/Gopurs.GoAst/index.js';
import * as RecordLoops from '../output/Gopurs.RecordAffineLoops/index.js';
import * as Affine from '../output/Gopurs.AffineLoops/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { TcoExpr } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import { localId } from '../output/PureScript.Backend.Optimizer.FreeVars/index.js';
import { Just } from '../output/Data.Maybe/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { printGoDecl, printGoExpr } from '../output/Gopurs.Printer/index.js';

const int = C.Int.value, i64 = Go.TypeInt64.value;
const expr = value => new TcoExpr(null, value);
const typed = value => expr(new S.Typed(int, value));
const local = (name, level) => expr(new S.Local(new Just(name), level));
const n = local('remaining', 0), record = local('state', 1);
const nName = localId(new Just('remaining'))(0), recordName = localId(new Just('state'))(1);
const nestedType = new Go.TypeRecord([new Tuple('value', i64)]);
const recordType = new Go.TypeRecord([
  new Tuple('sum', i64), new Tuple('nested', nestedType),
  new Tuple('tag', Go.TypeString.value), new Tuple('borrowed', new Go.TypeNativeArray(i64)),
]);
const params = [new Tuple(nName, i64), new Tuple(recordName, recordType)];
const integer = value => typed(expr(new S.Lit(new C.LitInt(value))));
const binary = (operator, a, b) => typed(expr(new S.PrimOp(new S.Op2(operator, a, b))));
const add = (a, b) => binary(new S.OpIntNum(S.OpAdd.value), a, b);
const sub = (a, b) => binary(new S.OpIntNum(S.OpSubtract.value), a, b);
const mod = (a, b) => binary(new S.OpIntNum(S.OpMod.value), a, b);
const eq = (a, b) => binary(new S.OpIntOrd(S.OpEq.value), a, b);
const get = (source, key) => expr(new S.Accessor(source, new S.GetProp(key)));
const update = (source, entries) => expr(new S.Update(source, entries.map(([k, v]) => new C.Prop(k, v))));
const variable = (name = 'step', owner = 'Renamed') => expr(new S.Var(new C.Qualified(new Just(owner), name)));
const call = (target, args) => expr(new S.App(target, args));
const nested = get(record, 'nested');
const sum = get(record, 'sum'), value = get(nested, 'value');
const fallback = new Go.GoReturn(new Go.GoVar(recordName + '_loop'));
const self = Affine.globalTarget('Renamed')('step');

function fixture(options = {}) {
  const increment = integer(options.increment ?? 3);
  const nestedInc = integer(options.nestedIncrement ?? 7);
  const changes = options.changes ?? update(record, [
    ['sum', options.sum ?? (options.commute ? add(increment, sum) : add(sum, increment))],
    ['nested', update(nested, [['value', options.value ?? (options.commute ? add(nestedInc, value) : add(value, nestedInc))]])],
  ]);
  const args = options.args ?? (options.reversed ? [changes, sub(n, integer(1))] : [sub(n, integer(1)), changes]);
  const step = options.step ?? (options.uncurried ? expr(new S.UncurriedApp(variable(), args)) : call(options.target ?? variable(), args));
  return expr(new S.Branch([new S.Pair(options.condition ?? (options.commute ? eq(integer(0), n) : eq(n, integer(0))), options.seed ?? record)], step));
}

function optimize(body, options = {}) {
  const before = JSON.stringify(body);
  const wrap = options.hook ? Affine.wrap(() => { throw new Error('record recurrence must not emit global reads'); })({}) : RecordLoops.wrap;
  const result = wrap(options.self ?? self)
    (options.params ?? params)(options.result ?? recordType)(body)(options.fallback ?? fallback);
  assert.equal(JSON.stringify(body), before, 'the source recurrence stays immutable');
  return result;
}

test('recognizes renamed nested records and equivalent parameter/operator orders', () => {
  for (const options of [{}, {commute:true}, {uncurried:true}, {reversed:true}, {nestedIncrement:0}, {nestedIncrement:2147483647}, {increment:-2147483648}]) {
    const result = optimize(fixture(options), {params:options.reversed ? [...params].reverse() : params, hook:true});
    assert.notEqual(result, fallback);
    assert.equal(result.value0.at(-1), fallback);
    const code = printGoExpr(result);
    assert.match(code, />.*int64\(0\)/s);
    assert.match(code, /2147483647/);
    assert.match(code, /int64\(int32\(/);
    assert.doesNotMatch(code, /IntMul|IntMod|IntAdd/);
  }
});

test('accepts untouched fields and an unchanged record without evaluating anything new', () => {
  assert.notEqual(optimize(fixture({changes:record})), fallback);
  assert.notEqual(optimize(fixture({changes:update(record,[['sum',sum],['tag',get(record,'tag')]])})), fallback);
});

test('rejects altered termination, decrement, base value, arity and recursion target', () => {
  for (const options of [
    {condition:eq(n,integer(1))}, {condition:eq(integer(0),sum)}, {seed:n},
    {args:[sub(n,integer(2)),record]}, {args:[add(n,integer(1)),record]},
    {args:[n,record]}, {args:[sub(n,integer(1)),record,integer(0)]},
    {target:variable('other')}, {target:variable('step','Foreign')},
    {step:expr(new S.EffectPure(record))},
  ]) assert.equal(optimize(fixture(options)), fallback);
});

test('rejects cross-field dependencies and changed access paths', () => {
  for (const options of [
    {sum:add(value,integer(3))}, {value:add(sum,mod(n,integer(5)))},
    {sum:add(sum,n)}, {value:add(value,mod(sum,integer(5)))},
    {changes:update(nested,[['value',add(value,integer(1))]])},
    {changes:update(local('state',2),[['sum',add(sum,integer(1))]])},
  ]) assert.equal(optimize(fixture(options)), fallback);
});

test('rejects effectful or opaque work even in otherwise ignored updates', () => {
  for (const opaque of [call(variable('observe'),[integer(1)]), expr(new S.Fail('must remain')), expr(new S.EffectPure(integer(3)))]) {
    assert.equal(optimize(fixture({sum:add(sum,opaque)})), fallback);
    assert.equal(optimize(fixture({changes:update(record,[['sum',add(sum,integer(1))],['tag',opaque]])})), fallback);
  }
  const wrapped = expr(new S.Let(new Just('seen'),4,call(variable('observe'),[record]),fixture()));
  assert.equal(optimize(wrapped), fallback);
});

test('rejects modulo/non-constant operations and captured increments', () => {
  for (const candidate of [add(value,mod(n,integer(5))), add(value,mod(n,sum)), sub(value,integer(2)), add(value,mod(add(n,integer(1)),integer(5)))]) {
    assert.equal(optimize(fixture({value:candidate})), fallback);
  }
  assert.equal(optimize(fixture({sum:add(sum,local('capture',7))})), fallback);
});

test('rejects duplicate or missing fields and incompatible layouts', () => {
  assert.equal(optimize(fixture({changes:update(record,[['sum',sum],['sum',add(sum,integer(1))]])})), fallback);
  assert.equal(optimize(fixture({changes:update(record,[['missing',integer(0)]])})), fallback);
  assert.equal(optimize(fixture(), {result:i64}), fallback);
  assert.equal(optimize(fixture(), {params:[new Tuple(nName,i64),new Tuple(recordName,Go.TypeValue.value)]}), fallback);
});

test('emitted Go matches the loop and preserves wide seeds, borrowed fields and fallback', t => {
  const directory = process.env.GOPURS_RECORD_AFFINE_TEST_OUTPUT ?? mkdtempSync(join(tmpdir(),'record-affine-'));
  mkdirSync(directory,{recursive:true});
  if (!process.env.GOPURS_RECORD_AFFINE_TEST_OUTPUT) t.after(() => rmSync(directory,{recursive:true,force:true}));
  const definitions = [], configs = [{increment:3,nestedIncrement:5},{increment:0,nestedIncrement:1},{increment:-7,nestedIncrement:7},{increment:2147483647,nestedIncrement:-2147483648}];
  const count = nName+'_loop', state = recordName+'_loop';
  const print = (name,body) => printGoDecl(new Go.GoFunctionDecl({name:{owner:'generated/Probe',name},params:params.map(p=>new Tuple(p.value0+'_loop',p.value1)),result:recordType,body}));
  for (const [i,config] of configs.entries()) {
    const raw = Go.rawGo(`for ${count} != 0 { ${state}.sum=int64(int32(${state}.sum+${config.increment})); ${state}.nested.value=int64(int32(${state}.nested.value+${config.nestedIncrement})); ${count}=int64(int32(${count}-1)) }; return ${state}`);
    definitions.push(print('Run'+i,optimize(fixture(config),{fallback:raw})));
    definitions.push(print('Guard'+i,optimize(fixture(config))));
  }
  const large=[];
  for(const [i,{increment,nestedIncrement}] of configs.entries())for(const count of [2147483646,2147483647])for(const seed of [-2147483648,0,2147483647,4294967295]){
    const n=BigInt(count);
    large.push(`{Run${i},${count},${seed},${BigInt.asIntN(32,BigInt(seed)+n*BigInt(increment))},${BigInt.asIntN(32,BigInt(seed)+n*BigInt(nestedIncrement))}}`);
  }
  writeFileSync(join(directory,'go.mod'),'module recordaffineprobe\n\ngo 1.22\n');
  writeFileSync(join(directory,'record.go'),`package recordaffineprobe\n${definitions.join('\n')}`);
  writeFileSync(join(directory,'record_test.go'),`package recordaffineprobe
import ("testing";"reflect")
type state = ${Go.goTypeToStr(recordType)}
func TestEmitted(t *testing.T) {
 cases:=0
 configs:=[]struct{f,g func(int64,state)state; increment,nestedIncrement int64}{${configs.map((c,i)=>`{Run${i},Guard${i},${c.increment},${c.nestedIncrement}}`).join(',')}}
 for _,c:=range configs { for _,n:=range []int64{0,1,2,4,5,6,7,30,1000,10000,4294967297,4294967303} { for _,seed:=range []int64{-9223372036854775808,-2147483648,-1,0,2147483647,2147483648,4294967295,9223372036854775807} {
  input:=state{sum:seed,nested:struct{value int64}{seed},tag:"retained",borrowed:[]int64{7,11}};want:=input
  for k:=n;k!=0;k=int64(int32(k-1)){want.sum=int64(int32(want.sum+c.increment));want.nested.value=int64(int32(want.nested.value+c.nestedIncrement))}
  got:=c.f(n,input); if !reflect.DeepEqual(got,want)||input.sum!=seed||input.nested.value!=seed||&got.borrowed[0]!=&input.borrowed[0]{t.Fatalf("n=%d seed=%d got=%v want=%v",n,seed,got,want)};cases++
 } } }
 for _,c:=range []struct{f func(int64,state)state;n,seed,a,b int64}{${large.join(',')}}{input:=state{sum:c.seed,nested:struct{value int64}{c.seed}};got:=c.f(c.n,input);if got.sum!=c.a||got.nested.value!=c.b{t.Fatalf("large n=%d got=%v want=%d/%d",c.n,got,c.a,c.b)};cases++}
 for _,c:=range configs{for _,n:=range []int64{-2147483648,-1,0,2147483648,4294967297}{input:=state{sum:4294967295,nested:struct{value int64}{-2147483649}};if got:=c.g(n,input);!reflect.DeepEqual(got,input){t.Fatalf("fallback n=%d got=%v",n,got)};cases++}}
 t.Logf("%d emitted semantic, boundary and fallback cases passed",cases)
}
`);
  const run=spawnSync('go',['test','-count=1','-v','.'],{cwd:directory,encoding:'utf8',timeout:60000,env:{...process.env,GOWORK:'off'}});
  assert.ifError(run.error); assert.equal(run.status,0,run.stdout+run.stderr);
  if(process.env.GOPURS_RECORD_AFFINE_TEST_OUTPUT)writeFileSync(join(directory,'go-test.log'),run.stdout+run.stderr);
});
