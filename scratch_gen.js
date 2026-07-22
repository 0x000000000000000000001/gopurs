let out = "";
for (let arity = 6; arity <= 10; arity++) {
    out += `\tif arity == ${arity} {\n`;
    let spaces = "";
    for (let i = 0; i <= arity; i++) {
        out += "\t\t" + spaces;
        if (i == 0) {
            out += "return gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {\n";
        } else if (i < arity) {
            out += `return gopurs_runtime.Func(func(${String.fromCharCode(96+i)} gopurs_runtime.Value) gopurs_runtime.Value {\n`;
        } else {
            let applyStr = "fn";
            for (let j = 1; j <= arity; j++) {
                applyStr = `gopurs_runtime.Apply(${applyStr}, ${String.fromCharCode(96+j)})`;
            }
            out += `return ${applyStr}\n`;
        }
        spaces += "\t";
    }
    for (let i = arity; i >= 0; i--) {
        spaces = spaces.substring(1);
        out += "\t\t" + spaces + "})\n";
    }
    out += `\t}\n`;
}
require('fs').writeFileSync('out.txt', out);
