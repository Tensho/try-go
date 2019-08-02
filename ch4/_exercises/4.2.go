package main

import (
  "flag"
  "fmt"
  "bufio"
  "os"
  "strings"
  "crypto/sha256"
  "crypto/sha512"
)

var (
  sha384_flag = flag.Bool("sha384", false, "SHA384")
  sha512_flag = flag.Bool("sha512", false, "SHA512")
)

func main() {
  flag.Parse()

  r := bufio.NewReader(os.Stdin)
  s, err := r.ReadString('\n')
  if err != nil {
    fmt.Fprintf(os.Stderr, "%v\n", err)
    os.Exit(1)
  }
  s = strings.Trim(s, "\n")

  fmt.Printf("SHA256: %x\n", sha256.Sum256([]byte(s)))

  if *sha384_flag {
    fmt.Printf("SHA384: %x\n", sha512.Sum384([]byte(s)))
  }

  if *sha512_flag {
    fmt.Printf("SHA512: %x\n", sha512.Sum512([]byte(s)))
  }
}
