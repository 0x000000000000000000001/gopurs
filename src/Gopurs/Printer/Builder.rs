#[derive(Debug)]
pub struct Builder(std::sync::Mutex<String>);

pub fn Gopurs_Printer_Builder_newBuilderImpl() -> Value {
    Value::Func1(Func1::Static(|_| {
        Value::Class(std::rc::Rc::new(std::rc::Rc::new(Builder(std::sync::Mutex::new(String::new())))))
    }))
}

pub fn Gopurs_Printer_Builder_pushImpl(builder: std::rc::Rc<Builder>, piece: String) -> std::rc::Rc<Builder> {
    builder.0.lock().unwrap().push_str(&piece);
    builder
}

pub fn Gopurs_Printer_Builder_toStringImpl(builder: std::rc::Rc<Builder>) -> String {
    builder.0.lock().unwrap().clone()
}
