package main

import (
	"fmt"
	"net/http"
	"github.com/codegangsta/negroni"
	"github.com/go-zoo/bone"
	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %v\n", bone.GetValue(r, "id"))
}

func GHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Hello %v, From gorilla\n", vars["id"])
}

func main() {
	bmux := bone.New()
	bmux.Get("/some/page/:id", http.HandlerFunc(handler))

	gmux := mux.NewRouter()
	gmux.HandleFunc("/other/page/{id}", GHandler).Methods("GET")

	n := negroni.Classic()
 	n.UseHandler(bmux)
	n.UseHandler(gmux)
	n.Run(":3004")
}

//Use two mux will get following result with
//=> http://localhost:3004/other/page/2323
//404 page not found
//Hello 2323, From gorilla