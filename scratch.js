const ast = {
  type: "EffectBind",
  ident: "record'",
  binding: {
    type: "EffectDefer",
    binding: {
      type: "Let",
      ident: "f'",
      value: { type: "Abs" },
      body: {
        type: "EffectDefer",
        binding: {
          type: "Let",
          ident: "a'",
          value: { type: "Literal", value: true },
          body: {
            type: "EffectPure",
            binding: { type: "App" }
          }
        }
      }
    }
  }
};
