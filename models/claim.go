package models

import jwt "github.com/dgrijalva/jwt-go"

// Claim es la estrcutura usada para decodifcar el jwt
type Claim struct {
	// El email del token, por eso colocamos el formato json
	Email string `json:"email"`
	// Status bool   `json:"status"`
	// StandardClaims es una estructura que guarda los datos
	// que se agragan por defecto al jwt
	jwt.StandardClaims
}
