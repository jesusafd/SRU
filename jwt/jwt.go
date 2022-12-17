package jwt

import (
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/jesusafd/SRU/models"
)

// GenerateJWT Genera un token jwt
func GenerateJWT(u models.User) (string, error) {
	// status, err := connections.Check(u.Email)
	// if err != nil {
	// 	log.Println("Error al buscar cuenta en cuentas verificadas")
	// 	return "", err
	// }
	// El payload solo contendra el email y el estado de la cuenta
	pyload := jwt.MapClaims{
		"email": u.Email,
		// "status": status,
	}
	// Creamos el token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, pyload)
	// Firmamos el token y obtenemos el token final
	tokenStr, err := token.SignedString(Clave)
	if err != nil {
		log.Println("Error al firmar el token " + err.Error())
	}
	return tokenStr, err
}
