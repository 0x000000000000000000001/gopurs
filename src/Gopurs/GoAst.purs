module Gopurs.GoAst where

import Prelude
import Data.Generic.Rep (class Generic)
import Data.Int as Int
import Data.Tuple (Tuple(..))
import Data.String as String
import Data.String (Pattern(..), Replacement(..))
import PureScript.Backend.Optimizer.CoreFn (ModuleName)
import Gopurs.GoCode (GoCode, opaqueCode)
import Gopurs.GoSymbol (GoReference, GoSymbol, generated)
import Data.Maybe (Maybe(..))
import Data.Newtype (unwrap)
import Data.Array as Array
import Data.Enum (fromEnum)
import Data.String.CodePoints as CodePoints

data GoExpr
  = GoVar String
  | GoRef GoReference
  | GoString String
  | GoInt Int
  | GoCall GoExpr (Array GoExpr)
  | GoTypeApply GoExpr (Array GoType)
  | GoSelector GoExpr String
  | GoBlock (Array GoExpr)
  | GoReturn GoExpr
  | GoAssign String GoExpr
  | GoRecordDict GoType (Array (Tuple String GoExpr))
  | GoRecordUpdateDict GoExpr (Array (Tuple String GoExpr))
  | GoRecordUpdateStatic GoExpr Int (Array (Tuple Int GoExpr)) (Array (Tuple String GoExpr))
  | GoRecordUpdateNative GoType GoExpr (Array (Tuple String GoExpr))
  | GoIIFE String GoExpr GoExpr
  | GoLetRec (Array (Tuple String GoExpr)) GoExpr
  | GoRecordAccess GoExpr String
  | GoStructAccess GoExpr String
  | GoRecordFieldAccess GoExpr String
  | GoRecordAccessStatic GoExpr Int Int
  | GoConstructor String GoSymbol (Array GoType) (Array GoExpr)
  | GoConstructorDict String (Array GoExpr)
  | GoConstructorAccess GoExpr GoSymbol (Array GoType) Int Boolean
  | GoBranch (Array (Tuple GoExpr GoExpr)) GoExpr
  | GoBinOp String GoExpr GoExpr
  | GoPrefixOp String GoExpr
  | GoTypeAssertion GoExpr GoType
  | GoIndex GoExpr GoExpr
  -- Retain the operand and its imports until the final rendering.
  | GoBoxStructPointer String GoExpr
  -- Keep these conversions structured until the array consumer is known.
  | GoBoxIntArray GoExpr
  | GoUnboxIntArray GoExpr
  -- Only the intrinsic filter's private []Value buffer may carry this marker.
  | GoFreshFilterArray GoExpr
  | GoRaw GoCode
  | GoFor String (Array GoExpr)
  | GoForRange String (Array GoExpr)
  | GoContinue String
  | GoMutate String GoExpr
  | GoIfElse GoExpr (Array GoExpr) (Array GoExpr)
  -- A native function body with explicit returns, including loop bodies.
  -- Parameter grouping and runtime wrappers are chosen before printing.
  | GoFuncBlock (Array (Tuple String GoType)) (Array GoExpr) GoType
  | GoFuncLit (Array (Tuple String GoType)) (Array GoExpr) GoExpr GoType
  | GoStructValue String (Array GoType) (Array GoExpr)

derive instance eqGoExpr :: Eq GoExpr

type GoBinding =
  { identifier :: GoSymbol
  , expression :: GoExpr
  , goType :: GoType
  }

data GoDecl
  = GoCachedValue GoBinding
  | GoStructDecl
      { name :: GoSymbol
      , typeParams :: Array (Tuple String GoType)
      , fields :: Array (Tuple String GoType)
      }
  | GoFunctionDecl
      { name :: GoSymbol
      , params :: Array (Tuple String GoType)
      , result :: GoType
      , body :: GoExpr
      }
  | GoInitDecl GoSymbol GoExpr
  | GoForeignGetter { name :: GoSymbol, value :: GoSymbol }

type GoFile =
  { packageName :: String
  , imports :: Array String
  , declarationGroups :: Array (Array GoDecl)
  }

-- Opaque syntax carries its dependencies alongside its text. Structured
-- expressions and declarations expose their dependencies to GoImports.
rawGo :: String -> GoExpr
rawGo = GoRaw <<< opaqueCode

