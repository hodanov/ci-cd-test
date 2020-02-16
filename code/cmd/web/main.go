package main

import (
	"fmt"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintf(w, "Hello")
}

func main() {
        http.HandleFunc("/", top)
        http.ListenAndServe(":8080", nil)
}
