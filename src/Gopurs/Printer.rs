pub fn Gopurs_Printer_escapeGoStringImpl(input: String) -> String {
    use std::fmt::Write;
    let mut output = String::with_capacity(input.len());
    let mut units = input.chars().peekable();
    while let Some(ch) = units.next() {
        let code = purust_char_to_code_unit(ch);
        match code {
            34 => output.push_str("\\\""),
            92 => output.push_str("\\\\"),
            0..=31 | 127 => write!(&mut output, "\\x{:02x}", code).unwrap(),
            0xd800..=0xdfff => {
                if code <= 0xdbff && units.peek().is_some_and(|next|
                    (0xdc00..=0xdfff).contains(&purust_char_to_code_unit(*next))) {
                    output.push(ch);
                    output.push(units.next().unwrap());
                } else {
                    write!(&mut output, "\\x{:02x}\\x{:02x}\\x{:02x}",
                        0xe0 | (code >> 12), 0x80 | ((code >> 6) & 0x3f), 0x80 | (code & 0x3f)).unwrap();
                }
            }
            _ => output.push(ch),
        }
    }
    output
}
