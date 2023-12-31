package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rodrigoreeis/api-go-rest/controllers"
	"github.com/rodrigoreeis/api-go-rest/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentType)
	r.HandleFunc("/api/v1/personalities/", controllers.GetPersonalities).Methods("GET")
	r.HandleFunc("/api/v1/personalities/{id}", controllers.GetPersonalityById).Methods("GET")
	r.HandleFunc("/api/v1/personalities/", controllers.CreatePersonality).Methods("POST")
	r.HandleFunc("/api/v1/personalities/{id}", controllers.DeletePersonality).Methods("DELETE")
	r.HandleFunc("/api/v1/personalities/{id}", controllers.UpdatePersonality).Methods("PUT")
	fmt.Println("server localhost port:8000")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
