    flattenForeignType (ADT ["Data", "Function", "Uncurried", "Fn2"] [a, b, ret]) =
      let rest = flattenForeignType ret
      in { args: [a, b] <> rest.args, ret: rest.ret, isEffect: rest.isEffect }
