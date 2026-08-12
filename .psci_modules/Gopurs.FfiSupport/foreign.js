import { execSync } from 'child_process';
import { fileURLToPath } from 'url';
import { dirname, join } from 'path';

export const extractFfiAstImpl = function(moduleName) {
    return function(content) {
        return function() {
            try {
            const currentDir = dirname(fileURLToPath(import.meta.url));
            const toolsDir = join(currentDir, '..', 'tools');
            const wasmExec = join(toolsDir, 'wasm_exec.js');
            const wasmFile = join(toolsDir, 'ffi_gen.wasm');
            
            // We need a runner script to execute the WASM and call global.parseFFI
            // Since we don't want to create temp files on the fly, we can use node -e
            const runnerCode = `
const fs = require('fs');
require('${wasmExec.replace(/\\/g, '\\\\')}');
const go = new Go();
const wasmBuffer = fs.readFileSync('${wasmFile.replace(/\\/g, '\\\\')}');
WebAssembly.instantiate(wasmBuffer, go.importObject).then((result) => {
    go.run(result.instance);
    const content = fs.readFileSync(0, 'utf8');
    const jsonStr = global.parseFFI(content);
    console.log(jsonStr);
    process.exit(0);
}).catch(console.error);
`;
            
            const output = execSync(`node -e "${runnerCode.replace(/"/g, '\\"')}"`, {
                input: content,
                encoding: 'utf-8',
                maxBuffer: 10 * 1024 * 1024
            });
            return output;
            } catch (e) {
                console.error("FFI AST Extraction Error in module", moduleName);
                if (e.stdout) console.error(e.stdout.toString());
                if (e.stderr) console.error(e.stderr.toString());
                throw e;
            }
        };
    };
};