-- Keep the runtime tag, PureScript identity and Go names distinct.
-- structPointer is the single constructor of the instantiated name.
type StructPointer =
  { baseStructName :: String -- Runtime tag identity.
  , fullName :: String -- Fully qualified PureScript type.
  , fullPath :: String -- Instantiated Go constructor name.
  , structName :: String -- Go constructor name without type arguments.
  , symbol :: GoSymbol -- Owner of that single named Go declaration.
  , typeArgs :: Array GoType
  }

-- The caller supplies its module prefix. In particular, metadata for elided
-- constructors retains module dots; normal Go declarations supply underscores.
constructorNames :: String -> String -> { baseStructName :: String, structName :: String, symbol :: GoSymbol }
constructorNames modulePrefix constructorName =
  let suffix = modulePrefix <> "_" <> sanitizeName constructorName
      structName = "Constructor_" <> suffix
      unit = String.replaceAll (Pattern ".") (Replacement "_") modulePrefix
  in { baseStructName: "Data_" <> suffix, structName, symbol: generated unit structName }

structPointer :: forall r. { baseStructName :: String, fullName :: String, structName :: String, symbol :: GoSymbol | r } -> Array GoType -> GoType
structPointer { baseStructName, fullName, structName, symbol } typeArgs =
  TypeStructPointer { baseStructName, fullName, structName, symbol, fullPath: structTypeName structName typeArgs, typeArgs }

structTypeName :: String -> Array GoType -> String
structTypeName name args = name <> if Array.null args then "" else "[" <> String.joinWith ", " (map goTypeToStr args) <> "]"

data GoType
  = TypeValue
  | TypeInt64
  | TypeFloat64
  | TypeString
  | TypeBool
  | TypeUint32
  | TypeStructPointer StructPointer
  | TypeRecord (Array (Tuple String GoType))
  | TypeInterface String
  | TypeNamed GoSymbol
  | TypePointer GoType
  | TypeNativeArray GoType
  | TypeGenericParam String
  | TypeFunc (Array GoType) GoType
  | TypeStructValue String (Array GoType)

derive instance eqGoType :: Eq GoType
derive instance ordGoType :: Ord GoType
derive instance Generic GoType _

goTypeToStr :: GoType -> String
goTypeToStr TypeInt64 = "int64"
goTypeToStr TypeFloat64 = "float64"
goTypeToStr TypeString = "string"
goTypeToStr TypeBool = "bool"
goTypeToStr TypeUint32 = "uint32"
goTypeToStr (TypeStructPointer { fullPath }) = "*" <> fullPath
goTypeToStr (TypeInterface name) = name
goTypeToStr (TypeNamed symbol) = symbol.name
goTypeToStr (TypePointer ty) = "*" <> goTypeToStr ty
goTypeToStr (TypeNativeArray inner) = "[]" <> goTypeToStr inner
goTypeToStr (TypeGenericParam name) = "T_" <> sanitizeName name
goTypeToStr (TypeFunc args ret) = "func(" <> String.joinWith ", " (map goTypeToStr args) <> ") " <> goTypeToStr ret
goTypeToStr (TypeStructValue _ fields) = "struct{" <> String.joinWith "; " (Array.mapWithIndex (\i t -> "V" <> show i <> " " <> goTypeToStr t) fields) <> "}"
goTypeToStr (TypeRecord fields) = goRecordStructName fields
goTypeToStr _ = "gopurs_runtime.Value"

-- Helper identities describe Go types, independently of transport annotations.
-- Recurse through every composite so nested records cannot rename a helper.
structTypeIdentity :: String -> Array GoType -> String
structTypeIdentity name args = name <> if Array.null args then "" else "[" <> String.joinWith ", " (map goTypeIdentity args) <> "]"

goTypeIdentity :: GoType -> String
goTypeIdentity (TypeStructPointer { structName, typeArgs }) = "*" <> structTypeIdentity structName typeArgs
goTypeIdentity (TypePointer ty) = "*" <> goTypeIdentity ty
goTypeIdentity (TypeNativeArray inner) = "[]" <> goTypeIdentity inner
goTypeIdentity (TypeFunc args ret) = "func(" <> String.joinWith ", " (map goTypeIdentity args) <> ") " <> goTypeIdentity ret
goTypeIdentity (TypeStructValue _ fields) = "struct{" <> String.joinWith "; " (Array.mapWithIndex (\i t -> "V" <> show i <> " " <> goTypeIdentity t) fields) <> "}"
goTypeIdentity (TypeRecord fields) =
  "struct{\n" <> String.joinWith "\n" (map (\(Tuple k v) -> "\t" <> recordFieldName k <> " " <> goTypeIdentity v) fields) <> "\n}"
