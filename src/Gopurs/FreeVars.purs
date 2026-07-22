module Gopurs.FreeVars where

import Prelude

import Data.Maybe (Maybe(..))
import Data.Set (Set)
import Data.Set as Set
import Data.Tuple (Tuple(..))
import Data.Array (foldl)
import Data.String (Pattern(..), Replacement(..))
import Data.String as String

import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), Level(..), Pair(..), BackendOperator(..))
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (Ident(..), Prop(..), Literal(..))
import Data.Array.NonEmpty (toArray)

sanitizeName :: String -> String
sanitizeName name = 
  let s1 = String.replaceAll (Pattern "'") (Replacement "_prime") (String.replaceAll (Pattern "$") (Replacement "_dollar") name)
  in if s1 == "break" || s1 == "default" || s1 == "func" || s1 == "interface" || s1 == "select" || s1 == "case" || s1 == "defer" || s1 == "go" || s1 == "map" || s1 == "struct" || s1 == "chan" || s1 == "else" || s1 == "goto" || s1 == "package" || s1 == "switch" || s1 == "const" || s1 == "fallthrough" || s1 == "if" || s1 == "range" || s1 == "type" || s1 == "continue" || s1 == "for" || s1 == "import" || s1 == "return" || s1 == "var" then s1 <> "_" else s1

localId :: Maybe Ident -> Level -> String
localId (Just (Ident i)) (Level l) = sanitizeName i <> "_" <> show l
localId Nothing (Level l) = "__local_var_" <> show l

-- | Computes the set of free variables (by unique localId) in a given `TcoExpr`.
freeVars :: TcoExpr -> Set String
freeVars (TcoExpr _ syntax) = case syntax of
  Var _ -> Set.empty
  Local mbIdent lvl -> Set.singleton (localId mbIdent lvl)
  Lit lit -> case lit of
    LitArray arr -> foldl (\acc e -> Set.union acc (freeVars e)) Set.empty arr
    LitRecord rec -> foldl (\acc (Prop _ e) -> Set.union acc (freeVars e)) Set.empty rec
    _ -> Set.empty
  App fn args -> 
    foldl (\acc e -> Set.union acc (freeVars e)) (freeVars fn) (toArray args)
  Abs args body ->
    let 
      argsSet = foldl (\acc (Tuple mbIdent lvl) -> Set.insert (localId mbIdent lvl) acc) Set.empty (toArray args)
      bodyVars = freeVars body
    in Set.difference bodyVars argsSet
  UncurriedApp fn args ->
    foldl (\acc e -> Set.union acc (freeVars e)) (freeVars fn) args
  UncurriedAbs args body ->
    let 
      argsSet = foldl (\acc (Tuple mbIdent lvl) -> Set.insert (localId mbIdent lvl) acc) Set.empty args
      bodyVars = freeVars body
    in Set.difference bodyVars argsSet
  UncurriedEffectApp fn args ->
    foldl (\acc e -> Set.union acc (freeVars e)) (freeVars fn) args
  UncurriedEffectAbs args body ->
    let 
      argsSet = foldl (\acc (Tuple mbIdent lvl) -> Set.insert (localId mbIdent lvl) acc) Set.empty args
      bodyVars = freeVars body
    in Set.difference bodyVars argsSet
  Accessor e _ -> freeVars e
  Update e props ->
    foldl (\acc (Prop _ val) -> Set.union acc (freeVars val)) (freeVars e) props
  CtorSaturated _ _ _ _ args ->
    foldl (\acc (Tuple _ e) -> Set.union acc (freeVars e)) Set.empty args
  CtorDef _ _ _ _ -> Set.empty
  LetRec lvl binds body ->
    let
      bindsSet = foldl (\acc (Tuple ident _) -> Set.insert (localId (Just ident) lvl) acc) Set.empty (toArray binds)
      bodyVars = freeVars body
      bindsVars = foldl (\acc (Tuple _ e) -> Set.union acc (freeVars e)) Set.empty (toArray binds)
    in Set.difference (Set.union bodyVars bindsVars) bindsSet
  Let mbIdent lvl val body ->
    Set.union (freeVars val) (Set.difference (freeVars body) (Set.singleton (localId mbIdent lvl)))
  EffectBind mbIdent lvl val body ->
    Set.union (freeVars val) (Set.difference (freeVars body) (Set.singleton (localId mbIdent lvl)))
  EffectPure e -> freeVars e
  EffectDefer e -> freeVars e
  Branch pairs def ->
    let
      defVars = freeVars def
      pairsVars = foldl (\acc (Pair cond body) -> Set.union acc (Set.union (freeVars cond) (freeVars body))) Set.empty (toArray pairs)
    in Set.union defVars pairsVars
  PrimOp op -> case op of
    Op1 _ e -> freeVars e
    Op2 _ e1 e2 -> Set.union (freeVars e1) (freeVars e2)
  PrimEffect _ -> Set.empty
  PrimUndefined -> Set.empty
  Fail _ -> Set.empty
