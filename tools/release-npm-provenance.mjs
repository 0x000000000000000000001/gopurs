#!/usr/bin/env node
import { execFileSync } from "node:child_process";
import { createHash } from "node:crypto";
import { accessSync, constants, readFileSync, realpathSync, writeFileSync } from "node:fs";
import { createRequire } from "node:module";
import { basename, delimiter, dirname, join, resolve } from "node:path";
import { fileURLToPath } from "node:url";

export const repository = "0x000000000000000000001/gopurs";
const repositoryURL = `https://github.com/${repository}`;
const workflowPath = ".github/workflows/ci.yml";
const payloadType = "application/vnd.in-toto+json";
const workflowBuildType = "https://slsa-framework.github.io/github-actions-buildtypes/workflow/v1";
const root = fileURLToPath(new URL("../", import.meta.url));
const hash = (bytes, algorithm) => createHash(algorithm).update(bytes).digest("hex");
const requireValue = (condition, message) => { if (!condition) throw new Error(message); };

// Keep the signed subject compatible with npm 11.6.0's libnpmpublish verifier.
// The public Sigstore API is loaded from that pinned CLI, not a project install.
export function loadPinnedSigstore(env = process.env) {
  const executable = (env.PATH ?? "").split(delimiter).map(path => join(path, "npm")).find(path => {
    try { accessSync(path, constants.X_OK); return true; } catch { return false; }
  });
  requireValue(executable, "npm 11.6.0 must be on PATH (use nix develop)");
  const npmRoot = dirname(dirname(realpathSync(executable)));
  const npmManifest = JSON.parse(readFileSync(join(npmRoot, "package.json"), "utf8"));
  requireValue(npmManifest.name === "npm" && npmManifest.version === "11.6.0", "Expected pinned npm 11.6.0");
  const npmRequire = createRequire(join(npmRoot, "package.json"));
  requireValue(npmRequire("sigstore/package.json").version === "3.1.0", "Expected npm's Sigstore 3.1.0");
  requireValue(npmRequire("@sigstore/sign/package.json").version === "3.1.0", "Expected npm's @sigstore/sign 3.1.0");
  requireValue(npmRequire("@sigstore/bundle/package.json").version === "3.1.0", "Expected npm's @sigstore/bundle 3.1.0");
  requireValue(npmRequire("libnpmpublish/package.json").version === "11.1.0", "Expected npm's libnpmpublish 11.1.0");
  return { ...npmRequire("sigstore"), attest: createAttestor(npmRequire("@sigstore/sign"), npmRequire("@sigstore/bundle").bundleToJSON) };
}

export function createAttestor({ DSSEBundleBuilder, FulcioSigner, CIContextProvider, RekorWitness }, bundleToJSON) {
  return async (data, type) => {
    // Sigstore 3.1.0's high-level attest hardcodes fetchOnConflict=false.
    // Compose its public APIs to recover a POST accepted before a lost response.
    // Keep its CI provider, DSSE format, timeout and retry defaults unchanged.
    const request = { timeout: 5000, retry: { retries: 2 } };
    const builder = new DSSEBundleBuilder({
      signer: new FulcioSigner({ ...request, identityProvider: new CIContextProvider("sigstore") }),
      witnesses: [new RekorWitness({ ...request, entryType: "dsse", fetchOnConflict: true })],
    });
    // A fetched entry is evidence only; generateBundle still verifies the
    // signature, transparency log and exact release/workflow identity below.
    return bundleToJSON(await builder.create({ data, type }));
  };
}

export function describeRelease({ manifest, sourceManifest, build, archiveName, bytes, git }) {
  requireValue(manifest.name === "gopurs" && sourceManifest.name === "gopurs", "Expected the gopurs package");
  const version = manifest.version;
  requireValue(/^\d+\.\d+\.\d+(?:-rc\.\d+)?$/.test(version), "Unsupported release version");
  requireValue(sourceManifest.version === version && build.version === version, "Package/build/source versions differ");
  for (const value of [manifest, sourceManifest]) {
    requireValue(value.repository?.url === `git+${repositoryURL}.git`, "Package repository does not match the release repository");
  }
  requireValue(archiveName === `gopurs-${version}.tgz`, "Unexpected release archive name");
  const tag = `v${version}`;
  requireValue(build.schema === 1 && build.preview === false && build.source?.dirty === false, "Preview or dirty builds cannot be attested");
  requireValue(build.source.tag === tag && /^[a-f0-9]{40}$/.test(build.source.commit), "Build source tag/commit is invalid");
  requireValue(git.status === "" && git.tagType === "tag" && git.commit === build.source.commit && git.tagCommit === git.commit,
    "A clean checkout at the matching annotated release tag is required");
  const artifacts = build.artifacts?.filter(artifact => artifact.name === archiveName) ?? [];
  requireValue(artifacts.length === 1 && artifacts[0].sha256 === hash(bytes, "sha256"), "Archive SHA-256 differs from the build metadata");
  return { version, tag, commit: git.commit, subject: { name: `pkg:npm/gopurs@${version}`, digest: { sha512: hash(bytes, "sha512") } } };
}

