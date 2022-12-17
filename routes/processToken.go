package routes

import (
	"errors"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jesusafd/SRU/connections"
	"github.com/jesusafd/SRU/models"
)

// Email almacena el email extraido del token
var Email string

// ProcessToken la funcion encargada de extraer los datos del token
func ProcessToken(token string) (*models.Claim, error) {
	clave := []byte("Jesus12345")

	// Nota es importante de clarar el objeto claim de la forma
	// claim := &models.Claim{}
	// Ya que al querer declararlo de otra forma como puede ser:
	// var claim *models.Claim
	// El token no podra mapearse en este objeto
	claim := &models.Claim{}

	// Quitamos los espacios en blanco
	token = strings.TrimSpace(token)

	_, err := jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (interface{}, error) { return clave, nil })
	// En caso de no existir error el token es valido
	if err == nil {
		// Verificamos que el email del token se encuentr en la bd
		_, err = connections.Read(claim.Email)
		if err == nil {
			Email = claim.Email
			return claim, nil
		}
	} else {
		log.Println("Error al mapear el token : " + err.Error())
		return claim, err
	}
	// en caso de que el token no sea invalido
	err = errors.New("token valido usuario no existe")
	return claim, err
}
