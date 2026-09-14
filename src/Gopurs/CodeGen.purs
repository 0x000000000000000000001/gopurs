module Gopurs.CodeGen
  ( translate
  , translateWithFunctions
  ) where

import Prelude
import Control.Alternative (guard)
import PureScript.Backend.Optimizer.Syntax (BackendAccessor(..), BackendOperator(..), BackendOperator1(..), BackendOperator2(..), BackendSyntax(PrimEffect, PrimOp, Branch, Fail, CtorDef, CtorSaturated, Lit, UncurriedAbs, Abs, Update, Accessor, LetRec, Let, EffectDefer, EffectPure, EffectBind, UncurriedEffectAbs, UncurriedEffectApp, UncurriedApp, App, Local, Var, Typed))
import PureScript.Backend.Optimizer.Syntax as Syn
import PureScript.Backend.Optimizer.Convert (BackendModule)
import Data.String as String
import Data.Array as Array
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Newtype (unwrap)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Ident(..), Literal(..), Prop(..), Qualified(..))
import Data.Tuple (Tuple(..))
import Data.Array.NonEmpty (toArray)
import Effect.Unsafe (unsafePerformEffect)
import Effect.Ref (Ref)
import Effect.Ref as Ref
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Map as Map
import Data.Set as Set
import Data.Foldable (foldl)
import Gopurs.GoAst (rawGo, GoExpr(..), GoDecl(..), GoType(..), capitalize, goTypeToStr, sanitizeName)
import Gopurs.Printer (printGoFile, printGoExpr)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.FfiSupport (hashString)
import Gopurs.ThunkFusion (optimizeThunkProducers)
import Gopurs.GoTypes (exprTypeToGenericGoType, exprTypeToGoType, instantiateGenericGoType)
import Gopurs.CodegenState (CodegenMetadata, CodegenState)
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr, generateReboxFunctions, unboxGoExpr)
import Gopurs.PrimitiveExprs as PrimitiveExprs
import Gopurs.RecordExprs as RecordExprs
import Gopurs.AdtExprs as AdtExprs
import Gopurs.CallAnalysis (isClosureNode)
import Gopurs.CallExprs as CallExprs
import Gopurs.FunctionExprs as FunctionExprs
import Gopurs.BindingExprs as BindingExprs
import Gopurs.ControlExprs as ControlExprs
import Gopurs.EffectExprs as EffectExprs
import Gopurs.ModuleBindings as ModuleBindings
import Gopurs.ModuleDeclarations as ModuleDeclarations
import Gopurs.GoImports (collectImports)
import Gopurs.ExprAnalysis (bindFieldFunctionParameters, extractExprFuncType, getExprType, hasTypeVars, isEffectNode, unwrapTcoExpr)
import Gopurs.ExprContext (ExprOptions, ExprResult, LocalEnv, LoopContext, ModuleFunctions, StmtTree(..), TranslateExpr)

translate :: CodegenMetadata -> BackendModule -> String
translate metadata inputMod = (translateWithFunctions metadata inputMod).code

