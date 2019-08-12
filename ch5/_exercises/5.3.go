package main

import (
  "fmt"
  "os"
  "strings"

  "golang.org/x/net/html"
)

func main() {
  doc, err := html.Parse(os.Stdin)
  if err != nil {
    fmt.Fprintf(os.Stderr, "traverse: %v\n", err)
    os.Exit(1)
  }
  html.Render(os.Stdout, doc)
  fmt.Println("\n\n========================================\n")

  traverse(doc)
}

func traverse(n *html.Node) {
  if n.Data == "head" || n.Data == "script" || n.Data == "style" {
    return
  }

  // fmt.Printf("\s", n.Data)
  if n.Type == html.TextNode && strings.TrimSpace(n.Data) != "" {
   fmt.Println(n.Data)
  }

  for c := n.FirstChild; c != nil; c = c.NextSibling {
    traverse(c)
  }
}
