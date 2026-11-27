import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join, resolve } from 'node:path';
import { fileURLToPath, pathToFileURL } from 'node:url';
import test from 'node:test';

const output = process.argv[2] ?? fileURLToPath(new URL('../output/', import.meta.url));
const load = name => import(pathToFileURL(resolve(output, name, 'index.js')));
const [Go, A, C, S, T, Maybe, M, Ord, Tuples, Free, Printer, Context] = await Promise.all([
    'Gopurs.GoAst', 'Gopurs.AffineLoops', 'PureScript.Backend.Optimizer.CoreFn',
    'PureScript.Backend.Optimizer.Syntax', 'PureScript.Backend.Optimizer.Codegen.Tco',
    'Data.Maybe', 'Data.Map', 'Data.Ord', 'Data.Tuple',
    'PureScript.Backend.Optimizer.FreeVars', 'Gopurs.Printer', 'Gopurs.ExprContext',
].map(load));
const { Just, Nothing } = Maybe, { Tuple } = Tuples;
const int = C.Int.value, i64 = Go.TypeInt64.value;
const expr = syntax => new T.TcoExpr(null, syntax);
const typed = (type, value) => expr(new S.Typed(type, value));
const local = (name, level) => typed(int, expr(new S.Local(new Just(name), level)));
const counter = local('remaining', 0), accumulator = local('seed', 1);
const counterName = Free.localId(new Just('remaining'))(0), accumulatorName = Free.localId(new Just('seed'))(1);
const params = [new Tuple(counterName, i64), new Tuple(accumulatorName, i64)];
const integer = n => typed(int, expr(new S.Lit(new C.LitInt(n))));
const variable = (name = 'repeat', owner = 'Probe') => expr(new S.Var(new C.Qualified(new Just(owner), name)));
const binary = (operator, left, right, type = int) => typed(type, expr(new S.PrimOp(new S.Op2(operator, left, right))));
const add = (left, right) => binary(new S.OpIntNum(S.OpAdd.value), left, right);
const subtract = (left, right) => binary(new S.OpIntNum(S.OpSubtract.value), left, right);
const modulo = (left, right) => binary(new S.OpIntNum(S.OpMod.value), left, right);
const condition = binary(new S.OpIntOrd(S.OpEq.value), counter, integer(0), C.Boolean.value);
const call = (target, args) => typed(int, expr(new S.App(target, args)));
const sentinel = new Go.GoReturn(new Go.GoInt(-98));
const self = A.globalTarget('Probe')('repeat');
function fixture(options = {}) {
    const conditionExpr = options.condition ?? condition;
    const decrement = options.decrement ?? subtract(counter, integer(1));
    const update = options.update ?? add(accumulator, integer(options.increment ?? 1));
    const step = options.step ?? call(options.target ?? variable(), [decrement, update]);
    return typed(int, expr(new S.Branch([new S.Pair(conditionExpr, options.seed ?? accumulator)], step)));
}
const readContext = { depth: 0, mbExpectedExprType: new Just(int), mbExpectedGoType: new Just(i64),
    options: { isTail: true, inEffectBlock: true }, tcoIdent: new Just('repeat'), loopCtx: [], recVars: [] };
const noReads = () => () => () => { throw new Error('unexpected global read'); };
function optimize(body, parameters = params, result = i64, target = self, fallback = sentinel, translate = noReads) {
    const before = JSON.stringify(body);
    const value = A.wrap(translate)(readContext)(target)(parameters)(result)(body)(fallback);
    assert.equal(JSON.stringify(body), before, 'input IR remains immutable');
    return value;
}

