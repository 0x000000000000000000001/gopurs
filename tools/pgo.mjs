// Production profile-guided optimization for the Rust-hosted gopurs compiler.
//
// The training workload is compiler self-compilation: a separate workspace
// builds this repository's src/ tree against the Go library checkouts, and the
// resulting CoreFn modules are frozen with their sibling FFI sources and
// relative modulePath values. The unprofiled bootstrap binary defines the
// oracle Go output; the instrumented binary must reproduce it byte for byte in
// every training pass before its profile is accepted. Test.Main and the
// held-out Aff benchmark fixtures are never part of the workload.
import assert from "node:assert/strict";
import { createHash } from "node:crypto";
import {
  accessSync, constants, copyFileSync, existsSync, mkdirSync, readFileSync,
  readdirSync, rmSync, statSync, symlinkSync, writeFileSync,
} from "node:fs";
import { dirname, join, relative, resolve, sep } from "node:path";
import { nativeWorkspaceConfig, verifyTypedOutput } from "./native-workspace.mjs";

export const PGO_TRAINING_PASSES = 3;
// The excluded entry is the fixture/benchmark test module, never compiler code.
export const PGO_EXCLUDED_MODULES = Object.freeze(["Test.Main"]);
// Requiring Main proves every frozen workload can produce an executable entry
// point; the nightly experiment lost it to an APFS case alias.
export const PGO_REQUIRED_MODULES = Object.freeze(["Main"]);
// Full parallelism for the first two passes, sequential prepare/PBO/emit for
// the final pass, matching the isolated night-pgo experiment.
export const PGO_JOBS = Object.freeze({
  GOPURS_JOBS: "8", GOPURS_PREPARE_JOBS: "8", GOPURS_PBO_JOBS: "8", GOPURS_EMIT_JOBS: "8", GOPURS_PIPELINE: "1",
});
export const PGO_FINAL_JOBS = Object.freeze({
  ...PGO_JOBS, GOPURS_PREPARE_JOBS: "1", GOPURS_PBO_JOBS: "1", GOPURS_EMIT_JOBS: "1",
});
const PGO_TRAINING_POLICY = Object.freeze({
  scope: "compiler self-compilation with the Go-runtime library checkouts",
  library_overlap: "library modules shared with the Go library checkouts are allowed",
  held_out: "no Aff benchmark fixtures are read; Test.Main is excluded from training",
});
const TRAINING_SOURCE_EXTENSIONS = Object.freeze(["purs", "go", "js"]);

export const sha256 = bytes => createHash("sha256").update(bytes).digest("hex");

// Deterministic, relative-path, content-addressed file inventory.
export function walk(directory) {
  if (!existsSync(directory)) return [];
  return readdirSync(directory, { withFileTypes: true })
    .sort((left, right) => left.name < right.name ? -1 : left.name > right.name ? 1 : 0)
    .flatMap(entry => {
      const path = join(directory, entry.name);
      return entry.isDirectory() ? walk(path) : entry.isFile() ? [path] : [];
    });
}

export function manifest(directory, filter = () => true) {
  return walk(directory).filter(filter).map(path => ({
    path: relative(directory, path).split(sep).join("/"),
    bytes: statSync(path).size,
    sha256: sha256(readFileSync(path)),
  }));
}

export function digestManifest(files) {
  return sha256(JSON.stringify(files));
}

// The compiler writes output/main/main.go plus one workspace Go tree; go.sum
// only appears if a Go tool ran in the training directory. CoreFn is an input,
// never part of the generated set.
export function emittedManifest(output) {
  return manifest(output, path => path.endsWith(".go") || path.split("/").at(-1) === "go.mod");
}

