package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handlerRoot)
	http.HandleFunc("/counter", handlerCount)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Patch = %q\n", r.URL.Path)
}

func handlerCount(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
