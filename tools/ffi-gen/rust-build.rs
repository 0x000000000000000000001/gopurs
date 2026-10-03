fn main() {
    let archive = std::path::Path::new(env!("CARGO_MANIFEST_DIR")).join("native-parser");
    println!("cargo:rerun-if-changed={}", archive.join("libgopurs_ffi.a").display());
    println!("cargo:rustc-link-search=native={}", archive.display());
    println!("cargo:rustc-link-lib=static=gopurs_ffi");
    match std::env::var("CARGO_CFG_TARGET_OS").unwrap().as_str() {
        "macos" => println!("cargo:rustc-link-lib=resolv"),
        "linux" => {
            println!("cargo:rustc-link-lib=pthread");
            println!("cargo:rustc-link-lib=dl");
            println!("cargo:rustc-link-lib=m");
        }
        _ => {}
    }
}
