module Gopurs.CongruenceSum (Congruence, recognize, body) where

import Prelude

import Control.Alternative (guard)
import Data.Maybe (Maybe(..))
import Gopurs.GoAst (GoExpr(..))
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (Literal(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (BackendSyntax)
import PureScript.Backend.Optimizer.Syntax as Syn

type Congruence = { modulus :: Int, residue :: Int }

-- Only the complete pure expression is admitted: no calls, let bindings or
-- additional operations may disappear when its per-element calls are removed.
recognize :: String -> TcoExpr -> Maybe Congruence
recognize parameter expression = do
  congruence <- case syntax expression of
    Syn.PrimOp (Syn.Op2 (Syn.OpIntOrd Syn.OpEq) left right) ->
      case intLiteral right, intLiteral left of
        Just residue, _ -> Just { value: left, residue }
        _, Just residue -> Just { value: right, residue }
        _, _ -> Nothing
    _ -> Nothing
  modulus <- case syntax congruence.value of
    Syn.PrimOp (Syn.Op2 (Syn.OpIntNum Syn.OpMod) left right)
      | localName left == Just parameter -> intLiteral right
    _ -> Nothing
  guard (modulus >= 1 && modulus <= top && congruence.residue >= 0 && congruence.residue < modulus)
  pure { modulus, residue: congruence.residue }

-- The caller owns an isolated Go function scope and proves
-- 1 <= lower <= upper <= MaxInt32, with a signed Int32 seed. All intermediates
-- use Go int64: even the first selected value can exceed MaxInt32.
body :: Congruence -> { lower :: GoExpr, upper :: GoExpr, seed :: GoExpr } -> Array GoExpr
body { modulus, residue } { lower, upper, seed } =
  let
    first = GoVar "first"
    count = GoVar "count"
    last = GoVar "last"
    step = integer modulus
    offset = GoBinOp "%" (GoBinOp "+"
      (GoBinOp "-" (integer residue) (GoBinOp "%" lower step)) step) step
    -- When a term exists, count <= top and first,last <= top. Therefore
    -- count*(first+last) <= 2*top^2 < MaxInt64; division is exact and adding
    -- seed still fits. IntMul would round via JS Number, so it is not used.
    -- A final Int32 conversion preserves the modular sum of all additions.
    total = GoBinOp "/" (GoBinOp "*" count (GoBinOp "+" first last)) (integer 2)
  in
    [ GoAssign "first" (GoBinOp "+" lower offset)
    , GoIfElse (GoBinOp ">" first upper) [ GoReturn seed ] []
    , GoAssign "count" (GoBinOp "+" (GoBinOp "/" (GoBinOp "-" upper first) step) (integer 1))
    , GoAssign "last" (GoBinOp "+" first (GoBinOp "*" (GoBinOp "-" count (integer 1)) step))
    , GoReturn (cast "int64" (cast "int32" (GoBinOp "+" seed total)))
    ]

syntax :: TcoExpr -> BackendSyntax TcoExpr
syntax (TcoExpr _ value) = case value of
  Syn.Typed _ inner -> syntax inner
  _ -> value

localName :: TcoExpr -> Maybe String
localName value = case syntax value of
  Syn.Local name level -> Just (localId name level)
  _ -> Nothing

intLiteral :: TcoExpr -> Maybe Int
intLiteral value = case syntax value of
  Syn.Lit (LitInt number) -> Just number
  _ -> Nothing

cast :: String -> GoExpr -> GoExpr
cast name value = GoCall (GoVar name) [ value ]

integer :: Int -> GoExpr
integer = cast "int64" <<< GoInt
