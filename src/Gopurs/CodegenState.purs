module Gopurs.CodegenState
  ( CodegenMetadataRow
  , CodegenMetadata
  , CodegenMetadataWith
  , SharedMetadataRow
  , CodegenState
  , FunctionInfo
  ) where

import Data.Map (Map)
import Data.Set (Set)
import Data.Set as Set
import Data.Tuple (Tuple)
import Gopurs.Cache.LookupTable (LookupTable)
import Gopurs.AdtMetadata (PointerAdtLeaf, PointerAdtPaths)
import Gopurs.ClassMetadata (ClassFields)
import Gopurs.ConstructorMetadata (ConstructorTypes)
import Gopurs.GoAst (GoDecl, GoType)
import Gopurs.GoSymbol (GoReference)
import Gopurs.ReboxMetadata (ReboxFieldIndex)
import PureScript.Backend.Optimizer.CoreFn (ExprType)

type FunctionInfo =
  { reference :: GoReference
  , fArgs :: Array GoType
  , fRet :: GoType
  , arity :: Int
  }

-- Prepared metadata is shared by the generator and value conversions.
-- It is passed directly to each translation, never stored in a mutable Ref.
-- Driver.Prepare builds declaration/type tables from the original TAST, with
-- synthetic class data declarations used only for pointer layouts. Driver.Build
-- supplies the published native signatures for the current emission batch;
-- Driver.Output adds the current module's FFI signatures before translation.
type SharedMetadataRow :: Row Type
type SharedMetadataRow =
  ( elidedCtors :: Set.Set String
  , ctorTypes :: ConstructorTypes
  , pointerAdtPaths :: PointerAdtPaths
  , pointerAdtNodes :: Set String
  , pointerAdtLeaves :: Map String PointerAdtLeaf
  , enumAdts :: Set.Set String
  , enumCtors :: Set.Set String
  , globalFunctions :: Map String FunctionInfo
  -- | Infos d'appel natif des FFI du module courant (clé : ident brut).
  , ffiFunctions :: Map String FunctionInfo
  , classDeclsFields :: ClassFields
  , reboxFields :: ReboxFieldIndex
  )

type CodegenMetadataWith globals = { globalTypes :: globals | SharedMetadataRow }

type CodegenMetadataRow = ( globalTypes :: LookupTable ExprType | SharedMetadataRow )

type CodegenMetadata = { | CodegenMetadataRow }

-- Only output accumulated during one translation belongs in the mutable state.
type CodegenState =
  { declarations :: Array GoDecl
  , globalId :: Int
  , reboxPairs :: Set.Set (Tuple GoType GoType)
  }
