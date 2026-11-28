import assert from 'node:assert/strict';
import {spawnSync} from 'node:child_process';
import {mkdirSync,mkdtempSync,rmSync,writeFileSync} from 'node:fs';
import {tmpdir} from 'node:os';
import {join} from 'node:path';
import test from 'node:test';
import {empty as emptyMap} from '../output/Data.Map/index.js';
import {Just,Nothing} from '../output/Data.Maybe/index.js';
import {empty as emptySet} from '../output/Data.Set/index.js';
import {Tuple} from '../output/Data.Tuple/index.js';
import * as CodeGen from '../output/Gopurs.CodeGen/index.js';
import {optimize,describe,validDescriptor} from '../output/Gopurs.CodePointFold/index.js';
import {runtimeGoCode} from '../output/Gopurs.Runtime.Source/index.js';
import {runtimeSimdFiles} from '../output/Gopurs.Runtime.Simd/index.js';
import {TcoExpr} from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as C from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import {NeutralExpr} from '../output/PureScript.Backend.Optimizer.Semantics/index.js';
import * as S from '../output/PureScript.Backend.Optimizer.Syntax/index.js';
import {withReboxFields} from './codegen-metadata.mjs';
const e=syntax=>new NeutralExpr(syntax);
const typed=(ty,x)=>e(new S.Typed(ty,x));
const int=x=>e(new S.Lit(new C.LitInt(x)));
const local=(level,ty)=>{const x=e(new S.Local(new Just('x'),level));return ty?typed(ty,x):x;};
const ref=(module,name)=>e(new S.Var(new C.Qualified(new Just(module),name)));
const app=(fn,...args)=>e(new S.App(fn,args));
const abs=(levels,body)=>e(new S.Abs(levels.map(l=>new Tuple(new Just('x'),l)),body));
const operation=(op,a,b)=>typed(C.Int.value,e(new S.PrimOp(new S.Op2(new S.OpIntNum(op),a,b))));
const add=(a,b)=>operation(S.OpAdd.value,a,b);
const array=new C.Array(C.Int.value), cbtype=new C.Func([C.Int.value,C.Int.value],C.Int.value);
const sum=typed(cbtype,abs([0,1],add(local(0),local(1))));
const fold=ref('Data.Foldable','foldlArray'),convert=ref('Data.String.CodePoints','toCodePointArray');
const stringFold=new C.Func([C.String.value],C.Int.value);
const appliedFunction=(arg,ret)=>new C.TypeApp(new C.ADT("Prim.Function",["Prim","Function"],[]),[arg,ret]);
const compose=(callback=sum,seed=int(0),conversion=convert,dictionary=ref('Control.Semigroupoid','semigroupoidFn'))=>typed(stringFold,app(ref('Control.Semigroupoid','compose'),dictionary,app(fold,callback,seed),conversion));
const moduleOf=(bindings,recursive=false)=>({name:'CodePointFixture',bindings:[{recursive,bindings:bindings.map(([n,x])=>new Tuple(n,x))}],comments:[],imports:emptySet,exports:emptySet,reExports:emptySet,dataTypes:emptyMap,dataDecls:[],classDecls:[],foreign:emptyMap,implementations:emptyMap,directives:emptyMap});
const transformed=(x,recursive=false)=>optimize(moduleOf([['fold',x]],recursive)).bindings[0].bindings[0].value1;
const metadata=withReboxFields({elidedCtors:emptySet,ctorTypes:emptyMap,pointerAdtPaths:emptyMap,pointerAdtNodes:emptySet,pointerAdtLeaves:emptyMap,enumAdts:emptySet,enumCtors:emptySet,globalTypes:emptyMap,globalFunctions:emptyMap,classDeclsFields:emptyMap});

import {emit} from '../output/Gopurs.StringConcatFold/index.js';
import {analyze as analyzeTco} from '../output/PureScript.Backend.Optimizer.Codegen.Tco/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import {StmtEmpty,flattenStmts} from '../output/Gopurs.ExprContext/index.js';
import {printGoExpr} from '../output/Gopurs.Printer/index.js';
const analyze=x=>analyzeTco([])(x),strings=new C.Array(C.String.value),append=ref('Data.Semigroup','concatString');
const concat=(initial=local(3,C.String.value),input=local(4,strings),callback=append,head=fold)=>typed(C.String.value,app(head,callback,initial,input));

