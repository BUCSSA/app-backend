package app_backend

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello Go Server!")
}

func PromorQueryHandler(w http.ResponseWriter, req *http.Request) {

}
