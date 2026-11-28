import assert from 'node:assert/strict';
import test from 'node:test';
import { empty as emptySet } from '../output/Data.Set/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import * as Ref from '../output/Effect.Ref/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { unboxGoExpr } from '../output/Gopurs.GoConversions/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime.Source/index.js';
import { goModule, runGo, writeGoFile } from './lib/go.mjs';

const state = () => Ref.new({ declarations: [], globalId: 0, reboxPairs: emptySet })();
const record = fields => new Go.TypeRecord(fields.map(([key, type = Go.TypeInt64.value]) => new Tuple(key, type)));
const projection = fields => printGoExpr(unboxGoExpr(state())('Probe')(Go.rawGo('once(value)'))(Go.TypeValue.value)(record(fields)));
const quote = value => printGoExpr(new Go.GoString(value));

test('record projection guards are limited to unique integer fields', () => {
    const integers = Array.from({ length: 8 }, (_, index) => [`field${index}`]);
    assert.match(projection(integers), /recordProjection_data/);
    assert.match(projection([['alpha'], ['omega']]), /RecordIntFields2\(orig\)/);
    for (let size = 1; size <= 5; size++) {
        assert.match(projection(integers.slice(0, size)), new RegExp(`orig.Type == gopurs_runtime.TypeRecord${size}`));
    }
    assert.doesNotMatch(projection([]), /recordProjection_data/);
    assert.doesNotMatch(projection([]), /orig.Type ==/);
    assert.doesNotMatch(projection([['alpha'], ['alpha']]), /orig.Type ==/);
    assert.doesNotMatch(projection([...integers, ['field0']]), /recordProjection_data/);
    for (const type of [Go.TypeBool.value, Go.TypeString.value, Go.TypeFloat64.value, Go.TypeValue.value]) {
        assert.doesNotMatch(projection([...integers, ['other', type]]), /recordProjection_data/);
        for (let size = 1; size <= 5; size++) {
            assert.doesNotMatch(projection([...integers.slice(0, size - 1), ['other', type]]), /orig.Type ==/);
        }
    }
});

