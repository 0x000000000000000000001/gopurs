import assert from 'node:assert/strict';
import test from 'node:test';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { optimize } from '../output/Gopurs.RecordReadGroups/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';
import { goModule, runGo, writeGoFile } from './lib/go.mjs';

const variable = name => new Go.GoVar(name);
const integer = value => new Go.GoCall(variable('int64'), [new Go.GoInt(value)]);
const call = (name, ...args) => new Go.GoCall(new Go.GoSelector(variable('gopurs_runtime'), name), args);
const read = (key, source = 'value') => new Go.GoSelector(new Go.GoRecordAccess(variable(source), key), 'IntVal');
const sum = (keys, source = 'value') => keys.map(key => read(key, source)).reduce((a, b) => call('IntAdd', a, b));
const keys = [...'abcdefgh'];
const simple = sum(keys);
const text = expression => printGoExpr(expression);
const marker = /recordReadGroup.*_data/;

test('integer regions use full unique shape guards and retain their exact fallback', () => {
    const original = text(simple), changed = text(optimize(simple));
    assert.match(changed, marker);
    assert.match(changed, /len\(recordReadGroup_data.Keys\) == 8/);
    assert.match(changed, /len\(recordReadGroup_data.Vals\) >= 8/);
    assert.ok(changed.includes(`return ${original}\n}`));
    assert.equal((changed.match(/func\(recordReadGroup_input/g) ?? []).length, 1);
    assert.deepEqual(optimize(optimize(simple)), optimize(simple));
    assert.match(text(optimize(sum([...keys.slice(0, 5), 'a', 'a']))), /TypeRecord5/);
    const compactTwo = text(optimize(sum(['a', 'b'])));
    assert.match(compactTwo, /RecordFields2\(recordReadGroup_input\)/);
    assert.match(compactTwo, /recordReadGroup_keys != nil/);
    assert.match(compactTwo, /int64\(int32\(/);
    assert.ok(compactTwo.includes(`return ${text(sum(['a', 'b']))}\n}`));
    for (const object of [new Go.GoCall(variable('source'), []), new Go.GoSelector(variable('nested'), 'field')]) {
        const reads = keys.map(key => new Go.GoSelector(new Go.GoRecordAccess(object, key), 'IntVal'));
        assert.doesNotMatch(text(optimize(reads.reduce((a, b) => call('IntAdd', a, b)))), marker);
    }
    assert.equal((text(optimize(call('IntAdd', sum(keys.slice(0, 4), 'left'), sum(keys.slice(4), 'right')))).match(/func\(recordReadGroup_input/g) ?? []).length, 2);
    assert.equal((text(optimize(new Go.GoCall(variable('unknown'), [sum(keys.slice(0, 4)), sum(keys.slice(4))]))).match(/func\(recordReadGroup_input/g) ?? []).length, 2);
    const collision = call('IntAdd', variable('recordReadGroup_input'), sum(keys, 'recordReadGroup_data'));
    assert.match(text(optimize(collision)), /recordReadGroup__input/);
    for (const suffix of ['_keys', '_value0', '_value1']) {
        assert.match(text(optimize(call('IntAdd', variable(`recordReadGroup${suffix}`), sum(['a', 'b'])))), /recordReadGroup__input/);
    }
});

test('generated grouped reads preserve arithmetic, strict order, getter failures, shapes and unknown-call boundaries', t => {
    const interleaved = call('IntAdd', call('IntAdd', sum(keys.slice(0, 3)), new Go.GoCall(variable('mark'), [variable('value')])), sum(keys.slice(3)));
    const afterMutation = call('IntAdd', call('IntAdd', sum(keys.slice(0, 6)), new Go.GoCall(variable('mutate'), [variable('value')])), read('a'));
    const collision = call('IntAdd', variable('recordReadGroup_input'), sum(keys, 'recordReadGroup_data'));
    const candidates = [
        ...Array.from({ length: 5 }, (_, i) => [`compact${i + 1}`, sum(i === 0 ? ['a', 'a'] : keys.slice(0, i + 1))]),
        ['sum', simple],
        ['repeated', sum([...keys].reverse().concat(['a', 'c', 'h']))],
        ['native', call('IntAdd', new Go.GoRecordFieldAccess(variable('native'), 'prefix'), simple)],
        ['arithmetic', call('IntMod', call('IntMul', call('IntSub', simple, integer(2147483647)), integer(3)), integer(7))],
        ['constantBoundary', call('IntMul', call('IntAdd', integer(2147483647), integer(2147483647)), simple)],
        ['constantIsland', call('IntAdd', call('IntAdd', integer(2147483647), integer(2147483647)), simple)],
        ['zero', call('IntDiv', simple, integer(0))],
        ['interleaved', interleaved],
        ['afterMutation', afterMutation],
        ['compactMutation', call('IntAdd', call('IntAdd', sum(['a', 'b']), new Go.GoCall(variable('mutateCompact'), [variable('value')])), sum(['a', 'b']))],
    ];
    assert.equal((text(optimize(interleaved)).match(/func\(recordReadGroup_input/g) ?? []).length, 2);
    assert.match(text(optimize(afterMutation)), marker);
    const functions = candidates.map(([name, expression]) => `
func ${name}Before(value gopurs_runtime.Value)int64{native:=struct{${Go.recordFieldName('prefix')} int64}{2147483647};_=native;return ${text(expression)}}
func ${name}After(value gopurs_runtime.Value)int64{native:=struct{${Go.recordFieldName('prefix')} int64}{2147483647};_=native;return ${text(optimize(expression))}}
`).join('\n');
    const cast = new Go.GoCall(variable('int32'), [simple]);
    const directory = goModule(t, 'gopurs-record-read-groups-', { runtime: runtimeGoCode });
    writeGoFile(directory, 'reads_test.go', `package main
import("fmt";"reflect";"strings";"testing";"unsafe";"gopurs/output/gopurs_runtime")
var trace []string
func mark(value gopurs_runtime.Value)int64{trace=append(trace,"call");return 13}
func mutateCompact(value gopurs_runtime.Value)int64{trace=append(trace,"compact mutation");if value.Type==gopurs_runtime.TypeRecord2&&value.UnsafePtr!=nil{keys,v0,v1:=gopurs_runtime.RecordFields2(value);if keys!=nil{if value.IntVal==1{data:=(*gopurs_runtime.RecordData2Int)(value.UnsafePtr);data.V0=v1.IntVal;data.V1=v0.IntVal+991}else{data:=(*gopurs_runtime.RecordData2)(value.UnsafePtr);data.V0=v1;data.V1=gopurs_runtime.Int(v0.IntVal+991)}}};return 19}
func mutate(value gopurs_runtime.Value)int64{trace=append(trace,"mutation");if value.Type==gopurs_runtime.TypeRecordData{data:=(*gopurs_runtime.RecordData)(value.UnsafePtr);if len(data.Vals)>0{data.Vals[0]=gopurs_runtime.Int(991)}};return 19}
${functions}
func collisionBefore(recordReadGroup_data gopurs_runtime.Value,recordReadGroup_input int64)int64{return ${text(collision)}}
func collisionAfter(recordReadGroup_data gopurs_runtime.Value,recordReadGroup_input int64)int64{return ${text(optimize(collision))}}
func castAfter(value gopurs_runtime.Value)int32{return ${text(optimize(cast))}}
type outcome struct{value int64;failure string;log []string}
func observe(f func(gopurs_runtime.Value)int64,input gopurs_runtime.Value)(out outcome){trace=nil;defer func(){if p:=recover();p!=nil{out.failure=strings.Split(fmt.Sprint(p),"Object:")[0]};out.log=append([]string{},trace...)}();out.value=f(input);return}
func fixture(mode,n int)gopurs_runtime.Value{keys:=[]string{"a","b","c","d","e","f","g","h"}[:n];keys=append([]string{},keys...);values:=make([]gopurs_runtime.Value,n);m:=map[string]gopurs_runtime.Value{};foreign:=map[string]any{};for i,key:=range keys{values[i]=gopurs_runtime.Int(int64(2147483647-i*997));m[key]=values[i];foreign[key]=values[i].IntVal};if n>1{values[1]=gopurs_runtime.Int(1<<40);m[keys[1]]=values[1];foreign[keys[1]]=int64(1<<40)}
switch mode{
case 1:for i,j:=0,n-1;i<j;i,j=i+1,j-1{keys[i],keys[j]=keys[j],keys[i];values[i],values[j]=values[j],values[i]}
case 2:keys=append(keys,"extra");values=append(values,gopurs_runtime.Int(999))
case 3:keys=append([]string{keys[n-1]},keys...);values=append([]gopurs_runtime.Value{gopurs_runtime.Int(-555)},values...)
case 4:keys[n-1]=keys[0]
case 5:keys=keys[:n-1];values=values[:n-1]
case 6:values=values[:n-1]
case 7:return gopurs_runtime.Record(m)
case 8:return gopurs_runtime.Any(foreign)
case 9:return gopurs_runtime.Record(map[string]gopurs_runtime.Value{})
case 10,11:gopurs_runtime.StructGetters[701]=func(_ unsafe.Pointer,key string)gopurs_runtime.Value{trace=append(trace,key);if mode==11&&key=="c"{panic("getter panic")};return m[key]};return gopurs_runtime.Value{Type:gopurs_runtime.TypeConstructor,IntVal:701}
case 12:return gopurs_runtime.Value{}
case 13:return gopurs_runtime.Value{Type:gopurs_runtime.TypeRecordData}
};return gopurs_runtime.RecordDict(keys,values)}
func fixtureCompact(n,mode int)gopurs_runtime.Value{keys:=[]string{"a","b","c","d","e"}[:n];keys=append([]string{},keys...);raw:=[]int64{9223372036854775807,-9223372036854775808,1<<40,2147483647,-2147483648};values:=make([]gopurs_runtime.Value,n);for i:=range values{values[i]=gopurs_runtime.Int(raw[i])};if mode==1{values[0]=gopurs_runtime.Value{Type:gopurs_runtime.TypeFloat,IntVal:raw[0]}};if mode==2{for i,j:=0,n-1;i<j;i,j=i+1,j-1{keys[i],keys[j]=keys[j],keys[i];values[i],values[j]=values[j],values[i]}};if mode==3&&n>1{keys[n-1]=keys[0]};var value gopurs_runtime.Value
switch n{
case 1:value=gopurs_runtime.RecordDict1(keys[0],values[0]);if mode==4{(*gopurs_runtime.RecordData1)(value.UnsafePtr).RecordKeys1=nil}
case 2:value=gopurs_runtime.RecordDict2(keys[0],keys[1],values[0],values[1]);if mode==4{value=gopurs_runtime.RecordShape2(nil,values[0],values[1])}
case 3:value=gopurs_runtime.RecordDict3(keys[0],keys[1],keys[2],values[0],values[1],values[2]);if mode==4{(*gopurs_runtime.RecordData3)(value.UnsafePtr).RecordKeys3=nil}
case 4:value=gopurs_runtime.RecordDict4(keys[0],keys[1],keys[2],keys[3],values[0],values[1],values[2],values[3]);if mode==4{(*gopurs_runtime.RecordData4)(value.UnsafePtr).RecordKeys4=nil}
case 5:value=gopurs_runtime.RecordDict5(keys[0],keys[1],keys[2],keys[3],keys[4],values[0],values[1],values[2],values[3],values[4]);if mode==4{(*gopurs_runtime.RecordData5)(value.UnsafePtr).RecordKeys5=nil}
};if mode==5{value.UnsafePtr=nil};return value}
func TestReadGroups(t *testing.T){checks:=0;cases:=[]struct{name string;before,after func(gopurs_runtime.Value)int64}{
${candidates.map(([name]) => `{"${name}",${name}Before,${name}After}`).join(',\n')},
};for _,n:=range []int{6,8}{for mode:=0;mode<14;mode++{for _,c:=range cases{expected:=observe(c.before,fixture(mode,n));actual:=observe(c.after,fixture(mode,n));if !reflect.DeepEqual(actual,expected){t.Fatalf("%s n=%d mode=%d got=%+v want=%+v",c.name,n,mode,actual,expected)};checks++}}};for n:=1;n<=5;n++{for mode:=0;mode<6;mode++{for _,c:=range cases{expected:=observe(c.before,fixtureCompact(n,mode));actual:=observe(c.after,fixtureCompact(n,mode));if !reflect.DeepEqual(actual,expected){t.Fatalf("compact %s n=%d mode=%d got=%+v want=%+v",c.name,n,mode,actual,expected)};checks++}}};value:=fixture(0,8);if collisionBefore(value,17)!=collisionAfter(value,17){t.Fatal("temporary name capture")};if castAfter(value)!=int32(sumBefore(value)){t.Fatal("int32 result type changed")};for repeat:=0;repeat<3;repeat++{input:=fixture(0,8);if sumBefore(input)!=sumAfter(input)||sumBefore(input)!=sumAfter(input){t.Fatal("replay or source mutation")}};changed:=fixture(0,8);if repeatedBefore(changed)!=repeatedAfter(changed){t.Fatal("initial repeated read")};data:=(*gopurs_runtime.RecordData)(changed.UnsafePtr);data.Keys[0],data.Keys[1]=data.Keys[1],data.Keys[0];if repeatedBefore(changed)!=repeatedAfter(changed){t.Fatal("shape proof cached across calls")};t.Logf("%d strict shape/arithmetic checks passed",checks)}
`);
    runGo(directory, ['test', '-count=1', '.']);
});
