# Generated with cabal2nix 2.21.3 from the TAST compiler revision
# 105f5752b39cc0f6c99a40064295338bf57ad223. The source is supplied by
# the locked flake input; keep this dependency description in sync with it.
{ src
, mkDerivation
, aeson
, aeson-better-errors
, ansi-terminal
, array
, base
, blaze-html
, bower-json
, boxes
, bytestring
, Cabal
, cborg
, cheapskate
, clock
, containers
, cryptonite
, data-ordlist
, deepseq
, directory
, dlist
, edit-distance
, exceptions
, file-embed
, filepath
, generic-random
, gitrev
, Glob
, happy
, haskeline
, hspec
, hspec-discover
, HUnit
, language-javascript
, lens
, lib
, lifted-async
, lifted-base
, memory
, monad-control
, monad-logger
, monoidal-containers
, mtl
, network
, newtype
, optparse-applicative
, parallel
, parsec
, prettyprinter
, prettyprinter-ansi-terminal
, process
, protolude
, QuickCheck
, regex-base
, regex-tdfa
, safe
, scientific
, semialign
, semigroups
, serialise
, sourcemap
, split
, stm
, stringsearch
, template-haskell
, text
, these
, time
, transformers
, transformers-base
, typed-process
, utf8-string
, vector
, witherable
}:
assert src.rev == "105f5752b39cc0f6c99a40064295338bf57ad223";
mkDerivation {
  pname = "purescript";
  version = "0.15.16";
  inherit src;
  isLibrary = true;
  isExecutable = true;
  libraryHaskellDepends = [
    aeson
    aeson-better-errors
    ansi-terminal
    array
    base
    blaze-html
    bower-json
    boxes
    bytestring
    Cabal
    cborg
    cheapskate
    clock
    containers
    cryptonite
    data-ordlist
    deepseq
    directory
    dlist
    edit-distance
    file-embed
    filepath
    Glob
    haskeline
    language-javascript
    lens
    lifted-async
    lifted-base
    memory
    monad-control
    monad-logger
    monoidal-containers
    mtl
    parallel
    parsec
    process
    protolude
    regex-tdfa
    safe
    scientific
    semialign
    semigroups
    serialise
    sourcemap
    stm
    stringsearch
    template-haskell
    text
    these
    time
    transformers
    transformers-base
    utf8-string
    vector
    witherable
  ];
  libraryToolDepends = [ happy ];
  executableHaskellDepends = [
    aeson
    aeson-better-errors
    ansi-terminal
    array
    base
    blaze-html
    bower-json
    boxes
    bytestring
    Cabal
    cborg
    cheapskate
    clock
    containers
    cryptonite
    data-ordlist
    deepseq
    directory
    dlist
    edit-distance
    exceptions
    file-embed
    filepath
    gitrev
    Glob
    haskeline
    language-javascript
    lens
    lifted-async
    lifted-base
    memory
    monad-control
    monad-logger
    monoidal-containers
    mtl
    network
    optparse-applicative
    parallel
    parsec
    prettyprinter
    prettyprinter-ansi-terminal
    process
    protolude
    regex-tdfa
    safe
    scientific
    semialign
    semigroups
    serialise
    sourcemap
    stm
    stringsearch
    template-haskell
    text
    these
    time
    transformers
    transformers-base
    utf8-string
    vector
    witherable
  ];
  testHaskellDepends = [
    aeson
    aeson-better-errors
    ansi-terminal
    array
    base
    blaze-html
    bower-json
    boxes
    bytestring
    Cabal
    cborg
    cheapskate
    clock
    containers
    cryptonite
    data-ordlist
    deepseq
    directory
    dlist
    edit-distance
    file-embed
    filepath
    generic-random
    Glob
    haskeline
    hspec
    HUnit
    language-javascript
    lens
    lifted-async
    lifted-base
    memory
    monad-control
    monad-logger
    monoidal-containers
    mtl
    newtype
    parallel
    parsec
    process
    protolude
    QuickCheck
    regex-base
    regex-tdfa
    safe
    scientific
    semialign
    semigroups
    serialise
    sourcemap
    split
    stm
    stringsearch
    template-haskell
    text
    these
    time
    transformers
    transformers-base
    typed-process
    utf8-string
    vector
    witherable
  ];
  testToolDepends = [ hspec-discover ];
  doCheck = false;
  homepage = "http://www.purescript.org/";
  description = "PureScript Programming Language Compiler";
  license = lib.meta.getLicenseFromSpdxId "BSD-3-Clause";
  mainProgram = "purs";
}
