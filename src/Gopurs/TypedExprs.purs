module Gopurs.TypedExprs (annotated) where

import Prelude

import Data.Array as Array
import Data.Foldable (foldl)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Tuple (Tuple(..))
import Gopurs.CallAnalysis (isClosureNode)
import Gopurs.ExprAnalysis (bindFieldFunctionParameters, unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, StmtTree(..), TranslateExpr, childContext)
import Gopurs.GoAst (GoExpr(..), GoType(..), StructPointer)
import Gopurs.GoConversions (coerceGoExpr, getUnboxedADT)
import Gopurs.GoTypes (exprTypeToGenericGoType, exprTypeToGoType, instantiateGenericGoType)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Literal(..), Prop(..))
import PureScript.Backend.Optimizer.FfiSupport (hashString)
import PureScript.Backend.Optimizer.Syntax (BackendSyntax(Let, LetRec, Lit, Local, Typed))

annotated :: TranslateExpr -> ExprContext -> Int -> ExprType -> TcoExpr -> ExprResult
annotated translate context@{ metadata, modNameStr, mbExpectedExprType } nextId type_ expression =
  let
    expectedGoType = exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr type_
  in
    case unwrapTcoExpr expression, expectedGoType of
      Lit (LitRecord props), TypeStructPointer pointer ->
        case classDictionary translate context nextId type_ pointer props of
          Just result -> result
          Nothing -> coerceResult context expression expectedGoType (translateValue expression)
      Let ident level value body, _ ->
        translateBinding (Let ident level value (annotate type_ body))
      LetRec level bindings body, _ ->
        translateBinding (LetRec level bindings (annotate type_ body))
      _, _ ->
        let
          result = translateValue expression
          preserveBoxedRecord = case expectedGoType, mbExpectedExprType of
            TypeRecord _, Just Any -> result.exprType == TypeValue
            _, _ -> false
          -- Open-row workers prove that every use is a known-field read.
          -- Keep their projected native parameter until that read occurs.
          preserveNativeRecord = case type_, unwrapTcoExpr expression, result.exprType of
            Record (Row _ (Just _)), Local _ _, TypeRecord _ -> true
            _, _, _ -> false
          -- An annotation describes the value, not its allocation. A native
          -- sum stays native unless a more precise pointer layout is required.
          preserveNativeSum = expectedGoType == TypeValue && case result.exprType, getUnboxedADT type_ of
            TypeStructValue actual _, Just (Tuple expected _) -> actual == expected
            _, _ -> false
        in
          if preserveBoxedRecord || preserveNativeRecord || preserveNativeSum then result
          else coerceResult context expression expectedGoType result
  where
  -- Translate only the selected branch: translation can register helpers in
  -- codegenStateRef even though the callback has a pure interface.
  translateValue value = translate (context { mbExpectedExprType = Just type_ }) nextId value

  -- Move the annotation onto a binding's result while preserving its TCO
  -- annotation, tail position and effect context.
  translateBinding shape = case expression of
    TcoExpr annotation _ ->
      translate (context { mbExpectedExprType = Nothing }) nextId (TcoExpr annotation shape)

annotate :: ExprType -> TcoExpr -> TcoExpr
annotate type_ expression@(TcoExpr annotation _) = TcoExpr annotation (Typed type_ expression)

coerceResult :: ExprContext -> TcoExpr -> GoType -> ExprResult -> ExprResult
coerceResult { metadata, codegenStateRef, modNameStr } expression expected result =
  case result.exprType of
    TypeStructPointer _ -> result
    _ ->
      if expected == result.exprType || isClosureNode metadata expression then result
      else result
        { expr = coerceGoExpr codegenStateRef modNameStr result.expr result.exprType expected
        , exprType = expected
        }

classDictionary :: TranslateExpr -> ExprContext -> Int -> ExprType -> StructPointer -> Array (Prop TcoExpr) -> Maybe ExprResult
classDictionary translate context@{ metadata, codegenStateRef, modNameStr, bound } nextId type_ pointer props = do
  classInfo <- Map.lookup pointer.fullName metadata.classDeclsFields
  let
    toGoType = exprTypeToGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors modNameStr
    typeArgs = case type_ of
      ADT className _ args ->
        let
          arity = case Map.lookup className metadata.pointerAdtPaths of
            Just info -> info.arity
            Nothing -> Array.length classInfo.vars
        in
          Array.take arity (map toGoType args)
      _ -> map (const TypeValue) classInfo.vars
    instantiations = Map.fromFoldable (Array.zip classInfo.vars typeArgs)
    properties = Map.fromFoldable (map (\(Prop key value) -> Tuple key value) props)
    -- Dictionary layout follows the class declaration, not record source order.
    fields = Array.mapMaybe (\field -> map (\value -> { field, value }) (Map.lookup field.name properties)) classInfo.fields
    translateField acc { field, value } =
      let
        genericType = exprTypeToGenericGoType metadata.pointerAdtPaths metadata.enumAdts metadata.elidedCtors classInfo.vars modNameStr field."type"
        expectedType = instantiateGenericGoType instantiations genericType
        fieldContext = (childContext context Nothing)
          { bound = bindFieldFunctionParameters toGoType bound field."type" value }
        result = translate fieldContext acc.nextId value
        -- Adapt each field before translating the next one: coercion can
        -- register reboxing helpers in the module's mutable codegen state.
        coerced = coerceGoExpr codegenStateRef modNameStr result.expr result.exprType expectedType
      in
        { stmts: acc.stmts <> result.stmts
        , exprs: Array.snoc acc.exprs coerced
        , nextId: result.nextId
        }
    translatedFields = foldl translateField { stmts: StmtEmpty, exprs: [], nextId } fields
  pure
    { stmts: translatedFields.stmts
    , expr: GoConstructor (hashString pointer.baseStructName) pointer.structName pointer.typeArgs translatedFields.exprs
    , exprType: TypeStructPointer pointer
    , nextId: translatedFields.nextId
    }
