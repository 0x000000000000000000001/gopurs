module Gopurs.CodeGen
  ( translate
  , translateWithFunctions
  ) where

import Prelude
import Control.Alternative (guard)
import PureScript.Backend.Optimizer.Syntax (BackendAccessor(..), BackendEffect(..), BackendOperator(..), BackendOperator1(..), BackendOperator2(..), BackendSyntax(PrimEffect, PrimOp, Fail, Branch, CtorDef, CtorSaturated, Lit, UncurriedAbs, Abs, Update, Accessor, LetRec, Let, EffectDefer, EffectPure, EffectBind, UncurriedEffectAbs, UncurriedEffectApp, UncurriedApp, App, Local, Var, Typed), Pair(..))
import PureScript.Backend.Optimizer.Syntax as Syn
import PureScript.Backend.Optimizer.Convert (BackendModule)
import Data.String as String
import Data.Array as Array
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Newtype (unwrap)
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Ident(..), Literal(..), Prop(..), Qualified(..))
import Data.Tuple (Tuple(..), fst, snd)
import Data.Array.NonEmpty (fromArray, toArray)
import Effect.Unsafe (unsafePerformEffect)
import Effect.Ref (Ref)
import Effect.Ref as Ref
import Data.String.Pattern (Pattern(..), Replacement(..))
import Debug as Debug
import Data.Map as Map
import Data.Set (Set)
import Data.Set as Set
import Data.Foldable (foldMap, foldl)
import Data.Traversable (traverse)
import Gopurs.GoAst (GoDecl, GoExpr(..), GoType(..), goTypeToStr, sanitizeName, capitalize)
import Gopurs.Printer (printGoFile, printGoExpr, printGoDeclVar)
import PureScript.Backend.Optimizer.Codegen.Tco as Tco
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.FreeVars (freeVars, localId)
import PureScript.Backend.Optimizer.FfiSupport (hashString)
import Gopurs.ThunkFusion (optimizeThunkProducers)
import Gopurs.GoTypes (exprTypeToGenericGoType, exprTypeToGoType, instantiateGenericGoType, printExprType, structFieldGoType)
import Gopurs.CodegenState (CodegenMetadata, CodegenState, FunctionInfo)
import Gopurs.GoConversions (boxGoExpr, unboxGoExpr, coerceGoExpr, generateReboxFunctions, getUnboxedADT)
import Gopurs.PrimitiveExprs as PrimitiveExprs
import Gopurs.RecordExprs as RecordExprs
import Gopurs.AdtExprs as AdtExprs
import Gopurs.CallAnalysis (extractUncurriedAbs, isClosureNode)
import Gopurs.CallExprs as CallExprs
import Gopurs.FunctionExprs as FunctionExprs
import Gopurs.ExprAnalysis (bindFieldFunctionParameters, executeIfOpaque, extractExprFuncType, extractFuncType, getExprType, hasTypeVars, isEffectNode, printTcoExprShape, unwrapTcoExpr)
import Gopurs.ExprContext (ExprOptions, ExprResult, LocalEnv, LoopContext, ModuleFunctions, StmtTree(..), TranslateExpr, flattenStmts, wrapInStmts)

foreign import memoizedFreeVarsImpl :: (TcoExpr -> Set String) -> TcoExpr -> Set String

memoizedFreeVars :: TcoExpr -> Set String
memoizedFreeVars = memoizedFreeVarsImpl freeVars

translate :: CodegenMetadata -> BackendModule -> String
translate metadata inputMod = (translateWithFunctions metadata inputMod).code

