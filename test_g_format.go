package main
import (
	"fmt"
	"strconv"
)
func main() {
	fmt.Println(strconv.FormatFloat(0.25996181067141905, 'g', -1, 64))
	fmt.Println(strconv.FormatFloat(0.25996181067141905, 'f', -1, 64))
}
