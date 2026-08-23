package main
import "fmt"

func makeStruct() *struct{ A int; B int } {
    return &struct{ A int; B int }{ A: 1, B: 2 }
}

func printStruct(s *struct{ A int; B int }) {
    fmt.Println(s.A, s.B)
}

func main() {
    printStruct(makeStruct())
}
