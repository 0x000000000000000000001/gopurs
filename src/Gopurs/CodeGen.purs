module Gopurs.CodeGen
  ( translate
  , translateWithFunctions
  , PreparedModule
  , prepareModule
  , workerInputs
  , workerOwners
  , globalTypeDemands
  , translatePreparedWithFunctions
  ) where

import Prelude
import Gopurs.Cache.LookupTable as LookupTable
import Control.Alternative (guard)
import PureScript.Backend.Optimizer.Syntax (BackendAccessor(..), BackendOperator(..), BackendOperator1(..), BackendOperator2(..), BackendSyntax(PrimEffect, PrimOp, Branch, Fail, CtorDef, CtorSaturated, Lit, UncurriedAbs, Abs, Update, Accessor, LetRec, Let, EffectDefer, EffectPure, EffectBind, UncurriedEffectAbs, UncurriedEffectApp, UncurriedApp, App, Local, Var, Typed))
import PureScript.Backend.Optimizer.Syntax as Syn
import PureScript.Backend.Optimizer.Convert (BackendModule)
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import Data.String as String
import Data.Array as Array
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Newtype (unwrap)
import PureScript.Backend.Optimizer.CoreFn (Ident(..), Literal(..), Prop(..), Qualified(..))
import PureScript.Backend.Optimizer.CoreFn as CoreFn
import Data.Tuple (Tuple(..))
import Effect.Unsafe (unsafePerformEffect)
import Effect.Ref (Ref)
import Effect.Ref as Ref
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Map as Map
import Data.Set as Set
import Data.Foldable (foldl)
import Gopurs.GoAst (GoExpr(..), GoDecl(..), GoType(..), GoFile, capitalize, sanitizeName)
import Gopurs.GoSymbol (GoReference(..), generated, foreignSymbol)
import Gopurs.Printer (printGoFile)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import Gopurs.ThunkFusion (optimizeThunkProducers)
import Gopurs.FunctionFusion (optimizeFunctionProducers)
import Gopurs.ListFusion as ListFusion
import Gopurs.ListProgressionFold as ListProgressionFold
import Gopurs.ImmediateApplications (optimizeImmediateApplications)
import Gopurs.RegexMatchProjection as RegexMatchProjection
import Gopurs.Ownership as Ownership
import Gopurs.GoTypes (exprTypeToGoType)
import Gopurs.CodegenState (CodegenMetadata, CodegenState, FunctionInfo)
import Gopurs.GoConversions (boxGoExpr, coerceGoExpr, generateReboxFunctions)
import Gopurs.PrimitiveExprs as PrimitiveExprs
import Gopurs.RecordExprs as RecordExprs
import Gopurs.RecordReadGroups as RecordReadGroups
import Gopurs.ClosedIndexedFolds as ClosedIndexedFolds
import Gopurs.CodePointFold as CodePointFold
import Gopurs.AdtExprs as AdtExprs
import Gopurs.ConstructorExprs as ConstructorExprs
import Gopurs.LiteralExprs as LiteralExprs
import Gopurs.ArrayIntrinsics as ArrayIntrinsics
import Gopurs.ClosedDictionaries (cacheClosedDictionaries)
import Gopurs.ClosedScalar as ClosedScalar
import Gopurs.ConstructorFold as ConstructorFold
import Gopurs.BorrowedObjects (borrowReadOnlyObjects)
import Gopurs.RefLoopRegions as RefLoopRegions
import Gopurs.DecoderSchemas (specializeDecoderSchemas)
import Gopurs.CallExprs as CallExprs
import Gopurs.FunctionExprs as FunctionExprs
import Gopurs.BindingExprs as BindingExprs
import Gopurs.ControlExprs as ControlExprs
import Gopurs.EffectExprs as EffectExprs
import Gopurs.FfiBridge.ImmediateEffects as ImmediateEffects
import Gopurs.TypedExprs as TypedExprs
import Gopurs.RecordConsumer as RecordConsumer
import Gopurs.ModuleBindings as ModuleBindings
import Gopurs.ModuleDeclarations as ModuleDeclarations
import Gopurs.GoImports (collectImports)
import Gopurs.ExprAnalysis (getExprType, isEffectNode)
import Gopurs.ExprContext (ModuleFunctions, StmtTree(..), TranslateExpr, childContext)

