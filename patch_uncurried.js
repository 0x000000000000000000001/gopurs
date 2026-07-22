import fs from 'fs';

let code = fs.readFileSync('src/Gopurs/CodeGen.purs', 'utf8');

code = code.replace(
  /isEffectNode :: forall a\. BackendSyntax a \-\> Boolean\nisEffectNode = case _ of\n  EffectBind _ _ _ _ \-\> true\n  EffectPure _ \-\> true\n  EffectDefer _ \-\> true\n  PrimEffect _ \-\> true\n  _ \-\> false/,
  `isEffectNode :: forall a. BackendSyntax a -> Boolean
isEffectNode = case _ of
  EffectBind _ _ _ _ -> true
  EffectPure _ -> true
  EffectDefer _ -> true
  PrimEffect _ -> true
  UncurriedEffectApp _ _ -> true
  _ -> false`
);

fs.writeFileSync('src/Gopurs/CodeGen.purs', code);
