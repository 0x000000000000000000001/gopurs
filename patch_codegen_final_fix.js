import fs from 'fs';
let content = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

content = content.replace(
  'res2 = translateExprImpl helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false res1.nextId e2',
  'res2 = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars namedBound bound Nothing [] false false res1.nextId e2'
);

// Also let's check for any others that might have a number or something!
content = content.replace(/translateExprImpl helpersRef/g, 'translateExprImpl_ helpersRef');
// Wait, if I just replace it globally now, it will lack the `false`!
// So I should just do a global replace of `translateExprImpl helpersRef` with `translateExprImpl_ helpersRef`, and then the compiler will tell me which ones are missing arguments!

fs.writeFileSync('src/Gopurs/CodeGen.purs', content);
