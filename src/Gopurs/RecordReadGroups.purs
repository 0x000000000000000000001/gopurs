module Gopurs.RecordReadGroups (optimize, declaration) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Maybe (Maybe(..))
import Data.Tuple (Tuple(..))
import Gopurs.GoAst (GoDecl(..), GoExpr(..), rawGo)
import Gopurs.Printer (printGoExpr)
import Gopurs.RecordShapeGuard as RecordShapeGuard

type Reads = { source :: Maybe String, keys :: Array String, variables :: Array String }

-- This pass runs after TAST-driven lowering. Only native Int operations and
-- their scalar leaves are admitted; an IntVal projection is the already chosen
-- Value -> Int representation, not an inference that arbitrary calls are pure.
-- Match the whole arithmetic tree before visiting children, so its fallback
-- remains exactly the original expression and is never optimized recursively.
optimize :: GoExpr -> GoExpr
optimize expression = case grouped expression of
  Just result -> result
  Nothing -> children optimize expression

declaration :: GoDecl -> GoDecl
declaration = case _ of
  GoCachedValue binding -> GoCachedValue (binding { expression = optimize binding.expression })
  GoFunctionDecl function -> GoFunctionDecl (function { body = optimize function.body })
  GoInitDecl name expression -> GoInitDecl name (optimize expression)
  other -> other

grouped :: GoExpr -> Maybe GoExpr
grouped original = do
  guard (case original of
    GoCall (GoVar "int64") [ _ ] -> true
    GoCall (GoSelector (GoVar "gopurs_runtime") operation) [ _, _ ] ->
      Array.elem operation [ "IntAdd", "IntSub", "IntMul", "IntDiv", "IntMod" ]
    _ -> false)
  reads <- scalar original
  source <- reads.source
  let keys = Array.sort (Array.nub reads.keys)
  guard (Array.length reads.keys >= 2)
  let
    prefix = fresh "recordReadGroup" reads.variables
    input = prefix <> "_input"
  pure (rawGo
    ("func(" <> input <> " gopurs_runtime.Value) int64 {\n"
      <> RecordShapeGuard.integerReads prefix input keys (\values ->
        "return " <> printGoExpr (additions (replace source keys values original)))
      <> "return " <> printGoExpr original <> "\n}(" <> source <> ")"))

fresh :: String -> Array String -> String
fresh prefix variables = if Array.any (\suffix -> Array.elem (prefix <> suffix) variables)
  [ "_input", "_data", "_keys", "_value0", "_value1" ]
  then fresh (prefix <> "_") variables else prefix

empty :: Reads
empty = { source: Nothing, keys: [], variables: [] }

scalar :: GoExpr -> Maybe Reads
scalar = case _ of
  GoInt _ -> Just empty
  GoVar name -> Just (empty { variables = [ name ] })
  GoSelector (GoVar name) "IntVal" -> Just (empty { variables = [ name ] })
  GoSelector (GoRecordAccess (GoVar source) key) "IntVal" ->
    Just { source: Just source, keys: [ key ], variables: [ source ] }
  field@(GoRecordFieldAccess _ _) -> nativeField field
  GoCall (GoVar cast) [ value ] | cast == "int64" || cast == "int32" -> scalar value
  GoCall (GoSelector (GoVar "gopurs_runtime") operation) [ left, right ]
    | Array.elem operation [ "IntAdd", "IntSub", "IntMul", "IntDiv", "IntMod" ] -> do
        leftReads <- scalar left
        rightReads <- scalar right
        combine leftReads rightReads
  _ -> Nothing

nativeField :: GoExpr -> Maybe Reads
nativeField = case _ of
  GoVar name -> Just (empty { variables = [ name ] })
  GoRecordFieldAccess object _ -> nativeField object
  _ -> Nothing

combine :: Reads -> Reads -> Maybe Reads
combine left right = do
  source <- case left.source, right.source of
    Nothing, source -> Just source
    source, Nothing -> Just source
    Just a, Just b | a == b -> Just (Just a)
    _, _ -> Nothing
  pure { source, keys: left.keys <> right.keys, variables: left.variables <> right.variables }

replace :: String -> Array String -> Array GoExpr -> GoExpr -> GoExpr
replace source keys values expression = case expression of
  GoSelector (GoRecordAccess (GoVar object) key) "IntVal" | object == source ->
    case Array.elemIndex key keys of
      Just index -> case Array.index values index of
        Just value -> value
        Nothing -> expression
      Nothing -> expression
  _ -> children (replace source keys values) expression

