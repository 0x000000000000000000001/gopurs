#!/usr/bin/env node
// Compare two output trees, or a tree against a recorded manifest, without
// reading or interpreting the generated Go. Paths and contents must match
// exactly; this is the guardrail for behavior-preserving cleanup.
import { createHash } from "node:crypto";
import {
  existsSync, lstatSync, mkdirSync, readdirSync, readFileSync, readlinkSync,
  statSync, writeFileSync,
} from "node:fs";
import { basename, dirname, join, relative, resolve, sep } from "node:path";

const VERSION = 1;
const MAX_REPORT = 50;

const usage = () => {
  console.error(`Usage:
  node tools/check-output-identity.mjs snapshot <root> <manifest>
  node tools/check-output-identity.mjs compare  <root> <manifest>
  node tools/check-output-identity.mjs diff     <baseline-root> <current-root>

Options:
  --ignore <glob>     Skip matching relative paths (repeatable).
  --max-report <n>    Show at most n differences (default ${MAX_REPORT}).

A manifest is JSON: { version, algorithm, files: { path: "sha256:<hex>" } }.
Symbolic links are recorded as "symlink:<target>". Ignored paths are not
represented in a manifest and must be passed again when comparing.`);
};

const toPosix = path => path.split(sep).join("/");
const sha256 = path => createHash("sha256").update(readFileSync(path)).digest("hex");

const globToRegExp = pattern => {
  let source = "";
  for (let index = 0; index < pattern.length; index += 1) {
    const character = pattern[index];
    if (character === "*") {
      if (pattern[index + 1] === "*") {
        if (pattern[index + 2] === "/") {
          source += "(?:.*/)?";
          index += 2;
        } else {
          source += ".*";
          index += 1;
        }
      } else {
        source += "[^/]*";
      }
    } else if (character === "?") {
      source += "[^/]";
    } else if ("\\^$+.()|{}[]".includes(character)) {
      source += `\\${character}`;
    } else {
      source += character;
    }
  }
  return new RegExp(`^${source}$`);
};

const compileIgnores = patterns => patterns.map(pattern => ({
  matches: globToRegExp(pattern),
  anyDepth: !pattern.includes("/"),
}));

const isIgnored = (path, ignores) => ignores.some(({ matches, anyDepth }) =>
  matches.test(path) || (anyDepth && matches.test(basename(path))));

const walk = root => {
  const entries = new Map();
  const visit = directory => {
    for (const name of readdirSync(directory).sort()) {
      const absolute = join(directory, name);
      const path = toPosix(relative(root, absolute));
      const stats = lstatSync(absolute);
      if (stats.isSymbolicLink()) {
        entries.set(path, `symlink:${readlinkSync(absolute)}`);
      } else if (stats.isDirectory()) {
        visit(absolute);
      } else if (stats.isFile()) {
        entries.set(path, `sha256:${sha256(absolute)}`);
      } else {
        throw new Error(`unsupported file type: ${absolute}`);
      }
    }
  };
  visit(root);
  return entries;
};

const readManifest = path => {
  const manifest = JSON.parse(readFileSync(path, "utf8"));
  if (manifest.version !== VERSION || manifest.algorithm !== "sha256" || typeof manifest.files !== "object") {
    throw new Error(`unsupported manifest: ${path}`);
  }
  return new Map(Object.entries(manifest.files));
};

const writeManifest = (path, entries) => {
  const files = Object.fromEntries([...entries.entries()].sort(([left], [right]) => left.localeCompare(right)));
  const manifest = { version: VERSION, algorithm: "sha256", files };
  mkdirSync(dirname(resolve(path)), { recursive: true });
  writeFileSync(path, `${JSON.stringify(manifest, null, 2)}\n`);
};

const compare = (expected, actual, ignores) => {
  const differences = [];
  for (const [path, value] of expected) {
    if (isIgnored(path, ignores)) continue;
    if (!actual.has(path)) differences.push({ kind: "missing", path });
    else if (actual.get(path) !== value) differences.push({ kind: "changed", path });
  }
  for (const path of actual.keys()) {
    if (expected.has(path) || isIgnored(path, ignores)) continue;
    differences.push({ kind: "extra", path });
  }
  differences.sort((left, right) => left.path.localeCompare(right.path) || left.kind.localeCompare(right.kind));
  return differences;
};

const printDifferences = (differences, maximum) => {
  const counts = differences.reduce((total, { kind }) => {
    total[kind] = (total[kind] ?? 0) + 1;
    return total;
  }, {});
  console.error(`[identity] ${differences.length} difference(s): ` +
    ["missing", "extra", "changed"].map(kind => `${counts[kind] ?? 0} ${kind}`).join(", "));
  for (const { kind, path } of differences.slice(0, maximum)) {
    console.error(`  ${kind.padEnd(7)} ${path}`);
  }
  if (differences.length > maximum) {
    console.error(`  ... ${differences.length - maximum} more`);
  }
};

const parseArguments = args => {
  const positional = [];
  const ignores = [];
  let maximum = MAX_REPORT;
  for (let index = 0; index < args.length; index += 1) {
    const argument = args[index];
    if (argument === "--ignore") {
      const pattern = args[index + 1];
      if (!pattern) throw new Error("--ignore needs a pattern");
      ignores.push(pattern);
      index += 1;
    } else if (argument === "--max-report") {
      const value = Number(args[index + 1]);
      if (!Number.isInteger(value) || value < 1) throw new Error("--max-report needs a positive integer");
      maximum = value;
      index += 1;
    } else if (argument.startsWith("--")) {
      throw new Error(`unknown option: ${argument}`);
    } else {
      positional.push(argument);
    }
  }
  return { positional, ignores: compileIgnores(ignores), maximum };
};

const requireDirectory = path => {
  if (!existsSync(path) || !statSync(path).isDirectory()) {
    throw new Error(`not a directory: ${path}`);
  }
  return resolve(path);
};

const main = () => {
  const [command, ...rest] = process.argv.slice(2);
  if (!command || command === "--help" || command === "-h" || command === "help") {
    usage();
    process.exit(command ? 0 : 2);
  }
  const { positional, ignores, maximum } = parseArguments(rest);

  if (command === "snapshot") {
    if (positional.length !== 2) throw new Error("snapshot needs <root> <manifest>");
    const root = requireDirectory(positional[0]);
    const entries = walk(root);
    const kept = new Map([...entries].filter(([path]) => !isIgnored(path, ignores)));
    writeManifest(positional[1], kept);
    console.log(`[identity] wrote ${kept.size} file(s) to ${positional[1]}`);
    return 0;
  }

  if (command === "compare" || command === "diff") {
    if (positional.length !== 2) throw new Error(`${command} needs two paths`);
    let expected;
    let actualRoot;
    if (command === "compare") {
      actualRoot = requireDirectory(positional[0]);
      expected = readManifest(positional[1]);
    } else {
      expected = walk(requireDirectory(positional[0]));
      actualRoot = requireDirectory(positional[1]);
    }
    const differences = compare(expected, walk(actualRoot), ignores);
    if (differences.length === 0) {
      console.log(`[identity] ${expected.size} file(s), identical`);
      return 0;
    }
    printDifferences(differences, maximum);
    return 1;
  }

  throw new Error(`unknown command: ${command}`);
};

try {
  process.exit(main());
} catch (error) {
  console.error(`[identity] ${error.message}`);
  usage();
  process.exit(2);
}
