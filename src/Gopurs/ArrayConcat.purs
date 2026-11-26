module Gopurs.ArrayConcat (nativeArgument) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Tuple (Tuple(..))
import Gopurs.ArrayIntrinsics.Source (arraySource)
import Gopurs.CallAnalysis (collectGoSpine, getGoSpineArgs, qualifiedTarget)
import Gopurs.CallArguments as CallArguments
import Gopurs.ExprAnalysis (getExprType)
import Gopurs.ExprContext (ExprContext, ExprResult, TranslateExpr, childContext, flattenStmts, StmtTree(..))
import Gopurs.GoAst (GoExpr(..), GoType(..), rawGo)
import Gopurs.GoConversions (coerceGoExpr)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), ModuleName(..))

-- Only a concrete native consumer asks for this representation. Both array
-- values are evaluated before either buffer is read, exactly as in concat's
-- saturated FFI call. In particular, evaluating the right input may mutate a
-- boxed left input through foreign code; no early unboxing is permitted.
nativeArgument :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> Maybe ExprResult
nativeArgument translate context nextId expression = do
  let Tuple function spine = collectGoSpine expression
  target <- qualifiedTarget function
  guard (target.mbMod == Just (ModuleName "Data.Semigroup") && target.name == "concatArray")
  let args = getGoSpineArgs spine
  guard (Array.length args == 2)
  guard (getExprType expression == Array Int)
  let
    captured = CallArguments.capture { prefix: "concatInt_arg_", markUsed: false }
      (\_ currentId argument -> translate (childContext context Nothing) currentId argument) nextId args
    suffix = show captured.nextId
    source argumentIndex =
      let name = "concatInt_source_" <> show argumentIndex <> "_" <> suffix
          inputName = case Array.index captured.exprs argumentIndex of
            Just (GoVar variable) -> variable
            _ -> ""
      in arraySource inputName name (fromMaybe TypeValue (Array.index captured.exprTypes argumentIndex))
    left = source 0
    right = source 1
    output = "concatInt_output_" <> suffix
    offset = "concatInt_offset_" <> suffix
    index = "concatInt_index_" <> suffix
    value = "concatInt_value_" <> suffix
    lengthOf targetName = GoCall (GoVar "len") [ rawGo targetName ]
    copy sourceInfo = GoForRange (index <> ", " <> value <> " := range " <> sourceInfo.target)
      [ GoMutate (output <> "[" <> offset <> "+" <> index <> "]")
          (coerceGoExpr context.codegenStateRef context.modNameStr (GoVar value) sourceInfo.elementType TypeInt64) ]
    body = flattenStmts captured.stmts <>
      [ left.assignment, right.assignment
      , GoAssign output (GoCall (GoVar "make") [ rawGo "[]int64", GoBinOp "+" (lengthOf left.target) (lengthOf right.target) ])
      , GoAssign offset (GoInt 0)
      , copy left
      , GoMutate offset (lengthOf left.target)
      , copy right
      ]
  pure
    { stmts: StmtEmpty
    , expr: GoCall (GoFuncLit [] body (GoVar output) (TypeNativeArray TypeInt64)) []
    , exprType: TypeNativeArray TypeInt64
    , nextId: captured.nextId + 1
    }
