#!/usr/bin/env node
// Keep installation examples aligned with this checkout's version and pins.
// Publication status and historical qualification records are not generated.
import { readFileSync, writeFileSync } from "node:fs";
import { join, resolve } from "node:path";
import { fileURLToPath, pathToFileURL } from "node:url";

const repository = "https://github.com/0x000000000000000000001/gopurs";
const defaultRoot = fileURLToPath(new URL("../", import.meta.url));
const fence = (language, lines) => "```" + language + "\n" + lines.join("\n") + "\n```";

function metadata(root) {
  const read = (file) => JSON.parse(readFileSync(join(root, file), "utf8"));
  const { version } = read("package.json");
  if (typeof version !== "string" || !/^\d+\.\d+\.\d+(?:-rc\.\d+)?$/.test(version))
    throw new Error("package.json: unsupported documentation version");
  const sources = read("bootstrap-sources.json");
  for (const [name, pin] of [
    ["compiler", sources.compiler],
    ...["prelude", "console", "arrays", "strings"].map((name) => [name, sources.libraries?.[name]]),
  ]) {
    if (
      !pin ||
      !/^[a-f0-9]{40}$/.test(pin.rev) ||
      !/^https:\/\/github\.com\/[A-Za-z0-9_.-]+\/[A-Za-z0-9_.-]+(?:\.git)?$/.test(pin.url)
    )
      throw new Error(`bootstrap-sources.json: invalid ${name} source pin`);
  }
  return { version, tag: `v${version}`, sources };
}

function blocks({ version, tag, sources }) {
  const npm = `npm install --save-dev --ignore-scripts gopurs@${version}`;
  const nix = `nix develop github:0x000000000000000000001/gopurs/${tag}`;
  const sourceVersion = `For installation, start with the [latest published release](${repository}/releases/latest)\nand the guides at its tag. This guide targets the version in this checkout:\n\`gopurs@${version}\`.`;
  return {
    "README.md": {
      "source-version": sourceVersion,
      install: fence("bash", [
        npm,
        "spago build                 # Compile PureScript with the typed fork",
        "npx gopurs --main Main      # Generate Go source and go.mod in output/",
        "(cd output && go build -o app ./main && ./app)",
      ]),
    },
    "docs/quickstart.md": {
      "source-version": sourceVersion,
      toolchain: fence("bash", [nix]),
      frontend: `- The [pinned typed \`purs\` compiler](${sources.compiler.url.replace(/\.git$/, "")}/tree/${sources.compiler.rev}). Spago uses it to emit the type information gopurs needs.`,
      install: fence("bash", [npm]),
      console: fence("bash", [
        "cd ..",
        `git clone ${sources.libraries.console.url}`,
        `git -C gopurs-console checkout --detach ${sources.libraries.console.rev}`,
        "cd hello",
      ]),
    },
    "docs/migration.md": {
      "source-version": sourceVersion,
      toolchain: fence("bash", [nix, "command -v purs", "purs --version", npm]),
      libraries:
        `These revisions come from the [source manifest for \`${tag}\`](${repository}/blob/${tag}/bootstrap-sources.json).\n\n` +
        fence("yaml", [
          "  extraPackages:",
          ...["prelude", "console", "arrays", "strings"].flatMap((name) => [
            `    ${name}:`,
            `      git: ${sources.libraries[name].url}`,
            `      ref: ${sources.libraries[name].rev}`,
          ]),
        ]),
    },
    "docs/publication-scope.md": {
      identity: [
        "| Item | Source of truth |",
        "| --- | --- |",
        "| Package | `gopurs` |",
        `| Documentation target | \`${version}\`, the version in this checkout |`,
        `| Published release | [Latest GitHub release](${repository}/releases/latest) |`,
        "| npm channels | [Current registry state](https://www.npmjs.com/package/gopurs?activeTab=versions) |",
        "| Source version | [`package.json`](../package.json) |",
        "| Versioning | [Semantic Versioning 2.0.0](https://semver.org/) |",
        `| Version tag | \`${tag}\`; release tags are annotated and signed |`,
      ].join("\n"),
    },
  };
}

function render(file, source, replacements) {
  const marker = /<!-- gopurs:([a-z-]+):(start|end) -->/g;
  const expected = Object.keys(replacements).flatMap((name) => [`${name}:start`, `${name}:end`]);
  const found = [...source.matchAll(marker)].map((match) => `${match[1]}:${match[2]}`);
  if (
    found.length !== expected.length ||
    new Set(found).size !== found.length ||
    expected.some((value) => !found.includes(value))
  )
    throw new Error(`${file}: missing, duplicate or unknown generated documentation markers`);
  let result = source;
  for (const [name, content] of Object.entries(replacements)) {
    const start = `<!-- gopurs:${name}:start -->`;
    const end = `<!-- gopurs:${name}:end -->`;
    const from = result.indexOf(start);
    const to = result.indexOf(end);
    if (to < from || result.slice(from + start.length, to).includes("<!-- gopurs:"))
      throw new Error(`${file}: reversed or nested generated documentation markers for ${name}`);
    result = result.slice(0, from) + `${start}\n${content}\n${end}` + result.slice(to + end.length);
  }
  return result;
}

function main(args) {
  let root = defaultRoot,
    mode = "check",
    selected = false;
  for (let index = 0; index < args.length; index++) {
    const arg = args[index];
    if (arg === "--help" || arg === "-h") {
      console.log(
        "Usage: node tools/sync-docs.mjs [--check | --write] [--root DIR]\nChecks generated installation examples against package.json and bootstrap-sources.json.\nThe default --check is read-only. --write updates only marked blocks; historical prose is preserved.",
      );
      return;
    }
    if (arg === "--root" && args[index + 1]) root = resolve(args[++index]);
    else if (["--check", "--write"].includes(arg) && !selected) {
      mode = arg.slice(2);
      selected = true;
    } else throw new Error(`Unknown, duplicate or incomplete option: ${arg}`);
  }
  // Validate every document before writing any of them.
  const changes = Object.entries(blocks(metadata(root))).flatMap(([file, replacements]) => {
    const path = join(root, file);
    const source = readFileSync(path, "utf8");
    const generated = render(file, source, replacements);
    return source === generated ? [] : [{ file, path, generated }];
  });
  if (mode === "check" && changes.length) {
    for (const { file } of changes) console.error(`[docs] out of date: ${file}`);
    throw new Error("Run npm run docs:sync and commit the updated documentation");
  }
  if (mode === "write") for (const { path, generated } of changes) writeFileSync(path, generated);
  console.log(
    `[docs] version and source pins ${mode === "write" ? `synchronized (${changes.length} changed)` : "match"}`,
  );
}

if (process.argv[1] && import.meta.url === pathToFileURL(resolve(process.argv[1])).href) {
  try {
    main(process.argv.slice(2));
  } catch (error) {
    console.error(`[docs] ${error.message}`);
    process.exitCode = 1;
  }
}