test('generated guarded projections preserve lookup order, duplicate keys, failures, representations and aliases', t => {
    const shapes = [0, 1, 2, 3, 4, 5, 8, 16].map(size => Array.from({ length: size }, (_, i) => `field ${i}`));
    const functions = shapes.map((keys, index) => {
        const expr = projection(keys.map(key => [key]));
        const fields = keys.map(key => printGoExpr(new Go.GoRecordFieldAccess(new Go.GoVar('native'), key)));
        return `func project${index}(value gopurs_runtime.Value) []int64 { native := ${expr}; _ = native; return []int64{${fields.join(',')}} }`;
    }).join('\n');
    const directory = goModule(t, 'gopurs-record-projection-guard-', { runtime: runtimeGoCode });
    writeGoFile(directory, 'projection_test.go', `package main
import("fmt";"reflect";"testing";"unsafe";"gopurs/output/gopurs_runtime")
var reads int
var trace []string
func once(value gopurs_runtime.Value) gopurs_runtime.Value {reads++;return value}
${functions}
type outcome struct { values []int64; failure string; log []string }
func observe(run func()[]int64)(out outcome){trace=nil;defer func(){if p:=recover();p!=nil{out.failure=fmt.Sprint(p)};out.log=append([]string{},trace...)}();out.values=run();return}
func compact(keys []string, values []gopurs_runtime.Value)gopurs_runtime.Value{switch len(keys){
case 0:return gopurs_runtime.RecordDict0()
case 1:return gopurs_runtime.RecordDict1(keys[0],values[0])
case 2:return gopurs_runtime.RecordDict2(keys[0],keys[1],values[0],values[1])
case 3:return gopurs_runtime.RecordDict3(keys[0],keys[1],keys[2],values[0],values[1],values[2])
case 4:return gopurs_runtime.RecordDict4(keys[0],keys[1],keys[2],keys[3],values[0],values[1],values[2],values[3])
case 5:return gopurs_runtime.RecordDict5(keys[0],keys[1],keys[2],keys[3],keys[4],values[0],values[1],values[2],values[3],values[4])
};return gopurs_runtime.RecordDict(keys,values)}
func TestProjectionMatrix(t *testing.T){cases:=[]struct{keys []string;project func(gopurs_runtime.Value)[]int64}{
${shapes.map((keys, i) => `{[]string{${keys.map(quote).join(',')}},project${i}}`).join(',\n')},
};checks:=0
for _,c:=range cases{values:=make([]gopurs_runtime.Value,len(c.keys));boxed:=map[string]gopurs_runtime.Value{};foreign:=map[string]any{};for i,k:=range c.keys{values[i]=gopurs_runtime.Int(int64(2147483647-i*99));if i==1{values[i]=gopurs_runtime.Int(1<<52+1)};boxed[k]=values[i];foreign[k]=int64(values[i].IntVal)}
reverseKeys:=append([]string{},c.keys...);reverseValues:=append([]gopurs_runtime.Value{},values...);for i,j:=0,len(c.keys)-1;i<j;i,j=i+1,j-1{reverseKeys[i],reverseKeys[j]=reverseKeys[j],reverseKeys[i];reverseValues[i],reverseValues[j]=reverseValues[j],reverseValues[i]}
inputs:=[]gopurs_runtime.Value{
gopurs_runtime.RecordDict(c.keys,values),compact(c.keys,values),gopurs_runtime.RecordDict(reverseKeys,reverseValues),compact(reverseKeys,reverseValues),
gopurs_runtime.RecordDict(append(append([]string{},c.keys...),"extra"),append(append([]gopurs_runtime.Value{},values...),gopurs_runtime.Int(1))),
gopurs_runtime.Record(boxed),gopurs_runtime.Any(foreign),gopurs_runtime.Record(map[string]gopurs_runtime.Value{}),
{Type:gopurs_runtime.TypeConstructor,IntVal:777},{Type:gopurs_runtime.TypeConstructor,IntVal:778},{Type:gopurs_runtime.TypeConstructor,IntVal:779},{},
{Type:gopurs_runtime.TypeRecord1},{Type:gopurs_runtime.TypeRecord2},{Type:gopurs_runtime.TypeRecord2,IntVal:1},
{Type:gopurs_runtime.TypeRecord3},{Type:gopurs_runtime.TypeRecord4},{Type:gopurs_runtime.TypeRecord5},{Type:gopurs_runtime.TypeRecordData},
gopurs_runtime.RecordShape1(nil,gopurs_runtime.Int(1)),gopurs_runtime.RecordShape2(nil,gopurs_runtime.Int(1),gopurs_runtime.Int(2)),
gopurs_runtime.RecordShape2(nil,gopurs_runtime.Int(1),gopurs_runtime.Bool(true)),
gopurs_runtime.RecordShape3(nil,gopurs_runtime.Int(1),gopurs_runtime.Int(2),gopurs_runtime.Int(3)),
gopurs_runtime.RecordShape4(nil,gopurs_runtime.Int(1),gopurs_runtime.Int(2),gopurs_runtime.Int(3),gopurs_runtime.Int(4)),
gopurs_runtime.RecordShape5(nil,gopurs_runtime.Int(1),gopurs_runtime.Int(2),gopurs_runtime.Int(3),gopurs_runtime.Int(4),gopurs_runtime.Int(5)),
}
if len(c.keys)>0{duplicateKeys:=append([]string{c.keys[len(c.keys)-1]},c.keys...);duplicateValues:=append([]gopurs_runtime.Value{gopurs_runtime.Int(-987)},values...);sameKeys:=append([]string{},c.keys...);sameKeys[len(sameKeys)-1]=sameKeys[0]
inputs=append(inputs,gopurs_runtime.RecordDict(duplicateKeys,duplicateValues),gopurs_runtime.RecordDict(sameKeys,values),compact(sameKeys,values),gopurs_runtime.RecordDict(c.keys[:len(c.keys)-1],values[:len(values)-1]),gopurs_runtime.RecordDict(c.keys,values[:len(values)-1]))}
if len(c.keys)==2{keys:=&gopurs_runtime.RecordKeys2{K0:c.keys[0],K1:c.keys[1]};legacy:=&gopurs_runtime.RecordData2{RecordKeys2:keys,V0:values[0],V1:values[1]};nilKeys:=&gopurs_runtime.RecordData2{V0:values[0],V1:values[1]}
inputs=append(inputs,gopurs_runtime.Value{Type:gopurs_runtime.TypeRecord2,UnsafePtr:unsafe.Pointer(legacy)},gopurs_runtime.Value{Type:gopurs_runtime.TypeRecord2,IntVal:77,UnsafePtr:unsafe.Pointer(legacy)},gopurs_runtime.Value{Type:gopurs_runtime.TypeRecord2,UnsafePtr:unsafe.Pointer(nilKeys)},gopurs_runtime.RecordShape2(keys,values[0],gopurs_runtime.Bool(true)))}
gopurs_runtime.StructGetters[777]=func(_ unsafe.Pointer,k string)gopurs_runtime.Value{trace=append(trace,k);return boxed[k]}
gopurs_runtime.StructGetters[778]=func(_ unsafe.Pointer,k string)gopurs_runtime.Value{trace=append(trace,k);if len(c.keys)>0&&k==c.keys[len(c.keys)/2]{panic("getter failure")};return boxed[k]}
gopurs_runtime.StructGetters[779]=func(_ unsafe.Pointer,k string)gopurs_runtime.Value{trace=append(trace,k);if len(c.keys)>0&&k==c.keys[0]{last:=c.keys[len(c.keys)-1];boxed[last]=gopurs_runtime.Int(boxed[last].IntVal+17)};return boxed[k]}
reset:=func(){for i,key:=range c.keys{boxed[key]=values[i]}}
for mode,input:=range inputs{reset();expected:=observe(func()[]int64{result:=make([]int64,len(c.keys));for i,k:=range c.keys{result[i]=gopurs_runtime.RecordGet(input,k).IntVal};return result});reset();reads=0;actual:=observe(func()[]int64{return c.project(input)});if !reflect.DeepEqual(actual,expected)||reads!=1{t.Fatalf("size=%d mode=%d actual=%+v expected=%+v reads=%d",len(c.keys),mode,actual,expected,reads)};checks++}
// A previous successful shape check is not cached across later key changes.
mutableKeys:=append([]string{},c.keys...);mutableValues:=append([]gopurs_runtime.Value{},values...);input:=gopurs_runtime.RecordDict(mutableKeys,mutableValues);first:=c.project(input);if len(c.keys)>1{last:=len(c.keys)-1;mutableKeys[0],mutableKeys[last]=mutableKeys[last],mutableKeys[0];mutableValues[0],mutableValues[last]=mutableValues[last],mutableValues[0]};if !reflect.DeepEqual(first,c.project(input)){t.Fatal("shape result was cached")};if len(c.keys)>0{first[0]=0;if gopurs_runtime.RecordGet(input,c.keys[0]).IntVal!=values[0].IntVal{t.Fatal("projection mutated source")}}
};t.Logf("%d representation/shape checks passed",checks)}
`);
    const result = runGo(directory, ['test', '-count=1', '-v', '.']);
    assert.match(result.stdout, /239 representation\/shape checks passed/);
    console.log(result.stdout.trim());
});
