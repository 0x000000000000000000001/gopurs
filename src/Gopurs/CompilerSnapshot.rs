// Handwritten FFI for Gopurs.CompilerSnapshot; keep in sync with the .purs foreign imports.
unsafe extern "C" {
    fn GopursFreezeCompiler() -> *mut std::ffi::c_char;
    fn GopursReleaseCompiler();
    fn GopursFreeFFI(result: *mut std::ffi::c_char);
}

pub fn Gopurs_CompilerSnapshot_withSnapshot(run: Value) -> Value {
    Value::Func1(Func1::Shared(std::rc::Rc::new(move |_| {
        unsafe {
            let error = GopursFreezeCompiler();
            if !error.is_null() {
                eprintln!("[gopurs] compiler snapshot: {}", std::ffi::CStr::from_ptr(error).to_string_lossy());
                GopursFreeFFI(error);
                std::process::exit(1);
            }
        }
        run.unwrap_func1()(Value::Unit)
    })))
}

pub fn Gopurs_CompilerSnapshot_release() -> Value {
    Value::Func1(Func1::Static(|_| { unsafe { GopursReleaseCompiler(); } Value::Unit }))
}
