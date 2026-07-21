import {time_, traceImpl} from "./foreign.js";
const traceWhen = () => bool => a => b => {
  if (bool) { return traceImpl(a)(v => b); }
  return b;
};
const time = () => time_;
const spyWhen = () => bool => a => {
  if (bool) { return traceImpl(a)(v => a); }
  return a;
};
export {spyWhen, time, traceWhen};
export * from "./foreign.js";
