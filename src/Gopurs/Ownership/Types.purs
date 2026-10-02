module Gopurs.Ownership.Types
  ( LocalRef
  , Path(..)
  , Scalar
  , Value(..)
  , Env
  , TreeSpec
  , Candidate
  , Candidates
  , Context
  , TreeTerm(..)
  , Argument(..)
  ) where

import Prelude

import Data.Map (Map)
import Data.Maybe (Maybe)
import Data.Tuple (Tuple)
import Gopurs.CodegenState (CodegenMetadata)
import Gopurs.GoAst (GoExpr, GoType)
import PureScript.Backend.Optimizer.CoreFn (ExprType, Ident, ModuleName, Qualified)
import PureScript.Backend.Optimizer.Semantics (NeutralExpr)
import PureScript.Backend.Optimizer.Syntax (Level)

type LocalRef = Tuple (Maybe Ident) Level

-- Aliases share a canonical root and field path, independent of local spelling.
data Path = Path String (Array Int)
derive instance eqPath :: Eq Path
derive instance ordPath :: Ord Path

type Scalar = { expr :: GoExpr, goType :: GoType, reads :: Array Path }
data Value = Tree Path | Scalar Scalar
type Env = Map LocalRef Value

type TreeSpec =
  { ty :: ExprType
  , goType :: GoType
  , constructor :: Qualified Ident
  , leaf :: Maybe (Qualified Ident)
  , fields :: Array GoType
  }

type Candidate =
  { original :: Ident
  , worker :: Ident
  , native :: String
  , consume :: String
  , spec :: TreeSpec
  , args :: Array LocalRef
  , argTypes :: Array GoType
  , body :: NeutralExpr
  }

type Candidates = Map Ident Candidate
type Context = { metadata :: CodegenMetadata, moduleName :: ModuleName, candidate :: Candidate, candidates :: Candidates }

-- Analysis produces Keep paths and borrowed scalar reads. Workers snapshot all
-- of them before mutation, replacing Keep with Existing and clearing reads.
-- Tree leaves must be disjoint throughout the admitted terminal term.
data TreeTerm
  = Keep Path
  | Empty
  | Construct (Array Argument)
  | Call Ident (Array Argument)
  | Existing GoExpr
data Argument = TreeArg TreeTerm | ScalarArg Scalar
