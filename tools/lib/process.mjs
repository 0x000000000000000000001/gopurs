import { spawn } from "node:child_process";
import { closeSync, openSync, writeFileSync } from "node:fs";

export class Interrupted extends Error {
  constructor(signal) {
    super(`Interrupted by ${signal}`);
    this.exitCode = signal === "SIGINT" ? 130 : 143;
  }
}

// Own one command at a time and its log descriptor. Each command gets a process
// group so interrupts also reach compiler grandchildren. Callers own phase
// names, diagnostics and interpretation of nonzero exit codes.
export class CommandRunner {
  constructor() {
    this.active = null;
    this.signal = null;
    this.handlers = new Map(["SIGINT", "SIGTERM"].map(signal => {
      const handler = () => {
        this.signal = signal;
        if (this.active?.pid) {
          try { process.kill(-this.active.pid, signal); }
          catch (error) { if (error.code !== "ESRCH") throw error; }
          this.active.stdout?.resume();
          this.active.stderr?.resume();
        }
      };
      process.on(signal, handler);
      return [signal, handler];
    }));
  }

  checkInterrupted() {
    if (this.signal) throw new Interrupted(this.signal);
  }

  async run(command, args, { cwd, env, log, stream = false } = {}) {
    this.checkInterrupted();
    const fd = log ? openSync(log, "w") : null;
    const cleanups = [];
    try {
      const child = spawn(command, args, {
        cwd, env, detached: true,
        stdio: fd === null ? "inherit" : stream ? ["ignore", "pipe", "pipe"] : ["ignore", fd, fd],
      });
      this.active = child;
      let failure;
      const pendingWrites = new Set();
      const fail = error => {
        if (failure) return;
        failure = error;
        if (child.pid) {
          try { process.kill(-child.pid, "SIGTERM"); }
          catch (error) { if (error.code !== "ESRCH") failure.cause ??= error; }
        }
        child.stdout?.resume();
        child.stderr?.resume();
      };
      if (fd !== null && stream) {
        for (const [input, output] of [[child.stdout, process.stdout], [child.stderr, process.stderr]]) {
          input.on("error", fail);
          output.on("error", fail);
          cleanups.push(() => output.off("error", fail));
          input.on("data", chunk => {
            if (failure) return;
            try {
              writeFileSync(fd, chunk);
              const pending = new Promise(resolve => {
                try {
                  const writable = output.write(chunk, error => {
                    if (error) fail(error);
                    input.resume();
                    resolve();
                  });
                  if (!writable && !this.signal) input.pause();
                } catch (error) {
                  fail(error);
                  resolve();
                }
              });
              pendingWrites.add(pending);
              pending.then(() => pendingWrites.delete(pending));
            } catch (error) { fail(error); }
          });
        }
      }
      const status = await new Promise(resolve => {
        child.once("error", fail);
        child.once("close", (code, signal) => resolve({ code, signal }));
      });
      await Promise.all(pendingWrites);
      if (failure) throw failure;
      return status;
    } finally {
      this.active = null;
      for (const cleanup of cleanups) cleanup();
      if (fd !== null) closeSync(fd);
    }
  }

  dispose() {
    for (const [signal, handler] of this.handlers) process.off(signal, handler);
  }
}
