import fs from 'fs';
import path from 'path';

let cachedScanDirs = null;

function getScanDirs(mbFfiDir) {
    if (cachedScanDirs !== null) return cachedScanDirs;
    
    const rootDir = process.cwd();
    const scanDirs = [];
    
    const spagoDirs = [
        path.join(rootDir, '.spago'),
        path.join(rootDir, 'spago.d')
    ];
    
    for (const spagoDir of spagoDirs) {
        if (fs.existsSync(spagoDir) && fs.statSync(spagoDir).isDirectory()) {
            const packages = fs.readdirSync(spagoDir);
            for (const pkg of packages) {
                const pkgDir = path.join(spagoDir, pkg);
                if (fs.statSync(pkgDir).isDirectory()) {
                    let hasVersion = false;
                    const subdirs = fs.readdirSync(pkgDir);
                    for (const subdir of subdirs) {
                        const versionDir = path.join(pkgDir, subdir);
                        if (subdir.startsWith('v') && fs.statSync(versionDir).isDirectory()) {
                            scanDirs.push(versionDir);
                            hasVersion = true;
                        }
                    }
                    if (!hasVersion) {
                        scanDirs.push(pkgDir);
                    }
                }
            }
        }
    }
    
    if (mbFfiDir) {
        scanDirs.push(path.join(rootDir, mbFfiDir));
    } else {
        // Fallback to searching the local src/ dir if mbFfiDir is not provided
        scanDirs.push(rootDir);
    }
    
    cachedScanDirs = scanDirs;
    return scanDirs;
}

let goFileIndex = null;

function buildGoFileIndex(scanDirs) {
    if (goFileIndex !== null) return;
    goFileIndex = new Set();
    
    function walk(dir) {
        let entries;
        try {
            entries = fs.readdirSync(dir, { withFileTypes: true });
        } catch (e) {
            return;
        }
        for (const entry of entries) {
            const res = path.join(dir, entry.name);
            if (entry.isDirectory()) {
                walk(res);
            } else if (entry.name.endsWith('.go')) {
                goFileIndex.add(res);
            }
        }
    }
    
    for (const d of scanDirs) {
        walk(d);
    }
}

export const findFfiFileImpl = function(mbFfiDir) {
    return function(modNameStr) {
        return function(mbModulePath) {
            return function() {
                if (mbModulePath) {
                    const goPath = mbModulePath.replace(/\.purs$/, '.go');
                    if (fs.existsSync(goPath)) {
                        return goPath;
                    }
                }
                
                const scanDirs = getScanDirs(mbFfiDir);
                buildGoFileIndex(scanDirs);
                
                for (const dir of scanDirs) {
                    // Search in dir/src/Data/Show.go, dir/src/Show.go, etc.
                    const searchPaths = [
                        path.join(dir, 'src', ...modNameStr.split('.')) + '.go',
                        path.join(dir, 'src', modNameStr + '.go'),
                        path.join(dir, modNameStr + '.go')
                    ];
                    for (const p of searchPaths) {
                        if (goFileIndex.has(p)) {
                            return p;
                        }
                    }
                }
                return null;
            };
        };
    };
};


export const appendFfiWrappersImpl = function(moduleName) {
    return function(content) {
        let text = content.replace(/^package\s+[a-zA-Z0-9_]+[\s\n]*/m, '');
        text = text.replace(/^[ \t]*import[ \t]+"gopurs\/output\/gopurs_runtime"[\s\n]*/gm, '');
        text = text.replace(/^[ \t]*"gopurs\/output\/gopurs_runtime"[\s\n]*/gm, '');
        text = text.replace(/import\s*\(\s*\)[\s\n]*/gm, '');

        const lines = text.split('\n');
        let newLines = [];

        for (let i = 0; i < lines.length; i++) {
            const line = lines[i];
            const basicMatch = line.match(/^func\s+(_?[A-Z][A-Za-z0-9_]*)\s*\(/);
            if (basicMatch) {
                const funcName = basicMatch[1];
                let exportName = "_Gopurs_" + funcName;
                
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
                if (remainder.endsWith('{')) {
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
                if (arity > 5) continue; // Not supported yet

                let funcConstructor = arity === 0 ? "Func" : (arity === 1 ? "Func" : `Func${arity}`);
                let goFuncArgs = parsedArgs.map((_, idx) => `arg${idx} gopurs_runtime.Value`).join(', ');
                if (arity === 0) goFuncArgs = `_ gopurs_runtime.Value`;
                if (arity === 1) goFuncArgs = `arg0 gopurs_runtime.Value`;
                
                newLines.push(`var ${exportName} = gopurs_runtime.${funcConstructor}(func(${goFuncArgs}) gopurs_runtime.Value {`);
                
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
                    for (let j = 0; j < argsStr.length; j++) {
                        let c = argsStr[j];
                        if (c === '(') parens++;
                        else if (c === ')') parens--;
                        else if (c === '[') brackets++;
                        else if (c === ']') brackets--;
                        else if (c === ',' && parens === 0 && brackets === 0) {
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
                        return `func(${params}) ${retStr} {\n\t\tinner_res${depth} := ${applyCall}\n\t\tres_arr${depth} := inner_res${depth}.PtrVal.([]gopurs_runtime.Value)\n\t\tres_go${depth} := make(${retStr}, len(res_arr${depth}))\n\t\tfor i, v := range res_arr${depth} { res_go${depth}[i] = gopurs_runtime.Unbox[${elemType}](v) }\n\t\treturn res_go${depth}\n\t}`;
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
                                let argUnwrap = "arg.PtrVal";
                                if (argT === "any" || argT === "interface{}") {
                                    argUnwrap = "arg.PtrVal";
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
                    let t = p.type;
                    
                    if (t.startsWith("func")) {
                        let unwrapStr = unwrapValueToFunc(t, 'arg' + idx, 0);
                        let indentedUnwrapStr = unwrapStr.split('\n').join('\n\t');
                        newLines.push(`\tgo_arg${idx} := ${indentedUnwrapStr}`);
                    } else if (t.startsWith("[]") && t !== "[]gopurs_runtime.Value") {
                        let elemType = t.substring(2);
                        if (elemType === "any") elemType = "interface{}";
                        newLines.push(`\targ${idx}_arr := arg${idx}.PtrVal.([]gopurs_runtime.Value)`);
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
                            newLines.push(`\tgo_arg${idx} := arg${idx}.PtrVal.(${t})`);
                        }
                    } else {
                        newLines.push(`\tgo_arg${idx} := gopurs_runtime.Unbox[${t}](arg${idx})`);
                    }
                    callArgs.push(`go_arg${idx}`);
                }
                
                if (retStr === '') {
                    newLines.push(`\t${funcName}(${callArgs.join(', ')})`);
                    newLines.push(`\treturn gopurs_runtime.Value{}`);
                } else {
                    newLines.push(`\tgo_res := ${funcName}(${callArgs.join(', ')})`);
                    
                    let wrapCode = wrapReturn(retStr, "go_res");
                    newLines.push(`\treturn ${wrapCode}`);
                }
                newLines.push(`})`);
            } else {
                const varMatch = line.match(/^var\s+([A-Z][A-Za-z0-9_]*)\s*(.*?)=\s*(.*)/);
                if (varMatch) {
                    const varName = varMatch[1];
                    let exportName = "_Gopurs_" + varName;
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
