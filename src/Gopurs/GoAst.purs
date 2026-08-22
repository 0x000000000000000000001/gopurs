module Gopurs.GoAst where

import Prelude
import Data.Tuple (Tuple(..))
import Data.String as String
import Data.String (Pattern(..), Replacement(..))
import PureScript.Backend.Optimizer.CoreFn (ExprType(..))
import Data.Maybe (Maybe(..))
import Data.Set as Set
import Data.Foldable (fold)
import Data.Array as Array
data GoExpr
  = GoVar String
  | GoString String
  | GoInt Int
  | GoCall GoExpr (Array GoExpr)
  | GoSelector GoExpr String
  | GoFunc String GoType GoType GoExpr
  | GoBlock (Array GoExpr)
  | GoReturn GoExpr
  | GoAssign String GoExpr
  | GoRecordDict (Array (Tuple String GoExpr))
  | GoRecordUpdateDict GoExpr (Array (Tuple String GoExpr))
  | GoRecordUpdateStatic GoExpr Int (Array (Tuple Int GoExpr)) (Array (Tuple String GoExpr))
  | GoIIFE String GoExpr GoExpr
  | GoLetRec (Array (Tuple String GoExpr)) GoExpr
  | GoRecordAccess GoExpr String
  | GoStructAccess GoExpr String
  | GoRecordAccessStatic GoExpr Int Int
  | GoConstructor String String (Array GoType) (Array GoExpr)
  | GoConstructorDict String (Array GoExpr)
  | GoConstructorAccess GoExpr String (Array GoType) Int Boolean
  | GoBranch (Array (Tuple GoExpr GoExpr)) GoExpr
  | GoBinOp String GoExpr GoExpr
  | GoPrefixOp String GoExpr
  | GoTypeAssertion GoExpr String
  | GoIndex GoExpr GoExpr
  | GoRaw String
  | GoFor String (Array GoExpr)
  | GoForRange String (Array GoExpr)
  | GoContinue String
  | GoMutate String GoExpr
  | GoIfElse GoExpr (Array GoExpr) (Array GoExpr)
  | GoFuncLit (Array (Tuple String GoType)) (Array GoExpr) GoExpr GoType

derive instance eqGoExpr :: Eq GoExpr

type GoDecl =
  { identifier :: String
  , expression :: GoExpr
  , goType :: GoType
  }

type GoFile =
  { packageName :: String
  , imports :: Array String
  , decls :: Array GoDecl
  , rawDecls :: Array String
  , foreigns :: Array { pursName :: String, goName :: String, exprType :: Maybe ExprType }
  }

data GoType
  = TypeValue
  | TypeInt64
  | TypeFloat64
  | TypeString
  | TypeBool
  | TypeUint32
  | TypeStructPointer String String String (Array GoType)
  | TypeRecord (Array (Tuple String GoType))
  | TypeInterface String
  | TypeNativeArray GoType
  | TypeGenericParam String
  | TypeFunc (Array GoType) GoType

derive instance eqGoType :: Eq GoType
derive instance ordGoType :: Ord GoType

goTypeToStr :: GoType -> String
goTypeToStr TypeInt64 = "int64"
goTypeToStr TypeFloat64 = "float64"
goTypeToStr TypeString = "string"
goTypeToStr TypeBool = "bool"
goTypeToStr TypeUint32 = "uint32"
goTypeToStr (TypeStructPointer _ _ monoStructName args) = "*" <> monoStructName
goTypeToStr (TypeInterface name) = name
goTypeToStr (TypeNativeArray inner) = "[]" <> goTypeToStr inner
goTypeToStr (TypeGenericParam name) = "gopurs_runtime.Value"
goTypeToStr (TypeFunc args ret) = "func(" <> String.joinWith ", " (map goTypeToStr args) <> ") " <> goTypeToStr ret
goTypeToStr _ = "gopurs_runtime.Value"

