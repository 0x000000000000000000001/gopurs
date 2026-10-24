#!/usr/bin/env python3
"""Compile the projection capsule regression through the typed native toolchain."""
from pathlib import Path
import argparse, hashlib, json, os, re, subprocess, time

parser = argparse.ArgumentParser(description=__doc__)
parser.add_argument('--native-workspace', type=Path, required=True,
                    help='Retained build:native workspace containing output/cache-db.json and typed-bin/purs')
parser.add_argument('--backend', type=Path, required=True, help='Explicit native gopurs executable')
parser.add_argument('--work-dir', type=Path, required=True, help='New directory for sources, logs and generated Go')
parser.add_argument('--go', default='go', help='Go executable (default: PATH)')
parser.add_argument('--fixture-dir', type=Path, help='Alternative fixture with Main.purs and adjacent foreign files')
parser.add_argument('--pass-marker', help='Required marker for an alternative fixture')
args = parser.parse_args()
project = Path(__file__).resolve().parents[3]
base = args.native_workspace.resolve()
backend = args.backend.resolve()
workspace = args.work_dir.resolve()
workspace.mkdir(parents=True, exist_ok=False)
sources = workspace / 'src'
sources.mkdir(parents=True, exist_ok=True)
cache = json.loads((base / 'output/cache-db.json').read_text())
paths = {name: next((base / p for p in files if p.endswith('.purs') and (base / p).is_file()), None) for name, files in cache.items()}
# Prefer current project sources over the retained dependency inventory. The
# closure below copies only imports actually needed by this regression.
for source in (project / 'src').rglob('*.purs'):
    declaration = re.search(r'^module\s+([A-Z][A-Za-z0-9_.]*)', source.read_text(), re.M)
    if declaration:
        paths[declaration.group(1)] = source
fixture = (args.fixture_dir.resolve() if args.fixture_dir else project / 'tools/fixtures/projection-capsule') / 'Main.purs'
main_module = re.search(r'^module\s+([A-Z][A-Za-z0-9_.]*)', fixture.read_text(), re.M).group(1)
paths[main_module] = fixture
selected = {}
def collect(name):
    if name == 'Prim' or name.startswith('Prim.') or name in selected:
        return
    source = paths.get(name)
    if source is None:
        raise RuntimeError(f'Missing source for {name}')
    text = source.read_text()
    selected[name] = source
    for imported in re.findall(r'^import\s+([A-Z][A-Za-z0-9_.]*)', text, re.M):
        collect(imported)
collect(main_module)
# NativeMaps' shared Go bridge names the TCO ADT even when only its map helpers
# occur in the PS import closure. Retain that bridge's explicit type dependency.
collect('PureScript.Backend.Optimizer.Codegen.Tco')
captured = []
for name, source in sorted(selected.items()):
    for suffix in ['.purs', '.js', '.go']:
        original = source.with_suffix(suffix)
        if not original.is_file():
            continue
        target = sources / (name + suffix)
        data = original.read_bytes()
        target.write_bytes(data)
        captured.append({'module': name, 'source': str(original.resolve()), 'copy': str(target), 'sha256': hashlib.sha256(data).hexdigest()})
(workspace / 'source-state.json').write_text(json.dumps(captured, indent=2) + '\n')
compiler = (base / 'typed-bin/purs').resolve()
env = {key: value for key, value in os.environ.items() if not key.startswith('GOPURS_') and key not in ['GOGC', 'GOMEMLIMIT', 'PPROF', 'GOFLAGS', 'GOMAXPROCS']}
env.update(TMPDIR='/private/tmp', GOCACHE='/private/tmp/gopurs-build-cache-gocache', GOWORK='off', GOPURS_FORCE_BUILD='1', GOPURS_BUILD_MANIFEST='0', GOPURS_GO_PACKAGES='0', GOPURS_PREPARE_JOBS='8', GOPURS_PBO_JOBS='8', GOPURS_EMIT_JOBS='8', GOGC='200', GOMEMLIMIT='10GiB')
commands = []
def command(label, args, cwd=workspace):
    started = time.monotonic()
    with (workspace / (label + '.stdout')).open('w') as out, (workspace / (label + '.stderr')).open('w') as err:
        result = subprocess.run([str(x) for x in args], cwd=cwd, env=env, stdout=out, stderr=err, timeout=300)
    commands.append({'label': label, 'args': [str(x) for x in args], 'cwd': str(cwd), 'seconds': time.monotonic() - started, 'returncode': result.returncode})
    (workspace / 'commands.json').write_text(json.dumps(commands, indent=2) + '\n')
    print(label, result.returncode, round(commands[-1]['seconds'], 3), flush=True)
    if result.returncode:
        raise SystemExit(result.returncode)

command('frontend', [compiler, 'compile', '--codegen', 'corefn', '--output', workspace / 'output', *sorted(sources.glob('*.purs'))])
typed = []
for path in sorted((workspace / 'output').glob('*/corefn.json')):
    data = json.loads(path.read_text())
    assert all(isinstance(data.get(key), list) for key in ['typeTable', 'dataDecls', 'classDecls']), path
    typed.append({'module': path.parent.name, 'types': len(data['typeTable'])})
(workspace / 'typed.json').write_text(json.dumps(typed, indent=2) + '\n')
command('native-support', ['node', project / 'tools/prepare-native-output.mjs', workspace / 'output'])
command('generate-go', [backend, '--main', main_module])
go = args.go
command('go-build', [go, 'build', '-p', '2', '-o', workspace / 'lookup-native-ps', './main/main.go'], workspace / 'output')
command('run', [workspace / 'lookup-native-ps'])
text = (workspace / 'run.stdout').read_text()
assert (args.pass_marker or 'PROJECTION-CAPSULE-NATIVE-PASS') in text
if args.fixture_dir is None:
    for jobs in [1, 8]:
        assert f'CAPSULE-DISPATCH-NATIVE-PASS jobs={jobs} started={jobs} finished={jobs} active=0' in text, text
generated = {}
for name in ['Gopurs.Driver.Cache', 'Gopurs.Cache.HeaderPreload', 'Gopurs.BuildCache', 'Gopurs.Driver.Prepare.ProjectionCapsule', 'Gopurs.Preparation', main_module]:
    path = workspace / 'output/purescript' / (name.replace('.', '_') + '.go')
    if path.is_file():
        generated[name] = hashlib.sha256(path.read_bytes()).hexdigest()
report = {'status': 'passed', 'scope': main_module + ' current PS and FFI through typed compiler and explicit native generator', 'compiler_sha256': hashlib.sha256(compiler.read_bytes()).hexdigest(), 'backend_sha256': hashlib.sha256(backend.read_bytes()).hexdigest(), 'source_state_sha256': hashlib.sha256((workspace / 'source-state.json').read_bytes()).hexdigest(), 'generated_sources': generated, 'binary_sha256': hashlib.sha256((workspace / 'lookup-native-ps').read_bytes()).hexdigest(), 'modules': len(typed), 'commands': commands}
(workspace / 'result.json').write_text(json.dumps(report, indent=2) + '\n')
