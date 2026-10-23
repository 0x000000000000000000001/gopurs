package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"testing"
)

func TestFingerprintStreamAndFrozenObservation(t *testing.T) {
	root := t.TempDir()
	data := bytes.Repeat([]byte("source bytes\x00\n"), 17000)
	path := filepath.Join(root, "source")
	put(t, path, string(data))
	put(t, filepath.Join(root, "empty"), "")
	if err := os.Symlink("source", filepath.Join(root, "link")); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink("absent", filepath.Join(root, "dangling")); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"source", "link", "empty", "absent", "dangling"} {
		input := Input{Name: name, Kind: "file", Path: name}
		frozen, content, _, err := readObservation(root, input)
		if err != nil {
			t.Fatal(err)
		}
		actual, err := fingerprint(root, input)
		if err != nil || !reflect.DeepEqual(actual, frozen) {
			t.Fatalf("stream/frozen observation differs for %s: %+v %+v %v", name, actual, frozen, err)
		}
		if frozen.State == "file" && digest(content) != frozen.Digest {
			t.Fatalf("consumer bytes do not match the fingerprint: %s", name)
		}
		if (name == "source" || name == "link") && !bytes.Equal(content, data) {
			t.Fatalf("freezing did not retain all bytes: %s", name)
		}
	}
	input := Input{Name: "source", Kind: "file", Path: "link"}
	before, err := fingerprint(root, input)
	if err != nil {
		t.Fatal(err)
	}
	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	data[65536] ^= 1
	put(t, path, string(data))
	if err := os.Chtimes(path, info.ModTime(), info.ModTime()); err != nil {
		t.Fatal(err)
	}
	changed, err := fingerprint(root, input)
	if err != nil || changed.Digest == before.Digest {
		t.Fatalf("same-size/mtime source edit was not observed: %v", err)
	}
	data[65536] ^= 1
	put(t, path, string(data))
	reverted, err := fingerprint(root, input)
	if err != nil || reverted != before {
		t.Fatalf("edit and revert did not restore exact observation: %v", err)
	}
	if _, err := fingerprint(root, Input{Name: "directory", Kind: "file", Path: root}); err == nil {
		t.Fatal("accepted a directory as a file")
	}
}

func TestVerifyObjectStreamIntegrity(t *testing.T) {
	data := bytes.Repeat([]byte("object bytes\n"), 17000)
	object := Object{digest(data), int64(len(data))}
	for _, packed := range []bool{false, true} {
		for _, mode := range []string{"valid", "corrupt", "truncated", "appended", "missing", "wrong-length"} {
			t.Run(fmt.Sprintf("packed-%t/%s", packed, mode), func(t *testing.T) {
				s := store{cache: t.TempDir()}
				if err := s.ensurePacks(); err != nil {
					t.Fatal(err)
				}
				t.Cleanup(func() {
					for _, file := range s.packs.files {
						_ = file.Close()
					}
				})
				path := filepath.Join(s.cache, "objects", object.SHA256)
				content := append([]byte{}, data...)
				offset := 0
				if packed {
					content = append([]byte("prefix\n"), content...)
					offset = len("prefix\n")
					path = filepath.Join(s.cache, "packs", object.SHA256+".pack")
					s.packs.entries[object.SHA256] = packEntry{Pack: object.SHA256, Offset: int64(offset), Bytes: object.Bytes}
				}
				request := object
				switch mode {
				case "corrupt":
					content[offset+65536] ^= 1
				case "truncated":
					content = content[:len(content)-1]
				case "appended":
					content = append(content, 'x')
				case "wrong-length":
					request.Bytes--
				}
				if mode != "missing" {
					put(t, path, string(content))
				}
				err := s.verifyObject(request)
				valid := mode == "valid" || (packed && mode == "appended")
				if valid && err != nil || !valid && (err == nil || !strings.Contains(err.Error(), "missing/corrupt object")) {
					t.Fatalf("valid=%t: %v", valid, err)
				}
			})
		}
	}
}

func TestVerifyObjectStreamRejectsLooseSymlink(t *testing.T) {
	s := store{cache: t.TempDir()}
	data := []byte("unchanged content")
	object := Object{digest(data), int64(len(data))}
	target := filepath.Join(t.TempDir(), "target")
	put(t, target, string(data))
	if err := os.Mkdir(filepath.Join(s.cache, "objects"), 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(target, filepath.Join(s.cache, "objects", object.SHA256)); err != nil {
		t.Fatal(err)
	}
	if err := s.verifyObject(object); err == nil {
		t.Fatal("accepted a loose object symlink")
	}
}

func TestObjectStreamScratchIsIndependentAcrossReaders(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			data := bytes.Repeat([]byte{byte(i)}, 200000+i)
			for j := 0; j < 3; j++ {
				actual, count, err := digestReader(bytes.NewReader(data))
				if err != nil || count != int64(len(data)) || actual != digest(data) {
					t.Errorf("reader %d did not hash its own complete bytes: %v", i, err)
				}
			}
		}(i)
	}
	wg.Wait()
}

type failingDigestReader struct{}

func (failingDigestReader) Read(buffer []byte) (int, error) {
	return copy(buffer, "partial"), io.ErrUnexpectedEOF
}

func TestDigestReaderPropagatesReadFailure(t *testing.T) {
	hash, count, err := digestReader(failingDigestReader{})
	if err != io.ErrUnexpectedEOF || hash != "" || count != int64(len("partial")) {
		t.Fatalf("read error was treated as a complete digest: %q %d %v", hash, count, err)
	}
	data := []byte("next complete read")
	hash, count, err = digestReader(bytes.NewReader(data))
	if err != nil || count != int64(len(data)) || hash != digest(data) {
		t.Fatalf("failed reader contaminated the next digest: %q %d %v", hash, count, err)
	}
}
