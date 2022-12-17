package connections

import (
	"context"
	"log"
	"time"
)

// Check es la funcion encardada de buscar en la bd,
// si el usuario esta o no verificado
func Check(email string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	status := false
	query := "SELECT * FROM verified WHERE email=?"
	// En la tabla verified de la base de datos se encuentran
	// las cuentas de verificadas
	rows, err := DBConn.QueryContext(
		ctx,
		query,
		email,
	)
	if err != nil {
		log.Println("Error al realizar la busqueda en la bd : " + err.Error())
		return status, err
	}
	// En caso de que la consulta devolviera una fila para leer,
	// quiere decir que la cuneta se encuentra verificada
	if rows.Next() {
		status = true
	}
	return status, nil
}
