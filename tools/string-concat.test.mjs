import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import * as Ref from '../output/Effect.Ref/index.js';
import { empty as emptySet } from '../output/Data.Set/index.js';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { binary } from '../output/Gopurs.PrimitiveExprs/index.js';
import { printGoExpr } from '../output/Gopurs.Printer/index.js';
import { runtimeGoCode } from '../output/Gopurs.Runtime/index.js';
import { OpStringAppend } from '../output/PureScript.Backend.Optimizer.Syntax/index.js';

test('generated and Prelude string concatenation match JavaScript UTF-16 boundaries', t => {
    const ref = Ref.new({ declarations: [], globalId: 0, reboxPairs: emptySet })();
    const emit = binary(ref)('Test')(OpStringAppend.value).value0;
    const operand = name => ({ expr: new Go.GoVar(name), exprType: Go.TypeString.value });
    const expression = printGoExpr(emit(operand('left'))(operand('right')).expr);
    const quote = text => printGoExpr(new Go.GoString(text));
    const strings = ['', 'A', '\0', 'é', '𝌆', '\ud800', '\udbff', '\ud834', '\udc00', '\udfff', '\udf06',
        'a\ud834', '\udf06z', '\ud834a', 'z\udf06', '\ud834\ud834', '\udf06\ud834'];
    const pairs = strings.flatMap(left => strings.map(right => [left, right]));
    let seed = 0x16c0de;
    const next = () => (seed = (Math.imul(seed, 1664525) + 1013904223) >>> 0);
    const units = [0, 34, 92, 65, 233, 0xd7ff, 0xd800, 0xd834, 0xdbff, 0xdc00, 0xdf06, 0xdfff, 0xe000, 0xffff];
    const sample = () => String.fromCharCode(...Array.from({ length: next() % 12 }, () => units[next() % units.length]));
    for (let i = 0; i < 512; i++) pairs.push([sample(), sample()]);
    const triples = [['\ud834', '', '\udf06'], ['\ud834', '\udf06', '\udf06'],
        ['\udf06', '\ud834', '\udf06'], ['a\ud834', '\udf06\ud800', '\udc00z']];
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-string-concat-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    mkdirSync(join(directory, 'gopurs_runtime'));
    mkdirSync(join(directory, 'purescript'));
    writeFileSync(join(directory, 'go.mod'), 'module gopurs/output\n\ngo 1.22\n');
    writeFileSync(join(directory, 'gopurs_runtime/runtime.go'), runtimeGoCode);
    const prelude = readFileSync(new URL('../../gopurs-prelude/src/Data/Semigroup.go', import.meta.url), 'utf8');
    writeFileSync(join(directory, 'purescript/semigroup.go'), 'package purescript\n\n' + prelude);
    writeFileSync(join(directory, 'concat_test.go'), `package concat_test
import ("testing"; "gopurs/output/gopurs_runtime"; "gopurs/output/purescript")
var _ = gopurs_runtime.Str
func generated(left, right string) string { return ${expression} }
func TestUTF16Concat(t *testing.T) {
    for name, concat := range map[string]func(string,string)string{"generated": generated, "Prelude": purescript.ConcatString} {
        t.Run(name, func(t *testing.T) {
            for i, c := range []struct { left, right, want string } {
                ${pairs.map(([left, right]) => `{${quote(left)}, ${quote(right)}, ${quote(left + right)}}`).join(',\n')},
            } {
                if got := concat(c.left, c.right); got != c.want { t.Fatalf("case %d: %x + %x = %x, want %x", i, c.left, c.right, got, c.want) }
            }
            ${triples.map(([a, b, c]) => `if concat(concat(${quote(a)}, ${quote(b)}), ${quote(c)}) != ${quote(a + b + c)} || concat(${quote(a)}, concat(${quote(b)}, ${quote(c)})) != ${quote(a + b + c)} { t.Fatal("UTF-16 concatenation is not associative") }`).join('\n')}
        })
    }
}
`);
    const result = spawnSync('go', ['test', '-race', '-count=1', '-v', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 120_000, env: { ...process.env, GOWORK: 'off' },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stdout + result.stderr);
    assert.match(result.stdout, /--- PASS: TestUTF16Concat/);
});
