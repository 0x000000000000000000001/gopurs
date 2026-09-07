const fs = require('fs');
let code = fs.readFileSync('bin/gopurs.js', 'utf8');

const hook = `
let _ = (function() {
  if (v3.size > 100000) {
    console.log("Analyzing AST of size", v3.size);
    let counts = {};
    function traverse(node) {
      if (!node) return;
      let name = node.constructor ? node.constructor.name : typeof node;
      if (name === "ExprSyntax") {
         let syn = node.value1;
         let sname = syn.constructor.name;
         counts[sname] = (counts[sname] || 0) + 1;
         if (sname === "App") { traverse(syn.value0); traverse(syn.value1); }
         else if (sname === "Var") { }
         else if (sname === "Local") { }
         else if (sname === "Lit") { }
         else if (sname === "Abs") { traverse(syn.value1); }
         else if (sname === "Let") { traverse(syn.value2); traverse(syn.value3); }
         else if (sname === "Branch") {
           syn.value0.forEach(pair => { traverse(pair.value0); traverse(pair.value1); });
           traverse(syn.value1);
         }
         else if (sname === "PrimOp") {
           let op = syn.value0;
           if (op.constructor.name === "Op1") { traverse(op.value1); }
           else if (op.constructor.name === "Op2") { traverse(op.value1); traverse(op.value2); }
         }
         else if (sname === "CtorDef") { }
         else if (sname === "CtorSaturated") { syn.value4.forEach(f => traverse(f)); }
         else if (sname === "Accessor") { traverse(syn.value0); }
         else if (sname === "Update") { traverse(syn.value0); syn.value1.forEach(f => traverse(f.value1)); }
         else if (sname === "Typed") { traverse(syn.value1); }
         else if (sname === "LetRec") { syn.value1.forEach(b => traverse(b.value1)); traverse(syn.value2); }
      }
    }
    traverse(expr22);
    console.log("Counts:", counts);
  }
})();
`;

code = code.replace(/var v3 = unwrap6\(analysisOf2\(expr22\)\);/, `var v3 = unwrap6(analysisOf2(expr22)); ${hook}`);
fs.writeFileSync('bin/gopurs.js', code);
