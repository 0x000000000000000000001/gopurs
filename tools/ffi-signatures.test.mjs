import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { copyFileSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { fileURLToPath } from 'node:url';
import test from 'node:test';

const runner = fileURLToPath(new URL('./ffi-runner.mjs', import.meta.url));
const run = (command, args, options = {}) => {
    const result = spawnSync(command, args, { encoding: 'utf8', timeout: 120_000, ...options });
    assert.ifError(result.error);
    return result;
};

test('WASM and native expose identical complete FFI signatures and source positions', t => {
    const directory = mkdtempSync(join(tmpdir(), 'gopurs-ffi-signatures-'));
    t.after(() => rmSync(directory, { recursive: true, force: true }));
    for (const name of ['parser.go', 'types.go', 'native_api.go', 'go.mod'])
        copyFileSync(new URL(`./ffi-gen/${name}`, import.meta.url), join(directory, name));
    writeFileSync(join(directory, 'main.go'), `package main
import("fmt"; "io"; "os")
func main() {
 input, err := io.ReadAll(os.Stdin)
 if err != nil { panic(err) }
 var result string
 if len(os.Args) > 1 { result, err = Prepare(os.Args[1], string(input)) } else { result, err = Extract(string(input)) }
 if err != nil { fmt.Fprint(os.Stderr, err); os.Exit(1) }
 fmt.Println(result)
}
`);
    const binary = join(directory, 'extract');
    const build = run('go', ['build', '-o', binary, '.'], { cwd: directory, env: { ...process.env, GOWORK: 'off' } });
    assert.equal(build.status, 0, build.stderr);
    const source = `package ffi
import c "context"
type C = c.Context
type E = error
type Handle struct{}
func Read(ctx C, callback func(left, right string) (accepted bool, err E), pointers []*Handle, fixed [3]int, rest ...string) (first, second string, err E) { panic("only extracted") }
func NoResult() {}
type Record struct { Label string; Nested struct { Count int }; Parent *Record }
func RecordValue(value Record) Record { return value }
`;
    for (const args of [[], ['Fixture_']]) {
        const wasm = run(process.execPath, [runner, ...args], { input: source });
        const native = run(binary, args, { input: source });
        assert.equal(wasm.status, 0, wasm.stderr);
        assert.equal(native.status, 0, native.stderr);
        assert.equal(native.stdout, wasm.stdout);
        const parsed = JSON.parse(wasm.stdout);
        const declarations = args.length ? parsed.decls : parsed;
        assert.equal(declarations[0].schemaVersion, 2);
        assert.equal(declarations[0].results.length, 3);
        assert.equal(declarations[0].ret, null);
        assert.deepEqual(declarations[0].resultNames, ['first', 'second', 'err']);
        assert.equal(declarations[0].args[0].identity, 'context.Context');
        assert.equal(declarations[0].results[2].identity, 'builtin.error');
        assert.equal(declarations[0].args[1].results.length, 2);
        assert.equal(declarations[0].args[3].type, 'FixedArray');
        assert.equal(declarations[0].args[4].type, 'Variadic');
        assert.deepEqual(declarations[1].results, []);
        assert.equal(declarations[0].line, 6);
        assert.equal(declarations[2].args[0].underlying.type, 'Struct');
        assert.deepEqual(declarations[2].args[0].underlying.fields.map(field => field.name), ['Label', 'Nested', 'Parent']);
        assert.equal(declarations[2].args[0].underlying.fields[1].type.type, 'Struct');
        assert.equal(declarations[2].args[0].underlying.fields[2].type.elem.name, 'Record');
        if (args.length) assert.ok(parsed.content.includes('func Fixture_Read('));
    }
    const invalid = 'func Valid() {}\nfunc Broken(';
    const native = run(binary, [], { input: invalid });
    const wasm = run(process.execPath, [runner], { input: invalid });
    assert.notEqual(native.status, 0);
    assert.notEqual(wasm.status, 0);
    assert.equal(native.stdout, '');
    assert.equal(wasm.stdout, '');
    assert.ok(wasm.stderr.includes(native.stderr), `${wasm.stderr} should contain ${native.stderr}`);
    assert.match(native.stderr, /2:13/);
});
