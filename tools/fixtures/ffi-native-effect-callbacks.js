let trace = "";

export const runInt = action => {
  trace += "F";
  return () => {
    trace += "B";
    const value = action();
    trace += "A";
    return value;
  };
};
export const runInt64 = runInt;
export const runValue = runInt;
export const tick = value => () => { trace += "T"; return value; };
export const reset = () => { trace = ""; };
export const snapshot = () => trace;
