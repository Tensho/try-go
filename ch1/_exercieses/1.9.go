package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	schema = "http"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, schema) {
			url = schema + "://" + url
		}
		fmt.Printf("Fetching from: %s...\n", url)

		resp, err := http.Get(url);
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n\n", resp.Status)
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch: copying %s: %v\n", url, err)
			os.Exit(1)
		}
		resp.Body.Close()
	}
}
