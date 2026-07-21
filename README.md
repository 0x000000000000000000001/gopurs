# gopurs

<img height="160" alt="Screenshot 2026-07-21 at 17 22 56" src="https://github.com/user-attachments/assets/b013e7c3-fac6-4ee8-9d4c-f39ac8c2c921" />
<br />
<br />

A super-optimized **PureScript-to-Go compiler**, leveraging its **blazing-fast execution**, **lightweight goroutines** and **huge ecosystem**.

> **Note:** This project is currently a work in progress (WIP).

`gopurs` leverages the intermediate `CoreFn` representation to compile your pure business logic into robust, modern Go code. It seamlessly integrates into your existing PureScript workflow as a custom backend.

## Why Go?

While the broader JS ecosystem has heavily leaned towards TypeScript, many backend services and infrastructure tools rely heavily on Go for its **raw performance**, **concurrency model**, and **deployment simplicity**.

`gopurs` aims to provide a bridge for developers who want the elegance and strict typing of a purely functional language like PureScript, while benefiting from Go's **massive ecosystem**. It opens a door for those who want to compile their pure business logic into a **single, zero-dependency static binary** that can run anywhere.

## Architecture and status

This project is currently in an experimental phase.

The main challenge and current focus is to successfully map PureScript's core concepts (currying, tail call optimization, structural records) into Go's nominal type system and execution model without sacrificing performance.

## License

MIT License. See [LICENSE](LICENSE) for details.
