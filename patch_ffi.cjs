const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/FfiSupport.js', 'utf8');

const startParse = code.indexOf('let parseFuncType = function(t) {');
const endWrap = code.indexOf('let goFuncArgsBoxed = parsedArgs', startParse);
const toMove = code.substring(startParse, endWrap);

code = code.substring(0, startParse) + code.substring(endWrap);

const insertPoint = code.indexOf('// 1. Generate Native Call_X proxy');
code = code.substring(0, insertPoint) + toMove + '\n                ' + code.substring(insertPoint);

const oldCallGen = `let callRetNative = callRet === 'interface{}' || callRet === 'any' ? 'gopurs_runtime.Value' : callRet;
                let goFuncArgsNativeProxy = parsedArgs.map((_, idx) => {
                    let t = parsedArgs[idx].type;
                    if (t === 'interface{}' || t === 'any') return \`arg\${idx} gopurs_runtime.Value\`;
                    return \`arg\${idx} \${t}\`;
                }).join(', ');
                
                newLines.push(\`func Call_\${pursName}\${typeParams}(\${goFuncArgsNativeProxy}) \${callRetNative} {\`);
                let nativeCallArgs = parsedArgs.map((_, idx) => \`arg\${idx}\`).join(', ');
                if (callRet === '') {
                    newLines.push(\`\\t\${nativeCallFunc}(\${nativeCallArgs})\`);
                } else if (callRet === 'interface{}' || callRet === 'any') {
                    newLines.push(\`\\treturn gopurs_runtime.Box(\${nativeCallFunc}(\${nativeCallArgs}))\`);
                } else {
                    newLines.push(\`\\treturn \${nativeCallFunc}(\${nativeCallArgs})\`);
                }
                newLines.push(\`}\`);`;

const newCallGen = `let callRetNative = callRet === 'interface{}' || callRet === 'any' || callRet.startsWith('func') ? 'gopurs_runtime.Value' : callRet;
                let goFuncArgsNativeProxy = parsedArgs.map((_, idx) => {
                    let t = parsedArgs[idx].type;
                    if (t === 'interface{}' || t === 'any' || t.startsWith('func')) return \`arg\${idx} gopurs_runtime.Value\`;
                    return \`arg\${idx} \${t}\`;
                }).join(', ');
                
                newLines.push(\`func Call_\${pursName}\${typeParams}(\${goFuncArgsNativeProxy}) \${callRetNative} {\`);
                let nativeCallArgs = parsedArgs.map((_, idx) => {
                    let t = parsedArgs[idx].type;
                    if (t.startsWith('func')) return unwrapValueToFunc(t, \`arg\${idx}\`, 0).replace(/\\n/g, "\\n\\t");
                    return \`arg\${idx}\`;
                }).join(', ');
                if (callRet === '') {
                    newLines.push(\`\\t\${nativeCallFunc}(\${nativeCallArgs})\`);
                } else if (callRetNative === 'gopurs_runtime.Value') {
                    newLines.push(\`\\treturn \${wrapReturn(callRet, nativeCallFunc + "(" + nativeCallArgs + ")").replace(/\\n/g, "\\n\\t")}\`);
                } else {
                    newLines.push(\`\\treturn \${nativeCallFunc}(\${nativeCallArgs})\`);
                }
                newLines.push(\`}\`);`;

code = code.replace(oldCallGen, newCallGen);
fs.writeFileSync('src/Gopurs/FfiSupport.js', code);
