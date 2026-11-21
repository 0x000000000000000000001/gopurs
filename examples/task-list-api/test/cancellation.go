package cancellation

import (
	"fmt"
	"io"
	"os"
)

// The compiled PureScript handler waits here in an EffectFnAff worker. Reading
// stdin is solely the test's synchronization gate, not part of the HTTP module.
func AwaitHandlerRelease() error {
	var token [1]byte
	if _, err := io.ReadFull(os.Stdin, token[:]); err != nil {
		return err
	}
	if token[0] != '\n' {
		return fmt.Errorf("unexpected handler release token: %q", token[0])
	}
	return nil
}
