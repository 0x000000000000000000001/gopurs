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
  ) where

import Prelude
import Data.Array as Array
import Data.List as List
import Data.Map (Map)
import Data.Maybe (Maybe)
import Effect.Ref (Ref)
import Gopurs.CodegenState (CodegenState, FunctionInfo)
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
  , params :: Array String
  , loopParams :: Array String
  , goTypes :: Array GoType
  , fRet :: GoType
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
  { codegenStateRef :: Ref CodegenState
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
