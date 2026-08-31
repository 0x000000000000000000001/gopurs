const fs = require('fs');
let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

// We want to trace fn.args and fArgs for lookup
code = code.replace(
  `                                  expectedRetType = case Map.lookup goName moduleArities of`,
  `                                  expectedRetType = unsafePerformEffect (do
                                      when (String.contains (Pattern "lookup") goName) do
                                        Console.log ("lookup fn.args: " <> show fn.args <> " fArgs length: " <> show (Array.length (fromMaybe [] (map _.fArgs (Map.lookup goName moduleArities)))) <> " fArgsGo length: " <> show (Array.length paramsWithTypes) <> " pTypes: " <> show (Array.fromFoldable (Map.keys pTypes)))
                                      pure unit) \`const\` case Map.lookup goName moduleArities of`
);

code = `import Effect.Console as Console
import Data.String (Pattern(Pattern))
import Data.String as String
import Debug as Debug
` + code;

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
