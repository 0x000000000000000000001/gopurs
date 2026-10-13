static GOPURS_DIAGNOSTIC: std::sync::atomic::AtomicBool = std::sync::atomic::AtomicBool::new(false);

pub fn Gopurs_Diagnostics_mark() -> Value {
    Value::Func1(Func1::Static(|_| {
        GOPURS_DIAGNOSTIC.store(true, std::sync::atomic::Ordering::Relaxed);
        Value::Unit
    }))
}

pub fn Gopurs_Diagnostics_recorded() -> Value {
    Value::Func1(Func1::Static(|_| Value::Bool(GOPURS_DIAGNOSTIC.load(std::sync::atomic::Ordering::Relaxed))))
}
