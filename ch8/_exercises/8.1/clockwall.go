// go run clockwall.go NewYork=localhost:8000 Tokyo=localhost:8010 London=localhost:8020

package main

import (
  "bufio"
  "log"
  "net"
  "os"
  "fmt"
  "strings"
)

func main() {
  done := make(chan int)

  for _, arg := range os.Args[1:] {
    s := strings.Split(arg, "=")
    city, addr := s[0], s[1]
    fmt.Println(city, addr)
    conn, err := net.Dial("tcp", addr)
    if err != nil {
      log.Fatal(err)
    }
    defer conn.Close()
    go func() {
      buf := bufio.NewScanner(conn)
      for buf.Scan() {
        fmt.Println(city, buf.Text())
      }
      done <- 1
    }()
  }

  <-done
}


