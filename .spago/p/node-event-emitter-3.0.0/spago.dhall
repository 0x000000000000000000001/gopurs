{ name = "node-event-emitters"
, dependencies =
  [ "effect"
  , "either"
  , "functions"
  , "maybe"
  , "nullable"
  , "prelude"
  , "unsafe-coerce"
  ]
, packages = ./packages.dhall
, sources = [ "src/**/*.purs" ]
}
