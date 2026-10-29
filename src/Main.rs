// Handwritten FFI for Main; keep in sync with the .purs foreign imports.
pub fn Main_exitFailure() -> purust_core::Value {
    purust_core::Value::Func1(purust_core::Func1::Shared(std::rc::Rc::new(|_| std::process::exit(1))))
}
