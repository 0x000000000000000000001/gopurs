package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	request, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(Exchange(string(request)))
}
