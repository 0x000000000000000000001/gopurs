#!/usr/bin/env node
// Check the normative documentation without a network or Markdown dependency:
// relative links and anchors, package.json.files entries, and module-map
// coverage against the real src/ tree. Historical records are out of scope.
import { existsSync, readdirSync, readFileSync, statSync } from "node:fs";
import { basename, dirname, extname, join, relative, resolve, sep } from "node:path";
import { fileURLToPath, pathToFileURL } from "node:url";

const ROOT = fileURLToPath(new URL("../", import.meta.url));
const EXTERNAL = /^(?:[a-z][a-z0-9+.-]*:|\/\/)/i;

const usage = () => {
  console.error(`Usage: node tools/check-docs.mjs [options] [files...]

Checks relative links and heading anchors, every package.json "files" entry,
and that docs/module-map.md has exactly one row per src/**/*.purs module.

Options:
  --root <dir>     Repository root to inspect (default: this checkout).
  --links-only     Skip package.json.files and module-map coverage.
  --quiet          Print only failures.
  -h, --help       Show this help.

With no files, checks README.md, CONTRIBUTING.md, cleanup.md, todo.md and
docs/**/*.md except docs/history/, docs/build-*, docs/plan-archive-* and
docs/testing.md (historical sections are split out by the cleanup plan).`);
};

const toPosix = path => path.split(sep).join("/");
const isHistorical = path => {
  const posix = toPosix(path);
  const name = basename(posix);
  return posix.startsWith("docs/history/")
    || name.startsWith("build-")
    || name.startsWith("plan-archive-")
    || name === "testing.md";
};

const walk = (directory, predicate) => {
  const found = [];
  for (const entry of readdirSync(directory, { withFileTypes: true }).sort((a, b) => a.name.localeCompare(b.name))) {
    const path = join(directory, entry.name);
    if (entry.isDirectory()) found.push(...walk(path, predicate));
    else if (entry.isFile() && predicate(path)) found.push(path);
  }
  return found;
};

const defaultDocuments = root => {
  const candidates = ["README.md", "CONTRIBUTING.md", "cleanup.md", "todo.md"]
    .map(name => join(root, name))
    .filter(existsSync);
  const docs = join(root, "docs");
  if (existsSync(docs)) {
    candidates.push(...walk(docs, path => extname(path) === ".md"));
  }
  return candidates
    .filter(path => !isHistorical(toPosix(relative(root, path))))
    .map(path => toPosix(relative(root, path)))
    .sort();
};

// GitHub-style heading anchor: keep Unicode letters and digits, drop the rest
// of the punctuation, turn spaces into hyphens, and suffix duplicates.
export const slugify = text => text
  .trim()
  .toLowerCase()
  .replace(/<[^>]*>/g, "")
  .replace(/!\[([^\]]*)\]\([^)]*\)/g, "$1")
  .replace(/\[([^\]]*)\]\([^)]*\)/g, "$1")
  .replace(/[`*~]/g, "")
  .replace(/[^\p{L}\p{N}\s_-]/gu, "")
  .replace(/\s+/g, "-");

// Strip fenced blocks and inline code so examples never look like links.
export const stripCode = source => {
  const lines = source.split("\n");
  let fence = null;
  const kept = [];
  for (const line of lines) {
    const match = line.match(/^\s*(`{3,}|~{3,})/);
    if (match) {
      if (fence === null) fence = match[1][0] + match[1].length;
      else if (match[1][0] + match[1].length === fence) fence = null;
      kept.push("");
      continue;
    }
    if (fence !== null) { kept.push(""); continue; }
    kept.push(line.replace(/(`+)[^`]*?\1/g, ""));
  }
  return kept.join("\n");
};

