package main

import (
  "fmt"
  "os"
  "strings"
  "net/http"
  "golang.org/x/net/html"
)

func main() {
  fmt.Println(CountWordsAndImages(os.Args[1]))
}

func CountWordsAndImages(url string) (words, images int, err error) {
  resp, err := http.Get(url)
  if err != nil {
      return
  }

  doc, err := html.Parse(resp.Body)
  resp.Body.Close()
  if err != nil {
      err = fmt.Errorf("parsing HTML: %s", err)
      return
  }
  words, images = countWordsAndImages(doc)
  return
}

func countWordsAndImages(n *html.Node) (words, images int) {
  if n.Type == html.ElementNode && n.Data == "img" {
    images++
  }

  if n.Type == html.TextNode {
    words += len(strings.Fields(n.Data))
  }

  for c := n.FirstChild; c != nil; c = c.NextSibling {
    w, i := countWordsAndImages(c)
    words += w
    images += i
  }

  return
}
