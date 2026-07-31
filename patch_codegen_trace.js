const fs = require('fs');
let code = fs.readFileSync('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
  /let\s+concreteArr = Set\.toUnfoldable concretes :: Array ExprType/,
  `let
                 _ = if qual == "Data.Void.absurd" then unsafePerformEffect (Console.log ("CONCRETES FOR ABSURD: " <> show (Array.length (Set.toUnfoldable concretes :: Array ExprType)))) else unit
                 concreteArr = Set.toUnfoldable concretes :: Array ExprType`
);

fs.writeFileSync('/Users/0x1/Documents/htdocs/gopurs/src/Gopurs/CodeGen.purs', code);
