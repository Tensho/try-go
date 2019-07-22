package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
  s, sep := "", ""
  for _, arg := range os.Args[1:] {
      s += sep + arg
      sep = " "
  }
  fmt.Println(s)
  fmt.Println("+:", time.Since(start))

	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println("strings.Join:", time.Since(start))
}
