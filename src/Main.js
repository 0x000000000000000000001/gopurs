export const takeMonomorphizedModules = (prepared) => () => {
    const list = prepared.monomorphizedModules;
    // Replace with a dummy empty List (Nil)
    prepared.monomorphizedModules = { value0: "Nil" };
    return list;
};
