import Language.PureScript.Parser
import Language.PureScript.AST
import qualified Data.Text.IO as T
import System.Environment

main = do
  txt <- T.readFile "src/Main.purs"
  case parseModuleFromFile id ("src/Main.purs", txt) of
    Right (_, m) -> print m
    Left err -> print err
