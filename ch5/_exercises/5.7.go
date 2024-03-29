package main

import (
  "fmt"
  "net/http"
  "os"
  "strings"

  "golang.org/x/net/html"
)

func main() {
  for _, url := range os.Args[1:] {
    outline(url)
  }
}

func outline(url string) error {
  resp, err := http.Get(url)
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  doc, err := html.Parse(resp.Body)
  if err != nil {
    return err
  }

  forEachNode(doc, startElement, endElement)

  return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
  if pre != nil {
    pre(n)
  }

  for c := n.FirstChild; c != nil; c = c.NextSibling {
    forEachNode(c, pre, post)
  }

  if post != nil {
    post(n)
  }
}

var depth int

func startElement(n *html.Node) {
  var s string

  if n.Type == html.ElementNode {
    if len(n.Attr) > 0 {
      for _, a := range n.Attr {
        s += fmt.Sprintf(" %s=\"%s\"", a.Key, a.Val)
      }
    }

    if n.FirstChild == nil {
      fmt.Printf("%*s<%s%s />\n", depth*2, "", n.Data, s)
    } else {
      fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, s)
    }

    depth++
  }

  if n.Type == html.CommentNode {
    fmt.Printf("%*s<!--%s-->\n", depth*2, "", n.Data)
  }

  if n.Type == html.TextNode && strings.TrimSpace(n.Data) != "" {
    fmt.Printf("%*s%s\n", depth*2, "", n.Data)
  }
}

func endElement(n *html.Node) {
  if n.Type == html.ElementNode {
    depth--

    if n.FirstChild != nil {
      fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
    }
  }
}
