package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/", HomeHandler)
	n := negroni.Classic()
	n.UseHandler(route)

	http.ListenAndServe(":3000", n)
}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello Go Server!")
}
