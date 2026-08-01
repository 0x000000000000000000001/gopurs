const fs = require('fs');
require('./tools/wasm_exec.js');

const go = new Go();
const wasmBuffer = fs.readFileSync('tools/ffi_gen.wasm');

WebAssembly.instantiate(wasmBuffer, go.importObject).then((result) => {
    // Run the go program, which registers parseFFI globally and then blocks
    go.run(result.instance);
    
    // Now we can call global.parseFFI
    const code = `
package main
var Default_ = 42
func Map_[T any](a int, b string) bool { return true }
`;
    const jsonStr = global.parseFFI(code);
    console.log(jsonStr);
    process.exit(0);
}).catch(console.error);