test('affine admission proves the entire recurrence and the exact self target', () => {
    for (const increment of [-2147483648, -7, 0, 1, 7, 2147483647]) {
        const body = fixture({ increment });
        const result = optimize(body);
        assert.notEqual(result, sentinel);
        const code = Printer.printGoExpr(result);
        assert.match(code, /int64\(int32\(/);
        assert.match(code, /\*.*int64\(/);
        assert.doesNotMatch(code, /IntMul/);
        assert.match(code, /2147483647/);
        assert.equal(result.value0.at(-1), sentinel, 'fallback is unchanged');
    }
    const swapped = fixture({ condition: binary(new S.OpIntOrd(S.OpEq.value), integer(0), counter, C.Boolean.value),
        update: add(integer(-3), accumulator) });
    assert.notEqual(optimize(swapped), sentinel);
    const localSelf = local('nested', 7), localKey = Free.localId(new Just('nested'))(7);
    const bound = M.insert(Ord.ordString)(localKey)({ name: 'allocated_worker', goType: i64 })(M.empty);
    assert.notEqual(optimize(fixture({ target: localSelf }), params, i64, A.localTarget(bound)('allocated_worker')), sentinel);
    assert.equal(optimize(fixture({ target: localSelf }), params, i64, A.localTarget(bound)('another_worker')), sentinel);
});

test('affine admission rejects calls, effects, changing increments and altered recurrences', () => {
    const options = [
        { target: variable('different') }, { target: variable('repeat', 'Foreign') },
        { decrement: subtract(counter, integer(2)) }, { decrement: add(counter, integer(1)) },
        { update: add(accumulator, counter) }, { update: subtract(accumulator, integer(1)) },
        { update: call(variable('observe'), [accumulator]) },
        { update: add(accumulator, integer(4294967295)) },
        { seed: counter }, { seed: add(accumulator, integer(1)) },
        { condition: binary(new S.OpIntOrd(S.OpLt.value), counter, integer(0), C.Boolean.value) },
        { step: call(variable(), [subtract(counter, integer(1)), accumulator, integer(1)]) },
        { step: expr(new S.EffectPure(integer(1))) },
    ];
    for (const option of options) assert.equal(optimize(fixture(option)), sentinel);
    const effectBefore = expr(new S.Let(Nothing.value, 9, call(variable('observe'), [accumulator]), fixture()));
    assert.equal(optimize(effectBefore), sentinel);
    assert.equal(optimize(fixture(), [...params, new Tuple('extra', i64)]), sentinel);
    assert.equal(optimize(fixture(), [new Tuple(counterName, Go.TypeFloat64.value), params[1]]), sentinel);
    assert.equal(optimize(fixture(), params, Go.TypeValue.value), sentinel);
});

test('remainder periods require the same counter and a positive canonical literal divisor', () => {
    for (const divisor of [1, 2, 3, 7, 65537, 2147483647]) {
        const result = optimize(fixture({update:add(accumulator,modulo(counter,integer(divisor)))}));
        assert.notEqual(result,sentinel);
        assert.equal(result.value0.at(-1),sentinel);
        assert.doesNotMatch(Printer.printGoExpr(result),/IntMul/);
        assert.notEqual(optimize(fixture({update:add(modulo(counter,integer(divisor)),accumulator)})),sentinel);
    }
    for (const value of [modulo(counter,integer(0)),modulo(counter,integer(-3)),
        modulo(counter,integer(2147483648)),modulo(accumulator,integer(3)),
        modulo(subtract(counter,integer(1)),integer(3)),modulo(counter,accumulator),
        modulo(counter,call(variable('divisor'),[integer(3)]))]) {
        assert.equal(optimize(fixture({update:add(accumulator,value)})),sentinel);
    }
    const stepped=letRead('first',2,call(variable(),[subtract(counter,integer(1)),
        add(accumulator,letRead('second',3,modulo(counter,integer(7))))]));
    const code=Printer.printGoExpr(optimize(fixture({step:stepped}),params,i64,self,sentinel,translateRead));
    assert.ok(code.indexOf('GetFirst()')<code.indexOf('GetSecond()'));
    assert.equal((code.match(/GetFirst\(\)/g)??[]).length,1);
    assert.equal((code.match(/GetSecond\(\)/g)??[]).length,1);
});

const selfStep = () => call(variable(), [subtract(counter, integer(1)), add(accumulator, integer(60))]);
const letRead = (name, level, body, value = variable(name, 'RenamedGlobals')) =>
    expr(new S.Let(new Just(name), level, value, body));
const twoReads = () => letRead('first', 2, letRead('second', 3, selfStep()));
const nestedReadStep = () => letRead('first', 2, call(variable(), [subtract(counter, integer(1)),
    add(accumulator, expr(new S.Let(Nothing.value, 3, variable('second', 'RenamedGlobals'), integer(60))))]));
const translateRead = context => nextId => value => {
    assert.equal(context.mbExpectedExprType, Nothing.value, 'the Int result expectation must not reach a global read');
    assert.equal(context.mbExpectedGoType, Nothing.value);
    assert.deepEqual(context.options, {isTail:false, inEffectBlock:false});
    while (value.value1 instanceof S.Typed) value = value.value1.value1;
    assert.ok(value.value1 instanceof S.Var);
    const name = value.value1.value0.value1;
    return {stmts: new Context.StmtLeaf(Go.rawGo(`trace += "${name === 'first' ? 'a' : 'b'}"`)),
        expr: new Go.GoCall(new Go.GoVar(name === 'first' ? 'GetFirst' : 'GetSecond'), []),
        exprType:i64, nextId:nextId + 1};
};

test('affine recurrence admits only unused strict global reads in the recursive branch', () => {
    const body = fixture({step:twoReads()});
    const result = optimize(body, params, i64, self, sentinel, translateRead);
    assert.notEqual(result, sentinel);
    assert.equal(result.value0.at(-1), sentinel);
    const code = Printer.printGoExpr(result);
    assert.ok(code.indexOf('>') < code.indexOf('GetFirst()'), 'positive guard precedes reads');
    assert.ok(code.indexOf('GetFirst()') < code.indexOf('GetSecond()'));
    assert.equal((code.match(/GetFirst\(\)/g) ?? []).length, 1);
    assert.equal((code.match(/GetSecond\(\)/g) ?? []).length, 1);

    const unsupported = [
        letRead('first', 2, selfStep(), call(variable('observe'), [integer(0)])),
        letRead('first', 2, selfStep(), local('captured', 9)),
        letRead('first', 2, selfStep(), expr(new S.Accessor(variable('object'), new S.GetProp('value')))),
        letRead('first', 2, selfStep(), expr(new S.EffectPure(integer(0)))),
        letRead('first', 2, call(variable(), [subtract(counter, integer(1)), add(accumulator, local('first', 2))])),
        // A rebinding of the exact counter level cannot masquerade as unused.
        letRead('remaining', 0, selfStep()),
    ];
    for (const step of unsupported) assert.equal(optimize(fixture({step})), sentinel);
    let tooMany = selfStep();
    for (let i=0; i<65; i++) tooMany = letRead('unused'+i, 20+i, tooMany);
    assert.equal(optimize(fixture({step:tooMany})), sentinel);
    assert.equal(optimize(letRead('first',2,fixture())),sentinel,'a read outside the branch cannot move under a positive guard');
});

test('affine constant operands retain strict demands after the step prefix', () => {
    const body = fixture({step:nestedReadStep()});
    const result = optimize(body, params, i64, self, sentinel, translateRead);
    assert.notEqual(result, sentinel);
    const code = Printer.printGoExpr(result);
    assert.ok(code.indexOf('GetFirst()') < code.indexOf('GetSecond()'));
    for (const update of [
        add(accumulator, letRead('first', 2, integer(60))),
        add(letRead('first', 2, integer(60)), accumulator),
    ]) assert.notEqual(optimize(fixture({update}),params,i64,self,sentinel,translateRead),sentinel);
    for (const value of [call(variable('observe'),[integer(0)]),local('captured',9),
        expr(new S.Accessor(variable('object'),new S.GetProp('value'))),expr(new S.EffectPure(integer(0)))]) {
        assert.equal(optimize(fixture({update:add(accumulator,letRead('unused',2,integer(60),value))})),sentinel);
    }
    assert.equal(optimize(fixture({update:add(accumulator,letRead('first',2,local('first',2)))})),sentinel);
    assert.equal(optimize(fixture({update:add(letRead('first',2,accumulator),integer(60))})),sentinel,
        'only the proven constant operand may contain reads');
    let constant = integer(60);
    for (let i=0;i<64;i++) constant=letRead('unused'+i,20+i,constant);
    assert.equal(optimize(fixture({step:letRead('first',2,call(variable(),[
        subtract(counter,integer(1)),add(accumulator,constant)]))})),sentinel,'the read budget spans both positions');
});

test('strict affine fast path keeps first initialization, order, failure and zero-demand behavior', t => {
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-affine-reads-'));
    t.after(() => rmSync(directory,{recursive:true,force:true}));
    const n = counterName + '_loop', a = accumulatorName + '_loop';
    const original = Go.rawGo(`for ${n} != 0 { _ = GetFirst(); _ = GetSecond(); ${n}=int64(int32(${n}-1)); ${a}=int64(int32(${a}+60)) }; return ${a}`);
    const print = (name, fallback) => Printer.printGoDecl(new Go.GoFunctionDecl({
        name:{owner:'generated/RenamedGlobals',name},
        params:params.map(p=>new Tuple(p.value0+'_loop',p.value1)), result:i64,
        body:optimize(fixture({step:nestedReadStep()}),params,i64,self,fallback,translateRead),
    }));
    writeFileSync(join(directory,'go.mod'),'module affineprobe\n\ngo 1.22\n');
    writeFileSync(join(directory,'affine.go'),`package affineprobe
import "sync"
var firstOnce,secondOnce sync.Once
var firstCalls,secondCalls int
var firstFailure bool
var trace string
func GetFirst()int64{firstCalls++;firstOnce.Do(func(){trace+="A";if firstFailure{panic("first initializer")}});return 41}
func GetSecond()int64{secondCalls++;secondOnce.Do(func(){trace+="B"});return 9}
${print('Run',original)}
${print('Guard',sentinel)}
`);
    writeFileSync(join(directory,'affine_test.go'),`package affineprobe
import ("sync";"testing")
func reset(){firstOnce=sync.Once{};secondOnce=sync.Once{};firstCalls=0;secondCalls=0;firstFailure=false;trace=""}
func invoke(n,a int64)(value int64,failed bool){defer func(){if recover()!=nil{failed=true}}();return Run(n,a),false}
func TestReads(t *testing.T){
 for _,n:=range []int64{1,2,20,1000,2147483647}{for _,seed:=range []int64{-2147483648,-1,0,2147483647}{
  reset();want:=int64(int32(seed+n*60));got,failed:=invoke(n,seed)
  if failed||got!=want||firstCalls!=1||secondCalls!=1||trace!="aAbB"{t.Fatalf("n=%d seed=%d got=%d want=%d calls=%d,%d trace=%s",n,seed,got,want,firstCalls,secondCalls,trace)}
 }}
 reset();firstFailure=true;if got,failed:=invoke(0,4294967296);failed||got!=4294967296||trace!=""||firstCalls!=0||secondCalls!=0{t.Fatal("zero demanded a global")}
 reset();firstFailure=true;if _,failed:=invoke(1,0);!failed||firstCalls!=1||secondCalls!=0||trace!="aA"{t.Fatal("first initialization failure did not stop later reads")}
 reset();if got:=Run(2,4294967296);got!=120||firstCalls!=2||secondCalls!=2{t.Fatal("wide seed must retain original loop and its reads")}
 for _,c:=range [][2]int64{{0,0},{-1,0},{-2147483648,0},{2147483648,0},{1,4294967296}}{reset();if Guard(c[0],c[1])!=-98||trace!=""||firstCalls!=0||secondCalls!=0{t.Fatal("unsupported input entered fast path")}}
}
`);
    const run=spawnSync('go',['test','-count=1','-v','.'],{cwd:directory,encoding:'utf8',timeout:60000,env:{...process.env,GOWORK:'off'}});
    assert.ifError(run.error);
    assert.equal(run.status,0,run.stdout+run.stderr);
});

test('generated affine fast paths preserve Int addition and leave unsupported inputs to the loop', t => {
    const directory = process.env.GOPURS_AFFINE_TEST_OUTPUT ?? mkdtempSync(join(tmpdir(), 'gopurs-affine-loops-'));
    if (!process.env.GOPURS_AFFINE_TEST_OUTPUT) t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(directory, { recursive: true });
    const definitions = [], cases = [], guards = [];
    const increments = [-2147483648, -2147483647, -7, 0, 1, 7, 2147483647];
    const seeds = [-2147483648, -1, 0, 1, 2147483647, 2147483648, 4294967295];
    const print = (name, body) => Printer.printGoDecl(new Go.GoFunctionDecl({
        name: { owner: 'generated/Probe', name },
        params: params.map(p => new Tuple(p.value0 + '_loop', p.value1)), result: i64, body,
    }));
    for (const [index, increment] of increments.entries()) {
        const n = counterName + '_loop', a = accumulatorName + '_loop';
        const fallback = Go.rawGo(`for ${n} != 0 { ${n} = int64(int32(${n} - 1)); ${a} = int64(int32(${a} + int64(${increment}))) }; return ${a}`);
        definitions.push(print('Run' + index, optimize(fixture({ increment }), params, i64, self, fallback)));
        definitions.push(print('Guard' + index, optimize(fixture({ increment }))));
        for (const count of [0, 1, 2, 7, 1000, 4294967297, -4294967295]) for (const seed of seeds) {
            let n = count, value = seed, iterations = 0;
            while (n !== 0) { assert.ok(iterations++ < 1002); n = (n - 1) | 0; value = (value + increment) | 0; }
            cases.push(`{Run${index}, ${count}, ${seed}, ${value}}`);
        }
        // Large counts execute only the proven path. BigInt supplies an exact
        // integer oracle independent of the Go formula and JS Number IntMul.
        for (const count of [2147483646, 2147483647]) for (const seed of seeds.slice(0, 5)) {
            const want = BigInt.asIntN(32, BigInt(seed) + BigInt(count) * BigInt(increment));
            cases.push(`{Run${index}, ${count}, ${seed}, ${want}}`);
        }
        for (const [count, seed] of [[0, 0], [-1, 0], [-2147483648, 0], [2147483648, 0], [4294967297, 0], [1, 2147483648]]) {
            guards.push(`{Guard${index}, ${count}, ${seed}, -98}`);
        }
    }
    writeFileSync(join(directory, 'go.mod'), 'module affineprobe\n\ngo 1.22\n');
    writeFileSync(join(directory, 'affine.go'), `package affineprobe\n${definitions.join('\n')}`);
    writeFileSync(join(directory, 'affine_test.go'), `package affineprobe
import "testing"
func TestAffine(t *testing.T) {
 for i,c := range []struct{f func(int64,int64)int64; n,a,want int64}{${[...cases, ...guards].join(',\n')}} {
  if got:=c.f(c.n,c.a); got!=c.want {t.Fatalf("case %d n=%d a=%d got=%d want=%d",i,c.n,c.a,got,c.want)}
 }
 t.Log("${cases.length} semantic cases and ${guards.length} guard refusals passed")
}
`);
    const run = spawnSync('go', ['test', '-count=1', '-v', '.'], { cwd: directory, encoding: 'utf8', timeout: 60000, env: { ...process.env, GOWORK: 'off' } });
    assert.ifError(run.error);
    assert.equal(run.status, 0, run.stdout + run.stderr);
    if (process.env.GOPURS_AFFINE_TEST_OUTPUT) writeFileSync(join(directory, 'go-test.log'), run.stdout + run.stderr);
});

test('generated remainder sums agree with iterative and BigInt floor-sum oracles', t => {
    const directory=mkdtempSync(join(tmpdir(),'gopurs-affine-modulo-'));
    t.after(()=>rmSync(directory,{recursive:true,force:true}));
    const definitions=[],cases=[],guards=[];
    const n=counterName+'_loop',a=accumulatorName+'_loop';
    const print=(name,body)=>Printer.printGoDecl(new Go.GoFunctionDecl({
        name:{owner:'generated/Modulo',name},params:params.map(p=>new Tuple(p.value0+'_loop',p.value1)),result:i64,body}));
    const seeds=[-2147483648n,-1n,0n,2147483647n,2147483648n,9223372036854775807n,-9223372036854775808n];
    for(const [index,d] of [1,2,3,7,65537,1073741824,2147483647].entries()) {
        const expression=fixture({update:add(accumulator,modulo(counter,integer(d)))});
        const fallback=Go.rawGo(`for ${n}!=0 { value:=${n}%int64(${d});if value<0 {value+=int64(${d})};${n}=int64(int32(${n}-1));${a}=int64(int32(${a}+value)) };return ${a}`);
        definitions.push(print('Run'+index,optimize(expression,params,i64,self,fallback)));
        definitions.push(print('Guard'+index,optimize(expression)));
        for(const count of [0n,1n,2n,3n,7n,31n,1000n,10000n,4294967297n,-4294967295n])for(const seed of seeds) {
            let k=count,value=seed,steps=0;
            while(k!==0n){assert.ok(steps++<10001);const rem=(k%BigInt(d)+BigInt(d))%BigInt(d);k=BigInt.asIntN(32,k-1n);value=BigInt.asIntN(32,value+rem);}
            cases.push(`{Run${index},${count},${seed},${value}}`);
        }
        for(const count of [2147483646n,2147483647n])for(const seed of seeds.slice(0,4)) {
            const divisor=BigInt(d),q=count/divisor;
            // Independent identity: sum(k) - d*sum(floor(k/d)), using BigInt.
            const floors=q*(count+1n)-divisor*q*(q+1n)/2n;
            const want=BigInt.asIntN(32,seed+count*(count+1n)/2n-divisor*floors);
            cases.push(`{Run${index},${count},${seed},${want}}`);
        }
        for(const [count,seed] of [[0,0],[-1,0],[-2147483648,0],[2147483648,0],[4294967297,0],[1,2147483648]])guards.push(`{Guard${index},${count},${seed},-98}`);
    }
    writeFileSync(join(directory,'go.mod'),'module moduloprobe\n\ngo 1.27\n');
    writeFileSync(join(directory,'modulo.go'),'package moduloprobe\n'+definitions.join('\n'));
    writeFileSync(join(directory,'modulo_test.go'),`package moduloprobe
import "testing"
func TestModulo(t *testing.T){for i,c:=range []struct{f func(int64,int64)int64;n,a,want int64}{${[...cases,...guards].join(',\n')}} {
if got:=c.f(c.n,c.a);got!=c.want{t.Fatalf("case%d n%d seed%d got%d want%d",i,c.n,c.a,got,c.want)}
};t.Log("${cases.length} semantic cases and ${guards.length} guard refusals")}
`);
    const run=spawnSync('go',['test','-count=1','-v','.'],{cwd:directory,encoding:'utf8',timeout:60000,env:{...process.env,GOWORK:'off'}});
    assert.ifError(run.error);assert.equal(run.status,0,run.stdout+run.stderr);
});
