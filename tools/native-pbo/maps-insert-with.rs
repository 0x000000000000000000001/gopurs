// Noncommutative merge oracle: existing value must precede incoming value,
// and the callback must only run when the key is already present.
fn check_native_insert_with() {
    use Purs_Data_Map_Internal::Data_Map_Internal_insertWith;
    use Purs_PureScript_Backend_Optimizer_NativeMaps::PureScript_Backend_Optimizer_NativeMaps_insertWithStringImpl as insert_with;
    use std::sync::atomic::{AtomicUsize, Ordering};
    let keys: Vec<_> = (0..70).map(|i| Value::String(purust_string_from_utf8(&format!("key{i}"))))
        .chain(["", "A", "a\0b", "é", "😀", "\u{e000}"].iter().map(|s| Value::String(purust_string_from_utf8(s))))
        .chain([Value::String(purust_string_from_utf16(&[0xd800])),
                Value::String(purust_string_from_utf16(&[0xdc00]))]).collect();
    let ord = Purs_Data_Ord::Data_Ord_ordString();
    let merge = Func2::Static(|old: Value, new: Value| Value::Int(old.unwrap_int() - 2 * new.unwrap_int()));
    for seed in 0..5_u64 {
        let calls = Rc::new(AtomicUsize::new(0));
        let counter = calls.clone();
        let counted_merge = Func2::Shared(Rc::new(move |old: Value, new: Value| {
            counter.fetch_add(1, Ordering::Relaxed);
            Value::Int(old.unwrap_int() - 2 * new.unwrap_int())
        }));
        let mut native = Rc::new(Map::Leaf);
        let mut oracle = native.clone();
        let mut expected = vec![None; keys.len()];
        let mut expected_calls = 0;
        let mut previous = Vec::new();
        let mut random = seed + 17;
        for index in 0..800 {
            random = random.wrapping_mul(6364136223846793005).wrapping_add(1);
            let slot = (random >> 32) as usize % keys.len();
            let key = keys[slot].clone();
            if index % 71 == 0 { previous.push((native.clone(), snapshot(&native))); }
            expected[slot] = Some(match expected[slot] {
                Some(old) => { expected_calls += 1; old - 2 * index },
                None => index,
            });
            native = insert_with(Value::Unit, counted_merge.clone(), key.clone(), Value::Int(index), native);
            oracle = Data_Map_Internal_insertWith(ord.clone(), merge.clone(), key, Value::Int(index), oracle);
            assert_eq!(snapshot(&native), snapshot(&oracle), "insertWith shape and value order");
            assert_eq!(calls.load(Ordering::Relaxed), expected_calls);
            if index % 31 == 0 {
                for (key, value) in keys.iter().zip(&expected) {
                    assert_eq!(lookup_result(&Data_Map_Internal_lookup(ord.clone(), key.clone(), native.clone())), *value);
                }
            }
        }
        for (map, before) in previous { assert_eq!(snapshot(&map), before, "old map mutated"); }
    }
    println!("Native Maps: 4000 insertWith updates, noncommutative merge order, callback counts, Unicode and persistence passed");
}
