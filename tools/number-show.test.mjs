import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { mkdtempSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import { showNumberImpl } from '../../gopurs-prelude/src/Data/Show.js';

test('Go Show Number matches the JavaScript Prelude across binary64 boundaries', t => {
    const view = new DataView(new ArrayBuffer(8));
    const bitsOf = n => { view.setFloat64(0, n); return view.getBigUint64(0); };
    const fromBits = bits => { view.setBigUint64(0, bits); return view.getFloat64(0); };
    const groups = {};
    const add = (group, bits) => {
        (groups[group] ??= new Map()).set(bits.toString(16), showNumberImpl(fromBits(bits)));
    };
    const both = (group, bits) => {
        add(group, bits);
        add(group, bits | 0x8000000000000000n);
    };
    for (const n of [0, -0, NaN, Infinity, -Infinity]) add('special', bitsOf(n));
    for (const bits of [0x7ff0000000000001n, 0x7ff8000000000001n, 0x7fffffffffffffffn]) both('special', bits);
    for (const n of [0.17, 0.25996181067141905, 0.3572019862807257, 0.46817723004874223,
        0.9640035681058178, 4.23808622486133, 4.540362294799751, 5.212384849884261,
        13.958257048123212, 32.96176575630599, 38.47735512322269, 1e10, 1e-5,
        1.5339794352098402e-118, 2.108934760892056e-59, 2.250634744599241e-19,
        5.960464477539063e-8, 5e-324]) both('literal-regression', bitsOf(n));
    for (const n of [1e-9, 1e-8, 1e-7, 1e-6, 1e20, 1e21, 1e23, Number.MIN_VALUE,
        Number.MAX_VALUE, Number.MAX_SAFE_INTEGER, Number.MAX_SAFE_INTEGER + 1,
        1000000000000000128, 1000000000000000100, 2.2250738585072014e-308]) {
        both('notation-and-shortest-roundtrip', bitsOf(n));
    }
    // Adjacent floats on both sides of each decimal notation/precision boundary.
    for (let exponent = -324; exponent <= 308; exponent++) {
        const bits = bitsOf(Number(`1e${exponent}`));
        for (const delta of [-1n, 0n, 1n]) if (bits + delta >= 0n) both('decimal-boundaries', bits + delta);
    }
    // Every finite binary exponent, including subnormals and mantissa carries.
    for (let exponent = 0n; exponent < 2047n; exponent++) {
        for (const fraction of [0n, 1n, 0x7ffffffffffffn, 0x8000000000000n, 0xffffffffffffen, 0xfffffffffffffn]) {
            both('binary-boundaries', (exponent << 52n) | fraction);
        }
    }
    let seed = 0x2115n;
    for (let i = 0; i < 4096; i++) {
        seed = BigInt.asUintN(64, seed * 6364136223846793005n + 1442695040888963407n);
        add('random-bit-patterns', seed);
    }
    const cases = Object.fromEntries(Object.entries(groups).map(([group, values]) => [group, [...values]]));
    const count = Object.values(cases).reduce((sum, values) => sum + values.length, 0);
    t.diagnostic(`${count} binary64 values compared with the actual JS FFI`);
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-number-show-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    writeFileSync(join(directory, 'go.mod'), 'module number-show\n\ngo 1.22\n');
    writeFileSync(join(directory, 'show.go'), 'package show\n' + readFileSync(new URL('../../gopurs-prelude/src/Data/Show.go', import.meta.url), 'utf8'));
    writeFileSync(join(directory, 'cases.json'), JSON.stringify(cases));
    writeFileSync(join(directory, 'show_test.go'), `package show
import ("encoding/json"; "math"; "os"; "strconv"; "testing")
func TestJavaScriptOracle(t *testing.T) {
    data, err := os.ReadFile("cases.json"); if err != nil { t.Fatal(err) }
    var groups map[string][][2]string
    if err := json.Unmarshal(data, &groups); err != nil { t.Fatal(err) }
    for group, cases := range groups {
        t.Run(group, func(t *testing.T) {
            failures := 0
            for _, c := range cases {
                bits, err := strconv.ParseUint(c[0], 16, 64); if err != nil { t.Fatal(err) }
                got := ShowNumberImpl(math.Float64frombits(bits))
                if got != c[1] {
                    t.Errorf("bits=%016x: got %q, JS wants %q", bits, got, c[1])
                    failures++
                    if failures == 5 { break }
                }
            }
        })
    }
}
`);
    const result = spawnSync('go', ['test', '-race', '-count=1', '-v', '.'], {
        cwd: directory, encoding: 'utf8', timeout: 120_000, env: { ...process.env, GOWORK: 'off' },
    });
    assert.ifError(result.error);
    assert.equal(result.status, 0, result.stdout + result.stderr);
    assert.match(result.stdout, /--- PASS: TestJavaScriptOracle /);
});
