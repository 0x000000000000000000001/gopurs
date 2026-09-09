import { readFileSync } from 'node:fs';
import './wasm_exec.js';

const go = new globalThis.Go();
const wasmBuffer = readFileSync(new URL('./ffi_gen.wasm', import.meta.url));

// Preserve the existing rejection handling until the FFI error contract changes.
WebAssembly.instantiate(wasmBuffer, go.importObject).then((result) => {
    go.run(result.instance);
    const content = readFileSync(0, 'utf8');
    console.log(globalThis.parseFFI(content));
    process.exit(0);
}).catch(console.error);
