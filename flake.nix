{
  description = "gopurs, a PureScript to Go compiler backend";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";

    # Only used by shell.nix for non-flake users.
    flake-compat = {
      url = "github:edolstra/flake-compat";
      flake = false;
    };

    # Canonical PureScript toolchain overlay (purs, spago, purs-tidy,
    # purs-backend-es, purescript-language-server).
    purescript-overlay = {
      url = "github:thomashoneyman/purescript-overlay";
      inputs.nixpkgs.follows = "nixpkgs";
    };

    # Overlay providing `mkSpagoDerivation`. Kept for parity with the phpurs
    # flake and for a future `packages` output: the current backend build is
    # not a pure derivation yet because spago.yaml references sibling
    # checkouts (`../../purescript-backend-optimizer-gopurs`, `../gopurs-st`,
    # `../gopurs-unsafe-coerce`) and applications need the TAST-capable
    # compiler fork, which is not packaged.
    mkSpagoDerivation = {
      url = "github:jeslie0/mkSpagoDerivation";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.ps-overlay.follows = "purescript-overlay";
    };
  };

  outputs = { self, nixpkgs, ... }@inputs:
    let
      supportedSystems = [ "x86_64-linux" "aarch64-linux" "aarch64-darwin" ];

      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;

      nixpkgsFor = forAllSystems (system:
        import nixpkgs {
          inherit system;
          overlays = builtins.attrValues self.overlays;
        });
    in
    {
      overlays = {
        purescript = inputs.purescript-overlay.overlays.default;
        spago = inputs.mkSpagoDerivation.overlays.default;
      };

      devShells = forAllSystems (system:
        let
          pkgs = nixpkgsFor.${system};
        in
        {
          default = pkgs.mkShell {
            name = "gopurs";

            packages = [
              # PureScript toolchain. Upstream `purs` is deliberately not
              # included: gopurs consumes the enriched TAST produced by the
              # compiler fork from ../../purescript (see README.md), and a
              # nix-provided purs would shadow it inside `nix develop`.
              pkgs.spago-unstable
              pkgs.purs-tidy
              pkgs.purescript-language-server

              # Building and running the backend. Full nodejs on purpose:
              # bin/test invokes `npm run build`, which needs npm/npx.
              pkgs.nodejs_24
              pkgs.esbuild

              # Generated Go output and the compiler's own native build.
              # Go 1.27.0 is required to rebuild the FFI parser; ordinary
              # builds use the checked-in WASM/JavaScript runtime.
              pkgs.go
              pkgs.gopls

              # Test and scratch helpers.
              pkgs.git
              pkgs.python3

              pkgs.nixpkgs-fmt
            ];

            shellHook = ''
              echo "gopurs dev shell"
              echo "  - sibling checkouts required: ../../purescript, ../../purescript-backend-optimizer-gopurs, ../gopurs-*"
              echo "  - put the TAST-capable purs fork (built from ../../purescript) ahead of any other purs in PATH"
              if ! command -v purs >/dev/null 2>&1; then
                echo "  - warning: no purs found; spago build needs the TAST-capable fork" >&2
              fi
            '';
          };
        });

      formatter = forAllSystems (system: nixpkgsFor.${system}.nixpkgs-fmt);
    };
}
