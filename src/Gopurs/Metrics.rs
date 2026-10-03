pub fn Gopurs_Metrics_now() -> Value {
    Value::Func1(Func1::Shared(std::rc::Rc::new(|_| {
        static START: std::sync::OnceLock<std::time::Instant> = std::sync::OnceLock::new();
        Value::Number(START.get_or_init(std::time::Instant::now).elapsed().as_secs_f64() * 1000.0)
    })))
}

// This option controls Go's allocation sampler, as in the JavaScript host.
pub fn Gopurs_Metrics_setMemProfileRate(_rate: i64) -> Value {
    Value::Func1(Func1::Static(|_| Value::Unit))
}
