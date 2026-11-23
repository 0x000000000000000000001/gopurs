module Gopurs.RegexMatchProjection (optimize, rewrite, helperName, helperType, adapterName) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (all)
import Data.Maybe (Maybe(..))
import Data.String as String
import Data.String.CodeUnits as CU
import Data.String.Pattern (Pattern(..))
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Gopurs.Cache.LookupTable as LookupTable
import Gopurs.CodegenState (CodegenMetadata)
import PureScript.Backend.Optimizer.Convert (BackendModule)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Ident(..), ModuleName(..), Qualified(..))
import PureScript.Backend.Optimizer.CoreFn as CoreFn
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendAccessor(..), BackendOperator(..), BackendOperator1(..), BackendSyntax(..), Level, Pair(..))
import PureScript.Backend.Optimizer.Syntax as Syn

helperName :: String
helperName = "Data.String.Regex.matchWholeImpl"

adapterName :: String
adapterName = "Data.Array.NonEmpty.adaptMaybe"

resultType :: ExprType
resultType = Record (Row [ Tuple "matched" Boolean, Tuple "text" String ] Nothing)

matchedStringType :: ExprType
matchedStringType = ADT "Data.Maybe.Maybe" [ "Data", "Maybe", "Maybe" ] [ String ]

helperType :: ExprType
helperType = Func [ ADT "Data.String.Regex.Regex" [ "Data", "String", "Regex", "Regex" ] [], String ] resultType

-- The port explicitly opts into this contract. Older ports retain Regex.match.
-- Source references introduced here are included in the usual worker/cache
-- dependency collection because this pass runs before prepareModule finishes.
optimize :: CodegenMetadata -> BackendModule -> BackendModule
optimize metadata mod
  | map normalize (LookupTable.lookup helperName metadata.globalTypes) /= Just helperType = mod
  | otherwise = mod { bindings = map (\group -> group { bindings = map (\(Tuple name value) -> Tuple name (visit value)) group.bindings }) mod.bindings }
  where
  visit original@(NeutralExpr syntax) = case rewriteWith (standardAdapter metadata) original of
    Just rewritten -> rewritten
    Nothing -> NeutralExpr (map visit syntax)

-- Admit the complete match/Just/head/Just consumer, then replace only the
-- observed whole-match value. Unknown uses, aliases, escapes, extra groups,
-- different constructors and unguarded payload accesses retain the original.
rewrite :: NeutralExpr -> Maybe NeutralExpr
rewrite = rewriteWith (named "Data.Array.NonEmpty" "adaptMaybe")

rewriteWith :: (NeutralExpr -> Boolean) -> NeutralExpr -> Maybe NeutralExpr
rewriteWith adapter original = do
  binding <- case strip original of
    Let ident level value body -> Just { ident, level, value, body }
    _ -> Nothing
  Tuple fn args <- application binding.value
  guard (named "Data.String.Regex" "match" fn && Array.length args == 2)
  arguments <- NEA.fromArray args
  outer <- justBranch binding.level binding.body
  guard (independent [ binding.level ] outer.failure)
  first <- case strip outer.success of
    Let ident level value body -> Just { ident, level, value, body }
    _ -> Nothing
  guard (first.level /= binding.level && wholeHead adapter binding.level first.value)
  inner <- justBranch first.level first.body
  guard (independent [ binding.level, first.level ] inner.failure)
  -- The only successful payload use is Just.value0. Reject references to the
  -- original array anywhere, including in callbacks and delayed computations.
  guard (independent [ binding.level ] inner.success)
  let result = typed resultType (NeutralExpr (Local binding.ident binding.level))
      text = typed String (NeutralExpr (Accessor result (GetProp "text")))
      matched = typed Boolean (NeutralExpr (Accessor result (GetProp "matched")))
  success <- replacePayload first.level text inner.success
  let call = typed resultType (NeutralExpr (App
        (typed helperType (NeutralExpr (Var (Qualified (Just (ModuleName "Data.String.Regex")) (Ident "matchWholeImpl")))))
        arguments))
      innerBody = replace first.body (NeutralExpr (Branch (NEA.singleton (Pair matched success)) inner.failure))
      body = replace binding.body (NeutralExpr (Branch (NEA.singleton (Pair matched innerBody)) outer.failure))
  pure (replace original (NeutralExpr (Let binding.ident binding.level call body)))

justBranch :: Level -> NeutralExpr -> Maybe { success :: NeutralExpr, failure :: NeutralExpr }
justBranch level expression = case strip expression of
  Branch branches failure -> case NEA.toArray branches of
    [ Pair condition success ] | justTest level condition -> Just { success, failure }
    _ -> Nothing
  _ -> Nothing

justTest :: Level -> NeutralExpr -> Boolean
justTest level expression = case strip expression of
  PrimOp (Op1 (OpIsTag ctor) value) -> ctor == just && local level value
  _ -> false

