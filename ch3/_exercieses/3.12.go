package main

import (
  "fmt"
  "os"
  "strings"
  "sort"
  "reflect"
)

// Anagrams
func main() {
  word_1 := os.Args[1]
  word_2 := os.Args[2]
  fmt.Printf("%s << Anagram >> %s ?\n", word_1, word_2)
  letters_1 := strings.Split(word_1, "")
  letters_2 := strings.Split(word_2, "")
  sort.Strings(letters_1)
  sort.Strings(letters_2)
  if reflect.DeepEqual(letters_1, letters_2) {
    fmt.Println("YES")
  } else {
    fmt.Println("NO")
  }
}
