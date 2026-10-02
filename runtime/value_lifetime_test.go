package gopurs_runtime

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
)

//go:noinline
func retainedStrings(seed int) Value {
	text := strings.Repeat("padding", 1024) + fmt.Sprintf("%d:é😀\xed\xa0\x80\x00", seed)
	return Array([]Value{Str(""), Str(text[len("padding")*1024:]), Box(text)})
}

func TestPackedStringsRetainStorage(t *testing.T) {
	created := make(chan Value, 16)
	var creator sync.WaitGroup
	creator.Add(1)
	go func() {
		defer creator.Done()
		for i := 0; i < cap(created); i++ {
			created <- retainedStrings(i)
		}
		close(created)
	}()
	creator.Wait()
	for i := 0; i < 3; i++ {
		runtime.GC()
	}
	index := 0
	for values := range created {
		want := fmt.Sprintf("%d:é😀\xed\xa0\x80\x00", index)
		empty, substring := ArrayAccess(values, 0), ArrayAccess(values, 1)
		if StrValue(empty) != "" || empty.UnsafePtr != nil || empty.IntVal != 0 {
			t.Fatal("empty string is not canonical")
		}
		if StrValue(substring) != want || substring.IntVal != int64(len(want)) {
			t.Fatalf("substring %d lost byte contents or length", index)
		}
		if Unbox[string](ArrayAccess(values, 2)) != strings.Repeat("padding", 1024)+want {
			t.Fatalf("boxed string %d lost its backing storage", index)
		}
		index++
	}
}

type lifetimePayload struct {
	text    string
	numbers []int64
}

//go:noinline
func retainedContainers() Value {
	payload := &lifetimePayload{strings.Repeat("owned", 128), []int64{37}}
	callback := Func(func(arg Value) Value { return Int(payload.numbers[0] + arg.IntVal) })
	return RecordDict([]string{"array", "opaque", "callback"}, []Value{
		Array([]Value{Record(map[string]Value{"text": Str(payload.text)})}),
		Any(payload), callback,
	})
}

func TestContainersRetainPayloadsAfterCreatorExit(t *testing.T) {
	created := make(chan Value, 1)
	var creator sync.WaitGroup
	creator.Add(1)
	go func() { defer creator.Done(); created <- retainedContainers() }()
	value := <-created
	creator.Wait()
	var readers sync.WaitGroup
	for i := 0; i < 8; i++ {
		readers.Add(1)
		go func() {
			defer readers.Done()
			for j := 0; j < 16; j++ {
				if j%8 == 0 {
					runtime.GC()
				}
				text := RecordGet(ArrayAccess(RecordGet(value, "array"), 0), "text").StrVal()
				payload := RecordGet(value, "opaque").AnyVal().(*lifetimePayload)
				if text != strings.Repeat("owned", 128) || text != payload.text || payload.numbers[0] != 37 {
					t.Error("nested container or opaque payload lost its storage")
				}
				if Apply(RecordGet(value, "callback"), Int(5)).IntVal != 42 {
					t.Error("container lost its callback capture")
				}
			}
		}()
	}
	readers.Wait()
}

func TestContainerConstructorsShareCallerStorage(t *testing.T) {
	values := []Value{Int(1)}
	fields := map[string]Value{"x": Int(2)}
	keys, entries := []string{"x"}, []Value{Int(3)}
	array, record, dict := Array(values), Record(fields), RecordDict(keys, entries)
	// Probe aliasing before publication: the low-level constructors do not copy.
	values[0], fields["x"], entries[0], keys[0] = Int(11), Int(12), Int(13), "y"
	if ArrayAccess(array, 0).IntVal != 11 || RecordGet(record, "x").IntVal != 12 || RecordGet(dict, "y").IntVal != 13 {
		t.Fatal("container constructor copied caller-owned storage")
	}
	// The read view of an already foreign map also borrows its container.
	foreign := map[string]any{"x": int64(4)}
	view, ok := ReadJSONObject(Any(foreign))
	foreign["x"] = int64(14)
	if value, present := view.Lookup("x"); !ok || !present || value != int64(14) {
		t.Fatal("foreign object view did not borrow the map")
	}
}

func TestEventLoopWaitIncludesRetainedWork(t *testing.T) {
	const workers = 8
	var finished atomic.Int32
	start := make(chan struct{})
	for i := 0; i < workers; i++ {
		Retain()
		go func() {
			defer Release()
			<-start
			finished.Add(1)
		}()
	}
	waited := make(chan struct{})
	go func() { EventLoopWait(); close(waited) }()
	select {
	case <-waited:
		t.Fatal("wait returned with retained work outstanding")
	default:
	}
	close(start)
	<-waited
	if finished.Load() != workers {
		t.Fatal("wait returned before retained work completed")
	}
	// A completed generation leaves the wait group reusable.
	Retain()
	Release()
	EventLoopWait()
}
