module Gopurs.Driver.Prepare.Proof
  ( NamedHash, ModuleProof, Payload, IndexEntry, State
  , moduleProof, moduleProofWithReads, permitsLiteralEdit, placeholder
  ) where

import Prelude
import Data.Array as Array
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Newtype (unwrap)
import Data.Set (Set)
import Data.Set as Set
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Effect (Effect)
import Gopurs.Cache.Wire as Wire
import PureScript.Backend.Optimizer.CoreFn (Ann(..), Bind(..), Binding(..), CaseAlternative(..), CaseGuard(..), Expr(..), Guard(..), Literal(..), Module(..), Prop(..), emptySpan)

type NamedHash = { name :: String, hash :: String }
type ModuleProof = { name :: String, shape :: String, bindings :: Array NamedHash }
type Payload = { name :: String, hash :: String }
type IndexEntry = { name :: String, keys :: Array String }
type State =
  { schema :: Int, context :: String, directives :: String, snapshot :: String
  , sources :: Array NamedHash, modules :: Array ModuleProof
  , index :: Array IndexEntry, payloads :: Array Payload, bodyReads :: Array String
  }

-- Source locations are diagnostic, never specialization inputs. Keep types,
-- metadata, binding structure, field labels, arities and static expressions.
-- Only expression scalar data values are erased: isStatic rejects all of them,
-- and neither collection nor monomorphization branches on their values.
locationless :: Ann -> Ann
locationless (Ann ann) = Ann (ann { span = emptySpan, sourceUsage = Nothing })

shapeExpr :: Expr Ann -> Expr Ann
shapeExpr = case _ of
  ExprLit ann lit -> ExprLit ann case lit of
    LitInt _ -> LitInt 0
    LitNumber _ -> LitNumber 0.0
    LitString _ -> LitString ""
    LitChar _ -> LitChar 'x'
    LitBoolean _ -> LitBoolean false
    other -> map shapeExpr other
  ExprApp ann fn arg -> ExprApp ann (shapeExpr fn) (shapeExpr arg)
  ExprAbs ann ident body -> ExprAbs ann ident (shapeExpr body)
  ExprLet ann binds body -> ExprLet ann (map shapeBind binds) (shapeExpr body)
  ExprTypeApp ann body ty -> ExprTypeApp ann (shapeExpr body) ty
  ExprCase ann exprs alts -> ExprCase ann (map shapeExpr exprs) (map shapeAlt alts)
  ExprAccessor ann expr prop -> ExprAccessor ann (shapeExpr expr) prop
  ExprUpdate ann expr props -> ExprUpdate ann (shapeExpr expr) (map (\(Prop p value) -> Prop p (shapeExpr value)) props)
  expr -> expr
  where
  shapeAlt (CaseAlternative binders guard) = CaseAlternative binders case guard of
    Unconditional expr -> Unconditional (shapeExpr expr)
    Guarded guards -> Guarded (map (\(Guard a b) -> Guard (shapeExpr a) (shapeExpr b)) guards)

shapeBind :: Bind Ann -> Bind Ann
shapeBind = case _ of
  NonRec (Binding ann ident expr) -> NonRec (Binding ann ident (shapeExpr expr))
  Rec bindings -> Rec (map (\(Binding ann ident expr) -> Binding ann ident (shapeExpr expr)) bindings)

moduleProof :: Module Ann -> Effect ModuleProof
moduleProof = moduleProofWith (const true)

-- Only bodies actually read during preparation can invalidate another prepared
-- module. The whole-module shape still checks every declaration and type.
moduleProofWithReads :: Set String -> Module Ann -> Effect ModuleProof
moduleProofWithReads reads = moduleProofWith (flip Set.member reads)

moduleProofWith :: (String -> Boolean) -> Module Ann -> Effect ModuleProof
moduleProofWith keep (Module mod) = do
  let decls = map (map locationless) mod.decls
  let canonical = Module (mod { span = emptySpan, imports = map (map locationless) mod.imports, decls = map shapeBind decls })
  shape <- Wire.hash <$> Wire.write canonical
  bindings <- traverse (\binding@(Binding _ ident _) -> do
    hash <- Wire.hash <$> Wire.write binding
    pure { name: unwrap mod.name <> "." <> unwrap ident, hash }) (Array.filter (\(Binding _ ident _) -> keep (unwrap mod.name <> "." <> unwrap ident)) (Array.concatMap flatten decls))
  pure { name: unwrap mod.name, shape, bindings }
  where
  flatten (NonRec binding) = [ binding ]
  flatten (Rec bindings) = bindings

-- The shape proof applies at every fixed-point round, so it cannot retain an
-- unsupported specialization cycle. A changed body unfolded as a dictionary
-- would change another module's emitted expression and must fall back.
permitsLiteralEdit :: Set String -> ModuleProof -> ModuleProof -> Boolean
permitsLiteralEdit bodyReads old current =
  old.name == current.name && old.shape == current.shape
    && Array.length old.bindings == Array.length current.bindings
    && Array.all unchangedOrUnread current.bindings
  where
  previous = Map.fromFoldable (map (\entry -> Tuple entry.name entry.hash) old.bindings)
  unchangedOrUnread entry = case Map.lookup entry.name previous of
    Nothing -> false
    Just oldHash -> oldHash == entry.hash || not (Set.member entry.name bodyReads)

-- monomorphize exports all final declarations. Generated clones are exported
-- too, hence no private untyped global is lost when an inert source placeholder
-- carries all its source declarations as exports.
placeholder :: Boolean -> Module Ann -> Module Ann
placeholder specialized (Module mod) =
  if specialized then Module (mod { exports = Array.concatMap names mod.decls })
  else Module mod
  where
  names (NonRec (Binding _ ident _)) = [ ident ]
  names (Rec bindings) = map (\(Binding _ ident _) -> ident) bindings
