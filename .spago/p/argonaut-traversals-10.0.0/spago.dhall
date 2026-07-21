{ name = "argonaut-traversals"
, dependencies =
  [ "argonaut-codecs"
  , "argonaut-core"
  , "arrays"
  , "console"
  , "control"
  , "effect"
  , "either"
  , "foldable-traversable"
  , "foreign-object"
  , "gen"
  , "integers"
  , "lists"
  , "maybe"
  , "prelude"
  , "profunctor-lenses"
  , "quickcheck"
  , "strings"
  , "tailrec"
  , "tuples"
  , "unfoldable"
  ]
, packages = ./packages.dhall
, sources = [ "src/**/*.purs", "test/**/*.purs" ]
}
