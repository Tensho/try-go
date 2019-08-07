// $ cat a.txt | go run 4.9.go

package main

import (
  "fmt"
  "os"
  "bufio"
)

func main() {
  distribution := make(map[string]int)

  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanWords)
  for scanner.Scan() {
    word := scanner.Text()
    distribution[word]++
  }

  for word, n := range distribution {
    fmt.Printf("%-30s %d\n", word, n)
  }
}
