package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jesusafd/SRU/connections"
	"github.com/jesusafd/SRU/models"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var u models.User

	// Decodificamos el contenido de la peticion en el objeto de usario
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Println("Error al decodificar el mensaje : " + err.Error())
		http.Error(w, "Error del servidor : "+err.Error(), http.StatusConflict)
		return
	}
	// Usamos el email extraido del token
	u.Email = Email
	// verificamos que exista un registro con el mismo correo
	_, err = connections.Read(u.Email)
	if err != nil {
		log.Println("Error al insertar registro, usuario ya existe")
		http.Error(w, "Error al insertar registro, usuario ya existe", http.StatusConflict)
		return
	}
	// Actulizamos el registro en la bd
	err = connections.Update(u)
	// Verificamos si la inserccion fue exitosa
	if err != nil {
		log.Println("Error al Actualizar en la base de datos : " + err.Error())
		http.Error(w, "Error al Actualizar en la base de datos : "+err.Error(), http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
