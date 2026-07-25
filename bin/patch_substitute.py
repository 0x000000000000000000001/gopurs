import re

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/Monomorphize/Substitute.purs', 'r') as f:
    content = f.read()

newCode = """import Gopurs.Monomorphize (InstantiationMap)
import PureScript.Backend.Optimizer.CoreFn (Ident(..), Qualified(..))

substituteAst :: InstantiationMap -> (ExprType -> String) -> (TcoExpr -> ExprType) -> TcoExpr -> TcoExpr
substituteAst insts mangle getTy = go
  where
  go expr@(TcoExpr a syn) = case syn of
    Typed ty inner -> TcoExpr a (Typed ty (go inner))
    Var (Qualified mbMn (Ident name)) ->
      case Map.lookup (case mbMn of { Just mn -> unwrap mn <> "." <> name; Nothing -> name }) insts of
        Just concretes ->
          let
            varTy = getTy expr
            -- If varTy matches one of the concretes, mangle the name!
            mangled = name <> "__" <> mangle varTy
          in TcoExpr a (Var (Qualified mbMn (Ident mangled)))
        Nothing -> expr
    App fn args -> TcoExpr a (App (go fn) (map go args))
    Abs args body -> TcoExpr a (Abs args (go body))
    Let bindings body -> TcoExpr a (Let (map (\(Tuple (Tuple ident val) alloc) -> Tuple (Tuple ident (go val)) alloc) bindings) (go body))
    Case items alts -> TcoExpr a (Case (map go items) (map (\alt -> alt { expr = go alt.expr, guards = map (\g -> g { expr = go g.expr }) alt.guards }) alts))
    Constructor mbMn ctor -> expr
    Literal lit -> TcoExpr a (Literal (case lit of
      ArrayLiteral arr -> ArrayLiteral (map go arr)
      ObjectLiteral props -> ObjectLiteral (map (\(Tuple k v) -> Tuple k (go v)) props)
      _ -> lit))
    Accessor prop obj -> TcoExpr a (Accessor prop (go obj))
    ObjectUpdate obj props -> TcoExpr a (ObjectUpdate (go obj) (map (\(Tuple k v) -> Tuple k (go v)) props))
    UncurriedApp fn args -> TcoExpr a (UncurriedApp (go fn) (map go args))
    UncurriedEffectApp fn args -> TcoExpr a (UncurriedEffectApp (go fn) (map go args))
    UncurriedEffectAbs args body -> TcoExpr a (UncurriedEffectAbs args (go body))
    PrimEffect eff -> TcoExpr a (PrimEffect (case eff of
      EffectRefNew val -> EffectRefNew (go val)
      EffectRefRead ref -> EffectRefRead (go ref)
      EffectRefWrite ref val -> EffectRefWrite (go ref) (go val)
      _ -> eff))
    PrimOp op -> TcoExpr a (PrimOp (case op of
      Op1 o1 v1 -> Op1 o1 (go v1)
      Op2 o2 v1 v2 -> Op2 o2 (go v1) (go v2)))
    Local mbIdent lvl -> expr
    Fail str -> expr

"""

if 'substituteAst' not in content:
    content = content + newCode

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/Monomorphize/Substitute.purs', 'w') as f:
    f.write(content)