test('concat fusion requires canonical saturation, total closed callbacks and compatible nested types',()=>{
 const never=()=>{throw Error('rejected candidate translated');};
 const reject=(p,d={callback:sum,seed:0},recVars=[])=>assert.deepEqual(emit(never)({recVars})(0)(d)(analyze(p)),Nothing.value);
 reject(concat(undefined,undefined,ref('Custom','concatString')));reject(concat(undefined,undefined,append,ref('Custom','foldlArray')));
 reject(app(fold,append,local(3,C.String.value)));reject(typed(C.String.value,local(3,C.String.value)));
 reject(concat(undefined,undefined,app(ref('Foreign','factory'),int(1))));
 reject(concat(undefined,undefined,append,e(new S.UncurriedApp(fold,[]))));
 reject(concat(local(3,C.Int.value)));reject(concat(undefined,local(4,array)));
 reject(typed(C.Int.value,concat()));reject(typed(C.String.value,typed(C.Int.value,concat())));
 reject(concat(undefined,undefined,typed(new C.Func([C.String.value,C.String.value],C.String.value),typed(C.Int.value,append))));
 reject(concat(undefined,undefined,append,typed(C.Int.value,fold)));
 reject(concat(),{callback:sum,seed:0},['recursive']);
 for(const callback of [typed(cbtype,abs([0,1],add(local(0),local(9,C.Int.value)))),typed(cbtype,abs([0,1],app(ref('Foreign','call'),local(1)))),typed(cbtype,abs([0,1],operation(S.OpDivide.value,local(0),local(1)))),typed(cbtype,abs([0],typed(new C.Func([C.String.value],C.Int.value),abs([1],add(local(0),local(1))))))])reject(concat(),{callback,seed:0});
});

function generated(){
 const consumer=(name,callback,seed)=>[name,compose(callback,int(seed))];
 const caller=name=>[name+'Concat',typed(new C.Func([C.String.value,strings],C.Int.value),abs([3,4],typed(C.Int.value,app(ref('CodePointFixture',name),concat()))))];
 const weighted=typed(cbtype,abs([0,1],add(operation(S.OpMultiply.value,local(0),int(33)),local(1))));
 const subtract=typed(cbtype,abs([0,1],operation(S.OpSubtract.value,local(0),local(1))));
 const wideProduct=typed(cbtype,abs([0,1],add(operation(S.OpMultiply.value,local(0),int(2147483647)),local(1))));
 return CodeGen.translate(metadata)(moduleOf([consumer('sum',sum,0),consumer('weighted',weighted,17),consumer('wide',sum,4294967295),consumer('subtract',subtract,-2147483648),consumer('wideProduct',wideProduct,2147483647),...['sum','weighted','wide','subtract','wideProduct'].map(caller)]));
}
function runGo(t,code,main){const dir=mkdtempSync(join(tmpdir(),'gopurs-concat-fold-'));t.after(()=>rmSync(dir,{recursive:true,force:true}));mkdirSync(join(dir,'gopurs_runtime'));writeFileSync(join(dir,'go.mod'),'module gopurs/output\n\ngo 1.22\n');writeFileSync(join(dir,'gopurs_runtime/runtime.go'),runtimeGoCode);for(const f of runtimeSimdFiles)writeFileSync(join(dir,'gopurs_runtime',f.name),f.content);writeFileSync(join(dir,'fixture.go'),code.replace('package purescript','package main'));writeFileSync(join(dir,'main.go'),main);const result=spawnSync('go',['run','.'],{cwd:dir,encoding:'utf8',timeout:60000,env:{...process.env,GOWORK:'off'}});assert.ifError(result.error);assert.equal(result.status,0,result.stdout+result.stderr);return result.stdout;}

