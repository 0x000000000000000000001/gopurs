import { readFileSync } from 'node:fs';
import './wasm_exec.js';

async function run() {
    const go = new globalThis.Go();
    const wasmBuffer = readFileSync(new URL('./ffi_gen.wasm', import.meta.url));
    const result = await WebAssembly.instantiate(wasmBuffer, go.importObject);
    // The Go main registers parseFFI, then waits for calls indefinitely.
    // Observe startup failures without waiting for the runtime to finish.
    go.run(result.instance).catch(fail);
    const content = readFileSync(0, 'utf8');
    const response = globalThis.parseFFI(content);
    if (!response || typeof response.json !== 'string' || typeof response.error !== 'string') {
        throw new Error('Invalid Go FFI parser response; rebuild with npm run build:ffi');
    }
    if (response.error) {
        throw new Error(`Go FFI parse failed: ${response.error}`);
    }
    process.stdout.write(response.json + '\n', () => process.exit(0));
}

function fail(error) {
    // Go may still be waiting; an exitCode alone would leave it alive.
    process.stderr.write(`FFI runner: ${error.message}\n`, () => process.exit(1));
}

run().catch(fail);
