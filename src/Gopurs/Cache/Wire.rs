unsafe extern "C" {
    fn GopursCacheDigest(content: *const u8, length: usize) -> *mut std::ffi::c_char;
    fn GopursFreeFFI(result: *mut std::ffi::c_char);
}

pub fn Gopurs_Cache_Wire_hash(text: String) -> String {
    let bytes = purust_string_to_utf8_lossy(&text);
    unsafe {
        let pointer = GopursCacheDigest(bytes.as_ptr(), bytes.len());
        let result = std::ffi::CStr::from_ptr(pointer).to_string_lossy().into_owned();
        GopursFreeFFI(pointer);
        purust_string_from_utf8(&result)
    }
}

pub fn Gopurs_Cache_Wire_stringUnits(text: String) -> String {
    text.chars().map(|c| format!("{:04x}", purust_char_to_code_unit(c))).collect()
}

pub fn Gopurs_Cache_Wire_hex(text: String) -> bool {
    text.bytes().all(|c| matches!(c, b'0'..=b'9' | b'a'..=b'f'))
}

pub fn Gopurs_Cache_Wire_unitsString(text: String) -> String {
    text.as_bytes().chunks_exact(4).map(|chunk| {
        let value = u16::from_str_radix(std::str::from_utf8(chunk).unwrap(), 16).unwrap();
        purust_char_from_code_unit(value)
    }).collect()
}

pub fn Gopurs_Cache_Wire_numberBits(value: f64) -> String { format!("{:016x}", value.to_bits()) }
pub fn Gopurs_Cache_Wire_bitsNumber(text: String) -> f64 { f64::from_bits(u64::from_str_radix(&text, 16).unwrap()) }
pub fn Gopurs_Cache_Wire_compareStrings(a: String, b: String) -> i64 {
    match a.chars().map(purust_char_to_code_unit).cmp(b.chars().map(purust_char_to_code_unit)) {
        std::cmp::Ordering::Less => -1, std::cmp::Ordering::Equal => 0, std::cmp::Ordering::Greater => 1,
    }
}
