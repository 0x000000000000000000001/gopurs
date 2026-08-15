module TestSanitize where
import Prelude
import Effect.Console (log)
import Data.String as String
import Data.String.Pattern (Pattern(..))

sanitizeName :: String -> String
sanitizeName name =
  let
    escapeOpChar = case _ of
      "/" -> "_slash_"
      "\\" -> "_bslash_"
      "<" -> "_less_"
      ">" -> "_greater_"
      "=" -> "_eq_"
      "+" -> "_plus_"
      "-" -> "_minus_"
      "*" -> "_times_"
      ":" -> "_colon_"
      "|" -> "_bar_"
      "&" -> "_amp_"
      "^" -> "_caret_"
      "~" -> "_tilde_"
      "?" -> "_qmark_"
      "!" -> "_bang_"
      "@" -> "_at_"
      "#" -> "_hash_"
      "%" -> "_percent_"
      "." -> "_dot_"
      "\"" -> "_quote_"
      "'" -> "_prime_"
      "$" -> "_dollar_"
      char -> char
    s1 = String.joinWith "" (map escapeOpChar (String.split (Pattern "") name))
  in s1

main = do
  log (sanitizeName "test/\\")
