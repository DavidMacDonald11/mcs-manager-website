package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
    port := "8000"
    fmt.Printf("Starting server on port %s\n", port)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, "Hello")
    })

    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
