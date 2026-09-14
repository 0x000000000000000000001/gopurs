module Gopurs.CallExprs
  ( application
  , uncurriedApplication
  , effectApplication
  ) where

import Prelude
import Data.Array as Array

import Data.Foldable (foldl, foldMap)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Newtype (unwrap)
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Tuple (Tuple(..))
import Gopurs.CallAnalysis (collectGoSpine, getGoSpineArgs)
import Gopurs.ExprAnalysis (extractFuncType, getExprType, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, TranslateExpr, StmtTree(..))
import Gopurs.GoAst (GoExpr(..), GoType(..), goTypeToStr, sanitizeName)
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr, unboxGoExpr)
import Gopurs.GoTypes (exprTypeToGoType)

import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Ident(..), Literal(..), ModuleName(..), Qualified(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..))

application :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> ExprResult
application translate context@{ metadata, codegenStateRef, depth, modNameStr, moduleFunctions, bound, loopCtx, mbExpectedExprType, options: { isTail } } nextId tcoExpr =
  let
    Tuple flatFn flatArgsSpine = collectGoSpine tcoExpr
    flatArgs = getGoSpineArgs flatArgsSpine

    isTailCallTo =
      if isTail then case unwrapTcoExpr flatFn of
        Local mbIdent lvl ->
          let
            v = fromMaybe { name: localId mbIdent lvl, goType: TypeValue } (Map.lookup (localId mbIdent lvl) bound)
          in
            Array.findIndex (\ctx -> ctx.ident == v.name) loopCtx
        Var (Qualified _ (Ident name)) ->
          let
            fullName = sanitizeName name
          in
            Array.findIndex (\ctx -> ctx.ident == fullName) loopCtx
        _ -> Nothing
      else Nothing

  in
    case isTailCallTo of
      Just index ->
        let
          accFinal = foldl
            ( \acc arg ->
                let
                  argRes = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing }) acc.nextId arg
                in
                  { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs argRes.expr, exprTypes: Array.snoc acc.exprTypes argRes.exprType, nextId: argRes.nextId }
            )
            { stmts: StmtEmpty, exprs: [], exprTypes: [], nextId }
            flatArgs
          targetCtx = fromMaybe { ident: "", params: [], loopParams: [], goTypes: [], fRet: TypeValue } (Array.index loopCtx index)
          assigns = Array.mapWithIndex
            ( \i paramName ->
                let
                  argExpr = fromMaybe (GoRaw "nil") (Array.index accFinal.exprs i)
                  argType = fromMaybe TypeValue (Array.index accFinal.exprTypes i)
                  expectedType = fromMaybe TypeValue (Array.index targetCtx.goTypes i)
                in
                  GoMutate paramName (coerceGoExpr codegenStateRef modNameStr argExpr argType expectedType)
            )
            targetCtx.loopParams
        in
          let
            expectedGoType = exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr
              ( case getExprType tcoExpr of
                  Any -> fromMaybe Any mbExpectedExprType
                  ty -> ty
              )
            expectedGoTypeStr = goTypeToStr expectedGoType
          in
            { stmts: accFinal.stmts <> foldMap StmtLeaf assigns <> StmtLeaf (GoContinue targetCtx.ident), expr: GoRaw ("func() " <> expectedGoTypeStr <> " { panic(\"unreachable\") }()"), exprType: expectedGoType, nextId: accFinal.nextId }

      Nothing ->
        let
          getVar :: BackendSyntax TcoExpr -> Maybe { mbMod :: Maybe ModuleName, name :: String }
          getVar (Typed _ inner) = getVar (unwrapTcoExpr inner)
          getVar (Var (Qualified mbMod (Ident name))) = Just { mbMod, name }
          getVar (Local mbIdent lvl) =
            let
              resolvedName = (fromMaybe { name: localId mbIdent lvl, goType: TypeValue } (Map.lookup (localId mbIdent lvl) bound)).name
            in
              Just { mbMod: Nothing, name: resolvedName }
          getVar (Lit _) = Just { mbMod: Nothing, name: "Lit" }
          getVar (App _ _) = Just { mbMod: Nothing, name: "App" }
          getVar (Abs _ _) = Just { mbMod: Nothing, name: "Abs" }
          getVar (UncurriedApp _ _) = Just { mbMod: Nothing, name: "UncurriedApp" }
          getVar (UncurriedAbs _ _) = Just { mbMod: Nothing, name: "UncurriedAbs" }
          getVar (UncurriedEffectApp _ _) = Just { mbMod: Nothing, name: "UncurriedEffectApp" }
          getVar (UncurriedEffectAbs _ _) = Just { mbMod: Nothing, name: "UncurriedEffectAbs" }
          getVar (Accessor _ _) = Just { mbMod: Nothing, name: "Accessor" }
          getVar (Update _ _) = Just { mbMod: Nothing, name: "Update" }
          getVar (CtorSaturated _ _ _ _ _) = Just { mbMod: Nothing, name: "CtorSaturated" }
          getVar (CtorDef _ _ _ _) = Just { mbMod: Nothing, name: "CtorDef" }
          getVar (LetRec _ _ _) = Just { mbMod: Nothing, name: "LetRec" }
          getVar (Let _ _ _ _) = Just { mbMod: Nothing, name: "Let" }
          getVar (EffectBind _ _ _ _) = Just { mbMod: Nothing, name: "EffectBind" }
          getVar (EffectPure _) = Just { mbMod: Nothing, name: "EffectPure" }
          getVar (EffectDefer _) = Just { mbMod: Nothing, name: "EffectDefer" }
          getVar _ = Just { mbMod: Nothing, name: "Unknown" }

          mbIntrinsic = case getVar (unwrapTcoExpr flatFn) of
            Just { name: "arrayMap" } ->
              if Array.length flatArgs >= 2 then Just "arrayMap" else Nothing
            Just { name: "foldlArray" } ->
              if Array.length flatArgs >= 3 then Just "foldlArray" else Nothing
            Just { mbMod, name: "filter" } | mbMod == Just (ModuleName "Data.Array") || (mbMod == Nothing && modNameStr == "Data.Array") ->
              if Array.length flatArgs >= 2 then Just "filter" else Nothing
            _ -> Nothing

          mbDirectCall = case getVar (unwrapTcoExpr flatFn) of
            Just { mbMod, name } ->
              let
                isLocal = map (String.replaceAll (Pattern ".") (Replacement "_") <<< unwrap) mbMod == Just modNameStr || mbMod == Nothing
                entry = if isLocal then Map.lookup name moduleFunctions
                  else do
                    mn <- mbMod
                    Map.lookup (unwrap mn <> "." <> name) metadata.globalFunctions
              in
                case entry of
                  Just e ->
                    if Array.length flatArgs >= e.arity && e.arity >= 1 then Just e else Nothing
                  Nothing -> Nothing
            Nothing -> Nothing
        in
          case mbIntrinsic of
            Just intrinsicName ->
              let
                accArgs = foldl
                  ( \acc arg ->
                      let
                        argRes = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing }) acc.nextId arg
                      in
                        { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs (boxGoExpr codegenStateRef modNameStr argRes.expr argRes.exprType), exprTypes: Array.snoc acc.exprTypes argRes.exprType, nextId: argRes.nextId }
                  )
                  { stmts: StmtEmpty, exprs: [], exprTypes: [], nextId }
                  flatArgs

                iifeName = intrinsicName <> show depth
                arrValName = "arr_val_" <> iifeName
                arrGoName = "arr_go_" <> iifeName
                resGoName = "res_go_" <> iifeName
                iName = "i_" <> iifeName
                vName = "v_" <> iifeName

                iifeExpr = case intrinsicName of
                  "arrayMap" ->
                    let
                      fExpr = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 0)
                      arrExpr = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 1)
                      loopBody = GoMutate (resGoName <> "[" <> iName <> "]") (GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ fExpr, GoVar vName ])
                      iifeBody = GoBlock
                        [ GoAssign arrGoName (GoCall (GoRaw "(*[]gopurs_runtime.Value)") [ GoSelector (GoVar arrValName) "UnsafePtr" ])
                        , GoAssign resGoName (GoCall (GoVar "make") [ GoRaw "[]gopurs_runtime.Value", GoCall (GoVar "len") [ GoRaw ("*" <> arrGoName) ] ])
                        , GoForRange (iName <> ", " <> vName <> " := range *" <> arrGoName) [ loopBody ]
                        , GoReturn (GoCall (GoSelector (GoVar "gopurs_runtime") "Array") [ GoVar resGoName ])
                        ]
                    in
                      GoIIFE arrValName arrExpr iifeBody

                  "foldlArray" ->
                    let
                      fExpr = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 0)
                      initExpr = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 1)
                      boxedArrExpr = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 2)
                      arrExpr =
                        if isIntArrayFold flatFn flatArgs then
                          normalizeFreshIntArrayRoundtrip (iifeName <> "_" <> show accArgs.nextId) boxedArrExpr
                        else boxedArrExpr
                      loopBody = GoMutate resGoName (GoCall (GoSelector (GoVar "gopurs_runtime") "Apply2") [ fExpr, GoVar resGoName, GoVar vName ])
                      iifeBody = GoBlock
                        [ GoAssign resGoName initExpr
                        , GoAssign arrGoName (GoCall (GoRaw "(*[]gopurs_runtime.Value)") [ GoSelector (GoVar arrValName) "UnsafePtr" ])
                        , GoForRange ("_, " <> vName <> " := range *" <> arrGoName) [ loopBody ]
                        , GoReturn (GoVar resGoName)
                        ]
                    in
                      GoIIFE arrValName arrExpr iifeBody

                  "filter" ->
                    let
                      fExpr = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 0)
                      arrExpr = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 1)
                      condExpr = GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ fExpr, GoVar vName ]
                      isTrueExpr = GoCall (GoSelector condExpr "BoolVal") []
                      loopBody = GoIfElse isTrueExpr [ GoMutate resGoName (GoCall (GoVar "append") [ GoVar resGoName, GoVar vName ]) ] []
                      iifeBody = GoBlock
                        [ GoAssign arrGoName (GoCall (GoRaw "(*[]gopurs_runtime.Value)") [ GoSelector (GoVar arrValName) "UnsafePtr" ])
                        , GoAssign resGoName (GoCall (GoVar "make") [ GoRaw "[]gopurs_runtime.Value", GoRaw "0" ])
                        , GoForRange ("_, " <> vName <> " := range *" <> arrGoName) [ loopBody ]
                        , GoReturn (GoCall (GoSelector (GoVar "gopurs_runtime") "Array") [ GoVar resGoName ])
                        ]
                    in
                      GoIIFE arrValName arrExpr iifeBody

                  _ -> GoRaw "nil"

                arity = if intrinsicName == "foldlArray" then 3 else 2
                accArgsRemaining = Array.drop arity accArgs.exprs
                accArgsRemainingTypes = Array.drop arity accArgs.exprTypes
                accArgsRemainingBoxed = Array.zipWith (\arg t -> boxGoExpr codegenStateRef modNameStr arg t) accArgsRemaining accArgsRemainingTypes

                finalExpr = applyBoxed iifeExpr accArgsRemainingBoxed
              in
                { stmts: accArgs.stmts, expr: finalExpr, exprType: TypeValue, nextId: accArgs.nextId }

            Nothing ->
              case mbDirectCall of
                Just { fullName, fArgs, fRet, arity } ->
                  let
                    accArgs = foldl
                      ( \acc arg ->
                          let
                            argRes = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing }) acc.nextId arg
                          in
                            { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs argRes.expr, exprTypes: Array.snoc acc.exprTypes argRes.exprType, nextId: argRes.nextId }
                      )
                      { stmts: StmtEmpty, exprs: [], exprTypes: [], nextId }
                      flatArgs

                    accArgsArity = Array.take arity accArgs.exprs
                    accArgsRemaining = Array.drop arity accArgs.exprs
                    accArgsRemainingTypes = Array.drop arity accArgs.exprTypes
                    accArgsRemainingBoxed = Array.zipWith (\arg t -> boxGoExpr codegenStateRef modNameStr arg t) accArgsRemaining accArgsRemainingTypes

                    callArgs = Array.mapWithIndex
                      ( \i argExprValue ->
                          let
                            expectedType = fromMaybe TypeValue (Array.index fArgs i)
                            actualType = fromMaybe TypeValue (Array.index accArgs.exprTypes i)
                          in
                            coerceGoExpr codegenStateRef modNameStr argExprValue actualType expectedType
                      )
                      accArgsArity

                    callExpr = GoCall (GoVar fullName) callArgs
                    finalExpr = if Array.length accArgsRemainingBoxed == 0 then callExpr else applyBoxed (boxGoExpr codegenStateRef modNameStr callExpr fRet) accArgsRemainingBoxed
                    finalExprType = if Array.length accArgsRemainingBoxed == 0 then fRet else TypeValue
                  in
                    { stmts: accArgs.stmts, expr: finalExpr, exprType: finalExprType, nextId: accArgs.nextId }

                Nothing ->
                  let
                    resFn = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing }) nextId flatFn
                    -- A dynamic call consumes boxed arguments. Keep
                    -- an existing record value instead of unboxing
                    -- and immediately rebuilding its representation.
                    argExpectedType = case resFn.exprType of
                      TypeFunc _ _ -> Nothing
                      _ -> Just Any
                    accArgs = foldl
                      ( \acc arg ->
                          let
                            argRes = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = argExpectedType }) acc.nextId arg
                          in
                            { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs argRes.expr, exprTypes: Array.snoc acc.exprTypes argRes.exprType, nextId: argRes.nextId }
                      )
                      { stmts: resFn.stmts, exprs: [], exprTypes: [], nextId: resFn.nextId }
                      flatArgs

                    finalExprType = case resFn.exprType of
                      TypeFunc fArgs fRet | Array.length fArgs == Array.length flatArgs -> fRet
                      _ -> TypeValue

                    finalExpr = case resFn.exprType of
                      TypeFunc fArgs _ | Array.length fArgs == Array.length flatArgs ->
                        let
                          callArgs = Array.mapWithIndex
                            ( \i expected ->
                                let
                                  arg = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs i)
                                  actual = fromMaybe TypeValue (Array.index accArgs.exprTypes i)
                                in
                                  coerceGoExpr codegenStateRef modNameStr arg actual expected
                            )
                            fArgs
                        in
                          GoCall resFn.expr callArgs
                      TypeFunc fArgs fRet | Array.length flatArgs > Array.length fArgs ->
                        let
                          arity = Array.length fArgs
                          accArgsArity = Array.take arity accArgs.exprs
                          accArgsRemaining = Array.drop arity accArgs.exprs
                          accArgsRemainingTypes = Array.drop arity accArgs.exprTypes
                          accArgsRemainingBoxed = Array.zipWith (\arg t -> boxGoExpr codegenStateRef modNameStr arg t) accArgsRemaining accArgsRemainingTypes

                          callArgs = Array.mapWithIndex
                            ( \i argExprValue ->
                                let
                                  expectedType = fromMaybe TypeValue (Array.index fArgs i)
                                  actualType = fromMaybe TypeValue (Array.index accArgs.exprTypes i)
                                in
                                  coerceGoExpr codegenStateRef modNameStr argExprValue actualType expectedType
                            )
                            accArgsArity

                          callExpr = GoCall resFn.expr callArgs
                        in
                          applyBoxed (boxGoExpr codegenStateRef modNameStr callExpr fRet) accArgsRemainingBoxed
                      _ ->
                        let
                          boxedArgs = Array.zipWith (\arg actual -> boxGoExpr codegenStateRef modNameStr arg actual) accArgs.exprs accArgs.exprTypes
                        in
                          applyBoxed (boxGoExpr codegenStateRef modNameStr resFn.expr resFn.exprType) boxedArgs
                  in
                    { stmts: accArgs.stmts, expr: finalExpr, exprType: finalExprType, nextId: accArgs.nextId }

