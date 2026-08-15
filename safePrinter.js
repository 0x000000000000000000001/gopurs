function safePrint(obj, depth = 0) {
  if (depth > 2) return "...";
  if (!obj) return "" + obj;
  if (typeof obj === "string" || typeof obj === "number" || typeof obj === "boolean") return "" + obj;
  if (Array.isArray(obj)) return "[" + obj.map(x => safePrint(x, depth + 1)).join(", ") + "]";
  return "{" + Object.keys(obj).map(k => k + ": " + safePrint(obj[k], depth + 1)).join(", ") + "}";
}
module.exports = safePrint;
