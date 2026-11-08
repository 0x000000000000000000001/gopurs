package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"syscall"
	"time"
)

type probeInput struct {
	Root     string     `json:"root"`
	Previous []Artifact `json:"previous"`
	Entries  []struct {
		Name   string `json:"name"`
		Source string `json:"source"`
		Object Object `json:"object"`
	} `json:"entries"`
}

// This private executable links the frozen store sources. It measures marginal
// object validation/staging plus one durable index, not a complete build commit.
func runProbe() error {
	var input probeInput
	if err := json.NewDecoder(os.Stdin).Decode(&input); err != nil {
		return err
	}
	if !filepath.IsAbs(input.Root) {
		return fmt.Errorf("absolute private store required")
	}
	s := store{root: input.Root, cache: filepath.Join(input.Root, ".gopurs-cache")}
	if err := os.MkdirAll(filepath.Join(s.cache, "objects"), 0700); err != nil {
		return err
	}
	validated := map[Artifact]bool{}
	for _, previous := range input.Previous {
		if _, err := s.readObject(previous.Object); err != nil {
			return err
		}
		validated[previous] = true
	}
	runtime.GC()
	started := time.Now()
	lock, err := os.OpenFile(filepath.Join(s.cache, "write.lock"), os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return err
	}
	defer lock.Close()
	if err := syscall.Flock(int(lock.Fd()), syscall.LOCK_EX); err != nil {
		return err
	}
	defer syscall.Flock(int(lock.Fd()), syscall.LOCK_UN)
	result := make([]Artifact, len(input.Entries))
	if err := boundedEach(len(input.Entries), func(index int) error {
		entry := input.Entries[index]
		value, err := s.stageArtifact("", StagedArtifact{Source: entry.Source, Codec: "p00/probe-v1"}, validated)
		if err != nil {
			return err
		}
		if value.Object != entry.Object {
			return fmt.Errorf("staged object differs: %s", entry.Name)
		}
		result[index] = value
		return nil
	}); err != nil {
		return err
	}
	if err := syncDirectory(filepath.Join(s.cache, "objects")); err != nil {
		return err
	}
	index, err := json.Marshal(result)
	if err != nil {
		return err
	}
	if err := atomicWrite(filepath.Join(s.cache, "probe-index.json"), index, 0600); err != nil {
		return err
	}
	seconds := time.Since(started).Seconds()
	var bytes int64
	unique := map[string]bool{}
	for _, value := range result {
		if !unique[value.Object.SHA256] {
			bytes += value.Object.Bytes
			unique[value.Object.SHA256] = true
		}
	}
	return json.NewEncoder(os.Stdout).Encode(struct {
		Seconds   float64    `json:"seconds"`
		Bytes     int64      `json:"bytes"`
		Unique    int        `json:"unique"`
		Artifacts []Artifact `json:"artifacts"`
	}{seconds, bytes, len(unique), result})
}

func main() {
	runtime.MemProfileRate = 0
	if err := runProbe(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