export function readRelease(archive, buildFile, cwd = root) {
  const bytes = readFileSync(archive);
  const manifest = JSON.parse(execFileSync("tar", ["-xOf", archive, "package/package.json"], { encoding: "utf8", maxBuffer: 1024 * 1024 }));
  const sourceManifest = JSON.parse(readFileSync(join(cwd, "package.json"), "utf8"));
  const build = JSON.parse(readFileSync(buildFile, "utf8"));
  requireValue(/^\d+\.\d+\.\d+(?:-rc\.\d+)?$/.test(manifest.version), "Unsupported release version");
  const gitCommand = args => execFileSync("git", args, { cwd, encoding: "utf8" }).trim();
  const tagRef = `refs/tags/v${manifest.version}`;
  return describeRelease({ manifest, sourceManifest, build, archiveName: basename(archive), bytes, git: {
    status: gitCommand(["status", "--porcelain", "--untracked-files=all"]),
    tagType: gitCommand(["cat-file", "-t", tagRef]),
    commit: gitCommand(["rev-parse", "HEAD"]),
    tagCommit: gitCommand(["rev-parse", `${tagRef}^{commit}`]),
  } });
}

export function requireGitHubTag(release, env = process.env) {
  requireValue(env.CI === "true" && env.GITHUB_ACTIONS === "true" && env.RUNNER_ENVIRONMENT === "github-hosted", "Generation requires a GitHub-hosted Actions job");
  requireValue(env.GITHUB_SERVER_URL === "https://github.com" && env.GITHUB_REPOSITORY === repository, "Unexpected GitHub repository");
  requireValue(["push", "workflow_dispatch"].includes(env.GITHUB_EVENT_NAME), "Generation requires a tag push or tag workflow dispatch");
  const ref = `refs/tags/${release.tag}`;
  requireValue(env.GITHUB_REF_TYPE === "tag" && env.GITHUB_REF === ref && env.GITHUB_REF_NAME === release.tag, "Generation requires the matching release tag");
  requireValue(env.GITHUB_SHA === release.commit && env.GITHUB_WORKFLOW_SHA === release.commit, "GitHub source/workflow commit differs from the release");
  requireValue(env.GITHUB_WORKFLOW_REF === `${repository}/${workflowPath}@${ref}`, "Unexpected release workflow identity");
  for (const key of ["GITHUB_REPOSITORY_ID", "GITHUB_REPOSITORY_OWNER_ID", "GITHUB_RUN_ID", "GITHUB_RUN_ATTEMPT"]) {
    requireValue(/^[1-9]\d*$/.test(env[key] ?? ""), `Missing or invalid ${key}`);
  }
  requireValue(!env.SIGSTORE_ID_TOKEN, "Pre-supplied Sigstore identity tokens are not accepted");
  requireValue(Boolean(env.ACTIONS_ID_TOKEN_REQUEST_TOKEN), "The job requires id-token: write");
  const oidcURL = new URL(env.ACTIONS_ID_TOKEN_REQUEST_URL ?? "https://invalid.invalid");
  requireValue(oidcURL.protocol === "https:" && oidcURL.hostname.endsWith(".actions.githubusercontent.com") && !oidcURL.username && !oidcURL.password,
    "Missing GitHub Actions OIDC request endpoint");
}

