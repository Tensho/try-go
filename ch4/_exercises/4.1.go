package main

import (
  "fmt"
  "crypto/sha256"
)

func main() {
  str1, str2 := "x", "X"
  c1, c2 := sha256.Sum256([]byte(str1)), sha256.Sum256([]byte(str2))
  fmt.Printf("%q : %x\n", str1, c1)
  fmt.Printf("%q : %x\n", str2, c2)
  fmt.Println(diffcount(c1, c2))
}

func diffcount(h1 [32]uint8, h2 [32]uint8) int {
  var cnt int
  for i := 0; i < len(h1); i++ {
    d := h1[i] ^ h2[i]

    for i := uint(0); i < 8; i++ {
      if d & (1 << i) != 0 {
        cnt++
      }
    }
  }
  return cnt
}
