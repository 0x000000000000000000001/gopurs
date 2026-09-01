const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf-8');

const target1 = `targetCtx = fromMaybe { ident: "", params: [], loopParams: [], goTypes: [] } (Array.index loopCtx index)`;
const replacement1 = `targetCtx = fromMaybe { ident: "", params: [], loopParams: [], goTypes: [], fRet: TypeValue } (Array.index loopCtx index)`;

if (!code.includes(target1)) {
  console.error("Target string not found!");
  process.exit(1);
}

fs.writeFileSync('src/Gopurs/CodeGen.purs', code.split(target1).join(replacement1));
console.log("Patched CodeGen.purs successfully.");
