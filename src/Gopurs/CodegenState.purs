module Gopurs.CodegenState
  ( CodegenMetadataRow
  , CodegenMetadata
  , CodegenState
  ) where

import Data.Map (Map)
import Data.Map as Map
import Data.Set (Set)
import Data.Set as Set
import Data.Tuple (Tuple)
import Gopurs.GoAst (GoDecl, GoType)
import PureScript.Backend.Optimizer.CoreFn (ExprType)

-- Prepared metadata is shared by the generator and value conversions.
-- Only CodegenState adds mutable data for one invocation of translate.
type CodegenMetadataRow :: Row Type
type CodegenMetadataRow =
  ( elidedCtors :: Set.Set String
  , ctorTypes :: Map String { vars :: Array String, fields :: Array ExprType }
  , pointerAdtPaths :: Map String { ctorName :: String, arity :: Int }
  , pointerAdtNodes :: Set String
  , pointerAdtLeaves :: Map String { nodeBaseStruct :: String, nodeCtor :: String }
  , enumAdts :: Set.Set String
  , enumCtors :: Set.Set String
  , globalTypes :: Map.Map String ExprType
  , classDeclsFields :: Map String { vars :: Array String, fields :: Array { name :: String, "type" :: ExprType } }
  )

type CodegenMetadata = { | CodegenMetadataRow }

type CodegenState =
  { decls :: Array GoDecl
  , rawDecls :: Array String
  , globalId :: Int
  , reboxPairs :: Set.Set (Tuple GoType GoType)
  | CodegenMetadataRow
  }