export function createStatement(release, env = process.env) {
  requireGitHubTag(release, env);
  const ref = `refs/tags/${release.tag}`;
  // SLSA v1 workflow fields follow npm CLI 11.6.0's provenance format:
  // https://github.com/npm/cli/blob/v11.6.0/workspaces/libnpmpublish/lib/provenance.js
  return {
    _type: "https://in-toto.io/Statement/v1",
    subject: [structuredClone(release.subject)],
    predicateType: "https://slsa.dev/provenance/v1",
    predicate: {
      buildDefinition: {
        buildType: workflowBuildType,
        externalParameters: { workflow: { ref, repository: repositoryURL, path: workflowPath } },
        internalParameters: { github: { event_name: env.GITHUB_EVENT_NAME, repository_id: env.GITHUB_REPOSITORY_ID, repository_owner_id: env.GITHUB_REPOSITORY_OWNER_ID } },
        resolvedDependencies: [{ uri: `git+${repositoryURL}@${ref}`, digest: { gitCommit: release.commit } }],
      },
      runDetails: {
        builder: { id: "https://github.com/actions/runner/github-hosted" },
        metadata: { invocationId: `${repositoryURL}/actions/runs/${env.GITHUB_RUN_ID}/attempts/${env.GITHUB_RUN_ATTEMPT}` },
      },
    },
  };
}

export async function verifyBundle(bundle, release, sigstore) {
  requireValue(bundle.dsseEnvelope?.payloadType === payloadType, "Expected an in-toto DSSE bundle");
  const statement = JSON.parse(Buffer.from(bundle.dsseEnvelope.payload, "base64").toString("utf8"));
  requireValue(statement._type === "https://in-toto.io/Statement/v1" && statement.predicateType === "https://slsa.dev/provenance/v1", "Expected SLSA v1 provenance");
  requireValue(statement.subject?.length === 1 && statement.subject[0].name === release.subject.name
    && statement.subject[0].digest?.sha512 === release.subject.digest.sha512, "Provenance must have exactly this npm subject and archive SHA-512");
  const definition = statement.predicate?.buildDefinition;
  const workflow = definition?.externalParameters?.workflow;
  const ref = `refs/tags/${release.tag}`;
  requireValue(definition?.buildType === workflowBuildType && workflow?.repository === repositoryURL && workflow?.path === workflowPath && workflow?.ref === ref,
    "Provenance repository/workflow/tag differs from the release");
  const dependencies = definition.resolvedDependencies;
  requireValue(dependencies?.length === 1 && dependencies[0].uri === `git+${repositoryURL}@${ref}` && dependencies[0].digest?.gitCommit === release.commit,
    "Provenance source commit differs from the release");
  requireValue(statement.predicate.runDetails?.builder?.id === "https://github.com/actions/runner/github-hosted", "Expected a GitHub-hosted builder");
  const invocation = statement.predicate.runDetails?.metadata?.invocationId;
  requireValue(typeof invocation === "string" && invocation.startsWith(`${repositoryURL}/actions/runs/`)
    && /^\d+\/attempts\/\d+$/.test(invocation.slice(`${repositoryURL}/actions/runs/`.length)), "Invalid provenance run identity");
  // These certificate constraints authenticate the workflow identity; parsing
  // its JSON payload alone is never considered signature verification.
  await sigstore.verify(bundle, {
    certificateIssuer: "https://token.actions.githubusercontent.com",
    certificateIdentityURI: new RegExp(`^${`${repositoryURL}/${workflowPath}@${ref}`.replace(/[.*+?^${}()|[\]\\]/g, "\\$&")}$`),
  });
  return statement;
}

export async function generateBundle(release, sigstore, env = process.env) {
  const statement = createStatement(release, env);
  // No identityToken, custom identityProvider or unsigned/local fallback.
  const bundle = await sigstore.attest(Buffer.from(JSON.stringify(statement)), payloadType);
  await verifyBundle(bundle, release, sigstore);
  return bundle;
}

async function publicRegistryJSON(url) {
  const response = await fetch(url, { redirect: "error", signal: AbortSignal.timeout(30_000), headers: { accept: "application/json" } });
  requireValue(response.ok, `Public npm registry returned HTTP ${response.status}`);
  return response.json();
}

