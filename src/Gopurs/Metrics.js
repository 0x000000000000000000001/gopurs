// Like altbak's Bench clock, use monotonic elapsed time, never Date.now().
import { performance } from "node:perf_hooks";

export const now = () => performance.now();

// No-op: seul le backend Go échantillonne le profil mémoire.
export const setMemProfileRate = (rate) => () => {};