export const anchorsOf = source => {
  const counts = new Map();
  const anchors = new Set();
  for (const line of stripCode(source).split("\n")) {
    const match = line.match(/^(#{1,6})\s+(.+?)\s*#*\s*$/);
    if (!match) continue;
    const base = slugify(match[2]);
    const count = counts.get(base) ?? 0;
    counts.set(base, count + 1);
    anchors.add(count === 0 ? base : `${base}-${count}`);
  }
  return anchors;
};

export const linksOf = source => {
  const links = [];
  for (const match of stripCode(source).matchAll(/!?\[[^\]]*\]\(\s*([^)\s]+)(?:\s+"[^"]*")?\s*\)/g)) {
    links.push(match[1]);
  }
  return links;
};

const readAnchors = (root, path) => anchorsOf(readFileSync(join(root, path), "utf8"));

export const checkLinks = (root, documents) => {
  const errors = [];
  const checked = [];
  let outside = 0;
  for (const document of documents) {
    const absolute = join(root, document);
    if (!existsSync(absolute)) {
      errors.push(`${document}: document does not exist`);
      continue;
    }
    const source = readFileSync(absolute, "utf8");
    checked.push(document);
    for (const link of linksOf(source)) {
      if (!link || EXTERNAL.test(link)) continue;
      let target = link;
      let anchor = "";
      const hash = target.indexOf("#");
      if (hash >= 0) { anchor = target.slice(hash + 1); target = target.slice(0, hash); }
      if (target.includes("?")) target = target.slice(0, target.indexOf("?"));
      let decoded;
      try { decoded = decodeURIComponent(target); }
      catch { errors.push(`${document}: invalid URL escape in ${link}`); continue; }
      let targetPath;
      if (decoded === "") targetPath = document;
      else if (decoded.startsWith("/")) targetPath = decoded.slice(1);
      else targetPath = toPosix(relative(root, resolve(dirname(join(root, document)), decoded)));
      if (targetPath.startsWith("../")) {
        // Cross-repository references cannot be verified from a fresh checkout.
        outside += 1;
        continue;
      }
      const absoluteTarget = join(root, targetPath);
      if (!existsSync(absoluteTarget)) {
        errors.push(`${document}: missing target ${link} (resolved to ${targetPath})`);
        continue;
      }
      if (!anchor || statSync(absoluteTarget).isDirectory()) continue;
      if (extname(absoluteTarget) !== ".md") {
        errors.push(`${document}: anchor ${link} points to a non-Markdown file`);
        continue;
      }
      let decodedAnchor;
      try { decodedAnchor = decodeURIComponent(anchor); }
      catch { errors.push(`${document}: invalid anchor escape in ${link}`); continue; }
      const anchors = readAnchors(root, targetPath);
      if (!anchors.has(decodedAnchor)) {
        errors.push(`${document}: missing anchor ${link} in ${targetPath}`);
      }
    }
  }
  return { errors, checked, outside };
};

const checkPackageFiles = root => {
  const errors = [];
  let manifest;
  try {
    manifest = JSON.parse(readFileSync(join(root, "package.json"), "utf8"));
  } catch (error) {
    return { errors: [`package.json: ${error.message}`], checked: 0 };
  }
  const files = Array.isArray(manifest.files) ? manifest.files : [];
  if (files.length === 0) errors.push("package.json: files is missing or empty");
  for (const entry of files) {
    if (typeof entry !== "string" || entry.length === 0) {
      errors.push(`package.json: invalid files entry ${JSON.stringify(entry)}`);
    } else if (!existsSync(join(root, entry))) {
      errors.push(`package.json: files entry does not exist: ${entry}`);
    }
  }
  return { errors, checked: files.length };
};

const readModuleMap = root => {
  const path = join(root, "docs/module-map.md");
  if (!existsSync(path)) return { rows: [], errors: ["docs/module-map.md: missing"] };
  const rows = [];
  const errors = [];
  for (const [index, line] of readFileSync(path, "utf8").split("\n").entries()) {
    const match = line.match(/^\| `([^`]+)`<br>`([^`]+)` \| `([^`]+)` \| ([A-Za-zÀ-ÿ]+) \|/);
    if (!match) continue;
    rows.push({ name: match[1], path: match[2], target: match[3], phase: match[4], line: index + 1 });
  }
  if (rows.length === 0) errors.push("docs/module-map.md: no module rows found");
  return { rows, errors };
};

const sourceModules = root => {
  const src = join(root, "src");
  if (!existsSync(src)) return new Map();
  const modules = new Map();
  for (const path of walk(src, file => extname(file) === ".purs")) {
    const source = readFileSync(path, "utf8");
    const match = source.match(/^module\s+([\w.]+)/m);
    if (!match) {
      modules.set(`?${toPosix(relative(root, path))}`, toPosix(relative(root, path)));
      continue;
    }
    modules.set(match[1], toPosix(relative(root, path)));
  }
  return modules;
};

export const checkModuleMap = root => {
  const errors = [];
  const { rows, errors: mapErrors } = readModuleMap(root);
  errors.push(...mapErrors);
  const modules = sourceModules(root);
  const seen = new Set();
  for (const row of rows) {
    if (seen.has(row.name)) errors.push(`docs/module-map.md:${row.line}: duplicate module ${row.name}`);
    seen.add(row.name);
    if (!modules.has(row.name)) {
      errors.push(`docs/module-map.md:${row.line}: ${row.name} is not a module under src/`);
      continue;
    }
    const actual = modules.get(row.name);
    if (actual !== row.path) {
      errors.push(`docs/module-map.md:${row.line}: ${row.name} declares ${row.path}, actual file is ${actual}`);
    }
  }
  for (const [name, path] of modules) {
    if (!seen.has(name)) errors.push(`docs/module-map.md: missing row for ${name} (${path})`);
  }
  const expected = modules.size;
  if (rows.length !== expected) {
    errors.push(`docs/module-map.md: ${rows.length} rows for ${expected} modules`);
  }
  return { errors, checked: rows.length };
};

export const checkRepository = (root, documents) => {
  const linkResult = checkLinks(root, documents);
  const packageResult = checkPackageFiles(root);
  const moduleResult = checkModuleMap(root);
  return {
    errors: [...linkResult.errors, ...packageResult.errors, ...moduleResult.errors],
    counts: {
      documents: linkResult.checked.length,
      files: packageResult.checked,
      modules: moduleResult.checked,
      outside: linkResult.outside,
    },
  };
};

const parseArgs = argv => {
  const options = { root: ROOT, linksOnly: false, quiet: false, help: false, files: [] };
  for (let index = 0; index < argv.length; index += 1) {
    const argument = argv[index];
    if (argument === "--root") {
      if (!argv[index + 1]) throw new Error("--root needs a directory");
      options.root = resolve(argv[++index]);
    } else if (argument === "--links-only") options.linksOnly = true;
    else if (argument === "--quiet") options.quiet = true;
    else if (argument === "-h" || argument === "--help") options.help = true;
    else if (argument.startsWith("-")) throw new Error(`unknown option ${argument}`);
    else options.files.push(argument);
  }
  return options;
};

const main = () => {
  let options;
  try { options = parseArgs(process.argv.slice(2)); }
  catch (error) { console.error(`[docs] ${error.message}`); usage(); process.exitCode = 2; return; }
  if (options.help) { usage(); return; }
  if (!existsSync(options.root)) { console.error(`[docs] root does not exist: ${options.root}`); process.exitCode = 2; return; }
  const documents = options.files.length ? options.files : defaultDocuments(options.root);
  if (documents.length === 0) { console.error("[docs] no documents selected"); process.exitCode = 2; return; }
  const result = options.linksOnly
    ? { errors: checkLinks(options.root, documents).errors, counts: { documents: documents.length } }
    : checkRepository(options.root, documents);
  if (result.errors.length > 0) {
    for (const error of result.errors) console.error(`[docs] ${error}`);
    console.error(`[docs] ${result.errors.length} error(s)`);
    process.exitCode = 1;
    return;
  }
  if (!options.quiet) {
    console.log(`[docs] ${result.counts.documents} document(s) checked`);
    if (!options.linksOnly) {
      console.log(`[docs] ${result.counts.files} packaged file(s) present`);
      console.log(`[docs] ${result.counts.modules} module row(s) match src/`);
      if (result.counts.outside) console.log(`[docs] ${result.counts.outside} cross-repository link(s) skipped`);
    }
  }
};

if (process.argv[1] && import.meta.url === pathToFileURL(resolve(process.argv[1])).href) main();
