module Gopurs.QuoteTAST where

import Prelude
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Array as Array
import Data.Array.NonEmpty as NonEmptyArray
import Data.Tuple (Tuple(..))
import Partial.Unsafe (unsafeCrashWith)

import PureScript.Backend.Optimizer.CoreFn (Ann(..), Bind(..), Binding(..), CaseAlternative(..), CaseGuard(..), Expr(..), ExprType(..), Guard(..), Ident(..), Prop(..), Qualified(..), Binder(..), Literal(..))
import PureScript.Backend.Optimizer.Semantics (BackendExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), Level(..), BackendAccessor(..), fstPair, sndPair)

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
          Branch pairs def -> 
            Array.foldr (\pair acc ->
              let cond = quoteTAST (fstPair pair)
                  res = quoteTAST (sndPair pair)
              in ExprCase ann [cond]
                   [ CaseAlternative [BinderLit ann (LitBoolean true)] (Unconditional res)
                   , CaseAlternative [BinderLit ann (LitBoolean false)] (Unconditional acc)
                   ]
            ) (quoteTAST def) (NonEmptyArray.toArray pairs)
          PrimOp _ -> unsafeCrashWith "quoteTAST: PrimOp not yet supported"
          PrimEffect _ -> unsafeCrashWith "quoteTAST: PrimEffect not yet supported"
          PrimUndefined -> unsafeCrashWith "quoteTAST: PrimUndefined not yet supported"
          Fail err -> unsafeCrashWith ("quoteTAST: Fail - " <> err)
          Typed _ e -> quoteTAST e
        ExprRewrite _ _ -> unsafeCrashWith "quoteTAST: expected syntax inside Typed, found rewrite"
    -- If it's not a Typed at the top level, we just recurse if possible or wrap it
    -- but usually Semantics wraps everything in Typed. If not, we lack the type.
    _ -> 
      let dummyAnn = Ann { span: { path: "", start: { line: 0, column: 0 }, end: { line: 0, column: 0 } }, meta: Nothing, type: Nothing }
      in case syntax of
        Var qual -> ExprVar dummyAnn qual
        Local ident _ -> ExprVar dummyAnn (Qualified Nothing (fromMaybe (Ident "_") ident))
        Lit lit -> ExprLit dummyAnn (map quoteTAST lit)
        App f args -> Array.foldl (ExprApp dummyAnn) (quoteTAST f) (map quoteTAST (NonEmptyArray.toArray args))
        UncurriedApp f args -> Array.foldl (ExprApp dummyAnn) (quoteTAST f) (map quoteTAST args)
        Abs args body -> Array.foldr (\(Tuple ident _) b -> ExprAbs dummyAnn (fromMaybe (Ident "_") ident) b) (quoteTAST body) (NonEmptyArray.toArray args)
        UncurriedAbs args body -> Array.foldr (\(Tuple ident _) b -> ExprAbs dummyAnn (fromMaybe (Ident "_") ident) b) (quoteTAST body) args
        UncurriedEffectApp f args -> Array.foldl (ExprApp dummyAnn) (quoteTAST f) (map quoteTAST args)
        UncurriedEffectAbs args body -> Array.foldr (\(Tuple ident _) b -> ExprAbs dummyAnn (fromMaybe (Ident "_") ident) b) (quoteTAST body) args
        Let ident _ val body -> ExprLet dummyAnn [ NonRec (Binding dummyAnn (fromMaybe (Ident "_") ident) (quoteTAST val)) ] (quoteTAST body)
        LetRec _ binds body -> ExprLet dummyAnn [ Rec (map (\(Tuple ident val) -> Binding dummyAnn ident (quoteTAST val)) (NonEmptyArray.toArray binds)) ] (quoteTAST body)
        Accessor e (GetProp prop) -> ExprAccessor dummyAnn (quoteTAST e) prop
        Accessor e (GetIndex idx) -> ExprAccessor dummyAnn (quoteTAST e) (show idx)
        Accessor e (GetCtorField _ _ _ _ prop _) -> ExprAccessor dummyAnn (quoteTAST e) prop
        Update e props -> ExprUpdate dummyAnn (quoteTAST e) (map (\(Prop p v) -> Prop p (quoteTAST v)) props)
        CtorSaturated _ _ pn ident fields -> 
          Array.foldl (ExprApp dummyAnn) (ExprConstructor dummyAnn pn ident (map (\(Tuple k _) -> k) fields)) (map (\(Tuple _ v) -> quoteTAST v) fields)
        CtorDef _ pn ident fields -> ExprConstructor dummyAnn pn ident fields
        EffectBind ident _ val body -> ExprLet dummyAnn [ NonRec (Binding dummyAnn (fromMaybe (Ident "_") ident) (quoteTAST val)) ] (quoteTAST body)
        EffectPure e -> quoteTAST e
        EffectDefer e -> quoteTAST e
        Branch pairs def -> 
          Array.foldr (\pair acc ->
            let cond = quoteTAST (fstPair pair)
                res = quoteTAST (sndPair pair)
            in ExprCase dummyAnn [cond]
                 [ CaseAlternative [BinderLit dummyAnn (LitBoolean true)] (Unconditional res)
                 , CaseAlternative [BinderLit dummyAnn (LitBoolean false)] (Unconditional acc)
                 ]
          ) (quoteTAST def) (NonEmptyArray.toArray pairs)
        PrimOp _ -> unsafeCrashWith "quoteTAST: PrimOp not yet supported"
        PrimEffect _ -> unsafeCrashWith "quoteTAST: PrimEffect not yet supported"
        PrimUndefined -> unsafeCrashWith "quoteTAST: PrimUndefined not yet supported"
        Fail err -> unsafeCrashWith ("quoteTAST: Fail - " <> err)
        Typed _ e -> quoteTAST e
  ExprRewrite _ _ ->
    unsafeCrashWith "quoteTAST: ExprRewrite encountered"
