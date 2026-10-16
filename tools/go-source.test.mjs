import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { createHash } from 'node:crypto';
import { copyFileSync, mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import test from 'node:test';
import * as Map from '../output/Data.Map/index.js';
import { Left, Right } from '../output/Data.Either/index.js';
import * as Symbols from '../output/Gopurs.GoSymbol/index.js';
import { bindSources, packageSources } from '../output/Gopurs.GoSource/index.js';
import { bindGoSourceImpl, packageGoSourceImpl, packageGoFilesImpl } from '../output/Gopurs.FfiSupport/index.js';
import { resolveGoPackagesHelper } from './go-packages.mjs';
import * as Go from '../output/Gopurs.GoAst/index.js';
import { Tuple } from '../output/Data.Tuple/index.js';
import { printBoundGoCode } from '../output/Gopurs.Printer/index.js';
import { boundImports } from '../output/Gopurs.GoImports/index.js';

const provider = `package purescript
type Box[T any] struct { Key T }
func (renamed *Box[T]) Read() T { return renamed.Key }
type IntMap map[int]int
const Key = 3
func Produce() int { return 21 }
type Readable[T any] interface { Read() T }
`;
const consumer = `package purescript
// 😀 Produce Key: never a source reference.
func Use() int {
  value := Box[int]{Key: Produce()}
  entries := IntMap{Key: 10}
  callback := func(Produce func() int) int { return Produce() }
  var boxed any = value
  { Key := 99; _ = Box[int]{Key: Key} }
  return boxed.(Box[int]).Key + entries[Key] + callback(Produce) - 10
}
func init() { _ = Produce() }
func ReadBox(value Readable[int]) int { return value.Read() }
type Wrapped struct { Box[int] }
func ReadWrapped(value Wrapped) int { return value.Box.Read() }
`;
const request = {
    packagePath: 'example/purescript', goVersion: 'go1.26', arch: 'arm64', exports: {},
    files: [
        { owner: 'ffi/Provider', path: 'Provider.go', content: provider, initializers: [] },
        { owner: 'generated/Consumer', path: 'Consumer.go', content: consumer,
            initializers: [{ owner: 'generated/Consumer', name: 'init/probe' }] },
    ],
};
// Use the compiler's derived ordering of symbol records, including ownership.
const symbolOrd = {
    Eq0: () => ({ eq: a => b => a.owner === b.owner && a.name === b.name }),
    compare: a => b => Symbols.ordGoReference.compare(new Symbols.SymbolReference(a))(new Symbols.SymbolReference(b)),
};
const right = value => { assert.ok(value instanceof Right, value.value0); return value.value0; };
const workspace = t => {
    const dir = mkdtempSync(join(process.env.GOPURS_TEST_EVIDENCE ?? tmpdir(), 'go-source-'));
    if (!process.env.GOPURS_TEST_EVIDENCE) t.after(() => rmSync(dir, { recursive: true, force: true }));
    else console.log(`Go source evidence: ${dir}`);
    return dir;
};
const run = (cwd, command, args, input) => {
    const result = spawnSync(command, args, { cwd, input, encoding: 'utf8', timeout: 120_000,
        maxBuffer: 32 * 1024 * 1024, env: { ...process.env, GOWORK: 'off' } });
    assert.ifError(result.error);
    return result;
};
const context = (files, split = false) => ({
    consumer: split ? 'example/consumer' : 'example/purescript',
    symbols: files.flatMap(file => file.declarations.flatMap(decl => decl.symbols))
        .reduce((map, symbol) => Map.insert(symbolOrd)(symbol)({
            packagePath: split ? symbol.owner === 'ffi/Provider' ? 'example/provider' : 'example/consumer' : 'example/purescript',
            name: symbol.name,
        })(map), Map.empty),
    aliases: Map.singleton('example/provider')('g123abc'),
});

test('typed source binding qualifies types, map keys and callbacks while preserving fields and shadowing', t => {
    const files = bindSources(request)();
    assert.equal(files.length, 2);
    const monolithic = context(files);
    for (const file of files) assert.equal(right(printBoundGoCode(monolithic)(file)), file.text);
    const file = files.find(file => file.path === 'Consumer.go');
    const split = context(files, true);
    const qualified = right(printBoundGoCode(split)(file));
    assert.ok(qualified.includes('g123abc.Box[int]{Key: g123abc.Produce()}'));
    assert.ok(qualified.includes('g123abc.IntMap{g123abc.Key: 10}'));
    assert.ok(qualified.includes('func(Produce func() int) int { return Produce() }'));
    assert.ok(qualified.includes('boxed.(g123abc.Box[int]).Key'));
    assert.ok(qualified.includes('type Wrapped struct { g123abc.Box[int] }'));
    assert.ok(qualified.includes('return value.Box.Read()'));
    assert.ok(qualified.includes('// 😀 Produce Key: never a source reference.'));
    assert.deepEqual(right(boundImports(split)(file)), ['example/provider']);
    const directory = workspace(t);
    for (const child of ['provider', 'consumer', 'broken']) mkdirSync(join(directory, child));
    writeFileSync(join(directory, 'go.mod'), 'module example\n\ngo 1.26\n');
    writeFileSync(join(directory, 'provider/provider.go'), provider.replace('package purescript', 'package provider'));
    writeFileSync(join(directory, 'consumer/consumer.go'), qualified.replace('package purescript', 'package consumer\nimport g123abc "example/provider"'));
    writeFileSync(join(directory, 'main.go'), 'package main\nimport("fmt"; "example/consumer")\nfunc main(){fmt.Print(consumer.Use())}\n');
    // The rejected name-only transformation is retained as a negative oracle.
    const broken = consumer.replace(/\b(Box|IntMap|Key|Produce)\b/g, 'g123abc.$1')
        .replace('package purescript', 'package broken\nimport g123abc "example/provider"');
    writeFileSync(join(directory, 'broken/broken.go'), broken);
    const rejected = run(directory, 'go', ['build', './broken']);
    writeFileSync(join(directory, 'before.log'), rejected.stdout + rejected.stderr);
    assert.notEqual(rejected.status, 0, 'the unscoped baseline must fail');
    const accepted = run(directory, 'go', ['run', '.']);
    writeFileSync(join(directory, 'after.log'), accepted.stdout + accepted.stderr);
    assert.equal(accepted.status, 0, accepted.stderr);
    assert.equal(accepted.stdout, '42');
    writeFileSync(join(directory, 'bindings.json'), JSON.stringify(files, null, 2));
    assert.ok(printBoundGoCode({ ...split, aliases: Map.singleton('example/provider')('Produce') })(file) instanceof Left);
    // An existing import does not prove the alias is visible at every use:
    // a nested callback may bind a local with the same name.
    assert.ok(printBoundGoCode({ ...split, aliases: Map.singleton('example/provider')('Produce') })({ ...file,
        imports: [...file.imports, { path: 'example/provider', alias: 'Produce' }] }) instanceof Left);
    assert.ok(printBoundGoCode({ ...split, symbols: Map.empty })(file) instanceof Left);
    const stale = { ...file, text: file.text.replace('Box[int]', 'Bad[int]') };
    assert.ok(printBoundGoCode(split)(stale) instanceof Left);
});

test('WASM, native and the Rust C bridge return identical complete binding bytes', t => {
    const directory = workspace(t);
    for (const name of ['parser.go', 'types.go', 'native_api.go', 'bindings.go', 'packages.go', 'workspace.go'])
        copyFileSync(new URL(`./ffi-gen/${name}`, import.meta.url), join(directory, name));
    writeFileSync(join(directory, 'go.mod'), 'module probe\n\ngo 1.27.0\n');
    writeFileSync(join(directory, 'main.go'), `package main
import("fmt"; "io"; "os")
func main(){ input, err := io.ReadAll(os.Stdin); if err != nil {panic(err)}; operation := Bind; if len(os.Args)>1 {switch os.Args[1] {case "packages": operation=Packages; case "files": operation=PackageFiles}}; output, err := operation(string(input)); if err != nil {fmt.Fprint(os.Stderr,err); os.Exit(1)}; fmt.Print(output) }
`);
    const compiled = run(directory, 'go', ['build', '-o', join(directory, 'bind'), '.']);
    assert.equal(compiled.status, 0, compiled.stderr);
    const json = JSON.stringify(request);
    const wasm = bindGoSourceImpl(json)().trimEnd();
    const native = run(directory, join(directory, 'bind'), [], json);
    assert.equal(native.status, 0, native.stderr);
    assert.equal(native.stdout, wasm);
    writeFileSync(join(directory, 'request.json'), json);
    writeFileSync(join(directory, 'native.json'), native.stdout);
    writeFileSync(join(directory, 'wasm.json'), wasm);
    const archive = join(directory, 'archive');
    mkdirSync(archive);
    for (const name of ['parser.go', 'types.go', 'native_api.go', 'bindings.go', 'packages.go', 'workspace.go', 'main_carchive.go', 'go.mod'])
        copyFileSync(new URL(`./ffi-gen/${name}`, import.meta.url), join(archive, name));
    const library = run(archive, 'go', ['build', '-tags=carchive', '-buildmode=c-archive', '-o', 'libgopurs_ffi.a', '.']);
    assert.equal(library.status, 0, library.stderr);
    // Compile the production Rust bridge itself. This small harness supplies
    // only the host's Value/effect shell; parsing and C memory ownership are real.
    writeFileSync(join(directory, 'transport.rs'), `
use std::io::Read;
enum Value { String(String), Func1(Func1) }
enum Func1 { Shared(std::rc::Rc<dyn Fn(Value) -> Value>) }
fn purust_string_to_utf8_lossy(value: &str) -> String { value.to_owned() }
fn purust_string_from_utf8(value: &str) -> String { value.to_owned() }
#[allow(non_snake_case)] mod Purs_Effect_Exception {
  pub fn Effect_Exception_error(value: String) -> String { value }
  pub fn purust_exception_raise(value: String) -> ! { panic!("{}", value) }
}
${readFileSync(new URL('../src/Gopurs/FfiSupport.rs', import.meta.url), 'utf8')}
fn main() {
  let mut input = String::new(); std::io::stdin().read_to_string(&mut input).unwrap();
  let operation = match std::env::args().nth(1).as_deref() {Some("packages") => Gopurs_FfiSupport_packageGoSourceImpl, Some("files") => Gopurs_FfiSupport_packageGoFilesImpl, _ => Gopurs_FfiSupport_bindGoSourceImpl};
  let Value::Func1(Func1::Shared(effect)) = operation(input) else { panic!("missing effect") };
  let Value::String(result) = effect(Value::String(String::new())) else { panic!("missing JSON") };
  print!("{}", result);
}
`);
    const rust = run(directory, 'rustc', ['--edition=2021', 'transport.rs', '-L', `native=${archive}`,
        '-l', 'static=gopurs_ffi', ...(process.platform === 'darwin' ? ['-l', 'resolv'] : ['-l', 'pthread', '-l', 'dl', '-l', 'm']), '-o', 'bind-rust']);
    assert.equal(rust.status, 0, rust.stderr);
    const response = run(directory, join(directory, 'bind-rust'), [], json);
    assert.equal(response.status, 0, response.stderr);
    assert.equal(response.stdout, wasm);
    writeFileSync(join(directory, 'rust.json'), response.stdout);

    const dependency = join(directory, 'dependency');
    mkdirSync(dependency);
    writeFileSync(join(dependency, 'go.mod'), 'module example/native\n\ngo 1.26\n');
    writeFileSync(join(dependency, 'native.go'), 'package native\ntype Box[T any] struct{ Value T }; func (b Box[T]) Get() T {return b.Value}\n');
    const exported = run(dependency, 'go', ['list', '-export', '-json', '.']);
    assert.equal(exported.status, 0, exported.stderr);
    const exportFile = JSON.parse(exported.stdout).Export;
    copyFileSync(exportFile, join(directory, 'native.a'));
    const withImport = {
        ...request, files: [{ owner: 'ffi/Bridge', path: 'Bridge.go', initializers: [],
            content: 'package purescript\nimport native "example/native"\nfunc Bridge() int { v := native.Box[int]{Value: 42}; return v.Get() }\n' }],
        exports: { 'example/native': readFileSync(exportFile).toString('base64') },
    };
    for (const [label, data] of [['exports', withImport], ['import-effect', { ...withImport,
        files: [{ ...withImport.files[0], content: 'package purescript\nimport _ "example/native"\nfunc Bridge() int {return 42}\n' }] }]]) {
        const input = JSON.stringify(data), expected = bindGoSourceImpl(input)().trimEnd();
        writeFileSync(join(directory, label + '-request.json'), input);
        writeFileSync(join(directory, label + '-wasm.json'), expected);
        for (const [host, executable] of [['native', 'bind'], ['rust', 'bind-rust']]) {
            const result = run(directory, join(directory, executable), [], input);
            assert.equal(result.status, 0, result.stderr);
            assert.equal(result.stdout, expected, label + ': ' + host);
            writeFileSync(join(directory, label + '-' + host + '.json'), result.stdout);
        }
        const [bound] = bindSources(data)();
        assert.deepEqual(right(boundImports(context([bound]))(bound)), ['example/native']);
        assert.equal(right(printBoundGoCode(context([bound]))(bound)), data.files[0].content);
    }
    const invalid = JSON.stringify({ ...withImport, exports: {} });
    assert.throws(() => bindGoSourceImpl(invalid)(), /missing Go export data: example\/native/);
    for (const executable of ['bind', 'bind-rust']) {
        const result = run(directory, join(directory, executable), [], invalid);
        assert.notEqual(result.status, 0);
        assert.match(result.stderr, /missing Go export data: example\/native/);
    }

    const nativeRecord = Go.goTypeToStr(new Go.TypeRecord([
        new Tuple('value', Go.TypeInt64.value), new Tuple('type', Go.TypeString.value),
        new Tuple('\ud800', Go.TypeInt64.value), new Tuple('💡', Go.TypeBool.value),
    ]));
    assert.match(nativeRecord, /gopurs:record-v1:4:00760061006c00750065,0074007900700065,d800,d83ddca1/);
    const packaged = { ...request,
        files: request.files.map(file => ({ ...file, moduleName: file.owner.split('/')[1], records: [] })),
        entries: [{ owner: 'generated/Consumer', name: 'Use' }],
    };
    packaged.files[0].content += `\nfunc MakeRecord() ${nativeRecord} {return ${nativeRecord}{21,"tag",7,true}}\n`;
    packaged.files[1].content += `\nfunc AcceptRecord(v ${nativeRecord}) int64 {return v.value}\nfunc UseRecord() int64 {return AcceptRecord(MakeRecord())}\n`;
    // Include two fixed anchor identities so the delivered 1/8192 density
    // exercises genuine boundaries without relying on the arbitrary fillers.
    packaged.files[0].content += [...Array.from({length:512}, (_, i) => i), 9972, 11041]
        .map(i => `\nfunc Filler${i}() int {return ${i}}\n`).join('');
    // Keep fillers generated rather than constrained by the entire FFI atom.
    packaged.files[1].content += packaged.files[0].content.slice(packaged.files[0].content.indexOf('\nfunc Filler0'));
    packaged.files[0].content = packaged.files[0].content.slice(0, packaged.files[0].content.indexOf('\nfunc Filler0'));
    const input = JSON.stringify(packaged), expected = packageGoSourceImpl(input)().trimEnd();
    const decoded = packageSources(packaged)();
    assert.ok(decoded.packages.length > 1);
    assert.equal(decoded.format, 'gopurs/go-packages-v1');
    assert.ok(decoded.files.some(file => file.content.includes('PX_d800')));
    assert.ok(decoded.files.every(file => !file.content.includes('/*gopurs:record-')));
    writeFileSync(join(directory, 'packages-request.json'), input);
    writeFileSync(join(directory, 'packages-wasm.json'), expected);
    for (const [host, executable] of [['native', 'bind'], ['rust', 'bind-rust']]) {
        const result = run(directory, join(directory, executable), ['packages'], input);
        assert.equal(result.status, 0, result.stderr);
        assert.equal(result.stdout, expected, 'package bytes: ' + host);
        writeFileSync(join(directory, 'packages-' + host + '.json'), result.stdout);
    }

    const sha256 = data => createHash('sha256').update(data).digest('hex');
    const sourceFiles = packaged.files.map(({ content, ...file }) => {
        const source = join(directory, 'input-' + file.path);
        writeFileSync(source, content);
        return { ...file, source, sha256: sha256(content) };
    });
    const frozen = join(directory, 'frozen');
    mkdirSync(join(frozen, 'tools'), { recursive: true });
    copyFileSync(resolveGoPackagesHelper(), join(frozen, 'tools/go-packages-helper'));
    let reference;
    for (const host of ['js', 'native', 'rust']) {
        const filesRequest = { format: 'gopurs/go-package-files-v1', packagePath: packaged.packagePath,
            goVersion: packaged.goVersion, arch: packaged.arch, files: sourceFiles, exports: [],
            entries: packaged.entries, output: join(directory, 'emitted-' + host) };
        let response;
        if (host === 'js') {
            const previous = process.env.GOPURS_COMPILER_SNAPSHOT;
            try {
                process.env.GOPURS_COMPILER_SNAPSHOT = frozen;
                response = packageGoFilesImpl(JSON.stringify(filesRequest))();
            } finally {
                if (previous === undefined) delete process.env.GOPURS_COMPILER_SNAPSHOT;
                else process.env.GOPURS_COMPILER_SNAPSHOT = previous;
            }
        } else {
            const result = run(directory, join(directory, host === 'rust' ? 'bind-rust' : 'bind'), ['files'], JSON.stringify(filesRequest));
            assert.equal(result.status, 0, result.stderr);
            response = result.stdout;
        }
        if (reference !== undefined) assert.equal(response, reference, host + ': compact file transport');
        else reference = response;
        writeFileSync(join(directory, 'files-' + host + '.json'), response);
        const manifest = JSON.parse(response);
        assert.equal(manifest.planHash, decoded.planHash);
        assert.deepEqual(manifest.packages, decoded.packages);
        for (const [i, file] of manifest.files.entries()) {
            const content = readFileSync(join(filesRequest.output, file.path));
            assert.equal(content.toString('utf8'), decoded.files[i].content);
            assert.equal(file.sha256, sha256(content));
            assert.equal(file.bytes, content.length);
            assert.ok(!Object.hasOwn(file, 'content'));
        }
    }
});
