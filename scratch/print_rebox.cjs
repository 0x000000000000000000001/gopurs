const fs = require('fs');
let code = fs.readFileSync('../../altbak.pub/output/purescript/Control_Comonad_Env_Class.go', 'utf-8');
let match = code.match(/func Rebox_Control_Comonad_Env_Class_.*?\{[\s\S]*?\}/);
if (match) {
  console.log(match[0]);
} else {
  console.log("Rebox not found");
}
