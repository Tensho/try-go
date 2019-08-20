package main

import (
  "bytes"
  "fmt"
)

type IntSet struct {
  words []uint64
}

func main() {
  var x IntSet
  x.AddAll(1, 144, 9)
  fmt.Println(x.String())
}

// Add adds the non-negative values x to the set.
func (s *IntSet) AddAll(values ...int) {
  for _, v := range values {
    word, bit := v/64, uint(v%64)
    for word >= len(s.words) {
      s.words = append(s.words, 0)
    }
    s.words[word] |= 1 << bit
  }
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
  var buf bytes.Buffer
  buf.WriteByte('{')
  for i, word := range s.words {
    if word == 0 {
      continue
    }
    for j := 0; j < 64; j++ {
      if word&(1<<uint(j)) != 0 {
        if buf.Len() > len("{") {
          buf.WriteByte(' ')
        }
        fmt.Fprintf(&buf, "%d", 64*i+j)
      }
    }
  }
  buf.WriteByte('}')
  return buf.String()
}
