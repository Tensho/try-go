package main

import (
  "fmt"
  "net/http"
  "os"

  "golang.org/x/net/html"
)

func main() {
  for _, url := range os.Args[1:] {
    links, err := findLinks(url)
    if err != nil {
      fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
      continue
    }
    for _, link := range links {
      fmt.Println(link)
    }
  }
}

func visit(links []string, n *html.Node) []string {
  if n.Type == html.ElementNode {
    if n.Data == "a" || n.Data == "link" {
      for _, a := range n.Attr {
        if a.Key == "href" && a.Val != "" {
          links = append(links, a.Val)
        }
      }
    }
    if n.Data == "script" || n.Data == "img" {
      for _, a := range n.Attr {
        if a.Key == "src" && a.Val != "" {
          links = append(links, a.Val)
        }
      }
    }
  }
  for c := n.FirstChild; c != nil; c = c.NextSibling {
    links = visit(links, c)
  }
  return links
}

func findLinks(url string) ([]string, error) {
  resp, err := http.Get(url)
  if err != nil {
    return nil, err
  }
  if resp.StatusCode != http.StatusOK {
    resp.Body.Close()
    return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
  }
  doc, err := html.Parse(resp.Body)
  resp.Body.Close()
  if err != nil {
    return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
  }
  return visit(nil, doc), nil
}
