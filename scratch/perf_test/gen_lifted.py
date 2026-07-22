import sys

def write_file(filename, n_total, chunk_size):
    with open(filename, "w") as f:
        f.write("package main\n")
        f.write("import \"gopurs/output/gopurs_runtime\"\n")
        
        chunks = n_total // chunk_size
        
        for c in range(chunks):
            f.write(f"func __helper_{c}() gopurs_runtime.Value {{\n")
            if c == chunks - 1:
                f.write("return gopurs_runtime.Int(42)\n")
            else:
                res = "return "
                for i in range(chunk_size):
                    res += "gopurs_runtime.Func(func(v gopurs_runtime.Value) gopurs_runtime.Value {\nreturn "
                res += f"__helper_{c+1}()\n"
                for i in range(chunk_size):
                    res += "})\n"
                f.write(res)
            f.write("}\n")
            
        f.write(f"""
func Get_foo() gopurs_runtime.Value {{
    return __helper_0()
}}
func main() {{ Get_foo() }}
""")

write_file("lifted/main.go", 4000, 100)
