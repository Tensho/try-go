// TZ=US/Eastern go run clock.go -port 8000
// TZ=Asia/Tokyo go run clock.go -port 8010
// TZ=Europe/London go run clock.go -port 8020

package main

import (
  "os"
  "io"
  "log"
  "net"
  "time"
  "flag"
)

var (
  port = flag.String("port", "8000", "port number")
)

func handleConn(c net.Conn) {
  defer c.Close()

  for {
    var now time.Time

    if tz, ok := os.LookupEnv("TZ"); ok {
      location, _ := time.LoadLocation(tz)
      now = time.Now().In(location)
    } else {
      now = time.Now()
    }

    _, err := io.WriteString(c, now.Format("15:04:05\n"))
    if err != nil {
      return // client disconnected
    }
    time.Sleep(1 * time.Second)
  }
}

func main() {
  flag.Parse()
  listener, err := net.Listen("tcp", "localhost:" + *port)
  if err != nil {
    log.Fatal(err)
  }
  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Print(err) // connection aborted
      continue
    }
    go handleConn(conn) // handle connections concurrently
  }
}