// Freeze the typed compiler-self output as a self-contained training project.
// Every corefn.json gets its sibling .purs/.go/.js sources and a relative
// modulePath, so the workload does not depend on the live checkouts. Excluded
// modules are recorded but never copied.
export function freezeTrainingOutput({ output, projectRoot, trainingDirectory }) {
  if (!existsSync(output)) throw new Error(`Missing typed output: ${output}`);
  const frozenOutput = join(trainingDirectory, "output");
  mkdirSync(frozenOutput, { recursive: true });
  const modules = [];
  const excluded = [];
  for (const entry of readdirSync(output, { withFileTypes: true })
    .sort((left, right) => left.name < right.name ? -1 : left.name > right.name ? 1 : 0)) {
    if (!entry.isDirectory()) continue;
    const corefn = join(output, entry.name, "corefn.json");
    if (!existsSync(corefn)) continue;
    const module = entry.name;
    if (PGO_EXCLUDED_MODULES.includes(module)) { excluded.push(module); continue; }
    const bytes = readFileSync(corefn);
    const metadata = JSON.parse(bytes);
    if (typeof metadata.modulePath !== "string" || metadata.modulePath === "") {
      throw new Error(`${corefn} has no modulePath`);
    }
    const original = resolve(projectRoot, metadata.modulePath);
    if (!existsSync(original)) throw new Error(`${corefn} points to missing source ${original}`);
    const source = metadata.modulePath;
    const frozenSource = `sources/${module}/${module.split(".").at(-1)}.purs`;
    const sources = [];
    for (const extension of TRAINING_SOURCE_EXTENSIONS) {
      const sibling = extension === "purs" ? original : original.replace(/\.purs$/, "." + extension);
      if (!existsSync(sibling)) continue;
      const relativeSibling = frozenSource.replace(/\.purs$/, "." + extension);
      const destination = join(trainingDirectory, relativeSibling);
      mkdirSync(dirname(destination), { recursive: true });
      copyFileSync(sibling, destination, constants.COPYFILE_FICLONE);
      sources.push(relativeSibling);
    }
    metadata.modulePath = frozenSource;
    const destination = join(frozenOutput, module, "corefn.json");
    mkdirSync(dirname(destination), { recursive: true });
    writeFileSync(destination, JSON.stringify(metadata, null, 2) + "\n");
    modules.push({
      module, source, corefn: relative(projectRoot, corefn).split(sep).join("/"),
      frozen_source: frozenSource, sources, sha256: sha256(bytes),
    });
  }
  for (const module of PGO_REQUIRED_MODULES) {
    if (!modules.some(record => record.module === module)) {
      throw new Error(`Training output is missing required module ${module}`);
    }
  }
  const frozen = manifest(trainingDirectory);
  return { modules, excluded, required: [...PGO_REQUIRED_MODULES], manifest: frozen, sha256: digestManifest(frozen) };
}

// Generated Go code lands in output/main on case-sensitive filesystems and
// inside output/Main on case-insensitive APFS (main aliases Main). Delete
// individual generated files, never module directories, so the frozen CoreFn
// and the FFI sources survive either aliasing. Only the compiler's generated
// Go output is removed; any other diff fails the manifest assertion instead of
// being cleaned away.
export function cleanGenerated(trainingDirectory) {
  const output = join(trainingDirectory, "output");
  let removed = 0;
  for (const path of walk(output)) {
    const base = path.split(sep).at(-1);
    if (!path.endsWith(".go") && base !== "go.mod" && base !== "go.sum") continue;
    rmSync(path);
    removed++;
  }
  // PBO metadata and caches are recreated per process, as in the benchmark.
  for (const name of [".cache", ".purmeta"]) {
    rmSync(join(trainingDirectory, name), { recursive: true, force: true });
  }
  return removed;
}

export function assertTrainingIntact(trainingDirectory, expectedManifest) {
  for (const module of PGO_REQUIRED_MODULES) {
    const corefn = join(trainingDirectory, "output", module, "corefn.json");
    if (!existsSync(corefn)) throw new Error(`Training input ${corefn} was removed`);
  }
  assert.deepEqual(manifest(trainingDirectory), expectedManifest);
}

// CARGO_ENCODED_RUSTFLAGS is unit-separated; paths keep their spaces in one
// argv entry and never pass through a shell.
export function encodedRustFlags(flags) {
  if (!Array.isArray(flags) || !flags.length) throw new TypeError("Expected at least one Rust flag");
  for (const flag of flags) {
    if (typeof flag !== "string" || flag.length === 0) throw new TypeError("Rust flags must be non-empty strings");
    if (flag.includes("\x1f")) throw new Error("Rust flags cannot contain the CARGO_ENCODED_RUSTFLAGS separator");
  }
  return flags.join("\x1f");
}