translate :: CodegenMetadata -> BackendModule -> String
translate metadata inputMod = (translateWithFunctions metadata inputMod).code

translateWithFunctions :: CodegenMetadata -> BackendModule -> { code :: String, functions :: ModuleFunctions, file :: GoFile }
translateWithFunctions metadata inputMod = translatePreparedWithFunctions metadata (prepareModule metadata inputMod)

type PreparedModule =
  { module :: BackendModule
  , declarations :: Array GoDecl
  , functions :: ModuleFunctions
  , workers :: Map.Map String (Maybe FunctionInfo)
  }

-- These transformations consume source metadata, not imported worker signatures
-- or module-local FFI signatures. Prepare once, before computing the emission
-- key, so references introduced by a transformation belong to its proof too.
prepareModule :: CodegenMetadata -> BackendModule -> PreparedModule
prepareModule metadata inputMod =
  let
    projected = ListFusion.optimize (RegexMatchProjection.optimize metadata (CodePointFold.optimize (ClosedIndexedFolds.optimize inputMod)))
    progressions = ListProgressionFold.prepare projected
    constants = ClosedScalar.optimize (RefLoopRegions.optimize progressions.module)
    constructors = ConstructorFold.optimize constants
    schemas = specializeDecoderSchemas metadata (borrowReadOnlyObjects metadata (cacheClosedDictionaries metadata (optimizeImmediateApplications (optimizeFunctionProducers (optimizeThunkProducers constructors)))))
    owned = Ownership.prepare metadata schemas.module
  in
    { module: owned.module
    , declarations: owned.declarations <> schemas.declarations <> progressions.declarations
    , functions: Map.union owned.functions progressions.functions
    , workers: workerInputs metadata.globalFunctions owned.module
    }

-- Imported ordinary and immediate-effect workers use qualified Var targets;
-- local targets use the module's
-- own signatures. Include every such Var, even if it is not called, and record
-- absences in the exact batch snapshot. TCO preserves these qualifications.
workerInputs :: ModuleFunctions -> BackendModule -> Map.Map String (Maybe FunctionInfo)
workerInputs available = foldWorkerReferences
  (\owner name acc -> foldl
    (\entries ident -> let key = unwrap owner <> "." <> ident in Map.insert key (Map.lookup key available) entries)
    acc [ name, ImmediateEffects.key name, ImmediateEffects.factoryKey name, ImmediateEffects.statelessFactoryKey name ])
  Map.empty

-- The emitter can resolve immutable owner contributions without decoding every
-- earlier module's signatures. Keep the exact qualified-Var inventory shared
-- with workerInputs, including FFI targets and references introduced by passes.
workerOwners :: BackendModule -> Set.Set String
workerOwners = foldWorkerReferences (\owner _ -> Set.insert (unwrap owner)) Set.empty

-- Conservative requests are collected explicitly by the emitter. They do not
-- need to be complete: LookupTable checks coverage on every actual lookup.
globalTypeDemands :: BackendModule -> Set.Set String
globalTypeDemands mod = foldl
  (\acc group -> foldl (\entries (Tuple _ value) -> visit entries value) acc group.bindings)
  (Set.fromFoldable
    [ "Data.Argonaut.Decode.Internal.Record.schemaDecoderABI1"
    , "Data.Argonaut.Decode.Parser.textDecoderABI1"
    , "Data.Argonaut.Decode.Internal.Record.borrowObject"
    , RegexMatchProjection.helperName
    , RegexMatchProjection.adapterName
    ]) mod.bindings
  where
  visit acc (NeutralExpr syntax) = case syntax of
    Var (Qualified (Just owner) (Ident name)) -> Set.insert (unwrap owner <> "." <> name) acc
    _ -> foldl visit acc syntax

foldWorkerReferences :: forall a. (CoreFn.ModuleName -> String -> a -> a) -> a -> BackendModule -> a
foldWorkerReferences insert initial mod = foldl
  (\acc group -> foldl (\entries (Tuple _ value) -> visit entries value) acc group.bindings)
  initial
  mod.bindings
  where
  visit acc (NeutralExpr syntax) = case syntax of
    Var (Qualified (Just owner) (Ident name)) | owner /= mod.name ->
      insert owner name acc
    _ -> foldl visit acc syntax

