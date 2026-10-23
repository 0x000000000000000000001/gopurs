package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"sync"
)

const (
	packIndexFormat = "gopurs-publication-packs"
	packIndexSchema = 1
	packPublication = "GOPURS_PACK_PUBLICATION"
	packGroups      = 64
)

type packEntry struct {
	Pack   string `json:"pack"`
	Offset int64  `json:"offset"`
	Bytes  int64  `json:"bytes"`
}

type packIndexFile struct {
	Format  string               `json:"format"`
	Schema  int                  `json:"schema"`
	Objects map[string]packEntry `json:"objects"`
}

type packDB struct {
	mu      sync.Mutex
	entries map[string]packEntry
	files   map[string]*os.File
}

type stagedPackObject struct {
	sha    string
	data   []byte
	offset int64
}

type packGroup struct {
	objects []stagedPackObject
}

type packWriter struct {
	store   *store
	enabled bool
	mu      sync.Mutex
	groups  [packGroups]*packGroup
	pending map[string]bool
}

func packPublicationEnabled() bool {
	return os.Getenv(packPublication) != "0"
}

func validPackEntry(entry packEntry) bool {
	return digestPattern.MatchString(entry.Pack) && entry.Offset >= 0 && entry.Bytes >= 0
}

// A missing index is an ordinary empty cache. A corrupt index is treated as
// empty too: packed objects become unreachable and the next publication can
// rebuild them. This is a cache miss, never a trust decision.
func (s *store) ensurePacks() error {
	if s.packs != nil {
		return nil
	}
	s.packs = &packDB{entries: map[string]packEntry{}, files: map[string]*os.File{}}
	path := filepath.Join(s.cache, "packs", "index.json")
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	var index packIndexFile
	if strictJSON(data, &index) != nil || index.Format != packIndexFormat || index.Schema != packIndexSchema || index.Objects == nil {
		return nil
	}
	for name, entry := range index.Objects {
		if !digestPattern.MatchString(name) || !validPackEntry(entry) {
			return nil
		}
	}
	s.packs.entries = index.Objects
	return nil
}

func (s store) packFile(pack string) (*os.File, error) {
	if s.packs == nil {
		return nil, fmt.Errorf("missing/corrupt pack: %s", pack)
	}
	s.packs.mu.Lock()
	defer s.packs.mu.Unlock()
	if f := s.packs.files[pack]; f != nil {
		return f, nil
	}
	f, err := os.Open(filepath.Join(s.cache, "packs", pack+".pack"))
	if err != nil {
		return nil, err
	}
	s.packs.files[pack] = f
	return f, nil
}

func (s store) readPackedObject(entry packEntry, obj Object) ([]byte, error) {
	if !validPackEntry(entry) || entry.Bytes != obj.Bytes {
		return nil, fmt.Errorf("missing/corrupt object: %s", obj.SHA256)
	}
	f, err := s.packFile(entry.Pack)
	if err != nil {
		return nil, fmt.Errorf("missing/corrupt object: %s", obj.SHA256)
	}
	info, err := f.Stat()
	if err != nil || !info.Mode().IsRegular() || info.Size() < entry.Offset+entry.Bytes {
		return nil, fmt.Errorf("missing/corrupt object: %s", obj.SHA256)
	}
	data := make([]byte, int(entry.Bytes))
	if _, err := f.ReadAt(data, entry.Offset); err != nil {
		return nil, fmt.Errorf("missing/corrupt object: %s", obj.SHA256)
	}
	if digest(data) != obj.SHA256 {
		return nil, fmt.Errorf("missing/corrupt object: %s", obj.SHA256)
	}
	return data, nil
}

func newPackWriter(s *store) *packWriter {
	return &packWriter{store: s, enabled: packPublicationEnabled(), pending: map[string]bool{}}
}

