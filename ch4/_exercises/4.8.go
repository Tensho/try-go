// $ cat a.txt | go run 4.8.go

package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "unicode"
  "unicode/utf8"
  "sort"
)

func main() {
  counts := make(map[rune]int)       // counts of Unicode characters
  categories := make(map[string]int) // counts of Unicode characters per category
  var utflen [utf8.UTFMax + 1]int    // count of lengths of UTF-8 encodings
  invalid := 0                       // count of invalid UTF-8 characters

  in := bufio.NewReader(os.Stdin)
  for {
    r, n, err := in.ReadRune() // returns rune, nbytes, error
    if err == io.EOF {
      break
    }
    if err != nil {
      fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
      os.Exit(1)
    }
    if r == unicode.ReplacementChar && n == 1 {
      invalid++
      continue
    }
    counts[r]++
    utflen[n]++

    if unicode.IsDigit(r) {
      categories["digit"]++
    }
    if unicode.IsGraphic(r) {
      categories["graphic"]++
    }
    if unicode.IsLetter(r) {
      categories["letter"]++
    }
    if unicode.IsLower(r) {
      categories["lower"]++
    }
    if unicode.IsMark(r) {
      categories["mark"]++
    }
    if unicode.IsNumber(r) {
      categories["number"]++
    }
    if unicode.IsPrint(r) {
      categories["print"]++
    }
    if unicode.IsSpace(r) {
      categories["space"]++
    }
    if unicode.IsSymbol(r) {
      categories["symbol"]++
    }
    if unicode.IsTitle(r) {
      categories["title"]++
    }
    if unicode.IsUpper(r) {
      categories["upper"]++
    }
  }

  fmt.Printf("rune\tcount\n")
  for c, n := range counts {
    fmt.Printf("%q\t%d\n", c, n)
  }
  fmt.Print("\nlen\tcount\n")
  for i, n := range utflen {
    if i > 0 {
      fmt.Printf("%d\t%d\n", i, n)
    }
  }

  // sort map keys for the consistent output
  var names []string
  for name := range categories { // take only keys here and omit values
      names = append(names, name)
  }
  sort.Strings(names)

  fmt.Print("\ncategory\tcount\n")
  for _, name := range names {
    fmt.Printf("%s\t\t%d\n", name, categories[name])
  }

  if invalid > 0 {
    fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
  }
}
