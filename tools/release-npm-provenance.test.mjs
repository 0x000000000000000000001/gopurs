import assert from "node:assert/strict";
import { execFileSync } from "node:child_process";
import { createHash, generateKeyPairSync, sign } from "node:crypto";
import { once } from "node:events";
import { existsSync, mkdirSync, mkdtempSync, readFileSync, realpathSync, rmSync, writeFileSync } from "node:fs";
import { createServer } from "node:http";
import { createRequire } from "node:module";
import { tmpdir } from "node:os";
import { delimiter, join } from "node:path";
import test from "node:test";
import { createAttestor, createStatement, describeRelease, generateBundle, loadPinnedSigstore, readRelease, repository, requireGitHubTag, verifyBundle, verifyPublished } from "./release-npm-provenance.mjs";

const sha256 = bytes => createHash("sha256").update(bytes).digest("hex");
const commit = "a".repeat(40);
const manifest = { name: "gopurs", version: "0.1.0", repository: { type: "git", url: `git+https://github.com/${repository}.git` } };
function inputs() {
  const bytes = Buffer.from("unit-test archive bytes, not a publishable tarball");
  return { manifest: structuredClone(manifest), sourceManifest: structuredClone(manifest), archiveName: "gopurs-0.1.0.tgz", bytes,
    build: { schema: 1, version: "0.1.0", preview: false, source: { dirty: false, tag: "v0.1.0", commit }, artifacts: [{ name: "gopurs-0.1.0.tgz", sha256: sha256(bytes) }] },
    git: { status: "", tagType: "tag", commit, tagCommit: commit } };
}
const release = describeRelease(inputs());
function ciEnvironment() {
  return { CI: "true", GITHUB_ACTIONS: "true", RUNNER_ENVIRONMENT: "github-hosted", GITHUB_SERVER_URL: "https://github.com", GITHUB_REPOSITORY: repository,
    GITHUB_EVENT_NAME: "push", GITHUB_REF_TYPE: "tag", GITHUB_REF: "refs/tags/v0.1.0", GITHUB_REF_NAME: "v0.1.0", GITHUB_SHA: commit, GITHUB_WORKFLOW_SHA: commit,
    GITHUB_WORKFLOW_REF: `${repository}/.github/workflows/ci.yml@refs/tags/v0.1.0`, GITHUB_REPOSITORY_ID: "123", GITHUB_REPOSITORY_OWNER_ID: "456",
    GITHUB_RUN_ID: "789", GITHUB_RUN_ATTEMPT: "1", ACTIONS_ID_TOKEN_REQUEST_URL: "https://unit-test.actions.githubusercontent.com/token",
    ACTIONS_ID_TOKEN_REQUEST_TOKEN: "unit-test-sentinel-not-a-token" };
}
// Only unit-test envelopes: no signature/certificate and never written as a
// release artifact. Cryptographic checks below use an explicit verifier stub.
function envelope(statement = createStatement(release, ciEnvironment())) {
  return { dsseEnvelope: { payloadType: "application/vnd.in-toto+json", payload: Buffer.from(JSON.stringify(statement)).toString("base64"), signatures: [{ keyid: "", sig: "unit-test-not-a-signature" }] } };
}

test("release identity binds the exact tarball bytes, build metadata and annotated tag", () => {
  assert.equal(release.subject.name, "pkg:npm/gopurs@0.1.0");
  assert.equal(release.subject.digest.sha512, createHash("sha512").update(inputs().bytes).digest("hex"));
  const changes = [
    value => { value.bytes = Buffer.from("changed"); },
    value => { value.build.artifacts.push(value.build.artifacts[0]); },
    value => { value.build.preview = true; },
    value => { value.build.source.dirty = true; },
    value => { value.git.status = " M source"; },
    value => { value.git.tagType = "commit"; },
    value => { value.git.tagCommit = "b".repeat(40); },
    value => { value.build.source.commit = "b".repeat(40); },
    value => { value.build.source.tag = "v0.2.0"; },
    value => { value.sourceManifest.version = "0.2.0"; },
    value => { value.manifest.repository.url = "git+https://github.com/other/gopurs.git"; },
    value => { value.archiveName = "other.tgz"; },
  ];
  for (const change of changes) { const value = inputs(); change(value); assert.throws(() => describeRelease(value)); }
});

