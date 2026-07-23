module Gopurs.GoAst where

import Prelude
import Data.Tuple (Tuple(..))

data GoExpr
  = GoVar String
  | GoString String
  | GoInt Int
  | GoCall GoExpr (Array GoExpr)
  | GoSelector GoExpr String
  | GoFunc String GoExpr
  | GoBlock (Array GoExpr)
  | GoReturn GoExpr
  | GoAssign String GoExpr
  | GoRecordDict (Array (Tuple String GoExpr))
  | GoRecordUpdateDict GoExpr (Array (Tuple String GoExpr))
  | GoIIFE String GoExpr GoExpr
  | GoLetRec (Array (Tuple String GoExpr)) GoExpr
  | GoRecordAccess GoExpr String
  | GoConstructor String (Array GoExpr)
  | GoConstructorAccess GoExpr Int
  | GoBranch (Array (Tuple GoExpr GoExpr)) GoExpr
  | GoBinOp String GoExpr GoExpr
  | GoTypeAssertion GoExpr String
  | GoRaw String
  | GoFor String (Array GoExpr)
  | GoContinue String
  | GoMutate String GoExpr
  | GoIfElse GoExpr (Array GoExpr) (Array GoExpr)

derive instance eqGoExpr :: Eq GoExpr

type GoDecl =
  { identifier :: String
  , expression :: GoExpr
  }

type GoFile =
  { packageName :: String
  , imports :: Array String
  , decls :: Array GoDecl
  , foreigns :: Array { pursName :: String, goName :: String }
  }
