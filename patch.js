const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');
code = code.replace(
  'unwrapFunc binds =',
  `unwrapFunc binds =
      let _ = unsafePerformEffect $ traverse_ (\\(Tuple (Ident name) val) -> if name == "altAppend" then Debug.trace ("unwrapFunc: altAppend " <> show (isJust (extractUncurriedAbs val))) \\_ -> pure unit else pure unit) binds in`
);
fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