test('generated suffix continuation agrees with complete concatenation and the original consumers',t=>{
 const code=generated();assert.match(code,/concatFold_finish/);assert.doesNotMatch(code,/Get_Data_Foldable_foldlArray|Get_Data_String_CodePoints_toCodePointArray/);
 const result=runGo(t,code,`package main
import("fmt";"math/rand";r "gopurs/output/gopurs_runtime")
func Get_Data_Semigroup_concatString()r.Value{return r.Func2(func(a,b r.Value)r.Value{return r.Str(r.ConcatString(a.StrVal(),b.StrVal()))})}
func boxed(parts []string)r.Value{values:=make([]r.Value,len(parts));for i,s:=range parts{values[i]=r.Str(s)};return r.Box(values)}
func main(){samples:=[]string{"","ab","é","𝄞","\\xed\\xa0\\x80","\\xed\\xb0\\x80","\\xed","\\xa0","\\x80","\\xf0\\x9f","\\xff","\\xc0\\x80"};checks:=0
check:=func(initial string,parts []string){joined:=r.ConcatStrings(initial,parts,func(x string)string{return x});pairs:=[][2]int64{
{Call_CodePointFixture_sumConcat(initial,parts),Call_CodePointFixture_sum(joined)},
{Call_CodePointFixture_weightedConcat(initial,parts),Call_CodePointFixture_weighted(joined)},
{Call_CodePointFixture_wideConcat(initial,parts),Call_CodePointFixture_wide(joined)},
{Call_CodePointFixture_subtractConcat(initial,parts),Call_CodePointFixture_subtract(joined)},
{Call_CodePointFixture_wideProductConcat(initial,parts),Call_CodePointFixture_wideProduct(joined)},
{r.Apply2(Get_CodePointFixture_weightedConcat(),r.Str(initial),boxed(parts)).IntVal,Call_CodePointFixture_weighted(joined)},
};for _,pair:=range pairs{if pair[0]!=pair[1]{panic(fmt.Sprintf("%q %q: %v",initial,parts,pair))};checks++}}
for _,initial:=range samples{check(initial,nil);check(initial,[]string{});for _,a:=range samples{for _,b:=range samples{check(initial,[]string{a,"",b})}}}
rng:=rand.New(rand.NewSource(133));for n:=0;n<250;n++{parts:=make([]string,rng.Intn(10));for i:=range parts{bytes:=make([]byte,rng.Intn(20));rng.Read(bytes);parts[i]=string(bytes)};check("ASCII prefix",parts)}
for _,s:=range []string{"\\xed\\xa0\\x80\\xed\\xb0\\x80","\\xf0\\x9d\\x84\\x9e","\\xff\\xed\\xa0\\x80\\xc0\\x80","ASCII\\xed\\xa0\\x80\\xed\\xb0\\x80tail"}{for a:=0;a<=len(s);a++{for b:=a;b<=len(s);b++{check("",[]string{s[:a],"",s[a:b],s[b:]})}}}
fmt.Println(checks,"generated native/public suffix checks passed")}
`);assert.match(result,/generated native\/public suffix checks passed/);console.log(result.trim());
});

test('producer callback, initial factory and input factory are captured once in order before reads',t=>{
 const context={metadata,modNameStr:'Capture',depth:0,bound:emptyMap,recVars:[],loopCtx:[],tcoIdent:Nothing.value,
  options:{isTail:false,inEffectBlock:false},mbExpectedExprType:Nothing.value,mbExpectedGoType:Nothing.value,codegenStateRef:{value:{declarations:[],globalId:0,reboxPairs:emptySet}}};
 const strip=x=>{let n=x.value1;while(n instanceof S.Typed)n=n.value1.value1;return n;};
 // This leaf translator deliberately returns the current code point; 81/90
 // observe which fragment was read, while checksum arithmetic is tested above.
 const translator=()=>id=>value=>{const n=strip(value);let expr,exprType;if(n instanceof S.PrimOp){const cp=strip(n.value0.value2);expr=Go.rawGo('x_'+cp.value1);exprType=Go.TypeInt64.value;}else if(n instanceof S.Lit){expr=Go.rawGo('int64('+n.value0.value0+')');exprType=Go.TypeInt64.value;}else if(n instanceof S.Var){expr=Go.rawGo('concatFactory()');exprType=Go.TypeValue.value;}else if(n instanceof S.Local&&n.value1===3){expr=Go.rawGo('initialFactory()');exprType=Go.TypeString.value;}else {expr=Go.rawGo('inputFactory()');exprType=new Go.TypeNativeArray(Go.TypeString.value);}return {stmts:StmtEmpty.value,expr,exprType,nextId:id+1};};
 const result=emit(translator)(context)(0)({callback:sum,seed:7})(analyze(concat()));assert.ok(result instanceof Just);
 const body=flattenStmts(result.value0.stmts).map(printGoExpr).join('\n')+'\nreturn '+printGoExpr(result.value0.expr);
 runGo(t,'package main\nimport gopurs_runtime "gopurs/output/gopurs_runtime"\nfunc run()int64{'+body+'}\n',`package main
import("fmt";r "gopurs/output/gopurs_runtime")
var trace string
var parts=[]string{"before"}
var emptyParts,emptyInitial bool
var fail string
var marker=&struct{}{}
func concatFactory()r.Value{trace+="C";return r.Value{}}
func initialFactory()string{trace+="I";if fail=="I"{panic(marker)};parts[0]="Q";if emptyInitial{return ""};return "Z"}
func inputFactory()[]string{trace+="P";if fail=="P"{panic(marker)};if emptyParts{return nil};return parts}
func main(){for _,c:=range []struct{emptyParts,emptyInitial bool;want int64}{{false,false,81},{true,false,90},{true,true,7}}{trace="";emptyParts=c.emptyParts;emptyInitial=c.emptyInitial;if result:=run();result!=c.want||trace!="CIP"{panic(fmt.Sprint(result,trace))}};for _,stage:=range []string{"I","P"}{fail=stage;trace="";func(){defer func(){if recover()!=marker{panic("wrong panic")}}();run();panic("missing panic")}();want:="CI";if stage=="P"{want="CIP"};if trace!=want{panic(trace)}};fmt.Println("capture order, empty array, factory panics and alias mutation: ok")}
`);
});
