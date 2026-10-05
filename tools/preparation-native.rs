use std::sync::{Arc, Mutex};
use std::sync::atomic::{AtomicUsize, Ordering};
use std::time::Duration;
use purust_core::*;
use Purs_Effect_Aff::*;
use Purs_Gopurs_Preparation::Gopurs_Preparation_runPreparationJobs;

fn main() {
    std::env::set_var("PURUST_AFF_WORKERS", "8");
    let mut executions = 0;
    for configured in [-2, 0, 1, 3, 8, 999] {
        for count in [0, 1, 17, 129, 1025] {
            let calls: Arc<Vec<AtomicUsize>> = Arc::new((0..count).map(|_| AtomicUsize::new(0)).collect());
            let active = Arc::new(AtomicUsize::new(0));
            let peak = Arc::new(AtomicUsize::new(0));
            let tasks = Value::Array(Arc::new((0..count).map(|index| {
                let calls = calls.clone(); let active = active.clone(); let peak = peak.clone();
                Value::Func1(Func1::Shared(Arc::new(move |_| {
                    let concurrent = active.fetch_add(1, Ordering::SeqCst) + 1;
                    peak.fetch_max(concurrent, Ordering::SeqCst);
                    let round = calls[index].fetch_add(1, Ordering::SeqCst) + 1;
                    // Intentionally uneven pure task costs, with different finish
                    // orders. Timing is diagnostic only; assertions concern the
                    // scheduler's concurrency bound, multiplicity and ordering.
                    std::thread::sleep(Duration::from_micros(if index % 11 == 0 { 1500 } else { 100 }));
                    active.fetch_sub(1, Ordering::SeqCst);
                    Value::Int((round * 1000 + index) as i64)
                })))
            }).collect()));
            let action = Gopurs_Preparation_runPreparationJobs(configured, tasks.clone());
            assert!(calls.iter().all(|count| count.load(Ordering::SeqCst) == 0), "construction evaluated jobs");
            for round in 1..=2 {
                let output = Arc::new(Mutex::new(None));
                let stored = output.clone();
                let checked = Effect_Aff__map(Func1::Shared(Arc::new(move |value: Value| {
                    *stored.lock().unwrap() = Some(value.unwrap_array().iter().map(|v| v.unwrap_int()).collect::<Vec<_>>());
                    Value::Unit
                })), action.clone());
                purust_aff_run_main(|| Effect_Aff_launchAff_(checked).unwrap_func1()(Value::Unit));
                assert_eq!(output.lock().unwrap().as_ref().unwrap(), &(0..count).map(|i| (round * 1000 + i) as i64).collect::<Vec<_>>());
                assert!(calls.iter().all(|count| count.load(Ordering::SeqCst) == round));
                assert_eq!(active.load(Ordering::SeqCst), 0);
                assert!(peak.load(Ordering::SeqCst) <= configured.clamp(1, 8) as usize);
                executions += 1;
            }
            assert_eq!(tasks.unwrap_array().len(), count, "caller input was consumed");
            if configured > 1 && count >= 17 { assert!(peak.load(Ordering::SeqCst) > 1, "native jobs did not overlap"); }
        }
    }
    println!("Native preparation: {executions} runs; deferred/replayable actions, exact ordered results and bounded overlapping workers passed");
}
