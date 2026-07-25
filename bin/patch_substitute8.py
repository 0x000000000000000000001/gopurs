import re

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/Monomorphize/Substitute.purs', 'r') as f:
    content = f.read()

new_code = """
mapTcoExprTypes :: (ExprType -> ExprType) -> TcoExpr -> TcoExpr
mapTcoExprTypes f = go
  where
  go (TcoExpr a syn) = case syn of
    Typed ty inner -> TcoExpr a (Typed (f ty) (go inner))
    Var v -> TcoExpr a (Var v)
    App fn args -> TcoExpr a (App (go fn) (map go args))
    Abs args body -> TcoExpr a (Abs args (go body))
    UncurriedApp fn args -> TcoExpr a (UncurriedApp (go fn) (map go args))
    UncurriedAbs args body -> TcoExpr a (UncurriedAbs args (go body))
    UncurriedEffectApp fn args -> TcoExpr a (UncurriedEffectApp (go fn) (map go args))
    UncurriedEffectAbs args body -> TcoExpr a (UncurriedEffectAbs args (go body))
    Accessor obj prop -> TcoExpr a (Accessor (go obj) prop)
    Update obj props -> TcoExpr a (Update (go obj) (map (\\(Prop k v) -> Prop k (go v)) props))
    CtorSaturated mbMn ty pn ident args -> TcoExpr a (CtorSaturated mbMn ty pn ident (map (\\(Tuple s v) -> Tuple s (go v)) args))
    CtorDef ty pn ident args -> TcoExpr a (CtorDef ty pn ident args)
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
    Local mbIdent lvl -> TcoExpr a (Local mbIdent lvl)
    Fail str -> TcoExpr a (Fail str)
    PrimUndefined -> TcoExpr a PrimUndefined
"""

if "mapTcoExprTypes" not in content:
    content = content + new_code

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/Monomorphize/Substitute.purs', 'w') as f:
    f.write(content)
