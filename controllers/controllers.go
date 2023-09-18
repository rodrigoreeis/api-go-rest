package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodrigoreeis/api-go-rest/database"
	"github.com/rodrigoreeis/api-go-rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

func Heath(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func GetPersonalities(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalities
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func GetPersonalityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var personality models.Personalities
	database.DB.First(&personality, vars["id"])
	json.NewEncoder(w).Encode(personality)
}
