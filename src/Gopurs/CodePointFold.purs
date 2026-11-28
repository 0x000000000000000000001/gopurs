module Gopurs.CodePointFold (optimize, emit, emitBounded, describe, validDescriptor, compatibleT, primitiveCallback) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (all, foldM)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe, isJust)
import Data.Tuple (Tuple(..))
import Gopurs.ArrayIndexedFold (readOnlyScalar, scalarType)
import Gopurs.CallAnalysis (collectCurriedAbs, collectGoSpine, getGoSpineArgs, qualifiedTarget)
import Gopurs.CallArguments as CallArguments
import Gopurs.CodegenState (CodePointFoldDescriptor)
import Gopurs.ExprAnalysis (unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, StmtTree(..), TranslateExpr, bindParameters, childContext, flattenStmts)
import Gopurs.GoAst (GoExpr(..), GoType(..), rawGo)
import Gopurs.GoConversions (coerceGoExpr)
import Gopurs.GoTypes (appliedAdt)
import Gopurs.ImmediateApplications.Scope (maximumLevel)
import PureScript.Backend.Optimizer.Codegen.Tco as Tco
import PureScript.Backend.Optimizer.Convert (BackendModule)
import PureScript.Backend.Optimizer.CoreFn (ExprType(Any, Array, Func, Int, String, ADT, TypeVar), Ident(..), Literal(..), ModuleName(..), Qualified(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), Level(..))

-- Expose only a closed primitive fold composed with the canonical conversion.
-- Its callback and literal seed have no creation effects to postpone. Running
-- before worker discovery also puts the new signature in the usual cache key.
optimize :: BackendModule -> BackendModule
optimize mod = mod { bindings = map (\group -> if group.recursive then group else
  group { bindings = map (\(Tuple name value) -> Tuple name (fromMaybe value (expand value))) group.bindings }) mod.bindings }

