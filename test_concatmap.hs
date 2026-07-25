import Data.List.NonEmpty (NonEmpty(..))
import qualified Data.List.NonEmpty as NEL
import Data.Foldable (concatMap)

data Decl = DataDecl String | OtherDecl String deriving Show

dataDeclToCoreFn :: Decl -> [String]
dataDeclToCoreFn (DataDecl name) = [name]
dataDeclToCoreFn (OtherDecl _) = []

main = print $ concatMap dataDeclToCoreFn (DataDecl "foo" :| [OtherDecl "bar", DataDecl "baz"])
