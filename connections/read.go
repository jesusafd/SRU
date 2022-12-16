package connections

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/jesusafd/SRU/models"
)

// Read es la funncion encargada de leer un registro en la bd
// tambien puede ser usado para saber si un usario existe
// en caso de que el usuario no exista la funcion devolvera un error
func Read(email string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var u models.User

	query := "SELECT * FROM user WHERE email=?"
	rows, err := DBConn.QueryContext(
		ctx,
		query,
		email,
	)
	if err != nil {
		log.Println("Error al buscar registro en la base de datos : " + err.Error())
		return u, err
	}
	// Leemos el registro, si es que se encontro
	if rows.Next() {
		err = rows.Scan(&u.Email, &u.Password, &u.Name, &u.LastName, &u.YearBirth, &u.MonthBirth, &u.DayBirth)
		if err != nil {
			log.Println("Error al leer datos extraidos de la bd : " + err.Error())
			return u, err
		}
	} else {
		err = errors.New("usuario no existe")
	}
	return u, err

}
