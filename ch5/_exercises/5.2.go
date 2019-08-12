package main

import (
  "fmt"
  "os"

  "golang.org/x/net/html"
)

func main() {
  doc, err := html.Parse(os.Stdin)
  if err != nil {
    fmt.Fprintf(os.Stderr, "traverse: %v\n", err)
    os.Exit(1)
  }
  html.Render(os.Stdout, doc)

  counts := make(map[string]int)

  traverse(counts, doc)

  fmt.Println("\n\n========================================\n")

  for t, n := range counts {
    fmt.Printf("%s\t%d\n", t, n)
  }
}

func traverse(counts map[string]int, n *html.Node) {
  if n.Data == "p" || n.Data == "div" || n.Data == "span" {
    counts[n.Data]++
  }

  for c := n.FirstChild; c != nil; c = c.NextSibling {
    traverse(counts, c)
  }
}
