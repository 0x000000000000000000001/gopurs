module Gopurs.ListFilter (worker, workerArguments) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (all)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Newtype (unwrap)
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Tuple (Tuple(..), fst, snd)
import Gopurs.CallAnalysis (extractUncurriedAbs)
import Gopurs.CodegenState (CodegenMetadata)
import Gopurs.ConstructorLayout as Layout
import Gopurs.ExprAnalysis (getExprType, unwrapTcoExpr)
import Gopurs.ExprContext (ExprResult, StmtTree(..))
import Gopurs.GoAst (GoExpr(..), GoType(..), StructPointer, constructorNames, goTypeToStr, rawGo)
import Gopurs.GoFunctions (Parameters)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Ident(..), Qualified(..))
import PureScript.Backend.Optimizer.FfiSupport (hashString)
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (BackendAccessor(..), BackendOperator(..), BackendOperator1(..), Pair(..))
import PureScript.Backend.Optimizer.Syntax as Syn

type Shape = { empty :: Qualified Ident, cons :: Qualified Ident }
type Cases = { empty :: TcoExpr, cons :: TcoExpr }
type Loop = { self :: String, input :: String, acc :: String, body :: TcoExpr, initial :: Array TcoExpr }
type Filter = { predicate :: String, source :: String, pointer :: StructPointer, predicateType :: GoType }

-- Publish the native callback only after the same complete proof used by the
-- builder. Imported callers receive this ABI through ordinary worker metadata.
workerArguments :: CodegenMetadata -> Parameters -> GoType -> TcoExpr -> Maybe (Array GoType)
workerArguments metadata params resultType value = do
  _ <- recognize metadata params resultType value
  pure [ TypeFunc [ TypeInt64 ] TypeBool, resultType ]

-- Recognize an entire filter/prepend/reverse implementation before emitting a
-- forward builder. The input remains borrowed and immutable. Retained runs are
-- copied only when a later rejection requires a new link; the final retained
-- suffix is shared. Only unpublished, freshly allocated output links are
-- assigned. No usage facts or ownership assumption enable this transformation.
worker :: CodegenMetadata -> Parameters -> GoType -> TcoExpr -> Maybe ExprResult
worker metadata params resultType value = do
  found <- recognize metadata params resultType value
  pure (builder found)

recognize :: CodegenMetadata -> Parameters -> GoType -> TcoExpr -> Maybe Filter
recognize metadata params resultType value = do
  signature <- signatureOf value
  listType <- case signature.args, signature.result of
    [ Func [ Int ] Boolean, ty ], returned | ty == returned -> Just ty
    _, _ -> Nothing
  lambda <- extractUncurriedAbs value
  Tuple predicate source <- two lambda.args
  guard (map fst params == lambda.args)
  Tuple predicateType pointer <- case params, resultType of
    [ Tuple _ callback, Tuple _ input ], TypeStructPointer pointer
      | input == resultType && (callback == TypeValue || callback == TypeFunc [ TypeInt64 ] TypeBool) -> Just (Tuple callback pointer)
    _, _ -> Nothing
  loop <- localLoop listType lambda.body
  Tuple initial seed <- two loop.initial
  guard (local source initial)
  shape <- shapeFrom loop.input loop.body
  guard (emptyList shape seed)
  validateLayout metadata listType pointer shape
  cases <- listCases shape loop.input loop.body
  guard (reverseOf listType shape loop.acc cases.empty)
  branch <- case unwrapTcoExpr cases.cons of
    Syn.Branch choices no -> case NEA.toArray choices of
      [ Pair condition yes ] -> Just { condition, yes, no }
      _ -> Nothing
    _ -> Nothing
  let condition = application branch.condition
  argument <- single condition.args
  guard (local predicate condition.head && field shape loop.input 0 argument)
  Tuple yesTail yesAcc <- selfCall loop branch.yes >>= two
  Tuple noTail noAcc <- selfCall loop branch.no >>= two
  Tuple head tail <- consFields shape yesAcc
  guard (field shape loop.input 1 yesTail && field shape loop.input 1 noTail)
  guard (field shape loop.input 0 head && local loop.acc tail && local loop.acc noAcc)
  pure { predicate, source, pointer, predicateType }

signatureOf :: TcoExpr -> Maybe { args :: Array ExprType, result :: ExprType }
signatureOf value = case getExprType value of
  Func args result -> Just (flatten args result)
  _ -> Nothing
  where
  flatten args (Func more result) = flatten (args <> more) result
  flatten args result = { args, result }

two :: forall a. Array a -> Maybe (Tuple a a)
two [ a, b ] = Just (Tuple a b)
two _ = Nothing

