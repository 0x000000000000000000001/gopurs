/** Print a CLI usage message and exit successfully. */
export function showUsage(text) {
  console.log(text);
  process.exit(0);
}

/** Record a CLI failure without hiding the original error. */
export function fail(message, code = 1) {
  console.error(message);
  process.exitCode = code;
}