// put keeps new private artifacts in memory until flush. Objects already
// present individually or in a committed pack are not rewritten. The writer is
// used only for module artifacts; public outputs keep their individual files.
func (w *packWriter) put(data []byte) (Object, error) {
	obj := Object{SHA256: digest(data), Bytes: int64(len(data))}
	if !w.enabled {
		return w.store.putStagedObject(data)
	}
	repairPacked := false
	if w.store.packs != nil {
		if entry, ok := w.store.packs.entries[obj.SHA256]; ok {
			// A lazy read can fail after lookup returned inputs-changed. The
			// full fallback must repair that indexed object from fresh bytes.
			// Keep the index immutable while producers run in parallel; flush
			// replaces the entry after every producer has joined.
			if _, err := w.store.readPackedObject(entry, obj); err == nil {
				return obj, nil
			}
			repairPacked = true
		}
	}
	if !repairPacked && fileMatches(filepath.Join(w.store.cache, "objects", obj.SHA256), obj) {
		return obj, nil
	}
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.pending[obj.SHA256] {
		return obj, nil
	}
	value, err := strconv.ParseUint(obj.SHA256[:2], 16, 8)
	if err != nil {
		return Object{}, fmt.Errorf("invalid object digest: %s", obj.SHA256)
	}
	group := int(value) % packGroups
	if w.groups[group] == nil {
		w.groups[group] = &packGroup{}
	}
	w.groups[group].objects = append(w.groups[group].objects, stagedPackObject{sha: obj.SHA256, data: data})
	w.pending[obj.SHA256] = true
	return obj, nil
}

func (w *packWriter) flush() error {
	if !w.enabled {
		return nil
	}
	directory := filepath.Join(w.store.cache, "packs")
	if err := privateDirectory(directory); err != nil {
		return err
	}
	changed := false
	for _, group := range w.groups {
		if group == nil || len(group.objects) == 0 {
			continue
		}
		sort.Slice(group.objects, func(i, j int) bool { return group.objects[i].sha < group.objects[j].sha })
		var content bytes.Buffer
		for i := range group.objects {
			group.objects[i].offset = int64(content.Len())
			content.Write(group.objects[i].data)
		}
		data := content.Bytes()
		pack := digest(data)
		path := filepath.Join(directory, pack+".pack")
		if !fileMatches(path, Object{SHA256: pack, Bytes: int64(len(data))}) {
			if err := atomicReplaceFile(path, data, 0600); err != nil {
				return err
			}
			// An existing handle still names the replaced inode. No producer
			// runs during flush; subsequent reads must open the repaired file.
			w.store.packs.mu.Lock()
			if file := w.store.packs.files[pack]; file != nil {
				_ = file.Close()
				delete(w.store.packs.files, pack)
			}
			w.store.packs.mu.Unlock()
		}
		for _, object := range group.objects {
			w.store.packs.entries[object.sha] = packEntry{Pack: pack, Offset: object.offset, Bytes: int64(len(object.data))}
		}
		changed = true
	}
	if !changed {
		return nil
	}
	return w.writeIndex()
}

func (s *store) writePackIndex() error {
	directory := filepath.Join(s.cache, "packs")
	if err := privateDirectory(directory); err != nil {
		return err
	}
	data, err := json.Marshal(packIndexFile{Format: packIndexFormat, Schema: packIndexSchema, Objects: s.packs.entries})
	if err != nil {
		return err
	}
	if err := atomicWrite(filepath.Join(directory, "index.json"), data, 0600); err != nil {
		return err
	}
	return syncDirectory(directory)
}

func (w *packWriter) writeIndex() error {
	return w.store.writePackIndex()
}

func (s *store) invalidatePack(pack string) {
	s.packs.mu.Lock()
	defer s.packs.mu.Unlock()
	for name, entry := range s.packs.entries {
		if entry.Pack == pack {
			delete(s.packs.entries, name)
		}
	}
	if file := s.packs.files[pack]; file != nil {
		_ = file.Close()
		delete(s.packs.files, pack)
	}
}

func (s *store) invalidateAllPacks() {
	if s.packs == nil {
		return
	}
	s.packs.mu.Lock()
	empty := len(s.packs.entries) == 0 && len(s.packs.files) == 0
	for _, file := range s.packs.files {
		_ = file.Close()
	}
	s.packs.entries = map[string]packEntry{}
	s.packs.files = map[string]*os.File{}
	s.packs.mu.Unlock()
	if !empty {
		_ = s.writePackIndex()
	}
}

func packGroupFor(sha string) int {
	value, err := strconv.ParseUint(sha[:2], 16, 8)
	if err != nil {
		return 0
	}
	return int(value) % packGroups
}

