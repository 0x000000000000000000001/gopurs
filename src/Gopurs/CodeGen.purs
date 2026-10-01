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
import PureScript.Backend.Optimizer.CoreFn (Ident(..), Literal(..), Prop(..), Qualified(..))
import Data.Tuple (Tuple(..))
import Effect.Unsafe (unsafePerformEffect)
import Effect.Ref (Ref)
import Effect.Ref as Ref
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Map as Map
import Data.Set as Set
import Data.Foldable (foldl)
import Gopurs.GoAst (GoExpr(..), GoDecl(..), GoType(..), capitalize, sanitizeName)
import Gopurs.Printer (printGoFile)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import Gopurs.ThunkFusion (optimizeThunkProducers)
import Gopurs.FunctionFusion (optimizeFunctionProducers)
import Gopurs.ImmediateApplications (optimizeImmediateApplications)
import Gopurs.Ownership as Ownership
import Gopurs.GoTypes (exprTypeToGoType)
import Gopurs.CodegenState (CodegenMetadata, CodegenState, FunctionInfo)
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr, generateReboxFunctions)
import Gopurs.PrimitiveExprs as PrimitiveExprs
import Gopurs.RecordExprs as RecordExprs
import Gopurs.AdtExprs as AdtExprs
import Gopurs.ConstructorExprs as ConstructorExprs
import Gopurs.LiteralExprs as LiteralExprs
import Gopurs.ArrayIntrinsics as ArrayIntrinsics
import Gopurs.ClosedDictionaries (cacheClosedDictionaries)
import Gopurs.BorrowedObjects (borrowReadOnlyObjects)
import Gopurs.DecoderSchemas (specializeDecoderSchemas)
import Gopurs.CallExprs as CallExprs
import Gopurs.FunctionExprs as FunctionExprs
import Gopurs.BindingExprs as BindingExprs
import Gopurs.ControlExprs as ControlExprs
import Gopurs.EffectExprs as EffectExprs
import Gopurs.TypedExprs as TypedExprs
import Gopurs.ModuleBindings as ModuleBindings
import Gopurs.ModuleDeclarations as ModuleDeclarations
import Gopurs.GoImports (collectImports)
import Gopurs.ExprAnalysis (getExprType, isEffectNode)
import Gopurs.ExprContext (ModuleFunctions, StmtTree(..), TranslateExpr, childContext)

translate :: CodegenMetadata -> BackendModule -> String
translate metadata inputMod = (translateWithFunctions metadata inputMod).code

translateWithFunctions :: CodegenMetadata -> BackendModule -> { code :: String, functions :: ModuleFunctions }
translateWithFunctions metadata inputMod =

  let
    schemas = specializeDecoderSchemas metadata (borrowReadOnlyObjects metadata (cacheClosedDictionaries metadata (optimizeImmediateApplications (optimizeFunctionProducers (optimizeThunkProducers inputMod)))))
    owned = Ownership.prepare metadata schemas.module
    mod = owned.module
    modNameStrOrig = unwrap mod.name
    modNameStr = String.replaceAll (Pattern ".") (Replacement "_") modNameStrOrig

    codegenStateRef :: Ref CodegenState
    codegenStateRef = unsafePerformEffect $
      Ref.new { declarations: ModuleDeclarations.constructors metadata modNameStr mod, globalId: 0, reboxPairs: Set.empty }

    preparedBindings = ModuleBindings.prepare metadata modNameStr mod
    -- Les FFI du module sont connues par ident brut ; les fonctions locales
    -- gardent la priorité en cas d'homonymie.
    ffiLocalFunctions = Map.fromFoldable (map (\(Tuple k v) -> Tuple (sanitizeName k) v) (Map.toUnfoldable metadata.ffiFunctions :: Array (Tuple String FunctionInfo)))
    moduleFunctions = Map.union (Map.union owned.functions preparedBindings.functions) ffiLocalFunctions
    normalFunctions = Map.fromFoldable $ Array.concatMap
      (\group -> Array.mapMaybe
        (\(Tuple (Ident name) _) -> do
          info <- Map.lookup (sanitizeName name) moduleFunctions
          guard (info.arity >= 1 && info.arity <= 10)
          pure (Tuple (unwrap mod.name <> "." <> name) info))
        group.bindings)
      preparedBindings.bindings
    -- FFI exportées sous `<Module>.<ident>` pour les appelants d'autres modules.
    ffiGlobalFunctions = Map.fromFoldable $ Array.mapMaybe
      (\(Tuple (Ident name) _) -> map (Tuple (unwrap mod.name <> "." <> name)) (Map.lookup name metadata.ffiFunctions))
      (Map.toUnfoldable mod.foreign)

    Tuple allDeclsAst helpers = unsafePerformEffect do
      let
        d = ModuleBindings.declarations translateExpr metadata codegenStateRef modNameStr moduleFunctions preparedBindings.bindings
      h <- Ref.read codegenStateRef
      pure (Tuple d h)

    foreignGetters = map
      (\(Tuple (Ident name) _) -> GoForeignGetter
        { name: "Get_" <> modNameStr <> "_" <> sanitizeName name
        , value: "_Gopurs_" <> modNameStr <> "_" <> capitalize (sanitizeName name)
        })
      (Map.toUnfoldable mod.foreign)
    declarationGroups =
      [ allDeclsAst <> owned.declarations <> schemas.declarations
      , helpers.declarations <> unsafePerformEffect (generateReboxFunctions metadata codegenStateRef modNameStr)
      , foreignGetters
      ]
    goFile = { packageName: "purescript", imports: collectImports declarationGroups, declarationGroups }

  in
    { code: printGoFile goFile
    , functions: Map.union normalFunctions ffiGlobalFunctions
    }

