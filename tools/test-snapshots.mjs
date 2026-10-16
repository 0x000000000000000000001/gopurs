import { copyFileSync, existsSync, mkdirSync, readFileSync, readdirSync, writeFileSync } from "node:fs";
import { dirname, join } from "node:path";

// Match comments as Go tokens, skipping raw/interpreted strings and runes.
// Source comments using this reserved namespace keep the fragmented layout.
export function packageSourceSections(content) {
  const tokens = /\/\/[^\n]*|\/\*[\s\S]*?\*\/|"(?:\\[\s\S]|[^"\\])*"|'(?:\\[\s\S]|[^'\\])*'|`[^`]*`/g;
  const sections = [];
  let active, header;
  for (const match of content.matchAll(tokens)) {
    if (match.index > 0 && content[match.index - 1] !== "\n") continue;
    const start = /^\/\/ gopurs-source-v1: ((?:[0-9a-f]{2})+)$/.exec(match[0]);
    if (start) {
      if (active) throw new Error("Nested package source section");
      header ??= content.slice(0, match.index);
      active = { name: Buffer.from(start[1], "hex").toString("utf8"), start: match.index };
    } else if (match[0] === "// gopurs-source-end-v1") {
      if (!active) throw new Error("Package source section has no start");
      const end = match.index + match[0].length + (content[match.index + match[0].length] === "\n" ? 1 : 0);
      sections.push({ name: active.name, content: content.slice(active.start, end) });
      active = undefined;
    }
  }
  if (active) throw new Error("Unterminated package source section");
  return { header, sections };
}

export function snapshotFiles(root, fixture, env = process.env) {
  if (env.GOPURS_GO_PACKAGES === "1") {
    const output = join(fixture.directory, "output");
    const sources = ["main/main.go", "purescript/init_all/init.go", "purescript/init_external/init.go"];
    let fragments = 0;
    const projections = {};
    for (const entry of readdirSync(join(output, "purescript"), { withFileTypes: true })) {
      if (!entry.isDirectory() || !/^p_(?:root|[0-9a-f]{32})$/.test(entry.name)) continue;
      const bundle = `purescript/${entry.name}/0_gopurs_package.go`;
      if (existsSync(join(output, bundle))) {
        const selected = packageSourceSections(readFileSync(join(output, bundle), "utf8")).sections.filter(section =>
          section.name === "Main.go" || (fixture.snapshotFfi && section.name === "Main_ffi.go"));
        if (selected.length === 0) continue;
        fragments += selected.filter(section => section.name === "Main.go").length;
        projections[bundle] = selected.map(section => section.name);
        sources.push(bundle, `purescript/${entry.name}/zz_gopurs_order.go`);
        continue;
      }
      const files = readdirSync(join(output, "purescript", entry.name)).filter(name =>
        name === "Main.go" || (fixture.snapshotFfi && name === "Main_ffi.go"));
      if (files.length === 0) continue;
      fragments += files.filter(name => name === "Main.go").length;
      sources.push(...[...files, "zz_gopurs_order.go"].map(name => `purescript/${entry.name}/${name}`));
    }
    if (fragments === 0) throw new Error("Missing generated package fragments for Main");
    return [{ generated: join(fixture.directory, "logs/package-snapshot.txt"),
      expected: join(root, `tests/passing-package-snapshots/${fixture.name}.txt`), sources: sources.sort(), projections }];
  }
  return (fixture.snapshotFfi ? ["", "_ffi"] : [""]).map(suffix => ({
    generated: join(fixture.directory, `output/purescript/Main${suffix}.go`),
    expected: join(root, `tests/passing-snapshots/${fixture.name}${suffix}.go`),
  }));
}

export async function formatSnapshots(files, processes, fixture, env) {
  const output = join(fixture.directory, "output");
  const sources = files.flatMap(file => file.sources?.map(name => join(output, name)) ?? [file.generated]);
  for (const source of sources) if (!existsSync(source)) throw new Error(`Missing generated snapshot source: ${source}`);
  await processes.run("format", "gofmt", ["-w", ...sources], {
    cwd: fixture.directory, env, log: join(fixture.directory, "logs/format.log"),
  });
  for (const file of files) if (file.sources) {
    writeFileSync(file.generated, file.sources.map(name => {
      let content = readFileSync(join(output, name), "utf8");
      if (file.projections?.[name]) {
        const { header, sections } = packageSourceSections(content);
        content = header + sections.filter(section => file.projections[name].includes(section.name)).map(section => section.content).join("\n");
      }
      return `// gopurs-package-file: ${name}\n${content}\n`;
    }).join(""));
  }
}

export async function verifySnapshots(files, processes, fixture, env) {
  for (const { generated, expected } of files) {
    if (!existsSync(expected)) throw new Error(`Missing snapshot: ${expected}. Use --update-snapshots to create it explicitly.`);
    if (!readFileSync(expected).equals(readFileSync(generated))) {
      await processes.run("snapshot diff", "diff", ["-u", expected, generated], { cwd: fixture.directory, env, log: join(fixture.directory, "logs/snapshot.diff"), display: true });
    }
  }
}

export function updateSnapshots(files) {
  // Called only after the fixture's Go build and execution have succeeded.
  for (const { generated, expected } of files) {
    mkdirSync(dirname(expected), { recursive: true });
    copyFileSync(generated, expected);
    console.log(`   Updated snapshot: ${expected}`);
  }
}
