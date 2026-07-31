export const escapeGoStringImpl = function(s) {
  let out = "";
  for (let i = 0; i < s.length; i++) {
    const code = s.charCodeAt(i);
    if (code === 34) out += "\\\"";
    else if (code === 92) out += "\\\\";
    else if (code >= 0xD800 && code <= 0xDFFF) {
       if (code >= 0xD800 && code <= 0xDBFF && i + 1 < s.length) {
         const next = s.charCodeAt(i + 1);
         if (next >= 0xDC00 && next <= 0xDFFF) {
           out += s.charAt(i) + s.charAt(i + 1);
           i++;
           continue;
         }
       }
       const b1 = 0xE0 | (code >> 12);
       const b2 = 0x80 | ((code >> 6) & 0x3F);
       const b3 = 0x80 | (code & 0x3F);
       let hex1 = b1.toString(16).toLowerCase();
       let hex2 = b2.toString(16).toLowerCase();
       let hex3 = b3.toString(16).toLowerCase();
       out += "\\x" + (hex1.length === 1 ? "0" + hex1 : hex1) + 
              "\\x" + (hex2.length === 1 ? "0" + hex2 : hex2) + 
              "\\x" + (hex3.length === 1 ? "0" + hex3 : hex3);
    } else if (code < 32 || code === 127) {
       let hex = code.toString(16).toLowerCase();
       out += "\\x" + (hex.length === 1 ? "0" + hex : hex);
    } else {
       out += s.charAt(i);
    }
  }
  return out;
};