single :: forall a. Array a -> Maybe a
single [ value ] = Just value
single _ = Nothing

localName :: TcoExpr -> Maybe String
localName value = case unwrapTcoExpr value of
  Syn.Local name level -> Just (localId name level)
  _ -> Nothing

local :: String -> TcoExpr -> Boolean
local name value = localName value == Just name

application :: TcoExpr -> { head :: TcoExpr, args :: Array TcoExpr }
application value = case unwrapTcoExpr value of
  Syn.App fn args -> let prior = application fn in { head: prior.head, args: prior.args <> NEA.toArray args }
  _ -> { head: value, args: [] }

-- The selected self-recursive function is the only executable initializer.
-- Other entries must be lambdas and the admitted body never calls them.
localLoop :: ExprType -> TcoExpr -> Maybe Loop
localLoop listType value = case unwrapTcoExpr value of
  Syn.LetRec level bindings body -> do
    let entry = application body
    target <- localName entry.head
    selected <- snd <$> Array.find (\(Tuple name _) -> localId (Just name) level == target) (NEA.toArray bindings)
    guard (all (\(Tuple _ binding) -> case extractUncurriedAbs binding of
      Just _ -> true
      Nothing -> false) bindings)
    signature <- signatureOf selected
    guard (signature.args == [ listType, listType ] && signature.result == listType)
    lambda <- extractUncurriedAbs selected
    Tuple input acc <- two lambda.args
    guard (input /= acc && target /= input && target /= acc)
    pure { self: target, input, acc, body: lambda.body, initial: entry.args }
  _ -> Nothing

selfCall :: Loop -> TcoExpr -> Maybe (Array TcoExpr)
selfCall loop value = do
  let call = application value
  guard (local loop.self call.head)
  pure call.args

testedTag :: String -> TcoExpr -> Maybe (Qualified Ident)
testedTag input value = case unwrapTcoExpr value of
  Syn.PrimOp (Op1 (OpIsTag ctor) argument) | local input argument -> Just ctor
  _ -> Nothing

shapeFrom :: String -> TcoExpr -> Maybe Shape
shapeFrom input value = case unwrapTcoExpr value of
  Syn.Branch choices failure -> case NEA.toArray choices, unwrapTcoExpr failure of
    [ Pair emptyTest _, Pair consTest _ ], Syn.Fail _ -> do
      empty <- testedTag input emptyTest
      cons <- testedTag input consTest
      guard (empty /= cons)
      pure { empty, cons }
    _, _ -> Nothing
  _ -> Nothing

listCases :: Shape -> String -> TcoExpr -> Maybe Cases
listCases shape input value = do
  found <- shapeFrom input value
  guard (found == shape)
  case unwrapTcoExpr value of
    Syn.Branch choices _ -> case NEA.toArray choices of
      [ Pair _ empty, Pair _ cons ] -> Just { empty, cons }
      _ -> Nothing
    _ -> Nothing

emptyList :: Shape -> TcoExpr -> Boolean
emptyList shape value = case unwrapTcoExpr value of
  Syn.CtorSaturated ctor _ _ _ fields -> ctor == shape.empty && Array.null fields
  Syn.Var ctor -> ctor == shape.empty
  _ -> false

field :: Shape -> String -> Int -> TcoExpr -> Boolean
field shape input index value = case unwrapTcoExpr value of
  Syn.Accessor base (GetCtorField ctor _ _ _ _ found) -> ctor == shape.cons && index == found && local input base
  _ -> false

consFields :: Shape -> TcoExpr -> Maybe (Tuple TcoExpr TcoExpr)
consFields shape value = case unwrapTcoExpr value of
  Syn.CtorSaturated ctor _ _ _ [ Tuple _ head, Tuple _ tail ] | ctor == shape.cons -> Just (Tuple head tail)
  _ -> Nothing

reverseOf :: ExprType -> Shape -> String -> TcoExpr -> Boolean
reverseOf listType shape source value = case reverseLoop of
  Just _ -> true
  Nothing -> false
  where
  reverseLoop = do
    loop <- localLoop listType value
    Tuple initial seed <- two loop.initial
    guard (local source initial && emptyList shape seed)
    cases <- listCases shape loop.input loop.body
    guard (local loop.acc cases.empty)
    Tuple next constructed <- selfCall loop cases.cons >>= two
    Tuple head tail <- consFields shape constructed
    guard (field shape loop.input 1 next && field shape loop.input 0 head && local loop.acc tail)

