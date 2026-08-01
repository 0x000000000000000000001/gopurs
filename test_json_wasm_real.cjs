const fs = require('fs');
require('./tools/wasm_exec.js');

const go = new Go();
const wasmBuffer = fs.readFileSync('tools/ffi_gen.wasm');

WebAssembly.instantiate(wasmBuffer, go.importObject).then((result) => {
    go.run(result.instance);
    const code = fs.readFileSync('tests/passing/FFIConstraintWorkaround.go', 'utf8');
    const jsonStr = global.parseFFI(code);
    console.log(JSON.stringify(JSON.parse(jsonStr), null, 2));
    process.exit(0);
}).catch(console.error);
