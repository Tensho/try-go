package main

// I don't know how to share packages yet, so please use the next command to run
// $ go run 1.12.go lissajous.go
import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
