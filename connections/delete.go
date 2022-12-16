package connections

import (
	"context"
	"errors"
	"log"
	"time"
)

// Delete es la funcion encargada de eliminar el registro de la bd, en caso
// de devolver un erro el regstro no se elmino o no existe
func Delete(email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := "DELETE FROM user WHERE email=?"
	result, err := DBConn.ExecContext(
		ctx,
		query,
		email,
	)
	if err != nil {
		log.Println("Error al eliminar registro en la base de datos : " + err.Error())
		return err
	}
	// Leemos el registro, si es que se encontro
	var rowsAffected int64
	// verificamos si se elimino algun registro
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Println("Error al eliminar el registro : " + err.Error())
		return err
	}
	if rowsAffected != 0 {
		err = errors.New("usuario no existe")
	}
	return err

}
