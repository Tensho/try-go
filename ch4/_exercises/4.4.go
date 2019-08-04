package main

import (
  "fmt"
)

func main() {
  a := []int{0, 1, 2, 3, 4, 5}
  n := 4
  fmt.Printf("%v << %v\n", a, n)
  b := rotate(a, n)
  fmt.Println(b)
}

// Left rotate
// [0 1 2 3 4 5] << 4
// [4 5 0 1 2 3]
func rotate(a []int, n int) []int {
	head, tail := a[:n], a[n:]
  return append(tail, head...)
}