translateExpr :: TranslateExpr
translateExpr context@{ metadata, codegenStateRef, modNameStr, bound, options: { inEffectBlock }, mbExpectedExprType } nextId tcoExpr@(TcoExpr _ expr) =
  let
    child = childContext context Nothing
    isEff = isEffectNode tcoExpr
  in
    if isEff && not inEffectBlock then
      EffectExprs.wrap translateExpr context nextId tcoExpr
    else
      case expr of
        Typed type_ a ->
          TypedExprs.annotated translateExpr context nextId type_ a
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
          LiteralExprs.array translateExpr context nextId xs

        Lit (LitRecord props) ->
          LiteralExprs.record translateExpr context nextId (getExprType tcoExpr) props

        expr_
          | ( case expr_ of
                App _ _ -> true
                Syn.TypeApp _ _ -> true
                _ -> false
            ) ->
              CallExprs.application translateExpr context nextId tcoExpr

        Abs args body ->
          FunctionExprs.abstraction translateExpr context nextId tcoExpr args body

        UncurriedApp fn args ->
          CallExprs.uncurriedApplication translateExpr context nextId tcoExpr fn args

        UncurriedAbs args body ->
          FunctionExprs.uncurriedAbstraction translateExpr context nextId tcoExpr args body

        UncurriedEffectApp fn args ->
          CallExprs.effectApplication translateExpr context nextId fn args

        UncurriedEffectAbs args body ->
          FunctionExprs.effectAbstraction translateExpr context nextId tcoExpr args body

        EffectBind mbIdent lvl binding body ->
          EffectExprs.bindEffect translateExpr context nextId mbIdent lvl binding body

        EffectPure binding ->
          EffectExprs.pureValue translateExpr context nextId binding

        EffectDefer binding ->
          EffectExprs.defer translateExpr context nextId binding

        Let mbIdent lvl binding body ->
          BindingExprs.nonRecursive translateExpr context nextId mbIdent lvl binding body

        LetRec lvl bindings body ->
          BindingExprs.recursive translateExpr context nextId tcoExpr lvl bindings body

        Accessor obj accessor ->
          let
            resObj = translateExpr child nextId obj
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
            resObj = translateExpr child nextId obj
            accProps = foldl
              ( \acc (Prop key val) ->
                  let
                    resVal = translateExpr child acc.nextId val
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
          ConstructorExprs.definition context nextId tcoExpr name fields

        CtorSaturated (Qualified mbMod _) _ _ (Ident name) props ->
          ConstructorExprs.saturated translateExpr context nextId tcoExpr mbMod name props

        Fail msg ->
          ControlExprs.failure context nextId tcoExpr msg

        Branch branches def ->
          ControlExprs.branch translateExpr context nextId branches def

        PrimOp (Op1 op1 e)
          | Just emit <- PrimitiveExprs.unary codegenStateRef modNameStr op1 ->
              let
                resE = translateExpr child nextId e
                result = emit { expr: resE.expr, exprType: resE.exprType }
              in
                { stmts: resE.stmts, expr: result.expr, exprType: result.exprType, nextId: resE.nextId }

        PrimOp (Op1 (OpIsTag (Qualified mbMod (Ident tag))) e) ->
          let
            -- A pointer ADT's tag is independent of its type arguments. Keep
            -- the operand's representation: coercing a boxed List here would
            -- copy its entire tail merely to distinguish Nil from Cons.
            translateOperand expected operand@(TcoExpr _ syntax) =
              case syntax of
                Typed type_ inner
                  | TypeStructPointer _ <- exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr type_ ->
                      translateOperand (Just type_) inner
                _ ->
                  translateExpr (child { mbExpectedExprType = expected }) nextId operand
            resE = translateOperand Nothing e
          in
            AdtExprs.isTag metadata codegenStateRef modNameStr mbMod tag resE

        PrimOp (Op1 OpArrayLength e) ->
          let
            resE = translateExpr child nextId e
          in
            case resE.exprType of
              TypeNativeArray _ -> { stmts: resE.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoCall (GoVar "int64") [ GoCall (GoVar "len") [ resE.expr ] ] ], exprType: TypeValue, nextId: resE.nextId }
              _ -> { stmts: resE.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoCall (GoVar "int64") [ GoCall (GoSelector (GoVar "gopurs_runtime") "ArrayLength") [ boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType ] ] ], exprType: TypeValue, nextId: resE.nextId }

        PrimOp (Op2 OpBooleanAnd e1 e2) ->
          ControlExprs.booleanAnd translateExpr context nextId e1 e2

        PrimOp (Op2 OpBooleanOr e1 e2) ->
          ControlExprs.booleanOr translateExpr context nextId e1 e2

        PrimOp (Op2 OpArrayIndex e1 e2) ->
          ArrayIntrinsics.unsafeIndex translateExpr context nextId e1 e2

        PrimOp (Op2 op2 e1 e2)
          | Just emit <- PrimitiveExprs.binary codegenStateRef modNameStr op2 ->
              let
                res1 = translateExpr child nextId e1
                res2 = translateExpr child res1.nextId e2
                result = emit { expr: res1.expr, exprType: res1.exprType } { expr: res2.expr, exprType: res2.exprType }
              in
                { stmts: res1.stmts <> res2.stmts, expr: result.expr, exprType: result.exprType, nextId: res2.nextId }

        PrimEffect eff ->
          EffectExprs.primitive translateExpr context nextId eff

        _ -> { stmts: StmtEmpty, expr: GoVar "gopurs_runtime.Value{}", exprType: TypeValue, nextId }
