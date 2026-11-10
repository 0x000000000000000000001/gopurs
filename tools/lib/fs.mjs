import { mkdirSync, readFileSync, renameSync, writeFileSync } from "node:fs";

/** Read and parse one UTF-8 JSON document. */
export function readJson(file) {
  return JSON.parse(readFileSync(file, "utf8"));
}

/** Create a directory and any missing parents. */
export function ensureDirectory(directory) {
  mkdirSync(directory, { recursive: true });
  return directory;
}

// Write the complete document beside its destination before replacing it, so
// an interrupted report never leaves a truncated JSON file behind.
export function writeJsonAtomic(file, value) {
  const temporary = `${file}.tmp`;
  writeFileSync(temporary, JSON.stringify(value, null, 2) + "\n");
  renameSync(temporary, file);
}
