import assert from 'node:assert/strict';
import test from 'node:test';
import { mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { spawnSync } from 'node:child_process';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { empty as emptyMap } from '../output/Data.Map/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import { TcoExpr } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import { StmtEmpty } from '../output/Gopurs.ExprContext/index.js';
import * as G from '../output/Gopurs.GoAst/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { worker } from '../output/Gopurs.TailAddLoops/index.js';

const expr = syntax => new TcoExpr(null, syntax);
const typed = (ty, value) => expr(new S.Typed(ty, value));
const lit = n => typed(C.Int.value, expr(new S.Lit(new C.LitInt(n))));
const local = (name, level = 0) => typed(C.Int.value, expr(new S.Local(new Just(name), level)));
const variable = (name, owner = 'Renamed') => expr(new S.Var(new C.Qualified(owner === null ? Nothing.value : new Just(owner), name)));
const call = (name, args, owner = 'Renamed') => typed(C.Int.value, expr(new S.App(variable(name, owner), args)));
const op = (kind, a, b) => typed(C.Int.value, expr(new S.PrimOp(new S.Op2(new S.OpIntNum(kind), a, b))));
const add = (a,b) => op(S.OpAdd.value, a,b);
const sub = (a,b) => op(S.OpSubtract.value, a,b);
const condition = value => typed(C.Boolean.value, expr(new S.App(variable('zero', 'Foreign'), [value])));
const branch = (cond, yes, no) => typed(C.Int.value, expr(new S.Branch([new S.Pair(cond, yes)], no)));
const params = names => names.map(([name,level]) => new Tuple(`${name}_${level}`, G.TypeInt64.value));
const baseContext = {modNameStr:'Renamed', codegenStateRef:{}, bound:emptyMap, depth:0,
  tcoIdent:Nothing.value, loopCtx:[], options:{isTail:false,inEffectBlock:false},
  mbExpectedExprType:Nothing.value, mbExpectedGoType:Nothing.value};

function unwrap(value) {
  while (value.value1 instanceof S.Typed || value.value1 instanceof S.TypeApp) {
    value = value.value1 instanceof S.Typed ? value.value1.value1 : value.value1.value0;
  }
  return value.value1;
}
const go = value => {
  const syntax = unwrap(value);
  if (syntax instanceof S.Lit) return new G.GoInt(syntax.value0.value0);
  if (syntax instanceof S.Local) return new G.GoVar(`${syntax.value0.value0}_${syntax.value1}`);
  if (syntax instanceof S.Var) return new G.GoVar(syntax.value0.value1);
  if (syntax instanceof S.App) return new G.GoCall(go(syntax.value0), syntax.value1.map(go));
  if (syntax instanceof S.PrimOp) {
    const op = syntax.value0;
    const name = op.value0.value0 instanceof S.OpAdd ? 'IntAdd' : 'IntSub';
    return new G.GoCall(new G.GoSelector(new G.GoVar('gopurs_runtime'), name), [go(op.value1),go(op.value2)]);
  }
  throw new Error(`unsupported test translator ${syntax.constructor.name}`);
};
const translate = _context => nextId => value => ({stmts:StmtEmpty.value,expr:go(value),
  exprType: value.value1 instanceof S.Typed && value.value1.value0 === C.Boolean.value ? G.TypeBool.value : G.TypeInt64.value,
  nextId});
function lower(body, options={}) {
  return worker(translate)(options.context ?? baseContext)(options.name ?? 'walk')(
    options.params ?? params([['n',0]]))(options.result ?? G.TypeInt64.value)(body);
}
const render = (body, options={}) => {
  const result = lower(body,options);
  assert.ok(result instanceof Just, 'expected tail-add loop');
  return printGoExpr(result.value0);
};
const recurrence = (right=call('walk',[sub(local('n'),lit(1))]), left=local('n')) =>
  branch(condition(local('n')),lit(7),add(left,right));

test('lowers renamed saturated global calls and retains a flag for noncanonical base values', () => {
  for (const owner of ['Renamed',null]) {
    const original=recurrence(call('walk',[sub(local('n'),lit(1))],owner));
    const saved=JSON.stringify(original), printed=render(original);
    assert.equal(JSON.stringify(original),saved);
    assert.match(printed,/__gopurs_tail_add_0_used := false/);
    assert.match(printed,/__gopurs_tail_add_0_used = true/);
    assert.match(printed,/continue __gopurs_tail_add_0_loop/);
    assert.match(printed,/if __gopurs_tail_add_0_used/);
    assert.match(printed,/return __gopurs_tail_add_0_result_/);
  }
});

test('requires the same module, binding, native result and exact arity', () => {
  for (const wrong of [call('different',[local('n')]),call('walk',[local('n')],'Other'),
    call('walk',[]),call('walk',[local('n'),lit(1)]),typed(C.Int.value,expr(new S.App(local('walk',9),[local('n')]))),
    add(call('walk',[local('n')]),lit(1))]) {
    assert.ok(lower(recurrence(wrong)) instanceof Nothing);
  }
  assert.ok(lower(recurrence(),{result:G.TypeFloat64.value}) instanceof Nothing);
  assert.ok(lower(recurrence(),{params:[]}) instanceof Nothing);
});

test('only addition in the final right position is admitted; other syntax keeps its path', () => {
  assert.ok(lower(branch(condition(local('n')),lit(0),sub(local('n'),call('walk',[local('n')])))) instanceof Nothing);
  assert.ok(lower(branch(condition(local('n')),lit(0),add(call('walk',[local('n')]),lit(1)))) instanceof Nothing);
  assert.ok(lower(expr(new S.Let(new Just('walk'),0,lit(1),recurrence()))) instanceof Nothing);
});

test('nested branches preserve their condition ordering and all arguments are captured first', () => {
  const body=branch(condition(local('n')),lit(0),branch(condition(sub(local('n'),lit(1))),lit(9),
    add(call('left',[local('n')],'Foreign'),call('walk',[sub(local('n'),lit(1)),local('b',2),local('a',1)]))));
  const printed=render(body,{params:params([['n',0],['a',1],['b',2]])});
  assert.ok(printed.indexOf('zero(n_0)') < printed.indexOf('zero(gopurs_runtime.IntSub'));
  assert.ok(printed.indexOf('left(n_0)') < printed.indexOf('__gopurs_tail_add_0_arg_'));
  const firstMutation=printed.indexOf('n_0_loop =');
  assert.ok(firstMutation > printed.lastIndexOf(' := int64(', firstMutation));
  assert.match(printed,/a_1_loop = __gopurs_tail_add_0_arg_/);
  assert.match(printed,/b_2_loop = __gopurs_tail_add_0_arg_/);
});

test('private names avoid source parameters and unused local bindings', () => {
  const source=recurrence();
  const printed=render(source,{params:[new Tuple('__gopurs_tail_add_0_acc_0',G.TypeInt64.value)]});
  assert.match(printed,/__gopurs_tail_add_1_used/);
  const opaqueBase=expr(new S.Let(new Just('__gopurs_tail_add_0_result'),4,lit(1),lit(7)));
  const fake=context=>nextId=>value=>unwrap(value) instanceof S.Let
    ? {stmts:StmtEmpty.value,expr:new G.GoInt(7),exprType:G.TypeInt64.value,nextId}
    : translate(context)(nextId)(value);
  const wrapped=branch(condition(local('n')),opaqueBase,add(lit(1),call('walk',[sub(local('n'),lit(1))])));
  const result=worker(fake)(baseContext)('walk')(params([['n',0]]))(G.TypeInt64.value)(wrapped);
  assert.match(printGoExpr(result.value0),/__gopurs_tail_add_1_used/);
});

test('emitted Go preserves Int64 base values, Int32 overflow, argument permutation and callback order', () => {
  const order=branch(condition(local('n')),call('base',[local('seed',1)],'Foreign'),
    add(call('left',[local('n')],'Foreign'),call('walk',[call('next',[sub(local('n'),lit(1))],'Foreign'),local('seed',1)])));
  const swap=branch(condition(local('n')),local('a',1),add(local('b',2),
    call('permute',[sub(local('n'),lit(1)),local('b',2),add(local('a',1),lit(1))])));
  const path=mkdtempSync(join(tmpdir(),'gopurs-tail-add-'));
  try {
    const source=`package main
import("encoding/json";"os")
var trace []int64
var gopurs_runtime=struct{IntAdd,IntSub func(int64,int64)int64}{func(a,b int64)int64{return int64(int32(a+b))},func(a,b int64)int64{return int64(int32(a-b))}}
func zero(n int64)bool{return n==0}
func left(n int64)int64{trace=append(trace,100+n);return 2147483647}
func next(n int64)int64{trace=append(trace,200+n);return n}
func base(n int64)int64{trace=append(trace,300);return n}
func walk(n_0_loop,seed_1_loop int64)int64 {${render(order,{params:params([['n',0],['seed',1]])})}}
func permute(n_0_loop,a_1_loop,b_2_loop int64)int64 {${render(swap,{name:'permute',params:params([['n',0],['a',1],['b',2]])})}}
func original(n,a,b int64)int64{if n==0{return a};return gopurs_runtime.IntAdd(b,original(n-1,b,gopurs_runtime.IntAdd(a,1)))}
func main(){
 enc:=json.NewEncoder(os.Stdout)
 for _,seed:=range []int64{0,2147483647,-2147483648,1099511627779,-1099511627779}{
  for _,n:=range []int64{0,1,3}{trace=nil;v:=walk(n,seed);_ = enc.Encode([]any{n,seed,v,trace})}
 }
 for n:=int64(0);n<10;n++{for _,a:=range []int64{3,2147483647,1099511627779}{if permute(n,a,-17)!=original(n,a,-17){panic("permutation")}}}
}
`;
    writeFileSync(join(path,'main.go'),source);
    const run=spawnSync('go',['run','main.go'],{cwd:path,encoding:'utf8',timeout:30000});
    assert.equal(run.status,0,run.stdout+run.stderr);
    const rows=run.stdout.trim().split('\n').map(JSON.parse);
    assert.equal(rows.length,15);
    for (const [n,seed,result,trace] of rows) {
      let wanted=BigInt(seed);
      for(let i=0;i<n;i++)wanted=BigInt.asIntN(32,wanted+2147483647n);
      assert.equal(result,Number(wanted));
      assert.deepEqual(trace,Array.from({length:n},(_,i)=>[100+n-i,200+n-i-1]).flat().concat(300));
    }
  } finally {rmSync(path,{recursive:true,force:true});}
});
