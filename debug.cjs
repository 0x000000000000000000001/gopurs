const fs = require('fs');
const { decodeModule } = require('./output/PureScript.Backend.Optimizer.CoreFn.Json/index.js');
const { jsonParser } = require('./output/Data.Argonaut.Parser/index.js');

const content = fs.readFileSync('output/Record.Builder/corefn.json', 'utf8');
const json = jsonParser(content);

console.log("JSON Parser result:", json);
try {
  const decoded = decodeModule(json.value0);
  console.log("Decoded:", JSON.stringify(decoded, null, 2));
} catch(e) {
  console.log("Exception:", e);
}
