package main

import (
  "io"
  "golang.org/x/net/html"
)

type stringReader struct {
  source string
}

func main() {
  s := "<html><body><div>ABC</div></body></html>"
  _, _ = html.Parse(NewReader(s))
}

func NewReader(s string) io.Reader {
  return &stringReader{source: s}
}

func (r *stringReader) Read(p []byte) (n int, err error) {
  n = copy(p, r.source)
  r.source = r.source[n:]
  if len(r.source) == 0 {
    err = io.EOF
  }
  return
}
