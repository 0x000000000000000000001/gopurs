package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestSessionPreloadIsOptionalAndRequestLocal(t *testing.T) {
	for _, kind := range []string{"hit", "forced", "cold", "miss", "panic", "begin-error"} {
		t.Run(kind, func(t *testing.T) {
			spec, initial, state := driverFixture(t)
			finishDriver(t, &initial, &state, []SourceModule{{Name: "Main", Path: "src/Main.purs", HasForeign: true}}, false)
			if kind == "cold" {
				if err := os.Remove(filepath.Join(spec.Output, ".gopurs-cache/recipe.json")); err != nil {
					t.Fatal(err)
				}
			}
			if kind == "miss" || kind == "panic" || kind == "begin-error" {
				put(t, filepath.Join(spec.Output, "Main/corefn.json"), "changed bytes")
			}
			calls := 0
			var seed struct{ Root, CacheRoot, Context string }
			callback := func(text string) {
				calls++
				if err := json.Unmarshal([]byte(text), &seed); err != nil {
					t.Fatal(err)
				}
				if seed.Root != spec.Output || seed.CacheRoot != cacheDirectory(spec) || seed.Context == "" {
					t.Fatal("incomplete speculative seed", seed)
				}
				// Callback runs after lookup has released its store lock.
				if _, err := execute(Request{Schema: Schema, Operation: "snapshot", Spec: spec}); err != nil {
					t.Fatal(err)
				}
				if kind == "panic" {
					panic("optional preload unavailable")
				}
				if kind == "begin-error" {
					parent := filepath.Join(spec.Output, ".gopurs-cache/stages")
					if err := os.RemoveAll(parent); err != nil {
						t.Fatal(err)
					}
					put(t, parent, "blocks private directory")
				}
			}
			result := SessionResponse{FFI: []FrozenFFI{}}
			err := beginSessionOnMiss(spec, kind == "forced", &result, callback)
			if result.Stage != "" {
				t.Cleanup(func() { os.RemoveAll(result.Stage) })
			}
			wanted := 0
			if kind == "miss" || kind == "panic" || kind == "begin-error" {
				wanted = 1
			}
			if calls != wanted {
				t.Fatalf("notifications %d != %d", calls, wanted)
			}
			if kind == "begin-error" {
				if err == nil {
					t.Fatal("expected begin failure after notification")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if kind == "hit" {
				if result.Status != "hit" || result.Stage != "" {
					t.Fatal(result)
				}
			} else if result.Status != "miss" {
				t.Fatal(result)
			}
			if calls == 1 && seed.Context != result.Context {
				t.Fatal("seed differs from accepted session context")
			}
		})
	}
}
