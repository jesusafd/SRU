package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jesusafd/SRU/connections"
	"github.com/jesusafd/SRU/models"
)

func InsertUser(w http.ResponseWriter, r *http.Request) {
	var u models.User

	// Decodificamos el contenido de la peticion en el objeto de usario
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Println("Error al decodificar el mensaje : " + err.Error())
		http.Error(w, "Error del servidor : "+err.Error(), http.StatusConflict)
		return
	}
	if u.Email == "" || u.Password == "" {
		log.Println("Error, Todos los datos son requeridos : " + err.Error())
		http.Error(w, "Error, Todos los datos son requeridos : "+err.Error(), http.StatusConflict)
		return
	}
	// Insertamos el registro en la bd
	err = connections.Create(u)
	// Verificamos si la inserccion fue exitosa
	if err != nil {
		log.Println("Error al insertar en la base de datos : " + err.Error())
		http.Error(w, "Error al insertar en la base de datos : "+err.Error(), http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
