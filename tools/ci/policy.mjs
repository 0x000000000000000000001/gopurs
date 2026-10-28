import { appendFileSync, readFileSync, writeFileSync } from 'node:fs';
import { execFileSync } from 'node:child_process';
import { resolve } from 'node:path';
import { fileURLToPath } from 'node:url';

const platforms = [
  { target: 'linux-amd64', os: 'ubuntu-24.04' },
  { target: 'linux-arm64', os: 'ubuntu-24.04-arm' },
  { target: 'darwin-arm64', os: 'macos-14' },
];
const rootDocs = new Set(['README.md', 'CHANGELOG.md', 'CODE_OF_CONDUCT.md',
  'CONTRIBUTING.md', 'SECURITY.md', 'THIRD_PARTY_NOTICES.md', 'todo.md']);

export function docsOnly(changes) {
  return changes.length > 0 && changes.every(({ path, oldMode, newMode }) =>
    [oldMode, newMode].every(mode => ['000000', '100644'].includes(mode)) &&
    !path.split('/').some(part => part === '.' || part === '..') &&
    (rootDocs.has(path) || /^docs\/(?:[A-Za-z0-9_.-]+\/)*[A-Za-z0-9_.-]+\.md$/.test(path)));
}

export function parseChanges(raw) {
  if (raw === '') return [];
  const fields = raw.split('\0');
  if (fields.pop() !== '' || fields.length % 2) throw new Error('Incomplete Git change list');
  return Array.from({ length: fields.length / 2 }, (_, index) => {
    const match = /^:([0-7]{6}) ([0-7]{6}) [a-f0-9]{40} [a-f0-9]{40} [ADM]$/.exec(fields[index * 2]);
    if (!match || !fields[index * 2 + 1]) throw new Error('Unsupported Git change');
    return { path: fields[index * 2 + 1], oldMode: match[1], newMode: match[2] };
  });
}

export function buildMatrix({ eventName, ref, full, changes }) {
  if (ref.startsWith('refs/tags/') || full === true) {
    return { build: true, reason: 'Complete platform qualification', matrix: { include: platforms } };
  }
  const documentation = ['pull_request', 'push'].includes(eventName) && docsOnly(changes);
  return { build: !documentation, reason: documentation ? 'Only regular Markdown documentation changed' : 'Linux quick qualification',
    matrix: { include: [platforms[0]] } };
}

export function validateTagVersion(ref, metadata, lockfile, changelog) {
  const version = metadata.version;
  if (typeof version !== 'string' || !/^\d+\.\d+\.\d+(?:-rc\.\d+)?$/.test(version)) throw new Error(`Unsupported release version: ${version}`);
  if (ref !== `refs/tags/v${version}`) throw new Error(`Workflow tag does not match v${version}`);
  if (lockfile.version !== version || lockfile.packages?.['']?.version !== version) throw new Error('Package and lockfile release versions differ');
  const heading = `## [${version}]`;
  if (!changelog.split(/\r?\n/).some(line => line === heading || line.startsWith(heading + ' '))) throw new Error(`Missing final changelog section for ${version}`);
}

async function main() {
  const env = process.env, event = JSON.parse(readFileSync(env.GITHUB_EVENT_PATH, 'utf8'));
  const git = (...args) => execFileSync('git', args, { encoding: 'utf8' }).trim();
  if (git('rev-parse', 'HEAD') !== env.GITHUB_SHA) throw new Error('Checkout differs from event commit');
  if (env.GITHUB_REF.startsWith('refs/tags/')) {
    validateTagVersion(env.GITHUB_REF, JSON.parse(readFileSync('package.json', 'utf8')),
      JSON.parse(readFileSync('package-lock.json', 'utf8')), readFileSync('CHANGELOG.md', 'utf8'));
  }
  let changes = [];
  try {
    const base = env.GITHUB_EVENT_NAME === 'pull_request' ? event.pull_request.base.sha : event.before;
    if (base && /^[a-f0-9]{40}$/.test(base) && !/^0+$/.test(base)) {
      changes = parseChanges(execFileSync('git', ['diff', '--raw', '-z', '--no-renames', '--no-abbrev', base, env.GITHUB_SHA], { encoding: 'utf8' }));
    }
  } catch (error) { console.log(`Changed paths unavailable; building: ${error.message}`); }
  const plan = buildMatrix({ eventName: env.GITHUB_EVENT_NAME, ref: env.GITHUB_REF,
    full: env.CI_FULL === 'true', changes });
  if (plan.build && env.GITHUB_EVENT_NAME === 'push' && env.GITHUB_REF === 'refs/heads/main' && env.CI_FULL !== 'true') {
    const { reuseQualification } = await import('./receipt.mjs');
    try {
      plan.reused = reuseQualification(env, git('rev-parse', 'HEAD^{tree}'));
      plan.build = false;
      plan.reason = `Reused successful PR qualification ${plan.reused.runId} of the identical source tree`;
    } catch (error) { plan.reuseUnavailable = error.message; }
  }
  writeFileSync(`${env.RUNNER_TEMP}/ci-plan.json`, JSON.stringify(plan, null, 2) + '\n');
  appendFileSync(env.GITHUB_OUTPUT, `build=${plan.build}\nmatrix=${JSON.stringify(plan.matrix)}\n`);
  console.log(JSON.stringify(plan, null, 2));
}

if (process.argv[1] && resolve(process.argv[1]) === fileURLToPath(import.meta.url)) await main();
