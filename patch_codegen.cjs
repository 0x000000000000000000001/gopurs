const fs = require('fs');

let ast = fs.readFileSync('src/Gopurs/GoAst.purs', 'utf8');
ast = ast.replace('| GoMap (Array (Tuple String GoExpr))', '| GoRecordDict (Array (Tuple String GoExpr))\n  | GoRecordUpdateDict GoExpr (Array (Tuple String GoExpr))');
fs.writeFileSync('src/Gopurs/GoAst.purs', ast);

let printer = fs.readFileSync('src/Gopurs/Printer.purs', 'utf8');
printer = printer.replace('  GoMap props ->\n    "map[string]gopurs_runtime.Value{" <> String.joinWith ", " (map (\\(Tuple k v) -> "\\"" <> k <> "\\": " <> printGoExpr v) props) <> "}"', `  GoRecordDict props ->
    let
      keysStr = String.joinWith ", " (map (\\(Tuple k _) -> "\\"" <> k <> "\\"") props)
      valsStr = String.joinWith ", " (map (\\(Tuple _ v) -> printGoExpr v) props)
    in
      "gopurs_runtime.RecordDict([]string{" <> keysStr <> "}, []gopurs_runtime.Value{" <> valsStr <> "})"
  GoRecordUpdateDict orig props ->
    let
      keysStr = String.joinWith ", " (map (\\(Tuple k _) -> "\\"" <> k <> "\\"") props)
      valsStr = String.joinWith ", " (map (\\(Tuple _ v) -> printGoExpr v) props)
    in
      "gopurs_runtime.RecordUpdateDict(" <> printGoExpr orig <> ", []string{" <> keysStr <> "}, []gopurs_runtime.Value{" <> valsStr <> "})"`);
printer = printer.replace('  GoRecordAccess obj prop ->\n    printGoExpr obj <> ".PtrVal.(map[string]gopurs_runtime.Value)[\\"" <> prop <> "\\"]"', `  GoRecordAccess obj prop ->
    "gopurs_runtime.RecordGet(" <> printGoExpr obj <> ", \\"" <> prop <> "\\")"`);
fs.writeFileSync('src/Gopurs/Printer.purs', printer);

let codegen = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');
codegen = codegen.replace('expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Record") [ GoMap accProps.exprs ]', 'expr: GoRecordDict accProps.exprs');
codegen = codegen.replace('expr: GoCall (GoSelector (GoVar "gopurs_runtime") "RecordUpdate") [ resObj.expr, GoMap accProps.exprs ]', 'expr: GoRecordUpdateDict resObj.expr accProps.exprs');
codegen = codegen.replace('GoCall (GoSelector (GoVar "gopurs_runtime") "Record") [ GoMap (Array.cons (Tuple "_tag" (GoCall (GoSelector (GoVar "gopurs_runtime") "Str") [ GoString name ])) (map (\\f -> Tuple f (GoVar (sanitizeName f))) fields)) ]', 'GoRecordDict (Array.cons (Tuple "_tag" (GoCall (GoSelector (GoVar "gopurs_runtime") "Str") [ GoString name ])) (map (\\f -> Tuple f (GoVar (sanitizeName f))) fields))');
codegen = codegen.replace('expr: GoCall (GoSelector (GoVar "gopurs_runtime") "Record") [ GoMap (Array.cons (Tuple "_tag" (GoCall (GoSelector (GoVar "gopurs_runtime") "Str") [ GoString name ])) accProps.exprs) ]', 'expr: GoRecordDict (Array.cons (Tuple "_tag" (GoCall (GoSelector (GoVar "gopurs_runtime") "Str") [ GoString name ])) accProps.exprs)');
fs.writeFileSync('src/Gopurs/CodeGen.purs', codegen);

