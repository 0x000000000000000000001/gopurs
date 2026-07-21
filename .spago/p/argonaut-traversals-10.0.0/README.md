# Argonaut Traversals

[![CI](https://github.com/purescript-contrib/purescript-argonaut-traversals/workflows/CI/badge.svg?branch=main)](https://github.com/purescript-contrib/purescript-argonaut-traversals/actions?query=workflow%3ACI+branch%3Amain)
[![Release](http://img.shields.io/github/release/purescript-contrib/purescript-argonaut-traversals.svg)](https://github.com/purescript-contrib/purescript-argonaut-traversals/releases)
[![Pursuit](http://pursuit.purescript.org/packages/purescript-argonaut-traversals/badge)](http://pursuit.purescript.org/packages/purescript-argonaut-traversals)
[![Maintainer: garyb](https://img.shields.io/badge/maintainer-garyb-teal.svg)](http://github.com/garyb)

Prisms, traversals, and zipper for the `Json` type. Part of the [Argonaut](https://github.com/purescript-contrib/purescript-argonaut) collection of libraries for working with JSON in PureScript.

## Installation

Install `argonaut-traversals` with [Spago](https://github.com/purescript/spago):

```sh
spago install argonaut-traversals
```

or install as part of the Argonaut bundle:

```sh
spago install argonaut
```

## Quick start

You can use the prisms defined in `Data.Argonaut.Prisms` with functions from the [`profunctor-lenses`](https://github.com/purescript-contrib/purescript-profunctor-lenses) library to work with nested `Json` structures. For example:

```js
// FFI file
exports.sampleJson = { "a": { "b": [ 10, 11, 12 ] } }
```

```purs
module Main where

import Prelude

import Effect (Effect)
import Data.Argonaut.Core (Json)
import Data.Argonaut.Prisms (_Array, _Number, _Object)
import Data.Maybe (Maybe(..))
import Data.Lens (preview)
import Data.Lens.Index (ix)
import Effect.Console (log)

foreign import sampleJson :: Json

main :: Effect Unit
main =
  -- Walk through an object at the key 'a', then an object at the key 'b', then
  -- get the first index of an array as a number.
  case preview (_Object <<< ix "a" <<< _Object <<< ix "b" <<< _Array <<< ix 0 <<< _Number) sampleJson of
    Nothing -> log "nothin' there"
    Just v -> log $ "This should be 10.0 " <> show v
```

You may also be interested in other libraries in the Argonaut ecosystem:

- [purescript-argonaut-core](https://github.com/purescript-contrib/purescript-argonaut-core) defines the `Json` type, along with basic parsing, printing, and folding functions
- [purescript-argonaut-codecs](https://github.com/purescript-contrib/purescript-argonaut-codecs) provides codecs based on `EncodeJson` and `DecodeJson` type classes, along with instances for common data types and combinators for encoding and decoding `Json` values.
- [purescript-codec-argonaut](https://github.com/garyb/purescript-codec-argonaut) supports an alternative approach for codecs, which are based on profunctors instead of type classes.
- [purescript-argonaut-generic](https://github.com/purescript-contrib/purescript-argonaut-generic) supports generic encoding and decoding for any type with a `Generic` instance.

## Documentation

`argonaut-traversals` documentation is stored in a few places:

1. Module documentation is [published on Pursuit](https://pursuit.purescript.org/packages/purescript-argonaut-traversals).
2. Written documentation is kept in [the docs directory](./docs).
3. Usage examples can be found in [the test suite](./test).

If you get stuck, there are several ways to get help:

- [Open an issue](https://github.com/purescript-contrib/purescript-argonaut-traversals/issues) if you have encountered a bug or problem.
- Ask general questions on the [PureScript Discourse](https://discourse.purescript.org) forum or the [PureScript Discord](https://purescript.org/chat) chat.

## Contributing

You can contribute to `argonaut-traversals` in several ways:

1. If you encounter a problem or have a question, please [open an issue](https://github.com/purescript-contrib/purescript-argonaut-traversals/issues). We'll do our best to work with you to resolve or answer it.

2. If you would like to contribute code, tests, or documentation, please [read the contributor guide](./CONTRIBUTING.md). It's a short, helpful introduction to contributing to this library, including development instructions.

3. If you have written a library, tutorial, guide, or other resource based on this package, please share it on the [PureScript Discourse](https://discourse.purescript.org)! Writing libraries and learning resources are a great way to help this library succeed.
