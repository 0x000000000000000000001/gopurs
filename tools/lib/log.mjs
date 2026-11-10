/** Print a scoped step heading. */
export function step(scope, message) {
  console.log(`\n[${scope}] ${message}`);
}

/** Print an informational line. */
export function info(message) {
  console.log(message);
}
