package main

// I don't know how to share packages yet, so please use the next command to run
// $ go run 1.12.go lissajous.go > out.gif
import (
  "os"
)

func main() {
  lissajous(os.Stdout)
}
