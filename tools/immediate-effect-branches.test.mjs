import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { empty as emptyMap } from '../output/Data.Map/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { StmtEmpty, StmtLeaf, flattenStmts } from '../output/Gopurs.ExprContext/index.js';
import { emit } from '../output/Gopurs.ImmediateEffectBranches/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';
import { TcoExpr } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';

const e=s=>new TcoExpr(null,s), token=s=>e(new S.Lit(new C.LitString(s))), pure=n=>e(new S.EffectPure(e(new S.Lit(new C.LitInt(n)))));
const branch=(condition,yes,no)=>e(new S.Branch([new S.Pair(token(condition),yes)],no));
const context=()=>({modNameStr:'ImmediateBranches',metadata:{},depth:0,bound:emptyMap,tcoIdent:Nothing.value,loopCtx:[],
  options:{isTail:true,inEffectBlock:true},mbExpectedExprType:Nothing.value,mbExpectedGoType:Nothing.value,
  codegenStateRef:{value:{declarations:[],globalId:0,reboxPairs:emptySet}}});
const translate=ctx=>id=>expression=>{
  const node=expression.value1;
  let expr,exprType=Go.TypeValue.value,stmts=StmtEmpty.value;
  if(node instanceof S.EffectPure){assert.equal(ctx.options.inEffectBlock,true);assert.equal(ctx.options.isTail,false);const n=node.value0.value1.value0.value0;expr=Go.rawGo(`int64(${n})`);exprType=Go.TypeInt64.value;stmts=new StmtLeaf(Go.rawGo(`trace=append(trace,"pure${n}")`));}
  else if(node instanceof S.Let||node instanceof S.LetRec){throw new Error('Let context uses its separate translator');}
  else if(node instanceof S.EffectDefer){assert.equal(ctx.options.inEffectBlock,false);expr=Go.rawGo('factory("defer")');}
  else {assert.equal(ctx.options.inEffectBlock,false);const label=node.value0.value0;if(label.startsWith('cond')){expr=Go.rawGo(`condition("${label}")`);exprType=Go.TypeBool.value;}else expr=Go.rawGo(`factory("${label}")`);}
  return {expr,exprType,stmts,nextId:id+1};
};
// Keep the fake leaf translator small: Let/LetRec tests below inspect contexts
// separately; Go execution uses primitive and opaque leaves.
const render=expression=>{const result=emit(translate)(context())(0)(expression);assert.ok(result instanceof Just);return flattenStmts(result.value0.stmts).map(printGoExpr).join('\n')+'\nreturn '+printGoExpr(result.value0.expr);};

test('only immediate branch roots with known effects are admitted; defer and opaque leaves keep Apply',()=>{
  assert.equal(emit(translate)(context())(0)(pure(1)),Nothing.value);
  assert.equal(emit(translate)(context())(0)(branch('cond0',token('left'),token('right'))),Nothing.value);
  assert.equal(emit(translate)(context())(0)(branch('cond0',e(new S.EffectDefer(token('x'))),token('right'))),Nothing.value);
  const code=render(branch('cond0',pure(1),e(new S.EffectDefer(token('x')))));
  assert.match(code,/Apply\(factory\("defer"\), gopurs_runtime.Value\{\}\)/);
  assert.doesNotMatch(code,/gopurs_runtime.Func\(/);
  for(const expression of [e(new S.Let(new Just('x'),0,token('ignored'),pure(3))),e(new S.LetRec(0,[new Tuple('x',token('ignored'))],pure(3)))]){
    let seen=false;
    const inspect=ctx=>id=>node=>{if(node===expression){seen=true;assert.deepEqual(ctx.options,{isTail:false,inEffectBlock:true});return {stmts:StmtEmpty.value,expr:Go.rawGo('int64(3)'),exprType:Go.TypeInt64.value,nextId:id+1};}return translate(ctx)(id)(node);};
    assert.ok(emit(inspect)(context())(0)(branch('cond0',expression,token('saved'))) instanceof Just);assert.ok(seen);
  }
});

test('generated branches preserve nested selection, one factory/effect, replay and panic order',t=>{
  const directory=mkdtempSync(join(tmpdir(),'gopurs-effect-branches-'));t.after(()=>rmSync(directory,{recursive:true,force:true}));mkdirSync(join(directory,'gopurs_runtime'));
  writeFileSync(join(directory,'go.mod'),'module gopurs/output\n\ngo 1.22\n');writeFileSync(join(directory,'gopurs_runtime/runtime.go'),runtimeGoCode);
  const code=render(branch('cond0',pure(11),branch('cond1',pure(22),token('saved'))));
  const deferred=render(branch('cond0',pure(11),e(new S.EffectDefer(token('x')))));
  writeFileSync(join(directory,'branches_test.go'),`package branches
import("testing";"reflect";r "gopurs/output/gopurs_runtime";gopurs_runtime "gopurs/output/gopurs_runtime")
var trace []string
var flags [2]bool
var marker=&struct{}{}
var throws bool
func condition(name string)bool{trace=append(trace,name);if throws{panic(marker)};if name=="cond0"{return flags[0]};return flags[1]}
func factory(name string)r.Value{trace=append(trace,"factory:"+name);return r.Func(func(r.Value)r.Value{trace=append(trace,"effect:"+name);return r.Int(33)})}
func run()r.Value{${code}}
func deferred()r.Value{${deferred}}
func TestOrdering(t *testing.T){cases:=[]struct{flags [2]bool;value int64;trace []string}{
 {[2]bool{true,false},11,[]string{"cond0","pure11"}},
 {[2]bool{false,true},22,[]string{"cond0","cond1","pure22"}},
 {[2]bool{false,false},33,[]string{"cond0","cond1","factory:saved","effect:saved"}},
 };for _,c:=range cases{flags=c.flags;for i:=0;i<3;i++{trace=nil;got:=run();if got.IntVal!=c.value||!reflect.DeepEqual(trace,c.trace){t.Fatalf("%v: %v %v",flags,got,trace)}}}
 flags=[2]bool{};trace=nil;if deferred().IntVal!=33||!reflect.DeepEqual(trace,[]string{"cond0","factory:defer","effect:defer"}){t.Fatal(trace)}
 throws=true;trace=nil;func(){defer func(){if recover()!=marker{t.Fatal("panic identity")}}();run();t.Fatal("missing panic")}();if !reflect.DeepEqual(trace,[]string{"cond0"}){t.Fatal(trace)}
}
`);
  const result=spawnSync('go',['test','.'],{cwd:directory,encoding:'utf8',env:{...process.env,GOWORK:'off'},timeout:120000});assert.ifError(result.error);assert.equal(result.status,0,result.stdout+result.stderr);
});
