import { readFileSync } from "node:fs";
import { CommandRunner, Interrupted } from "./command-runner.mjs";

export { Interrupted };

// Test-specific presentation and failure messages around shared process ownership.
export class TestProcesses {
  constructor() {
    this.commands = new CommandRunner();
  }

  checkInterrupted() {
    this.commands.checkInterrupted();
  }

  async run(label, command, args, { cwd, log, display = false } = {}) {
    this.checkInterrupted();
    console.log(`   [${label}] ${command} ${args.join(" ")}`);
    let status;
    try {
      status = await this.commands.run(command, args, { cwd, log });
    } catch (error) {
      if (error instanceof Interrupted) throw error;
      throw new Error(`${label}: ${error.message}`);
    }
    const output = log ? readFileSync(log, "utf8") : "";
    if ((display || status.code !== 0) && output) process.stdout.write(output.endsWith("\n") ? output : output + "\n");
    this.checkInterrupted();
    if (status.code !== 0) throw new Error(`${label} failed (${status.signal ?? "exit " + status.code})${log ? `; log: ${log}` : ""}`);
    return output;
  }

  dispose() {
    this.commands.dispose();
  }
}
