{ pkgs }:
let
  inherit (pkgs) lib;
  system = pkgs.stdenv.hostPlatform.system;
  # Official distribution SHA-256 values, pinned independently of nixpkgs.
  # https://nodejs.org/dist/v24.8.0/SHASUMS256.txt
  # https://go.dev/dl/?mode=json&include=all
  archives = {
    x86_64-linux = {
      node = "linux-x64.tar.xz";
      nodeHash = "2598641d188b41793930917f1a99a81c9615856b4205d408a44ab676c1acbb3d";
      go = "linux-amd64.tar.gz";
      goHash = "675c26c449cbb18fc24b74650de1eabbae6e16f64326fd85a283fb3b58280685";
    };
    aarch64-linux = {
      node = "linux-arm64.tar.xz";
      nodeHash = "323e6bc34bfe82a21f42d9eb94d9a2b6c8082d8c03cb510cf2f95dd1dcaa1531";
      go = "linux-arm64.tar.gz";
      goHash = "51798d2c42d0e1c6ed7fd9f48728b4193abac9e8aad6dbac2fe96a81f5909bda";
    };
    aarch64-darwin = {
      node = "darwin-arm64.tar.gz";
      nodeHash = "d81191a1866760eb918caa976c023036bc1fc7405ea31b148905211522045767";
      go = "darwin-arm64.tar.gz";
      goHash = "90493b3bbd5e10f91d12153198bf1994fd756399b4fec93b49b0c6e2acdeeb3e";
    };
  }.${system};
in
{
  nodejs = pkgs.stdenv.mkDerivation {
    pname = "nodejs-gopurs";
    version = "24.8.0";
    src = pkgs.fetchurl {
      url = "https://nodejs.org/dist/v24.8.0/node-v24.8.0-${archives.node}";
      sha256 = archives.nodeHash;
    };
    nativeBuildInputs = lib.optionals pkgs.stdenv.isLinux [ pkgs.autoPatchelfHook ];
    buildInputs = lib.optionals pkgs.stdenv.isLinux [ pkgs.stdenv.cc.cc.lib ];
    dontBuild = true;
    installPhase = ''
      mkdir -p $out
      cp -R bin include lib share $out/
      mkdir -p $out/share/doc/nodejs
      cp LICENSE $out/share/doc/nodejs/LICENSE
    '';
    meta = { license = lib.licenses.mit; };
  };
  go = pkgs.stdenvNoCC.mkDerivation {
    pname = "go-gopurs";
    version = "1.27.0";
    src = pkgs.fetchurl {
      url = "https://go.dev/dl/go1.27.0.${archives.go}";
      sha256 = archives.goHash;
    };
    nativeBuildInputs = [ pkgs.makeWrapper ];
    dontBuild = true;
    dontStrip = true;
    installPhase = ''
      mkdir -p $out/share/go $out/bin
      cp -R . $out/share/go/
      makeWrapper $out/share/go/bin/go $out/bin/go --set GOROOT $out/share/go
      makeWrapper $out/share/go/bin/gofmt $out/bin/gofmt
    '';
    meta = { license = lib.licenses.bsd3; };
  };
}
