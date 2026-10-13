let observed = false;
export const mark = () => { observed = true; };
export const recorded = () => observed;
