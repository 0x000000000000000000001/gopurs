import { buildReboxFieldIndex } from '../output/Gopurs.ReboxMetadata/index.js';

// Build derived metadata after a fixture has applied its constructor/class overrides.
export const withReboxFields = metadata => ({
    ...metadata,
    reboxFields: buildReboxFieldIndex(metadata.ctorTypes)(metadata.classDeclsFields),
});
