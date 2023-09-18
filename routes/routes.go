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
	r.HandleFunc("/api/v1/personalities/", controllers.GetPersonalities).Methods("GET")
	r.HandleFunc("/api/v1/personalities/{id}", controllers.GetPersonalityById).Methods("GET")
	r.HandleFunc("/api/v1/personalities/", controllers.CreatePersonality).Methods("POST")
	r.HandleFunc("/api/v1/personalities/{id}", controllers.DeletePersonality).Methods("DELETE")
	r.HandleFunc("/api/v1/personalities/{id}", controllers.UpdatePersonality).Methods("PUT")
	fmt.Println("server localhost port:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