-- Source metadata must be the same as at preparation. Output supplies the
-- freshly parsed local FFI signatures here. Restrict imported signatures to
-- exactly the snapshot covered by the emission key, including on fresh builds.
translatePreparedWithFunctions :: CodegenMetadata -> PreparedModule -> { code :: String, functions :: ModuleFunctions, file :: GoFile }
translatePreparedWithFunctions sourceMetadata prepared =
  let
    metadata = sourceMetadata { globalFunctions = Map.mapMaybe identity prepared.workers }
    mod = prepared.module
    modNameStrOrig = unwrap mod.name
    modNameStr = String.replaceAll (Pattern ".") (Replacement "_") modNameStrOrig

    codegenStateRef :: Ref CodegenState
    codegenStateRef = unsafePerformEffect $
      Ref.new { declarations: ModuleDeclarations.constructors metadata modNameStr mod, globalId: 0, reboxPairs: Set.empty }

    preparedBindings = ModuleBindings.prepare metadata modNameStr mod
    -- The module's FFI functions are keyed by raw identifier; local functions
    -- take precedence when names collide.
    -- Reserved immediate-effect signatures must never become ordinary local
    -- calls after sanitizing their otherwise impossible PureScript names.
    ffiLocalFunctions = Map.fromFoldable $ Array.mapMaybe
      (\(Tuple (Ident name) _) -> map (Tuple (sanitizeName name)) (Map.lookup name metadata.ffiFunctions))
      (Map.toUnfoldable mod.foreign)
    foldDescriptors = Map.fromFoldable (Array.concatMap (\group -> if group.recursive then [] else
      map (\(Tuple (Ident name) value) -> Tuple (sanitizeName name) (ClosedIndexedFolds.describe value)) group.bindings) mod.bindings)
    codePointDescriptors = Map.fromFoldable (Array.concatMap (\group -> if group.recursive then [] else
      map (\(Tuple (Ident name) value) -> Tuple (sanitizeName name) (CodePointFold.describe value)) group.bindings) mod.bindings)
    describedFunctions = Map.fromFoldable (map (\(Tuple name info) -> Tuple name (info
      { arrayFold = if info.fArgs == [ TypeNativeArray TypeInt64 ] && info.fRet == TypeInt64 && info.arity == 1 then fromMaybe Nothing (Map.lookup name foldDescriptors) else Nothing
      , codePointFold = if info.fArgs == [ TypeString ] && info.fRet == TypeInt64 && info.arity == 1 then fromMaybe Nothing (Map.lookup name codePointDescriptors) else Nothing
      })) (Map.toUnfoldable preparedBindings.functions :: Array (Tuple String FunctionInfo)))
    moduleFunctions = Map.union (Map.union prepared.functions describedFunctions) ffiLocalFunctions
    normalFunctions = Map.fromFoldable $ Array.concatMap
      (\group -> Array.mapMaybe
        (\(Tuple (Ident name) _) -> do
          info <- Map.lookup (sanitizeName name) moduleFunctions
          guard (info.arity >= 1 && info.arity <= 10)
          pure (Tuple (unwrap mod.name <> "." <> name) info))
        group.bindings)
      preparedBindings.bindings
    -- Export FFI functions as `<Module>.<ident>` for callers in other modules.
    ffiGlobalFunctions = Map.fromFoldable $ map
      (\(Tuple name info) -> Tuple (unwrap mod.name <> "." <> name) info)
      (Map.toUnfoldable metadata.ffiFunctions :: Array (Tuple String FunctionInfo))

    Tuple allDeclsAst helpers = unsafePerformEffect do
      let
        proofs = RecordConsumer.prepare modNameStr moduleFunctions preparedBindings.bindings
        translator = if Map.isEmpty proofs then translateExpr else translateExprWith proofs
        d = ModuleBindings.declarations translator metadata codegenStateRef modNameStr moduleFunctions preparedBindings.bindings
      h <- Ref.read codegenStateRef
      pure (Tuple d h)

    foreignGetters = map
      (\(Tuple (Ident name) _) -> GoForeignGetter
        { name: generated modNameStr ("Get_" <> modNameStr <> "_" <> sanitizeName name)
        , value: foreignSymbol modNameStr ("_Gopurs_" <> modNameStr <> "_" <> capitalize (sanitizeName name))
        })
      (Map.toUnfoldable mod.foreign)
    declarationGroups = map (map RecordReadGroups.declaration)
      [ allDeclsAst <> prepared.declarations
      , helpers.declarations <> unsafePerformEffect (generateReboxFunctions metadata codegenStateRef modNameStr)
      , foreignGetters
      ]
    goFile = { packageName: "purescript", imports: collectImports declarationGroups, declarationGroups }

  in
    { code: printGoFile goFile
    , functions: Map.union normalFunctions ffiGlobalFunctions
    , file: goFile
    }

