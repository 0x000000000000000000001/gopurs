import assert from "node:assert/strict";
import test from "node:test";
import { withSbomSerialNumber } from "./release-sbom.mjs";

const inventory = { bomFormat: "CycloneDX", specVersion: "1.6", version: 1,
  components: [{ type: "library", name: "example", version: "1.0.0" }] };

test("CycloneDX inventory has the serial required by actions/attest", () => {
  // Exact format discriminator from actions/attest src/sbom.ts at
  // 1e69f48acb82d1966a394da916b4c1698aa569d6 (the failed RC5 action).
  const accepted = value => Boolean(value.bomFormat && value.serialNumber && value.specVersion);
  assert.equal(accepted(inventory), false);
  const result = withSbomSerialNumber(inventory);
  assert.equal(accepted(result), true);
  assert.match(result.serialNumber, /^urn:uuid:[a-f0-9]{8}-[a-f0-9]{4}-5[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}$/);
  // Independent Python uuid.uuid5(uuid.NAMESPACE_URL, compact JSON) fixture.
  assert.equal(result.serialNumber, "urn:uuid:3a99af45-df0b-596a-8a9e-7479fe6a1416");
  const { serialNumber, ...unchanged } = result;
  assert.deepEqual(unchanged, inventory);
  assert.equal(inventory.serialNumber, undefined);
});

test("inventory identity is reproducible and changes with dependency contents", () => {
  const result = withSbomSerialNumber(inventory);
  assert.deepEqual(withSbomSerialNumber(structuredClone(inventory)), result);
  assert.deepEqual(withSbomSerialNumber(result), result);
  const changed = structuredClone(inventory);
  changed.components[0].version = "1.0.1";
  assert.notEqual(withSbomSerialNumber(changed).serialNumber, result.serialNumber);
});
