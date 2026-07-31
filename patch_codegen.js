const fs = require('fs');
let code = fs.readFileSync('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
  /getAdtModules \(ADT parts args\) = \s*let modPart = String\.joinWith "\." \(fromMaybe \[\] \(Array\.init parts\)\)\s*in \[ modPart \] <> Array\.concatMap getAdtModules args/,
  `getAdtModules (ADT parts args) = 
          let modPart = String.joinWith "." (fromMaybe [] (Array.init parts))
              isPtr = Map.member (String.joinWith "." parts) pointerAdtPaths
              modList = if isPtr then [ modPart ] else []
          in modList <> Array.concatMap getAdtModules args`
);

fs.writeFileSync('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', code);
