package main

import (
	"fmt"
	"net/http"
	"github.com/codegangsta/negroni"
	"github.com/go-zoo/bone"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %v\n", bone.GetValue(r, "id"))
}

func main() {
	mux := bone.New()
	mux.Get("/some/page/:id", http.HandlerFunc(handler))
	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3003")
}