translateWithFunctions :: CodegenMetadata -> BackendModule -> { code :: String, functions :: ModuleFunctions }
translateWithFunctions { enumAdts, enumCtors, pointerAdtPaths, pointerAdtNodes, pointerAdtLeaves, elidedCtors, ctorTypes, globalTypes, globalFunctions, classDeclsFields } inputMod =

  let
    mod = optimizeThunkProducers inputMod
    modNameStrOrig = unwrap mod.name
    modNameStr = String.replaceAll (Pattern ".") (Replacement "_") modNameStrOrig

    codegenStateRef :: Ref CodegenState
    codegenStateRef = unsafePerformEffect do
      let
        structDecls = Array.concatMap
          ( \decl ->
              Array.concatMap
                ( \ctor ->
                    let
                      fieldTypes = ctor.fields
                      goFieldTypes = map (structFieldGoType pointerAdtPaths enumAdts elidedCtors decl.vars modNameStr) fieldTypes
                      structName = "Constructor_" <> modNameStr <> "_" <> sanitizeName ctor.name

                      typeParams =
                        if Array.length decl.vars > 0 then
                          "[" <> String.joinWith ", " (map (\v -> "T_" <> sanitizeName v <> " any") decl.vars) <> "]"
                        else ""

                      fieldsStr = Array.cons "Rc uint32" (Array.mapWithIndex (\i ty -> "V" <> show i <> " " <> goTypeToStr ty) goFieldTypes)
                      structDecl = "type " <> structName <> typeParams <> " struct {\n\t" <> String.joinWith "\n\t" fieldsStr <> "\n}\n"

                      fullName = unwrap mod.name <> "." <> ctor.name
                      getterDecl = case Map.lookup fullName classDeclsFields of
                        Just info ->
                          let
                            typeParamsGetter =
                              if Array.length decl.vars > 0 then
                                "[" <> String.joinWith ", " (map (const "gopurs_runtime.Value") decl.vars) <> "]"
                              else ""
                            cases = Array.mapWithIndex (\i f -> "\t\tcase \"" <> f.name <> "\": return gopurs_runtime.Box(c.V" <> show i <> ")") info.fields
                            pkgNameStr = String.replaceAll (Pattern ".") (Replacement "_") (unwrap mod.name)
                            baseStructName = "Data_" <> pkgNameStr <> "_" <> sanitizeName ctor.name
                            hashStr = hashString baseStructName
                          in
                            "func init() {\n\tgopurs_runtime.StructGetters[" <> hashStr <> "] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {\n\t\tc := (*" <> structName <> typeParamsGetter <> ")(ptr)\n\t\t_ = c\n\t\tswitch key {\n" <> String.joinWith "\n" cases <> "\n\t\tdefault: panic(\"Key not found in dictionary " <> structName <> ": \" + key)\n\t\t}\n\t}\n}\n"
                        Nothing -> ""

                    in
                      if getterDecl == "" then [ structDecl ] else [ structDecl, getterDecl ]
                )
                decl.constructors
          )
          mod.dataDecls

      -- Rebox requests belong to this invocation of translate.
      Ref.new { decls: [], rawDecls: structDecls, elidedCtors, ctorTypes, pointerAdtPaths, pointerAdtNodes, pointerAdtLeaves, enumAdts, enumCtors, globalTypes, globalFunctions, classDeclsFields, globalId: 0, reboxPairs: Set.empty }

    Tuple _ tcoBindings = foldl
      ( \(Tuple env acc) group ->
          let
            neBindings = fromArray group.bindings

            env' = case neBindings of
              Just ne | group.recursive -> Tco.topLevelTcoEnvGroup mod.name ne <> env
              _ -> env
            tcoBinds = map
              (\(Tuple id val) -> Tuple id (Tco.analyze env' val))
              group.bindings
          in
            Tuple env' (Array.snoc acc { recursive: group.recursive, bindings: tcoBinds })
      )
      (Tuple [] [])
      mod.bindings

    getExprType :: TcoExpr -> ExprType
    getExprType (TcoExpr _ (Typed ty _)) = ty
    getExprType _ = Any -- fallback

    tcoBindingsExpanded = tcoBindings

    unwrapFunc :: Array (Tuple Ident TcoExpr) -> Array (Tuple String FunctionInfo)
    unwrapFunc binds =
      Array.concatMap
        ( \(Tuple (Ident name) val) ->
            case extractUncurriedAbs val of
              Just { args, body } ->
                let
                  typeSig = extractFuncType val
                  fArgsGo = case typeSig of
                    Just { fArgs } -> map (exprTypeToGoType pointerAdtPaths enumAdts elidedCtors modNameStr) (Array.take (Array.length args) fArgs) <> Array.replicate (Array.length args - Array.length fArgs) TypeValue
                    Nothing -> Array.replicate (Array.length args) TypeValue
                  fRetGo = case typeSig of
                    Just { fArgs, fRet } ->
                      if Array.length args < Array.length fArgs then TypeValue
                      else
                        case getUnboxedADT fRet of
                          Just (Tuple adtName adt) -> TypeStructValue adtName adt.signature
                          Nothing -> exprTypeToGoType pointerAdtPaths enumAdts elidedCtors modNameStr fRet
                    Nothing -> TypeValue
                  fullName = "Call_" <> modNameStr <> "_" <> sanitizeName name
                in
                  [ Tuple (sanitizeName name) { fullName, fArgs: fArgsGo, fRet: fRetGo, arity: Array.length args } ]
              Nothing ->
                let
                  fullName = "Call_" <> modNameStr <> "_" <> sanitizeName name
                in
                  [ Tuple (sanitizeName name) { fullName, fArgs: [], fRet: TypeValue, arity: 0 } ]
        )
        binds

    moduleFunctions :: ModuleFunctions
    moduleFunctions = Map.fromFoldable $ Array.concatMap
      ( \group ->
          if group.recursive then
            let
              mutRecBinds = traverse (\(Tuple _ val) -> extractUncurriedAbs val) group.bindings
            in
              case mutRecBinds of
                Just _ -> unwrapFunc group.bindings
                Nothing -> map (\(Tuple (Ident name) _) -> Tuple (sanitizeName name) { fullName: "Call_" <> modNameStr <> "_" <> sanitizeName name, fArgs: [], fRet: TypeValue, arity: 0 }) group.bindings
          else
            unwrapFunc group.bindings
      )
      tcoBindingsExpanded

    Tuple decls helpers = unsafePerformEffect do
      let
        d = Array.concatMap
          ( \group ->
              let
                recVars = if group.recursive then map (\(Tuple (Ident name) _) -> sanitizeName name) group.bindings else []

                processBindingGroup :: Array (Tuple Ident TcoExpr) -> Boolean -> Array GoDecl
                processBindingGroup binds _ =
                  let
                    mutRecBinds = traverse (\(Tuple (Ident name) val) -> map (\abs -> { ident: sanitizeName name, args: abs.args, body: abs.body, fvs: memoizedFreeVars val, val: val }) (extractUncurriedAbs val)) binds
                  in
                    case mutRecBinds of
                      Just fns ->
                        let
                          fnWrapperStmts = map
                            ( \fn ->
                                let
                                  paramsWithTypes = case extractExprFuncType (getExprType fn.val) of
                                    Just { fArgs } -> Array.zipWith (\p goType -> Tuple p goType) fn.args (map (exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr) fArgs <> Array.replicate (Array.length fn.args - Array.length fArgs) TypeValue)
                                    Nothing -> map (\p -> Tuple p TypeValue) fn.args

                                  newBound = foldl (\acc (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } acc) Map.empty paramsWithTypes

                                  isSelfRecursiveLoop = group.recursive && Array.length group.bindings == 1
                                  mbExpectedRet = case extractExprFuncType (getExprType fn.val) of
                                    Just { fArgs, fRet: rt } -> Just case Array.drop (Array.length fn.args) fArgs of
                                      [] -> rt
                                      remaining -> Func remaining rt
                                    Nothing -> Nothing
                                  fRet = case mbExpectedRet of
                                    Just rt -> exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr rt
                                    Nothing -> TypeValue

                                  currentLoopCtx :: LoopContext
                                  currentLoopCtx = if isSelfRecursiveLoop then [ { ident: fn.ident, params: map fst paramsWithTypes, loopParams: map (\p -> fst p <> "_loop") paramsWithTypes, goTypes: map snd paramsWithTypes, fRet } ] else []
                                  resBodyMut = translateExprWithExpectedType codegenStateRef 0 modNameStr recVars moduleFunctions newBound (Just fn.ident) currentLoopCtx { isTail: isSelfRecursiveLoop, inEffectBlock: false } mbExpectedRet 0 fn.body

                                  goName = fn.ident
                                  initVars = Array.concatMap (\(Tuple p goT) -> [ GoRaw ("var " <> p <> " " <> goTypeToStr goT <> " = " <> p <> "_loop"), GoRaw ("_ = " <> p) ]) paramsWithTypes

                                  arity = Array.length fn.args

                                  expectedRetType = case Map.lookup goName moduleFunctions of
                                    Just { fRet, fArgs } | arity < Array.length fArgs -> TypeValue
                                    Just { fRet } -> fRet
                                    Nothing -> TypeValue

                                  goParams = String.joinWith ", " (map (\(Tuple p goT) -> p <> "_loop " <> goTypeToStr goT) paramsWithTypes)

                                  funcExpr =
                                    if arity >= 1 && arity <= 10 then
                                      let
                                        coercedExpr = coerceGoExpr codegenStateRef modNameStr resBodyMut.expr resBodyMut.exprType expectedRetType
                                        bodyStmts = initVars <> flattenStmts resBodyMut.stmts <> [ GoReturn coercedExpr ]
                                        funcBody = if isSelfRecursiveLoop then GoFor goName bodyStmts else GoBlock bodyStmts
                                      in
                                        unsafePerformEffect do
                                          let callFuncDecl = "func Call_" <> modNameStr <> "_" <> goName <> "(" <> goParams <> ") " <> goTypeToStr expectedRetType <> " {\n" <> printGoExpr funcBody <> "\n}"
                                          Ref.modify_ (\r -> r { rawDecls = Array.snoc r.rawDecls callFuncDecl }) codegenStateRef
                                          let wrapperParams = map (\(Tuple p _) -> p <> "_box") paramsWithTypes
                                          let callExpr = GoCall (GoVar ("Call_" <> modNameStr <> "_" <> goName)) (map (\(Tuple p goT) -> coerceGoExpr codegenStateRef modNameStr (GoVar (p <> "_box")) TypeValue goT) paramsWithTypes)
                                          let boxedRes = boxGoExpr codegenStateRef modNameStr callExpr expectedRetType
                                          let wrapperFunc = GoRaw ("func(" <> String.joinWith ", " (map (\p -> p <> " gopurs_runtime.Value") wrapperParams) <> ") gopurs_runtime.Value {\nreturn " <> printGoExpr boxedRes <> "\n}")
                                          let funcWrapperName = if arity == 1 then "gopurs_runtime.Func" else "gopurs_runtime.Func" <> show arity
                                          pure $ GoRaw (funcWrapperName <> "(" <> printGoExpr wrapperFunc <> ")")
                                    else
                                      let
                                        bodyStmts = initVars <> flattenStmts resBodyMut.stmts <> [ GoReturn (boxGoExpr codegenStateRef modNameStr resBodyMut.expr resBodyMut.exprType) ]
                                        funcBody = if isSelfRecursiveLoop then GoFor goName bodyStmts else GoBlock bodyStmts
                                        iife = GoRaw ("func() gopurs_runtime.Value {\n" <> printGoExpr funcBody <> "\n}()")
                                      in
                                        if arity == 0 then
                                          GoFunc "_" TypeValue TypeValue funcBody
                                        else
                                          Array.foldr (\(Tuple p goT) acc -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> p <> "_box gopurs_runtime.Value) gopurs_runtime.Value {\nvar " <> p <> "_loop " <> goTypeToStr goT <> " = " <> printGoExpr (coerceGoExpr codegenStateRef modNameStr (GoVar (p <> "_box")) TypeValue goT) <> "\nreturn " <> printGoExpr acc <> "\n}") ]) iife paramsWithTypes
                                in
                                  { identifier: modNameStr <> "_" <> goName, expression: funcExpr, goType: TypeValue }
                            )
                            fns
                        in
                          fnWrapperStmts
                      Nothing ->
                        Array.concatMap
                          ( \(Tuple (Ident name) expr) ->
                              let
                                res = translateExprWithExpectedType codegenStateRef 0 modNameStr recVars moduleFunctions Map.empty (Just (sanitizeName name)) [] { isTail: false, inEffectBlock: false } (Just (getExprType expr)) 0 expr
                              in
                                [ { identifier: modNameStr <> "_" <> sanitizeName name, expression: wrapInStmts [] res.stmts TypeValue (boxGoExpr codegenStateRef modNameStr res.expr res.exprType), goType: TypeValue } ]
                          )
                          binds
              in
                if group.recursive then
                  processBindingGroup group.bindings true
                else
                  Array.concatMap (\b -> processBindingGroup [ b ] false) group.bindings
          )
          tcoBindingsExpanded
      h <- Ref.read codegenStateRef
      pure (Tuple d h)

    allDeclsAst = decls <> helpers.decls
    declsStr = String.joinWith "\\n" (map printGoDeclVar allDeclsAst) <> "\\n" <> String.joinWith "\\n" helpers.rawDecls

    parts = [ declsStr ]
    usedPkgNames =
      Set.toUnfoldable $ Set.fromFoldable $ Array.mapMaybe
        ( \part ->
            let
              subParts = String.split (Pattern ".") part
            in
              Array.head subParts
        )
        (fromMaybe [] (Array.tail parts)) :: Array String

    goImports = Set.toUnfoldable $ Set.fromFoldable $
      (if Array.length allDeclsAst > 0 || Array.length (Array.fromFoldable mod.foreign) > 0 then [ "gopurs/output/gopurs_runtime" ] else [])
        <> (if Array.length allDeclsAst > 0 then [ "sync" ] else [])
        <> (if String.contains (Pattern "math.") declsStr then [ "math" ] else [])
        <> Array.mapMaybe
          ( \pkg ->
              if pkg /= modNameStr && pkg /= "Prim" && not (String.indexOf (Pattern "Prim_") pkg == Just 0) then Just ("gopurs/output/" <> String.replaceAll (Pattern "_") (Replacement ".") pkg)
              else Nothing
          )
          usedPkgNames

    goFile =
      { packageName: "purescript"
      , imports: goImports
      , decls: allDeclsAst
      , rawDecls: helpers.rawDecls <> unsafePerformEffect (generateReboxFunctions codegenStateRef modNameStr)
      , foreigns: map (\(Tuple (Ident name) type_) -> { pursName: modNameStr <> "_" <> sanitizeName name, goName: "_Gopurs_" <> modNameStr <> "_" <> capitalize (sanitizeName name), exprType: type_ }) (Map.toUnfoldable mod.foreign)
      }
  in
    { code: printGoFile goFile
    , functions: Map.fromFoldable $ Array.concatMap
        (\group -> Array.mapMaybe
          (\(Tuple (Ident name) _) -> do
            info <- Map.lookup (sanitizeName name) moduleFunctions
            guard (info.arity >= 1 && info.arity <= 10)
            pure (Tuple (unwrap mod.name <> "." <> name) info))
          group.bindings)
        tcoBindingsExpanded
    }

