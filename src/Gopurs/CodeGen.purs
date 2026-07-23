module Gopurs.CodeGen where

import Prelude
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(..), BackendAccessor(..), Pair(..), Level(..), BackendOperator(..), BackendOperator1(..), BackendOperator2(..), BackendOperatorOrd(..), BackendOperatorNum(..), BackendEffect(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Convert (BackendModule, BackendBindingGroup)
import Data.String as String
import Data.Array as Array
import Data.Maybe (Maybe(..), fromMaybe, isJust)
import Data.Newtype (unwrap)
import PureScript.Backend.Optimizer.CoreFn (Literal(..), Ident(..), Qualified(..), Prop(..), ModuleName(..))
import Data.Tuple (Tuple(..))
import Data.Array.NonEmpty as NonEmptyArray
import Data.Array.NonEmpty (NonEmptyArray, fromArray, toArray)
import Effect.Console as Console
import Effect.Unsafe (unsafePerformEffect)
import Effect.Ref (Ref)
import Effect.Ref as Ref
import Gopurs.FreeVars (freeVars)
import Data.Set as Set
import Data.String.CodeUnits as SCU
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Map (Map)
import Debug as Debug
import Data.Map as Map
import Data.Set (Set)
import Data.Set as Set
import Data.Foldable (foldl, foldMap)
import Data.List as List
import Data.Traversable (traverse)

import Gopurs.GoAst (GoFile, GoDecl, GoExpr(..))
import Gopurs.Printer (printGoFile, printGoExpr, printGoDeclVar)
import PureScript.Backend.Optimizer.Codegen.Tco as Tco
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..), tcoAnalysisOf)
import Gopurs.FreeVars (freeVars, localId)

capitalize :: String -> String
capitalize "" = ""
capitalize s =
  let
    firstChar = String.take 1 s
  in
    if firstChar >= "a" && firstChar <= "z" then String.toUpper firstChar <> String.drop 1 s
    else if firstChar == "_" then "_" <> capitalize (String.drop 1 s)
    else s

sanitizeName :: String -> String
sanitizeName name =
  let
    s1 = String.replaceAll (Pattern "'") (Replacement "_prime") (String.replaceAll (Pattern "$") (Replacement "_dollar") name)
  in
    if s1 == "break" || s1 == "default" || s1 == "func" || s1 == "interface" || s1 == "select" || s1 == "case" || s1 == "defer" || s1 == "go" || s1 == "map" || s1 == "struct" || s1 == "chan" || s1 == "else" || s1 == "goto" || s1 == "package" || s1 == "switch" || s1 == "const" || s1 == "fallthrough" || s1 == "if" || s1 == "range" || s1 == "type" || s1 == "continue" || s1 == "for" || s1 == "import" || s1 == "return" || s1 == "var" || s1 == "init" || s1 == "append" || s1 == "make" || s1 == "len" || s1 == "cap" || s1 == "new" || s1 == "close" || s1 == "delete" || s1 == "complex" || s1 == "real" || s1 == "imag" || s1 == "panic" || s1 == "recover" || s1 == "print" || s1 == "println" then s1 <> "_" else s1

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

wrapInStmts :: Array String -> StmtTree -> GoExpr -> GoExpr
wrapInStmts _ stmts expr =
  let
    stmtsArr = flattenStmts stmts
  in
    if Array.length stmtsArr == 0 then expr
    else GoRaw ("func() gopurs_runtime.Value {\n" <> printGoExpr (GoBlock (stmtsArr <> [ GoReturn expr ])) <> "\n}()")

extractUncurriedAbs :: TcoExpr -> Maybe { args :: Array String, body :: TcoExpr, fvs :: Set String }
extractUncurriedAbs tcoExpr@(TcoExpr _ syntax) = case syntax of
  UncurriedAbs args body ->
    Just { args: map (\(Tuple mbI lvl) -> localId mbI lvl) args, body, fvs: freeVars tcoExpr }
  Abs args body ->
    Just { args: map (\(Tuple mbI lvl) -> localId mbI lvl) (toArray args), body, fvs: freeVars tcoExpr }
  _ -> Nothing

unwrapExpr :: TcoExpr -> BackendSyntax TcoExpr
unwrapExpr (TcoExpr _ e) = e

flattenApp :: TcoExpr -> Tuple TcoExpr (Array TcoExpr)
flattenApp e@(TcoExpr _ (App f args)) =
  let
    Tuple f' args' = flattenApp f
  in
    Tuple f' (args' <> toArray args)
flattenApp e = Tuple e []