translateWithFunctions :: CodegenMetadata -> BackendModule -> { code :: String, functions :: ModuleFunctions }
translateWithFunctions metadata inputMod =

  let
    mod = optimizeThunkProducers inputMod
    modNameStrOrig = unwrap mod.name
    modNameStr = String.replaceAll (Pattern ".") (Replacement "_") modNameStrOrig

    codegenStateRef :: Ref CodegenState
    codegenStateRef = unsafePerformEffect $
      Ref.new { declarations: ModuleDeclarations.constructors metadata modNameStr mod, globalId: 0, reboxPairs: Set.empty }

    preparedBindings = ModuleBindings.prepare metadata modNameStr mod
    moduleFunctions = preparedBindings.functions

    Tuple allDeclsAst helpers = unsafePerformEffect do
      let
        d = ModuleBindings.declarations translateInContext metadata codegenStateRef modNameStr moduleFunctions preparedBindings.bindings
      h <- Ref.read codegenStateRef
      pure (Tuple d h)

    foreignGetters = map
      (\(Tuple (Ident name) _) -> GoForeignGetter
        { name: "Get_" <> modNameStr <> "_" <> sanitizeName name
        , value: "_Gopurs_" <> modNameStr <> "_" <> capitalize (sanitizeName name)
        })
      (Map.toUnfoldable mod.foreign)
    declarationGroups =
      [ allDeclsAst
      , helpers.declarations <> unsafePerformEffect (generateReboxFunctions metadata codegenStateRef modNameStr)
      , foreignGetters
      ]
    goFile = { packageName: "purescript", imports: collectImports declarationGroups, declarationGroups }

  in
    { code: printGoFile goFile
    , functions: Map.fromFoldable $ Array.concatMap
        (\group -> Array.mapMaybe
          (\(Tuple (Ident name) _) -> do
            info <- Map.lookup (sanitizeName name) moduleFunctions
            guard (info.arity >= 1 && info.arity <= 10)
            pure (Tuple (unwrap mod.name <> "." <> name) info))
          group.bindings)
        preparedBindings.bindings
    }

translateInContext :: TranslateExpr
translateInContext { metadata, codegenStateRef, depth, modNameStr, recVars, moduleFunctions, bound, tcoIdent, loopCtx, options, mbExpectedExprType } nextId expression =
  translateExprWithExpectedType metadata codegenStateRef depth modNameStr recVars moduleFunctions bound tcoIdent loopCtx options mbExpectedExprType nextId expression

translateExpr :: CodegenMetadata -> Ref CodegenState -> Int -> String -> Array String -> ModuleFunctions -> LocalEnv -> Maybe String -> LoopContext -> ExprOptions -> Int -> TcoExpr -> ExprResult
translateExpr metadata codegenStateRef depth modNameStr recVars moduleFunctions bound tcoIdent loopCtx options nextId tcoExpr =
  translateExprWithExpectedType metadata codegenStateRef depth modNameStr recVars moduleFunctions bound tcoIdent loopCtx options Nothing nextId tcoExpr

