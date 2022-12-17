package middlewares

import (
	"net/http"

	"github.com/jesusafd/SRU/routes"
)

// TokenValidation valida el jwt que vine cn la peticion
func TokenValidation(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := routes.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error al procesar token jwt : "+err.Error(), http.StatusConflict)
			return
		}
		next.ServeHTTP(w, r)
	}
}
