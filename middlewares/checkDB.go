package middlewares

import (
	"net/http"

	"github.com/jesusafd/SRU/connections"
)

// CheckDB Se envarga de verificar que la conexion con la bd este activa
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !connections.CheckConnection() {
			http.Error(w, "Error, conexion con la bd perdida", http.StatusConflict)
			return
		}
		next.ServeHTTP(w, r)
	}
}
