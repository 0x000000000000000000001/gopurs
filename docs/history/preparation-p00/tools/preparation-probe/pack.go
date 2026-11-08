package main

import (
	"bytes"
	"compress/gzip"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"syscall"
	"time"
)

const packSchema = "p01-pack-study-v1"

type packEntry struct {
	Kind   string `json:"kind"`
	Name   string `json:"name"`
	Source string `json:"source,omitempty"`
	Object Object `json:"object"`
}

type packGroup struct {
	ID      string      `json:"id"`
	Input   string      `json:"input"`
	Body    Object      `json:"body"`
	Members []packEntry `json:"members"`
}

type packManifest struct {
	Schema string      `json:"schema"`
	Layout string      `json:"layout"`
	Fence  string      `json:"fence"`
	Groups []packGroup `json:"groups"`
}

type packRequest struct {
	Action           string      `json:"action"`
	Root             string      `json:"root"`
	Layout           string      `json:"layout"`
	Fence            string      `json:"fence"`
	Entries          []packEntry `json:"entries"`
	Target           string      `json:"target"`
	FailBeforeCommit bool        `json:"failBeforeCommit"`
}

func packDigest(value string) bool {
	decoded, err := hex.DecodeString(value)
	return err == nil && len(decoded) == 32 && hex.EncodeToString(decoded) == value
}

func packKind(value string) bool {
	switch value {
	case "source", "prepared", "metadata", "direct", "callers", "transitive", "contributions", "compact":
		return true
	}
	return false
}

func packIdentity(entry packEntry) string {
	data, _ := json.Marshal([2]string{entry.Kind, entry.Name})
	return digest(data)
}

func packRead(s store) (*packManifest, error) {
	data, err := os.ReadFile(filepath.Join(s.cache, "pack-current.json"))
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var pointer Object
	if err := strictJSON(data, &pointer); err != nil {
		return nil, err
	}
	data, err = s.readObject(pointer)
	if err != nil {
		return nil, err
	}
	var manifest packManifest
	if err := strictJSON(data, &manifest); err != nil {
		return nil, err
	}
	if manifest.Schema != packSchema || !packDigest(manifest.Fence) || (manifest.Layout != "raw" && manifest.Layout != "packed64") {
		return nil, fmt.Errorf("invalid pack manifest")
	}
	seen, groups := map[string]bool{}, map[string]bool{}
	for _, group := range manifest.Groups {
		if groups[group.ID] || len(group.Members) == 0 || (manifest.Layout == "raw" && len(group.Members) != 1) {
			return nil, fmt.Errorf("invalid pack group")
		}
		groups[group.ID] = true
		input, _ := json.Marshal(group.Members)
		if digest(input) != group.Input {
			return nil, fmt.Errorf("pack group identity")
		}
		for _, member := range group.Members {
			key := packIdentity(member)
			if seen[key] || !packKind(member.Kind) || member.Name == "" || member.Source != "" || !packDigest(member.Object.SHA256) || member.Object.Bytes <= 0 || member.Object.Bytes > 512<<20 {
				return nil, fmt.Errorf("invalid pack member")
			}
			seen[key] = true
		}
	}
	return &manifest, nil
}

func packPublish(s store, request packRequest) (int, int64, error) {
	if request.Layout != "raw" && request.Layout != "packed64" {
		return 0, 0, fmt.Errorf("unknown layout")
	}
	previous, err := packRead(s)
	if err != nil {
		return 0, 0, err
	}
	priorGroups, validated := map[string]packGroup{}, map[Object]bool{}
	if previous != nil {
		unique := map[Object]bool{}
		for _, group := range previous.Groups {
			unique[group.Body] = true
			if previous.Layout == request.Layout {
				priorGroups[group.ID] = group
			}
			for _, member := range group.Members {
				validated[member.Object] = true
			}
		}
		objects := make([]Object, 0, len(unique))
		for object := range unique {
			objects = append(objects, object)
		}
		if err := boundedEach(len(objects), func(i int) error { _, err := s.readObject(objects[i]); return err }); err != nil {
			return 0, 0, err
		}
	}
	entries := append([]packEntry(nil), request.Entries...)
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Kind != entries[j].Kind {
			return entries[i].Kind < entries[j].Kind
		}
		return entries[i].Name < entries[j].Name
	})
	sources, grouped := map[string]string{}, map[string][]packEntry{}
	for _, entry := range entries {
		key := packIdentity(entry)
		if !packKind(entry.Kind) || entry.Name == "" || !filepath.IsAbs(entry.Source) || !packDigest(entry.Object.SHA256) || entry.Object.Bytes <= 0 || entry.Object.Bytes > 512<<20 || sources[key] != "" {
			return 0, 0, fmt.Errorf("invalid publication entry")
		}
		sources[key] = entry.Source
		entry.Source = ""
		id := entry.Kind + "/" + key
		if request.Layout == "packed64" {
			hash, _ := hex.DecodeString(key)
			id = fmt.Sprintf("%s/%02x", entry.Kind, hash[0]&63)
		}
		grouped[id] = append(grouped[id], entry)
	}
	ids := make([]string, 0, len(grouped))
	for id := range grouped {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	manifest := packManifest{Schema: packSchema, Layout: request.Layout, Fence: request.Fence, Groups: make([]packGroup, len(ids))}
	if err := boundedEach(len(ids), func(i int) error {
		id := ids[i]
		members := grouped[id]
		input, _ := json.Marshal(members)
		group := packGroup{ID: id, Input: digest(input), Members: members}
		if old, ok := priorGroups[id]; ok && old.Input == group.Input {
			group.Body = old.Body
		} else {
			var buffer bytes.Buffer
			var writer io.Writer = &buffer
			var zipped *gzip.Writer
			if request.Layout == "packed64" {
				zipped, _ = gzip.NewWriterLevel(&buffer, gzip.BestSpeed)
				writer = zipped
			}
			for _, member := range members {
				data, err := os.ReadFile(sources[packIdentity(member)])
				if err != nil {
					return err
				}
				if int64(len(data)) != member.Object.Bytes || digest(data) != member.Object.SHA256 {
					return fmt.Errorf("input object differs: %s", member.Name)
				}
				if !validated[member.Object] {
					if err := validateWire(data); err != nil {
						return err
					}
				}
				if _, err := writer.Write(data); err != nil {
					return err
				}
			}
			if zipped != nil {
				if err := zipped.Close(); err != nil {
					return err
				}
			}
			object, err := s.putStagedObject(buffer.Bytes())
			if err != nil {
				return err
			}
			group.Body = object
		}
		manifest.Groups[i] = group
		return nil
	}); err != nil {
		return 0, 0, err
	}
	data, err := json.Marshal(manifest)
	if err != nil {
		return 0, 0, err
	}
	object, err := s.putStagedObject(data)
	if err != nil {
		return 0, 0, err
	}
	if err := syncDirectory(filepath.Join(s.cache, "objects")); err != nil {
		return 0, 0, err
	}
	if request.FailBeforeCommit {
		return 0, 0, fmt.Errorf("injected pre-commit failure")
	}
	pointer, _ := json.Marshal(object)
	if err := atomicWrite(filepath.Join(s.cache, "pack-current.json"), pointer, 0600); err != nil {
		return 0, 0, err
	}
	unique := map[Object]bool{object: true}
	for _, group := range manifest.Groups {
		unique[group.Body] = true
	}
	var count int64
	for object := range unique {
		count += object.Bytes
	}
	return len(manifest.Groups), count + int64(len(pointer)), nil
}

