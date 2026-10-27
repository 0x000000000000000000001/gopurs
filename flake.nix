{
  description = "gopurs, a PureScript to Go compiler backend";
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
    flake-compat = { url = "github:edolstra/flake-compat"; flake = false; };
    purescript-overlay = {
      url = "github:thomashoneyman/purescript-overlay";
      inputs.nixpkgs.follows = "nixpkgs";
    };
    mkSpagoDerivation = {
      url = "github:jeslie0/mkSpagoDerivation";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.ps-overlay.follows = "purescript-overlay";
    };
    purs-tast = {
      url = "github:0x000000000000000000001/purescript/105f5752b39cc0f6c99a40064295338bf57ad223";
      flake = false;
    };
  };
  outputs = { self, nixpkgs, ... }@inputs:
    let
      systems = [ "x86_64-linux" "aarch64-linux" "aarch64-darwin" ];
      forAllSystems = nixpkgs.lib.genAttrs systems;
      pkgsFor = system: import nixpkgs { inherit system; overlays = builtins.attrValues self.overlays; };
      toolsFor = system: import ./nix/toolchains.nix { pkgs = pkgsFor system; };
      bootstrap = builtins.fromJSON (builtins.readFile ./bootstrap-sources.json);
    in
    assert inputs.purs-tast.rev == bootstrap.compiler.rev; {
      overlays = {
        purescript = inputs.purescript-overlay.overlays.default;
        spago = inputs.mkSpagoDerivation.overlays.default;
      };
      packages = forAllSystems (system:
        let
          pkgs = pkgsFor system;
          toolchains = toolsFor system;
          purs-tast = import ./nix/purs-tast.nix { inherit pkgs; src = inputs.purs-tast; };
          gopurs = import ./nix/gopurs.nix { inherit pkgs toolchains; };
        in
        {
          inherit gopurs purs-tast;
          inherit (toolchains) nodejs go;
          default = gopurs;
        });
      apps = forAllSystems (system:
        let
          app = {
            type = "app";
            program = "${self.packages.${system}.gopurs}/bin/gopurs";
            meta.description = "Compile typed PureScript to Go";
          };
        in
        {
          gopurs = app;
          default = app;
          purs = {
            type = "app";
            program = "${self.packages.${system}.purs-tast}/bin/purs";
            meta.description = "PureScript compiler with typed TAST output";
          };
        });
      checks = forAllSystems (system: import ./nix/checks.nix {
        pkgs = pkgsFor system;
        packages = self.packages.${system};
      });
      devShells = forAllSystems (system:
        let pkgs = pkgsFor system; toolchains = toolsFor system;
        in {
          default = pkgs.mkShell {
            name = "gopurs";
            packages = [
              self.packages.${system}.purs-tast
              pkgs.spago-bin.spago-1_0_3
              pkgs.purs-tidy
              pkgs.purescript-language-server
              toolchains.nodejs
              toolchains.go
              pkgs.esbuild
              pkgs.gopls
              pkgs.git
              pkgs.python3
              pkgs.nixpkgs-fmt
            ];
            GOPURS_PURS = "${self.packages.${system}.purs-tast}/bin/purs";
            GOTOOLCHAIN = "local";
            shellHook = ''
              echo "gopurs: pinned TAST frontend, Go 1.27.0, Node 24.8.0, Spago 1.0.3"
              echo "npm ci builds the JS backend with its separate purs/Spago lockfile."
            '';
          };
        });
      formatter = forAllSystems (system: (pkgsFor system).nixpkgs-fmt);
    };
}