goTypeIdentity ty = goTypeToStr ty

goRecordStructName :: Array (Tuple String GoType) -> String
goRecordStructName fields =
  "struct{\n" <> String.joinWith "\n" (map (\(Tuple k v) -> "\t" <> recordFieldName k <> " " <> goTypeToStr v) fields)
    <> "\n/*gopurs:record-v1:" <> show (Array.length fields) <> ":"
    <> String.joinWith "," (map (\(Tuple label _) -> recordLabelUnits label) fields) <> "*/\n}"

-- Opaque expressions also print Go types. Carry the original labels through
-- that textual boundary; the typed package emitter consumes this annotation.
-- Hex UTF-16 preserves isolated surrogates and cannot terminate a Go comment.
recordLabelUnits :: String -> String
recordLabelUnits = String.joinWith "" <<< map (encodePoint <<< fromEnum) <<< CodePoints.toCodePointArray
  where
  encodePoint point
    | point > 65535 = encode (55296 + div (point - 65536) 1024) <> encode (56320 + mod (point - 65536) 1024)
    | otherwise = encode point
  encode value =
    let hex = Int.toStringAs Int.hexadecimal value
    in String.take (4 - String.length hex) "0000" <> hex

-- Labels are arbitrary PSStrings, unlike identifiers. Keep the established
-- ASCII spelling, encoding other labels into a portable Go field name. Reserve
-- the encoding prefix so a literal label cannot collide with an encoded one.
recordFieldName :: String -> String
recordFieldName = memoizeName \label ->
  let
    name = sanitizeName label
    prefix = "gopurs_field_"
    asciiIdentifier code =
      let n = fromEnum code
      in n == 95 || (n >= 48 && n <= 57) || (n >= 65 && n <= 90) || (n >= 97 && n <= 122)
  in
    if label /= "" && name /= "_" && String.take (String.length prefix) label /= prefix
        && String.take 4 label /= "go__"
        && not (Array.elem (String.take 1 label) [ "0", "1", "2", "3", "4", "5", "6", "7", "8", "9" ])
        && Array.all asciiIdentifier (CodePoints.toCodePointArray label) then name
    else prefix <> String.joinWith "_" (map (show <<< fromEnum) (CodePoints.toCodePointArray label))



capitalize :: String -> String
capitalize "" = ""
capitalize s =
  let
    firstChar = String.take 1 s
  in
    if firstChar >= "a" && firstChar <= "z" then String.toUpper firstChar <> String.drop 1 s
    else if firstChar == "_" then "_" <> capitalize (String.drop 1 s)
    else s

sanitizeName :: String -> String
sanitizeName = memoizeName sanitizeNameUncached

foreign import memoizeName :: (String -> String) -> String -> String

sanitizeNameUncached :: String -> String
sanitizeNameUncached name =
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
    
    s2 = if s1 == "break" || s1 == "default" || s1 == "func" || s1 == "interface" || s1 == "select" || s1 == "case" || s1 == "defer" || s1 == "go" || s1 == "map" || s1 == "struct" || s1 == "chan" || s1 == "else" || s1 == "goto" || s1 == "package" || s1 == "switch" || s1 == "const" || s1 == "fallthrough" || s1 == "if" || s1 == "range" || s1 == "type" || s1 == "continue" || s1 == "for" || s1 == "import" || s1 == "return" || s1 == "var" || s1 == "init" || s1 == "append" || s1 == "make" || s1 == "len" || s1 == "cap" || s1 == "new" || s1 == "close" || s1 == "delete" || s1 == "complex" || s1 == "real" || s1 == "imag" || s1 == "panic" || s1 == "recover" || s1 == "print" || s1 == "println" then "go__" <> s1 else s1
    
    firstChar = String.take 1 s2
  in
    if firstChar == "0" || firstChar == "1" || firstChar == "2" || firstChar == "3" || firstChar == "4" || firstChar == "5" || firstChar == "6" || firstChar == "7" || firstChar == "8" || firstChar == "9" then "X_" <> s2
    else if s2 == "" then "X_empty"
    else s2

getStructName :: String -> Maybe ModuleName -> String -> String
getStructName modNameStr mbMod ctorName =
  let
    modNamePart = case mbMod of
      Just mn -> sanitizeName (String.replaceAll (Pattern ".") (Replacement "_") (unwrap mn))
      Nothing -> modNameStr
  in
    (constructorNames modNamePart ctorName).baseStructName
