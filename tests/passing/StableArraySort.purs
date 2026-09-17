-- @dependencies: prelude arrays effect console
module Main where

import Prelude

import Data.Array as Array
import Effect (Effect)
import Effect.Console (log)

type Entry = { key :: Int, name :: String }

sortEntries :: Array Entry -> Array Entry
sortEntries = Array.sortBy (comparing _.key)

defaultContext :: { depth :: Int, tcoIdent :: Boolean, options :: { isTail :: Boolean }, bound :: Int }
defaultContext = { depth: 0, tcoIdent: false, options: { isTail: false }, bound: 0 }

updateContext :: Int -> { depth :: Int, tcoIdent :: Boolean, options :: { isTail :: Boolean }, bound :: Int }
updateContext value = defaultContext { tcoIdent = true, options = { isTail: true }, bound = value }

check :: String -> Boolean -> Effect Unit
check name condition = log (if condition then name else "Fail: " <> name)

main :: Effect Unit
main = do
  let
    input = [ { key: 2, name: "new" }, { key: 2, name: "old" }, { key: 1, name: "first" } ]
    result = sortEntries input
    context = updateContext 42
  check "stable duplicate keys" (map _.name result == [ "first", "new", "old" ])
  check "persistent input" (map _.name input == [ "new", "old", "first" ])
  check "constant record updates" (context.tcoIdent && context.options.isTail && context.bound == 42)
  log "Done"
