package main

import "fmt"

// Return 42 without "return" statements
func tmp() (res int) {
  defer func() {
    if r := recover(); r != nil {
      res = 42
    }
  }()

  panic("A")
}

func main() {
  fmt.Println(tmp())
}
