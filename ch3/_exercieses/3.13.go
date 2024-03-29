package main

import (
  "fmt"
)

const (
  KB = 1000
  MB = KB * KB
  GB = MB * KB
  TB = GB * KB
  PB = TB * KB
  EB = PB * KB
  ZB = EB * KB
  YB = ZB * KB
)

func main() {
  fmt.Println(KB)
  fmt.Println(MB)
  fmt.Println(GB)
  fmt.Println(TB)
  fmt.Println(PB)
  fmt.Println(EB)
  // fmt.Println(ZB) // constant 1000000000000000000000 overflows int
  // fmt.Println(YB)
}
