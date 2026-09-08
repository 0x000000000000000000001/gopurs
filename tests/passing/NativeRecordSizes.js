export const opaqueValue = () => "opaque";

export const nativeShape = record => {
  const keys = Object.keys(record).sort();
  return `${keys.length <= 5 ? "compact" : "generic"}${keys.length}:${keys.join("|")}`;
};

const render = value => {
  if (typeof value === "string") return `s:${value}`;
  if (typeof value === "boolean") return `b:${value}`;
  if (typeof value === "number") return `${Number.isInteger(value) ? "i" : "n"}:${value}`;
  if (Array.isArray(value)) return `[${value.map(render).join(",")}]`;
  return `{${Object.keys(value).sort().map(key => `${key}=${render(value[key])}`).join(",")}}`;
};

export const nativeMap = record => render(record);
