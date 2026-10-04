#![allow(warnings)]
#![recursion_limit = "512"]
// Differential harness for the native CoreFn text decoder (Text.rs candidate).
//
// The candidate FFI is injected verbatim at // NATIVE_TEXT_FFI and compiled in
// this crate; the oracle is the validated PureScript decoder
// `jsonParser >=> decodeModulePS >=> lmap printJsonDecodeError`, built from the
// generated crates. Inputs come from CASES/manifest.tsv as raw UTF-8 files and
// are converted once through purust_string_from_utf8, exactly like
// FS.readTextFile UTF8 feeds parseModule.
//
// Value comparison is a stable canonical structural dump of the whole Module
// (types, annotations, source usage, expressions, binders, literals, cold
// declarations), never pointer equality and never the same decoder twice.
// Every Class/ClassShared handle is read through unwrap_class_shared so both
// carriers qualify; the counters prove which carriers ran.
//
// Failure cases compare the printed Left strings byte for byte. Success cases
// compare canonical dumps byte for byte and report an FNV-1a fingerprint.
//
// Usage: gopurs_text_differential CASES_DIR REPORT_JSON

use purust_core::*;
use std::sync::atomic::{AtomicUsize, Ordering};
use std::sync::Arc as Rc;
use Purs_Data_Argonaut_Decode_Error::{JsonDecodeError, Data_Argonaut_Decode_Error_printJsonDecodeError};
use Purs_Data_Either::Either;
use Purs_Data_Maybe::Maybe;
use Purs_PureScript_Backend_Optimizer_CoreFn::*;
use Purs_PureScript_Backend_Optimizer_CoreFn_Json::*;
use Purs_PureScript_Backend_Optimizer_CoreFn_Usage::*;
// Full PS oracle: the runner copies the generated Json/Usage crates into
// Oracle packages whose FFI entry points call their PureScript bodies. The
// candidate below still links the original crates for cold decoders and its
// validate function; only the reference composition uses the oracle copy.
use Purs_PureScript_Backend_Optimizer_CoreFn_JsonOracle as oracle_json;

mod candidate {
    use super::*;
    use purust_core::*;
    use Purs_PureScript_Backend_Optimizer_CoreFn_Json::*;
    // NATIVE_TEXT_FFI
}

// ---------------------------------------------------------------------------
// Class/ClassShared tolerant access

static CLASS_CARRIERS: AtomicUsize = AtomicUsize::new(0);
static SHARED_CARRIERS: AtomicUsize = AtomicUsize::new(0);

fn reset_carriers() {
    CLASS_CARRIERS.store(0, Ordering::Relaxed);
    SHARED_CARRIERS.store(0, Ordering::Relaxed);
}

fn carriers() -> (usize, usize) {
    (CLASS_CARRIERS.load(Ordering::Relaxed), SHARED_CARRIERS.load(Ordering::Relaxed))
}

// Mirrors Value::unwrap_class_shared while counting the carrier in use. Both
// the legacy nested Class box and the erasure-optimized ClassShared owner
// resolve to the same logical node.
static CASE_NAME: std::sync::Mutex<String> = std::sync::Mutex::new(String::new());
static FIRST_FAILURE: std::sync::atomic::AtomicBool = std::sync::atomic::AtomicBool::new(false);

fn set_case(name: &str) { *CASE_NAME.lock().unwrap() = name.to_owned(); }
fn case_name() -> String { CASE_NAME.lock().unwrap().clone() }

fn value_variant(value: &Value) -> &'static str {
    match value.resolve() {
        Value::Unit => "Unit", Value::Null => "Null", Value::Int(_) => "Int", Value::Number(_) => "Number",
        Value::Bool(_) => "Bool", Value::String(_) => "String", Value::Char(_) => "Char",
        Value::Array(_) => "Array", Value::Class(_) => "Class", Value::ClassShared(_) => "ClassShared",
        Value::Record_bindingUsage_variableUse(_) => "Record_bindingUsage_variableUse",
        Value::Record_binding_hasEscapingUseContext_maxUses(_) => "Record_binding_hasEscapingUseContext_maxUses",
        Value::Record_binding_lastLocalUse(_) => "Record_binding_lastLocalUse",
        Value::Record_bindingId_moduleName(_) => "Record_bindingId_moduleName",
        Value::Record_meta_sourceUsage_span_type_kw(_) => "Record_meta_sourceUsage_span_type_kw",
        Value::Record_meta_type_kw(_) => "Record_meta_type_kw",
        Value::Record_meta(_) => "Record_meta",
        Value::Record_a(_) => "Record_a",
        Value::DynamicRecord(_) => "DynamicRecord",
        _ => "other",
    }
}

fn node<T: Clone + std::any::Any + Send + Sync + 'static>(value: &Value) -> Rc<T> {
    match value.resolve() {
        Value::ClassShared(_) => {
            SHARED_CARRIERS.fetch_add(1, Ordering::Relaxed);
            value.unwrap_class_shared::<T>()
        }
        Value::Class(_) => {
            CLASS_CARRIERS.fetch_add(1, Ordering::Relaxed);
            value.unwrap_class_shared::<T>()
        }
        _ => {
            eprintln!(
                "node<{}> unexpected {} in case {}",
                std::any::type_name::<T>(),
                value_variant(value),
                case_name());
            panic!("expected a Class/ClassShared node");
        }
    }
}

