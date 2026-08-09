# gopurs

<img height="160" alt="Screenshot 2026-07-21 at 17 22 56" src="https://github.com/user-attachments/assets/b013e7c3-fac6-4ee8-9d4c-f39ac8c2c921" />
<br />
<br />

_Mature experimental WIP. The core ideas have been valided by the facts/results. That will be official soon. You can [find a complete devlog here](https://discourse.purescript.org/t/leveraging-a-blazing-fast-runtime-a-new-go-backend-for-purescript)._

A super-optimized **PureScript-to-Go compiler**, entirely written in PureScript, leveraging Go's **blazing-fast execution**, **lightweight goroutines** and **huge ecosystem**. 

`gopurs` leverages an enriched `tcorefn` (Typed CoreFn) representation to compile your pure business logic into robust, modern Go code. It seamlessly integrates into your existing PureScript workflow as a custom backend.

## Why Go?

While the broader JS ecosystem has heavily leaned towards TypeScript, many backend services, CLIs, and infrastructure tools rely heavily on Go for its **raw performance**, **concurrency model** (goroutines), and **deployment simplicity** (single static binaries).

`gopurs` aims to provide a bridge for developers who want the elegance and strict typing of a purely functional language like PureScript, while benefiting from Go's massive ecosystem. It opens a door for those who want to compile their pure business logic into a single, zero-dependency static binary that can run anywhere.

## Why a new Go backend?

The [`purescript-native`](https://github.com/andyarvanitis/purescript-native) project already provides a Go compiler backend. I want to deeply acknowledge the fantastic work done by Andy Arvanitis on that project, which paved the way for compiling PureScript to native targets. This new project is largely inspired by his initial effort, and my gratitude for his pioneering work is very real. It is always easier to come second and learn from the technical limits encountered by the pioneers.

Reading through the discussions and challenges raised by users over the years (initialization orders, performance overhead of `interface{}`, module qualifications), it became clear that the ecosystem has evolved drastically. This evolution unlocked new architectural paradigms that make building a completely new Go backend highly relevant today, specifically to address these past limitations:

### 1. The optimizer & bootstrapping
While previous native compilers were often written in Haskell and parsed raw `CoreFn`, `gopurs` is written 100% in PureScript. It integrates directly with the [`purescript-backend-optimizer`](https://github.com/aristanetworks/purescript-backend-optimizer) (just like `purs-backend-es` or `phpurs`). This allows the compiler to instantly benefit from classical optimizations such as aggressive uncurrying, magic-do, and Tail Call Optimization (TCO) at the AST level. The `gopurs` compiler can then strictly focus on translating this highly-optimized AST into idiomatic, performant Go code. Being built in PureScript also ensures it remains fully accessible to anyone in the ecosystem (installable via `spago` and `npm`).

### 2. Heap vs stack: a new memory layout for Go
Dynamic typing in statically typed languages like Go often relies heavily on `interface{}` (or `any`). Previous compilers represented PureScript values as `any` and PureScript records as `map[string]any`. However, assigning primitive values to interfaces forces them to escape to the heap (boxing), generating massive Garbage Collector pressure. For `gopurs`, I ran extensive benchmarks and decided to completely ditch `any`. Instead, the runtime uses a universal flat `Value` struct (a tagged union). This ensures that dynamic operations stay mostly on the stack.

> **Benchmark context:** On a 1 billion operations benchmark, native static Go took ~250ms, a dynamic `any` approach took ~9 seconds, and my `Value` struct solution completed in **~240ms**. This memory model completely bypasses the GC overhead in heavy iterative loops.

### 3. Up-to-date with modern PureScript
`gopurs` aims to be fully aligned with the current v0.15+ ecosystem (and v0.16+ soon). I am currently mirroring the standard libraries ([`gopurs-prelude`](https://github.com/0x000000000000000000001/gopurs-prelude), [`gopurs-effect`](https://github.com/0x000000000000000000001/gopurs-effect), etc.) to provide native Go FFIs.

### 4. Concurrency: Aff mapped to goroutines
The previous compiler listed goroutine parallelism as a "future idea", which meant it lacked an implementation to run `Aff` asynchronously out of the box. `gopurs` introduces an event loop emulator and natively maps `Aff` to goroutines. This brings true, shared-memory parallelism to PureScript, allowing your asynchronous code to scale effortlessly across multiple CPU cores.

## How to use

If you wish to configure an existing project, `gopurs` acts as a drop-in backend for the Spago build system.

1. **Install the `gopurs` backend compiler:**
   You can install the compiler directly from GitHub. NPM will automatically compile it in the background during installation.
   ```bash
   npm install --save-dev github:0x000000000000000000001/gopurs
   ```

2. **Manage Core Library Overrides (`spago.yaml`):**
   Because standard PureScript libraries use JavaScript FFI, you must override them with their `gopurs-*` counterparts. Keep using the official PureScript registry as your base, and manually define all Go overrides using the `extraPackages` directive.

   ```yaml
   workspace:
     packageSet:
       registry: 77.10.1
     extraPackages:
       prelude:
         git: "https://github.com/0x000000000000000000001/gopurs-prelude.git"
         ref: "master"
         dependencies: []
       # ... all other gopurs-* packages
     backend:
       cmd: gopurs
   ```
   *Alternatively, you can pass the backend directly via CLI:*
   ```bash
   spago build --backend gopurs
   ```

3. **Build and execute:**
   The compiler will parse all `tcorefn.json` files generated by `purs` (via a TAST-enabled fork) and output native Go files in the `output/` directory.
   
   An executable `main.go` entrypoint will be automatically generated. You can run it directly by initializing a Go module in the output folder:
   
   ```bash
   spago build
   cd output
   go mod init gopurs
   go mod tidy
   go run main.go
   ```

### Compiler configuration options

The `gopurs` compiler is entirely **zero-config by default**. It will automatically scan your `tcorefn` ASTs and generate a ready-to-execute `main.go` entrypoint.

If you need advanced behavior, you can pass arguments to the `gopurs` compiler by appending them to the `spago build --backend-args` command:

```bash
spago build --backend gopurs --backend-args "--main App.Main"
```

| Option | Description |
|---|---|
| `--main <Module>` | *Optional*. Explicitly sets the entrypoint module. Without this flag, `gopurs` automatically targets the `Main` module. |

## Local development & testing

If you plan to contribute to the compiler or run the official test suite locally, you will have to follow a specific "sibling-checkout" directory layout. 

Because `gopurs` replaces the JS ecosystem with Go, it requires custom Go-compatible forks of the core PureScript libraries (e.g. `purescript-prelude` becomes `gopurs-prelude`). The internal test runner (`bin/test`) expects these core `gopurs-*` repositories to be cloned side-by-side in the same parent directory as the main `gopurs` repository.

```
workspace/
├── gopurs/
├── gopurs-prelude/
├── gopurs-effect/
├── gopurs-console/
├── gopurs-assert/
└── ... (all other core gopurs-* forks)
```

To easily clone all these required dependencies, you can simply run the provided setup script:
```bash
cd gopurs
./bin/setup
```

To run the test suite:
```bash
./bin/test
```

## Current status & milestones

Since its inception, `gopurs` has reached several major milestones:

- [x] **100% of the [official tests](https://github.com/purescript/purescript/tree/master/tests/purs/passing) are green.**
- [x] **Typed AST (TAST):** By consuming an enriched `tcorefn.json` (Typed CoreFn) instead of standard `corefn`, `gopurs` preserves deep structural typing. Combined with partial monomorphization, this unlocks massive performance gains (x10).
- [x] **Zero boilerplate FFI:** A complete overhaul of the FFI developer experience via a WebAssembly parser (`ffi_gen.wasm`). It analyzes your Go signatures on the fly, allowing you to write idiomatic Go (uncurried functions, native types, flexible Effect semantics) without manual boxing or closures.
- [x] **Native `Aff` via goroutines:** Full support for `Aff` mapped directly to Go's goroutines. This emulates the event loop while providing **true multi-core parallelism for free**, meaning your async PureScript code gets exponentially faster on multi-core systems.
- [x] **Real world validation (unit):** Successful validation on 100% of the unit tests for a complex, full scale project involving Postgres, S3, RabbitMQ, and deep Aff nesting.
- [x] **Real world validation (integration):** Successful validation on 100% of the integration tests for a complex, full scale project involving Postgres, S3, RabbitMQ, and deep Aff nesting.
- [ ] **Module validation:** Validate tests module by module (`gopurs-*`).
- [ ] General code **cleanup** (it’s still quite messy)

_(maybe more to come)_

## Architecture

`gopurs` is built on top of [Arista's purescript-backend-optimizer](https://github.com/aristanetworks/purescript-backend-optimizer) to avoid reinventing the optimization wheel. The compilation pipeline is functionally decoupled:

1. **Optimization**: The optimizer reads the `tcorefn.json` generated by `purs`, performs aggressive Dead Code Elimination (DCE), typeclass dictionary resolution, inlining, and constant folding at the AST level, and outputs an optimized `BackendModule`.
2. **Code Generation**: `Gopurs.CodeGen` maps this heavily optimized PureScript AST to our native `GoAst`.
3. **Printing**: `Gopurs.Printer` formats the Go AST into valid, modern Go syntax.
4. **Caching & CLI**: `Main` orchestrates the CLI, writing the generated `.go` files to their respective module directories. 

## License

MIT License. See [LICENSE](LICENSE) for details.
