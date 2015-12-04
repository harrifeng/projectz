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

func CommonMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("All----------->")
	next(rw, r)
	fmt.Println("All----------->")
}

func AdminMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("Admin----------->")
	next(rw, r)
	fmt.Println("Admin----------->")
}

func main() {
	mux := bone.New()
	mux.Get("/some/page/:id", http.HandlerFunc(handler))
	// Special treat, extra middleware for admin
	mux.Get("/admin", negroni.New(negroni.HandlerFunc(AdminMiddleware)))

	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(CommonMiddleware))
	n.UseHandler(mux)
	n.Run(":3003")
}
