# Argonaut codecs bootstrap input

This is the Go adaptation of `purescript-argonaut-codecs` 9.1.0 used to
bootstrap the native gopurs compiler. It is included in the source release so
the build does not depend on an unversioned sibling checkout.

Upstream: https://github.com/purescript-contrib/purescript-argonaut-codecs

The PureScript and JavaScript sources derive from the MIT-licensed upstream
package. The Go foreign implementations and the accompanying tests were
developed in the gopurs workspace. The original license is retained in
`LICENSE`; gopurs contributions use the repository's MIT license.

The `src/`, `test/` and `spago.yaml` contents were copied from the validated
`gopurs-argonaut-codecs` adaptation on 26 October 2026. The two JavaScript test
runners have only their relative compiler paths adapted to this directory.
The package configuration declares an independent workspace so Spago does not
auto-discover its tests while building the JavaScript compiler. The source and
test cases are unchanged. This directory is a build input, not a separately
published npm or registry package.
