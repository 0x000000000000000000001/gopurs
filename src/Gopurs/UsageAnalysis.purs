module Gopurs.UsageAnalysis where

import Prelude

import Data.Array as Array
import Data.Map (Map)
import Data.Map as Map
import Data.Set (Set)
import Data.Set as Set
import Data.Maybe (Maybe(..))
import Data.Tuple (Tuple(..))
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (Ident(..), Prop(..), Literal(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(Var, Local, Lit, App, Abs, UncurriedApp, UncurriedAbs, UncurriedEffectApp, UncurriedEffectAbs, Accessor, Update, CtorSaturated, CtorDef, LetRec, Let, EffectBind, EffectPure, EffectDefer, Branch, PrimOp, PrimEffect, PrimUndefined, Fail, Typed), Level(..), Pair(..), BackendOperator(..))
import PureScript.Backend.Optimizer.Syntax as Syn
import Data.Array.NonEmpty (toArray)
import Data.String (Pattern(..), Replacement(..))
import Data.String as String

sanitizeName :: String -> String
sanitizeName name = 
  let
    s1 = String.replaceAll (Pattern "/") (Replacement "_slash_")
      $ String.replaceAll (Pattern "\\") (Replacement "_bslash_")
      $ String.replaceAll (Pattern "<") (Replacement "_less_")
      $ String.replaceAll (Pattern ">") (Replacement "_greater_")
      $ String.replaceAll (Pattern "=") (Replacement "_eq_")
      $ String.replaceAll (Pattern "+") (Replacement "_plus_")
      $ String.replaceAll (Pattern "-") (Replacement "_minus_")
      $ String.replaceAll (Pattern "*") (Replacement "_times_")
      $ String.replaceAll (Pattern ":") (Replacement "_colon_")
      $ String.replaceAll (Pattern "|") (Replacement "_bar_")
      $ String.replaceAll (Pattern "&") (Replacement "_amp_")
      $ String.replaceAll (Pattern "^") (Replacement "_caret_")
      $ String.replaceAll (Pattern "~") (Replacement "_tilde_")
      $ String.replaceAll (Pattern "?") (Replacement "_qmark_")
      $ String.replaceAll (Pattern "!") (Replacement "_bang_")
      $ String.replaceAll (Pattern "@") (Replacement "_at_")
      $ String.replaceAll (Pattern "#") (Replacement "_hash_")
      $ String.replaceAll (Pattern "%") (Replacement "_percent_")
      $ String.replaceAll (Pattern "\"") (Replacement "_quote_")
      $ String.replaceAll (Pattern ".") (Replacement "_dot_")
      $ String.replaceAll (Pattern "'") (Replacement "_prime_")
      $ String.replaceAll (Pattern "$") (Replacement "_dollar_") name
  in if s1 == "break" || s1 == "default" || s1 == "func" || s1 == "interface" || s1 == "select" || s1 == "case" || s1 == "defer" || s1 == "go" || s1 == "map" || s1 == "struct" || s1 == "chan" || s1 == "else" || s1 == "goto" || s1 == "package" || s1 == "switch" || s1 == "const" || s1 == "fallthrough" || s1 == "if" || s1 == "range" || s1 == "type" || s1 == "continue" || s1 == "for" || s1 == "import" || s1 == "return" || s1 == "var" then "go__" <> s1 else s1

localId :: Maybe Ident -> Level -> String
localId (Just (Ident i)) (Level l) = sanitizeName i <> "_" <> show l
localId Nothing (Level l) = "__local_var_" <> show l

addUsages :: Map String Int -> Map String Int -> Map String Int
addUsages = Map.unionWith (+)

maxUsages :: Map String Int -> Map String Int -> Map String Int
maxUsages = Map.unionWith max

-- | Computes the usage count of each free variable in a given `TcoExpr`.
usageCount :: TcoExpr -> Map String Int
usageCount (TcoExpr _ syntax) = case syntax of
  Var _ -> Map.empty
  Local mbIdent lvl -> Map.singleton (localId mbIdent lvl) 1
  Lit lit -> case lit of
    LitArray arr -> Array.foldl (\acc e -> addUsages acc (usageCount e)) Map.empty arr
    LitRecord rec -> Array.foldl (\acc (Prop _ e) -> addUsages acc (usageCount e)) Map.empty rec
    _ -> Map.empty
  App fn args -> 
    Array.foldl (\acc e -> addUsages acc (usageCount e)) (usageCount fn) (toArray args)
  Syn.TypeApp fn _ ->
    usageCount fn
  Abs args body ->
    let 
      argsList = map (\(Tuple mbIdent lvl) -> localId mbIdent lvl) (toArray args)
      bodyVars = usageCount body
    in Array.foldl (flip Map.delete) bodyVars argsList
  UncurriedApp fn args ->
    Array.foldl (\acc e -> addUsages acc (usageCount e)) (usageCount fn) args
  UncurriedAbs args body ->
    let 
      argsList = map (\(Tuple mbIdent lvl) -> localId mbIdent lvl) args
      bodyVars = usageCount body
    in Array.foldl (flip Map.delete) bodyVars argsList
  UncurriedEffectApp fn args ->
    Array.foldl (\acc e -> addUsages acc (usageCount e)) (usageCount fn) args
  UncurriedEffectAbs args body ->
    let 
      argsList = map (\(Tuple mbIdent lvl) -> localId mbIdent lvl) args
      bodyVars = usageCount body
    in Array.foldl (flip Map.delete) bodyVars argsList
  Accessor e _ -> usageCount e
  Update e props ->
    Array.foldl (\acc (Prop _ val) -> addUsages acc (usageCount val)) (usageCount e) props
  CtorSaturated _ _ _ _ args ->
    Array.foldl (\acc (Tuple _ e) -> addUsages acc (usageCount e)) Map.empty args
  CtorDef _ _ _ _ -> Map.empty
  LetRec lvl binds body ->
    let
      bindsList = map (\(Tuple ident _) -> localId (Just ident) lvl) (toArray binds)
      bodyVars = usageCount body
      bindsVars = Array.foldl (\acc (Tuple _ e) -> addUsages acc (usageCount e)) Map.empty (toArray binds)
    in Array.foldl (flip Map.delete) (addUsages bodyVars bindsVars) bindsList
  Let mbIdent lvl val body ->
    addUsages (usageCount val) (Map.delete (localId mbIdent lvl) (usageCount body))
  EffectBind mbIdent lvl val body ->
    addUsages (usageCount val) (Map.delete (localId mbIdent lvl) (usageCount body))
  EffectPure e -> usageCount e
  EffectDefer e -> usageCount e
  Branch pairs def ->
    let
      defVars = usageCount def
      pairsVars = Array.foldl (\acc (Pair cond body) -> addUsages acc (maxUsages (usageCount cond) (usageCount body))) Map.empty (toArray pairs)
      -- Wait, for `if cond then true else false`, it should be `usage(cond) + max(usage(true), usage(false))`.
      -- But Branch has multiple pairs. It's like `cond1 ? body1 : cond2 ? body2 : def`.
      -- So usages = cond1 + max(body1, cond2 + max(body2, def))
    in foldBranchPairs (toArray pairs) (usageCount def)
  PrimOp op -> case op of
    Op1 _ e -> usageCount e
    Op2 _ e1 e2 -> addUsages (usageCount e1) (usageCount e2)
  PrimEffect _ -> Map.empty
  PrimUndefined -> Map.empty
  Fail _ -> Map.empty
  Typed _ a -> usageCount a

foldBranchPairs :: Array (Pair TcoExpr) -> Map String Int -> Map String Int
foldBranchPairs pairs def = case Array.uncons pairs of
  Nothing -> def
  Just { head: Pair cond body, tail } ->
    let next = foldBranchPairs tail def
    in addUsages (usageCount cond) (maxUsages (usageCount body) next)

freeVars :: TcoExpr -> Set String
freeVars (TcoExpr _ syntax) = case syntax of
  Var _ -> Set.empty
  Local mbIdent lvl -> Set.singleton (localId mbIdent lvl)
  Lit lit -> case lit of
    LitArray arr -> Array.foldl (\acc e -> Set.union acc (freeVars e)) Set.empty arr
    LitRecord rec -> Array.foldl (\acc (Prop _ e) -> Set.union acc (freeVars e)) Set.empty rec
    _ -> Set.empty
  App fn args ->
    Array.foldl (\acc e -> Set.union acc (freeVars e)) (freeVars fn) (toArray args)
  Syn.TypeApp fn _ ->
    freeVars fn
  Abs args body ->
    let 
      argsList = map (\(Tuple mbIdent lvl) -> localId mbIdent lvl) (toArray args)
      bodyVars = freeVars body
    in Array.foldl (flip Set.delete) bodyVars argsList
  UncurriedApp fn args ->
    Array.foldl (\acc e -> Set.union acc (freeVars e)) (freeVars fn) args
  UncurriedAbs args body ->
    let 
      argsList = map (\(Tuple mbIdent lvl) -> localId mbIdent lvl) args
      bodyVars = freeVars body
    in Array.foldl (flip Set.delete) bodyVars argsList
  UncurriedEffectApp fn args ->
    Array.foldl (\acc e -> Set.union acc (freeVars e)) (freeVars fn) args
  UncurriedEffectAbs args body ->
    let 
      argsList = map (\(Tuple mbIdent lvl) -> localId mbIdent lvl) args
      bodyVars = freeVars body
    in Array.foldl (flip Set.delete) bodyVars argsList
  Accessor e _ -> freeVars e
  Update e props ->
    Array.foldl (\acc (Prop _ val) -> Set.union acc (freeVars val)) (freeVars e) props
  CtorSaturated _ _ _ _ args ->
    Array.foldl (\acc (Tuple _ e) -> Set.union acc (freeVars e)) Set.empty args
  CtorDef _ _ _ _ -> Set.empty
  LetRec lvl binds body ->
    let
      bindsList = map (\(Tuple ident _) -> localId (Just ident) lvl) (toArray binds)
      bodyVars = freeVars body
      bindsVars = Array.foldl (\acc (Tuple _ e) -> Set.union acc (freeVars e)) Set.empty (toArray binds)
    in Array.foldl (flip Set.delete) (Set.union bodyVars bindsVars) bindsList
  Let mbIdent lvl val body ->
    Set.union (freeVars val) (Set.delete (localId mbIdent lvl) (freeVars body))
  EffectBind mbIdent lvl val body ->
    Set.union (freeVars val) (Set.delete (localId mbIdent lvl) (freeVars body))
  EffectPure e -> freeVars e
  EffectDefer e -> freeVars e
  Branch pairs def ->
    let
      pairsVars = Array.foldl (\acc (Pair cond body) -> Set.union acc (Set.union (freeVars cond) (freeVars body))) Set.empty (toArray pairs)
    in Set.union pairsVars (freeVars def)
  PrimOp op -> case op of
    Op1 _ e -> freeVars e
    Op2 _ e1 e2 -> Set.union (freeVars e1) (freeVars e2)
  PrimEffect _ -> Set.empty
  PrimUndefined -> Set.empty
  Fail _ -> Set.empty
  Typed _ a -> freeVars a

