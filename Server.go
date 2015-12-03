package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("H")
	http.HandleFunc("/", handler)
	http.ListenAndServe("127.0.0.1:3333", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello", r.URL.Path)
}
