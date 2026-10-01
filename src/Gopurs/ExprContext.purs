module Gopurs.ExprContext
  ( LocalBinding
  , LocalEnv
  , ModuleFunctions
  , LoopTarget
  , LoopContext
  , ExprOptions
  , ExprResult
  , StmtTree(..)
  , flattenStmts
  , wrapInStmts
  , ExprContext
  , TranslateExpr
  , childContext
  , bindParameters
  ) where

import Prelude
import Data.Array as Array
import Data.List as List
import Data.Foldable (foldl)
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Tuple (Tuple(..))
import Effect.Ref (Ref)
import Gopurs.CodegenState (CodegenMetadata, CodegenState, FunctionInfo)
import Gopurs.GoAst (GoExpr(..), GoType)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.CoreFn (ExprType)

type LocalBinding =
  { name :: String
  , goType :: GoType
  }

-- Keys are original localId values; each binding's name is the emitted Go name,
-- which may have been renamed.
type LocalEnv = Map String LocalBinding

type ModuleFunctions = Map String FunctionInfo

type LoopTarget =
  { ident :: String
  , loopParams :: Array String
  , goTypes :: Array GoType
  }

type LoopContext = Array LoopTarget

type ExprOptions =
  { isTail :: Boolean
  , inEffectBlock :: Boolean
  }

type ExprResult =
  { stmts :: StmtTree
  , expr :: GoExpr
  , exprType :: GoType
  , nextId :: Int
  }

data StmtTree = StmtEmpty | StmtLeaf GoExpr | StmtAppend StmtTree StmtTree

instance Semigroup StmtTree where
  append StmtEmpty a = a
  append a StmtEmpty = a
  append a b = StmtAppend a b

instance Monoid StmtTree where
  mempty = StmtEmpty

flattenStmts :: StmtTree -> Array GoExpr
flattenStmts tree = Array.fromFoldable (go List.Nil tree)
  where
  go acc StmtEmpty = acc
  go acc (StmtLeaf s) = List.Cons s acc
  go acc (StmtAppend a b) =
    let
      acc' = go acc b
    in
      go acc' a

wrapInStmts :: Array String -> StmtTree -> GoType -> GoExpr -> GoExpr
wrapInStmts _ stmts retType expr =
  let
    stmtsArr = flattenStmts stmts
  in
    if Array.length stmtsArr == 0 then expr
    else GoCall (GoFuncLit [] stmtsArr expr retType) []

-- Child translation receives an explicit context and returns the next free
-- identifier with its statements. Family emitters never import CodeGen.
type ExprContext =
  { metadata :: CodegenMetadata
  , codegenStateRef :: Ref CodegenState
  , depth :: Int
  , modNameStr :: String
  , recVars :: Array String
  , moduleFunctions :: ModuleFunctions
  , bound :: LocalEnv
  , tcoIdent :: Maybe String
  , loopCtx :: LoopContext
  , options :: ExprOptions
  , mbExpectedExprType :: Maybe ExprType
  }

type TranslateExpr = ExprContext -> Int -> TcoExpr -> ExprResult

-- Extend a lexical environment without discarding captured outer bindings.
-- Only parameters shadow their original localId keys.
bindParameters :: Array (Tuple String GoType) -> LocalEnv -> LocalEnv
bindParameters params bound =
  foldl (\acc (Tuple name goType) -> Map.insert name { name, goType } acc) bound params

-- Ordinary operands are non-tail values, outside the enclosing effect block.
-- Their expected type is explicit rather than inherited from the parent.
childContext :: ExprContext -> Maybe ExprType -> ExprContext
childContext context expected = context
  { depth = context.depth + 1
  , tcoIdent = Nothing
  , loopCtx = []
  , options = { isTail: false, inEffectBlock: false }
  , mbExpectedExprType = expected
  }
