module Gopurs.Int32Loops (loop) where

import Prelude
import Data.Array as Array
import Data.Foldable (foldl, sum)
import Data.Maybe (Maybe(..))
import Data.String (Pattern(..), contains)
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Gopurs.GoAst (GoExpr(..), GoType(..))
import Gopurs.GoFunctions (Parameters, iterationBindings)

-- Int's public representation remains int64: zshr and bottom / -1 can produce
-- values outside the signed 32-bit range. Version only scalar tail loops whose
-- selected slots are written by proven narrowing operations. Check their initial
-- values once, then keep both slots and iteration aliases native inside the loop.
-- The original loop handles wide inputs, including an immediate wide return.
loop :: String -> Parameters -> GoType -> Array GoExpr -> GoExpr
loop label params result statements =
  let
    original = GoFor label (iterationBindings params <> statements)
    candidates = Array.mapMaybe (\(Tuple name ty) ->
      if ty /= TypeInt64 then Nothing
      else case traverse (inspect label name) statements of
        Just counts | sum counts > 0 -> Just name
        _ -> Nothing) params
  in
    case Array.uncons candidates of
      Just { head, tail } | result == TypeInt64 ->
        let
          fits name = GoBinOp "==" (GoVar (slot name)) (cast "int64" (cast "int32" (GoVar (slot name))))
          guard = foldl (GoBinOp "&&") (fits head) (map fits tail)
          initial = map (\name -> GoAssign (slot name) (cast "int32" (GoVar (slot name)))) candidates
          aliases = Array.concatMap (\param@(Tuple name _) ->
            if Array.elem name candidates then [ GoAssign name (GoVar (slot name)) ]
            else iterationBindings [ param ]) params
          fast = GoCall (GoFuncBlock []
            (initial <> [ GoFor label (aliases <> map (rewrite candidates) statements) ]) result) []
        in
          -- Go labels have function scope. An IIFE gives the copied branch/loop
          -- labels a separate scope without renaming opaque goto syntax.
          GoBlock [ GoIfElse guard [ GoReturn fast ] [], original ]
      _ -> original

slot :: String -> String
slot name = name <> "_loop"

cast :: String -> GoExpr -> GoExpr
cast name value = GoCall (GoVar name) [ value ]

arithmetic :: GoExpr -> Maybe { operator :: String, left :: GoExpr, right :: GoExpr }
arithmetic = case _ of
  GoCall (GoSelector (GoVar "gopurs_runtime") name) [ left, right ]
    | name == "IntAdd" -> Just { operator: "+", left, right }
    | name == "IntSub" -> Just { operator: "-", left, right }
  _ -> Nothing

-- Deliberately exclude division, unsigned shifts and arbitrary Int-returning
-- calls. Types alone do not prove a signed 32-bit result.
normalized :: GoExpr -> Boolean
normalized expr = case expr of
  GoInt value -> signedLiteral value
  GoCall (GoVar "int64") [ GoInt value ] -> signedLiteral value
  _ -> case arithmetic expr of
    Just _ -> true
    Nothing -> false

-- PBO can fold zshr or bottom / -1 to an Int literal outside this range.
signedLiteral :: Int -> Boolean
signedLiteral value = value >= bottom && value <= top

-- Fail closed on opaque references, shadowing, closures, nested loops and
-- unrecognized AST forms. Ordinary calls retain the iteration's int64 values;
-- only the private, directly assigned tail slots change representation.
inspect :: String -> String -> GoExpr -> Maybe Int
inspect label name = go
  where
  reserved value = value == name || value == slot name
  many values = map sum (traverse go values)
  go = case _ of
    GoVar value | value == slot name -> Nothing
    GoVar _ -> Just 0
    GoString _ -> Just 0
    GoInt _ -> Just 0
    GoRaw code | contains (Pattern name) code.text -> Nothing
    GoRaw _ -> Just 0
    GoAssign target _ | reserved target -> Nothing
    GoAssign _ value -> go value
    GoMutate target value
      | target == name -> Nothing
      | target == slot name -> if normalized value then map (_ + 1) (go value) else Nothing
      | otherwise -> go value
    GoContinue target -> if target == label then Just 0 else Nothing
    GoCall fn args -> many (Array.cons fn args)
    GoSelector obj _ -> go obj
    GoBinOp _ left right -> many [ left, right ]
    GoPrefixOp _ value -> go value
    GoTypeAssertion value _ -> go value
    GoIndex obj index -> many [ obj, index ]
    GoReturn value -> go value
    GoBlock values -> many values
    GoIfElse cond yes no -> many (Array.cons cond (yes <> no))
    _ -> Nothing

children :: GoExpr -> Array GoExpr
children = case _ of
  GoCall fn args -> Array.cons fn args
  GoSelector obj _ -> [ obj ]
  GoBinOp _ left right -> [ left, right ]
  GoPrefixOp _ value -> [ value ]
  GoTypeAssertion value _ -> [ value ]
  GoIndex obj index -> [ obj, index ]
  _ -> []

readsNative :: Array String -> GoExpr -> Boolean
readsNative names = go
  where
  go (GoVar name) = Array.elem name names
  go expr = Array.any go (children expr)

known32 :: Array String -> GoExpr -> Boolean
known32 names (GoVar name) = Array.elem name names
known32 _ expr = normalized expr

comparison :: String -> Boolean
comparison op = Array.elem op [ "==", "!=", "<", "<=", ">", ">=" ]

rewrite :: Array String -> GoExpr -> GoExpr
rewrite names = wide
  where
  narrow expr = case expr of
    GoVar name | Array.elem name names -> expr
    GoInt value | signedLiteral value -> cast "int32" expr
    GoCall (GoVar "int64") [ GoInt value ] -> narrow (GoInt value)
    _ -> case arithmetic expr of
      -- A native operand also prevents Go constant expressions overflowing at
      -- type checking time. Constant-only additions retain the runtime helper.
      Just { operator, left, right } | readsNative names expr ->
        GoBinOp operator (narrow left) (narrow right)
      -- The operand may itself be an unsigned constant expression. A helper
      -- permits its reduction without a Go compile-time conversion overflow.
      _ -> cast "int32" (GoCall (GoSelector (GoVar "gopurs_runtime") "IntAdd")
        [ wide expr, cast "int64" (GoInt 0) ])

  wide expr = case expr of
    GoVar name | Array.elem name names -> cast "int64" expr
    GoCall fn args -> case arithmetic expr of
      Just _ | readsNative names expr -> cast "int64" (narrow expr)
      _ -> GoCall (wide fn) (map wide args)
    GoBinOp op left right
      | comparison op && known32 names left && known32 names right -> GoBinOp op (narrow left) (narrow right)
      | otherwise -> GoBinOp op (wide left) (wide right)
    GoSelector obj field -> GoSelector (wide obj) field
    GoPrefixOp op value -> GoPrefixOp op (wide value)
    GoTypeAssertion value ty -> GoTypeAssertion (wide value) ty
    GoIndex obj index -> GoIndex (wide obj) (wide index)
    GoAssign target value -> GoAssign target (wide value)
    GoMutate target value
      | Array.any (\name -> target == slot name) names -> GoMutate target (narrow value)
      | otherwise -> GoMutate target (wide value)
    GoReturn value -> GoReturn (wide value)
    GoBlock values -> GoBlock (map wide values)
    GoIfElse cond yes no -> GoIfElse (wide cond) (map wide yes) (map wide no)
    _ -> expr