export function rustFlagsEnvironment(environment, flags = []) {
  const env = { ...environment };
  if (flags.length) env.CARGO_ENCODED_RUSTFLAGS = encodedRustFlags(flags);
  return env;
}

// Drive the whole PGO stage from the bootstrap binary. The caller owns command
// execution through `run` (CommandRunner logging and interruption) and the
// cargo/linker policy through `buildBinary(label, targetDirectory, rustflags)`.
export async function buildProfileGuidedBinary({
  run, root, compiler, workspace, rust, bootstrapBinary,
  environment, typedEnvironment, targetRoot, buildBinary, toolchain,
}) {
  const pgoRoot = join(workspace, "pgo");
  const trainingWorkspace = join(pgoRoot, "workspace");
  const trainingDirectory = join(pgoRoot, "training");
  const rawDirectory = join(pgoRoot, "raw");
  const mergedProfile = join(pgoRoot, "training.profdata");
  const metadataPath = join(workspace, "pgo-profile.json");
  const sourceManifest = manifest(rust, path => path.endsWith(".rs") || path.endsWith(".toml"));
  const metadata = {
    status: "pending",
    enabled: true,
    started_at: new Date().toISOString(),
    source: {
      directory: relative(workspace, rust),
      files: sourceManifest.length, sha256: digestManifest(sourceManifest), manifest: sourceManifest,
    },
    training: {
      workspace: relative(workspace, trainingWorkspace),
      output: relative(workspace, join(trainingWorkspace, "output")),
      compiler,
      required_modules: [...PGO_REQUIRED_MODULES],
      excluded_modules: [...PGO_EXCLUDED_MODULES],
      policy: PGO_TRAINING_POLICY,
    },
    builds: {
      bootstrap: { binary: relative(workspace, bootstrapBinary), sha256: sha256(readFileSync(bootstrapBinary)) },
    },
    profile: null,
    passes: [],
    binary: null,
  };
  let stage = "training-workspace";
  const save = () => writeFileSync(metadataPath, JSON.stringify(metadata, null, 2) + "\n");
  save();
  try {
    mkdirSync(trainingWorkspace, { recursive: true });
    writeFileSync(join(trainingWorkspace, "spago.yaml"), nativeWorkspaceConfig(root, { runtime: "go" }));
    symlinkSync(join(root, "src"), join(trainingWorkspace, "src"), "dir");
    await run("pgo-typed-corefn", "spago", ["build"], trainingWorkspace, typedEnvironment);
    stage = "freeze-training-inputs";
    const output = join(trainingWorkspace, "output");
    metadata.training.verified = verifyTypedOutput(output, compiler);
    const frozen = freezeTrainingOutput({ output, projectRoot: trainingWorkspace, trainingDirectory });
    metadata.training.modules = frozen.modules;
    metadata.training.excluded_present = frozen.excluded;
    metadata.training.manifest = frozen.manifest;
    metadata.training.sha256 = frozen.sha256;
    metadata.training.files = frozen.manifest.length;
    save();

    stage = "training-oracle";
    cleanGenerated(trainingDirectory);
    assertTrainingIntact(trainingDirectory, frozen.manifest);
    await run("pgo-oracle", bootstrapBinary, ["--main", "Main"], trainingDirectory, { ...environment, ...PGO_JOBS });
    const oracle = emittedManifest(join(trainingDirectory, "output"));
    if (!oracle.length) throw new Error("The PGO oracle produced no generated Go files");
    metadata.training.oracle = { files: oracle.length, sha256: digestManifest(oracle), manifest: oracle };
    save();
    console.log(`[pgo] oracle: ${oracle.length} generated files from the unprofiled compiler`);

    stage = "instrument";
    mkdirSync(rawDirectory, { recursive: true });
    const instrumented = await buildBinary("pgo-instrument", join(targetRoot, "pgo-generate"),
      ["-C", `profile-generate=${rawDirectory}`]);
    metadata.builds.instrumented = { binary: relative(workspace, instrumented), sha256: sha256(readFileSync(instrumented)) };
    save();

    stage = "training";
    for (let round = 1; round <= PGO_TRAINING_PASSES; round++) {
      const jobs = round === PGO_TRAINING_PASSES ? PGO_FINAL_JOBS : PGO_JOBS;
      cleanGenerated(trainingDirectory);
      assertTrainingIntact(trainingDirectory, frozen.manifest);
      await run(`pgo-train-${round}`, instrumented, ["--main", "Main"], trainingDirectory, {
        ...environment, ...jobs, LLVM_PROFILE_FILE: join(rawDirectory, `round-${round}-%m-%p.profraw`),
      });
      const emitted = emittedManifest(join(trainingDirectory, "output"));
      assert.deepEqual(emitted, oracle, `PGO training pass ${round} generated different Go sources`);
      cleanGenerated(trainingDirectory);
      assertTrainingIntact(trainingDirectory, frozen.manifest);
      metadata.passes.push({ round, jobs, files: emitted.length, sha256: digestManifest(emitted) });
      save();
      console.log(`[pgo] training pass ${round}/${PGO_TRAINING_PASSES}: ${emitted.length} files identical to the oracle`);
    }

    stage = "merge";
    const profdata = toolchain.profdata;
    accessSync(profdata, constants.X_OK);
    const rawProfiles = manifest(rawDirectory, path => path.endsWith(".profraw"));
    if (!rawProfiles.length) throw new Error("PGO training produced no raw profiles");
    await run("pgo-profdata", profdata, ["--version"], pgoRoot, environment);
    await run("pgo-merge", profdata,
      ["merge", "-o", mergedProfile, ...rawProfiles.map(file => join(rawDirectory, file.path))], pgoRoot, environment);
    metadata.profile = {
      tool: { path: profdata, sha256: sha256(readFileSync(profdata)) },
      generate_flags: ["-C", `profile-generate=${rawDirectory}`],
      raw: rawProfiles, raw_sha256: digestManifest(rawProfiles),
      merged: {
        path: relative(workspace, mergedProfile), bytes: statSync(mergedProfile).size,
        sha256: sha256(readFileSync(mergedProfile)),
      },
      use_flags: ["-C", `profile-use=${mergedProfile}`],
    };
    save();
    console.log(`[pgo] merged ${rawProfiles.length} raw profiles`);

    stage = "profile-use";
    const binary = await buildBinary("pgo-profile-use", join(targetRoot, "pgo-use"),
      ["-C", `profile-use=${mergedProfile}`]);
    metadata.binary = {
      path: relative(workspace, binary), bytes: statSync(binary).size, sha256: sha256(readFileSync(binary)),
      bootstrap_sha256: metadata.builds.bootstrap.sha256, instrumented_sha256: metadata.builds.instrumented.sha256,
    };
    assert.deepEqual(manifest(rust, path => path.endsWith(".rs") || path.endsWith(".toml")), sourceManifest);
    assertTrainingIntact(trainingDirectory, frozen.manifest);
    metadata.status = "passed";
    metadata.finished_at = new Date().toISOString();
    save();
    console.log(`[pgo] source sha256: ${metadata.source.sha256} (${metadata.source.manifest.length} Rust files)`);
    console.log(`[pgo] training sha256: ${metadata.training.sha256} (${frozen.modules.length} modules, ` +
      `${frozen.manifest.length} frozen files, excluded ${frozen.excluded.length ? frozen.excluded.join(", ") : "none"})`);
    console.log(`[pgo] profile sha256: ${metadata.profile.merged.sha256} (${rawProfiles.length} raw profiles)`);
    console.log(`[pgo] binary sha256: ${metadata.binary.sha256} (${binary})`);
    console.log(`[pgo] metadata: ${metadataPath}`);
    return { binary, profilePath: metadataPath, metadata };
  } catch (error) {
    metadata.status = "failed";
    metadata.stage = stage;
    metadata.failure = error?.stack ?? String(error);
    metadata.finished_at = new Date().toISOString();
    save();
    throw error;
  }
}