test("readRelease examines a real tarball and git tag without extracting into the checkout", t => {
  const temp = mkdtempSync(join(tmpdir(), "gopurs-provenance-test-"));
  t.after(() => rmSync(temp, { recursive: true, force: true }));
  const checkout = join(temp, "source"), packed = join(temp, "packed");
  mkdirSync(checkout); mkdirSync(join(packed, "package"), { recursive: true });
  for (const file of [join(checkout, "package.json"), join(packed, "package/package.json")]) writeFileSync(file, JSON.stringify(manifest));
  const git = args => execFileSync("git", args, { cwd: checkout, encoding: "utf8", stdio: ["ignore", "pipe", "pipe"] }).trim();
  git(["init", "-q"]); git(["config", "user.name", "Provenance unit test"]); git(["config", "user.email", "test@example.invalid"]);
  git(["add", "package.json"]); git(["-c", "commit.gpgsign=false", "commit", "-qm", "unit test"]);
  git(["-c", "tag.gpgsign=false", "tag", "-a", "v0.1.0", "-m", "unit test"]);
  const archive = join(temp, "gopurs-0.1.0.tgz"), buildFile = join(temp, "build.json");
  execFileSync("tar", ["-czf", archive, "-C", packed, "package"]);
  const build = inputs().build;
  build.source.commit = git(["rev-parse", "HEAD"]);
  build.artifacts[0].sha256 = sha256(readFileSync(archive));
  writeFileSync(buildFile, JSON.stringify(build));
  assert.equal(readRelease(archive, buildFile, checkout).commit, build.source.commit);
  assert.equal(git(["status", "--porcelain"]), "");
  writeFileSync(join(checkout, "untracked"), "changed checkout");
  assert.throws(() => readRelease(archive, buildFile, checkout), /clean checkout/);
});

test("generation refuses non-CI, forks, branches, mismatched commits and supplied identity tokens before signing", async () => {
  const changes = [
    { CI: undefined }, { GITHUB_ACTIONS: "false" }, { RUNNER_ENVIRONMENT: "self-hosted" },
    { GITHUB_REPOSITORY: "fork/gopurs" }, { GITHUB_EVENT_NAME: "pull_request" }, { GITHUB_REF_TYPE: "branch" },
    { GITHUB_REF: "refs/heads/main" }, { GITHUB_REF_NAME: "v0.2.0" }, { GITHUB_SHA: "b".repeat(40) },
    { GITHUB_WORKFLOW_SHA: "b".repeat(40) }, { GITHUB_WORKFLOW_REF: `${repository}/.github/workflows/publish.yml@refs/tags/v0.1.0` },
    { GITHUB_RUN_ID: undefined }, { SIGSTORE_ID_TOKEN: "unit-test-rejected-token" }, { ACTIONS_ID_TOKEN_REQUEST_TOKEN: undefined },
    { ACTIONS_ID_TOKEN_REQUEST_URL: "http://unit-test.actions.githubusercontent.com/token" },
    { ACTIONS_ID_TOKEN_REQUEST_URL: "https://actions.githubusercontent.com.evil.invalid/token" },
  ];
  let calls = 0;
  const signer = { attest() { calls++; throw new Error("must not sign"); } };
  for (const change of changes) await assert.rejects(generateBundle(release, signer, { ...ciEnvironment(), ...change }));
  await assert.rejects(generateBundle(release, signer, {}));
  assert.equal(calls, 0);
  assert.doesNotThrow(() => requireGitHubTag(release, { ...ciEnvironment(), GITHUB_EVENT_NAME: "workflow_dispatch" }));
});