fn maybe_of(value: &Value) -> Rc<Maybe> { node::<Maybe>(value) }

// ---------------------------------------------------------------------------
// Canonical structural dump

struct Dump(String);
impl Dump {
    fn new() -> Self { Dump(String::with_capacity(1 << 16)) }
    fn put(&mut self, text: &str) { self.0.push_str(text); }
    fn end(&mut self) { self.0.push(';'); }
}

fn dump_string(out: &mut Dump, value: &str) {
    let units = purust_string_to_utf16(value);
    out.put(&format!("s{}:", units.len()));
    for unit in units { out.put(&format!("{:04x}", unit)); }
    out.end();
}

fn dump_i64(out: &mut Dump, value: i64) { out.put(&format!("i{}", value)); out.end(); }
fn dump_bool(out: &mut Dump, value: bool) { out.put(if value { "b1" } else { "b0" }); out.end(); }
fn dump_f64(out: &mut Dump, value: f64) { out.put(&format!("n{:016x}", value.to_bits())); out.end(); }
fn dump_char(out: &mut Dump, value: char) { out.put(&format!("c{:06x}", value as u32)); out.end(); }

fn dump_string_value(out: &mut Dump, value: &Value) { dump_string(out, &value.unwrap_string()); }
fn dump_int_value(out: &mut Dump, value: &Value) { dump_i64(out, value.unwrap_int()); }
fn dump_bool_value(out: &mut Dump, value: &Value) { dump_bool(out, value.unwrap_bool()); }

fn dump_maybe(out: &mut Dump, value: &Value, element: fn(&mut Dump, &Value)) {
    match maybe_of(value).as_ref() {
        Maybe::Nothing => out.put("nothing"),
        Maybe::Just(item) => { out.put("just("); element(out, item); out.put(")"); }
    }
    out.end();
}

fn dump_string_array(out: &mut Dump, value: &Value) {
    out.put("[");
    for index in 0..value.array_len() {
        dump_string_value(out, &value.array_get(index));
    }
    out.put("]");
    out.end();
}

fn dump_span(out: &mut Dump, value: &Value) {
    dump_string_value(out, &value.get_path());
    dump_int_value(out, &value.get_start().get_line());
    dump_int_value(out, &value.get_start().get_column());
    dump_int_value(out, &value.get_end().get_line());
    dump_int_value(out, &value.get_end().get_column());
}

fn dump_binding_id(out: &mut Dump, value: &Value) {
    dump_string_value(out, &value.get_moduleName());
    dump_int_value(out, &value.get_bindingId());
}

fn dump_binding_usage(out: &mut Dump, value: &Value) {
    dump_binding_id(out, &value.get_binding());
    dump_maybe(out, &value.get_maxUses(), dump_int_value);
    dump_maybe(out, &value.get_hasEscapingUseContext(), dump_bool_value);
}

fn dump_variable_use(out: &mut Dump, value: &Value) {
    dump_binding_id(out, &value.get_binding());
    dump_maybe(out, &value.get_lastLocalUse(), dump_bool_value);
}

fn dump_source_usage(out: &mut Dump, value: &Value) {
    // `value` is the payload of the sourceUsage Maybe: the usage record itself.
    // Every field is compared, including maxUses, hasEscapingUseContext and
    // lastLocalUse, through the nested Maybe dumps below.
    out.put("usage(");
    dump_maybe(out, &value.get_bindingUsage(), dump_binding_usage);
    dump_maybe(out, &value.get_variableUse(), dump_variable_use);
    out.put(")");
    out.end();
}

fn dump_expr_type_value(out: &mut Dump, value: &Value) {
    dump_expr_type(out, node::<ExprType>(value).as_ref());
}

fn dump_expr_type_array(out: &mut Dump, value: &Value) {
    out.put("[");
    for index in 0..value.array_len() {
        dump_expr_type_value(out, &value.array_get(index));
    }
    out.put("]");
    out.end();
}

fn dump_tuple_type(out: &mut Dump, value: &Value) {
    match node::<Purs_Data_Tuple::Tuple>(value).as_ref() {
        Purs_Data_Tuple::Tuple::Tuple(first, second) => {
            dump_string_value(out, first);
            dump_expr_type_value(out, second);
        }
    }
}

fn dump_constraint(out: &mut Dump, value: &Value) {
    match node::<Purs_Data_Tuple::Tuple>(value).as_ref() {
        Purs_Data_Tuple::Tuple::Tuple(first, second) => {
            dump_string_array(out, first);
            dump_expr_type_array(out, second);
        }
    }
}

fn dump_row_field(out: &mut Dump, value: &Value) {
    match node::<Purs_Data_Tuple::Tuple>(value).as_ref() {
        Purs_Data_Tuple::Tuple::Tuple(first, second) => {
            dump_string_value(out, first);
            dump_expr_type_value(out, second);
        }
    }
}

