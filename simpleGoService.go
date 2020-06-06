package main

// We import any external packages required
import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// An object to represent the JSON we will return
type pingResponse struct {
	Code   string `json:"code"`
	Status string `json:"status"`
}

// sayHello implements the GET method of this REST API
func sayHello(w http.ResponseWriter, r *http.Request) {

	// Create a response object
	responseData := &pingResponse{}
	responseData.Code = "ok"
	responseData.Status = "Hello, world!"

	// Return HTTP200...
	w.WriteHeader(http.StatusOK)

	// ...and attempt to serialize and send the response
	if err := json.NewEncoder(w).Encode(responseData); err != nil {
		panic(err)
	}
}

// Main method, this runs when our app executes
func main() {

	router := *mux.NewRouter().StrictSlash(true)
	router.Methods("GET").
		Path("/hello").
		HandlerFunc(sayHello)

	log.Fatal(http.ListenAndServe("0.0.0.0:80", &router))

}