translate :: Array (Array String) -> BackendModule -> String
translate importsArray mod =
  let
    modNameStr = String.replaceAll (Pattern ".") (Replacement "_") (unwrap mod.name)
    _ = unsafePerformEffect (Console.log ("Translating module " <> modNameStr))
    helpersRef = unsafePerformEffect (Ref.new [])

    Tuple _ tcoBindings = foldl
      ( \(Tuple env acc) group ->
          let
            neBindings = fromArray group.bindings
            _ = unsafePerformEffect (Console.log ("Translating binding group"))
            env' = case neBindings of
              Just ne | group.recursive -> Tco.topLevelTcoEnvGroup mod.name ne <> env
              _ -> env
            tcoBinds = map (\(Tuple k v) -> Tuple k (Tco.analyze env' v)) group.bindings
          in
            Tuple env' (Array.snoc acc { recursive: group.recursive, bindings: tcoBinds })
      )
      (Tuple [] [])
      mod.bindings

    decls = Array.concatMap
      ( \group ->
          let
            recVars = if group.recursive then map (\(Tuple (Ident name) _) -> sanitizeName name) group.bindings else []
          in
            if group.recursive && Array.length group.bindings == 1 then
              let
                mutRecBinds = traverse (\(Tuple (Ident name) val) -> map (\abs -> { ident: sanitizeName name, args: abs.args, body: abs.body, fvs: abs.fvs }) (extractUncurriedAbs val)) group.bindings
              in
                case mutRecBinds of
                  Just fns ->
                    let
                      loopCtxs = map (\fn -> { ident: fn.ident, params: fn.args, loopParams: map (\p -> p <> "_loop") fn.args }) fns

                      fnWrapperStmts = map
                        ( \fn ->
                            let
                              currentLoopCtx = [ { ident: fn.ident, params: fn.args, loopParams: map (\p -> p <> "_loop") fn.args } ]
                              resBodyMut = translateExprImpl_ helpersRef 0 modNameStr recVars Map.empty Map.empty Nothing currentLoopCtx true false 0 fn.body
                              goName = fn.ident
                              loopParams = map (\p -> p <> "_loop") fn.args
                              initVars = Array.concatMap (\p -> [ GoRaw ("var " <> p <> " = " <> p <> "_loop"), GoRaw ("_ = " <> p) ]) fn.args
                              funcBody = GoFor goName (initVars <> flattenStmts resBodyMut.stmts <> [ GoReturn resBodyMut.expr ])
                              iife = GoRaw ("func() gopurs_runtime.Value {\n" <> printGoExpr funcBody <> "\n}()")
                              -- If fn is UncurriedAbs we might need different logic, but assuming curried for now:
                              funcExpr = Array.foldr (\p acc -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> p <> " gopurs_runtime.Value) gopurs_runtime.Value {\nreturn " <> printGoExpr acc <> "\n}") ]) iife loopParams
                            in
                              { identifier: goName, expression: funcExpr }
                        )
                        fns
                    in 
                      fnWrapperStmts
                  Nothing ->
                    Array.concatMap
                      ( \(Tuple (Ident name) expr) ->
                          let
                            res = translateExprImpl_ helpersRef 0 modNameStr recVars Map.empty Map.empty (Just (sanitizeName name)) [] false false 0 expr
                          in
                            [ { identifier: sanitizeName name, expression: wrapInStmts [] res.stmts res.expr } ]
                      )
                      group.bindings
            else
              Array.concatMap
                ( \(Tuple (Ident name) expr) ->
                    let
                      res = translateExprImpl_ helpersRef 0 modNameStr [] Map.empty Map.empty (Just (sanitizeName name)) [] false false 0 expr
                    in
                      [ { identifier: sanitizeName name, expression: wrapInStmts [] res.stmts res.expr } ]
                )
                group.bindings
      )
      tcoBindings

    allDeclsAst = decls <> unsafePerformEffect (Ref.read helpersRef)
    declsStr = String.joinWith "\\n" (map printGoDeclVar allDeclsAst)

    parts = String.split (Pattern "pkg_") declsStr
    usedPkgNames = Array.nub $ Array.mapMaybe
      ( \part ->
          let
            subParts = String.split (Pattern ".") part
          in
            Array.head subParts
      )
      (fromMaybe [] (Array.tail parts))

    goImports = Array.nub $
      (if Array.length allDeclsAst > 0 || Array.length (Array.fromFoldable mod.foreign) > 0 then [ "gopurs/output/gopurs_runtime" ] else [])
        <> (if Array.length allDeclsAst > 0 then [ "sync" ] else [])
        <> Array.mapMaybe
          ( \pkg ->
              if pkg /= modNameStr && pkg /= "Prim" && not (String.indexOf (Pattern "Prim_") pkg == Just 0) then Just ("gopurs/output/" <> String.replaceAll (Pattern "_") (Replacement ".") pkg)
              else Nothing
          )
          usedPkgNames

    goFile =
      { packageName: modNameStr
      , imports: goImports
      , decls: allDeclsAst
      , foreigns: map (\(Ident name) -> { pursName: sanitizeName name, goName: "_Gopurs_" <> capitalize (sanitizeName name) }) (Array.fromFoldable mod.foreign)
      }
  in
    printGoFile goFile


isEffectNode :: forall a. BackendSyntax a -> Boolean
isEffectNode = case _ of
  EffectBind _ _ _ _ -> true
  EffectPure _ -> true
  EffectDefer _ -> false
  PrimEffect _ -> true
  UncurriedEffectApp _ _ -> true
  _ -> false

unwrapTcoExpr :: TcoExpr -> BackendSyntax TcoExpr
unwrapTcoExpr (TcoExpr _ expr) = expr

executeIfOpaque :: forall a. BackendSyntax a -> GoExpr -> GoExpr

executeIfOpaque expr goExpr =
  if isEffectNode expr then goExpr
  else GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ goExpr, GoRaw "gopurs_runtime.Value{}" ]


translateExprImpl :: Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String } -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }
translateExprImpl helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail nextId tcoExpr =
  translateExprImpl_ helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail false nextId tcoExpr

