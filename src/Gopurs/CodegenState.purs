module Gopurs.CodegenState
  ( CodegenMetadataRow
  , CodegenMetadata
  , CodegenState
  , FunctionInfo
  ) where

import Data.Map (Map)
import Data.Map as Map
import Data.Set (Set)
import Data.Set as Set
import Data.Tuple (Tuple)
import Gopurs.GoAst (GoDecl, GoType)
import PureScript.Backend.Optimizer.CoreFn (ExprType)

type FunctionInfo =
  { fullName :: String
  , fArgs :: Array GoType
  , fRet :: GoType
  , arity :: Int
  }

-- Prepared metadata is shared by the generator and value conversions.
-- It is passed directly to each translation, never stored in a mutable Ref.
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
  , globalFunctions :: Map String FunctionInfo
  , classDeclsFields :: Map String { vars :: Array String, fields :: Array { name :: String, "type" :: ExprType } }
  )

type CodegenMetadata = { | CodegenMetadataRow }

-- Only output accumulated during one translation belongs in the mutable state.
type CodegenState =
  { declarations :: Array GoDecl
  , globalId :: Int
  , reboxPairs :: Set.Set (Tuple GoType GoType)
  }
