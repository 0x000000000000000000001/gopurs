package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

const inputPackFormat = "gopurs/frozen-inputs-v1"

type frozenInput struct {
	Name     string `json:"name"`
	Resolved string `json:"resolved"`
	Object   Object `json:"object"`
	Offset   int64  `json:"offset"`
}

type frozenInputIndex struct {
	Format  string        `json:"format"`
	Bytes   int64         `json:"bytes"`
	Modules []frozenInput `json:"modules"`
}

func packedInputs(spec Spec) bool {
	for _, option := range spec.Options {
		if option.Name == "input-layout" {
			return option.Value == "packed-v1"
		}
	}
	return false
}

func inputLayout(spec Spec) string {
	if packedInputs(spec) {
		return "packed-v1"
	}
	return ""
}

// Readers hash their own observed bytes concurrently. Only append is serialized;
// at most eight source buffers are live and no per-module inode is created.
type inputPackWriter struct {
	file  *os.File
	mu    sync.Mutex
	bytes int64
}

func newInputPackWriter(stage string) (*inputPackWriter, error) {
	file, err := os.OpenFile(filepath.Join(stage, "input.pack"), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		return nil, err
	}
	return &inputPackWriter{file: file}, nil
}

func (writer *inputPackWriter) append(f Fingerprint, data []byte) (frozenInput, error) {
	if !validModule(f.Name) {
		return frozenInput{}, fmt.Errorf("invalid frozen input module: %s", f.Name)
	}
	writer.mu.Lock()
	defer writer.mu.Unlock()
	entry := frozenInput{Name: f.Name, Resolved: f.Resolved, Object: Object{f.Digest, int64(len(data))}, Offset: writer.bytes}
	n, err := writer.file.Write(data)
	writer.bytes += int64(n)
	if err == nil && n != len(data) {
		err = io.ErrShortWrite
	}
	return entry, err
}

func (writer *inputPackWriter) finish(stage string, modules []frozenInput) error {
	if err := writer.file.Chmod(0400); err != nil {
		return err
	}
	if err := writer.file.Close(); err != nil {
		return err
	}
	data, err := json.Marshal(frozenInputIndex{Format: inputPackFormat, Bytes: writer.bytes, Modules: modules})
	if err != nil {
		return err
	}
	return writePrivateFile(filepath.Join(stage, "input-index.json"), data)
}

