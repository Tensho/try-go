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
  x.Add(1)
  x.Add(144)
  x.Add(9)
  fmt.Println(x.String())
  fmt.Println(x.Len())
  x.Remove(9)
  fmt.Println(x.String())
  x.Clear()
  fmt.Println(x.String())
  x.Add(5)
  x.Add(40)
  fmt.Println(x.String())

  y := x.Copy()
  fmt.Println(y.String())
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
  word, bit := x/64, uint(x%64)
  return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
  word, bit := x/64, uint(x%64)
  for word >= len(s.words) {
    s.words = append(s.words, 0)
  }
  s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
  for i, tword := range t.words {
    if i < len(s.words) {
      s.words[i] |= tword
    } else {
      s.words = append(s.words, tword)
    }
  }
}

// Len returns the number of elements.
func (s *IntSet) Len() (count int) {
  for _, word := range s.words {
    for i := uint(0); i < 64; i++ {
      if word & (1 << i) != 0 {
        count++
      }
    }
  }
  return
}

// Remove removes x from the set.
func (s *IntSet) Remove(x int) {
  word, bit := x/64, uint(x%64)
  s.words[word] &= ^(1<<bit)
}

// Clear removes all elements from the set.
func (s *IntSet) Clear() {
  for i, _ := range s.words {
    s.words[i] = 0
  }
}

// Copy returns a copy of the set.
func (s *IntSet) Copy() (sc *IntSet) {
  sc = &IntSet{}
  sc.words = make([]uint64, len(s.words))
  copy(sc.words, s.words)
  return
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
