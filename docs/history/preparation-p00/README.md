# Preparation P00/P01 archive

These scripts were moved out of `tools/` on 8 November 2026. They are the
historical harnesses behind [the P00 report](../../build-preparation-baseline.md)
and [the P01 compact study](../../build-preparation-compact.md); they are kept as
evidence, not as maintained tooling.

They are not covered by `tools/README.md`, `node --test tools/*.test.mjs` or the
CI quick gates. They assume a private campaign layout, a `b8x` checkout, the
recorded Docker image and a parent P00 directory, so they are not expected to run
from a fresh clone.

The active preparation helpers remain in `tools/`:

- `preparation-source-variants.mjs`
- `preparation-transition-oracles.mjs`
- `preparation-native.test.mjs`
- `preparation-persistent.test.mjs`
- `preparation.test.mjs`
