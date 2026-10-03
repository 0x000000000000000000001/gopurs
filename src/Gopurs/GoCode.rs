// Same lexical boundaries and sorted result as referencedImportsPS and the Go
// native scanner. Strings, rune literals and comments do not introduce imports.
pub fn Gopurs_GoCode_referencedImportsImpl(_fallback: Func1<String, Value>, text: String) -> Value {
    let bytes = text.as_bytes();
    let mut index = 0;
    let mut imports = std::collections::BTreeSet::new();
    let ident = |b: u8| b.is_ascii_alphanumeric() || b == b'_' || b > 127;
    while index < bytes.len() {
        match bytes[index] {
            quote @ (b'"' | b'\'' | b'`') => {
                index += 1;
                while index < bytes.len() {
                    let b = bytes[index];
                    index += 1;
                    if b == quote { break; }
                    if b == b'\\' && quote != b'`' { index += 1; }
                }
            }
            b'/' if bytes.get(index + 1) == Some(&b'/') => {
                index += 2;
                while index < bytes.len() && bytes[index] != b'\n' { index += 1; }
            }
            b'/' if bytes.get(index + 1) == Some(&b'*') => {
                index += 2;
                while index < bytes.len() {
                    if bytes[index] == b'*' && bytes.get(index + 1) == Some(&b'/') { index += 2; break; }
                    index += 1;
                }
            }
            b if ident(b) => {
                let start = index;
                index += 1;
                while index < bytes.len() && ident(bytes[index]) { index += 1; }
                if bytes.get(index) == Some(&b'.') {
                    let path = match &text[start..index] {
                        "gopurs_runtime" => Some("gopurs/output/gopurs_runtime"),
                        "math" => Some("math"), "sync" => Some("sync"), "unsafe" => Some("unsafe"),
                        _ => None,
                    };
                    if let Some(path) = path { imports.insert(path); }
                }
            }
            _ => index += 1,
        }
    }
    mk_array(imports.into_iter().map(|path| Value::String(path.to_owned())).collect())
}