fn dump_expr_type(out: &mut Dump, value: &ExprType) {
    match value {
        ExprType::Int => out.put("Int"),
        ExprType::Number => out.put("Number"),
        ExprType::String => out.put("String"),
        ExprType::Char => out.put("Char"),
        ExprType::Boolean => out.put("Boolean"),
        ExprType::Unit => out.put("Unit"),
        ExprType::Any => out.put("Any"),
        ExprType::TypeLevelString(name) => { out.put("TypeLevelString("); dump_string(out, name); out.put(")"); }
        ExprType::TypeVar(name) => { out.put("TypeVar("); dump_string(out, name); out.put(")"); }
        ExprType::Array(element) => { out.put("Array("); dump_expr_type(out, element.as_ref()); out.put(")"); }
        ExprType::Record(row) => { out.put("Record("); dump_expr_type(out, row.as_ref()); out.put(")"); }
        ExprType::ADT(name, fqn, args) => {
            out.put("ADT("); dump_string(out, name);
            dump_string_array(out, fqn);
            dump_expr_type_array(out, args);
            out.put(")");
        }
        ExprType::TypeApp(base, args) => {
            out.put("TypeApp("); dump_expr_type(out, base.as_ref());
            dump_expr_type_array(out, args);
            out.put(")");
        }
        ExprType::Func(args, ret) => {
            out.put("Func("); dump_expr_type_array(out, args);
            dump_expr_type(out, ret.as_ref());
            out.put(")");
        }
        ExprType::Row(fields, tail) => {
            out.put("Row(");
            out.put("[");
            for index in 0..fields.array_len() { dump_row_field(out, &fields.array_get(index)); }
            out.put("]");
            out.end();
            match tail.as_ref() {
                Maybe::Nothing => out.put("nothing"),
                Maybe::Just(item) => { out.put("just("); dump_expr_type_value(out, item); out.put(")"); }
            }
            out.end();
            out.put(")");
        }
        ExprType::ForAll(vars, body) => {
            out.put("ForAll("); dump_string_array(out, vars);
            dump_expr_type(out, body.as_ref());
            out.put(")");
        }
        ExprType::ConstrainedType(constraints, body) => {
            out.put("ConstrainedType(");
            out.put("[");
            for index in 0..constraints.array_len() { dump_constraint(out, &constraints.array_get(index)); }
            out.put("]");
            out.end();
            dump_expr_type(out, body.as_ref());
            out.put(")");
        }
    }
}

fn dump_constructor_type(out: &mut Dump, value: &ConstructorType) {
    match value { ConstructorType::ProductType => out.put("ProductType"), ConstructorType::SumType => out.put("SumType") }
    out.end();
}

fn dump_meta(out: &mut Dump, value: &Value) {
    match node::<Meta>(value).as_ref() {
        Meta::IsConstructor(constructor, identifiers) => {
            out.put("IsConstructor(");
            dump_constructor_type(out, constructor);
            dump_string_array(out, identifiers);
            out.put(")");
        }
        Meta::IsNewtype => out.put("IsNewtype"),
        Meta::IsTypeClassConstructor => out.put("IsTypeClassConstructor"),
        Meta::IsForeign => out.put("IsForeign"),
        Meta::IsWhere => out.put("IsWhere"),
        Meta::IsSyntheticApp => out.put("IsSyntheticApp"),
    }
    out.end();
}

fn dump_ann(out: &mut Dump, value: &Value) {
    dump_span(out, &value.get_span());
    dump_maybe(out, &value.get_meta(), dump_meta);
    dump_maybe(out, &value.get_type_kw(), dump_expr_type_value);
    dump_maybe(out, &value.get_sourceUsage(), dump_source_usage);
}

fn dump_qualified(out: &mut Dump, value: &Qualified) {
    match value {
        Qualified::Qualified(module, identifier) => {
            match module.as_ref() {
                Maybe::Nothing => out.put("nothing"),
                Maybe::Just(name) => { out.put("just("); dump_string_value(out, name); out.put(")"); }
            }
            out.end();
            dump_string_value(out, identifier);
        }
    }
}

fn dump_qualified_value(out: &mut Dump, value: &Value) {
    dump_qualified(out, node::<Qualified>(value).as_ref());
}

#[derive(Clone, Copy)]
enum Element { Expr, Binder }

fn dump_literal(out: &mut Dump, value: &Literal, element: Element) {
    match value {
        Literal::LitInt(item) => { out.put("LitInt("); dump_i64(out, *item); out.put(")"); }
        Literal::LitNumber(item) => { out.put("LitNumber("); dump_f64(out, *item); out.put(")"); }
        Literal::LitString(item) => { out.put("LitString("); dump_string(out, item); out.put(")"); }
        Literal::LitChar(item) => { out.put("LitChar("); dump_char(out, *item); out.put(")"); }
        Literal::LitBoolean(item) => { out.put("LitBoolean("); dump_bool(out, *item); out.put(")"); }
        Literal::LitArray(items) => {
            out.put("LitArray[");
            for index in 0..items.array_len() {
                match element {
                    Element::Expr => dump_expr_value(out, &items.array_get(index)),
                    Element::Binder => dump_binder_value(out, &items.array_get(index)),
                }
            }
            out.put("]");
        }
        Literal::LitRecord(items) => {
            out.put("LitRecord[");
            for index in 0..items.array_len() { dump_prop(out, &items.array_get(index), element); }
            out.put("]");
        }
    }
    out.end();
}

