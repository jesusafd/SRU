package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jesusafd/SRU/connections"
	"github.com/jesusafd/SRU/jwt"
	"github.com/jesusafd/SRU/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// u es la varible encargada de extraer los datos de la request
	var u models.User
	json.NewDecoder(r.Body).Decode(&u)
	if len(u.Email) == 0 {

		http.Error(w, "Email no enviado", 400)
		return
	}
	if len(u.Password) == 0 {
		http.Error(w, "Contraseña no enviada", 400)
		return
	}
	// usr es el objeto usado para extaer los datos de la bd
	usr, err := connections.Read(u.Email)
	// En aso de haber un error lo mas probable es que
	// el usuario no exista
	if err != nil {
		log.Println("Error : " + err.Error())
		http.Error(w, "Error : "+err.Error(), http.StatusConflict)
		return
	}
	// comprobamos las credenciales
	err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(u.Password))
	if err != nil {
		log.Println("Error en el login, credenciales invalidas" + err.Error())
		http.Error(w, "Error en el login, credenciales invalidas"+err.Error(), http.StatusConflict)
		return
	}
	// Cerdenciales validas
	var token string
	token, err = jwt.GenerateJWT(usr)
	if err != nil {
		log.Println("Error al iniciar sesion : " + err.Error())
		http.Error(w, "Error al iniciar sesion : "+err.Error(), http.StatusConflict)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(token)

}
