-- Fuse `let parts = Data.String.Common.split sep text in foldlArray (\acc p -> acc + cpFold p) seed parts`
-- (including when `seed` is `Array.length parts`) into a single allocation-free
-- scan over `text`. Preserves Go's `strings.Split` semantics for both empty
-- and non-empty separators, including UTF-8/WTF-8 code point folding per part.
module Gopurs.StringSplitFold (emit) where

import Prelude

import Control.Alternative (guard)
import Control.Monad.State (evalState)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (any)
import Data.Map as Map
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Newtype (unwrap)
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Tuple (Tuple(..))
import Gopurs.ArrayProducerFold (onlyCurried)
import Gopurs.ArrayTraverseFold as ArrayTraverseFold
import Gopurs.CallAnalysis (collectCurriedAbs, collectGoSpine, getGoSpineArgs, qualifiedTarget)
import Gopurs.CallArguments as CallArguments
import Gopurs.CodePointFold as CodePointFold
import Gopurs.CodegenState (CodePointFoldDescriptor)
import Gopurs.ExprAnalysis (unwrapTcoExpr)
import Gopurs.ExprContext (ExprContext, ExprResult, TranslateExpr, bindParameters, childContext, flattenStmts)
import Gopurs.GoAst (GoExpr(..), GoType(..), rawGo)
import Gopurs.GoConversions (coerceGoExpr)
import Gopurs.ImmediateApplications.Scope (freshen, maximumLevel)
import PureScript.Backend.Optimizer.Codegen.Tco (TcoExpr(..))
import PureScript.Backend.Optimizer.Codegen.Tco as Tco
import PureScript.Backend.Optimizer.CoreFn (ExprType(..), Ident, ModuleName(..))
import PureScript.Backend.Optimizer.FreeVars (localId)
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendOperator(..), BackendOperator1(..), BackendOperator2(..), BackendOperatorNum(..), BackendSyntax(..), Level)

data SeedMode
  = SeedLength
  | SeedExpr TcoExpr

emit :: TranslateExpr -> ExprContext -> Int -> Maybe Ident -> Level -> TcoExpr -> TcoExpr -> Maybe ExprResult
emit translate context nextId mbIdent level producer continuation = do
  guard (Array.null context.recVars && onlyCurried producer)
  Tuple sep text <- recognizeSplit producer
  body <- case unwrapTcoExpr continuation of
    EffectPure value -> do
      guard context.options.inEffectBlock
      guard (ArrayTraverseFold.annotated (ADT "Effect.Effect" [ "Effect", "Effect" ] [ Int ]) continuation)
      pure value
    _ -> pure continuation
  guard (onlyCurried body && ArrayTraverseFold.annotated Int body)
  let Tuple foldFn foldSpine = collectGoSpine body
  foldTarget <- qualifiedTarget foldFn
  guard (foldTarget.mbMod == Just (ModuleName "Data.Foldable") && foldTarget.name == "foldlArray")
  Tuple callback (Tuple seed partsArg) <- case getGoSpineArgs foldSpine of
    [ callback, seed, partsArg ] -> Just (Tuple callback (Tuple seed partsArg))
    _ -> Nothing
  guard (isLocalLevel level partsArg)
  cpDescriptor <- recognizeAdditiveCodePointCallback context level callback
  seedMode <- recognizeSeed level seed
  let whole = NeutralExpr (Let mbIdent level (neutral producer) (neutral continuation))
  pure (emitFused translate context nextId sep text seedMode cpDescriptor (maximumLevel whole))

recognizeSplit :: TcoExpr -> Maybe (Tuple TcoExpr TcoExpr)
recognizeSplit producer = do
  let Tuple fn spine = collectGoSpine producer
  target <- qualifiedTarget fn
  guard (target.mbMod == Just (ModuleName "Data.String.Common") && target.name == "split")
  guard (ArrayTraverseFold.annotated (Array String) producer)
  case getGoSpineArgs spine of
    [ sep, text ] -> do
      guard (ArrayTraverseFold.annotated String text)
      guard (ArrayTraverseFold.annotated String sep || ArrayTraverseFold.annotated (ADT "Data.String.Pattern.Pattern" [ "Data", "String", "Pattern", "Pattern" ] []) sep)
      pure (Tuple sep text)
    _ -> Nothing

