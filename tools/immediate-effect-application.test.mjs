import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { executeIfOpaque } from '../output/Gopurs.ExprAnalysis/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';
import { TcoExpr } from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';

const opaque = new TcoExpr(null, new S.Local(null, 0));
const runtimeCall = (name, args) => new Go.GoCall(new Go.GoSelector(new Go.GoVar('gopurs_runtime'), name), args);
const arg = name => Go.rawGo(`capture("${name}", ${name})`);
const unit = Go.rawGo('gopurs_runtime.Value{}');
const render = expr => printGoExpr(executeIfOpaque(opaque)(expr));

test('only immediate exact runtime applications consume the extra unit argument', () => {
    assert.equal(render(runtimeCall('Apply', [arg('f'), arg('a')])),
        printGoExpr(runtimeCall('Apply2', [arg('f'), arg('a'), unit])));
    assert.equal(render(runtimeCall('Apply2', [arg('f'), arg('a'), arg('b')])),
        printGoExpr(runtimeCall('Apply3', [arg('f'), arg('a'), arg('b'), unit])));
    for (const expr of [new Go.GoVar('savedAction'),
        new Go.GoCall(new Go.GoSelector(new Go.GoVar('other'), 'Apply2'), [arg('f'), arg('a'), arg('b')]),
        runtimeCall('Apply2', [arg('f'), arg('a')]), runtimeCall('Apply3', [arg('f'), arg('a'), arg('b'), unit])]) {
        assert.equal(render(expr), printGoExpr(runtimeCall('Apply', [expr, unit])));
    }
    const alreadyExecuted = new TcoExpr(null, new S.EffectPure(opaque));
    const call = runtimeCall('Apply2', [arg('f'), arg('a'), arg('b')]);
    assert.equal(executeIfOpaque(alreadyExecuted)(call), call);
});

test('immediate saturation preserves capture order, currying, replay and exceptions', t => {
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-immediate-effect-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    const cases = [1, 2].map(arity => {
        const args = [arg('f'), arg('a'), ...(arity === 2 ? [arg('b')] : [])];
        return `func execute${arity}(f,a,b gopurs_runtime.Value)gopurs_runtime.Value{return ${render(runtimeCall(arity === 1 ? 'Apply' : 'Apply2', args))}}`;
    });
    writeFileSync(join(directory, 'main.go'), `package main
import("fmt";"reflect";"gopurs/output/gopurs_runtime")
var trace []string
func capture(label string,v gopurs_runtime.Value)gopurs_runtime.Value{trace=append(trace,label);return v}
${cases.join('\n')}
func main(){
 for _,arity:=range []int{1,2}{for _,curried:=range []bool{false,true}{
  calls:=0
  end:=func(_ gopurs_runtime.Value)gopurs_runtime.Value{calls++;trace=append(trace,"effect");return gopurs_runtime.Int(int64(calls))}
  var f gopurs_runtime.Value
  if arity==1 {if curried {f=gopurs_runtime.Func(func(_ gopurs_runtime.Value)gopurs_runtime.Value{trace=append(trace,"partial");return gopurs_runtime.Func(end)})}else{f=gopurs_runtime.Func2(func(_,u gopurs_runtime.Value)gopurs_runtime.Value{return end(u)})}}else{
   if curried {f=gopurs_runtime.Func2(func(_,_ gopurs_runtime.Value)gopurs_runtime.Value{trace=append(trace,"partial");return gopurs_runtime.Func(end)})}else{f=gopurs_runtime.Func3(func(_,_,u gopurs_runtime.Value)gopurs_runtime.Value{return end(u)})}
  }
  for replay:=1;replay<=2;replay++{trace=nil;var result gopurs_runtime.Value;if arity==1{result=execute1(f,gopurs_runtime.Int(1),gopurs_runtime.Int(2))}else{result=execute2(f,gopurs_runtime.Int(1),gopurs_runtime.Int(2))}
   want:=[]string{"f","a"};if arity==2{want=append(want,"b")};if curried{want=append(want,"partial")};want=append(want,"effect")
   if !reflect.DeepEqual(trace,want)||result.IntVal!=int64(replay){panic(fmt.Sprint(trace,result.IntVal))}
  }
 }}
 marker:=&struct{}{};func(){defer func(){if recover()!=marker{panic("exception identity")}}();execute2(gopurs_runtime.Func3(func(_,_,_ gopurs_runtime.Value)gopurs_runtime.Value{panic(marker)}),gopurs_runtime.Value{},gopurs_runtime.Value{})}()
}`);
    const result = spawnSync('go', ['run', '.'], { cwd: directory, encoding: 'utf8' });
    assert.equal(result.status, 0, result.stderr || result.stdout);
});
