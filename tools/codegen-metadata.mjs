import { Leaf, Node } from "../output/Data.Map.Internal/index.js";
import { fromMap } from "../output/Gopurs.Cache.LookupTable/index.js";
import { empty as emptyMap } from "../output/Data.Map/index.js";
import { map } from "../output/Data.Functor/index.js";
import { functorMap } from "../output/Data.Map.Internal/index.js";
import { Nothing } from "../output/Data.Maybe/index.js";
import { empty as emptySet } from "../output/Data.Set/index.js";
import { buildReboxFieldIndex } from "../output/Gopurs.ReboxMetadata/index.js";

export const withGlobalTypes = (metadata) => ({
  ...metadata,
  globalTypes:
    metadata.globalTypes instanceof Leaf || metadata.globalTypes instanceof Node
      ? fromMap(metadata.globalTypes)
      : metadata.globalTypes,
});

// Fixtures have no local FFI workers unless specified. Build derived metadata
// after applying their constructor/class overrides.
export const withReboxFields = (metadata) => ({
  ffiFunctions: emptyMap,
  sharedRecords: emptySet,
  ...withGlobalTypes(metadata),
  globalFunctions: map(functorMap)(info => ({ arrayFold: Nothing.value, codePointFold: Nothing.value, ...info }))(metadata.globalFunctions ?? emptyMap),
  reboxFields: buildReboxFieldIndex(metadata.ctorTypes)(metadata.classDeclsFields),
});
