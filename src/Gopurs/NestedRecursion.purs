-- Optimize pure two-parameter nested self-recursive Int functions where the
-- outer parameter steps toward an affine base case (e.g. f(c0, b) = b + k,
-- f(a, b0) = f(a + stepA, bReset), f(a, b) = f(a + stepA, f(a, b + stepB))).
-- Emits an internal native int32 worker guarded at entry for wide inputs, and
-- jump-threads the slice where `a + stepA == c0` into a zero-stack modular sum.
module Gopurs.NestedRecursion (plan) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Maybe (Maybe(..))
import Data.Newtype (unwrap)
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Tuple (Tuple(..))
import Gopurs.CallAnalysis (collectGoSpine, getGoSpineArgs)
import Gopurs.ExprAnalysis (unwrapTcoExpr)
import Gopurs.GoAst (GoDecl(..), GoExpr(..), GoType(..), rawGo)
import Gopurs.GoFunctions (Parameters)
import Gopurs.GoSymbol (GoReference(..), GoSymbol)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.CoreFn (Ident(..), Literal(..), Qualified(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (BackendOperator(..), BackendOperator2(..), BackendOperatorNum(..), BackendOperatorOrd(..), BackendSyntax(..), Pair(..))

type Spec =
  { swapped :: Boolean
  , c0 :: Int
  , k :: Int
  , b0 :: Int
  , stepA :: Int
  , bReset :: Int
  , stepB :: Int
  }

plan
  :: String
  -> String
  -> Parameters
  -> GoType
  -> TcoExpr
  -> GoSymbol
  -> GoExpr
  -> Maybe { worker :: GoDecl, fast :: GoExpr }
plan modName ident params resultType body i32Symbol coldCall = do
  guard (resultType == TypeInt64)
  Tuple p0 p1 <- case params of
    [ Tuple a TypeInt64, Tuple b TypeInt64 ] -> Just (Tuple a b)
    _ -> Nothing
  spec <- recognize modName ident p0 p1 body
  let
    p0Loop = p0 <> "_loop"
    p1Loop = p1 <> "_loop"
    fits name = GoBinOp "==" (GoCall (GoVar "int64") [ GoCall (GoVar "int32") [ GoVar name ] ]) (GoVar name)
    entryGuard = GoBinOp "&&" (fits p0Loop) (fits p1Loop)
    callArgs =
      if spec.swapped then
        [ GoCall (GoVar "int32") [ GoVar p1Loop ], GoCall (GoVar "int32") [ GoVar p0Loop ] ]
      else
        [ GoCall (GoVar "int32") [ GoVar p0Loop ], GoCall (GoVar "int32") [ GoVar p1Loop ] ]
    fast = GoBlock
      [ GoIfElse entryGuard
          [ GoReturn (GoCall (GoVar "int64") [ GoCall (GoRef (SymbolReference i32Symbol)) callArgs ]) ]
          []
      , GoReturn coldCall
      ]
    workerDecl = buildWorker i32Symbol spec
  pure { worker: workerDecl, fast }

buildWorker :: GoSymbol -> Spec -> GoDecl
buildWorker sym spec =
  let
    a = GoVar "a"
    b = GoVar "b"
    i32 n = GoCall (GoVar "int32") [ GoInt n ]
    baseReturn =
      if spec.k == 0 then b
      else GoBinOp "+" b (i32 spec.k)
    aStep = GoBinOp "+" a (i32 spec.stepA)
    bStep = GoBinOp "+" b (i32 spec.stepB)
    -- Only jump-thread the terminating direction toward b0 so non-terminating
    -- inputs on the opposite side retain the original call semantics.
    bDirGuard =
      if spec.stepB == -1 then GoBinOp ">" b (i32 spec.b0)
      else GoBinOp "<" b (i32 spec.b0)
    distExpr =
      if spec.stepB == -1 then GoBinOp "-" b (i32 spec.b0)
      else GoBinOp "-" (i32 spec.b0) b
    accExpr =
      if spec.k == 0 then i32 0
      else if spec.k == 1 then distExpr
      else GoBinOp "*" distExpr (i32 spec.k)
    jumpResult =
      GoBinOp "+" (GoBinOp "+" accExpr (i32 spec.bReset)) (i32 spec.k)
  in
    GoFunctionDecl
      { name: sym
      , params: [ Tuple "a" (TypeInterface "int32"), Tuple "b" (TypeInterface "int32") ]
      , result: TypeInterface "int32"
      , body: GoBlock
          [ GoForRange ""
              [ GoIfElse (GoBinOp "==" a (i32 spec.c0))
                  [ GoReturn baseReturn ] []
              , GoIfElse (GoBinOp "==" b (i32 spec.b0))
                  [ GoMutate "a" aStep
                  , GoMutate "b" (i32 spec.bReset)
                  , rawGo "continue"
                  ] []
              , GoIfElse (GoBinOp "&&" (GoBinOp "==" aStep (i32 spec.c0)) bDirGuard)
                  [ GoReturn jumpResult ] []
              , GoMutate "b" (GoCall (GoRef (SymbolReference sym)) [ a, bStep ])
              , GoMutate "a" aStep
              ]
          ]
      }

recognize :: String -> String -> String -> String -> TcoExpr -> Maybe Spec
recognize modName ident p0 p1 body = do
  { c1, e1, c2, e2, fallback } <- flattenTwoBranches body
  case tryOrder false p0 p1 c1 e1 c2 e2 fallback of
    Just s -> Just s
    Nothing -> tryOrder true p1 p0 c1 e1 c2 e2 fallback
  where
  tryOrder swapped aName bName c1 e1 c2 e2 fallback = do
    c0 <- matchEqConst aName c1
    k <- matchAffine bName e1
    b0 <- matchEqConst bName c2
    tailArgs <- selfCallArgs modName ident 2 e2
    nestedArgs <- selfCallArgs modName ident 2 fallback
    Tuple aNext1 bResetExpr <- orderedPair swapped tailArgs
    Tuple aNext2 innerExpr <- orderedPair swapped nestedArgs
    stepA1 <- matchStep aName aNext1
    stepA2 <- matchStep aName aNext2
    guard (stepA1 /= 0 && stepA1 == stepA2)
    bReset <- matchSignedInt32 bResetExpr
    innerArgs <- selfCallArgs modName ident 2 innerExpr
    Tuple innerA innerB <- orderedPair swapped innerArgs
    guard (isLocalName aName innerA)
    stepB <- matchStep bName innerB
    guard (stepB == -1 || stepB == 1)
    pure { swapped, c0, k, b0, stepA: stepA1, bReset, stepB }

orderedPair :: forall a. Boolean -> Array a -> Maybe (Tuple a a)
orderedPair swapped = case _ of
  [ x, y ] -> Just (if swapped then Tuple y x else Tuple x y)
  _ -> Nothing

flattenTwoBranches :: TcoExpr -> Maybe { c1 :: TcoExpr, e1 :: TcoExpr, c2 :: TcoExpr, e2 :: TcoExpr, fallback :: TcoExpr }
flattenTwoBranches expr = case unwrapTcoExpr expr of
  Branch branches fallback -> case NEA.toArray branches of
    [ Pair c1 e1, Pair c2 e2 ] -> Just { c1, e1, c2, e2, fallback }
    [ Pair c1 e1 ] -> case unwrapTcoExpr fallback of
      Branch branches2 fallback2 -> case NEA.toArray branches2 of
        [ Pair c2 e2 ] -> Just { c1, e1, c2, e2, fallback: fallback2 }
        _ -> Nothing
      _ -> Nothing
    _ -> Nothing
  _ -> Nothing

selfCallArgs :: String -> String -> Int -> TcoExpr -> Maybe (Array TcoExpr)
selfCallArgs owner name arity expr =
  let
    Tuple callee spine = collectGoSpine expr
    args = getGoSpineArgs spine
  in
    case unwrapTcoExpr callee of
      Var (Qualified moduleName (Ident ident))
        | ident == name && Array.length args == arity
          && (moduleName == Nothing || map (String.replaceAll (Pattern ".") (Replacement "_") <<< unwrap) moduleName == Just owner) ->
            Just args
      _ -> Nothing

isLocalName :: String -> TcoExpr -> Boolean
isLocalName expected expr = case unwrapTcoExpr expr of
  Local mbIdent level -> localId mbIdent level == expected
  _ -> false

matchSignedInt32 :: TcoExpr -> Maybe Int
matchSignedInt32 expr = case unwrapTcoExpr expr of
  Lit (LitInt n) | n >= bottom && n <= top -> Just n
  _ -> Nothing

matchEqConst :: String -> TcoExpr -> Maybe Int
matchEqConst param expr = case unwrapTcoExpr expr of
  PrimOp (Op2 (OpIntOrd OpEq) left right)
    | isLocalName param left -> matchSignedInt32 right
    | isLocalName param right -> matchSignedInt32 left
  _ -> Nothing

matchAffine :: String -> TcoExpr -> Maybe Int
matchAffine param expr = case unwrapTcoExpr expr of
  Local mbIdent level | localId mbIdent level == param -> Just 0
  PrimOp (Op2 (OpIntNum OpAdd) left right)
    | isLocalName param left -> matchSignedInt32 right
    | isLocalName param right -> matchSignedInt32 left
  PrimOp (Op2 (OpIntNum OpSubtract) left right)
    | isLocalName param left -> do
        n <- matchSignedInt32 right
        guard (n > bottom)
        pure (-n)
  _ -> Nothing

matchStep :: String -> TcoExpr -> Maybe Int
matchStep param expr = case unwrapTcoExpr expr of
  PrimOp (Op2 (OpIntNum OpAdd) left right)
    | isLocalName param left -> matchSignedInt32 right
    | isLocalName param right -> matchSignedInt32 left
  PrimOp (Op2 (OpIntNum OpSubtract) left right)
    | isLocalName param left -> do
        n <- matchSignedInt32 right
        guard (n > bottom)
        pure (-n)
  _ -> Nothing