func openFrozenInputFile(path string) (*os.File, error) {
	before, err := os.Lstat(path)
	if err != nil || !before.Mode().IsRegular() {
		return nil, fmt.Errorf("frozen input is not a regular file: %s", path)
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	info, err := file.Stat()
	after, afterErr := os.Lstat(path)
	if err != nil || afterErr != nil || !after.Mode().IsRegular() || !os.SameFile(before, info) || !os.SameFile(info, after) {
		file.Close()
		return nil, fmt.Errorf("frozen input file changed while opening: %s", path)
	}
	return file, nil
}

// The complete index is bound to the original TAST snapshot, including resolved
// symlink targets. A damaged offset/hash inventory cannot redefine that input.
func readFrozenInputIndex(stage string, state buildSession) (*frozenInputIndex, error) {
	file, err := openFrozenInputFile(filepath.Join(stage, "input-index.json"))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var index frozenInputIndex
	if err := strictJSON(data, &index); err != nil {
		return nil, err
	}
	if index.Format != inputPackFormat || index.Bytes < 0 || len(index.Modules) != state.Count {
		return nil, fmt.Errorf("invalid frozen input inventory")
	}
	parts := make([]string, 0, 3*len(index.Modules))
	previous := ""
	for _, entry := range index.Modules {
		if !validModule(entry.Name) || entry.Name <= previous || !filepath.IsAbs(entry.Resolved) ||
			!validObject(entry.Object) || entry.Offset < 0 || entry.Offset > index.Bytes || entry.Object.Bytes > index.Bytes-entry.Offset {
			return nil, fmt.Errorf("invalid frozen input entry: %s", entry.Name)
		}
		previous = entry.Name
		parts = append(parts, entry.Name, entry.Resolved, entry.Object.SHA256)
	}
	expected := ""
	for _, input := range state.Snapshot.Inputs {
		if input.Name == "tast" && input.Kind == "tast" {
			if expected != "" {
				return nil, fmt.Errorf("ambiguous frozen TAST snapshot")
			}
			expected = input.Digest
		}
	}
	if expected == "" || Key("tast", parts...) != expected {
		return nil, fmt.Errorf("frozen input inventory does not match TAST snapshot")
	}
	// Append completion order is independent of canonical module order. Require
	// every byte to belong to exactly one indexed source, without gaps/overlap.
	ordered := append([]frozenInput{}, index.Modules...)
	sort.Slice(ordered, func(i, j int) bool {
		if ordered[i].Offset == ordered[j].Offset {
			return ordered[i].Object.Bytes < ordered[j].Object.Bytes
		}
		return ordered[i].Offset < ordered[j].Offset
	})
	end := int64(0)
	for _, entry := range ordered {
		if entry.Offset != end {
			return nil, fmt.Errorf("frozen input ranges overlap or have gaps")
		}
		end += entry.Object.Bytes
	}
	if end != index.Bytes {
		return nil, fmt.Errorf("incomplete frozen input ranges")
	}
	return &index, nil
}

func materializeInputs(stage string, state buildSession, owners []string, all bool) error {
	if !packedInputs(state.Spec) {
		return nil // Older clients/sessions already have a complete input tree.
	}
	if all && len(owners) != 0 {
		return fmt.Errorf("all frozen inputs cannot be combined with named inputs")
	}
	if err := privateDirectory(stage); err != nil {
		return err
	}
	index, err := readFrozenInputIndex(stage, state)
	if err != nil {
		return err
	}
	byName := make(map[string]frozenInput, len(index.Modules))
	for _, entry := range index.Modules {
		byName[entry.Name] = entry
	}
	selected := []frozenInput{}
	if all {
		selected = index.Modules
	} else {
		seen := map[string]bool{}
		for _, name := range owners {
			entry, ok := byName[name]
			if !validModule(name) || !ok || seen[name] {
				return fmt.Errorf("invalid or duplicate frozen input owner: %s", name)
			}
			seen[name] = true
			selected = append(selected, entry)
		}
	}
	if len(selected) == 0 {
		return nil
	}
	input := filepath.Join(stage, "input")
	if err := privateDirectory(input); err != nil {
		return err
	}
	pack, err := openFrozenInputFile(filepath.Join(stage, "input.pack"))
	if err != nil {
		return err
	}
	defer pack.Close()
	info, err := pack.Stat()
	if err != nil || info.Size() != index.Bytes {
		return fmt.Errorf("frozen input pack size changed")
	}
	return boundedEach(len(selected), func(i int) error {
		entry := selected[i]
		directory := filepath.Join(input, entry.Name)
		if err := privateDirectory(directory); err != nil {
			return err
		}
		destination := filepath.Join(directory, "corefn.json")
		if info, err := os.Lstat(destination); err == nil && !info.Mode().IsRegular() {
			return fmt.Errorf("frozen input destination is not a regular file: %s", entry.Name)
		} else if err != nil && !os.IsNotExist(err) {
			return err
		}
		if fileMatches(destination, entry.Object) {
			return nil
		}
		data := make([]byte, int(entry.Object.Bytes))
		if _, err := pack.ReadAt(data, entry.Offset); err != nil && !(err == io.EOF && len(data) == 0) {
			return fmt.Errorf("reading frozen input %s: %w", entry.Name, err)
		}
		if digest(data) != entry.Object.SHA256 {
			return fmt.Errorf("frozen input bytes changed: %s", entry.Name)
		}
		return writePrivateFile(destination, data)
	})
}
