import re

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/Monomorphize/Substitute.purs', 'r') as f:
    content = f.read()

old_func = """substituteAst :: Map String (Array ExprType) -> (ExprType -> String) -> (TcoExpr -> ExprType) -> TcoExpr -> TcoExpr
substituteAst insts mangle getTy = go
  where
  go expr@(TcoExpr a syn) = case syn of
    Typed ty inner -> TcoExpr a (Typed ty (go inner))
    Var (Qualified mbMn (Ident name)) ->
      let 
        fullName = case mbMn of
          Just mn -> unwrap mn <> "." <> name
          Nothing -> name
      in case Map.lookup fullName insts of
        Just concretes ->
          let
            varTy = getTy expr
            mangled = name <> "__" <> mangle varTy
          in TcoExpr a (Var (Qualified mbMn (Ident mangled)))
        Nothing -> expr
    App fn args -> TcoExpr a (App (go fn) (map go args))
    Abs args body -> TcoExpr a (Abs args (go body))
    UncurriedApp fn args -> TcoExpr a (UncurriedApp (go fn) (map go args))
    UncurriedAbs args body -> TcoExpr a (UncurriedAbs args (go body))
    UncurriedEffectApp fn args -> TcoExpr a (UncurriedEffectApp (go fn) (map go args))
    UncurriedEffectAbs args body -> TcoExpr a (UncurriedEffectAbs args (go body))
    Accessor obj prop -> TcoExpr a (Accessor (go obj) prop)
    Update obj props -> TcoExpr a (Update (go obj) (map (\\(Prop k v) -> Prop k (go v)) props))
    CtorSaturated mbMn ty pn ident args -> TcoExpr a (CtorSaturated mbMn ty pn ident (map (\\(Tuple s v) -> Tuple s (go v)) args))
    CtorDef _ _ _ _ -> expr
    LetRec lvl bindings body -> TcoExpr a (LetRec lvl (map (\\(Tuple ident val) -> Tuple ident (go val)) bindings) (go body))
    Let ident lvl val body -> TcoExpr a (Let ident lvl (go val) (go body))
    EffectBind ident lvl val body -> TcoExpr a (EffectBind ident lvl (go val) (go body))
    EffectPure val -> TcoExpr a (EffectPure (go val))
    EffectDefer val -> TcoExpr a (EffectDefer (go val))
    Branch pairs def -> TcoExpr a (Branch (map (\\(Pair p r) -> Pair (go p) (go r)) pairs) (go def))
    Lit lit -> TcoExpr a (Lit (case lit of
      LitArray arr -> LitArray (map go arr)
      LitRecord props -> LitRecord (map (\\(Prop k v) -> Prop k (go v)) props)
      _ -> lit))
    PrimEffect eff -> TcoExpr a (PrimEffect (case eff of
      EffectRefNew val -> EffectRefNew (go val)
      EffectRefRead ref -> EffectRefRead (go ref)
      EffectRefWrite ref val -> EffectRefWrite (go ref) (go val)
      _ -> eff))
    PrimOp op -> TcoExpr a (PrimOp (case op of
      Op1 o1 v1 -> Op1 o1 (go v1)
      Op2 o2 v1 v2 -> Op2 o2 (go v1) (go v2)))
    Local _ _ -> expr
    Fail _ -> expr
    PrimUndefined -> expr"""

new_func = """substituteAst :: Map String (Array ExprType) -> (ExprType -> String) -> TcoExpr -> TcoExpr
substituteAst insts mangle = go Nothing
  where
  go mbTy (TcoExpr a syn) = case syn of
    Typed ty inner -> TcoExpr a (Typed ty (go (Just ty) inner))
    Var (Qualified mbMn (Ident name)) ->
      let 
        fullName = case mbMn of
          Just mn -> unwrap mn <> "." <> name
          Nothing -> name
      in case Map.lookup fullName insts of
        Just concretes ->
          let
            varTy = case mbTy of
                      Just ty -> ty
                      Nothing -> TypeVar "gopurs_runtime.Value"
            mangled = name <> "__" <> mangle varTy
          in TcoExpr a (Var (Qualified mbMn (Ident mangled)))
        Nothing -> TcoExpr a (Var (Qualified mbMn (Ident name)))
    App fn args -> TcoExpr a (App (go Nothing fn) (map (go Nothing) args))
    Abs args body -> TcoExpr a (Abs args (go Nothing body))
    UncurriedApp fn args -> TcoExpr a (UncurriedApp (go Nothing fn) (map (go Nothing) args))
    UncurriedAbs args body -> TcoExpr a (UncurriedAbs args (go Nothing body))
    UncurriedEffectApp fn args -> TcoExpr a (UncurriedEffectApp (go Nothing fn) (map (go Nothing) args))
    UncurriedEffectAbs args body -> TcoExpr a (UncurriedEffectAbs args (go Nothing body))
    Accessor obj prop -> TcoExpr a (Accessor (go Nothing obj) prop)
    Update obj props -> TcoExpr a (Update (go Nothing obj) (map (\\(Prop k v) -> Prop k (go Nothing v)) props))
    CtorSaturated mbMn ty pn ident args -> TcoExpr a (CtorSaturated mbMn ty pn ident (map (\\(Tuple s v) -> Tuple s (go Nothing v)) args))
    CtorDef _ _ _ _ -> TcoExpr a syn
    LetRec lvl bindings body -> TcoExpr a (LetRec lvl (map (\\(Tuple ident val) -> Tuple ident (go Nothing val)) bindings) (go Nothing body))
    Let ident lvl val body -> TcoExpr a (Let ident lvl (go Nothing val) (go Nothing body))
    EffectBind ident lvl val body -> TcoExpr a (EffectBind ident lvl (go Nothing val) (go Nothing body))
    EffectPure val -> TcoExpr a (EffectPure (go Nothing val))
    EffectDefer val -> TcoExpr a (EffectDefer (go Nothing val))
    Branch pairs def -> TcoExpr a (Branch (map (\\(Pair p r) -> Pair (go Nothing p) (go Nothing r)) pairs) (go Nothing def))
    Lit lit -> TcoExpr a (Lit (case lit of
      LitArray arr -> LitArray (map (go Nothing) arr)
      LitRecord props -> LitRecord (map (\\(Prop k v) -> Prop k (go Nothing v)) props)
      _ -> lit))
    PrimEffect eff -> TcoExpr a (PrimEffect (case eff of
      EffectRefNew val -> EffectRefNew (go Nothing val)
      EffectRefRead ref -> EffectRefRead (go Nothing ref)
      EffectRefWrite ref val -> EffectRefWrite (go Nothing ref) (go Nothing val)
      _ -> eff))
    PrimOp op -> TcoExpr a (PrimOp (case op of
      Op1 o1 v1 -> Op1 o1 (go Nothing v1)
      Op2 o2 v1 v2 -> Op2 o2 (go Nothing v1) (go Nothing v2)))
    Local _ _ -> TcoExpr a syn
    Fail _ -> TcoExpr a syn
    PrimUndefined -> TcoExpr a syn"""

content = content.replace(old_func, new_func)

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/Monomorphize/Substitute.purs', 'w') as f:
    f.write(content)
