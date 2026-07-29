import fs from 'fs';

let code = fs.readFileSync('src/Gopurs/FfiSupport.js', 'utf8');

let helpersStart = code.indexOf('let parseFuncType = function(t) {');
let helpersEnd = code.indexOf('let goFuncArgsBoxed = parsedArgs', helpersStart);
let helpersCode = code.substring(helpersStart, helpersEnd);
code = code.substring(0, helpersStart) + code.substring(helpersEnd);

let insertPoint = code.indexOf('// 1. Generate Native Call_X proxy');
code = code.substring(0, insertPoint) + helpersCode + '\n                ' + code.substring(insertPoint);

const oldCallGen = `                let callRet = retStr.trim() !== '' ? retStr : '';
                newLines.push(\`func Call_\${pursName}\${typeParams}(\${goFuncArgsNative}) \${callRet} {\`);
                let nativeCallArgs = parsedArgs.map((_, idx) => \`arg\${idx}\`).join(', ');
                if (callRet === '') {
                    newLines.push(\`\\t\${nativeCallFunc}(\${nativeCallArgs})\`);
                } else {
                    newLines.push(\`\\treturn \${nativeCallFunc}(\${nativeCallArgs})\`);
                }`;

const newCallGen = `                let callRet = retStr.trim() !== '' ? retStr : '';
                let callRetNative = callRet === 'interface{}' || callRet === 'any' || callRet.startsWith('func') ? 'gopurs_runtime.Value' : callRet;
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
                }`;

code = code.replace(oldCallGen, newCallGen);
fs.writeFileSync('src/Gopurs/FfiSupport.js', code);
