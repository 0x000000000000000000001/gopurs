module Gopurs.Ownership.Candidates (collect) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (foldl)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Newtype (unwrap)
import Data.Set as Set
import Data.String as String
import Data.String.Pattern (Pattern(..))
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..), fst)
import Gopurs.CodegenState (CodegenMetadata)
import Gopurs.GoAst (GoType(..), sanitizeName)
import Gopurs.GoTypes (exprTypeToGoType)
import Gopurs.Ownership.Analysis (modulePrefix, strip)
import Gopurs.Ownership.Types (Candidate, Candidates, LocalRef, TreeSpec)
import PureScript.Backend.Optimizer.Convert (BackendModule)
import PureScript.Backend.Optimizer.CoreFn (ExprType(ADT, Func), Ident(..), ModuleName, Qualified(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..))

type CandidateInput =
  { original :: Ident
  , spec :: TreeSpec
  , args :: Array LocalRef
  , argTypes :: Array GoType
  , body :: NeutralExpr
  }

-- Shape/signature admission precedes body validation. Reserve names in source
-- order before building the lookup map, even for workers later rejected by the
-- fixed point: changing that order would change other workers' suffixes.
collect :: CodegenMetadata -> BackendModule -> Candidates
collect metadata mod =
  let specs = treeSpecs metadata mod
      found = Array.mapMaybe (candidate metadata mod specs) (Array.concatMap _.bindings mod.bindings)
  in Map.fromFoldable $ map (\fn -> Tuple fn.original fn) (reserveNames mod found)

annotation :: NeutralExpr -> Maybe ExprType
annotation (NeutralExpr syn) = case syn of
  Typed ty _ -> Just ty
  TypeApp inner _ -> annotation inner
  _ -> Nothing

arrow :: ExprType -> { args :: Array ExprType, result :: ExprType }
arrow (Func args result) = let next = arrow result in { args: args <> next.args, result: next.result }
arrow ty = { args: [], result: ty }

abstractions :: NeutralExpr -> { args :: Array LocalRef, body :: NeutralExpr }
abstractions expr = case strip expr of
  Abs args body -> let next = abstractions body in { args: NEA.toArray args <> next.args, body: next.body }
  UncurriedAbs args body | not (Array.null args) ->
    let next = abstractions body in { args: args <> next.args, body: next.body }
  _ -> { args: [], body: expr }

scalarType :: CodegenMetadata -> ModuleName -> ExprType -> Maybe GoType
scalarType metadata current ty = do
  let goType = exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors (modulePrefix current) ty
  guard case goType of
    TypeInt64 -> true
    TypeFloat64 -> true
    TypeBool -> true
    TypeString -> true
    TypeUint32 -> true
    _ -> false
  pure goType

treeSpecs :: CodegenMetadata -> BackendModule -> Array TreeSpec
treeSpecs metadata mod = Array.mapMaybe make mod.dataDecls
  where
  make decl = do
    guard (Array.null decl.vars)
    let fullName = unwrap mod.name <> "." <> decl.name
        ty = ADT fullName (String.split (Pattern ".") fullName) []
        nodes = Array.filter (not <<< Array.null <<< _.fields) decl.constructors
        nullary = Array.filter (Array.null <<< _.fields) decl.constructors
    node <- case nodes of
      [ node ] | Array.length nullary <= 1 -> Just node
      _ -> Nothing
    pointer <- case exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors (modulePrefix mod.name) ty of
      result@(TypeStructPointer _) -> Just result
      _ -> Nothing
    fields <- traverse (\field -> if field == ty then Just pointer else scalarType metadata mod.name field) node.fields
    guard (Array.elem pointer fields)
    pure
      { ty, goType: pointer, fields
      , constructor: Qualified (Just mod.name) (Ident node.name)
      , leaf: map (\leaf -> Qualified (Just mod.name) (Ident leaf.name)) (Array.head nullary)
      }

candidate :: CodegenMetadata -> BackendModule -> Array TreeSpec -> Tuple Ident NeutralExpr -> Maybe CandidateInput
candidate metadata mod specs (Tuple original expr) = do
  signature <- arrow <$> annotation expr
  spec <- Array.find (\spec -> spec.ty == signature.result) specs
  let lambda = abstractions expr
  guard (not (Array.null lambda.args) && Array.length lambda.args == Array.length signature.args)
  argTypes <- traverse (\ty -> if ty == spec.ty then Just spec.goType else scalarType metadata mod.name ty) signature.args
  guard (Array.elem spec.goType argTypes)
  pure { original, spec, args: lambda.args, argTypes, body: lambda.body }

-- Ownership reserves an entry/consume pair, including foreign and sanitized
-- user names. Its paired Go namespace and suffix policy differ from the fusion
-- workers' single source-name reservation.
reserveNames :: BackendModule -> Array CandidateInput -> Array Candidate
reserveNames mod functions = _.functions $ foldl assign { reserved, functions: [] } functions
  where
  names = map fst (Array.concatMap _.bindings mod.bindings)
    <> Array.fromFoldable (Map.keys mod.foreign)
  reserved = Set.fromFoldable $ map (\name -> "Call_" <> modulePrefix mod.name <> "_" <> sanitizeName (unwrap name)) names
  assign acc fn =
    let choose index =
          let worker = Ident ("__gopurs_owned_" <> unwrap fn.original <> "_" <> show index)
              native = "Call_" <> modulePrefix mod.name <> "_" <> sanitizeName (unwrap worker)
              consume = native <> "_consume"
          in if Set.member native acc.reserved || Set.member consume acc.reserved then choose (index + 1)
             else { original: fn.original, spec: fn.spec, args: fn.args, argTypes: fn.argTypes, body: fn.body
                  , worker, native, consume }
        renamed = choose 0
    in { reserved: Set.insert renamed.native $ Set.insert renamed.consume acc.reserved
       , functions: Array.snoc acc.functions renamed }
