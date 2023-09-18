package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodrigoreeis/api-go-rest/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/heath", controllers.Heath)
	r.HandleFunc("/api/v1/personalities/", controllers.GetPersonalities).Methods("GET")
	r.HandleFunc("/api/v1/personalities/{id}", controllers.GetPersonalityById).Methods("GET")
	fmt.Println("server localhost port:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
