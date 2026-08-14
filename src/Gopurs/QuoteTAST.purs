module Gopurs.QuoteTAST where

import Prelude
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Array as Array
import Data.Array.NonEmpty as NonEmptyArray
import Data.Tuple (Tuple(..))
import Partial.Unsafe (unsafeCrashWith)

import PureScript.Backend.Optimizer.CoreFn (Ann(..), Bind(..), Binding(..), CaseAlternative(..), CaseGuard(..), Expr(..), ExprType(..), Guard(..), Ident(..), Prop(..), Qualified(..))
import PureScript.Backend.Optimizer.Semantics (BackendExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), Level(..), BackendAccessor(..))

quoteTAST :: BackendExpr -> Expr Ann
quoteTAST = case _ of
  ExprSyntax _ syntax -> case syntax of
    Typed ty inner ->
      let ann = Ann { span: { path: "", start: { line: 0, column: 0 }, end: { line: 0, column: 0 } }, meta: Nothing, type: Just ty }
      in case inner of
        ExprSyntax _ innerSyntax -> case innerSyntax of
          Var qual -> ExprVar ann qual

          Local ident _ -> ExprVar ann (Qualified Nothing (fromMaybe (Ident "_") ident))
          Lit lit -> ExprLit ann (map quoteTAST lit)
          App f args -> Array.foldl (ExprApp ann) (quoteTAST f) (map quoteTAST (NonEmptyArray.toArray args))
          UncurriedApp f args -> Array.foldl (ExprApp ann) (quoteTAST f) (map quoteTAST args)
          Abs args body -> Array.foldr (\(Tuple ident _) b -> ExprAbs ann (fromMaybe (Ident "_") ident) b) (quoteTAST body) (NonEmptyArray.toArray args)
          UncurriedAbs args body -> Array.foldr (\(Tuple ident _) b -> ExprAbs ann (fromMaybe (Ident "_") ident) b) (quoteTAST body) args
          UncurriedEffectApp f args -> Array.foldl (ExprApp ann) (quoteTAST f) (map quoteTAST args)
          UncurriedEffectAbs args body -> Array.foldr (\(Tuple ident _) b -> ExprAbs ann (fromMaybe (Ident "_") ident) b) (quoteTAST body) args
          Let ident _ val body -> ExprLet ann [ NonRec (Binding ann (fromMaybe (Ident "_") ident) (quoteTAST val)) ] (quoteTAST body)
          LetRec _ binds body -> ExprLet ann [ Rec (map (\(Tuple ident val) -> Binding ann ident (quoteTAST val)) (NonEmptyArray.toArray binds)) ] (quoteTAST body)
          Accessor e (GetProp prop) -> ExprAccessor ann (quoteTAST e) prop
          Accessor e (GetIndex idx) -> ExprAccessor ann (quoteTAST e) (show idx)
          Accessor e (GetCtorField _ _ _ _ prop _) -> ExprAccessor ann (quoteTAST e) prop
          Update e props -> ExprUpdate ann (quoteTAST e) (map (\(Prop p v) -> Prop p (quoteTAST v)) props)
          CtorSaturated _ _ pn ident fields -> 
            Array.foldl (ExprApp ann) (ExprConstructor ann pn ident (map (\(Tuple k _) -> k) fields)) (map (\(Tuple _ v) -> quoteTAST v) fields)
          CtorDef _ pn ident fields -> ExprConstructor ann pn ident fields
          EffectBind ident _ val body -> ExprLet ann [ NonRec (Binding ann (fromMaybe (Ident "_") ident) (quoteTAST val)) ] (quoteTAST body)
          EffectPure e -> quoteTAST e
          EffectDefer e -> quoteTAST e
          Branch _ _ -> 
            -- Note: Branch translation is non-trivial if it comes from Case. 
            -- But we don't have pattern matching natively in Branch, it's just booleans or simple conditions.
            -- For now we just crash, since we need to see how to reconstruct ExprCase.
            unsafeCrashWith "quoteTAST: Branch translation is not yet fully implemented"
          PrimOp _ -> unsafeCrashWith "quoteTAST: PrimOp not yet supported"
          PrimEffect _ -> unsafeCrashWith "quoteTAST: PrimEffect not yet supported"
          PrimUndefined -> unsafeCrashWith "quoteTAST: PrimUndefined not yet supported"
          Fail err -> unsafeCrashWith ("quoteTAST: Fail - " <> err)
          Typed _ _ -> unsafeCrashWith "quoteTAST: nested Typed"
        ExprRewrite _ _ -> unsafeCrashWith "quoteTAST: expected syntax inside Typed, found rewrite"
    _ -> unsafeCrashWith "quoteTAST: Expected Typed syntax"
  ExprRewrite _ _ ->
    unsafeCrashWith "quoteTAST: ExprRewrite encountered"
