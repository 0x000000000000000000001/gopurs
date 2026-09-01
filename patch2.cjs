const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');
code = code.replace(/mbClass = Map\.lookup fullName classDeclsFields/g, 'mbClass = Debug.trace ("Looking up " <> fullName <> ", found: " <> show (isJust (Map.lookup fullName classDeclsFields))) \\_ -> Map.lookup fullName classDeclsFields');
fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
