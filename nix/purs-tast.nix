{ pkgs, src }:
let
  # Match the fork's GHC 9.8 series. Nix locks the complete dependency set;
  # stack.yaml is the separate, non-Nix development build plan.
  haskell = pkgs.haskell.packages.ghc984.override {
    overrides = self: super: {
      language-javascript = self.language-javascript_0_7_0_0;
    };
  };
  compiler = haskell.callPackage ./purs-tast-package.nix { inherit src; };
in
pkgs.haskell.lib.justStaticExecutables (pkgs.haskell.lib.overrideCabal
  (pkgs.haskell.lib.doJailbreak compiler)
  (old: {
    configureFlags = (old.configureFlags or [ ]) ++ [ "-frelease" ];
    doCheck = false;
    doHaddock = false;
    enableLibraryProfiling = false;
    description = "PureScript fork preserving typed TAST for AOT backends";
  }))
