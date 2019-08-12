package main

import (
  "fmt"
  "strings"
  "bufio"
)

func main() {
  var res string
  s := "Alpha $one and beta $two"
  fmt.Println(s)
  res = expand(s, func(s string) string {
    return strings.ToUpper(s)
  })
  fmt.Println(res)
  res = expand(s, func(s string) string {
    return strings.ReplaceAll(s, "o", "0")
  })
  fmt.Println(res)
}

// Replaces "$foo" with f("foo")
// Example:
// expand("Alpha $abc and beta $xyz", func(s) {
//   return strings.Title(s)
// })
func expand(s string, f func(string) string) string {
  var buf []string
  scanner := bufio.NewScanner(strings.NewReader(s))
  scanner.Split(bufio.ScanWords)
  for scanner.Scan() {
    var word string
    if strings.HasPrefix(scanner.Text(), "$") {
      word = f(strings.TrimPrefix(scanner.Text(), "$"))
    } else {
      word = scanner.Text()
    }
    buf = append(buf, word)
  }
  return strings.Join(buf, " ")
}
