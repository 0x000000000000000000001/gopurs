// Handwritten FFI for Gopurs.BuildCache; keep in sync with the .purs foreign imports.
unsafe extern "C" {
    fn GopursBuildCache(content: *const u8, length: usize, session: i32) -> *mut std::ffi::c_char;
    fn GopursFreeFFI(result: *mut std::ffi::c_char);
}

fn gopurs_cache_exchange(request: String, session: i32) -> Value {
    Value::Func1(Func1::Shared(std::rc::Rc::new(move |_| {
        let bytes = purust_string_to_utf8_lossy(&request);
        let response = unsafe {
            let pointer = GopursBuildCache(bytes.as_ptr(), bytes.len(), session);
            let text = std::ffi::CStr::from_ptr(pointer).to_string_lossy().into_owned();
            GopursFreeFFI(pointer);
            text
        };
        Value::String(purust_string_from_utf8(&response))
    })))
}

pub fn Gopurs_BuildCache_exchange(request: String) -> Value { gopurs_cache_exchange(request, 0) }
pub fn Gopurs_BuildCache_sessionExchange(request: String) -> Value { gopurs_cache_exchange(request, 1) }

// Preloading is optional. The helper bridge retains the sequential protocol.
pub fn Gopurs_BuildCache_sessionExchangeOnMiss(request: String, _on_miss: Func1<String, Value>) -> Value { gopurs_cache_exchange(request, 1) }
