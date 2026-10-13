type Type = std::rc::Rc<Purs_PureScript_Backend_Optimizer_CoreFn::ExprType>;
pub struct TypeMemo(std::sync::Mutex<std::collections::HashMap<usize, (Type, i64)>>);

pub fn Gopurs_Cache_TypeMemo_create() -> Value {
    Value::Func1(Func1::Static(|_| Value::Class(std::rc::Rc::new(std::rc::Rc::new(
        TypeMemo(std::sync::Mutex::new(Default::default()))
    )))))
}

pub fn Gopurs_Cache_TypeMemo_lookup(memo: std::rc::Rc<TypeMemo>, value: Type) -> Value {
    Value::Func1(Func1::Shared(std::rc::Rc::new(move |_| {
        Value::Int(memo.0.lock().unwrap().get(&(std::rc::Rc::as_ptr(&value) as usize))
            .map_or(-1, |(_, index)| *index))
    })))
}

pub fn Gopurs_Cache_TypeMemo_insert(memo: std::rc::Rc<TypeMemo>, value: Type, index: i64) -> Value {
    Value::Func1(Func1::Shared(std::rc::Rc::new(move |_| {
        // Retain the type, preventing address reuse for the lifetime of the key.
        memo.0.lock().unwrap().insert(std::rc::Rc::as_ptr(&value) as usize, (value.clone(), index));
        Value::Unit
    })))
}
