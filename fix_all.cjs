const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

// Replace any missing Set.empty calls by searching for common patterns of boolean arguments
code = code.replace(
  'Nothing [] false false resRef.nextId val',
  'Nothing [] Set.empty false false resRef.nextId val'
);

code = code.replace(
  'Nothing [] false false (nextId + 1) e1',
  'Nothing [] Set.empty false false (nextId + 1) e1'
);

code = code.replace(
  'Nothing [] false false res1.nextId e2',
  'Nothing [] Set.empty false false res1.nextId e2'
);

// Any other `[] false false` that might be translateExprImpl_
code = code.replace(
  'bound Nothing [] false false',
  'bound Nothing [] Set.empty false false'
);

code = code.replace(
  'loopCtx false false',
  'loopCtx Set.empty false false'
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
