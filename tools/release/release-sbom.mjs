import { createHash } from "node:crypto";

// UUIDv5 in the URL namespace, using the complete inventory as its name.
// Identical inventories retain their identity; changed dependencies get a new
// serial even when the release version, source commit and target are unchanged.
export function withSbomSerialNumber({ serialNumber: _previous, ...inventory }) {
  const namespace = Buffer.from("6ba7b8119dad11d180b400c04fd430c8", "hex");
  const bytes = createHash("sha1").update(namespace).update(JSON.stringify(inventory)).digest().subarray(0, 16);
  bytes[6] = (bytes[6] & 0x0f) | 0x50;
  bytes[8] = (bytes[8] & 0x3f) | 0x80;
  const hex = bytes.toString("hex");
  return { ...inventory, serialNumber: `urn:uuid:${hex.slice(0, 8)}-${hex.slice(8, 12)}-${hex.slice(12, 16)}-${hex.slice(16, 20)}-${hex.slice(20)}` };
}
