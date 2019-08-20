package main

import (
  "io"
  "golang.org/x/net/html"
)

type StringReader struct {
  source string
}

func main() {
  s := "<html><body><div>ABC</div></body></html>"
  _, _ = html.Parse(NewReader(s))
}

func NewReader(s string) io.Reader {
  return &StringReader{source: s}
}

func (r *StringReader) Read(p []byte) (n int, err error) {
  n = copy(p, r.source)
  r.source = r.source[n:]
  if len(r.source) == 0 {
    err = io.EOF
  }
  return
}
