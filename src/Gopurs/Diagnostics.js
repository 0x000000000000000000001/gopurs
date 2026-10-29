// Handwritten FFI for Gopurs.Diagnostics; keep in sync with the .purs foreign imports.
let observed = false;
export const mark = () => { observed = true; };
export const recorded = () => observed;
