import fs from 'fs';

let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
  /EffectDefer binding \-\>\n        let\n          resBinding = translateExprImpl helpersRef \(depth \+ 1\) modNameStr recVars namedBound bound Nothing \[\] false nextId binding\n          funcExpr = GoRaw \("gopurs_runtime.Func\\(func\\(_ gopurs_runtime.Value\\) gopurs_runtime.Value \{\\n" \<\> printGoExpr \(GoBlock \(flattenStmts resBinding.stmts \<\> \[ GoReturn resBinding.expr \]\)\) \<\> "\\n\}\\)"\)\n        in\n          \{ stmts: StmtEmpty, expr: funcExpr, nextId: resBinding.nextId \}/,
  `EffectDefer binding ->
        let
          resBinding = translateExprImpl helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false nextId binding
          bindingExpr = executeIfOpaque (unwrapTcoExpr binding) resBinding.expr
          funcExpr = GoRaw ("gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {\\n" <> printGoExpr (GoBlock (flattenStmts resBinding.stmts <> [ GoReturn bindingExpr ])) <> "\\n})")
        in
          { stmts: StmtEmpty, expr: funcExpr, nextId: resBinding.nextId }`
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
