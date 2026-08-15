const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(/(import Node\.Path as Path\n)/, "$1import Effect.Console as Effect.Console\n");

code = code.replace(/expectedGoType = exprTypeToGoType([^]*?)newBound = Map\.insert originalName \{ name, goType: expectedGoType \} bound/, 
`expectedGoType = exprTypeToGoType$1          synName = (\\(TcoExpr _ syn) -> case syn of { Typed _ _ -> "Typed"; PrimOp _ -> "PrimOp"; App _ _ -> "App"; Var _ -> "Var"; Let _ _ _ _ -> "Let"; LetRec _ _ _ -> "LetRec"; Abs _ _ -> "Abs"; Lit _ -> "Lit"; CtorDef _ _ _ _ -> "CtorDef"; _ -> "Other" }) binding
          _ = unsafePerformEffect (if String.contains (Pattern "__local_var") name then Effect.Console.log ("MakeLet " <> name <> " (" <> modNameStr <> ") -> " <> synName <> ", expectedGoType: " <> goTypeToStr expectedGoType) else pure unit)\n          newBound = Map.insert originalName { name, goType: expectedGoType } bound`);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