fn dump_prop(out: &mut Dump, value: &Value, element: Element) {
    match node::<Prop>(value).as_ref() {
        Prop::Prop(key, item) => {
            dump_string(out, key);
            match element {
                Element::Expr => dump_expr_value(out, item),
                Element::Binder => dump_binder_value(out, item),
            }
        }
    }
}

fn dump_expr_value(out: &mut Dump, value: &Value) { dump_expr(out, node::<Expr>(value).as_ref()); }

fn dump_expr(out: &mut Dump, value: &Expr) {
    match value {
        Expr::ExprVar(ann, qualified) => {
            out.put("ExprVar("); dump_ann(out, ann); dump_qualified(out, qualified.as_ref()); out.put(")");
        }
        Expr::ExprLit(ann, literal) => {
            out.put("ExprLit("); dump_ann(out, ann); dump_literal(out, literal.as_ref(), Element::Expr); out.put(")");
        }
        Expr::ExprConstructor(ann, type_name, constructor, fields) => {
            out.put("ExprConstructor("); dump_ann(out, ann);
            dump_string(out, type_name); dump_string(out, constructor); dump_string_array(out, fields);
            out.put(")");
        }
        Expr::ExprAccessor(ann, expression, field) => {
            out.put("ExprAccessor("); dump_ann(out, ann);
            dump_expr(out, expression.as_ref()); dump_string(out, field); out.put(")");
        }
        Expr::ExprUpdate(ann, expression, updates) => {
            out.put("ExprUpdate("); dump_ann(out, ann);
            dump_expr(out, expression.as_ref());
            out.put("[");
            for index in 0..updates.array_len() { dump_prop(out, &updates.array_get(index), Element::Expr); }
            out.put("]"); out.end();
            out.put(")");
        }
        Expr::ExprAbs(ann, argument, body) => {
            out.put("ExprAbs("); dump_ann(out, ann); dump_string(out, argument);
            dump_expr(out, body.as_ref()); out.put(")");
        }
        Expr::ExprApp(ann, abstraction, argument) => {
            out.put("ExprApp("); dump_ann(out, ann);
            dump_expr(out, abstraction.as_ref()); dump_expr(out, argument.as_ref()); out.put(")");
        }
        Expr::ExprCase(ann, expressions, alternatives) => {
            out.put("ExprCase("); dump_ann(out, ann);
            out.put("[");
            for index in 0..expressions.array_len() { dump_expr_value(out, &expressions.array_get(index)); }
            out.put("]"); out.end();
            out.put("[");
            for index in 0..alternatives.array_len() { dump_alternative(out, &alternatives.array_get(index)); }
            out.put("]"); out.end();
            out.put(")");
        }
        Expr::ExprLet(ann, binds, body) => {
            out.put("ExprLet("); dump_ann(out, ann);
            out.put("[");
            for index in 0..binds.array_len() { dump_bind(out, &binds.array_get(index)); }
            out.put("]"); out.end();
            dump_expr(out, body.as_ref()); out.put(")");
        }
        Expr::ExprTypeApp(ann, expression, ty) => {
            out.put("ExprTypeApp("); dump_ann(out, ann);
            dump_expr(out, expression.as_ref()); dump_expr_type(out, ty.as_ref()); out.put(")");
        }
    }
    out.end();
}

fn dump_binder_value(out: &mut Dump, value: &Value) { dump_binder(out, node::<Binder>(value).as_ref()); }

fn dump_binder(out: &mut Dump, value: &Binder) {
    match value {
        Binder::BinderNull(ann) => { out.put("BinderNull("); dump_ann(out, ann); out.put(")"); }
        Binder::BinderVar(ann, name) => {
            out.put("BinderVar("); dump_ann(out, ann); dump_string(out, name); out.put(")");
        }
        Binder::BinderNamed(ann, name, inner) => {
            out.put("BinderNamed("); dump_ann(out, ann); dump_string(out, name);
            dump_binder(out, inner.as_ref()); out.put(")");
        }
        Binder::BinderLit(ann, literal) => {
            out.put("BinderLit("); dump_ann(out, ann); dump_literal(out, literal.as_ref(), Element::Binder); out.put(")");
        }
        Binder::BinderConstructor(ann, type_name, constructor, binders) => {
            out.put("BinderConstructor("); dump_ann(out, ann);
            dump_qualified(out, type_name.as_ref()); dump_qualified(out, constructor.as_ref());
            out.put("[");
            for index in 0..binders.array_len() { dump_binder_value(out, &binders.array_get(index)); }
            out.put("]"); out.end();
            out.put(")");
        }
    }
    out.end();
}

fn dump_binding_value(out: &mut Dump, value: &Value) {
    dump_binding(out, node::<Binding>(value).as_ref());
}

fn dump_binding(out: &mut Dump, value: &Binding) {
    match value {
        Binding::Binding(ann, name, expression) => {
            out.put("Binding("); dump_ann(out, ann); dump_string(out, name);
            dump_expr(out, expression.as_ref()); out.put(")");
        }
    }
}

