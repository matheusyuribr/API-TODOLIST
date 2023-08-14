package handlers

import (
	"api-postgresql/models"
	"encoding/json"
	"log"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {

	todos, err := models.GetAll()

	if err != nil {
		log.Printf("Error ao obter reegistros: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)

}