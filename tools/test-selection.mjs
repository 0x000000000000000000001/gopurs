import { accessSync, constants, readdirSync, statSync } from "node:fs";
import { basename, dirname, isAbsolute, join, resolve } from "node:path";

// Rechecked on 2026-10-02; StringEdgeCases reintegrated on 2026-10-03.
// Evidence and exact failures are in docs/testing.md.
const excluded = new Set([
  // The current frontend rejects these with CannotDeriveInvalidConstructorArg.
  "DerivingContravariant.purs", "DerivingFunctorFromBi.purs",
  "DerivingFunctorFromPro.purs", "DerivingProfunctor.purs",
  "NumberLiterals.purs", // Number Show spelling differs from the fixture's oracle.
  "StringEscapes.purs", // Folding concatenated surrogate halves differs from JS.
  "2136.purs", // 32-bit boundary overflow, with native 64-bit integers in gopurs.
]);

export class UsageError extends Error {}

export function parseOptions(args, { modules = false, env = process.env } = {}) {
  const options = { targets: [], clean: false, list: false, all: false, keep: false, keepGoing: false, update: false, resume: null, resumeFailed: null, help: false };
  if (!modules) {
    const update = env.UPDATE_SNAPSHOTS ?? "0";
    if (update !== "0" && update !== "1") throw new UsageError("UPDATE_SNAPSHOTS must be 0 or 1.");
    options.update = update === "1";
  }
  for (let i = 0; i < args.length; i++) {
    const arg = args[i];
    if (arg === "--") { options.targets.push(...args.slice(i + 1)); break; }
    if (arg === "-c" || arg === "--clean") options.clean = true;
    else if (arg === "--list") options.list = true;
    else if (arg === "--all") options.all = true;
    else if (arg === "--help" || arg === "-h") options.help = true;
    else if (!modules && arg === "--update-snapshots") options.update = true;
    else if (arg === "--keep-workspace") options.keep = true;
    else if (arg === "--keep-going") options.keepGoing = true;
    else if (arg === "--resume-failed") {
      const value = args[++i];
      if (!value || value.startsWith("-")) throw new UsageError("--resume-failed needs a report file.");
      options.resumeFailed = value;
    }
    else if (arg === "--skip-before") {
      const value = args[++i];
      if (!value || value.startsWith("-")) throw new UsageError("--skip-before needs a name.");
      options.resume = value;
    } else if (arg.startsWith("--skip-before=") || arg.startsWith("skip_before=")) {
      options.resume = arg.slice(arg.indexOf("=") + 1);
      if (!options.resume) throw new UsageError("--skip-before needs a name.");
    } else if (arg.startsWith("-")) throw new UsageError(`Unknown option: ${arg}`);
    else options.targets.push(arg);
  }
  if (options.all && options.targets.length) throw new UsageError("Use --all or explicit names, not both.");
  if (options.resumeFailed && (options.all || options.targets.length || options.resume)) {
    throw new UsageError("Use --resume-failed without targets, --all or --skip-before.");
  }
  return options;
}

function resumeFrom(items, options, matches) {
  if (!options.resume) return items;
  const index = items.findIndex(item => matches(item, options.resume));
  if (index < 0) throw new UsageError(`Resume target not found in selection: ${options.resume}`);
  return items.slice(index);
}

function isFile(path) {
  return statSync(path, { throwIfNoEntry: false })?.isFile() ?? false;
}

export function selectFixtures(root, options, cwd = process.cwd()) {
  const passing = join(root, "tests/passing");
  const candidates = options.targets.length ? options.targets.map(target => {
    const paths = isAbsolute(target) ? [target] : [resolve(root, target), resolve(cwd, target), join(passing, target), join(passing, target + ".purs")];
    const file = paths.find(path => path.endsWith(".purs") && isFile(path));
    if (!file) throw new UsageError(`Test file not found: ${target}`);
    return resolve(file);
  }) : readdirSync(passing).filter(name => name.endsWith(".purs") && isFile(join(passing, name))).sort().map(name => join(passing, name));
  const resumed = resumeFrom([...new Set(candidates)], options, (file, name) => basename(file) === name || basename(file) === name + ".purs");
  const skipped = resumed.filter(file => excluded.has(basename(file)));
  const selected = resumed.filter(file => !excluded.has(basename(file)));
  if (!selected.length) throw new UsageError("No runnable fixtures selected (the selection may contain only excluded fixtures).");
  return { selected, skipped };
}

export function selectModules(root, options) {
  const parent = dirname(root);
  const available = readdirSync(parent).filter(name => {
    if (!name.startsWith("gopurs-")) return false;
    const script = join(parent, name, "bin/test");
    if (!isFile(script)) return false;
    try { accessSync(script, constants.X_OK); return true; } catch { return false; }
  }).sort();
  const normalize = name => name.startsWith("gopurs-") ? name : "gopurs-" + name;
  const candidates = options.targets.length ? options.targets.map(target => {
    const name = normalize(target);
    if (!available.includes(name)) throw new UsageError(`No executable module test: ${target}`);
    return name;
  }) : available;
  const selected = resumeFrom([...new Set(candidates)], options, (name, target) => name === normalize(target));
  if (!selected.length) throw new UsageError("No module tests selected.");
  return selected.map(name => join(parent, name));
}