fn dump_bind(out: &mut Dump, value: &Value) {
    match node::<Bind>(value).as_ref() {
        Bind::NonRec(binding) => {
            out.put("NonRec(");
            dump_binding(out, binding.as_ref());
            out.put(")");
        }
        Bind::Rec(group) => {
            out.put("Rec[");
            for index in 0..group.array_len() { dump_binding_value(out, &group.array_get(index)); }
            out.put("]");
        }
    }
    out.end();
}

fn dump_guard(out: &mut Dump, value: &Value) {
    match node::<Guard>(value).as_ref() {
        Guard::Guard(condition, result) => {
            out.put("Guard(");
            dump_expr(out, condition.as_ref());
            dump_expr(out, result.as_ref());
            out.put(")");
        }
    }
    out.end();
}

fn dump_alternative(out: &mut Dump, value: &Value) {
    match node::<CaseAlternative>(value).as_ref() {
        CaseAlternative::CaseAlternative(binders, guard) => {
            out.put("CaseAlternative(");
            out.put("[");
            for index in 0..binders.array_len() { dump_binder_value(out, &binders.array_get(index)); }
            out.put("]"); out.end();
            match guard.as_ref() {
                CaseGuard::Unconditional(expression) => {
                    out.put("Unconditional("); dump_expr(out, expression.as_ref()); out.put(")");
                }
                CaseGuard::Guarded(guards) => {
                    out.put("Guarded[");
                    for index in 0..guards.array_len() { dump_guard(out, &guards.array_get(index)); }
                    out.put("]");
                }
            }
            out.end();
            out.put(")");
        }
    }
    out.end();
}

fn dump_import(out: &mut Dump, value: &Value) {
    match node::<Import>(value).as_ref() {
        Import::Import(ann, name) => {
            out.put("Import("); dump_ann(out, ann); dump_string(out, name); out.put(")");
        }
    }
    out.end();
}

fn dump_re_export(out: &mut Dump, value: &Value) {
    match node::<ReExport>(value).as_ref() {
        ReExport::ReExport(module, idents) => {
            out.put("ReExport("); dump_string(out, module); dump_string(out, idents); out.put(")");
        }
    }
    out.end();
}

fn dump_comment(out: &mut Dump, value: &Value) {
    match node::<Comment>(value).as_ref() {
        Comment::LineComment(text) => { out.put("LineComment("); dump_string(out, text); out.put(")"); }
        Comment::BlockComment(text) => { out.put("BlockComment("); dump_string(out, text); out.put(")"); }
    }
    out.end();
}

fn dump_data_decl(out: &mut Dump, value: &Value) {
    out.put("DataDecl(");
    dump_string_value(out, &value.get_name());
    dump_string_array(out, &value.get_vars());
    let constructors = value.get_constructors();
    out.put("[");
    for index in 0..constructors.array_len() {
        let constructor = constructors.array_get(index);
        out.put("DataConstructor(");
        dump_string_value(out, &constructor.get_name());
        dump_expr_type_array(out, &constructor.get_fields());
        out.put(")");
        out.end();
    }
    out.put("]"); out.end();
    out.put(")");
    out.end();
}

fn dump_class_decl(out: &mut Dump, value: &Value) {
    out.put("ClassDecl(");
    dump_string_value(out, &value.get_name());
    dump_string_array(out, &value.get_vars());
    let superclasses = value.get_superclasses();
    out.put("[");
    for index in 0..superclasses.array_len() { dump_constraint(out, &superclasses.array_get(index)); }
    out.put("]"); out.end();
    let methods = value.get_methods();
    out.put("[");
    for index in 0..methods.array_len() {
        match node::<Purs_Data_Tuple::Tuple>(&methods.array_get(index)).as_ref() {
            Purs_Data_Tuple::Tuple::Tuple(name, ty) => {
                out.put("Method("); dump_string_value(out, name); dump_expr_type_value(out, ty); out.put(")");
            }
        }
        out.end();
    }
    out.put("]"); out.end();
    out.put(")");
    out.end();
}

fn foreign_entries(map: &Value) -> Vec<(String, Value)> {
    let map = node::<Purs_Data_Map_Internal::Map>(map);
    let entries = Purs_Data_Map_Internal::Data_Map_Internal_toUnfoldableUnordered(
        Purs_Data_Unfoldable::Data_Unfoldable_unfoldableArray(), map);
    let mut result = Vec::with_capacity(entries.array_len());
    for index in 0..entries.array_len() {
        match node::<Purs_Data_Tuple::Tuple>(&entries.array_get(index)).as_ref() {
            Purs_Data_Tuple::Tuple::Tuple(key, value) => {
                result.push((key.unwrap_string(), value.clone()));
            }
        }
    }
    result.sort_by(|left, right| left.0.cmp(&right.0));
    result
}

fn dump_foreign(out: &mut Dump, value: &Value) {
    out.put("{");
    for (key, ty) in foreign_entries(value) {
        dump_string(out, &key);
        dump_maybe(out, &ty, dump_expr_type_value);
    }
    out.put("}");
    out.end();
}