-- IntAdd is addition modulo 2^32, including noncanonical Int64 leaves.
-- Inside this validated read-only region, connected additions need only their
-- outer narrowing. Other arithmetic keeps its original narrowing boundary;
-- in particular IntMul's JS Number rounding is never treated as ring multiply.
additions :: GoExpr -> GoExpr
additions expression = case expression of
  GoCall (GoSelector (GoVar "gopurs_runtime") "IntAdd") [ _, _ ] ->
    if constantAddition expression then expression
    else GoCall (GoVar "int64") [ GoCall (GoVar "int32") [ chain expression ] ]
  _ -> children additions expression
  where
  chain value = case value of
    GoCall (GoSelector (GoVar "gopurs_runtime") "IntAdd") [ left, right ] ->
      if constantAddition value then value else GoBinOp "+" (chain left) (chain right)
    other -> additions other

-- Go rejects an overflowing constant conversion instead of applying runtime
-- Int32 wrapping. Keep constant-only islands as runtime calls, including when
-- they are nested inside a larger addition with dynamic leaves.
constantAddition :: GoExpr -> Boolean
constantAddition = case _ of
  GoInt _ -> true
  GoCall (GoVar cast) [ value ] | cast == "int64" || cast == "int32" -> constantAddition value
  GoCall (GoSelector (GoVar "gopurs_runtime") "IntAdd") [ left, right ] ->
    constantAddition left && constantAddition right
  _ -> false

-- Walk inside existing containers only. No expression is moved into or out of
-- a branch, binding, callback, loop or argument position. Opaque generated code
-- is intentionally left alone, including the complete original fallback.
children :: (GoExpr -> GoExpr) -> GoExpr -> GoExpr
children visit = case _ of
  GoCall function args -> GoCall (visit function) (map visit args)
  GoTypeApply function types -> GoTypeApply (visit function) types
  GoSelector object key -> GoSelector (visit object) key
  GoBlock values -> GoBlock (map visit values)
  GoReturn value -> GoReturn (visit value)
  GoAssign name value -> GoAssign name (visit value)
  GoRecordDict ty fields -> GoRecordDict ty (map field fields)
  GoRecordUpdateDict object fields -> GoRecordUpdateDict (visit object) (map field fields)
  GoRecordUpdateStatic object size fields fallback -> GoRecordUpdateStatic (visit object) size (map field fields) (map field fallback)
  GoRecordUpdateNative ty object fields -> GoRecordUpdateNative ty (visit object) (map field fields)
  GoIIFE name value body -> GoIIFE name (visit value) (visit body)
  GoLetRec bindings body -> GoLetRec (map field bindings) (visit body)
  GoRecordAccess object key -> GoRecordAccess (visit object) key
  GoStructAccess object key -> GoStructAccess (visit object) key
  GoRecordFieldAccess object key -> GoRecordFieldAccess (visit object) key
  GoRecordAccessStatic object size index -> GoRecordAccessStatic (visit object) size index
  GoConstructor tag name types args -> GoConstructor tag name types (map visit args)
  GoConstructorDict tag args -> GoConstructorDict tag (map visit args)
  GoConstructorAccess object name types index native -> GoConstructorAccess (visit object) name types index native
  GoBranch branches otherwise -> GoBranch (map (\(Tuple condition value) -> Tuple (visit condition) (visit value)) branches) (visit otherwise)
  GoBinOp operation left right -> GoBinOp operation (visit left) (visit right)
  GoPrefixOp operation value -> GoPrefixOp operation (visit value)
  GoTypeAssertion value ty -> GoTypeAssertion (visit value) ty
  GoIndex object index -> GoIndex (visit object) (visit index)
  GoBoxStructPointer name value -> GoBoxStructPointer name (visit value)
  GoBoxIntArray value -> GoBoxIntArray (visit value)
  GoUnboxIntArray value -> GoUnboxIntArray (visit value)
  GoFreshFilterArray value -> GoFreshFilterArray (visit value)
  GoFor condition body -> GoFor condition (map visit body)
  GoForRange condition body -> GoForRange condition (map visit body)
  GoMutate name value -> GoMutate name (visit value)
  GoIfElse condition yes no -> GoIfElse (visit condition) (map visit yes) (map visit no)
  GoFuncBlock params body ty -> GoFuncBlock params (map visit body) ty
  GoFuncLit params statements result ty -> GoFuncLit params (map visit statements) (visit result) ty
  GoStructValue name types fields -> GoStructValue name types (map visit fields)
  other -> other
  where
  field :: forall a. Tuple a GoExpr -> Tuple a GoExpr
  field (Tuple key value) = Tuple key (visit value)