test("statement has npm's single purl/SHA-512 subject, real CI identifiers and no bearer token", () => {
  const statement = createStatement(release, ciEnvironment());
  assert.deepEqual(statement.subject, [release.subject]);
  assert.equal(statement.predicate.buildDefinition.resolvedDependencies[0].digest.gitCommit, commit);
  assert.equal(statement.predicate.runDetails.metadata.invocationId, `https://github.com/${repository}/actions/runs/789/attempts/1`);
  const serialized = JSON.stringify(statement);
  assert(!serialized.includes("unit-test-sentinel-not-a-token"));
  assert(!serialized.includes("ACTIONS_ID_TOKEN_REQUEST"));
});

test("verification rejects substituted artifacts or source identities before calling the signature verifier", async () => {
  const changes = [
    value => { value.subject.push(value.subject[0]); },
    value => { value.subject[0].digest.sha512 = "b".repeat(128); },
    value => { value.subject[0].name = "pkg:npm/other@0.1.0"; },
    value => { value.predicateType = "https://slsa.dev/provenance/v0.2"; },
    value => { value.predicate.buildDefinition.externalParameters.workflow.ref = "refs/heads/main"; },
    value => { value.predicate.buildDefinition.externalParameters.workflow.repository = "https://github.com/fork/gopurs"; },
    value => { value.predicate.buildDefinition.resolvedDependencies[0].digest.gitCommit = "b".repeat(40); },
    value => { value.predicate.runDetails.builder.id = "https://github.com/actions/runner/self-hosted"; },
    value => { value.predicate.runDetails.metadata.invocationId = "https://example.invalid/run"; },
  ];
  let calls = 0;
  for (const change of changes) {
    const statement = createStatement(release, ciEnvironment()); change(statement);
    await assert.rejects(verifyBundle(envelope(statement), release, { verify() { calls++; } }));
  }
  assert.equal(calls, 0);
});

test("generation verifies the signed result with exact issuer and workflow certificate identity", async () => {
  const calls = [];
  const sigstore = {
    async attest(bytes, type, ...options) {
      calls.push("attest");
      assert.equal(type, "application/vnd.in-toto+json");
      assert.deepEqual(options, [], "use the real CI identity provider without a token override");
      return envelope(JSON.parse(bytes.toString("utf8")));
    },
    async verify(bundle, options) {
      calls.push("verify");
      assert.equal(options.certificateIssuer, "https://token.actions.githubusercontent.com");
      const identity = `https://github.com/${repository}/.github/workflows/ci.yml@refs/tags/v0.1.0`;
      assert(options.certificateIdentityURI.test(identity));
      for (const wrong of [`prefix${identity}`, `${identity}/suffix`, identity.replace("github.com", "githubXcom"), identity.replace("v0.1.0", "v0X1X0")]) {
        assert.equal(options.certificateIdentityURI.test(wrong), false);
      }
      assert.equal(bundle.dsseEnvelope.payloadType, "application/vnd.in-toto+json");
    },
  };
  await generateBundle(release, sigstore, ciEnvironment());
  assert.deepEqual(calls, ["attest", "verify"]);
  await assert.rejects(verifyBundle(envelope(), release, { async verify() { throw new Error("invalid certificate"); } }), /invalid certificate/);
});

test("the pinned npm distribution provides the public Sigstore APIs used by the release", () => {
  const sigstore = loadPinnedSigstore();
  assert.equal(typeof sigstore.attest, "function");
  assert.equal(typeof sigstore.verify, "function");
});