fn dump_module(out: &mut Dump, value: &Value) {
    out.put("Module(");
    dump_string_value(out, &value.get_name());
    dump_string_value(out, &value.get_path());
    dump_span(out, &value.get_span());
    let imports = value.get_imports();
    out.put("[");
    for index in 0..imports.array_len() { dump_import(out, &imports.array_get(index)); }
    out.put("]"); out.end();
    dump_string_array(out, &value.get_exports());
    let re_exports = value.get_reExports();
    out.put("[");
    for index in 0..re_exports.array_len() { dump_re_export(out, &re_exports.array_get(index)); }
    out.put("]"); out.end();
    let data_decls = value.get_dataDecls();
    out.put("[");
    for index in 0..data_decls.array_len() { dump_data_decl(out, &data_decls.array_get(index)); }
    out.put("]"); out.end();
    let class_decls = value.get_classDecls();
    out.put("[");
    for index in 0..class_decls.array_len() { dump_class_decl(out, &class_decls.array_get(index)); }
    out.put("]"); out.end();
    let decls = value.get_decls();
    out.put("[");
    for index in 0..decls.array_len() { dump_bind(out, &decls.array_get(index)); }
    out.put("]"); out.end();
    dump_foreign(out, &value.get_foreign());
    let comments = value.get_comments();
    out.put("[");
    for index in 0..comments.array_len() { dump_comment(out, &comments.array_get(index)); }
    out.put("]"); out.end();
    out.put(")");
    out.end();
}

fn canonical_module(value: &Value) -> String {
    let mut out = Dump::new();
    dump_module(&mut out, value);
    out.0
}

fn fnv1a(bytes: &[u8]) -> u64 {
    let mut hash = 0xcbf29ce484222325u64;
    for byte in bytes {
        hash ^= *byte as u64;
        hash = hash.wrapping_mul(0x100000001b3);
    }
    hash
}

// ---------------------------------------------------------------------------
// Oracle and run

// Pure PS reference: the oracle crate's decodeModulePS runs the generated PS
// composition with its FFI entry points rewritten to call decodeAnnWithUsagePS,
// decodeArrayPS, decodeTypeTablePS and validateSourceUsageModulePS.
fn oracle(text: &str) -> Rc<Either> {
    match Purs_Data_Argonaut_Core::purust_json_parse_text(text) {
        Ok(json) => match oracle_json::PureScript_Backend_Optimizer_CoreFn_Json_decodeModulePS(json).as_ref() {
            Either::Left(error) => {
                let error = node::<JsonDecodeError>(error);
                Rc::new(Either::Left(Value::String(Data_Argonaut_Decode_Error_printJsonDecodeError(error))))
            }
            Either::Right(module) => Rc::new(Either::Right(module.clone())),
        },
        Err(message) => Rc::new(Either::Left(Value::String(message))),
    }
}

struct Run {
    candidate: Rc<Either>,
    reference: Rc<Either>,
    fallbacks: usize,
    validates: usize,
}

fn run_case(text: &str) -> Run {
    let fallback_calls = Rc::new(AtomicUsize::new(0));
    let fallback_counter = fallback_calls.clone();
    let fallback = Func1::Shared(Rc::new(move |input: String| -> Rc<Either> {
        fallback_counter.fetch_add(1, Ordering::SeqCst);
        oracle(&input)
    }));
    let validate_calls = Rc::new(AtomicUsize::new(0));
    let validate_counter = validate_calls.clone();
    let validate = Func1::Shared(Rc::new(move |module: Value| -> Rc<Either> {
        validate_counter.fetch_add(1, Ordering::SeqCst);
        PureScript_Backend_Optimizer_CoreFn_Usage_validateSourceUsageModule(module)
    }));
    let print_error = Func1::Shared(Rc::new(|error: Rc<JsonDecodeError>| -> String {
        Data_Argonaut_Decode_Error_printJsonDecodeError(error)
    }));
    let candidate = candidate::PureScript_Backend_Optimizer_CoreFn_Json_Text_parseModuleTextImpl(
        fallback, validate, print_error, text.to_owned());
    let reference = oracle(text);
    Run {
        candidate,
        reference,
        fallbacks: fallback_calls.load(Ordering::SeqCst),
        validates: validate_calls.load(Ordering::SeqCst),
    }
}

