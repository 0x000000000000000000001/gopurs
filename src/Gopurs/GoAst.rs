// Purust flattens the two curried arguments at the native FFI boundary. Retain
// the callback with its cache so a shared closure's address cannot be recycled.
// Per-thread caches also allow callbacks to re-enter without holding a lock.
type GopursNameMemo = std::collections::HashMap<(u8, usize),
    (Func1<String, String>, std::collections::HashMap<String, String>)>;
std::thread_local! {
    static GOPURS_NAMES: std::cell::RefCell<GopursNameMemo> = Default::default();
}

pub fn Gopurs_GoAst_memoizeName(sanitize: Func1<String, String>, name: String) -> String {
    let key = match &sanitize {
        Func1::Static(function) => (0, *function as usize),
        Func1::Shared(function) => (1, std::rc::Rc::as_ptr(function) as *const () as usize),
    };
    let cached = GOPURS_NAMES.with(|cache| cache.borrow().get(&key)
        .and_then(|(_, names)| names.get(&name)).cloned());
    if let Some(result) = cached { return result; }
    let result = sanitize(name.clone());
    GOPURS_NAMES.with(|cache| {
        cache.borrow_mut().entry(key).or_insert_with(|| (sanitize, Default::default()))
            .1.insert(name, result.clone());
    });
    result
}
