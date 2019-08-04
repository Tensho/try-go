package main

import "fmt"

func main() {
  s := []string{"A", "A", "B", "C", "C", "C", "D"}
  fmt.Println(s)
  res := dedup(s)
  fmt.Println(res)
}

// In-place function to eliminate adjacent duplicates in a string slice
func dedup(s []string) []string {
  out := s[:0]
  for i := 0; i < len(s) - 1; i++ {
    if s[i] != s[i+1] {
      out = append(out, s[i])
    }
  }
  return out
}
