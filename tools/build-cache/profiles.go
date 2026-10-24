package main

import (
	"os"
	"path/filepath"
)

// Profiles keep their reusable state independently, while publishing to the
// same owned output tree. The historical release namespace stays unchanged.
func cacheDirectory(spec Spec) string {
	root := filepath.Join(spec.Output, ".gopurs-cache")
	for _, option := range spec.Options {
		// Keep the earlier policy in the development namespace too. Its
		// version remains part of the build/module keys, so it cannot reuse
		// results produced by the current policy.
		if option.Name == "optimization-profile" && (option.Value == "dev-v1" || option.Value == "dev-v2" || option.Value == "dev-v3" || option.Value == "dev-v4" || option.Value == "dev-v5") {
			return filepath.Join(root, "profiles", "dev")
		}
	}
	return root
}

func publicLockPath(output string) string {
	return filepath.Join(output, ".gopurs-cache", "write.lock")
}

func prepareCacheDirectories(spec Spec) error {
	root := filepath.Join(spec.Output, ".gopurs-cache")
	directories := []string{root}
	cache := cacheDirectory(spec)
	if cache != root {
		directories = append(directories, filepath.Join(root, "profiles"), cache)
	}
	directories = append(directories, filepath.Join(cache, "objects"))
	for _, directory := range directories {
		if err := privateDirectory(directory); err != nil {
			return err
		}
	}
	return nil
}

// A disposable or older ownership journal can omit files published by the
// other profile. A verified committed manifest remains ownership evidence;
// absent historical paths do not describe the current public generation.
// The caller holds the shared publication lock throughout this observation.
func (s store) otherProfileOutputs(workspace string, current *Manifest) ([]OwnedOutput, error) {
	root := filepath.Join(s.root, ".gopurs-cache")
	known := map[string]bool{}
	if current != nil && current.Snapshot.Output == s.root && current.Snapshot.Workspace == workspace {
		for _, out := range current.Outputs {
			known[out.Path] = true
		}
	}
	var outputs []OwnedOutput
	for _, cache := range []string{root, filepath.Join(root, "profiles", "dev")} {
		if cache == s.cache {
			continue
		}
		other := store{root: s.root, cache: cache}
		manifest, _ := other.readManifest()
		if manifest == nil || manifest.Snapshot.Output != s.root || manifest.Snapshot.Workspace != workspace {
			continue
		}
		for _, out := range manifest.Outputs {
			// The selected manifest already validates these public paths and
			// supplies their ownership on recovery. Avoid duplicate filesystem
			// observations for the common case where profiles share the layout.
			if known[out.Path] {
				continue
			}
			path, err := s.destination(out.Path)
			if err != nil {
				return nil, err
			}
			if _, err := os.Lstat(path); os.IsNotExist(err) {
				continue
			} else if err != nil {
				return nil, err
			}
			// Keep the recorded digest, never the observed external bytes.
			// Publication verifies obsolete bytes before removing anything.
			outputs = append(outputs, out)
		}
	}
	return outputs, nil
}
