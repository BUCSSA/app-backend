package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set!")
	}

	route := mux.NewRouter()
	route.HandleFunc("/", HomeHandler)
	n := negroni.Classic()
	n.UseHandler(route)

	http.ListenAndServe(":"+port, n)
}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello Go Server!")
}
