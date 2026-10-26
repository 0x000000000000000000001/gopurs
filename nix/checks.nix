{ pkgs, packages }:
{
  build = packages.gopurs;
  smoke = pkgs.runCommand "gopurs-wrapper-smoke" { } ''
    export HOME=$TMPDIR/home
    mkdir -p "$HOME"
    if ${packages.gopurs}/bin/gopurs 2>error.log; then
      echo "gopurs unexpectedly succeeded without TAST input" >&2
      exit 1
    fi
    cat error.log
    grep -q '\[gopurs\] error:.*output' error.log
    touch $out
  '';
  application = pkgs.runCommand "gopurs-tast-go-application"
    {
      nativeBuildInputs = [ packages.purs-tast packages.gopurs packages.go packages.nodejs ];
    } ''
    export HOME=$TMPDIR/home GOCACHE=$TMPDIR/go-cache GOTOOLCHAIN=local
    mkdir -p "$HOME" src
    cat >src/Effect.purs <<'PURS'
    module Effect where
    foreign import data Effect :: Type -> Type
    PURS
    cat >src/Main.purs <<'PURS'
    module Main where
    import Effect (Effect)
    data Unit = Unit
    foreign import main :: Effect Unit
    PURS
    cat >src/Main.go <<'GO'
    package Main
    import "fmt"
    func Main(_ any) any { fmt.Println("Hello from gopurs Nix"); return nil }
    GO
    purs compile 'src/*.purs' --codegen corefn
    node --input-type=module <<'JS'
    import assert from 'node:assert/strict';
    import { readFileSync } from 'node:fs';
    const tast = JSON.parse(readFileSync('output/Main/corefn.json', 'utf8'));
    for (const key of ['typeTable', 'dataDecls', 'classDecls']) assert.ok(key in tast, key);
    JS
    for layout in 0 1; do
      GOPURS_GO_PACKAGES=$layout gopurs --main Main --force
      (cd output && go build -trimpath -o ../hello ./main)
      test "$(./hello)" = 'Hello from gopurs Nix'
    done
    touch $out
  '';
}
