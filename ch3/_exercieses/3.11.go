package main

import (
  "fmt"
  "os"
  "strings"
)

//  $ go run 3.11.go 1.2 12.3 123.456 1234.5678 12345.67890
//  1.0
//  12.4
//  123.501
//  1,234.567,8
//  12,345.678,90
func main() {
  for i := 1; i < len(os.Args); i++ {
    fmt.Printf("  %s\n", comma(os.Args[i]))
  }
}

func comma(s string) string {
  parts := strings.Split(s, ".")
  i, f := parts[0], parts[1]

  return integer_comma(i) + "." + fraction_comma(f)
}

func integer_comma(s string) string {
  n := len(s)
  if n <= 3 {
    return s
  }
  return integer_comma(s[:n-3]) + "," + s[n-3:]
}

func fraction_comma(s string) string {
  n := len(s)
  if n <= 3 {
    return s
  }
  return s[:3] + "," + fraction_comma(s[3:])
}
