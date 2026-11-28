-- Specialize pure self-recursive Int -> ... -> Int functions containing at
-- least one non-tail self-call into an internal native int32 worker guarded
-- once at the int64 ABI entry boundary. Tail self-calls compile to a loop,
-- and right-operand IntAdd tail self-calls compile to an int32 accumulator
-- loop, preserving exact recursive structure and modulo-2^32 semantics.
module Gopurs.Int32Recursion (plan) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (foldl)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Newtype (unwrap)
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Tuple (Tuple(..), fst, snd)
import Gopurs.CallAnalysis (collectGoSpine, getGoSpineArgs)
import Gopurs.ExprAnalysis (unwrapTcoExpr)
import Gopurs.GoAst (GoDecl(..), GoExpr(..), GoType(..), rawGo)
import Gopurs.GoFunctions (Parameters)
import Gopurs.GoSymbol (GoReference(..), GoSymbol)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.CoreFn (Ident(..), Literal(..), Qualified(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax
  ( BackendOperator(..)
  , BackendOperator1(..)
  , BackendOperator2(..)
  , BackendOperatorNum(..)
  , BackendOperatorOrd(..)
  , BackendSyntax(..)
  , Pair(..)
  )

data ValType
  = VInt
  | VBool

derive instance eqValType :: Eq ValType

type Env = Map.Map String { goName :: String, ty :: ValType }

type Stats =
  { hasNonTailSelf :: Boolean
  , hasTailLoop :: Boolean
  , hasAcc :: Boolean
  }

emptyStats :: Stats
emptyStats = { hasNonTailSelf: false, hasTailLoop: false, hasAcc: false }

mergeStats :: Stats -> Stats -> Stats
mergeStats a b =
  { hasNonTailSelf: a.hasNonTailSelf || b.hasNonTailSelf
  , hasTailLoop: a.hasTailLoop || b.hasTailLoop
  , hasAcc: a.hasAcc || b.hasAcc
  }

data TailPlan
  = ReturnLeaf GoExpr
  | TailSelf (Array GoExpr)
  | TailAddSelf GoExpr (Array GoExpr)
  | TailLet String GoExpr TailPlan
  | TailBranch (Array (Tuple GoExpr TailPlan)) TailPlan

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
  guard (resultType == TypeInt64 && not (Array.null params))
  guard (Array.all (\p -> snd p == TypeInt64) params)
  let
    paramSourceNames = map fst params
    arity = Array.length params
  guard (Array.length (Array.nub paramSourceNames) == arity)
  let
    workerParamNames = Array.mapWithIndex (\i _ -> "p" <> show i) params
    initEnv = foldl
      (\acc (Tuple srcName goName) -> Map.insert srcName { goName, ty: VInt } acc)
      Map.empty
      (Array.zip paramSourceNames workerParamNames)
    ctx = { modName, ident, arity, i32Symbol }
  analyzed <- lowerTail ctx initEnv 0 body
  guard analyzed.stats.hasNonTailSelf
  let
    emitted = emitTail workerParamNames analyzed.stats.hasAcc analyzed.nextId analyzed.plan
    workerBody =
      if analyzed.stats.hasAcc then
        GoBlock [ rawGo "var __acc int32", GoForRange "" emitted.stmts ]
      else if analyzed.stats.hasTailLoop then
        GoBlock [ GoForRange "" emitted.stmts ]
      else
        GoBlock emitted.stmts
    workerDecl = GoFunctionDecl
      { name: i32Symbol
      , params: map (\p -> Tuple p (TypeInterface "int32")) workerParamNames
      , result: TypeInterface "int32"
      , body: workerBody
      }
    loopNames = map (\p -> fst p <> "_loop") params
    fits name = GoBinOp "==" (GoCall (GoVar "int64") [ GoCall (GoVar "int32") [ GoVar name ] ]) (GoVar name)
  entryGuard <- case Array.uncons loopNames of
    Just { head, tail } -> Just (foldl (GoBinOp "&&") (fits head) (map fits tail))
    Nothing -> Nothing
  let
    callArgs = map (\name -> GoCall (GoVar "int32") [ GoVar name ]) loopNames
    fast = GoBlock
      [ GoIfElse entryGuard
          [ GoReturn (GoCall (GoVar "int64") [ GoCall (GoRef (SymbolReference i32Symbol)) callArgs ]) ]
          []
      , GoReturn coldCall
      ]
  pure { worker: workerDecl, fast }

type LowerCtx =
  { modName :: String
  , ident :: String
  , arity :: Int
  , i32Symbol :: GoSymbol
  }

lowerTail
  :: LowerCtx
  -> Env
  -> Int
  -> TcoExpr
  -> Maybe { plan :: TailPlan, nextId :: Int, stats :: Stats }
lowerTail ctx env nextId expr = case unwrapTcoExpr expr of
  Branch branches fallback -> do
    branchRes <- foldMArray
      (\acc (Pair cond body) -> do
        c <- lowerExpr ctx env VBool cond
        b <- lowerTail ctx env acc.nextId body
        pure
          { pairs: Array.snoc acc.pairs (Tuple c.goExpr b.plan)
          , nextId: b.nextId
          , stats: mergeStats acc.stats (mergeStats c.stats b.stats)
          })
      { pairs: [], nextId, stats: emptyStats }
      (NEA.toArray branches)
    fb <- lowerTail ctx env branchRes.nextId fallback
    pure
      { plan: TailBranch branchRes.pairs fb.plan
      , nextId: fb.nextId
      , stats: mergeStats branchRes.stats fb.stats
      }
  Let mbIdent level boundVal rest -> do
    let key = localId mbIdent level
    bound <- lowerAny ctx env boundVal
    let
      goName = "__v" <> show nextId
      env' = Map.insert key { goName, ty: bound.ty } env
    body <- lowerTail ctx env' (nextId + 1) rest
    pure
      { plan: TailLet goName bound.goExpr body.plan
      , nextId: body.nextId
      , stats: mergeStats bound.stats body.stats
      }
  PrimOp (Op2 (OpIntNum OpAdd) left right) ->
    case selfCallArgs ctx.modName ctx.ident ctx.arity right of
      Just rightArgs -> do
        l <- lowerExpr ctx env VInt left
        rs <- lowerArgs ctx env rightArgs
        pure
          { plan: TailAddSelf l.goExpr rs.goExprs
          , nextId
          , stats: mergeStats l.stats (mergeStats rs.stats { hasNonTailSelf: false, hasTailLoop: true, hasAcc: true })
          }
      Nothing -> do
        v <- lowerExpr ctx env VInt expr
        pure { plan: ReturnLeaf v.goExpr, nextId, stats: v.stats }
  _ ->
    case selfCallArgs ctx.modName ctx.ident ctx.arity expr of
      Just args -> do
        rs <- lowerArgs ctx env args
        pure
          { plan: TailSelf rs.goExprs
          , nextId
          , stats: mergeStats rs.stats { hasNonTailSelf: false, hasTailLoop: true, hasAcc: false }
          }
      Nothing -> do
        v <- lowerExpr ctx env VInt expr
        pure { plan: ReturnLeaf v.goExpr, nextId, stats: v.stats }

lowerArgs :: LowerCtx -> Env -> Array TcoExpr -> Maybe { goExprs :: Array GoExpr, stats :: Stats }
lowerArgs ctx env =
  foldMArray
    (\acc arg -> do
      r <- lowerExpr ctx env VInt arg
      pure { goExprs: Array.snoc acc.goExprs r.goExpr, stats: mergeStats acc.stats r.stats })
    { goExprs: [], stats: emptyStats }

lowerAny :: LowerCtx -> Env -> TcoExpr -> Maybe { goExpr :: GoExpr, ty :: ValType, stats :: Stats }
lowerAny ctx env expr = case unwrapTcoExpr expr of
  Lit (LitInt n) | n >= bottom && n <= top ->
    Just { goExpr: GoCall (GoVar "int32") [ GoInt n ], ty: VInt, stats: emptyStats }
  Lit (LitBoolean b) ->
    Just { goExpr: GoVar (if b then "true" else "false"), ty: VBool, stats: emptyStats }
  Local mbIdent level -> do
    entry <- Map.lookup (localId mbIdent level) env
    pure { goExpr: GoVar entry.goName, ty: entry.ty, stats: emptyStats }
  PrimOp (Op1 op inner) -> case op of
    OpIntNegate -> do
      r <- lowerExpr ctx env VInt inner
      pure { goExpr: GoPrefixOp "-" r.goExpr, ty: VInt, stats: r.stats }
    OpIntBitNot -> do
      r <- lowerExpr ctx env VInt inner
      pure { goExpr: GoPrefixOp "^" r.goExpr, ty: VInt, stats: r.stats }
    OpBooleanNot -> do
      r <- lowerExpr ctx env VBool inner
      pure { goExpr: GoPrefixOp "!" r.goExpr, ty: VBool, stats: r.stats }
    _ -> Nothing
  PrimOp (Op2 op a b) -> case op of
    OpIntNum OpAdd -> intBin "+" a b
    OpIntNum OpSubtract -> intBin "-" a b
    OpIntBitAnd -> intBin "&" a b
    OpIntBitOr -> intBin "|" a b
    OpIntBitXor -> intBin "^" a b
    OpIntBitShiftLeft -> shiftBin "<<" a b
    OpIntBitShiftRight -> shiftBin ">>" a b
    OpIntOrd ord -> do
      ra <- lowerExpr ctx env VInt a
      rb <- lowerExpr ctx env VInt b
      pure { goExpr: GoBinOp (ordOp ord) ra.goExpr rb.goExpr, ty: VBool, stats: mergeStats ra.stats rb.stats }
    OpBooleanOrd OpEq -> boolBin "==" a b
    OpBooleanOrd OpNotEq -> boolBin "!=" a b
    OpBooleanAnd -> boolBin "&&" a b
    OpBooleanOr -> boolBin "||" a b
    _ -> Nothing
  _ -> case selfCallArgs ctx.modName ctx.ident ctx.arity expr of
    Just args -> do
      rs <- lowerArgs ctx env args
      pure
        { goExpr: GoCall (GoRef (SymbolReference ctx.i32Symbol)) rs.goExprs
        , ty: VInt
        , stats: mergeStats rs.stats { hasNonTailSelf: true, hasTailLoop: false, hasAcc: false }
        }
    Nothing -> Nothing
  where
  intBin sym a b = do
    ra <- lowerExpr ctx env VInt a
    rb <- lowerExpr ctx env VInt b
    pure { goExpr: GoBinOp sym ra.goExpr rb.goExpr, ty: VInt, stats: mergeStats ra.stats rb.stats }

  shiftBin sym a b = do
    ra <- lowerExpr ctx env VInt a
    rb <- lowerExpr ctx env VInt b
    let mask = GoBinOp "&" (GoCall (GoVar "uint32") [ rb.goExpr ]) (GoInt 31)
    pure { goExpr: GoBinOp sym ra.goExpr mask, ty: VInt, stats: mergeStats ra.stats rb.stats }

  boolBin sym a b = do
    ra <- lowerExpr ctx env VBool a
    rb <- lowerExpr ctx env VBool b
    pure { goExpr: GoBinOp sym ra.goExpr rb.goExpr, ty: VBool, stats: mergeStats ra.stats rb.stats }

lowerExpr :: LowerCtx -> Env -> ValType -> TcoExpr -> Maybe { goExpr :: GoExpr, stats :: Stats }
lowerExpr ctx env expected expr = do
  r <- lowerAny ctx env expr
  guard (r.ty == expected)
  pure { goExpr: r.goExpr, stats: r.stats }

ordOp :: BackendOperatorOrd -> String
ordOp = case _ of
  OpEq -> "=="
  OpNotEq -> "!="
  OpLt -> "<"
  OpLte -> "<="
  OpGt -> ">"
  OpGte -> ">="

emitTail :: Array String -> Boolean -> Int -> TailPlan -> { stmts :: Array GoExpr, nextId :: Int }
emitTail paramNames useAcc nextId = case _ of
  ReturnLeaf expr ->
    let
      retExpr =
        if not useAcc then expr
        else if expr == GoCall (GoVar "int32") [ GoInt 0 ] then GoVar "__acc"
        else GoBinOp "+" (GoVar "__acc") expr
    in { stmts: [ GoReturn retExpr ], nextId }
  TailSelf args ->
    let u = emitParamUpdates paramNames args nextId
    in { stmts: u.stmts <> [ rawGo "continue" ], nextId: u.nextId }
  TailAddSelf left args ->
    let u = emitParamUpdates paramNames args nextId
    in
      { stmts: [ GoMutate "__acc" (GoBinOp "+" (GoVar "__acc") left) ] <> u.stmts <> [ rawGo "continue" ]
      , nextId: u.nextId
      }
  TailLet goName valExpr rest ->
    let r = emitTail paramNames useAcc nextId rest
    in { stmts: Array.cons (GoAssign goName valExpr) r.stmts, nextId: r.nextId }
  TailBranch pairs fallback ->
    let
      emittedBranches = foldl
        (\acc (Tuple cond body) ->
          let r = emitTail paramNames useAcc acc.nextId body
          in { stmts: Array.snoc acc.stmts (GoIfElse cond r.stmts []), nextId: r.nextId })
        { stmts: [], nextId }
        pairs
      fb = emitTail paramNames useAcc emittedBranches.nextId fallback
    in { stmts: emittedBranches.stmts <> fb.stmts, nextId: fb.nextId }

emitParamUpdates :: Array String -> Array GoExpr -> Int -> { stmts :: Array GoExpr, nextId :: Int }
emitParamUpdates paramNames args nextId =
  let
    changed = Array.mapMaybe
      (\(Tuple p a) -> if a == GoVar p then Nothing else Just (Tuple p a))
      (Array.zip paramNames args)
  in case changed of
    [] -> { stmts: [], nextId }
    [ Tuple p a ] -> { stmts: [ GoMutate p a ], nextId }
    many ->
      let
        captured = foldl
          (\acc (Tuple p a) ->
            let tmp = "__t" <> show acc.nextId
            in
              { assigns: Array.snoc acc.assigns (GoAssign tmp a)
              , mutates: Array.snoc acc.mutates (GoMutate p (GoVar tmp))
              , nextId: acc.nextId + 1
              })
          { assigns: [], mutates: [], nextId }
          many
      in { stmts: captured.assigns <> captured.mutates, nextId: captured.nextId }

selfCallArgs :: String -> String -> Int -> TcoExpr -> Maybe (Array TcoExpr)
selfCallArgs owner name arity expr =
  let
    Tuple callee spine = collectGoSpine expr
    args = getGoSpineArgs spine
  in case unwrapTcoExpr callee of
    Var (Qualified moduleName (Ident ident))
      | ident == name && Array.length args == arity
        && (moduleName == Nothing || map (String.replaceAll (Pattern ".") (Replacement "_") <<< unwrap) moduleName == Just owner) ->
          Just args
    _ -> Nothing

foldMArray :: forall a b. (b -> a -> Maybe b) -> b -> Array a -> Maybe b
foldMArray f init = go init 0
  where
  go acc idx arr = case Array.index arr idx of
    Nothing -> Just acc
    Just x -> case f acc x of
      Nothing -> Nothing
      Just acc' -> go acc' (idx + 1) arr