translateExpr :: TranslateExpr
translateExpr = translateExprWith Map.empty

translateExprWith :: RecordConsumer.Plans -> TranslateExpr
translateExprWith proofs context@{ metadata, codegenStateRef, modNameStr, bound, options: { inEffectBlock }, mbExpectedExprType } nextId tcoExpr@(TcoExpr _ expr) =
  let
    recur = translateExprWith proofs
    child = childContext context Nothing
    isEff = isEffectNode tcoExpr
  in
    if isEff && not inEffectBlock then
      EffectExprs.wrap recur context nextId tcoExpr
    else
      case expr of
        Typed type_ a ->
          TypedExprs.annotated recur context nextId type_ a
        Var (Qualified mbMn (Ident i)) ->
          let
            safeName = sanitizeName i
            vType = case mbMn of
              Just mn ->
                let
                  modStr = unwrap mn
                in
                  case LookupTable.lookup (modStr <> "." <> i) metadata.globalTypes of
                    Just ty -> exprTypeToGoType metadata.sharedRecords metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr ty
                    Nothing -> TypeValue
              Nothing -> TypeValue
          in
            case mbMn of
              Just mn ->
                let
                  modStr = unwrap mn
                  modPkg = String.replaceAll (Pattern ".") (Replacement "_") modStr
                  rawCall = GoCall (GoRef (SymbolReference (generated modPkg ("Get_" <> modPkg <> "_" <> safeName)))) []
                in
                  { stmts: StmtEmpty, expr: coerceGoExpr codegenStateRef modNameStr rawCall TypeValue vType, exprType: vType, nextId }
              Nothing ->
                let
                  rawCall = GoCall (GoRef (SymbolReference (generated modNameStr ("Get_" <> modNameStr <> "_" <> safeName)))) []
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
          LiteralExprs.array recur context nextId xs

        Lit (LitRecord props) ->
          LiteralExprs.record recur context nextId (getExprType tcoExpr) props

        expr_
          | ( case expr_ of
                App _ _ -> true
                Syn.TypeApp _ _ -> true
                _ -> false
            ) ->
              CallExprs.application recur context nextId tcoExpr

        Abs args body ->
          FunctionExprs.abstraction recur context nextId tcoExpr args body

        UncurriedApp fn args ->
          CallExprs.uncurriedApplication recur context nextId tcoExpr fn args

        UncurriedAbs args body ->
          FunctionExprs.uncurriedAbstraction recur context nextId tcoExpr args body

        UncurriedEffectApp fn args ->
          CallExprs.effectApplication recur context nextId fn args

        UncurriedEffectAbs args body ->
          FunctionExprs.effectAbstraction recur context nextId tcoExpr args body

        EffectBind mbIdent lvl binding body ->
          EffectExprs.bindEffect recur context nextId mbIdent lvl binding body

        EffectPure binding ->
          EffectExprs.pureValue recur context nextId binding

        EffectDefer binding ->
          EffectExprs.defer recur context nextId binding

        Let mbIdent lvl binding body ->
          case RecordConsumer.emit proofs recur context nextId mbIdent lvl binding body of
            Just result -> result
            Nothing -> BindingExprs.nonRecursive recur context nextId mbIdent lvl binding body

        LetRec lvl bindings body ->
          BindingExprs.recursive recur context nextId tcoExpr lvl bindings body

        Accessor obj accessor ->
          let
            resObj = recur child nextId obj
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
            resObj = recur child nextId obj
            accProps = foldl
              ( \acc (Prop key val) ->
                  let
                    resVal = recur child acc.nextId val
                  in
                    { stmts: acc.stmts <> resVal.stmts, exprs: Array.snoc acc.exprs { key, expr: resVal.expr, goType: resVal.exprType }, exprType: TypeValue, nextId: resVal.nextId }
              )
              { stmts: StmtEmpty, exprs: [], exprType: TypeValue, nextId: resObj.nextId }
              props
            resultType = map (exprTypeToGoType metadata.sharedRecords metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr) mbExpectedExprType
            result = RecordExprs.update codegenStateRef modNameStr resultType { expr: resObj.expr, exprType: resObj.exprType } accProps.exprs
          in
            { stmts: resObj.stmts <> accProps.stmts, expr: result.expr, exprType: result.exprType, nextId: accProps.nextId }

        CtorDef _ _ (Ident name) fields ->
          ConstructorExprs.definition context nextId tcoExpr name fields

        CtorSaturated (Qualified mbMod _) _ _ (Ident name) props ->
          ConstructorExprs.saturated recur context nextId tcoExpr mbMod name props

        Fail msg ->
          ControlExprs.failure context nextId tcoExpr msg

        Branch branches def ->
          ControlExprs.branch recur context nextId branches def

        PrimOp (Op1 op1 e)
          | Just emit <- PrimitiveExprs.unary codegenStateRef modNameStr op1 ->
              let
                resE = recur child nextId e
                result = emit { expr: resE.expr, exprType: resE.exprType }
              in
                { stmts: resE.stmts, expr: result.expr, exprType: result.exprType, nextId: resE.nextId }

        PrimOp (Op1 (OpIsTag (Qualified mbMod (Ident tag))) e) ->
          let
            -- A pointer ADT's tag is independent of its type arguments. Keep
            -- the operand's representation: coercing a boxed List here would
            -- copy its entire tail merely to distinguish Nil from Cons.
            recurOperand expected operand@(TcoExpr _ syntax) =
              case syntax of
                Typed type_ inner
                  | TypeStructPointer _ <- exprTypeToGoType metadata.sharedRecords metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr type_ ->
                      recurOperand (Just type_) inner
                _ ->
                  recur (child { mbExpectedExprType = expected }) nextId operand
            resE = recurOperand Nothing e
          in
            AdtExprs.isTag metadata codegenStateRef modNameStr mbMod tag resE

        PrimOp (Op1 OpArrayLength e) ->
          let
            resE = recur child nextId e
          in
            case resE.exprType of
              TypeNativeArray _ -> { stmts: resE.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoCall (GoVar "int64") [ GoCall (GoVar "len") [ resE.expr ] ] ], exprType: TypeValue, nextId: resE.nextId }
              _ -> { stmts: resE.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoCall (GoVar "int64") [ GoCall (GoSelector (GoVar "gopurs_runtime") "ArrayLength") [ boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType ] ] ], exprType: TypeValue, nextId: resE.nextId }

        PrimOp (Op2 OpBooleanAnd e1 e2) ->
          ControlExprs.booleanAnd recur context nextId e1 e2

        PrimOp (Op2 OpBooleanOr e1 e2) ->
          ControlExprs.booleanOr recur context nextId e1 e2

        PrimOp (Op2 OpArrayIndex e1 e2) ->
          ArrayIntrinsics.unsafeIndex recur context nextId e1 e2

        PrimOp (Op2 op2 e1 e2)
          | Just result <- PrimitiveExprs.stringLengthComparison recur context nextId op2 e1 e2 ->
              result

        PrimOp (Op2 op2 e1 e2)
          | Just result <- PrimitiveExprs.modularIntProduct recur context nextId op2 e1 e2 ->
              result
          | Just result <- PrimitiveExprs.boundedIntProduct recur context nextId op2 e1 e2 ->
              result
          | Just result <- PrimitiveExprs.guardedIntProduct recur context nextId op2 e1 e2 ->
              result
          | Just result <- PrimitiveExprs.powerOfTwoModulo recur context nextId op2 e1 e2 ->
              result

        PrimOp (Op2 op2 e1 e2)
          | Just emit <- PrimitiveExprs.binary codegenStateRef modNameStr op2 ->
              let
                res1 = recur child nextId e1
                res2 = recur child res1.nextId e2
                result = emit { expr: res1.expr, exprType: res1.exprType } { expr: res2.expr, exprType: res2.exprType }
              in
                { stmts: res1.stmts <> res2.stmts, expr: result.expr, exprType: result.exprType, nextId: res2.nextId }

        PrimEffect eff ->
          EffectExprs.primitive recur context nextId eff

        _ -> { stmts: StmtEmpty, expr: GoVar "gopurs_runtime.Value{}", exprType: TypeValue, nextId }