// Rewrite only groups that contain more than one pack or bytes no longer
// referenced. This bounds pack growth without a full-cache rewrite per edit.
func (s *store) compactPackGroups() (bool, error) {
	if len(s.packs.entries) == 0 {
		return false, nil
	}
	type groupState struct {
		shas      []string
		packs     map[string]bool
		fileBytes int64
		refBytes  int64
	}
	groups := map[int]*groupState{}
	sizes := map[string]int64{}
	for sha, entry := range s.packs.entries {
		group := groups[packGroupFor(sha)]
		if group == nil {
			group = &groupState{packs: map[string]bool{}}
			groups[packGroupFor(sha)] = group
		}
		group.shas = append(group.shas, sha)
		group.packs[entry.Pack] = true
		group.refBytes += entry.Bytes
		if _, ok := sizes[entry.Pack]; !ok {
			info, err := os.Stat(filepath.Join(s.cache, "packs", entry.Pack+".pack"))
			if err != nil {
				return false, err
			}
			sizes[entry.Pack] = info.Size()
		}
	}
	for _, group := range groups {
		for pack := range group.packs {
			group.fileBytes += sizes[pack]
		}
	}
	oldPacks := map[string]bool{}
	newPacks := map[string]bool{}
	changed := false
	for _, group := range groups {
		if len(group.packs) <= 1 && group.fileBytes == group.refBytes {
			continue
		}
		sort.Strings(group.shas)
		var content bytes.Buffer
		for _, sha := range group.shas {
			entry := s.packs.entries[sha]
			data, err := s.readObject(Object{SHA256: sha, Bytes: entry.Bytes})
			if err != nil {
				return false, err
			}
			offset := int64(content.Len())
			content.Write(data)
			s.packs.entries[sha] = packEntry{Pack: "", Offset: offset, Bytes: entry.Bytes}
		}
		pack := digest(content.Bytes())
		path := filepath.Join(s.cache, "packs", pack+".pack")
		if !fileMatches(path, Object{SHA256: pack, Bytes: int64(content.Len())}) {
			if err := atomicReplaceFile(path, content.Bytes(), 0600); err != nil {
				return false, err
			}
		}
		for _, sha := range group.shas {
			entry := s.packs.entries[sha]
			entry.Pack = pack
			s.packs.entries[sha] = entry
		}
		for pack := range group.packs {
			oldPacks[pack] = true
		}
		newPacks[pack] = true
		changed = true
	}
	if !changed {
		return false, nil
	}
	if err := s.writePackIndex(); err != nil {
		return false, err
	}
	directory := filepath.Join(s.cache, "packs")
	s.packs.mu.Lock()
	for pack := range oldPacks {
		if newPacks[pack] {
			continue
		}
		if file := s.packs.files[pack]; file != nil {
			_ = file.Close()
			delete(s.packs.files, pack)
		}
	}
	s.packs.mu.Unlock()
	for pack := range oldPacks {
		if newPacks[pack] {
			continue
		}
		if err := os.Remove(filepath.Join(directory, pack+".pack")); err != nil && !os.IsNotExist(err) {
			return false, err
		}
	}
	return true, syncDirectory(directory)
}

func (s *store) retirePackObjects(manifest *Manifest) error {
	if s.packs == nil || len(s.packs.entries) == 0 {
		return nil
	}
	retained := map[string]bool{}
	for _, artifact := range artifacts(*manifest) {
		retained[artifact.Object.SHA256] = true
	}
	packs := map[string]bool{}
	changed := false
	for name, entry := range s.packs.entries {
		if retained[name] {
			packs[entry.Pack] = true
			continue
		}
		delete(s.packs.entries, name)
		changed = true
	}
	directory := filepath.Join(s.cache, "packs")
	entries, err := os.ReadDir(directory)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	for _, entry := range entries {
		if entry.Type().IsRegular() && filepath.Ext(entry.Name()) == ".pack" && !packs[entry.Name()[:len(entry.Name())-len(".pack")]] {
			if err := os.Remove(filepath.Join(directory, entry.Name())); err != nil {
				return err
			}
			changed = true
		}
	}
	compacted, err := s.compactPackGroups()
	if err != nil {
		return err
	}
	if changed && !compacted {
		if err := os.MkdirAll(directory, 0700); err != nil {
			return err
		}
		if err := s.writePackIndex(); err != nil {
			return err
		}
	}
	return nil
}