fn compare(run: &Run) -> Result<bool, String> {
    // Error-only cases have no AST carriers. Do not inherit the previous
    // successful case's counters, and include both dumps for successful cases.
    reset_carriers();
    match (run.candidate.as_ref(), run.reference.as_ref()) {
        (Either::Left(candidate), Either::Left(reference)) => {
            let candidate = candidate.unwrap_string();
            let reference = reference.unwrap_string();
            if candidate != reference {
                return Err(format!(
                    "different Left strings: candidate {} vs reference {}",
                    purust_string_to_utf8_lossy(&candidate),
                    purust_string_to_utf8_lossy(&reference)));
            }
            Ok(false)
        }
        (Either::Right(candidate), Either::Right(reference)) => {
            reset_carriers();
            let candidate_dump = canonical_module(candidate);
            let candidate_carriers = carriers();
            reset_carriers();
            let reference_dump = canonical_module(reference);
            let reference_carriers = carriers();
            CLASS_CARRIERS.fetch_add(candidate_carriers.0, Ordering::Relaxed);
            SHARED_CARRIERS.fetch_add(candidate_carriers.1, Ordering::Relaxed);
            if candidate_dump != reference_dump {
                let at = candidate_dump.bytes().zip(reference_dump.bytes())
                    .position(|(left, right)| left != right)
                    .unwrap_or_else(|| candidate_dump.len().min(reference_dump.len()));
                return Err(format!(
                    "structural mismatch at byte {} (candidate {} bytes, reference {} bytes; \
                     candidate fp {:016x}, reference fp {:016x}; carriers native {:?}/{:?}, \
                     oracle {:?}/{:?})",
                    at, candidate_dump.len(), reference_dump.len(),
                    fnv1a(candidate_dump.as_bytes()), fnv1a(reference_dump.as_bytes()),
                    candidate_carriers.0, candidate_carriers.1,
                    reference_carriers.0, reference_carriers.1));
            }
            Ok(true)
        }
        (Either::Left(candidate), Either::Right(_)) => Err(format!(
            "native returned Left {} but the oracle succeeded",
            purust_string_to_utf8_lossy(&candidate.unwrap_string()))),
        (Either::Right(_), Either::Left(reference)) => Err(format!(
            "native succeeded but the oracle returned Left {}",
            purust_string_to_utf8_lossy(&reference.unwrap_string()))),
    }
}

// Type-table entries must stay one shared owner across annotations and
// ExprTypeApp under both Class and ClassShared carriers.
fn sharing_check(module: &Value) -> Result<(), String> {
    let decls = module.get_decls();
    let annotation_type = |index: usize| -> Rc<ExprType> {
        let bind = decls.array_get(index);
        let binding = match node::<Bind>(&bind).as_ref() {
            Bind::NonRec(binding) => binding.clone(),
            _ => panic!("expected NonRec binding"),
        };
        let annotation = match binding.as_ref() {
            Binding::Binding(annotation, _, _) => annotation.clone(),
        };
        match maybe_of(&annotation.get_type_kw()).as_ref() {
            Maybe::Just(entry) => node::<ExprType>(entry),
            Maybe::Nothing => panic!("annotation has no type"),
        }
    };
    let first = annotation_type(0);
    let second = annotation_type(1);
    if !Rc::ptr_eq(&first, &second) {
        return Err("type-table entry is not shared between annotations".into());
    }
    let bind = decls.array_get(1);
    let expression = match node::<Bind>(&bind).as_ref() {
        Bind::NonRec(binding) => match binding.as_ref() {
            Binding::Binding(_, _, expression) => expression.clone(),
        },
        _ => panic!("expected NonRec binding"),
    };
    match expression.as_ref() {
        Expr::ExprTypeApp(_, _, ty) => {
            if !Rc::ptr_eq(&first, ty) {
                return Err("ExprTypeApp does not share the table entry".into());
            }
        }
        _ => return Err("expected ExprTypeApp".into()),
    }
    Ok(())
}

// ---------------------------------------------------------------------------
// Report

fn json_string(value: &str) -> String { Purs_Data_Argonaut_Core::purust_json_quote(value) }

struct Stats {
    total: usize,
    ok: usize,
    right: usize,
    left: usize,
    fast: usize,
    declined: usize,
    class_carriers: usize,
    shared_carriers: usize,
}

impl Stats {
    fn new() -> Self {
        Stats { total: 0, ok: 0, right: 0, left: 0, fast: 0, declined: 0, class_carriers: 0, shared_carriers: 0 }
    }
}

// Keep the first failing snapshot: the case name plus both canonical dumps
// (or both Left strings) beside the report, so a retained run is self-contained.
fn preserve_failure(report_path: &std::path::Path, context: &str, run: &Run) {
    if FIRST_FAILURE.swap(true, std::sync::atomic::Ordering::SeqCst) { return; }
    let directory = report_path.parent().unwrap_or_else(|| std::path::Path::new("."));
    let _ = std::fs::write(directory.join("failure-case.txt"), format!("{context}\n"));
    match (run.candidate.as_ref(), run.reference.as_ref()) {
        (Either::Right(candidate), Either::Right(reference)) => {
            let _ = std::fs::write(directory.join("failure-candidate.dump.txt"), canonical_module(candidate));
            let _ = std::fs::write(directory.join("failure-reference.dump.txt"), canonical_module(reference));
        }
        _ => {
            let candidate = match run.candidate.as_ref() {
                Either::Left(error) => purust_string_to_utf8_lossy(&error.unwrap_string()),
                Either::Right(_) => String::from("<right>"),
            };
            let reference = match run.reference.as_ref() {
                Either::Left(error) => purust_string_to_utf8_lossy(&error.unwrap_string()),
                Either::Right(_) => String::from("<right>"),
            };
            let _ = std::fs::write(directory.join("failure-errors.txt"), format!("candidate: {candidate}\nreference: {reference}\n"));
        }
    }
}

