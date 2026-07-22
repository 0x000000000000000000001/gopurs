package main
func main() {
  lbl:
  for {
    if false { continue lbl }
    break
  }
}