translateInContext :: TranslateExpr
translateInContext { codegenStateRef, depth, modNameStr, recVars, moduleFunctions, bound, tcoIdent, loopCtx, options, mbExpectedExprType } nextId expression =
  translateExprWithExpectedType codegenStateRef depth modNameStr recVars moduleFunctions bound tcoIdent loopCtx options mbExpectedExprType nextId expression

translateExpr :: Ref CodegenState -> Int -> String -> Array String -> ModuleFunctions -> LocalEnv -> Maybe String -> LoopContext -> ExprOptions -> Int -> TcoExpr -> ExprResult
translateExpr codegenStateRef depth modNameStr recVars moduleFunctions bound tcoIdent loopCtx options nextId tcoExpr =
  translateExprWithExpectedType codegenStateRef depth modNameStr recVars moduleFunctions bound tcoIdent loopCtx options Nothing nextId tcoExpr

translateExprWithExpectedType :: Ref CodegenState -> Int -> String -> Array String -> ModuleFunctions -> LocalEnv -> Maybe String -> LoopContext -> ExprOptions -> Maybe ExprType -> Int -> TcoExpr -> ExprResult
translateExprWithExpectedType codegenStateRef depth modNameStr recVars moduleFunctions bound tcoIdent loopCtx options@{ isTail, inEffectBlock } mbExpectedExprType nextId tcoExpr@(TcoExpr tcoAnalysis expr) =
  let
    context = { codegenStateRef, depth, modNameStr, recVars, moduleFunctions, bound, tcoIdent, loopCtx, options, mbExpectedExprType }
    elidedCtors = (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors
    isEff = isEffectNode tcoExpr
  in
    if isEff && not inEffectBlock then
      let
        res = translateExpr codegenStateRef depth modNameStr recVars moduleFunctions bound tcoIdent loopCtx { isTail: false, inEffectBlock: true } nextId tcoExpr
        funcExpr = GoRaw ("gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {\n" <> printGoExpr (GoBlock (flattenStmts res.stmts <> [ GoReturn (boxGoExpr codegenStateRef modNameStr res.expr res.exprType) ])) <> "\n})")
      in
        { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: res.nextId }
    else
      case expr of
        Typed type_ a ->
          let
            expectedGoType = exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr type_
            h = unsafePerformEffect (Ref.read codegenStateRef)

          in
            case unwrapTcoExpr a, expectedGoType of
              Lit (LitRecord props), TypeStructPointer baseStructName fullName fullPath tArgs ->
                case Map.lookup fullName h.classDeclsFields of
                  Just classInfo ->
                    let
                      classFields = classInfo.fields
                      typeArgs = case type_ of
                        ADT fullName _ tArgs ->
                          let
                            mapped = map (exprTypeToGoType h.pointerAdtPaths h.enumAdts h.elidedCtors modNameStr) tArgs
                            arity = case Map.lookup fullName h.pointerAdtPaths of
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
                              genericGoType = exprTypeToGenericGoType h.pointerAdtPaths h.enumAdts h.elidedCtors classInfo.vars modNameStr item.field."type"
                              expectedType = instantiateGenericGoType instMap genericGoType
                              expectedExprType = item.field."type"

                              newBound = bindFieldFunctionParameters
                                (exprTypeToGoType h.pointerAdtPaths h.enumAdts h.elidedCtors modNameStr)
                                bound
                                expectedExprType
                                item.val

                              resVal = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions newBound Nothing [] { isTail: false, inEffectBlock: false } acc.nextId item.val
                              coercedExpr = coerceGoExpr codegenStateRef modNameStr resVal.expr resVal.exprType expectedType
                            in
                              { stmts: acc.stmts <> resVal.stmts, exprs: Array.snoc acc.exprs coercedExpr, exprType: TypeValue, nextId: resVal.nextId }
                        )
                        { stmts: StmtEmpty, exprs: [], exprType: TypeValue, nextId }
                        sortedVals
                    in
                      let
                        typeArgsForDict = case expectedGoType of
                          TypeStructPointer _ _ _ args -> args
                          _ -> []
                        monoStructName = case String.indexOf (Pattern "[") fullPath of
                          Just idx -> String.take idx fullPath
                          Nothing -> fullPath
                      in
                        { stmts: accProps.stmts, expr: GoConstructor (hashString baseStructName) monoStructName typeArgsForDict accProps.exprs, exprType: expectedGoType, nextId: accProps.nextId }
                  Nothing ->
                    let
                      res = translateExprWithExpectedType codegenStateRef depth modNameStr recVars moduleFunctions bound tcoIdent loopCtx options (Just type_) nextId a
                    in
                      case res.exprType of
                        TypeStructPointer _ _ _ _ -> res
                        _ ->
                          if expectedGoType == res.exprType then res
                          else if isClosureNode codegenStateRef a then res
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
                  translateExpr codegenStateRef depth modNameStr recVars moduleFunctions bound tcoIdent loopCtx options nextId newA
              LetRec lvl bindings body, _ ->
                let
                  newBody = case body of
                    TcoExpr bAnn _ -> TcoExpr bAnn (Typed type_ body)
                  newLetShape = LetRec lvl bindings newBody
                  newA = case a of
                    TcoExpr ann _ -> TcoExpr ann newLetShape
                in
                  translateExpr codegenStateRef depth modNameStr recVars moduleFunctions bound tcoIdent loopCtx options nextId newA
              _, _ ->
                let
                  res = translateExprWithExpectedType codegenStateRef depth modNameStr recVars moduleFunctions bound tcoIdent loopCtx options (Just type_) nextId a
                  preserveBoxedRecord = case expectedGoType, mbExpectedExprType of
                    TypeRecord _, Just Any -> res.exprType == TypeValue
                    _, _ -> false
                in
                  case res.exprType of
                    TypeStructPointer _ _ _ _ -> res
                    _ ->
                      if expectedGoType == res.exprType || preserveBoxedRecord then res
                      else if isClosureNode codegenStateRef a then res
                      else
                        { stmts: res.stmts, expr: coerceGoExpr codegenStateRef modNameStr res.expr res.exprType expectedGoType, exprType: expectedGoType, nextId: res.nextId }
        Var (Qualified mbMn (Ident i)) ->
          let
            safeName = sanitizeName i
            h = unsafePerformEffect (Ref.read codegenStateRef)
            vType = case mbMn of
              Just mn ->
                let
                  modStr = unwrap mn
                in
                  case Map.lookup (modStr <> "." <> i) h.globalTypes of
                    Just ty -> exprTypeToGoType h.pointerAdtPaths h.enumAdts h.elidedCtors modNameStr ty
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
                  rawCall = Debug.trace ("mbMn is Nothing for safeName: " <> safeName) (\_ -> GoCall (GoVar ("Get_" <> modNameStr <> "_" <> safeName)) [])
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
                    resVal = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } acc.nextId val
                  in
                    { stmts: acc.stmts <> resVal.stmts, exprs: Array.snoc acc.exprs resVal.expr, exprTypes: Array.snoc acc.exprTypes resVal.exprType, nextId: resVal.nextId }
              )
              { stmts: StmtEmpty, exprs: [], exprTypes: [], nextId }
              xs

            mbElemType = Array.head accXs.exprTypes
            isAllSame = Array.all (\t -> Just t == mbElemType) accXs.exprTypes

            expectedElemType = case mbExpectedExprType of
              Just exTy ->
                let
                  h = unsafePerformEffect (Ref.read codegenStateRef)
                in
                  case exprTypeToGenericGoType h.pointerAdtPaths h.enumAdts h.elidedCtors [] modNameStr exTy of
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
                  { stmts: accXs.stmts, expr: GoRaw (goTypeToStr goTypeArr <> "{" <> String.joinWith ", " (map printGoExpr accXs.exprs) <> "}"), exprType: goTypeArr, nextId: accXs.nextId }
              _ ->
                let
                  boxedExprs = Array.zipWith (\expr ty -> boxGoExpr codegenStateRef modNameStr expr ty) accXs.exprs accXs.exprTypes
                in
                  { stmts: accXs.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Array") [ GoRaw ("[]gopurs_runtime.Value{" <> String.joinWith ", " (map printGoExpr boxedExprs) <> "}") ], exprType: TypeValue, nextId: accXs.nextId }

        Lit (LitRecord props) ->
          let
            sortedProps = Array.sortBy (comparing \(Prop k _) -> k) props
            recordInfo = RecordExprs.prepareLiteral codegenStateRef modNameStr (getExprType tcoExpr) mbExpectedExprType

            accProps = foldl
              ( \acc (Prop key val) ->
                  let
                    expectedExprType = fromMaybe Any (Map.lookup key recordInfo.fields)
                    newBound = bindFieldFunctionParameters
                      (\fArgTy -> exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr fArgTy)
                      bound
                      expectedExprType
                      val

                    resVal = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions newBound Nothing [] { isTail: false, inEffectBlock: false } acc.nextId val
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
          let
            stripEffectDefer (TcoExpr a syn) = case unwrapTcoExpr (TcoExpr a syn) of
              EffectDefer inner -> let Tuple _ i = stripEffectDefer inner in Tuple true i
              Abs _ inner -> let Tuple _ i = stripEffectDefer inner in Tuple true i
              Let ident lvl_ val body_ -> let Tuple stripped i = stripEffectDefer body_ in Tuple stripped (TcoExpr a (Let ident lvl_ val i))
              LetRec lvl_ bindings_ body_ -> let Tuple stripped i = stripEffectDefer body_ in Tuple stripped (TcoExpr a (LetRec lvl_ bindings_ i))
              _ -> Tuple false (TcoExpr a syn)
            Tuple wasStripped realBinding = stripEffectDefer binding
            originalName = localId mbIdent lvl
            name = originalName <> "_" <> show nextId
            newBound = Map.insert originalName { name, goType: TypeValue } bound
            resBinding = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: true } (nextId + 1) realBinding
            resBody = translateExprWithExpectedType codegenStateRef (depth + 1) modNameStr recVars moduleFunctions newBound Nothing loopCtx { isTail, inEffectBlock: true } mbExpectedExprType resBinding.nextId body
            bindingExpr = if wasStripped then boxGoExpr codegenStateRef modNameStr resBinding.expr resBinding.exprType else executeIfOpaque realBinding (boxGoExpr codegenStateRef modNameStr resBinding.expr resBinding.exprType)
            bodyExpr = executeIfOpaque body resBody.expr
          in
            { stmts: resBinding.stmts <> StmtLeaf (GoAssign name bindingExpr) <> resBody.stmts, expr: bodyExpr, exprType: resBody.exprType, nextId: resBody.nextId }

        EffectPure binding ->
          translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId binding

        EffectDefer binding ->
          let
            resBinding = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: true } nextId binding
            funcExpr = GoRaw ("gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {\n" <> printGoExpr (GoBlock (flattenStmts resBinding.stmts <> [ GoReturn (boxGoExpr codegenStateRef modNameStr resBinding.expr resBinding.exprType) ])) <> "\n})")
          in
            { stmts: StmtEmpty, expr: funcExpr, exprType: TypeValue, nextId: resBinding.nextId }

        Let mbIdent lvl binding body ->
          let
            originalName = localId mbIdent lvl
            name = originalName <> "_" <> show nextId
            expectedGoTypeFromAst = exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr (getExprType binding)

            mbFunc = extractUncurriedAbs binding
          in
            case mbFunc of
              Just abs | Array.length abs.args > 0 ->
                let
                  fArgsAst = case extractExprFuncType (getExprType binding) of
                    Just { fArgs } -> fArgs
                    Nothing -> []
                  paramsWithTypes = Array.zipWith
                    ( \idStr ty ->
                        Tuple idStr (exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr ty)
                    )
                    abs.args
                    (fArgsAst <> Array.replicate (max 0 (Array.length abs.args - Array.length fArgsAst)) Any)
                  goTypes = map snd paramsWithTypes

                  localModuleFunctions = Map.insert name { fullName: "Call_local_" <> modNameStr <> "_" <> name, fArgs: goTypes, fRet: TypeValue, arity: Array.length abs.args } moduleFunctions
                  declStmts = [ GoRaw ("var Call_local_" <> modNameStr <> "_" <> name <> " func(" <> String.joinWith ", " (map goTypeToStr goTypes) <> ") gopurs_runtime.Value"), GoRaw ("_ = Call_local_" <> modNameStr <> "_" <> name), GoRaw ("var " <> name <> " gopurs_runtime.Value"), GoRaw ("_ = " <> name) ]

                  loopBound = foldl (\acc (Tuple idStr goT) -> Map.insert idStr { name: idStr, goType: goT } acc) bound paramsWithTypes
                  resBodyMut = translateExpr codegenStateRef (depth + 1) modNameStr recVars localModuleFunctions loopBound (Just name) [] { isTail: true, inEffectBlock: false } (nextId + 1) abs.body

                  goParamsNative = String.joinWith ", " (map (\(Tuple p goT) -> p <> "_loop " <> goTypeToStr goT) paramsWithTypes)
                  initVars = Array.concatMap (\(Tuple p goT) -> [ GoRaw ("var " <> p <> " " <> goTypeToStr goT <> " = " <> p <> "_loop"), GoRaw ("_ = " <> p) ]) paramsWithTypes
                  funcBody = GoBlock (initVars <> flattenStmts resBodyMut.stmts <> [ GoReturn (boxGoExpr codegenStateRef modNameStr resBodyMut.expr resBodyMut.exprType) ])
                  nativeAssignment = GoMutate ("Call_local_" <> modNameStr <> "_" <> name) (GoRaw ("func(" <> goParamsNative <> ") gopurs_runtime.Value {\n" <> printGoExpr funcBody <> "\n}"))

                  nativeCallExpr = GoCall (GoVar ("Call_local_" <> modNameStr <> "_" <> name)) (map (\(Tuple p goT) -> coerceGoExpr codegenStateRef modNameStr (GoVar (p <> "_loop_val")) TypeValue goT) paramsWithTypes)
                  funcExpr = Array.foldr (\(Tuple p goT) acc -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> p <> "_loop_val gopurs_runtime.Value) gopurs_runtime.Value {\nreturn " <> printGoExpr acc <> "\n}") ]) nativeCallExpr paramsWithTypes

                  newBound = Map.insert originalName { name, goType: TypeValue } bound
                  resBodyOuter = translateExpr codegenStateRef (depth + 1) modNameStr recVars localModuleFunctions newBound Nothing loopCtx options resBodyMut.nextId body
                in
                  { stmts: foldMap StmtLeaf declStmts <> StmtLeaf nativeAssignment <> StmtLeaf (GoMutate name funcExpr) <> resBodyOuter.stmts, expr: resBodyOuter.expr, exprType: resBodyOuter.exprType, nextId: resBodyOuter.nextId }

              _ ->
                let
                  resBinding = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } (nextId + 1) binding
                  actualGoType = if expectedGoTypeFromAst == TypeValue then resBinding.exprType else expectedGoTypeFromAst
                  newBound = Map.insert originalName { name, goType: actualGoType } bound
                  resBody = translateExprWithExpectedType codegenStateRef (depth + 1) modNameStr recVars moduleFunctions newBound Nothing loopCtx options mbExpectedExprType resBinding.nextId body
                  letStmt =
                    if actualGoType == resBinding.exprType then
                      StmtLeaf (GoAssign name resBinding.expr)
                    else
                      StmtLeaf (GoRaw ("var " <> name <> " " <> goTypeToStr actualGoType <> " = " <> printGoExpr (coerceGoExpr codegenStateRef modNameStr resBinding.expr resBinding.exprType actualGoType)))
                in
                  { stmts: resBinding.stmts <> StmtLeaf (GoRaw ("// TAST (Let): " <> name <> " shape=" <> printTcoExprShape binding <> " bindingType=" <> printExprType (getExprType binding))) <> letStmt <> resBody.stmts, expr: resBody.expr, exprType: resBody.exprType, nextId: resBody.nextId }

        LetRec lvl bindings body ->
          let
            allocRes = foldl
              ( \acc (Tuple (Ident ident) val) ->
                  let
                    oldName = localId (Just (Ident ident)) lvl
                    gId = unsafePerformEffect do
                      curr <- Ref.read codegenStateRef
                      Ref.modify_ (\r -> r { globalId = r.globalId + 1 }) codegenStateRef
                      pure curr.globalId
                    newName = oldName <> "_" <> show acc.nextId <> "_" <> show gId
                    expectedGoTypeFromAst = exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr (getExprType val)
                  in
                    { newBound: Map.insert oldName { name: newName, goType: expectedGoTypeFromAst } acc.newBound, newNames: Array.snoc acc.newNames { oldName, newName }, exprType: TypeValue, nextId: acc.nextId + 1 }
              )
              { newBound: bound, newNames: [], exprType: TypeValue, nextId }
              (toArray bindings)

            combinedRecVars = recVars <> map (\(Tuple (Ident i) _) -> sanitizeName i) (toArray bindings)

            isLoop = (unwrap tcoAnalysis).role.isLoop
            mutRecBinds =
              if isLoop then
                traverse (\(Tuple (Ident name) val) -> map (\abs -> { ident: sanitizeName name, args: abs.args, body: abs.body, fvs: memoizedFreeVars val, val }) (extractUncurriedAbs val)) (toArray bindings)
              else Nothing
          in
            case mutRecBinds of
              Just fns ->
                let
                  loopCtxs = map
                    ( \fn ->
                        let
                          oldName = localId (Just (Ident fn.ident)) lvl
                          boundInfo = fromMaybe { name: oldName, goType: TypeValue } (Map.lookup oldName allocRes.newBound)
                          newName = boundInfo.name
                          goType = boundInfo.goType
                          fRet = case goType of
                            TypeFunc _ r -> r
                            _ -> TypeValue
                          fArgs = case extractExprFuncType (getExprType fn.val) of
                            Just { fArgs: a } -> map (exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr) a
                            Nothing -> []
                          paramsWithTypes = Array.zipWith (\idStr goT -> Tuple idStr goT) fn.args (fArgs <> Array.replicate (max 0 (Array.length fn.args - Array.length fArgs)) TypeValue)
                        in
                          { ident: newName, params: fn.args, loopParams: map (\p -> p <> "_loop") fn.args, goTypes: map snd paramsWithTypes, fRet }
                    )
                    fns

                  combinedLoopCtx = loopCtxs <> loopCtx

                  prepopulatedFunctions = foldl
                    ( \accCtx fn ->
                        let
                          oldName = localId (Just (Ident fn.ident)) lvl
                          boundInfo = fromMaybe { name: oldName, goType: TypeValue } (Map.lookup oldName allocRes.newBound)
                          newName = boundInfo.name
                          fArgs = case extractExprFuncType (getExprType fn.val) of
                            Just { fArgs: a } -> map (exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr) a
                            Nothing -> []
                          fRet = case extractExprFuncType (getExprType fn.val) of
                            Just { fRet: r } -> exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr r
                            Nothing -> TypeValue
                          paramsWithTypes = Array.zipWith (\idStr goT -> Tuple idStr goT) fn.args (fArgs <> Array.replicate (max 0 (Array.length fn.args - Array.length fArgs)) TypeValue)
                        in
                          Map.insert newName { fullName: "Call_local_" <> modNameStr <> "_" <> newName, fArgs: map snd paramsWithTypes, fRet: fRet, arity: Array.length fn.args } accCtx
                    )
                    moduleFunctions
                    fns

                  prepopulatedBound = foldl
                    ( \accCtx fn ->
                        let
                          oldName = localId (Just (Ident fn.ident)) lvl
                          boundInfo = fromMaybe { name: oldName, goType: TypeValue } (Map.lookup oldName allocRes.newBound)
                          newName = boundInfo.name
                          fArgs = case extractExprFuncType (getExprType fn.val) of
                            Just { fArgs: a } -> map (exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr) a
                            Nothing -> []
                          fRet = case extractExprFuncType (getExprType fn.val) of
                            Just { fRet: r } -> exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr r
                            Nothing -> TypeValue
                          paramsWithTypes = Array.zipWith (\idStr goT -> Tuple idStr goT) fn.args (fArgs <> Array.replicate (max 0 (Array.length fn.args - Array.length fArgs)) TypeValue)
                        in
                          Map.insert oldName { name: newName, goType: TypeFunc (map snd paramsWithTypes) fRet } accCtx
                    )
                    allocRes.newBound
                    fns

                  resData = foldl
                    ( \acc fn ->
                        let
                          oldName = localId (Just (Ident fn.ident)) lvl
                          boundInfo = fromMaybe { name: oldName, goType: TypeValue } (Map.lookup oldName prepopulatedBound)
                          newName = boundInfo.name
                          fArgs = case extractExprFuncType (getExprType fn.val) of
                            Just { fArgs: a } -> map (exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr) a
                            Nothing -> []
                          paramsWithTypes = Array.zipWith (\idStr goT -> Tuple idStr goT) fn.args (fArgs <> Array.replicate (max 0 (Array.length fn.args - Array.length fArgs)) TypeValue)
                          currentLoopCtx = [ { ident: newName, params: fn.args, loopParams: map (\p -> p <> "_loop") fn.args, goTypes: map snd paramsWithTypes, fRet: TypeValue } ]
                          loopBound = foldl (\acc2 (Tuple idStr goT) -> Map.insert idStr { name: idStr, goType: goT } acc2) prepopulatedBound paramsWithTypes
                          mbExpectedRet = case extractExprFuncType (getExprType fn.val) of
                            Just { fRet: r } -> Just r
                            Nothing -> Nothing
                          resBodyMut = translateExprWithExpectedType codegenStateRef (depth + 1) modNameStr combinedRecVars acc.moduleFunctions loopBound (Just newName) currentLoopCtx { isTail: true, inEffectBlock: false } mbExpectedRet acc.nextId fn.body
                          trueFRet = resBodyMut.exprType

                          initVars = Array.concatMap (\(Tuple p goT) -> [ GoRaw ("var " <> p <> " " <> goTypeToStr goT <> " = " <> p <> "_loop"), GoRaw ("_ = " <> p) ]) paramsWithTypes

                          funcBody = GoFor newName (initVars <> flattenStmts resBodyMut.stmts <> [ GoReturn resBodyMut.expr ])

                          goParamsNative = String.joinWith ", " (map (\(Tuple p goT) -> p <> "_loop " <> goTypeToStr goT) paramsWithTypes)
                          nativeAssignment = GoMutate ("Call_local_" <> modNameStr <> "_" <> newName) (GoRaw ("func(" <> goParamsNative <> ") " <> goTypeToStr trueFRet <> " {\n" <> printGoExpr funcBody <> "\n}"))

                          nativeCallExpr = GoCall (GoVar ("Call_local_" <> modNameStr <> "_" <> newName)) (map (\(Tuple p goT) -> coerceGoExpr codegenStateRef modNameStr (GoVar (p <> "_loop_val")) TypeValue goT) paramsWithTypes)
                          funcExpr =
                            if Array.null paramsWithTypes then
                              GoFunc "_" TypeValue TypeValue (GoBlock [ GoReturn (boxGoExpr codegenStateRef modNameStr nativeCallExpr trueFRet) ])
                            else
                              Array.foldr (\(Tuple p goT) accExpr -> GoCall (GoSelector (GoVar "gopurs_runtime") "Func") [ GoRaw ("func(" <> p <> "_loop_val gopurs_runtime.Value) gopurs_runtime.Value {\nreturn " <> printGoExpr accExpr <> "\n}") ]) (boxGoExpr codegenStateRef modNameStr nativeCallExpr trueFRet) paramsWithTypes

                          newFunctions = Map.insert newName { fullName: "Call_local_" <> modNameStr <> "_" <> newName, fArgs: map snd paramsWithTypes, fRet: trueFRet, arity: Array.length fn.args } acc.moduleFunctions
                          newBound2 = Map.insert oldName { name: newName, goType: TypeFunc (map snd paramsWithTypes) trueFRet } acc.newBound
                          declStmtsLocal = [ GoRaw ("var Call_local_" <> modNameStr <> "_" <> newName <> " func(" <> String.joinWith ", " (map goTypeToStr (map snd paramsWithTypes)) <> ") " <> goTypeToStr trueFRet), GoRaw ("_ = Call_local_" <> modNameStr <> "_" <> newName), GoRaw ("var " <> newName <> " gopurs_runtime.Value"), GoRaw ("_ = " <> newName) ]
                        in
                          { declarations: acc.declarations <> declStmtsLocal, stmts: acc.stmts <> [ nativeAssignment, GoMutate newName funcExpr ], nextId: resBodyMut.nextId, moduleFunctions: newFunctions, newBound: newBound2 }
                    )
                    { declarations: [], stmts: [], nextId: allocRes.nextId, moduleFunctions: prepopulatedFunctions, newBound: prepopulatedBound }
                    fns

                  resBodyOuter = translateExpr codegenStateRef (depth + 1) modNameStr combinedRecVars resData.moduleFunctions resData.newBound Nothing loopCtx options resData.nextId body
                in
                  -- Every function in the recursive group must be in scope
                  -- before emitting bodies that can refer to its peers.
                  { stmts: foldMap StmtLeaf (resData.declarations <> resData.stmts) <> resBodyOuter.stmts, expr: resBodyOuter.expr, exprType: resBodyOuter.exprType, nextId: resBodyOuter.nextId }

              Nothing ->
                let
                  -- An initializer must not observe Go's zero value for an
                  -- uninitialized recursive binding. Publish a pointer only
                  -- after its complete initializer has run; closures capture
                  -- that cell and can safely read it once initialization ends.
                  initializingBound = foldl
                    ( \acc alloc -> Map.update
                        (\binding -> Just (binding { name = "(*" <> alloc.newName <> "_cell)" }))
                        alloc.oldName
                        acc
                    )
                    allocRes.newBound
                    allocRes.newNames

                  accBindings = foldl
                    ( \acc (Tuple (Tuple (Ident ident) val) alloc) ->
                        let
                          res = translateExpr codegenStateRef (depth + 1) modNameStr combinedRecVars moduleFunctions initializingBound Nothing [] { isTail: false, inEffectBlock: false } acc.nextId val
                          expectedGoType = (fromMaybe { name: alloc.newName, goType: TypeValue } (Map.lookup alloc.oldName allocRes.newBound)).goType
                          assignedVal = coerceGoExpr codegenStateRef modNameStr res.expr res.exprType expectedGoType
                        in
                          { stmts: acc.stmts <> res.stmts
                              <> StmtLeaf (GoMutate alloc.newName assignedVal)
                              <> StmtLeaf (GoMutate (alloc.newName <> "_cell") (GoRaw ("&" <> alloc.newName)))
                          , exprs: Array.snoc acc.exprs { key: alloc.newName, goType: expectedGoType }
                          , exprType: TypeValue
                          , nextId: res.nextId
                          }
                    )
                    { stmts: StmtEmpty, exprs: [], exprType: TypeValue, nextId: allocRes.nextId }
                    (Array.zip (toArray bindings) allocRes.newNames)

                  declStmts = map (\b -> GoRaw ("var " <> b.key <> " " <> goTypeToStr b.goType <> "\n_ = " <> b.key
                    <> "\nvar " <> b.key <> "_cell *" <> goTypeToStr b.goType <> "\n_ = " <> b.key <> "_cell"
                    <> "\n// FALLBACK TCO: isLoop=" <> show isLoop <> " len=" <> show (Array.length (toArray bindings)))) accBindings.exprs

                  resBody = translateExprWithExpectedType codegenStateRef (depth + 1) modNameStr combinedRecVars moduleFunctions allocRes.newBound Nothing loopCtx options mbExpectedExprType accBindings.nextId body
                in
                  { stmts: foldMap StmtLeaf declStmts <> accBindings.stmts <> resBody.stmts, expr: resBody.expr, exprType: resBody.exprType, nextId: resBody.nextId }

        Accessor obj accessor ->
          let
            resObj = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId obj
          in
            case accessor of
              GetProp prop ->
                let
                  result = RecordExprs.getProp codegenStateRef modNameStr prop { expr: resObj.expr, exprType: resObj.exprType }
                in
                  { stmts: resObj.stmts, expr: result.expr, exprType: result.exprType, nextId: resObj.nextId }
              GetIndex idx -> { stmts: resObj.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "ArrayAccess") [ (boxGoExpr codegenStateRef modNameStr resObj.expr resObj.exprType), GoInt idx ], exprType: TypeValue, nextId: resObj.nextId }
              GetCtorField (Qualified mbMod _) _ _ (Ident ctorName) _ idx ->
                let
                  result = AdtExprs.getField codegenStateRef modNameStr elidedCtors
                    { moduleName: mbMod, ctorName, index: idx }
                    { expr: resObj.expr, exprType: resObj.exprType, sourceType: getExprType obj }
                in
                  { stmts: resObj.stmts, expr: result.expr, exprType: result.exprType, nextId: resObj.nextId }

        Update obj props ->
          let
            resObj = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId obj
            accProps = foldl
              ( \acc (Prop key val) ->
                  let
                    resVal = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } acc.nextId val
                  in
                    { stmts: acc.stmts <> resVal.stmts, exprs: Array.snoc acc.exprs { key, expr: resVal.expr, goType: resVal.exprType }, exprType: TypeValue, nextId: resVal.nextId }
              )
              { stmts: StmtEmpty, exprs: [], exprType: TypeValue, nextId: resObj.nextId }
              props
            result = RecordExprs.update codegenStateRef modNameStr { expr: resObj.expr, exprType: resObj.exprType } accProps.exprs
          in
            { stmts: resObj.stmts <> accProps.stmts, expr: result.expr, exprType: result.exprType, nextId: accProps.nextId }

        CtorDef _ _ (Ident name) fields ->
          let
            annotatedType = fromMaybe (getExprType tcoExpr) mbExpectedExprType
            ctorType = case extractExprFuncType annotatedType of
              Just { fRet } -> fRet
              Nothing -> annotatedType
            result = AdtExprs.definition codegenStateRef modNameStr name fields ctorType
          in
            { stmts: StmtEmpty, expr: result.expr, exprType: result.exprType, nextId }

        CtorSaturated (Qualified mbMod _) _ _ (Ident name) props ->
          let
            helpers = unsafePerformEffect (Ref.read codegenStateRef)

            ctorType = case getExprType tcoExpr of
              Any -> fromMaybe Any mbExpectedExprType
              ty ->
                if hasTypeVars ty then
                  case mbExpectedExprType of
                    Just expectedTy | not (hasTypeVars expectedTy) -> expectedTy
                    _ -> ty
                else ty

            prepared = AdtExprs.prepareSaturated codegenStateRef modNameStr mbMod name ctorType

            accProps = foldl
              ( \acc (Tuple _ val) ->
                  let
                    { exprType: expectedExprType, goType: expectedType } = AdtExprs.saturatedFieldType prepared acc.fieldIdx

                    newBound = case unwrapTcoExpr val, extractExprFuncType expectedExprType of
                      Abs args _, Just { fArgs } ->
                        let
                          paramsWithTypes = Array.zipWith (\(Tuple mbI lvl) fArgTy -> Tuple (localId mbI lvl) (exprTypeToGoType helpers.pointerAdtPaths helpers.enumAdts helpers.elidedCtors modNameStr fArgTy)) (toArray args) (fArgs <> Array.replicate (Array.length (toArray args) - Array.length fArgs) Any)
                        in
                          foldl (\b (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } b) bound paramsWithTypes
                      UncurriedAbs args _, Just { fArgs } ->
                        let
                          paramsWithTypes = Array.zipWith (\(Tuple mbI lvl) fArgTy -> Tuple (localId mbI lvl) (exprTypeToGoType helpers.pointerAdtPaths helpers.enumAdts helpers.elidedCtors modNameStr fArgTy)) args (fArgs <> Array.replicate (Array.length args - Array.length fArgs) Any)
                        in
                          foldl (\b (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } b) bound paramsWithTypes
                      _, _ -> bound

                    resVal = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions newBound Nothing [] { isTail: false, inEffectBlock: false } acc.nextId val
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
                  declare = GoRaw ("var " <> resultName <> " " <> goTypeToStr res.exprType)
                  choose = GoIfElse condition [ GoMutate resultName source ] [ GoMutate resultName res.expr ]
                in
                  { stmts: StmtLeaf declare <> StmtLeaf choose, expr: GoVar resultName, exprType: res.exprType, nextId: accProps.nextId + 1 }
              Nothing ->
                { stmts: accProps.stmts, expr: res.expr, exprType: res.exprType, nextId: accProps.nextId }

        Fail msg ->
          let
            expectedGoType = exprTypeToGoType (unsafePerformEffect (Ref.read codegenStateRef)).pointerAdtPaths (unsafePerformEffect (Ref.read codegenStateRef)).enumAdts (unsafePerformEffect (Ref.read codegenStateRef)).elidedCtors modNameStr
              ( case getExprType tcoExpr of
                  Any -> fromMaybe Any mbExpectedExprType
                  ty -> ty
              )
            expectedGoTypeStr = goTypeToStr expectedGoType
          in
            { stmts: StmtEmpty, expr: GoRaw ("func() " <> expectedGoTypeStr <> " { panic(" <> printGoExpr (GoString msg) <> ") }()"), exprType: expectedGoType, nextId }

        Branch branches def ->
          let
            resDef = translateExprWithExpectedType codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing loopCtx { isTail, inEffectBlock: false } mbExpectedExprType nextId def

            computedBranches = foldl
              ( \acc (Pair condExpr bodyExpr) ->
                  let
                    resCond = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } acc.nextId condExpr
                    resBody = translateExprWithExpectedType codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing loopCtx { isTail, inEffectBlock: false } mbExpectedExprType resCond.nextId bodyExpr
                  in
                    { nextId: resBody.nextId, results: acc.results <> [ { cond: resCond, body: resBody } ] }
              )
              { nextId: resDef.nextId, results: [] }
              (toArray branches)

            isFailNode = case unwrapTcoExpr def of
              Fail _ -> true
              _ -> false

            allTypes = (if isFailNode then [] else [ resDef.exprType ]) <> map (\r -> r.body.exprType) computedBranches.results

            hasTypeValue = Array.any
              ( \t -> case t of
                  TypeValue -> true
                  _ -> false
              )
              allTypes

            expectedGoType =
              if hasTypeValue then TypeValue
              else
                let
                  nubbed = Array.nub (map goTypeToStr allTypes)
                in
                  if Array.length nubbed == 1 then fromMaybe TypeValue (Array.head allTypes)
                  else TypeValue

            tmpVar = "__t" <> show computedBranches.nextId
            declTmp = StmtLeaf (GoRaw ("var " <> tmpVar <> " " <> goTypeToStr expectedGoType))
            labelName = "end_branch_" <> show computedBranches.nextId

            buildIfs = foldl
              ( \acc r ->
                  let
                    goIf = GoIfElse (unboxGoExpr codegenStateRef modNameStr r.cond.expr r.cond.exprType TypeBool) (flattenStmts r.body.stmts <> [ GoMutate tmpVar (coerceGoExpr codegenStateRef modNameStr r.body.expr r.body.exprType expectedGoType), GoRaw ("goto " <> labelName) ]) []
                  in
                    acc <> StmtLeaf (GoRaw "{") <> r.cond.stmts <> StmtLeaf goIf <> StmtLeaf (GoRaw "}")
              )
              StmtEmpty
              computedBranches.results
          in
            { stmts: declTmp <> buildIfs <> StmtLeaf (GoRaw "{") <> resDef.stmts <> StmtLeaf (GoMutate tmpVar (coerceGoExpr codegenStateRef modNameStr resDef.expr resDef.exprType expectedGoType)) <> StmtLeaf (GoRaw "}") <> StmtLeaf (GoRaw (labelName <> ":")), expr: GoVar tmpVar, exprType: expectedGoType, nextId: computedBranches.nextId + 1 }

        PrimOp (Op1 op1 e)
          | Just emit <- PrimitiveExprs.unary codegenStateRef modNameStr op1 ->
              let
                resE = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId e
                result = emit { expr: resE.expr, exprType: resE.exprType }
              in
                { stmts: resE.stmts, expr: result.expr, exprType: result.exprType, nextId: resE.nextId }

        PrimOp (Op1 (OpIsTag (Qualified mbMod (Ident tag))) e) ->
          let
            resE = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId e
          in
            AdtExprs.isTag codegenStateRef modNameStr mbMod tag resE

        PrimOp (Op1 OpArrayLength e) ->
          let
            resE = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId e
          in
            case resE.exprType of
              TypeNativeArray _ -> { stmts: resE.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoCall (GoVar "int64") [ GoCall (GoVar "len") [ resE.expr ] ] ], exprType: TypeValue, nextId: resE.nextId }
              _ -> { stmts: resE.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Int") [ GoCall (GoVar "int64") [ GoCall (GoSelector (GoVar "gopurs_runtime") "ArrayLength") [ boxGoExpr codegenStateRef modNameStr resE.expr resE.exprType ] ] ], exprType: TypeValue, nextId: resE.nextId }

        PrimOp (Op2 OpBooleanAnd e1 e2) ->
          let
            res1 = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId e1
            res2 = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } res1.nextId e2
            isEmptyStmts StmtEmpty = true
            isEmptyStmts (StmtAppend s1 s2) = isEmptyStmts s1 && isEmptyStmts s2
            isEmptyStmts _ = false
          in
            if isEmptyStmts res2.stmts then
              { expr: GoBinOp "&&" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeBool) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeBool), exprType: TypeBool, stmts: res1.stmts <> res2.stmts, nextId: res2.nextId }
            else
              let
                tmpVar = "__t_and_" <> show res2.nextId
                declTmp = StmtLeaf (GoRaw ("var " <> tmpVar <> " bool = false\nif " <> printGoExpr (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeBool) <> " {\n"))
                assignTmp = StmtLeaf (GoRaw (tmpVar <> " = " <> printGoExpr (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeBool) <> "\n}"))
              in
                { expr: GoRaw tmpVar, exprType: TypeBool, stmts: res1.stmts <> declTmp <> res2.stmts <> assignTmp, nextId: res2.nextId + 1 }

        PrimOp (Op2 OpBooleanOr e1 e2) ->
          let
            res1 = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId e1
            res2 = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } res1.nextId e2
            isEmptyStmts StmtEmpty = true
            isEmptyStmts (StmtAppend s1 s2) = isEmptyStmts s1 && isEmptyStmts s2
            isEmptyStmts _ = false
          in
            if isEmptyStmts res2.stmts then
              { expr: GoBinOp "||" (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeBool) (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeBool), exprType: TypeBool, stmts: res1.stmts <> res2.stmts, nextId: res2.nextId }
            else
              let
                tmpVar = "__t_or_" <> show res2.nextId
                declTmp = StmtLeaf (GoRaw ("var " <> tmpVar <> " bool = true\nif !(" <> printGoExpr (unboxGoExpr codegenStateRef modNameStr res1.expr res1.exprType TypeBool) <> ") {\n"))
                assignTmp = StmtLeaf (GoRaw (tmpVar <> " = " <> printGoExpr (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeBool) <> "\n}"))
              in
                { expr: GoRaw tmpVar, exprType: TypeBool, stmts: res1.stmts <> declTmp <> res2.stmts <> assignTmp, nextId: res2.nextId + 1 }

        PrimOp (Op2 OpArrayIndex e1 e2) ->
          let
            res1 = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId e1
            res2 = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } res1.nextId e2
            result = case res1.exprType of
              TypeNativeArray innerType -> { expr: boxGoExpr codegenStateRef modNameStr (GoRaw (printGoExpr res1.expr <> "[" <> printGoExpr (unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64) <> "]")) innerType, exprType: TypeValue }
              _ -> { expr: GoCall (GoSelector (GoVar "gopurs_runtime") "ArrayAccess") [ boxGoExpr codegenStateRef modNameStr res1.expr res1.exprType, GoCall (GoVar "int") [ unboxGoExpr codegenStateRef modNameStr res2.expr res2.exprType TypeInt64 ] ], exprType: TypeValue }
          in
            { stmts: res1.stmts <> res2.stmts, expr: result.expr, exprType: result.exprType, nextId: res2.nextId }

        PrimOp (Op2 op2 e1 e2)
          | Just emit <- PrimitiveExprs.binary codegenStateRef modNameStr op2 ->
              let
                res1 = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId e1
                res2 = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } res1.nextId e2
                result = emit { expr: res1.expr, exprType: res1.exprType } { expr: res2.expr, exprType: res2.exprType }
              in
                { stmts: res1.stmts <> res2.stmts, expr: result.expr, exprType: result.exprType, nextId: res2.nextId }

        PrimEffect eff -> case eff of
          EffectRefNew a ->
            let
              resA = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId a
              refIdent = "__local_ref_" <> show resA.nextId
              declStmt = GoAssign refIdent (boxGoExpr codegenStateRef modNameStr resA.expr resA.exprType)
              ifaceIdent = "__local_iface_" <> show resA.nextId
              ifaceStmt = GoRaw ("var " <> ifaceIdent <> " interface{} = " <> refIdent)
            in
              { stmts: resA.stmts <> StmtLeaf declStmt <> StmtLeaf ifaceStmt
              , expr: GoRaw ("gopurs_runtime.Any(&" <> ifaceIdent <> ")")
              , exprType: TypeValue
              , nextId: resA.nextId + 1
              }
          EffectRefRead a ->
            let
              resA = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId a
            in
              { stmts: resA.stmts
              , expr: GoRaw ("(*(" <> printGoExpr resA.expr <> ".PtrVal().(*interface{}))).(gopurs_runtime.Value)")
              , exprType: TypeValue
              , nextId: resA.nextId
              }
          EffectRefWrite ref val ->
            let
              resRef = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } nextId ref
              resVal = translateExpr codegenStateRef (depth + 1) modNameStr recVars moduleFunctions bound Nothing [] { isTail: false, inEffectBlock: false } resRef.nextId val
              writeStmt = GoRaw ("*(" <> printGoExpr resRef.expr <> ".PtrVal().(*interface{})) = " <> printGoExpr (boxGoExpr codegenStateRef modNameStr resVal.expr resVal.exprType))
            in
              { stmts: resRef.stmts <> resVal.stmts <> StmtLeaf writeStmt
              , expr: boxGoExpr codegenStateRef modNameStr resVal.expr resVal.exprType
              , exprType: TypeValue
              , nextId: resVal.nextId
              }

        _ -> { stmts: StmtEmpty, expr: GoVar "gopurs_runtime.Value{}", exprType: TypeValue, nextId }
