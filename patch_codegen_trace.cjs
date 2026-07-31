const fs = require('fs');
let code = fs.readFileSync('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
  /expandBind \(Tuple id@\(Ident name\) val\) =/,
  `expandBind (Tuple id@(Ident name) val) =
      let _ = if modNameStrOrig <> "." <> name == "Data.Void.absurd" then unsafePerformEffect (Console.log "EXPANDBIND ABSURD CALLED") else unit in`
);

fs.writeFileSync('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', code);
