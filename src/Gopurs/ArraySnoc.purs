module Gopurs.ArraySnoc (worker) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Maybe (Maybe(..))
import Data.String as String
import Data.String.Pattern (Pattern(..))
import Data.Tuple (Tuple(..))
import Gopurs.CallAnalysis (collectGoSpine, getGoSpineArgs, qualifiedTarget)
import Gopurs.ExprAnalysis (unwrapTcoExpr)
import Gopurs.ExprContext (ExprResult, StmtTree(..))
import Gopurs.GoAst (GoExpr(..), GoType(..), goTypeToStr, rawGo)
import Gopurs.GoFunctions (Parameters)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.CoreFn (ModuleName(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax as Syn

-- Match the final worker, including specialized copies. Its TAST-derived ABI
-- already supplies the native element layout. A name suffix alone is not an
-- origin proof: require the complete canonical thaw/push/freeze body as well.
worker :: String -> String -> Parameters -> GoType -> TcoExpr -> Maybe ExprResult
worker moduleName name params resultType body = do
  guard (moduleName == "Data_Array" && (name == "snoc" || String.indexOf (Pattern "snoc__") name == Just 0))
  { source, element, elementType } <- case params, resultType of
    [ Tuple source (TypeNativeArray item), Tuple element added ], TypeNativeArray returned
      | item == added && item == returned && source /= element -> Just { source, element, elementType: item }
    _, _ -> Nothing
  guard (canonicalBody source element body)
  let
    sourceExpr = GoVar source
    length = GoCall (GoVar "len") [ sourceExpr ]
    output = "snoc_output"
    statements =
      [ GoAssign output (GoCall (GoVar "make")
          [ rawGo ("[]" <> goTypeToStr elementType), GoBinOp "+" length (GoInt 1) ])
      , GoCall (GoVar "copy") [ GoVar output, sourceExpr ]
      , GoMutate (output <> "[len(" <> source <> ")]") (GoVar element)
      ]
  -- Always allocate fresh storage, even if the input slice has spare capacity.
  -- copy retains element references; there is no deep boxing or reconstruction.
  pure { stmts: StmtEmpty, expr: GoCall (GoFuncLit [] statements (GoVar output) resultType) []
    , exprType: resultType, nextId: 0 }

canonicalBody :: String -> String -> TcoExpr -> Boolean
canonicalBody source element body = case curriedArguments "Control.Monad.ST.Internal" "run" body of
  Just [ effect ] -> case unwrapTcoExpr effect of
    Syn.Let pushIdent pushLevel pushValue afterPush ->
      case curriedArguments "Data.Array.ST" "push" pushValue, unwrapTcoExpr afterPush of
        Just [ item ], Syn.EffectBind arrayIdent arrayLevel thaw afterThaw
          | local element item && effectArgument "thawImpl" source thaw ->
          let
            pushName = localId pushIdent pushLevel
            arrayName = localId arrayIdent arrayLevel
          in case unwrapTcoExpr afterThaw of
            Syn.EffectBind unusedIdent unusedLevel push freeze ->
              let names = [ source, element, pushName, arrayName, localId unusedIdent unusedLevel ]
              in Array.length (Array.nub names) == 5
                && localApplication pushName arrayName push
                && effectArgument "unsafeFreezeImpl" arrayName freeze
            _ -> false
        _, _ -> false
    _ -> false
  _ -> false

curriedArguments :: String -> String -> TcoExpr -> Maybe (Array TcoExpr)
curriedArguments moduleName name expr = do
  let Tuple fn spine = collectGoSpine expr
  target <- qualifiedTarget fn
  guard (target.mbMod == Just (ModuleName moduleName) && target.name == name)
  pure (getGoSpineArgs spine)

localApplication :: String -> String -> TcoExpr -> Boolean
localApplication fn argument expr = case unwrapTcoExpr expr of
  Syn.App target args -> case NEA.toArray args of
    [ value ] -> local fn target && local argument value
    _ -> false
  _ -> false

effectArgument :: String -> String -> TcoExpr -> Boolean
effectArgument name argument expr = case unwrapTcoExpr expr of
  Syn.UncurriedEffectApp target [ value ] -> case qualifiedTarget target of
    Just resolved -> resolved.mbMod == Just (ModuleName "Data.Array.ST") && resolved.name == name && local argument value
    _ -> false
  _ -> false

local :: String -> TcoExpr -> Boolean
local name expr = case unwrapTcoExpr expr of
  Syn.Local ident level -> localId ident level == name
  _ -> false
