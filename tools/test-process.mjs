import { existsSync, readFileSync } from "node:fs";
import { delimiter, resolve } from "node:path";
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

  async run(label, command, args, { cwd, env, log, display = false, stream = false } = {}) {
    this.checkInterrupted();
    console.log(`   [${label}] ${command} ${args.join(" ")}`);
    let status;
    try {
      status = await this.commands.run(command, args, { cwd, env, log, stream });
    } catch (error) {
      if (error instanceof Interrupted) throw error;
      if (error.code === "ENOENT") {
        const directory = cwd ?? process.cwd();
        const search = command.includes("/") ? [resolve(directory, command)]
          : ((env ?? process.env).PATH ?? "/usr/bin:/bin").split(delimiter).map(path => resolve(directory, path, command));
        const describe = path => ({ path, exists: existsSync(path) });
        error.message += `; spawn context: ${JSON.stringify({ cwd: describe(directory), command: search.map(describe) })}`;
      }
      throw new Error(`${label}: ${error.message}`);
    }
    const output = log ? readFileSync(log, "utf8") : "";
    if (!stream && (display || status.code !== 0) && output) process.stdout.write(output.endsWith("\n") ? output : output + "\n");
    this.checkInterrupted();
    if (status.code !== 0) throw new Error(`${label} failed (${status.signal ?? "exit " + status.code})${log ? `; log: ${log}` : ""}`);
    return output;
  }

  dispose() {
    this.commands.dispose();
  }
}