test("Rekor conflicts fetch the existing entry and still require matching signatures, payload and identity", { timeout: 10_000 }, async t => {
  loadPinnedSigstore(); // Check the exact npm/SDK versions before loading their public test APIs.
  const npm = process.env.PATH.split(delimiter).map(path => join(path, "npm")).find(existsSync);
  const npmRequire = createRequire(realpathSync(npm));
  const signing = npmRequire("@sigstore/sign"), bundles = npmRequire("@sigstore/bundle");
  const { Verifier, toSignedEntity } = npmRequire("@sigstore/verify");
  const { json } = npmRequire("@sigstore/core");
  // Ephemeral test keys only: no Fulcio/OIDC calls or public transparency-log writes.
  const signerKey = generateKeyPairSync("ec", { namedCurve: "prime256v1" });
  const logKey = generateKeyPairSync("ec", { namedCurve: "prime256v1" });
  const signerPEM = signerKey.publicKey.export({ type: "spki", format: "pem" });
  const logID = sha256(logKey.publicKey.export({ type: "spki", format: "der" }));
  const uuid = "a".repeat(64), requests = [];
  let mode = "equivalent", proposed;
  function existingEntry() {
    const dsse = JSON.parse(proposed.spec.proposedContent.envelope);
    const body = { apiVersion: "0.0.1", kind: "dsse", spec: {
      payloadHash: { algorithm: "sha256", value: sha256(Buffer.from(dsse.payload, "base64")) },
      signatures: [{ signature: dsse.signatures[0].sig, verifier: proposed.spec.proposedContent.verifiers[0] }],
    } };
    if (mode === "payload") body.spec.payloadHash.value = "b".repeat(64);
    if (mode === "signature") body.spec.signatures[0].signature = Buffer.from("wrong signature").toString("base64");
    const canonical = Buffer.from(json.canonicalize(body));
    const entry = { body: canonical.toString("base64"), integratedTime: 1700000000, logIndex: 0, logID };
    const set = sign("sha256", Buffer.from(json.canonicalize(entry)), logKey.privateKey);
    const rootHash = sha256(Buffer.concat([Buffer.from([0]), canonical]));
    const note = `test-rekor\n1\n${Buffer.from(rootHash, "hex").toString("base64")}\n`;
    const checkpointSignature = Buffer.concat([Buffer.from(logID, "hex").subarray(0, 4), sign("sha256", Buffer.from(note), logKey.privateKey)]);
    entry.verification = {
      signedEntryTimestamp: set.toString("base64"),
      inclusionProof: { logIndex: 0, treeSize: 1, rootHash, hashes: [], checkpoint: `${note}\n— test-rekor ${checkpointSignature.toString("base64")}\n` },
    };
    if (mode === "log-signature") entry.verification.signedEntryTimestamp = Buffer.from("wrong SET").toString("base64");
    return { [uuid]: entry };
  }
  const server = createServer(async (request, response) => {
    requests.push(`${request.method} ${request.url}`);
    response.setHeader("content-type", "application/json");
    if (request.method === "POST" && request.url === "/api/v1/log/entries") {
      const chunks = [];
      for await (const chunk of request) chunks.push(chunk);
      proposed = JSON.parse(Buffer.concat(chunks));
      response.statusCode = 409;
      if (mode !== "no-location") response.setHeader("location", `/api/v1/log/entries/${uuid}`);
      response.end(JSON.stringify({ message: "an equivalent entry already exists" }));
    } else if (request.method === "GET" && request.url === `/api/v1/log/entries/${uuid}`) {
      response.statusCode = mode === "missing" ? 404 : 200;
      response.end(JSON.stringify(mode === "missing" ? { message: "entry missing" } : existingEntry()));
    } else {
      response.statusCode = 404; response.end("{}");
    }
  });
  server.listen(0, "127.0.0.1");
  await once(server, "listening");
  t.after(() => new Promise(resolve => { server.close(resolve); server.closeAllConnections(); }));
  class TestSigner {
    constructor(options) {
      assert(options.identityProvider instanceof signing.CIContextProvider);
      assert.equal(options.identityProvider.audience, "sigstore");
    }
    async sign(data) {
      return { signature: sign("sha256", data, signerKey.privateKey), key: { $case: "publicKey", publicKey: signerPEM, hint: "test-key" } };
    }
  }
  class LocalRekorWitness extends signing.RekorWitness {
    constructor(options) { super({ ...options, rekorBaseURL: `http://127.0.0.1:${server.address().port}`, retry: 0 }); }
  }
  const attest = createAttestor({ ...signing, FulcioSigner: TestSigner, RekorWitness: LocalRekorWitness }, bundles.bundleToJSON);
  const verifier = new Verifier({
    certificateAuthorities: [], timestampAuthorities: [], ctlogs: [],
    tlogs: [{ logID: Buffer.from(logID, "hex"), publicKey: logKey.publicKey, validFor: { start: new Date(0), end: new Date(8640000000000000) } }],
    publicKey: hint => { assert.equal(hint, "test-key"); return { publicKey: signerKey.publicKey, validFor: () => true }; },
  });
  const verifyCrypto = (value, policy) => verifier.verify(toSignedEntity(bundles.bundleFromJSON(value)), policy);
  const bytes = Buffer.from(JSON.stringify(createStatement(release, ciEnvironment())));
  const type = "application/vnd.in-toto+json";
  const recovered = await attest(bytes, type);
  assert.deepEqual(requests, ["POST /api/v1/log/entries", `GET /api/v1/log/entries/${uuid}`]);
  assert.doesNotThrow(() => verifyCrypto(recovered));
  for (const [bad, error] of [["payload", /payload hash mismatch/], ["signature", /signature mismatch/], ["log-signature", /inclusion promise/]]) {
    mode = bad;
    const substituted = await attest(bytes, type);
    assert.throws(() => verifyCrypto(substituted), error);
  }
  // A valid recovered log entry with no authenticated GitHub certificate must
  // still fail generateBundle's unchanged issuer/workflow identity policy.
  mode = "equivalent";
  await assert.rejects(generateBundle(release, { attest, verify(value, policy) {
    return verifyCrypto(value, { subjectAlternativeName: policy.certificateIdentityURI, extensions: { issuer: policy.certificateIssuer } });
  } }, ciEnvironment()), /certificate identity error/);
  for (const bad of ["no-location", "missing"]) {
    mode = bad; requests.length = 0;
    await assert.rejects(attest(bytes, type), bad === "missing" ? /fetching tlog entry/ : /creating tlog entry/);
    assert.equal(requests.length, bad === "missing" ? 2 : 1);
  }
});

