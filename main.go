package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/go-zoo/bone"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %v\n", bone.GetValue(r, "id"))
}

func MyMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("Hello----------->")
	next(rw, r)
	fmt.Println("World===========>")
}

func main() {
	mux := bone.New()
	mux.Get("/some/page/:id", http.HandlerFunc(handler))
	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(MyMiddleware))
	n.UseHandler(mux)
	n.Run(":3003")
}
