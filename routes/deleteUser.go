package routes

import (
	"log"
	"net/http"

	"github.com/jesusafd/SRU/connections"
)

// DeleteUser es el endpoint para eliminar un perfil de la bd
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Hacemos uso del email, extraido del token jwt
	// Un usuario normal solo puede eliminar su propia cuenta
	err := connections.Delete(Email)
	if err != nil {
		log.Println("Error al eliminar el registro en la bd : " + err.Error())
		http.Error(w, "Error al eliminar el registro en la bd : "+err.Error(), http.StatusConflict)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}