expand :: NeutralExpr -> Maybe NeutralExpr
expand original = do
  guard (compatible (Func [ String ] Int) original)
  let composition = application original
  guard (namedN "Control.Semigroupoid" "compose" composition.head)
  Tuple folder conversion <- case composition.args of
    [ dictionary, folder, conversion ] | namedN "Control.Semigroupoid" "semigroupoidFn" dictionary -> Just (Tuple folder conversion)
    _ -> Nothing
  guard (namedN "Data.String.CodePoints" "toCodePointArray" conversion)
  guard (compatible (Func [ String ] (Array Int)) conversion && compatible (Func [ Array Int ] Int) folder)
  let folding = application folder
  guard (namedN "Data.Foldable" "foldlArray" folding.head)
  guard (compatible (Func [ Func [ Int, Int ] Int, Int, Array Int ] Int) folding.head)
  Tuple callback seed <- case folding.args of
    [ callback, seed ] -> Just (Tuple callback seed)
    _ -> Nothing
  guard (compatible Int seed && case strip seed of
    Lit (LitInt _) -> true
    _ -> false)
  guard (closed [] callback && isJust (primitiveCallback (Tco.analyze [] callback)))
  let
    level = Level (maximumLevel original + 1)
    name = Just (Ident "__gopurs_codepoint_input")
    input = NeutralExpr (Typed String (NeutralExpr (Local name level)))
    converted = NeutralExpr (Typed (Array Int) (NeutralExpr (App conversion (NEA.singleton input))))
    result = NeutralExpr (Typed Int (NeutralExpr (App folding.head (NEA.cons' callback [ seed, converted ]))))
  pure (NeutralExpr (Typed (Func [ String ] Int) (NeutralExpr (Abs (NEA.singleton (Tuple name level)) result))))

-- The only removed allocation belongs to an immediately consumed, canonical
-- conversion of an immutable string. Decoding is total; callback invocation
-- order and each Int operation remain unchanged. Unknown callbacks keep the
-- ordinary array path, even when their annotations claim a pure function.
emit :: TranslateExpr -> ExprContext -> Int -> Tco.TcoExpr -> Array Tco.TcoExpr -> Maybe ExprResult
emit translate context nextId fn args =
  if named "Data.String.CodePoints" "length" fn && Array.length args == 1 && Array.null context.recVars && compatibleT (Func [ String ] Int) fn then do
    input <- Array.head args
    guard (compatibleT String input)
    let
      translated = translate ((childContext context (Just String)) { mbExpectedGoType = Just TypeString }) nextId input
      strExpr = coerceGoExpr context.codegenStateRef context.modNameStr translated.expr translated.exprType TypeString
    pure
      { stmts: translated.stmts
      , expr: GoCall (GoSelector (GoVar "gopurs_runtime") "CountUTF8CodePoints") [ strExpr ]
      , exprType: TypeInt64
      , nextId: translated.nextId
      }
  else do
    guard (named "Data.Foldable" "foldlArray" fn && Array.null context.recVars)
    Tuple callback (Tuple seed converted) <- case args of
      [ callback, seed, converted ] -> Just (Tuple callback (Tuple seed converted))
      _ -> Nothing
    let Tuple conversion spine = collectGoSpine converted
    guard (named "Data.String.CodePoints" "toCodePointArray" conversion)
    guard (compatibleT (Func [ String ] (Array Int)) conversion)
    input <- case getGoSpineArgs spine of
      [ input ] -> Just input
      _ -> Nothing
    grouped <- primitiveCallback callback
    guard (compatibleT Int seed && compatibleT String input && compatibleT (Array Int) converted)
    pure (emitLoop translate context nextId callback seed input Nothing grouped)

-- The take consumer preserves the callback and seed proof and supplies a count
-- captured before the string. The ordinary unbounded path remains unchanged.
emitBounded :: TranslateExpr -> ExprContext -> Int -> Tco.TcoExpr -> Tco.TcoExpr -> Tco.TcoExpr -> Tco.TcoExpr -> Maybe ExprResult
emitBounded translate context nextId callback seed count input = do
  guard (Array.null context.recVars && compatibleT Int seed && compatibleT Int count && compatibleT String input)
  grouped <- primitiveCallback callback
  pure (emitLoop translate context nextId callback seed input (Just count) grouped)

emitLoop :: TranslateExpr -> ExprContext -> Int -> Tco.TcoExpr -> Tco.TcoExpr -> Tco.TcoExpr -> Maybe Tco.TcoExpr -> { args :: NEA.NonEmptyArray (Tuple (Maybe Ident) Level), body :: Tco.TcoExpr } -> ExprResult
emitLoop translate context nextId callback seed input limit grouped =
  let
    params = map (\(Tuple name level) -> Tuple (localId name level) TypeInt64) (NEA.toArray grouped.args)
    callbackContext = context
      { depth = context.depth + 1, bound = bindParameters params context.bound
      , tcoIdent = Nothing, loopCtx = [], options = { isTail: true, inEffectBlock: false }
      , mbExpectedExprType = Just Int, mbExpectedGoType = Just TypeInt64 }
    capture index currentId value = if index == 0 then
      let callbackBody = translate callbackContext currentId grouped.body
      in { stmts: StmtEmpty, expr: GoFuncLit params (flattenStmts callbackBody.stmts) (integer callbackBody.expr callbackBody.exprType) TypeInt64
         , exprType: TypeFunc [ TypeInt64, TypeInt64 ] TypeInt64, nextId: callbackBody.nextId }
      else translate (childContext context Nothing) currentId value
    captured = CallArguments.capture { prefix: if isJust limit then "codePointTake_arg_" else "codePointFold_arg_", markUsed: true } capture nextId
      (case limit of
        Just count -> [ callback, seed, count, input ]
        Nothing -> [ callback, seed, input ])
    directAdd = case Array.head captured.exprs of
      -- Inspect the translated callback AST directly (not just source syntax)
      -- so synthetic test translators retain their exact step semantics.
      Just (GoVar _) ->
        let
          checkBody = translate callbackContext captured.nextId grouped.body
          accName = fromMaybe "" (map (\(Tuple name _) -> name) (Array.index params 0))
          cpName = fromMaybe "" (map (\(Tuple name _) -> name) (Array.index params 1))
        in Array.null (flattenStmts checkBody.stmts) && case integer checkBody.expr checkBody.exprType of
          GoCall (GoSelector (GoVar "gopurs_runtime") "IntAdd") [ GoVar a, GoVar b ] ->
            (a == accName && b == cpName) || (a == cpName && b == accName)
          _ -> false
      _ -> false
    arg index = fromMaybe (rawGo "nil") (Array.index captured.exprs index)
    argType index = fromMaybe TypeValue (Array.index captured.exprTypes index)
    inputIndex = if isJust limit then 3 else 2
    str = coerceGoExpr context.codegenStateRef context.modNameStr (arg inputIndex) (argType inputIndex) TypeString
    stepStmt valueExpr =
      if directAdd then
        GoMutate "codePointFold_acc" (GoBinOp "+" (GoVar "codePointFold_acc") valueExpr)
      else
        GoMutate "codePointFold_result" (GoCall (arg 0) [ GoVar "codePointFold_result", valueExpr ])
    bounded = isJust limit
    preparation =
      [ GoAssign "codePointFold_rem" (GoVar "codePointFold_input") ]
        <> (if directAdd then [ GoAssign "codePointFold_acc" (rawGo "int64(0)") ] else [])
        <> (if bounded then [ GoAssign "codePointTake_limit" (integer (arg 2) (argType 2)), GoAssign "codePointTake_count" (rawGo "int64(0)") ] else [])
    mainHeader = if bounded then
      "; codePointTake_count < codePointTake_limit && len(codePointFold_rem) >= 4; codePointTake_count++"
      else "; len(codePointFold_rem) >= 4;"
    tailHeader = if bounded then
      "; codePointTake_count < codePointTake_limit && len(codePointFold_rem) > 0; codePointTake_count++"
      else "; len(codePointFold_rem) > 0;"
    -- Shrinking string slice `rem` with `len(rem) >= 4` eliminates all bounds
    -- checks on `rem[0..3]` and `rem[1..4:]` inside Go's prove pass.
    -- Surrogates (0xED) and malformed sequences delegate to DecodeStringCodePoint.
    mainLoop = GoForRange mainHeader
      [ GoAssign "codePointFold_b0" (GoIndex (GoVar "codePointFold_rem") (GoInt 0))
      , GoIfElse (rawGo "codePointFold_b0 < 0x80")
          [ stepStmt (rawGo "int64(codePointFold_b0)")
          , GoMutate "codePointFold_rem" (rawGo "codePointFold_rem[1:]")
          ]
          [ GoIfElse (rawGo "codePointFold_b0 < 0xE0")
              [ GoAssign "codePointFold_b1" (GoIndex (GoVar "codePointFold_rem") (GoInt 1))
              , GoIfElse (rawGo "codePointFold_b0 >= 0xC2 && codePointFold_b1&0xC0 == 0x80")
                  [ stepStmt (rawGo "int64(codePointFold_b0&0x1F)<<6 | int64(codePointFold_b1&0x3F)")
                  , GoMutate "codePointFold_rem" (rawGo "codePointFold_rem[2:]")
                  ]
                  [ stepStmt (rawGo "int64(0xFFFD)")
                  , GoMutate "codePointFold_rem" (rawGo "codePointFold_rem[1:]")
                  ]
              ]
              [ GoIfElse (rawGo "codePointFold_b0 < 0xF0")
                  [ GoAssign "codePointFold_b1" (GoIndex (GoVar "codePointFold_rem") (GoInt 1))
                  , GoAssign "codePointFold_b2" (GoIndex (GoVar "codePointFold_rem") (GoInt 2))
                  , GoIfElse (rawGo "codePointFold_b0 != 0xED && ((codePointFold_b1^0x80)|(codePointFold_b2^0x80)) < 0x40 && (codePointFold_b0 != 0xE0 || codePointFold_b1 >= 0xA0)")
                      [ stepStmt (rawGo "int64(codePointFold_b0&0x0F)<<12 | int64(codePointFold_b1&0x3F)<<6 | int64(codePointFold_b2&0x3F)")
                      , GoMutate "codePointFold_rem" (rawGo "codePointFold_rem[3:]")
                      ]
                      [ rawGo "codePointFold_value, codePointFold_width := gopurs_runtime.DecodeStringCodePoint(codePointFold_rem)"
                      , stepStmt (rawGo "int64(codePointFold_value)")
                      , GoMutate "codePointFold_rem" (rawGo "codePointFold_rem[codePointFold_width:]")
                      ]
                  ]
                  [ GoAssign "codePointFold_b1" (GoIndex (GoVar "codePointFold_rem") (GoInt 1))
                  , GoAssign "codePointFold_b2" (GoIndex (GoVar "codePointFold_rem") (GoInt 2))
                  , GoAssign "codePointFold_b3" (GoIndex (GoVar "codePointFold_rem") (GoInt 3))
                  , GoIfElse (rawGo "codePointFold_b0 <= 0xF4 && ((codePointFold_b1^0x80)|(codePointFold_b2^0x80)|(codePointFold_b3^0x80)) < 0x40 && (codePointFold_b0 != 0xF0 || codePointFold_b1 >= 0x90) && (codePointFold_b0 != 0xF4 || codePointFold_b1 <= 0x8F)")
                      [ stepStmt (rawGo "int64(codePointFold_b0&0x07)<<18 | int64(codePointFold_b1&0x3F)<<12 | int64(codePointFold_b2&0x3F)<<6 | int64(codePointFold_b3&0x3F)")
                      , GoMutate "codePointFold_rem" (rawGo "codePointFold_rem[4:]")
                      ]
                      [ stepStmt (rawGo "int64(0xFFFD)")
                      , GoMutate "codePointFold_rem" (rawGo "codePointFold_rem[1:]")
                      ]
                  ]
              ]
          ]
      ]
    tailLoop = GoForRange tailHeader
      [ rawGo "codePointFold_value, codePointFold_width := gopurs_runtime.DecodeStringCodePoint(codePointFold_rem)"
      , stepStmt (rawGo "int64(codePointFold_value)")
      , GoMutate "codePointFold_rem" (rawGo "codePointFold_rem[codePointFold_width:]")
      ]
    finalize = if directAdd then
      [ GoIfElse (rawGo "len(codePointFold_rem) < len(codePointFold_input)")
          [ GoMutate "codePointFold_result" (GoCall (GoVar "int64") [ GoCall (GoVar "int32") [ GoBinOp "+" (GoVar "codePointFold_result") (GoVar "codePointFold_acc") ] ]) ] []
      ] else []
    body = if directAdd then
      [ GoAssign "codePointFold_input" str
      , GoAssign "codePointFold_result" (integer (arg 1) (argType 1))
      ] <> (if bounded then
              [ GoAssign "codePointTake_limit" (integer (arg 2) (argType 2))
              , rawGo "codePointFold_acc, codePointFold_taken := gopurs_runtime.SumCodePointsLimit(codePointFold_input, codePointTake_limit)"
              ]
            else
              [ rawGo "codePointFold_acc, codePointFold_taken := gopurs_runtime.SumCodePoints(codePointFold_input)" ])
        <> [ GoIfElse (GoVar "codePointFold_taken")
               [ GoMutate "codePointFold_result" (GoCall (GoVar "int64") [ GoCall (GoVar "int32") [ GoBinOp "+" (GoVar "codePointFold_result") (GoVar "codePointFold_acc") ] ]) ] []
           ]
      else
        [ GoAssign "codePointFold_input" str
        , GoAssign "codePointFold_result" (integer (arg 1) (argType 1))
        ] <> preparation <> [ mainLoop, tailLoop ] <> finalize
  in { stmts: captured.stmts, expr: GoCall (GoFuncLit [] body (GoVar "codePointFold_result") TypeInt64) []
       , exprType: TypeInt64, nextId: captured.nextId }
  where
  integer expression from = coerceGoExpr context.codegenStateRef context.modNameStr expression from TypeInt64

primitiveCallback :: Tco.TcoExpr -> Maybe { args :: NEA.NonEmptyArray (Tuple (Maybe Ident) Level), body :: Tco.TcoExpr }
primitiveCallback callback = do
  guard (lambdaCompatible [ Int, Int ] callback)
  case unwrapTcoExpr callback of
    Abs args body -> do
      let grouped = collectCurriedAbs args body
      guard (NEA.length grouped.args == 2)
      let levels = map (\(Tuple _ level) -> level) (NEA.toArray grouped.args)
      guard (readOnlyScalar levels grouped.body)
      guard (scalarAnnotations (Map.fromFoldable (map (\level -> Tuple level Int) levels)) grouped.body)
      pure grouped
    _ -> Nothing

-- Each adjacent lambda may carry a TAST suffix. Checking only the outer
-- annotation would miss a contradictory type on the second parameter.
lambdaCompatible :: Array ExprType -> Tco.TcoExpr -> Boolean
lambdaCompatible arguments expression =
  compatibleT (if Array.null arguments then Int else Func arguments Int) expression &&
    if Array.null arguments then true else case unwrapTcoExpr expression of
      Abs parameters body -> NEA.length parameters <= Array.length arguments &&
        lambdaCompatible (Array.drop (NEA.length parameters) arguments) body
      _ -> false

named :: String -> String -> Tco.TcoExpr -> Boolean
named moduleName name expression = case qualifiedTarget expression of
  Just target -> target.mbMod == Just (ModuleName moduleName) && target.name == name
  Nothing -> false

namedN :: String -> String -> NeutralExpr -> Boolean
namedN moduleName name value = case strip value of
  Var (Qualified (Just (ModuleName owner)) (Ident identifier)) -> owner == moduleName && identifier == name
  _ -> false

application :: NeutralExpr -> { head :: NeutralExpr, args :: Array NeutralExpr }
application value = case strip value of
  App fn args -> let prior = application fn in { head: prior.head, args: prior.args <> NEA.toArray args }
  _ -> { head: value, args: [] }

-- Residual generic annotations may wrap the instantiated TAST. Require every
-- wrapper to admit the same specialization, never overwrite a concrete type.
compatible :: ExprType -> NeutralExpr -> Boolean
compatible expected (NeutralExpr syntax) = case syntax of
  Typed actual value -> isJust (refine expected actual Map.empty) && compatible expected value
  TypeApp value _ -> compatible expected value
  _ -> true

compatibleT :: ExprType -> Tco.TcoExpr -> Boolean
compatibleT expected (Tco.TcoExpr _ syntax) = case syntax of
  Typed actual value -> isJust (refine expected actual Map.empty) && compatibleT expected value
  TypeApp value _ -> compatibleT expected value
  _ -> true

refine :: ExprType -> ExprType -> Map.Map String ExprType -> Maybe (Map.Map String ExprType)
refine expected supplied variables = case normalize supplied, expected of
  Any, _ -> Just variables
  TypeVar name, _ -> case Map.lookup name variables of
    Just previous | previous /= expected -> Nothing
    _ -> Just (Map.insert name expected variables)
  Array element, Array wanted -> refine wanted element variables
  ADT "Prim.Array" [ "Prim", "Array" ] [ element ], Array wanted -> refine wanted element variables
  Func arguments result, Func wantedArguments wantedResult -> do
    let actualSignature = arrows arguments result
        wantedSignature = arrows wantedArguments wantedResult
    guard (Array.length actualSignature.args == Array.length wantedSignature.args)
    foldM (\known (Tuple wanted argument) -> refine wanted argument known) variables
      (Array.snoc (Array.zip wantedSignature.args actualSignature.args) (Tuple wantedSignature.result actualSignature.result))
  actual, _ | actual == expected -> Just variables
  _, _ -> Nothing
  where
  arrows args result = case normalize result of
    Func more remaining -> arrows (args <> more) remaining
    normalized -> { args, result: normalized }

-- The TAST also spells instantiated function/array constructors through
-- TypeApp. These are the same concrete types, not permission to erase an
-- incompatible wrapper or to reinterpret a user ADT with a similar name.
normalize :: ExprType -> ExprType
normalize ty = case appliedAdt ty of
  Just { fullName: "Prim.Function", path: [ "Prim", "Function" ], args: [ argument, result ] } -> Func [ argument ] result
  Just { fullName: "Prim.Array", path: [ "Prim", "Array" ], args: [ element ] } -> Array element
  _ -> ty

closed :: Array Level -> NeutralExpr -> Boolean
closed locals expression = case strip expression of
  Local _ level -> Array.elem level locals
  Var _ -> false
  Abs parameters body -> closed (locals <> map (\(Tuple _ level) -> level) (NEA.toArray parameters)) body
  Let _ level value body -> closed locals value && closed (Array.snoc locals level) body
  LetRec _ _ _ -> false
  syntax -> all (closed locals) syntax

strip :: NeutralExpr -> BackendSyntax NeutralExpr
strip (NeutralExpr syntax) = case syntax of
  Typed _ expression -> strip expression
  TypeApp expression _ -> strip expression
  _ -> syntax

-- Cache descriptors are stricter than immediate folds: a replayed worker must
-- neither capture another scope nor move any callback/seed creation effects.
describe :: NeutralExpr -> Maybe CodePointFoldDescriptor
describe original = do
  guard (compatible (Func [ String ] Int) original)
  Tuple level body <- case strip original of
    Abs parameters body | NEA.length parameters == 1 -> Just (Tuple (case NEA.head parameters of Tuple _ level -> level) body)
    _ -> Nothing
  guard (compatible Int body)
  let folding = application body
  guard (namedN "Data.Foldable" "foldlArray" folding.head)
  guard (compatible (Func [ Func [ Int, Int ] Int, Int, Array Int ] Int) folding.head)
  descriptor <- case folding.args of
    [ callback, seed, converted ] -> do
      guard (compatible Int seed && compatible (Array Int) converted)
      initial <- case strip seed of
        Lit (LitInt value) -> Just value
        _ -> Nothing
      let conversion = application converted
      guard (namedN "Data.String.CodePoints" "toCodePointArray" conversion.head)
      guard (compatible (Func [ String ] (Array Int)) conversion.head)
      case conversion.args of
        [ input ] -> do
          guard (compatible String input)
          guard (case strip input of
            Local _ found -> found == level
            _ -> false)
        _ -> Nothing
      pure { callback, seed: initial }
    _ -> Nothing
  guard (validDescriptor descriptor)
  pure descriptor

validDescriptor :: CodePointFoldDescriptor -> Boolean
validDescriptor descriptor = closed [] descriptor.callback && isJust (primitiveCallback (Tco.analyze [] descriptor.callback))

-- Inspect all Typed layers rather than allowing an outer Int to hide an inner
-- contradiction. Captures still use the existing scalar annotation proof for
-- an immediate fold; exported descriptors additionally require closure.
scalarAnnotations :: Map.Map Level ExprType -> Tco.TcoExpr -> Boolean
scalarAnnotations locals value@(Tco.TcoExpr _ syntax) = isJust (scalarType locals value) && case syntax of
  Let _ level initial body -> case scalarType locals initial of
    Just ty -> scalarAnnotations locals initial && scalarAnnotations (Map.insert level ty locals) body
    Nothing -> false
  _ -> all (scalarAnnotations locals) syntax
