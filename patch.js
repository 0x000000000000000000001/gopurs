const fs = require('fs');
const file = 'src/Gopurs/CodeGen.purs';
let code = fs.readFileSync(file, 'utf8');

if (!code.includes("APP ARRAYMAP in Test_Main")) {
    // already removed or not here?
}

// I will just add a JSON stringify of the expr when it matches arrayMap!
// But wait, the expr is a PureScript data structure! I can't JSON stringify it easily.
