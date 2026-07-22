import fs from "fs";
let code = fs.readFileSync("../../output/Gopurs.CodeGen/index.js", "utf8");
code = code.replace(/var translateExprImpl = function /g, 'var translateExprImplCount = 0; var translateExprImpl = function ');
code = code.replace(/var resE = translateExprImpl/g, 'if (++translateExprImplCount % 1000 === 0) console.log("translateExprImpl calls:", translateExprImplCount); var resE = translateExprImpl');
// Wait, index.js has multiple calls to translateExprImpl. I will just replace the function body beginning:
code = code.replace(/var translateExprImpl = function \([^)]+\) {\n    return function \([^)]+\) {/g, 'var translateExprImplCount = 0;\nvar translateExprImpl = function(mod) { return function(r1) { return function(r2) { return function(r3) { return function(r4) { return function(r5) { return function(r6) { return function(r7) { return function(r8) { if (++translateExprImplCount % 10000 === 0) console.log("translateExprImpl", translateExprImplCount);');
// Actually, too complex to inject manually.

