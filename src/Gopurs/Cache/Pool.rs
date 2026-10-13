#[derive(Debug)]
pub struct Pool(std::sync::Mutex<(std::collections::HashMap<String, i64>, Vec<String>)>);

pub fn Gopurs_Cache_Pool_create() -> Value {
    Value::Func1(Func1::Static(|_| Value::Class(std::rc::Rc::new(std::rc::Rc::new(
        Pool(std::sync::Mutex::new((std::collections::HashMap::new(), Vec::new())))
    )))))
}

pub fn Gopurs_Cache_Pool_intern(pool: std::rc::Rc<Pool>, text: String) -> Value {
    Value::Func1(Func1::Shared(std::rc::Rc::new(move |_| {
        let mut state = pool.0.lock().unwrap();
        if let Some(index) = state.0.get(&text) { return Value::Int(*index); }
        let index = state.1.len() as i64;
        state.0.insert(text.clone(), index);
        state.1.push(text.clone());
        Value::Int(index)
    })))
}

pub fn Gopurs_Cache_Pool_finish(pool: std::rc::Rc<Pool>) -> Value {
    Value::Func1(Func1::Shared(std::rc::Rc::new(move |_| Value::String(pool.0.lock().unwrap().1.join(",")))))
}
