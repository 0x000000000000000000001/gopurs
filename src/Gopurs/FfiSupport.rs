// The existing Go parser is linked into this executable through a small C ABI.
// Both compiler hosts therefore use exactly the same Go syntax/type handling.
unsafe extern "C" {
    fn GopursParseFFI(content: *const u8, content_len: usize,
        prefix: *const u8, prefix_len: usize, prepare: i32, failed: *mut i32) -> *mut std::ffi::c_char;
    fn GopursFreeFFI(result: *mut std::ffi::c_char);
}

fn gopurs_parse_ffi(content: &str, prefix: Option<&str>) -> Value {
    let content = purust_string_to_utf8_lossy(content);
    let prepare = i32::from(prefix.is_some());
    let prefix = purust_string_to_utf8_lossy(prefix.unwrap_or(""));
    let mut failed = 0;
    let response = unsafe {
        let pointer = GopursParseFFI(content.as_ptr(), content.len(), prefix.as_ptr(), prefix.len(), prepare, &mut failed);
        let result = std::ffi::CStr::from_ptr(pointer).to_string_lossy().into_owned();
        GopursFreeFFI(pointer);
        result
    };
    if failed != 0 {
        Purs_Effect_Exception::purust_exception_raise(Purs_Effect_Exception::Effect_Exception_error(
            purust_string_from_utf8(&format!("Go FFI parse failed: {}", response))));
    }
    Value::String(purust_string_from_utf8(&response))
}

pub fn Gopurs_FfiSupport_extractFfiAstImpl(content: String) -> Value {
    Value::Func1(Func1::Shared(std::rc::Rc::new(move |_| gopurs_parse_ffi(&content, None))))
}

pub fn Gopurs_FfiSupport_prepareFfiAstImpl(prefix: String, content: String) -> Value {
    Value::Func1(Func1::Shared(std::rc::Rc::new(move |_| gopurs_parse_ffi(&content, Some(&prefix)))))
}
