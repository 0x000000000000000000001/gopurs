const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

const regex = /          deadVarOpt = Array\.head \([\s\S]*?         \)\n\n          result =/m;
const deadVarOptMatch = code.match(regex);

if (deadVarOptMatch) {
  let deadVarOptCode = deadVarOptMatch[0].replace('          result =\n', '');
  code = code.replace(regex, '          result =');
  
  const allocBlockRegex = /              allocBlock = if isElided then[\s\S]*?                 else GoConstructor \(hashString baseStructName\) \(pkgPrefix <> monoStructName\) typeArgs accProps\.exprs\n                 \n/;
  const allocBlockMatch = code.match(allocBlockRegex);
  
  if (allocBlockMatch) {
    let typeArgsStrCode = '              typeArgsStr = if Array.length typeArgs > 0 then "[" <> String.joinWith ", " (map goTypeToStr typeArgs) <> "]" else ""\n';
    
    // insert deadVarOpt after typeArgsStrCode
    code = code.replace(allocBlockRegex, allocBlockMatch[0] + typeArgsStrCode + '    ' + deadVarOptCode.trim() + '\n\n');
  }
}
fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
