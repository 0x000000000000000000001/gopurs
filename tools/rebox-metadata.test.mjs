import assert from 'node:assert/strict';
import test from 'node:test';
import { empty, insert, lookup, toUnfoldable } from '../output/Data.Map/index.js';
import { Just, Nothing } from '../output/Data.Maybe/index.js';
import { ordString } from '../output/Data.Ord/index.js';
import { unfoldableArray } from '../output/Data.Unfoldable/index.js';
import { sanitizeName } from '../output/Gopurs.GoAst/index.js';
import { buildReboxFieldIndex } from '../output/Gopurs.ReboxMetadata/index.js';
import * as Core from '../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { withReboxFields } from './codegen-metadata.mjs';

const map = entries => entries.reduce((result, [key, value]) =>
    insert(ordString)(key)(value)(result), empty);
const fields = type => ({ vars: ['a'], fields: [type] });
const classFields = type => ({ vars: ['a'], fields: [{ name: 'value', type }] });
const get = (index, key) => lookup(ordString)(key)(index);
const aliases = suffix => [`Constructor_${suffix}`, `Data_${suffix}`];
const assertAliases = (index, suffix, expected) => {
    for (const key of aliases(suffix)) assert.deepEqual(get(index, key), new Just(expected), key);
};

// Reference behavior of the previous per-lookup scan, including first-match order.
function scan(constructors, classes, name) {
    const matches = pair => {
        const parts = pair.value0.split('.');
        if (parts.length < 2) return false;
        const constructor = sanitizeName(parts.pop());
        return aliases(`${parts.join('_')}_${constructor}`).includes(name);
    };
    const constructor = toUnfoldable(unfoldableArray)(constructors).find(matches);
    if (constructor) return new Just(constructor.value1);
    const klass = toUnfoldable(unfoldableArray)(classes).find(matches);
    return klass ? new Just({ vars: klass.value1.vars, fields: klass.value1.fields.map(field => field.type) })
        : Nothing.value;
}

test('rebox aliases retain constructor variables and field order', () => {
    const info = { vars: ['z', 'a'], fields: [Core.String.value, new Core.TypeVar('z'), Core.Int.value] };
    const index = buildReboxFieldIndex(map([['Data_Tuple.Tuple', info]]))(empty);
    assertAliases(index, 'Data_Tuple_Tuple', info);
});

test('rebox class aliases retain declared field order without sorting names', () => {
    const info = { vars: ['z', 'a'], fields: [
        { name: 'zLast', type: new Core.TypeVar('z') },
        { name: 'aFirst', type: Core.String.value },
        { name: 'middle', type: new Core.TypeVar('a') },
    ] };
    const index = buildReboxFieldIndex(empty)(map([['Fixture.Nested.Dictionary', info]]));
    assertAliases(index, 'Fixture_Nested_Dictionary', {
        vars: ['z', 'a'], fields: [new Core.TypeVar('z'), Core.String.value, new Core.TypeVar('a')],
    });
});

test('rebox constructors take priority over classes even when the class key sorts first', () => {
    const constructor = fields(Core.Int.value);
    const constructors = map([['A_B.C', constructor]]);
    const classes = map([['A.B.C', classFields(Core.String.value)]]);
    assertAliases(buildReboxFieldIndex(constructors)(classes), 'A_B_C', constructor);
});

for (const kind of ['constructor', 'class']) {
    for (const [label, first, second, suffix] of [
        ['module separators', 'A.B.C', 'A_B.C', 'A_B_C'],
        ['sanitized constructor names', "Fixture.A'", 'Fixture.A_prime_', 'Fixture_A_prime_'],
    ]) {
        test(`rebox ${kind} collisions preserve the first Map key for ${label}`, () => {
            const info = kind === 'constructor' ? fields : classFields;
            // Insertion order is deliberately opposite to Map key order.
            const entries = map([[second, info(Core.String.value)], [first, info(Core.Int.value)]]);
            const index = kind === 'constructor'
                ? buildReboxFieldIndex(entries)(empty)
                : buildReboxFieldIndex(empty)(entries);
            assertAliases(index, suffix, fields(Core.Int.value));
        });
    }
}

test('rebox indexes ignore keys without a module separator and return Nothing for misses', () => {
    const constructors = map([['', fields(Core.Int.value)], ['Orphan', fields(Core.Int.value)]]);
    const classes = map([['', classFields(Core.String.value)], ['OrphanClass', classFields(Core.String.value)]]);
    const index = buildReboxFieldIndex(constructors)(classes);
    assert.deepEqual(toUnfoldable(unfoldableArray)(index), []);
    for (const key of ['Data_Orphan', 'Constructor_OrphanClass', 'Data_Missing_Value', '']) {
        assert.deepEqual(get(index, key), Nothing.value);
    }
});

test('rebox indexes agree with the previous scan across mixed aliases, collisions and empty segments', () => {
    const constructors = map([
        ['A_B.C', fields(Core.Boolean.value)], ['A.B.C', fields(Core.Int.value)],
        ["Fixture.A'", fields(Core.String.value)], ['Fixture.A_prime_', fields(Core.Int.value)],
        ['.EmptyModule', fields(Core.Int.value)], ['Fixture.', fields(Core.Boolean.value)],
        ['invalid', fields(Core.String.value)],
    ]);
    const classes = map([
        ['A.B.C', classFields(Core.String.value)], ['Dictionary.Nested.Only', classFields(Core.Boolean.value)],
        ['Class.Case+', classFields(Core.Int.value)], ['Class.Case_plus_', classFields(Core.String.value)],
        ['NoSeparator', classFields(Core.String.value)],
    ]);
    const index = buildReboxFieldIndex(constructors)(classes);
    const names = [
        ...['A_B_C', 'Fixture_A_prime_', '_EmptyModule', 'Fixture_X_empty',
            'Dictionary_Nested_Only', 'Class_Case_plus_', 'Missing_Value'].flatMap(aliases),
        'Data_invalid', 'Constructor_NoSeparator', '',
    ];
    for (const name of names) assert.deepEqual(get(index, name), scan(constructors, classes, name), name);
});

test('rebox metadata is independent between compilations and rebuilt after fixture overrides', () => {
    const constructors = map([['Fixture.Value', fields(Core.Int.value)]]);
    const first = withReboxFields({ ctorTypes: constructors, classDeclsFields: empty });
    const second = withReboxFields({ ...first, ctorTypes: empty,
        classDeclsFields: map([['Fixture.Value', classFields(Core.String.value)]]) });
    const third = withReboxFields({ ...second, classDeclsFields: empty });
    assertAliases(first.reboxFields, 'Fixture_Value', fields(Core.Int.value));
    assertAliases(second.reboxFields, 'Fixture_Value', fields(Core.String.value));
    assert.deepEqual(get(third.reboxFields, 'Data_Fixture_Value'), Nothing.value);
    assert.deepEqual(get(first.ctorTypes, 'Fixture.Value'), new Just(fields(Core.Int.value)));
});