test("first-publication recovery requires npm to serve the approved signed CI envelope and exact integrity", async () => {
  const approved = envelope();
  const published = { name: "gopurs", version: "0.1.0", dist: {
    integrity: `sha512-${Buffer.from(release.subject.digest.sha512, "hex").toString("base64")}`,
    attestations: { url: "https://registry.npmjs.org/-/npm/v1/attestations/gopurs@0.1.0" },
  } };
  const evidence = { attestations: [{ predicateType: "https://slsa.dev/provenance/v1", bundle: approved }] };
  let verifies = 0;
  const sigstore = { async verify() { verifies++; } };
  const fetchJSON = async url => {
    if (url === "https://registry.npmjs.org/gopurs/0.1.0") return published;
    assert.equal(url, published.dist.attestations.url);
    return evidence;
  };
  const result = await verifyPublished(release, approved, sigstore, fetchJSON);
  assert.equal(result.integrity, published.dist.integrity);
  assert.equal(verifies, 2);
  const changes = [
    value => { value.name = "different"; },
    value => { value.version = "0.2.0"; },
    value => { value.dist.integrity = "sha512-wrong"; },
    value => { delete value.dist.attestations; },
    value => { value.dist.attestations.url = "https://untrusted.invalid/-/npm/v1/attestations/gopurs@0.1.0"; },
  ];
  for (const change of changes) {
    const bad = structuredClone(published); change(bad);
    await assert.rejects(verifyPublished(release, approved, sigstore, async () => bad));
  }
  await assert.rejects(verifyPublished(release, approved, sigstore, async url => url.includes("/-/") ? { attestations: [] } : published), /exactly one/);
  const differentRun = structuredClone(evidence);
  differentRun.attestations[0].bundle = envelope(createStatement(release, { ...ciEnvironment(), GITHUB_RUN_ID: "999" }));
  await assert.rejects(verifyPublished(release, approved, sigstore, async url => url.includes("/-/") ? differentRun : published), /approved signed CI envelope/);
});

