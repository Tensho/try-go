package main

import (
  "bytes"
  "fmt"
)

type IntSet struct {
  words []uint64
}

func main() {
  var a, b IntSet
  a.AddAll(1, 65, 129, 1)
  b.AddAll(1, 130, 65, 65)

  a.UnionWith(&b)
  fmt.Printf("Union: %s\n", a.String())

  var c, d IntSet
  c.AddAll(1, 65, 129, 131, 250)
  d.AddAll(1, 130, 65)

  c.IntersectWith(&d)
  fmt.Printf("Intersect 1: %s\n", c.String())

  var e, f IntSet
  e.AddAll(1, 65, 129)
  f.AddAll(1, 130, 65, 244)

  e.IntersectWith(&f)
  fmt.Printf("Intersect 2: %s\n", e.String())

  var g, h IntSet
  g.AddAll(1, 65, 129, 131, 250)
  h.AddAll(1, 130, 65)

  g.DifferenceWith(&h)
  fmt.Printf("Difference 1: %s\n", g.String())

  var i, j IntSet
  i.AddAll(1, 65, 129)
  j.AddAll(1, 130, 65, 244, 250)

  i.DifferenceWith(&j)
  fmt.Printf("Difference 2: %s\n", i.String())

  var k, l IntSet
  k.AddAll(1, 65, 129, 131, 249)
  l.AddAll(1, 130, 65)

  k.SymmetricDifference(&l)
  fmt.Printf("Symmetric Difference 1: %s\n", k.String())

  var m, n IntSet
  m.AddAll(1, 65, 129)
  n.AddAll(1, 130, 65, 244, 250)

  m.SymmetricDifference(&n)
  fmt.Printf("Symmetric Difference 2: %s\n", m.String())
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

// IntersectWith sets s to the intersect of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
  for i, tword := range t.words {
    if i < len(s.words) {
      s.words[i] &= tword
    } else {
      // Nullify the rest of the argument's vector in case if it's longer
      t.words[i] = 0
    }
  }
  // Nullify the rest of the receiver's vector in case if it's longer
  for i := len(t.words); i < len(s.words); i++ {
    s.words[i] = 0
  }
}

// DifferenceWith sets s to the difference of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) {
  for i, tword := range t.words {
    if i < len(s.words) {
      s.words[i] &^= tword
    }
  }
}

// SymmetricDifference sets s to the symmetric difference of s and t.
func (s *IntSet) SymmetricDifference(t *IntSet) {
  for i, tword := range t.words {
    if i < len(s.words) {
      s.words[i] ^= tword
    } else {
      s.words = append(s.words, tword)
    }
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
