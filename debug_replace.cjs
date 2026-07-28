const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace("module Gopurs.CodeGen where\n", "module Gopurs.CodeGen where\nimport Debug as Debug\n");

code = code.replace(
  /deadVarOpt = Array\.head \(Array\.mapMaybe \\\(\\\(Tuple _ v\\\) ->[^{]*if v\.goType == expectedGoType[^{]*&& not \\\(Set\.member v\.name liveOut\\\)[^{]*&& not \\\(Set\.member v\.name \\\(freeVars tcoExpr\\\)\\\)[^{]*then Just v\.name[^{]*else Nothing/m,
  `deadVarOpt = Array.head (Array.mapMaybe (\\(Tuple _ v) -> 
                  let
                    _ = Debug.trace { msg: "FBIP Check", vname: v.name, vtype: goTypeToStr v.goType, exType: goTypeToStr expectedGoType, isLive: Set.member v.name liveOut, isFree: Set.member v.name (freeVars tcoExpr) } (\\_ -> unit)
                  in
                  if v.goType == expectedGoType 
                     && not (Set.member v.name liveOut) 
                     && not (Set.member v.name (freeVars tcoExpr)) 
                  then Just v.name 
                  else Nothing`
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
