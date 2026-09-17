import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { runtimeGoCode } from '../output/Gopurs.Runtime/index.js';

const arrays = new URL('../../gopurs-arrays/src/Data/', import.meta.url);
const javascript = readFileSync(new URL('Array.js', arrays), 'utf8');
const { sortByImpl } = await import('data:text/javascript;base64,' + Buffer.from(javascript).toString('base64'));

test('native Array.sortBy preserves equal-key order and matches its JS implementation', t => {
    const inputs = [[], [1], [2, 2, 1], [2, 1, 2, 0, 1, 2], [0, 0, 0, 0]];
    for (let seed = 0; seed < 40; seed++) {
        let random = seed + 1;
        inputs.push(Array.from({ length: seed + 2 }, () => {
            random = (Math.imul(random, 1664525) + 1013904223) >>> 0;
            return random % 5;
        }));
    }
    const cases = inputs.map(keys => {
        const input = keys.map((key, id) => ({ key, id }));
        const expected = sortByImpl(a => b => a.key - b.key, value => value, input);
        return { input, expected };
    });
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-native-stable-sort-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    writeFileSync(join(directory, 'array.go'), 'package main\nimport "gopurs/output/gopurs_runtime"\n' + readFileSync(new URL('Array.go', arrays), 'utf8'));
    writeFileSync(join(directory, 'cases.json'), JSON.stringify(cases));
    writeFileSync(join(directory, 'main.go'), `package main
import ("encoding/json"; "fmt"; "os"; "reflect")
type row struct { Key int; Id int }
type fixture struct { Input []row; Expected []row }
func main() {
    data, err := os.ReadFile("cases.json"); if err != nil { panic(err) }
    var fixtures []fixture; if err := json.Unmarshal(data, &fixtures); err != nil { panic(err) }
    for number, fixture := range fixtures {
        input := make([]interface{}, len(fixture.Input))
        for i, value := range fixture.Input { input[i] = value }
        result := SortByImpl(func(a, b interface{}) interface{} { return int64(a.(row).Key - b.(row).Key) }, func(value interface{}) int64 { return value.(int64) }, input)
        actual := make([]row, len(result))
        for i, value := range result { actual[i] = value.(row) }
        if !reflect.DeepEqual(actual, fixture.Expected) { panic(fmt.Sprintf("case %d: got %v, want %v", number, actual, fixture.Expected)) }
        for i, value := range input { if value.(row) != fixture.Input[i] { panic("input mutated") } }
    }
    fmt.Printf("%d stable sorts match JS; inputs unchanged\\n", len(fixtures))
}
`);
    const result = spawnSync('go', ['run', '.'], { cwd: directory, encoding: 'utf8', timeout: 30_000, env: { ...process.env, GOWORK: 'off' } });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stderr);
    assert.match(result.stdout, /45 stable sorts match JS; inputs unchanged/);
});
