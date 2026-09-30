import { empty as emptyMap } from '../output/Data.Map/index.js';
import { buildReboxFieldIndex } from '../output/Gopurs.ReboxMetadata/index.js';

// Fixtures have no local FFI workers unless specified. Build derived metadata
// after applying their constructor/class overrides.
export const withReboxFields = metadata => ({
    ffiFunctions: emptyMap,
    ...metadata,
    reboxFields: buildReboxFieldIndex(metadata.ctorTypes)(metadata.classDeclsFields),
});