-- TAST metadata establishes one payload constructor plus a nil alternative,
-- their defining module, native instantiation and exact two-field layout.
-- Restrict this first lowering to Int payloads; erased/generic layouts retain
-- the ordinary implementation and its normal boxing conversions.
validateLayout :: CodegenMetadata -> ExprType -> StructPointer -> Shape -> Maybe Unit
validateLayout metadata listType pointer shape = do
  owner <- case shape.cons, shape.empty of
    Qualified (Just owner) _, Qualified (Just emptyOwner) _ | owner == emptyOwner -> Just owner
    _, _ -> Nothing
  let Qualified _ (Ident cons) = shape.cons
      Qualified _ (Ident empty) = shape.empty
      moduleName = String.replaceAll (Pattern ".") (Replacement "_") (unwrap owner)
      consNames = constructorNames moduleName cons
      emptyNames = constructorNames moduleName empty
      nativeType = TypeStructPointer pointer
      layout = Layout.layout metadata moduleName cons (Just nativeType)
  guard (pointer.symbol == consNames.symbol && pointer.baseStructName == consNames.baseStructName)
  guard case listType of
    ADT fullName _ args -> fullName == pointer.fullName && all (_ == Int) args
    _ -> false
  path <- Map.lookup pointer.fullName metadata.pointerAdtPaths
  guard (path.ctorName == cons && path.arity == Array.length pointer.typeArgs)
  leaf <- Map.lookup emptyNames.baseStructName metadata.pointerAdtLeaves
  guard (leaf.nodeBaseStruct == pointer.baseStructName && leaf.nodeCtor == cons)
  fields <- layout.constructorFields
  guard (Array.length fields.vars == Array.length pointer.typeArgs && Array.length fields.fields == 2)
  guard (Layout.fieldType metadata moduleName fields pointer.typeArgs 0 == TypeInt64)
  guard (Layout.fieldType metadata moduleName fields pointer.typeArgs 1 == nativeType)

builder :: Filter -> ExprResult
builder { predicate, source, pointer, predicateType } =
  let
    resultType = TypeStructPointer pointer
    nil = rawGo ("(" <> goTypeToStr resultType <> ")(nil)")
    cursor = GoVar "cursor"
    run = GoVar "run"
    output = GoVar "output"
    node = GoVar "node"
    head = GoStructAccess cursor "V0"
    selected = if predicateType == TypeFunc [ TypeInt64 ] TypeBool then
      GoCall (GoVar "predicate") [ head ]
      else GoBinOp "!="
        (GoStructAccess (GoCall (GoSelector (GoVar "gopurs_runtime") "Apply")
          [ GoVar "predicate", GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ head ] ]) "IntVal")
        (GoInt 0)
    append =
      [ GoAssign "node" (GoConstructor (hashString pointer.baseStructName) pointer.symbol pointer.typeArgs
          [ GoStructAccess run "V0", nil ])
      , GoIfElse (GoBinOp "==" output nil)
          [ GoMutate "output" node ] [ GoMutate "tail.V1" node ]
      , GoMutate "tail" node
      ]
    advance = GoMutate "cursor" (GoStructAccess cursor "V1")
    -- Each scan leaves its first opposite element at cursor. Consume that
    -- element explicitly before switching phases, so neither callback ABI
    -- ever sees it twice. Rejected runs need no copying or accumulator work.
    scanRejected = GoForRange "cursor != nil"
      [ GoIfElse selected [ rawGo "break" ] [], advance ]
    scanRetained = GoForRange "cursor != nil"
      [ GoIfElse (GoPrefixOp "!" selected) [ rawGo "break" ] [], advance ]
    finish =
      [ GoIfElse (GoBinOp "==" output nil) [ GoReturn run ] []
      , GoMutate "tail.V1" run
      , GoReturn output
      ]
    body =
      [ GoAssign "output" nil
      , GoAssign "tail" nil
      , GoFor "filter_forward"
          [ scanRejected
          , GoIfElse (GoBinOp "==" cursor nil) [ GoReturn output ] []
          , GoAssign "run" cursor
          , advance
          , scanRetained
          , GoIfElse (GoBinOp "==" cursor nil) finish []
          , GoForRange "run != cursor" (append <> [ GoMutate "run" (GoStructAccess run "V1") ])
          , advance
          ]
      ]
    -- An explicit-parameter IIFE gives private names their own scope and
    -- evaluates both arguments before allocating any output. The predicate's
    -- invocation count and traversal order stay intact for either ABI. Copying
    -- a retained run never calls the predicate again. Until the final return,
    -- tail always denotes a fresh node, never a node of the borrowed suffix.
    expr = GoCall (GoFuncBlock [ Tuple "predicate" predicateType, Tuple "cursor" resultType ] body resultType)
      [ GoVar predicate, GoVar source ]
  in { stmts: StmtEmpty, expr, exprType: resultType, nextId: 0 }
