package middlewares

import (
	"net/http"

	"github.com/jesusafd/SRU/connections"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !connections.CheckConnection() {
			http.Error(w, "Error, conexion con la bd perdida", http.StatusConflict)
			return
		}
		next(w, r)
	}
}