erasedGoTypeToStr :: GoType -> String
erasedGoTypeToStr TypeInt64 = "int64"
erasedGoTypeToStr TypeFloat64 = "float64"
erasedGoTypeToStr TypeString = "string"
erasedGoTypeToStr TypeBool = "bool"
erasedGoTypeToStr TypeUint32 = "uint32"
erasedGoTypeToStr (TypeStructPointer _ _ monoStructName args) = "*" <> monoStructName
erasedGoTypeToStr (TypeInterface name) = name
erasedGoTypeToStr (TypeNativeArray inner) = "[]" <> erasedGoTypeToStr inner
erasedGoTypeToStr (TypeGenericParam _) = "gopurs_runtime.Value"
erasedGoTypeToStr (TypeFunc args ret) = "func(" <> String.joinWith ", " (map erasedGoTypeToStr args) <> ") " <> erasedGoTypeToStr ret
erasedGoTypeToStr _ = "gopurs_runtime.Value"

extractTypeVars :: GoType -> Set.Set String
extractTypeVars (TypeStructPointer _ _ _ args) = fold (map extractTypeVars args)
extractTypeVars (TypeRecord fields) = fold (map (\(Tuple _ v) -> extractTypeVars v) fields)
extractTypeVars (TypeNativeArray inner) = extractTypeVars inner
extractTypeVars (TypeGenericParam name) = Set.singleton name
extractTypeVars (TypeFunc args ret) = fold (map extractTypeVars args) <> extractTypeVars ret
extractTypeVars _ = Set.empty

sanitizeName :: String -> String
sanitizeName name =
  let
    s1 = String.replaceAll (Pattern "/") (Replacement "_slash_")
      $ String.replaceAll (Pattern "\\") (Replacement "_bslash_")
      $ String.replaceAll (Pattern "<") (Replacement "_less_")
      $ String.replaceAll (Pattern ">") (Replacement "_greater_")
      $ String.replaceAll (Pattern "=") (Replacement "_eq_")
      $ String.replaceAll (Pattern "+") (Replacement "_plus_")
      $ String.replaceAll (Pattern "-") (Replacement "_minus_")
      $ String.replaceAll (Pattern "*") (Replacement "_times_")
      $ String.replaceAll (Pattern ":") (Replacement "_colon_")
      $ String.replaceAll (Pattern "|") (Replacement "_bar_")
      $ String.replaceAll (Pattern "&") (Replacement "_amp_")
      $ String.replaceAll (Pattern "^") (Replacement "_caret_")
      $ String.replaceAll (Pattern "~") (Replacement "_tilde_")
      $ String.replaceAll (Pattern "?") (Replacement "_qmark_")
      $ String.replaceAll (Pattern "!") (Replacement "_bang_")
      $ String.replaceAll (Pattern "@") (Replacement "_at_")
      $ String.replaceAll (Pattern "#") (Replacement "_hash_")
      $ String.replaceAll (Pattern "%") (Replacement "_percent_")
      $ String.replaceAll (Pattern "\"") (Replacement "_quote_")
      $ String.replaceAll (Pattern ".") (Replacement "_dot_")
      $ String.replaceAll (Pattern "'") (Replacement "_prime_")
      $ String.replaceAll (Pattern "$") (Replacement "_dollar_") name
  in
    if s1 == "break" || s1 == "default" || s1 == "func" || s1 == "interface" || s1 == "select" || s1 == "case" || s1 == "defer" || s1 == "go" || s1 == "map" || s1 == "struct" || s1 == "chan" || s1 == "else" || s1 == "goto" || s1 == "package" || s1 == "switch" || s1 == "const" || s1 == "fallthrough" || s1 == "if" || s1 == "range" || s1 == "type" || s1 == "continue" || s1 == "for" || s1 == "import" || s1 == "return" || s1 == "var" || s1 == "init" || s1 == "append" || s1 == "make" || s1 == "len" || s1 == "cap" || s1 == "new" || s1 == "close" || s1 == "delete" || s1 == "complex" || s1 == "real" || s1 == "imag" || s1 == "panic" || s1 == "recover" || s1 == "print" || s1 == "println" then "go__" <> s1 else s1

