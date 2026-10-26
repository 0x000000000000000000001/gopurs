# Security policy

## Supported versions

During `0.x`, security fixes target the latest minor release. There is no LTS
branch before `1.0.0`. Old checkouts, generated outputs and caches are not
supported after the release that supersedes them.

## Reporting a vulnerability

Report suspected vulnerabilities privately through GitHub's security advisory
form:

<https://github.com/0x000000000000000000001/gopurs/security/advisories/new>

Do not open a public issue, discussion or pull request before the maintainers
have had a chance to assess the report.

Please include, when possible:

- affected gopurs version or commit;
- host (`js`, `native` or `rust`) and platform;
- typed compiler fork revision and Go/Node versions;
- a minimal reproduction, generated output or corpus hash;
- expected and actual impact;
- whether the issue affects the compiler, generated code, runtime, FFI bridge
  or the release artifacts.

## Scope

Examples of in-scope reports:

- arbitrary code execution or unexpected file writes caused by the compiler;
- incorrect FFI bridging or runtime behavior with a security impact;
- cache or output publication that can execute stale or attacker-controlled
  code;
- vulnerabilities in the published npm package or release artifacts;
- supply-chain issues in the pinned build inputs.

General performance, compatibility and missing-FFI reports are normal bugs and
can use the public issue tracker. Do not include secrets, private keys or
personal data in any report.

## Response

Reports are handled on a best-effort basis. The maintainers will acknowledge the
report, reproduce it, agree on a disclosure timeline where practical and credit
the reporter unless anonymity is requested. No response-time SLA is offered for
the experimental `0.x` series.

## Safe harbor

Good-faith research that respects user privacy, avoids destructive testing and
does not access data beyond what is needed to demonstrate the issue is welcome.
Do not test against systems or data you do not own or have permission to use.