test("published DSSE envelopes allow only absent/empty keyid normalization", async () => {
  const manifest = { name: "gopurs", version: release.version, dist: {
    integrity: `sha512-${Buffer.from(release.subject.digest.sha512, "hex").toString("base64")}`,
    attestations: { url: "https://registry.npmjs.org/-/npm/v1/attestations/gopurs@0.1.0" },
  } };
  const approved = envelope(), published = structuredClone(approved);
  const fetchJSON = async url => url.includes("/-/")
    ? { attestations: [{ predicateType: "https://slsa.dev/provenance/v1", bundle: published }] } : manifest;
  const verified = [];
  const sigstore = { async verify(bundle) { verified.push(bundle); } };
  const check = () => verifyPublished(release, approved, sigstore, fetchJSON);
  delete approved.dsseEnvelope.signatures[0].keyid;
  await check();
  assert.deepEqual(verified, [approved, published], "both original bundles still undergo signature verification");
  approved.dsseEnvelope.signatures[0].keyid = "";
  delete published.dsseEnvelope.signatures[0].keyid;
  await check();
  approved.dsseEnvelope.signatures[0].keyid = "same-key";
  published.dsseEnvelope.signatures[0].keyid = "same-key";
  await check();

  const approvedEnvelope = structuredClone(approved.dsseEnvelope);
  const changes = [
    value => { value.payload = Buffer.from(`${Buffer.from(value.payload, "base64")} `).toString("base64"); },
    value => { value.payloadType = "different"; },
    value => { value.signatures[0].sig = "different"; },
    value => { value.signatures[0].keyid = "different-key"; },
    value => { value.signatures[0].keyid = ""; },
    value => { delete value.signatures[0].keyid; },
    value => { value.signatures[0].extra = "unexpected"; },
    value => { value.signatures.push(value.signatures[0]); },
    value => { value.signatures = []; },
  ];
  for (const change of changes) {
    published.dsseEnvelope = structuredClone(approvedEnvelope); change(published.dsseEnvelope);
    await assert.rejects(check(), /approved signed CI envelope|in-toto DSSE bundle/);
  }
  // Unknown fields and malformed signatures must not become accepted just
  // because canonicalization drops them, even if both bundles contain them.
  for (const signature of [{ sig: "s", extra: "unexpected" }, { sig: "s", keyid: null }, { keyid: "" }]) {
    approved.dsseEnvelope.signatures = [signature];
    published.dsseEnvelope = structuredClone(approved.dsseEnvelope);
    await assert.rejects(check(), /approved signed CI envelope/);
  }
  approved.dsseEnvelope.signatures = [{ sig: "first" }, { sig: "second" }];
  published.dsseEnvelope = structuredClone(approved.dsseEnvelope);
  published.dsseEnvelope.signatures.reverse();
  await assert.rejects(check(), /approved signed CI envelope/);
  approved.dsseEnvelope = structuredClone(approvedEnvelope);
  published.dsseEnvelope = structuredClone(approvedEnvelope);
  for (const rejected of [approved, published]) {
    await assert.rejects(verifyPublished(release, approved, { async verify(bundle) {
      if (bundle === rejected) throw new Error("invalid signature");
    } }, fetchJSON), /invalid signature/);
  }
});
