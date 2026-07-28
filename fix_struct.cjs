const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
  /fieldsStr = Array\.mapWithIndex \\\(\\i ty -> "V" <> show i <> " " <> goTypeToStr ty\\\) goFieldTypes/,
  `fieldsStr = Array.cons "Rc uint32" (Array.mapWithIndex (\\i ty -> "V" <> show i <> " " <> goTypeToStr ty) goFieldTypes)`
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
