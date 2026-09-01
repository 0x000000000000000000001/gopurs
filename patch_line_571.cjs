const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf-8');

const target1 = `loopCtxs = map (\\fn -> { ident: fn.ident, params: fn.args, loopParams: map (\\p -> p <> "_loop") fn.args }) fns`;
const replacement1 = `loopCtxs = map (\\fn -> { ident: fn.ident, params: fn.args, loopParams: map (\\p -> p <> "_loop") fn.args, goTypes: [], fRet: TypeValue }) fns`;

if (!code.includes(target1)) {
  console.error("Target string not found!");
  process.exit(1);
}

fs.writeFileSync('src/Gopurs/CodeGen.purs', code.split(target1).join(replacement1));
console.log("Patched CodeGen.purs successfully.");
