package main

import (
  "fmt"
  "unicode"
  "unicode/utf8"
)

func main() {
  s := "\tHellΩ \u00A0 ∀lpha  B\tCℷrca\u0085\n"
  fmt.Println(s)
  res := squashSpaces([]byte(s))
  fmt.Println(string(res))
}

// Squash UTF-8 spaces in byte slice
func squashSpaces(s []byte) []byte {
  out := s[:0]
  for i := 0; i < len(s); {
    current_rune, size := utf8.DecodeRune(s[i:])
    next_rune, _ := utf8.DecodeRune(s[i+size:])

    switch {
    case unicode.IsSpace(current_rune) && unicode.IsSpace(next_rune):
      // skip asjucent spaces
    case unicode.IsSpace(current_rune):
      out = append(out, 32) // normalize to ASCII space
    default:
      out = append(out, s[i:i+size]...)
    }

    i += size
  }
  return out
}
