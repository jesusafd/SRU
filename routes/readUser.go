package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jesusafd/SRU/connections"
)

func ReadUser(w http.ResponseWriter, r *http.Request) {
	// Extraemos el email de la url
	email := r.URL.Query().Get("email")

	u, err := connections.Read(email)
	if email == "" {
		log.Println("Error, el email es requerido : " + err.Error())
		http.Error(w, "Error, el email es requerido : "+err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		log.Println("Error al buscar el registro en la bd : " + err.Error())
		http.Error(w, "Error al buscar el registro en la bd : "+err.Error(), http.StatusConflict)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(u)
}
