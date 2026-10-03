pub fn Gopurs_GoImports_concatStringArrays(arrays: Value) -> Value {
    let mut result = Vec::new();
    for array in arrays.unwrap_array().iter() {
        result.extend(array.unwrap_array().iter().cloned());
    }
    mk_array(result)
}
