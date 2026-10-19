package main

import (
	"os"
	"testing"
)

// The unit suite keeps the legacy individual-object path by default. Pack
// tests opt in with t.Setenv, and the production default remains packed.
func TestMain(m *testing.M) {
	if os.Getenv(packPublication) == "" {
		_ = os.Setenv(packPublication, "0")
	}
	os.Exit(m.Run())
}
