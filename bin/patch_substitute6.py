import re

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/Monomorphize/Substitute.purs', 'r') as f:
    content = f.read()

# Replace the go function body completely
new_go = """  go expr@(TcoExpr a syn) = case syn of
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
    Update obj props -> TcoExpr a (Update (go obj) (map (\\p -> { key: p.key, value: go p.value }) props))
    CtorSaturated mbMn ty pn ident args -> TcoExpr a (CtorSaturated mbMn ty pn ident (map (\\(Tuple s v) -> Tuple s (go v)) args))
    CtorDef ty pn ident args -> expr
    LetRec lvl bindings body -> TcoExpr a (LetRec lvl (map (\\(Tuple ident val) -> Tuple ident (go val)) bindings) (go body))
    Let ident lvl val body -> TcoExpr a (Let ident lvl (go val) (go body))
    EffectBind ident lvl val body -> TcoExpr a (EffectBind ident lvl (go val) (go body))
    EffectPure val -> TcoExpr a (EffectPure (go val))
    EffectDefer val -> TcoExpr a (EffectDefer (go val))
    Branch pairs def -> TcoExpr a (Branch (map (\\p -> { predicate: go p.predicate, result: go p.result }) pairs) (go def))
    Lit lit -> TcoExpr a (Lit (case lit of
      ArrayLiteral arr -> ArrayLiteral (map go arr)
      ObjectLiteral props -> ObjectLiteral (map (\\(Tuple k v) -> Tuple k (go v)) props)
      _ -> lit))
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
    PrimUndefined -> expr
    Var q -> expr
"""

# Replace the part from `  go expr@(TcoExpr a syn) = case syn of` to the end of the function
import re
content = re.sub(r'  go expr@\(TcoExpr a syn\) = case syn of\n(.*?)(?=\n\n|\Z)', new_go, content, flags=re.DOTALL)

# Add imports for types used
imports = "import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), BackendEffect(..), BackendOperator(..))\nimport PureScript.Backend.Optimizer.CoreFn (Literal(..), ExprType(..), Ident(..), Qualified(..), Prop)\n"

# Remove duplicate BackendSyntax import
content = re.sub(r'import PureScript.Backend.Optimizer.Syntax.*?\n', '', content)
content = re.sub(r'import PureScript.Backend.Optimizer.CoreFn.*?\n', '', content)
content = content.replace('import Data.Newtype (unwrap)\n', 'import Data.Newtype (unwrap)\n' + imports)

with open('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/Monomorphize/Substitute.purs', 'w') as f:
    f.write(content)
