import { findFfiFileImpl as f2 } from '../purescript-backend-optimizer/src/PureScript/Backend/Optimizer/FfiSupport.js';

let res = f2(".go")([])(null)("Control.Extend")(".spago/p/control-6.0.0/src/Control/Extend.purs")();
console.log("Result:", res);
