import assert from "node:assert/strict";
import { spawnSync } from "node:child_process";
import { copyFileSync, mkdirSync, mkdtempSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { test } from "node:test";
import * as Wire from "../output/Gopurs.Cache.Wire/index.js";
import * as Maybe from "../output/Data.Maybe/index.js";
import * as Ord from "../output/Data.Ord/index.js";
import * as Scalars from "../src/Gopurs/Cache/Wire.js";
import * as CoreFn from "../output/PureScript.Backend.Optimizer.CoreFn/index.js";

const ctor = (name, ...args) => ["ctor", name, args];
const array = values => ["array", values];
const string = value => ["string", Scalars.stringUnits(value)];
const int = value => ["int", String(value)];
const ref = index => ctor("TypeRef", int(index));
// Independent, tiny encoder for test vectors. Production decoding must accept
// these portable bytes; production encoding must reproduce their exact order.
const artifact = (value, types = []) => {
  const nodes = [], indices = new Map(), table = Array(types.length), active = new Set();
  const graphRef = int;
  const intern = value => {
    const text = JSON.stringify(value);
    if (!indices.has(text)) { indices.set(text, nodes.length); nodes.push(value); }
    return graphRef(indices.get(text));
  };
  const encode = value => {
    if (value[0] === "ctor") {
      if (value[1] === "TypeRef") {
        const index = Number(value[2][0][1]);
        if (types[index] && !table[index] && !active.has(index)) {
          active.add(index); table[index] = encode(types[index]); active.delete(index);
        }
      }
      return intern(ctor(value[1], ...value[2].map(encode)));
    }
    if (value[0] === "array") return intern(array(value[1].map(encode)));
    if (value[0] === "record") return intern(["record", value[1].map(([name, v]) => [name, encode(v)])]);
    return intern(value);
  };
  const root = encode(value);
  // Keep invalid/unused type-table entries in negative fixtures too.
  for (let i = 0; i < types.length; i++) if (!table[i]) table[i] = encode(types[i]);
  return ctor("WithGraph", array(nodes), ctor("WithTypes", array(table), root));
};
function roundTrip(codec, value, types = []) {
  const text = JSON.stringify(artifact(value, types));
  const decoded = Wire.read(codec)(text);
  assert.ok(decoded instanceof Maybe.Just, text);
  assert.equal(Wire.write(codec)(decoded.value0)(), text);
}
function reject(codec, value) {
  assert.equal(Wire.read(codec)(JSON.stringify(artifact(value))), Maybe.Nothing.value);
}

test("portable codec preserves scalar bits and UTF-16 units", () => {
  for (const value of ["", "abc", "💻", "\ud800x\udfff", "\ud800\udfff", "\ufffd", "\u0000"]) roundTrip(Wire.codecString, string(value));
  for (const bits of ["0000000000000000", "8000000000000000", "7ff0000000000000", "fff0000000000000", "7ff8000000000042", "3ff0000000000000"]) roundTrip(Wire.codecNumber, ["number", bits]);
  for (const value of [0, 1, -1, 2147483647, -2147483648]) roundTrip(Wire.codecInt, int(value));
  for (const invalid of [["int", "01"], ["int", "-0"], ["int", "9223372036854775808"], ["int", 1]]) reject(Wire.codecInt, invalid);
  for (const invalid of [["number", "x"], ["number", "7FF0000000000000"]]) reject(Wire.codecNumber, invalid);
  for (const invalid of [["string", "d80"], ["string", "D800"]]) reject(Wire.codecString, invalid);
});

test("scalar hex validation agrees on every UTF-16 code unit", () => {
  const alphabet = "0123456789abcdef";
  for (let unit = 0; unit < 65536; unit++) {
    const text = String.fromCharCode(unit);
    assert.equal(Scalars.hex(text), alphabet.includes(text), `code unit ${unit}`);
  }
  assert.ok(Scalars.hex(""));
  const long = alphabet.repeat(4096);
  assert.ok(Scalars.hex(long));
  for (const invalid of ["A", "g", "\u0000", "\ud800", "💻", "０"]) {
    assert.equal(Scalars.hex(long + invalid), false);
  }
});

test("business codecs validate constructor names, arities and ordered containers", () => {
  const types = [ctor("Int"), ctor("String"), ctor("Array", ref(1)), ctor("Boolean"),
    ctor("Row", array([ctor("Tuple", string("label"), ref(3))]), ctor("Nothing")),
    ctor("Record", ref(4)), ctor("Func", array([ref(0), ref(2)]), ref(5))];
  roundTrip(Wire.codecExprType, ref(6), types);
  roundTrip(Wire.codecGoType, ctor("TypeFunc", array([ctor("TypeInt64")]), ctor("TypeNativeArray", ctor("TypeString"))));
  const literal = ctor("NeutralExpr", ctor("Lit", ctor("LitInt", int(42))));
  roundTrip(Wire.codecNeutralExpr, ctor("NeutralExpr", ctor("Typed", ref(0), literal)), [ctor("Int")]);
  reject(Wire.codecExprType, ctor("UnknownType"));
  reject(Wire.codecExprType, ctor("Int", int(1)));
  reject(Wire.codecExprType, ctor("Func", array([])));
  reject(Wire.codecGoType, ctor("TypeNativeArray", ctor("Int")));
  const invalidApp = ctor("App", ctor("NeutralExpr", ctor("PrimUndefined")), ctor("NonEmptyArray", array([])));
  reject(Wire.codecNeutralExpr, ctor("NeutralExpr", invalidApp));
  const codec = Wire.codecMap(Ord.ordString)(Wire.keyOrderString)(Wire.codecString)(Wire.codecInt);
  const entries = [ctor("Tuple", string("a"), int(1)), ctor("Tuple", string("b"), int(2))];
  roundTrip(codec, ctor("Map", array(entries)));
  reject(codec, ctor("Map", array([entries[0], entries[0]])));
  reject(codec, ctor("Map", array(entries.toReversed())));
  roundTrip(codec, ctor("Map", array([ctor("Tuple", string("💻"), int(1)), ctor("Tuple", string("\ue000"), int(2))])));
});

test("type tables reject missing, forward and cyclic references", () => {
  for (const [nodes, root] of [
    [[], ref(0)], [[ctor("Array", ref(0))], ref(0)],
    [[ctor("Array", ref(1)), ctor("Int")], ref(0)],
    [[ctor("Unknown")], ref(0)], [[ctor("Int", int(1))], ref(0)],
    [[ctor("Int")], ref(-1)], [[ctor("Int")], ref(1)],
  ]) assert.equal(Wire.read(Wire.codecExprType)(JSON.stringify(artifact(root, nodes))), Maybe.Nothing.value);
});

test("type indices depend on structure and encounter order, not object sharing or prior artifacts", () => {
  const shared = new CoreFn.Array(CoreFn.Int.value);
  const equal = new CoreFn.Array(CoreFn.Int.value);
  const codec = Wire.codecArray(Wire.codecExprType);
  const expected = JSON.stringify(artifact(array([ref(1), ref(1), ref(1)]), [ctor("Int"), ctor("Array", ref(0))]));
  assert.equal(Wire.write(codec)([shared, shared, equal])(), expected);
  assert.equal(Wire.write(codec)([equal, shared, equal])(), expected);
  Wire.write(codec)([CoreFn.String.value, shared])();
  assert.equal(Wire.write(codec)([shared, equal, shared])(), expected);
});

test("nested equal type graphs retain canonical indices across sharing patterns", () => {
  const codec = Wire.codecArray(Wire.codecExprType);
  const values = [], roots = [], types = [ctor("Int")];
  let child = CoreFn.Int.value;
  for (let depth = 0; depth < 9; depth++) {
    const shared = new CoreFn.Func([child, child], child);
    const equal = new CoreFn.Func([child, child], child);
    types.push(ctor("Func", array([ref(depth), ref(depth)]), ref(depth)));
    roots.push(ref(depth + 1), ref(depth + 1), ref(depth + 1));
    values.push(shared, equal, shared);
    child = shared;
  }
  const expected = JSON.stringify(artifact(array(roots), types));
  assert.equal(Wire.write(codec)(values)(), expected);
  const decoded = Wire.read(codec)(expected);
  assert.ok(decoded instanceof Maybe.Just);
  assert.equal(Wire.write(codec)(decoded.value0)(), expected);
  // A prior artifact may encounter the same physical types in another order.
  Wire.write(codec)([CoreFn.String.value, child, ...values.toReversed()])();
  assert.equal(Wire.write(codec)(values)(), expected);
});

test("object graphs reject forward, cyclic and missing references before business decoding", () => {
  const gr = int;
  for (const [nodes, root] of [
    [[], gr(0)], [[ctor("Array", gr(0))], gr(0)],
    [[ctor("Array", gr(1)), int(1)], gr(0)], [[int(1)], gr(-1)],
  ]) assert.equal(Wire.read(Wire.codecInt)(JSON.stringify(ctor("WithGraph", array(nodes), ctor("WithTypes", array([]), root)))), Maybe.Nothing.value);
});

test("shared backend expressions are decoded once per graph node", () => {
  // The depth-1 assertion fails before attempting the larger graph if sharing
  // regresses. At depth 24, 100 wire nodes represent over 16 million leaves.
  for (const depth of [1, 24]) {
    const nodes = [int(1), ctor("LitInt", int(0)), ctor("Lit", int(1)), ctor("NeutralExpr", int(2))];
    const add = value => { nodes.push(value); return nodes.length - 1; };
    let root = 3;
    for (let i = 0; i < depth; i++) {
      const children = add(array([int(root), int(root)]));
      const literal = add(ctor("LitArray", int(children)));
      const syntax = add(ctor("Lit", int(literal)));
      root = add(ctor("NeutralExpr", int(syntax)));
    }
    const decoded = Wire.read(Wire.codecNeutralExpr)(JSON.stringify(ctor("WithGraph", array(nodes), ctor("WithTypes", array([]), int(root)))));
    assert.ok(decoded instanceof Maybe.Just);
    let value = decoded.value0;
    for (let i = 0; i < depth; i++) {
      const children = value.value0.value0;
      assert.equal(children.length, 2);
      assert.equal(children[0], children[1], `shared subtree duplicated at depth ${depth - i}`);
      value = children[0];
    }
    assert.equal(value.value0.value0, 1);
  }
});

test("the actual Go scalar bridge agrees with JS on surrogate and IEEE-754 boundaries", t => {
  const directory = mkdtempSync(join(tmpdir(), "gopurs-wire-"));
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  copyFileSync(new URL("../src/Gopurs/Cache/Wire.go", import.meta.url), join(directory, "wire.go"));
  writeFileSync(join(directory, "go.mod"), "module wire.test\n\ngo 1.22\n");
  const units = ["", "0061d8000062dfff", "d83ddcbb", "0000ffff", "d800dfff"];
  const bits = ["8000000000000000", "7ff0000000000000", "fff0000000000000", "7ff8000000000042"];
  writeFileSync(join(directory, "wire_test.go"), `package Gopurs_Cache_Wire
import (
  "fmt"
  "strings"
  "testing"
)
func TestWire(t *testing.T) {
  for _, v := range []string{${units.map(JSON.stringify).join(",")}} {
    if StringUnits(UnitsString(v)) != v { t.Fatal(v) }
  }
  for _, v := range []string{${bits.map(JSON.stringify).join(",")}} {
    if NumberBits(BitsNumber(v)) != v { t.Fatal(v) }
  }
  if Hash("portable-wire-v1") != ${JSON.stringify(Scalars.hash("portable-wire-v1"))} { t.Fatal("digest differs") }
  if CompareStrings("💻", "\ue000") >= 0 { t.Fatal("UTF-8 order leaked into wire map") }
}
func TestHex(t *testing.T) {
  const alphabet = "0123456789abcdef"
  for unit := 0; unit < 65536; unit++ {
    text := UnitsString(fmt.Sprintf("%04x", unit))
    if Hex(text) != strings.ContainsRune(alphabet, rune(unit)) { t.Fatalf("code unit %d", unit) }
  }
  if !Hex("") { t.Fatal("empty hex") }
  long := strings.Repeat(alphabet, 4096)
  if !Hex(long) { t.Fatal("long hex") }
  for _, invalid := range []string{"A", "g", string([]byte{0}), UnitsString("d800"), "💻", "０", string([]byte{255})} {
    if Hex(long + invalid) { t.Fatal("invalid hex suffix") }
  }
}
`);
  const result = spawnSync("go", ["test", "-race", "-count=1", "./..."], { cwd: directory, env: { ...process.env, GOWORK: "off" }, encoding: "utf8", timeout: 120000 });
  assert.equal(result.status, 0, result.stdout + result.stderr);
});

test("the native type memo retains opaque owners and isolates artifacts", t => {
  const directory = mkdtempSync(join(tmpdir(), "gopurs-type-memo-"));
  t.after(() => rmSync(directory, { recursive: true, force: true }));
  mkdirSync(join(directory, "gopurs_runtime"));
  copyFileSync(new URL("../runtime/runtime.go", import.meta.url), join(directory, "gopurs_runtime/runtime.go"));
  copyFileSync(new URL("../src/Gopurs/Cache/TypeMemo.go", import.meta.url), join(directory, "memo.go"));
  writeFileSync(join(directory, "go.mod"), "module gopurs/output\n\ngo 1.22\n");
  writeFileSync(join(directory, "memo_test.go"), `package Gopurs_Cache_TypeMemo
import (
  "gopurs/output/gopurs_runtime"
  "runtime"
  "testing"
  "unsafe"
)
func TestIdentity(t *testing.T) {
  type payload struct { tag uint32; value int64 }
  key := func() gopurs_runtime.Value {
    return gopurs_runtime.Value{Type: gopurs_runtime.TypeConstructor, IntVal: 42, UnsafePtr: unsafe.Pointer(&payload{1, 123})}
  }
  m, a, b := Create(nil), key(), key()
  if Lookup(m, a, nil) != -1 { t.Fatal("unexpected entry") }
  Insert(m, a, 7, nil)
  if Lookup(m, a, nil) != 7 || Lookup(m, b, nil) != -1 { t.Fatal("identity conflated") }
  if Lookup(Create(nil), a, nil) != -1 { t.Fatal("artifact leaked") }
  for tag := int64(1); tag <= 2; tag++ {
    v := gopurs_runtime.Value{Type: gopurs_runtime.TypeConstructor, IntVal: tag}
    Insert(m, v, tag, nil)
    if Lookup(m, v, nil) != tag { t.Fatal("nullary tag conflated") }
  }
  a = gopurs_runtime.Value{}
  runtime.GC()
  retained := false
  for v, index := range m.AnyVal().(memo) {
    if index == 7 {
      retained = (*payload)(v.UnsafePtr).value == 123
      if Lookup(m, v, nil) != 7 { t.Fatal("entry lost after GC") }
    }
  }
  if !retained { t.Fatal("owner not retained") }
}
`);
  const result = spawnSync("go", ["test", "-race", "-count=1", "./..."], { cwd: directory, env: { ...process.env, GOWORK: "off" }, encoding: "utf8", timeout: 120000 });
  assert.equal(result.status, 0, result.stdout + result.stderr);
});