translateExprWithExpectedType :: CodegenMetadata -> Ref CodegenState -> Int -> String -> Array String -> ModuleFunctions -> LocalEnv -> Maybe String -> LoopContext -> ExprOptions -> Maybe ExprType -> Int -> TcoExpr -> ExprResult
translateExprWithExpectedType metadata codegenStateRef depth modNameStr recVars moduleFunctions bound tcoIdent loopCtx options@{ inEffectBlock } mbExpectedExprType nextId tcoExpr@(TcoExpr _ expr) =
  let
    context = { metadata, codegenStateRef, depth, modNameStr, recVars, moduleFunctions, bound, tcoIdent, loopCtx, options, mbExpectedExprType }
    isEff = isEffectNode tcoExpr
  in
    if isEff && not inEffectBlock then
      EffectExprs.wrap translateInContext context nextId tcoExpr
    else
      case expr of
        Typed type_ a ->
          let
            expectedGoType = exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr type_

          in
            case unwrapTcoExpr a, expectedGoType of
              Lit (LitRecord props), TypeStructPointer { baseStructName, fullName, structName: monoStructName, typeArgs: typeArgsForDict } ->
                case Map.lookup fullName metadata.classDeclsFields of
                  Just classInfo ->
                    let
                      classFields = classInfo.fields
                      typeArgs = case type_ of
                        ADT className _ tArgs ->
                          let
                            mapped = map (exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr) tArgs
                            arity = case Map.lookup className metadata.pointerAdtPaths of
                              Just info -> info.arity
                              Nothing -> Array.length classInfo.vars
                          in
                            Array.take arity mapped
                        _ -> map (const TypeValue) classInfo.vars
                      instMap = Map.fromFoldable (Array.zip classInfo.vars typeArgs)
                      propMap = Map.fromFoldable (map (\(Prop k v) -> Tuple k v) props)
                      sortedVals = Array.mapMaybe (\f -> map (\v -> { field: f, val: v }) (Map.lookup f.name propMap)) classFields
                      accProps = foldl
                        ( \acc item ->
                            let
                              genericGoType = exprTypeToGenericGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors classInfo.vars modNameStr item.field."type"
                              expectedType = instantiateGenericGoType instMap genericGoType
                              expectedExprType = item.field."type"

                              newBound = bindFieldFunctionParameters
                                (exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr)
                                bound
                                expectedExprType
                                item.val

                              resVal = translateExpr metadata codegenStateRef (depth + 1) modNameStr recVars moduleFunctions newBound Nothing [] { isTail: false, inEffectBlock: false } acc.nextId item.val
                              coercedExpr = coerceGoExpr codegenStateRef modNameStr resVal.expr resVal.exprType expectedType
                            in
                              { stmts: acc.stmts <> resVal.stmts, exprs: Array.snoc acc.exprs coercedExpr, exprType: TypeValue, nextId: resVal.nextId }
                        )
                        { stmts: StmtEmpty, exprs: [], exprType: TypeValue, nextId }
                        sortedVals
                    in
                      { stmts: accProps.stmts, expr: GoConstructor (hashString baseStructName) monoStructName typeArgsForDict accProps.exprs, exprType: expectedGoType, nextId: accProps.nextId }
                  Nothing ->
                    let
                      res = translateExprWithExpectedType metadata codegenStateRef depth modNameStr recVars moduleFunctions bound tcoIdent loopCtx options (Just type_) nextId a
                    in
                      case res.exprType of
                        TypeStructPointer _ -> res
                        _ ->
                          if expectedGoType == res.exprType then res
                          else if isClosureNode metadata a then res
                          else
                            { stmts: res.stmts, expr: coerceGoExpr codegenStateRef modNameStr res.expr res.exprType expectedGoType, exprType: expectedGoType, nextId: res.nextId }
              Let ident lvl val body, _ ->
                let
                  newBody = case body of
                    TcoExpr bAnn _ -> TcoExpr bAnn (Typed type_ body)
                  newLetShape = Let ident lvl val newBody
                  newA = case a of
                    TcoExpr ann _ -> TcoExpr ann newLetShape
                in
                  translateExpr metadata codegenStateRef depth modNameStr recVars moduleFunctions bound tcoIdent loopCtx options nextId newA
              LetRec lvl bindings body, _ ->
                let
                  newBody = case body of
                    TcoExpr bAnn _ -> TcoExpr bAnn (Typed type_ body)
                  newLetShape = LetRec lvl bindings newBody
                  newA = case a of
                    TcoExpr ann _ -> TcoExpr ann newLetShape
                in
                  translateExpr metadata codegenStateRef depth modNameStr recVars moduleFunctions bound tcoIdent loopCtx options nextId newA
              _, _ ->
                let
                  res = translateExprWithExpectedType metadata codegenStateRef depth modNameStr recVars moduleFunctions bound tcoIdent loopCtx options (Just type_) nextId a
                  preserveBoxedRecord = case expectedGoType, mbExpectedExprType of
                    TypeRecord _, Just Any -> res.exprType == TypeValue
                    _, _ -> false
                in
                  case res.exprType of
                    TypeStructPointer _ -> res
                    _ ->
                      if expectedGoType == res.exprType || preserveBoxedRecord then res
                      else if isClosureNode metadata a then res
                      else
                        { stmts: res.stmts, expr: coerceGoExpr codegenStateRef modNameStr res.expr res.exprType expectedGoType, exprType: expectedGoType, nextId: res.nextId }
        Var (Qualified mbMn (Ident i)) ->
          let
            safeName = sanitizeName i
            vType = case mbMn of
              Just mn ->
                let
                  modStr = unwrap mn
                in
                  case Map.lookup (modStr <> "." <> i) metadata.globalTypes of
                    Just ty -> exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr ty
                    Nothing -> TypeValue
              Nothing -> TypeValue
          in
            case mbMn of
              Just mn ->
                let
                  modStr = unwrap mn
                  modPkg = String.replaceAll (Pattern ".") (Replacement "_") modStr
                  rawCall = GoCall (GoVar ("Get_" <> modPkg <> "_" <> safeName)) []
                in
                  { stmts: StmtEmpty, expr: coerceGoExpr codegenStateRef modNameStr rawCall TypeValue vType, exprType: vType, nextId }
              Nothing ->
                let
                  rawCall = GoCall (GoVar ("Get_" <> modNameStr <> "_" <> safeName)) []
                in
                  { stmts: StmtEmpty, expr: coerceGoExpr codegenStateRef modNameStr rawCall TypeValue vType, exprType: vType, nextId }

        Local mbIdent lvl ->
          let
            v = fromMaybe { name: localId mbIdent lvl, goType: TypeValue } (Map.lookup (localId mbIdent lvl) bound)
          in
            { stmts: StmtEmpty, expr: GoVar v.name, exprType: v.goType, nextId }

        Lit value
          | Just result <- PrimitiveExprs.literal value ->
              { stmts: StmtEmpty, expr: result.expr, exprType: result.exprType, nextId }

        Lit (LitArray xs) ->
          let
            accXs = foldl
              ( \acc val ->
                  let
                    resVal = translateExpr metadata codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } acc.nextId val
                  in
                    { stmts: acc.stmts <> resVal.stmts, exprs: Array.snoc acc.exprs resVal.expr, exprTypes: Array.snoc acc.exprTypes resVal.exprType, nextId: resVal.nextId }
              )
              { stmts: StmtEmpty, exprs: [], exprTypes: [], nextId }
              xs

            mbElemType = Array.head accXs.exprTypes
            isAllSame = Array.all (\t -> Just t == mbElemType) accXs.exprTypes

            expectedElemType = case mbExpectedExprType of
              Just exTy ->
                case exprTypeToGenericGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors [] modNameStr exTy of
                  TypeNativeArray et -> Just et
                  _ -> Nothing
              Nothing -> Nothing

            finalElemType = case expectedElemType of
              Just t -> Just t
              Nothing -> if isAllSame then mbElemType else Nothing
          in
            case finalElemType of
              Just elemType | (isAllSame || Array.length xs == 0) && elemType /= TypeValue ->
                let
                  goTypeArr = TypeNativeArray elemType
                in
                  { stmts: accXs.stmts, expr: rawGo (goTypeToStr goTypeArr <> "{" <> String.joinWith ", " (map printGoExpr accXs.exprs) <> "}"), exprType: goTypeArr, nextId: accXs.nextId }
              _ ->
                let
                  boxedExprs = Array.zipWith (\itemExpr ty -> boxGoExpr codegenStateRef modNameStr itemExpr ty) accXs.exprs accXs.exprTypes
                in
                  { stmts: accXs.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Array") [ rawGo ("[]gopurs_runtime.Value{" <> String.joinWith ", " (map printGoExpr boxedExprs) <> "}") ], exprType: TypeValue, nextId: accXs.nextId }

        Lit (LitRecord props) ->
          let
            sortedProps = Array.sortBy (comparing \(Prop k _) -> k) props
            recordInfo = RecordExprs.prepareLiteral metadata modNameStr (getExprType tcoExpr) mbExpectedExprType

            accProps = foldl
              ( \acc (Prop key val) ->
                  let
                    expectedExprType = fromMaybe Any (Map.lookup key recordInfo.fields)
                    newBound = bindFieldFunctionParameters
                      (\fArgTy -> exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr fArgTy)
                      bound
                      expectedExprType
                      val

                    resVal = translateExpr metadata codegenStateRef (depth + 1) modNameStr recVars moduleFunctions newBound Nothing [] { isTail: false, inEffectBlock: false } acc.nextId val
                    field = RecordExprs.coerceLiteralField codegenStateRef modNameStr key recordInfo.recordType { expr: resVal.expr, exprType: resVal.exprType }
                  in
                    { stmts: acc.stmts <> resVal.stmts, exprs: Array.snoc acc.exprs field, exprType: TypeValue, nextId: resVal.nextId }
              )
              { stmts: StmtEmpty, exprs: [], exprType: TypeValue, nextId }
              sortedProps
            result = RecordExprs.literal recordInfo.recordType accProps.exprs
          in
            { stmts: accProps.stmts, expr: result.expr, exprType: result.exprType, nextId: accProps.nextId }

        expr_
          | ( case expr_ of
                App _ _ -> true
                Syn.TypeApp _ _ -> true
                _ -> false
            ) ->
              CallExprs.application translateInContext context nextId tcoExpr

        Abs args body ->
          FunctionExprs.abstraction translateInContext context nextId tcoExpr args body

        UncurriedApp fn args ->
          CallExprs.uncurriedApplication translateInContext context nextId tcoExpr fn args

        UncurriedAbs args body ->
          FunctionExprs.uncurriedAbstraction translateInContext context nextId tcoExpr args body

        UncurriedEffectApp fn args ->
          CallExprs.effectApplication translateInContext context nextId fn args

        UncurriedEffectAbs args body ->
          FunctionExprs.effectAbstraction translateInContext context nextId tcoExpr args body

        EffectBind mbIdent lvl binding body ->
          EffectExprs.bindEffect translateInContext context nextId mbIdent lvl binding body

        EffectPure binding ->
          EffectExprs.pureValue translateInContext context nextId binding

        EffectDefer binding ->
          EffectExprs.defer translateInContext context nextId binding

        Let mbIdent lvl binding body ->
          BindingExprs.nonRecursive translateInContext context nextId mbIdent lvl binding body

        LetRec lvl bindings body ->
          BindingExprs.recursive translateInContext context nextId tcoExpr lvl bindings body

        Accessor obj accessor ->
          let
            resObj = translateExpr metadata codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId obj
          in
            case accessor of
              GetProp prop ->
                let
                  result = RecordExprs.getProp metadata codegenStateRef modNameStr prop { expr: resObj.expr, exprType: resObj.exprType }
                in
                  { stmts: resObj.stmts, expr: result.expr, exprType: result.exprType, nextId: resObj.nextId }
              GetIndex idx -> { stmts: resObj.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "ArrayAccess") [ (boxGoExpr codegenStateRef modNameStr resObj.expr resObj.exprType), GoInt idx ], exprType: TypeValue, nextId: resObj.nextId }
              GetCtorField (Qualified mbMod _) _ _ (Ident ctorName) _ idx ->
                let
                  result = AdtExprs.getField metadata codegenStateRef modNameStr
                    { moduleName: mbMod, ctorName, index: idx }
                    { expr: resObj.expr, exprType: resObj.exprType, sourceType: getExprType obj }
                in
                  { stmts: resObj.stmts, expr: result.expr, exprType: result.exprType, nextId: resObj.nextId }

        Update obj props ->
          let
            resObj = translateExpr metadata codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId obj
            accProps = foldl
              ( \acc (Prop key val) ->
                  let
                    resVal = translateExpr metadata codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } acc.nextId val
                  in
                    { stmts: acc.stmts <> resVal.stmts, exprs: Array.snoc acc.exprs { key, expr: resVal.expr, goType: resVal.exprType }, exprType: TypeValue, nextId: resVal.nextId }
              )
              { stmts: StmtEmpty, exprs: [], exprType: TypeValue, nextId: resObj.nextId }
              props
            resultType = map (exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr) mbExpectedExprType
            result = RecordExprs.update codegenStateRef modNameStr resultType { expr: resObj.expr, exprType: resObj.exprType } accProps.exprs
          in
            { stmts: resObj.stmts <> accProps.stmts, expr: result.expr, exprType: result.exprType, nextId: accProps.nextId }

        CtorDef _ _ (Ident name) fields ->
          let
            annotatedType = fromMaybe (getExprType tcoExpr) mbExpectedExprType
            ctorType = case extractExprFuncType annotatedType of
              Just { fRet } -> fRet
              Nothing -> annotatedType
            result = AdtExprs.definition metadata codegenStateRef modNameStr name fields ctorType
          in
            { stmts: StmtEmpty, expr: result.expr, exprType: result.exprType, nextId }

        CtorSaturated (Qualified mbMod _) _ _ (Ident name) props ->
          let
            ctorType = case getExprType tcoExpr of
              Any -> fromMaybe Any mbExpectedExprType
              ty ->
                if hasTypeVars ty then
                  case mbExpectedExprType of
                    Just expectedTy | not (hasTypeVars expectedTy) -> expectedTy
                    _ -> ty
                else ty

            prepared = AdtExprs.prepareSaturated metadata modNameStr mbMod name ctorType

            accProps = foldl
              ( \acc (Tuple _ val) ->
                  let
                    { exprType: expectedExprType, goType: expectedType } = AdtExprs.saturatedFieldType prepared acc.fieldIdx

                    newBound = case unwrapTcoExpr val, extractExprFuncType expectedExprType of
                      Abs args _, Just { fArgs } ->
                        let
                          paramsWithTypes = Array.zipWith (\(Tuple mbI lvl) fArgTy -> Tuple (localId mbI lvl) (exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr fArgTy)) (toArray args) (fArgs <> Array.replicate (Array.length (toArray args) - Array.length fArgs) Any)
                        in
                          foldl (\b (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } b) bound paramsWithTypes
                      UncurriedAbs args _, Just { fArgs } ->
                        let
                          paramsWithTypes = Array.zipWith (\(Tuple mbI lvl) fArgTy -> Tuple (localId mbI lvl) (exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr fArgTy)) args (fArgs <> Array.replicate (Array.length args - Array.length fArgs) Any)
                        in
                          foldl (\b (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } b) bound paramsWithTypes
                      _, _ -> bound

                    resVal = translateExpr metadata codegenStateRef (depth + 1) modNameStr recVars moduleFunctions newBound Nothing [] { isTail: false, inEffectBlock: false } acc.nextId val
                    coercedExpr = coerceGoExpr codegenStateRef modNameStr resVal.expr resVal.exprType expectedType
                    isConstant = expectedType == resVal.exprType && case expectedType, unwrapTcoExpr val of
                      TypeInt64, Lit (LitInt _) -> true
                      TypeBool, Lit (LitBoolean _) -> true
                      TypeUint32, CtorSaturated _ _ _ _ ctorFields -> Array.null ctorFields
                      TypeUint32, CtorDef _ _ _ ctorFields -> Array.null ctorFields
                      _, _ -> false
                  in
                    { stmts: acc.stmts <> resVal.stmts, exprs: Array.snoc acc.exprs coercedExpr, exprTypes: Array.snoc acc.exprTypes expectedType, exprType: TypeValue, nextId: resVal.nextId, fieldIdx: acc.fieldIdx + 1, constants: Array.snoc acc.constants isConstant }
              )
              { stmts: StmtEmpty, exprs: [], exprTypes: [], exprType: TypeValue, nextId, fieldIdx: 0, constants: [] }
              props

            res = AdtExprs.saturated codegenStateRef prepared { exprs: accProps.exprs, exprTypes: accProps.exprTypes }
            reuse = case accProps.stmts of
              StmtEmpty -> AdtExprs.constructorReuse bound res.exprType accProps.constants res.expr
              _ -> Nothing
          in
            case reuse of
              Just { source, condition } ->
                let
                  resultName = "__reuse_" <> show accProps.nextId
                  declare = rawGo ("var " <> resultName <> " " <> goTypeToStr res.exprType)
                  choose = GoIfElse condition [ GoMutate resultName source ] [ GoMutate resultName res.expr ]
                in
                  { stmts: StmtLeaf declare <> StmtLeaf choose, expr: GoVar resultName, exprType: res.exprType, nextId: accProps.nextId + 1 }
              Nothing ->
                { stmts: accProps.stmts, expr: res.expr, exprType: res.exprType, nextId: accProps.nextId }

        Fail msg ->
          ControlExprs.failure context nextId tcoExpr msg

        Branch branches def ->
          ControlExprs.branch translateInContext context nextId branches def

        PrimOp (Op1 op1 e)
          | Just emit <- PrimitiveExprs.unary codegenStateRef modNameStr op1 ->
              let
                resE = translateExpr metadata codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId e
                result = emit { expr: resE.expr, exprType: resE.exprType }
              in
                { stmts: resE.stmts, expr: result.expr, exprType: result.exprType, nextId: resE.nextId }

        PrimOp (Op1 (OpIsTag (Qualified mbMod (Ident tag))) e) ->
          let
            resE = translateExpr metadata codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId e
          in
            AdtExprs.isTag metadata codegenStateRef modNameStr mbMod tag resE

        PrimOp (Op1 OpArrayLength e) ->
          let
            resE = translateExpr metadata codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId e
          in
            case resE.exprType of
              TypeNativeArray _ -> { stmts: resE.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoCall (GoVar "int64") [ GoCall (GoVar "len") [ resE.expr ] ] ], exprType: TypeValue, nextId: resE.nextId }
              _ -> { stmts: resE.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoCall (GoVar "int64") [ GoCall (GoSelector (GoVar "gopurs_runtime") "ArrayLength") [ boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType ] ] ], exprType: TypeValue, nextId: resE.nextId }

        PrimOp (Op2 OpBooleanAnd e1 e2) ->
          ControlExprs.booleanAnd translateInContext context nextId e1 e2

        PrimOp (Op2 OpBooleanOr e1 e2) ->
          ControlExprs.booleanOr translateInContext context nextId e1 e2

        PrimOp (Op2 OpArrayIndex e1 e2) ->
          let
            res1 = translateExpr metadata codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId e1
            res2 = translateExpr metadata codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } res1.nextId e2
            result = case res1.exprType of
              TypeNativeArray innerType -> { expr: boxGoExpr codegenStateRef modNameStr (rawGo (printGoExpr res1.expr <> "[" <> printGoExpr (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64) <> "]")) innerType, exprType: TypeValue }
              _ -> { expr: GoCall (GoSelector (GoVar "gopurs_runtime") "ArrayAccess") [ boxGoExpr codegenStateRef modNameStr res1.expr res1.exprType, GoCall (GoVar "int") [ unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64 ] ], exprType: TypeValue }
          in
            { stmts: res1.stmts <> res2.stmts, expr: result.expr, exprType: result.exprType, nextId: res2.nextId }

        PrimOp (Op2 op2 e1 e2)
          | Just emit <- PrimitiveExprs.binary codegenStateRef modNameStr op2 ->
              let
                res1 = translateExpr metadata codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId e1
                res2 = translateExpr metadata codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } res1.nextId e2
                result = emit { expr: res1.expr, exprType: res1.exprType } { expr: res2.expr, exprType: res2.exprType }
              in
                { stmts: res1.stmts <> res2.stmts, expr: result.expr, exprType: result.exprType, nextId: res2.nextId }

        PrimEffect eff ->
          EffectExprs.primitive translateInContext context nextId eff

        _ -> { stmts: StmtEmpty, expr: GoVar "gopurs_runtime.Value{}", exprType: TypeValue, nextId }
