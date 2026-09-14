import { spawn } from "node:child_process";
import { closeSync, openSync, readFileSync } from "node:fs";

export class Interrupted extends Error {
  constructor(signal) {
    super(`Interrupted by ${signal}`);
    this.exitCode = signal === "SIGINT" ? 130 : 143;
  }
}

// Give each command its own process group so interrupting the runner also
// stops compilers and grandchildren, without touching another test run.
export class TestProcesses {
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

  async run(label, command, args, { cwd, log, display = false } = {}) {
    this.checkInterrupted();
    console.log(`   [${label}] ${command} ${args.join(" ")}`);
    const fd = log ? openSync(log, "w") : null;
    let status;
    try {
      const child = spawn(command, args, { cwd, detached: true, stdio: fd === null ? "inherit" : ["ignore", fd, fd] });
      this.active = child;
      status = await new Promise((resolve, reject) => {
        child.once("error", reject);
        child.once("close", (code, signal) => resolve({ code, signal }));
      });
    } catch (error) {
      throw new Error(`${label}: ${error.message}`);
    } finally {
      this.active = null;
      if (fd !== null) closeSync(fd);
    }
    const output = log ? readFileSync(log, "utf8") : "";
    if ((display || status.code !== 0) && output) process.stdout.write(output.endsWith("\n") ? output : output + "\n");
    this.checkInterrupted();
    if (status.code !== 0) throw new Error(`${label} failed (${status.signal ?? "exit " + status.code})${log ? `; log: ${log}` : ""}`);
    return output;
  }

  dispose() {
    for (const [signal, handler] of this.handlers) process.off(signal, handler);
  }
}