fn main() {
    let args: Vec<String> = std::env::args().collect();
    assert_eq!(args.len(), 3, "Usage: gopurs_text_differential CASES_DIR REPORT_JSON");
    let cases = std::path::PathBuf::from(&args[1]);
    let report_path = std::path::PathBuf::from(&args[2]);
    let manifest = std::fs::read_to_string(cases.join("manifest.tsv")).expect("cases manifest.tsv");

    let mut kinds: Vec<(String, Stats)> = Vec::new();
    let mut failures: Vec<String> = Vec::new();

    for line in manifest.lines() {
        if line.is_empty() { continue; }
        let mut cells = line.splitn(3, '\t');
        let kind = cells.next().unwrap_or("case").to_owned();
        let name = cells.next().unwrap_or("").to_owned();
        let file = cells.next().unwrap_or("");
        let slot = match kinds.iter().position(|entry| entry.0 == kind) {
            Some(index) => index,
            None => { kinds.push((kind.clone(), Stats::new())); kinds.len() - 1 }
        };
        let stats = &mut kinds[slot].1;
        stats.total += 1;

        let raw = std::fs::read_to_string(cases.join(file))
            .unwrap_or_else(|error| panic!("cannot read {}: {}", file, error));
        let text = purust_string_from_utf8(&raw);
        set_case(&format!("{kind} {name}"));
        let run = run_case(&text);
        let context = format!("{} {}", kind, name);
        if run.fallbacks == 0 { stats.fast += 1; } else { stats.declined += 1; }

        match compare(&run) {
            Ok(right) => {
                if right { stats.right += 1; } else { stats.left += 1; }
                let mut failure = None;
                if kind == "corpus" {
                    if run.fallbacks != 0 { failure = Some(format!("{context}: valid module declined the native path")); }
                    else if run.validates != 1 { failure = Some(format!("{context}: validate must run exactly once, ran {}", run.validates)); }
                    else if !right { failure = Some(format!("{context}: expected a successful decode")); }
                } else if kind == "boundary-fast" {
                    if run.fallbacks != 0 { failure = Some(format!("{context}: expected the native fast path")); }
                    else if run.validates != 1 { failure = Some(format!("{context}: validate must run exactly once, ran {}", run.validates)); }
                    else if !right { failure = Some(format!("{context}: expected a successful decode")); }
                } else if kind == "boundary-validate" {
                    if run.fallbacks != 0 { failure = Some(format!("{context}: expected the native fast path")); }
                    else if run.validates != 1 { failure = Some(format!("{context}: validate must run exactly once, ran {}", run.validates)); }
                    else if right { failure = Some(format!("{context}: expected the validate error")); }
                } else if kind == "boundary-fallback" {
                    if run.fallbacks != 1 { failure = Some(format!("{context}: expected exactly one fallback, saw {}", run.fallbacks)); }
                    else if run.validates != 0 { failure = Some(format!("{context}: declined input must never reach validate")); }
                    else if right { failure = Some(format!("{context}: invalid fixture unexpectedly succeeded")); }
                }
                if failure.is_none() && name == "valid-sharing" {
                    if let Either::Right(module) = run.candidate.as_ref() {
                        if let Err(error) = sharing_check(module) { failure = Some(format!("{context}: {error}")); }
                    } else {
                        failure = Some(format!("{context}: expected a decoded module for sharing"));
                    }
                }
                match failure {
                    Some(error) => {
                        preserve_failure(&report_path, &context, &run);
                        failures.push(error)
                    }
                    None => stats.ok += 1,
                }
            }
            Err(error) => {
                preserve_failure(&report_path, &context, &run);
                failures.push(format!("{context}: {error}"))
            }
        }
        let (class, shared) = carriers();
        stats.class_carriers += class;
        stats.shared_carriers += shared;
        if name == "valid-sharing" {
            reset_carriers();
            if let Either::Right(module) = run.candidate.as_ref() { let _ = sharing_check(module); }
            let (class, shared) = carriers();
            stats.class_carriers += class;
            stats.shared_carriers += shared;
        }
    }

    let mut report = String::from("{\"kinds\":{");
    for (index, (kind, stats)) in kinds.iter().enumerate() {
        if index > 0 { report.push(','); }
        report.push_str(&format!(
            "{}:{{\"total\":{},\"ok\":{},\"right\":{},\"left\":{},\"native\":{},\"declined\":{},\
             \"class_carriers\":{},\"shared_carriers\":{}}}",
            json_string(kind), stats.total, stats.ok, stats.right, stats.left, stats.fast,
            stats.declined, stats.class_carriers, stats.shared_carriers));
    }
    report.push_str("},\"failures\":[");
    for (index, failure) in failures.iter().enumerate() {
        if index > 0 { report.push(','); }
        report.push_str(&json_string(failure));
    }
    report.push_str("]}\n");
    std::fs::write(&report_path, &report).expect("write report");

    let total: usize = kinds.iter().map(|(_, stats)| stats.total).sum();
    let ok: usize = kinds.iter().map(|(_, stats)| stats.ok).sum();
    println!("text-native differential: {ok}/{total} cases equal");
    for (kind, stats) in &kinds {
        println!(
            "  {kind}: total {} ok {} right {} left {} native {} declined {} class {} shared {}",
            stats.total, stats.ok, stats.right, stats.left, stats.fast, stats.declined,
            stats.class_carriers, stats.shared_carriers);
    }
    if !failures.is_empty() {
        eprintln!("{} failures:", failures.len());
        for failure in &failures { eprintln!("  {failure}"); }
        std::process::exit(1);
    }
}