function sameSignatures(actual, expected) {
  const valid = signatures => Array.isArray(signatures) && signatures.length > 0 && signatures.every(signature =>
    signature !== null && typeof signature === "object" && !Array.isArray(signature)
    && Object.keys(signature).every(key => key === "sig" || key === "keyid")
    && typeof signature.sig === "string"
    && (!Object.hasOwn(signature, "keyid") || typeof signature.keyid === "string"));
  // npm serializes an absent DSSE keyid as "". Keep every signature and any
  // nonempty key identifier exact; reject unknown fields before comparing.
  return valid(actual) && valid(expected) && actual.length === expected.length
    && actual.every((signature, index) => signature.sig === expected[index].sig
      && (signature.keyid ?? "") === (expected[index].keyid ?? ""));
}

export async function verifyPublished(release, approvedBundle, sigstore, fetchJSON = publicRegistryJSON) {
  await verifyBundle(approvedBundle, release, sigstore);
  const manifest = await fetchJSON(`https://registry.npmjs.org/gopurs/${release.version}`);
  const integrity = `sha512-${Buffer.from(release.subject.digest.sha512, "hex").toString("base64")}`;
  requireValue(manifest.name === "gopurs" && manifest.version === release.version
    && manifest.dist?.integrity?.split(/\s+/).includes(integrity), "Published package/version/SHA-512 differs from the approved archive");
  const url = new URL(manifest.dist.attestations?.url ?? "https://invalid.invalid");
  requireValue(url.origin === "https://registry.npmjs.org" && !url.username && !url.password
    && url.pathname.startsWith("/-/npm/v1/attestations/"), "Published package has no public npm attestation endpoint");
  const { attestations } = await fetchJSON(url.href);
  const provenance = attestations?.filter(value => value.predicateType === "https://slsa.dev/provenance/v1") ?? [];
  requireValue(provenance.length === 1, "Expected exactly one published SLSA v1 provenance bundle");
  const bundle = provenance[0].bundle;
  await verifyBundle(bundle, release, sigstore);
  requireValue(bundle.dsseEnvelope.payloadType === approvedBundle.dsseEnvelope.payloadType
    && bundle.dsseEnvelope.payload === approvedBundle.dsseEnvelope.payload
    && sameSignatures(bundle.dsseEnvelope.signatures, approvedBundle.dsseEnvelope.signatures), "Published provenance differs from the approved signed CI envelope");
  return { package: manifest.name, version: manifest.version, integrity, attestationURL: url.href, provenanceBundle: bundle };
}

async function main(args) {
  const mode = args.shift();
  if (mode === "--help") {
    console.log("Usage: node tools/release-npm-provenance.mjs generate|verify|verify-published --package FILE.tgz --build FILE-build.json --bundle FILE.sigstore.json\nGeneration requires the actual tagged GitHub CI job. Verification is also available locally from the clean tagged checkout; all modes use npm 11.6.0's Sigstore. verify-published additionally requires npm to serve the same signed CI envelope.");
    return;
  }
  requireValue(["generate", "verify", "verify-published"].includes(mode), "Expected generate, verify or verify-published");
  const options = {};
  for (let index = 0; index < args.length; index += 2) {
    requireValue(["--package", "--build", "--bundle"].includes(args[index]) && args[index + 1] && !options[args[index]], "Unknown, duplicate or incomplete option");
    options[args[index]] = resolve(args[index + 1]);
  }
  requireValue(options["--package"] && options["--build"] && options["--bundle"], "--package, --build and --bundle are required");
  const release = readRelease(options["--package"], options["--build"]);
  const sigstore = loadPinnedSigstore();
  let published;
  if (mode === "generate") {
    const bundle = await generateBundle(release, sigstore);
    requireValue(hash(readFileSync(options["--package"]), "sha512") === release.subject.digest.sha512, "Archive changed during signing");
    writeFileSync(options["--bundle"], JSON.stringify(bundle, null, 2) + "\n", { flag: "wx" });
  } else {
    const bundle = JSON.parse(readFileSync(options["--bundle"], "utf8"));
    if (mode === "verify-published") published = await verifyPublished(release, bundle, sigstore);
    else await verifyBundle(bundle, release, sigstore);
    requireValue(hash(readFileSync(options["--package"]), "sha512") === release.subject.digest.sha512, "Archive changed during verification");
  }
  console.log(JSON.stringify({ mode, version: release.version, tag: release.tag, commit: release.commit, subject: release.subject, bundle: options["--bundle"], published }, null, 2));
}

if (process.argv[1] && resolve(process.argv[1]) === fileURLToPath(import.meta.url)) {
  main(process.argv.slice(2)).catch(error => { console.error(error.message); process.exitCode = 1; });
}
