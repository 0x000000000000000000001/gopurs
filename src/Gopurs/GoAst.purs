module Gopurs.GoAst where

import Prelude

data GoExpr
  = GoVar String
  | GoString String
  | GoInt Int
  | GoCall GoExpr (Array GoExpr)
  | GoSelector GoExpr String
  | GoFunc (Array String) GoExpr
  | GoBlock (Array GoExpr)
  | GoReturn GoExpr
  | GoAssign String GoExpr
  | GoRaw String

derive instance eqGoExpr :: Eq GoExpr

type GoDecl =
  { identifier :: String
  , expression :: GoExpr
  }

type GoFile =
  { packageName :: String
  , imports :: Array String
  , decls :: Array GoDecl
  }
