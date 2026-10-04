#![allow(warnings)]
// Differential contract for the borrowed `Qualified Ident` comparison embedded
// from CoreFn.rs. Every pair is checked against three real generated paths:
// the pure PureScript oracle (`...PS`), the public wrappers (the calls the
// optimizer makes) and the untouched generic `ordQualified` instance. The Rust
// native receives a panicking fallback, so a fallback call aborts the fixture.
use purust_core::*;
use std::sync::Arc as Rc;
use Purs_Data_Maybe::Maybe;
use Purs_Data_Ordering::Ordering;
use Purs_PureScript_Backend_Optimizer_CoreFn::{
    Qualified,
    PureScript_Backend_Optimizer_CoreFn_compareQualifiedIdent as compare_wrapper,
    PureScript_Backend_Optimizer_CoreFn_compareQualifiedIdentImpl as compare_native,
    PureScript_Backend_Optimizer_CoreFn_compareQualifiedIdentPS as compare_oracle,
    PureScript_Backend_Optimizer_CoreFn_eqQualifiedIdent as eq_wrapper,
    PureScript_Backend_Optimizer_CoreFn_eqQualifiedIdentImpl as eq_native,
    PureScript_Backend_Optimizer_CoreFn_eqQualifiedIdentPS as eq_oracle,
    PureScript_Backend_Optimizer_CoreFn_ordQualified as generic_ord_qualified,
};

fn forbidden_compare(_: Rc<Qualified>, _: Rc<Qualified>) -> Ordering {
    panic!("the Rust native comparison must not call the PureScript fallback")
}

fn forbidden_eq(_: Rc<Qualified>, _: Rc<Qualified>) -> bool {
    panic!("the Rust native equality must not call the PureScript fallback")
}

fn module_names() -> Vec<Option<String>> {
    let mut modules: Vec<Option<String>> = [
        None::<&str>,
        Some(""),
        Some("A"),
        Some("B"),
        Some("A.B"),
        Some("A.B.C"),
        Some("Data.Map"),
        Some("Data.Map.Internal"),
        Some("a"),
        Some("z"),
        Some("é"),
        Some("État"),
        Some("😀"),
        Some("\u{0}"),
    ]
    .into_iter()
    .map(|module| module.map(purust_string_from_utf8))
    .collect();
    // Module keys that only exist as UTF-16 code units, including isolated
    // surrogates. They are encoded once here; `qualified` must clone them.
    for units in [
        &[0xd800u16][..],
        &[0xdc00u16][..],
        &[0xd800, 0xdc00][..],
        &[0xd83d, 0xde00][..],
    ] {
        modules.push(Some(purust_string_from_utf16(units)));
    }
    modules
}

fn identifiers() -> Vec<String> {
    let mut idents: Vec<String> = [
        "", "a", "A", "z", "_", "'", "a'", "a_b", "a.b", "a::b",
        "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
        "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
        "é", "É", "e\u{301}", "État", "λ", "🦀", "😀", "\u{e000}",
        "\u{0}", "a\u{0}", "a\u{0}b", "\u{7f}", "\u{80}", "\u{7ff}", "\u{800}",
        "\u{ffff}", "\u{10000}", "\u{10ffff}",
    ]
    .iter()
    .map(|text| purust_string_from_utf8(text))
    .collect();
    // The runtime preserves isolated UTF-16 surrogates, so the differential
    // corpus must contain code-unit-only strings as well.
    for units in [
        &[0xd800u16][..],
        &[0xdc00u16][..],
        &[0xd800, 0xdc00][..],
        &[0xdc00, 0xd800][..],
        &[0x0041, 0xd800, 0x0042][..],
    ] {
        idents.push(purust_string_from_utf16(units));
    }
    idents
}

// `module` and `ident` already carry the runtime's internal encoding (one
// scalar per UTF-16 code unit); clone them, never re-encode, or the surrogate
// corpus would be corrupted by a second pass.
fn qualified(module: Option<&str>, ident: &str) -> Rc<Qualified> {
    Rc::new(Qualified::Qualified(
        Rc::new(match module {
            None => Maybe::Nothing,
            Some(name) => Maybe::Just(Value::String(name.to_owned())),
        }),
        Value::String(ident.to_owned()),
    ))
}

// The generic instance travels through boxed `Value` arguments, exactly like
// the optimizer's despecialized call sites. Sharing the inner `Rc` is the same
// representation as cloning the `Qualified` enum into a fresh `Rc`.
fn boxed(value: &Rc<Qualified>) -> Value {
    Value::Class(Rc::new(value.clone()))
}

// Generator-produced shared owner. The generic Ord instance must read it and
// mixed carriers through the same `unwrap_class_shared` path.
fn boxed_shared(value: &Rc<Qualified>) -> Value {
    Value::ClassShared(value.clone())
}

