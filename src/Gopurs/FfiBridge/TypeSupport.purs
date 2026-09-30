module Gopurs.FfiBridge.TypeSupport
  ( printTypeNode
  , instantiateDeclaration
  , getTastArgType
  , getTastReturnType
  , resolveNewtype
  , flattenFuncArgs
  , isStandardPursFunc
  , isRuntimeValueNode
  ) where

import Prelude

import Data.Array as Array
import Data.Maybe (Maybe(..), fromMaybe)
import Data.String as String
import Data.Tuple (Tuple(..))
import Gopurs.FfiTypes (FfiDecl, TypeNode(..))
import PureScript.Backend.Optimizer.CoreFn (DataDecl, ExprType(..))

printTypeNode :: TypeNode -> String
printTypeNode (TNamed n) = n
printTypeNode (TFunc args ret) =
  let
    argsStr = String.joinWith ", " (map printTypeNode args)
    retStr = case ret of
      Nothing -> ""
      Just r -> " " <> printTypeNode r
  in
    "func(" <> argsStr <> ")" <> retStr
printTypeNode (TArray elem) = "[]" <> printTypeNode elem
printTypeNode (TMap k v) = "map[" <> printTypeNode k <> "]" <> printTypeNode v
printTypeNode (TUnknown s) = s

-- Wrapper calls instantiate Go type parameters with Value. The same
-- substitution must reach callbacks and containers before adapting values.
instantiateDeclaration :: FfiDecl -> FfiDecl
instantiateDeclaration declaration = declaration
  { args = map instantiate declaration.args
  , ret = map instantiate declaration.ret
  }
  where
  instantiate = case _ of
    TNamed name | Array.elem name declaration.typeParams -> TNamed "gopurs_runtime.Value"
    TFunc args ret -> TFunc (map instantiate args) (map instantiate ret)
    TArray item -> TArray (instantiate item)
    TMap key value -> TMap (instantiate key) (instantiate value)
    other -> other

getTastArgType :: ExprType -> Int -> Maybe ExprType
getTastArgType (Func args _) i = Array.index args i
getTastArgType _ _ = Nothing

getTastReturnType :: ExprType -> Maybe ExprType
getTastReturnType (Func _ ret) = Just ret
getTastReturnType _ = Nothing

resolveNewtype :: Array DataDecl -> ExprType -> ExprType
resolveNewtype dataDecls (ADT fullName path args) =
  let
    mbDecl = Array.find (\d -> d.name == fullName || (Array.length path > 0 && d.name == fromMaybe "" (Array.last path))) dataDecls
  in
    case mbDecl of
      Just decl ->
        if Array.length decl.constructors == 1 then
          case Array.head decl.constructors of
            Just ctor ->
              if Array.length ctor.fields == 1 then
                case Array.head ctor.fields of
                  Just fieldT -> resolveNewtype dataDecls fieldT
                  Nothing -> ADT fullName path args
              else ADT fullName path args
            Nothing -> ADT fullName path args
        else ADT fullName path args
      Nothing -> ADT fullName path args
resolveNewtype dataDecls (Func fArgs ret) = Func (map (resolveNewtype dataDecls) fArgs) (resolveNewtype dataDecls ret)
resolveNewtype dataDecls (Array elem) = Array (resolveNewtype dataDecls elem)
resolveNewtype dataDecls (Record row) = Record (resolveNewtype dataDecls row)
resolveNewtype dataDecls (Row fields tail) = Row (map (\(Tuple k v) -> Tuple k (resolveNewtype dataDecls v)) fields) (map (resolveNewtype dataDecls) tail)
resolveNewtype dataDecls (TypeApp c args) = TypeApp (resolveNewtype dataDecls c) (map (resolveNewtype dataDecls) args)
resolveNewtype dataDecls (ForAll vars body) = ForAll vars (resolveNewtype dataDecls body)
resolveNewtype dataDecls (ConstrainedType constraints body) = ConstrainedType (map (\(Tuple c a) -> Tuple c (map (resolveNewtype dataDecls) a)) constraints) (resolveNewtype dataDecls body)
resolveNewtype _ other = other

flattenFuncArgs :: ExprType -> Array ExprType
flattenFuncArgs (Func args ret) = args <> flattenFuncArgs ret
flattenFuncArgs _ = []

isStandardPursFunc :: TypeNode -> Boolean
isStandardPursFunc (TFunc args ret) =
  let
    argsAny = Array.all (\a -> printTypeNode a == "any" || printTypeNode a == "interface{}" || printTypeNode a == "gopurs_runtime.Value") args
    retStr = case ret of
      Nothing -> ""
      Just r -> printTypeNode r
  in
    if Array.length args == 0 then
      retStr == "" || retStr == "bool" || retStr == "int" || retStr == "int64" || retStr == "string" || retStr == "float64" || retStr == "gopurs_runtime.Value" || retStr == "any" || retStr == "interface{}"
    else if argsAny then
      case ret of
        Nothing -> true
        Just r@(TFunc _ _) -> isStandardPursFunc r
        Just _ -> retStr == "any" || retStr == "interface{}" || retStr == "gopurs_runtime.Value"
    else
      false
isStandardPursFunc _ = false

isRuntimeValueNode :: TypeNode -> Boolean
isRuntimeValueNode = case _ of
  TNamed "gopurs_runtime.Value" -> true
  _ -> false
