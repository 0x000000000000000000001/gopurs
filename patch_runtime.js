const fs = require('fs');
let code = fs.readFileSync('tests/runner/output/gopurs_runtime/runtime.go', 'utf8');

code = code.replace(
    'return Apply(Apply(f, arg1), arg2)',
    'fmt.Printf("Apply2 fallback! f.Type=%d, f.UnsafePtr=%v\\n", f.Type, f.UnsafePtr)\n\tres1 := Apply(f, arg1)\n\tfmt.Printf("Apply2 fallback res1! res1.Type=%d, res1.UnsafePtr=%v\\n", res1.Type, res1.UnsafePtr)\n\treturn Apply(res1, arg2)'
);

fs.writeFileSync('tests/runner/output/gopurs_runtime/runtime.go', code);
