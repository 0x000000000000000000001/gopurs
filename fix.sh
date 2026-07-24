#!/bin/bash
sed -i '' -e 's/import Data.Semigroup.Foldable (foldMap1)/import Data.Semigroup.Foldable (foldMap1)\nimport Debug as Debug/' src/Gopurs/CodeGen.purs

# Add stripTyped and fix flattenApp
sed -i '' -e 's/flattenApp e@(TcoExpr _ (App f args)) =/flattenApp e@(TcoExpr analysis expr) = case expr of\n  App f args ->/' src/Gopurs/CodeGen.purs
sed -i '' -e 's/  let/    let/' src/Gopurs/CodeGen.purs
sed -i '' -e 's/    Tuple f'"'"' args'"'"' = flattenApp f/      Tuple f'"'"' args'"'"' = flattenApp f/' src/Gopurs/CodeGen.purs
sed -i '' -e 's/  in/    in/' src/Gopurs/CodeGen.purs
sed -i '' -e 's/    Tuple f'"'"' (args'"'"' <> toArray args)/      Tuple f'"'"' (args'"'"' <> toArray args)\n  Typed t inner ->\n    let\n      Tuple f'"'"' args'"'"' = flattenApp inner\n    in\n      Tuple (TcoExpr analysis (Typed t f'"'"')) args'"'"'\n  _ -> Tuple e []\n\nstripTyped :: BackendSyntax TcoExpr -> BackendSyntax TcoExpr\nstripTyped (Typed _ (TcoExpr _ e)) = stripTyped e\nstripTyped e = e/' src/Gopurs/CodeGen.purs
sed -i '' -e 's/flattenApp e = Tuple e \[\]//' src/Gopurs/CodeGen.purs

# Fix isTailCallTo
sed -i '' -e 's/if isTail then case unwrapExpr flatFn of/if isTail then case stripTyped (unwrapExpr flatFn) of/' src/Gopurs/CodeGen.purs

# Add Debug trace to isTailCallTo
sed -i '' -e 's/Array.findIndex (\\ctx -> ctx.ident == fullName) loopCtx/if name == "deepTailRec" then Debug.trace ("TCO check for deepTailRec. isTail = " <> show isTail <> " fn: " <> show (getVar (unwrapExpr flatFn))) \\_ -> Array.findIndex (\\ctx -> ctx.ident == fullName) loopCtx else Array.findIndex (\\ctx -> ctx.ident == fullName) loopCtx/' src/Gopurs/CodeGen.purs

# Fix getVar to support Let, Case, Fail
sed -i '' -e 's/getVar (CtorDef _ _ _ _) = Just { mbMod: Nothing, name: "CtorDef" }/getVar (CtorDef _ _ _ _) = Just { mbMod: Nothing, name: "CtorDef" }\n                getVar (Let _ _) = Just { mbMod: Nothing, name: "Let" }\n                getVar (Case _ _) = Just { mbMod: Nothing, name: "Case" }\n                getVar (Fail _) = Just { mbMod: Nothing, name: "Fail" }/' src/Gopurs/CodeGen.purs

# Fix mbDirectCall logic
sed -i '' -e 's/getVar (Typed _ inner) = getVar (unwrapTcoExpr inner)/getVar (Typed _ inner) = getVar (unwrapExpr inner)/' src/Gopurs/CodeGen.purs
sed -i '' -e 's/getFuncType (Typed _ inner) = getFuncType (unwrapTcoExpr inner)/getFuncType (Typed _ inner) = getFuncType (unwrapExpr inner)/' src/Gopurs/CodeGen.purs

# Fix mbDirectCall condition
sed -i '' -e 's/arity >= 2/arity >= 1/' src/Gopurs/CodeGen.purs
sed -i '' -e 's/fArgs: \[\], fRet: Any/fArgs, fRet/' src/Gopurs/CodeGen.purs
sed -i '' -e 's/mbDirectCall = case getVar (unwrapTcoExpr flatFn) of/mbDirectCall = case getFuncType (unwrapExpr flatFn) of\n                  Just { fArgs, fRet } ->\n                    case getVar (unwrapExpr flatFn) of/' src/Gopurs/CodeGen.purs
sed -i '' -e 's/result = if cond then Just { fullName: "Call_" <> sanitizeName name, fArgs, fRet, arity } else Nothing/result = if cond then Just { fullName: "Call_" <> sanitizeName name, fArgs, fRet, arity } else Nothing\n                        in result\n                      Nothing -> Nothing/' src/Gopurs/CodeGen.purs

# Fix unboxGoExpr in mbDirectCall
sed -i '' -e 's/accArgsRemaining = Array.drop arity accArgs.exprs/accArgsRemaining = Array.drop arity accArgs.exprs\n                      accArgsRemainingTypes = Array.drop arity accArgs.exprTypes/' src/Gopurs/CodeGen.purs
sed -i '' -e 's/GoCall (GoSelector (GoVar "gopurs_runtime") goFuncName) (Array.cons callExpr accArgsRemaining)/let expectedRetType = exprTypeToGoType fRet\n                          boxedRemaining = Array.zipWith (\\expr typ -> boxGoExpr expr typ) accArgsRemaining accArgsRemainingTypes\n                        in GoCall (GoSelector (GoVar "gopurs_runtime") goFuncName) (Array.cons (boxGoExpr callExpr expectedRetType) boxedRemaining)/' src/Gopurs/CodeGen.purs
sed -i '' -e 's/exprType: TypeValue/exprType: if Array.length accArgsRemaining == 0 then exprTypeToGoType fRet else TypeValue/' src/Gopurs/CodeGen.purs

