{ pkgs, toolchains }:
let
  manifest = builtins.fromJSON (builtins.readFile ../package.json);
  source = pkgs.lib.fileset.toSource {
    root = ../.;
    fileset = pkgs.lib.fileset.unions [
      ../src
      ../runtime
      ../tools
      ../docs
      ../package.json
      ../spago.yaml
      ../spago.lock
      ../LICENSE
      ../README.md
      ../CHANGELOG.md
      ../CODE_OF_CONDUCT.md
      ../CONTRIBUTING.md
      ../NOTICE
      ../SECURITY.md
      ../THIRD_PARTY_NOTICES.md
      ../bin/gopurs
    ];
  };
in
pkgs.mkSpagoDerivation {
  pname = "gopurs";
  version = manifest.version;
  src = source;
  spagoYaml = ../spago.yaml;
  spagoLock = ../spago.lock;
  nativeBuildInputs = [
    pkgs.purs-bin.purs-0_15_16
    pkgs.spago-bin.spago-0_93_45
    toolchains.nodejs
    toolchains.go
    pkgs.esbuild
    pkgs.makeWrapper
  ];
  GOTOOLCHAIN = "local";
  buildPhase = ''
    export GOCACHE=$TMPDIR/go-cache
    node tools/embed-runtime.mjs
    spago build --offline
    spago bundle --module Main --platform node --outfile bin/gopurs.js --bundle-type app --offline
    # The Nix store is immutable. Build the content-addressed helpers now,
    # then install the same source bytes and their ready caches.
    node tools/build-cache.mjs --helper
    node tools/go-packages.mjs
  '';
  installPhase = ''
    export destination=$out/libexec/gopurs
    mkdir -p "$destination" $out/bin
    node --input-type=module <<'JS'
    import { cpSync, mkdirSync, readFileSync } from 'node:fs';
    import { dirname, join } from 'node:path';
    const manifest = JSON.parse(readFileSync('package.json', 'utf8'));
    for (const path of [...manifest.files, 'package.json', 'LICENSE', 'README.md']) {
      const target = join(process.env.destination, path);
      mkdirSync(dirname(target), { recursive: true });
      cpSync(path, target, { recursive: true });
    }
    JS
    mkdir -p "$destination/node_modules"
    cp -R node_modules/.cache "$destination/node_modules/"
    makeWrapper ${toolchains.nodejs}/bin/node $out/bin/gopurs \
      --add-flags "$destination/bin/gopurs.js"
  '';
  meta = {
    description = "PureScript compiler backend targeting Go";
    homepage = "https://github.com/0x000000000000000000001/gopurs";
    license = pkgs.lib.licenses.mit;
    mainProgram = "gopurs";
  };
}
