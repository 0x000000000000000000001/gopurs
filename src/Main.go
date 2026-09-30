package Main

import "os"

// Main calls this after compilation cleanup and the synchronous Go diagnostic.
// The native entry point does not read Node.Process's stored exit code.
func ExitFailure(_ any) any {
	os.Exit(1)
	return nil
}
