const nativeOnly = () => {
  throw new Error("This example requires the gopurs Go backend.");
};
export const startImpl = nativeOnly;
export const waitImpl = nativeOnly;
export const stopImpl = nativeOnly;