recognizeAdditiveCodePointCallback :: ExprContext -> Level -> TcoExpr -> Maybe CodePointFoldDescriptor
recognizeAdditiveCodePointCallback context partsLevel callback = case unwrapTcoExpr callback of
  Abs params body -> do
    let grouped = collectCurriedAbs params body
    Tuple accLevel partLevel <- case NEA.toArray grouped.args of
      [ Tuple _ a, Tuple _ p ] -> Just (Tuple a p)
      _ -> Nothing
    guard (not (usesLevel partsLevel grouped.body))
    consumerCall <- case unwrapTcoExpr grouped.body of
      PrimOp (Op2 (OpIntNum OpAdd) lhs rhs)
        | isLocalLevel accLevel lhs -> Just rhs
        | isLocalLevel accLevel rhs -> Just lhs
      _ -> Nothing
    guard (onlyCurried consumerCall && ArrayTraverseFold.annotated Int consumerCall)
    let Tuple fn spine = collectGoSpine consumerCall
    arg <- case getGoSpineArgs spine of
      [ a ] -> Just a
      _ -> Nothing
    guard (isLocalLevel partLevel arg)
    target <- qualifiedTarget fn
    let isLocal = target.mbMod == Nothing || map (String.replaceAll (Pattern ".") (Replacement "_") <<< unwrap) target.mbMod == Just context.modNameStr
    info <- if isLocal then Map.lookup target.name context.moduleFunctions else do
      owner <- target.mbMod
      Map.lookup (unwrap owner <> "." <> target.name) context.metadata.globalFunctions
    guard (info.arity == 1 && info.fArgs == [ TypeString ] && info.fRet == TypeInt64)
    descriptor <- info.codePointFold
    guard (CodePointFold.validDescriptor descriptor)
    pure descriptor
  _ -> Nothing

recognizeSeed :: Level -> TcoExpr -> Maybe SeedMode
recognizeSeed partsLevel seed = case unwrapTcoExpr seed of
  PrimOp (Op1 OpArrayLength arr) | isLocalLevel partsLevel arr -> Just SeedLength
  _ ->
    if not (usesLevel partsLevel seed) && ArrayTraverseFold.annotated Int seed then
      Just (SeedExpr seed)
    else do
      guard (onlyCurried seed)
      let Tuple fn spine = collectGoSpine seed
      target <- qualifiedTarget fn
      guard (target.mbMod == Just (ModuleName "Data.Array") && target.name == "length")
      case getGoSpineArgs spine of
        [ arr ] | isLocalLevel partsLevel arr -> Just SeedLength
        _ -> Nothing

emitFused
  :: TranslateExpr
  -> ExprContext
  -> Int
  -> TcoExpr
  -> TcoExpr
  -> SeedMode
  -> CodePointFoldDescriptor
  -> Int
  -> ExprResult
