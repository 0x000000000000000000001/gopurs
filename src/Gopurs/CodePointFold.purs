module Gopurs.CodePointFold (optimize, emit) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (all, foldM)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe, isJust)
import Data.Tuple (Tuple(..))
import Gopurs.ArrayIndexedFold (readOnlyScalar)
import Gopurs.CallAnalysis (collectCurriedAbs, collectGoSpine, getGoSpineArgs, qualifiedTarget)
import Gopurs.CallArguments as CallArguments
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
emit translate context nextId fn args = do
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
    captured = CallArguments.capture { prefix: "codePointFold_arg_", markUsed: true } capture nextId [ callback, seed, input ]
    arg index = fromMaybe (rawGo "nil") (Array.index captured.exprs index)
    argType index = fromMaybe TypeValue (Array.index captured.exprTypes index)
    str = coerceGoExpr context.codegenStateRef context.modNameStr (arg 2) (argType 2) TypeString
    body =
      [ GoAssign "codePointFold_input" str
      , GoAssign "codePointFold_result" (integer (arg 1) (argType 1))
      , GoForRange "codePointFold_offset := 0; codePointFold_offset < len(codePointFold_input);"
          [ GoAssign "codePointFold_value" (GoCall (GoVar "rune")
              [ GoIndex (GoVar "codePointFold_input") (GoVar "codePointFold_offset") ])
          , GoAssign "codePointFold_width" (GoInt 1)
          -- ASCII is exactly one code point of width one. Keep the more
          -- complex decoder off this common path; every other byte retains
          -- the complete UTF-8/WTF-8 and replacement-character semantics.
          , GoIfElse (GoBinOp ">=" (GoVar "codePointFold_value") (GoInt 128))
              [ GoMutate "codePointFold_value, codePointFold_width"
                  (GoCall (GoSelector (GoVar "gopurs_runtime") "DecodeStringCodePoint") [ rawGo "codePointFold_input[codePointFold_offset:]" ]) ] []
          , GoMutate "codePointFold_result" (GoCall (arg 0)
              [ GoVar "codePointFold_result", GoCall (GoVar "int64") [ GoVar "codePointFold_value" ] ])
          , GoMutate "codePointFold_offset" (GoBinOp "+" (GoVar "codePointFold_offset") (GoVar "codePointFold_width"))
          ]
      ]
  pure { stmts: captured.stmts, expr: GoCall (GoFuncLit [] body (GoVar "codePointFold_result") TypeInt64) []
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
      guard (readOnlyScalar (map (\(Tuple _ level) -> level) (NEA.toArray grouped.args)) grouped.body)
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
