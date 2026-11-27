module Gopurs.ArraySort (emit) where

import Prelude

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (foldM)
import Data.Maybe (Maybe(..), fromMaybe, isJust)
import Data.Tuple (Tuple(..))
import Gopurs.AdtExprs as AdtExprs
import Gopurs.ArrayIntrinsics.Callback as Callback
import Gopurs.ArrayIntrinsics.Source (arraySource)
import Gopurs.CallAnalysis (collectGoSpine, getGoSpineArgs, qualifiedTarget)
import Gopurs.CallArguments (applyBoxed)
import Gopurs.CallArguments as CallArguments
import Gopurs.ExprAnalysis (unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, StmtTree(..), TranslateExpr, childContext)
import Gopurs.GoAst (GoExpr(..), GoType(..), rawGo)
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Ident(..), Literal(..), ModuleName(..), Qualified(..))
import PureScript.Backend.Optimizer.Syntax (BackendAccessor(GetProp), BackendOperator(..), BackendOperator1(..), BackendSyntax(..), Pair(..))

-- The standard Int comparator is pure and fully known. Preserve any dynamic
-- Ordering conversion, including all of its calls; closed pure Ordering -> Int
-- literal lambdas (such as Data.Array.sortBy's) specialize directly to native
-- comparisons without per-element closures. Custom Ord or comparison functions
-- keep their ordinary FFI path.
emit :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> Array TcoExpr -> Maybe ExprResult
emit translate context@{ metadata, codegenStateRef, modNameStr } nextId fn args = do
  guard (named "Data.Array" "sortByImpl" fn && Array.length args == 3)
  comparator <- Array.head args
  guard (intComparator comparator)
  let
    staticConversion = Array.index args 1 >>= staticOrderingConversion
    captured = CallArguments.capture { prefix: "arraySort_arg_", markUsed: true }
      (\index currentId expression -> case index of
        1 | isJust staticConversion -> { stmts: StmtEmpty, expr: rawGo "true", exprType: TypeBool, nextId: currentId }
        1 -> case Callback.nativeLambda translate context currentId 1 expression of
          Just native -> native
          Nothing -> translate (childContext context Nothing) currentId expression
        2 ->
          let result = translate (childContext context Nothing) currentId expression
          in case result.expr of
            GoUnboxIntArray boxedSource -> result { expr = boxedSource, exprType = TypeValue }
            _ -> result
        _ -> translate (childContext context Nothing) currentId expression)
      nextId args
    argument index = fromMaybe (rawGo "nil") (Array.index captured.exprs index)
    argumentType index = fromMaybe TypeValue (Array.index captured.exprTypes index)
    source = arraySource "arraySort_input" "arraySort_values" (argumentType 2)
    orderingType = ADT "Data.Ordering.Ordering" [ "Data", "Ordering", "Ordering" ] []
    converted name =
      let
        prepared = AdtExprs.prepareSaturated metadata modNameStr (Just (ModuleName "Data.Ordering")) name orderingType
        ordering = AdtExprs.saturated codegenStateRef prepared { exprs: [], exprTypes: [] }
      in case argumentType 1 of
        TypeFunc [ expected ] resultType -> coerceGoExpr codegenStateRef modNameStr
          (GoCall (argument 1) [ coerceGoExpr codegenStateRef modNameStr ordering.expr ordering.exprType expected ]) resultType TypeInt64
        _ -> coerceGoExpr codegenStateRef modNameStr
          (applyBoxed (boxGoExpr codegenStateRef modNameStr (argument 1) (argumentType 1))
            [ boxGoExpr codegenStateRef modNameStr ordering.expr ordering.exprType ]) TypeValue TypeInt64
    copyOutput = case source.elementType of
      TypeInt64 -> GoCall (GoVar "copy") [ GoVar "arraySort_output", rawGo source.target ]
      _ -> GoForRange ("index, value := range " <> source.target)
        [ GoMutate "arraySort_output[index]" (coerceGoExpr codegenStateRef modNameStr (GoVar "value") source.elementType TypeInt64) ]
    length = GoCall (GoVar "len") [ GoVar "arraySort_output" ]
    at index = GoIndex (GoVar "arraySort_output") index
    comparison a b = GoCall (GoVar "arraySort_compare") [ a, b ]
    dynamicSortBody ltExpr eqExpr gtExpr =
      let
        compare = GoFuncBlock [ Tuple "a" TypeInt64, Tuple "b" TypeInt64 ]
          [ GoIfElse (GoBinOp "<" (GoVar "a") (GoVar "b")) [ GoReturn ltExpr ] []
          , GoIfElse (GoBinOp "==" (GoVar "a") (GoVar "b")) [ GoReturn eqExpr ] []
          , GoReturn gtExpr
          ] TypeInt64
      in
        [ GoAssign "arraySort_compare" compare
        , GoAssign "arraySort_descending" (GoBinOp "<" (comparison (at (GoInt 1)) (at (GoInt 0))) (GoInt 0))
        , GoAssign "arraySort_monotone" (rawGo "true")
        , GoForRange "index := 2; index < len(arraySort_output); index++"
            [ GoAssign "order" (comparison (at (GoVar "index")) (at (GoBinOp "-" (GoVar "index") (GoInt 1))))
            , GoIfElse (GoBinOp "||"
                (GoBinOp "&&" (GoVar "arraySort_descending") (GoBinOp ">=" (GoVar "order") (GoInt 0)))
                (GoBinOp "&&" (GoPrefixOp "!" (GoVar "arraySort_descending")) (GoBinOp "<" (GoVar "order") (GoInt 0))))
                [ GoMutate "arraySort_monotone" (rawGo "false"), rawGo "break" ] []
            ]
        , GoIfElse (GoVar "arraySort_monotone")
            [ GoIfElse (GoVar "arraySort_descending") [ GoCall (slices "Reverse") [ GoVar "arraySort_output" ] ] []
            , GoReturn (GoVar "arraySort_output") ] []
        , GoCall (slices "SortStableFunc")
            [ GoVar "arraySort_output"
            , GoFuncLit [ Tuple "a" TypeInt64, Tuple "b" TypeInt64 ] []
                (GoCall (GoVar "int") [ comparison (GoVar "a") (GoVar "b") ]) (TypeInterface "int")
            ]
        ]
    directSortFrom srcName ascending =
      let
        descInitOp = if ascending then "<" else ">"
        descBreakOp = if ascending then ">=" else "<="
        ascBreakOp = if ascending then "<" else ">"
        sortCall = if ascending then
          GoCall (slices "Sort") [ GoVar "arraySort_output" ]
          else GoCall (slices "SortFunc")
            [ GoVar "arraySort_output"
            , GoFuncBlock [ Tuple "a" TypeInt64, Tuple "b" TypeInt64 ]
                [ GoIfElse (GoBinOp ">" (GoVar "a") (GoVar "b")) [ GoReturn (GoInt (-1)) ] []
                , GoIfElse (GoBinOp "<" (GoVar "a") (GoVar "b")) [ GoReturn (GoInt 1) ] []
                , GoReturn (GoInt 0)
                ] (TypeInterface "int")
            ]
      in
        [ GoAssign "arraySort_n" (GoCall (GoVar "len") [ rawGo srcName ])
        , GoAssign "arraySort_output" (GoCall (GoVar "make") [ rawGo "[]int64", GoVar "arraySort_n" ])
        , GoIfElse (GoBinOp "<" (GoVar "arraySort_n") (GoInt 2))
            [ GoCall (GoVar "copy") [ GoVar "arraySort_output", rawGo srcName ]
            , GoReturn (GoVar "arraySort_output")
            ] []
        , GoAssign "arraySort_prev" (GoIndex (rawGo srcName) (GoInt 1))
        , GoIfElse (GoBinOp descInitOp (GoVar "arraySort_prev") (GoIndex (rawGo srcName) (GoInt 0)))
            [ GoAssign "arraySort_monotone" (rawGo "true")
            , GoForRange ("_, arraySort_cur := range " <> srcName <> "[2:]")
                [ GoIfElse (GoBinOp descBreakOp (GoVar "arraySort_cur") (GoVar "arraySort_prev"))
                    [ GoMutate "arraySort_monotone" (rawGo "false"), rawGo "break" ] []
                , GoMutate "arraySort_prev" (GoVar "arraySort_cur")
                ]
            , GoIfElse (GoVar "arraySort_monotone")
                [ GoAssign "arraySort_last" (rawGo "arraySort_n - 1")
                , GoForRange ("arraySort_i, arraySort_v := range " <> srcName)
                    [ GoMutate "arraySort_output[arraySort_last-arraySort_i]" (GoVar "arraySort_v") ]
                , GoReturn (GoVar "arraySort_output")
                ] []
            ]
            [ GoAssign "arraySort_monotone" (rawGo "true")
            , GoForRange ("_, arraySort_cur := range " <> srcName <> "[2:]")
                [ GoIfElse (GoBinOp ascBreakOp (GoVar "arraySort_cur") (GoVar "arraySort_prev"))
                    [ GoMutate "arraySort_monotone" (rawGo "false"), rawGo "break" ] []
                , GoMutate "arraySort_prev" (GoVar "arraySort_cur")
                ]
            , GoIfElse (GoVar "arraySort_monotone")
                [ GoCall (GoVar "copy") [ GoVar "arraySort_output", rawGo srcName ]
                , GoReturn (GoVar "arraySort_output")
                ] []
            ]
        , GoCall (GoVar "copy") [ GoVar "arraySort_output", rawGo srcName ]
        , sortCall
        ]
    directSortBody ascending =
      let
        descInitOp = if ascending then "<" else ">"
        descBreakOp = if ascending then ">=" else "<="
        ascBreakOp = if ascending then "<" else ">"
        sortCall = if ascending then
          GoCall (slices "Sort") [ GoVar "arraySort_output" ]
          else GoCall (slices "SortFunc")
            [ GoVar "arraySort_output"
            , GoFuncBlock [ Tuple "a" TypeInt64, Tuple "b" TypeInt64 ]
                [ GoIfElse (GoBinOp ">" (GoVar "a") (GoVar "b")) [ GoReturn (GoInt (-1)) ] []
                , GoIfElse (GoBinOp "<" (GoVar "a") (GoVar "b")) [ GoReturn (GoInt 1) ] []
                , GoReturn (GoInt 0)
                ] (TypeInterface "int")
            ]
      in
        [ GoAssign "arraySort_prev" (at (GoInt 1))
        , GoIfElse (GoBinOp descInitOp (GoVar "arraySort_prev") (at (GoInt 0)))
            [ GoAssign "arraySort_monotone" (rawGo "true")
            , GoForRange "_, arraySort_cur := range arraySort_output[2:]"
                [ GoIfElse (GoBinOp descBreakOp (GoVar "arraySort_cur") (GoVar "arraySort_prev"))
                    [ GoMutate "arraySort_monotone" (rawGo "false"), rawGo "break" ] []
                , GoMutate "arraySort_prev" (GoVar "arraySort_cur")
                ]
            , GoIfElse (GoVar "arraySort_monotone")
                [ GoCall (slices "Reverse") [ GoVar "arraySort_output" ]
                , GoReturn (GoVar "arraySort_output")
                ] []
            ]
            [ GoAssign "arraySort_monotone" (rawGo "true")
            , GoForRange "_, arraySort_cur := range arraySort_output[2:]"
                [ GoIfElse (GoBinOp ascBreakOp (GoVar "arraySort_cur") (GoVar "arraySort_prev"))
                    [ GoMutate "arraySort_monotone" (rawGo "false"), rawGo "break" ] []
                , GoMutate "arraySort_prev" (GoVar "arraySort_cur")
                ]
            , GoIfElse (GoVar "arraySort_monotone")
                [ GoReturn (GoVar "arraySort_output") ] []
            ]
        , sortCall
        ]
    sortBody = case staticConversion of
      Just { lt, eq, gt }
        | lt < 0 && eq == 0 && gt > 0 -> directSortBody true
        | lt > 0 && eq == 0 && gt < 0 -> directSortBody false
        | otherwise -> dynamicSortBody (GoCall (GoVar "int64") [ GoInt lt ]) (GoCall (GoVar "int64") [ GoInt eq ]) (GoCall (GoVar "int64") [ GoInt gt ])
      Nothing -> dynamicSortBody (converted "LT") (converted "EQ") (converted "GT")
    directOrder = case staticConversion of
      Just { lt, eq, gt }
        | lt < 0 && eq == 0 && gt > 0 -> Just true
        | lt > 0 && eq == 0 && gt < 0 -> Just false
      _ -> Nothing
    body = case argumentType 2, directOrder of
      TypeNativeArray TypeInt64, Just ascending ->
        [ GoAssign "arraySort_input" (argument 2), source.assignment ] <> directSortFrom source.target ascending
      TypeNativeArray _, _ ->
        [ GoAssign "arraySort_input" (argument 2), source.assignment
        , GoAssign "arraySort_output" (GoCall (GoVar "make") [ rawGo "[]int64", GoCall (GoVar "len") [ rawGo source.target ] ])
        , copyOutput
        , GoIfElse (GoBinOp "<" length (GoInt 2)) [ GoReturn (GoVar "arraySort_output") ] []
        ] <> sortBody
      _, Just ascending ->
        [ GoAssign "arraySort_input" (argument 2)
        , GoIfElse (rawGo "arraySort_input.Type == gopurs_runtime.TypeIntArray")
            ([ GoAssign "arraySort_ints" (rawGo "*(*[]int64)(arraySort_input.UnsafePtr)") ]
              <> directSortFrom "arraySort_ints" ascending
              <> [ GoReturn (GoVar "arraySort_output") ])
            []
        , GoAssign "arraySort_values" (rawGo "*(*[]gopurs_runtime.Value)(arraySort_input.UnsafePtr)")
        , GoAssign "arraySort_output" (GoCall (GoVar "make") [ rawGo "[]int64", GoCall (GoVar "len") [ GoVar "arraySort_values" ] ])
        , GoForRange "index, value := range arraySort_values"
            [ GoMutate "arraySort_output[index]" (GoSelector (GoVar "value") "IntVal") ]
        , GoIfElse (GoBinOp "<" length (GoInt 2)) [ GoReturn (GoVar "arraySort_output") ] []
        ] <> directSortBody ascending
      _, Nothing ->
        [ GoAssign "arraySort_input" (argument 2)
        , rawGo "var arraySort_output []int64"
        , GoIfElse (rawGo "arraySort_input.Type == gopurs_runtime.TypeIntArray")
            [ GoAssign "arraySort_ints" (rawGo "*(*[]int64)(arraySort_input.UnsafePtr)")
            , GoMutate "arraySort_output" (GoCall (GoVar "make") [ rawGo "[]int64", GoCall (GoVar "len") [ GoVar "arraySort_ints" ] ])
            , GoCall (GoVar "copy") [ GoVar "arraySort_output", GoVar "arraySort_ints" ]
            ]
            [ GoAssign "arraySort_values" (rawGo "*(*[]gopurs_runtime.Value)(arraySort_input.UnsafePtr)")
            , GoMutate "arraySort_output" (GoCall (GoVar "make") [ rawGo "[]int64", GoCall (GoVar "len") [ GoVar "arraySort_values" ] ])
            , GoForRange "index, value := range arraySort_values"
                [ GoMutate "arraySort_output[index]" (GoSelector (GoVar "value") "IntVal") ]
            ]
        , GoIfElse (GoBinOp "<" length (GoInt 2)) [ GoReturn (GoVar "arraySort_output") ] []
        ] <> sortBody
  pure { stmts: captured.stmts
    , expr: GoCall (GoFuncLit [] body (GoVar "arraySort_output") (TypeNativeArray TypeInt64)) []
    , exprType: TypeNativeArray TypeInt64, nextId: captured.nextId }
  where
  slices name = GoRaw { text: "slices." <> name, imports: [ "slices" ] }

staticOrderingConversion :: TcoExpr -> Maybe { lt :: Int, eq :: Int, gt :: Int }
staticOrderingConversion expression = case unwrapTcoExpr expression of
  Abs params body | NEA.length params == 1 ->
    let Tuple _ paramLevel = NEA.head params
    in case unwrapTcoExpr body of
      Branch pairs otherwise -> do
        collected <- foldM (collectBranch paramLevel) { lt: Nothing, eq: Nothing, gt: Nothing } (NEA.toArray pairs)
        let fallback = case unwrapTcoExpr otherwise of
              Lit (LitInt n) -> Just n
              _ -> Nothing
            lt = case collected.lt of
              Just v -> Just v
              Nothing -> fallback
            eq = case collected.eq of
              Just v -> Just v
              Nothing -> fallback
            gt = case collected.gt of
              Just v -> Just v
              Nothing -> fallback
        { lt: _, eq: _, gt: _ } <$> lt <*> eq <*> gt
      _ -> Nothing
  _ -> Nothing
  where
  collectBranch paramLevel acc (Pair cond result) = do
    tag <- case unwrapTcoExpr cond of
      PrimOp (Op1 (OpIsTag (Qualified (Just (ModuleName "Data.Ordering")) (Ident name))) arg) ->
        case unwrapTcoExpr arg of
          Local _ level | level == paramLevel -> Just name
          _ -> Nothing
      _ -> Nothing
    value <- case unwrapTcoExpr result of
      Lit (LitInt n) -> Just n
      _ -> Nothing
    case tag of
      "LT" -> Just (acc { lt = case acc.lt of
        Just _ -> acc.lt
        Nothing -> Just value })
      "EQ" -> Just (acc { eq = case acc.eq of
        Just _ -> acc.eq
        Nothing -> Just value })
      "GT" -> Just (acc { gt = case acc.gt of
        Just _ -> acc.gt
        Nothing -> Just value })
      _ -> Nothing

intComparator :: TcoExpr -> Boolean
intComparator expression = case unwrapTcoExpr expression of
  Accessor dictionary (GetProp "compare") -> named "Data.Ord" "ordInt" dictionary
  _ ->
    let Tuple fn spine = collectGoSpine expression
    in named "Data.Ord" "compare" fn && case getGoSpineArgs spine of
      [ dictionary ] -> named "Data.Ord" "ordInt" dictionary
      _ -> false

named :: String -> String -> TcoExpr -> Boolean
named moduleName name expression = case qualifiedTarget expression of
  Just target -> target.mbMod == Just (ModuleName moduleName) && target.name == name
  Nothing -> false