emitFused translate context nextId sepExpr textExpr seedMode cpDesc lastLevel =
  let
    captureList = case seedMode of
      SeedLength -> [ sepExpr, textExpr ]
      SeedExpr s -> [ sepExpr, textExpr, s ]
    captured = CallArguments.capture { prefix: "stringSplitFold_arg_", markUsed: true }
      (\_ currentId value -> translate (childContext context Nothing) currentId value) nextId captureList
    arg index = fromMaybe (rawGo "nil") (Array.index captured.exprs index)
    argType index = fromMaybe TypeValue (Array.index captured.exprTypes index)
    sepGo = coerceGoExpr context.codegenStateRef context.modNameStr (arg 0) (argType 0) TypeString
    textGo = coerceGoExpr context.codegenStateRef context.modNameStr (arg 1) (argType 1) TypeString
    initAcc = case seedMode of
      SeedLength -> rawGo "int64(0)"
      SeedExpr _ -> coerceGoExpr context.codegenStateRef context.modNameStr (arg 2) (argType 2) TypeInt64
    start = max lastLevel (maximumLevel cpDesc.callback) + captured.nextId + 1
    freshenedCb = Tco.analyze [] (evalState (freshen cpDesc.callback) start)
    cbResult = case unwrapTcoExpr freshenedCb of
      Abs params cbBody ->
        let
          grouped = collectCurriedAbs params cbBody
          goParams = map (\(Tuple name lvl) -> Tuple (localId name lvl) TypeInt64) (NEA.toArray grouped.args)
          cbCtx = context
            { depth = context.depth + 1, bound = bindParameters goParams context.bound
            , tcoIdent = Nothing, loopCtx = [], options = { isTail: true, inEffectBlock: false }
            , mbExpectedExprType = Just Int, mbExpectedGoType = Just TypeInt64 }
          res = translate cbCtx captured.nextId grouped.body
        in { expr: GoFuncLit goParams (flattenStmts res.stmts) (coerceGoExpr context.codegenStateRef context.modNameStr res.expr res.exprType TypeInt64) TypeInt64
           , nextId: res.nextId }
      _ -> { expr: rawGo "nil", nextId: captured.nextId }
    suffix = show cbResult.nextId
    v label = "stringSplitFold_" <> label <> "_" <> suffix
    cbName = v "cb"
    foldPartName = v "part"
    sepName = v "sep"
    remName = v "rem"
    accName = v "acc"
    countName = v "count"
    idxName = v "idx"
    sepByteName = v "sepByte"
    -- Inline fast UTF-8 1/2/3/4-byte decoder per fragment with DecodeStringCodePoint fallback.
    foldPartFn = GoFuncLit [ Tuple "frag" TypeString ]
      [ GoAssign "res" (GoCall (GoVar "int64") [ GoInt cpDesc.seed ])
      , GoForRange "off := 0; off < len(frag);"
          [ GoAssign "b0" (rawGo "frag[off]")
          , rawGo "var r rune"
          , rawGo "var w int"
          , GoIfElse (rawGo "b0 < 0x80")
              [ GoMutate "r" (rawGo "rune(b0)"), GoMutate "w" (GoInt 1) ]
              [ GoIfElse (rawGo "b0 >= 0xC2 && b0 < 0xE0 && off+1 < len(frag) && frag[off+1]&0xC0 == 0x80")
                  [ GoMutate "r" (rawGo "rune(b0&0x1F)<<6 | rune(frag[off+1]&0x3F)"), GoMutate "w" (GoInt 2) ]
                  [ GoIfElse (rawGo "b0 >= 0xE0 && b0 < 0xF0 && b0 != 0xED && off+2 < len(frag) && frag[off+1]&0xC0 == 0x80 && frag[off+2]&0xC0 == 0x80 && (b0 != 0xE0 || frag[off+1] >= 0xA0)")
                      [ GoMutate "r" (rawGo "rune(b0&0x0F)<<12 | rune(frag[off+1]&0x3F)<<6 | rune(frag[off+2]&0x3F)"), GoMutate "w" (GoInt 3) ]
                      [ GoIfElse (rawGo "b0 >= 0xF0 && b0 <= 0xF4 && off+3 < len(frag) && frag[off+1]&0xC0 == 0x80 && frag[off+2]&0xC0 == 0x80 && frag[off+3]&0xC0 == 0x80 && (b0 != 0xF0 || frag[off+1] >= 0x90) && (b0 != 0xF4 || frag[off+1] <= 0x8F)")
                          [ GoMutate "r" (rawGo "rune(b0&0x07)<<18 | rune(frag[off+1]&0x3F)<<12 | rune(frag[off+2]&0x3F)<<6 | rune(frag[off+3]&0x3F)"), GoMutate "w" (GoInt 4) ]
                          [ GoMutate "r, w" (GoCall (GoSelector (GoVar "gopurs_runtime") "DecodeStringCodePoint") [ rawGo "frag[off:]" ]) ]
                      ]
                  ]
              ]
          , GoMutate "res" (GoCall (GoVar cbName) [ GoVar "res", GoCall (GoVar "int64") [ GoVar "r" ] ])
          , GoMutate "off" (rawGo "off + w")
          ]
      ] (GoVar "res") TypeInt64
    addPart fragExpr =
      [ GoMutate accName (GoBinOp "+" (GoVar accName) (GoCall (GoVar foldPartName) [ fragExpr ]))
      , GoMutate countName (GoBinOp "+" (GoVar countName) (rawGo "int64(1)"))
      ]
    stringsIndex = GoRaw { text: "strings.Index", imports: [ "strings" ] }
    stringsCount = GoRaw { text: "strings.Count", imports: [ "strings" ] }
    emptySepLoop =
      [ GoForRange ("len(" <> remName <> ") > 0")
          ([ GoRaw { text: "_, w := utf8.DecodeRuneInString(" <> remName <> ")", imports: [ "unicode/utf8" ] } ]
            <> addPart (rawGo (remName <> "[:w]"))
            <> [ GoMutate remName (rawGo (remName <> "[w:]")) ])
      ]
    isDirectAdd = case unwrapTcoExpr freshenedCb of
      Abs params cbBody ->
        let grouped = collectCurriedAbs params cbBody
        in case NEA.toArray grouped.args, unwrapTcoExpr grouped.body of
          [ Tuple _ a, Tuple _ b ], PrimOp (Op2 (OpIntNum OpAdd) l r) ->
            (isLocalLevel a l && isLocalLevel b r) || (isLocalLevel b l && isLocalLevel a r)
          _, _ -> false
      _ -> false
    -- When the per-part code-point fold is direct Int32 addition, splitting by
    -- any 1-byte ASCII separator (which cannot appear inside a multi-byte UTF-8
    -- or WTF-8 sequence) produces `1 + strings.Count(rem, sep)` fragments whose
    -- combined modular code-point sum equals
    -- `SumCodePoints(rem) - nSep * sepByte + count * seed` (mod 2^32).
    singleByteAsciiScan =
      if isDirectAdd then
        let
          nSepName = v "nSep"
          cpSumName = v "cpSum"
          seedTerm =
            if cpDesc.seed == 0 then ""
            else " + int64(int32(" <> countName <> "))*int64(" <> show cpDesc.seed <> ")"
        in
          [ GoAssign nSepName (GoCall (GoVar "int64") [ GoCall stringsCount [ GoVar remName, GoVar sepName ] ])
          , GoMutate countName (GoBinOp "+" (GoVar nSepName) (rawGo "int64(1)"))
          , rawGo (cpSumName <> ", _ := gopurs_runtime.SumCodePoints(" <> remName <> ")")
          , GoMutate accName (rawGo (accName <> " + " <> cpSumName <> " - int64(int32(" <> nSepName <> "))*int64(" <> sepName <> "[0])" <> seedTerm))
          ]
      else
        [ GoAssign sepByteName (rawGo (sepName <> "[0]"))
      , GoAssign "partRes" (GoCall (GoVar "int64") [ GoInt cpDesc.seed ])
      , GoMutate countName (rawGo "int64(1)")
      , GoForRange ("off := 0; off < len(" <> remName <> ");")
          [ GoAssign "b0" (rawGo (remName <> "[off]"))
          , GoIfElse (rawGo "b0 < 0x80")
              [ GoIfElse (rawGo ("b0 == " <> sepByteName))
                  [ GoMutate accName (GoBinOp "+" (GoVar accName) (GoVar "partRes"))
                  , GoMutate countName (GoBinOp "+" (GoVar countName) (rawGo "int64(1)"))
                  , GoMutate "partRes" (GoCall (GoVar "int64") [ GoInt cpDesc.seed ])
                  ]
                  [ GoMutate "partRes" (GoCall (GoVar cbName) [ GoVar "partRes", rawGo "int64(b0)" ]) ]
              , GoMutate "off" (rawGo "off + 1")
              , rawGo "continue"
              ]
              []
          , rawGo "var r rune"
          , rawGo "var w int"
          , GoIfElse (rawGo ("b0 >= 0xC2 && b0 < 0xE0 && off+1 < len(" <> remName <> ") && " <> remName <> "[off+1]&0xC0 == 0x80"))
              [ GoMutate "r" (rawGo ("rune(b0&0x1F)<<6 | rune(" <> remName <> "[off+1]&0x3F)")), GoMutate "w" (GoInt 2) ]
              [ GoIfElse (rawGo ("b0 >= 0xE0 && b0 < 0xF0 && b0 != 0xED && off+2 < len(" <> remName <> ") && " <> remName <> "[off+1]&0xC0 == 0x80 && " <> remName <> "[off+2]&0xC0 == 0x80 && (b0 != 0xE0 || " <> remName <> "[off+1] >= 0xA0)"))
                  [ GoMutate "r" (rawGo ("rune(b0&0x0F)<<12 | rune(" <> remName <> "[off+1]&0x3F)<<6 | rune(" <> remName <> "[off+2]&0x3F)")), GoMutate "w" (GoInt 3) ]
                  [ GoIfElse (rawGo ("b0 >= 0xF0 && b0 <= 0xF4 && off+3 < len(" <> remName <> ") && " <> remName <> "[off+1]&0xC0 == 0x80 && " <> remName <> "[off+2]&0xC0 == 0x80 && " <> remName <> "[off+3]&0xC0 == 0x80 && (b0 != 0xF0 || " <> remName <> "[off+1] >= 0x90) && (b0 != 0xF4 || " <> remName <> "[off+1] <= 0x8F)"))
                      [ GoMutate "r" (rawGo ("rune(b0&0x07)<<18 | rune(" <> remName <> "[off+1]&0x3F)<<12 | rune(" <> remName <> "[off+2]&0x3F)<<6 | rune(" <> remName <> "[off+3]&0x3F)")), GoMutate "w" (GoInt 4) ]
                      [ GoMutate "r, w" (GoCall (GoSelector (GoVar "gopurs_runtime") "DecodeStringCodePoint") [ rawGo (remName <> "[off:]") ]) ]
                  ]
              ]
          , GoMutate "partRes" (GoCall (GoVar cbName) [ GoVar "partRes", GoCall (GoVar "int64") [ GoVar "r" ] ])
          , GoMutate "off" (rawGo "off + w")
          ]
      , GoMutate accName (GoBinOp "+" (GoVar accName) (GoVar "partRes"))
      ]
    multiByteScan =
      [ GoForRange ""
          ([ GoAssign idxName (GoCall stringsIndex [ GoVar remName, GoVar sepName ])
           , GoIfElse (rawGo (idxName <> " < 0"))
               (addPart (GoVar remName) <> [ rawGo "break" ]) []
           ] <> addPart (rawGo (remName <> "[:" <> idxName <> "]"))
             <> [ GoMutate remName (rawGo (remName <> "[" <> idxName <> "+len(" <> sepName <> "):]")) ])
      ]
    finalizeStmts = case seedMode of
      SeedLength ->
        [ GoIfElse (rawGo (countName <> " > 0"))
            [ GoMutate accName (rawGo ("int64(int32(" <> accName <> " + " <> countName <> "))")) ] [] ]
      SeedExpr _ ->
        [ GoIfElse (rawGo (countName <> " > 0"))
            [ GoMutate accName (rawGo ("int64(int32(" <> accName <> "))")) ] [] ]
    body =
      [ GoAssign cbName cbResult.expr
      , GoAssign foldPartName foldPartFn
      , GoAssign sepName sepGo
      , GoAssign remName textGo
      , GoAssign accName initAcc
      , GoAssign countName (rawGo "int64(0)")
      , GoIfElse (rawGo ("len(" <> sepName <> ") == 1 && " <> sepName <> "[0] < 0x80"))
          singleByteAsciiScan
          [ GoIfElse (rawGo ("len(" <> sepName <> ") == 0")) emptySepLoop multiByteScan ]
      ] <> finalizeStmts
  in { stmts: captured.stmts
     , expr: GoCall (GoFuncLit [] body (GoVar accName) TypeInt64) []
     , exprType: TypeInt64
     , nextId: cbResult.nextId + 1 }

isLocalLevel :: Level -> TcoExpr -> Boolean
isLocalLevel expected expression = case unwrapTcoExpr expression of
  Local _ found -> found == expected
  _ -> false

usesLevel :: Level -> TcoExpr -> Boolean
usesLevel target (TcoExpr _ syntax) = case syntax of
  Local _ level -> level == target
  _ -> any (usesLevel target) syntax

neutral :: TcoExpr -> NeutralExpr
neutral (TcoExpr _ syntax) = NeutralExpr (map neutral syntax)
