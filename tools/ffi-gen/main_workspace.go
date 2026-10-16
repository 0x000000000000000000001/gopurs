//go:build packagecli && !js && !carchive

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	request, err := io.ReadAll(os.Stdin)
	if err == nil {
		var result string
		result, err = PackageFiles(string(request))
		if err == nil {
			fmt.Println(result)
			return
		}
	}
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
