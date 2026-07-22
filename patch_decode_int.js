const fs = require('fs');
const file = '../gopurs-backend-optimizer/src/PureScript/Backend/Optimizer/CoreFn/Json.purs';
let code = fs.readFileSync(file, 'utf8');

// Replace `LitInt <$> getField decodeInt obj "value"`
const target = `    "IntLiteral" ->
      LitInt <$> getField decodeInt obj "value"`;

const replacement = `    "IntLiteral" -> do
      num <- getField decodeNumber obj "value"
      let num32 = if num > 2147483647.0 then num - 4294967296.0 else if num < -2147483648.0 then num + 4294967296.0 else num
      case Int.fromNumber num32 of
        Just i -> pure (LitInt i)
        Nothing -> Left (TypeMismatch "Int")`;

code = code.replace(target, replacement);

fs.writeFileSync(file, code);
