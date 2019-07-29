package main

import (
  "fmt"
  "os"
  "bytes"
)

//  $ go run 3.10.go 1 12 123 1234 1234567890
//  1
//  12
//  123
//  1,234
//  1,234,567,890
func main() {
  for i := 1; i < len(os.Args); i++ {
    fmt.Printf("  %s\n", comma(os.Args[i]))
  }
}

func comma(s string) string {
  var buf bytes.Buffer
  cnt := len(s) % 3
  start := true // skip comma for the module 3 number
  for _, r := range s {
    if cnt % 3 == 0 && start == false {
      buf.WriteByte(',')
      cnt = 3
    }
    buf.WriteByte(byte(r))
    start = false
    cnt--
  }
  return buf.String()
}