translateExprImpl_ :: Ref (Array GoDecl) -> Int -> String -> Array String -> Map String String -> Map String String -> Maybe String -> Array { ident :: String, params :: Array String, loopParams :: Array String } -> Boolean -> Boolean -> Int -> TcoExpr -> { stmts :: StmtTree, expr :: GoExpr, nextId :: Int }
translateExprImpl_ helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx isTail inEffectBlock nextId tcoExpr@(TcoExpr tcoAnalysis expr) =
  let
    isEff = isEffectNode expr
  in
    if isEff && not inEffectBlock then
      let
        res = translateExprImpl_ helpersRef depth modNameStr recVars namedBound bound tcoIdent loopCtx false true nextId tcoExpr
        funcExpr = GoRaw ("gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {\n" <> printGoExpr (GoBlock (flattenStmts res.stmts <> [ GoReturn res.expr ])) <> "\n})")
      in
        { stmts: StmtEmpty, expr: funcExpr, nextId: res.nextId }
    else
  let
    liftIfNeeded mkNodeThunk =
      if depth > 10 then unsafePerformEffect do
        let fvsSet = freeVars tcoExpr
        let fvs = Array.fromFoldable fvsSet
        let helperName = "__helper_" <> show nextId
        let newNextId = nextId + 1
        let res = translateExprImpl_ helpersRef 0 modNameStr recVars namedBound bound Nothing [] false inEffectBlock newNextId tcoExpr

        let
          helperExpr =
            if Array.length fvs == 0 then GoFunc "_" (wrapInStmts [] res.stmts res.expr)
            else Array.foldr (\fv accFunc -> GoFunc fv accFunc) (wrapInStmts [] res.stmts res.expr) fvs

        Ref.modify_ (\arr -> Array.snoc arr { identifier: helperName, expression: helperExpr }) helpersRef

        let
          callExpr =
            if Array.length fvs == 0 then GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ GoVar helperName, GoRaw "gopurs_runtime.Int(0)" ]
            else Array.foldl (\accCall fv -> GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ accCall, GoVar fv ]) (GoVar helperName) fvs

        pure { stmts: StmtEmpty, expr: callExpr, nextId: res.nextId }
      else mkNodeThunk unit
  in
    case expr of
      Var (Qualified mbMn (Ident i)) ->
        let
          safeName = sanitizeName i
        in
          case mbMn of
            Just mn ->
              let
                modStr = unwrap mn
                modPkg = String.replaceAll (Pattern ".") (Replacement "_") modStr
              in
                if modPkg == modNameStr then { stmts: StmtEmpty, expr: GoCall (GoVar ("Get_" <> safeName)) [], nextId } else { stmts: StmtEmpty, expr: GoCall (GoSelector (GoVar ("pkg_" <> modPkg)) ("Get_" <> safeName)) [], nextId }
            Nothing -> { stmts: StmtEmpty, expr: GoCall (GoVar ("Get_" <> safeName)) [], nextId }

      Local mbIdent lvl ->
        let
          name = fromMaybe (localId mbIdent lvl) (Map.lookup (localId mbIdent lvl) bound)
        in
          { stmts: StmtEmpty, expr: GoVar name, nextId }

      Lit (LitString s) -> { stmts: StmtEmpty, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Str") [ GoString s ], nextId }
      Lit (LitInt i) -> { stmts: StmtEmpty, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoInt i ], nextId }
      Lit (LitNumber n) -> { stmts: StmtEmpty, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Float") [ GoRaw (show n) ], nextId }
      Lit (LitBoolean b) -> { stmts: StmtEmpty, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoRaw (if b then "true" else "false") ], nextId }
      Lit (LitChar c) -> { stmts: StmtEmpty, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Str") [ GoString (SCU.singleton c) ], nextId }

      Lit (LitArray xs) ->
        let
          accXs = foldl
            ( \acc val ->
                let
                  resVal = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false false acc.nextId val
                in
                  { stmts: acc.stmts <> resVal.stmts, exprs: Array.snoc acc.exprs resVal.expr, nextId: resVal.nextId }
            )
            { stmts: StmtEmpty, exprs: [], nextId }
            xs
        in
          { stmts: accXs.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Array") [ GoRaw ("[]gopurs_runtime.Value{" <> String.joinWith ", " (map printGoExpr accXs.exprs) <> "}") ], nextId: accXs.nextId }

      Lit (LitRecord props) ->
        let
          accProps = foldl
            ( \acc (Prop key val) ->
                let
                  resVal = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false false acc.nextId val
                in
                  { stmts: acc.stmts <> resVal.stmts, exprs: Array.snoc acc.exprs (Tuple key resVal.expr), nextId: resVal.nextId }
            )
            { stmts: StmtEmpty, exprs: [], nextId }
            props
        in
          { stmts: accProps.stmts, expr: GoRecordDict accProps.exprs, nextId: accProps.nextId }

      App fn args ->
        let
          argsArr = toArray args
          Tuple flatFn flatArgs = flattenApp tcoExpr

          isTailCallTo =
            if isTail then case unwrapExpr flatFn of
              Local mbIdent lvl ->
                let
                  v = fromMaybe (localId mbIdent lvl) (Map.lookup (localId mbIdent lvl) bound)
                in
                  Array.findIndex (\ctx -> ctx.ident == v) loopCtx
              Var (Qualified mbMod (Ident name)) ->
                let
                  fullName = fromMaybe "" (map (\(ModuleName m) -> String.joinWith "_" (String.split (Pattern ".") m) <> "_") mbMod) <> sanitizeName name
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
                        argRes = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false false acc.nextId arg
                      in
                        { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs argRes.expr, nextId: argRes.nextId }
                  )
                  { stmts: StmtEmpty, exprs: [], nextId }
                  flatArgs
                targetCtx = fromMaybe { ident: "", params: [], loopParams: [] } (Array.index loopCtx index)
                assigns = Array.mapWithIndex (\i paramName -> GoMutate paramName (fromMaybe (GoRaw "nil") (Array.index accFinal.exprs i))) targetCtx.loopParams
              in
                { stmts: accFinal.stmts <> foldMap StmtLeaf assigns <> StmtLeaf (GoContinue targetCtx.ident), expr: GoRaw "gopurs_runtime.Value{}", nextId: accFinal.nextId }

            Nothing ->
              let
                resFn = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false false nextId flatFn
                accArgs = foldl
                  ( \acc arg ->
                      let
                        argRes = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false false acc.nextId arg
                      in
                        { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs argRes.expr, nextId: argRes.nextId }
                  )
                  { stmts: resFn.stmts, exprs: [], nextId: resFn.nextId }
                  flatArgs

                buildApp :: GoExpr -> Array GoExpr -> GoExpr
                buildApp fExpr argExprs =
                  let len = Array.length argExprs
                  in
                    if len == 0 then fExpr
                    else if len == 1 then GoCall (GoSelector (GoVar "gopurs_runtime") "Apply") [ fExpr, fromMaybe (GoRaw "nil") (Array.index argExprs 0) ]
                    else if len >= 2 && len <= 5 then
                      GoCall (GoSelector (GoVar "gopurs_runtime") ("Apply" <> show len)) (Array.cons fExpr argExprs)
                    else
                      let chunk = Array.take 5 argExprs
                          rest = Array.drop 5 argExprs
                      in buildApp (buildApp fExpr chunk) rest

                finalExpr = buildApp resFn.expr accArgs.exprs
              in
                { stmts: accArgs.stmts, expr: finalExpr, nextId: accArgs.nextId }

      Abs args body ->
        let
          params = map (\(Tuple mbI lvl) -> localId mbI lvl) (toArray args)
          resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing loopCtx isTail false nextId body
          
          buildFunc :: Array String -> GoExpr -> GoExpr
          buildFunc ps innerExpr =
            let
              len = Array.length ps
              bodyStr = case innerExpr of
                GoBlock _ -> printGoExpr innerExpr
                _ -> "return " <> printGoExpr innerExpr
            in
              if len == 1 then
                GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> fromMaybe "" (Array.index ps 0) <> " gopurs_runtime.Value) gopurs_runtime.Value {\n" <> bodyStr <> "\n}") ]
              else if len >= 2 && len <= 5 then
                let goParams = String.joinWith ", " (map (\p -> p <> " gopurs_runtime.Value") ps)
                in GoCall (GoSelector (GoVar "gopurs_runtime") ("Func" <> show len)) [ GoRaw ("func(" <> goParams <> ") gopurs_runtime.Value {\n" <> bodyStr <> "\n}") ]
              else
                let chunk = Array.take 5 ps
                    rest = Array.drop 5 ps
                in buildFunc chunk (buildFunc rest innerExpr)
                
          funcExpr = buildFunc params (GoBlock (flattenStmts resBody.stmts <> [ GoReturn resBody.expr ]))
        in
          { stmts: StmtEmpty, expr: funcExpr, nextId: resBody.nextId }

      UncurriedApp fn args ->
        let
          resFn = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false false nextId fn
          accArgs = foldl
            ( \acc arg ->
                let
                  argRes = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false false acc.nextId arg
                in
                  { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs argRes.expr, nextId: argRes.nextId }
            )
            { stmts: resFn.stmts, exprs: [], nextId: resFn.nextId }
            args
        in
          let
            len = Array.length args
            goFuncName = if len >= 2 && len <= 5 then "UncurriedApp" <> show len else "UncurriedApp"
          in
            { stmts: accArgs.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") goFuncName) (Array.cons resFn.expr accArgs.exprs), nextId: accArgs.nextId }

      UncurriedAbs args body -> liftIfNeeded \_ ->
        let
          params = map (\(Tuple mbI lvl) -> localId mbI lvl) args
          goParams = String.joinWith ", " (map (\p -> p <> " gopurs_runtime.Value") params)
          resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing loopCtx isTail false nextId body
          arity = Array.length args
        in if arity >= 2 && arity <= 5 then
          let
            funcExpr = GoRaw ("gopurs_runtime.Func" <> show arity <> "(func(" <> goParams <> ") gopurs_runtime.Value {\n" <> printGoExpr (GoBlock (flattenStmts resBody.stmts <> [ GoReturn resBody.expr ])) <> "\n})")
          in { stmts: StmtEmpty, expr: funcExpr, nextId: resBody.nextId }
        else
          let
            makeCurried [] = resBody.expr
            makeCurried [p] = GoFunc p (GoBlock (flattenStmts resBody.stmts <> [ GoReturn resBody.expr ]))
            makeCurried ps = case Array.uncons ps of
              Just { head: p, tail: rest } -> GoFunc p (makeCurried rest)
              Nothing -> resBody.expr
          in { stmts: StmtEmpty, expr: makeCurried params, nextId: resBody.nextId }

      UncurriedEffectApp fn args ->
        let
          resFn = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false false nextId fn
          accArgs = foldl
            ( \acc arg ->
                let
                  argRes = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false false acc.nextId arg
                in
                  { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs argRes.expr, nextId: argRes.nextId }
            )
            { stmts: resFn.stmts, exprs: [], nextId: resFn.nextId }
            args
        in
          let
            len = Array.length args
            goFuncName = if len >= 2 && len <= 5 then "UncurriedApp" <> show len else "UncurriedApp"
          in
            { stmts: accArgs.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") goFuncName) (Array.cons resFn.expr accArgs.exprs), nextId: accArgs.nextId }

      UncurriedEffectAbs args body -> liftIfNeeded \_ ->
        let
          params = map (\(Tuple mbI lvl) -> localId mbI lvl) args
          goParams = String.joinWith ", " (map (\p -> p <> " gopurs_runtime.Value") params)
          resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing loopCtx isTail false nextId body
          arity = Array.length args
        in if arity >= 2 && arity <= 5 then
          let
            funcExpr = GoRaw ("gopurs_runtime.Func" <> show arity <> "(func(" <> goParams <> ") gopurs_runtime.Value {\n" <> printGoExpr (GoBlock (flattenStmts resBody.stmts <> [ GoReturn resBody.expr ])) <> "\n})")
          in { stmts: StmtEmpty, expr: funcExpr, nextId: resBody.nextId }
        else
          let
            makeCurried [] = resBody.expr
            makeCurried [p] = GoFunc p (GoBlock (flattenStmts resBody.stmts <> [ GoReturn resBody.expr ]))
            makeCurried ps = case Array.uncons ps of
              Just { head: p, tail: rest } -> GoFunc p (makeCurried rest)
              Nothing -> resBody.expr
          in { stmts: StmtEmpty, expr: makeCurried params, nextId: resBody.nextId }

      EffectBind mbIdent lvl binding body ->
        let
          originalName = localId mbIdent lvl
          name = originalName <> "_" <> show nextId
          newBound = Map.insert originalName name bound
          resBinding = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false true (nextId + 1) binding
          resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound newBound Nothing loopCtx isTail true resBinding.nextId body
          bindingExpr = executeIfOpaque (unwrapTcoExpr binding) resBinding.expr
          bodyExpr = executeIfOpaque (unwrapTcoExpr body) resBody.expr
        in
          { stmts: resBinding.stmts <> StmtLeaf (GoAssign name bindingExpr) <> resBody.stmts, expr: bodyExpr, nextId: resBody.nextId }

      EffectPure binding ->
        translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false false nextId binding

      EffectDefer binding ->
        let
          resBinding = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false true nextId binding
          funcExpr = GoRaw ("gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {\n" <> printGoExpr (GoBlock (flattenStmts resBinding.stmts <> [ GoReturn resBinding.expr ])) <> "\n})")
        in
          { stmts: StmtEmpty, expr: funcExpr, nextId: resBinding.nextId }

      Let mbIdent lvl binding body ->
        let
          originalName = localId mbIdent lvl
          name = originalName <> "_" <> show nextId
          newBound = Map.insert originalName name bound
          resBinding = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false false (nextId + 1) binding
          resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound newBound Nothing loopCtx isTail false resBinding.nextId body
        in
          { stmts: resBinding.stmts <> StmtLeaf (GoAssign name resBinding.expr) <> resBody.stmts, expr: resBody.expr, nextId: resBody.nextId }

      LetRec lvl bindings body ->
        let
          allocRes = foldl
            ( \acc (Tuple (Ident ident) _) ->
                let
                  oldName = localId (Just (Ident ident)) lvl
                  newName = oldName <> "_" <> show acc.nextId
                in
                  { newBound: Map.insert oldName newName acc.newBound, newNames: Array.snoc acc.newNames { oldName, newName }, nextId: acc.nextId + 1 }
            )
            { newBound: bound, newNames: [], nextId }
            (toArray bindings)

          combinedRecVars = recVars <> map (\(Tuple (Ident i) _) -> sanitizeName i) (toArray bindings)
          
          isLoop = (unwrap tcoAnalysis).role.isLoop
          mutRecBinds = if isLoop && Array.length (toArray bindings) == 1 then
              traverse (\(Tuple (Ident name) val) -> map (\abs -> { ident: sanitizeName name, args: abs.args, body: abs.body, fvs: abs.fvs }) (extractUncurriedAbs val)) (toArray bindings)
            else Nothing
        in
          case mutRecBinds of
            Just fns ->
              let
                loopCtxs = map (\fn ->
                    let
                      oldName = localId (Just (Ident fn.ident)) lvl
                      newName = fromMaybe oldName (Map.lookup oldName allocRes.newBound)
                    in
                      { ident: newName, params: fn.args, loopParams: map (\p -> p <> "_loop") fn.args }
                  ) fns
                
                combinedLoopCtx = loopCtxs <> loopCtx
                
                declStmts = map (\ctx -> GoRaw ("var " <> ctx.ident <> " gopurs_runtime.Value")) loopCtxs
                
                Tuple fnWrapperStmts nextId' = foldl
                  ( \(Tuple accStmts currNextId) fn ->
                      let
                        oldName = localId (Just (Ident fn.ident)) lvl
                        newName = fromMaybe oldName (Map.lookup oldName allocRes.newBound)
                        currentLoopCtx = [ { ident: newName, params: fn.args, loopParams: map (\p -> p <> "_loop") fn.args } ]
                        resBodyMut = translateExprImpl_ helpersRef (depth + 1) modNameStr combinedRecVars namedBound allocRes.newBound Nothing currentLoopCtx true false currNextId fn.body
                        
                        loopParams = map (\p -> p <> "_loop") fn.args
                        initVars = Array.concatMap (\p -> [ GoRaw ("var " <> p <> " = " <> p <> "_loop"), GoRaw ("_ = " <> p) ]) fn.args
                        
                        funcBody = GoFor newName (initVars <> flattenStmts resBodyMut.stmts <> [ GoReturn resBodyMut.expr ])
                        iife = GoRaw ("func() gopurs_runtime.Value {\n" <> printGoExpr funcBody <> "\n}()")
                        funcExpr = Array.foldr (\p acc -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> p <> " gopurs_runtime.Value) gopurs_runtime.Value {\nreturn " <> printGoExpr acc <> "\n}") ]) iife loopParams
                      in
                        Tuple (Array.snoc accStmts (GoMutate newName funcExpr)) resBodyMut.nextId
                  )
                  (Tuple [] allocRes.nextId)
                  fns
                
                resBodyOuter = translateExprImpl_ helpersRef (depth + 1) modNameStr combinedRecVars namedBound allocRes.newBound Nothing loopCtx isTail false nextId' body
              in
                { stmts: foldMap StmtLeaf declStmts <> foldMap StmtLeaf fnWrapperStmts <> resBodyOuter.stmts, expr: resBodyOuter.expr, nextId: resBodyOuter.nextId }
            
            Nothing ->
              let
                accBindings = foldl
                  ( \acc (Tuple (Tuple (Ident ident) val) alloc) ->
                      let
                        res = translateExprImpl_ helpersRef (depth + 1) modNameStr combinedRecVars namedBound allocRes.newBound (Just alloc.newName) [] false false acc.nextId val
                      in
                        { stmts: acc.stmts <> res.stmts, exprs: Array.snoc acc.exprs { key: alloc.newName, value: res.expr }, nextId: res.nextId }
                  )
                  { stmts: StmtEmpty, exprs: [], nextId: allocRes.nextId }
                  (Array.zip (toArray bindings) allocRes.newNames)

                declStmts = map (\b -> GoRaw ("var " <> b.key <> " gopurs_runtime.Value\n_ = " <> b.key)) accBindings.exprs
                assignStmts = map (\b -> GoMutate b.key b.value) accBindings.exprs

                resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr combinedRecVars namedBound allocRes.newBound Nothing loopCtx isTail false accBindings.nextId body
              in
                { stmts: foldMap StmtLeaf declStmts <> accBindings.stmts <> foldMap StmtLeaf assignStmts <> resBody.stmts, expr: resBody.expr, nextId: resBody.nextId }

      Accessor obj accessor ->
        let
          resObj = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false false nextId obj
        in
          case accessor of
            GetProp prop -> { stmts: resObj.stmts, expr: GoRecordAccess resObj.expr prop, nextId: resObj.nextId }
            GetIndex idx -> { stmts: resObj.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "ArrayAccess") [ resObj.expr, GoInt idx ], nextId: resObj.nextId }
            GetCtorField _ _ _ _ fieldName _ -> { stmts: resObj.stmts, expr: GoRecordAccess resObj.expr fieldName, nextId: resObj.nextId }

      Update obj props ->
        let
          resObj = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false false nextId obj
          accProps = foldl
            ( \acc (Prop key val) ->
                let
                  resVal = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false false acc.nextId val
                in
                  { stmts: acc.stmts <> resVal.stmts, exprs: Array.snoc acc.exprs (Tuple key resVal.expr), nextId: resVal.nextId }
            )
            { stmts: StmtEmpty, exprs: [], nextId: resObj.nextId }
            props
        in
          { stmts: resObj.stmts <> accProps.stmts, expr: GoRecordUpdateDict resObj.expr accProps.exprs, nextId: accProps.nextId }

      CtorDef _ _ (Ident name) fields ->
        let
          recordMap = GoRecordDict (Array.cons (Tuple "_tag" (GoCall (GoSelector (GoVar "gopurs_runtime") "Str") [ GoString name ])) (map (\f -> Tuple f (GoVar (sanitizeName f))) fields))
          funcExpr = Array.foldr (\f inner -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> sanitizeName f <> " gopurs_runtime.Value) gopurs_runtime.Value {\nreturn " <> printGoExpr inner <> "\n}") ]) recordMap fields
        in
          { stmts: StmtEmpty, expr: funcExpr, nextId }

      CtorSaturated _ _ _ (Ident name) props ->
        let
          accProps = foldl
            ( \acc (Tuple key val) ->
                let
                  resVal = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false false acc.nextId val
                in
                  { stmts: acc.stmts <> resVal.stmts, exprs: Array.snoc acc.exprs (Tuple key resVal.expr), nextId: resVal.nextId }
            )
            { stmts: StmtEmpty, exprs: [], nextId }
            props
        in
          { stmts: accProps.stmts, expr: GoRecordDict (Array.cons (Tuple "_tag" (GoCall (GoSelector (GoVar "gopurs_runtime") "Str") [ GoString name ])) accProps.exprs), nextId: accProps.nextId }

      Fail msg ->
        { stmts: StmtEmpty, expr: GoRaw ("func() gopurs_runtime.Value { panic(" <> printGoExpr (GoString msg) <> ") }()"), nextId }

      Branch branches def ->
        let
          resDef = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing loopCtx isTail false nextId def
          tmpVar = "__t" <> show resDef.nextId
          declTmp = StmtLeaf (GoRaw ("var " <> tmpVar <> " gopurs_runtime.Value"))
          labelName = "end_branch_" <> show resDef.nextId

          buildIfs = foldl
            ( \acc (Pair condExpr bodyExpr) ->
                let
                  resCond = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false false acc.nextId condExpr
                  resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing loopCtx isTail false resCond.nextId bodyExpr
                  goIf = GoIfElse resCond.expr (flattenStmts resBody.stmts <> [ GoMutate tmpVar resBody.expr, GoRaw ("goto " <> labelName) ]) []
                in
                  { stmts: acc.stmts <> StmtLeaf (GoRaw "{") <> resCond.stmts <> StmtLeaf goIf <> StmtLeaf (GoRaw "}"), nextId: resBody.nextId }
            )
            { stmts: StmtEmpty, nextId: resDef.nextId + 1 }
            (toArray branches)
        in
          { stmts: declTmp <> buildIfs.stmts <> StmtLeaf (GoRaw "{") <> resDef.stmts <> StmtLeaf (GoMutate tmpVar resDef.expr) <> StmtLeaf (GoRaw "}") <> StmtLeaf (GoRaw (labelName <> ":")), expr: GoVar tmpVar, nextId: buildIfs.nextId }

      PrimOp op -> case op of
        Op1 op1 e ->
          let
            resE = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false false nextId e
            goOp = case op1 of
              OpBooleanNot -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp "==" (GoSelector resE.expr "IntVal") (GoInt 0) ]
              OpIntNegate -> GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoBinOp "-" (GoInt 0) (GoSelector resE.expr "IntVal") ]
              OpIntBitNot -> GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoBinOp "^" (GoRaw "^0") (GoSelector resE.expr "IntVal") ]
              OpNumberNegate -> GoCall (GoSelector (GoVar "gopurs_runtime") "FloatNeg") [ resE.expr ]
              OpIsTag (Qualified _ (Ident tag)) -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp "==" (GoSelector (GoRecordAccess resE.expr "_tag") "StrVal") (GoString tag) ]
              OpArrayLength -> GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoCall (GoVar "int64") [ GoCall (GoVar "len") [ GoTypeAssertion (GoSelector resE.expr "PtrVal") "[]gopurs_runtime.Value" ] ] ]
          in
            { stmts: resE.stmts, expr: goOp, nextId: resE.nextId }
        Op2 op2 e1 e2 ->
          let
            res1 = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false false nextId e1
            res2 = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false false res1.nextId e2
            goOp = case op2 of
              OpArrayIndex -> GoCall (GoSelector (GoVar "gopurs_runtime") "ArrayAccess") [ res1.expr, GoCall (GoVar "int") [ GoSelector res2.expr "IntVal" ] ]
              OpIntNum OpAdd -> GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoBinOp "+" (GoSelector res1.expr "IntVal") (GoSelector res2.expr "IntVal") ]
              OpIntNum OpSubtract -> GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoBinOp "-" (GoSelector res1.expr "IntVal") (GoSelector res2.expr "IntVal") ]
              OpIntNum OpMultiply -> GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoBinOp "*" (GoSelector res1.expr "IntVal") (GoSelector res2.expr "IntVal") ]
              OpIntNum OpDivide -> GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoBinOp "/" (GoSelector res1.expr "IntVal") (GoSelector res2.expr "IntVal") ]
              OpIntBitAnd -> GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoBinOp "&" (GoSelector res1.expr "IntVal") (GoSelector res2.expr "IntVal") ]
              OpIntBitOr -> GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoBinOp "|" (GoSelector res1.expr "IntVal") (GoSelector res2.expr "IntVal") ]
              OpIntBitXor -> GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoBinOp "^" (GoSelector res1.expr "IntVal") (GoSelector res2.expr "IntVal") ]
              OpIntBitShiftLeft -> GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoBinOp "<<" (GoSelector res1.expr "IntVal") (GoSelector res2.expr "IntVal") ]
              OpIntBitShiftRight -> GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoBinOp ">>" (GoSelector res1.expr "IntVal") (GoSelector res2.expr "IntVal") ]
              OpIntBitZeroFillShiftRight -> GoCall (GoSelector (GoVar "gopurs_runtime") "Zshr") [ res1.expr, res2.expr ]
              OpIntOrd OpEq -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp "==" (GoSelector res1.expr "IntVal") (GoSelector res2.expr "IntVal") ]
              OpIntOrd OpNotEq -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp "!=" (GoSelector res1.expr "IntVal") (GoSelector res2.expr "IntVal") ]
              OpIntOrd OpLt -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp "<" (GoSelector res1.expr "IntVal") (GoSelector res2.expr "IntVal") ]
              OpIntOrd OpLte -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp "<=" (GoSelector res1.expr "IntVal") (GoSelector res2.expr "IntVal") ]
              OpIntOrd OpGt -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp ">" (GoSelector res1.expr "IntVal") (GoSelector res2.expr "IntVal") ]
              OpIntOrd OpGte -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp ">=" (GoSelector res1.expr "IntVal") (GoSelector res2.expr "IntVal") ]
              OpNumberNum OpAdd -> GoCall (GoSelector (GoVar "gopurs_runtime") "FloatAdd") [ res1.expr, res2.expr ]
              OpNumberNum OpSubtract -> GoCall (GoSelector (GoVar "gopurs_runtime") "FloatSub") [ res1.expr, res2.expr ]
              OpNumberNum OpMultiply -> GoCall (GoSelector (GoVar "gopurs_runtime") "FloatMul") [ res1.expr, res2.expr ]
              OpNumberNum OpDivide -> GoCall (GoSelector (GoVar "gopurs_runtime") "FloatDiv") [ res1.expr, res2.expr ]
              OpNumberOrd OpEq -> GoCall (GoSelector (GoVar "gopurs_runtime") "FloatEq") [ res1.expr, res2.expr ]
              OpNumberOrd OpNotEq -> GoCall (GoSelector (GoVar "gopurs_runtime") "FloatNeq") [ res1.expr, res2.expr ]
              OpNumberOrd OpLt -> GoCall (GoSelector (GoVar "gopurs_runtime") "FloatLt") [ res1.expr, res2.expr ]
              OpNumberOrd OpLte -> GoCall (GoSelector (GoVar "gopurs_runtime") "FloatLte") [ res1.expr, res2.expr ]
              OpNumberOrd OpGt -> GoCall (GoSelector (GoVar "gopurs_runtime") "FloatGt") [ res1.expr, res2.expr ]
              OpNumberOrd OpGte -> GoCall (GoSelector (GoVar "gopurs_runtime") "FloatGte") [ res1.expr, res2.expr ]
              OpStringAppend -> GoCall (GoSelector (GoVar "gopurs_runtime") "Str") [ GoBinOp "+" (GoSelector res1.expr "StrVal") (GoSelector res2.expr "StrVal") ]
              OpStringOrd OpEq -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp "==" (GoSelector res1.expr "StrVal") (GoSelector res2.expr "StrVal") ]
              OpStringOrd OpNotEq -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp "!=" (GoSelector res1.expr "StrVal") (GoSelector res2.expr "StrVal") ]
              OpStringOrd OpLt -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp "<" (GoSelector res1.expr "StrVal") (GoSelector res2.expr "StrVal") ]
              OpStringOrd OpLte -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp "<=" (GoSelector res1.expr "StrVal") (GoSelector res2.expr "StrVal") ]
              OpStringOrd OpGt -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp ">" (GoSelector res1.expr "StrVal") (GoSelector res2.expr "StrVal") ]
              OpStringOrd OpGte -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp ">=" (GoSelector res1.expr "StrVal") (GoSelector res2.expr "StrVal") ]
              OpCharOrd OpEq -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp "==" (GoSelector res1.expr "StrVal") (GoSelector res2.expr "StrVal") ]
              OpCharOrd OpNotEq -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp "!=" (GoSelector res1.expr "StrVal") (GoSelector res2.expr "StrVal") ]
              OpCharOrd OpLt -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp "<" (GoSelector res1.expr "StrVal") (GoSelector res2.expr "StrVal") ]
              OpCharOrd OpLte -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp "<=" (GoSelector res1.expr "StrVal") (GoSelector res2.expr "StrVal") ]
              OpCharOrd OpGt -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp ">" (GoSelector res1.expr "StrVal") (GoSelector res2.expr "StrVal") ]
              OpCharOrd OpGte -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp ">=" (GoSelector res1.expr "StrVal") (GoSelector res2.expr "StrVal") ]
              OpBooleanOrd OpEq -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp "==" (GoSelector res1.expr "IntVal") (GoSelector res2.expr "IntVal") ]
              OpBooleanOrd OpNotEq -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp "!=" (GoSelector res1.expr "IntVal") (GoSelector res2.expr "IntVal") ]
              OpBooleanOrd OpLt -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp "<" (GoSelector res1.expr "IntVal") (GoSelector res2.expr "IntVal") ]
              OpBooleanOrd OpLte -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp "<=" (GoSelector res1.expr "IntVal") (GoSelector res2.expr "IntVal") ]
              OpBooleanOrd OpGt -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp ">" (GoSelector res1.expr "IntVal") (GoSelector res2.expr "IntVal") ]
              OpBooleanOrd OpGte -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp ">=" (GoSelector res1.expr "IntVal") (GoSelector res2.expr "IntVal") ]
              OpBooleanAnd -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp "&&" (GoBinOp "!=" (GoSelector res1.expr "IntVal") (GoInt 0)) (GoBinOp "!=" (GoSelector res2.expr "IntVal") (GoInt 0)) ]
              OpBooleanOr -> GoCall (GoSelector (GoVar "gopurs_runtime") "Bool") [ GoBinOp "||" (GoBinOp "!=" (GoSelector res1.expr "IntVal") (GoInt 0)) (GoBinOp "!=" (GoSelector res2.expr "IntVal") (GoInt 0)) ]
          in
            { stmts: res1.stmts <> res2.stmts, expr: goOp, nextId: res2.nextId }

      PrimEffect eff -> case eff of
        EffectRefNew a ->
          let
            resA = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false false nextId a
            refIdent = "__local_ref_" <> show resA.nextId
            declStmt = GoAssign refIdent resA.expr
          in
            { stmts: resA.stmts <> StmtLeaf declStmt
            , expr: GoRaw ("gopurs_runtime.Value{PtrVal: &" <> refIdent <> "}")
            , nextId: resA.nextId + 1
            }
        EffectRefRead a ->
          let
            resA = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false false nextId a
          in
            { stmts: resA.stmts
            , expr: GoRaw ("*(" <> printGoExpr resA.expr <> ".PtrVal.(*gopurs_runtime.Value))")
            , nextId: resA.nextId
            }
        EffectRefWrite ref val ->
          let
            resRef = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false false nextId ref
            resVal = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false false resRef.nextId val
            writeStmt = GoRaw ("*(" <> printGoExpr resRef.expr <> ".PtrVal.(*gopurs_runtime.Value)) = " <> printGoExpr resVal.expr)
          in
            { stmts: resRef.stmts <> resVal.stmts <> StmtLeaf writeStmt
            , expr: resVal.expr
            , nextId: resVal.nextId
            }

      _ -> { stmts: StmtEmpty, expr: GoVar "gopurs_runtime.Value{}", nextId }

