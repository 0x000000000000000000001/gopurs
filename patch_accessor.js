import fs from 'fs';
let code = fs.readFileSync('bin/gopurs.js', 'utf8');
code = code.replace('return evalAccessor(v)(dictEval.eval(v)(v1._1))(v1._2);', `
      console.log("Accessor:", v1._2);
      return evalAccessor(v)(dictEval.eval(v)(v1._1))(v1._2);
`);
fs.writeFileSync('bin/gopurs.js', code);
