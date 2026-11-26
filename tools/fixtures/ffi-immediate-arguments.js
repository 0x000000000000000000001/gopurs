let trace = "";

export const identity = value => {
  trace += "F";
  return () => {
    trace += "E";
    return value;
  };
};

export const snapshot = () => trace;
export const reset = () => { trace = ""; };
