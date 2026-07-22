const fs = require('fs');

let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

const replacementImports = `
    declsAst = Array.mapMaybe (\\b -> translateBinding modNameStr b) (Array.fromFoldable mod.bindings)
    allDeclsAst = declsAst <> unsafePerformEffect (Ref.read helpersRef)
    declsStr = String.joinWith "\\n" (map printGoDecl allDeclsAst)
    
    parts = String.split (Pattern "pkg_") declsStr
    usedPkgNames = Array.nub $ Array.mapMaybe (\\part ->
      let subParts = String.split (Pattern ".") part
      in Array.head subParts
    ) (fromMaybe [] (Array.tail parts))

    goImports = Array.nub $ 
      (if Array.length allDeclsAst > 0 || Array.length (Array.fromFoldable mod.foreign) > 0 then [ "gopurs/output/gopurs_runtime" ] else []) <>
      (if Array.length allDeclsAst > 0 then [ "sync" ] else []) <>
      Array.mapMaybe (\\pkg -> 
        if pkg /= modNameStr && pkg /= "Prim"
           then Just ("gopurs/output/" <> pkg)
           else Nothing
      ) usedPkgNames

    goFile = { packageName: modNameStr
             , imports: goImports
             , decls: allDeclsAst
`;

code = code.replace(/    declsAst = Array\.mapMaybe[\s\S]*?, decls: decls <> unsafePerformEffect \(Ref\.read helpersRef\)/m, replacementImports.trim());

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
console.log('Patched CodeGen.purs with imports logic');