func packRestore(s store, request packRequest) (bool, error) {
	manifest, err := packRead(s)
	if err != nil {
		return false, err
	}
	if manifest == nil || manifest.Fence != request.Fence {
		return false, nil
	}
	if !filepath.IsAbs(request.Target) {
		return false, fmt.Errorf("absolute restoration target required")
	}
	if err := os.Mkdir(request.Target, 0700); err != nil {
		return false, err
	}
	indices := map[string][]any{}
	for _, group := range manifest.Groups {
		for _, member := range group.Members {
			indices[member.Kind] = append(indices[member.Kind], []any{member.Name, []any{member.Object.SHA256, member.Object.Bytes}})
		}
	}
	for kind := range indices {
		if err := os.Mkdir(filepath.Join(request.Target, kind), 0700); err != nil {
			return false, err
		}
	}
	if err := boundedEach(len(manifest.Groups), func(i int) error {
		group := manifest.Groups[i]
		data, err := s.readObject(group.Body)
		if err != nil {
			return err
		}
		buffer := bytes.NewReader(data)
		var reader io.Reader = buffer
		var zipped *gzip.Reader
		if manifest.Layout == "packed64" {
			zipped, err = gzip.NewReader(buffer)
			if err != nil {
				return err
			}
			defer zipped.Close()
			zipped.Multistream(false)
			reader = zipped
		}
		for _, member := range group.Members {
			value := make([]byte, int(member.Object.Bytes))
			if _, err := io.ReadFull(reader, value); err != nil {
				return err
			}
			if digest(value) != member.Object.SHA256 {
				return fmt.Errorf("restored member hash differs")
			}
			name := digest([]byte(member.Name)) + ".json"
			if err := os.WriteFile(filepath.Join(request.Target, member.Kind, name), value, 0600); err != nil {
				return err
			}
		}
		trailing, err := io.ReadAll(io.LimitReader(reader, 1))
		if err != nil || len(trailing) != 0 || buffer.Len() != 0 {
			return fmt.Errorf("invalid pack length/checksum: %v", err)
		}
		return nil
	}); err != nil {
		return false, err
	}
	for kind, entries := range indices {
		sort.Slice(entries, func(i, j int) bool { return entries[i].([]any)[0].(string) < entries[j].([]any)[0].(string) })
		data, _ := json.Marshal(entries)
		if err := os.WriteFile(filepath.Join(request.Target, kind, "index.json"), data, 0600); err != nil {
			return false, err
		}
	}
	return true, nil
}

func runPackProbe() error {
	var request packRequest
	if err := json.NewDecoder(os.Stdin).Decode(&request); err != nil {
		return err
	}
	if !filepath.IsAbs(request.Root) || !packDigest(request.Fence) {
		return fmt.Errorf("absolute store and valid input fence required")
	}
	s := store{root: request.Root, cache: filepath.Join(request.Root, ".gopurs-cache")}
	if err := os.MkdirAll(filepath.Join(s.cache, "objects"), 0700); err != nil {
		return err
	}
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
	result := map[string]any{"status": "ok"}
	switch request.Action {
	case "publish":
		groups, count, err := packPublish(s, request)
		if err != nil {
			return err
		}
		result["groups"], result["bytes"] = groups, count
	case "restore":
		hit, err := packRestore(s, request)
		if err != nil {
			return err
		}
		if !hit {
			result["status"] = "miss"
		}
	default:
		return fmt.Errorf("unknown pack action")
	}
	result["seconds"] = time.Since(started).Seconds()
	return json.NewEncoder(os.Stdout).Encode(result)
}

func main() {
	runtime.MemProfileRate = 0
	if err := runPackProbe(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
