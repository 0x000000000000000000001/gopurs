// Handwritten FFI for Gopurs.Cache.Wire; keep in sync with the .purs foreign imports.
import { createHash } from "node:crypto";

export const hash = text => createHash("sha256").update(text, "utf8").digest("hex");
export const compareStrings = a => b => a < b ? -1 : a > b ? 1 : 0;
export const hex = text => {
  for (let i = 0; i < text.length; i++) {
    const c = text.charCodeAt(i);
    if (!((c >= 48 && c <= 57) || (c >= 97 && c <= 102))) return false;
  }
  return true;
};
export const stringUnits = text => {
  let result = "";
  for (let i = 0; i < text.length; i++) result += text.charCodeAt(i).toString(16).padStart(4, "0");
  return result;
};
export const unitsString = text => {
  let result = "";
  for (let i = 0; i < text.length; i += 4) result += String.fromCharCode(parseInt(text.slice(i, i + 4), 16));
  return result;
};
export const numberBits = value => {
  const bytes = Buffer.allocUnsafe(8);
  bytes.writeDoubleBE(value);
  return bytes.toString("hex");
};
export const bitsNumber = text => Buffer.from(text, "hex").readDoubleBE();
