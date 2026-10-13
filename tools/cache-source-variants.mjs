import assert from "node:assert/strict";
import { readFileSync, readdirSync } from "node:fs";
import { join } from "node:path";

function readSources(directory, prefix = "") {
  const result = {};
  for (const entry of readdirSync(directory, { withFileTypes: true })) {
    const name = prefix + entry.name, path = join(directory, entry.name);
    if (entry.isDirectory()) Object.assign(result, readSources(path, name + "/"));
    else result[name] = readFileSync(path, "utf8");
  }
  return result;
}

export function sourceVariants(directory) {
  const original = readSources(directory);
  let sources = { ...original };
  let values = [14, 5, 8, 6, 5, 9];
  const states = [];
  const edit = (file, before, after) => {
    assert.equal(sources[file].split(before).length, 2, `${file}: expected one edit target`);
    sources[file] = sources[file].replace(before, after);
  };
  const add = (label, extra = {}) => states.push({ label, sources: { ...sources }, stdout: values.join("\n") + "\nDone\n", ...extra });
  add("initial");
  edit("Cache/Body.purs", "privateStep n = n + 2", "privateStep n = n + 3");
  values[0] = 16;
  add("private-body", { changedModules: ["Cache.Body"] });
  edit("Cache/Body.purs", "value n = privateStep n * 2", "value n = privateStep n * 2 + 1");
  values[0] = 17;
  add("exported-body", { changedModules: ["Cache.Body"] });
  edit("Cache/Body.purs", "@inline export value never", "@inline export value always");
  add("inline-directive", { changedModules: ["Cache.Body"] });
  edit("Cache/Shapes.purs", "data Payload = Empty | Payload Int", "data Payload = Empty | Payload Int String");
  edit("Cache/Shapes.purs", "pack n = Payload n", 'pack n = Payload n "payload"');
  edit("Cache/Shapes.purs", "unpack (Payload n) = n", 'unpack (Payload n label) = if label == "payload" then n + 10 else n');
  values[1] = 15;
  add("constructor-layout", { changedModules: ["Cache.Shapes"] });
  edit("Cache/Shapes.purs", "type Packet = { left :: Int, right :: Int }", "type Packet = { active :: Boolean, left :: Int, right :: Int }");
  edit("Cache/Shapes.purs", "packet n = { left: n, right: 3 }", "packet n = { active: true, left: n, right: 3 }");
  add("record-layout", { changedModules: ["Cache.Shapes"] });
  edit("Cache/Shapes.purs", "right :: Int", "right :: Boolean");
  edit("Cache/Shapes.purs", "right: 3", "right: true");
  edit("Cache/Shapes.purs", "scorePacket p = p.left + p.right", "scorePacket p = if p.right then p.left + 10 else p.left");
  values[2] = 15;
  add("record-field-type", { changedModules: ["Cache.Shapes"] });
  edit("Cache/Dictionary.purs", "  meter :: a -> Int", "  meter :: a -> Int\n  bonus :: a -> Int");
  edit("Cache/Dictionary.purs", "  meter n = n + 1", "  meter n = n + 1\n  bonus _ = 2");
  edit("Cache/Dictionary.purs", "measure value = meter value", "measure value = meter value + bonus value");
  values[3] = 8;
  add("dictionary-layout", { changedModules: ["Cache.Dictionary"] });
  edit("Cache/Calls.purs", "value n = Relay.relay n", "value n = Relay.relay n + if Relay.relay (n > 0) then 7 else 8");
  values[4] = 12;
  add("specialization-added", { sameTast: ["Cache.Poly", "Cache.Relay"], changedPrepared: ["Cache.Poly", "Cache.Relay"] });
  edit("Cache/Calls.purs", "Relay.relay (n > 0)", 'Relay.relay (if n > 0 then "yes" else "no") == "yes"');
  add("specialization-replaced", { sameTast: ["Cache.Poly", "Cache.Relay"], changedPrepared: ["Cache.Poly", "Cache.Relay"] });
  sources["Cache/Calls.purs"] = original["Cache/Calls.purs"];
  values[4] = 5;
  add("specialization-removed", { sameTast: ["Cache.Poly", "Cache.Relay"], changedPrepared: ["Cache.Poly", "Cache.Relay"], restoredPrepared: ["Cache.Poly", "Cache.Relay"] });
  for (const extension of ["purs", "go", "js"]) {
    sources["Cache/Moved." + extension] = sources["Cache/Move." + extension];
    delete sources["Cache/Move." + extension];
  }
  edit("Cache/Moved.purs", "module Cache.Move where", "module Cache.Moved where");
  edit("Main.purs", "import Cache.Move as Move", "import Cache.Moved as Move");
  add("module-renamed", { removedModules: ["Cache.Move"], addedModules: ["Cache.Moved"] });
  for (const extension of ["purs", "go", "js"]) {
    sources["Cache/Move." + extension] = original["Cache/Move." + extension];
    delete sources["Cache/Moved." + extension];
  }
  edit("Main.purs", "import Cache.Moved as Move", "import Cache.Move as Move");
  add("module-restored", { removedModules: ["Cache.Moved"], addedModules: ["Cache.Move"] });
  sources = { ...original };
  values = [14, 5, 8, 6, 5, 9];
  add("all-restored", { restoredAll: true });
  assert.equal(states.length, 14);
  return states;
}
