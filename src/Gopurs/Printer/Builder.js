// Handwritten FFI for Gopurs.Printer.Builder; keep in sync with the .purs foreign imports.
// Each rendering owns its buffer. pushImpl returns the same mutable handle;
// toStringImpl publishes an immutable string, independent of later appends.
export const newBuilderImpl = () => ({ parts: [] });
export const pushImpl = builder => chunk => {
  builder.parts.push(chunk);
  return builder;
};
export const toStringImpl = builder => builder.parts.join("");
