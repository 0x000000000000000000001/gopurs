// Handwritten FFI for Main; keep in sync with the .purs foreign imports.
// Let Node drain pending stderr writes before leaving with a failing status.
export const exitFailure = () => {
  process.exitCode = 1;
};
