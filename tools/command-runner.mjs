import { spawn } from "node:child_process";
import { closeSync, openSync } from "node:fs";

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
        }
      };
      process.on(signal, handler);
      return [signal, handler];
    }));
  }

  checkInterrupted() {
    if (this.signal) throw new Interrupted(this.signal);
  }

  async run(command, args, { cwd, env, log } = {}) {
    this.checkInterrupted();
    const fd = log ? openSync(log, "w") : null;
    try {
      const child = spawn(command, args, {
        cwd, env, detached: true,
        stdio: fd === null ? "inherit" : ["ignore", fd, fd],
      });
      this.active = child;
      return await new Promise((resolve, reject) => {
        child.once("error", reject);
        child.once("close", (code, signal) => resolve({ code, signal }));
      });
    } finally {
      this.active = null;
      if (fd !== null) closeSync(fd);
    }
  }

  dispose() {
    for (const [signal, handler] of this.handlers) process.off(signal, handler);
  }
}
