const fs = require('fs');
let code = fs.readFileSync('../altbak.pub/output/Test.RBTree/Test_RBTree.go', 'utf8');

code = code.replace(
  'var Call_main func() gopurs_runtime.Value = func() func() gopurs_runtime.Value {',
  'func printTreeDepth(v gopurs_runtime.Value) int { if v.UnsafePtr == nil { return 0 }; t := (*Constructor_T)(v.UnsafePtr); left := printTreeDepth(t.V1); right := printTreeDepth(t.V3); if left > right { return left + 1 } else { return right + 1 } }\n\nvar Call_main func() gopurs_runtime.Value = func() func() gopurs_runtime.Value {'
);

code = code.replace(
  '	fmt.Printf("%v\\n", __t61)\n',
  '	fmt.Printf("%v\\n", __t61)\n	fmt.Printf("Depth: %v\\n", printTreeDepth(__t61))\n'
);

fs.writeFileSync('../altbak.pub/output/Test.RBTree/Test_RBTree.go', code);
