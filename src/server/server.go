package app_backend

import (
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
)

var (
	port   = os.Getenv("PORT")
	db_uri = os.Getenv("MONGODB_URI")
)

func init() {

	if port == "" {
		log.Fatal("$PORT must be set!")
	}

	if db_uri == "" {
		log.Fatal("DB URI must be set!")
	}

}

func main() {

	route := NewRouter()

	n := negroni.Classic()
	n.UseHandler(route)

	log.Fatal(http.ListenAndServe(":"+port, n))
}
