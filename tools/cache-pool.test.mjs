import assert from 'node:assert/strict';
import test from 'node:test';
import { create, finish, intern, internRef } from '../src/Gopurs/Cache/Pool.js';

test('pool effects assign indices when executed and keep insertion order', () => {
    const pool = create();
    const referenceNode = '["unit"]';
    const numericNode = '["boolean",true]';
    const reference = internRef(pool)(referenceNode);
    const numeric = intern(pool)(numericNode);

    assert.equal(finish(pool)(), '');
    assert.equal(numeric(), 0);
    assert.equal(reference(), '["int","1"]');
    assert.equal(reference(), '["int","1"]');
    assert.equal(numeric(), 0);
    assert.equal(finish(pool)(), `${numericNode},${referenceNode}`);
});

test('numeric indices and references share one namespace without duplicate nodes', () => {
    const pool = create();
    const nodes = Array.from({ length: 12 }, (_, index) => `["int","${index}"]`);
    for (const [index, node] of nodes.entries()) assert.equal(intern(pool)(node)(), index);

    assert.equal(internRef(pool)(nodes[11])(), '["int","11"]');
    assert.equal(internRef(pool)(nodes[0])(), '["int","0"]');
    const next = '["string","0061"]';
    assert.equal(internRef(pool)(next)(), '["int","12"]');
    assert.equal(intern(pool)(next)(), 12);
    assert.equal(internRef(pool)(nodes[11])(), '["int","11"]');
    assert.equal(finish(pool)(), [...nodes, next].join(','));
});

test('separate artifact pools assign independent references', () => {
    const first = create();
    const second = create();
    const shared = '["unit"]';
    const prefix = '["boolean",false]';

    assert.equal(internRef(first)(shared)(), '["int","0"]');
    assert.equal(intern(second)(prefix)(), 0);
    assert.equal(internRef(second)(shared)(), '["int","1"]');
    assert.equal(internRef(first)(shared)(), '["int","0"]');
    assert.equal(finish(first)(), shared);
    assert.equal(finish(second)(), `${prefix},${shared}`);
});
