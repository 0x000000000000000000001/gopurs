# gopurs

<img height="160" alt="Screenshot 2026-07-21 at 17 22 56" src="https://github.com/user-attachments/assets/b013e7c3-fac6-4ee8-9d4c-f39ac8c2c921" />
<br />
<br />

A super-optimized **PureScript-to-Go compiler**, entirely written in PureScript, leveraging Go's **blazing-fast execution**, **lightweight goroutines** and **huge ecosystem**.

> **Note:** This project is currently a work in progress (WIP) and highly experimental, but the first official tests are already showing green and positive results! 🟢 The repo may be ready for production soon, this summer 🌻.

## Why a new Go backend?

You might legitimately ask: *"Wait, doesn't `purescript-native` (and its `psgo` tool) already do this?"*. And yes, it does! 

I want to deeply acknowledge the fantastic work done by Andy Arvanitis on `purescript-native`. It really paved the way. It is always much easier to come second and learn from the limits encountered by the pioneers. This project is a response, an extension of his work, an apology, a love letter. It simply couldn’t exist without it. All my gratitude for his initial effort. ❤️

However, the ecosystem has evolved drastically over the last few years. This unlocked new architectural paradigms that make building a completely new Go backend highly relevant today, specifically to address previous limitations:

### 1. The Optimizer & Bootstrapping 🧠
While `purescript-native` was written in Haskell and parsed raw `CoreFn`, `gopurs` is written 100% in PureScript. It plugs directly into the `purescript-backend-optimizer` (which powers `purs-backend-es`). This means we do not reinvent the wheel for classical optimizations. The `gopurs` compiler can strictly focus on translating this highly-optimized AST into idiomatic, performant Go code, adding further optimizations specifically for Go. And it remains fully accessible to anyone in the PureScript ecosystem (installable via `spago` and `npm`).

### 2. Heap vs. Stack: a new memory layout for Go ⚡
Dynamic typing in statically typed languages like Go often relies heavily on `interface{}` (or `any`). However, assigning primitive values to interfaces forces them to escape to the heap (Boxing), generating massive Garbage Collector pressure. For `gopurs`, I ran extensive benchmarks and decided to completely ditch `any`. Instead, the runtime uses a universal flat `Value` struct (a tagged union). This ensures that dynamic operations stay mostly on the stack.

> To give you an idea, on a **1 billion operations benchmark**, native static Go took ~250ms, a dynamic `any` approach took ~9 seconds, and our `Value` struct solution completed in **~240ms**! The performance difference is staggering, completely bypassing the GC overhead in heavy iterative loops.

### 3. Up-to-date with modern PureScript 📦
`gopurs` aims to be fully aligned with the current v0.15+ ecosystem (and v0.16+ soon). We are currently mirroring the standard libraries (`gopurs-prelude`, `gopurs-effect`, etc.) to provide native Go FFIs.

Go offers exceptional concurrency (goroutines fit perfectly with PureScript's `Aff`) and rock-solid native binaries. 

## License

MIT License. See [LICENSE](LICENSE) for details.
