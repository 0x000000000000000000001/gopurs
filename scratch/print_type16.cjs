const fs = require('fs');
const json = JSON.parse(fs.readFileSync('../../altbak.pub/output/Control.Comonad.Env.Class/corefn.json', 'utf-8'));

function printType(types, id, depth=0) {
  const t = types[id];
  if (!t) return "UNKNOWN";
  const indent = "  ".repeat(depth);
  console.log(indent + JSON.stringify(t));
  if (t.type === "TypeApp") {
    printType(types, t.fn, depth+1);
    printType(types, t.arg, depth+1);
  } else if (t.type === "ForAll") {
    printType(types, t.body, depth+1);
  } else if (t.type === "ConstrainedType") {
    printType(types, t.body, depth+1);
  }
}

printType(json.typeTable, 16);
