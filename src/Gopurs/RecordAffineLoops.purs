-- Collapse pure countdown recurrences over immutable native records. The
-- complete update tree is proved before changing the number of iterations.
module Gopurs.RecordAffineLoops (wrap) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Maybe (Maybe(..))
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..), fst)
import Gopurs.GoAst (GoExpr(..), GoType(..))
import Gopurs.GoFunctions (Parameters)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (Literal(..), Prop(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (BackendSyntax, Pair(..))
import PureScript.Backend.Optimizer.Syntax as Syn

data Delta = Constant Int | Remainders Int
data Change = Keep | Fields GoType (Array (Tuple String Change)) | Add Delta

type Recurrence = { counter :: String, record :: String, changes :: Change }

wrap :: (TcoExpr -> Boolean) -> Parameters -> GoType -> TcoExpr -> GoExpr -> GoExpr
wrap self params result body original = case recognize self params result body of
  Just recurrence ->
    let counter = GoVar (recurrence.counter <> "_loop")
        record = GoVar (recurrence.record <> "_loop")
        admitted = GoBinOp "&&" (GoBinOp ">" counter (integer 0))
          (GoBinOp "<=" counter (integer top))
    in GoBlock [ GoIfElse admitted [ GoReturn (emit counter record recurrence.changes) ] [], original ]
  Nothing -> original

recognize :: (TcoExpr -> Boolean) -> Parameters -> GoType -> TcoExpr -> Maybe Recurrence
recognize self params result body = do
  parameters <- case params of
    [ Tuple counter TypeInt64, Tuple record ty@(TypeRecord _) ] ->
      Just { counter, record, ty, counterFirst: true }
    [ Tuple record ty@(TypeRecord _), Tuple counter TypeInt64 ] ->
      Just { counter, record, ty, counterFirst: false }
    _ -> Nothing
  guard (parameters.counter /= parameters.record && result == parameters.ty)
  branch <- case syntax body of
    Syn.Branch branches step -> case NEA.toArray branches of
      [ Pair condition seed ] -> Just { condition, seed, step }
      _ -> Nothing
    _ -> Nothing
  guard (case syntax branch.condition of
    Syn.PrimOp (Syn.Op2 (Syn.OpIntOrd Syn.OpEq) left right) ->
      (local parameters.counter left && literal 0 right)
        || (literal 0 left && local parameters.counter right)
    _ -> false)
  guard (local parameters.record branch.seed)
  arguments <- case syntax branch.step of
    Syn.App target args | self target -> Just (NEA.toArray args)
    Syn.UncurriedApp target args | self target -> Just args
    _ -> Nothing
  Tuple nextCounter nextRecord <- case arguments of
    [ first, second ] -> Just (if parameters.counterFirst then Tuple first second else Tuple second first)
    _ -> Nothing
  guard (case syntax nextCounter of
    Syn.PrimOp (Syn.Op2 (Syn.OpIntNum Syn.OpSubtract) left right) ->
      local parameters.counter left && literal 1 right
    _ -> false)
  changes <- recordChanges parameters.counter parameters.record [] parameters.ty nextRecord
  pure { counter: parameters.counter, record: parameters.record, changes }

-- Every updated leaf reads its own previous value. Cross-field dependencies,
-- calls, effects, duplicate updates and layout changes retain the original loop.
-- Unchanged fields are copied by value, including any borrowed references.
recordChanges :: String -> String -> Array String -> GoType -> TcoExpr -> Maybe Change
recordChanges counter record path ty expression
  | recordPath record expression == Just path = Just Keep
  | otherwise = case ty, syntax expression of
      recordType@(TypeRecord fields), Syn.Update source props -> do
        guard (recordPath record source == Just path)
        let keys = map (\(Prop key _) -> key) props
        guard (Array.length (Array.nub keys) == Array.length keys)
        guard (Array.length (Array.nub (map fst fields)) == Array.length fields)
        updates <- traverse (\(Prop key value) -> do
          Tuple _ fieldType <- Array.find (\(Tuple label _) -> label == key) fields
          change <- recordChanges counter record (Array.snoc path key) fieldType value
          pure (Tuple key change)) props
        pure (Fields recordType updates)
      TypeInt64, Syn.PrimOp (Syn.Op2 (Syn.OpIntNum Syn.OpAdd) left right)
        | recordPath record left == Just path -> Add <$> increment counter right
        | recordPath record right == Just path -> Add <$> increment counter left
      _, _ -> Nothing

increment :: String -> TcoExpr -> Maybe Delta
increment counter expression = case syntax expression of
  Syn.Lit (LitInt value) | value >= bottom && value <= top -> Just (Constant value)
  Syn.PrimOp (Syn.Op2 (Syn.OpIntNum Syn.OpMod) value divisor) | local counter value -> do
    modulus <- intLiteral divisor
    guard (modulus > 0 && modulus <= top)
    pure (Remainders modulus)
  _ -> Nothing

emit :: GoExpr -> GoExpr -> Change -> GoExpr
emit counter original = case _ of
  Keep -> original
  Fields ty fields -> GoRecordUpdateNative ty original
    (map (\(Tuple key change) -> Tuple key (emit counter (GoRecordFieldAccess original key) change)) fields)
  Add delta ->
    -- At least one IntAdd executes on the admitted path. Normalize its seed
    -- first, so even noncanonical Int64 FFI inputs need no overflow assumption.
    -- n and each literal fit Int32; products and their sum fit signed Int64.
    narrow (GoBinOp "+" (narrow original) (total counter delta))

total :: GoExpr -> Delta -> GoExpr
total count = case _ of
  Constant value -> GoBinOp "*" count (integer value)
  Remainders modulus ->
    let divisor = integer modulus
        quotient = GoBinOp "/" count divisor
        remainder = GoBinOp "%" count divisor
        cycles = GoBinOp "/" (GoBinOp "*" (GoBinOp "*" quotient divisor) (integer (modulus - 1))) (integer 2)
        suffix = GoBinOp "/" (GoBinOp "*" remainder (GoBinOp "+" remainder (integer 1))) (integer 2)
    -- Sum k mod d for k=1..n: complete periods plus the final 1..remainder.
    in GoBinOp "+" cycles suffix

syntax :: TcoExpr -> BackendSyntax TcoExpr
syntax (TcoExpr _ value) = case value of
  Syn.Typed _ inner -> syntax inner
  Syn.TypeApp inner _ -> syntax inner
  _ -> value

local :: String -> TcoExpr -> Boolean
local name expression = case syntax expression of
  Syn.Local ident level -> localId ident level == name
  _ -> false

recordPath :: String -> TcoExpr -> Maybe (Array String)
recordPath root expression = case syntax expression of
  Syn.Local ident level | localId ident level == root -> Just []
  Syn.Accessor source (Syn.GetProp key) -> map (flip Array.snoc key) (recordPath root source)
  _ -> Nothing

intLiteral :: TcoExpr -> Maybe Int
intLiteral expression = case syntax expression of
  Syn.Lit (LitInt value) -> Just value
  _ -> Nothing

literal :: Int -> TcoExpr -> Boolean
literal value expression = intLiteral expression == Just value

narrow :: GoExpr -> GoExpr
narrow value = GoCall (GoVar "int64") [ GoCall (GoVar "int32") [ value ] ]

integer :: Int -> GoExpr
integer value = GoCall (GoVar "int64") [ GoInt value ]
