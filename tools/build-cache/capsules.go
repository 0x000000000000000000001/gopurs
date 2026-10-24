package main

import (
	"fmt"
	"os"
	"strings"
)

const bindingCapsuleCodec = "pbo/binding-capsule-v1"

type capsuleAnchor struct {
	Owner  string `json:"owner"`
	SHA256 string `json:"sha256"`
}

type capsuleCoverage struct {
	Owner   string `json:"owner"`
	Ident   string `json:"ident"`
	Present *bool  `json:"present"`
}

type bindingCapsule struct {
	Codec         string            `json:"codec"`
	Context       string            `json:"context"`
	Reader        string            `json:"reader"`
	Anchors       []capsuleAnchor   `json:"anchors"`
	Coverage      []capsuleCoverage `json:"coverage"`
	Payload       string            `json:"payload"`
	PayloadSHA256 string            `json:"payloadSha256"`
}

func moduleContext(spec Spec, snapshot Snapshot) (string, error) {
	options, err := sortedSettings(spec.Options)
	if err != nil {
		return "", err
	}
	return Key("module-context", snapshot.Compiler, Key("options", options...)), nil
}

func validCapsuleArtifact(part *Artifact) bool {
	return part != nil && part.Codec == bindingCapsuleCodec && validObject(part.Object)
}

// Optional accelerators are retained like other objects, but are deliberately
// absent from artifacts(): losing one must not invalidate mandatory PBO parts.
func retainedArtifacts(manifest Manifest) []Artifact {
	result := artifacts(manifest)
	for _, memo := range manifest.Modules {
		if validCapsuleArtifact(memo.Capsule) {
			result = append(result, *memo.Capsule)
		}
	}
	return result
}

func validateCapsule(data []byte, context, reader string, implementation func(string) string) bool {
	var capsule bindingCapsule
	if strictJSON(data, &capsule) != nil || capsule.Codec != bindingCapsuleCodec ||
		capsule.Context != context || capsule.Reader != reader || !validModule(reader) ||
		capsule.Anchors == nil || capsule.Coverage == nil ||
		!digestPattern.MatchString(capsule.PayloadSHA256) || digest([]byte(capsule.Payload)) != capsule.PayloadSHA256 ||
		validateWire([]byte(capsule.Payload)) != nil {
		return false
	}
	owners := map[string]bool{}
	for _, anchor := range capsule.Anchors {
		if !validModule(anchor.Owner) || owners[anchor.Owner] || !digestPattern.MatchString(anchor.SHA256) || implementation(anchor.Owner) != anchor.SHA256 {
			return false
		}
		owners[anchor.Owner] = true
	}
	covered := map[string]bool{}
	coveredOwners := map[string]bool{}
	for _, item := range capsule.Coverage {
		key := item.Owner + "\x00" + item.Ident
		if !owners[item.Owner] || item.Ident == "" || strings.ContainsRune(item.Ident, '\x00') || item.Present == nil || covered[key] {
			return false
		}
		covered[key] = true
		coveredOwners[item.Owner] = true
	}
	if len(coveredOwners) != len(owners) {
		return false
	}
	// PBO additionally validates canonical order, typed payload and exact
	// coverage before using any binding. Go never interprets the business IR.
	return true
}

func implementationHashes(modules []ModuleMemo) map[string]string {
	result := make(map[string]string, len(modules))
	for _, memo := range modules {
		if len(memo.Parts) > 1 && memo.Parts[1].Codec == moduleCodecs[1] {
			result[memo.Name] = memo.Parts[1].Object.SHA256
		}
	}
	return result
}

// A full producer has no candidates, but publication still owns the preceding
// committed manifest under write.lock. Keep only its existing capsules whose
// anchors also describe all newly staged implementations; create none otherwise.
func (s store) stageCapsules(req Request, snapshot Snapshot, previous *Manifest, manifest *Manifest, packer *packWriter) error {
	type candidate struct {
		part   *Artifact
		source *store
	}
	old := map[string][]candidate{}
	for _, origin := range []struct {
		manifest *Manifest
		store    *store
	}{{previous, &s}, {req.inherited, req.inheritedStore}} {
		if origin.manifest == nil || origin.store == nil {
			continue
		}
		for _, memo := range origin.manifest.Modules {
			if validCapsuleArtifact(memo.Capsule) {
				old[memo.Name] = append(old[memo.Name], candidate{memo.Capsule, origin.store})
			}
		}
	}
	needed := len(old) != 0
	for _, staged := range req.Modules {
		needed = needed || staged.Capsule != nil
	}
	if !needed {
		return nil
	}
	context, err := moduleContext(req.Spec, snapshot)
	if err != nil {
		return err
	}
	hashes := implementationHashes(manifest.Modules)
	lookup := func(name string) string { return hashes[name] }
	for i := range manifest.Modules {
		memo := &manifest.Modules[i]
		memo.Capsule = nil
		if staged := req.Modules[i].Capsule; staged != nil && staged.Codec == bindingCapsuleCodec {
			data, err := os.ReadFile(absolute(req.Spec.Workspace, staged.Source))
			if err == nil && validateCapsule(data, context, memo.Name, lookup) {
				object, err := packer.put(data)
				if err != nil {
					return err
				}
				memo.Capsule = &Artifact{Codec: bindingCapsuleCodec, Object: object}
				continue
			}
		}
		for _, prior := range old[memo.Name] {
			data, err := prior.source.validateObject(prior.part.Object)
			if err != nil || !validateCapsule(data, context, memo.Name, lookup) {
				continue
			}
			if prior.source != &s {
				if err := s.verifyObject(prior.part.Object); err != nil {
					// An indexed pack remains authoritative even for optional data.
					if _, indexed := s.packs.entries[prior.part.Object.SHA256]; indexed {
						continue
					}
					if _, err := packer.put(data); err != nil {
						return err
					}
				}
			}
			part := *prior.part
			memo.Capsule = &part
			break
		}
	}
	return nil
}

func readSessionCapsule(stage string, state buildSession, req SessionRequest, result *SessionResponse) error {
	if len(req.Owners) != 1 || !validModule(req.Owners[0]) {
		return fmt.Errorf("invalid capsule reader")
	}
	result.Status = "ready"
	if !lazyModules(state.Spec) {
		return nil
	}
	prior, err := getLazyPrior(stage, state)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	prior.mu.Lock()
	defer prior.mu.Unlock()
	name := req.Owners[0]
	part := prior.modules[name].Capsule
	if !validCapsuleArtifact(part) {
		return nil
	}
	data, err := prior.store.validateObject(part.Object)
	if err != nil {
		return nil
	}
	context, err := moduleContext(state.Spec, state.Snapshot)
	if err != nil {
		return err
	}
	lookup := func(owner string) string {
		memo := prior.modules[owner]
		if len(memo.Parts) > 1 && memo.Parts[1].Codec == moduleCodecs[1] {
			return memo.Parts[1].Object.SHA256
		}
		return ""
	}
	if validateCapsule(data, context, name, lookup) {
		text := string(data)
		result.Capsule = &text
	}
	return nil
}
