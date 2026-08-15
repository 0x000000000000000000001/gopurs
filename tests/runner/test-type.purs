module TestType where
import Prelude
import Effect (Effect)
import Effect.Console (logShow)
import PureScript.Backend.Optimizer.Convert (toBackendExpr)
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), ExprType(..), Ident(..))
import Gopurs.CodeGen (getExprType)
import Data.Maybe (Maybe(..))
-- It is too complex to reconstruct the AST just to test it...
