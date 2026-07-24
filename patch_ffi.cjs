const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/FfiSupport.js', 'utf8');

code = code.replace(/\.PtrVal/g, '.PtrVal()');

fs.writeFileSync('src/Gopurs/FfiSupport.js', code);