fn rank(ordering: &Ordering) -> i8 {
    match ordering {
        Ordering::LT => -1,
        Ordering::EQ => 0,
        Ordering::GT => 1,
    }
}

fn main() {
    let modules = module_names();
    let idents = identifiers();
    let ord = generic_ord_qualified(Purs_Data_Ord::Data_Ord_ordString());
    let generic_compare = ord.compare.clone();
    let mut pairs = 0usize;
    for module_a in &modules {
        for module_b in &modules {
            for ident_a in &idents {
                for ident_b in &idents {
                    let a = qualified(module_a.as_deref(), ident_a);
                    let b = qualified(module_b.as_deref(), ident_b);
                    let label = format!("{module_a:?}:{ident_a:?} vs {module_b:?}:{ident_b:?}");
                    let oracle_ordering = compare_oracle(a.clone(), b.clone());
                    let native_ordering =
                        compare_native(Func2::Static(forbidden_compare), a.clone(), b.clone());
                    let wrapped_ordering = compare_wrapper(a.clone(), b.clone());
                    let generic_ordering = generic_compare(boxed(&a), boxed(&b));
                    assert_eq!(
                        rank(&native_ordering),
                        rank(&oracle_ordering),
                        "native compare differs from the PS oracle: {label}"
                    );
                    assert_eq!(
                        rank(&wrapped_ordering),
                        rank(&oracle_ordering),
                        "generated wrapper differs from the PS oracle: {label}"
                    );
                    assert_eq!(
                        rank(&generic_ordering),
                        rank(&oracle_ordering),
                        "generic Ord Qualified drifted from the monomorphic oracle: {label}"
                    );
                    let oracle_eq = eq_oracle(a.clone(), b.clone());
                    let native_eq = eq_native(Func2::Static(forbidden_eq), a.clone(), b.clone());
                    let wrapped_eq = eq_wrapper(a.clone(), b.clone());
                    assert_eq!(native_eq, oracle_eq, "native eq differs from the PS oracle: {label}");
                    assert_eq!(wrapped_eq, oracle_eq, "generated eq wrapper differs from the PS oracle: {label}");
                    assert_eq!(
                        oracle_eq,
                        rank(&oracle_ordering) == 0,
                        "the PS Ord and Eq oracles disagree: {label}"
                    );
                    assert_eq!(
                        native_eq,
                        rank(&native_ordering) == 0,
                        "the native Ord and Eq paths disagree: {label}"
                    );
                    let reversed = compare_native(Func2::Static(forbidden_compare), b.clone(), a.clone());
                    assert_eq!(
                        rank(&reversed),
                        -rank(&native_ordering),
                        "antisymmetry violated: {label}"
                    );
                    assert_eq!(
                        rank(&compare_native(Func2::Static(forbidden_compare), a.clone(), a.clone())),
                        0,
                        "reflexivity violated: {label}"
                    );
                    pairs += 1;
                }
            }
        }
    }
    // ClassShared and mixed carriers on the generic instance: the same pairs
    // must keep the PS oracle ordering through the shared reader.
    let mut shared_pairs = 0usize;
    for module_a in modules.iter().step_by(4) {
        for module_b in modules.iter().step_by(5) {
            for ident_a in idents.iter().step_by(6) {
                for ident_b in idents.iter().step_by(7) {
                    let a = qualified(module_a.as_deref(), ident_a);
                    let b = qualified(module_b.as_deref(), ident_b);
                    let label = format!("{module_a:?}:{ident_a:?} vs {module_b:?}:{ident_b:?}");
                    let oracle_ordering = compare_oracle(a.clone(), b.clone());
                    assert_eq!(
                        rank(&generic_compare(boxed_shared(&a), boxed_shared(&b))),
                        rank(&oracle_ordering),
                        "shared generic Ord differs from the PS oracle: {label}"
                    );
                    assert_eq!(
                        rank(&generic_compare(boxed(&a), boxed_shared(&b))),
                        rank(&oracle_ordering),
                        "mixed generic Ord differs from the PS oracle: {label}"
                    );
                    assert_eq!(
                        rank(&generic_compare(boxed_shared(&a), boxed(&b))),
                        rank(&oracle_ordering),
                        "mixed generic Ord differs from the PS oracle: {label}"
                    );
                    shared_pairs += 1;
                }
            }
        }
    }
    println!(
        "Native Qualified comparison: {pairs} pairs against the PS oracle, generated wrappers and generic Ord; {shared_pairs} shared/mixed-carrier pairs; ASCII, Unicode, NUL and surrogate module/ident keys passed"
    );
}
