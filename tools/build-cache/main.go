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
	if len(os.Args) == 2 && os.Args[1] == "--session" {
		fmt.Println(SessionExchange(string(request)))
	} else {
		fmt.Println(Exchange(string(request)))
	}
}