uncurriedApplication :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> TcoExpr -> Array TcoExpr -> ExprResult
uncurriedApplication translate context@{ metadata, codegenStateRef, depth, modNameStr, moduleFunctions, bound, loopCtx, mbExpectedExprType, options: { isTail } } nextId tcoExpr fn args =
  let
    getVar :: BackendSyntax TcoExpr -> Maybe { mbMod :: Maybe ModuleName, name :: String }
    getVar (Typed _ inner) = getVar (unwrapTcoExpr inner)
    getVar (Var (Qualified mbMod (Ident name))) = Just { mbMod, name }
    getVar _ = Nothing

    mbIntrinsic = case getVar (unwrapTcoExpr fn) of
      Just { name: "arrayMap" } ->
        if Array.length args >= 2 then Just "arrayMap" else Nothing
      Just { name: "foldlArray" } ->
        if Array.length args >= 3 then Just "foldlArray" else Nothing
      Just { mbMod, name: "filterImpl" } | mbMod == Just (ModuleName "Data.Array") || (mbMod == Nothing && modNameStr == "Data.Array") ->
        if Array.length args >= 2 then Just "filterImpl" else Nothing
      _ -> Nothing
  in
    case mbIntrinsic of
      Just intrinsicName ->
        let
          accArgs = foldl
            ( \acc arg ->
                let
                  argRes = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing }) acc.nextId arg
                in
                  { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs argRes.expr, exprTypes: Array.snoc acc.exprTypes argRes.exprType, nextId: argRes.nextId }
            )
            { stmts: StmtEmpty, exprs: [], exprTypes: [], nextId }
            args

          iifeName = intrinsicName <> show depth
          arrValName = "arr_val_" <> iifeName
          arrGoName = "arr_go_" <> iifeName
          resGoName = "res_go_" <> iifeName
          iName = "i_" <> iifeName
          vName = "v_" <> iifeName

          iifeExpr = case intrinsicName of
            "arrayMap" ->
              let
                fExprRaw = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 0)
                fExprType = fromMaybe TypeValue (Array.index accArgs.exprTypes 0)
                arrExprRaw = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 1)
                arrExprType = fromMaybe TypeValue (Array.index accArgs.exprTypes 1)

                mbFnVar = case Array.index args 0 of
                  Just fArg -> getVar (unwrapTcoExpr fArg)
                  Nothing -> Nothing

                fnFullName = case mbFnVar of
                  Just { mbMod: Just (ModuleName mn), name } -> String.replaceAll (Pattern ".") (Replacement "_") mn <> "." <> name
                  Just { mbMod: Nothing, name } -> modNameStr <> "." <> name
                  Nothing -> ""

                mbFnArityInfo = Map.lookup fnFullName moduleFunctions

                elemType = case arrExprType of
                  TypeNativeArray inner -> inner
                  _ -> TypeValue

                retType = case fExprType of
                  TypeFunc _ ret -> ret
                  _ -> case mbFnArityInfo of
                    Just info -> info.fRet
                    _ -> TypeValue

                finalRetType = TypeNativeArray retType

                arrGoAssignment = case arrExprType of
                  TypeNativeArray _ -> GoAssign arrGoName (GoVar arrValName)
                  _ -> GoAssign arrGoName (GoCall (GoRaw "(*[]gopurs_runtime.Value)") [ GoSelector (GoVar arrValName) "UnsafePtr" ])

                arrGoRangeTarget = case arrExprType of
                  TypeNativeArray _ -> arrGoName
                  _ -> "*" <> arrGoName

                loopBody = case mbFnArityInfo of
                  Just info | info.arity == 1 ->
                    let
                      expectedArgType = fromMaybe TypeValue (Array.index info.fArgs 0)
                    in
                      GoMutate (resGoName <> "[" <> iName <> "]") (unboxGoExpr codegenStateRef modNameStr (GoCall (GoVar ("Call_" <> String.replaceAll (Pattern ".") (Replacement "_") fnFullName)) [ unboxGoExpr codegenStateRef modNameStr (GoVar vName) elemType expectedArgType ]) info.fRet retType)
                  _ ->
                    GoMutate (resGoName <> "[" <> iName <> "]") (unboxGoExpr codegenStateRef modNameStr (GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ boxGoExpr codegenStateRef modNameStr fExprRaw fExprType, boxGoExpr codegenStateRef modNameStr (GoVar vName) elemType ]) TypeValue retType)

                iifeBodyStmts =
                  [ arrGoAssignment
                  , GoAssign resGoName (GoCall (GoVar "make") [ GoRaw ("[]" <> goTypeToStr retType), GoCall (GoVar "len") [ GoRaw arrGoRangeTarget ] ])
                  , GoForRange (iName <> ", " <> vName <> " := range " <> arrGoRangeTarget) [ loopBody ]
                  ]
              in
                { stmts: accArgs.stmts, expr: GoCall (GoFuncLit [] (Array.cons (GoAssign arrValName arrExprRaw) (Array.cons (GoMutate "_" (GoVar arrValName)) iifeBodyStmts)) (GoVar resGoName) finalRetType) [], exprType: finalRetType, nextId: accArgs.nextId }

            "foldlArray" ->
              let
                fExpr = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 0)
                fExprType = fromMaybe TypeValue (Array.index accArgs.exprTypes 0)
                initExpr = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 1)
                initExprType = fromMaybe TypeValue (Array.index accArgs.exprTypes 1)
                arrExpr = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 2)
                arrExprType = fromMaybe TypeValue (Array.index accArgs.exprTypes 2)

                mbFnVar = case Array.index args 0 of
                  Just fArg -> getVar (unwrapTcoExpr fArg)
                  Nothing -> Nothing

                fnFullName = case mbFnVar of
                  Just { mbMod: Just (ModuleName mn), name } -> String.replaceAll (Pattern ".") (Replacement "_") mn <> "." <> name
                  Just { mbMod: Nothing, name } -> modNameStr <> "." <> name
                  Nothing -> ""

                mbFnArityInfo = Map.lookup fnFullName moduleFunctions

                elemType = case arrExprType of
                  TypeNativeArray inner -> inner
                  _ -> TypeValue

                loopBody = case mbFnArityInfo of
                  Just info | info.arity == 2 ->
                    let
                      expectedArg0 = fromMaybe TypeValue (Array.index info.fArgs 0)
                      expectedArg1 = fromMaybe TypeValue (Array.index info.fArgs 1)
                    in
                      GoMutate resGoName (unboxGoExpr codegenStateRef modNameStr (GoCall (GoVar ("Call_" <> String.replaceAll (Pattern ".") (Replacement "_") fnFullName)) [ unboxGoExpr codegenStateRef modNameStr (GoVar resGoName) initExprType expectedArg0, unboxGoExpr codegenStateRef modNameStr (GoVar vName) elemType expectedArg1 ]) info.fRet initExprType)
                  _ ->
                    GoMutate resGoName (unboxGoExpr codegenStateRef modNameStr (GoCall (GoSelector (GoVar "gopurs_runtime") "Apply2") [ boxGoExpr codegenStateRef modNameStr fExpr fExprType, boxGoExpr codegenStateRef modNameStr (GoVar resGoName) initExprType, boxGoExpr codegenStateRef modNameStr (GoVar vName) elemType ]) TypeValue initExprType)

                arrGoAssignment = case arrExprType of
                  TypeNativeArray _ -> GoAssign arrGoName (GoVar arrValName)
                  _ -> GoAssign arrGoName (GoCall (GoRaw "(*[]gopurs_runtime.Value)") [ GoSelector (GoVar arrValName) "UnsafePtr" ])

                arrGoRangeTarget = case arrExprType of
                  TypeNativeArray _ -> arrGoName
                  _ -> "*" <> arrGoName

                iifeBody = GoBlock
                  [ GoAssign resGoName initExpr
                  , arrGoAssignment
                  , GoForRange ("_, " <> vName <> " := range " <> arrGoRangeTarget) [ loopBody ]
                  , GoReturn (GoVar resGoName)
                  ]
              in
                { stmts: accArgs.stmts, expr: GoIIFE arrValName arrExpr iifeBody, exprType: initExprType, nextId: accArgs.nextId }

            "filterImpl" ->
              let
                fExprRaw = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 0)
                fExprType = fromMaybe TypeValue (Array.index accArgs.exprTypes 0)
                arrExprRaw = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs 1)
                arrExprType = fromMaybe TypeValue (Array.index accArgs.exprTypes 1)

                mbFnVar = case Array.index args 0 of
                  Just fArg -> getVar (unwrapTcoExpr fArg)
                  Nothing -> Nothing

                fnFullName = case mbFnVar of
                  Just { mbMod: Just (ModuleName mn), name } -> String.replaceAll (Pattern ".") (Replacement "_") mn <> "." <> name
                  Just { mbMod: Nothing, name } -> modNameStr <> "." <> name
                  Nothing -> ""

                mbFnArityInfo = Map.lookup fnFullName moduleFunctions

                elemType = case arrExprType of
                  TypeNativeArray inner -> inner
                  _ -> TypeValue

                isTrueExpr = case mbFnArityInfo of
                  Just info | info.arity == 1 ->
                    let
                      expectedArgType = fromMaybe TypeValue (Array.index info.fArgs 0)
                    in
                      GoCall (GoVar ("Call_" <> String.replaceAll (Pattern ".") (Replacement "_") fnFullName)) [ unboxGoExpr codegenStateRef modNameStr (GoVar vName) elemType expectedArgType ]
                  _ ->
                    let
                      condExpr = GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ boxGoExpr codegenStateRef modNameStr fExprRaw fExprType, boxGoExpr codegenStateRef modNameStr (GoVar vName) elemType ]
                    in
                      GoCall (GoSelector condExpr "BoolVal") []

                loopBody = GoIfElse isTrueExpr [ GoMutate resGoName (GoCall (GoVar "append") [ GoVar resGoName, GoVar vName ]) ] []

                arrGoAssignment = case arrExprType of
                  TypeNativeArray _ -> GoAssign arrGoName (GoVar arrValName)
                  _ -> GoAssign arrGoName (GoCall (GoRaw "(*[]gopurs_runtime.Value)") [ GoSelector (GoVar arrValName) "UnsafePtr" ])

                arrGoRangeTarget = case arrExprType of
                  TypeNativeArray _ -> arrGoName
                  _ -> "*" <> arrGoName

                iifeBodyStmts =
                  [ arrGoAssignment
                  , GoAssign resGoName (GoCall (GoVar "make") [ GoRaw ("[]" <> goTypeToStr elemType), GoRaw "0" ])
                  , GoForRange ("_, " <> vName <> " := range " <> arrGoRangeTarget) [ loopBody ]
                  ]
                filterExpr = GoCall (GoFuncLit [] (Array.cons (GoAssign arrValName arrExprRaw) (Array.cons (GoMutate "_" (GoVar arrValName)) iifeBodyStmts)) (GoVar resGoName) (TypeNativeArray elemType)) []
                freshFilterExpr = if elemType == TypeValue && Array.length args == 2 then GoFreshFilterArray filterExpr else filterExpr
              in
                { stmts: accArgs.stmts, expr: freshFilterExpr, exprType: TypeNativeArray elemType, nextId: accArgs.nextId }

            _ -> { stmts: accArgs.stmts, expr: GoRaw "nil", exprType: TypeValue, nextId: accArgs.nextId }
        in
          iifeExpr
      Nothing ->
        let
          Tuple flatFn flatArgsSpine = collectGoSpine tcoExpr
          flatArgs = getGoSpineArgs flatArgsSpine
          isTailCallTo =
            if isTail then case unwrapTcoExpr flatFn of
              Local mbIdent lvl ->
                let
                  v = fromMaybe { name: localId mbIdent lvl, goType: TypeValue } (Map.lookup (localId mbIdent lvl) bound)
                in
                  Array.findIndex (\ctx -> ctx.ident == v.name) loopCtx
              Var (Qualified _ (Ident name)) ->
                let
                  fullName = sanitizeName name
                in
                  Array.findIndex (\ctx -> ctx.ident == fullName) loopCtx
              _ -> Nothing
            else Nothing
        in
          case isTailCallTo of
            Just index ->
              let
                accFinal = foldl
                  ( \acc arg ->
                      let
                        argRes = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing }) acc.nextId arg
                      in
                        { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs argRes.expr, exprTypes: Array.snoc acc.exprTypes argRes.exprType, nextId: argRes.nextId }
                  )
                  { stmts: StmtEmpty, exprs: [], exprTypes: [], nextId }
                  flatArgs
                targetCtx = fromMaybe { ident: "", params: [], loopParams: [], goTypes: [], fRet: TypeValue } (Array.index loopCtx index)
                assigns = Array.mapWithIndex
                  ( \i paramName ->
                      let
                        argExpr = fromMaybe (GoRaw "nil") (Array.index accFinal.exprs i)
                        argType = fromMaybe TypeValue (Array.index accFinal.exprTypes i)
                        expectedType = fromMaybe TypeValue (Array.index targetCtx.goTypes i)
                      in
                        GoMutate paramName (coerceGoExpr codegenStateRef modNameStr argExpr argType expectedType)
                  )
                  targetCtx.loopParams
              in
                let
                  expectedGoType = exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr
                    ( case getExprType tcoExpr of
                        Any -> fromMaybe Any mbExpectedExprType
                        ty -> ty
                    )
                  expectedGoTypeStr = goTypeToStr expectedGoType
                in
                  { stmts: accFinal.stmts <> foldMap StmtLeaf assigns <> StmtLeaf (GoContinue targetCtx.ident), expr: GoRaw ("func() " <> expectedGoTypeStr <> " { panic(\"unreachable\") }()"), exprType: expectedGoType, nextId: accFinal.nextId }
            Nothing ->
              let
                resFn = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing }) nextId fn
                accArgs = foldl
                  ( \acc arg ->
                      let
                        argRes = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing }) acc.nextId arg
                      in
                        { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs argRes.expr, exprTypes: Array.snoc acc.exprTypes argRes.exprType, nextId: argRes.nextId }
                  )
                  { stmts: resFn.stmts, exprs: [], exprTypes: [], nextId: resFn.nextId }
                  args
                len = Array.length args
                goFuncName = if len >= 2 && len <= 10 then "UncurriedApp" <> show len else "UncurriedApp"
              in
                case resFn.exprType of
                  TypeFunc fArgs fRet | Array.length fArgs == len ->
                    let
                      callArgs = Array.mapWithIndex
                        ( \i expected ->
                            let
                              arg = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs i)
                              actual = fromMaybe TypeValue (Array.index accArgs.exprTypes i)
                            in
                              coerceGoExpr codegenStateRef modNameStr arg actual expected
                        )
                        fArgs
                    in
                      { stmts: accArgs.stmts, expr: boxGoExpr codegenStateRef modNameStr (GoCall resFn.expr callArgs) fRet, exprType: TypeValue, nextId: accArgs.nextId }
                  _ ->
                    let
                      boxedArgs = Array.zipWith (\arg actual -> boxGoExpr codegenStateRef modNameStr arg actual) accArgs.exprs accArgs.exprTypes
                    in
                      { stmts: accArgs.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") goFuncName) (Array.cons (boxGoExpr codegenStateRef modNameStr resFn.expr resFn.exprType) boxedArgs), exprType: TypeValue, nextId: accArgs.nextId }

effectApplication :: TranslateExpr -> ExprContext -> Int -> TcoExpr -> Array TcoExpr -> ExprResult
effectApplication translate context@{ codegenStateRef, depth, modNameStr } nextId fn args =
  let
    resFn = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing }) nextId fn
    accArgs = foldl
      ( \acc arg ->
          let
            argRes = translate (context { depth = (depth + 1), tcoIdent = Nothing, loopCtx = [], options = { isTail: false, inEffectBlock: false }, mbExpectedExprType = Nothing }) acc.nextId arg
          in
            { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs argRes.expr, exprTypes: Array.snoc acc.exprTypes argRes.exprType, nextId: argRes.nextId }
      )
      { stmts: resFn.stmts, exprs: [], exprTypes: [], nextId: resFn.nextId }
      args
  in
    let
      len = Array.length args
      goFuncName = if len >= 2 && len <= 5 then "UncurriedApp" <> show len else "UncurriedApp"
    in
      case resFn.exprType of
        TypeFunc fArgs fRet | Array.length fArgs == len ->
          let
            callArgs = Array.mapWithIndex
              ( \i expected ->
                  let
                    arg = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs i)
                    actual = fromMaybe TypeValue (Array.index accArgs.exprTypes i)
                  in
                    coerceGoExpr codegenStateRef modNameStr arg actual expected
              )
              fArgs
          in
            { stmts: accArgs.stmts, expr: boxGoExpr codegenStateRef modNameStr (GoCall resFn.expr callArgs) fRet, exprType: TypeValue, nextId: accArgs.nextId }
        _ ->
          let
            boxedArgs = Array.zipWith (\arg actual -> boxGoExpr codegenStateRef modNameStr arg actual) accArgs.exprs accArgs.exprTypes
          in
            { stmts: accArgs.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") goFuncName) (Array.cons (boxGoExpr codegenStateRef modNameStr resFn.expr resFn.exprType) boxedArgs), exprType: TypeValue, nextId: accArgs.nextId }

-- Require the resolved intrinsic and concrete scalar callback. Outer Typed
-- annotations on the seed can describe the enclosing application, so only a
-- literal establishes its Int type here; opaque/dynamic seeds keep the copies.
isIntArrayFold :: TcoExpr -> Array TcoExpr -> Boolean
isIntArrayFold fn args = case unwrapTcoExpr fn, args of
  Var (Qualified (Just (ModuleName "Data.Foldable")) (Ident "foldlArray")), [ callback, seed, _ ] ->
    case extractFuncType callback, unwrapTcoExpr seed of
      Just { fArgs: [ Int, Int ], fRet: Int }, Lit (LitInt _) -> true
      _, _ -> false
  _, _ -> false

normalizeFreshIntArrayRoundtrip :: String -> GoExpr -> GoExpr
normalizeFreshIntArrayRoundtrip suffix expr = case expr of
  GoBoxIntArray (GoUnboxIntArray (GoCall (GoSelector (GoVar "gopurs_runtime") "Array") [ GoFreshFilterArray filtered ])) ->
    let
      sourceName = "source_int_array_" <> suffix
      itemsName = "items_int_array_" <> suffix
      indexName = "i_int_array_" <> suffix
      valueName = "v_int_array_" <> suffix
      source = GoCall (GoSelector (GoVar "gopurs_runtime") "Array") [ filtered ]
      -- The source runs once, before these IIFE-local names enter scope. The
      -- filter's make/append owns this buffer; do not search through variables,
      -- calls or storage for a marker. Normalize every Value before the fold
      -- to preserve IntVal/tag/pointer semantics of the two original copies.
      body = GoBlock
        [ GoAssign itemsName (GoCall (GoRaw "(*[]gopurs_runtime.Value)") [ GoSelector (GoVar sourceName) "UnsafePtr" ])
        , GoForRange (indexName <> ", " <> valueName <> " := range *" <> itemsName)
            [ GoMutate ("(*" <> itemsName <> ")[" <> indexName <> "]")
                (GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoSelector (GoVar valueName) "IntVal" ])
            ]
        , GoReturn (GoVar sourceName)
        ]
    in
      GoIIFE sourceName source body
  _ -> expr

applyBoxed :: GoExpr -> Array GoExpr -> GoExpr
applyBoxed fExpr argExprs =
  let
    len = Array.length argExprs
  in
    if len == 0 then fExpr
    else if len == 1 then GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ fExpr, fromMaybe (GoRaw "nil") (Array.index argExprs 0) ]
    else if len >= 2 && len <= 10 then
      GoCall (GoSelector (GoVar "gopurs_runtime") ("Apply" <> show len)) (Array.cons fExpr argExprs)
    else
      let
        chunk = Array.take 10 argExprs
        rest = Array.drop 10 argExprs
      in
        applyBoxed (applyBoxed fExpr chunk) rest
