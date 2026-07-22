import sys

def gen_nested(n):
    res = "func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {\n"
    for i in range(1, n):
        res += f"return gopurs_runtime.Func(func(v_{i} gopurs_runtime.Value) gopurs_runtime.Value {{\n"
    res += f"return gopurs_runtime.Int(42)\n"
    for i in range(1, n):
        res += "})\n"
    res += "}"
    return res

def write_file(filename, content):
    with open(filename, "w") as f:
        f.write("package main\n")
        f.write("import \"sync\"\n")
        f.write("import \"gopurs/output/gopurs_runtime\"\n")
        f.write(content)

n = 4000
nested = gen_nested(n)

write_file("closure/main.go", f"""
var foo gopurs_runtime.Value
var once_foo sync.Once

func Get_foo() gopurs_runtime.Value {{
    once_foo.Do(func() {{
        foo = gopurs_runtime.Func({nested})
    }})
    return foo
}}

func main() {{ Get_foo() }}
""")

write_file("named/main.go", f"""
var foo gopurs_runtime.Value
var once_foo sync.Once

func init_foo() {{
    foo = gopurs_runtime.Func({nested})
}}

func Get_foo() gopurs_runtime.Value {{
    once_foo.Do(init_foo)
    return foo
}}

func main() {{ Get_foo() }}
""")

write_file("direct/main.go", f"""
func Get_foo() gopurs_runtime.Value {{
    return gopurs_runtime.Func({nested})
}}

func main() {{ Get_foo() }}
""")

