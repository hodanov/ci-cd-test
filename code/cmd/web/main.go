package main

import (
	"fmt"
	"net/http"
)

func hoge(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hogehoge")
}

func top(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func main() {
	http.HandleFunc("/", top)
	http.HandleFunc("/hoge", hoge)
	http.ListenAndServe(":8080", nil)
}
