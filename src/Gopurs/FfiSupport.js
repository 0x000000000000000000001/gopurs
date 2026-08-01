
export const appendFfiWrappersImpl = function(moduleName) {
    return function(content) {
        let ffiMarkerIdx = content.indexOf("// --- Auto-generated FFI wrappers ---");
        if (ffiMarkerIdx !== -1) {
            content = content.substring(0, ffiMarkerIdx);
        }
        let text = content.replace(/^package\s+[a-zA-Z0-9_]+[\s\n]*/m, '');
        text = text.replace(/^[ \t]*import[ \t]+"gopurs\/output\/gopurs_runtime"[\s\n]*/gm, '');
        text = text.replace(/^[ \t]*"gopurs\/output\/gopurs_runtime"[\s\n]*/gm, '');
        text = text.replace(/import\s*\(\s*\)[\s\n]*/gm, '');

        const goKeywords = ["break", "default", "func", "interface", "select", "case", "defer", "go", "map", "struct", "chan", "else", "goto", "package", "switch", "const", "fallthrough", "if", "range", "type", "continue", "for", "import", "return", "var"];
        function getExportName(n) {
            let base = n.endsWith('_') ? n.slice(0, -1).toLowerCase() : n.toLowerCase();
            if (goKeywords.includes(base)) return "_Gopurs_Go__" + base;
            return "_Gopurs_" + n;
        }

        const lines = text.split('\n');
        let newLines = [];

        for (let i = 0; i < lines.length; i++) {
            const line = lines[i];
            const basicMatch = line.match(/^(?:func\s+(_?[A-Z][A-Za-z0-9_]*)(?:\s*\[([^\]]+)\])?\s*\(|var\s+(_?[A-Z][A-Za-z0-9_]*)\s*=\s*func\s*\()/);
            if (basicMatch) {
                const funcName = basicMatch[1] || basicMatch[3];
                const typeParams = basicMatch[2] ? `[${basicMatch[2]}]` : ``;
                let exportName = getExportName(funcName);
                
                let startIdx = basicMatch[0].length;
                let parenDepth = 1;
                let i = startIdx;
                for (; i < line.length; i++) {
                    if (line[i] === '(') parenDepth++;
                    else if (line[i] === ')') {
                        parenDepth--;
                        if (parenDepth === 0) break;
                    }
                }
                
                const argsStr = line.substring(startIdx, i);
                let remainder = line.substring(i + 1).trim();
                let braceIdx = remainder.indexOf(' {');
                if (braceIdx !== -1) {
                    remainder = remainder.substring(0, braceIdx).trim();
                } else if (remainder.endsWith('{')) {
                    remainder = remainder.substring(0, remainder.length - 1).trim();
                }
                const retStr = remainder;

                // Parse arguments
                let args = [];
                if (argsStr.trim() !== '') {
                    let currentArg = '';
                    let parenDepth = 0;
                    for (let j = 0; j < argsStr.length; j++) {
                        let char = argsStr[j];
                        if (char === '(') parenDepth++;
                        else if (char === ')') parenDepth--;
                        
                        if (char === ',' && parenDepth === 0) {
                            args.push(currentArg.trim());
                            currentArg = '';
                        } else {
                            currentArg += char;
                        }
                    }
                    if (currentArg.trim() !== '') {
                        args.push(currentArg.trim());
                    }
                }
                
                let parsedArgs = [];
                for (let argStr of args) {
                    let spaceIdx = argStr.indexOf(' ');
                    if (spaceIdx === -1) {
                        parsedArgs.push({ name: `_arg${parsedArgs.length}`, type: argStr });
                    } else {
                        parsedArgs.push({ name: argStr.substring(0, spaceIdx).trim(), type: argStr.substring(spaceIdx + 1).trim() });
                    }
                }
                
                let arity = parsedArgs.length;
                if (arity > 11) continue; // Not supported yet

                let funcConstructor = arity === 0 ? "Func" : (arity === 1 ? "Func" : `Func${arity}`);
                let goFuncArgsNative = parsedArgs.map((_, idx) => `arg${idx} ${parsedArgs[idx].type}`).join(', ');
                
                let pursName = funcName;
                if (pursName.charAt(0) >= 'A' && pursName.charAt(0) <= 'Z') {
                    pursName = pursName.charAt(0).toLowerCase() + pursName.substring(1);
                }
                
                // 1. Generate Native Call_X proxy
                let nativeCallFunc = funcName;
                let typeParamNames = [];
                if (basicMatch[2]) {
                    basicMatch[2].split(',').forEach(tp => {
                        let parts = tp.trim().split(/\s+/);
                        if (parts.length > 0) typeParamNames.push(parts[0]);
                    });
                    nativeCallFunc = `${funcName}[${typeParamNames.join(', ')}]`;
                }

                let callRet = retStr.trim() !== '' ? retStr : '';
                newLines.push(`func Call_${pursName}${typeParams}(${goFuncArgsNative}) ${callRet} {`);
                let nativeCallArgs = parsedArgs.map((_, idx) => `arg${idx}`).join(', ');
                if (callRet === '') {
                    newLines.push(`\t${nativeCallFunc}(${nativeCallArgs})`);
                } else {
                    newLines.push(`\treturn ${nativeCallFunc}(${nativeCallArgs})`);
                }
                newLines.push(`}`);
                
                // 2. Generate boxed wrapper for dynamic dispatch
                let goFuncArgsBoxed = parsedArgs.map((_, idx) => `arg${idx} gopurs_runtime.Value`).join(', ');
                if (arity === 0) goFuncArgsBoxed = `_ gopurs_runtime.Value`;
                if (arity === 1) goFuncArgsBoxed = `arg0 gopurs_runtime.Value`;
                
                newLines.push(`var ${exportName} = gopurs_runtime.${funcConstructor}(func(${goFuncArgsBoxed}) gopurs_runtime.Value {`);
                let callFunc = funcName;
                let substituteGeneric = function(t) { return t; };
                if (basicMatch[2]) {
                    callFunc = `${funcName}[${typeParamNames.map(() => 'gopurs_runtime.Value').join(', ')}]`;
                    substituteGeneric = function(t) {
                        let res = t;
                        typeParamNames.forEach(tp => {
                            let re = new RegExp("\\b" + tp + "\\b", 'g');
                            res = res.replace(re, 'gopurs_runtime.Value');
                        });
                        return res;
                    };
                }
                let callArgs = [];
                let parseFuncType = function(t) {
                    let match = t.match(/^func\s*\((.*)/);
                    if (!match) return null;
                    let rest = match[1];
                    let parens = 1;
                    let i = 0;
                    while (i < rest.length && parens > 0) {
                        if (rest[i] === '(') parens++;
                        else if (rest[i] === ')') parens--;
                        i++;
                    }
                    let argsStr = rest.substring(0, i - 1);
                    let retStr = rest.substring(i).trim();
                    
                    let args = [];
                    let currentArg = "";
                    parens = 0;
                    let brackets = 0;
                    let braces = 0;
                    for (let j = 0; j < argsStr.length; j++) {
                        let c = argsStr[j];
                        if (c === '(') parens++;
                        else if (c === ')') parens--;
                        else if (c === '[') brackets++;
                        else if (c === ']') brackets--;
                        else if (c === '{') braces++;
                        else if (c === '}') braces--;
                        else if (c === ',' && parens === 0 && brackets === 0 && braces === 0) {
                            if (currentArg.trim() !== "") args.push(currentArg.trim());
                            currentArg = "";
                            continue;
                        }
                        currentArg += c;
                    }
                    if (currentArg.trim() !== "") args.push(currentArg.trim());
                    
                    return { args: args, retStr: retStr };
                };

                let unwrapValueToFunc = function(t, valName, depth) {
                    let parsed = parseFuncType(t);
                    if (!parsed) return `gopurs_runtime.Unbox[${t}](${valName})`;
                    
                    let args = parsed.args;
                    let retStr = parsed.retStr;
                    
                    let params = args.map((_, cidx) => `p${depth}_${cidx} ${args[cidx]}`).join(', ');
                    let applyArgs = args.map((atype, cidx) => {
                        if (atype === "gopurs_runtime.Value") return `p${depth}_${cidx}`;
                        if (atype.startsWith("func")) return wrapReturn(atype, `p${depth}_${cidx}`).replace(/\n/g, "\n\t\t");
                        return `gopurs_runtime.Box(p${depth}_${cidx})`;
                    }).join(', ');
                    
                    let applyCall = '';
                    if (args.length === 1) applyCall = `gopurs_runtime.Apply(${valName}, ${applyArgs})`;
                    else if (args.length > 1) applyCall = `gopurs_runtime.Apply${args.length}(${valName}, ${applyArgs})`;
                    else applyCall = `gopurs_runtime.Apply(${valName}, gopurs_runtime.Value{})`;

                    if (retStr === "") {
                        return `func(${params}) {\n\t\t${applyCall}\n\t}`;
                    } else if (retStr.startsWith("[]") && retStr !== "[]gopurs_runtime.Value") {
                        let elemType = retStr.substring(2);
                        return `func(${params}) ${retStr} {\n\t\tinner_res${depth} := ${applyCall}\n\t\tres_arr${depth} := *(*[]gopurs_runtime.Value)(inner_res${depth}.UnsafePtr)\n\t\tres_go${depth} := make(${retStr}, len(res_arr${depth}))\n\t\tfor i, v := range res_arr${depth} { res_go${depth}[i] = gopurs_runtime.Unbox[${elemType}](v) }\n\t\treturn res_go${depth}\n\t}`;
                    } else if (retStr === "any" || retStr === "interface{}") {
                        return `func(${params}) ${retStr} {\n\t\treturn ${applyCall}\n\t}`;
                    } else if (retStr === "gopurs_runtime.Value") {
                        return `func(${params}) ${retStr} {\n\t\treturn ${applyCall}\n\t}`;
                    } else if (retStr.startsWith("func")) {
                        let innerUnwrap = unwrapValueToFunc(retStr, `inner_res${depth}`, depth+1);
                        return `func(${params}) ${retStr} {\n\t\tinner_res${depth} := ${applyCall}\n\t\treturn ${innerUnwrap}\n\t}`;
                    } else {
                        return `func(${params}) ${retStr} {\n\t\tinner_res${depth} := ${applyCall}\n\t\treturn gopurs_runtime.Unbox[${retStr}](inner_res${depth})\n\t}`;
                    }
                };

                let wrapReturn = function(t, valName) {
                    if (t.startsWith("func")) {
                        let parsed = parseFuncType(t);
                        if (parsed) {
                            let innerT = parsed.retStr;
                            let argT = parsed.args.length > 0 ? parsed.args[0] : null;
                            
                            if (argT === null) {
                                if (innerT === "") {
                                    return `gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {\n` +
                                           `\t\t\t${valName}()\n` +
                                           `\t\t\treturn gopurs_runtime.Value{}\n` +
                                           `\t\t})`;
                                } else {
                                    let innerWrap = wrapReturn(innerT, "inner_res");
                                    return `gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {\n` +
                                           `\t\t\tinner_res := ${valName}()\n` +
                                           `\t\t\treturn ${innerWrap}\n` +
                                           `\t\t})`;
                                }
                            } else {
                                let argUnwrap = "arg.PtrVal()";
                                if (argT === "any" || argT === "interface{}") {
                                    argUnwrap = "arg";
                                } else if (argT === "gopurs_runtime.Value") {
                                    argUnwrap = "arg";
                                } else if (argT.startsWith("func")) {
                                    argUnwrap = unwrapValueToFunc(argT, "arg", 99).replace(/\n/g, "\n\t\t\t");
                                } else {
                                    argUnwrap = `gopurs_runtime.Unbox[${argT}](arg)`;
                                }
                                
                                if (innerT === "") {
                                    return `gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {\n` +
                                           `\t\t\t${valName}(${argUnwrap})\n` +
                                           `\t\t\treturn gopurs_runtime.Value{}\n` +
                                           `\t\t})`;
                                } else {
                                    let innerWrap = wrapReturn(innerT, "inner_res");
                                    return `gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {\n` +
                                           `\t\t\tinner_res := ${valName}(${argUnwrap})\n` +
                                           `\t\t\treturn ${innerWrap}\n` +
                                           `\t\t})`;
                                }
                            }
                        }
                    }
                    if (t.startsWith("[]") && t !== "[]gopurs_runtime.Value") {
                        return `func() gopurs_runtime.Value {\n` +
                               `\t\t\tres_arr := make([]gopurs_runtime.Value, len(${valName}))\n` +
                               `\t\t\tfor i, v := range ${valName} { res_arr[i] = gopurs_runtime.Box(v) }\n` +
                               `\t\t\treturn gopurs_runtime.Array(res_arr)\n` +
                               `\t\t}()`;
                    } else if (t === "gopurs_runtime.Value") {
                        return valName;
                    } else if (t === "") {
                        return `gopurs_runtime.Value{}`;
                    } else if (t.startsWith("map[")) {
                        return `func() gopurs_runtime.Value {\n` +
                               `\t\t\tres_map := make(map[string]gopurs_runtime.Value)\n` +
                               `\t\t\tfor k, v := range ${valName} { res_map[k] = gopurs_runtime.Box(v) }\n` +
                               `\t\t\treturn gopurs_runtime.Record(res_map)\n` +
                               `\t\t}()`;
                    } else {
                        return `gopurs_runtime.Box(${valName})`;
                    }
                };

                for (let idx = 0; idx < parsedArgs.length; idx++) {
                    let p = parsedArgs[idx];
                    let t = substituteGeneric(p.type);
                    
                    if (t.startsWith("func")) {
                        let unwrapStr = unwrapValueToFunc(t, 'arg' + idx, 0);
                        let indentedUnwrapStr = unwrapStr.split('\n').join('\n\t');
                        newLines.push(`\tgo_arg${idx} := ${indentedUnwrapStr}`);
                    } else if (t.startsWith("[]") && t !== "[]gopurs_runtime.Value") {
                        let elemType = t.substring(2);
                        if (elemType === "any") elemType = "interface{}";
                        newLines.push(`\targ${idx}_arr := *(*[]gopurs_runtime.Value)(arg${idx}.UnsafePtr)`);
                        newLines.push(`\tgo_arg${idx} := make(${t}, len(arg${idx}_arr))`);
                        if (elemType === "interface{}") {
                            newLines.push(`\tfor i, v := range arg${idx}_arr { go_arg${idx}[i] = v }`);
                        } else {
                            newLines.push(`\tfor i, v := range arg${idx}_arr { go_arg${idx}[i] = gopurs_runtime.Unbox[${elemType}](v) }`);
                        }
                    } else if (t === "any" || t === "interface{}") {
                        newLines.push(`\tgo_arg${idx} := arg${idx}`);
                    } else if (t === "gopurs_runtime.Value") {
                        newLines.push(`\tgo_arg${idx} := arg${idx}`);
                    } else if (t.startsWith("map[")) {
                        let elemType = t.substring(t.indexOf(']')+1);
                        if (elemType === "any" || elemType === "interface{}") {
                            newLines.push(`\targ${idx}_map := gopurs_runtime.RecordToMap(arg${idx})`);
                            newLines.push(`\tgo_arg${idx} := make(${t})`);
                            newLines.push(`\tfor k, v := range arg${idx}_map { go_arg${idx}[k] = v }`);
                        } else {
                            newLines.push(`\tgo_arg${idx} := arg${idx}.PtrVal().(${t})`);
                        }
                    } else {
                        newLines.push(`\tgo_arg${idx} := gopurs_runtime.Unbox[${t}](arg${idx})`);
                    }
                    callArgs.push(`go_arg${idx}`);
                }
                
                if (retStr === '') {
                    newLines.push(`\t${callFunc}(${callArgs.join(', ')})`);
                    newLines.push(`\treturn gopurs_runtime.Value{}`);
                } else {
                    newLines.push(`\tgo_res := ${callFunc}(${callArgs.join(', ')})`);
                    
                    let wrapCode = wrapReturn(substituteGeneric(retStr.trim()), "go_res");
                    newLines.push(`\treturn ${wrapCode}`);
                }
                newLines.push(`})`);
            } else {
                const varMatch = line.match(/^var\s+([A-Z_][A-Za-z0-9_]*)\s*(.*?)=\s*(.*)/);
                if (varMatch) {
                    const varName = varMatch[1];
                    let exportName = getExportName(varName);
                    newLines.push(`var ${exportName} = gopurs_runtime.Box(${varName})`);
                }
            }
        }

        let finalContent = text;
        if (newLines.length > 0) {
            finalContent = text + "\n\n// --- Auto-generated FFI wrappers ---\n" + newLines.join("\n") + "\n";
        }
        
        let goPackageName = moduleName.replace(/\./g, '_');
        let header = `package ${goPackageName}\n\n`;
        
        if (finalContent.indexOf("gopurs_runtime.") !== -1) {
            header += `import "gopurs/output/gopurs_runtime"\n\n`;
        }
        
        return header + finalContent;
    };
};