wholeHead :: (NeutralExpr -> Boolean) -> Level -> NeutralExpr -> Boolean
wholeHead adapter level expression = case application expression of
  Just (Tuple fn [ value ]) -> named "Data.Array.NonEmpty" "head" fn && payload level value
  Just (Tuple fn [ head, value ]) -> adapter fn
    && named "Data.Array" "head" head && payload level value
    -- PBO retains the result annotation on this application but can introduce
    -- an unannotated Var for its monomorphized callee.
    && (named "Data.Array.NonEmpty" "adaptMaybe" fn
      || map normalize (annotation expression) == Just matchedStringType)
  _ -> false

payload :: Level -> NeutralExpr -> Boolean
payload level expression = case strip expression of
  Accessor value (GetCtorField ctor _ _ (Ident "Just") _ 0) -> ctor == just && local level value
  _ -> false

replacePayload :: Level -> NeutralExpr -> NeutralExpr -> Maybe NeutralExpr
replacePayload level value expression@(NeutralExpr syntax)
  | payload level expression = Just value
  | otherwise = case strip expression of
      Local _ found | found == level -> Nothing
      _ -> NeutralExpr <$> traverse (replacePayload level value) syntax

independent :: Array Level -> NeutralExpr -> Boolean
independent levels (NeutralExpr syntax) = case syntax of
  Local _ level -> not (Array.elem level levels)
  _ -> all (independent levels) syntax

local :: Level -> NeutralExpr -> Boolean
local level expression = case strip expression of
  Local _ found -> level == found
  _ -> false

just :: Qualified Ident
just = Qualified (Just (ModuleName "Data.Maybe")) (Ident "Just")

named :: String -> String -> NeutralExpr -> Boolean
named owner name expression = case strip expression of
  Var found -> found == Qualified (Just (ModuleName owner)) (Ident name)
  _ -> false

application :: NeutralExpr -> Maybe (Tuple NeutralExpr (Array NeutralExpr))
application expression = case strip expression of
  App fn args -> case application fn of
    Just (Tuple head prior) -> Just (Tuple head (prior <> NEA.toArray args))
    Nothing -> Just (Tuple fn (NEA.toArray args))
  _ -> Nothing

typed :: ExprType -> NeutralExpr -> NeutralExpr
typed ty = NeutralExpr <<< Typed ty

strip :: NeutralExpr -> BackendSyntax NeutralExpr
strip (NeutralExpr syntax) = case syntax of
  Typed _ value -> strip value
  Syn.TypeApp value _ -> strip value
  _ -> syntax

replace :: NeutralExpr -> NeutralExpr -> NeutralExpr
replace (NeutralExpr syntax) replacement = case syntax of
  Typed ty value -> typed ty (replace value replacement)
  Syn.TypeApp value ty -> NeutralExpr (Syn.TypeApp (replace value replacement) ty)
  _ -> replacement

normalize :: ExprType -> ExprType
normalize (ForAll _ ty) = normalize ty
normalize (Func args ret) = case normalize ret of
  Func more result -> Func (map normalize args <> more) result
  result -> Func (map normalize args) result
normalize ty = ty

-- Monomorphization appends a decimal hash to the original name. globalTypes
-- indexes original TAST declarations, before that pass: a source declaration
-- bearing the same suffix must not acquire the adapter's semantics. Require
-- the precise instantiated adapter type as well as the original declaration.
standardAdapter :: CodegenMetadata -> NeutralExpr -> Boolean
standardAdapter metadata expression
  | named "Data.Array.NonEmpty" "adaptMaybe" expression = true
  | otherwise = case strip expression of
      Var (Qualified (Just (ModuleName "Data.Array.NonEmpty")) (Ident name)) ->
        case String.stripPrefix (Pattern "adaptMaybe__") name of
          Just suffix -> suffix /= ""
            && Array.all (\c -> c >= '0' && c <= '9') (CU.toCharArray suffix)
            && (case map normalize (LookupTable.lookup adapterName metadata.globalTypes) of
                  Just Any -> false
                  Just _ -> true
                  Nothing -> false)
            && not (LookupTable.member ("Data.Array.NonEmpty." <> name) metadata.globalTypes)
            && (case map normalize (annotation expression) of
                  Nothing -> true
                  Just ty -> ty == adapterType)
          Nothing -> false
      _ -> false
  where
  maybeString = matchedStringType
  maybeMaybeString = ADT "Data.Maybe.Maybe" [ "Data", "Maybe", "Maybe" ] [ maybeString ]
  values = CoreFn.Array maybeString
  adapterType = Func [ Func [ values ] maybeMaybeString, values ] maybeString

annotation :: NeutralExpr -> Maybe ExprType
annotation (NeutralExpr syntax) = case syntax of
  Typed ty _ -> Just ty
  Syn.TypeApp value _ -> annotation value
  _ -> Nothing
