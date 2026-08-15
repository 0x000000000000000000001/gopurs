const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
`          resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities newBound Nothing loopCtx isTail false resBinding.nextId body`,
`          resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities newBound Nothing loopCtx isTail inEffectBlock resBinding.nextId body`
);

code = code.replace(
`                resBodyOuter = translateExprImpl_ helpersRef (depth + 1) modNameStr combinedRecVars moduleArities allocRes.newBound Nothing loopCtx isTail false nextId' body`,
`                resBodyOuter = translateExprImpl_ helpersRef (depth + 1) modNameStr combinedRecVars moduleArities allocRes.newBound Nothing loopCtx isTail inEffectBlock nextId' body`
);

code = code.replace(
`                resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr combinedRecVars moduleArities allocRes.newBound Nothing loopCtx isTail false accBindings.nextId body`,
`                resBody = translateExprImpl_ helpersRef (depth + 1) modNameStr combinedRecVars moduleArities allocRes.newBound Nothing loopCtx isTail inEffectBlock accBindings.nextId body`
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
